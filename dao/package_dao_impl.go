package dao

import (
	"auto-rev/config"
	"auto-rev/model"
)

type PackageDaoImpl struct{}

func (PackageDaoImpl) GetPackageById(id string) (u model.Packages, e error) {
	defer config.CatchError(&e)
	data := model.Packages{BaseModel: model.BaseModel{Id: id}}
	result := g.First(&data)

	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil
	}
	return data, result.Error
}
