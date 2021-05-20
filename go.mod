module github.com/eyedeekay/i2pbrowser

go 1.14

require (
	github.com/atotto/clipboard v0.1.4
	github.com/dimfeld/httptreemux v5.0.1+incompatible
	//	github.com/eyedeekay/GingerShrew v0.0.0-20210508032440-8cc02b7866b3
	github.com/eyedeekay/I2P-Configuration-for-Chromium v0.0.0-20200802063209-8973270c836e
	github.com/eyedeekay/go-fpw v0.0.0-20210510061537-1b2dcea2a5f3
	github.com/eyedeekay/httptunnel v0.0.0-20210508192603-23be3f2cdfaa
	github.com/eyedeekay/sam3 v0.32.33-0.20210313224934-b9e4186119b8
	github.com/eyedeekay/zerobundle v0.0.0-20210508181003-86893b4491fd
	github.com/getlantern/systray v1.1.0
	github.com/kabukky/journey v0.2.0
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
	github.com/webview/webview v0.0.0-20210330151455-f540d88dde4e
	github.com/zserge/lorca v0.1.9
	i2pgit.org/idk/libbrb v0.0.0-20210508203107-2ad39d4f85c6
	i2pgit.org/idk/railroad v0.0.0-20210415002900-6492d5d1dbd8
	i2pgit.org/idk/zerocontrol v0.0.0-20210415002655-ac2964c74407
)

replace github.com/zserge/lorca => github.com/eyedeekay/lorca v0.1.9-0.20200403221704-ea2ffcadfc1b

replace github.com/prologic/eris => github.com/prologic/eris v1.6.7-0.20210430033226-64d4acc46ca7

replace github.com/eyedeekay/zerobundle => ./zerobundle

//replace github.com/eyedeekay/GingerShrew => ./GingerShrew

replace github.com/khlieng/dispatch => github.com/khlieng/dispatch v0.6.5-0.20201210080608-721492cae225

replace github.com/russross/blackfriday => github.com/russross/blackfriday v1.6.0