package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type CurrentState struct {
  Method          httpMethod `json:"method"`
  Url             string     `json:"url"`
  RequestBody     string     `json:"request_body"`
  RequestHeaders  string     `json:"request_headers"`
  ResponseBody    string     `json:"response_body"`
  ResponseHeaders string     `json:"response_headers"`
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

func fileExists(path string) bool {
  _, err := os.Stat(path)
  if os.IsNotExist(err) {
    return false
  }
  return err == nil
}

func (m mainModel) storeCurrentState() {
  configDir, err := os.UserConfigDir()
  if err != nil {
    log.Println("Error finding the config dir: ", err)
    return
  }

	filePath := filepath.Join(configDir, "tgorq")
	savedStateFile := filepath.Join(configDir, "tgorq", "state.json")

  // Check if the path exists. If not, create it.
  if !fileExists(filePath) {
    err := os.Mkdir(filePath, 0755)
    if err != nil {
      log.Println("Error making directory: ", err)
      return
    }
  }

  // Checking if the file exists and create it if not
  if !fileExists(savedStateFile) {
    f, err := os.Create(savedStateFile)
    if err != nil {
      log.Println("Error creating the state file: ", err)
      return
    }
    defer f.Close()
  }

  f, err := os.OpenFile(savedStateFile, os.O_SYNC|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
   log.Println(err)
  }
  writeToFile(f, "test")
}
