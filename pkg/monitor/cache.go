package monitor

import (
	"fmt"
	"sync"
)

var (
	cache = sync.Map{}
)

func AddMonitorDeploy(namespace string, name string) {
	key := geneKey(namespace, name)
	if _, ok := cache.Load(key); !ok {
		cache.Store(key, true)
	}
}

func RemoveMonitoredDeploy(namespace string, name string) {
	key := geneKey(namespace, name)
	cache.Delete(key)
}

func CheckMonitored(namespace string, name string) bool {
	key := geneKey(namespace, name)
	_, ok := cache.Load(key)
	return ok
}

func geneKey(namespace string, name string) string {
	return fmt.Sprintf("%s_%s", namespace, name)
}
