
VERSION=0.73

build: setup assets.go
	go build
	
assets.go:
	go run -tags generate gen.go

setup: i2pfox/extensions/i2ppb@eyedeekay.github.io.xpi


i2pfox/extensions/i2ppb@eyedeekay.github.io.xpi:
	mkdir -p i2pfox/extensions/
	wget -c -O i2pfox/extensions/i2ppb@eyedeekay.github.io.xpi https://github.com/eyedeekay/I2P-in-Private-Browsing-Mode-Firefox/releases/download/$(VERSION)/i2ppb@eyedeekay.github.io.xpi