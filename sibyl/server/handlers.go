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
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/entryPoints/pollingHandlers"
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

	// fullRevert handler
	bindHandler(banHandlers.FullRevertHandler, "fullRevert")

	// getInfo handlers
	bindHandler(infoHandlers.GetInfoHandler, "getInfo", "fetchInfo")

	// getInfo handlers
	bindHandler(infoHandlers.GeneralInfoHandler, "generalInfo", "getGeneralInfo")

	// get all bans handlers
	bindHandler(infoHandlers.GetAllBansHandler, "getAll", "getAllBans", "getBans")

	// stats handlers
	bindHandler(infoHandlers.GetStatsHandler, "getStats", "stats")

	// checkToken handlers
	bindHandler(infoHandlers.CheckTokenHandler, "checkToken")

	// report handlers
	bindHandler(reportHandlers.ReportUserHandler, "report", "reportUser", "scan", "scanUser")

	// multiReport handlers
	bindPostHandler(reportHandlers.MultiReportHandler, "multiReport", "multiScan")

	// getUpdates handlers
	bindHandler(pollingHandlers.GetUpdatesHandler, "getUpdates")

	// startPolling handlers
	bindHandler(pollingHandlers.StartPollingHandler, "startPolling")

	bindNoRoot()
}

func bindHandler(handler gin.HandlerFunc, paths ...string) {
	var currentPath string
	for _, path := range paths {
		// ServerEngine.GET(path, handler)
		// ServerEngine.POST(path, handler)

		currentPath = strings.ToLower(path)
		getHandlers[currentPath] = handler
		postHandlers[currentPath] = handler
	}
}

func bindPostHandler(handler gin.HandlerFunc, paths ...string) {
	var currentPath string
	for _, path := range paths {
		// ServerEngine.POST(path, handler)
		currentPath = strings.ToLower(path)
		postHandlers[currentPath] = handler
	}
}

func bindNoRoot() {
	ServerEngine.NoRoute(noRootHandler)
}

func loadDocs() {
	ServerEngine.Static("/docs", "docs"+string(os.PathSeparator)+"site")
}

func addHeaders() {
	ServerEngine.Use(gin.HandlerFunc(func(ctx *gin.Context) {
		// ctx.Header("Content-Security-Policy", "default-src 'none'; style-src 'self'")
		ctx.Header("Access-Control-Allow-Origin", "*")
	}))
}

func noRootHandler(c *gin.Context) {
	path := strings.TrimSpace(c.Request.URL.Path)
	path = strings.Trim(path, "/")
	path = strings.ToLower(path)
	switch c.Request.Method {
	case http.MethodGet:
		h := getHandlers[path]
		if h != nil {
			h(c)
			return
		}
		c.Redirect(http.StatusPermanentRedirect, DocsPath)
		return
	case http.MethodPost:
		h := postHandlers[path]
		if h != nil {
			h(c)
			return
		}
		c.Redirect(http.StatusPermanentRedirect, DocsPath)
		return
	default:
		goto redirect_req
	}

redirect_req: /* redirect requests to /docs */
	c.Redirect(http.StatusPermanentRedirect, DocsPath)
}
