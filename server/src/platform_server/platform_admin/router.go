package main

import "net/http"

func initRouter(mux MyServeMux) {
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))

	mux.HandleFunc("/api/login", login)
	mux.HandleFunc("/api/logout", logout)
	mux.HandleFunc("/api/register", register)

	mux.HandleFunc("/api/serverall", getServerList)
	mux.HandleFunc("/api/editServer", editServer)

	mux.HandleFunc("/api/patches", getPatches)
	mux.HandleFunc("/api/editPatch", editPatch)

	mux.HandleFunc("/api/getPatchUrl", getPatchUrl)
	mux.HandleFunc("/api/editPatchUrl", editPatchUrl)

	mux.HandleFunc("/api/getTownUrl", getTownUrl)
	mux.HandleFunc("/api/editTownUrl", editTownUrl)

	mux.HandleFunc("/api/towns", getTowns)
	mux.HandleFunc("/api/editTown", editTown)

	mux.HandleFunc("/api/setVersionInfo", setVersionInfo)
	mux.HandleFunc("/api/getVersionInfo", getVersionInfo)

	mux.HandleFunc("/api/getDarkInfo", getDarkInfo)
	mux.HandleFunc("/api/delDarkInfo", delDarkInfo)
	mux.HandleFunc("/api/setDarkInfo", setDarkInfo)

	mux.HandleFunc("/api/isCloseType", isCloseType)
	mux.HandleFunc("/api/setIsCloseType", setIsCloseType)
	mux.HandleFunc("/test", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("test"))
	})
}
