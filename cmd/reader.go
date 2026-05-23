package cmd

import (
	"encoding/json"
	"io"
)

func Read(r io.Reader) ([]Happiness, error) {
  var records []Happiness

  err := json.NewDecoder(r).Decode(&records)

  if err != nil {
    return nil, err
  }

  return records, nil
}
