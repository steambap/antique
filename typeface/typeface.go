package typeface

import (
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

const (
	smSize = 14
	mdSize = 24
)

var (
	GoRegular14 font.Face
	GoMedium24  font.Face
)

func init() {
	tt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal("Error parsing go font.")
	}

	GoRegular14 = truetype.NewFace(tt, &truetype.Options{
		Size:    smSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	GoMedium24 = truetype.NewFace(tt, &truetype.Options{
		Size:    mdSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}
