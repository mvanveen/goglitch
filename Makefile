glitch:
	GOPATH=$(PWD) go get code.google.com/p/draw2d/draw2d
	GOPATH=$(PWD) go build .
clean:
	rm -f *.png *.gif
