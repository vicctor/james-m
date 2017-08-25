package client

import (
	"log"
	"strings"
)

type Queue struct {
	QueueType         string `json:"queueType"`
	QueueSize         int    `json:"queueSize"`
	RemainingCapacity int    `json:"remainingCapacity"`
	DroppedItemsCount int    `json:"droppedItemsCount"`
}

type InformationPoint struct {
	ClassName                          string   `json:"className"`
	MethodName                         string   `json:"methodName"`
	Script                             []string `json:"script"`
	IncludeAbstractClass               bool     `json:"includeAbstractClass"`
	IncludeNonAbstractClassDescendants bool     `json:"includeNonAbstractClassDescendants"`
	SampleRate                         int      `json:"sampleRate"`
}

type MethodReference struct {
	ClassName  string
	MethodName string
}

func CreateMethodReference(encodedMethodReference string) (methodReference *MethodReference) {
	parts := strings.Split(encodedMethodReference, "#")
	if len(parts) != 2 {
		log.Fatal("Invalid method reference format, expected className#methodName")
	}

	methodReference = &MethodReference{
		ClassName:  parts[0],
		MethodName: parts[1],
	}
	return
}
