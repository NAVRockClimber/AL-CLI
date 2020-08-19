package newproject

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*InitManifest sets the flags for the manifest*/
func InitManifest(newCmd *cobra.Command) {
	newCmd.Flags().StringP("Version", "V", "1.0.0.0", "Set Version field in app manifest. If empty the from the settings is taken.")
	newCmd.Flags().StringP("Description", "D", "", "Set Description field in app manifest. If empty the from the settings is taken.")
	newCmd.Flags().StringP("Brief", "B", "", "Set Brief field in app manifest. If empty the from the settings is taken.")
	newCmd.Flags().String("PrivacyStatement", "", "Set PrivacyStatement field in app manifest. If empty the from the settings is taken.")
	newCmd.Flags().String("EULA", "", "Set EULA field in app manifest. If empty the from the settings is taken.")
	newCmd.Flags().String("Help", "", "Set Help field in app manifest. If empty the from the settings is taken.")
	newCmd.Flags().String("URL", "", "Set URL field in app manifest. If empty the from the settings is taken.")
	newCmd.Flags().String("Logo", "", "Set Logo field in app manifest. If empty the from the settings is taken.")
	newCmd.Flags().StringP("Platform", "P", "", "Set Platform field in app manifest. If empty the from the settings is taken.")
	newCmd.Flags().StringP("SupportedLocales", "L", "TranslationFile", "Set SupportedLocales field in app manifest. If empty the from the settings is taken.")
	defaultFeatures := []string{"TranslationFile"}
	newCmd.Flags().StringSliceP("Features", "F", defaultFeatures, "Set Features field in app manifest. If empty the from the settings is taken.")
	newCmd.Flags().Bool("ShowMyCode", true, "Set ShowMyCode field in app manifest. If empty the from the settings is taken.")
	newCmd.Flags().StringP("Runtime", "R", "", "Set Runtime field in app manifest. If empty the from the settings is taken.")
	newCmd.Flags().StringP("Publisher", "p", "", "Set company field in app manifest. If empty the default from setup is taken.")
	defaultRange := []int{0, 0}
	newCmd.Flags().IntSliceP("Range", "I", defaultRange, "Define From field in app manifests id range")
}

/*CreateAppJSON create a new app manifest file*/
func CreateAppJSON(cmd *cobra.Command, appName string, appFolder string) appManifest {
	appID := uuid.New().String()
	manifest := appManifest{
		ID:   appID,
		Name: appName,
	}

	manifest.Name = getStringValue(cmd, "Name", "App Manifest field", false)
	manifest.Publisher = getStringValue(cmd, "Publisher", "App Manifest field", false)
	manifest.Version = getStringValue(cmd, "Version", "App Manifest field", true)
	manifest.Description = getStringValue(cmd, "Description", "App Manifest field", false)
	manifest.Brief = getStringValue(cmd, "Brief", "App Manifest field", true)
	manifest.PrivacyStatement = getStringValue(cmd, "PrivacyStatement", "App Manifest field", false)
	manifest.EULA = getStringValue(cmd, "EULA", "App Manifest field", false)
	manifest.Help = getStringValue(cmd, "Help", "App Manifest field", false)
	manifest.URL = getStringValue(cmd, "URL", "App Manifest field", false)
	manifest.Logo = getStringValue(cmd, "Logo", "App Manifest field", false)
	manifest.Platform = getStringValue(cmd, "Platform", "App Manifest field", true)
	manifest.Runtime = getStringValue(cmd, "Runtime", "App Manifest field", true)

	createAppJSON(appFolder, manifest)
	return manifest
}

func getStringValue(cmd *cobra.Command, key, question string, request bool) string {
	value, _ := cmd.Flags().GetString(key)
	if value == "" {
		value = viper.GetString(key)
	}
	if request {
		questionWithField := question + " " + key
		value = requestString(questionWithField, value)
	}
	return value
}

func requestString(question, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question, defaultValue)
	answer, _ := reader.ReadString('\n')
	if answer == "\r\n" {
		answer = defaultValue
		fmt.Printf("Reurning:  %s", answer)
	}
	return answer
}

func requestIntDefault(question string, defaultValue int) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(question, defaultValue)
	answer, _ := reader.ReadString('\n')
	answerInt := 0
	if answer != "" {
		answerInt, _ = strconv.Atoi(answer)
	}
	return answerInt
}

func requestInt(question string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	answer, _ := reader.ReadString('\n')
	answerInt := 0
	if answer != "" {
		answerInt, _ = strconv.Atoi(answer)
	}
	return answerInt
}

func createAppJSON(appFolderName string, manifest appManifest) {
	var jsonFileName string
	curpath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	jsonFileName = filepath.Join(curpath, appFolderName, "App.json")
	jsonManifest, _ := converToJSONString(manifest)
	fmt.Printf("Writing manifest to: %s \n", jsonFileName)
	ioerr := ioutil.WriteFile(jsonFileName, jsonManifest, os.ModePerm)
	if ioerr != nil {
		fmt.Printf("Writing manifest to: %s", ioerr.Error())
	}
}

func converToJSONString(manifest appManifest) ([]byte, error) {
	jsonManifest, marshalErr := json.MarshalIndent(manifest, "", "  ")
	if marshalErr != nil {
		fmt.Println(marshalErr.Error())
	}
	return jsonManifest, marshalErr
}
