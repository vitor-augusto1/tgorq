package main

import (
	"log"
	"os"
	"path/filepath"
)

type CurrentState struct {
  method          httpMethod
  url             string
  requestBody     string
  requestHeaders  string
  responseBody    string
  responseHeaders string
}

func writeToFile(f *os.File, content string) {
  _, err := f.WriteString(content + "\n")
  if err != nil {
    log.Fatal(err)
  }
  if err := f.Sync(); err != nil {
    log.Println("Error syncing the file: ", err)
    return
  }
  defer f.Close()
}

