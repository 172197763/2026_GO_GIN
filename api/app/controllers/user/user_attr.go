package user

import (
	"gin_test/api/app/daos"
	"gin_test/api/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *User) MyName(ctx *gin.Context) {
	data := "i am tom"
	dao := daos.Newstore_supplier_dao()
	data2 := []models.StoreSupplier{}
	dao.GetList(&data2)
	ctx.JSON(http.StatusOK, gin.H{"data": data, "data2": data2})
}
func (u *User) MySex(ctx *gin.Context) {
	data := "i am boy"
	ctx.JSON(http.StatusOK, gin.H{"name": data})
}
