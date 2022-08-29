package cmd

import (
	"fmt"

	"framework_v1/core"
	"framework_v1/logger"

	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts the API Server",
	Long: `
		
	Environment Variables
	---------------------------------
	APP_PORT: The port, the server listens on (default: 8080)
	MYSQL_USER: The MYSQL database user (default: root)
	MYSQL_PASSWORD: The MYSQL database password (default: my$qlroot)
	MYSQL_HOST: Host of the MYSQL databases (default: localhost)
	MYSQL_PORT: Port for connecting to the MYSQL database (default: 3306)
	`,
	Run: apiServerRun,
}

func apiServerRun(*cobra.Command, []string) {
	fmt.Println("Starting API Server...  Done!")

	c := new(core.Config)
	s := core.NewServer(c)
	fmt.Println("Starting API Server...  Done-- 2!")

	s.Start()
}

func init() {
	logger.InitLogger()
	fmt.Println("Hello API")
	rootCmd.AddCommand(apiCmd)

}
