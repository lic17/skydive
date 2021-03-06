/*
 * Copyright (C) 2017 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package graph

import (
	"sort"

	"github.com/skydive-project/skydive/graffiti/filters"
	"github.com/skydive-project/skydive/graffiti/getter"
)

// Sort criteria
const (
	SortByInt64 int = iota + 1
	SortByString
)

type slice interface {
	Get(i int) getter.Getter
	Swap(i, j int)
}

type sortableSlice struct {
	sortBy     string
	sortOrder  filters.SortOrder
	sortByType int
	length     int
	items      slice
}

type sortableNodeSlice struct {
	nodes []*Node
}

type sortableEdgeSlice struct {
	edges []*Edge
}

func (s sortableSlice) Len() int {
	return s.length
}

func (s sortableSlice) Swap(i, j int) {
	s.items.Swap(i, j)
}

func (s sortableSlice) lessInt64(i, j int) bool {
	i1, _ := s.items.Get(i).GetFieldInt64(s.sortBy)
	i2, _ := s.items.Get(j).GetFieldInt64(s.sortBy)

	if s.sortOrder == filters.SortOrder_Ascending {
		return i1 < i2
	}
	return i1 > i2
}

func (s sortableSlice) lessString(i, j int) bool {
	s1, _ := s.items.Get(i).GetFieldString(s.sortBy)
	s2, _ := s.items.Get(j).GetFieldString(s.sortBy)

	if s.sortOrder == filters.SortOrder_Ascending {
		return s1 < s2
	}

	return s1 > s2
}

func (s sortableSlice) Less(i, j int) bool {
	switch s.sortByType {
	case SortByInt64:
		return s.lessInt64(i, j)
	case SortByString:
		return s.lessString(i, j)
	}

	// detection of type
	if _, err := s.items.Get(i).GetFieldInt64(s.sortBy); err == nil {
		s.sortByType = SortByInt64
		return s.lessInt64(i, j)
	}

	s.sortByType = SortByInt64
	return s.lessString(i, j)
}

func (s sortableNodeSlice) Swap(i, j int) {
	s.nodes[i], s.nodes[j] = s.nodes[j], s.nodes[i]
}

func (s sortableNodeSlice) Get(i int) getter.Getter {
	return s.nodes[i]
}

func (s sortableEdgeSlice) Swap(i, j int) {
	s.edges[i], s.edges[j] = s.edges[j], s.edges[i]
}

func (s sortableEdgeSlice) Get(i int) getter.Getter {
	return s.edges[i]
}

func sortSlice(items slice, length int, sortBy string, sortOrder filters.SortOrder) {
	sort.Sort(sortableSlice{
		sortBy:    sortBy,
		sortOrder: sortOrder,
		length:    length,
		items:     items,
	})
}

// SortNodes sorts a set of nodes
func SortNodes(nodes []*Node, sortBy string, sortOrder filters.SortOrder) {
	sortSlice(sortableNodeSlice{nodes: nodes}, len(nodes), sortBy, sortOrder)
}

// SortEdges sorts a set of edges
func SortEdges(edges []*Edge, sortBy string, sortOrder filters.SortOrder) {
	sortSlice(sortableEdgeSlice{edges: edges}, len(edges), sortBy, sortOrder)
}
