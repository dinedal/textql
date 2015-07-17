//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const IN = 57363
const IS = 57364
const LIKE = 57365
const BETWEEN = 57366
const NULL = 57367
const ASC = 57368
const DESC = 57369
const VALUES = 57370
const INTO = 57371
const DUPLICATE = 57372
const KEY = 57373
const DEFAULT = 57374
const SET = 57375
const LOCK = 57376
const KEYRANGE = 57377
const ID = 57378
const STRING = 57379
const NUMBER = 57380
const VALUE_ARG = 57381
const LIST_ARG = 57382
const COMMENT = 57383
const LE = 57384
const GE = 57385
const NE = 57386
const NULL_SAFE_EQUAL = 57387
const UNION = 57388
const MINUS = 57389
const EXCEPT = 57390
const INTERSECT = 57391
const JOIN = 57392
const STRAIGHT_JOIN = 57393
const LEFT = 57394
const RIGHT = 57395
const INNER = 57396
const OUTER = 57397
const CROSS = 57398
const NATURAL = 57399
const USE = 57400
const FORCE = 57401
const ON = 57402
const OR = 57403
const AND = 57404
const NOT = 57405
const UNARY = 57406
const CASE = 57407
const WHEN = 57408
const THEN = 57409
const ELSE = 57410
const END = 57411
const CREATE = 57412
const ALTER = 57413
const DROP = 57414
const RENAME = 57415
const ANALYZE = 57416
const TABLE = 57417
const INDEX = 57418
const VIEW = 57419
const TO = 57420
const IGNORE = 57421
const IF = 57422
const UNIQUE = 57423
const USING = 57424
const SHOW = 57425
const DESCRIBE = 57426
const EXPLAIN = 57427

