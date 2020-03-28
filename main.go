package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-libvterm"
	"log"
	"os"
	"regexp"
)

var isDebug = false

const BELL = "\x07"
const ESC = "\x1b"
const LF = "\x0a"
const CR = "\x0d"

// ESC + k + .... ESC + \  :  Title Definition String
const TITLE_DEFINITION = ESC + "k" + `.*?\\`

// OSC begins with ESC + ] and ends with BELL
const OSC_PREFIX = ESC + `]`
const OSC_ALL = OSC_PREFIX + ` .*? ` + BELL

// CSI begins with ESC + [
const CSI_PREFIX = ESC + `\[`

// ESC[ + Pm + m : Settings for Color, BackGround, Font and so on
const CSI_SGR_X = CSI_PREFIX + `[0-9;]+?m`

// move line sequences
const CSI_MOVE_LINE = CSI_PREFIX + `[0-9;]+?[ABEFGHST]`

// DEC/Xterm
const CSI_DEC_XTERM = CSI_PREFIX + `\?[0-9]+?[hl]`
const CSI_DEC_XTERM_XT_EXTSCRN = CSI_PREFIX + `\?1049[hl]`

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " typescript")
		os.Exit(0)
	}

	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("cannot open file :", os.Args[1])
		os.Exit(0)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)

	vt := vterm.New(25, 80)
	defer vt.Close()

	vt.SetUTF8(true)
	screen := vt.ObtainScreen()
	rows, cols := vt.Size()

	for scanner.Scan() {
		screen.Reset(true)
		line := scanner.Text()

		debug("Before Convert                    : ")
		debugln(line)

		_, err := vt.Write([]byte(delNonControllableChar(line)))
		if err != nil {
			log.Fatal(err)
		}
		screen.Flush()
		out := ""
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				cell, err := screen.GetCellAt(row, col)
				if err != nil {
					log.Fatal(err)
				}
				chars := cell.Chars()

				out = out + string(chars)
			}
		}

		out = delSpaces(out)
		out = delNullAndReplacementChar(out)
		fmt.Println(out)
		debugln()
	}
}

func delNonControllableChar(line string) string {
	debugln("Before delete Control chars       : " + line)

	// delete OSC sequence (OSC squence has BELL so delete it before deleting bell)
	rep := regexp.MustCompile(OSC_ALL)
	line = rep.ReplaceAllString(line, "")

	// delete bell
	rep = regexp.MustCompile(BELL)
	line = rep.ReplaceAllString(line, "")

	// delete title definition
	rep = regexp.MustCompile(TITLE_DEFINITION)
	line = rep.ReplaceAllString(line, "")

	// delete color
	rep = regexp.MustCompile(CSI_SGR_X)
	line = rep.ReplaceAllString(line, "")

	// replace from clear screen to <CR><LF>
	rep = regexp.MustCompile(CSI_DEC_XTERM_XT_EXTSCRN)
	line = rep.ReplaceAllString(line, "")

	// delete DEC/Xterm
	rep = regexp.MustCompile(CSI_DEC_XTERM)
	line = rep.ReplaceAllString(line, "")

	// delete move line sequences
	rep = regexp.MustCompile(CSI_MOVE_LINE)
	line = rep.ReplaceAllString(line, "")

	debugln("After delete Control chars        : " + line)
	return line
}

func delNullAndReplacementChar(line string) string {
	debugln("Before delete Null and DecodeError: " + line)
	var rep *regexp.Regexp

	// delete \x00
	rep = regexp.MustCompile("\x00")
	line = rep.ReplaceAllString(line, "")

	// delete \xef \xbf \xbd <- REPLACEMENT CHARACTER
	rep = regexp.MustCompile("\xef\xbf\xbd")
	line = rep.ReplaceAllString(line, "")

	debugln("After delete Null and DecodeError : " + line)

	return line
}

// delete spaces + ~ or @  for deletin vi output
func delSpaces(line string) string {
	debugln("Before delete spaces              : " + line)
	var rep *regexp.Regexp

	// delete 5 spaces + ~
	rep = regexp.MustCompile("     *~")
	line = rep.ReplaceAllString(line, "")

	// delete 5 spaces @ ~
	rep = regexp.MustCompile("     *@")
	line = rep.ReplaceAllString(line, "")

	debugln("After delete spaces               : " + line)

	return line
}

func debugln(v ...interface{}) {
	if isDebug {
		fmt.Println(v...)
	}
}

func debug(v ...interface{}) {
	if isDebug {
		fmt.Print(v...)
	}
}
