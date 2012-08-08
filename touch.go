package gowl

import (
	"bytes"
)

var _ bytes.Buffer

type Touch struct {
//	*WlObject
	id int32
	listeners map[int16][]chan interface{}
	events []func (t *Touch, msg message)
}

//// Requests
//// Events
func (t *Touch) HandleEvent(msg message) {
	if t.events[msg.opcode] != nil {
		t.events[msg.opcode](t, msg)
	}
}

type TouchDown struct {
	Serial uint32
	Time uint32
	Surface *Surface
	Id int32
	X int32
	Y int32
}

func (t *Touch) AddDownListener(channel chan interface{}) {
	t.listeners[0] = append(t.listeners[0], channel)
	addListener(channel)
}

func touch_down(t *Touch, msg message) {
	var data TouchDown

	serial,err := readUint32(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	data.Serial = serial

	time,err := readUint32(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	data.Time = time

	surfaceid, err := readInt32(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	surface := new(Surface)
	surfaceobj := getObject(surfaceid)
	if surfaceobj == nil {
		return
	}
	surface = surfaceobj.(*Surface)
	data.Surface = surface

	id,err := readInt32(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	data.Id = id

	x,err := readFixed(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	data.X = x

	y,err := readFixed(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	data.Y = y

	for _,channel := range t.listeners[0] {
		go func() {
			channel <- data
		} ()
	}
	printEvent("touch", t, "down", serial, time, surface.Id(), id, x, y)
}

type TouchUp struct {
	Serial uint32
	Time uint32
	Id int32
}

func (t *Touch) AddUpListener(channel chan interface{}) {
	t.listeners[1] = append(t.listeners[1], channel)
	addListener(channel)
}

func touch_up(t *Touch, msg message) {
	var data TouchUp

	serial,err := readUint32(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	data.Serial = serial

	time,err := readUint32(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	data.Time = time

	id,err := readInt32(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	data.Id = id

	for _,channel := range t.listeners[1] {
		go func() {
			channel <- data
		} ()
	}
	printEvent("touch", t, "up", serial, time, id)
}

type TouchMotion struct {
	Time uint32
	Id int32
	X int32
	Y int32
}

func (t *Touch) AddMotionListener(channel chan interface{}) {
	t.listeners[2] = append(t.listeners[2], channel)
	addListener(channel)
}

func touch_motion(t *Touch, msg message) {
	var data TouchMotion

	time,err := readUint32(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	data.Time = time

	id,err := readInt32(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	data.Id = id

	x,err := readFixed(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	data.X = x

	y,err := readFixed(msg.buf)
	if err != nil {
		// XXX Error handling
	}
	data.Y = y

	for _,channel := range t.listeners[2] {
		go func() {
			channel <- data
		} ()
	}
	printEvent("touch", t, "motion", time, id, x, y)
}

type TouchFrame struct {
}

func (t *Touch) AddFrameListener(channel chan interface{}) {
	t.listeners[3] = append(t.listeners[3], channel)
	addListener(channel)
}

func touch_frame(t *Touch, msg message) {
	var data TouchFrame

	for _,channel := range t.listeners[3] {
		go func() {
			channel <- data
		} ()
	}
	printEvent("touch", t, "frame")
}

type TouchCancel struct {
}

func (t *Touch) AddCancelListener(channel chan interface{}) {
	t.listeners[4] = append(t.listeners[4], channel)
	addListener(channel)
}

func touch_cancel(t *Touch, msg message) {
	var data TouchCancel

	for _,channel := range t.listeners[4] {
		go func() {
			channel <- data
		} ()
	}
	printEvent("touch", t, "cancel")
}

func NewTouch() (t *Touch) {
	t = new(Touch)
	t.listeners = make(map[int16][]chan interface{}, 0)

	t.events = append(t.events, touch_down)
	t.events = append(t.events, touch_up)
	t.events = append(t.events, touch_motion)
	t.events = append(t.events, touch_frame)
	t.events = append(t.events, touch_cancel)
	return
}

func (t *Touch) SetId(id int32) {
	t.id = id
}

func (t *Touch) Id() int32 {
	return t.id
}