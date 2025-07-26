# ASCII Art Web Dockerize

## Description
A web application that converts text to ASCII art using different banners (standard, shadow, thinkertoy), with more decoration for better look, with the addition of docker.

## Authors
Hussain Abdulrasool (habdulras)

## Usage
1. Clone the repository
2. Run the server: `go run main.go`
3. Open your browser and navigate to `http://localhost:8080`

## Usage with Docker
Run the following in Command Line.
```
git clone https://learn.reboot01.com/git/oaljamal/ascii-art-web-
cd ascii-art-web-dockerize

1. docker build -f Dockerfile -t ascii-art-web-image .
2. docker run -p 8080:8080 --detach --name ascii-art-web-container ascii-art-web-image
3. docker exec -it ascii-art-web-container /bin/bash
```
## Implementation Details
The application consists of:
- A Go HTTP server that handles requests
- HTML templates for the user interface
- ASCII art generation logic that takes text as an input and converts it to the art

The algorithm works by:
1. Reading the selected banner file (standard, shadow, or thinkertoy)
2. Processing each character in the input text
3. Building the ASCII art line by line
4. Formatting the output with proper line breaks

## Features
- Three different ASCII art styles
- Web interface with form input
- Error handling for invalid input

## References
https://hub.docker.com/_/golang

https://docs.docker.com/reference/dockerfile/