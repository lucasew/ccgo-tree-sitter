// Code generated for linux/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-html/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-html -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && amd64

package grammar_html

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 0
const BUFSIZ = 1024
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXTERNAL_TOKEN_COUNT = 9
const FIELD_COUNT = 0
const FILENAME_MAX = 4096
const FOPEN_MAX = 1000
const INT16_MAX = 0x7fff
const INT32_MAX = 0x7fffffff
const INT64_MAX = 0x7fffffffffffffff
const INT8_MAX = 0x7f
const INTMAX_MAX = "INT64_MAX"
const INTMAX_MIN = "INT64_MIN"
const INTPTR_MAX = "INT64_MAX"
const INTPTR_MIN = "INT64_MIN"
const INT_FAST16_MAX = "INT32_MAX"
const INT_FAST16_MIN = "INT32_MIN"
const INT_FAST32_MAX = "INT32_MAX"
const INT_FAST32_MIN = "INT32_MIN"
const INT_FAST64_MAX = "INT64_MAX"
const INT_FAST64_MIN = "INT64_MIN"
const INT_FAST8_MAX = "INT8_MAX"
const INT_FAST8_MIN = "INT8_MIN"
const INT_LEAST16_MAX = "INT16_MAX"
const INT_LEAST16_MIN = "INT16_MIN"
const INT_LEAST32_MAX = "INT32_MAX"
const INT_LEAST32_MIN = "INT32_MIN"
const INT_LEAST64_MAX = "INT64_MAX"
const INT_LEAST64_MIN = "INT64_MIN"
const INT_LEAST8_MAX = "INT8_MAX"
const INT_LEAST8_MIN = "INT8_MIN"
const LANGUAGE_VERSION = 14
const LARGE_STATE_COUNT = 2
const L_ctermid = 20
const L_cuserid = 20
const L_tmpnam = 20
const MAX_ALIAS_SEQUENCE_LENGTH = 4
const PRODUCTION_ID_COUNT = 1
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const P_tmpdir = "/tmp"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 94
const SYMBOL_COUNT = 41
const TMP_MAX = 10000
const TOKEN_COUNT = 25
const TREE_SITTER_SERIALIZATION_BUFFER_SIZE = 1024
const UINT16_MAX = 65535
const UINT32_MAX = "0xffffffffu"
const UINT64_MAX = "0xffffffffffffffffu"
const UINT8_MAX = 255
const UINTMAX_MAX = "UINT64_MAX"
const UINTPTR_MAX = "UINT64_MAX"
const UINT_FAST16_MAX = "UINT32_MAX"
const UINT_FAST32_MAX = "UINT32_MAX"
const UINT_FAST64_MAX = "UINT64_MAX"
const UINT_FAST8_MAX = "UINT8_MAX"
const UINT_LEAST16_MAX = "UINT16_MAX"
const UINT_LEAST32_MAX = "UINT32_MAX"
const UINT_LEAST64_MAX = "UINT64_MAX"
const UINT_LEAST8_MAX = "UINT8_MAX"
const WINT_MAX = "UINT32_MAX"
const WINT_MIN = 0
const WNOHANG = 1
const WUNTRACED = 2
const _GNU_SOURCE = 1
const _IOFBF = 0
const _IOLBF = 1
const _IONBF = 2
const _LP64 = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BIGGEST_ALIGNMENT__ = 16
const __BIG_ENDIAN = 4321
const __BITINT_MAXWIDTH__ = 8388608
const __BOOL_WIDTH__ = 8
const __BYTE_ORDER = 1234
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CHAR_BIT__ = 8
const __CLANG_ATOMIC_BOOL_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR_LOCK_FREE = 2
const __CLANG_ATOMIC_INT_LOCK_FREE = 2
const __CLANG_ATOMIC_LLONG_LOCK_FREE = 2
const __CLANG_ATOMIC_LONG_LOCK_FREE = 2
const __CLANG_ATOMIC_POINTER_LOCK_FREE = 2
const __CLANG_ATOMIC_SHORT_LOCK_FREE = 2
const __CLANG_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __CONSTANT_CFSTRINGS__ = 1
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DENORM_MIN__ = 4.9406564584124654e-324
const __DBL_DIG__ = 15
const __DBL_EPSILON__ = 2.2204460492503131e-16
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DBL_MAX__ = 1.7976931348623157e+308
const __DBL_MIN__ = 2.2250738585072014e-308
const __DECIMAL_DIG__ = "__LDBL_DECIMAL_DIG__"
const __ELF__ = 1
const __FINITE_MATH_ONLY__ = 0
const __FLOAT128__ = 1
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.9604644775390625e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.765625e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.5504e+4
const __FLT16_MIN__ = 6.103515625e-5
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209290e-7
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282347e+38
const __FLT_MIN__ = 1.17549435e-38
const __FLT_RADIX__ = 2
const __FPCLASS_NEGINF = 0x0004
const __FPCLASS_NEGNORMAL = 0x0008
const __FPCLASS_NEGSUBNORMAL = 0x0010
const __FPCLASS_NEGZERO = 0x0020
const __FPCLASS_POSINF = 0x0200
const __FPCLASS_POSNORMAL = 0x0100
const __FPCLASS_POSSUBNORMAL = 0x0080
const __FPCLASS_POSZERO = 0x0040
const __FPCLASS_QNAN = 0x0002
const __FPCLASS_SNAN = 0x0001
const __FUNCTION__ = "__func__"
const __FXSR__ = 1
const __GCC_ASM_FLAG_OUTPUTS__ = 1
const __GCC_ATOMIC_BOOL_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR_LOCK_FREE = 2
const __GCC_ATOMIC_INT_LOCK_FREE = 2
const __GCC_ATOMIC_LLONG_LOCK_FREE = 2
const __GCC_ATOMIC_LONG_LOCK_FREE = 2
const __GCC_ATOMIC_POINTER_LOCK_FREE = 2
const __GCC_ATOMIC_SHORT_LOCK_FREE = 2
const __GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const __GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GNUC_MINOR__ = 2
const __GNUC_PATCHLEVEL__ = 1
const __GNUC_STDC_INLINE__ = 1
const __GNUC__ = 4
const __GXX_ABI_VERSION = 1002
const __INT16_FMTd__ = "hd"
const __INT16_FMTi__ = "hi"
const __INT16_MAX__ = 32767
const __INT16_TYPE__ = "short"
const __INT32_FMTd__ = "d"
const __INT32_FMTi__ = "i"
const __INT32_MAX__ = 2147483647
const __INT32_TYPE__ = "int"
const __INT64_C_SUFFIX__ = "L"
const __INT64_FMTd__ = "ld"
const __INT64_FMTi__ = "li"
const __INT64_MAX__ = 9223372036854775807
const __INT8_FMTd__ = "hhd"
const __INT8_FMTi__ = "hhi"
const __INT8_MAX__ = 127
const __INTMAX_C_SUFFIX__ = "L"
const __INTMAX_FMTd__ = "ld"
const __INTMAX_FMTi__ = "li"
const __INTMAX_MAX__ = 9223372036854775807
const __INTMAX_WIDTH__ = 64
const __INTPTR_FMTd__ = "ld"
const __INTPTR_FMTi__ = "li"
const __INTPTR_MAX__ = 9223372036854775807
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_FMTd__ = "hd"
const __INT_FAST16_FMTi__ = "hi"
const __INT_FAST16_MAX__ = 32767
const __INT_FAST16_TYPE__ = "short"
const __INT_FAST16_WIDTH__ = 16
const __INT_FAST32_FMTd__ = "d"
const __INT_FAST32_FMTi__ = "i"
const __INT_FAST32_MAX__ = 2147483647
const __INT_FAST32_TYPE__ = "int"
const __INT_FAST32_WIDTH__ = 32
const __INT_FAST64_FMTd__ = "ld"
const __INT_FAST64_FMTi__ = "li"
const __INT_FAST64_MAX__ = 9223372036854775807
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_FMTd__ = "hhd"
const __INT_FAST8_FMTi__ = "hhi"
const __INT_FAST8_MAX__ = 127
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_FMTd__ = "hd"
const __INT_LEAST16_FMTi__ = "hi"
const __INT_LEAST16_MAX__ = 32767
const __INT_LEAST16_TYPE__ = "short"
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_FMTd__ = "d"
const __INT_LEAST32_FMTi__ = "i"
const __INT_LEAST32_MAX__ = 2147483647
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_FMTd__ = "ld"
const __INT_LEAST64_FMTi__ = "li"
const __INT_LEAST64_MAX__ = 9223372036854775807
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_FMTd__ = "hhd"
const __INT_LEAST8_FMTi__ = "hhi"
const __INT_LEAST8_MAX__ = 127
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 2147483647
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 4.9406564584124654e-324
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 2.2204460492503131e-16
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 1.7976931348623157e+308
const __LDBL_MIN__ = 2.2250738585072014e-308
const __LITTLE_ENDIAN = 1234
const __LITTLE_ENDIAN__ = 1
const __LLONG_WIDTH__ = 64
const __LONG_LONG_MAX__ = 9223372036854775807
const __LONG_MAX = 0x7fffffffffffffff
const __LONG_MAX__ = 9223372036854775807
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MEMORY_SCOPE_DEVICE = 1
const __MEMORY_SCOPE_SINGLE = 4
const __MEMORY_SCOPE_SYSTEM = 0
const __MEMORY_SCOPE_WRKGRP = 2
const __MEMORY_SCOPE_WVFRNT = 3
const __MMX__ = 1
const __NO_INLINE__ = 1
const __NO_MATH_INLINES = 1
const __OBJC_BOOL_IS_BOOL = 0
const __OPENCL_MEMORY_SCOPE_ALL_SVM_DEVICES = 3
const __OPENCL_MEMORY_SCOPE_DEVICE = 2
const __OPENCL_MEMORY_SCOPE_SUB_GROUP = 4
const __OPENCL_MEMORY_SCOPE_WORK_GROUP = 1
const __OPENCL_MEMORY_SCOPE_WORK_ITEM = 0
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PIE__ = 2
const __POINTER_WIDTH__ = 64
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_FMTd__ = "ld"
const __PTRDIFF_FMTi__ = "li"
const __PTRDIFF_MAX__ = 9223372036854775807
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 127
const __SEG_FS = 1
const __SEG_GS = 1
const __SHRT_MAX__ = 32767
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 2147483647
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT128__ = 16
const __SIZEOF_FLOAT__ = 4
const __SIZEOF_INT128__ = 16
const __SIZEOF_INT__ = 4
const __SIZEOF_LONG_DOUBLE__ = 8
const __SIZEOF_LONG_LONG__ = 8
const __SIZEOF_LONG__ = 8
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 4
const __SIZEOF_WINT_T__ = 4
const __SIZE_FMTX__ = "lX"
const __SIZE_FMTo__ = "lo"
const __SIZE_FMTu__ = "lu"
const __SIZE_FMTx__ = "lx"
const __SIZE_MAX__ = 18446744073709551615
const __SIZE_WIDTH__ = 64
const __SSE2_MATH__ = 1
const __SSE2__ = 1
const __SSE_MATH__ = 1
const __SSE__ = 1
const __STDC_HOSTED__ = 1
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201112
const __STDC__ = 1
const __UINT16_FMTX__ = "hX"
const __UINT16_FMTo__ = "ho"
const __UINT16_FMTu__ = "hu"
const __UINT16_FMTx__ = "hx"
const __UINT16_MAX__ = 65535
const __UINT32_C_SUFFIX__ = "U"
const __UINT32_FMTX__ = "X"
const __UINT32_FMTo__ = "o"
const __UINT32_FMTu__ = "u"
const __UINT32_FMTx__ = "x"
const __UINT32_MAX__ = 4294967295
const __UINT64_C_SUFFIX__ = "UL"
const __UINT64_FMTX__ = "lX"
const __UINT64_FMTo__ = "lo"
const __UINT64_FMTu__ = "lu"
const __UINT64_FMTx__ = "lx"
const __UINT64_MAX__ = 18446744073709551615
const __UINT8_FMTX__ = "hhX"
const __UINT8_FMTo__ = "hho"
const __UINT8_FMTu__ = "hhu"
const __UINT8_FMTx__ = "hhx"
const __UINT8_MAX__ = 255
const __UINTMAX_C_SUFFIX__ = "UL"
const __UINTMAX_FMTX__ = "lX"
const __UINTMAX_FMTo__ = "lo"
const __UINTMAX_FMTu__ = "lu"
const __UINTMAX_FMTx__ = "lx"
const __UINTMAX_MAX__ = 18446744073709551615
const __UINTMAX_WIDTH__ = 64
const __UINTPTR_FMTX__ = "lX"
const __UINTPTR_FMTo__ = "lo"
const __UINTPTR_FMTu__ = "lu"
const __UINTPTR_FMTx__ = "lx"
const __UINTPTR_MAX__ = 18446744073709551615
const __UINTPTR_WIDTH__ = 64
const __UINT_FAST16_FMTX__ = "hX"
const __UINT_FAST16_FMTo__ = "ho"
const __UINT_FAST16_FMTu__ = "hu"
const __UINT_FAST16_FMTx__ = "hx"
const __UINT_FAST16_MAX__ = 65535
const __UINT_FAST32_FMTX__ = "X"
const __UINT_FAST32_FMTo__ = "o"
const __UINT_FAST32_FMTu__ = "u"
const __UINT_FAST32_FMTx__ = "x"
const __UINT_FAST32_MAX__ = 4294967295
const __UINT_FAST64_FMTX__ = "lX"
const __UINT_FAST64_FMTo__ = "lo"
const __UINT_FAST64_FMTu__ = "lu"
const __UINT_FAST64_FMTx__ = "lx"
const __UINT_FAST64_MAX__ = 18446744073709551615
const __UINT_FAST8_FMTX__ = "hhX"
const __UINT_FAST8_FMTo__ = "hho"
const __UINT_FAST8_FMTu__ = "hhu"
const __UINT_FAST8_FMTx__ = "hhx"
const __UINT_FAST8_MAX__ = 255
const __UINT_LEAST16_FMTX__ = "hX"
const __UINT_LEAST16_FMTo__ = "ho"
const __UINT_LEAST16_FMTu__ = "hu"
const __UINT_LEAST16_FMTx__ = "hx"
const __UINT_LEAST16_MAX__ = 65535
const __UINT_LEAST32_FMTX__ = "X"
const __UINT_LEAST32_FMTo__ = "o"
const __UINT_LEAST32_FMTu__ = "u"
const __UINT_LEAST32_FMTx__ = "x"
const __UINT_LEAST32_MAX__ = 4294967295
const __UINT_LEAST64_FMTX__ = "lX"
const __UINT_LEAST64_FMTo__ = "lo"
const __UINT_LEAST64_FMTu__ = "lu"
const __UINT_LEAST64_FMTx__ = "lx"
const __UINT_LEAST64_MAX__ = 18446744073709551615
const __UINT_LEAST8_FMTX__ = "hhX"
const __UINT_LEAST8_FMTo__ = "hho"
const __UINT_LEAST8_FMTu__ = "hhu"
const __UINT_LEAST8_FMTx__ = "hhx"
const __UINT_LEAST8_MAX__ = 255
const __USE_TIME_BITS64 = 1
const __VERSION__ = "Ubuntu Clang 18.1.3 (1ubuntu1)"
const __WCHAR_MAX__ = 2147483647
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 4294967295
const __WINT_UNSIGNED__ = 1
const __WINT_WIDTH__ = 32
const __amd64 = 1
const __amd64__ = 1
const __bool_true_false_are_defined = 1
const __clang__ = 1
const __clang_literal_encoding__ = "UTF-8"
const __clang_major__ = 18
const __clang_minor__ = 1
const __clang_patchlevel__ = 3
const __clang_version__ = "18.1.3 (1ubuntu1)"
const __clang_wide_literal_encoding__ = "UTF-32"
const __code_model_small__ = 1
const __gnu_linux__ = 1
const __inline = "inline"
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __llvm__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict = "restrict"
const __restrict_arr = "restrict"
const __tune_k8__ = 1
const __unix = 1
const __unix__ = 1
const __x86_64 = 1
const __x86_64__ = 1
const alloca1 = "__builtin_alloca"
const bool1 = "_Bool"
const chan1 = "chan_token"
const defer1 = "defer_token"
const fallthrough1 = "fallthrough_token"
const false1 = 0
const func1 = "func_token"
const go1 = "go_token"
const import1 = "import_token"
const interface1 = "interface_token"
const linux = 1
const map1 = "map_token"
const package1 = "package_token"
const range1 = "range_token"
const select2 = "select_token"
const static_assert = "_Static_assert"
const true1 = 1
const ts_builtin_sym_end = 0
const ts_calloc = "calloc"
const ts_free = "free"
const ts_malloc = "malloc"
const ts_realloc = "realloc"
const type1 = "type_token"
const unix = 1
const var1 = "var_token"

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type size_t = uint64

type ssize_t = int64

type off_t = int64

type va_list = uintptr

type __isoc_va_list = uintptr

type fpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]int8
}

type _G_fpos64_t = fpos_t

type cookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type _IO_cookie_io_functions_t = cookie_io_functions_t

type wchar_t = int32

type div_t = struct {
	Fquot int32
	Frem  int32
}

type ldiv_t = struct {
	Fquot int64
	Frem  int64
}

type lldiv_t = struct {
	Fquot int64
	Frem  int64
}

type uintptr_t = uint64

type intptr_t = int64

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type intmax_t = int64

type uint8_t = uint8

type uint16_t = uint16

type uint32_t = uint32

type uint64_t = uint64

type uintmax_t = uint64

type int_fast8_t = int8

type int_fast64_t = int64

type int_least8_t = int8

type int_least16_t = int16

type int_least32_t = int32

type int_least64_t = int64

type uint_fast8_t = uint8

type uint_fast64_t = uint64

type uint_least8_t = uint8

type uint_least16_t = uint16

type uint_least32_t = uint32

type uint_least64_t = uint64

type int_fast16_t = int32

type int_fast32_t = int32

type uint_fast16_t = uint32

type uint_fast32_t = uint32

type locale_t = uintptr

type Array = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type TagType = int32

