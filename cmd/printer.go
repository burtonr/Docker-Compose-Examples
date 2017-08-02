package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// whereAmICmd represents the version command
var printerCmd = &cobra.Command{
	Use:   "printer",
	Short: "Show the location of the demo-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The variable TEST_ENV_ITEM was: " + os.Getenv("TEST_ENV_ITEM"))
		fmt.Println("The variable ANOTHER_ENV_ITEM was: " + os.Getenv("ANOTHER_ENV_ITEM"))
		if os.Getenv("PROD_ENV_ITEM") != "" {
			fmt.Println("The variable PROD_ENV_ITEM was: " + os.Getenv("PROD_ENV_ITEM"))
		}
	},
}

func init() {
	RootCmd.AddCommand(printerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
