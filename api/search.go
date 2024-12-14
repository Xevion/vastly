package api

import (
	"encoding/json"
	"net/http"
)

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

type SearchResponse struct {
	Offers []Offer `json:"offers"`
}

func (c *Client) Search(search *AdvancedSearch) (*SearchResponse, error) {
	resp, err := c.makeRequest(http.MethodPost, "/bundles/", search)
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
