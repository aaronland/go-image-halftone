package main

import (
	"context"
	"log"

	"github.com/aaronland/go-image-halftone/app/halftone"
	_ "github.com/aaronland/go-image/common"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := halftone.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to halftone images, %v", err)
	}
}
