package server

import (
	"strings"

	"github.com/AnimeKaizoku/PsychoPass/docs/in"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/entryPoints/banHandlers"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/entryPoints/infoHandlers"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/entryPoints/reportHandlers"
	"github.com/AnimeKaizoku/PsychoPass/sibyl/entryPoints/tokenHandlers"
	"github.com/gin-gonic/gin"
)

func LoadHandlers() {
	// documentation
	in.LoadDocs(ServerEngine)
	// create token handlers
	bindHandler(tokenHandlers.CreateTokenHandler, "create",
		"createToken", "generate")

	// revoke token handlers
	bindHandler(tokenHandlers.RevokeTokenHandler, "revoke", "revokeToken")

	// change token permission handlers
	bindHandler(tokenHandlers.ChangeTokenPermHandler, "changePerm", "promote")

	// get token handlers
	bindHandler(tokenHandlers.GetTokenHandler, "getToken")

	// get all register handlers
	bindHandler(tokenHandlers.GetAllRegisteredUsersHandler, "getRegistered",
		"registeredUsers", "getAllregistered")

	// addBan handlers
	bindHandler(banHandlers.AddBanHandler, "addBan", "ban", "banUser")

	// multiBan handlers
	bindPostHandler(banHandlers.MultiBanHandler, "multiBan", "addMultiBan")

	// multiUnBan handlers
	bindPostHandler(banHandlers.MultiUnBanHandler, "multiUnBan", "removeMultiBan")

	// deleteBan handlers
	bindHandler(banHandlers.RemoveBanHandler, "deleteBan", "removeBan",
		"revertBan", "remBan", "rmBan")

	// getInfo handlers
	bindHandler(infoHandlers.GetInfoHandler, "getInfo", "fetchInfo")

	// get all bans handlers
	bindHandler(infoHandlers.GetAllBansHandler, "getAll", "getAllBans", "getbans")

	// stats handlers
	bindHandler(infoHandlers.GetStatsHandler, "getStats", "stats")

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

func bindPostHandler(handler gin.HandlerFunc, paths ...string) {
	for _, path := range paths {
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
		tokenHandlers.GetTokenHandler(c)
	case "getregistered", "registeredusers", "getallregistered":
		tokenHandlers.GetAllRegisteredUsersHandler(c)
	case "addban", "ban", "banuser":
		banHandlers.AddBanHandler(c)
	case "multiban", "addmultiban":
		banHandlers.MultiBanHandler(c)
	case "multiunban", "removemultiban":
		banHandlers.MultiUnBanHandler(c)
	case "deleteban", "removeban", "revertban", "remban":
		banHandlers.RemoveBanHandler(c)
	case "getinfo", "fetchinfo":
		infoHandlers.GetInfoHandler(c)
	case "checktoken":
		infoHandlers.CheckTokenHandler(c)
	case "getall", "getallbans", "getbans":
		infoHandlers.GetAllBansHandler(c)
	case "getstats", "stats":
		infoHandlers.GetStatsHandler(c)
	case "report", "reportuser":
		reportHandlers.ReportUserHandler(c)
	default:
		// TODO: send docs or redirect to docs or something else.
		return
	}
}
