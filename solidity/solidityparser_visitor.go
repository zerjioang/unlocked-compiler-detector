package solidity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SolidityParserVisitor is a complete Visitor for a parse tree produced by SolidityParser.
type SolidityParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SolidityParser#sourceUnit.
	VisitSourceUnit(ctx *SourceUnitContext) interface{}

	// Visit a parse tree produced by SolidityParser#pragmaDirective.
	VisitPragmaDirective(ctx *PragmaDirectiveContext) interface{}

	// Visit a parse tree produced by SolidityParser#importDirective.
	VisitImportDirective(ctx *ImportDirectiveContext) interface{}

	// Visit a parse tree produced by SolidityParser#importAliases.
	VisitImportAliases(ctx *ImportAliasesContext) interface{}

	// Visit a parse tree produced by SolidityParser#path.
	VisitPath(ctx *PathContext) interface{}

	// Visit a parse tree produced by SolidityParser#symbolAliases.
	VisitSymbolAliases(ctx *SymbolAliasesContext) interface{}

	// Visit a parse tree produced by SolidityParser#contractDefinition.
	VisitContractDefinition(ctx *ContractDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#interfaceDefinition.
	VisitInterfaceDefinition(ctx *InterfaceDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#libraryDefinition.
	VisitLibraryDefinition(ctx *LibraryDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#inheritanceSpecifierList.
	VisitInheritanceSpecifierList(ctx *InheritanceSpecifierListContext) interface{}

	// Visit a parse tree produced by SolidityParser#inheritanceSpecifier.
	VisitInheritanceSpecifier(ctx *InheritanceSpecifierContext) interface{}

	// Visit a parse tree produced by SolidityParser#contractBodyElement.
	VisitContractBodyElement(ctx *ContractBodyElementContext) interface{}

	// Visit a parse tree produced by SolidityParser#namedArgument.
	VisitNamedArgument(ctx *NamedArgumentContext) interface{}

	// Visit a parse tree produced by SolidityParser#callArgumentList.
	VisitCallArgumentList(ctx *CallArgumentListContext) interface{}

	// Visit a parse tree produced by SolidityParser#identifierPath.
	VisitIdentifierPath(ctx *IdentifierPathContext) interface{}

	// Visit a parse tree produced by SolidityParser#modifierInvocation.
	VisitModifierInvocation(ctx *ModifierInvocationContext) interface{}

	// Visit a parse tree produced by SolidityParser#visibility.
	VisitVisibility(ctx *VisibilityContext) interface{}

	// Visit a parse tree produced by SolidityParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by SolidityParser#parameterDeclaration.
	VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{}

	// Visit a parse tree produced by SolidityParser#constructorDefinition.
	VisitConstructorDefinition(ctx *ConstructorDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#stateMutability.
	VisitStateMutability(ctx *StateMutabilityContext) interface{}

	// Visit a parse tree produced by SolidityParser#overrideSpecifier.
	VisitOverrideSpecifier(ctx *OverrideSpecifierContext) interface{}

	// Visit a parse tree produced by SolidityParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#modifierDefinition.
	VisitModifierDefinition(ctx *ModifierDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#fallbackFunctionDefinition.
	VisitFallbackFunctionDefinition(ctx *FallbackFunctionDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#receiveFunctionDefinition.
	VisitReceiveFunctionDefinition(ctx *ReceiveFunctionDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#structDefinition.
	VisitStructDefinition(ctx *StructDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#structMember.
	VisitStructMember(ctx *StructMemberContext) interface{}

	// Visit a parse tree produced by SolidityParser#enumDefinition.
	VisitEnumDefinition(ctx *EnumDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#userDefinedValueTypeDefinition.
	VisitUserDefinedValueTypeDefinition(ctx *UserDefinedValueTypeDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#stateVariableDeclaration.
	VisitStateVariableDeclaration(ctx *StateVariableDeclarationContext) interface{}

	// Visit a parse tree produced by SolidityParser#constantVariableDeclaration.
	VisitConstantVariableDeclaration(ctx *ConstantVariableDeclarationContext) interface{}

	// Visit a parse tree produced by SolidityParser#eventParameter.
	VisitEventParameter(ctx *EventParameterContext) interface{}

	// Visit a parse tree produced by SolidityParser#eventDefinition.
	VisitEventDefinition(ctx *EventDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#errorParameter.
	VisitErrorParameter(ctx *ErrorParameterContext) interface{}

	// Visit a parse tree produced by SolidityParser#errorDefinition.
	VisitErrorDefinition(ctx *ErrorDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#usingDirective.
	VisitUsingDirective(ctx *UsingDirectiveContext) interface{}

	// Visit a parse tree produced by SolidityParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by SolidityParser#elementaryTypeName.
	VisitElementaryTypeName(ctx *ElementaryTypeNameContext) interface{}

	// Visit a parse tree produced by SolidityParser#functionTypeName.
	VisitFunctionTypeName(ctx *FunctionTypeNameContext) interface{}

	// Visit a parse tree produced by SolidityParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by SolidityParser#dataLocation.
	VisitDataLocation(ctx *DataLocationContext) interface{}

	// Visit a parse tree produced by SolidityParser#UnaryPrefixOperation.
	VisitUnaryPrefixOperation(ctx *UnaryPrefixOperationContext) interface{}

	// Visit a parse tree produced by SolidityParser#PrimaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by SolidityParser#OrderComparison.
	VisitOrderComparison(ctx *OrderComparisonContext) interface{}

	// Visit a parse tree produced by SolidityParser#Conditional.
	VisitConditional(ctx *ConditionalContext) interface{}

	// Visit a parse tree produced by SolidityParser#PayableConversion.
	VisitPayableConversion(ctx *PayableConversionContext) interface{}

	// Visit a parse tree produced by SolidityParser#Assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by SolidityParser#UnarySuffixOperation.
	VisitUnarySuffixOperation(ctx *UnarySuffixOperationContext) interface{}

	// Visit a parse tree produced by SolidityParser#ShiftOperation.
	VisitShiftOperation(ctx *ShiftOperationContext) interface{}

	// Visit a parse tree produced by SolidityParser#BitAndOperation.
	VisitBitAndOperation(ctx *BitAndOperationContext) interface{}

	// Visit a parse tree produced by SolidityParser#FunctionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by SolidityParser#IndexRangeAccess.
	VisitIndexRangeAccess(ctx *IndexRangeAccessContext) interface{}

	// Visit a parse tree produced by SolidityParser#IndexAccess.
	VisitIndexAccess(ctx *IndexAccessContext) interface{}

	// Visit a parse tree produced by SolidityParser#AddSubOperation.
	VisitAddSubOperation(ctx *AddSubOperationContext) interface{}

	// Visit a parse tree produced by SolidityParser#BitOrOperation.
	VisitBitOrOperation(ctx *BitOrOperationContext) interface{}

	// Visit a parse tree produced by SolidityParser#ExpOperation.
	VisitExpOperation(ctx *ExpOperationContext) interface{}

	// Visit a parse tree produced by SolidityParser#AndOperation.
	VisitAndOperation(ctx *AndOperationContext) interface{}

	// Visit a parse tree produced by SolidityParser#InlineArray.
	VisitInlineArray(ctx *InlineArrayContext) interface{}

	// Visit a parse tree produced by SolidityParser#OrOperation.
	VisitOrOperation(ctx *OrOperationContext) interface{}

	// Visit a parse tree produced by SolidityParser#MemberAccess.
	VisitMemberAccess(ctx *MemberAccessContext) interface{}

	// Visit a parse tree produced by SolidityParser#MulDivModOperation.
	VisitMulDivModOperation(ctx *MulDivModOperationContext) interface{}

	// Visit a parse tree produced by SolidityParser#FunctionCallOptions.
	VisitFunctionCallOptions(ctx *FunctionCallOptionsContext) interface{}

	// Visit a parse tree produced by SolidityParser#BitXorOperation.
	VisitBitXorOperation(ctx *BitXorOperationContext) interface{}

	// Visit a parse tree produced by SolidityParser#Tuple.
	VisitTuple(ctx *TupleContext) interface{}

	// Visit a parse tree produced by SolidityParser#EqualityComparison.
	VisitEqualityComparison(ctx *EqualityComparisonContext) interface{}

	// Visit a parse tree produced by SolidityParser#MetaType.
	VisitMetaType(ctx *MetaTypeContext) interface{}

	// Visit a parse tree produced by SolidityParser#CreateNewExpression.
	VisitCreateNewExpression(ctx *CreateNewExpressionContext) interface{}

	// Visit a parse tree produced by SolidityParser#createNewExpressionContext.
	VisitCreateNewExpressionContext(ctx *CreateNewExpressionContextContext) interface{}

	// Visit a parse tree produced by SolidityParser#assignOp.
	VisitAssignOp(ctx *AssignOpContext) interface{}

	// Visit a parse tree produced by SolidityParser#tupleExpression.
	VisitTupleExpression(ctx *TupleExpressionContext) interface{}

	// Visit a parse tree produced by SolidityParser#inlineArrayExpression.
	VisitInlineArrayExpression(ctx *InlineArrayExpressionContext) interface{}

	// Visit a parse tree produced by SolidityParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by SolidityParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by SolidityParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by SolidityParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by SolidityParser#hexStringLiteral.
	VisitHexStringLiteral(ctx *HexStringLiteralContext) interface{}

	// Visit a parse tree produced by SolidityParser#unicodeStringLiteral.
	VisitUnicodeStringLiteral(ctx *UnicodeStringLiteralContext) interface{}

	// Visit a parse tree produced by SolidityParser#numberLiteral.
	VisitNumberLiteral(ctx *NumberLiteralContext) interface{}

	// Visit a parse tree produced by SolidityParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by SolidityParser#uncheckedBlock.
	VisitUncheckedBlock(ctx *UncheckedBlockContext) interface{}

	// Visit a parse tree produced by SolidityParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#simpleStatement.
	VisitSimpleStatement(ctx *SimpleStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#doWhileStatement.
	VisitDoWhileStatement(ctx *DoWhileStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#tryStatement.
	VisitTryStatement(ctx *TryStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#catchClause.
	VisitCatchClause(ctx *CatchClauseContext) interface{}

	// Visit a parse tree produced by SolidityParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#emitStatement.
	VisitEmitStatement(ctx *EmitStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#revertStatement.
	VisitRevertStatement(ctx *RevertStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#assemblyStatement.
	VisitAssemblyStatement(ctx *AssemblyStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#assemblyFlags.
	VisitAssemblyFlags(ctx *AssemblyFlagsContext) interface{}

	// Visit a parse tree produced by SolidityParser#variableDeclarationList.
	VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{}

	// Visit a parse tree produced by SolidityParser#variableDeclarationTuple.
	VisitVariableDeclarationTuple(ctx *VariableDeclarationTupleContext) interface{}

	// Visit a parse tree produced by SolidityParser#variableDeclarationStatement.
	VisitVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#mappingType.
	VisitMappingType(ctx *MappingTypeContext) interface{}

	// Visit a parse tree produced by SolidityParser#mappingKeyType.
	VisitMappingKeyType(ctx *MappingKeyTypeContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulStatement.
	VisitYulStatement(ctx *YulStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulBlock.
	VisitYulBlock(ctx *YulBlockContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulVariableDeclaration.
	VisitYulVariableDeclaration(ctx *YulVariableDeclarationContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulAssignment.
	VisitYulAssignment(ctx *YulAssignmentContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulIfStatement.
	VisitYulIfStatement(ctx *YulIfStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulForStatement.
	VisitYulForStatement(ctx *YulForStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulSwitchCase.
	VisitYulSwitchCase(ctx *YulSwitchCaseContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulSwitchStatement.
	VisitYulSwitchStatement(ctx *YulSwitchStatementContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulFunctionDefinition.
	VisitYulFunctionDefinition(ctx *YulFunctionDefinitionContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulPath.
	VisitYulPath(ctx *YulPathContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulFunctionCall.
	VisitYulFunctionCall(ctx *YulFunctionCallContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulBoolean.
	VisitYulBoolean(ctx *YulBooleanContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulLiteral.
	VisitYulLiteral(ctx *YulLiteralContext) interface{}

	// Visit a parse tree produced by SolidityParser#yulExpression.
	VisitYulExpression(ctx *YulExpressionContext) interface{}
}
