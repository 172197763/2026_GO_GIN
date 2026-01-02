package daos

import (
	"gin_test/api/app/models"
	"gin_test/api/database"

	"xorm.io/xorm"
)

type store_supplier_dao struct {
	Session *xorm.Session
}

func Newstore_supplier_dao() *store_supplier_dao {
	return &store_supplier_dao{Session: database.GetDbSession()}
}

// 查找列表
func (dao *store_supplier_dao) GetList(model *[]models.StoreSupplier) {
	err := dao.Session.Find(model)
	if err != nil {
		panic(err)
	}
}

// 批量更新
func (dao *store_supplier_dao) UpdateBatch(model *[]models.StoreSupplier) {
	for _, item := range *model {
		_, err := dao.Session.ID(item.Id).Update(item)
		if err != nil {
			// 处理错误
			panic(err)
		}
	}
}
