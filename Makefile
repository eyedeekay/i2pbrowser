
VERSION=0.73
SNOW_VERSION=0.2.2
LAUNCH_VERSION=$(VERSION).01

build: setup assets.go
	go build

assets.go:
	go run -tags generate gen.go

clean:
	rm -rf ifox i2pfirefox* assets.go i2pfox i2p-fox
	gofmt -w -s *.go

setup: i2ppb snowflake ublock umatrix

i2ppb: ifox/i2ppb@eyedeekay.github.io.xpi 
snowflake: ifox/snowflake@torproject.org.xpi 
ublock: ifox/uBlock0@raymondhill.net.xpi 
umatrix: ifox/uMatrix@raymondhill.net.xpi

ifox:
	mkdir -p ifox

ifox/i2ppb@eyedeekay.github.io.xpi: ifox
	wget -c -O ifox/i2ppb@eyedeekay.github.io.xpi https://github.com/eyedeekay/I2P-in-Private-Browsing-Mode-Firefox/releases/download/$(VERSION)/i2ppb@eyedeekay.github.io.xpi

ifox/snowflake@torproject.org.xpi: ifox/{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi

ifox/{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi: ifox
	wget -c -O 'ifox/{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi' https://addons.mozilla.org/firefox/downloads/file/3519836/snowflake-0.2.2-fx.xpi

ifox/uBlock0@raymondhill.net.xpi: ifox
	wget -c -O ifox/uBlock0@raymondhill.net.xpi https://addons.mozilla.org/firefox/downloads/file/3521827/ublock_origin-1.25.2-an+fx.xpi

ifox/uMatrix@raymondhill.net.xpi: ifox
	wget -c -O ifox/uMatrix@raymondhill.net.xpi https://addons.mozilla.org/firefox/downloads/file/3396815/umatrix-1.4.0-an+fx.xpi

sums: setup
	sha256sum ifox/i2ppb@eyedeekay.github.io.xpi
	sha256sum 'ifox/{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi'
	sha256sum ifox/uBlock0@raymondhill.net.xpi
	sha256sum ifox/uMatrix@raymondhill.net.xpi

all: setup assets.go
	GOOS=windows go build -o i2pfirefox.exe
	GOOS=darwin go build -o i2pfirefox-darwin
	GOOS=linux go build -o i2pfirefox

release:
	gothub release -p -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "Launchers" -d "A self-configuring launcher for mixed I2P and clearnet Browsing with Firefox"; true
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "i2pfirefox.exe" -f "i2pfirefox.exe"
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "i2pfirefox-darwin" -f "i2pfirefox-darwin"
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "i2pfirefox" -f "i2pfirefox"

linux-release:
	gothub release -p -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "Launchers" -d "A self-configuring launcher for mixed I2P and clearnet Browsing with Firefox"; true
	gothub upload -R -u eyedeekay -r "i2pfirefox" -t $(LAUNCH_VERSION) -n "i2pfirefox" -f "i2pfirefox"