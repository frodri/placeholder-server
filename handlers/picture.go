// Copyright 2016 Francisco Rodriguez. All rights reserved.
// Use of this source code is governed by The MIT license.

package handlers

import (
    "log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/frodri/placeholder-server/util"
)



// PicGenerate is the handler responsible for generating the
// images. It parses the URL data into a pictureData object,
// which then takes care of returning an encodedData struct 
// (contains both a byte buffer containing the encoded image 
// and a content type) used for writing the handler's response. 
func PicGenerate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vars["text"] = r.URL.Query().Get("text")
	
	pd := util.BuildPictureData(vars)
	
	data, err := pd.Process()
	if err != nil {
		log.Println(err)
	}
	
	buffer := data.Buffer
	
	w.Header().Set("Content-Type", data.ContentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println(err)
	}
	
}
