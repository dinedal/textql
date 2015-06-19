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

const yyNprod = 205
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 618

var yyAct = []int{

	94, 85, 160, 368, 91, 335, 297, 92, 62, 163,
	250, 289, 241, 90, 201, 211, 63, 376, 80, 162,
	3, 137, 136, 376, 81, 103, 376, 50, 179, 262,
	263, 264, 265, 266, 348, 267, 268, 187, 65, 130,
	232, 70, 64, 295, 73, 315, 317, 53, 77, 28,
	29, 30, 31, 51, 52, 230, 130, 130, 86, 232,
	38, 76, 40, 378, 68, 257, 41, 124, 346, 377,
	120, 43, 375, 44, 319, 316, 46, 47, 48, 128,
	14, 15, 16, 17, 133, 324, 321, 345, 344, 294,
	159, 161, 164, 326, 69, 72, 165, 49, 121, 231,
	45, 123, 283, 281, 242, 233, 287, 242, 18, 352,
	273, 135, 173, 65, 117, 113, 65, 64, 183, 182,
	64, 177, 119, 169, 144, 145, 146, 147, 148, 149,
	150, 151, 181, 86, 220, 206, 183, 136, 208, 209,
	71, 210, 341, 290, 218, 219, 184, 222, 223, 224,
	225, 226, 227, 228, 229, 199, 197, 205, 198, 253,
	19, 20, 22, 21, 23, 213, 149, 150, 151, 234,
	86, 86, 115, 24, 25, 26, 65, 65, 221, 207,
	64, 248, 252, 290, 246, 239, 236, 238, 254, 193,
	137, 136, 127, 322, 245, 144, 145, 146, 147, 148,
	149, 150, 151, 343, 342, 14, 249, 313, 191, 137,
	136, 312, 194, 311, 180, 234, 255, 115, 272, 276,
	277, 309, 274, 59, 328, 102, 310, 307, 108, 205,
	232, 275, 308, 131, 353, 280, 66, 99, 100, 101,
	86, 213, 330, 284, 288, 260, 166, 14, 75, 214,
	106, 204, 292, 286, 116, 212, 282, 361, 293, 115,
	296, 203, 190, 192, 189, 147, 148, 149, 150, 151,
	175, 355, 356, 104, 105, 305, 306, 360, 204, 130,
	109, 323, 111, 176, 359, 114, 205, 205, 203, 166,
	327, 28, 29, 30, 31, 107, 65, 78, 325, 171,
	331, 333, 336, 332, 170, 339, 168, 234, 167, 110,
	271, 71, 134, 338, 340, 144, 145, 146, 147, 148,
	149, 150, 151, 66, 320, 347, 318, 270, 337, 71,
	301, 349, 300, 262, 263, 264, 265, 266, 351, 267,
	268, 196, 195, 358, 278, 357, 144, 145, 146, 147,
	148, 149, 150, 151, 178, 363, 336, 125, 122, 364,
	369, 369, 369, 65, 370, 371, 237, 64, 97, 118,
	60, 372, 79, 102, 74, 379, 108, 380, 102, 366,
	381, 350, 112, 98, 84, 99, 100, 101, 329, 14,
	99, 100, 101, 58, 89, 279, 374, 367, 106, 144,
	145, 146, 147, 148, 149, 150, 151, 215, 32, 216,
	217, 185, 244, 126, 56, 54, 298, 88, 304, 299,
	251, 104, 105, 82, 34, 35, 36, 37, 109, 97,
	303, 259, 180, 61, 102, 373, 362, 108, 14, 33,
	186, 39, 256, 107, 98, 84, 99, 100, 101, 235,
	188, 42, 67, 247, 174, 89, 365, 354, 334, 106,
	302, 258, 285, 172, 240, 14, 96, 93, 95, 291,
	243, 138, 87, 314, 202, 261, 129, 200, 88, 83,
	97, 269, 104, 105, 82, 102, 132, 55, 108, 109,
	27, 57, 13, 12, 11, 98, 66, 99, 100, 101,
	97, 10, 9, 8, 107, 102, 89, 7, 108, 6,
	106, 5, 4, 2, 1, 98, 66, 99, 100, 101,
	0, 0, 0, 0, 0, 102, 89, 0, 108, 88,
	106, 0, 0, 104, 105, 0, 66, 99, 100, 101,
	109, 0, 0, 0, 0, 0, 166, 0, 0, 88,
	106, 0, 0, 104, 105, 107, 0, 0, 0, 0,
	109, 139, 143, 141, 142, 0, 0, 0, 0, 0,
	0, 0, 0, 104, 105, 107, 0, 0, 0, 0,
	109, 0, 155, 156, 157, 158, 0, 152, 153, 154,
	0, 0, 0, 0, 0, 107, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 140,
	144, 145, 146, 147, 148, 149, 150, 151,
}
var yyPact = []int{

	75, -1000, -1000, 240, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -30,
	-21, 10, -14, 7, -1000, -1000, -1000, 433, 398, -1000,
	-1000, -1000, 396, -1000, 364, 334, 424, 287, -31, 3,
	275, -1000, 5, 275, -1000, 338, -34, 275, -34, 336,
	-1000, -1000, -1000, -1000, -1000, 409, -1000, 268, 334, 349,
	37, 334, 162, -1000, 207, -1000, 36, 333, 53, 275,
	-1000, -1000, 322, -1000, -26, 321, 393, 126, 275, -1000,
	224, -1000, -1000, 293, 33, 123, 540, -1000, 480, 460,
	-1000, -1000, -1000, 500, 262, 260, -1000, 258, 253, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 500,
	-1000, 237, 287, 318, 422, 287, 500, 275, -1000, 391,
	-60, -1000, 176, -1000, 306, -1000, -1000, 305, -1000, 422,
	409, 215, -1000, -1000, 275, 104, 480, 480, 500, 209,
	386, 500, 500, 109, 500, 500, 500, 500, 500, 500,
	500, 500, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	540, -46, -2, 4, 540, -1000, 200, 348, 409, -1000,
	433, 353, 26, 329, 384, 287, 287, 204, -1000, 407,
	480, -1000, 329, -1000, -1000, -1000, 93, 275, -1000, -28,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 420, -1000,
	190, 277, 291, 242, 32, -1000, -1000, -1000, -1000, 69,
	329, -1000, 200, -1000, -1000, 209, 500, 500, 329, 276,
	-1000, 370, 192, 192, 192, 91, 91, -1000, -1000, -1000,
	-1000, -1000, 500, -1000, 329, -1000, 2, 409, 1, 188,
	23, -1000, 480, 77, 243, 240, 117, -12, -1000, 407,
	401, 405, 123, 296, -1000, -1000, 294, -1000, 418, 404,
	215, 215, -1000, -1000, 171, 165, 157, 155, 151, -19,
	-1000, 290, -27, 288, -15, -1000, 329, 125, 500, -1000,
	329, -1000, -16, -1000, 353, 9, -1000, 500, 142, -1000,
	358, 187, -1000, -1000, -1000, 287, 401, -1000, 500, 500,
	-1000, -1000, 407, 480, 500, 277, 76, -1000, 148, -1000,
	147, -1000, -1000, -1000, -1000, -3, -4, -23, -1000, -1000,
	-1000, -1000, 500, 329, -1000, -67, -1000, 329, 500, 350,
	243, -1000, -1000, 54, 179, -1000, 245, -1000, 401, 123,
	175, 480, -1000, -1000, 238, 231, 211, 329, -1000, 329,
	429, -1000, 500, 500, -1000, -1000, -1000, 363, 123, 275,
	275, 275, 287, 329, -1000, -1000, 428, 375, -29, -1000,
	-32, -38, 162, -1000, 275, -1000, 275, -1000, -1000, 275,
	-1000, -1000,
}
var yyPgo = []int{

	0, 514, 513, 19, 512, 511, 509, 507, 503, 502,
	501, 494, 493, 492, 408, 491, 490, 487, 18, 24,
	486, 481, 479, 477, 14, 476, 475, 474, 223, 473,
	3, 28, 1, 472, 471, 470, 13, 2, 15, 9,
	469, 7, 468, 25, 467, 4, 466, 464, 12, 463,
	462, 461, 460, 10, 458, 5, 457, 6, 456, 454,
	453, 11, 8, 16, 248, 452, 451, 450, 442, 441,
	440, 0, 27, 439,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 3, 4, 4, 5, 6, 7,
	8, 8, 8, 9, 9, 9, 10, 11, 11, 11,
	12, 13, 13, 13, 73, 14, 15, 15, 16, 16,
	16, 16, 16, 17, 17, 18, 18, 19, 19, 19,
	22, 22, 20, 20, 20, 23, 23, 24, 24, 24,
	24, 21, 21, 21, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 27, 27, 27, 28, 28, 29, 29,
	29, 29, 30, 30, 25, 25, 31, 31, 32, 32,
	32, 32, 32, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 34, 34, 34, 34, 34, 34,
	34, 38, 38, 38, 43, 39, 39, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 37, 42, 42, 44, 44, 44, 46,
	49, 49, 47, 47, 48, 50, 50, 45, 45, 36,
	36, 36, 36, 51, 51, 52, 52, 53, 53, 54,
	54, 55, 56, 56, 56, 57, 57, 57, 58, 58,
	58, 59, 59, 60, 60, 61, 61, 35, 35, 40,
	40, 41, 41, 62, 62, 63, 64, 64, 65, 65,
	66, 66, 67, 67, 67, 67, 67, 68, 68, 69,
	69, 70, 70, 71, 72,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 11, 3, 7, 7, 8, 7, 3,
	5, 8, 4, 6, 7, 4, 5, 4, 5, 5,
	3, 2, 2, 2, 0, 2, 0, 2, 1, 2,
	1, 1, 1, 0, 1, 1, 3, 1, 2, 3,
	1, 1, 0, 1, 2, 1, 3, 3, 3, 3,
	5, 0, 1, 2, 1, 1, 2, 3, 2, 3,
	2, 2, 2, 1, 3, 1, 1, 3, 0, 5,
	5, 5, 1, 3, 0, 2, 0, 2, 1, 3,
	3, 2, 3, 3, 3, 4, 3, 4, 5, 6,
	3, 4, 2, 6, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 1, 3, 1, 3, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 3,
	4, 5, 4, 1, 1, 1, 1, 1, 1, 5,
	0, 1, 1, 2, 4, 0, 2, 1, 3, 1,
	1, 1, 1, 0, 3, 0, 2, 0, 3, 1,
	3, 2, 0, 1, 1, 0, 2, 4, 0, 2,
	4, 0, 3, 1, 3, 0, 5, 2, 1, 1,
	3, 3, 1, 1, 3, 3, 0, 2, 0, 3,
	0, 1, 1, 1, 1, 1, 1, 0, 1, 0,
	1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, 5, 6, 7, 8, 33, 85,
	86, 88, 87, 89, 98, 99, 100, -16, 51, 52,
	53, 54, -14, -73, -14, -14, -14, -14, 90, -69,
	92, 96, -66, 92, 94, 90, 90, 91, 92, 90,
	-72, -72, -72, -3, 17, -17, 18, -15, 29, -28,
	36, 9, -62, -63, -45, -71, 36, -65, 95, 91,
	-71, 36, 90, -71, 36, -64, 95, -71, -64, 36,
	-18, -19, 75, -22, 36, -32, -37, -33, 69, 46,
	-36, -45, -41, -44, -71, -42, -46, 20, 35, 37,
	38, 39, 25, -43, 73, 74, 50, 95, 28, 80,
	41, -28, 33, 78, -28, 55, 47, 78, 36, 69,
	-71, -72, 36, -72, 93, 36, 20, 66, -71, -25,
	55, 9, -20, -71, 19, 78, 68, 67, -34, 21,
	69, 23, 24, 22, 70, 71, 72, 73, 74, 75,
	76, 77, 47, 48, 49, 42, 43, 44, 45, -32,
	-37, -32, -3, -39, -37, -37, 46, 46, 46, -43,
	46, 46, -49, -37, -59, 33, 46, -62, 36, -31,
	10, -63, -37, -71, -72, 20, -70, 97, -67, 88,
	86, 32, 87, 13, 36, 36, 36, -72, -31, -19,
	-23, -24, -27, 46, 36, -43, -71, 75, -32, -32,
	-37, -38, 46, -43, 40, 21, 23, 24, -37, -37,
	25, 69, -37, -37, -37, -37, -37, -37, -37, -37,
	101, 101, 55, 101, -37, 101, -18, 18, -18, -36,
	-47, -48, 81, -35, 28, -3, -62, -60, -45, -31,
	-53, 13, -32, 66, -71, -72, -68, 93, -51, 11,
	55, -26, 56, 57, 58, 59, 60, 62, 63, -21,
	36, 19, -24, 78, -39, -38, -37, -37, 68, 25,
	-37, 101, -18, 101, 55, -50, -48, 83, -32, -61,
	66, -40, -41, -61, 101, 55, -53, -57, 15, 14,
	36, 36, -52, 12, 14, -24, -24, 56, 61, 56,
	61, 56, 56, 56, -29, 64, 94, 65, 36, 101,
	36, 101, 68, -37, 101, -36, 84, -37, 82, 30,
	55, -45, -57, -37, -54, -55, -37, -72, -53, -32,
	-39, 66, 56, 56, 91, 91, 91, -37, 101, -37,
	31, -41, 55, 55, -56, 26, 27, -57, -32, 46,
	46, 46, 7, -37, -55, -58, 16, 34, -30, -71,
	-30, -30, -62, 7, 21, 101, 55, 101, 101, -71,
	-71, -71,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 34, 34, 34, 34, 34, 199,
	190, 0, 0, 0, 204, 204, 204, 0, 38, 40,
	41, 42, 43, 36, 0, 0, 0, 0, 188, 0,
	0, 200, 0, 0, 191, 0, 186, 0, 186, 0,
	31, 32, 33, 14, 39, 0, 44, 35, 0, 0,
	76, 0, 19, 183, 0, 147, 203, 0, 0, 0,
	204, 203, 0, 204, 0, 0, 0, 0, 0, 30,
	84, 45, 47, 52, 203, 50, 51, 88, 0, 0,
	117, 118, 119, 0, 147, 0, 133, 0, 0, 149,
	150, 151, 152, 182, 136, 137, 138, 134, 135, 140,
	37, 171, 0, 0, 86, 0, 0, 0, 204, 0,
	201, 22, 0, 25, 0, 27, 187, 0, 204, 86,
	0, 0, 48, 53, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 104, 105, 106, 107, 108, 109, 110, 91,
	0, 0, 0, 0, 115, 128, 0, 0, 0, 102,
	0, 0, 0, 141, 0, 0, 0, 86, 77, 157,
	0, 184, 185, 148, 20, 189, 0, 0, 204, 197,
	192, 193, 194, 195, 196, 26, 28, 29, 153, 46,
	85, 55, 61, 0, 73, 75, 54, 49, 89, 90,
	93, 94, 0, 112, 113, 0, 0, 0, 96, 0,
	100, 0, 120, 121, 122, 123, 124, 125, 126, 127,
	92, 114, 0, 181, 115, 129, 0, 0, 0, 0,
	145, 142, 0, 175, 0, 178, 175, 0, 173, 157,
	165, 0, 87, 0, 202, 23, 0, 198, 155, 0,
	0, 0, 64, 65, 0, 0, 0, 0, 0, 78,
	62, 0, 0, 0, 0, 95, 97, 0, 0, 101,
	116, 130, 0, 132, 0, 0, 143, 0, 0, 15,
	0, 177, 179, 16, 172, 0, 165, 18, 0, 0,
	204, 24, 157, 0, 0, 56, 59, 66, 0, 68,
	0, 70, 71, 72, 57, 0, 0, 0, 63, 58,
	74, 111, 0, 98, 131, 0, 139, 146, 0, 0,
	0, 174, 17, 166, 158, 159, 162, 21, 165, 156,
	154, 0, 67, 69, 0, 0, 0, 99, 103, 144,
	0, 180, 0, 0, 161, 163, 164, 168, 60, 0,
	0, 0, 0, 167, 160, 13, 0, 0, 0, 82,
	0, 0, 176, 169, 0, 79, 0, 80, 81, 0,
	83, 170,
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
		//line sql.y:152
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:158
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 3:
		yyVAL.statement = yyS[yypt-0].statement
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
		//line sql.y:174
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-9].bytes2), Distinct: yyS[yypt-8].str, SelectExprs: yyS[yypt-7].selectExprs, From: NewFrom(AST_FROM, yyS[yypt-6].tableExprs), Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 14:
		//line sql.y:178
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 15:
		//line sql.y:184
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 16:
		//line sql.y:188
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 17:
		//line sql.y:200
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 18:
		//line sql.y:206
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 19:
		//line sql.y:212
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 20:
		//line sql.y:218
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 21:
		//line sql.y:222
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 22:
		//line sql.y:227
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 23:
		//line sql.y:233
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 24:
		//line sql.y:237
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 25:
		//line sql.y:242
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 26:
		//line sql.y:248
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 27:
		//line sql.y:254
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 28:
		//line sql.y:258
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 29:
		//line sql.y:263
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 30:
		//line sql.y:269
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 31:
		//line sql.y:275
		{
			yyVAL.statement = &Other{}
		}
	case 32:
		//line sql.y:279
		{
			yyVAL.statement = &Other{}
		}
	case 33:
		//line sql.y:283
		{
			yyVAL.statement = &Other{}
		}
	case 34:
		//line sql.y:288
		{
			SetAllowComments(yylex, true)
		}
	case 35:
		//line sql.y:292
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 36:
		//line sql.y:298
		{
			yyVAL.bytes2 = nil
		}
	case 37:
		//line sql.y:302
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 38:
		//line sql.y:308
		{
			yyVAL.str = AST_UNION
		}
	case 39:
		//line sql.y:312
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 40:
		//line sql.y:316
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 41:
		//line sql.y:320
		{
			yyVAL.str = AST_EXCEPT
		}
	case 42:
		//line sql.y:324
		{
			yyVAL.str = AST_INTERSECT
		}
	case 43:
		//line sql.y:329
		{
			yyVAL.str = ""
		}
	case 44:
		//line sql.y:333
		{
			yyVAL.str = AST_DISTINCT
		}
	case 45:
		//line sql.y:339
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 46:
		//line sql.y:343
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 47:
		//line sql.y:349
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 48:
		//line sql.y:353
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 49:
		//line sql.y:357
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 50:
		//line sql.y:363
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 51:
		//line sql.y:367
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 52:
		//line sql.y:372
		{
			yyVAL.bytes = nil
		}
	case 53:
		//line sql.y:376
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 54:
		//line sql.y:380
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 55:
		//line sql.y:386
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 56:
		//line sql.y:390
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 57:
		//line sql.y:396
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 58:
		//line sql.y:400
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 59:
		//line sql.y:404
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 60:
		//line sql.y:408
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 61:
		//line sql.y:413
		{
			yyVAL.bytes = nil
		}
	case 62:
		//line sql.y:417
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 63:
		//line sql.y:421
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 64:
		//line sql.y:427
		{
			yyVAL.str = AST_JOIN
		}
	case 65:
		//line sql.y:431
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 66:
		//line sql.y:435
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 67:
		//line sql.y:439
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 68:
		//line sql.y:443
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 69:
		//line sql.y:447
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 70:
		//line sql.y:451
		{
			yyVAL.str = AST_JOIN
		}
	case 71:
		//line sql.y:455
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 72:
		//line sql.y:459
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 73:
		//line sql.y:465
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 74:
		//line sql.y:469
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 75:
		//line sql.y:473
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 76:
		//line sql.y:479
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 77:
		//line sql.y:483
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 78:
		//line sql.y:488
		{
			yyVAL.indexHints = nil
		}
	case 79:
		//line sql.y:492
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 80:
		//line sql.y:496
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 81:
		//line sql.y:500
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 82:
		//line sql.y:506
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 83:
		//line sql.y:510
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 84:
		//line sql.y:515
		{
			yyVAL.tableExprs = nil
		}
	case 85:
		//line sql.y:519
		{
			yyVAL.tableExprs = yyS[yypt-0].tableExprs
		}
	case 86:
		//line sql.y:525
		{
			yyVAL.boolExpr = nil
		}
	case 87:
		//line sql.y:529
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 88:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 89:
		//line sql.y:536
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 90:
		//line sql.y:540
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 91:
		//line sql.y:544
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 92:
		//line sql.y:548
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 93:
		//line sql.y:554
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 94:
		//line sql.y:558
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].colTuple}
		}
	case 95:
		//line sql.y:562
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].colTuple}
		}
	case 96:
		//line sql.y:566
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 97:
		//line sql.y:570
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 98:
		//line sql.y:574
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 99:
		//line sql.y:578
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 100:
		//line sql.y:582
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 101:
		//line sql.y:586
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 102:
		//line sql.y:590
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 103:
		//line sql.y:594
		{
			yyVAL.boolExpr = &KeyrangeExpr{Start: yyS[yypt-3].valExpr, End: yyS[yypt-1].valExpr}
		}
	case 104:
		//line sql.y:600
		{
			yyVAL.str = AST_EQ
		}
	case 105:
		//line sql.y:604
		{
			yyVAL.str = AST_LT
		}
	case 106:
		//line sql.y:608
		{
			yyVAL.str = AST_GT
		}
	case 107:
		//line sql.y:612
		{
			yyVAL.str = AST_LE
		}
	case 108:
		//line sql.y:616
		{
			yyVAL.str = AST_GE
		}
	case 109:
		//line sql.y:620
		{
			yyVAL.str = AST_NE
		}
	case 110:
		//line sql.y:624
		{
			yyVAL.str = AST_NSE
		}
	case 111:
		//line sql.y:630
		{
			yyVAL.colTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 112:
		//line sql.y:634
		{
			yyVAL.colTuple = yyS[yypt-0].subquery
		}
	case 113:
		//line sql.y:638
		{
			yyVAL.colTuple = ListArg(yyS[yypt-0].bytes)
		}
	case 114:
		//line sql.y:644
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 115:
		//line sql.y:650
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 116:
		//line sql.y:654
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 117:
		//line sql.y:660
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 118:
		//line sql.y:664
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 119:
		//line sql.y:668
		{
			yyVAL.valExpr = yyS[yypt-0].rowTuple
		}
	case 120:
		//line sql.y:672
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 121:
		//line sql.y:676
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 122:
		//line sql.y:680
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 123:
		//line sql.y:684
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 124:
		//line sql.y:688
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 125:
		//line sql.y:692
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 126:
		//line sql.y:696
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 127:
		//line sql.y:700
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 128:
		//line sql.y:704
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
	case 129:
		//line sql.y:719
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 130:
		//line sql.y:723
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 131:
		//line sql.y:727
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 132:
		//line sql.y:731
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 133:
		//line sql.y:735
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 134:
		//line sql.y:741
		{
			yyVAL.bytes = IF_BYTES
		}
	case 135:
		//line sql.y:745
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 136:
		//line sql.y:751
		{
			yyVAL.byt = AST_UPLUS
		}
	case 137:
		//line sql.y:755
		{
			yyVAL.byt = AST_UMINUS
		}
	case 138:
		//line sql.y:759
		{
			yyVAL.byt = AST_TILDA
		}
	case 139:
		//line sql.y:765
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 140:
		//line sql.y:770
		{
			yyVAL.valExpr = nil
		}
	case 141:
		//line sql.y:774
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 142:
		//line sql.y:780
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 143:
		//line sql.y:784
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 144:
		//line sql.y:790
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 145:
		//line sql.y:795
		{
			yyVAL.valExpr = nil
		}
	case 146:
		//line sql.y:799
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 147:
		//line sql.y:805
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 148:
		//line sql.y:809
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 149:
		//line sql.y:815
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 150:
		//line sql.y:819
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 151:
		//line sql.y:823
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 152:
		//line sql.y:827
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 153:
		//line sql.y:832
		{
			yyVAL.valExprs = nil
		}
	case 154:
		//line sql.y:836
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 155:
		//line sql.y:841
		{
			yyVAL.boolExpr = nil
		}
	case 156:
		//line sql.y:845
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 157:
		//line sql.y:850
		{
			yyVAL.orderBy = nil
		}
	case 158:
		//line sql.y:854
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 159:
		//line sql.y:860
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 160:
		//line sql.y:864
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 161:
		//line sql.y:870
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 162:
		//line sql.y:875
		{
			yyVAL.str = AST_ASC
		}
	case 163:
		//line sql.y:879
		{
			yyVAL.str = AST_ASC
		}
	case 164:
		//line sql.y:883
		{
			yyVAL.str = AST_DESC
		}
	case 165:
		//line sql.y:888
		{
			yyVAL.limit = nil
		}
	case 166:
		//line sql.y:892
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 167:
		//line sql.y:896
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 168:
		//line sql.y:901
		{
			yyVAL.str = ""
		}
	case 169:
		//line sql.y:905
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 170:
		//line sql.y:909
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
	case 171:
		//line sql.y:922
		{
			yyVAL.columns = nil
		}
	case 172:
		//line sql.y:926
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 173:
		//line sql.y:932
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 174:
		//line sql.y:936
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 175:
		//line sql.y:941
		{
			yyVAL.updateExprs = nil
		}
	case 176:
		//line sql.y:945
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 177:
		//line sql.y:951
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 178:
		//line sql.y:955
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 179:
		//line sql.y:961
		{
			yyVAL.values = Values{yyS[yypt-0].rowTuple}
		}
	case 180:
		//line sql.y:965
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].rowTuple)
		}
	case 181:
		//line sql.y:971
		{
			yyVAL.rowTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 182:
		//line sql.y:975
		{
			yyVAL.rowTuple = yyS[yypt-0].subquery
		}
	case 183:
		//line sql.y:981
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 184:
		//line sql.y:985
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 185:
		//line sql.y:991
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 186:
		//line sql.y:996
		{
			yyVAL.empty = struct{}{}
		}
	case 187:
		//line sql.y:998
		{
			yyVAL.empty = struct{}{}
		}
	case 188:
		//line sql.y:1001
		{
			yyVAL.empty = struct{}{}
		}
	case 189:
		//line sql.y:1003
		{
			yyVAL.empty = struct{}{}
		}
	case 190:
		//line sql.y:1006
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		//line sql.y:1008
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		//line sql.y:1012
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		//line sql.y:1014
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		//line sql.y:1016
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1018
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1020
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1023
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1025
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1028
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1030
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1033
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1035
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1039
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 204:
		//line sql.y:1044
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
