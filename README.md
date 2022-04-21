A simple configurable logging facade for Go, using the [config](https://github.com/alt-golang/config) package
=====================================


![Language Badge](https://img.shields.io/github/languages/top/alt-golang/logger) <br/>
[release notes](https://github.com/alt-golang/logger/blob/main/HISTORY.md)

<a name="intro">Introduction</a>
--------------------------------
A simple configurable logging facade for Go, using the [config](https://github.com/alt-golang/config)
package.

<a name="usage">Usage</a>
-------------------------

To use the module, import the module and call the `GetLogger` function with a logging category (your module
url and package path is a sensible choice).

```go
import  ( "github.com/alt-golang/logger") ;

Logger := logger.GetLogger("github.com/MyModule/mypackage");
Logger.Info('Hello world!');

```

The logger's DefaultFactory will create a ConsoleLogger (uses `stdout and stderr`) object instance, with the root logging level
set to `info` by default.  To change the logging level for your module (category), add something similar to the
following in your [config](https://github.com/alt-golang/config) files.

`local-development.json`
```json

  "logging" : {
     "format" : "json",
     "level" : {
       "/" : "info",
       "github.com/MyModule/mypackage" : "debug"
     }
  }
}
```

<a name="license">License</a>
-----------------------------

May be freely distributed under the [MIT license](https://raw.githubusercontent.com/alt-javascript/config/main/LICENSE).

Copyright (c) 2022 Craig Parravicini    
