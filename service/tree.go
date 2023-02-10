package service

import (
	"fiber-nuzn-api/validator/form"
)

type TreeService struct {
}

func NewTreeService() *TreeService {
	return &TreeService{}
}

func (r *TreeService) GetList(TreesStruct form.TreesStruct) ([]form.TreesStruct, error) {
	var result []form.TreesStruct
	switch TreesStruct.Types {
	case "1":
		tree1 := []form.TreesStruct{
			{
				ID:           "1",
				Name:         "张三",
				VoltageLevel: "",
				Types:        "1",
				Children:     "",
				Status:       "",
				LockTag:      "",
				ParentId:     "",
				OrgId:        "",
			},
		}
		result = tree1
	case "2":
		tree2 := []form.TreesStruct{
			{
				ID:           "10",
				Name:         "关羽",
				VoltageLevel: "",
				Types:        "2",
				Children:     "",
				Status:       "",
				LockTag:      "",
				ParentId:     "",
				OrgId:        "",
			},
			{
				ID:           "11",
				Name:         "这歌",
				VoltageLevel: "",
				Types:        "2",
				Children:     "",
				Status:       "",
				LockTag:      "",
				ParentId:     "",
				OrgId:        "",
			},
			{
				ID:           "12",
				Name:         "破解",
				VoltageLevel: "",
				Types:        "2",
				Children:     "",
				Status:       "",
				LockTag:      "",
				ParentId:     "",
				OrgId:        "",
			},
		}
		result = tree2
	case "3":
		tree3 := []form.TreesStruct{
			{
				ID:           "10" + "00",
				Name:         "关羽" + TreesStruct.ID,
				VoltageLevel: "",
				Types:        "2",
				Children:     "",
				Status:       "",
				LockTag:      "",
				ParentId:     "",
				OrgId:        "",
			},
			{
				ID:           "11" + "00",
				Name:         "这歌" + TreesStruct.ID,
				VoltageLevel: "",
				Types:        "2",
				Children:     "",
				Status:       "",
				LockTag:      "",
				ParentId:     "",
				OrgId:        "",
			},
			{
				ID:           "12" + "00",
				Name:         "破解" + TreesStruct.ID,
				VoltageLevel: "",
				Types:        "2",
				Children:     "",
				Status:       "",
				LockTag:      "",
				ParentId:     "",
				OrgId:        "",
			},
		}
		result = tree3
	}
	return result, nil
}
