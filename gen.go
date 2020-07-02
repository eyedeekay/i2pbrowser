//+build generate

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/zserge/lorca"
)

var i2ppb = []string{
	"i2ppb@eyedeekay.github.io.xpi",
	"https://github.com/eyedeekay/I2P-in-Private-Browsing-Mode-Firefox/releases/download/",
	VERSION,
}
var snowflake = []string{
	"{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi",
	"https://addons.mozilla.org/firefox/downloads/file/3519836/",
	SNOW_VERSION,
	"snowflake",
}
var noscript = []string{
	"{73a6fe31-595d-460b-a920-fcc0f8843232}.xpi",
	"https://addons.mozilla.org/firefox/downloads/file/3534184/",
	NOSS_VERSION,
	"noscript_security_suite",
}
var umatrix = []string{
	"uMatrix@raymondhill.net.xpi",
	"https://addons.mozilla.org/firefox/downloads/file/3396815/",
	UMAT_VERSION,
	"umatrix",
}
var ublock = []string{
	"uBlock0@raymondhill.net.xpi",
	"https://addons.mozilla.org/firefox/downloads/file/3519836/",
	UBLO_VERSION,
	"ublock",
}

func fetch() error {
	if err := get(i2ppb); err != nil {
		return err
	}
	if err := get(snowflake); err != nil {
		return err
	}
	if err := get(noscript); err != nil {
		return err
	}
	if err := get(umatrix); err != nil {
		return err
	}
	if err := get(ublock); err != nil {
		return err
	}
	return nil
}

func download(filepath string, url string) error {
	os.MkdirAll("ifox", 0755)
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Create the file
	out, err := os.Create("ifox/" + filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func get(extension []string) error {
	if len(extension) == 3 {
		return download(extension[0], extension[1]+extension[2])
	} else if len(extension) == 4 {
		return download(extension[0], extension[1]+extension[3]+extension[2]+"-an+fx.xpi")
	}
	return fmt.Errorf("Error fetching extension for build.")
}

/// wget -nv -c -O

func main() {
	if err := fetch(); err != nil {
		log.Fatal(err)
	}
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.
	lorca.Embed("main", "assets.go", "ifox/")
}
