**GOWAL**

**gowal: A Wallpaper Color Extractor**
=====================================

**Overview**

gowal is a command-line tool that extracts dominant colors from a wallpaper image file and saves them to a file. This tool is built using Go and uses the `image/jpeg` and `image/png` packages to handle image processing.

**Installation**

To install Dynamica, follow these steps:

1. **Install Go**: You need to have Go installed on your system. You can download the latest version from the official Go website: <https://golang.org/dl/>
2. **Get the code**: Clone this repository using Git: `git clone https://github.com/amsynist/gowal.git` && `cd gowal`
3. **Install the Dependencies**: `go install`
3. **Build the binary**: Navigate to the cloned repository and run: `go build main.go`

**Usage**

To use Gowal, run the following command:
```
./gowal -file <path-to-wallpaper-file> -size <color-pallete-size> -out <color-pallete-output-path> 
```
**Flags**

* `-file`: The path to the wallpaper image file (required)
* `-size`: The color pallete size (n of colors) (optional, default is 8)
* `-out`: The output path for pallete (optional, default is `colors`)


**Example**

Extract dominant colors from a wallpaper image file `wallpaper.jpg` with an intensity of 12:
```
./gowal -file wallpaper.jpg -size 12 -out output_file/path
```
This will save the extracted colors to a file.



If you encounter any issues or have feedback, please create an issue in this repository.
