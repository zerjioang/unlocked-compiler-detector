package solidity_version_check

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/zerjioang/solidity-version-check/datatype"
	"github.com/zerjioang/solidity-version-check/solidity"
)

func CheckVersion(filepath string, code []byte) (datatype.VersionStatus, error) {
	var v datatype.VersionStatus
	v.File = filepath
	// Setup the input
	is := antlr.NewInputStream(string(code))

	// Create the Lexer
	lexer := solidity.NewSolidityLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := solidity.NewSolidityParser(stream)

	// Finally, parse the expression
	listener := solidity.NewCustomListener(true, &v)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.SourceUnit())

	return v, nil
}
