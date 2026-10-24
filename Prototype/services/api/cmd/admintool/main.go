// Command admintool creates admin secrets (12.4).
//
//	admintool password   reads a password from stdin and prints its bcrypt hash
//	admintool key        prints a random 32-byte key, base64url without padding
package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: admintool password|key")
		os.Exit(2)
	}
	switch os.Args[1] {
	case "password":
		fmt.Fprintln(os.Stderr, "Password (read from stdin, one line):")
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil && line == "" {
			fmt.Fprintln(os.Stderr, "no password read")
			os.Exit(1)
		}
		pw := strings.TrimRight(line, "\r\n")
		if len(pw) < 12 {
			fmt.Fprintln(os.Stderr, "use at least 12 characters")
			os.Exit(1)
		}
		h, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(string(h))
	case "key":
		b := make([]byte, 32)
		if _, err := rand.Read(b); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(base64.RawURLEncoding.EncodeToString(b))
	default:
		fmt.Fprintln(os.Stderr, "usage: admintool password|key")
		os.Exit(2)
	}
}
