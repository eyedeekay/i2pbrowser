package i2pbrowser

import (
	"i2pgit.org/idk/libbrb"

	. "github.com/eyedeekay/i2pbrowser/lib"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var brb *trayirc.BRB

func irc(sam, dir string, i2pdispatch bool) {
	var err error
	brb, err = trayirc.NewBRBFromOptions(
		trayirc.SetSAMPort(sam),
		trayirc.SetSAMHost("127.0.0.1"),
		trayirc.SetHost("127.0.0.1"),
		trayirc.SetPort("7669"),
		trayirc.SetSaveFile(true),
		trayirc.SetName("brb"),
		trayirc.SetType("server"),
		trayirc.SetFilePath(dir),
		trayirc.SetBRBConfigDirectory(dir),
		trayirc.SetBRBServerConfig("ircd.yaml"),
		trayirc.SetBRBServerName("iirc"),
		trayirc.SetHostInI2P(i2pdispatch),
		trayirc.SetInLength(3),
		trayirc.SetOutLength(3),
		trayirc.SetInVariance(0),
		trayirc.SetOutVariance(0),
		trayirc.SetInQuantity(3),
		trayirc.SetOutQuantity(3),
		trayirc.SetInBackups(1),
		trayirc.SetOutBackups(1),
		trayirc.SetEncrypt(false),
		trayirc.SetAllowZeroIn(false),
		trayirc.SetAllowZeroOut(false),
		trayirc.SetCompress(true),
		trayirc.SetReduceIdle(false),
		trayirc.SetReduceIdleTimeMs(3000000),
		trayirc.SetReduceIdleQuantity(2),
		trayirc.SetAccessListType("none"),
		trayirc.SetAccessList([]string{}),
	)
	if err != nil {
		log.Println("failed to start embedded IRC server, proceeding anyway.", err)
	}
	brb.Run()
	for len(brb.OutputAutoLink()) > len("http://localhost:7669/connect?host=?name=invisibleirc") {
		time.Sleep(time.Second * 1)
	}
	bookmarks := filepath.Join(UserDir, "/bookmarks.html")
	if _, err := os.Stat(bookmarks); os.IsNotExist(err) {
		if err := ioutil.WriteFile(bookmarks, []byte(GenerateDefaultBookmarks(UserDir, "ircd.yml")), 0644); err == nil {
			log.Println("wrote", bookmarks)
		} else {
			log.Fatal(err)
		}
	}

}
