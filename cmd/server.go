package cmd

import (
	"fmt"
	"os"

	"github.com/MokkeMeguru/PCA_Pi_Server/internal/router"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var AppName = "PCA_Pi_Server"

func runServer(cmd *cobra.Command, args []string) {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("file read err: %s", err)
		os.Exit(1)
	}
	port := fmt.Sprintf(":%d", viper.Get("port"))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r := gin.Default()
	r.Use(cors.New(config))
	router.AddAlarmRoute(r)
	router.AddSoundRoute(r)
	r.Run(port)
}

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Run:   runServer,
}

func init() {
	rootCmd.AddCommand(serverCommand)
}
