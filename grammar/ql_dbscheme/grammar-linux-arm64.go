// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-ql-dbscheme/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-ql-dbscheme -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-ql-dbscheme/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_ql_dbscheme

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 0
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXTERNAL_TOKEN_COUNT = 0
const FIELD_COUNT = 12
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
const MAX_ALIAS_SEQUENCE_LENGTH = 8
const PRODUCTION_ID_COUNT = 17
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 99
const SYMBOL_COUNT = 49
const TOKEN_COUNT = 29
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
const _LP64 = 1
const __AARCH64EL__ = 1
const __AARCH64_CMODEL_SMALL__ = 1
const __ARM_64BIT_STATE = 1
const __ARM_ACLE = 200
const __ARM_ALIGN_MAX_STACK_PWR = 4
const __ARM_ARCH = 8
const __ARM_ARCH_ISA_A64 = 1
const __ARM_ARCH_PROFILE = 'A'
const __ARM_FEATURE_CLZ = 1
const __ARM_FEATURE_DIRECTED_ROUNDING = 1
const __ARM_FEATURE_DIV = 1
const __ARM_FEATURE_FMA = 1
const __ARM_FEATURE_IDIV = 1
const __ARM_FEATURE_LDREX = 0xF
const __ARM_FEATURE_NUMERIC_MAXMIN = 1
const __ARM_FEATURE_UNALIGNED = 1
const __ARM_FP = 0xE
const __ARM_FP16_ARGS = 1
const __ARM_FP16_FORMAT_IEEE = 1
const __ARM_NEON = 1
const __ARM_NEON_FP = 0xE
const __ARM_PCS_AAPCS64 = 1
const __ARM_SIZEOF_MINIMAL_ENUM = 4
const __ARM_SIZEOF_WCHAR_T = 4
const __ARM_STATE_ZA = 1
const __ARM_STATE_ZT0 = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BIGGEST_ALIGNMENT__ = 16
const __BIG_ENDIAN = 4321
const __BITINT_MAXWIDTH__ = 128
const __BOOL_WIDTH__ = 8
const __BYTE_ORDER = 1234
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CHAR_BIT__ = 8
const __CHAR_UNSIGNED__ = 1
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
const __FP_FAST_FMA = 1
const __FP_FAST_FMAF = 1
const __FUNCTION__ = "__func__"
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
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_16 = 1
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
const __LDBL_DECIMAL_DIG__ = 36
const __LDBL_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __LDBL_DIG__ = 33
const __LDBL_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_MANT_DIG__ = 113
const __LDBL_MAX_10_EXP__ = 4932
const __LDBL_MAX_EXP__ = 16384
const __LDBL_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __LDBL_MIN__ = 3.36210314311209350626267781732175260e-4932
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
const __NO_INLINE__ = 1
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
const __SHRT_MAX__ = 32767
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 2147483647
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
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
const __WCHAR_MAX__ = 4294967295
const __WCHAR_UNSIGNED__ = 1
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 4294967295
const __WINT_UNSIGNED__ = 1
const __WINT_WIDTH__ = 32
const __aarch64__ = 1
const __bool_true_false_are_defined = 1
const __clang__ = 1
const __clang_literal_encoding__ = "UTF-8"
const __clang_major__ = 18
const __clang_minor__ = 1
const __clang_patchlevel__ = 3
const __clang_version__ = "18.1.3 (1ubuntu1)"
const __clang_wide_literal_encoding__ = "UTF-32"
const __gnu_linux__ = 1
const __inline = "inline"
const __linux = 1
const __linux__ = 1
const __llvm__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict = "restrict"
const __restrict_arr = "restrict"
const __unix = 1
const __unix__ = 1
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
const true1 = 1
const ts_builtin_sym_end = 0
const type1 = "type_token"
const unix = 1
const var1 = "var_token"

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = uint32

type __predefined_ptrdiff_t = int64

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

type wchar_t = uint32

type size_t = uint64

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

type ts_symbol_identifiers = int32

const sym__lower_id = 1
const anon_sym_LPAREN = 2
const anon_sym_COMMA = 3
const anon_sym_RPAREN = 4
const anon_sym_SEMI = 5
const anon_sym_POUND = 6
const anon_sym_LBRACK = 7
const anon_sym_RBRACK = 8
const anon_sym_COLON = 9
const anon_sym_EQ = 10
const anon_sym_PIPE = 11
const anon_sym_case = 12
const anon_sym_DOT = 13
const anon_sym_of = 14
const sym_ref = 15
const sym_int = 16
const sym_float = 17
const sym_boolean = 18
const sym_date = 19
const sym_varchar = 20
const sym_string = 21
const sym_unique = 22
const sym__upper_id = 23
const sym_integer = 24
const sym_dbtype = 25
const sym_qldoc = 26
const sym_line_comment = 27
const sym_block_comment = 28
const sym_dbscheme = 29
const sym_entry = 30
const sym_table = 31
const sym_annotation = 32
const sym_argsAnnotation = 33
const sym_tableName = 34
const sym_column = 35
const sym_unionDecl = 36
const sym_caseDecl = 37
const sym_branch = 38
const sym_colType = 39
const sym_reprType = 40
const sym_annotName = 41
const sym_simpleId = 42
const aux_sym_dbscheme_repeat1 = 43
const aux_sym_table_repeat1 = 44
const aux_sym_table_repeat2 = 45
const aux_sym_argsAnnotation_repeat1 = 46
const aux_sym_unionDecl_repeat1 = 47
const aux_sym_caseDecl_repeat1 = 48

var ts_symbol_names = [49]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 14,
	3:  __ccgo_ts + 16,
	4:  __ccgo_ts + 18,
	5:  __ccgo_ts + 20,
	6:  __ccgo_ts + 22,
	7:  __ccgo_ts + 24,
	8:  __ccgo_ts + 26,
	9:  __ccgo_ts + 28,
	10: __ccgo_ts + 30,
	11: __ccgo_ts + 32,
	12: __ccgo_ts + 34,
	13: __ccgo_ts + 39,
	14: __ccgo_ts + 41,
	15: __ccgo_ts + 44,
	16: __ccgo_ts + 48,
	17: __ccgo_ts + 52,
	18: __ccgo_ts + 58,
	19: __ccgo_ts + 66,
	20: __ccgo_ts + 71,
	21: __ccgo_ts + 79,
	22: __ccgo_ts + 86,
	23: __ccgo_ts + 93,
	24: __ccgo_ts + 103,
	25: __ccgo_ts + 111,
	26: __ccgo_ts + 118,
	27: __ccgo_ts + 124,
	28: __ccgo_ts + 137,
	29: __ccgo_ts + 151,
	30: __ccgo_ts + 160,
	31: __ccgo_ts + 166,
	32: __ccgo_ts + 172,
	33: __ccgo_ts + 183,
	34: __ccgo_ts + 198,
	35: __ccgo_ts + 208,
	36: __ccgo_ts + 215,
	37: __ccgo_ts + 225,
	38: __ccgo_ts + 234,
	39: __ccgo_ts + 241,
	40: __ccgo_ts + 249,
	41: __ccgo_ts + 258,
	42: __ccgo_ts + 268,
	43: __ccgo_ts + 277,
	44: __ccgo_ts + 294,
	45: __ccgo_ts + 308,
	46: __ccgo_ts + 322,
	47: __ccgo_ts + 345,
	48: __ccgo_ts + 363,
}

var ts_symbol_map = [49]TSSymbol{
	1:  uint16(sym__lower_id),
	2:  uint16(anon_sym_LPAREN),
	3:  uint16(anon_sym_COMMA),
	4:  uint16(anon_sym_RPAREN),
	5:  uint16(anon_sym_SEMI),
	6:  uint16(anon_sym_POUND),
	7:  uint16(anon_sym_LBRACK),
	8:  uint16(anon_sym_RBRACK),
	9:  uint16(anon_sym_COLON),
	10: uint16(anon_sym_EQ),
	11: uint16(anon_sym_PIPE),
	12: uint16(anon_sym_case),
	13: uint16(anon_sym_DOT),
	14: uint16(anon_sym_of),
	15: uint16(sym_ref),
	16: uint16(sym_int),
	17: uint16(sym_float),
	18: uint16(sym_boolean),
	19: uint16(sym_date),
	20: uint16(sym_varchar),
	21: uint16(sym_string),
	22: uint16(sym_unique),
	23: uint16(sym__upper_id),
	24: uint16(sym_integer),
	25: uint16(sym_dbtype),
	26: uint16(sym_qldoc),
	27: uint16(sym_line_comment),
	28: uint16(sym_block_comment),
	29: uint16(sym_dbscheme),
	30: uint16(sym_entry),
	31: uint16(sym_table),
	32: uint16(sym_annotation),
	33: uint16(sym_argsAnnotation),
	34: uint16(sym_tableName),
	35: uint16(sym_column),
	36: uint16(sym_unionDecl),
	37: uint16(sym_caseDecl),
	38: uint16(sym_branch),
	39: uint16(sym_colType),
	40: uint16(sym_reprType),
	41: uint16(sym_annotName),
	42: uint16(sym_simpleId),
	43: uint16(aux_sym_dbscheme_repeat1),
	44: uint16(aux_sym_table_repeat1),
	45: uint16(aux_sym_table_repeat2),
	46: uint16(aux_sym_argsAnnotation_repeat1),
	47: uint16(aux_sym_unionDecl_repeat1),
	48: uint16(aux_sym_caseDecl_repeat1),
}

var ts_symbol_metadata = [49]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	2: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
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
	},
	10: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	11: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	12: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	13: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	23: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	39: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	40: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	41: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	42: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	43: {},
	44: {},
	45: {},
	46: {},
	47: {},
	48: {},
}

type ts_field_identifiers = int32

const field_argsAnnotation = 1
const field_base = 2
const field_colName = 3
const field_colType = 4
const field_discriminator = 5
const field_isRef = 6
const field_isUnique = 7
const field_name = 8
const field_qldoc = 9
const field_reprType = 10
const field_simpleAnnotation = 11
const field_tableName = 12

var ts_field_names = [13]uintptr{
	0:  libc.UintptrFromInt32(0),
	1:  __ccgo_ts + 183,
	2:  __ccgo_ts + 380,
	3:  __ccgo_ts + 385,
	4:  __ccgo_ts + 241,
	5:  __ccgo_ts + 393,
	6:  __ccgo_ts + 407,
	7:  __ccgo_ts + 413,
	8:  __ccgo_ts + 422,
	9:  __ccgo_ts + 118,
	10: __ccgo_ts + 249,
	11: __ccgo_ts + 427,
	12: __ccgo_ts + 198,
}

