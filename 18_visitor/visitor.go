package visitor

import (
	"fmt"
	"path"
)

type IResourceFile interface {
	Accpet(visitor Visitor) error
}

type Visitor interface {
	Visit(file IResourceFile) error
}

type PDFFile struct {
	path string
}
func (f *PDFFile) Accpet(visitor Visitor) error {
	return visitor.Visit(f)
}

type PPTFile struct {
	path string
}
func (f *PPTFile) Accpet(visitor Visitor) error {
	return visitor.Visit(f)
}

func NewResourceFile(filepath string) (IResourceFile, error) {
	switch path.Ext(filepath) {
	case ".ppt":
		return &PPTFile{path: filepath}, nil
	case ".pdf":
		return &PDFFile{path: filepath}, nil
	default:
		return nil, fmt.Errorf("not found file type: %s", filepath)
	}
}

type Compressor struct{}

func (c *Compressor) Visit(r IResourceFile) error {
	switch f := r.(type) {
	case *PDFFile:
		return c.VisitPDFFile(f)
	case *PPTFile:
		return c.VisitPPTFile(f)
	default:
		return fmt.Errorf("not found resource typr: %#v", r)
	}
}

func (c *Compressor) VisitPPTFile(f *PPTFile) error {
	fmt.Println("this is ppt file")
	return nil
}
func (c *Compressor) VisitPDFFile(f *PDFFile) error {
	fmt.Println("this is pdf file")
	return nil
}