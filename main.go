//go:generate go run -tags generate gen.go

package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/eyedeekay/checki2cp"
	. "github.com/eyedeekay/go-fpw"
)

var userdir = "./i2pfox"

func writeExtension(val os.FileInfo, system *fs) {
	if len(val.Name()) > 3 {
		if !val.IsDir() {
			file, err := system.Open(val.Name())
			if err != nil {
				log.Fatal(err.Error())
			}
			sys := bytes.NewBuffer(nil)
			if _, err := io.Copy(sys, file); err != nil {
				log.Fatal(err.Error())
			}
			log.Println(val.Name()[len(val.Name())-3:], "== xpi")
			if val.Name()[len(val.Name())-3:] == "xpi" {
				extension := userdir + "/extensions/" + val.Name()
				if _, err := os.Stat(extension); !os.IsNotExist(err) {
					os.Remove(extension)
				}
				if err := ioutil.WriteFile(extension, sys.Bytes(), val.Mode()); err == nil {
					ARGS = append(ARGS, extension)
					log.Println("wrote", extension)
				} else {
					log.Fatal(err)
				}
			} else {
				log.Println("'"+userdir+"/"+val.Name()+"'", "ignored")
			}
		}
	} else {
		log.Println("'"+userdir+"/"+val.Name()+"'", "ignored", "contents", val.Sys())
	}
}

func writeProfile(FS *fs) {
	if embedded, err := FS.Readdir(-1); err != nil {
		log.Fatal("Extension error, embedded extension not read.", err)
	} else {
		os.MkdirAll(userdir+"/extensions", 0755)
		for _, val := range embedded {
			if val.IsDir() {
				os.MkdirAll(userdir+"/"+val.Name(), val.Mode())
			} else {
				writeExtension(val, FS)
			}
		}
	}
}

func main() {
	if ok, err := checki2p.ConditionallyLaunchI2P(); ok {
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Undefined I2P launching error")
	}
	writeProfile(FS)
	prefs := userdir + "/user.js"
	if _, err := os.Stat(prefs); os.IsNotExist(err) {
		if err := ioutil.WriteFile(prefs, []byte(PREFS), 0644); err == nil {
			log.Println("wrote", prefs)
		} else {
			log.Fatal(err)
		}
	}
	FIREFOX, ERROR := SecureExtendedFirefox(userdir, false, EXTENSIONS, EXTENSIONHASHES, ARGS...)
	if ERROR != nil {
		log.Fatal(ERROR)
	}
	defer FIREFOX.Close()
	<-FIREFOX.Done()
}
