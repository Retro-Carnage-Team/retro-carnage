package util

type ChangeListener struct {
	Callback      func(value interface{}, property string)
	PropertyNames []string
}

func (cl *ChangeListener) handlesProperty(property string) bool {
	if 0 == len(cl.PropertyNames) {
		return true
	}
	for _, propertyName := range cl.PropertyNames {
		if propertyName == property {
			return true
		}
	}
	return false
}

func (cl *ChangeListener) Call(value interface{}, property string) {
	if cl.handlesProperty(property) {
		cl.Callback(value, property)
	}
}
