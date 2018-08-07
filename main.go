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

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_- "

var tlds = []string{
	"com",
	"net",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		domainName := generateDomainName(scanner.Text())
		fmt.Println(domainName)
	}
}

func generateDomainName(plainText string) string {
	domainName := adjustToRFC3986(plainText)
	return domainName + "." + tlds[rand.Intn(len(tlds))]
}

func adjustToRFC3986(plainText string) string {
	lowerText := strings.ToLower(plainText)
	domainName := make([]rune, 0)
	for _, r := range lowerText {
		if !isAllowedChar(r) {
			continue
		}
		if unicode.IsSpace(r) {
			r = '-'
		}

		domainName = append(domainName, r)
	}

	return string(domainName)
}

func isAllowedChar(r rune) bool {
	return strings.ContainsRune(allowedChars, r)
}
