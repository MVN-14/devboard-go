package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/MVN-14/devboard-go/devboard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var updateCmd = &cobra.Command{
	Use:   "update PROJECT",
	Short: "Update a project",
	Long: `Update a project being tracked by devboard.

Project argument must be a json string in the following format:
{
    id: 1                                   - ID of project to update
	name: "My Project",                     - name of your project
    path: "/home/joe/projectDir",           - the path to your project directory
    command: "alacritty -e nvim $projPath"  - (optional) startup command for project
}
	`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), isValidUpdateProjectArg),
	RunE: updateProject,
}

func updateProject(cmd *cobra.Command, args []string) error {
	p := devboard.Project{}
	json.Unmarshal([]byte(args[0]), &p)
	
	err := board.Update(p)
	if err != nil {
		return err
	}

	if viper.GetBool("verbose") {
		fmt.Printf("Successfully updated project %s\n", p.Name)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
