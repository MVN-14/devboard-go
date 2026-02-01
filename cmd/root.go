package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

	"database/sql"

	_ "github.com/glebarez/go-sqlite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var db *sql.DB
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "devboard",
	Short: "CLI for devboard",
	Long: `
Command Line Interface to interact with devboard. 
Can add, remove, update, and list projects tracked by devboard`,
	Version:           "1.0.0",
	PersistentPreRunE: setUpDevboard,
	PersistentPostRun: closeDb,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func setUpDevboard(cmd *cobra.Command, args []string) error {
	err := initConfig(cmd)
	if err != nil {
		return err
	}

	err = initDb()
	if err != nil {
		return err
	}
	return nil
}

func initDb() error {
	var err error
	dbPath := viper.GetString("dbpath")
	if dbPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		dbPath = path.Join(home, ".devboard", "devboard.db")
	}
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS projects (
		command 	TEXT,
		id 			INTEGER PRIMARY KEY,
		name 		TEXT NOT NULL,
		path 		TEXT NOT NULL,
		created_at	DATETIME NOT NULL DEFAULT (datetime('now', 'localtime')),
		updated_at	DATETIME NOT NULL DEFAULT (datetime('now', 'localtime')),
		deleted_at	DATETIME 
	);`)

	if err != nil {
		return err
	}
	return nil
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

func closeDb(*cobra.Command, []string) {
	db.Close()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.devboard/devboard.yaml)")
	rootCmd.PersistentFlags().String("dbpath", "", "Path to sqlite database")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Path to sqlite database")
}
