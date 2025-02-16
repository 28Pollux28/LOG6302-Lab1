package pretty_print

import (
	"github.com/28Pollux28/log6302-parser/utils"
)

// Primitive Blocks that are used in multiple other blocks
var (
	WHITESPACE_BLOCK              = &PrimitiveBlock{Content: " ", BlockType: WhitespaceBlockType}
	PHP_TAG_BLOCK                 = &PrimitiveBlock{Content: "<?php\n", BlockType: PhpTagBlockType}
	PHP_CLOSE_TAG_BLOCK           = &PrimitiveBlock{Content: "?>", BlockType: PhpCloseTagBlockType}
	STATIC_BLOCK                  = &PrimitiveBlock{Content: "static", BlockType: StaticBlockType}
	EQUALS_BLOCK                  = &PrimitiveBlock{Content: "=", BlockType: EqualsBlockType}
	GLOBAL_BLOCK                  = &PrimitiveBlock{Content: "global", BlockType: GlobalBlockType}
	NAMESPACE_BLOCK               = &PrimitiveBlock{Content: "namespace", BlockType: NamespaceBlockType}
	USE_BLOCK                     = &PrimitiveBlock{Content: "use", BlockType: UseBlockType}
	AS_BLOCK                      = &PrimitiveBlock{Content: "as", BlockType: AsBlockType}
	FUNCTION_BLOCK                = &PrimitiveBlock{Content: "function", BlockType: FunctionBlockType}
	CONST_BLOCK                   = &PrimitiveBlock{Content: "const", BlockType: ConstBlockType}
	BACKSLASH_BLOCK               = &PrimitiveBlock{Content: "\\", BlockType: BackslashBlockType}
	OPEN_BRACE_BLOCK              = &PrimitiveBlock{Content: "{", BlockType: OpenBraceBlockType}
	CLOSE_BRACE_BLOCK             = &PrimitiveBlock{Content: "}", BlockType: CloseBraceBlockType}
	TRAIT_BLOCK                   = &PrimitiveBlock{Content: "trait", BlockType: TraitBlockType}
	INTERFACE_BLOCK               = &PrimitiveBlock{Content: "interface", BlockType: InterfaceBlockType}
	EXTENDS_BLOCK                 = &PrimitiveBlock{Content: "extends", BlockType: ExtendsBlockType}
	ENUM_BLOCK                    = &PrimitiveBlock{Content: "enum", BlockType: EnumBlockType}
	COLON_BLOCK                   = &PrimitiveBlock{Content: ":", BlockType: ColonBlockType}
	ARRAY_TYPE_BLOCK              = &PrimitiveBlock{Content: "array", BlockType: ArrayTypeBlockType}
	CALLABLE_TYPE_BLOCK           = &PrimitiveBlock{Content: "callable", BlockType: CallableTypeBlockType}
	ITERABLE_TYPE_BLOCK           = &PrimitiveBlock{Content: "iterable", BlockType: IterableTypeBlockType}
	BOOL_TYPE_BLOCK               = &PrimitiveBlock{Content: "bool", BlockType: BoolTypeBlockType}
	FLOAT_TYPE_BLOCK              = &PrimitiveBlock{Content: "float", BlockType: FloatTypeBlockType}
	STRING_TYPE_BLOCK             = &PrimitiveBlock{Content: "string", BlockType: StringTypeBlockType}
	INT_TYPE_BLOCK                = &PrimitiveBlock{Content: "int", BlockType: IntTypeBlockType}
	VOID_TYPE_BLOCK               = &PrimitiveBlock{Content: "void", BlockType: VoidTypeBlockType}
	MIXED_TYPE_BLOCK              = &PrimitiveBlock{Content: "mixed", BlockType: MixedTypeBlockType}
	NULL_TYPE_BLOCK               = &PrimitiveBlock{Content: "null", BlockType: NullTypeBlockType}
	CASE_BLOCK                    = &PrimitiveBlock{Content: "case", BlockType: CaseBlockType}
	CLASS_BLOCK                   = &PrimitiveBlock{Content: "class", BlockType: ClassBlockType}
	FINAL_BLOCK                   = &PrimitiveBlock{Content: "final", BlockType: FinalBlockType}
	ABSTRACT_BLOCK                = &PrimitiveBlock{Content: "abstract", BlockType: AbstractBlockType}
	READONLY_BLOCK                = &PrimitiveBlock{Content: "readonly", BlockType: ReadonlyBlockType}
	IMPLEMENTS_BLOCK              = &PrimitiveBlock{Content: "implements", BlockType: ImplementsBlockType}
	ARROW_FUNCTION_SEQUENCE_BLOCK = &PrimitiveBlock{Content: "=>", BlockType: ArrowFunctionSequenceBlockType}
	VAR_BLOCK                     = &PrimitiveBlock{Content: "var", BlockType: VarBlockType}
	INSTEAD_OF_BLOCK              = &PrimitiveBlock{Content: "insteadof", BlockType: InsteadOfBlockType}
	PUBLIC_BLOCK                  = &PrimitiveBlock{Content: "public", BlockType: PublicBlockType}
	PROTECTED_BLOCK               = &PrimitiveBlock{Content: "protected", BlockType: ProtectedBlockType}
	PRIVATE_BLOCK                 = &PrimitiveBlock{Content: "private", BlockType: PrivateBlockType}
	OPEN_PAREN_BLOCK              = &PrimitiveBlock{Content: "(", BlockType: OpenParenBlockType}
	CLOSE_PAREN_BLOCK             = &PrimitiveBlock{Content: ")", BlockType: CloseParenBlockType}
	COMMA_BLOCK                   = &PrimitiveBlock{Content: ",", BlockType: CommaBlockType}
	FN_BLOCK                      = &PrimitiveBlock{Content: "fn", BlockType: FnBlockType}
	THREE_DOT_BLOCK               = &PrimitiveBlock{Content: "...", BlockType: ThreeDotBlockType}
	QUESTION_BLOCK                = &PrimitiveBlock{Content: "?", BlockType: QuestionBlockType}
	NEVER_BLOCK                   = &PrimitiveBlock{Content: "never", BlockType: NeverBlockType}
	ECHO_BLOCK                    = &PrimitiveBlock{Content: "echo", BlockType: EchoBlockType}
	EXIT_BLOCK                    = &PrimitiveBlock{Content: "exit", BlockType: ExitBlockType}
	UNSET_BLOCK                   = &PrimitiveBlock{Content: "unset", BlockType: UnsetBlockType}
	DECLARE_BLOCK                 = &PrimitiveBlock{Content: "declare", BlockType: DeclareBlockType}
	ENDDECLARE_BLOCK              = &PrimitiveBlock{Content: "enddeclare", BlockType: EnddeclareBlockType}
	TICKS_BLOCK                   = &PrimitiveBlock{Content: "ticks", BlockType: TicksBlockType}
	ENCODING_BLOCK                = &PrimitiveBlock{Content: "encoding", BlockType: EncodingBlockType}
	STRIC_TYPES_BLOCK             = &PrimitiveBlock{Content: "strict_types", BlockType: StrictTypesBlockType}
	TRY_BLOCK                     = &PrimitiveBlock{Content: "try", BlockType: TryBlockType}
	CATCH_BLOCK                   = &PrimitiveBlock{Content: "catch", BlockType: CatchBlockType}
	FINALLY_BLOCK                 = &PrimitiveBlock{Content: "finally", BlockType: FinallyBlockType}
	GOTO_BLOCK                    = &PrimitiveBlock{Content: "goto", BlockType: GotoBlockType}
	CONTINUE_BLOCK                = &PrimitiveBlock{Content: "continue", BlockType: ContinueBlockType}
	BREAK_BLOCK                   = &PrimitiveBlock{Content: "break", BlockType: BreakBlockType}
	RETURN_BLOCK                  = &PrimitiveBlock{Content: "return", BlockType: ReturnBlockType}
	THROW_BLOCK                   = &PrimitiveBlock{Content: "throw", BlockType: ThrowBlockType}
	WHILE_BLOCK                   = &PrimitiveBlock{Content: "while", BlockType: WhileBlockType}
	ENDWHILE_BLOCK                = &PrimitiveBlock{Content: "endwhile", BlockType: EndwhileBlockType}
	DO_BLOCK                      = &PrimitiveBlock{Content: "do", BlockType: DoBlockType}
	FOR_BLOCK                     = &PrimitiveBlock{Content: "for", BlockType: ForBlockType}
	SEMICOLON_BLOCK               = &PrimitiveBlock{Content: ";", BlockType: SemicolonBlockType}
	ENDFOR_BLOCK                  = &PrimitiveBlock{Content: "endfor", BlockType: EndforBlockType}
	FOREACH_BLOCK                 = &PrimitiveBlock{Content: "foreach", BlockType: ForeachBlockType}
	ENDFOREACH_BLOCK              = &PrimitiveBlock{Content: "endforeach", BlockType: EndforeachBlockType}
	IF_BLOCK                      = &PrimitiveBlock{Content: "if", BlockType: IfBlockType}
	ENDIF_BLOCK                   = &PrimitiveBlock{Content: "endif", BlockType: EndifBlockType}
	ELSEIF_BLOCK                  = &PrimitiveBlock{Content: "elseif", BlockType: ElseifBlockType}
	ELSE_BLOCK                    = &PrimitiveBlock{Content: "else", BlockType: ElseBlockType}
	MATCH_BLOCK                   = &PrimitiveBlock{Content: "match", BlockType: MatchBlockType}
	DEFAULT_BLOCK                 = &PrimitiveBlock{Content: "default", BlockType: DefaultBlockType}
	SWITCH_BLOCK                  = &PrimitiveBlock{Content: "switch", BlockType: SwitchBlockType}
	ENDSWITCH_BLOCK               = &PrimitiveBlock{Content: "endswitch", BlockType: EndswitchBlockType}
	PLUS_BLOCK                    = &PrimitiveBlock{Content: "+", BlockType: PlusBlockType}
	MINUS_BLOCK                   = &PrimitiveBlock{Content: "-", BlockType: MinusBlockType}
	TILDE_BLOCK                   = &PrimitiveBlock{Content: "~", BlockType: TildeBlockType}
	EXCLAMATION_BLOCK             = &PrimitiveBlock{Content: "!", BlockType: ExclamationBlockType}
	AT_BLOCK                      = &PrimitiveBlock{Content: "@", BlockType: AtBlockType}
	CLONE_BLOCK                   = &PrimitiveBlock{Content: "clone", BlockType: CloneBlockType}
	DOUBLE_COLON_BLOCK            = &PrimitiveBlock{Content: "::", BlockType: DoubleColonBlockType}
	PRINT_BLOCK                   = &PrimitiveBlock{Content: "print", BlockType: PrintBlockType}
	NEW_BLOCK                     = &PrimitiveBlock{Content: "new", BlockType: NewBlockType}
	DOUBLE_MINUS_BLOCK            = &PrimitiveBlock{Content: "--", BlockType: DoubleMinusBlockType}
	DOUBLE_PLUS_BLOCK             = &PrimitiveBlock{Content: "++", BlockType: DoublePlusBlockType}
	AMPERSAND_BLOCK               = &PrimitiveBlock{Content: "&", BlockType: AmpersandBlockType}
	EXPONENT_EQUAL_BLOCK          = &PrimitiveBlock{Content: "**=", BlockType: ExponentEqualBlockType}
	MULTIPLY_EQUAL_BLOCK          = &PrimitiveBlock{Content: "*=", BlockType: MultiplyEqualBlockType}
	DIVIDE_EQUAL_BLOCK            = &PrimitiveBlock{Content: "/=", BlockType: DivideEqualBlockType}
	MODULO_EQUAL_BLOCK            = &PrimitiveBlock{Content: "%=", BlockType: ModuloEqualBlockType}
	PLUS_EQUAL_BLOCK              = &PrimitiveBlock{Content: "+=", BlockType: PlusEqualBlockType}
	MINUS_EQUAL_BLOCK             = &PrimitiveBlock{Content: "-=", BlockType: MinusEqualBlockType}
	DOT_EQUAL_BLOCK               = &PrimitiveBlock{Content: ".=", BlockType: DotEqualBlockType}
	LEFT_SHIFT_EQUAL_BLOCK        = &PrimitiveBlock{Content: "<<=", BlockType: LeftShiftEqualBlockType}
	RIGHT_SHIFT_EQUAL_BLOCK       = &PrimitiveBlock{Content: ">>=", BlockType: RightShiftEqualBlockType}
	AND_EQUAL_BLOCK               = &PrimitiveBlock{Content: "&=", BlockType: AndEqualBlockType}
	XOR_EQUAL_BLOCK               = &PrimitiveBlock{Content: "^=", BlockType: XorEqualBlockType}
	OR_EQUAL_BLOCK                = &PrimitiveBlock{Content: "|=", BlockType: OrEqualBlockType}
	NULL_COALESCE_EQUAL_BLOCK     = &PrimitiveBlock{Content: "??=", BlockType: NullCoalesceEqualBlockType}
	MEMBER_ACCESS_BLOCK           = &PrimitiveBlock{Content: "->", BlockType: MemberAccessBlockType}
	MEMBER_ACCESS_NULL_BLOCK      = &PrimitiveBlock{Content: "?->", BlockType: MemberAccessNullBlockType}
	OPEN_BRACKET_BLOCK            = &PrimitiveBlock{Content: "[", BlockType: OpenBracketBlockType}
	CLOSE_BRACKET_BLOCK           = &PrimitiveBlock{Content: "]", BlockType: CloseBracketBlockType}
	SELF_BLOCK                    = &PrimitiveBlock{Content: "self", BlockType: SelfBlockType}
	PARENT_BLOCK                  = &PrimitiveBlock{Content: "parent", BlockType: ParentBlockType}
	TRUE_BLOCK                    = &PrimitiveBlock{Content: "true", BlockType: TrueBlockType}
	FALSE_BLOCK                   = &PrimitiveBlock{Content: "false", BlockType: FalseBlockType}
	OPEN_ATTRIBUTE_BLOCK          = &PrimitiveBlock{Content: "#[", BlockType: OpenAttributeBlockType}
	BYTE_STRING_BLOCK             = &PrimitiveBlock{Content: "b'", BlockType: ByteStringBlockType}
	SINGLE_QUOTE_BLOCK            = &PrimitiveBlock{Content: "'", BlockType: SingleQuoteBlockType}
	DOUBLE_QUOTE_BLOCK            = &PrimitiveBlock{Content: "\"", BlockType: DoubleQuoteBlockType}
	HEREDOC_OPEN_BLOCK            = &PrimitiveBlock{Content: "<<<", BlockType: HeredocOpenBlockType}
	NEWLINE_BLOCK                 = &PrimitiveBlock{Content: "\n", BlockType: NewlineBlockType}
	SHELL_EXEC_BLOCK              = &PrimitiveBlock{Content: "`", BlockType: ShellExecBlockType}
	DOLLAR_BLOCK                  = &PrimitiveBlock{Content: "$", BlockType: DollarBlockType}
	YIELD_BLOCK                   = &PrimitiveBlock{Content: "yield", BlockType: YieldBlockType}
	FROM_BLOCK                    = &PrimitiveBlock{Content: "from", BlockType: FromBlockType}
	INSTANCEOF_BLOCK              = &PrimitiveBlock{Content: "instanceof", BlockType: InstanceofBlockType}
	NULL_COALESCE_BLOCK           = &PrimitiveBlock{Content: "??", BlockType: NullCoalesceBlockType}
	EXPONENT_BLOCK                = &PrimitiveBlock{Content: "**", BlockType: ExponentBlockType}
	AND_BLOCK                     = &PrimitiveBlock{Content: "and", BlockType: AndBlockType}
	OR_BLOCK                      = &PrimitiveBlock{Content: "or", BlockType: OrBlockType}
	XOR_BLOCK                     = &PrimitiveBlock{Content: "xor", BlockType: XorBlockType}
	LOGICAL_OR_BLOCK              = &PrimitiveBlock{Content: "||", BlockType: LogicalOrBlockType}
	LOGICAL_AND_BLOCK             = &PrimitiveBlock{Content: "&&", BlockType: LogicalAndBlockType}
	BITWISE_OR_BLOCK              = &PrimitiveBlock{Content: "|", BlockType: BitwiseOrBlockType}
	BITWISE_XOR_BLOCK             = &PrimitiveBlock{Content: "^", BlockType: BitwiseXorBlockType}
	EQUALITY_BLOCK                = &PrimitiveBlock{Content: "==", BlockType: EqualBlockType}
	NOT_EQUAL_BLOCK               = &PrimitiveBlock{Content: "!=", BlockType: NotEqualBlockType} // Equivalent to "<>"
	IDENTICAL_BLOCK               = &PrimitiveBlock{Content: "===", BlockType: IdenticalBlockType}
	NOT_IDENTICAL_BLOCK           = &PrimitiveBlock{Content: "!==", BlockType: NotIdenticalBlockType}
	LESS_THAN_BLOCK               = &PrimitiveBlock{Content: "<", BlockType: LessThanBlockType}
	GREATER_THAN_BLOCK            = &PrimitiveBlock{Content: ">", BlockType: GreaterThanBlockType}
	LESS_THAN_OR_EQUAL_BLOCK      = &PrimitiveBlock{Content: "<=", BlockType: LessThanOrEqualBlockType}
	GREATER_THAN_OR_EQUAL_BLOCK   = &PrimitiveBlock{Content: ">=", BlockType: GreaterThanOrEqualBlockType}
	SPACESHIP_BLOCK               = &PrimitiveBlock{Content: "<=>", BlockType: SpaceshipBlockType}
	LEFT_SHIFT_BLOCK              = &PrimitiveBlock{Content: "<<", BlockType: LeftShiftBlockType}
	RIGHT_SHIFT_BLOCK             = &PrimitiveBlock{Content: ">>", BlockType: RightShiftBlockType}
	CONCATENATION_BLOCK           = &PrimitiveBlock{Content: ".", BlockType: ConcatenationBlockType}
	MULTIPLY_BLOCK                = &PrimitiveBlock{Content: "*", BlockType: MultiplyBlockType}
	DIVIDE_BLOCK                  = &PrimitiveBlock{Content: "/", BlockType: DivideBlockType}
	MODULO_BLOCK                  = &PrimitiveBlock{Content: "%", BlockType: ModuloBlockType}
	INCLUDE_BLOCK                 = &PrimitiveBlock{Content: "include", BlockType: IncludeBlockType}
	INCLUDE_ONCE_BLOCK            = &PrimitiveBlock{Content: "include_once", BlockType: IncludeOnceBlockType}
	REQUIRE_BLOCK                 = &PrimitiveBlock{Content: "require", BlockType: RequireBlockType}
	REQUIRE_ONCE_BLOCK            = &PrimitiveBlock{Content: "require_once", BlockType: RequireOnceBlockType}
	HASHTAG_BLOCK                 = &PrimitiveBlock{Content: "#", BlockType: HashtagBlockType}
)

