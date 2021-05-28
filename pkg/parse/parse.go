package parse

import (
	"errors"
	"fmt"
	"strings"
)

/**
1. resolve user defined variables pointed at @@()
2. output
*/
func Execute(in []byte, data map[string]string) ([]byte, error) {
	s := string(in)
	for k, v := range data {
		val := fmt.Sprintf("@@(%s)", k)
		s = strings.Replace(s, val, v, -1)
	}
	if strings.Index(s, "@@(") > 0 {
		return nil, errors.New("unresolved variables found")
	}
	return []byte(s), nil
}
