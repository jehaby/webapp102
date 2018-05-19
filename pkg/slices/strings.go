package slices

func StringInSlice(s []string, t string) bool {
	for _, v := range s {
		if t == v {
			return true
		}
	}
	return false
}
