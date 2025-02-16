package pretty_print

type BlockType int

const (
	ProgramBlockType BlockType = iota
	PhpTagBlockType
	PhpCloseTagBlockType
	TextInterpolationBlockType
	TextBlockType
	EmptyStatementBlockType
	ReferenceModifierBlockType
	FunctionStaticDeclarationBlockType
	StaticVariableDeclarationBlockType
	GlobalDeclarationBlockType
	NamespaceDefinitionBlockType
	NamespaceUseDeclarationBlockType
	NamespaceUseClauseBlockType
	QualifiedNameBlockType
	NamespaceNameBlockType
	NamespaceUseGroupBlockType
	TraitDeclarationBlockType
	InterfaceDeclarationBlockType
	BaseClauseBlockType
	EnumDeclarationBlockType
	EnumDeclarationListBlockType
	EnumCaseBlockType
	ClassDeclarationBlockType
	DeclarationListBlockType
	FinalModifierBlockType
	AbstractModifierBlockType
	ReadonlyModifierBlockType
	ClassInterfaceClauseBlockType
	ConstDeclarationBlockType
	PropertyDeclarationBlockType
	PropertyElementBlockType
	PropertyHookListBlockType
	PropertyHookBlockType
	MethodDeclarationBlockType
	VarModifierBlockType
	StaticModifierBlockType
	UseDeclarationBlockType
	UseListBlockType
	UseInsteadOfClauseBlockType
	UseAsClauseBlockType
	VisibilityModifierBlockType
	FunctionDefinitionBlockType
	AnonymousFunctionBlockType
	AnonymousFunctionUseClauseBlockType
	ArrowFunctionBlockType
	FormalParametersBlockType
	PropertyPromotionParameterBlockType
	SimpleParameterBlockType
	VariadicParameterBlockType
	NamedTypeBlockType
	OptionalTypeBlockType
	BottomTypeBlockType
	UnionTypeBlockType
	IntersectionTypeBlockType
	DisjunctiveNormalFormTypeBlockType
	PrimitiveTypeBlockType
	CastTypeBlockType
	ConstElementBlockType
	EchoStatementBlockType
	ExitStatementBlockType
	UnsetStatementBlockType
	DeclareStatementBlockType
	DeclareDirectiveBlockType
	FloatBlockType
	TryStatementBlockType
	CatchClauseBlockType
	TypeListBlockType
	FinallyClauseBlockType
	GotoStatementBlockType
	ContinueStatementBlockType
	BreakStatementBlockType
	IntegerBlockType
	ReturnStatementBlockType
	ThrowExpressionBlockType
	WhileStatementBlockType
	DoStatementBlockType
	ForStatementBlockType
	SequenceExpressionBlockType
	ForeachStatementBlockType
	PairBlockType
	IfStatementBlockType
	ColonBlockBlockType
	ElseIfClauseBlockType
	ElseClauseBlockType
	MatchExpressionBlockType
	MatchBlockBlockType
	MatchConditionListBlockType
	MatchConditionalExpressionBlockType
	MatchDefaultExpressionBlockType
	SwitchStatementBlockType
	SwitchBlockBlockType
	CaseStatementBlockType
	DefaultStatementBlockType
	CompoundStatementBlockType
	NamedLabelStatementBlockType
	ExpressionStatementBlockType
	UnaryOpExpressionBlockType
	ErrorSuppressionExpressionBlockType
	CloneExpressionBlockType
	ParenthesizedExpressionBlockType
	ClassConstantAccessExpressionBlockType
	PrintIntrinsicBlockType
	ObjectCreationExpressionBlockType
	AnonymousClassBlockType
	UpdateExpressionBlockType
	CastExpressionBlockType
	CastVariableBlockType
	AssignmentExpressionBlockType
	ReferenceAssignmentExpressionBlockType
	ConditionalExpressionBlockType
	AugmentedAssignmentExpressionBlockType
	MemberAccessExpressionBlockType
	NullsafeMemberAccessExpressionBlockType
	ScopedPropertyAccessExpressionBlockType
	ListLiteralBlockType
	FunctionCallExpressionBlockType
	ScopedCallExpressionBlockType
	RelativeScopeBlockType
	VariadicPlaceholderBlockType
	ArgumentsBlockType
	ArgumentBlockType
	MemberCallExpressionBlockType
	NullsafeMemberCallExpressionBlockType
	VariadicUnpackingBlockType
	ArrayCreationExpressionBlockType
	AttributeGroupBlockType
	AttributeListBlockType
	AttributeBlockType
	EscapeSequenceBlockType
	EncapsedStringBlockType
	StringBlockType
	StringContentBlockType
	HeredocBodyBlockType
	HeredocBlockType
	HeredocStartBlockType
	HeredocEndBlockType
	NowdocBodyBlockType
	NowdocBlockType
	ShellCommandExpressionBlockType
	BooleanBlockType
	NullBlockType
	DynamicVariableNameBlockType
	VariableNameBlockType
	ByRefBlockType
	YieldExpressionBlockType
	ArrayElementInitializerBlockType
	BinaryExpressionBlockType
	IncludeExpressionBlockType
	IncludeOnceExpressionBlockType
	RequireExpressionBlockType
	RequireOnceExpressionBlockType
	NameBlockType
	CommentBlockType
	WhitespaceBlockType
	StaticBlockType
	EqualsBlockType
	GlobalBlockType
	NamespaceBlockType
	UseBlockType
	AsBlockType
	FunctionBlockType
	ConstBlockType
	BackslashBlockType
	OpenBraceBlockType
	CloseBraceBlockType
	TraitBlockType
	InterfaceBlockType
	ExtendsBlockType
	EnumBlockType
	ColonBlockType
	ArrayTypeBlockType
	CallableTypeBlockType
	IterableTypeBlockType
	BoolTypeBlockType
	FloatTypeBlockType
	StringTypeBlockType
	IntTypeBlockType
	VoidTypeBlockType
	MixedTypeBlockType
	NullTypeBlockType
	CaseBlockType
	ClassBlockType
	FinalBlockType
	AbstractBlockType
	ReadonlyBlockType
	ImplementsBlockType
	ArrowFunctionSequenceBlockType
	VarBlockType
	InsteadOfBlockType
	PublicBlockType
	ProtectedBlockType
	PrivateBlockType
	OpenParenBlockType
	CloseParenBlockType
	CommaBlockType
	FnBlockType
	ThreeDotBlockType
	QuestionBlockType
	NeverBlockType
	EchoBlockType
	ExitBlockType
	UnsetBlockType
	DeclareBlockType
	EnddeclareBlockType
	TicksBlockType
	EncodingBlockType
	StrictTypesBlockType
	TryBlockType
	CatchBlockType
	FinallyBlockType
	GotoBlockType
	ContinueBlockType
	BreakBlockType
	ReturnBlockType
	ThrowBlockType
	WhileBlockType
	EndwhileBlockType
	DoBlockType
	ForBlockType
	SemicolonBlockType
	EndforBlockType
	ForeachBlockType
	EndforeachBlockType
	IfBlockType
	EndifBlockType
	ElseifBlockType
	ElseBlockType
	MatchBlockType
	DefaultBlockType
	SwitchBlockType
	EndswitchBlockType
	PlusBlockType
	MinusBlockType
	TildeBlockType
	ExclamationBlockType
	AtBlockType
	CloneBlockType
	DoubleColonBlockType
	PrintBlockType
	NewBlockType
	DoubleMinusBlockType
	DoublePlusBlockType
	AmpersandBlockType
	ExponentEqualBlockType
	MultiplyEqualBlockType
	DivideEqualBlockType
	ModuloEqualBlockType
	PlusEqualBlockType
	MinusEqualBlockType
	DotEqualBlockType
	LeftShiftEqualBlockType
	RightShiftEqualBlockType
	AndEqualBlockType
	XorEqualBlockType
	OrEqualBlockType
	NullCoalesceEqualBlockType
	MemberAccessBlockType
	MemberAccessNullBlockType
	OpenBracketBlockType
	CloseBracketBlockType
	SelfBlockType
	ParentBlockType
	TrueBlockType
	FalseBlockType
	OpenAttributeBlockType
	ByteStringBlockType
	SingleQuoteBlockType
	DoubleQuoteBlockType
	HeredocOpenBlockType
	NewlineBlockType
	ShellExecBlockType
	DollarBlockType
	YieldBlockType
	FromBlockType
	InstanceofBlockType
	NullCoalesceBlockType
	ExponentBlockType
	AndBlockType
	OrBlockType
	XorBlockType
	LogicalOrBlockType
	LogicalAndBlockType
	BitwiseOrBlockType
	BitwiseXorBlockType
	EqualBlockType
	NotEqualBlockType
	IdenticalBlockType
	NotIdenticalBlockType
	LessThanBlockType
	GreaterThanBlockType
	LessThanOrEqualBlockType
	GreaterThanOrEqualBlockType
	SpaceshipBlockType
	LeftShiftBlockType
	RightShiftBlockType
	ConcatenationBlockType
	MultiplyBlockType
	DivideBlockType
	ModuloBlockType
	IncludeBlockType
	IncludeOnceBlockType
	RequireBlockType
	RequireOnceBlockType
	HashtagBlockType
	SubscriptExpressionBlockType
	NULL // NULL is a special BlockType used for IndentBlock
	COMPOSITE
)

