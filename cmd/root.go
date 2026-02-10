package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/MVN-14/devboard-go/devboard"
	_ "github.com/glebarez/go-sqlite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var board *devboard.Board

var rootCmd = &cobra.Command{
	Use:   "devboard",
	Short: "CLI for devboard",
	Long: `
Command line interface to interact with dev projects. 
Can open, add, remove, update, and list projects tracked by devboard`,
	Version: "1.0.0",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := initConfig(cmd)
		if err != nil {
			return err
		}
		
		dbPath := viper.GetString("dbpath")
		if dbPath == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			dbPath = path.Join(home, ".devboard", "devboard.db")
		}

		board, err = devboard.New(dbPath)
		if err != nil {
			return err
		}

		if viper.GetBool("verbose") {
			fmt.Println("Connected to db at", dbPath)
		}
		return nil
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		board.Close()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig(cmd *cobra.Command) error {
	viper.SetEnvPrefix("DEVBOARD")
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(".")
		viper.AddConfigPath(path.Join(home, ".devboard"))
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return err
	}

	if viper.GetBool("verbose") {
		fmt.Println("Read config from", viper.ConfigFileUsed())
		fmt.Println("Config is:", viper.AllSettings())
	}

	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file priority: flag value (.yaml file), config.yaml in project directory, $HOME/.devboard/devboard.yaml")
	rootCmd.PersistentFlags().String("dbpath", "", "path to sqlite database (default - $HOME/.devboard/devboard.db)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}
