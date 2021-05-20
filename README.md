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

What things does it do?
-----------------------

- It can run from a flash drive, without requiring installation.
- It manages a Firefox profile, and a whole Firefox browser on Linux. It
 falls back to Chrome if Firefox is not available.
- It makes sure an I2P router is ready to work for I2P browsing, including
 launching an embedded I2P Zero router if an I2P router is not available to
 start.
- It installs browser extensions in Firefox.
- It embeds its own HTTP Proxy for browsing inside I2P
- Easily integrate Go and Webassembly based web applications as well enhance Java I2P apps
- Embed web applications(Like a blog or a WebIRC client), under human-readable domain-like aliases

It *can eventually*
-------------------

- Also browse Tor, use Tor as an outproxy
- Bring along TLS certificates for those domain-like aliases
- Update itself via bittorrent
- Behave as a distributed social network by providing social applications like groupchat,
 microblogging, feeds and even voice calls.

It also features:

 - a built-in IRC client
 - a built-in Blog
 - NoScript, uBlock Origin
 - The Snowflake plugin, to help Tor users in oppressive regimes by donating
  back the the privacy community at large.

If you're a beginner, who uses this and is interested in learning more about
I2P, I advise you to install a fully-featured I2P router by following my
detailed install guide here: https://github.com/eyedeekay/Install-Java-And-I2P-on-Windows.

Installation
------------

This is experimental software right now, in an early release. Only you will be
able to evaluate if it fits your threat model.

It should 'just work' if you download the executable from the releases page
and double-click it, but expect some rough still, especially on Windows. I
would advise you to check the sha256sum against the sha256sum in the release
description.

```Bash
# Linux
wget -O i2pfirefox https://github.com/eyedeekay/i2pfirefox/releases/download/116.0.73.099/i2pfirefox
sha256sum i2pfirefox

# Windows
wget -O i2pfirefox.exe https://github.com/eyedeekay/i2pfirefox/releases/download/116.0.73.099/i2pfirefox.exe
does-windows-even-have-sha25sum i2pfirefox.exe

# OSX
wget -O i2pfirefox-darwin https://github.com/eyedeekay/i2pfirefox/releases/download/116.0.73.099/i2pfirefox-darwin
sha256sum i2pfirefox-darwin
```

You may need to mark the downloaded file as executable, like this:

```Bash
chmod +x i2pfirefox*
```

Go developers with at least go 1.14 installed can build it from source with:

```Bash
if [ -z $GOPATH ]; then
  GOPATH=$HOME/go
fi
mkdir -p $GOPATH/src/github.com/eyedeekay/
git clone https://github.com/eyedeekay/i2pbrowser $GOPATH/src/github.com/eyedeekay/i2pbrowser
cd $GOPATH/src/github.com/eyedeekay/i2pbrowser
make build
```
