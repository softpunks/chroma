package x

import (
	. "github.com/maxinekrebs/chroma" // nolint
	"github.com/maxinekrebs/chroma/lexers/internal"
)

// Xorg lexer.
var Xorg = internal.Register(MustNewLexer(
	&Config{
		Name:      "Xorg",
		Aliases:   []string{"xorg.conf"},
		Filenames: []string{"xorg.conf"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {
			{`\s+`, TextWhitespace, nil},
			{`#.*$`, Comment, nil},
			{`((|Sub)Section)(\s+)("\w+")`, ByGroups(KeywordNamespace, LiteralStringEscape, TextWhitespace, LiteralStringEscape), nil},
			{`(End(|Sub)Section)`, KeywordNamespace, nil},
			{`(\w+)(\s+)([^\n#]+)`, ByGroups(NameKeyword, TextWhitespace, LiteralString), nil},
		},
	},
))
