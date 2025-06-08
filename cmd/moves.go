/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cmp"
	"fmt"
	"strings"

	// "os"
	"slices"

	"github.com/IMBoBx/pokedex-cli/pkg"
	"github.com/IMBoBx/pokedex-cli/structs"
	"github.com/spf13/cobra"
)

// movesCmd represents the moves command
var movesCmd = &cobra.Command{
	Use:   "moves [name/number]",
	Short: "View the level-up moveset of any Pokémon (in the latest game)",
	Long: `This command is used to check the moves that can be obtained on levelling up your Pokémon (in the latest game). For example:
pokedex-cli moves infernape
pokedex-cli moves 392`,
	Run: func(cmd *cobra.Command, args []string) {
		game := cmd.Flag("game").Value.String()
		method := cmd.Flag("method").Value.String()

		movesRunner(args[0], game, method)
	},
}

func init() {
	rootCmd.AddCommand(movesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// movesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// movesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	movesCmd.Flags().StringP("game", "g", "scarlet-violet", `Choose which game you want your moveset for (defaults to scarlet-violet)`)
	movesCmd.Flags().StringP("method", "m", "level-up", `Choose the move-learning method (defaults to level-up)`)
}

func movesRunner(id, game, method string) {
	var p structs.Pokemon
	pkg.FetchApi(pkg.Base+"pokemon/"+id, &p)

	// println := fmt.Println
	printf := fmt.Printf

	type Move struct {
		name  string
		level int
	}

	movesLearned := make([]Move, 0)
	moveset := p.Moves
	for _, move := range moveset {
		for _, versionGroup := range move.VersionGroupDetails {
			if versionGroup.VersionGroup.Name != game || versionGroup.MoveLearnMethod.Name != method {
				continue
			}

			movesLearned = append(movesLearned, Move{move.Move.Name, versionGroup.LevelLearnedAt})
			break
		}
	}

	slices.SortFunc(movesLearned, func(a, b Move) int {
		return cmp.Compare(a.level, b.level)
	})

	if len(movesLearned) == 0 {
		var vgList, methodList structs.GenericAPIResponse
		pkg.FetchApi(pkg.Base+"version-group/?limit=50", &vgList)
		pkg.FetchApi(pkg.Base+"move-learn-method/?limit=50", &methodList)

		println(pkg.Red("No learnable moves for", pkg.ToTitle(p.Name), "in selected game using selected method. (or Pokémon not available in selected game)"))
		println(pkg.Yellow("\nAvailable games:"), strings.Join(vgList.GetNames(), ", "))
		println(pkg.Yellow("\nAvailable methods:"), strings.Join(methodList.GetNames(), ", "))
		return
	}

	printf("%s's Moveset:\n", pkg.ToTitle(p.Name))
	for _, move := range movesLearned {
		printf("  └── Level %-2d - %s\n", move.level, pkg.ToTitle(move.name))
	}
}
