## Tech stack
- Server: Go
- Web UI: Reactjs, Rsbuild

## Build
### Requirements
- Nodejs 24.12.0+
- Go 1.25.5+
### Steps
- Step 1: Install packages for Nodejs
    - <code>npm install --omit=dev</code>
- Step 2: Build web UI static files
    - <code>npx rsbuild build</code>
- Step 3: Build executable (optional, depends on which <a href="#methods">method</a> you use to run app)
    - <code>GOOS=[target_OS] GOARCH=[target_architecture] go build -o [output_filepath] main.go</code>

## Run
### Methods
- With executable
    - Add execute permission to the executable
    - Open terminal and run the executable from there
- Without building executable
    - Run command <code>go run main.go</code>
### Open app
- Go to <a href="http://localhost:3000">http://localhost:3000</a>