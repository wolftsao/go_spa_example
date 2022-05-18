# Go SPA example
This is a simple project to demo how to serve Vue SPA with Go standard library

### Requirement
1. Linux environment 
2. npm (optionally Vue CLI)
3. Go compile (at lease 1.16)
4. Docker (optional, if you like to run in docker container)

### Steps to run this program
> The **main** branch use go embed to serve frontend bundle
>
> There's another branch **noembed** which serve fronted with normal filesystem
1. go to ui directory, install npm packages then build:
    ```bash
    cd ui
    npm install
    npm run build
    ```
    then there's frontend bundle under **ui/dist**
2. go back to project root then build Go binary:
    ```bash
    cd ../
    go build -o go-spa
    # or just run go run main.go
    ```
3. if you like to try docker container, just build docker image then run:
    ```bash
    # also in project root
    docker build -t go-spa:demo .
    docker run --rm -p "8888:8888" go-spa:demo
    ```
4. after running the program, you will have below resources:
    * backend:
        * /api/v1/greeting
    * frontend:
        * /
        * /about
        * anything else will show NotFound page