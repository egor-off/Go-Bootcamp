package DB

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

func (reader File) readJSON() (*CookBook, error) {
  data, err := os.ReadFile(reader.Filename)
  if err != nil {
    return nil, err
  }

  var cookbook CookBook
  if err = json.Unmarshal(data, &cookbook); err != nil {
    fmt.Fprintln(os.Stderr, err)
    return nil, err
  }

  return &cookbook, nil
}

func (reader File) readXML() (*CookBook, error) {
  data, err := os.ReadFile(reader.Filename)
  if err != nil {
    return nil, err
  }

  var cookbook CookBook
  if err = xml.Unmarshal(data, &cookbook); err != nil {
    fmt.Fprintln(os.Stderr, err)
    return nil, err
  }

  return &cookbook, nil
}

func (reader File) Read() (*CookBook, error) {
  if strings.HasSuffix(reader.Filename, ".json") {
    return reader.readJSON()
  } else {
    return reader.readXML()
  }
}
