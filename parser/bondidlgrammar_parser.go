// Code generated from unofficial-bond-grammar/BondIdlGrammar.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // BondIdlGrammar
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 48, 334,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 105, 10, 2, 12, 2, 14, 2, 108,
	11, 2, 3, 2, 3, 2, 3, 2, 5, 2, 113, 10, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 4, 5, 4, 121, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 135, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	5, 7, 142, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 148, 10, 8, 3, 8, 3, 8,
	5, 8, 152, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 158, 10, 8, 5, 8, 160,
	10, 8, 3, 9, 5, 9, 163, 10, 9, 3, 9, 5, 9, 166, 10, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 5, 9, 172, 10, 9, 3, 10, 3, 10, 3, 10, 5, 10, 177, 10, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	14, 7, 14, 191, 10, 14, 12, 14, 14, 14, 194, 11, 14, 3, 15, 3, 15, 3, 15,
	5, 15, 199, 10, 15, 3, 15, 3, 15, 3, 15, 5, 15, 204, 10, 15, 3, 15, 5,
	15, 207, 10, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 5, 22, 229, 10, 22, 5, 22, 231, 10, 22, 3, 23, 3, 23, 3, 24,
	3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 7, 27, 245,
	10, 27, 12, 27, 14, 27, 248, 11, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 30, 3, 30, 5, 30, 257, 10, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 33, 3, 33, 7, 33, 267, 10, 33, 12, 33, 14, 33, 270, 11, 33, 3, 34,
	3, 34, 3, 34, 3, 35, 3, 35, 5, 35, 277, 10, 35, 3, 36, 3, 36, 3, 37, 3,
	37, 3, 38, 7, 38, 284, 10, 38, 12, 38, 14, 38, 287, 11, 38, 3, 39, 3, 39,
	3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3,
	41, 5, 41, 302, 10, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44,
	3, 44, 7, 44, 312, 10, 44, 12, 44, 14, 44, 315, 11, 44, 3, 45, 3, 45, 5,
	45, 319, 10, 45, 3, 46, 3, 46, 5, 46, 323, 10, 46, 3, 47, 3, 47, 3, 48,
	3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 2, 2, 51, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
	46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
	82, 84, 86, 88, 90, 92, 94, 96, 98, 2, 8, 3, 2, 28, 29, 3, 2, 20, 22, 4,
	2, 19, 19, 29, 29, 6, 2, 19, 19, 24, 25, 27, 27, 29, 30, 4, 2, 11, 13,
	15, 16, 4, 2, 24, 24, 27, 27, 2, 315, 2, 112, 3, 2, 2, 2, 4, 114, 3, 2,
	2, 2, 6, 117, 3, 2, 2, 2, 8, 122, 3, 2, 2, 2, 10, 134, 3, 2, 2, 2, 12,
	136, 3, 2, 2, 2, 14, 159, 3, 2, 2, 2, 16, 171, 3, 2, 2, 2, 18, 173, 3,
	2, 2, 2, 20, 180, 3, 2, 2, 2, 22, 182, 3, 2, 2, 2, 24, 185, 3, 2, 2, 2,
	26, 192, 3, 2, 2, 2, 28, 195, 3, 2, 2, 2, 30, 208, 3, 2, 2, 2, 32, 211,
	3, 2, 2, 2, 34, 213, 3, 2, 2, 2, 36, 215, 3, 2, 2, 2, 38, 217, 3, 2, 2,
	2, 40, 220, 3, 2, 2, 2, 42, 230, 3, 2, 2, 2, 44, 232, 3, 2, 2, 2, 46, 234,
	3, 2, 2, 2, 48, 236, 3, 2, 2, 2, 50, 238, 3, 2, 2, 2, 52, 242, 3, 2, 2,
	2, 54, 249, 3, 2, 2, 2, 56, 252, 3, 2, 2, 2, 58, 254, 3, 2, 2, 2, 60, 258,
	3, 2, 2, 2, 62, 260, 3, 2, 2, 2, 64, 264, 3, 2, 2, 2, 66, 271, 3, 2, 2,
	2, 68, 274, 3, 2, 2, 2, 70, 278, 3, 2, 2, 2, 72, 280, 3, 2, 2, 2, 74, 285,
	3, 2, 2, 2, 76, 288, 3, 2, 2, 2, 78, 292, 3, 2, 2, 2, 80, 297, 3, 2, 2,
	2, 82, 303, 3, 2, 2, 2, 84, 305, 3, 2, 2, 2, 86, 309, 3, 2, 2, 2, 88, 316,
	3, 2, 2, 2, 90, 320, 3, 2, 2, 2, 92, 324, 3, 2, 2, 2, 94, 326, 3, 2, 2,
	2, 96, 329, 3, 2, 2, 2, 98, 331, 3, 2, 2, 2, 100, 101, 5, 6, 4, 2, 101,
	102, 5, 10, 6, 2, 102, 113, 3, 2, 2, 2, 103, 105, 5, 4, 3, 2, 104, 103,
	3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2,
	2, 2, 107, 109, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 110, 5, 6, 4, 2,
	110, 111, 5, 10, 6, 2, 111, 113, 3, 2, 2, 2, 112, 100, 3, 2, 2, 2, 112,
	106, 3, 2, 2, 2, 113, 3, 3, 2, 2, 2, 114, 115, 7, 23, 2, 2, 115, 116, 7,
	30, 2, 2, 116, 5, 3, 2, 2, 2, 117, 118, 7, 7, 2, 2, 118, 120, 5, 8, 5,
	2, 119, 121, 7, 35, 2, 2, 120, 119, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121,
	7, 3, 2, 2, 2, 122, 123, 9, 2, 2, 2, 123, 9, 3, 2, 2, 2, 124, 125, 5, 14,
	8, 2, 125, 126, 5, 10, 6, 2, 126, 135, 3, 2, 2, 2, 127, 128, 5, 80, 41,
	2, 128, 129, 5, 10, 6, 2, 129, 135, 3, 2, 2, 2, 130, 131, 5, 12, 7, 2,
	131, 132, 5, 10, 6, 2, 132, 135, 3, 2, 2, 2, 133, 135, 7, 2, 2, 3, 134,
	124, 3, 2, 2, 2, 134, 127, 3, 2, 2, 2, 134, 130, 3, 2, 2, 2, 134, 133,
	3, 2, 2, 2, 135, 11, 3, 2, 2, 2, 136, 137, 7, 5, 2, 2, 137, 138, 5, 58,
	30, 2, 138, 139, 7, 3, 2, 2, 139, 141, 5, 42, 22, 2, 140, 142, 7, 35, 2,
	2, 141, 140, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 13, 3, 2, 2, 2, 143,
	144, 5, 74, 38, 2, 144, 145, 7, 8, 2, 2, 145, 147, 5, 58, 30, 2, 146, 148,
	5, 22, 12, 2, 147, 146, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 149, 3,
	2, 2, 2, 149, 151, 5, 24, 13, 2, 150, 152, 7, 35, 2, 2, 151, 150, 3, 2,
	2, 2, 151, 152, 3, 2, 2, 2, 152, 160, 3, 2, 2, 2, 153, 154, 7, 8, 2, 2,
	154, 155, 5, 58, 30, 2, 155, 157, 5, 16, 9, 2, 156, 158, 7, 35, 2, 2, 157,
	156, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 160, 3, 2, 2, 2, 159, 143,
	3, 2, 2, 2, 159, 153, 3, 2, 2, 2, 160, 15, 3, 2, 2, 2, 161, 163, 5, 22,
	12, 2, 162, 161, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 165, 3, 2, 2, 2,
	164, 166, 5, 24, 13, 2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166,
	172, 3, 2, 2, 2, 167, 168, 7, 6, 2, 2, 168, 169, 5, 42, 22, 2, 169, 170,
	5, 18, 10, 2, 170, 172, 3, 2, 2, 2, 171, 162, 3, 2, 2, 2, 171, 167, 3,
	2, 2, 2, 172, 17, 3, 2, 2, 2, 173, 174, 7, 33, 2, 2, 174, 176, 5, 20, 11,
	2, 175, 177, 7, 35, 2, 2, 176, 175, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177,
	178, 3, 2, 2, 2, 178, 179, 7, 34, 2, 2, 179, 19, 3, 2, 2, 2, 180, 181,
	7, 29, 2, 2, 181, 21, 3, 2, 2, 2, 182, 183, 7, 36, 2, 2, 183, 184, 5, 42,
	22, 2, 184, 23, 3, 2, 2, 2, 185, 186, 7, 33, 2, 2, 186, 187, 5, 26, 14,
	2, 187, 188, 7, 34, 2, 2, 188, 25, 3, 2, 2, 2, 189, 191, 5, 28, 15, 2,
	190, 189, 3, 2, 2, 2, 191, 194, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 192,
	193, 3, 2, 2, 2, 193, 27, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 195, 196, 5,
	74, 38, 2, 196, 198, 5, 30, 16, 2, 197, 199, 5, 32, 17, 2, 198, 197, 3,
	2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 201, 5, 34, 18,
	2, 201, 203, 5, 36, 19, 2, 202, 204, 5, 38, 20, 2, 203, 202, 3, 2, 2, 2,
	203, 204, 3, 2, 2, 2, 204, 206, 3, 2, 2, 2, 205, 207, 7, 35, 2, 2, 206,
	205, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 29, 3, 2, 2, 2, 208, 209, 5,
	98, 50, 2, 209, 210, 7, 36, 2, 2, 210, 31, 3, 2, 2, 2, 211, 212, 9, 3,
	2, 2, 212, 33, 3, 2, 2, 2, 213, 214, 5, 42, 22, 2, 214, 35, 3, 2, 2, 2,
	215, 216, 9, 4, 2, 2, 216, 37, 3, 2, 2, 2, 217, 218, 7, 3, 2, 2, 218, 219,
	5, 40, 21, 2, 219, 39, 3, 2, 2, 2, 220, 221, 9, 5, 2, 2, 221, 41, 3, 2,
	2, 2, 222, 231, 5, 44, 23, 2, 223, 224, 5, 46, 24, 2, 224, 225, 5, 50,
	26, 2, 225, 231, 3, 2, 2, 2, 226, 228, 5, 48, 25, 2, 227, 229, 5, 50, 26,
	2, 228, 227, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 231, 3, 2, 2, 2, 230,
	222, 3, 2, 2, 2, 230, 223, 3, 2, 2, 2, 230, 226, 3, 2, 2, 2, 231, 43, 3,
	2, 2, 2, 232, 233, 9, 6, 2, 2, 233, 45, 3, 2, 2, 2, 234, 235, 7, 14, 2,
	2, 235, 47, 3, 2, 2, 2, 236, 237, 9, 2, 2, 2, 237, 49, 3, 2, 2, 2, 238,
	239, 7, 38, 2, 2, 239, 240, 5, 52, 27, 2, 240, 241, 7, 39, 2, 2, 241, 51,
	3, 2, 2, 2, 242, 246, 5, 56, 29, 2, 243, 245, 5, 54, 28, 2, 244, 243, 3,
	2, 2, 2, 245, 248, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2,
	2, 247, 53, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 249, 250, 7, 37, 2, 2, 250,
	251, 5, 56, 29, 2, 251, 55, 3, 2, 2, 2, 252, 253, 5, 42, 22, 2, 253, 57,
	3, 2, 2, 2, 254, 256, 5, 60, 31, 2, 255, 257, 5, 62, 32, 2, 256, 255, 3,
	2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 59, 3, 2, 2, 2, 258, 259, 7, 29, 2,
	2, 259, 61, 3, 2, 2, 2, 260, 261, 7, 38, 2, 2, 261, 262, 5, 64, 33, 2,
	262, 263, 7, 39, 2, 2, 263, 63, 3, 2, 2, 2, 264, 268, 5, 68, 35, 2, 265,
	267, 5, 66, 34, 2, 266, 265, 3, 2, 2, 2, 267, 270, 3, 2, 2, 2, 268, 266,
	3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 65, 3, 2, 2, 2, 270, 268, 3, 2,
	2, 2, 271, 272, 7, 37, 2, 2, 272, 273, 5, 68, 35, 2, 273, 67, 3, 2, 2,
	2, 274, 276, 5, 70, 36, 2, 275, 277, 5, 72, 37, 2, 276, 275, 3, 2, 2, 2,
	276, 277, 3, 2, 2, 2, 277, 69, 3, 2, 2, 2, 278, 279, 7, 29, 2, 2, 279,
	71, 3, 2, 2, 2, 280, 281, 7, 4, 2, 2, 281, 73, 3, 2, 2, 2, 282, 284, 5,
	76, 39, 2, 283, 282, 3, 2, 2, 2, 284, 287, 3, 2, 2, 2, 285, 283, 3, 2,
	2, 2, 285, 286, 3, 2, 2, 2, 286, 75, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2,
	288, 289, 7, 41, 2, 2, 289, 290, 5, 78, 40, 2, 290, 291, 7, 42, 2, 2, 291,
	77, 3, 2, 2, 2, 292, 293, 7, 29, 2, 2, 293, 294, 7, 43, 2, 2, 294, 295,
	7, 30, 2, 2, 295, 296, 7, 44, 2, 2, 296, 79, 3, 2, 2, 2, 297, 298, 7, 9,
	2, 2, 298, 299, 5, 82, 42, 2, 299, 301, 5, 84, 43, 2, 300, 302, 7, 35,
	2, 2, 301, 300, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 81, 3, 2, 2, 2,
	303, 304, 7, 29, 2, 2, 304, 83, 3, 2, 2, 2, 305, 306, 7, 33, 2, 2, 306,
	307, 5, 86, 44, 2, 307, 308, 7, 34, 2, 2, 308, 85, 3, 2, 2, 2, 309, 313,
	5, 90, 46, 2, 310, 312, 5, 88, 45, 2, 311, 310, 3, 2, 2, 2, 312, 315, 3,
	2, 2, 2, 313, 311, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 87, 3, 2, 2,
	2, 315, 313, 3, 2, 2, 2, 316, 318, 7, 37, 2, 2, 317, 319, 5, 86, 44, 2,
	318, 317, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 89, 3, 2, 2, 2, 320, 322,
	5, 92, 47, 2, 321, 323, 5, 94, 48, 2, 322, 321, 3, 2, 2, 2, 322, 323, 3,
	2, 2, 2, 323, 91, 3, 2, 2, 2, 324, 325, 7, 29, 2, 2, 325, 93, 3, 2, 2,
	2, 326, 327, 7, 3, 2, 2, 327, 328, 5, 96, 49, 2, 328, 95, 3, 2, 2, 2, 329,
	330, 5, 98, 50, 2, 330, 97, 3, 2, 2, 2, 331, 332, 9, 7, 2, 2, 332, 99,
	3, 2, 2, 2, 30, 106, 112, 120, 134, 141, 147, 151, 157, 159, 162, 165,
	171, 176, 192, 198, 203, 206, 228, 230, 246, 256, 268, 276, 285, 301, 313,
	318, 322,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "", "", "", "", "", "", "'bool'", "", "", "", "", "'string'",
	"'wstring'", "", "", "'nothing'", "'required'", "'optional'", "'required_optional'",
	"", "", "", "", "", "", "", "", "", "", "'{'", "'}'", "';'", "':'", "','",
	"'<'", "'>'", "'.'", "'['", "']'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "VALUE_CONSTRAINT_STMT", "TYPE_ALIAS_KEYWORD", "VIEW_OF_KEYWORD",
	"NAMESPACE_KEYWORD", "STRUCT_TYPE_KEYWORD", "ENUM_TYPE_KEYWORD", "BOOL_TYPE_KEYWORD",
	"SIGNED_INTEGER_TYPE_KEYWORD", "UNSIGNED_INTEGER_TYPE_KEYWORD", "FLOAT_POINT_TYPE_KEYWORD",
	"CONTAINER_TYPE_KEYWORD", "BYTESTRING_TYPE_KEYWORD", "WSTRING_TYPE_KEYWORD",
	"NON_CONTAINER_TYPE_KEYWORD", "BUILTIN_TYPE_KEYWORD", "NOTHING_KEYWORD",
	"REQUIRED_KEYWORD", "OPTIONAL_KEYWORD", "REQUIRED_OPTIONAL_KEYWORD", "IMPORT_KEYWORD",
	"DEC_NUMBER", "FLOAT_NUMBER", "FLOAT_HALF", "HEX_NUMBER", "MULTI_SECTION_IDENTIFIER",
	"IDENTIFIER", "QUOTED_STRING", "NON_DIGIT", "DIGIT", "BLOCK_BODY_BEGIN",
	"BLOCK_BODY_END", "STMT_END", "FIELD_TYPENAME_DELIMITER", "DEFINITION_DELIMITER",
	"GENERIC_TYPELIST_BEGIN", "GENERIC_TYPELIST_END", "NAMESPACE_DELIMITER",
	"ATTRIBUTE_BEGIN", "ATTRIBUTE_END", "ATTRIBUTE_PARAM_BEGIN", "ATTRIBUTE_PARAM_END",
	"WS", "NEWLINE", "MULTI_LINE_COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"bondIdl", "importDecl", "namespaceDecl", "namespaceName", "bondTypeDef",
	"typeAliasDef", "structDef", "structOrViewDef", "viewOfBody", "singleViewOfField",
	"structBaseClassDef", "structBody", "structFieldDefList", "singleStructFieldDef",
	"fieldID", "fieldModifier", "fieldType", "fieldName", "defaultValueSpec",
	"defaultValues", "typeNameWithGeneric", "builtinPrimitiveType", "builtinContainerType",
	"customType", "genericTypeArgs", "typeArgsList", "moreTypeArgsList", "typeArgName",
	"structNameDeclWithGeneric", "typeDeclName", "genericTypeArgsDecl", "typeParamList",
	"moreTypeParamList", "typeParamDef", "typeParamName", "typeParamValueConstraint",
	"attributeDefList", "attributeDef", "attributeBody", "enumDef", "enumName",
	"enumBody", "enumSymbolDefList", "moreEnumSymbolDef", "singleEnumSymbolDef",
	"enumSymbol", "enumSymbolValueAssignment", "enumSymbolValue", "integerLiteral",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type BondIdlGrammarParser struct {
	*antlr.BaseParser
}

