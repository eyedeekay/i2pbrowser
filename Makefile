
export GO111MODULE=on
GO111MODULE=on
GONOPROXY=github.com/eyedeekay/GingerShrew/*,github.com/eyedeekay/zerobundle/*
export GONOPROXY=github.com/eyedeekay/GingerShrew/*,github.com/eyedeekay/zerobundle/*
#GOPROXY=https://goproxy.dev,direct"
#export GOPROXY=https://goproxy.dev,direct
GOPROXY=direct
export GOPROXY=direct
GONOSUMDB=github.com/eyedeekay/GingerShrew/*,github.com/eyedeekay/zerobundle/*
export GONOSUMDB=github.com/eyedeekay/GingerShrew/*,github.com/eyedeekay/zerobundle/*

BIN_NAME=i2pbrowser

EXT_VERSION=`./amo-version.sh i2p-in-private-browsing`
SNOW_VERSION=`./amo-version.sh torproject-snowflake`
UMAT_VERSION=`./amo-version.sh umatrix`
UBLO_VERSION=`./amo-version.sh ublock-origin`
NOSS_VERSION=`./amo-version.sh noscript`
HTSV_VERSION=`./amo-version.sh https-everywhere`
ZERO_VERSION=`./get_latest_release.sh "i2p-zero/i2p-zero"`
ZERO_VERSION_B=`./get_latest_release.sh "i2p-zero/i2p-zero" | tr -d 'v.'`
PREV_VERSION=.099
PROD_VERSION=.0991
LAST_VERSION=$(ZERO_VERSION_B).$(EXT_VERSION).$(PREV_VERSION)
LAUNCH_VERSION=$(ZERO_VERSION_B).$(EXT_VERSION)$(PROD_VERSION)

GO_COMPILER_OPTS = -a -tags netgo #-ldflags '-w -extldflags "-static"'
CGO_ENABLED=1
export CGO_ENABLED=1
GOPATH=$(HOME)/go
export GOPATH=$(HOME)/go

echo: fmt
	echo $(LAUNCH_VERSION) $(ZERO_VERSION) $(EXT_VERSION) $(PROD_VERSION)

extensions.go:
	@echo "//+build generate" | tee extensions.go
	@echo "" | tee -a extensions.go
	@echo "package main" | tee -a extensions.go
	@echo "" | tee -a extensions.go
	@echo "/*" | tee -a extensions.go
	@echo "Released under the The MIT License (MIT)" | tee -a extensions.go
	@echo "see ./LICENSE" | tee -a extensions.go
	@echo "*/" | tee -a extensions.go
	@echo "var VERSION = \"$(EXT_VERSION)\"" | tee -a extensions.go
	@echo "var SNOW_VERSION = \"$(SNOW_VERSION)\"" | tee -a extensions.go
	@echo "var UMAT_VERSION = \"$(UMAT_VERSION)\"" | tee -a extensions.go
	@echo "var UBLO_VERSION = \"$(UBLO_VERSION)\"" | tee -a extensions.go
	@echo "var NOSS_VERSION = \"$(NOSS_VERSION)\"" | tee -a extensions.go
	@echo "var HTSV_VERSION = \"$(HTSV_VERSION)\"" | tee -a extensions.go
	@echo "" | tee -a extensions.go

build: clean echo deps gen
	go build $(GO_COMPILER_OPTS)

assets: fmt lib/assets.go

gen: extensions.go
	go run $(GO_COMPILER_OPTS) -tags generate gen.go extensions.go

clean: fmt
	rm -fr $(BIN_NAME)* ifox extensions.go

fmt:
	gofmt -w -s *.go
	gofmt -w -s \
		lib/firefox.go \
		lib/pure.go \
		lib/pureextensions.go \
		lib/variant.go \
		lib/variantextensions.go
	gofmt -w -s \
		import/chromium.go \
		import/firefox.go \
		import/httptunnel.go \
		import/import.go \
		import/defaultbookmarks.go

deps: GingerShrew zerobundle

GingerShrew:
	#git clone https://github.com/eyedeekay/GingerShrew

zerobundle:
	git clone https://github.com/eyedeekay/zerobundle; cd zerobundle && git pull origin master; cd ..

