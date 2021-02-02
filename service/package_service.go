package service

import (
	"auto-rev/model"
)

type PackageService interface {
	GetPackageById(id string) (model.Packages, error)
}
