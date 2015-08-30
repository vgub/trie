package trie

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Test_JSON_Unmarshal(t *testing.T) {
	var tr Trie

	reader, err := os.Open("json.txt")
	if err != nil {
		t.Fail()
	}
	defer reader.Close()

	tr.Unmarshal(reader)
	dict := tr.FindEntries("автор")

	for i := 0; i < len(dict); i++ {
		fmt.Println(dict[i])
	}
}
