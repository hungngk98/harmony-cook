## Description
- Server: Go
- Web UI: React, Rsbuild

## Build
### Requirements
- Nodejs 24.12.0+
- Go 1.25.5+
### Steps
1. Install packages
    - <code>npm install</code>
2. Build web UI static files
    - <code>npx rsbuild build</code>
3. Build executable
    - <code>go build -o [output-filepath]</code>

## Run
- Run the built executable in terminal
- Use the app at <a href="http://localhost:5000">http://localhost:5000</a>