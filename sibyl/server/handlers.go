package server

import (
	"strings"

	"github.com/AnimeKaizoku/sibylapi-go/sibyl/entryPoints/banHandlers"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/entryPoints/infoHandlers"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/entryPoints/reportHandlers"
	"github.com/AnimeKaizoku/sibylapi-go/sibyl/entryPoints/tokenHandlers"
	"github.com/gin-gonic/gin"
)

func LoadHandlers() {
	// create token handlers
	bindHandler(tokenHandlers.CreateTokenHandler, "create",
		"createToken", "generate")

	// revoke token handlers
	bindHandler(tokenHandlers.RevokeTokenHandler, "revoke", "revokeToken")

	// change token permission handlers
	bindHandler(tokenHandlers.ChangeTokenPermHandler, "changePerm", "promote")

	// get token handlers
	bindHandler(tokenHandlers.GetTokenHandler, "getToken")

	// addBan handlers
	bindHandler(banHandlers.AddBanHandler, "addBan", "ban", "banUser")

	// deleteBan handlers
	bindHandler(banHandlers.RemoveBanHandler, "deleteBan", "removeBan",
		"revertBan", "remBan")

	// getInfo handlers
	bindHandler(infoHandlers.GetInfoHandler, "getInfo", "fetchInfo")

	// checkToken handlers
	bindHandler(infoHandlers.CheckTokenHandler, "checkToken")

	// report handlers
	bindHandler(reportHandlers.ReportUserHandler, "report", "reportUser")

	bindNoRoot()
}

func bindHandler(handler gin.HandlerFunc, paths ...string) {
	for _, path := range paths {
		ServerEngine.GET(path, handler)
		ServerEngine.POST(path, handler)
	}
}

func bindNoRoot() {
	ServerEngine.NoRoute(noRootHandler)
}

func noRootHandler(c *gin.Context) {
	path := strings.TrimSpace(c.Request.URL.Path)
	path = strings.Trim(path, "/")
	path = strings.ToLower(path)
	switch path {
	case "create", "createtoken", "generate":
		tokenHandlers.CreateTokenHandler(c)
	case "revoke", "revoketoken":
		tokenHandlers.RevokeTokenHandler(c)
	case "changeperm", "promote":
		tokenHandlers.ChangeTokenPermHandler(c)
	case "gettoken":
		banHandlers.AddBanHandler(c)
	case "addban", "ban", "banuser":
		banHandlers.AddBanHandler(c)
	case "deleteban", "removeban", "revertban", "remban":
		banHandlers.RemoveBanHandler(c)
	case "getinfo", "fetchinfo":
		infoHandlers.GetInfoHandler(c)
	case "checktoken":
		infoHandlers.CheckTokenHandler(c)
	case "report", "reportuser":
		reportHandlers.ReportUserHandler(c)
	default:
		// TODO: send docs or redirect to docs or something else.
		return
	}
}
