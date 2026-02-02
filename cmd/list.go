package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Projects",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		projects, err := board.List(viper.GetBool("deleted"))
		if err != nil {
			return err
		}
		if viper.GetBool("verbose") {
			fmt.Println("Project List:")
		}
		fmt.Println(projects)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("deleted", "d", false, "List deleted projects")
}
