![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/62fb0odpbge4mlkqkoxx.jpg)

# Whats is Dotfm?
Dotfm is a simple config file manager, which uses git as it's backend to manage and track dotfiles.

# Installation

## Download binary
### Using cURL
```console
$ curl -L  https://git.io/JfQTV -o dotfm
```
### Using wget
```console
$ wget -O dotfm https://git.io/JfQTV
```

**Note:** Make sure to add this binary to your `$PATH` otherwise you won't be able to access it globally.

## Build from source

**Note:** Make sure you have `go` installed.

- Clone this repository
- `cd` into the downloaded repository and run `go build`.


# Usage

## Initialize

```console
$ dotfm init
```

This will initialize the `dotfm` repository in `$HOME/.dotfm`

## Clone
Download an existing dotfm repository or use dotfm to track a non dotfm repository.
```console
$ dotfm clone <url>
```

## Track
This adds a file to the dotfm tracker.
```console
$ dotfm track /path/to/file
```

## List
Lists tracked files.
```console
$ dotfm list
```
## Status
Shows the status of the repository.
```console
$ dotfm status
```

## Stage
Stage files to be commited.
```console
$ dotfm add <filename>
```

## Commit
Commit your changes.
```console
$ dotfm commit
```
This will open your get editor.

## Remote

### Add
Add a new Remote.
```console
$ dotfm remote add <remote name> <url>
```
### Remove
Remove an existing Remote.
```console
$ dotfm remote remove <remote name>
```
### List
List remotes.
```console
$ dotfm remote list
```

## Push
Push local changes to remote repository.
```console
$ dotfm push <remote name> <branch>
```

Example:
```console
$ dotfm push origin master
```

## Sync
Sync local repository with remote.
```console
$ dotfm sync <remote name> <branch>
```
Example:
```console
$ dotfm sync origin master
```

## Environments
Environments are nothing but git branches but they can be very powerful. You can create different environments for different machines and use them accordingly.

## Create an Environment
```console
$ dotfm env create <env name>
```
## List Environments
```console
$ dotfm env list
```
## Switch to an Environment
```console
$ dotfm env switch <env name>
```
## Delete an Environment
```console
$ dotfm env delete <env name>
```
