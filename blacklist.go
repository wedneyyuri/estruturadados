package estruturadados

import (
	"sort"
	"strings"
)

// Node .
type Node struct {
	Value    string
	Children map[string]*Node
}

func (n *Node) isBlacklisted(termParts map[string]bool) bool {
	if n == nil {
		panic("como isso foi acontecer, o que fiz de errado?")
	}

	if len(n.Children) == 0 {
		return true
	}
	if _, ok := termParts[n.Value]; !ok {
		return false
	}
	for _, v := range n.Children {
		if v.isBlacklisted(termParts) {
			return true
		}
	}
	return false
}

func (n *Node) addWords(termParts []string) {
	if len(termParts) == 0 {
		return
	}

	value := termParts[0]

	if _, ok := n.Children[value]; !ok {
		n.Children[value] = &Node{
			Value:    value,
			Children: make(map[string]*Node),
		}
	}

	n.Children[value].addWords(termParts[1:])
}

// NewBlacklist returns a initialized blacklist.
func NewBlacklist() *Blacklist {
	return &Blacklist{
		tree: make(map[string]*Node),
	}
}

// Blacklist .
type Blacklist struct {
	tree map[string]*Node
}

func (b *Blacklist) addWords(termParts []string) {
	value := termParts[0]

	if _, ok := b.tree[value]; !ok {
		b.tree[value] = &Node{
			Value:    value,
			Children: make(map[string]*Node),
		}
	}

	b.tree[value].addWords(termParts[1:])
}

// AddTerm .
func (b *Blacklist) AddTerm(v string) {
	if b.IsBlacklisted(v) {
		return
	}

	term := Slugfy(v)
	termParts := strings.Split(term, "-")
	sort.Strings(termParts)

	b.addWords(termParts)

	// strings.Split()
	// add novo termo na blacklist
}

// IsBlacklisted .
func (b *Blacklist) IsBlacklisted(v string) bool {
	term := Slugfy(v)
	termParts := strings.Split(term, "-")
	m := make(map[string]bool)
	for _, v := range termParts {
		m[v] = true
	}
	for _, v := range termParts {
		if n, ok := b.tree[v]; ok {
			if n.isBlacklisted(m) {
				return true
			}
		}
	}
	return false
}

// ContentOnBlackList returns whether a term is blacklisted or not.
func ContentOnBlackList(searchBlacklist []string, content string) bool {
	contentWords := strings.Split(Slugfy(content), "-")
	_ = contentWords

	return false
}
