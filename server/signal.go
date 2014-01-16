package server

type Signal struct{}

func (s *Signal) Init()              {}
func (s *Signal) Add(l *Listener)    {}
func (s *Signal) Get() *Listener     {}
func (s *Signal) Emit(d interface{}) {}
