package main

import "fmt"

type TLDs []string

func (TLDs) String() string {
	return fmt.Sprint("")
}

func (ds *TLDs) Set(v string) error {
	*ds = append(*ds, v)
	return nil
}
