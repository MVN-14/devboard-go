package cmd

import (
	"fmt"

	"github.com/MVN-14/devboard-go/devboard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Projects",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		projects := devboard.ProjectList{}

		sql := "SELECT * FROM projects"
		if !viper.GetBool("deleted") {
			sql += " WHERE deleted_at IS NULL"
		}

		rows, err := db.Query(sql)
		if err != nil {
			return err
		}
		defer rows.Close()
		
		if err := projects.FromRows(rows); err != nil {
			return err
		}
		if viper.GetBool("verbose") {
			fmt.Println("Projects List:")
		}
		fmt.Println(projects)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().Bool("deleted", false, "List deleted projects")
}
