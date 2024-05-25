package routes

import (
	"github.com/ddvalim/go-pdf-builder/cmd/handler/pdf"
	"github.com/ddvalim/go-pdf-builder/core/ports"
	"net/http"
)

var PDF = []ports.Route{
	{
		URI:     "/pdf",
		Method:  http.MethodPost,
		Handler: pdf.NewHandler().Create,
	},
}
