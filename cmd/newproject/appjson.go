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
	newCmd.Flags().StringP("Publisher", "p", "", "Set company field in app manifest. If empty the default from setup is taken.")
	newCmd.Flags().Int("FromRange", 0, "Define From field in app manifests id range")
	newCmd.Flags().Int("ToRange", 0, "Define To field in app manifests id range")
}

/*CreateAppJSON create a new app manifest file*/
func CreateAppJSON(cmd *cobra.Command, appName string, appFolder string) {
	appID := uuid.New().String()
	manifest := appManifest{
		ID:   appID,
		Name: appName,
	}
	publisher, _ := cmd.Flags().GetString("Publisher")
	if publisher == "" {
		viper.GetString("Publisher")
	}
	manifest.Name = requestString("App name (%s): ", manifest.Name)
	manifest.Publisher = publisher
	fromRange, _ := cmd.Flags().GetInt("FromRange")
	toRange, _ := cmd.Flags().GetInt("FromTo")

	appIDRange := idRange{
		From: fromRange,
		To:   toRange,
	}

	if fromRange <= 0 {
		appIDRange.From = requestInt("ID range start:")
	}
	if toRange <= 0 {
		appIDRange.To = requestInt("ID range end: ")
	}

	manifest.IDRanges = append(manifest.IDRanges, appIDRange)
	createAppJSON(appFolder, manifest)
}

func requestString(question, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(question, defaultValue)
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
