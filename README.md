# f1-cli
A command line interface to keep up with Formula 1 race events. Made using [Cobra](https://github.com/spf13/cobra) and [Termui](https://github.com/gizak/termui).

## Features
- Responsive and clear interface
- Display standings, race details and Formula 1 schedule
- Automatic caching to boost data retrieval speeds

## Requirements
- [Go](https://go.dev/)
- [Wget](https://www.gnu.org/software/wget/)

## Installation
Copy and paste the following command in your terminal
```
wget -O - https://raw.githubusercontent.com/MrVSiK/f1-cli/main/install.sh | bash
```

## Commands
| Command | Flags | Description |
| --- | --- | --- |
| f1 race | --year or -y, --round or -r | List race results based on the year and round given. |
| f1 sch | --year or -y | Show the race schedule of the given year. |
| f1 std | --year or -y, --round or -r (optional), --constructors or -c (optional) | List the driver's or constructors championship standings for the given year. |

## Acknowledgements
- This cli uses [Ergast Developer API](https://ergast.com/mrd).
