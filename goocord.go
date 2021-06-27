/*
Package goocord contains GooCord is a simple, robust and versatile Discord-API
library made for Go.
*/
package goocord

import "fmt"

const VERSION = "0.1.0"

type Client struct{}

func (c *Client) Hello() {
	fmt.Println("Hello, world!")
}
