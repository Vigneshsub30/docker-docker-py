package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ClusterInfo represents the ClusterInfo schema from the OpenAPI specification
type ClusterInfo struct {
	Createdat string `json:"CreatedAt,omitempty"` // Date and time at which the swarm was initialised in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
	Id string `json:"ID,omitempty"` // The ID of the swarm.
	Rootrotationinprogress bool `json:"RootRotationInProgress,omitempty"` // Whether there is currently a root CA rotation in progress for the swarm
	Spec SwarmSpec `json:"Spec,omitempty"` // User modifiable swarm configuration.
	Tlsinfo TLSInfo `json:"TLSInfo,omitempty"` // Information about the issuer of leaf TLS certificates and the trusted root CA certificate
	Updatedat string `json:"UpdatedAt,omitempty"` // Date and time at which the swarm was last updated in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
	Version ObjectVersion `json:"Version,omitempty"` // The version number of the object such as node, service, etc. This is needed to avoid conflicting writes. The client must send the version number along with the modified specification when updating these objects. This approach ensures safe concurrency and determinism in that the change on the object may not be applied if the version number has changed from the last read. In other words, if two update requests specify the same base version, only one of the requests can succeed. As a result, two separate update requests that happen at the same time will not unintentionally overwrite each other.
}

// ErrorDetail represents the ErrorDetail schema from the OpenAPI specification
type ErrorDetail struct {
	Code int `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// TLSInfo represents the TLSInfo schema from the OpenAPI specification
type TLSInfo struct {
	Certissuersubject string `json:"CertIssuerSubject,omitempty"` // The base64-url-safe-encoded raw subject bytes of the issuer
	Trustroot string `json:"TrustRoot,omitempty"` // The root CA certificate(s) that are used to validate leaf TLS certificates
	Certissuerpublickey string `json:"CertIssuerPublicKey,omitempty"` // The base64-url-safe-encoded raw public key bytes of the issuer
}

// PeerNode represents the PeerNode schema from the OpenAPI specification
type PeerNode struct {
	Addr string `json:"Addr,omitempty"` // IP address and ports at which this node can be reached.
	Nodeid string `json:"NodeID,omitempty"` // Unique identifier of for this node in the swarm.
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	Message string `json:"message"` // The error message.
}

// PortBinding represents the PortBinding schema from the OpenAPI specification
type PortBinding struct {
	Hostip string `json:"HostIp,omitempty"` // Host IP address that the container's port is mapped to.
	Hostport string `json:"HostPort,omitempty"` // Host port number that the container's port is mapped to.
}

// SecretSpec represents the SecretSpec schema from the OpenAPI specification
type SecretSpec struct {
	Driver Driver `json:"Driver,omitempty"` // Driver represents a driver (network, logging, secrets).
	Labels map[string]interface{} `json:"Labels,omitempty"` // User-defined key/value metadata.
	Name string `json:"Name,omitempty"` // User-defined name of the secret.
	Data string `json:"Data,omitempty"` // Base64-url-safe-encoded ([RFC 4648](https://tools.ietf.org/html/rfc4648#section-3.2)) data to store as secret. This field is only used to _create_ a secret, and is not returned by other endpoints.
}

// PluginInterfaceType represents the PluginInterfaceType schema from the OpenAPI specification
type PluginInterfaceType struct {
	Capability string `json:"Capability"`
	Prefix string `json:"Prefix"`
	Version string `json:"Version"`
}

// ProcessConfig represents the ProcessConfig schema from the OpenAPI specification
type ProcessConfig struct {
	User string `json:"user,omitempty"`
	Arguments []string `json:"arguments,omitempty"`
	Entrypoint string `json:"entrypoint,omitempty"`
	Privileged bool `json:"privileged,omitempty"`
	Tty bool `json:"tty,omitempty"`
}

// Mount represents the Mount schema from the OpenAPI specification
type Mount struct {
	Bindoptions map[string]interface{} `json:"BindOptions,omitempty"` // Optional configuration for the `bind` type.
	Consistency string `json:"Consistency,omitempty"` // The consistency requirement for the mount: `default`, `consistent`, `cached`, or `delegated`.
	Readonly bool `json:"ReadOnly,omitempty"` // Whether the mount should be read-only.
	Source string `json:"Source,omitempty"` // Mount source (e.g. a volume name, a host path).
	Target string `json:"Target,omitempty"` // Container path.
	Tmpfsoptions map[string]interface{} `json:"TmpfsOptions,omitempty"` // Optional configuration for the `tmpfs` type.
	TypeField string `json:"Type,omitempty"` // The mount type. Available types: - `bind` Mounts a file or directory from the host into the container. Must exist prior to creating the container. - `volume` Creates a volume with the given name and options (or uses a pre-existing volume with the same name and options). These are **not** removed when the container is removed. - `tmpfs` Create a tmpfs with the given options. The mount source cannot be specified for tmpfs.
	Volumeoptions map[string]interface{} `json:"VolumeOptions,omitempty"` // Optional configuration for the `volume` type.
}

// Resources represents the Resources schema from the OpenAPI specification
type Resources struct {
	Diskquota int64 `json:"DiskQuota,omitempty"` // Disk limit (in bytes).
	Blkiodevicewritebps []ThrottleDevice `json:"BlkioDeviceWriteBps,omitempty"` // Limit write rate (bytes per second) to a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	Cpushares int `json:"CpuShares,omitempty"` // An integer value representing this container's relative CPU weight versus other containers.
	Devices []DeviceMapping `json:"Devices,omitempty"` // A list of devices to add to the container.
	Cpupercent int64 `json:"CpuPercent,omitempty"` // The usable percentage of the available CPUs (Windows only). On Windows Server containers, the processor resource controls are mutually exclusive. The order of precedence is `CPUCount` first, then `CPUShares`, and `CPUPercent` last.
	Cpurealtimeruntime int64 `json:"CpuRealtimeRuntime,omitempty"` // The length of a CPU real-time runtime in microseconds. Set to 0 to allocate no time allocated to real-time tasks.
	Pidslimit int64 `json:"PidsLimit,omitempty"` // Tune a container's pids limit. Set -1 for unlimited.
	Cpuperiod int64 `json:"CpuPeriod,omitempty"` // The length of a CPU period in microseconds.
	Blkiodevicewriteiops []ThrottleDevice `json:"BlkioDeviceWriteIOps,omitempty"` // Limit write rate (IO per second) to a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	Cpusetmems string `json:"CpusetMems,omitempty"` // Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.
	Blkioweight int `json:"BlkioWeight,omitempty"` // Block IO weight (relative weight).
	Kernelmemory int64 `json:"KernelMemory,omitempty"` // Kernel memory limit in bytes.
	Cpurealtimeperiod int64 `json:"CpuRealtimePeriod,omitempty"` // The length of a CPU real-time period in microseconds. Set to 0 to allocate no time allocated to real-time tasks.
	Cpuquota int64 `json:"CpuQuota,omitempty"` // Microseconds of CPU time that the container can get in a CPU period.
	Cgroupparent string `json:"CgroupParent,omitempty"` // Path to `cgroups` under which the container's `cgroup` is created. If the path is not absolute, the path is considered to be relative to the `cgroups` path of the init process. Cgroups are created if they do not already exist.
	Blkiodevicereadbps []ThrottleDevice `json:"BlkioDeviceReadBps,omitempty"` // Limit read rate (bytes per second) from a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	Memory int `json:"Memory,omitempty"` // Memory limit in bytes.
	Memoryreservation int64 `json:"MemoryReservation,omitempty"` // Memory soft limit in bytes.
	Blkioweightdevice []map[string]interface{} `json:"BlkioWeightDevice,omitempty"` // Block IO weight (relative device weight) in the form `[{"Path": "device_path", "Weight": weight}]`.
	Ulimits []map[string]interface{} `json:"Ulimits,omitempty"` // A list of resource limits to set in the container. For example: `{"Name": "nofile", "Soft": 1024, "Hard": 2048}`"
	Devicecgrouprules []string `json:"DeviceCgroupRules,omitempty"` // a list of cgroup rules to apply to the container
	Memoryswap int64 `json:"MemorySwap,omitempty"` // Total memory limit (memory + swap). Set as `-1` to enable unlimited swap.
	Cpusetcpus string `json:"CpusetCpus,omitempty"` // CPUs in which to allow execution (e.g., `0-3`, `0,1`)
	Iomaximumbandwidth int64 `json:"IOMaximumBandwidth,omitempty"` // Maximum IO in bytes per second for the container system drive (Windows only)
	Nanocpus int64 `json:"NanoCPUs,omitempty"` // CPU quota in units of 10<sup>-9</sup> CPUs.
	Blkiodevicereadiops []ThrottleDevice `json:"BlkioDeviceReadIOps,omitempty"` // Limit read rate (IO per second) from a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	Memoryswappiness int64 `json:"MemorySwappiness,omitempty"` // Tune a container's memory swappiness behavior. Accepts an integer between 0 and 100.
	Oomkilldisable bool `json:"OomKillDisable,omitempty"` // Disable OOM Killer for the container.
	Iomaximumiops int64 `json:"IOMaximumIOps,omitempty"` // Maximum IOps for the container system drive (Windows only)
	Cpucount int64 `json:"CpuCount,omitempty"` // The number of usable CPUs (Windows only). On Windows Server containers, the processor resource controls are mutually exclusive. The order of precedence is `CPUCount` first, then `CPUShares`, and `CPUPercent` last.
}

// ContainerConfig represents the ContainerConfig schema from the OpenAPI specification
type ContainerConfig struct {
	Entrypoint string `json:"Entrypoint,omitempty"` // The entry point for the container as a string or an array of strings. If the array consists of exactly one empty string (`[""]`) then the entry point is reset to system default (i.e., the entry point used by docker when there is no `ENTRYPOINT` instruction in the `Dockerfile`).
	Macaddress string `json:"MacAddress,omitempty"` // MAC address of the container.
	Networkdisabled bool `json:"NetworkDisabled,omitempty"` // Disable networking for the container.
	Exposedports map[string]interface{} `json:"ExposedPorts,omitempty"` // An object mapping ports to an empty object in the form: `{"<port>/<tcp|udp>": {}}`
	Shell []string `json:"Shell,omitempty"` // Shell for when `RUN`, `CMD`, and `ENTRYPOINT` uses a shell.
	Workingdir string `json:"WorkingDir,omitempty"` // The working directory for commands to run in.
	Hostname string `json:"Hostname,omitempty"` // The hostname to use for the container, as a valid RFC 1123 hostname.
	Stdinonce bool `json:"StdinOnce,omitempty"` // Close `stdin` after one attached client disconnects
	Labels map[string]interface{} `json:"Labels,omitempty"` // User-defined key/value metadata.
	Attachstderr bool `json:"AttachStderr,omitempty"` // Whether to attach to `stderr`.
	Attachstdin bool `json:"AttachStdin,omitempty"` // Whether to attach to `stdin`.
	Image string `json:"Image,omitempty"` // The name of the image to use when creating the container
	Stopsignal string `json:"StopSignal,omitempty"` // Signal to stop a container as a string or unsigned integer.
	User string `json:"User,omitempty"` // The user that commands are run as inside the container.
	Onbuild []string `json:"OnBuild,omitempty"` // `ONBUILD` metadata that were defined in the image's `Dockerfile`.
	Volumes map[string]interface{} `json:"Volumes,omitempty"` // An object mapping mount point paths inside the container to empty objects.
	Healthcheck HealthConfig `json:"Healthcheck,omitempty"` // A test to perform to check that the container is healthy.
	Tty bool `json:"Tty,omitempty"` // Attach standard streams to a TTY, including `stdin` if it is not closed.
	Domainname string `json:"Domainname,omitempty"` // The domain name to use for the container.
	Argsescaped bool `json:"ArgsEscaped,omitempty"` // Command is already escaped (Windows only)
	Cmd string `json:"Cmd,omitempty"` // Command to run specified as a string or an array of strings.
	Openstdin bool `json:"OpenStdin,omitempty"` // Open `stdin`
	Env []string `json:"Env,omitempty"` // A list of environment variables to set inside the container in the form `["VAR=value", ...]`. A variable without `=` is removed from the environment, rather than to have an empty value.
	Attachstdout bool `json:"AttachStdout,omitempty"` // Whether to attach to `stdout`.
	Stoptimeout int `json:"StopTimeout,omitempty"` // Timeout to stop a container in seconds.
}

// Config represents the Config schema from the OpenAPI specification
type Config struct {
	Updatedat string `json:"UpdatedAt,omitempty"`
	Version ObjectVersion `json:"Version,omitempty"` // The version number of the object such as node, service, etc. This is needed to avoid conflicting writes. The client must send the version number along with the modified specification when updating these objects. This approach ensures safe concurrency and determinism in that the change on the object may not be applied if the version number has changed from the last read. In other words, if two update requests specify the same base version, only one of the requests can succeed. As a result, two separate update requests that happen at the same time will not unintentionally overwrite each other.
	Createdat string `json:"CreatedAt,omitempty"`
	Id string `json:"ID,omitempty"`
	Spec ConfigSpec `json:"Spec,omitempty"`
}

// ImageSummary represents the ImageSummary schema from the OpenAPI specification
type ImageSummary struct {
	Repodigests []string `json:"RepoDigests"`
	Repotags []string `json:"RepoTags"`
	Containers int `json:"Containers"`
	Id string `json:"Id"`
	Sharedsize int `json:"SharedSize"`
	Created int `json:"Created"`
	Labels map[string]interface{} `json:"Labels"`
	Parentid string `json:"ParentId"`
	Virtualsize int `json:"VirtualSize"`
	Size int `json:"Size"`
}

// EngineDescription represents the EngineDescription schema from the OpenAPI specification
type EngineDescription struct {
	Engineversion string `json:"EngineVersion,omitempty"`
	Labels map[string]interface{} `json:"Labels,omitempty"`
	Plugins []map[string]interface{} `json:"Plugins,omitempty"`
}

// PushImageInfo represents the PushImageInfo schema from the OpenAPI specification
type PushImageInfo struct {
	Progress string `json:"progress,omitempty"`
	Progressdetail ProgressDetail `json:"progressDetail,omitempty"`
	Status string `json:"status,omitempty"`
	ErrorField string `json:"error,omitempty"`
}

// HealthConfig represents the HealthConfig schema from the OpenAPI specification
type HealthConfig struct {
	Startperiod int `json:"StartPeriod,omitempty"` // Start period for the container to initialize before starting health-retries countdown in nanoseconds. It should be 0 or at least 1000000 (1 ms). 0 means inherit.
	Test []string `json:"Test,omitempty"` // The test to perform. Possible values are: - `[]` inherit healthcheck from image or parent image - `["NONE"]` disable healthcheck - `["CMD", args...]` exec arguments directly - `["CMD-SHELL", command]` run command with system's default shell
	Timeout int `json:"Timeout,omitempty"` // The time to wait before considering the check to have hung. It should be 0 or at least 1000000 (1 ms). 0 means inherit.
	Interval int `json:"Interval,omitempty"` // The time to wait between checks in nanoseconds. It should be 0 or at least 1000000 (1 ms). 0 means inherit.
	Retries int `json:"Retries,omitempty"` // The number of consecutive failures needed to consider a container as unhealthy. 0 means inherit.
}

// EndpointSpec represents the EndpointSpec schema from the OpenAPI specification
type EndpointSpec struct {
	Mode string `json:"Mode,omitempty"` // The mode of resolution to use for internal load balancing between tasks.
	Ports []EndpointPortConfig `json:"Ports,omitempty"` // List of exposed ports that this service is accessible on from the outside. Ports can only be provided if `vip` resolution mode is used.
}

// Platform represents the Platform schema from the OpenAPI specification
type Platform struct {
	Architecture string `json:"Architecture,omitempty"` // Architecture represents the hardware architecture (for example, `x86_64`).
	Os string `json:"OS,omitempty"` // OS represents the Operating System (for example, `linux` or `windows`).
}

// Network represents the Network schema from the OpenAPI specification
type Network struct {
	Options map[string]interface{} `json:"Options,omitempty"`
	Created string `json:"Created,omitempty"`
	Labels map[string]interface{} `json:"Labels,omitempty"`
	Scope string `json:"Scope,omitempty"`
	Driver string `json:"Driver,omitempty"`
	Ipam IPAM `json:"IPAM,omitempty"`
	Id string `json:"Id,omitempty"`
	Internal bool `json:"Internal,omitempty"`
	Attachable bool `json:"Attachable,omitempty"`
	Containers map[string]interface{} `json:"Containers,omitempty"`
	Enableipv6 bool `json:"EnableIPv6,omitempty"`
	Ingress bool `json:"Ingress,omitempty"`
	Name string `json:"Name,omitempty"`
}

// Plugin represents the Plugin schema from the OpenAPI specification
type Plugin struct {
	Id string `json:"Id,omitempty"`
	Name string `json:"Name"`
	Pluginreference string `json:"PluginReference,omitempty"` // plugin remote reference used to push/pull the plugin
	Settings map[string]interface{} `json:"Settings"` // Settings that can be modified by users.
	Config map[string]interface{} `json:"Config"` // The config of a plugin.
	Enabled bool `json:"Enabled"` // True if the plugin is running. False if the plugin is not running, only installed.
}

// IPAM represents the IPAM schema from the OpenAPI specification
type IPAM struct {
	Options []map[string]interface{} `json:"Options,omitempty"` // Driver-specific options, specified as a map.
	Config []map[string]interface{} `json:"Config,omitempty"` // List of IPAM configuration options, specified as a map: `{"Subnet": <CIDR>, "IPRange": <CIDR>, "Gateway": <IP address>, "AuxAddress": <device_name:IP address>}`
	Driver string `json:"Driver,omitempty"` // Name of the IPAM driver to use.
}

// ConfigSpec represents the ConfigSpec schema from the OpenAPI specification
type ConfigSpec struct {
	Name string `json:"Name,omitempty"` // User-defined name of the config.
	Data string `json:"Data,omitempty"` // Base64-url-safe-encoded ([RFC 4648](https://tools.ietf.org/html/rfc4648#section-3.2)) config data.
	Labels map[string]interface{} `json:"Labels,omitempty"` // User-defined key/value metadata.
}

// NetworkContainer represents the NetworkContainer schema from the OpenAPI specification
type NetworkContainer struct {
	Name string `json:"Name,omitempty"`
	Endpointid string `json:"EndpointID,omitempty"`
	Ipv4address string `json:"IPv4Address,omitempty"`
	Ipv6address string `json:"IPv6Address,omitempty"`
	Macaddress string `json:"MacAddress,omitempty"`
}

// Swarm represents the Swarm schema from the OpenAPI specification
type Swarm struct {
	Createdat string `json:"CreatedAt,omitempty"` // Date and time at which the swarm was initialised in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
	Id string `json:"ID,omitempty"` // The ID of the swarm.
	Rootrotationinprogress bool `json:"RootRotationInProgress,omitempty"` // Whether there is currently a root CA rotation in progress for the swarm
	Spec SwarmSpec `json:"Spec,omitempty"` // User modifiable swarm configuration.
	Tlsinfo TLSInfo `json:"TLSInfo,omitempty"` // Information about the issuer of leaf TLS certificates and the trusted root CA certificate
	Updatedat string `json:"UpdatedAt,omitempty"` // Date and time at which the swarm was last updated in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
	Version ObjectVersion `json:"Version,omitempty"` // The version number of the object such as node, service, etc. This is needed to avoid conflicting writes. The client must send the version number along with the modified specification when updating these objects. This approach ensures safe concurrency and determinism in that the change on the object may not be applied if the version number has changed from the last read. In other words, if two update requests specify the same base version, only one of the requests can succeed. As a result, two separate update requests that happen at the same time will not unintentionally overwrite each other.
	Jointokens JoinTokens `json:"JoinTokens,omitempty"` // JoinTokens contains the tokens workers and managers need to join the swarm.
}

// EndpointPortConfig represents the EndpointPortConfig schema from the OpenAPI specification
type EndpointPortConfig struct {
	Targetport int `json:"TargetPort,omitempty"` // The port inside the container.
	Name string `json:"Name,omitempty"`
	Protocol string `json:"Protocol,omitempty"`
	Publishedport int `json:"PublishedPort,omitempty"` // The port on the swarm hosts.
}

// BuildInfo represents the BuildInfo schema from the OpenAPI specification
type BuildInfo struct {
	ErrorField string `json:"error,omitempty"`
	Errordetail ErrorDetail `json:"errorDetail,omitempty"`
	Id string `json:"id,omitempty"`
	Progress string `json:"progress,omitempty"`
	Progressdetail ProgressDetail `json:"progressDetail,omitempty"`
	Status string `json:"status,omitempty"`
	Stream string `json:"stream,omitempty"`
}

// TaskSpec represents the TaskSpec schema from the OpenAPI specification
type TaskSpec struct {
	Placement map[string]interface{} `json:"Placement,omitempty"`
	Pluginspec map[string]interface{} `json:"PluginSpec,omitempty"` // Invalid when specified with `ContainerSpec`. *(Experimental release only.)*
	Forceupdate int `json:"ForceUpdate,omitempty"` // A counter that triggers an update even if no relevant parameters have been changed.
	Restartpolicy map[string]interface{} `json:"RestartPolicy,omitempty"` // Specification for the restart policy which applies to containers created as part of this service.
	Resources map[string]interface{} `json:"Resources,omitempty"` // Resource requirements which apply to each individual container created as part of the service.
	Runtime string `json:"Runtime,omitempty"` // Runtime is the type of runtime specified for the task executor.
	Containerspec map[string]interface{} `json:"ContainerSpec,omitempty"` // Invalid when specified with `PluginSpec`.
	Logdriver map[string]interface{} `json:"LogDriver,omitempty"` // Specifies the log driver to use for tasks created from this spec. If not present, the default one for the swarm will be used, finally falling back to the engine default if not specified.
	Networks []map[string]interface{} `json:"Networks,omitempty"`
}

// IdResponse represents the IdResponse schema from the OpenAPI specification
type IdResponse struct {
	Id string `json:"Id"` // The id of the newly created object.
}

// Commit represents the Commit schema from the OpenAPI specification
type Commit struct {
	Expected string `json:"Expected,omitempty"` // Commit ID of external tool expected by dockerd as set at build time.
	Id string `json:"ID,omitempty"` // Actual commit ID of external tool.
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	Addr string `json:"Addr,omitempty"` // IP address.
	Prefixlen int `json:"PrefixLen,omitempty"` // Mask length of the IP address.
}

// RestartPolicy represents the RestartPolicy schema from the OpenAPI specification
type RestartPolicy struct {
	Name string `json:"Name,omitempty"` // - Empty string means not to restart - `always` Always restart - `unless-stopped` Restart always except when the user has manually stopped the container - `on-failure` Restart only when the container exit code is non-zero
	Maximumretrycount int `json:"MaximumRetryCount,omitempty"` // If `on-failure` is used, the number of times to retry before giving up
}

// Driver represents the Driver schema from the OpenAPI specification
type Driver struct {
	Name string `json:"Name"` // Name of the driver.
	Options map[string]interface{} `json:"Options,omitempty"` // Key/value map of driver-specific options.
}

// ImageDeleteResponseItem represents the ImageDeleteResponseItem schema from the OpenAPI specification
type ImageDeleteResponseItem struct {
	Deleted string `json:"Deleted,omitempty"` // The image ID of an image that was deleted
	Untagged string `json:"Untagged,omitempty"` // The image ID of an image that was untagged
}

// PluginsInfo represents the PluginsInfo schema from the OpenAPI specification
type PluginsInfo struct {
	Network []string `json:"Network,omitempty"` // Names of available network-drivers, and network-driver plugins.
	Volume []string `json:"Volume,omitempty"` // Names of available volume-drivers, and network-driver plugins.
	Authorization []string `json:"Authorization,omitempty"` // Names of available authorization plugins.
	Log []string `json:"Log,omitempty"` // Names of available logging-drivers, and logging-driver plugins.
}

// Secret represents the Secret schema from the OpenAPI specification
type Secret struct {
	Version ObjectVersion `json:"Version,omitempty"` // The version number of the object such as node, service, etc. This is needed to avoid conflicting writes. The client must send the version number along with the modified specification when updating these objects. This approach ensures safe concurrency and determinism in that the change on the object may not be applied if the version number has changed from the last read. In other words, if two update requests specify the same base version, only one of the requests can succeed. As a result, two separate update requests that happen at the same time will not unintentionally overwrite each other.
	Createdat string `json:"CreatedAt,omitempty"`
	Id string `json:"ID,omitempty"`
	Spec SecretSpec `json:"Spec,omitempty"`
	Updatedat string `json:"UpdatedAt,omitempty"`
}

// EndpointSettings represents the EndpointSettings schema from the OpenAPI specification
type EndpointSettings struct {
	Aliases []string `json:"Aliases,omitempty"`
	Endpointid string `json:"EndpointID,omitempty"` // Unique ID for the service endpoint in a Sandbox.
	Ipprefixlen int `json:"IPPrefixLen,omitempty"` // Mask length of the IPv4 address.
	Macaddress string `json:"MacAddress,omitempty"` // MAC address for the endpoint on this network.
	Driveropts map[string]interface{} `json:"DriverOpts,omitempty"` // DriverOpts is a mapping of driver options and values. These options are passed directly to the driver and are driver specific.
	Ipamconfig EndpointIPAMConfig `json:"IPAMConfig,omitempty"` // EndpointIPAMConfig represents an endpoint's IPAM configuration.
	Globalipv6address string `json:"GlobalIPv6Address,omitempty"` // Global IPv6 address.
	Globalipv6prefixlen int64 `json:"GlobalIPv6PrefixLen,omitempty"` // Mask length of the global IPv6 address.
	Ipaddress string `json:"IPAddress,omitempty"` // IPv4 address.
	Links []string `json:"Links,omitempty"`
	Gateway string `json:"Gateway,omitempty"` // Gateway address for this network.
	Ipv6gateway string `json:"IPv6Gateway,omitempty"` // IPv6 gateway address.
	Networkid string `json:"NetworkID,omitempty"` // Unique ID of the network.
}

// Volume represents the Volume schema from the OpenAPI specification
type Volume struct {
	Usagedata map[string]interface{} `json:"UsageData,omitempty"` // Usage details about the volume. This information is used by the `GET /system/df` endpoint, and omitted in other endpoints.
	Createdat string `json:"CreatedAt,omitempty"` // Date/Time the volume was created.
	Options map[string]interface{} `json:"Options"` // The driver specific options used when creating the volume.
	Mountpoint string `json:"Mountpoint"` // Mount path of the volume on the host.
	Driver string `json:"Driver"` // Name of the volume driver used by the volume.
	Labels map[string]interface{} `json:"Labels"` // User-defined key/value metadata.
	Name string `json:"Name"` // Name of the volume.
	Scope string `json:"Scope"` // The level at which the volume exists. Either `global` for cluster-wide, or `local` for machine level.
	Status map[string]interface{} `json:"Status,omitempty"` // Low-level details about the volume, provided by the volume driver. Details are returned as a map with key/value pairs: `{"key":"value","key2":"value2"}`. The `Status` field is optional, and is omitted if the volume driver does not support this feature.
}

// ResourceObject represents the ResourceObject schema from the OpenAPI specification
type ResourceObject struct {
	Genericresources []map[string]interface{} `json:"GenericResources,omitempty"` // User-defined resources can be either Integer resources (e.g, `SSD=3`) or String resources (e.g, `GPU=UUID1`)
	Memorybytes int64 `json:"MemoryBytes,omitempty"`
	Nanocpus int64 `json:"NanoCPUs,omitempty"`
}

// DeviceMapping represents the DeviceMapping schema from the OpenAPI specification
type DeviceMapping struct {
	Cgrouppermissions string `json:"CgroupPermissions,omitempty"`
	Pathincontainer string `json:"PathInContainer,omitempty"`
	Pathonhost string `json:"PathOnHost,omitempty"`
}

// GraphDriverData represents the GraphDriverData schema from the OpenAPI specification
type GraphDriverData struct {
	Data map[string]interface{} `json:"Data"`
	Name string `json:"Name"`
}

// SwarmSpec represents the SwarmSpec schema from the OpenAPI specification
type SwarmSpec struct {
	Orchestration map[string]interface{} `json:"Orchestration,omitempty"` // Orchestration configuration.
	Raft map[string]interface{} `json:"Raft,omitempty"` // Raft configuration.
	Taskdefaults map[string]interface{} `json:"TaskDefaults,omitempty"` // Defaults for creating tasks in this cluster.
	Caconfig map[string]interface{} `json:"CAConfig,omitempty"` // CA configuration.
	Dispatcher map[string]interface{} `json:"Dispatcher,omitempty"` // Dispatcher configuration.
	Encryptionconfig map[string]interface{} `json:"EncryptionConfig,omitempty"` // Parameters related to encryption-at-rest.
	Labels map[string]interface{} `json:"Labels,omitempty"` // User-defined key/value metadata.
	Name string `json:"Name,omitempty"` // Name of the swarm.
}

// CreateImageInfo represents the CreateImageInfo schema from the OpenAPI specification
type CreateImageInfo struct {
	Status string `json:"status,omitempty"`
	ErrorField string `json:"error,omitempty"`
	Progress string `json:"progress,omitempty"`
	Progressdetail ProgressDetail `json:"progressDetail,omitempty"`
}

// IndexInfo represents the IndexInfo schema from the OpenAPI specification
type IndexInfo struct {
	Secure bool `json:"Secure,omitempty"` // Indicates if the the registry is part of the list of insecure registries. If `false`, the registry is insecure. Insecure registries accept un-encrypted (HTTP) and/or untrusted (HTTPS with certificates from unknown CAs) communication. > **Warning**: Insecure registries can be useful when running a local > registry. However, because its use creates security vulnerabilities > it should ONLY be enabled for testing purposes. For increased > security, users should add their CA to their system's list of > trusted CAs instead of enabling this option.
	Mirrors []string `json:"Mirrors,omitempty"` // List of mirrors, expressed as URIs.
	Name string `json:"Name,omitempty"` // Name of the registry, such as "docker.io".
	Official bool `json:"Official,omitempty"` // Indicates whether this is an official registry (i.e., Docker Hub / docker.io)
}

// NodeSpec represents the NodeSpec schema from the OpenAPI specification
type NodeSpec struct {
	Role string `json:"Role,omitempty"` // Role of the node.
	Availability string `json:"Availability,omitempty"` // Availability of the node.
	Labels map[string]interface{} `json:"Labels,omitempty"` // User-defined key/value metadata.
	Name string `json:"Name,omitempty"` // Name for the node.
}

// NodeDescription represents the NodeDescription schema from the OpenAPI specification
type NodeDescription struct {
	Engine EngineDescription `json:"Engine,omitempty"` // EngineDescription provides information about an engine.
	Hostname string `json:"Hostname,omitempty"`
	Platform Platform `json:"Platform,omitempty"` // Platform represents the platform (Arch/OS).
	Resources ResourceObject `json:"Resources,omitempty"` // An object describing the resources which can be advertised by a node and requested by a task
	Tlsinfo TLSInfo `json:"TLSInfo,omitempty"` // Information about the issuer of leaf TLS certificates and the trusted root CA certificate
}

// ObjectVersion represents the ObjectVersion schema from the OpenAPI specification
type ObjectVersion struct {
	Index int `json:"Index,omitempty"`
}

// JoinTokens represents the JoinTokens schema from the OpenAPI specification
type JoinTokens struct {
	Manager string `json:"Manager,omitempty"` // The token managers can use to join the swarm.
	Worker string `json:"Worker,omitempty"` // The token workers can use to join the swarm.
}

// Image represents the Image schema from the OpenAPI specification
type Image struct {
	Container string `json:"Container"`
	Os string `json:"Os"`
	Created string `json:"Created"`
	Config ContainerConfig `json:"Config,omitempty"` // Configuration for a container that is portable between hosts
	Osversion string `json:"OsVersion,omitempty"`
	Virtualsize int64 `json:"VirtualSize"`
	Graphdriver GraphDriverData `json:"GraphDriver"` // Information about a container's graph driver.
	Id string `json:"Id"`
	Containerconfig ContainerConfig `json:"ContainerConfig,omitempty"` // Configuration for a container that is portable between hosts
	Author string `json:"Author"`
	Comment string `json:"Comment"`
	Parent string `json:"Parent"`
	Repotags []string `json:"RepoTags,omitempty"`
	Metadata map[string]interface{} `json:"Metadata,omitempty"`
	Size int64 `json:"Size"`
	Dockerversion string `json:"DockerVersion"`
	Repodigests []string `json:"RepoDigests,omitempty"`
	Rootfs map[string]interface{} `json:"RootFS"`
	Architecture string `json:"Architecture"`
}

// Task represents the Task schema from the OpenAPI specification
type Task struct {
	Status map[string]interface{} `json:"Status,omitempty"`
	Spec TaskSpec `json:"Spec,omitempty"` // User modifiable task configuration.
	Updatedat string `json:"UpdatedAt,omitempty"`
	Version ObjectVersion `json:"Version,omitempty"` // The version number of the object such as node, service, etc. This is needed to avoid conflicting writes. The client must send the version number along with the modified specification when updating these objects. This approach ensures safe concurrency and determinism in that the change on the object may not be applied if the version number has changed from the last read. In other words, if two update requests specify the same base version, only one of the requests can succeed. As a result, two separate update requests that happen at the same time will not unintentionally overwrite each other.
	Assignedgenericresources []map[string]interface{} `json:"AssignedGenericResources,omitempty"` // User-defined resources can be either Integer resources (e.g, `SSD=3`) or String resources (e.g, `GPU=UUID1`)
	Createdat string `json:"CreatedAt,omitempty"`
	Name string `json:"Name,omitempty"` // Name of the task.
	Desiredstate string `json:"DesiredState,omitempty"`
	Id string `json:"ID,omitempty"` // The ID of the task.
	Nodeid string `json:"NodeID,omitempty"` // The ID of the node that this task is on.
	Serviceid string `json:"ServiceID,omitempty"` // The ID of the service this task is part of.
	Labels map[string]interface{} `json:"Labels,omitempty"` // User-defined key/value metadata.
	Slot int `json:"Slot,omitempty"`
}

// ServiceUpdateResponse represents the ServiceUpdateResponse schema from the OpenAPI specification
type ServiceUpdateResponse struct {
	Warnings []string `json:"Warnings,omitempty"` // Optional warning messages
}

// HostConfig represents the HostConfig schema from the OpenAPI specification
type HostConfig struct {
	Blkioweightdevice []map[string]interface{} `json:"BlkioWeightDevice,omitempty"` // Block IO weight (relative device weight) in the form `[{"Path": "device_path", "Weight": weight}]`.
	Ulimits []map[string]interface{} `json:"Ulimits,omitempty"` // A list of resource limits to set in the container. For example: `{"Name": "nofile", "Soft": 1024, "Hard": 2048}`"
	Devicecgrouprules []string `json:"DeviceCgroupRules,omitempty"` // a list of cgroup rules to apply to the container
	Memoryswap int64 `json:"MemorySwap,omitempty"` // Total memory limit (memory + swap). Set as `-1` to enable unlimited swap.
	Cpusetcpus string `json:"CpusetCpus,omitempty"` // CPUs in which to allow execution (e.g., `0-3`, `0,1`)
	Iomaximumbandwidth int64 `json:"IOMaximumBandwidth,omitempty"` // Maximum IO in bytes per second for the container system drive (Windows only)
	Nanocpus int64 `json:"NanoCPUs,omitempty"` // CPU quota in units of 10<sup>-9</sup> CPUs.
	Blkiodevicereadiops []ThrottleDevice `json:"BlkioDeviceReadIOps,omitempty"` // Limit read rate (IO per second) from a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	Memoryswappiness int64 `json:"MemorySwappiness,omitempty"` // Tune a container's memory swappiness behavior. Accepts an integer between 0 and 100.
	Oomkilldisable bool `json:"OomKillDisable,omitempty"` // Disable OOM Killer for the container.
	Iomaximumiops int64 `json:"IOMaximumIOps,omitempty"` // Maximum IOps for the container system drive (Windows only)
	Cpucount int64 `json:"CpuCount,omitempty"` // The number of usable CPUs (Windows only). On Windows Server containers, the processor resource controls are mutually exclusive. The order of precedence is `CPUCount` first, then `CPUShares`, and `CPUPercent` last.
	Diskquota int64 `json:"DiskQuota,omitempty"` // Disk limit (in bytes).
	Blkiodevicewritebps []ThrottleDevice `json:"BlkioDeviceWriteBps,omitempty"` // Limit write rate (bytes per second) to a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	Cpushares int `json:"CpuShares,omitempty"` // An integer value representing this container's relative CPU weight versus other containers.
	Devices []DeviceMapping `json:"Devices,omitempty"` // A list of devices to add to the container.
	Cpupercent int64 `json:"CpuPercent,omitempty"` // The usable percentage of the available CPUs (Windows only). On Windows Server containers, the processor resource controls are mutually exclusive. The order of precedence is `CPUCount` first, then `CPUShares`, and `CPUPercent` last.
	Cpurealtimeruntime int64 `json:"CpuRealtimeRuntime,omitempty"` // The length of a CPU real-time runtime in microseconds. Set to 0 to allocate no time allocated to real-time tasks.
	Pidslimit int64 `json:"PidsLimit,omitempty"` // Tune a container's pids limit. Set -1 for unlimited.
	Cpuperiod int64 `json:"CpuPeriod,omitempty"` // The length of a CPU period in microseconds.
	Blkiodevicewriteiops []ThrottleDevice `json:"BlkioDeviceWriteIOps,omitempty"` // Limit write rate (IO per second) to a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	Cpusetmems string `json:"CpusetMems,omitempty"` // Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.
	Blkioweight int `json:"BlkioWeight,omitempty"` // Block IO weight (relative weight).
	Kernelmemory int64 `json:"KernelMemory,omitempty"` // Kernel memory limit in bytes.
	Cpurealtimeperiod int64 `json:"CpuRealtimePeriod,omitempty"` // The length of a CPU real-time period in microseconds. Set to 0 to allocate no time allocated to real-time tasks.
	Cpuquota int64 `json:"CpuQuota,omitempty"` // Microseconds of CPU time that the container can get in a CPU period.
	Cgroupparent string `json:"CgroupParent,omitempty"` // Path to `cgroups` under which the container's `cgroup` is created. If the path is not absolute, the path is considered to be relative to the `cgroups` path of the init process. Cgroups are created if they do not already exist.
	Blkiodevicereadbps []ThrottleDevice `json:"BlkioDeviceReadBps,omitempty"` // Limit read rate (bytes per second) from a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	Memory int `json:"Memory,omitempty"` // Memory limit in bytes.
	Memoryreservation int64 `json:"MemoryReservation,omitempty"` // Memory soft limit in bytes.
	Privileged bool `json:"Privileged,omitempty"` // Gives the container full access to the host.
	Cgroup string `json:"Cgroup,omitempty"` // Cgroup to use for the container.
	Dns []string `json:"Dns,omitempty"` // A list of DNS servers for the container to use.
	Usernsmode string `json:"UsernsMode,omitempty"` // Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
	Securityopt []string `json:"SecurityOpt,omitempty"` // A list of string values to customize labels for MLS systems, such as SELinux.
	Capadd []string `json:"CapAdd,omitempty"` // A list of kernel capabilities to add to the container.
	Isolation string `json:"Isolation,omitempty"` // Isolation technology of the container. (Windows only)
	Extrahosts []string `json:"ExtraHosts,omitempty"` // A list of hostnames/IP mappings to add to the container's `/etc/hosts` file. Specified in the form `["hostname:IP"]`.
	Containeridfile string `json:"ContainerIDFile,omitempty"` // Path to a file where the container ID is written
	Ipcmode string `json:"IpcMode,omitempty"` // IPC sharing mode for the container. Possible values are: - `"none"`: own private IPC namespace, with /dev/shm not mounted - `"private"`: own private IPC namespace - `"shareable"`: own private IPC namespace, with a possibility to share it with other containers - `"container:<name|id>"`: join another (shareable) container's IPC namespace - `"host"`: use the host system's IPC namespace If not specified, daemon default is used, which can either be `"private"` or `"shareable"`, depending on daemon version and configuration.
	Binds []string `json:"Binds,omitempty"` // A list of volume bindings for this container. Each volume binding is a string in one of these forms: - `host-src:container-dest` to bind-mount a host path into the container. Both `host-src`, and `container-dest` must be an _absolute_ path. - `host-src:container-dest:ro` to make the bind mount read-only inside the container. Both `host-src`, and `container-dest` must be an _absolute_ path. - `volume-name:container-dest` to bind-mount a volume managed by a volume driver into the container. `container-dest` must be an _absolute_ path. - `volume-name:container-dest:ro` to mount the volume read-only inside the container. `container-dest` must be an _absolute_ path.
	Logconfig map[string]interface{} `json:"LogConfig,omitempty"` // The logging configuration for this container
	Mounts []Mount `json:"Mounts,omitempty"` // Specification for mounts to be added to the container.
	Capdrop []string `json:"CapDrop,omitempty"` // A list of kernel capabilities to drop from the container.
	Dnsoptions []string `json:"DnsOptions,omitempty"` // A list of DNS options.
	Networkmode string `json:"NetworkMode,omitempty"` // Network mode to use for this container. Supported standard values are: `bridge`, `host`, `none`, and `container:<name|id>`. Any other value is taken as a custom network's name to which this container should connect to.
	Readonlyrootfs bool `json:"ReadonlyRootfs,omitempty"` // Mount the container's root filesystem as read only.
	Pidmode string `json:"PidMode,omitempty"` // Set the PID (Process) Namespace mode for the container. It can be either: - `"container:<name|id>"`: joins another container's PID namespace - `"host"`: use the host's PID namespace inside the container
	Oomscoreadj int `json:"OomScoreAdj,omitempty"` // An integer value containing the score given to the container in order to tune OOM killer preferences.
	Restartpolicy RestartPolicy `json:"RestartPolicy,omitempty"` // The behavior to apply when the container exits. The default is not to restart. An ever increasing delay (double the previous delay, starting at 100ms) is added before each restart to prevent flooding the server.
	Consolesize []int `json:"ConsoleSize,omitempty"` // Initial console size, as an `[height, width]` array. (Windows only)
	Portbindings map[string]interface{} `json:"PortBindings,omitempty"` // A map of exposed container ports and the host port they should map to.
	Storageopt map[string]interface{} `json:"StorageOpt,omitempty"` // Storage driver options for this container, in the form `{"size": "120G"}`.
	Volumesfrom []string `json:"VolumesFrom,omitempty"` // A list of volumes to inherit from another container, specified in the form `<container name>[:<ro|rw>]`.
	Publishallports bool `json:"PublishAllPorts,omitempty"` // Allocates a random host port for all of a container's exposed ports.
	Utsmode string `json:"UTSMode,omitempty"` // UTS namespace to use for the container.
	Links []string `json:"Links,omitempty"` // A list of links for the container in the form `container_name:alias`.
	Shmsize int `json:"ShmSize,omitempty"` // Size of `/dev/shm` in bytes. If omitted, the system uses 64MB.
	Sysctls map[string]interface{} `json:"Sysctls,omitempty"` // A list of kernel parameters (sysctls) to set in the container. For example: `{"net.ipv4.ip_forward": "1"}`
	Tmpfs map[string]interface{} `json:"Tmpfs,omitempty"` // A map of container directories which should be replaced by tmpfs mounts, and their corresponding mount options. For example: `{ "/run": "rw,noexec,nosuid,size=65536k" }`.
	Groupadd []string `json:"GroupAdd,omitempty"` // A list of additional groups that the container process will run as.
	Autoremove bool `json:"AutoRemove,omitempty"` // Automatically remove the container when the container's process exits. This has no effect if `RestartPolicy` is set.
	Runtime string `json:"Runtime,omitempty"` // Runtime to use with this container.
	Dnssearch []string `json:"DnsSearch,omitempty"` // A list of DNS search domains.
	Volumedriver string `json:"VolumeDriver,omitempty"` // Driver that this container uses to mount volumes.
}

// ServiceSpec represents the ServiceSpec schema from the OpenAPI specification
type ServiceSpec struct {
	Tasktemplate TaskSpec `json:"TaskTemplate,omitempty"` // User modifiable task configuration.
	Updateconfig map[string]interface{} `json:"UpdateConfig,omitempty"` // Specification for the update strategy of the service.
	Endpointspec EndpointSpec `json:"EndpointSpec,omitempty"` // Properties that can be configured to access and load balance a service.
	Labels map[string]interface{} `json:"Labels,omitempty"` // User-defined key/value metadata.
	Mode map[string]interface{} `json:"Mode,omitempty"` // Scheduling mode for the service.
	Name string `json:"Name,omitempty"` // Name of the service.
	Networks []map[string]interface{} `json:"Networks,omitempty"` // Array of network names or IDs to attach the service to.
	Rollbackconfig map[string]interface{} `json:"RollbackConfig,omitempty"` // Specification for the rollback strategy of the service.
}

// PluginMount represents the PluginMount schema from the OpenAPI specification
type PluginMount struct {
	Destination string `json:"Destination"`
	Name string `json:"Name"`
	Options []string `json:"Options"`
	Settable []string `json:"Settable"`
	Source string `json:"Source"`
	TypeField string `json:"Type"`
	Description string `json:"Description"`
}

// RegistryServiceConfig represents the RegistryServiceConfig schema from the OpenAPI specification
type RegistryServiceConfig struct {
	Mirrors []string `json:"Mirrors,omitempty"` // List of registry URLs that act as a mirror for the official (`docker.io`) registry.
	Allownondistributableartifactscidrs []string `json:"AllowNondistributableArtifactsCIDRs,omitempty"` // List of IP ranges to which nondistributable artifacts can be pushed, using the CIDR syntax [RFC 4632](https://tools.ietf.org/html/4632). Some images (for example, Windows base images) contain artifacts whose distribution is restricted by license. When these images are pushed to a registry, restricted artifacts are not included. This configuration override this behavior, and enables the daemon to push nondistributable artifacts to all registries whose resolved IP address is within the subnet described by the CIDR syntax. This option is useful when pushing images containing nondistributable artifacts to a registry on an air-gapped network so hosts on that network can pull the images without connecting to another server. > **Warning**: Nondistributable artifacts typically have restrictions > on how and where they can be distributed and shared. Only use this > feature to push artifacts to private registries and ensure that you > are in compliance with any terms that cover redistributing > nondistributable artifacts.
	Allownondistributableartifactshostnames []string `json:"AllowNondistributableArtifactsHostnames,omitempty"` // List of registry hostnames to which nondistributable artifacts can be pushed, using the format `<hostname>[:<port>]` or `<IP address>[:<port>]`. Some images (for example, Windows base images) contain artifacts whose distribution is restricted by license. When these images are pushed to a registry, restricted artifacts are not included. This configuration override this behavior for the specified registries. This option is useful when pushing images containing nondistributable artifacts to a registry on an air-gapped network so hosts on that network can pull the images without connecting to another server. > **Warning**: Nondistributable artifacts typically have restrictions > on how and where they can be distributed and shared. Only use this > feature to push artifacts to private registries and ensure that you > are in compliance with any terms that cover redistributing > nondistributable artifacts.
	Indexconfigs map[string]interface{} `json:"IndexConfigs,omitempty"`
	Insecureregistrycidrs []string `json:"InsecureRegistryCIDRs,omitempty"` // List of IP ranges of insecure registries, using the CIDR syntax ([RFC 4632](https://tools.ietf.org/html/4632)). Insecure registries accept un-encrypted (HTTP) and/or untrusted (HTTPS with certificates from unknown CAs) communication. By default, local registries (`127.0.0.0/8`) are configured as insecure. All other registries are secure. Communicating with an insecure registry is not possible if the daemon assumes that registry is secure. This configuration override this behavior, insecure communication with registries whose resolved IP address is within the subnet described by the CIDR syntax. Registries can also be marked insecure by hostname. Those registries are listed under `IndexConfigs` and have their `Secure` field set to `false`. > **Warning**: Using this option can be useful when running a local > registry, but introduces security vulnerabilities. This option > should therefore ONLY be used for testing purposes. For increased > security, users should add their CA to their system's list of trusted > CAs instead of enabling this option.
}

// AuthConfig represents the AuthConfig schema from the OpenAPI specification
type AuthConfig struct {
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Serveraddress string `json:"serveraddress,omitempty"`
	Username string `json:"username,omitempty"`
}

// ProgressDetail represents the ProgressDetail schema from the OpenAPI specification
type ProgressDetail struct {
	Code int `json:"code,omitempty"`
	Message int `json:"message,omitempty"`
}

// NetworkSettings represents the NetworkSettings schema from the OpenAPI specification
type NetworkSettings struct {
	Ipv6gateway string `json:"IPv6Gateway,omitempty"` // IPv6 gateway address for this network. <p><br /></p> > **Deprecated**: This field is only propagated when attached to the > default "bridge" network. Use the information from the "bridge" > network inside the `Networks` map instead, which contains the same > information. This field was deprecated in Docker 1.9 and is scheduled > to be removed in Docker 17.12.0
	Hairpinmode bool `json:"HairpinMode,omitempty"` // Indicates if hairpin NAT should be enabled on the virtual interface.
	Ports PortMap `json:"Ports,omitempty"` // PortMap describes the mapping of container ports to host ports, using the container's port-number and protocol as key in the format `<port>/<protocol>`, for example, `80/udp`. If a container's port is mapped for both `tcp` and `udp`, two separate entries are added to the mapping table.
	Sandboxid string `json:"SandboxID,omitempty"` // SandboxID uniquely represents a container's network stack.
	Gateway string `json:"Gateway,omitempty"` // Gateway address for the default "bridge" network. <p><br /></p> > **Deprecated**: This field is only propagated when attached to the > default "bridge" network. Use the information from the "bridge" > network inside the `Networks` map instead, which contains the same > information. This field was deprecated in Docker 1.9 and is scheduled > to be removed in Docker 17.12.0
	Globalipv6address string `json:"GlobalIPv6Address,omitempty"` // Global IPv6 address for the default "bridge" network. <p><br /></p> > **Deprecated**: This field is only propagated when attached to the > default "bridge" network. Use the information from the "bridge" > network inside the `Networks` map instead, which contains the same > information. This field was deprecated in Docker 1.9 and is scheduled > to be removed in Docker 17.12.0
	Globalipv6prefixlen int `json:"GlobalIPv6PrefixLen,omitempty"` // Mask length of the global IPv6 address. <p><br /></p> > **Deprecated**: This field is only propagated when attached to the > default "bridge" network. Use the information from the "bridge" > network inside the `Networks` map instead, which contains the same > information. This field was deprecated in Docker 1.9 and is scheduled > to be removed in Docker 17.12.0
	Macaddress string `json:"MacAddress,omitempty"` // MAC address for the container on the default "bridge" network. <p><br /></p> > **Deprecated**: This field is only propagated when attached to the > default "bridge" network. Use the information from the "bridge" > network inside the `Networks` map instead, which contains the same > information. This field was deprecated in Docker 1.9 and is scheduled > to be removed in Docker 17.12.0
	Secondaryipaddresses []Address `json:"SecondaryIPAddresses,omitempty"`
	Linklocalipv6address string `json:"LinkLocalIPv6Address,omitempty"` // IPv6 unicast address using the link-local prefix.
	Networks map[string]interface{} `json:"Networks,omitempty"` // Information about all networks that the container is connected to.
	Linklocalipv6prefixlen int `json:"LinkLocalIPv6PrefixLen,omitempty"` // Prefix length of the IPv6 unicast address.
	Sandboxkey string `json:"SandboxKey,omitempty"` // SandboxKey identifies the sandbox
	Bridge string `json:"Bridge,omitempty"` // Name of the network'a bridge (for example, `docker0`).
	Endpointid string `json:"EndpointID,omitempty"` // EndpointID uniquely represents a service endpoint in a Sandbox. <p><br /></p> > **Deprecated**: This field is only propagated when attached to the > default "bridge" network. Use the information from the "bridge" > network inside the `Networks` map instead, which contains the same > information. This field was deprecated in Docker 1.9 and is scheduled > to be removed in Docker 17.12.0
	Ipprefixlen int `json:"IPPrefixLen,omitempty"` // Mask length of the IPv4 address. <p><br /></p> > **Deprecated**: This field is only propagated when attached to the > default "bridge" network. Use the information from the "bridge" > network inside the `Networks` map instead, which contains the same > information. This field was deprecated in Docker 1.9 and is scheduled > to be removed in Docker 17.12.0
	Ipaddress string `json:"IPAddress,omitempty"` // IPv4 address for the default "bridge" network. <p><br /></p> > **Deprecated**: This field is only propagated when attached to the > default "bridge" network. Use the information from the "bridge" > network inside the `Networks` map instead, which contains the same > information. This field was deprecated in Docker 1.9 and is scheduled > to be removed in Docker 17.12.0
	Secondaryipv6addresses []Address `json:"SecondaryIPv6Addresses,omitempty"`
}

// SwarmInfo represents the SwarmInfo schema from the OpenAPI specification
type SwarmInfo struct {
	Nodeid string `json:"NodeID,omitempty"` // Unique identifier of for this node in the swarm.
	Remotemanagers []PeerNode `json:"RemoteManagers,omitempty"` // List of ID's and addresses of other managers in the swarm.
	Localnodestate string `json:"LocalNodeState,omitempty"` // Current local status of this node.
	Nodes int `json:"Nodes,omitempty"` // Total number of nodes in the swarm.
	ErrorField string `json:"Error,omitempty"`
	Nodeaddr string `json:"NodeAddr,omitempty"` // IP address at which this node can be reached by other nodes in the swarm.
	Cluster ClusterInfo `json:"Cluster,omitempty"` // ClusterInfo represents information about the swarm as is returned by the "/info" endpoint. Join-tokens are not included.
	Controlavailable bool `json:"ControlAvailable,omitempty"`
	Managers int `json:"Managers,omitempty"` // Total number of managers in the swarm.
}

// Runtime represents the Runtime schema from the OpenAPI specification
type Runtime struct {
	Path string `json:"path,omitempty"` // Name and, optional, path, of the OCI executable binary. If the path is omitted, the daemon searches the host's `$PATH` for the binary and uses the first result.
	Runtimeargs []string `json:"runtimeArgs,omitempty"` // List of command-line arguments to pass to the runtime when invoked.
}

// PortMap represents the PortMap schema from the OpenAPI specification
type PortMap struct {
}

// Service represents the Service schema from the OpenAPI specification
type Service struct {
	Createdat string `json:"CreatedAt,omitempty"`
	Endpoint map[string]interface{} `json:"Endpoint,omitempty"`
	Id string `json:"ID,omitempty"`
	Spec ServiceSpec `json:"Spec,omitempty"` // User modifiable configuration for a service.
	Updatestatus map[string]interface{} `json:"UpdateStatus,omitempty"` // The status of a service update.
	Updatedat string `json:"UpdatedAt,omitempty"`
	Version ObjectVersion `json:"Version,omitempty"` // The version number of the object such as node, service, etc. This is needed to avoid conflicting writes. The client must send the version number along with the modified specification when updating these objects. This approach ensures safe concurrency and determinism in that the change on the object may not be applied if the version number has changed from the last read. In other words, if two update requests specify the same base version, only one of the requests can succeed. As a result, two separate update requests that happen at the same time will not unintentionally overwrite each other.
}

// ManagerStatus represents the ManagerStatus schema from the OpenAPI specification
type ManagerStatus struct {
	Addr string `json:"Addr,omitempty"` // The IP address and port at which the manager is reachable.
	Leader bool `json:"Leader,omitempty"`
	Reachability string `json:"Reachability,omitempty"` // Reachability represents the reachability of a node.
}

// PluginEnv represents the PluginEnv schema from the OpenAPI specification
type PluginEnv struct {
	Description string `json:"Description"`
	Name string `json:"Name"`
	Settable []string `json:"Settable"`
	Value string `json:"Value"`
}

// ThrottleDevice represents the ThrottleDevice schema from the OpenAPI specification
type ThrottleDevice struct {
	Path string `json:"Path,omitempty"` // Device path
	Rate int64 `json:"Rate,omitempty"` // Rate
}

// SystemInfo represents the SystemInfo schema from the OpenAPI specification
type SystemInfo struct {
	Runtimes map[string]interface{} `json:"Runtimes,omitempty"` // List of [OCI compliant](https://github.com/opencontainers/runtime-spec) runtimes configured on the daemon. Keys hold the "name" used to reference the runtime. The Docker daemon relies on an OCI compliant runtime (invoked via the `containerd` daemon) as its interface to the Linux kernel namespaces, cgroups, and SELinux. The default runtime is `runc`, and automatically configured. Additional runtimes can be configured by the user and will be listed here.
	Containersstopped int `json:"ContainersStopped,omitempty"` // Number of containers with status `"stopped"`.
	Debug bool `json:"Debug,omitempty"` // Indicates if the daemon is running in debug-mode / with debug-level logging enabled.
	Clusteradvertise string `json:"ClusterAdvertise,omitempty"` // The network endpoint that the Engine advertises for the purpose of node discovery. ClusterAdvertise is a `host:port` combination on which the daemon is reachable by other hosts. <p><br /></p> > **Note**: This field is only propagated when using standalone Swarm > mode, and overlay networking using an external k/v store. Overlay > networks with Swarm mode enabled use the built-in raft store, and > this field will be empty.
	Liverestoreenabled bool `json:"LiveRestoreEnabled,omitempty"` // Indicates if live restore is enabled. If enabled, containers are kept running when the daemon is shutdown or upon daemon start if running containers are detected.
	Nfd int `json:"NFd,omitempty"` // The total number of file Descriptors in use by the daemon process. This information is only returned if debug-mode is enabled.
	Id string `json:"ID,omitempty"` // Unique identifier of the daemon. <p><br /></p> > **Note**: The format of the ID itself is not part of the API, and > should not be considered stable.
	Ipv4forwarding bool `json:"IPv4Forwarding,omitempty"` // Indicates IPv4 forwarding is enabled.
	Genericresources []map[string]interface{} `json:"GenericResources,omitempty"` // User-defined resources can be either Integer resources (e.g, `SSD=3`) or String resources (e.g, `GPU=UUID1`)
	Httpproxy string `json:"HttpProxy,omitempty"` // HTTP-proxy configured for the daemon. This value is obtained from the [`HTTP_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html) environment variable. Containers do not automatically inherit this configuration.
	Securityoptions []string `json:"SecurityOptions,omitempty"` // List of security features that are enabled on the daemon, such as apparmor, seccomp, SELinux, and user-namespaces (userns). Additional configuration options for each security feature may be present, and are included as a comma-separated list of key/value pairs.
	Name string `json:"Name,omitempty"` // Hostname of the host.
	Ncpu int `json:"NCPU,omitempty"` // The number of logical CPUs usable by the daemon. The number of available CPUs is checked by querying the operating system when the daemon starts. Changes to operating system CPU allocation after the daemon is started are not reflected.
	Containerdcommit Commit `json:"ContainerdCommit,omitempty"` // Commit holds the Git-commit (SHA1) that a binary was built from, as reported in the version-string of external tools, such as `containerd`, or `runC`.
	Initbinary string `json:"InitBinary,omitempty"` // Name and, optional, path of the the `docker-init` binary. If the path is omitted, the daemon searches the host's `$PATH` for the binary and uses the first result.
	Plugins PluginsInfo `json:"Plugins,omitempty"` // Available plugins per type. <p><br /></p> > **Note**: Only unmanaged (V1) plugins are included in this list. > V1 plugins are "lazily" loaded, and are not returned in this list > if there is no resource using the plugin.
	Systemtime string `json:"SystemTime,omitempty"` // Current system-time in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
	Swaplimit bool `json:"SwapLimit,omitempty"` // Indicates if the host has memory swap limit support enabled.
	Memorylimit bool `json:"MemoryLimit,omitempty"` // Indicates if the host has memory limit support enabled.
	Runccommit Commit `json:"RuncCommit,omitempty"` // Commit holds the Git-commit (SHA1) that a binary was built from, as reported in the version-string of external tools, such as `containerd`, or `runC`.
	Noproxy string `json:"NoProxy,omitempty"` // Comma-separated list of domain extensions for which no proxy should be used. This value is obtained from the [`NO_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html) environment variable. Containers do not automatically inherit this configuration.
	Swarm SwarmInfo `json:"Swarm,omitempty"` // Represents generic information about swarm.
	Cpucfsperiod bool `json:"CpuCfsPeriod,omitempty"` // Indicates if CPU CFS(Completely Fair Scheduler) period is supported by the host.
	Cgroupdriver string `json:"CgroupDriver,omitempty"` // The driver to use for managing cgroups.
	Memtotal int64 `json:"MemTotal,omitempty"` // Total amount of physical memory available on the host, in kilobytes (kB).
	Cpuset bool `json:"CPUSet,omitempty"` // Indicates if CPUsets (cpuset.cpus, cpuset.mems) are supported by the host. See [cpuset(7)](https://www.kernel.org/doc/Documentation/cgroup-v1/cpusets.txt)
	Kernelmemory bool `json:"KernelMemory,omitempty"` // Indicates if the host has kernel memory limit support enabled.
	Systemstatus [][]string `json:"SystemStatus,omitempty"` // Status information about this node (standalone Swarm API). <p><br /></p> > **Note**: The information returned in this field is only propagated > by the Swarm standalone API, and is empty (`null`) when using > built-in swarm mode.
	Indexserveraddress string `json:"IndexServerAddress,omitempty"` // Address / URL of the index server that is used for image search, and as a default for user authentication for Docker Hub and Docker Cloud.
	Operatingsystem string `json:"OperatingSystem,omitempty"` // Name of the host's operating system, for example: "Ubuntu 16.04.2 LTS" or "Windows Server 2016 Datacenter"
	Initcommit Commit `json:"InitCommit,omitempty"` // Commit holds the Git-commit (SHA1) that a binary was built from, as reported in the version-string of external tools, such as `containerd`, or `runC`.
	Bridgenfip6tables bool `json:"BridgeNfIp6tables,omitempty"` // Indicates if `bridge-nf-call-ip6tables` is available on the host.
	Architecture string `json:"Architecture,omitempty"` // Hardware architecture of the host, as returned by the Go runtime (`GOARCH`). A full list of possible values can be found in the [Go documentation](https://golang.org/doc/install/source#environment).
	Kernelversion string `json:"KernelVersion,omitempty"` // Kernel version of the host. On Linux, this information obtained from `uname`. On Windows this information is queried from the <kbd>HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\</kbd> registry value, for example _"10.0 14393 (14393.1198.amd64fre.rs1_release_sec.170427-1353)"_.
	Driver string `json:"Driver,omitempty"` // Name of the storage driver in use.
	Ngoroutines int `json:"NGoroutines,omitempty"` // The number of goroutines that currently exist. This information is only returned if debug-mode is enabled.
	Ostype string `json:"OSType,omitempty"` // Generic type of the operating system of the host, as returned by the Go runtime (`GOOS`). Currently returned values are "linux" and "windows". A full list of possible values can be found in the [Go documentation](https://golang.org/doc/install/source#environment).
	Dockerrootdir string `json:"DockerRootDir,omitempty"` // Root directory of persistent Docker state. Defaults to `/var/lib/docker` on Linux, and `C:\ProgramData\docker` on Windows.
	Labels []string `json:"Labels,omitempty"` // User-defined labels (key/value metadata) as set on the daemon. <p><br /></p> > **Note**: When part of a Swarm, nodes can both have _daemon_ labels, > set through the daemon configuration, and _node_ labels, set from a > manager node in the Swarm. Node labels are not included in this > field. Node labels can be retrieved using the `/nodes/(id)` endpoint > on a manager node in the Swarm.
	Loggingdriver string `json:"LoggingDriver,omitempty"` // The logging driver to use as a default for new containers.
	Experimentalbuild bool `json:"ExperimentalBuild,omitempty"` // Indicates if experimental features are enabled on the daemon.
	Cpushares bool `json:"CPUShares,omitempty"` // Indicates if CPU Shares limiting is supported by the host.
	Isolation string `json:"Isolation,omitempty"` // Represents the isolation technology to use as a default for containers. The supported values are platform-specific. If no isolation value is specified on daemon start, on Windows client, the default is `hyperv`, and on Windows server, the default is `process`. This option is currently not used on other platforms.
	Bridgenfiptables bool `json:"BridgeNfIptables,omitempty"` // Indicates if `bridge-nf-call-iptables` is available on the host.
	Neventslistener int `json:"NEventsListener,omitempty"` // Number of event listeners subscribed.
	Serverversion string `json:"ServerVersion,omitempty"` // Version string of the daemon. > **Note**: the [standalone Swarm API](https://docs.docker.com/swarm/swarm-api/) > returns the Swarm version instead of the daemon version, for example > `swarm/1.2.8`.
	Registryconfig RegistryServiceConfig `json:"RegistryConfig,omitempty"` // RegistryServiceConfig stores daemon registry services configuration.
	Defaultruntime string `json:"DefaultRuntime,omitempty"` // Name of the default OCI runtime that is used when starting containers. The default can be overridden per-container at create time.
	Httpsproxy string `json:"HttpsProxy,omitempty"` // HTTPS-proxy configured for the daemon. This value is obtained from the [`HTTPS_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html) environment variable. Containers do not automatically inherit this configuration.
	Oomkilldisable bool `json:"OomKillDisable,omitempty"` // Indicates if OOM killer disable is supported on the host.
	Clusterstore string `json:"ClusterStore,omitempty"` // URL of the distributed storage backend. The storage backend is used for multihost networking (to store network and endpoint information) and by the node discovery mechanism. <p><br /></p> > **Note**: This field is only propagated when using standalone Swarm > mode, and overlay networking using an external k/v store. Overlay > networks with Swarm mode enabled use the built-in raft store, and > this field will be empty.
	Containersrunning int `json:"ContainersRunning,omitempty"` // Number of containers with status `"running"`.
	Containerspaused int `json:"ContainersPaused,omitempty"` // Number of containers with status `"paused"`.
	Driverstatus [][]string `json:"DriverStatus,omitempty"` // Information specific to the storage driver, provided as "label" / "value" pairs. This information is provided by the storage driver, and formatted in a way consistent with the output of `docker info` on the command line. <p><br /></p> > **Note**: The information returned in this field, including the > formatting of values and labels, should not be considered stable, > and may change without notice.
	Cpucfsquota bool `json:"CpuCfsQuota,omitempty"` // Indicates if CPU CFS(Completely Fair Scheduler) quota is supported by the host.
	Images int `json:"Images,omitempty"` // Total number of images on the host. Both _tagged_ and _untagged_ (dangling) images are counted.
	Containers int `json:"Containers,omitempty"` // Total number of containers on the host.
}

