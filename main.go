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
				Name:  "intensity",
				Value: "8",
				Usage: "theming intensity",
			},
		},
		Action: func(cCtx *cli.Context) error {
			intensity, err := strconv.Atoi(cCtx.String("intensity"))
			if err != nil {
				fmt.Printf("intensity=%d, type: %T\n", intensity, intensity)
			}

			if cCtx.String("file") != "" {
				colors := utils.ExtractColors(cCtx.String("file"), intensity)
				colors_file := utils.SaveColors(colors)
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
