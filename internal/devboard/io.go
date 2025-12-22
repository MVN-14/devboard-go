package devboard

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var homeDirectory string

func GetHomeDirectory() string {
	if homeDirectory != "" {
		return homeDirectory
	}

	home, err := os.UserHomeDir()
	if err != nil {
		panic("Could not get user home directory")
	}

	homeDirectory = home
	return home
}

func GetDataFilePath() string {
	return filepath.Join(GetHomeDirectory(), ApplicationDirName, DataFileName)
}

func GetDevboardData() (DevboardData, error) {
	var applicationData DevboardData
	dataFilePath := GetDataFilePath()

	file, err := os.Open(dataFilePath)
	if err != nil {
		fmt.Println("Error reading application data file at", GetDataFilePath())
		return applicationData, err
	}
	defer file.Close()
	
	err = json.NewDecoder(file).Decode(&applicationData)
	if err != nil {
		fmt.Println("Error parsing json file at", GetDataFilePath())
		return applicationData, err
	}
	
	return applicationData, nil
}

func WriteDevboardData(data DevboardData) error {
	dataFilePath := GetDataFilePath()

	file, err := os.OpenFile(dataFilePath, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()
	
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}

	return nil
}
