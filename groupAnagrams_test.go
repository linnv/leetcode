package demo

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"normal", args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}}, [][]string{
			[]string{"eat", "tea", "ate"},
			[]string{"tan", "nat"},
			[]string{"bat"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}

			if got := groupAnagramsTimeout(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagramsTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"normal", args{[]byte("cba")}, []byte("abc")},
		{"normal", args{[]byte("cbja")}, []byte("abcj")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.bs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGroupAnagramsTimeout(b *testing.B) {
	for i := 0; i < b.N; i++ {
		groupAnagramsTimeout([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	}
}