func NewBondIdlGrammarParser(input antlr.TokenStream) *BondIdlGrammarParser {
	this := new(BondIdlGrammarParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "BondIdlGrammar.g4"

	return this
}

// BondIdlGrammarParser tokens.
const (
	BondIdlGrammarParserEOF                           = antlr.TokenEOF
	BondIdlGrammarParserT__0                          = 1
	BondIdlGrammarParserVALUE_CONSTRAINT_STMT         = 2
	BondIdlGrammarParserTYPE_ALIAS_KEYWORD            = 3
	BondIdlGrammarParserVIEW_OF_KEYWORD               = 4
	BondIdlGrammarParserNAMESPACE_KEYWORD             = 5
	BondIdlGrammarParserSTRUCT_TYPE_KEYWORD           = 6
	BondIdlGrammarParserENUM_TYPE_KEYWORD             = 7
	BondIdlGrammarParserBOOL_TYPE_KEYWORD             = 8
	BondIdlGrammarParserSIGNED_INTEGER_TYPE_KEYWORD   = 9
	BondIdlGrammarParserUNSIGNED_INTEGER_TYPE_KEYWORD = 10
	BondIdlGrammarParserFLOAT_POINT_TYPE_KEYWORD      = 11
	BondIdlGrammarParserCONTAINER_TYPE_KEYWORD        = 12
	BondIdlGrammarParserBYTESTRING_TYPE_KEYWORD       = 13
	BondIdlGrammarParserWSTRING_TYPE_KEYWORD          = 14
	BondIdlGrammarParserNON_CONTAINER_TYPE_KEYWORD    = 15
	BondIdlGrammarParserBUILTIN_TYPE_KEYWORD          = 16
	BondIdlGrammarParserNOTHING_KEYWORD               = 17
	BondIdlGrammarParserREQUIRED_KEYWORD              = 18
	BondIdlGrammarParserOPTIONAL_KEYWORD              = 19
	BondIdlGrammarParserREQUIRED_OPTIONAL_KEYWORD     = 20
	BondIdlGrammarParserIMPORT_KEYWORD                = 21
	BondIdlGrammarParserDEC_NUMBER                    = 22
	BondIdlGrammarParserFLOAT_NUMBER                  = 23
	BondIdlGrammarParserFLOAT_HALF                    = 24
	BondIdlGrammarParserHEX_NUMBER                    = 25
	BondIdlGrammarParserMULTI_SECTION_IDENTIFIER      = 26
	BondIdlGrammarParserIDENTIFIER                    = 27
	BondIdlGrammarParserQUOTED_STRING                 = 28
	BondIdlGrammarParserNON_DIGIT                     = 29
	BondIdlGrammarParserDIGIT                         = 30
	BondIdlGrammarParserBLOCK_BODY_BEGIN              = 31
	BondIdlGrammarParserBLOCK_BODY_END                = 32
	BondIdlGrammarParserSTMT_END                      = 33
	BondIdlGrammarParserFIELD_TYPENAME_DELIMITER      = 34
	BondIdlGrammarParserDEFINITION_DELIMITER          = 35
	BondIdlGrammarParserGENERIC_TYPELIST_BEGIN        = 36
	BondIdlGrammarParserGENERIC_TYPELIST_END          = 37
	BondIdlGrammarParserNAMESPACE_DELIMITER           = 38
	BondIdlGrammarParserATTRIBUTE_BEGIN               = 39
	BondIdlGrammarParserATTRIBUTE_END                 = 40
	BondIdlGrammarParserATTRIBUTE_PARAM_BEGIN         = 41
	BondIdlGrammarParserATTRIBUTE_PARAM_END           = 42
	BondIdlGrammarParserWS                            = 43
	BondIdlGrammarParserNEWLINE                       = 44
	BondIdlGrammarParserMULTI_LINE_COMMENT            = 45
	BondIdlGrammarParserLINE_COMMENT                  = 46
)

// BondIdlGrammarParser rules.
const (
	BondIdlGrammarParserRULE_bondIdl                   = 0
	BondIdlGrammarParserRULE_importDecl                = 1
	BondIdlGrammarParserRULE_namespaceDecl             = 2
	BondIdlGrammarParserRULE_namespaceName             = 3
	BondIdlGrammarParserRULE_bondTypeDef               = 4
	BondIdlGrammarParserRULE_typeAliasDef              = 5
	BondIdlGrammarParserRULE_structDef                 = 6
	BondIdlGrammarParserRULE_structOrViewDef           = 7
	BondIdlGrammarParserRULE_viewOfBody                = 8
	BondIdlGrammarParserRULE_singleViewOfField         = 9
	BondIdlGrammarParserRULE_structBaseClassDef        = 10
	BondIdlGrammarParserRULE_structBody                = 11
	BondIdlGrammarParserRULE_structFieldDefList        = 12
	BondIdlGrammarParserRULE_singleStructFieldDef      = 13
	BondIdlGrammarParserRULE_fieldID                   = 14
	BondIdlGrammarParserRULE_fieldModifier             = 15
	BondIdlGrammarParserRULE_fieldType                 = 16
	BondIdlGrammarParserRULE_fieldName                 = 17
	BondIdlGrammarParserRULE_defaultValueSpec          = 18
	BondIdlGrammarParserRULE_defaultValues             = 19
	BondIdlGrammarParserRULE_typeNameWithGeneric       = 20
	BondIdlGrammarParserRULE_builtinPrimitiveType      = 21
	BondIdlGrammarParserRULE_builtinContainerType      = 22
	BondIdlGrammarParserRULE_customType                = 23
	BondIdlGrammarParserRULE_genericTypeArgs           = 24
	BondIdlGrammarParserRULE_typeArgsList              = 25
	BondIdlGrammarParserRULE_moreTypeArgsList          = 26
	BondIdlGrammarParserRULE_typeArgName               = 27
	BondIdlGrammarParserRULE_structNameDeclWithGeneric = 28
	BondIdlGrammarParserRULE_typeDeclName              = 29
	BondIdlGrammarParserRULE_genericTypeArgsDecl       = 30
	BondIdlGrammarParserRULE_typeParamList             = 31
	BondIdlGrammarParserRULE_moreTypeParamList         = 32
	BondIdlGrammarParserRULE_typeParamDef              = 33
	BondIdlGrammarParserRULE_typeParamName             = 34
	BondIdlGrammarParserRULE_typeParamValueConstraint  = 35
	BondIdlGrammarParserRULE_attributeDefList          = 36
	BondIdlGrammarParserRULE_attributeDef              = 37
	BondIdlGrammarParserRULE_attributeBody             = 38
	BondIdlGrammarParserRULE_enumDef                   = 39
	BondIdlGrammarParserRULE_enumName                  = 40
	BondIdlGrammarParserRULE_enumBody                  = 41
	BondIdlGrammarParserRULE_enumSymbolDefList         = 42
	BondIdlGrammarParserRULE_moreEnumSymbolDef         = 43
	BondIdlGrammarParserRULE_singleEnumSymbolDef       = 44
	BondIdlGrammarParserRULE_enumSymbol                = 45
	BondIdlGrammarParserRULE_enumSymbolValueAssignment = 46
	BondIdlGrammarParserRULE_enumSymbolValue           = 47
	BondIdlGrammarParserRULE_integerLiteral            = 48
)

// IBondIdlContext is an interface to support dynamic dispatch.
type IBondIdlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBondIdlContext differentiates from other interfaces.
	IsBondIdlContext()
}

type BondIdlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBondIdlContext() *BondIdlContext {
	var p = new(BondIdlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_bondIdl
	return p
}

func (*BondIdlContext) IsBondIdlContext() {}

func NewBondIdlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BondIdlContext {
	var p = new(BondIdlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_bondIdl

	return p
}

func (s *BondIdlContext) GetParser() antlr.Parser { return s.parser }

func (s *BondIdlContext) NamespaceDecl() INamespaceDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceDeclContext)
}

func (s *BondIdlContext) BondTypeDef() IBondTypeDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBondTypeDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBondTypeDefContext)
}

func (s *BondIdlContext) AllImportDecl() []IImportDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportDeclContext)(nil)).Elem())
	var tst = make([]IImportDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportDeclContext)
		}
	}

	return tst
}

func (s *BondIdlContext) ImportDecl(i int) IImportDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *BondIdlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BondIdlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BondIdlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterBondIdl(s)
	}
}

func (s *BondIdlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitBondIdl(s)
	}
}

func (p *BondIdlGrammarParser) BondIdl() (localctx IBondIdlContext) {
	localctx = NewBondIdlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BondIdlGrammarParserRULE_bondIdl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.NamespaceDecl()
		}
		{
			p.SetState(99)
			p.BondTypeDef()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BondIdlGrammarParserIMPORT_KEYWORD {
			{
				p.SetState(101)
				p.ImportDecl()
			}

			p.SetState(106)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(107)
			p.NamespaceDecl()
		}
		{
			p.SetState(108)
			p.BondTypeDef()
		}

	}

	return localctx
}

// IImportDeclContext is an interface to support dynamic dispatch.
type IImportDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclContext differentiates from other interfaces.
	IsImportDeclContext()
}

type ImportDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclContext() *ImportDeclContext {
	var p = new(ImportDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_importDecl
	return p
}

func (*ImportDeclContext) IsImportDeclContext() {}

func NewImportDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclContext {
	var p = new(ImportDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_importDecl

	return p
}

func (s *ImportDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclContext) IMPORT_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserIMPORT_KEYWORD, 0)
}

func (s *ImportDeclContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserQUOTED_STRING, 0)
}

func (s *ImportDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterImportDecl(s)
	}
}

func (s *ImportDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitImportDecl(s)
	}
}

func (p *BondIdlGrammarParser) ImportDecl() (localctx IImportDeclContext) {
	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BondIdlGrammarParserRULE_importDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(BondIdlGrammarParserIMPORT_KEYWORD)
	}
	{
		p.SetState(113)
		p.Match(BondIdlGrammarParserQUOTED_STRING)
	}

	return localctx
}

// INamespaceDeclContext is an interface to support dynamic dispatch.
type INamespaceDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceDeclContext differentiates from other interfaces.
	IsNamespaceDeclContext()
}

type NamespaceDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceDeclContext() *NamespaceDeclContext {
	var p = new(NamespaceDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_namespaceDecl
	return p
}

func (*NamespaceDeclContext) IsNamespaceDeclContext() {}

func NewNamespaceDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceDeclContext {
	var p = new(NamespaceDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_namespaceDecl

	return p
}

func (s *NamespaceDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceDeclContext) NAMESPACE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserNAMESPACE_KEYWORD, 0)
}

func (s *NamespaceDeclContext) NamespaceName() INamespaceNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceNameContext)
}

func (s *NamespaceDeclContext) STMT_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserSTMT_END, 0)
}

func (s *NamespaceDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterNamespaceDecl(s)
	}
}

func (s *NamespaceDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitNamespaceDecl(s)
	}
}

func (p *BondIdlGrammarParser) NamespaceDecl() (localctx INamespaceDeclContext) {
	localctx = NewNamespaceDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BondIdlGrammarParserRULE_namespaceDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(BondIdlGrammarParserNAMESPACE_KEYWORD)
	}
	{
		p.SetState(116)
		p.NamespaceName()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BondIdlGrammarParserSTMT_END {
		{
			p.SetState(117)
			p.Match(BondIdlGrammarParserSTMT_END)
		}

	}

	return localctx
}

// INamespaceNameContext is an interface to support dynamic dispatch.
type INamespaceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceNameContext differentiates from other interfaces.
	IsNamespaceNameContext()
}

type NamespaceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceNameContext() *NamespaceNameContext {
	var p = new(NamespaceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_namespaceName
	return p
}

func (*NamespaceNameContext) IsNamespaceNameContext() {}

func NewNamespaceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceNameContext {
	var p = new(NamespaceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_namespaceName

	return p
}

func (s *NamespaceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceNameContext) MULTI_SECTION_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserMULTI_SECTION_IDENTIFIER, 0)
}

func (s *NamespaceNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserIDENTIFIER, 0)
}

func (s *NamespaceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterNamespaceName(s)
	}
}

func (s *NamespaceNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitNamespaceName(s)
	}
}

func (p *BondIdlGrammarParser) NamespaceName() (localctx INamespaceNameContext) {
	localctx = NewNamespaceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BondIdlGrammarParserRULE_namespaceName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BondIdlGrammarParserMULTI_SECTION_IDENTIFIER || _la == BondIdlGrammarParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBondTypeDefContext is an interface to support dynamic dispatch.
type IBondTypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBondTypeDefContext differentiates from other interfaces.
	IsBondTypeDefContext()
}

type BondTypeDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBondTypeDefContext() *BondTypeDefContext {
	var p = new(BondTypeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_bondTypeDef
	return p
}

func (*BondTypeDefContext) IsBondTypeDefContext() {}

func NewBondTypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BondTypeDefContext {
	var p = new(BondTypeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_bondTypeDef

	return p
}

func (s *BondTypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *BondTypeDefContext) StructDef() IStructDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructDefContext)
}

func (s *BondTypeDefContext) BondTypeDef() IBondTypeDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBondTypeDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBondTypeDefContext)
}

func (s *BondTypeDefContext) EnumDef() IEnumDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *BondTypeDefContext) TypeAliasDef() ITypeAliasDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAliasDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAliasDefContext)
}

func (s *BondTypeDefContext) EOF() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserEOF, 0)
}

func (s *BondTypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BondTypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BondTypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterBondTypeDef(s)
	}
}

func (s *BondTypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitBondTypeDef(s)
	}
}

func (p *BondIdlGrammarParser) BondTypeDef() (localctx IBondTypeDefContext) {
	localctx = NewBondTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BondIdlGrammarParserRULE_bondTypeDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BondIdlGrammarParserSTRUCT_TYPE_KEYWORD, BondIdlGrammarParserATTRIBUTE_BEGIN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.StructDef()
		}
		{
			p.SetState(123)
			p.BondTypeDef()
		}

	case BondIdlGrammarParserENUM_TYPE_KEYWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.EnumDef()
		}
		{
			p.SetState(126)
			p.BondTypeDef()
		}

	case BondIdlGrammarParserTYPE_ALIAS_KEYWORD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.TypeAliasDef()
		}
		{
			p.SetState(129)
			p.BondTypeDef()
		}

	case BondIdlGrammarParserEOF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			p.Match(BondIdlGrammarParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeAliasDefContext is an interface to support dynamic dispatch.
type ITypeAliasDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeAliasDefContext differentiates from other interfaces.
	IsTypeAliasDefContext()
}

type TypeAliasDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAliasDefContext() *TypeAliasDefContext {
	var p = new(TypeAliasDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_typeAliasDef
	return p
}

func (*TypeAliasDefContext) IsTypeAliasDefContext() {}

func NewTypeAliasDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAliasDefContext {
	var p = new(TypeAliasDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_typeAliasDef

	return p
}

func (s *TypeAliasDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAliasDefContext) TYPE_ALIAS_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserTYPE_ALIAS_KEYWORD, 0)
}

func (s *TypeAliasDefContext) StructNameDeclWithGeneric() IStructNameDeclWithGenericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructNameDeclWithGenericContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructNameDeclWithGenericContext)
}

func (s *TypeAliasDefContext) TypeNameWithGeneric() ITypeNameWithGenericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameWithGenericContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameWithGenericContext)
}

func (s *TypeAliasDefContext) STMT_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserSTMT_END, 0)
}

func (s *TypeAliasDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAliasDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAliasDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterTypeAliasDef(s)
	}
}

func (s *TypeAliasDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitTypeAliasDef(s)
	}
}

func (p *BondIdlGrammarParser) TypeAliasDef() (localctx ITypeAliasDefContext) {
	localctx = NewTypeAliasDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BondIdlGrammarParserRULE_typeAliasDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(BondIdlGrammarParserTYPE_ALIAS_KEYWORD)
	}
	{
		p.SetState(135)
		p.StructNameDeclWithGeneric()
	}
	{
		p.SetState(136)
		p.Match(BondIdlGrammarParserT__0)
	}
	{
		p.SetState(137)
		p.TypeNameWithGeneric()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BondIdlGrammarParserSTMT_END {
		{
			p.SetState(138)
			p.Match(BondIdlGrammarParserSTMT_END)
		}

	}

	return localctx
}

// IStructDefContext is an interface to support dynamic dispatch.
type IStructDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructDefContext differentiates from other interfaces.
	IsStructDefContext()
}

type StructDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDefContext() *StructDefContext {
	var p = new(StructDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_structDef
	return p
}

func (*StructDefContext) IsStructDefContext() {}

func NewStructDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDefContext {
	var p = new(StructDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_structDef

	return p
}

func (s *StructDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDefContext) AttributeDefList() IAttributeDefListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeDefListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeDefListContext)
}

func (s *StructDefContext) STRUCT_TYPE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserSTRUCT_TYPE_KEYWORD, 0)
}

func (s *StructDefContext) StructNameDeclWithGeneric() IStructNameDeclWithGenericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructNameDeclWithGenericContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructNameDeclWithGenericContext)
}

func (s *StructDefContext) StructBody() IStructBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *StructDefContext) StructBaseClassDef() IStructBaseClassDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructBaseClassDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructBaseClassDefContext)
}

func (s *StructDefContext) STMT_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserSTMT_END, 0)
}

func (s *StructDefContext) StructOrViewDef() IStructOrViewDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructOrViewDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructOrViewDefContext)
}

func (s *StructDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterStructDef(s)
	}
}

func (s *StructDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitStructDef(s)
	}
}

