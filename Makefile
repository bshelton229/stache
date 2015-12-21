default:
	go build -o build/stache

build_all:
	mkdir -p build
	# Darwin 64
	GOOS=darwin GOARCH=amd64 go build -o build/stache-tmp
	mv ./build/stache-tmp ./build/stache.darwin.amd64
	# Linux 64
	GOOS=linux GOARCH=amd64 go build -o build/stache-tmp
	mv ./build/stache-tmp ./build/stache.linux.amd64
	# Darwin 386
	GOOS=darwin GOARCH=386 go build -o build/stache-tmp
	mv ./build/stache-tmp ./build/stache.darwin.386
	# Linux 386
	GOOS=linux GOARCH=386 go build -o build/stache-tmp
	mv ./build/stache-tmp ./build/stache.linux.386