var ts_field_map_slices = [17]TSFieldMapSlice{
	1: {
		Flength: uint16(1),
	},
	2: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	3: {
		Findex:  uint16(2),
		Flength: uint16(1),
	},
	4: {
		Findex:  uint16(3),
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(4),
		Flength: uint16(1),
	},
	6: {
		Findex:  uint16(5),
		Flength: uint16(1),
	},
	7: {
		Findex:  uint16(6),
		Flength: uint16(2),
	},
	8: {
		Findex:  uint16(8),
		Flength: uint16(3),
	},
	9: {
		Findex:  uint16(11),
		Flength: uint16(4),
	},
	10: {
		Findex:  uint16(15),
		Flength: uint16(4),
	},
	11: {
		Findex:  uint16(19),
		Flength: uint16(4),
	},
	12: {
		Findex:  uint16(23),
		Flength: uint16(5),
	},
	13: {
		Findex:  uint16(28),
		Flength: uint16(5),
	},
	14: {
		Findex:  uint16(33),
		Flength: uint16(5),
	},
	15: {
		Findex:  uint16(38),
		Flength: uint16(1),
	},
	16: {
		Findex:  uint16(39),
		Flength: uint16(6),
	},
}

var ts_field_map_entries = [45]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_argsAnnotation),
	},
	1: {
		Ffield_id:    uint16(field_simpleAnnotation),
		Fchild_index: uint8(1),
	},
	2: {
		Ffield_id: uint16(field_base),
	},
	3: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	4: {
		Ffield_id: uint16(field_tableName),
	},
	5: {
		Ffield_id:    uint16(field_tableName),
		Fchild_index: uint8(1),
	},
	6: {
		Ffield_id:    uint16(field_base),
		Fchild_index: uint8(1),
	},
	7: {
		Ffield_id:    uint16(field_discriminator),
		Fchild_index: uint8(3),
	},
	8: {
		Ffield_id:    uint16(field_colName),
		Fchild_index: uint8(1),
	},
	9: {
		Ffield_id:    uint16(field_colType),
		Fchild_index: uint8(3),
	},
	10: {
		Ffield_id: uint16(field_reprType),
	},
	11: {
		Ffield_id:    uint16(field_colName),
		Fchild_index: uint8(2),
	},
	12: {
		Ffield_id:    uint16(field_colType),
		Fchild_index: uint8(4),
	},
	13: {
		Ffield_id: uint16(field_isUnique),
	},
	14: {
		Ffield_id:    uint16(field_reprType),
		Fchild_index: uint8(1),
	},
	15: {
		Ffield_id:    uint16(field_colName),
		Fchild_index: uint8(2),
	},
	16: {
		Ffield_id:    uint16(field_colType),
		Fchild_index: uint8(4),
	},
	17: {
		Ffield_id: uint16(field_qldoc),
	},
	18: {
		Ffield_id:    uint16(field_reprType),
		Fchild_index: uint8(1),
	},
	19: {
		Ffield_id:    uint16(field_colName),
		Fchild_index: uint8(1),
	},
	20: {
		Ffield_id:    uint16(field_colType),
		Fchild_index: uint8(3),
	},
	21: {
		Ffield_id:    uint16(field_isRef),
		Fchild_index: uint8(4),
	},
	22: {
		Ffield_id: uint16(field_reprType),
	},
	23: {
		Ffield_id:    uint16(field_colName),
		Fchild_index: uint8(2),
	},
	24: {
		Ffield_id:    uint16(field_colType),
		Fchild_index: uint8(4),
	},
	25: {
		Ffield_id:    uint16(field_isRef),
		Fchild_index: uint8(5),
	},
	26: {
		Ffield_id: uint16(field_isUnique),
	},
	27: {
		Ffield_id:    uint16(field_reprType),
		Fchild_index: uint8(1),
	},
	28: {
		Ffield_id:    uint16(field_colName),
		Fchild_index: uint8(3),
	},
	29: {
		Ffield_id:    uint16(field_colType),
		Fchild_index: uint8(5),
	},
	30: {
		Ffield_id:    uint16(field_isUnique),
		Fchild_index: uint8(1),
	},
	31: {
		Ffield_id: uint16(field_qldoc),
	},
	32: {
		Ffield_id:    uint16(field_reprType),
		Fchild_index: uint8(2),
	},
	33: {
		Ffield_id:    uint16(field_colName),
		Fchild_index: uint8(2),
	},
	34: {
		Ffield_id:    uint16(field_colType),
		Fchild_index: uint8(4),
	},
	35: {
		Ffield_id:    uint16(field_isRef),
		Fchild_index: uint8(5),
	},
	36: {
		Ffield_id: uint16(field_qldoc),
	},
	37: {
		Ffield_id:    uint16(field_reprType),
		Fchild_index: uint8(1),
	},
	38: {
		Ffield_id: uint16(field_qldoc),
	},
	39: {
		Ffield_id:    uint16(field_colName),
		Fchild_index: uint8(3),
	},
	40: {
		Ffield_id:    uint16(field_colType),
		Fchild_index: uint8(5),
	},
	41: {
		Ffield_id:    uint16(field_isRef),
		Fchild_index: uint8(6),
	},
	42: {
		Ffield_id:    uint16(field_isUnique),
		Fchild_index: uint8(1),
	},
	43: {
		Ffield_id: uint16(field_qldoc),
	},
	44: {
		Ffield_id:    uint16(field_reprType),
		Fchild_index: uint8(2),
	},
}

var ts_alias_sequences = [17][8]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [99]TSStateId{
	1:  uint16(1),
	2:  uint16(2),
	3:  uint16(3),
	4:  uint16(4),
	5:  uint16(5),
	6:  uint16(6),
	7:  uint16(7),
	8:  uint16(8),
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
	20: uint16(20),
	21: uint16(21),
	22: uint16(22),
	23: uint16(23),
	24: uint16(24),
	25: uint16(25),
	26: uint16(26),
	27: uint16(27),
	28: uint16(28),
	29: uint16(29),
	30: uint16(30),
	31: uint16(31),
	32: uint16(32),
	33: uint16(33),
	34: uint16(34),
	35: uint16(35),
	36: uint16(36),
	37: uint16(37),
	38: uint16(38),
	39: uint16(39),
	40: uint16(40),
	41: uint16(41),
	42: uint16(42),
	43: uint16(43),
	44: uint16(44),
	45: uint16(45),
	46: uint16(46),
	47: uint16(47),
	48: uint16(48),
	49: uint16(49),
	50: uint16(50),
	51: uint16(51),
	52: uint16(52),
	53: uint16(53),
	54: uint16(54),
	55: uint16(55),
	56: uint16(56),
	57: uint16(57),
	58: uint16(58),
	59: uint16(59),
	60: uint16(60),
	61: uint16(61),
	62: uint16(62),
	63: uint16(63),
	64: uint16(64),
	65: uint16(65),
	66: uint16(66),
	67: uint16(67),
	68: uint16(68),
	69: uint16(69),
	70: uint16(70),
	71: uint16(71),
	72: uint16(72),
	73: uint16(73),
	74: uint16(74),
	75: uint16(75),
	76: uint16(76),
	77: uint16(77),
	78: uint16(78),
	79: uint16(79),
	80: uint16(80),
	81: uint16(81),
	82: uint16(82),
	83: uint16(83),
	84: uint16(84),
	85: uint16(85),
	86: uint16(86),
	87: uint16(87),
	88: uint16(88),
	89: uint16(89),
	90: uint16(90),
	91: uint16(91),
	92: uint16(92),
	93: uint16(93),
	94: uint16(94),
	95: uint16(95),
	96: uint16(96),
	97: uint16(97),
	98: uint16(98),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _ = eof, i, i1, lookahead, result, skip
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
			state = uint16(13)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(27)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(26)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('*') {
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('*') {
			state = uint16(5)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('*') {
			state = uint16(10)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('*') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('*') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(32)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('*') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(29)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('*') {
			state = uint16(6)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('*') {
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(31)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('*') {
			state = uint16(8)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('/') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(11):
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(12):
		if eof != 0 {
			state = uint16(13)
			goto next_state
		}
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token1[i1]) == lookahead {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _2
		_2:
			;
			i1 = i1 + uint32(2)
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(27)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(26)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(13):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(14):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(15):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(16):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(17):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(18):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(19):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(20):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(22):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(23):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(24):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__lower_id)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__upper_id)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_dbtype)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_qldoc)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_line_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_block_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_block_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(6)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(7)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [26]uint16_t{
	0:  uint16('#'),
	1:  uint16(18),
	2:  uint16('('),
	3:  uint16(14),
	4:  uint16(')'),
	5:  uint16(16),
	6:  uint16(','),
	7:  uint16(15),
	8:  uint16('.'),
	9:  uint16(24),
	10: uint16('/'),
	11: uint16(1),
	12: uint16(':'),
	13: uint16(21),
	14: uint16(';'),
	15: uint16(17),
	16: uint16('='),
	17: uint16(22),
	18: uint16('@'),
	19: uint16(11),
	20: uint16('['),
	21: uint16(19),
	22: uint16(']'),
	23: uint16(20),
	24: uint16('|'),
	25: uint16(23),
}

var map_token1 = [22]uint16_t{
	0:  uint16('#'),
	1:  uint16(18),
	2:  uint16('('),
	3:  uint16(14),
	4:  uint16(')'),
	5:  uint16(16),
	6:  uint16(','),
	7:  uint16(15),
	8:  uint16('.'),
	9:  uint16(24),
	10: uint16('/'),
	11: uint16(4),
	12: uint16(':'),
	13: uint16(21),
	14: uint16('='),
	15: uint16(22),
	16: uint16('@'),
	17: uint16(11),
	18: uint16('['),
	19: uint16(19),
	20: uint16(']'),
	21: uint16(20),
}

func ts_lex_keywords(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
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
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token2[i]) == lookahead {
				state = map_token2[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('o') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('a') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('a') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('l') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('n') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('f') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('e') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('t') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('n') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('a') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('o') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('s') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('t') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('o') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('t') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(16):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_of)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(17):
		if lookahead == int32('f') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('r') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('i') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('r') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('l') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('e') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('e') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('a') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_ref)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		if lookahead == int32('i') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('q') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('c') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('e') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_case)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_date)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		if lookahead == int32('t') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('n') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('u') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('h') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('a') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		if lookahead == int32('g') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('e') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('a') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('n') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unique)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		if lookahead == int32('r') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_boolean)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_varchar)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token2 = [20]uint16_t{
	0:  uint16('b'),
	1:  uint16(1),
	2:  uint16('c'),
	3:  uint16(2),
	4:  uint16('d'),
	5:  uint16(3),
	6:  uint16('f'),
	7:  uint16(4),
	8:  uint16('i'),
	9:  uint16(5),
	10: uint16('o'),
	11: uint16(6),
	12: uint16('r'),
	13: uint16(7),
	14: uint16('s'),
	15: uint16(8),
	16: uint16('u'),
	17: uint16(9),
	18: uint16('v'),
	19: uint16(10),
}

