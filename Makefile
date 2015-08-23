compile:
	CGO_ENABLED=0 GOOS=linux godep go build -a -installsuffix cgo -o main .
osx:
	CGO_ENABLED=0 GOOS=darwin godep go build -a -installsuffix cgo -o main .
