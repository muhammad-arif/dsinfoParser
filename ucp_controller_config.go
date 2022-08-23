package dsinfoParsingLibrary

import (
	"encoding/json"
	"fmt"
)

type UCPControllerConfigType struct {
	Analytics struct {
		DisableUsageInfo bool   `json:"disableUsageInfo"`
		DisableTracking  bool   `json:"disableTracking"`
		ClusterLabel     string `json:"clusterLabel"`
		OpsCare          bool   `json:"opsCare"`
	} `json:"analytics"`
	BackupSchedule struct {
		Enabled           bool   `json:"enabled"`
		Path              string `json:"path"`
		Passphrase        string `json:"passphrase"`
		NoPassphrase      bool   `json:"noPassphrase"`
		CronSpec          string `json:"cronSpec"`
		IncludeLogs       bool   `json:"includeLogs"`
		BackupsLimit      int    `json:"backupsLimit"`
		NotificationDelay int    `json:"notificationDelay"`
	} `json:"backupSchedule"`
	Connectivity struct {
		ExternalServiceLB          string      `json:"externalServiceLB"`
		ControllerPort             int         `json:"controllerPort"`
		KubeAPIServerPort          int         `json:"kubeAPIServerPort"`
		SwarmPort                  int         `json:"swarmPort"`
		Dns                        interface{} `json:"dns"`
		DnsOpt                     interface{} `json:"dnsOpt"`
		DnsSearch                  interface{} `json:"dnsSearch"`
		NetworksQuotaPerCollection int         `json:"NetworksQuotaPerCollection"`
	} `json:"connectivity"`
	DockerContentTrust struct {
		RequireContentTrust  bool          `json:"requireContentTrust"`
		RequireSignatureFrom []interface{} `json:"requireSignatureFrom"`
		AllowRepos           []interface{} `json:"AllowRepos"`
	} `json:"dockerContentTrust"`
	DockerTrustedRegistries []struct {
		HostAddress              string `json:"hostAddress"`
		ServiceID                string `json:"serviceID"`
		CaBundle                 string `json:"caBundle"`
		BatchScanningDataEnabled bool   `json:"batchScanningDataEnabled"`
	} `json:"dockerTrustedRegistries"`
	ExternalIdentityProvider struct {
		Issuer             string      `json:"issuer"`
		UserServiceId      string      `json:"userServiceId"`
		ClientId           string      `json:"clientId"`
		WellKnownConfigUrl string      `json:"wellKnownConfigUrl"`
		CaBundle           string      `json:"caBundle"`
		AdminRoleCriteria  interface{} `json:"adminRoleCriteria"`
		SignInCriteria     interface{} `json:"signInCriteria"`
		UsernameClaim      string      `json:"usernameClaim"`
		HttpProxy          string      `json:"httpProxy"`
		HttpsProxy         string      `json:"httpsProxy"`
	} `json:"externalIdentityProvider"`
	HttpServer struct {
		CustomAPIServerHeaders       interface{} `json:"customAPIServerHeaders"`
		ExcludeServerIdentityHeaders bool        `json:"excludeServerIdentityHeaders"`
		ClientCABundle               string      `json:"clientCABundle"`
		PreLogonMessage              string      `json:"preLogonMessage"`
	} `json:"httpServer"`
	Kubernetes struct {
		CloudProvider             string `json:"cloudProvider"`
		CniPodCIDR                string `json:"cniPodCIDR"`
		ServiceClusterIPRange     string `json:"serviceClusterIPRange"`
		NodePortRange             string `json:"nodePortRange"`
		UnmanagedCNI              bool   `json:"unmanagedCNI"`
		CalicoIPAutoMethod        string `json:"CalicoIPAutoMethod"`
		CalicoMTU                 string `json:"calicoMTU"`
		CalicoStrictAffinity      bool   `json:"calicoStrictAffinity"`
		CalicoVXLANEnabled        bool   `json:"CalicoVXLANEnabled"`
		CalicoVXLANVNI            int    `json:"CalicoVXLANVNI"`
		CalicoVXLANMTU            string `json:"CalicoVXLANMTU"`
		CalicoVXLANPort           string `json:"CalicoVXLANPort"`
		CalicoEbpfEnabled         bool   `json:"CalicoEbpfEnabled"`
		KubeProxyMode             string `json:"KubeProxyMode"`
		KubeDefaultDropMasqBits   bool   `json:"KubeDefaultDropMasqBits"`
		KubeProxyNoCleanupOnStart bool   `json:"KubeProxyNoCleanupOnStart"`
		IPVSExcludeCIDRs          string `json:"IPVSExcludeCIDRs"`
		IPVSMinSyncPeriod         string `json:"IPVSMinSyncPeriod"`
		IPVSScheduler             string `json:"IPVSScheduler"`
		IPVSStrictARP             bool   `json:"IPVSStrictARP"`
		IPVSSyncPeriod            string `json:"IPVSSyncPeriod"`
		IPVSTCPTimeout            string `json:"IPVSTCPTimeout"`
		IPVSTCPFinTimeout         string `json:"IPVSTCPFinTimeout"`
		IPVSUDPTimeout            string `json:"IPVSUDPTimeout"`
		IpipMTU                   string `json:"ipipMTU"`
		AzureIPCount              string `json:"azureIPCount"`
		ServiceMesh               struct {
			Enabled                 bool        `json:"enabled"`
			IngressExposedPorts     interface{} `json:"ingressExposedPorts"`
			IngressNumReplicas      int         `json:"ingressNumReplicas"`
			IngressExternalIPs      interface{} `json:"ingressExternalIPs"`
			IngressNodeAffinity     interface{} `json:"ingressNodeAffinity"`
			IngressNodeToleration   interface{} `json:"ingressNodeToleration"`
			IngressEnableLB         bool        `json:"ingressEnableLB"`
			IngressPreserveClientIP bool        `json:"ingressPreserveClientIP"`
			IngressConfigMap        interface{} `json:"ingressConfigMap"`
		} `json:"serviceMesh"`
		IngressController struct {
			Enabled             bool `json:"enabled"`
			IngressExposedPorts []struct {
				Name       string `json:"name"`
				Port       int    `json:"port"`
				TargetPort int    `json:"targetPort"`
				NodePort   int    `json:"nodePort"`
			} `json:"ingressExposedPorts"`
			IngressNumReplicas    int         `json:"ingressNumReplicas"`
			IngressExternalIPs    interface{} `json:"ingressExternalIPs"`
			IngressNodeAffinity   interface{} `json:"ingressNodeAffinity"`
			IngressNodeToleration []struct {
				Key      string `json:"key"`
				Value    string `json:"value"`
				Operator string `json:"operator"`
				Effect   string `json:"effect"`
			} `json:"ingressNodeToleration"`
			IngressEnableLB         bool `json:"ingressEnableLB"`
			IngressPreserveClientIP bool `json:"ingressPreserveClientIP"`
			IngressConfigMap        struct {
			} `json:"ingressConfigMap"`
		} `json:"ingressController"`
		SecureOverlay                           bool        `json:"secureOverlay"`
		ManagerReservedResources                string      `json:"managerReservedResources"`
		WorkerReservedResources                 string      `json:"workerReservedResources"`
		KubeletMaxPods                          int         `json:"kubeletMaxPods"`
		KubeletPodsPerCore                      int         `json:"kubeletPodsPerCore"`
		WindowsGMSA                             bool        `json:"windowsGMSA"`
		NvidiaDevicePlugin                      bool        `json:"nvidiaDevicePlugin"`
		ClusterName                             string      `json:"clusterName"`
		PrivAttributesAllowedForUserAccounts    interface{} `json:"privAttributesAllowedForUserAccounts"`
		PrivAttributesUserAccounts              interface{} `json:"privAttributesUserAccounts"`
		PrivAttributesAllowedForServiceAccounts interface{} `json:"privAttributesAllowedForServiceAccounts"`
		PrivAttributesServiceAccounts           interface{} `json:"privAttributesServiceAccounts"`
		KubeletDataRoot                         string      `json:"kubeletDataRoot"`
		WindowsKubeletDataRoot                  string      `json:"windowsKubeletDataRoot"`
		Storage                                 struct {
			Iscsi struct {
				Enabled bool `json:"enabled"`
			} `json:"iscsi"`
			StorageExptEnabled bool `json:"storageExptEnabled"`
		} `json:"storage"`
		IgnorePodUpdatesForNodeSelector bool        `json:"ignorePodUpdatesForNodeSelector"`
		KMSEnabled                      bool        `json:"KMSEnabled"`
		KMSName                         string      `json:"KMSName"`
		KMSEndpoint                     string      `json:"KMSEndpoint"`
		KMSCachesize                    int         `json:"KMSCachesize"`
		CipherSuitesForAPIServer        string      `json:"cipherSuitesForAPIServer"`
		CipherSuitesForKubelet          string      `json:"cipherSuitesForKubelet"`
		CipherSuitesForEtcdServer       string      `json:"cipherSuitesForEtcdServer"`
		CustomAPIServerFlags            interface{} `json:"customAPIServerFlags"`
		CustomControllerManagerFlags    interface{} `json:"customControllerManagerFlags"`
		CustomKubeletFlags              interface{} `json:"customKubeletFlags"`
		CustomSchedulerFlags            interface{} `json:"customSchedulerFlags"`
		CustomProxyFlags                interface{} `json:"customProxyFlags"`
	} `json:"kubernetes"`
	LicenseConfig struct {
		AutoRefresh bool `json:"autoRefresh"`
	} `json:"licenseConfig"`
	Logging struct {
		LogLevel                      string `json:"logLevel"`
		AuditLevel                    string `json:"auditLevel"`
		SupportBundleIncludeAuditLogs bool   `json:"supportBundleIncludeAuditLogs"`
	} `json:"logging"`
	Orchestration struct {
		EnableAdminSystemScheduling              bool        `json:"enableAdminSystemScheduling"`
		DefaultNodeOrchestrator                  string      `json:"defaultNodeOrchestrator"`
		SwarmSchedulingStrategy                  string      `json:"swarmSchedulingStrategy"`
		SwarmLocalVolumeCollectionMapping        bool        `json:"swarmLocalVolumeCollectionMapping"`
		DefaultNewUserSwarmPrivateCollectionRole string      `json:"defaultNewUserSwarmPrivateCollectionRole"`
		SwarmDefaults                            interface{} `json:"swarmDefaults"`
		ContainerdRootDir                        string      `json:"containerdRootDir"`
		WindowsContainerdRootDir                 string      `json:"windowsContainerdRootDir"`
	} `json:"orchestration"`
	Tuning struct {
		ProfilingEnabled            bool   `json:"profilingEnabled"`
		EtcdTimeout                 int    `json:"etcdTimeout"`
		EtcdSnapshotCount           int    `json:"etcdSnapshotCount"`
		RethinkDBCacheSize          string `json:"rethinkDBCacheSize"`
		MetricsRetentionTime        string `json:"metricsRetentionTime"`
		MetricsScrapeInterval       string `json:"metricsScrapeInterval"`
		SwarmPollingDisabled        bool   `json:"swarmPollingDisabled"`
		HideSwarmUI                 bool   `json:"hideSwarmUI"`
		SwarmOnly                   bool   `json:"swarmOnly"`
		ImageScanAggregationEnabled bool   `json:"imageScanAggregationEnabled"`
		ImageScanCvssVersion        int    `json:"imageScanCvssVersion"`
	} `json:"tuning"`
}

func NewUCPControllerConfigType() *UCPControllerConfigType {
	return &UCPControllerConfigType{}
}
func (c UCPControllerConfigType) GetAll(i *json.RawMessage) UCPControllerConfigType {
	var data UCPControllerConfigType
	err := json.Unmarshal(*i, &data)
	if err != nil {
		fmt.Errorf("[ucp_controller_config.go.GetAll()] : cannot unmarshal %s", err)
	}
	return data

}