const AREA = 0
const BASE = 1
const BASEFONT = 2
const BGSOUND = 3
const BR = 4
const COL = 5
const COMMAND = 6
const EMBED = 7
const FRAME = 8
const HR = 9
const IMAGE = 10
const IMG = 11
const INPUT = 12
const ISINDEX = 13
const KEYGEN = 14
const LINK = 15
const MENUITEM = 16
const META = 17
const NEXTID = 18
const PARAM = 19
const SOURCE = 20
const TRACK = 21
const WBR = 22
const END_OF_VOID_TAGS = 23
const A = 24
const ABBR = 25
const ADDRESS = 26
const ARTICLE = 27
const ASIDE = 28
const AUDIO = 29
const B = 30
const BDI = 31
const BDO = 32
const BLOCKQUOTE = 33
const BODY = 34
const BUTTON = 35
const CANVAS = 36
const CAPTION = 37
const CITE = 38
const CODE = 39
const COLGROUP = 40
const DATA = 41
const DATALIST = 42
const DD = 43
const DEL = 44
const DETAILS = 45
const DFN = 46
const DIALOG = 47
const DIV = 48
const DL = 49
const DT = 50
const EM = 51
const FIELDSET = 52
const FIGCAPTION = 53
const FIGURE = 54
const FOOTER = 55
const FORM = 56
const H1 = 57
const H2 = 58
const H3 = 59
const H4 = 60
const H5 = 61
const H6 = 62
const HEAD = 63
const HEADER = 64
const HGROUP = 65
const HTML = 66
const I = 67
const IFRAME = 68
const INS = 69
const KBD = 70
const LABEL = 71
const LEGEND = 72
const LI = 73
const MAIN = 74
const MAP = 75
const MARK = 76
const MATH = 77
const MENU = 78
const METER = 79
const NAV = 80
const NOSCRIPT = 81
const OBJECT = 82
const OL = 83
const OPTGROUP = 84
const OPTION = 85
const OUTPUT = 86
const P = 87
const PICTURE = 88
const PRE = 89
const PROGRESS = 90
const Q = 91
const RB = 92
const RP = 93
const RT = 94
const RTC = 95
const RUBY = 96
const S = 97
const SAMP = 98
const SCRIPT = 99
const SECTION = 100
const SELECT = 101
const SLOT = 102
const SMALL = 103
const SPAN = 104
const STRONG = 105
const STYLE = 106
const SUB = 107
const SUMMARY = 108
const SUP = 109
const SVG = 110
const TABLE = 111
const TBODY = 112
const TD = 113
const TEMPLATE = 114
const TEXTAREA = 115
const TFOOT = 116
const TH = 117
const THEAD = 118
const TIME = 119
const TITLE = 120
const TR = 121
const U = 122
const UL = 123
const VAR = 124
const VIDEO = 125
const CUSTOM = 126
const END_ = 127

type String = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type TagMapEntry = struct {
	Ftag_name [16]int8
	Ftag_type TagType
}

type Tag = struct {
	Ftype_token      TagType
	Fcustom_tag_name String
}

var TAG_TYPES_BY_TAG_NAME = [126]TagMapEntry{
	0: {
		Ftag_name: [16]int8{'A', 'R', 'E', 'A'},
	},
	1: {
		Ftag_name: [16]int8{'B', 'A', 'S', 'E'},
		Ftag_type: int32(BASE),
	},
	2: {
		Ftag_name: [16]int8{'B', 'A', 'S', 'E', 'F', 'O', 'N', 'T'},
		Ftag_type: int32(BASEFONT),
	},
	3: {
		Ftag_name: [16]int8{'B', 'G', 'S', 'O', 'U', 'N', 'D'},
		Ftag_type: int32(BGSOUND),
	},
	4: {
		Ftag_name: [16]int8{'B', 'R'},
		Ftag_type: int32(BR),
	},
	5: {
		Ftag_name: [16]int8{'C', 'O', 'L'},
		Ftag_type: int32(COL),
	},
	6: {
		Ftag_name: [16]int8{'C', 'O', 'M', 'M', 'A', 'N', 'D'},
		Ftag_type: int32(COMMAND),
	},
	7: {
		Ftag_name: [16]int8{'E', 'M', 'B', 'E', 'D'},
		Ftag_type: int32(EMBED),
	},
	8: {
		Ftag_name: [16]int8{'F', 'R', 'A', 'M', 'E'},
		Ftag_type: int32(FRAME),
	},
	9: {
		Ftag_name: [16]int8{'H', 'R'},
		Ftag_type: int32(HR),
	},
	10: {
		Ftag_name: [16]int8{'I', 'M', 'A', 'G', 'E'},
		Ftag_type: int32(IMAGE),
	},
	11: {
		Ftag_name: [16]int8{'I', 'M', 'G'},
		Ftag_type: int32(IMG),
	},
	12: {
		Ftag_name: [16]int8{'I', 'N', 'P', 'U', 'T'},
		Ftag_type: int32(INPUT),
	},
	13: {
		Ftag_name: [16]int8{'I', 'S', 'I', 'N', 'D', 'E', 'X'},
		Ftag_type: int32(ISINDEX),
	},
	14: {
		Ftag_name: [16]int8{'K', 'E', 'Y', 'G', 'E', 'N'},
		Ftag_type: int32(KEYGEN),
	},
	15: {
		Ftag_name: [16]int8{'L', 'I', 'N', 'K'},
		Ftag_type: int32(LINK),
	},
	16: {
		Ftag_name: [16]int8{'M', 'E', 'N', 'U', 'I', 'T', 'E', 'M'},
		Ftag_type: int32(MENUITEM),
	},
	17: {
		Ftag_name: [16]int8{'M', 'E', 'T', 'A'},
		Ftag_type: int32(META),
	},
	18: {
		Ftag_name: [16]int8{'N', 'E', 'X', 'T', 'I', 'D'},
		Ftag_type: int32(NEXTID),
	},
	19: {
		Ftag_name: [16]int8{'P', 'A', 'R', 'A', 'M'},
		Ftag_type: int32(PARAM),
	},
	20: {
		Ftag_name: [16]int8{'S', 'O', 'U', 'R', 'C', 'E'},
		Ftag_type: int32(SOURCE),
	},
	21: {
		Ftag_name: [16]int8{'T', 'R', 'A', 'C', 'K'},
		Ftag_type: int32(TRACK),
	},
	22: {
		Ftag_name: [16]int8{'W', 'B', 'R'},
		Ftag_type: int32(WBR),
	},
	23: {
		Ftag_name: [16]int8{'A'},
		Ftag_type: int32(A),
	},
	24: {
		Ftag_name: [16]int8{'A', 'B', 'B', 'R'},
		Ftag_type: int32(ABBR),
	},
	25: {
		Ftag_name: [16]int8{'A', 'D', 'D', 'R', 'E', 'S', 'S'},
		Ftag_type: int32(ADDRESS),
	},
	26: {
		Ftag_name: [16]int8{'A', 'R', 'T', 'I', 'C', 'L', 'E'},
		Ftag_type: int32(ARTICLE),
	},
	27: {
		Ftag_name: [16]int8{'A', 'S', 'I', 'D', 'E'},
		Ftag_type: int32(ASIDE),
	},
	28: {
		Ftag_name: [16]int8{'A', 'U', 'D', 'I', 'O'},
		Ftag_type: int32(AUDIO),
	},
	29: {
		Ftag_name: [16]int8{'B'},
		Ftag_type: int32(B),
	},
	30: {
		Ftag_name: [16]int8{'B', 'D', 'I'},
		Ftag_type: int32(BDI),
	},
	31: {
		Ftag_name: [16]int8{'B', 'D', 'O'},
		Ftag_type: int32(BDO),
	},
	32: {
		Ftag_name: [16]int8{'B', 'L', 'O', 'C', 'K', 'Q', 'U', 'O', 'T', 'E'},
		Ftag_type: int32(BLOCKQUOTE),
	},
	33: {
		Ftag_name: [16]int8{'B', 'O', 'D', 'Y'},
		Ftag_type: int32(BODY),
	},
	34: {
		Ftag_name: [16]int8{'B', 'U', 'T', 'T', 'O', 'N'},
		Ftag_type: int32(BUTTON),
	},
	35: {
		Ftag_name: [16]int8{'C', 'A', 'N', 'V', 'A', 'S'},
		Ftag_type: int32(CANVAS),
	},
	36: {
		Ftag_name: [16]int8{'C', 'A', 'P', 'T', 'I', 'O', 'N'},
		Ftag_type: int32(CAPTION),
	},
	37: {
		Ftag_name: [16]int8{'C', 'I', 'T', 'E'},
		Ftag_type: int32(CITE),
	},
	38: {
		Ftag_name: [16]int8{'C', 'O', 'D', 'E'},
		Ftag_type: int32(CODE),
	},
	39: {
		Ftag_name: [16]int8{'C', 'O', 'L', 'G', 'R', 'O', 'U', 'P'},
		Ftag_type: int32(COLGROUP),
	},
	40: {
		Ftag_name: [16]int8{'D', 'A', 'T', 'A'},
		Ftag_type: int32(DATA),
	},
	41: {
		Ftag_name: [16]int8{'D', 'A', 'T', 'A', 'L', 'I', 'S', 'T'},
		Ftag_type: int32(DATALIST),
	},
	42: {
		Ftag_name: [16]int8{'D', 'D'},
		Ftag_type: int32(DD),
	},
	43: {
		Ftag_name: [16]int8{'D', 'E', 'L'},
		Ftag_type: int32(DEL),
	},
	44: {
		Ftag_name: [16]int8{'D', 'E', 'T', 'A', 'I', 'L', 'S'},
		Ftag_type: int32(DETAILS),
	},
	45: {
		Ftag_name: [16]int8{'D', 'F', 'N'},
		Ftag_type: int32(DFN),
	},
	46: {
		Ftag_name: [16]int8{'D', 'I', 'A', 'L', 'O', 'G'},
		Ftag_type: int32(DIALOG),
	},
	47: {
		Ftag_name: [16]int8{'D', 'I', 'V'},
		Ftag_type: int32(DIV),
	},
	48: {
		Ftag_name: [16]int8{'D', 'L'},
		Ftag_type: int32(DL),
	},
	49: {
		Ftag_name: [16]int8{'D', 'T'},
		Ftag_type: int32(DT),
	},
	50: {
		Ftag_name: [16]int8{'E', 'M'},
		Ftag_type: int32(EM),
	},
	51: {
		Ftag_name: [16]int8{'F', 'I', 'E', 'L', 'D', 'S', 'E', 'T'},
		Ftag_type: int32(FIELDSET),
	},
	52: {
		Ftag_name: [16]int8{'F', 'I', 'G', 'C', 'A', 'P', 'T', 'I', 'O', 'N'},
		Ftag_type: int32(FIGCAPTION),
	},
	53: {
		Ftag_name: [16]int8{'F', 'I', 'G', 'U', 'R', 'E'},
		Ftag_type: int32(FIGURE),
	},
	54: {
		Ftag_name: [16]int8{'F', 'O', 'O', 'T', 'E', 'R'},
		Ftag_type: int32(FOOTER),
	},
	55: {
		Ftag_name: [16]int8{'F', 'O', 'R', 'M'},
		Ftag_type: int32(FORM),
	},
	56: {
		Ftag_name: [16]int8{'H', '1'},
		Ftag_type: int32(H1),
	},
	57: {
		Ftag_name: [16]int8{'H', '2'},
		Ftag_type: int32(H2),
	},
	58: {
		Ftag_name: [16]int8{'H', '3'},
		Ftag_type: int32(H3),
	},
	59: {
		Ftag_name: [16]int8{'H', '4'},
		Ftag_type: int32(H4),
	},
	60: {
		Ftag_name: [16]int8{'H', '5'},
		Ftag_type: int32(H5),
	},
	61: {
		Ftag_name: [16]int8{'H', '6'},
		Ftag_type: int32(H6),
	},
	62: {
		Ftag_name: [16]int8{'H', 'E', 'A', 'D'},
		Ftag_type: int32(HEAD),
	},
	63: {
		Ftag_name: [16]int8{'H', 'E', 'A', 'D', 'E', 'R'},
		Ftag_type: int32(HEADER),
	},
	64: {
		Ftag_name: [16]int8{'H', 'G', 'R', 'O', 'U', 'P'},
		Ftag_type: int32(HGROUP),
	},
	65: {
		Ftag_name: [16]int8{'H', 'T', 'M', 'L'},
		Ftag_type: int32(HTML),
	},
	66: {
		Ftag_name: [16]int8{'I'},
		Ftag_type: int32(I),
	},
	67: {
		Ftag_name: [16]int8{'I', 'F', 'R', 'A', 'M', 'E'},
		Ftag_type: int32(IFRAME),
	},
	68: {
		Ftag_name: [16]int8{'I', 'N', 'S'},
		Ftag_type: int32(INS),
	},
	69: {
		Ftag_name: [16]int8{'K', 'B', 'D'},
		Ftag_type: int32(KBD),
	},
	70: {
		Ftag_name: [16]int8{'L', 'A', 'B', 'E', 'L'},
		Ftag_type: int32(LABEL),
	},
	71: {
		Ftag_name: [16]int8{'L', 'E', 'G', 'E', 'N', 'D'},
		Ftag_type: int32(LEGEND),
	},
	72: {
		Ftag_name: [16]int8{'L', 'I'},
		Ftag_type: int32(LI),
	},
	73: {
		Ftag_name: [16]int8{'M', 'A', 'I', 'N'},
		Ftag_type: int32(MAIN),
	},
	74: {
		Ftag_name: [16]int8{'M', 'A', 'P'},
		Ftag_type: int32(MAP),
	},
	75: {
		Ftag_name: [16]int8{'M', 'A', 'R', 'K'},
		Ftag_type: int32(MARK),
	},
	76: {
		Ftag_name: [16]int8{'M', 'A', 'T', 'H'},
		Ftag_type: int32(MATH),
	},
	77: {
		Ftag_name: [16]int8{'M', 'E', 'N', 'U'},
		Ftag_type: int32(MENU),
	},
	78: {
		Ftag_name: [16]int8{'M', 'E', 'T', 'E', 'R'},
		Ftag_type: int32(METER),
	},
	79: {
		Ftag_name: [16]int8{'N', 'A', 'V'},
		Ftag_type: int32(NAV),
	},
	80: {
		Ftag_name: [16]int8{'N', 'O', 'S', 'C', 'R', 'I', 'P', 'T'},
		Ftag_type: int32(NOSCRIPT),
	},
	81: {
		Ftag_name: [16]int8{'O', 'B', 'J', 'E', 'C', 'T'},
		Ftag_type: int32(OBJECT),
	},
	82: {
		Ftag_name: [16]int8{'O', 'L'},
		Ftag_type: int32(OL),
	},
	83: {
		Ftag_name: [16]int8{'O', 'P', 'T', 'G', 'R', 'O', 'U', 'P'},
		Ftag_type: int32(OPTGROUP),
	},
	84: {
		Ftag_name: [16]int8{'O', 'P', 'T', 'I', 'O', 'N'},
		Ftag_type: int32(OPTION),
	},
	85: {
		Ftag_name: [16]int8{'O', 'U', 'T', 'P', 'U', 'T'},
		Ftag_type: int32(OUTPUT),
	},
	86: {
		Ftag_name: [16]int8{'P'},
		Ftag_type: int32(P),
	},
	87: {
		Ftag_name: [16]int8{'P', 'I', 'C', 'T', 'U', 'R', 'E'},
		Ftag_type: int32(PICTURE),
	},
	88: {
		Ftag_name: [16]int8{'P', 'R', 'E'},
		Ftag_type: int32(PRE),
	},
	89: {
		Ftag_name: [16]int8{'P', 'R', 'O', 'G', 'R', 'E', 'S', 'S'},
		Ftag_type: int32(PROGRESS),
	},
	90: {
		Ftag_name: [16]int8{'Q'},
		Ftag_type: int32(Q),
	},
	91: {
		Ftag_name: [16]int8{'R', 'B'},
		Ftag_type: int32(RB),
	},
	92: {
		Ftag_name: [16]int8{'R', 'P'},
		Ftag_type: int32(RP),
	},
	93: {
		Ftag_name: [16]int8{'R', 'T'},
		Ftag_type: int32(RT),
	},
	94: {
		Ftag_name: [16]int8{'R', 'T', 'C'},
		Ftag_type: int32(RTC),
	},
	95: {
		Ftag_name: [16]int8{'R', 'U', 'B', 'Y'},
		Ftag_type: int32(RUBY),
	},
	96: {
		Ftag_name: [16]int8{'S'},
		Ftag_type: int32(S),
	},
	97: {
		Ftag_name: [16]int8{'S', 'A', 'M', 'P'},
		Ftag_type: int32(SAMP),
	},
	98: {
		Ftag_name: [16]int8{'S', 'C', 'R', 'I', 'P', 'T'},
		Ftag_type: int32(SCRIPT),
	},
	99: {
		Ftag_name: [16]int8{'S', 'E', 'C', 'T', 'I', 'O', 'N'},
		Ftag_type: int32(SECTION),
	},
	100: {
		Ftag_name: [16]int8{'S', 'E', 'L', 'E', 'C', 'T'},
		Ftag_type: int32(SELECT),
	},
	101: {
		Ftag_name: [16]int8{'S', 'L', 'O', 'T'},
		Ftag_type: int32(SLOT),
	},
	102: {
		Ftag_name: [16]int8{'S', 'M', 'A', 'L', 'L'},
		Ftag_type: int32(SMALL),
	},
	103: {
		Ftag_name: [16]int8{'S', 'P', 'A', 'N'},
		Ftag_type: int32(SPAN),
	},
	104: {
		Ftag_name: [16]int8{'S', 'T', 'R', 'O', 'N', 'G'},
		Ftag_type: int32(STRONG),
	},
	105: {
		Ftag_name: [16]int8{'S', 'T', 'Y', 'L', 'E'},
		Ftag_type: int32(STYLE),
	},
	106: {
		Ftag_name: [16]int8{'S', 'U', 'B'},
		Ftag_type: int32(SUB),
	},
	107: {
		Ftag_name: [16]int8{'S', 'U', 'M', 'M', 'A', 'R', 'Y'},
		Ftag_type: int32(SUMMARY),
	},
	108: {
		Ftag_name: [16]int8{'S', 'U', 'P'},
		Ftag_type: int32(SUP),
	},
	109: {
		Ftag_name: [16]int8{'S', 'V', 'G'},
		Ftag_type: int32(SVG),
	},
	110: {
		Ftag_name: [16]int8{'T', 'A', 'B', 'L', 'E'},
		Ftag_type: int32(TABLE),
	},
	111: {
		Ftag_name: [16]int8{'T', 'B', 'O', 'D', 'Y'},
		Ftag_type: int32(TBODY),
	},
	112: {
		Ftag_name: [16]int8{'T', 'D'},
		Ftag_type: int32(TD),
	},
	113: {
		Ftag_name: [16]int8{'T', 'E', 'M', 'P', 'L', 'A', 'T', 'E'},
		Ftag_type: int32(TEMPLATE),
	},
	114: {
		Ftag_name: [16]int8{'T', 'E', 'X', 'T', 'A', 'R', 'E', 'A'},
		Ftag_type: int32(TEXTAREA),
	},
	115: {
		Ftag_name: [16]int8{'T', 'F', 'O', 'O', 'T'},
		Ftag_type: int32(TFOOT),
	},
	116: {
		Ftag_name: [16]int8{'T', 'H'},
		Ftag_type: int32(TH),
	},
	117: {
		Ftag_name: [16]int8{'T', 'H', 'E', 'A', 'D'},
		Ftag_type: int32(THEAD),
	},
	118: {
		Ftag_name: [16]int8{'T', 'I', 'M', 'E'},
		Ftag_type: int32(TIME),
	},
	119: {
		Ftag_name: [16]int8{'T', 'I', 'T', 'L', 'E'},
		Ftag_type: int32(TITLE),
	},
	120: {
		Ftag_name: [16]int8{'T', 'R'},
		Ftag_type: int32(TR),
	},
	121: {
		Ftag_name: [16]int8{'U'},
		Ftag_type: int32(U),
	},
	122: {
		Ftag_name: [16]int8{'U', 'L'},
		Ftag_type: int32(UL),
	},
	123: {
		Ftag_name: [16]int8{'V', 'A', 'R'},
		Ftag_type: int32(VAR),
	},
	124: {
		Ftag_name: [16]int8{'V', 'I', 'D', 'E', 'O'},
		Ftag_type: int32(VIDEO),
	},
	125: {
		Ftag_name: [16]int8{'C', 'U', 'S', 'T', 'O', 'M'},
		Ftag_type: int32(CUSTOM),
	},
}
var TAG_TYPES_NOT_ALLOWED_IN_PARAGRAPHS = [26]TagType{
	0:  int32(ADDRESS),
	1:  int32(ARTICLE),
	2:  int32(ASIDE),
	3:  int32(BLOCKQUOTE),
	4:  int32(DETAILS),
	5:  int32(DIV),
	6:  int32(DL),
	7:  int32(FIELDSET),
	8:  int32(FIGCAPTION),
	9:  int32(FIGURE),
	10: int32(FOOTER),
	11: int32(FORM),
	12: int32(H1),
	13: int32(H2),
	14: int32(H3),
	15: int32(H4),
	16: int32(H5),
	17: int32(H6),
	18: int32(HEADER),
	19: int32(HR),
	20: int32(MAIN),
	21: int32(NAV),
	22: int32(OL),
	23: int32(P),
	24: int32(PRE),
	25: int32(SECTION),
}

