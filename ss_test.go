package randomString

import "testing"

func Test_String(t *testing.T) {
	s := String()

	ok := false
	for _, v := range ss {
		if v == s {
			ok = true
			break
		}
	}
	if ok == false {
		t.Fatalf("%s is not in %v", s, ss)
	}
}
