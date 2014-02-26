package gowl

import (
	"bytes"
	"encoding/binary"
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

func TestString(t *testing.T) {

	checkEnc := func(s String, err error) {
		if err != nil {
			t.Error(err)
			return
		}
		var slen, n uint32
		if err = binary.Read(bytes.NewBuffer(s[:4]), binary.LittleEndian, &slen); err != nil {
			t.Error(err)
			return
		}

		var aligned bool = len(s)%4 == 0

		n = slen + 4
		if n > 75 {
			t.Logf("StringEncode(\"%s\"...)\n\nlen:%v\n"+
				"Padding:%v\n32bitAligned?:%v\n\n",
				s[4:75], slen, s[n:], aligned)
		} else {
			t.Logf("StringEncode(\"%s\")\n\nlen:%v\n"+
				"Padding:%v\n32bitAligned?:%v\n\n",
				s[4:n-1], slen, s[n:], aligned)
		}
	}
	gowlstr := make([]String, len(lorem))
	for i, s := range lorem {

		var err error
		gowlstr[i], err = StringEncode(s)
		checkEnc(gowlstr[i], err)

		gos, err := StringDecode(gowlstr[i])
		if err != nil {
			t.Error(err)
			continue
		}
		if gos == s{
		t.Log("StringDecode(gowlstr[i]) == lorem[i]: True\n")
		}else{
			t.Error("StringDecode(gowlstr[i]) == lorem[i]: False\n")
		} 
	}
}
