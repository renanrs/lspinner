# Lspinner
An easy and light weight spinner to use in your go console applications.

# Installing
Using Lspinner is easy. First, use `go get` to install the latest version
of the library. This command will install the `lspinner` along with the library and its dependencies:

    go get -u github.com/renanrs/lspinner

Next, include Lspinner in your application:

```go
import "github.com/renanrs/lspinner"
```

# Getting Started

In your console application, initialize `Lspinner`, use `Wait` function to show the spinner and wait message and user `Stop` function to erase the message.

```go
package main

import (
	"time"

	"github.com/renanrs/lspinner"
)

func main() {
	spinner := lspinner.New(lspinner.Options{})
	spinner.Wait()
  // code to wait the execution

	time.Sleep(time.Second * 5)
	spinner.Stop()
}
```
