package gowl

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

var lorem = []string{
	"Lorem",
	"Lorem ipsum Qui Excepteur in irure.",

	"Lorem ipsum Amet sint sunt laboris qui " +
		"est nostrud consectetur velit deserunt elit tempor id minim " +
		"tempor occaecat minim qui sunt ex laboris ad enim minim est " +
		"deserunt in nulla nulla ex sit sit ea dolore eu ut et magna " +
		"aute nisi incididunt aliquip aliqua velit fugiat eu irure " +
		"ullamco laboris anim commodo velit in mollit magna qui Ut " +
		"fugiat ea mollit sit veniam occaecat aliqua quis minim in sint " +
		"sit nisi dolore adipisicing voluptate sit dolor sit sed irure " +
		"quis incididunt culpa in dolore ad amet id sed ullamco occaecat " +
		"irure consectetur mollit sunt do cillum enim laborum laborum " +
		"eiusmod ex occaecat commodo tempor ex esse elit esse consectetur " +
		"sit sunt sed aute Duis commodo nulla officia consequat dolor sint " +
		"consectetur veniam voluptate aute consequat exercitation dolore " +
		"dolore fugiat deserunt laboris do occaecat quis."}

func TestStringGob(t *testing.T) {
	for _, s := range lorem {
		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		if err := enc.Encode((*String)(&s)); err != nil {
			t.Fatal(err)
		}
		var s1 String
		dec := gob.NewDecoder(buf)
		if err := dec.Decode(&s1); err != nil {
			t.Fatal(err)
		}
		var sfmt string
		if len(s) > 40 {
			sfmt = fmt.Sprintf("\"%v...\" == \"%v...\": %v", s[:40], string(s1[:40]), s == string(s1))
		} else {
			sfmt = fmt.Sprintf("\"%v\" == \"%v\": %v", s, string(s1), s == string(s1))
		}
		if s != string(s1) {
			t.Fatal(sfmt)
		}
		t.Log(sfmt)

	}
}
