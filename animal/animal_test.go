package animal

import "testing"

func TestSpeaker(t *testing.T) {
	tests := []struct {
		s    Speaker
		want string
	}{
		{Dog{Name: "D"}, "D: 汪汪!"},
		{Cat{Name: "C"}, "C: 喵喵~"},
		{Cow{Name: "W"}, "W: 哞哞"},
	}

	for _, tc := range tests {
		got := tc.s.Speak()
		if got != tc.want {
			t.Errorf("Speak() = %q, want %q", got, tc.want)
		}
	}
}

func TestTypeAssertion(t *testing.T) {
	var s Speaker = Dog{Name: "X"}
	if d, ok := s.(Dog); ok {
		t.Logf("类型断言成功: %T %+v", d, d)
	} else {
		t.Fatal("类型断言失败")
	}
}