func tag_type_for_name(tls *libc.TLS, tag_name uintptr) (r TagType) {
	var entry uintptr
	var i int32
	_, _ = entry, i
	i = 0
	for {
		if !(i < int32(126)) {
			break
		}
		entry = uintptr(unsafe.Pointer(&TAG_TYPES_BY_TAG_NAME)) + uintptr(i)*20
		if libc.Xstrlen(tls, entry) == uint64((*String)(unsafe.Pointer(tag_name)).Fsize) && libc.Xmemcmp(tls, (*String)(unsafe.Pointer(tag_name)).Fcontents, entry, uint64((*String)(unsafe.Pointer(tag_name)).Fsize)) == 0 {
			return (*TagMapEntry)(unsafe.Pointer(entry)).Ftag_type
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return int32(CUSTOM)
}

func tag_can_contain(tls *libc.TLS, self uintptr, other uintptr) (r uint8) {
	var child TagType
	var i int32
	_, _ = child, i
	child = (*Tag)(unsafe.Pointer(other)).Ftype_token
	switch (*Tag)(unsafe.Pointer(self)).Ftype_token {
	case int32(LI):
		goto _1
	case int32(DD):
		goto _2
	case int32(DT):
		goto _3
	case int32(P):
		goto _4
	case int32(COLGROUP):
		goto _5
	case int32(RP):
		goto _6
	case int32(RT):
		goto _7
	case int32(RB):
		goto _8
	case int32(OPTGROUP):
		goto _9
	case int32(TR):
		goto _10
	case int32(TH):
		goto _11
	case int32(TD):
		goto _12
	default:
		goto _13
	}
	goto _14
_1:
	;
	return libc.BoolUint8(child != int32(LI))
_3:
	;
_2:
	;
	return libc.BoolUint8(child != int32(DT) && child != int32(DD))
_4:
	;
	i = 0
_17:
	;
	if !(i < int32(26)) {
		goto _15
	}
	if child == TAG_TYPES_NOT_ALLOWED_IN_PARAGRAPHS[i] {
		return libc.BoolUint8(false1 != 0)
	}
	goto _16
_16:
	;
	i = i + 1
	goto _17
_15:
	;
	return libc.BoolUint8(true1 != 0)
_5:
	;
	return libc.BoolUint8(child == int32(COL))
_8:
	;
_7:
	;
_6:
	;
	return libc.BoolUint8(child != int32(RB) && child != int32(RT) && child != int32(RP))
_9:
	;
	return libc.BoolUint8(child != int32(OPTGROUP))
_10:
	;
	return libc.BoolUint8(child != int32(TR))
_12:
	;
_11:
	;
	return libc.BoolUint8(child != int32(TD) && child != int32(TH) && child != int32(TR))
_13:
	;
	return libc.BoolUint8(true1 != 0)
_14:
	;
	return r
}

type TSStateId = uint16

type TSSymbol = uint16

type TSFieldId = uint16

type TSLanguage = struct {
	Fversion                   uint32_t
	Fsymbol_count              uint32_t
	Falias_count               uint32_t
	Ftoken_count               uint32_t
	Fexternal_token_count      uint32_t
	Fstate_count               uint32_t
	Flarge_state_count         uint32_t
	Fproduction_id_count       uint32_t
	Ffield_count               uint32_t
	Fmax_alias_sequence_length uint16_t
	Fparse_table               uintptr
	Fsmall_parse_table         uintptr
	Fsmall_parse_table_map     uintptr
	Fparse_actions             uintptr
	Fsymbol_names              uintptr
	Ffield_names               uintptr
	Ffield_map_slices          uintptr
	Ffield_map_entries         uintptr
	Fsymbol_metadata           uintptr
	Fpublic_symbol_map         uintptr
	Falias_map                 uintptr
	Falias_sequences           uintptr
	Flex_modes                 uintptr
	Flex_fn                    uintptr
	Fkeyword_lex_fn            uintptr
	Fkeyword_capture_token     TSSymbol
	Fexternal_scanner          struct {
		Fstates      uintptr
		Fsymbol_map  uintptr
		Fcreate      uintptr
		Fdestroy     uintptr
		Fscan        uintptr
		Fserialize   uintptr
		Fdeserialize uintptr
	}
	Fprimary_state_ids uintptr
}

type TSFieldMapEntry = struct {
	Ffield_id    TSFieldId
	Fchild_index uint8_t
	Finherited   uint8
}

type TSFieldMapSlice = struct {
	Findex  uint16_t
	Flength uint16_t
}

type TSSymbolMetadata = struct {
	Fvisible   uint8
	Fnamed     uint8
	Fsupertype uint8
}

type TSLexer = struct {
	Flookahead                  int32_t
	Fresult_symbol              TSSymbol
	Fadvance                    uintptr
	Fmark_end                   uintptr
	Fget_column                 uintptr
	Fis_at_included_range_start uintptr
	Feof                        uintptr
	Flog                        uintptr
}

type TSParseActionType = int32

const TSParseActionTypeShift = 0
const TSParseActionTypeReduce = 1
const TSParseActionTypeAccept = 2
const TSParseActionTypeRecover = 3

type TSParseAction = struct {
	Freduce [0]struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}
	Ftype_token [0]uint8_t
	Fshift      struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}
	F__ccgo_pad3 [2]byte
}

type TSLexMode = struct {
	Flex_state          uint16_t
	Fexternal_lex_state uint16_t
}

type TSParseActionEntry = struct {
	Fentry [0]struct {
		Fcount    uint8_t
		Freusable uint8
	}
	Faction TSParseAction
}

type TSCharacterRange = struct {
	Fstart int32_t
	Fend   int32_t
}

type wint_t = uint32

type wctype_t = uint64

type wctrans_t = uintptr

type TokenType = int32

const START_TAG_NAME = 0
const SCRIPT_START_TAG_NAME = 1
const STYLE_START_TAG_NAME = 2
const END_TAG_NAME = 3
const ERRONEOUS_END_TAG_NAME = 4
const SELF_CLOSING_TAG_DELIMITER = 5
const IMPLICIT_END_TAG = 6
const RAW_TEXT = 7
const COMMENT = 8

type Scanner = struct {
	Ftags struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
}

func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
}

func skip(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
}

func serialize(tls *libc.TLS, scanner uintptr, buffer uintptr) (r uint32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var name_length, size, v1 uint32
	var tag Tag
	var _ /* serialized_tag_count at bp+2 */ uint16_t
	var _ /* tag_count at bp+0 */ uint16_t
	_, _, _, _ = name_length, size, tag, v1
	if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > libc.Uint32FromInt32(libc.Int32FromInt32(UINT16_MAX)) {
		v1 = libc.Uint32FromInt32(libc.Int32FromInt32(UINT16_MAX))
	} else {
		v1 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize
	}
	*(*uint16_t)(unsafe.Pointer(bp)) = uint16(v1)
	*(*uint16_t)(unsafe.Pointer(bp + 2)) = uint16(0)
	size = uint32(2)
	libc.Xmemcpy(tls, buffer+uintptr(size), bp, uint64(2))
	size = uint32(uint64(size) + libc.Uint64FromInt64(2))
	for {
		if !(libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(bp + 2))) < libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(bp)))) {
			break
		}
		tag = *(*Tag)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Ftags.Fcontents + uintptr(*(*uint16_t)(unsafe.Pointer(bp + 2)))*24))
		if tag.Ftype_token == int32(CUSTOM) {
			name_length = tag.Fcustom_tag_name.Fsize
			if name_length > libc.Uint32FromInt32(libc.Int32FromInt32(UINT8_MAX)) {
				name_length = libc.Uint32FromInt32(libc.Int32FromInt32(UINT8_MAX))
			}
			if size+uint32(2)+name_length >= uint32(TREE_SITTER_SERIALIZATION_BUFFER_SIZE) {
				break
			}
			v1 = size
			size = size + 1
			*(*int8)(unsafe.Pointer(buffer + uintptr(v1))) = int8(tag.Ftype_token)
			v1 = size
			size = size + 1
			*(*int8)(unsafe.Pointer(buffer + uintptr(v1))) = libc.Int8FromUint32(name_length)
			libc.Xstrncpy(tls, buffer+uintptr(size), tag.Fcustom_tag_name.Fcontents, uint64(name_length))
			size = size + name_length
		} else {
			if size+uint32(1) >= uint32(TREE_SITTER_SERIALIZATION_BUFFER_SIZE) {
				break
			}
			v1 = size
			size = size + 1
			*(*int8)(unsafe.Pointer(buffer + uintptr(v1))) = int8(tag.Ftype_token)
		}
		goto _2
	_2:
		;
		*(*uint16_t)(unsafe.Pointer(bp + 2)) = *(*uint16_t)(unsafe.Pointer(bp + 2)) + 1
	}
	libc.Xmemcpy(tls, buffer, bp+2, uint64(2))
	return size
}

func deserialize(tls *libc.TLS, scanner uintptr, buffer uintptr, length uint32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i, iter, size, v10 uint32
	var name_length uint16_t
	var new_capacity1, new_size, v14, v6 uint32_t
	var tag, v8 Tag
	var v2, v3, v4 uintptr
	var v5 size_t
	var _ /* serialized_tag_count at bp+2 */ uint16_t
	var _ /* tag at bp+8 */ Tag
	var _ /* tag_count at bp+0 */ uint16_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, iter, name_length, new_capacity1, new_size, size, tag, v10, v14, v2, v3, v4, v5, v6, v8
	i = uint32(0)
	for {
		if !(i < (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize) {
			break
		}
		v2 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fcontents + uintptr(i)*24
		if (*Tag)(unsafe.Pointer(v2)).Ftype_token == int32(CUSTOM) {
			v3 = v2 + 8
			if (*Array)(unsafe.Pointer(v3)).Fcontents != 0 {
				libc.Xfree(tls, (*Array)(unsafe.Pointer(v3)).Fcontents)
				(*Array)(unsafe.Pointer(v3)).Fcontents = libc.UintptrFromInt32(0)
				(*Array)(unsafe.Pointer(v3)).Fsize = uint32(0)
				(*Array)(unsafe.Pointer(v3)).Fcapacity = uint32(0)
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fsize = uint32(0)
	if length > uint32(0) {
		size = uint32(0)
		*(*uint16_t)(unsafe.Pointer(bp)) = uint16(0)
		*(*uint16_t)(unsafe.Pointer(bp + 2)) = uint16(0)
		libc.Xmemcpy(tls, bp+2, buffer+uintptr(size), uint64(2))
		size = uint32(uint64(size) + libc.Uint64FromInt64(2))
		libc.Xmemcpy(tls, bp, buffer+uintptr(size), uint64(2))
		size = uint32(uint64(size) + libc.Uint64FromInt64(2))
		v2 = scanner
		v5 = libc.Uint64FromInt64(24)
		v6 = uint32(*(*uint16_t)(unsafe.Pointer(bp)))
		if v6 > (*Array)(unsafe.Pointer(v2)).Fcapacity {
			if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
				(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v2)).Fcontents, uint64(v6)*v5)
			} else {
				(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xmalloc(tls, uint64(v6)*v5)
			}
			(*Array)(unsafe.Pointer(v2)).Fcapacity = v6
		}
		if libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(bp))) > 0 {
			iter = uint32(0)
			iter = uint32(0)
			for {
				if !(iter < uint32(*(*uint16_t)(unsafe.Pointer(bp + 2)))) {
					break
				}
				tag.Ftype_token = int32(END_)
				tag.Fcustom_tag_name = String{}
				v8 = tag
				goto _9
			_9:
				*(*Tag)(unsafe.Pointer(bp + 8)) = v8
				v10 = size
				size = size + 1
				(*(*Tag)(unsafe.Pointer(bp + 8))).Ftype_token = int32(*(*int8)(unsafe.Pointer(buffer + uintptr(v10))))
				if (*(*Tag)(unsafe.Pointer(bp + 8))).Ftype_token == int32(CUSTOM) {
					v10 = size
					size = size + 1
					name_length = uint16(libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(buffer + uintptr(v10)))))
					v2 = bp + 8 + 8
					v5 = libc.Uint64FromInt64(1)
					v6 = uint32(name_length)
					if v6 > (*Array)(unsafe.Pointer(v2)).Fcapacity {
						if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
							(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v2)).Fcontents, uint64(v6)*v5)
						} else {
							(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xmalloc(tls, uint64(v6)*v5)
						}
						(*Array)(unsafe.Pointer(v2)).Fcapacity = v6
					}
					(*(*Tag)(unsafe.Pointer(bp + 8))).Fcustom_tag_name.Fsize = uint32(name_length)
					libc.Xmemcpy(tls, (*(*Tag)(unsafe.Pointer(bp + 8))).Fcustom_tag_name.Fcontents, buffer+uintptr(size), uint64(name_length))
					size = size + uint32(name_length)
				}
				v2 = scanner
				new_size = (*Array)(unsafe.Pointer(v2)).Fsize + uint32(1)
				if new_size > (*Array)(unsafe.Pointer(v2)).Fcapacity {
					new_capacity1 = (*Array)(unsafe.Pointer(v2)).Fcapacity * uint32(2)
					if new_capacity1 < libc.Uint32FromInt32(8) {
						new_capacity1 = uint32(8)
					}
					if new_capacity1 < new_size {
						new_capacity1 = new_size
					}
					v3 = v2
					v5 = libc.Uint64FromInt64(24)
					v6 = new_capacity1
					if v6 > (*Array)(unsafe.Pointer(v3)).Fcapacity {
						if (*Array)(unsafe.Pointer(v3)).Fcontents != 0 {
							(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v3)).Fcontents, uint64(v6)*v5)
						} else {
							(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xmalloc(tls, uint64(v6)*v5)
						}
						(*Array)(unsafe.Pointer(v3)).Fcapacity = v6
					}
				}
				v4 = scanner + 8
				v14 = *(*uint32_t)(unsafe.Pointer(v4))
				*(*uint32_t)(unsafe.Pointer(v4)) = *(*uint32_t)(unsafe.Pointer(v4)) + 1
				*(*Tag)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(scanner)).Fcontents + uintptr(v14)*24)) = *(*Tag)(unsafe.Pointer(bp + 8))
				goto _7
			_7:
				;
				iter = iter + 1
			}
			// add zero tags if we didn't read enough, this is because the
			// buffer had no more room but we held more tags.
			for {
				if !(iter < uint32(*(*uint16_t)(unsafe.Pointer(bp)))) {
					break
				}
				v2 = scanner
				new_size = (*Array)(unsafe.Pointer(v2)).Fsize + uint32(1)
				if new_size > (*Array)(unsafe.Pointer(v2)).Fcapacity {
					new_capacity1 = (*Array)(unsafe.Pointer(v2)).Fcapacity * uint32(2)
					if new_capacity1 < libc.Uint32FromInt32(8) {
						new_capacity1 = uint32(8)
					}
					if new_capacity1 < new_size {
						new_capacity1 = new_size
					}
					v3 = v2
					v5 = libc.Uint64FromInt64(24)
					v6 = new_capacity1
					if v6 > (*Array)(unsafe.Pointer(v3)).Fcapacity {
						if (*Array)(unsafe.Pointer(v3)).Fcontents != 0 {
							(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v3)).Fcontents, uint64(v6)*v5)
						} else {
							(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xmalloc(tls, uint64(v6)*v5)
						}
						(*Array)(unsafe.Pointer(v3)).Fcapacity = v6
					}
				}
				v4 = scanner + 8
				v14 = *(*uint32_t)(unsafe.Pointer(v4))
				*(*uint32_t)(unsafe.Pointer(v4)) = *(*uint32_t)(unsafe.Pointer(v4)) + 1
				tag.Ftype_token = int32(END_)
				tag.Fcustom_tag_name = String{}
				v8 = tag
				goto _29
			_29:
				*(*Tag)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(scanner)).Fcontents + uintptr(v14)*24)) = v8
				goto _21
			_21:
				;
				iter = iter + 1
			}
		}
	}
}

