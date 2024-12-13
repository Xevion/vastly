package api

func Pointer[T any](d T) *T {
	return &d
}
