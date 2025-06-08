/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pokecli",
	Short: "A CLI-based Pokedex tool, written in Go",
	Long: `pokecli is a command-line Pokédex tool that lets you fetch and display information about any Pokémon directly from your terminal.

Features:
  - Get detailed stats, types, and descriptions for any Pokémon by name or National Dex number
  - View Effort Values (EVs) and base stats
  - Display level-up movesets for Pokémon in the latest or selected games
  - Supports multiple move-learning methods and game versions
  - Clean, human-readable output

Examples:
  pokecli pokemon pikachu
  pokecli pokemon 25
  pokecli ev garchomp
  pokecli moves charizard -g sword-shield -m level-up

Powered by PokeAPI.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pokedex-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
