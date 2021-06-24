package goocord

import "fmt"

type Client struct{}

func (c *Client) Hello() {
	fmt.Println("Hello, world!")
}