func (p *BondIdlGrammarParser) StructDef() (localctx IStructDefContext) {
	localctx = NewStructDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BondIdlGrammarParserRULE_structDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.AttributeDefList()
		}
		{
			p.SetState(142)
			p.Match(BondIdlGrammarParserSTRUCT_TYPE_KEYWORD)
		}
		{
			p.SetState(143)
			p.StructNameDeclWithGeneric()
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BondIdlGrammarParserFIELD_TYPENAME_DELIMITER {
			{
				p.SetState(144)
				p.StructBaseClassDef()
			}

		}
		{
			p.SetState(147)
			p.StructBody()
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BondIdlGrammarParserSTMT_END {
			{
				p.SetState(148)
				p.Match(BondIdlGrammarParserSTMT_END)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.Match(BondIdlGrammarParserSTRUCT_TYPE_KEYWORD)
		}
		{
			p.SetState(152)
			p.StructNameDeclWithGeneric()
		}
		{
			p.SetState(153)
			p.StructOrViewDef()
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BondIdlGrammarParserSTMT_END {
			{
				p.SetState(154)
				p.Match(BondIdlGrammarParserSTMT_END)
			}

		}

	}

	return localctx
}

// IStructOrViewDefContext is an interface to support dynamic dispatch.
type IStructOrViewDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructOrViewDefContext differentiates from other interfaces.
	IsStructOrViewDefContext()
}

type StructOrViewDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructOrViewDefContext() *StructOrViewDefContext {
	var p = new(StructOrViewDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_structOrViewDef
	return p
}

func (*StructOrViewDefContext) IsStructOrViewDefContext() {}

func NewStructOrViewDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructOrViewDefContext {
	var p = new(StructOrViewDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_structOrViewDef

	return p
}

func (s *StructOrViewDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StructOrViewDefContext) StructBaseClassDef() IStructBaseClassDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructBaseClassDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructBaseClassDefContext)
}

func (s *StructOrViewDefContext) StructBody() IStructBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *StructOrViewDefContext) VIEW_OF_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserVIEW_OF_KEYWORD, 0)
}

func (s *StructOrViewDefContext) TypeNameWithGeneric() ITypeNameWithGenericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameWithGenericContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameWithGenericContext)
}

func (s *StructOrViewDefContext) ViewOfBody() IViewOfBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IViewOfBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IViewOfBodyContext)
}

func (s *StructOrViewDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructOrViewDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructOrViewDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterStructOrViewDef(s)
	}
}

func (s *StructOrViewDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitStructOrViewDef(s)
	}
}

func (p *BondIdlGrammarParser) StructOrViewDef() (localctx IStructOrViewDefContext) {
	localctx = NewStructOrViewDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BondIdlGrammarParserRULE_structOrViewDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(169)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BondIdlGrammarParserEOF, BondIdlGrammarParserTYPE_ALIAS_KEYWORD, BondIdlGrammarParserSTRUCT_TYPE_KEYWORD, BondIdlGrammarParserENUM_TYPE_KEYWORD, BondIdlGrammarParserBLOCK_BODY_BEGIN, BondIdlGrammarParserSTMT_END, BondIdlGrammarParserFIELD_TYPENAME_DELIMITER, BondIdlGrammarParserATTRIBUTE_BEGIN:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BondIdlGrammarParserFIELD_TYPENAME_DELIMITER {
			{
				p.SetState(159)
				p.StructBaseClassDef()
			}

		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BondIdlGrammarParserBLOCK_BODY_BEGIN {
			{
				p.SetState(162)
				p.StructBody()
			}

		}

	case BondIdlGrammarParserVIEW_OF_KEYWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Match(BondIdlGrammarParserVIEW_OF_KEYWORD)
		}
		{
			p.SetState(166)
			p.TypeNameWithGeneric()
		}
		{
			p.SetState(167)
			p.ViewOfBody()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IViewOfBodyContext is an interface to support dynamic dispatch.
type IViewOfBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsViewOfBodyContext differentiates from other interfaces.
	IsViewOfBodyContext()
}

type ViewOfBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyViewOfBodyContext() *ViewOfBodyContext {
	var p = new(ViewOfBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_viewOfBody
	return p
}

func (*ViewOfBodyContext) IsViewOfBodyContext() {}

func NewViewOfBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ViewOfBodyContext {
	var p = new(ViewOfBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_viewOfBody

	return p
}

func (s *ViewOfBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ViewOfBodyContext) BLOCK_BODY_BEGIN() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserBLOCK_BODY_BEGIN, 0)
}

func (s *ViewOfBodyContext) SingleViewOfField() ISingleViewOfFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleViewOfFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleViewOfFieldContext)
}

func (s *ViewOfBodyContext) BLOCK_BODY_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserBLOCK_BODY_END, 0)
}

func (s *ViewOfBodyContext) STMT_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserSTMT_END, 0)
}

func (s *ViewOfBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ViewOfBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ViewOfBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterViewOfBody(s)
	}
}

func (s *ViewOfBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitViewOfBody(s)
	}
}

func (p *BondIdlGrammarParser) ViewOfBody() (localctx IViewOfBodyContext) {
	localctx = NewViewOfBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BondIdlGrammarParserRULE_viewOfBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(BondIdlGrammarParserBLOCK_BODY_BEGIN)
	}
	{
		p.SetState(172)
		p.SingleViewOfField()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BondIdlGrammarParserSTMT_END {
		{
			p.SetState(173)
			p.Match(BondIdlGrammarParserSTMT_END)
		}

	}
	{
		p.SetState(176)
		p.Match(BondIdlGrammarParserBLOCK_BODY_END)
	}

	return localctx
}

// ISingleViewOfFieldContext is an interface to support dynamic dispatch.
type ISingleViewOfFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleViewOfFieldContext differentiates from other interfaces.
	IsSingleViewOfFieldContext()
}

type SingleViewOfFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleViewOfFieldContext() *SingleViewOfFieldContext {
	var p = new(SingleViewOfFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_singleViewOfField
	return p
}

func (*SingleViewOfFieldContext) IsSingleViewOfFieldContext() {}

func NewSingleViewOfFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleViewOfFieldContext {
	var p = new(SingleViewOfFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_singleViewOfField

	return p
}

func (s *SingleViewOfFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleViewOfFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserIDENTIFIER, 0)
}

func (s *SingleViewOfFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleViewOfFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleViewOfFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterSingleViewOfField(s)
	}
}

func (s *SingleViewOfFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitSingleViewOfField(s)
	}
}

func (p *BondIdlGrammarParser) SingleViewOfField() (localctx ISingleViewOfFieldContext) {
	localctx = NewSingleViewOfFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BondIdlGrammarParserRULE_singleViewOfField)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(BondIdlGrammarParserIDENTIFIER)
	}

	return localctx
}

// IStructBaseClassDefContext is an interface to support dynamic dispatch.
type IStructBaseClassDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructBaseClassDefContext differentiates from other interfaces.
	IsStructBaseClassDefContext()
}

type StructBaseClassDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructBaseClassDefContext() *StructBaseClassDefContext {
	var p = new(StructBaseClassDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_structBaseClassDef
	return p
}

func (*StructBaseClassDefContext) IsStructBaseClassDefContext() {}

func NewStructBaseClassDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBaseClassDefContext {
	var p = new(StructBaseClassDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_structBaseClassDef

	return p
}

func (s *StructBaseClassDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBaseClassDefContext) FIELD_TYPENAME_DELIMITER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserFIELD_TYPENAME_DELIMITER, 0)
}

func (s *StructBaseClassDefContext) TypeNameWithGeneric() ITypeNameWithGenericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameWithGenericContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameWithGenericContext)
}

func (s *StructBaseClassDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBaseClassDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBaseClassDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterStructBaseClassDef(s)
	}
}

func (s *StructBaseClassDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitStructBaseClassDef(s)
	}
}

func (p *BondIdlGrammarParser) StructBaseClassDef() (localctx IStructBaseClassDefContext) {
	localctx = NewStructBaseClassDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BondIdlGrammarParserRULE_structBaseClassDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(BondIdlGrammarParserFIELD_TYPENAME_DELIMITER)
	}
	{
		p.SetState(181)
		p.TypeNameWithGeneric()
	}

	return localctx
}

// IStructBodyContext is an interface to support dynamic dispatch.
type IStructBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructBodyContext differentiates from other interfaces.
	IsStructBodyContext()
}

type StructBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructBodyContext() *StructBodyContext {
	var p = new(StructBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_structBody
	return p
}

func (*StructBodyContext) IsStructBodyContext() {}

func NewStructBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBodyContext {
	var p = new(StructBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_structBody

	return p
}

func (s *StructBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBodyContext) BLOCK_BODY_BEGIN() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserBLOCK_BODY_BEGIN, 0)
}

func (s *StructBodyContext) StructFieldDefList() IStructFieldDefListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructFieldDefListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructFieldDefListContext)
}

func (s *StructBodyContext) BLOCK_BODY_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserBLOCK_BODY_END, 0)
}

func (s *StructBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterStructBody(s)
	}
}

func (s *StructBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitStructBody(s)
	}
}

func (p *BondIdlGrammarParser) StructBody() (localctx IStructBodyContext) {
	localctx = NewStructBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BondIdlGrammarParserRULE_structBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(BondIdlGrammarParserBLOCK_BODY_BEGIN)
	}
	{
		p.SetState(184)
		p.StructFieldDefList()
	}
	{
		p.SetState(185)
		p.Match(BondIdlGrammarParserBLOCK_BODY_END)
	}

	return localctx
}

// IStructFieldDefListContext is an interface to support dynamic dispatch.
type IStructFieldDefListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructFieldDefListContext differentiates from other interfaces.
	IsStructFieldDefListContext()
}

type StructFieldDefListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructFieldDefListContext() *StructFieldDefListContext {
	var p = new(StructFieldDefListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_structFieldDefList
	return p
}

func (*StructFieldDefListContext) IsStructFieldDefListContext() {}

func NewStructFieldDefListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldDefListContext {
	var p = new(StructFieldDefListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_structFieldDefList

	return p
}

func (s *StructFieldDefListContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldDefListContext) AllSingleStructFieldDef() []ISingleStructFieldDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleStructFieldDefContext)(nil)).Elem())
	var tst = make([]ISingleStructFieldDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleStructFieldDefContext)
		}
	}

	return tst
}

func (s *StructFieldDefListContext) SingleStructFieldDef(i int) ISingleStructFieldDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleStructFieldDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleStructFieldDefContext)
}

func (s *StructFieldDefListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldDefListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructFieldDefListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterStructFieldDefList(s)
	}
}

func (s *StructFieldDefListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitStructFieldDefList(s)
	}
}

func (p *BondIdlGrammarParser) StructFieldDefList() (localctx IStructFieldDefListContext) {
	localctx = NewStructFieldDefListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BondIdlGrammarParserRULE_structFieldDefList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(BondIdlGrammarParserDEC_NUMBER-22))|(1<<(BondIdlGrammarParserHEX_NUMBER-22))|(1<<(BondIdlGrammarParserATTRIBUTE_BEGIN-22)))) != 0 {
		{
			p.SetState(187)
			p.SingleStructFieldDef()
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISingleStructFieldDefContext is an interface to support dynamic dispatch.
type ISingleStructFieldDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleStructFieldDefContext differentiates from other interfaces.
	IsSingleStructFieldDefContext()
}

type SingleStructFieldDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleStructFieldDefContext() *SingleStructFieldDefContext {
	var p = new(SingleStructFieldDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_singleStructFieldDef
	return p
}

func (*SingleStructFieldDefContext) IsSingleStructFieldDefContext() {}

func NewSingleStructFieldDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleStructFieldDefContext {
	var p = new(SingleStructFieldDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_singleStructFieldDef

	return p
}

func (s *SingleStructFieldDefContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleStructFieldDefContext) AttributeDefList() IAttributeDefListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeDefListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeDefListContext)
}

func (s *SingleStructFieldDefContext) FieldID() IFieldIDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldIDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldIDContext)
}

