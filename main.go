//go:generate go run -tags generate gen.go

package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	. "github.com/eyedeekay/go-fpw"
)

var EXTENSIONS = []string{"i2ppb@eyedeekay.github.io.xpi"}
var EXTENSIONHASHES = []string{"bca6f385637c76445775af6271f8e9621966283f2f648cef9db9c635b6662f6d"}
var ARGS = []string{
	/*"--example-arg",*/
}

var PREFS = `user_pref("privacy.firstparty.isolate", true);                      // [SET] [SAFE=false] [!PRIV=true] whether to enable First Party Isolation (FPI) - higly suggested to set this to true- IF DISABLING FPI, READ RELEVANT SECTIONS OF USER.JS!
user_pref("privacy.resistFingerprinting", true);                    // [SET] [SAFE=false] [!PRIV=true] whether to enable Firefox built-in ability to resist fingerprinting by web servers (used to uniquely identify the browser)) - higly suggested to set this to true
user_pref("privacy.resistFingerprinting.letterboxing", true);       // [SET] [!PRIV=true] whether to set the viewport size to a generic dimension in order to resist fingerprinting) - suggested to set this to true, however doing so may make the viewport smaller than the window
user_pref("browser.display.use_document_fonts", 0);                 // [SET] [SAFE=1] [!PRIV=0] whether to allow websites to use fonts they specify - 0=no, 1=yes - setting this to 0 will uglify many websites - value can be easily flipped with the Toggle Fonts add-on
user_pref("browser.download.forbid_open_with", true);               // whether to allow the 'open with' option when downloading a file
user_pref("browser.library.activity-stream.enabled", false);        // whether to enable Activity Stream recent Highlights in the Library`

var userdir = "i2pfox"

func writeExtension(val os.FileInfo, system http.FileSystem) {
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

func main() {
	if embedded, err := FS.Readdir(-1); err != nil {
		log.Fatal("Extension error, embedded extension not read.", err)
	} else {
		os.MkdirAll(userdir+"/extensions", FS.Mode())
		for _, val := range embedded {
			if val.IsDir() {
				if embedded, err := FS.Readdir(-1); err != nil {
					log.Fatal("Extension error, embedded extension not read.", err)
				} else {
					os.MkdirAll(userdir+"/extensions", FS.Mode())
					for _, val := range embedded {
						writeExtension(val, FS)
					}
				}
			} else {
				writeExtension(val, FS)
			}
		}
	}
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
