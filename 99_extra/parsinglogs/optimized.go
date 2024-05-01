package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Open the file
	filePath := "/home/yash/Desktop/original.log"

	file, err := os.Open(filePath)

	if err != nil {

		fmt.Println("Error opening file:", err)

		return

	}

	defer file.Close()

	// Create a new Scanner for reading the file line by line

	scanner := bufio.NewScanner(file)

	var jsonBuf bytes.Buffer

	jsonBuf.WriteString("[{}")

	fmt.Println("Starting scanning...")

	// Read the file line by line

	for scanner.Scan() {

		line := scanner.Text()

		if strings.Contains(line, "metric poll response: {") {

			jsonBuf.WriteString(",{")

		} else {

			jsonBuf.WriteString(line)

		}

	}

	jsonBuf.WriteString("]")

	fmt.Println("Scanning complete...")

	// Check if there was an error during the scan

	if err := scanner.Err(); err != nil {

		fmt.Println("Error reading file:", err)

	}

	var jsonArr []interface{}

	err = json.Unmarshal(jsonBuf.Bytes(), &jsonArr)

	if err != nil {

		fmt.Println("Error unmarshaling byte string:", err)

		return

	}

	objectTarget := os.Args[1]

	metricType := os.Args[2]

	for _, obj := range jsonArr {

		myMap, ok := obj.(map[string]interface{})

		if !ok || len(myMap) == 0 {
			continue
		}

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
