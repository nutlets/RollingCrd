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
			klog.Info("Add: ")
			printReplicaInfo(rs)
			return
		}
	}
}

func (h RSEventHandler) OnUpdate(oldObj, newObj interface{}) {
	rs1 := oldObj.(*appsv1.ReplicaSet)
	rs := newObj.(*appsv1.ReplicaSet)
	for _, v := range rs.OwnerReferences {
		if v.Kind != "Deployment" {
			klog.Warningf("unexpected type %s", v.Kind)
			continue
		}
		if CheckMonitored(rs.Namespace, v.Name) {
			klog.Info("Update: ")
			printReplicaInfo1(rs, rs1)
			return
		}
	}
}
func printReplicaInfo1(rs *appsv1.ReplicaSet, rs1 *appsv1.ReplicaSet) {

	klog.Infof("oldInfo %s/%s rs info: expected replica %d, current replica %d, ready replica %d %n", rs1.Namespace, rs1.Name, *rs1.Spec.Replicas, rs1.Status.Replicas, rs1.Status.ReadyReplicas)

	klog.Infof("newInfo %s/%s rs info: expected replica %d, current replica %d, ready replica %d", rs.Namespace, rs.Name, *rs.Spec.Replicas, rs.Status.Replicas, rs.Status.ReadyReplicas)
}
func printReplicaInfo(rs *appsv1.ReplicaSet) {

	klog.Infof("%s/%s rs info: expected replica %d, current replica %d, ready replica %d", rs.Namespace, rs.Name, *rs.Spec.Replicas, rs.Status.Replicas, rs.Status.ReadyReplicas)
}

func (h RSEventHandler) OnDelete(obj interface{}) {
	rs := obj.(*appsv1.ReplicaSet)
	for _, v := range rs.OwnerReferences {
		if v.Kind != "Deployment" {
			klog.Warningf("unexpected type %s", v.Kind)
			continue
		}
		if CheckMonitored(rs.Namespace, v.Name) {
			klog.Info("Delete: ")
			printReplicaInfo(rs)
			return
		}
	}
}
