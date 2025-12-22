package devboard

type DevboardArgs struct {
	Add     bool
	Help    bool `docopt:"-h, --help"`
	List    bool
	Project string `docopt:"PROJECT"`
	Path string `docopt:"PATH"`
	Remove  bool
	Update  bool
}

type Project struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type DevboardData struct {
	Projects []Project `json:"projects"`
}

func (d *DevboardData) ContainsProject(path string) bool {
	for _, v := range d.Projects {
		if v.Path == path {
			return true
		}
	}
	return false
}
