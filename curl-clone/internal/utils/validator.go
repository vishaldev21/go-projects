package utils

func MethodValidator(method string) bool {
	methodAllowed := []string{"get", "post", "put", "patch", "delete"}
	for _, val := range methodAllowed {
		if method == val {
			return true
		}
	}
	return false
}
