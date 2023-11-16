version := "v1.0.2"

compile:

	@echo compile with version ${version}
	go get
	go build -x --ldflags="-X 'github.com/dkzippa/prometheus-kbot/cmd.appVersion=${version}'"

format:
	gofmt -s -w ./

clean:
	rm -f ./prometheus-kbot
