package models

type Machine struct {
	IP          string   `json:"ip"`
	PingTime    Duration `json:"ping_time"`
	Success     bool     `json:"success"`
	LastSuccess Time     `json:"last_success"`
}
