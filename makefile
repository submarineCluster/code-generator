.PHONY: build

build:
	go fmt ./...
	goimports -w .
	golint ./...
	pkger -include /template
	go-xray -d .
	go build -ldflags "-X github.com/submarineCluster/code-generator/ldflags.Version=1.0.0 \
	-X 'github.com/submarineCluster/code-generator/ldflags.GOVersion=`go version`' \
	-X 'github.com/submarineCluster/code-generator/ldflags.BuildTime=`date`'" .
#	./code-generator clean -n demo
#	./code-generator resource -n demo -s mongo --apiServer true
	go vet ./...
	rm pkged.go

clean:
	rm -rf params
	rm -rf model
	rm -rf service
	rm -rf controller
	rm -rf dao