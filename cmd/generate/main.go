package main

import (
	"log"

	"github.com/byteorbit/terragrunt-atlantis-config/internal/generate"
)

func main() {
	cmd := generate.New()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
