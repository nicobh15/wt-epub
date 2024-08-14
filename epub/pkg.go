package epub

// pkg implements the package document file (package.opf), which contains
// metadata about the EPUB (title, author, etc) as well as a list of files the
// EPUB contains.
type pkg struct {
	Title    string
	Author   string
	Language string
	Chapters []Chapter
}

type Chapter struct {
	Title string
	Path  string
}
