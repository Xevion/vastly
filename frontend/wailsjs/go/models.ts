export namespace api {
	
	export class Offer {
	    is_bid: boolean;
	    inet_up_billed?: number;
	    inet_down_billed?: number;
	    external: boolean;
	    webpage?: string;
	    logo: string;
	    rentable: boolean;
	    compute_cap: number;
	    driver_version: string;
	    cuda_max_good: number;
	    machine_id: number;
	    hosting_type?: number;
	    public_ipaddr: string;
	    geolocation: string;
	    geolocode?: number;
	    flops_per_dphtotal: number;
	    dlperf_per_dphtotal: number;
	    reliability2: number;
	    host_run_time: number;
	    host_id: number;
	    id: number;
	    bundle_id: number;
	    num_gpus: number;
	    total_flops: number;
	    min_bid: number;
	    dph_base: number;
	    dph_total: number;
	    gpu_name: string;
	    gpu_ram: number;
	    gpu_display_active: boolean;
	    gpu_mem_bw: number;
	    bw_nvlink: number;
	    direct_port_count: number;
	    gpu_lanes: number;
	    pcie_bw: number;
	    pci_gen: number;
	    dlperf: number;
	    cpu_name: string;
	    mobo_name: string;
	    cpu_ram: number;
	    cpu_cores: number;
	    cpu_cores_effective: number;
	    gpu_frac: number;
	    has_avx: number;
	    disk_space: number;
	    disk_name: string;
	    disk_bw: number;
	    inet_up: number;
	    inet_down: number;
	    start_date: number;
	    end_date?: number;
	    duration?: number;
	    storage_cost: number;
	    inet_up_cost: number;
	    inet_down_cost: number;
	    storage_total_cost: number;
	    verification: string;
	    score: number;
	    rented: boolean;
	    bundled_results: number;
	    pending_count: number;
	
	    static createFrom(source: any = {}) {
	        return new Offer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.is_bid = source["is_bid"];
	        this.inet_up_billed = source["inet_up_billed"];
	        this.inet_down_billed = source["inet_down_billed"];
	        this.external = source["external"];
	        this.webpage = source["webpage"];
	        this.logo = source["logo"];
	        this.rentable = source["rentable"];
	        this.compute_cap = source["compute_cap"];
	        this.driver_version = source["driver_version"];
	        this.cuda_max_good = source["cuda_max_good"];
	        this.machine_id = source["machine_id"];
	        this.hosting_type = source["hosting_type"];
	        this.public_ipaddr = source["public_ipaddr"];
	        this.geolocation = source["geolocation"];
	        this.geolocode = source["geolocode"];
	        this.flops_per_dphtotal = source["flops_per_dphtotal"];
	        this.dlperf_per_dphtotal = source["dlperf_per_dphtotal"];
	        this.reliability2 = source["reliability2"];
	        this.host_run_time = source["host_run_time"];
	        this.host_id = source["host_id"];
	        this.id = source["id"];
	        this.bundle_id = source["bundle_id"];
	        this.num_gpus = source["num_gpus"];
	        this.total_flops = source["total_flops"];
	        this.min_bid = source["min_bid"];
	        this.dph_base = source["dph_base"];
	        this.dph_total = source["dph_total"];
	        this.gpu_name = source["gpu_name"];
	        this.gpu_ram = source["gpu_ram"];
	        this.gpu_display_active = source["gpu_display_active"];
	        this.gpu_mem_bw = source["gpu_mem_bw"];
	        this.bw_nvlink = source["bw_nvlink"];
	        this.direct_port_count = source["direct_port_count"];
	        this.gpu_lanes = source["gpu_lanes"];
	        this.pcie_bw = source["pcie_bw"];
	        this.pci_gen = source["pci_gen"];
	        this.dlperf = source["dlperf"];
	        this.cpu_name = source["cpu_name"];
	        this.mobo_name = source["mobo_name"];
	        this.cpu_ram = source["cpu_ram"];
	        this.cpu_cores = source["cpu_cores"];
	        this.cpu_cores_effective = source["cpu_cores_effective"];
	        this.gpu_frac = source["gpu_frac"];
	        this.has_avx = source["has_avx"];
	        this.disk_space = source["disk_space"];
	        this.disk_name = source["disk_name"];
	        this.disk_bw = source["disk_bw"];
	        this.inet_up = source["inet_up"];
	        this.inet_down = source["inet_down"];
	        this.start_date = source["start_date"];
	        this.end_date = source["end_date"];
	        this.duration = source["duration"];
	        this.storage_cost = source["storage_cost"];
	        this.inet_up_cost = source["inet_up_cost"];
	        this.inet_down_cost = source["inet_down_cost"];
	        this.storage_total_cost = source["storage_total_cost"];
	        this.verification = source["verification"];
	        this.score = source["score"];
	        this.rented = source["rented"];
	        this.bundled_results = source["bundled_results"];
	        this.pending_count = source["pending_count"];
	    }
	}
	export class ScoredOffer {
	    Offer: Offer;
	    Score: number;
	
	    static createFrom(source: any = {}) {
	        return new ScoredOffer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Offer = this.convertValues(source["Offer"], Offer);
	        this.Score = source["Score"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

