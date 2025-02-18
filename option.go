package huh

import "fmt"

// Option is an option for select fields.
type Option[T comparable] struct {
	Key      string
	Value    T
	selected bool
	children []Option[T] // Add this field
}

// NewOptions returns new options from a list of values.
func NewOptions[T comparable](values ...T) []Option[T] {
	options := make([]Option[T], len(values))
	for i, o := range values {
		options[i] = Option[T]{
			Key:   fmt.Sprint(o),
			Value: o,
		}
	}
	return options
}

// NewOption returns a new select option.
func NewOption[T comparable](key string, value T) Option[T] {
	return Option[T]{Key: key, Value: value}
}

// Selected sets whether the option is currently selected.
func (o Option[T]) Selected(selected bool) Option[T] {
	o.selected = selected
	return o
}

// Children adds nested options and returns the parent for chaining.
func (o Option[T]) Children(child ...Option[T]) Option[T] {
	o.children = append(o.children, child...)
	return o
}

// Helper to recursively apply a function to all children.
func (o *Option[T]) traverseChildren(fn func(*Option[T])) {
	for i := range o.children {
		fn(&o.children[i])
		o.children[i].traverseChildren(fn)
	}
}

// String returns the key of the option.
func (o Option[T]) String() string {
	return o.Key
}
