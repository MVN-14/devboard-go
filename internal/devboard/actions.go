package devboard

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
)

func AddProject(projectStr string) error {
	data, err := GetDevboardData()
	if err != nil {
		return err
	}

	project := Project{}
	err = json.Unmarshal([]byte(projectStr), &project)
	if err != nil {
		return err
	}

	if data.ContainsProject(project.Path) {
		return errors.New("Project path " + project.Path + " already exists in devboard. Did you mean to update?")
	}

	if project.Name == "" {
		return errors.New("Project is missing value \"name\"")
	} else if project.Path == "" {
		return errors.New("Project is missing value \"path\"")
	}

	info, err := os.Stat(project.Path)
	if !info.IsDir() {
		return errors.New("Project path " + project.Path + "does not point to directory.")
	}
	//if os.IsNotExist(err) {
	//	return errors.New("Project path \"" + project.Path + "\" does not exist")
	//}

	data.Projects = append(data.Projects, project)
	err = WriteDevboardData(data)
	if err != nil {
		return err
	}
	return nil
}

func ListProjects() (string, error) {
	data, err := GetDevboardData()
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(data.Projects)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func RemoveProject(path string) error {
	data, err := GetDevboardData()
	if err != nil {
		return err
	}

	projectIdx := slices.IndexFunc(data.Projects, func(v Project) bool {
		return v.Path == path
	})
	if projectIdx == -1 {
		return errors.New("Project with path " + path + " not found.")
	}

	data.Projects = slices.Concat(data.Projects[:projectIdx], data.Projects[projectIdx + 1:])
	fmt.Println("New Project List", data.Projects)

	return nil
}
