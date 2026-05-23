package cmd

import (
	"encoding/json"
	"io"
	"log"
)

func Read(r io.Reader) ([]Happiness, error) {
  var records []Happiness

  err := json.NewDecoder(r).Decode(&records)

  if err != nil {
    log.Fatal(err)
    return nil, err
  }

  return records, nil
}
