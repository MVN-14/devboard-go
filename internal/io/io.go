package io

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MVN-14/devboard-go/internal/constants"
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
	return filepath.Join(GetHomeDirectory(), constants.ApplicationDirName, constants.DataFileName)
}

func GetApplicationData() (ApplicationData, error) {
	var applicationData ApplicationData
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


