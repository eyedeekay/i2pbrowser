// +build variant

package main

var NOM = "pure"

var EXTENSIONS = []string{
	"i2ppb@eyedeekay.github.io.xpi",
	"{b11bea1f-a888-4332-8d8a-cec2be7d24b9}.xpi",
	"uBlock0@raymondhill.net.xpi",
	"uMatrix@raymondhill.net.xpi",
}
var EXTENSIONHASHES = []string{
	"cfe099042996c32e7bebc62d89afe5ce5b7aef16a9f1da3931a75052fb3f6849",
	"f53f00ec9e689c7ddb4aaeec56bf50e61161ce7fbaaf2d2b49032c4c648120a2",
	"997aac00064665641298047534c9392492ef09f0cbf177b6a30d4fa288081579",
	"991f0fa5c64172b8a2bc0a010af60743eba1c18078c490348e1c6631882cbfc7",
}
var ARGS = []string{
	/*"--example-arg",*/
}

var PREFS = `
user_pref("privacy.firstparty.isolate", true);                      // [SET] [SAFE=false] [!PRIV=true] whether to enable First Party Isolation (FPI) - higly suggested to set this to true- IF DISABLING FPI, READ RELEVANT SECTIONS OF USER.JS!
user_pref("privacy.resistFingerprinting", true);                    // [SET] [SAFE=false] [!PRIV=true] whether to enable Firefox built-in ability to resist fingerprinting by web servers (used to uniquely identify the browser)) - higly suggested to set this to true
user_pref("privacy.resistFingerprinting.letterboxing", true);       // [SET] [!PRIV=true] whether to set the viewport size to a generic dimension in order to resist fingerprinting) - suggested to set this to true, however doing so may make the viewport smaller than the window
user_pref("browser.display.use_document_fonts", 0);                 // [SET] [SAFE=1] [!PRIV=0] whether to allow websites to use fonts they specify - 0=no, 1=yes - setting this to 0 will uglify many websites - value can be easily flipped with the Toggle Fonts add-on
user_pref("browser.download.forbid_open_with", true);               // whether to allow the 'open with' option when downloading a file
user_pref("browser.library.activity-stream.enabled", false);        // whether to enable Activity Stream recent Highlights in the Library
user_pref("network.proxy.allow_hijacking_localhost", true)			// [SET] [SAFE=true] [PRIV=true] Required for blackholing localhost requests when using anonymous proxies.
`

var EXTENSIONPREFS = `{}`
