package halftone

import (
	"context"
	"flag"
	"fmt"
	"log"

	_ "github.com/aaronland/go-image-halftone"
	"github.com/aaronland/go-image/app/transform"
	"github.com/sfomuseum/go-flags/flagset"
)

func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	resize_uri := fmt.Sprintf("halftone://?process=%s&scale-factor=%d", process, scale_factor)
	suffix := fmt.Sprintf("-%s-%d", process, scale_factor)

	transformation_uris := []string{
		resize_uri,
	}

	for _, e := range extra_transformations {
		transformation_uris = append(transformation_uris, e)
	}

	opts := &transform.RunOptions{
		TransformationURIs: transformation_uris,
		ApplySuffix:        suffix,
		SourceURI:          source_uri,
		TargetURI:          target_uri,
		Logger:             logger,
	}

	paths := fs.Args()

	return transform.RunWithOptions(ctx, opts, paths...)
}
