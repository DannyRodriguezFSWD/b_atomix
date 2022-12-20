//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta2

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSinkConfig) DeepCopyInto(out *FileSinkConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSinkConfig.
func (in *FileSinkConfig) DeepCopy() *FileSinkConfig {
	if in == nil {
		return nil
	}
	out := new(FileSinkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggerConfig) DeepCopyInto(out *LoggerConfig) {
	*out = *in
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = make(map[string]OutputConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggerConfig.
func (in *LoggerConfig) DeepCopy() *LoggerConfig {
	if in == nil {
		return nil
	}
	out := new(LoggerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfig) DeepCopyInto(out *LoggingConfig) {
	*out = *in
	if in.Loggers != nil {
		in, out := &in.Loggers, &out.Loggers
		*out = make(map[string]LoggerConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Sinks != nil {
		in, out := &in.Sinks, &out.Sinks
		*out = make(map[string]SinkConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfig.
func (in *LoggingConfig) DeepCopy() *LoggingConfig {
	if in == nil {
		return nil
	}
	out := new(LoggingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRaftCluster) DeepCopyInto(out *MultiRaftCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRaftCluster.
func (in *MultiRaftCluster) DeepCopy() *MultiRaftCluster {
	if in == nil {
		return nil
	}
	out := new(MultiRaftCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiRaftCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRaftClusterConfig) DeepCopyInto(out *MultiRaftClusterConfig) {
	*out = *in
	out.RTT = in.RTT
	in.Server.DeepCopyInto(&out.Server)
	in.Logging.DeepCopyInto(&out.Logging)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRaftClusterConfig.
func (in *MultiRaftClusterConfig) DeepCopy() *MultiRaftClusterConfig {
	if in == nil {
		return nil
	}
	out := new(MultiRaftClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRaftClusterList) DeepCopyInto(out *MultiRaftClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiRaftCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRaftClusterList.
func (in *MultiRaftClusterList) DeepCopy() *MultiRaftClusterList {
	if in == nil {
		return nil
	}
	out := new(MultiRaftClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiRaftClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRaftClusterPartitionStatus) DeepCopyInto(out *MultiRaftClusterPartitionStatus) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRaftClusterPartitionStatus.
func (in *MultiRaftClusterPartitionStatus) DeepCopy() *MultiRaftClusterPartitionStatus {
	if in == nil {
		return nil
	}
	out := new(MultiRaftClusterPartitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRaftClusterSpec) DeepCopyInto(out *MultiRaftClusterSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeClaimTemplate != nil {
		in, out := &in.VolumeClaimTemplate, &out.VolumeClaimTemplate
		*out = new(v1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRaftClusterSpec.
func (in *MultiRaftClusterSpec) DeepCopy() *MultiRaftClusterSpec {
	if in == nil {
		return nil
	}
	out := new(MultiRaftClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRaftClusterStatus) DeepCopyInto(out *MultiRaftClusterStatus) {
	*out = *in
	if in.PartitionStatuses != nil {
		in, out := &in.PartitionStatuses, &out.PartitionStatuses
		*out = make([]MultiRaftClusterPartitionStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRaftClusterStatus.
func (in *MultiRaftClusterStatus) DeepCopy() *MultiRaftClusterStatus {
	if in == nil {
		return nil
	}
	out := new(MultiRaftClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRaftServerConfig) DeepCopyInto(out *MultiRaftServerConfig) {
	*out = *in
	if in.ReadBufferSize != nil {
		in, out := &in.ReadBufferSize, &out.ReadBufferSize
		*out = new(int)
		**out = **in
	}
	if in.WriteBufferSize != nil {
		in, out := &in.WriteBufferSize, &out.WriteBufferSize
		*out = new(int)
		**out = **in
	}
	if in.MaxRecvMsgSize != nil {
		in, out := &in.MaxRecvMsgSize, &out.MaxRecvMsgSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.MaxSendMsgSize != nil {
		in, out := &in.MaxSendMsgSize, &out.MaxSendMsgSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.NumStreamWorkers != nil {
		in, out := &in.NumStreamWorkers, &out.NumStreamWorkers
		*out = new(uint32)
		**out = **in
	}
	if in.MaxConcurrentStreams != nil {
		in, out := &in.MaxConcurrentStreams, &out.MaxConcurrentStreams
		*out = new(uint32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRaftServerConfig.
func (in *MultiRaftServerConfig) DeepCopy() *MultiRaftServerConfig {
	if in == nil {
		return nil
	}
	out := new(MultiRaftServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRaftStore) DeepCopyInto(out *MultiRaftStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRaftStore.
func (in *MultiRaftStore) DeepCopy() *MultiRaftStore {
	if in == nil {
		return nil
	}
	out := new(MultiRaftStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiRaftStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRaftStoreList) DeepCopyInto(out *MultiRaftStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiRaftStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRaftStoreList.
func (in *MultiRaftStoreList) DeepCopy() *MultiRaftStoreList {
	if in == nil {
		return nil
	}
	out := new(MultiRaftStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiRaftStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRaftStoreSpec) DeepCopyInto(out *MultiRaftStoreSpec) {
	*out = *in
	in.RaftConfig.DeepCopyInto(&out.RaftConfig)
	out.Cluster = in.Cluster
	if in.ReplicationFactor != nil {
		in, out := &in.ReplicationFactor, &out.ReplicationFactor
		*out = new(uint32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRaftStoreSpec.
func (in *MultiRaftStoreSpec) DeepCopy() *MultiRaftStoreSpec {
	if in == nil {
		return nil
	}
	out := new(MultiRaftStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRaftStoreStatus) DeepCopyInto(out *MultiRaftStoreStatus) {
	*out = *in
	if in.ReplicationFactor != nil {
		in, out := &in.ReplicationFactor, &out.ReplicationFactor
		*out = new(uint32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRaftStoreStatus.
func (in *MultiRaftStoreStatus) DeepCopy() *MultiRaftStoreStatus {
	if in == nil {
		return nil
	}
	out := new(MultiRaftStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputConfig) DeepCopyInto(out *OutputConfig) {
	*out = *in
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(string)
		**out = **in
	}
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputConfig.
func (in *OutputConfig) DeepCopy() *OutputConfig {
	if in == nil {
		return nil
	}
	out := new(OutputConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftConfig) DeepCopyInto(out *RaftConfig) {
	*out = *in
	if in.ElectionRTT != nil {
		in, out := &in.ElectionRTT, &out.ElectionRTT
		*out = new(uint64)
		**out = **in
	}
	if in.HeartbeatRTT != nil {
		in, out := &in.HeartbeatRTT, &out.HeartbeatRTT
		*out = new(uint64)
		**out = **in
	}
	if in.SnapshotEntries != nil {
		in, out := &in.SnapshotEntries, &out.SnapshotEntries
		*out = new(int64)
		**out = **in
	}
	if in.CompactionOverhead != nil {
		in, out := &in.CompactionOverhead, &out.CompactionOverhead
		*out = new(int64)
		**out = **in
	}
	if in.MaxInMemLogSize != nil {
		in, out := &in.MaxInMemLogSize, &out.MaxInMemLogSize
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftConfig.
func (in *RaftConfig) DeepCopy() *RaftConfig {
	if in == nil {
		return nil
	}
	out := new(RaftConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftMember) DeepCopyInto(out *RaftMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftMember.
func (in *RaftMember) DeepCopy() *RaftMember {
	if in == nil {
		return nil
	}
	out := new(RaftMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftMemberConfig) DeepCopyInto(out *RaftMemberConfig) {
	*out = *in
	in.RaftConfig.DeepCopyInto(&out.RaftConfig)
	if in.Peers != nil {
		in, out := &in.Peers, &out.Peers
		*out = make([]RaftMemberReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftMemberConfig.
func (in *RaftMemberConfig) DeepCopy() *RaftMemberConfig {
	if in == nil {
		return nil
	}
	out := new(RaftMemberConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftMemberList) DeepCopyInto(out *RaftMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RaftMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftMemberList.
func (in *RaftMemberList) DeepCopy() *RaftMemberList {
	if in == nil {
		return nil
	}
	out := new(RaftMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftMemberReference) DeepCopyInto(out *RaftMemberReference) {
	*out = *in
	out.Pod = in.Pod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftMemberReference.
func (in *RaftMemberReference) DeepCopy() *RaftMemberReference {
	if in == nil {
		return nil
	}
	out := new(RaftMemberReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftMemberSpec) DeepCopyInto(out *RaftMemberSpec) {
	*out = *in
	out.Cluster = in.Cluster
	out.Pod = in.Pod
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftMemberSpec.
func (in *RaftMemberSpec) DeepCopy() *RaftMemberSpec {
	if in == nil {
		return nil
	}
	out := new(RaftMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftMemberStatus) DeepCopyInto(out *RaftMemberStatus) {
	*out = *in
	if in.PodRef != nil {
		in, out := &in.PodRef, &out.PodRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(int32)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(RaftMemberRole)
		**out = **in
	}
	if in.Leader != nil {
		in, out := &in.Leader, &out.Leader
		*out = new(MemberID)
		**out = **in
	}
	if in.Term != nil {
		in, out := &in.Term, &out.Term
		*out = new(uint64)
		**out = **in
	}
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = (*in).DeepCopy()
	}
	if in.LastSnapshotIndex != nil {
		in, out := &in.LastSnapshotIndex, &out.LastSnapshotIndex
		*out = new(uint64)
		**out = **in
	}
	if in.LastSnapshotTime != nil {
		in, out := &in.LastSnapshotTime, &out.LastSnapshotTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftMemberStatus.
func (in *RaftMemberStatus) DeepCopy() *RaftMemberStatus {
	if in == nil {
		return nil
	}
	out := new(RaftMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftPartition) DeepCopyInto(out *RaftPartition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftPartition.
func (in *RaftPartition) DeepCopy() *RaftPartition {
	if in == nil {
		return nil
	}
	out := new(RaftPartition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftPartition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftPartitionList) DeepCopyInto(out *RaftPartitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RaftPartition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftPartitionList.
func (in *RaftPartitionList) DeepCopy() *RaftPartitionList {
	if in == nil {
		return nil
	}
	out := new(RaftPartitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftPartitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftPartitionMemberStatus) DeepCopyInto(out *RaftPartitionMemberStatus) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftPartitionMemberStatus.
func (in *RaftPartitionMemberStatus) DeepCopy() *RaftPartitionMemberStatus {
	if in == nil {
		return nil
	}
	out := new(RaftPartitionMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftPartitionSpec) DeepCopyInto(out *RaftPartitionSpec) {
	*out = *in
	in.RaftConfig.DeepCopyInto(&out.RaftConfig)
	out.Cluster = in.Cluster
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftPartitionSpec.
func (in *RaftPartitionSpec) DeepCopy() *RaftPartitionSpec {
	if in == nil {
		return nil
	}
	out := new(RaftPartitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftPartitionStatus) DeepCopyInto(out *RaftPartitionStatus) {
	*out = *in
	if in.Term != nil {
		in, out := &in.Term, &out.Term
		*out = new(uint64)
		**out = **in
	}
	if in.Leader != nil {
		in, out := &in.Leader, &out.Leader
		*out = new(MemberID)
		**out = **in
	}
	if in.Followers != nil {
		in, out := &in.Followers, &out.Followers
		*out = make([]MemberID, len(*in))
		copy(*out, *in)
	}
	if in.MemberStatuses != nil {
		in, out := &in.MemberStatuses, &out.MemberStatuses
		*out = make([]RaftPartitionMemberStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftPartitionStatus.
func (in *RaftPartitionStatus) DeepCopy() *RaftPartitionStatus {
	if in == nil {
		return nil
	}
	out := new(RaftPartitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftStore) DeepCopyInto(out *RaftStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftStore.
func (in *RaftStore) DeepCopy() *RaftStore {
	if in == nil {
		return nil
	}
	out := new(RaftStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftStoreList) DeepCopyInto(out *RaftStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RaftStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftStoreList.
func (in *RaftStoreList) DeepCopy() *RaftStoreList {
	if in == nil {
		return nil
	}
	out := new(RaftStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RaftStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftStoreSpec) DeepCopyInto(out *RaftStoreSpec) {
	*out = *in
	in.RaftConfig.DeepCopyInto(&out.RaftConfig)
	out.Cluster = in.Cluster
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftStoreSpec.
func (in *RaftStoreSpec) DeepCopy() *RaftStoreSpec {
	if in == nil {
		return nil
	}
	out := new(RaftStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftStoreStatus) DeepCopyInto(out *RaftStoreStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftStoreStatus.
func (in *RaftStoreStatus) DeepCopy() *RaftStoreStatus {
	if in == nil {
		return nil
	}
	out := new(RaftStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkConfig) DeepCopyInto(out *SinkConfig) {
	*out = *in
	if in.Encoding != nil {
		in, out := &in.Encoding, &out.Encoding
		*out = new(string)
		**out = **in
	}
	if in.Stdout != nil {
		in, out := &in.Stdout, &out.Stdout
		*out = new(StdoutSinkConfig)
		**out = **in
	}
	if in.Stderr != nil {
		in, out := &in.Stderr, &out.Stderr
		*out = new(StderrSinkConfig)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(FileSinkConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkConfig.
func (in *SinkConfig) DeepCopy() *SinkConfig {
	if in == nil {
		return nil
	}
	out := new(SinkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StderrSinkConfig) DeepCopyInto(out *StderrSinkConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StderrSinkConfig.
func (in *StderrSinkConfig) DeepCopy() *StderrSinkConfig {
	if in == nil {
		return nil
	}
	out := new(StderrSinkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StdoutSinkConfig) DeepCopyInto(out *StdoutSinkConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StdoutSinkConfig.
func (in *StdoutSinkConfig) DeepCopy() *StdoutSinkConfig {
	if in == nil {
		return nil
	}
	out := new(StdoutSinkConfig)
	in.DeepCopyInto(out)
	return out
}
