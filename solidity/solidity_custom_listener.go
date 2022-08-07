package solidity

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/zerjioang/solidity-version-check/datatype"
	"regexp"
	"strings"
)

var (
	space = regexp.MustCompile(`\s+`)
)

// CustomSolidityListener is a complete listener for a parse tree produced by the solidity parser.
type CustomSolidityListener struct {
	BaseSolidityParserListener
	debug  bool
	result *datatype.VersionStatus
}

var _ SolidityParserListener = &CustomSolidityListener{}

func NewCustomListener(debug bool, status *datatype.VersionStatus) *CustomSolidityListener {
	c := CustomSolidityListener{
		debug:  debug,
		result: status,
	}
	return &c
}

// implementation of ANTLR listener methods

// VisitTerminal is called when a terminal node is visited.
func (s *CustomSolidityListener) VisitTerminal(node antlr.TerminalNode) {
}

// VisitErrorNode is called when an error node is visited.
func (s *CustomSolidityListener) VisitErrorNode(node antlr.ErrorNode) {
}

// EnterPragmaDirective is called when production pragmaDirective is entered.
func (s *CustomSolidityListener) EnterPragmaDirective(ctx *PragmaDirectiveContext) {
	allLine := ctx.GetText()
	tk := ctx.PragmaToken(0)
	if tk != nil {
		tktext := tk.GetText()
		content := space.ReplaceAllString(tktext, " ")
		if strings.Index(content, "solidity ^") != -1 {
			if !s.result.Errored {
				//alert
				s.result.Errored = true
				s.result.AffectedLine = allLine
				s.result.Start = ctx.GetStart().GetStart()
				s.result.End = ctx.GetStop().GetStop()
				s.result.LineNo = ctx.GetStart().GetLine()
				s.result.SuggestedFix = strings.Replace(allLine, "^", "", 1)
			}
		}
	}
}

// EnterEveryRule is called when any rule is entered.
func (s *CustomSolidityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

// ExitEveryRule is called when any rule is exited.
func (s *CustomSolidityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {

}
