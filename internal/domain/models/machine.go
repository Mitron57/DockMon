package models

type Machine struct {
    IP        string `json:"ip"`
    PingTime  int32  `json:"ping_time"`
    LastCheck Time   `json:"last_check"`
}
