package devboard

type DevboardArgs struct {
	Add     bool
	Id      int `docopt:"ID"`
	Help    bool `docopt:"-h, --help"`
	List    bool
	Project string `docopt:"PROJECT"`
	Remove  bool
	Update  bool
}

const Usage = `Devboard.

Usage: devboard (add | update) PROJECT
       devboard remove ID
       devboard list
       devboard (--help | -h)
       devboard --version

Arguments:
    ID         ID of project to update.
    PROJECT    Json string with the following properties:
                   name    - project name
                   path    - path to project directory

Options:
    -h --help    Show this screen.
    --version    Show version.
`
