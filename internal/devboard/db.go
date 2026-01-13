package devboard

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBProject struct {
	gorm.Model
	Name      string
	Path      string `gorm:"uniqueIndex"`
}

var db *gorm.DB
var ctx = context.Background()

func InitDB() error {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	applicationDir := path.Join(home, ApplicationDirName)

	_, err = os.Stat(applicationDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.Mkdir(applicationDir, 0755)
			if err != nil {
				panic(err)
			}
			fmt.Println("Created application directory at ", applicationDir)
		} else {
			return err
		}
	}

	dbPath := path.Join(home, ApplicationDirName, DbName)
	fmt.Println(dbPath)
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&DBProject{})
	if err != nil {
		return err
	}

	return nil
}

func PathExists(path string) (bool, error) {
	table := gorm.G[DBProject](db)
	count, err := table.Where("path = ?", path).Count(ctx, "path")
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

func AddProject(project DBProject) error {
	if project.Path == "" {
		return errors.New("Project is missing 'path' value")
	}
	if project.Name == "" {
		return errors.New("Project is missing 'name' value")
	}
	
	exists, err := PathExists(project.Path)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("Project with path " + project.Path + " already exists.")
	}

	err = gorm.G[DBProject](db).Create(ctx, &project)
	
	return err
}

func UpdateProject(id uint, newProject DBProject) error {
	// project, err := gorm.G[DBProject](db).Where("id = ?", id).First(ctx)
	// if err != nil {
	// 	return err
	// }
	
	rows, err := gorm.G[DBProject](db).Where("id = ?", id).Updates(ctx, newProject)
	if err != nil {
		return err
	}
	if rows <= 0 {
		return errors.New("No rows affected by update")
	}
	
	// if newProject.Path != "" && newProject.Path != project.Path {
	// 	exists, err := PathExists(newProject.Path)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if exists {
	// 		return errors.New("Can't update path to " + newProject.Path + ". A project already exists with that path.")
	// 	}
	// 	project.Path = newProject.Path
	// }
	// if newProject.Name != "" {
	// 	project.Name = newProject.Name
	// }
	
	return nil
}

func DeleteProject(path string) error {
	_, err := gorm.G[DBProject](db).Where("path = ?", path).Delete(ctx)
	return err
}

func GetProjects() ([]DBProject, error) {
	projects, err := gorm.G[DBProject](db).Find(ctx)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return []DBProject{}, nil
		}
		return nil, err
	}

	return projects, nil
}
