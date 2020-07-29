package main

import (
	"github.com/eyedeekay/zerobundle"
	"github.com/eyedeekay/httptunnel"
  "github.com/eyedeekay/httptunnel/multiproxy"

"context"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"net"
	"net/http"
	"time"
)

var i2cpConf = `i2cp.tcp.host=127.0.0.1
i2cp.tcp.port=7654
`

var (
	tunnelName           = flag.String("service-name", "sam-browser-proxy", "Name of the service(can be anything)")
	aggressiveIsolation  = flag.Bool("mode-aggressive", false, "Create a new client for every single eepSite, rather than making use of contextual identities")
	controlPortString        = flag.String("control-addr", "127.0.0.1:7951", ":port of the SAM bridge")
	proxyPortString        = flag.String("proxy-addr", "127.0.0.1:4444", ":port of the SAM bridge")
	samHostString        = flag.String("bridge-host", "127.0.0.1", "host: of the SAM bridge")
	samPortString        = flag.String("bridge-port", "7656", ":port of the SAM bridge")
	watchProfiles        = flag.String("watch-profiles", "~/.mozilla/.firefox.profile.i2p.default/user.js,~/.mozilla/.firefox.profile.i2p.debug/user.js", "Monitor and control these Firefox profiles")
	destfile             = flag.String("dest-file", "invalid.tunkey", "Use a long-term destination key")
	debugConnection      = flag.Bool("conn-debug", true, "Print connection debug info")
	inboundTunnelLength  = flag.Int("in-tun-length", 2, "Tunnel Length(default 3)")
	outboundTunnelLength = flag.Int("out-tun-length", 2, "Tunnel Length(default 3)")
	inboundTunnels       = flag.Int("in-tunnels", 2, "Inbound Tunnel Count(default 2)")
	outboundTunnels      = flag.Int("out-tunnels", 2, "Outbound Tunnel Count(default 2)")
	inboundBackups       = flag.Int("in-backups", 1, "Inbound Backup Count(default 1)")
	outboundBackups      = flag.Int("out-backups", 1, "Inbound Backup Count(default 1)")
	inboundVariance      = flag.Int("in-variance", 0, "Inbound Backup Count(default 0)")
	outboundVariance     = flag.Int("out-variance", 0, "Inbound Backup Count(default 0)")
	dontPublishLease     = flag.Bool("no-publish", true, "Don't publish the leaseset(Client mode)")
	encryptLease         = flag.Bool("encrypt-lease", false, "Encrypt the leaseset(default false, inert)")
	reduceIdle           = flag.Bool("reduce-idle", false, "Reduce tunnels on extended idle time")
	useCompression       = flag.Bool("use-compression", true, "Enable gzip compression")
	reduceIdleTime       = flag.Int("reduce-idle-time", 2000000, "Reduce tunnels after time(Ms)")
	reduceIdleQuantity   = flag.Int("reduce-idle-tunnels", 1, "Reduce tunnels to this level")
	runCommand           = flag.String("run-command", "", "Execute command using the *_PROXY environment variables")
	runArguments         = flag.String("run-arguments", "", "Pass arguments to run-command")
	suppressLifetime     = flag.Bool("suppress-lifetime-output", false, "Suppress \"Tunnel lifetime\" output")
)

func WriteI2CPConf() error {
	dir, err := zerobundle.UnpackZeroDir()
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

func proxyMain(ctx context.Context) {
  profiles := strings.Split(*watchProfiles, ",")

	srv := &http.Server{
		ReadTimeout:  600 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         *proxyPortString,
	}
	var err error
	srv.Handler, err = i2pbrowserproxy.NewHttpProxy(
		i2pbrowserproxy.SetHost(*samHostString),
		i2pbrowserproxy.SetPort(*samPortString),
		i2pbrowserproxy.SetProxyAddr(*proxyPortString),
		i2pbrowserproxy.SetControlAddr(*controlPortString),
		i2pbrowserproxy.SetDebug(*debugConnection),
		i2pbrowserproxy.SetInLength(uint(*inboundTunnelLength)),
		i2pbrowserproxy.SetOutLength(uint(*outboundTunnelLength)),
		i2pbrowserproxy.SetInQuantity(uint(*inboundTunnels)),
		i2pbrowserproxy.SetOutQuantity(uint(*outboundTunnels)),
		i2pbrowserproxy.SetInBackups(uint(*inboundBackups)),
		i2pbrowserproxy.SetOutBackups(uint(*outboundBackups)),
		i2pbrowserproxy.SetInVariance(*inboundVariance),
		i2pbrowserproxy.SetOutVariance(*outboundVariance),
		i2pbrowserproxy.SetUnpublished(*dontPublishLease),
		i2pbrowserproxy.SetReduceIdle(*reduceIdle),
		i2pbrowserproxy.SetCompression(*useCompression),
		i2pbrowserproxy.SetReduceIdleTime(uint(*reduceIdleTime)),
		i2pbrowserproxy.SetReduceIdleQuantity(uint(*reduceIdleQuantity)),
		i2pbrowserproxy.SetKeysPath(*destfile),
		i2pbrowserproxy.SetProxyMode(*aggressiveIsolation),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctrlsrv := &http.Server{
		ReadHeaderTimeout: 600 * time.Second,
		WriteTimeout:      600 * time.Second,
		Addr:              *controlPortString,
	}
	ctrlsrv.Handler, err = i2phttpproxy.NewSAMHTTPController(profiles, srv)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for sig := range c {
			if sig == os.Interrupt {
				srv.Handler.(*i2pbrowserproxy.SAMMultiProxy).Close()
				srv.Shutdown(ctx)
				ctrlsrv.Shutdown(ctx)
			}
		}
	}()

  cln, err := net.Listen("tcp4", *controlPortString)
  if err != nil {
    log.Fatal(err)
  }
	go func() {
		log.Println("Starting control server on", cln.Addr())
		if err := ctrlsrv.Serve(cln); err != nil {
			if err == http.ErrServerClosed {
				return
			}
			log.Fatal("Serve:", err)
		}
		log.Println("Stopping control server on", cln.Addr())
	}()

  ln, err := net.Listen("tcp4", *controlPortString)
  if err != nil {
    log.Fatal(err)
  }
	go func() {
		log.Println("Starting proxy server on", ln.Addr())
		if err := srv.Serve(ln); err != nil {
			if err == http.ErrServerClosed {
				return
			}
			log.Fatal("Serve:", err)
		}
		log.Println("Stopping proxy server on", ln.Addr())
	}()

	go counter()

	<-ctx.Done()
}

func counter() {
	var x int
	for {
		if !*suppressLifetime {
			log.Println("Identity is", x, "minute(s) old")
			time.Sleep(1 * time.Minute)
			x++
		}
	}
}
