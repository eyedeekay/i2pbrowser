I2P Profile Configuring Launcher for Firefox, Multiplatform
===========================================================

Configures a Firefox profile for use with I2P Sites, I2P applications, and
I2P configuration. Provides a bundled I2P Router and a pre-configured
browser. Notably, it is completely compatible with any pre-installed router
while also not requiring a pre-installed I2P router, which means that it
can be used as a one-click introduction to I2P use for beginners while *also*
being a convenient, automatic way to configure I2P Browsing for I2P users
who have been around for years. It is a project to minimize I2P browsing
misconfiguration and explore options to improve our response to various
threat models, in a way similar to Tor Browser did but with the ability to
fulfill I2P's specific needs.

If you're a beginner, who uses this and is interested in learning more about
I2P, I advise you to install a fully-featured I2P router by following my
detailed install guide here: https://github.com/eyedeekay/Install-Java-And-I2P-on-Windows.

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
wget -O i2pfirefox-darwin https://github.com/eyedeekay/i2pfirefox/releases/download/0.73.092/i2pfirefox

# Windows
wget -O i2pfirefox-darwin https://github.com/eyedeekay/i2pfirefox/releases/download/0.73.092/i2pfirefox.exe

# OSX
wget -O i2pfirefox-darwin https://github.com/eyedeekay/i2pfirefox/releases/download/0.73.092/i2pfirefox-darwin
```

Go developers with at least go 1.14 installed can build it from source with:

```Go
go get -u github.com/eyedeekay/i2pfirefox
```

What things does it do?
-----------------------

- It manages a Firefox profile, and a whole Firefox browser on Linux. It falls back to Chrome if Firefox is not
 available.
- It makes sure an I2P router is ready to work for I2P browsing, including launching an embedded I2P Zero router if an I2P router is
 not available to start.
- It installs browser extensions in Firefox.
- It embeds its own HTTP Proxy for browsing inside I2P
- Easily integrate Go and Webassembly based web applications as well enhance Java I2P apps

It *can eventually*
-------------------

- Also browse Tor, use Tor as an outproxy
- Embed web applications(Like webtorrent, webIRC), under human-readable domain-like aliases
- Bring along TLS certificates for those domain-like aliases
- Update itself via bittorrent

