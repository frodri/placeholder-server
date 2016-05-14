// Copyright 2016 Francisco Rodriguez. All rights reserved.
// Use of this source code is governed by The MIT license.

package handlers

import (
    "net/http"
)

// NotFound is your typical 404 error handler.
// Used to render the 404 error page from one of our templates.
func NotFound(w http.ResponseWriter, r *http.Request) {
    type Page struct {
        Title string
    }
    
    p := Page{
        Title: "404",
    }

    renderTemplate(w, "templates/notfound.html", p)
}