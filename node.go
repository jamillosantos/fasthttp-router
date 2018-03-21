package fasthttp_router

import (
	"github.com/valyala/fasthttp"
	"bytes"
	"fmt"
)

type node struct {
	wildcard *node
	children map[string]*node
	handler  fasthttp.RequestHandler
	names    []string
}

func newNode() *node {
	return &node{
		children: make(map[string]*node),
	}
}

func (n *node) Add(path string, handler fasthttp.RequestHandler, names []string) {
	pathBytes := bytes.Split([]byte(path), []byte{'/'})
	lpath := len(pathBytes)
	if names == nil {
		names = make([]string, 0)
	}
	parent := n
	for i := 0; i < lpath; i++ {
		token := pathBytes[i]
		if len(token) > 0 {
			if token[0] == ':' {
				name := string(token[1:])
				node := parent.wildcard
				nodeCreated := false
				if node == nil {
					node = newNode()
					parent.wildcard = node
					nodeCreated = true
				}
				names = append(names, name)
				parent = node
				if i+1 >= lpath {
					if !nodeCreated && node.handler != nil {
						panic(fmt.Sprintf("conflict adding '%s'", path))
					}
					node.handler = handler
					node.names = names
				}
				continue
			} else {
				spath := string(pathBytes[i])
				node, ok := parent.children[spath]
				if !ok {
					node = newNode()
					parent.children[spath] = node
				}
				if i+1 < lpath {
					parent = node
				} else {
					if ok && node.handler != nil {
						panic(fmt.Sprintf("conflict adding '%s'", path))
					}
					node.handler = handler
					node.names = names
					return
				}
			}
		} else {
			panic("empty token")
		}
	}
}

func (n *node) Matches(path [][]byte, values [][]byte) (bool, *node, [][]byte) {
	lpath := len(path)
	for i := 0; i < lpath; i++ {
		token := string(path[i])
		node, ok := n.children[token]
		if ok {
			if i+1 < lpath {
				return node.Matches(path[i+1:], values)
			} else if node.handler == nil {
				return true, n.wildcard, values
			} else {
				return true, node, values
			}
		} else if n.wildcard != nil {
			values = append(values, path[i])
			if i+1 < lpath {
				return n.wildcard.Matches(path[i+1:], values)
			} else if n.wildcard.handler == nil {
				return false, nil, nil
			} else {
				return true, n.wildcard, values
			}
		} else {
			return false, nil, nil
		}
	}
	return false, nil, nil
}
