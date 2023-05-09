# LTT - structured logger with routing

## Getting Started

### Installation
for install package use command:
```bash
go get -u github.com/nooize/smug
```

### Quick start code

POG provide `log` package with predefined static `Mux` for global logging.
so you can use it in same way as standard `log` package:
```go
package main

import (
 "github.com/nooize/ltr/log"
)

func main() {
 log.Info("hello world")
}

// Output: {"time":1516134303,"level":"info","message":"hello world"}
```
> Note: By default log writes to `os.Stderr`
> 
> Note: The default log level is *debug*


## Motivation



### Send log to one of 3 servers via http

```go
package main

import (
 "github.com/nooize/smug/log"
)

func main() {
	log.Append(targets.)
 log.Ctx().Info("hello world")
}

// Output: {"time":1516134303,"level":"info","message":"hello world"}
```
