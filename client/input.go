package client

type Seat struct{}

func (s *Seat) Marshal() {}

type Pointer struct{}

func (p *Pointer) Marshal() {}

type Keyboard struct{}

func (k *Keyboard) Marshal() {}

type Touch struct{}

func (t *Touch) Marshal() {}
