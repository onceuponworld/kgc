package cmd

import (
	"encoding/json"
	"fmt"
	//"log"
	"os"
	//"strings"

	"github.com/spf13/cobra"
)


var (
	
	fOut			string
	fForce    bool
	
	initCmd = &cobra.Command{
		Use: "init",
		Short: "initialize json kingdom profile",
		Long: "initialize json kingdom profile",
		Run: func(cmd *cobra.Command, args []string) {
			initProfile()
		},
	}

)


func init() {

	initCmd.PersistentFlags().StringVarP(&fOut, "out", "o",
    DEFAULT_PROFILE, "output filename")
	initCmd.PersistentFlags().BoolVarP(&fForce, "force", "f",
	  DEFAULT_OVERWRITE, "force file overwrite")

} // init


func createProfile() {

	k := Kingdom {
		Municipals: []Municipal{Municipal{}},
	}

	j, err := json.MarshalIndent(k, EMPTY_STR, CHAR_TAB)

	if err != nil {
		fmt.Println(err)
	} else {

		err := os.WriteFile(fOut, j, 0644)

		if err != nil {
			fmt.Println(err)
		}
			
	}

} // createProfile


func initProfile() {

	_, err := os.Stat(fOut)
	
	if os.IsNotExist(err) {
		createProfile()
	} else {

		if(fForce) {
			createProfile()
		} else {
			fmt.Println(ERR_FILE_EXISTS)
		}

	}

} // generateKingdom
