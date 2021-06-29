/*
Package goocord contains GooCord - a simple, robust and versatile Discord-API
library made for Go.
*/
package goocord

import "fmt"

const Version = "0.1.0"

type Client struct{}

func (c *Client) Hello() {
	fmt.Println("Hello, world!")
}
