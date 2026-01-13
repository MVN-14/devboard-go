package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MVN-14/devboard-go/internal/devboard"
	"github.com/docopt/docopt-go"
)

func HandleError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	args, err := docopt.ParseArgs(devboard.Usage, os.Args[1:], "1.0.0")
	if err != nil {
		fmt.Println("Error parsing usage:", err)
		return
	}

	err = devboard.InitDB()
	if err != nil {
		HandleError(err.Error())
	}

	devboardArgs := devboard.DevboardArgs{}
	args.Bind(&devboardArgs)

	if devboardArgs.List {
		projects, err := devboard.GetProjects()
		if err != nil {
			HandleError(err.Error())
		}
		bytes, err := json.Marshal(projects)
		if err != nil {
			HandleError(err.Error())
		}
		fmt.Println(string(bytes))
	} else if devboardArgs.Add {
		var project devboard.DBProject
		err := json.Unmarshal([]byte(devboardArgs.Project), &project)
		if err != nil {
			HandleError(err.Error())
		}

		err = devboard.AddProject(project)
		if err != nil {
			HandleError(err.Error())
		}
		fmt.Println("Successfully added project", project.Name, "at path", project.Path)
	} else if devboardArgs.Remove {
		err := devboard.DeleteProject(devboardArgs.Path)
		if err != nil {
			HandleError(err.Error())
		}
		fmt.Println("Deleted project at path", devboardArgs.Path)
	} else if devboardArgs.Update {
		var project devboard.DBProject
		err := json.Unmarshal([]byte(devboardArgs.Project), &project)
		if err != nil {
			HandleError(err.Error())
		}
			
		err = devboard.UpdateProject(devboardArgs.Id, project)
		if err != nil {
			HandleError(err.Error())
		}
	}
}

// func tmuxLs() {
// 	tmux, err := exec.LookPath("tmux")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
//
// 	cmd := exec.Command(tmux, "ls")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
//
// 	var stdout, stderr bytes.Buffer
// 	cmd.Stderr = &stderr
// 	cmd.Stdout = &stdout
//
// 	err = cmd.Run()
//
// 	fmt.Printf("stdout: %s\nstderr: %s\n", stdout.String(), stderr.String())
//
// }
