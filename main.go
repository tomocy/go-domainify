package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

var tlds = []string{
	"com",
	"net",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lowerText := strings.ToLower(scanner.Text())
		domain := make([]rune, 0)
		for _, r := range lowerText {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}

			domain = append(domain, r)
		}

		fmt.Println(string(domain) + "." + tlds[rand.Intn(len(tlds))])
	}
}
