# avalon-nano-cli

> Makes commanding the Avalon Nano 3 Bitcoin little home miners fun

Using `avalon-nano-cli` you can get information and set all parameters of your miners right from the command line, a script, or a cron job.

## Install for Developers

If you have Go installed on your machine and want to run the program locally, then it's just:

```
$ go install github.com/prcolaco/avalon-nano-cli@latest
```

## Download

Check the [releases](releases) page for binaries for your OS and architecture.

## Build

There's a `Makefile` for that, try it out... it should build the binaries for the various supported OS's and architectures:

```
$ make all
```

The compressed binaries will be in the `dist` folder by OS and architecture.

## Run

Just run the command and you'll get all the help:

```
$ avalon-nano-cli
Using avalon-nano-cli you can get information and set all parameters of your miners right from the command line, a script, or a cron job

Usage:
  avalon-nano-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  device      Request device information from the miner
  help        Help about any command
  led         Gets led information form the miner
  mining      Gets mining pools information from the miner
  password    Miner password commands
  reboot      Reboots the miner
  stats       Gets statistics form the miner
  version     Show program version
  wifi        Gets wifi information form the miner
  work        Gets work level information from the miner

Flags:
  -h, --help   help for avalon-nano-cli

Use "avalon-nano-cli [command] --help" for more information about a command.
```

To set the work level:

```
$ avalon-nano-cli work level -h
Sets the work level of the miner, possible level values are low, med or high

Usage:
  avalon-nano-cli work level [flags] miner_host level

Flags:
  -h, --help   help for level
```

Or to set the led mode:

```
$ avalon-nano-cli led mode -h
Sets the mode of the miner led, possible modes are off, fixed, flash, pulse or loop

Usage:
  avalon-nano-cli led mode [flags] miner_host mode

Flags:
  -h, --help   help for mode
```

Or the led color:

```
$ avalon-nano-cli led color -h
Sets the color and brightness of the miner led, brightness in percent, and RGB color in hexadecimal (HTML like)

Usage:
  avalon-nano-cli led color [flags] miner_host brightness rgb_hex

Flags:
  -h, --help   help for color
```

It's this easy... Have fun with your Avalon Nano 3! :)

### Wanna buy me a coffee?

If this project made your day, you can buy me a coffee as a token of appreciation...

BTC Address: `bc1qv3s04m40vpskea9c7f9h3ph6jy2w5c7sv8j5yj`

Thank you!
