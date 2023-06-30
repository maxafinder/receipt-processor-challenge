# receipt-processor-challenge
## Author: Max Finder
## Table of contents
1. [Description](#description)
2. [Installation steps](#installation-steps)
3. [Steps to run application](#steps-to-run-using-docker)
6. [How to run tests](#how-to-run-tests-with-go-using-make)
7. [Documentation referenced](#documentation-referenced)
---
## Description
For this project I have provided mutliple ways to run the webservice, so the user can choose the one that is easiest for them. This project uses the Gin web framework for its high performance. Since the size of the project was small I decided to kept everything in one package. Here is an overview of the project structure that I used: <br />

### Main project structure:
`/receipt-processor-challenge` <br />
&emsp;`/main.go` - The main application file where the server starts <br />
&emsp;`/routes.go` - Where the application routes are defined <br />
&emsp;`/handlers.go` - Request handlers <br />
&emsp;`/models.go` - Data models <br />
&emsp;`/middleware.go` - Middleware components <br />
&emsp;`/tools.go` - Code related to auxiliary tools that aren't directly part of the main functionality of the application but support it <br />
&emsp;`/config.go` - Configurations <br />

### Other files:
`*_test.go` files - Contains test code <br />
`.gitignore` - Contains files and directories that should not be pushed to a remote repository <br />
`Dockerfile` - Specifications to build and run the service in a Docker container <br />
`.dockerignore` - Contains files and directories that should not be built into Docker image to reduce image size <br />
`Makefile` - Contains the __make__ targets and tasks to clean, build, run, and test the application <br />
`go.mod` and `go.sum` - Files used in Go Modules <br />
### Extra notes:
I also wanted to note that this is my first time using Go, but I wanted to use if for this project to make it easier for the reviewers. It was simiar to other high-level languages that I have experience with, so I figured it wouldn't be too difficult to use it. I included links to the documentation I referred to in the last section of this document. 

---
## Installation steps
1. Install __Docker__ (assumed that you already have it).
2. Install __Go__ (assumed that you already have it).
3. [Optional] Install __make__ if it is not already installed (should be installed by default on Unix-like systems).
---
## Steps to run using Docker
### [Option 1]
1. Build the Docker image and name it __receipt-processor-challenge__. <br />
	`docker build -t receipt-processor-challenge .`
2. Run a Docker container using the image __receipt-processor-challenge__ mapping port 8080 inside the container to port 8080 on your host machine. <br />
	`docker run -p 8080:8080 receipt-processor-challenge`

## Steps to run with Go using make
### [Option 2]
1. If you have __make__ installed then run there service with Go using the following command: <br />
	`make run`
## Steps to run with Go manually
### [Option 3]
1. Build the Go project to generate binary. <br />
	`go build -o receipt-processor-challenge.out`
2. Execute the binary generated. <br />
	`./receipt-processor-challenge.out`	
---
## How to run tests with Go using make
### [Option 1]
1. If you have __make__ is installed then run the tests using the following command: <br />
	`make test`	
## How to run tests with Go manually 
### [Option 2]
1. To run the test cases with the following command: <br />
	`go test -cover -v ./...`
---
## Documentation referenced
* https://go.dev/doc/ 
* https://github.com/golang-standards/project-layout
* https://gin-gonic.com/docs/quickstart/ 