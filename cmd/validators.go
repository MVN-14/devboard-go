package cmd

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/MVN-14/devboard-go/devboard"
	"github.com/spf13/cobra"
)

func isValidAddProjectArg(cmd *cobra.Command, args []string) error {
	p := devboard.Project{}
	if err := json.Unmarshal([]byte(args[0]), &p); err != nil {
		return err
	}

	if p.Name == "" {
		return errors.New("Error: project is missing Name field\n")
	} 
	if p.Path == "" {
		return errors.New("Error: project is missing Path field\n")
	}
	
	return nil
}

func isValidUpdateProjectArg(cmd *cobra.Command, args []string) error {
	p := devboard.Project{}
	if err := json.Unmarshal([]byte(args[0]), &p); err != nil {
		return err
	}

	if p.Name == "" {
		return errors.New("Error: project is missing Name field\n")
	} 
	if p.Path == "" {
		return errors.New("Error: project is missing Path field\n")
	}
	if p.Id == 0 {
		return errors.New("Error: project is missing Id field\n")
	}
	return nil
}

func isValidIDArg(cmd *cobra.Command, args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	if id <= 0 {
		return errors.New("Error: ID argument must be greater than 0")
	}

	return nil
}
