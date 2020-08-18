package newproject

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

/*InitAlFolder sets up command flags concerning the folder structure*/
func InitAlFolder(newCmd *cobra.Command) {
	newCmd.Flags().BoolP("NoResource", "r", false, "Do not create a resource folder")
	newCmd.Flags().String("ResourceFolder", "res", "Set the name for the resource folder. Default: res")
	newCmd.Flags().BoolP("NoTest", "t", false, "Do not create a test folder")
	newCmd.Flags().String("TestFolder", "test", "Set the name for the test folder. Default: test")
	newCmd.Flags().String("AppFolder", "app", "Set the name for the app folder. Default: app")
}

/*
CreateFolderStructure creates the al folder structure
*/
func CreateFolderStructure(cmd *cobra.Command, appName string) (string, string) {
	var noResourceFlag, noTestFlag bool
	var appFolder, resourceFolder, testFolder string
	noResourceFlag, _ = cmd.Flags().GetBool("NoResource")
	noTestFlag, _ = cmd.Flags().GetBool("NoTest")
	appFolder, _ = cmd.Flags().GetString("AppFolder")
	resourceFolder, _ = cmd.Flags().GetString("ResourceFolder")
	testFolder, _ = cmd.Flags().GetString("TestFolder")
	return createFolderStructure(appName, true, !noTestFlag, !noResourceFlag, appFolder, resourceFolder, testFolder)
}

func createFolderStructure(appName string, app, test, res bool, appFolderName, resourceFolderName, testFolderName string) (string, string) {
	var appFolder, testFolder, resFolder string
	appFolder = path.Join(appName, appFolderName)
	testFolder = path.Join(appName, testFolderName)
	resFolder = path.Join(appName, resourceFolderName)
	checkAndCreateFolder(appName)
	checkAndCreateFolder(appFolder)
	checkAndCreateFolder(testFolder)
	checkAndCreateFolder(resFolder)
	return appFolder, testFolder
}

func checkAndCreateFolder(folderName string) {
	_, err := os.Stat(folderName)
	if os.IsNotExist(err) {
		os.Mkdir(folderName, os.ModeDir)
	}
}
