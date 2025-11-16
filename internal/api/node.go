package api

import (
	"time"
)

type Node struct {
	ID           string            `json:"id"`
	Hostname     string            `json:"hostname"`
	CPUCores     int64             `json:"cpu_cores"`
	HostMemoryMB int64             `json:"host_memory_mb"`
	Architecture string            `json:"architecture"`
	Labels       map[string]string `json:"labels"`
}

type Keepalive struct {
	Node      Node      `json:"node"`
	Timestamp time.Time `json:"timestamp"`
	Status    string    `json:"status"`
}
