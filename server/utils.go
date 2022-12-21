package main

import "encoding/json"

func isJson(data []byte) bool {
	return json.Valid(data)
}