func scan_tag_name(tls *libc.TLS, lexer uintptr) (r String) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var new_capacity1, new_size, v4, v5 uint32_t
	var v1, v2, v6 uintptr
	var v3 size_t
	var _ /* tag_name at bp+0 */ String
	_, _, _, _, _, _, _, _ = new_capacity1, new_size, v1, v2, v3, v4, v5, v6
	*(*String)(unsafe.Pointer(bp)) = String{}
	for libc.Xiswalnum(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('-') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(':') {
		v1 = bp
		new_size = (*Array)(unsafe.Pointer(v1)).Fsize + uint32(1)
		if new_size > (*Array)(unsafe.Pointer(v1)).Fcapacity {
			new_capacity1 = (*Array)(unsafe.Pointer(v1)).Fcapacity * uint32(2)
			if new_capacity1 < libc.Uint32FromInt32(8) {
				new_capacity1 = uint32(8)
			}
			if new_capacity1 < new_size {
				new_capacity1 = new_size
			}
			v2 = v1
			v3 = libc.Uint64FromInt64(1)
			v4 = new_capacity1
			if v4 > (*Array)(unsafe.Pointer(v2)).Fcapacity {
				if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
					(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v2)).Fcontents, uint64(v4)*v3)
				} else {
					(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xmalloc(tls, uint64(v4)*v3)
				}
				(*Array)(unsafe.Pointer(v2)).Fcapacity = v4
			}
		}
		v6 = bp + 8
		v5 = *(*uint32_t)(unsafe.Pointer(v6))
		*(*uint32_t)(unsafe.Pointer(v6)) = *(*uint32_t)(unsafe.Pointer(v6)) + 1
		*(*int8)(unsafe.Pointer((*String)(unsafe.Pointer(bp)).Fcontents + uintptr(v5))) = libc.Int8FromUint32(libc.Xtowupper(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)))
		advance(tls, lexer)
	}
	return *(*String)(unsafe.Pointer(bp))
}

func scan_comment(tls *libc.TLS, lexer uintptr) (r uint8) {
	var dashes uint32
	_ = dashes
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('-') {
		return libc.BoolUint8(false1 != 0)
	}
	advance(tls, lexer)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('-') {
		return libc.BoolUint8(false1 != 0)
	}
	advance(tls, lexer)
	dashes = uint32(0)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
		case int32('-'):
			dashes = dashes + 1
		case int32('>'):
			if dashes >= uint32(2) {
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(COMMENT)
				advance(tls, lexer)
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				return libc.BoolUint8(true1 != 0)
			}
			fallthrough
		default:
			dashes = uint32(0)
		}
		advance(tls, lexer)
	}
	return libc.BoolUint8(false1 != 0)
}

func scan_raw_text(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	var delimiter_index uint32
	var end_delimiter, v1 uintptr
	var v2 bool
	_, _, _, _ = delimiter_index, end_delimiter, v1, v2
	if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize == uint32(0) {
		return libc.BoolUint8(false1 != 0)
	}
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
	if v2 = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fsize; !v2 {
		libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+64, int32(150), uintptr(unsafe.Pointer(&__func__)))
	}
	_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
	if (*Tag)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fcontents+uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24)).Ftype_token == int32(SCRIPT) {
		v1 = __ccgo_ts + 75
	} else {
		v1 = __ccgo_ts + 84
	}
	end_delimiter = v1
	delimiter_index = uint32(0)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		if libc.Xtowupper(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) == libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(end_delimiter + uintptr(delimiter_index)))) {
			delimiter_index = delimiter_index + 1
			if uint64(delimiter_index) == libc.Xstrlen(tls, end_delimiter) {
				break
			}
			advance(tls, lexer)
		} else {
			delimiter_index = uint32(0)
			advance(tls, lexer)
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		}
	}
	(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(RAW_TEXT)
	return libc.BoolUint8(true1 != 0)
}

var __func__ = [14]int8{'s', 'c', 'a', 'n', '_', 'r', 'a', 'w', '_', 't', 'e', 'x', 't'}

func pop_tag(tls *libc.TLS, scanner uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var v1 uint32_t
	var v2, v3, v4 uintptr
	var _ /* popped_tag at bp+0 */ Tag
	_, _, _, _ = v1, v2, v3, v4
	v2 = scanner + 8
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*Tag)(unsafe.Pointer(bp)) = *(*Tag)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fcontents + uintptr(v1)*24))
	v3 = bp
	if (*Tag)(unsafe.Pointer(v3)).Ftype_token == int32(CUSTOM) {
		v4 = v3 + 8
		if (*Array)(unsafe.Pointer(v4)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v4)).Fcontents)
			(*Array)(unsafe.Pointer(v4)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v4)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v4)).Fcapacity = uint32(0)
		}
	}
}

func scan_implicit_end_tag(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var i uint32
	var is_closing_tag, v3 uint8
	var parent, v1, v10, v14, v15, v6 uintptr
	var tag, tag1, v11, v8 Tag
	var v2, v5 bool
	var v7 String
	var _ /* next_tag at bp+32 */ Tag
	var _ /* tag_name at bp+16 */ String
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, is_closing_tag, parent, tag, tag1, v1, v10, v11, v14, v15, v2, v3, v5, v6, v7, v8
	if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize == uint32(0) {
		v1 = libc.UintptrFromInt32(0)
	} else {
		if v2 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize; !v2 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+64, int32(177), uintptr(unsafe.Pointer(&__func__6)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v1 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fcontents + uintptr((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24
	}
	parent = v1
	is_closing_tag = libc.BoolUint8(false1 != 0)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
		is_closing_tag = libc.BoolUint8(true1 != 0)
		advance(tls, lexer)
	} else {
		if v2 = parent != 0; v2 {
			v3 = libc.BoolUint8((*Tag)(unsafe.Pointer(parent)).Ftype_token < int32(END_OF_VOID_TAGS))
			goto _4
		_4:
		}
		if v2 && v3 != 0 {
			pop_tag(tls, scanner)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(IMPLICIT_END_TAG)
			return libc.BoolUint8(true1 != 0)
		}
	}
	*(*String)(unsafe.Pointer(bp + 16)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp + 16))).Fsize == uint32(0) && !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
		v1 = bp + 16
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
		return libc.BoolUint8(false1 != 0)
	}
	v7 = *(*String)(unsafe.Pointer(bp + 16))
	*(*String)(unsafe.Pointer(bp)) = v7
	tag.Ftype_token = int32(END_)
	tag.Fcustom_tag_name = String{}
	v8 = tag
	goto _9
_9:
	tag1 = v8
	tag1.Ftype_token = tag_type_for_name(tls, bp)
	if tag1.Ftype_token == int32(CUSTOM) {
		tag1.Fcustom_tag_name = v7
	} else {
		v1 = bp
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
	}
	v11 = tag1
	goto _12
_12:
	*(*Tag)(unsafe.Pointer(bp + 32)) = v11
	if is_closing_tag != 0 {
		// The tag correctly closes the topmost element on the stack
		if v5 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0); v5 {
			if v2 = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize; !v2 {
				libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+64, int32(201), uintptr(unsafe.Pointer(&__func__6)))
			}
			_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
			v6 = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fcontents + uintptr((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24
			v10 = bp + 32
			if (*Tag)(unsafe.Pointer(v6)).Ftype_token != (*Tag)(unsafe.Pointer(v10)).Ftype_token {
				v3 = libc.BoolUint8(false1 != 0)
				goto _17
			}
			if (*Tag)(unsafe.Pointer(v6)).Ftype_token == int32(CUSTOM) {
				if (*Tag)(unsafe.Pointer(v6)).Fcustom_tag_name.Fsize != (*Tag)(unsafe.Pointer(v10)).Fcustom_tag_name.Fsize {
					v3 = libc.BoolUint8(false1 != 0)
					goto _17
				}
				if libc.Xmemcmp(tls, (*Tag)(unsafe.Pointer(v6)).Fcustom_tag_name.Fcontents, (*Tag)(unsafe.Pointer(v10)).Fcustom_tag_name.Fcontents, uint64((*Tag)(unsafe.Pointer(v6)).Fcustom_tag_name.Fsize)) != libc.Int32FromInt32(0) {
					v3 = libc.BoolUint8(false1 != 0)
					goto _17
				}
			}
			v3 = libc.BoolUint8(true1 != 0)
			goto _17
		_17:
		}
		if v5 && v3 != 0 {
			v14 = bp + 32
			if (*Tag)(unsafe.Pointer(v14)).Ftype_token == int32(CUSTOM) {
				v15 = v14 + 8
				if (*Array)(unsafe.Pointer(v15)).Fcontents != 0 {
					libc.Xfree(tls, (*Array)(unsafe.Pointer(v15)).Fcontents)
					(*Array)(unsafe.Pointer(v15)).Fcontents = libc.UintptrFromInt32(0)
					(*Array)(unsafe.Pointer(v15)).Fsize = uint32(0)
					(*Array)(unsafe.Pointer(v15)).Fcapacity = uint32(0)
				}
			}
			return libc.BoolUint8(false1 != 0)
		}
		// Otherwise, dig deeper and queue implicit end tags (to be nice in
		// the case of malformed HTML)
		i = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize
		for {
			if !(i > uint32(0)) {
				break
			}
			if (*(*Tag)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Ftags.Fcontents + uintptr(i-uint32(1))*24))).Ftype_token == (*(*Tag)(unsafe.Pointer(bp + 32))).Ftype_token {
				pop_tag(tls, scanner)
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(IMPLICIT_END_TAG)
				v1 = bp + 32
				if (*Tag)(unsafe.Pointer(v1)).Ftype_token == int32(CUSTOM) {
					v6 = v1 + 8
					if (*Array)(unsafe.Pointer(v6)).Fcontents != 0 {
						libc.Xfree(tls, (*Array)(unsafe.Pointer(v6)).Fcontents)
						(*Array)(unsafe.Pointer(v6)).Fcontents = libc.UintptrFromInt32(0)
						(*Array)(unsafe.Pointer(v6)).Fsize = uint32(0)
						(*Array)(unsafe.Pointer(v6)).Fcapacity = uint32(0)
					}
				}
				return libc.BoolUint8(true1 != 0)
			}
			goto _21
		_21:
			;
			i = i - 1
		}
	} else {
		if parent != 0 && (!(tag_can_contain(tls, parent, bp+32) != 0) || ((*Tag)(unsafe.Pointer(parent)).Ftype_token == int32(HTML) || (*Tag)(unsafe.Pointer(parent)).Ftype_token == int32(HEAD) || (*Tag)(unsafe.Pointer(parent)).Ftype_token == int32(BODY)) && (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
			pop_tag(tls, scanner)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(IMPLICIT_END_TAG)
			v1 = bp + 32
			if (*Tag)(unsafe.Pointer(v1)).Ftype_token == int32(CUSTOM) {
				v6 = v1 + 8
				if (*Array)(unsafe.Pointer(v6)).Fcontents != 0 {
					libc.Xfree(tls, (*Array)(unsafe.Pointer(v6)).Fcontents)
					(*Array)(unsafe.Pointer(v6)).Fcontents = libc.UintptrFromInt32(0)
					(*Array)(unsafe.Pointer(v6)).Fsize = uint32(0)
					(*Array)(unsafe.Pointer(v6)).Fcapacity = uint32(0)
				}
			}
			return libc.BoolUint8(true1 != 0)
		}
	}
	v1 = bp + 32
	if (*Tag)(unsafe.Pointer(v1)).Ftype_token == int32(CUSTOM) {
		v6 = v1 + 8
		if (*Array)(unsafe.Pointer(v6)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v6)).Fcontents)
			(*Array)(unsafe.Pointer(v6)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v6)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v6)).Fcapacity = uint32(0)
		}
	}
	return libc.BoolUint8(false1 != 0)
}

var __func__6 = [22]int8{'s', 'c', 'a', 'n', '_', 'i', 'm', 'p', 'l', 'i', 'c', 'i', 't', '_', 'e', 'n', 'd', '_', 't', 'a', 'g'}

func scan_start_tag_name(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var new_capacity1, new_size, v11, v12 uint32_t
	var tag, tag1, tag2, v3, v6 Tag
	var v1, v5, v8, v9 uintptr
	var v10 size_t
	var v2 String
	var _ /* tag_name at bp+16 */ String
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = new_capacity1, new_size, tag, tag1, tag2, v1, v10, v11, v12, v2, v3, v5, v6, v8, v9
	*(*String)(unsafe.Pointer(bp + 16)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp + 16))).Fsize == uint32(0) {
		v1 = bp + 16
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
		return libc.BoolUint8(false1 != 0)
	}
	v2 = *(*String)(unsafe.Pointer(bp + 16))
	*(*String)(unsafe.Pointer(bp)) = v2
	tag.Ftype_token = int32(END_)
	tag.Fcustom_tag_name = String{}
	v3 = tag
	goto _4
_4:
	tag1 = v3
	tag1.Ftype_token = tag_type_for_name(tls, bp)
	if tag1.Ftype_token == int32(CUSTOM) {
		tag1.Fcustom_tag_name = v2
	} else {
		v1 = bp
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
	}
	v6 = tag1
	goto _7
_7:
	tag2 = v6
	v5 = scanner
	new_size = (*Array)(unsafe.Pointer(v5)).Fsize + uint32(1)
	if new_size > (*Array)(unsafe.Pointer(v5)).Fcapacity {
		new_capacity1 = (*Array)(unsafe.Pointer(v5)).Fcapacity * uint32(2)
		if new_capacity1 < libc.Uint32FromInt32(8) {
			new_capacity1 = uint32(8)
		}
		if new_capacity1 < new_size {
			new_capacity1 = new_size
		}
		v8 = v5
		v10 = libc.Uint64FromInt64(24)
		v11 = new_capacity1
		if v11 > (*Array)(unsafe.Pointer(v8)).Fcapacity {
			if (*Array)(unsafe.Pointer(v8)).Fcontents != 0 {
				(*Array)(unsafe.Pointer(v8)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v8)).Fcontents, uint64(v11)*v10)
			} else {
				(*Array)(unsafe.Pointer(v8)).Fcontents = libc.Xmalloc(tls, uint64(v11)*v10)
			}
			(*Array)(unsafe.Pointer(v8)).Fcapacity = v11
		}
	}
	v9 = scanner + 8
	v12 = *(*uint32_t)(unsafe.Pointer(v9))
	*(*uint32_t)(unsafe.Pointer(v9)) = *(*uint32_t)(unsafe.Pointer(v9)) + 1
	*(*Tag)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(scanner)).Fcontents + uintptr(v12)*24)) = tag2
	switch tag2.Ftype_token {
	case int32(SCRIPT):
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(SCRIPT_START_TAG_NAME)
	case int32(STYLE):
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(STYLE_START_TAG_NAME)
	default:
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(START_TAG_NAME)
		break
	}
	return libc.BoolUint8(true1 != 0)
}

func scan_end_tag_name(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var tag, tag1, v3, v6 Tag
	var v1, v5, v9 uintptr
	var v11 uint8
	var v13, v8 bool
	var v2 String
	var _ /* tag at bp+32 */ Tag
	var _ /* tag_name at bp+16 */ String
	_, _, _, _, _, _, _, _, _, _, _ = tag, tag1, v1, v11, v13, v2, v3, v5, v6, v8, v9
	*(*String)(unsafe.Pointer(bp + 16)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp + 16))).Fsize == uint32(0) {
		v1 = bp + 16
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
		return libc.BoolUint8(false1 != 0)
	}
	v2 = *(*String)(unsafe.Pointer(bp + 16))
	*(*String)(unsafe.Pointer(bp)) = v2
	tag.Ftype_token = int32(END_)
	tag.Fcustom_tag_name = String{}
	v3 = tag
	goto _4
_4:
	tag1 = v3
	tag1.Ftype_token = tag_type_for_name(tls, bp)
	if tag1.Ftype_token == int32(CUSTOM) {
		tag1.Fcustom_tag_name = v2
	} else {
		v1 = bp
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
	}
	v6 = tag1
	goto _7
