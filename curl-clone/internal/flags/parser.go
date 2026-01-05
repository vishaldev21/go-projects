package flags

import "strings"

func HeaderParser(data []string) map[string]string {
	result := map[string]string{}
	for _, val := range data {
		element := strings.Split(val, ":")
		result[element[0]] = element[1]
	}
	return result
}
func ParamParser(data []string) map[string]string {
	result := map[string]string{}
	for _, val := range data {
		element := strings.Split(val, "=")
		result[element[0]] = element[1]
	}
	return result
}
