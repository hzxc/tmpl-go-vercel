package project

import (
	"fmt"
	"testing"
)

func Test_Str(t *testing.T) {
	name := "annie"
	s := fmt.Sprintf("hello %s", name)
	fmt.Println(s)
}
