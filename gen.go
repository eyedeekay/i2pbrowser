//+build generate

package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"os/exec"
	"time"

	"github.com/zserge/lorca"
)

var x = `<dependency>
        <dependentAssembly>
            <assemblyIdentity type="win32" name="Microsoft.VC140.CRT" version="1.0.0.0"></assemblyIdentity>
        </dependentAssembly>
    </dependency>`

var manifest = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns:asm.v3="urn:schemas-microsoft-com:asm.v3" xmlns="urn:schemas-microsoft-com:asm.v3" manifestVersion="1.0">
    <assemblyIdentity
    version="0.0.0.1"
    name="i2pfirefox.exe"
    type="win32"/>
    <description>I2P Browsing Bundle</description>
</assembly>
`

var pureExtensions = `// +build !variant

package main

var EXTENSIONS = []string{
	"i2ppb@eyedeekay.github.io.xpi",
	"uBlock0@raymondhill.net.xpi",
	"uMatrix@raymondhill.net.xpi",
}
var EXTENSIONHASHES = []string{
`

var variantExtensions = `// +build variant

package main

var NOM = "variant"

var EXTENSIONS = []string{
	"i2ppb@eyedeekay.github.io.xpi",
	"{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi",
	"uBlock0@raymondhill.net.xpi",
	"uMatrix@raymondhill.net.xpi",
}
var EXTENSIONHASHES = []string{
`

var i2ppbHash = ""
var snowflakeHash = ""
var ublockHash = ""
var umatrixHash = ""

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
	"https://addons.mozilla.org/firefox/downloads/file/3521827/",
	UBLO_VERSION,
	"ublock",
}

func variantFile() error {
	value := variantExtensions
	value += "\t\"" + i2ppbHash + "\",\n"
	value += "\t\"" + snowflakeHash + "\",\n"
	value += "\t\"" + ublockHash + "\",\n"
	value += "\t\"" + umatrixHash + "\",\n"
	value += "}\n"
	return ioutil.WriteFile("variantextensions.go", []byte(value), 0644)
}

func pureFile() error {
	value := pureExtensions
	value += "\t\"" + i2ppbHash + "\",\n"
	value += "\t\"" + ublockHash + "\",\n"
	value += "\t\"" + umatrixHash + "\",\n"
	value += "}\n"
	return ioutil.WriteFile("pureextensions.go", []byte(value), 0644)
}

func sha256sum(path string) (string, error) {
	bytes, err := ioutil.ReadFile("ifox/" + path)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(bytes)
	var s []byte
	for _, c := range sum {
		s = append(s, c)
	}
	log.Println(fmt.Sprintf("%x", sum))
	return fmt.Sprintf("%x", sum), nil
}

func fetch() error {
	if err := get(i2ppb); err != nil {
		return err
	}
	if tmp, err := sha256sum(i2ppb[0]); err != nil {
		return err
	} else {
		i2ppbHash = tmp
	}
	if err := get(snowflake); err != nil {
		return err
	}
	if tmp, err := sha256sum(snowflake[0]); err != nil {
		return err
	} else {
		snowflakeHash = tmp
	}
	/*
		//	if err := get(noscript); err != nil {
		//		return err
		//	}
		//	if tmp, err := sha256sum(noscript[0]); err != nil {
		//		return err
		//	}else{
		//		i2ppbHash = tmp
		//	}
	*/
	if err := get(umatrix); err != nil {
		return err
	}
	if tmp, err := sha256sum(umatrix[0]); err != nil {
		return err
	} else {
		umatrixHash = tmp
	}
	if err := get(ublock); err != nil {
		return err
	}
	if tmp, err := sha256sum(ublock[0]); err != nil {
		return err
	} else {
		ublockHash = tmp
	}
	return nil
}

func determinate(path string) error {
	t := time.Date(1970, time.January, 1, 1, 0, 0, 0, time.UTC)
	if err := os.Chtimes("ifox/"+path, t, t); err != nil {
		return err
	}
	return nil
}

func download(path string, url string) error {
	if _, err := os.Stat("ifox/" + path); os.IsNotExist(err) {
		os.MkdirAll("ifox", 0755)
		log.Println("fetching", path, "from", url)
		// Get the data
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		// Create the file
		out, err := os.Create("ifox/" + path)
		if err != nil {
			return err
		}
		defer out.Close()
		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		return err
	}
	return nil
}

func get(extension []string) error {
	if len(extension) == 3 {
		path := extension[1] + extension[2] + "/" + extension[0]
		err := download(extension[0], path)
		if err != nil {
			return err
		}
		return determinate(extension[0])
	} else if len(extension) == 4 {
		path := extension[1] + extension[3] + extension[2] + "-an+fx.xpi"
		err := download(extension[0], path)
		if err != nil {
			return err
		}
		return determinate(extension[0])
	}
	return fmt.Errorf("Error fetching extension for build.")
}

func main() {
	//os.RemoveAll("ifox")
	//ioutil.WriteFile("i2pfirefox.manifest", []byte(manifest), 0644)
	//if err := exec.Command("rsrc", "-manifest", "i2pfirefox.manifest", "-o", "rsrc.syso").Run(); err != nil {
	//log.Fatal(err)
	//}
	if err := fetch(); err != nil {
		log.Fatal(err)
	}
	if err := pureFile(); err != nil {
		log.Fatal(err)
	}
	if err := variantFile(); err != nil {
		log.Fatal(err)
	}
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.
	lorca.Embed("i2pfirefox", "lib/assets.go", "ifox/")
}
