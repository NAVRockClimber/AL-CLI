package newproject

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

/*InitCompose sets the flags for the compose file*/
func InitCompose(newCmd *cobra.Command) {
	newCmd.Flags().String("ComposeVersion", "2.4", "Defines field Version in compose file")
	newCmd.Flags().String("Image", "", "Defines field Image in compose file")
	newCmd.Flags().StringP("ContainerName", "N", "", "Defines field ContainerName in compose file")
	newCmd.Flags().String("MemLimit", "8GB", "Defines field MemLimit in compose file")
	newCmd.Flags().String("AcceptEula", "Y", "Defines field AcceptEula in compose file")
	newCmd.Flags().String("AcceptOutdated", "Y", "Defines field AcceptOutdated in compose file")
	newCmd.Flags().String("UseSSL", "N", "Defines field UseSSL in compose file")
	newCmd.Flags().String("Licensefile", "C:\\run\\my\\license.flf", "Defines field Licensefile in compose file")
	newCmd.Flags().StringP("Username", "U", "admin", "Defines field Username in compose file")
	newCmd.Flags().String("Password", "", "Defines field Password in compose file")
}

/*CreateComposeFile create a new docker compose file*/
func CreateComposeFile(cmd *cobra.Command, rootFolder string) {
	var compose simpleCompose

	compose.Version = getStringValue(cmd, "ComposeVersion", "Compose file field ", false)
	compose.Services.BC.Image = getStringValue(cmd, "Image", "Compose file field ", true)
	compose.Services.BC.ContainerName = getStringValue(cmd, "ContainerName", "Compose file field ", false)
	compose.Services.BC.MemLimit = getStringValue(cmd, "MemLimit", "Compose file field ", false)
	compose.Services.BC.Environment.AcceptEula = getStringValue(cmd, "AcceptEula", "Compose file field ", false)
	compose.Services.BC.Environment.AcceptOutdated = getStringValue(cmd, "AcceptOutdated", "Compose file field ", false)
	compose.Services.BC.Environment.UseSSL = getStringValue(cmd, "UseSSL", "Compose file field ", false)
	compose.Services.BC.Environment.Licensefile = getStringValue(cmd, "Licensefile", "Compose file field ", false)
	compose.Services.BC.Environment.Username = getStringValue(cmd, "Username", "Compose file field ", true)
	compose.Services.BC.Environment.Password = getStringValue(cmd, "Password", "Compose file field ", true)

	createComposeFile(rootFolder, compose)
}

func createComposeFile(rootFolder string, compose simpleCompose) {
	var composeFileName string
	curpath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	composeFileName = filepath.Join(curpath, rootFolder, "docker-compose.yaml")
	jsonManifest, _ := convertToYAMLString(compose)
	ioerr := ioutil.WriteFile(composeFileName, jsonManifest, os.ModePerm)
	if ioerr != nil {
		fmt.Print("Error writing compose file: ", ioerr.Error())
	}
}

func convertToYAMLString(compose simpleCompose) ([]byte, error) {
	yamlString, marshalErr := yaml.Marshal(compose)
	if marshalErr != nil {
		fmt.Println(marshalErr.Error())
	}
	return yamlString, marshalErr
}
