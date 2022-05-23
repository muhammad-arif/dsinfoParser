## Purpose
The purpose of this package/library/types is to interpret the contents of `dsinfo.json` file. the `dsinfo.json` file consist the complete structure and contents of the support bundle. So we can utilze this file to automate our investigation process. But the contents of `dsinfo.json` is not directly usable because it has different files of different format. This package attempts to interpret the various contents of the `dsinfo.json` file with proper structure.

## Use case
This package could be utilized while creating an API server or a command line tool to parse/interpret `dsinfo.json` file.

## How to

At first, we need to read the file and unmarshal it with the core structure. The core structure won't unmarshal the fileds  with proper sturcture but it will keep most of the fileds as byte array. But when we need to interpret the contents of the specific filed we can unmarshal them with appropriate structure. We have choose this path to reduce memory footprint, as sometimes `dsinfo.json` file could be big as 10GB.

So initially, we can do something like the following to unmarshal the core structure,
```go
var file = "dsinfo.json"
dsinfoJson, err := ioutil.ReadFile(file)
if err != nil {
fmt.Errorf("cannot Read File %s", err)
}
var core CoreDsinfo
err = json.Unmarshal(dsinfoJson, &core)
if err != nil {
fmt.Errorf("Cannot unmarshal %s", err)
}
```
After that we can see the contents of the `core`'s filed with the following ways.
### To collect the info from `ucp-nodes.txt`
To collect at Json format , easily parsable by jq
```go
nodeInspect := NewNodeParser().GetAllJSON(&core.UcpNodesTxt)
fmt.Println(nodeInspect)
```
To collect the data with proper structure
```go
nodeInspect := NewNodeParser().GetAll(&core.UcpNodesTxt)
fmt.Println(nodeInspect[0].Description.Engine.EngineVersion)
```
You can also iterate all nodes and collect the info. For details, go to ucpnodestxt.go->type UCPNodesTxtType
```go
for k, _ := range nodeInspect {
	fmt.Printf("Hostname: %s\n\tEngine Version: %s\n\tRole: %s\n", nodeInspect[k].Description.Hostname, nodeInspect[k].Description.Engine.EngineVersion, nodeInspect[k].Spec.Role)
}
```
### To collect the info from `ucp-telemetry.json`
Here data is already in json format.

To collect at json format
```go
fmt.Println(string(core.UcpTelemetryJson))
```
To collect the data with proper structure. For details go to ucp_controller_config.go
```go
data := NewUCPTelemetry().GetAll(&core.UcpTelemetryJson)
fmt.Printf("Cluster ID: %s\n\tIs Licensed: %v\n\tOS: %s\n\tKernel Version:%s\n", data.ClusterId, data.IsLicensed, data.Os, data.KernelVersion)
fmt.Println(data.Kubernetes.ManagerReservedResources)
```

### To collect the info from `ucp-*-tasks.txt`
- Use `UcpAuthApiTasksTxt` for `ucp-auth-api-tasks.txt`
- Use `UcpAuthWorkerTasksTxt` for `ucp-auth-worker-tasks.txt`
- Use `UcpClusterAgentTasksTxt` for `ucp-cluster-agent-tasks.txt`
- Use `UcpInterlockConfigTasksTxt` for `ucp-interlock-config-tasks.txt`
- Use `UcpInterlockExtensionTasksTxt` for `ucp-interlock-extension-tasks.txt`
- Use `UcpInterlockProxyTasksTxt` for `ucp-interlock-proxy-tasks.txt`
- Use `UcpInterlockTasksTxt` for `ucp-interlock-tasks.txt`
- Use `UcpManagerAgentTasksTxt` for `ucp-manager-agent-tasks.txt`
- Use `UcpWorkerAgentWinXTasksTxt` for `ucp-worker-agent-win-x-tasks.tx`
- Use `UcpWorkerAgentXTasksTxt` for `ucp-worker-agent-x-service.txt`