// MountPoint represents the MountPoint schema from the OpenAPI specification
type MountPoint struct {
	Driver string `json:"Driver,omitempty"`
	Mode string `json:"Mode,omitempty"`
	Name string `json:"Name,omitempty"`
	Propagation string `json:"Propagation,omitempty"`
	Rw bool `json:"RW,omitempty"`
	Source string `json:"Source,omitempty"`
	TypeField string `json:"Type,omitempty"`
	Destination string `json:"Destination,omitempty"`
}

// NodeStatus represents the NodeStatus schema from the OpenAPI specification
type NodeStatus struct {
	Addr string `json:"Addr,omitempty"` // IP address of the node.
	Message string `json:"Message,omitempty"`
	State string `json:"State,omitempty"` // NodeState represents the state of a node.
}

// EndpointIPAMConfig represents the EndpointIPAMConfig schema from the OpenAPI specification
type EndpointIPAMConfig struct {
	Ipv6address string `json:"IPv6Address,omitempty"`
	Linklocalips []string `json:"LinkLocalIPs,omitempty"`
	Ipv4address string `json:"IPv4Address,omitempty"`
}

// Port represents the Port schema from the OpenAPI specification
type Port struct {
	Ip string `json:"IP,omitempty"`
	Privateport int `json:"PrivatePort"` // Port on the container
	Publicport int `json:"PublicPort,omitempty"` // Port exposed on the host
	TypeField string `json:"Type"`
}

