package main

import (
	"os"

	"github.com/byteorbit/terragrunt-atlantis-config/internal/root"
)

// This variable is set at build time using -ldflags parameters.
// But we still set a default here for those using plain `go get` downloads
// For more info, see: http://stackoverflow.com/a/11355611/483528
var VERSION string = "1.23.1-BO"

func main() {
	rootCmd := root.New(VERSION)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
