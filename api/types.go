package api

type PortMapping struct {
	HostIp   string `json:"HostIp"`
	HostPort string `json:"HostPort"`
}

type Ports struct {
	TCP22   []PortMapping `json:"22/tcp"`
	TCP8080 []PortMapping `json:"8080/tcp"`
	UDP8080 []PortMapping `json:"8080/udp"`
}

type Instance struct {
	IsBid             bool     `json:"is_bid"`
	InetUpBilled      float64  `json:"inet_up_billed"`
	InetDownBilled    float64  `json:"inet_down_billed"`
	External          bool     `json:"external"`
	Webpage           *string  `json:"webpage"`
	Logo              string   `json:"logo"`
	Rentable          bool     `json:"rentable"`
	ComputeCap        int      `json:"compute_cap"`
	DriverVersion     string   `json:"driver_version"`
	CudaMaxGood       int      `json:"cuda_max_good"`
	MachineID         int      `json:"machine_id"`
	HostingType       *string  `json:"hosting_type"`
	PublicIPAddr      string   `json:"public_ipaddr"`
	Geolocation       string   `json:"geolocation"`
	FlopsPerDPHTotal  float64  `json:"flops_per_dphtotal"`
	DLPerfPerDPHTotal float64  `json:"dlperf_per_dphtotal"`
	Reliability2      float64  `json:"reliability2"`
	HostRunTime       int64    `json:"host_run_time"`
	HostID            int      `json:"host_id"`
	ID                int      `json:"id"`
	BundleID          int      `json:"bundle_id"`
	NumGPUs           int      `json:"num_gpus"`
	TotalFlops        float64  `json:"total_flops"`
	MinBid            float64  `json:"min_bid"`
	DPHBase           float64  `json:"dph_base"`
	DPHTotal          float64  `json:"dph_total"`
	GPUName           string   `json:"gpu_name"`
	GPURam            int      `json:"gpu_ram"`
	GPUDisplayActive  bool     `json:"gpu_display_active"`
	GPUMemBW          float64  `json:"gpu_mem_bw"`
	BWNVLink          int      `json:"bw_nvlink"`
	DirectPortCount   int      `json:"direct_port_count"`
	GPULanes          int      `json:"gpu_lanes"`
	PCIeBW            float64  `json:"pcie_bw"`
	PCIGen            int      `json:"pci_gen"`
	DLPerf            float64  `json:"dlperf"`
	CPUName           string   `json:"cpu_name"`
	MoboName          string   `json:"mobo_name"`
	CPURam            int      `json:"cpu_ram"`
	CPUCores          int      `json:"cpu_cores"`
	CPUCoresEffective float64  `json:"cpu_cores_effective"`
	GPUFrac           float64  `json:"gpu_frac"`
	HasAVX            int      `json:"has_avx"`
	DiskSpace         float64  `json:"disk_space"`
	DiskName          string   `json:"disk_name"`
	DiskBW            float64  `json:"disk_bw"`
	InetUp            float64  `json:"inet_up"`
	InetDown          float64  `json:"inet_down"`
	StartDate         float64  `json:"start_date"`
	EndDate           int64    `json:"end_date"`
	Duration          float64  `json:"duration"`
	StorageCost       float64  `json:"storage_cost"`
	InetUpCost        float64  `json:"inet_up_cost"`
	InetDownCost      float64  `json:"inet_down_cost"`
	StorageTotalCost  float64  `json:"storage_total_cost"`
	Verification      string   `json:"verification"`
	Score             float64  `json:"score"`
	SSHIdx            string   `json:"ssh_idx"`
	SSHHost           string   `json:"ssh_host"`
	SSHPort           int      `json:"ssh_port"`
	ActualStatus      string   `json:"actual_status"`
	IntendedStatus    string   `json:"intended_status"`
	CurState          string   `json:"cur_state"`
	NextState         string   `json:"next_state"`
	ImageUUID         string   `json:"image_uuid"`
	ImageArgs         []string `json:"image_args"`
	ImageRuntype      string   `json:"image_runtype"`
	ExtraEnv          string   `json:"extra_env"`
	OnStart           string   `json:"onstart"`
	Label             *string  `json:"label"`
	JupyterToken      string   `json:"jupyter_token"`
	StatusMsg         string   `json:"status_msg"`
	GPUUtil           float64  `json:"gpu_util"`
	DiskUtil          float64  `json:"disk_util"`
	GPUTemp           float64  `json:"gpu_temp"`
	LocalIPAddrs      string   `json:"local_ipaddrs"`
	DirectPortEnd     int      `json:"direct_port_end"`
	DirectPortStart   int      `json:"direct_port_start"`
	CPUUtil           float64  `json:"cpu_util"`
	MemUsage          float64  `json:"mem_usage"`
	MemLimit          float64  `json:"mem_limit"`
	VMemUsage         float64  `json:"vmem_usage"`
	MachineDirSSHPort int      `json:"machine_dir_ssh_port"`
	Ports             Ports    `json:"ports"`
}