_7:
	*(*Tag)(unsafe.Pointer(bp + 32)) = v6
	if v13 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0); v13 {
		if v8 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize; !v8 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+64, int32(265), uintptr(unsafe.Pointer(&__func__5)))
		}
		_ = v8 || libc.Bool(libc.Int32FromInt32(0) != 0)
		v5 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fcontents + uintptr((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(scanner)).Fsize-uint32(1))*24
		v9 = bp + 32
		if (*Tag)(unsafe.Pointer(v5)).Ftype_token != (*Tag)(unsafe.Pointer(v9)).Ftype_token {
			v11 = libc.BoolUint8(false1 != 0)
			goto _12
		}
		if (*Tag)(unsafe.Pointer(v5)).Ftype_token == int32(CUSTOM) {
			if (*Tag)(unsafe.Pointer(v5)).Fcustom_tag_name.Fsize != (*Tag)(unsafe.Pointer(v9)).Fcustom_tag_name.Fsize {
				v11 = libc.BoolUint8(false1 != 0)
				goto _12
			}
			if libc.Xmemcmp(tls, (*Tag)(unsafe.Pointer(v5)).Fcustom_tag_name.Fcontents, (*Tag)(unsafe.Pointer(v9)).Fcustom_tag_name.Fcontents, uint64((*Tag)(unsafe.Pointer(v5)).Fcustom_tag_name.Fsize)) != libc.Int32FromInt32(0) {
				v11 = libc.BoolUint8(false1 != 0)
				goto _12
			}
		}
		v11 = libc.BoolUint8(true1 != 0)
		goto _12
	_12:
	}
	if v13 && v11 != 0 {
		pop_tag(tls, scanner)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(END_TAG_NAME)
	} else {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ERRONEOUS_END_TAG_NAME)
	}
	v1 = bp + 32
	if (*Tag)(unsafe.Pointer(v1)).Ftype_token == int32(CUSTOM) {
		v5 = v1 + 8
		if (*Array)(unsafe.Pointer(v5)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v5)).Fcontents)
			(*Array)(unsafe.Pointer(v5)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v5)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v5)).Fcapacity = uint32(0)
		}
	}
	return libc.BoolUint8(true1 != 0)
}

var __func__5 = [18]int8{'s', 'c', 'a', 'n', '_', 'e', 'n', 'd', '_', 't', 'a', 'g', '_', 'n', 'a', 'm', 'e'}

func scan_self_closing_tag_delimiter(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	advance(tls, lexer)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('>') {
		advance(tls, lexer)
		if (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize > uint32(0) {
			pop_tag(tls, scanner)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(SELF_CLOSING_TAG_DELIMITER)
		}
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

func scan(tls *libc.TLS, scanner uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var v1 int32
	_ = v1
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(RAW_TEXT))) != 0 && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(START_TAG_NAME))) != 0) && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(END_TAG_NAME))) != 0) {
		return scan_raw_text(tls, scanner, lexer)
	}
	for libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
		skip(tls, lexer)
	}
	switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
	case int32('<'):
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		advance(tls, lexer)
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('!') {
			advance(tls, lexer)
			return scan_comment(tls, lexer)
		}
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(IMPLICIT_END_TAG))) != 0 {
			return scan_implicit_end_tag(tls, scanner, lexer)
		}
	case int32('\000'):
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(IMPLICIT_END_TAG))) != 0 {
			return scan_implicit_end_tag(tls, scanner, lexer)
		}
	case int32('/'):
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(SELF_CLOSING_TAG_DELIMITER))) != 0 {
			return scan_self_closing_tag_delimiter(tls, scanner, lexer)
		}
	default:
		if (*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(START_TAG_NAME))) != 0 || *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(END_TAG_NAME))) != 0) && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(RAW_TEXT))) != 0) {
			if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(START_TAG_NAME))) != 0 {
				v1 = libc.Int32FromUint8(scan_start_tag_name(tls, scanner, lexer))
			} else {
				v1 = libc.Int32FromUint8(scan_end_tag_name(tls, scanner, lexer))
			}
			return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
		}
	}
	return libc.BoolUint8(false1 != 0)
}

func tree_sitter_html_external_scanner_create(tls *libc.TLS) (r uintptr) {
	var scanner uintptr
	_ = scanner
	scanner = libc.Xcalloc(tls, uint64(1), uint64(16))
	return scanner
}

func tree_sitter_html_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	return scan(tls, scanner, lexer, valid_symbols)
}

func tree_sitter_html_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	return serialize(tls, scanner, buffer)
}

func tree_sitter_html_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	deserialize(tls, scanner, buffer, length)
}

func tree_sitter_html_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
	var i uint32
	var scanner, v2, v3 uintptr
	_, _, _, _ = i, scanner, v2, v3
	scanner = payload
	i = uint32(0)
	for {
		if !(i < (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fsize) {
			break
		}
		v2 = (*Scanner)(unsafe.Pointer(scanner)).Ftags.Fcontents + uintptr(i)*24
		if (*Tag)(unsafe.Pointer(v2)).Ftype_token == int32(CUSTOM) {
			v3 = v2 + 8
			if (*Array)(unsafe.Pointer(v3)).Fcontents != 0 {
				libc.Xfree(tls, (*Array)(unsafe.Pointer(v3)).Fcontents)
				(*Array)(unsafe.Pointer(v3)).Fcontents = libc.UintptrFromInt32(0)
				(*Array)(unsafe.Pointer(v3)).Fsize = uint32(0)
				(*Array)(unsafe.Pointer(v3)).Fcapacity = uint32(0)
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	v2 = scanner
	if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
		libc.Xfree(tls, (*Array)(unsafe.Pointer(v2)).Fcontents)
		(*Array)(unsafe.Pointer(v2)).Fcontents = libc.UintptrFromInt32(0)
		(*Array)(unsafe.Pointer(v2)).Fsize = uint32(0)
		(*Array)(unsafe.Pointer(v2)).Fcapacity = uint32(0)
	}
	libc.Xfree(tls, scanner)
}

type ts_symbol_identifiers = int32

const anon_sym_LT_BANG = 1
const aux_sym_doctype_token1 = 2
const anon_sym_GT = 3
const sym__doctype = 4
const anon_sym_LT = 5
const anon_sym_SLASH_GT = 6
const anon_sym_LT_SLASH = 7
const anon_sym_EQ = 8
const sym_attribute_name = 9
const sym_attribute_value = 10
const sym_entity = 11
const anon_sym_SQUOTE = 12
const aux_sym_quoted_attribute_value_token1 = 13
const anon_sym_DQUOTE = 14
const aux_sym_quoted_attribute_value_token2 = 15
const sym_text = 16
const sym__start_tag_name = 17
const sym__script_start_tag_name = 18
const sym__style_start_tag_name = 19
const sym__end_tag_name = 20
const sym_erroneous_end_tag_name = 21
const sym__implicit_end_tag = 22
const sym_raw_text = 23
const sym_comment = 24
const sym_document = 25
const sym_doctype = 26
const sym__node = 27
const sym_element = 28
const sym_script_element = 29
const sym_style_element = 30
const sym_start_tag = 31
const sym_script_start_tag = 32
const sym_style_start_tag = 33
const sym_self_closing_tag = 34
const sym_end_tag = 35
const sym_erroneous_end_tag = 36
const sym_attribute = 37
const sym_quoted_attribute_value = 38
const aux_sym_document_repeat1 = 39
const aux_sym_start_tag_repeat1 = 40

var ts_symbol_names = [41]uintptr{
	0:  __ccgo_ts + 92,
	1:  __ccgo_ts + 96,
	2:  __ccgo_ts + 99,
	3:  __ccgo_ts + 114,
	4:  __ccgo_ts + 116,
	5:  __ccgo_ts + 124,
	6:  __ccgo_ts + 126,
	7:  __ccgo_ts + 129,
	8:  __ccgo_ts + 132,
	9:  __ccgo_ts + 134,
	10: __ccgo_ts + 149,
	11: __ccgo_ts + 165,
	12: __ccgo_ts + 172,
	13: __ccgo_ts + 149,
	14: __ccgo_ts + 174,
	15: __ccgo_ts + 149,
	16: __ccgo_ts + 176,
	17: __ccgo_ts + 181,
	18: __ccgo_ts + 181,
	19: __ccgo_ts + 181,
	20: __ccgo_ts + 181,
	21: __ccgo_ts + 190,
	22: __ccgo_ts + 213,
	23: __ccgo_ts + 231,
	24: __ccgo_ts + 240,
	25: __ccgo_ts + 248,
	26: __ccgo_ts + 116,
	27: __ccgo_ts + 257,
	28: __ccgo_ts + 263,
	29: __ccgo_ts + 271,
	30: __ccgo_ts + 286,
	31: __ccgo_ts + 300,
	32: __ccgo_ts + 300,
	33: __ccgo_ts + 300,
	34: __ccgo_ts + 310,
	35: __ccgo_ts + 327,
	36: __ccgo_ts + 335,
	37: __ccgo_ts + 353,
	38: __ccgo_ts + 363,
	39: __ccgo_ts + 386,
	40: __ccgo_ts + 403,
}

var ts_symbol_map = [41]TSSymbol{
	1:  uint16(anon_sym_LT_BANG),
	2:  uint16(aux_sym_doctype_token1),
	3:  uint16(anon_sym_GT),
	4:  uint16(sym__doctype),
	5:  uint16(anon_sym_LT),
	6:  uint16(anon_sym_SLASH_GT),
	7:  uint16(anon_sym_LT_SLASH),
	8:  uint16(anon_sym_EQ),
	9:  uint16(sym_attribute_name),
	10: uint16(sym_attribute_value),
	11: uint16(sym_entity),
	12: uint16(anon_sym_SQUOTE),
	13: uint16(sym_attribute_value),
	14: uint16(anon_sym_DQUOTE),
	15: uint16(sym_attribute_value),
	16: uint16(sym_text),
	17: uint16(sym__start_tag_name),
	18: uint16(sym__start_tag_name),
	19: uint16(sym__start_tag_name),
	20: uint16(sym__start_tag_name),
	21: uint16(sym_erroneous_end_tag_name),
	22: uint16(sym__implicit_end_tag),
	23: uint16(sym_raw_text),
	24: uint16(sym_comment),
	25: uint16(sym_document),
	26: uint16(sym_doctype),
	27: uint16(sym__node),
	28: uint16(sym_element),
	29: uint16(sym_script_element),
	30: uint16(sym_style_element),
	31: uint16(sym_start_tag),
	32: uint16(sym_start_tag),
	33: uint16(sym_start_tag),
	34: uint16(sym_self_closing_tag),
	35: uint16(sym_end_tag),
	36: uint16(sym_erroneous_end_tag),
	37: uint16(sym_attribute),
	38: uint16(sym_quoted_attribute_value),
	39: uint16(aux_sym_document_repeat1),
	40: uint16(aux_sym_start_tag_repeat1),
}

var ts_symbol_metadata = [41]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	2: {},
	3: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	4: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	5: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	6: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	7: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	8: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	9: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	10: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	11: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	12: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	13: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	14: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	15: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	16: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	17: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	18: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	19: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	20: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	21: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	22: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	23: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	24: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	25: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	26: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	27: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	28: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	29: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	30: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	31: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	32: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	33: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	35: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	36: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	37: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	38: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	39: {},
	40: {},
}

var ts_alias_sequences = [1][4]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [94]TSStateId{
	1:  uint16(1),
	2:  uint16(2),
	3:  uint16(3),
	4:  uint16(3),
	5:  uint16(2),
	6:  uint16(6),
	7:  uint16(7),
	8:  uint16(6),
	9:  uint16(9),
	10: uint16(10),
	11: uint16(11),
	12: uint16(12),
	13: uint16(13),
	14: uint16(14),
	15: uint16(15),
	16: uint16(16),
	17: uint16(17),
	18: uint16(18),
	19: uint16(19),
	20: uint16(9),
	21: uint16(21),
	22: uint16(22),
	23: uint16(19),
	24: uint16(10),
	25: uint16(12),
	26: uint16(13),
	27: uint16(14),
	28: uint16(15),
	29: uint16(16),
	30: uint16(17),
	31: uint16(31),
	32: uint16(22),
	33: uint16(21),
	34: uint16(31),
	35: uint16(35),
	36: uint16(36),
	37: uint16(37),
	38: uint16(36),
	39: uint16(37),
	40: uint16(40),
	41: uint16(35),
	42: uint16(42),
	43: uint16(43),
	44: uint16(44),
	45: uint16(45),
	46: uint16(46),
	47: uint16(46),
	48: uint16(48),
	49: uint16(49),
	50: uint16(50),
	51: uint16(45),
	52: uint16(52),
	53: uint16(53),
	54: uint16(48),
	55: uint16(53),
	56: uint16(56),
	57: uint16(52),
	58: uint16(56),
	59: uint16(59),
	60: uint16(60),
	61: uint16(61),
	62: uint16(62),
	63: uint16(49),
	64: uint16(50),
	65: uint16(65),
	66: uint16(66),
	67: uint16(67),
	68: uint16(59),
	69: uint16(69),
	70: uint16(60),
	71: uint16(69),
	72: uint16(72),
	73: uint16(72),
	74: uint16(62),
	75: uint16(75),
	76: uint16(75),
	77: uint16(77),
	78: uint16(78),
	79: uint16(79),
	80: uint16(79),
	81: uint16(81),
	82: uint16(82),
	83: uint16(83),
	84: uint16(84),
	85: uint16(85),
	86: uint16(78),
	87: uint16(82),
	88: uint16(88),
	89: uint16(85),
	90: uint16(81),
	91: uint16(84),
	92: uint16(88),
	93: uint16(77),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i uint32_t
	var lookahead int32_t
	_, _, _, _, _ = eof, i, lookahead, result, skip
	result = libc.BoolUint8(false1 != 0)
	skip = libc.BoolUint8(false1 != 0)
	eof = libc.BoolUint8(false1 != 0)
	goto start
	goto next_state
next_state:
	;
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, skip)
	goto start
start:
	;
	skip = libc.BoolUint8(false1 != 0)
	lookahead = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
	eof = (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer)
	switch libc.Int32FromUint16(state) {
	case 0:
		if eof != 0 {
			state = uint16(18)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token[i]) == lookahead {
				state = map_token[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('"') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(70)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('"') {
			state = uint16(73)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(74)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('#') {
			state = uint16(12)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('\'') {
			state = uint16(70)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(71)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('/') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(22)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\'') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('>') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('C') || lookahead == int32('c') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('T') || lookahead == int32('t') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(16)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('Y') || lookahead == int32('y') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(14):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(14)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('&') && lookahead != int32('<') && lookahead != int32('>') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(15):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(20)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(16):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(17):
		if eof != 0 {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(24)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(18):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(19):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(20):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_doctype_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(20)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_doctype_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(22):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(23):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__doctype)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(24):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && lookahead != int32('/') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('\'') && (lookahead < int32('<') || int32('>') < lookahead) {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(71)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\'') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(74)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(14)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('&') && lookahead != int32('<') && lookahead != int32('>') {
			state = uint16(76)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [18]uint16_t{
	0:  uint16('"'),
	1:  uint16(73),
	2:  uint16('&'),
	3:  uint16(3),
	4:  uint16('\''),
	5:  uint16(70),
	6:  uint16('/'),
	7:  uint16(6),
	8:  uint16('<'),
	9:  uint16(24),
	10: uint16('='),
	11: uint16(27),
	12: uint16('>'),
	13: uint16(22),
	14: uint16('D'),
	15: uint16(9),
	16: uint16('d'),
	17: uint16(9),
}

var ts_lex_modes = [94]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	2: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	3: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	4: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	5: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	6: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	7: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	8: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	9: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	10: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	11: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	12: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	13: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	14: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	15: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	16: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	17: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	18: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	19: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	20: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	21: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	22: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	23: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	24: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	25: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	26: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	27: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	28: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	29: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	30: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	31: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(3),
	},
	32: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	33: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	34: {
		Flex_state:          uint16(17),
		Fexternal_lex_state: uint16(2),
	},
	35: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	36: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	37: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	38: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	39: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	40: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	41: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	42: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	43: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	44: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	45: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	46: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	47: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	48: {
		Fexternal_lex_state: uint16(5),
	},
	49: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	50: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	51: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	52: {
		Fexternal_lex_state: uint16(6),
	},
	53: {
		Fexternal_lex_state: uint16(5),
	},
	54: {
		Fexternal_lex_state: uint16(5),
	},
	55: {
		Fexternal_lex_state: uint16(5),
	},
	56: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	57: {
		Fexternal_lex_state: uint16(6),
	},
	58: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	59: {
		Fexternal_lex_state: uint16(7),
	},
	60: {
		Fexternal_lex_state: uint16(2),
	},
	61: {
		Fexternal_lex_state: uint16(5),
	},
	62: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	63: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	64: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(2),
	},
	65: {
		Fexternal_lex_state: uint16(5),
	},
	66: {
		Fexternal_lex_state: uint16(5),
	},
	67: {
		Fexternal_lex_state: uint16(5),
	},
	68: {
		Fexternal_lex_state: uint16(7),
	},
	69: {
		Fexternal_lex_state: uint16(2),
	},
	70: {
		Fexternal_lex_state: uint16(2),
	},
	71: {
		Fexternal_lex_state: uint16(2),
	},
	72: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(2),
	},
	73: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(2),
	},
	74: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	75: {
		Fexternal_lex_state: uint16(2),
	},
	76: {
		Fexternal_lex_state: uint16(2),
	},
	77: {
		Fexternal_lex_state: uint16(2),
	},
	78: {
		Fexternal_lex_state: uint16(2),
	},
	79: {
		Fexternal_lex_state: uint16(2),
	},
	80: {
		Fexternal_lex_state: uint16(2),
	},
	81: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(2),
	},
	82: {
		Fexternal_lex_state: uint16(2),
	},
	83: {
		Fexternal_lex_state: uint16(2),
	},
	84: {
		Fexternal_lex_state: uint16(2),
	},
	85: {
		Fexternal_lex_state: uint16(8),
	},
	86: {
		Fexternal_lex_state: uint16(2),
	},
	87: {
		Fexternal_lex_state: uint16(2),
	},
	88: {
		Fexternal_lex_state: uint16(9),
	},
	89: {
		Fexternal_lex_state: uint16(8),
	},
	90: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(2),
	},
	91: {
		Fexternal_lex_state: uint16(2),
	},
	92: {
		Fexternal_lex_state: uint16(9),
	},
	93: {
		Fexternal_lex_state: uint16(2),
	},
}

var ts_parse_table = [2][41]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		3:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		6:  uint16(1),
		7:  uint16(1),
		8:  uint16(1),
		11: uint16(1),
		12: uint16(1),
		14: uint16(1),
		17: uint16(1),
		18: uint16(1),
		19: uint16(1),
		20: uint16(1),
		21: uint16(1),
		22: uint16(1),
		23: uint16(1),
		24: uint16(3),
	},
	1: {
		0:  uint16(5),
		1:  uint16(7),
		5:  uint16(9),
		7:  uint16(11),
		11: uint16(13),
		16: uint16(13),
		24: uint16(3),
		25: uint16(83),
		26: uint16(7),
		27: uint16(7),
		28: uint16(7),
		29: uint16(7),
		30: uint16(7),
		31: uint16(2),
		32: uint16(48),
		33: uint16(53),
		34: uint16(33),
		36: uint16(7),
		39: uint16(7),
	},
}

