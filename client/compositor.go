package client

type Compositor struct {
	object
}

func (c *Compositor) Marshal() {}

type SubCompositor struct {
	object
}

func (c *SubCompositor) Marshal() {}
