package __27

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	s := Constructor()
	s.AddText("bxyackuncqzcqo")
	s.CursorLeft(12)
	s.DeleteText(3)
	s.CursorLeft(5)
	s.AddText("osdhyvqxf")
	fmt.Println(s.CursorRight(10))
}
