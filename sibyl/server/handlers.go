package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/entryPoints/banHandlers"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/entryPoints/infoHandlers"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/entryPoints/reportHandlers"
	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/entryPoints/tokenHandlers"
)

func LoadHandlers() {
	// create token handlers
	bindHandler(tokenHandlers.CreateTokenHandler, "create",
		"createToken", "generate")

	// revoke token handlers
	bindHandler(tokenHandlers.RevokeTokenHandler,
		"revoke", "revokeToken")

	// get token handlers
	bindHandler(tokenHandlers.GetTokenHandler, "getToken")

	// addBan handlers
	bindHandler(banHandlers.AddBanHandler, "addBan")

	// deleteBan handlers
	bindHandler(banHandlers.RemoveBanHandler, "deleteBan", "removeBan",
		"revertBan")

	// getInfo handlers
	bindHandler(infoHandlers.GetInfoHandler, "getInfo", "fetchInfo")

	// report handlers
	bindHandler(reportHandlers.ReportUserHandler, "report", "reportUser")

	// get update handlers
	bindHandler(reportHandlers.GetUpdateHandler, "getUpdate", "fetchUpdate")
}

func bindHandler(handler gin.HandlerFunc, paths ...string) {
	for _, path := range paths {
		ServerEngine.GET(path, handler)
		ServerEngine.POST(path, handler)
	}
}
