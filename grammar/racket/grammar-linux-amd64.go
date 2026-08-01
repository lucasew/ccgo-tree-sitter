// Code generated for linux/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-racket/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-racket -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && amd64

package grammar_racket

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
const EXTERNAL_TOKEN_COUNT = 1
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
const LARGE_STATE_COUNT = 93
const L_ctermid = 20
const L_cuserid = 20
const L_tmpnam = 20
const MAX_ALIAS_SEQUENCE_LENGTH = 5
const PRODUCTION_ID_COUNT = 1
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const P_tmpdir = "/tmp"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 257
const SYMBOL_COUNT = 81
const TMP_MAX = 10000
const TOKEN_COUNT = 49
const TREE_SITTER_SERIALIZATION_BUFFER_SIZE = 1024
const UINT16_MAX = 0xffff
const UINT32_MAX = "0xffffffffu"
const UINT64_MAX = "0xffffffffffffffffu"
const UINT8_MAX = 0xff
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

type locale_t = uintptr

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

/*
 *  Lexer Macros
 */

/*
 *  Parse Table Macros
 */

type TokenType = int32

const HERE_STRING_BODY = 0

// C documentation
//
//	// a hand written string implmentation
//	// data[0], data[1], ..., data[len-1] is the content of string
//	// data[len] is `\0` for typical string `char*` compatibility
//	// So 0 <= len < cap
type String = struct {
	Flen1 size_t
	Fcap1 size_t
	Fdata uintptr
}

func check_alloc(tls *libc.TLS, ptr uintptr) {
	if ptr == libc.UintptrFromInt32(0) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, 0)
		libc.Xabort(tls)
	}
}

func string_new(tls *libc.TLS) (r String) {
	var cap1, init_len size_t
	var tmp uintptr
	_, _, _ = cap1, init_len, tmp
	init_len = uint64(16)
	// (init_len + 1) for null terminator
	cap1 = init_len + uint64(1)
	tmp = libc.Xcalloc(tls, uint64(1), uint64(1)*cap1)
	check_alloc(tls, tmp)
	return String{
		Fcap1: cap1,
		Fdata: tmp,
	}
}

func string_resize(tls *libc.TLS, str uintptr, new_cap size_t) {
	var block uintptr
	_ = block
	block = libc.Xrealloc(tls, (*String)(unsafe.Pointer(str)).Fdata, new_cap*uint64(1))
	check_alloc(tls, block)
	(*String)(unsafe.Pointer(str)).Fdata = block
	(*String)(unsafe.Pointer(str)).Fcap1 = new_cap
	libc.Xmemset(tls, (*String)(unsafe.Pointer(str)).Fdata+uintptr((*String)(unsafe.Pointer(str)).Flen1), 0, (new_cap-(*String)(unsafe.Pointer(str)).Flen1)*uint64(1))
}

func string_push(tls *libc.TLS, str uintptr, _elem int32_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*int32_t)(unsafe.Pointer(bp)) = _elem
	if (*String)(unsafe.Pointer(str)).Flen1+uint64(4) >= (*String)(unsafe.Pointer(str)).Fcap1 {
		// str->cap * 2 + 1 > str->len + sizeof(elem) always holds
		// as str->cap > 16
		string_resize(tls, str, (*String)(unsafe.Pointer(str)).Fcap1*uint64(2)+uint64(1))
	}
	// NOTE: we don't consider little-endian/big-endian here
	// the character in string is only for compare.
	// They only need to be store in consistent way
	libc.Xmemcpy(tls, (*String)(unsafe.Pointer(str)).Fdata+uintptr((*String)(unsafe.Pointer(str)).Flen1), bp, uint64(4))
	*(*size_t)(unsafe.Pointer(str)) += uint64(4)
}

func string_free(tls *libc.TLS, str uintptr) {
	if (*String)(unsafe.Pointer(str)).Fdata != libc.UintptrFromInt32(0) {
		libc.Xfree(tls, (*String)(unsafe.Pointer(str)).Fdata)
		(*String)(unsafe.Pointer(str)).Fdata = libc.UintptrFromInt32(0)
		(*String)(unsafe.Pointer(str)).Flen1 = uint64(0)
		(*String)(unsafe.Pointer(str)).Fcap1 = uint64(0)
	}
}

func string_clear(tls *libc.TLS, str uintptr) {
	libc.Xmemset(tls, (*String)(unsafe.Pointer(str)).Fdata, 0, (*String)(unsafe.Pointer(str)).Flen1*uint64(1))
	(*String)(unsafe.Pointer(str)).Flen1 = uint64(0)
}

func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
}

func skip(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
}

// C documentation
//
//	// NOTE: only "\n" is allowed as newline here,
//	// It implies that "\r" can also be terminator.
func isnewline(tls *libc.TLS, chr int32_t) (r uint8) {
	return libc.BoolUint8(chr == int32('\n'))
}

// C documentation
//
//	// `read_line` read strings until a newline or EOF
func read_line(tls *libc.TLS, str uintptr, lexer uintptr) {
	for !(isnewline(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0) && !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
		string_push(tls, str, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead)
		advance(tls, lexer)
	}
}

// C documentation
//
//	// Suppose terminator is `T`, newline (\n) is `$`,
//	// It should accept "#<<T$T" or "#<<T$...$T", where `...`
//	// is the string content.
func scan(tls *libc.TLS, lexer uintptr, valid_symbols uintptr) (r uint8) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* current_line at bp+24 */ String
	var _ /* terminator at bp+0 */ String
	if !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(HERE_STRING_BODY))) != 0) {
		return libc.BoolUint8(false1 != 0)
	}
	*(*String)(unsafe.Pointer(bp)) = string_new(tls)
	read_line(tls, bp, lexer)
	if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
		string_free(tls, bp)
		return libc.BoolUint8(false1 != 0)
	}
	// skip `\n`
	skip(tls, lexer)
	*(*String)(unsafe.Pointer(bp + 24)) = string_new(tls)
	for int32(true1) != 0 {
		read_line(tls, bp+24, lexer)
		if libc.Xstrcmp(tls, (*(*String)(unsafe.Pointer(bp))).Fdata, (*(*String)(unsafe.Pointer(bp + 24))).Fdata) == 0 {
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(HERE_STRING_BODY)
			string_free(tls, bp)
			string_free(tls, bp+24)
			return libc.BoolUint8(true1 != 0)
		}
		if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
			string_free(tls, bp)
			string_free(tls, bp+24)
			return libc.BoolUint8(false1 != 0)
		}
		string_clear(tls, bp+24)
		// skip `\n`
		skip(tls, lexer)
	}
	return r
}

func tree_sitter_racket_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_racket_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_racket_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
}

func tree_sitter_racket_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	return scan(tls, lexer, valid_symbols)
}

func tree_sitter_racket_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
}

type ts_symbol_identifiers = int32

const aux_sym__skip_token1 = 1
const sym_dot = 2
const aux_sym_comment_token1 = 3
const anon_sym_POUND_PIPE = 4
const aux_sym_block_comment_token1 = 5
const anon_sym_PIPE_POUND = 6
const anon_sym_POUND_SEMI = 7
const sym__line_comment = 8
const sym_boolean = 9
const anon_sym_POUND = 10
const anon_sym_POUND_LT_LT = 11
const aux_sym_regex_token1 = 12
const anon_sym_DQUOTE = 13
const aux_sym__real_string_token1 = 14
const sym_escape_sequence = 15
const sym_number = 16
const sym_decimal = 17
const sym_character = 18
const sym_symbol = 19
const sym_keyword = 20
const anon_sym_POUND_AMP = 21
const anon_sym_LPAREN = 22
const anon_sym_RPAREN = 23
const anon_sym_LBRACK = 24
const anon_sym_RBRACK = 25
const anon_sym_LBRACE = 26
const anon_sym_RBRACE = 27
const anon_sym_POUNDfl = 28
const anon_sym_POUNDfx = 29
const anon_sym_POUNDs = 30
const anon_sym_POUNDhash = 31
const anon_sym_POUNDhashalw = 32
const anon_sym_POUNDhasheq = 33
const anon_sym_POUNDhasheqv = 34
const anon_sym_EQ = 35
const anon_sym_SQUOTE = 36
const anon_sym_BQUOTE = 37
const anon_sym_POUND_SQUOTE = 38
const anon_sym_POUND_BQUOTE = 39
const anon_sym_COMMA = 40
const anon_sym_COMMA_AT = 41
const anon_sym_POUND_COMMA = 42
const anon_sym_POUND_COMMA_AT = 43
const anon_sym_POUNDreader = 44
const anon_sym_POUNDlang = 45
const anon_sym_POUND_BANG = 46
const sym_lang_name = 47
const sym__here_string_body = 48
const sym_program = 49
const sym__token = 50
const sym__skip = 51
const sym_comment = 52
const sym_block_comment = 53
const sym_sexp_comment = 54
const sym__datum = 55
const sym_string = 56
const sym_byte_string = 57
const sym_here_string = 58
const sym_regex = 59
const sym__real_string = 60
const sym_box = 61
const sym_list = 62
const sym_vector = 63
const sym_structure = 64
const sym_hash = 65
const sym_graph = 66
const sym_quote = 67
const sym_quasiquote = 68
const sym_syntax = 69
const sym_quasisyntax = 70
const sym_unquote = 71
const sym_unquote_splicing = 72
const sym_unsyntax = 73
const sym_unsyntax_splicing = 74
const sym_extension = 75
const aux_sym_program_repeat1 = 76
const aux_sym_block_comment_repeat1 = 77
const aux_sym_sexp_comment_repeat1 = 78
const aux_sym__real_string_repeat1 = 79
const aux_sym_list_repeat1 = 80

var ts_symbol_names = [81]uintptr{
	0:  __ccgo_ts + 36,
	1:  __ccgo_ts + 40,
	2:  __ccgo_ts + 53,
	3:  __ccgo_ts + 57,
	4:  __ccgo_ts + 72,
	5:  __ccgo_ts + 75,
	6:  __ccgo_ts + 96,
	7:  __ccgo_ts + 99,
	8:  __ccgo_ts + 102,
	9:  __ccgo_ts + 116,
	10: __ccgo_ts + 124,
	11: __ccgo_ts + 126,
	12: __ccgo_ts + 130,
	13: __ccgo_ts + 143,
	14: __ccgo_ts + 145,
	15: __ccgo_ts + 165,
	16: __ccgo_ts + 181,
	17: __ccgo_ts + 188,
	18: __ccgo_ts + 196,
	19: __ccgo_ts + 206,
	20: __ccgo_ts + 213,
	21: __ccgo_ts + 221,
	22: __ccgo_ts + 224,
	23: __ccgo_ts + 226,
	24: __ccgo_ts + 228,
	25: __ccgo_ts + 230,
	26: __ccgo_ts + 232,
	27: __ccgo_ts + 234,
	28: __ccgo_ts + 236,
	29: __ccgo_ts + 240,
	30: __ccgo_ts + 244,
	31: __ccgo_ts + 247,
	32: __ccgo_ts + 253,
	33: __ccgo_ts + 262,
	34: __ccgo_ts + 270,
	35: __ccgo_ts + 279,
	36: __ccgo_ts + 281,
	37: __ccgo_ts + 283,
	38: __ccgo_ts + 285,
	39: __ccgo_ts + 288,
	40: __ccgo_ts + 291,
	41: __ccgo_ts + 293,
	42: __ccgo_ts + 296,
	43: __ccgo_ts + 299,
	44: __ccgo_ts + 303,
	45: __ccgo_ts + 311,
	46: __ccgo_ts + 318,
	47: __ccgo_ts + 321,
	48: __ccgo_ts + 331,
	49: __ccgo_ts + 349,
	50: __ccgo_ts + 357,
	51: __ccgo_ts + 364,
	52: __ccgo_ts + 370,
	53: __ccgo_ts + 378,
	54: __ccgo_ts + 392,
	55: __ccgo_ts + 405,
	56: __ccgo_ts + 412,
	57: __ccgo_ts + 419,
	58: __ccgo_ts + 431,
	59: __ccgo_ts + 443,
	60: __ccgo_ts + 449,
	61: __ccgo_ts + 462,
	62: __ccgo_ts + 466,
	63: __ccgo_ts + 471,
	64: __ccgo_ts + 478,
	65: __ccgo_ts + 488,
	66: __ccgo_ts + 493,
	67: __ccgo_ts + 499,
	68: __ccgo_ts + 505,
	69: __ccgo_ts + 516,
	70: __ccgo_ts + 523,
	71: __ccgo_ts + 535,
	72: __ccgo_ts + 543,
	73: __ccgo_ts + 560,
	74: __ccgo_ts + 569,
	75: __ccgo_ts + 587,
	76: __ccgo_ts + 597,
	77: __ccgo_ts + 613,
	78: __ccgo_ts + 635,
	79: __ccgo_ts + 656,
	80: __ccgo_ts + 677,
}

var ts_symbol_map = [81]TSSymbol{
	1:  uint16(aux_sym__skip_token1),
	2:  uint16(sym_dot),
	3:  uint16(aux_sym_comment_token1),
	4:  uint16(anon_sym_POUND_PIPE),
	5:  uint16(aux_sym_block_comment_token1),
	6:  uint16(anon_sym_PIPE_POUND),
	7:  uint16(anon_sym_POUND_SEMI),
	8:  uint16(sym__line_comment),
	9:  uint16(sym_boolean),
	10: uint16(anon_sym_POUND),
	11: uint16(anon_sym_POUND_LT_LT),
	12: uint16(aux_sym_regex_token1),
	13: uint16(anon_sym_DQUOTE),
	14: uint16(aux_sym__real_string_token1),
	15: uint16(sym_escape_sequence),
	16: uint16(sym_number),
	17: uint16(sym_decimal),
	18: uint16(sym_character),
	19: uint16(sym_symbol),
	20: uint16(sym_keyword),
	21: uint16(anon_sym_POUND_AMP),
	22: uint16(anon_sym_LPAREN),
	23: uint16(anon_sym_RPAREN),
	24: uint16(anon_sym_LBRACK),
	25: uint16(anon_sym_RBRACK),
	26: uint16(anon_sym_LBRACE),
	27: uint16(anon_sym_RBRACE),
	28: uint16(anon_sym_POUNDfl),
	29: uint16(anon_sym_POUNDfx),
	30: uint16(anon_sym_POUNDs),
	31: uint16(anon_sym_POUNDhash),
	32: uint16(anon_sym_POUNDhashalw),
	33: uint16(anon_sym_POUNDhasheq),
	34: uint16(anon_sym_POUNDhasheqv),
	35: uint16(anon_sym_EQ),
	36: uint16(anon_sym_SQUOTE),
	37: uint16(anon_sym_BQUOTE),
	38: uint16(anon_sym_POUND_SQUOTE),
	39: uint16(anon_sym_POUND_BQUOTE),
	40: uint16(anon_sym_COMMA),
	41: uint16(anon_sym_COMMA_AT),
	42: uint16(anon_sym_POUND_COMMA),
	43: uint16(anon_sym_POUND_COMMA_AT),
	44: uint16(anon_sym_POUNDreader),
	45: uint16(anon_sym_POUNDlang),
	46: uint16(anon_sym_POUND_BANG),
	47: uint16(sym_lang_name),
	48: uint16(sym__here_string_body),
	49: uint16(sym_program),
	50: uint16(sym__token),
	51: uint16(sym__skip),
	52: uint16(sym_comment),
	53: uint16(sym_block_comment),
	54: uint16(sym_sexp_comment),
	55: uint16(sym__datum),
	56: uint16(sym_string),
	57: uint16(sym_byte_string),
	58: uint16(sym_here_string),
	59: uint16(sym_regex),
	60: uint16(sym__real_string),
	61: uint16(sym_box),
	62: uint16(sym_list),
	63: uint16(sym_vector),
	64: uint16(sym_structure),
	65: uint16(sym_hash),
	66: uint16(sym_graph),
	67: uint16(sym_quote),
	68: uint16(sym_quasiquote),
	69: uint16(sym_syntax),
	70: uint16(sym_quasisyntax),
	71: uint16(sym_unquote),
	72: uint16(sym_unquote_splicing),
	73: uint16(sym_unsyntax),
	74: uint16(sym_unsyntax_splicing),
	75: uint16(sym_extension),
	76: uint16(aux_sym_program_repeat1),
	77: uint16(aux_sym_block_comment_repeat1),
	78: uint16(aux_sym_sexp_comment_repeat1),
	79: uint16(aux_sym__real_string_repeat1),
	80: uint16(aux_sym_list_repeat1),
}

var ts_symbol_metadata = [81]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {},
	2: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	3: {},
	4: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	5: {},
	6: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	7: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	8: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	9: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	10: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	11: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	12: {},
	13: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	14: {},
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
	},
	22: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	23: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	24: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	25: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	26: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	27: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	28: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	29: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	30: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	31: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	32: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	33: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	35: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	36: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	37: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	38: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	39: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	40: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	41: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	42: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	43: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	44: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	45: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	46: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	47: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	48: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	49: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	50: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	51: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	52: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	53: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	54: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	55: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	57: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	58: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	59: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	60: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	61: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	62: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	63: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	64: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	65: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	66: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	67: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	68: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	69: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	70: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	71: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	72: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	73: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	74: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	75: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	76: {},
	77: {},
	78: {},
	79: {},
	80: {},
}

var ts_alias_sequences = [1][5]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [257]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(3),
	4:   uint16(4),
	5:   uint16(5),
	6:   uint16(6),
	7:   uint16(7),
	8:   uint16(8),
	9:   uint16(3),
	10:  uint16(4),
	11:  uint16(8),
	12:  uint16(5),
	13:  uint16(6),
	14:  uint16(7),
	15:  uint16(3),
	16:  uint16(4),
	17:  uint16(8),
	18:  uint16(5),
	19:  uint16(6),
	20:  uint16(7),
	21:  uint16(21),
	22:  uint16(22),
	23:  uint16(23),
	24:  uint16(24),
	25:  uint16(25),
	26:  uint16(26),
	27:  uint16(27),
	28:  uint16(28),
	29:  uint16(29),
	30:  uint16(30),
	31:  uint16(31),
	32:  uint16(32),
	33:  uint16(33),
	34:  uint16(34),
	35:  uint16(35),
	36:  uint16(36),
	37:  uint16(37),
	38:  uint16(38),
	39:  uint16(37),
	40:  uint16(40),
	41:  uint16(41),
	42:  uint16(42),
	43:  uint16(43),
	44:  uint16(44),
	45:  uint16(45),
	46:  uint16(24),
	47:  uint16(25),
	48:  uint16(48),
	49:  uint16(35),
	50:  uint16(23),
	51:  uint16(26),
	52:  uint16(27),
	53:  uint16(28),
	54:  uint16(29),
	55:  uint16(30),
	56:  uint16(31),
	57:  uint16(48),
	58:  uint16(33),
	59:  uint16(34),
	60:  uint16(36),
	61:  uint16(38),
	62:  uint16(37),
	63:  uint16(40),
	64:  uint16(41),
	65:  uint16(42),
	66:  uint16(43),
	67:  uint16(44),
	68:  uint16(45),
	69:  uint16(24),
	70:  uint16(25),
	71:  uint16(35),
	72:  uint16(23),
	73:  uint16(26),
	74:  uint16(27),
	75:  uint16(28),
	76:  uint16(29),
	77:  uint16(30),
	78:  uint16(31),
	79:  uint16(32),
	80:  uint16(33),
	81:  uint16(36),
	82:  uint16(38),
	83:  uint16(83),
	84:  uint16(83),
	85:  uint16(83),
	86:  uint16(40),
	87:  uint16(41),
	88:  uint16(42),
	89:  uint16(43),
	90:  uint16(44),
	91:  uint16(45),
	92:  uint16(32),
	93:  uint16(93),
	94:  uint16(94),
	95:  uint16(95),
	96:  uint16(96),
	97:  uint16(97),
	98:  uint16(98),
	99:  uint16(99),
	100: uint16(100),
	101: uint16(101),
	102: uint16(102),
	103: uint16(103),
	104: uint16(104),
	105: uint16(105),
	106: uint16(106),
	107: uint16(107),
	108: uint16(108),
	109: uint16(109),
	110: uint16(110),
	111: uint16(111),
	112: uint16(112),
	113: uint16(113),
	114: uint16(114),
	115: uint16(115),
	116: uint16(116),
	117: uint16(117),
	118: uint16(118),
	119: uint16(119),
	120: uint16(120),
	121: uint16(121),
	122: uint16(122),
	123: uint16(123),
	124: uint16(124),
	125: uint16(125),
	126: uint16(126),
	127: uint16(127),
	128: uint16(128),
	129: uint16(129),
	130: uint16(130),
	131: uint16(131),
	132: uint16(132),
	133: uint16(133),
	134: uint16(101),
	135: uint16(128),
	136: uint16(129),
	137: uint16(130),
	138: uint16(131),
	139: uint16(111),
	140: uint16(112),
	141: uint16(94),
	142: uint16(95),
	143: uint16(114),
	144: uint16(97),
	145: uint16(116),
	146: uint16(119),
	147: uint16(96),
	148: uint16(99),
	149: uint16(100),
	150: uint16(127),
	151: uint16(102),
	152: uint16(104),
	153: uint16(107),
	154: uint16(106),
	155: uint16(109),
	156: uint16(93),
	157: uint16(113),
	158: uint16(103),
	159: uint16(115),
	160: uint16(105),
	161: uint16(108),
	162: uint16(122),
	163: uint16(118),
	164: uint16(110),
	165: uint16(117),
	166: uint16(98),
	167: uint16(120),
	168: uint16(123),
	169: uint16(124),
	170: uint16(125),
	171: uint16(132),
	172: uint16(126),
	173: uint16(121),
	174: uint16(120),
	175: uint16(119),
	176: uint16(96),
	177: uint16(99),
	178: uint16(100),
	179: uint16(102),
	180: uint16(104),
	181: uint16(106),
	182: uint16(107),
	183: uint16(109),
	184: uint16(98),
	185: uint16(93),
	186: uint16(113),
	187: uint16(122),
	188: uint16(117),
	189: uint16(123),
	190: uint16(124),
	191: uint16(125),
	192: uint16(126),
	193: uint16(127),
	194: uint16(116),
	195: uint16(129),
	196: uint16(130),
	197: uint16(131),
	198: uint16(111),
	199: uint16(112),
	200: uint16(95),
	201: uint16(114),
	202: uint16(97),
	203: uint16(121),
	204: uint16(101),
	205: uint16(103),
	206: uint16(105),
	207: uint16(108),
	208: uint16(118),
	209: uint16(110),
	210: uint16(132),
	211: uint16(128),
	212: uint16(212),
	213: uint16(212),
	214: uint16(212),
	215: uint16(215),
	216: uint16(215),
	217: uint16(215),
	218: uint16(218),
	219: uint16(219),
	220: uint16(218),
	221: uint16(221),
	222: uint16(219),
	223: uint16(223),
	224: uint16(221),
	225: uint16(218),
	226: uint16(219),
	227: uint16(221),
	228: uint16(219),
	229: uint16(218),
	230: uint16(230),
	231: uint16(231),
	232: uint16(232),
	233: uint16(233),
	234: uint16(234),
	235: uint16(234),
	236: uint16(233),
	237: uint16(237),
	238: uint16(231),
	239: uint16(230),
	240: uint16(233),
	241: uint16(234),
	242: uint16(232),
	243: uint16(230),
	244: uint16(231),
	245: uint16(232),
	246: uint16(121),
	247: uint16(122),
	248: uint16(248),
	249: uint16(248),
	250: uint16(248),
	251: uint16(251),
	252: uint16(252),
	253: uint16(251),
	254: uint16(254),
	255: uint16(252),
	256: uint16(251),
}

var aux_sym__skip_token1_character_set_1 = [10]TSCharacterRange{
	0: {
		Fstart: int32('\t'),
		Fend:   int32('\r'),
	},
	1: {
		Fstart: int32(' '),
		Fend:   int32(' '),
	},
	2: {
		Fstart: int32(0x85),
		Fend:   int32(0x85),
	},
	3: {
		Fstart: int32(0xa0),
		Fend:   int32(0xa0),
	},
	4: {
		Fstart: int32(0x1680),
		Fend:   int32(0x1680),
	},
	5: {
		Fstart: int32(0x2000),
		Fend:   int32(0x200a),
	},
	6: {
		Fstart: int32(0x2028),
		Fend:   int32(0x2029),
	},
	7: {
		Fstart: int32(0x202f),
		Fend:   int32(0x202f),
	},
	8: {
		Fstart: int32(0x205f),
		Fend:   int32(0x205f),
	},
	9: {
		Fstart: int32(0x3000),
		Fend:   int32(0x3000),
	},
}

var sym_escape_sequence_character_set_1 = [13]TSCharacterRange{
	0: {
		Fstart: int32('\n'),
		Fend:   int32('\n'),
	},
	1: {
		Fstart: int32('\r'),
		Fend:   int32('\r'),
	},
	2: {
		Fstart: int32('"'),
		Fend:   int32('"'),
	},
	3: {
		Fstart: int32('\''),
		Fend:   int32('\''),
	},
	4: {
		Fstart: int32('0'),
		Fend:   int32('7'),
	},
	5: {
		Fstart: int32('U'),
		Fend:   int32('U'),
	},
	6: {
		Fstart: int32('\\'),
		Fend:   int32('\\'),
	},
	7: {
		Fstart: int32('a'),
		Fend:   int32('b'),
	},
	8: {
		Fstart: int32('e'),
		Fend:   int32('f'),
	},
	9: {
		Fstart: int32('n'),
		Fend:   int32('n'),
	},
	10: {
		Fstart: int32('r'),
		Fend:   int32('r'),
	},
	11: {
		Fstart: int32('t'),
		Fend:   int32('v'),
	},
	12: {
		Fstart: int32('x'),
		Fend:   int32('x'),
	},
}

var sym_symbol_character_set_1 = [21]TSCharacterRange{
	0: {
		Fend: int32(0x08),
	},
	1: {
		Fstart: int32(0x0e),
		Fend:   int32(0x1f),
	},
	2: {
		Fstart: int32('!'),
		Fend:   int32('!'),
	},
	3: {
		Fstart: int32('#'),
		Fend:   int32('&'),
	},
	4: {
		Fstart: int32('*'),
		Fend:   int32('+'),
	},
	5: {
		Fstart: int32('-'),
		Fend:   int32(':'),
	},
	6: {
		Fstart: int32('<'),
		Fend:   int32('Z'),
	},
	7: {
		Fstart: int32('\\'),
		Fend:   int32('\\'),
	},
	8: {
		Fstart: int32('^'),
		Fend:   int32('_'),
	},
	9: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	10: {
		Fstart: int32('|'),
		Fend:   int32('|'),
	},
	11: {
		Fstart: int32('~'),
		Fend:   int32(0x84),
	},
	12: {
		Fstart: int32(0x86),
		Fend:   int32(0x9f),
	},
	13: {
		Fstart: int32(0xa1),
		Fend:   int32(0x167f),
	},
	14: {
		Fstart: int32(0x1681),
		Fend:   int32(0x1fff),
	},
	15: {
		Fstart: int32(0x200b),
		Fend:   int32(0x2027),
	},
	16: {
		Fstart: int32(0x202a),
		Fend:   int32(0x202e),
	},
	17: {
		Fstart: int32(0x2030),
		Fend:   int32(0x205e),
	},
	18: {
		Fstart: int32(0x2060),
		Fend:   int32(0x2fff),
	},
	19: {
		Fstart: int32(0x3001),
		Fend:   int32(0xfefe),
	},
	20: {
		Fstart: int32(0xff00),
		Fend:   int32(0x10ffff),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v4 uint8
	var half_size, i, i1, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i2, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29, i3, i30, i31, i32, i33, i34, i35, i36, i37, i38, i39, i4, i40, i41, i42, i43, i44, i45, i46, i47, i48, i49, i5, i50, i51, i52, i53, i54, i55, i56, i57, i6, i7, i8, i9, index, mid_index, size uint32_t
	var lookahead1, v3 int32_t
	var range_token, range_token1, v2 uintptr
	var v27 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i2, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29, i3, i30, i31, i32, i33, i34, i35, i36, i37, i38, i39, i4, i40, i41, i42, i43, i44, i45, i46, i47, i48, i49, i5, i50, i51, i52, i53, i54, i55, i56, i57, i6, i7, i8, i9, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v2, v27, v3, v4
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
	lookahead1 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
	eof = (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer)
	switch libc.Int32FromUint16(state) {
	case 0:
		if eof != 0 {
			state = uint16(233)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token[i]) == lookahead1 {
				state = map_token[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym__skip_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(10) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _5
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _5
	_5:
		if v4 != 0 {
			state = uint16(234)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead1 == int32('\r') {
			state = uint16(259)
			goto next_state
		}
		if lookahead1 == int32('U') {
			state = uint16(226)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(228)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(227)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(261)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_escape_sequence_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(13) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _9
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _9
	_9:
		if v4 != 0 {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead1 == int32(' ') {
			state = uint16(480)
			goto next_state
		}
		return result
	case int32(3):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token1[i1]) == lookahead1 {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _10
		_10:
			;
			i1 = i1 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(270)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym__skip_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(10) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _14
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _14
	_14:
		if v4 != 0 {
			state = uint16(234)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32(0xfeff) {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(4):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _15
		_15:
			;
			i2 = i2 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(270)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym__skip_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(10) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _19
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _19
	_19:
		if v4 != 0 {
			state = uint16(234)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\'') || int32(')') < lookahead1) && (lookahead1 < int32('[') || int32(']') < lookahead1) && (lookahead1 < int32('{') || int32('}') < lookahead1) && lookahead1 != int32(0xfeff) {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead1 == int32('"') {
			state = uint16(256)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(249)
			goto next_state
		}
		if lookahead1 == int32('(') {
			state = uint16(456)
			goto next_state
		}
		if lookahead1 == int32('=') {
			state = uint16(469)
			goto next_state
		}
		if lookahead1 == int32('[') {
			state = uint16(458)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(460)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead1 == int32('"') {
			state = uint16(256)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(1)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead1 == int32('#') {
			state = uint16(190)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(174)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(31)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead1 == int32('#') {
			state = uint16(143)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(207)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(42)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(280)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead1 == int32('#') {
			state = uint16(10)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(13)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(181)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(9)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead1 == int32('#') {
			state = uint16(10)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(14)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(181)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(11):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token3[i3]) == lookahead1 {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _20
		_20:
			;
			i3 = i3 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(12):
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token4[i4]) == lookahead1 {
				state = map_token4[i4+uint32(1)]
				goto next_state
			}
			goto _21
		_21:
			;
			i4 = i4 + uint32(2)
		}
		return result
	case int32(13):
		if lookahead1 == int32('#') {
			state = uint16(14)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(13)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead1 == int32('#') {
			state = uint16(14)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead1 == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('L') || lookahead1 == int32('S') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead1 == int32('#') {
			state = uint16(15)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('L') || lookahead1 == int32('S') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(131)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead1 == int32('#') {
			state = uint16(241)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(240)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead1 == int32('#') {
			state = uint16(191)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(206)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(36)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(322)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead1 == int32('#') {
			state = uint16(20)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(22)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(214)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(129)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead1 == int32('#') {
			state = uint16(20)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(21)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(214)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead1 == int32('#') {
			state = uint16(21)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead1 == int32('#') {
			state = uint16(21)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(129)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead1 == int32('#') {
			state = uint16(192)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(196)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(37)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead1 == int32('#') {
			state = uint16(25)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(26)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(203)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(24)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead1 == int32('#') {
			state = uint16(25)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(27)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(203)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead1 == int32('#') {
			state = uint16(27)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(26)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead1 == int32('#') {
			state = uint16(27)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead1 == int32('#') {
			state = uint16(193)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(216)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(33)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead1 == int32('.') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead1 == int32('.') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead1 == int32('.') {
			state = uint16(175)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(381)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(133)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead1 == int32('.') {
			state = uint16(207)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(42)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(280)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead1 == int32('.') {
			state = uint16(217)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(389)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(142)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead1 == int32('.') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead1 == int32('.') {
			state = uint16(178)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(54)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(324)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead1 == int32('.') {
			state = uint16(208)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(387)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(328)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead1 == int32('.') {
			state = uint16(197)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(388)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(141)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(348)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead1 == int32('.') {
			state = uint16(220)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(55)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(304)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead1 == int32('.') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead1 == int32('.') {
			state = uint16(180)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(41)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(302)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead1 == int32('.') {
			state = uint16(180)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(166)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(137)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(302)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead1 == int32('.') {
			state = uint16(211)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(382)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(134)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead1 == int32('.') {
			state = uint16(200)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(56)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(344)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead1 == int32('.') {
			state = uint16(222)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(45)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(308)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead1 == int32('.') {
			state = uint16(222)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(166)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(137)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(308)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead1 == int32('.') {
			state = uint16(181)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(383)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(135)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead1 == int32('.') {
			state = uint16(213)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(48)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(336)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead1 == int32('.') {
			state = uint16(213)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(166)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(137)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(336)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead1 == int32('.') {
			state = uint16(202)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(50)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(356)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead1 == int32('.') {
			state = uint16(202)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(166)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(137)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(356)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead1 == int32('.') {
			state = uint16(223)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(383)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(135)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead1 == int32('.') {
			state = uint16(214)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(383)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(135)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead1 == int32('.') {
			state = uint16(203)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(383)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(135)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead1 == int32('.') {
			state = uint16(182)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(384)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(136)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(332)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead1 == int32('.') {
			state = uint16(224)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(386)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(139)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(316)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead1 == int32('.') {
			state = uint16(204)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(385)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(138)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(352)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead1 == int32('.') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead1 == int32('.') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead1 == int32('.') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead1 == int32('.') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead1 == int32('.') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead1 == int32('.') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead1 == int32('<') {
			state = uint16(253)
			goto next_state
		}
		return result
	case int32(64):
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token5[i5]) == lookahead1 {
				state = map_token5[i5+uint32(1)]
				goto next_state
			}
			goto _22
		_22:
			;
			i5 = i5 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(405)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead1 == int32('\\') {
			state = uint16(231)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(115)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _26
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _26
		_26:
		}
		if v27 && v4 != 0 {
			state = uint16(454)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead1 == int32('a') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead1 == int32('a') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead1 == int32('a') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead1 == int32('a') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead1 == int32('a') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead1 == int32('b') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead1 == int32('b') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead1 == int32('c') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead1 == int32('c') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead1 == int32('d') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead1 == int32('d') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead1 == int32('e') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead1 == int32('e') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead1 == int32('e') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead1 == int32('e') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead1 == int32('e') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead1 == int32('e') {
			state = uint16(68)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead1 == int32('e') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead1 == int32('f') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead1 == int32('g') {
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead1 == int32('g') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead1 == int32('h') {
			state = uint16(465)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead1 == int32('i') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead1 == int32('k') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead1 == int32('l') {
			state = uint16(402)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead1 == int32('l') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead1 == int32('l') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead1 == int32('l') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead1 == int32('n') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead1 == int32('n') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead1 == int32('n') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead1 == int32('n') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead1 == int32('o') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead1 == int32('p') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead1 == int32('q') {
			state = uint16(467)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead1 == int32('r') {
			state = uint16(479)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead1 == int32('r') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead1 == int32('s') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead1 == int32('s') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead1 == int32('s') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead1 == int32('t') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead1 == int32('t') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead1 == int32('u') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead1 == int32('u') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead1 == int32('u') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead1 == int32('w') {
			state = uint16(466)
			goto next_state
		}
		return result
	case int32(112):
		if lookahead1 == int32('w') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead1 == int32('x') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead1 == int32('|') {
			state = uint16(453)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead1 == int32('|') {
			state = uint16(454)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead1 == int32(' ') || lookahead1 == int32('/') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(176)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(390)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(218)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(177)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(367)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(209)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(392)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(198)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(391)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(219)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(376)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(210)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(370)
			goto next_state
		}
		return result
	case int32(124):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(199)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(373)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(179)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(366)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(212)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(369)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(221)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(183)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(129):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(215)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(130):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(201)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(372)
			goto next_state
		}
		return result
	case int32(131):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(225)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(132):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(205)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(133):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(134):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(136):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(163)
			goto next_state
		}
		return result
	case int32(137):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(162)
			goto next_state
		}
		return result
	case int32(138):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(139):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(140):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(142):
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(143):
		i6 = uint32(0)
		for {
			if !(uint64(i6) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token6[i6]) == lookahead1 {
				state = map_token6[i6+uint32(1)]
				goto next_state
			}
			goto _28
		_28:
			;
			i6 = i6 + uint32(2)
		}
		return result
	case int32(144):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(145):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(146):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(147):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(148):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(149):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(150):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(151):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(152):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(154):
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(155):
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(156):
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(157):
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(158):
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(159):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(160):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(161):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(162):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(163):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(164):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(165):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(166):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(167):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(168):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(169):
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(170):
		if lookahead1 == int32('T') || lookahead1 == int32('t') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(365)
			goto next_state
		}
		return result
	case int32(171):
		if lookahead1 == int32('T') || lookahead1 == int32('t') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(368)
			goto next_state
		}
		return result
	case int32(172):
		if lookahead1 == int32('T') || lookahead1 == int32('t') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(173):
		if lookahead1 == int32('T') || lookahead1 == int32('t') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(374)
			goto next_state
		}
		return result
	case int32(174):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(286)
			goto next_state
		}
		return result
	case int32(175):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(300)
			goto next_state
		}
		return result
	case int32(176):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(390)
			goto next_state
		}
		return result
	case int32(177):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(367)
			goto next_state
		}
		return result
	case int32(178):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(330)
			goto next_state
		}
		return result
	case int32(179):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(366)
			goto next_state
		}
		return result
	case int32(180):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(312)
			goto next_state
		}
		return result
	case int32(181):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(182):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(338)
			goto next_state
		}
		return result
	case int32(183):
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(184):
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(185):
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(365)
			goto next_state
		}
		return result
	case int32(186):
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(187):
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(368)
			goto next_state
		}
		return result
	case int32(188):
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(189):
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(374)
			goto next_state
		}
		return result
	case int32(190):
		if lookahead1 == int32('E') || lookahead1 == int32('I') || lookahead1 == int32('e') || lookahead1 == int32('i') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(191):
		if lookahead1 == int32('E') || lookahead1 == int32('I') || lookahead1 == int32('e') || lookahead1 == int32('i') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(192):
		if lookahead1 == int32('E') || lookahead1 == int32('I') || lookahead1 == int32('e') || lookahead1 == int32('i') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(193):
		if lookahead1 == int32('E') || lookahead1 == int32('I') || lookahead1 == int32('e') || lookahead1 == int32('i') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(194):
		if lookahead1 == int32('I') || lookahead1 == int32('S') || lookahead1 == int32('i') || lookahead1 == int32('s') {
			state = uint16(414)
			goto next_state
		}
		return result
	case int32(195):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(196):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(197):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(354)
			goto next_state
		}
		return result
	case int32(198):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(391)
			goto next_state
		}
		return result
	case int32(199):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(373)
			goto next_state
		}
		return result
	case int32(200):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(201):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(372)
			goto next_state
		}
		return result
	case int32(202):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(360)
			goto next_state
		}
		return result
	case int32(203):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(204):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(358)
			goto next_state
		}
		return result
	case int32(205):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(206):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(327)
			goto next_state
		}
		return result
	case int32(207):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(295)
			goto next_state
		}
		return result
	case int32(208):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(335)
			goto next_state
		}
		return result
	case int32(209):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(392)
			goto next_state
		}
		return result
	case int32(210):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(370)
			goto next_state
		}
		return result
	case int32(211):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(311)
			goto next_state
		}
		return result
	case int32(212):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(369)
			goto next_state
		}
		return result
	case int32(213):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(341)
			goto next_state
		}
		return result
	case int32(214):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(215):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(216):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(291)
			goto next_state
		}
		return result
	case int32(217):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(307)
			goto next_state
		}
		return result
	case int32(218):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(219):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(376)
			goto next_state
		}
		return result
	case int32(220):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(315)
			goto next_state
		}
		return result
	case int32(221):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(222):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(319)
			goto next_state
		}
		return result
	case int32(223):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(224):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(321)
			goto next_state
		}
		return result
	case int32(225):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(226):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(227):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(228):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(264)
			goto next_state
		}
		return result
	case int32(229):
		if lookahead1 == int32('+') || lookahead1 == int32('-') || int32('/') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(482)
			goto next_state
		}
		return result
	case int32(230):
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(231):
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(454)
			goto next_state
		}
		return result
	case int32(232):
		if eof != 0 {
			state = uint16(233)
			goto next_state
		}
		i7 = uint32(0)
		for {
			if !(uint64(i7) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token7[i7]) == lookahead1 {
				state = map_token7[i7+uint32(1)]
				goto next_state
			}
			goto _29
		_29:
			;
			i7 = i7 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(270)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym__skip_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(10) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _33
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _33
	_33:
		if v4 != 0 {
			state = uint16(234)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\'') || int32(')') < lookahead1) && (lookahead1 < int32('[') || int32(']') < lookahead1) && (lookahead1 < int32('{') || int32('}') < lookahead1) && lookahead1 != int32(0xfeff) {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(233):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(234):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__skip_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&aux_sym__skip_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(10) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _37
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _37
	_37:
		if v4 != 0 {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(235):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_dot)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(236):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_dot)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(272)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _41
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _41
		_41:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(237):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\n') && lookahead1 != int32('\r') && lookahead1 != int32(0x85) && lookahead1 != int32(0x2028) && lookahead1 != int32(0x2029) {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(238):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(239):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_block_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(240):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_block_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(241):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_block_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(242):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(243):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(244):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__line_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(244)
			goto next_state
		}
		if lookahead1 == int32('\n') || lookahead1 == int32('\r') || lookahead1 == int32(0x85) || lookahead1 == int32(0x2028) || lookahead1 == int32(0x2029) {
			state = uint16(245)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(245):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__line_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(244)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') && lookahead1 != int32('\r') && lookahead1 != int32(0x85) && lookahead1 != int32(0x2028) && lookahead1 != int32(0x2029) {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(246):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_boolean)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(247):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_boolean)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(92)
			goto next_state
		}
		if lookahead1 == int32('l') {
			state = uint16(462)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(463)
			goto next_state
		}
		return result
	case int32(248):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_boolean)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('r') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(249):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(250):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i8 = uint32(0)
		for {
			if !(uint64(i8) < libc.Uint64FromInt64(136)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token8[i8]) == lookahead1 {
				state = map_token8[i8+uint32(1)]
				goto next_state
			}
			goto _43
		_43:
			;
			i8 = i8 + uint32(2)
		}
		return result
	case int32(251):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i9 = uint32(0)
		for {
			if !(uint64(i9) < libc.Uint64FromInt64(132)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token9[i9]) == lookahead1 {
				state = map_token9[i9+uint32(1)]
				goto next_state
			}
			goto _44
		_44:
			;
			i9 = i9 + uint32(2)
		}
		return result
	case int32(252):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('|') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(253):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_LT_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(254):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_regex_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(255):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_regex_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(254)
			goto next_state
		}
		return result
	case int32(256):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(257):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__real_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('"') && lookahead1 != int32('\\') {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(258):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(259):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\n') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(260):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(261):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(260)
			goto next_state
		}
		return result
	case int32(262):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(263):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(264):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(263)
			goto next_state
		}
		return result
	case int32(265):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(264)
			goto next_state
		}
		return result
	case int32(266):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(265)
			goto next_state
		}
		return result
	case int32(267):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		return result
	case int32(268):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(269):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(270):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i10 = uint32(0)
		for {
			if !(uint64(i10) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token10[i10]) == lookahead1 {
				state = map_token10[i10+uint32(1)]
				goto next_state
			}
			goto _45
		_45:
			;
			i10 = i10 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(427)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(270)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _49
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _49
		_49:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(271):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i11 = uint32(0)
		for {
			if !(uint64(i11) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token11[i11]) == lookahead1 {
				state = map_token11[i11+uint32(1)]
				goto next_state
			}
			goto _51
		_51:
			;
			i11 = i11 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(427)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _55
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _55
		_55:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(272):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i12 = uint32(0)
		for {
			if !(uint64(i12) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token12[i12]) == lookahead1 {
				state = map_token12[i12+uint32(1)]
				goto next_state
			}
			goto _57
		_57:
			;
			i12 = i12 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(427)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(272)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _61
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _61
		_61:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(273):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i13 = uint32(0)
		for {
			if !(uint64(i13) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token13[i13]) == lookahead1 {
				state = map_token13[i13+uint32(1)]
				goto next_state
			}
			goto _63
		_63:
			;
			i13 = i13 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(427)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _67
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _67
		_67:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(274):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i14 = uint32(0)
		for {
			if !(uint64(i14) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token14[i14]) == lookahead1 {
				state = map_token14[i14+uint32(1)]
				goto next_state
			}
			goto _69
		_69:
			;
			i14 = i14 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(428)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(274)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _73
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _73
		_73:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(275):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i15 = uint32(0)
		for {
			if !(uint64(i15) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token15[i15]) == lookahead1 {
				state = map_token15[i15+uint32(1)]
				goto next_state
			}
			goto _75
		_75:
			;
			i15 = i15 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(428)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _79
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _79
		_79:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(276):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i16 = uint32(0)
		for {
			if !(uint64(i16) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token16[i16]) == lookahead1 {
				state = map_token16[i16+uint32(1)]
				goto next_state
			}
			goto _81
		_81:
			;
			i16 = i16 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(277):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i17 = uint32(0)
		for {
			if !(uint64(i17) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token17[i17]) == lookahead1 {
				state = map_token17[i17+uint32(1)]
				goto next_state
			}
			goto _82
		_82:
			;
			i17 = i17 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(278):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i18 = uint32(0)
		for {
			if !(uint64(i18) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token18[i18]) == lookahead1 {
				state = map_token18[i18+uint32(1)]
				goto next_state
			}
			goto _83
		_83:
			;
			i18 = i18 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(279):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i19 = uint32(0)
		for {
			if !(uint64(i19) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token19[i19]) == lookahead1 {
				state = map_token19[i19+uint32(1)]
				goto next_state
			}
			goto _84
		_84:
			;
			i19 = i19 + uint32(2)
		}
		return result
	case int32(280):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(281)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(295)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(207)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(52)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(123)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(280)
			goto next_state
		}
		return result
	case int32(281):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(281)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(294)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(207)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(52)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(282):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i20 = uint32(0)
		for {
			if !(uint64(i20) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token20[i20]) == lookahead1 {
				state = map_token20[i20+uint32(1)]
				goto next_state
			}
			goto _85
		_85:
			;
			i20 = i20 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(428)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(282)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _89
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _89
		_89:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(283):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i21 = uint32(0)
		for {
			if !(uint64(i21) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token21[i21]) == lookahead1 {
				state = map_token21[i21+uint32(1)]
				goto next_state
			}
			goto _91
		_91:
			;
			i21 = i21 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(428)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _95
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _95
		_95:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(284):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(285)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(298)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(450)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(426)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(284)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _100
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _100
		_100:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(285):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(285)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(299)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(450)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(426)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _105
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _105
		_105:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(286):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(287)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 == int32('T') || lookahead1 == int32('t') {
			state = uint16(117)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(286)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(287):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(287)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 == int32('T') || lookahead1 == int32('t') {
			state = uint16(117)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(288):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i22 = uint32(0)
		for {
			if !(uint64(i22) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token22[i22]) == lookahead1 {
				state = map_token22[i22+uint32(1)]
				goto next_state
			}
			goto _107
		_107:
			;
			i22 = i22 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(289):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i23 = uint32(0)
		for {
			if !(uint64(i23) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token23[i23]) == lookahead1 {
				state = map_token23[i23+uint32(1)]
				goto next_state
			}
			goto _108
		_108:
			;
			i23 = i23 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(290):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i24 = uint32(0)
		for {
			if !(uint64(i24) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token24[i24]) == lookahead1 {
				state = map_token24[i24+uint32(1)]
				goto next_state
			}
			goto _109
		_109:
			;
			i24 = i24 + uint32(2)
		}
		return result
	case int32(291):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i25 = uint32(0)
		for {
			if !(uint64(i25) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token25[i25]) == lookahead1 {
				state = map_token25[i25+uint32(1)]
				goto next_state
			}
			goto _110
		_110:
			;
			i25 = i25 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(291)
			goto next_state
		}
		return result
	case int32(292):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i26 = uint32(0)
		for {
			if !(uint64(i26) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token26[i26]) == lookahead1 {
				state = map_token26[i26+uint32(1)]
				goto next_state
			}
			goto _111
		_111:
			;
			i26 = i26 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(293):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i27 = uint32(0)
		for {
			if !(uint64(i27) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token27[i27]) == lookahead1 {
				state = map_token27[i27+uint32(1)]
				goto next_state
			}
			goto _112
		_112:
			;
			i27 = i27 + uint32(2)
		}
		return result
	case int32(294):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(294)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(52)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(295):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(294)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(52)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(123)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(295)
			goto next_state
		}
		return result
	case int32(296):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i28 = uint32(0)
		for {
			if !(uint64(i28) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token28[i28]) == lookahead1 {
				state = map_token28[i28+uint32(1)]
				goto next_state
			}
			goto _113
		_113:
			;
			i28 = i28 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(126)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(297):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i29 = uint32(0)
		for {
			if !(uint64(i29) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token29[i29]) == lookahead1 {
				state = map_token29[i29+uint32(1)]
				goto next_state
			}
			goto _114
		_114:
			;
			i29 = i29 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(298):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(299)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(426)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(298)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _118
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _118
		_118:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(299):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(299)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(426)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _123
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _123
		_123:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(300):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i30 = uint32(0)
		for {
			if !(uint64(i30) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token30[i30]) == lookahead1 {
				state = map_token30[i30+uint32(1)]
				goto next_state
			}
			goto _125
		_125:
			;
			i30 = i30 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(301):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i31 = uint32(0)
		for {
			if !(uint64(i31) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token31[i31]) == lookahead1 {
				state = map_token31[i31+uint32(1)]
				goto next_state
			}
			goto _126
		_126:
			;
			i31 = i31 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(302):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(303)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(312)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(180)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(302)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(303):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(303)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(313)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(180)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(304):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i32 = uint32(0)
		for {
			if !(uint64(i32) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token32[i32]) == lookahead1 {
				state = map_token32[i32+uint32(1)]
				goto next_state
			}
			goto _127
		_127:
			;
			i32 = i32 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(304)
			goto next_state
		}
		return result
	case int32(305):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i33 = uint32(0)
		for {
			if !(uint64(i33) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token33[i33]) == lookahead1 {
				state = map_token33[i33+uint32(1)]
				goto next_state
			}
			goto _128
		_128:
			;
			i33 = i33 + uint32(2)
		}
		return result
	case int32(306):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i34 = uint32(0)
		for {
			if !(uint64(i34) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token34[i34]) == lookahead1 {
				state = map_token34[i34+uint32(1)]
				goto next_state
			}
			goto _129
		_129:
			;
			i34 = i34 + uint32(2)
		}
		return result
	case int32(307):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i35 = uint32(0)
		for {
			if !(uint64(i35) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token35[i35]) == lookahead1 {
				state = map_token35[i35+uint32(1)]
				goto next_state
			}
			goto _130
		_130:
			;
			i35 = i35 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(307)
			goto next_state
		}
		return result
	case int32(308):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(309)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(319)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(222)
			goto next_state
		}
		if lookahead1 == int32('L') || lookahead1 == int32('S') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(118)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(308)
			goto next_state
		}
		return result
	case int32(309):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(309)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(318)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(222)
			goto next_state
		}
		if lookahead1 == int32('L') || lookahead1 == int32('S') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(310):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(52)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(311):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(310)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(52)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(126)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(311)
			goto next_state
		}
		return result
	case int32(312):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(313)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(312)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(313):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(313)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(314):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i36 = uint32(0)
		for {
			if !(uint64(i36) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token36[i36]) == lookahead1 {
				state = map_token36[i36+uint32(1)]
				goto next_state
			}
			goto _131
		_131:
			;
			i36 = i36 + uint32(2)
		}
		return result
	case int32(315):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i37 = uint32(0)
		for {
			if !(uint64(i37) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token37[i37]) == lookahead1 {
				state = map_token37[i37+uint32(1)]
				goto next_state
			}
			goto _132
		_132:
			;
			i37 = i37 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(315)
			goto next_state
		}
		return result
	case int32(316):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i38 = uint32(0)
		for {
			if !(uint64(i38) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token38[i38]) == lookahead1 {
				state = map_token38[i38+uint32(1)]
				goto next_state
			}
			goto _133
		_133:
			;
			i38 = i38 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(316)
			goto next_state
		}
		return result
	case int32(317):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i39 = uint32(0)
		for {
			if !(uint64(i39) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token39[i39]) == lookahead1 {
				state = map_token39[i39+uint32(1)]
				goto next_state
			}
			goto _134
		_134:
			;
			i39 = i39 + uint32(2)
		}
		return result
	case int32(318):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(318)
			goto next_state
		}
		if lookahead1 == int32('L') || lookahead1 == int32('S') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(319):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(318)
			goto next_state
		}
		if lookahead1 == int32('L') || lookahead1 == int32('S') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(118)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(319)
			goto next_state
		}
		return result
	case int32(320):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i40 = uint32(0)
		for {
			if !(uint64(i40) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token40[i40]) == lookahead1 {
				state = map_token40[i40+uint32(1)]
				goto next_state
			}
			goto _135
		_135:
			;
			i40 = i40 + uint32(2)
		}
		return result
	case int32(321):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i41 = uint32(0)
		for {
			if !(uint64(i41) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token41[i41]) == lookahead1 {
				state = map_token41[i41+uint32(1)]
				goto next_state
			}
			goto _136
		_136:
			;
			i41 = i41 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(321)
			goto next_state
		}
		return result
	case int32(322):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i42 = uint32(0)
		for {
			if !(uint64(i42) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token42[i42]) == lookahead1 {
				state = map_token42[i42+uint32(1)]
				goto next_state
			}
			goto _137
		_137:
			;
			i42 = i42 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(123)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(322)
			goto next_state
		}
		return result
	case int32(323):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i43 = uint32(0)
		for {
			if !(uint64(i43) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token43[i43]) == lookahead1 {
				state = map_token43[i43+uint32(1)]
				goto next_state
			}
			goto _138
		_138:
			;
			i43 = i43 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(324):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(325)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(330)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(178)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(324)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(325):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(325)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(331)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(178)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(326):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(326)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(52)
			goto next_state
		}
		if lookahead1 == int32('T') || lookahead1 == int32('t') {
			state = uint16(120)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(327):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(326)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(52)
			goto next_state
		}
		if lookahead1 == int32('T') || lookahead1 == int32('t') {
			state = uint16(120)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(123)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(327)
			goto next_state
		}
		return result
	case int32(328):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i44 = uint32(0)
		for {
			if !(uint64(i44) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token44[i44]) == lookahead1 {
				state = map_token44[i44+uint32(1)]
				goto next_state
			}
			goto _139
		_139:
			;
			i44 = i44 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(126)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(328)
			goto next_state
		}
		return result
	case int32(329):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i45 = uint32(0)
		for {
			if !(uint64(i45) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token45[i45]) == lookahead1 {
				state = map_token45[i45+uint32(1)]
				goto next_state
			}
			goto _140
		_140:
			;
			i45 = i45 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(330):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(331)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(330)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(331):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(331)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(332):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i46 = uint32(0)
		for {
			if !(uint64(i46) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token46[i46]) == lookahead1 {
				state = map_token46[i46+uint32(1)]
				goto next_state
			}
			goto _141
		_141:
			;
			i46 = i46 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(333):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i47 = uint32(0)
		for {
			if !(uint64(i47) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token47[i47]) == lookahead1 {
				state = map_token47[i47+uint32(1)]
				goto next_state
			}
			goto _142
		_142:
			;
			i47 = i47 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(334):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i48 = uint32(0)
		for {
			if !(uint64(i48) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token48[i48]) == lookahead1 {
				state = map_token48[i48+uint32(1)]
				goto next_state
			}
			goto _143
		_143:
			;
			i48 = i48 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(335):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i49 = uint32(0)
		for {
			if !(uint64(i49) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token49[i49]) == lookahead1 {
				state = map_token49[i49+uint32(1)]
				goto next_state
			}
			goto _144
		_144:
			;
			i49 = i49 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(126)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(335)
			goto next_state
		}
		return result
	case int32(336):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(337)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(341)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(213)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(120)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(336)
			goto next_state
		}
		return result
	case int32(337):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(337)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(340)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(213)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(338):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(339)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(338)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(339):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(339)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(340):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(340)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(341):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(340)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(120)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(341)
			goto next_state
		}
		return result
	case int32(342):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i50 = uint32(0)
		for {
			if !(uint64(i50) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token50[i50]) == lookahead1 {
				state = map_token50[i50+uint32(1)]
				goto next_state
			}
			goto _145
		_145:
			;
			i50 = i50 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(342)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(343):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i51 = uint32(0)
		for {
			if !(uint64(i51) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token51[i51]) == lookahead1 {
				state = map_token51[i51+uint32(1)]
				goto next_state
			}
			goto _146
		_146:
			;
			i51 = i51 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(344):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(345)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(350)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(200)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(53)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(344)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(345):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(345)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(351)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(200)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(53)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(346):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(347)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(53)
			goto next_state
		}
		if lookahead1 == int32('T') || lookahead1 == int32('t') {
			state = uint16(121)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(346)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(347):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(347)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(53)
			goto next_state
		}
		if lookahead1 == int32('T') || lookahead1 == int32('t') {
			state = uint16(121)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(348):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i52 = uint32(0)
		for {
			if !(uint64(i52) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token52[i52]) == lookahead1 {
				state = map_token52[i52+uint32(1)]
				goto next_state
			}
			goto _147
		_147:
			;
			i52 = i52 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(348)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(349):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i53 = uint32(0)
		for {
			if !(uint64(i53) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token53[i53]) == lookahead1 {
				state = map_token53[i53+uint32(1)]
				goto next_state
			}
			goto _148
		_148:
			;
			i53 = i53 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(350):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(351)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(53)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(350)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(351):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(351)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(53)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(352):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i54 = uint32(0)
		for {
			if !(uint64(i54) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token54[i54]) == lookahead1 {
				state = map_token54[i54+uint32(1)]
				goto next_state
			}
			goto _149
		_149:
			;
			i54 = i54 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(352)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(353):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i55 = uint32(0)
		for {
			if !(uint64(i55) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token55[i55]) == lookahead1 {
				state = map_token55[i55+uint32(1)]
				goto next_state
			}
			goto _150
		_150:
			;
			i55 = i55 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(354):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i56 = uint32(0)
		for {
			if !(uint64(i56) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token56[i56]) == lookahead1 {
				state = map_token56[i56+uint32(1)]
				goto next_state
			}
			goto _151
		_151:
			;
			i56 = i56 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(354)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(355):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i57 = uint32(0)
		for {
			if !(uint64(i57) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token57[i57]) == lookahead1 {
				state = map_token57[i57+uint32(1)]
				goto next_state
			}
			goto _152
		_152:
			;
			i57 = i57 + uint32(2)
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(356):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(357)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(360)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(202)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(356)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(357):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(357)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(361)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(202)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(358):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(359)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(53)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(358)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(359):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(359)
			goto next_state
		}
		if lookahead1 == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(53)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(360):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(361)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(360)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(361):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(361)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(362):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(423)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(425)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(380)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(362)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _156
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _156
		_156:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(363):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(423)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(425)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(380)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _161
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _161
		_161:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(364):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(423)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(425)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(364)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _166
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _166
		_166:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(365):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(366):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(366)
			goto next_state
		}
		return result
	case int32(367):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(367)
			goto next_state
		}
		return result
	case int32(368):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(52)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(369):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(52)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(369)
			goto next_state
		}
		return result
	case int32(370):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(52)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(370)
			goto next_state
		}
		return result
	case int32(371):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(53)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(372):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(53)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(372)
			goto next_state
		}
		return result
	case int32(373):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(53)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(373)
			goto next_state
		}
		return result
	case int32(374):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(44)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(51)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(375):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(44)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(51)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(269)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(376):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(44)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(51)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(376)
			goto next_state
		}
		return result
	case int32(377):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(433)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _171
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _171
		_171:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(378):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(434)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _176
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _176
		_176:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(379):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(379)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _181
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _181
		_181:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(380):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _186
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _186
		_186:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(381):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(382):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(383):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(384):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(385):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(386):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(387):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(388):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(389):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(390):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') || lookahead1 == int32('1') {
			state = uint16(390)
			goto next_state
		}
		return result
	case int32(391):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(391)
			goto next_state
		}
		return result
	case int32(392):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(392)
			goto next_state
		}
		return result
	case int32(393):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(394):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_decimal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(395):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(396):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(397):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(398):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(399):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(112)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(400):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(107)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(401):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('i') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(402):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(403):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('p') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(404):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('t') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(405):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(406):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(407):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(406)
			goto next_state
		}
		return result
	case int32(408):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(407)
			goto next_state
		}
		return result
	case int32(409):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(408)
			goto next_state
		}
		return result
	case int32(410):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(409)
			goto next_state
		}
		return result
	case int32(411):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(412):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(411)
			goto next_state
		}
		return result
	case int32(413):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_character)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(412)
			goto next_state
		}
		return result
	case int32(414):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(415):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(416)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(417)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(451)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(380)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(429)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(415)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _191
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _191
		_191:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(416):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(416)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(418)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(451)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(380)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(429)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _196
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _196
		_196:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(417):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(418)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(380)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(429)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(417)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _201
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _201
		_201:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(418):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(418)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(380)
			goto next_state
		}
		if int32('D') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('L') || lookahead1 == int32('S') || int32('d') <= lookahead1 && lookahead1 <= int32('f') || lookahead1 == int32('l') || lookahead1 == int32('s') {
			state = uint16(429)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _206
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _206
		_206:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(419):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(442)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _211
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _211
		_211:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(420):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(444)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _216
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _216
		_216:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(421):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(446)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(377)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(430)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(274)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _221
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _221
		_221:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(422):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(443)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _226
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _226
		_226:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(423):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(450)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(424)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(284)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _231
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _231
		_231:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(424):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(450)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(441)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(432)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(284)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _236
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _236
		_236:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(425):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(451)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(378)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(431)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(415)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _241
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _241
		_241:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(426):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(447)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(379)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _246
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _246
		_246:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(427):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(448)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(364)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _251
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _251
		_251:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(428):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(449)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(362)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _256
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _256
		_256:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(429):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(452)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(436)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _261
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _261
		_261:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(430):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(438)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _266
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _266
		_266:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(431):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(439)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _271
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _271
		_271:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(432):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('A') || lookahead1 == int32('a') {
			state = uint16(440)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _276
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _276
		_276:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(433):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(419)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _281
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _281
		_281:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(434):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(420)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _286
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _286
		_286:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(435):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(422)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _291
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _291
		_291:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(436):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(380)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(436)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _296
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _296
		_296:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(437):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('I') || lookahead1 == int32('i') {
			state = uint16(380)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _301
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _301
		_301:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(438):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(419)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _306
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _306
		_306:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(439):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(420)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _311
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _311
		_311:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(440):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(422)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _316
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _316
		_316:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(441):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('N') || lookahead1 == int32('n') {
			state = uint16(435)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _321
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _321
		_321:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(442):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('T') || lookahead1 == int32('t') {
			state = uint16(380)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(363)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _326
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _326
		_326:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(443):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(380)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _331
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _331
		_331:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(444):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('0') || lookahead1 == int32('F') || lookahead1 == int32('f') {
			state = uint16(437)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _336
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _336
		_336:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(445):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(272)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _341
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _341
		_341:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(446):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(282)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _346
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _346
		_346:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(447):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(379)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _351
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _351
		_351:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(448):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(364)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _356
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _356
		_356:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(449):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(362)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _361
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _361
		_361:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(450):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(298)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _366
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _366
		_366:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(451):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(417)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _371
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _371
		_371:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(452):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(436)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _376
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _376
		_376:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(453):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(230)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(114)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _381
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _381
		_381:
		}
		if v27 && v4 != 0 {
			state = uint16(453)
			goto next_state
		}
		return result
	case int32(454):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_keyword)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(231)
			goto next_state
		}
		if lookahead1 == int32('|') {
			state = uint16(115)
			goto next_state
		}
		if v27 = !(eof != 0); v27 {
			v2 = uintptr(unsafe.Pointer(&sym_symbol_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(21) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _386
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _386
		_386:
		}
		if v27 && v4 != 0 {
			state = uint16(454)
			goto next_state
		}
		return result
	case int32(455):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(456):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(457):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(458):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(459):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(460):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(461):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(462):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDfl)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(463):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDfx)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(464):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDs)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(465):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDhash)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(93)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(466):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDhashalw)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(467):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDhasheq)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('v') {
			state = uint16(468)
			goto next_state
		}
		return result
	case int32(468):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDhasheqv)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(469):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(470):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(471):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(472):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(473):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_BQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(474):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(475):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(476)
			goto next_state
		}
		return result
	case int32(476):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(477):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('@') {
			state = uint16(478)
			goto next_state
		}
		return result
	case int32(478):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_COMMA_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(479):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDreader)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(480):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDlang)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(481):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(' ') || lookahead1 == int32('/') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(482):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_lang_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('+') || lookahead1 == int32('-') || int32('/') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(482)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [30]uint16_t{
	0:  uint16('"'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(252),
	4:  uint16('\''),
	5:  uint16(470),
	6:  uint16('('),
	7:  uint16(456),
	8:  uint16(')'),
	9:  uint16(457),
	10: uint16(','),
	11: uint16(474),
	12: uint16('.'),
	13: uint16(235),
	14: uint16(';'),
	15: uint16(237),
	16: uint16('='),
	17: uint16(469),
	18: uint16('['),
	19: uint16(458),
	20: uint16(']'),
	21: uint16(459),
	22: uint16('`'),
	23: uint16(471),
	24: uint16('{'),
	25: uint16(460),
	26: uint16('|'),
	27: uint16(240),
	28: uint16('}'),
	29: uint16(461),
}

var map_token1 = [34]uint16_t{
	0:  uint16('"'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(250),
	4:  uint16('\''),
	5:  uint16(470),
	6:  uint16('('),
	7:  uint16(456),
	8:  uint16(')'),
	9:  uint16(457),
	10: uint16(','),
	11: uint16(475),
	12: uint16('.'),
	13: uint16(236),
	14: uint16(';'),
	15: uint16(237),
	16: uint16('['),
	17: uint16(458),
	18: uint16('\\'),
	19: uint16(230),
	20: uint16(']'),
	21: uint16(459),
	22: uint16('`'),
	23: uint16(471),
	24: uint16('{'),
	25: uint16(460),
	26: uint16('|'),
	27: uint16(114),
	28: uint16('}'),
	29: uint16(461),
	30: uint16('+'),
	31: uint16(421),
	32: uint16('-'),
	33: uint16(421),
}

var map_token2 = [28]uint16_t{
	0:  uint16('"'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(251),
	4:  uint16('\''),
	5:  uint16(470),
	6:  uint16('('),
	7:  uint16(456),
	8:  uint16(','),
	9:  uint16(475),
	10: uint16('.'),
	11: uint16(445),
	12: uint16(';'),
	13: uint16(237),
	14: uint16('['),
	15: uint16(458),
	16: uint16('\\'),
	17: uint16(230),
	18: uint16('`'),
	19: uint16(471),
	20: uint16('{'),
	21: uint16(460),
	22: uint16('|'),
	23: uint16(114),
	24: uint16('+'),
	25: uint16(421),
	26: uint16('-'),
	27: uint16(421),
}

var map_token3 = [18]uint16_t{
	0:  uint16('#'),
	1:  uint16(12),
	2:  uint16('.'),
	3:  uint16(16),
	4:  uint16('/'),
	5:  uint16(223),
	6:  uint16('I'),
	7:  uint16(269),
	8:  uint16('i'),
	9:  uint16(269),
	10: uint16('L'),
	11: uint16(131),
	12: uint16('S'),
	13: uint16(131),
	14: uint16('l'),
	15: uint16(131),
	16: uint16('s'),
	17: uint16(131),
}

var map_token4 = [18]uint16_t{
	0:  uint16('#'),
	1:  uint16(12),
	2:  uint16('.'),
	3:  uint16(15),
	4:  uint16('/'),
	5:  uint16(223),
	6:  uint16('I'),
	7:  uint16(269),
	8:  uint16('i'),
	9:  uint16(269),
	10: uint16('L'),
	11: uint16(131),
	12: uint16('S'),
	13: uint16(131),
	14: uint16('l'),
	15: uint16(131),
	16: uint16('s'),
	17: uint16(131),
}

var map_token5 = [20]uint16_t{
	0:  uint16('U'),
	1:  uint16(413),
	2:  uint16('b'),
	3:  uint16(396),
	4:  uint16('l'),
	5:  uint16(401),
	6:  uint16('n'),
	7:  uint16(399),
	8:  uint16('p'),
	9:  uint16(397),
	10: uint16('r'),
	11: uint16(400),
	12: uint16('s'),
	13: uint16(403),
	14: uint16('t'),
	15: uint16(398),
	16: uint16('u'),
	17: uint16(409),
	18: uint16('v'),
	19: uint16(404),
}

var map_token6 = [16]uint16_t{
	0:  uint16('B'),
	1:  uint16(35),
	2:  uint16('b'),
	3:  uint16(35),
	4:  uint16('D'),
	5:  uint16(32),
	6:  uint16('d'),
	7:  uint16(32),
	8:  uint16('O'),
	9:  uint16(43),
	10: uint16('o'),
	11: uint16(43),
	12: uint16('X'),
	13: uint16(38),
	14: uint16('x'),
	15: uint16(38),
}

var map_token7 = [28]uint16_t{
	0:  uint16('"'),
	1:  uint16(256),
	2:  uint16('#'),
	3:  uint16(250),
	4:  uint16('\''),
	5:  uint16(470),
	6:  uint16('('),
	7:  uint16(456),
	8:  uint16(','),
	9:  uint16(475),
	10: uint16('.'),
	11: uint16(445),
	12: uint16(';'),
	13: uint16(237),
	14: uint16('['),
	15: uint16(458),
	16: uint16('\\'),
	17: uint16(230),
	18: uint16('`'),
	19: uint16(471),
	20: uint16('{'),
	21: uint16(460),
	22: uint16('|'),
	23: uint16(114),
	24: uint16('+'),
	25: uint16(421),
	26: uint16('-'),
	27: uint16(421),
}

var map_token8 = [68]uint16_t{
	0:  uint16('!'),
	1:  uint16(481),
	2:  uint16('%'),
	3:  uint16(453),
	4:  uint16('&'),
	5:  uint16(455),
	6:  uint16('\''),
	7:  uint16(472),
	8:  uint16(','),
	9:  uint16(477),
	10: uint16(':'),
	11: uint16(65),
	12: uint16(';'),
	13: uint16(243),
	14: uint16('<'),
	15: uint16(63),
	16: uint16('\\'),
	17: uint16(64),
	18: uint16('`'),
	19: uint16(473),
	20: uint16('f'),
	21: uint16(247),
	22: uint16('h'),
	23: uint16(66),
	24: uint16('l'),
	25: uint16(67),
	26: uint16('p'),
	27: uint16(113),
	28: uint16('r'),
	29: uint16(82),
	30: uint16('s'),
	31: uint16(464),
	32: uint16('t'),
	33: uint16(248),
	34: uint16('|'),
	35: uint16(238),
	36: uint16('B'),
	37: uint16(7),
	38: uint16('b'),
	39: uint16(7),
	40: uint16('C'),
	41: uint16(194),
	42: uint16('c'),
	43: uint16(194),
	44: uint16('D'),
	45: uint16(18),
	46: uint16('d'),
	47: uint16(18),
	48: uint16('F'),
	49: uint16(246),
	50: uint16('T'),
	51: uint16(246),
	52: uint16('O'),
	53: uint16(23),
	54: uint16('o'),
	55: uint16(23),
	56: uint16('X'),
	57: uint16(28),
	58: uint16('x'),
	59: uint16(28),
	60: uint16('E'),
	61: uint16(8),
	62: uint16('I'),
	63: uint16(8),
	64: uint16('e'),
	65: uint16(8),
	66: uint16('i'),
	67: uint16(8),
}

var map_token9 = [66]uint16_t{
	0:  uint16('!'),
	1:  uint16(116),
	2:  uint16('%'),
	3:  uint16(453),
	4:  uint16('&'),
	5:  uint16(455),
	6:  uint16('\''),
	7:  uint16(472),
	8:  uint16(','),
	9:  uint16(477),
	10: uint16(':'),
	11: uint16(65),
	12: uint16(';'),
	13: uint16(243),
	14: uint16('<'),
	15: uint16(63),
	16: uint16('\\'),
	17: uint16(64),
	18: uint16('`'),
	19: uint16(473),
	20: uint16('f'),
	21: uint16(247),
	22: uint16('h'),
	23: uint16(66),
	24: uint16('p'),
	25: uint16(113),
	26: uint16('r'),
	27: uint16(113),
	28: uint16('s'),
	29: uint16(464),
	30: uint16('t'),
	31: uint16(248),
	32: uint16('|'),
	33: uint16(238),
	34: uint16('B'),
	35: uint16(7),
	36: uint16('b'),
	37: uint16(7),
	38: uint16('C'),
	39: uint16(194),
	40: uint16('c'),
	41: uint16(194),
	42: uint16('D'),
	43: uint16(18),
	44: uint16('d'),
	45: uint16(18),
	46: uint16('F'),
	47: uint16(246),
	48: uint16('T'),
	49: uint16(246),
	50: uint16('O'),
	51: uint16(23),
	52: uint16('o'),
	53: uint16(23),
	54: uint16('X'),
	55: uint16(28),
	56: uint16('x'),
	57: uint16(28),
	58: uint16('E'),
	59: uint16(8),
	60: uint16('I'),
	61: uint16(8),
	62: uint16('e'),
	63: uint16(8),
	64: uint16('i'),
	65: uint16(8),
}

var map_token10 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(271),
	2:  uint16('.'),
	3:  uint16(272),
	4:  uint16('/'),
	5:  uint16(445),
	6:  uint16('@'),
	7:  uint16(423),
	8:  uint16('\\'),
	9:  uint16(230),
	10: uint16('|'),
	11: uint16(114),
	12: uint16('+'),
	13: uint16(425),
	14: uint16('-'),
	15: uint16(425),
	16: uint16('T'),
	17: uint16(426),
	18: uint16('t'),
	19: uint16(426),
}

var map_token11 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(271),
	2:  uint16('.'),
	3:  uint16(273),
	4:  uint16('/'),
	5:  uint16(445),
	6:  uint16('@'),
	7:  uint16(423),
	8:  uint16('\\'),
	9:  uint16(230),
	10: uint16('|'),
	11: uint16(114),
	12: uint16('+'),
	13: uint16(425),
	14: uint16('-'),
	15: uint16(425),
	16: uint16('T'),
	17: uint16(426),
	18: uint16('t'),
	19: uint16(426),
}

var map_token12 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(273),
	2:  uint16('@'),
	3:  uint16(423),
	4:  uint16('\\'),
	5:  uint16(230),
	6:  uint16('|'),
	7:  uint16(114),
	8:  uint16('+'),
	9:  uint16(425),
	10: uint16('-'),
	11: uint16(425),
	12: uint16('T'),
	13: uint16(426),
	14: uint16('t'),
	15: uint16(426),
}

var map_token13 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(273),
	2:  uint16('@'),
	3:  uint16(423),
	4:  uint16('\\'),
	5:  uint16(230),
	6:  uint16('|'),
	7:  uint16(114),
	8:  uint16('+'),
	9:  uint16(425),
	10: uint16('-'),
	11: uint16(425),
	12: uint16('T'),
	13: uint16(426),
	14: uint16('t'),
	15: uint16(426),
}

var map_token14 = [24]uint16_t{
	0:  uint16('#'),
	1:  uint16(275),
	2:  uint16('.'),
	3:  uint16(282),
	4:  uint16('/'),
	5:  uint16(446),
	6:  uint16('@'),
	7:  uint16(423),
	8:  uint16('\\'),
	9:  uint16(230),
	10: uint16('|'),
	11: uint16(114),
	12: uint16('+'),
	13: uint16(425),
	14: uint16('-'),
	15: uint16(425),
	16: uint16('I'),
	17: uint16(380),
	18: uint16('i'),
	19: uint16(380),
	20: uint16('T'),
	21: uint16(426),
	22: uint16('t'),
	23: uint16(426),
}

var map_token15 = [24]uint16_t{
	0:  uint16('#'),
	1:  uint16(275),
	2:  uint16('.'),
	3:  uint16(283),
	4:  uint16('/'),
	5:  uint16(446),
	6:  uint16('@'),
	7:  uint16(423),
	8:  uint16('\\'),
	9:  uint16(230),
	10: uint16('|'),
	11: uint16(114),
	12: uint16('+'),
	13: uint16(425),
	14: uint16('-'),
	15: uint16(425),
	16: uint16('I'),
	17: uint16(380),
	18: uint16('i'),
	19: uint16(380),
	20: uint16('T'),
	21: uint16(426),
	22: uint16('t'),
	23: uint16(426),
}

var map_token16 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(277),
	2:  uint16('.'),
	3:  uint16(286),
	4:  uint16('/'),
	5:  uint16(174),
	6:  uint16('@'),
	7:  uint16(40),
	8:  uint16('+'),
	9:  uint16(46),
	10: uint16('-'),
	11: uint16(46),
	12: uint16('T'),
	13: uint16(117),
	14: uint16('t'),
	15: uint16(117),
	16: uint16('0'),
	17: uint16(276),
	18: uint16('1'),
	19: uint16(276),
}

var map_token17 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(277),
	2:  uint16('.'),
	3:  uint16(287),
	4:  uint16('/'),
	5:  uint16(174),
	6:  uint16('@'),
	7:  uint16(40),
	8:  uint16('+'),
	9:  uint16(46),
	10: uint16('-'),
	11: uint16(46),
	12: uint16('T'),
	13: uint16(117),
	14: uint16('t'),
	15: uint16(117),
}

var map_token18 = [24]uint16_t{
	0:  uint16('#'),
	1:  uint16(279),
	2:  uint16('.'),
	3:  uint16(291),
	4:  uint16('/'),
	5:  uint16(216),
	6:  uint16('@'),
	7:  uint16(44),
	8:  uint16('+'),
	9:  uint16(51),
	10: uint16('-'),
	11: uint16(51),
	12: uint16('T'),
	13: uint16(118),
	14: uint16('t'),
	15: uint16(118),
	16: uint16('L'),
	17: uint16(122),
	18: uint16('S'),
	19: uint16(122),
	20: uint16('l'),
	21: uint16(122),
	22: uint16('s'),
	23: uint16(122),
}

var map_token19 = [24]uint16_t{
	0:  uint16('#'),
	1:  uint16(279),
	2:  uint16('.'),
	3:  uint16(290),
	4:  uint16('/'),
	5:  uint16(216),
	6:  uint16('@'),
	7:  uint16(44),
	8:  uint16('+'),
	9:  uint16(51),
	10: uint16('-'),
	11: uint16(51),
	12: uint16('T'),
	13: uint16(118),
	14: uint16('t'),
	15: uint16(118),
	16: uint16('L'),
	17: uint16(122),
	18: uint16('S'),
	19: uint16(122),
	20: uint16('l'),
	21: uint16(122),
	22: uint16('s'),
	23: uint16(122),
}

var map_token20 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(283),
	2:  uint16('@'),
	3:  uint16(423),
	4:  uint16('\\'),
	5:  uint16(230),
	6:  uint16('|'),
	7:  uint16(114),
	8:  uint16('+'),
	9:  uint16(425),
	10: uint16('-'),
	11: uint16(425),
	12: uint16('I'),
	13: uint16(380),
	14: uint16('i'),
	15: uint16(380),
	16: uint16('T'),
	17: uint16(426),
	18: uint16('t'),
	19: uint16(426),
}

var map_token21 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(283),
	2:  uint16('@'),
	3:  uint16(423),
	4:  uint16('\\'),
	5:  uint16(230),
	6:  uint16('|'),
	7:  uint16(114),
	8:  uint16('+'),
	9:  uint16(425),
	10: uint16('-'),
	11: uint16(425),
	12: uint16('I'),
	13: uint16(380),
	14: uint16('i'),
	15: uint16(380),
	16: uint16('T'),
	17: uint16(426),
	18: uint16('t'),
	19: uint16(426),
}

var map_token22 = [24]uint16_t{
	0:  uint16('#'),
	1:  uint16(289),
	2:  uint16('.'),
	3:  uint16(300),
	4:  uint16('/'),
	5:  uint16(175),
	6:  uint16('@'),
	7:  uint16(40),
	8:  uint16('+'),
	9:  uint16(46),
	10: uint16('-'),
	11: uint16(46),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
	16: uint16('T'),
	17: uint16(117),
	18: uint16('t'),
	19: uint16(117),
	20: uint16('0'),
	21: uint16(288),
	22: uint16('1'),
	23: uint16(288),
}

var map_token23 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(289),
	2:  uint16('.'),
	3:  uint16(301),
	4:  uint16('/'),
	5:  uint16(175),
	6:  uint16('@'),
	7:  uint16(40),
	8:  uint16('+'),
	9:  uint16(46),
	10: uint16('-'),
	11: uint16(46),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
	16: uint16('T'),
	17: uint16(117),
	18: uint16('t'),
	19: uint16(117),
}

var map_token24 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(290),
	2:  uint16('@'),
	3:  uint16(44),
	4:  uint16('+'),
	5:  uint16(51),
	6:  uint16('-'),
	7:  uint16(51),
	8:  uint16('T'),
	9:  uint16(118),
	10: uint16('t'),
	11: uint16(118),
	12: uint16('L'),
	13: uint16(122),
	14: uint16('S'),
	15: uint16(122),
	16: uint16('l'),
	17: uint16(122),
	18: uint16('s'),
	19: uint16(122),
}

var map_token25 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(290),
	2:  uint16('@'),
	3:  uint16(44),
	4:  uint16('+'),
	5:  uint16(51),
	6:  uint16('-'),
	7:  uint16(51),
	8:  uint16('T'),
	9:  uint16(118),
	10: uint16('t'),
	11: uint16(118),
	12: uint16('L'),
	13: uint16(122),
	14: uint16('S'),
	15: uint16(122),
	16: uint16('l'),
	17: uint16(122),
	18: uint16('s'),
	19: uint16(122),
}

var map_token26 = [28]uint16_t{
	0:  uint16('#'),
	1:  uint16(293),
	2:  uint16('.'),
	3:  uint16(307),
	4:  uint16('/'),
	5:  uint16(217),
	6:  uint16('@'),
	7:  uint16(44),
	8:  uint16('+'),
	9:  uint16(51),
	10: uint16('-'),
	11: uint16(51),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
	16: uint16('T'),
	17: uint16(118),
	18: uint16('t'),
	19: uint16(118),
	20: uint16('L'),
	21: uint16(127),
	22: uint16('S'),
	23: uint16(127),
	24: uint16('l'),
	25: uint16(127),
	26: uint16('s'),
	27: uint16(127),
}

var map_token27 = [28]uint16_t{
	0:  uint16('#'),
	1:  uint16(293),
	2:  uint16('.'),
	3:  uint16(306),
	4:  uint16('/'),
	5:  uint16(217),
	6:  uint16('@'),
	7:  uint16(44),
	8:  uint16('+'),
	9:  uint16(51),
	10: uint16('-'),
	11: uint16(51),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
	16: uint16('T'),
	17: uint16(118),
	18: uint16('t'),
	19: uint16(118),
	20: uint16('L'),
	21: uint16(127),
	22: uint16('S'),
	23: uint16(127),
	24: uint16('l'),
	25: uint16(127),
	26: uint16('s'),
	27: uint16(127),
}

var map_token28 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(297),
	2:  uint16('.'),
	3:  uint16(311),
	4:  uint16('/'),
	5:  uint16(211),
	6:  uint16('@'),
	7:  uint16(47),
	8:  uint16('+'),
	9:  uint16(52),
	10: uint16('-'),
	11: uint16(52),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
}

var map_token29 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(297),
	2:  uint16('.'),
	3:  uint16(310),
	4:  uint16('/'),
	5:  uint16(211),
	6:  uint16('@'),
	7:  uint16(47),
	8:  uint16('+'),
	9:  uint16(52),
	10: uint16('-'),
	11: uint16(52),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
}

var map_token30 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(301),
	2:  uint16('@'),
	3:  uint16(40),
	4:  uint16('+'),
	5:  uint16(46),
	6:  uint16('-'),
	7:  uint16(46),
	8:  uint16('I'),
	9:  uint16(269),
	10: uint16('i'),
	11: uint16(269),
	12: uint16('T'),
	13: uint16(117),
	14: uint16('t'),
	15: uint16(117),
	16: uint16('0'),
	17: uint16(300),
	18: uint16('1'),
	19: uint16(300),
}

var map_token31 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(301),
	2:  uint16('@'),
	3:  uint16(40),
	4:  uint16('+'),
	5:  uint16(46),
	6:  uint16('-'),
	7:  uint16(46),
	8:  uint16('I'),
	9:  uint16(269),
	10: uint16('i'),
	11: uint16(269),
	12: uint16('T'),
	13: uint16(117),
	14: uint16('t'),
	15: uint16(117),
}

var map_token32 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(305),
	2:  uint16('.'),
	3:  uint16(315),
	4:  uint16('/'),
	5:  uint16(220),
	6:  uint16('@'),
	7:  uint16(44),
	8:  uint16('+'),
	9:  uint16(51),
	10: uint16('-'),
	11: uint16(51),
	12: uint16('L'),
	13: uint16(122),
	14: uint16('S'),
	15: uint16(122),
	16: uint16('l'),
	17: uint16(122),
	18: uint16('s'),
	19: uint16(122),
}

var map_token33 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(305),
	2:  uint16('.'),
	3:  uint16(314),
	4:  uint16('/'),
	5:  uint16(220),
	6:  uint16('@'),
	7:  uint16(44),
	8:  uint16('+'),
	9:  uint16(51),
	10: uint16('-'),
	11: uint16(51),
	12: uint16('L'),
	13: uint16(122),
	14: uint16('S'),
	15: uint16(122),
	16: uint16('l'),
	17: uint16(122),
	18: uint16('s'),
	19: uint16(122),
}

var map_token34 = [24]uint16_t{
	0:  uint16('#'),
	1:  uint16(306),
	2:  uint16('@'),
	3:  uint16(44),
	4:  uint16('+'),
	5:  uint16(51),
	6:  uint16('-'),
	7:  uint16(51),
	8:  uint16('I'),
	9:  uint16(269),
	10: uint16('i'),
	11: uint16(269),
	12: uint16('T'),
	13: uint16(118),
	14: uint16('t'),
	15: uint16(118),
	16: uint16('L'),
	17: uint16(127),
	18: uint16('S'),
	19: uint16(127),
	20: uint16('l'),
	21: uint16(127),
	22: uint16('s'),
	23: uint16(127),
}

var map_token35 = [24]uint16_t{
	0:  uint16('#'),
	1:  uint16(306),
	2:  uint16('@'),
	3:  uint16(44),
	4:  uint16('+'),
	5:  uint16(51),
	6:  uint16('-'),
	7:  uint16(51),
	8:  uint16('I'),
	9:  uint16(269),
	10: uint16('i'),
	11: uint16(269),
	12: uint16('T'),
	13: uint16(118),
	14: uint16('t'),
	15: uint16(118),
	16: uint16('L'),
	17: uint16(127),
	18: uint16('S'),
	19: uint16(127),
	20: uint16('l'),
	21: uint16(127),
	22: uint16('s'),
	23: uint16(127),
}

var map_token36 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(314),
	2:  uint16('@'),
	3:  uint16(44),
	4:  uint16('+'),
	5:  uint16(51),
	6:  uint16('-'),
	7:  uint16(51),
	8:  uint16('L'),
	9:  uint16(122),
	10: uint16('S'),
	11: uint16(122),
	12: uint16('l'),
	13: uint16(122),
	14: uint16('s'),
	15: uint16(122),
}

var map_token37 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(314),
	2:  uint16('@'),
	3:  uint16(44),
	4:  uint16('+'),
	5:  uint16(51),
	6:  uint16('-'),
	7:  uint16(51),
	8:  uint16('L'),
	9:  uint16(122),
	10: uint16('S'),
	11: uint16(122),
	12: uint16('l'),
	13: uint16(122),
	14: uint16('s'),
	15: uint16(122),
}

var map_token38 = [24]uint16_t{
	0:  uint16('#'),
	1:  uint16(317),
	2:  uint16('.'),
	3:  uint16(321),
	4:  uint16('/'),
	5:  uint16(224),
	6:  uint16('@'),
	7:  uint16(44),
	8:  uint16('+'),
	9:  uint16(51),
	10: uint16('-'),
	11: uint16(51),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
	16: uint16('L'),
	17: uint16(127),
	18: uint16('S'),
	19: uint16(127),
	20: uint16('l'),
	21: uint16(127),
	22: uint16('s'),
	23: uint16(127),
}

var map_token39 = [24]uint16_t{
	0:  uint16('#'),
	1:  uint16(317),
	2:  uint16('.'),
	3:  uint16(320),
	4:  uint16('/'),
	5:  uint16(224),
	6:  uint16('@'),
	7:  uint16(44),
	8:  uint16('+'),
	9:  uint16(51),
	10: uint16('-'),
	11: uint16(51),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
	16: uint16('L'),
	17: uint16(127),
	18: uint16('S'),
	19: uint16(127),
	20: uint16('l'),
	21: uint16(127),
	22: uint16('s'),
	23: uint16(127),
}

var map_token40 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(320),
	2:  uint16('@'),
	3:  uint16(44),
	4:  uint16('+'),
	5:  uint16(51),
	6:  uint16('-'),
	7:  uint16(51),
	8:  uint16('I'),
	9:  uint16(269),
	10: uint16('i'),
	11: uint16(269),
	12: uint16('L'),
	13: uint16(127),
	14: uint16('S'),
	15: uint16(127),
	16: uint16('l'),
	17: uint16(127),
	18: uint16('s'),
	19: uint16(127),
}

var map_token41 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(320),
	2:  uint16('@'),
	3:  uint16(44),
	4:  uint16('+'),
	5:  uint16(51),
	6:  uint16('-'),
	7:  uint16(51),
	8:  uint16('I'),
	9:  uint16(269),
	10: uint16('i'),
	11: uint16(269),
	12: uint16('L'),
	13: uint16(127),
	14: uint16('S'),
	15: uint16(127),
	16: uint16('l'),
	17: uint16(127),
	18: uint16('s'),
	19: uint16(127),
}

var map_token42 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(323),
	2:  uint16('.'),
	3:  uint16(327),
	4:  uint16('/'),
	5:  uint16(206),
	6:  uint16('@'),
	7:  uint16(47),
	8:  uint16('+'),
	9:  uint16(52),
	10: uint16('-'),
	11: uint16(52),
	12: uint16('T'),
	13: uint16(120),
	14: uint16('t'),
	15: uint16(120),
}

var map_token43 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(323),
	2:  uint16('.'),
	3:  uint16(326),
	4:  uint16('/'),
	5:  uint16(206),
	6:  uint16('@'),
	7:  uint16(47),
	8:  uint16('+'),
	9:  uint16(52),
	10: uint16('-'),
	11: uint16(52),
	12: uint16('T'),
	13: uint16(120),
	14: uint16('t'),
	15: uint16(120),
}

var map_token44 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(329),
	2:  uint16('.'),
	3:  uint16(335),
	4:  uint16('/'),
	5:  uint16(208),
	6:  uint16('@'),
	7:  uint16(47),
	8:  uint16('+'),
	9:  uint16(52),
	10: uint16('-'),
	11: uint16(52),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
	16: uint16('T'),
	17: uint16(120),
	18: uint16('t'),
	19: uint16(120),
}

var map_token45 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(329),
	2:  uint16('.'),
	3:  uint16(334),
	4:  uint16('/'),
	5:  uint16(208),
	6:  uint16('@'),
	7:  uint16(47),
	8:  uint16('+'),
	9:  uint16(52),
	10: uint16('-'),
	11: uint16(52),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
	16: uint16('T'),
	17: uint16(120),
	18: uint16('t'),
	19: uint16(120),
}

var map_token46 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(333),
	2:  uint16('.'),
	3:  uint16(338),
	4:  uint16('/'),
	5:  uint16(182),
	6:  uint16('@'),
	7:  uint16(40),
	8:  uint16('+'),
	9:  uint16(46),
	10: uint16('-'),
	11: uint16(46),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
	16: uint16('0'),
	17: uint16(332),
	18: uint16('1'),
	19: uint16(332),
}

var map_token47 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(333),
	2:  uint16('.'),
	3:  uint16(339),
	4:  uint16('/'),
	5:  uint16(182),
	6:  uint16('@'),
	7:  uint16(40),
	8:  uint16('+'),
	9:  uint16(46),
	10: uint16('-'),
	11: uint16(46),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
}

var map_token48 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(334),
	2:  uint16('@'),
	3:  uint16(47),
	4:  uint16('+'),
	5:  uint16(52),
	6:  uint16('-'),
	7:  uint16(52),
	8:  uint16('I'),
	9:  uint16(269),
	10: uint16('i'),
	11: uint16(269),
	12: uint16('T'),
	13: uint16(120),
	14: uint16('t'),
	15: uint16(120),
}

var map_token49 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(334),
	2:  uint16('@'),
	3:  uint16(47),
	4:  uint16('+'),
	5:  uint16(52),
	6:  uint16('-'),
	7:  uint16(52),
	8:  uint16('I'),
	9:  uint16(269),
	10: uint16('i'),
	11: uint16(269),
	12: uint16('T'),
	13: uint16(120),
	14: uint16('t'),
	15: uint16(120),
}

var map_token50 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(343),
	2:  uint16('.'),
	3:  uint16(346),
	4:  uint16('/'),
	5:  uint16(196),
	6:  uint16('@'),
	7:  uint16(49),
	8:  uint16('+'),
	9:  uint16(53),
	10: uint16('-'),
	11: uint16(53),
	12: uint16('T'),
	13: uint16(121),
	14: uint16('t'),
	15: uint16(121),
}

var map_token51 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(343),
	2:  uint16('.'),
	3:  uint16(347),
	4:  uint16('/'),
	5:  uint16(196),
	6:  uint16('@'),
	7:  uint16(49),
	8:  uint16('+'),
	9:  uint16(53),
	10: uint16('-'),
	11: uint16(53),
	12: uint16('T'),
	13: uint16(121),
	14: uint16('t'),
	15: uint16(121),
}

var map_token52 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(349),
	2:  uint16('.'),
	3:  uint16(354),
	4:  uint16('/'),
	5:  uint16(197),
	6:  uint16('@'),
	7:  uint16(49),
	8:  uint16('+'),
	9:  uint16(53),
	10: uint16('-'),
	11: uint16(53),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
	16: uint16('T'),
	17: uint16(121),
	18: uint16('t'),
	19: uint16(121),
}

var map_token53 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(349),
	2:  uint16('.'),
	3:  uint16(355),
	4:  uint16('/'),
	5:  uint16(197),
	6:  uint16('@'),
	7:  uint16(49),
	8:  uint16('+'),
	9:  uint16(53),
	10: uint16('-'),
	11: uint16(53),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
	16: uint16('T'),
	17: uint16(121),
	18: uint16('t'),
	19: uint16(121),
}

var map_token54 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(353),
	2:  uint16('.'),
	3:  uint16(358),
	4:  uint16('/'),
	5:  uint16(204),
	6:  uint16('@'),
	7:  uint16(49),
	8:  uint16('+'),
	9:  uint16(53),
	10: uint16('-'),
	11: uint16(53),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
}

var map_token55 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(353),
	2:  uint16('.'),
	3:  uint16(359),
	4:  uint16('/'),
	5:  uint16(204),
	6:  uint16('@'),
	7:  uint16(49),
	8:  uint16('+'),
	9:  uint16(53),
	10: uint16('-'),
	11: uint16(53),
	12: uint16('I'),
	13: uint16(269),
	14: uint16('i'),
	15: uint16(269),
}

var map_token56 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(355),
	2:  uint16('@'),
	3:  uint16(49),
	4:  uint16('+'),
	5:  uint16(53),
	6:  uint16('-'),
	7:  uint16(53),
	8:  uint16('I'),
	9:  uint16(269),
	10: uint16('i'),
	11: uint16(269),
	12: uint16('T'),
	13: uint16(121),
	14: uint16('t'),
	15: uint16(121),
}

var map_token57 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(355),
	2:  uint16('@'),
	3:  uint16(49),
	4:  uint16('+'),
	5:  uint16(53),
	6:  uint16('-'),
	7:  uint16(53),
	8:  uint16('I'),
	9:  uint16(269),
	10: uint16('i'),
	11: uint16(269),
	12: uint16('T'),
	13: uint16(121),
	14: uint16('t'),
	15: uint16(121),
}

var ts_lex_modes = [257]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state: uint16(232),
	},
	2: {
		Flex_state: uint16(3),
	},
	3: {
		Flex_state: uint16(3),
	},
	4: {
		Flex_state: uint16(3),
	},
	5: {
		Flex_state: uint16(3),
	},
	6: {
		Flex_state: uint16(3),
	},
	7: {
		Flex_state: uint16(3),
	},
	8: {
		Flex_state: uint16(3),
	},
	9: {
		Flex_state: uint16(3),
	},
	10: {
		Flex_state: uint16(3),
	},
	11: {
		Flex_state: uint16(3),
	},
	12: {
		Flex_state: uint16(3),
	},
	13: {
		Flex_state: uint16(3),
	},
	14: {
		Flex_state: uint16(3),
	},
	15: {
		Flex_state: uint16(3),
	},
	16: {
		Flex_state: uint16(3),
	},
	17: {
		Flex_state: uint16(3),
	},
	18: {
		Flex_state: uint16(3),
	},
	19: {
		Flex_state: uint16(3),
	},
	20: {
		Flex_state: uint16(3),
	},
	21: {
		Flex_state: uint16(232),
	},
	22: {
		Flex_state: uint16(232),
	},
	23: {
		Flex_state: uint16(4),
	},
	24: {
		Flex_state: uint16(4),
	},
	25: {
		Flex_state: uint16(4),
	},
	26: {
		Flex_state: uint16(4),
	},
	27: {
		Flex_state: uint16(4),
	},
	28: {
		Flex_state: uint16(4),
	},
	29: {
		Flex_state: uint16(4),
	},
	30: {
		Flex_state: uint16(4),
	},
	31: {
		Flex_state: uint16(4),
	},
	32: {
		Flex_state: uint16(4),
	},
	33: {
		Flex_state: uint16(4),
	},
	34: {
		Flex_state: uint16(4),
	},
	35: {
		Flex_state: uint16(4),
	},
	36: {
		Flex_state: uint16(4),
	},
	37: {
		Flex_state: uint16(4),
	},
	38: {
		Flex_state: uint16(4),
	},
	39: {
		Flex_state: uint16(4),
	},
	40: {
		Flex_state: uint16(4),
	},
	41: {
		Flex_state: uint16(4),
	},
	42: {
		Flex_state: uint16(4),
	},
	43: {
		Flex_state: uint16(4),
	},
	44: {
		Flex_state: uint16(4),
	},
	45: {
		Flex_state: uint16(4),
	},
	46: {
		Flex_state: uint16(4),
	},
	47: {
		Flex_state: uint16(4),
	},
	48: {
		Flex_state: uint16(4),
	},
	49: {
		Flex_state: uint16(4),
	},
	50: {
		Flex_state: uint16(4),
	},
	51: {
		Flex_state: uint16(4),
	},
	52: {
		Flex_state: uint16(4),
	},
	53: {
		Flex_state: uint16(4),
	},
	54: {
		Flex_state: uint16(4),
	},
	55: {
		Flex_state: uint16(4),
	},
	56: {
		Flex_state: uint16(4),
	},
	57: {
		Flex_state: uint16(4),
	},
	58: {
		Flex_state: uint16(4),
	},
	59: {
		Flex_state: uint16(4),
	},
	60: {
		Flex_state: uint16(4),
	},
	61: {
		Flex_state: uint16(4),
	},
	62: {
		Flex_state: uint16(4),
	},
	63: {
		Flex_state: uint16(4),
	},
	64: {
		Flex_state: uint16(4),
	},
	65: {
		Flex_state: uint16(4),
	},
	66: {
		Flex_state: uint16(4),
	},
	67: {
		Flex_state: uint16(4),
	},
	68: {
		Flex_state: uint16(4),
	},
	69: {
		Flex_state: uint16(4),
	},
	70: {
		Flex_state: uint16(4),
	},
	71: {
		Flex_state: uint16(4),
	},
	72: {
		Flex_state: uint16(4),
	},
	73: {
		Flex_state: uint16(4),
	},
	74: {
		Flex_state: uint16(4),
	},
	75: {
		Flex_state: uint16(4),
	},
	76: {
		Flex_state: uint16(4),
	},
	77: {
		Flex_state: uint16(4),
	},
	78: {
		Flex_state: uint16(4),
	},
	79: {
		Flex_state: uint16(4),
	},
	80: {
		Flex_state: uint16(4),
	},
	81: {
		Flex_state: uint16(4),
	},
	82: {
		Flex_state: uint16(4),
	},
	83: {
		Flex_state: uint16(4),
	},
	84: {
		Flex_state: uint16(4),
	},
	85: {
		Flex_state: uint16(4),
	},
	86: {
		Flex_state: uint16(4),
	},
	87: {
		Flex_state: uint16(4),
	},
	88: {
		Flex_state: uint16(4),
	},
	89: {
		Flex_state: uint16(4),
	},
	90: {
		Flex_state: uint16(4),
	},
	91: {
		Flex_state: uint16(4),
	},
	92: {
		Flex_state: uint16(4),
	},
	93: {
		Flex_state: uint16(3),
	},
	94: {
		Flex_state: uint16(3),
	},
	95: {
		Flex_state: uint16(3),
	},
	96: {
		Flex_state: uint16(3),
	},
	97: {
		Flex_state: uint16(3),
	},
	98: {
		Flex_state: uint16(3),
	},
	99: {
		Flex_state: uint16(3),
	},
	100: {
		Flex_state: uint16(3),
	},
	101: {
		Flex_state: uint16(3),
	},
	102: {
		Flex_state: uint16(3),
	},
	103: {
		Flex_state: uint16(3),
	},
	104: {
		Flex_state: uint16(3),
	},
	105: {
		Flex_state: uint16(3),
	},
	106: {
		Flex_state: uint16(3),
	},
	107: {
		Flex_state: uint16(3),
	},
	108: {
		Flex_state: uint16(3),
	},
	109: {
		Flex_state: uint16(3),
	},
	110: {
		Flex_state: uint16(3),
	},
	111: {
		Flex_state: uint16(3),
	},
	112: {
		Flex_state: uint16(3),
	},
	113: {
		Flex_state: uint16(3),
	},
	114: {
		Flex_state: uint16(3),
	},
	115: {
		Flex_state: uint16(3),
	},
	116: {
		Flex_state: uint16(3),
	},
	117: {
		Flex_state: uint16(3),
	},
	118: {
		Flex_state: uint16(3),
	},
	119: {
		Flex_state: uint16(3),
	},
	120: {
		Flex_state: uint16(3),
	},
	121: {
		Flex_state: uint16(3),
	},
	122: {
		Flex_state: uint16(3),
	},
	123: {
		Flex_state: uint16(3),
	},
	124: {
		Flex_state: uint16(3),
	},
	125: {
		Flex_state: uint16(3),
	},
	126: {
		Flex_state: uint16(3),
	},
	127: {
		Flex_state: uint16(3),
	},
	128: {
		Flex_state: uint16(3),
	},
	129: {
		Flex_state: uint16(3),
	},
	130: {
		Flex_state: uint16(3),
	},
	131: {
		Flex_state: uint16(3),
	},
	132: {
		Flex_state: uint16(3),
	},
	133: {
		Flex_state: uint16(4),
	},
	134: {
		Flex_state: uint16(232),
	},
	135: {
		Flex_state: uint16(232),
	},
	136: {
		Flex_state: uint16(232),
	},
	137: {
		Flex_state: uint16(232),
	},
	138: {
		Flex_state: uint16(232),
	},
	139: {
		Flex_state: uint16(232),
	},
	140: {
		Flex_state: uint16(232),
	},
	141: {
		Flex_state: uint16(232),
	},
	142: {
		Flex_state: uint16(232),
	},
	143: {
		Flex_state: uint16(232),
	},
	144: {
		Flex_state: uint16(232),
	},
	145: {
		Flex_state: uint16(232),
	},
	146: {
		Flex_state: uint16(232),
	},
	147: {
		Flex_state: uint16(232),
	},
	148: {
		Flex_state: uint16(232),
	},
	149: {
		Flex_state: uint16(232),
	},
	150: {
		Flex_state: uint16(232),
	},
	151: {
		Flex_state: uint16(232),
	},
	152: {
		Flex_state: uint16(232),
	},
	153: {
		Flex_state: uint16(232),
	},
	154: {
		Flex_state: uint16(232),
	},
	155: {
		Flex_state: uint16(232),
	},
	156: {
		Flex_state: uint16(232),
	},
	157: {
		Flex_state: uint16(232),
	},
	158: {
		Flex_state: uint16(232),
	},
	159: {
		Flex_state: uint16(232),
	},
	160: {
		Flex_state: uint16(232),
	},
	161: {
		Flex_state: uint16(232),
	},
	162: {
		Flex_state: uint16(232),
	},
	163: {
		Flex_state: uint16(232),
	},
	164: {
		Flex_state: uint16(232),
	},
	165: {
		Flex_state: uint16(232),
	},
	166: {
		Flex_state: uint16(232),
	},
	167: {
		Flex_state: uint16(232),
	},
	168: {
		Flex_state: uint16(232),
	},
	169: {
		Flex_state: uint16(232),
	},
	170: {
		Flex_state: uint16(232),
	},
	171: {
		Flex_state: uint16(232),
	},
	172: {
		Flex_state: uint16(232),
	},
	173: {
		Flex_state: uint16(232),
	},
	174: {
		Flex_state: uint16(4),
	},
	175: {
		Flex_state: uint16(4),
	},
	176: {
		Flex_state: uint16(4),
	},
	177: {
		Flex_state: uint16(4),
	},
	178: {
		Flex_state: uint16(4),
	},
	179: {
		Flex_state: uint16(4),
	},
	180: {
		Flex_state: uint16(4),
	},
	181: {
		Flex_state: uint16(4),
	},
	182: {
		Flex_state: uint16(4),
	},
	183: {
		Flex_state: uint16(4),
	},
	184: {
		Flex_state: uint16(4),
	},
	185: {
		Flex_state: uint16(4),
	},
	186: {
		Flex_state: uint16(4),
	},
	187: {
		Flex_state: uint16(4),
	},
	188: {
		Flex_state: uint16(4),
	},
	189: {
		Flex_state: uint16(4),
	},
	190: {
		Flex_state: uint16(4),
	},
	191: {
		Flex_state: uint16(4),
	},
	192: {
		Flex_state: uint16(4),
	},
	193: {
		Flex_state: uint16(4),
	},
	194: {
		Flex_state: uint16(4),
	},
	195: {
		Flex_state: uint16(4),
	},
	196: {
		Flex_state: uint16(4),
	},
	197: {
		Flex_state: uint16(4),
	},
	198: {
		Flex_state: uint16(4),
	},
	199: {
		Flex_state: uint16(4),
	},
	200: {
		Flex_state: uint16(4),
	},
	201: {
		Flex_state: uint16(4),
	},
	202: {
		Flex_state: uint16(4),
	},
	203: {
		Flex_state: uint16(4),
	},
	204: {
		Flex_state: uint16(4),
	},
	205: {
		Flex_state: uint16(4),
	},
	206: {
		Flex_state: uint16(4),
	},
	207: {
		Flex_state: uint16(4),
	},
	208: {
		Flex_state: uint16(4),
	},
	209: {
		Flex_state: uint16(4),
	},
	210: {
		Flex_state: uint16(4),
	},
	211: {
		Flex_state: uint16(4),
	},
	212: {
		Flex_state: uint16(5),
	},
	213: {
		Flex_state: uint16(5),
	},
	214: {
		Flex_state: uint16(5),
	},
	215: {
		Flex_state: uint16(5),
	},
	216: {
		Flex_state: uint16(5),
	},
	217: {
		Flex_state: uint16(5),
	},
	218: {
		Flex_state: uint16(17),
	},
	219: {
		Flex_state: uint16(17),
	},
	220: {
		Flex_state: uint16(17),
	},
	221: {
		Flex_state: uint16(5),
	},
	222: {
		Flex_state: uint16(17),
	},
	223: {
		Flex_state: uint16(17),
	},
	224: {
		Flex_state: uint16(5),
	},
	225: {
		Flex_state: uint16(17),
	},
	226: {
		Flex_state: uint16(17),
	},
	227: {
		Flex_state: uint16(5),
	},
	228: {
		Flex_state: uint16(17),
	},
	229: {
		Flex_state: uint16(17),
	},
	230: {
		Flex_state: uint16(232),
	},
	231: {
		Flex_state: uint16(6),
	},
	232: {
		Flex_state: uint16(6),
	},
	233: {
		Flex_state: uint16(232),
	},
	234: {
		Flex_state: uint16(232),
	},
	235: {
		Flex_state: uint16(232),
	},
	236: {
		Flex_state: uint16(232),
	},
	237: {
		Flex_state: uint16(6),
	},
	238: {
		Flex_state: uint16(6),
	},
	239: {
		Flex_state: uint16(232),
	},
	240: {
		Flex_state: uint16(232),
	},
	241: {
		Flex_state: uint16(232),
	},
	242: {
		Flex_state: uint16(6),
	},
	243: {
		Flex_state: uint16(232),
	},
	244: {
		Flex_state: uint16(6),
	},
	245: {
		Flex_state: uint16(6),
	},
	246: {
		Flex_state: uint16(17),
	},
	247: {
		Flex_state: uint16(17),
	},
	248: {
		Flex_state: uint16(232),
	},
	249: {
		Flex_state: uint16(232),
	},
	250: {
		Flex_state: uint16(232),
	},
	251: {
		Fexternal_lex_state: uint16(1),
	},
	252: {
		Flex_state: uint16(229),
	},
	253: {
		Fexternal_lex_state: uint16(1),
	},
	254: {},
	255: {
		Flex_state: uint16(229),
	},
	256: {
		Fexternal_lex_state: uint16(1),
	},
}

var ts_parse_table = [93][81]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		3:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		6:  uint16(1),
		10: uint16(1),
		13: uint16(1),
		22: uint16(1),
		23: uint16(1),
		24: uint16(1),
		25: uint16(1),
		26: uint16(1),
		27: uint16(1),
		35: uint16(1),
		36: uint16(1),
		37: uint16(1),
		40: uint16(1),
		48: uint16(1),
	},
	1: {
		0:  uint16(3),
		1:  uint16(5),
		3:  uint16(7),
		4:  uint16(9),
		7:  uint16(11),
		8:  uint16(7),
		9:  uint16(13),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(13),
		18: uint16(5),
		19: uint16(13),
		20: uint16(5),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		44: uint16(55),
		45: uint16(57),
		46: uint16(59),
		49: uint16(254),
		50: uint16(21),
		51: uint16(21),
		52: uint16(21),
		53: uint16(21),
		54: uint16(21),
		55: uint16(21),
		56: uint16(21),
		57: uint16(21),
		58: uint16(21),
		59: uint16(21),
		60: uint16(143),
		61: uint16(21),
		62: uint16(21),
		63: uint16(21),
		64: uint16(21),
		65: uint16(21),
		66: uint16(21),
		67: uint16(21),
		68: uint16(21),
		69: uint16(21),
		70: uint16(21),
		71: uint16(21),
		72: uint16(21),
		73: uint16(21),
		74: uint16(21),
		75: uint16(21),
		76: uint16(21),
	},
	2: {
		1:  uint16(61),
		2:  uint16(64),
		3:  uint16(67),
		4:  uint16(70),
		7:  uint16(73),
		8:  uint16(67),
		9:  uint16(64),
		10: uint16(76),
		11: uint16(79),
		12: uint16(82),
		13: uint16(85),
		16: uint16(64),
		18: uint16(61),
		19: uint16(64),
		20: uint16(61),
		21: uint16(88),
		22: uint16(91),
		23: uint16(94),
		24: uint16(96),
		25: uint16(94),
		26: uint16(99),
		27: uint16(94),
		28: uint16(102),
		29: uint16(102),
		30: uint16(105),
		31: uint16(108),
		32: uint16(111),
		33: uint16(108),
		34: uint16(111),
		36: uint16(114),
		37: uint16(117),
		38: uint16(120),
		39: uint16(123),
		40: uint16(126),
		41: uint16(129),
		42: uint16(132),
		43: uint16(135),
		44: uint16(138),
		45: uint16(141),
		46: uint16(144),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
		53: uint16(2),
		54: uint16(2),
		55: uint16(2),
		56: uint16(2),
		57: uint16(2),
		58: uint16(2),
		59: uint16(2),
		60: uint16(114),
		61: uint16(2),
		62: uint16(2),
		63: uint16(2),
		64: uint16(2),
		65: uint16(2),
		66: uint16(2),
		67: uint16(2),
		68: uint16(2),
		69: uint16(2),
		70: uint16(2),
		71: uint16(2),
		72: uint16(2),
		73: uint16(2),
		74: uint16(2),
		75: uint16(2),
		80: uint16(2),
	},
	3: {
		1:  uint16(147),
		2:  uint16(149),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(149),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(149),
		18: uint16(147),
		19: uint16(149),
		20: uint16(147),
		21: uint16(165),
		22: uint16(167),
		23: uint16(169),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(5),
		51: uint16(5),
		52: uint16(5),
		53: uint16(5),
		54: uint16(5),
		55: uint16(5),
		56: uint16(5),
		57: uint16(5),
		58: uint16(5),
		59: uint16(5),
		60: uint16(114),
		61: uint16(5),
		62: uint16(5),
		63: uint16(5),
		64: uint16(5),
		65: uint16(5),
		66: uint16(5),
		67: uint16(5),
		68: uint16(5),
		69: uint16(5),
		70: uint16(5),
		71: uint16(5),
		72: uint16(5),
		73: uint16(5),
		74: uint16(5),
		75: uint16(5),
		80: uint16(5),
	},
	4: {
		1:  uint16(205),
		2:  uint16(207),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(207),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(207),
		18: uint16(205),
		19: uint16(207),
		20: uint16(205),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		25: uint16(169),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(6),
		51: uint16(6),
		52: uint16(6),
		53: uint16(6),
		54: uint16(6),
		55: uint16(6),
		56: uint16(6),
		57: uint16(6),
		58: uint16(6),
		59: uint16(6),
		60: uint16(114),
		61: uint16(6),
		62: uint16(6),
		63: uint16(6),
		64: uint16(6),
		65: uint16(6),
		66: uint16(6),
		67: uint16(6),
		68: uint16(6),
		69: uint16(6),
		70: uint16(6),
		71: uint16(6),
		72: uint16(6),
		73: uint16(6),
		74: uint16(6),
		75: uint16(6),
		80: uint16(6),
	},
	5: {
		1:  uint16(209),
		2:  uint16(211),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(211),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(211),
		18: uint16(209),
		19: uint16(211),
		20: uint16(209),
		21: uint16(165),
		22: uint16(167),
		23: uint16(213),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
		53: uint16(2),
		54: uint16(2),
		55: uint16(2),
		56: uint16(2),
		57: uint16(2),
		58: uint16(2),
		59: uint16(2),
		60: uint16(114),
		61: uint16(2),
		62: uint16(2),
		63: uint16(2),
		64: uint16(2),
		65: uint16(2),
		66: uint16(2),
		67: uint16(2),
		68: uint16(2),
		69: uint16(2),
		70: uint16(2),
		71: uint16(2),
		72: uint16(2),
		73: uint16(2),
		74: uint16(2),
		75: uint16(2),
		80: uint16(2),
	},
	6: {
		1:  uint16(209),
		2:  uint16(211),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(211),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(211),
		18: uint16(209),
		19: uint16(211),
		20: uint16(209),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		25: uint16(213),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
		53: uint16(2),
		54: uint16(2),
		55: uint16(2),
		56: uint16(2),
		57: uint16(2),
		58: uint16(2),
		59: uint16(2),
		60: uint16(114),
		61: uint16(2),
		62: uint16(2),
		63: uint16(2),
		64: uint16(2),
		65: uint16(2),
		66: uint16(2),
		67: uint16(2),
		68: uint16(2),
		69: uint16(2),
		70: uint16(2),
		71: uint16(2),
		72: uint16(2),
		73: uint16(2),
		74: uint16(2),
		75: uint16(2),
		80: uint16(2),
	},
	7: {
		1:  uint16(209),
		2:  uint16(211),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(211),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(211),
		18: uint16(209),
		19: uint16(211),
		20: uint16(209),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		27: uint16(213),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
		53: uint16(2),
		54: uint16(2),
		55: uint16(2),
		56: uint16(2),
		57: uint16(2),
		58: uint16(2),
		59: uint16(2),
		60: uint16(114),
		61: uint16(2),
		62: uint16(2),
		63: uint16(2),
		64: uint16(2),
		65: uint16(2),
		66: uint16(2),
		67: uint16(2),
		68: uint16(2),
		69: uint16(2),
		70: uint16(2),
		71: uint16(2),
		72: uint16(2),
		73: uint16(2),
		74: uint16(2),
		75: uint16(2),
		80: uint16(2),
	},
	8: {
		1:  uint16(215),
		2:  uint16(217),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(217),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(217),
		18: uint16(215),
		19: uint16(217),
		20: uint16(215),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		27: uint16(169),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(7),
		51: uint16(7),
		52: uint16(7),
		53: uint16(7),
		54: uint16(7),
		55: uint16(7),
		56: uint16(7),
		57: uint16(7),
		58: uint16(7),
		59: uint16(7),
		60: uint16(114),
		61: uint16(7),
		62: uint16(7),
		63: uint16(7),
		64: uint16(7),
		65: uint16(7),
		66: uint16(7),
		67: uint16(7),
		68: uint16(7),
		69: uint16(7),
		70: uint16(7),
		71: uint16(7),
		72: uint16(7),
		73: uint16(7),
		74: uint16(7),
		75: uint16(7),
		80: uint16(7),
	},
	9: {
		1:  uint16(219),
		2:  uint16(221),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(221),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(221),
		18: uint16(219),
		19: uint16(221),
		20: uint16(219),
		21: uint16(165),
		22: uint16(167),
		23: uint16(223),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(12),
		51: uint16(12),
		52: uint16(12),
		53: uint16(12),
		54: uint16(12),
		55: uint16(12),
		56: uint16(12),
		57: uint16(12),
		58: uint16(12),
		59: uint16(12),
		60: uint16(114),
		61: uint16(12),
		62: uint16(12),
		63: uint16(12),
		64: uint16(12),
		65: uint16(12),
		66: uint16(12),
		67: uint16(12),
		68: uint16(12),
		69: uint16(12),
		70: uint16(12),
		71: uint16(12),
		72: uint16(12),
		73: uint16(12),
		74: uint16(12),
		75: uint16(12),
		80: uint16(12),
	},
	10: {
		1:  uint16(225),
		2:  uint16(227),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(227),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(227),
		18: uint16(225),
		19: uint16(227),
		20: uint16(225),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		25: uint16(223),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(13),
		51: uint16(13),
		52: uint16(13),
		53: uint16(13),
		54: uint16(13),
		55: uint16(13),
		56: uint16(13),
		57: uint16(13),
		58: uint16(13),
		59: uint16(13),
		60: uint16(114),
		61: uint16(13),
		62: uint16(13),
		63: uint16(13),
		64: uint16(13),
		65: uint16(13),
		66: uint16(13),
		67: uint16(13),
		68: uint16(13),
		69: uint16(13),
		70: uint16(13),
		71: uint16(13),
		72: uint16(13),
		73: uint16(13),
		74: uint16(13),
		75: uint16(13),
		80: uint16(13),
	},
	11: {
		1:  uint16(229),
		2:  uint16(231),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(231),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(231),
		18: uint16(229),
		19: uint16(231),
		20: uint16(229),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		27: uint16(223),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(14),
		51: uint16(14),
		52: uint16(14),
		53: uint16(14),
		54: uint16(14),
		55: uint16(14),
		56: uint16(14),
		57: uint16(14),
		58: uint16(14),
		59: uint16(14),
		60: uint16(114),
		61: uint16(14),
		62: uint16(14),
		63: uint16(14),
		64: uint16(14),
		65: uint16(14),
		66: uint16(14),
		67: uint16(14),
		68: uint16(14),
		69: uint16(14),
		70: uint16(14),
		71: uint16(14),
		72: uint16(14),
		73: uint16(14),
		74: uint16(14),
		75: uint16(14),
		80: uint16(14),
	},
	12: {
		1:  uint16(209),
		2:  uint16(211),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(211),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(211),
		18: uint16(209),
		19: uint16(211),
		20: uint16(209),
		21: uint16(165),
		22: uint16(167),
		23: uint16(233),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
		53: uint16(2),
		54: uint16(2),
		55: uint16(2),
		56: uint16(2),
		57: uint16(2),
		58: uint16(2),
		59: uint16(2),
		60: uint16(114),
		61: uint16(2),
		62: uint16(2),
		63: uint16(2),
		64: uint16(2),
		65: uint16(2),
		66: uint16(2),
		67: uint16(2),
		68: uint16(2),
		69: uint16(2),
		70: uint16(2),
		71: uint16(2),
		72: uint16(2),
		73: uint16(2),
		74: uint16(2),
		75: uint16(2),
		80: uint16(2),
	},
	13: {
		1:  uint16(209),
		2:  uint16(211),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(211),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(211),
		18: uint16(209),
		19: uint16(211),
		20: uint16(209),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		25: uint16(233),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
		53: uint16(2),
		54: uint16(2),
		55: uint16(2),
		56: uint16(2),
		57: uint16(2),
		58: uint16(2),
		59: uint16(2),
		60: uint16(114),
		61: uint16(2),
		62: uint16(2),
		63: uint16(2),
		64: uint16(2),
		65: uint16(2),
		66: uint16(2),
		67: uint16(2),
		68: uint16(2),
		69: uint16(2),
		70: uint16(2),
		71: uint16(2),
		72: uint16(2),
		73: uint16(2),
		74: uint16(2),
		75: uint16(2),
		80: uint16(2),
	},
	14: {
		1:  uint16(209),
		2:  uint16(211),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(211),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(211),
		18: uint16(209),
		19: uint16(211),
		20: uint16(209),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		27: uint16(233),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
		53: uint16(2),
		54: uint16(2),
		55: uint16(2),
		56: uint16(2),
		57: uint16(2),
		58: uint16(2),
		59: uint16(2),
		60: uint16(114),
		61: uint16(2),
		62: uint16(2),
		63: uint16(2),
		64: uint16(2),
		65: uint16(2),
		66: uint16(2),
		67: uint16(2),
		68: uint16(2),
		69: uint16(2),
		70: uint16(2),
		71: uint16(2),
		72: uint16(2),
		73: uint16(2),
		74: uint16(2),
		75: uint16(2),
		80: uint16(2),
	},
	15: {
		1:  uint16(235),
		2:  uint16(237),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(237),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(237),
		18: uint16(235),
		19: uint16(237),
		20: uint16(235),
		21: uint16(165),
		22: uint16(167),
		23: uint16(239),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(18),
		51: uint16(18),
		52: uint16(18),
		53: uint16(18),
		54: uint16(18),
		55: uint16(18),
		56: uint16(18),
		57: uint16(18),
		58: uint16(18),
		59: uint16(18),
		60: uint16(114),
		61: uint16(18),
		62: uint16(18),
		63: uint16(18),
		64: uint16(18),
		65: uint16(18),
		66: uint16(18),
		67: uint16(18),
		68: uint16(18),
		69: uint16(18),
		70: uint16(18),
		71: uint16(18),
		72: uint16(18),
		73: uint16(18),
		74: uint16(18),
		75: uint16(18),
		80: uint16(18),
	},
	16: {
		1:  uint16(241),
		2:  uint16(243),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(243),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(243),
		18: uint16(241),
		19: uint16(243),
		20: uint16(241),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		25: uint16(239),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(19),
		51: uint16(19),
		52: uint16(19),
		53: uint16(19),
		54: uint16(19),
		55: uint16(19),
		56: uint16(19),
		57: uint16(19),
		58: uint16(19),
		59: uint16(19),
		60: uint16(114),
		61: uint16(19),
		62: uint16(19),
		63: uint16(19),
		64: uint16(19),
		65: uint16(19),
		66: uint16(19),
		67: uint16(19),
		68: uint16(19),
		69: uint16(19),
		70: uint16(19),
		71: uint16(19),
		72: uint16(19),
		73: uint16(19),
		74: uint16(19),
		75: uint16(19),
		80: uint16(19),
	},
	17: {
		1:  uint16(245),
		2:  uint16(247),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(247),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(247),
		18: uint16(245),
		19: uint16(247),
		20: uint16(245),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		27: uint16(239),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(20),
		51: uint16(20),
		52: uint16(20),
		53: uint16(20),
		54: uint16(20),
		55: uint16(20),
		56: uint16(20),
		57: uint16(20),
		58: uint16(20),
		59: uint16(20),
		60: uint16(114),
		61: uint16(20),
		62: uint16(20),
		63: uint16(20),
		64: uint16(20),
		65: uint16(20),
		66: uint16(20),
		67: uint16(20),
		68: uint16(20),
		69: uint16(20),
		70: uint16(20),
		71: uint16(20),
		72: uint16(20),
		73: uint16(20),
		74: uint16(20),
		75: uint16(20),
		80: uint16(20),
	},
	18: {
		1:  uint16(209),
		2:  uint16(211),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(211),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(211),
		18: uint16(209),
		19: uint16(211),
		20: uint16(209),
		21: uint16(165),
		22: uint16(167),
		23: uint16(249),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
		53: uint16(2),
		54: uint16(2),
		55: uint16(2),
		56: uint16(2),
		57: uint16(2),
		58: uint16(2),
		59: uint16(2),
		60: uint16(114),
		61: uint16(2),
		62: uint16(2),
		63: uint16(2),
		64: uint16(2),
		65: uint16(2),
		66: uint16(2),
		67: uint16(2),
		68: uint16(2),
		69: uint16(2),
		70: uint16(2),
		71: uint16(2),
		72: uint16(2),
		73: uint16(2),
		74: uint16(2),
		75: uint16(2),
		80: uint16(2),
	},
	19: {
		1:  uint16(209),
		2:  uint16(211),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(211),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(211),
		18: uint16(209),
		19: uint16(211),
		20: uint16(209),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		25: uint16(249),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
		53: uint16(2),
		54: uint16(2),
		55: uint16(2),
		56: uint16(2),
		57: uint16(2),
		58: uint16(2),
		59: uint16(2),
		60: uint16(114),
		61: uint16(2),
		62: uint16(2),
		63: uint16(2),
		64: uint16(2),
		65: uint16(2),
		66: uint16(2),
		67: uint16(2),
		68: uint16(2),
		69: uint16(2),
		70: uint16(2),
		71: uint16(2),
		72: uint16(2),
		73: uint16(2),
		74: uint16(2),
		75: uint16(2),
		80: uint16(2),
	},
	20: {
		1:  uint16(209),
		2:  uint16(211),
		3:  uint16(151),
		4:  uint16(153),
		7:  uint16(155),
		8:  uint16(151),
		9:  uint16(211),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(211),
		18: uint16(209),
		19: uint16(211),
		20: uint16(209),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		27: uint16(249),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		44: uint16(199),
		45: uint16(201),
		46: uint16(203),
		50: uint16(2),
		51: uint16(2),
		52: uint16(2),
		53: uint16(2),
		54: uint16(2),
		55: uint16(2),
		56: uint16(2),
		57: uint16(2),
		58: uint16(2),
		59: uint16(2),
		60: uint16(114),
		61: uint16(2),
		62: uint16(2),
		63: uint16(2),
		64: uint16(2),
		65: uint16(2),
		66: uint16(2),
		67: uint16(2),
		68: uint16(2),
		69: uint16(2),
		70: uint16(2),
		71: uint16(2),
		72: uint16(2),
		73: uint16(2),
		74: uint16(2),
		75: uint16(2),
		80: uint16(2),
	},
	21: {
		0:  uint16(251),
		1:  uint16(253),
		3:  uint16(7),
		4:  uint16(9),
		7:  uint16(11),
		8:  uint16(7),
		9:  uint16(255),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(255),
		18: uint16(253),
		19: uint16(255),
		20: uint16(253),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		44: uint16(55),
		45: uint16(57),
		46: uint16(59),
		50: uint16(22),
		51: uint16(22),
		52: uint16(22),
		53: uint16(22),
		54: uint16(22),
		55: uint16(22),
		56: uint16(22),
		57: uint16(22),
		58: uint16(22),
		59: uint16(22),
		60: uint16(143),
		61: uint16(22),
		62: uint16(22),
		63: uint16(22),
		64: uint16(22),
		65: uint16(22),
		66: uint16(22),
		67: uint16(22),
		68: uint16(22),
		69: uint16(22),
		70: uint16(22),
		71: uint16(22),
		72: uint16(22),
		73: uint16(22),
		74: uint16(22),
		75: uint16(22),
		76: uint16(22),
	},
	22: {
		0:  uint16(257),
		1:  uint16(259),
		3:  uint16(262),
		4:  uint16(265),
		7:  uint16(268),
		8:  uint16(262),
		9:  uint16(271),
		10: uint16(274),
		11: uint16(277),
		12: uint16(280),
		13: uint16(283),
		16: uint16(271),
		18: uint16(259),
		19: uint16(271),
		20: uint16(259),
		21: uint16(286),
		22: uint16(289),
		24: uint16(292),
		26: uint16(295),
		28: uint16(298),
		29: uint16(298),
		30: uint16(301),
		31: uint16(304),
		32: uint16(307),
		33: uint16(304),
		34: uint16(307),
		36: uint16(310),
		37: uint16(313),
		38: uint16(316),
		39: uint16(319),
		40: uint16(322),
		41: uint16(325),
		42: uint16(328),
		43: uint16(331),
		44: uint16(334),
		45: uint16(337),
		46: uint16(340),
		50: uint16(22),
		51: uint16(22),
		52: uint16(22),
		53: uint16(22),
		54: uint16(22),
		55: uint16(22),
		56: uint16(22),
		57: uint16(22),
		58: uint16(22),
		59: uint16(22),
		60: uint16(143),
		61: uint16(22),
		62: uint16(22),
		63: uint16(22),
		64: uint16(22),
		65: uint16(22),
		66: uint16(22),
		67: uint16(22),
		68: uint16(22),
		69: uint16(22),
		70: uint16(22),
		71: uint16(22),
		72: uint16(22),
		73: uint16(22),
		74: uint16(22),
		75: uint16(22),
		76: uint16(22),
	},
	23: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(351),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(351),
		18: uint16(353),
		19: uint16(351),
		20: uint16(353),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(169),
		56: uint16(169),
		57: uint16(169),
		58: uint16(169),
		59: uint16(169),
		60: uint16(143),
		61: uint16(169),
		62: uint16(169),
		63: uint16(169),
		64: uint16(169),
		65: uint16(169),
		66: uint16(169),
		67: uint16(169),
		68: uint16(169),
		69: uint16(169),
		70: uint16(169),
		71: uint16(169),
		72: uint16(169),
		73: uint16(169),
		74: uint16(169),
		78: uint16(133),
	},
	24: {
		1:  uint16(355),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(357),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(357),
		18: uint16(359),
		19: uint16(357),
		20: uint16(359),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(32),
		52: uint16(32),
		53: uint16(32),
		54: uint16(32),
		55: uint16(156),
		56: uint16(156),
		57: uint16(156),
		58: uint16(156),
		59: uint16(156),
		60: uint16(143),
		61: uint16(156),
		62: uint16(156),
		63: uint16(156),
		64: uint16(156),
		65: uint16(156),
		66: uint16(156),
		67: uint16(156),
		68: uint16(156),
		69: uint16(156),
		70: uint16(156),
		71: uint16(156),
		72: uint16(156),
		73: uint16(156),
		74: uint16(156),
		78: uint16(32),
	},
	25: {
		1:  uint16(361),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(363),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(363),
		18: uint16(365),
		19: uint16(363),
		20: uint16(365),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(33),
		52: uint16(33),
		53: uint16(33),
		54: uint16(33),
		55: uint16(157),
		56: uint16(157),
		57: uint16(157),
		58: uint16(157),
		59: uint16(157),
		60: uint16(143),
		61: uint16(157),
		62: uint16(157),
		63: uint16(157),
		64: uint16(157),
		65: uint16(157),
		66: uint16(157),
		67: uint16(157),
		68: uint16(157),
		69: uint16(157),
		70: uint16(157),
		71: uint16(157),
		72: uint16(157),
		73: uint16(157),
		74: uint16(157),
		78: uint16(33),
	},
	26: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(367),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(367),
		18: uint16(369),
		19: uint16(367),
		20: uint16(369),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(172),
		56: uint16(172),
		57: uint16(172),
		58: uint16(172),
		59: uint16(172),
		60: uint16(143),
		61: uint16(172),
		62: uint16(172),
		63: uint16(172),
		64: uint16(172),
		65: uint16(172),
		66: uint16(172),
		67: uint16(172),
		68: uint16(172),
		69: uint16(172),
		70: uint16(172),
		71: uint16(172),
		72: uint16(172),
		73: uint16(172),
		74: uint16(172),
		78: uint16(133),
	},
	27: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(371),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(371),
		18: uint16(373),
		19: uint16(371),
		20: uint16(373),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(150),
		56: uint16(150),
		57: uint16(150),
		58: uint16(150),
		59: uint16(150),
		60: uint16(143),
		61: uint16(150),
		62: uint16(150),
		63: uint16(150),
		64: uint16(150),
		65: uint16(150),
		66: uint16(150),
		67: uint16(150),
		68: uint16(150),
		69: uint16(150),
		70: uint16(150),
		71: uint16(150),
		72: uint16(150),
		73: uint16(150),
		74: uint16(150),
		78: uint16(133),
	},
	28: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(375),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(375),
		18: uint16(377),
		19: uint16(375),
		20: uint16(377),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(135),
		56: uint16(135),
		57: uint16(135),
		58: uint16(135),
		59: uint16(135),
		60: uint16(143),
		61: uint16(135),
		62: uint16(135),
		63: uint16(135),
		64: uint16(135),
		65: uint16(135),
		66: uint16(135),
		67: uint16(135),
		68: uint16(135),
		69: uint16(135),
		70: uint16(135),
		71: uint16(135),
		72: uint16(135),
		73: uint16(135),
		74: uint16(135),
		78: uint16(133),
	},
	29: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(379),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(379),
		18: uint16(381),
		19: uint16(379),
		20: uint16(381),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(136),
		56: uint16(136),
		57: uint16(136),
		58: uint16(136),
		59: uint16(136),
		60: uint16(143),
		61: uint16(136),
		62: uint16(136),
		63: uint16(136),
		64: uint16(136),
		65: uint16(136),
		66: uint16(136),
		67: uint16(136),
		68: uint16(136),
		69: uint16(136),
		70: uint16(136),
		71: uint16(136),
		72: uint16(136),
		73: uint16(136),
		74: uint16(136),
		78: uint16(133),
	},
	30: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(383),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(383),
		18: uint16(385),
		19: uint16(383),
		20: uint16(385),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(137),
		56: uint16(137),
		57: uint16(137),
		58: uint16(137),
		59: uint16(137),
		60: uint16(143),
		61: uint16(137),
		62: uint16(137),
		63: uint16(137),
		64: uint16(137),
		65: uint16(137),
		66: uint16(137),
		67: uint16(137),
		68: uint16(137),
		69: uint16(137),
		70: uint16(137),
		71: uint16(137),
		72: uint16(137),
		73: uint16(137),
		74: uint16(137),
		78: uint16(133),
	},
	31: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(387),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(387),
		18: uint16(389),
		19: uint16(387),
		20: uint16(389),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(138),
		56: uint16(138),
		57: uint16(138),
		58: uint16(138),
		59: uint16(138),
		60: uint16(143),
		61: uint16(138),
		62: uint16(138),
		63: uint16(138),
		64: uint16(138),
		65: uint16(138),
		66: uint16(138),
		67: uint16(138),
		68: uint16(138),
		69: uint16(138),
		70: uint16(138),
		71: uint16(138),
		72: uint16(138),
		73: uint16(138),
		74: uint16(138),
		78: uint16(133),
	},
	32: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(391),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(391),
		18: uint16(393),
		19: uint16(391),
		20: uint16(393),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(139),
		56: uint16(139),
		57: uint16(139),
		58: uint16(139),
		59: uint16(139),
		60: uint16(143),
		61: uint16(139),
		62: uint16(139),
		63: uint16(139),
		64: uint16(139),
		65: uint16(139),
		66: uint16(139),
		67: uint16(139),
		68: uint16(139),
		69: uint16(139),
		70: uint16(139),
		71: uint16(139),
		72: uint16(139),
		73: uint16(139),
		74: uint16(139),
		78: uint16(133),
	},
	33: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(395),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(395),
		18: uint16(397),
		19: uint16(395),
		20: uint16(397),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(140),
		56: uint16(140),
		57: uint16(140),
		58: uint16(140),
		59: uint16(140),
		60: uint16(143),
		61: uint16(140),
		62: uint16(140),
		63: uint16(140),
		64: uint16(140),
		65: uint16(140),
		66: uint16(140),
		67: uint16(140),
		68: uint16(140),
		69: uint16(140),
		70: uint16(140),
		71: uint16(140),
		72: uint16(140),
		73: uint16(140),
		74: uint16(140),
		78: uint16(133),
	},
	34: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(399),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(399),
		18: uint16(401),
		19: uint16(399),
		20: uint16(401),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(141),
		56: uint16(141),
		57: uint16(141),
		58: uint16(141),
		59: uint16(141),
		60: uint16(143),
		61: uint16(141),
		62: uint16(141),
		63: uint16(141),
		64: uint16(141),
		65: uint16(141),
		66: uint16(141),
		67: uint16(141),
		68: uint16(141),
		69: uint16(141),
		70: uint16(141),
		71: uint16(141),
		72: uint16(141),
		73: uint16(141),
		74: uint16(141),
		78: uint16(133),
	},
	35: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(403),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(403),
		18: uint16(405),
		19: uint16(403),
		20: uint16(405),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(163),
		56: uint16(163),
		57: uint16(163),
		58: uint16(163),
		59: uint16(163),
		60: uint16(143),
		61: uint16(163),
		62: uint16(163),
		63: uint16(163),
		64: uint16(163),
		65: uint16(163),
		66: uint16(163),
		67: uint16(163),
		68: uint16(163),
		69: uint16(163),
		70: uint16(163),
		71: uint16(163),
		72: uint16(163),
		73: uint16(163),
		74: uint16(163),
		78: uint16(133),
	},
	36: {
		1:  uint16(407),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(409),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(409),
		18: uint16(411),
		19: uint16(409),
		20: uint16(411),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(38),
		52: uint16(38),
		53: uint16(38),
		54: uint16(38),
		55: uint16(142),
		56: uint16(142),
		57: uint16(142),
		58: uint16(142),
		59: uint16(142),
		60: uint16(143),
		61: uint16(142),
		62: uint16(142),
		63: uint16(142),
		64: uint16(142),
		65: uint16(142),
		66: uint16(142),
		67: uint16(142),
		68: uint16(142),
		69: uint16(142),
		70: uint16(142),
		71: uint16(142),
		72: uint16(142),
		73: uint16(142),
		74: uint16(142),
		78: uint16(38),
	},
	37: {
		1:  uint16(413),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(415),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(415),
		18: uint16(417),
		19: uint16(415),
		20: uint16(417),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(23),
		52: uint16(23),
		53: uint16(23),
		54: uint16(23),
		55: uint16(171),
		56: uint16(171),
		57: uint16(171),
		58: uint16(171),
		59: uint16(171),
		60: uint16(143),
		61: uint16(171),
		62: uint16(171),
		63: uint16(171),
		64: uint16(171),
		65: uint16(171),
		66: uint16(171),
		67: uint16(171),
		68: uint16(171),
		69: uint16(171),
		70: uint16(171),
		71: uint16(171),
		72: uint16(171),
		73: uint16(171),
		74: uint16(171),
		78: uint16(23),
	},
	38: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(419),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(419),
		18: uint16(421),
		19: uint16(419),
		20: uint16(421),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(144),
		56: uint16(144),
		57: uint16(144),
		58: uint16(144),
		59: uint16(144),
		60: uint16(143),
		61: uint16(144),
		62: uint16(144),
		63: uint16(144),
		64: uint16(144),
		65: uint16(144),
		66: uint16(144),
		67: uint16(144),
		68: uint16(144),
		69: uint16(144),
		70: uint16(144),
		71: uint16(144),
		72: uint16(144),
		73: uint16(144),
		74: uint16(144),
		78: uint16(133),
	},
	39: {
		1:  uint16(423),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(425),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(425),
		18: uint16(427),
		19: uint16(425),
		20: uint16(427),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(50),
		52: uint16(50),
		53: uint16(50),
		54: uint16(50),
		55: uint16(132),
		56: uint16(132),
		57: uint16(132),
		58: uint16(132),
		59: uint16(132),
		60: uint16(114),
		61: uint16(132),
		62: uint16(132),
		63: uint16(132),
		64: uint16(132),
		65: uint16(132),
		66: uint16(132),
		67: uint16(132),
		68: uint16(132),
		69: uint16(132),
		70: uint16(132),
		71: uint16(132),
		72: uint16(132),
		73: uint16(132),
		74: uint16(132),
		78: uint16(50),
	},
	40: {
		1:  uint16(429),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(431),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(431),
		18: uint16(433),
		19: uint16(431),
		20: uint16(433),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(51),
		52: uint16(51),
		53: uint16(51),
		54: uint16(51),
		55: uint16(99),
		56: uint16(99),
		57: uint16(99),
		58: uint16(99),
		59: uint16(99),
		60: uint16(114),
		61: uint16(99),
		62: uint16(99),
		63: uint16(99),
		64: uint16(99),
		65: uint16(99),
		66: uint16(99),
		67: uint16(99),
		68: uint16(99),
		69: uint16(99),
		70: uint16(99),
		71: uint16(99),
		72: uint16(99),
		73: uint16(99),
		74: uint16(99),
		78: uint16(51),
	},
	41: {
		1:  uint16(435),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(437),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(437),
		18: uint16(439),
		19: uint16(437),
		20: uint16(439),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(52),
		52: uint16(52),
		53: uint16(52),
		54: uint16(52),
		55: uint16(100),
		56: uint16(100),
		57: uint16(100),
		58: uint16(100),
		59: uint16(100),
		60: uint16(114),
		61: uint16(100),
		62: uint16(100),
		63: uint16(100),
		64: uint16(100),
		65: uint16(100),
		66: uint16(100),
		67: uint16(100),
		68: uint16(100),
		69: uint16(100),
		70: uint16(100),
		71: uint16(100),
		72: uint16(100),
		73: uint16(100),
		74: uint16(100),
		78: uint16(52),
	},
	42: {
		1:  uint16(441),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(443),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(443),
		18: uint16(445),
		19: uint16(443),
		20: uint16(445),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(53),
		52: uint16(53),
		53: uint16(53),
		54: uint16(53),
		55: uint16(102),
		56: uint16(102),
		57: uint16(102),
		58: uint16(102),
		59: uint16(102),
		60: uint16(114),
		61: uint16(102),
		62: uint16(102),
		63: uint16(102),
		64: uint16(102),
		65: uint16(102),
		66: uint16(102),
		67: uint16(102),
		68: uint16(102),
		69: uint16(102),
		70: uint16(102),
		71: uint16(102),
		72: uint16(102),
		73: uint16(102),
		74: uint16(102),
		78: uint16(53),
	},
	43: {
		1:  uint16(447),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(449),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(449),
		18: uint16(451),
		19: uint16(449),
		20: uint16(451),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(54),
		52: uint16(54),
		53: uint16(54),
		54: uint16(54),
		55: uint16(104),
		56: uint16(104),
		57: uint16(104),
		58: uint16(104),
		59: uint16(104),
		60: uint16(114),
		61: uint16(104),
		62: uint16(104),
		63: uint16(104),
		64: uint16(104),
		65: uint16(104),
		66: uint16(104),
		67: uint16(104),
		68: uint16(104),
		69: uint16(104),
		70: uint16(104),
		71: uint16(104),
		72: uint16(104),
		73: uint16(104),
		74: uint16(104),
		78: uint16(54),
	},
	44: {
		1:  uint16(453),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(455),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(455),
		18: uint16(457),
		19: uint16(455),
		20: uint16(457),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(55),
		52: uint16(55),
		53: uint16(55),
		54: uint16(55),
		55: uint16(106),
		56: uint16(106),
		57: uint16(106),
		58: uint16(106),
		59: uint16(106),
		60: uint16(114),
		61: uint16(106),
		62: uint16(106),
		63: uint16(106),
		64: uint16(106),
		65: uint16(106),
		66: uint16(106),
		67: uint16(106),
		68: uint16(106),
		69: uint16(106),
		70: uint16(106),
		71: uint16(106),
		72: uint16(106),
		73: uint16(106),
		74: uint16(106),
		78: uint16(55),
	},
	45: {
		1:  uint16(459),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(461),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(461),
		18: uint16(463),
		19: uint16(461),
		20: uint16(463),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(56),
		52: uint16(56),
		53: uint16(56),
		54: uint16(56),
		55: uint16(109),
		56: uint16(109),
		57: uint16(109),
		58: uint16(109),
		59: uint16(109),
		60: uint16(114),
		61: uint16(109),
		62: uint16(109),
		63: uint16(109),
		64: uint16(109),
		65: uint16(109),
		66: uint16(109),
		67: uint16(109),
		68: uint16(109),
		69: uint16(109),
		70: uint16(109),
		71: uint16(109),
		72: uint16(109),
		73: uint16(109),
		74: uint16(109),
		78: uint16(56),
	},
	46: {
		1:  uint16(465),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(467),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(467),
		18: uint16(469),
		19: uint16(467),
		20: uint16(469),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(92),
		52: uint16(92),
		53: uint16(92),
		54: uint16(92),
		55: uint16(93),
		56: uint16(93),
		57: uint16(93),
		58: uint16(93),
		59: uint16(93),
		60: uint16(114),
		61: uint16(93),
		62: uint16(93),
		63: uint16(93),
		64: uint16(93),
		65: uint16(93),
		66: uint16(93),
		67: uint16(93),
		68: uint16(93),
		69: uint16(93),
		70: uint16(93),
		71: uint16(93),
		72: uint16(93),
		73: uint16(93),
		74: uint16(93),
		78: uint16(92),
	},
	47: {
		1:  uint16(471),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(473),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(473),
		18: uint16(475),
		19: uint16(473),
		20: uint16(475),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(58),
		52: uint16(58),
		53: uint16(58),
		54: uint16(58),
		55: uint16(113),
		56: uint16(113),
		57: uint16(113),
		58: uint16(113),
		59: uint16(113),
		60: uint16(114),
		61: uint16(113),
		62: uint16(113),
		63: uint16(113),
		64: uint16(113),
		65: uint16(113),
		66: uint16(113),
		67: uint16(113),
		68: uint16(113),
		69: uint16(113),
		70: uint16(113),
		71: uint16(113),
		72: uint16(113),
		73: uint16(113),
		74: uint16(113),
		78: uint16(58),
	},
	48: {
		1:  uint16(477),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(479),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(479),
		18: uint16(481),
		19: uint16(479),
		20: uint16(481),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(59),
		52: uint16(59),
		53: uint16(59),
		54: uint16(59),
		55: uint16(115),
		56: uint16(115),
		57: uint16(115),
		58: uint16(115),
		59: uint16(115),
		60: uint16(114),
		61: uint16(115),
		62: uint16(115),
		63: uint16(115),
		64: uint16(115),
		65: uint16(115),
		66: uint16(115),
		67: uint16(115),
		68: uint16(115),
		69: uint16(115),
		70: uint16(115),
		71: uint16(115),
		72: uint16(115),
		73: uint16(115),
		74: uint16(115),
		78: uint16(59),
	},
	49: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(483),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(483),
		18: uint16(493),
		19: uint16(483),
		20: uint16(493),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(208),
		56: uint16(208),
		57: uint16(208),
		58: uint16(208),
		59: uint16(208),
		60: uint16(201),
		61: uint16(208),
		62: uint16(208),
		63: uint16(208),
		64: uint16(208),
		65: uint16(208),
		66: uint16(208),
		67: uint16(208),
		68: uint16(208),
		69: uint16(208),
		70: uint16(208),
		71: uint16(208),
		72: uint16(208),
		73: uint16(208),
		74: uint16(208),
		78: uint16(133),
	},
	50: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(527),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(527),
		18: uint16(529),
		19: uint16(527),
		20: uint16(529),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(124),
		56: uint16(124),
		57: uint16(124),
		58: uint16(124),
		59: uint16(124),
		60: uint16(114),
		61: uint16(124),
		62: uint16(124),
		63: uint16(124),
		64: uint16(124),
		65: uint16(124),
		66: uint16(124),
		67: uint16(124),
		68: uint16(124),
		69: uint16(124),
		70: uint16(124),
		71: uint16(124),
		72: uint16(124),
		73: uint16(124),
		74: uint16(124),
		78: uint16(133),
	},
	51: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(531),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(531),
		18: uint16(533),
		19: uint16(531),
		20: uint16(533),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(126),
		56: uint16(126),
		57: uint16(126),
		58: uint16(126),
		59: uint16(126),
		60: uint16(114),
		61: uint16(126),
		62: uint16(126),
		63: uint16(126),
		64: uint16(126),
		65: uint16(126),
		66: uint16(126),
		67: uint16(126),
		68: uint16(126),
		69: uint16(126),
		70: uint16(126),
		71: uint16(126),
		72: uint16(126),
		73: uint16(126),
		74: uint16(126),
		78: uint16(133),
	},
	52: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(535),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(535),
		18: uint16(537),
		19: uint16(535),
		20: uint16(537),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(127),
		56: uint16(127),
		57: uint16(127),
		58: uint16(127),
		59: uint16(127),
		60: uint16(114),
		61: uint16(127),
		62: uint16(127),
		63: uint16(127),
		64: uint16(127),
		65: uint16(127),
		66: uint16(127),
		67: uint16(127),
		68: uint16(127),
		69: uint16(127),
		70: uint16(127),
		71: uint16(127),
		72: uint16(127),
		73: uint16(127),
		74: uint16(127),
		78: uint16(133),
	},
	53: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(539),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(539),
		18: uint16(541),
		19: uint16(539),
		20: uint16(541),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(128),
		56: uint16(128),
		57: uint16(128),
		58: uint16(128),
		59: uint16(128),
		60: uint16(114),
		61: uint16(128),
		62: uint16(128),
		63: uint16(128),
		64: uint16(128),
		65: uint16(128),
		66: uint16(128),
		67: uint16(128),
		68: uint16(128),
		69: uint16(128),
		70: uint16(128),
		71: uint16(128),
		72: uint16(128),
		73: uint16(128),
		74: uint16(128),
		78: uint16(133),
	},
	54: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(543),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(543),
		18: uint16(545),
		19: uint16(543),
		20: uint16(545),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(129),
		56: uint16(129),
		57: uint16(129),
		58: uint16(129),
		59: uint16(129),
		60: uint16(114),
		61: uint16(129),
		62: uint16(129),
		63: uint16(129),
		64: uint16(129),
		65: uint16(129),
		66: uint16(129),
		67: uint16(129),
		68: uint16(129),
		69: uint16(129),
		70: uint16(129),
		71: uint16(129),
		72: uint16(129),
		73: uint16(129),
		74: uint16(129),
		78: uint16(133),
	},
	55: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(547),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(547),
		18: uint16(549),
		19: uint16(547),
		20: uint16(549),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(130),
		56: uint16(130),
		57: uint16(130),
		58: uint16(130),
		59: uint16(130),
		60: uint16(114),
		61: uint16(130),
		62: uint16(130),
		63: uint16(130),
		64: uint16(130),
		65: uint16(130),
		66: uint16(130),
		67: uint16(130),
		68: uint16(130),
		69: uint16(130),
		70: uint16(130),
		71: uint16(130),
		72: uint16(130),
		73: uint16(130),
		74: uint16(130),
		78: uint16(133),
	},
	56: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(551),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(551),
		18: uint16(553),
		19: uint16(551),
		20: uint16(553),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(131),
		56: uint16(131),
		57: uint16(131),
		58: uint16(131),
		59: uint16(131),
		60: uint16(114),
		61: uint16(131),
		62: uint16(131),
		63: uint16(131),
		64: uint16(131),
		65: uint16(131),
		66: uint16(131),
		67: uint16(131),
		68: uint16(131),
		69: uint16(131),
		70: uint16(131),
		71: uint16(131),
		72: uint16(131),
		73: uint16(131),
		74: uint16(131),
		78: uint16(133),
	},
	57: {
		1:  uint16(555),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(557),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(557),
		18: uint16(559),
		19: uint16(557),
		20: uint16(559),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(34),
		52: uint16(34),
		53: uint16(34),
		54: uint16(34),
		55: uint16(159),
		56: uint16(159),
		57: uint16(159),
		58: uint16(159),
		59: uint16(159),
		60: uint16(143),
		61: uint16(159),
		62: uint16(159),
		63: uint16(159),
		64: uint16(159),
		65: uint16(159),
		66: uint16(159),
		67: uint16(159),
		68: uint16(159),
		69: uint16(159),
		70: uint16(159),
		71: uint16(159),
		72: uint16(159),
		73: uint16(159),
		74: uint16(159),
		78: uint16(34),
	},
	58: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(561),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(561),
		18: uint16(563),
		19: uint16(561),
		20: uint16(563),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(112),
		56: uint16(112),
		57: uint16(112),
		58: uint16(112),
		59: uint16(112),
		60: uint16(114),
		61: uint16(112),
		62: uint16(112),
		63: uint16(112),
		64: uint16(112),
		65: uint16(112),
		66: uint16(112),
		67: uint16(112),
		68: uint16(112),
		69: uint16(112),
		70: uint16(112),
		71: uint16(112),
		72: uint16(112),
		73: uint16(112),
		74: uint16(112),
		78: uint16(133),
	},
	59: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(565),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(565),
		18: uint16(567),
		19: uint16(565),
		20: uint16(567),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(94),
		56: uint16(94),
		57: uint16(94),
		58: uint16(94),
		59: uint16(94),
		60: uint16(114),
		61: uint16(94),
		62: uint16(94),
		63: uint16(94),
		64: uint16(94),
		65: uint16(94),
		66: uint16(94),
		67: uint16(94),
		68: uint16(94),
		69: uint16(94),
		70: uint16(94),
		71: uint16(94),
		72: uint16(94),
		73: uint16(94),
		74: uint16(94),
		78: uint16(133),
	},
	60: {
		1:  uint16(569),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(571),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(571),
		18: uint16(573),
		19: uint16(571),
		20: uint16(573),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(61),
		52: uint16(61),
		53: uint16(61),
		54: uint16(61),
		55: uint16(95),
		56: uint16(95),
		57: uint16(95),
		58: uint16(95),
		59: uint16(95),
		60: uint16(114),
		61: uint16(95),
		62: uint16(95),
		63: uint16(95),
		64: uint16(95),
		65: uint16(95),
		66: uint16(95),
		67: uint16(95),
		68: uint16(95),
		69: uint16(95),
		70: uint16(95),
		71: uint16(95),
		72: uint16(95),
		73: uint16(95),
		74: uint16(95),
		78: uint16(61),
	},
	61: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(575),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(575),
		18: uint16(577),
		19: uint16(575),
		20: uint16(577),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(97),
		56: uint16(97),
		57: uint16(97),
		58: uint16(97),
		59: uint16(97),
		60: uint16(114),
		61: uint16(97),
		62: uint16(97),
		63: uint16(97),
		64: uint16(97),
		65: uint16(97),
		66: uint16(97),
		67: uint16(97),
		68: uint16(97),
		69: uint16(97),
		70: uint16(97),
		71: uint16(97),
		72: uint16(97),
		73: uint16(97),
		74: uint16(97),
		78: uint16(133),
	},
	62: {
		1:  uint16(579),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(581),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(581),
		18: uint16(583),
		19: uint16(581),
		20: uint16(583),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(72),
		52: uint16(72),
		53: uint16(72),
		54: uint16(72),
		55: uint16(210),
		56: uint16(210),
		57: uint16(210),
		58: uint16(210),
		59: uint16(210),
		60: uint16(201),
		61: uint16(210),
		62: uint16(210),
		63: uint16(210),
		64: uint16(210),
		65: uint16(210),
		66: uint16(210),
		67: uint16(210),
		68: uint16(210),
		69: uint16(210),
		70: uint16(210),
		71: uint16(210),
		72: uint16(210),
		73: uint16(210),
		74: uint16(210),
		78: uint16(72),
	},
	63: {
		1:  uint16(585),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(587),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(587),
		18: uint16(589),
		19: uint16(587),
		20: uint16(589),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(73),
		52: uint16(73),
		53: uint16(73),
		54: uint16(73),
		55: uint16(177),
		56: uint16(177),
		57: uint16(177),
		58: uint16(177),
		59: uint16(177),
		60: uint16(201),
		61: uint16(177),
		62: uint16(177),
		63: uint16(177),
		64: uint16(177),
		65: uint16(177),
		66: uint16(177),
		67: uint16(177),
		68: uint16(177),
		69: uint16(177),
		70: uint16(177),
		71: uint16(177),
		72: uint16(177),
		73: uint16(177),
		74: uint16(177),
		78: uint16(73),
	},
	64: {
		1:  uint16(591),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(593),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(593),
		18: uint16(595),
		19: uint16(593),
		20: uint16(595),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(74),
		52: uint16(74),
		53: uint16(74),
		54: uint16(74),
		55: uint16(178),
		56: uint16(178),
		57: uint16(178),
		58: uint16(178),
		59: uint16(178),
		60: uint16(201),
		61: uint16(178),
		62: uint16(178),
		63: uint16(178),
		64: uint16(178),
		65: uint16(178),
		66: uint16(178),
		67: uint16(178),
		68: uint16(178),
		69: uint16(178),
		70: uint16(178),
		71: uint16(178),
		72: uint16(178),
		73: uint16(178),
		74: uint16(178),
		78: uint16(74),
	},
	65: {
		1:  uint16(597),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(599),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(599),
		18: uint16(601),
		19: uint16(599),
		20: uint16(601),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(75),
		52: uint16(75),
		53: uint16(75),
		54: uint16(75),
		55: uint16(179),
		56: uint16(179),
		57: uint16(179),
		58: uint16(179),
		59: uint16(179),
		60: uint16(201),
		61: uint16(179),
		62: uint16(179),
		63: uint16(179),
		64: uint16(179),
		65: uint16(179),
		66: uint16(179),
		67: uint16(179),
		68: uint16(179),
		69: uint16(179),
		70: uint16(179),
		71: uint16(179),
		72: uint16(179),
		73: uint16(179),
		74: uint16(179),
		78: uint16(75),
	},
	66: {
		1:  uint16(603),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(605),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(605),
		18: uint16(607),
		19: uint16(605),
		20: uint16(607),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(76),
		52: uint16(76),
		53: uint16(76),
		54: uint16(76),
		55: uint16(180),
		56: uint16(180),
		57: uint16(180),
		58: uint16(180),
		59: uint16(180),
		60: uint16(201),
		61: uint16(180),
		62: uint16(180),
		63: uint16(180),
		64: uint16(180),
		65: uint16(180),
		66: uint16(180),
		67: uint16(180),
		68: uint16(180),
		69: uint16(180),
		70: uint16(180),
		71: uint16(180),
		72: uint16(180),
		73: uint16(180),
		74: uint16(180),
		78: uint16(76),
	},
	67: {
		1:  uint16(609),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(611),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(611),
		18: uint16(613),
		19: uint16(611),
		20: uint16(613),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(77),
		52: uint16(77),
		53: uint16(77),
		54: uint16(77),
		55: uint16(181),
		56: uint16(181),
		57: uint16(181),
		58: uint16(181),
		59: uint16(181),
		60: uint16(201),
		61: uint16(181),
		62: uint16(181),
		63: uint16(181),
		64: uint16(181),
		65: uint16(181),
		66: uint16(181),
		67: uint16(181),
		68: uint16(181),
		69: uint16(181),
		70: uint16(181),
		71: uint16(181),
		72: uint16(181),
		73: uint16(181),
		74: uint16(181),
		78: uint16(77),
	},
	68: {
		1:  uint16(615),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(617),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(617),
		18: uint16(619),
		19: uint16(617),
		20: uint16(619),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(78),
		52: uint16(78),
		53: uint16(78),
		54: uint16(78),
		55: uint16(183),
		56: uint16(183),
		57: uint16(183),
		58: uint16(183),
		59: uint16(183),
		60: uint16(201),
		61: uint16(183),
		62: uint16(183),
		63: uint16(183),
		64: uint16(183),
		65: uint16(183),
		66: uint16(183),
		67: uint16(183),
		68: uint16(183),
		69: uint16(183),
		70: uint16(183),
		71: uint16(183),
		72: uint16(183),
		73: uint16(183),
		74: uint16(183),
		78: uint16(78),
	},
	69: {
		1:  uint16(621),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(623),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(623),
		18: uint16(625),
		19: uint16(623),
		20: uint16(625),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(79),
		52: uint16(79),
		53: uint16(79),
		54: uint16(79),
		55: uint16(185),
		56: uint16(185),
		57: uint16(185),
		58: uint16(185),
		59: uint16(185),
		60: uint16(201),
		61: uint16(185),
		62: uint16(185),
		63: uint16(185),
		64: uint16(185),
		65: uint16(185),
		66: uint16(185),
		67: uint16(185),
		68: uint16(185),
		69: uint16(185),
		70: uint16(185),
		71: uint16(185),
		72: uint16(185),
		73: uint16(185),
		74: uint16(185),
		78: uint16(79),
	},
	70: {
		1:  uint16(627),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(629),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(629),
		18: uint16(631),
		19: uint16(629),
		20: uint16(631),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(80),
		52: uint16(80),
		53: uint16(80),
		54: uint16(80),
		55: uint16(186),
		56: uint16(186),
		57: uint16(186),
		58: uint16(186),
		59: uint16(186),
		60: uint16(201),
		61: uint16(186),
		62: uint16(186),
		63: uint16(186),
		64: uint16(186),
		65: uint16(186),
		66: uint16(186),
		67: uint16(186),
		68: uint16(186),
		69: uint16(186),
		70: uint16(186),
		71: uint16(186),
		72: uint16(186),
		73: uint16(186),
		74: uint16(186),
		78: uint16(80),
	},
	71: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(633),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(633),
		18: uint16(635),
		19: uint16(633),
		20: uint16(635),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(118),
		56: uint16(118),
		57: uint16(118),
		58: uint16(118),
		59: uint16(118),
		60: uint16(114),
		61: uint16(118),
		62: uint16(118),
		63: uint16(118),
		64: uint16(118),
		65: uint16(118),
		66: uint16(118),
		67: uint16(118),
		68: uint16(118),
		69: uint16(118),
		70: uint16(118),
		71: uint16(118),
		72: uint16(118),
		73: uint16(118),
		74: uint16(118),
		78: uint16(133),
	},
	72: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(637),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(637),
		18: uint16(639),
		19: uint16(637),
		20: uint16(639),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(190),
		56: uint16(190),
		57: uint16(190),
		58: uint16(190),
		59: uint16(190),
		60: uint16(201),
		61: uint16(190),
		62: uint16(190),
		63: uint16(190),
		64: uint16(190),
		65: uint16(190),
		66: uint16(190),
		67: uint16(190),
		68: uint16(190),
		69: uint16(190),
		70: uint16(190),
		71: uint16(190),
		72: uint16(190),
		73: uint16(190),
		74: uint16(190),
		78: uint16(133),
	},
	73: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(641),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(641),
		18: uint16(643),
		19: uint16(641),
		20: uint16(643),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(192),
		56: uint16(192),
		57: uint16(192),
		58: uint16(192),
		59: uint16(192),
		60: uint16(201),
		61: uint16(192),
		62: uint16(192),
		63: uint16(192),
		64: uint16(192),
		65: uint16(192),
		66: uint16(192),
		67: uint16(192),
		68: uint16(192),
		69: uint16(192),
		70: uint16(192),
		71: uint16(192),
		72: uint16(192),
		73: uint16(192),
		74: uint16(192),
		78: uint16(133),
	},
	74: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(645),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(645),
		18: uint16(647),
		19: uint16(645),
		20: uint16(647),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(193),
		56: uint16(193),
		57: uint16(193),
		58: uint16(193),
		59: uint16(193),
		60: uint16(201),
		61: uint16(193),
		62: uint16(193),
		63: uint16(193),
		64: uint16(193),
		65: uint16(193),
		66: uint16(193),
		67: uint16(193),
		68: uint16(193),
		69: uint16(193),
		70: uint16(193),
		71: uint16(193),
		72: uint16(193),
		73: uint16(193),
		74: uint16(193),
		78: uint16(133),
	},
	75: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(649),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(649),
		18: uint16(651),
		19: uint16(649),
		20: uint16(651),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(211),
		56: uint16(211),
		57: uint16(211),
		58: uint16(211),
		59: uint16(211),
		60: uint16(201),
		61: uint16(211),
		62: uint16(211),
		63: uint16(211),
		64: uint16(211),
		65: uint16(211),
		66: uint16(211),
		67: uint16(211),
		68: uint16(211),
		69: uint16(211),
		70: uint16(211),
		71: uint16(211),
		72: uint16(211),
		73: uint16(211),
		74: uint16(211),
		78: uint16(133),
	},
	76: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(653),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(653),
		18: uint16(655),
		19: uint16(653),
		20: uint16(655),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(195),
		56: uint16(195),
		57: uint16(195),
		58: uint16(195),
		59: uint16(195),
		60: uint16(201),
		61: uint16(195),
		62: uint16(195),
		63: uint16(195),
		64: uint16(195),
		65: uint16(195),
		66: uint16(195),
		67: uint16(195),
		68: uint16(195),
		69: uint16(195),
		70: uint16(195),
		71: uint16(195),
		72: uint16(195),
		73: uint16(195),
		74: uint16(195),
		78: uint16(133),
	},
	77: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(657),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(657),
		18: uint16(659),
		19: uint16(657),
		20: uint16(659),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(196),
		56: uint16(196),
		57: uint16(196),
		58: uint16(196),
		59: uint16(196),
		60: uint16(201),
		61: uint16(196),
		62: uint16(196),
		63: uint16(196),
		64: uint16(196),
		65: uint16(196),
		66: uint16(196),
		67: uint16(196),
		68: uint16(196),
		69: uint16(196),
		70: uint16(196),
		71: uint16(196),
		72: uint16(196),
		73: uint16(196),
		74: uint16(196),
		78: uint16(133),
	},
	78: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(661),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(661),
		18: uint16(663),
		19: uint16(661),
		20: uint16(663),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(197),
		56: uint16(197),
		57: uint16(197),
		58: uint16(197),
		59: uint16(197),
		60: uint16(201),
		61: uint16(197),
		62: uint16(197),
		63: uint16(197),
		64: uint16(197),
		65: uint16(197),
		66: uint16(197),
		67: uint16(197),
		68: uint16(197),
		69: uint16(197),
		70: uint16(197),
		71: uint16(197),
		72: uint16(197),
		73: uint16(197),
		74: uint16(197),
		78: uint16(133),
	},
	79: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(665),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(665),
		18: uint16(667),
		19: uint16(665),
		20: uint16(667),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(198),
		56: uint16(198),
		57: uint16(198),
		58: uint16(198),
		59: uint16(198),
		60: uint16(201),
		61: uint16(198),
		62: uint16(198),
		63: uint16(198),
		64: uint16(198),
		65: uint16(198),
		66: uint16(198),
		67: uint16(198),
		68: uint16(198),
		69: uint16(198),
		70: uint16(198),
		71: uint16(198),
		72: uint16(198),
		73: uint16(198),
		74: uint16(198),
		78: uint16(133),
	},
	80: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(669),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(669),
		18: uint16(671),
		19: uint16(669),
		20: uint16(671),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(199),
		56: uint16(199),
		57: uint16(199),
		58: uint16(199),
		59: uint16(199),
		60: uint16(201),
		61: uint16(199),
		62: uint16(199),
		63: uint16(199),
		64: uint16(199),
		65: uint16(199),
		66: uint16(199),
		67: uint16(199),
		68: uint16(199),
		69: uint16(199),
		70: uint16(199),
		71: uint16(199),
		72: uint16(199),
		73: uint16(199),
		74: uint16(199),
		78: uint16(133),
	},
	81: {
		1:  uint16(673),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(675),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(675),
		18: uint16(677),
		19: uint16(675),
		20: uint16(677),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(82),
		52: uint16(82),
		53: uint16(82),
		54: uint16(82),
		55: uint16(200),
		56: uint16(200),
		57: uint16(200),
		58: uint16(200),
		59: uint16(200),
		60: uint16(201),
		61: uint16(200),
		62: uint16(200),
		63: uint16(200),
		64: uint16(200),
		65: uint16(200),
		66: uint16(200),
		67: uint16(200),
		68: uint16(200),
		69: uint16(200),
		70: uint16(200),
		71: uint16(200),
		72: uint16(200),
		73: uint16(200),
		74: uint16(200),
		78: uint16(82),
	},
	82: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(679),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(679),
		18: uint16(681),
		19: uint16(679),
		20: uint16(681),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(202),
		56: uint16(202),
		57: uint16(202),
		58: uint16(202),
		59: uint16(202),
		60: uint16(201),
		61: uint16(202),
		62: uint16(202),
		63: uint16(202),
		64: uint16(202),
		65: uint16(202),
		66: uint16(202),
		67: uint16(202),
		68: uint16(202),
		69: uint16(202),
		70: uint16(202),
		71: uint16(202),
		72: uint16(202),
		73: uint16(202),
		74: uint16(202),
		78: uint16(133),
	},
	83: {
		1:  uint16(683),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(685),
		10: uint16(485),
		11: uint16(487),
		12: uint16(489),
		13: uint16(491),
		16: uint16(685),
		18: uint16(687),
		19: uint16(685),
		20: uint16(687),
		21: uint16(495),
		22: uint16(497),
		24: uint16(499),
		26: uint16(501),
		28: uint16(503),
		29: uint16(503),
		30: uint16(505),
		31: uint16(507),
		32: uint16(509),
		33: uint16(507),
		34: uint16(509),
		36: uint16(511),
		37: uint16(513),
		38: uint16(515),
		39: uint16(517),
		40: uint16(519),
		41: uint16(521),
		42: uint16(523),
		43: uint16(525),
		51: uint16(49),
		52: uint16(49),
		53: uint16(49),
		54: uint16(49),
		55: uint16(182),
		56: uint16(182),
		57: uint16(182),
		58: uint16(182),
		59: uint16(182),
		60: uint16(201),
		61: uint16(182),
		62: uint16(182),
		63: uint16(182),
		64: uint16(182),
		65: uint16(182),
		66: uint16(182),
		67: uint16(182),
		68: uint16(182),
		69: uint16(182),
		70: uint16(182),
		71: uint16(182),
		72: uint16(182),
		73: uint16(182),
		74: uint16(182),
		78: uint16(49),
	},
	84: {
		1:  uint16(689),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(691),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(691),
		18: uint16(693),
		19: uint16(691),
		20: uint16(693),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(35),
		52: uint16(35),
		53: uint16(35),
		54: uint16(35),
		55: uint16(153),
		56: uint16(153),
		57: uint16(153),
		58: uint16(153),
		59: uint16(153),
		60: uint16(143),
		61: uint16(153),
		62: uint16(153),
		63: uint16(153),
		64: uint16(153),
		65: uint16(153),
		66: uint16(153),
		67: uint16(153),
		68: uint16(153),
		69: uint16(153),
		70: uint16(153),
		71: uint16(153),
		72: uint16(153),
		73: uint16(153),
		74: uint16(153),
		78: uint16(35),
	},
	85: {
		1:  uint16(695),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(697),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(697),
		18: uint16(699),
		19: uint16(697),
		20: uint16(699),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(71),
		52: uint16(71),
		53: uint16(71),
		54: uint16(71),
		55: uint16(107),
		56: uint16(107),
		57: uint16(107),
		58: uint16(107),
		59: uint16(107),
		60: uint16(114),
		61: uint16(107),
		62: uint16(107),
		63: uint16(107),
		64: uint16(107),
		65: uint16(107),
		66: uint16(107),
		67: uint16(107),
		68: uint16(107),
		69: uint16(107),
		70: uint16(107),
		71: uint16(107),
		72: uint16(107),
		73: uint16(107),
		74: uint16(107),
		78: uint16(71),
	},
	86: {
		1:  uint16(701),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(703),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(703),
		18: uint16(705),
		19: uint16(703),
		20: uint16(705),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(26),
		52: uint16(26),
		53: uint16(26),
		54: uint16(26),
		55: uint16(148),
		56: uint16(148),
		57: uint16(148),
		58: uint16(148),
		59: uint16(148),
		60: uint16(143),
		61: uint16(148),
		62: uint16(148),
		63: uint16(148),
		64: uint16(148),
		65: uint16(148),
		66: uint16(148),
		67: uint16(148),
		68: uint16(148),
		69: uint16(148),
		70: uint16(148),
		71: uint16(148),
		72: uint16(148),
		73: uint16(148),
		74: uint16(148),
		78: uint16(26),
	},
	87: {
		1:  uint16(707),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(709),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(709),
		18: uint16(711),
		19: uint16(709),
		20: uint16(711),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(27),
		52: uint16(27),
		53: uint16(27),
		54: uint16(27),
		55: uint16(149),
		56: uint16(149),
		57: uint16(149),
		58: uint16(149),
		59: uint16(149),
		60: uint16(143),
		61: uint16(149),
		62: uint16(149),
		63: uint16(149),
		64: uint16(149),
		65: uint16(149),
		66: uint16(149),
		67: uint16(149),
		68: uint16(149),
		69: uint16(149),
		70: uint16(149),
		71: uint16(149),
		72: uint16(149),
		73: uint16(149),
		74: uint16(149),
		78: uint16(27),
	},
	88: {
		1:  uint16(713),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(715),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(715),
		18: uint16(717),
		19: uint16(715),
		20: uint16(717),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(28),
		52: uint16(28),
		53: uint16(28),
		54: uint16(28),
		55: uint16(151),
		56: uint16(151),
		57: uint16(151),
		58: uint16(151),
		59: uint16(151),
		60: uint16(143),
		61: uint16(151),
		62: uint16(151),
		63: uint16(151),
		64: uint16(151),
		65: uint16(151),
		66: uint16(151),
		67: uint16(151),
		68: uint16(151),
		69: uint16(151),
		70: uint16(151),
		71: uint16(151),
		72: uint16(151),
		73: uint16(151),
		74: uint16(151),
		78: uint16(28),
	},
	89: {
		1:  uint16(719),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(721),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(721),
		18: uint16(723),
		19: uint16(721),
		20: uint16(723),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(29),
		52: uint16(29),
		53: uint16(29),
		54: uint16(29),
		55: uint16(152),
		56: uint16(152),
		57: uint16(152),
		58: uint16(152),
		59: uint16(152),
		60: uint16(143),
		61: uint16(152),
		62: uint16(152),
		63: uint16(152),
		64: uint16(152),
		65: uint16(152),
		66: uint16(152),
		67: uint16(152),
		68: uint16(152),
		69: uint16(152),
		70: uint16(152),
		71: uint16(152),
		72: uint16(152),
		73: uint16(152),
		74: uint16(152),
		78: uint16(29),
	},
	90: {
		1:  uint16(725),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(727),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(727),
		18: uint16(729),
		19: uint16(727),
		20: uint16(729),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(30),
		52: uint16(30),
		53: uint16(30),
		54: uint16(30),
		55: uint16(154),
		56: uint16(154),
		57: uint16(154),
		58: uint16(154),
		59: uint16(154),
		60: uint16(143),
		61: uint16(154),
		62: uint16(154),
		63: uint16(154),
		64: uint16(154),
		65: uint16(154),
		66: uint16(154),
		67: uint16(154),
		68: uint16(154),
		69: uint16(154),
		70: uint16(154),
		71: uint16(154),
		72: uint16(154),
		73: uint16(154),
		74: uint16(154),
		78: uint16(30),
	},
	91: {
		1:  uint16(731),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(733),
		10: uint16(15),
		11: uint16(17),
		12: uint16(19),
		13: uint16(21),
		16: uint16(733),
		18: uint16(735),
		19: uint16(733),
		20: uint16(735),
		21: uint16(23),
		22: uint16(25),
		24: uint16(27),
		26: uint16(29),
		28: uint16(31),
		29: uint16(31),
		30: uint16(33),
		31: uint16(35),
		32: uint16(37),
		33: uint16(35),
		34: uint16(37),
		36: uint16(39),
		37: uint16(41),
		38: uint16(43),
		39: uint16(45),
		40: uint16(47),
		41: uint16(49),
		42: uint16(51),
		43: uint16(53),
		51: uint16(31),
		52: uint16(31),
		53: uint16(31),
		54: uint16(31),
		55: uint16(155),
		56: uint16(155),
		57: uint16(155),
		58: uint16(155),
		59: uint16(155),
		60: uint16(143),
		61: uint16(155),
		62: uint16(155),
		63: uint16(155),
		64: uint16(155),
		65: uint16(155),
		66: uint16(155),
		67: uint16(155),
		68: uint16(155),
		69: uint16(155),
		70: uint16(155),
		71: uint16(155),
		72: uint16(155),
		73: uint16(155),
		74: uint16(155),
		78: uint16(31),
	},
	92: {
		1:  uint16(343),
		3:  uint16(345),
		4:  uint16(347),
		7:  uint16(349),
		8:  uint16(345),
		9:  uint16(737),
		10: uint16(157),
		11: uint16(159),
		12: uint16(161),
		13: uint16(163),
		16: uint16(737),
		18: uint16(739),
		19: uint16(737),
		20: uint16(739),
		21: uint16(165),
		22: uint16(167),
		24: uint16(171),
		26: uint16(173),
		28: uint16(175),
		29: uint16(175),
		30: uint16(177),
		31: uint16(179),
		32: uint16(181),
		33: uint16(179),
		34: uint16(181),
		36: uint16(183),
		37: uint16(185),
		38: uint16(187),
		39: uint16(189),
		40: uint16(191),
		41: uint16(193),
		42: uint16(195),
		43: uint16(197),
		51: uint16(133),
		52: uint16(133),
		53: uint16(133),
		54: uint16(133),
		55: uint16(111),
		56: uint16(111),
		57: uint16(111),
		58: uint16(111),
		59: uint16(111),
		60: uint16(114),
		61: uint16(111),
		62: uint16(111),
		63: uint16(111),
		64: uint16(111),
		65: uint16(111),
		66: uint16(111),
		67: uint16(111),
		68: uint16(111),
		69: uint16(111),
		70: uint16(111),
		71: uint16(111),
		72: uint16(111),
		73: uint16(111),
		74: uint16(111),
		78: uint16(133),
	},
}

var ts_small_parse_table = [5529]uint16_t{
	0:    uint16(2),
	1:    uint16(743),
	2:    uint16(10),
	3:    uint16(sym_dot),
	4:    uint16(sym_boolean),
	5:    uint16(anon_sym_POUND),
	6:    uint16(sym_number),
	7:    uint16(sym_symbol),
	8:    uint16(anon_sym_POUNDhash),
	9:    uint16(anon_sym_POUNDhasheq),
	10:   uint16(anon_sym_COMMA),
	11:   uint16(anon_sym_POUND_COMMA),
	12:   uint16(anon_sym_POUND_BANG),
	13:   uint16(741),
	14:   uint16(30),
	15:   uint16(aux_sym__skip_token1),
	16:   uint16(aux_sym_comment_token1),
	17:   uint16(anon_sym_POUND_PIPE),
	18:   uint16(anon_sym_POUND_SEMI),
	19:   uint16(sym__line_comment),
	20:   uint16(anon_sym_POUND_LT_LT),
	21:   uint16(aux_sym_regex_token1),
	22:   uint16(anon_sym_DQUOTE),
	23:   uint16(sym_character),
	24:   uint16(sym_keyword),
	25:   uint16(anon_sym_POUND_AMP),
	26:   uint16(anon_sym_LPAREN),
	27:   uint16(anon_sym_RPAREN),
	28:   uint16(anon_sym_LBRACK),
	29:   uint16(anon_sym_RBRACK),
	30:   uint16(anon_sym_LBRACE),
	31:   uint16(anon_sym_RBRACE),
	32:   uint16(anon_sym_POUNDfl),
	33:   uint16(anon_sym_POUNDfx),
	34:   uint16(anon_sym_POUNDs),
	35:   uint16(anon_sym_POUNDhashalw),
	36:   uint16(anon_sym_POUNDhasheqv),
	37:   uint16(anon_sym_SQUOTE),
	38:   uint16(anon_sym_BQUOTE),
	39:   uint16(anon_sym_POUND_SQUOTE),
	40:   uint16(anon_sym_POUND_BQUOTE),
	41:   uint16(anon_sym_COMMA_AT),
	42:   uint16(anon_sym_POUND_COMMA_AT),
	43:   uint16(anon_sym_POUNDreader),
	44:   uint16(anon_sym_POUNDlang),
	45:   uint16(2),
	46:   uint16(747),
	47:   uint16(10),
	48:   uint16(sym_dot),
	49:   uint16(sym_boolean),
	50:   uint16(anon_sym_POUND),
	51:   uint16(sym_number),
	52:   uint16(sym_symbol),
	53:   uint16(anon_sym_POUNDhash),
	54:   uint16(anon_sym_POUNDhasheq),
	55:   uint16(anon_sym_COMMA),
	56:   uint16(anon_sym_POUND_COMMA),
	57:   uint16(anon_sym_POUND_BANG),
	58:   uint16(745),
	59:   uint16(30),
	60:   uint16(aux_sym__skip_token1),
	61:   uint16(aux_sym_comment_token1),
	62:   uint16(anon_sym_POUND_PIPE),
	63:   uint16(anon_sym_POUND_SEMI),
	64:   uint16(sym__line_comment),
	65:   uint16(anon_sym_POUND_LT_LT),
	66:   uint16(aux_sym_regex_token1),
	67:   uint16(anon_sym_DQUOTE),
	68:   uint16(sym_character),
	69:   uint16(sym_keyword),
	70:   uint16(anon_sym_POUND_AMP),
	71:   uint16(anon_sym_LPAREN),
	72:   uint16(anon_sym_RPAREN),
	73:   uint16(anon_sym_LBRACK),
	74:   uint16(anon_sym_RBRACK),
	75:   uint16(anon_sym_LBRACE),
	76:   uint16(anon_sym_RBRACE),
	77:   uint16(anon_sym_POUNDfl),
	78:   uint16(anon_sym_POUNDfx),
	79:   uint16(anon_sym_POUNDs),
	80:   uint16(anon_sym_POUNDhashalw),
	81:   uint16(anon_sym_POUNDhasheqv),
	82:   uint16(anon_sym_SQUOTE),
	83:   uint16(anon_sym_BQUOTE),
	84:   uint16(anon_sym_POUND_SQUOTE),
	85:   uint16(anon_sym_POUND_BQUOTE),
	86:   uint16(anon_sym_COMMA_AT),
	87:   uint16(anon_sym_POUND_COMMA_AT),
	88:   uint16(anon_sym_POUNDreader),
	89:   uint16(anon_sym_POUNDlang),
	90:   uint16(2),
	91:   uint16(751),
	92:   uint16(10),
	93:   uint16(sym_dot),
	94:   uint16(sym_boolean),
	95:   uint16(anon_sym_POUND),
	96:   uint16(sym_number),
	97:   uint16(sym_symbol),
	98:   uint16(anon_sym_POUNDhash),
	99:   uint16(anon_sym_POUNDhasheq),
	100:  uint16(anon_sym_COMMA),
	101:  uint16(anon_sym_POUND_COMMA),
	102:  uint16(anon_sym_POUND_BANG),
	103:  uint16(749),
	104:  uint16(30),
	105:  uint16(aux_sym__skip_token1),
	106:  uint16(aux_sym_comment_token1),
	107:  uint16(anon_sym_POUND_PIPE),
	108:  uint16(anon_sym_POUND_SEMI),
	109:  uint16(sym__line_comment),
	110:  uint16(anon_sym_POUND_LT_LT),
	111:  uint16(aux_sym_regex_token1),
	112:  uint16(anon_sym_DQUOTE),
	113:  uint16(sym_character),
	114:  uint16(sym_keyword),
	115:  uint16(anon_sym_POUND_AMP),
	116:  uint16(anon_sym_LPAREN),
	117:  uint16(anon_sym_RPAREN),
	118:  uint16(anon_sym_LBRACK),
	119:  uint16(anon_sym_RBRACK),
	120:  uint16(anon_sym_LBRACE),
	121:  uint16(anon_sym_RBRACE),
	122:  uint16(anon_sym_POUNDfl),
	123:  uint16(anon_sym_POUNDfx),
	124:  uint16(anon_sym_POUNDs),
	125:  uint16(anon_sym_POUNDhashalw),
	126:  uint16(anon_sym_POUNDhasheqv),
	127:  uint16(anon_sym_SQUOTE),
	128:  uint16(anon_sym_BQUOTE),
	129:  uint16(anon_sym_POUND_SQUOTE),
	130:  uint16(anon_sym_POUND_BQUOTE),
	131:  uint16(anon_sym_COMMA_AT),
	132:  uint16(anon_sym_POUND_COMMA_AT),
	133:  uint16(anon_sym_POUNDreader),
	134:  uint16(anon_sym_POUNDlang),
	135:  uint16(2),
	136:  uint16(755),
	137:  uint16(10),
	138:  uint16(sym_dot),
	139:  uint16(sym_boolean),
	140:  uint16(anon_sym_POUND),
	141:  uint16(sym_number),
	142:  uint16(sym_symbol),
	143:  uint16(anon_sym_POUNDhash),
	144:  uint16(anon_sym_POUNDhasheq),
	145:  uint16(anon_sym_COMMA),
	146:  uint16(anon_sym_POUND_COMMA),
	147:  uint16(anon_sym_POUND_BANG),
	148:  uint16(753),
	149:  uint16(30),
	150:  uint16(aux_sym__skip_token1),
	151:  uint16(aux_sym_comment_token1),
	152:  uint16(anon_sym_POUND_PIPE),
	153:  uint16(anon_sym_POUND_SEMI),
	154:  uint16(sym__line_comment),
	155:  uint16(anon_sym_POUND_LT_LT),
	156:  uint16(aux_sym_regex_token1),
	157:  uint16(anon_sym_DQUOTE),
	158:  uint16(sym_character),
	159:  uint16(sym_keyword),
	160:  uint16(anon_sym_POUND_AMP),
	161:  uint16(anon_sym_LPAREN),
	162:  uint16(anon_sym_RPAREN),
	163:  uint16(anon_sym_LBRACK),
	164:  uint16(anon_sym_RBRACK),
	165:  uint16(anon_sym_LBRACE),
	166:  uint16(anon_sym_RBRACE),
	167:  uint16(anon_sym_POUNDfl),
	168:  uint16(anon_sym_POUNDfx),
	169:  uint16(anon_sym_POUNDs),
	170:  uint16(anon_sym_POUNDhashalw),
	171:  uint16(anon_sym_POUNDhasheqv),
	172:  uint16(anon_sym_SQUOTE),
	173:  uint16(anon_sym_BQUOTE),
	174:  uint16(anon_sym_POUND_SQUOTE),
	175:  uint16(anon_sym_POUND_BQUOTE),
	176:  uint16(anon_sym_COMMA_AT),
	177:  uint16(anon_sym_POUND_COMMA_AT),
	178:  uint16(anon_sym_POUNDreader),
	179:  uint16(anon_sym_POUNDlang),
	180:  uint16(2),
	181:  uint16(759),
	182:  uint16(10),
	183:  uint16(sym_dot),
	184:  uint16(sym_boolean),
	185:  uint16(anon_sym_POUND),
	186:  uint16(sym_number),
	187:  uint16(sym_symbol),
	188:  uint16(anon_sym_POUNDhash),
	189:  uint16(anon_sym_POUNDhasheq),
	190:  uint16(anon_sym_COMMA),
	191:  uint16(anon_sym_POUND_COMMA),
	192:  uint16(anon_sym_POUND_BANG),
	193:  uint16(757),
	194:  uint16(30),
	195:  uint16(aux_sym__skip_token1),
	196:  uint16(aux_sym_comment_token1),
	197:  uint16(anon_sym_POUND_PIPE),
	198:  uint16(anon_sym_POUND_SEMI),
	199:  uint16(sym__line_comment),
	200:  uint16(anon_sym_POUND_LT_LT),
	201:  uint16(aux_sym_regex_token1),
	202:  uint16(anon_sym_DQUOTE),
	203:  uint16(sym_character),
	204:  uint16(sym_keyword),
	205:  uint16(anon_sym_POUND_AMP),
	206:  uint16(anon_sym_LPAREN),
	207:  uint16(anon_sym_RPAREN),
	208:  uint16(anon_sym_LBRACK),
	209:  uint16(anon_sym_RBRACK),
	210:  uint16(anon_sym_LBRACE),
	211:  uint16(anon_sym_RBRACE),
	212:  uint16(anon_sym_POUNDfl),
	213:  uint16(anon_sym_POUNDfx),
	214:  uint16(anon_sym_POUNDs),
	215:  uint16(anon_sym_POUNDhashalw),
	216:  uint16(anon_sym_POUNDhasheqv),
	217:  uint16(anon_sym_SQUOTE),
	218:  uint16(anon_sym_BQUOTE),
	219:  uint16(anon_sym_POUND_SQUOTE),
	220:  uint16(anon_sym_POUND_BQUOTE),
	221:  uint16(anon_sym_COMMA_AT),
	222:  uint16(anon_sym_POUND_COMMA_AT),
	223:  uint16(anon_sym_POUNDreader),
	224:  uint16(anon_sym_POUNDlang),
	225:  uint16(2),
	226:  uint16(763),
	227:  uint16(10),
	228:  uint16(sym_dot),
	229:  uint16(sym_boolean),
	230:  uint16(anon_sym_POUND),
	231:  uint16(sym_number),
	232:  uint16(sym_symbol),
	233:  uint16(anon_sym_POUNDhash),
	234:  uint16(anon_sym_POUNDhasheq),
	235:  uint16(anon_sym_COMMA),
	236:  uint16(anon_sym_POUND_COMMA),
	237:  uint16(anon_sym_POUND_BANG),
	238:  uint16(761),
	239:  uint16(30),
	240:  uint16(aux_sym__skip_token1),
	241:  uint16(aux_sym_comment_token1),
	242:  uint16(anon_sym_POUND_PIPE),
	243:  uint16(anon_sym_POUND_SEMI),
	244:  uint16(sym__line_comment),
	245:  uint16(anon_sym_POUND_LT_LT),
	246:  uint16(aux_sym_regex_token1),
	247:  uint16(anon_sym_DQUOTE),
	248:  uint16(sym_character),
	249:  uint16(sym_keyword),
	250:  uint16(anon_sym_POUND_AMP),
	251:  uint16(anon_sym_LPAREN),
	252:  uint16(anon_sym_RPAREN),
	253:  uint16(anon_sym_LBRACK),
	254:  uint16(anon_sym_RBRACK),
	255:  uint16(anon_sym_LBRACE),
	256:  uint16(anon_sym_RBRACE),
	257:  uint16(anon_sym_POUNDfl),
	258:  uint16(anon_sym_POUNDfx),
	259:  uint16(anon_sym_POUNDs),
	260:  uint16(anon_sym_POUNDhashalw),
	261:  uint16(anon_sym_POUNDhasheqv),
	262:  uint16(anon_sym_SQUOTE),
	263:  uint16(anon_sym_BQUOTE),
	264:  uint16(anon_sym_POUND_SQUOTE),
	265:  uint16(anon_sym_POUND_BQUOTE),
	266:  uint16(anon_sym_COMMA_AT),
	267:  uint16(anon_sym_POUND_COMMA_AT),
	268:  uint16(anon_sym_POUNDreader),
	269:  uint16(anon_sym_POUNDlang),
	270:  uint16(2),
	271:  uint16(767),
	272:  uint16(10),
	273:  uint16(sym_dot),
	274:  uint16(sym_boolean),
	275:  uint16(anon_sym_POUND),
	276:  uint16(sym_number),
	277:  uint16(sym_symbol),
	278:  uint16(anon_sym_POUNDhash),
	279:  uint16(anon_sym_POUNDhasheq),
	280:  uint16(anon_sym_COMMA),
	281:  uint16(anon_sym_POUND_COMMA),
	282:  uint16(anon_sym_POUND_BANG),
	283:  uint16(765),
	284:  uint16(30),
	285:  uint16(aux_sym__skip_token1),
	286:  uint16(aux_sym_comment_token1),
	287:  uint16(anon_sym_POUND_PIPE),
	288:  uint16(anon_sym_POUND_SEMI),
	289:  uint16(sym__line_comment),
	290:  uint16(anon_sym_POUND_LT_LT),
	291:  uint16(aux_sym_regex_token1),
	292:  uint16(anon_sym_DQUOTE),
	293:  uint16(sym_character),
	294:  uint16(sym_keyword),
	295:  uint16(anon_sym_POUND_AMP),
	296:  uint16(anon_sym_LPAREN),
	297:  uint16(anon_sym_RPAREN),
	298:  uint16(anon_sym_LBRACK),
	299:  uint16(anon_sym_RBRACK),
	300:  uint16(anon_sym_LBRACE),
	301:  uint16(anon_sym_RBRACE),
	302:  uint16(anon_sym_POUNDfl),
	303:  uint16(anon_sym_POUNDfx),
	304:  uint16(anon_sym_POUNDs),
	305:  uint16(anon_sym_POUNDhashalw),
	306:  uint16(anon_sym_POUNDhasheqv),
	307:  uint16(anon_sym_SQUOTE),
	308:  uint16(anon_sym_BQUOTE),
	309:  uint16(anon_sym_POUND_SQUOTE),
	310:  uint16(anon_sym_POUND_BQUOTE),
	311:  uint16(anon_sym_COMMA_AT),
	312:  uint16(anon_sym_POUND_COMMA_AT),
	313:  uint16(anon_sym_POUNDreader),
	314:  uint16(anon_sym_POUNDlang),
	315:  uint16(2),
	316:  uint16(771),
	317:  uint16(10),
	318:  uint16(sym_dot),
	319:  uint16(sym_boolean),
	320:  uint16(anon_sym_POUND),
	321:  uint16(sym_number),
	322:  uint16(sym_symbol),
	323:  uint16(anon_sym_POUNDhash),
	324:  uint16(anon_sym_POUNDhasheq),
	325:  uint16(anon_sym_COMMA),
	326:  uint16(anon_sym_POUND_COMMA),
	327:  uint16(anon_sym_POUND_BANG),
	328:  uint16(769),
	329:  uint16(30),
	330:  uint16(aux_sym__skip_token1),
	331:  uint16(aux_sym_comment_token1),
	332:  uint16(anon_sym_POUND_PIPE),
	333:  uint16(anon_sym_POUND_SEMI),
	334:  uint16(sym__line_comment),
	335:  uint16(anon_sym_POUND_LT_LT),
	336:  uint16(aux_sym_regex_token1),
	337:  uint16(anon_sym_DQUOTE),
	338:  uint16(sym_character),
	339:  uint16(sym_keyword),
	340:  uint16(anon_sym_POUND_AMP),
	341:  uint16(anon_sym_LPAREN),
	342:  uint16(anon_sym_RPAREN),
	343:  uint16(anon_sym_LBRACK),
	344:  uint16(anon_sym_RBRACK),
	345:  uint16(anon_sym_LBRACE),
	346:  uint16(anon_sym_RBRACE),
	347:  uint16(anon_sym_POUNDfl),
	348:  uint16(anon_sym_POUNDfx),
	349:  uint16(anon_sym_POUNDs),
	350:  uint16(anon_sym_POUNDhashalw),
	351:  uint16(anon_sym_POUNDhasheqv),
	352:  uint16(anon_sym_SQUOTE),
	353:  uint16(anon_sym_BQUOTE),
	354:  uint16(anon_sym_POUND_SQUOTE),
	355:  uint16(anon_sym_POUND_BQUOTE),
	356:  uint16(anon_sym_COMMA_AT),
	357:  uint16(anon_sym_POUND_COMMA_AT),
	358:  uint16(anon_sym_POUNDreader),
	359:  uint16(anon_sym_POUNDlang),
	360:  uint16(2),
	361:  uint16(775),
	362:  uint16(10),
	363:  uint16(sym_dot),
	364:  uint16(sym_boolean),
	365:  uint16(anon_sym_POUND),
	366:  uint16(sym_number),
	367:  uint16(sym_symbol),
	368:  uint16(anon_sym_POUNDhash),
	369:  uint16(anon_sym_POUNDhasheq),
	370:  uint16(anon_sym_COMMA),
	371:  uint16(anon_sym_POUND_COMMA),
	372:  uint16(anon_sym_POUND_BANG),
	373:  uint16(773),
	374:  uint16(30),
	375:  uint16(aux_sym__skip_token1),
	376:  uint16(aux_sym_comment_token1),
	377:  uint16(anon_sym_POUND_PIPE),
	378:  uint16(anon_sym_POUND_SEMI),
	379:  uint16(sym__line_comment),
	380:  uint16(anon_sym_POUND_LT_LT),
	381:  uint16(aux_sym_regex_token1),
	382:  uint16(anon_sym_DQUOTE),
	383:  uint16(sym_character),
	384:  uint16(sym_keyword),
	385:  uint16(anon_sym_POUND_AMP),
	386:  uint16(anon_sym_LPAREN),
	387:  uint16(anon_sym_RPAREN),
	388:  uint16(anon_sym_LBRACK),
	389:  uint16(anon_sym_RBRACK),
	390:  uint16(anon_sym_LBRACE),
	391:  uint16(anon_sym_RBRACE),
	392:  uint16(anon_sym_POUNDfl),
	393:  uint16(anon_sym_POUNDfx),
	394:  uint16(anon_sym_POUNDs),
	395:  uint16(anon_sym_POUNDhashalw),
	396:  uint16(anon_sym_POUNDhasheqv),
	397:  uint16(anon_sym_SQUOTE),
	398:  uint16(anon_sym_BQUOTE),
	399:  uint16(anon_sym_POUND_SQUOTE),
	400:  uint16(anon_sym_POUND_BQUOTE),
	401:  uint16(anon_sym_COMMA_AT),
	402:  uint16(anon_sym_POUND_COMMA_AT),
	403:  uint16(anon_sym_POUNDreader),
	404:  uint16(anon_sym_POUNDlang),
	405:  uint16(2),
	406:  uint16(779),
	407:  uint16(10),
	408:  uint16(sym_dot),
	409:  uint16(sym_boolean),
	410:  uint16(anon_sym_POUND),
	411:  uint16(sym_number),
	412:  uint16(sym_symbol),
	413:  uint16(anon_sym_POUNDhash),
	414:  uint16(anon_sym_POUNDhasheq),
	415:  uint16(anon_sym_COMMA),
	416:  uint16(anon_sym_POUND_COMMA),
	417:  uint16(anon_sym_POUND_BANG),
	418:  uint16(777),
	419:  uint16(30),
	420:  uint16(aux_sym__skip_token1),
	421:  uint16(aux_sym_comment_token1),
	422:  uint16(anon_sym_POUND_PIPE),
	423:  uint16(anon_sym_POUND_SEMI),
	424:  uint16(sym__line_comment),
	425:  uint16(anon_sym_POUND_LT_LT),
	426:  uint16(aux_sym_regex_token1),
	427:  uint16(anon_sym_DQUOTE),
	428:  uint16(sym_character),
	429:  uint16(sym_keyword),
	430:  uint16(anon_sym_POUND_AMP),
	431:  uint16(anon_sym_LPAREN),
	432:  uint16(anon_sym_RPAREN),
	433:  uint16(anon_sym_LBRACK),
	434:  uint16(anon_sym_RBRACK),
	435:  uint16(anon_sym_LBRACE),
	436:  uint16(anon_sym_RBRACE),
	437:  uint16(anon_sym_POUNDfl),
	438:  uint16(anon_sym_POUNDfx),
	439:  uint16(anon_sym_POUNDs),
	440:  uint16(anon_sym_POUNDhashalw),
	441:  uint16(anon_sym_POUNDhasheqv),
	442:  uint16(anon_sym_SQUOTE),
	443:  uint16(anon_sym_BQUOTE),
	444:  uint16(anon_sym_POUND_SQUOTE),
	445:  uint16(anon_sym_POUND_BQUOTE),
	446:  uint16(anon_sym_COMMA_AT),
	447:  uint16(anon_sym_POUND_COMMA_AT),
	448:  uint16(anon_sym_POUNDreader),
	449:  uint16(anon_sym_POUNDlang),
	450:  uint16(2),
	451:  uint16(783),
	452:  uint16(10),
	453:  uint16(sym_dot),
	454:  uint16(sym_boolean),
	455:  uint16(anon_sym_POUND),
	456:  uint16(sym_number),
	457:  uint16(sym_symbol),
	458:  uint16(anon_sym_POUNDhash),
	459:  uint16(anon_sym_POUNDhasheq),
	460:  uint16(anon_sym_COMMA),
	461:  uint16(anon_sym_POUND_COMMA),
	462:  uint16(anon_sym_POUND_BANG),
	463:  uint16(781),
	464:  uint16(30),
	465:  uint16(aux_sym__skip_token1),
	466:  uint16(aux_sym_comment_token1),
	467:  uint16(anon_sym_POUND_PIPE),
	468:  uint16(anon_sym_POUND_SEMI),
	469:  uint16(sym__line_comment),
	470:  uint16(anon_sym_POUND_LT_LT),
	471:  uint16(aux_sym_regex_token1),
	472:  uint16(anon_sym_DQUOTE),
	473:  uint16(sym_character),
	474:  uint16(sym_keyword),
	475:  uint16(anon_sym_POUND_AMP),
	476:  uint16(anon_sym_LPAREN),
	477:  uint16(anon_sym_RPAREN),
	478:  uint16(anon_sym_LBRACK),
	479:  uint16(anon_sym_RBRACK),
	480:  uint16(anon_sym_LBRACE),
	481:  uint16(anon_sym_RBRACE),
	482:  uint16(anon_sym_POUNDfl),
	483:  uint16(anon_sym_POUNDfx),
	484:  uint16(anon_sym_POUNDs),
	485:  uint16(anon_sym_POUNDhashalw),
	486:  uint16(anon_sym_POUNDhasheqv),
	487:  uint16(anon_sym_SQUOTE),
	488:  uint16(anon_sym_BQUOTE),
	489:  uint16(anon_sym_POUND_SQUOTE),
	490:  uint16(anon_sym_POUND_BQUOTE),
	491:  uint16(anon_sym_COMMA_AT),
	492:  uint16(anon_sym_POUND_COMMA_AT),
	493:  uint16(anon_sym_POUNDreader),
	494:  uint16(anon_sym_POUNDlang),
	495:  uint16(2),
	496:  uint16(787),
	497:  uint16(10),
	498:  uint16(sym_dot),
	499:  uint16(sym_boolean),
	500:  uint16(anon_sym_POUND),
	501:  uint16(sym_number),
	502:  uint16(sym_symbol),
	503:  uint16(anon_sym_POUNDhash),
	504:  uint16(anon_sym_POUNDhasheq),
	505:  uint16(anon_sym_COMMA),
	506:  uint16(anon_sym_POUND_COMMA),
	507:  uint16(anon_sym_POUND_BANG),
	508:  uint16(785),
	509:  uint16(30),
	510:  uint16(aux_sym__skip_token1),
	511:  uint16(aux_sym_comment_token1),
	512:  uint16(anon_sym_POUND_PIPE),
	513:  uint16(anon_sym_POUND_SEMI),
	514:  uint16(sym__line_comment),
	515:  uint16(anon_sym_POUND_LT_LT),
	516:  uint16(aux_sym_regex_token1),
	517:  uint16(anon_sym_DQUOTE),
	518:  uint16(sym_character),
	519:  uint16(sym_keyword),
	520:  uint16(anon_sym_POUND_AMP),
	521:  uint16(anon_sym_LPAREN),
	522:  uint16(anon_sym_RPAREN),
	523:  uint16(anon_sym_LBRACK),
	524:  uint16(anon_sym_RBRACK),
	525:  uint16(anon_sym_LBRACE),
	526:  uint16(anon_sym_RBRACE),
	527:  uint16(anon_sym_POUNDfl),
	528:  uint16(anon_sym_POUNDfx),
	529:  uint16(anon_sym_POUNDs),
	530:  uint16(anon_sym_POUNDhashalw),
	531:  uint16(anon_sym_POUNDhasheqv),
	532:  uint16(anon_sym_SQUOTE),
	533:  uint16(anon_sym_BQUOTE),
	534:  uint16(anon_sym_POUND_SQUOTE),
	535:  uint16(anon_sym_POUND_BQUOTE),
	536:  uint16(anon_sym_COMMA_AT),
	537:  uint16(anon_sym_POUND_COMMA_AT),
	538:  uint16(anon_sym_POUNDreader),
	539:  uint16(anon_sym_POUNDlang),
	540:  uint16(2),
	541:  uint16(791),
	542:  uint16(10),
	543:  uint16(sym_dot),
	544:  uint16(sym_boolean),
	545:  uint16(anon_sym_POUND),
	546:  uint16(sym_number),
	547:  uint16(sym_symbol),
	548:  uint16(anon_sym_POUNDhash),
	549:  uint16(anon_sym_POUNDhasheq),
	550:  uint16(anon_sym_COMMA),
	551:  uint16(anon_sym_POUND_COMMA),
	552:  uint16(anon_sym_POUND_BANG),
	553:  uint16(789),
	554:  uint16(30),
	555:  uint16(aux_sym__skip_token1),
	556:  uint16(aux_sym_comment_token1),
	557:  uint16(anon_sym_POUND_PIPE),
	558:  uint16(anon_sym_POUND_SEMI),
	559:  uint16(sym__line_comment),
	560:  uint16(anon_sym_POUND_LT_LT),
	561:  uint16(aux_sym_regex_token1),
	562:  uint16(anon_sym_DQUOTE),
	563:  uint16(sym_character),
	564:  uint16(sym_keyword),
	565:  uint16(anon_sym_POUND_AMP),
	566:  uint16(anon_sym_LPAREN),
	567:  uint16(anon_sym_RPAREN),
	568:  uint16(anon_sym_LBRACK),
	569:  uint16(anon_sym_RBRACK),
	570:  uint16(anon_sym_LBRACE),
	571:  uint16(anon_sym_RBRACE),
	572:  uint16(anon_sym_POUNDfl),
	573:  uint16(anon_sym_POUNDfx),
	574:  uint16(anon_sym_POUNDs),
	575:  uint16(anon_sym_POUNDhashalw),
	576:  uint16(anon_sym_POUNDhasheqv),
	577:  uint16(anon_sym_SQUOTE),
	578:  uint16(anon_sym_BQUOTE),
	579:  uint16(anon_sym_POUND_SQUOTE),
	580:  uint16(anon_sym_POUND_BQUOTE),
	581:  uint16(anon_sym_COMMA_AT),
	582:  uint16(anon_sym_POUND_COMMA_AT),
	583:  uint16(anon_sym_POUNDreader),
	584:  uint16(anon_sym_POUNDlang),
	585:  uint16(2),
	586:  uint16(795),
	587:  uint16(10),
	588:  uint16(sym_dot),
	589:  uint16(sym_boolean),
	590:  uint16(anon_sym_POUND),
	591:  uint16(sym_number),
	592:  uint16(sym_symbol),
	593:  uint16(anon_sym_POUNDhash),
	594:  uint16(anon_sym_POUNDhasheq),
	595:  uint16(anon_sym_COMMA),
	596:  uint16(anon_sym_POUND_COMMA),
	597:  uint16(anon_sym_POUND_BANG),
	598:  uint16(793),
	599:  uint16(30),
	600:  uint16(aux_sym__skip_token1),
	601:  uint16(aux_sym_comment_token1),
	602:  uint16(anon_sym_POUND_PIPE),
	603:  uint16(anon_sym_POUND_SEMI),
	604:  uint16(sym__line_comment),
	605:  uint16(anon_sym_POUND_LT_LT),
	606:  uint16(aux_sym_regex_token1),
	607:  uint16(anon_sym_DQUOTE),
	608:  uint16(sym_character),
	609:  uint16(sym_keyword),
	610:  uint16(anon_sym_POUND_AMP),
	611:  uint16(anon_sym_LPAREN),
	612:  uint16(anon_sym_RPAREN),
	613:  uint16(anon_sym_LBRACK),
	614:  uint16(anon_sym_RBRACK),
	615:  uint16(anon_sym_LBRACE),
	616:  uint16(anon_sym_RBRACE),
	617:  uint16(anon_sym_POUNDfl),
	618:  uint16(anon_sym_POUNDfx),
	619:  uint16(anon_sym_POUNDs),
	620:  uint16(anon_sym_POUNDhashalw),
	621:  uint16(anon_sym_POUNDhasheqv),
	622:  uint16(anon_sym_SQUOTE),
	623:  uint16(anon_sym_BQUOTE),
	624:  uint16(anon_sym_POUND_SQUOTE),
	625:  uint16(anon_sym_POUND_BQUOTE),
	626:  uint16(anon_sym_COMMA_AT),
	627:  uint16(anon_sym_POUND_COMMA_AT),
	628:  uint16(anon_sym_POUNDreader),
	629:  uint16(anon_sym_POUNDlang),
	630:  uint16(2),
	631:  uint16(799),
	632:  uint16(10),
	633:  uint16(sym_dot),
	634:  uint16(sym_boolean),
	635:  uint16(anon_sym_POUND),
	636:  uint16(sym_number),
	637:  uint16(sym_symbol),
	638:  uint16(anon_sym_POUNDhash),
	639:  uint16(anon_sym_POUNDhasheq),
	640:  uint16(anon_sym_COMMA),
	641:  uint16(anon_sym_POUND_COMMA),
	642:  uint16(anon_sym_POUND_BANG),
	643:  uint16(797),
	644:  uint16(30),
	645:  uint16(aux_sym__skip_token1),
	646:  uint16(aux_sym_comment_token1),
	647:  uint16(anon_sym_POUND_PIPE),
	648:  uint16(anon_sym_POUND_SEMI),
	649:  uint16(sym__line_comment),
	650:  uint16(anon_sym_POUND_LT_LT),
	651:  uint16(aux_sym_regex_token1),
	652:  uint16(anon_sym_DQUOTE),
	653:  uint16(sym_character),
	654:  uint16(sym_keyword),
	655:  uint16(anon_sym_POUND_AMP),
	656:  uint16(anon_sym_LPAREN),
	657:  uint16(anon_sym_RPAREN),
	658:  uint16(anon_sym_LBRACK),
	659:  uint16(anon_sym_RBRACK),
	660:  uint16(anon_sym_LBRACE),
	661:  uint16(anon_sym_RBRACE),
	662:  uint16(anon_sym_POUNDfl),
	663:  uint16(anon_sym_POUNDfx),
	664:  uint16(anon_sym_POUNDs),
	665:  uint16(anon_sym_POUNDhashalw),
	666:  uint16(anon_sym_POUNDhasheqv),
	667:  uint16(anon_sym_SQUOTE),
	668:  uint16(anon_sym_BQUOTE),
	669:  uint16(anon_sym_POUND_SQUOTE),
	670:  uint16(anon_sym_POUND_BQUOTE),
	671:  uint16(anon_sym_COMMA_AT),
	672:  uint16(anon_sym_POUND_COMMA_AT),
	673:  uint16(anon_sym_POUNDreader),
	674:  uint16(anon_sym_POUNDlang),
	675:  uint16(2),
	676:  uint16(803),
	677:  uint16(10),
	678:  uint16(sym_dot),
	679:  uint16(sym_boolean),
	680:  uint16(anon_sym_POUND),
	681:  uint16(sym_number),
	682:  uint16(sym_symbol),
	683:  uint16(anon_sym_POUNDhash),
	684:  uint16(anon_sym_POUNDhasheq),
	685:  uint16(anon_sym_COMMA),
	686:  uint16(anon_sym_POUND_COMMA),
	687:  uint16(anon_sym_POUND_BANG),
	688:  uint16(801),
	689:  uint16(30),
	690:  uint16(aux_sym__skip_token1),
	691:  uint16(aux_sym_comment_token1),
	692:  uint16(anon_sym_POUND_PIPE),
	693:  uint16(anon_sym_POUND_SEMI),
	694:  uint16(sym__line_comment),
	695:  uint16(anon_sym_POUND_LT_LT),
	696:  uint16(aux_sym_regex_token1),
	697:  uint16(anon_sym_DQUOTE),
	698:  uint16(sym_character),
	699:  uint16(sym_keyword),
	700:  uint16(anon_sym_POUND_AMP),
	701:  uint16(anon_sym_LPAREN),
	702:  uint16(anon_sym_RPAREN),
	703:  uint16(anon_sym_LBRACK),
	704:  uint16(anon_sym_RBRACK),
	705:  uint16(anon_sym_LBRACE),
	706:  uint16(anon_sym_RBRACE),
	707:  uint16(anon_sym_POUNDfl),
	708:  uint16(anon_sym_POUNDfx),
	709:  uint16(anon_sym_POUNDs),
	710:  uint16(anon_sym_POUNDhashalw),
	711:  uint16(anon_sym_POUNDhasheqv),
	712:  uint16(anon_sym_SQUOTE),
	713:  uint16(anon_sym_BQUOTE),
	714:  uint16(anon_sym_POUND_SQUOTE),
	715:  uint16(anon_sym_POUND_BQUOTE),
	716:  uint16(anon_sym_COMMA_AT),
	717:  uint16(anon_sym_POUND_COMMA_AT),
	718:  uint16(anon_sym_POUNDreader),
	719:  uint16(anon_sym_POUNDlang),
	720:  uint16(2),
	721:  uint16(807),
	722:  uint16(10),
	723:  uint16(sym_dot),
	724:  uint16(sym_boolean),
	725:  uint16(anon_sym_POUND),
	726:  uint16(sym_number),
	727:  uint16(sym_symbol),
	728:  uint16(anon_sym_POUNDhash),
	729:  uint16(anon_sym_POUNDhasheq),
	730:  uint16(anon_sym_COMMA),
	731:  uint16(anon_sym_POUND_COMMA),
	732:  uint16(anon_sym_POUND_BANG),
	733:  uint16(805),
	734:  uint16(30),
	735:  uint16(aux_sym__skip_token1),
	736:  uint16(aux_sym_comment_token1),
	737:  uint16(anon_sym_POUND_PIPE),
	738:  uint16(anon_sym_POUND_SEMI),
	739:  uint16(sym__line_comment),
	740:  uint16(anon_sym_POUND_LT_LT),
	741:  uint16(aux_sym_regex_token1),
	742:  uint16(anon_sym_DQUOTE),
	743:  uint16(sym_character),
	744:  uint16(sym_keyword),
	745:  uint16(anon_sym_POUND_AMP),
	746:  uint16(anon_sym_LPAREN),
	747:  uint16(anon_sym_RPAREN),
	748:  uint16(anon_sym_LBRACK),
	749:  uint16(anon_sym_RBRACK),
	750:  uint16(anon_sym_LBRACE),
	751:  uint16(anon_sym_RBRACE),
	752:  uint16(anon_sym_POUNDfl),
	753:  uint16(anon_sym_POUNDfx),
	754:  uint16(anon_sym_POUNDs),
	755:  uint16(anon_sym_POUNDhashalw),
	756:  uint16(anon_sym_POUNDhasheqv),
	757:  uint16(anon_sym_SQUOTE),
	758:  uint16(anon_sym_BQUOTE),
	759:  uint16(anon_sym_POUND_SQUOTE),
	760:  uint16(anon_sym_POUND_BQUOTE),
	761:  uint16(anon_sym_COMMA_AT),
	762:  uint16(anon_sym_POUND_COMMA_AT),
	763:  uint16(anon_sym_POUNDreader),
	764:  uint16(anon_sym_POUNDlang),
	765:  uint16(2),
	766:  uint16(811),
	767:  uint16(10),
	768:  uint16(sym_dot),
	769:  uint16(sym_boolean),
	770:  uint16(anon_sym_POUND),
	771:  uint16(sym_number),
	772:  uint16(sym_symbol),
	773:  uint16(anon_sym_POUNDhash),
	774:  uint16(anon_sym_POUNDhasheq),
	775:  uint16(anon_sym_COMMA),
	776:  uint16(anon_sym_POUND_COMMA),
	777:  uint16(anon_sym_POUND_BANG),
	778:  uint16(809),
	779:  uint16(30),
	780:  uint16(aux_sym__skip_token1),
	781:  uint16(aux_sym_comment_token1),
	782:  uint16(anon_sym_POUND_PIPE),
	783:  uint16(anon_sym_POUND_SEMI),
	784:  uint16(sym__line_comment),
	785:  uint16(anon_sym_POUND_LT_LT),
	786:  uint16(aux_sym_regex_token1),
	787:  uint16(anon_sym_DQUOTE),
	788:  uint16(sym_character),
	789:  uint16(sym_keyword),
	790:  uint16(anon_sym_POUND_AMP),
	791:  uint16(anon_sym_LPAREN),
	792:  uint16(anon_sym_RPAREN),
	793:  uint16(anon_sym_LBRACK),
	794:  uint16(anon_sym_RBRACK),
	795:  uint16(anon_sym_LBRACE),
	796:  uint16(anon_sym_RBRACE),
	797:  uint16(anon_sym_POUNDfl),
	798:  uint16(anon_sym_POUNDfx),
	799:  uint16(anon_sym_POUNDs),
	800:  uint16(anon_sym_POUNDhashalw),
	801:  uint16(anon_sym_POUNDhasheqv),
	802:  uint16(anon_sym_SQUOTE),
	803:  uint16(anon_sym_BQUOTE),
	804:  uint16(anon_sym_POUND_SQUOTE),
	805:  uint16(anon_sym_POUND_BQUOTE),
	806:  uint16(anon_sym_COMMA_AT),
	807:  uint16(anon_sym_POUND_COMMA_AT),
	808:  uint16(anon_sym_POUNDreader),
	809:  uint16(anon_sym_POUNDlang),
	810:  uint16(2),
	811:  uint16(815),
	812:  uint16(10),
	813:  uint16(sym_dot),
	814:  uint16(sym_boolean),
	815:  uint16(anon_sym_POUND),
	816:  uint16(sym_number),
	817:  uint16(sym_symbol),
	818:  uint16(anon_sym_POUNDhash),
	819:  uint16(anon_sym_POUNDhasheq),
	820:  uint16(anon_sym_COMMA),
	821:  uint16(anon_sym_POUND_COMMA),
	822:  uint16(anon_sym_POUND_BANG),
	823:  uint16(813),
	824:  uint16(30),
	825:  uint16(aux_sym__skip_token1),
	826:  uint16(aux_sym_comment_token1),
	827:  uint16(anon_sym_POUND_PIPE),
	828:  uint16(anon_sym_POUND_SEMI),
	829:  uint16(sym__line_comment),
	830:  uint16(anon_sym_POUND_LT_LT),
	831:  uint16(aux_sym_regex_token1),
	832:  uint16(anon_sym_DQUOTE),
	833:  uint16(sym_character),
	834:  uint16(sym_keyword),
	835:  uint16(anon_sym_POUND_AMP),
	836:  uint16(anon_sym_LPAREN),
	837:  uint16(anon_sym_RPAREN),
	838:  uint16(anon_sym_LBRACK),
	839:  uint16(anon_sym_RBRACK),
	840:  uint16(anon_sym_LBRACE),
	841:  uint16(anon_sym_RBRACE),
	842:  uint16(anon_sym_POUNDfl),
	843:  uint16(anon_sym_POUNDfx),
	844:  uint16(anon_sym_POUNDs),
	845:  uint16(anon_sym_POUNDhashalw),
	846:  uint16(anon_sym_POUNDhasheqv),
	847:  uint16(anon_sym_SQUOTE),
	848:  uint16(anon_sym_BQUOTE),
	849:  uint16(anon_sym_POUND_SQUOTE),
	850:  uint16(anon_sym_POUND_BQUOTE),
	851:  uint16(anon_sym_COMMA_AT),
	852:  uint16(anon_sym_POUND_COMMA_AT),
	853:  uint16(anon_sym_POUNDreader),
	854:  uint16(anon_sym_POUNDlang),
	855:  uint16(2),
	856:  uint16(819),
	857:  uint16(10),
	858:  uint16(sym_dot),
	859:  uint16(sym_boolean),
	860:  uint16(anon_sym_POUND),
	861:  uint16(sym_number),
	862:  uint16(sym_symbol),
	863:  uint16(anon_sym_POUNDhash),
	864:  uint16(anon_sym_POUNDhasheq),
	865:  uint16(anon_sym_COMMA),
	866:  uint16(anon_sym_POUND_COMMA),
	867:  uint16(anon_sym_POUND_BANG),
	868:  uint16(817),
	869:  uint16(30),
	870:  uint16(aux_sym__skip_token1),
	871:  uint16(aux_sym_comment_token1),
	872:  uint16(anon_sym_POUND_PIPE),
	873:  uint16(anon_sym_POUND_SEMI),
	874:  uint16(sym__line_comment),
	875:  uint16(anon_sym_POUND_LT_LT),
	876:  uint16(aux_sym_regex_token1),
	877:  uint16(anon_sym_DQUOTE),
	878:  uint16(sym_character),
	879:  uint16(sym_keyword),
	880:  uint16(anon_sym_POUND_AMP),
	881:  uint16(anon_sym_LPAREN),
	882:  uint16(anon_sym_RPAREN),
	883:  uint16(anon_sym_LBRACK),
	884:  uint16(anon_sym_RBRACK),
	885:  uint16(anon_sym_LBRACE),
	886:  uint16(anon_sym_RBRACE),
	887:  uint16(anon_sym_POUNDfl),
	888:  uint16(anon_sym_POUNDfx),
	889:  uint16(anon_sym_POUNDs),
	890:  uint16(anon_sym_POUNDhashalw),
	891:  uint16(anon_sym_POUNDhasheqv),
	892:  uint16(anon_sym_SQUOTE),
	893:  uint16(anon_sym_BQUOTE),
	894:  uint16(anon_sym_POUND_SQUOTE),
	895:  uint16(anon_sym_POUND_BQUOTE),
	896:  uint16(anon_sym_COMMA_AT),
	897:  uint16(anon_sym_POUND_COMMA_AT),
	898:  uint16(anon_sym_POUNDreader),
	899:  uint16(anon_sym_POUNDlang),
	900:  uint16(2),
	901:  uint16(823),
	902:  uint16(10),
	903:  uint16(sym_dot),
	904:  uint16(sym_boolean),
	905:  uint16(anon_sym_POUND),
	906:  uint16(sym_number),
	907:  uint16(sym_symbol),
	908:  uint16(anon_sym_POUNDhash),
	909:  uint16(anon_sym_POUNDhasheq),
	910:  uint16(anon_sym_COMMA),
	911:  uint16(anon_sym_POUND_COMMA),
	912:  uint16(anon_sym_POUND_BANG),
	913:  uint16(821),
	914:  uint16(30),
	915:  uint16(aux_sym__skip_token1),
	916:  uint16(aux_sym_comment_token1),
	917:  uint16(anon_sym_POUND_PIPE),
	918:  uint16(anon_sym_POUND_SEMI),
	919:  uint16(sym__line_comment),
	920:  uint16(anon_sym_POUND_LT_LT),
	921:  uint16(aux_sym_regex_token1),
	922:  uint16(anon_sym_DQUOTE),
	923:  uint16(sym_character),
	924:  uint16(sym_keyword),
	925:  uint16(anon_sym_POUND_AMP),
	926:  uint16(anon_sym_LPAREN),
	927:  uint16(anon_sym_RPAREN),
	928:  uint16(anon_sym_LBRACK),
	929:  uint16(anon_sym_RBRACK),
	930:  uint16(anon_sym_LBRACE),
	931:  uint16(anon_sym_RBRACE),
	932:  uint16(anon_sym_POUNDfl),
	933:  uint16(anon_sym_POUNDfx),
	934:  uint16(anon_sym_POUNDs),
	935:  uint16(anon_sym_POUNDhashalw),
	936:  uint16(anon_sym_POUNDhasheqv),
	937:  uint16(anon_sym_SQUOTE),
	938:  uint16(anon_sym_BQUOTE),
	939:  uint16(anon_sym_POUND_SQUOTE),
	940:  uint16(anon_sym_POUND_BQUOTE),
	941:  uint16(anon_sym_COMMA_AT),
	942:  uint16(anon_sym_POUND_COMMA_AT),
	943:  uint16(anon_sym_POUNDreader),
	944:  uint16(anon_sym_POUNDlang),
	945:  uint16(2),
	946:  uint16(827),
	947:  uint16(10),
	948:  uint16(sym_dot),
	949:  uint16(sym_boolean),
	950:  uint16(anon_sym_POUND),
	951:  uint16(sym_number),
	952:  uint16(sym_symbol),
	953:  uint16(anon_sym_POUNDhash),
	954:  uint16(anon_sym_POUNDhasheq),
	955:  uint16(anon_sym_COMMA),
	956:  uint16(anon_sym_POUND_COMMA),
	957:  uint16(anon_sym_POUND_BANG),
	958:  uint16(825),
	959:  uint16(30),
	960:  uint16(aux_sym__skip_token1),
	961:  uint16(aux_sym_comment_token1),
	962:  uint16(anon_sym_POUND_PIPE),
	963:  uint16(anon_sym_POUND_SEMI),
	964:  uint16(sym__line_comment),
	965:  uint16(anon_sym_POUND_LT_LT),
	966:  uint16(aux_sym_regex_token1),
	967:  uint16(anon_sym_DQUOTE),
	968:  uint16(sym_character),
	969:  uint16(sym_keyword),
	970:  uint16(anon_sym_POUND_AMP),
	971:  uint16(anon_sym_LPAREN),
	972:  uint16(anon_sym_RPAREN),
	973:  uint16(anon_sym_LBRACK),
	974:  uint16(anon_sym_RBRACK),
	975:  uint16(anon_sym_LBRACE),
	976:  uint16(anon_sym_RBRACE),
	977:  uint16(anon_sym_POUNDfl),
	978:  uint16(anon_sym_POUNDfx),
	979:  uint16(anon_sym_POUNDs),
	980:  uint16(anon_sym_POUNDhashalw),
	981:  uint16(anon_sym_POUNDhasheqv),
	982:  uint16(anon_sym_SQUOTE),
	983:  uint16(anon_sym_BQUOTE),
	984:  uint16(anon_sym_POUND_SQUOTE),
	985:  uint16(anon_sym_POUND_BQUOTE),
	986:  uint16(anon_sym_COMMA_AT),
	987:  uint16(anon_sym_POUND_COMMA_AT),
	988:  uint16(anon_sym_POUNDreader),
	989:  uint16(anon_sym_POUNDlang),
	990:  uint16(2),
	991:  uint16(831),
	992:  uint16(10),
	993:  uint16(sym_dot),
	994:  uint16(sym_boolean),
	995:  uint16(anon_sym_POUND),
	996:  uint16(sym_number),
	997:  uint16(sym_symbol),
	998:  uint16(anon_sym_POUNDhash),
	999:  uint16(anon_sym_POUNDhasheq),
	1000: uint16(anon_sym_COMMA),
	1001: uint16(anon_sym_POUND_COMMA),
	1002: uint16(anon_sym_POUND_BANG),
	1003: uint16(829),
	1004: uint16(30),
	1005: uint16(aux_sym__skip_token1),
	1006: uint16(aux_sym_comment_token1),
	1007: uint16(anon_sym_POUND_PIPE),
	1008: uint16(anon_sym_POUND_SEMI),
	1009: uint16(sym__line_comment),
	1010: uint16(anon_sym_POUND_LT_LT),
	1011: uint16(aux_sym_regex_token1),
	1012: uint16(anon_sym_DQUOTE),
	1013: uint16(sym_character),
	1014: uint16(sym_keyword),
	1015: uint16(anon_sym_POUND_AMP),
	1016: uint16(anon_sym_LPAREN),
	1017: uint16(anon_sym_RPAREN),
	1018: uint16(anon_sym_LBRACK),
	1019: uint16(anon_sym_RBRACK),
	1020: uint16(anon_sym_LBRACE),
	1021: uint16(anon_sym_RBRACE),
	1022: uint16(anon_sym_POUNDfl),
	1023: uint16(anon_sym_POUNDfx),
	1024: uint16(anon_sym_POUNDs),
	1025: uint16(anon_sym_POUNDhashalw),
	1026: uint16(anon_sym_POUNDhasheqv),
	1027: uint16(anon_sym_SQUOTE),
	1028: uint16(anon_sym_BQUOTE),
	1029: uint16(anon_sym_POUND_SQUOTE),
	1030: uint16(anon_sym_POUND_BQUOTE),
	1031: uint16(anon_sym_COMMA_AT),
	1032: uint16(anon_sym_POUND_COMMA_AT),
	1033: uint16(anon_sym_POUNDreader),
	1034: uint16(anon_sym_POUNDlang),
	1035: uint16(2),
	1036: uint16(835),
	1037: uint16(10),
	1038: uint16(sym_dot),
	1039: uint16(sym_boolean),
	1040: uint16(anon_sym_POUND),
	1041: uint16(sym_number),
	1042: uint16(sym_symbol),
	1043: uint16(anon_sym_POUNDhash),
	1044: uint16(anon_sym_POUNDhasheq),
	1045: uint16(anon_sym_COMMA),
	1046: uint16(anon_sym_POUND_COMMA),
	1047: uint16(anon_sym_POUND_BANG),
	1048: uint16(833),
	1049: uint16(30),
	1050: uint16(aux_sym__skip_token1),
	1051: uint16(aux_sym_comment_token1),
	1052: uint16(anon_sym_POUND_PIPE),
	1053: uint16(anon_sym_POUND_SEMI),
	1054: uint16(sym__line_comment),
	1055: uint16(anon_sym_POUND_LT_LT),
	1056: uint16(aux_sym_regex_token1),
	1057: uint16(anon_sym_DQUOTE),
	1058: uint16(sym_character),
	1059: uint16(sym_keyword),
	1060: uint16(anon_sym_POUND_AMP),
	1061: uint16(anon_sym_LPAREN),
	1062: uint16(anon_sym_RPAREN),
	1063: uint16(anon_sym_LBRACK),
	1064: uint16(anon_sym_RBRACK),
	1065: uint16(anon_sym_LBRACE),
	1066: uint16(anon_sym_RBRACE),
	1067: uint16(anon_sym_POUNDfl),
	1068: uint16(anon_sym_POUNDfx),
	1069: uint16(anon_sym_POUNDs),
	1070: uint16(anon_sym_POUNDhashalw),
	1071: uint16(anon_sym_POUNDhasheqv),
	1072: uint16(anon_sym_SQUOTE),
	1073: uint16(anon_sym_BQUOTE),
	1074: uint16(anon_sym_POUND_SQUOTE),
	1075: uint16(anon_sym_POUND_BQUOTE),
	1076: uint16(anon_sym_COMMA_AT),
	1077: uint16(anon_sym_POUND_COMMA_AT),
	1078: uint16(anon_sym_POUNDreader),
	1079: uint16(anon_sym_POUNDlang),
	1080: uint16(2),
	1081: uint16(839),
	1082: uint16(10),
	1083: uint16(sym_dot),
	1084: uint16(sym_boolean),
	1085: uint16(anon_sym_POUND),
	1086: uint16(sym_number),
	1087: uint16(sym_symbol),
	1088: uint16(anon_sym_POUNDhash),
	1089: uint16(anon_sym_POUNDhasheq),
	1090: uint16(anon_sym_COMMA),
	1091: uint16(anon_sym_POUND_COMMA),
	1092: uint16(anon_sym_POUND_BANG),
	1093: uint16(837),
	1094: uint16(30),
	1095: uint16(aux_sym__skip_token1),
	1096: uint16(aux_sym_comment_token1),
	1097: uint16(anon_sym_POUND_PIPE),
	1098: uint16(anon_sym_POUND_SEMI),
	1099: uint16(sym__line_comment),
	1100: uint16(anon_sym_POUND_LT_LT),
	1101: uint16(aux_sym_regex_token1),
	1102: uint16(anon_sym_DQUOTE),
	1103: uint16(sym_character),
	1104: uint16(sym_keyword),
	1105: uint16(anon_sym_POUND_AMP),
	1106: uint16(anon_sym_LPAREN),
	1107: uint16(anon_sym_RPAREN),
	1108: uint16(anon_sym_LBRACK),
	1109: uint16(anon_sym_RBRACK),
	1110: uint16(anon_sym_LBRACE),
	1111: uint16(anon_sym_RBRACE),
	1112: uint16(anon_sym_POUNDfl),
	1113: uint16(anon_sym_POUNDfx),
	1114: uint16(anon_sym_POUNDs),
	1115: uint16(anon_sym_POUNDhashalw),
	1116: uint16(anon_sym_POUNDhasheqv),
	1117: uint16(anon_sym_SQUOTE),
	1118: uint16(anon_sym_BQUOTE),
	1119: uint16(anon_sym_POUND_SQUOTE),
	1120: uint16(anon_sym_POUND_BQUOTE),
	1121: uint16(anon_sym_COMMA_AT),
	1122: uint16(anon_sym_POUND_COMMA_AT),
	1123: uint16(anon_sym_POUNDreader),
	1124: uint16(anon_sym_POUNDlang),
	1125: uint16(2),
	1126: uint16(843),
	1127: uint16(10),
	1128: uint16(sym_dot),
	1129: uint16(sym_boolean),
	1130: uint16(anon_sym_POUND),
	1131: uint16(sym_number),
	1132: uint16(sym_symbol),
	1133: uint16(anon_sym_POUNDhash),
	1134: uint16(anon_sym_POUNDhasheq),
	1135: uint16(anon_sym_COMMA),
	1136: uint16(anon_sym_POUND_COMMA),
	1137: uint16(anon_sym_POUND_BANG),
	1138: uint16(841),
	1139: uint16(30),
	1140: uint16(aux_sym__skip_token1),
	1141: uint16(aux_sym_comment_token1),
	1142: uint16(anon_sym_POUND_PIPE),
	1143: uint16(anon_sym_POUND_SEMI),
	1144: uint16(sym__line_comment),
	1145: uint16(anon_sym_POUND_LT_LT),
	1146: uint16(aux_sym_regex_token1),
	1147: uint16(anon_sym_DQUOTE),
	1148: uint16(sym_character),
	1149: uint16(sym_keyword),
	1150: uint16(anon_sym_POUND_AMP),
	1151: uint16(anon_sym_LPAREN),
	1152: uint16(anon_sym_RPAREN),
	1153: uint16(anon_sym_LBRACK),
	1154: uint16(anon_sym_RBRACK),
	1155: uint16(anon_sym_LBRACE),
	1156: uint16(anon_sym_RBRACE),
	1157: uint16(anon_sym_POUNDfl),
	1158: uint16(anon_sym_POUNDfx),
	1159: uint16(anon_sym_POUNDs),
	1160: uint16(anon_sym_POUNDhashalw),
	1161: uint16(anon_sym_POUNDhasheqv),
	1162: uint16(anon_sym_SQUOTE),
	1163: uint16(anon_sym_BQUOTE),
	1164: uint16(anon_sym_POUND_SQUOTE),
	1165: uint16(anon_sym_POUND_BQUOTE),
	1166: uint16(anon_sym_COMMA_AT),
	1167: uint16(anon_sym_POUND_COMMA_AT),
	1168: uint16(anon_sym_POUNDreader),
	1169: uint16(anon_sym_POUNDlang),
	1170: uint16(2),
	1171: uint16(847),
	1172: uint16(10),
	1173: uint16(sym_dot),
	1174: uint16(sym_boolean),
	1175: uint16(anon_sym_POUND),
	1176: uint16(sym_number),
	1177: uint16(sym_symbol),
	1178: uint16(anon_sym_POUNDhash),
	1179: uint16(anon_sym_POUNDhasheq),
	1180: uint16(anon_sym_COMMA),
	1181: uint16(anon_sym_POUND_COMMA),
	1182: uint16(anon_sym_POUND_BANG),
	1183: uint16(845),
	1184: uint16(30),
	1185: uint16(aux_sym__skip_token1),
	1186: uint16(aux_sym_comment_token1),
	1187: uint16(anon_sym_POUND_PIPE),
	1188: uint16(anon_sym_POUND_SEMI),
	1189: uint16(sym__line_comment),
	1190: uint16(anon_sym_POUND_LT_LT),
	1191: uint16(aux_sym_regex_token1),
	1192: uint16(anon_sym_DQUOTE),
	1193: uint16(sym_character),
	1194: uint16(sym_keyword),
	1195: uint16(anon_sym_POUND_AMP),
	1196: uint16(anon_sym_LPAREN),
	1197: uint16(anon_sym_RPAREN),
	1198: uint16(anon_sym_LBRACK),
	1199: uint16(anon_sym_RBRACK),
	1200: uint16(anon_sym_LBRACE),
	1201: uint16(anon_sym_RBRACE),
	1202: uint16(anon_sym_POUNDfl),
	1203: uint16(anon_sym_POUNDfx),
	1204: uint16(anon_sym_POUNDs),
	1205: uint16(anon_sym_POUNDhashalw),
	1206: uint16(anon_sym_POUNDhasheqv),
	1207: uint16(anon_sym_SQUOTE),
	1208: uint16(anon_sym_BQUOTE),
	1209: uint16(anon_sym_POUND_SQUOTE),
	1210: uint16(anon_sym_POUND_BQUOTE),
	1211: uint16(anon_sym_COMMA_AT),
	1212: uint16(anon_sym_POUND_COMMA_AT),
	1213: uint16(anon_sym_POUNDreader),
	1214: uint16(anon_sym_POUNDlang),
	1215: uint16(2),
	1216: uint16(851),
	1217: uint16(10),
	1218: uint16(sym_dot),
	1219: uint16(sym_boolean),
	1220: uint16(anon_sym_POUND),
	1221: uint16(sym_number),
	1222: uint16(sym_symbol),
	1223: uint16(anon_sym_POUNDhash),
	1224: uint16(anon_sym_POUNDhasheq),
	1225: uint16(anon_sym_COMMA),
	1226: uint16(anon_sym_POUND_COMMA),
	1227: uint16(anon_sym_POUND_BANG),
	1228: uint16(849),
	1229: uint16(30),
	1230: uint16(aux_sym__skip_token1),
	1231: uint16(aux_sym_comment_token1),
	1232: uint16(anon_sym_POUND_PIPE),
	1233: uint16(anon_sym_POUND_SEMI),
	1234: uint16(sym__line_comment),
	1235: uint16(anon_sym_POUND_LT_LT),
	1236: uint16(aux_sym_regex_token1),
	1237: uint16(anon_sym_DQUOTE),
	1238: uint16(sym_character),
	1239: uint16(sym_keyword),
	1240: uint16(anon_sym_POUND_AMP),
	1241: uint16(anon_sym_LPAREN),
	1242: uint16(anon_sym_RPAREN),
	1243: uint16(anon_sym_LBRACK),
	1244: uint16(anon_sym_RBRACK),
	1245: uint16(anon_sym_LBRACE),
	1246: uint16(anon_sym_RBRACE),
	1247: uint16(anon_sym_POUNDfl),
	1248: uint16(anon_sym_POUNDfx),
	1249: uint16(anon_sym_POUNDs),
	1250: uint16(anon_sym_POUNDhashalw),
	1251: uint16(anon_sym_POUNDhasheqv),
	1252: uint16(anon_sym_SQUOTE),
	1253: uint16(anon_sym_BQUOTE),
	1254: uint16(anon_sym_POUND_SQUOTE),
	1255: uint16(anon_sym_POUND_BQUOTE),
	1256: uint16(anon_sym_COMMA_AT),
	1257: uint16(anon_sym_POUND_COMMA_AT),
	1258: uint16(anon_sym_POUNDreader),
	1259: uint16(anon_sym_POUNDlang),
	1260: uint16(2),
	1261: uint16(855),
	1262: uint16(10),
	1263: uint16(sym_dot),
	1264: uint16(sym_boolean),
	1265: uint16(anon_sym_POUND),
	1266: uint16(sym_number),
	1267: uint16(sym_symbol),
	1268: uint16(anon_sym_POUNDhash),
	1269: uint16(anon_sym_POUNDhasheq),
	1270: uint16(anon_sym_COMMA),
	1271: uint16(anon_sym_POUND_COMMA),
	1272: uint16(anon_sym_POUND_BANG),
	1273: uint16(853),
	1274: uint16(30),
	1275: uint16(aux_sym__skip_token1),
	1276: uint16(aux_sym_comment_token1),
	1277: uint16(anon_sym_POUND_PIPE),
	1278: uint16(anon_sym_POUND_SEMI),
	1279: uint16(sym__line_comment),
	1280: uint16(anon_sym_POUND_LT_LT),
	1281: uint16(aux_sym_regex_token1),
	1282: uint16(anon_sym_DQUOTE),
	1283: uint16(sym_character),
	1284: uint16(sym_keyword),
	1285: uint16(anon_sym_POUND_AMP),
	1286: uint16(anon_sym_LPAREN),
	1287: uint16(anon_sym_RPAREN),
	1288: uint16(anon_sym_LBRACK),
	1289: uint16(anon_sym_RBRACK),
	1290: uint16(anon_sym_LBRACE),
	1291: uint16(anon_sym_RBRACE),
	1292: uint16(anon_sym_POUNDfl),
	1293: uint16(anon_sym_POUNDfx),
	1294: uint16(anon_sym_POUNDs),
	1295: uint16(anon_sym_POUNDhashalw),
	1296: uint16(anon_sym_POUNDhasheqv),
	1297: uint16(anon_sym_SQUOTE),
	1298: uint16(anon_sym_BQUOTE),
	1299: uint16(anon_sym_POUND_SQUOTE),
	1300: uint16(anon_sym_POUND_BQUOTE),
	1301: uint16(anon_sym_COMMA_AT),
	1302: uint16(anon_sym_POUND_COMMA_AT),
	1303: uint16(anon_sym_POUNDreader),
	1304: uint16(anon_sym_POUNDlang),
	1305: uint16(2),
	1306: uint16(859),
	1307: uint16(10),
	1308: uint16(sym_dot),
	1309: uint16(sym_boolean),
	1310: uint16(anon_sym_POUND),
	1311: uint16(sym_number),
	1312: uint16(sym_symbol),
	1313: uint16(anon_sym_POUNDhash),
	1314: uint16(anon_sym_POUNDhasheq),
	1315: uint16(anon_sym_COMMA),
	1316: uint16(anon_sym_POUND_COMMA),
	1317: uint16(anon_sym_POUND_BANG),
	1318: uint16(857),
	1319: uint16(30),
	1320: uint16(aux_sym__skip_token1),
	1321: uint16(aux_sym_comment_token1),
	1322: uint16(anon_sym_POUND_PIPE),
	1323: uint16(anon_sym_POUND_SEMI),
	1324: uint16(sym__line_comment),
	1325: uint16(anon_sym_POUND_LT_LT),
	1326: uint16(aux_sym_regex_token1),
	1327: uint16(anon_sym_DQUOTE),
	1328: uint16(sym_character),
	1329: uint16(sym_keyword),
	1330: uint16(anon_sym_POUND_AMP),
	1331: uint16(anon_sym_LPAREN),
	1332: uint16(anon_sym_RPAREN),
	1333: uint16(anon_sym_LBRACK),
	1334: uint16(anon_sym_RBRACK),
	1335: uint16(anon_sym_LBRACE),
	1336: uint16(anon_sym_RBRACE),
	1337: uint16(anon_sym_POUNDfl),
	1338: uint16(anon_sym_POUNDfx),
	1339: uint16(anon_sym_POUNDs),
	1340: uint16(anon_sym_POUNDhashalw),
	1341: uint16(anon_sym_POUNDhasheqv),
	1342: uint16(anon_sym_SQUOTE),
	1343: uint16(anon_sym_BQUOTE),
	1344: uint16(anon_sym_POUND_SQUOTE),
	1345: uint16(anon_sym_POUND_BQUOTE),
	1346: uint16(anon_sym_COMMA_AT),
	1347: uint16(anon_sym_POUND_COMMA_AT),
	1348: uint16(anon_sym_POUNDreader),
	1349: uint16(anon_sym_POUNDlang),
	1350: uint16(2),
	1351: uint16(863),
	1352: uint16(10),
	1353: uint16(sym_dot),
	1354: uint16(sym_boolean),
	1355: uint16(anon_sym_POUND),
	1356: uint16(sym_number),
	1357: uint16(sym_symbol),
	1358: uint16(anon_sym_POUNDhash),
	1359: uint16(anon_sym_POUNDhasheq),
	1360: uint16(anon_sym_COMMA),
	1361: uint16(anon_sym_POUND_COMMA),
	1362: uint16(anon_sym_POUND_BANG),
	1363: uint16(861),
	1364: uint16(30),
	1365: uint16(aux_sym__skip_token1),
	1366: uint16(aux_sym_comment_token1),
	1367: uint16(anon_sym_POUND_PIPE),
	1368: uint16(anon_sym_POUND_SEMI),
	1369: uint16(sym__line_comment),
	1370: uint16(anon_sym_POUND_LT_LT),
	1371: uint16(aux_sym_regex_token1),
	1372: uint16(anon_sym_DQUOTE),
	1373: uint16(sym_character),
	1374: uint16(sym_keyword),
	1375: uint16(anon_sym_POUND_AMP),
	1376: uint16(anon_sym_LPAREN),
	1377: uint16(anon_sym_RPAREN),
	1378: uint16(anon_sym_LBRACK),
	1379: uint16(anon_sym_RBRACK),
	1380: uint16(anon_sym_LBRACE),
	1381: uint16(anon_sym_RBRACE),
	1382: uint16(anon_sym_POUNDfl),
	1383: uint16(anon_sym_POUNDfx),
	1384: uint16(anon_sym_POUNDs),
	1385: uint16(anon_sym_POUNDhashalw),
	1386: uint16(anon_sym_POUNDhasheqv),
	1387: uint16(anon_sym_SQUOTE),
	1388: uint16(anon_sym_BQUOTE),
	1389: uint16(anon_sym_POUND_SQUOTE),
	1390: uint16(anon_sym_POUND_BQUOTE),
	1391: uint16(anon_sym_COMMA_AT),
	1392: uint16(anon_sym_POUND_COMMA_AT),
	1393: uint16(anon_sym_POUNDreader),
	1394: uint16(anon_sym_POUNDlang),
	1395: uint16(2),
	1396: uint16(867),
	1397: uint16(10),
	1398: uint16(sym_dot),
	1399: uint16(sym_boolean),
	1400: uint16(anon_sym_POUND),
	1401: uint16(sym_number),
	1402: uint16(sym_symbol),
	1403: uint16(anon_sym_POUNDhash),
	1404: uint16(anon_sym_POUNDhasheq),
	1405: uint16(anon_sym_COMMA),
	1406: uint16(anon_sym_POUND_COMMA),
	1407: uint16(anon_sym_POUND_BANG),
	1408: uint16(865),
	1409: uint16(30),
	1410: uint16(aux_sym__skip_token1),
	1411: uint16(aux_sym_comment_token1),
	1412: uint16(anon_sym_POUND_PIPE),
	1413: uint16(anon_sym_POUND_SEMI),
	1414: uint16(sym__line_comment),
	1415: uint16(anon_sym_POUND_LT_LT),
	1416: uint16(aux_sym_regex_token1),
	1417: uint16(anon_sym_DQUOTE),
	1418: uint16(sym_character),
	1419: uint16(sym_keyword),
	1420: uint16(anon_sym_POUND_AMP),
	1421: uint16(anon_sym_LPAREN),
	1422: uint16(anon_sym_RPAREN),
	1423: uint16(anon_sym_LBRACK),
	1424: uint16(anon_sym_RBRACK),
	1425: uint16(anon_sym_LBRACE),
	1426: uint16(anon_sym_RBRACE),
	1427: uint16(anon_sym_POUNDfl),
	1428: uint16(anon_sym_POUNDfx),
	1429: uint16(anon_sym_POUNDs),
	1430: uint16(anon_sym_POUNDhashalw),
	1431: uint16(anon_sym_POUNDhasheqv),
	1432: uint16(anon_sym_SQUOTE),
	1433: uint16(anon_sym_BQUOTE),
	1434: uint16(anon_sym_POUND_SQUOTE),
	1435: uint16(anon_sym_POUND_BQUOTE),
	1436: uint16(anon_sym_COMMA_AT),
	1437: uint16(anon_sym_POUND_COMMA_AT),
	1438: uint16(anon_sym_POUNDreader),
	1439: uint16(anon_sym_POUNDlang),
	1440: uint16(2),
	1441: uint16(871),
	1442: uint16(10),
	1443: uint16(sym_dot),
	1444: uint16(sym_boolean),
	1445: uint16(anon_sym_POUND),
	1446: uint16(sym_number),
	1447: uint16(sym_symbol),
	1448: uint16(anon_sym_POUNDhash),
	1449: uint16(anon_sym_POUNDhasheq),
	1450: uint16(anon_sym_COMMA),
	1451: uint16(anon_sym_POUND_COMMA),
	1452: uint16(anon_sym_POUND_BANG),
	1453: uint16(869),
	1454: uint16(30),
	1455: uint16(aux_sym__skip_token1),
	1456: uint16(aux_sym_comment_token1),
	1457: uint16(anon_sym_POUND_PIPE),
	1458: uint16(anon_sym_POUND_SEMI),
	1459: uint16(sym__line_comment),
	1460: uint16(anon_sym_POUND_LT_LT),
	1461: uint16(aux_sym_regex_token1),
	1462: uint16(anon_sym_DQUOTE),
	1463: uint16(sym_character),
	1464: uint16(sym_keyword),
	1465: uint16(anon_sym_POUND_AMP),
	1466: uint16(anon_sym_LPAREN),
	1467: uint16(anon_sym_RPAREN),
	1468: uint16(anon_sym_LBRACK),
	1469: uint16(anon_sym_RBRACK),
	1470: uint16(anon_sym_LBRACE),
	1471: uint16(anon_sym_RBRACE),
	1472: uint16(anon_sym_POUNDfl),
	1473: uint16(anon_sym_POUNDfx),
	1474: uint16(anon_sym_POUNDs),
	1475: uint16(anon_sym_POUNDhashalw),
	1476: uint16(anon_sym_POUNDhasheqv),
	1477: uint16(anon_sym_SQUOTE),
	1478: uint16(anon_sym_BQUOTE),
	1479: uint16(anon_sym_POUND_SQUOTE),
	1480: uint16(anon_sym_POUND_BQUOTE),
	1481: uint16(anon_sym_COMMA_AT),
	1482: uint16(anon_sym_POUND_COMMA_AT),
	1483: uint16(anon_sym_POUNDreader),
	1484: uint16(anon_sym_POUNDlang),
	1485: uint16(2),
	1486: uint16(875),
	1487: uint16(10),
	1488: uint16(sym_dot),
	1489: uint16(sym_boolean),
	1490: uint16(anon_sym_POUND),
	1491: uint16(sym_number),
	1492: uint16(sym_symbol),
	1493: uint16(anon_sym_POUNDhash),
	1494: uint16(anon_sym_POUNDhasheq),
	1495: uint16(anon_sym_COMMA),
	1496: uint16(anon_sym_POUND_COMMA),
	1497: uint16(anon_sym_POUND_BANG),
	1498: uint16(873),
	1499: uint16(30),
	1500: uint16(aux_sym__skip_token1),
	1501: uint16(aux_sym_comment_token1),
	1502: uint16(anon_sym_POUND_PIPE),
	1503: uint16(anon_sym_POUND_SEMI),
	1504: uint16(sym__line_comment),
	1505: uint16(anon_sym_POUND_LT_LT),
	1506: uint16(aux_sym_regex_token1),
	1507: uint16(anon_sym_DQUOTE),
	1508: uint16(sym_character),
	1509: uint16(sym_keyword),
	1510: uint16(anon_sym_POUND_AMP),
	1511: uint16(anon_sym_LPAREN),
	1512: uint16(anon_sym_RPAREN),
	1513: uint16(anon_sym_LBRACK),
	1514: uint16(anon_sym_RBRACK),
	1515: uint16(anon_sym_LBRACE),
	1516: uint16(anon_sym_RBRACE),
	1517: uint16(anon_sym_POUNDfl),
	1518: uint16(anon_sym_POUNDfx),
	1519: uint16(anon_sym_POUNDs),
	1520: uint16(anon_sym_POUNDhashalw),
	1521: uint16(anon_sym_POUNDhasheqv),
	1522: uint16(anon_sym_SQUOTE),
	1523: uint16(anon_sym_BQUOTE),
	1524: uint16(anon_sym_POUND_SQUOTE),
	1525: uint16(anon_sym_POUND_BQUOTE),
	1526: uint16(anon_sym_COMMA_AT),
	1527: uint16(anon_sym_POUND_COMMA_AT),
	1528: uint16(anon_sym_POUNDreader),
	1529: uint16(anon_sym_POUNDlang),
	1530: uint16(2),
	1531: uint16(879),
	1532: uint16(10),
	1533: uint16(sym_dot),
	1534: uint16(sym_boolean),
	1535: uint16(anon_sym_POUND),
	1536: uint16(sym_number),
	1537: uint16(sym_symbol),
	1538: uint16(anon_sym_POUNDhash),
	1539: uint16(anon_sym_POUNDhasheq),
	1540: uint16(anon_sym_COMMA),
	1541: uint16(anon_sym_POUND_COMMA),
	1542: uint16(anon_sym_POUND_BANG),
	1543: uint16(877),
	1544: uint16(30),
	1545: uint16(aux_sym__skip_token1),
	1546: uint16(aux_sym_comment_token1),
	1547: uint16(anon_sym_POUND_PIPE),
	1548: uint16(anon_sym_POUND_SEMI),
	1549: uint16(sym__line_comment),
	1550: uint16(anon_sym_POUND_LT_LT),
	1551: uint16(aux_sym_regex_token1),
	1552: uint16(anon_sym_DQUOTE),
	1553: uint16(sym_character),
	1554: uint16(sym_keyword),
	1555: uint16(anon_sym_POUND_AMP),
	1556: uint16(anon_sym_LPAREN),
	1557: uint16(anon_sym_RPAREN),
	1558: uint16(anon_sym_LBRACK),
	1559: uint16(anon_sym_RBRACK),
	1560: uint16(anon_sym_LBRACE),
	1561: uint16(anon_sym_RBRACE),
	1562: uint16(anon_sym_POUNDfl),
	1563: uint16(anon_sym_POUNDfx),
	1564: uint16(anon_sym_POUNDs),
	1565: uint16(anon_sym_POUNDhashalw),
	1566: uint16(anon_sym_POUNDhasheqv),
	1567: uint16(anon_sym_SQUOTE),
	1568: uint16(anon_sym_BQUOTE),
	1569: uint16(anon_sym_POUND_SQUOTE),
	1570: uint16(anon_sym_POUND_BQUOTE),
	1571: uint16(anon_sym_COMMA_AT),
	1572: uint16(anon_sym_POUND_COMMA_AT),
	1573: uint16(anon_sym_POUNDreader),
	1574: uint16(anon_sym_POUNDlang),
	1575: uint16(2),
	1576: uint16(883),
	1577: uint16(10),
	1578: uint16(sym_dot),
	1579: uint16(sym_boolean),
	1580: uint16(anon_sym_POUND),
	1581: uint16(sym_number),
	1582: uint16(sym_symbol),
	1583: uint16(anon_sym_POUNDhash),
	1584: uint16(anon_sym_POUNDhasheq),
	1585: uint16(anon_sym_COMMA),
	1586: uint16(anon_sym_POUND_COMMA),
	1587: uint16(anon_sym_POUND_BANG),
	1588: uint16(881),
	1589: uint16(30),
	1590: uint16(aux_sym__skip_token1),
	1591: uint16(aux_sym_comment_token1),
	1592: uint16(anon_sym_POUND_PIPE),
	1593: uint16(anon_sym_POUND_SEMI),
	1594: uint16(sym__line_comment),
	1595: uint16(anon_sym_POUND_LT_LT),
	1596: uint16(aux_sym_regex_token1),
	1597: uint16(anon_sym_DQUOTE),
	1598: uint16(sym_character),
	1599: uint16(sym_keyword),
	1600: uint16(anon_sym_POUND_AMP),
	1601: uint16(anon_sym_LPAREN),
	1602: uint16(anon_sym_RPAREN),
	1603: uint16(anon_sym_LBRACK),
	1604: uint16(anon_sym_RBRACK),
	1605: uint16(anon_sym_LBRACE),
	1606: uint16(anon_sym_RBRACE),
	1607: uint16(anon_sym_POUNDfl),
	1608: uint16(anon_sym_POUNDfx),
	1609: uint16(anon_sym_POUNDs),
	1610: uint16(anon_sym_POUNDhashalw),
	1611: uint16(anon_sym_POUNDhasheqv),
	1612: uint16(anon_sym_SQUOTE),
	1613: uint16(anon_sym_BQUOTE),
	1614: uint16(anon_sym_POUND_SQUOTE),
	1615: uint16(anon_sym_POUND_BQUOTE),
	1616: uint16(anon_sym_COMMA_AT),
	1617: uint16(anon_sym_POUND_COMMA_AT),
	1618: uint16(anon_sym_POUNDreader),
	1619: uint16(anon_sym_POUNDlang),
	1620: uint16(2),
	1621: uint16(887),
	1622: uint16(10),
	1623: uint16(sym_dot),
	1624: uint16(sym_boolean),
	1625: uint16(anon_sym_POUND),
	1626: uint16(sym_number),
	1627: uint16(sym_symbol),
	1628: uint16(anon_sym_POUNDhash),
	1629: uint16(anon_sym_POUNDhasheq),
	1630: uint16(anon_sym_COMMA),
	1631: uint16(anon_sym_POUND_COMMA),
	1632: uint16(anon_sym_POUND_BANG),
	1633: uint16(885),
	1634: uint16(30),
	1635: uint16(aux_sym__skip_token1),
	1636: uint16(aux_sym_comment_token1),
	1637: uint16(anon_sym_POUND_PIPE),
	1638: uint16(anon_sym_POUND_SEMI),
	1639: uint16(sym__line_comment),
	1640: uint16(anon_sym_POUND_LT_LT),
	1641: uint16(aux_sym_regex_token1),
	1642: uint16(anon_sym_DQUOTE),
	1643: uint16(sym_character),
	1644: uint16(sym_keyword),
	1645: uint16(anon_sym_POUND_AMP),
	1646: uint16(anon_sym_LPAREN),
	1647: uint16(anon_sym_RPAREN),
	1648: uint16(anon_sym_LBRACK),
	1649: uint16(anon_sym_RBRACK),
	1650: uint16(anon_sym_LBRACE),
	1651: uint16(anon_sym_RBRACE),
	1652: uint16(anon_sym_POUNDfl),
	1653: uint16(anon_sym_POUNDfx),
	1654: uint16(anon_sym_POUNDs),
	1655: uint16(anon_sym_POUNDhashalw),
	1656: uint16(anon_sym_POUNDhasheqv),
	1657: uint16(anon_sym_SQUOTE),
	1658: uint16(anon_sym_BQUOTE),
	1659: uint16(anon_sym_POUND_SQUOTE),
	1660: uint16(anon_sym_POUND_BQUOTE),
	1661: uint16(anon_sym_COMMA_AT),
	1662: uint16(anon_sym_POUND_COMMA_AT),
	1663: uint16(anon_sym_POUNDreader),
	1664: uint16(anon_sym_POUNDlang),
	1665: uint16(2),
	1666: uint16(891),
	1667: uint16(10),
	1668: uint16(sym_dot),
	1669: uint16(sym_boolean),
	1670: uint16(anon_sym_POUND),
	1671: uint16(sym_number),
	1672: uint16(sym_symbol),
	1673: uint16(anon_sym_POUNDhash),
	1674: uint16(anon_sym_POUNDhasheq),
	1675: uint16(anon_sym_COMMA),
	1676: uint16(anon_sym_POUND_COMMA),
	1677: uint16(anon_sym_POUND_BANG),
	1678: uint16(889),
	1679: uint16(30),
	1680: uint16(aux_sym__skip_token1),
	1681: uint16(aux_sym_comment_token1),
	1682: uint16(anon_sym_POUND_PIPE),
	1683: uint16(anon_sym_POUND_SEMI),
	1684: uint16(sym__line_comment),
	1685: uint16(anon_sym_POUND_LT_LT),
	1686: uint16(aux_sym_regex_token1),
	1687: uint16(anon_sym_DQUOTE),
	1688: uint16(sym_character),
	1689: uint16(sym_keyword),
	1690: uint16(anon_sym_POUND_AMP),
	1691: uint16(anon_sym_LPAREN),
	1692: uint16(anon_sym_RPAREN),
	1693: uint16(anon_sym_LBRACK),
	1694: uint16(anon_sym_RBRACK),
	1695: uint16(anon_sym_LBRACE),
	1696: uint16(anon_sym_RBRACE),
	1697: uint16(anon_sym_POUNDfl),
	1698: uint16(anon_sym_POUNDfx),
	1699: uint16(anon_sym_POUNDs),
	1700: uint16(anon_sym_POUNDhashalw),
	1701: uint16(anon_sym_POUNDhasheqv),
	1702: uint16(anon_sym_SQUOTE),
	1703: uint16(anon_sym_BQUOTE),
	1704: uint16(anon_sym_POUND_SQUOTE),
	1705: uint16(anon_sym_POUND_BQUOTE),
	1706: uint16(anon_sym_COMMA_AT),
	1707: uint16(anon_sym_POUND_COMMA_AT),
	1708: uint16(anon_sym_POUNDreader),
	1709: uint16(anon_sym_POUNDlang),
	1710: uint16(2),
	1711: uint16(895),
	1712: uint16(10),
	1713: uint16(sym_dot),
	1714: uint16(sym_boolean),
	1715: uint16(anon_sym_POUND),
	1716: uint16(sym_number),
	1717: uint16(sym_symbol),
	1718: uint16(anon_sym_POUNDhash),
	1719: uint16(anon_sym_POUNDhasheq),
	1720: uint16(anon_sym_COMMA),
	1721: uint16(anon_sym_POUND_COMMA),
	1722: uint16(anon_sym_POUND_BANG),
	1723: uint16(893),
	1724: uint16(30),
	1725: uint16(aux_sym__skip_token1),
	1726: uint16(aux_sym_comment_token1),
	1727: uint16(anon_sym_POUND_PIPE),
	1728: uint16(anon_sym_POUND_SEMI),
	1729: uint16(sym__line_comment),
	1730: uint16(anon_sym_POUND_LT_LT),
	1731: uint16(aux_sym_regex_token1),
	1732: uint16(anon_sym_DQUOTE),
	1733: uint16(sym_character),
	1734: uint16(sym_keyword),
	1735: uint16(anon_sym_POUND_AMP),
	1736: uint16(anon_sym_LPAREN),
	1737: uint16(anon_sym_RPAREN),
	1738: uint16(anon_sym_LBRACK),
	1739: uint16(anon_sym_RBRACK),
	1740: uint16(anon_sym_LBRACE),
	1741: uint16(anon_sym_RBRACE),
	1742: uint16(anon_sym_POUNDfl),
	1743: uint16(anon_sym_POUNDfx),
	1744: uint16(anon_sym_POUNDs),
	1745: uint16(anon_sym_POUNDhashalw),
	1746: uint16(anon_sym_POUNDhasheqv),
	1747: uint16(anon_sym_SQUOTE),
	1748: uint16(anon_sym_BQUOTE),
	1749: uint16(anon_sym_POUND_SQUOTE),
	1750: uint16(anon_sym_POUND_BQUOTE),
	1751: uint16(anon_sym_COMMA_AT),
	1752: uint16(anon_sym_POUND_COMMA_AT),
	1753: uint16(anon_sym_POUNDreader),
	1754: uint16(anon_sym_POUNDlang),
	1755: uint16(2),
	1756: uint16(899),
	1757: uint16(10),
	1758: uint16(sym_dot),
	1759: uint16(sym_boolean),
	1760: uint16(anon_sym_POUND),
	1761: uint16(sym_number),
	1762: uint16(sym_symbol),
	1763: uint16(anon_sym_POUNDhash),
	1764: uint16(anon_sym_POUNDhasheq),
	1765: uint16(anon_sym_COMMA),
	1766: uint16(anon_sym_POUND_COMMA),
	1767: uint16(anon_sym_POUND_BANG),
	1768: uint16(897),
	1769: uint16(30),
	1770: uint16(aux_sym__skip_token1),
	1771: uint16(aux_sym_comment_token1),
	1772: uint16(anon_sym_POUND_PIPE),
	1773: uint16(anon_sym_POUND_SEMI),
	1774: uint16(sym__line_comment),
	1775: uint16(anon_sym_POUND_LT_LT),
	1776: uint16(aux_sym_regex_token1),
	1777: uint16(anon_sym_DQUOTE),
	1778: uint16(sym_character),
	1779: uint16(sym_keyword),
	1780: uint16(anon_sym_POUND_AMP),
	1781: uint16(anon_sym_LPAREN),
	1782: uint16(anon_sym_RPAREN),
	1783: uint16(anon_sym_LBRACK),
	1784: uint16(anon_sym_RBRACK),
	1785: uint16(anon_sym_LBRACE),
	1786: uint16(anon_sym_RBRACE),
	1787: uint16(anon_sym_POUNDfl),
	1788: uint16(anon_sym_POUNDfx),
	1789: uint16(anon_sym_POUNDs),
	1790: uint16(anon_sym_POUNDhashalw),
	1791: uint16(anon_sym_POUNDhasheqv),
	1792: uint16(anon_sym_SQUOTE),
	1793: uint16(anon_sym_BQUOTE),
	1794: uint16(anon_sym_POUND_SQUOTE),
	1795: uint16(anon_sym_POUND_BQUOTE),
	1796: uint16(anon_sym_COMMA_AT),
	1797: uint16(anon_sym_POUND_COMMA_AT),
	1798: uint16(anon_sym_POUNDreader),
	1799: uint16(anon_sym_POUNDlang),
	1800: uint16(7),
	1801: uint16(901),
	1802: uint16(1),
	1803: uint16(aux_sym__skip_token1),
	1804: uint16(907),
	1805: uint16(1),
	1806: uint16(anon_sym_POUND_PIPE),
	1807: uint16(910),
	1808: uint16(1),
	1809: uint16(anon_sym_POUND_SEMI),
	1810: uint16(904),
	1811: uint16(2),
	1812: uint16(aux_sym_comment_token1),
	1813: uint16(sym__line_comment),
	1814: uint16(133),
	1815: uint16(5),
	1816: uint16(sym__skip),
	1817: uint16(sym_comment),
	1818: uint16(sym_block_comment),
	1819: uint16(sym_sexp_comment),
	1820: uint16(aux_sym_sexp_comment_repeat1),
	1821: uint16(913),
	1822: uint16(8),
	1823: uint16(sym_boolean),
	1824: uint16(anon_sym_POUND),
	1825: uint16(sym_number),
	1826: uint16(sym_symbol),
	1827: uint16(anon_sym_POUNDhash),
	1828: uint16(anon_sym_POUNDhasheq),
	1829: uint16(anon_sym_COMMA),
	1830: uint16(anon_sym_POUND_COMMA),
	1831: uint16(915),
	1832: uint16(20),
	1833: uint16(anon_sym_POUND_LT_LT),
	1834: uint16(aux_sym_regex_token1),
	1835: uint16(anon_sym_DQUOTE),
	1836: uint16(sym_character),
	1837: uint16(sym_keyword),
	1838: uint16(anon_sym_POUND_AMP),
	1839: uint16(anon_sym_LPAREN),
	1840: uint16(anon_sym_LBRACK),
	1841: uint16(anon_sym_LBRACE),
	1842: uint16(anon_sym_POUNDfl),
	1843: uint16(anon_sym_POUNDfx),
	1844: uint16(anon_sym_POUNDs),
	1845: uint16(anon_sym_POUNDhashalw),
	1846: uint16(anon_sym_POUNDhasheqv),
	1847: uint16(anon_sym_SQUOTE),
	1848: uint16(anon_sym_BQUOTE),
	1849: uint16(anon_sym_POUND_SQUOTE),
	1850: uint16(anon_sym_POUND_BQUOTE),
	1851: uint16(anon_sym_COMMA_AT),
	1852: uint16(anon_sym_POUND_COMMA_AT),
	1853: uint16(2),
	1854: uint16(775),
	1855: uint16(9),
	1856: uint16(sym_boolean),
	1857: uint16(anon_sym_POUND),
	1858: uint16(sym_number),
	1859: uint16(sym_symbol),
	1860: uint16(anon_sym_POUNDhash),
	1861: uint16(anon_sym_POUNDhasheq),
	1862: uint16(anon_sym_COMMA),
	1863: uint16(anon_sym_POUND_COMMA),
	1864: uint16(anon_sym_POUND_BANG),
	1865: uint16(773),
	1866: uint16(28),
	1868: uint16(aux_sym__skip_token1),
	1869: uint16(aux_sym_comment_token1),
	1870: uint16(anon_sym_POUND_PIPE),
	1871: uint16(anon_sym_POUND_SEMI),
	1872: uint16(sym__line_comment),
	1873: uint16(anon_sym_POUND_LT_LT),
	1874: uint16(aux_sym_regex_token1),
	1875: uint16(anon_sym_DQUOTE),
	1876: uint16(sym_character),
	1877: uint16(sym_keyword),
	1878: uint16(anon_sym_POUND_AMP),
	1879: uint16(anon_sym_LPAREN),
	1880: uint16(anon_sym_LBRACK),
	1881: uint16(anon_sym_LBRACE),
	1882: uint16(anon_sym_POUNDfl),
	1883: uint16(anon_sym_POUNDfx),
	1884: uint16(anon_sym_POUNDs),
	1885: uint16(anon_sym_POUNDhashalw),
	1886: uint16(anon_sym_POUNDhasheqv),
	1887: uint16(anon_sym_SQUOTE),
	1888: uint16(anon_sym_BQUOTE),
	1889: uint16(anon_sym_POUND_SQUOTE),
	1890: uint16(anon_sym_POUND_BQUOTE),
	1891: uint16(anon_sym_COMMA_AT),
	1892: uint16(anon_sym_POUND_COMMA_AT),
	1893: uint16(anon_sym_POUNDreader),
	1894: uint16(anon_sym_POUNDlang),
	1895: uint16(2),
	1896: uint16(883),
	1897: uint16(9),
	1898: uint16(sym_boolean),
	1899: uint16(anon_sym_POUND),
	1900: uint16(sym_number),
	1901: uint16(sym_symbol),
	1902: uint16(anon_sym_POUNDhash),
	1903: uint16(anon_sym_POUNDhasheq),
	1904: uint16(anon_sym_COMMA),
	1905: uint16(anon_sym_POUND_COMMA),
	1906: uint16(anon_sym_POUND_BANG),
	1907: uint16(881),
	1908: uint16(28),
	1910: uint16(aux_sym__skip_token1),
	1911: uint16(aux_sym_comment_token1),
	1912: uint16(anon_sym_POUND_PIPE),
	1913: uint16(anon_sym_POUND_SEMI),
	1914: uint16(sym__line_comment),
	1915: uint16(anon_sym_POUND_LT_LT),
	1916: uint16(aux_sym_regex_token1),
	1917: uint16(anon_sym_DQUOTE),
	1918: uint16(sym_character),
	1919: uint16(sym_keyword),
	1920: uint16(anon_sym_POUND_AMP),
	1921: uint16(anon_sym_LPAREN),
	1922: uint16(anon_sym_LBRACK),
	1923: uint16(anon_sym_LBRACE),
	1924: uint16(anon_sym_POUNDfl),
	1925: uint16(anon_sym_POUNDfx),
	1926: uint16(anon_sym_POUNDs),
	1927: uint16(anon_sym_POUNDhashalw),
	1928: uint16(anon_sym_POUNDhasheqv),
	1929: uint16(anon_sym_SQUOTE),
	1930: uint16(anon_sym_BQUOTE),
	1931: uint16(anon_sym_POUND_SQUOTE),
	1932: uint16(anon_sym_POUND_BQUOTE),
	1933: uint16(anon_sym_COMMA_AT),
	1934: uint16(anon_sym_POUND_COMMA_AT),
	1935: uint16(anon_sym_POUNDreader),
	1936: uint16(anon_sym_POUNDlang),
	1937: uint16(2),
	1938: uint16(887),
	1939: uint16(9),
	1940: uint16(sym_boolean),
	1941: uint16(anon_sym_POUND),
	1942: uint16(sym_number),
	1943: uint16(sym_symbol),
	1944: uint16(anon_sym_POUNDhash),
	1945: uint16(anon_sym_POUNDhasheq),
	1946: uint16(anon_sym_COMMA),
	1947: uint16(anon_sym_POUND_COMMA),
	1948: uint16(anon_sym_POUND_BANG),
	1949: uint16(885),
	1950: uint16(28),
	1952: uint16(aux_sym__skip_token1),
	1953: uint16(aux_sym_comment_token1),
	1954: uint16(anon_sym_POUND_PIPE),
	1955: uint16(anon_sym_POUND_SEMI),
	1956: uint16(sym__line_comment),
	1957: uint16(anon_sym_POUND_LT_LT),
	1958: uint16(aux_sym_regex_token1),
	1959: uint16(anon_sym_DQUOTE),
	1960: uint16(sym_character),
	1961: uint16(sym_keyword),
	1962: uint16(anon_sym_POUND_AMP),
	1963: uint16(anon_sym_LPAREN),
	1964: uint16(anon_sym_LBRACK),
	1965: uint16(anon_sym_LBRACE),
	1966: uint16(anon_sym_POUNDfl),
	1967: uint16(anon_sym_POUNDfx),
	1968: uint16(anon_sym_POUNDs),
	1969: uint16(anon_sym_POUNDhashalw),
	1970: uint16(anon_sym_POUNDhasheqv),
	1971: uint16(anon_sym_SQUOTE),
	1972: uint16(anon_sym_BQUOTE),
	1973: uint16(anon_sym_POUND_SQUOTE),
	1974: uint16(anon_sym_POUND_BQUOTE),
	1975: uint16(anon_sym_COMMA_AT),
	1976: uint16(anon_sym_POUND_COMMA_AT),
	1977: uint16(anon_sym_POUNDreader),
	1978: uint16(anon_sym_POUNDlang),
	1979: uint16(2),
	1980: uint16(891),
	1981: uint16(9),
	1982: uint16(sym_boolean),
	1983: uint16(anon_sym_POUND),
	1984: uint16(sym_number),
	1985: uint16(sym_symbol),
	1986: uint16(anon_sym_POUNDhash),
	1987: uint16(anon_sym_POUNDhasheq),
	1988: uint16(anon_sym_COMMA),
	1989: uint16(anon_sym_POUND_COMMA),
	1990: uint16(anon_sym_POUND_BANG),
	1991: uint16(889),
	1992: uint16(28),
	1994: uint16(aux_sym__skip_token1),
	1995: uint16(aux_sym_comment_token1),
	1996: uint16(anon_sym_POUND_PIPE),
	1997: uint16(anon_sym_POUND_SEMI),
	1998: uint16(sym__line_comment),
	1999: uint16(anon_sym_POUND_LT_LT),
	2000: uint16(aux_sym_regex_token1),
	2001: uint16(anon_sym_DQUOTE),
	2002: uint16(sym_character),
	2003: uint16(sym_keyword),
	2004: uint16(anon_sym_POUND_AMP),
	2005: uint16(anon_sym_LPAREN),
	2006: uint16(anon_sym_LBRACK),
	2007: uint16(anon_sym_LBRACE),
	2008: uint16(anon_sym_POUNDfl),
	2009: uint16(anon_sym_POUNDfx),
	2010: uint16(anon_sym_POUNDs),
	2011: uint16(anon_sym_POUNDhashalw),
	2012: uint16(anon_sym_POUNDhasheqv),
	2013: uint16(anon_sym_SQUOTE),
	2014: uint16(anon_sym_BQUOTE),
	2015: uint16(anon_sym_POUND_SQUOTE),
	2016: uint16(anon_sym_POUND_BQUOTE),
	2017: uint16(anon_sym_COMMA_AT),
	2018: uint16(anon_sym_POUND_COMMA_AT),
	2019: uint16(anon_sym_POUNDreader),
	2020: uint16(anon_sym_POUNDlang),
	2021: uint16(2),
	2022: uint16(895),
	2023: uint16(9),
	2024: uint16(sym_boolean),
	2025: uint16(anon_sym_POUND),
	2026: uint16(sym_number),
	2027: uint16(sym_symbol),
	2028: uint16(anon_sym_POUNDhash),
	2029: uint16(anon_sym_POUNDhasheq),
	2030: uint16(anon_sym_COMMA),
	2031: uint16(anon_sym_POUND_COMMA),
	2032: uint16(anon_sym_POUND_BANG),
	2033: uint16(893),
	2034: uint16(28),
	2036: uint16(aux_sym__skip_token1),
	2037: uint16(aux_sym_comment_token1),
	2038: uint16(anon_sym_POUND_PIPE),
	2039: uint16(anon_sym_POUND_SEMI),
	2040: uint16(sym__line_comment),
	2041: uint16(anon_sym_POUND_LT_LT),
	2042: uint16(aux_sym_regex_token1),
	2043: uint16(anon_sym_DQUOTE),
	2044: uint16(sym_character),
	2045: uint16(sym_keyword),
	2046: uint16(anon_sym_POUND_AMP),
	2047: uint16(anon_sym_LPAREN),
	2048: uint16(anon_sym_LBRACK),
	2049: uint16(anon_sym_LBRACE),
	2050: uint16(anon_sym_POUNDfl),
	2051: uint16(anon_sym_POUNDfx),
	2052: uint16(anon_sym_POUNDs),
	2053: uint16(anon_sym_POUNDhashalw),
	2054: uint16(anon_sym_POUNDhasheqv),
	2055: uint16(anon_sym_SQUOTE),
	2056: uint16(anon_sym_BQUOTE),
	2057: uint16(anon_sym_POUND_SQUOTE),
	2058: uint16(anon_sym_POUND_BQUOTE),
	2059: uint16(anon_sym_COMMA_AT),
	2060: uint16(anon_sym_POUND_COMMA_AT),
	2061: uint16(anon_sym_POUNDreader),
	2062: uint16(anon_sym_POUNDlang),
	2063: uint16(2),
	2064: uint16(815),
	2065: uint16(9),
	2066: uint16(sym_boolean),
	2067: uint16(anon_sym_POUND),
	2068: uint16(sym_number),
	2069: uint16(sym_symbol),
	2070: uint16(anon_sym_POUNDhash),
	2071: uint16(anon_sym_POUNDhasheq),
	2072: uint16(anon_sym_COMMA),
	2073: uint16(anon_sym_POUND_COMMA),
	2074: uint16(anon_sym_POUND_BANG),
	2075: uint16(813),
	2076: uint16(28),
	2078: uint16(aux_sym__skip_token1),
	2079: uint16(aux_sym_comment_token1),
	2080: uint16(anon_sym_POUND_PIPE),
	2081: uint16(anon_sym_POUND_SEMI),
	2082: uint16(sym__line_comment),
	2083: uint16(anon_sym_POUND_LT_LT),
	2084: uint16(aux_sym_regex_token1),
	2085: uint16(anon_sym_DQUOTE),
	2086: uint16(sym_character),
	2087: uint16(sym_keyword),
	2088: uint16(anon_sym_POUND_AMP),
	2089: uint16(anon_sym_LPAREN),
	2090: uint16(anon_sym_LBRACK),
	2091: uint16(anon_sym_LBRACE),
	2092: uint16(anon_sym_POUNDfl),
	2093: uint16(anon_sym_POUNDfx),
	2094: uint16(anon_sym_POUNDs),
	2095: uint16(anon_sym_POUNDhashalw),
	2096: uint16(anon_sym_POUNDhasheqv),
	2097: uint16(anon_sym_SQUOTE),
	2098: uint16(anon_sym_BQUOTE),
	2099: uint16(anon_sym_POUND_SQUOTE),
	2100: uint16(anon_sym_POUND_BQUOTE),
	2101: uint16(anon_sym_COMMA_AT),
	2102: uint16(anon_sym_POUND_COMMA_AT),
	2103: uint16(anon_sym_POUNDreader),
	2104: uint16(anon_sym_POUNDlang),
	2105: uint16(2),
	2106: uint16(819),
	2107: uint16(9),
	2108: uint16(sym_boolean),
	2109: uint16(anon_sym_POUND),
	2110: uint16(sym_number),
	2111: uint16(sym_symbol),
	2112: uint16(anon_sym_POUNDhash),
	2113: uint16(anon_sym_POUNDhasheq),
	2114: uint16(anon_sym_COMMA),
	2115: uint16(anon_sym_POUND_COMMA),
	2116: uint16(anon_sym_POUND_BANG),
	2117: uint16(817),
	2118: uint16(28),
	2120: uint16(aux_sym__skip_token1),
	2121: uint16(aux_sym_comment_token1),
	2122: uint16(anon_sym_POUND_PIPE),
	2123: uint16(anon_sym_POUND_SEMI),
	2124: uint16(sym__line_comment),
	2125: uint16(anon_sym_POUND_LT_LT),
	2126: uint16(aux_sym_regex_token1),
	2127: uint16(anon_sym_DQUOTE),
	2128: uint16(sym_character),
	2129: uint16(sym_keyword),
	2130: uint16(anon_sym_POUND_AMP),
	2131: uint16(anon_sym_LPAREN),
	2132: uint16(anon_sym_LBRACK),
	2133: uint16(anon_sym_LBRACE),
	2134: uint16(anon_sym_POUNDfl),
	2135: uint16(anon_sym_POUNDfx),
	2136: uint16(anon_sym_POUNDs),
	2137: uint16(anon_sym_POUNDhashalw),
	2138: uint16(anon_sym_POUNDhasheqv),
	2139: uint16(anon_sym_SQUOTE),
	2140: uint16(anon_sym_BQUOTE),
	2141: uint16(anon_sym_POUND_SQUOTE),
	2142: uint16(anon_sym_POUND_BQUOTE),
	2143: uint16(anon_sym_COMMA_AT),
	2144: uint16(anon_sym_POUND_COMMA_AT),
	2145: uint16(anon_sym_POUNDreader),
	2146: uint16(anon_sym_POUNDlang),
	2147: uint16(2),
	2148: uint16(747),
	2149: uint16(9),
	2150: uint16(sym_boolean),
	2151: uint16(anon_sym_POUND),
	2152: uint16(sym_number),
	2153: uint16(sym_symbol),
	2154: uint16(anon_sym_POUNDhash),
	2155: uint16(anon_sym_POUNDhasheq),
	2156: uint16(anon_sym_COMMA),
	2157: uint16(anon_sym_POUND_COMMA),
	2158: uint16(anon_sym_POUND_BANG),
	2159: uint16(745),
	2160: uint16(28),
	2162: uint16(aux_sym__skip_token1),
	2163: uint16(aux_sym_comment_token1),
	2164: uint16(anon_sym_POUND_PIPE),
	2165: uint16(anon_sym_POUND_SEMI),
	2166: uint16(sym__line_comment),
	2167: uint16(anon_sym_POUND_LT_LT),
	2168: uint16(aux_sym_regex_token1),
	2169: uint16(anon_sym_DQUOTE),
	2170: uint16(sym_character),
	2171: uint16(sym_keyword),
	2172: uint16(anon_sym_POUND_AMP),
	2173: uint16(anon_sym_LPAREN),
	2174: uint16(anon_sym_LBRACK),
	2175: uint16(anon_sym_LBRACE),
	2176: uint16(anon_sym_POUNDfl),
	2177: uint16(anon_sym_POUNDfx),
	2178: uint16(anon_sym_POUNDs),
	2179: uint16(anon_sym_POUNDhashalw),
	2180: uint16(anon_sym_POUNDhasheqv),
	2181: uint16(anon_sym_SQUOTE),
	2182: uint16(anon_sym_BQUOTE),
	2183: uint16(anon_sym_POUND_SQUOTE),
	2184: uint16(anon_sym_POUND_BQUOTE),
	2185: uint16(anon_sym_COMMA_AT),
	2186: uint16(anon_sym_POUND_COMMA_AT),
	2187: uint16(anon_sym_POUNDreader),
	2188: uint16(anon_sym_POUNDlang),
	2189: uint16(2),
	2190: uint16(751),
	2191: uint16(9),
	2192: uint16(sym_boolean),
	2193: uint16(anon_sym_POUND),
	2194: uint16(sym_number),
	2195: uint16(sym_symbol),
	2196: uint16(anon_sym_POUNDhash),
	2197: uint16(anon_sym_POUNDhasheq),
	2198: uint16(anon_sym_COMMA),
	2199: uint16(anon_sym_POUND_COMMA),
	2200: uint16(anon_sym_POUND_BANG),
	2201: uint16(749),
	2202: uint16(28),
	2204: uint16(aux_sym__skip_token1),
	2205: uint16(aux_sym_comment_token1),
	2206: uint16(anon_sym_POUND_PIPE),
	2207: uint16(anon_sym_POUND_SEMI),
	2208: uint16(sym__line_comment),
	2209: uint16(anon_sym_POUND_LT_LT),
	2210: uint16(aux_sym_regex_token1),
	2211: uint16(anon_sym_DQUOTE),
	2212: uint16(sym_character),
	2213: uint16(sym_keyword),
	2214: uint16(anon_sym_POUND_AMP),
	2215: uint16(anon_sym_LPAREN),
	2216: uint16(anon_sym_LBRACK),
	2217: uint16(anon_sym_LBRACE),
	2218: uint16(anon_sym_POUNDfl),
	2219: uint16(anon_sym_POUNDfx),
	2220: uint16(anon_sym_POUNDs),
	2221: uint16(anon_sym_POUNDhashalw),
	2222: uint16(anon_sym_POUNDhasheqv),
	2223: uint16(anon_sym_SQUOTE),
	2224: uint16(anon_sym_BQUOTE),
	2225: uint16(anon_sym_POUND_SQUOTE),
	2226: uint16(anon_sym_POUND_BQUOTE),
	2227: uint16(anon_sym_COMMA_AT),
	2228: uint16(anon_sym_POUND_COMMA_AT),
	2229: uint16(anon_sym_POUNDreader),
	2230: uint16(anon_sym_POUNDlang),
	2231: uint16(2),
	2232: uint16(827),
	2233: uint16(9),
	2234: uint16(sym_boolean),
	2235: uint16(anon_sym_POUND),
	2236: uint16(sym_number),
	2237: uint16(sym_symbol),
	2238: uint16(anon_sym_POUNDhash),
	2239: uint16(anon_sym_POUNDhasheq),
	2240: uint16(anon_sym_COMMA),
	2241: uint16(anon_sym_POUND_COMMA),
	2242: uint16(anon_sym_POUND_BANG),
	2243: uint16(825),
	2244: uint16(28),
	2246: uint16(aux_sym__skip_token1),
	2247: uint16(aux_sym_comment_token1),
	2248: uint16(anon_sym_POUND_PIPE),
	2249: uint16(anon_sym_POUND_SEMI),
	2250: uint16(sym__line_comment),
	2251: uint16(anon_sym_POUND_LT_LT),
	2252: uint16(aux_sym_regex_token1),
	2253: uint16(anon_sym_DQUOTE),
	2254: uint16(sym_character),
	2255: uint16(sym_keyword),
	2256: uint16(anon_sym_POUND_AMP),
	2257: uint16(anon_sym_LPAREN),
	2258: uint16(anon_sym_LBRACK),
	2259: uint16(anon_sym_LBRACE),
	2260: uint16(anon_sym_POUNDfl),
	2261: uint16(anon_sym_POUNDfx),
	2262: uint16(anon_sym_POUNDs),
	2263: uint16(anon_sym_POUNDhashalw),
	2264: uint16(anon_sym_POUNDhasheqv),
	2265: uint16(anon_sym_SQUOTE),
	2266: uint16(anon_sym_BQUOTE),
	2267: uint16(anon_sym_POUND_SQUOTE),
	2268: uint16(anon_sym_POUND_BQUOTE),
	2269: uint16(anon_sym_COMMA_AT),
	2270: uint16(anon_sym_POUND_COMMA_AT),
	2271: uint16(anon_sym_POUNDreader),
	2272: uint16(anon_sym_POUNDlang),
	2273: uint16(2),
	2274: uint16(759),
	2275: uint16(9),
	2276: uint16(sym_boolean),
	2277: uint16(anon_sym_POUND),
	2278: uint16(sym_number),
	2279: uint16(sym_symbol),
	2280: uint16(anon_sym_POUNDhash),
	2281: uint16(anon_sym_POUNDhasheq),
	2282: uint16(anon_sym_COMMA),
	2283: uint16(anon_sym_POUND_COMMA),
	2284: uint16(anon_sym_POUND_BANG),
	2285: uint16(757),
	2286: uint16(28),
	2288: uint16(aux_sym__skip_token1),
	2289: uint16(aux_sym_comment_token1),
	2290: uint16(anon_sym_POUND_PIPE),
	2291: uint16(anon_sym_POUND_SEMI),
	2292: uint16(sym__line_comment),
	2293: uint16(anon_sym_POUND_LT_LT),
	2294: uint16(aux_sym_regex_token1),
	2295: uint16(anon_sym_DQUOTE),
	2296: uint16(sym_character),
	2297: uint16(sym_keyword),
	2298: uint16(anon_sym_POUND_AMP),
	2299: uint16(anon_sym_LPAREN),
	2300: uint16(anon_sym_LBRACK),
	2301: uint16(anon_sym_LBRACE),
	2302: uint16(anon_sym_POUNDfl),
	2303: uint16(anon_sym_POUNDfx),
	2304: uint16(anon_sym_POUNDs),
	2305: uint16(anon_sym_POUNDhashalw),
	2306: uint16(anon_sym_POUNDhasheqv),
	2307: uint16(anon_sym_SQUOTE),
	2308: uint16(anon_sym_BQUOTE),
	2309: uint16(anon_sym_POUND_SQUOTE),
	2310: uint16(anon_sym_POUND_BQUOTE),
	2311: uint16(anon_sym_COMMA_AT),
	2312: uint16(anon_sym_POUND_COMMA_AT),
	2313: uint16(anon_sym_POUNDreader),
	2314: uint16(anon_sym_POUNDlang),
	2315: uint16(2),
	2316: uint16(835),
	2317: uint16(9),
	2318: uint16(sym_boolean),
	2319: uint16(anon_sym_POUND),
	2320: uint16(sym_number),
	2321: uint16(sym_symbol),
	2322: uint16(anon_sym_POUNDhash),
	2323: uint16(anon_sym_POUNDhasheq),
	2324: uint16(anon_sym_COMMA),
	2325: uint16(anon_sym_POUND_COMMA),
	2326: uint16(anon_sym_POUND_BANG),
	2327: uint16(833),
	2328: uint16(28),
	2330: uint16(aux_sym__skip_token1),
	2331: uint16(aux_sym_comment_token1),
	2332: uint16(anon_sym_POUND_PIPE),
	2333: uint16(anon_sym_POUND_SEMI),
	2334: uint16(sym__line_comment),
	2335: uint16(anon_sym_POUND_LT_LT),
	2336: uint16(aux_sym_regex_token1),
	2337: uint16(anon_sym_DQUOTE),
	2338: uint16(sym_character),
	2339: uint16(sym_keyword),
	2340: uint16(anon_sym_POUND_AMP),
	2341: uint16(anon_sym_LPAREN),
	2342: uint16(anon_sym_LBRACK),
	2343: uint16(anon_sym_LBRACE),
	2344: uint16(anon_sym_POUNDfl),
	2345: uint16(anon_sym_POUNDfx),
	2346: uint16(anon_sym_POUNDs),
	2347: uint16(anon_sym_POUNDhashalw),
	2348: uint16(anon_sym_POUNDhasheqv),
	2349: uint16(anon_sym_SQUOTE),
	2350: uint16(anon_sym_BQUOTE),
	2351: uint16(anon_sym_POUND_SQUOTE),
	2352: uint16(anon_sym_POUND_BQUOTE),
	2353: uint16(anon_sym_COMMA_AT),
	2354: uint16(anon_sym_POUND_COMMA_AT),
	2355: uint16(anon_sym_POUNDreader),
	2356: uint16(anon_sym_POUNDlang),
	2357: uint16(2),
	2358: uint16(847),
	2359: uint16(9),
	2360: uint16(sym_boolean),
	2361: uint16(anon_sym_POUND),
	2362: uint16(sym_number),
	2363: uint16(sym_symbol),
	2364: uint16(anon_sym_POUNDhash),
	2365: uint16(anon_sym_POUNDhasheq),
	2366: uint16(anon_sym_COMMA),
	2367: uint16(anon_sym_POUND_COMMA),
	2368: uint16(anon_sym_POUND_BANG),
	2369: uint16(845),
	2370: uint16(28),
	2372: uint16(aux_sym__skip_token1),
	2373: uint16(aux_sym_comment_token1),
	2374: uint16(anon_sym_POUND_PIPE),
	2375: uint16(anon_sym_POUND_SEMI),
	2376: uint16(sym__line_comment),
	2377: uint16(anon_sym_POUND_LT_LT),
	2378: uint16(aux_sym_regex_token1),
	2379: uint16(anon_sym_DQUOTE),
	2380: uint16(sym_character),
	2381: uint16(sym_keyword),
	2382: uint16(anon_sym_POUND_AMP),
	2383: uint16(anon_sym_LPAREN),
	2384: uint16(anon_sym_LBRACK),
	2385: uint16(anon_sym_LBRACE),
	2386: uint16(anon_sym_POUNDfl),
	2387: uint16(anon_sym_POUNDfx),
	2388: uint16(anon_sym_POUNDs),
	2389: uint16(anon_sym_POUNDhashalw),
	2390: uint16(anon_sym_POUNDhasheqv),
	2391: uint16(anon_sym_SQUOTE),
	2392: uint16(anon_sym_BQUOTE),
	2393: uint16(anon_sym_POUND_SQUOTE),
	2394: uint16(anon_sym_POUND_BQUOTE),
	2395: uint16(anon_sym_COMMA_AT),
	2396: uint16(anon_sym_POUND_COMMA_AT),
	2397: uint16(anon_sym_POUNDreader),
	2398: uint16(anon_sym_POUNDlang),
	2399: uint16(2),
	2400: uint16(755),
	2401: uint16(9),
	2402: uint16(sym_boolean),
	2403: uint16(anon_sym_POUND),
	2404: uint16(sym_number),
	2405: uint16(sym_symbol),
	2406: uint16(anon_sym_POUNDhash),
	2407: uint16(anon_sym_POUNDhasheq),
	2408: uint16(anon_sym_COMMA),
	2409: uint16(anon_sym_POUND_COMMA),
	2410: uint16(anon_sym_POUND_BANG),
	2411: uint16(753),
	2412: uint16(28),
	2414: uint16(aux_sym__skip_token1),
	2415: uint16(aux_sym_comment_token1),
	2416: uint16(anon_sym_POUND_PIPE),
	2417: uint16(anon_sym_POUND_SEMI),
	2418: uint16(sym__line_comment),
	2419: uint16(anon_sym_POUND_LT_LT),
	2420: uint16(aux_sym_regex_token1),
	2421: uint16(anon_sym_DQUOTE),
	2422: uint16(sym_character),
	2423: uint16(sym_keyword),
	2424: uint16(anon_sym_POUND_AMP),
	2425: uint16(anon_sym_LPAREN),
	2426: uint16(anon_sym_LBRACK),
	2427: uint16(anon_sym_LBRACE),
	2428: uint16(anon_sym_POUNDfl),
	2429: uint16(anon_sym_POUNDfx),
	2430: uint16(anon_sym_POUNDs),
	2431: uint16(anon_sym_POUNDhashalw),
	2432: uint16(anon_sym_POUNDhasheqv),
	2433: uint16(anon_sym_SQUOTE),
	2434: uint16(anon_sym_BQUOTE),
	2435: uint16(anon_sym_POUND_SQUOTE),
	2436: uint16(anon_sym_POUND_BQUOTE),
	2437: uint16(anon_sym_COMMA_AT),
	2438: uint16(anon_sym_POUND_COMMA_AT),
	2439: uint16(anon_sym_POUNDreader),
	2440: uint16(anon_sym_POUNDlang),
	2441: uint16(2),
	2442: uint16(767),
	2443: uint16(9),
	2444: uint16(sym_boolean),
	2445: uint16(anon_sym_POUND),
	2446: uint16(sym_number),
	2447: uint16(sym_symbol),
	2448: uint16(anon_sym_POUNDhash),
	2449: uint16(anon_sym_POUNDhasheq),
	2450: uint16(anon_sym_COMMA),
	2451: uint16(anon_sym_POUND_COMMA),
	2452: uint16(anon_sym_POUND_BANG),
	2453: uint16(765),
	2454: uint16(28),
	2456: uint16(aux_sym__skip_token1),
	2457: uint16(aux_sym_comment_token1),
	2458: uint16(anon_sym_POUND_PIPE),
	2459: uint16(anon_sym_POUND_SEMI),
	2460: uint16(sym__line_comment),
	2461: uint16(anon_sym_POUND_LT_LT),
	2462: uint16(aux_sym_regex_token1),
	2463: uint16(anon_sym_DQUOTE),
	2464: uint16(sym_character),
	2465: uint16(sym_keyword),
	2466: uint16(anon_sym_POUND_AMP),
	2467: uint16(anon_sym_LPAREN),
	2468: uint16(anon_sym_LBRACK),
	2469: uint16(anon_sym_LBRACE),
	2470: uint16(anon_sym_POUNDfl),
	2471: uint16(anon_sym_POUNDfx),
	2472: uint16(anon_sym_POUNDs),
	2473: uint16(anon_sym_POUNDhashalw),
	2474: uint16(anon_sym_POUNDhasheqv),
	2475: uint16(anon_sym_SQUOTE),
	2476: uint16(anon_sym_BQUOTE),
	2477: uint16(anon_sym_POUND_SQUOTE),
	2478: uint16(anon_sym_POUND_BQUOTE),
	2479: uint16(anon_sym_COMMA_AT),
	2480: uint16(anon_sym_POUND_COMMA_AT),
	2481: uint16(anon_sym_POUNDreader),
	2482: uint16(anon_sym_POUNDlang),
	2483: uint16(2),
	2484: uint16(771),
	2485: uint16(9),
	2486: uint16(sym_boolean),
	2487: uint16(anon_sym_POUND),
	2488: uint16(sym_number),
	2489: uint16(sym_symbol),
	2490: uint16(anon_sym_POUNDhash),
	2491: uint16(anon_sym_POUNDhasheq),
	2492: uint16(anon_sym_COMMA),
	2493: uint16(anon_sym_POUND_COMMA),
	2494: uint16(anon_sym_POUND_BANG),
	2495: uint16(769),
	2496: uint16(28),
	2498: uint16(aux_sym__skip_token1),
	2499: uint16(aux_sym_comment_token1),
	2500: uint16(anon_sym_POUND_PIPE),
	2501: uint16(anon_sym_POUND_SEMI),
	2502: uint16(sym__line_comment),
	2503: uint16(anon_sym_POUND_LT_LT),
	2504: uint16(aux_sym_regex_token1),
	2505: uint16(anon_sym_DQUOTE),
	2506: uint16(sym_character),
	2507: uint16(sym_keyword),
	2508: uint16(anon_sym_POUND_AMP),
	2509: uint16(anon_sym_LPAREN),
	2510: uint16(anon_sym_LBRACK),
	2511: uint16(anon_sym_LBRACE),
	2512: uint16(anon_sym_POUNDfl),
	2513: uint16(anon_sym_POUNDfx),
	2514: uint16(anon_sym_POUNDs),
	2515: uint16(anon_sym_POUNDhashalw),
	2516: uint16(anon_sym_POUNDhasheqv),
	2517: uint16(anon_sym_SQUOTE),
	2518: uint16(anon_sym_BQUOTE),
	2519: uint16(anon_sym_POUND_SQUOTE),
	2520: uint16(anon_sym_POUND_BQUOTE),
	2521: uint16(anon_sym_COMMA_AT),
	2522: uint16(anon_sym_POUND_COMMA_AT),
	2523: uint16(anon_sym_POUNDreader),
	2524: uint16(anon_sym_POUNDlang),
	2525: uint16(2),
	2526: uint16(879),
	2527: uint16(9),
	2528: uint16(sym_boolean),
	2529: uint16(anon_sym_POUND),
	2530: uint16(sym_number),
	2531: uint16(sym_symbol),
	2532: uint16(anon_sym_POUNDhash),
	2533: uint16(anon_sym_POUNDhasheq),
	2534: uint16(anon_sym_COMMA),
	2535: uint16(anon_sym_POUND_COMMA),
	2536: uint16(anon_sym_POUND_BANG),
	2537: uint16(877),
	2538: uint16(28),
	2540: uint16(aux_sym__skip_token1),
	2541: uint16(aux_sym_comment_token1),
	2542: uint16(anon_sym_POUND_PIPE),
	2543: uint16(anon_sym_POUND_SEMI),
	2544: uint16(sym__line_comment),
	2545: uint16(anon_sym_POUND_LT_LT),
	2546: uint16(aux_sym_regex_token1),
	2547: uint16(anon_sym_DQUOTE),
	2548: uint16(sym_character),
	2549: uint16(sym_keyword),
	2550: uint16(anon_sym_POUND_AMP),
	2551: uint16(anon_sym_LPAREN),
	2552: uint16(anon_sym_LBRACK),
	2553: uint16(anon_sym_LBRACE),
	2554: uint16(anon_sym_POUNDfl),
	2555: uint16(anon_sym_POUNDfx),
	2556: uint16(anon_sym_POUNDs),
	2557: uint16(anon_sym_POUNDhashalw),
	2558: uint16(anon_sym_POUNDhasheqv),
	2559: uint16(anon_sym_SQUOTE),
	2560: uint16(anon_sym_BQUOTE),
	2561: uint16(anon_sym_POUND_SQUOTE),
	2562: uint16(anon_sym_POUND_BQUOTE),
	2563: uint16(anon_sym_COMMA_AT),
	2564: uint16(anon_sym_POUND_COMMA_AT),
	2565: uint16(anon_sym_POUNDreader),
	2566: uint16(anon_sym_POUNDlang),
	2567: uint16(2),
	2568: uint16(779),
	2569: uint16(9),
	2570: uint16(sym_boolean),
	2571: uint16(anon_sym_POUND),
	2572: uint16(sym_number),
	2573: uint16(sym_symbol),
	2574: uint16(anon_sym_POUNDhash),
	2575: uint16(anon_sym_POUNDhasheq),
	2576: uint16(anon_sym_COMMA),
	2577: uint16(anon_sym_POUND_COMMA),
	2578: uint16(anon_sym_POUND_BANG),
	2579: uint16(777),
	2580: uint16(28),
	2582: uint16(aux_sym__skip_token1),
	2583: uint16(aux_sym_comment_token1),
	2584: uint16(anon_sym_POUND_PIPE),
	2585: uint16(anon_sym_POUND_SEMI),
	2586: uint16(sym__line_comment),
	2587: uint16(anon_sym_POUND_LT_LT),
	2588: uint16(aux_sym_regex_token1),
	2589: uint16(anon_sym_DQUOTE),
	2590: uint16(sym_character),
	2591: uint16(sym_keyword),
	2592: uint16(anon_sym_POUND_AMP),
	2593: uint16(anon_sym_LPAREN),
	2594: uint16(anon_sym_LBRACK),
	2595: uint16(anon_sym_LBRACE),
	2596: uint16(anon_sym_POUNDfl),
	2597: uint16(anon_sym_POUNDfx),
	2598: uint16(anon_sym_POUNDs),
	2599: uint16(anon_sym_POUNDhashalw),
	2600: uint16(anon_sym_POUNDhasheqv),
	2601: uint16(anon_sym_SQUOTE),
	2602: uint16(anon_sym_BQUOTE),
	2603: uint16(anon_sym_POUND_SQUOTE),
	2604: uint16(anon_sym_POUND_BQUOTE),
	2605: uint16(anon_sym_COMMA_AT),
	2606: uint16(anon_sym_POUND_COMMA_AT),
	2607: uint16(anon_sym_POUNDreader),
	2608: uint16(anon_sym_POUNDlang),
	2609: uint16(2),
	2610: uint16(787),
	2611: uint16(9),
	2612: uint16(sym_boolean),
	2613: uint16(anon_sym_POUND),
	2614: uint16(sym_number),
	2615: uint16(sym_symbol),
	2616: uint16(anon_sym_POUNDhash),
	2617: uint16(anon_sym_POUNDhasheq),
	2618: uint16(anon_sym_COMMA),
	2619: uint16(anon_sym_POUND_COMMA),
	2620: uint16(anon_sym_POUND_BANG),
	2621: uint16(785),
	2622: uint16(28),
	2624: uint16(aux_sym__skip_token1),
	2625: uint16(aux_sym_comment_token1),
	2626: uint16(anon_sym_POUND_PIPE),
	2627: uint16(anon_sym_POUND_SEMI),
	2628: uint16(sym__line_comment),
	2629: uint16(anon_sym_POUND_LT_LT),
	2630: uint16(aux_sym_regex_token1),
	2631: uint16(anon_sym_DQUOTE),
	2632: uint16(sym_character),
	2633: uint16(sym_keyword),
	2634: uint16(anon_sym_POUND_AMP),
	2635: uint16(anon_sym_LPAREN),
	2636: uint16(anon_sym_LBRACK),
	2637: uint16(anon_sym_LBRACE),
	2638: uint16(anon_sym_POUNDfl),
	2639: uint16(anon_sym_POUNDfx),
	2640: uint16(anon_sym_POUNDs),
	2641: uint16(anon_sym_POUNDhashalw),
	2642: uint16(anon_sym_POUNDhasheqv),
	2643: uint16(anon_sym_SQUOTE),
	2644: uint16(anon_sym_BQUOTE),
	2645: uint16(anon_sym_POUND_SQUOTE),
	2646: uint16(anon_sym_POUND_BQUOTE),
	2647: uint16(anon_sym_COMMA_AT),
	2648: uint16(anon_sym_POUND_COMMA_AT),
	2649: uint16(anon_sym_POUNDreader),
	2650: uint16(anon_sym_POUNDlang),
	2651: uint16(2),
	2652: uint16(799),
	2653: uint16(9),
	2654: uint16(sym_boolean),
	2655: uint16(anon_sym_POUND),
	2656: uint16(sym_number),
	2657: uint16(sym_symbol),
	2658: uint16(anon_sym_POUNDhash),
	2659: uint16(anon_sym_POUNDhasheq),
	2660: uint16(anon_sym_COMMA),
	2661: uint16(anon_sym_POUND_COMMA),
	2662: uint16(anon_sym_POUND_BANG),
	2663: uint16(797),
	2664: uint16(28),
	2666: uint16(aux_sym__skip_token1),
	2667: uint16(aux_sym_comment_token1),
	2668: uint16(anon_sym_POUND_PIPE),
	2669: uint16(anon_sym_POUND_SEMI),
	2670: uint16(sym__line_comment),
	2671: uint16(anon_sym_POUND_LT_LT),
	2672: uint16(aux_sym_regex_token1),
	2673: uint16(anon_sym_DQUOTE),
	2674: uint16(sym_character),
	2675: uint16(sym_keyword),
	2676: uint16(anon_sym_POUND_AMP),
	2677: uint16(anon_sym_LPAREN),
	2678: uint16(anon_sym_LBRACK),
	2679: uint16(anon_sym_LBRACE),
	2680: uint16(anon_sym_POUNDfl),
	2681: uint16(anon_sym_POUNDfx),
	2682: uint16(anon_sym_POUNDs),
	2683: uint16(anon_sym_POUNDhashalw),
	2684: uint16(anon_sym_POUNDhasheqv),
	2685: uint16(anon_sym_SQUOTE),
	2686: uint16(anon_sym_BQUOTE),
	2687: uint16(anon_sym_POUND_SQUOTE),
	2688: uint16(anon_sym_POUND_BQUOTE),
	2689: uint16(anon_sym_COMMA_AT),
	2690: uint16(anon_sym_POUND_COMMA_AT),
	2691: uint16(anon_sym_POUNDreader),
	2692: uint16(anon_sym_POUNDlang),
	2693: uint16(2),
	2694: uint16(795),
	2695: uint16(9),
	2696: uint16(sym_boolean),
	2697: uint16(anon_sym_POUND),
	2698: uint16(sym_number),
	2699: uint16(sym_symbol),
	2700: uint16(anon_sym_POUNDhash),
	2701: uint16(anon_sym_POUNDhasheq),
	2702: uint16(anon_sym_COMMA),
	2703: uint16(anon_sym_POUND_COMMA),
	2704: uint16(anon_sym_POUND_BANG),
	2705: uint16(793),
	2706: uint16(28),
	2708: uint16(aux_sym__skip_token1),
	2709: uint16(aux_sym_comment_token1),
	2710: uint16(anon_sym_POUND_PIPE),
	2711: uint16(anon_sym_POUND_SEMI),
	2712: uint16(sym__line_comment),
	2713: uint16(anon_sym_POUND_LT_LT),
	2714: uint16(aux_sym_regex_token1),
	2715: uint16(anon_sym_DQUOTE),
	2716: uint16(sym_character),
	2717: uint16(sym_keyword),
	2718: uint16(anon_sym_POUND_AMP),
	2719: uint16(anon_sym_LPAREN),
	2720: uint16(anon_sym_LBRACK),
	2721: uint16(anon_sym_LBRACE),
	2722: uint16(anon_sym_POUNDfl),
	2723: uint16(anon_sym_POUNDfx),
	2724: uint16(anon_sym_POUNDs),
	2725: uint16(anon_sym_POUNDhashalw),
	2726: uint16(anon_sym_POUNDhasheqv),
	2727: uint16(anon_sym_SQUOTE),
	2728: uint16(anon_sym_BQUOTE),
	2729: uint16(anon_sym_POUND_SQUOTE),
	2730: uint16(anon_sym_POUND_BQUOTE),
	2731: uint16(anon_sym_COMMA_AT),
	2732: uint16(anon_sym_POUND_COMMA_AT),
	2733: uint16(anon_sym_POUNDreader),
	2734: uint16(anon_sym_POUNDlang),
	2735: uint16(2),
	2736: uint16(807),
	2737: uint16(9),
	2738: uint16(sym_boolean),
	2739: uint16(anon_sym_POUND),
	2740: uint16(sym_number),
	2741: uint16(sym_symbol),
	2742: uint16(anon_sym_POUNDhash),
	2743: uint16(anon_sym_POUNDhasheq),
	2744: uint16(anon_sym_COMMA),
	2745: uint16(anon_sym_POUND_COMMA),
	2746: uint16(anon_sym_POUND_BANG),
	2747: uint16(805),
	2748: uint16(28),
	2750: uint16(aux_sym__skip_token1),
	2751: uint16(aux_sym_comment_token1),
	2752: uint16(anon_sym_POUND_PIPE),
	2753: uint16(anon_sym_POUND_SEMI),
	2754: uint16(sym__line_comment),
	2755: uint16(anon_sym_POUND_LT_LT),
	2756: uint16(aux_sym_regex_token1),
	2757: uint16(anon_sym_DQUOTE),
	2758: uint16(sym_character),
	2759: uint16(sym_keyword),
	2760: uint16(anon_sym_POUND_AMP),
	2761: uint16(anon_sym_LPAREN),
	2762: uint16(anon_sym_LBRACK),
	2763: uint16(anon_sym_LBRACE),
	2764: uint16(anon_sym_POUNDfl),
	2765: uint16(anon_sym_POUNDfx),
	2766: uint16(anon_sym_POUNDs),
	2767: uint16(anon_sym_POUNDhashalw),
	2768: uint16(anon_sym_POUNDhasheqv),
	2769: uint16(anon_sym_SQUOTE),
	2770: uint16(anon_sym_BQUOTE),
	2771: uint16(anon_sym_POUND_SQUOTE),
	2772: uint16(anon_sym_POUND_BQUOTE),
	2773: uint16(anon_sym_COMMA_AT),
	2774: uint16(anon_sym_POUND_COMMA_AT),
	2775: uint16(anon_sym_POUNDreader),
	2776: uint16(anon_sym_POUNDlang),
	2777: uint16(2),
	2778: uint16(743),
	2779: uint16(9),
	2780: uint16(sym_boolean),
	2781: uint16(anon_sym_POUND),
	2782: uint16(sym_number),
	2783: uint16(sym_symbol),
	2784: uint16(anon_sym_POUNDhash),
	2785: uint16(anon_sym_POUNDhasheq),
	2786: uint16(anon_sym_COMMA),
	2787: uint16(anon_sym_POUND_COMMA),
	2788: uint16(anon_sym_POUND_BANG),
	2789: uint16(741),
	2790: uint16(28),
	2792: uint16(aux_sym__skip_token1),
	2793: uint16(aux_sym_comment_token1),
	2794: uint16(anon_sym_POUND_PIPE),
	2795: uint16(anon_sym_POUND_SEMI),
	2796: uint16(sym__line_comment),
	2797: uint16(anon_sym_POUND_LT_LT),
	2798: uint16(aux_sym_regex_token1),
	2799: uint16(anon_sym_DQUOTE),
	2800: uint16(sym_character),
	2801: uint16(sym_keyword),
	2802: uint16(anon_sym_POUND_AMP),
	2803: uint16(anon_sym_LPAREN),
	2804: uint16(anon_sym_LBRACK),
	2805: uint16(anon_sym_LBRACE),
	2806: uint16(anon_sym_POUNDfl),
	2807: uint16(anon_sym_POUNDfx),
	2808: uint16(anon_sym_POUNDs),
	2809: uint16(anon_sym_POUNDhashalw),
	2810: uint16(anon_sym_POUNDhasheqv),
	2811: uint16(anon_sym_SQUOTE),
	2812: uint16(anon_sym_BQUOTE),
	2813: uint16(anon_sym_POUND_SQUOTE),
	2814: uint16(anon_sym_POUND_BQUOTE),
	2815: uint16(anon_sym_COMMA_AT),
	2816: uint16(anon_sym_POUND_COMMA_AT),
	2817: uint16(anon_sym_POUNDreader),
	2818: uint16(anon_sym_POUNDlang),
	2819: uint16(2),
	2820: uint16(823),
	2821: uint16(9),
	2822: uint16(sym_boolean),
	2823: uint16(anon_sym_POUND),
	2824: uint16(sym_number),
	2825: uint16(sym_symbol),
	2826: uint16(anon_sym_POUNDhash),
	2827: uint16(anon_sym_POUNDhasheq),
	2828: uint16(anon_sym_COMMA),
	2829: uint16(anon_sym_POUND_COMMA),
	2830: uint16(anon_sym_POUND_BANG),
	2831: uint16(821),
	2832: uint16(28),
	2834: uint16(aux_sym__skip_token1),
	2835: uint16(aux_sym_comment_token1),
	2836: uint16(anon_sym_POUND_PIPE),
	2837: uint16(anon_sym_POUND_SEMI),
	2838: uint16(sym__line_comment),
	2839: uint16(anon_sym_POUND_LT_LT),
	2840: uint16(aux_sym_regex_token1),
	2841: uint16(anon_sym_DQUOTE),
	2842: uint16(sym_character),
	2843: uint16(sym_keyword),
	2844: uint16(anon_sym_POUND_AMP),
	2845: uint16(anon_sym_LPAREN),
	2846: uint16(anon_sym_LBRACK),
	2847: uint16(anon_sym_LBRACE),
	2848: uint16(anon_sym_POUNDfl),
	2849: uint16(anon_sym_POUNDfx),
	2850: uint16(anon_sym_POUNDs),
	2851: uint16(anon_sym_POUNDhashalw),
	2852: uint16(anon_sym_POUNDhasheqv),
	2853: uint16(anon_sym_SQUOTE),
	2854: uint16(anon_sym_BQUOTE),
	2855: uint16(anon_sym_POUND_SQUOTE),
	2856: uint16(anon_sym_POUND_BQUOTE),
	2857: uint16(anon_sym_COMMA_AT),
	2858: uint16(anon_sym_POUND_COMMA_AT),
	2859: uint16(anon_sym_POUNDreader),
	2860: uint16(anon_sym_POUNDlang),
	2861: uint16(2),
	2862: uint16(783),
	2863: uint16(9),
	2864: uint16(sym_boolean),
	2865: uint16(anon_sym_POUND),
	2866: uint16(sym_number),
	2867: uint16(sym_symbol),
	2868: uint16(anon_sym_POUNDhash),
	2869: uint16(anon_sym_POUNDhasheq),
	2870: uint16(anon_sym_COMMA),
	2871: uint16(anon_sym_POUND_COMMA),
	2872: uint16(anon_sym_POUND_BANG),
	2873: uint16(781),
	2874: uint16(28),
	2876: uint16(aux_sym__skip_token1),
	2877: uint16(aux_sym_comment_token1),
	2878: uint16(anon_sym_POUND_PIPE),
	2879: uint16(anon_sym_POUND_SEMI),
	2880: uint16(sym__line_comment),
	2881: uint16(anon_sym_POUND_LT_LT),
	2882: uint16(aux_sym_regex_token1),
	2883: uint16(anon_sym_DQUOTE),
	2884: uint16(sym_character),
	2885: uint16(sym_keyword),
	2886: uint16(anon_sym_POUND_AMP),
	2887: uint16(anon_sym_LPAREN),
	2888: uint16(anon_sym_LBRACK),
	2889: uint16(anon_sym_LBRACE),
	2890: uint16(anon_sym_POUNDfl),
	2891: uint16(anon_sym_POUNDfx),
	2892: uint16(anon_sym_POUNDs),
	2893: uint16(anon_sym_POUNDhashalw),
	2894: uint16(anon_sym_POUNDhasheqv),
	2895: uint16(anon_sym_SQUOTE),
	2896: uint16(anon_sym_BQUOTE),
	2897: uint16(anon_sym_POUND_SQUOTE),
	2898: uint16(anon_sym_POUND_BQUOTE),
	2899: uint16(anon_sym_COMMA_AT),
	2900: uint16(anon_sym_POUND_COMMA_AT),
	2901: uint16(anon_sym_POUNDreader),
	2902: uint16(anon_sym_POUNDlang),
	2903: uint16(2),
	2904: uint16(831),
	2905: uint16(9),
	2906: uint16(sym_boolean),
	2907: uint16(anon_sym_POUND),
	2908: uint16(sym_number),
	2909: uint16(sym_symbol),
	2910: uint16(anon_sym_POUNDhash),
	2911: uint16(anon_sym_POUNDhasheq),
	2912: uint16(anon_sym_COMMA),
	2913: uint16(anon_sym_POUND_COMMA),
	2914: uint16(anon_sym_POUND_BANG),
	2915: uint16(829),
	2916: uint16(28),
	2918: uint16(aux_sym__skip_token1),
	2919: uint16(aux_sym_comment_token1),
	2920: uint16(anon_sym_POUND_PIPE),
	2921: uint16(anon_sym_POUND_SEMI),
	2922: uint16(sym__line_comment),
	2923: uint16(anon_sym_POUND_LT_LT),
	2924: uint16(aux_sym_regex_token1),
	2925: uint16(anon_sym_DQUOTE),
	2926: uint16(sym_character),
	2927: uint16(sym_keyword),
	2928: uint16(anon_sym_POUND_AMP),
	2929: uint16(anon_sym_LPAREN),
	2930: uint16(anon_sym_LBRACK),
	2931: uint16(anon_sym_LBRACE),
	2932: uint16(anon_sym_POUNDfl),
	2933: uint16(anon_sym_POUNDfx),
	2934: uint16(anon_sym_POUNDs),
	2935: uint16(anon_sym_POUNDhashalw),
	2936: uint16(anon_sym_POUNDhasheqv),
	2937: uint16(anon_sym_SQUOTE),
	2938: uint16(anon_sym_BQUOTE),
	2939: uint16(anon_sym_POUND_SQUOTE),
	2940: uint16(anon_sym_POUND_BQUOTE),
	2941: uint16(anon_sym_COMMA_AT),
	2942: uint16(anon_sym_POUND_COMMA_AT),
	2943: uint16(anon_sym_POUNDreader),
	2944: uint16(anon_sym_POUNDlang),
	2945: uint16(2),
	2946: uint16(791),
	2947: uint16(9),
	2948: uint16(sym_boolean),
	2949: uint16(anon_sym_POUND),
	2950: uint16(sym_number),
	2951: uint16(sym_symbol),
	2952: uint16(anon_sym_POUNDhash),
	2953: uint16(anon_sym_POUNDhasheq),
	2954: uint16(anon_sym_COMMA),
	2955: uint16(anon_sym_POUND_COMMA),
	2956: uint16(anon_sym_POUND_BANG),
	2957: uint16(789),
	2958: uint16(28),
	2960: uint16(aux_sym__skip_token1),
	2961: uint16(aux_sym_comment_token1),
	2962: uint16(anon_sym_POUND_PIPE),
	2963: uint16(anon_sym_POUND_SEMI),
	2964: uint16(sym__line_comment),
	2965: uint16(anon_sym_POUND_LT_LT),
	2966: uint16(aux_sym_regex_token1),
	2967: uint16(anon_sym_DQUOTE),
	2968: uint16(sym_character),
	2969: uint16(sym_keyword),
	2970: uint16(anon_sym_POUND_AMP),
	2971: uint16(anon_sym_LPAREN),
	2972: uint16(anon_sym_LBRACK),
	2973: uint16(anon_sym_LBRACE),
	2974: uint16(anon_sym_POUNDfl),
	2975: uint16(anon_sym_POUNDfx),
	2976: uint16(anon_sym_POUNDs),
	2977: uint16(anon_sym_POUNDhashalw),
	2978: uint16(anon_sym_POUNDhasheqv),
	2979: uint16(anon_sym_SQUOTE),
	2980: uint16(anon_sym_BQUOTE),
	2981: uint16(anon_sym_POUND_SQUOTE),
	2982: uint16(anon_sym_POUND_BQUOTE),
	2983: uint16(anon_sym_COMMA_AT),
	2984: uint16(anon_sym_POUND_COMMA_AT),
	2985: uint16(anon_sym_POUNDreader),
	2986: uint16(anon_sym_POUNDlang),
	2987: uint16(2),
	2988: uint16(803),
	2989: uint16(9),
	2990: uint16(sym_boolean),
	2991: uint16(anon_sym_POUND),
	2992: uint16(sym_number),
	2993: uint16(sym_symbol),
	2994: uint16(anon_sym_POUNDhash),
	2995: uint16(anon_sym_POUNDhasheq),
	2996: uint16(anon_sym_COMMA),
	2997: uint16(anon_sym_POUND_COMMA),
	2998: uint16(anon_sym_POUND_BANG),
	2999: uint16(801),
	3000: uint16(28),
	3002: uint16(aux_sym__skip_token1),
	3003: uint16(aux_sym_comment_token1),
	3004: uint16(anon_sym_POUND_PIPE),
	3005: uint16(anon_sym_POUND_SEMI),
	3006: uint16(sym__line_comment),
	3007: uint16(anon_sym_POUND_LT_LT),
	3008: uint16(aux_sym_regex_token1),
	3009: uint16(anon_sym_DQUOTE),
	3010: uint16(sym_character),
	3011: uint16(sym_keyword),
	3012: uint16(anon_sym_POUND_AMP),
	3013: uint16(anon_sym_LPAREN),
	3014: uint16(anon_sym_LBRACK),
	3015: uint16(anon_sym_LBRACE),
	3016: uint16(anon_sym_POUNDfl),
	3017: uint16(anon_sym_POUNDfx),
	3018: uint16(anon_sym_POUNDs),
	3019: uint16(anon_sym_POUNDhashalw),
	3020: uint16(anon_sym_POUNDhasheqv),
	3021: uint16(anon_sym_SQUOTE),
	3022: uint16(anon_sym_BQUOTE),
	3023: uint16(anon_sym_POUND_SQUOTE),
	3024: uint16(anon_sym_POUND_BQUOTE),
	3025: uint16(anon_sym_COMMA_AT),
	3026: uint16(anon_sym_POUND_COMMA_AT),
	3027: uint16(anon_sym_POUNDreader),
	3028: uint16(anon_sym_POUNDlang),
	3029: uint16(2),
	3030: uint16(859),
	3031: uint16(9),
	3032: uint16(sym_boolean),
	3033: uint16(anon_sym_POUND),
	3034: uint16(sym_number),
	3035: uint16(sym_symbol),
	3036: uint16(anon_sym_POUNDhash),
	3037: uint16(anon_sym_POUNDhasheq),
	3038: uint16(anon_sym_COMMA),
	3039: uint16(anon_sym_POUND_COMMA),
	3040: uint16(anon_sym_POUND_BANG),
	3041: uint16(857),
	3042: uint16(28),
	3044: uint16(aux_sym__skip_token1),
	3045: uint16(aux_sym_comment_token1),
	3046: uint16(anon_sym_POUND_PIPE),
	3047: uint16(anon_sym_POUND_SEMI),
	3048: uint16(sym__line_comment),
	3049: uint16(anon_sym_POUND_LT_LT),
	3050: uint16(aux_sym_regex_token1),
	3051: uint16(anon_sym_DQUOTE),
	3052: uint16(sym_character),
	3053: uint16(sym_keyword),
	3054: uint16(anon_sym_POUND_AMP),
	3055: uint16(anon_sym_LPAREN),
	3056: uint16(anon_sym_LBRACK),
	3057: uint16(anon_sym_LBRACE),
	3058: uint16(anon_sym_POUNDfl),
	3059: uint16(anon_sym_POUNDfx),
	3060: uint16(anon_sym_POUNDs),
	3061: uint16(anon_sym_POUNDhashalw),
	3062: uint16(anon_sym_POUNDhasheqv),
	3063: uint16(anon_sym_SQUOTE),
	3064: uint16(anon_sym_BQUOTE),
	3065: uint16(anon_sym_POUND_SQUOTE),
	3066: uint16(anon_sym_POUND_BQUOTE),
	3067: uint16(anon_sym_COMMA_AT),
	3068: uint16(anon_sym_POUND_COMMA_AT),
	3069: uint16(anon_sym_POUNDreader),
	3070: uint16(anon_sym_POUNDlang),
	3071: uint16(2),
	3072: uint16(843),
	3073: uint16(9),
	3074: uint16(sym_boolean),
	3075: uint16(anon_sym_POUND),
	3076: uint16(sym_number),
	3077: uint16(sym_symbol),
	3078: uint16(anon_sym_POUNDhash),
	3079: uint16(anon_sym_POUNDhasheq),
	3080: uint16(anon_sym_COMMA),
	3081: uint16(anon_sym_POUND_COMMA),
	3082: uint16(anon_sym_POUND_BANG),
	3083: uint16(841),
	3084: uint16(28),
	3086: uint16(aux_sym__skip_token1),
	3087: uint16(aux_sym_comment_token1),
	3088: uint16(anon_sym_POUND_PIPE),
	3089: uint16(anon_sym_POUND_SEMI),
	3090: uint16(sym__line_comment),
	3091: uint16(anon_sym_POUND_LT_LT),
	3092: uint16(aux_sym_regex_token1),
	3093: uint16(anon_sym_DQUOTE),
	3094: uint16(sym_character),
	3095: uint16(sym_keyword),
	3096: uint16(anon_sym_POUND_AMP),
	3097: uint16(anon_sym_LPAREN),
	3098: uint16(anon_sym_LBRACK),
	3099: uint16(anon_sym_LBRACE),
	3100: uint16(anon_sym_POUNDfl),
	3101: uint16(anon_sym_POUNDfx),
	3102: uint16(anon_sym_POUNDs),
	3103: uint16(anon_sym_POUNDhashalw),
	3104: uint16(anon_sym_POUNDhasheqv),
	3105: uint16(anon_sym_SQUOTE),
	3106: uint16(anon_sym_BQUOTE),
	3107: uint16(anon_sym_POUND_SQUOTE),
	3108: uint16(anon_sym_POUND_BQUOTE),
	3109: uint16(anon_sym_COMMA_AT),
	3110: uint16(anon_sym_POUND_COMMA_AT),
	3111: uint16(anon_sym_POUNDreader),
	3112: uint16(anon_sym_POUNDlang),
	3113: uint16(2),
	3114: uint16(811),
	3115: uint16(9),
	3116: uint16(sym_boolean),
	3117: uint16(anon_sym_POUND),
	3118: uint16(sym_number),
	3119: uint16(sym_symbol),
	3120: uint16(anon_sym_POUNDhash),
	3121: uint16(anon_sym_POUNDhasheq),
	3122: uint16(anon_sym_COMMA),
	3123: uint16(anon_sym_POUND_COMMA),
	3124: uint16(anon_sym_POUND_BANG),
	3125: uint16(809),
	3126: uint16(28),
	3128: uint16(aux_sym__skip_token1),
	3129: uint16(aux_sym_comment_token1),
	3130: uint16(anon_sym_POUND_PIPE),
	3131: uint16(anon_sym_POUND_SEMI),
	3132: uint16(sym__line_comment),
	3133: uint16(anon_sym_POUND_LT_LT),
	3134: uint16(aux_sym_regex_token1),
	3135: uint16(anon_sym_DQUOTE),
	3136: uint16(sym_character),
	3137: uint16(sym_keyword),
	3138: uint16(anon_sym_POUND_AMP),
	3139: uint16(anon_sym_LPAREN),
	3140: uint16(anon_sym_LBRACK),
	3141: uint16(anon_sym_LBRACE),
	3142: uint16(anon_sym_POUNDfl),
	3143: uint16(anon_sym_POUNDfx),
	3144: uint16(anon_sym_POUNDs),
	3145: uint16(anon_sym_POUNDhashalw),
	3146: uint16(anon_sym_POUNDhasheqv),
	3147: uint16(anon_sym_SQUOTE),
	3148: uint16(anon_sym_BQUOTE),
	3149: uint16(anon_sym_POUND_SQUOTE),
	3150: uint16(anon_sym_POUND_BQUOTE),
	3151: uint16(anon_sym_COMMA_AT),
	3152: uint16(anon_sym_POUND_COMMA_AT),
	3153: uint16(anon_sym_POUNDreader),
	3154: uint16(anon_sym_POUNDlang),
	3155: uint16(2),
	3156: uint16(839),
	3157: uint16(9),
	3158: uint16(sym_boolean),
	3159: uint16(anon_sym_POUND),
	3160: uint16(sym_number),
	3161: uint16(sym_symbol),
	3162: uint16(anon_sym_POUNDhash),
	3163: uint16(anon_sym_POUNDhasheq),
	3164: uint16(anon_sym_COMMA),
	3165: uint16(anon_sym_POUND_COMMA),
	3166: uint16(anon_sym_POUND_BANG),
	3167: uint16(837),
	3168: uint16(28),
	3170: uint16(aux_sym__skip_token1),
	3171: uint16(aux_sym_comment_token1),
	3172: uint16(anon_sym_POUND_PIPE),
	3173: uint16(anon_sym_POUND_SEMI),
	3174: uint16(sym__line_comment),
	3175: uint16(anon_sym_POUND_LT_LT),
	3176: uint16(aux_sym_regex_token1),
	3177: uint16(anon_sym_DQUOTE),
	3178: uint16(sym_character),
	3179: uint16(sym_keyword),
	3180: uint16(anon_sym_POUND_AMP),
	3181: uint16(anon_sym_LPAREN),
	3182: uint16(anon_sym_LBRACK),
	3183: uint16(anon_sym_LBRACE),
	3184: uint16(anon_sym_POUNDfl),
	3185: uint16(anon_sym_POUNDfx),
	3186: uint16(anon_sym_POUNDs),
	3187: uint16(anon_sym_POUNDhashalw),
	3188: uint16(anon_sym_POUNDhasheqv),
	3189: uint16(anon_sym_SQUOTE),
	3190: uint16(anon_sym_BQUOTE),
	3191: uint16(anon_sym_POUND_SQUOTE),
	3192: uint16(anon_sym_POUND_BQUOTE),
	3193: uint16(anon_sym_COMMA_AT),
	3194: uint16(anon_sym_POUND_COMMA_AT),
	3195: uint16(anon_sym_POUNDreader),
	3196: uint16(anon_sym_POUNDlang),
	3197: uint16(2),
	3198: uint16(763),
	3199: uint16(9),
	3200: uint16(sym_boolean),
	3201: uint16(anon_sym_POUND),
	3202: uint16(sym_number),
	3203: uint16(sym_symbol),
	3204: uint16(anon_sym_POUNDhash),
	3205: uint16(anon_sym_POUNDhasheq),
	3206: uint16(anon_sym_COMMA),
	3207: uint16(anon_sym_POUND_COMMA),
	3208: uint16(anon_sym_POUND_BANG),
	3209: uint16(761),
	3210: uint16(28),
	3212: uint16(aux_sym__skip_token1),
	3213: uint16(aux_sym_comment_token1),
	3214: uint16(anon_sym_POUND_PIPE),
	3215: uint16(anon_sym_POUND_SEMI),
	3216: uint16(sym__line_comment),
	3217: uint16(anon_sym_POUND_LT_LT),
	3218: uint16(aux_sym_regex_token1),
	3219: uint16(anon_sym_DQUOTE),
	3220: uint16(sym_character),
	3221: uint16(sym_keyword),
	3222: uint16(anon_sym_POUND_AMP),
	3223: uint16(anon_sym_LPAREN),
	3224: uint16(anon_sym_LBRACK),
	3225: uint16(anon_sym_LBRACE),
	3226: uint16(anon_sym_POUNDfl),
	3227: uint16(anon_sym_POUNDfx),
	3228: uint16(anon_sym_POUNDs),
	3229: uint16(anon_sym_POUNDhashalw),
	3230: uint16(anon_sym_POUNDhasheqv),
	3231: uint16(anon_sym_SQUOTE),
	3232: uint16(anon_sym_BQUOTE),
	3233: uint16(anon_sym_POUND_SQUOTE),
	3234: uint16(anon_sym_POUND_BQUOTE),
	3235: uint16(anon_sym_COMMA_AT),
	3236: uint16(anon_sym_POUND_COMMA_AT),
	3237: uint16(anon_sym_POUNDreader),
	3238: uint16(anon_sym_POUNDlang),
	3239: uint16(2),
	3240: uint16(851),
	3241: uint16(9),
	3242: uint16(sym_boolean),
	3243: uint16(anon_sym_POUND),
	3244: uint16(sym_number),
	3245: uint16(sym_symbol),
	3246: uint16(anon_sym_POUNDhash),
	3247: uint16(anon_sym_POUNDhasheq),
	3248: uint16(anon_sym_COMMA),
	3249: uint16(anon_sym_POUND_COMMA),
	3250: uint16(anon_sym_POUND_BANG),
	3251: uint16(849),
	3252: uint16(28),
	3254: uint16(aux_sym__skip_token1),
	3255: uint16(aux_sym_comment_token1),
	3256: uint16(anon_sym_POUND_PIPE),
	3257: uint16(anon_sym_POUND_SEMI),
	3258: uint16(sym__line_comment),
	3259: uint16(anon_sym_POUND_LT_LT),
	3260: uint16(aux_sym_regex_token1),
	3261: uint16(anon_sym_DQUOTE),
	3262: uint16(sym_character),
	3263: uint16(sym_keyword),
	3264: uint16(anon_sym_POUND_AMP),
	3265: uint16(anon_sym_LPAREN),
	3266: uint16(anon_sym_LBRACK),
	3267: uint16(anon_sym_LBRACE),
	3268: uint16(anon_sym_POUNDfl),
	3269: uint16(anon_sym_POUNDfx),
	3270: uint16(anon_sym_POUNDs),
	3271: uint16(anon_sym_POUNDhashalw),
	3272: uint16(anon_sym_POUNDhasheqv),
	3273: uint16(anon_sym_SQUOTE),
	3274: uint16(anon_sym_BQUOTE),
	3275: uint16(anon_sym_POUND_SQUOTE),
	3276: uint16(anon_sym_POUND_BQUOTE),
	3277: uint16(anon_sym_COMMA_AT),
	3278: uint16(anon_sym_POUND_COMMA_AT),
	3279: uint16(anon_sym_POUNDreader),
	3280: uint16(anon_sym_POUNDlang),
	3281: uint16(2),
	3282: uint16(863),
	3283: uint16(9),
	3284: uint16(sym_boolean),
	3285: uint16(anon_sym_POUND),
	3286: uint16(sym_number),
	3287: uint16(sym_symbol),
	3288: uint16(anon_sym_POUNDhash),
	3289: uint16(anon_sym_POUNDhasheq),
	3290: uint16(anon_sym_COMMA),
	3291: uint16(anon_sym_POUND_COMMA),
	3292: uint16(anon_sym_POUND_BANG),
	3293: uint16(861),
	3294: uint16(28),
	3296: uint16(aux_sym__skip_token1),
	3297: uint16(aux_sym_comment_token1),
	3298: uint16(anon_sym_POUND_PIPE),
	3299: uint16(anon_sym_POUND_SEMI),
	3300: uint16(sym__line_comment),
	3301: uint16(anon_sym_POUND_LT_LT),
	3302: uint16(aux_sym_regex_token1),
	3303: uint16(anon_sym_DQUOTE),
	3304: uint16(sym_character),
	3305: uint16(sym_keyword),
	3306: uint16(anon_sym_POUND_AMP),
	3307: uint16(anon_sym_LPAREN),
	3308: uint16(anon_sym_LBRACK),
	3309: uint16(anon_sym_LBRACE),
	3310: uint16(anon_sym_POUNDfl),
	3311: uint16(anon_sym_POUNDfx),
	3312: uint16(anon_sym_POUNDs),
	3313: uint16(anon_sym_POUNDhashalw),
	3314: uint16(anon_sym_POUNDhasheqv),
	3315: uint16(anon_sym_SQUOTE),
	3316: uint16(anon_sym_BQUOTE),
	3317: uint16(anon_sym_POUND_SQUOTE),
	3318: uint16(anon_sym_POUND_BQUOTE),
	3319: uint16(anon_sym_COMMA_AT),
	3320: uint16(anon_sym_POUND_COMMA_AT),
	3321: uint16(anon_sym_POUNDreader),
	3322: uint16(anon_sym_POUNDlang),
	3323: uint16(2),
	3324: uint16(867),
	3325: uint16(9),
	3326: uint16(sym_boolean),
	3327: uint16(anon_sym_POUND),
	3328: uint16(sym_number),
	3329: uint16(sym_symbol),
	3330: uint16(anon_sym_POUNDhash),
	3331: uint16(anon_sym_POUNDhasheq),
	3332: uint16(anon_sym_COMMA),
	3333: uint16(anon_sym_POUND_COMMA),
	3334: uint16(anon_sym_POUND_BANG),
	3335: uint16(865),
	3336: uint16(28),
	3338: uint16(aux_sym__skip_token1),
	3339: uint16(aux_sym_comment_token1),
	3340: uint16(anon_sym_POUND_PIPE),
	3341: uint16(anon_sym_POUND_SEMI),
	3342: uint16(sym__line_comment),
	3343: uint16(anon_sym_POUND_LT_LT),
	3344: uint16(aux_sym_regex_token1),
	3345: uint16(anon_sym_DQUOTE),
	3346: uint16(sym_character),
	3347: uint16(sym_keyword),
	3348: uint16(anon_sym_POUND_AMP),
	3349: uint16(anon_sym_LPAREN),
	3350: uint16(anon_sym_LBRACK),
	3351: uint16(anon_sym_LBRACE),
	3352: uint16(anon_sym_POUNDfl),
	3353: uint16(anon_sym_POUNDfx),
	3354: uint16(anon_sym_POUNDs),
	3355: uint16(anon_sym_POUNDhashalw),
	3356: uint16(anon_sym_POUNDhasheqv),
	3357: uint16(anon_sym_SQUOTE),
	3358: uint16(anon_sym_BQUOTE),
	3359: uint16(anon_sym_POUND_SQUOTE),
	3360: uint16(anon_sym_POUND_BQUOTE),
	3361: uint16(anon_sym_COMMA_AT),
	3362: uint16(anon_sym_POUND_COMMA_AT),
	3363: uint16(anon_sym_POUNDreader),
	3364: uint16(anon_sym_POUNDlang),
	3365: uint16(2),
	3366: uint16(871),
	3367: uint16(9),
	3368: uint16(sym_boolean),
	3369: uint16(anon_sym_POUND),
	3370: uint16(sym_number),
	3371: uint16(sym_symbol),
	3372: uint16(anon_sym_POUNDhash),
	3373: uint16(anon_sym_POUNDhasheq),
	3374: uint16(anon_sym_COMMA),
	3375: uint16(anon_sym_POUND_COMMA),
	3376: uint16(anon_sym_POUND_BANG),
	3377: uint16(869),
	3378: uint16(28),
	3380: uint16(aux_sym__skip_token1),
	3381: uint16(aux_sym_comment_token1),
	3382: uint16(anon_sym_POUND_PIPE),
	3383: uint16(anon_sym_POUND_SEMI),
	3384: uint16(sym__line_comment),
	3385: uint16(anon_sym_POUND_LT_LT),
	3386: uint16(aux_sym_regex_token1),
	3387: uint16(anon_sym_DQUOTE),
	3388: uint16(sym_character),
	3389: uint16(sym_keyword),
	3390: uint16(anon_sym_POUND_AMP),
	3391: uint16(anon_sym_LPAREN),
	3392: uint16(anon_sym_LBRACK),
	3393: uint16(anon_sym_LBRACE),
	3394: uint16(anon_sym_POUNDfl),
	3395: uint16(anon_sym_POUNDfx),
	3396: uint16(anon_sym_POUNDs),
	3397: uint16(anon_sym_POUNDhashalw),
	3398: uint16(anon_sym_POUNDhasheqv),
	3399: uint16(anon_sym_SQUOTE),
	3400: uint16(anon_sym_BQUOTE),
	3401: uint16(anon_sym_POUND_SQUOTE),
	3402: uint16(anon_sym_POUND_BQUOTE),
	3403: uint16(anon_sym_COMMA_AT),
	3404: uint16(anon_sym_POUND_COMMA_AT),
	3405: uint16(anon_sym_POUNDreader),
	3406: uint16(anon_sym_POUNDlang),
	3407: uint16(2),
	3408: uint16(899),
	3409: uint16(9),
	3410: uint16(sym_boolean),
	3411: uint16(anon_sym_POUND),
	3412: uint16(sym_number),
	3413: uint16(sym_symbol),
	3414: uint16(anon_sym_POUNDhash),
	3415: uint16(anon_sym_POUNDhasheq),
	3416: uint16(anon_sym_COMMA),
	3417: uint16(anon_sym_POUND_COMMA),
	3418: uint16(anon_sym_POUND_BANG),
	3419: uint16(897),
	3420: uint16(28),
	3422: uint16(aux_sym__skip_token1),
	3423: uint16(aux_sym_comment_token1),
	3424: uint16(anon_sym_POUND_PIPE),
	3425: uint16(anon_sym_POUND_SEMI),
	3426: uint16(sym__line_comment),
	3427: uint16(anon_sym_POUND_LT_LT),
	3428: uint16(aux_sym_regex_token1),
	3429: uint16(anon_sym_DQUOTE),
	3430: uint16(sym_character),
	3431: uint16(sym_keyword),
	3432: uint16(anon_sym_POUND_AMP),
	3433: uint16(anon_sym_LPAREN),
	3434: uint16(anon_sym_LBRACK),
	3435: uint16(anon_sym_LBRACE),
	3436: uint16(anon_sym_POUNDfl),
	3437: uint16(anon_sym_POUNDfx),
	3438: uint16(anon_sym_POUNDs),
	3439: uint16(anon_sym_POUNDhashalw),
	3440: uint16(anon_sym_POUNDhasheqv),
	3441: uint16(anon_sym_SQUOTE),
	3442: uint16(anon_sym_BQUOTE),
	3443: uint16(anon_sym_POUND_SQUOTE),
	3444: uint16(anon_sym_POUND_BQUOTE),
	3445: uint16(anon_sym_COMMA_AT),
	3446: uint16(anon_sym_POUND_COMMA_AT),
	3447: uint16(anon_sym_POUNDreader),
	3448: uint16(anon_sym_POUNDlang),
	3449: uint16(2),
	3450: uint16(875),
	3451: uint16(9),
	3452: uint16(sym_boolean),
	3453: uint16(anon_sym_POUND),
	3454: uint16(sym_number),
	3455: uint16(sym_symbol),
	3456: uint16(anon_sym_POUNDhash),
	3457: uint16(anon_sym_POUNDhasheq),
	3458: uint16(anon_sym_COMMA),
	3459: uint16(anon_sym_POUND_COMMA),
	3460: uint16(anon_sym_POUND_BANG),
	3461: uint16(873),
	3462: uint16(28),
	3464: uint16(aux_sym__skip_token1),
	3465: uint16(aux_sym_comment_token1),
	3466: uint16(anon_sym_POUND_PIPE),
	3467: uint16(anon_sym_POUND_SEMI),
	3468: uint16(sym__line_comment),
	3469: uint16(anon_sym_POUND_LT_LT),
	3470: uint16(aux_sym_regex_token1),
	3471: uint16(anon_sym_DQUOTE),
	3472: uint16(sym_character),
	3473: uint16(sym_keyword),
	3474: uint16(anon_sym_POUND_AMP),
	3475: uint16(anon_sym_LPAREN),
	3476: uint16(anon_sym_LBRACK),
	3477: uint16(anon_sym_LBRACE),
	3478: uint16(anon_sym_POUNDfl),
	3479: uint16(anon_sym_POUNDfx),
	3480: uint16(anon_sym_POUNDs),
	3481: uint16(anon_sym_POUNDhashalw),
	3482: uint16(anon_sym_POUNDhasheqv),
	3483: uint16(anon_sym_SQUOTE),
	3484: uint16(anon_sym_BQUOTE),
	3485: uint16(anon_sym_POUND_SQUOTE),
	3486: uint16(anon_sym_POUND_BQUOTE),
	3487: uint16(anon_sym_COMMA_AT),
	3488: uint16(anon_sym_POUND_COMMA_AT),
	3489: uint16(anon_sym_POUNDreader),
	3490: uint16(anon_sym_POUNDlang),
	3491: uint16(2),
	3492: uint16(855),
	3493: uint16(9),
	3494: uint16(sym_boolean),
	3495: uint16(anon_sym_POUND),
	3496: uint16(sym_number),
	3497: uint16(sym_symbol),
	3498: uint16(anon_sym_POUNDhash),
	3499: uint16(anon_sym_POUNDhasheq),
	3500: uint16(anon_sym_COMMA),
	3501: uint16(anon_sym_POUND_COMMA),
	3502: uint16(anon_sym_POUND_BANG),
	3503: uint16(853),
	3504: uint16(28),
	3506: uint16(aux_sym__skip_token1),
	3507: uint16(aux_sym_comment_token1),
	3508: uint16(anon_sym_POUND_PIPE),
	3509: uint16(anon_sym_POUND_SEMI),
	3510: uint16(sym__line_comment),
	3511: uint16(anon_sym_POUND_LT_LT),
	3512: uint16(aux_sym_regex_token1),
	3513: uint16(anon_sym_DQUOTE),
	3514: uint16(sym_character),
	3515: uint16(sym_keyword),
	3516: uint16(anon_sym_POUND_AMP),
	3517: uint16(anon_sym_LPAREN),
	3518: uint16(anon_sym_LBRACK),
	3519: uint16(anon_sym_LBRACE),
	3520: uint16(anon_sym_POUNDfl),
	3521: uint16(anon_sym_POUNDfx),
	3522: uint16(anon_sym_POUNDs),
	3523: uint16(anon_sym_POUNDhashalw),
	3524: uint16(anon_sym_POUNDhasheqv),
	3525: uint16(anon_sym_SQUOTE),
	3526: uint16(anon_sym_BQUOTE),
	3527: uint16(anon_sym_POUND_SQUOTE),
	3528: uint16(anon_sym_POUND_BQUOTE),
	3529: uint16(anon_sym_COMMA_AT),
	3530: uint16(anon_sym_POUND_COMMA_AT),
	3531: uint16(anon_sym_POUNDreader),
	3532: uint16(anon_sym_POUNDlang),
	3533: uint16(2),
	3534: uint16(851),
	3535: uint16(8),
	3536: uint16(sym_boolean),
	3537: uint16(anon_sym_POUND),
	3538: uint16(sym_number),
	3539: uint16(sym_symbol),
	3540: uint16(anon_sym_POUNDhash),
	3541: uint16(anon_sym_POUNDhasheq),
	3542: uint16(anon_sym_COMMA),
	3543: uint16(anon_sym_POUND_COMMA),
	3544: uint16(849),
	3545: uint16(25),
	3546: uint16(aux_sym__skip_token1),
	3547: uint16(aux_sym_comment_token1),
	3548: uint16(anon_sym_POUND_PIPE),
	3549: uint16(anon_sym_POUND_SEMI),
	3550: uint16(sym__line_comment),
	3551: uint16(anon_sym_POUND_LT_LT),
	3552: uint16(aux_sym_regex_token1),
	3553: uint16(anon_sym_DQUOTE),
	3554: uint16(sym_character),
	3555: uint16(sym_keyword),
	3556: uint16(anon_sym_POUND_AMP),
	3557: uint16(anon_sym_LPAREN),
	3558: uint16(anon_sym_LBRACK),
	3559: uint16(anon_sym_LBRACE),
	3560: uint16(anon_sym_POUNDfl),
	3561: uint16(anon_sym_POUNDfx),
	3562: uint16(anon_sym_POUNDs),
	3563: uint16(anon_sym_POUNDhashalw),
	3564: uint16(anon_sym_POUNDhasheqv),
	3565: uint16(anon_sym_SQUOTE),
	3566: uint16(anon_sym_BQUOTE),
	3567: uint16(anon_sym_POUND_SQUOTE),
	3568: uint16(anon_sym_POUND_BQUOTE),
	3569: uint16(anon_sym_COMMA_AT),
	3570: uint16(anon_sym_POUND_COMMA_AT),
	3571: uint16(2),
	3572: uint16(847),
	3573: uint16(8),
	3574: uint16(sym_boolean),
	3575: uint16(anon_sym_POUND),
	3576: uint16(sym_number),
	3577: uint16(sym_symbol),
	3578: uint16(anon_sym_POUNDhash),
	3579: uint16(anon_sym_POUNDhasheq),
	3580: uint16(anon_sym_COMMA),
	3581: uint16(anon_sym_POUND_COMMA),
	3582: uint16(845),
	3583: uint16(25),
	3584: uint16(aux_sym__skip_token1),
	3585: uint16(aux_sym_comment_token1),
	3586: uint16(anon_sym_POUND_PIPE),
	3587: uint16(anon_sym_POUND_SEMI),
	3588: uint16(sym__line_comment),
	3589: uint16(anon_sym_POUND_LT_LT),
	3590: uint16(aux_sym_regex_token1),
	3591: uint16(anon_sym_DQUOTE),
	3592: uint16(sym_character),
	3593: uint16(sym_keyword),
	3594: uint16(anon_sym_POUND_AMP),
	3595: uint16(anon_sym_LPAREN),
	3596: uint16(anon_sym_LBRACK),
	3597: uint16(anon_sym_LBRACE),
	3598: uint16(anon_sym_POUNDfl),
	3599: uint16(anon_sym_POUNDfx),
	3600: uint16(anon_sym_POUNDs),
	3601: uint16(anon_sym_POUNDhashalw),
	3602: uint16(anon_sym_POUNDhasheqv),
	3603: uint16(anon_sym_SQUOTE),
	3604: uint16(anon_sym_BQUOTE),
	3605: uint16(anon_sym_POUND_SQUOTE),
	3606: uint16(anon_sym_POUND_BQUOTE),
	3607: uint16(anon_sym_COMMA_AT),
	3608: uint16(anon_sym_POUND_COMMA_AT),
	3609: uint16(2),
	3610: uint16(755),
	3611: uint16(8),
	3612: uint16(sym_boolean),
	3613: uint16(anon_sym_POUND),
	3614: uint16(sym_number),
	3615: uint16(sym_symbol),
	3616: uint16(anon_sym_POUNDhash),
	3617: uint16(anon_sym_POUNDhasheq),
	3618: uint16(anon_sym_COMMA),
	3619: uint16(anon_sym_POUND_COMMA),
	3620: uint16(753),
	3621: uint16(25),
	3622: uint16(aux_sym__skip_token1),
	3623: uint16(aux_sym_comment_token1),
	3624: uint16(anon_sym_POUND_PIPE),
	3625: uint16(anon_sym_POUND_SEMI),
	3626: uint16(sym__line_comment),
	3627: uint16(anon_sym_POUND_LT_LT),
	3628: uint16(aux_sym_regex_token1),
	3629: uint16(anon_sym_DQUOTE),
	3630: uint16(sym_character),
	3631: uint16(sym_keyword),
	3632: uint16(anon_sym_POUND_AMP),
	3633: uint16(anon_sym_LPAREN),
	3634: uint16(anon_sym_LBRACK),
	3635: uint16(anon_sym_LBRACE),
	3636: uint16(anon_sym_POUNDfl),
	3637: uint16(anon_sym_POUNDfx),
	3638: uint16(anon_sym_POUNDs),
	3639: uint16(anon_sym_POUNDhashalw),
	3640: uint16(anon_sym_POUNDhasheqv),
	3641: uint16(anon_sym_SQUOTE),
	3642: uint16(anon_sym_BQUOTE),
	3643: uint16(anon_sym_POUND_SQUOTE),
	3644: uint16(anon_sym_POUND_BQUOTE),
	3645: uint16(anon_sym_COMMA_AT),
	3646: uint16(anon_sym_POUND_COMMA_AT),
	3647: uint16(2),
	3648: uint16(767),
	3649: uint16(8),
	3650: uint16(sym_boolean),
	3651: uint16(anon_sym_POUND),
	3652: uint16(sym_number),
	3653: uint16(sym_symbol),
	3654: uint16(anon_sym_POUNDhash),
	3655: uint16(anon_sym_POUNDhasheq),
	3656: uint16(anon_sym_COMMA),
	3657: uint16(anon_sym_POUND_COMMA),
	3658: uint16(765),
	3659: uint16(25),
	3660: uint16(aux_sym__skip_token1),
	3661: uint16(aux_sym_comment_token1),
	3662: uint16(anon_sym_POUND_PIPE),
	3663: uint16(anon_sym_POUND_SEMI),
	3664: uint16(sym__line_comment),
	3665: uint16(anon_sym_POUND_LT_LT),
	3666: uint16(aux_sym_regex_token1),
	3667: uint16(anon_sym_DQUOTE),
	3668: uint16(sym_character),
	3669: uint16(sym_keyword),
	3670: uint16(anon_sym_POUND_AMP),
	3671: uint16(anon_sym_LPAREN),
	3672: uint16(anon_sym_LBRACK),
	3673: uint16(anon_sym_LBRACE),
	3674: uint16(anon_sym_POUNDfl),
	3675: uint16(anon_sym_POUNDfx),
	3676: uint16(anon_sym_POUNDs),
	3677: uint16(anon_sym_POUNDhashalw),
	3678: uint16(anon_sym_POUNDhasheqv),
	3679: uint16(anon_sym_SQUOTE),
	3680: uint16(anon_sym_BQUOTE),
	3681: uint16(anon_sym_POUND_SQUOTE),
	3682: uint16(anon_sym_POUND_BQUOTE),
	3683: uint16(anon_sym_COMMA_AT),
	3684: uint16(anon_sym_POUND_COMMA_AT),
	3685: uint16(2),
	3686: uint16(771),
	3687: uint16(8),
	3688: uint16(sym_boolean),
	3689: uint16(anon_sym_POUND),
	3690: uint16(sym_number),
	3691: uint16(sym_symbol),
	3692: uint16(anon_sym_POUNDhash),
	3693: uint16(anon_sym_POUNDhasheq),
	3694: uint16(anon_sym_COMMA),
	3695: uint16(anon_sym_POUND_COMMA),
	3696: uint16(769),
	3697: uint16(25),
	3698: uint16(aux_sym__skip_token1),
	3699: uint16(aux_sym_comment_token1),
	3700: uint16(anon_sym_POUND_PIPE),
	3701: uint16(anon_sym_POUND_SEMI),
	3702: uint16(sym__line_comment),
	3703: uint16(anon_sym_POUND_LT_LT),
	3704: uint16(aux_sym_regex_token1),
	3705: uint16(anon_sym_DQUOTE),
	3706: uint16(sym_character),
	3707: uint16(sym_keyword),
	3708: uint16(anon_sym_POUND_AMP),
	3709: uint16(anon_sym_LPAREN),
	3710: uint16(anon_sym_LBRACK),
	3711: uint16(anon_sym_LBRACE),
	3712: uint16(anon_sym_POUNDfl),
	3713: uint16(anon_sym_POUNDfx),
	3714: uint16(anon_sym_POUNDs),
	3715: uint16(anon_sym_POUNDhashalw),
	3716: uint16(anon_sym_POUNDhasheqv),
	3717: uint16(anon_sym_SQUOTE),
	3718: uint16(anon_sym_BQUOTE),
	3719: uint16(anon_sym_POUND_SQUOTE),
	3720: uint16(anon_sym_POUND_BQUOTE),
	3721: uint16(anon_sym_COMMA_AT),
	3722: uint16(anon_sym_POUND_COMMA_AT),
	3723: uint16(2),
	3724: uint16(779),
	3725: uint16(8),
	3726: uint16(sym_boolean),
	3727: uint16(anon_sym_POUND),
	3728: uint16(sym_number),
	3729: uint16(sym_symbol),
	3730: uint16(anon_sym_POUNDhash),
	3731: uint16(anon_sym_POUNDhasheq),
	3732: uint16(anon_sym_COMMA),
	3733: uint16(anon_sym_POUND_COMMA),
	3734: uint16(777),
	3735: uint16(25),
	3736: uint16(aux_sym__skip_token1),
	3737: uint16(aux_sym_comment_token1),
	3738: uint16(anon_sym_POUND_PIPE),
	3739: uint16(anon_sym_POUND_SEMI),
	3740: uint16(sym__line_comment),
	3741: uint16(anon_sym_POUND_LT_LT),
	3742: uint16(aux_sym_regex_token1),
	3743: uint16(anon_sym_DQUOTE),
	3744: uint16(sym_character),
	3745: uint16(sym_keyword),
	3746: uint16(anon_sym_POUND_AMP),
	3747: uint16(anon_sym_LPAREN),
	3748: uint16(anon_sym_LBRACK),
	3749: uint16(anon_sym_LBRACE),
	3750: uint16(anon_sym_POUNDfl),
	3751: uint16(anon_sym_POUNDfx),
	3752: uint16(anon_sym_POUNDs),
	3753: uint16(anon_sym_POUNDhashalw),
	3754: uint16(anon_sym_POUNDhasheqv),
	3755: uint16(anon_sym_SQUOTE),
	3756: uint16(anon_sym_BQUOTE),
	3757: uint16(anon_sym_POUND_SQUOTE),
	3758: uint16(anon_sym_POUND_BQUOTE),
	3759: uint16(anon_sym_COMMA_AT),
	3760: uint16(anon_sym_POUND_COMMA_AT),
	3761: uint16(2),
	3762: uint16(787),
	3763: uint16(8),
	3764: uint16(sym_boolean),
	3765: uint16(anon_sym_POUND),
	3766: uint16(sym_number),
	3767: uint16(sym_symbol),
	3768: uint16(anon_sym_POUNDhash),
	3769: uint16(anon_sym_POUNDhasheq),
	3770: uint16(anon_sym_COMMA),
	3771: uint16(anon_sym_POUND_COMMA),
	3772: uint16(785),
	3773: uint16(25),
	3774: uint16(aux_sym__skip_token1),
	3775: uint16(aux_sym_comment_token1),
	3776: uint16(anon_sym_POUND_PIPE),
	3777: uint16(anon_sym_POUND_SEMI),
	3778: uint16(sym__line_comment),
	3779: uint16(anon_sym_POUND_LT_LT),
	3780: uint16(aux_sym_regex_token1),
	3781: uint16(anon_sym_DQUOTE),
	3782: uint16(sym_character),
	3783: uint16(sym_keyword),
	3784: uint16(anon_sym_POUND_AMP),
	3785: uint16(anon_sym_LPAREN),
	3786: uint16(anon_sym_LBRACK),
	3787: uint16(anon_sym_LBRACE),
	3788: uint16(anon_sym_POUNDfl),
	3789: uint16(anon_sym_POUNDfx),
	3790: uint16(anon_sym_POUNDs),
	3791: uint16(anon_sym_POUNDhashalw),
	3792: uint16(anon_sym_POUNDhasheqv),
	3793: uint16(anon_sym_SQUOTE),
	3794: uint16(anon_sym_BQUOTE),
	3795: uint16(anon_sym_POUND_SQUOTE),
	3796: uint16(anon_sym_POUND_BQUOTE),
	3797: uint16(anon_sym_COMMA_AT),
	3798: uint16(anon_sym_POUND_COMMA_AT),
	3799: uint16(2),
	3800: uint16(795),
	3801: uint16(8),
	3802: uint16(sym_boolean),
	3803: uint16(anon_sym_POUND),
	3804: uint16(sym_number),
	3805: uint16(sym_symbol),
	3806: uint16(anon_sym_POUNDhash),
	3807: uint16(anon_sym_POUNDhasheq),
	3808: uint16(anon_sym_COMMA),
	3809: uint16(anon_sym_POUND_COMMA),
	3810: uint16(793),
	3811: uint16(25),
	3812: uint16(aux_sym__skip_token1),
	3813: uint16(aux_sym_comment_token1),
	3814: uint16(anon_sym_POUND_PIPE),
	3815: uint16(anon_sym_POUND_SEMI),
	3816: uint16(sym__line_comment),
	3817: uint16(anon_sym_POUND_LT_LT),
	3818: uint16(aux_sym_regex_token1),
	3819: uint16(anon_sym_DQUOTE),
	3820: uint16(sym_character),
	3821: uint16(sym_keyword),
	3822: uint16(anon_sym_POUND_AMP),
	3823: uint16(anon_sym_LPAREN),
	3824: uint16(anon_sym_LBRACK),
	3825: uint16(anon_sym_LBRACE),
	3826: uint16(anon_sym_POUNDfl),
	3827: uint16(anon_sym_POUNDfx),
	3828: uint16(anon_sym_POUNDs),
	3829: uint16(anon_sym_POUNDhashalw),
	3830: uint16(anon_sym_POUNDhasheqv),
	3831: uint16(anon_sym_SQUOTE),
	3832: uint16(anon_sym_BQUOTE),
	3833: uint16(anon_sym_POUND_SQUOTE),
	3834: uint16(anon_sym_POUND_BQUOTE),
	3835: uint16(anon_sym_COMMA_AT),
	3836: uint16(anon_sym_POUND_COMMA_AT),
	3837: uint16(2),
	3838: uint16(799),
	3839: uint16(8),
	3840: uint16(sym_boolean),
	3841: uint16(anon_sym_POUND),
	3842: uint16(sym_number),
	3843: uint16(sym_symbol),
	3844: uint16(anon_sym_POUNDhash),
	3845: uint16(anon_sym_POUNDhasheq),
	3846: uint16(anon_sym_COMMA),
	3847: uint16(anon_sym_POUND_COMMA),
	3848: uint16(797),
	3849: uint16(25),
	3850: uint16(aux_sym__skip_token1),
	3851: uint16(aux_sym_comment_token1),
	3852: uint16(anon_sym_POUND_PIPE),
	3853: uint16(anon_sym_POUND_SEMI),
	3854: uint16(sym__line_comment),
	3855: uint16(anon_sym_POUND_LT_LT),
	3856: uint16(aux_sym_regex_token1),
	3857: uint16(anon_sym_DQUOTE),
	3858: uint16(sym_character),
	3859: uint16(sym_keyword),
	3860: uint16(anon_sym_POUND_AMP),
	3861: uint16(anon_sym_LPAREN),
	3862: uint16(anon_sym_LBRACK),
	3863: uint16(anon_sym_LBRACE),
	3864: uint16(anon_sym_POUNDfl),
	3865: uint16(anon_sym_POUNDfx),
	3866: uint16(anon_sym_POUNDs),
	3867: uint16(anon_sym_POUNDhashalw),
	3868: uint16(anon_sym_POUNDhasheqv),
	3869: uint16(anon_sym_SQUOTE),
	3870: uint16(anon_sym_BQUOTE),
	3871: uint16(anon_sym_POUND_SQUOTE),
	3872: uint16(anon_sym_POUND_BQUOTE),
	3873: uint16(anon_sym_COMMA_AT),
	3874: uint16(anon_sym_POUND_COMMA_AT),
	3875: uint16(2),
	3876: uint16(807),
	3877: uint16(8),
	3878: uint16(sym_boolean),
	3879: uint16(anon_sym_POUND),
	3880: uint16(sym_number),
	3881: uint16(sym_symbol),
	3882: uint16(anon_sym_POUNDhash),
	3883: uint16(anon_sym_POUNDhasheq),
	3884: uint16(anon_sym_COMMA),
	3885: uint16(anon_sym_POUND_COMMA),
	3886: uint16(805),
	3887: uint16(25),
	3888: uint16(aux_sym__skip_token1),
	3889: uint16(aux_sym_comment_token1),
	3890: uint16(anon_sym_POUND_PIPE),
	3891: uint16(anon_sym_POUND_SEMI),
	3892: uint16(sym__line_comment),
	3893: uint16(anon_sym_POUND_LT_LT),
	3894: uint16(aux_sym_regex_token1),
	3895: uint16(anon_sym_DQUOTE),
	3896: uint16(sym_character),
	3897: uint16(sym_keyword),
	3898: uint16(anon_sym_POUND_AMP),
	3899: uint16(anon_sym_LPAREN),
	3900: uint16(anon_sym_LBRACK),
	3901: uint16(anon_sym_LBRACE),
	3902: uint16(anon_sym_POUNDfl),
	3903: uint16(anon_sym_POUNDfx),
	3904: uint16(anon_sym_POUNDs),
	3905: uint16(anon_sym_POUNDhashalw),
	3906: uint16(anon_sym_POUNDhasheqv),
	3907: uint16(anon_sym_SQUOTE),
	3908: uint16(anon_sym_BQUOTE),
	3909: uint16(anon_sym_POUND_SQUOTE),
	3910: uint16(anon_sym_POUND_BQUOTE),
	3911: uint16(anon_sym_COMMA_AT),
	3912: uint16(anon_sym_POUND_COMMA_AT),
	3913: uint16(2),
	3914: uint16(763),
	3915: uint16(8),
	3916: uint16(sym_boolean),
	3917: uint16(anon_sym_POUND),
	3918: uint16(sym_number),
	3919: uint16(sym_symbol),
	3920: uint16(anon_sym_POUNDhash),
	3921: uint16(anon_sym_POUNDhasheq),
	3922: uint16(anon_sym_COMMA),
	3923: uint16(anon_sym_POUND_COMMA),
	3924: uint16(761),
	3925: uint16(25),
	3926: uint16(aux_sym__skip_token1),
	3927: uint16(aux_sym_comment_token1),
	3928: uint16(anon_sym_POUND_PIPE),
	3929: uint16(anon_sym_POUND_SEMI),
	3930: uint16(sym__line_comment),
	3931: uint16(anon_sym_POUND_LT_LT),
	3932: uint16(aux_sym_regex_token1),
	3933: uint16(anon_sym_DQUOTE),
	3934: uint16(sym_character),
	3935: uint16(sym_keyword),
	3936: uint16(anon_sym_POUND_AMP),
	3937: uint16(anon_sym_LPAREN),
	3938: uint16(anon_sym_LBRACK),
	3939: uint16(anon_sym_LBRACE),
	3940: uint16(anon_sym_POUNDfl),
	3941: uint16(anon_sym_POUNDfx),
	3942: uint16(anon_sym_POUNDs),
	3943: uint16(anon_sym_POUNDhashalw),
	3944: uint16(anon_sym_POUNDhasheqv),
	3945: uint16(anon_sym_SQUOTE),
	3946: uint16(anon_sym_BQUOTE),
	3947: uint16(anon_sym_POUND_SQUOTE),
	3948: uint16(anon_sym_POUND_BQUOTE),
	3949: uint16(anon_sym_COMMA_AT),
	3950: uint16(anon_sym_POUND_COMMA_AT),
	3951: uint16(2),
	3952: uint16(743),
	3953: uint16(8),
	3954: uint16(sym_boolean),
	3955: uint16(anon_sym_POUND),
	3956: uint16(sym_number),
	3957: uint16(sym_symbol),
	3958: uint16(anon_sym_POUNDhash),
	3959: uint16(anon_sym_POUNDhasheq),
	3960: uint16(anon_sym_COMMA),
	3961: uint16(anon_sym_POUND_COMMA),
	3962: uint16(741),
	3963: uint16(25),
	3964: uint16(aux_sym__skip_token1),
	3965: uint16(aux_sym_comment_token1),
	3966: uint16(anon_sym_POUND_PIPE),
	3967: uint16(anon_sym_POUND_SEMI),
	3968: uint16(sym__line_comment),
	3969: uint16(anon_sym_POUND_LT_LT),
	3970: uint16(aux_sym_regex_token1),
	3971: uint16(anon_sym_DQUOTE),
	3972: uint16(sym_character),
	3973: uint16(sym_keyword),
	3974: uint16(anon_sym_POUND_AMP),
	3975: uint16(anon_sym_LPAREN),
	3976: uint16(anon_sym_LBRACK),
	3977: uint16(anon_sym_LBRACE),
	3978: uint16(anon_sym_POUNDfl),
	3979: uint16(anon_sym_POUNDfx),
	3980: uint16(anon_sym_POUNDs),
	3981: uint16(anon_sym_POUNDhashalw),
	3982: uint16(anon_sym_POUNDhasheqv),
	3983: uint16(anon_sym_SQUOTE),
	3984: uint16(anon_sym_BQUOTE),
	3985: uint16(anon_sym_POUND_SQUOTE),
	3986: uint16(anon_sym_POUND_BQUOTE),
	3987: uint16(anon_sym_COMMA_AT),
	3988: uint16(anon_sym_POUND_COMMA_AT),
	3989: uint16(2),
	3990: uint16(823),
	3991: uint16(8),
	3992: uint16(sym_boolean),
	3993: uint16(anon_sym_POUND),
	3994: uint16(sym_number),
	3995: uint16(sym_symbol),
	3996: uint16(anon_sym_POUNDhash),
	3997: uint16(anon_sym_POUNDhasheq),
	3998: uint16(anon_sym_COMMA),
	3999: uint16(anon_sym_POUND_COMMA),
	4000: uint16(821),
	4001: uint16(25),
	4002: uint16(aux_sym__skip_token1),
	4003: uint16(aux_sym_comment_token1),
	4004: uint16(anon_sym_POUND_PIPE),
	4005: uint16(anon_sym_POUND_SEMI),
	4006: uint16(sym__line_comment),
	4007: uint16(anon_sym_POUND_LT_LT),
	4008: uint16(aux_sym_regex_token1),
	4009: uint16(anon_sym_DQUOTE),
	4010: uint16(sym_character),
	4011: uint16(sym_keyword),
	4012: uint16(anon_sym_POUND_AMP),
	4013: uint16(anon_sym_LPAREN),
	4014: uint16(anon_sym_LBRACK),
	4015: uint16(anon_sym_LBRACE),
	4016: uint16(anon_sym_POUNDfl),
	4017: uint16(anon_sym_POUNDfx),
	4018: uint16(anon_sym_POUNDs),
	4019: uint16(anon_sym_POUNDhashalw),
	4020: uint16(anon_sym_POUNDhasheqv),
	4021: uint16(anon_sym_SQUOTE),
	4022: uint16(anon_sym_BQUOTE),
	4023: uint16(anon_sym_POUND_SQUOTE),
	4024: uint16(anon_sym_POUND_BQUOTE),
	4025: uint16(anon_sym_COMMA_AT),
	4026: uint16(anon_sym_POUND_COMMA_AT),
	4027: uint16(2),
	4028: uint16(859),
	4029: uint16(8),
	4030: uint16(sym_boolean),
	4031: uint16(anon_sym_POUND),
	4032: uint16(sym_number),
	4033: uint16(sym_symbol),
	4034: uint16(anon_sym_POUNDhash),
	4035: uint16(anon_sym_POUNDhasheq),
	4036: uint16(anon_sym_COMMA),
	4037: uint16(anon_sym_POUND_COMMA),
	4038: uint16(857),
	4039: uint16(25),
	4040: uint16(aux_sym__skip_token1),
	4041: uint16(aux_sym_comment_token1),
	4042: uint16(anon_sym_POUND_PIPE),
	4043: uint16(anon_sym_POUND_SEMI),
	4044: uint16(sym__line_comment),
	4045: uint16(anon_sym_POUND_LT_LT),
	4046: uint16(aux_sym_regex_token1),
	4047: uint16(anon_sym_DQUOTE),
	4048: uint16(sym_character),
	4049: uint16(sym_keyword),
	4050: uint16(anon_sym_POUND_AMP),
	4051: uint16(anon_sym_LPAREN),
	4052: uint16(anon_sym_LBRACK),
	4053: uint16(anon_sym_LBRACE),
	4054: uint16(anon_sym_POUNDfl),
	4055: uint16(anon_sym_POUNDfx),
	4056: uint16(anon_sym_POUNDs),
	4057: uint16(anon_sym_POUNDhashalw),
	4058: uint16(anon_sym_POUNDhasheqv),
	4059: uint16(anon_sym_SQUOTE),
	4060: uint16(anon_sym_BQUOTE),
	4061: uint16(anon_sym_POUND_SQUOTE),
	4062: uint16(anon_sym_POUND_BQUOTE),
	4063: uint16(anon_sym_COMMA_AT),
	4064: uint16(anon_sym_POUND_COMMA_AT),
	4065: uint16(2),
	4066: uint16(839),
	4067: uint16(8),
	4068: uint16(sym_boolean),
	4069: uint16(anon_sym_POUND),
	4070: uint16(sym_number),
	4071: uint16(sym_symbol),
	4072: uint16(anon_sym_POUNDhash),
	4073: uint16(anon_sym_POUNDhasheq),
	4074: uint16(anon_sym_COMMA),
	4075: uint16(anon_sym_POUND_COMMA),
	4076: uint16(837),
	4077: uint16(25),
	4078: uint16(aux_sym__skip_token1),
	4079: uint16(aux_sym_comment_token1),
	4080: uint16(anon_sym_POUND_PIPE),
	4081: uint16(anon_sym_POUND_SEMI),
	4082: uint16(sym__line_comment),
	4083: uint16(anon_sym_POUND_LT_LT),
	4084: uint16(aux_sym_regex_token1),
	4085: uint16(anon_sym_DQUOTE),
	4086: uint16(sym_character),
	4087: uint16(sym_keyword),
	4088: uint16(anon_sym_POUND_AMP),
	4089: uint16(anon_sym_LPAREN),
	4090: uint16(anon_sym_LBRACK),
	4091: uint16(anon_sym_LBRACE),
	4092: uint16(anon_sym_POUNDfl),
	4093: uint16(anon_sym_POUNDfx),
	4094: uint16(anon_sym_POUNDs),
	4095: uint16(anon_sym_POUNDhashalw),
	4096: uint16(anon_sym_POUNDhasheqv),
	4097: uint16(anon_sym_SQUOTE),
	4098: uint16(anon_sym_BQUOTE),
	4099: uint16(anon_sym_POUND_SQUOTE),
	4100: uint16(anon_sym_POUND_BQUOTE),
	4101: uint16(anon_sym_COMMA_AT),
	4102: uint16(anon_sym_POUND_COMMA_AT),
	4103: uint16(2),
	4104: uint16(863),
	4105: uint16(8),
	4106: uint16(sym_boolean),
	4107: uint16(anon_sym_POUND),
	4108: uint16(sym_number),
	4109: uint16(sym_symbol),
	4110: uint16(anon_sym_POUNDhash),
	4111: uint16(anon_sym_POUNDhasheq),
	4112: uint16(anon_sym_COMMA),
	4113: uint16(anon_sym_POUND_COMMA),
	4114: uint16(861),
	4115: uint16(25),
	4116: uint16(aux_sym__skip_token1),
	4117: uint16(aux_sym_comment_token1),
	4118: uint16(anon_sym_POUND_PIPE),
	4119: uint16(anon_sym_POUND_SEMI),
	4120: uint16(sym__line_comment),
	4121: uint16(anon_sym_POUND_LT_LT),
	4122: uint16(aux_sym_regex_token1),
	4123: uint16(anon_sym_DQUOTE),
	4124: uint16(sym_character),
	4125: uint16(sym_keyword),
	4126: uint16(anon_sym_POUND_AMP),
	4127: uint16(anon_sym_LPAREN),
	4128: uint16(anon_sym_LBRACK),
	4129: uint16(anon_sym_LBRACE),
	4130: uint16(anon_sym_POUNDfl),
	4131: uint16(anon_sym_POUNDfx),
	4132: uint16(anon_sym_POUNDs),
	4133: uint16(anon_sym_POUNDhashalw),
	4134: uint16(anon_sym_POUNDhasheqv),
	4135: uint16(anon_sym_SQUOTE),
	4136: uint16(anon_sym_BQUOTE),
	4137: uint16(anon_sym_POUND_SQUOTE),
	4138: uint16(anon_sym_POUND_BQUOTE),
	4139: uint16(anon_sym_COMMA_AT),
	4140: uint16(anon_sym_POUND_COMMA_AT),
	4141: uint16(2),
	4142: uint16(867),
	4143: uint16(8),
	4144: uint16(sym_boolean),
	4145: uint16(anon_sym_POUND),
	4146: uint16(sym_number),
	4147: uint16(sym_symbol),
	4148: uint16(anon_sym_POUNDhash),
	4149: uint16(anon_sym_POUNDhasheq),
	4150: uint16(anon_sym_COMMA),
	4151: uint16(anon_sym_POUND_COMMA),
	4152: uint16(865),
	4153: uint16(25),
	4154: uint16(aux_sym__skip_token1),
	4155: uint16(aux_sym_comment_token1),
	4156: uint16(anon_sym_POUND_PIPE),
	4157: uint16(anon_sym_POUND_SEMI),
	4158: uint16(sym__line_comment),
	4159: uint16(anon_sym_POUND_LT_LT),
	4160: uint16(aux_sym_regex_token1),
	4161: uint16(anon_sym_DQUOTE),
	4162: uint16(sym_character),
	4163: uint16(sym_keyword),
	4164: uint16(anon_sym_POUND_AMP),
	4165: uint16(anon_sym_LPAREN),
	4166: uint16(anon_sym_LBRACK),
	4167: uint16(anon_sym_LBRACE),
	4168: uint16(anon_sym_POUNDfl),
	4169: uint16(anon_sym_POUNDfx),
	4170: uint16(anon_sym_POUNDs),
	4171: uint16(anon_sym_POUNDhashalw),
	4172: uint16(anon_sym_POUNDhasheqv),
	4173: uint16(anon_sym_SQUOTE),
	4174: uint16(anon_sym_BQUOTE),
	4175: uint16(anon_sym_POUND_SQUOTE),
	4176: uint16(anon_sym_POUND_BQUOTE),
	4177: uint16(anon_sym_COMMA_AT),
	4178: uint16(anon_sym_POUND_COMMA_AT),
	4179: uint16(2),
	4180: uint16(871),
	4181: uint16(8),
	4182: uint16(sym_boolean),
	4183: uint16(anon_sym_POUND),
	4184: uint16(sym_number),
	4185: uint16(sym_symbol),
	4186: uint16(anon_sym_POUNDhash),
	4187: uint16(anon_sym_POUNDhasheq),
	4188: uint16(anon_sym_COMMA),
	4189: uint16(anon_sym_POUND_COMMA),
	4190: uint16(869),
	4191: uint16(25),
	4192: uint16(aux_sym__skip_token1),
	4193: uint16(aux_sym_comment_token1),
	4194: uint16(anon_sym_POUND_PIPE),
	4195: uint16(anon_sym_POUND_SEMI),
	4196: uint16(sym__line_comment),
	4197: uint16(anon_sym_POUND_LT_LT),
	4198: uint16(aux_sym_regex_token1),
	4199: uint16(anon_sym_DQUOTE),
	4200: uint16(sym_character),
	4201: uint16(sym_keyword),
	4202: uint16(anon_sym_POUND_AMP),
	4203: uint16(anon_sym_LPAREN),
	4204: uint16(anon_sym_LBRACK),
	4205: uint16(anon_sym_LBRACE),
	4206: uint16(anon_sym_POUNDfl),
	4207: uint16(anon_sym_POUNDfx),
	4208: uint16(anon_sym_POUNDs),
	4209: uint16(anon_sym_POUNDhashalw),
	4210: uint16(anon_sym_POUNDhasheqv),
	4211: uint16(anon_sym_SQUOTE),
	4212: uint16(anon_sym_BQUOTE),
	4213: uint16(anon_sym_POUND_SQUOTE),
	4214: uint16(anon_sym_POUND_BQUOTE),
	4215: uint16(anon_sym_COMMA_AT),
	4216: uint16(anon_sym_POUND_COMMA_AT),
	4217: uint16(2),
	4218: uint16(875),
	4219: uint16(8),
	4220: uint16(sym_boolean),
	4221: uint16(anon_sym_POUND),
	4222: uint16(sym_number),
	4223: uint16(sym_symbol),
	4224: uint16(anon_sym_POUNDhash),
	4225: uint16(anon_sym_POUNDhasheq),
	4226: uint16(anon_sym_COMMA),
	4227: uint16(anon_sym_POUND_COMMA),
	4228: uint16(873),
	4229: uint16(25),
	4230: uint16(aux_sym__skip_token1),
	4231: uint16(aux_sym_comment_token1),
	4232: uint16(anon_sym_POUND_PIPE),
	4233: uint16(anon_sym_POUND_SEMI),
	4234: uint16(sym__line_comment),
	4235: uint16(anon_sym_POUND_LT_LT),
	4236: uint16(aux_sym_regex_token1),
	4237: uint16(anon_sym_DQUOTE),
	4238: uint16(sym_character),
	4239: uint16(sym_keyword),
	4240: uint16(anon_sym_POUND_AMP),
	4241: uint16(anon_sym_LPAREN),
	4242: uint16(anon_sym_LBRACK),
	4243: uint16(anon_sym_LBRACE),
	4244: uint16(anon_sym_POUNDfl),
	4245: uint16(anon_sym_POUNDfx),
	4246: uint16(anon_sym_POUNDs),
	4247: uint16(anon_sym_POUNDhashalw),
	4248: uint16(anon_sym_POUNDhasheqv),
	4249: uint16(anon_sym_SQUOTE),
	4250: uint16(anon_sym_BQUOTE),
	4251: uint16(anon_sym_POUND_SQUOTE),
	4252: uint16(anon_sym_POUND_BQUOTE),
	4253: uint16(anon_sym_COMMA_AT),
	4254: uint16(anon_sym_POUND_COMMA_AT),
	4255: uint16(2),
	4256: uint16(879),
	4257: uint16(8),
	4258: uint16(sym_boolean),
	4259: uint16(anon_sym_POUND),
	4260: uint16(sym_number),
	4261: uint16(sym_symbol),
	4262: uint16(anon_sym_POUNDhash),
	4263: uint16(anon_sym_POUNDhasheq),
	4264: uint16(anon_sym_COMMA),
	4265: uint16(anon_sym_POUND_COMMA),
	4266: uint16(877),
	4267: uint16(25),
	4268: uint16(aux_sym__skip_token1),
	4269: uint16(aux_sym_comment_token1),
	4270: uint16(anon_sym_POUND_PIPE),
	4271: uint16(anon_sym_POUND_SEMI),
	4272: uint16(sym__line_comment),
	4273: uint16(anon_sym_POUND_LT_LT),
	4274: uint16(aux_sym_regex_token1),
	4275: uint16(anon_sym_DQUOTE),
	4276: uint16(sym_character),
	4277: uint16(sym_keyword),
	4278: uint16(anon_sym_POUND_AMP),
	4279: uint16(anon_sym_LPAREN),
	4280: uint16(anon_sym_LBRACK),
	4281: uint16(anon_sym_LBRACE),
	4282: uint16(anon_sym_POUNDfl),
	4283: uint16(anon_sym_POUNDfx),
	4284: uint16(anon_sym_POUNDs),
	4285: uint16(anon_sym_POUNDhashalw),
	4286: uint16(anon_sym_POUNDhasheqv),
	4287: uint16(anon_sym_SQUOTE),
	4288: uint16(anon_sym_BQUOTE),
	4289: uint16(anon_sym_POUND_SQUOTE),
	4290: uint16(anon_sym_POUND_BQUOTE),
	4291: uint16(anon_sym_COMMA_AT),
	4292: uint16(anon_sym_POUND_COMMA_AT),
	4293: uint16(2),
	4294: uint16(835),
	4295: uint16(8),
	4296: uint16(sym_boolean),
	4297: uint16(anon_sym_POUND),
	4298: uint16(sym_number),
	4299: uint16(sym_symbol),
	4300: uint16(anon_sym_POUNDhash),
	4301: uint16(anon_sym_POUNDhasheq),
	4302: uint16(anon_sym_COMMA),
	4303: uint16(anon_sym_POUND_COMMA),
	4304: uint16(833),
	4305: uint16(25),
	4306: uint16(aux_sym__skip_token1),
	4307: uint16(aux_sym_comment_token1),
	4308: uint16(anon_sym_POUND_PIPE),
	4309: uint16(anon_sym_POUND_SEMI),
	4310: uint16(sym__line_comment),
	4311: uint16(anon_sym_POUND_LT_LT),
	4312: uint16(aux_sym_regex_token1),
	4313: uint16(anon_sym_DQUOTE),
	4314: uint16(sym_character),
	4315: uint16(sym_keyword),
	4316: uint16(anon_sym_POUND_AMP),
	4317: uint16(anon_sym_LPAREN),
	4318: uint16(anon_sym_LBRACK),
	4319: uint16(anon_sym_LBRACE),
	4320: uint16(anon_sym_POUNDfl),
	4321: uint16(anon_sym_POUNDfx),
	4322: uint16(anon_sym_POUNDs),
	4323: uint16(anon_sym_POUNDhashalw),
	4324: uint16(anon_sym_POUNDhasheqv),
	4325: uint16(anon_sym_SQUOTE),
	4326: uint16(anon_sym_BQUOTE),
	4327: uint16(anon_sym_POUND_SQUOTE),
	4328: uint16(anon_sym_POUND_BQUOTE),
	4329: uint16(anon_sym_COMMA_AT),
	4330: uint16(anon_sym_POUND_COMMA_AT),
	4331: uint16(2),
	4332: uint16(887),
	4333: uint16(8),
	4334: uint16(sym_boolean),
	4335: uint16(anon_sym_POUND),
	4336: uint16(sym_number),
	4337: uint16(sym_symbol),
	4338: uint16(anon_sym_POUNDhash),
	4339: uint16(anon_sym_POUNDhasheq),
	4340: uint16(anon_sym_COMMA),
	4341: uint16(anon_sym_POUND_COMMA),
	4342: uint16(885),
	4343: uint16(25),
	4344: uint16(aux_sym__skip_token1),
	4345: uint16(aux_sym_comment_token1),
	4346: uint16(anon_sym_POUND_PIPE),
	4347: uint16(anon_sym_POUND_SEMI),
	4348: uint16(sym__line_comment),
	4349: uint16(anon_sym_POUND_LT_LT),
	4350: uint16(aux_sym_regex_token1),
	4351: uint16(anon_sym_DQUOTE),
	4352: uint16(sym_character),
	4353: uint16(sym_keyword),
	4354: uint16(anon_sym_POUND_AMP),
	4355: uint16(anon_sym_LPAREN),
	4356: uint16(anon_sym_LBRACK),
	4357: uint16(anon_sym_LBRACE),
	4358: uint16(anon_sym_POUNDfl),
	4359: uint16(anon_sym_POUNDfx),
	4360: uint16(anon_sym_POUNDs),
	4361: uint16(anon_sym_POUNDhashalw),
	4362: uint16(anon_sym_POUNDhasheqv),
	4363: uint16(anon_sym_SQUOTE),
	4364: uint16(anon_sym_BQUOTE),
	4365: uint16(anon_sym_POUND_SQUOTE),
	4366: uint16(anon_sym_POUND_BQUOTE),
	4367: uint16(anon_sym_COMMA_AT),
	4368: uint16(anon_sym_POUND_COMMA_AT),
	4369: uint16(2),
	4370: uint16(891),
	4371: uint16(8),
	4372: uint16(sym_boolean),
	4373: uint16(anon_sym_POUND),
	4374: uint16(sym_number),
	4375: uint16(sym_symbol),
	4376: uint16(anon_sym_POUNDhash),
	4377: uint16(anon_sym_POUNDhasheq),
	4378: uint16(anon_sym_COMMA),
	4379: uint16(anon_sym_POUND_COMMA),
	4380: uint16(889),
	4381: uint16(25),
	4382: uint16(aux_sym__skip_token1),
	4383: uint16(aux_sym_comment_token1),
	4384: uint16(anon_sym_POUND_PIPE),
	4385: uint16(anon_sym_POUND_SEMI),
	4386: uint16(sym__line_comment),
	4387: uint16(anon_sym_POUND_LT_LT),
	4388: uint16(aux_sym_regex_token1),
	4389: uint16(anon_sym_DQUOTE),
	4390: uint16(sym_character),
	4391: uint16(sym_keyword),
	4392: uint16(anon_sym_POUND_AMP),
	4393: uint16(anon_sym_LPAREN),
	4394: uint16(anon_sym_LBRACK),
	4395: uint16(anon_sym_LBRACE),
	4396: uint16(anon_sym_POUNDfl),
	4397: uint16(anon_sym_POUNDfx),
	4398: uint16(anon_sym_POUNDs),
	4399: uint16(anon_sym_POUNDhashalw),
	4400: uint16(anon_sym_POUNDhasheqv),
	4401: uint16(anon_sym_SQUOTE),
	4402: uint16(anon_sym_BQUOTE),
	4403: uint16(anon_sym_POUND_SQUOTE),
	4404: uint16(anon_sym_POUND_BQUOTE),
	4405: uint16(anon_sym_COMMA_AT),
	4406: uint16(anon_sym_POUND_COMMA_AT),
	4407: uint16(2),
	4408: uint16(895),
	4409: uint16(8),
	4410: uint16(sym_boolean),
	4411: uint16(anon_sym_POUND),
	4412: uint16(sym_number),
	4413: uint16(sym_symbol),
	4414: uint16(anon_sym_POUNDhash),
	4415: uint16(anon_sym_POUNDhasheq),
	4416: uint16(anon_sym_COMMA),
	4417: uint16(anon_sym_POUND_COMMA),
	4418: uint16(893),
	4419: uint16(25),
	4420: uint16(aux_sym__skip_token1),
	4421: uint16(aux_sym_comment_token1),
	4422: uint16(anon_sym_POUND_PIPE),
	4423: uint16(anon_sym_POUND_SEMI),
	4424: uint16(sym__line_comment),
	4425: uint16(anon_sym_POUND_LT_LT),
	4426: uint16(aux_sym_regex_token1),
	4427: uint16(anon_sym_DQUOTE),
	4428: uint16(sym_character),
	4429: uint16(sym_keyword),
	4430: uint16(anon_sym_POUND_AMP),
	4431: uint16(anon_sym_LPAREN),
	4432: uint16(anon_sym_LBRACK),
	4433: uint16(anon_sym_LBRACE),
	4434: uint16(anon_sym_POUNDfl),
	4435: uint16(anon_sym_POUNDfx),
	4436: uint16(anon_sym_POUNDs),
	4437: uint16(anon_sym_POUNDhashalw),
	4438: uint16(anon_sym_POUNDhasheqv),
	4439: uint16(anon_sym_SQUOTE),
	4440: uint16(anon_sym_BQUOTE),
	4441: uint16(anon_sym_POUND_SQUOTE),
	4442: uint16(anon_sym_POUND_BQUOTE),
	4443: uint16(anon_sym_COMMA_AT),
	4444: uint16(anon_sym_POUND_COMMA_AT),
	4445: uint16(2),
	4446: uint16(815),
	4447: uint16(8),
	4448: uint16(sym_boolean),
	4449: uint16(anon_sym_POUND),
	4450: uint16(sym_number),
	4451: uint16(sym_symbol),
	4452: uint16(anon_sym_POUNDhash),
	4453: uint16(anon_sym_POUNDhasheq),
	4454: uint16(anon_sym_COMMA),
	4455: uint16(anon_sym_POUND_COMMA),
	4456: uint16(813),
	4457: uint16(25),
	4458: uint16(aux_sym__skip_token1),
	4459: uint16(aux_sym_comment_token1),
	4460: uint16(anon_sym_POUND_PIPE),
	4461: uint16(anon_sym_POUND_SEMI),
	4462: uint16(sym__line_comment),
	4463: uint16(anon_sym_POUND_LT_LT),
	4464: uint16(aux_sym_regex_token1),
	4465: uint16(anon_sym_DQUOTE),
	4466: uint16(sym_character),
	4467: uint16(sym_keyword),
	4468: uint16(anon_sym_POUND_AMP),
	4469: uint16(anon_sym_LPAREN),
	4470: uint16(anon_sym_LBRACK),
	4471: uint16(anon_sym_LBRACE),
	4472: uint16(anon_sym_POUNDfl),
	4473: uint16(anon_sym_POUNDfx),
	4474: uint16(anon_sym_POUNDs),
	4475: uint16(anon_sym_POUNDhashalw),
	4476: uint16(anon_sym_POUNDhasheqv),
	4477: uint16(anon_sym_SQUOTE),
	4478: uint16(anon_sym_BQUOTE),
	4479: uint16(anon_sym_POUND_SQUOTE),
	4480: uint16(anon_sym_POUND_BQUOTE),
	4481: uint16(anon_sym_COMMA_AT),
	4482: uint16(anon_sym_POUND_COMMA_AT),
	4483: uint16(2),
	4484: uint16(819),
	4485: uint16(8),
	4486: uint16(sym_boolean),
	4487: uint16(anon_sym_POUND),
	4488: uint16(sym_number),
	4489: uint16(sym_symbol),
	4490: uint16(anon_sym_POUNDhash),
	4491: uint16(anon_sym_POUNDhasheq),
	4492: uint16(anon_sym_COMMA),
	4493: uint16(anon_sym_POUND_COMMA),
	4494: uint16(817),
	4495: uint16(25),
	4496: uint16(aux_sym__skip_token1),
	4497: uint16(aux_sym_comment_token1),
	4498: uint16(anon_sym_POUND_PIPE),
	4499: uint16(anon_sym_POUND_SEMI),
	4500: uint16(sym__line_comment),
	4501: uint16(anon_sym_POUND_LT_LT),
	4502: uint16(aux_sym_regex_token1),
	4503: uint16(anon_sym_DQUOTE),
	4504: uint16(sym_character),
	4505: uint16(sym_keyword),
	4506: uint16(anon_sym_POUND_AMP),
	4507: uint16(anon_sym_LPAREN),
	4508: uint16(anon_sym_LBRACK),
	4509: uint16(anon_sym_LBRACE),
	4510: uint16(anon_sym_POUNDfl),
	4511: uint16(anon_sym_POUNDfx),
	4512: uint16(anon_sym_POUNDs),
	4513: uint16(anon_sym_POUNDhashalw),
	4514: uint16(anon_sym_POUNDhasheqv),
	4515: uint16(anon_sym_SQUOTE),
	4516: uint16(anon_sym_BQUOTE),
	4517: uint16(anon_sym_POUND_SQUOTE),
	4518: uint16(anon_sym_POUND_BQUOTE),
	4519: uint16(anon_sym_COMMA_AT),
	4520: uint16(anon_sym_POUND_COMMA_AT),
	4521: uint16(2),
	4522: uint16(751),
	4523: uint16(8),
	4524: uint16(sym_boolean),
	4525: uint16(anon_sym_POUND),
	4526: uint16(sym_number),
	4527: uint16(sym_symbol),
	4528: uint16(anon_sym_POUNDhash),
	4529: uint16(anon_sym_POUNDhasheq),
	4530: uint16(anon_sym_COMMA),
	4531: uint16(anon_sym_POUND_COMMA),
	4532: uint16(749),
	4533: uint16(25),
	4534: uint16(aux_sym__skip_token1),
	4535: uint16(aux_sym_comment_token1),
	4536: uint16(anon_sym_POUND_PIPE),
	4537: uint16(anon_sym_POUND_SEMI),
	4538: uint16(sym__line_comment),
	4539: uint16(anon_sym_POUND_LT_LT),
	4540: uint16(aux_sym_regex_token1),
	4541: uint16(anon_sym_DQUOTE),
	4542: uint16(sym_character),
	4543: uint16(sym_keyword),
	4544: uint16(anon_sym_POUND_AMP),
	4545: uint16(anon_sym_LPAREN),
	4546: uint16(anon_sym_LBRACK),
	4547: uint16(anon_sym_LBRACE),
	4548: uint16(anon_sym_POUNDfl),
	4549: uint16(anon_sym_POUNDfx),
	4550: uint16(anon_sym_POUNDs),
	4551: uint16(anon_sym_POUNDhashalw),
	4552: uint16(anon_sym_POUNDhasheqv),
	4553: uint16(anon_sym_SQUOTE),
	4554: uint16(anon_sym_BQUOTE),
	4555: uint16(anon_sym_POUND_SQUOTE),
	4556: uint16(anon_sym_POUND_BQUOTE),
	4557: uint16(anon_sym_COMMA_AT),
	4558: uint16(anon_sym_POUND_COMMA_AT),
	4559: uint16(2),
	4560: uint16(827),
	4561: uint16(8),
	4562: uint16(sym_boolean),
	4563: uint16(anon_sym_POUND),
	4564: uint16(sym_number),
	4565: uint16(sym_symbol),
	4566: uint16(anon_sym_POUNDhash),
	4567: uint16(anon_sym_POUNDhasheq),
	4568: uint16(anon_sym_COMMA),
	4569: uint16(anon_sym_POUND_COMMA),
	4570: uint16(825),
	4571: uint16(25),
	4572: uint16(aux_sym__skip_token1),
	4573: uint16(aux_sym_comment_token1),
	4574: uint16(anon_sym_POUND_PIPE),
	4575: uint16(anon_sym_POUND_SEMI),
	4576: uint16(sym__line_comment),
	4577: uint16(anon_sym_POUND_LT_LT),
	4578: uint16(aux_sym_regex_token1),
	4579: uint16(anon_sym_DQUOTE),
	4580: uint16(sym_character),
	4581: uint16(sym_keyword),
	4582: uint16(anon_sym_POUND_AMP),
	4583: uint16(anon_sym_LPAREN),
	4584: uint16(anon_sym_LBRACK),
	4585: uint16(anon_sym_LBRACE),
	4586: uint16(anon_sym_POUNDfl),
	4587: uint16(anon_sym_POUNDfx),
	4588: uint16(anon_sym_POUNDs),
	4589: uint16(anon_sym_POUNDhashalw),
	4590: uint16(anon_sym_POUNDhasheqv),
	4591: uint16(anon_sym_SQUOTE),
	4592: uint16(anon_sym_BQUOTE),
	4593: uint16(anon_sym_POUND_SQUOTE),
	4594: uint16(anon_sym_POUND_BQUOTE),
	4595: uint16(anon_sym_COMMA_AT),
	4596: uint16(anon_sym_POUND_COMMA_AT),
	4597: uint16(2),
	4598: uint16(759),
	4599: uint16(8),
	4600: uint16(sym_boolean),
	4601: uint16(anon_sym_POUND),
	4602: uint16(sym_number),
	4603: uint16(sym_symbol),
	4604: uint16(anon_sym_POUNDhash),
	4605: uint16(anon_sym_POUNDhasheq),
	4606: uint16(anon_sym_COMMA),
	4607: uint16(anon_sym_POUND_COMMA),
	4608: uint16(757),
	4609: uint16(25),
	4610: uint16(aux_sym__skip_token1),
	4611: uint16(aux_sym_comment_token1),
	4612: uint16(anon_sym_POUND_PIPE),
	4613: uint16(anon_sym_POUND_SEMI),
	4614: uint16(sym__line_comment),
	4615: uint16(anon_sym_POUND_LT_LT),
	4616: uint16(aux_sym_regex_token1),
	4617: uint16(anon_sym_DQUOTE),
	4618: uint16(sym_character),
	4619: uint16(sym_keyword),
	4620: uint16(anon_sym_POUND_AMP),
	4621: uint16(anon_sym_LPAREN),
	4622: uint16(anon_sym_LBRACK),
	4623: uint16(anon_sym_LBRACE),
	4624: uint16(anon_sym_POUNDfl),
	4625: uint16(anon_sym_POUNDfx),
	4626: uint16(anon_sym_POUNDs),
	4627: uint16(anon_sym_POUNDhashalw),
	4628: uint16(anon_sym_POUNDhasheqv),
	4629: uint16(anon_sym_SQUOTE),
	4630: uint16(anon_sym_BQUOTE),
	4631: uint16(anon_sym_POUND_SQUOTE),
	4632: uint16(anon_sym_POUND_BQUOTE),
	4633: uint16(anon_sym_COMMA_AT),
	4634: uint16(anon_sym_POUND_COMMA_AT),
	4635: uint16(2),
	4636: uint16(855),
	4637: uint16(8),
	4638: uint16(sym_boolean),
	4639: uint16(anon_sym_POUND),
	4640: uint16(sym_number),
	4641: uint16(sym_symbol),
	4642: uint16(anon_sym_POUNDhash),
	4643: uint16(anon_sym_POUNDhasheq),
	4644: uint16(anon_sym_COMMA),
	4645: uint16(anon_sym_POUND_COMMA),
	4646: uint16(853),
	4647: uint16(25),
	4648: uint16(aux_sym__skip_token1),
	4649: uint16(aux_sym_comment_token1),
	4650: uint16(anon_sym_POUND_PIPE),
	4651: uint16(anon_sym_POUND_SEMI),
	4652: uint16(sym__line_comment),
	4653: uint16(anon_sym_POUND_LT_LT),
	4654: uint16(aux_sym_regex_token1),
	4655: uint16(anon_sym_DQUOTE),
	4656: uint16(sym_character),
	4657: uint16(sym_keyword),
	4658: uint16(anon_sym_POUND_AMP),
	4659: uint16(anon_sym_LPAREN),
	4660: uint16(anon_sym_LBRACK),
	4661: uint16(anon_sym_LBRACE),
	4662: uint16(anon_sym_POUNDfl),
	4663: uint16(anon_sym_POUNDfx),
	4664: uint16(anon_sym_POUNDs),
	4665: uint16(anon_sym_POUNDhashalw),
	4666: uint16(anon_sym_POUNDhasheqv),
	4667: uint16(anon_sym_SQUOTE),
	4668: uint16(anon_sym_BQUOTE),
	4669: uint16(anon_sym_POUND_SQUOTE),
	4670: uint16(anon_sym_POUND_BQUOTE),
	4671: uint16(anon_sym_COMMA_AT),
	4672: uint16(anon_sym_POUND_COMMA_AT),
	4673: uint16(2),
	4674: uint16(775),
	4675: uint16(8),
	4676: uint16(sym_boolean),
	4677: uint16(anon_sym_POUND),
	4678: uint16(sym_number),
	4679: uint16(sym_symbol),
	4680: uint16(anon_sym_POUNDhash),
	4681: uint16(anon_sym_POUNDhasheq),
	4682: uint16(anon_sym_COMMA),
	4683: uint16(anon_sym_POUND_COMMA),
	4684: uint16(773),
	4685: uint16(25),
	4686: uint16(aux_sym__skip_token1),
	4687: uint16(aux_sym_comment_token1),
	4688: uint16(anon_sym_POUND_PIPE),
	4689: uint16(anon_sym_POUND_SEMI),
	4690: uint16(sym__line_comment),
	4691: uint16(anon_sym_POUND_LT_LT),
	4692: uint16(aux_sym_regex_token1),
	4693: uint16(anon_sym_DQUOTE),
	4694: uint16(sym_character),
	4695: uint16(sym_keyword),
	4696: uint16(anon_sym_POUND_AMP),
	4697: uint16(anon_sym_LPAREN),
	4698: uint16(anon_sym_LBRACK),
	4699: uint16(anon_sym_LBRACE),
	4700: uint16(anon_sym_POUNDfl),
	4701: uint16(anon_sym_POUNDfx),
	4702: uint16(anon_sym_POUNDs),
	4703: uint16(anon_sym_POUNDhashalw),
	4704: uint16(anon_sym_POUNDhasheqv),
	4705: uint16(anon_sym_SQUOTE),
	4706: uint16(anon_sym_BQUOTE),
	4707: uint16(anon_sym_POUND_SQUOTE),
	4708: uint16(anon_sym_POUND_BQUOTE),
	4709: uint16(anon_sym_COMMA_AT),
	4710: uint16(anon_sym_POUND_COMMA_AT),
	4711: uint16(2),
	4712: uint16(783),
	4713: uint16(8),
	4714: uint16(sym_boolean),
	4715: uint16(anon_sym_POUND),
	4716: uint16(sym_number),
	4717: uint16(sym_symbol),
	4718: uint16(anon_sym_POUNDhash),
	4719: uint16(anon_sym_POUNDhasheq),
	4720: uint16(anon_sym_COMMA),
	4721: uint16(anon_sym_POUND_COMMA),
	4722: uint16(781),
	4723: uint16(25),
	4724: uint16(aux_sym__skip_token1),
	4725: uint16(aux_sym_comment_token1),
	4726: uint16(anon_sym_POUND_PIPE),
	4727: uint16(anon_sym_POUND_SEMI),
	4728: uint16(sym__line_comment),
	4729: uint16(anon_sym_POUND_LT_LT),
	4730: uint16(aux_sym_regex_token1),
	4731: uint16(anon_sym_DQUOTE),
	4732: uint16(sym_character),
	4733: uint16(sym_keyword),
	4734: uint16(anon_sym_POUND_AMP),
	4735: uint16(anon_sym_LPAREN),
	4736: uint16(anon_sym_LBRACK),
	4737: uint16(anon_sym_LBRACE),
	4738: uint16(anon_sym_POUNDfl),
	4739: uint16(anon_sym_POUNDfx),
	4740: uint16(anon_sym_POUNDs),
	4741: uint16(anon_sym_POUNDhashalw),
	4742: uint16(anon_sym_POUNDhasheqv),
	4743: uint16(anon_sym_SQUOTE),
	4744: uint16(anon_sym_BQUOTE),
	4745: uint16(anon_sym_POUND_SQUOTE),
	4746: uint16(anon_sym_POUND_BQUOTE),
	4747: uint16(anon_sym_COMMA_AT),
	4748: uint16(anon_sym_POUND_COMMA_AT),
	4749: uint16(2),
	4750: uint16(791),
	4751: uint16(8),
	4752: uint16(sym_boolean),
	4753: uint16(anon_sym_POUND),
	4754: uint16(sym_number),
	4755: uint16(sym_symbol),
	4756: uint16(anon_sym_POUNDhash),
	4757: uint16(anon_sym_POUNDhasheq),
	4758: uint16(anon_sym_COMMA),
	4759: uint16(anon_sym_POUND_COMMA),
	4760: uint16(789),
	4761: uint16(25),
	4762: uint16(aux_sym__skip_token1),
	4763: uint16(aux_sym_comment_token1),
	4764: uint16(anon_sym_POUND_PIPE),
	4765: uint16(anon_sym_POUND_SEMI),
	4766: uint16(sym__line_comment),
	4767: uint16(anon_sym_POUND_LT_LT),
	4768: uint16(aux_sym_regex_token1),
	4769: uint16(anon_sym_DQUOTE),
	4770: uint16(sym_character),
	4771: uint16(sym_keyword),
	4772: uint16(anon_sym_POUND_AMP),
	4773: uint16(anon_sym_LPAREN),
	4774: uint16(anon_sym_LBRACK),
	4775: uint16(anon_sym_LBRACE),
	4776: uint16(anon_sym_POUNDfl),
	4777: uint16(anon_sym_POUNDfx),
	4778: uint16(anon_sym_POUNDs),
	4779: uint16(anon_sym_POUNDhashalw),
	4780: uint16(anon_sym_POUNDhasheqv),
	4781: uint16(anon_sym_SQUOTE),
	4782: uint16(anon_sym_BQUOTE),
	4783: uint16(anon_sym_POUND_SQUOTE),
	4784: uint16(anon_sym_POUND_BQUOTE),
	4785: uint16(anon_sym_COMMA_AT),
	4786: uint16(anon_sym_POUND_COMMA_AT),
	4787: uint16(2),
	4788: uint16(803),
	4789: uint16(8),
	4790: uint16(sym_boolean),
	4791: uint16(anon_sym_POUND),
	4792: uint16(sym_number),
	4793: uint16(sym_symbol),
	4794: uint16(anon_sym_POUNDhash),
	4795: uint16(anon_sym_POUNDhasheq),
	4796: uint16(anon_sym_COMMA),
	4797: uint16(anon_sym_POUND_COMMA),
	4798: uint16(801),
	4799: uint16(25),
	4800: uint16(aux_sym__skip_token1),
	4801: uint16(aux_sym_comment_token1),
	4802: uint16(anon_sym_POUND_PIPE),
	4803: uint16(anon_sym_POUND_SEMI),
	4804: uint16(sym__line_comment),
	4805: uint16(anon_sym_POUND_LT_LT),
	4806: uint16(aux_sym_regex_token1),
	4807: uint16(anon_sym_DQUOTE),
	4808: uint16(sym_character),
	4809: uint16(sym_keyword),
	4810: uint16(anon_sym_POUND_AMP),
	4811: uint16(anon_sym_LPAREN),
	4812: uint16(anon_sym_LBRACK),
	4813: uint16(anon_sym_LBRACE),
	4814: uint16(anon_sym_POUNDfl),
	4815: uint16(anon_sym_POUNDfx),
	4816: uint16(anon_sym_POUNDs),
	4817: uint16(anon_sym_POUNDhashalw),
	4818: uint16(anon_sym_POUNDhasheqv),
	4819: uint16(anon_sym_SQUOTE),
	4820: uint16(anon_sym_BQUOTE),
	4821: uint16(anon_sym_POUND_SQUOTE),
	4822: uint16(anon_sym_POUND_BQUOTE),
	4823: uint16(anon_sym_COMMA_AT),
	4824: uint16(anon_sym_POUND_COMMA_AT),
	4825: uint16(2),
	4826: uint16(843),
	4827: uint16(8),
	4828: uint16(sym_boolean),
	4829: uint16(anon_sym_POUND),
	4830: uint16(sym_number),
	4831: uint16(sym_symbol),
	4832: uint16(anon_sym_POUNDhash),
	4833: uint16(anon_sym_POUNDhasheq),
	4834: uint16(anon_sym_COMMA),
	4835: uint16(anon_sym_POUND_COMMA),
	4836: uint16(841),
	4837: uint16(25),
	4838: uint16(aux_sym__skip_token1),
	4839: uint16(aux_sym_comment_token1),
	4840: uint16(anon_sym_POUND_PIPE),
	4841: uint16(anon_sym_POUND_SEMI),
	4842: uint16(sym__line_comment),
	4843: uint16(anon_sym_POUND_LT_LT),
	4844: uint16(aux_sym_regex_token1),
	4845: uint16(anon_sym_DQUOTE),
	4846: uint16(sym_character),
	4847: uint16(sym_keyword),
	4848: uint16(anon_sym_POUND_AMP),
	4849: uint16(anon_sym_LPAREN),
	4850: uint16(anon_sym_LBRACK),
	4851: uint16(anon_sym_LBRACE),
	4852: uint16(anon_sym_POUNDfl),
	4853: uint16(anon_sym_POUNDfx),
	4854: uint16(anon_sym_POUNDs),
	4855: uint16(anon_sym_POUNDhashalw),
	4856: uint16(anon_sym_POUNDhasheqv),
	4857: uint16(anon_sym_SQUOTE),
	4858: uint16(anon_sym_BQUOTE),
	4859: uint16(anon_sym_POUND_SQUOTE),
	4860: uint16(anon_sym_POUND_BQUOTE),
	4861: uint16(anon_sym_COMMA_AT),
	4862: uint16(anon_sym_POUND_COMMA_AT),
	4863: uint16(2),
	4864: uint16(811),
	4865: uint16(8),
	4866: uint16(sym_boolean),
	4867: uint16(anon_sym_POUND),
	4868: uint16(sym_number),
	4869: uint16(sym_symbol),
	4870: uint16(anon_sym_POUNDhash),
	4871: uint16(anon_sym_POUNDhasheq),
	4872: uint16(anon_sym_COMMA),
	4873: uint16(anon_sym_POUND_COMMA),
	4874: uint16(809),
	4875: uint16(25),
	4876: uint16(aux_sym__skip_token1),
	4877: uint16(aux_sym_comment_token1),
	4878: uint16(anon_sym_POUND_PIPE),
	4879: uint16(anon_sym_POUND_SEMI),
	4880: uint16(sym__line_comment),
	4881: uint16(anon_sym_POUND_LT_LT),
	4882: uint16(aux_sym_regex_token1),
	4883: uint16(anon_sym_DQUOTE),
	4884: uint16(sym_character),
	4885: uint16(sym_keyword),
	4886: uint16(anon_sym_POUND_AMP),
	4887: uint16(anon_sym_LPAREN),
	4888: uint16(anon_sym_LBRACK),
	4889: uint16(anon_sym_LBRACE),
	4890: uint16(anon_sym_POUNDfl),
	4891: uint16(anon_sym_POUNDfx),
	4892: uint16(anon_sym_POUNDs),
	4893: uint16(anon_sym_POUNDhashalw),
	4894: uint16(anon_sym_POUNDhasheqv),
	4895: uint16(anon_sym_SQUOTE),
	4896: uint16(anon_sym_BQUOTE),
	4897: uint16(anon_sym_POUND_SQUOTE),
	4898: uint16(anon_sym_POUND_BQUOTE),
	4899: uint16(anon_sym_COMMA_AT),
	4900: uint16(anon_sym_POUND_COMMA_AT),
	4901: uint16(2),
	4902: uint16(899),
	4903: uint16(8),
	4904: uint16(sym_boolean),
	4905: uint16(anon_sym_POUND),
	4906: uint16(sym_number),
	4907: uint16(sym_symbol),
	4908: uint16(anon_sym_POUNDhash),
	4909: uint16(anon_sym_POUNDhasheq),
	4910: uint16(anon_sym_COMMA),
	4911: uint16(anon_sym_POUND_COMMA),
	4912: uint16(897),
	4913: uint16(25),
	4914: uint16(aux_sym__skip_token1),
	4915: uint16(aux_sym_comment_token1),
	4916: uint16(anon_sym_POUND_PIPE),
	4917: uint16(anon_sym_POUND_SEMI),
	4918: uint16(sym__line_comment),
	4919: uint16(anon_sym_POUND_LT_LT),
	4920: uint16(aux_sym_regex_token1),
	4921: uint16(anon_sym_DQUOTE),
	4922: uint16(sym_character),
	4923: uint16(sym_keyword),
	4924: uint16(anon_sym_POUND_AMP),
	4925: uint16(anon_sym_LPAREN),
	4926: uint16(anon_sym_LBRACK),
	4927: uint16(anon_sym_LBRACE),
	4928: uint16(anon_sym_POUNDfl),
	4929: uint16(anon_sym_POUNDfx),
	4930: uint16(anon_sym_POUNDs),
	4931: uint16(anon_sym_POUNDhashalw),
	4932: uint16(anon_sym_POUNDhasheqv),
	4933: uint16(anon_sym_SQUOTE),
	4934: uint16(anon_sym_BQUOTE),
	4935: uint16(anon_sym_POUND_SQUOTE),
	4936: uint16(anon_sym_POUND_BQUOTE),
	4937: uint16(anon_sym_COMMA_AT),
	4938: uint16(anon_sym_POUND_COMMA_AT),
	4939: uint16(2),
	4940: uint16(883),
	4941: uint16(8),
	4942: uint16(sym_boolean),
	4943: uint16(anon_sym_POUND),
	4944: uint16(sym_number),
	4945: uint16(sym_symbol),
	4946: uint16(anon_sym_POUNDhash),
	4947: uint16(anon_sym_POUNDhasheq),
	4948: uint16(anon_sym_COMMA),
	4949: uint16(anon_sym_POUND_COMMA),
	4950: uint16(881),
	4951: uint16(25),
	4952: uint16(aux_sym__skip_token1),
	4953: uint16(aux_sym_comment_token1),
	4954: uint16(anon_sym_POUND_PIPE),
	4955: uint16(anon_sym_POUND_SEMI),
	4956: uint16(sym__line_comment),
	4957: uint16(anon_sym_POUND_LT_LT),
	4958: uint16(aux_sym_regex_token1),
	4959: uint16(anon_sym_DQUOTE),
	4960: uint16(sym_character),
	4961: uint16(sym_keyword),
	4962: uint16(anon_sym_POUND_AMP),
	4963: uint16(anon_sym_LPAREN),
	4964: uint16(anon_sym_LBRACK),
	4965: uint16(anon_sym_LBRACE),
	4966: uint16(anon_sym_POUNDfl),
	4967: uint16(anon_sym_POUNDfx),
	4968: uint16(anon_sym_POUNDs),
	4969: uint16(anon_sym_POUNDhashalw),
	4970: uint16(anon_sym_POUNDhasheqv),
	4971: uint16(anon_sym_SQUOTE),
	4972: uint16(anon_sym_BQUOTE),
	4973: uint16(anon_sym_POUND_SQUOTE),
	4974: uint16(anon_sym_POUND_BQUOTE),
	4975: uint16(anon_sym_COMMA_AT),
	4976: uint16(anon_sym_POUND_COMMA_AT),
	4977: uint16(7),
	4978: uint16(491),
	4979: uint16(1),
	4980: uint16(anon_sym_DQUOTE),
	4981: uint16(497),
	4982: uint16(1),
	4983: uint16(anon_sym_LPAREN),
	4984: uint16(499),
	4985: uint16(1),
	4986: uint16(anon_sym_LBRACK),
	4987: uint16(501),
	4988: uint16(1),
	4989: uint16(anon_sym_LBRACE),
	4990: uint16(917),
	4991: uint16(1),
	4992: uint16(sym_decimal),
	4993: uint16(204),
	4994: uint16(1),
	4995: uint16(sym__real_string),
	4996: uint16(205),
	4997: uint16(1),
	4998: uint16(sym_list),
	4999: uint16(7),
	5000: uint16(21),
	5001: uint16(1),
	5002: uint16(anon_sym_DQUOTE),
	5003: uint16(25),
	5004: uint16(1),
	5005: uint16(anon_sym_LPAREN),
	5006: uint16(27),
	5007: uint16(1),
	5008: uint16(anon_sym_LBRACK),
	5009: uint16(29),
	5010: uint16(1),
	5011: uint16(anon_sym_LBRACE),
	5012: uint16(919),
	5013: uint16(1),
	5014: uint16(sym_decimal),
	5015: uint16(134),
	5016: uint16(1),
	5017: uint16(sym__real_string),
	5018: uint16(158),
	5019: uint16(1),
	5020: uint16(sym_list),
	5021: uint16(7),
	5022: uint16(163),
	5023: uint16(1),
	5024: uint16(anon_sym_DQUOTE),
	5025: uint16(167),
	5026: uint16(1),
	5027: uint16(anon_sym_LPAREN),
	5028: uint16(171),
	5029: uint16(1),
	5030: uint16(anon_sym_LBRACK),
	5031: uint16(173),
	5032: uint16(1),
	5033: uint16(anon_sym_LBRACE),
	5034: uint16(921),
	5035: uint16(1),
	5036: uint16(sym_decimal),
	5037: uint16(101),
	5038: uint16(1),
	5039: uint16(sym__real_string),
	5040: uint16(103),
	5041: uint16(1),
	5042: uint16(sym_list),
	5043: uint16(6),
	5044: uint16(497),
	5045: uint16(1),
	5046: uint16(anon_sym_LPAREN),
	5047: uint16(499),
	5048: uint16(1),
	5049: uint16(anon_sym_LBRACK),
	5050: uint16(501),
	5051: uint16(1),
	5052: uint16(anon_sym_LBRACE),
	5053: uint16(923),
	5054: uint16(1),
	5055: uint16(anon_sym_POUND),
	5056: uint16(925),
	5057: uint16(1),
	5058: uint16(anon_sym_EQ),
	5059: uint16(174),
	5060: uint16(1),
	5061: uint16(sym_list),
	5062: uint16(6),
	5063: uint16(25),
	5064: uint16(1),
	5065: uint16(anon_sym_LPAREN),
	5066: uint16(27),
	5067: uint16(1),
	5068: uint16(anon_sym_LBRACK),
	5069: uint16(29),
	5070: uint16(1),
	5071: uint16(anon_sym_LBRACE),
	5072: uint16(927),
	5073: uint16(1),
	5074: uint16(anon_sym_POUND),
	5075: uint16(929),
	5076: uint16(1),
	5077: uint16(anon_sym_EQ),
	5078: uint16(167),
	5079: uint16(1),
	5080: uint16(sym_list),
	5081: uint16(6),
	5082: uint16(167),
	5083: uint16(1),
	5084: uint16(anon_sym_LPAREN),
	5085: uint16(171),
	5086: uint16(1),
	5087: uint16(anon_sym_LBRACK),
	5088: uint16(173),
	5089: uint16(1),
	5090: uint16(anon_sym_LBRACE),
	5091: uint16(931),
	5092: uint16(1),
	5093: uint16(anon_sym_POUND),
	5094: uint16(933),
	5095: uint16(1),
	5096: uint16(anon_sym_EQ),
	5097: uint16(120),
	5098: uint16(1),
	5099: uint16(sym_list),
	5100: uint16(4),
	5101: uint16(935),
	5102: uint16(1),
	5103: uint16(anon_sym_POUND_PIPE),
	5104: uint16(937),
	5105: uint16(1),
	5106: uint16(aux_sym_block_comment_token1),
	5107: uint16(939),
	5108: uint16(1),
	5109: uint16(anon_sym_PIPE_POUND),
	5110: uint16(222),
	5111: uint16(2),
	5112: uint16(sym_block_comment),
	5113: uint16(aux_sym_block_comment_repeat1),
	5114: uint16(4),
	5115: uint16(935),
	5116: uint16(1),
	5117: uint16(anon_sym_POUND_PIPE),
	5118: uint16(941),
	5119: uint16(1),
	5120: uint16(aux_sym_block_comment_token1),
	5121: uint16(943),
	5122: uint16(1),
	5123: uint16(anon_sym_PIPE_POUND),
	5124: uint16(223),
	5125: uint16(2),
	5126: uint16(sym_block_comment),
	5127: uint16(aux_sym_block_comment_repeat1),
	5128: uint16(4),
	5129: uint16(935),
	5130: uint16(1),
	5131: uint16(anon_sym_POUND_PIPE),
	5132: uint16(945),
	5133: uint16(1),
	5134: uint16(aux_sym_block_comment_token1),
	5135: uint16(947),
	5136: uint16(1),
	5137: uint16(anon_sym_PIPE_POUND),
	5138: uint16(226),
	5139: uint16(2),
	5140: uint16(sym_block_comment),
	5141: uint16(aux_sym_block_comment_repeat1),
	5142: uint16(5),
	5143: uint16(167),
	5144: uint16(1),
	5145: uint16(anon_sym_LPAREN),
	5146: uint16(171),
	5147: uint16(1),
	5148: uint16(anon_sym_LBRACK),
	5149: uint16(173),
	5150: uint16(1),
	5151: uint16(anon_sym_LBRACE),
	5152: uint16(949),
	5153: uint16(1),
	5154: uint16(sym_decimal),
	5155: uint16(103),
	5156: uint16(1),
	5157: uint16(sym_list),
	5158: uint16(4),
	5159: uint16(935),
	5160: uint16(1),
	5161: uint16(anon_sym_POUND_PIPE),
	5162: uint16(941),
	5163: uint16(1),
	5164: uint16(aux_sym_block_comment_token1),
	5165: uint16(951),
	5166: uint16(1),
	5167: uint16(anon_sym_PIPE_POUND),
	5168: uint16(223),
	5169: uint16(2),
	5170: uint16(sym_block_comment),
	5171: uint16(aux_sym_block_comment_repeat1),
	5172: uint16(4),
	5173: uint16(953),
	5174: uint16(1),
	5175: uint16(anon_sym_POUND_PIPE),
	5176: uint16(956),
	5177: uint16(1),
	5178: uint16(aux_sym_block_comment_token1),
	5179: uint16(959),
	5180: uint16(1),
	5181: uint16(anon_sym_PIPE_POUND),
	5182: uint16(223),
	5183: uint16(2),
	5184: uint16(sym_block_comment),
	5185: uint16(aux_sym_block_comment_repeat1),
	5186: uint16(5),
	5187: uint16(25),
	5188: uint16(1),
	5189: uint16(anon_sym_LPAREN),
	5190: uint16(27),
	5191: uint16(1),
	5192: uint16(anon_sym_LBRACK),
	5193: uint16(29),
	5194: uint16(1),
	5195: uint16(anon_sym_LBRACE),
	5196: uint16(961),
	5197: uint16(1),
	5198: uint16(sym_decimal),
	5199: uint16(158),
	5200: uint16(1),
	5201: uint16(sym_list),
	5202: uint16(4),
	5203: uint16(935),
	5204: uint16(1),
	5205: uint16(anon_sym_POUND_PIPE),
	5206: uint16(963),
	5207: uint16(1),
	5208: uint16(aux_sym_block_comment_token1),
	5209: uint16(965),
	5210: uint16(1),
	5211: uint16(anon_sym_PIPE_POUND),
	5212: uint16(228),
	5213: uint16(2),
	5214: uint16(sym_block_comment),
	5215: uint16(aux_sym_block_comment_repeat1),
	5216: uint16(4),
	5217: uint16(935),
	5218: uint16(1),
	5219: uint16(anon_sym_POUND_PIPE),
	5220: uint16(941),
	5221: uint16(1),
	5222: uint16(aux_sym_block_comment_token1),
	5223: uint16(967),
	5224: uint16(1),
	5225: uint16(anon_sym_PIPE_POUND),
	5226: uint16(223),
	5227: uint16(2),
	5228: uint16(sym_block_comment),
	5229: uint16(aux_sym_block_comment_repeat1),
	5230: uint16(5),
	5231: uint16(497),
	5232: uint16(1),
	5233: uint16(anon_sym_LPAREN),
	5234: uint16(499),
	5235: uint16(1),
	5236: uint16(anon_sym_LBRACK),
	5237: uint16(501),
	5238: uint16(1),
	5239: uint16(anon_sym_LBRACE),
	5240: uint16(969),
	5241: uint16(1),
	5242: uint16(sym_decimal),
	5243: uint16(205),
	5244: uint16(1),
	5245: uint16(sym_list),
	5246: uint16(4),
	5247: uint16(935),
	5248: uint16(1),
	5249: uint16(anon_sym_POUND_PIPE),
	5250: uint16(941),
	5251: uint16(1),
	5252: uint16(aux_sym_block_comment_token1),
	5253: uint16(971),
	5254: uint16(1),
	5255: uint16(anon_sym_PIPE_POUND),
	5256: uint16(223),
	5257: uint16(2),
	5258: uint16(sym_block_comment),
	5259: uint16(aux_sym_block_comment_repeat1),
	5260: uint16(4),
	5261: uint16(935),
	5262: uint16(1),
	5263: uint16(anon_sym_POUND_PIPE),
	5264: uint16(973),
	5265: uint16(1),
	5266: uint16(aux_sym_block_comment_token1),
	5267: uint16(975),
	5268: uint16(1),
	5269: uint16(anon_sym_PIPE_POUND),
	5270: uint16(219),
	5271: uint16(2),
	5272: uint16(sym_block_comment),
	5273: uint16(aux_sym_block_comment_repeat1),
	5274: uint16(4),
	5275: uint16(167),
	5276: uint16(1),
	5277: uint16(anon_sym_LPAREN),
	5278: uint16(171),
	5279: uint16(1),
	5280: uint16(anon_sym_LBRACK),
	5281: uint16(173),
	5282: uint16(1),
	5283: uint16(anon_sym_LBRACE),
	5284: uint16(120),
	5285: uint16(1),
	5286: uint16(sym_list),
	5287: uint16(3),
	5288: uint16(977),
	5289: uint16(1),
	5290: uint16(anon_sym_DQUOTE),
	5291: uint16(237),
	5292: uint16(1),
	5293: uint16(aux_sym__real_string_repeat1),
	5294: uint16(979),
	5295: uint16(2),
	5296: uint16(aux_sym__real_string_token1),
	5297: uint16(sym_escape_sequence),
	5298: uint16(3),
	5299: uint16(981),
	5300: uint16(1),
	5301: uint16(anon_sym_DQUOTE),
	5302: uint16(244),
	5303: uint16(1),
	5304: uint16(aux_sym__real_string_repeat1),
	5305: uint16(983),
	5306: uint16(2),
	5307: uint16(aux_sym__real_string_token1),
	5308: uint16(sym_escape_sequence),
	5309: uint16(4),
	5310: uint16(25),
	5311: uint16(1),
	5312: uint16(anon_sym_LPAREN),
	5313: uint16(27),
	5314: uint16(1),
	5315: uint16(anon_sym_LBRACK),
	5316: uint16(29),
	5317: uint16(1),
	5318: uint16(anon_sym_LBRACE),
	5319: uint16(146),
	5320: uint16(1),
	5321: uint16(sym_list),
	5322: uint16(4),
	5323: uint16(25),
	5324: uint16(1),
	5325: uint16(anon_sym_LPAREN),
	5326: uint16(27),
	5327: uint16(1),
	5328: uint16(anon_sym_LBRACK),
	5329: uint16(29),
	5330: uint16(1),
	5331: uint16(anon_sym_LBRACE),
	5332: uint16(147),
	5333: uint16(1),
	5334: uint16(sym_list),
	5335: uint16(4),
	5336: uint16(167),
	5337: uint16(1),
	5338: uint16(anon_sym_LPAREN),
	5339: uint16(171),
	5340: uint16(1),
	5341: uint16(anon_sym_LBRACK),
	5342: uint16(173),
	5343: uint16(1),
	5344: uint16(anon_sym_LBRACE),
	5345: uint16(96),
	5346: uint16(1),
	5347: uint16(sym_list),
	5348: uint16(4),
	5349: uint16(167),
	5350: uint16(1),
	5351: uint16(anon_sym_LPAREN),
	5352: uint16(171),
	5353: uint16(1),
	5354: uint16(anon_sym_LBRACK),
	5355: uint16(173),
	5356: uint16(1),
	5357: uint16(anon_sym_LBRACE),
	5358: uint16(119),
	5359: uint16(1),
	5360: uint16(sym_list),
	5361: uint16(3),
	5362: uint16(985),
	5363: uint16(1),
	5364: uint16(anon_sym_DQUOTE),
	5365: uint16(237),
	5366: uint16(1),
	5367: uint16(aux_sym__real_string_repeat1),
	5368: uint16(987),
	5369: uint16(2),
	5370: uint16(aux_sym__real_string_token1),
	5371: uint16(sym_escape_sequence),
	5372: uint16(3),
	5373: uint16(990),
	5374: uint16(1),
	5375: uint16(anon_sym_DQUOTE),
	5376: uint16(237),
	5377: uint16(1),
	5378: uint16(aux_sym__real_string_repeat1),
	5379: uint16(979),
	5380: uint16(2),
	5381: uint16(aux_sym__real_string_token1),
	5382: uint16(sym_escape_sequence),
	5383: uint16(4),
	5384: uint16(497),
	5385: uint16(1),
	5386: uint16(anon_sym_LPAREN),
	5387: uint16(499),
	5388: uint16(1),
	5389: uint16(anon_sym_LBRACK),
	5390: uint16(501),
	5391: uint16(1),
	5392: uint16(anon_sym_LBRACE),
	5393: uint16(174),
	5394: uint16(1),
	5395: uint16(sym_list),
	5396: uint16(4),
	5397: uint16(497),
	5398: uint16(1),
	5399: uint16(anon_sym_LPAREN),
	5400: uint16(499),
	5401: uint16(1),
	5402: uint16(anon_sym_LBRACK),
	5403: uint16(501),
	5404: uint16(1),
	5405: uint16(anon_sym_LBRACE),
	5406: uint16(175),
	5407: uint16(1),
	5408: uint16(sym_list),
	5409: uint16(4),
	5410: uint16(497),
	5411: uint16(1),
	5412: uint16(anon_sym_LPAREN),
	5413: uint16(499),
	5414: uint16(1),
	5415: uint16(anon_sym_LBRACK),
	5416: uint16(501),
	5417: uint16(1),
	5418: uint16(anon_sym_LBRACE),
	5419: uint16(176),
	5420: uint16(1),
	5421: uint16(sym_list),
	5422: uint16(3),
	5423: uint16(992),
	5424: uint16(1),
	5425: uint16(anon_sym_DQUOTE),
	5426: uint16(238),
	5427: uint16(1),
	5428: uint16(aux_sym__real_string_repeat1),
	5429: uint16(994),
	5430: uint16(2),
	5431: uint16(aux_sym__real_string_token1),
	5432: uint16(sym_escape_sequence),
	5433: uint16(4),
	5434: uint16(25),
	5435: uint16(1),
	5436: uint16(anon_sym_LPAREN),
	5437: uint16(27),
	5438: uint16(1),
	5439: uint16(anon_sym_LBRACK),
	5440: uint16(29),
	5441: uint16(1),
	5442: uint16(anon_sym_LBRACE),
	5443: uint16(167),
	5444: uint16(1),
	5445: uint16(sym_list),
	5446: uint16(3),
	5447: uint16(996),
	5448: uint16(1),
	5449: uint16(anon_sym_DQUOTE),
	5450: uint16(237),
	5451: uint16(1),
	5452: uint16(aux_sym__real_string_repeat1),
	5453: uint16(979),
	5454: uint16(2),
	5455: uint16(aux_sym__real_string_token1),
	5456: uint16(sym_escape_sequence),
	5457: uint16(3),
	5458: uint16(998),
	5459: uint16(1),
	5460: uint16(anon_sym_DQUOTE),
	5461: uint16(231),
	5462: uint16(1),
	5463: uint16(aux_sym__real_string_repeat1),
	5464: uint16(1000),
	5465: uint16(2),
	5466: uint16(aux_sym__real_string_token1),
	5467: uint16(sym_escape_sequence),
	5468: uint16(2),
	5469: uint16(855),
	5470: uint16(1),
	5471: uint16(aux_sym_block_comment_token1),
	5472: uint16(853),
	5473: uint16(2),
	5474: uint16(anon_sym_POUND_PIPE),
	5475: uint16(anon_sym_PIPE_POUND),
	5476: uint16(2),
	5477: uint16(859),
	5478: uint16(1),
	5479: uint16(aux_sym_block_comment_token1),
	5480: uint16(857),
	5481: uint16(2),
	5482: uint16(anon_sym_POUND_PIPE),
	5483: uint16(anon_sym_PIPE_POUND),
	5484: uint16(2),
	5485: uint16(491),
	5486: uint16(1),
	5487: uint16(anon_sym_DQUOTE),
	5488: uint16(207),
	5489: uint16(1),
	5490: uint16(sym__real_string),
	5491: uint16(2),
	5492: uint16(163),
	5493: uint16(1),
	5494: uint16(anon_sym_DQUOTE),
	5495: uint16(108),
	5496: uint16(1),
	5497: uint16(sym__real_string),
	5498: uint16(2),
	5499: uint16(21),
	5500: uint16(1),
	5501: uint16(anon_sym_DQUOTE),
	5502: uint16(161),
	5503: uint16(1),
	5504: uint16(sym__real_string),
	5505: uint16(1),
	5506: uint16(1002),
	5507: uint16(1),
	5508: uint16(sym__here_string_body),
	5509: uint16(1),
	5510: uint16(559),
	5511: uint16(1),
	5512: uint16(sym_lang_name),
	5513: uint16(1),
	5514: uint16(1004),
	5515: uint16(1),
	5516: uint16(sym__here_string_body),
	5517: uint16(1),
	5518: uint16(1006),
	5519: uint16(1),
	5521: uint16(1),
	5522: uint16(481),
	5523: uint16(1),
	5524: uint16(sym_lang_name),
	5525: uint16(1),
	5526: uint16(1008),
	5527: uint16(1),
	5528: uint16(sym__here_string_body),
}

var ts_small_parse_table_map = [164]uint32_t{
	1:   uint32(45),
	2:   uint32(90),
	3:   uint32(135),
	4:   uint32(180),
	5:   uint32(225),
	6:   uint32(270),
	7:   uint32(315),
	8:   uint32(360),
	9:   uint32(405),
	10:  uint32(450),
	11:  uint32(495),
	12:  uint32(540),
	13:  uint32(585),
	14:  uint32(630),
	15:  uint32(675),
	16:  uint32(720),
	17:  uint32(765),
	18:  uint32(810),
	19:  uint32(855),
	20:  uint32(900),
	21:  uint32(945),
	22:  uint32(990),
	23:  uint32(1035),
	24:  uint32(1080),
	25:  uint32(1125),
	26:  uint32(1170),
	27:  uint32(1215),
	28:  uint32(1260),
	29:  uint32(1305),
	30:  uint32(1350),
	31:  uint32(1395),
	32:  uint32(1440),
	33:  uint32(1485),
	34:  uint32(1530),
	35:  uint32(1575),
	36:  uint32(1620),
	37:  uint32(1665),
	38:  uint32(1710),
	39:  uint32(1755),
	40:  uint32(1800),
	41:  uint32(1853),
	42:  uint32(1895),
	43:  uint32(1937),
	44:  uint32(1979),
	45:  uint32(2021),
	46:  uint32(2063),
	47:  uint32(2105),
	48:  uint32(2147),
	49:  uint32(2189),
	50:  uint32(2231),
	51:  uint32(2273),
	52:  uint32(2315),
	53:  uint32(2357),
	54:  uint32(2399),
	55:  uint32(2441),
	56:  uint32(2483),
	57:  uint32(2525),
	58:  uint32(2567),
	59:  uint32(2609),
	60:  uint32(2651),
	61:  uint32(2693),
	62:  uint32(2735),
	63:  uint32(2777),
	64:  uint32(2819),
	65:  uint32(2861),
	66:  uint32(2903),
	67:  uint32(2945),
	68:  uint32(2987),
	69:  uint32(3029),
	70:  uint32(3071),
	71:  uint32(3113),
	72:  uint32(3155),
	73:  uint32(3197),
	74:  uint32(3239),
	75:  uint32(3281),
	76:  uint32(3323),
	77:  uint32(3365),
	78:  uint32(3407),
	79:  uint32(3449),
	80:  uint32(3491),
	81:  uint32(3533),
	82:  uint32(3571),
	83:  uint32(3609),
	84:  uint32(3647),
	85:  uint32(3685),
	86:  uint32(3723),
	87:  uint32(3761),
	88:  uint32(3799),
	89:  uint32(3837),
	90:  uint32(3875),
	91:  uint32(3913),
	92:  uint32(3951),
	93:  uint32(3989),
	94:  uint32(4027),
	95:  uint32(4065),
	96:  uint32(4103),
	97:  uint32(4141),
	98:  uint32(4179),
	99:  uint32(4217),
	100: uint32(4255),
	101: uint32(4293),
	102: uint32(4331),
	103: uint32(4369),
	104: uint32(4407),
	105: uint32(4445),
	106: uint32(4483),
	107: uint32(4521),
	108: uint32(4559),
	109: uint32(4597),
	110: uint32(4635),
	111: uint32(4673),
	112: uint32(4711),
	113: uint32(4749),
	114: uint32(4787),
	115: uint32(4825),
	116: uint32(4863),
	117: uint32(4901),
	118: uint32(4939),
	119: uint32(4977),
	120: uint32(4999),
	121: uint32(5021),
	122: uint32(5043),
	123: uint32(5062),
	124: uint32(5081),
	125: uint32(5100),
	126: uint32(5114),
	127: uint32(5128),
	128: uint32(5142),
	129: uint32(5158),
	130: uint32(5172),
	131: uint32(5186),
	132: uint32(5202),
	133: uint32(5216),
	134: uint32(5230),
	135: uint32(5246),
	136: uint32(5260),
	137: uint32(5274),
	138: uint32(5287),
	139: uint32(5298),
	140: uint32(5309),
	141: uint32(5322),
	142: uint32(5335),
	143: uint32(5348),
	144: uint32(5361),
	145: uint32(5372),
	146: uint32(5383),
	147: uint32(5396),
	148: uint32(5409),
	149: uint32(5422),
	150: uint32(5433),
	151: uint32(5446),
	152: uint32(5457),
	153: uint32(5468),
	154: uint32(5476),
	155: uint32(5484),
	156: uint32(5491),
	157: uint32(5498),
	158: uint32(5505),
	159: uint32(5509),
	160: uint32(5513),
	161: uint32(5517),
	162: uint32(5521),
	163: uint32(5525),
}

var ts_parse_actions = [1010]TSParseActionEntry{
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token: uint8(TSParseActionTypeReduce),
		Fsymbol:     uint16(sym_program),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(166)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(220)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(213)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(256)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(250)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(245)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(224)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(233)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(234)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	38: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(234)),
	}})))),
	39: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	40: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	41: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	44: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	45: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	46: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(89)),
	}})))),
	47: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Fcount: uint8(1),
	}})),
	52: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	56: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	57: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	58: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(252)),
	}})))),
	59: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(252)),
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
		Fsymbol:      uint16(aux_sym_list_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_list_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	69: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(98)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	70: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	71: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	72: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(229)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	75: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	76: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	77: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	78: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(214)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	81: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(251)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	82: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	83: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	84: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(249)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	86: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	87: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(242)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	88: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	89: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	90: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(39)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	93: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(9)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	94: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	95: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	96: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	97: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	98: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(10)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	101: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(11)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	102: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	103: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	104: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(221)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	107: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(236)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	108: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	109: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	110: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(235)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	113: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(235)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	114: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	115: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	116: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(40)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	119: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(41)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	120: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	121: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	122: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(42)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	125: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(43)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	126: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	127: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	128: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(44)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	129: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	130: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(45)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	133: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(46)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	135: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	136: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(47)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	139: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	140: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(48)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	141: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	142: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(255)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	145: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_repeat1),
	})))),
	146: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(255)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	147: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(98)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(229)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(214)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(251)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(249)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(242)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(145)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(221)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(236)),
	}})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(235)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(235)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(41)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Fcount: uint8(1),
	}})),
	196: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(255)),
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
		Fcount: uint8(1),
	}})),
	204: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(255)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
	}})))),
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
		Fcount: uint8(1),
	}})),
	212: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(170)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(116)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(125)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
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
	240: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(194)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(191)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_program),
	})))),
	253: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	254: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	255: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	256: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	257: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	258: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	259: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	260: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	261: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(22)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	262: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	263: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	264: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(166)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	265: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	266: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	267: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(220)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	268: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	269: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	270: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(84)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	271: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	272: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	273: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(22)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	274: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	275: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	276: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(213)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	277: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	278: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	279: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(256)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	280: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	281: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	282: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(250)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	283: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	284: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	285: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(245)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	286: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	287: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	288: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(37)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	289: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	290: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	291: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(3)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	292: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	293: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	294: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(4)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	295: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	296: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	297: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	298: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	299: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	300: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(224)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	301: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	302: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	303: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(233)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	304: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	305: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	306: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(234)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	307: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	308: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	309: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(234)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	310: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	311: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	312: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(86)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	313: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	314: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	315: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(87)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	316: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	317: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	318: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(88)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	319: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	320: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	321: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	322: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	323: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	324: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(90)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	325: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	326: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	327: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(91)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	328: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	329: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	330: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(24)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	331: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	332: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	333: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(25)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	334: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	335: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	336: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	337: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	338: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	339: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(252)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	340: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	341: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	342: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(252)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	343: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	344: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(133)),
	}})))),
	345: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	346: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(184)),
	}})))),
	347: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	348: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(225)),
	}})))),
	349: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	350: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
	}})))),
	351: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	352: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(169)),
	}})))),
	353: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	354: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(169)),
	}})))),
	355: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	356: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	357: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	358: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(156)),
	}})))),
	359: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	360: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(156)),
	}})))),
	361: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	362: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
	}})))),
	363: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	364: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(157)),
	}})))),
	365: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	366: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(157)),
	}})))),
	367: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	368: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(172)),
	}})))),
	369: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	370: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(172)),
	}})))),
	371: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	372: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(150)),
	}})))),
	373: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	374: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(150)),
	}})))),
	375: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	376: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(135)),
	}})))),
	377: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	378: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(135)),
	}})))),
	379: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	380: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(136)),
	}})))),
	381: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	382: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(136)),
	}})))),
	383: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	384: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(137)),
	}})))),
	385: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	386: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(137)),
	}})))),
	387: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	388: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(138)),
	}})))),
	389: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	390: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(138)),
	}})))),
	391: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	392: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(139)),
	}})))),
	393: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	394: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(139)),
	}})))),
	395: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	396: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(140)),
	}})))),
	397: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	398: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(140)),
	}})))),
	399: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	400: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(141)),
	}})))),
	401: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	402: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(141)),
	}})))),
	403: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	404: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(163)),
	}})))),
	405: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	406: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(163)),
	}})))),
	407: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	408: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
	}})))),
	409: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	410: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(142)),
	}})))),
	411: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	412: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(142)),
	}})))),
	413: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	414: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
	}})))),
	415: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	416: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(171)),
	}})))),
	417: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	418: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(171)),
	}})))),
	419: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	420: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(144)),
	}})))),
	421: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	422: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(144)),
	}})))),
	423: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	424: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	425: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	426: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(132)),
	}})))),
	427: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	428: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(132)),
	}})))),
	429: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	430: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	431: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	432: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(99)),
	}})))),
	433: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	434: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(99)),
	}})))),
	435: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	436: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	437: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	438: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
	}})))),
	439: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	440: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
	}})))),
	441: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	442: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
	}})))),
	443: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	444: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(102)),
	}})))),
	445: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	446: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(102)),
	}})))),
	447: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	448: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(54)),
	}})))),
	449: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	450: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(104)),
	}})))),
	451: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	452: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(104)),
	}})))),
	453: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	454: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(55)),
	}})))),
	455: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	456: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(106)),
	}})))),
	457: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	458: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(106)),
	}})))),
	459: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	460: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	461: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	462: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(109)),
	}})))),
	463: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	464: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(109)),
	}})))),
	465: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	466: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	467: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	468: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	469: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	470: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	471: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	472: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	473: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	474: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
	}})))),
	475: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	476: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
	}})))),
	477: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	478: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	479: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	480: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(115)),
	}})))),
	481: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	482: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(115)),
	}})))),
	483: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	484: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(208)),
	}})))),
	485: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	486: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(212)),
	}})))),
	487: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	488: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(253)),
	}})))),
	489: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	490: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(248)),
	}})))),
	491: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	492: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(232)),
	}})))),
	493: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	494: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(208)),
	}})))),
	495: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	496: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	497: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	498: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
	}})))),
	499: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	500: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(16)),
	}})))),
	501: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	502: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	503: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	504: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(227)),
	}})))),
	505: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	506: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(240)),
	}})))),
	507: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	508: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(241)),
	}})))),
	509: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	510: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(241)),
	}})))),
	511: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	512: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	513: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	514: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	515: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	516: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	517: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	518: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	519: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	520: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	521: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	522: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	523: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	524: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	525: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	526: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	527: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	528: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(124)),
	}})))),
	529: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	530: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(124)),
	}})))),
	531: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	532: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(126)),
	}})))),
	533: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	534: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(126)),
	}})))),
	535: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	536: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(127)),
	}})))),
	537: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	538: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(127)),
	}})))),
	539: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	540: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(128)),
	}})))),
	541: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	542: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(128)),
	}})))),
	543: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	544: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(129)),
	}})))),
	545: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	546: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(129)),
	}})))),
	547: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	548: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(130)),
	}})))),
	549: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	550: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(130)),
	}})))),
	551: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	552: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(131)),
	}})))),
	553: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	554: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(131)),
	}})))),
	555: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	556: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	557: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	558: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(159)),
	}})))),
	559: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	560: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(159)),
	}})))),
	561: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	562: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
	}})))),
	563: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	564: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
	}})))),
	565: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	566: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
	}})))),
	567: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	568: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
	}})))),
	569: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	570: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	571: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	572: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
	}})))),
	573: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	574: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
	}})))),
	575: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	576: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(97)),
	}})))),
	577: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	578: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(97)),
	}})))),
	579: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	580: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	581: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	582: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(210)),
	}})))),
	583: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	584: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(210)),
	}})))),
	585: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	586: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	587: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	588: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(177)),
	}})))),
	589: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	590: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(177)),
	}})))),
	591: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	592: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	593: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	594: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(178)),
	}})))),
	595: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	596: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(178)),
	}})))),
	597: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	598: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	599: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	600: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(179)),
	}})))),
	601: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	602: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(179)),
	}})))),
	603: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	604: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	605: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	606: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(180)),
	}})))),
	607: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	608: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(180)),
	}})))),
	609: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	610: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	611: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	612: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(181)),
	}})))),
	613: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	614: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(181)),
	}})))),
	615: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	616: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	617: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	618: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(183)),
	}})))),
	619: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	620: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(183)),
	}})))),
	621: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	622: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	623: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	624: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(185)),
	}})))),
	625: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	626: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(185)),
	}})))),
	627: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	628: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	629: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	630: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(186)),
	}})))),
	631: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	632: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(186)),
	}})))),
	633: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	634: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(118)),
	}})))),
	635: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	636: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(118)),
	}})))),
	637: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	638: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(190)),
	}})))),
	639: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	640: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(190)),
	}})))),
	641: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	642: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(192)),
	}})))),
	643: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	644: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(192)),
	}})))),
	645: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	646: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(193)),
	}})))),
	647: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	648: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(193)),
	}})))),
	649: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	650: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(211)),
	}})))),
	651: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	652: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(211)),
	}})))),
	653: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	654: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(195)),
	}})))),
	655: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	656: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(195)),
	}})))),
	657: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	658: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(196)),
	}})))),
	659: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	660: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(196)),
	}})))),
	661: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	662: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(197)),
	}})))),
	663: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	664: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(197)),
	}})))),
	665: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	666: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(198)),
	}})))),
	667: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	668: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(198)),
	}})))),
	669: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	670: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(199)),
	}})))),
	671: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	672: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(199)),
	}})))),
	673: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	674: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	675: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	676: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(200)),
	}})))),
	677: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	678: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(200)),
	}})))),
	679: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	680: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(202)),
	}})))),
	681: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	682: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(202)),
	}})))),
	683: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	684: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	685: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	686: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(182)),
	}})))),
	687: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	688: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(182)),
	}})))),
	689: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	690: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(35)),
	}})))),
	691: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	692: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(153)),
	}})))),
	693: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	694: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(153)),
	}})))),
	695: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	696: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	697: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	698: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(107)),
	}})))),
	699: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	700: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(107)),
	}})))),
	701: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	702: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	703: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	704: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(148)),
	}})))),
	705: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	706: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(148)),
	}})))),
	707: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	708: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	709: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	710: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(149)),
	}})))),
	711: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	712: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(149)),
	}})))),
	713: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	714: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(28)),
	}})))),
	715: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	716: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(151)),
	}})))),
	717: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	718: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(151)),
	}})))),
	719: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	720: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
	}})))),
	721: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	722: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(152)),
	}})))),
	723: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	724: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(152)),
	}})))),
	725: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	726: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	727: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	728: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(154)),
	}})))),
	729: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	730: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(154)),
	}})))),
	731: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	732: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	733: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	734: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(155)),
	}})))),
	735: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	736: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(155)),
	}})))),
	737: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	738: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(111)),
	}})))),
	739: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	740: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(111)),
	}})))),
	741: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	742: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unsyntax),
	})))),
	743: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	744: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unsyntax),
	})))),
	745: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	746: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_extension),
	})))),
	747: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	748: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_extension),
	})))),
	749: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	750: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_graph),
	})))),
	751: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	752: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_graph),
	})))),
	753: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	754: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_hash),
	})))),
	755: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	756: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_hash),
	})))),
	757: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	758: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_graph),
	})))),
	759: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	760: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_graph),
	})))),
	761: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	762: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_comment),
	})))),
	763: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	764: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_comment),
	})))),
	765: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	766: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quote),
	})))),
	767: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	768: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quote),
	})))),
	769: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	770: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quasiquote),
	})))),
	771: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	772: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quasiquote),
	})))),
	773: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	774: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_byte_string),
	})))),
	775: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	776: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_byte_string),
	})))),
	777: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	778: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_syntax),
	})))),
	779: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	780: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_syntax),
	})))),
	781: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	782: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_vector),
	})))),
	783: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	784: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_vector),
	})))),
	785: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	786: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quasisyntax),
	})))),
	787: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	788: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_quasisyntax),
	})))),
	789: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	790: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_here_string),
	})))),
	791: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	792: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_here_string),
	})))),
	793: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	794: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unquote),
	})))),
	795: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	796: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unquote),
	})))),
	797: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	798: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_sexp_comment),
	})))),
	799: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	800: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_sexp_comment),
	})))),
	801: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	802: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_regex),
	})))),
	803: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	804: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_regex),
	})))),
	805: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	806: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unquote_splicing),
	})))),
	807: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	808: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unquote_splicing),
	})))),
	809: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	810: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__real_string),
	})))),
	811: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	812: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__real_string),
	})))),
	813: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	814: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unsyntax),
	})))),
	815: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	816: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unsyntax),
	})))),
	817: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	818: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unsyntax_splicing),
	})))),
	819: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	820: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unsyntax_splicing),
	})))),
	821: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	822: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unsyntax_splicing),
	})))),
	823: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	824: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unsyntax_splicing),
	})))),
	825: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	826: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_string),
	})))),
	827: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	828: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_string),
	})))),
	829: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	830: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_extension),
	})))),
	831: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	832: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_extension),
	})))),
	833: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	834: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_list),
	})))),
	835: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	836: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_list),
	})))),
	837: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	838: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_graph),
	})))),
	839: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	840: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_graph),
	})))),
	841: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	842: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_sexp_comment),
	})))),
	843: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	844: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_sexp_comment),
	})))),
	845: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	846: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_structure),
	})))),
	847: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	848: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_structure),
	})))),
	849: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	850: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_vector),
	})))),
	851: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	852: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_vector),
	})))),
	853: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	854: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_block_comment),
	})))),
	855: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	856: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_block_comment),
	})))),
	857: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	858: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_block_comment),
	})))),
	859: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	860: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_block_comment),
	})))),
	861: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	862: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__real_string),
	})))),
	863: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	864: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__real_string),
	})))),
	865: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	866: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_box),
	})))),
	867: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	868: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_box),
	})))),
	869: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	870: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list),
	})))),
	871: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	872: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list),
	})))),
	873: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	874: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quote),
	})))),
	875: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	876: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quote),
	})))),
	877: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	878: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quasiquote),
	})))),
	879: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	880: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quasiquote),
	})))),
	881: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	882: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_syntax),
	})))),
	883: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	884: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_syntax),
	})))),
	885: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	886: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quasisyntax),
	})))),
	887: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	888: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quasisyntax),
	})))),
	889: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	890: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unquote),
	})))),
	891: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	892: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unquote),
	})))),
	893: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	894: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unquote_splicing),
	})))),
	895: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	896: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_unquote_splicing),
	})))),
	897: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	898: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_box),
	})))),
	899: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	900: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_box),
	})))),
	901: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	902: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_sexp_comment_repeat1),
	})))),
	903: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(133)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	904: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	905: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_sexp_comment_repeat1),
	})))),
	906: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(184)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	907: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	908: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_sexp_comment_repeat1),
	})))),
	909: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(225)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	910: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	911: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_sexp_comment_repeat1),
	})))),
	912: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(83)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	913: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	914: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_sexp_comment_repeat1),
	})))),
	915: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	916: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_sexp_comment_repeat1),
	})))),
	917: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	918: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(215)),
	}})))),
	919: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	920: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(216)),
	}})))),
	921: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	922: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(217)),
	}})))),
	923: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	924: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(188)),
	}})))),
	925: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	926: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	927: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	928: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(165)),
	}})))),
	929: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	930: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
	}})))),
	931: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	932: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(117)),
	}})))),
	933: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	934: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	935: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	936: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(218)),
	}})))),
	937: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	938: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(222)),
	}})))),
	939: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	940: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(246)),
	}})))),
	941: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	942: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(223)),
	}})))),
	943: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	944: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(122)),
	}})))),
	945: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	946: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(226)),
	}})))),
	947: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	948: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(173)),
	}})))),
	949: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	950: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(230)),
	}})))),
	951: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	952: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(247)),
	}})))),
	953: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	954: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_comment_repeat1),
	})))),
	955: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(218)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	956: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	957: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_comment_repeat1),
	})))),
	958: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(223)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	959: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	960: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_comment_repeat1),
	})))),
	961: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	962: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(243)),
	}})))),
	963: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	964: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(228)),
	}})))),
	965: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	966: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(203)),
	}})))),
	967: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	968: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(162)),
	}})))),
	969: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	970: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(239)),
	}})))),
	971: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	972: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(187)),
	}})))),
	973: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	974: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(219)),
	}})))),
	975: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	976: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(121)),
	}})))),
	977: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	978: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(168)),
	}})))),
	979: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	980: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(237)),
	}})))),
	981: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	982: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(209)),
	}})))),
	983: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	984: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(244)),
	}})))),
	985: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	986: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__real_string_repeat1),
	})))),
	987: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	988: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__real_string_repeat1),
	})))),
	989: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(237)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	990: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	991: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(123)),
	}})))),
	992: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	993: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
	}})))),
	994: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	995: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(238)),
	}})))),
	996: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	997: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(189)),
	}})))),
	998: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	999: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(164)),
	}})))),
	1000: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1001: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(231)),
	}})))),
	1002: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1003: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(105)),
	}})))),
	1004: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1005: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(206)),
	}})))),
	1006: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1007: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	1008: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1009: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(160)),
	}})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token__here_string_body = 0

var ts_external_scanner_symbol_map = [1]TSSymbol{
	0: uint16(sym__here_string_body),
}

var ts_external_scanner_states = [2][1]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_racket(tls *libc.TLS) (r uintptr) {
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
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_racket_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_racket_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_racket_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_racket_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_racket_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "Scanner: Failed to allocate memory\n\x00end\x00_skip_token1\x00dot\x00comment_token1\x00#|\x00block_comment_token1\x00|#\x00#;\x00_line_comment\x00boolean\x00#\x00#<<\x00regex_token1\x00\"\x00_real_string_token1\x00escape_sequence\x00number\x00decimal\x00character\x00symbol\x00keyword\x00#&\x00(\x00)\x00[\x00]\x00{\x00}\x00#fl\x00#fx\x00#s\x00#hash\x00#hashalw\x00#hasheq\x00#hasheqv\x00=\x00'\x00`\x00#'\x00#`\x00,\x00,@\x00#,\x00#,@\x00#reader\x00#lang \x00#!\x00lang_name\x00_here_string_body\x00program\x00_token\x00_skip\x00comment\x00block_comment\x00sexp_comment\x00_datum\x00string\x00byte_string\x00here_string\x00regex\x00_real_string\x00box\x00list\x00vector\x00structure\x00hash\x00graph\x00quote\x00quasiquote\x00syntax\x00quasisyntax\x00unquote\x00unquote_splicing\x00unsyntax\x00unsyntax_splicing\x00extension\x00program_repeat1\x00block_comment_repeat1\x00sexp_comment_repeat1\x00_real_string_repeat1\x00list_repeat1\x00"
