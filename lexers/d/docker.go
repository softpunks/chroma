package d

import (
	. "github.com/maxinekrebs/chroma" // nolint
	. "github.com/maxinekrebs/chroma/lexers/b"
	"github.com/maxinekrebs/chroma/lexers/internal"
)

// Docker lexer.
var Docker = internal.Register(MustNewLexer(
	&Config{
		Name:            "Docker",
		Aliases:         []string{"docker", "dockerfile"},
		Filenames:       []string{"Dockerfile", "*.docker"},
		MimeTypes:       []string{"text/x-dockerfile-config"},
		CaseInsensitive: true,
	},
	Rules{
		"root": {
			{`^(ONBUILD)(\s+)((?:FROM|MAINTAINER|CMD|EXPOSE|ENV|ADD|ENTRYPOINT|VOLUME|WORKDIR))\b`, ByGroups(NameKeyword, TextWhitespace, Keyword), nil},
			{`^((?:FROM|MAINTAINER|CMD|EXPOSE|ENV|ADD|ENTRYPOINT|VOLUME|WORKDIR))\b(.*)`, ByGroups(Keyword, LiteralString), nil},
			{`#.*`, Comment, nil},
			{`RUN`, Keyword, nil},
			{`(.*\\\n)*.+`, Using(Bash), nil},
		},
	},
))
