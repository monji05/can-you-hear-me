/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	// "charm.land/lipgloss/v2"
	"github.com/spf13/cobra"
)

type Happiness struct {
  Date string `json:"date"` // jsonのキーのaliasみたいな
  Content []Content `json:"content"`
}

type Content struct {
  Detail string `json:"detail"`
}


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "can-you-hear-me",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
  Run: func(cmd *cobra.Command, args []string) {

    fileByte, err := os.ReadFile("data.json")

    if err != nil {
      log.Fatal(err)
    }

    var records []Happiness

    err = json.Unmarshal(fileByte, &records)

    if err != nil {
      log.Fatal(err)
    }

    // NOTE: 2006-01-02じゃないとだめ
    // NOTE: これは2006年の1月2日としているわけではなく、 1月2日 3時4分5秒 2006年=123456と並べられるもの
    // NOTE: 独特すぎる。。
    today := time.Now().Format("2006-01-02")

    var contents []Content

    if len(args) == 0 {
      return
    }

    // TODO:同じ日付があればmerge
    for _, arg := range args {
      content := Content {
        Detail: arg,
      }

      contents = append(contents, content)
    }

    var happiness = Happiness {
      Date: today,
      Content: contents,
    }

    records = append(records, happiness)
    buf, err := json.Marshal(records)

    if err != nil {
      log.Fatal(err)
    }

    os.WriteFile("data.json", buf, 0640)
    ShowTip()
	},
}

func ShowTip() {
  fmt.Println("hello")

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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.can-you-hear-me.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


