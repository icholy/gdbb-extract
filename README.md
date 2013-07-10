# GDB Breakpoint Extract

This tools searches your source files for specially formatted comments and converts them to gdb breakpoints.

**Example:**

``` go
package main

import "fmt"

func main() {
  //break
  fmt.Println("Hello World")
}
```

if you invoke `gdbb-extract` on this file, it will output the following

``` sh
$ gdbb-extract *.go

break /path/to/main.go:6
```

Here's how you would debug this example.

``` sh
// build with debug flags
$ go build -gcflags "-N -l" -o out`

// extract the breakpoints
$ gdbb-extract *.go > .breakpoints

// run gdb
$ gdb -x .breakpoints -ex run
```

**Conditional** breakpoints work too

```
//break if $len(x) > 5
```

**Commands** can be placed after a `:` and are separated by `;`

```
//break : print x; continue
```

Combine them

```
//break if y < 2 : set y = 2; continue
```
