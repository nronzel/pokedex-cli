# PokedexCLI

## Description

PokedexCLI is a command-line interface tool designed for Pokemon enthusiasts.
Whether you're looking to explore the world of Pokemon, inspect individual
Pokemon, or manage your own virtual Pokedex, this tool has got you covered.

## Installation

To install PokedexCLI:

Clone the repository:

```bash
git clone https://github.com/nronzel/pokedex-cli.git
```

Navigate to the project directory:

```bash
cd pokedex-cli
```

Build the project:

```bash
go build
```

## Usage

Run the pokedex-cli executable to start the program:

```bash
./pokedex-cli
```

## Commands

Here's a list of available commands:

    help: Display help information about the available commands.
    map: Gets next page of locations
    mapb: Gets previous page of locations
    explore: Dive into the world of Pokemon and discover new ones.
    inspect [PokemonName]: Get detailed information about a specific Pokemon.
    catch [PokemonName]: Add a Pokemon to your Pokedex.
    pokedex: Display your current Pokedex.
    exit: Exit the program.

## Tests

Run the supplied tests with:

```bash
go test ./... -v
```

## License

PokedexCLI is distributed under the terms of the GPL-3.0 License. For more details, see the LICENSE file.

## Contribution

Contributions are welcome! If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are warmly welcomed.
