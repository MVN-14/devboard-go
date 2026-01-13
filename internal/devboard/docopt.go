package devboard

type DevboardArgs struct {
	Add     bool
	Id      uint `docopt:"ID"`
	Help    bool `docopt:"-h, --help"`
	List    bool
	Project string `docopt:"PROJECT"`
	Path    string `docopt:"PATH"`
	Remove  bool
	ResetDB bool `docopt:"resetdb"`
	Update  bool
}

const Usage = `Devboard.

Usage: devboard add PROJECT
       devboard update ID PROJECT
       devboard remove ID
       devboard list
       devboard resetdb
       devboard (--help | -h)
       devboard --version

Arguments:
    ID         ID of project to update.
    PATH       Path to project directory.
    PROJECT    Json string with the following properties:
                   name    - project name
                   path    - path to project directory

Options:
    -h --help    Show this screen.
    --version    Show version.
`
