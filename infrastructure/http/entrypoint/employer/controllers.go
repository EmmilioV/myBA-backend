package employer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	customerEntity "go.mod/domain/customer/entity"
	employeeEntity "go.mod/domain/employee/entity"
	employerUsecase "go.mod/domain/employer/usecase"
)

func registerEmployee(
	hireEmployee *employerUsecase.HireEmployee,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var employee *employeeEntity.Employee

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
		var employee *employeeEntity.Employee

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

func registerCustomer(
	registerCustomer *employerUsecase.RegisterCustomer,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var customer *customerEntity.Customer

		if err := ctx.BindJSON(&customer); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		employerID := ctx.Request.Header.Get("employer_id")

		err := registerCustomer.UseCase(
			ctx.Request.Context(), employerID, customer,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		ctx.JSON(http.StatusOK, "ok!")
	}
}

func searchWithAppointmentsByID(
	searchWithAppointmentsByID *employerUsecase.SearchWithAppointmentsByID,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employerID := ctx.Params.ByName("id")

		response, err := searchWithAppointmentsByID.UseCase(
			ctx.Request.Context(), employerID,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}
