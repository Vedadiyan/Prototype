package protoutil

import (
	"encoding/json"
	"testing"
)

func TestUnmarshall(t *testing.T) {
	const input = `
	{
		"Test":  {
			"value": "ok"
		}
	}
	`
	mapper := make(map[string]any)
	_ = json.Unmarshal([]byte(input), &mapper)
	var testMessage TestMessage
	err := Unmarshal(mapper, &testMessage)
	_ = err
}