var yyToknames = []string{
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"KEYRANGE",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
	"'~'",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"OR",
	"AND",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 207
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 678

var yyAct = []int{

	51, 122, 141, 363, 359, 342, 49, 48, 83, 314,
	366, 280, 126, 334, 225, 159, 142, 42, 43, 125,
	3, 163, 205, 174, 332, 60, 47, 90, 138, 381,
	381, 37, 31, 32, 33, 34, 84, 85, 99, 98,
	268, 94, 15, 17, 18, 19, 153, 146, 127, 381,
	86, 340, 128, 91, 196, 29, 320, 91, 91, 219,
	71, 196, 73, 121, 124, 76, 74, 77, 136, 330,
	20, 144, 194, 38, 148, 383, 382, 150, 143, 329,
	132, 154, 195, 328, 147, 230, 231, 232, 233, 234,
	43, 235, 236, 43, 149, 380, 168, 339, 170, 302,
	299, 82, 173, 251, 249, 181, 182, 197, 185, 186,
	187, 188, 189, 190, 191, 192, 171, 172, 167, 304,
	157, 78, 21, 22, 24, 23, 25, 176, 293, 295,
	297, 198, 43, 43, 206, 26, 27, 28, 79, 80,
	81, 206, 183, 255, 241, 99, 98, 152, 215, 109,
	110, 111, 112, 113, 214, 209, 223, 216, 294, 218,
	306, 203, 200, 202, 193, 161, 207, 97, 96, 210,
	274, 111, 112, 113, 99, 98, 98, 198, 227, 211,
	335, 244, 245, 327, 325, 224, 184, 240, 242, 272,
	335, 167, 316, 275, 170, 222, 287, 326, 248, 291,
	285, 288, 243, 43, 176, 286, 290, 169, 160, 144,
	289, 260, 144, 211, 264, 373, 143, 354, 196, 143,
	252, 228, 265, 212, 256, 92, 262, 254, 263, 155,
	129, 352, 278, 250, 258, 230, 231, 232, 233, 234,
	279, 235, 236, 271, 273, 270, 166, 259, 301, 351,
	283, 284, 15, 211, 167, 167, 165, 305, 350, 144,
	144, 310, 31, 32, 33, 34, 143, 312, 177, 317,
	134, 91, 133, 131, 175, 130, 313, 309, 318, 303,
	158, 239, 96, 166, 198, 123, 346, 345, 298, 95,
	322, 59, 296, 165, 321, 324, 277, 276, 238, 323,
	261, 220, 331, 56, 57, 58, 96, 217, 333, 213,
	139, 156, 151, 16, 364, 337, 208, 370, 353, 15,
	137, 247, 379, 341, 338, 266, 29, 178, 348, 179,
	180, 221, 365, 347, 87, 343, 344, 282, 315, 281,
	226, 144, 308, 349, 160, 357, 360, 356, 355, 88,
	140, 367, 367, 367, 361, 384, 378, 362, 15, 36,
	267, 371, 368, 369, 72, 319, 269, 75, 377, 145,
	311, 257, 374, 358, 385, 360, 375, 376, 386, 35,
	388, 387, 389, 253, 201, 144, 54, 390, 135, 391,
	204, 59, 143, 53, 65, 50, 52, 67, 68, 69,
	70, 55, 41, 56, 57, 58, 336, 307, 100, 44,
	292, 164, 46, 229, 162, 40, 63, 237, 93, 30,
	106, 107, 108, 109, 110, 111, 112, 113, 89, 14,
	13, 12, 11, 10, 9, 45, 8, 7, 6, 61,
	62, 39, 5, 4, 2, 1, 66, 54, 0, 0,
	0, 0, 59, 0, 0, 65, 0, 0, 0, 0,
	0, 64, 55, 41, 56, 57, 58, 199, 0, 372,
	0, 0, 0, 46, 0, 0, 0, 63, 0, 0,
	0, 0, 0, 15, 106, 107, 108, 109, 110, 111,
	112, 113, 0, 0, 0, 0, 45, 0, 54, 0,
	61, 62, 39, 59, 0, 0, 65, 66, 0, 0,
	0, 0, 0, 55, 123, 56, 57, 58, 54, 0,
	0, 0, 64, 59, 46, 0, 65, 0, 63, 0,
	0, 0, 0, 55, 123, 56, 57, 58, 0, 0,
	0, 0, 0, 15, 46, 0, 0, 45, 63, 0,
	0, 61, 62, 0, 0, 0, 0, 0, 66, 0,
	0, 0, 0, 59, 0, 0, 65, 45, 0, 0,
	0, 61, 62, 64, 123, 56, 57, 58, 66, 0,
	0, 0, 0, 59, 129, 0, 65, 0, 63, 0,
	0, 0, 0, 64, 123, 56, 57, 58, 0, 0,
	0, 0, 0, 0, 129, 0, 0, 0, 63, 0,
	0, 61, 62, 101, 105, 103, 104, 300, 66, 106,
	107, 108, 109, 110, 111, 112, 113, 0, 0, 0,
	0, 61, 62, 64, 117, 118, 119, 120, 66, 114,
	115, 116, 246, 0, 106, 107, 108, 109, 110, 111,
	112, 113, 0, 64, 0, 0, 0, 0, 0, 0,
	0, 102, 106, 107, 108, 109, 110, 111, 112, 113,
	106, 107, 108, 109, 110, 111, 112, 113,
}
var yyPact = []int{

	37, -1000, -1000, 211, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 427, -1000, -1000, -1000,
	-1000, -30, -27, 31, 48, 11, -1000, -1000, -1000, -1000,
	353, 317, -1000, -1000, -1000, 308, -1000, 216, -1000, -1000,
	270, 89, 107, 592, -1000, 498, 478, -1000, -1000, -1000,
	558, 229, 227, -1000, 226, 224, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 558, 291, 274, 341,
	249, -48, -7, 246, -1000, 4, 246, -1000, 276, -49,
	246, -49, 275, -1000, -1000, -1000, -1000, -1000, 427, 239,
	334, 427, 210, -1000, -1000, 246, -1000, 132, 498, 498,
	558, 228, 306, 558, 558, 117, 558, 558, 558, 558,
	558, 558, 558, 558, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 592, 86, -29, -19, 6, 592, -1000, 538,
	366, 427, -1000, 353, 266, 53, 600, 274, 283, 77,
	274, 158, -1000, 176, -1000, 273, 85, 246, -1000, 271,
	-1000, -34, 265, 311, 129, 246, -1000, 216, -1000, 329,
	498, -1000, 166, 179, 262, 247, 66, -1000, -1000, -1000,
	-1000, -1000, 108, 600, -1000, 538, -1000, -1000, 228, 558,
	558, 600, 574, -1000, 296, 76, 76, 76, 96, 96,
	-1000, -1000, -1000, 246, -1000, -1000, 558, -1000, 600, -1000,
	3, 427, 2, 165, 60, -1000, 498, 201, 249, 264,
	334, 249, 558, -1000, 305, -57, -1000, 157, -1000, 261,
	-1000, -1000, 260, -1000, 334, 327, 323, 107, 210, 210,
	-1000, -1000, 144, 140, 154, 150, 143, 64, -1000, 256,
	29, 252, -1, -1000, 600, 549, 558, -1000, 600, -1000,
	-2, -1000, 266, 35, -1000, 558, 78, 314, 249, 249,
	198, -1000, 325, -1000, 600, -1000, -1000, 126, 246, -1000,
	-37, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 329,
	325, 498, 558, 179, 118, -1000, 141, -1000, 127, -1000,
	-1000, -1000, -1000, -8, -12, -22, -1000, -1000, -1000, -1000,
	558, 600, -1000, -77, -1000, 600, 558, 114, 184, 211,
	124, -4, -1000, 325, 320, 322, 251, -1000, -1000, 250,
	-1000, 327, 320, 107, 163, 498, -1000, -1000, 212, 203,
	185, 600, -1000, 600, -1000, 288, 162, -1000, -1000, -1000,
	249, 320, -1000, 558, 558, -1000, -1000, 325, 298, 107,
	246, 246, 246, 286, 184, -1000, -1000, 414, 160, -1000,
	350, -1000, 320, -1000, 349, 301, -6, -1000, -25, -26,
	348, -1000, 558, 558, -1000, -1000, -1000, 298, -1000, 246,
	-1000, 246, -1000, -1000, 249, 600, -1000, -1000, 246, -1000,
	158, -1000,
}
var yyPgo = []int{

	0, 445, 444, 19, 443, 442, 438, 437, 436, 434,
	433, 432, 431, 430, 429, 379, 428, 419, 313, 31,
	73, 418, 417, 415, 414, 21, 27, 413, 411, 28,
	410, 10, 15, 17, 409, 408, 407, 26, 1, 23,
	12, 406, 6, 396, 25, 395, 7, 393, 390, 22,
	388, 383, 14, 11, 9, 373, 4, 372, 5, 3,
	371, 370, 13, 2, 16, 147, 369, 367, 366, 365,
	364, 360, 0, 8, 359,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 4, 3, 3, 5, 5, 6,
	7, 8, 9, 9, 9, 10, 10, 10, 11, 12,
	12, 12, 13, 14, 14, 14, 74, 15, 16, 16,
	17, 17, 17, 17, 17, 18, 18, 19, 19, 20,
	20, 20, 23, 23, 21, 21, 21, 24, 24, 25,
	25, 25, 25, 22, 22, 22, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 28, 28, 28, 29, 29,
	30, 30, 30, 30, 31, 31, 26, 26, 32, 32,
	33, 33, 33, 33, 33, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 35, 35, 35, 35,
	35, 35, 35, 39, 39, 39, 44, 40, 40, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 43, 43, 45, 45,
	45, 47, 50, 50, 48, 48, 49, 51, 51, 46,
	46, 37, 37, 37, 37, 52, 52, 53, 53, 54,
	54, 55, 55, 56, 57, 57, 57, 58, 58, 58,
	59, 59, 59, 60, 60, 61, 61, 62, 62, 36,
	36, 41, 41, 42, 42, 63, 63, 64, 65, 65,
	66, 66, 67, 67, 68, 68, 68, 68, 68, 69,
	69, 70, 70, 71, 71, 72, 73,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 9, 11, 3, 7, 7, 8,
	7, 3, 5, 8, 4, 6, 7, 4, 5, 4,
	5, 5, 3, 2, 2, 2, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 3,
	3, 3, 5, 0, 1, 2, 1, 1, 2, 3,
	2, 3, 2, 2, 2, 1, 3, 1, 1, 3,
	0, 5, 5, 5, 1, 3, 0, 2, 0, 2,
	1, 3, 3, 2, 3, 3, 3, 4, 3, 4,
	5, 6, 3, 4, 2, 6, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 1, 3, 1, 3, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 3, 4, 5, 4, 1, 1, 1, 1, 1,
	1, 5, 0, 1, 1, 2, 4, 0, 2, 1,
	3, 1, 1, 1, 1, 0, 3, 0, 2, 0,
	3, 1, 3, 2, 0, 1, 1, 0, 2, 4,
	0, 2, 4, 0, 3, 1, 3, 0, 5, 2,
	1, 1, 3, 3, 1, 1, 3, 3, 0, 2,
	0, 3, 0, 1, 1, 1, 1, 1, 1, 0,
	1, 0, 1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, -14, 5, -18, 6, 7, 8,
	33, 85, 86, 88, 87, 89, 98, 99, 100, 18,
	-17, 51, 52, 53, 54, -15, -74, -19, -20, 75,
	-23, 36, -33, -38, -34, 69, 46, -37, -46, -42,
	-45, -72, -43, -47, 20, 35, 37, 38, 39, 25,
	-44, 73, 74, 50, 95, 28, 80, -15, -15, -15,
	-15, 90, -70, 92, 96, -67, 92, 94, 90, 90,
	91, 92, 90, -73, -73, -73, -3, 17, -18, -16,
	-26, 55, 9, -21, -72, 19, 36, 78, 68, 67,
	-35, 21, 69, 23, 24, 22, 70, 71, 72, 73,
	74, 75, 76, 77, 47, 48, 49, 42, 43, 44,
	45, -33, -38, 36, -33, -3, -40, -38, -38, 46,
	46, 46, -44, 46, 46, -50, -38, 29, -29, 36,
	9, -63, -64, -46, -72, -66, 95, 91, -72, 90,
	-72, 36, -65, 95, -72, -65, 36, -19, 41, -32,
	10, -20, -24, -25, -28, 46, 36, -44, -72, 75,
	-72, -33, -33, -38, -39, 46, -44, 40, 21, 23,
	24, -38, -38, 25, 69, -38, -38, -38, -38, -38,
	-38, -38, -38, 78, 101, 101, 55, 101, -38, 101,
	-19, 18, -19, -37, -48, -49, 81, -29, 33, 78,
	-29, 55, 47, 36, 69, -72, -73, 36, -73, 93,
	36, 20, 66, -72, -26, -52, 11, -33, 55, -27,
	56, 57, 58, 59, 60, 62, 63, -22, 36, 19,
	-25, 78, -40, -39, -38, -38, 68, 25, -38, 101,
	-19, 101, 55, -51, -49, 83, -33, -60, 33, 46,
	-63, 36, -32, -64, -38, -73, 20, -71, 97, -68,
	88, 86, 32, 87, 13, 36, 36, 36, -73, -32,
	-53, 12, 14, -25, -25, 56, 61, 56, 61, 56,
	56, 56, -30, 64, 94, 65, 36, 101, 36, 101,
	68, -38, 101, -37, 84, -38, 82, -36, 28, -3,
	-63, -61, -46, -32, -54, 13, 66, -72, -73, -69,
	93, -52, -54, -33, -40, 66, 56, 56, 91, 91,
	91, -38, 101, -38, -62, 66, -41, -42, -62, 101,
	55, -54, -58, 15, 14, 36, 36, -53, -58, -33,
	46, 46, 46, 30, 55, -46, -58, -38, -55, -56,
	-38, -73, -54, -59, 16, 34, -31, -72, -31, -31,
	31, -42, 55, 55, -57, 26, 27, -58, 7, 21,
	101, 55, 101, 101, 7, -38, -56, -59, -72, -72,
	-63, -72,
}
var yyDef = []int{

	45, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 36, 0, 36, 36, 36,
	36, 201, 192, 0, 0, 0, 206, 206, 206, 46,
	0, 40, 42, 43, 44, 45, 38, 86, 47, 49,
	54, 205, 52, 53, 90, 0, 0, 119, 120, 121,
	0, 149, 0, 135, 0, 0, 151, 152, 153, 154,
	184, 138, 139, 140, 136, 137, 142, 0, 0, 0,
	0, 190, 0, 0, 202, 0, 0, 193, 0, 188,
	0, 188, 0, 33, 34, 35, 16, 41, 0, 37,
	88, 0, 0, 50, 55, 0, 205, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 106, 107, 108, 109, 110, 111,
	112, 93, 0, 205, 0, 0, 0, 117, 130, 0,
	0, 0, 104, 0, 0, 0, 143, 0, 0, 78,
	0, 21, 185, 0, 149, 0, 0, 0, 206, 0,
	206, 0, 0, 0, 0, 0, 32, 86, 39, 155,
	0, 48, 87, 57, 63, 0, 75, 77, 56, 51,
	150, 91, 92, 95, 96, 0, 114, 115, 0, 0,
	0, 98, 0, 102, 0, 122, 123, 124, 125, 126,
	127, 128, 129, 0, 94, 116, 0, 183, 117, 131,
	0, 0, 0, 0, 147, 144, 0, 173, 0, 0,
	88, 0, 0, 206, 0, 203, 24, 0, 27, 0,
	29, 189, 0, 206, 88, 157, 0, 89, 0, 0,
	66, 67, 0, 0, 0, 0, 0, 80, 64, 0,
	0, 0, 0, 97, 99, 0, 0, 103, 118, 132,
	0, 134, 0, 0, 145, 0, 0, 0, 0, 0,
	88, 79, 159, 186, 187, 22, 191, 0, 0, 206,
	199, 194, 195, 196, 197, 198, 28, 30, 31, 155,
	159, 0, 0, 58, 61, 68, 0, 70, 0, 72,
	73, 74, 59, 0, 0, 0, 65, 60, 76, 113,
	0, 100, 133, 0, 141, 148, 0, 177, 0, 180,
	177, 0, 175, 159, 167, 0, 0, 204, 25, 0,
	200, 157, 167, 158, 156, 0, 69, 71, 0, 0,
	0, 101, 105, 146, 17, 0, 179, 181, 18, 174,
	0, 167, 20, 0, 0, 206, 26, 159, 170, 62,
	0, 0, 0, 0, 0, 176, 19, 168, 160, 161,
	164, 23, 167, 14, 0, 0, 0, 84, 0, 0,
	0, 182, 0, 0, 163, 165, 166, 170, 171, 0,
	81, 0, 82, 83, 0, 169, 162, 15, 0, 85,
	178, 172,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 77, 70, 3,
	46, 101, 75, 73, 55, 74, 78, 76, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	48, 47, 49, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 72, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 71, 3, 50,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 51, 52, 53, 54, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line sql.y:154
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:160
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 3:
		//line sql.y:164
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 4:
		yyVAL.statement = yyS[yypt-0].statement
	case 5:
		yyVAL.statement = yyS[yypt-0].statement
	case 6:
		yyVAL.statement = yyS[yypt-0].statement
	case 7:
		yyVAL.statement = yyS[yypt-0].statement
	case 8:
		yyVAL.statement = yyS[yypt-0].statement
	case 9:
		yyVAL.statement = yyS[yypt-0].statement
	case 10:
		yyVAL.statement = yyS[yypt-0].statement
	case 11:
		yyVAL.statement = yyS[yypt-0].statement
	case 12:
		yyVAL.statement = yyS[yypt-0].statement
	case 13:
		yyVAL.statement = yyS[yypt-0].statement
	case 14:
		//line sql.y:180
		{
			yyVAL.selStmt = &Select{Comments: nil, Distinct: yyS[yypt-8].str, SelectExprs: yyS[yypt-7].selectExprs, From: NewFrom(AST_FROM, yyS[yypt-6].tableExprs), Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 15:
		//line sql.y:186
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-9].bytes2), Distinct: yyS[yypt-8].str, SelectExprs: yyS[yypt-7].selectExprs, From: NewFrom(AST_FROM, yyS[yypt-6].tableExprs), Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 16:
		//line sql.y:190
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 17:
		//line sql.y:196
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 18:
		//line sql.y:200
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 19:
		//line sql.y:212
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 20:
		//line sql.y:218
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 21:
		//line sql.y:224
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 22:
		//line sql.y:230
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 23:
		//line sql.y:234
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 24:
		//line sql.y:239
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 25:
		//line sql.y:245
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 26:
		//line sql.y:249
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 27:
		//line sql.y:254
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 28:
		//line sql.y:260
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 29:
		//line sql.y:266
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 30:
		//line sql.y:270
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 31:
		//line sql.y:275
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 32:
		//line sql.y:281
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 33:
		//line sql.y:287
		{
			yyVAL.statement = &Other{}
		}
	case 34:
		//line sql.y:291
		{
			yyVAL.statement = &Other{}
		}
	case 35:
		//line sql.y:295
		{
			yyVAL.statement = &Other{}
		}
	case 36:
		//line sql.y:300
		{
			SetAllowComments(yylex, true)
		}
	case 37:
		//line sql.y:304
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 38:
		//line sql.y:310
		{
			yyVAL.bytes2 = nil
		}
	case 39:
		//line sql.y:314
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 40:
		//line sql.y:320
		{
			yyVAL.str = AST_UNION
		}
	case 41:
		//line sql.y:324
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 42:
		//line sql.y:328
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 43:
		//line sql.y:332
		{
			yyVAL.str = AST_EXCEPT
		}
	case 44:
		//line sql.y:336
		{
			yyVAL.str = AST_INTERSECT
		}
	case 45:
		//line sql.y:341
		{
			yyVAL.str = ""
		}
	case 46:
		//line sql.y:345
		{
			yyVAL.str = AST_DISTINCT
		}
	case 47:
		//line sql.y:351
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 48:
		//line sql.y:355
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 49:
		//line sql.y:361
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 50:
		//line sql.y:365
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 51:
		//line sql.y:369
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 52:
		//line sql.y:375
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 53:
		//line sql.y:379
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 54:
		//line sql.y:384
		{
			yyVAL.bytes = nil
		}
	case 55:
		//line sql.y:388
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 56:
		//line sql.y:392
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 57:
		//line sql.y:398
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 58:
		//line sql.y:402
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 59:
		//line sql.y:408
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 60:
		//line sql.y:412
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 61:
		//line sql.y:416
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 62:
		//line sql.y:420
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 63:
		//line sql.y:425
		{
			yyVAL.bytes = nil
		}
	case 64:
		//line sql.y:429
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 65:
		//line sql.y:433
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 66:
		//line sql.y:439
		{
			yyVAL.str = AST_JOIN
		}
	case 67:
		//line sql.y:443
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 68:
		//line sql.y:447
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 69:
		//line sql.y:451
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 70:
		//line sql.y:455
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 71:
		//line sql.y:459
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 72:
		//line sql.y:463
		{
			yyVAL.str = AST_JOIN
		}
	case 73:
		//line sql.y:467
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 74:
		//line sql.y:471
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 75:
		//line sql.y:477
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 76:
		//line sql.y:481
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 77:
		//line sql.y:485
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 78:
		//line sql.y:491
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 79:
		//line sql.y:495
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 80:
		//line sql.y:500
		{
			yyVAL.indexHints = nil
		}
	case 81:
		//line sql.y:504
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 82:
		//line sql.y:508
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 83:
		//line sql.y:512
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 84:
		//line sql.y:518
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 85:
		//line sql.y:522
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 86:
		//line sql.y:527
		{
			yyVAL.tableExprs = nil
		}
	case 87:
		//line sql.y:531
		{
			yyVAL.tableExprs = yyS[yypt-0].tableExprs
		}
	case 88:
		//line sql.y:537
		{
			yyVAL.boolExpr = nil
		}
	case 89:
		//line sql.y:541
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 90:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 91:
		//line sql.y:548
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 92:
		//line sql.y:552
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 93:
		//line sql.y:556
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 94:
		//line sql.y:560
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 95:
		//line sql.y:566
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 96:
		//line sql.y:570
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].colTuple}
		}
	case 97:
		//line sql.y:574
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].colTuple}
		}
	case 98:
		//line sql.y:578
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 99:
		//line sql.y:582
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 100:
		//line sql.y:586
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 101:
		//line sql.y:590
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 102:
		//line sql.y:594
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 103:
		//line sql.y:598
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 104:
		//line sql.y:602
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 105:
		//line sql.y:606
		{
			yyVAL.boolExpr = &KeyrangeExpr{Start: yyS[yypt-3].valExpr, End: yyS[yypt-1].valExpr}
		}
	case 106:
		//line sql.y:612
		{
			yyVAL.str = AST_EQ
		}
	case 107:
		//line sql.y:616
		{
			yyVAL.str = AST_LT
		}
	case 108:
		//line sql.y:620
		{
			yyVAL.str = AST_GT
		}
	case 109:
		//line sql.y:624
		{
			yyVAL.str = AST_LE
		}
	case 110:
		//line sql.y:628
		{
			yyVAL.str = AST_GE
		}
	case 111:
		//line sql.y:632
		{
			yyVAL.str = AST_NE
		}
	case 112:
		//line sql.y:636
		{
			yyVAL.str = AST_NSE
		}
	case 113:
		//line sql.y:642
		{
			yyVAL.colTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 114:
		//line sql.y:646
		{
			yyVAL.colTuple = yyS[yypt-0].subquery
		}
	case 115:
		//line sql.y:650
		{
			yyVAL.colTuple = ListArg(yyS[yypt-0].bytes)
		}
	case 116:
		//line sql.y:656
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 117:
		//line sql.y:662
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 118:
		//line sql.y:666
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 119:
		//line sql.y:672
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 120:
		//line sql.y:676
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 121:
		//line sql.y:680
		{
			yyVAL.valExpr = yyS[yypt-0].rowTuple
		}
	case 122:
		//line sql.y:684
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 123:
		//line sql.y:688
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 124:
		//line sql.y:692
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 125:
		//line sql.y:696
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 126:
		//line sql.y:700
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 127:
		//line sql.y:704
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 128:
		//line sql.y:708
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 129:
		//line sql.y:712
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 130:
		//line sql.y:716
		{
			if num, ok := yyS[yypt-0].valExpr.(NumVal); ok {
				switch yyS[yypt-1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
			}
		}
	case 131:
		//line sql.y:731
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 132:
		//line sql.y:735
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 133:
		//line sql.y:739
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 134:
		//line sql.y:743
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 135:
		//line sql.y:747
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 136:
		//line sql.y:753
		{
			yyVAL.bytes = IF_BYTES
		}
	case 137:
		//line sql.y:757
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 138:
		//line sql.y:763
		{
			yyVAL.byt = AST_UPLUS
		}
	case 139:
		//line sql.y:767
		{
			yyVAL.byt = AST_UMINUS
		}
	case 140:
		//line sql.y:771
		{
			yyVAL.byt = AST_TILDA
		}
	case 141:
		//line sql.y:777
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 142:
		//line sql.y:782
		{
			yyVAL.valExpr = nil
		}
	case 143:
		//line sql.y:786
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 144:
		//line sql.y:792
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 145:
		//line sql.y:796
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 146:
		//line sql.y:802
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 147:
		//line sql.y:807
		{
			yyVAL.valExpr = nil
		}
	case 148:
		//line sql.y:811
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 149:
		//line sql.y:817
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 150:
		//line sql.y:821
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 151:
		//line sql.y:827
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 152:
		//line sql.y:831
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 153:
		//line sql.y:835
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 154:
		//line sql.y:839
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 155:
		//line sql.y:844
		{
			yyVAL.valExprs = nil
		}
	case 156:
		//line sql.y:848
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 157:
		//line sql.y:853
		{
			yyVAL.boolExpr = nil
		}
	case 158:
		//line sql.y:857
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 159:
		//line sql.y:862
		{
			yyVAL.orderBy = nil
		}
	case 160:
		//line sql.y:866
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 161:
		//line sql.y:872
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 162:
		//line sql.y:876
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 163:
		//line sql.y:882
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 164:
		//line sql.y:887
		{
			yyVAL.str = AST_ASC
		}
	case 165:
		//line sql.y:891
		{
			yyVAL.str = AST_ASC
		}
	case 166:
		//line sql.y:895
		{
			yyVAL.str = AST_DESC
		}
	case 167:
		//line sql.y:900
		{
			yyVAL.limit = nil
		}
	case 168:
		//line sql.y:904
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 169:
		//line sql.y:908
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 170:
		//line sql.y:913
		{
			yyVAL.str = ""
		}
	case 171:
		//line sql.y:917
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 172:
		//line sql.y:921
		{
			if !bytes.Equal(yyS[yypt-1].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyS[yypt-0].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 173:
		//line sql.y:934
		{
			yyVAL.columns = nil
		}
	case 174:
		//line sql.y:938
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 175:
		//line sql.y:944
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 176:
		//line sql.y:948
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 177:
		//line sql.y:953
		{
			yyVAL.updateExprs = nil
		}
	case 178:
		//line sql.y:957
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 179:
		//line sql.y:963
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 180:
		//line sql.y:967
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 181:
		//line sql.y:973
		{
			yyVAL.values = Values{yyS[yypt-0].rowTuple}
		}
	case 182:
		//line sql.y:977
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].rowTuple)
		}
	case 183:
		//line sql.y:983
		{
			yyVAL.rowTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 184:
		//line sql.y:987
		{
			yyVAL.rowTuple = yyS[yypt-0].subquery
		}
	case 185:
		//line sql.y:993
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 186:
		//line sql.y:997
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 187:
		//line sql.y:1003
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 188:
		//line sql.y:1008
		{
			yyVAL.empty = struct{}{}
		}
	case 189:
		//line sql.y:1010
		{
			yyVAL.empty = struct{}{}
		}
	case 190:
		//line sql.y:1013
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		//line sql.y:1015
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		//line sql.y:1018
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		//line sql.y:1020
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		//line sql.y:1024
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1026
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1028
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1030
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1032
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1035
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1037
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1040
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1042
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1045
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1047
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1051
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 206:
		//line sql.y:1056
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
