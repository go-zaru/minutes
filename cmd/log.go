/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/go-zaru/minutes/internal"
	"github.com/spf13/cobra"
)

// LogEditFileName ...
const LogDirName = ".mins"
const LogFileName = "logs"
const isoDateFormat = "2006-01-02"
const isoDateTimeFormat = "2006-01-02T15:04:05-0700"

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Help track decisions",
	Long:  `Helps log decisions for later records, NSTF`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get Data and Buf
		buf := bytes.NewBuffer(nil)
		buf.WriteString(fmt.Sprintln("===", time.Now().Format(isoDateTimeFormat)))
		tags, _ := internal.PromptString("Tags")
		buf.WriteString(fmt.Sprintln("[tags]", tags))
		who, _ := internal.PromptString("Who")
		buf.WriteString(fmt.Sprintln("[who]", who))
		comments, err := internal.CaptureInputFromEditor(
			internal.GetPreferredEditorFromEnvironment,
		)
		buf.WriteString(fmt.Sprintln("[comments:start]"))
		buf.Write(comments)
		buf.WriteString(fmt.Sprintln("[comments:end]"))
		buf.WriteString(fmt.Sprintln(""))

		homedir, err := os.UserHomeDir()
		// Create logdir and file
		logdir := path.Join(homedir, LogDirName)
		os.Mkdir(logdir, os.ModePerm)
		d := time.Now().Format(isoDateFormat)
		filepath := path.Join(logdir, d)
		logFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer logFile.Close()
		// Write to file
		writer := bufio.NewWriter(logFile)
		buf.WriteTo(writer)
		err = writer.Flush()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
