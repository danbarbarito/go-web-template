# go-web-template

### A template for creating web applications in Go


## What's Included

`httprouter` for all your routing needs

`mgo` for all your MongoDB needs

`pongo2` for all your templating needs

## Setup

`go get github.com/danielbarbarito/go-web-template`

`cd $GOPATH/src/github.com/danielbarbarito/`

`cp -r go-web-template <Your App Name>`

`cd <Your App Name>`

And then change all the import paths in all the `.go` files by replacing any instance of `go-web-template` with `Your App Name`


## Running the server

If you want to simply run the server, just run `go run App.go`.

This has it's downsides, however. You have to stop and restart the server every time you change a file in order to see your changes.

If you want an auto restarting server, use CompileDaemon.

Run `go get github.com/danielbarbarito/CompileDaemon` and then just run `CompileDaemon -command="./App"` in your apps folder. This watches for changes in your files, and then restarts the server when there is a change.
