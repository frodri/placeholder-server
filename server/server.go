// Copyright 2016 Francisco Rodriguez. All rights reserved.
// Use of this source code is governed by The MIT license.

package server

import (
    "os"
    "log"
    "strconv"
    "net/http"
    
    "github.com/spf13/cobra"
    "github.com/gorilla/mux"
    
    "github.com/frodri/placeholder-server/util"
    "github.com/frodri/placeholder-server/handlers"
)

// ServerCmd is a cobra Root Command used to start the server.
// RunE is used instead of the standard Run because we want to catch
// things like font load errors before starting the server.
var ServerCmd = &cobra.Command{
	Use:   "placeholder-server",
	Short: "Holy Canoli! This server creates placeholders!",
	Long: `
placeholder-server is a configurable placeholder generator server.

Complete documentation is available at http://github.com/frodri/placeholder-server/
    `,
	RunE: func(cmd *cobra.Command, args []string) error {
		return launch()
	},
}

var port int
var fontfile string


// Execute is the launch point for the server. Once the flags are
// initialized, we launch the server using the data from those flags.
func Execute() {
    flagInit()
    
    if _, err := ServerCmd.ExecuteC(); err != nil {
		os.Exit(-1)
	}
}


func flagInit() {
    ServerCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port in which to host the server")
    ServerCmd.Flags().StringVar(&fontfile, "fontfile", "fonts/Roboto-Regular.ttf", "TrueType font file to use for the generated images.")
}

func launch() error {    
    
    // We make sure the font is loaded before starting the server.
    // This will catch things typos on the fontfile path along
    // with attempts to load unsupported fonts (e.g.: OpenType)
    if err := util.SetFont(fontfile); err != nil {
        return err
    }    
    
    r := mux.NewRouter()
    
    // This pattern matches one to four digits to determine width.
    // It then looks for two optional patterns - the first one being
    // height (so you can do things like 640x480), and the second one
    // being a file extension like .gif or .jpeg.
    dimensionpattern := "\\d{1,4}(x\\d{1,4})?(\\.\\w+)?";
    // This pattern accepts CSS-style hex color notation (e.g. "0000FF"
    // or "00F" would represent "blue") along woth an optional extension.
    colorpattern := "([\\da-f]{6}|[\\da-f]{3})(\\.\\w+)?"
    
    // Http routes. We use the regexes we declare above to build 'em.
    // We do break them down further once we reach the http.Handlers.
    r.HandleFunc("/", handlers.Home)
    r.HandleFunc("/{dimensions:"+dimensionpattern+"}", handlers.PicGenerate)
    r.HandleFunc("/{dimensions:"+dimensionpattern+"}/{bgcolor:"+colorpattern+"}", handlers.PicGenerate)
    r.HandleFunc("/{dimensions:"+dimensionpattern+"}/{bgcolor:"+colorpattern+"}/{fgcolor:"+colorpattern+"}", handlers.PicGenerate)
    r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

    if port > 0 && port <= 65535 {
        log.Println("Listening on port", port, "- Press Ctrl-C to close the server")    
    }
    
    return http.ListenAndServe(":"+strconv.Itoa(port), r)
}