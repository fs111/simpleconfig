simpleconfig - a simple configuration file library for go
=========================================================

`simpleconfig` is a trivial configuration file library for configuration files
following the `key=value` format. `key = value` is also permitted.  The format
uses `#` as the comment character. 

Example config file
-------------------

This is an example config file used by
[kurz.go](http://github.com/fs111/kurz.go) from which the `simpleconfig` library
was extracted:

    # port to listen on
    port=9999

    # listen only on loopback
    listen=127.0.0.1

    # hostname to use
    hostname=localhost:9999

    # redis related settings
    redis.netaddress=tcp:localhost:6379
    redis.database=0
    redis.password=

    # directory to serve static files
    static-directory=static


Installation and usage
======================

To install the library with the go comand, simply do:

    go get github.com/fs111/simpleconfig

To stay in sync with the latest developments, use:

    go get -u github.com/fs111/simpleconfig

Code
----

    import (
        "github.com/fs111/simpleconfig"
    )


    func main(){
        cfg, err := simpleconfig.NewConfig("/path/to/config")
        if err != nil{
            // handle error
        }

        // get a string value from the config
        // handle error, if necessary...
        val, e := cfg.GetString("key")


        // get a string, but return a default, if it's not set
        // if the key "ook?" does not exist, it will return "ook!", otherwise it
        // will return the value from the config file
        val, e := cfg.GetStringDefault("ook?", "ook!")

        // same API for integers
        intVal, e := cfg.GetInt("somekey")

        // or

        intVal := cfg.GetIntDefault("answer", 42)


        // and of course bool

        flag, err := cfg.GetBool("someflag")

        // or

        flag := cfg.GetBoolDefault("someflag", false)

    }


CI
==
[![Build Status](https://travis-ci.org/fs111/simpleconfig.png)](https://travis-ci.org/fs111/simpleconfig)


Contributing
============
If you have problems, suggestions or anything else that you would like to add,
please open a pull request here on github.


License
=======
Copyright (c) 2014 André Kelpe

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
