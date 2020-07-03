// +build !variant

package main

var NOM = "pure"

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

user_pref("network.proxy.type", 1);                                 // [SET] [SAFE=1] [PRIV=1] [ALTSAFE=5] [ALTPRIV=5] 1==http, 5=socks
user_pref("network.proxy.http", "127.0.0.1");                       // [SET] 127.0.0.1
user_pref("network.proxy.http_port", 4444);                         // [SET] 4444
user_pref("network.proxy.ssl", "127.0.0.1");                        // [SET] 127.0.0.1
user_pref("network.proxy.ssl_port", 4444);                          // [SET] 4444
user_pref("network.proxy.ftp", "127.0.0.1");                        // [SET] 127.0.0.1
user_pref("network.proxy.ftp_port", 4444);                          // [SET] 4444
user_pref("network.proxy.socks", "127.0.0.1");                      // [SET] 127.0.0.1
user_pref("network.proxy.socks_port", 4444);                        // [SET] 4444
user_pref("network.proxy.share_proxy_settings", true);              // [SET] [SAFE=true] [PRIV=true] does not actually apply to http, but leave it set anyway so that we preserve the SOCKS option
user_pref("network.proxy.socks_remote_dns", true);                  // [SET] [SAFE=true] [PRIV=true]
user_pref("network.proxy.allow_hijacking_localhost", true)			// [SET] [SAFE=true] [PRIV=true] Required for blackholing localhost requests when using anonymous proxies.
`

var EXTENSIONPREFS = `{}`
