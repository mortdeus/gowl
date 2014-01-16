package server

import (
	"os"
)

type Client struct {
}

func NewClient(disp *Display, fd os.File) *Client {}
func (c *Client) Destroy()                        {}
func (c *Client) Flush()                          {}

func (c *Client) GetDisplay() *Display                {}
func (c *Client) GetCredentials() (pid, uid, gid int) {} 
func (c *Client) GetObj(id uint) *Resource            {}

func (c *Client) PostNoMemory() {}

func (c *Client) AddDestroyListener(l *Listener)             {}
func (c *Client) GetDestroyListener(notify func()) *Listener {}


