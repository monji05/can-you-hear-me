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
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/spf13/cobra"
)

type Happiness struct {
  Date string `json:"date"` // jsonのキーのaliasみたいな
  Contents []Content `json:"contents"`
  Count int `json:"count"`
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

		var records []Happiness
    // NOTE: 初回実行時はdata.jsonなんかない、どうする
    fileByte, err := os.ReadFile("data.json")

    if err != nil  && !errors.Is(err, os.ErrNotExist) {
			// ファイルが存在しない場合は初回実行とみなし、空のスライスのまま進める
			log.Fatal(err)
    } else {
			err = json.Unmarshal(fileByte, &records)

			if err != nil {
				log.Fatal(err)
			}
		}

    today := time.Now().Format(time.DateOnly)

    if len(args) == 0 {
			showGraph(records, today)
      return
    }

		newContents := make([]Content, 0, len(args))
    for _, arg := range args {
      newContents = append(newContents, Content {
				Detail: arg,
			})
    }

		isTodayFlg := false
    for index, record := range records {
      if today == record.Date {
				isTodayFlg = true
        records[index].Contents = append(records[index].Contents, newContents...)
        fmt.Println(records[index].Contents)
        records[index].Count = len(records[index].Contents)
				break
      }
    }

		if !isTodayFlg {
			records = append(records, Happiness {
				Date: today,
				Contents: newContents,
        Count: len(newContents),
			})
		}

    buf, err := json.MarshalIndent(records, "", "  ")

    if err != nil {
      log.Fatal(err)
    }

    os.WriteFile("data.json", buf, 0640)
    showGraph(records, today)
	},
}

func showGraph(records []Happiness, today string) {
  grassChar := "■ "
  level0 := lipgloss.NewStyle().Foreground(lipgloss.Black)
  level1 := lipgloss.NewStyle().Foreground(lipgloss.Color("#9be9a8"))
  level2 := lipgloss.NewStyle().Foreground(lipgloss.Color("#40c463"))
  level3 := lipgloss.NewStyle().Foreground(lipgloss.Color("#30a14e"))
  level4 := lipgloss.NewStyle().Foreground(lipgloss.Color("#216e39"))

  for _, record := range records {
    switch record.Count {
    case 0:
      fmt.Println(level0.Render(grassChar))
    case 1:
      fmt.Println(level1.Render(grassChar))
    case 2:
      fmt.Println(level2.Render(grassChar))
    case 3:
      fmt.Println(level3.Render(grassChar))
    default:
      fmt.Println(level4.Render(grassChar))
    }
  }
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


