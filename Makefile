
VERSION=0.73

build: setup assets.go
	go build

assets.go:
	go run -tags generate gen.go

clean:
	rm -rf ifox i2pfirefox* assets.go
	gofmt -w -s *.go

setup: ifox/i2ppb@eyedeekay.github.io.xpi

ifox/i2ppb@eyedeekay.github.io.xpi:
	mkdir -p ifox/
	wget -c -O ifox/i2ppb@eyedeekay.github.io.xpi https://github.com/eyedeekay/I2P-in-Private-Browsing-Mode-Firefox/releases/download/$(VERSION)/i2ppb@eyedeekay.github.io.xpi

all: setup assets.go
	GOOS=windows go build -o i2pfirefox.exe
	GOOS=darwin go build -o i2pfirefox-darwin
	GOOS=linux go build -o i2pfirefox

release:
	gothub release -p -u eyedeekay -r "i2pfirefox" -t $(VERSION) -n "Launchers" -d "A self-configuring launcher for mixed I2P and clearnet Browsing with Firefox"; true
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(VERSION) -n "i2pfirefox.exe" -f "i2pfirefox.exe"
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(VERSION) -n "i2pfirefox-darwin" -f "i2pfirefox-darwin"
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(VERSION) -n "i2pfirefox" -f "i2pfirefox"

linux-release:
	gothub release -p -u eyedeekay -r "i2pfirefox" -t $(VERSION) -n "Launchers" -d "A self-configuring launcher for mixed I2P and clearnet Browsing with Firefox"; true
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(VERSION) -n "i2pfirefox" -f "i2pfirefox"