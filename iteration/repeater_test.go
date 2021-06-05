package iteration

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	rep := Repeat("a", 10)
	fmt.Println(rep)
	//Output: aaaaaaaaaa
}
