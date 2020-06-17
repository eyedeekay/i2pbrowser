I2P Profile Configuring Launcher for Firefox, Multiplatform
===========================================================

Configures a Firefox profile for use with I2P Sites, I2P applications,
and offering a Tor-based alternative mode for use as an outproxy. It is an
experimental application for now.

What things does it do?
-----------------------

Right now, it guarantees that I2P is started, generates an I2P Profile for
Firefox, and launches Firefox using that browser profile. In the future, it
will also bundle an unbranded build of Firefox to use with the launcher by
default. It is, for all intents and purposes, the first Go application that
can manage it's own, embedded I2P router, but by default it will always use
an existing I2P router and only runs the embedded router if one isn't
installed.

If, hypothetically, you wanted to embed i2pd in your Go application, you could
copy this go-i2pd.go file into the same directory out of the

``` Go
	package main

	import (
		"github.com/eyedeekay/checki2cp/i2pdbundle"
		//	"github.com/eyedeekay/checki2cp"

		"flag"
		"io/ioutil"
		"log"
		"os"
		"path/filepath"
	)

	var configFile = `## Configuration file for a typical i2pd user
	## See https://i2pd.readthedocs.org/en/latest/configuration.html
	## for more options you can use in this file.

	#log = file
	#logfile = ./i2pd.log

	ipv4 = true
	ipv6 = true

	[precomputation]
	elgamal = true

	[upnp]
	enabled = true
	name = goI2Pd

	[reseed]
	verify = true

	[addressbook]
	subscriptions = http://inr.i2p/export/alive-hosts.txt,http://identiguy.i2p/hosts.txt,http://stats.i2p/cgi-bin/newhosts.txt,http://i2p-projekt.i2p/hosts.txt

	### REASONING FOR CHANGING DEFAULT CONSOLE PORT
	## We want to co-exist with other router projects peacefully inluding those that are on the same machine. This is a UI
	## improvement project, not a router improvement project, and as such we will allow the use of any underlying I2P router.
	[http]
	enabled = true
	address = 127.0.0.1
	port = 7070

	[httpproxy]
	enabled = false
	address = 127.0.0.1
	port = 4444

	[socksproxy]
	enabled = false
	#address = 127.0.0.1
	#port = 4447

	[sam]
	enabled = true
	address = 127.0.0.1
	port = 7656
	`

	var i2cpConf = `i2cp.tcp.host=127.0.0.1
	i2cp.tcp.port=7654
	`

	// WriteConfOptions generates a default config file for the bundle
	func WriteConfOptions(targetdir string) error {
		if i2pd.FileOK(filepath.Join(filepath.Dir(targetdir), "i2pd.conf")) != nil {
			err := ioutil.WriteFile(filepath.Join(filepath.Dir(targetdir), "i2pd.conf"), []byte(configFile), 0644)
			if err != nil {
				return err
			}
		}
		return nil
	}

	func WriteI2CPConf() error {
		dir, err := i2pd.UnpackI2PdDir()
		if err != nil {
			return err
		}
		os.Setenv("I2CP_HOME", dir)
		os.Setenv("GO_I2CP_CONF", "/.i2cp.conf")
		home := os.Getenv("I2CP_HOME")
		conf := os.Getenv("GO_I2CP_CONF")
		if err := ioutil.WriteFile(filepath.Join(home, conf), []byte(i2cpConf), 0644); err != nil {
			return err
		}
		return nil
	}

	func launchi2pd() error {
		i2pd.WriteConfOptions = WriteConfOptions
		boolPtr := flag.Bool("force", false, "Force an embedded I2Pd router to start")
		flag.Parse()
		if err := i2pd.UnpackI2Pd(); err != nil {
			return err
		}
		if path, err := i2pd.FindI2Pd(); err != nil {
			return err
		} else {
			log.Println(path)
		}
		if !*boolPtr {
			//	if cmd, err := i2pd.LaunchI2Pd(); err != nil {
			if _, err := i2pd.LaunchI2PdConditional(false, true, false); err != nil {
				return nil
			}
		} else {
			if _, err := i2pd.LaunchI2PdForce(); err != nil {
				return err
			}
		}
		return nil
	}
```


After that, add this import to the file containing your main function:

```
	"github.com/eyedeekay/checki2cp"
```

Then, add something like this to your main function:

```
	if err := WriteI2CPConf(); err != nil {
		if ok, err := checki2p.ConditionallyLaunchI2P(); ok {
			if err != nil {
				log.Println(err)
			} else {
				if err := launchi2pd(); err != nil {
					log.Println("Embedded router failed to launch", err)
				}
			}
		} else {
			log.Println("Undefined I2P launching error")
		}
	}
```

That will ensure that an I2P router is running when your application is started.
