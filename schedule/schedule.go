package schedule

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Execute will run the schedule
func Execute(target string, image string) {
	// Build the payload
	// The payload is a string to pass information into your worker as part of a task
	// It generally is a JSON-serialized string (which is what we're doing here) that can be deserialized in the worker
	payload := map[string]string{
		"text": "This is the comment that I'm testing. A beautiful story.",
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err.Error())
	}
	payloadStr := string(payloadBytes)

	// Build the job
	job := &Job{
		Image:   image,
		Payload: payloadStr,
	}

	// Build a request containing the task
	jsonData := &ReqData{
		Jobs: []*Job{job},
	}
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		panic(err.Error())
	}
	jsonStr := string(jsonBytes)

	// Post expects a Reader
	jsonBuf := bytes.NewBufferString(jsonStr)

	// Make the request
	resp, err := http.Post(target, "application/json", jsonBuf)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	// Read the response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	// Print the response to STDOUT
	fmt.Println(string(respBody))
}
