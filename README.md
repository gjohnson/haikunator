
# haikunator

Generates server names in a similar fasion to Heroku. _Inspired by [usmanbashir/haikunator](https://github.com/usmanbashir/haikunator)._


### Usage

As a package.

```go
package main

import (
	"fmt"
	"github.com/gjohnson/haikunator"
)

func main() {
	name, _ := haikunator.Haikunate()
	fmt.Printf(name)
}
```

As a CLI tool.

```sh
$ ~ haikunator
quiet-moon-1513
```

View the [docs](http://godoc.org/github.com/gjohnson/haikunator).

### License

MIT