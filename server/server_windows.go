// Copyright 2016 Francisco Rodriguez. All rights reserved.
// Use of this source code is governed by The MIT license.

package server

import "github.com/spf13/cobra"

func init() {
// Help text that appears when placeholder-server is opened
// in Windows Explorer - a precaution adapted from 
// spf13's Hugo project
	cobra.MousetrapHelpText = `
placeholder-server is a configurable placeholder generator server.

As this is a command line application, you must use a command prompt to run it.
    `
}