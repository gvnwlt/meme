package main

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "github.com/jteeuwen/go-bindata"
	"github.com/nomad-software/meme/cli"
	"github.com/nomad-software/meme/image"
	"github.com/nomad-software/meme/output"
)

func main() {
	opt := cli.ParseOptions()

	if opt.Help {
		opt.PrintUsage()

	} else if opt.Valid() {
		st := image.Load(opt)
		st = image.RenderImage(opt, st)

		if opt.ClientID != "" {
			url := image.Upload(opt, st)
			output.Info(url)
		} else {
			file := image.Save(opt, st)
			output.Info(file)
		}
	}
}
