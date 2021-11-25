package main

import (
	"fmt"
	"github.com/Navid2zp/dups"
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	var dir = os.Getenv("USERPROFILE") + "\\AppData\\Local\\Packages\\Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy\\LocalState\\Assets"
	var dirWallpapers = os.Getenv("USERPROFILE") + "\\Pictures\\WallSpot"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("Diretório não existe")
		os.Exit(1)
	}
	if _, err := os.Stat(dirWallpapers); os.IsNotExist(err) {
		os.Mkdir(dirWallpapers, 0777)
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	i := 1
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		item := filepath.Join(dir, file.Name())
		src := dir + "\\" + file.Name()
		f, _ := os.Open(item)
		img, _ := jpeg.Decode(f)
		if img != nil {
			if img.Bounds().Dx() == 1920 && img.Bounds().Dy() == 1080 {
				newName := "w" + time.Now().Format("020106") + "_" + strconv.Itoa(i) + ".jpg"
				dst := dirWallpapers + "\\" + newName
				copyFile(src, dst)
				i++
			}
		}
	}
	clean(dirWallpapers)
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

//Based on README example from "github.com/Navid2zp/dups"
func clean(dir string) {
	files, err := dups.GetFiles(dir, true)
	if err != nil {
		panic(err)
	}
	minSize := 1
	hashes := dups.CollectHashes(files, false, minSize, dups.XXHash, true)
	duplicates, _, _ := dups.GetDuplicates(hashes)
	_, _, error := dups.RemoveDuplicates(duplicates)
	if error != nil {
		panic(error)
	}
}
