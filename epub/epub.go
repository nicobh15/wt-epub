package epub

const (
	CSSFolderName     = "css"
	ImageFolderName   = "images"
	defaultLang       = "en"
	imageFileFormat   = "image%04d%s"
	sectionFileFormat = "section%04d.xhtml"
	urnUUIDPrefix     = "urn:uuid:"
)

type Epub struct {
	toc        *toc
	pkg        *pkg
	css        map[string]string
	author     string
	title      string
	cover      *epubCover
	identifier string
	lang       string
	desc       string
	ppd        string
	images     map[string]string
}

type epubCover struct {
	cssFilename   string
	cssTempFile   string
	imageFilename string
	xhtmlFilename string
}