var ts_small_parse_table = [1265]uint16_t{
	0:    uint16(12),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(15),
	5:    uint16(1),
	6:    uint16(anon_sym_LT_BANG),
	7:    uint16(17),
	8:    uint16(1),
	9:    uint16(anon_sym_LT),
	10:   uint16(19),
	11:   uint16(1),
	12:   uint16(anon_sym_LT_SLASH),
	13:   uint16(23),
	14:   uint16(1),
	15:   uint16(sym__implicit_end_tag),
	16:   uint16(5),
	17:   uint16(1),
	18:   uint16(sym_start_tag),
	19:   uint16(21),
	20:   uint16(1),
	21:   uint16(sym_self_closing_tag),
	22:   uint16(32),
	23:   uint16(1),
	24:   uint16(sym_end_tag),
	25:   uint16(54),
	26:   uint16(1),
	27:   uint16(sym_script_start_tag),
	28:   uint16(55),
	29:   uint16(1),
	30:   uint16(sym_style_start_tag),
	31:   uint16(21),
	32:   uint16(2),
	33:   uint16(sym_entity),
	34:   uint16(sym_text),
	35:   uint16(3),
	36:   uint16(7),
	37:   uint16(sym_doctype),
	38:   uint16(sym__node),
	39:   uint16(sym_element),
	40:   uint16(sym_script_element),
	41:   uint16(sym_style_element),
	42:   uint16(sym_erroneous_end_tag),
	43:   uint16(aux_sym_document_repeat1),
	44:   uint16(12),
	45:   uint16(3),
	46:   uint16(1),
	47:   uint16(sym_comment),
	48:   uint16(15),
	49:   uint16(1),
	50:   uint16(anon_sym_LT_BANG),
	51:   uint16(17),
	52:   uint16(1),
	53:   uint16(anon_sym_LT),
	54:   uint16(19),
	55:   uint16(1),
	56:   uint16(anon_sym_LT_SLASH),
	57:   uint16(27),
	58:   uint16(1),
	59:   uint16(sym__implicit_end_tag),
	60:   uint16(5),
	61:   uint16(1),
	62:   uint16(sym_start_tag),
	63:   uint16(14),
	64:   uint16(1),
	65:   uint16(sym_end_tag),
	66:   uint16(21),
	67:   uint16(1),
	68:   uint16(sym_self_closing_tag),
	69:   uint16(54),
	70:   uint16(1),
	71:   uint16(sym_script_start_tag),
	72:   uint16(55),
	73:   uint16(1),
	74:   uint16(sym_style_start_tag),
	75:   uint16(25),
	76:   uint16(2),
	77:   uint16(sym_entity),
	78:   uint16(sym_text),
	79:   uint16(6),
	80:   uint16(7),
	81:   uint16(sym_doctype),
	82:   uint16(sym__node),
	83:   uint16(sym_element),
	84:   uint16(sym_script_element),
	85:   uint16(sym_style_element),
	86:   uint16(sym_erroneous_end_tag),
	87:   uint16(aux_sym_document_repeat1),
	88:   uint16(12),
	89:   uint16(3),
	90:   uint16(1),
	91:   uint16(sym_comment),
	92:   uint16(15),
	93:   uint16(1),
	94:   uint16(anon_sym_LT_BANG),
	95:   uint16(17),
	96:   uint16(1),
	97:   uint16(anon_sym_LT),
	98:   uint16(29),
	99:   uint16(1),
	100:  uint16(anon_sym_LT_SLASH),
	101:  uint16(31),
	102:  uint16(1),
	103:  uint16(sym__implicit_end_tag),
	104:  uint16(5),
	105:  uint16(1),
	106:  uint16(sym_start_tag),
	107:  uint16(21),
	108:  uint16(1),
	109:  uint16(sym_self_closing_tag),
	110:  uint16(27),
	111:  uint16(1),
	112:  uint16(sym_end_tag),
	113:  uint16(54),
	114:  uint16(1),
	115:  uint16(sym_script_start_tag),
	116:  uint16(55),
	117:  uint16(1),
	118:  uint16(sym_style_start_tag),
	119:  uint16(25),
	120:  uint16(2),
	121:  uint16(sym_entity),
	122:  uint16(sym_text),
	123:  uint16(6),
	124:  uint16(7),
	125:  uint16(sym_doctype),
	126:  uint16(sym__node),
	127:  uint16(sym_element),
	128:  uint16(sym_script_element),
	129:  uint16(sym_style_element),
	130:  uint16(sym_erroneous_end_tag),
	131:  uint16(aux_sym_document_repeat1),
	132:  uint16(12),
	133:  uint16(3),
	134:  uint16(1),
	135:  uint16(sym_comment),
	136:  uint16(15),
	137:  uint16(1),
	138:  uint16(anon_sym_LT_BANG),
	139:  uint16(17),
	140:  uint16(1),
	141:  uint16(anon_sym_LT),
	142:  uint16(29),
	143:  uint16(1),
	144:  uint16(anon_sym_LT_SLASH),
	145:  uint16(35),
	146:  uint16(1),
	147:  uint16(sym__implicit_end_tag),
	148:  uint16(5),
	149:  uint16(1),
	150:  uint16(sym_start_tag),
	151:  uint16(21),
	152:  uint16(1),
	153:  uint16(sym_self_closing_tag),
	154:  uint16(22),
	155:  uint16(1),
	156:  uint16(sym_end_tag),
	157:  uint16(54),
	158:  uint16(1),
	159:  uint16(sym_script_start_tag),
	160:  uint16(55),
	161:  uint16(1),
	162:  uint16(sym_style_start_tag),
	163:  uint16(33),
	164:  uint16(2),
	165:  uint16(sym_entity),
	166:  uint16(sym_text),
	167:  uint16(4),
	168:  uint16(7),
	169:  uint16(sym_doctype),
	170:  uint16(sym__node),
	171:  uint16(sym_element),
	172:  uint16(sym_script_element),
	173:  uint16(sym_style_element),
	174:  uint16(sym_erroneous_end_tag),
	175:  uint16(aux_sym_document_repeat1),
	176:  uint16(11),
	177:  uint16(3),
	178:  uint16(1),
	179:  uint16(sym_comment),
	180:  uint16(37),
	181:  uint16(1),
	182:  uint16(anon_sym_LT_BANG),
	183:  uint16(40),
	184:  uint16(1),
	185:  uint16(anon_sym_LT),
	186:  uint16(43),
	187:  uint16(1),
	188:  uint16(anon_sym_LT_SLASH),
	189:  uint16(49),
	190:  uint16(1),
	191:  uint16(sym__implicit_end_tag),
	192:  uint16(5),
	193:  uint16(1),
	194:  uint16(sym_start_tag),
	195:  uint16(21),
	196:  uint16(1),
	197:  uint16(sym_self_closing_tag),
	198:  uint16(54),
	199:  uint16(1),
	200:  uint16(sym_script_start_tag),
	201:  uint16(55),
	202:  uint16(1),
	203:  uint16(sym_style_start_tag),
	204:  uint16(46),
	205:  uint16(2),
	206:  uint16(sym_entity),
	207:  uint16(sym_text),
	208:  uint16(6),
	209:  uint16(7),
	210:  uint16(sym_doctype),
	211:  uint16(sym__node),
	212:  uint16(sym_element),
	213:  uint16(sym_script_element),
	214:  uint16(sym_style_element),
	215:  uint16(sym_erroneous_end_tag),
	216:  uint16(aux_sym_document_repeat1),
	217:  uint16(11),
	218:  uint16(3),
	219:  uint16(1),
	220:  uint16(sym_comment),
	221:  uint16(7),
	222:  uint16(1),
	223:  uint16(anon_sym_LT_BANG),
	224:  uint16(9),
	225:  uint16(1),
	226:  uint16(anon_sym_LT),
	227:  uint16(11),
	228:  uint16(1),
	229:  uint16(anon_sym_LT_SLASH),
	230:  uint16(51),
	231:  uint16(1),
	233:  uint16(2),
	234:  uint16(1),
	235:  uint16(sym_start_tag),
	236:  uint16(33),
	237:  uint16(1),
	238:  uint16(sym_self_closing_tag),
	239:  uint16(48),
	240:  uint16(1),
	241:  uint16(sym_script_start_tag),
	242:  uint16(53),
	243:  uint16(1),
	244:  uint16(sym_style_start_tag),
	245:  uint16(53),
	246:  uint16(2),
	247:  uint16(sym_entity),
	248:  uint16(sym_text),
	249:  uint16(8),
	250:  uint16(7),
	251:  uint16(sym_doctype),
	252:  uint16(sym__node),
	253:  uint16(sym_element),
	254:  uint16(sym_script_element),
	255:  uint16(sym_style_element),
	256:  uint16(sym_erroneous_end_tag),
	257:  uint16(aux_sym_document_repeat1),
	258:  uint16(11),
	259:  uint16(3),
	260:  uint16(1),
	261:  uint16(sym_comment),
	262:  uint16(49),
	263:  uint16(1),
	265:  uint16(55),
	266:  uint16(1),
	267:  uint16(anon_sym_LT_BANG),
	268:  uint16(58),
	269:  uint16(1),
	270:  uint16(anon_sym_LT),
	271:  uint16(61),
	272:  uint16(1),
	273:  uint16(anon_sym_LT_SLASH),
	274:  uint16(2),
	275:  uint16(1),
	276:  uint16(sym_start_tag),
	277:  uint16(33),
	278:  uint16(1),
	279:  uint16(sym_self_closing_tag),
	280:  uint16(48),
	281:  uint16(1),
	282:  uint16(sym_script_start_tag),
	283:  uint16(53),
	284:  uint16(1),
	285:  uint16(sym_style_start_tag),
	286:  uint16(64),
	287:  uint16(2),
	288:  uint16(sym_entity),
	289:  uint16(sym_text),
	290:  uint16(8),
	291:  uint16(7),
	292:  uint16(sym_doctype),
	293:  uint16(sym__node),
	294:  uint16(sym_element),
	295:  uint16(sym_script_element),
	296:  uint16(sym_style_element),
	297:  uint16(sym_erroneous_end_tag),
	298:  uint16(aux_sym_document_repeat1),
	299:  uint16(3),
	300:  uint16(3),
	301:  uint16(1),
	302:  uint16(sym_comment),
	303:  uint16(69),
	304:  uint16(1),
	305:  uint16(anon_sym_LT),
	306:  uint16(67),
	307:  uint16(5),
	308:  uint16(sym__implicit_end_tag),
	309:  uint16(anon_sym_LT_BANG),
	310:  uint16(anon_sym_LT_SLASH),
	311:  uint16(sym_entity),
	312:  uint16(sym_text),
	313:  uint16(3),
	314:  uint16(3),
	315:  uint16(1),
	316:  uint16(sym_comment),
	317:  uint16(73),
	318:  uint16(1),
	319:  uint16(anon_sym_LT),
	320:  uint16(71),
	321:  uint16(5),
	323:  uint16(anon_sym_LT_BANG),
	324:  uint16(anon_sym_LT_SLASH),
	325:  uint16(sym_entity),
	326:  uint16(sym_text),
	327:  uint16(3),
	328:  uint16(3),
	329:  uint16(1),
	330:  uint16(sym_comment),
	331:  uint16(77),
	332:  uint16(1),
	333:  uint16(anon_sym_LT),
	334:  uint16(75),
	335:  uint16(5),
	336:  uint16(sym__implicit_end_tag),
	337:  uint16(anon_sym_LT_BANG),
	338:  uint16(anon_sym_LT_SLASH),
	339:  uint16(sym_entity),
	340:  uint16(sym_text),
	341:  uint16(3),
	342:  uint16(3),
	343:  uint16(1),
	344:  uint16(sym_comment),
	345:  uint16(81),
	346:  uint16(1),
	347:  uint16(anon_sym_LT),
	348:  uint16(79),
	349:  uint16(5),
	351:  uint16(anon_sym_LT_BANG),
	352:  uint16(anon_sym_LT_SLASH),
	353:  uint16(sym_entity),
	354:  uint16(sym_text),
	355:  uint16(3),
	356:  uint16(3),
	357:  uint16(1),
	358:  uint16(sym_comment),
	359:  uint16(85),
	360:  uint16(1),
	361:  uint16(anon_sym_LT),
	362:  uint16(83),
	363:  uint16(5),
	365:  uint16(anon_sym_LT_BANG),
	366:  uint16(anon_sym_LT_SLASH),
	367:  uint16(sym_entity),
	368:  uint16(sym_text),
	369:  uint16(3),
	370:  uint16(3),
	371:  uint16(1),
	372:  uint16(sym_comment),
	373:  uint16(89),
	374:  uint16(1),
	375:  uint16(anon_sym_LT),
	376:  uint16(87),
	377:  uint16(5),
	379:  uint16(anon_sym_LT_BANG),
	380:  uint16(anon_sym_LT_SLASH),
	381:  uint16(sym_entity),
	382:  uint16(sym_text),
	383:  uint16(3),
	384:  uint16(3),
	385:  uint16(1),
	386:  uint16(sym_comment),
	387:  uint16(93),
	388:  uint16(1),
	389:  uint16(anon_sym_LT),
	390:  uint16(91),
	391:  uint16(5),
	393:  uint16(anon_sym_LT_BANG),
	394:  uint16(anon_sym_LT_SLASH),
	395:  uint16(sym_entity),
	396:  uint16(sym_text),
	397:  uint16(3),
	398:  uint16(3),
	399:  uint16(1),
	400:  uint16(sym_comment),
	401:  uint16(97),
	402:  uint16(1),
	403:  uint16(anon_sym_LT),
	404:  uint16(95),
	405:  uint16(5),
	407:  uint16(anon_sym_LT_BANG),
	408:  uint16(anon_sym_LT_SLASH),
	409:  uint16(sym_entity),
	410:  uint16(sym_text),
	411:  uint16(3),
	412:  uint16(3),
	413:  uint16(1),
	414:  uint16(sym_comment),
	415:  uint16(101),
	416:  uint16(1),
	417:  uint16(anon_sym_LT),
	418:  uint16(99),
	419:  uint16(5),
	421:  uint16(anon_sym_LT_BANG),
	422:  uint16(anon_sym_LT_SLASH),
	423:  uint16(sym_entity),
	424:  uint16(sym_text),
	425:  uint16(3),
	426:  uint16(3),
	427:  uint16(1),
	428:  uint16(sym_comment),
	429:  uint16(105),
	430:  uint16(1),
	431:  uint16(anon_sym_LT),
	432:  uint16(103),
	433:  uint16(5),
	434:  uint16(sym__implicit_end_tag),
	435:  uint16(anon_sym_LT_BANG),
	436:  uint16(anon_sym_LT_SLASH),
	437:  uint16(sym_entity),
	438:  uint16(sym_text),
	439:  uint16(3),
	440:  uint16(3),
	441:  uint16(1),
	442:  uint16(sym_comment),
	443:  uint16(109),
	444:  uint16(1),
	445:  uint16(anon_sym_LT),
	446:  uint16(107),
	447:  uint16(5),
	449:  uint16(anon_sym_LT_BANG),
	450:  uint16(anon_sym_LT_SLASH),
	451:  uint16(sym_entity),
	452:  uint16(sym_text),
	453:  uint16(3),
	454:  uint16(3),
	455:  uint16(1),
	456:  uint16(sym_comment),
	457:  uint16(69),
	458:  uint16(1),
	459:  uint16(anon_sym_LT),
	460:  uint16(67),
	461:  uint16(5),
	463:  uint16(anon_sym_LT_BANG),
	464:  uint16(anon_sym_LT_SLASH),
	465:  uint16(sym_entity),
	466:  uint16(sym_text),
	467:  uint16(3),
	468:  uint16(3),
	469:  uint16(1),
	470:  uint16(sym_comment),
	471:  uint16(113),
	472:  uint16(1),
	473:  uint16(anon_sym_LT),
	474:  uint16(111),
	475:  uint16(5),
	476:  uint16(sym__implicit_end_tag),
	477:  uint16(anon_sym_LT_BANG),
	478:  uint16(anon_sym_LT_SLASH),
	479:  uint16(sym_entity),
	480:  uint16(sym_text),
	481:  uint16(3),
	482:  uint16(3),
	483:  uint16(1),
	484:  uint16(sym_comment),
	485:  uint16(117),
	486:  uint16(1),
	487:  uint16(anon_sym_LT),
	488:  uint16(115),
	489:  uint16(5),
	490:  uint16(sym__implicit_end_tag),
	491:  uint16(anon_sym_LT_BANG),
	492:  uint16(anon_sym_LT_SLASH),
	493:  uint16(sym_entity),
	494:  uint16(sym_text),
	495:  uint16(3),
	496:  uint16(3),
	497:  uint16(1),
	498:  uint16(sym_comment),
	499:  uint16(109),
	500:  uint16(1),
	501:  uint16(anon_sym_LT),
	502:  uint16(107),
	503:  uint16(5),
	504:  uint16(sym__implicit_end_tag),
	505:  uint16(anon_sym_LT_BANG),
	506:  uint16(anon_sym_LT_SLASH),
	507:  uint16(sym_entity),
	508:  uint16(sym_text),
	509:  uint16(3),
	510:  uint16(3),
	511:  uint16(1),
	512:  uint16(sym_comment),
	513:  uint16(73),
	514:  uint16(1),
	515:  uint16(anon_sym_LT),
	516:  uint16(71),
	517:  uint16(5),
	518:  uint16(sym__implicit_end_tag),
	519:  uint16(anon_sym_LT_BANG),
	520:  uint16(anon_sym_LT_SLASH),
	521:  uint16(sym_entity),
	522:  uint16(sym_text),
	523:  uint16(3),
	524:  uint16(3),
	525:  uint16(1),
	526:  uint16(sym_comment),
	527:  uint16(81),
	528:  uint16(1),
	529:  uint16(anon_sym_LT),
	530:  uint16(79),
	531:  uint16(5),
	532:  uint16(sym__implicit_end_tag),
	533:  uint16(anon_sym_LT_BANG),
	534:  uint16(anon_sym_LT_SLASH),
	535:  uint16(sym_entity),
	536:  uint16(sym_text),
	537:  uint16(3),
	538:  uint16(3),
	539:  uint16(1),
	540:  uint16(sym_comment),
	541:  uint16(85),
	542:  uint16(1),
	543:  uint16(anon_sym_LT),
	544:  uint16(83),
	545:  uint16(5),
	546:  uint16(sym__implicit_end_tag),
	547:  uint16(anon_sym_LT_BANG),
	548:  uint16(anon_sym_LT_SLASH),
	549:  uint16(sym_entity),
	550:  uint16(sym_text),
	551:  uint16(3),
	552:  uint16(3),
	553:  uint16(1),
	554:  uint16(sym_comment),
	555:  uint16(89),
	556:  uint16(1),
	557:  uint16(anon_sym_LT),
	558:  uint16(87),
	559:  uint16(5),
	560:  uint16(sym__implicit_end_tag),
	561:  uint16(anon_sym_LT_BANG),
	562:  uint16(anon_sym_LT_SLASH),
	563:  uint16(sym_entity),
	564:  uint16(sym_text),
	565:  uint16(3),
	566:  uint16(3),
	567:  uint16(1),
	568:  uint16(sym_comment),
	569:  uint16(93),
	570:  uint16(1),
	571:  uint16(anon_sym_LT),
	572:  uint16(91),
	573:  uint16(5),
	574:  uint16(sym__implicit_end_tag),
	575:  uint16(anon_sym_LT_BANG),
	576:  uint16(anon_sym_LT_SLASH),
	577:  uint16(sym_entity),
	578:  uint16(sym_text),
	579:  uint16(3),
	580:  uint16(3),
	581:  uint16(1),
	582:  uint16(sym_comment),
	583:  uint16(97),
	584:  uint16(1),
	585:  uint16(anon_sym_LT),
	586:  uint16(95),
	587:  uint16(5),
	588:  uint16(sym__implicit_end_tag),
	589:  uint16(anon_sym_LT_BANG),
	590:  uint16(anon_sym_LT_SLASH),
	591:  uint16(sym_entity),
	592:  uint16(sym_text),
	593:  uint16(3),
	594:  uint16(3),
	595:  uint16(1),
	596:  uint16(sym_comment),
	597:  uint16(101),
	598:  uint16(1),
	599:  uint16(anon_sym_LT),
	600:  uint16(99),
	601:  uint16(5),
	602:  uint16(sym__implicit_end_tag),
	603:  uint16(anon_sym_LT_BANG),
	604:  uint16(anon_sym_LT_SLASH),
	605:  uint16(sym_entity),
	606:  uint16(sym_text),
	607:  uint16(3),
	608:  uint16(3),
	609:  uint16(1),
	610:  uint16(sym_comment),
	611:  uint16(121),
	612:  uint16(1),
	613:  uint16(anon_sym_LT),
	614:  uint16(119),
	615:  uint16(5),
	616:  uint16(sym__implicit_end_tag),
	617:  uint16(anon_sym_LT_BANG),
	618:  uint16(anon_sym_LT_SLASH),
	619:  uint16(sym_entity),
	620:  uint16(sym_text),
	621:  uint16(3),
	622:  uint16(3),
	623:  uint16(1),
	624:  uint16(sym_comment),
	625:  uint16(117),
	626:  uint16(1),
	627:  uint16(anon_sym_LT),
	628:  uint16(115),
	629:  uint16(5),
	631:  uint16(anon_sym_LT_BANG),
	632:  uint16(anon_sym_LT_SLASH),
	633:  uint16(sym_entity),
	634:  uint16(sym_text),
	635:  uint16(3),
	636:  uint16(3),
	637:  uint16(1),
	638:  uint16(sym_comment),
	639:  uint16(113),
	640:  uint16(1),
	641:  uint16(anon_sym_LT),
	642:  uint16(111),
	643:  uint16(5),
	645:  uint16(anon_sym_LT_BANG),
	646:  uint16(anon_sym_LT_SLASH),
	647:  uint16(sym_entity),
	648:  uint16(sym_text),
	649:  uint16(3),
	650:  uint16(3),
	651:  uint16(1),
	652:  uint16(sym_comment),
	653:  uint16(121),
	654:  uint16(1),
	655:  uint16(anon_sym_LT),
	656:  uint16(119),
	657:  uint16(5),
	659:  uint16(anon_sym_LT_BANG),
	660:  uint16(anon_sym_LT_SLASH),
	661:  uint16(sym_entity),
	662:  uint16(sym_text),
	663:  uint16(4),
	664:  uint16(3),
	665:  uint16(1),
	666:  uint16(sym_comment),
	667:  uint16(125),
	668:  uint16(1),
	669:  uint16(sym_attribute_name),
	670:  uint16(123),
	671:  uint16(2),
	672:  uint16(anon_sym_GT),
	673:  uint16(anon_sym_SLASH_GT),
	674:  uint16(35),
	675:  uint16(2),
	676:  uint16(sym_attribute),
	677:  uint16(aux_sym_start_tag_repeat1),
	678:  uint16(5),
	679:  uint16(3),
	680:  uint16(1),
	681:  uint16(sym_comment),
	682:  uint16(128),
	683:  uint16(1),
	684:  uint16(anon_sym_GT),
	685:  uint16(130),
	686:  uint16(1),
	687:  uint16(anon_sym_SLASH_GT),
	688:  uint16(132),
	689:  uint16(1),
	690:  uint16(sym_attribute_name),
	691:  uint16(35),
	692:  uint16(2),
	693:  uint16(sym_attribute),
	694:  uint16(aux_sym_start_tag_repeat1),
	695:  uint16(5),
	696:  uint16(3),
	697:  uint16(1),
	698:  uint16(sym_comment),
	699:  uint16(132),
	700:  uint16(1),
	701:  uint16(sym_attribute_name),
	702:  uint16(134),
	703:  uint16(1),
	704:  uint16(anon_sym_GT),
	705:  uint16(136),
	706:  uint16(1),
	707:  uint16(anon_sym_SLASH_GT),
	708:  uint16(38),
	709:  uint16(2),
	710:  uint16(sym_attribute),
	711:  uint16(aux_sym_start_tag_repeat1),
	712:  uint16(5),
	713:  uint16(3),
	714:  uint16(1),
	715:  uint16(sym_comment),
	716:  uint16(128),
	717:  uint16(1),
	718:  uint16(anon_sym_GT),
	719:  uint16(132),
	720:  uint16(1),
	721:  uint16(sym_attribute_name),
	722:  uint16(138),
	723:  uint16(1),
	724:  uint16(anon_sym_SLASH_GT),
	725:  uint16(35),
	726:  uint16(2),
	727:  uint16(sym_attribute),
	728:  uint16(aux_sym_start_tag_repeat1),
	729:  uint16(5),
	730:  uint16(3),
	731:  uint16(1),
	732:  uint16(sym_comment),
	733:  uint16(132),
	734:  uint16(1),
	735:  uint16(sym_attribute_name),
	736:  uint16(134),
	737:  uint16(1),
	738:  uint16(anon_sym_GT),
	739:  uint16(140),
	740:  uint16(1),
	741:  uint16(anon_sym_SLASH_GT),
	742:  uint16(36),
	743:  uint16(2),
	744:  uint16(sym_attribute),
	745:  uint16(aux_sym_start_tag_repeat1),
	746:  uint16(4),
	747:  uint16(3),
	748:  uint16(1),
	749:  uint16(sym_comment),
	750:  uint16(142),
	751:  uint16(1),
	752:  uint16(anon_sym_GT),
	753:  uint16(144),
	754:  uint16(1),
	755:  uint16(sym_attribute_name),
	756:  uint16(41),
	757:  uint16(2),
	758:  uint16(sym_attribute),
	759:  uint16(aux_sym_start_tag_repeat1),
	760:  uint16(4),
	761:  uint16(3),
	762:  uint16(1),
	763:  uint16(sym_comment),
	764:  uint16(123),
	765:  uint16(1),
	766:  uint16(anon_sym_GT),
	767:  uint16(146),
	768:  uint16(1),
	769:  uint16(sym_attribute_name),
	770:  uint16(41),
	771:  uint16(2),
	772:  uint16(sym_attribute),
	773:  uint16(aux_sym_start_tag_repeat1),
	774:  uint16(4),
	775:  uint16(3),
	776:  uint16(1),
	777:  uint16(sym_comment),
	778:  uint16(144),
	779:  uint16(1),
	780:  uint16(sym_attribute_name),
	781:  uint16(149),
	782:  uint16(1),
	783:  uint16(anon_sym_GT),
	784:  uint16(40),
	785:  uint16(2),
	786:  uint16(sym_attribute),
	787:  uint16(aux_sym_start_tag_repeat1),
	788:  uint16(4),
	789:  uint16(3),
	790:  uint16(1),
	791:  uint16(sym_comment),
	792:  uint16(144),
	793:  uint16(1),
	794:  uint16(sym_attribute_name),
	795:  uint16(151),
	796:  uint16(1),
	797:  uint16(anon_sym_GT),
	798:  uint16(44),
	799:  uint16(2),
	800:  uint16(sym_attribute),
	801:  uint16(aux_sym_start_tag_repeat1),
	802:  uint16(4),
	803:  uint16(3),
	804:  uint16(1),
	805:  uint16(sym_comment),
	806:  uint16(144),
	807:  uint16(1),
	808:  uint16(sym_attribute_name),
	809:  uint16(153),
	810:  uint16(1),
	811:  uint16(anon_sym_GT),
	812:  uint16(41),
	813:  uint16(2),
	814:  uint16(sym_attribute),
	815:  uint16(aux_sym_start_tag_repeat1),
	816:  uint16(3),
	817:  uint16(3),
	818:  uint16(1),
	819:  uint16(sym_comment),
	820:  uint16(157),
	821:  uint16(1),
	822:  uint16(anon_sym_EQ),
	823:  uint16(155),
	824:  uint16(3),
	825:  uint16(anon_sym_GT),
	826:  uint16(anon_sym_SLASH_GT),
	827:  uint16(sym_attribute_name),
	828:  uint16(5),
	829:  uint16(3),
	830:  uint16(1),
	831:  uint16(sym_comment),
	832:  uint16(159),
	833:  uint16(1),
	834:  uint16(sym_attribute_value),
	835:  uint16(161),
	836:  uint16(1),
	837:  uint16(anon_sym_SQUOTE),
	838:  uint16(163),
	839:  uint16(1),
	840:  uint16(anon_sym_DQUOTE),
	841:  uint16(56),
	842:  uint16(1),
	843:  uint16(sym_quoted_attribute_value),
	844:  uint16(5),
	845:  uint16(3),
	846:  uint16(1),
	847:  uint16(sym_comment),
	848:  uint16(165),
	849:  uint16(1),
	850:  uint16(sym_attribute_value),
	851:  uint16(167),
	852:  uint16(1),
	853:  uint16(anon_sym_SQUOTE),
	854:  uint16(169),
	855:  uint16(1),
	856:  uint16(anon_sym_DQUOTE),
	857:  uint16(58),
	858:  uint16(1),
	859:  uint16(sym_quoted_attribute_value),
	860:  uint16(4),
	861:  uint16(3),
	862:  uint16(1),
	863:  uint16(sym_comment),
	864:  uint16(171),
	865:  uint16(1),
	866:  uint16(anon_sym_LT_SLASH),
	867:  uint16(173),
	868:  uint16(1),
	869:  uint16(sym_raw_text),
	870:  uint16(19),
	871:  uint16(1),
	872:  uint16(sym_end_tag),
	873:  uint16(2),
	874:  uint16(3),
	875:  uint16(1),
	876:  uint16(sym_comment),
	877:  uint16(175),
	878:  uint16(3),
	879:  uint16(anon_sym_GT),
	880:  uint16(anon_sym_SLASH_GT),
	881:  uint16(sym_attribute_name),
	882:  uint16(2),
	883:  uint16(3),
	884:  uint16(1),
	885:  uint16(sym_comment),
	886:  uint16(177),
	887:  uint16(3),
	888:  uint16(anon_sym_GT),
	889:  uint16(anon_sym_SLASH_GT),
	890:  uint16(sym_attribute_name),
	891:  uint16(3),
	892:  uint16(3),
	893:  uint16(1),
	894:  uint16(sym_comment),
	895:  uint16(179),
	896:  uint16(1),
	897:  uint16(anon_sym_EQ),
	898:  uint16(155),
	899:  uint16(2),
	900:  uint16(anon_sym_GT),
	901:  uint16(sym_attribute_name),
	902:  uint16(4),
	903:  uint16(3),
	904:  uint16(1),
	905:  uint16(sym_comment),
	906:  uint16(181),
	907:  uint16(1),
	908:  uint16(sym__start_tag_name),
	909:  uint16(183),
	910:  uint16(1),
	911:  uint16(sym__script_start_tag_name),
	912:  uint16(185),
	913:  uint16(1),
	914:  uint16(sym__style_start_tag_name),
	915:  uint16(4),
	916:  uint16(3),
	917:  uint16(1),
	918:  uint16(sym_comment),
	919:  uint16(171),
	920:  uint16(1),
	921:  uint16(anon_sym_LT_SLASH),
	922:  uint16(187),
	923:  uint16(1),
	924:  uint16(sym_raw_text),
	925:  uint16(10),
	926:  uint16(1),
	927:  uint16(sym_end_tag),
	928:  uint16(4),
	929:  uint16(3),
	930:  uint16(1),
	931:  uint16(sym_comment),
	932:  uint16(189),
	933:  uint16(1),
	934:  uint16(anon_sym_LT_SLASH),
	935:  uint16(191),
	936:  uint16(1),
	937:  uint16(sym_raw_text),
	938:  uint16(23),
	939:  uint16(1),
	940:  uint16(sym_end_tag),
	941:  uint16(4),
	942:  uint16(3),
	943:  uint16(1),
	944:  uint16(sym_comment),
	945:  uint16(189),
	946:  uint16(1),
	947:  uint16(anon_sym_LT_SLASH),
	948:  uint16(193),
	949:  uint16(1),
	950:  uint16(sym_raw_text),
	951:  uint16(24),
	952:  uint16(1),
	953:  uint16(sym_end_tag),
	954:  uint16(2),
	955:  uint16(3),
	956:  uint16(1),
	957:  uint16(sym_comment),
	958:  uint16(195),
	959:  uint16(3),
	960:  uint16(anon_sym_GT),
	961:  uint16(anon_sym_SLASH_GT),
	962:  uint16(sym_attribute_name),
	963:  uint16(4),
	964:  uint16(3),
	965:  uint16(1),
	966:  uint16(sym_comment),
	967:  uint16(183),
	968:  uint16(1),
	969:  uint16(sym__script_start_tag_name),
	970:  uint16(185),
	971:  uint16(1),
	972:  uint16(sym__style_start_tag_name),
	973:  uint16(197),
	974:  uint16(1),
	975:  uint16(sym__start_tag_name),
	976:  uint16(2),
	977:  uint16(3),
	978:  uint16(1),
	979:  uint16(sym_comment),
	980:  uint16(195),
	981:  uint16(2),
	982:  uint16(anon_sym_GT),
	983:  uint16(sym_attribute_name),
	984:  uint16(3),
	985:  uint16(3),
	986:  uint16(1),
	987:  uint16(sym_comment),
	988:  uint16(199),
	989:  uint16(1),
	990:  uint16(sym__end_tag_name),
	991:  uint16(201),
	992:  uint16(1),
	993:  uint16(sym_erroneous_end_tag_name),
	994:  uint16(3),
	995:  uint16(3),
	996:  uint16(1),
	997:  uint16(sym_comment),
	998:  uint16(171),
	999:  uint16(1),
	1000: uint16(anon_sym_LT_SLASH),
	1001: uint16(16),
	1002: uint16(1),
	1003: uint16(sym_end_tag),
	1004: uint16(2),
	1005: uint16(3),
	1006: uint16(1),
	1007: uint16(sym_comment),
	1008: uint16(203),
	1009: uint16(2),
	1010: uint16(sym_raw_text),
	1011: uint16(anon_sym_LT_SLASH),
	1012: uint16(3),
	1013: uint16(3),
	1014: uint16(1),
	1015: uint16(sym_comment),
	1016: uint16(205),
	1017: uint16(1),
	1018: uint16(anon_sym_DQUOTE),
	1019: uint16(207),
	1020: uint16(1),
	1021: uint16(aux_sym_quoted_attribute_value_token2),
	1022: uint16(2),
	1023: uint16(3),
	1024: uint16(1),
	1025: uint16(sym_comment),
	1026: uint16(175),
	1027: uint16(2),
	1028: uint16(anon_sym_GT),
	1029: uint16(sym_attribute_name),
	1030: uint16(2),
	1031: uint16(3),
	1032: uint16(1),
	1033: uint16(sym_comment),
	1034: uint16(177),
	1035: uint16(2),
	1036: uint16(anon_sym_GT),
	1037: uint16(sym_attribute_name),
	1038: uint16(2),
	1039: uint16(3),
	1040: uint16(1),
	1041: uint16(sym_comment),
	1042: uint16(209),
	1043: uint16(2),
	1044: uint16(sym_raw_text),
	1045: uint16(anon_sym_LT_SLASH),
	1046: uint16(2),
	1047: uint16(3),
	1048: uint16(1),
	1049: uint16(sym_comment),
	1050: uint16(211),
	1051: uint16(2),
	1052: uint16(sym_raw_text),
	1053: uint16(anon_sym_LT_SLASH),
	1054: uint16(2),
	1055: uint16(3),
	1056: uint16(1),
	1057: uint16(sym_comment),
	1058: uint16(213),
	1059: uint16(2),
	1060: uint16(sym_raw_text),
	1061: uint16(anon_sym_LT_SLASH),
	1062: uint16(3),
	1063: uint16(3),
	1064: uint16(1),
	1065: uint16(sym_comment),
	1066: uint16(201),
	1067: uint16(1),
	1068: uint16(sym_erroneous_end_tag_name),
	1069: uint16(215),
	1070: uint16(1),
	1071: uint16(sym__end_tag_name),
	1072: uint16(3),
	1073: uint16(3),
	1074: uint16(1),
	1075: uint16(sym_comment),
	1076: uint16(189),
	1077: uint16(1),
	1078: uint16(anon_sym_LT_SLASH),
	1079: uint16(28),
	1080: uint16(1),
	1081: uint16(sym_end_tag),
	1082: uint16(3),
	1083: uint16(3),
	1084: uint16(1),
	1085: uint16(sym_comment),
	1086: uint16(189),
	1087: uint16(1),
	1088: uint16(anon_sym_LT_SLASH),
	1089: uint16(29),
	1090: uint16(1),
	1091: uint16(sym_end_tag),
	1092: uint16(3),
	1093: uint16(3),
	1094: uint16(1),
	1095: uint16(sym_comment),
	1096: uint16(171),
	1097: uint16(1),
	1098: uint16(anon_sym_LT_SLASH),
	1099: uint16(15),
	1100: uint16(1),
	1101: uint16(sym_end_tag),
	1102: uint16(3),
	1103: uint16(3),
	1104: uint16(1),
	1105: uint16(sym_comment),
	1106: uint16(205),
	1107: uint16(1),
	1108: uint16(anon_sym_SQUOTE),
	1109: uint16(217),
	1110: uint16(1),
	1111: uint16(aux_sym_quoted_attribute_value_token1),
	1112: uint16(3),
	1113: uint16(3),
	1114: uint16(1),
	1115: uint16(sym_comment),
	1116: uint16(219),
	1117: uint16(1),
	1118: uint16(anon_sym_SQUOTE),
	1119: uint16(221),
	1120: uint16(1),
	1121: uint16(aux_sym_quoted_attribute_value_token1),
	1122: uint16(3),
	1123: uint16(3),
	1124: uint16(1),
	1125: uint16(sym_comment),
	1126: uint16(219),
	1127: uint16(1),
	1128: uint16(anon_sym_DQUOTE),
	1129: uint16(223),
	1130: uint16(1),
	1131: uint16(aux_sym_quoted_attribute_value_token2),
	1132: uint16(2),
	1133: uint16(3),
	1134: uint16(1),
	1135: uint16(sym_comment),
	1136: uint16(225),
	1137: uint16(1),
	1138: uint16(anon_sym_DQUOTE),
	1139: uint16(2),
	1140: uint16(3),
	1141: uint16(1),
	1142: uint16(sym_comment),
	1143: uint16(227),
	1144: uint16(1),
	1145: uint16(anon_sym_DQUOTE),
	1146: uint16(2),
	1147: uint16(3),
	1148: uint16(1),
	1149: uint16(sym_comment),
	1150: uint16(229),
	1151: uint16(1),
	1152: uint16(sym__doctype),
	1153: uint16(2),
	1154: uint16(3),
	1155: uint16(1),
	1156: uint16(sym_comment),
	1157: uint16(227),
	1158: uint16(1),
	1159: uint16(anon_sym_SQUOTE),
	1160: uint16(2),
	1161: uint16(3),
	1162: uint16(1),
	1163: uint16(sym_comment),
	1164: uint16(231),
	1165: uint16(1),
	1166: uint16(anon_sym_GT),
	1167: uint16(2),
	1168: uint16(3),
	1169: uint16(1),
	1170: uint16(sym_comment),
	1171: uint16(233),
	1172: uint16(1),
	1173: uint16(anon_sym_GT),
	1174: uint16(2),
	1175: uint16(3),
	1176: uint16(1),
	1177: uint16(sym_comment),
	1178: uint16(235),
	1179: uint16(1),
	1180: uint16(aux_sym_doctype_token1),
	1181: uint16(2),
	1182: uint16(3),
	1183: uint16(1),
	1184: uint16(sym_comment),
	1185: uint16(237),
	1186: uint16(1),
	1187: uint16(anon_sym_GT),
	1188: uint16(2),
	1189: uint16(3),
	1190: uint16(1),
	1191: uint16(sym_comment),
	1192: uint16(239),
	1193: uint16(1),
	1195: uint16(2),
	1196: uint16(3),
	1197: uint16(1),
	1198: uint16(sym_comment),
	1199: uint16(241),
	1200: uint16(1),
	1201: uint16(anon_sym_GT),
	1202: uint16(2),
	1203: uint16(3),
	1204: uint16(1),
	1205: uint16(sym_comment),
	1206: uint16(243),
	1207: uint16(1),
	1208: uint16(sym_erroneous_end_tag_name),
	1209: uint16(2),
	1210: uint16(3),
	1211: uint16(1),
	1212: uint16(sym_comment),
	1213: uint16(225),
	1214: uint16(1),
	1215: uint16(anon_sym_SQUOTE),
	1216: uint16(2),
	1217: uint16(3),
	1218: uint16(1),
	1219: uint16(sym_comment),
	1220: uint16(245),
	1221: uint16(1),
	1222: uint16(anon_sym_GT),
	1223: uint16(2),
	1224: uint16(3),
	1225: uint16(1),
	1226: uint16(sym_comment),
	1227: uint16(215),
	1228: uint16(1),
	1229: uint16(sym__end_tag_name),
	1230: uint16(2),
	1231: uint16(3),
	1232: uint16(1),
	1233: uint16(sym_comment),
	1234: uint16(201),
	1235: uint16(1),
	1236: uint16(sym_erroneous_end_tag_name),
	1237: uint16(2),
	1238: uint16(3),
	1239: uint16(1),
	1240: uint16(sym_comment),
	1241: uint16(247),
	1242: uint16(1),
	1243: uint16(aux_sym_doctype_token1),
	1244: uint16(2),
	1245: uint16(3),
	1246: uint16(1),
	1247: uint16(sym_comment),
	1248: uint16(249),
	1249: uint16(1),
	1250: uint16(anon_sym_GT),
	1251: uint16(2),
	1252: uint16(3),
	1253: uint16(1),
	1254: uint16(sym_comment),
	1255: uint16(199),
	1256: uint16(1),
	1257: uint16(sym__end_tag_name),
	1258: uint16(2),
	1259: uint16(3),
	1260: uint16(1),
	1261: uint16(sym_comment),
	1262: uint16(251),
	1263: uint16(1),
	1264: uint16(sym__doctype),
}

