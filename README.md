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
	- `docker run -ti dkzippa/prometheus-kbot:v1.0.6-ca17012-amd64`
	- `docker run -ti -e TELE_TOKEN=... dkzippa/prometheus-kbot:v1.0.6-ca17012-amd64`
- test image in deployment 
	- `k create deploy kbot-test --image dkzippa/prometheus-kbot:v1.0.6-ca17012-amd64`
	

# ADD HELM
- helm create helm
- change Chart.yaml `name` to app name `prometheus-kbot`
- test deploy by helm `helm template prometheus-kbot ./helm -s templates/deployment.yaml | k apply -f -`
- commit, push
- set tag, push tags `git push --tags`
- `helm package ./helm`
- `gh release create`
- `gh release list`
- `gh release upload v1.0.6 prometheus-kbot-0.0.6.tgz`
- `helm install prometheus-kbot https://github.com/dkzippa/prometheus-kbot/releases/download/v1.0.6/prometheus-kbot-0.0.6.tgz`

- helper commands:
	- work with secret: 
		- `k create secret generic kbot --from-literal=token=...` # converts to base64 automatically
		- to recheck: `echo -n "..." | base64`
		- see if it is correct: `k get secrets kbot -o yaml`
		- `k delete secret/kbot`
	- check logs of 1st pod: 
		- `k get po && POD="pod/$(k get po -o jsonpath='{.items[0].metadata.name}')" && k describe $POD && k logs $POD -f`		
	- delete 1st deploy:
		- `k get deploy && k delete "deployment.apps/$(k get deploy -o jsonpath='{.items[0].metadata.name}')"`


# use github pipeline

- test ghcr.io/dkzippa			
	- create personal token with permissions for packages operations
	- `CR_PAT=... && echo $CR_PAT | docker login ghcr.io -u USERNAME --password-stdin`	
	- `make image push REGISTRY=ghcr.io/dkzippa APP=prometheus-kbot`
	- `docker inspect ghcr.io/dkzippa/prometheus-kbot:v1.0.6-2eea280-arm64`
	- `docker run -e TELE_TOKEN=... -ti ghcr.io/dkzippa/prometheus-kbot:v1.0.6-2eea280-arm64`

- add github worflows and actions
	- test updating version
		- `VERSION=$(git describe --tags --abbrev=0)-$(git rev-parse --short HEAD) && echo $VERSION`
		- `export TEST_VERSION2=$(git describe --tags --abbrev=0)-$(git rev-parse --short HEAD) && echo $TEST_VERSION2 && yq -i '.image.tag=strenv(TEST_VERSION2)' ./helm/values.yaml`
	- commit and push with tags 
		- `MSG="github ci/cd implemented" && git add --all && git commit -m $MSG && git push`
		- `MSG="github ci/cd implemented" && git add --all && git commit -m $MSG && git tag v.1.0.7 -m $MSG && git push && git push --tags`