var ts_lex_modes = [99]TSLexMode{
	0:  {},
	1:  {},
	2:  {},
	3:  {},
	4:  {},
	5:  {},
	6:  {},
	7:  {},
	8:  {},
	9:  {},
	10: {},
	11: {},
	12: {},
	13: {},
	14: {},
	15: {},
	16: {},
	17: {},
	18: {},
	19: {
		Flex_state: uint16(12),
	},
	20: {},
	21: {},
	22: {
		Flex_state: uint16(12),
	},
	23: {},
	24: {
		Flex_state: uint16(12),
	},
	25: {},
	26: {},
	27: {
		Flex_state: uint16(12),
	},
	28: {
		Flex_state: uint16(12),
	},
	29: {},
	30: {},
	31: {},
	32: {},
	33: {
		Flex_state: uint16(12),
	},
	34: {},
	35: {
		Flex_state: uint16(12),
	},
	36: {},
	37: {
		Flex_state: uint16(12),
	},
	38: {
		Flex_state: uint16(12),
	},
	39: {
		Flex_state: uint16(12),
	},
	40: {
		Flex_state: uint16(12),
	},
	41: {
		Flex_state: uint16(12),
	},
	42: {
		Flex_state: uint16(12),
	},
	43: {
		Flex_state: uint16(12),
	},
	44: {
		Flex_state: uint16(12),
	},
	45: {},
	46: {
		Flex_state: uint16(12),
	},
	47: {
		Flex_state: uint16(12),
	},
	48: {
		Flex_state: uint16(12),
	},
	49: {
		Flex_state: uint16(12),
	},
	50: {
		Flex_state: uint16(12),
	},
	51: {
		Flex_state: uint16(12),
	},
	52: {
		Flex_state: uint16(12),
	},
	53: {
		Flex_state: uint16(12),
	},
	54: {
		Flex_state: uint16(12),
	},
	55: {
		Flex_state: uint16(12),
	},
	56: {
		Flex_state: uint16(12),
	},
	57: {
		Flex_state: uint16(12),
	},
	58: {},
	59: {
		Flex_state: uint16(12),
	},
	60: {
		Flex_state: uint16(12),
	},
	61: {
		Flex_state: uint16(12),
	},
	62: {
		Flex_state: uint16(12),
	},
	63: {
		Flex_state: uint16(12),
	},
	64: {
		Flex_state: uint16(12),
	},
	65: {
		Flex_state: uint16(12),
	},
	66: {
		Flex_state: uint16(12),
	},
	67: {
		Flex_state: uint16(12),
	},
	68: {
		Flex_state: uint16(12),
	},
	69: {
		Flex_state: uint16(12),
	},
	70: {
		Flex_state: uint16(12),
	},
	71: {
		Flex_state: uint16(12),
	},
	72: {
		Flex_state: uint16(12),
	},
	73: {
		Flex_state: uint16(12),
	},
	74: {
		Flex_state: uint16(12),
	},
	75: {
		Flex_state: uint16(12),
	},
	76: {
		Flex_state: uint16(12),
	},
	77: {
		Flex_state: uint16(12),
	},
	78: {
		Flex_state: uint16(12),
	},
	79: {
		Flex_state: uint16(12),
	},
	80: {
		Flex_state: uint16(12),
	},
	81: {
		Flex_state: uint16(12),
	},
	82: {
		Flex_state: uint16(12),
	},
	83: {
		Flex_state: uint16(12),
	},
	84: {
		Flex_state: uint16(12),
	},
	85: {
		Flex_state: uint16(12),
	},
	86: {
		Flex_state: uint16(12),
	},
	87: {
		Flex_state: uint16(12),
	},
	88: {
		Flex_state: uint16(12),
	},
	89: {
		Flex_state: uint16(12),
	},
	90: {
		Flex_state: uint16(12),
	},
	91: {
		Flex_state: uint16(12),
	},
	92: {
		Flex_state: uint16(12),
	},
	93: {
		Flex_state: uint16(12),
	},
	94: {
		Flex_state: uint16(12),
	},
	95: {
		Flex_state: uint16(12),
	},
	96: {
		Flex_state: uint16(12),
	},
	97: {
		Flex_state: uint16(12),
	},
	98: {
		Flex_state: uint16(12),
	},
}

var ts_parse_table = [2][49]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		3:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		6:  uint16(1),
		7:  uint16(1),
		8:  uint16(1),
		9:  uint16(1),
		10: uint16(1),
		11: uint16(1),
		12: uint16(1),
		13: uint16(1),
		14: uint16(1),
		15: uint16(1),
		16: uint16(1),
		17: uint16(1),
		18: uint16(1),
		19: uint16(1),
		20: uint16(1),
		21: uint16(1),
		22: uint16(1),
		23: uint16(1),
		24: uint16(1),
		25: uint16(1),
		26: uint16(1),
		27: uint16(3),
		28: uint16(5),
	},
	1: {
		0:  uint16(7),
		1:  uint16(9),
		6:  uint16(11),
		12: uint16(13),
		23: uint16(15),
		25: uint16(17),
		26: uint16(19),
		27: uint16(3),
		28: uint16(5),
		29: uint16(81),
		30: uint16(3),
		31: uint16(36),
		32: uint16(19),
		33: uint16(46),
		34: uint16(93),
		36: uint16(36),
		37: uint16(36),
		42: uint16(88),
		43: uint16(3),
		44: uint16(19),
	},
}

