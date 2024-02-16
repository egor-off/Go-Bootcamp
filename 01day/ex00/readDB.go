package main

import (
  "os"
  "fmt"
  "ex00/DB"
  "flag"
  "strings"
  "errors")

type filename []string

func (f *filename) String() string {
  return fmt.Sprint(*f)
}

func (f *filename) Set(value string) error {
  if !strings.HasSuffix(value, ".xml") &&
     !strings.HasSuffix(value, ".json") {
    return errors.New("expected xml or json extension")
  }
  *f = append(*f, value)
  return nil
}

var filenameFlag filename

func init() {
  flag.Var(&filenameFlag, "f",
           "A string. Set filename json or xml extension for DB")
}

func main() {
  flag.Parse()
  if flag.NArg() != 0 {
    fmt.Fprintln(os.Stderr,
                 "No arguments are expected except for the -f option")
    flag.PrintDefaults()
    return
  }

  for _, f := range filenameFlag {
    var reader DB.File
    reader.Filename = f
    cookbook, err := reader.Read()
    if err != nil {
      fmt.Fprintf(os.Stderr, "%s: %s\n", f, err)
      return
    }
    if cookbook != nil {
      fmt.Println(*cookbook)
    }

    if err := cookbook.Write("xml"); err != nil {
      fmt.Fprintln(os.Stderr, err)
    }
    if err := cookbook.Write("json"); err != nil {
      fmt.Fprintln(os.Stderr, err)
    }
    if err := cookbook.Write("asasa"); err != nil {
      fmt.Fprintln(os.Stderr, err)
    }
  }
}
