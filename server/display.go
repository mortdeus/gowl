package server

import (
	"os"
	"strings"
)

type Display struct{}

func (d *Display) Create()  {}
func (d *Display) Destroy() {}

func (d *Display) GetEventLoop() {}

func (d *Display) AddSocket(s string) {

}
func (d *Display) Terminate() {

}
func (d *Display) Run() {

}
func (d *Display) FlushClients() {

}
func (d *Display) GetSerial() uint {

}
func (d *Display) NextSerial() uint {

}
func (d *Display) AddDestroyListner(l *Listner) {

}
func (d *Display) GetDestroyListner(n NotifyFunc) {

}
func (d *Display) AddShmFormat(fmt uint)           {}
func (d *Display) GetAdditionalShmFormats() []uint {}
