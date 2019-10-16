package struct_to_json

import (
	"encoding/json"
	"fmt"
)

type Sample struct {
	Name string `json:"name"`
}

func main() {
	sample := Sample{}
	sample.Name = "sample"

	j, err := json.Marshal(sample)
	if err != nil {
		fmt.Print(err)
		return
	}

	jsonStr := string(j)
	fmt.Printf("%s\n", jsonStr)
}