var ts_small_parse_table = [1460]uint16_t{
	0:    uint16(15),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_line_comment),
	4:    uint16(5),
	5:    uint16(1),
	6:    uint16(sym_block_comment),
	7:    uint16(21),
	8:    uint16(1),
	10:   uint16(23),
	11:   uint16(1),
	12:   uint16(sym__lower_id),
	13:   uint16(26),
	14:   uint16(1),
	15:   uint16(anon_sym_POUND),
	16:   uint16(29),
	17:   uint16(1),
	18:   uint16(anon_sym_case),
	19:   uint16(32),
	20:   uint16(1),
	21:   uint16(sym__upper_id),
	22:   uint16(35),
	23:   uint16(1),
	24:   uint16(sym_dbtype),
	25:   uint16(38),
	26:   uint16(1),
	27:   uint16(sym_qldoc),
	28:   uint16(46),
	29:   uint16(1),
	30:   uint16(sym_argsAnnotation),
	31:   uint16(88),
	32:   uint16(1),
	33:   uint16(sym_simpleId),
	34:   uint16(93),
	35:   uint16(1),
	36:   uint16(sym_tableName),
	37:   uint16(2),
	38:   uint16(2),
	39:   uint16(sym_entry),
	40:   uint16(aux_sym_dbscheme_repeat1),
	41:   uint16(19),
	42:   uint16(2),
	43:   uint16(sym_annotation),
	44:   uint16(aux_sym_table_repeat1),
	45:   uint16(36),
	46:   uint16(3),
	47:   uint16(sym_table),
	48:   uint16(sym_unionDecl),
	49:   uint16(sym_caseDecl),
	50:   uint16(15),
	51:   uint16(3),
	52:   uint16(1),
	53:   uint16(sym_line_comment),
	54:   uint16(5),
	55:   uint16(1),
	56:   uint16(sym_block_comment),
	57:   uint16(9),
	58:   uint16(1),
	59:   uint16(sym__lower_id),
	60:   uint16(11),
	61:   uint16(1),
	62:   uint16(anon_sym_POUND),
	63:   uint16(13),
	64:   uint16(1),
	65:   uint16(anon_sym_case),
	66:   uint16(15),
	67:   uint16(1),
	68:   uint16(sym__upper_id),
	69:   uint16(17),
	70:   uint16(1),
	71:   uint16(sym_dbtype),
	72:   uint16(19),
	73:   uint16(1),
	74:   uint16(sym_qldoc),
	75:   uint16(41),
	76:   uint16(1),
	78:   uint16(46),
	79:   uint16(1),
	80:   uint16(sym_argsAnnotation),
	81:   uint16(88),
	82:   uint16(1),
	83:   uint16(sym_simpleId),
	84:   uint16(93),
	85:   uint16(1),
	86:   uint16(sym_tableName),
	87:   uint16(2),
	88:   uint16(2),
	89:   uint16(sym_entry),
	90:   uint16(aux_sym_dbscheme_repeat1),
	91:   uint16(19),
	92:   uint16(2),
	93:   uint16(sym_annotation),
	94:   uint16(aux_sym_table_repeat1),
	95:   uint16(36),
	96:   uint16(3),
	97:   uint16(sym_table),
	98:   uint16(sym_unionDecl),
	99:   uint16(sym_caseDecl),
	100:  uint16(7),
	101:  uint16(3),
	102:  uint16(1),
	103:  uint16(sym_line_comment),
	104:  uint16(5),
	105:  uint16(1),
	106:  uint16(sym_block_comment),
	107:  uint16(47),
	108:  uint16(1),
	109:  uint16(anon_sym_SEMI),
	110:  uint16(49),
	111:  uint16(1),
	112:  uint16(anon_sym_PIPE),
	113:  uint16(8),
	114:  uint16(1),
	115:  uint16(aux_sym_unionDecl_repeat1),
	116:  uint16(45),
	117:  uint16(2),
	118:  uint16(anon_sym_case),
	119:  uint16(sym__lower_id),
	120:  uint16(43),
	121:  uint16(5),
	123:  uint16(anon_sym_POUND),
	124:  uint16(sym__upper_id),
	125:  uint16(sym_dbtype),
	126:  uint16(sym_qldoc),
	127:  uint16(8),
	128:  uint16(3),
	129:  uint16(1),
	130:  uint16(sym_line_comment),
	131:  uint16(5),
	132:  uint16(1),
	133:  uint16(sym_block_comment),
	134:  uint16(53),
	135:  uint16(1),
	136:  uint16(sym_varchar),
	137:  uint16(55),
	138:  uint16(1),
	139:  uint16(sym_unique),
	140:  uint16(57),
	141:  uint16(1),
	142:  uint16(sym_qldoc),
	143:  uint16(48),
	144:  uint16(1),
	145:  uint16(sym_column),
	146:  uint16(56),
	147:  uint16(1),
	148:  uint16(sym_reprType),
	149:  uint16(51),
	150:  uint16(5),
	151:  uint16(sym_int),
	152:  uint16(sym_float),
	153:  uint16(sym_boolean),
	154:  uint16(sym_date),
	155:  uint16(sym_string),
	156:  uint16(8),
	157:  uint16(3),
	158:  uint16(1),
	159:  uint16(sym_line_comment),
	160:  uint16(5),
	161:  uint16(1),
	162:  uint16(sym_block_comment),
	163:  uint16(53),
	164:  uint16(1),
	165:  uint16(sym_varchar),
	166:  uint16(55),
	167:  uint16(1),
	168:  uint16(sym_unique),
	169:  uint16(57),
	170:  uint16(1),
	171:  uint16(sym_qldoc),
	172:  uint16(56),
	173:  uint16(1),
	174:  uint16(sym_reprType),
	175:  uint16(67),
	176:  uint16(1),
	177:  uint16(sym_column),
	178:  uint16(51),
	179:  uint16(5),
	180:  uint16(sym_int),
	181:  uint16(sym_float),
	182:  uint16(sym_boolean),
	183:  uint16(sym_date),
	184:  uint16(sym_string),
	185:  uint16(8),
	186:  uint16(3),
	187:  uint16(1),
	188:  uint16(sym_line_comment),
	189:  uint16(5),
	190:  uint16(1),
	191:  uint16(sym_block_comment),
	192:  uint16(53),
	193:  uint16(1),
	194:  uint16(sym_varchar),
	195:  uint16(55),
	196:  uint16(1),
	197:  uint16(sym_unique),
	198:  uint16(57),
	199:  uint16(1),
	200:  uint16(sym_qldoc),
	201:  uint16(56),
	202:  uint16(1),
	203:  uint16(sym_reprType),
	204:  uint16(73),
	205:  uint16(1),
	206:  uint16(sym_column),
	207:  uint16(51),
	208:  uint16(5),
	209:  uint16(sym_int),
	210:  uint16(sym_float),
	211:  uint16(sym_boolean),
	212:  uint16(sym_date),
	213:  uint16(sym_string),
	214:  uint16(7),
	215:  uint16(3),
	216:  uint16(1),
	217:  uint16(sym_line_comment),
	218:  uint16(5),
	219:  uint16(1),
	220:  uint16(sym_block_comment),
	221:  uint16(49),
	222:  uint16(1),
	223:  uint16(anon_sym_PIPE),
	224:  uint16(63),
	225:  uint16(1),
	226:  uint16(anon_sym_SEMI),
	227:  uint16(9),
	228:  uint16(1),
	229:  uint16(aux_sym_unionDecl_repeat1),
	230:  uint16(61),
	231:  uint16(2),
	232:  uint16(anon_sym_case),
	233:  uint16(sym__lower_id),
	234:  uint16(59),
	235:  uint16(5),
	237:  uint16(anon_sym_POUND),
	238:  uint16(sym__upper_id),
	239:  uint16(sym_dbtype),
	240:  uint16(sym_qldoc),
	241:  uint16(6),
	242:  uint16(3),
	243:  uint16(1),
	244:  uint16(sym_line_comment),
	245:  uint16(5),
	246:  uint16(1),
	247:  uint16(sym_block_comment),
	248:  uint16(69),
	249:  uint16(1),
	250:  uint16(anon_sym_PIPE),
	251:  uint16(9),
	252:  uint16(1),
	253:  uint16(aux_sym_unionDecl_repeat1),
	254:  uint16(67),
	255:  uint16(2),
	256:  uint16(anon_sym_case),
	257:  uint16(sym__lower_id),
	258:  uint16(65),
	259:  uint16(6),
	261:  uint16(anon_sym_SEMI),
	262:  uint16(anon_sym_POUND),
	263:  uint16(sym__upper_id),
	264:  uint16(sym_dbtype),
	265:  uint16(sym_qldoc),
	266:  uint16(7),
	267:  uint16(3),
	268:  uint16(1),
	269:  uint16(sym_line_comment),
	270:  uint16(5),
	271:  uint16(1),
	272:  uint16(sym_block_comment),
	273:  uint16(76),
	274:  uint16(1),
	275:  uint16(anon_sym_SEMI),
	276:  uint16(78),
	277:  uint16(1),
	278:  uint16(anon_sym_PIPE),
	279:  uint16(12),
	280:  uint16(1),
	281:  uint16(aux_sym_caseDecl_repeat1),
	282:  uint16(74),
	283:  uint16(2),
	284:  uint16(anon_sym_case),
	285:  uint16(sym__lower_id),
	286:  uint16(72),
	287:  uint16(5),
	289:  uint16(anon_sym_POUND),
	290:  uint16(sym__upper_id),
	291:  uint16(sym_dbtype),
	292:  uint16(sym_qldoc),
	293:  uint16(7),
	294:  uint16(3),
	295:  uint16(1),
	296:  uint16(sym_line_comment),
	297:  uint16(5),
	298:  uint16(1),
	299:  uint16(sym_block_comment),
	300:  uint16(78),
	301:  uint16(1),
	302:  uint16(anon_sym_PIPE),
	303:  uint16(84),
	304:  uint16(1),
	305:  uint16(anon_sym_SEMI),
	306:  uint16(10),
	307:  uint16(1),
	308:  uint16(aux_sym_caseDecl_repeat1),
	309:  uint16(82),
	310:  uint16(2),
	311:  uint16(anon_sym_case),
	312:  uint16(sym__lower_id),
	313:  uint16(80),
	314:  uint16(5),
	316:  uint16(anon_sym_POUND),
	317:  uint16(sym__upper_id),
	318:  uint16(sym_dbtype),
	319:  uint16(sym_qldoc),
	320:  uint16(6),
	321:  uint16(3),
	322:  uint16(1),
	323:  uint16(sym_line_comment),
	324:  uint16(5),
	325:  uint16(1),
	326:  uint16(sym_block_comment),
	327:  uint16(90),
	328:  uint16(1),
	329:  uint16(anon_sym_PIPE),
	330:  uint16(12),
	331:  uint16(1),
	332:  uint16(aux_sym_caseDecl_repeat1),
	333:  uint16(88),
	334:  uint16(2),
	335:  uint16(anon_sym_case),
	336:  uint16(sym__lower_id),
	337:  uint16(86),
	338:  uint16(6),
	340:  uint16(anon_sym_SEMI),
	341:  uint16(anon_sym_POUND),
	342:  uint16(sym__upper_id),
	343:  uint16(sym_dbtype),
	344:  uint16(sym_qldoc),
	345:  uint16(4),
	346:  uint16(3),
	347:  uint16(1),
	348:  uint16(sym_line_comment),
	349:  uint16(5),
	350:  uint16(1),
	351:  uint16(sym_block_comment),
	352:  uint16(67),
	353:  uint16(2),
	354:  uint16(anon_sym_case),
	355:  uint16(sym__lower_id),
	356:  uint16(65),
	357:  uint16(7),
	359:  uint16(anon_sym_SEMI),
	360:  uint16(anon_sym_POUND),
	361:  uint16(anon_sym_PIPE),
	362:  uint16(sym__upper_id),
	363:  uint16(sym_dbtype),
	364:  uint16(sym_qldoc),
	365:  uint16(4),
	366:  uint16(3),
	367:  uint16(1),
	368:  uint16(sym_line_comment),
	369:  uint16(5),
	370:  uint16(1),
	371:  uint16(sym_block_comment),
	372:  uint16(95),
	373:  uint16(2),
	374:  uint16(anon_sym_case),
	375:  uint16(sym__lower_id),
	376:  uint16(93),
	377:  uint16(7),
	379:  uint16(anon_sym_SEMI),
	380:  uint16(anon_sym_POUND),
	381:  uint16(anon_sym_PIPE),
	382:  uint16(sym__upper_id),
	383:  uint16(sym_dbtype),
	384:  uint16(sym_qldoc),
	385:  uint16(4),
	386:  uint16(3),
	387:  uint16(1),
	388:  uint16(sym_line_comment),
	389:  uint16(5),
	390:  uint16(1),
	391:  uint16(sym_block_comment),
	392:  uint16(88),
	393:  uint16(2),
	394:  uint16(anon_sym_case),
	395:  uint16(sym__lower_id),
	396:  uint16(86),
	397:  uint16(7),
	399:  uint16(anon_sym_SEMI),
	400:  uint16(anon_sym_POUND),
	401:  uint16(anon_sym_PIPE),
	402:  uint16(sym__upper_id),
	403:  uint16(sym_dbtype),
	404:  uint16(sym_qldoc),
	405:  uint16(4),
	406:  uint16(3),
	407:  uint16(1),
	408:  uint16(sym_line_comment),
	409:  uint16(5),
	410:  uint16(1),
	411:  uint16(sym_block_comment),
	412:  uint16(99),
	413:  uint16(2),
	414:  uint16(anon_sym_case),
	415:  uint16(sym__lower_id),
	416:  uint16(97),
	417:  uint16(7),
	419:  uint16(anon_sym_SEMI),
	420:  uint16(anon_sym_POUND),
	421:  uint16(anon_sym_PIPE),
	422:  uint16(sym__upper_id),
	423:  uint16(sym_dbtype),
	424:  uint16(sym_qldoc),
	425:  uint16(5),
	426:  uint16(3),
	427:  uint16(1),
	428:  uint16(sym_line_comment),
	429:  uint16(5),
	430:  uint16(1),
	431:  uint16(sym_block_comment),
	432:  uint16(105),
	433:  uint16(1),
	434:  uint16(anon_sym_SEMI),
	435:  uint16(103),
	436:  uint16(2),
	437:  uint16(anon_sym_case),
	438:  uint16(sym__lower_id),
	439:  uint16(101),
	440:  uint16(5),
	442:  uint16(anon_sym_POUND),
	443:  uint16(sym__upper_id),
	444:  uint16(sym_dbtype),
	445:  uint16(sym_qldoc),
	446:  uint16(5),
	447:  uint16(3),
	448:  uint16(1),
	449:  uint16(sym_line_comment),
	450:  uint16(5),
	451:  uint16(1),
	452:  uint16(sym_block_comment),
	453:  uint16(111),
	454:  uint16(1),
	455:  uint16(anon_sym_SEMI),
	456:  uint16(109),
	457:  uint16(2),
	458:  uint16(anon_sym_case),
	459:  uint16(sym__lower_id),
	460:  uint16(107),
	461:  uint16(5),
	463:  uint16(anon_sym_POUND),
	464:  uint16(sym__upper_id),
	465:  uint16(sym_dbtype),
	466:  uint16(sym_qldoc),
	467:  uint16(7),
	468:  uint16(11),
	469:  uint16(1),
	470:  uint16(anon_sym_POUND),
	471:  uint16(46),
	472:  uint16(1),
	473:  uint16(sym_argsAnnotation),
	474:  uint16(82),
	475:  uint16(1),
	476:  uint16(sym_tableName),
	477:  uint16(88),
	478:  uint16(1),
	479:  uint16(sym_simpleId),
	480:  uint16(3),
	481:  uint16(2),
	482:  uint16(sym_line_comment),
	483:  uint16(sym_block_comment),
	484:  uint16(15),
	485:  uint16(2),
	486:  uint16(sym__lower_id),
	487:  uint16(sym__upper_id),
	488:  uint16(38),
	489:  uint16(2),
	490:  uint16(sym_annotation),
	491:  uint16(aux_sym_table_repeat1),
	492:  uint16(5),
	493:  uint16(3),
	494:  uint16(1),
	495:  uint16(sym_line_comment),
	496:  uint16(5),
	497:  uint16(1),
	498:  uint16(sym_block_comment),
	499:  uint16(117),
	500:  uint16(1),
	501:  uint16(anon_sym_SEMI),
	502:  uint16(115),
	503:  uint16(2),
	504:  uint16(anon_sym_case),
	505:  uint16(sym__lower_id),
	506:  uint16(113),
	507:  uint16(5),
	509:  uint16(anon_sym_POUND),
	510:  uint16(sym__upper_id),
	511:  uint16(sym_dbtype),
	512:  uint16(sym_qldoc),
	513:  uint16(5),
	514:  uint16(3),
	515:  uint16(1),
	516:  uint16(sym_line_comment),
	517:  uint16(5),
	518:  uint16(1),
	519:  uint16(sym_block_comment),
	520:  uint16(123),
	521:  uint16(1),
	522:  uint16(anon_sym_SEMI),
	523:  uint16(121),
	524:  uint16(2),
	525:  uint16(anon_sym_case),
	526:  uint16(sym__lower_id),
	527:  uint16(119),
	528:  uint16(5),
	530:  uint16(anon_sym_POUND),
	531:  uint16(sym__upper_id),
	532:  uint16(sym_dbtype),
	533:  uint16(sym_qldoc),
	534:  uint16(5),
	535:  uint16(53),
	536:  uint16(1),
	537:  uint16(sym_varchar),
	538:  uint16(125),
	539:  uint16(1),
	540:  uint16(sym_unique),
	541:  uint16(63),
	542:  uint16(1),
	543:  uint16(sym_reprType),
	544:  uint16(3),
	545:  uint16(2),
	546:  uint16(sym_line_comment),
	547:  uint16(sym_block_comment),
	548:  uint16(51),
	549:  uint16(5),
	550:  uint16(sym_int),
	551:  uint16(sym_float),
	552:  uint16(sym_boolean),
	553:  uint16(sym_date),
	554:  uint16(sym_string),
	555:  uint16(4),
	556:  uint16(3),
	557:  uint16(1),
	558:  uint16(sym_line_comment),
	559:  uint16(5),
	560:  uint16(1),
	561:  uint16(sym_block_comment),
	562:  uint16(129),
	563:  uint16(2),
	564:  uint16(anon_sym_case),
	565:  uint16(sym__lower_id),
	566:  uint16(127),
	567:  uint16(5),
	569:  uint16(anon_sym_POUND),
	570:  uint16(sym__upper_id),
	571:  uint16(sym_dbtype),
	572:  uint16(sym_qldoc),
	573:  uint16(3),
	574:  uint16(55),
	575:  uint16(1),
	576:  uint16(sym_colType),
	577:  uint16(3),
	578:  uint16(2),
	579:  uint16(sym_line_comment),
	580:  uint16(sym_block_comment),
	581:  uint16(131),
	582:  uint16(6),
	583:  uint16(sym_int),
	584:  uint16(sym_float),
	585:  uint16(sym_boolean),
	586:  uint16(sym_date),
	587:  uint16(sym_string),
	588:  uint16(sym_dbtype),
	589:  uint16(4),
	590:  uint16(3),
	591:  uint16(1),
	592:  uint16(sym_line_comment),
	593:  uint16(5),
	594:  uint16(1),
	595:  uint16(sym_block_comment),
	596:  uint16(109),
	597:  uint16(2),
	598:  uint16(anon_sym_case),
	599:  uint16(sym__lower_id),
	600:  uint16(107),
	601:  uint16(5),
	603:  uint16(anon_sym_POUND),
	604:  uint16(sym__upper_id),
	605:  uint16(sym_dbtype),
	606:  uint16(sym_qldoc),
	607:  uint16(4),
	608:  uint16(3),
	609:  uint16(1),
	610:  uint16(sym_line_comment),
	611:  uint16(5),
	612:  uint16(1),
	613:  uint16(sym_block_comment),
	614:  uint16(61),
	615:  uint16(2),
	616:  uint16(anon_sym_case),
	617:  uint16(sym__lower_id),
	618:  uint16(59),
	619:  uint16(5),
	621:  uint16(anon_sym_POUND),
	622:  uint16(sym__upper_id),
	623:  uint16(sym_dbtype),
	624:  uint16(sym_qldoc),
	625:  uint16(3),
	626:  uint16(60),
	627:  uint16(1),
	628:  uint16(sym_colType),
	629:  uint16(3),
	630:  uint16(2),
	631:  uint16(sym_line_comment),
	632:  uint16(sym_block_comment),
	633:  uint16(131),
	634:  uint16(6),
	635:  uint16(sym_int),
	636:  uint16(sym_float),
	637:  uint16(sym_boolean),
	638:  uint16(sym_date),
	639:  uint16(sym_string),
	640:  uint16(sym_dbtype),
	641:  uint16(3),
	642:  uint16(61),
	643:  uint16(1),
	644:  uint16(sym_colType),
	645:  uint16(3),
	646:  uint16(2),
	647:  uint16(sym_line_comment),
	648:  uint16(sym_block_comment),
	649:  uint16(131),
	650:  uint16(6),
	651:  uint16(sym_int),
	652:  uint16(sym_float),
	653:  uint16(sym_boolean),
	654:  uint16(sym_date),
	655:  uint16(sym_string),
	656:  uint16(sym_dbtype),
	657:  uint16(4),
	658:  uint16(3),
	659:  uint16(1),
	660:  uint16(sym_line_comment),
	661:  uint16(5),
	662:  uint16(1),
	663:  uint16(sym_block_comment),
	664:  uint16(135),
	665:  uint16(2),
	666:  uint16(anon_sym_case),
	667:  uint16(sym__lower_id),
	668:  uint16(133),
	669:  uint16(5),
	671:  uint16(anon_sym_POUND),
	672:  uint16(sym__upper_id),
	673:  uint16(sym_dbtype),
	674:  uint16(sym_qldoc),
	675:  uint16(4),
	676:  uint16(3),
	677:  uint16(1),
	678:  uint16(sym_line_comment),
	679:  uint16(5),
	680:  uint16(1),
	681:  uint16(sym_block_comment),
	682:  uint16(121),
	683:  uint16(2),
	684:  uint16(anon_sym_case),
	685:  uint16(sym__lower_id),
	686:  uint16(119),
	687:  uint16(5),
	689:  uint16(anon_sym_POUND),
	690:  uint16(sym__upper_id),
	691:  uint16(sym_dbtype),
	692:  uint16(sym_qldoc),
	693:  uint16(4),
	694:  uint16(3),
	695:  uint16(1),
	696:  uint16(sym_line_comment),
	697:  uint16(5),
	698:  uint16(1),
	699:  uint16(sym_block_comment),
	700:  uint16(74),
	701:  uint16(2),
	702:  uint16(anon_sym_case),
	703:  uint16(sym__lower_id),
	704:  uint16(72),
	705:  uint16(5),
	707:  uint16(anon_sym_POUND),
	708:  uint16(sym__upper_id),
	709:  uint16(sym_dbtype),
	710:  uint16(sym_qldoc),
	711:  uint16(4),
	712:  uint16(3),
	713:  uint16(1),
	714:  uint16(sym_line_comment),
	715:  uint16(5),
	716:  uint16(1),
	717:  uint16(sym_block_comment),
	718:  uint16(139),
	719:  uint16(2),
	720:  uint16(anon_sym_case),
	721:  uint16(sym__lower_id),
	722:  uint16(137),
	723:  uint16(5),
	725:  uint16(anon_sym_POUND),
	726:  uint16(sym__upper_id),
	727:  uint16(sym_dbtype),
	728:  uint16(sym_qldoc),
	729:  uint16(3),
	730:  uint16(66),
	731:  uint16(1),
	732:  uint16(sym_colType),
	733:  uint16(3),
	734:  uint16(2),
	735:  uint16(sym_line_comment),
	736:  uint16(sym_block_comment),
	737:  uint16(131),
	738:  uint16(6),
	739:  uint16(sym_int),
	740:  uint16(sym_float),
	741:  uint16(sym_boolean),
	742:  uint16(sym_date),
	743:  uint16(sym_string),
	744:  uint16(sym_dbtype),
	745:  uint16(4),
	746:  uint16(3),
	747:  uint16(1),
	748:  uint16(sym_line_comment),
	749:  uint16(5),
	750:  uint16(1),
	751:  uint16(sym_block_comment),
	752:  uint16(143),
	753:  uint16(2),
	754:  uint16(anon_sym_case),
	755:  uint16(sym__lower_id),
	756:  uint16(141),
	757:  uint16(5),
	759:  uint16(anon_sym_POUND),
	760:  uint16(sym__upper_id),
	761:  uint16(sym_dbtype),
	762:  uint16(sym_qldoc),
	763:  uint16(4),
	764:  uint16(53),
	765:  uint16(1),
	766:  uint16(sym_varchar),
	767:  uint16(47),
	768:  uint16(1),
	769:  uint16(sym_reprType),
	770:  uint16(3),
	771:  uint16(2),
	772:  uint16(sym_line_comment),
	773:  uint16(sym_block_comment),
	774:  uint16(51),
	775:  uint16(5),
	776:  uint16(sym_int),
	777:  uint16(sym_float),
	778:  uint16(sym_boolean),
	779:  uint16(sym_date),
	780:  uint16(sym_string),
	781:  uint16(4),
	782:  uint16(3),
	783:  uint16(1),
	784:  uint16(sym_line_comment),
	785:  uint16(5),
	786:  uint16(1),
	787:  uint16(sym_block_comment),
	788:  uint16(147),
	789:  uint16(2),
	790:  uint16(anon_sym_case),
	791:  uint16(sym__lower_id),
	792:  uint16(145),
	793:  uint16(5),
	795:  uint16(anon_sym_POUND),
	796:  uint16(sym__upper_id),
	797:  uint16(sym_dbtype),
	798:  uint16(sym_qldoc),
	799:  uint16(4),
	800:  uint16(53),
	801:  uint16(1),
	802:  uint16(sym_varchar),
	803:  uint16(65),
	804:  uint16(1),
	805:  uint16(sym_reprType),
	806:  uint16(3),
	807:  uint16(2),
	808:  uint16(sym_line_comment),
	809:  uint16(sym_block_comment),
	810:  uint16(51),
	811:  uint16(5),
	812:  uint16(sym_int),
	813:  uint16(sym_float),
	814:  uint16(sym_boolean),
	815:  uint16(sym_date),
	816:  uint16(sym_string),
	817:  uint16(5),
	818:  uint16(151),
	819:  uint16(1),
	820:  uint16(anon_sym_POUND),
	821:  uint16(46),
	822:  uint16(1),
	823:  uint16(sym_argsAnnotation),
	824:  uint16(3),
	825:  uint16(2),
	826:  uint16(sym_line_comment),
	827:  uint16(sym_block_comment),
	828:  uint16(149),
	829:  uint16(2),
	830:  uint16(sym__lower_id),
	831:  uint16(sym__upper_id),
	832:  uint16(38),
	833:  uint16(2),
	834:  uint16(sym_annotation),
	835:  uint16(aux_sym_table_repeat1),
	836:  uint16(2),
	837:  uint16(3),
	838:  uint16(2),
	839:  uint16(sym_line_comment),
	840:  uint16(sym_block_comment),
	841:  uint16(154),
	842:  uint16(5),
	843:  uint16(anon_sym_LPAREN),
	844:  uint16(anon_sym_COMMA),
	845:  uint16(anon_sym_RBRACK),
	846:  uint16(anon_sym_COLON),
	847:  uint16(anon_sym_of),
	848:  uint16(3),
	849:  uint16(158),
	850:  uint16(1),
	851:  uint16(anon_sym_LBRACK),
	852:  uint16(3),
	853:  uint16(2),
	854:  uint16(sym_line_comment),
	855:  uint16(sym_block_comment),
	856:  uint16(156),
	857:  uint16(3),
	858:  uint16(anon_sym_POUND),
	859:  uint16(sym__lower_id),
	860:  uint16(sym__upper_id),
	861:  uint16(4),
	862:  uint16(160),
	863:  uint16(1),
	864:  uint16(anon_sym_RBRACK),
	865:  uint16(59),
	866:  uint16(1),
	867:  uint16(sym_simpleId),
	868:  uint16(3),
	869:  uint16(2),
	870:  uint16(sym_line_comment),
	871:  uint16(sym_block_comment),
	872:  uint16(15),
	873:  uint16(2),
	874:  uint16(sym__lower_id),
	875:  uint16(sym__upper_id),
	876:  uint16(2),
	877:  uint16(3),
	878:  uint16(2),
	879:  uint16(sym_line_comment),
	880:  uint16(sym_block_comment),
	881:  uint16(162),
	882:  uint16(4),
	883:  uint16(anon_sym_POUND),
	884:  uint16(anon_sym_LBRACK),
	885:  uint16(sym__lower_id),
	886:  uint16(sym__upper_id),
	887:  uint16(2),
	888:  uint16(3),
	889:  uint16(2),
	890:  uint16(sym_line_comment),
	891:  uint16(sym_block_comment),
	892:  uint16(164),
	893:  uint16(3),
	894:  uint16(anon_sym_POUND),
	895:  uint16(sym__lower_id),
	896:  uint16(sym__upper_id),
	897:  uint16(4),
	898:  uint16(166),
	899:  uint16(1),
	900:  uint16(anon_sym_COMMA),
	901:  uint16(168),
	902:  uint16(1),
	903:  uint16(anon_sym_RBRACK),
	904:  uint16(52),
	905:  uint16(1),
	906:  uint16(aux_sym_argsAnnotation_repeat1),
	907:  uint16(3),
	908:  uint16(2),
	909:  uint16(sym_line_comment),
	910:  uint16(sym_block_comment),
	911:  uint16(5),
	912:  uint16(3),
	913:  uint16(1),
	914:  uint16(sym_line_comment),
	915:  uint16(5),
	916:  uint16(1),
	917:  uint16(sym_block_comment),
	918:  uint16(170),
	919:  uint16(1),
	920:  uint16(sym_integer),
	921:  uint16(172),
	922:  uint16(1),
	923:  uint16(sym_qldoc),
	924:  uint16(11),
	925:  uint16(1),
	926:  uint16(sym_branch),
	927:  uint16(2),
	928:  uint16(3),
	929:  uint16(2),
	930:  uint16(sym_line_comment),
	931:  uint16(sym_block_comment),
	932:  uint16(174),
	933:  uint16(3),
	934:  uint16(anon_sym_POUND),
	935:  uint16(sym__lower_id),
	936:  uint16(sym__upper_id),
	937:  uint16(3),
	938:  uint16(89),
	939:  uint16(1),
	940:  uint16(sym_simpleId),
	941:  uint16(3),
	942:  uint16(2),
	943:  uint16(sym_line_comment),
	944:  uint16(sym_block_comment),
	945:  uint16(15),
	946:  uint16(2),
	947:  uint16(sym__lower_id),
	948:  uint16(sym__upper_id),
	949:  uint16(4),
	950:  uint16(176),
	951:  uint16(1),
	952:  uint16(anon_sym_COMMA),
	953:  uint16(178),
	954:  uint16(1),
	955:  uint16(anon_sym_RPAREN),
	956:  uint16(51),
	957:  uint16(1),
	958:  uint16(aux_sym_table_repeat2),
	959:  uint16(3),
	960:  uint16(2),
	961:  uint16(sym_line_comment),
	962:  uint16(sym_block_comment),
	963:  uint16(3),
	964:  uint16(85),
	965:  uint16(1),
	966:  uint16(sym_simpleId),
	967:  uint16(3),
	968:  uint16(2),
	969:  uint16(sym_line_comment),
	970:  uint16(sym_block_comment),
	971:  uint16(15),
	972:  uint16(2),
	973:  uint16(sym__lower_id),
	974:  uint16(sym__upper_id),
	975:  uint16(4),
	976:  uint16(180),
	977:  uint16(1),
	978:  uint16(anon_sym_COMMA),
	979:  uint16(183),
	980:  uint16(1),
	981:  uint16(anon_sym_RPAREN),
	982:  uint16(50),
	983:  uint16(1),
	984:  uint16(aux_sym_table_repeat2),
	985:  uint16(3),
	986:  uint16(2),
	987:  uint16(sym_line_comment),
	988:  uint16(sym_block_comment),
	989:  uint16(4),
	990:  uint16(176),
	991:  uint16(1),
	992:  uint16(anon_sym_COMMA),
	993:  uint16(185),
	994:  uint16(1),
	995:  uint16(anon_sym_RPAREN),
	996:  uint16(50),
	997:  uint16(1),
	998:  uint16(aux_sym_table_repeat2),
	999:  uint16(3),
	1000: uint16(2),
	1001: uint16(sym_line_comment),
	1002: uint16(sym_block_comment),
	1003: uint16(4),
	1004: uint16(187),
	1005: uint16(1),
	1006: uint16(anon_sym_COMMA),
	1007: uint16(190),
	1008: uint16(1),
	1009: uint16(anon_sym_RBRACK),
	1010: uint16(52),
	1011: uint16(1),
	1012: uint16(aux_sym_argsAnnotation_repeat1),
	1013: uint16(3),
	1014: uint16(2),
	1015: uint16(sym_line_comment),
	1016: uint16(sym_block_comment),
	1017: uint16(4),
	1018: uint16(176),
	1019: uint16(1),
	1020: uint16(anon_sym_COMMA),
	1021: uint16(192),
	1022: uint16(1),
	1023: uint16(anon_sym_RPAREN),
	1024: uint16(50),
	1025: uint16(1),
	1026: uint16(aux_sym_table_repeat2),
	1027: uint16(3),
	1028: uint16(2),
	1029: uint16(sym_line_comment),
	1030: uint16(sym_block_comment),
	1031: uint16(2),
	1032: uint16(3),
	1033: uint16(2),
	1034: uint16(sym_line_comment),
	1035: uint16(sym_block_comment),
	1036: uint16(194),
	1037: uint16(3),
	1038: uint16(anon_sym_COMMA),
	1039: uint16(anon_sym_RPAREN),
	1040: uint16(sym_ref),
	1041: uint16(3),
	1042: uint16(198),
	1043: uint16(1),
	1044: uint16(sym_ref),
	1045: uint16(3),
	1046: uint16(2),
	1047: uint16(sym_line_comment),
	1048: uint16(sym_block_comment),
	1049: uint16(196),
	1050: uint16(2),
	1051: uint16(anon_sym_COMMA),
	1052: uint16(anon_sym_RPAREN),
	1053: uint16(3),
	1054: uint16(87),
	1055: uint16(1),
	1056: uint16(sym_simpleId),
	1057: uint16(3),
	1058: uint16(2),
	1059: uint16(sym_line_comment),
	1060: uint16(sym_block_comment),
	1061: uint16(15),
	1062: uint16(2),
	1063: uint16(sym__lower_id),
	1064: uint16(sym__upper_id),
	1065: uint16(2),
	1066: uint16(3),
	1067: uint16(2),
	1068: uint16(sym_line_comment),
	1069: uint16(sym_block_comment),
	1070: uint16(200),
	1071: uint16(3),
	1072: uint16(anon_sym_POUND),
	1073: uint16(sym__lower_id),
	1074: uint16(sym__upper_id),
	1075: uint16(5),
	1076: uint16(3),
	1077: uint16(1),
	1078: uint16(sym_line_comment),
	1079: uint16(5),
	1080: uint16(1),
	1081: uint16(sym_block_comment),
	1082: uint16(170),
	1083: uint16(1),
	1084: uint16(sym_integer),
	1085: uint16(172),
	1086: uint16(1),
	1087: uint16(sym_qldoc),
	1088: uint16(15),
	1089: uint16(1),
	1090: uint16(sym_branch),
	1091: uint16(4),
	1092: uint16(166),
	1093: uint16(1),
	1094: uint16(anon_sym_COMMA),
	1095: uint16(202),
	1096: uint16(1),
	1097: uint16(anon_sym_RBRACK),
	1098: uint16(44),
	1099: uint16(1),
	1100: uint16(aux_sym_argsAnnotation_repeat1),
	1101: uint16(3),
	1102: uint16(2),
	1103: uint16(sym_line_comment),
	1104: uint16(sym_block_comment),
	1105: uint16(3),
	1106: uint16(206),
	1107: uint16(1),
	1108: uint16(sym_ref),
	1109: uint16(3),
	1110: uint16(2),
	1111: uint16(sym_line_comment),
	1112: uint16(sym_block_comment),
	1113: uint16(204),
	1114: uint16(2),
	1115: uint16(anon_sym_COMMA),
	1116: uint16(anon_sym_RPAREN),
	1117: uint16(3),
	1118: uint16(210),
	1119: uint16(1),
	1120: uint16(sym_ref),
	1121: uint16(3),
	1122: uint16(2),
	1123: uint16(sym_line_comment),
	1124: uint16(sym_block_comment),
	1125: uint16(208),
	1126: uint16(2),
	1127: uint16(anon_sym_COMMA),
	1128: uint16(anon_sym_RPAREN),
	1129: uint16(3),
	1130: uint16(75),
	1131: uint16(1),
	1132: uint16(sym_simpleId),
	1133: uint16(3),
	1134: uint16(2),
	1135: uint16(sym_line_comment),
	1136: uint16(sym_block_comment),
	1137: uint16(15),
	1138: uint16(2),
	1139: uint16(sym__lower_id),
	1140: uint16(sym__upper_id),
	1141: uint16(3),
	1142: uint16(96),
	1143: uint16(1),
	1144: uint16(sym_simpleId),
	1145: uint16(3),
	1146: uint16(2),
	1147: uint16(sym_line_comment),
	1148: uint16(sym_block_comment),
	1149: uint16(15),
	1150: uint16(2),
	1151: uint16(sym__lower_id),
	1152: uint16(sym__upper_id),
	1153: uint16(2),
	1154: uint16(3),
	1155: uint16(2),
	1156: uint16(sym_line_comment),
	1157: uint16(sym_block_comment),
	1158: uint16(212),
	1159: uint16(3),
	1160: uint16(anon_sym_POUND),
	1161: uint16(sym__lower_id),
	1162: uint16(sym__upper_id),
	1163: uint16(3),
	1164: uint16(94),
	1165: uint16(1),
	1166: uint16(sym_simpleId),
	1167: uint16(3),
	1168: uint16(2),
	1169: uint16(sym_line_comment),
	1170: uint16(sym_block_comment),
	1171: uint16(15),
	1172: uint16(2),
	1173: uint16(sym__lower_id),
	1174: uint16(sym__upper_id),
	1175: uint16(3),
	1176: uint16(216),
	1177: uint16(1),
	1178: uint16(sym_ref),
	1179: uint16(3),
	1180: uint16(2),
	1181: uint16(sym_line_comment),
	1182: uint16(sym_block_comment),
	1183: uint16(214),
	1184: uint16(2),
	1185: uint16(anon_sym_COMMA),
	1186: uint16(anon_sym_RPAREN),
	1187: uint16(4),
	1188: uint16(176),
	1189: uint16(1),
	1190: uint16(anon_sym_COMMA),
	1191: uint16(218),
	1192: uint16(1),
	1193: uint16(anon_sym_RPAREN),
	1194: uint16(53),
	1195: uint16(1),
	1196: uint16(aux_sym_table_repeat2),
	1197: uint16(3),
	1198: uint16(2),
	1199: uint16(sym_line_comment),
	1200: uint16(sym_block_comment),
	1201: uint16(2),
	1202: uint16(3),
	1203: uint16(2),
	1204: uint16(sym_line_comment),
	1205: uint16(sym_block_comment),
	1206: uint16(220),
	1207: uint16(2),
	1208: uint16(anon_sym_COMMA),
	1209: uint16(anon_sym_RPAREN),
	1210: uint16(2),
	1211: uint16(3),
	1212: uint16(2),
	1213: uint16(sym_line_comment),
	1214: uint16(sym_block_comment),
	1215: uint16(222),
	1216: uint16(2),
	1217: uint16(anon_sym_COMMA),
	1218: uint16(anon_sym_RPAREN),
	1219: uint16(2),
	1220: uint16(3),
	1221: uint16(2),
	1222: uint16(sym_line_comment),
	1223: uint16(sym_block_comment),
	1224: uint16(224),
	1225: uint16(2),
	1226: uint16(sym__lower_id),
	1227: uint16(sym__upper_id),
	1228: uint16(2),
	1229: uint16(3),
	1230: uint16(2),
	1231: uint16(sym_line_comment),
	1232: uint16(sym_block_comment),
	1233: uint16(226),
	1234: uint16(2),
	1235: uint16(sym__lower_id),
	1236: uint16(sym__upper_id),
	1237: uint16(3),
	1238: uint16(228),
	1239: uint16(1),
	1240: uint16(sym__lower_id),
	1241: uint16(40),
	1242: uint16(1),
	1243: uint16(sym_annotName),
	1244: uint16(3),
	1245: uint16(2),
	1246: uint16(sym_line_comment),
	1247: uint16(sym_block_comment),
	1248: uint16(2),
	1249: uint16(3),
	1250: uint16(2),
	1251: uint16(sym_line_comment),
	1252: uint16(sym_block_comment),
	1253: uint16(183),
	1254: uint16(2),
	1255: uint16(anon_sym_COMMA),
	1256: uint16(anon_sym_RPAREN),
	1257: uint16(2),
	1258: uint16(3),
	1259: uint16(2),
	1260: uint16(sym_line_comment),
	1261: uint16(sym_block_comment),
	1262: uint16(230),
	1263: uint16(2),
	1264: uint16(anon_sym_COMMA),
	1265: uint16(anon_sym_RPAREN),
	1266: uint16(2),
	1267: uint16(3),
	1268: uint16(2),
	1269: uint16(sym_line_comment),
	1270: uint16(sym_block_comment),
	1271: uint16(190),
	1272: uint16(2),
	1273: uint16(anon_sym_COMMA),
	1274: uint16(anon_sym_RBRACK),
	1275: uint16(2),
	1276: uint16(3),
	1277: uint16(2),
	1278: uint16(sym_line_comment),
	1279: uint16(sym_block_comment),
	1280: uint16(232),
	1281: uint16(2),
	1282: uint16(anon_sym_COMMA),
	1283: uint16(anon_sym_RPAREN),
	1284: uint16(2),
	1285: uint16(234),
	1286: uint16(1),
	1287: uint16(sym_dbtype),
	1288: uint16(3),
	1289: uint16(2),
	1290: uint16(sym_line_comment),
	1291: uint16(sym_block_comment),
	1292: uint16(2),
	1293: uint16(236),
	1294: uint16(1),
	1295: uint16(sym_dbtype),
	1296: uint16(3),
	1297: uint16(2),
	1298: uint16(sym_line_comment),
	1299: uint16(sym_block_comment),
	1300: uint16(2),
	1301: uint16(238),
	1302: uint16(1),
	1303: uint16(sym_dbtype),
	1304: uint16(3),
	1305: uint16(2),
	1306: uint16(sym_line_comment),
	1307: uint16(sym_block_comment),
	1308: uint16(2),
	1309: uint16(240),
	1310: uint16(1),
	1311: uint16(anon_sym_EQ),
	1312: uint16(3),
	1313: uint16(2),
	1314: uint16(sym_line_comment),
	1315: uint16(sym_block_comment),
	1316: uint16(2),
	1317: uint16(242),
	1318: uint16(1),
	1320: uint16(3),
	1321: uint16(2),
	1322: uint16(sym_line_comment),
	1323: uint16(sym_block_comment),
	1324: uint16(2),
	1325: uint16(244),
	1326: uint16(1),
	1327: uint16(anon_sym_LPAREN),
	1328: uint16(3),
	1329: uint16(2),
	1330: uint16(sym_line_comment),
	1331: uint16(sym_block_comment),
	1332: uint16(2),
	1333: uint16(246),
	1334: uint16(1),
	1335: uint16(anon_sym_EQ),
	1336: uint16(3),
	1337: uint16(2),
	1338: uint16(sym_line_comment),
	1339: uint16(sym_block_comment),
	1340: uint16(2),
	1341: uint16(248),
	1342: uint16(1),
	1343: uint16(sym_integer),
	1344: uint16(3),
	1345: uint16(2),
	1346: uint16(sym_line_comment),
	1347: uint16(sym_block_comment),
	1348: uint16(2),
	1349: uint16(250),
	1350: uint16(1),
	1351: uint16(anon_sym_of),
	1352: uint16(3),
	1353: uint16(2),
	1354: uint16(sym_line_comment),
	1355: uint16(sym_block_comment),
	1356: uint16(2),
	1357: uint16(252),
	1358: uint16(1),
	1359: uint16(anon_sym_LPAREN),
	1360: uint16(3),
	1361: uint16(2),
	1362: uint16(sym_line_comment),
	1363: uint16(sym_block_comment),
	1364: uint16(2),
	1365: uint16(254),
	1366: uint16(1),
	1367: uint16(anon_sym_COLON),
	1368: uint16(3),
	1369: uint16(2),
	1370: uint16(sym_line_comment),
	1371: uint16(sym_block_comment),
	1372: uint16(2),
	1373: uint16(256),
	1374: uint16(1),
	1375: uint16(anon_sym_LPAREN),
	1376: uint16(3),
	1377: uint16(2),
	1378: uint16(sym_line_comment),
	1379: uint16(sym_block_comment),
	1380: uint16(2),
	1381: uint16(258),
	1382: uint16(1),
	1383: uint16(anon_sym_COLON),
	1384: uint16(3),
	1385: uint16(2),
	1386: uint16(sym_line_comment),
	1387: uint16(sym_block_comment),
	1388: uint16(2),
	1389: uint16(260),
	1390: uint16(1),
	1391: uint16(sym_dbtype),
	1392: uint16(3),
	1393: uint16(2),
	1394: uint16(sym_line_comment),
	1395: uint16(sym_block_comment),
	1396: uint16(2),
	1397: uint16(262),
	1398: uint16(1),
	1399: uint16(anon_sym_RPAREN),
	1400: uint16(3),
	1401: uint16(2),
	1402: uint16(sym_line_comment),
	1403: uint16(sym_block_comment),
	1404: uint16(2),
	1405: uint16(264),
	1406: uint16(1),
	1407: uint16(sym_integer),
	1408: uint16(3),
	1409: uint16(2),
	1410: uint16(sym_line_comment),
	1411: uint16(sym_block_comment),
	1412: uint16(2),
	1413: uint16(266),
	1414: uint16(1),
	1415: uint16(anon_sym_LPAREN),
	1416: uint16(3),
	1417: uint16(2),
	1418: uint16(sym_line_comment),
	1419: uint16(sym_block_comment),
	1420: uint16(2),
	1421: uint16(268),
	1422: uint16(1),
	1423: uint16(anon_sym_COLON),
	1424: uint16(3),
	1425: uint16(2),
	1426: uint16(sym_line_comment),
	1427: uint16(sym_block_comment),
	1428: uint16(2),
	1429: uint16(270),
	1430: uint16(1),
	1431: uint16(sym_dbtype),
	1432: uint16(3),
	1433: uint16(2),
	1434: uint16(sym_line_comment),
	1435: uint16(sym_block_comment),
	1436: uint16(2),
	1437: uint16(272),
	1438: uint16(1),
	1439: uint16(anon_sym_COLON),
	1440: uint16(3),
	1441: uint16(2),
	1442: uint16(sym_line_comment),
	1443: uint16(sym_block_comment),
	1444: uint16(2),
	1445: uint16(274),
	1446: uint16(1),
	1447: uint16(anon_sym_DOT),
	1448: uint16(3),
	1449: uint16(2),
	1450: uint16(sym_line_comment),
	1451: uint16(sym_block_comment),
	1452: uint16(2),
	1453: uint16(276),
	1454: uint16(1),
	1455: uint16(anon_sym_EQ),
	1456: uint16(3),
	1457: uint16(2),
	1458: uint16(sym_line_comment),
	1459: uint16(sym_block_comment),
}

