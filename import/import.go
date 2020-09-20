//go:generate go run -tags generate gen.go extensions.go

/*
Released under the The MIT License (MIT)
see ./LICENSE
*/

package i2pbrowser

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/eyedeekay/GingerShrew/import"
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

func Main() {
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
			if os.Getenv("FIREFOX_BIN") == "" {
				if err := gingershrew.UnpackTBZ(GingerDir); err != nil {
					log.Fatal("Error unpacking embedded browser")
				} else {
					os.Setenv("LD_LIBRARY_PATH", filepath.Join(GingerDir, "lib/x86_64-linux-gnu")+","+filepath.Join(GingerDir, "usr/lib/x86_64-linux-gnu"))
					log.Println("LD_LIBRARY_PATH", filepath.Join(GingerDir, "lib/x86_64-linux-gnu")+","+filepath.Join(GingerDir, "usr/lib/x86_64-linux-gnu"))
					os.Setenv("FIREFOX_BIN", filepath.Join(GingerDir, "gingershrew", "gingershrew"))
				}
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
	if !proxyCheck() {
		go proxyMain(ctx)
	}
	if !*chromium {
		firefoxMain()
	} else {
		chromiumMain()
	}
}

func MainNoEmbeddedStuff(args []string) {
	if err := hello(); err != nil {
		log.Fatal(err)
	}
	userdir := UserDir
	for _, arg := range args {
		UserDir = filepath.Join(UserFind(), "i2p", "firefox-profiles", "webapps")
		if arg == "--app" {
			err := os.MkdirAll(filepath.Join(UserFind(), "i2p", "firefox-profiles", "webapps", "chrome"), 0755)
			if err != nil {
				log.Fatal(err)
			}
			prefs := filepath.Join(UserDir, "chrome/userChrome.css")
			if _, err := os.Stat(prefs); os.IsNotExist(err) {
				if err := ioutil.WriteFile(prefs, []byte(APPCHROME), 0644); err == nil {
					log.Println("wrote", prefs)
				} else {
					log.Fatal(err)
				}
			}
		}
	}
	chromium := false
	if args != nil {
		ARGS = append(ARGS, args...)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if LocateFirefox() == "" {
		chromium = true
	}
	if !proxyCheck() {
		go proxyMain(ctx)
	}
	if !chromium {
		firefoxLaunch()
	} else {
		chromiumMain()
	}
	UserDir = userdir
}
