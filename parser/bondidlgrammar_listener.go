// Code generated from unofficial-bond-grammar/BondIdlGrammar.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // BondIdlGrammar
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BondIdlGrammarListener is a complete listener for a parse tree produced by BondIdlGrammarParser.
type BondIdlGrammarListener interface {
	antlr.ParseTreeListener

	// EnterBondIdl is called when entering the bondIdl production.
	EnterBondIdl(c *BondIdlContext)

	// EnterImportDecl is called when entering the importDecl production.
	EnterImportDecl(c *ImportDeclContext)

	// EnterNamespaceDecl is called when entering the namespaceDecl production.
	EnterNamespaceDecl(c *NamespaceDeclContext)

	// EnterNamespaceName is called when entering the namespaceName production.
	EnterNamespaceName(c *NamespaceNameContext)

	// EnterBondTypeDef is called when entering the bondTypeDef production.
	EnterBondTypeDef(c *BondTypeDefContext)

	// EnterTypeAliasDef is called when entering the typeAliasDef production.
	EnterTypeAliasDef(c *TypeAliasDefContext)

	// EnterStructDef is called when entering the structDef production.
	EnterStructDef(c *StructDefContext)

	// EnterStructOrViewDef is called when entering the structOrViewDef production.
	EnterStructOrViewDef(c *StructOrViewDefContext)

	// EnterViewOfBody is called when entering the viewOfBody production.
	EnterViewOfBody(c *ViewOfBodyContext)

	// EnterSingleViewOfField is called when entering the singleViewOfField production.
	EnterSingleViewOfField(c *SingleViewOfFieldContext)

	// EnterStructBaseClassDef is called when entering the structBaseClassDef production.
	EnterStructBaseClassDef(c *StructBaseClassDefContext)

	// EnterStructBody is called when entering the structBody production.
	EnterStructBody(c *StructBodyContext)

	// EnterStructFieldDefList is called when entering the structFieldDefList production.
	EnterStructFieldDefList(c *StructFieldDefListContext)

	// EnterSingleStructFieldDef is called when entering the singleStructFieldDef production.
	EnterSingleStructFieldDef(c *SingleStructFieldDefContext)

	// EnterFieldID is called when entering the fieldID production.
	EnterFieldID(c *FieldIDContext)

	// EnterFieldModifier is called when entering the fieldModifier production.
	EnterFieldModifier(c *FieldModifierContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterDefaultValueSpec is called when entering the defaultValueSpec production.
	EnterDefaultValueSpec(c *DefaultValueSpecContext)

	// EnterDefaultValues is called when entering the defaultValues production.
	EnterDefaultValues(c *DefaultValuesContext)

	// EnterTypeNameWithGeneric is called when entering the typeNameWithGeneric production.
	EnterTypeNameWithGeneric(c *TypeNameWithGenericContext)

	// EnterBuiltinPrimitiveType is called when entering the builtinPrimitiveType production.
	EnterBuiltinPrimitiveType(c *BuiltinPrimitiveTypeContext)

	// EnterBuiltinContainerType is called when entering the builtinContainerType production.
	EnterBuiltinContainerType(c *BuiltinContainerTypeContext)

	// EnterCustomType is called when entering the customType production.
	EnterCustomType(c *CustomTypeContext)

	// EnterGenericTypeArgs is called when entering the genericTypeArgs production.
	EnterGenericTypeArgs(c *GenericTypeArgsContext)

	// EnterTypeArgsList is called when entering the typeArgsList production.
	EnterTypeArgsList(c *TypeArgsListContext)

	// EnterMoreTypeArgsList is called when entering the moreTypeArgsList production.
	EnterMoreTypeArgsList(c *MoreTypeArgsListContext)

	// EnterTypeArgName is called when entering the typeArgName production.
	EnterTypeArgName(c *TypeArgNameContext)

	// EnterStructNameDeclWithGeneric is called when entering the structNameDeclWithGeneric production.
	EnterStructNameDeclWithGeneric(c *StructNameDeclWithGenericContext)

	// EnterTypeDeclName is called when entering the typeDeclName production.
	EnterTypeDeclName(c *TypeDeclNameContext)

	// EnterGenericTypeArgsDecl is called when entering the genericTypeArgsDecl production.
	EnterGenericTypeArgsDecl(c *GenericTypeArgsDeclContext)

	// EnterTypeParamList is called when entering the typeParamList production.
	EnterTypeParamList(c *TypeParamListContext)

	// EnterMoreTypeParamList is called when entering the moreTypeParamList production.
	EnterMoreTypeParamList(c *MoreTypeParamListContext)

	// EnterTypeParamDef is called when entering the typeParamDef production.
	EnterTypeParamDef(c *TypeParamDefContext)

	// EnterTypeParamName is called when entering the typeParamName production.
	EnterTypeParamName(c *TypeParamNameContext)

	// EnterTypeParamValueConstraint is called when entering the typeParamValueConstraint production.
	EnterTypeParamValueConstraint(c *TypeParamValueConstraintContext)

	// EnterAttributeDefList is called when entering the attributeDefList production.
	EnterAttributeDefList(c *AttributeDefListContext)

	// EnterAttributeDef is called when entering the attributeDef production.
	EnterAttributeDef(c *AttributeDefContext)

	// EnterAttributeBody is called when entering the attributeBody production.
	EnterAttributeBody(c *AttributeBodyContext)

	// EnterEnumDef is called when entering the enumDef production.
	EnterEnumDef(c *EnumDefContext)

	// EnterEnumName is called when entering the enumName production.
	EnterEnumName(c *EnumNameContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterEnumSymbolDefList is called when entering the enumSymbolDefList production.
	EnterEnumSymbolDefList(c *EnumSymbolDefListContext)

	// EnterMoreEnumSymbolDef is called when entering the moreEnumSymbolDef production.
	EnterMoreEnumSymbolDef(c *MoreEnumSymbolDefContext)

	// EnterSingleEnumSymbolDef is called when entering the singleEnumSymbolDef production.
	EnterSingleEnumSymbolDef(c *SingleEnumSymbolDefContext)

	// EnterEnumSymbol is called when entering the enumSymbol production.
	EnterEnumSymbol(c *EnumSymbolContext)

	// EnterEnumSymbolValueAssignment is called when entering the enumSymbolValueAssignment production.
	EnterEnumSymbolValueAssignment(c *EnumSymbolValueAssignmentContext)

	// EnterEnumSymbolValue is called when entering the enumSymbolValue production.
	EnterEnumSymbolValue(c *EnumSymbolValueContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// ExitBondIdl is called when exiting the bondIdl production.
	ExitBondIdl(c *BondIdlContext)

	// ExitImportDecl is called when exiting the importDecl production.
	ExitImportDecl(c *ImportDeclContext)

	// ExitNamespaceDecl is called when exiting the namespaceDecl production.
	ExitNamespaceDecl(c *NamespaceDeclContext)

	// ExitNamespaceName is called when exiting the namespaceName production.
	ExitNamespaceName(c *NamespaceNameContext)

	// ExitBondTypeDef is called when exiting the bondTypeDef production.
	ExitBondTypeDef(c *BondTypeDefContext)

	// ExitTypeAliasDef is called when exiting the typeAliasDef production.
	ExitTypeAliasDef(c *TypeAliasDefContext)

	// ExitStructDef is called when exiting the structDef production.
	ExitStructDef(c *StructDefContext)

	// ExitStructOrViewDef is called when exiting the structOrViewDef production.
	ExitStructOrViewDef(c *StructOrViewDefContext)

	// ExitViewOfBody is called when exiting the viewOfBody production.
	ExitViewOfBody(c *ViewOfBodyContext)

	// ExitSingleViewOfField is called when exiting the singleViewOfField production.
	ExitSingleViewOfField(c *SingleViewOfFieldContext)

	// ExitStructBaseClassDef is called when exiting the structBaseClassDef production.
	ExitStructBaseClassDef(c *StructBaseClassDefContext)

	// ExitStructBody is called when exiting the structBody production.
	ExitStructBody(c *StructBodyContext)

	// ExitStructFieldDefList is called when exiting the structFieldDefList production.
	ExitStructFieldDefList(c *StructFieldDefListContext)

	// ExitSingleStructFieldDef is called when exiting the singleStructFieldDef production.
	ExitSingleStructFieldDef(c *SingleStructFieldDefContext)

	// ExitFieldID is called when exiting the fieldID production.
	ExitFieldID(c *FieldIDContext)

	// ExitFieldModifier is called when exiting the fieldModifier production.
	ExitFieldModifier(c *FieldModifierContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitDefaultValueSpec is called when exiting the defaultValueSpec production.
	ExitDefaultValueSpec(c *DefaultValueSpecContext)

	// ExitDefaultValues is called when exiting the defaultValues production.
	ExitDefaultValues(c *DefaultValuesContext)

	// ExitTypeNameWithGeneric is called when exiting the typeNameWithGeneric production.
	ExitTypeNameWithGeneric(c *TypeNameWithGenericContext)

	// ExitBuiltinPrimitiveType is called when exiting the builtinPrimitiveType production.
	ExitBuiltinPrimitiveType(c *BuiltinPrimitiveTypeContext)

	// ExitBuiltinContainerType is called when exiting the builtinContainerType production.
	ExitBuiltinContainerType(c *BuiltinContainerTypeContext)

	// ExitCustomType is called when exiting the customType production.
	ExitCustomType(c *CustomTypeContext)

	// ExitGenericTypeArgs is called when exiting the genericTypeArgs production.
	ExitGenericTypeArgs(c *GenericTypeArgsContext)

	// ExitTypeArgsList is called when exiting the typeArgsList production.
	ExitTypeArgsList(c *TypeArgsListContext)

	// ExitMoreTypeArgsList is called when exiting the moreTypeArgsList production.
	ExitMoreTypeArgsList(c *MoreTypeArgsListContext)

	// ExitTypeArgName is called when exiting the typeArgName production.
	ExitTypeArgName(c *TypeArgNameContext)

	// ExitStructNameDeclWithGeneric is called when exiting the structNameDeclWithGeneric production.
	ExitStructNameDeclWithGeneric(c *StructNameDeclWithGenericContext)

	// ExitTypeDeclName is called when exiting the typeDeclName production.
	ExitTypeDeclName(c *TypeDeclNameContext)

	// ExitGenericTypeArgsDecl is called when exiting the genericTypeArgsDecl production.
	ExitGenericTypeArgsDecl(c *GenericTypeArgsDeclContext)

	// ExitTypeParamList is called when exiting the typeParamList production.
	ExitTypeParamList(c *TypeParamListContext)

	// ExitMoreTypeParamList is called when exiting the moreTypeParamList production.
	ExitMoreTypeParamList(c *MoreTypeParamListContext)

	// ExitTypeParamDef is called when exiting the typeParamDef production.
	ExitTypeParamDef(c *TypeParamDefContext)

	// ExitTypeParamName is called when exiting the typeParamName production.
	ExitTypeParamName(c *TypeParamNameContext)

	// ExitTypeParamValueConstraint is called when exiting the typeParamValueConstraint production.
	ExitTypeParamValueConstraint(c *TypeParamValueConstraintContext)

	// ExitAttributeDefList is called when exiting the attributeDefList production.
	ExitAttributeDefList(c *AttributeDefListContext)

	// ExitAttributeDef is called when exiting the attributeDef production.
	ExitAttributeDef(c *AttributeDefContext)

	// ExitAttributeBody is called when exiting the attributeBody production.
	ExitAttributeBody(c *AttributeBodyContext)

	// ExitEnumDef is called when exiting the enumDef production.
	ExitEnumDef(c *EnumDefContext)

	// ExitEnumName is called when exiting the enumName production.
	ExitEnumName(c *EnumNameContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitEnumSymbolDefList is called when exiting the enumSymbolDefList production.
	ExitEnumSymbolDefList(c *EnumSymbolDefListContext)

	// ExitMoreEnumSymbolDef is called when exiting the moreEnumSymbolDef production.
	ExitMoreEnumSymbolDef(c *MoreEnumSymbolDefContext)

	// ExitSingleEnumSymbolDef is called when exiting the singleEnumSymbolDef production.
	ExitSingleEnumSymbolDef(c *SingleEnumSymbolDefContext)

	// ExitEnumSymbol is called when exiting the enumSymbol production.
	ExitEnumSymbol(c *EnumSymbolContext)

	// ExitEnumSymbolValueAssignment is called when exiting the enumSymbolValueAssignment production.
	ExitEnumSymbolValueAssignment(c *EnumSymbolValueAssignmentContext)

	// ExitEnumSymbolValue is called when exiting the enumSymbolValue production.
	ExitEnumSymbolValue(c *EnumSymbolValueContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)
}
