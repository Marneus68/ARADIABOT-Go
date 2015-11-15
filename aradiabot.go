package main

import (
    "os"
    "fmt"
    "aradiabot/usage"
    "aradiabot/ircbot"
)

func main() {
    if len(os.Args) < 4 {
        fmt.Fprintln(os.Stderr, "Not enough paramters provided.")
        usage.Usage()
        os.Exit(0)
    }

    os.Exit(ircbot.Run(os.Args[1], os.Args[2], os.Args[3]))
}

