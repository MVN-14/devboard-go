package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var openCmd = &cobra.Command{
	Use:   "open ID",
	Short: "Open devboard project",
	Long:  `Open a project in devboard based on the value of "command"

	Looks for "command" in the following order until it finds a value:
	1. --command flag passed to program at runtime
	2. "command" field on the project
	3. DEVBOARD_CMD env var
	4. "cmd" value in config

	if none are present it will throw an error
	
	*command value can reference project path using '$projPath' variable
	e.g. "command": "alacritty nvim $projPath"
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		
		if viper.GetBool("verbose") {
			fmt.Println("Opening project with id", id)
		}
		p, err := board.Get(id)
		if err != nil {
			return err
		}

		err = board.Open(p)
		if err != nil {
			return err
		}
		if viper.GetBool("verbose") {
			fmt.Println("Opened project", p.Name)
		}
		return nil
	},
	Args: cobra.MatchAll(cobra.ExactArgs(1), isValidIDArg),
}

func init() {
	rootCmd.AddCommand(openCmd)
	openCmd.Flags().String("command", "", "commmand to use to open the project (will override 'command' value of project)")
}
