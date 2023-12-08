package fuzz

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverseByBytes(f *testing.F) {
	strs := []string{"hello", "特朗不", " ", "!123"}

	//模糊测试种子，用于自动生成测试用例
	for _, str := range strs {
		f.Add(str)
	}

	f.Fuzz(func(t *testing.T, orig string) {

		//判断初始字符串
		if !utf8.ValidString(orig) {
			return
		}
		rev, err := ReverseByBytes(orig)
		if err != nil {
			t.Fatalf("Reverse failed: %v", err)
		}
		//判断一次反转字符串
		if !utf8.ValidString(rev) {
			return
		}
		doubleRev, err := ReverseByBytes(rev)

		if err != nil {
			t.Fatalf("Reverse failed: %v", err)
		}

		//两次反转过后应该和源字符串相等
		if orig != doubleRev {
			t.Fatalf("Reverse(%q) = %q, want: %q", orig, rev, orig)
		}

	})
}

//func FuzzReverseByRunes(f *testing.F) {
//}
