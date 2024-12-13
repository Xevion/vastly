package api

func Pointer[T any](d T) *T {
	return &d
}

type Comparable[T comparable] struct {
	Eq *T `json:"eq,omitempty"`
	Lt *T `json:"lt,omitempty"`
	Le *T `json:"le,omitempty"`
	Gt *T `json:"gt,omitempty"`
	Ge *T `json:"ge,omitempty"`
}

// Then you can use these type aliases for convenience
type ComparableInteger = Comparable[int]
type ComparableFloat = Comparable[float64]

func Ge[T comparable](d T) *Comparable[T] {
	return &Comparable[T]{Ge: Pointer(d)}
}
