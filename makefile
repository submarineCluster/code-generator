.PHONY: build

build:
	go fmt ./...
	goimports -w .
	golint ./...
	pkger -include /template
	go-xray -d .
	go build -ldflags "-X git.code.oa.com/tencent_abtest/code-generator/ldflags.Version=1.0.0 \
	-X 'git.code.oa.com/tencent_abtest/code-generator/ldflags.GOVersion=`go version`' \
	-X 'git.code.oa.com/tencent_abtest/code-generator/ldflags.BuildTime=`date`'" .
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