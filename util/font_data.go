// Copyright 2016 Francisco Rodriguez. All rights reserved.
// Use of this source code is governed by The MIT license.

package util

import (
    "errors"
    "io/ioutil"
    "github.com/golang/freetype/truetype"
)

var fontfile *truetype.Font

// SetFont is called on server launch to load in the
// font used by the pd.Process calls used during 
// image generation.
func SetFont(loc string) error {
    if fontfile != nil {
        return errors.New("Font already set")
    }
    
    fontBytes, err := ioutil.ReadFile(loc);
	if err != nil {
		return err
	}
    
    fontfile, err = truetype.Parse(fontBytes)
	if err != nil {
		return err
	}
    
    return nil
}
