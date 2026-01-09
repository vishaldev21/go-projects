package utils

import "slices"

func MethodValidator(method string) bool {
	methodAllowed := []string{"get", "post", "put", "patch", "delete"}
	return slices.Contains(methodAllowed, method)
}