var ts_small_parse_table_map = [92]uint32_t{
	1:  uint32(44),
	2:  uint32(88),
	3:  uint32(132),
	4:  uint32(176),
	5:  uint32(217),
	6:  uint32(258),
	7:  uint32(299),
	8:  uint32(313),
	9:  uint32(327),
	10: uint32(341),
	11: uint32(355),
	12: uint32(369),
	13: uint32(383),
	14: uint32(397),
	15: uint32(411),
	16: uint32(425),
	17: uint32(439),
	18: uint32(453),
	19: uint32(467),
	20: uint32(481),
	21: uint32(495),
	22: uint32(509),
	23: uint32(523),
	24: uint32(537),
	25: uint32(551),
	26: uint32(565),
	27: uint32(579),
	28: uint32(593),
	29: uint32(607),
	30: uint32(621),
	31: uint32(635),
	32: uint32(649),
	33: uint32(663),
	34: uint32(678),
	35: uint32(695),
	36: uint32(712),
	37: uint32(729),
	38: uint32(746),
	39: uint32(760),
	40: uint32(774),
	41: uint32(788),
	42: uint32(802),
	43: uint32(816),
	44: uint32(828),
	45: uint32(844),
	46: uint32(860),
	47: uint32(873),
	48: uint32(882),
	49: uint32(891),
	50: uint32(902),
	51: uint32(915),
	52: uint32(928),
	53: uint32(941),
	54: uint32(954),
	55: uint32(963),
	56: uint32(976),
	57: uint32(984),
	58: uint32(994),
	59: uint32(1004),
	60: uint32(1012),
	61: uint32(1022),
	62: uint32(1030),
	63: uint32(1038),
	64: uint32(1046),
	65: uint32(1054),
	66: uint32(1062),
	67: uint32(1072),
	68: uint32(1082),
	69: uint32(1092),
	70: uint32(1102),
	71: uint32(1112),
	72: uint32(1122),
	73: uint32(1132),
	74: uint32(1139),
	75: uint32(1146),
	76: uint32(1153),
	77: uint32(1160),
	78: uint32(1167),
	79: uint32(1174),
	80: uint32(1181),
	81: uint32(1188),
	82: uint32(1195),
	83: uint32(1202),
	84: uint32(1209),
	85: uint32(1216),
	86: uint32(1223),
	87: uint32(1230),
	88: uint32(1237),
	89: uint32(1244),
	90: uint32(1251),
	91: uint32(1258),
}

