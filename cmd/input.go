/*
Copyright © 2019 Dustin Hemard <dphemard@gmail.com>

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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// inputCmd represents the input command
var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "enter MPG records one by one",
	Long:  `For each MPG record the user is prompted for each field.`,
	RunE:  inputPrompt,
}

func init() {
	rootCmd.AddCommand(inputCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inputCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inputCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func inputPrompt(cmd *cobra.Command, args []string) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("- - - - - Enter A Record - - - - -")
		fmt.Print("Odometer Reading: ")
		odometer, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		odometer = strings.TrimSpace(odometer)

		fmt.Print("Estimated MPG: ")
		estMPG, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		estMPG = strings.TrimSpace(estMPG)

		fmt.Println()
		fmt.Println("Save it?")
		fmt.Println("yes (y), no (n), quit (q)")

		response, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		response = strings.ToLower(strings.TrimSpace(response))
		switch response {
		case "y":
			// save it
		case "n":
			// don't save it
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Please select either y, n, or q")
			os.Exit(1)
		}

	}
}
