/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/IMBoBx/pokedex-cli/pkg"
	"github.com/IMBoBx/pokedex-cli/structs"
	"github.com/spf13/cobra"
)

// pokemonCmd represents the pokemon command
var pokemonCmd = &cobra.Command{
	Use:   "pokemon [name/number]",
	Short: "Get details for any pokemon",
	Long: `This command fetches the details for any Pokemon using its name or its National Pokedex number. Example:
		pokedex-cli pokemon infernape
		pokedex-cli pokemon`,
	Run: func(cmd *cobra.Command, args []string) {
		pokemonFunc(args[0])
	},
}

func init() {
	rootCmd.AddCommand(pokemonCmd)
}

func pokemonFunc(id string) {
	println := fmt.Println

	var p structs.Pokemon
	var ps structs.PokemonSpecies

	pkg.FetchApi(pkg.Base+"pokemon/"+id, &p)
	pkg.FetchApi(pkg.Base+"pokemon-species/"+id, &ps)

	var (
		name        string = pkg.ToTitle(p.Name)
		description string
		types       []string
	)

	for _, t := range p.Types {
		types = append(types, pkg.ToTitle(t.Type.Name))
	}

	for _, d := range ps.FlavorTextEntries {
		if d.Language.Name == "en" {
			description = d.FlavorText
		}
	}

	println("Name:", name)
	println(description)
	println("Types:", strings.Join(types, ", "))

}
