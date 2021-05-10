module github.com/eyedeekay/i2pbrowser

go 1.14

require (
	github.com/eyedeekay/GingerShrew v0.0.0-20210508032440-8cc02b7866b3
	github.com/eyedeekay/I2P-Configuration-for-Chromium v0.0.0-20200802063209-8973270c836e
	github.com/eyedeekay/go-fpw v0.0.0-20210510061537-1b2dcea2a5f3
	github.com/eyedeekay/httptunnel v0.0.0-20210508192603-23be3f2cdfaa
	github.com/eyedeekay/zerobundle v0.0.0-20210508181003-86893b4491fd
	github.com/prometheus/client_golang v1.1.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/zserge/lorca v0.1.9
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	i2pgit.org/idk/libbrb v0.0.0-20210508203107-2ad39d4f85c6
	i2pgit.org/idk/zerocontrol v0.0.0-20210415002655-ac2964c74407 // indirect
)

replace github.com/zserge/lorca => github.com/eyedeekay/lorca v0.1.9-0.20200403221704-ea2ffcadfc1b

replace github.com/prologic/eris => github.com/prologic/eris v1.6.7-0.20210430033226-64d4acc46ca7

replace github.com/eyedeekay/zerobundle => ./zerobundle

replace github.com/eyedeekay/GingerShrew => ./GingerShrew

replace github.com/khlieng/dispatch => github.com/khlieng/dispatch v0.6.5-0.20201210080608-721492cae225
