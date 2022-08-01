package monitor

import (
	"k8s.io/client-go/informers"
)

func Start(factory informers.SharedInformerFactory) {
	rsInformer := factory.Apps().V1().ReplicaSets().Informer()
	rsInformer.AddEventHandler(&RSEventHandler{})
}