sum:
	sha256sum ifox/i2ppb@eyedeekay.github.io.xpi
	sha256sum 'ifox/{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi'
	sha256sum ifox/uBlock0@raymondhill.net.xpi
	sha256sum ifox/uMatrix@raymondhill.net.xpi
	sha256sum 'ifox/{73a6fe31-595d-460b-a920-fcc0f8843232}.xpi'

all: pure variant

pure: fmt lib/assets.go windows osx linux

havelibs: fmt lib/assets.go windows linux
vhavelibs: fmt lib/assets.go vwindows vlinux


windows: fmt
	GOOS=windows CC=x86_64-w64-mingw32-gcc-win32 go build $(GO_COMPILER_OPTS) -ldflags="-H windowsgui" -o $(BIN_NAME).exe

osx: fmt
	GOOS=darwin go build $(GO_COMPILER_OPTS) -o $(BIN_NAME)-osx

linux: fmt
	GOOS=linux go build $(GO_COMPILER_OPTS) -o $(BIN_NAME)

variant: fmt lib/assets.go vwindows vosx vlinux

vwindows: fmt
	GOOS=windows CC=x86_64-w64-mingw32-gcc-win32 go build $(GO_COMPILER_OPTS) -ldflags="-H windowsgui" -tags variant -o $(BIN_NAME)-variant.exe

vosx: fmt
	GOOS=darwin go build $(GO_COMPILER_OPTS) -tags variant -o $(BIN_NAME)-variant-osx

vlinux: fmt
	GOOS=linux go build $(GO_COMPILER_OPTS) -tags variant -o $(BIN_NAME)-variant

sumwindows=`sha256sum $(BIN_NAME).exe`
sumlinux=`sha256sum $(BIN_NAME)`
sumdarwin=`sha256sum $(BIN_NAME)-osx`
sumvwindows=`sha256sum $(BIN_NAME)-variant.exe`
sumvlinux=`sha256sum $(BIN_NAME)-variant`
sumvdarwin=`sha256sum $(BIN_NAME)-variant-osx`

check:
	echo "$(sumwindows)"
	echo "$(sumlinux)"
	echo "$(sumdarwin)"
	echo "$(sumvwindows)"
	echo "$(sumvlinux)"
	echo "$(sumvdarwin)"

release:
	gothub release -p -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -n "Launchers" -d "A self-configuring launcher for mixed I2P and clearnet Browsing with Firefox"
	sed -i "s|$(LAST_VERSION)|$(LAUNCH_VERSION)|g" README.md
	sed -i "s|$(LAST_VERSION)|$(LAUNCH_VERSION)|g" Makefile
	git commit -am "Make release version $(LAUNCH_VERSION)" && git push

upload: upload-windows upload-osx upload-linux

upload-windows:
	gothub upload -R -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -l "$(sumwindows)" -n "$(BIN_NAME).exe" -f "$(BIN_NAME).exe"

upload-osx:
	gothub upload -R -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -l "$(sumdarwin)" -n "$(BIN_NAME)-osx" -f "$(BIN_NAME)-osx"

upload-linux:
	gothub upload -R -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -l "$(sumlinux)" -n "$(BIN_NAME)" -f "$(BIN_NAME)"

upload-variant: upload-variant-windows upload-variant-osx upload-variant-linux

upload-variant-windows:
	gothub upload -R -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -l "$(sumvwindows)" -n "$(BIN_NAME)-variant.exe" -f "$(BIN_NAME)-variant.exe"

upload-variant-osx:
	gothub upload -R -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -l "$(sumvdarwin)" -n "$(BIN_NAME)-variant-osx" -f "$(BIN_NAME)-variant-osx"

upload-variant-linux:
	gothub upload -R -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -l "$(sumvlinux)" -n "$(BIN_NAME)-variant" -f "$(BIN_NAME)-variant"

upload-all: upload upload-variant

release-all: release upload-all

release-pure: havelibs
	make release; true
	make linux upload-linux
	make windows upload-windows
	#make osx upload-osx

release-variant: vhavelibs
	make release; true
	make vlinux upload-variant-linux
	make vwindows upload-variant-windows
	#make vosx upload-variant-osx

clean-release: clean release-pure release-variant

