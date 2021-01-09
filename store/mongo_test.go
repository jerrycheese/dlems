package store

import (
	"fmt"
	"testing"
)

func TestMongo(t *testing.T) {

	res := mapToD(map[string]interface{}{
		"Name": "tan",
	})

	fmt.Printf("%+v", res)
}
