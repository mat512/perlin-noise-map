BINARY_NAME=perlin-noise-map

build:
	go build -o ${BINARY_NAME} .

run: build
	./${BINARY_NAME}

clean:
	go clean
