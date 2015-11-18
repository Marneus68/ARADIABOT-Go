// Package deciding on the location to use to read and store users file, history
// and configuration
package filoc

import (
    "os"
    "fmt"
)

var Dir string = ""

var dirs = [...]string{ "~/.aradiabot", "." }

func init() {
    for _, dirpath := range dirs {
        //fmt.Println(line)
        _, err := os.Stat(dirpath)
        if err == nil {
            fmt.Printf("Using %s as the working directory.\n", dirpath)
            Dir = dirpath
            break
        }
    }
}

func Foo() {

}

