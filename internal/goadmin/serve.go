package goadmin

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"os"
	"os/signal"
	"strconv"

	ginAdaptor "github.com/GoAdminGroup/go-admin/adapter/gin"
	goAdminContext "github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	goAdminConfig "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres" // for importing postgres driver
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	_ "github.com/GoAdminGroup/themes/adminlte" // for importing themes
	_ "github.com/GoAdminGroup/themes/sword"    // for importing themes
	"github.com/mohammadVatandoost/ingbusiness/internal/database"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func (c *Controller) ServeAdmin(conf Config, postgresConf database.Config, r *gin.Engine) error {

	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = ioutil.Discard

	engine.Register(new(ginAdaptor.Gin))
	eng := engine.Default()

	cfg := goAdminConfig.Config{
		Env: goAdminConfig.EnvLocal,
		Databases: goAdminConfig.DatabaseList{
			"default": {
				Host:       postgresConf.Host,
				Port:       strconv.Itoa(postgresConf.Port),
				User:       postgresConf.Username,
				Pwd:        postgresConf.Password,
				Name:       postgresConf.Database,
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     goAdminConfig.DriverPostgresql,
			},
		},
		UrlPrefix: "admin",
		IndexUrl:  "/",
		Store: goAdminConfig.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		AccessLogPath:      "./logs/access.log",
		ErrorLogPath:       "./logs/error.log",
		InfoLogPath:        "./logs/info.log",
		AccessAssetsLogOff: true,
		//BootstrapFilePath: "./bootstrap.go",
		Debug:    true,
		Language: language.EN,
	}

	//examplePlugin := example.NewExample()

	template.AddComp(chartjs.NewChart())
	//eng.AddAdapter()
	if err := eng.AddConfig(&cfg).
		//AddGenerators(users.Generators).
		AddGenerator("external", GetExternalTable).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")
	eng.HTML("GET", "/admin", c.DashboardPage)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(conf.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Printf("goadmin listen: %s\n", err.Error())
		}
	}()

	//go e.Logger.Fatal(e.Start(":"+strconv.Itoa(conf.Port)))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Print("closing database connection")
	eng.PostgresqlConnection().Close()

	return nil
}

func GetExternalTable(ctx *goAdminContext.Context) (externalTable table.Table) {

	externalTable = table.NewDefaultTable(table.DefaultConfigWithDriver(goAdminConfig.DriverPostgresql))
	//externalTable = table.NewDefaultTable(table.DefaultConfig())

	info := externalTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Title", "title", db.Varchar)

	info.SetTable("external").
		SetTitle("Externals").
		SetDescription("Externals").
		SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
			return []map[string]interface{}{
				{
					"id":    10,
					"title": "this is a title",
				}, {
					"id":    11,
					"title": "this is a title2",
				}, {
					"id":    12,
					"title": "this is a title3",
				}, {
					"id":    13,
					"title": "this is a title4",
				},
			}, 10
		})

	formList := externalTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("Title", "title", db.Varchar, form.Text)

	formList.SetTable("external").SetTitle("Externals").SetDescription("Externals")

	detail := externalTable.GetDetail()

	detail.SetTable("external").
		SetTitle("Externals").
		SetDescription("Externals").
		SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
			return []map[string]interface{}{
				{
					"id":    10,
					"title": "this is a title",
				},
			}, 1
		})

	return
}
