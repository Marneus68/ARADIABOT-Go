// Package responsible for writing in the history file
//
// This package holds the functions needed to write to the history file as well as
// getting a list of lines, called excerpt, located between two lones for the purpose
// of displaying a selected portion o the historu.
package hist

import (
    "os"
    "fmt"
    "bufio"
    "aradiabot/filoc"
)

var File = "history.log"

func Write(line string) {
    var filepath = filoc.Dir + "/" + File
    f, err := os.OpenFile(filepath, os.O_APPEND | os.O_WRONLY, 0600)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to open history file (%s) for witing: %s", filepath, err)
        return
    }
    defer f.Close()
    f.WriteString(line)
}

func Excerpt(start, end int64) []string {
    if start > end {
        start, end = end, start
    } else if start == end {
        return []string {}
    }
    var size = end - start
    var ret = make([]string, size)

    var filepath = filoc.Dir + "/" + File
    file, err := os.Open(filepath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to open history file for reading: %s\n", err)
        return []string {}
    }
    defer file.Close()

    scn := bufio.NewScanner(file)
    var i int64 = 0
    for scn.Scan() {
        var line = scn.Text()
        if i == end {
            break
        }
        if i >= start {
            ret = append(ret, line)
        }
        i++
    }
    return ret
}

