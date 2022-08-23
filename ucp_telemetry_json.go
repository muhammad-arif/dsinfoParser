package dsinfoParsingLibrary

import (
	"encoding/json"
	"fmt"
	"time"
)

type UCPTelemetry struct {
	AuditLogs                string      `json:"audit_logs"`
	ClusterId                string      `json:"cluster_id"`
	ClusterLabel             string      `json:"cluster_label"`
	ClusterStart             time.Time   `json:"cluster_start"`
	ContainerCount           int         `json:"container_count"`
	ContentTrust             bool        `json:"content_trust"`
	ControllerCount          int         `json:"controller_count"`
	Cpus                     int         `json:"cpus"`
	EngineVersion            string      `json:"engine_version"`
	GraphDriver              string      `json:"graph_driver"`
	ImageCount               int         `json:"image_count"`
	IsLicensed               bool        `json:"is_licensed"`
	KernelVersion            string      `json:"kernel_version"`
	KubeContainerCount       int         `json:"kube_container_count"`
	KubeCsidriverCount       int         `json:"kube_csidriver_count"`
	KubeCsidriverList        interface{} `json:"kube_csidriver_list"`
	KubeCustomerPodCount     int         `json:"kube_customer_pod_count"`
	KubeDaemonSetCount       int         `json:"kube_daemon_set_count"`
	KubeDeploymentCount      int         `json:"kube_deployment_count"`
	KubeGmsaEnabled          bool        `json:"kube_gmsa_enabled"`
	KubeGmsaServiceCount     int         `json:"kube_gmsa_service_count"`
	KubeGpuNodeCount         int         `json:"kube_gpu_node_count"`
	KubeGpuPodCount          int         `json:"kube_gpu_pod_count"`
	KubeIngressesCount       int         `json:"kube_ingresses_count"`
	KubeJobCount             int         `json:"kube_job_count"`
	KubeLinuxWorkerNodeCount int         `json:"kube_linux_worker_node_count"`
	KubeNamespaceCount       int         `json:"kube_namespace_count"`
	KubePodCount             int         `json:"kube_pod_count"`
	KubePvCount              int         `json:"kube_pv_count"`
	KubePvTypeCount          struct {
		AwsEBS     int `json:"awsEBS"`
		AzureDisk  int `json:"azureDisk"`
		AzureFile  int `json:"azureFile"`
		Ceph       int `json:"ceph"`
		Cinder     int `json:"cinder"`
		Csi        int `json:"csi"`
		Fc         int `json:"fc"`
		FlexVolume int `json:"flexVolume"`
		Flocker    int `json:"flocker"`
		GcePD      int `json:"gcePD"`
		Gluster    int `json:"gluster"`
		HostPath   int `json:"hostPath"`
		ISCSI      int `json:"iSCSI"`
		Local      int `json:"local"`
		Nfs        int `json:"nfs"`
		Photon     int `json:"photon"`
		Portworx   int `json:"portworx"`
		Quobyte    int `json:"quobyte"`
		Rbd        int `json:"rbd"`
		Scaleio    int `json:"scaleio"`
		StorageOS  int `json:"storageOS"`
		VSphere    int `json:"vSphere"`
	} `json:"kube_pv_type_count"`
	KubePvcCount                   int     `json:"kube_pvc_count"`
	KubeReplicaSetCount            int     `json:"kube_replica_set_count"`
	KubeReplicationControllerCount int     `json:"kube_replication_controller_count"`
	KubeSecretCount                int     `json:"kube_secret_count"`
	KubeServiceAccountCount        int     `json:"kube_service_account_count"`
	KubeServiceCount               int     `json:"kube_service_count"`
	KubeStackCount                 int     `json:"kube_stack_count"`
	KubeStatefulSetCount           int     `json:"kube_stateful_set_count"`
	KubeSystemPodCount             int     `json:"kube_system_pod_count"`
	KubeWinWorkerNodeCount         int     `json:"kube_win_worker_node_count"`
	KubeWorkerNodeCount            int     `json:"kube_worker_node_count"`
	LdapEnabled                    bool    `json:"ldap_enabled"`
	LicenseId                      string  `json:"license_id"`
	LicenseTier                    string  `json:"license_tier"`
	LoggingDriver                  string  `json:"logging_driver"`
	Memory                         float64 `json:"memory"`
	MirantisLicenseSubject         string  `json:"mirantis_license_subject"`
	NetworkCount                   int     `json:"network_count"`
	NginxIngressEnabled            bool    `json:"nginx_ingress_enabled"`
	NodeCount                      int     `json:"node_count"`
	NodeDiskInformation            struct {
		Nodes []struct {
			NumDisks           int   `json:"num_disks"`
			TotalDiskSizeBytes int64 `json:"total_disk_size_bytes"`
			NumNvmeDisks       int   `json:"num_nvme_disks"`
			TotalNvmeSizeBytes int   `json:"total_nvme_size_bytes"`
			RawDisks           []struct {
				Name              string `json:"name"`
				SizeBytes         int64  `json:"size_bytes"`
				DriveType         string `json:"drive_type"`
				StorageController string `json:"storage_controller"`
			} `json:"raw_disks"`
		} `json:"nodes"`
	} `json:"node_disk_information"`
	NodeStatistics []struct {
		Name    string `json:"Name"`
		ID      string `json:"ID"`
		CPUUsed string `json:"CPUUsed"`
		MemUsed string `json:"MemUsed"`
	} `json:"node_statistics"`
	OrcaVersion                string `json:"orca_version"`
	OrgCount                   int    `json:"org_count"`
	Os                         string `json:"os"`
	OsArchitecture             string `json:"os_architecture"`
	OsType                     string `json:"os_type"`
	SamlEnabled                bool   `json:"saml_enabled"`
	SamlIdp                    string `json:"saml_idp"`
	ScimEnabled                bool   `json:"scim_enabled"`
	SecretCount                int    `json:"secret_count"`
	ServiceCount               int    `json:"service_count"`
	SwarmGmsaServiceCount      int    `json:"swarm_gmsa_service_count"`
	SwarmGpuNodeCount          int    `json:"swarm_gpu_node_count"`
	SwarmLinuxWorkerNodeCount  int    `json:"swarm_linux_worker_node_count"`
	SwarmOnly                  bool   `json:"swarm_only"`
	SwarmStackCount            int    `json:"swarm_stack_count"`
	SwarmVersion               string `json:"swarm_version"`
	SwarmWinWorkerNodeCount    int    `json:"swarm_win_worker_node_count"`
	SwarmWorkerNodeCount       int    `json:"swarm_worker_node_count"`
	SwarmclassicContainerCount int    `json:"swarmclassic_container_count"`
	SwarmkitContainerCount     int    `json:"swarmkit_container_count"`
	SystemContainerCount       int    `json:"system_container_count"`
	TaskCount                  int    `json:"task_count"`
	UserCount                  int    `json:"user_count"`
	UsersLoggedIntoUcpCount    int    `json:"users_logged_into_ucp_count"`
	VolumeCount                int    `json:"volume_count"`
}

func NewUCPTelemetry() *UCPTelemetry {
	return &UCPTelemetry{}
}
func (c UCPTelemetry) GetAll(i *json.RawMessage) UCPTelemetry {
	var data UCPTelemetry
	err := json.Unmarshal(*i, &data)
	if err != nil {
		fmt.Errorf("[ucp_controller_config.go.GetAll()] : cannot unmarshal %s", err)
	}
	return data

}
