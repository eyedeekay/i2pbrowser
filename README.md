I2P Profile Configuring Launcher for Firefox, Multiplatform
===========================================================

Configures a Firefox profile in the working directory with extensions and 
configuration settings for I2P. At this time, it is configured for mixed I2P
and clearnet browsing via my browser plugin. In the near future, other plugins
may be configured as well as additional settings in user.js.

Experimental
------------

This is experimental software, but I do use it successfully under these
circumstances:

### Computer 1

        Ubuntu Linux
        Java I2P router installed from Project PPA
        Mozilla Firefox installed from official Ubuntu repository

### Computer 2

        Debian GNU/Linux
        Java I2P router built from source and installed as *my user*
        Mozilla Firefox installed from official Debian repository

Note that *I have not personally used this with i2pd yet, nor has it been*
*extensively tested on Windows or OSX* but, if I did everything right the first
time, and we can pretty much rely on it to run i2pd with the system-provided
settings, it should usually work, and it should work on Windows as long as you
have Firefox installed to the default path but if it doesn't work, create an
issue here or yell at me on reddit(r/i2p) or i2pforum.net and we'll figure it
out.

Even More Experimental
----------------------

There are two logical places to take this in the medium-close future. One is to
embed a router inside this, along with some kind of console, probably based on a
soon-to-be-usable SWIG binding for i2pd, or a possibly even sooner-to-be-usable
Go binding for Java I2P and an embedded runtime and class files? I dunno which
one's going to be easier but it turns out it's comparatively easy to talk to
Java from Go, much easier than I thought before about 2AM today. Speaking of
which, I need to remember these commands:

        /usr/lib/jvm/java-11-openjdk-amd64/bin/java \
            -DloggerFilenameOverride=logs/log-router-@.txt \
            -Di2p.dir.base=$HOME/i2p \
            -Xmx512m \
            -Djava.library.path=$HOME/i2p:$HOME/i2p/lib \
            -classpath $HOME/i2p/lib/BOB.jar:$HOME/i2p/lib/addressbook.jar:$HOME/i2p/lib/commons-el.jar:$HOME/i2p/lib/commons-logging.jar:$HOME/i2p/lib/desktopgui.jar:$HOME/i2p/lib/i2p.jar:$HOME/i2p/lib/i2psnark.jar:$HOME/i2p/lib/i2ptunnel.jar:$HOME/i2p/lib/jasper-compiler.jar:$HOME/i2p/lib/jasper-runtime.jar:$HOME/i2p/lib/javax.servlet.jar:$HOME/i2p/lib/jbigi.jar:$HOME/i2p/lib/jetty-continuation.jar:$HOME/i2p/lib/jetty-deploy.jar:$HOME/i2p/lib/jetty-http.jar:$HOME/i2p/lib/jetty-i2p.jar:$HOME/i2p/lib/jetty-io.jar:$HOME/i2p/lib/jetty-java5-threadpool.jar:$HOME/i2p/lib/jetty-rewrite-handler.jar:$HOME/i2p/lib/jetty-security.jar:$HOME/i2p/lib/jetty-servlet.jar:$HOME/i2p/lib/jetty-servlets.jar:$HOME/i2p/lib/jetty-sslengine.jar:$HOME/i2p/lib/jetty-start.jar:$HOME/i2p/lib/jetty-util.jar:$HOME/i2p/lib/jetty-webapp.jar:$HOME/i2p/lib/jetty-xml.jar:$HOME/i2p/lib/jrobin.jar:$HOME/i2p/lib/jstl.jar:$HOME/i2p/lib/mstreaming.jar:$HOME/i2p/lib/org.mortbay.jetty.jar:$HOME/i2p/lib/org.mortbay.jmx.jar:$HOME/i2p/lib/router.jar:$HOME/i2p/lib/routerconsole.jar:$HOME/i2p/lib/sam.jar:$HOME/i2p/lib/standard.jar:$HOME/i2p/lib/streaming.jar:$HOME/i2p/lib/systray.jar:$HOME/i2p/lib/wrapper.jar \
            -Dwrapper.key=key \
            -Dwrapper.port=port \
            -Dwrapper.jvm.port.min=31000 \
            -Dwrapper.jvm.port.max=31999 \
            -Dwrapper.disable_console_input=TRUE \
            -Dwrapper.pid=pid \
            -Dwrapper.version=3.5.39 \
            -Dwrapper.native_library=wrapper \
            -Dwrapper.arch=x86 \
            -Dwrapper.service=TRUE \
            -Dwrapper.cpu.timeout=10 \
            -Dwrapper.jvmid=3 org.tanukisoftware.wrapper.WrapperSimpleApp \
            net.i2p.router.Router
        $HOME/i2p/i2psvc $HOME/i2p/wrapper.config \
            wrapper.syslog.ident=i2p \
            wrapper.java.command=java \
            wrapper.pidfile=$HOME/.i2p/i2p.pid \
            wrapper.name=i2p \
            wrapper.displayname=I2P Service \
            wrapper.daemonize=TRUE \
            wrapper.statusfile=$HOME/.i2p/i2p.status \
            wrapper.java.statusfile=$HOME/.i2p/i2p.java.status \
            wrapper.logfile=$HOME/.i2p/wrapper.log


