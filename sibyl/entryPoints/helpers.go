package entryPoints

import (
	"net/http"

	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/core/utils/timeUtils"
	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, err *EndpointError, code int) {
	c.JSON(code, &EndpointResponse{
		Success: false,
		Error:   err,
	})
}

func SendBadGateAway(c *gin.Context, message, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
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
	c.JSON(http.StatusOK, &EndpointResponse{
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
	c.JSON(http.StatusOK, &EndpointResponse{
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
	c.JSON(http.StatusOK, &EndpointResponse{
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
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadGateway,
			Message:   ErrInvalidUserId,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendInvalidPermError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadGateway,
			Message:   ErrInvalidPerm,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendCannotChangePermError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusConflict,
			Message:   ErrCannotChangePerm,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendSamePermError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusConflict,
			Message:   ErrSamePerm,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendNoReasonError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadGateway,
			Message:   ErrNoReason,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendCannotReportYourselfError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadGateway,
			Message:   ErrCannotReportYourself,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendCannotBanYourselfError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadGateway,
			Message:   ErrCannotBanYourself,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendCannotBeReportedError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadGateway,
			Message:   ErrCannotBeReported,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendCannotBeBannedError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadGateway,
			Message:   ErrCannotBeBanned,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendUserNotFoundError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusNotFound,
			Message:   ErrUserNotFound,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendUserNotRegisteredError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusMethodNotAllowed,
			Message:   ErrUserNotRegistered,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendUserNotBannedError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusConflict,
			Message:   ErrUserNotBanned,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendNoDataError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadRequest,
			Message:   ErrNoData,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendBadDataError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadRequest,
			Message:   ErrBadData,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendCannotBeRevokedError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadRequest,
			Message:   ErrCannotBeRevoked,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendTooManyError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadRequest,
			Message:   ErrTooManyUsers,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendNoMessageError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusConflict,
			Message:   ErrNoMessage,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendNoSourceError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusConflict,
			Message:   ErrNoSource,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendUserAlreadyBannedError(c *gin.Context, origin string,
	u *sibylValues.User, banReason, banMsg, srcUrl string) {
	str := "reasons: [" + u.Reason + "-" + banReason + "] | "
	str += "messages: [" + u.Message + "-" + banMsg + "] | "
	str += "urls: [" + u.BanSourceUrl + "-" + srcUrl + "]"
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusAccepted,
			Message:   ErrUserAlreadyBanned + str,
			Origin:    origin,
			Date:      timeUtils.GenerateCurrentDateTime(),
		},
	})
}

func SendPermissionDenied(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
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
