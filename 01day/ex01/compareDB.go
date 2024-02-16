package main

import (
  "os"
  "fmt"
  "ex01/DB"
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

var fileNameOld, fileNameNew filename

func init() {
  flag.Var(&fileNameOld, "old", "A string. Set filename json or xml extension for DB")
  flag.Var(&fileNameNew, "new", "A string. Set filename json or xml extension for DB")
}

func main() {
  flag.Parse()
  if flag.NArg() != 0 {
    fmt.Fprintln(os.Stderr,
                 "No arguments are expected except for the --old and --new option")
    flag.PrintDefaults()
    return
  } else if len(fileNameNew) != len(fileNameOld) {
    fmt.Fprintln(os.Stderr, "Wrong number of flags: expected the same number of --old and --new arguements")
    flag.PrintDefaults()
    return
  }

  for i := 0; i < len(fileNameNew); i++ {
    var old, fresh DB.File
    old.Filename = fileNameOld[i]
    fresh.Filename = fileNameNew[i]
    oldBook, err := old.Read()
    if err != nil {
      fmt.Fprintf(os.Stderr, "%s: %s\n", old.Filename, err)
      return
    }
    freshBook, err := fresh.Read()
    if err != nil {
      fmt.Fprintf(os.Stderr, "%s: %s\n", &fresh.Filename, err)
      return
    }
    DB.CompareBook(oldBook, freshBook)
    // fmt.Fprintln(os.Stdout, strings.Repeat("-", 20))
  }
}
