package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strconv"

	"github.com/amsynist/gowal/utils"

	"github.com/urfave/cli"
)

func main() {

	app := &cli.App{
		Name: "dynamica",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Value: "",
				Usage: "path to the wallpaper file",
			},
			&cli.StringFlag{
				Name:  "size",
				Value: "8",
				Usage: "number of colors to generate from the image",
			},
			&cli.StringFlag{
				Name:  "out",
				Value: "colors",
				Usage: "color pallete output path",
			},
		},
		Action: func(cCtx *cli.Context) error {
			intensity, err := strconv.Atoi(cCtx.String("size"))
			palleteDest := cCtx.String("out")

			if err != nil {
				fmt.Printf("pallete-size=%d, type: %T\n", intensity, intensity)
			}

			if cCtx.String("file") != "" {
				colors, err := utils.ExtractColors(cCtx.String("file"), intensity)
				if err != nil {
					log.Fatal("Error extracting colors: ", err)
				}
				colors_file, err := utils.SaveColors(colors, palleteDest)
				if err != nil {
					log.Fatal("Error saving colors: ", err)
				}
				fmt.Printf("Saved colors with filename : %s", colors_file)

			} else {
				log.Fatal("Missing required flags : file")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
