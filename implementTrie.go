package demo

func (t *Trie) Insert(word string) {
	operatingNode := t
	wordLen := len(word)
	for i := 0; i < wordLen; i++ {
		if n, ok := operatingNode.chartsForSearching[word[i]]; ok {
			operatingNode = n
			continue
		}
		operatingNode.chartsForSearching[word[i]] = NewTrie()
		operatingNode = operatingNode.chartsForSearching[word[i]]
	}
	operatingNode.oneWordHere = true
}

func (t *Trie) reachTrieNode(prefix string) *Trie {
	operatingNode := t
	wordLen := len(prefix)
	for i := 0; i < wordLen; i++ {
		if n, ok := operatingNode.chartsForSearching[prefix[i]]; ok {
			operatingNode = n
			continue
		} else {
			return nil
		}
	}
	return operatingNode
}

func (t *Trie) Search(word string) bool {
	if operatingNode := t.reachTrieNode(word); operatingNode != nil {
		return operatingNode.oneWordHere
	}
	return false
}

func (t *Trie) StartWith(prefix string) bool {
	if operatingNode := t.reachTrieNode(prefix); operatingNode != nil {
		return true
	}
	return false
}

func getAllWords(node *Trie, words *[]string, thisLevelWord []byte) {
	if len(node.chartsForSearching) < 1 {
		return
	}
	makeSlice := func(n int) []byte {
		defer func() {
			if recover() != nil {
				panic("ErrTooLarge")
			}
		}()
		return make([]byte, n)
	}
	for k, v := range node.chartsForSearching {
		if v.oneWordHere {
			*words = append(*words, string(append(thisLevelWord, k)))
		}
		nextLevelWord := makeSlice(len(thisLevelWord))
		copy(nextLevelWord, thisLevelWord)
		nextLevelWord = append(nextLevelWord, k)
		getAllWords(v, words, nextLevelWord)
	}

}

func (t *Trie) GetWordsStartWith(prefix string) (words []string) {
	if node := t.reachTrieNode(prefix); node != nil {
		prefixByte := []byte(prefix)
		getAllWords(node, &words, prefixByte)
		return
	}
	return
}

func (t *Trie) GetAllWords() []string {
	var levelWord []byte
	var words []string
	getAllWords(t, &words, levelWord)
	return words
}
