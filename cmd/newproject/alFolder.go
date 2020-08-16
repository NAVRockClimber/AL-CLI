package newproject

import (
	"os"
	"path"
)

/*
CreateFolderStructure creates the al folder structure
*/
func CreateFolderStructure(appName string, app, test, res bool, appFolderName, resourceFolderName, testFolderName string) {
	var appFolder, testFolder, resFolder string
	appFolder = path.Join(appName, appFolderName)
	testFolder = path.Join(appName, testFolderName)
	resFolder = path.Join(appName, resourceFolderName)
	checkAndCreateFolder(appName)
	checkAndCreateFolder(appFolder)
	checkAndCreateFolder(testFolder)
	checkAndCreateFolder(resFolder)
}

func checkAndCreateFolder(folderName string) {
	_, err := os.Stat(folderName)
	if os.IsNotExist(err) {
		os.Mkdir(folderName, os.ModeDir)
	}
}
