package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:   "delete ID",
	Short: "Delete a project from devboard",
	Long: `Delete a project from devboard by ID`,
	RunE: deleteProject,
	Args: cobra.MatchAll(cobra.ExactArgs(1), isValidIDArg),
}

func deleteProject(cmd *cobra.Command, args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	err = board.Delete(id)
	if err != nil {
		return err
	}

	if viper.GetBool("verbose") {
		fmt.Printf("Successfully deleted project with id %s", args[0])
	}

	return nil
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