func GetPrimitiveBlockRenders() map[string]func(*utils.Stack, Node) IBlock {
	return map[string]func(*utils.Stack, Node) IBlock{
		"<?php": func(s *utils.Stack, n Node) IBlock {
			return PHP_TAG_BLOCK
		},
		"?>": func(s *utils.Stack, n Node) IBlock {
			return PHP_CLOSE_TAG_BLOCK
		},
		"static": func(s *utils.Stack, n Node) IBlock {
			return STATIC_BLOCK
		},
		"=": func(s *utils.Stack, n Node) IBlock {
			return EQUALS_BLOCK
		},
		"global": func(s *utils.Stack, n Node) IBlock {
			return GLOBAL_BLOCK
		},
		"namespcace": func(s *utils.Stack, n Node) IBlock {
			return NAMESPACE_BLOCK
		},
		"use": func(s *utils.Stack, n Node) IBlock {
			return USE_BLOCK
		},
		"as": func(s *utils.Stack, n Node) IBlock {
			return AS_BLOCK
		},
		"function": func(s *utils.Stack, n Node) IBlock {
			return FUNCTION_BLOCK
		},
		"const": func(s *utils.Stack, n Node) IBlock {
			return CONST_BLOCK
		},
		"\\": func(s *utils.Stack, n Node) IBlock {
			return BACKSLASH_BLOCK
		},
		"{": func(s *utils.Stack, n Node) IBlock {
			return OPEN_BRACE_BLOCK
		},
		"}": func(s *utils.Stack, n Node) IBlock {
			return CLOSE_BRACE_BLOCK
		},
		"trait": func(s *utils.Stack, n Node) IBlock {
			return TRAIT_BLOCK
		},
		"interface": func(s *utils.Stack, n Node) IBlock {
			return INTERFACE_BLOCK
		},
		"extends": func(s *utils.Stack, n Node) IBlock {
			return EXTENDS_BLOCK
		},
		"enum": func(s *utils.Stack, n Node) IBlock {
			return ENUM_BLOCK
		},
		":": func(s *utils.Stack, n Node) IBlock {
			return COLON_BLOCK
		},
		"array": func(s *utils.Stack, n Node) IBlock {
			return ARRAY_TYPE_BLOCK
		},
		"callable": func(s *utils.Stack, n Node) IBlock {
			return CALLABLE_TYPE_BLOCK
		},
		"iterable": func(s *utils.Stack, n Node) IBlock {
			return ITERABLE_TYPE_BLOCK
		},
		"bool": func(s *utils.Stack, n Node) IBlock {
			return BOOL_TYPE_BLOCK
		},
		"float": func(s *utils.Stack, n Node) IBlock {
			return FLOAT_TYPE_BLOCK
		},
		//"string": func(s *utils.Stack, n Node) IBlock {
		//	return STRING_TYPE_BLOCK
		//},
		"int": func(s *utils.Stack, n Node) IBlock {
			return INT_TYPE_BLOCK
		},
		"void": func(s *utils.Stack, n Node) IBlock {
			return VOID_TYPE_BLOCK
		},
		"mixed": func(s *utils.Stack, n Node) IBlock {
			return MIXED_TYPE_BLOCK
		},
		"false": func(s *utils.Stack, n Node) IBlock {
			return FALSE_BLOCK
		},
		"null": func(s *utils.Stack, n Node) IBlock {
			return NULL_TYPE_BLOCK
		},
		"true": func(s *utils.Stack, n Node) IBlock {
			return TRUE_BLOCK
		},
		"case": func(s *utils.Stack, n Node) IBlock {
			return CASE_BLOCK
		},
		"class": func(s *utils.Stack, n Node) IBlock {
			return CLASS_BLOCK
		},
		"final": func(s *utils.Stack, n Node) IBlock {
			return FINAL_BLOCK
		},
		"abstract": func(s *utils.Stack, n Node) IBlock {
			return ABSTRACT_BLOCK
		},
		"readonly": func(s *utils.Stack, n Node) IBlock {
			return READONLY_BLOCK
		},
		"implements": func(s *utils.Stack, n Node) IBlock {
			return IMPLEMENTS_BLOCK
		},
		"=>": func(s *utils.Stack, n Node) IBlock {
			return ARROW_FUNCTION_SEQUENCE_BLOCK
		},
		"var": func(s *utils.Stack, n Node) IBlock {
			return VAR_BLOCK
		},
		"insteadof": func(s *utils.Stack, n Node) IBlock {
			return INSTEAD_OF_BLOCK
		},
		"public": func(s *utils.Stack, n Node) IBlock {
			return PUBLIC_BLOCK
		},
		"protected": func(s *utils.Stack, n Node) IBlock {
			return PROTECTED_BLOCK
		},
		"private": func(s *utils.Stack, n Node) IBlock {
			return PRIVATE_BLOCK
		},
		"(": func(s *utils.Stack, n Node) IBlock {
			return OPEN_PAREN_BLOCK
		},
		")": func(s *utils.Stack, n Node) IBlock {
			return CLOSE_PAREN_BLOCK
		},
		",": func(s *utils.Stack, n Node) IBlock {
			return COMMA_BLOCK
		},
		"fn": func(s *utils.Stack, n Node) IBlock {
			return FN_BLOCK
		},
		"...": func(s *utils.Stack, n Node) IBlock {
			return THREE_DOT_BLOCK
		},
		"?": func(s *utils.Stack, n Node) IBlock {
			return QUESTION_BLOCK
		},
		"never": func(s *utils.Stack, n Node) IBlock {
			return NEVER_BLOCK
		},
		"echo": func(s *utils.Stack, n Node) IBlock {
			return ECHO_BLOCK
		},
		"exit": func(s *utils.Stack, n Node) IBlock {
			return EXIT_BLOCK
		},
		"unset": func(s *utils.Stack, n Node) IBlock {
			return UNSET_BLOCK
		},
		"declare": func(s *utils.Stack, n Node) IBlock {
			return DECLARE_BLOCK
		},
		"enddeclare": func(s *utils.Stack, n Node) IBlock {
			return ENDDECLARE_BLOCK
		},
		"ticks": func(s *utils.Stack, n Node) IBlock {
			return TICKS_BLOCK
		},
		"encoding": func(s *utils.Stack, n Node) IBlock {
			return ENCODING_BLOCK
		},
		"strict_types": func(s *utils.Stack, n Node) IBlock {
			return STRIC_TYPES_BLOCK
		},
		"try": func(s *utils.Stack, n Node) IBlock {
			return TRY_BLOCK
		},
		"catch": func(s *utils.Stack, n Node) IBlock {
			return CATCH_BLOCK
		},
		"finally": func(s *utils.Stack, n Node) IBlock {
			return FINALLY_BLOCK
		},
		"goto": func(s *utils.Stack, n Node) IBlock {
			return GOTO_BLOCK
		},
		"continue": func(s *utils.Stack, n Node) IBlock {
			return CONTINUE_BLOCK
		},
		"break": func(s *utils.Stack, n Node) IBlock {
			return BREAK_BLOCK
		},
		"return": func(s *utils.Stack, n Node) IBlock {
			return RETURN_BLOCK
		},
		"throw": func(s *utils.Stack, n Node) IBlock {
			return THROW_BLOCK
		},
		"while": func(s *utils.Stack, n Node) IBlock {
			return WHILE_BLOCK
		},
		"endwhile": func(s *utils.Stack, n Node) IBlock {
			return ENDWHILE_BLOCK
		},
		"do": func(s *utils.Stack, n Node) IBlock {
			return DO_BLOCK
		},
		"for": func(s *utils.Stack, n Node) IBlock {
			return FOR_BLOCK
		},
		";": func(s *utils.Stack, n Node) IBlock {
			return SEMICOLON_BLOCK
		},
		"endfor": func(s *utils.Stack, n Node) IBlock {
			return ENDFOR_BLOCK
		},
		"foreach": func(s *utils.Stack, n Node) IBlock {
			return FOREACH_BLOCK
		},
		"endforeach": func(s *utils.Stack, n Node) IBlock {
			return ENDFOREACH_BLOCK
		},
		"if": func(s *utils.Stack, n Node) IBlock {
			return IF_BLOCK
		},
		"endif": func(s *utils.Stack, n Node) IBlock {
			return ENDIF_BLOCK
		},
		"elseif": func(s *utils.Stack, n Node) IBlock {
			return ELSEIF_BLOCK
		},
		"else": func(s *utils.Stack, n Node) IBlock {
			return ELSE_BLOCK
		},
		"match": func(s *utils.Stack, n Node) IBlock {
			return MATCH_BLOCK
		},
		"default": func(s *utils.Stack, n Node) IBlock {
			return DEFAULT_BLOCK
		},
		"switch": func(s *utils.Stack, n Node) IBlock {
			return SWITCH_BLOCK
		},
		"endswitch": func(s *utils.Stack, n Node) IBlock {
			return ENDSWITCH_BLOCK
		},
		"+": func(s *utils.Stack, n Node) IBlock {
			return PLUS_BLOCK
		},
		"-": func(s *utils.Stack, n Node) IBlock {
			return MINUS_BLOCK
		},
		"~": func(s *utils.Stack, n Node) IBlock {
			return TILDE_BLOCK
		},
		"!": func(s *utils.Stack, n Node) IBlock {
			return EXCLAMATION_BLOCK
		},
		"@": func(s *utils.Stack, n Node) IBlock {
			return AT_BLOCK
		},
		"clone": func(s *utils.Stack, n Node) IBlock {
			return CLONE_BLOCK
		},
		"::": func(s *utils.Stack, n Node) IBlock {
			return DOUBLE_COLON_BLOCK
		},
		"print": func(s *utils.Stack, n Node) IBlock {
			return PRINT_BLOCK
		},
		"new": func(s *utils.Stack, n Node) IBlock {
			return NEW_BLOCK
		},
		"--": func(s *utils.Stack, n Node) IBlock {
			return DOUBLE_MINUS_BLOCK
		},
		"++": func(s *utils.Stack, n Node) IBlock {
			return DOUBLE_PLUS_BLOCK
		},
		"&": func(s *utils.Stack, n Node) IBlock {
			return AMPERSAND_BLOCK
		},
		"**=": func(s *utils.Stack, n Node) IBlock {
			return EXPONENT_EQUAL_BLOCK
		},
		"*=": func(s *utils.Stack, n Node) IBlock {
			return MULTIPLY_EQUAL_BLOCK
		},
		"/=": func(s *utils.Stack, n Node) IBlock {
			return DIVIDE_EQUAL_BLOCK
		},
		"%=": func(s *utils.Stack, n Node) IBlock {
			return MODULO_EQUAL_BLOCK
		},
		"+=": func(s *utils.Stack, n Node) IBlock {
			return PLUS_EQUAL_BLOCK
		},
		"-=": func(s *utils.Stack, n Node) IBlock {
			return MINUS_EQUAL_BLOCK
		},
		".=": func(s *utils.Stack, n Node) IBlock {
			return DOT_EQUAL_BLOCK
		},
		"<<=": func(s *utils.Stack, n Node) IBlock {
			return LEFT_SHIFT_EQUAL_BLOCK
		},
		">>=": func(s *utils.Stack, n Node) IBlock {
			return RIGHT_SHIFT_EQUAL_BLOCK
		},
		"&=": func(s *utils.Stack, n Node) IBlock {
			return AND_EQUAL_BLOCK
		},
		"^=": func(s *utils.Stack, n Node) IBlock {
			return XOR_EQUAL_BLOCK
		},
		"|=": func(s *utils.Stack, n Node) IBlock {
			return OR_EQUAL_BLOCK
		},
		"??=": func(s *utils.Stack, n Node) IBlock {
			return NULL_COALESCE_EQUAL_BLOCK
		},
		"->": func(s *utils.Stack, n Node) IBlock {
			return MEMBER_ACCESS_BLOCK
		},
		"?->": func(s *utils.Stack, n Node) IBlock {
			return MEMBER_ACCESS_NULL_BLOCK
		},
		"[": func(s *utils.Stack, n Node) IBlock {
			return OPEN_BRACKET_BLOCK
		},
		"]": func(s *utils.Stack, n Node) IBlock {
			return CLOSE_BRACKET_BLOCK
		},
		"self": func(s *utils.Stack, n Node) IBlock {
			return SELF_BLOCK
		},
		"parent": func(s *utils.Stack, n Node) IBlock {
			return PARENT_BLOCK
		},
		"#[": func(s *utils.Stack, n Node) IBlock {
			return OPEN_ATTRIBUTE_BLOCK
		},
		"b'": func(s *utils.Stack, n Node) IBlock {
			return BYTE_STRING_BLOCK
		},
		"'": func(s *utils.Stack, n Node) IBlock {
			return SINGLE_QUOTE_BLOCK
		},
		"\"": func(s *utils.Stack, n Node) IBlock {
			return DOUBLE_QUOTE_BLOCK
		},
		"<<<": func(s *utils.Stack, n Node) IBlock {
			return HEREDOC_OPEN_BLOCK
		},
		"\n": func(s *utils.Stack, n Node) IBlock {
			return NEWLINE_BLOCK
		},
		"`": func(s *utils.Stack, n Node) IBlock {
			return SHELL_EXEC_BLOCK
		},
		"$": func(s *utils.Stack, n Node) IBlock {
			return DOLLAR_BLOCK
		},
		"yield": func(s *utils.Stack, n Node) IBlock {
			return YIELD_BLOCK
		},
		"from": func(s *utils.Stack, n Node) IBlock {
			return FROM_BLOCK
		},
		"instanceof": func(s *utils.Stack, n Node) IBlock {
			return INSTANCEOF_BLOCK
		},
		"??": func(s *utils.Stack, n Node) IBlock {
			return NULL_COALESCE_BLOCK
		},
		"**": func(s *utils.Stack, n Node) IBlock {
			return EXPONENT_BLOCK
		},
		"and": func(s *utils.Stack, n Node) IBlock {
			return AND_BLOCK
		},
		"or": func(s *utils.Stack, n Node) IBlock {
			return OR_BLOCK
		},
		"xor": func(s *utils.Stack, n Node) IBlock {
			return XOR_BLOCK
		},
		"||": func(s *utils.Stack, n Node) IBlock {
			return LOGICAL_OR_BLOCK
		},
		"&&": func(s *utils.Stack, n Node) IBlock {
			return LOGICAL_AND_BLOCK
		},
		"|": func(s *utils.Stack, n Node) IBlock {
			return BITWISE_OR_BLOCK
		},
		"^": func(s *utils.Stack, n Node) IBlock {
			return BITWISE_XOR_BLOCK
		},

		"==": func(s *utils.Stack, n Node) IBlock {
			return EQUALITY_BLOCK
		},
		"!=": func(s *utils.Stack, n Node) IBlock {
			return NOT_EQUAL_BLOCK
		},
		"<>": func(s *utils.Stack, n Node) IBlock {
			return NOT_EQUAL_BLOCK
		},
		"===": func(s *utils.Stack, n Node) IBlock {
			return IDENTICAL_BLOCK
		},
		"!==": func(s *utils.Stack, n Node) IBlock {
			return NOT_IDENTICAL_BLOCK
		},
		"<": func(s *utils.Stack, n Node) IBlock {
			return LESS_THAN_BLOCK
		},
		">": func(s *utils.Stack, n Node) IBlock {
			return GREATER_THAN_BLOCK
		},
		"<=": func(s *utils.Stack, n Node) IBlock {
			return LESS_THAN_OR_EQUAL_BLOCK
		},
		">=": func(s *utils.Stack, n Node) IBlock {
			return GREATER_THAN_OR_EQUAL_BLOCK
		},
		"<=>": func(s *utils.Stack, n Node) IBlock {
			return SPACESHIP_BLOCK
		},
		"<<": func(s *utils.Stack, n Node) IBlock {
			return LEFT_SHIFT_BLOCK
		},
		">>": func(s *utils.Stack, n Node) IBlock {
			return RIGHT_SHIFT_BLOCK
		},
		".": func(s *utils.Stack, n Node) IBlock {
			return CONCATENATION_BLOCK
		},
		"*": func(s *utils.Stack, n Node) IBlock {
			return MULTIPLY_BLOCK
		},
		"/": func(s *utils.Stack, n Node) IBlock {
			return DIVIDE_BLOCK
		},
		"%": func(s *utils.Stack, n Node) IBlock {
			return MODULO_BLOCK
		},
		"include": func(s *utils.Stack, n Node) IBlock {
			return INCLUDE_BLOCK
		},
		"include_once": func(s *utils.Stack, n Node) IBlock {
			return INCLUDE_ONCE_BLOCK
		},
		"require": func(s *utils.Stack, n Node) IBlock {
			return REQUIRE_BLOCK
		},
		"require_once": func(s *utils.Stack, n Node) IBlock {
			return REQUIRE_ONCE_BLOCK
		},
		"#": func(s *utils.Stack, n Node) IBlock {
			return HASHTAG_BLOCK
		},
	}
}
