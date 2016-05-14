// Copyright 2016 Francisco Rodriguez. All rights reserved.
// Use of this source code is governed by The MIT license.

package handlers

import (
    "html/template"
    "net/http"
)

var templates *template.Template
const mainlayout string = "templates/layout.html"

// Let's just say writing the template boilerplate gets real annoying
// after a while.
func renderTemplate(w http.ResponseWriter, filename string, data interface{}) {
    templates = template.Must(template.ParseFiles(filename, mainlayout))
    err := templates.ExecuteTemplate(w, "base", data)
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}