func (s *SingleStructFieldDefContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *SingleStructFieldDefContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *SingleStructFieldDefContext) FieldModifier() IFieldModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldModifierContext)
}

func (s *SingleStructFieldDefContext) DefaultValueSpec() IDefaultValueSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValueSpecContext)
}

func (s *SingleStructFieldDefContext) STMT_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserSTMT_END, 0)
}

func (s *SingleStructFieldDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleStructFieldDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleStructFieldDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterSingleStructFieldDef(s)
	}
}

func (s *SingleStructFieldDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitSingleStructFieldDef(s)
	}
}

func (p *BondIdlGrammarParser) SingleStructFieldDef() (localctx ISingleStructFieldDefContext) {
	localctx = NewSingleStructFieldDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BondIdlGrammarParserRULE_singleStructFieldDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.AttributeDefList()
	}
	{
		p.SetState(194)
		p.FieldID()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BondIdlGrammarParserREQUIRED_KEYWORD)|(1<<BondIdlGrammarParserOPTIONAL_KEYWORD)|(1<<BondIdlGrammarParserREQUIRED_OPTIONAL_KEYWORD))) != 0 {
		{
			p.SetState(195)
			p.FieldModifier()
		}

	}
	{
		p.SetState(198)
		p.FieldType()
	}
	{
		p.SetState(199)
		p.FieldName()
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BondIdlGrammarParserT__0 {
		{
			p.SetState(200)
			p.DefaultValueSpec()
		}

	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BondIdlGrammarParserSTMT_END {
		{
			p.SetState(203)
			p.Match(BondIdlGrammarParserSTMT_END)
		}

	}

	return localctx
}

// IFieldIDContext is an interface to support dynamic dispatch.
type IFieldIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldIDContext differentiates from other interfaces.
	IsFieldIDContext()
}

type FieldIDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldIDContext() *FieldIDContext {
	var p = new(FieldIDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_fieldID
	return p
}

func (*FieldIDContext) IsFieldIDContext() {}

func NewFieldIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldIDContext {
	var p = new(FieldIDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_fieldID

	return p
}

func (s *FieldIDContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldIDContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *FieldIDContext) FIELD_TYPENAME_DELIMITER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserFIELD_TYPENAME_DELIMITER, 0)
}

func (s *FieldIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterFieldID(s)
	}
}

func (s *FieldIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitFieldID(s)
	}
}

func (p *BondIdlGrammarParser) FieldID() (localctx IFieldIDContext) {
	localctx = NewFieldIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BondIdlGrammarParserRULE_fieldID)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.IntegerLiteral()
	}
	{
		p.SetState(207)
		p.Match(BondIdlGrammarParserFIELD_TYPENAME_DELIMITER)
	}

	return localctx
}

// IFieldModifierContext is an interface to support dynamic dispatch.
type IFieldModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldModifierContext differentiates from other interfaces.
	IsFieldModifierContext()
}

type FieldModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldModifierContext() *FieldModifierContext {
	var p = new(FieldModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_fieldModifier
	return p
}

func (*FieldModifierContext) IsFieldModifierContext() {}

func NewFieldModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldModifierContext {
	var p = new(FieldModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_fieldModifier

	return p
}

func (s *FieldModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldModifierContext) REQUIRED_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserREQUIRED_KEYWORD, 0)
}

func (s *FieldModifierContext) OPTIONAL_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserOPTIONAL_KEYWORD, 0)
}

func (s *FieldModifierContext) REQUIRED_OPTIONAL_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserREQUIRED_OPTIONAL_KEYWORD, 0)
}

func (s *FieldModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterFieldModifier(s)
	}
}

func (s *FieldModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitFieldModifier(s)
	}
}

func (p *BondIdlGrammarParser) FieldModifier() (localctx IFieldModifierContext) {
	localctx = NewFieldModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BondIdlGrammarParserRULE_fieldModifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BondIdlGrammarParserREQUIRED_KEYWORD)|(1<<BondIdlGrammarParserOPTIONAL_KEYWORD)|(1<<BondIdlGrammarParserREQUIRED_OPTIONAL_KEYWORD))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_fieldType
	return p
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) TypeNameWithGeneric() ITypeNameWithGenericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameWithGenericContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameWithGenericContext)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *BondIdlGrammarParser) FieldType() (localctx IFieldTypeContext) {
	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BondIdlGrammarParserRULE_fieldType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.TypeNameWithGeneric()
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) NOTHING_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserNOTHING_KEYWORD, 0)
}

func (s *FieldNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserIDENTIFIER, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *BondIdlGrammarParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BondIdlGrammarParserRULE_fieldName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BondIdlGrammarParserNOTHING_KEYWORD || _la == BondIdlGrammarParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDefaultValueSpecContext is an interface to support dynamic dispatch.
type IDefaultValueSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultValueSpecContext differentiates from other interfaces.
	IsDefaultValueSpecContext()
}

type DefaultValueSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultValueSpecContext() *DefaultValueSpecContext {
	var p = new(DefaultValueSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_defaultValueSpec
	return p
}

func (*DefaultValueSpecContext) IsDefaultValueSpecContext() {}

func NewDefaultValueSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultValueSpecContext {
	var p = new(DefaultValueSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_defaultValueSpec

	return p
}

func (s *DefaultValueSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultValueSpecContext) DefaultValues() IDefaultValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValuesContext)
}

func (s *DefaultValueSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultValueSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultValueSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterDefaultValueSpec(s)
	}
}

func (s *DefaultValueSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitDefaultValueSpec(s)
	}
}

func (p *BondIdlGrammarParser) DefaultValueSpec() (localctx IDefaultValueSpecContext) {
	localctx = NewDefaultValueSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BondIdlGrammarParserRULE_defaultValueSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(BondIdlGrammarParserT__0)
	}
	{
		p.SetState(216)
		p.DefaultValues()
	}

	return localctx
}

// IDefaultValuesContext is an interface to support dynamic dispatch.
type IDefaultValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultValuesContext differentiates from other interfaces.
	IsDefaultValuesContext()
}

type DefaultValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultValuesContext() *DefaultValuesContext {
	var p = new(DefaultValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_defaultValues
	return p
}

func (*DefaultValuesContext) IsDefaultValuesContext() {}

func NewDefaultValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultValuesContext {
	var p = new(DefaultValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_defaultValues

	return p
}

func (s *DefaultValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultValuesContext) HEX_NUMBER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserHEX_NUMBER, 0)
}

func (s *DefaultValuesContext) DEC_NUMBER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserDEC_NUMBER, 0)
}

func (s *DefaultValuesContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserFLOAT_NUMBER, 0)
}

func (s *DefaultValuesContext) NOTHING_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserNOTHING_KEYWORD, 0)
}

func (s *DefaultValuesContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserIDENTIFIER, 0)
}

func (s *DefaultValuesContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserQUOTED_STRING, 0)
}

func (s *DefaultValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterDefaultValues(s)
	}
}

func (s *DefaultValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitDefaultValues(s)
	}
}

func (p *BondIdlGrammarParser) DefaultValues() (localctx IDefaultValuesContext) {
	localctx = NewDefaultValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BondIdlGrammarParserRULE_defaultValues)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BondIdlGrammarParserNOTHING_KEYWORD)|(1<<BondIdlGrammarParserDEC_NUMBER)|(1<<BondIdlGrammarParserFLOAT_NUMBER)|(1<<BondIdlGrammarParserHEX_NUMBER)|(1<<BondIdlGrammarParserIDENTIFIER)|(1<<BondIdlGrammarParserQUOTED_STRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeNameWithGenericContext is an interface to support dynamic dispatch.
type ITypeNameWithGenericContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameWithGenericContext differentiates from other interfaces.
	IsTypeNameWithGenericContext()
}

type TypeNameWithGenericContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameWithGenericContext() *TypeNameWithGenericContext {
	var p = new(TypeNameWithGenericContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_typeNameWithGeneric
	return p
}

func (*TypeNameWithGenericContext) IsTypeNameWithGenericContext() {}

func NewTypeNameWithGenericContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameWithGenericContext {
	var p = new(TypeNameWithGenericContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_typeNameWithGeneric

	return p
}

func (s *TypeNameWithGenericContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameWithGenericContext) BuiltinPrimitiveType() IBuiltinPrimitiveTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltinPrimitiveTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltinPrimitiveTypeContext)
}

func (s *TypeNameWithGenericContext) BuiltinContainerType() IBuiltinContainerTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltinContainerTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltinContainerTypeContext)
}

func (s *TypeNameWithGenericContext) GenericTypeArgs() IGenericTypeArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGenericTypeArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGenericTypeArgsContext)
}

func (s *TypeNameWithGenericContext) CustomType() ICustomTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICustomTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICustomTypeContext)
}

func (s *TypeNameWithGenericContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameWithGenericContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameWithGenericContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterTypeNameWithGeneric(s)
	}
}

func (s *TypeNameWithGenericContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitTypeNameWithGeneric(s)
	}
}

func (p *BondIdlGrammarParser) TypeNameWithGeneric() (localctx ITypeNameWithGenericContext) {
	localctx = NewTypeNameWithGenericContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BondIdlGrammarParserRULE_typeNameWithGeneric)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(228)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BondIdlGrammarParserSIGNED_INTEGER_TYPE_KEYWORD, BondIdlGrammarParserUNSIGNED_INTEGER_TYPE_KEYWORD, BondIdlGrammarParserFLOAT_POINT_TYPE_KEYWORD, BondIdlGrammarParserBYTESTRING_TYPE_KEYWORD, BondIdlGrammarParserWSTRING_TYPE_KEYWORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.BuiltinPrimitiveType()
		}

	case BondIdlGrammarParserCONTAINER_TYPE_KEYWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(221)
			p.BuiltinContainerType()
		}
		{
			p.SetState(222)
			p.GenericTypeArgs()
		}

	case BondIdlGrammarParserMULTI_SECTION_IDENTIFIER, BondIdlGrammarParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(224)
			p.CustomType()
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BondIdlGrammarParserGENERIC_TYPELIST_BEGIN {
			{
				p.SetState(225)
				p.GenericTypeArgs()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBuiltinPrimitiveTypeContext is an interface to support dynamic dispatch.
type IBuiltinPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltinPrimitiveTypeContext differentiates from other interfaces.
	IsBuiltinPrimitiveTypeContext()
}

type BuiltinPrimitiveTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltinPrimitiveTypeContext() *BuiltinPrimitiveTypeContext {
	var p = new(BuiltinPrimitiveTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_builtinPrimitiveType
	return p
}

func (*BuiltinPrimitiveTypeContext) IsBuiltinPrimitiveTypeContext() {}

func NewBuiltinPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinPrimitiveTypeContext {
	var p = new(BuiltinPrimitiveTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_builtinPrimitiveType

	return p
}

func (s *BuiltinPrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltinPrimitiveTypeContext) SIGNED_INTEGER_TYPE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserSIGNED_INTEGER_TYPE_KEYWORD, 0)
}

func (s *BuiltinPrimitiveTypeContext) UNSIGNED_INTEGER_TYPE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserUNSIGNED_INTEGER_TYPE_KEYWORD, 0)
}

func (s *BuiltinPrimitiveTypeContext) FLOAT_POINT_TYPE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserFLOAT_POINT_TYPE_KEYWORD, 0)
}

func (s *BuiltinPrimitiveTypeContext) BYTESTRING_TYPE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserBYTESTRING_TYPE_KEYWORD, 0)
}

func (s *BuiltinPrimitiveTypeContext) WSTRING_TYPE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserWSTRING_TYPE_KEYWORD, 0)
}

func (s *BuiltinPrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinPrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltinPrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterBuiltinPrimitiveType(s)
	}
}

func (s *BuiltinPrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitBuiltinPrimitiveType(s)
	}
}

