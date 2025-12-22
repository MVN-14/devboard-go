package devboard

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func CheckOrCreateApplicationDir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	basePath := filepath.Join(home, "." + ApplicationName)
	
	_, err = os.Stat(basePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("Creating project directory at", basePath)
			err = os.Mkdir(basePath, 0777)
			if err != nil {
				fmt.Println("Error creating directory at", basePath)
				return err
			}
			fmt.Println("Successfully created application directory at", basePath)
		}
	}
	
	dataFilePath := GetDataFilePath()
	_, err = os.Stat(dataFilePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("Creating application data file at", dataFilePath)
			file, err := os.Create(dataFilePath)
			if err != nil {
					fmt.Println("Error creating data file at", dataFilePath)
				return err
			}
			defer file.Close()
			file.WriteString("{}")
			fmt.Println("Successfully created data file at", dataFilePath)
		}
	}
	
	return nil
}

