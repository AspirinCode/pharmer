package gce

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/appscode/go/crypto/rand"
	"github.com/appscode/mergo"
	api "github.com/appscode/pharmer/apis/v1alpha1"
	. "github.com/appscode/pharmer/cloud"
	"github.com/appscode/pharmer/data/files"
	"github.com/appscode/pharmer/phid"
	semver "github.com/hashicorp/go-version"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1"
)

func (cm *ClusterManager) CreateMasterNodeGroup(cluster *api.Cluster) (*api.NodeGroup, error) {
	ig := api.NodeGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:              "master",
			ClusterName:       cluster.Name,
			UID:               phid.NewNodeGroup(),
			CreationTimestamp: metav1.Time{Time: time.Now()},
			Labels: map[string]string{
				api.RoleMasterKey: "true",
			},
		},
		Spec: api.NodeGroupSpec{
			Nodes: 1,
			Template: api.NodeTemplateSpec{
				Spec: api.NodeSpec{
					SKU:           "n1-standard-2", // assign at the time of apply
					SpotInstances: false,
					DiskType:      "pd-standard",
					DiskSize:      100,
				},
			},
		},
	}
	return Store(cm.ctx).NodeGroups(cluster.Name).Create(&ig)
}

func (cm *ClusterManager) DefaultSpec(in *api.Cluster) (*api.Cluster, error) {
	// Load default spec from data files
	kv, err := semver.NewVersion(in.Spec.KubernetesVersion)
	if err != nil {
		return nil, err
	}
	defaultSpec, err := files.GetDefaultClusterSpec(in.Spec.Cloud.CloudProvider, kv)
	if err != nil {
		return nil, err
	}
	cluster := &api.Cluster{
		Spec: *defaultSpec,
	}

	// Copy default spec into return value
	err = mergo.MergeWithOverwrite(cluster, in)
	if err != nil {
		return nil, err
	}
	n := namer{cluster: cluster}

	// Init object meta
	cluster.ObjectMeta.UID = phid.NewKubeCluster()
	cluster.ObjectMeta.CreationTimestamp = metav1.Time{Time: time.Now()}
	cluster.ObjectMeta.Generation = time.Now().UnixNano()
	api.AssignTypeKind(cluster)

	// Init spec
	cluster.Spec.Cloud.Region = cluster.Spec.Cloud.Zone[0:strings.LastIndex(cluster.Spec.Cloud.Zone, "-")]
	cluster.Spec.API.BindPort = kubeadmapi.DefaultAPIBindPort
	// REGISTER_MASTER_KUBELET = false // always false, keep master lightweight
	// PREEMPTIBLE_NODE = false // Removed Support

	// Init status
	cluster.Status = api.ClusterStatus{
		Phase:            api.ClusterPending,
		SSHKeyExternalID: n.GenSSHKeyExternalID(),
		Cloud: api.CloudStatus{
			GCE: &api.GCEStatus{
				BucketName: rand.WithUniqSuffix("pharmer-config"),
			},
		},
	}
	cluster.Spec.Cloud.InstanceImageProject = "ubuntu-os-cloud"
	cluster.Spec.Cloud.InstanceImage = "ubuntu-1604-xenial-v20170721"
	cluster.Spec.Networking.NonMasqueradeCIDR = "10.0.0.0/8"
	cluster.Spec.Networking.PodSubnet = "10.244.0.0/16"

	return cluster, nil
}

func (cm *ClusterManager) IsValid(cluster *api.Cluster) (bool, error) {
	return false, ErrNotImplemented
}

func (cm *ClusterManager) GetSSHConfig(cluster *api.Cluster, node *core.Node) (*api.SSHConfig, error) {
	n := namer{cluster: cluster}
	cfg := &api.SSHConfig{
		PrivateKey:   SSHKey(cm.ctx).PrivateKey,
		User:         n.AdminUsername(),
		InstancePort: int32(22),
	}
	for _, addr := range node.Status.Addresses {
		if addr.Type == core.NodeExternalIP {
			cfg.InstanceAddress = addr.Address
		}
	}
	if net.ParseIP(cfg.InstanceAddress) == nil {
		return nil, fmt.Errorf("failed to detect external Ip for node %s of cluster %s", node.Name, cluster.Name)
	}
	return cfg, nil
}
