![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/62fb0odpbge4mlkqkoxx.jpg)

# Whats is Dotfm?
Dotfm is a simple config file manager, which uses git as it's backend to manage and track dotfiles

# Installation
**Note:** Make sure `$GOPATH/bin` is added to your `$PATH`.

To Intsall Dotfm run the following command.
`$ go get -u github.com/thetinygoat/dotfm`

# Usage
- Intialize the dotfm repository using
`$ dotfm init`
It will initialize dotfm repository in `$HOME/.dotfm`
- Add a file to dotfm's tracker
`$dotfm add /path/to/file`
It will add the file to dotfm's tracker and create symlink to the original file.

More features coming soon :)

