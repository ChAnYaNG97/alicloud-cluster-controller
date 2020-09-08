package v1

const (

	//VPC的状态，取值：
	//Pending：配置中。
	//Available：可用。
	StatusPending   = "Pending"
	StatusAvailable = "Available"
)

// VPCSpec 专有网络
// 使用云资源前, 必须先创建一个专有网络和交换机
// 详细文档见 [CreateVpc](https://help.aliyun.com/document_detail/35737.html)
type VpcSpec struct {

	// 使用一个已经存在的VPC
	VpcId string `json:"vpc_id,omitempty"`

	// 专有网络名称。长度为2-128个字符，必须以字母或中文开头，可包含数字，点号（.），下划线（_）和短横线（-），但不能以http://或https://开头。
	VpcName string `json:"vpc_ame,omitempty"`
	// VPC的网段。您可以使用以下网段或其子集：
	//   10.0.0.0/8。
	//   172.16.0.0/12（默认值）。
	//   192.168.0.0/16。
	CidrBlock string `json:"cidr_block,omitempty"`

	// VPC的描述信息。长度为2-256个字符，必须以字母或中文开头，但不能以http://或https://开头。
	Description string `json:"description,omitempty"`

	// 用户侧网络的网段，如需定义多个网段请使用半角逗号隔开，最多支持3个网段。
	//
	// VPC定义的默认私网转发网段为10.0.0.0/8、172.16.0.0/12、192.168.0.0/16、100.64.0.0/10和VPC CIDR网段。
	// 如果ECS实例或弹性网卡已经具备了公网访问能力（ECS实例分配了固定公网IP、ECS实例或弹性网卡绑定了公网IP、ECS实例或弹性网卡设置了DNAT IP映射规则），
	// 这类资源访问非上述默认私网转发网段的请求均会通过公网IP直接转发至公网。
	// 当希望按照路由表在私网（如VPC内、通过VPN/高速通道/云企业网搭建的混合云网络）转发访问非上述默认私网网段的请求时，
	// 需要将网络请求的目的网段设置为ECS或弹性网卡所在VPC的UserCidr。为VPC设置UserCidr后，
	// 该VPC中访问UserCidr地址的请求将按照路由表进行转发，而不通过公网IP转发。
	// UserCidr string `json:"userCidr,omitempty"`

	// 是否开启IPv6网段，取值：
	//   false（默认值）：不开启。
	//   true：开启。
	// EnableIpv6 string `json:"enableIpv6,omitempty"`
	// VPC的IPv6网段
	// Ipv6CidrBlock string `json:"ipv6CidrBlock,omitempty"`

	// 资源组ID
	// ResourceGroupId	string `json:"resourceGroupId,omitempty"`

	// 是否只预检此次请求，取值：
	//
	// true：发送检查请求，不会创建VPC。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码DryRunOperation。
	// false（默认值）：发送正常请求，通过检查后返回2XX HTTP状态码并直接创建VPC。
	// DryRun bool `json:"dryRun,omitempty"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。更多详情，请参见如何保证幂等性。
	// ClientToken	string `json:"clientToken,omitempty"`
}

type VSwitchSpec struct {
	// 使用存在的VSwitch
	VSwitchId string `json:"vswitch_id,omitempty"`

	// 交换机的网段。交换机网段要求如下：
	// 交换机的网段的掩码长度范围为16~29位。
	// 交换机的网段必须从属于所在VPC的网段。
	// 交换机的网段不能与所在VPC中路由条目的目标网段相同，但可以是目标网段的子集。
	CidrBlock string `json:"cidr_block,omitempty"`

	// 要创建的交换机所属的VPC ID。
	VpcId string `json:"vpc_id,omitempty"`

	// 交换机的名称。
	// 名称长度为2~128个字符，必须以字母或中文开头，但不能以http://或https://开头。
	VSwitchName string `json:"vswitch_name,omitempty"`

	// 交换机的描述信息。
	//描述长度为2~256个字符，必须以字母或中文开头，但不能以http://或https://开头。
	Description string `json:"description,omitempty"`

	// 可用区ID
	ZoneId string `json:"zone_id,omitempty"`
}

type ClusterSpec struct {
	ClusterType              string      `json:"cluster_type"`
	Name                     string      `json:"name"`
	RegionID                 string      `json:"region_id"`
	DisableRollback          bool        `json:"disable_rollback"`
	TimeoutMins              int         `json:"timeout_mins"`
	KubernetesVersion        string      `json:"kubernetes_version"`
	SnatEntry                bool        `json:"snat_entry"`
	EndpointPublicAccess     bool        `json:"endpoint_public_access"`
	SSHFlags                 bool        `json:"ssh_flags"`
	CloudMonitorFlags        bool        `json:"cloud_monitor_flags"`
	DeletionProtection       bool        `json:"deletion_protection"`
	NodeCidrMask             string      `json:"node_cidr_mask"`
	ProxyMode                string      `json:"proxy_mode"`
	Tags                     []Tag       `json:"tags"`
	Addons                   []Addon     `json:"addons"`
	OsType                   string      `json:"os_type"`
	Platform                 string      `json:"platform"`
	NodePortRange            string      `json:"node_port_range"`
	KeyPair                  string      `json:"key_pair"`
	CPUPolicy                string      `json:"cpu_policy"`
	MasterCount              int         `json:"master_count"`
	MasterVswitchIds         []string    `json:"master_vswitch_ids"`
	MasterInstanceTypes      []string    `json:"master_instance_types"`
	MasterSystemDiskCategory string      `json:"master_system_disk_category"`
	MasterSystemDiskSize     int         `json:"master_system_disk_size"`
	Runtime                  RuntimeType `json:"runtime"`
	WorkerInstanceTypes      []string    `json:"worker_instance_types"`
	NumOfNodes               int         `json:"num_of_nodes"`
	WorkerSystemDiskCategory string      `json:"worker_system_disk_category"`
	WorkerSystemDiskSize     int         `json:"worker_system_disk_size"`
	Vpcid                    string      `json:"vpcid"`
	WorkerVswitchIds         []string    `json:"worker_vswitch_ids"`
	ContainerCidr            string      `json:"container_cidr"`
	ServiceCidr              string      `json:"service_cidr"`
}

type Addon struct {
	Name     string `json:"name"`
	Config   string `json:"config"`
	Disabled string `json:"disabled"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RuntimeType struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}