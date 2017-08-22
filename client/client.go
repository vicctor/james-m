package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"bytes"
)

type JamesClient struct {
	Host string
	Port uint16
}

func (jamesClient *JamesClient) GetInformationPoints() (ips []InformationPoint) {
	jsonResponse := httpGet(fmt.Sprintf("http://%s:%d/v1/information-point", jamesClient.Host, jamesClient.Port))
	json.Unmarshal(jsonResponse, &ips)
	return
}

func (jamesClient *JamesClient) GetInformationPoint(methodReference *MethodReference) (ip InformationPoint) {
	jsonResponse := httpGet(fmt.Sprintf("http://%s:%d/v1/information-point/%s/%s",
		jamesClient.Host, jamesClient.Port, methodReference.ClassName, methodReference.MethodName))
	json.Unmarshal(jsonResponse, &ip)
	return
}

func (jamesClient *JamesClient) AddInformationPoint(
	methodReference *MethodReference,
	script []byte,
	includeAbstractClass bool,
	includeNonAbstractClassDescendants bool) {

	noCrScript := strings.Replace(string(script), "\r", "", -1)
	scriptLines := strings.Split(noCrScript, "\n")
	filteredLines := filter(scriptLines, func(s string) bool {
		return len(strings.TrimSpace(s)) > 0
	})

	informationPoint := &InformationPoint{
		ClassName:                          methodReference.ClassName,
		MethodName:                         methodReference.MethodName,
		Script:                             filteredLines,
		IncludeAbstractClass:               includeAbstractClass,
		IncludeNonAbstractClassDescendants: includeNonAbstractClassDescendants,
	}

	url := fmt.Sprintf("http://%s:%d/v1/information-point", jamesClient.Host, jamesClient.Port)
	informationPointJson, _ := json.Marshal(informationPoint)
	httpPost(url, informationPointJson)
}

func (jamesClient *JamesClient) RemoveInformationPoint(methodReference *MethodReference) {
	httpDelete(fmt.Sprintf("http://%s:%d/v1/information-point/%s/%s",
		jamesClient.Host, jamesClient.Port, methodReference.ClassName, methodReference.MethodName))
}

func (jamesClient *JamesClient) GetAllQueues() (queues []Queue) {
	jsonResponse := httpGet(fmt.Sprintf("http://%s:%d/v1/queue", jamesClient.Host, jamesClient.Port))
	json.Unmarshal(jsonResponse, &queues)
	return
}

func (jamesClient *JamesClient) GetScriptEngineQueue() (queue Queue) {
	jsonResponse := httpGet(fmt.Sprintf("http://%s:%d/v1/queue/script-engine", jamesClient.Host, jamesClient.Port))
	json.Unmarshal(jsonResponse, &queue)
	return
}

func (jamesClient *JamesClient) GetEventPublisherQueue() (queue Queue) {
	jsonResponse := httpGet(fmt.Sprintf("http://%s:%d/v1/queue/event-publisher", jamesClient.Host, jamesClient.Port))
	json.Unmarshal(jsonResponse, &queue)
	return
}

func httpGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	checkStatusCode(resp.StatusCode)
	return readResponseBody(resp.Body)
}

func httpPost(url string, body []byte) {
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	checkStatusCode(resp.StatusCode)
}

func httpDelete(url string) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	checkStatusCode(resp.StatusCode)
}

func checkStatusCode(code int) {
	if code != http.StatusOK {
		log.Fatal(fmt.Sprintf("Service returned status code %d (%s)", code, http.StatusText(code)))
	}
}

func readResponseBody(reader io.Reader) []byte {
	result, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
