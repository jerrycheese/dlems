package model

import (
	"encoding/json"
	"time"
)

// DRun is a run of Deep Learning
type DRun struct {
	ID        string                 `json:"_id"`
	Name      string                 `json:"name"`      // 名称
	ExecStr   string                 `json:"execStr"`   // 完整执行命令
	Args      map[string]interface{} `json:"args"`      // 程序参数
	MapInfo   map[string]interface{} `json:"mapInfo"`   // 附加信息（键值对）
	Info      string                 `json:"info"`      // 附加信息（文本）
	Device    string                 `json:"device"`    // 使用的设备。"cpu"，"gpu"
	StartTime time.Time              `json:"startTime"` // 创建时间
	EndTime   time.Time              `json:"endTime"`   // 结束时间
}

// MakeDRun return a DRun struct
func MakeDRun(name, execStr string, args map[string]interface{}) DRun {
	return DRun{
		ID:        "",
		Name:      name,
		ExecStr:   execStr,
		Args:      args,
		Device:    "cpu",
		StartTime: time.Now(),
	}
}

// MakeDRunFromMap return a DRun struct according to a map
func MakeDRunFromMap(m map[string]interface{}) (d DRun) {
	if b, err := json.Marshal(m); err == nil {
		json.Unmarshal(b, &d)
	}
	return
}

// AsMap return a map type of DRun
func (d *DRun) AsMap() (m map[string]interface{}) {
	b, err := json.Marshal(d)
	if err != nil {
		return
	}
	json.Unmarshal(b, &m)
	return
}
