## Tech stack
- Server: Go
- Web UI: Reactjs, Rsbuild

## Build
### Notes
- If you don't want to build an executable, use the other method mentioned in section <a href="#Run">"Run"</a> below
### Requirements
- Nodejs 24.12.0+
- Go 1.25.5+
### Steps
- Install packages 
    - <code>npm install --omit=dev</code>
- Build web UI static files
    - <code>npx rsbuild build</code>
- Build executable
    - <code>go build -o [output_filepath] main.go</code>

## Run
### Methods
- With executable
    - Add execute permission to the executable
    - Open terminal and run the executable from there
- Without building executable
    - Unix-like systems: <code>npx rsbuild build && go run main.go</code>
    - Windows: Basically run <code>npx rsbuild build</code> and then <code>go run main.go</code>. I just don't know how to chain commands in Windows. Please help yourself :")
### Open app
- Go to <a href="http://localhost:3000">http://localhost:3000</a>