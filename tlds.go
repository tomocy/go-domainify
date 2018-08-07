package main

import "fmt"

type topLevelDomains []string

func (tlds topLevelDomains) String() string {
	return fmt.Sprint([]string(tlds))
}

func (tlds *topLevelDomains) Set(v string) error {
	*tlds = append(*tlds, v)
	return nil
}
