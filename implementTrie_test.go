package demo

import (
	"reflect"
	"testing"
)

func TestTrie_Search(t *testing.T) {
	root := NewTrie()
	word := "ab"
	if e := root.Search(word); e {
		t.Errorf(" %s word should not exist \n", word)
	}
	if e := root.StartWith(word[:1]); e {
		t.Errorf(" %s prfix should not exist \n", word[:1])
	}
	root.Insert(word)
	if e := root.Search(word); !e {
		t.Errorf(" %s word should exist \n", word)
	}
	if e := root.StartWith(word[:1]); !e {
		t.Errorf(" %s prfix should exist \n", word[:1])
	}
	if e := root.StartWith(word); !e {
		t.Errorf(" %s prefix should exist \n", word)
	}

	word = "abc"
	if e := root.Search(word); e {
		t.Errorf(" %s word should not exist \n", word)
	}
	root.Insert(word)
	if e := root.Search(word); !e {
		t.Errorf(" %s word should exist \n", word)
	}
	if ws := root.GetAllWords(); !reflect.DeepEqual(ws, []string{"ab", "abc"}) {
		t.Errorf(" %s words not match %s \n", ws, []string{"ab", "abc"})
	}
	if ws := root.GetWordsStartWith(word[:2]); !reflect.DeepEqual(ws, []string{word}) {
		t.Errorf(" %s words not match %s \n", ws, []string{word})
	}
}
