// Copyright 2016 Francisco Rodriguez. All rights reserved.
// Use of this source code is governed by The MIT license.

package handlers

import (
    "net/http"
)


// Home renders the homepage for the site root.
// This page is used to provide instructions on how
// to use placeholder-server.
func Home(w http.ResponseWriter, r *http.Request) {
    type Page struct {
        Title string
        Host string
    }
    
    p := Page{
        Title: "home",
        Host: r.Host,
    }
    

    renderTemplate(w, "templates/home.html", p)
    
}