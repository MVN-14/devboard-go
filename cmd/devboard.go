package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"os/exec"

	"github.com/MVN-14/devboard-go/internal/initialize"
	"github.com/MVN-14/devboard-go/internal/io"
	_ "github.com/glebarez/go-sqlite"
)

type DevboardJSON struct {
	Projects Project `json:"projects"`
}

type Project struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func main() {
	fmt.Println("Checking application files...")
	err := initialize.CheckOrCreateApplicationDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Succesfully checked application files.")
	
	fmt.Println("Loading application data")
	data, err := io.GetApplicationData()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("data: %+v", data)
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
