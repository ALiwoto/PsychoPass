package server

import "gitlab.com/Dank-del/SibylAPI-Go/sibyl/entryPoints"

func LoadHandlers() {
	// create token handlers
	ServerEngine.GET("create", entryPoints.CreateToken)
	ServerEngine.POST("create", entryPoints.CreateToken)
	ServerEngine.GET("createToken", entryPoints.CreateToken)
	ServerEngine.POST("createToken", entryPoints.CreateToken)

	// revoke token handlers
	ServerEngine.GET("revoke", entryPoints.RevokeToken)
	ServerEngine.POST("revoke", entryPoints.RevokeToken)
	ServerEngine.GET("revokeToken", entryPoints.RevokeToken)
	ServerEngine.POST("revokeToken", entryPoints.RevokeToken)

	// get token handlers
	ServerEngine.GET("getToken", entryPoints.GetToken)
	ServerEngine.POST("getToken", entryPoints.RevokeToken)

	// addBan handlers
	ServerEngine.GET("addBan", entryPoints.AddBan)
	ServerEngine.POST("addBan", entryPoints.AddBan)

	// deleteBan handlers
	ServerEngine.GET("deleteBan", entryPoints.DeleteBan)
	ServerEngine.POST("deleteBan", entryPoints.DeleteBan)

	// getInfo handlers
	ServerEngine.GET("getInfo", entryPoints.GetInfo)
	ServerEngine.POST("getInfo", entryPoints.GetInfo)
}
