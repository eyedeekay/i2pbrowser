package i2pbrowser

import (
	"log"
)

func GenerateDefaultBookmarks(dir, configfile string) string {
	_, err := brb.OutputServerConfigFile()
	if err != nil {
		log.Fatalf("Config file generation error, %s", err)
	}
	for len(brb.OutputAutoLink()) < len("http://localhost:7669/connect?host=?name=invisibleirc") {
	}
	return `<!DOCTYPE NETSCAPE-Bookmark-file-1>
  <!-- This is an automatically generated file.
       It will be read and overwritten.
       DO NOT EDIT! -->
  <META HTTP-EQUIV="Content-Type" CONTENT="text/html; charset=UTF-8">
  <TITLE>Bookmarks</TITLE>
  <H1>Bookmarks Menu</H1>

  <DL><p>
      <DT><A HREF="place:type=6&sort=14&maxResults=10" ADD_DATE="1539649423" LAST_MODIFIED="1539649424">Recent Tags</A>
      <HR>    <DT><A HREF="http://legwork.i2p/yacysearch.html?query=%s&verify=cacheonly&contentdom=text&nav=location%2Chosts%2Cauthors%2Cnamespace%2Ctopics%2Cfiletype%2Cprotocol%2Clanguage&startRecord=0&indexof=off&meanCount=5&resource=global&prefermaskfilter=&maximumRecords=10&timezoneOffset=0" ADD_DATE="1539652098" LAST_MODIFIED="1539652098" SHORTCUTURL="legwork.i2p">Search YaCy &#39;legwork&#39;: Search Page</A>
      <DD>Software HTTP Freeware Home Page
      <DT><H3 ADD_DATE="1539649419" LAST_MODIFIED="1539649423" PERSONAL_TOOLBAR_FOLDER="true">Bookmarks Toolbar</H3>
      <DL><p>
          <DT><A HREF="place:sort=8&maxResults=10" ADD_DATE="1539649423" LAST_MODIFIED="1539649423">Most Visited</A>
          <DT><A HREF="http://i2p-projekt.i2p/" ADD_DATE="1538511080" LAST_MODIFIED="1538511080">I2P Anonymous Network</A>
          <DD>Anonymous peer-to-peer distributed communication layer built with open source tools and designed to run any traditional Internet service such as email, IRC or web hosting.
          <DT><A HREF="http://127.0.0.1:7669/connect" ADD_DATE="1538511080" LAST_MODIFIED="1538511080">I2P IRC Chat</A>
          <DD>Connect to the I2P IRC chatrooms to chat with other i2p people
          <DT><A HREF="` + brb.OutputAutoLink() + `ADD_DATE="1538511080" LAST_MODIFIED="1538511080">Personal IRC Chat</A>
          <DD>Host a group chat with other i2p people, from your own device
      </DL><p>
  </DL>
`
	/**



	 */
}
