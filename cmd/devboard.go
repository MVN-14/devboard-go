package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"os/exec"

	"github.com/MVN-14/devboard-go/internal/devboard"
	"github.com/docopt/docopt-go"
	_ "github.com/glebarez/go-sqlite"
)

const usage = `Devboard.

Usage: devboard (add | update) PROJECT
       devboard remove PATH
       devboard list
       devboard (--help | -h)
       devboard --version

Arguments:
    PATH       Path string to the project directory
    PROJECT    Project json string with the following properties:
                   name    - project name
                   path    - path to project directory

Options:
    -h --help    Show this screen.
    --version    Show version.
`

func main() {
	args, err := docopt.ParseArgs(usage, os.Args[1:], "1.0.0")
	if err != nil {
		fmt.Println("Error parsing usage:", err)
		return
	}
	
	devboardArgs := devboard.DevboardArgs{}
	args.Bind(&devboardArgs)
	
	if devboardArgs.Add {
		err := devboard.AddProject(devboardArgs.Project)
		if err != nil {
			panic(err)
		}
	} else if devboardArgs.List {
		projects, err := devboard.ListProjects()
		if err != nil {
			panic(err)
		}
		
		fmt.Println(projects)
	}

	// for k, v := range args {
	// 	fmt.Printf("%s = %v\n", k, v)
	// }

	// if len(os.Args) <= 1 || len(os.Args) > 1 && os.Args[1] == "help" {
	// 	return
	// }
	//
	// fmt.Println("Checking application files...")
	// err = devboard.CheckOrCreateApplicationDir()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Succesfully checked application files.")
	//
	//
	// switch os.Args[1] {
	// case "add":
	// 	if len(os.Args) < 3 {
	// 		printHelp()
	// 		break
	// 	}
	//
	// 	err := devboard.AddProject(os.Args[2])
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }
}

func connectDB() {
	db, err := sql.Open("sqlite", "./my.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	var sqliteVersion string
	err = db.QueryRow("select sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Connected to database with sqlite version %s", sqliteVersion)
}

func tmuxLs() {
	tmux, err := exec.LookPath("tmux")
	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command(tmux, "ls")
	if err != nil {
		fmt.Println(err)
	}

	var stdout, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err = cmd.Run()

	fmt.Printf("stdout: %s\nstderr: %s\n", stdout.String(), stderr.String())

}
