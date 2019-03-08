package main

import "fmt"

type RssStream struct {
	Name        string
	Description string
	Code        string
}

func (rs RssStream) getUrl(c Config) string {
	return c.Target_website + rs.Code
}

func (rs RssStream) print() {
	fmt.Println(rs)
}
