package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Open the file
	file, err := os.Open("/home/yash/Desktop/test.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new Scanner for reading the file line by line
	scanner := bufio.NewScanner(file)

	jsonStrData := "[{}"

	fmt.Println("Starting scanning...")
	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "metric poll response: {") {
			jsonStrData += ",{"
		} else {
			jsonStrData += line
		}
	}

	jsonStrData += "]"

	fmt.Println("Scanning complete...")

	// Check if there was an error during the scan
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

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
