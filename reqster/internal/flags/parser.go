package flags

import "strings"

func Parser(data []string,sep string) map[string]string {
	result := map[string]string{}
	for _, val := range data {
		element := strings.Split(val, sep)
		result[element[0]] = element[1]
	}
	return result
}
