package arrays

import "testing"

func TestDefineArray(t *testing.T) {
	test := struct {
		name    string
		wantOut [10]int
	}{
		name:    "define array test",
		wantOut: [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	t.Run(test.name, func(t *testing.T) {
		out := defineArray()
		if out != test.wantOut {
			t.Errorf("want %v, got %v", test.wantOut, out)
		}
	})
}

func TestCompareArrays(t *testing.T) {
	test := struct {
		name     string
		wantOut1 bool
		wantOut2 bool
		wantOut3 bool
	}{
		name:     "compare arrays",
		wantOut1: true,
		wantOut2: true,
		wantOut3: false,
	}

	t.Run(test.name, func(t *testing.T) {
		out1, out2, out3 := compArrays()
		if out1 != test.wantOut1 || out2 != test.wantOut2 || out3 != test.wantOut3 {
			t.Errorf("\nwant1 %v, got1 %v\nwant2 %v, got2 %v\nwant3 %v, got3 %v",
				test.wantOut1, out1, test.wantOut2, out2, test.wantOut3, out3)
		}
	})
}

func TestInitArrayWithKeys(t *testing.T) {
	test := struct {
		name     string
		wantOut1 bool
		wantOut2 bool
		wantOut3 [10]int
	}{
		name:     "compare arrays",
		wantOut1: false,
		wantOut2: false,
		wantOut3: [10]int{1, 0, 0, 0, 5, 0, 0, 0, 0, 10},
	}

	t.Run(test.name, func(t *testing.T) {
		out1, out2, out3 := initArrayWithKeys()
		if out1 != test.wantOut1 || out2 != test.wantOut2 || out3 != test.wantOut3 {
			t.Errorf("\nwant1 %v, got1 %v\nwant2 %v, got2 %v\nwant3 %v, got3 %v",
				test.wantOut1, out1, test.wantOut2, out2, test.wantOut3, out3)
		}
	})
}
