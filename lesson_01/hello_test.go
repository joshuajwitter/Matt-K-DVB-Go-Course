package lesson_01

import "testing"

func TestSayHello(t *testing.T) {

	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Josh"},
			result: "Hello, Josh!",
		},
		{
			items:  []string{"Josh", "Gina", "Charlotte"},
			result: "Hello, Josh, Gina, Charlotte!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}
}
