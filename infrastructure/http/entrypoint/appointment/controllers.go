package appointment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appointmentUseCase "go.mod/domain/appointment/usecase"
	serviceEntity "go.mod/domain/service/entity"
)

func addService(
	addService *appointmentUseCase.AddService,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var service *serviceEntity.Service

		if err := ctx.BindJSON(&service); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		employerID := ctx.Request.Header.Get("employer_id")

		err := addService.UseCase(
			ctx.Request.Context(), service, employerID,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		ctx.JSON(http.StatusOK, "ok!")
	}
}
