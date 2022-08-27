package dsinfoParsingLibrary

import (
	"encoding/json"
)

type CoreDsinfo struct {
	AuthConfigTxt                 []interface{}   `json:"auth-config.txt"`
	Configs                       json.RawMessage `json:"configs"`
	IngressDescribeAllTxt         json.RawMessage `json:"ingress-describe-all.txt"`
	IngressDescribeIngressTxt     json.RawMessage `json:"ingress-describe-ingress.txt"`
	IngressGetAllTxt              json.RawMessage `json:"ingress-get-all.txt"`
	IngressGetIngressTxt          json.RawMessage `json:"ingress-get-ingress.txt"`
	IngressNginxConfigTxt         json.RawMessage `json:"ingress-nginx-config.txt"`
	IngressNginxControllerTxt     json.RawMessage `json:"ingress-nginx-controller.txt"`
	KubeDescribeAllTxt            json.RawMessage `json:"kube-describe-all.txt"`
	KubeDescribeNodesTxt          json.RawMessage `json:"kube-describe-nodes.txt"`
	MirantisLic                   json.RawMessage `json:"mirantis.lic"`
	TasksCountTxt                 []string        `json:"tasks-count.txt"`
	UcpAuthApiServiceTxt          json.RawMessage `json:"ucp-auth-api-service.txt"`
	UcpAuthApiTasksTxt            json.RawMessage `json:"ucp-auth-api-tasks.txt"`
	UcpAuthWorkerServiceTxt       json.RawMessage `json:"ucp-auth-worker-service.txt"`
	UcpAuthWorkerTasksTxt         json.RawMessage `json:"ucp-auth-worker-tasks.txt"`
	UcpClusterAgentServiceTxt     json.RawMessage `json:"ucp-cluster-agent-service.txt"`
	UcpClusterAgentTasksTxt       json.RawMessage `json:"ucp-cluster-agent-tasks.txt"`
	UcpControllerConfigJson       json.RawMessage `json:"ucp-controller-config.json"`
	UcpInstanceIdTxt              json.RawMessage `json:"ucp-instance-id.txt"`
	UcpInterlockConfigTasksTxt    json.RawMessage `json:"ucp-interlock-config-tasks.txt"`
	UcpInterlockConfigTxt         json.RawMessage `json:"ucp-interlock-config.txt"`
	UcpInterlockExtensionTasksTxt json.RawMessage `json:"ucp-interlock-extension-tasks.txt"`
	UcpInterlockExtensionTxt      json.RawMessage `json:"ucp-interlock-extension.txt"`
	UcpInterlockNetworkTxt        json.RawMessage `json:"ucp-interlock-network.txt"`
	UcpInterlockProxyTasksTxt     json.RawMessage `json:"ucp-interlock-proxy-tasks.txt"`
	UcpInterlockProxyTxt          json.RawMessage `json:"ucp-interlock-proxy.txt"`
	UcpInterlockServiceTxt        json.RawMessage `json:"ucp-interlock-service.txt"`
	UcpInterlockTasksTxt          json.RawMessage `json:"ucp-interlock-tasks.txt"`
	UcpLicenseIdTxt               []string        `json:"ucp-license-id.txt"`
	UcpManagerAgentServiceTxt     json.RawMessage `json:"ucp-manager-agent-service.txt"`
	UcpManagerAgentTasksTxt       json.RawMessage `json:"ucp-manager-agent-tasks.txt"`
	UcpNodesTxt                   json.RawMessage `json:"ucp-nodes.txt"`
	UcpTelemetryJson              json.RawMessage `json:"ucp-telemetry.json"`
	UcpWorkerAgentWinXServiceTxt  json.RawMessage `json:"ucp-worker-agent-win-x-service.txt"`
	UcpWorkerAgentWinXTasksTxt    json.RawMessage `json:"ucp-worker-agent-win-x-tasks.txt"`
	UcpWorkerAgentXServiceTxt     json.RawMessage `json:"ucp-worker-agent-x-service.txt"`
	UcpWorkerAgentXTasksTxt       json.RawMessage `json:"ucp-worker-agent-x-tasks.txt"`
}
type DsinfoSlashDsinfoDotJson struct {
	DsinfoContents json.RawMessage `json:"dsinfo/dsinfo.json"`
}
type NestedDsinfo struct {
	Calico              json.RawMessage `json:"calico"`
	ClusterInfo         json.RawMessage `json:"cluster_info"`
	ContainerInfo       json.RawMessage `json:"container_info"`
	DaemonStackTrace    json.RawMessage `json:"daemon_stack_trace"`
	DciDeployment       json.RawMessage `json:"dci_deployment"`
	DeadContainerMounts json.RawMessage `json:"dead_container_mounts"`
	Df                  json.RawMessage `json:"df"`
	Dmesg               json.RawMessage `json:"dmesg"`
	Dmidecode           json.RawMessage `json:"dmidecode"`
	DockerDaemonJson    json.RawMessage `json:"docker_daemon_json"`
	DockerImages        json.RawMessage `json:"docker_images"`
	DockerInfo          json.RawMessage `json:"docker_info"`
	DockerPidLimits     json.RawMessage `json:"docker_pid_limits"`
	DockerStats         json.RawMessage `json:"docker_stats"`
	DockerVersion       json.RawMessage `json:"docker_version"`
	Dtr                 json.RawMessage `json:"dtr"`
	FileMax             json.RawMessage `json:"file_max"`
	Hostname            json.RawMessage `json:"hostname"`
	Ifconfig            json.RawMessage `json:"ifconfig"`
	Iostat              json.RawMessage `json:"iostat"`
	JournalctlDaemon    json.RawMessage `json:"journalctl_daemon"`
	JournalctlKernel    json.RawMessage `json:"journalctl_kernel"`
	KernelConfig        json.RawMessage `json:"kernel_config"`
	KernelVersion       json.RawMessage `json:"kernel_version"`
	Lscpu               json.RawMessage `json:"lscpu"`
	Meminfo             json.RawMessage `json:"meminfo"`
	Mount               json.RawMessage `json:"mount"`
	Netstat             json.RawMessage `json:"netstat"`
	Networking          json.RawMessage `json:"networking"`
	PsAuxGrepDocker     json.RawMessage `json:"ps_aux_grep_docker"`
	RethinkStatus       json.RawMessage `json:"rethink_status"`
	RunningCgroup       json.RawMessage `json:"running_cgroup"`
	Sestatus            json.RawMessage `json:"sestatus"`
	ShimLogs            json.RawMessage `json:"shim_logs"`
	Ssd                 json.RawMessage `json:"ssd"`
	SystemCgroups       json.RawMessage `json:"system_cgroups"`
	SystemRelease       json.RawMessage `json:"system_release"`
	UcpControllerDiag   json.RawMessage `json:"ucp_controller_diag"`
	Uptime              json.RawMessage `json:"uptime"`
	Vmstat              json.RawMessage `json:"vmstat"`
}
