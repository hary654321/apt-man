package router

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"zrDispatch/core/middleware"
	"zrDispatch/core/router/api/v1/finger"
	"zrDispatch/core/router/api/v1/res"
	sysP "zrDispatch/core/router/api/v1/sys"

	"zrDispatch/core/config"

	"zrDispatch/core/router/api/v1/host"
	"zrDispatch/core/router/api/v1/hostgroup"
	"zrDispatch/core/router/api/v1/install"
	"zrDispatch/core/router/api/v1/notify"
	"zrDispatch/core/router/api/v1/plug"
	"zrDispatch/core/router/api/v1/probe"
	"zrDispatch/core/router/api/v1/task"
	"zrDispatch/core/router/api/v1/user"
	"zrDispatch/core/utils/asset"
	"zrDispatch/core/utils/define"

	_ "zrDispatch/core/docs" // init swagger docs

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
)

// NewHTTPRouter create http.Server
func NewHTTPRouter() *http.Server {
	//gin.SetMode("release")
	router := gin.New()

	fs := &assetfs.AssetFS{
		Asset:     asset.Asset,
		AssetDir:  asset.AssetDir,
		AssetInfo: asset.AssetInfo,
		Prefix:    "web/crocodile",
	}

	router.StaticFS("/crocodile", fs)
	router.GET("/static/*url", func(c *gin.Context) {
		pre, exist := c.Params.Get("url")
		if !exist {
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/crocodile/static"+pre)
	})
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/crocodile/favicon.ico")
	})
	router.GET("/probeTem.csv", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/crocodile/probeTem.csv")
	})
	router.GET("/index.html", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/crocodile")
	})

	v1 := router.Group("/api/v1")
	v1.Use(gin.Recovery(), middleware.ZapLogger(), middleware.PermissionControl(), middleware.Oprtation())
	ru := v1.Group("/user")
	{
		ru.POST("/registry", user.RegistryUser) // only admin // 管理员创建了新的用户。。。
		ru.GET("/info", user.GetUser)
		ru.GET("/all", user.GetUsers)             // only admin
		ru.PUT("/admin", user.AdminChangeUser)    // only admin  // 管理员修改了某某用户
		ru.DELETE("/admin", user.AdminDeleteUser) // only admin  // 管理员删除普通用户
		ru.PUT("/info", user.ChangeUserInfo)      // 某某修改了个人信息
		ru.POST("/login", user.LoginUser)
		ru.POST("/logout", user.LogoutUser) // 某某注销登陆
		ru.GET("/select", user.GetSelect)
		ru.GET("/alarmstatus", user.GetAlarmStatus)
		ru.GET("/operate", user.GetOperateLog)
	}
	rhg := v1.Group("/hostgroup")
	{
		rhg.GET("", hostgroup.GetHostGroups)
		rhg.POST("", hostgroup.CreateHostGroup)
		rhg.PUT("", hostgroup.ChangeHostGroup)
		rhg.DELETE("", hostgroup.DeleteHostGroup)
		rhg.GET("/select", hostgroup.GetSelect)
		rhg.GET("/hosts", hostgroup.GetHostsByIHGID)
	}
	pg := v1.Group("/probegroup")
	{
		pg.POST("", probe.CreateProbeGroup)
		pg.GET("", probe.GetProbeGroup)
		pg.GET("select", probe.GetPgSelect)
		pg.DELETE("", probe.DelProbeGroup)
		pg.PUT("", probe.EditProbeGroup)
	}
	pi := v1.Group("/probeinfo")
	{
		pi.POST("", probe.CreateProbe)
		pi.GET("", probe.GetProbe)
		pi.GET("select", probe.GetPiSelect)
		pi.DELETE("", probe.DelProbe)
		pi.PUT("", probe.EditProbe)
		pi.POST("import", probe.Import)
		pi.GET("/exportCsv", probe.ExportProbeCsv)
		pi.GET("/temcsv", probe.ExportTem)
	}
	sr := v1.Group("/res")
	{
		sr.GET("/port", res.GetPoertRes)
		sr.GET("/getOsSelect", res.GetOsSelect)
		sr.GET("/cert", res.GetCertRes)
		sr.GET("/probe", res.GetProbeRes)
		sr.PUT("/probe", res.UpdateRemark)
		sr.PUT("/port", res.UpdatePortRemark)
		sr.GET("/exportProbeCsv", res.ExportProbeCsv)
		sr.GET("/match", res.Getmatches)
		sr.GET("/matchExport", res.ExportCsv)
	}
	rt := v1.Group("/task")
	{
		rt.GET("", task.GetTasks)
		rt.GET("/info", task.GetTask)
		rt.POST("", task.CreateTask)
		rt.POST("/clone", task.CloneTask)
		rt.PUT("", task.ChangeTask)
		rt.DELETE("", task.DeleteTask)
		rt.PUT("/run", task.RunTask)
		rt.PUT("/kill", task.KillTask)
		rt.PUT("/changestate", task.Changestate)
		rt.GET("/running", task.GetRunningTask)
		rt.DELETE("/log", task.CleanTaskLog)
		rt.GET("/log", task.LogTask)
		rt.GET("/log/tree", task.LogTreeData)
		rt.GET("/log/res", task.TaskRes)
		rt.GET("/log/websocket", task.RealRunTaskLog)
		rt.GET("/status/websocket", task.RealRunTaskStatus)

		rt.GET("/cron", task.ParseCron)
		rt.GET("/select", task.GetSelect)
		rt.GET("/stop", task.Stop)
		rt.GET("/report", task.Report)
	}
	rh := v1.Group("/host")
	{
		rh.GET("", host.GetHost)
		rh.POST("", host.CreateHost)
		rh.PUT("", host.ChangeHost)
		rh.GET("/info", host.GetHostInfo)
		rh.PUT("/stop", host.ChangeHostState)
		rh.DELETE("", host.DeleteHost)
		rh.GET("/select", host.GetSelect)
		rh.GET("/serviceLog", host.ServiceLog)
		rh.GET("/binDeploy", host.BinDeploy)
		rh.GET("/restart", host.Restart)
		rh.GET("/bash", host.Bash)
		rh.GET("/cleanLog", host.CleanLog)
	}
	rn := v1.Group("/notify")
	{
		rn.GET("", notify.GetNotify)
		rn.PUT("", notify.ReadNotify)
	}
	ri := v1.Group("/install")
	{
		ri.GET("/status", install.QueryIsInstall)
		ri.POST("", install.StartInstall)
		ri.GET("/version", install.QueryVersion)
	}
	rf := v1.Group("/finger")
	{
		rf.GET("", finger.GetFinger)
		rf.POST("", finger.AddFinger)
		// 测试指纹
		rf.POST("/test", finger.TestFinger)
		rf.GET("/:id", finger.ScanFinger)
		rf.DELETE("/:id", finger.DeleteFinger)
	}
	sys := v1.Group("/sys")
	{
		sys.GET("log", sysP.RunLog)

		sys.POST("upload", sysP.Upload)
	}

	plugr := v1.Group("/plug")
	{
		plugr.POST("", plug.CreatePlug)
		plugr.GET("", plug.GetPlug)
		plugr.GET("select", plug.GetPiSelect)
		plugr.DELETE("", plug.DelPlug)
		plugr.PUT("", plug.EditPlug)
	}
	// if nor find router, will rediret to /crocodile/
	router.NoRoute(func(c *gin.Context) {
		//c.Redirect(http.StatusMovedPermanently, "/crocodile/")
	})

	httpSrv := &http.Server{
		Handler: router,
		// ReadTimeout:  config.CoreConf.Server.MaxHTTPTime.Duration,
		// WriteTimeout: config.CoreConf.Server.MaxHTTPTime.Duration,
	}
	return httpSrv
}

// GetListen get listen addr by server or client
func GetListen(mode define.RunMode) (net.Listener, error) {
	var (
		addr string
	)
	switch mode {
	case define.Server:
		if os.Getenv("PORT") != "" {
			addr = ":" + os.Getenv("PORT")
		} else {
			addr = fmt.Sprintf(":%d", config.CoreConf.Server.Port)
		}

	case define.Client:
		addr = fmt.Sprintf(":%d", config.CoreConf.Client.Port)

	default:
		return nil, errors.New("Unsupport mode")
	}
	lis, err := net.Listen("tcp", addr)

	return lis, err
}

// Run start run http or grpc Server
func Run(lis net.Listener) error {
	var (
		httpServer *http.Server
		err        error
		m          cmux.CMux
	)

	//gRPCServer, err = schedule.NewgRPCServer(mode)
	if err != nil {
		return err
	}

	m = cmux.New(lis)

	httpServer = NewHTTPRouter()
	httpL := m.Match(cmux.HTTP1Fast())
	go httpServer.Serve(httpL)

	return m.Serve()
}
