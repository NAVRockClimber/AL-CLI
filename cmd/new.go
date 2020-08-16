/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"path"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new <projectname>",
	Short: "create a new AL project",
	Long:  `Create a new al project including folder structure, manifest, compose files and git.`,
	Run: func(cmd *cobra.Command, args []string) {
		var noResourceFlag, noTestFlag bool
		var appFolder, resourceFolder, testFolder string
		parameterCheck(cmd, args)
		noResourceFlag, _ = cmd.Flags().GetBool("NoResource")
		noTestFlag, _ = cmd.Flags().GetBool("NoTest")
		appFolder, _ = cmd.Flags().GetString("AppFolder")
		resourceFolder, _ = cmd.Flags().GetString("ResourceFolder")
		testFolder, _ = cmd.Flags().GetString("TestFolder")
		createFolderStructure(args[0], true, !noTestFlag, !noResourceFlag, appFolder, resourceFolder, testFolder)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().BoolP("NoResource", "r", false, "Do not create a resource folder")
	newCmd.Flags().String("ResourceFolder", "res", "Set the name for the resource folder. Default: res")
	newCmd.Flags().BoolP("NoTest", "t", false, "Do not create a test folder")
	newCmd.Flags().String("TestFolder", "test", "Set the name for the test folder. Default: test")
	newCmd.Flags().String("AppFolder", "app", "Set the name for the app folder. Default: app")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func parameterCheck(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		fmt.Println("New command takes exactly one parameter")
		os.Exit(0)
	}
}

func createFolderStructure(appName string, app, test, res bool, appFolderName, resourceFolderName, testFolderName string) {
	var appFolder, testFolder, resFolder string
	appFolder = path.Join(appName, appFolderName)
	testFolder = path.Join(appName, testFolderName)
	resFolder = path.Join(appName, resourceFolderName)
	_, rootFolderErr := os.Stat(appName)
	if os.IsNotExist(rootFolderErr) {
		os.Mkdir(appName, os.ModeDir)
	}
	_, appErr := os.Stat(appFolder)
	if app && os.IsNotExist(appErr) {
		os.Mkdir(appFolder, os.ModeDir)
	}
	_, testErr := os.Stat(testFolder)
	if test && os.IsNotExist(testErr) {
		os.Mkdir(testFolder, os.ModeDir)
	}
	_, err := os.Stat(resFolder)
	if res && os.IsNotExist(err) {
		os.Mkdir(resFolder, os.ModeDir)
	}
}