var ts_small_parse_table_map = [97]uint32_t{
	1:  uint32(50),
	2:  uint32(100),
	3:  uint32(127),
	4:  uint32(156),
	5:  uint32(185),
	6:  uint32(214),
	7:  uint32(241),
	8:  uint32(266),
	9:  uint32(293),
	10: uint32(320),
	11: uint32(345),
	12: uint32(365),
	13: uint32(385),
	14: uint32(405),
	15: uint32(425),
	16: uint32(446),
	17: uint32(467),
	18: uint32(492),
	19: uint32(513),
	20: uint32(534),
	21: uint32(555),
	22: uint32(573),
	23: uint32(589),
	24: uint32(607),
	25: uint32(625),
	26: uint32(641),
	27: uint32(657),
	28: uint32(675),
	29: uint32(693),
	30: uint32(711),
	31: uint32(729),
	32: uint32(745),
	33: uint32(763),
	34: uint32(781),
	35: uint32(799),
	36: uint32(817),
	37: uint32(836),
	38: uint32(848),
	39: uint32(861),
	40: uint32(876),
	41: uint32(887),
	42: uint32(897),
	43: uint32(911),
	44: uint32(927),
	45: uint32(937),
	46: uint32(949),
	47: uint32(963),
	48: uint32(975),
	49: uint32(989),
	50: uint32(1003),
	51: uint32(1017),
	52: uint32(1031),
	53: uint32(1041),
	54: uint32(1053),
	55: uint32(1065),
	56: uint32(1075),
	57: uint32(1091),
	58: uint32(1105),
	59: uint32(1117),
	60: uint32(1129),
	61: uint32(1141),
	62: uint32(1153),
	63: uint32(1163),
	64: uint32(1175),
	65: uint32(1187),
	66: uint32(1201),
	67: uint32(1210),
	68: uint32(1219),
	69: uint32(1228),
	70: uint32(1237),
	71: uint32(1248),
	72: uint32(1257),
	73: uint32(1266),
	74: uint32(1275),
	75: uint32(1284),
	76: uint32(1292),
	77: uint32(1300),
	78: uint32(1308),
	79: uint32(1316),
	80: uint32(1324),
	81: uint32(1332),
	82: uint32(1340),
	83: uint32(1348),
	84: uint32(1356),
	85: uint32(1364),
	86: uint32(1372),
	87: uint32(1380),
	88: uint32(1388),
	89: uint32(1396),
	90: uint32(1404),
	91: uint32(1412),
	92: uint32(1420),
	93: uint32(1428),
	94: uint32(1436),
	95: uint32(1444),
	96: uint32(1452),
}

