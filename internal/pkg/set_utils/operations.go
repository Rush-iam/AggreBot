package set_utils

func Difference[K comparable, V1, V2 any](base map[K]V1, subtract map[K]V2) []K {
	var difference []K
	for key := range base {
		if _, ok := subtract[key]; !ok {
			difference = append(difference, key)
		}
	}
	return difference
}

func DifferenceMap[K comparable, V1, V2 any](base map[K]V1, subtract map[K]V2) map[K]V1 {
	difference := make(map[K]V1)
	for key, value := range base {
		if _, ok := subtract[key]; !ok {
			difference[key] = value
		}
	}
	return difference
}
