package cmd

import (
	"encoding/json"

	"github.com/MVN-14/devboard-go/devboard"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update PROJECT",
	Short: "Update a project",
	Long: `Update a project being tracked by devboard.

Project argument must be a json string in the following format:
{
    id: 1                                 - ID of project to update
	name: "My Project",                   - name of your project
    path: "/home/joe/projectDir",         - the path to your project directory
    command: "vim /home/joe/projectDir",  - (optional) startup command for project
}
	`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), isValidUpdateProjectArg),
	RunE: updateProject,
}

func updateProject(cmd *cobra.Command, args []string) error {
	p := devboard.Project{}
	json.Unmarshal([]byte(args[0]), &p)

	_, err := db.Exec(`
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

func init() {
	rootCmd.AddCommand(updateCmd)
}