var ts_parse_actions = [278]TSParseActionEntry{
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
		Fcount: uint8(1),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token: uint8(TSParseActionTypeReduce),
		Fsymbol:     uint16(sym_dbscheme),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(98)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dbscheme_repeat1),
	})))),
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
		Fcount: uint8(2),
	}})),
	24: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dbscheme_repeat1),
	})))),
	25: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	26: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	27: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dbscheme_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(72)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	30: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dbscheme_repeat1),
	})))),
	31: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(95)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	32: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	33: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dbscheme_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(39)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	36: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dbscheme_repeat1),
	})))),
	37: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	38: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	39: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dbscheme_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(36)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_dbscheme),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_unionDecl),
		Fproduction_id: uint16(3),
	})))),
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
		Fcount: uint8(1),
	}})),
	46: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_unionDecl),
		Fproduction_id: uint16(3),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(70)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	60: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_unionDecl),
		Fproduction_id: uint16(3),
	})))),
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
		Fcount: uint8(1),
	}})),
	62: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_unionDecl),
		Fproduction_id: uint16(3),
	})))),
	63: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	64: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	65: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	66: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_unionDecl_repeat1),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_unionDecl_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	70: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_unionDecl_repeat1),
	})))),
	71: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	72: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	73: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_caseDecl),
		Fproduction_id: uint16(7),
	})))),
	74: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	75: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_caseDecl),
		Fproduction_id: uint16(7),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	77: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	78: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	79: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	80: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	81: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_caseDecl),
		Fproduction_id: uint16(7),
	})))),
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
		Fcount: uint8(1),
	}})),
	83: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_caseDecl),
		Fproduction_id: uint16(7),
	})))),
	84: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	85: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	86: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	87: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_caseDecl_repeat1),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_caseDecl_repeat1),
	})))),
	90: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	91: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_caseDecl_repeat1),
	})))),
	92: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(58)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_branch),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_branch),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	98: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_branch),
		Fproduction_id: uint16(15),
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
		Fcount: uint8(1),
	}})),
	100: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_branch),
		Fproduction_id: uint16(15),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	102: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_table),
		Fproduction_id: uint16(5),
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
		Fcount: uint8(1),
	}})),
	104: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_table),
		Fproduction_id: uint16(5),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_table),
		Fproduction_id: uint16(5),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_table),
		Fproduction_id: uint16(5),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	114: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_table),
		Fproduction_id: uint16(6),
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
		Fcount: uint8(1),
	}})),
	116: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_table),
		Fproduction_id: uint16(6),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	118: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_table),
		Fproduction_id: uint16(6),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_table),
		Fproduction_id: uint16(6),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	126: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	127: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	128: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_unionDecl),
		Fproduction_id: uint16(3),
	})))),
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
		Fcount: uint8(1),
	}})),
	130: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_unionDecl),
		Fproduction_id: uint16(3),
	})))),
	131: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	132: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	133: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_table),
		Fproduction_id: uint16(5),
	})))),
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
		Fcount: uint8(1),
	}})),
	136: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_table),
		Fproduction_id: uint16(5),
	})))),
	137: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	138: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_caseDecl),
		Fproduction_id: uint16(7),
	})))),
	139: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	140: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_caseDecl),
		Fproduction_id: uint16(7),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	142: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_table),
		Fproduction_id: uint16(6),
	})))),
	143: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	144: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_table),
		Fproduction_id: uint16(6),
	})))),
	145: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	146: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_entry),
	})))),
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
		Fcount: uint8(1),
	}})),
	148: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_entry),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_table_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	152: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_table_repeat1),
	})))),
	153: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(72)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	154: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	155: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_simpleId),
	})))),
	156: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	157: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_annotation),
		Fproduction_id: uint16(2),
	})))),
	158: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	159: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	160: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	161: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	162: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	163: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_annotName),
	})))),
	164: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_argsAnnotation),
		Fproduction_id: uint16(4),
	})))),
	166: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	167: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	168: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	169: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	170: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	171: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	172: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	173: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	174: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	175: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_annotation),
		Fproduction_id: uint16(1),
	})))),
	176: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	177: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	178: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	179: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	180: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	181: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_table_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(7)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_table_repeat2),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	188: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_argsAnnotation_repeat1),
	})))),
	189: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(62)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	190: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	191: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_argsAnnotation_repeat1),
	})))),
	192: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	193: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	194: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	195: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_colType),
	})))),
	196: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	197: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_column),
		Fproduction_id: uint16(8),
	})))),
	198: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	199: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	200: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	201: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_argsAnnotation),
		Fproduction_id: uint16(4),
	})))),
	202: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	203: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	204: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	205: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_column),
		Fproduction_id: uint16(9),
	})))),
	206: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	207: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	208: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	209: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_column),
		Fproduction_id: uint16(10),
	})))),
	210: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	211: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	212: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	213: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_argsAnnotation),
		Fproduction_id: uint16(4),
	})))),
	214: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	215: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_column),
		Fproduction_id: uint16(13),
	})))),
	216: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	217: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	218: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	219: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	220: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	221: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_column),
		Fproduction_id: uint16(16),
	})))),
	222: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	223: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_column),
		Fproduction_id: uint16(11),
	})))),
	224: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	225: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_reprType),
	})))),
	226: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	227: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_reprType),
	})))),
	228: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	229: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	230: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_column),
		Fproduction_id: uint16(12),
	})))),
	232: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	233: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_column),
		Fproduction_id: uint16(14),
	})))),
	234: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	235: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	236: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	237: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	238: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	239: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	240: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	241: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	242: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	243: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	244: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	245: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	246: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	247: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	248: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	249: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	250: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	251: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	252: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	253: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	254: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	255: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	256: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	257: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_tableName),
	})))),
	258: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	259: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	260: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	263: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	264: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	265: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	266: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	269: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	270: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	271: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	272: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(28)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	275: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	276: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	277: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

