package model

type Models struct {
	GroupModel          GroupModel
	OfflineMessageModel OfflineMessageModel
	// 其他需要的 model
}

func NewModels(group GroupModel, offline OfflineMessageModel) *Models {
	return &Models{
		GroupModel:          group,
		OfflineMessageModel: offline,
	}
}
