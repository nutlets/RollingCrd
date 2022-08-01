package monitor

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/klog/v2"
)

type RSEventHandler struct{}

func (h RSEventHandler) OnAdd(obj interface{}) {
	rs := obj.(*appsv1.ReplicaSet)
	for _, v := range rs.OwnerReferences {
		if v.Kind != "Deployment" {
			klog.Warningf("unexpected type %s", v.Kind)
			continue
		}
		if CheckMonitored(rs.Namespace, v.Name) {
			printReplicaInfo(rs)
			return
		}
	}
}

func (h RSEventHandler) OnUpdate(oldObj, newObj interface{}) {
	rs := newObj.(*appsv1.ReplicaSet)
	for _, v := range rs.OwnerReferences {
		if v.Kind != "Deployment" {
			klog.Warningf("unexpected type %s", v.Kind)
			continue
		}
		if CheckMonitored(rs.Namespace, v.Name) {
			printReplicaInfo(rs)
			return
		}
	}
}

func printReplicaInfo(rs *appsv1.ReplicaSet) {

	klog.Infof("%s/%s rs info: expected replica %d, current replica %d, ready replica %d", rs.Namespace, rs.Name, *rs.Spec.Replicas, rs.Status.Replicas, rs.Status.ReadyReplicas)
}

func (h RSEventHandler) OnDelete(obj interface{}) {

}
