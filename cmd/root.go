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
	"strconv"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/spf13/cobra"
)

type Happiness struct {
  Date string `json:"date"` // jsonのキーのaliasみたいな
  Contents []Content `json:"contents"`
  Count int `json:"count"`
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
    records, err := readFile()

		if err != nil {
			log.Fatal(err)
			return
		}

    today := time.Now().Format(time.DateOnly)

    if len(args) == 0 {
			if records == nil {
        fmt.Println("󰛩 今日あった嬉しかったこと、良かったこと、頑張ったことを教えてください󰛩 ")
				return
			}
			showGraph(records, today)
      return
    }

    newContents := AddContents(args)
    newRecords := AddRecords(records, newContents, today)

    buf, err := json.MarshalIndent(newRecords, "", "  ")

    if err != nil {
      log.Fatal(err)
    }

    os.WriteFile("data.json", buf, 0640)
    showGraph(newRecords, today)
	},
}

func readFile() ([]Happiness, error){
  file, err := os.Open("data.json")

  if err != nil {
    // ファイルが存在しない場合は初回実行とみなす
    if errors.Is(err, os.ErrNotExist) {
      return nil, nil
    }

    log.Fatal(err)
		return nil, err
  }

  r, err := Read(file)

	if err != nil {
		log.Fatal(err)
    return nil, err
	}

  defer file.Close()
  return r, nil
}

func showGraph(records []Happiness, today string) {
  grassChar := " "
  darkGray := lipgloss.Color("#3C3C3C")
  darkenGray := lipgloss.Darken(darkGray, 0.25)
  level0 := lipgloss.NewStyle().Foreground(darkenGray)
  level1 := lipgloss.NewStyle().Foreground(lipgloss.Color("#9be9a8"))
  level2 := lipgloss.NewStyle().Foreground(lipgloss.Color("#4d8b60"))
  level3 := lipgloss.NewStyle().Foreground(lipgloss.Color("#216e39"))
  level4 := lipgloss.NewStyle().Foreground(lipgloss.Color("#1a582d"))

  now := time.Now()
  year := now.Year()

  happinessMap := make(map[string]int)
  for _, record := range records {
    happinessMap[record.Date] = record.Count
  }

  var date string = "    "
  for day := 1; day <= 31; day++ {
		if day % 5 == 0 {
			date += fmt.Sprintf("%02d", day)
		} else {
			date += "  "
		}
  }
  fmt.Printf("Happiness in %s \n", strconv.Itoa(year))
  fmt.Println(date)

  var grassRow string
  for m := 1; m <= 12; m ++ {
    daysInMonth := time.Date(year, time.Month(m), 0 ,0, 0, 0, 0, time.Local).Day()
    grassRow = fmt.Sprintf("%02d  ", m)

    for day := 1;  day <= daysInMonth; day++ {
      key := fmt.Sprintf("%04d-%02d-%02d", year, time.Month(m), day)
      count := happinessMap[key]
      var rendered string
      switch {
      case count == 0:
        rendered = level0.Render(grassChar)
      case count == 1:
        rendered = level1.Render(grassChar)
      case count == 2:
        rendered = level2.Render(grassChar)
      case count == 3:
        rendered = level3.Render(grassChar)
      default:
        rendered = level4.Render(grassChar)
      }
      grassRow += rendered
    }
    fmt.Println(grassRow)
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


