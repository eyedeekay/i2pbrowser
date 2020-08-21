// +build !variant

package i2pfirefox

/*
Released under the The MIT License (MIT)
see ./LICENSE
*/

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


//These come from pyllyukko user.js.
/*
The MIT License (MIT)

Copyright (c) 2016 pyllyukko

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
// PREF: Disable Mozilla telemetry/experiments
// https://wiki.mozilla.org/Platform/Features/Telemetry
// https://wiki.mozilla.org/Privacy/Reviews/Telemetry
// https://wiki.mozilla.org/Telemetry
// https://www.mozilla.org/en-US/legal/privacy/firefox.html#telemetry
// https://support.mozilla.org/t5/Firefox-crashes/Mozilla-Crash-Reporter/ta-p/1715
// https://wiki.mozilla.org/Security/Reviews/Firefox6/ReviewNotes/telemetry
// https://gecko.readthedocs.io/en/latest/browser/experiments/experiments/manifest.html
// https://wiki.mozilla.org/Telemetry/Experiments
// https://support.mozilla.org/en-US/questions/1197144
// https://firefox-source-docs.mozilla.org/toolkit/components/telemetry/telemetry/internals/preferences.html#id1
user_pref("toolkit.telemetry.enabled",				false);
user_pref("toolkit.telemetry.unified",				false);
user_pref("toolkit.telemetry.archive.enabled",			false);
user_pref("experiments.supported",				false);
user_pref("experiments.enabled",				false);
user_pref("experiments.manifest.uri",				"");

// PREF: Disallow Necko to do A/B testing
// https://trac.torproject.org/projects/tor/ticket/13170
user_pref("network.allow-experiments",				false);

// PREF: Disable sending Firefox crash reports to Mozilla servers
// https://wiki.mozilla.org/Breakpad
// http://kb.mozillazine.org/Breakpad
// https://dxr.mozilla.org/mozilla-central/source/toolkit/crashreporter
// https://bugzilla.mozilla.org/show_bug.cgi?id=411490
// A list of submitted crash reports can be found at about:crashes
user_pref("breakpad.reportURL",					"");

// PREF: Disable sending reports of tab crashes to Mozilla (about:tabcrashed), don't nag user about unsent crash reports
// https://hg.mozilla.org/mozilla-central/file/tip/browser/app/profile/firefox.js
user_pref("browser.tabs.crashReporting.sendReport",		false);
user_pref("browser.crashReports.unsubmittedCheck.enabled",	false);

// PREF: Disable FlyWeb (discovery of LAN/proximity IoT devices that expose a Web interface)
// https://wiki.mozilla.org/FlyWeb
// https://wiki.mozilla.org/FlyWeb/Security_scenarios
// https://docs.google.com/document/d/1eqLb6cGjDL9XooSYEEo7mE-zKQ-o-AuDTcEyNhfBMBM/edit
// http://www.ghacks.net/2016/07/26/firefox-flyweb
user_pref("dom.flyweb.enabled",					false);

// PREF: Disable the UITour backend
// https://trac.torproject.org/projects/tor/ticket/19047#comment:3
user_pref("browser.uitour.enabled",				false);


// PREF: Disable collection/sending of the health report (healthreport.sqlite*)
// https://support.mozilla.org/en-US/kb/firefox-health-report-understand-your-browser-perf
// https://gecko.readthedocs.org/en/latest/toolkit/components/telemetry/telemetry/preferences.html
user_pref("datareporting.healthreport.uploadEnabled",		false);
user_pref("datareporting.healthreport.service.enabled",		false);
user_pref("datareporting.policy.dataSubmissionEnabled",		false);
// "Allow Firefox to make personalized extension recommendations"
user_pref("browser.discovery.enabled",				false);

// PREF: Disable Heartbeat  (Mozilla user rating telemetry)
// https://wiki.mozilla.org/Advocacy/heartbeat
// https://trac.torproject.org/projects/tor/ticket/19047
user_pref("browser.selfsupport.url",				"");

// PREF: Disable Firefox Hello (disabled) (Firefox < 49)
// https://wiki.mozilla.org/Loop
// https://support.mozilla.org/t5/Chat-and-share/Support-for-Hello-discontinued-in-Firefox-49/ta-p/37946
// NOTICE-DISABLED: Firefox Hello requires setting "media.peerconnection.enabled" and "media.getusermedia.screensharing.enabled" to true, "security.OCSP.require" to false to work.
//user_pref("loop.enabled",		false);

// PREF: Disable Firefox Hello metrics collection
// https://groups.google.com/d/topic/mozilla.dev.platform/nyVkCx-_sFw/discussion
user_pref("loop.logDomains",					false);


// PREF: Disable "Recommended by Pocket" in Firefox Quantum
user_pref("browser.newtabpage.activity-stream.feeds.section.topstories",	false);

// PREF: Limit the connection keep-alive timeout to 15 seconds (disabled)
// https://github.com/pyllyukko/user.js/issues/387
// http://kb.mozillazine.org/Network.http.keep-alive.timeout
// https://httpd.apache.org/docs/current/mod/core.html#keepalivetimeout
//user_pref("network.http.keep-alive.timeout",			15);

// PREF: Disable prefetching of <link rel="next"> URLs
// http://kb.mozillazine.org/Network.prefetch-next
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Link_prefetching_FAQ#Is_there_a_preference_to_disable_link_prefetching.3F
user_pref("network.prefetch-next",				false);

// PREF: Disable DNS prefetching
// http://kb.mozillazine.org/Network.dns.disablePrefetch
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Controlling_DNS_prefetching
user_pref("network.dns.disablePrefetch",			true);
user_pref("network.dns.disablePrefetchFromHTTPS",		true);

// PREF: Disable the predictive service (Necko)
// https://wiki.mozilla.org/Privacy/Reviews/Necko
user_pref("network.predictor.enabled",				false);

// PREF: Reject .onion hostnames before passing the to DNS
// https://bugzilla.mozilla.org/show_bug.cgi?id=1228457
// RFC 7686
user_pref("network.dns.blockDotOnion",				true);

// PREF: Disable search suggestions in the search bar
// http://kb.mozillazine.org/Browser.search.suggest.enabled
user_pref("browser.search.suggest.enabled",			false);

// PREF: Disable "Show search suggestions in location bar results"
user_pref("browser.urlbar.suggest.searches",			false);
// PREF: When using the location bar, don't suggest URLs from browsing history
user_pref("browser.urlbar.suggest.history",			false);

// PREF: Disable SSDP
// https://bugzilla.mozilla.org/show_bug.cgi?id=1111967
user_pref("browser.casting.enabled",				false);

// PREF: Disable automatic downloading of OpenH264 codec
// https://support.mozilla.org/en-US/kb/how-stop-firefox-making-automatic-connections#w_media-capabilities
// https://andreasgal.com/2014/10/14/openh264-now-in-firefox/
user_pref("media.gmp-gmpopenh264.enabled",			false);
user_pref("media.gmp-manager.url",				"");


// PREF: Disable downloading homepage snippets/messages from Mozilla
// https://support.mozilla.org/en-US/kb/how-stop-firefox-making-automatic-connections#w_mozilla-content
// https://wiki.mozilla.org/Firefox/Projects/Firefox_Start/Snippet_Service
user_pref("browser.aboutHomeSnippets.updateUrl",		"");

// PREF: Never check updates for search engines
// https://support.mozilla.org/en-US/kb/how-stop-firefox-making-automatic-connections#w_auto-update-checking
user_pref("browser.search.update",				false);

// PREF: Disable automatic captive portal detection (Firefox >= 52.0)
// https://support.mozilla.org/en-US/questions/1157121
user_pref("network.captive-portal-service.enabled",		false);

/******************************************************************************
 * SECTION: HTTP                                                              *
 ******************************************************************************/

