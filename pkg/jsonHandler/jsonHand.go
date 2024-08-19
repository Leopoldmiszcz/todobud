// * JŚON handler
package jsonHand

//TODO: Try to fix the json formating
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func FileCheck() {
	_, err := os.ReadFile("todobudData.json")

	if err != nil {
		fmt.Println("There's been a error while reading a file:", err)
	}
}

func WriteJSONFile(value string) {

	filename := "todobudData.json"

	data := map[string]interface{}{
		"todo": map[string]interface{}{
			"done": false,
			"task": value,
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("sad boi")
		return
	}

	fmt.Println(string(jsonData))
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.SetFlags(log.Ldate | log.Lshortfile)
		log.Println(err)
	}

	f.WriteString(string(jsonData))
	f.Close()
}
