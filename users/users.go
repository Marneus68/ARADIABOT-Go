// Structure and functions relating to the list of registered users
//
// This package holds the Users structure representing the list of registered 
// users to the history service of the bot well as the functions necessary to
// read it from and write it to a file.
package users

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "aradiabot/filoc"
)

var File = "users.log"

type Users struct {
   Map map[string]int64
}

func New() *Users {
    fmt.Println("Loading users file...")
    return Read()
}

func Read() *Users {
    var user *Users = &Users{ Map : make(map[string]int64) }
    user.Read()
    return user
}

func (u Users) Read() {
    var filepath = filoc.Dir + "/" +  File
    _, err := os.Stat(filepath)
    if err != nil {
        fmt.Println("Users file doesn't exist. Creating it at " + 
            filepath)
        // If file doesn't exist we create it
        os.Create(filepath)
    }
    // Then we read it
    file, err := os.Open(filepath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to open users file for reading: %s\n", err)
        return
    }
    defer file.Close()

    scn := bufio.NewScanner(file)
    for scn.Scan() {
        s := strings.Split(scn.Text(), " ")
        var key = s[0]
        var val, err = strconv.ParseInt(s[1], 10, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Invalid line in users file: %s\n", err)
        } else {
            u.Map[key] = val
        }
    }
}

func (u Users) Write() {
    var filepath = filoc.Dir + "/" +  File
    file, err := os.Create(filepath)
    if (err != nil) {
        fmt.Fprintf(os.Stderr, "Unable to open users file for writing: %s\n", err)
        return
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for key, val := range u.Map {
        var wl string = key + " " + strconv.FormatInt(val, 10)
        fmt.Fprintln(w, wl)
    }
    w.Flush()
}

func (u Users) Add(key string, val int64) {
    u.Map[key] = val
}

func (u Users) Remove(key string) {
    delete(u.Map, key)
}

