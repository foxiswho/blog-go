package editor

import (
	"bytes"
	"github.com/russross/blackfriday"
	"regexp"
	"strings"
)

var (
	tab    = []byte("\t")
	spaces = []byte("    ")
)

// markdownRender sets some additions instead of default Render
type markdownRender struct {
	blackfriday.Renderer
}

// BlockCode overrides code block
func (mr *markdownRender) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	var tmp bytes.Buffer
	mr.Renderer.BlockCode(&tmp, text, strings.ToLower(lang))
	out.Write(bytes.Replace(tmp.Bytes(), tab, spaces, -1))
}
func Markdown(raw []byte) []byte {
	htmlFlags := 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES |
		blackfriday.HTML_TOC |
		blackfriday.HTML_HREF_TARGET_BLANK

	renderer := &markdownRender{
		Renderer: blackfriday.HtmlRenderer(htmlFlags, "", ""),
	}

	extensions := 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_AUTO_HEADER_IDS |
		blackfriday.EXTENSION_HEADER_IDS

	return blackfriday.Markdown(raw, renderer, extensions)
}
/**
	Markdown自动换行
 */
func MarkdownAutoNewline(str string) string {
	re, _ := regexp.Compile("\\ *\\n")
	str = re.ReplaceAllLiteralString(str, "  \n")
	//m.Content=strings.Replace(m.Content, "\n", "  \n", -1)
	reg := regexp.MustCompile("```([\\s\\S]*)```")
	//返回str中第一个匹配reg的字符串
	data := reg.Find([]byte(str))
	strs := strings.Replace(string(data), "  \n", "\n", -1)
	re, _ = regexp.Compile("```([\\s\\S]*)```")
	return re.ReplaceAllLiteralString(str, strs)
}
