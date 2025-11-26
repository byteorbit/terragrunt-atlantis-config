package main

import (
	"log"

	"github.com/transcend-io/terragrunt-atlantis-config/internal/diff"
)

func main() {
	cmd := diff.New()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