This *almost* launches a router on a pure-Go JVM, could be a configuration
detail or two away from working.

        export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64/
        JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64/
        $HOME/go/bin/java \
            -XuseJavaHome \
            -classpath $HOME/i2p:$HOME/i2p/lib:$HOME/i2p/lib/BOB.jar:$HOME/i2p/lib/addressbook.jar:$HOME/i2p/lib/commons-el.jar:$HOME/i2p/lib/commons-logging.jar:$HOME/i2p/lib/desktopgui.jar:$HOME/i2p/lib/i2p.jar:$HOME/i2p/lib/i2psnark.jar:$HOME/i2p/lib/i2ptunnel.jar:$HOME/i2p/lib/jasper-compiler.jar:$HOME/i2p/lib/jasper-runtime.jar:$HOME/i2p/lib/javax.servlet.jar:$HOME/i2p/lib/jbigi.jar:$HOME/i2p/lib/jetty-continuation.jar:$HOME/i2p/lib/jetty-deploy.jar:$HOME/i2p/lib/jetty-http.jar:$HOME/i2p/lib/jetty-i2p.jar:$HOME/i2p/lib/jetty-io.jar:$HOME/i2p/lib/jetty-java5-threadpool.jar:$HOME/i2p/lib/jetty-rewrite-handler.jar:$HOME/i2p/lib/jetty-security.jar:$HOME/i2p/lib/jetty-servlet.jar:$HOME/i2p/lib/jetty-servlets.jar:$HOME/i2p/lib/jetty-sslengine.jar:$HOME/i2p/lib/jetty-start.jar:$HOME/i2p/lib/jetty-util.jar:$HOME/i2p/lib/jetty-webapp.jar:$HOME/i2p/lib/jetty-xml.jar:$HOME/i2p/lib/jrobin.jar:$HOME/i2p/lib/jstl.jar:$HOME/i2p/lib/mstreaming.jar:$HOME/i2p/lib/org.mortbay.jetty.jar:$HOME/i2p/lib/org.mortbay.jmx.jar:$HOME/i2p/lib/router.jar:$HOME/i2p/lib/routerconsole.jar:$HOME/i2p/lib/sam.jar:$HOME/i2p/lib/standard.jar:$HOME/i2p/lib/streaming.jar:$HOME/i2p/lib/systray.jar:$HOME/i2p/lib/wrapper.jar \
            net.i2p.router.Router

The other is to embed a Firefox Portable, possibly one that is enhanced in terms
of privacy by inheriting non-Branding patches from Tor Browser. How this is to
be achieved remains to be seen, but extracting Tor from their own build system
is not sustainable or advisable, this may be better.

Reproducibility?
----------------

What definitions of reproducibility this satisfies remains to be seen. I will
think about it later, but I *think* that it should be reproducible and since
the resources it embeds should also be reproducible, then it should be possible
to build this reproducibly.

