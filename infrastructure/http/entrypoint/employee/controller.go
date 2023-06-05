package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mod/domain/employee/usecase"
)

func getEmployeeWithServicesInfo(
	searchByIDWithServices *usecase.SearchByIDWithServices,
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
