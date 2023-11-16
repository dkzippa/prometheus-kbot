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

