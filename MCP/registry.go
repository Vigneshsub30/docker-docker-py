package main

import (
	"github.com/docker-engine-api/mcp-server/config"
	"github.com/docker-engine-api/mcp-server/models"
	tools_container "github.com/docker-engine-api/mcp-server/tools/container"
	tools_secret "github.com/docker-engine-api/mcp-server/tools/secret"
	tools_image "github.com/docker-engine-api/mcp-server/tools/image"
	tools_system "github.com/docker-engine-api/mcp-server/tools/system"
	tools_service "github.com/docker-engine-api/mcp-server/tools/service"
	tools_swarm "github.com/docker-engine-api/mcp-server/tools/swarm"
	tools_task "github.com/docker-engine-api/mcp-server/tools/task"
	tools_volume "github.com/docker-engine-api/mcp-server/tools/volume"
	tools_node "github.com/docker-engine-api/mcp-server/tools/node"
	tools_config "github.com/docker-engine-api/mcp-server/tools/config"
	tools_plugin "github.com/docker-engine-api/mcp-server/tools/plugin"
	tools_exec "github.com/docker-engine-api/mcp-server/tools/exec"
	tools_distribution "github.com/docker-engine-api/mcp-server/tools/distribution"
	tools_network "github.com/docker-engine-api/mcp-server/tools/network"
	tools_tasks "github.com/docker-engine-api/mcp-server/tools/tasks"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_container.CreateContainerwaitTool(cfg),
		tools_secret.CreateSecretcreateTool(cfg),
		tools_image.CreateImagepruneTool(cfg),
		tools_system.CreateSysteminfoTool(cfg),
		tools_service.CreateServicelistTool(cfg),
		tools_swarm.CreateSwarmupdateTool(cfg),
		tools_swarm.CreateSwarmleaveTool(cfg),
		tools_task.CreateTasklistTool(cfg),
		tools_service.CreateServicecreateTool(cfg),
		tools_volume.CreateVolumepruneTool(cfg),
		tools_node.CreateNodelistTool(cfg),
		tools_image.CreateImagedeleteTool(cfg),
		tools_container.CreateContainerlistTool(cfg),
		tools_container.CreateContainerstartTool(cfg),
		tools_container.CreateContainerinspectTool(cfg),
		tools_service.CreateServiceupdateTool(cfg),
		tools_image.CreateImagelistTool(cfg),
		tools_image.CreateImageinspectTool(cfg),
		tools_config.CreateConfigcreateTool(cfg),
		tools_secret.CreateSecretupdateTool(cfg),
		tools_config.CreateConfigupdateTool(cfg),
		tools_volume.CreateVolumelistTool(cfg),
		tools_service.CreateServiceinspectTool(cfg),
		tools_service.CreateServicedeleteTool(cfg),
		tools_swarm.CreateSwarmunlockkeyTool(cfg),
		tools_container.CreateContainerunpauseTool(cfg),
		tools_plugin.CreatePluginupgradeTool(cfg),
		tools_image.CreateImagetagTool(cfg),
		tools_swarm.CreateSwarmunlockTool(cfg),
		tools_container.CreateContainerstatsTool(cfg),
		tools_secret.CreateSecretlistTool(cfg),
		tools_service.CreateServicelogsTool(cfg),
		tools_task.CreateTaskinspectTool(cfg),
		tools_volume.CreateVolumedeleteTool(cfg),
		tools_volume.CreateVolumeinspectTool(cfg),
		tools_container.CreateContainerresizeTool(cfg),
		tools_container.CreateContainerupdateTool(cfg),
		tools_plugin.CreatePluginsetTool(cfg),
		tools_config.CreateConfiglistTool(cfg),
		tools_plugin.CreatePluginenableTool(cfg),
		tools_system.CreateSystempingTool(cfg),
		tools_exec.CreateContainerexecTool(cfg),
		tools_distribution.CreateDistributioninspectTool(cfg),
		tools_plugin.CreatePlugininspectTool(cfg),
		tools_container.CreateContainerrenameTool(cfg),
		tools_container.CreateContainerstopTool(cfg),
		tools_container.CreateContainerarchiveinfoTool(cfg),
		tools_network.CreateNetworkcreateTool(cfg),
		tools_network.CreateNetworkpruneTool(cfg),
		tools_secret.CreateSecretdeleteTool(cfg),
		tools_secret.CreateSecretinspectTool(cfg),
		tools_image.CreateImagepushTool(cfg),
		tools_plugin.CreatePlugindisableTool(cfg),
		tools_image.CreateImagesearchTool(cfg),
		tools_system.CreateSystemversionTool(cfg),
		tools_swarm.CreateSwarminspectTool(cfg),
		tools_network.CreateNetworkdeleteTool(cfg),
		tools_network.CreateNetworkinspectTool(cfg),
		tools_container.CreateContainercreateTool(cfg),
		tools_plugin.CreateGetpluginprivilegesTool(cfg),
		tools_exec.CreateExecinspectTool(cfg),
		tools_system.CreateSystemdatausageTool(cfg),
		tools_tasks.CreateTasklogsTool(cfg),
		tools_image.CreateImagecreateTool(cfg),
		tools_swarm.CreateSwarmjoinTool(cfg),
		tools_node.CreateNodeupdateTool(cfg),
		tools_container.CreateContainerrestartTool(cfg),
		tools_network.CreateNetworkdisconnectTool(cfg),
		tools_image.CreateImagecommitTool(cfg),
		tools_exec.CreateExecresizeTool(cfg),
		tools_container.CreateContainerlogsTool(cfg),
		tools_system.CreateSystemauthTool(cfg),
		tools_swarm.CreateSwarminitTool(cfg),
		tools_volume.CreateVolumecreateTool(cfg),
		tools_container.CreateContainerchangesTool(cfg),
		tools_network.CreateNetworklistTool(cfg),
		tools_container.CreateContainerpauseTool(cfg),
		tools_image.CreateBuildpruneTool(cfg),
		tools_container.CreateContainerattachwebsocketTool(cfg),
		tools_container.CreateContainerpruneTool(cfg),
		tools_container.CreateContainerkillTool(cfg),
		tools_image.CreateImagehistoryTool(cfg),
		tools_plugin.CreatePluginlistTool(cfg),
		tools_container.CreateContainertopTool(cfg),
		tools_plugin.CreatePluginpullTool(cfg),
		tools_system.CreateSystemeventsTool(cfg),
		tools_plugin.CreatePlugindeleteTool(cfg),
		tools_container.CreateContainerdeleteTool(cfg),
		tools_plugin.CreatePluginpushTool(cfg),
		tools_node.CreateNodedeleteTool(cfg),
		tools_node.CreateNodeinspectTool(cfg),
		tools_config.CreateConfigdeleteTool(cfg),
		tools_config.CreateConfiginspectTool(cfg),
	}
}
