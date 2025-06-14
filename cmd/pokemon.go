/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/IMBoBx/pokedex-cli/pkg"
	"github.com/IMBoBx/pokedex-cli/structs"
	"github.com/spf13/cobra"
)

// pokemonCmd represents the pokemon command
var pokemonCmd = &cobra.Command{
	Use:   "pokemon [name/number]",
	Short: "Get details for any pokemon",
	Long: `This command fetches the details for any Pokemon using its name or its National 
Pokedex number. For example:
pokedex-cli pokemon infernape
pokedex-cli pokemon`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Too many args. Use quotes if name has a space.")
			os.Exit(1)
		}
		pokemonFunc(args[0])
	},
}

func init() {
	rootCmd.AddCommand(pokemonCmd)
}

func pokemonFunc(id string) {
	id = strings.ReplaceAll(strings.TrimSpace(id), " ", "-")

	println := fmt.Println
	printf := fmt.Printf

	var (
		p  structs.Pokemon
		ps structs.PokemonSpecies
	)

	pkg.FetchApi(pkg.Base+"pokemon/"+id, &p)
	pkg.FetchApi(pkg.Base+"pokemon-species/"+id, &ps)

	var (
		name        string = pkg.ToTitle(p.Name)
		description string
		types       []string
		evolvesFrom string = pkg.ToTitle(ps.EvolvesFromSpecies.Name)
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

	println("\nTypes:", strings.Join(types, ", "))
	if evolvesFrom != "" {
		println("Evolves from:", evolvesFrom)
	}
	println("\nStats:")

	for _, stat := range p.Stats {
		printf("  └── %-15s - %d\n", pkg.ToTitle(stat.Stat.Name), stat.BaseStat)
	}

}

// func getEvoFromTo(url, name string) (string, []string) {
// 	var (
// 		evoChain structs.EvolutionChain
// 		evoFrom  string
// 		evoTo    []string
// 	)

// 	pkg.FetchApi(url, &evoChain)

// 	for link 



// }
