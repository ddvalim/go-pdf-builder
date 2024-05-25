package pdf

import (
	"fmt"
	"github.com/ddvalim/go-pdf-builder/cmd/response"
	"github.com/ddvalim/go-pdf-builder/internal/pdf"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	PDF pdf.Service
}

func NewHandler() Handler {
	pdfService := pdf.NewService()

	return Handler{
		PDF: pdfService,
	}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		response.Error(w, "something went wrong on parsing multipart form", http.StatusInternalServerError)
	}

	file, handler, err := r.FormFile("sample")
	if err != nil {
		response.Error(w, "something went wrong on getting file", http.StatusInternalServerError)
	}

	defer file.Close()

	if handler.Header.Get("Content-Type") != "text/plain" {
		fmt.Println(handler.Header.Get("Content-Type"))
		response.Error(w, "invalid content type", http.StatusBadRequest)
	}

	bF, err := ioutil.ReadAll(file)
	if err != nil {
		response.Error(w, "something went wrong on reading file", http.StatusInternalServerError)
	}

	inString := string(bF)

	pdfFile := h.PDF.Create(inString)

	defer pdfFile.Close()

	err = pdfFile.Output(w)
	if err != nil {
		response.Error(w, "something went wrong on exporting pdf file", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment")

	response.Write(w, http.StatusOK, nil)
}
