package devboard

import (
	"database/sql"
	"errors"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

type Board struct {
	db *sql.DB
}

const createTableStmt = `CREATE TABLE IF NOT EXISTS projects (
    command 	TEXT,
	id 			INTEGER PRIMARY KEY,
	name 		TEXT NOT NULL,
	path 		TEXT NOT NULL,
	created_at	DATETIME NOT NULL DEFAULT (datetime('now', 'localtime')),
	updated_at	DATETIME NOT NULL DEFAULT (datetime('now', 'localtime')),
	deleted_at	DATETIME 
);`

func New(dbpath string) (*Board, error) {
	db, err := sql.Open("sqlite", dbpath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(createTableStmt)
	if err != nil {
		return nil, err
	}

	board := &Board{db}
	return board, nil
}

func (b *Board) Open(p Project) error {
	cmdStr := viper.GetString("command")
	if cmdStr == "" {
		cmdStr = p.Command
	}
	if cmdStr == "" {
		return errors.New("Error, no command provided to open project")
	}

	cmdStr = strings.ReplaceAll(cmdStr, "$projPath", p.Path)
	fields := strings.Fields(cmdStr)
	prog := fields[0]
	
	var cmd *exec.Cmd
	if len(fields) == 1 {
		cmd = exec.Command(prog)
	} else {
		args := fields[1:]
		cmd = exec.Command(prog, args...)
	}
	cmd.Env = append(cmd.Environ(), "projPath=" + p.Path)
	
	if err := cmd.Start(); err != nil {
		return err
	}
	
	return nil
}

func (b *Board) Update(p Project) error {
	_, err := b.db.Exec(`
		UPDATE projects 
		SET 
			name = ?,
			path = ?,
			command = ?,
			updated_at = (datetime('now', 'localtime'))
		WHERE 
			id = ?;`, p.Name, p.Path, p.Command, p.Id)

	if err != nil {
		return err
	}

	return nil
}

func (b *Board) List(deleted bool) (ProjectList, error) {
	projects := ProjectList{}

	sql := "SELECT * FROM projects"
	if !deleted {
		sql += " WHERE deleted_at IS NULL"
	}

	rows, err := b.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err := projects.FromRows(rows); err != nil {
		return nil, err
	}

	return projects, nil
}

func (b *Board) Get(id int) (Project, error) {
	p := Project{}
	
	row := b.db.QueryRow(`SELECT * FROM projects WHERE id = ?`, id)
	err := p.FromRow(row)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (b *Board) Add(p Project) error {
	_, err := b.db.Exec(
		`INSERT INTO projects (name, path, command) VALUES (?, ?, ?)`,
		p.Name, p.Path, p.Command)

	if err != nil {
		return err
	}
	return nil
}

func (b *Board) Delete(id int) error {
	_, err := b.db.Exec(`UPDATE projects SET deleted_at = (datetime('now', 'localtime')) WHERE id = ?; `, id)
	if err != nil {
		return err
	}
	return nil
}

func (b *Board) Close() {
	b.db.Close()
}
