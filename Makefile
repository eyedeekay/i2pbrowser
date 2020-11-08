
export GO111MODULE=on
GO111MODULE=on

BIN_NAME=i2pbrowser

EXT_VERSION=`amo-version -v`
SNOW_VERSION=`amo-version -v -n torproject-snowflake`
UMAT_VERSION=`amo-version -v -n umatrix`
UBLO_VERSION=`amo-version -v -n ublock-origin`
NOSS_VERSION=`amo-version -v -n noscript`
ZERO_VERSION=v1.17
ZERO_VERSION_B=`echo $(ZERO_VERSION) | tr -d 'v.'`
LAST_VERSION=$(ZERO_VERSION_B).$(EXT_VERSION).097
LAUNCH_VERSION=$(ZERO_VERSION_B).$(EXT_VERSION).098

GO_COMPILER_OPTS = -a -tags netgo -ldflags '-w -extldflags "-static"'

echo:
	echo $(LAUNCH_VERSION)

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
	@echo 

build:
	go build $(GO_COMPILER_OPTS)

assets: fmt lib/assets.go

gen:
	go run $(GO_COMPILER_OPTS) -tags generate gen.go extensions.go

clean: fmt
	rm -f $(BIN_NAME)*

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
		import/import.go


sum:
	sha256sum ifox/i2ppb@eyedeekay.github.io.xpi
	sha256sum 'ifox/{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi'
	sha256sum ifox/uBlock0@raymondhill.net.xpi
	sha256sum ifox/uMatrix@raymondhill.net.xpi
	sha256sum 'ifox/{73a6fe31-595d-460b-a920-fcc0f8843232}.xpi'

all: pure variant

pure: fmt lib/assets.go windows osx linux

windows: fmt
	GOOS=windows go build $(GO_COMPILER_OPTS) -o $(BIN_NAME).exe

osx: fmt
	GOOS=darwin go build $(GO_COMPILER_OPTS) -o $(BIN_NAME)-darwin

linux: fmt
	GOOS=linux go build $(GO_COMPILER_OPTS) -o $(BIN_NAME)

variant: fmt lib/assets.go vwindows vosx vlinux

vwindows: fmt
	GOOS=windows go build $(GO_COMPILER_OPTS) -tags variant -o $(BIN_NAME)-variant.exe

vosx: fmt
	GOOS=darwin go build $(GO_COMPILER_OPTS) -tags variant -o $(BIN_NAME)-variant-darwin

vlinux: fmt
	GOOS=linux go build $(GO_COMPILER_OPTS) -tags variant -o $(BIN_NAME)-variant

sumwindows=`sha256sum $(BIN_NAME).exe`
sumlinux=`sha256sum $(BIN_NAME)`
sumdarwin=`sha256sum $(BIN_NAME)-darwin`
sumvwindows=`sha256sum $(BIN_NAME)-variant.exe`
sumvlinux=`sha256sum $(BIN_NAME)-variant`
sumvdarwin=`sha256sum $(BIN_NAME)-variant-darwin`

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

upload: upload-windows upload-darwin upload-linux

upload-windows:
	gothub upload -R -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -l "$(sumwindows)" -n "$(BIN_NAME).exe" -f "$(BIN_NAME).exe"

upload-darwin:
	gothub upload -R -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -l "$(sumdarwin)" -n "$(BIN_NAME)-darwin" -f "$(BIN_NAME)-darwin"

upload-linux:
	gothub upload -R -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -l "$(sumlinux)" -n "$(BIN_NAME)" -f "$(BIN_NAME)"

upload-variant: upload-variant-windows upload-variant-darwin upload-variant-linux

upload-variant-windows:
	gothub upload -R -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -l "$(sumvwindows)" -n "$(BIN_NAME)-variant.exe" -f "$(BIN_NAME)-variant.exe"

upload-variant-darwin:
	gothub upload -R -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -l "$(sumvdarwin)" -n "$(BIN_NAME)-variant-darwin" -f "$(BIN_NAME)-variant-darwin"

upload-variant-linux:
	gothub upload -R -u eyedeekay -r "$(BIN_NAME)" -t $(LAUNCH_VERSION) -l "$(sumvlinux)" -n "$(BIN_NAME)-variant" -f "$(BIN_NAME)-variant"

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