func isBlockOfTypes(t BlockType, types ...BlockType) bool {
	for _, ty := range types {
		if t == ty {
			return true
		}
	}
	return false
}

func isBlockOfType(t BlockType, ty BlockType) bool {
	return t == ty
}

func isStatementBlockType(t BlockType) bool {
	return t == EmptyStatementBlockType ||
		t == CompoundStatementBlockType ||
		t == NamedLabelStatementBlockType ||
		t == ExpressionStatementBlockType ||
		t == IfStatementBlockType ||
		t == SwitchStatementBlockType ||
		t == WhileStatementBlockType ||
		t == DoStatementBlockType ||
		t == ForStatementBlockType ||
		t == ForeachStatementBlockType ||
		t == GotoStatementBlockType ||
		t == ContinueStatementBlockType ||
		t == BreakStatementBlockType ||
		t == ReturnStatementBlockType ||
		t == TryStatementBlockType ||
		t == DeclareStatementBlockType ||
		t == EchoStatementBlockType ||
		t == ExitStatementBlockType ||
		t == UnsetStatementBlockType ||
		t == ConstDeclarationBlockType ||
		t == FunctionDefinitionBlockType ||
		t == ClassDeclarationBlockType ||
		t == InterfaceDeclarationBlockType ||
		t == TraitDeclarationBlockType ||
		t == EnumDeclarationBlockType ||
		t == NamespaceDefinitionBlockType ||
		t == NamespaceUseDeclarationBlockType ||
		t == GlobalDeclarationBlockType ||
		t == FunctionStaticDeclarationBlockType
}

