package dao

import (
	"auto-rev/model"
)

type PackageDao interface {
	GetPackageById(id string) (model.Packages, error)
}
