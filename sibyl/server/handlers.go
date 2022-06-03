/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package server

import (
	"net/http"
	"os"
	"strings"

	"github.com/MinistryOfWelfare/PsychoPass/sibyl/entryPoints/banHandlers"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/entryPoints/infoHandlers"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/entryPoints/reportHandlers"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/entryPoints/tokenHandlers"
	"github.com/gin-gonic/gin"
)

func LoadHandlers() {
	// add csp and access-control-allow-origin to all responses
	addHeaders()

	// documentations handlers
	loadDocs()

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
		"registeredUsers", "getAllRegistered")

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

	// getInfo handlers
	bindHandler(infoHandlers.GeneralInfoHandler, "generalInfo", "getGeneralInfo")

	// get all bans handlers
	bindHandler(infoHandlers.GetAllBansHandler, "getAll", "getAllBans", "getbans")

	// stats handlers
	bindHandler(infoHandlers.GetStatsHandler, "getStats", "stats")

	// checkToken handlers
	bindHandler(infoHandlers.CheckTokenHandler, "checkToken")

	// report handlers
	bindHandler(reportHandlers.ReportUserHandler, "report", "reportUser", "scan", "scanUser")

	// multiReport handlers
	bindPostHandler(reportHandlers.MultiReportHandler, "multiReport", "multiScan")

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

func loadDocs() {
	ServerEngine.Static("/docs", "docs"+string(os.PathSeparator)+"out")
}

func addHeaders() {
	ServerEngine.Use(gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.Header("Content-Security-Policy", "default-src 'none'; style-src 'self'")
		ctx.Header("Access-Control-Allow-Origin", "*")
	}))
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
		if c.Request.Method != http.MethodPost {
			goto redirect_req
		}
		banHandlers.MultiUnBanHandler(c)

	case "deleteban", "removeban", "revertban", "remban":
		banHandlers.RemoveBanHandler(c)
	case "getinfo", "fetchinfo":
		infoHandlers.GetInfoHandler(c)
	case "generalinfo", "getgeneralinfo":
		infoHandlers.GeneralInfoHandler(c)
	case "checktoken":
		infoHandlers.CheckTokenHandler(c)
	case "getall", "getallbans", "getbans":
		infoHandlers.GetAllBansHandler(c)
	case "getstats", "stats":
		infoHandlers.GetStatsHandler(c)
	case "report", "reportuser":
		reportHandlers.ReportUserHandler(c)
	case "multireport", "multiscan":
		if c.Request.Method != http.MethodPost {
			goto redirect_req
		}
		reportHandlers.MultiReportHandler(c)
	default:
		c.Redirect(http.StatusPermanentRedirect, "/docs/")
		return
	}

	return

redirect_req: /* redirect requests to /docs */
	c.Redirect(http.StatusPermanentRedirect, "/docs/")
}
