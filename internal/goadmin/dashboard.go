package goadmin

import (
	adminContext "github.com/GoAdminGroup/go-admin/context"
	tmpl "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

func (c *Controller) DashboardPage(ctx *adminContext.Context) (types.Panel, error) {

	components := tmpl.Default()
	colComp := components.Col()

	boxInfo := components.Box().
		WithHeadBorder().
		SetHeader("Generator").
		SetHeadColor("#f7f7f7").
		SetBody(`<div class="clearfix"><a href="/v1/config/generate" 
		class="btn btn-sm btn-info btn-flat pull-left"  target="_blank">Disable All Tests</a>
		<a href="/v1/config/deploy" 
		class="btn btn-sm btn-default btn-flat pull-right" target="_blank">Disable All Tests</a> </div>`).
		GetContent()

	tableCol := colComp.SetSize(types.SizeMD(8)).SetContent(boxInfo).GetContent()
	row1 := components.Row().SetContent(tableCol).GetContent()

	return types.Panel{
		Content:     row1,
		Title:       "Dashboard",
		Description: "",
	}, nil
}
