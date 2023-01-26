package output

import (
	"encoding/json"
	"fmt"

	"github.com/owenrumney/gtail/pkg/logger"
)

func (o *output) writeJson(value interface{}) error {
	content, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		logger.Error("Error marshalling json: %v", err)
		fmt.Printf("%v\n", value)
	}
	fmt.Println(string(content))
	return nil
}
