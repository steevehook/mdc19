package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func listen(c net.Conn, mem cache) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		l := strings.ToLower(strings.TrimSpace(scanner.Text()))
		values := strings.Split(l, " ")
		switch {
		case len(values) == 3 && values[0] == "set":
			mem.Set(values[1], values[2])
			fmt.Fprintf(c, "OK\n-> ")
		case len(values) == 2 && values[0] == "get":
			fmt.Fprintf(c, "%s\n-> ",
				mem.Get(values[1]))
		case len(values) == 1 && values[0] == "exit":
			c.Close()
		default:
			fmt.Fprintf(c, "UNKNOWN: %s\n-> ", l)
		}
	}
}
