package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
var rootCmd = &cobra.Command{
	Use: "myFirstCli",
	Short: "My first cli app",
	Long:  "My first CLI app is a sample application built with Go, Cobra, and Viper.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world from root command")
	},
}
func Execute(){
	cobra.OnInitialize(initConfig)
	err := rootCmd.Execute()
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
func initConfig(){
	fmt.Println("setting up initConfig with viper...")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err!= nil {
		fmt.Println("Failed to read config file:", err)
		os.Exit(1)
	}
	appName:=viper.GetString("app.name")
	appVersion:=viper.GetString("app.version")
	fmt.Println("Application Name:", appName)
	fmt.Println("Application Version:", appVersion)
}