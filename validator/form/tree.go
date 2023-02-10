package form

type TreesStruct struct {
	ID           string `form:"id" json:"id" binding:"required"` // tree的id
	Name         string `json:"name"`
	VoltageLevel string `form:"voltageLevel" json:"voltageLevel" binding:"required"`
	Types        string `json:"types"` // 1 是变电站 2 是 变电线 3是台区
	Children     string `json:"children"`
	Status       string `json:"status"`
	LockTag      string `json:"lockTag"`
	ParentId     string `json:"parentId"`
	OrgId        string `json:"orgId"`
}
