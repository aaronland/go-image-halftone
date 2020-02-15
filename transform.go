package halftone

import (
	"context"
	"github.com/aaronland/go-image-transform"
	"image"
	_ "net/url"
)

type HalftoneTransformation struct {
	transform.Transformation
	options *HalftoneOptions
}

func init() {

	ctx := context.Background()
	err := transform.RegisterTransformation(ctx, "Halftone", NewHalftoneTransformation)

	if err != nil {
		panic(err)
	}
}

func NewHalftoneTransformation(ctx context.Context, str_url string) (transform.Transformation, error) {

	opts := NewDefaultHalftoneOptions()

	tr := &HalftoneTransformation{
		options: opts,
	}

	return tr, nil
}

func (tr *HalftoneTransformation) Transform(ctx context.Context, im image.Image) (image.Image, error) {
	return HalftoneImage(ctx, im, tr.options)
}