// Node represents the Node schema from the OpenAPI specification
type Node struct {
	Status NodeStatus `json:"Status,omitempty"` // NodeStatus represents the status of a node. It provides the current status of the node, as seen by the manager.
	Updatedat string `json:"UpdatedAt,omitempty"` // Date and time at which the node was last updated in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
	Version ObjectVersion `json:"Version,omitempty"` // The version number of the object such as node, service, etc. This is needed to avoid conflicting writes. The client must send the version number along with the modified specification when updating these objects. This approach ensures safe concurrency and determinism in that the change on the object may not be applied if the version number has changed from the last read. In other words, if two update requests specify the same base version, only one of the requests can succeed. As a result, two separate update requests that happen at the same time will not unintentionally overwrite each other.
	Createdat string `json:"CreatedAt,omitempty"` // Date and time at which the node was added to the swarm in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
	Description NodeDescription `json:"Description,omitempty"` // NodeDescription encapsulates the properties of the Node as reported by the agent.
	Id string `json:"ID,omitempty"`
	Managerstatus ManagerStatus `json:"ManagerStatus,omitempty"` // ManagerStatus represents the status of a manager. It provides the current status of a node's manager component, if the node is a manager.
	Spec NodeSpec `json:"Spec,omitempty"`
}

// PluginDevice represents the PluginDevice schema from the OpenAPI specification
type PluginDevice struct {
	Path string `json:"Path"`
	Settable []string `json:"Settable"`
	Description string `json:"Description"`
	Name string `json:"Name"`
}
