package cmd

import (
	"github.com/r05323028/anonymous/db"
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database Command Line Tools",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var initDBCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize database",
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		database, _ := cmd.Flags().GetString("database")
		db.InitDB(user, password, host, port, database)
	},
}

var dropDBCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop database",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	dbCmd.AddCommand(initDBCmd)
	dbCmd.AddCommand(dropDBCmd)
	rootCmd.AddCommand(dbCmd)

	// add flags
	dbCmd.PersistentFlags().String("host", "127.0.0.1", "Database host")
	dbCmd.PersistentFlags().String("port", "3306", "Database port")
	dbCmd.PersistentFlags().String("user", "anonymous", "Database user")
	dbCmd.PersistentFlags().String("password", "example", "Database password")
	dbCmd.MarkPersistentFlagRequired("password")
}
