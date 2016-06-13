package ec2instances

type Region string

type Instance struct {
	Family                   string   `json:"family"`
	EnhancedNetworking       bool     `json:"enhanced_networking"`
	VCPU                     int      `json:"vCPU"`
	Generation               string   `json:"generation"`
	EbsIOPS                  int      `json:"ebs_iops"`
	NetworkPerformance       string   `json:"network_performance"`
	EbsThroughput            int      `json:"ebs_throughput"`
	PrettyName               string   `json:"pretty_name"`
	InstanceType             string   `json:"instance_type"`
	ECU                      float64  `json:"ECU"`
	Memory                   float64  `json:"memory"`
	EbsMaxBandwidth          int      `json:"ebs_max_bandwidth"`
	VPC                      VPC      `json:"vpc"`
	Arch                     []string `json:"arch"`
	VPCOnly                  bool     `json:"vpc_only"`
	LinuxVirtualizationTypes []string `json:"linux_virtualization_types"`
	EBSOptimized             bool     `json:"ebs_optimized"`
	Storage                  Storage  `json:"storage"`

	Pricing map[Region]RegionPricing `json:"pricing"`
}

type RegionPricing struct {
	Linux         Pricing `json:"linux"`
	Windows       Pricing `json:"mswin"`
	WindowsSQL    Pricing `json:"mswinSQL"`
	WindowsSQLWeb Pricing `json:"mswinSQLWeb"`
}

type Pricing struct {
	Reserved ReservedPricing `json:"reserved"`
	OnDemand string          `json:"ondemand"`
}

type ReservedPricing struct {
	YearTerm1NoUpFront      string `json:"yrTerm1.noUpfront"`
	YearTerm1PartialUpFront string `json:"yrTerm1.partialUpfront"`
	YearTerm1AllUpFront     string `json:"yrTerm1.allUpfront"`
	YearTerm3NoUpFront      string `json:"yrTerm3.allUpfront"`
	YearTerm3PartialUpFront string `json:"yrTerm3.partialUpfront"`
}

type VPC struct {
	IPsPerENI int `json:"ips_per_eni"`
	MaxENIs   int `json:"max_enis"`
}

type Storage struct {
	SSD     bool    `json:"ssd"`
	Devices int     `json:"devices"`
	Size    float64 `json:"size"`
}
