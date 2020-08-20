package newproject

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
)

/*InitGit sets the flags for git*/
func InitGit(newCmd *cobra.Command) {
	newCmd.Flags().Bool("DontInitGit", false, "Defines field Image in compose file")
}

/*CreateGit create a new docker compose file*/
func CreateGit(cmd *cobra.Command, rootFolder string) {
	curpath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	gitroot := filepath.Join(curpath, rootFolder)

	DontInitGit, _ := cmd.Flags().GetBool("DontInitGit")
	if !DontInitGit {
		git.PlainInit(gitroot, false)
	}
}
