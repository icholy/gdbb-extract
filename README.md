# gdbb-extract [![Build Status](https://travis-ci.org/icholy/gdbb-extract.png?branch=master)](https://travis-ci.org/icholy/gdbb-extract)

> Inline GDB Breakpoints 

This tool searches your go source files for specially formatted comments and converts them to gdb breakpoints.

**Example:**

``` go
package main

import "fmt"

func main() {
  for i := 0; i < 10; i++ {
    fmt.Println("Hello World") //break if i == 5
  }
}
```

**Usage:**

``` sh
# debug application
$ gdbb

# debug tests
$ gdbbtest
```

**Demo:**

![](http://i.imgur.com/GEEmHSs.gif)
