package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:AsignacionEspacioFisicoDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:AsignacionEspacioFisicoDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:AsignacionEspacioFisicoDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:AsignacionEspacioFisicoDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:AsignacionEspacioFisicoDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EdificioController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EdificioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EdificioController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EdificioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EdificioController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EdificioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EdificioController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EdificioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EdificioController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EdificioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:SedeController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:SedeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:SedeController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:SedeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:SedeController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:SedeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:SedeController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:SedeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:SedeController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:SedeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:TipoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:TipoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:TipoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:TipoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/oikos_api/controllers:TipoEspacioFisicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
