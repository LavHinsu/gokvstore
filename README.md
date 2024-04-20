[![Build](https://github.com/LavHinsu/gokvstore/actions/workflows/buildandrelease.yml/badge.svg)](https://github.com/LavHinsu/gokvstore/actions/workflows/buildandrelease.yml) [![Release](https://github.com/LavHinsu/gokvstore/actions/workflows/release.yml/badge.svg)](https://github.com/LavHinsu/gokvstore/actions/workflows/release.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/LavHinsu/gokvstore)](https://goreportcard.com/report/github.com/LavHinsu/gokvstore)
# gokvstore
 An in memory key value store written in go, with no external imports

### why build this? 
i like golang, and this seemed like a cool project. it is maintained, if you create a issue i will look at it. 

### build
 to build this program, in the program directory run,
 go build . 
 this will build a executable in your directory

### usage instructions
 there are three enpoints, which are present in the server/server.go file.
 - GET /healthcheck ( this is a probe to see if the server is running )
 - GET /key/{key} ( where you replace your key to the value in brackets, ie GET /key/thisisatestkey )
 - POST /writekey ( create a key. json body for both writekey and updatekey is describe below)
 - PATCH /updatekey ( update a key. )
 {
    "Keyname" : "Thisisakey",
    "Value": "Thisisavalue"
 }

### development setup
 - main branch is where development takes place
 - release branch build will create a release through github actions.

### todo
 - Add capability to set keys to expire
 - Add created_at automatically to all keys
 - Add capability to log ip of the requester
 - Write tests
