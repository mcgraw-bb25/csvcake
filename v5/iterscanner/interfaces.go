package iterscanner

// Bakeable interfaces implement Bake(), which is a
// close cousin of the more general Clone() concept.
// However Bake here is used specifically for CSV rows
// which have been prepared, and therefore are typed
// to match a destination struct.  The interface returned
// should be a newly created struct that conforms to the
// underlying type of the Bakeable.  Bake is used in
// calls to Next() in order to return an interface to
// the caller that is safe to cast into a struct
// of the underlying type.
type Bakeable interface {
	Bake(map[string]interface{}) interface{}
}

type preparable interface {
	prepare(string) (interface{}, error)
}
