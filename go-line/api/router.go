package api

func (a *App) SetRoute() {
	a.router.HandleFunc("/ping", a.PingHandler)
}
