package diff

import (
	"errors"
	"io"
	"os"

	"github.com/runatlantis/atlantis/server/core/config/raw"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func UnmarshalFile(filename string, v interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open yaml file: %v", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	decoder := yaml.NewDecoder(f)
	decoder.KnownFields(true)

	err = decoder.Decode(v)
	if errors.Is(err, io.EOF) {
		return nil
	}
	return err
}

func LoadRawRepoCfg(filename string) (*raw.RepoCfg, error) {
	rawCfg := &raw.RepoCfg{}
	err := UnmarshalFile(filename, rawCfg)
	if err != nil {
		return nil, err
	}
	err = rawCfg.Validate()
	if err != nil {
		return nil, err
	}
	return rawCfg, nil
}
