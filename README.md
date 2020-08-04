I2P Profile Configuring Launcher for Firefox, Multiplatform
===========================================================

Configures a Firefox profile for use with I2P Sites, I2P applications,
and offering a Tor-based alternative mode for use as an outproxy. It is an
experimental application for now.

Installation
------------

This is very experimental software right now. Use it at your own
risk.

It should 'just work' if you download the executable from the releases page
and double-click it, but expect some rough still, especially on Windows. I
would advise you to check the sha256sum against the sha256sum in the release
description.

```Bash
# Linux
wget -O i2pfirefox-darwin https://github.com/eyedeekay/i2pfirefox/releases/download/0.73.091/i2pfirefox

# Windows
wget -O i2pfirefox-darwin https://github.com/eyedeekay/i2pfirefox/releases/download/0.73.091/i2pfirefox.exe

# OSX
wget -O i2pfirefox-darwin https://github.com/eyedeekay/i2pfirefox/releases/download/0.73.091/i2pfirefox-darwin
```

Go developers with at least go 1.14 installed can build it from source with:

```Go
go get -u github.com/eyedeekay/i2pfirefox
```

What things does it do?
-----------------------

Rapidly growing in number. It manages a Firefox profile, and a whole
Firefox browser on Linux. It makes sure an I2P router is ready to work with
I2P, including launching an embedded I2P Zero router if an I2P router is
not available to start. It installs browser extensions in Firefox and
installs browser extensions.

Much more to be added soon.
