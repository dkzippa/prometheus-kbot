# Kbot

#### live url: 
    t.me/dkzippa_bot


##### optional:
- use https://github.com/cosmtrek/air
- change/use `Makefile`
- use `.env` file with `godotenv`


##### mandatory:
- create repo `github.com/dkzippa/prometheus-kbot`
- `go mod init github.com/dkzippa/prometheus-kbot`
- `go get` when have dependencies
- `go install github.com/spf13/cobra-cli@latest`
- `cobra-cli init`
- `cobra-cli add version`
- `cobra-cli add kbot`
- `go run main.go`
- `go run main.go version`
- compile `make compile` or `go build -x --ldflags="-X 'github.com/dkzippa/prometheus-kbot/cmd.appVersion=v1.0.0'"`
- `./prometheus-kbot version`
- use `gopkg.in` for telebot
- `gofmt -s -w ./`
- `go get`
- `go build -x --ldflags="-X 'github.com/dkzippa/prometheus-kbot/cmd.appVersion=v1.0.1'"`
- create new bot 'kbot' 'dkzippa_kbot' in godfather 
- `read -s  TELE_TOKEN` 
- (command+v)
- `export TELE_TOKEN`
- go to [t.me/dkzippa_bot](https://t.me/dkzippa_bot) and test


# Build in Docker image for different platforms
- add Makefile
	- build app for diff platforms: linux, macos, windows, arm
	- add clean, format for the kbot app
	- add docker image building

- add Dockerfile
	- use multistage build
	- create image from goland and build the app with `make <platform>`
	- create image from scratch and run kbot app in it(copy from dev image)
- add .dockerignore
- commit, push git codebase
- don't forget to add tags
- push image to registry(dockerhub login)
- test all on other stages



# RUN in docker or kubernetes
- change `ENTRYPOINT ["./kbot"]` to `ENTRYPOINT ["./kbot", "prometheusKbot"]` to have it running
- pass `TELE_TOKEN` as env
- commit, push
- add version tag like `v1.0.5`, push tags
- test image in docker 
	- `docker run -ti dkzippa/prometheus-kbot:v1.0.5-50482dd-amd64`
	- `docker run -ti -e TELE_TOKEN=... dkzippa/prometheus-kbot:v1.0.5-50482dd-amd64`
- test image in deployment 
	- `k create deploy kbot-test --image dkzippa/prometheus-kbot:v1.0.5-50482dd-amd64`
	

