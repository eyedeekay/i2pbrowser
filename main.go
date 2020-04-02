//go:generate go run -tags generate gen.go

package main

import (
	"io/ioutil"
	"log"
	"os"

	. "github.com/eyedeekay/go-fpw"
)

var EXTENSIONS = []string{"i2ppb@eyedeekay.github.io.xpi"}
var EXTENSIONHASHES = []string{"bca6f385637c76445775af6271f8e9621966283f2f648cef9db9c635b6662f6d"}
var ARGS = []string{
	/*"--example-arg",*/
}

var userdir = "i2pfox"

func main() {
	if embedded, err := FS.Readdir(0); err != nil {
		log.Fatal("Extension error, embedded extension not read.", err)
	} else {
		os.MkdirAll(userdir, FS.Mode())
		os.MkdirAll(userdir+"/extensions", FS.Mode())
		for _, val := range embedded {
			sys := val.Sys()
			if sys != nil {
                if _, err := os.Stat(userdir+"/extensions/"+val.Name()); !os.IsNotExist(err) {
                    os.Remove(userdir+"/extensions/"+val.Name())
                }
				if err := ioutil.WriteFile(userdir+"/extensions/"+val.Name(), sys.([]byte), val.Mode()); err != nil {
					log.Fatal(err)
					ARGS = append(ARGS, userdir+"/extensions/"+val.Name())
					log.Println(userdir + "/extensions/" + val.Name())
				}else{
                    log.Fatal(err)
                }
			}
		}

	}
	FIREFOX, ERROR := SecureExtendedFirefox(userdir, false, EXTENSIONS, EXTENSIONHASHES, ARGS...)
	if ERROR != nil {
		log.Fatal(ERROR)
	}
	defer FIREFOX.Close()
	<-FIREFOX.Done()
}
