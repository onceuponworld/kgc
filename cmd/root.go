package cmd

import (

	"github.com/spf13/cobra"
)


const (
	DEFAULT_OVERWRITE			= false
  DEFAULT_PATH					= "."
	DEFAULT_PROFILE				= "profile.json"
)


const (
	CHAR_TAB							= "\t"
	EMPTY_STR							= ""
)


const (
	EXT_JSON							= ".json"
)


var (

	fPath			string

	rootCmd = &cobra.Command{
		Use: "kgc",
		Short: "kgc command line tool",
		Long: "kgc kingdom generator command line generates kingdoms with municipals and plots",
		Version: "0.1",
	}

)


const (
	ERR_FILE_EXISTS					= "File already exists, aborting...  Use '-f' to force overwrite."
)


type Kingdom struct {
	Name									string      `json:"name"`
	Capital               string      `json:"capital"`
	MaleRatio							int					`json:"maleRatio"`
	MedianAge							int					`json:"medianAge"`
	BirthRate             int					`json:"birthRate"`
	DeathRate      				int					`json:"deathRate"`
	Population            int					`json:"population"`
	Land                  int         `json:"land"`
	Wealth                int         `json:"wealth"`
	Trees                 int         `json:"trees"`
	Rocks                 int         `json:"rocks"`
	Grass                 int         `json:"grass"`
	Cows                 	int         `json:"cows"`
	TaxRate               int         `json:"taxRate"`
	ConscriptAge    			int         `json:"conscriptAge"`
	Municipals            [] string 	`json"municipals"`
}


func init() {

	cobra.OnInitialize()

	rootCmd.AddCommand(kingdomCmd)
	rootCmd.AddCommand(initCmd)

} // init


func Execute() error {
	return rootCmd.Execute()
} // Execute
