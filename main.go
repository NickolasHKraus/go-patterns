package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type OpenCostAPIResponse struct {
	Code   int                                  `json:"code"`
	Status string                               `json:"status"`
	Data   []map[string]OpenCostAPIResponseData `json:"data"`
}

type OpenCostAPIResponseData struct {
	Name                       string                 `json:"name"`
	Properties                 map[string]interface{} `json:"properties"`
	Window                     map[string]interface{} `json:"window"`
	Start                      string                 `json:"start"`
	End                        string                 `json:"end"`
	Minutes                    float64                `json:"minutes"`
	CpuCores                   float64                `json:"cpuCores"`
	CpuCoreRequestAverage      float64                `json:"cpuCoreRequestAverage"`
	CpuCoreUsageAverage        float64                `json:"cpuCoreUsageAverage"`
	CpuCoreHours               float64                `json:"cpuCoreHours"`
	CpuCost                    float64                `json:"cpuCost"`
	CpuCostAdjustment          float64                `json:"cpuCostAdjustment"`
	CpuEfficiency              float64                `json:"cpuEfficiency"`
	GpuCount                   float64                `json:"gpuCount"`
	GpuHours                   float64                `json:"gpuHours"`
	GpuCost                    float64                `json:"gpuCost"`
	GpuCostAdjustment          float64                `json:"gpuCostAdjustment"`
	NetworkTransferBytes       float64                `json:"networkTransferBytes"`
	NetworkReceiveBytes        float64                `json:"networkReceiveBytes"`
	NetworkCost                float64                `json:"networkCost"`
	NetworkCostAdjustment      float64                `json:"networkCostAdjustment"`
	LoadBalancerCost           float64                `json:"loadBalancerCost"`
	LoadBalancerCostAdjustment float64                `json:"loadBalancerCostAdjustment"`
	PvBytes                    float64                `json:"pvBytes"`
	PvByteHours                float64                `json:"pvByteHours"`
	PvCost                     float64                `json:"pvCost"`
	Pvs                        float64                `json:"pvs"`
	PvCostAdjustment           float64                `json:"pvCostAdjustment"`
	RamBytes                   float64                `json:"ramBytes"`
	RamByteRequestAverage      float64                `json:"ramByteRequestAverage"`
	RamByteUsageAverage        float64                `json:"ramByteUsageAverage"`
	RamByteHours               float64                `json:"ramByteHours"`
	RamCost                    float64                `json:"ramCost"`
	RamCostAdjustment          float64                `json:"ramCostAdjustment"`
	RamEfficiency              float64                `json:"ramEfficiency"`
	SharedCost                 float64                `json:"sharedCost"`
	ExternalCost               float64                `json:"externalCost"`
	TotalCost                  float64                `json:"totalCost"`
	TotalEfficiency            float64                `json:"totalEfficiency"`
	RawAllocationOnly          float64                `json:"rawAllocationOnly"`
}

func main() {
	body := `{
  "code": 200,
  "status": "success",
  "data": [
    {
      "cluster-one": {
        "name": "cluster-one",
        "properties": {
          "cluster": "cluster-one",
          "node": "minikube"
        },
        "window": {
          "start": "2022-10-24T16:04:11Z",
          "end": "2022-10-24T16:05:11Z"
        },
        "start": "2022-10-24T16:04:11Z",
        "end": "2022-10-24T16:05:00Z",
        "minutes": 0.803359,
        "cpuCores": 0.76,
        "cpuCoreRequestAverage": 0.76,
        "cpuCoreUsageAverage": 0,
        "cpuCoreHours": 0.010176,
        "cpuCost": 0.000322,
        "cpuCostAdjustment": 0,
        "cpuEfficiency": 0,
        "gpuCount": 0,
        "gpuHours": 0,
        "gpuCost": 0,
        "gpuCostAdjustment": 0,
        "networkTransferBytes": 0,
        "networkReceiveBytes": 0,
        "networkCost": 0,
        "networkCostAdjustment": 0,
        "loadBalancerCost": 0,
        "loadBalancerCostAdjustment": 0,
        "pvBytes": 8589934592,
        "pvByteHours": 115013300.610935,
        "pvCost": 6e-06,
        "pvs": null,
        "pvCostAdjustment": 0,
        "ramBytes": 233257920,
        "ramByteRequestAverage": 233257920,
        "ramByteUsageAverage": 0,
        "ramByteHours": 3123162.695305,
        "ramCost": 1.2e-05,
        "ramCostAdjustment": 0,
        "ramEfficiency": 0,
        "sharedCost": 0,
        "externalCost": 0,
        "totalCost": 0.00034,
        "totalEfficiency": 0,
        "rawAllocationOnly": null
      }
    }
  ]
}`
	var o OpenCostAPIResponse
	if err := json.Unmarshal([]byte(body), &o); err != nil {
		fmt.Printf("%s\n", err)
	}

	d := OpenCostAPIResponseData{
		CpuCores: 1337.0,
	}
	v := reflect.ValueOf(d)
	fmt.Println("value:", v.FieldByName("CpuCorex").Float())
}
