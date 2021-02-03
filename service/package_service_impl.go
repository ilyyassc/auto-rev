package service

import (
	"auto-rev/config"
	"auto-rev/dao"
	"auto-rev/model"
)

var packageDao dao.PackageDao = dao.PackageDaoImpl{}

type PackageServiceImpl struct{}

func (PackageServiceImpl) GetPackageById(id string) (p model.Packages, e error) {
	defer config.CatchError(&e)
	
	p, err := packageDao.GetPackageById(id)
	return p, err
}
