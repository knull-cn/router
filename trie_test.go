package router

import (
	"testing"
)

type CheckNode struct {
	url    string
	checkv string
}

func TestTrie(t *testing.T) {
	var tt TrieTree
	urls := []string{"/a/b/c/d", "/a/b", "/"}
	for _, url := range urls {
		tt.AddPath(url, url)
	}
	//test;
	urls = []string{"/", "/a", "/a/b", "/a/b/c", "/a/b/c/d", "/b", "/a/b/c/d/e"}
	testnodes := []CheckNode{
		CheckNode{"/", "/"},
		CheckNode{"/a", "/"},
		CheckNode{"/a/b", "/a/b"},
		CheckNode{"/a/b/c", "/a/b"},
		CheckNode{"/a/b/c/d", "/a/b/c/d"},
		CheckNode{"/a/b/c/d/e", "/a/b/c/d"},
		CheckNode{"/a/b/c/e", "/a/b"},
		CheckNode{"/b", "/"},
	}
	for _, node := range testnodes {
		v := tt.GetValue(node.url)
		if v.(string) != node.checkv {
			t.Errorf("getv(%v) != checkv(%s)", v, node.checkv)
		}
	}
}
