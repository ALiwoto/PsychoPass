package entryPoints

import (
	"net/http"

	"github.com/AnimeKaizoku/sibylapi-go/sibyl/core/utils/timeUtils"
	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, err *EndpointError, code int) {
	c.JSON(code, &EndpointResponse{
		Success: false,
		Error:   err,
	})
}

func SendBadGateAway(c *gin.Context, message, origin string) {
	c.JSON(http.StatusBadGateway, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadGateway,
			Message:   message,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendNoTokenError(c *gin.Context, origin string) {
	c.JSON(http.StatusUnauthorized, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusUnauthorized,
			Message:   ErrNoToken,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendInvalidTokenError(c *gin.Context, origin string) {
	c.JSON(http.StatusBadGateway, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadGateway,
			Message:   ErrInvalidToken,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendInternalServerError(c *gin.Context, origin string) {
	c.JSON(http.StatusInternalServerError, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusInternalServerError,
			Message:   ErrInternalServerError,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendInvalidUserIdError(c *gin.Context, origin string) {
	c.JSON(http.StatusBadGateway, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadGateway,
			Message:   ErrInvalidUserId,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendUserNotFoundError(c *gin.Context, origin string) {
	c.JSON(http.StatusNotFound, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusNotFound,
			Message:   ErrUserNotFound,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendUserNotBannedError(c *gin.Context, origin string) {
	c.JSON(http.StatusConflict, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusConflict,
			Message:   ErrUserNotBanned,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendUserAlreadyBannedError(c *gin.Context, origin string) {
	c.JSON(http.StatusAccepted, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusAccepted,
			Message:   ErrUserAlreadyBanned,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendPermissionDenied(c *gin.Context, origin string) {
	c.JSON(http.StatusBadGateway, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadGateway,
			Message:   ErrPermissionDenied,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendResult(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: true,
		Result:  result,
	})
}
