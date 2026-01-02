package user

import (
	"gin_test/api/app/daos"
	"gin_test/api/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *User) UpdateBatch(ctx *gin.Context) {
	dao := daos.Newstore_supplier_dao()
	data2 := []models.StoreSupplier{}
	dao.GetList(&data2)
	for k, v := range data2 {
		data2[k].Sort_num = v.Sort_num + 1
	}
	dao.UpdateBatch(&data2)
	ctx.JSON(http.StatusOK, gin.H{"data2": data2})
}
