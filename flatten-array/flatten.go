package flatten

// Helper is a recursive function that flattens element array and append to accumulator (accu)
func Helper(element interface{}) interface{} {
	// Type assertion
	nestedArray, ok := element.([]interface{})
	acc := []interface{}{}
	if ok == true {
		for _, e := range nestedArray {
			// Type switches
			switch v := e.(type) {
			case int:
				acc = append(acc, v)
			case []interface{}:
				flattenArray := Helper(v).([]interface{})
				for _, e1 := range flattenArray {
					acc = append(acc, e1)
				}
			}
		}
	}
	return acc
}

// Flatten accepts a nested number array and return a one-dimensional array back
func Flatten(input interface{}) interface{} {
	return Helper(input)
}
