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