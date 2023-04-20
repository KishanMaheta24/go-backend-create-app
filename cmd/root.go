package cmd

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"os/exec"
	"reflect"

	"github.com/spf13/cobra"
)

var StubStorage = map[string]interface{}{
	"Beego App": beegoApp,
	"Beego API": beegoAPI,
	"Gin":       gin,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "create-backend-app",
	Short: "initializes structure for golang frameworks",
	Long:  `initializes structure for golang frameworks 1.gin 2.beego-app 3.beego-api 4.mux`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Choose framework",
			Items: []string{"Gin", "Beego App", "Beego API", "MUX"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		dir := takeDirName()
		res, _ := Call(result, dir)
		var prnt string
		prnt = res.(string)
		fmt.Println(prnt)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-backend-create-app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func Call(funcName string, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(StubStorage[funcName])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is out of index.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	result = res[0].Interface()
	return
}

func beegoApp(dir string) string {

	if _, err := os.Stat("./" + dir); errors.Is(err, os.ErrNotExist) {
		fmt.Println("creating......")

	} else {
		fmt.Println("directory already exists... please try different name again")
		os.Exit(0)
	}
	cmd := exec.Command("bee", "new", dir)

	if !execCmd(cmd) {
		return ""
	} else {

		cmd.Run()
		return "successfully created"
	}
	return "BEE APP"
}

func beegoAPI(dir string) string {
	if _, err := os.Stat("./" + dir); errors.Is(err, os.ErrNotExist) {
		fmt.Println("creating......")

	} else {
		fmt.Println("directory already exists... please try different name again")
		os.Exit(0)
	}
	cmd := exec.Command("bee", "api", dir)

	if !execCmd(cmd) {
		return ""
	} else {
		cmd.Run()
		return "successfully created"
	}
	return "BEE API"
}

func gin(dir string) string {
	cmd := exec.Command("git", "clone", "https://github.com/code4mk/go-gin-boilerplate")

	if !execCmd(cmd) {
		return ""
	}

	cmd = exec.Command("mv", "go-gin-boilerplate", dir)

	if !execCmd(cmd) {
		return ""
	}

	cmd = exec.Command("bash", "-c", "cd", dir, ";rm -fr .git")
	fmt.Println(cmd.String())
	if !execCmd(cmd) {
		return ""
	}

	return "GIN"
}

func takeDirName() string {
	prompt := promptui.Prompt{
		Label: "Parent Dir Name",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func execCmd(cmd *exec.Cmd) bool {
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
