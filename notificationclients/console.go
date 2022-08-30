package notificationclients

import (
	"fmt"

	"k8s.io/apimachinery/pkg/watch"
)

func Console(data watch.Event)  bool  {
	fmt.Println(data)
	return true
}