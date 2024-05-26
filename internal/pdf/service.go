package pdf

import (
	"github.com/jung-kurt/gofpdf"
	"strings"
)

type Service interface {
	Create(message string) gofpdf.Fpdf
}

type ServiceImpl struct {
}

func NewService() ServiceImpl {
	return ServiceImpl{}
}

func (s ServiceImpl) Create(message string) gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetFont("Arial", "B", 16)
	pdf.SetMargins(10, 20, 10)

	pdf.AddPage()

	for _, line := range splitIntoLines(message) {
		tr := pdf.UnicodeTranslatorFromDescriptor("")

		pdf.MultiCell(190, 10, tr(line), "", "", false)
	}

	return *pdf
}

func splitIntoLines(message string) []string {
	return strings.Split(message, "\n")
}
