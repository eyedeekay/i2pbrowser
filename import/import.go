//go:generate go run -tags generate gen.go extensions.go

/*
Released under the The MIT License (MIT)
see ./LICENSE
*/

package i2pbrowser

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	. "github.com/eyedeekay/go-fpw"
	. "github.com/eyedeekay/i2pbrowser/lib"
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

func proxyCheck() bool {
	conn, err := net.Dial("tcp4", "localhost:4444")
	log.Println("Doing dial check")
	if err != nil {
		return false
	}
	log.Println("Dial check true, proxy is up")
	defer conn.Close()
	return true
}

func Main(chromium, chat, blog, app bool, rundir string, args []string) {
	zerobundle.JAVA_I2P_OPT_DIR = filepath.Join(UserFind(rundir), "i2p", "rhizome")
	zerobundle.I2P_DIRECTORY_PATH = filepath.Join(UserFind(rundir), "i2p", "router")
	if err := hello(); err != nil {
		log.Fatal(err)
	}
	if rundir != "" {
		UserDir = filepath.Join(UserFind(rundir), "i2p", "firefox-profiles", NOM)
		GingerDir = filepath.Join(UserFind(rundir), "i2p", "rhizome")
		err := os.Setenv("RHZ_PROFILE_OVERRIDE", rundir)
		if err != nil {
			log.Fatal("Unable to set profile directory.", err)
		}
	}
	userdir := UserDir
	if app {
		UserDir = filepath.Join(UserFind(rundir), "i2p", "firefox-profiles", "webapps")
		err := os.MkdirAll(filepath.Join(UserFind(rundir), "i2p", "firefox-profiles", "webapps", "chrome"), 0755)
		if err != nil {
			UserDir = userdir
			log.Fatal(err)
		}
		prefs := filepath.Join(UserDir, "chrome/userChrome.css")
		if _, err := os.Stat(prefs); os.IsNotExist(err) {
			if err := ioutil.WriteFile(prefs, []byte(APPCHROME), 0644); err == nil {
				log.Println("wrote", prefs)
			} else {
				UserDir = userdir
				log.Fatal(err)
			}
		}
	}
	for _, arg := range args {
		ARGS = append(ARGS, arg)
	}
	//	ARGS = append(ARGS, flag.Args()...)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if LocateFirefox() == "" {
		chromium = true
	}
	if err := WriteI2CPConf(); err != nil {
		log.Println(err)
	}
	if err := zerobundle.ZeroMain(); err != nil {
		log.Println(err)
	}
	time.Sleep(time.Second * 2)
	if !proxyCheck() {
		go proxyMain(ctx)
	}
	if chat {
		irc("7656", userdir, false)
	}
	if blog {
		go Railroad(rundir)
	}
	if !chromium {
		firefoxMain()
	} else {
		chromiumMain()
	}
	UserDir = userdir
}