func isExpressionBlockType(t BlockType) bool {
	return t == ConditionalExpressionBlockType ||
		t == MatchConditionalExpressionBlockType ||
		t == AugmentedAssignmentExpressionBlockType ||
		t == AssignmentExpressionBlockType ||
		t == ReferenceAssignmentExpressionBlockType ||
		t == YieldExpressionBlockType ||
		t == CloneExpressionBlockType ||
		isPrimaryExpressionBlockType(t) ||
		t == UnaryOpExpressionBlockType ||
		t == CastExpressionBlockType ||
		t == ErrorSuppressionExpressionBlockType ||
		t == BinaryExpressionBlockType ||
		t == IncludeExpressionBlockType ||
		t == IncludeOnceExpressionBlockType ||
		t == RequireExpressionBlockType ||
		t == RequireOnceExpressionBlockType
}

func isPrimaryExpressionBlockType(t BlockType) bool {
	return /*isVariableBlockType(t) ||*/ t == ClassConstantAccessExpressionBlockType ||
		isLiteralBlockType(t) ||
		t == QualifiedNameBlockType ||
		t == NameBlockType ||
		t == ArrayCreationExpressionBlockType ||
		t == PrintIntrinsicBlockType ||
		t == AnonymousFunctionBlockType ||
		t == ObjectCreationExpressionBlockType ||
		t == UpdateExpressionBlockType ||
		t == ShellCommandExpressionBlockType ||
		t == ParenthesizedExpressionBlockType ||
		t == ThrowExpressionBlockType ||
		t == ArrowFunctionBlockType
}

//func isVariableBlockType(t BlockType) bool {
//	return t == CastVariableBlockType || //TODO: In the grammar it aliases to CastExpressionBlockType
//		t == VariableNameBlockType ||
//		t == DynamicVariableNameBlockType ||
//
//}

func isLiteralBlockType(t BlockType) bool {
	return t == IntegerBlockType ||
		t == FloatBlockType ||
		t == EncapsedStringBlockType ||
		t == StringBlockType ||
		t == HeredocBlockType ||
		t == NowdocBlockType ||
		t == BooleanBlockType ||
		t == NullBlockType
}

func isTypeBlockType(t BlockType) bool {
	return t == OptionalTypeBlockType ||
		t == NamedTypeBlockType ||
		t == PrimitiveTypeBlockType ||
		t == UnionTypeBlockType ||
		t == IntersectionTypeBlockType ||
		t == DisjunctiveNormalFormTypeBlockType
}
