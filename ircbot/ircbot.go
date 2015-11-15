package ircbot

import (
    "os"
    "fmt"
    "net"
    "bufio"
    "strings"
    "strconv"
    "aradiabot/users"
)

func Run(ipport, channel, name string) int {
    var u = users.New()

    fmt.Println("Loaded %s registered users.", strconv.Itoa(len(u.Map)))

    i := 0
    c, err := net.Dial("tcp", ipport)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating the TCP socket: %s\n", err)
        return 1
    }

    cbuf := bufio.NewReader(c)
    for {
        line, err := cbuf.ReadString('\n')
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error reading incomming lines in the socket: %s\n", err)
            return 1
        }

        if i == 2 {
            go identify(c, channel, name)
        }

        split := strings.Split(line, " ")

        if len(split) > 1 {
            if strings.HasPrefix(line, "PING") {
                go ping(c, strings.Join(split[1:], " "))
                continue
            }
        }

        if len(split) > 2 {

        }

        fmt.Fprintf(os.Stdout, "%s", line)
        i++
    }
    return 0
}

func identify(conn net.Conn, channel string, name string) {
    fmt.Fprintf(conn, "NICK %s\r\n", name)
    fmt.Fprintf(conn, "USER %s somewhere me :Aradia Megido\r\n", name)
    fmt.Fprintf(conn, "JOIN %s\r\n", fmt.Sprintf("#%s", channel))
}

func ping(conn net.Conn, answer string) {
    fmt.Fprintf(conn, "PING %s", answer)
}

