package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appointmentEntity "go.mod/domain/appointment/entity"
	customerUsecase "go.mod/domain/customer/usecase"
)

func scheduleAnAppointment(
	scheduleAnAppointment *customerUsecase.ScheduleAnAppointment,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var appointment *appointmentEntity.Appoinment

		if err := ctx.BindJSON(&appointment); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		err := scheduleAnAppointment.UseCase(
			ctx.Request.Context(), appointment,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		ctx.JSON(http.StatusOK, "ok!")
	}
}
