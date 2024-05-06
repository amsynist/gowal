package utils

import (
	"fmt"
	"image"
	"os"

	colorkit "github.com/gookit/color"

	"github.com/cenkalti/dominantcolor"
)

func ExtractDomiantColors(fileInput string, count int) ([]string, error) {
	f, err := os.Open(fileInput)
	defer f.Close()
	if err != nil {
		fmt.Println("File not found:", fileInput)
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	var dominantColors []string
	for _, color := range dominantcolor.FindN(img, count) {
		dominantColors = append(dominantColors, dominantcolor.Hex(color))
	}
	return dominantColors, nil
}
func ExtractColors(fileInput string, count int) []string {
	fallBackColors := []string{"#ffffff", "#000000", "00ffff", "0f0f0f", "#ffffff", "#000000", "00ffff", "0f0f0f"}
	colors, err := ExtractDomiantColors(fileInput, count)
	if err != nil {
		colors = fallBackColors
	}
	if len(colors) != count {
		colors = append(colors, fallBackColors...)

	}
	return colors[:count]
}

func SaveColors(colorList []string) string {
	// Create or open the file named "colors"
	file, err := os.Create("colors")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return file.Name()
	}
	defer file.Close()

	// Write each string to the file with the naming convention COLOR_ONE, COLOR_TWO, etc.
	file.WriteString("#!/bin/bash \n")
	fmt.Println("━━━━━━━━━━Color Sample Pallete━━━━━━━━━━")
	for i, color := range colorList {
		colorkit.Hex(color).Printf("┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃┃\n")
		line := fmt.Sprintf("COLOR_%d=0xff%s\n", i+1, color[1:])
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}
	fmt.Println("\n━━━━━━━━━━Color Sample Pallete━━━━━━━━━━")

	return file.Name()

}
