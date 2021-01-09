package model

import (
	"encoding/json"
	"time"
)

// DValue is a metric/loss value of Deep Learning
type DValue struct {
	RunID   string                 `json:"runID"`   // DRun的ID
	Name    string                 `json:"name"`    // 值名称
	Value   float64                `json:"value"`   // 值
	MapInfo map[string]interface{} `json:"mapInfo"` // 附加信息（键值对）
	Info    string                 `json:"info"`    // 附加信息（文本）
	Time    time.Time              `json:"time"`    // 创建时间
}

// MakeDValue return a DValue struct
func MakeDValue(RunID, name string, value float64) DValue {
	return DValue{
		RunID: RunID,
		Name:  name,
		Value: value,
		Time:  time.Now(),
	}
}

// MakeDValueFromMap return a DValue struct according to a map
func MakeDValueFromMap(m map[string]interface{}) (d DValue) {
	if b, err := json.Marshal(m); err == nil {
		json.Unmarshal(b, &d)
	}
	return
}

// AsMap return a map type of DRun
func (d *DValue) AsMap() (m map[string]interface{}) {
	b, err := json.Marshal(d)
	if err != nil {
		return
	}
	json.Unmarshal(b, &m)
	return
}
