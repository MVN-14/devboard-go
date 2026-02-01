package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete ID",
	Short: "Delete a project from devboard",
	Long: `Delete a project from devboard by ID`,
	RunE: deleteProject,
	Args: cobra.MatchAll(cobra.ExactArgs(1), isValidIDArg),
}

func deleteProject(cmd *cobra.Command, args []string) error {
	res, err := db.Exec(`
		UPDATE projects SET deleted_at = (datetime('now', 'localtime')) WHERE id = ?;
	`, args[0])	
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	fmt.Printf("Successfully deleted %d rows", rows)

	return nil
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

