package api

import (
	"encoding/json"
	"net/http"
)

type ComparableInteger struct {
	eq *int `json:"eq,omitempty"`
	lt *int `json:"lt,omitempty"`
	le *int `json:"le,omitempty"`
	gt *int `json:"gt,omitempty"`
	ge *int `json:"ge,omitempty"`
}

type ComparableFloat struct {
	eq *float64 `json:"eq,omitempty"`
	lt *float64 `json:"lt,omitempty"`
	le *float64 `json:"le,omitempty"`
	gt *float64 `json:"gt,omitempty"`
	ge *float64 `json:"ge,omitempty"`
}

type AdvancedSearch struct {
	Verified          *bool              `json:"verified,omitempty"`
	ComputeCap        *ComparableInteger `json:"compute_cap,omitempty"`
	DiskSpace         *ComparableInteger `json:"disk_space,omitempty"`
	Order             []string           `json:"order,omitempty"`
	Type              string             `json:"type,omitempty"`
	BwNVLink          *ComparableInteger `json:"bw_nvlink,omitempty"`
	CPUCores          *ComparableInteger `json:"cpu_cores,omitempty"`
	CPUCoresEffective *ComparableInteger `json:"cpu_cores_effective,omitempty"`
	CPURam            *ComparableInteger `json:"cpu_ram,omitempty"`
	CudaVers          *ComparableInteger `json:"cuda_vers,omitempty"`
	DirectPortCount   *ComparableInteger `json:"direct_port_count,omitempty"`
	DiskBw            *ComparableFloat   `json:"disk_bw,omitempty"`
	DLPerf            *ComparableFloat   `json:"dlperf,omitempty"`
	DLPerfUSD         *ComparableFloat   `json:"dlperf_usd,omitempty"`
	DPH               *ComparableFloat   `json:"dph,omitempty"`
	DriverVersion     *string            `json:"driver_version,omitempty"`
	Duration          *ComparableFloat   `json:"duration,omitempty"`
	External          *bool              `json:"external,omitempty"`
	FlopsUSD          *ComparableFloat   `json:"flops_usd,omitempty"`
	GPUMemBw          *ComparableFloat   `json:"gpu_mem_bw,omitempty"`
	GPUName           *string            `json:"gpu_name,omitempty"`
	GPURam            *ComparableInteger `json:"gpu_ram,omitempty"`
	GPUFrac           *ComparableFloat   `json:"gpu_frac,omitempty"`
	HasAVX            *bool              `json:"has_avx,omitempty"`
	ID                *ComparableInteger `json:"id,omitempty"`
	InetDown          *ComparableFloat   `json:"inet_down,omitempty"`
	InetDownCost      *ComparableFloat   `json:"inet_down_cost,omitempty"`
	InetUp            *ComparableFloat   `json:"inet_up,omitempty"`
	InetUpCost        *ComparableFloat   `json:"inet_up_cost,omitempty"`
	MachineID         *ComparableInteger `json:"machine_id,omitempty"`
	MinBid            *ComparableFloat   `json:"min_bid,omitempty"`
	NumGPUs           *ComparableInteger `json:"num_gpus,omitempty"`
	PCIGen            *ComparableInteger `json:"pci_gen,omitempty"`
	PCIeBw            *ComparableFloat   `json:"pcie_bw,omitempty"`
	Reliability       *ComparableFloat   `json:"reliability,omitempty"`
	Rentable          *bool              `json:"rentable,omitempty"`
	Rented            *bool              `json:"rented,omitempty"`
	StorageCost       *ComparableFloat   `json:"storage_cost,omitempty"`
	TotalFlops        *ComparableFloat   `json:"total_flops,omitempty"`
}

func NewSearch() *AdvancedSearch {
	return &AdvancedSearch{
		Rented: Pointer(false),
	}
}

type Offer struct {
	IsBid             bool     `json:"is_bid"`
	InetUpBilled      *float64 `json:"inet_up_billed"`
	InetDownBilled    *float64 `json:"inet_down_billed"`
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
	HostRunTime       int      `json:"host_run_time"`
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
	GPUMemBw          float64  `json:"gpu_mem_bw"`
	BwNVLink          int      `json:"bw_nvlink"`
	DirectPortCount   int      `json:"direct_port_count"`
	GPULanes          int      `json:"gpu_lanes"`
	PCIeBw            float64  `json:"pcie_bw"`
	PCIGen            int      `json:"pci_gen"`
	DLPerf            float64  `json:"dlperf"`
	CPUName           string   `json:"cpu_name"`
	MoboName          string   `json:"mobo_name"`
	CPURam            int      `json:"cpu_ram"`
	CPUCores          int      `json:"cpu_cores"`
	CPUCoresEffective int      `json:"cpu_cores_effective"`
	GPUFrac           float64  `json:"gpu_frac"`
	HasAVX            int      `json:"has_avx"`
	DiskSpace         float64  `json:"disk_space"`
	DiskName          string   `json:"disk_name"`
	DiskBw            float64  `json:"disk_bw"`
	InetUp            float64  `json:"inet_up"`
	InetDown          float64  `json:"inet_down"`
	StartDate         float64  `json:"start_date"`
	EndDate           *float64 `json:"end_date"`
	Duration          *float64 `json:"duration"`
	StorageCost       float64  `json:"storage_cost"`
	InetUpCost        float64  `json:"inet_up_cost"`
	InetDownCost      float64  `json:"inet_down_cost"`
	StorageTotalCost  float64  `json:"storage_total_cost"`
	Verification      string   `json:"verification"`
	Score             float64  `json:"score"`
	Rented            bool     `json:"rented"`
	BundledResults    int      `json:"bundled_results"`
	PendingCount      int      `json:"pending_count"`
}

type SearchResponse struct {
	Offers []Offer `json:"offers"`
}

func (c *Client) Search(search *AdvancedSearch) (*SearchResponse, error) {
	resp, err := c.makeRequest(http.MethodPost, "/bundles/", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
