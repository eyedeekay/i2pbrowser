package main

import (
	"github.com/eyedeekay/checki2cp/i2pdbundle"
	
	"path/filepath"
	"io/ioutil"
	"log"
	"flag"
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
		if _, err := i2pd.LaunchI2Pd(); err != nil {
			return nil
		}
	} else {
		if _, err := i2pd.LaunchI2PdForce(); err != nil {
			return err
		}
	}
	return nil
}
