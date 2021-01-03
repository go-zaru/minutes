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
	"fmt"

	"github.com/go-zaru/minutes/internal"
	"github.com/spf13/cobra"
)

// LogEditFileName ...
const LogEditFileName = "LOG_EDIT"
const isoDateFormat = "2006-01-02"
const isoDateTimeFormat = "2006-01-02T15:04:05-0700"

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Help track decisions",
	Long:  `Helps log decisions for later records, NSTF`,
	Run: func(cmd *cobra.Command, args []string) {

		editorBytes, err := internal.CaptureInputFromEditor(
			internal.GetPreferredEditorFromEnvironment,
		)

		// buf := bytes.NewBuffer(nil)
		// consoleReader := bufio.NewReader(os.Stdin)

		// now := time.Now()
		// now.Format(isoDateFormat)

		// fmt.Println("When:")
		// text, _ := consoleReader.ReadString('\n')
		// buf.WriteString(text)

		// fmt.Println("Who:")
		// text, _ := consoleReader.ReadString('¥')

		fmt.Println(editorBytes)
		fmt.Println(err)
		// When
		// Who
		// Comments

		// Create Log Directory
		// logDir := time.Now().Format(isoDateFormat)
		// os.Mkdir(logDir, os.ModePerm)
		// Create Scratch File Path
		// fp := path.Join(logDir, LogEditFileName)
		// consoleReader := bufio.NewReader(os.Stdin)
		// text, _ := consoleReader.ReadString('\n')
		// fmt.Println(text)
		// // Create command to start vi with scratch file
		// // can I use a Channel here?
		// vi := exec.Command("vi", fp)
		// vi.Stdin = os.Stdin
		// vi.Stdout = os.Stdout
		// err := vi.Run()
		// // Waits to completion
		// fpl := path.Join(logDir, "log")
		// logFile, err := os.OpenFile(fpl, os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
		// fmt.Println(err)

		// defer logFile.Close()
		// writer := bufio.NewWriter(logFile)

		// // Gather Meta
		// buf := bytes.NewBuffer(nil)
		// timestring := time.Now().Format(isoDateTimeFormat)
		// buf.WriteString("===\n")
		// buf.WriteString(fmt.Sprintf("%s\n", timestring))
		// buf.WriteString("---\n")
		// buf.WriteString("\n")
		// buf.WriteTo(writer)
		// buf.WriteString("\n")
		// buf.WriteString("===\n")
		// buf.WriteString("\n")

		// // Gather Content
		// logByte, err := ioutil.ReadFile(fp)
		// buf.Write(logByte)
		// buf.WriteString("\n")
		// buf.WriteTo(writer)
		// err = writer.Flush()
		// fmt.Println(err)
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
