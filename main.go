//go:generate go run -tags generate gen.go extensions.go

package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	. "github.com/eyedeekay/GingerShrew/import"
	. "github.com/eyedeekay/go-fpw"
	"github.com/eyedeekay/zerobundle"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	/*In the future, we may use this as a sort of loopback for privately testing the browser
	fingerprint. At first this will be 100% user-initiated, but it may be useful to query such
	a service periodically, in order to inform the user when a fingerprint change has occurred
	and prompt them to potentially re-set their browser to it's original state.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}*/
}

func hello() error {
	server := &http.Server{Addr: "localhost:", Handler: &handler{}}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Setting up signal capturing
	//    stop := make(chan os.Signal, 1)
	//    signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	//    <-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println(err)
	}
	return nil
}

func userFind() string {
	if os.Geteuid() == 0 {
		log.Fatal("Do not run this application as root!")
	}
	if un, err := os.UserHomeDir(); err == nil {
		os.MkdirAll(filepath.Join(un, "i2p"), 0755)
		os.MkdirAll(filepath.Join(un, "i2p", "opt"), 0755)
		os.MkdirAll(filepath.Join(un, "i2p", "firefox-profiles", NOM), 0755)
		os.MkdirAll(filepath.Join(un, "i2p", "rhizome"), 0755)
		return un
	}
	return ""
}

var userdir = filepath.Join(userFind(), "i2p", "firefox-profiles", NOM)
var gingerdir = filepath.Join(userFind(), "i2p", "rhizome")

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

func firefoxMain() {
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
		if ERROR != nil {
			log.Fatal(ERROR)
		}
		defer FIREFOX.Close()
		<-FIREFOX.Done()
	} else {
		FIREFOX, ERROR := BasicFirefox(userdir, false, ARGS...)
		if ERROR != nil {
			log.Fatal(ERROR)
		}
		defer FIREFOX.Close()

		sigs := make(chan os.Signal, 1)
		done := make(chan bool, 1)

		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			sig := <-sigs
			fmt.Println()
			fmt.Println(sig)
			done <- true
		}()

		fmt.Println("awaiting signal")
		<-done
		fmt.Println("exiting")
		<-FIREFOX.Done()
	}
}

func main() {
	if err := hello(); err != nil {
		log.Fatal(err)
	}
	chromium := flag.Bool("chromium", false, "use a chromium-based browser instead of a firefox-based browser.")
	flag.Parse()
	ARGS = append(ARGS, flag.Args()...)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if runtime.GOOS == "linux" {
		if !*chromium {
			if err := UnpackTBZ(gingerdir); err != nil {
				log.Fatal("Error unpacking embedded browser")
			} else {
				os.Setenv("FIREFOX_BIN", filepath.Join(gingerdir, "gingershrew", "gingershrew"))
			}
		}
	} else {
		if LocateFirefox() == "" {
			*chromium = true
		}
	}
	if err := WriteI2CPConf(); err != nil {
		log.Println(err)
	}
	if err := zerobundle.ZeroMain(); err != nil {
		log.Println(err)
	}
	time.Sleep(time.Second * 2)
	go proxyMain(ctx)
	if !*chromium {
		firefoxMain()
	} else {
		chromiumMain()
	}
}
