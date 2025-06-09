# pokedex-cli
A CLI-based Pokédex tool, written in Go


## Features:
- Get detailed stats, types, and descriptions for any Pokémon by name or National Dex number
- View Effort Values (EVs) and base stats
- Display level-up movesets for Pokémon in the latest or selected games
- Supports multiple move-learning methods and game versions
- Clean, human-readable output


## Requirements
Go 1.18 or higher. Install [here](https://go.dev/doc/install).


## Installation
1. Clone the repository to your system
```sh
$ git clone https://github.com/IMBoBx/pokedex-cli.git
$ cd pokedex-cli
```
2. Build
- Linux/macOS
```sh
$ go build -o pokedex-cli
```
- Windows
```sh
$ go build -o pokedex-cli.exe
```
3.  *(Optional)* For global access from the Terminal
- Linux/macOS
```sh
$ sudo mv pokedex-cli /usr/local/bin/
```
- Windows: 
Add the folder to your `PATH`.



## Usage
*(Note: if added to `PATH` or `/usr/local/bin`, drop the `./` at the beginning. If using Windows, use `.\` instead of `./`)*

### Get Pokémon Details
```sh
./pokedex-cli pokemon [name|number]
```
Example:
```sh
./pokedex-cli pokemon infernape
./pokedex-cli pokemon 392
```

### Get Effort Values (EVs)

```sh
./pokedex-cli ev [name|number]
```
Example:
```sh
./pokedex-cli ev infernape
./pokedex-cli ev 392
```

### View Moveset
```sh
./pokedex-cli moves [name|number] [flags]
```
Available flags:
- `--game, -g`:  Select game version (default: `scarlet-violet`)
- `--method, -m`: Select move-learning method (default: `level-up`)

Example:
```sh
./pokedex moves infernape -g black-white
./pokedex moves 392 -m machine --game sun-moon
```

## Credits
- Powered by [PokeAPI](https://pokeapi.co/).
- Built with [Cobra CLI](https://github.com/spf13/cobra-cli).
