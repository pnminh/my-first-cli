# my-first-cli
A cli buit with Golang and uses Cobra and Viper
## Setup Golang dev environment
Initialize a Go module:
```
$ go mod init github.com/pnminh/my-first-cli
```
Install the Cobra and Viper libraries:
```
$ go get github.com/spf13/cobra
$ go get github.com/spf13/viper
```
First iteration: Only uses Cobra. Example with `root.go` in `cmd` package
```go
package cmd
import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
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
	err := rootCmd.Execute()
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
}  
```
Add `main.go`, which belongs to `main` package and has `main` function as entry point.
```go
package main
import "github.com/pnminh/my-first-cli/cmd"
func main(){
	cmd.Execute()
}
```
Run build from root folder
```bash
$ go build
$ ./my-first-cli
```
Iteration 2: Add configs with viper
Add `config.yaml` file
```yaml
app:
  name: My CLI App
  version: 1.0

```
Update `root.go` with viper. First read in config file
```go
func initConfig(){
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err!= nil {
        fmt.Println("Failed to read config file:", err)
        os.Exit(1)
    }
}
```
Then add this as part of initialization for cobra command
```go
cobra.OnInitialize(initConfig)
```