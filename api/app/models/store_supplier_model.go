package models

type StoreSupplier struct {
	Id       int    `json:"id" xorm:"'id' pk autoincr"`
	Name     string `json:"name"`
	Sort_num int    `json:"sort_num"`
}

func (*StoreSupplier) TableName() string {
	return "t_store_supplier"
}
