package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
	employeeUsecase "go.mod/domain/employee/usecase"
	serviceEntity "go.mod/domain/service/entity"
)

func getEmployeeWithServicesInfo(
	searchByIDWithServices *employeeUsecase.SearchByIDWithServices,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employeeID := ctx.Params.ByName("id")

		response, err := searchByIDWithServices.UseCase(
			ctx.Request.Context(), employeeID,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}

func updateEmployeeService(
	updateServiceInfo *employeeUsecase.UpdateServiceInfo,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var service *serviceEntity.Service

		if err := ctx.BindJSON(&service); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		employeeID := ctx.Request.Header.Get("user_id")

		err := updateServiceInfo.UseCase(
			ctx.Request.Context(), employeeID, service,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		ctx.JSON(http.StatusOK, "ok!")
	}
}
