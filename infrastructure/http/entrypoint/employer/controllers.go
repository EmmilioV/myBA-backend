package employer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mod/domain/employee/entity"
	employerUsecase "go.mod/domain/employer/usecase"
)

func registerEmployee(
	hireEmployee *employerUsecase.HireEmployee,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var employee *entity.Employee

		if err := ctx.BindJSON(&employee); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		employerID := ctx.Request.Header.Get("employer_id")

		err := hireEmployee.UseCase(
			ctx.Request.Context(), employerID, employee,
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, "ok!")
	}
}
