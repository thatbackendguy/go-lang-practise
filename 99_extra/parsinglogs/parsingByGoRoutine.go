package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("/home/yash/Desktop/original.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	dataString := string(data)

	dataSlices := strings.Split(dataString, "\n")

	jsonStrData := "[{}"

	fmt.Println("Starting scanning...")

	// Read the file line by line
	for _, v := range dataSlices {
		fmt.Println(v)
		if strings.Contains(v, "metric poll response: {") {
			jsonStrData += ",{"
		} else {
			jsonStrData += v
		}
	}

	jsonStrData += "]"

	fmt.Println("Scanning complete...")

	//fmt.Println(jsonStrData)

	var jsonArr []interface{}
	err = json.Unmarshal([]byte(jsonStrData), &jsonArr)
	if err != nil {
		fmt.Println("Error unmarshaling byte string:", err)
		return
	}

	objectTarget := os.Args[1]
	metricType := os.Args[2]

	for _, obj := range jsonArr {
		myMap := obj.(map[string]interface{})

		if len(myMap) != 0 {
			if myMap["object.target"].(string) == objectTarget && myMap["metric.name"].(string) == metricType {

				timestamp := myMap["event.timestamp"].(float64)
				unixSeconds := int64(timestamp)
				unixNanoseconds := int64((timestamp - float64(unixSeconds)) * 1e9)
				t := time.Unix(unixSeconds, unixNanoseconds)
				formattedTime := t.Format("02-January-2006 03:04:05.000 pm")

				resultMap, _ := json.Marshal(myMap["result"])

				fmt.Printf("%v\t%v\t%v\n", formattedTime, myMap["status"], string(resultMap))
			}
		}

	}
}
