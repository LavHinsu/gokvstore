[![Build](https://github.com/LavHinsu/gokvstore/actions/workflows/build.yaml/badge.svg)](https://github.com/LavHinsu/gokvstore/actions/workflows/build.yaml) [![Release](https://github.com/LavHinsu/gokvstore/actions/workflows/release.yml/badge.svg)](https://github.com/LavHinsu/gokvstore/actions/workflows/release.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/LavHinsu/gokvstore)](https://goreportcard.com/report/github.com/LavHinsu/gokvstore) [![CodeQL](https://github.com/LavHinsu/gokvstore/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/LavHinsu/gokvstore/actions/workflows/github-code-scanning/codeql)
# gokvstore
 An in memory key value store written in go, with no external imports

### why build this? 
i like golang, and this seemed like a cool project. it is maintained, if you create a issue i will look at it. 

### build
 to build this program, in the program directory run
 go build . 
 this will build a executable in your directory

### usage instructions
 there are three enpoints, which are present in the server/server.go file.
 - GET /healthcheck ( this is a probe to see if the server is running )
 - GET /key/{key} ( where you replace your key to the value in brackets, ie GET /key/thisisatestkey )
 - POST /writekey ( create a key. json body for both writekey and updatekey is describe below)
```
    {
	    "Keyname": "String",
	    "Value" : "String",
	    "E_AT" : int64 // unixtimestamp for when to expire a key
    }
```
 - PATCH /updatekey. (the E_AT timestamp is optional, if it is not given, it will not be updated/added.)
 - DELETE /deletekey/{key} (delete a key from the map)

### keystore init
 - the keystore is init with a key named foo that will expire and be deleted in 10 seconds.

### development setup
 - main branch is where development takes place
 - release action will create a release through github actions, although i've set it up as a manual trigger to not spam release whenever i push code.

