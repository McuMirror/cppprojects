package modules

import (
	"net/http"
	"os"

	"server/engine/fetdata"
	"server/engine/utils"
	"server/engine/wrapper"
)

func (this *Modules) RegisterModule_Api() *Module {
	return this.newModule(MInfo{
		Mount:  "api",
		Name:   "Api",
		Order:  803,
		System: true,
		Icon:   "<i class=\"material-icons notranslate\">api</i>", //assets.SysSvgIconPage,
		Sub:    &[]MISub{},
	}, func(wrap *wrapper.Wrapper) {
		if len(wrap.UrlArgs) == 2 && wrap.UrlArgs[0] == "api" && wrap.UrlArgs[1] == "products" {
			if (*wrap.Config).Api.Xml.Enabled == 1 {
				// Fix url
				if wrap.R.URL.Path[len(wrap.R.URL.Path)-1] != '/' {
					http.Redirect(wrap.W, wrap.R, wrap.R.URL.Path+"/"+utils.ExtractGetParams(wrap.R.RequestURI), 301)
					return
				}

				// Response
				target_file := wrap.DHtdocs + string(os.PathSeparator) + "products.xml"
				if !utils.IsFileExists(target_file) {
					wrap.W.WriteHeader(http.StatusServiceUnavailable)
					wrap.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
					wrap.W.Header().Set("Content-Type", "text/xml; charset=utf-8")
					wrap.W.Write([]byte("In progress..."))
				} else {
					http.ServeFile(wrap.W, wrap.R, target_file)
				}
			} else {
				wrap.W.WriteHeader(http.StatusServiceUnavailable)
				wrap.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				wrap.W.Write([]byte("Disabled!"))
			}
		} else if len(wrap.UrlArgs) == 1 {
			// Fix url
			if wrap.R.URL.Path[len(wrap.R.URL.Path)-1] != '/' {
				http.Redirect(wrap.W, wrap.R, wrap.R.URL.Path+"/"+utils.ExtractGetParams(wrap.R.RequestURI), 301)
				return
			}

			// Some info
			wrap.W.WriteHeader(http.StatusOK)
			wrap.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			wrap.W.Write([]byte("Fave engine API mount point!"))
		} else {
			// User error 404 page
			wrap.RenderFrontEnd("404", fetdata.New(wrap, true, nil, nil), http.StatusNotFound)
			return
		}
	}, func(wrap *wrapper.Wrapper) (string, string, string) {
		// No any page for back-end
		return "", "", ""
	})
}
