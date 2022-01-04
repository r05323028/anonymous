package cmd

import (
	"github.com/r05323028/anonymous/db"
	"github.com/spf13/cobra"
)

var user string
var password string
var database string
var host string
var port string

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
	dbCmd.PersistentFlags().StringVar(&host, "host", "127.0.0.1", "host")
	dbCmd.PersistentFlags().StringVar(&port, "port", "3306", "Port")
	dbCmd.PersistentFlags().StringVarP(&user, "user", "u", "root", "user name")
	dbCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password")
	dbCmd.PersistentFlags().StringVarP(&database, "database", "d", "anonymous", "database")
	dbCmd.MarkFlagRequired("password")
}
