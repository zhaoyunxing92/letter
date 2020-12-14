package domain

//基础的数据
type base struct {
	Status   uint8 `json:"status"`   //应用状态
	Version  uint  `json:"version"`  //数据版本
	CreateIn int64 `json:"createIn"` //创建时间
	UpdateIn int64 `json:"updateIn"` //创建时间
}
