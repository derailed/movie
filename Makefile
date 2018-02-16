VERSION := 0.0.1
IMAGE   := quay.io/imhotepio/f1

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

watch:
	@reflex -g '*.go' go test

stamp:       ## Stamp the build with a sha rev
	@git rev-parse --short HEAD > /tmp/build

linux: stamp ## Build linux exec
	@GOOS=linux go build -ldflags "-w -X main.Build=`cat /tmp/build`" -a -tags netgo -o execs/f1 f1.go

img: linux   ## Build Docker Image
	@docker build --rm -t $(IMAGE):$(VERSION) .

push: img    ## Push image to Quay
	docker push $(IMAGE)

run:         ## Run the CLI on Linux. Wat!!
	@docker run -it f1:0.0.1 ./f1 -a 4000 --log

docs:        ## Local docs...
	@godoc -http :3000 --index