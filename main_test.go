package main

import (
	//"fmt"
	"testing"
)

var str string
var match string

func TestDelNonControllableChar_OSC(t *testing.T) {
	str = "abc\x1b] 4;10 \x07def"
	out := delNonControllableChar(str)
	if out != "abcdef" {
		t.Errorf("returned string is \"%s\"", out)
	}
}

func TestDelNonControllableChar_BELL(t *testing.T) {
	str = "abc\x07def"
	out := delNonControllableChar(str)
	if out != "abcdef" {
		t.Errorf("returned string is \"%s\"", out)
	}
}

func TestDelNonControllableChar_TITLE(t *testing.T) {
	str = `kroot@localhost:~/\[root@localhost:1:~/]$ echo`
	out := delNonControllableChar(str)
	if out != "[root@localhost:1:~/]$ echo" {
		t.Errorf("returned string is \"%s\"", out)
	}
}

func TestDelNonControllableChar_COLOR(t *testing.T) {
	str = `[01;34mNetworkManager[0m`
	out := delNonControllableChar(str)
	if out != "NetworkManager" {
		t.Errorf("returned string is \"%s\"", out)
	}
}

func TestDelNonControllableChar_XT_EXTSCRN(t *testing.T) {
	str = `abcd[?1049hefgh[?1049lijkl`
	//str = `[?1049habc`
	out := delNonControllableChar(str)
	if out != "abcdefghijkl" {
		t.Errorf("returned string is \"%s\"", out)
	}
}

func TestDelNullAndReplacementCharSpaces_NULL(t *testing.T) {
	str = "ab\x00\x00\x00cd"
	out := delNullAndReplacementChar(str)
	if out != "abcd" {
		t.Errorf("returned string is \"%s\"", out)
	}
}

func TestDelNullAndReplacementCharSpaces_DecodeError(t *testing.T) {
	str = "ab\xef\xbf\xbdcd"
	out := delNullAndReplacementChar(str)
	if out != "abcd" {
		t.Errorf("returned string is \"%s\"", out)
	}
}

func TestDelSpaces_Tilde(t *testing.T) {
	str = "          ~abc"
	out := delSpaces(str)
	if out != "abc" {
		t.Errorf("returned string is \"%s\"", out)
	}
}

func TestDelSpaces_At(t *testing.T) {
	str = "          @def"
	out := delSpaces(str)
	if out != "def" {
		t.Errorf("returned string is \"%s\"", out)
	}
}
