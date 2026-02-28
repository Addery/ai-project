package main

import "testing"

//func TestAdd(t *testing.T) {
//	if ans := Add(1, 2); ans != 3 {
//		// 如果不符合预期，那就是测试不通过
//		t.Errorf("1 + 2 expected be 3, but %d got", ans)
//	}
//
//	if ans := Add(-10, -20); ans != -30 {
//		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
//	}
//}

// 子测试
//func TestAdd(t1 *testing.T) {
//	t1.Run("add1", func(t *testing.T) {
//		if ans := Add(1, 2); ans != 3 {
//			// 如果不符合预期，那就是测试不通过
//			t.Fatalf("1 + 2 expected be 3, but %d got", ans)
//		}
//	})
//	t1.Run("add2", func(t *testing.T) {
//		if ans := Add(-10, -20); ans != -30 {
//			t.Fatalf("-10 + -20 expected be -30, but %d got", ans)
//		}
//	})
//
//}

func TestAdd(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"a1", 2, 3, 5},
		{"a2", 2, -3, -1},
		{"a3", 2, 0, 2},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Add(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}
