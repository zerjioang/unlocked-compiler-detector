package solidity

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSolidityParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSolidityParserVisitor) VisitSourceUnit(ctx *SourceUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitPragmaDirective(ctx *PragmaDirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitImportDirective(ctx *ImportDirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitImportAliases(ctx *ImportAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitPath(ctx *PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitSymbolAliases(ctx *SymbolAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitContractDefinition(ctx *ContractDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitInterfaceDefinition(ctx *InterfaceDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitLibraryDefinition(ctx *LibraryDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitInheritanceSpecifierList(ctx *InheritanceSpecifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitInheritanceSpecifier(ctx *InheritanceSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitContractBodyElement(ctx *ContractBodyElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitNamedArgument(ctx *NamedArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitCallArgumentList(ctx *CallArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitIdentifierPath(ctx *IdentifierPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitModifierInvocation(ctx *ModifierInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitVisibility(ctx *VisibilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitConstructorDefinition(ctx *ConstructorDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitStateMutability(ctx *StateMutabilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitOverrideSpecifier(ctx *OverrideSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitModifierDefinition(ctx *ModifierDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitFallbackFunctionDefinition(ctx *FallbackFunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitReceiveFunctionDefinition(ctx *ReceiveFunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitStructDefinition(ctx *StructDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitStructMember(ctx *StructMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitEnumDefinition(ctx *EnumDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitUserDefinedValueTypeDefinition(ctx *UserDefinedValueTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitStateVariableDeclaration(ctx *StateVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitConstantVariableDeclaration(ctx *ConstantVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitEventParameter(ctx *EventParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitEventDefinition(ctx *EventDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitErrorParameter(ctx *ErrorParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitErrorDefinition(ctx *ErrorDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitUsingDirective(ctx *UsingDirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitElementaryTypeName(ctx *ElementaryTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitFunctionTypeName(ctx *FunctionTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitDataLocation(ctx *DataLocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitUnaryPrefixOperation(ctx *UnaryPrefixOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitOrderComparison(ctx *OrderComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitConditional(ctx *ConditionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitPayableConversion(ctx *PayableConversionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitUnarySuffixOperation(ctx *UnarySuffixOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitShiftOperation(ctx *ShiftOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitBitAndOperation(ctx *BitAndOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitIndexRangeAccess(ctx *IndexRangeAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitIndexAccess(ctx *IndexAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitAddSubOperation(ctx *AddSubOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitBitOrOperation(ctx *BitOrOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitExpOperation(ctx *ExpOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitAndOperation(ctx *AndOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitInlineArray(ctx *InlineArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitOrOperation(ctx *OrOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitMemberAccess(ctx *MemberAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitMulDivModOperation(ctx *MulDivModOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitFunctionCallOptions(ctx *FunctionCallOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitBitXorOperation(ctx *BitXorOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitEqualityComparison(ctx *EqualityComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitMetaType(ctx *MetaTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitCreateNewExpression(ctx *CreateNewExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitCreateNewExpressionContext(ctx *CreateNewExpressionContextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitAssignOp(ctx *AssignOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitTupleExpression(ctx *TupleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitInlineArrayExpression(ctx *InlineArrayExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitHexStringLiteral(ctx *HexStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitUnicodeStringLiteral(ctx *UnicodeStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitNumberLiteral(ctx *NumberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitUncheckedBlock(ctx *UncheckedBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitSimpleStatement(ctx *SimpleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitDoWhileStatement(ctx *DoWhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitTryStatement(ctx *TryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitCatchClause(ctx *CatchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitEmitStatement(ctx *EmitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitRevertStatement(ctx *RevertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitAssemblyStatement(ctx *AssemblyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitAssemblyFlags(ctx *AssemblyFlagsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitVariableDeclarationTuple(ctx *VariableDeclarationTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitMappingType(ctx *MappingTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitMappingKeyType(ctx *MappingKeyTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulStatement(ctx *YulStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulBlock(ctx *YulBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulVariableDeclaration(ctx *YulVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulAssignment(ctx *YulAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulIfStatement(ctx *YulIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulForStatement(ctx *YulForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulSwitchCase(ctx *YulSwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulSwitchStatement(ctx *YulSwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulFunctionDefinition(ctx *YulFunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulPath(ctx *YulPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulFunctionCall(ctx *YulFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulBoolean(ctx *YulBooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulLiteral(ctx *YulLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSolidityParserVisitor) VisitYulExpression(ctx *YulExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
