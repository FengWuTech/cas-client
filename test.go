package cas_client

import (
	"encoding/json"
	"fmt"
)

func TestNew() *Cas {
	return New("http://127.0.0.1:8129", "test", "testestestestestestszetsztrsssssss")
}

func LogV(value interface{}) {
	switch value.(type) {
	case interface{}:
		str, _ := json.Marshal(value)
		fmt.Printf("%v", string(str))
	default:
		fmt.Printf("%v", value)
	}
	fmt.Printf("%v\n", value)
}
