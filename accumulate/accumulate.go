package accumulate

// Accumulate takes a collection, applies function to each element and return output collection
func Accumulate(inputs []string, fn func(string) string) []string {
	outputs := make([]string, len(inputs), len(inputs))
	for i, input := range inputs {
		outputs[i] = fn(input)
	}
	return outputs
}
