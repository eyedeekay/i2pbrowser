
export GO111MODULE=on
GO111MODULE=on

VERSION=0.73
SNOW_VERSION=0.2.2
UMAT_VERSION=1.25.2
UBLO_VERSION=1.4.0
NOSS_VERSION=11.0.23
ZERO_VERSION=v1.16
LAUNCH_VERSION=$(VERSION).09

build: setup assets.go
	go build

assets: clean setup assets.go

assets.go:
	go run -tags generate gen.go

clean:
	@echo CLEANING
	rm -rf ifox i2pfox
	gofmt -w -s main.go pure.go variant.go gen.go
	@echo CLEANED

setup: 
	@echo CLEANING
	rm -rf ifox i2pfox
	gofmt -w -s main.go pure.go variant.go gen.go
	@echo CLEANED
	make i2ppb ublock noscript
	go run -tags generate gen.go

setup-variant: 
	@echo CLEANING
	rm -rf ifox i2pfox
	gofmt -w -s main.go pure.go variant.go gen.go
	@echo CLEANED
	make i2ppb snowflake ublock umatrix
	go run -tags generate gen.go

exts: noscript i2ppb snowflake ublock umatrix

noscript: ifox/noscript@noscript.org
i2ppb: ifox/i2ppb@eyedeekay.github.io.xpi 
snowflake: ifox/snowflake@torproject.org.xpi 
ublock: ifox/uBlock0@raymondhill.net.xpi 
umatrix: ifox/uMatrix@raymondhill.net.xpi

ifox:
	mkdir -p ifox

ifox/i2ppb@eyedeekay.github.io.xpi: ifox
	wget -nv -c -O ifox/i2ppb@eyedeekay.github.io.xpi https://github.com/eyedeekay/I2P-in-Private-Browsing-Mode-Firefox/releases/download/$(VERSION)/i2ppb@eyedeekay.github.io.xpi

ifox/snowflake@torproject.org.xpi: ifox/{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi

ifox/{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi: ifox
	wget -nv -c -O 'ifox/{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi' https://addons.mozilla.org/firefox/downloads/file/3519836/snowflake-0.2.2-fx.xpi

ifox/uBlock0@raymondhill.net.xpi: ifox
	wget -nv -c -O ifox/uBlock0@raymondhill.net.xpi https://addons.mozilla.org/firefox/downloads/file/3521827/ublock_origin-$(UBLO_VERSION)-an+fx.xpi

ifox/uMatrix@raymondhill.net.xpi: ifox
	wget -nv -c -O ifox/uMatrix@raymondhill.net.xpi https://addons.mozilla.org/firefox/downloads/file/3396815/umatrix-$(UMAT_VERSION)-an+fx.xpi

ifox/{73a6fe31-595d-460b-a920-fcc0f8843232}.xpi:
	wget -nv -c -O 'ifox/{73a6fe31-595d-460b-a920-fcc0f8843232}.xpi' https://addons.mozilla.org/firefox/downloads/file/3534184/noscript_security_suite-$(NOSS_VERSION)-an+fx.xpi

ifox/noscript@noscript.org: ifox/{73a6fe31-595d-460b-a920-fcc0f8843232}.xpi

sums: exts
	sha256sum ifox/i2ppb@eyedeekay.github.io.xpi
	sha256sum 'ifox/{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi'
	sha256sum ifox/uBlock0@raymondhill.net.xpi
	sha256sum ifox/uMatrix@raymondhill.net.xpi
	sha256sum 'ifox/{73a6fe31-595d-460b-a920-fcc0f8843232}.xpi'

all: pure variant

pure: clean setup assets.go
	GOOS=windows go build -o i2pfirefox.exe
	GOOS=darwin go build -o i2pfirefox-darwin
	GOOS=linux go build -o i2pfirefox

variant: clean setup-variant assets.go
	GOOS=windows go build -tags variant -o i2pfirefox-variant.exe
	GOOS=darwin go build -tags variant -o i2pfirefox-variant-darwin
	GOOS=linux go build -tags variant -o i2pfirefox-variant

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