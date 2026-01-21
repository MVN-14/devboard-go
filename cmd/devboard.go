package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MVN-14/devboard-go/internal/devboard"
)

func HandleError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	board, err := devboard.NewDevboard(&devboard.SqliteStore{})
	if err != nil {
		HandleError(err.Error())
	}
	defer board.Close()

	devboardArgs := devboard.ParseArgs(os.Args[1:])

	if devboardArgs.List {
		projects, err := board.GetProjects()
		if err != nil {
			HandleError(err.Error())
		}
		bytes, err := json.Marshal(projects)
		if err != nil {
			HandleError(err.Error())
		}
		fmt.Println(string(bytes))

	} else if devboardArgs.Add {
		var project devboard.Project
		err := json.Unmarshal([]byte(devboardArgs.Project), &project)
		if err != nil {
			HandleError(err.Error())
		}

		ok, err := board.AddProject(project)
		if err != nil {
			HandleError(err.Error())
		}
		if ok {
			fmt.Println("Successfully added project", project.Name, "at path", project.Path)
		} else {
			fmt.Println("Something went wrong")
		}

	} else if devboardArgs.Remove {
		ok, err := board.DeleteProject(devboardArgs.Id)
		if err != nil {
			HandleError(err.Error())
		}
		if ok {
			fmt.Println("Successfully deleted project with id", devboardArgs.Id)
		} else {
			fmt.Println("Something went wrong")
		}

	} else if devboardArgs.Update {
		var project devboard.Project
		err := json.Unmarshal([]byte(devboardArgs.Project), &project)
		if err != nil {
			HandleError(err.Error())
		}

		ok, err := board.UpdateProject(project)
		if err != nil {
			HandleError(err.Error())
		}
		if ok {
			fmt.Println("Successfully deleted project with id", devboardArgs.Id)
		} else {
			fmt.Println("Something went wrong")
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
