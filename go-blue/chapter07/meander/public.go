package meander

// Facade defines a Public method to view a struct view
type Facade interface {
	Public() interface{}
}

// Public returns a struct view
func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return o
}
