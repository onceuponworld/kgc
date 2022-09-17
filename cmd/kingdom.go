package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"os"
	//"strings"

	"github.com/onceuponworld/ouw-sdk"
	"github.com/spf13/cobra"
)


var (
	
	fFile			string
	
	kingdomCmd = &cobra.Command{
		Use: "kingdom",
		Short: "create municipals and plots",
		Long: "generate random plots within municipals",
		Run: func(cmd *cobra.Command, args []string) {
			generateKingdom()
		},
	}

)


func init() {

	kingdomCmd.PersistentFlags().StringVarP(&fFile, "file", "f",
    DEFAULT_PROFILE, "kingdom profile")

} // init


func generateKingdom() {

	_, err := os.Stat(fFile)
	
	if err != nil || os.IsNotExist(err) {
		fmt.Println(err)
	} else {

		buf, err := ioutil.ReadFile(fFile)

		if err != nil {
			fmt.Println(err)
		} else {

			var k Kingdom

			err := json.Unmarshal(buf, &k)

			if err != nil {
				fmt.Println(err)
			} else {

				ouwsdk.Init(true)

				ouwsdk.SetAdd(ouwsdk.KEY_KINGDOMS, k.Name)

			}

		}

	}

} // generateKingdom