// PREF: Disallow NTLMv1
// https://bugzilla.mozilla.org/show_bug.cgi?id=828183
user_pref("network.negotiate-auth.allow-insecure-ntlm-v1",	false);
// it is still allowed through HTTPS. uncomment the following to disable it completely.
//user_pref("network.negotiate-auth.allow-insecure-ntlm-v1-https",		false);

// PREF: Enable CSP 1.1 script-nonce directive support
// https://bugzilla.mozilla.org/show_bug.cgi?id=855326
user_pref("security.csp.experimentalEnabled",			true);

// PREF: Enable Content Security Policy (CSP)
// https://developer.mozilla.org/en-US/docs/Web/Security/CSP/Introducing_Content_Security_Policy
// https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP
user_pref("security.csp.enable",				true);

// PREF: Enable Subresource Integrity
// https://developer.mozilla.org/en-US/docs/Web/Security/Subresource_Integrity
// https://wiki.mozilla.org/Security/Subresource_Integrity
user_pref("security.sri.enable",				true);

// PREF: DNT HTTP header (disabled)
// https://www.mozilla.org/en-US/firefox/dnt/
// https://en.wikipedia.org/wiki/Do_not_track_header
// https://dnt-dashboard.mozilla.org
// https://github.com/pyllyukko/user.js/issues/11
// NOTICE: Do No Track must be enabled manually
//user_pref("privacy.donottrackheader.enabled",		true);

// PREF: Send a referer header with the target URI as the source
// https://bugzilla.mozilla.org/show_bug.cgi?id=822869
// https://github.com/pyllyukko/user.js/issues/227
// NOTICE: Spoofing referers breaks functionality on websites relying on authentic referer headers
// NOTICE: Spoofing referers breaks visualisation of 3rd-party sites on the Lightbeam addon
// NOTICE: Spoofing referers disables CSRF protection on some login pages not implementing origin-header/cookie+token based CSRF protection
// TODO: https://github.com/pyllyukko/user.js/issues/94, commented-out XOriginPolicy/XOriginTrimmingPolicy = 2 prefs
user_pref("network.http.referer.spoofSource",			true);

// PREF: Don't send referer headers when following links across different domains (disabled)
// https://github.com/pyllyukko/user.js/issues/227
// user_pref("network.http.referer.XOriginPolicy",		2);


`

var EXTENSIONPREFS = `{}`
