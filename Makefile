
export GO111MODULE=on
GO111MODULE=on

EXT_VERSION=0.73
SNOW_VERSION=0.2.2
UMAT_VERSION=1.25.2
UBLO_VERSION=1.4.0
NOSS_VERSION=11.0.23
ZERO_VERSION=v1.16
ZERO_VERSION_B=`echo $(ZERO_VERSION) | tr -d 'v.'`
LAST_VERSION=$(ZERO_VERSION_B).$(EXT_VERSION).096
LAUNCH_VERSION=$(ZERO_VERSION_B).$(EXT_VERSION).097

GO_COMPILER_OPTS = -a -tags netgo -ldflags '-w -extldflags "-static"'

echo:
	echo $(LAUNCH_VERSION)

build:
	go build $(GO_COMPILER_OPTS)

assets: fmt lib/assets.go

gen:
	go run $(GO_COMPILER_OPTS) -tags generate gen.go extensions.go

clean: fmt
	rm -f i2pfirefox*

fmt:
	gofmt -w -s *.go
	gofmt -w -s \
		lib/firefox.go \
		lib/pure.go \
		lib/pureextensions.go \
		lib/variant.go \
		lib/variantextensions.go

sum:
	sha256sum ifox/i2ppb@eyedeekay.github.io.xpi
	sha256sum 'ifox/{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi'
	sha256sum ifox/uBlock0@raymondhill.net.xpi
	sha256sum ifox/uMatrix@raymondhill.net.xpi
	#sha256sum 'ifox/{73a6fe31-595d-460b-a920-fcc0f8843232}.xpi'

all: pure variant

pure: fmt lib/assets.go windows osx linux

windows: fmt
	GOOS=windows go build $(GO_COMPILER_OPTS) -o i2pfirefox.exe

osx: fmt
	GOOS=darwin go build $(GO_COMPILER_OPTS) -o i2pfirefox-darwin

linux: fmt
	GOOS=linux go build $(GO_COMPILER_OPTS) -o i2pfirefox

variant: fmt lib/assets.go vwindows vosx vlinux

vwindows: fmt
	GOOS=windows go build $(GO_COMPILER_OPTS) -tags variant -o i2pfirefox-variant.exe

vosx: fmt
	GOOS=darwin go build $(GO_COMPILER_OPTS) -tags variant -o i2pfirefox-variant-darwin

vlinux: fmt
	GOOS=linux go build $(GO_COMPILER_OPTS) -tags variant -o i2pfirefox-variant

sumwindows=`sha256sum i2pfirefox.exe`
sumlinux=`sha256sum i2pfirefox`
sumdarwin=`sha256sum i2pfirefox-darwin`
sumvwindows=`sha256sum i2pfirefox-variant.exe`
sumvlinux=`sha256sum i2pfirefox-variant`
sumvdarwin=`sha256sum i2pfirefox-variant-darwin`

check:
	echo "$(sumwindows)"
	echo "$(sumlinux)"
	echo "$(sumdarwin)"
	echo "$(sumvwindows)"
	echo "$(sumvlinux)"
	echo "$(sumvdarwin)"

release:
	gothub release -p -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "Launchers" -d "A self-configuring launcher for mixed I2P and clearnet Browsing with Firefox"
	sed -i "s|$(LAST_VERSION)|$(LAUNCH_VERSION)|g" README.md
	sed -i "s|$(LAST_VERSION)|$(LAUNCH_VERSION)|g" Makefile
	git commit -am "Make release version $(LAUNCH_VERSION)" && git push

upload: upload-windows upload-darwin upload-linux

upload-windows:
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -l "$(sumwindows)" -n "i2pfirefox.exe" -f "i2pfirefox.exe"

upload-darwin:
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -l "$(sumdarwin)" -n "i2pfirefox-darwin" -f "i2pfirefox-darwin"

upload-linux:
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -l "$(sumlinux)" -n "i2pfirefox" -f "i2pfirefox"

upload-variant: upload-variant-windows upload-variant-darwin upload-variant-linux

upload-variant-windows:
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -l "$(sumvwindows)" -n "i2pfirefox-variant.exe" -f "i2pfirefox-variant.exe"

upload-variant-darwin:
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -l "$(sumvdarwin)" -n "i2pfirefox-variant-darwin" -f "i2pfirefox-variant-darwin"

upload-variant-linux:
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -l "$(sumvlinux)" -n "i2pfirefox-variant" -f "i2pfirefox-variant"

upload-all: upload upload-variant

release-all: release upload-all

release-pure:
	make release; true
	make linux upload-linux
	make windows upload-windows
	make osx upload-darwin

release-variant: 
	make release; true
	make vlinux upload-variant-linux
	make vwindows upload-variant-windows
	make vosx upload-variant-darwin

clean-release: clean release-pure release-variant

