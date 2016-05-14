// Copyright 2016 Francisco Rodriguez. All rights reserved.
// Use of this source code is governed by The MIT license.

package util

import (
    "image"
    "image/draw"
    "image/color"
    "image/jpeg"
	"image/gif"
	"image/png"
    "bytes"
    
    "golang.org/x/image/font"
    "golang.org/x/image/math/fixed"
    
    "github.com/golang/freetype/truetype"
    "github.com/lucasb-eyer/go-colorful"
)

type encodedData struct {
    ContentType string
    Buffer *bytes.Buffer
}

func (pd *pictureData) Process() (data encodedData, err error)  {
    
    // Code here is derived from the "drawer" example in the
    // golang/freetype package.
    
    face := truetype.NewFace(fontfile, &truetype.Options{
		Size:   float64(pd.FontSize),
		DPI:     72,
		Hinting: font.HintingNone,
	})
    
    // The one major difference from the freetype example
    // comes from the use of the go-colorful library to create
    // the color.Color implementing objects required by 
    // image.NewUniform. Note that we use a var declaration 
    // instead of := because we want to actually avoid the 
    // colorful.Color cast the := operator performs.
    var BGcolor color.Color
    var FGcolor color.Color
    
    BGcolor, _ = colorful.Hex("#"+pd.BGcolor)
    FGcolor, _ = colorful.Hex("#"+pd.FGcolor)
    
    BGimage := image.NewUniform(BGcolor)
    FGimage := image.NewUniform(FGcolor)
    
    rgba := image.NewRGBA(image.Rect(0, 0, pd.Width, pd.Height))
    draw.Draw(rgba, rgba.Bounds(), BGimage, image.ZP, draw.Src)
    
    d := &font.Drawer{
		Dst: rgba,
		Src: FGimage,
		Face: face,
	}
    
	d.Dot = fixed.Point26_6{
		X: (fixed.I(pd.Width) - d.MeasureString(pd.Text)) / 2,
		Y: fixed.I(pd.Height) / 2 ,
	}
	d.DrawString(pd.Text)
   
   // Image encoding starts here.
   buffer := new(bytes.Buffer)

   switch pd.ImgType {
	case ".gif":
		err = gif.Encode(buffer, rgba, nil)
		data.ContentType = "image/gif"
	case ".jpg", "jpeg":
		err = jpeg.Encode(buffer, rgba, nil)
		data.ContentType = "image/jpeg"
	default:
		err = png.Encode(buffer, rgba)
        data.ContentType = "image/png"
	}
    
    data.Buffer = buffer
    
    return data, err
}