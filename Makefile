default:
	go build -o build/stache

build_all:
	mkdir -p build
	GOOS=darwin go build -o build/stache-tmp
	mv ./build/stache-tmp ./build/stache.osx
	GOOS=linux go build -o build/stache-tmp
	mv ./build/stache-tmp ./build/stache.linux
