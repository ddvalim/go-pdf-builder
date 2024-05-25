package pdf

import "github.com/jung-kurt/gofpdf"

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

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, message)

	return *pdf
}
