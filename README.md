# Golang

Learning Go

OS: Ubuntu 12.04, 32bit VM

## Installation
https://wiki.ubuntu.com/Go  
go-tip package

### Run Hello World

```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```   

go run hello.go  
```
fork/exec /tmp/go-build173576991/command-line-arguments/_obj/exe/hello: permission denied
```

issue: /tmp is mounted in nonexec mode  
solution: remount /tmp 
```
sudo mount -o remount exec /tmp

```
## Book & Reference  

Learning Go (free ebook): https://github.com/miekg/gobook  
Go Web 编程 （free ebook): https://github.com/astaxie/build-web-application-with-golang