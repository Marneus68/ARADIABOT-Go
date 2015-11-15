package users

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

type UsersFileLocation int

const (
    Default UsersFileLocation = iota
    Local UsersFileLocation = iota
)

var FileLocation UsersFileLocation = 3

var DefaultFileFullPath string = "./users.log"
var LocalFileFullPath string = "~/.aradiabot/users.log"

type Users struct {
   Map map[string]int64
}

func checkUsersFileLocation() {
    _, err := os.Stat(LocalFileFullPath)
    if err == nil {
        fmt.Printf("Local users file (%s) found.\n", LocalFileFullPath)
        FileLocation = Local
    } else {
        _, eerr := os.Stat(DefaultFileFullPath)
        if eerr == nil {
            fmt.Printf("Default users file (%s) found.\n", DefaultFileFullPath)
            FileLocation = Default
        } else {
            fmt.Printf("No user file found, creating default users file (%s).\n", DefaultFileFullPath)
            os.Create(DefaultFileFullPath)
        }
    }
}

func New() *Users {
    fmt.Println("Loading users file...")
    checkUsersFileLocation()
    return Read()
}

func Read() *Users {
    var user *Users = &Users{ Map : make(map[string]int64) }
    user.Read()
    return user
}

func (u Users) Read() {
    var fileName string
    switch FileLocation {
        case Default:
        fileName = DefaultFileFullPath
        case Local:
        fileName = LocalFileFullPath
    }
    file, err := os.Open(fileName)
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
    var fileName string
    switch FileLocation {
        case Default:
        fileName = DefaultFileFullPath
        case Local:
        fileName = LocalFileFullPath
    }
    file, err := os.Create(fileName)
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