var ts_parse_actions = [253]TSParseActionEntry{
	0: {},
	1: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	2: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeRecover)})),
	3: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	4: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fextra: libc.BoolUint8(true1 != 0),
	}})))),
	5: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	6: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token: uint8(TSParseActionTypeReduce),
		Fsymbol:     uint16(sym_document),
	})))),
	7: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	8: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
	}})))),
	9: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	10: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
	}})))),
	11: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	12: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
	}})))),
	13: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	14: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
	}})))),
	15: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	16: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(93)),
	}})))),
	17: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	18: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(57)),
	}})))),
	19: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	20: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(68)),
	}})))),
	21: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	22: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
	}})))),
	23: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	24: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
	}})))),
	25: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	26: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
	}})))),
	27: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	28: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
	}})))),
	29: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	30: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
	}})))),
	31: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	32: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(27)),
	}})))),
	33: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	34: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
	}})))),
	35: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	36: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
	}})))),
	37: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	38: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	39: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(93)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	40: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	41: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	42: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(57)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	43: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	44: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	45: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(89)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	46: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	47: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	48: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(6)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	49: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	50: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	51: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	52: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_document),
	})))),
	53: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	54: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
	}})))),
	55: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	56: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	57: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(77)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	58: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(2),
	}})),
	59: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	60: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(52)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	61: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	62: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	63: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(85)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	64: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	65: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	66: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(8)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	67: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	68: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_end_tag),
	})))),
	69: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	70: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_end_tag),
	})))),
	71: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	72: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_style_element),
	})))),
	73: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	74: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_style_element),
	})))),
	75: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	76: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_start_tag),
	})))),
	77: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	78: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_start_tag),
	})))),
	79: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	80: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_tag),
	})))),
	81: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	82: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_tag),
	})))),
	83: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	84: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_erroneous_end_tag),
	})))),
	85: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	86: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_erroneous_end_tag),
	})))),
	87: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	88: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_element),
	})))),
	89: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	90: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_element),
	})))),
	91: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	92: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_script_element),
	})))),
	93: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	94: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_script_element),
	})))),
	95: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	96: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_style_element),
	})))),
	97: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	98: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_style_element),
	})))),
	99: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	100: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_doctype),
	})))),
	101: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	102: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_doctype),
	})))),
	103: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	104: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_start_tag),
	})))),
	105: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_start_tag),
	})))),
	107: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	108: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_script_element),
	})))),
	109: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	110: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_script_element),
	})))),
	111: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	112: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_element),
	})))),
	113: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	114: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_element),
	})))),
	115: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	116: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_element),
	})))),
	117: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	118: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_element),
	})))),
	119: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	120: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_tag),
	})))),
	121: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	122: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_tag),
	})))),
	123: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	124: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
	125: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	126: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
	127: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(45)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	128: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	129: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
	}})))),
	130: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	131: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(34)),
	}})))),
	132: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	133: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
	}})))),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	135: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
	}})))),
	136: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	137: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
	}})))),
	138: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	139: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
	}})))),
	140: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	141: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
	}})))),
	142: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	143: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(67)),
	}})))),
	144: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	145: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(51)),
	}})))),
	146: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	147: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
	148: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(51)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	149: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	150: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(66)),
	}})))),
	151: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	152: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(61)),
	}})))),
	153: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	154: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(65)),
	}})))),
	155: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	156: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_attribute),
	})))),
	157: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	158: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(46)),
	}})))),
	159: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	160: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(56)),
	}})))),
	161: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	162: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
	}})))),
	163: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	164: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(62)),
	}})))),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	166: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
	}})))),
	167: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	168: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
	}})))),
	169: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
	}})))),
	171: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	172: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(88)),
	}})))),
	173: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	174: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(71)),
	}})))),
	175: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	176: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quoted_attribute_value),
	})))),
	177: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	178: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quoted_attribute_value),
	})))),
	179: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	180: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
	}})))),
	181: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	182: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
	}})))),
	183: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	184: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
	}})))),
	185: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	186: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
	}})))),
	187: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	188: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
	}})))),
	189: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	190: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(92)),
	}})))),
	191: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	192: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
	}})))),
	193: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	194: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(70)),
	}})))),
	195: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	196: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_attribute),
	})))),
	197: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	198: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
	}})))),
	199: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	200: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
	}})))),
	201: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	202: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
	}})))),
	203: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	204: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_script_start_tag),
	})))),
	205: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	206: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
	}})))),
	207: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	208: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
	}})))),
	209: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	210: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_script_start_tag),
	})))),
	211: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	212: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_style_start_tag),
	})))),
	213: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	214: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_style_start_tag),
	})))),
	215: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	216: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
	}})))),
	217: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	218: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
	}})))),
	219: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	220: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
	}})))),
	221: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	222: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
	}})))),
	223: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	224: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
	}})))),
	225: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	226: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(64)),
	}})))),
	227: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	228: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
	}})))),
	229: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	230: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
	}})))),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	232: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(17)),
	}})))),
	233: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	234: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
	}})))),
	235: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	236: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(79)),
	}})))),
	237: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	238: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
	}})))),
	239: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	240: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	241: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	242: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
	}})))),
	243: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	244: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
	}})))),
	245: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	246: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
	}})))),
	247: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	248: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
	}})))),
	249: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	250: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
	}})))),
	251: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	252: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
	}})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token__start_tag_name = 0
const ts_external_token__script_start_tag_name = 1
const ts_external_token__style_start_tag_name = 2
const ts_external_token__end_tag_name = 3
const ts_external_token_erroneous_end_tag_name = 4
const ts_external_token_SLASH_GT = 5
const ts_external_token__implicit_end_tag = 6
const ts_external_token_raw_text = 7
const ts_external_token_comment = 8

var ts_external_scanner_symbol_map = [9]TSSymbol{
	0: uint16(sym__start_tag_name),
	1: uint16(sym__script_start_tag_name),
	2: uint16(sym__style_start_tag_name),
	3: uint16(sym__end_tag_name),
	4: uint16(sym_erroneous_end_tag_name),
	5: uint16(anon_sym_SLASH_GT),
	6: uint16(sym__implicit_end_tag),
	7: uint16(sym_raw_text),
	8: uint16(sym_comment),
}

var ts_external_scanner_states = [10][9]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
		3: libc.BoolUint8(true1 != 0),
		4: libc.BoolUint8(true1 != 0),
		5: libc.BoolUint8(true1 != 0),
		6: libc.BoolUint8(true1 != 0),
		7: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	2: {
		8: libc.BoolUint8(true1 != 0),
	},
	3: {
		6: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	4: {
		5: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	5: {
		7: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	6: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	7: {
		3: libc.BoolUint8(true1 != 0),
		4: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	8: {
		4: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
	9: {
		3: libc.BoolUint8(true1 != 0),
		8: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_html(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fversion:                   uint32(LANGUAGE_VERSION),
	Fsymbol_count:              uint32(SYMBOL_COUNT),
	Ftoken_count:               uint32(TOKEN_COUNT),
	Fexternal_token_count:      uint32(EXTERNAL_TOKEN_COUNT),
	Fstate_count:               uint32(STATE_COUNT),
	Flarge_state_count:         uint32(LARGE_STATE_COUNT),
	Fproduction_id_count:       uint32(PRODUCTION_ID_COUNT),
	Fmax_alias_sequence_length: uint16(MAX_ALIAS_SEQUENCE_LENGTH),
	Fparse_table:               uintptr(unsafe.Pointer(&ts_parse_table)),
	Fsmall_parse_table:         uintptr(unsafe.Pointer(&ts_small_parse_table)),
	Fsmall_parse_table_map:     uintptr(unsafe.Pointer(&ts_small_parse_table_map)),
	Fparse_actions:             uintptr(unsafe.Pointer(&ts_parse_actions)),
	Fsymbol_names:              uintptr(unsafe.Pointer(&ts_symbol_names)),
	Fsymbol_metadata:           uintptr(unsafe.Pointer(&ts_symbol_metadata)),
	Fpublic_symbol_map:         uintptr(unsafe.Pointer(&ts_symbol_map)),
	Falias_map:                 uintptr(unsafe.Pointer(&ts_non_terminal_alias_map)),
	Falias_sequences:           uintptr(unsafe.Pointer(&ts_alias_sequences)),
	Flex_modes:                 uintptr(unsafe.Pointer(&ts_lex_modes)),
	Fexternal_scanner: struct {
		Fstates      uintptr
		Fsymbol_map  uintptr
		Fcreate      uintptr
		Fdestroy     uintptr
		Fscan        uintptr
		Fserialize   uintptr
		Fdeserialize uintptr
	}{
		Fstates:     uintptr(unsafe.Pointer(&ts_external_scanner_states)),
		Fsymbol_map: uintptr(unsafe.Pointer(&ts_external_scanner_symbol_map)),
	},
	Fprimary_state_ids: uintptr(unsafe.Pointer(&ts_primary_state_ids)),
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_html_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_html_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_html_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_html_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_html_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "(uint32_t)((&scanner->tags)->size - 1) < (&scanner->tags)->size\x00combined.c\x00</SCRIPT\x00</STYLE\x00end\x00<!\x00doctype_token1\x00>\x00doctype\x00<\x00/>\x00</\x00=\x00attribute_name\x00attribute_value\x00entity\x00'\x00\"\x00text\x00tag_name\x00erroneous_end_tag_name\x00_implicit_end_tag\x00raw_text\x00comment\x00document\x00_node\x00element\x00script_element\x00style_element\x00start_tag\x00self_closing_tag\x00end_tag\x00erroneous_end_tag\x00attribute\x00quoted_attribute_value\x00document_repeat1\x00start_tag_repeat1\x00"
