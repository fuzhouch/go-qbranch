// Code generated from unofficial-bond-grammar/BondIdlGrammar.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // BondIdlGrammar
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBondIdlGrammarListener is a complete listener for a parse tree produced by BondIdlGrammarParser.
type BaseBondIdlGrammarListener struct{}

var _ BondIdlGrammarListener = &BaseBondIdlGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBondIdlGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBondIdlGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBondIdlGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBondIdlGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBondIdl is called when production bondIdl is entered.
func (s *BaseBondIdlGrammarListener) EnterBondIdl(ctx *BondIdlContext) {}

// ExitBondIdl is called when production bondIdl is exited.
func (s *BaseBondIdlGrammarListener) ExitBondIdl(ctx *BondIdlContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseBondIdlGrammarListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseBondIdlGrammarListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterNamespaceDecl is called when production namespaceDecl is entered.
func (s *BaseBondIdlGrammarListener) EnterNamespaceDecl(ctx *NamespaceDeclContext) {}

// ExitNamespaceDecl is called when production namespaceDecl is exited.
func (s *BaseBondIdlGrammarListener) ExitNamespaceDecl(ctx *NamespaceDeclContext) {}

// EnterNamespaceName is called when production namespaceName is entered.
func (s *BaseBondIdlGrammarListener) EnterNamespaceName(ctx *NamespaceNameContext) {}

// ExitNamespaceName is called when production namespaceName is exited.
func (s *BaseBondIdlGrammarListener) ExitNamespaceName(ctx *NamespaceNameContext) {}

// EnterBondTypeDef is called when production bondTypeDef is entered.
func (s *BaseBondIdlGrammarListener) EnterBondTypeDef(ctx *BondTypeDefContext) {}

// ExitBondTypeDef is called when production bondTypeDef is exited.
func (s *BaseBondIdlGrammarListener) ExitBondTypeDef(ctx *BondTypeDefContext) {}

// EnterTypeAliasDef is called when production typeAliasDef is entered.
func (s *BaseBondIdlGrammarListener) EnterTypeAliasDef(ctx *TypeAliasDefContext) {}

// ExitTypeAliasDef is called when production typeAliasDef is exited.
func (s *BaseBondIdlGrammarListener) ExitTypeAliasDef(ctx *TypeAliasDefContext) {}

// EnterStructDef is called when production structDef is entered.
func (s *BaseBondIdlGrammarListener) EnterStructDef(ctx *StructDefContext) {}

// ExitStructDef is called when production structDef is exited.
func (s *BaseBondIdlGrammarListener) ExitStructDef(ctx *StructDefContext) {}

// EnterStructOrViewDef is called when production structOrViewDef is entered.
func (s *BaseBondIdlGrammarListener) EnterStructOrViewDef(ctx *StructOrViewDefContext) {}

// ExitStructOrViewDef is called when production structOrViewDef is exited.
func (s *BaseBondIdlGrammarListener) ExitStructOrViewDef(ctx *StructOrViewDefContext) {}

// EnterViewOfBody is called when production viewOfBody is entered.
func (s *BaseBondIdlGrammarListener) EnterViewOfBody(ctx *ViewOfBodyContext) {}

// ExitViewOfBody is called when production viewOfBody is exited.
func (s *BaseBondIdlGrammarListener) ExitViewOfBody(ctx *ViewOfBodyContext) {}

// EnterSingleViewOfField is called when production singleViewOfField is entered.
func (s *BaseBondIdlGrammarListener) EnterSingleViewOfField(ctx *SingleViewOfFieldContext) {}

// ExitSingleViewOfField is called when production singleViewOfField is exited.
func (s *BaseBondIdlGrammarListener) ExitSingleViewOfField(ctx *SingleViewOfFieldContext) {}

// EnterStructBaseClassDef is called when production structBaseClassDef is entered.
func (s *BaseBondIdlGrammarListener) EnterStructBaseClassDef(ctx *StructBaseClassDefContext) {}

// ExitStructBaseClassDef is called when production structBaseClassDef is exited.
func (s *BaseBondIdlGrammarListener) ExitStructBaseClassDef(ctx *StructBaseClassDefContext) {}

// EnterStructBody is called when production structBody is entered.
func (s *BaseBondIdlGrammarListener) EnterStructBody(ctx *StructBodyContext) {}

// ExitStructBody is called when production structBody is exited.
func (s *BaseBondIdlGrammarListener) ExitStructBody(ctx *StructBodyContext) {}

// EnterStructFieldDefList is called when production structFieldDefList is entered.
func (s *BaseBondIdlGrammarListener) EnterStructFieldDefList(ctx *StructFieldDefListContext) {}

// ExitStructFieldDefList is called when production structFieldDefList is exited.
func (s *BaseBondIdlGrammarListener) ExitStructFieldDefList(ctx *StructFieldDefListContext) {}

// EnterSingleStructFieldDef is called when production singleStructFieldDef is entered.
func (s *BaseBondIdlGrammarListener) EnterSingleStructFieldDef(ctx *SingleStructFieldDefContext) {}

// ExitSingleStructFieldDef is called when production singleStructFieldDef is exited.
func (s *BaseBondIdlGrammarListener) ExitSingleStructFieldDef(ctx *SingleStructFieldDefContext) {}

// EnterFieldID is called when production fieldID is entered.
func (s *BaseBondIdlGrammarListener) EnterFieldID(ctx *FieldIDContext) {}

// ExitFieldID is called when production fieldID is exited.
func (s *BaseBondIdlGrammarListener) ExitFieldID(ctx *FieldIDContext) {}

// EnterFieldModifier is called when production fieldModifier is entered.
func (s *BaseBondIdlGrammarListener) EnterFieldModifier(ctx *FieldModifierContext) {}

// ExitFieldModifier is called when production fieldModifier is exited.
func (s *BaseBondIdlGrammarListener) ExitFieldModifier(ctx *FieldModifierContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseBondIdlGrammarListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseBondIdlGrammarListener) ExitFieldType(ctx *FieldTypeContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseBondIdlGrammarListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseBondIdlGrammarListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterDefaultValueSpec is called when production defaultValueSpec is entered.
func (s *BaseBondIdlGrammarListener) EnterDefaultValueSpec(ctx *DefaultValueSpecContext) {}

// ExitDefaultValueSpec is called when production defaultValueSpec is exited.
func (s *BaseBondIdlGrammarListener) ExitDefaultValueSpec(ctx *DefaultValueSpecContext) {}

// EnterDefaultValues is called when production defaultValues is entered.
func (s *BaseBondIdlGrammarListener) EnterDefaultValues(ctx *DefaultValuesContext) {}

// ExitDefaultValues is called when production defaultValues is exited.
func (s *BaseBondIdlGrammarListener) ExitDefaultValues(ctx *DefaultValuesContext) {}

// EnterTypeNameWithGeneric is called when production typeNameWithGeneric is entered.
func (s *BaseBondIdlGrammarListener) EnterTypeNameWithGeneric(ctx *TypeNameWithGenericContext) {}

// ExitTypeNameWithGeneric is called when production typeNameWithGeneric is exited.
func (s *BaseBondIdlGrammarListener) ExitTypeNameWithGeneric(ctx *TypeNameWithGenericContext) {}

// EnterBuiltinPrimitiveType is called when production builtinPrimitiveType is entered.
func (s *BaseBondIdlGrammarListener) EnterBuiltinPrimitiveType(ctx *BuiltinPrimitiveTypeContext) {}

// ExitBuiltinPrimitiveType is called when production builtinPrimitiveType is exited.
func (s *BaseBondIdlGrammarListener) ExitBuiltinPrimitiveType(ctx *BuiltinPrimitiveTypeContext) {}

// EnterBuiltinContainerType is called when production builtinContainerType is entered.
func (s *BaseBondIdlGrammarListener) EnterBuiltinContainerType(ctx *BuiltinContainerTypeContext) {}

// ExitBuiltinContainerType is called when production builtinContainerType is exited.
func (s *BaseBondIdlGrammarListener) ExitBuiltinContainerType(ctx *BuiltinContainerTypeContext) {}

// EnterCustomType is called when production customType is entered.
func (s *BaseBondIdlGrammarListener) EnterCustomType(ctx *CustomTypeContext) {}

// ExitCustomType is called when production customType is exited.
func (s *BaseBondIdlGrammarListener) ExitCustomType(ctx *CustomTypeContext) {}

// EnterGenericTypeArgs is called when production genericTypeArgs is entered.
func (s *BaseBondIdlGrammarListener) EnterGenericTypeArgs(ctx *GenericTypeArgsContext) {}

// ExitGenericTypeArgs is called when production genericTypeArgs is exited.
func (s *BaseBondIdlGrammarListener) ExitGenericTypeArgs(ctx *GenericTypeArgsContext) {}

// EnterTypeArgsList is called when production typeArgsList is entered.
func (s *BaseBondIdlGrammarListener) EnterTypeArgsList(ctx *TypeArgsListContext) {}

// ExitTypeArgsList is called when production typeArgsList is exited.
func (s *BaseBondIdlGrammarListener) ExitTypeArgsList(ctx *TypeArgsListContext) {}

// EnterMoreTypeArgsList is called when production moreTypeArgsList is entered.
func (s *BaseBondIdlGrammarListener) EnterMoreTypeArgsList(ctx *MoreTypeArgsListContext) {}

// ExitMoreTypeArgsList is called when production moreTypeArgsList is exited.
func (s *BaseBondIdlGrammarListener) ExitMoreTypeArgsList(ctx *MoreTypeArgsListContext) {}

// EnterTypeArgName is called when production typeArgName is entered.
func (s *BaseBondIdlGrammarListener) EnterTypeArgName(ctx *TypeArgNameContext) {}

// ExitTypeArgName is called when production typeArgName is exited.
func (s *BaseBondIdlGrammarListener) ExitTypeArgName(ctx *TypeArgNameContext) {}

// EnterStructNameDeclWithGeneric is called when production structNameDeclWithGeneric is entered.
func (s *BaseBondIdlGrammarListener) EnterStructNameDeclWithGeneric(ctx *StructNameDeclWithGenericContext) {
}

// ExitStructNameDeclWithGeneric is called when production structNameDeclWithGeneric is exited.
func (s *BaseBondIdlGrammarListener) ExitStructNameDeclWithGeneric(ctx *StructNameDeclWithGenericContext) {
}

// EnterTypeDeclName is called when production typeDeclName is entered.
func (s *BaseBondIdlGrammarListener) EnterTypeDeclName(ctx *TypeDeclNameContext) {}

// ExitTypeDeclName is called when production typeDeclName is exited.
func (s *BaseBondIdlGrammarListener) ExitTypeDeclName(ctx *TypeDeclNameContext) {}

// EnterGenericTypeArgsDecl is called when production genericTypeArgsDecl is entered.
func (s *BaseBondIdlGrammarListener) EnterGenericTypeArgsDecl(ctx *GenericTypeArgsDeclContext) {}

// ExitGenericTypeArgsDecl is called when production genericTypeArgsDecl is exited.
func (s *BaseBondIdlGrammarListener) ExitGenericTypeArgsDecl(ctx *GenericTypeArgsDeclContext) {}

// EnterTypeParamList is called when production typeParamList is entered.
func (s *BaseBondIdlGrammarListener) EnterTypeParamList(ctx *TypeParamListContext) {}

// ExitTypeParamList is called when production typeParamList is exited.
func (s *BaseBondIdlGrammarListener) ExitTypeParamList(ctx *TypeParamListContext) {}

// EnterMoreTypeParamList is called when production moreTypeParamList is entered.
func (s *BaseBondIdlGrammarListener) EnterMoreTypeParamList(ctx *MoreTypeParamListContext) {}

// ExitMoreTypeParamList is called when production moreTypeParamList is exited.
func (s *BaseBondIdlGrammarListener) ExitMoreTypeParamList(ctx *MoreTypeParamListContext) {}

// EnterTypeParamDef is called when production typeParamDef is entered.
func (s *BaseBondIdlGrammarListener) EnterTypeParamDef(ctx *TypeParamDefContext) {}

// ExitTypeParamDef is called when production typeParamDef is exited.
func (s *BaseBondIdlGrammarListener) ExitTypeParamDef(ctx *TypeParamDefContext) {}

// EnterTypeParamName is called when production typeParamName is entered.
func (s *BaseBondIdlGrammarListener) EnterTypeParamName(ctx *TypeParamNameContext) {}

// ExitTypeParamName is called when production typeParamName is exited.
func (s *BaseBondIdlGrammarListener) ExitTypeParamName(ctx *TypeParamNameContext) {}

// EnterTypeParamValueConstraint is called when production typeParamValueConstraint is entered.
func (s *BaseBondIdlGrammarListener) EnterTypeParamValueConstraint(ctx *TypeParamValueConstraintContext) {
}

// ExitTypeParamValueConstraint is called when production typeParamValueConstraint is exited.
func (s *BaseBondIdlGrammarListener) ExitTypeParamValueConstraint(ctx *TypeParamValueConstraintContext) {
}

// EnterAttributeDefList is called when production attributeDefList is entered.
func (s *BaseBondIdlGrammarListener) EnterAttributeDefList(ctx *AttributeDefListContext) {}

// ExitAttributeDefList is called when production attributeDefList is exited.
func (s *BaseBondIdlGrammarListener) ExitAttributeDefList(ctx *AttributeDefListContext) {}

// EnterAttributeDef is called when production attributeDef is entered.
func (s *BaseBondIdlGrammarListener) EnterAttributeDef(ctx *AttributeDefContext) {}

// ExitAttributeDef is called when production attributeDef is exited.
func (s *BaseBondIdlGrammarListener) ExitAttributeDef(ctx *AttributeDefContext) {}

// EnterAttributeBody is called when production attributeBody is entered.
func (s *BaseBondIdlGrammarListener) EnterAttributeBody(ctx *AttributeBodyContext) {}

// ExitAttributeBody is called when production attributeBody is exited.
func (s *BaseBondIdlGrammarListener) ExitAttributeBody(ctx *AttributeBodyContext) {}

// EnterEnumDef is called when production enumDef is entered.
func (s *BaseBondIdlGrammarListener) EnterEnumDef(ctx *EnumDefContext) {}

// ExitEnumDef is called when production enumDef is exited.
func (s *BaseBondIdlGrammarListener) ExitEnumDef(ctx *EnumDefContext) {}

// EnterEnumName is called when production enumName is entered.
func (s *BaseBondIdlGrammarListener) EnterEnumName(ctx *EnumNameContext) {}

// ExitEnumName is called when production enumName is exited.
func (s *BaseBondIdlGrammarListener) ExitEnumName(ctx *EnumNameContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BaseBondIdlGrammarListener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BaseBondIdlGrammarListener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterEnumSymbolDefList is called when production enumSymbolDefList is entered.
func (s *BaseBondIdlGrammarListener) EnterEnumSymbolDefList(ctx *EnumSymbolDefListContext) {}

// ExitEnumSymbolDefList is called when production enumSymbolDefList is exited.
func (s *BaseBondIdlGrammarListener) ExitEnumSymbolDefList(ctx *EnumSymbolDefListContext) {}

// EnterMoreEnumSymbolDef is called when production moreEnumSymbolDef is entered.
func (s *BaseBondIdlGrammarListener) EnterMoreEnumSymbolDef(ctx *MoreEnumSymbolDefContext) {}

// ExitMoreEnumSymbolDef is called when production moreEnumSymbolDef is exited.
func (s *BaseBondIdlGrammarListener) ExitMoreEnumSymbolDef(ctx *MoreEnumSymbolDefContext) {}

// EnterSingleEnumSymbolDef is called when production singleEnumSymbolDef is entered.
func (s *BaseBondIdlGrammarListener) EnterSingleEnumSymbolDef(ctx *SingleEnumSymbolDefContext) {}

// ExitSingleEnumSymbolDef is called when production singleEnumSymbolDef is exited.
func (s *BaseBondIdlGrammarListener) ExitSingleEnumSymbolDef(ctx *SingleEnumSymbolDefContext) {}

// EnterEnumSymbol is called when production enumSymbol is entered.
func (s *BaseBondIdlGrammarListener) EnterEnumSymbol(ctx *EnumSymbolContext) {}

// ExitEnumSymbol is called when production enumSymbol is exited.
func (s *BaseBondIdlGrammarListener) ExitEnumSymbol(ctx *EnumSymbolContext) {}

// EnterEnumSymbolValueAssignment is called when production enumSymbolValueAssignment is entered.
func (s *BaseBondIdlGrammarListener) EnterEnumSymbolValueAssignment(ctx *EnumSymbolValueAssignmentContext) {
}

// ExitEnumSymbolValueAssignment is called when production enumSymbolValueAssignment is exited.
func (s *BaseBondIdlGrammarListener) ExitEnumSymbolValueAssignment(ctx *EnumSymbolValueAssignmentContext) {
}

// EnterEnumSymbolValue is called when production enumSymbolValue is entered.
func (s *BaseBondIdlGrammarListener) EnterEnumSymbolValue(ctx *EnumSymbolValueContext) {}

// ExitEnumSymbolValue is called when production enumSymbolValue is exited.
func (s *BaseBondIdlGrammarListener) ExitEnumSymbolValue(ctx *EnumSymbolValueContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseBondIdlGrammarListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseBondIdlGrammarListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}