func (p *BondIdlGrammarParser) BuiltinPrimitiveType() (localctx IBuiltinPrimitiveTypeContext) {
	localctx = NewBuiltinPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BondIdlGrammarParserRULE_builtinPrimitiveType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BondIdlGrammarParserSIGNED_INTEGER_TYPE_KEYWORD)|(1<<BondIdlGrammarParserUNSIGNED_INTEGER_TYPE_KEYWORD)|(1<<BondIdlGrammarParserFLOAT_POINT_TYPE_KEYWORD)|(1<<BondIdlGrammarParserBYTESTRING_TYPE_KEYWORD)|(1<<BondIdlGrammarParserWSTRING_TYPE_KEYWORD))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBuiltinContainerTypeContext is an interface to support dynamic dispatch.
type IBuiltinContainerTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltinContainerTypeContext differentiates from other interfaces.
	IsBuiltinContainerTypeContext()
}

type BuiltinContainerTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltinContainerTypeContext() *BuiltinContainerTypeContext {
	var p = new(BuiltinContainerTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_builtinContainerType
	return p
}

func (*BuiltinContainerTypeContext) IsBuiltinContainerTypeContext() {}

func NewBuiltinContainerTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinContainerTypeContext {
	var p = new(BuiltinContainerTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_builtinContainerType

	return p
}

func (s *BuiltinContainerTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltinContainerTypeContext) CONTAINER_TYPE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserCONTAINER_TYPE_KEYWORD, 0)
}

func (s *BuiltinContainerTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinContainerTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltinContainerTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterBuiltinContainerType(s)
	}
}

func (s *BuiltinContainerTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitBuiltinContainerType(s)
	}
}

func (p *BondIdlGrammarParser) BuiltinContainerType() (localctx IBuiltinContainerTypeContext) {
	localctx = NewBuiltinContainerTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BondIdlGrammarParserRULE_builtinContainerType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(BondIdlGrammarParserCONTAINER_TYPE_KEYWORD)
	}

	return localctx
}

// ICustomTypeContext is an interface to support dynamic dispatch.
type ICustomTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCustomTypeContext differentiates from other interfaces.
	IsCustomTypeContext()
}

type CustomTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCustomTypeContext() *CustomTypeContext {
	var p = new(CustomTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_customType
	return p
}

func (*CustomTypeContext) IsCustomTypeContext() {}

func NewCustomTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CustomTypeContext {
	var p = new(CustomTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_customType

	return p
}

func (s *CustomTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CustomTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserIDENTIFIER, 0)
}

func (s *CustomTypeContext) MULTI_SECTION_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserMULTI_SECTION_IDENTIFIER, 0)
}

func (s *CustomTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CustomTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CustomTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterCustomType(s)
	}
}

func (s *CustomTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitCustomType(s)
	}
}

func (p *BondIdlGrammarParser) CustomType() (localctx ICustomTypeContext) {
	localctx = NewCustomTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BondIdlGrammarParserRULE_customType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BondIdlGrammarParserMULTI_SECTION_IDENTIFIER || _la == BondIdlGrammarParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IGenericTypeArgsContext is an interface to support dynamic dispatch.
type IGenericTypeArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenericTypeArgsContext differentiates from other interfaces.
	IsGenericTypeArgsContext()
}

type GenericTypeArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericTypeArgsContext() *GenericTypeArgsContext {
	var p = new(GenericTypeArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_genericTypeArgs
	return p
}

func (*GenericTypeArgsContext) IsGenericTypeArgsContext() {}

func NewGenericTypeArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericTypeArgsContext {
	var p = new(GenericTypeArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_genericTypeArgs

	return p
}

func (s *GenericTypeArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericTypeArgsContext) GENERIC_TYPELIST_BEGIN() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserGENERIC_TYPELIST_BEGIN, 0)
}

func (s *GenericTypeArgsContext) TypeArgsList() ITypeArgsListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeArgsListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeArgsListContext)
}

func (s *GenericTypeArgsContext) GENERIC_TYPELIST_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserGENERIC_TYPELIST_END, 0)
}

func (s *GenericTypeArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericTypeArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericTypeArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterGenericTypeArgs(s)
	}
}

func (s *GenericTypeArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitGenericTypeArgs(s)
	}
}

func (p *BondIdlGrammarParser) GenericTypeArgs() (localctx IGenericTypeArgsContext) {
	localctx = NewGenericTypeArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BondIdlGrammarParserRULE_genericTypeArgs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(BondIdlGrammarParserGENERIC_TYPELIST_BEGIN)
	}
	{
		p.SetState(237)
		p.TypeArgsList()
	}
	{
		p.SetState(238)
		p.Match(BondIdlGrammarParserGENERIC_TYPELIST_END)
	}

	return localctx
}

// ITypeArgsListContext is an interface to support dynamic dispatch.
type ITypeArgsListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeArgsListContext differentiates from other interfaces.
	IsTypeArgsListContext()
}

type TypeArgsListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeArgsListContext() *TypeArgsListContext {
	var p = new(TypeArgsListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_typeArgsList
	return p
}

func (*TypeArgsListContext) IsTypeArgsListContext() {}

func NewTypeArgsListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeArgsListContext {
	var p = new(TypeArgsListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_typeArgsList

	return p
}

func (s *TypeArgsListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeArgsListContext) TypeArgName() ITypeArgNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeArgNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeArgNameContext)
}

func (s *TypeArgsListContext) AllMoreTypeArgsList() []IMoreTypeArgsListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMoreTypeArgsListContext)(nil)).Elem())
	var tst = make([]IMoreTypeArgsListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMoreTypeArgsListContext)
		}
	}

	return tst
}

func (s *TypeArgsListContext) MoreTypeArgsList(i int) IMoreTypeArgsListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMoreTypeArgsListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMoreTypeArgsListContext)
}

func (s *TypeArgsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeArgsListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeArgsListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterTypeArgsList(s)
	}
}

func (s *TypeArgsListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitTypeArgsList(s)
	}
}

func (p *BondIdlGrammarParser) TypeArgsList() (localctx ITypeArgsListContext) {
	localctx = NewTypeArgsListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BondIdlGrammarParserRULE_typeArgsList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.TypeArgName()
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BondIdlGrammarParserDEFINITION_DELIMITER {
		{
			p.SetState(241)
			p.MoreTypeArgsList()
		}

		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMoreTypeArgsListContext is an interface to support dynamic dispatch.
type IMoreTypeArgsListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoreTypeArgsListContext differentiates from other interfaces.
	IsMoreTypeArgsListContext()
}

type MoreTypeArgsListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoreTypeArgsListContext() *MoreTypeArgsListContext {
	var p = new(MoreTypeArgsListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_moreTypeArgsList
	return p
}

func (*MoreTypeArgsListContext) IsMoreTypeArgsListContext() {}

func NewMoreTypeArgsListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoreTypeArgsListContext {
	var p = new(MoreTypeArgsListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_moreTypeArgsList

	return p
}

func (s *MoreTypeArgsListContext) GetParser() antlr.Parser { return s.parser }

func (s *MoreTypeArgsListContext) DEFINITION_DELIMITER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserDEFINITION_DELIMITER, 0)
}

func (s *MoreTypeArgsListContext) TypeArgName() ITypeArgNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeArgNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeArgNameContext)
}

func (s *MoreTypeArgsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoreTypeArgsListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoreTypeArgsListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterMoreTypeArgsList(s)
	}
}

func (s *MoreTypeArgsListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitMoreTypeArgsList(s)
	}
}

func (p *BondIdlGrammarParser) MoreTypeArgsList() (localctx IMoreTypeArgsListContext) {
	localctx = NewMoreTypeArgsListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BondIdlGrammarParserRULE_moreTypeArgsList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(BondIdlGrammarParserDEFINITION_DELIMITER)
	}
	{
		p.SetState(248)
		p.TypeArgName()
	}

	return localctx
}

// ITypeArgNameContext is an interface to support dynamic dispatch.
type ITypeArgNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeArgNameContext differentiates from other interfaces.
	IsTypeArgNameContext()
}

type TypeArgNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeArgNameContext() *TypeArgNameContext {
	var p = new(TypeArgNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_typeArgName
	return p
}

func (*TypeArgNameContext) IsTypeArgNameContext() {}

func NewTypeArgNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeArgNameContext {
	var p = new(TypeArgNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_typeArgName

	return p
}

func (s *TypeArgNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeArgNameContext) TypeNameWithGeneric() ITypeNameWithGenericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameWithGenericContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameWithGenericContext)
}

func (s *TypeArgNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeArgNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeArgNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterTypeArgName(s)
	}
}

func (s *TypeArgNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitTypeArgName(s)
	}
}

func (p *BondIdlGrammarParser) TypeArgName() (localctx ITypeArgNameContext) {
	localctx = NewTypeArgNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BondIdlGrammarParserRULE_typeArgName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.TypeNameWithGeneric()
	}

	return localctx
}

// IStructNameDeclWithGenericContext is an interface to support dynamic dispatch.
type IStructNameDeclWithGenericContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructNameDeclWithGenericContext differentiates from other interfaces.
	IsStructNameDeclWithGenericContext()
}

type StructNameDeclWithGenericContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructNameDeclWithGenericContext() *StructNameDeclWithGenericContext {
	var p = new(StructNameDeclWithGenericContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_structNameDeclWithGeneric
	return p
}

func (*StructNameDeclWithGenericContext) IsStructNameDeclWithGenericContext() {}

func NewStructNameDeclWithGenericContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructNameDeclWithGenericContext {
	var p = new(StructNameDeclWithGenericContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_structNameDeclWithGeneric

	return p
}

func (s *StructNameDeclWithGenericContext) GetParser() antlr.Parser { return s.parser }

func (s *StructNameDeclWithGenericContext) TypeDeclName() ITypeDeclNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclNameContext)
}

func (s *StructNameDeclWithGenericContext) GenericTypeArgsDecl() IGenericTypeArgsDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGenericTypeArgsDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGenericTypeArgsDeclContext)
}

func (s *StructNameDeclWithGenericContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructNameDeclWithGenericContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructNameDeclWithGenericContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterStructNameDeclWithGeneric(s)
	}
}

func (s *StructNameDeclWithGenericContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitStructNameDeclWithGeneric(s)
	}
}

func (p *BondIdlGrammarParser) StructNameDeclWithGeneric() (localctx IStructNameDeclWithGenericContext) {
	localctx = NewStructNameDeclWithGenericContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BondIdlGrammarParserRULE_structNameDeclWithGeneric)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.TypeDeclName()
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BondIdlGrammarParserGENERIC_TYPELIST_BEGIN {
		{
			p.SetState(253)
			p.GenericTypeArgsDecl()
		}

	}

	return localctx
}

// ITypeDeclNameContext is an interface to support dynamic dispatch.
type ITypeDeclNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclNameContext differentiates from other interfaces.
	IsTypeDeclNameContext()
}

type TypeDeclNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclNameContext() *TypeDeclNameContext {
	var p = new(TypeDeclNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_typeDeclName
	return p
}

func (*TypeDeclNameContext) IsTypeDeclNameContext() {}

func NewTypeDeclNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclNameContext {
	var p = new(TypeDeclNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_typeDeclName

	return p
}

func (s *TypeDeclNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserIDENTIFIER, 0)
}

func (s *TypeDeclNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterTypeDeclName(s)
	}
}

func (s *TypeDeclNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitTypeDeclName(s)
	}
}

func (p *BondIdlGrammarParser) TypeDeclName() (localctx ITypeDeclNameContext) {
	localctx = NewTypeDeclNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BondIdlGrammarParserRULE_typeDeclName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(BondIdlGrammarParserIDENTIFIER)
	}

	return localctx
}

// IGenericTypeArgsDeclContext is an interface to support dynamic dispatch.
type IGenericTypeArgsDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenericTypeArgsDeclContext differentiates from other interfaces.
	IsGenericTypeArgsDeclContext()
}

type GenericTypeArgsDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericTypeArgsDeclContext() *GenericTypeArgsDeclContext {
	var p = new(GenericTypeArgsDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_genericTypeArgsDecl
	return p
}

func (*GenericTypeArgsDeclContext) IsGenericTypeArgsDeclContext() {}

func NewGenericTypeArgsDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericTypeArgsDeclContext {
	var p = new(GenericTypeArgsDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_genericTypeArgsDecl

	return p
}

func (s *GenericTypeArgsDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericTypeArgsDeclContext) GENERIC_TYPELIST_BEGIN() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserGENERIC_TYPELIST_BEGIN, 0)
}

func (s *GenericTypeArgsDeclContext) TypeParamList() ITypeParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParamListContext)
}

func (s *GenericTypeArgsDeclContext) GENERIC_TYPELIST_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserGENERIC_TYPELIST_END, 0)
}

func (s *GenericTypeArgsDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericTypeArgsDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericTypeArgsDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterGenericTypeArgsDecl(s)
	}
}

func (s *GenericTypeArgsDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitGenericTypeArgsDecl(s)
	}
}

func (p *BondIdlGrammarParser) GenericTypeArgsDecl() (localctx IGenericTypeArgsDeclContext) {
	localctx = NewGenericTypeArgsDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BondIdlGrammarParserRULE_genericTypeArgsDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(BondIdlGrammarParserGENERIC_TYPELIST_BEGIN)
	}
	{
		p.SetState(259)
		p.TypeParamList()
	}
	{
		p.SetState(260)
		p.Match(BondIdlGrammarParserGENERIC_TYPELIST_END)
	}

	return localctx
}

// ITypeParamListContext is an interface to support dynamic dispatch.
type ITypeParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeParamListContext differentiates from other interfaces.
	IsTypeParamListContext()
}

type TypeParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeParamListContext() *TypeParamListContext {
	var p = new(TypeParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_typeParamList
	return p
}

func (*TypeParamListContext) IsTypeParamListContext() {}

func NewTypeParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParamListContext {
	var p = new(TypeParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_typeParamList

	return p
}

func (s *TypeParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeParamListContext) TypeParamDef() ITypeParamDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParamDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParamDefContext)
}

func (s *TypeParamListContext) AllMoreTypeParamList() []IMoreTypeParamListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMoreTypeParamListContext)(nil)).Elem())
	var tst = make([]IMoreTypeParamListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMoreTypeParamListContext)
		}
	}

	return tst
}

func (s *TypeParamListContext) MoreTypeParamList(i int) IMoreTypeParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMoreTypeParamListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMoreTypeParamListContext)
}

func (s *TypeParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterTypeParamList(s)
	}
}

func (s *TypeParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitTypeParamList(s)
	}
}

func (p *BondIdlGrammarParser) TypeParamList() (localctx ITypeParamListContext) {
	localctx = NewTypeParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BondIdlGrammarParserRULE_typeParamList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.TypeParamDef()
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BondIdlGrammarParserDEFINITION_DELIMITER {
		{
			p.SetState(263)
			p.MoreTypeParamList()
		}

		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMoreTypeParamListContext is an interface to support dynamic dispatch.
type IMoreTypeParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoreTypeParamListContext differentiates from other interfaces.
	IsMoreTypeParamListContext()
}

type MoreTypeParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoreTypeParamListContext() *MoreTypeParamListContext {
	var p = new(MoreTypeParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_moreTypeParamList
	return p
}

func (*MoreTypeParamListContext) IsMoreTypeParamListContext() {}

func NewMoreTypeParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoreTypeParamListContext {
	var p = new(MoreTypeParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_moreTypeParamList

	return p
}

func (s *MoreTypeParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *MoreTypeParamListContext) DEFINITION_DELIMITER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserDEFINITION_DELIMITER, 0)
}

func (s *MoreTypeParamListContext) TypeParamDef() ITypeParamDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParamDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParamDefContext)
}

func (s *MoreTypeParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoreTypeParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoreTypeParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterMoreTypeParamList(s)
	}
}

func (s *MoreTypeParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitMoreTypeParamList(s)
	}
}

func (p *BondIdlGrammarParser) MoreTypeParamList() (localctx IMoreTypeParamListContext) {
	localctx = NewMoreTypeParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BondIdlGrammarParserRULE_moreTypeParamList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(BondIdlGrammarParserDEFINITION_DELIMITER)
	}
	{
		p.SetState(270)
		p.TypeParamDef()
	}

	return localctx
}

// ITypeParamDefContext is an interface to support dynamic dispatch.
type ITypeParamDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeParamDefContext differentiates from other interfaces.
	IsTypeParamDefContext()
}

type TypeParamDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeParamDefContext() *TypeParamDefContext {
	var p = new(TypeParamDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_typeParamDef
	return p
}

func (*TypeParamDefContext) IsTypeParamDefContext() {}

func NewTypeParamDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParamDefContext {
	var p = new(TypeParamDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_typeParamDef

	return p
}

func (s *TypeParamDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeParamDefContext) TypeParamName() ITypeParamNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParamNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParamNameContext)
}

func (s *TypeParamDefContext) TypeParamValueConstraint() ITypeParamValueConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParamValueConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParamValueConstraintContext)
}

func (s *TypeParamDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeParamDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterTypeParamDef(s)
	}
}

func (s *TypeParamDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitTypeParamDef(s)
	}
}

func (p *BondIdlGrammarParser) TypeParamDef() (localctx ITypeParamDefContext) {
	localctx = NewTypeParamDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, BondIdlGrammarParserRULE_typeParamDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.TypeParamName()
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BondIdlGrammarParserVALUE_CONSTRAINT_STMT {
		{
			p.SetState(273)
			p.TypeParamValueConstraint()
		}

	}

	return localctx
}

// ITypeParamNameContext is an interface to support dynamic dispatch.
type ITypeParamNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeParamNameContext differentiates from other interfaces.
	IsTypeParamNameContext()
}

type TypeParamNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeParamNameContext() *TypeParamNameContext {
	var p = new(TypeParamNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_typeParamName
	return p
}

func (*TypeParamNameContext) IsTypeParamNameContext() {}

func NewTypeParamNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParamNameContext {
	var p = new(TypeParamNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_typeParamName

	return p
}

func (s *TypeParamNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeParamNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserIDENTIFIER, 0)
}

func (s *TypeParamNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeParamNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterTypeParamName(s)
	}
}

func (s *TypeParamNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitTypeParamName(s)
	}
}

func (p *BondIdlGrammarParser) TypeParamName() (localctx ITypeParamNameContext) {
	localctx = NewTypeParamNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, BondIdlGrammarParserRULE_typeParamName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(BondIdlGrammarParserIDENTIFIER)
	}

	return localctx
}

// ITypeParamValueConstraintContext is an interface to support dynamic dispatch.
type ITypeParamValueConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeParamValueConstraintContext differentiates from other interfaces.
	IsTypeParamValueConstraintContext()
}

type TypeParamValueConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeParamValueConstraintContext() *TypeParamValueConstraintContext {
	var p = new(TypeParamValueConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_typeParamValueConstraint
	return p
}

func (*TypeParamValueConstraintContext) IsTypeParamValueConstraintContext() {}

func NewTypeParamValueConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParamValueConstraintContext {
	var p = new(TypeParamValueConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_typeParamValueConstraint

	return p
}

func (s *TypeParamValueConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeParamValueConstraintContext) VALUE_CONSTRAINT_STMT() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserVALUE_CONSTRAINT_STMT, 0)
}

func (s *TypeParamValueConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamValueConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeParamValueConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterTypeParamValueConstraint(s)
	}
}

func (s *TypeParamValueConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitTypeParamValueConstraint(s)
	}
}

func (p *BondIdlGrammarParser) TypeParamValueConstraint() (localctx ITypeParamValueConstraintContext) {
	localctx = NewTypeParamValueConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, BondIdlGrammarParserRULE_typeParamValueConstraint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(BondIdlGrammarParserVALUE_CONSTRAINT_STMT)
	}

	return localctx
}

// IAttributeDefListContext is an interface to support dynamic dispatch.
type IAttributeDefListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeDefListContext differentiates from other interfaces.
	IsAttributeDefListContext()
}

type AttributeDefListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeDefListContext() *AttributeDefListContext {
	var p = new(AttributeDefListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_attributeDefList
	return p
}

func (*AttributeDefListContext) IsAttributeDefListContext() {}

func NewAttributeDefListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeDefListContext {
	var p = new(AttributeDefListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_attributeDefList

	return p
}

func (s *AttributeDefListContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeDefListContext) AllAttributeDef() []IAttributeDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeDefContext)(nil)).Elem())
	var tst = make([]IAttributeDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeDefContext)
		}
	}

	return tst
}

func (s *AttributeDefListContext) AttributeDef(i int) IAttributeDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeDefContext)
}

func (s *AttributeDefListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeDefListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeDefListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterAttributeDefList(s)
	}
}

func (s *AttributeDefListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitAttributeDefList(s)
	}
}

func (p *BondIdlGrammarParser) AttributeDefList() (localctx IAttributeDefListContext) {
	localctx = NewAttributeDefListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, BondIdlGrammarParserRULE_attributeDefList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BondIdlGrammarParserATTRIBUTE_BEGIN {
		{
			p.SetState(280)
			p.AttributeDef()
		}

		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAttributeDefContext is an interface to support dynamic dispatch.
type IAttributeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeDefContext differentiates from other interfaces.
	IsAttributeDefContext()
}

type AttributeDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeDefContext() *AttributeDefContext {
	var p = new(AttributeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_attributeDef
	return p
}

func (*AttributeDefContext) IsAttributeDefContext() {}

func NewAttributeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeDefContext {
	var p = new(AttributeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_attributeDef

	return p
}

func (s *AttributeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeDefContext) ATTRIBUTE_BEGIN() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserATTRIBUTE_BEGIN, 0)
}

func (s *AttributeDefContext) AttributeBody() IAttributeBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeBodyContext)
}

func (s *AttributeDefContext) ATTRIBUTE_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserATTRIBUTE_END, 0)
}

func (s *AttributeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterAttributeDef(s)
	}
}

func (s *AttributeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitAttributeDef(s)
	}
}

func (p *BondIdlGrammarParser) AttributeDef() (localctx IAttributeDefContext) {
	localctx = NewAttributeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, BondIdlGrammarParserRULE_attributeDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(BondIdlGrammarParserATTRIBUTE_BEGIN)
	}
	{
		p.SetState(287)
		p.AttributeBody()
	}
	{
		p.SetState(288)
		p.Match(BondIdlGrammarParserATTRIBUTE_END)
	}

	return localctx
}

// IAttributeBodyContext is an interface to support dynamic dispatch.
type IAttributeBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeBodyContext differentiates from other interfaces.
	IsAttributeBodyContext()
}

type AttributeBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeBodyContext() *AttributeBodyContext {
	var p = new(AttributeBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_attributeBody
	return p
}

func (*AttributeBodyContext) IsAttributeBodyContext() {}

func NewAttributeBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeBodyContext {
	var p = new(AttributeBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_attributeBody

	return p
}

func (s *AttributeBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeBodyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserIDENTIFIER, 0)
}

func (s *AttributeBodyContext) ATTRIBUTE_PARAM_BEGIN() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserATTRIBUTE_PARAM_BEGIN, 0)
}

func (s *AttributeBodyContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserQUOTED_STRING, 0)
}

func (s *AttributeBodyContext) ATTRIBUTE_PARAM_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserATTRIBUTE_PARAM_END, 0)
}

func (s *AttributeBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterAttributeBody(s)
	}
}

func (s *AttributeBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitAttributeBody(s)
	}
}

func (p *BondIdlGrammarParser) AttributeBody() (localctx IAttributeBodyContext) {
	localctx = NewAttributeBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, BondIdlGrammarParserRULE_attributeBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(BondIdlGrammarParserIDENTIFIER)
	}
	{
		p.SetState(291)
		p.Match(BondIdlGrammarParserATTRIBUTE_PARAM_BEGIN)
	}
	{
		p.SetState(292)
		p.Match(BondIdlGrammarParserQUOTED_STRING)
	}
	{
		p.SetState(293)
		p.Match(BondIdlGrammarParserATTRIBUTE_PARAM_END)
	}

	return localctx
}

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumDefContext differentiates from other interfaces.
	IsEnumDefContext()
}

type EnumDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefContext() *EnumDefContext {
	var p = new(EnumDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_enumDef
	return p
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) ENUM_TYPE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserENUM_TYPE_KEYWORD, 0)
}

func (s *EnumDefContext) EnumName() IEnumNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumDefContext) EnumBody() IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumDefContext) STMT_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserSTMT_END, 0)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterEnumDef(s)
	}
}

func (s *EnumDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitEnumDef(s)
	}
}

func (p *BondIdlGrammarParser) EnumDef() (localctx IEnumDefContext) {
	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, BondIdlGrammarParserRULE_enumDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Match(BondIdlGrammarParserENUM_TYPE_KEYWORD)
	}
	{
		p.SetState(296)
		p.EnumName()
	}
	{
		p.SetState(297)
		p.EnumBody()
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BondIdlGrammarParserSTMT_END {
		{
			p.SetState(298)
			p.Match(BondIdlGrammarParserSTMT_END)
		}

	}

	return localctx
}

// IEnumNameContext is an interface to support dynamic dispatch.
type IEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumNameContext differentiates from other interfaces.
	IsEnumNameContext()
}

type EnumNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumNameContext() *EnumNameContext {
	var p = new(EnumNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_enumName
	return p
}

func (*EnumNameContext) IsEnumNameContext() {}

func NewEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumNameContext {
	var p = new(EnumNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_enumName

	return p
}

func (s *EnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserIDENTIFIER, 0)
}

func (s *EnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterEnumName(s)
	}
}

func (s *EnumNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitEnumName(s)
	}
}

func (p *BondIdlGrammarParser) EnumName() (localctx IEnumNameContext) {
	localctx = NewEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, BondIdlGrammarParserRULE_enumName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(BondIdlGrammarParserIDENTIFIER)
	}

	return localctx
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) BLOCK_BODY_BEGIN() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserBLOCK_BODY_BEGIN, 0)
}

func (s *EnumBodyContext) EnumSymbolDefList() IEnumSymbolDefListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumSymbolDefListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumSymbolDefListContext)
}

func (s *EnumBodyContext) BLOCK_BODY_END() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserBLOCK_BODY_END, 0)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (p *BondIdlGrammarParser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, BondIdlGrammarParserRULE_enumBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(BondIdlGrammarParserBLOCK_BODY_BEGIN)
	}
	{
		p.SetState(304)
		p.EnumSymbolDefList()
	}
	{
		p.SetState(305)
		p.Match(BondIdlGrammarParserBLOCK_BODY_END)
	}

	return localctx
}

// IEnumSymbolDefListContext is an interface to support dynamic dispatch.
type IEnumSymbolDefListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumSymbolDefListContext differentiates from other interfaces.
	IsEnumSymbolDefListContext()
}

type EnumSymbolDefListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumSymbolDefListContext() *EnumSymbolDefListContext {
	var p = new(EnumSymbolDefListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_enumSymbolDefList
	return p
}

func (*EnumSymbolDefListContext) IsEnumSymbolDefListContext() {}

func NewEnumSymbolDefListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumSymbolDefListContext {
	var p = new(EnumSymbolDefListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_enumSymbolDefList

	return p
}

func (s *EnumSymbolDefListContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumSymbolDefListContext) SingleEnumSymbolDef() ISingleEnumSymbolDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleEnumSymbolDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleEnumSymbolDefContext)
}

func (s *EnumSymbolDefListContext) AllMoreEnumSymbolDef() []IMoreEnumSymbolDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMoreEnumSymbolDefContext)(nil)).Elem())
	var tst = make([]IMoreEnumSymbolDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMoreEnumSymbolDefContext)
		}
	}

	return tst
}

func (s *EnumSymbolDefListContext) MoreEnumSymbolDef(i int) IMoreEnumSymbolDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMoreEnumSymbolDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMoreEnumSymbolDefContext)
}

func (s *EnumSymbolDefListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumSymbolDefListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumSymbolDefListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterEnumSymbolDefList(s)
	}
}

func (s *EnumSymbolDefListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitEnumSymbolDefList(s)
	}
}

func (p *BondIdlGrammarParser) EnumSymbolDefList() (localctx IEnumSymbolDefListContext) {
	localctx = NewEnumSymbolDefListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, BondIdlGrammarParserRULE_enumSymbolDefList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.SingleEnumSymbolDef()
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(308)
				p.MoreEnumSymbolDef()
			}

		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IMoreEnumSymbolDefContext is an interface to support dynamic dispatch.
type IMoreEnumSymbolDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoreEnumSymbolDefContext differentiates from other interfaces.
	IsMoreEnumSymbolDefContext()
}

type MoreEnumSymbolDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoreEnumSymbolDefContext() *MoreEnumSymbolDefContext {
	var p = new(MoreEnumSymbolDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_moreEnumSymbolDef
	return p
}

func (*MoreEnumSymbolDefContext) IsMoreEnumSymbolDefContext() {}

func NewMoreEnumSymbolDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoreEnumSymbolDefContext {
	var p = new(MoreEnumSymbolDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_moreEnumSymbolDef

	return p
}

func (s *MoreEnumSymbolDefContext) GetParser() antlr.Parser { return s.parser }

func (s *MoreEnumSymbolDefContext) DEFINITION_DELIMITER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserDEFINITION_DELIMITER, 0)
}

func (s *MoreEnumSymbolDefContext) EnumSymbolDefList() IEnumSymbolDefListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumSymbolDefListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumSymbolDefListContext)
}

func (s *MoreEnumSymbolDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoreEnumSymbolDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoreEnumSymbolDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterMoreEnumSymbolDef(s)
	}
}

func (s *MoreEnumSymbolDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitMoreEnumSymbolDef(s)
	}
}

func (p *BondIdlGrammarParser) MoreEnumSymbolDef() (localctx IMoreEnumSymbolDefContext) {
	localctx = NewMoreEnumSymbolDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, BondIdlGrammarParserRULE_moreEnumSymbolDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(BondIdlGrammarParserDEFINITION_DELIMITER)
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BondIdlGrammarParserIDENTIFIER {
		{
			p.SetState(315)
			p.EnumSymbolDefList()
		}

	}

	return localctx
}

// ISingleEnumSymbolDefContext is an interface to support dynamic dispatch.
type ISingleEnumSymbolDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleEnumSymbolDefContext differentiates from other interfaces.
	IsSingleEnumSymbolDefContext()
}

type SingleEnumSymbolDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleEnumSymbolDefContext() *SingleEnumSymbolDefContext {
	var p = new(SingleEnumSymbolDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_singleEnumSymbolDef
	return p
}

func (*SingleEnumSymbolDefContext) IsSingleEnumSymbolDefContext() {}

func NewSingleEnumSymbolDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleEnumSymbolDefContext {
	var p = new(SingleEnumSymbolDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_singleEnumSymbolDef

	return p
}

func (s *SingleEnumSymbolDefContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleEnumSymbolDefContext) EnumSymbol() IEnumSymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumSymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumSymbolContext)
}

func (s *SingleEnumSymbolDefContext) EnumSymbolValueAssignment() IEnumSymbolValueAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumSymbolValueAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumSymbolValueAssignmentContext)
}

func (s *SingleEnumSymbolDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleEnumSymbolDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleEnumSymbolDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterSingleEnumSymbolDef(s)
	}
}

func (s *SingleEnumSymbolDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitSingleEnumSymbolDef(s)
	}
}

func (p *BondIdlGrammarParser) SingleEnumSymbolDef() (localctx ISingleEnumSymbolDefContext) {
	localctx = NewSingleEnumSymbolDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, BondIdlGrammarParserRULE_singleEnumSymbolDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.EnumSymbol()
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BondIdlGrammarParserT__0 {
		{
			p.SetState(319)
			p.EnumSymbolValueAssignment()
		}

	}

	return localctx
}

// IEnumSymbolContext is an interface to support dynamic dispatch.
type IEnumSymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumSymbolContext differentiates from other interfaces.
	IsEnumSymbolContext()
}

type EnumSymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumSymbolContext() *EnumSymbolContext {
	var p = new(EnumSymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_enumSymbol
	return p
}

func (*EnumSymbolContext) IsEnumSymbolContext() {}

func NewEnumSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumSymbolContext {
	var p = new(EnumSymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_enumSymbol

	return p
}

func (s *EnumSymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumSymbolContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserIDENTIFIER, 0)
}

func (s *EnumSymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumSymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumSymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterEnumSymbol(s)
	}
}

func (s *EnumSymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitEnumSymbol(s)
	}
}

func (p *BondIdlGrammarParser) EnumSymbol() (localctx IEnumSymbolContext) {
	localctx = NewEnumSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, BondIdlGrammarParserRULE_enumSymbol)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(BondIdlGrammarParserIDENTIFIER)
	}

	return localctx
}

// IEnumSymbolValueAssignmentContext is an interface to support dynamic dispatch.
type IEnumSymbolValueAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumSymbolValueAssignmentContext differentiates from other interfaces.
	IsEnumSymbolValueAssignmentContext()
}

type EnumSymbolValueAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumSymbolValueAssignmentContext() *EnumSymbolValueAssignmentContext {
	var p = new(EnumSymbolValueAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_enumSymbolValueAssignment
	return p
}

func (*EnumSymbolValueAssignmentContext) IsEnumSymbolValueAssignmentContext() {}

func NewEnumSymbolValueAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumSymbolValueAssignmentContext {
	var p = new(EnumSymbolValueAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_enumSymbolValueAssignment

	return p
}

func (s *EnumSymbolValueAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumSymbolValueAssignmentContext) EnumSymbolValue() IEnumSymbolValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumSymbolValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumSymbolValueContext)
}

func (s *EnumSymbolValueAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumSymbolValueAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumSymbolValueAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterEnumSymbolValueAssignment(s)
	}
}

func (s *EnumSymbolValueAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitEnumSymbolValueAssignment(s)
	}
}

func (p *BondIdlGrammarParser) EnumSymbolValueAssignment() (localctx IEnumSymbolValueAssignmentContext) {
	localctx = NewEnumSymbolValueAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, BondIdlGrammarParserRULE_enumSymbolValueAssignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.Match(BondIdlGrammarParserT__0)
	}
	{
		p.SetState(325)
		p.EnumSymbolValue()
	}

	return localctx
}

// IEnumSymbolValueContext is an interface to support dynamic dispatch.
type IEnumSymbolValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumSymbolValueContext differentiates from other interfaces.
	IsEnumSymbolValueContext()
}

type EnumSymbolValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumSymbolValueContext() *EnumSymbolValueContext {
	var p = new(EnumSymbolValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_enumSymbolValue
	return p
}

func (*EnumSymbolValueContext) IsEnumSymbolValueContext() {}

func NewEnumSymbolValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumSymbolValueContext {
	var p = new(EnumSymbolValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_enumSymbolValue

	return p
}

func (s *EnumSymbolValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumSymbolValueContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *EnumSymbolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumSymbolValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumSymbolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterEnumSymbolValue(s)
	}
}

func (s *EnumSymbolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitEnumSymbolValue(s)
	}
}

func (p *BondIdlGrammarParser) EnumSymbolValue() (localctx IEnumSymbolValueContext) {
	localctx = NewEnumSymbolValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, BondIdlGrammarParserRULE_enumSymbolValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.IntegerLiteral()
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BondIdlGrammarParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BondIdlGrammarParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) HEX_NUMBER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserHEX_NUMBER, 0)
}

func (s *IntegerLiteralContext) DEC_NUMBER() antlr.TerminalNode {
	return s.GetToken(BondIdlGrammarParserDEC_NUMBER, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BondIdlGrammarListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (p *BondIdlGrammarParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, BondIdlGrammarParserRULE_integerLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BondIdlGrammarParserDEC_NUMBER || _la == BondIdlGrammarParserHEX_NUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