func tree_sitter_ql_dbscheme(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fversion:                   uint32(LANGUAGE_VERSION),
	Fsymbol_count:              uint32(SYMBOL_COUNT),
	Ftoken_count:               uint32(TOKEN_COUNT),
	Fstate_count:               uint32(STATE_COUNT),
	Flarge_state_count:         uint32(LARGE_STATE_COUNT),
	Fproduction_id_count:       uint32(PRODUCTION_ID_COUNT),
	Ffield_count:               uint32(FIELD_COUNT),
	Fmax_alias_sequence_length: uint16(MAX_ALIAS_SEQUENCE_LENGTH),
	Fparse_table:               uintptr(unsafe.Pointer(&ts_parse_table)),
	Fsmall_parse_table:         uintptr(unsafe.Pointer(&ts_small_parse_table)),
	Fsmall_parse_table_map:     uintptr(unsafe.Pointer(&ts_small_parse_table_map)),
	Fparse_actions:             uintptr(unsafe.Pointer(&ts_parse_actions)),
	Fsymbol_names:              uintptr(unsafe.Pointer(&ts_symbol_names)),
	Ffield_names:               uintptr(unsafe.Pointer(&ts_field_names)),
	Ffield_map_slices:          uintptr(unsafe.Pointer(&ts_field_map_slices)),
	Ffield_map_entries:         uintptr(unsafe.Pointer(&ts_field_map_entries)),
	Fsymbol_metadata:           uintptr(unsafe.Pointer(&ts_symbol_metadata)),
	Fpublic_symbol_map:         uintptr(unsafe.Pointer(&ts_symbol_map)),
	Falias_map:                 uintptr(unsafe.Pointer(&ts_non_terminal_alias_map)),
	Falias_sequences:           uintptr(unsafe.Pointer(&ts_alias_sequences)),
	Flex_modes:                 uintptr(unsafe.Pointer(&ts_lex_modes)),
	Fkeyword_capture_token:     uint16(sym__lower_id),
	Fprimary_state_ids:         uintptr(unsafe.Pointer(&ts_primary_state_ids)),
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(ts_lex_keywords)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00_lower_id\x00(\x00,\x00)\x00;\x00#\x00[\x00]\x00:\x00=\x00|\x00case\x00.\x00of\x00ref\x00int\x00float\x00boolean\x00date\x00varchar\x00string\x00unique\x00_upper_id\x00integer\x00dbtype\x00qldoc\x00line_comment\x00block_comment\x00dbscheme\x00entry\x00table\x00annotation\x00argsAnnotation\x00tableName\x00column\x00unionDecl\x00caseDecl\x00branch\x00colType\x00reprType\x00annotName\x00simpleId\x00dbscheme_repeat1\x00table_repeat1\x00table_repeat2\x00argsAnnotation_repeat1\x00unionDecl_repeat1\x00caseDecl_repeat1\x00base\x00colName\x00discriminator\x00isRef\x00isUnique\x00name\x00simpleAnnotation\x00"
