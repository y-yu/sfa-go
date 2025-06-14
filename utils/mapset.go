package utils

import (
	"github.com/y-yu/sfa-go/common"
	"iter"
	"maps"
	"slices"
)

type MapSet[T comparable] map[T]struct{}

func NewMapSet[T comparable](ts ...T) MapSet[T] {
	s := make(map[T]struct{}, len(ts))
	for _, t := range ts {
		s[t] = struct{}{}
	}
	return s
}

func (s MapSet[T]) Contains(t T) bool {
	_, exists := s[t]
	return exists
}

func (s MapSet[T]) Add(t T) {
	s[t] = struct{}{}
}

func (s MapSet[T]) Union(rhs MapSet[T]) MapSet[T] {
	result := NewMapSet(slices.Collect(maps.Keys(s))...)
	for k, _ := range rhs {
		result[k] = struct{}{}
	}
	return result
}

func (s MapSet[T]) IsSubset(rhs MapSet[T]) bool {
	for k, _ := range s {
		if !rhs.Contains(k) {
			return false
		}
	}
	return true
}

func (s MapSet[T]) Cardinality() int {
	return len(s)
}

func (s MapSet[T]) IsSuperset(rhs MapSet[T]) bool {
	for k, _ := range rhs {
		if !s.Contains(k) {
			return false
		}
	}
	return true
}

func (s MapSet[T]) Iter() iter.Seq[T] {
	return maps.Keys(s)
}

func (s MapSet[T]) Intersect(rhs MapSet[T]) MapSet[T] {
	result := NewMapSet[T]()
	for k, _ := range s {
		if rhs.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

func (s MapSet[T]) Equal(rhs MapSet[T]) bool {
	return s.IsSubset(rhs) && rhs.IsSubset(s)
}

func (s MapSet[T]) Remove(v T) {
	delete(s, v)
}

type Set = MapSet[common.State]

func NewSet(ss ...common.State) Set {
	return NewMapSet(ss...)
}
