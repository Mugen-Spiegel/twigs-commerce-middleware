package services

import "strings"

func ErrorMessage(e_message string) map[string]map[string]string {
	slice_e_message := strings.Split(e_message, "\n")
	json_e_message := make(map[string]map[string]string)
	nested_map_message := make(map[string]string)
	for _, v := range slice_e_message {
		map_value := strings.Split(v, "Error:Field")
		key := strings.Split(map_value[0], "'")
		key = strings.Split(key[1], ".")

		nested_map_message[key[1]] = map_value[1]
	}
	json_e_message["error"] = nested_map_message
	return json_e_message

}
