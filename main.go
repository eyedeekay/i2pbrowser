//go:generate go run -tags generate gen.go

package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/eyedeekay/checki2cp"
	. "github.com/eyedeekay/go-fpw"
)

func userFind() string {
	if os.Geteuid() == 0 {
		log.Fatal("Do not run this application as root!")
	}
	if un, err := os.UserHomeDir(); err == nil {
		os.MkdirAll(filepath.Join(un, "i2p"), 0755)
		return un
	}
	return ""
}

var userdir = filepath.Join(userFind(), "/i2p/firefox-profiles")

func writeExtension(val os.FileInfo, system *fs) bool {
	var firstrun = false
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
				extension := filepath.Join(userdir, "extensions", val.Name())
				if _, err := os.Stat(extension); os.IsNotExist(err) {
					if err := ioutil.WriteFile(extension, sys.Bytes(), val.Mode()); err == nil {
						ARGS = append(ARGS, extension)
						log.Println("wrote", extension)
					} else {
						log.Fatal(err)
					}
					firstrun = true
				}
			} else {
				log.Println(filepath.Join(userdir, val.Name()), "ignored")
			}
		}
	} else {
		log.Println(filepath.Join(userdir, val.Name()), "ignored", "contents", val.Sys())
	}
	return firstrun
}

func writeProfile(FS *fs) bool {
	var firstrun = false
	if embedded, err := FS.Readdir(-1); err != nil {
		log.Fatal("Extension error, embedded extension not read.", err)
	} else {
		os.MkdirAll(filepath.Join(userdir, "extensions"), 0755)
		/*err := ioutil.WriteFile(filepath.Join(userdir, "extension-settings.json"), []byte(EXTENSIONPREFS), 0644)
		if err != nil {
			log.Fatal(err)
		}*/
		for _, val := range embedded {
			if val.IsDir() {
				os.MkdirAll(filepath.Join(userdir, val.Name()), val.Mode())
			} else {
				firstrun = writeExtension(val, FS)
			}
		}
	}
	return firstrun
}

func main() {
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
	firstrun := writeProfile(FS)
	prefs := filepath.Join(userdir, "/user.js")
	if _, err := os.Stat(prefs); os.IsNotExist(err) {
		if err := ioutil.WriteFile(prefs, []byte(PREFS), 0644); err == nil {
			log.Println("wrote", prefs)
		} else {
			log.Fatal(err)
		}
	}
	if firstrun {
		FIREFOX, ERROR := SecureExtendedFirefox(userdir, false, EXTENSIONS, EXTENSIONHASHES, ARGS...)
		defer FIREFOX.Close()
		if ERROR != nil {
			log.Fatal(ERROR)
		}
		<-FIREFOX.Done()
	} else {
		FIREFOX, ERROR := BasicFirefox(userdir, false, ARGS...)
		defer FIREFOX.Close()
		if ERROR != nil {
			log.Fatal(ERROR)
		}
		<-FIREFOX.Done()
	}
}
