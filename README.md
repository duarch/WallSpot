![heroicon](assets/wallspotog.png)

## WallSpot

![CodeQL Status](https://img.shields.io/github/workflow/status/duarch/WallSpot/CodeQL?style=flat-square) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/duarch/WallSpot?style=flat-square) ![GitHub](https://img.shields.io/github/license/duarch/WallSpot?style=flat-square)

Use **[Windows Spotlight](https://docs.microsoft.com/en-us/windows/configuration/windows-spotlight#:~:text=Windows%20Spotlight%20is%20an%20option,desktop%20editions%20of%20Windows%2010.)** images as workspace (background) wallpaper. This Windows 10 app recover and save all images from this Windows feature to use as slideshow within Windows theme personalization.

## Motivation
When I start my computer at Login Screen or at Lock Screen I stare at some beautiful pictures and I wish I could use it as Wallpaper. But it is not easily available. Surfing on internet we can find some solutions, but all of them are complicated and takes too much time to accomplish. It motivated me to create a program to do this complete task automatically.

## Tech/framework used

**Built with:**

* [Windows .theme file](https://docs.microsoft.com/en-us/windows/win32/controls/themesfileformat-overview)
* [Inno Setup](https://jrsoftware.org/isinfo.php)
* [Golang](https://go.dev/)

## Features
* Find Spotlight Images
* Separate them from others images
* Rename pictures with the date it loaded
* Copy them to dedicated folder on `/Pictures/WallSpot`
* Easily available to use as slideshow
* Simply delete the ones you don’t like
* Comes with an initial stock of samples to you don’t start from zero
* Not disturbing or modifying any system file
* Avoid duplicate images
* Everything with just one click (Actually don't need even to click!)


## Installation
Pick your option here in [releases](https://github.com/duarch/WallSpot/releases)
I recommend the installer option `WSsetup.exe`

## How to use?
Simply execute the file WallSpot.exe and then choose Slideshow on Windows personalization.

To get a desktop background (wallpaper), right-click on an empty space on workspace, then select *Personalize* it will open Settings Window on Background Options. Then follow these step just once:

**Setting up a desktop slideshow**

1. Using the "Background" drop-down menu, select the Slideshow option.
2. Click the Browse button to select the folder *...\Pictures\wallspot* with the picture collection.
3. Use the "Change picture every" drop-down menu and select how often images should rotate.
4. (Optional) Turn on the Shuffle toggle switch to show images in random order.
5. Using the "Choose a fit" drop-down menu, select the fit that best suits the images:
    (Obs.: All imagens are available in 1920 X 1080 pixels)
    * Fill. (Recommended)
    * Fit.
    * Stretch.
    * Center.
    * Span .

## Initial Samples
If you don't wan't to begin from "0" you can download some sample packs from [Google Drive WallSpot Vault](https://drive.google.com/drive/folders/119iPqPXfjNvandmJwMXryqtCixrfYeaP?usp=sharing). The program doesn't include any pictures. Each pack contains 100 files to make download easier. Eventually, packs receive new files updates.



## Credits
**Thanks to:**

- [David J. Malan](https://cs.harvard.edu/malan/)
- [Brian Yu](https://brianyu.me/)
- [CS50's staff](https://cs50.harvard.edu/college/2020/fall/staff/)
- [My lovely wife for her patience](https://www.instagram.com/clicks_of_mylife/)
- [pythoncentral.io](https://www.pythoncentral.io/finding-duplicate-files-with-python/)
- [Programação Dinâmica - Youtube Channel](https://www.youtube.com/c/Programa%C3%A7%C3%A3oDin%C3%A2mica)(Portuguese)

#### Anything else that seems useful
Plan to future implamentations:
- Add an icon on context menu
- Auto add Windows Theme

## License
MIT License

Copyright (c) 2020 Andre Duarte

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
