package server

import (
	"os"
)

type Client struct {
}

func NewClient(disp *Display, fd os.File) *Client { return nil }
func (c *Client) Destroy()                        {}
func (c *Client) Flush()                          {}

func (c *Client) GetDisplay() *Display                { return nil }
func (c *Client) GetCredentials() (pid, uid, gid int) { return }
func (c *Client) GetObj(id uint) *Resource            { return nil }

func (c *Client) PostNoMemory() {}

func (c *Client) AddDestroyListener(l *Listener)             {}
func (c *Client) GetDestroyListener(notify func()) *Listener { return nil }
