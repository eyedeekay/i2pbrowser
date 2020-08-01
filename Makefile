
export GO111MODULE=on
GO111MODULE=on

VERSION=0.73
SNOW_VERSION=0.2.2
UMAT_VERSION=1.25.2
UBLO_VERSION=1.4.0
NOSS_VERSION=11.0.23
ZERO_VERSION=v1.16
LAUNCH_VERSION=$(VERSION).09

GO_COMPILER_OPTS = -a -tags netgo -ldflags '-w -extldflags "-static"'

build:
	go build $(GO_COMPILER_OPTS)

assets: clean assets.go

gen:
	go run $(GO_COMPILER_OPTS) -tags generate gen.go extensions.go

clean:
	gofmt -w -s main.go pure.go variant.go gen.go

sum: exts
	sha256sum ifox/i2ppb@eyedeekay.github.io.xpi
	sha256sum 'ifox/{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi'
	sha256sum ifox/uBlock0@raymondhill.net.xpi
	sha256sum ifox/uMatrix@raymondhill.net.xpi
	#sha256sum 'ifox/{73a6fe31-595d-460b-a920-fcc0f8843232}.xpi'

all: pure variant

pure: clean assets.go windows osx linux

windows:
	GOOS=windows go build $(GO_COMPILER_OPTS) -o i2pfirefox.exe

osx:
	#GOOS=darwin go build $(GO_COMPILER_OPTS) -o i2pfirefox-darwin

linux:
	GOOS=linux go build $(GO_COMPILER_OPTS) -o i2pfirefox

variant: clean assets.go vwindows vosx vlinux

vwindows:
	GOOS=windows go build $(GO_COMPILER_OPTS) -tags variant -o i2pfirefox-variant.exe

vosx:
	#GOOS=darwin go build $(GO_COMPILER_OPTS) -tags variant -o i2pfirefox-variant-darwin

vlinux:
	GOOS=linux go build $(GO_COMPILER_OPTS) -tags variant -o i2pfirefox-variant

release:
	gothub release -p -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "Launchers" -d "A self-configuring launcher for mixed I2P and clearnet Browsing with Firefox"
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "i2pfirefox.exe" -f "i2pfirefox.exe"
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "i2pfirefox-darwin" -f "i2pfirefox-darwin"
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "i2pfirefox" -f "i2pfirefox"

release-variant:
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "i2pfirefox-variant.exe" -f "i2pfirefox-variant.exe"
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "i2pfirefox-variant-darwin" -f "i2pfirefox-variant-darwin"
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "i2pfirefox-variant" -f "i2pfirefox-variant"

i2p-zero:
	git clone https://github.com/i2p-zero/i2p-zero.git; \
		cd i2p-zero && \
		git fetch --all --tags && \
		git checkout $(ZERO_VERSION)

zero-build: i2p-zero
	cd i2p-zero && \
		./bin/build-all-and-zip.sh

zero-copy:
	cp -rv i2p-zero ifox
