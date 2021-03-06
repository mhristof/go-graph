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
					"a": {"b"},
					"b": {"c"},
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
					"a": {"b", "c"},
					"b": {"d"},
					"c": {"d"},
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
					"a": {"c"},
					"b": {"c", "d"},
					"c": {"e"},
					"d": {"f"},
					"e": {"h", "f"},
					"f": {"g"},
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
					"a": {"c"},
					"b": {"c", "d"},
					"c": {"e"},
					"d": {"f"},
					"e": {"h", "f"},
					"f": {"g"},
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

func TestSortAll(t *testing.T) {
	var cases = []struct {
		name string
		in   DependencyMap
		out  []string
	}{
		{
			name: "include unvisited nodes",
			in: DependencyMap{
				Nodes: strings.Fields("a b c d e f g h"),
				edge: map[string][]string{
					"a": {"c"},
					"b": {"c", "d"},
					"c": {"e"},
					"d": {"f"},
					"e": {"h", "f"},
					"f": {"g"},
				},
			},
			out: []string{"h", "g", "f", "e", "c", "a", "d", "b"},
		},
	}

	for _, test := range cases {
		assert.Equal(t, test.out, test.in.SortAll(test.in.Nodes[0]), test.name)
	}
}

func TestSortExported(t *testing.T) {
	var cases = []struct {
		name string
		in   DependencyMap
		out  []string
	}{
		{
			name: "include unvisited nodes",
			in: DependencyMap{
				Nodes: strings.Fields("a b c d e f g h"),
				edge: map[string][]string{
					"a": {"c"},
					"b": {"c", "d"},
					"c": {"e"},
					"d": {"f"},
					"e": {"h", "f"},
					"f": {"g"},
				},
			},
			out: []string{"h", "g", "f", "e", "c", "a"},
		},
	}

	for _, test := range cases {
		assert.Equal(t, test.out, test.in.Sort(test.in.Nodes[0]), test.name)
	}
}

func TestEdge(t *testing.T) {
	var cases = []struct {
		name   string
		in     DependencyMap
		source string
		dest   string
		out    DependencyMap
	}{
		{
			name: "new edge",
			in: DependencyMap{
				Nodes: []string{"a", "b"},
			},
			source: "a",
			dest:   "b",
			out: DependencyMap{
				Nodes: []string{"a", "b"},
				edge: map[string][]string{
					"a": {"b"},
				},
			},
		},
		{
			name: "duplicate edge",
			in: DependencyMap{
				Nodes: []string{"a", "b"},
				edge: map[string][]string{
					"a": {"b"},
				},
			},
			source: "a",
			dest:   "b",
			out: DependencyMap{
				Nodes: []string{"a", "b"},
				edge: map[string][]string{
					"a": {"b"},
				},
			},
		},
	}

	for _, test := range cases {
		test.in.Edge(test.source, test.dest)
		assert.Equal(t, test.out, test.in, test.name)
	}
}
