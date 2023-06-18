[![butter logo](http://microbox.rocks/assets/readme-headers/butter.png)](http://microbox.cloud/open-source#butter)
[![Build Status](https://github.com/mu-box/butter/actions/workflows/ci.yaml/badge.svg)](https://github.com/mu-box/butter/actions)

## Butter

A small, version controll based deployment service with pluggable authentication and deployment strategies.

### Status
Experimental/Unstable/Incomplete

## Routes

| Route | Description | Payload | Response |
| --- | --- | --- | --- |
| `/files?ref={ref}` | Show the names of all the files at the specific ref, or MAIN | nil | `{file contents}` |
| `/files/{file}?ref={ref}` | Get the content of the file at the specific ref, or MAIN | nil | `{file contents}` |
| `/branches` | Get the names of all branches pushed | nil | `["main"]` |
| `/commits` | Get a list of all the commits | nil | `[{"id":"sha","message":"this is a message","author":"me"}]` |
| `/commits/{commit}` | Get details about a specific commit | nil | `[{"id":"sha","author":"me","message":"this is a message","author_date":"jan","author_email":"me@me.com"}]` |

[![butter logo](http://microbox.rocks/assets/open-src/microbox-open-src.png)](http://microbox.cloud/open-source)


## TODO
build a cli
Write tests
