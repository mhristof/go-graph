package graph

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNode(t *testing.T) {
	var cases = []struct {
		name string
		node string
		m    DependencyMap
		exp  DependencyMap
	}{
		{
			name: "add the first node",
			node: "new",
			m:    DependencyMap{},
			exp: DependencyMap{
				Nodes: []string{"new"},
			},
		},
		{
			name: "duplicate item",
			node: "new",
			m: DependencyMap{
				Nodes: []string{"new"},
			},
			exp: DependencyMap{
				Nodes: []string{"new"},
			},
		},
		{
			name: "add the second node",
			node: "second",
			m: DependencyMap{
				Nodes: []string{"first"},
			},
			exp: DependencyMap{
				Nodes: []string{"first", "second"},
			},
		},
	}

	for _, test := range cases {
		test.m.AddNode(test.node)

		assert.Equal(t, test.exp, test.m, test.name)
	}
}

func TestSort(t *testing.T) {
	var cases = []struct {
		name string
		node string
		all  bool
		in   DependencyMap
		out  []string
	}{
		{
			name: "simple map",
			node: "a",
			in: DependencyMap{
				Nodes: []string{"a", "b", "c"},
				edge: map[string][]string{
					"a": []string{"b"},
					"b": []string{"c"},
				},
			},
			out: []string{
				"c", "b", "a",
			},
			all: true,
		},
		{
			name: "simple map",
			node: "a",
			in: DependencyMap{
				Nodes: []string{"a", "b", "c", "d"},
				edge: map[string][]string{
					"a": []string{"b", "c"},
					"b": []string{"d"},
					"c": []string{"d"},
				},
			},
			all: true,
			out: []string{"d", "b", "c", "a"},
		},
		{
			name: "complicated dep list, https://www.youtube.com/watch?v=ddTC4Zovtbc",
			node: "a",
			in: DependencyMap{
				Nodes: strings.Fields("a b c d e f g h"),
				edge: map[string][]string{
					"a": []string{"c"},
					"b": []string{"c", "d"},
					"c": []string{"e"},
					"d": []string{"f"},
					"e": []string{"h", "f"},
					"f": []string{"g"},
				},
			},
			all: true,
			out: []string{"h", "g", "f", "e", "c", "a", "d", "b"},
		},
		{
			name: "all: false, dont show nodes that are not connected with the initial one",
			node: "a",
			in: DependencyMap{
				Nodes: strings.Fields("a b c d e f g h"),
				edge: map[string][]string{
					"a": []string{"c"},
					"b": []string{"c", "d"},
					"c": []string{"e"},
					"d": []string{"f"},
					"e": []string{"h", "f"},
					"f": []string{"g"},
				},
			},
			all: false,
			out: []string{"h", "g", "f", "e", "c", "a"},
		},
	}

	for _, test := range cases {
		sorted := test.in.sort(test.node, test.all)

		assert.Equal(t, test.out, sorted, test.name)
	}
}

func TestDiff(t *testing.T) {
	var cases = []struct {
		name string
		this []string
		that []string
		out  []string
	}{
		{
			name: "same list",
			this: []string{"1", "2"},
			that: []string{"1", "2"},
			out:  nil,
		},
		{
			name: "that list has more items",
			this: []string{"1", "2"},
			that: []string{"1", "2", "3"},
			out:  []string{"3"},
		},
	}

	for _, test := range cases {
		assert.Equal(t, test.out, diff(test.this, test.that), test.name)
	}
}