To collect at Json format , easily parsable by jq
```go
fmt.Println(NewTasksParser().GetAllJSON(&core.UcpAuthApiTasksTxt))
fmt.Println(NewTasksParser().GetAllJSON(&core.UcpAuthWorkerTasksTxt))
fmt.Println(NewTasksParser().GetAllJSON(&core.UcpClusterAgentTasksTxt))
fmt.Println(NewTasksParser().GetAllJSON(&core.UcpInterlockConfigTasksTxt))
fmt.Println(NewTasksParser().GetAllJSON(&core.UcpInterlockExtensionTasksTxt))
fmt.Println(NewTasksParser().GetAllJSON(&core.UcpInterlockProxyTasksTxt))
fmt.Println(NewTasksParser().GetAllJSON(&core.UcpInterlockTasksTxt))
fmt.Println(NewTasksParser().GetAllJSON(&core.UcpManagerAgentTasksTxt))
fmt.Println(NewTasksParser().GetAllJSON(&core.UcpWorkerAgentWinXTasksTxt))
fmt.Println(NewTasksParser().GetAllJSON(&core.UcpWorkerAgentXTasksTxt))
```
To collect the data with proper structure
```go
data := NewTasksParser().GetAll(&core.UcpAuthApiTasksTxt)
fmt.Println(data[0].ServiceID)
```
You can also iterate all fields and collect the info. For details, go to `task_func.go`
```go
for k, _ := range data {
    fmt.Printf("Task ID: %s\nImage: %s\n", data[k].ID, data[k].Spec.ContainerSpec.Image)
}
```
### To collect the info from `ucp-*-service.txt`
- Use `UcpAuthApiServiceTxt` for `ucp-auth-api-service.txt`
- Use `UcpAuthWorkerServiceTxt` for `ucp-auth-worker-service.txt`
- Use `UcpClusterAgentServiceTxt` for `ucp-cluster-agent-service.txt`
- Use `UcpInterlockServiceTxt` for `ucp-interlock-service.txt`
- Use `UcpManagerAgentServiceTxt` for `ucp-manager-agent-service.txt`
- Use `UcpWorkerAgentWinXServiceTxt` for `ucp-worker-agent-win-x-service.txt`
- Use `UcpWorkerAgentXServiceTxt` for `ucp-worker-agent-x-service.txt`

To collect at json format
```go
fmt.Println(string(core.UcpAuthApiServiceTxt))
fmt.Println(string(core.UcpAuthWorkerServiceTxt))
fmt.Println(string(core.UcpClusterAgentServiceTxt))
fmt.Println(string(core.UcpInterlockServiceTxt))
fmt.Println(string(core.UcpManagerAgentServiceTxt))
fmt.Println(string(core.UcpWorkerAgentWinXServiceTxt))
fmt.Println(string(core.UcpWorkerAgentXServiceTxt))
```
To collect the data with proper structure. For details go to service_func.go
```go
data := NewServiceParser().GetAll(&core.UcpAuthApiServiceTxt)
fmt.Printf("Name: %s\nID: %s\nImage: %s\n", data.Spec.Name, data.ID, data.Spec.TaskTemplate.ContainerSpec.Image)
```
### To collect the info from `mirantis.lic`
Here data is already in json format

To collect at json format
```go
fmt.Println(string(core.MirantisLic))
```
To collect the data with proper structure. For details go to `mirantis_lic.go`
```go
data := NewLicenseType().GetAll(&core.MirantisLic)
fmt.Printf("Max Engine: %d\nExpiration Date: %s\n", data.Details.MaxEngines, data.Details.Expiration)
```
### To collect the info from `ucp-controller-config.json`
Here data is already in json format

To collect at json format
```go
fmt.Println(string(core.UcpControllerConfigJson))
```
To collect the data with proper structure. For details go to ucp_controller_config.go
```go
data := NewUCPControllerConfigType().GetAll(&core.UcpControllerConfigJson)
fmt.Println(data.IngressController.Enabled)
fmt.Println(data.Kubernetes.ManagerReservedResources)
```
### To collect the info from txt type files
Following files are not in specific format,
- Use `UcpLicenseIdTxt` for `ucp-license-id.txt`
- Use `TasksCountTxt` for `tasks-count.txt`
- Use `IngressDescribeAllTxt` for `ingress-describe-all.txt`
- Use `IngressDescribeIngressTxt` for `ingress-describe-ingress.txt`
- Use `IngressGetAllTxt` for `ingress-get-all.txt`
- Use `IngressGetIngressTxt` for `ingress-get-ingress.txt`
- Use `IngressNginxConfigTxt` for `ingress-nginx-config.txt`
- Use `IngressNginxControllerTxt` for `ingress-nginx-controller.txt`
- Use `KubeDescribeAllTxt` for `kube-describe-all.txt`
- Use `KubeDescribeNodesTxt` for `kube-describe-nodes.txt`

To just print the string conents,
```go
fmt.Println(core.UcpLicenseIdTxt)
fmt.Println(core.TasksCountTxt)
fmt.Println(core.IngressDescribeAllTxt)
fmt.Println(core.IngressDescribeIngressTxt)
fmt.Println(core.IngressGetAllTxt)
fmt.Println(core.IngressGetIngressTxt)
fmt.Println(core.IngressNginxControllerTxt)
fmt.Println(core.IngressNginxControllerTxt)
fmt.Println(core.KubeDescribeNodesTxt)
fmt.Println(core.KubeDescribeAllTxt)
```