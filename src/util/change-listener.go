package util

type ChangeListener struct {
	callback      func(value interface{}, property string)
	propertyNames []string
}

func (cl *ChangeListener) handlesProperty(property string) bool {
	if 0 == len(cl.propertyNames) {
		return true
	}

	for _, propertyName := range cl.propertyNames {
		if propertyName == property {
			return true
		}
	}

	return false
}

func (cl *ChangeListener) Call(value interface{}, property string) {
	if cl.handlesProperty(property) {
		cl.callback(value, property)
	}
}
