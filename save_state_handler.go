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

func (m mainModel) returnCurrentValues() *CurrentState {
  return &CurrentState{
    Method: m.url.chosenMethod,
    Url: m.url.textInput.Value(),
    RequestBody: m.request.body.Value(),
    RequestHeaders: m.request.headers.Value(),
    ResponseBody: m.rawResponse.rawResponse,
    ResponseHeaders: m.rawResponse.headers,
  }
}

func writeToFile(f *os.File, content *CurrentState) {
  byts, err := json.Marshal(content)
  if err != nil {
    log.Println("Error marshilling content to JSON: ", err)
    return 
  }
  _, err = f.WriteString(string(byts) + "\n")
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
  writeToFile(f, m.returnCurrentValues())
}


func (m mainModel) stateFileExists() bool {
  configDir, err := os.UserConfigDir()
  if err != nil {
    log.Println("Error finding the config dir: ", err)
    return false
  }
	savedStateFile := filepath.Join(configDir, "tgorq", "state.json")
  return fileExists(savedStateFile)
}


func (m mainModel) restorePreviousState() {
  configDir, err := os.UserConfigDir()
  if err != nil {
    log.Println("Error finding the config dir: ", err)
    return
  }
	filePath := filepath.Join(configDir, "tgorq")
	savedStateFile := filepath.Join(configDir, "tgorq", "state.json")

  // Check if the path exists. If not, create it.
  if !fileExists(filePath) {
      return
  }

  // Checking if the file exists and create it if not
  if !fileExists(savedStateFile) {
      return
  }

  jsonFile, err := os.Open(savedStateFile)
  if err != nil {
    log.Println("Error open the JSON file: ", err)
    return
  }
  defer jsonFile.Close()

  byteValue, _ := io.ReadAll(jsonFile)

  var currentState CurrentState
  if err := json.Unmarshal(byteValue, &currentState); err != nil {
    log.Println("Error unmarshilling the JSON file: ", err)
    return
  }
  log.Println("This is the current State variable: ", currentState)

  m.url.textInput.SetValue(currentState.Url)
  m.request.body.SetValue(currentState.RequestBody)
  m.request.headers.SetValue(currentState.RequestHeaders)
  m.response.body.SetContent(currentState.ResponseBody)
  m.response.headers.SetContent(currentState.ResponseHeaders)
}
