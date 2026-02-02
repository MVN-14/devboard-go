package devboard

import (
	"database/sql"
	"encoding/json"
	"slices"
	"time"
)

type Project struct {
	Command   string     `json:"command"`
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (p *Project) FromRow(r *sql.Row) error {
	r.Scan(&p.Command, &p.Id, &p.Name, &p.Path, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}

type ProjectList []Project

func (pl ProjectList) String() string {
	b, _ := json.Marshal(pl)
	return string(b)
}

func (pl *ProjectList) FromRows(r *sql.Rows) error {
	cols, _ := r.Columns()
	deleted := false
	if slices.Contains(cols, "deleted_at") {
		deleted = true
	}

	for r.Next() {
		p := Project{}
		var err error
		if deleted {
			err = r.Scan(&p.Command, &p.Id, &p.Name, &p.Path, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
		} else {
			err = r.Scan(&p.Command, &p.Id, &p.Name, &p.Path, &p.CreatedAt, &p.UpdatedAt)
		}
		if err != nil {
			return err
		}
		*pl = append(*pl, p)
	}
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}
