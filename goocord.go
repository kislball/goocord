package goocord

import "fmt"

const VERSION = "0.1.0"

type Client struct{}

func (c *Client) Hello() {
	fmt.Println("Hello, world!")
}
