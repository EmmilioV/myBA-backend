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
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		ctx.JSON(http.StatusOK, "ok!")
	}
}

func removeEmployee(
	unhireEmployee *employerUsecase.UnhireEmployee,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employeeID := ctx.Params.ByName("id")

		employerID := ctx.Request.Header.Get("employer_id")

		err := unhireEmployee.UseCase(
			ctx.Request.Context(), employerID, employeeID,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		ctx.JSON(http.StatusOK, "ok!")
	}
}

func updateEmployee(
	updateEmployeeInfo *employerUsecase.UpdateEmployeeInfo,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var employee *entity.Employee

		if err := ctx.BindJSON(&employee); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		employerID := ctx.Request.Header.Get("employer_id")

		err := updateEmployeeInfo.UseCase(
			ctx.Request.Context(), employerID, employee,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		ctx.JSON(http.StatusOK, "ok!")
	}
}
