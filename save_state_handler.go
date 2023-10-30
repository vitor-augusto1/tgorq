package main

import (
	"encoding/json"
	"io"
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
		Method:          m.url.chosenMethod,
		Url:             m.url.textInput.Value(),
		RequestBody:     m.request.body.Value(),
		RequestHeaders:  m.request.headers.Value(),
		ResponseBody:    m.rawResponse.body,
		ResponseHeaders: m.rawResponse.headers,
	}
}

func writeToFile(f *os.File, content *CurrentState) {
	bytes, err := json.MarshalIndent(content, "", " ")
	if err != nil {
		log.Println("Error marshilling content to JSON: ", err)
		return
	}
	_, err = f.Write(bytes)
	if err != nil {
		log.Println("Error writing content to JSON file: ", err)
		return
	}
	if err := f.Sync(); err != nil {
		log.Println("Error syncing the file: ", err)
		return
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func (m mainModel) storeCurrentState() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Println("Error finding the config dir: ", err)
		return
	}

	filePath := filepath.Join(userConfigDir, "tgorq")
	savedStateFile := filepath.Join(userConfigDir, "tgorq", "state.json")

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
	defer f.Close()
	fstat, _ := f.Stat()
	if fstat.Size() != 0 {
		err = f.Truncate(0)
		if err != nil {
			log.Println("Error truncating the file: ", err)
			return
		}
		_, err = f.Seek(0, 0)
		if err != nil {
			log.Println("Error seeking the file: ", err)
			return
		}
	}
	writeToFile(f, m.returnCurrentValues())
}

func (m mainModel) stateFileExists() bool {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Println("Error finding the config dir: ", err)
		return false
	}
	savedStateFile := filepath.Join(userConfigDir, "tgorq", "state.json")
	return fileExists(savedStateFile)
}

func (m mainModel) restorePreviousState() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Println("Error finding the config dir: ", err)
		return
	}
	filePath := filepath.Join(userConfigDir, "tgorq")
	savedStateFile := filepath.Join(userConfigDir, "tgorq", "state.json")

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

	m.url.httpMethodPaginator.Page = int(currentState.Method)
	m.url.chosenMethod = currentState.Method
	m.url.textInput.SetValue(currentState.Url)
	m.request.body.SetValue(currentState.RequestBody)
	m.request.headers.SetValue(currentState.RequestHeaders)
	m.response.body.SetContent(currentState.ResponseBody)
	m.response.headers.SetContent(currentState.ResponseHeaders)
}
