/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/IMBoBx/pokedex-cli/pkg"
	"github.com/IMBoBx/pokedex-cli/structs"
	"github.com/spf13/cobra"
)

// evCmd represents the ev command
var evCmd = &cobra.Command{
	Use:   "ev [name/number]",
	Short: "Get Effort Values (EVs for any pokemon)",

	Run: func(cmd *cobra.Command, args []string) {
		evRunner(args[0])
	},
}

func init() {
	rootCmd.AddCommand(evCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// evCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// evCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func evRunner(id string) {
	var p structs.Pokemon
	pkg.FetchApi(pkg.Base + "pokemon/" + id, &p)

	sumEV := 0

	fmt.Printf("Effort Values for %s:\n", pkg.ToTitle(p.Name))
	
	for _, stat := range p.Stats {
		sumEV += stat.Effort
		fmt.Printf("  └── %-15s - %d\n", pkg.ToTitle(stat.Stat.Name), stat.Effort)
	}

	fmt.Println("----------------------------\n|       Total EV -", sumEV, "      |\n----------------------------")
}