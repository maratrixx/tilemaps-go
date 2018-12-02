package utils

func SliceKeyExists(key int, sl []string) bool {
	for k := range sl {
		if k == key {
			return true
		}
	}
	return false
}

func SliceValueExists(val string, sl []string) bool {
	for _, v := range sl {
		if v == val {
			return true
		}
	}
	return false
}
