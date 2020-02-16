package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-image-cli"
	"github.com/aaronland/go-image-halftone"
	"image"
	"log"
	"path/filepath"
	"strings"
)

func main() {

	mode := flag.String("mode", "atkinson", "...")
	scale := flag.Float64("scale-factor", 2.0, "...")

	flag.Parse()

	ctx := context.Background()

	opts := halftone.NewDefaultHalftoneOptions()

	opts.Mode = *mode
	opts.ScaleFactor = *scale

	cb := func(ctx context.Context, im image.Image, path string) (image.Image, string, error) {

		new_im, err := halftone.HalftoneImage(ctx, im, opts)

		if err != nil {
			return nil, "", err
		}

		root := filepath.Dir(path)

		fname := filepath.Base(path)
		ext := filepath.Ext(path)

		short_name := strings.Replace(fname, ext, "", 1)

		scale_label := fmt.Sprintf("%f", opts.ScaleFactor)
		scale_label = strings.TrimRight(scale_label, "0")
		scale_label = strings.TrimRight(scale_label, ".")

		label := fmt.Sprintf("halftone-%s-%s", opts.Mode, scale_label)

		new_name := fmt.Sprintf("%s-%s%s", short_name, label, ext)

		new_path := filepath.Join(root, new_name)

		log.Println(new_path)
		return new_im, new_path, nil
	}

	paths := flag.Args()

	err := cli.Process(ctx, cb, paths...)

	if err != nil {
		log.Fatal(err)
	}

}
