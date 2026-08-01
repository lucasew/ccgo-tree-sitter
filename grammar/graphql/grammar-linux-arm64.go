// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-graphql/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-graphql -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-graphql/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_graphql

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
const FIELD_COUNT = 0
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
const LANGUAGE_VERSION = 13
const LARGE_STATE_COUNT = 2
const MAX_ALIAS_SEQUENCE_LENGTH = 8
const PRODUCTION_ID_COUNT = 1
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 396
const SYMBOL_COUNT = 142
const TOKEN_COUNT = 62
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

const anon_sym_schema = 1
const anon_sym_LBRACE = 2
const anon_sym_RBRACE = 3
const anon_sym_extend = 4
const anon_sym_scalar = 5
const anon_sym_type = 6
const anon_sym_interface = 7
const anon_sym_union = 8
const anon_sym_enum = 9
const anon_sym_input = 10
const anon_sym_AMP = 11
const anon_sym_implements = 12
const anon_sym_COLON = 13
const anon_sym_LPAREN = 14
const anon_sym_RPAREN = 15
const anon_sym_EQ = 16
const anon_sym_PIPE = 17
const anon_sym_query = 18
const anon_sym_mutation = 19
const anon_sym_subscription = 20
const anon_sym_DOLLAR = 21
const anon_sym_DQUOTE_DQUOTE_DQUOTE = 22
const aux_sym_string_value_token1 = 23
const anon_sym_DQUOTE = 24
const aux_sym_string_value_token2 = 25
const sym_int_value = 26
const sym_float_value = 27
const anon_sym_true = 28
const anon_sym_false = 29
const sym_null_value = 30
const anon_sym_LBRACK = 31
const anon_sym_RBRACK = 32
const anon_sym_DOT_DOT_DOT = 33
const anon_sym_fragment = 34
const anon_sym_on = 35
const anon_sym_AT = 36
const anon_sym_directive = 37
const anon_sym_repeatable = 38
const anon_sym_QUERY = 39
const anon_sym_MUTATION = 40
const anon_sym_SUBSCRIPTION = 41
const anon_sym_FIELD = 42
const anon_sym_FRAGMENT_DEFINITION = 43
const anon_sym_FRAGMENT_SPREAD = 44
const anon_sym_INLINE_FRAGMENT = 45
const anon_sym_VARIABLE_DEFINITION = 46
const anon_sym_SCHEMA = 47
const anon_sym_SCALAR = 48
const anon_sym_OBJECT = 49
const anon_sym_FIELD_DEFINITION = 50
const anon_sym_ARGUMENT_DEFINITION = 51
const anon_sym_INTERFACE = 52
const anon_sym_UNION = 53
const anon_sym_ENUM = 54
const anon_sym_ENUM_VALUE = 55
const anon_sym_INPUT_OBJECT = 56
const anon_sym_INPUT_FIELD_DEFINITION = 57
const anon_sym_BANG = 58
const sym_name = 59
const sym_comment = 60
const sym_comma = 61
const sym_source_file = 62
const sym_document = 63
const sym_definition = 64
const sym_executable_definition = 65
const sym_type_system_definition = 66
const sym_type_system_extension = 67
const sym_schema_definition = 68
const sym_schema_extension = 69
const sym_type_extension = 70
const sym_scalar_type_extension = 71
const sym_object_type_extension = 72
const sym_interface_type_extension = 73
const sym_union_type_extension = 74
const sym_enum_type_extension = 75
const sym_input_object_type_extension = 76
const sym_input_fields_definition = 77
const sym_enum_values_definition = 78
const sym_enum_value_definition = 79
const sym_implements_interfaces = 80
const sym_fields_definition = 81
const sym_field_definition = 82
const sym_arguments_definition = 83
const sym_input_value_definition = 84
const sym_default_value = 85
const sym_union_member_types = 86
const sym_root_operation_type_definition = 87
const sym_operation_definition = 88
const sym_operation_type = 89
const sym_type_definition = 90
const sym_scalar_type_definition = 91
const sym_object_type_definition = 92
const sym_interface_type_definition = 93
const sym_union_type_definition = 94
const sym_enum_type_definition = 95
const sym_input_object_type_definition = 96
const sym_variable_definitions = 97
const sym_variable_definition = 98
const sym_selection_set = 99
const sym_selection = 100
const sym_field = 101
const sym_alias = 102
const sym_arguments = 103
const sym_argument = 104
const sym_value = 105
const sym_variable = 106
const sym_string_value = 107
const sym_boolean_value = 108
const sym_enum_value = 109
const sym_list_value = 110
const sym_object_value = 111
const sym_object_field = 112
const sym_fragment_spread = 113
const sym_fragment_definition = 114
const sym_fragment_name = 115
const sym_inline_fragment = 116
const sym_type_condition = 117
const sym_directives = 118
const sym_directive = 119
const sym_directive_definition = 120
const sym_directive_locations = 121
const sym_directive_location = 122
const sym_executable_directive_location = 123
const sym_type_system_directive_location = 124
const sym_type = 125
const sym_named_type = 126
const sym_list_type = 127
const sym_non_null_type = 128
const sym_description = 129
const aux_sym_document_repeat1 = 130
const aux_sym_schema_definition_repeat1 = 131
const aux_sym_input_object_type_extension_repeat1 = 132
const aux_sym_input_fields_definition_repeat1 = 133
const aux_sym_enum_values_definition_repeat1 = 134
const aux_sym_fields_definition_repeat1 = 135
const aux_sym_variable_definitions_repeat1 = 136
const aux_sym_selection_set_repeat1 = 137
const aux_sym_arguments_repeat1 = 138
const aux_sym_list_value_repeat1 = 139
const aux_sym_object_value_repeat1 = 140
const aux_sym_directives_repeat1 = 141

var ts_symbol_names = [142]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 11,
	3:   __ccgo_ts + 13,
	4:   __ccgo_ts + 15,
	5:   __ccgo_ts + 22,
	6:   __ccgo_ts + 29,
	7:   __ccgo_ts + 34,
	8:   __ccgo_ts + 44,
	9:   __ccgo_ts + 50,
	10:  __ccgo_ts + 55,
	11:  __ccgo_ts + 61,
	12:  __ccgo_ts + 63,
	13:  __ccgo_ts + 74,
	14:  __ccgo_ts + 76,
	15:  __ccgo_ts + 78,
	16:  __ccgo_ts + 80,
	17:  __ccgo_ts + 82,
	18:  __ccgo_ts + 84,
	19:  __ccgo_ts + 90,
	20:  __ccgo_ts + 99,
	21:  __ccgo_ts + 112,
	22:  __ccgo_ts + 114,
	23:  __ccgo_ts + 118,
	24:  __ccgo_ts + 138,
	25:  __ccgo_ts + 140,
	26:  __ccgo_ts + 160,
	27:  __ccgo_ts + 170,
	28:  __ccgo_ts + 182,
	29:  __ccgo_ts + 187,
	30:  __ccgo_ts + 193,
	31:  __ccgo_ts + 204,
	32:  __ccgo_ts + 206,
	33:  __ccgo_ts + 208,
	34:  __ccgo_ts + 212,
	35:  __ccgo_ts + 221,
	36:  __ccgo_ts + 224,
	37:  __ccgo_ts + 226,
	38:  __ccgo_ts + 236,
	39:  __ccgo_ts + 247,
	40:  __ccgo_ts + 253,
	41:  __ccgo_ts + 262,
	42:  __ccgo_ts + 275,
	43:  __ccgo_ts + 281,
	44:  __ccgo_ts + 301,
	45:  __ccgo_ts + 317,
	46:  __ccgo_ts + 333,
	47:  __ccgo_ts + 353,
	48:  __ccgo_ts + 360,
	49:  __ccgo_ts + 367,
	50:  __ccgo_ts + 374,
	51:  __ccgo_ts + 391,
	52:  __ccgo_ts + 411,
	53:  __ccgo_ts + 421,
	54:  __ccgo_ts + 427,
	55:  __ccgo_ts + 432,
	56:  __ccgo_ts + 443,
	57:  __ccgo_ts + 456,
	58:  __ccgo_ts + 479,
	59:  __ccgo_ts + 481,
	60:  __ccgo_ts + 486,
	61:  __ccgo_ts + 494,
	62:  __ccgo_ts + 500,
	63:  __ccgo_ts + 512,
	64:  __ccgo_ts + 521,
	65:  __ccgo_ts + 532,
	66:  __ccgo_ts + 554,
	67:  __ccgo_ts + 577,
	68:  __ccgo_ts + 599,
	69:  __ccgo_ts + 617,
	70:  __ccgo_ts + 634,
	71:  __ccgo_ts + 649,
	72:  __ccgo_ts + 671,
	73:  __ccgo_ts + 693,
	74:  __ccgo_ts + 718,
	75:  __ccgo_ts + 739,
	76:  __ccgo_ts + 759,
	77:  __ccgo_ts + 787,
	78:  __ccgo_ts + 811,
	79:  __ccgo_ts + 834,
	80:  __ccgo_ts + 856,
	81:  __ccgo_ts + 878,
	82:  __ccgo_ts + 896,
	83:  __ccgo_ts + 913,
	84:  __ccgo_ts + 934,
	85:  __ccgo_ts + 957,
	86:  __ccgo_ts + 971,
	87:  __ccgo_ts + 990,
	88:  __ccgo_ts + 1021,
	89:  __ccgo_ts + 1042,
	90:  __ccgo_ts + 1057,
	91:  __ccgo_ts + 1073,
	92:  __ccgo_ts + 1096,
	93:  __ccgo_ts + 1119,
	94:  __ccgo_ts + 1145,
	95:  __ccgo_ts + 1167,
	96:  __ccgo_ts + 1188,
	97:  __ccgo_ts + 1217,
	98:  __ccgo_ts + 1238,
	99:  __ccgo_ts + 1258,
	100: __ccgo_ts + 1272,
	101: __ccgo_ts + 1282,
	102: __ccgo_ts + 1288,
	103: __ccgo_ts + 1294,
	104: __ccgo_ts + 1304,
	105: __ccgo_ts + 1313,
	106: __ccgo_ts + 1319,
	107: __ccgo_ts + 1328,
	108: __ccgo_ts + 1341,
	109: __ccgo_ts + 1355,
	110: __ccgo_ts + 1366,
	111: __ccgo_ts + 1377,
	112: __ccgo_ts + 1390,
	113: __ccgo_ts + 1403,
	114: __ccgo_ts + 1419,
	115: __ccgo_ts + 1439,
	116: __ccgo_ts + 1453,
	117: __ccgo_ts + 1469,
	118: __ccgo_ts + 1484,
	119: __ccgo_ts + 226,
	120: __ccgo_ts + 1495,
	121: __ccgo_ts + 1516,
	122: __ccgo_ts + 1536,
	123: __ccgo_ts + 1555,
	124: __ccgo_ts + 1585,
	125: __ccgo_ts + 29,
	126: __ccgo_ts + 1616,
	127: __ccgo_ts + 1627,
	128: __ccgo_ts + 1637,
	129: __ccgo_ts + 1651,
	130: __ccgo_ts + 1663,
	131: __ccgo_ts + 1680,
	132: __ccgo_ts + 1706,
	133: __ccgo_ts + 1742,
	134: __ccgo_ts + 1774,
	135: __ccgo_ts + 1805,
	136: __ccgo_ts + 1831,
	137: __ccgo_ts + 1860,
	138: __ccgo_ts + 1882,
	139: __ccgo_ts + 1900,
	140: __ccgo_ts + 1919,
	141: __ccgo_ts + 1940,
}

var ts_symbol_map = [142]TSSymbol{
	1:   uint16(anon_sym_schema),
	2:   uint16(anon_sym_LBRACE),
	3:   uint16(anon_sym_RBRACE),
	4:   uint16(anon_sym_extend),
	5:   uint16(anon_sym_scalar),
	6:   uint16(anon_sym_type),
	7:   uint16(anon_sym_interface),
	8:   uint16(anon_sym_union),
	9:   uint16(anon_sym_enum),
	10:  uint16(anon_sym_input),
	11:  uint16(anon_sym_AMP),
	12:  uint16(anon_sym_implements),
	13:  uint16(anon_sym_COLON),
	14:  uint16(anon_sym_LPAREN),
	15:  uint16(anon_sym_RPAREN),
	16:  uint16(anon_sym_EQ),
	17:  uint16(anon_sym_PIPE),
	18:  uint16(anon_sym_query),
	19:  uint16(anon_sym_mutation),
	20:  uint16(anon_sym_subscription),
	21:  uint16(anon_sym_DOLLAR),
	22:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	23:  uint16(aux_sym_string_value_token1),
	24:  uint16(anon_sym_DQUOTE),
	25:  uint16(aux_sym_string_value_token2),
	26:  uint16(sym_int_value),
	27:  uint16(sym_float_value),
	28:  uint16(anon_sym_true),
	29:  uint16(anon_sym_false),
	30:  uint16(sym_null_value),
	31:  uint16(anon_sym_LBRACK),
	32:  uint16(anon_sym_RBRACK),
	33:  uint16(anon_sym_DOT_DOT_DOT),
	34:  uint16(anon_sym_fragment),
	35:  uint16(anon_sym_on),
	36:  uint16(anon_sym_AT),
	37:  uint16(anon_sym_directive),
	38:  uint16(anon_sym_repeatable),
	39:  uint16(anon_sym_QUERY),
	40:  uint16(anon_sym_MUTATION),
	41:  uint16(anon_sym_SUBSCRIPTION),
	42:  uint16(anon_sym_FIELD),
	43:  uint16(anon_sym_FRAGMENT_DEFINITION),
	44:  uint16(anon_sym_FRAGMENT_SPREAD),
	45:  uint16(anon_sym_INLINE_FRAGMENT),
	46:  uint16(anon_sym_VARIABLE_DEFINITION),
	47:  uint16(anon_sym_SCHEMA),
	48:  uint16(anon_sym_SCALAR),
	49:  uint16(anon_sym_OBJECT),
	50:  uint16(anon_sym_FIELD_DEFINITION),
	51:  uint16(anon_sym_ARGUMENT_DEFINITION),
	52:  uint16(anon_sym_INTERFACE),
	53:  uint16(anon_sym_UNION),
	54:  uint16(anon_sym_ENUM),
	55:  uint16(anon_sym_ENUM_VALUE),
	56:  uint16(anon_sym_INPUT_OBJECT),
	57:  uint16(anon_sym_INPUT_FIELD_DEFINITION),
	58:  uint16(anon_sym_BANG),
	59:  uint16(sym_name),
	60:  uint16(sym_comment),
	61:  uint16(sym_comma),
	62:  uint16(sym_source_file),
	63:  uint16(sym_document),
	64:  uint16(sym_definition),
	65:  uint16(sym_executable_definition),
	66:  uint16(sym_type_system_definition),
	67:  uint16(sym_type_system_extension),
	68:  uint16(sym_schema_definition),
	69:  uint16(sym_schema_extension),
	70:  uint16(sym_type_extension),
	71:  uint16(sym_scalar_type_extension),
	72:  uint16(sym_object_type_extension),
	73:  uint16(sym_interface_type_extension),
	74:  uint16(sym_union_type_extension),
	75:  uint16(sym_enum_type_extension),
	76:  uint16(sym_input_object_type_extension),
	77:  uint16(sym_input_fields_definition),
	78:  uint16(sym_enum_values_definition),
	79:  uint16(sym_enum_value_definition),
	80:  uint16(sym_implements_interfaces),
	81:  uint16(sym_fields_definition),
	82:  uint16(sym_field_definition),
	83:  uint16(sym_arguments_definition),
	84:  uint16(sym_input_value_definition),
	85:  uint16(sym_default_value),
	86:  uint16(sym_union_member_types),
	87:  uint16(sym_root_operation_type_definition),
	88:  uint16(sym_operation_definition),
	89:  uint16(sym_operation_type),
	90:  uint16(sym_type_definition),
	91:  uint16(sym_scalar_type_definition),
	92:  uint16(sym_object_type_definition),
	93:  uint16(sym_interface_type_definition),
	94:  uint16(sym_union_type_definition),
	95:  uint16(sym_enum_type_definition),
	96:  uint16(sym_input_object_type_definition),
	97:  uint16(sym_variable_definitions),
	98:  uint16(sym_variable_definition),
	99:  uint16(sym_selection_set),
	100: uint16(sym_selection),
	101: uint16(sym_field),
	102: uint16(sym_alias),
	103: uint16(sym_arguments),
	104: uint16(sym_argument),
	105: uint16(sym_value),
	106: uint16(sym_variable),
	107: uint16(sym_string_value),
	108: uint16(sym_boolean_value),
	109: uint16(sym_enum_value),
	110: uint16(sym_list_value),
	111: uint16(sym_object_value),
	112: uint16(sym_object_field),
	113: uint16(sym_fragment_spread),
	114: uint16(sym_fragment_definition),
	115: uint16(sym_fragment_name),
	116: uint16(sym_inline_fragment),
	117: uint16(sym_type_condition),
	118: uint16(sym_directives),
	119: uint16(sym_directive),
	120: uint16(sym_directive_definition),
	121: uint16(sym_directive_locations),
	122: uint16(sym_directive_location),
	123: uint16(sym_executable_directive_location),
	124: uint16(sym_type_system_directive_location),
	125: uint16(sym_type),
	126: uint16(sym_named_type),
	127: uint16(sym_list_type),
	128: uint16(sym_non_null_type),
	129: uint16(sym_description),
	130: uint16(aux_sym_document_repeat1),
	131: uint16(aux_sym_schema_definition_repeat1),
	132: uint16(aux_sym_input_object_type_extension_repeat1),
	133: uint16(aux_sym_input_fields_definition_repeat1),
	134: uint16(aux_sym_enum_values_definition_repeat1),
	135: uint16(aux_sym_fields_definition_repeat1),
	136: uint16(aux_sym_variable_definitions_repeat1),
	137: uint16(aux_sym_selection_set_repeat1),
	138: uint16(aux_sym_arguments_repeat1),
	139: uint16(aux_sym_list_value_repeat1),
	140: uint16(aux_sym_object_value_repeat1),
	141: uint16(aux_sym_directives_repeat1),
}

var ts_symbol_metadata = [142]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
	},
	16: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	17: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	18: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	19: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	20: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	21: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	22: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	23: {},
	24: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	25: {},
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
	},
	29: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	30: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	},
	48: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	49: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	50: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	51: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	52: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	53: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	54: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	55: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	57: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	58: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	59: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	60: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	76: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	77: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	78: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	79: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	80: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	81: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	82: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	83: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	84: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	85: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	86: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	87: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	88: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	89: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	90: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	91: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	92: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	93: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	94: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	95: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	96: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	97: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	98: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	99: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	100: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	101: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	102: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	103: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	104: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	105: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	106: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	107: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	108: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	109: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	110: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	111: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	112: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	113: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	114: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	115: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	116: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	117: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	118: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	119: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	120: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	121: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	122: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	123: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	124: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	125: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	126: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	127: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	128: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	129: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	130: {},
	131: {},
	132: {},
	133: {},
	134: {},
	135: {},
	136: {},
	137: {},
	138: {},
	139: {},
	140: {},
	141: {},
}

var ts_alias_sequences = [1][8]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var lookahead int32_t
	_, _, _, _ = eof, lookahead, result, skip
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
			state = uint16(269)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(338)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(296)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(290)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(280)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(283)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(284)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(354)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(300)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(282)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(285)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(316)
			goto next_state
		}
		if lookahead == int32('A') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('E') {
			state = uint16(114)
			goto next_state
		}
		if lookahead == int32('F') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('I') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('M') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('Q') {
			state = uint16(159)
			goto next_state
		}
		if lookahead == int32('S') {
			state = uint16(24)
			goto next_state
		}
		if lookahead == int32('U') {
			state = uint16(105)
			goto next_state
		}
		if lookahead == int32('V') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(310)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(311)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(205)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(228)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(171)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(217)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(259)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(260)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(222)
			goto next_state
		}
		if lookahead == int32('q') {
			state = uint16(258)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(195)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(182)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(244)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(227)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(271)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(286)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(272)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') || lookahead == int32(65279) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('!') {
			state = uint16(338)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(296)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(290)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(280)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(283)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(284)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(354)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(282)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(285)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(316)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(310)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(311)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(271)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(286)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(272)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') || lookahead == int32(65279) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('"') {
			state = uint16(296)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(290)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(354)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(300)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(310)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(311)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(339)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(349)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(346)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(271)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') || lookahead == int32(65279) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(2)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(301)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('"') {
			state = uint16(291)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('"') {
			state = uint16(295)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(354)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(316)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(271)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') || lookahead == int32(65279) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('"') {
			state = uint16(268)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(294)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('.') {
			state = uint16(312)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('.') {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('0') {
			state = uint16(300)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('A') {
			state = uint16(327)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('A') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('A') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('A') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('H') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('A') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('A') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('A') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('A') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('A') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('A') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('A') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('B') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('B') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('B') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('B') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('C') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('U') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('C') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('C') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('C') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('C') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('D') {
			state = uint16(322)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('D') {
			state = uint16(324)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('D') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('D') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('D') {
			state = uint16(53)
			goto next_state
		}
		if lookahead == int32('S') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('D') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('D') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('D') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('E') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('E') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('E') {
			state = uint16(332)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('E') {
			state = uint16(335)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('E') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('E') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('E') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('E') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('E') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('E') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('E') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('E') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('E') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('E') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('E') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('E') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('E') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('E') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('E') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('E') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('F') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('F') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('F') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('F') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('F') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('F') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('F') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('F') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('G') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('G') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('G') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('I') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('R') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('I') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('I') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('I') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('I') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('I') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('I') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('I') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('I') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('I') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('I') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead == int32('I') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead == int32('I') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead == int32('I') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead == int32('I') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('I') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('I') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('I') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('I') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead == int32('I') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead == int32('I') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead == int32('I') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('I') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('J') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('J') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('L') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead == int32('L') {
			state = uint16(71)
			goto next_state
		}
		if lookahead == int32('P') {
			state = uint16(160)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead == int32('L') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead == int32('L') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead == int32('L') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead == int32('L') {
			state = uint16(162)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead == int32('M') {
			state = uint16(334)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead == int32('M') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('M') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead == int32('M') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead == int32('M') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead == int32('N') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead == int32('N') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead == int32('N') {
			state = uint16(333)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead == int32('N') {
			state = uint16(320)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead == int32('N') {
			state = uint16(321)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('N') {
			state = uint16(330)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead == int32('N') {
			state = uint16(331)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead == int32('N') {
			state = uint16(323)
			goto next_state
		}
		return result
	case int32(112):
		if lookahead == int32('N') {
			state = uint16(326)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead == int32('N') {
			state = uint16(337)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('N') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('N') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead == int32('N') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead == int32('N') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead == int32('N') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('N') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('N') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('N') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('N') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead == int32('N') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(124):
		if lookahead == int32('O') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead == int32('O') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead == int32('O') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead == int32('O') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead == int32('O') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(129):
		if lookahead == int32('O') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(130):
		if lookahead == int32('O') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(131):
		if lookahead == int32('O') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(132):
		if lookahead == int32('P') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(133):
		if lookahead == int32('P') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(134):
		if lookahead == int32('R') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead == int32('R') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(136):
		if lookahead == int32('R') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(137):
		if lookahead == int32('R') {
			state = uint16(328)
			goto next_state
		}
		return result
	case int32(138):
		if lookahead == int32('R') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(139):
		if lookahead == int32('R') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(140):
		if lookahead == int32('R') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead == int32('R') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(142):
		if lookahead == int32('S') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(143):
		if lookahead == int32('T') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(144):
		if lookahead == int32('T') {
			state = uint16(329)
			goto next_state
		}
		return result
	case int32(145):
		if lookahead == int32('T') {
			state = uint16(336)
			goto next_state
		}
		return result
	case int32(146):
		if lookahead == int32('T') {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(147):
		if lookahead == int32('T') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(148):
		if lookahead == int32('T') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(149):
		if lookahead == int32('T') {
			state = uint16(166)
			goto next_state
		}
		return result
	case int32(150):
		if lookahead == int32('T') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(151):
		if lookahead == int32('T') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(152):
		if lookahead == int32('T') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead == int32('T') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(154):
		if lookahead == int32('T') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(155):
		if lookahead == int32('T') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(156):
		if lookahead == int32('T') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(157):
		if lookahead == int32('U') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(158):
		if lookahead == int32('U') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(159):
		if lookahead == int32('U') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(160):
		if lookahead == int32('U') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(161):
		if lookahead == int32('U') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(162):
		if lookahead == int32('U') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(163):
		if lookahead == int32('V') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(164):
		if lookahead == int32('Y') {
			state = uint16(319)
			goto next_state
		}
		return result
	case int32(165):
		if lookahead == int32('_') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(166):
		if lookahead == int32('_') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(167):
		if lookahead == int32('_') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(168):
		if lookahead == int32('_') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(169):
		if lookahead == int32('_') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(170):
		if lookahead == int32('_') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(171):
		if lookahead == int32('a') {
			state = uint16(211)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(172):
		if lookahead == int32('a') {
			state = uint16(204)
			goto next_state
		}
		return result
	case int32(173):
		if lookahead == int32('a') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(174):
		if lookahead == int32('a') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(175):
		if lookahead == int32('a') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(176):
		if lookahead == int32('a') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(177):
		if lookahead == int32('a') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(178):
		if lookahead == int32('a') {
			state = uint16(214)
			goto next_state
		}
		if lookahead == int32('h') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(179):
		if lookahead == int32('a') {
			state = uint16(254)
			goto next_state
		}
		return result
	case int32(180):
		if lookahead == int32('b') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(181):
		if lookahead == int32('b') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(182):
		if lookahead == int32('c') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(183):
		if lookahead == int32('c') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(184):
		if lookahead == int32('c') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(185):
		if lookahead == int32('c') {
			state = uint16(192)
			goto next_state
		}
		return result
	case int32(186):
		if lookahead == int32('d') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(187):
		if lookahead == int32('e') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(188):
		if lookahead == int32('e') {
			state = uint16(304)
			goto next_state
		}
		return result
	case int32(189):
		if lookahead == int32('e') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(190):
		if lookahead == int32('e') {
			state = uint16(306)
			goto next_state
		}
		return result
	case int32(191):
		if lookahead == int32('e') {
			state = uint16(317)
			goto next_state
		}
		return result
	case int32(192):
		if lookahead == int32('e') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(193):
		if lookahead == int32('e') {
			state = uint16(318)
			goto next_state
		}
		return result
	case int32(194):
		if lookahead == int32('e') {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(195):
		if lookahead == int32('e') {
			state = uint16(235)
			goto next_state
		}
		return result
	case int32(196):
		if lookahead == int32('e') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(197):
		if lookahead == int32('e') {
			state = uint16(240)
			goto next_state
		}
		return result
	case int32(198):
		if lookahead == int32('e') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(199):
		if lookahead == int32('e') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(200):
		if lookahead == int32('e') {
			state = uint16(229)
			goto next_state
		}
		return result
	case int32(201):
		if lookahead == int32('e') {
			state = uint16(230)
			goto next_state
		}
		return result
	case int32(202):
		if lookahead == int32('e') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(203):
		if lookahead == int32('f') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(204):
		if lookahead == int32('g') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(205):
		if lookahead == int32('i') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(206):
		if lookahead == int32('i') {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(207):
		if lookahead == int32('i') {
			state = uint16(263)
			goto next_state
		}
		return result
	case int32(208):
		if lookahead == int32('i') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(209):
		if lookahead == int32('i') {
			state = uint16(232)
			goto next_state
		}
		return result
	case int32(210):
		if lookahead == int32('i') {
			state = uint16(233)
			goto next_state
		}
		return result
	case int32(211):
		if lookahead == int32('l') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(212):
		if lookahead == int32('l') {
			state = uint16(308)
			goto next_state
		}
		return result
	case int32(213):
		if lookahead == int32('l') {
			state = uint16(212)
			goto next_state
		}
		return result
	case int32(214):
		if lookahead == int32('l') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(215):
		if lookahead == int32('l') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(216):
		if lookahead == int32('l') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(217):
		if lookahead == int32('m') {
			state = uint16(234)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(218):
		if lookahead == int32('m') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(219):
		if lookahead == int32('m') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(220):
		if lookahead == int32('m') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(221):
		if lookahead == int32('m') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(222):
		if lookahead == int32('n') {
			state = uint16(314)
			goto next_state
		}
		return result
	case int32(223):
		if lookahead == int32('n') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(224):
		if lookahead == int32('n') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(225):
		if lookahead == int32('n') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(226):
		if lookahead == int32('n') {
			state = uint16(289)
			goto next_state
		}
		return result
	case int32(227):
		if lookahead == int32('n') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(228):
		if lookahead == int32('n') {
			state = uint16(257)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(253)
			goto next_state
		}
		return result
	case int32(229):
		if lookahead == int32('n') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(230):
		if lookahead == int32('n') {
			state = uint16(252)
			goto next_state
		}
		return result
	case int32(231):
		if lookahead == int32('o') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(232):
		if lookahead == int32('o') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(233):
		if lookahead == int32('o') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(234):
		if lookahead == int32('p') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(235):
		if lookahead == int32('p') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(236):
		if lookahead == int32('p') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(237):
		if lookahead == int32('p') {
			state = uint16(261)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(238):
		if lookahead == int32('p') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(239):
		if lookahead == int32('r') {
			state = uint16(264)
			goto next_state
		}
		return result
	case int32(240):
		if lookahead == int32('r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(241):
		if lookahead == int32('r') {
			state = uint16(274)
			goto next_state
		}
		return result
	case int32(242):
		if lookahead == int32('r') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(243):
		if lookahead == int32('r') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(244):
		if lookahead == int32('r') {
			state = uint16(262)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(236)
			goto next_state
		}
		return result
	case int32(245):
		if lookahead == int32('s') {
			state = uint16(281)
			goto next_state
		}
		return result
	case int32(246):
		if lookahead == int32('s') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(247):
		if lookahead == int32('s') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(248):
		if lookahead == int32('t') {
			state = uint16(279)
			goto next_state
		}
		return result
	case int32(249):
		if lookahead == int32('t') {
			state = uint16(313)
			goto next_state
		}
		return result
	case int32(250):
		if lookahead == int32('t') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(251):
		if lookahead == int32('t') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(252):
		if lookahead == int32('t') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(253):
		if lookahead == int32('t') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(254):
		if lookahead == int32('t') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(255):
		if lookahead == int32('t') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(256):
		if lookahead == int32('t') {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(257):
		if lookahead == int32('u') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(258):
		if lookahead == int32('u') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(259):
		if lookahead == int32('u') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(260):
		if lookahead == int32('u') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(261):
		if lookahead == int32('u') {
			state = uint16(248)
			goto next_state
		}
		return result
	case int32(262):
		if lookahead == int32('u') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(263):
		if lookahead == int32('v') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(264):
		if lookahead == int32('y') {
			state = uint16(287)
			goto next_state
		}
		return result
	case int32(265):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(267)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(266):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(302)
			goto next_state
		}
		return result
	case int32(267):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(268):
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(294)
			goto next_state
		}
		return result
	case int32(269):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(270):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_schema)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(271):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(272):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(273):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_extend)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(274):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_scalar)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(275):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(276):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_interface)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(277):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_union)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(278):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(279):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_input)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(280):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(281):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_implements)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(282):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(283):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(284):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(285):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(286):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(287):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_query)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(288):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mutation)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(289):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_subscription)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(290):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(291):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(292):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(294)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(352)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(293):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') || lookahead == int32(65279) {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(292)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(356)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(294)
			goto next_state
		}
		return result
	case int32(294):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(294)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(295):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(296):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(297):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(298)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(355)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\r') || lookahead == int32(' ') || lookahead == int32(65279) {
			state = uint16(297)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(299)
			goto next_state
		}
		return result
	case int32(298):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') || lookahead == int32('\\') {
			state = uint16(353)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(298)
			goto next_state
		}
		return result
	case int32(299):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(299)
			goto next_state
		}
		return result
	case int32(300):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(266)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(265)
			goto next_state
		}
		return result
	case int32(301):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(266)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(265)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(302):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(265)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(302)
			goto next_state
		}
		return result
	case int32(303):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(304):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(305):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(306):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(307):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(308):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_null_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(309):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_null_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(310):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(311):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(312):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT_DOT_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(313):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(314):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_on)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(315):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_on)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(316):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(317):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_directive)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(318):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_repeatable)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(319):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QUERY)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(320):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MUTATION)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(321):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SUBSCRIPTION)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(322):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_FIELD)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(323):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_FRAGMENT_DEFINITION)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(324):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_FRAGMENT_SPREAD)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(325):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INLINE_FRAGMENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(326):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_VARIABLE_DEFINITION)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(327):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SCHEMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(328):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SCALAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(329):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_OBJECT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(330):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_FIELD_DEFINITION)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(331):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ARGUMENT_DEFINITION)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(332):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INTERFACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(333):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_UNION)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(334):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ENUM)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(163)
			goto next_state
		}
		return result
	case int32(335):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ENUM_VALUE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(336):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INPUT_OBJECT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(337):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INPUT_FIELD_DEFINITION)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(338):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(339):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(342)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(340):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(305)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(341):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(307)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(342):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(347)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(343):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(309)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(344):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(343)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(345):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(315)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(346):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(348)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(347):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(341)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(348):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(340)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(349):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(344)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(350):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(351):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(294)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(353)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(352):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(294)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(351)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(353):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(354):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comma)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(355):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comma)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(299)
			goto next_state
		}
		return result
	case int32(356):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comma)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') {
			state = uint16(294)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(5)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [396]TSLexMode{
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
	19: {},
	20: {},
	21: {},
	22: {},
	23: {},
	24: {},
	25: {},
	26: {},
	27: {},
	28: {},
	29: {},
	30: {},
	31: {},
	32: {},
	33: {},
	34: {},
	35: {},
	36: {},
	37: {
		Flex_state: uint16(2),
	},
	38: {
		Flex_state: uint16(2),
	},
	39: {
		Flex_state: uint16(2),
	},
	40: {
		Flex_state: uint16(2),
	},
	41: {
		Flex_state: uint16(2),
	},
	42: {
		Flex_state: uint16(2),
	},
	43: {
		Flex_state: uint16(2),
	},
	44: {},
	45: {},
	46: {},
	47: {},
	48: {},
	49: {},
	50: {},
	51: {},
	52: {},
	53: {},
	54: {},
	55: {
		Flex_state: uint16(2),
	},
	56: {},
	57: {},
	58: {},
	59: {},
	60: {},
	61: {
		Flex_state: uint16(2),
	},
	62: {},
	63: {},
	64: {},
	65: {},
	66: {},
	67: {},
	68: {},
	69: {},
	70: {},
	71: {},
	72: {},
	73: {},
	74: {},
	75: {},
	76: {},
	77: {},
	78: {},
	79: {},
	80: {},
	81: {},
	82: {},
	83: {
		Flex_state: uint16(2),
	},
	84: {},
	85: {},
	86: {},
	87: {},
	88: {
		Flex_state: uint16(2),
	},
	89:  {},
	90:  {},
	91:  {},
	92:  {},
	93:  {},
	94:  {},
	95:  {},
	96:  {},
	97:  {},
	98:  {},
	99:  {},
	100: {},
	101: {},
	102: {},
	103: {},
	104: {},
	105: {},
	106: {},
	107: {},
	108: {},
	109: {},
	110: {},
	111: {},
	112: {},
	113: {},
	114: {},
	115: {},
	116: {},
	117: {},
	118: {},
	119: {},
	120: {},
	121: {},
	122: {},
	123: {},
	124: {},
	125: {},
	126: {},
	127: {},
	128: {},
	129: {},
	130: {},
	131: {},
	132: {},
	133: {},
	134: {},
	135: {},
	136: {},
	137: {},
	138: {
		Flex_state: uint16(2),
	},
	139: {
		Flex_state: uint16(2),
	},
	140: {
		Flex_state: uint16(2),
	},
	141: {
		Flex_state: uint16(2),
	},
	142: {
		Flex_state: uint16(2),
	},
	143: {
		Flex_state: uint16(2),
	},
	144: {
		Flex_state: uint16(2),
	},
	145: {
		Flex_state: uint16(2),
	},
	146: {
		Flex_state: uint16(2),
	},
	147: {
		Flex_state: uint16(1),
	},
	148: {
		Flex_state: uint16(1),
	},
	149: {
		Flex_state: uint16(4),
	},
	150: {
		Flex_state: uint16(1),
	},
	151: {
		Flex_state: uint16(1),
	},
	152: {
		Flex_state: uint16(1),
	},
	153: {
		Flex_state: uint16(1),
	},
	154: {
		Flex_state: uint16(1),
	},
	155: {
		Flex_state: uint16(1),
	},
	156: {
		Flex_state: uint16(1),
	},
	157: {
		Flex_state: uint16(1),
	},
	158: {
		Flex_state: uint16(1),
	},
	159: {
		Flex_state: uint16(1),
	},
	160: {
		Flex_state: uint16(1),
	},
	161: {
		Flex_state: uint16(1),
	},
	162: {
		Flex_state: uint16(1),
	},
	163: {
		Flex_state: uint16(1),
	},
	164: {
		Flex_state: uint16(1),
	},
	165: {
		Flex_state: uint16(1),
	},
	166: {},
	167: {
		Flex_state: uint16(1),
	},
	168: {
		Flex_state: uint16(1),
	},
	169: {
		Flex_state: uint16(1),
	},
	170: {
		Flex_state: uint16(1),
	},
	171: {
		Flex_state: uint16(1),
	},
	172: {},
	173: {
		Flex_state: uint16(1),
	},
	174: {
		Flex_state: uint16(1),
	},
	175: {
		Flex_state: uint16(1),
	},
	176: {},
	177: {
		Flex_state: uint16(1),
	},
	178: {
		Flex_state: uint16(1),
	},
	179: {
		Flex_state: uint16(1),
	},
	180: {
		Flex_state: uint16(1),
	},
	181: {
		Flex_state: uint16(1),
	},
	182: {},
	183: {
		Flex_state: uint16(1),
	},
	184: {},
	185: {
		Flex_state: uint16(1),
	},
	186: {
		Flex_state: uint16(1),
	},
	187: {
		Flex_state: uint16(1),
	},
	188: {
		Flex_state: uint16(1),
	},
	189: {
		Flex_state: uint16(1),
	},
	190: {},
	191: {
		Flex_state: uint16(1),
	},
	192: {},
	193: {},
	194: {
		Flex_state: uint16(1),
	},
	195: {
		Flex_state: uint16(1),
	},
	196: {},
	197: {
		Flex_state: uint16(1),
	},
	198: {
		Flex_state: uint16(1),
	},
	199: {},
	200: {},
	201: {
		Flex_state: uint16(1),
	},
	202: {},
	203: {},
	204: {
		Flex_state: uint16(1),
	},
	205: {
		Flex_state: uint16(1),
	},
	206: {
		Flex_state: uint16(1),
	},
	207: {},
	208: {},
	209: {
		Flex_state: uint16(1),
	},
	210: {
		Flex_state: uint16(1),
	},
	211: {
		Flex_state: uint16(1),
	},
	212: {},
	213: {
		Flex_state: uint16(1),
	},
	214: {
		Flex_state: uint16(1),
	},
	215: {},
	216: {
		Flex_state: uint16(1),
	},
	217: {},
	218: {
		Flex_state: uint16(1),
	},
	219: {
		Flex_state: uint16(1),
	},
	220: {
		Flex_state: uint16(1),
	},
	221: {
		Flex_state: uint16(1),
	},
	222: {
		Flex_state: uint16(1),
	},
	223: {
		Flex_state: uint16(1),
	},
	224: {
		Flex_state: uint16(1),
	},
	225: {
		Flex_state: uint16(1),
	},
	226: {},
	227: {},
	228: {},
	229: {
		Flex_state: uint16(1),
	},
	230: {
		Flex_state: uint16(1),
	},
	231: {},
	232: {
		Flex_state: uint16(1),
	},
	233: {
		Flex_state: uint16(1),
	},
	234: {
		Flex_state: uint16(1),
	},
	235: {
		Flex_state: uint16(1),
	},
	236: {
		Flex_state: uint16(1),
	},
	237: {},
	238: {},
	239: {},
	240: {},
	241: {},
	242: {
		Flex_state: uint16(1),
	},
	243: {},
	244: {
		Flex_state: uint16(1),
	},
	245: {
		Flex_state: uint16(1),
	},
	246: {},
	247: {
		Flex_state: uint16(1),
	},
	248: {},
	249: {},
	250: {
		Flex_state: uint16(1),
	},
	251: {
		Flex_state: uint16(1),
	},
	252: {
		Flex_state: uint16(1),
	},
	253: {
		Flex_state: uint16(1),
	},
	254: {
		Flex_state: uint16(1),
	},
	255: {},
	256: {},
	257: {
		Flex_state: uint16(1),
	},
	258: {
		Flex_state: uint16(1),
	},
	259: {
		Flex_state: uint16(1),
	},
	260: {},
	261: {
		Flex_state: uint16(1),
	},
	262: {
		Flex_state: uint16(1),
	},
	263: {
		Flex_state: uint16(1),
	},
	264: {
		Flex_state: uint16(1),
	},
	265: {},
	266: {},
	267: {
		Flex_state: uint16(1),
	},
	268: {
		Flex_state: uint16(1),
	},
	269: {},
	270: {
		Flex_state: uint16(1),
	},
	271: {
		Flex_state: uint16(1),
	},
	272: {
		Flex_state: uint16(1),
	},
	273: {
		Flex_state: uint16(1),
	},
	274: {},
	275: {},
	276: {
		Flex_state: uint16(1),
	},
	277: {
		Flex_state: uint16(1),
	},
	278: {
		Flex_state: uint16(1),
	},
	279: {},
	280: {
		Flex_state: uint16(1),
	},
	281: {},
	282: {
		Flex_state: uint16(1),
	},
	283: {},
	284: {},
	285: {
		Flex_state: uint16(1),
	},
	286: {
		Flex_state: uint16(1),
	},
	287: {
		Flex_state: uint16(1),
	},
	288: {
		Flex_state: uint16(1),
	},
	289: {
		Flex_state: uint16(1),
	},
	290: {
		Flex_state: uint16(1),
	},
	291: {
		Flex_state: uint16(1),
	},
	292: {},
	293: {
		Flex_state: uint16(1),
	},
	294: {
		Flex_state: uint16(1),
	},
	295: {},
	296: {
		Flex_state: uint16(1),
	},
	297: {
		Flex_state: uint16(1),
	},
	298: {
		Flex_state: uint16(1),
	},
	299: {
		Flex_state: uint16(1),
	},
	300: {},
	301: {},
	302: {},
	303: {},
	304: {},
	305: {},
	306: {
		Flex_state: uint16(1),
	},
	307: {},
	308: {},
	309: {
		Flex_state: uint16(1),
	},
	310: {},
	311: {},
	312: {
		Flex_state: uint16(1),
	},
	313: {},
	314: {},
	315: {},
	316: {},
	317: {
		Flex_state: uint16(1),
	},
	318: {
		Flex_state: uint16(1),
	},
	319: {
		Flex_state: uint16(1),
	},
	320: {
		Flex_state: uint16(1),
	},
	321: {
		Flex_state: uint16(1),
	},
	322: {},
	323: {},
	324: {},
	325: {
		Flex_state: uint16(1),
	},
	326: {},
	327: {
		Flex_state: uint16(4),
	},
	328: {},
	329: {
		Flex_state: uint16(1),
	},
	330: {
		Flex_state: uint16(1),
	},
	331: {
		Flex_state: uint16(1),
	},
	332: {
		Flex_state: uint16(1),
	},
	333: {
		Flex_state: uint16(293),
	},
	334: {
		Flex_state: uint16(1),
	},
	335: {
		Flex_state: uint16(1),
	},
	336: {
		Flex_state: uint16(1),
	},
	337: {},
	338: {},
	339: {},
	340: {
		Flex_state: uint16(1),
	},
	341: {},
	342: {},
	343: {
		Flex_state: uint16(297),
	},
	344: {},
	345: {
		Flex_state: uint16(1),
	},
	346: {},
	347: {
		Flex_state: uint16(1),
	},
	348: {},
	349: {
		Flex_state: uint16(1),
	},
	350: {
		Flex_state: uint16(1),
	},
	351: {},
	352: {},
	353: {},
	354: {},
	355: {
		Flex_state: uint16(1),
	},
	356: {},
	357: {
		Flex_state: uint16(1),
	},
	358: {},
	359: {
		Flex_state: uint16(1),
	},
	360: {},
	361: {
		Flex_state: uint16(1),
	},
	362: {
		Flex_state: uint16(1),
	},
	363: {
		Flex_state: uint16(1),
	},
	364: {},
	365: {},
	366: {
		Flex_state: uint16(4),
	},
	367: {
		Flex_state: uint16(1),
	},
	368: {
		Flex_state: uint16(1),
	},
	369: {},
	370: {},
	371: {
		Flex_state: uint16(1),
	},
	372: {},
	373: {
		Flex_state: uint16(1),
	},
	374: {
		Flex_state: uint16(1),
	},
	375: {
		Flex_state: uint16(1),
	},
	376: {},
	377: {
		Flex_state: uint16(4),
	},
	378: {
		Flex_state: uint16(1),
	},
	379: {},
	380: {},
	381: {
		Flex_state: uint16(1),
	},
	382: {
		Flex_state: uint16(1),
	},
	383: {},
	384: {},
	385: {
		Flex_state: uint16(4),
	},
	386: {
		Flex_state: uint16(1),
	},
	387: {
		Flex_state: uint16(293),
	},
	388: {
		Flex_state: uint16(297),
	},
	389: {
		Flex_state: uint16(1),
	},
	390: {
		Flex_state: uint16(1),
	},
	391: {
		Flex_state: uint16(293),
	},
	392: {
		Flex_state: uint16(297),
	},
	393: {
		Flex_state: uint16(1),
	},
	394: {
		Flex_state: uint16(293),
	},
	395: {
		Flex_state: uint16(297),
	},
}

var ts_parse_table = [2][142]uint16_t{
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
		24: uint16(1),
		26: uint16(1),
		27: uint16(1),
		28: uint16(1),
		29: uint16(1),
		30: uint16(1),
		31: uint16(1),
		32: uint16(1),
		33: uint16(1),
		34: uint16(1),
		35: uint16(1),
		36: uint16(1),
		37: uint16(1),
		38: uint16(1),
		39: uint16(1),
		40: uint16(1),
		41: uint16(1),
		42: uint16(1),
		43: uint16(1),
		44: uint16(1),
		45: uint16(1),
		46: uint16(1),
		47: uint16(1),
		48: uint16(1),
		49: uint16(1),
		50: uint16(1),
		51: uint16(1),
		52: uint16(1),
		53: uint16(1),
		54: uint16(1),
		55: uint16(1),
		56: uint16(1),
		57: uint16(1),
		58: uint16(1),
		60: uint16(3),
		61: uint16(3),
	},
	1: {
		1:   uint16(5),
		2:   uint16(7),
		4:   uint16(9),
		5:   uint16(11),
		6:   uint16(13),
		7:   uint16(15),
		8:   uint16(17),
		9:   uint16(19),
		10:  uint16(21),
		18:  uint16(23),
		19:  uint16(23),
		20:  uint16(23),
		22:  uint16(25),
		24:  uint16(27),
		34:  uint16(29),
		37:  uint16(31),
		60:  uint16(3),
		61:  uint16(3),
		62:  uint16(352),
		63:  uint16(354),
		64:  uint16(2),
		65:  uint16(131),
		66:  uint16(131),
		67:  uint16(131),
		68:  uint16(123),
		69:  uint16(117),
		70:  uint16(117),
		71:  uint16(113),
		72:  uint16(113),
		73:  uint16(113),
		74:  uint16(113),
		75:  uint16(113),
		76:  uint16(113),
		88:  uint16(110),
		89:  uint16(156),
		90:  uint16(123),
		91:  uint16(107),
		92:  uint16(107),
		93:  uint16(107),
		94:  uint16(107),
		95:  uint16(107),
		96:  uint16(107),
		99:  uint16(134),
		107: uint16(166),
		114: uint16(110),
		120: uint16(123),
		129: uint16(182),
		130: uint16(2),
	},
}

var ts_small_parse_table = [8350]uint16_t{
	0:    uint16(27),
	1:    uint16(5),
	2:    uint16(1),
	3:    uint16(anon_sym_schema),
	4:    uint16(7),
	5:    uint16(1),
	6:    uint16(anon_sym_LBRACE),
	7:    uint16(9),
	8:    uint16(1),
	9:    uint16(anon_sym_extend),
	10:   uint16(11),
	11:   uint16(1),
	12:   uint16(anon_sym_scalar),
	13:   uint16(13),
	14:   uint16(1),
	15:   uint16(anon_sym_type),
	16:   uint16(15),
	17:   uint16(1),
	18:   uint16(anon_sym_interface),
	19:   uint16(17),
	20:   uint16(1),
	21:   uint16(anon_sym_union),
	22:   uint16(19),
	23:   uint16(1),
	24:   uint16(anon_sym_enum),
	25:   uint16(21),
	26:   uint16(1),
	27:   uint16(anon_sym_input),
	28:   uint16(25),
	29:   uint16(1),
	30:   uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	31:   uint16(27),
	32:   uint16(1),
	33:   uint16(anon_sym_DQUOTE),
	34:   uint16(29),
	35:   uint16(1),
	36:   uint16(anon_sym_fragment),
	37:   uint16(31),
	38:   uint16(1),
	39:   uint16(anon_sym_directive),
	40:   uint16(33),
	41:   uint16(1),
	43:   uint16(134),
	44:   uint16(1),
	45:   uint16(sym_selection_set),
	46:   uint16(156),
	47:   uint16(1),
	48:   uint16(sym_operation_type),
	49:   uint16(166),
	50:   uint16(1),
	51:   uint16(sym_string_value),
	52:   uint16(182),
	53:   uint16(1),
	54:   uint16(sym_description),
	55:   uint16(3),
	56:   uint16(2),
	57:   uint16(sym_comment),
	58:   uint16(sym_comma),
	59:   uint16(3),
	60:   uint16(2),
	61:   uint16(sym_definition),
	62:   uint16(aux_sym_document_repeat1),
	63:   uint16(110),
	64:   uint16(2),
	65:   uint16(sym_operation_definition),
	66:   uint16(sym_fragment_definition),
	67:   uint16(117),
	68:   uint16(2),
	69:   uint16(sym_schema_extension),
	70:   uint16(sym_type_extension),
	71:   uint16(23),
	72:   uint16(3),
	73:   uint16(anon_sym_query),
	74:   uint16(anon_sym_mutation),
	75:   uint16(anon_sym_subscription),
	76:   uint16(123),
	77:   uint16(3),
	78:   uint16(sym_schema_definition),
	79:   uint16(sym_type_definition),
	80:   uint16(sym_directive_definition),
	81:   uint16(131),
	82:   uint16(3),
	83:   uint16(sym_executable_definition),
	84:   uint16(sym_type_system_definition),
	85:   uint16(sym_type_system_extension),
	86:   uint16(107),
	87:   uint16(6),
	88:   uint16(sym_scalar_type_definition),
	89:   uint16(sym_object_type_definition),
	90:   uint16(sym_interface_type_definition),
	91:   uint16(sym_union_type_definition),
	92:   uint16(sym_enum_type_definition),
	93:   uint16(sym_input_object_type_definition),
	94:   uint16(113),
	95:   uint16(6),
	96:   uint16(sym_scalar_type_extension),
	97:   uint16(sym_object_type_extension),
	98:   uint16(sym_interface_type_extension),
	99:   uint16(sym_union_type_extension),
	100:  uint16(sym_enum_type_extension),
	101:  uint16(sym_input_object_type_extension),
	102:  uint16(27),
	103:  uint16(35),
	104:  uint16(1),
	106:  uint16(37),
	107:  uint16(1),
	108:  uint16(anon_sym_schema),
	109:  uint16(40),
	110:  uint16(1),
	111:  uint16(anon_sym_LBRACE),
	112:  uint16(43),
	113:  uint16(1),
	114:  uint16(anon_sym_extend),
	115:  uint16(46),
	116:  uint16(1),
	117:  uint16(anon_sym_scalar),
	118:  uint16(49),
	119:  uint16(1),
	120:  uint16(anon_sym_type),
	121:  uint16(52),
	122:  uint16(1),
	123:  uint16(anon_sym_interface),
	124:  uint16(55),
	125:  uint16(1),
	126:  uint16(anon_sym_union),
	127:  uint16(58),
	128:  uint16(1),
	129:  uint16(anon_sym_enum),
	130:  uint16(61),
	131:  uint16(1),
	132:  uint16(anon_sym_input),
	133:  uint16(67),
	134:  uint16(1),
	135:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	136:  uint16(70),
	137:  uint16(1),
	138:  uint16(anon_sym_DQUOTE),
	139:  uint16(73),
	140:  uint16(1),
	141:  uint16(anon_sym_fragment),
	142:  uint16(76),
	143:  uint16(1),
	144:  uint16(anon_sym_directive),
	145:  uint16(134),
	146:  uint16(1),
	147:  uint16(sym_selection_set),
	148:  uint16(156),
	149:  uint16(1),
	150:  uint16(sym_operation_type),
	151:  uint16(166),
	152:  uint16(1),
	153:  uint16(sym_string_value),
	154:  uint16(182),
	155:  uint16(1),
	156:  uint16(sym_description),
	157:  uint16(3),
	158:  uint16(2),
	159:  uint16(sym_comment),
	160:  uint16(sym_comma),
	161:  uint16(3),
	162:  uint16(2),
	163:  uint16(sym_definition),
	164:  uint16(aux_sym_document_repeat1),
	165:  uint16(110),
	166:  uint16(2),
	167:  uint16(sym_operation_definition),
	168:  uint16(sym_fragment_definition),
	169:  uint16(117),
	170:  uint16(2),
	171:  uint16(sym_schema_extension),
	172:  uint16(sym_type_extension),
	173:  uint16(64),
	174:  uint16(3),
	175:  uint16(anon_sym_query),
	176:  uint16(anon_sym_mutation),
	177:  uint16(anon_sym_subscription),
	178:  uint16(123),
	179:  uint16(3),
	180:  uint16(sym_schema_definition),
	181:  uint16(sym_type_definition),
	182:  uint16(sym_directive_definition),
	183:  uint16(131),
	184:  uint16(3),
	185:  uint16(sym_executable_definition),
	186:  uint16(sym_type_system_definition),
	187:  uint16(sym_type_system_extension),
	188:  uint16(107),
	189:  uint16(6),
	190:  uint16(sym_scalar_type_definition),
	191:  uint16(sym_object_type_definition),
	192:  uint16(sym_interface_type_definition),
	193:  uint16(sym_union_type_definition),
	194:  uint16(sym_enum_type_definition),
	195:  uint16(sym_input_object_type_definition),
	196:  uint16(113),
	197:  uint16(6),
	198:  uint16(sym_scalar_type_extension),
	199:  uint16(sym_object_type_extension),
	200:  uint16(sym_interface_type_extension),
	201:  uint16(sym_union_type_extension),
	202:  uint16(sym_enum_type_extension),
	203:  uint16(sym_input_object_type_extension),
	204:  uint16(10),
	205:  uint16(81),
	206:  uint16(1),
	207:  uint16(anon_sym_LBRACE),
	208:  uint16(83),
	209:  uint16(1),
	210:  uint16(anon_sym_implements),
	211:  uint16(85),
	212:  uint16(1),
	213:  uint16(anon_sym_DQUOTE),
	214:  uint16(87),
	215:  uint16(1),
	216:  uint16(anon_sym_AT),
	217:  uint16(20),
	218:  uint16(1),
	219:  uint16(sym_implements_interfaces),
	220:  uint16(81),
	221:  uint16(1),
	222:  uint16(sym_directives),
	223:  uint16(91),
	224:  uint16(1),
	225:  uint16(sym_fields_definition),
	226:  uint16(3),
	227:  uint16(2),
	228:  uint16(sym_comment),
	229:  uint16(sym_comma),
	230:  uint16(36),
	231:  uint16(2),
	232:  uint16(sym_directive),
	233:  uint16(aux_sym_directives_repeat1),
	234:  uint16(79),
	235:  uint16(15),
	237:  uint16(anon_sym_schema),
	238:  uint16(anon_sym_extend),
	239:  uint16(anon_sym_scalar),
	240:  uint16(anon_sym_type),
	241:  uint16(anon_sym_interface),
	242:  uint16(anon_sym_union),
	243:  uint16(anon_sym_enum),
	244:  uint16(anon_sym_input),
	245:  uint16(anon_sym_query),
	246:  uint16(anon_sym_mutation),
	247:  uint16(anon_sym_subscription),
	248:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	249:  uint16(anon_sym_fragment),
	250:  uint16(anon_sym_directive),
	251:  uint16(9),
	252:  uint16(89),
	253:  uint16(1),
	254:  uint16(anon_sym_PIPE),
	255:  uint16(93),
	256:  uint16(1),
	257:  uint16(anon_sym_FIELD),
	258:  uint16(97),
	259:  uint16(1),
	260:  uint16(anon_sym_ENUM),
	261:  uint16(73),
	262:  uint16(1),
	263:  uint16(sym_directive_locations),
	264:  uint16(74),
	265:  uint16(1),
	266:  uint16(sym_directive_location),
	267:  uint16(3),
	268:  uint16(2),
	269:  uint16(sym_comment),
	270:  uint16(sym_comma),
	271:  uint16(75),
	272:  uint16(2),
	273:  uint16(sym_executable_directive_location),
	274:  uint16(sym_type_system_directive_location),
	275:  uint16(91),
	276:  uint16(7),
	277:  uint16(anon_sym_QUERY),
	278:  uint16(anon_sym_MUTATION),
	279:  uint16(anon_sym_SUBSCRIPTION),
	280:  uint16(anon_sym_FRAGMENT_DEFINITION),
	281:  uint16(anon_sym_FRAGMENT_SPREAD),
	282:  uint16(anon_sym_INLINE_FRAGMENT),
	283:  uint16(anon_sym_VARIABLE_DEFINITION),
	284:  uint16(95),
	285:  uint16(10),
	286:  uint16(anon_sym_SCHEMA),
	287:  uint16(anon_sym_SCALAR),
	288:  uint16(anon_sym_OBJECT),
	289:  uint16(anon_sym_FIELD_DEFINITION),
	290:  uint16(anon_sym_ARGUMENT_DEFINITION),
	291:  uint16(anon_sym_INTERFACE),
	292:  uint16(anon_sym_UNION),
	293:  uint16(anon_sym_ENUM_VALUE),
	294:  uint16(anon_sym_INPUT_OBJECT),
	295:  uint16(anon_sym_INPUT_FIELD_DEFINITION),
	296:  uint16(10),
	297:  uint16(81),
	298:  uint16(1),
	299:  uint16(anon_sym_LBRACE),
	300:  uint16(83),
	301:  uint16(1),
	302:  uint16(anon_sym_implements),
	303:  uint16(87),
	304:  uint16(1),
	305:  uint16(anon_sym_AT),
	306:  uint16(101),
	307:  uint16(1),
	308:  uint16(anon_sym_DQUOTE),
	309:  uint16(21),
	310:  uint16(1),
	311:  uint16(sym_implements_interfaces),
	312:  uint16(66),
	313:  uint16(1),
	314:  uint16(sym_directives),
	315:  uint16(136),
	316:  uint16(1),
	317:  uint16(sym_fields_definition),
	318:  uint16(3),
	319:  uint16(2),
	320:  uint16(sym_comment),
	321:  uint16(sym_comma),
	322:  uint16(36),
	323:  uint16(2),
	324:  uint16(sym_directive),
	325:  uint16(aux_sym_directives_repeat1),
	326:  uint16(99),
	327:  uint16(15),
	329:  uint16(anon_sym_schema),
	330:  uint16(anon_sym_extend),
	331:  uint16(anon_sym_scalar),
	332:  uint16(anon_sym_type),
	333:  uint16(anon_sym_interface),
	334:  uint16(anon_sym_union),
	335:  uint16(anon_sym_enum),
	336:  uint16(anon_sym_input),
	337:  uint16(anon_sym_query),
	338:  uint16(anon_sym_mutation),
	339:  uint16(anon_sym_subscription),
	340:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	341:  uint16(anon_sym_fragment),
	342:  uint16(anon_sym_directive),
	343:  uint16(9),
	344:  uint16(89),
	345:  uint16(1),
	346:  uint16(anon_sym_PIPE),
	347:  uint16(93),
	348:  uint16(1),
	349:  uint16(anon_sym_FIELD),
	350:  uint16(97),
	351:  uint16(1),
	352:  uint16(anon_sym_ENUM),
	353:  uint16(74),
	354:  uint16(1),
	355:  uint16(sym_directive_location),
	356:  uint16(82),
	357:  uint16(1),
	358:  uint16(sym_directive_locations),
	359:  uint16(3),
	360:  uint16(2),
	361:  uint16(sym_comment),
	362:  uint16(sym_comma),
	363:  uint16(75),
	364:  uint16(2),
	365:  uint16(sym_executable_directive_location),
	366:  uint16(sym_type_system_directive_location),
	367:  uint16(91),
	368:  uint16(7),
	369:  uint16(anon_sym_QUERY),
	370:  uint16(anon_sym_MUTATION),
	371:  uint16(anon_sym_SUBSCRIPTION),
	372:  uint16(anon_sym_FRAGMENT_DEFINITION),
	373:  uint16(anon_sym_FRAGMENT_SPREAD),
	374:  uint16(anon_sym_INLINE_FRAGMENT),
	375:  uint16(anon_sym_VARIABLE_DEFINITION),
	376:  uint16(95),
	377:  uint16(10),
	378:  uint16(anon_sym_SCHEMA),
	379:  uint16(anon_sym_SCALAR),
	380:  uint16(anon_sym_OBJECT),
	381:  uint16(anon_sym_FIELD_DEFINITION),
	382:  uint16(anon_sym_ARGUMENT_DEFINITION),
	383:  uint16(anon_sym_INTERFACE),
	384:  uint16(anon_sym_UNION),
	385:  uint16(anon_sym_ENUM_VALUE),
	386:  uint16(anon_sym_INPUT_OBJECT),
	387:  uint16(anon_sym_INPUT_FIELD_DEFINITION),
	388:  uint16(10),
	389:  uint16(81),
	390:  uint16(1),
	391:  uint16(anon_sym_LBRACE),
	392:  uint16(83),
	393:  uint16(1),
	394:  uint16(anon_sym_implements),
	395:  uint16(87),
	396:  uint16(1),
	397:  uint16(anon_sym_AT),
	398:  uint16(105),
	399:  uint16(1),
	400:  uint16(anon_sym_DQUOTE),
	401:  uint16(16),
	402:  uint16(1),
	403:  uint16(sym_implements_interfaces),
	404:  uint16(59),
	405:  uint16(1),
	406:  uint16(sym_directives),
	407:  uint16(128),
	408:  uint16(1),
	409:  uint16(sym_fields_definition),
	410:  uint16(3),
	411:  uint16(2),
	412:  uint16(sym_comment),
	413:  uint16(sym_comma),
	414:  uint16(36),
	415:  uint16(2),
	416:  uint16(sym_directive),
	417:  uint16(aux_sym_directives_repeat1),
	418:  uint16(103),
	419:  uint16(15),
	421:  uint16(anon_sym_schema),
	422:  uint16(anon_sym_extend),
	423:  uint16(anon_sym_scalar),
	424:  uint16(anon_sym_type),
	425:  uint16(anon_sym_interface),
	426:  uint16(anon_sym_union),
	427:  uint16(anon_sym_enum),
	428:  uint16(anon_sym_input),
	429:  uint16(anon_sym_query),
	430:  uint16(anon_sym_mutation),
	431:  uint16(anon_sym_subscription),
	432:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	433:  uint16(anon_sym_fragment),
	434:  uint16(anon_sym_directive),
	435:  uint16(10),
	436:  uint16(81),
	437:  uint16(1),
	438:  uint16(anon_sym_LBRACE),
	439:  uint16(83),
	440:  uint16(1),
	441:  uint16(anon_sym_implements),
	442:  uint16(87),
	443:  uint16(1),
	444:  uint16(anon_sym_AT),
	445:  uint16(109),
	446:  uint16(1),
	447:  uint16(anon_sym_DQUOTE),
	448:  uint16(17),
	449:  uint16(1),
	450:  uint16(sym_implements_interfaces),
	451:  uint16(68),
	452:  uint16(1),
	453:  uint16(sym_directives),
	454:  uint16(106),
	455:  uint16(1),
	456:  uint16(sym_fields_definition),
	457:  uint16(3),
	458:  uint16(2),
	459:  uint16(sym_comment),
	460:  uint16(sym_comma),
	461:  uint16(36),
	462:  uint16(2),
	463:  uint16(sym_directive),
	464:  uint16(aux_sym_directives_repeat1),
	465:  uint16(107),
	466:  uint16(15),
	468:  uint16(anon_sym_schema),
	469:  uint16(anon_sym_extend),
	470:  uint16(anon_sym_scalar),
	471:  uint16(anon_sym_type),
	472:  uint16(anon_sym_interface),
	473:  uint16(anon_sym_union),
	474:  uint16(anon_sym_enum),
	475:  uint16(anon_sym_input),
	476:  uint16(anon_sym_query),
	477:  uint16(anon_sym_mutation),
	478:  uint16(anon_sym_subscription),
	479:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	480:  uint16(anon_sym_fragment),
	481:  uint16(anon_sym_directive),
	482:  uint16(9),
	483:  uint16(89),
	484:  uint16(1),
	485:  uint16(anon_sym_PIPE),
	486:  uint16(93),
	487:  uint16(1),
	488:  uint16(anon_sym_FIELD),
	489:  uint16(97),
	490:  uint16(1),
	491:  uint16(anon_sym_ENUM),
	492:  uint16(74),
	493:  uint16(1),
	494:  uint16(sym_directive_location),
	495:  uint16(87),
	496:  uint16(1),
	497:  uint16(sym_directive_locations),
	498:  uint16(3),
	499:  uint16(2),
	500:  uint16(sym_comment),
	501:  uint16(sym_comma),
	502:  uint16(75),
	503:  uint16(2),
	504:  uint16(sym_executable_directive_location),
	505:  uint16(sym_type_system_directive_location),
	506:  uint16(91),
	507:  uint16(7),
	508:  uint16(anon_sym_QUERY),
	509:  uint16(anon_sym_MUTATION),
	510:  uint16(anon_sym_SUBSCRIPTION),
	511:  uint16(anon_sym_FRAGMENT_DEFINITION),
	512:  uint16(anon_sym_FRAGMENT_SPREAD),
	513:  uint16(anon_sym_INLINE_FRAGMENT),
	514:  uint16(anon_sym_VARIABLE_DEFINITION),
	515:  uint16(95),
	516:  uint16(10),
	517:  uint16(anon_sym_SCHEMA),
	518:  uint16(anon_sym_SCALAR),
	519:  uint16(anon_sym_OBJECT),
	520:  uint16(anon_sym_FIELD_DEFINITION),
	521:  uint16(anon_sym_ARGUMENT_DEFINITION),
	522:  uint16(anon_sym_INTERFACE),
	523:  uint16(anon_sym_UNION),
	524:  uint16(anon_sym_ENUM_VALUE),
	525:  uint16(anon_sym_INPUT_OBJECT),
	526:  uint16(anon_sym_INPUT_FIELD_DEFINITION),
	527:  uint16(10),
	528:  uint16(81),
	529:  uint16(1),
	530:  uint16(anon_sym_LBRACE),
	531:  uint16(83),
	532:  uint16(1),
	533:  uint16(anon_sym_implements),
	534:  uint16(87),
	535:  uint16(1),
	536:  uint16(anon_sym_AT),
	537:  uint16(113),
	538:  uint16(1),
	539:  uint16(anon_sym_DQUOTE),
	540:  uint16(24),
	541:  uint16(1),
	542:  uint16(sym_implements_interfaces),
	543:  uint16(58),
	544:  uint16(1),
	545:  uint16(sym_directives),
	546:  uint16(135),
	547:  uint16(1),
	548:  uint16(sym_fields_definition),
	549:  uint16(3),
	550:  uint16(2),
	551:  uint16(sym_comment),
	552:  uint16(sym_comma),
	553:  uint16(36),
	554:  uint16(2),
	555:  uint16(sym_directive),
	556:  uint16(aux_sym_directives_repeat1),
	557:  uint16(111),
	558:  uint16(15),
	560:  uint16(anon_sym_schema),
	561:  uint16(anon_sym_extend),
	562:  uint16(anon_sym_scalar),
	563:  uint16(anon_sym_type),
	564:  uint16(anon_sym_interface),
	565:  uint16(anon_sym_union),
	566:  uint16(anon_sym_enum),
	567:  uint16(anon_sym_input),
	568:  uint16(anon_sym_query),
	569:  uint16(anon_sym_mutation),
	570:  uint16(anon_sym_subscription),
	571:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	572:  uint16(anon_sym_fragment),
	573:  uint16(anon_sym_directive),
	574:  uint16(9),
	575:  uint16(89),
	576:  uint16(1),
	577:  uint16(anon_sym_PIPE),
	578:  uint16(93),
	579:  uint16(1),
	580:  uint16(anon_sym_FIELD),
	581:  uint16(97),
	582:  uint16(1),
	583:  uint16(anon_sym_ENUM),
	584:  uint16(74),
	585:  uint16(1),
	586:  uint16(sym_directive_location),
	587:  uint16(89),
	588:  uint16(1),
	589:  uint16(sym_directive_locations),
	590:  uint16(3),
	591:  uint16(2),
	592:  uint16(sym_comment),
	593:  uint16(sym_comma),
	594:  uint16(75),
	595:  uint16(2),
	596:  uint16(sym_executable_directive_location),
	597:  uint16(sym_type_system_directive_location),
	598:  uint16(91),
	599:  uint16(7),
	600:  uint16(anon_sym_QUERY),
	601:  uint16(anon_sym_MUTATION),
	602:  uint16(anon_sym_SUBSCRIPTION),
	603:  uint16(anon_sym_FRAGMENT_DEFINITION),
	604:  uint16(anon_sym_FRAGMENT_SPREAD),
	605:  uint16(anon_sym_INLINE_FRAGMENT),
	606:  uint16(anon_sym_VARIABLE_DEFINITION),
	607:  uint16(95),
	608:  uint16(10),
	609:  uint16(anon_sym_SCHEMA),
	610:  uint16(anon_sym_SCALAR),
	611:  uint16(anon_sym_OBJECT),
	612:  uint16(anon_sym_FIELD_DEFINITION),
	613:  uint16(anon_sym_ARGUMENT_DEFINITION),
	614:  uint16(anon_sym_INTERFACE),
	615:  uint16(anon_sym_UNION),
	616:  uint16(anon_sym_ENUM_VALUE),
	617:  uint16(anon_sym_INPUT_OBJECT),
	618:  uint16(anon_sym_INPUT_FIELD_DEFINITION),
	619:  uint16(10),
	620:  uint16(81),
	621:  uint16(1),
	622:  uint16(anon_sym_LBRACE),
	623:  uint16(83),
	624:  uint16(1),
	625:  uint16(anon_sym_implements),
	626:  uint16(87),
	627:  uint16(1),
	628:  uint16(anon_sym_AT),
	629:  uint16(117),
	630:  uint16(1),
	631:  uint16(anon_sym_DQUOTE),
	632:  uint16(18),
	633:  uint16(1),
	634:  uint16(sym_implements_interfaces),
	635:  uint16(84),
	636:  uint16(1),
	637:  uint16(sym_directives),
	638:  uint16(111),
	639:  uint16(1),
	640:  uint16(sym_fields_definition),
	641:  uint16(3),
	642:  uint16(2),
	643:  uint16(sym_comment),
	644:  uint16(sym_comma),
	645:  uint16(36),
	646:  uint16(2),
	647:  uint16(sym_directive),
	648:  uint16(aux_sym_directives_repeat1),
	649:  uint16(115),
	650:  uint16(15),
	652:  uint16(anon_sym_schema),
	653:  uint16(anon_sym_extend),
	654:  uint16(anon_sym_scalar),
	655:  uint16(anon_sym_type),
	656:  uint16(anon_sym_interface),
	657:  uint16(anon_sym_union),
	658:  uint16(anon_sym_enum),
	659:  uint16(anon_sym_input),
	660:  uint16(anon_sym_query),
	661:  uint16(anon_sym_mutation),
	662:  uint16(anon_sym_subscription),
	663:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	664:  uint16(anon_sym_fragment),
	665:  uint16(anon_sym_directive),
	666:  uint16(8),
	667:  uint16(87),
	668:  uint16(1),
	669:  uint16(anon_sym_AT),
	670:  uint16(121),
	671:  uint16(1),
	672:  uint16(anon_sym_EQ),
	673:  uint16(123),
	674:  uint16(1),
	675:  uint16(anon_sym_DQUOTE),
	676:  uint16(54),
	677:  uint16(1),
	678:  uint16(sym_directives),
	679:  uint16(56),
	680:  uint16(1),
	681:  uint16(sym_union_member_types),
	682:  uint16(3),
	683:  uint16(2),
	684:  uint16(sym_comment),
	685:  uint16(sym_comma),
	686:  uint16(36),
	687:  uint16(2),
	688:  uint16(sym_directive),
	689:  uint16(aux_sym_directives_repeat1),
	690:  uint16(119),
	691:  uint16(16),
	693:  uint16(anon_sym_schema),
	694:  uint16(anon_sym_LBRACE),
	695:  uint16(anon_sym_extend),
	696:  uint16(anon_sym_scalar),
	697:  uint16(anon_sym_type),
	698:  uint16(anon_sym_interface),
	699:  uint16(anon_sym_union),
	700:  uint16(anon_sym_enum),
	701:  uint16(anon_sym_input),
	702:  uint16(anon_sym_query),
	703:  uint16(anon_sym_mutation),
	704:  uint16(anon_sym_subscription),
	705:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	706:  uint16(anon_sym_fragment),
	707:  uint16(anon_sym_directive),
	708:  uint16(3),
	709:  uint16(127),
	710:  uint16(1),
	711:  uint16(anon_sym_DQUOTE),
	712:  uint16(3),
	713:  uint16(2),
	714:  uint16(sym_comment),
	715:  uint16(sym_comma),
	716:  uint16(125),
	717:  uint16(22),
	719:  uint16(anon_sym_schema),
	720:  uint16(anon_sym_LBRACE),
	721:  uint16(anon_sym_RBRACE),
	722:  uint16(anon_sym_extend),
	723:  uint16(anon_sym_scalar),
	724:  uint16(anon_sym_type),
	725:  uint16(anon_sym_interface),
	726:  uint16(anon_sym_union),
	727:  uint16(anon_sym_enum),
	728:  uint16(anon_sym_input),
	729:  uint16(anon_sym_AMP),
	730:  uint16(anon_sym_PIPE),
	731:  uint16(anon_sym_query),
	732:  uint16(anon_sym_mutation),
	733:  uint16(anon_sym_subscription),
	734:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	735:  uint16(anon_sym_RBRACK),
	736:  uint16(anon_sym_fragment),
	737:  uint16(anon_sym_AT),
	738:  uint16(anon_sym_directive),
	739:  uint16(anon_sym_BANG),
	740:  uint16(9),
	741:  uint16(81),
	742:  uint16(1),
	743:  uint16(anon_sym_LBRACE),
	744:  uint16(87),
	745:  uint16(1),
	746:  uint16(anon_sym_AT),
	747:  uint16(131),
	748:  uint16(1),
	749:  uint16(anon_sym_AMP),
	750:  uint16(133),
	751:  uint16(1),
	752:  uint16(anon_sym_DQUOTE),
	753:  uint16(76),
	754:  uint16(1),
	755:  uint16(sym_directives),
	756:  uint16(102),
	757:  uint16(1),
	758:  uint16(sym_fields_definition),
	759:  uint16(3),
	760:  uint16(2),
	761:  uint16(sym_comment),
	762:  uint16(sym_comma),
	763:  uint16(36),
	764:  uint16(2),
	765:  uint16(sym_directive),
	766:  uint16(aux_sym_directives_repeat1),
	767:  uint16(129),
	768:  uint16(15),
	770:  uint16(anon_sym_schema),
	771:  uint16(anon_sym_extend),
	772:  uint16(anon_sym_scalar),
	773:  uint16(anon_sym_type),
	774:  uint16(anon_sym_interface),
	775:  uint16(anon_sym_union),
	776:  uint16(anon_sym_enum),
	777:  uint16(anon_sym_input),
	778:  uint16(anon_sym_query),
	779:  uint16(anon_sym_mutation),
	780:  uint16(anon_sym_subscription),
	781:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	782:  uint16(anon_sym_fragment),
	783:  uint16(anon_sym_directive),
	784:  uint16(9),
	785:  uint16(81),
	786:  uint16(1),
	787:  uint16(anon_sym_LBRACE),
	788:  uint16(87),
	789:  uint16(1),
	790:  uint16(anon_sym_AT),
	791:  uint16(131),
	792:  uint16(1),
	793:  uint16(anon_sym_AMP),
	794:  uint16(137),
	795:  uint16(1),
	796:  uint16(anon_sym_DQUOTE),
	797:  uint16(64),
	798:  uint16(1),
	799:  uint16(sym_directives),
	800:  uint16(112),
	801:  uint16(1),
	802:  uint16(sym_fields_definition),
	803:  uint16(3),
	804:  uint16(2),
	805:  uint16(sym_comment),
	806:  uint16(sym_comma),
	807:  uint16(36),
	808:  uint16(2),
	809:  uint16(sym_directive),
	810:  uint16(aux_sym_directives_repeat1),
	811:  uint16(135),
	812:  uint16(15),
	814:  uint16(anon_sym_schema),
	815:  uint16(anon_sym_extend),
	816:  uint16(anon_sym_scalar),
	817:  uint16(anon_sym_type),
	818:  uint16(anon_sym_interface),
	819:  uint16(anon_sym_union),
	820:  uint16(anon_sym_enum),
	821:  uint16(anon_sym_input),
	822:  uint16(anon_sym_query),
	823:  uint16(anon_sym_mutation),
	824:  uint16(anon_sym_subscription),
	825:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	826:  uint16(anon_sym_fragment),
	827:  uint16(anon_sym_directive),
	828:  uint16(9),
	829:  uint16(81),
	830:  uint16(1),
	831:  uint16(anon_sym_LBRACE),
	832:  uint16(87),
	833:  uint16(1),
	834:  uint16(anon_sym_AT),
	835:  uint16(113),
	836:  uint16(1),
	837:  uint16(anon_sym_DQUOTE),
	838:  uint16(131),
	839:  uint16(1),
	840:  uint16(anon_sym_AMP),
	841:  uint16(58),
	842:  uint16(1),
	843:  uint16(sym_directives),
	844:  uint16(135),
	845:  uint16(1),
	846:  uint16(sym_fields_definition),
	847:  uint16(3),
	848:  uint16(2),
	849:  uint16(sym_comment),
	850:  uint16(sym_comma),
	851:  uint16(36),
	852:  uint16(2),
	853:  uint16(sym_directive),
	854:  uint16(aux_sym_directives_repeat1),
	855:  uint16(111),
	856:  uint16(15),
	858:  uint16(anon_sym_schema),
	859:  uint16(anon_sym_extend),
	860:  uint16(anon_sym_scalar),
	861:  uint16(anon_sym_type),
	862:  uint16(anon_sym_interface),
	863:  uint16(anon_sym_union),
	864:  uint16(anon_sym_enum),
	865:  uint16(anon_sym_input),
	866:  uint16(anon_sym_query),
	867:  uint16(anon_sym_mutation),
	868:  uint16(anon_sym_subscription),
	869:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	870:  uint16(anon_sym_fragment),
	871:  uint16(anon_sym_directive),
	872:  uint16(8),
	873:  uint16(87),
	874:  uint16(1),
	875:  uint16(anon_sym_AT),
	876:  uint16(121),
	877:  uint16(1),
	878:  uint16(anon_sym_EQ),
	879:  uint16(141),
	880:  uint16(1),
	881:  uint16(anon_sym_DQUOTE),
	882:  uint16(50),
	883:  uint16(1),
	884:  uint16(sym_directives),
	885:  uint16(86),
	886:  uint16(1),
	887:  uint16(sym_union_member_types),
	888:  uint16(3),
	889:  uint16(2),
	890:  uint16(sym_comment),
	891:  uint16(sym_comma),
	892:  uint16(36),
	893:  uint16(2),
	894:  uint16(sym_directive),
	895:  uint16(aux_sym_directives_repeat1),
	896:  uint16(139),
	897:  uint16(16),
	899:  uint16(anon_sym_schema),
	900:  uint16(anon_sym_LBRACE),
	901:  uint16(anon_sym_extend),
	902:  uint16(anon_sym_scalar),
	903:  uint16(anon_sym_type),
	904:  uint16(anon_sym_interface),
	905:  uint16(anon_sym_union),
	906:  uint16(anon_sym_enum),
	907:  uint16(anon_sym_input),
	908:  uint16(anon_sym_query),
	909:  uint16(anon_sym_mutation),
	910:  uint16(anon_sym_subscription),
	911:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	912:  uint16(anon_sym_fragment),
	913:  uint16(anon_sym_directive),
	914:  uint16(9),
	915:  uint16(81),
	916:  uint16(1),
	917:  uint16(anon_sym_LBRACE),
	918:  uint16(87),
	919:  uint16(1),
	920:  uint16(anon_sym_AT),
	921:  uint16(105),
	922:  uint16(1),
	923:  uint16(anon_sym_DQUOTE),
	924:  uint16(131),
	925:  uint16(1),
	926:  uint16(anon_sym_AMP),
	927:  uint16(59),
	928:  uint16(1),
	929:  uint16(sym_directives),
	930:  uint16(128),
	931:  uint16(1),
	932:  uint16(sym_fields_definition),
	933:  uint16(3),
	934:  uint16(2),
	935:  uint16(sym_comment),
	936:  uint16(sym_comma),
	937:  uint16(36),
	938:  uint16(2),
	939:  uint16(sym_directive),
	940:  uint16(aux_sym_directives_repeat1),
	941:  uint16(103),
	942:  uint16(15),
	944:  uint16(anon_sym_schema),
	945:  uint16(anon_sym_extend),
	946:  uint16(anon_sym_scalar),
	947:  uint16(anon_sym_type),
	948:  uint16(anon_sym_interface),
	949:  uint16(anon_sym_union),
	950:  uint16(anon_sym_enum),
	951:  uint16(anon_sym_input),
	952:  uint16(anon_sym_query),
	953:  uint16(anon_sym_mutation),
	954:  uint16(anon_sym_subscription),
	955:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	956:  uint16(anon_sym_fragment),
	957:  uint16(anon_sym_directive),
	958:  uint16(9),
	959:  uint16(81),
	960:  uint16(1),
	961:  uint16(anon_sym_LBRACE),
	962:  uint16(87),
	963:  uint16(1),
	964:  uint16(anon_sym_AT),
	965:  uint16(131),
	966:  uint16(1),
	967:  uint16(anon_sym_AMP),
	968:  uint16(145),
	969:  uint16(1),
	970:  uint16(anon_sym_DQUOTE),
	971:  uint16(67),
	972:  uint16(1),
	973:  uint16(sym_directives),
	974:  uint16(109),
	975:  uint16(1),
	976:  uint16(sym_fields_definition),
	977:  uint16(3),
	978:  uint16(2),
	979:  uint16(sym_comment),
	980:  uint16(sym_comma),
	981:  uint16(36),
	982:  uint16(2),
	983:  uint16(sym_directive),
	984:  uint16(aux_sym_directives_repeat1),
	985:  uint16(143),
	986:  uint16(15),
	988:  uint16(anon_sym_schema),
	989:  uint16(anon_sym_extend),
	990:  uint16(anon_sym_scalar),
	991:  uint16(anon_sym_type),
	992:  uint16(anon_sym_interface),
	993:  uint16(anon_sym_union),
	994:  uint16(anon_sym_enum),
	995:  uint16(anon_sym_input),
	996:  uint16(anon_sym_query),
	997:  uint16(anon_sym_mutation),
	998:  uint16(anon_sym_subscription),
	999:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1000: uint16(anon_sym_fragment),
	1001: uint16(anon_sym_directive),
	1002: uint16(8),
	1003: uint16(87),
	1004: uint16(1),
	1005: uint16(anon_sym_AT),
	1006: uint16(149),
	1007: uint16(1),
	1008: uint16(anon_sym_LBRACE),
	1009: uint16(151),
	1010: uint16(1),
	1011: uint16(anon_sym_DQUOTE),
	1012: uint16(48),
	1013: uint16(1),
	1014: uint16(sym_directives),
	1015: uint16(3),
	1016: uint16(2),
	1017: uint16(sym_comment),
	1018: uint16(sym_comma),
	1019: uint16(36),
	1020: uint16(2),
	1021: uint16(sym_directive),
	1022: uint16(aux_sym_directives_repeat1),
	1023: uint16(49),
	1024: uint16(2),
	1025: uint16(sym_input_fields_definition),
	1026: uint16(aux_sym_input_object_type_extension_repeat1),
	1027: uint16(147),
	1028: uint16(15),
	1030: uint16(anon_sym_schema),
	1031: uint16(anon_sym_extend),
	1032: uint16(anon_sym_scalar),
	1033: uint16(anon_sym_type),
	1034: uint16(anon_sym_interface),
	1035: uint16(anon_sym_union),
	1036: uint16(anon_sym_enum),
	1037: uint16(anon_sym_input),
	1038: uint16(anon_sym_query),
	1039: uint16(anon_sym_mutation),
	1040: uint16(anon_sym_subscription),
	1041: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1042: uint16(anon_sym_fragment),
	1043: uint16(anon_sym_directive),
	1044: uint16(8),
	1045: uint16(87),
	1046: uint16(1),
	1047: uint16(anon_sym_AT),
	1048: uint16(121),
	1049: uint16(1),
	1050: uint16(anon_sym_EQ),
	1051: uint16(155),
	1052: uint16(1),
	1053: uint16(anon_sym_DQUOTE),
	1054: uint16(47),
	1055: uint16(1),
	1056: uint16(sym_directives),
	1057: uint16(65),
	1058: uint16(1),
	1059: uint16(sym_union_member_types),
	1060: uint16(3),
	1061: uint16(2),
	1062: uint16(sym_comment),
	1063: uint16(sym_comma),
	1064: uint16(36),
	1065: uint16(2),
	1066: uint16(sym_directive),
	1067: uint16(aux_sym_directives_repeat1),
	1068: uint16(153),
	1069: uint16(16),
	1071: uint16(anon_sym_schema),
	1072: uint16(anon_sym_LBRACE),
	1073: uint16(anon_sym_extend),
	1074: uint16(anon_sym_scalar),
	1075: uint16(anon_sym_type),
	1076: uint16(anon_sym_interface),
	1077: uint16(anon_sym_union),
	1078: uint16(anon_sym_enum),
	1079: uint16(anon_sym_input),
	1080: uint16(anon_sym_query),
	1081: uint16(anon_sym_mutation),
	1082: uint16(anon_sym_subscription),
	1083: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1084: uint16(anon_sym_fragment),
	1085: uint16(anon_sym_directive),
	1086: uint16(9),
	1087: uint16(81),
	1088: uint16(1),
	1089: uint16(anon_sym_LBRACE),
	1090: uint16(87),
	1091: uint16(1),
	1092: uint16(anon_sym_AT),
	1093: uint16(131),
	1094: uint16(1),
	1095: uint16(anon_sym_AMP),
	1096: uint16(159),
	1097: uint16(1),
	1098: uint16(anon_sym_DQUOTE),
	1099: uint16(77),
	1100: uint16(1),
	1101: uint16(sym_directives),
	1102: uint16(101),
	1103: uint16(1),
	1104: uint16(sym_fields_definition),
	1105: uint16(3),
	1106: uint16(2),
	1107: uint16(sym_comment),
	1108: uint16(sym_comma),
	1109: uint16(36),
	1110: uint16(2),
	1111: uint16(sym_directive),
	1112: uint16(aux_sym_directives_repeat1),
	1113: uint16(157),
	1114: uint16(15),
	1116: uint16(anon_sym_schema),
	1117: uint16(anon_sym_extend),
	1118: uint16(anon_sym_scalar),
	1119: uint16(anon_sym_type),
	1120: uint16(anon_sym_interface),
	1121: uint16(anon_sym_union),
	1122: uint16(anon_sym_enum),
	1123: uint16(anon_sym_input),
	1124: uint16(anon_sym_query),
	1125: uint16(anon_sym_mutation),
	1126: uint16(anon_sym_subscription),
	1127: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1128: uint16(anon_sym_fragment),
	1129: uint16(anon_sym_directive),
	1130: uint16(8),
	1131: uint16(87),
	1132: uint16(1),
	1133: uint16(anon_sym_AT),
	1134: uint16(163),
	1135: uint16(1),
	1136: uint16(anon_sym_LBRACE),
	1137: uint16(165),
	1138: uint16(1),
	1139: uint16(anon_sym_DQUOTE),
	1140: uint16(60),
	1141: uint16(1),
	1142: uint16(sym_directives),
	1143: uint16(137),
	1144: uint16(1),
	1145: uint16(sym_enum_values_definition),
	1146: uint16(3),
	1147: uint16(2),
	1148: uint16(sym_comment),
	1149: uint16(sym_comma),
	1150: uint16(36),
	1151: uint16(2),
	1152: uint16(sym_directive),
	1153: uint16(aux_sym_directives_repeat1),
	1154: uint16(161),
	1155: uint16(15),
	1157: uint16(anon_sym_schema),
	1158: uint16(anon_sym_extend),
	1159: uint16(anon_sym_scalar),
	1160: uint16(anon_sym_type),
	1161: uint16(anon_sym_interface),
	1162: uint16(anon_sym_union),
	1163: uint16(anon_sym_enum),
	1164: uint16(anon_sym_input),
	1165: uint16(anon_sym_query),
	1166: uint16(anon_sym_mutation),
	1167: uint16(anon_sym_subscription),
	1168: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1169: uint16(anon_sym_fragment),
	1170: uint16(anon_sym_directive),
	1171: uint16(8),
	1172: uint16(87),
	1173: uint16(1),
	1174: uint16(anon_sym_AT),
	1175: uint16(149),
	1176: uint16(1),
	1177: uint16(anon_sym_LBRACE),
	1178: uint16(169),
	1179: uint16(1),
	1180: uint16(anon_sym_DQUOTE),
	1181: uint16(62),
	1182: uint16(1),
	1183: uint16(sym_directives),
	1184: uint16(133),
	1185: uint16(1),
	1186: uint16(sym_input_fields_definition),
	1187: uint16(3),
	1188: uint16(2),
	1189: uint16(sym_comment),
	1190: uint16(sym_comma),
	1191: uint16(36),
	1192: uint16(2),
	1193: uint16(sym_directive),
	1194: uint16(aux_sym_directives_repeat1),
	1195: uint16(167),
	1196: uint16(15),
	1198: uint16(anon_sym_schema),
	1199: uint16(anon_sym_extend),
	1200: uint16(anon_sym_scalar),
	1201: uint16(anon_sym_type),
	1202: uint16(anon_sym_interface),
	1203: uint16(anon_sym_union),
	1204: uint16(anon_sym_enum),
	1205: uint16(anon_sym_input),
	1206: uint16(anon_sym_query),
	1207: uint16(anon_sym_mutation),
	1208: uint16(anon_sym_subscription),
	1209: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1210: uint16(anon_sym_fragment),
	1211: uint16(anon_sym_directive),
	1212: uint16(7),
	1213: uint16(93),
	1214: uint16(1),
	1215: uint16(anon_sym_FIELD),
	1216: uint16(97),
	1217: uint16(1),
	1218: uint16(anon_sym_ENUM),
	1219: uint16(85),
	1220: uint16(1),
	1221: uint16(sym_directive_location),
	1222: uint16(3),
	1223: uint16(2),
	1224: uint16(sym_comment),
	1225: uint16(sym_comma),
	1226: uint16(75),
	1227: uint16(2),
	1228: uint16(sym_executable_directive_location),
	1229: uint16(sym_type_system_directive_location),
	1230: uint16(91),
	1231: uint16(7),
	1232: uint16(anon_sym_QUERY),
	1233: uint16(anon_sym_MUTATION),
	1234: uint16(anon_sym_SUBSCRIPTION),
	1235: uint16(anon_sym_FRAGMENT_DEFINITION),
	1236: uint16(anon_sym_FRAGMENT_SPREAD),
	1237: uint16(anon_sym_INLINE_FRAGMENT),
	1238: uint16(anon_sym_VARIABLE_DEFINITION),
	1239: uint16(95),
	1240: uint16(10),
	1241: uint16(anon_sym_SCHEMA),
	1242: uint16(anon_sym_SCALAR),
	1243: uint16(anon_sym_OBJECT),
	1244: uint16(anon_sym_FIELD_DEFINITION),
	1245: uint16(anon_sym_ARGUMENT_DEFINITION),
	1246: uint16(anon_sym_INTERFACE),
	1247: uint16(anon_sym_UNION),
	1248: uint16(anon_sym_ENUM_VALUE),
	1249: uint16(anon_sym_INPUT_OBJECT),
	1250: uint16(anon_sym_INPUT_FIELD_DEFINITION),
	1251: uint16(8),
	1252: uint16(87),
	1253: uint16(1),
	1254: uint16(anon_sym_AT),
	1255: uint16(163),
	1256: uint16(1),
	1257: uint16(anon_sym_LBRACE),
	1258: uint16(173),
	1259: uint16(1),
	1260: uint16(anon_sym_DQUOTE),
	1261: uint16(90),
	1262: uint16(1),
	1263: uint16(sym_directives),
	1264: uint16(114),
	1265: uint16(1),
	1266: uint16(sym_enum_values_definition),
	1267: uint16(3),
	1268: uint16(2),
	1269: uint16(sym_comment),
	1270: uint16(sym_comma),
	1271: uint16(36),
	1272: uint16(2),
	1273: uint16(sym_directive),
	1274: uint16(aux_sym_directives_repeat1),
	1275: uint16(171),
	1276: uint16(15),
	1278: uint16(anon_sym_schema),
	1279: uint16(anon_sym_extend),
	1280: uint16(anon_sym_scalar),
	1281: uint16(anon_sym_type),
	1282: uint16(anon_sym_interface),
	1283: uint16(anon_sym_union),
	1284: uint16(anon_sym_enum),
	1285: uint16(anon_sym_input),
	1286: uint16(anon_sym_query),
	1287: uint16(anon_sym_mutation),
	1288: uint16(anon_sym_subscription),
	1289: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1290: uint16(anon_sym_fragment),
	1291: uint16(anon_sym_directive),
	1292: uint16(7),
	1293: uint16(93),
	1294: uint16(1),
	1295: uint16(anon_sym_FIELD),
	1296: uint16(97),
	1297: uint16(1),
	1298: uint16(anon_sym_ENUM),
	1299: uint16(80),
	1300: uint16(1),
	1301: uint16(sym_directive_location),
	1302: uint16(3),
	1303: uint16(2),
	1304: uint16(sym_comment),
	1305: uint16(sym_comma),
	1306: uint16(75),
	1307: uint16(2),
	1308: uint16(sym_executable_directive_location),
	1309: uint16(sym_type_system_directive_location),
	1310: uint16(91),
	1311: uint16(7),
	1312: uint16(anon_sym_QUERY),
	1313: uint16(anon_sym_MUTATION),
	1314: uint16(anon_sym_SUBSCRIPTION),
	1315: uint16(anon_sym_FRAGMENT_DEFINITION),
	1316: uint16(anon_sym_FRAGMENT_SPREAD),
	1317: uint16(anon_sym_INLINE_FRAGMENT),
	1318: uint16(anon_sym_VARIABLE_DEFINITION),
	1319: uint16(95),
	1320: uint16(10),
	1321: uint16(anon_sym_SCHEMA),
	1322: uint16(anon_sym_SCALAR),
	1323: uint16(anon_sym_OBJECT),
	1324: uint16(anon_sym_FIELD_DEFINITION),
	1325: uint16(anon_sym_ARGUMENT_DEFINITION),
	1326: uint16(anon_sym_INTERFACE),
	1327: uint16(anon_sym_UNION),
	1328: uint16(anon_sym_ENUM_VALUE),
	1329: uint16(anon_sym_INPUT_OBJECT),
	1330: uint16(anon_sym_INPUT_FIELD_DEFINITION),
	1331: uint16(8),
	1332: uint16(87),
	1333: uint16(1),
	1334: uint16(anon_sym_AT),
	1335: uint16(149),
	1336: uint16(1),
	1337: uint16(anon_sym_LBRACE),
	1338: uint16(177),
	1339: uint16(1),
	1340: uint16(anon_sym_DQUOTE),
	1341: uint16(78),
	1342: uint16(1),
	1343: uint16(sym_directives),
	1344: uint16(116),
	1345: uint16(1),
	1346: uint16(sym_input_fields_definition),
	1347: uint16(3),
	1348: uint16(2),
	1349: uint16(sym_comment),
	1350: uint16(sym_comma),
	1351: uint16(36),
	1352: uint16(2),
	1353: uint16(sym_directive),
	1354: uint16(aux_sym_directives_repeat1),
	1355: uint16(175),
	1356: uint16(15),
	1358: uint16(anon_sym_schema),
	1359: uint16(anon_sym_extend),
	1360: uint16(anon_sym_scalar),
	1361: uint16(anon_sym_type),
	1362: uint16(anon_sym_interface),
	1363: uint16(anon_sym_union),
	1364: uint16(anon_sym_enum),
	1365: uint16(anon_sym_input),
	1366: uint16(anon_sym_query),
	1367: uint16(anon_sym_mutation),
	1368: uint16(anon_sym_subscription),
	1369: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1370: uint16(anon_sym_fragment),
	1371: uint16(anon_sym_directive),
	1372: uint16(8),
	1373: uint16(87),
	1374: uint16(1),
	1375: uint16(anon_sym_AT),
	1376: uint16(163),
	1377: uint16(1),
	1378: uint16(anon_sym_LBRACE),
	1379: uint16(181),
	1380: uint16(1),
	1381: uint16(anon_sym_DQUOTE),
	1382: uint16(63),
	1383: uint16(1),
	1384: uint16(sym_directives),
	1385: uint16(126),
	1386: uint16(1),
	1387: uint16(sym_enum_values_definition),
	1388: uint16(3),
	1389: uint16(2),
	1390: uint16(sym_comment),
	1391: uint16(sym_comma),
	1392: uint16(36),
	1393: uint16(2),
	1394: uint16(sym_directive),
	1395: uint16(aux_sym_directives_repeat1),
	1396: uint16(179),
	1397: uint16(15),
	1399: uint16(anon_sym_schema),
	1400: uint16(anon_sym_extend),
	1401: uint16(anon_sym_scalar),
	1402: uint16(anon_sym_type),
	1403: uint16(anon_sym_interface),
	1404: uint16(anon_sym_union),
	1405: uint16(anon_sym_enum),
	1406: uint16(anon_sym_input),
	1407: uint16(anon_sym_query),
	1408: uint16(anon_sym_mutation),
	1409: uint16(anon_sym_subscription),
	1410: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1411: uint16(anon_sym_fragment),
	1412: uint16(anon_sym_directive),
	1413: uint16(5),
	1414: uint16(185),
	1415: uint16(1),
	1416: uint16(anon_sym_LPAREN),
	1417: uint16(187),
	1418: uint16(1),
	1419: uint16(anon_sym_DQUOTE),
	1420: uint16(53),
	1421: uint16(1),
	1422: uint16(sym_arguments),
	1423: uint16(3),
	1424: uint16(2),
	1425: uint16(sym_comment),
	1426: uint16(sym_comma),
	1427: uint16(183),
	1428: uint16(18),
	1430: uint16(anon_sym_schema),
	1431: uint16(anon_sym_LBRACE),
	1432: uint16(anon_sym_extend),
	1433: uint16(anon_sym_scalar),
	1434: uint16(anon_sym_type),
	1435: uint16(anon_sym_interface),
	1436: uint16(anon_sym_union),
	1437: uint16(anon_sym_enum),
	1438: uint16(anon_sym_input),
	1439: uint16(anon_sym_EQ),
	1440: uint16(anon_sym_query),
	1441: uint16(anon_sym_mutation),
	1442: uint16(anon_sym_subscription),
	1443: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1444: uint16(anon_sym_fragment),
	1445: uint16(anon_sym_AT),
	1446: uint16(anon_sym_directive),
	1447: uint16(6),
	1448: uint16(87),
	1449: uint16(1),
	1450: uint16(anon_sym_AT),
	1451: uint16(191),
	1452: uint16(1),
	1453: uint16(anon_sym_DQUOTE),
	1454: uint16(125),
	1455: uint16(1),
	1456: uint16(sym_directives),
	1457: uint16(3),
	1458: uint16(2),
	1459: uint16(sym_comment),
	1460: uint16(sym_comma),
	1461: uint16(36),
	1462: uint16(2),
	1463: uint16(sym_directive),
	1464: uint16(aux_sym_directives_repeat1),
	1465: uint16(189),
	1466: uint16(16),
	1468: uint16(anon_sym_schema),
	1469: uint16(anon_sym_LBRACE),
	1470: uint16(anon_sym_extend),
	1471: uint16(anon_sym_scalar),
	1472: uint16(anon_sym_type),
	1473: uint16(anon_sym_interface),
	1474: uint16(anon_sym_union),
	1475: uint16(anon_sym_enum),
	1476: uint16(anon_sym_input),
	1477: uint16(anon_sym_query),
	1478: uint16(anon_sym_mutation),
	1479: uint16(anon_sym_subscription),
	1480: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1481: uint16(anon_sym_fragment),
	1482: uint16(anon_sym_directive),
	1483: uint16(5),
	1484: uint16(195),
	1485: uint16(1),
	1486: uint16(anon_sym_DQUOTE),
	1487: uint16(197),
	1488: uint16(1),
	1489: uint16(anon_sym_AT),
	1490: uint16(3),
	1491: uint16(2),
	1492: uint16(sym_comment),
	1493: uint16(sym_comma),
	1494: uint16(34),
	1495: uint16(2),
	1496: uint16(sym_directive),
	1497: uint16(aux_sym_directives_repeat1),
	1498: uint16(193),
	1499: uint16(17),
	1501: uint16(anon_sym_schema),
	1502: uint16(anon_sym_LBRACE),
	1503: uint16(anon_sym_extend),
	1504: uint16(anon_sym_scalar),
	1505: uint16(anon_sym_type),
	1506: uint16(anon_sym_interface),
	1507: uint16(anon_sym_union),
	1508: uint16(anon_sym_enum),
	1509: uint16(anon_sym_input),
	1510: uint16(anon_sym_EQ),
	1511: uint16(anon_sym_query),
	1512: uint16(anon_sym_mutation),
	1513: uint16(anon_sym_subscription),
	1514: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1515: uint16(anon_sym_fragment),
	1516: uint16(anon_sym_directive),
	1517: uint16(6),
	1518: uint16(87),
	1519: uint16(1),
	1520: uint16(anon_sym_AT),
	1521: uint16(202),
	1522: uint16(1),
	1523: uint16(anon_sym_DQUOTE),
	1524: uint16(108),
	1525: uint16(1),
	1526: uint16(sym_directives),
	1527: uint16(3),
	1528: uint16(2),
	1529: uint16(sym_comment),
	1530: uint16(sym_comma),
	1531: uint16(36),
	1532: uint16(2),
	1533: uint16(sym_directive),
	1534: uint16(aux_sym_directives_repeat1),
	1535: uint16(200),
	1536: uint16(16),
	1538: uint16(anon_sym_schema),
	1539: uint16(anon_sym_LBRACE),
	1540: uint16(anon_sym_extend),
	1541: uint16(anon_sym_scalar),
	1542: uint16(anon_sym_type),
	1543: uint16(anon_sym_interface),
	1544: uint16(anon_sym_union),
	1545: uint16(anon_sym_enum),
	1546: uint16(anon_sym_input),
	1547: uint16(anon_sym_query),
	1548: uint16(anon_sym_mutation),
	1549: uint16(anon_sym_subscription),
	1550: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1551: uint16(anon_sym_fragment),
	1552: uint16(anon_sym_directive),
	1553: uint16(5),
	1554: uint16(87),
	1555: uint16(1),
	1556: uint16(anon_sym_AT),
	1557: uint16(206),
	1558: uint16(1),
	1559: uint16(anon_sym_DQUOTE),
	1560: uint16(3),
	1561: uint16(2),
	1562: uint16(sym_comment),
	1563: uint16(sym_comma),
	1564: uint16(34),
	1565: uint16(2),
	1566: uint16(sym_directive),
	1567: uint16(aux_sym_directives_repeat1),
	1568: uint16(204),
	1569: uint16(17),
	1571: uint16(anon_sym_schema),
	1572: uint16(anon_sym_LBRACE),
	1573: uint16(anon_sym_extend),
	1574: uint16(anon_sym_scalar),
	1575: uint16(anon_sym_type),
	1576: uint16(anon_sym_interface),
	1577: uint16(anon_sym_union),
	1578: uint16(anon_sym_enum),
	1579: uint16(anon_sym_input),
	1580: uint16(anon_sym_EQ),
	1581: uint16(anon_sym_query),
	1582: uint16(anon_sym_mutation),
	1583: uint16(anon_sym_subscription),
	1584: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1585: uint16(anon_sym_fragment),
	1586: uint16(anon_sym_directive),
	1587: uint16(13),
	1588: uint16(208),
	1589: uint16(1),
	1590: uint16(anon_sym_LBRACE),
	1591: uint16(211),
	1592: uint16(1),
	1593: uint16(anon_sym_DOLLAR),
	1594: uint16(214),
	1595: uint16(1),
	1596: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1597: uint16(217),
	1598: uint16(1),
	1599: uint16(anon_sym_DQUOTE),
	1600: uint16(223),
	1601: uint16(1),
	1602: uint16(sym_float_value),
	1603: uint16(229),
	1604: uint16(1),
	1605: uint16(anon_sym_LBRACK),
	1606: uint16(232),
	1607: uint16(1),
	1608: uint16(anon_sym_RBRACK),
	1609: uint16(234),
	1610: uint16(1),
	1611: uint16(sym_name),
	1612: uint16(3),
	1613: uint16(2),
	1614: uint16(sym_comment),
	1615: uint16(sym_comma),
	1616: uint16(220),
	1617: uint16(2),
	1618: uint16(sym_int_value),
	1619: uint16(sym_null_value),
	1620: uint16(226),
	1621: uint16(2),
	1622: uint16(anon_sym_true),
	1623: uint16(anon_sym_false),
	1624: uint16(37),
	1625: uint16(2),
	1626: uint16(sym_value),
	1627: uint16(aux_sym_list_value_repeat1),
	1628: uint16(145),
	1629: uint16(6),
	1630: uint16(sym_variable),
	1631: uint16(sym_string_value),
	1632: uint16(sym_boolean_value),
	1633: uint16(sym_enum_value),
	1634: uint16(sym_list_value),
	1635: uint16(sym_object_value),
	1636: uint16(13),
	1637: uint16(237),
	1638: uint16(1),
	1639: uint16(anon_sym_LBRACE),
	1640: uint16(239),
	1641: uint16(1),
	1642: uint16(anon_sym_DOLLAR),
	1643: uint16(241),
	1644: uint16(1),
	1645: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1646: uint16(243),
	1647: uint16(1),
	1648: uint16(anon_sym_DQUOTE),
	1649: uint16(247),
	1650: uint16(1),
	1651: uint16(sym_float_value),
	1652: uint16(251),
	1653: uint16(1),
	1654: uint16(anon_sym_LBRACK),
	1655: uint16(253),
	1656: uint16(1),
	1657: uint16(anon_sym_RBRACK),
	1658: uint16(255),
	1659: uint16(1),
	1660: uint16(sym_name),
	1661: uint16(3),
	1662: uint16(2),
	1663: uint16(sym_comment),
	1664: uint16(sym_comma),
	1665: uint16(245),
	1666: uint16(2),
	1667: uint16(sym_int_value),
	1668: uint16(sym_null_value),
	1669: uint16(249),
	1670: uint16(2),
	1671: uint16(anon_sym_true),
	1672: uint16(anon_sym_false),
	1673: uint16(37),
	1674: uint16(2),
	1675: uint16(sym_value),
	1676: uint16(aux_sym_list_value_repeat1),
	1677: uint16(145),
	1678: uint16(6),
	1679: uint16(sym_variable),
	1680: uint16(sym_string_value),
	1681: uint16(sym_boolean_value),
	1682: uint16(sym_enum_value),
	1683: uint16(sym_list_value),
	1684: uint16(sym_object_value),
	1685: uint16(13),
	1686: uint16(237),
	1687: uint16(1),
	1688: uint16(anon_sym_LBRACE),
	1689: uint16(239),
	1690: uint16(1),
	1691: uint16(anon_sym_DOLLAR),
	1692: uint16(241),
	1693: uint16(1),
	1694: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1695: uint16(243),
	1696: uint16(1),
	1697: uint16(anon_sym_DQUOTE),
	1698: uint16(247),
	1699: uint16(1),
	1700: uint16(sym_float_value),
	1701: uint16(251),
	1702: uint16(1),
	1703: uint16(anon_sym_LBRACK),
	1704: uint16(255),
	1705: uint16(1),
	1706: uint16(sym_name),
	1707: uint16(257),
	1708: uint16(1),
	1709: uint16(anon_sym_RBRACK),
	1710: uint16(3),
	1711: uint16(2),
	1712: uint16(sym_comment),
	1713: uint16(sym_comma),
	1714: uint16(245),
	1715: uint16(2),
	1716: uint16(sym_int_value),
	1717: uint16(sym_null_value),
	1718: uint16(249),
	1719: uint16(2),
	1720: uint16(anon_sym_true),
	1721: uint16(anon_sym_false),
	1722: uint16(40),
	1723: uint16(2),
	1724: uint16(sym_value),
	1725: uint16(aux_sym_list_value_repeat1),
	1726: uint16(145),
	1727: uint16(6),
	1728: uint16(sym_variable),
	1729: uint16(sym_string_value),
	1730: uint16(sym_boolean_value),
	1731: uint16(sym_enum_value),
	1732: uint16(sym_list_value),
	1733: uint16(sym_object_value),
	1734: uint16(13),
	1735: uint16(237),
	1736: uint16(1),
	1737: uint16(anon_sym_LBRACE),
	1738: uint16(239),
	1739: uint16(1),
	1740: uint16(anon_sym_DOLLAR),
	1741: uint16(241),
	1742: uint16(1),
	1743: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1744: uint16(243),
	1745: uint16(1),
	1746: uint16(anon_sym_DQUOTE),
	1747: uint16(247),
	1748: uint16(1),
	1749: uint16(sym_float_value),
	1750: uint16(251),
	1751: uint16(1),
	1752: uint16(anon_sym_LBRACK),
	1753: uint16(255),
	1754: uint16(1),
	1755: uint16(sym_name),
	1756: uint16(259),
	1757: uint16(1),
	1758: uint16(anon_sym_RBRACK),
	1759: uint16(3),
	1760: uint16(2),
	1761: uint16(sym_comment),
	1762: uint16(sym_comma),
	1763: uint16(245),
	1764: uint16(2),
	1765: uint16(sym_int_value),
	1766: uint16(sym_null_value),
	1767: uint16(249),
	1768: uint16(2),
	1769: uint16(anon_sym_true),
	1770: uint16(anon_sym_false),
	1771: uint16(37),
	1772: uint16(2),
	1773: uint16(sym_value),
	1774: uint16(aux_sym_list_value_repeat1),
	1775: uint16(145),
	1776: uint16(6),
	1777: uint16(sym_variable),
	1778: uint16(sym_string_value),
	1779: uint16(sym_boolean_value),
	1780: uint16(sym_enum_value),
	1781: uint16(sym_list_value),
	1782: uint16(sym_object_value),
	1783: uint16(13),
	1784: uint16(237),
	1785: uint16(1),
	1786: uint16(anon_sym_LBRACE),
	1787: uint16(239),
	1788: uint16(1),
	1789: uint16(anon_sym_DOLLAR),
	1790: uint16(241),
	1791: uint16(1),
	1792: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1793: uint16(243),
	1794: uint16(1),
	1795: uint16(anon_sym_DQUOTE),
	1796: uint16(247),
	1797: uint16(1),
	1798: uint16(sym_float_value),
	1799: uint16(251),
	1800: uint16(1),
	1801: uint16(anon_sym_LBRACK),
	1802: uint16(255),
	1803: uint16(1),
	1804: uint16(sym_name),
	1805: uint16(261),
	1806: uint16(1),
	1807: uint16(anon_sym_RBRACK),
	1808: uint16(3),
	1809: uint16(2),
	1810: uint16(sym_comment),
	1811: uint16(sym_comma),
	1812: uint16(245),
	1813: uint16(2),
	1814: uint16(sym_int_value),
	1815: uint16(sym_null_value),
	1816: uint16(249),
	1817: uint16(2),
	1818: uint16(anon_sym_true),
	1819: uint16(anon_sym_false),
	1820: uint16(42),
	1821: uint16(2),
	1822: uint16(sym_value),
	1823: uint16(aux_sym_list_value_repeat1),
	1824: uint16(145),
	1825: uint16(6),
	1826: uint16(sym_variable),
	1827: uint16(sym_string_value),
	1828: uint16(sym_boolean_value),
	1829: uint16(sym_enum_value),
	1830: uint16(sym_list_value),
	1831: uint16(sym_object_value),
	1832: uint16(13),
	1833: uint16(237),
	1834: uint16(1),
	1835: uint16(anon_sym_LBRACE),
	1836: uint16(239),
	1837: uint16(1),
	1838: uint16(anon_sym_DOLLAR),
	1839: uint16(241),
	1840: uint16(1),
	1841: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1842: uint16(243),
	1843: uint16(1),
	1844: uint16(anon_sym_DQUOTE),
	1845: uint16(247),
	1846: uint16(1),
	1847: uint16(sym_float_value),
	1848: uint16(251),
	1849: uint16(1),
	1850: uint16(anon_sym_LBRACK),
	1851: uint16(255),
	1852: uint16(1),
	1853: uint16(sym_name),
	1854: uint16(263),
	1855: uint16(1),
	1856: uint16(anon_sym_RBRACK),
	1857: uint16(3),
	1858: uint16(2),
	1859: uint16(sym_comment),
	1860: uint16(sym_comma),
	1861: uint16(245),
	1862: uint16(2),
	1863: uint16(sym_int_value),
	1864: uint16(sym_null_value),
	1865: uint16(249),
	1866: uint16(2),
	1867: uint16(anon_sym_true),
	1868: uint16(anon_sym_false),
	1869: uint16(37),
	1870: uint16(2),
	1871: uint16(sym_value),
	1872: uint16(aux_sym_list_value_repeat1),
	1873: uint16(145),
	1874: uint16(6),
	1875: uint16(sym_variable),
	1876: uint16(sym_string_value),
	1877: uint16(sym_boolean_value),
	1878: uint16(sym_enum_value),
	1879: uint16(sym_list_value),
	1880: uint16(sym_object_value),
	1881: uint16(13),
	1882: uint16(237),
	1883: uint16(1),
	1884: uint16(anon_sym_LBRACE),
	1885: uint16(239),
	1886: uint16(1),
	1887: uint16(anon_sym_DOLLAR),
	1888: uint16(241),
	1889: uint16(1),
	1890: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1891: uint16(243),
	1892: uint16(1),
	1893: uint16(anon_sym_DQUOTE),
	1894: uint16(247),
	1895: uint16(1),
	1896: uint16(sym_float_value),
	1897: uint16(251),
	1898: uint16(1),
	1899: uint16(anon_sym_LBRACK),
	1900: uint16(255),
	1901: uint16(1),
	1902: uint16(sym_name),
	1903: uint16(265),
	1904: uint16(1),
	1905: uint16(anon_sym_RBRACK),
	1906: uint16(3),
	1907: uint16(2),
	1908: uint16(sym_comment),
	1909: uint16(sym_comma),
	1910: uint16(245),
	1911: uint16(2),
	1912: uint16(sym_int_value),
	1913: uint16(sym_null_value),
	1914: uint16(249),
	1915: uint16(2),
	1916: uint16(anon_sym_true),
	1917: uint16(anon_sym_false),
	1918: uint16(38),
	1919: uint16(2),
	1920: uint16(sym_value),
	1921: uint16(aux_sym_list_value_repeat1),
	1922: uint16(145),
	1923: uint16(6),
	1924: uint16(sym_variable),
	1925: uint16(sym_string_value),
	1926: uint16(sym_boolean_value),
	1927: uint16(sym_enum_value),
	1928: uint16(sym_list_value),
	1929: uint16(sym_object_value),
	1930: uint16(5),
	1931: uint16(269),
	1932: uint16(1),
	1933: uint16(anon_sym_LBRACE),
	1934: uint16(272),
	1935: uint16(1),
	1936: uint16(anon_sym_DQUOTE),
	1937: uint16(3),
	1938: uint16(2),
	1939: uint16(sym_comment),
	1940: uint16(sym_comma),
	1941: uint16(44),
	1942: uint16(2),
	1943: uint16(sym_input_fields_definition),
	1944: uint16(aux_sym_input_object_type_extension_repeat1),
	1945: uint16(267),
	1946: uint16(15),
	1948: uint16(anon_sym_schema),
	1949: uint16(anon_sym_extend),
	1950: uint16(anon_sym_scalar),
	1951: uint16(anon_sym_type),
	1952: uint16(anon_sym_interface),
	1953: uint16(anon_sym_union),
	1954: uint16(anon_sym_enum),
	1955: uint16(anon_sym_input),
	1956: uint16(anon_sym_query),
	1957: uint16(anon_sym_mutation),
	1958: uint16(anon_sym_subscription),
	1959: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1960: uint16(anon_sym_fragment),
	1961: uint16(anon_sym_directive),
	1962: uint16(5),
	1963: uint16(149),
	1964: uint16(1),
	1965: uint16(anon_sym_LBRACE),
	1966: uint16(276),
	1967: uint16(1),
	1968: uint16(anon_sym_DQUOTE),
	1969: uint16(3),
	1970: uint16(2),
	1971: uint16(sym_comment),
	1972: uint16(sym_comma),
	1973: uint16(44),
	1974: uint16(2),
	1975: uint16(sym_input_fields_definition),
	1976: uint16(aux_sym_input_object_type_extension_repeat1),
	1977: uint16(274),
	1978: uint16(15),
	1980: uint16(anon_sym_schema),
	1981: uint16(anon_sym_extend),
	1982: uint16(anon_sym_scalar),
	1983: uint16(anon_sym_type),
	1984: uint16(anon_sym_interface),
	1985: uint16(anon_sym_union),
	1986: uint16(anon_sym_enum),
	1987: uint16(anon_sym_input),
	1988: uint16(anon_sym_query),
	1989: uint16(anon_sym_mutation),
	1990: uint16(anon_sym_subscription),
	1991: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1992: uint16(anon_sym_fragment),
	1993: uint16(anon_sym_directive),
	1994: uint16(3),
	1995: uint16(280),
	1996: uint16(1),
	1997: uint16(anon_sym_DQUOTE),
	1998: uint16(3),
	1999: uint16(2),
	2000: uint16(sym_comment),
	2001: uint16(sym_comma),
	2002: uint16(278),
	2003: uint16(18),
	2005: uint16(anon_sym_schema),
	2006: uint16(anon_sym_LBRACE),
	2007: uint16(anon_sym_extend),
	2008: uint16(anon_sym_scalar),
	2009: uint16(anon_sym_type),
	2010: uint16(anon_sym_interface),
	2011: uint16(anon_sym_union),
	2012: uint16(anon_sym_enum),
	2013: uint16(anon_sym_input),
	2014: uint16(anon_sym_AMP),
	2015: uint16(anon_sym_query),
	2016: uint16(anon_sym_mutation),
	2017: uint16(anon_sym_subscription),
	2018: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2019: uint16(anon_sym_fragment),
	2020: uint16(anon_sym_AT),
	2021: uint16(anon_sym_directive),
	2022: uint16(5),
	2023: uint16(121),
	2024: uint16(1),
	2025: uint16(anon_sym_EQ),
	2026: uint16(284),
	2027: uint16(1),
	2028: uint16(anon_sym_DQUOTE),
	2029: uint16(69),
	2030: uint16(1),
	2031: uint16(sym_union_member_types),
	2032: uint16(3),
	2033: uint16(2),
	2034: uint16(sym_comment),
	2035: uint16(sym_comma),
	2036: uint16(282),
	2037: uint16(16),
	2039: uint16(anon_sym_schema),
	2040: uint16(anon_sym_LBRACE),
	2041: uint16(anon_sym_extend),
	2042: uint16(anon_sym_scalar),
	2043: uint16(anon_sym_type),
	2044: uint16(anon_sym_interface),
	2045: uint16(anon_sym_union),
	2046: uint16(anon_sym_enum),
	2047: uint16(anon_sym_input),
	2048: uint16(anon_sym_query),
	2049: uint16(anon_sym_mutation),
	2050: uint16(anon_sym_subscription),
	2051: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2052: uint16(anon_sym_fragment),
	2053: uint16(anon_sym_directive),
	2054: uint16(5),
	2055: uint16(149),
	2056: uint16(1),
	2057: uint16(anon_sym_LBRACE),
	2058: uint16(288),
	2059: uint16(1),
	2060: uint16(anon_sym_DQUOTE),
	2061: uint16(3),
	2062: uint16(2),
	2063: uint16(sym_comment),
	2064: uint16(sym_comma),
	2065: uint16(45),
	2066: uint16(2),
	2067: uint16(sym_input_fields_definition),
	2068: uint16(aux_sym_input_object_type_extension_repeat1),
	2069: uint16(286),
	2070: uint16(15),
	2072: uint16(anon_sym_schema),
	2073: uint16(anon_sym_extend),
	2074: uint16(anon_sym_scalar),
	2075: uint16(anon_sym_type),
	2076: uint16(anon_sym_interface),
	2077: uint16(anon_sym_union),
	2078: uint16(anon_sym_enum),
	2079: uint16(anon_sym_input),
	2080: uint16(anon_sym_query),
	2081: uint16(anon_sym_mutation),
	2082: uint16(anon_sym_subscription),
	2083: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2084: uint16(anon_sym_fragment),
	2085: uint16(anon_sym_directive),
	2086: uint16(5),
	2087: uint16(149),
	2088: uint16(1),
	2089: uint16(anon_sym_LBRACE),
	2090: uint16(288),
	2091: uint16(1),
	2092: uint16(anon_sym_DQUOTE),
	2093: uint16(3),
	2094: uint16(2),
	2095: uint16(sym_comment),
	2096: uint16(sym_comma),
	2097: uint16(44),
	2098: uint16(2),
	2099: uint16(sym_input_fields_definition),
	2100: uint16(aux_sym_input_object_type_extension_repeat1),
	2101: uint16(286),
	2102: uint16(15),
	2104: uint16(anon_sym_schema),
	2105: uint16(anon_sym_extend),
	2106: uint16(anon_sym_scalar),
	2107: uint16(anon_sym_type),
	2108: uint16(anon_sym_interface),
	2109: uint16(anon_sym_union),
	2110: uint16(anon_sym_enum),
	2111: uint16(anon_sym_input),
	2112: uint16(anon_sym_query),
	2113: uint16(anon_sym_mutation),
	2114: uint16(anon_sym_subscription),
	2115: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2116: uint16(anon_sym_fragment),
	2117: uint16(anon_sym_directive),
	2118: uint16(5),
	2119: uint16(121),
	2120: uint16(1),
	2121: uint16(anon_sym_EQ),
	2122: uint16(123),
	2123: uint16(1),
	2124: uint16(anon_sym_DQUOTE),
	2125: uint16(56),
	2126: uint16(1),
	2127: uint16(sym_union_member_types),
	2128: uint16(3),
	2129: uint16(2),
	2130: uint16(sym_comment),
	2131: uint16(sym_comma),
	2132: uint16(119),
	2133: uint16(16),
	2135: uint16(anon_sym_schema),
	2136: uint16(anon_sym_LBRACE),
	2137: uint16(anon_sym_extend),
	2138: uint16(anon_sym_scalar),
	2139: uint16(anon_sym_type),
	2140: uint16(anon_sym_interface),
	2141: uint16(anon_sym_union),
	2142: uint16(anon_sym_enum),
	2143: uint16(anon_sym_input),
	2144: uint16(anon_sym_query),
	2145: uint16(anon_sym_mutation),
	2146: uint16(anon_sym_subscription),
	2147: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2148: uint16(anon_sym_fragment),
	2149: uint16(anon_sym_directive),
	2150: uint16(3),
	2151: uint16(292),
	2152: uint16(1),
	2153: uint16(anon_sym_DQUOTE),
	2154: uint16(3),
	2155: uint16(2),
	2156: uint16(sym_comment),
	2157: uint16(sym_comma),
	2158: uint16(290),
	2159: uint16(18),
	2161: uint16(anon_sym_schema),
	2162: uint16(anon_sym_LBRACE),
	2163: uint16(anon_sym_extend),
	2164: uint16(anon_sym_scalar),
	2165: uint16(anon_sym_type),
	2166: uint16(anon_sym_interface),
	2167: uint16(anon_sym_union),
	2168: uint16(anon_sym_enum),
	2169: uint16(anon_sym_input),
	2170: uint16(anon_sym_AMP),
	2171: uint16(anon_sym_query),
	2172: uint16(anon_sym_mutation),
	2173: uint16(anon_sym_subscription),
	2174: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2175: uint16(anon_sym_fragment),
	2176: uint16(anon_sym_AT),
	2177: uint16(anon_sym_directive),
	2178: uint16(3),
	2179: uint16(296),
	2180: uint16(1),
	2181: uint16(anon_sym_DQUOTE),
	2182: uint16(3),
	2183: uint16(2),
	2184: uint16(sym_comment),
	2185: uint16(sym_comma),
	2186: uint16(294),
	2187: uint16(18),
	2189: uint16(anon_sym_schema),
	2190: uint16(anon_sym_LBRACE),
	2191: uint16(anon_sym_extend),
	2192: uint16(anon_sym_scalar),
	2193: uint16(anon_sym_type),
	2194: uint16(anon_sym_interface),
	2195: uint16(anon_sym_union),
	2196: uint16(anon_sym_enum),
	2197: uint16(anon_sym_input),
	2198: uint16(anon_sym_EQ),
	2199: uint16(anon_sym_query),
	2200: uint16(anon_sym_mutation),
	2201: uint16(anon_sym_subscription),
	2202: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2203: uint16(anon_sym_fragment),
	2204: uint16(anon_sym_AT),
	2205: uint16(anon_sym_directive),
	2206: uint16(3),
	2207: uint16(300),
	2208: uint16(1),
	2209: uint16(anon_sym_DQUOTE),
	2210: uint16(3),
	2211: uint16(2),
	2212: uint16(sym_comment),
	2213: uint16(sym_comma),
	2214: uint16(298),
	2215: uint16(18),
	2217: uint16(anon_sym_schema),
	2218: uint16(anon_sym_LBRACE),
	2219: uint16(anon_sym_extend),
	2220: uint16(anon_sym_scalar),
	2221: uint16(anon_sym_type),
	2222: uint16(anon_sym_interface),
	2223: uint16(anon_sym_union),
	2224: uint16(anon_sym_enum),
	2225: uint16(anon_sym_input),
	2226: uint16(anon_sym_EQ),
	2227: uint16(anon_sym_query),
	2228: uint16(anon_sym_mutation),
	2229: uint16(anon_sym_subscription),
	2230: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2231: uint16(anon_sym_fragment),
	2232: uint16(anon_sym_AT),
	2233: uint16(anon_sym_directive),
	2234: uint16(5),
	2235: uint16(121),
	2236: uint16(1),
	2237: uint16(anon_sym_EQ),
	2238: uint16(304),
	2239: uint16(1),
	2240: uint16(anon_sym_DQUOTE),
	2241: uint16(79),
	2242: uint16(1),
	2243: uint16(sym_union_member_types),
	2244: uint16(3),
	2245: uint16(2),
	2246: uint16(sym_comment),
	2247: uint16(sym_comma),
	2248: uint16(302),
	2249: uint16(16),
	2251: uint16(anon_sym_schema),
	2252: uint16(anon_sym_LBRACE),
	2253: uint16(anon_sym_extend),
	2254: uint16(anon_sym_scalar),
	2255: uint16(anon_sym_type),
	2256: uint16(anon_sym_interface),
	2257: uint16(anon_sym_union),
	2258: uint16(anon_sym_enum),
	2259: uint16(anon_sym_input),
	2260: uint16(anon_sym_query),
	2261: uint16(anon_sym_mutation),
	2262: uint16(anon_sym_subscription),
	2263: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2264: uint16(anon_sym_fragment),
	2265: uint16(anon_sym_directive),
	2266: uint16(12),
	2267: uint16(306),
	2268: uint16(1),
	2269: uint16(anon_sym_LBRACE),
	2270: uint16(308),
	2271: uint16(1),
	2272: uint16(anon_sym_DOLLAR),
	2273: uint16(310),
	2274: uint16(1),
	2275: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2276: uint16(312),
	2277: uint16(1),
	2278: uint16(anon_sym_DQUOTE),
	2279: uint16(316),
	2280: uint16(1),
	2281: uint16(sym_float_value),
	2282: uint16(320),
	2283: uint16(1),
	2284: uint16(anon_sym_LBRACK),
	2285: uint16(322),
	2286: uint16(1),
	2287: uint16(sym_name),
	2288: uint16(317),
	2289: uint16(1),
	2290: uint16(sym_value),
	2291: uint16(3),
	2292: uint16(2),
	2293: uint16(sym_comment),
	2294: uint16(sym_comma),
	2295: uint16(314),
	2296: uint16(2),
	2297: uint16(sym_int_value),
	2298: uint16(sym_null_value),
	2299: uint16(318),
	2300: uint16(2),
	2301: uint16(anon_sym_true),
	2302: uint16(anon_sym_false),
	2303: uint16(205),
	2304: uint16(6),
	2305: uint16(sym_variable),
	2306: uint16(sym_string_value),
	2307: uint16(sym_boolean_value),
	2308: uint16(sym_enum_value),
	2309: uint16(sym_list_value),
	2310: uint16(sym_object_value),
	2311: uint16(4),
	2312: uint16(304),
	2313: uint16(1),
	2314: uint16(anon_sym_DQUOTE),
	2315: uint16(324),
	2316: uint16(1),
	2317: uint16(anon_sym_PIPE),
	2318: uint16(3),
	2319: uint16(2),
	2320: uint16(sym_comment),
	2321: uint16(sym_comma),
	2322: uint16(302),
	2323: uint16(16),
	2325: uint16(anon_sym_schema),
	2326: uint16(anon_sym_LBRACE),
	2327: uint16(anon_sym_extend),
	2328: uint16(anon_sym_scalar),
	2329: uint16(anon_sym_type),
	2330: uint16(anon_sym_interface),
	2331: uint16(anon_sym_union),
	2332: uint16(anon_sym_enum),
	2333: uint16(anon_sym_input),
	2334: uint16(anon_sym_query),
	2335: uint16(anon_sym_mutation),
	2336: uint16(anon_sym_subscription),
	2337: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2338: uint16(anon_sym_fragment),
	2339: uint16(anon_sym_directive),
	2340: uint16(3),
	2341: uint16(328),
	2342: uint16(1),
	2343: uint16(anon_sym_DQUOTE),
	2344: uint16(3),
	2345: uint16(2),
	2346: uint16(sym_comment),
	2347: uint16(sym_comma),
	2348: uint16(326),
	2349: uint16(17),
	2351: uint16(anon_sym_schema),
	2352: uint16(anon_sym_LBRACE),
	2353: uint16(anon_sym_extend),
	2354: uint16(anon_sym_scalar),
	2355: uint16(anon_sym_type),
	2356: uint16(anon_sym_interface),
	2357: uint16(anon_sym_union),
	2358: uint16(anon_sym_enum),
	2359: uint16(anon_sym_input),
	2360: uint16(anon_sym_PIPE),
	2361: uint16(anon_sym_query),
	2362: uint16(anon_sym_mutation),
	2363: uint16(anon_sym_subscription),
	2364: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2365: uint16(anon_sym_fragment),
	2366: uint16(anon_sym_directive),
	2367: uint16(5),
	2368: uint16(81),
	2369: uint16(1),
	2370: uint16(anon_sym_LBRACE),
	2371: uint16(159),
	2372: uint16(1),
	2373: uint16(anon_sym_DQUOTE),
	2374: uint16(101),
	2375: uint16(1),
	2376: uint16(sym_fields_definition),
	2377: uint16(3),
	2378: uint16(2),
	2379: uint16(sym_comment),
	2380: uint16(sym_comma),
	2381: uint16(157),
	2382: uint16(15),
	2384: uint16(anon_sym_schema),
	2385: uint16(anon_sym_extend),
	2386: uint16(anon_sym_scalar),
	2387: uint16(anon_sym_type),
	2388: uint16(anon_sym_interface),
	2389: uint16(anon_sym_union),
	2390: uint16(anon_sym_enum),
	2391: uint16(anon_sym_input),
	2392: uint16(anon_sym_query),
	2393: uint16(anon_sym_mutation),
	2394: uint16(anon_sym_subscription),
	2395: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2396: uint16(anon_sym_fragment),
	2397: uint16(anon_sym_directive),
	2398: uint16(5),
	2399: uint16(81),
	2400: uint16(1),
	2401: uint16(anon_sym_LBRACE),
	2402: uint16(133),
	2403: uint16(1),
	2404: uint16(anon_sym_DQUOTE),
	2405: uint16(102),
	2406: uint16(1),
	2407: uint16(sym_fields_definition),
	2408: uint16(3),
	2409: uint16(2),
	2410: uint16(sym_comment),
	2411: uint16(sym_comma),
	2412: uint16(129),
	2413: uint16(15),
	2415: uint16(anon_sym_schema),
	2416: uint16(anon_sym_extend),
	2417: uint16(anon_sym_scalar),
	2418: uint16(anon_sym_type),
	2419: uint16(anon_sym_interface),
	2420: uint16(anon_sym_union),
	2421: uint16(anon_sym_enum),
	2422: uint16(anon_sym_input),
	2423: uint16(anon_sym_query),
	2424: uint16(anon_sym_mutation),
	2425: uint16(anon_sym_subscription),
	2426: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2427: uint16(anon_sym_fragment),
	2428: uint16(anon_sym_directive),
	2429: uint16(5),
	2430: uint16(163),
	2431: uint16(1),
	2432: uint16(anon_sym_LBRACE),
	2433: uint16(332),
	2434: uint16(1),
	2435: uint16(anon_sym_DQUOTE),
	2436: uint16(94),
	2437: uint16(1),
	2438: uint16(sym_enum_values_definition),
	2439: uint16(3),
	2440: uint16(2),
	2441: uint16(sym_comment),
	2442: uint16(sym_comma),
	2443: uint16(330),
	2444: uint16(15),
	2446: uint16(anon_sym_schema),
	2447: uint16(anon_sym_extend),
	2448: uint16(anon_sym_scalar),
	2449: uint16(anon_sym_type),
	2450: uint16(anon_sym_interface),
	2451: uint16(anon_sym_union),
	2452: uint16(anon_sym_enum),
	2453: uint16(anon_sym_input),
	2454: uint16(anon_sym_query),
	2455: uint16(anon_sym_mutation),
	2456: uint16(anon_sym_subscription),
	2457: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2458: uint16(anon_sym_fragment),
	2459: uint16(anon_sym_directive),
	2460: uint16(12),
	2461: uint16(306),
	2462: uint16(1),
	2463: uint16(anon_sym_LBRACE),
	2464: uint16(308),
	2465: uint16(1),
	2466: uint16(anon_sym_DOLLAR),
	2467: uint16(310),
	2468: uint16(1),
	2469: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2470: uint16(312),
	2471: uint16(1),
	2472: uint16(anon_sym_DQUOTE),
	2473: uint16(316),
	2474: uint16(1),
	2475: uint16(sym_float_value),
	2476: uint16(320),
	2477: uint16(1),
	2478: uint16(anon_sym_LBRACK),
	2479: uint16(322),
	2480: uint16(1),
	2481: uint16(sym_name),
	2482: uint16(223),
	2483: uint16(1),
	2484: uint16(sym_value),
	2485: uint16(3),
	2486: uint16(2),
	2487: uint16(sym_comment),
	2488: uint16(sym_comma),
	2489: uint16(314),
	2490: uint16(2),
	2491: uint16(sym_int_value),
	2492: uint16(sym_null_value),
	2493: uint16(318),
	2494: uint16(2),
	2495: uint16(anon_sym_true),
	2496: uint16(anon_sym_false),
	2497: uint16(205),
	2498: uint16(6),
	2499: uint16(sym_variable),
	2500: uint16(sym_string_value),
	2501: uint16(sym_boolean_value),
	2502: uint16(sym_enum_value),
	2503: uint16(sym_list_value),
	2504: uint16(sym_object_value),
	2505: uint16(5),
	2506: uint16(149),
	2507: uint16(1),
	2508: uint16(anon_sym_LBRACE),
	2509: uint16(336),
	2510: uint16(1),
	2511: uint16(anon_sym_DQUOTE),
	2512: uint16(93),
	2513: uint16(1),
	2514: uint16(sym_input_fields_definition),
	2515: uint16(3),
	2516: uint16(2),
	2517: uint16(sym_comment),
	2518: uint16(sym_comma),
	2519: uint16(334),
	2520: uint16(15),
	2522: uint16(anon_sym_schema),
	2523: uint16(anon_sym_extend),
	2524: uint16(anon_sym_scalar),
	2525: uint16(anon_sym_type),
	2526: uint16(anon_sym_interface),
	2527: uint16(anon_sym_union),
	2528: uint16(anon_sym_enum),
	2529: uint16(anon_sym_input),
	2530: uint16(anon_sym_query),
	2531: uint16(anon_sym_mutation),
	2532: uint16(anon_sym_subscription),
	2533: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2534: uint16(anon_sym_fragment),
	2535: uint16(anon_sym_directive),
	2536: uint16(5),
	2537: uint16(163),
	2538: uint16(1),
	2539: uint16(anon_sym_LBRACE),
	2540: uint16(340),
	2541: uint16(1),
	2542: uint16(anon_sym_DQUOTE),
	2543: uint16(105),
	2544: uint16(1),
	2545: uint16(sym_enum_values_definition),
	2546: uint16(3),
	2547: uint16(2),
	2548: uint16(sym_comment),
	2549: uint16(sym_comma),
	2550: uint16(338),
	2551: uint16(15),
	2553: uint16(anon_sym_schema),
	2554: uint16(anon_sym_extend),
	2555: uint16(anon_sym_scalar),
	2556: uint16(anon_sym_type),
	2557: uint16(anon_sym_interface),
	2558: uint16(anon_sym_union),
	2559: uint16(anon_sym_enum),
	2560: uint16(anon_sym_input),
	2561: uint16(anon_sym_query),
	2562: uint16(anon_sym_mutation),
	2563: uint16(anon_sym_subscription),
	2564: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2565: uint16(anon_sym_fragment),
	2566: uint16(anon_sym_directive),
	2567: uint16(5),
	2568: uint16(81),
	2569: uint16(1),
	2570: uint16(anon_sym_LBRACE),
	2571: uint16(344),
	2572: uint16(1),
	2573: uint16(anon_sym_DQUOTE),
	2574: uint16(98),
	2575: uint16(1),
	2576: uint16(sym_fields_definition),
	2577: uint16(3),
	2578: uint16(2),
	2579: uint16(sym_comment),
	2580: uint16(sym_comma),
	2581: uint16(342),
	2582: uint16(15),
	2584: uint16(anon_sym_schema),
	2585: uint16(anon_sym_extend),
	2586: uint16(anon_sym_scalar),
	2587: uint16(anon_sym_type),
	2588: uint16(anon_sym_interface),
	2589: uint16(anon_sym_union),
	2590: uint16(anon_sym_enum),
	2591: uint16(anon_sym_input),
	2592: uint16(anon_sym_query),
	2593: uint16(anon_sym_mutation),
	2594: uint16(anon_sym_subscription),
	2595: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2596: uint16(anon_sym_fragment),
	2597: uint16(anon_sym_directive),
	2598: uint16(4),
	2599: uint16(284),
	2600: uint16(1),
	2601: uint16(anon_sym_DQUOTE),
	2602: uint16(324),
	2603: uint16(1),
	2604: uint16(anon_sym_PIPE),
	2605: uint16(3),
	2606: uint16(2),
	2607: uint16(sym_comment),
	2608: uint16(sym_comma),
	2609: uint16(282),
	2610: uint16(16),
	2612: uint16(anon_sym_schema),
	2613: uint16(anon_sym_LBRACE),
	2614: uint16(anon_sym_extend),
	2615: uint16(anon_sym_scalar),
	2616: uint16(anon_sym_type),
	2617: uint16(anon_sym_interface),
	2618: uint16(anon_sym_union),
	2619: uint16(anon_sym_enum),
	2620: uint16(anon_sym_input),
	2621: uint16(anon_sym_query),
	2622: uint16(anon_sym_mutation),
	2623: uint16(anon_sym_subscription),
	2624: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2625: uint16(anon_sym_fragment),
	2626: uint16(anon_sym_directive),
	2627: uint16(5),
	2628: uint16(81),
	2629: uint16(1),
	2630: uint16(anon_sym_LBRACE),
	2631: uint16(145),
	2632: uint16(1),
	2633: uint16(anon_sym_DQUOTE),
	2634: uint16(109),
	2635: uint16(1),
	2636: uint16(sym_fields_definition),
	2637: uint16(3),
	2638: uint16(2),
	2639: uint16(sym_comment),
	2640: uint16(sym_comma),
	2641: uint16(143),
	2642: uint16(15),
	2644: uint16(anon_sym_schema),
	2645: uint16(anon_sym_extend),
	2646: uint16(anon_sym_scalar),
	2647: uint16(anon_sym_type),
	2648: uint16(anon_sym_interface),
	2649: uint16(anon_sym_union),
	2650: uint16(anon_sym_enum),
	2651: uint16(anon_sym_input),
	2652: uint16(anon_sym_query),
	2653: uint16(anon_sym_mutation),
	2654: uint16(anon_sym_subscription),
	2655: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2656: uint16(anon_sym_fragment),
	2657: uint16(anon_sym_directive),
	2658: uint16(5),
	2659: uint16(81),
	2660: uint16(1),
	2661: uint16(anon_sym_LBRACE),
	2662: uint16(348),
	2663: uint16(1),
	2664: uint16(anon_sym_DQUOTE),
	2665: uint16(100),
	2666: uint16(1),
	2667: uint16(sym_fields_definition),
	2668: uint16(3),
	2669: uint16(2),
	2670: uint16(sym_comment),
	2671: uint16(sym_comma),
	2672: uint16(346),
	2673: uint16(15),
	2675: uint16(anon_sym_schema),
	2676: uint16(anon_sym_extend),
	2677: uint16(anon_sym_scalar),
	2678: uint16(anon_sym_type),
	2679: uint16(anon_sym_interface),
	2680: uint16(anon_sym_union),
	2681: uint16(anon_sym_enum),
	2682: uint16(anon_sym_input),
	2683: uint16(anon_sym_query),
	2684: uint16(anon_sym_mutation),
	2685: uint16(anon_sym_subscription),
	2686: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2687: uint16(anon_sym_fragment),
	2688: uint16(anon_sym_directive),
	2689: uint16(5),
	2690: uint16(81),
	2691: uint16(1),
	2692: uint16(anon_sym_LBRACE),
	2693: uint16(137),
	2694: uint16(1),
	2695: uint16(anon_sym_DQUOTE),
	2696: uint16(112),
	2697: uint16(1),
	2698: uint16(sym_fields_definition),
	2699: uint16(3),
	2700: uint16(2),
	2701: uint16(sym_comment),
	2702: uint16(sym_comma),
	2703: uint16(135),
	2704: uint16(15),
	2706: uint16(anon_sym_schema),
	2707: uint16(anon_sym_extend),
	2708: uint16(anon_sym_scalar),
	2709: uint16(anon_sym_type),
	2710: uint16(anon_sym_interface),
	2711: uint16(anon_sym_union),
	2712: uint16(anon_sym_enum),
	2713: uint16(anon_sym_input),
	2714: uint16(anon_sym_query),
	2715: uint16(anon_sym_mutation),
	2716: uint16(anon_sym_subscription),
	2717: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2718: uint16(anon_sym_fragment),
	2719: uint16(anon_sym_directive),
	2720: uint16(4),
	2721: uint16(324),
	2722: uint16(1),
	2723: uint16(anon_sym_PIPE),
	2724: uint16(352),
	2725: uint16(1),
	2726: uint16(anon_sym_DQUOTE),
	2727: uint16(3),
	2728: uint16(2),
	2729: uint16(sym_comment),
	2730: uint16(sym_comma),
	2731: uint16(350),
	2732: uint16(16),
	2734: uint16(anon_sym_schema),
	2735: uint16(anon_sym_LBRACE),
	2736: uint16(anon_sym_extend),
	2737: uint16(anon_sym_scalar),
	2738: uint16(anon_sym_type),
	2739: uint16(anon_sym_interface),
	2740: uint16(anon_sym_union),
	2741: uint16(anon_sym_enum),
	2742: uint16(anon_sym_input),
	2743: uint16(anon_sym_query),
	2744: uint16(anon_sym_mutation),
	2745: uint16(anon_sym_subscription),
	2746: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2747: uint16(anon_sym_fragment),
	2748: uint16(anon_sym_directive),
	2749: uint16(3),
	2750: uint16(356),
	2751: uint16(1),
	2752: uint16(anon_sym_DQUOTE),
	2753: uint16(3),
	2754: uint16(2),
	2755: uint16(sym_comment),
	2756: uint16(sym_comma),
	2757: uint16(354),
	2758: uint16(17),
	2760: uint16(anon_sym_schema),
	2761: uint16(anon_sym_LBRACE),
	2762: uint16(anon_sym_extend),
	2763: uint16(anon_sym_scalar),
	2764: uint16(anon_sym_type),
	2765: uint16(anon_sym_interface),
	2766: uint16(anon_sym_union),
	2767: uint16(anon_sym_enum),
	2768: uint16(anon_sym_input),
	2769: uint16(anon_sym_PIPE),
	2770: uint16(anon_sym_query),
	2771: uint16(anon_sym_mutation),
	2772: uint16(anon_sym_subscription),
	2773: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2774: uint16(anon_sym_fragment),
	2775: uint16(anon_sym_directive),
	2776: uint16(3),
	2777: uint16(360),
	2778: uint16(1),
	2779: uint16(anon_sym_DQUOTE),
	2780: uint16(3),
	2781: uint16(2),
	2782: uint16(sym_comment),
	2783: uint16(sym_comma),
	2784: uint16(358),
	2785: uint16(17),
	2787: uint16(anon_sym_schema),
	2788: uint16(anon_sym_LBRACE),
	2789: uint16(anon_sym_extend),
	2790: uint16(anon_sym_scalar),
	2791: uint16(anon_sym_type),
	2792: uint16(anon_sym_interface),
	2793: uint16(anon_sym_union),
	2794: uint16(anon_sym_enum),
	2795: uint16(anon_sym_input),
	2796: uint16(anon_sym_PIPE),
	2797: uint16(anon_sym_query),
	2798: uint16(anon_sym_mutation),
	2799: uint16(anon_sym_subscription),
	2800: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2801: uint16(anon_sym_fragment),
	2802: uint16(anon_sym_directive),
	2803: uint16(3),
	2804: uint16(364),
	2805: uint16(1),
	2806: uint16(anon_sym_DQUOTE),
	2807: uint16(3),
	2808: uint16(2),
	2809: uint16(sym_comment),
	2810: uint16(sym_comma),
	2811: uint16(362),
	2812: uint16(17),
	2814: uint16(anon_sym_schema),
	2815: uint16(anon_sym_LBRACE),
	2816: uint16(anon_sym_extend),
	2817: uint16(anon_sym_scalar),
	2818: uint16(anon_sym_type),
	2819: uint16(anon_sym_interface),
	2820: uint16(anon_sym_union),
	2821: uint16(anon_sym_enum),
	2822: uint16(anon_sym_input),
	2823: uint16(anon_sym_PIPE),
	2824: uint16(anon_sym_query),
	2825: uint16(anon_sym_mutation),
	2826: uint16(anon_sym_subscription),
	2827: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2828: uint16(anon_sym_fragment),
	2829: uint16(anon_sym_directive),
	2830: uint16(4),
	2831: uint16(368),
	2832: uint16(1),
	2833: uint16(anon_sym_PIPE),
	2834: uint16(370),
	2835: uint16(1),
	2836: uint16(anon_sym_DQUOTE),
	2837: uint16(3),
	2838: uint16(2),
	2839: uint16(sym_comment),
	2840: uint16(sym_comma),
	2841: uint16(366),
	2842: uint16(16),
	2844: uint16(anon_sym_schema),
	2845: uint16(anon_sym_LBRACE),
	2846: uint16(anon_sym_extend),
	2847: uint16(anon_sym_scalar),
	2848: uint16(anon_sym_type),
	2849: uint16(anon_sym_interface),
	2850: uint16(anon_sym_union),
	2851: uint16(anon_sym_enum),
	2852: uint16(anon_sym_input),
	2853: uint16(anon_sym_query),
	2854: uint16(anon_sym_mutation),
	2855: uint16(anon_sym_subscription),
	2856: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2857: uint16(anon_sym_fragment),
	2858: uint16(anon_sym_directive),
	2859: uint16(3),
	2860: uint16(374),
	2861: uint16(1),
	2862: uint16(anon_sym_DQUOTE),
	2863: uint16(3),
	2864: uint16(2),
	2865: uint16(sym_comment),
	2866: uint16(sym_comma),
	2867: uint16(372),
	2868: uint16(17),
	2870: uint16(anon_sym_schema),
	2871: uint16(anon_sym_LBRACE),
	2872: uint16(anon_sym_extend),
	2873: uint16(anon_sym_scalar),
	2874: uint16(anon_sym_type),
	2875: uint16(anon_sym_interface),
	2876: uint16(anon_sym_union),
	2877: uint16(anon_sym_enum),
	2878: uint16(anon_sym_input),
	2879: uint16(anon_sym_PIPE),
	2880: uint16(anon_sym_query),
	2881: uint16(anon_sym_mutation),
	2882: uint16(anon_sym_subscription),
	2883: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2884: uint16(anon_sym_fragment),
	2885: uint16(anon_sym_directive),
	2886: uint16(3),
	2887: uint16(378),
	2888: uint16(1),
	2889: uint16(anon_sym_DQUOTE),
	2890: uint16(3),
	2891: uint16(2),
	2892: uint16(sym_comment),
	2893: uint16(sym_comma),
	2894: uint16(376),
	2895: uint16(17),
	2897: uint16(anon_sym_schema),
	2898: uint16(anon_sym_LBRACE),
	2899: uint16(anon_sym_extend),
	2900: uint16(anon_sym_scalar),
	2901: uint16(anon_sym_type),
	2902: uint16(anon_sym_interface),
	2903: uint16(anon_sym_union),
	2904: uint16(anon_sym_enum),
	2905: uint16(anon_sym_input),
	2906: uint16(anon_sym_PIPE),
	2907: uint16(anon_sym_query),
	2908: uint16(anon_sym_mutation),
	2909: uint16(anon_sym_subscription),
	2910: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2911: uint16(anon_sym_fragment),
	2912: uint16(anon_sym_directive),
	2913: uint16(5),
	2914: uint16(81),
	2915: uint16(1),
	2916: uint16(anon_sym_LBRACE),
	2917: uint16(382),
	2918: uint16(1),
	2919: uint16(anon_sym_DQUOTE),
	2920: uint16(119),
	2921: uint16(1),
	2922: uint16(sym_fields_definition),
	2923: uint16(3),
	2924: uint16(2),
	2925: uint16(sym_comment),
	2926: uint16(sym_comma),
	2927: uint16(380),
	2928: uint16(15),
	2930: uint16(anon_sym_schema),
	2931: uint16(anon_sym_extend),
	2932: uint16(anon_sym_scalar),
	2933: uint16(anon_sym_type),
	2934: uint16(anon_sym_interface),
	2935: uint16(anon_sym_union),
	2936: uint16(anon_sym_enum),
	2937: uint16(anon_sym_input),
	2938: uint16(anon_sym_query),
	2939: uint16(anon_sym_mutation),
	2940: uint16(anon_sym_subscription),
	2941: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2942: uint16(anon_sym_fragment),
	2943: uint16(anon_sym_directive),
	2944: uint16(5),
	2945: uint16(81),
	2946: uint16(1),
	2947: uint16(anon_sym_LBRACE),
	2948: uint16(386),
	2949: uint16(1),
	2950: uint16(anon_sym_DQUOTE),
	2951: uint16(92),
	2952: uint16(1),
	2953: uint16(sym_fields_definition),
	2954: uint16(3),
	2955: uint16(2),
	2956: uint16(sym_comment),
	2957: uint16(sym_comma),
	2958: uint16(384),
	2959: uint16(15),
	2961: uint16(anon_sym_schema),
	2962: uint16(anon_sym_extend),
	2963: uint16(anon_sym_scalar),
	2964: uint16(anon_sym_type),
	2965: uint16(anon_sym_interface),
	2966: uint16(anon_sym_union),
	2967: uint16(anon_sym_enum),
	2968: uint16(anon_sym_input),
	2969: uint16(anon_sym_query),
	2970: uint16(anon_sym_mutation),
	2971: uint16(anon_sym_subscription),
	2972: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2973: uint16(anon_sym_fragment),
	2974: uint16(anon_sym_directive),
	2975: uint16(5),
	2976: uint16(149),
	2977: uint16(1),
	2978: uint16(anon_sym_LBRACE),
	2979: uint16(169),
	2980: uint16(1),
	2981: uint16(anon_sym_DQUOTE),
	2982: uint16(133),
	2983: uint16(1),
	2984: uint16(sym_input_fields_definition),
	2985: uint16(3),
	2986: uint16(2),
	2987: uint16(sym_comment),
	2988: uint16(sym_comma),
	2989: uint16(167),
	2990: uint16(15),
	2992: uint16(anon_sym_schema),
	2993: uint16(anon_sym_extend),
	2994: uint16(anon_sym_scalar),
	2995: uint16(anon_sym_type),
	2996: uint16(anon_sym_interface),
	2997: uint16(anon_sym_union),
	2998: uint16(anon_sym_enum),
	2999: uint16(anon_sym_input),
	3000: uint16(anon_sym_query),
	3001: uint16(anon_sym_mutation),
	3002: uint16(anon_sym_subscription),
	3003: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3004: uint16(anon_sym_fragment),
	3005: uint16(anon_sym_directive),
	3006: uint16(4),
	3007: uint16(324),
	3008: uint16(1),
	3009: uint16(anon_sym_PIPE),
	3010: uint16(390),
	3011: uint16(1),
	3012: uint16(anon_sym_DQUOTE),
	3013: uint16(3),
	3014: uint16(2),
	3015: uint16(sym_comment),
	3016: uint16(sym_comma),
	3017: uint16(388),
	3018: uint16(16),
	3020: uint16(anon_sym_schema),
	3021: uint16(anon_sym_LBRACE),
	3022: uint16(anon_sym_extend),
	3023: uint16(anon_sym_scalar),
	3024: uint16(anon_sym_type),
	3025: uint16(anon_sym_interface),
	3026: uint16(anon_sym_union),
	3027: uint16(anon_sym_enum),
	3028: uint16(anon_sym_input),
	3029: uint16(anon_sym_query),
	3030: uint16(anon_sym_mutation),
	3031: uint16(anon_sym_subscription),
	3032: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3033: uint16(anon_sym_fragment),
	3034: uint16(anon_sym_directive),
	3035: uint16(3),
	3036: uint16(394),
	3037: uint16(1),
	3038: uint16(anon_sym_DQUOTE),
	3039: uint16(3),
	3040: uint16(2),
	3041: uint16(sym_comment),
	3042: uint16(sym_comma),
	3043: uint16(392),
	3044: uint16(17),
	3046: uint16(anon_sym_schema),
	3047: uint16(anon_sym_LBRACE),
	3048: uint16(anon_sym_extend),
	3049: uint16(anon_sym_scalar),
	3050: uint16(anon_sym_type),
	3051: uint16(anon_sym_interface),
	3052: uint16(anon_sym_union),
	3053: uint16(anon_sym_enum),
	3054: uint16(anon_sym_input),
	3055: uint16(anon_sym_PIPE),
	3056: uint16(anon_sym_query),
	3057: uint16(anon_sym_mutation),
	3058: uint16(anon_sym_subscription),
	3059: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3060: uint16(anon_sym_fragment),
	3061: uint16(anon_sym_directive),
	3062: uint16(5),
	3063: uint16(81),
	3064: uint16(1),
	3065: uint16(anon_sym_LBRACE),
	3066: uint16(105),
	3067: uint16(1),
	3068: uint16(anon_sym_DQUOTE),
	3069: uint16(128),
	3070: uint16(1),
	3071: uint16(sym_fields_definition),
	3072: uint16(3),
	3073: uint16(2),
	3074: uint16(sym_comment),
	3075: uint16(sym_comma),
	3076: uint16(103),
	3077: uint16(15),
	3079: uint16(anon_sym_schema),
	3080: uint16(anon_sym_extend),
	3081: uint16(anon_sym_scalar),
	3082: uint16(anon_sym_type),
	3083: uint16(anon_sym_interface),
	3084: uint16(anon_sym_union),
	3085: uint16(anon_sym_enum),
	3086: uint16(anon_sym_input),
	3087: uint16(anon_sym_query),
	3088: uint16(anon_sym_mutation),
	3089: uint16(anon_sym_subscription),
	3090: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3091: uint16(anon_sym_fragment),
	3092: uint16(anon_sym_directive),
	3093: uint16(4),
	3094: uint16(368),
	3095: uint16(1),
	3096: uint16(anon_sym_PIPE),
	3097: uint16(398),
	3098: uint16(1),
	3099: uint16(anon_sym_DQUOTE),
	3100: uint16(3),
	3101: uint16(2),
	3102: uint16(sym_comment),
	3103: uint16(sym_comma),
	3104: uint16(396),
	3105: uint16(16),
	3107: uint16(anon_sym_schema),
	3108: uint16(anon_sym_LBRACE),
	3109: uint16(anon_sym_extend),
	3110: uint16(anon_sym_scalar),
	3111: uint16(anon_sym_type),
	3112: uint16(anon_sym_interface),
	3113: uint16(anon_sym_union),
	3114: uint16(anon_sym_enum),
	3115: uint16(anon_sym_input),
	3116: uint16(anon_sym_query),
	3117: uint16(anon_sym_mutation),
	3118: uint16(anon_sym_subscription),
	3119: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3120: uint16(anon_sym_fragment),
	3121: uint16(anon_sym_directive),
	3122: uint16(12),
	3123: uint16(400),
	3124: uint16(1),
	3125: uint16(anon_sym_LBRACE),
	3126: uint16(402),
	3127: uint16(1),
	3128: uint16(anon_sym_DOLLAR),
	3129: uint16(404),
	3130: uint16(1),
	3131: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3132: uint16(406),
	3133: uint16(1),
	3134: uint16(anon_sym_DQUOTE),
	3135: uint16(410),
	3136: uint16(1),
	3137: uint16(sym_float_value),
	3138: uint16(414),
	3139: uint16(1),
	3140: uint16(anon_sym_LBRACK),
	3141: uint16(416),
	3142: uint16(1),
	3143: uint16(sym_name),
	3144: uint16(300),
	3145: uint16(1),
	3146: uint16(sym_value),
	3147: uint16(3),
	3148: uint16(2),
	3149: uint16(sym_comment),
	3150: uint16(sym_comma),
	3151: uint16(408),
	3152: uint16(2),
	3153: uint16(sym_int_value),
	3154: uint16(sym_null_value),
	3155: uint16(412),
	3156: uint16(2),
	3157: uint16(anon_sym_true),
	3158: uint16(anon_sym_false),
	3159: uint16(251),
	3160: uint16(6),
	3161: uint16(sym_variable),
	3162: uint16(sym_string_value),
	3163: uint16(sym_boolean_value),
	3164: uint16(sym_enum_value),
	3165: uint16(sym_list_value),
	3166: uint16(sym_object_value),
	3167: uint16(5),
	3168: uint16(81),
	3169: uint16(1),
	3170: uint16(anon_sym_LBRACE),
	3171: uint16(113),
	3172: uint16(1),
	3173: uint16(anon_sym_DQUOTE),
	3174: uint16(135),
	3175: uint16(1),
	3176: uint16(sym_fields_definition),
	3177: uint16(3),
	3178: uint16(2),
	3179: uint16(sym_comment),
	3180: uint16(sym_comma),
	3181: uint16(111),
	3182: uint16(15),
	3184: uint16(anon_sym_schema),
	3185: uint16(anon_sym_extend),
	3186: uint16(anon_sym_scalar),
	3187: uint16(anon_sym_type),
	3188: uint16(anon_sym_interface),
	3189: uint16(anon_sym_union),
	3190: uint16(anon_sym_enum),
	3191: uint16(anon_sym_input),
	3192: uint16(anon_sym_query),
	3193: uint16(anon_sym_mutation),
	3194: uint16(anon_sym_subscription),
	3195: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3196: uint16(anon_sym_fragment),
	3197: uint16(anon_sym_directive),
	3198: uint16(3),
	3199: uint16(420),
	3200: uint16(1),
	3201: uint16(anon_sym_DQUOTE),
	3202: uint16(3),
	3203: uint16(2),
	3204: uint16(sym_comment),
	3205: uint16(sym_comma),
	3206: uint16(418),
	3207: uint16(17),
	3209: uint16(anon_sym_schema),
	3210: uint16(anon_sym_LBRACE),
	3211: uint16(anon_sym_extend),
	3212: uint16(anon_sym_scalar),
	3213: uint16(anon_sym_type),
	3214: uint16(anon_sym_interface),
	3215: uint16(anon_sym_union),
	3216: uint16(anon_sym_enum),
	3217: uint16(anon_sym_input),
	3218: uint16(anon_sym_PIPE),
	3219: uint16(anon_sym_query),
	3220: uint16(anon_sym_mutation),
	3221: uint16(anon_sym_subscription),
	3222: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3223: uint16(anon_sym_fragment),
	3224: uint16(anon_sym_directive),
	3225: uint16(4),
	3226: uint16(123),
	3227: uint16(1),
	3228: uint16(anon_sym_DQUOTE),
	3229: uint16(324),
	3230: uint16(1),
	3231: uint16(anon_sym_PIPE),
	3232: uint16(3),
	3233: uint16(2),
	3234: uint16(sym_comment),
	3235: uint16(sym_comma),
	3236: uint16(119),
	3237: uint16(16),
	3239: uint16(anon_sym_schema),
	3240: uint16(anon_sym_LBRACE),
	3241: uint16(anon_sym_extend),
	3242: uint16(anon_sym_scalar),
	3243: uint16(anon_sym_type),
	3244: uint16(anon_sym_interface),
	3245: uint16(anon_sym_union),
	3246: uint16(anon_sym_enum),
	3247: uint16(anon_sym_input),
	3248: uint16(anon_sym_query),
	3249: uint16(anon_sym_mutation),
	3250: uint16(anon_sym_subscription),
	3251: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3252: uint16(anon_sym_fragment),
	3253: uint16(anon_sym_directive),
	3254: uint16(4),
	3255: uint16(368),
	3256: uint16(1),
	3257: uint16(anon_sym_PIPE),
	3258: uint16(424),
	3259: uint16(1),
	3260: uint16(anon_sym_DQUOTE),
	3261: uint16(3),
	3262: uint16(2),
	3263: uint16(sym_comment),
	3264: uint16(sym_comma),
	3265: uint16(422),
	3266: uint16(16),
	3268: uint16(anon_sym_schema),
	3269: uint16(anon_sym_LBRACE),
	3270: uint16(anon_sym_extend),
	3271: uint16(anon_sym_scalar),
	3272: uint16(anon_sym_type),
	3273: uint16(anon_sym_interface),
	3274: uint16(anon_sym_union),
	3275: uint16(anon_sym_enum),
	3276: uint16(anon_sym_input),
	3277: uint16(anon_sym_query),
	3278: uint16(anon_sym_mutation),
	3279: uint16(anon_sym_subscription),
	3280: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3281: uint16(anon_sym_fragment),
	3282: uint16(anon_sym_directive),
	3283: uint16(12),
	3284: uint16(400),
	3285: uint16(1),
	3286: uint16(anon_sym_LBRACE),
	3287: uint16(402),
	3288: uint16(1),
	3289: uint16(anon_sym_DOLLAR),
	3290: uint16(404),
	3291: uint16(1),
	3292: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3293: uint16(406),
	3294: uint16(1),
	3295: uint16(anon_sym_DQUOTE),
	3296: uint16(410),
	3297: uint16(1),
	3298: uint16(sym_float_value),
	3299: uint16(414),
	3300: uint16(1),
	3301: uint16(anon_sym_LBRACK),
	3302: uint16(416),
	3303: uint16(1),
	3304: uint16(sym_name),
	3305: uint16(318),
	3306: uint16(1),
	3307: uint16(sym_value),
	3308: uint16(3),
	3309: uint16(2),
	3310: uint16(sym_comment),
	3311: uint16(sym_comma),
	3312: uint16(408),
	3313: uint16(2),
	3314: uint16(sym_int_value),
	3315: uint16(sym_null_value),
	3316: uint16(412),
	3317: uint16(2),
	3318: uint16(anon_sym_true),
	3319: uint16(anon_sym_false),
	3320: uint16(251),
	3321: uint16(6),
	3322: uint16(sym_variable),
	3323: uint16(sym_string_value),
	3324: uint16(sym_boolean_value),
	3325: uint16(sym_enum_value),
	3326: uint16(sym_list_value),
	3327: uint16(sym_object_value),
	3328: uint16(4),
	3329: uint16(368),
	3330: uint16(1),
	3331: uint16(anon_sym_PIPE),
	3332: uint16(428),
	3333: uint16(1),
	3334: uint16(anon_sym_DQUOTE),
	3335: uint16(3),
	3336: uint16(2),
	3337: uint16(sym_comment),
	3338: uint16(sym_comma),
	3339: uint16(426),
	3340: uint16(16),
	3342: uint16(anon_sym_schema),
	3343: uint16(anon_sym_LBRACE),
	3344: uint16(anon_sym_extend),
	3345: uint16(anon_sym_scalar),
	3346: uint16(anon_sym_type),
	3347: uint16(anon_sym_interface),
	3348: uint16(anon_sym_union),
	3349: uint16(anon_sym_enum),
	3350: uint16(anon_sym_input),
	3351: uint16(anon_sym_query),
	3352: uint16(anon_sym_mutation),
	3353: uint16(anon_sym_subscription),
	3354: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3355: uint16(anon_sym_fragment),
	3356: uint16(anon_sym_directive),
	3357: uint16(5),
	3358: uint16(163),
	3359: uint16(1),
	3360: uint16(anon_sym_LBRACE),
	3361: uint16(165),
	3362: uint16(1),
	3363: uint16(anon_sym_DQUOTE),
	3364: uint16(137),
	3365: uint16(1),
	3366: uint16(sym_enum_values_definition),
	3367: uint16(3),
	3368: uint16(2),
	3369: uint16(sym_comment),
	3370: uint16(sym_comma),
	3371: uint16(161),
	3372: uint16(15),
	3374: uint16(anon_sym_schema),
	3375: uint16(anon_sym_extend),
	3376: uint16(anon_sym_scalar),
	3377: uint16(anon_sym_type),
	3378: uint16(anon_sym_interface),
	3379: uint16(anon_sym_union),
	3380: uint16(anon_sym_enum),
	3381: uint16(anon_sym_input),
	3382: uint16(anon_sym_query),
	3383: uint16(anon_sym_mutation),
	3384: uint16(anon_sym_subscription),
	3385: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3386: uint16(anon_sym_fragment),
	3387: uint16(anon_sym_directive),
	3388: uint16(3),
	3389: uint16(105),
	3390: uint16(1),
	3391: uint16(anon_sym_DQUOTE),
	3392: uint16(3),
	3393: uint16(2),
	3394: uint16(sym_comment),
	3395: uint16(sym_comma),
	3396: uint16(103),
	3397: uint16(16),
	3399: uint16(anon_sym_schema),
	3400: uint16(anon_sym_LBRACE),
	3401: uint16(anon_sym_extend),
	3402: uint16(anon_sym_scalar),
	3403: uint16(anon_sym_type),
	3404: uint16(anon_sym_interface),
	3405: uint16(anon_sym_union),
	3406: uint16(anon_sym_enum),
	3407: uint16(anon_sym_input),
	3408: uint16(anon_sym_query),
	3409: uint16(anon_sym_mutation),
	3410: uint16(anon_sym_subscription),
	3411: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3412: uint16(anon_sym_fragment),
	3413: uint16(anon_sym_directive),
	3414: uint16(3),
	3415: uint16(432),
	3416: uint16(1),
	3417: uint16(anon_sym_DQUOTE),
	3418: uint16(3),
	3419: uint16(2),
	3420: uint16(sym_comment),
	3421: uint16(sym_comma),
	3422: uint16(430),
	3423: uint16(16),
	3425: uint16(anon_sym_schema),
	3426: uint16(anon_sym_LBRACE),
	3427: uint16(anon_sym_extend),
	3428: uint16(anon_sym_scalar),
	3429: uint16(anon_sym_type),
	3430: uint16(anon_sym_interface),
	3431: uint16(anon_sym_union),
	3432: uint16(anon_sym_enum),
	3433: uint16(anon_sym_input),
	3434: uint16(anon_sym_query),
	3435: uint16(anon_sym_mutation),
	3436: uint16(anon_sym_subscription),
	3437: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3438: uint16(anon_sym_fragment),
	3439: uint16(anon_sym_directive),
	3440: uint16(3),
	3441: uint16(436),
	3442: uint16(1),
	3443: uint16(anon_sym_DQUOTE),
	3444: uint16(3),
	3445: uint16(2),
	3446: uint16(sym_comment),
	3447: uint16(sym_comma),
	3448: uint16(434),
	3449: uint16(16),
	3451: uint16(anon_sym_schema),
	3452: uint16(anon_sym_LBRACE),
	3453: uint16(anon_sym_extend),
	3454: uint16(anon_sym_scalar),
	3455: uint16(anon_sym_type),
	3456: uint16(anon_sym_interface),
	3457: uint16(anon_sym_union),
	3458: uint16(anon_sym_enum),
	3459: uint16(anon_sym_input),
	3460: uint16(anon_sym_query),
	3461: uint16(anon_sym_mutation),
	3462: uint16(anon_sym_subscription),
	3463: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3464: uint16(anon_sym_fragment),
	3465: uint16(anon_sym_directive),
	3466: uint16(3),
	3467: uint16(440),
	3468: uint16(1),
	3469: uint16(anon_sym_DQUOTE),
	3470: uint16(3),
	3471: uint16(2),
	3472: uint16(sym_comment),
	3473: uint16(sym_comma),
	3474: uint16(438),
	3475: uint16(16),
	3477: uint16(anon_sym_schema),
	3478: uint16(anon_sym_LBRACE),
	3479: uint16(anon_sym_extend),
	3480: uint16(anon_sym_scalar),
	3481: uint16(anon_sym_type),
	3482: uint16(anon_sym_interface),
	3483: uint16(anon_sym_union),
	3484: uint16(anon_sym_enum),
	3485: uint16(anon_sym_input),
	3486: uint16(anon_sym_query),
	3487: uint16(anon_sym_mutation),
	3488: uint16(anon_sym_subscription),
	3489: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3490: uint16(anon_sym_fragment),
	3491: uint16(anon_sym_directive),
	3492: uint16(3),
	3493: uint16(444),
	3494: uint16(1),
	3495: uint16(anon_sym_DQUOTE),
	3496: uint16(3),
	3497: uint16(2),
	3498: uint16(sym_comment),
	3499: uint16(sym_comma),
	3500: uint16(442),
	3501: uint16(16),
	3503: uint16(anon_sym_schema),
	3504: uint16(anon_sym_LBRACE),
	3505: uint16(anon_sym_extend),
	3506: uint16(anon_sym_scalar),
	3507: uint16(anon_sym_type),
	3508: uint16(anon_sym_interface),
	3509: uint16(anon_sym_union),
	3510: uint16(anon_sym_enum),
	3511: uint16(anon_sym_input),
	3512: uint16(anon_sym_query),
	3513: uint16(anon_sym_mutation),
	3514: uint16(anon_sym_subscription),
	3515: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3516: uint16(anon_sym_fragment),
	3517: uint16(anon_sym_directive),
	3518: uint16(3),
	3519: uint16(448),
	3520: uint16(1),
	3521: uint16(anon_sym_DQUOTE),
	3522: uint16(3),
	3523: uint16(2),
	3524: uint16(sym_comment),
	3525: uint16(sym_comma),
	3526: uint16(446),
	3527: uint16(16),
	3529: uint16(anon_sym_schema),
	3530: uint16(anon_sym_LBRACE),
	3531: uint16(anon_sym_extend),
	3532: uint16(anon_sym_scalar),
	3533: uint16(anon_sym_type),
	3534: uint16(anon_sym_interface),
	3535: uint16(anon_sym_union),
	3536: uint16(anon_sym_enum),
	3537: uint16(anon_sym_input),
	3538: uint16(anon_sym_query),
	3539: uint16(anon_sym_mutation),
	3540: uint16(anon_sym_subscription),
	3541: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3542: uint16(anon_sym_fragment),
	3543: uint16(anon_sym_directive),
	3544: uint16(3),
	3545: uint16(452),
	3546: uint16(1),
	3547: uint16(anon_sym_DQUOTE),
	3548: uint16(3),
	3549: uint16(2),
	3550: uint16(sym_comment),
	3551: uint16(sym_comma),
	3552: uint16(450),
	3553: uint16(16),
	3555: uint16(anon_sym_schema),
	3556: uint16(anon_sym_LBRACE),
	3557: uint16(anon_sym_extend),
	3558: uint16(anon_sym_scalar),
	3559: uint16(anon_sym_type),
	3560: uint16(anon_sym_interface),
	3561: uint16(anon_sym_union),
	3562: uint16(anon_sym_enum),
	3563: uint16(anon_sym_input),
	3564: uint16(anon_sym_query),
	3565: uint16(anon_sym_mutation),
	3566: uint16(anon_sym_subscription),
	3567: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3568: uint16(anon_sym_fragment),
	3569: uint16(anon_sym_directive),
	3570: uint16(3),
	3571: uint16(456),
	3572: uint16(1),
	3573: uint16(anon_sym_DQUOTE),
	3574: uint16(3),
	3575: uint16(2),
	3576: uint16(sym_comment),
	3577: uint16(sym_comma),
	3578: uint16(454),
	3579: uint16(16),
	3581: uint16(anon_sym_schema),
	3582: uint16(anon_sym_LBRACE),
	3583: uint16(anon_sym_extend),
	3584: uint16(anon_sym_scalar),
	3585: uint16(anon_sym_type),
	3586: uint16(anon_sym_interface),
	3587: uint16(anon_sym_union),
	3588: uint16(anon_sym_enum),
	3589: uint16(anon_sym_input),
	3590: uint16(anon_sym_query),
	3591: uint16(anon_sym_mutation),
	3592: uint16(anon_sym_subscription),
	3593: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3594: uint16(anon_sym_fragment),
	3595: uint16(anon_sym_directive),
	3596: uint16(3),
	3597: uint16(460),
	3598: uint16(1),
	3599: uint16(anon_sym_DQUOTE),
	3600: uint16(3),
	3601: uint16(2),
	3602: uint16(sym_comment),
	3603: uint16(sym_comma),
	3604: uint16(458),
	3605: uint16(16),
	3607: uint16(anon_sym_schema),
	3608: uint16(anon_sym_LBRACE),
	3609: uint16(anon_sym_extend),
	3610: uint16(anon_sym_scalar),
	3611: uint16(anon_sym_type),
	3612: uint16(anon_sym_interface),
	3613: uint16(anon_sym_union),
	3614: uint16(anon_sym_enum),
	3615: uint16(anon_sym_input),
	3616: uint16(anon_sym_query),
	3617: uint16(anon_sym_mutation),
	3618: uint16(anon_sym_subscription),
	3619: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3620: uint16(anon_sym_fragment),
	3621: uint16(anon_sym_directive),
	3622: uint16(3),
	3623: uint16(464),
	3624: uint16(1),
	3625: uint16(anon_sym_DQUOTE),
	3626: uint16(3),
	3627: uint16(2),
	3628: uint16(sym_comment),
	3629: uint16(sym_comma),
	3630: uint16(462),
	3631: uint16(16),
	3633: uint16(anon_sym_schema),
	3634: uint16(anon_sym_LBRACE),
	3635: uint16(anon_sym_extend),
	3636: uint16(anon_sym_scalar),
	3637: uint16(anon_sym_type),
	3638: uint16(anon_sym_interface),
	3639: uint16(anon_sym_union),
	3640: uint16(anon_sym_enum),
	3641: uint16(anon_sym_input),
	3642: uint16(anon_sym_query),
	3643: uint16(anon_sym_mutation),
	3644: uint16(anon_sym_subscription),
	3645: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3646: uint16(anon_sym_fragment),
	3647: uint16(anon_sym_directive),
	3648: uint16(3),
	3649: uint16(386),
	3650: uint16(1),
	3651: uint16(anon_sym_DQUOTE),
	3652: uint16(3),
	3653: uint16(2),
	3654: uint16(sym_comment),
	3655: uint16(sym_comma),
	3656: uint16(384),
	3657: uint16(16),
	3659: uint16(anon_sym_schema),
	3660: uint16(anon_sym_LBRACE),
	3661: uint16(anon_sym_extend),
	3662: uint16(anon_sym_scalar),
	3663: uint16(anon_sym_type),
	3664: uint16(anon_sym_interface),
	3665: uint16(anon_sym_union),
	3666: uint16(anon_sym_enum),
	3667: uint16(anon_sym_input),
	3668: uint16(anon_sym_query),
	3669: uint16(anon_sym_mutation),
	3670: uint16(anon_sym_subscription),
	3671: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3672: uint16(anon_sym_fragment),
	3673: uint16(anon_sym_directive),
	3674: uint16(3),
	3675: uint16(382),
	3676: uint16(1),
	3677: uint16(anon_sym_DQUOTE),
	3678: uint16(3),
	3679: uint16(2),
	3680: uint16(sym_comment),
	3681: uint16(sym_comma),
	3682: uint16(380),
	3683: uint16(16),
	3685: uint16(anon_sym_schema),
	3686: uint16(anon_sym_LBRACE),
	3687: uint16(anon_sym_extend),
	3688: uint16(anon_sym_scalar),
	3689: uint16(anon_sym_type),
	3690: uint16(anon_sym_interface),
	3691: uint16(anon_sym_union),
	3692: uint16(anon_sym_enum),
	3693: uint16(anon_sym_input),
	3694: uint16(anon_sym_query),
	3695: uint16(anon_sym_mutation),
	3696: uint16(anon_sym_subscription),
	3697: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3698: uint16(anon_sym_fragment),
	3699: uint16(anon_sym_directive),
	3700: uint16(3),
	3701: uint16(468),
	3702: uint16(1),
	3703: uint16(anon_sym_DQUOTE),
	3704: uint16(3),
	3705: uint16(2),
	3706: uint16(sym_comment),
	3707: uint16(sym_comma),
	3708: uint16(466),
	3709: uint16(16),
	3711: uint16(anon_sym_schema),
	3712: uint16(anon_sym_LBRACE),
	3713: uint16(anon_sym_extend),
	3714: uint16(anon_sym_scalar),
	3715: uint16(anon_sym_type),
	3716: uint16(anon_sym_interface),
	3717: uint16(anon_sym_union),
	3718: uint16(anon_sym_enum),
	3719: uint16(anon_sym_input),
	3720: uint16(anon_sym_query),
	3721: uint16(anon_sym_mutation),
	3722: uint16(anon_sym_subscription),
	3723: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3724: uint16(anon_sym_fragment),
	3725: uint16(anon_sym_directive),
	3726: uint16(3),
	3727: uint16(472),
	3728: uint16(1),
	3729: uint16(anon_sym_DQUOTE),
	3730: uint16(3),
	3731: uint16(2),
	3732: uint16(sym_comment),
	3733: uint16(sym_comma),
	3734: uint16(470),
	3735: uint16(16),
	3737: uint16(anon_sym_schema),
	3738: uint16(anon_sym_LBRACE),
	3739: uint16(anon_sym_extend),
	3740: uint16(anon_sym_scalar),
	3741: uint16(anon_sym_type),
	3742: uint16(anon_sym_interface),
	3743: uint16(anon_sym_union),
	3744: uint16(anon_sym_enum),
	3745: uint16(anon_sym_input),
	3746: uint16(anon_sym_query),
	3747: uint16(anon_sym_mutation),
	3748: uint16(anon_sym_subscription),
	3749: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3750: uint16(anon_sym_fragment),
	3751: uint16(anon_sym_directive),
	3752: uint16(3),
	3753: uint16(476),
	3754: uint16(1),
	3755: uint16(anon_sym_DQUOTE),
	3756: uint16(3),
	3757: uint16(2),
	3758: uint16(sym_comment),
	3759: uint16(sym_comma),
	3760: uint16(474),
	3761: uint16(16),
	3763: uint16(anon_sym_schema),
	3764: uint16(anon_sym_LBRACE),
	3765: uint16(anon_sym_extend),
	3766: uint16(anon_sym_scalar),
	3767: uint16(anon_sym_type),
	3768: uint16(anon_sym_interface),
	3769: uint16(anon_sym_union),
	3770: uint16(anon_sym_enum),
	3771: uint16(anon_sym_input),
	3772: uint16(anon_sym_query),
	3773: uint16(anon_sym_mutation),
	3774: uint16(anon_sym_subscription),
	3775: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3776: uint16(anon_sym_fragment),
	3777: uint16(anon_sym_directive),
	3778: uint16(3),
	3779: uint16(137),
	3780: uint16(1),
	3781: uint16(anon_sym_DQUOTE),
	3782: uint16(3),
	3783: uint16(2),
	3784: uint16(sym_comment),
	3785: uint16(sym_comma),
	3786: uint16(135),
	3787: uint16(16),
	3789: uint16(anon_sym_schema),
	3790: uint16(anon_sym_LBRACE),
	3791: uint16(anon_sym_extend),
	3792: uint16(anon_sym_scalar),
	3793: uint16(anon_sym_type),
	3794: uint16(anon_sym_interface),
	3795: uint16(anon_sym_union),
	3796: uint16(anon_sym_enum),
	3797: uint16(anon_sym_input),
	3798: uint16(anon_sym_query),
	3799: uint16(anon_sym_mutation),
	3800: uint16(anon_sym_subscription),
	3801: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3802: uint16(anon_sym_fragment),
	3803: uint16(anon_sym_directive),
	3804: uint16(3),
	3805: uint16(480),
	3806: uint16(1),
	3807: uint16(anon_sym_DQUOTE),
	3808: uint16(3),
	3809: uint16(2),
	3810: uint16(sym_comment),
	3811: uint16(sym_comma),
	3812: uint16(478),
	3813: uint16(16),
	3815: uint16(anon_sym_schema),
	3816: uint16(anon_sym_LBRACE),
	3817: uint16(anon_sym_extend),
	3818: uint16(anon_sym_scalar),
	3819: uint16(anon_sym_type),
	3820: uint16(anon_sym_interface),
	3821: uint16(anon_sym_union),
	3822: uint16(anon_sym_enum),
	3823: uint16(anon_sym_input),
	3824: uint16(anon_sym_query),
	3825: uint16(anon_sym_mutation),
	3826: uint16(anon_sym_subscription),
	3827: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3828: uint16(anon_sym_fragment),
	3829: uint16(anon_sym_directive),
	3830: uint16(3),
	3831: uint16(191),
	3832: uint16(1),
	3833: uint16(anon_sym_DQUOTE),
	3834: uint16(3),
	3835: uint16(2),
	3836: uint16(sym_comment),
	3837: uint16(sym_comma),
	3838: uint16(189),
	3839: uint16(16),
	3841: uint16(anon_sym_schema),
	3842: uint16(anon_sym_LBRACE),
	3843: uint16(anon_sym_extend),
	3844: uint16(anon_sym_scalar),
	3845: uint16(anon_sym_type),
	3846: uint16(anon_sym_interface),
	3847: uint16(anon_sym_union),
	3848: uint16(anon_sym_enum),
	3849: uint16(anon_sym_input),
	3850: uint16(anon_sym_query),
	3851: uint16(anon_sym_mutation),
	3852: uint16(anon_sym_subscription),
	3853: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3854: uint16(anon_sym_fragment),
	3855: uint16(anon_sym_directive),
	3856: uint16(3),
	3857: uint16(348),
	3858: uint16(1),
	3859: uint16(anon_sym_DQUOTE),
	3860: uint16(3),
	3861: uint16(2),
	3862: uint16(sym_comment),
	3863: uint16(sym_comma),
	3864: uint16(346),
	3865: uint16(16),
	3867: uint16(anon_sym_schema),
	3868: uint16(anon_sym_LBRACE),
	3869: uint16(anon_sym_extend),
	3870: uint16(anon_sym_scalar),
	3871: uint16(anon_sym_type),
	3872: uint16(anon_sym_interface),
	3873: uint16(anon_sym_union),
	3874: uint16(anon_sym_enum),
	3875: uint16(anon_sym_input),
	3876: uint16(anon_sym_query),
	3877: uint16(anon_sym_mutation),
	3878: uint16(anon_sym_subscription),
	3879: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3880: uint16(anon_sym_fragment),
	3881: uint16(anon_sym_directive),
	3882: uint16(3),
	3883: uint16(484),
	3884: uint16(1),
	3885: uint16(anon_sym_DQUOTE),
	3886: uint16(3),
	3887: uint16(2),
	3888: uint16(sym_comment),
	3889: uint16(sym_comma),
	3890: uint16(482),
	3891: uint16(16),
	3893: uint16(anon_sym_schema),
	3894: uint16(anon_sym_LBRACE),
	3895: uint16(anon_sym_extend),
	3896: uint16(anon_sym_scalar),
	3897: uint16(anon_sym_type),
	3898: uint16(anon_sym_interface),
	3899: uint16(anon_sym_union),
	3900: uint16(anon_sym_enum),
	3901: uint16(anon_sym_input),
	3902: uint16(anon_sym_query),
	3903: uint16(anon_sym_mutation),
	3904: uint16(anon_sym_subscription),
	3905: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3906: uint16(anon_sym_fragment),
	3907: uint16(anon_sym_directive),
	3908: uint16(3),
	3909: uint16(113),
	3910: uint16(1),
	3911: uint16(anon_sym_DQUOTE),
	3912: uint16(3),
	3913: uint16(2),
	3914: uint16(sym_comment),
	3915: uint16(sym_comma),
	3916: uint16(111),
	3917: uint16(16),
	3919: uint16(anon_sym_schema),
	3920: uint16(anon_sym_LBRACE),
	3921: uint16(anon_sym_extend),
	3922: uint16(anon_sym_scalar),
	3923: uint16(anon_sym_type),
	3924: uint16(anon_sym_interface),
	3925: uint16(anon_sym_union),
	3926: uint16(anon_sym_enum),
	3927: uint16(anon_sym_input),
	3928: uint16(anon_sym_query),
	3929: uint16(anon_sym_mutation),
	3930: uint16(anon_sym_subscription),
	3931: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3932: uint16(anon_sym_fragment),
	3933: uint16(anon_sym_directive),
	3934: uint16(3),
	3935: uint16(344),
	3936: uint16(1),
	3937: uint16(anon_sym_DQUOTE),
	3938: uint16(3),
	3939: uint16(2),
	3940: uint16(sym_comment),
	3941: uint16(sym_comma),
	3942: uint16(342),
	3943: uint16(16),
	3945: uint16(anon_sym_schema),
	3946: uint16(anon_sym_LBRACE),
	3947: uint16(anon_sym_extend),
	3948: uint16(anon_sym_scalar),
	3949: uint16(anon_sym_type),
	3950: uint16(anon_sym_interface),
	3951: uint16(anon_sym_union),
	3952: uint16(anon_sym_enum),
	3953: uint16(anon_sym_input),
	3954: uint16(anon_sym_query),
	3955: uint16(anon_sym_mutation),
	3956: uint16(anon_sym_subscription),
	3957: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3958: uint16(anon_sym_fragment),
	3959: uint16(anon_sym_directive),
	3960: uint16(3),
	3961: uint16(488),
	3962: uint16(1),
	3963: uint16(anon_sym_DQUOTE),
	3964: uint16(3),
	3965: uint16(2),
	3966: uint16(sym_comment),
	3967: uint16(sym_comma),
	3968: uint16(486),
	3969: uint16(16),
	3971: uint16(anon_sym_schema),
	3972: uint16(anon_sym_LBRACE),
	3973: uint16(anon_sym_extend),
	3974: uint16(anon_sym_scalar),
	3975: uint16(anon_sym_type),
	3976: uint16(anon_sym_interface),
	3977: uint16(anon_sym_union),
	3978: uint16(anon_sym_enum),
	3979: uint16(anon_sym_input),
	3980: uint16(anon_sym_query),
	3981: uint16(anon_sym_mutation),
	3982: uint16(anon_sym_subscription),
	3983: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	3984: uint16(anon_sym_fragment),
	3985: uint16(anon_sym_directive),
	3986: uint16(3),
	3987: uint16(165),
	3988: uint16(1),
	3989: uint16(anon_sym_DQUOTE),
	3990: uint16(3),
	3991: uint16(2),
	3992: uint16(sym_comment),
	3993: uint16(sym_comma),
	3994: uint16(161),
	3995: uint16(16),
	3997: uint16(anon_sym_schema),
	3998: uint16(anon_sym_LBRACE),
	3999: uint16(anon_sym_extend),
	4000: uint16(anon_sym_scalar),
	4001: uint16(anon_sym_type),
	4002: uint16(anon_sym_interface),
	4003: uint16(anon_sym_union),
	4004: uint16(anon_sym_enum),
	4005: uint16(anon_sym_input),
	4006: uint16(anon_sym_query),
	4007: uint16(anon_sym_mutation),
	4008: uint16(anon_sym_subscription),
	4009: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4010: uint16(anon_sym_fragment),
	4011: uint16(anon_sym_directive),
	4012: uint16(3),
	4013: uint16(492),
	4014: uint16(1),
	4015: uint16(anon_sym_DQUOTE),
	4016: uint16(3),
	4017: uint16(2),
	4018: uint16(sym_comment),
	4019: uint16(sym_comma),
	4020: uint16(490),
	4021: uint16(16),
	4023: uint16(anon_sym_schema),
	4024: uint16(anon_sym_LBRACE),
	4025: uint16(anon_sym_extend),
	4026: uint16(anon_sym_scalar),
	4027: uint16(anon_sym_type),
	4028: uint16(anon_sym_interface),
	4029: uint16(anon_sym_union),
	4030: uint16(anon_sym_enum),
	4031: uint16(anon_sym_input),
	4032: uint16(anon_sym_query),
	4033: uint16(anon_sym_mutation),
	4034: uint16(anon_sym_subscription),
	4035: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4036: uint16(anon_sym_fragment),
	4037: uint16(anon_sym_directive),
	4038: uint16(3),
	4039: uint16(169),
	4040: uint16(1),
	4041: uint16(anon_sym_DQUOTE),
	4042: uint16(3),
	4043: uint16(2),
	4044: uint16(sym_comment),
	4045: uint16(sym_comma),
	4046: uint16(167),
	4047: uint16(16),
	4049: uint16(anon_sym_schema),
	4050: uint16(anon_sym_LBRACE),
	4051: uint16(anon_sym_extend),
	4052: uint16(anon_sym_scalar),
	4053: uint16(anon_sym_type),
	4054: uint16(anon_sym_interface),
	4055: uint16(anon_sym_union),
	4056: uint16(anon_sym_enum),
	4057: uint16(anon_sym_input),
	4058: uint16(anon_sym_query),
	4059: uint16(anon_sym_mutation),
	4060: uint16(anon_sym_subscription),
	4061: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4062: uint16(anon_sym_fragment),
	4063: uint16(anon_sym_directive),
	4064: uint16(3),
	4065: uint16(496),
	4066: uint16(1),
	4067: uint16(anon_sym_DQUOTE),
	4068: uint16(3),
	4069: uint16(2),
	4070: uint16(sym_comment),
	4071: uint16(sym_comma),
	4072: uint16(494),
	4073: uint16(16),
	4075: uint16(anon_sym_schema),
	4076: uint16(anon_sym_LBRACE),
	4077: uint16(anon_sym_extend),
	4078: uint16(anon_sym_scalar),
	4079: uint16(anon_sym_type),
	4080: uint16(anon_sym_interface),
	4081: uint16(anon_sym_union),
	4082: uint16(anon_sym_enum),
	4083: uint16(anon_sym_input),
	4084: uint16(anon_sym_query),
	4085: uint16(anon_sym_mutation),
	4086: uint16(anon_sym_subscription),
	4087: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4088: uint16(anon_sym_fragment),
	4089: uint16(anon_sym_directive),
	4090: uint16(3),
	4091: uint16(500),
	4092: uint16(1),
	4093: uint16(anon_sym_DQUOTE),
	4094: uint16(3),
	4095: uint16(2),
	4096: uint16(sym_comment),
	4097: uint16(sym_comma),
	4098: uint16(498),
	4099: uint16(16),
	4101: uint16(anon_sym_schema),
	4102: uint16(anon_sym_LBRACE),
	4103: uint16(anon_sym_extend),
	4104: uint16(anon_sym_scalar),
	4105: uint16(anon_sym_type),
	4106: uint16(anon_sym_interface),
	4107: uint16(anon_sym_union),
	4108: uint16(anon_sym_enum),
	4109: uint16(anon_sym_input),
	4110: uint16(anon_sym_query),
	4111: uint16(anon_sym_mutation),
	4112: uint16(anon_sym_subscription),
	4113: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4114: uint16(anon_sym_fragment),
	4115: uint16(anon_sym_directive),
	4116: uint16(3),
	4117: uint16(504),
	4118: uint16(1),
	4119: uint16(anon_sym_DQUOTE),
	4120: uint16(3),
	4121: uint16(2),
	4122: uint16(sym_comment),
	4123: uint16(sym_comma),
	4124: uint16(502),
	4125: uint16(16),
	4127: uint16(anon_sym_schema),
	4128: uint16(anon_sym_LBRACE),
	4129: uint16(anon_sym_extend),
	4130: uint16(anon_sym_scalar),
	4131: uint16(anon_sym_type),
	4132: uint16(anon_sym_interface),
	4133: uint16(anon_sym_union),
	4134: uint16(anon_sym_enum),
	4135: uint16(anon_sym_input),
	4136: uint16(anon_sym_query),
	4137: uint16(anon_sym_mutation),
	4138: uint16(anon_sym_subscription),
	4139: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4140: uint16(anon_sym_fragment),
	4141: uint16(anon_sym_directive),
	4142: uint16(3),
	4143: uint16(508),
	4144: uint16(1),
	4145: uint16(anon_sym_DQUOTE),
	4146: uint16(3),
	4147: uint16(2),
	4148: uint16(sym_comment),
	4149: uint16(sym_comma),
	4150: uint16(506),
	4151: uint16(16),
	4153: uint16(anon_sym_schema),
	4154: uint16(anon_sym_LBRACE),
	4155: uint16(anon_sym_extend),
	4156: uint16(anon_sym_scalar),
	4157: uint16(anon_sym_type),
	4158: uint16(anon_sym_interface),
	4159: uint16(anon_sym_union),
	4160: uint16(anon_sym_enum),
	4161: uint16(anon_sym_input),
	4162: uint16(anon_sym_query),
	4163: uint16(anon_sym_mutation),
	4164: uint16(anon_sym_subscription),
	4165: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4166: uint16(anon_sym_fragment),
	4167: uint16(anon_sym_directive),
	4168: uint16(3),
	4169: uint16(512),
	4170: uint16(1),
	4171: uint16(anon_sym_DQUOTE),
	4172: uint16(3),
	4173: uint16(2),
	4174: uint16(sym_comment),
	4175: uint16(sym_comma),
	4176: uint16(510),
	4177: uint16(16),
	4179: uint16(anon_sym_schema),
	4180: uint16(anon_sym_LBRACE),
	4181: uint16(anon_sym_extend),
	4182: uint16(anon_sym_scalar),
	4183: uint16(anon_sym_type),
	4184: uint16(anon_sym_interface),
	4185: uint16(anon_sym_union),
	4186: uint16(anon_sym_enum),
	4187: uint16(anon_sym_input),
	4188: uint16(anon_sym_query),
	4189: uint16(anon_sym_mutation),
	4190: uint16(anon_sym_subscription),
	4191: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4192: uint16(anon_sym_fragment),
	4193: uint16(anon_sym_directive),
	4194: uint16(3),
	4195: uint16(516),
	4196: uint16(1),
	4197: uint16(anon_sym_DQUOTE),
	4198: uint16(3),
	4199: uint16(2),
	4200: uint16(sym_comment),
	4201: uint16(sym_comma),
	4202: uint16(514),
	4203: uint16(16),
	4205: uint16(anon_sym_schema),
	4206: uint16(anon_sym_LBRACE),
	4207: uint16(anon_sym_extend),
	4208: uint16(anon_sym_scalar),
	4209: uint16(anon_sym_type),
	4210: uint16(anon_sym_interface),
	4211: uint16(anon_sym_union),
	4212: uint16(anon_sym_enum),
	4213: uint16(anon_sym_input),
	4214: uint16(anon_sym_query),
	4215: uint16(anon_sym_mutation),
	4216: uint16(anon_sym_subscription),
	4217: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4218: uint16(anon_sym_fragment),
	4219: uint16(anon_sym_directive),
	4220: uint16(3),
	4221: uint16(520),
	4222: uint16(1),
	4223: uint16(anon_sym_DQUOTE),
	4224: uint16(3),
	4225: uint16(2),
	4226: uint16(sym_comment),
	4227: uint16(sym_comma),
	4228: uint16(518),
	4229: uint16(16),
	4231: uint16(anon_sym_schema),
	4232: uint16(anon_sym_LBRACE),
	4233: uint16(anon_sym_extend),
	4234: uint16(anon_sym_scalar),
	4235: uint16(anon_sym_type),
	4236: uint16(anon_sym_interface),
	4237: uint16(anon_sym_union),
	4238: uint16(anon_sym_enum),
	4239: uint16(anon_sym_input),
	4240: uint16(anon_sym_query),
	4241: uint16(anon_sym_mutation),
	4242: uint16(anon_sym_subscription),
	4243: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4244: uint16(anon_sym_fragment),
	4245: uint16(anon_sym_directive),
	4246: uint16(3),
	4247: uint16(524),
	4248: uint16(1),
	4249: uint16(anon_sym_DQUOTE),
	4250: uint16(3),
	4251: uint16(2),
	4252: uint16(sym_comment),
	4253: uint16(sym_comma),
	4254: uint16(522),
	4255: uint16(16),
	4257: uint16(anon_sym_schema),
	4258: uint16(anon_sym_LBRACE),
	4259: uint16(anon_sym_extend),
	4260: uint16(anon_sym_scalar),
	4261: uint16(anon_sym_type),
	4262: uint16(anon_sym_interface),
	4263: uint16(anon_sym_union),
	4264: uint16(anon_sym_enum),
	4265: uint16(anon_sym_input),
	4266: uint16(anon_sym_query),
	4267: uint16(anon_sym_mutation),
	4268: uint16(anon_sym_subscription),
	4269: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4270: uint16(anon_sym_fragment),
	4271: uint16(anon_sym_directive),
	4272: uint16(3),
	4273: uint16(528),
	4274: uint16(1),
	4275: uint16(anon_sym_DQUOTE),
	4276: uint16(3),
	4277: uint16(2),
	4278: uint16(sym_comment),
	4279: uint16(sym_comma),
	4280: uint16(526),
	4281: uint16(16),
	4283: uint16(anon_sym_schema),
	4284: uint16(anon_sym_LBRACE),
	4285: uint16(anon_sym_extend),
	4286: uint16(anon_sym_scalar),
	4287: uint16(anon_sym_type),
	4288: uint16(anon_sym_interface),
	4289: uint16(anon_sym_union),
	4290: uint16(anon_sym_enum),
	4291: uint16(anon_sym_input),
	4292: uint16(anon_sym_query),
	4293: uint16(anon_sym_mutation),
	4294: uint16(anon_sym_subscription),
	4295: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4296: uint16(anon_sym_fragment),
	4297: uint16(anon_sym_directive),
	4298: uint16(3),
	4299: uint16(340),
	4300: uint16(1),
	4301: uint16(anon_sym_DQUOTE),
	4302: uint16(3),
	4303: uint16(2),
	4304: uint16(sym_comment),
	4305: uint16(sym_comma),
	4306: uint16(338),
	4307: uint16(16),
	4309: uint16(anon_sym_schema),
	4310: uint16(anon_sym_LBRACE),
	4311: uint16(anon_sym_extend),
	4312: uint16(anon_sym_scalar),
	4313: uint16(anon_sym_type),
	4314: uint16(anon_sym_interface),
	4315: uint16(anon_sym_union),
	4316: uint16(anon_sym_enum),
	4317: uint16(anon_sym_input),
	4318: uint16(anon_sym_query),
	4319: uint16(anon_sym_mutation),
	4320: uint16(anon_sym_subscription),
	4321: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4322: uint16(anon_sym_fragment),
	4323: uint16(anon_sym_directive),
	4324: uint16(3),
	4325: uint16(532),
	4326: uint16(1),
	4327: uint16(anon_sym_DQUOTE),
	4328: uint16(3),
	4329: uint16(2),
	4330: uint16(sym_comment),
	4331: uint16(sym_comma),
	4332: uint16(530),
	4333: uint16(16),
	4335: uint16(anon_sym_schema),
	4336: uint16(anon_sym_LBRACE),
	4337: uint16(anon_sym_extend),
	4338: uint16(anon_sym_scalar),
	4339: uint16(anon_sym_type),
	4340: uint16(anon_sym_interface),
	4341: uint16(anon_sym_union),
	4342: uint16(anon_sym_enum),
	4343: uint16(anon_sym_input),
	4344: uint16(anon_sym_query),
	4345: uint16(anon_sym_mutation),
	4346: uint16(anon_sym_subscription),
	4347: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4348: uint16(anon_sym_fragment),
	4349: uint16(anon_sym_directive),
	4350: uint16(3),
	4351: uint16(133),
	4352: uint16(1),
	4353: uint16(anon_sym_DQUOTE),
	4354: uint16(3),
	4355: uint16(2),
	4356: uint16(sym_comment),
	4357: uint16(sym_comma),
	4358: uint16(129),
	4359: uint16(16),
	4361: uint16(anon_sym_schema),
	4362: uint16(anon_sym_LBRACE),
	4363: uint16(anon_sym_extend),
	4364: uint16(anon_sym_scalar),
	4365: uint16(anon_sym_type),
	4366: uint16(anon_sym_interface),
	4367: uint16(anon_sym_union),
	4368: uint16(anon_sym_enum),
	4369: uint16(anon_sym_input),
	4370: uint16(anon_sym_query),
	4371: uint16(anon_sym_mutation),
	4372: uint16(anon_sym_subscription),
	4373: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4374: uint16(anon_sym_fragment),
	4375: uint16(anon_sym_directive),
	4376: uint16(3),
	4377: uint16(536),
	4378: uint16(1),
	4379: uint16(anon_sym_DQUOTE),
	4380: uint16(3),
	4381: uint16(2),
	4382: uint16(sym_comment),
	4383: uint16(sym_comma),
	4384: uint16(534),
	4385: uint16(16),
	4387: uint16(anon_sym_schema),
	4388: uint16(anon_sym_LBRACE),
	4389: uint16(anon_sym_extend),
	4390: uint16(anon_sym_scalar),
	4391: uint16(anon_sym_type),
	4392: uint16(anon_sym_interface),
	4393: uint16(anon_sym_union),
	4394: uint16(anon_sym_enum),
	4395: uint16(anon_sym_input),
	4396: uint16(anon_sym_query),
	4397: uint16(anon_sym_mutation),
	4398: uint16(anon_sym_subscription),
	4399: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4400: uint16(anon_sym_fragment),
	4401: uint16(anon_sym_directive),
	4402: uint16(3),
	4403: uint16(540),
	4404: uint16(1),
	4405: uint16(anon_sym_DQUOTE),
	4406: uint16(3),
	4407: uint16(2),
	4408: uint16(sym_comment),
	4409: uint16(sym_comma),
	4410: uint16(538),
	4411: uint16(16),
	4413: uint16(anon_sym_schema),
	4414: uint16(anon_sym_LBRACE),
	4415: uint16(anon_sym_extend),
	4416: uint16(anon_sym_scalar),
	4417: uint16(anon_sym_type),
	4418: uint16(anon_sym_interface),
	4419: uint16(anon_sym_union),
	4420: uint16(anon_sym_enum),
	4421: uint16(anon_sym_input),
	4422: uint16(anon_sym_query),
	4423: uint16(anon_sym_mutation),
	4424: uint16(anon_sym_subscription),
	4425: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4426: uint16(anon_sym_fragment),
	4427: uint16(anon_sym_directive),
	4428: uint16(3),
	4429: uint16(544),
	4430: uint16(1),
	4431: uint16(anon_sym_DQUOTE),
	4432: uint16(3),
	4433: uint16(2),
	4434: uint16(sym_comment),
	4435: uint16(sym_comma),
	4436: uint16(542),
	4437: uint16(16),
	4439: uint16(anon_sym_schema),
	4440: uint16(anon_sym_LBRACE),
	4441: uint16(anon_sym_extend),
	4442: uint16(anon_sym_scalar),
	4443: uint16(anon_sym_type),
	4444: uint16(anon_sym_interface),
	4445: uint16(anon_sym_union),
	4446: uint16(anon_sym_enum),
	4447: uint16(anon_sym_input),
	4448: uint16(anon_sym_query),
	4449: uint16(anon_sym_mutation),
	4450: uint16(anon_sym_subscription),
	4451: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4452: uint16(anon_sym_fragment),
	4453: uint16(anon_sym_directive),
	4454: uint16(3),
	4455: uint16(548),
	4456: uint16(1),
	4457: uint16(anon_sym_DQUOTE),
	4458: uint16(3),
	4459: uint16(2),
	4460: uint16(sym_comment),
	4461: uint16(sym_comma),
	4462: uint16(546),
	4463: uint16(16),
	4465: uint16(anon_sym_schema),
	4466: uint16(anon_sym_LBRACE),
	4467: uint16(anon_sym_extend),
	4468: uint16(anon_sym_scalar),
	4469: uint16(anon_sym_type),
	4470: uint16(anon_sym_interface),
	4471: uint16(anon_sym_union),
	4472: uint16(anon_sym_enum),
	4473: uint16(anon_sym_input),
	4474: uint16(anon_sym_query),
	4475: uint16(anon_sym_mutation),
	4476: uint16(anon_sym_subscription),
	4477: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4478: uint16(anon_sym_fragment),
	4479: uint16(anon_sym_directive),
	4480: uint16(3),
	4481: uint16(336),
	4482: uint16(1),
	4483: uint16(anon_sym_DQUOTE),
	4484: uint16(3),
	4485: uint16(2),
	4486: uint16(sym_comment),
	4487: uint16(sym_comma),
	4488: uint16(334),
	4489: uint16(16),
	4491: uint16(anon_sym_schema),
	4492: uint16(anon_sym_LBRACE),
	4493: uint16(anon_sym_extend),
	4494: uint16(anon_sym_scalar),
	4495: uint16(anon_sym_type),
	4496: uint16(anon_sym_interface),
	4497: uint16(anon_sym_union),
	4498: uint16(anon_sym_enum),
	4499: uint16(anon_sym_input),
	4500: uint16(anon_sym_query),
	4501: uint16(anon_sym_mutation),
	4502: uint16(anon_sym_subscription),
	4503: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4504: uint16(anon_sym_fragment),
	4505: uint16(anon_sym_directive),
	4506: uint16(3),
	4507: uint16(552),
	4508: uint16(1),
	4509: uint16(anon_sym_DQUOTE),
	4510: uint16(3),
	4511: uint16(2),
	4512: uint16(sym_comment),
	4513: uint16(sym_comma),
	4514: uint16(550),
	4515: uint16(16),
	4517: uint16(anon_sym_schema),
	4518: uint16(anon_sym_LBRACE),
	4519: uint16(anon_sym_extend),
	4520: uint16(anon_sym_scalar),
	4521: uint16(anon_sym_type),
	4522: uint16(anon_sym_interface),
	4523: uint16(anon_sym_union),
	4524: uint16(anon_sym_enum),
	4525: uint16(anon_sym_input),
	4526: uint16(anon_sym_query),
	4527: uint16(anon_sym_mutation),
	4528: uint16(anon_sym_subscription),
	4529: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4530: uint16(anon_sym_fragment),
	4531: uint16(anon_sym_directive),
	4532: uint16(3),
	4533: uint16(159),
	4534: uint16(1),
	4535: uint16(anon_sym_DQUOTE),
	4536: uint16(3),
	4537: uint16(2),
	4538: uint16(sym_comment),
	4539: uint16(sym_comma),
	4540: uint16(157),
	4541: uint16(16),
	4543: uint16(anon_sym_schema),
	4544: uint16(anon_sym_LBRACE),
	4545: uint16(anon_sym_extend),
	4546: uint16(anon_sym_scalar),
	4547: uint16(anon_sym_type),
	4548: uint16(anon_sym_interface),
	4549: uint16(anon_sym_union),
	4550: uint16(anon_sym_enum),
	4551: uint16(anon_sym_input),
	4552: uint16(anon_sym_query),
	4553: uint16(anon_sym_mutation),
	4554: uint16(anon_sym_subscription),
	4555: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4556: uint16(anon_sym_fragment),
	4557: uint16(anon_sym_directive),
	4558: uint16(3),
	4559: uint16(145),
	4560: uint16(1),
	4561: uint16(anon_sym_DQUOTE),
	4562: uint16(3),
	4563: uint16(2),
	4564: uint16(sym_comment),
	4565: uint16(sym_comma),
	4566: uint16(143),
	4567: uint16(16),
	4569: uint16(anon_sym_schema),
	4570: uint16(anon_sym_LBRACE),
	4571: uint16(anon_sym_extend),
	4572: uint16(anon_sym_scalar),
	4573: uint16(anon_sym_type),
	4574: uint16(anon_sym_interface),
	4575: uint16(anon_sym_union),
	4576: uint16(anon_sym_enum),
	4577: uint16(anon_sym_input),
	4578: uint16(anon_sym_query),
	4579: uint16(anon_sym_mutation),
	4580: uint16(anon_sym_subscription),
	4581: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4582: uint16(anon_sym_fragment),
	4583: uint16(anon_sym_directive),
	4584: uint16(3),
	4585: uint16(332),
	4586: uint16(1),
	4587: uint16(anon_sym_DQUOTE),
	4588: uint16(3),
	4589: uint16(2),
	4590: uint16(sym_comment),
	4591: uint16(sym_comma),
	4592: uint16(330),
	4593: uint16(16),
	4595: uint16(anon_sym_schema),
	4596: uint16(anon_sym_LBRACE),
	4597: uint16(anon_sym_extend),
	4598: uint16(anon_sym_scalar),
	4599: uint16(anon_sym_type),
	4600: uint16(anon_sym_interface),
	4601: uint16(anon_sym_union),
	4602: uint16(anon_sym_enum),
	4603: uint16(anon_sym_input),
	4604: uint16(anon_sym_query),
	4605: uint16(anon_sym_mutation),
	4606: uint16(anon_sym_subscription),
	4607: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4608: uint16(anon_sym_fragment),
	4609: uint16(anon_sym_directive),
	4610: uint16(3),
	4611: uint16(3),
	4612: uint16(2),
	4613: uint16(sym_comment),
	4614: uint16(sym_comma),
	4615: uint16(554),
	4616: uint16(6),
	4617: uint16(anon_sym_LBRACE),
	4618: uint16(anon_sym_DOLLAR),
	4619: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4620: uint16(sym_float_value),
	4621: uint16(anon_sym_LBRACK),
	4622: uint16(anon_sym_RBRACK),
	4623: uint16(556),
	4624: uint16(6),
	4625: uint16(anon_sym_DQUOTE),
	4626: uint16(sym_int_value),
	4627: uint16(anon_sym_true),
	4628: uint16(anon_sym_false),
	4629: uint16(sym_null_value),
	4630: uint16(sym_name),
	4631: uint16(3),
	4632: uint16(3),
	4633: uint16(2),
	4634: uint16(sym_comment),
	4635: uint16(sym_comma),
	4636: uint16(558),
	4637: uint16(6),
	4638: uint16(anon_sym_LBRACE),
	4639: uint16(anon_sym_DOLLAR),
	4640: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4641: uint16(sym_float_value),
	4642: uint16(anon_sym_LBRACK),
	4643: uint16(anon_sym_RBRACK),
	4644: uint16(560),
	4645: uint16(6),
	4646: uint16(anon_sym_DQUOTE),
	4647: uint16(sym_int_value),
	4648: uint16(anon_sym_true),
	4649: uint16(anon_sym_false),
	4650: uint16(sym_null_value),
	4651: uint16(sym_name),
	4652: uint16(3),
	4653: uint16(3),
	4654: uint16(2),
	4655: uint16(sym_comment),
	4656: uint16(sym_comma),
	4657: uint16(562),
	4658: uint16(6),
	4659: uint16(anon_sym_LBRACE),
	4660: uint16(anon_sym_DOLLAR),
	4661: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4662: uint16(sym_float_value),
	4663: uint16(anon_sym_LBRACK),
	4664: uint16(anon_sym_RBRACK),
	4665: uint16(564),
	4666: uint16(6),
	4667: uint16(anon_sym_DQUOTE),
	4668: uint16(sym_int_value),
	4669: uint16(anon_sym_true),
	4670: uint16(anon_sym_false),
	4671: uint16(sym_null_value),
	4672: uint16(sym_name),
	4673: uint16(3),
	4674: uint16(3),
	4675: uint16(2),
	4676: uint16(sym_comment),
	4677: uint16(sym_comma),
	4678: uint16(566),
	4679: uint16(6),
	4680: uint16(anon_sym_LBRACE),
	4681: uint16(anon_sym_DOLLAR),
	4682: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4683: uint16(sym_float_value),
	4684: uint16(anon_sym_LBRACK),
	4685: uint16(anon_sym_RBRACK),
	4686: uint16(568),
	4687: uint16(6),
	4688: uint16(anon_sym_DQUOTE),
	4689: uint16(sym_int_value),
	4690: uint16(anon_sym_true),
	4691: uint16(anon_sym_false),
	4692: uint16(sym_null_value),
	4693: uint16(sym_name),
	4694: uint16(3),
	4695: uint16(3),
	4696: uint16(2),
	4697: uint16(sym_comment),
	4698: uint16(sym_comma),
	4699: uint16(570),
	4700: uint16(6),
	4701: uint16(anon_sym_LBRACE),
	4702: uint16(anon_sym_DOLLAR),
	4703: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4704: uint16(sym_float_value),
	4705: uint16(anon_sym_LBRACK),
	4706: uint16(anon_sym_RBRACK),
	4707: uint16(572),
	4708: uint16(6),
	4709: uint16(anon_sym_DQUOTE),
	4710: uint16(sym_int_value),
	4711: uint16(anon_sym_true),
	4712: uint16(anon_sym_false),
	4713: uint16(sym_null_value),
	4714: uint16(sym_name),
	4715: uint16(3),
	4716: uint16(3),
	4717: uint16(2),
	4718: uint16(sym_comment),
	4719: uint16(sym_comma),
	4720: uint16(574),
	4721: uint16(6),
	4722: uint16(anon_sym_LBRACE),
	4723: uint16(anon_sym_DOLLAR),
	4724: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4725: uint16(sym_float_value),
	4726: uint16(anon_sym_LBRACK),
	4727: uint16(anon_sym_RBRACK),
	4728: uint16(576),
	4729: uint16(6),
	4730: uint16(anon_sym_DQUOTE),
	4731: uint16(sym_int_value),
	4732: uint16(anon_sym_true),
	4733: uint16(anon_sym_false),
	4734: uint16(sym_null_value),
	4735: uint16(sym_name),
	4736: uint16(3),
	4737: uint16(3),
	4738: uint16(2),
	4739: uint16(sym_comment),
	4740: uint16(sym_comma),
	4741: uint16(578),
	4742: uint16(6),
	4743: uint16(anon_sym_LBRACE),
	4744: uint16(anon_sym_DOLLAR),
	4745: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4746: uint16(sym_float_value),
	4747: uint16(anon_sym_LBRACK),
	4748: uint16(anon_sym_RBRACK),
	4749: uint16(580),
	4750: uint16(6),
	4751: uint16(anon_sym_DQUOTE),
	4752: uint16(sym_int_value),
	4753: uint16(anon_sym_true),
	4754: uint16(anon_sym_false),
	4755: uint16(sym_null_value),
	4756: uint16(sym_name),
	4757: uint16(3),
	4758: uint16(3),
	4759: uint16(2),
	4760: uint16(sym_comment),
	4761: uint16(sym_comma),
	4762: uint16(582),
	4763: uint16(6),
	4764: uint16(anon_sym_LBRACE),
	4765: uint16(anon_sym_DOLLAR),
	4766: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4767: uint16(sym_float_value),
	4768: uint16(anon_sym_LBRACK),
	4769: uint16(anon_sym_RBRACK),
	4770: uint16(584),
	4771: uint16(6),
	4772: uint16(anon_sym_DQUOTE),
	4773: uint16(sym_int_value),
	4774: uint16(anon_sym_true),
	4775: uint16(anon_sym_false),
	4776: uint16(sym_null_value),
	4777: uint16(sym_name),
	4778: uint16(3),
	4779: uint16(3),
	4780: uint16(2),
	4781: uint16(sym_comment),
	4782: uint16(sym_comma),
	4783: uint16(586),
	4784: uint16(6),
	4785: uint16(anon_sym_LBRACE),
	4786: uint16(anon_sym_DOLLAR),
	4787: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4788: uint16(sym_float_value),
	4789: uint16(anon_sym_LBRACK),
	4790: uint16(anon_sym_RBRACK),
	4791: uint16(588),
	4792: uint16(6),
	4793: uint16(anon_sym_DQUOTE),
	4794: uint16(sym_int_value),
	4795: uint16(anon_sym_true),
	4796: uint16(anon_sym_false),
	4797: uint16(sym_null_value),
	4798: uint16(sym_name),
	4799: uint16(8),
	4800: uint16(592),
	4801: uint16(1),
	4802: uint16(anon_sym_EQ),
	4803: uint16(594),
	4804: uint16(1),
	4805: uint16(anon_sym_DQUOTE),
	4806: uint16(596),
	4807: uint16(1),
	4808: uint16(anon_sym_AT),
	4809: uint16(163),
	4810: uint16(1),
	4811: uint16(sym_default_value),
	4812: uint16(230),
	4813: uint16(1),
	4814: uint16(sym_directives),
	4815: uint16(3),
	4816: uint16(2),
	4817: uint16(sym_comment),
	4818: uint16(sym_comma),
	4819: uint16(152),
	4820: uint16(2),
	4821: uint16(sym_directive),
	4822: uint16(aux_sym_directives_repeat1),
	4823: uint16(590),
	4824: uint16(4),
	4825: uint16(anon_sym_RBRACE),
	4826: uint16(anon_sym_RPAREN),
	4827: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4828: uint16(sym_name),
	4829: uint16(8),
	4830: uint16(592),
	4831: uint16(1),
	4832: uint16(anon_sym_EQ),
	4833: uint16(596),
	4834: uint16(1),
	4835: uint16(anon_sym_AT),
	4836: uint16(600),
	4837: uint16(1),
	4838: uint16(anon_sym_DQUOTE),
	4839: uint16(154),
	4840: uint16(1),
	4841: uint16(sym_default_value),
	4842: uint16(242),
	4843: uint16(1),
	4844: uint16(sym_directives),
	4845: uint16(3),
	4846: uint16(2),
	4847: uint16(sym_comment),
	4848: uint16(sym_comma),
	4849: uint16(152),
	4850: uint16(2),
	4851: uint16(sym_directive),
	4852: uint16(aux_sym_directives_repeat1),
	4853: uint16(598),
	4854: uint16(4),
	4855: uint16(anon_sym_RBRACE),
	4856: uint16(anon_sym_RPAREN),
	4857: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4858: uint16(sym_name),
	4859: uint16(10),
	4860: uint16(596),
	4861: uint16(1),
	4862: uint16(anon_sym_AT),
	4863: uint16(602),
	4864: uint16(1),
	4865: uint16(anon_sym_LBRACE),
	4866: uint16(604),
	4867: uint16(1),
	4868: uint16(anon_sym_on),
	4869: uint16(606),
	4870: uint16(1),
	4871: uint16(sym_name),
	4872: uint16(197),
	4873: uint16(1),
	4874: uint16(sym_fragment_name),
	4875: uint16(215),
	4876: uint16(1),
	4877: uint16(sym_type_condition),
	4878: uint16(286),
	4879: uint16(1),
	4880: uint16(sym_selection_set),
	4881: uint16(324),
	4882: uint16(1),
	4883: uint16(sym_directives),
	4884: uint16(3),
	4885: uint16(2),
	4886: uint16(sym_comment),
	4887: uint16(sym_comma),
	4888: uint16(256),
	4889: uint16(2),
	4890: uint16(sym_directive),
	4891: uint16(aux_sym_directives_repeat1),
	4892: uint16(9),
	4893: uint16(596),
	4894: uint16(1),
	4895: uint16(anon_sym_AT),
	4896: uint16(602),
	4897: uint16(1),
	4898: uint16(anon_sym_LBRACE),
	4899: uint16(610),
	4900: uint16(1),
	4901: uint16(anon_sym_COLON),
	4902: uint16(612),
	4903: uint16(1),
	4904: uint16(anon_sym_LPAREN),
	4905: uint16(198),
	4906: uint16(1),
	4907: uint16(sym_arguments),
	4908: uint16(229),
	4909: uint16(1),
	4910: uint16(sym_directive),
	4911: uint16(280),
	4912: uint16(1),
	4913: uint16(sym_selection_set),
	4914: uint16(3),
	4915: uint16(2),
	4916: uint16(sym_comment),
	4917: uint16(sym_comma),
	4918: uint16(608),
	4919: uint16(3),
	4920: uint16(anon_sym_RBRACE),
	4921: uint16(anon_sym_DOT_DOT_DOT),
	4922: uint16(sym_name),
	4923: uint16(5),
	4924: uint16(187),
	4925: uint16(1),
	4926: uint16(anon_sym_DQUOTE),
	4927: uint16(612),
	4928: uint16(1),
	4929: uint16(anon_sym_LPAREN),
	4930: uint16(187),
	4931: uint16(1),
	4932: uint16(sym_arguments),
	4933: uint16(3),
	4934: uint16(2),
	4935: uint16(sym_comment),
	4936: uint16(sym_comma),
	4937: uint16(183),
	4938: uint16(7),
	4939: uint16(anon_sym_LBRACE),
	4940: uint16(anon_sym_RBRACE),
	4941: uint16(anon_sym_RPAREN),
	4942: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4943: uint16(anon_sym_DOT_DOT_DOT),
	4944: uint16(anon_sym_AT),
	4945: uint16(sym_name),
	4946: uint16(5),
	4947: uint16(206),
	4948: uint16(1),
	4949: uint16(anon_sym_DQUOTE),
	4950: uint16(596),
	4951: uint16(1),
	4952: uint16(anon_sym_AT),
	4953: uint16(3),
	4954: uint16(2),
	4955: uint16(sym_comment),
	4956: uint16(sym_comma),
	4957: uint16(155),
	4958: uint16(2),
	4959: uint16(sym_directive),
	4960: uint16(aux_sym_directives_repeat1),
	4961: uint16(204),
	4962: uint16(5),
	4963: uint16(anon_sym_RBRACE),
	4964: uint16(anon_sym_RPAREN),
	4965: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4966: uint16(anon_sym_DOT_DOT_DOT),
	4967: uint16(sym_name),
	4968: uint16(8),
	4969: uint16(616),
	4970: uint16(1),
	4971: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	4972: uint16(619),
	4973: uint16(1),
	4974: uint16(anon_sym_DQUOTE),
	4975: uint16(622),
	4976: uint16(1),
	4977: uint16(sym_name),
	4978: uint16(345),
	4979: uint16(1),
	4980: uint16(sym_description),
	4981: uint16(361),
	4982: uint16(1),
	4983: uint16(sym_string_value),
	4984: uint16(3),
	4985: uint16(2),
	4986: uint16(sym_comment),
	4987: uint16(sym_comma),
	4988: uint16(614),
	4989: uint16(2),
	4990: uint16(anon_sym_RBRACE),
	4991: uint16(anon_sym_RPAREN),
	4992: uint16(153),
	4993: uint16(2),
	4994: uint16(sym_input_value_definition),
	4995: uint16(aux_sym_input_fields_definition_repeat1),
	4996: uint16(6),
	4997: uint16(594),
	4998: uint16(1),
	4999: uint16(anon_sym_DQUOTE),
	5000: uint16(596),
	5001: uint16(1),
	5002: uint16(anon_sym_AT),
	5003: uint16(230),
	5004: uint16(1),
	5005: uint16(sym_directives),
	5006: uint16(3),
	5007: uint16(2),
	5008: uint16(sym_comment),
	5009: uint16(sym_comma),
	5010: uint16(152),
	5011: uint16(2),
	5012: uint16(sym_directive),
	5013: uint16(aux_sym_directives_repeat1),
	5014: uint16(590),
	5015: uint16(4),
	5016: uint16(anon_sym_RBRACE),
	5017: uint16(anon_sym_RPAREN),
	5018: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5019: uint16(sym_name),
	5020: uint16(5),
	5021: uint16(195),
	5022: uint16(1),
	5023: uint16(anon_sym_DQUOTE),
	5024: uint16(625),
	5025: uint16(1),
	5026: uint16(anon_sym_AT),
	5027: uint16(3),
	5028: uint16(2),
	5029: uint16(sym_comment),
	5030: uint16(sym_comma),
	5031: uint16(155),
	5032: uint16(2),
	5033: uint16(sym_directive),
	5034: uint16(aux_sym_directives_repeat1),
	5035: uint16(193),
	5036: uint16(5),
	5037: uint16(anon_sym_RBRACE),
	5038: uint16(anon_sym_RPAREN),
	5039: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5040: uint16(anon_sym_DOT_DOT_DOT),
	5041: uint16(sym_name),
	5042: uint16(9),
	5043: uint16(7),
	5044: uint16(1),
	5045: uint16(anon_sym_LBRACE),
	5046: uint16(596),
	5047: uint16(1),
	5048: uint16(anon_sym_AT),
	5049: uint16(628),
	5050: uint16(1),
	5051: uint16(anon_sym_LPAREN),
	5052: uint16(630),
	5053: uint16(1),
	5054: uint16(sym_name),
	5055: uint16(129),
	5056: uint16(1),
	5057: uint16(sym_selection_set),
	5058: uint16(217),
	5059: uint16(1),
	5060: uint16(sym_variable_definitions),
	5061: uint16(323),
	5062: uint16(1),
	5063: uint16(sym_directives),
	5064: uint16(3),
	5065: uint16(2),
	5066: uint16(sym_comment),
	5067: uint16(sym_comma),
	5068: uint16(256),
	5069: uint16(2),
	5070: uint16(sym_directive),
	5071: uint16(aux_sym_directives_repeat1),
	5072: uint16(7),
	5073: uint16(632),
	5074: uint16(1),
	5075: uint16(anon_sym_RBRACE),
	5076: uint16(634),
	5077: uint16(1),
	5078: uint16(anon_sym_DOT_DOT_DOT),
	5079: uint16(637),
	5080: uint16(1),
	5081: uint16(sym_name),
	5082: uint16(359),
	5083: uint16(1),
	5084: uint16(sym_alias),
	5085: uint16(3),
	5086: uint16(2),
	5087: uint16(sym_comment),
	5088: uint16(sym_comma),
	5089: uint16(157),
	5090: uint16(2),
	5091: uint16(sym_selection),
	5092: uint16(aux_sym_selection_set_repeat1),
	5093: uint16(297),
	5094: uint16(3),
	5095: uint16(sym_field),
	5096: uint16(sym_fragment_spread),
	5097: uint16(sym_inline_fragment),
	5098: uint16(8),
	5099: uint16(596),
	5100: uint16(1),
	5101: uint16(anon_sym_AT),
	5102: uint16(602),
	5103: uint16(1),
	5104: uint16(anon_sym_LBRACE),
	5105: uint16(612),
	5106: uint16(1),
	5107: uint16(anon_sym_LPAREN),
	5108: uint16(195),
	5109: uint16(1),
	5110: uint16(sym_arguments),
	5111: uint16(250),
	5112: uint16(1),
	5113: uint16(sym_directive),
	5114: uint16(288),
	5115: uint16(1),
	5116: uint16(sym_selection_set),
	5117: uint16(3),
	5118: uint16(2),
	5119: uint16(sym_comment),
	5120: uint16(sym_comma),
	5121: uint16(640),
	5122: uint16(3),
	5123: uint16(anon_sym_RBRACE),
	5124: uint16(anon_sym_DOT_DOT_DOT),
	5125: uint16(sym_name),
	5126: uint16(7),
	5127: uint16(642),
	5128: uint16(1),
	5129: uint16(anon_sym_RBRACE),
	5130: uint16(644),
	5131: uint16(1),
	5132: uint16(anon_sym_DOT_DOT_DOT),
	5133: uint16(646),
	5134: uint16(1),
	5135: uint16(sym_name),
	5136: uint16(359),
	5137: uint16(1),
	5138: uint16(sym_alias),
	5139: uint16(3),
	5140: uint16(2),
	5141: uint16(sym_comment),
	5142: uint16(sym_comma),
	5143: uint16(157),
	5144: uint16(2),
	5145: uint16(sym_selection),
	5146: uint16(aux_sym_selection_set_repeat1),
	5147: uint16(297),
	5148: uint16(3),
	5149: uint16(sym_field),
	5150: uint16(sym_fragment_spread),
	5151: uint16(sym_inline_fragment),
	5152: uint16(3),
	5153: uint16(650),
	5154: uint16(1),
	5155: uint16(anon_sym_DQUOTE),
	5156: uint16(3),
	5157: uint16(2),
	5158: uint16(sym_comment),
	5159: uint16(sym_comma),
	5160: uint16(648),
	5161: uint16(8),
	5162: uint16(anon_sym_RBRACE),
	5163: uint16(anon_sym_RPAREN),
	5164: uint16(anon_sym_EQ),
	5165: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5166: uint16(anon_sym_RBRACK),
	5167: uint16(anon_sym_AT),
	5168: uint16(anon_sym_BANG),
	5169: uint16(sym_name),
	5170: uint16(4),
	5171: uint16(654),
	5172: uint16(1),
	5173: uint16(anon_sym_DQUOTE),
	5174: uint16(656),
	5175: uint16(1),
	5176: uint16(anon_sym_BANG),
	5177: uint16(3),
	5178: uint16(2),
	5179: uint16(sym_comment),
	5180: uint16(sym_comma),
	5181: uint16(652),
	5182: uint16(7),
	5183: uint16(anon_sym_RBRACE),
	5184: uint16(anon_sym_RPAREN),
	5185: uint16(anon_sym_EQ),
	5186: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5187: uint16(anon_sym_RBRACK),
	5188: uint16(anon_sym_AT),
	5189: uint16(sym_name),
	5190: uint16(9),
	5191: uint16(310),
	5192: uint16(1),
	5193: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5194: uint16(312),
	5195: uint16(1),
	5196: uint16(anon_sym_DQUOTE),
	5197: uint16(658),
	5198: uint16(1),
	5199: uint16(anon_sym_RBRACE),
	5200: uint16(660),
	5201: uint16(1),
	5202: uint16(sym_name),
	5203: uint16(174),
	5204: uint16(1),
	5205: uint16(sym_enum_value),
	5206: uint16(319),
	5207: uint16(1),
	5208: uint16(sym_description),
	5209: uint16(361),
	5210: uint16(1),
	5211: uint16(sym_string_value),
	5212: uint16(3),
	5213: uint16(2),
	5214: uint16(sym_comment),
	5215: uint16(sym_comma),
	5216: uint16(164),
	5217: uint16(2),
	5218: uint16(sym_enum_value_definition),
	5219: uint16(aux_sym_enum_values_definition_repeat1),
	5220: uint16(6),
	5221: uint16(596),
	5222: uint16(1),
	5223: uint16(anon_sym_AT),
	5224: uint16(664),
	5225: uint16(1),
	5226: uint16(anon_sym_DQUOTE),
	5227: uint16(236),
	5228: uint16(1),
	5229: uint16(sym_directives),
	5230: uint16(3),
	5231: uint16(2),
	5232: uint16(sym_comment),
	5233: uint16(sym_comma),
	5234: uint16(152),
	5235: uint16(2),
	5236: uint16(sym_directive),
	5237: uint16(aux_sym_directives_repeat1),
	5238: uint16(662),
	5239: uint16(4),
	5240: uint16(anon_sym_RBRACE),
	5241: uint16(anon_sym_RPAREN),
	5242: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5243: uint16(sym_name),
	5244: uint16(9),
	5245: uint16(666),
	5246: uint16(1),
	5247: uint16(anon_sym_RBRACE),
	5248: uint16(668),
	5249: uint16(1),
	5250: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5251: uint16(671),
	5252: uint16(1),
	5253: uint16(anon_sym_DQUOTE),
	5254: uint16(674),
	5255: uint16(1),
	5256: uint16(sym_name),
	5257: uint16(174),
	5258: uint16(1),
	5259: uint16(sym_enum_value),
	5260: uint16(319),
	5261: uint16(1),
	5262: uint16(sym_description),
	5263: uint16(361),
	5264: uint16(1),
	5265: uint16(sym_string_value),
	5266: uint16(3),
	5267: uint16(2),
	5268: uint16(sym_comment),
	5269: uint16(sym_comma),
	5270: uint16(164),
	5271: uint16(2),
	5272: uint16(sym_enum_value_definition),
	5273: uint16(aux_sym_enum_values_definition_repeat1),
	5274: uint16(7),
	5275: uint16(644),
	5276: uint16(1),
	5277: uint16(anon_sym_DOT_DOT_DOT),
	5278: uint16(646),
	5279: uint16(1),
	5280: uint16(sym_name),
	5281: uint16(677),
	5282: uint16(1),
	5283: uint16(anon_sym_RBRACE),
	5284: uint16(359),
	5285: uint16(1),
	5286: uint16(sym_alias),
	5287: uint16(3),
	5288: uint16(2),
	5289: uint16(sym_comment),
	5290: uint16(sym_comma),
	5291: uint16(157),
	5292: uint16(2),
	5293: uint16(sym_selection),
	5294: uint16(aux_sym_selection_set_repeat1),
	5295: uint16(297),
	5296: uint16(3),
	5297: uint16(sym_field),
	5298: uint16(sym_fragment_spread),
	5299: uint16(sym_inline_fragment),
	5300: uint16(2),
	5301: uint16(3),
	5302: uint16(2),
	5303: uint16(sym_comment),
	5304: uint16(sym_comma),
	5305: uint16(679),
	5306: uint16(8),
	5307: uint16(anon_sym_schema),
	5308: uint16(anon_sym_scalar),
	5309: uint16(anon_sym_type),
	5310: uint16(anon_sym_interface),
	5311: uint16(anon_sym_union),
	5312: uint16(anon_sym_enum),
	5313: uint16(anon_sym_input),
	5314: uint16(anon_sym_directive),
	5315: uint16(6),
	5316: uint16(644),
	5317: uint16(1),
	5318: uint16(anon_sym_DOT_DOT_DOT),
	5319: uint16(646),
	5320: uint16(1),
	5321: uint16(sym_name),
	5322: uint16(359),
	5323: uint16(1),
	5324: uint16(sym_alias),
	5325: uint16(3),
	5326: uint16(2),
	5327: uint16(sym_comment),
	5328: uint16(sym_comma),
	5329: uint16(165),
	5330: uint16(2),
	5331: uint16(sym_selection),
	5332: uint16(aux_sym_selection_set_repeat1),
	5333: uint16(297),
	5334: uint16(3),
	5335: uint16(sym_field),
	5336: uint16(sym_fragment_spread),
	5337: uint16(sym_inline_fragment),
	5338: uint16(3),
	5339: uint16(127),
	5340: uint16(1),
	5341: uint16(anon_sym_DQUOTE),
	5342: uint16(3),
	5343: uint16(2),
	5344: uint16(sym_comment),
	5345: uint16(sym_comma),
	5346: uint16(125),
	5347: uint16(7),
	5348: uint16(anon_sym_RBRACE),
	5349: uint16(anon_sym_RPAREN),
	5350: uint16(anon_sym_EQ),
	5351: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5352: uint16(anon_sym_AT),
	5353: uint16(anon_sym_BANG),
	5354: uint16(sym_name),
	5355: uint16(8),
	5356: uint16(310),
	5357: uint16(1),
	5358: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5359: uint16(312),
	5360: uint16(1),
	5361: uint16(anon_sym_DQUOTE),
	5362: uint16(681),
	5363: uint16(1),
	5364: uint16(anon_sym_RPAREN),
	5365: uint16(683),
	5366: uint16(1),
	5367: uint16(sym_name),
	5368: uint16(345),
	5369: uint16(1),
	5370: uint16(sym_description),
	5371: uint16(361),
	5372: uint16(1),
	5373: uint16(sym_string_value),
	5374: uint16(3),
	5375: uint16(2),
	5376: uint16(sym_comment),
	5377: uint16(sym_comma),
	5378: uint16(153),
	5379: uint16(2),
	5380: uint16(sym_input_value_definition),
	5381: uint16(aux_sym_input_fields_definition_repeat1),
	5382: uint16(8),
	5383: uint16(310),
	5384: uint16(1),
	5385: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5386: uint16(312),
	5387: uint16(1),
	5388: uint16(anon_sym_DQUOTE),
	5389: uint16(685),
	5390: uint16(1),
	5391: uint16(anon_sym_RBRACE),
	5392: uint16(687),
	5393: uint16(1),
	5394: uint16(sym_name),
	5395: uint16(329),
	5396: uint16(1),
	5397: uint16(sym_description),
	5398: uint16(361),
	5399: uint16(1),
	5400: uint16(sym_string_value),
	5401: uint16(3),
	5402: uint16(2),
	5403: uint16(sym_comment),
	5404: uint16(sym_comma),
	5405: uint16(185),
	5406: uint16(2),
	5407: uint16(sym_field_definition),
	5408: uint16(aux_sym_fields_definition_repeat1),
	5409: uint16(6),
	5410: uint16(596),
	5411: uint16(1),
	5412: uint16(anon_sym_AT),
	5413: uint16(691),
	5414: uint16(1),
	5415: uint16(anon_sym_DQUOTE),
	5416: uint16(272),
	5417: uint16(1),
	5418: uint16(sym_directives),
	5419: uint16(3),
	5420: uint16(2),
	5421: uint16(sym_comment),
	5422: uint16(sym_comma),
	5423: uint16(152),
	5424: uint16(2),
	5425: uint16(sym_directive),
	5426: uint16(aux_sym_directives_repeat1),
	5427: uint16(689),
	5428: uint16(3),
	5429: uint16(anon_sym_RBRACE),
	5430: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5431: uint16(sym_name),
	5432: uint16(8),
	5433: uint16(7),
	5434: uint16(1),
	5435: uint16(anon_sym_LBRACE),
	5436: uint16(596),
	5437: uint16(1),
	5438: uint16(anon_sym_AT),
	5439: uint16(628),
	5440: uint16(1),
	5441: uint16(anon_sym_LPAREN),
	5442: uint16(124),
	5443: uint16(1),
	5444: uint16(sym_selection_set),
	5445: uint16(202),
	5446: uint16(1),
	5447: uint16(sym_variable_definitions),
	5448: uint16(304),
	5449: uint16(1),
	5450: uint16(sym_directives),
	5451: uint16(3),
	5452: uint16(2),
	5453: uint16(sym_comment),
	5454: uint16(sym_comma),
	5455: uint16(256),
	5456: uint16(2),
	5457: uint16(sym_directive),
	5458: uint16(aux_sym_directives_repeat1),
	5459: uint16(6),
	5460: uint16(596),
	5461: uint16(1),
	5462: uint16(anon_sym_AT),
	5463: uint16(695),
	5464: uint16(1),
	5465: uint16(anon_sym_DQUOTE),
	5466: uint16(271),
	5467: uint16(1),
	5468: uint16(sym_directives),
	5469: uint16(3),
	5470: uint16(2),
	5471: uint16(sym_comment),
	5472: uint16(sym_comma),
	5473: uint16(152),
	5474: uint16(2),
	5475: uint16(sym_directive),
	5476: uint16(aux_sym_directives_repeat1),
	5477: uint16(693),
	5478: uint16(3),
	5479: uint16(anon_sym_RBRACE),
	5480: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5481: uint16(sym_name),
	5482: uint16(6),
	5483: uint16(596),
	5484: uint16(1),
	5485: uint16(anon_sym_AT),
	5486: uint16(699),
	5487: uint16(1),
	5488: uint16(anon_sym_DQUOTE),
	5489: uint16(268),
	5490: uint16(1),
	5491: uint16(sym_directives),
	5492: uint16(3),
	5493: uint16(2),
	5494: uint16(sym_comment),
	5495: uint16(sym_comma),
	5496: uint16(152),
	5497: uint16(2),
	5498: uint16(sym_directive),
	5499: uint16(aux_sym_directives_repeat1),
	5500: uint16(697),
	5501: uint16(3),
	5502: uint16(anon_sym_RBRACE),
	5503: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5504: uint16(sym_name),
	5505: uint16(3),
	5506: uint16(703),
	5507: uint16(1),
	5508: uint16(anon_sym_DQUOTE),
	5509: uint16(3),
	5510: uint16(2),
	5511: uint16(sym_comment),
	5512: uint16(sym_comma),
	5513: uint16(701),
	5514: uint16(7),
	5515: uint16(anon_sym_RBRACE),
	5516: uint16(anon_sym_RPAREN),
	5517: uint16(anon_sym_EQ),
	5518: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5519: uint16(anon_sym_RBRACK),
	5520: uint16(anon_sym_AT),
	5521: uint16(sym_name),
	5522: uint16(2),
	5523: uint16(3),
	5524: uint16(2),
	5525: uint16(sym_comment),
	5526: uint16(sym_comma),
	5527: uint16(570),
	5528: uint16(8),
	5529: uint16(anon_sym_schema),
	5530: uint16(anon_sym_scalar),
	5531: uint16(anon_sym_type),
	5532: uint16(anon_sym_interface),
	5533: uint16(anon_sym_union),
	5534: uint16(anon_sym_enum),
	5535: uint16(anon_sym_input),
	5536: uint16(anon_sym_directive),
	5537: uint16(6),
	5538: uint16(644),
	5539: uint16(1),
	5540: uint16(anon_sym_DOT_DOT_DOT),
	5541: uint16(646),
	5542: uint16(1),
	5543: uint16(sym_name),
	5544: uint16(359),
	5545: uint16(1),
	5546: uint16(sym_alias),
	5547: uint16(3),
	5548: uint16(2),
	5549: uint16(sym_comment),
	5550: uint16(sym_comma),
	5551: uint16(159),
	5552: uint16(2),
	5553: uint16(sym_selection),
	5554: uint16(aux_sym_selection_set_repeat1),
	5555: uint16(297),
	5556: uint16(3),
	5557: uint16(sym_field),
	5558: uint16(sym_fragment_spread),
	5559: uint16(sym_inline_fragment),
	5560: uint16(3),
	5561: uint16(654),
	5562: uint16(1),
	5563: uint16(anon_sym_DQUOTE),
	5564: uint16(3),
	5565: uint16(2),
	5566: uint16(sym_comment),
	5567: uint16(sym_comma),
	5568: uint16(652),
	5569: uint16(7),
	5570: uint16(anon_sym_RBRACE),
	5571: uint16(anon_sym_RPAREN),
	5572: uint16(anon_sym_EQ),
	5573: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5574: uint16(anon_sym_RBRACK),
	5575: uint16(anon_sym_AT),
	5576: uint16(sym_name),
	5577: uint16(6),
	5578: uint16(596),
	5579: uint16(1),
	5580: uint16(anon_sym_AT),
	5581: uint16(707),
	5582: uint16(1),
	5583: uint16(anon_sym_DQUOTE),
	5584: uint16(254),
	5585: uint16(1),
	5586: uint16(sym_directives),
	5587: uint16(3),
	5588: uint16(2),
	5589: uint16(sym_comment),
	5590: uint16(sym_comma),
	5591: uint16(152),
	5592: uint16(2),
	5593: uint16(sym_directive),
	5594: uint16(aux_sym_directives_repeat1),
	5595: uint16(705),
	5596: uint16(3),
	5597: uint16(anon_sym_RBRACE),
	5598: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5599: uint16(sym_name),
	5600: uint16(8),
	5601: uint16(310),
	5602: uint16(1),
	5603: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5604: uint16(312),
	5605: uint16(1),
	5606: uint16(anon_sym_DQUOTE),
	5607: uint16(683),
	5608: uint16(1),
	5609: uint16(sym_name),
	5610: uint16(709),
	5611: uint16(1),
	5612: uint16(anon_sym_RBRACE),
	5613: uint16(345),
	5614: uint16(1),
	5615: uint16(sym_description),
	5616: uint16(361),
	5617: uint16(1),
	5618: uint16(sym_string_value),
	5619: uint16(3),
	5620: uint16(2),
	5621: uint16(sym_comment),
	5622: uint16(sym_comma),
	5623: uint16(153),
	5624: uint16(2),
	5625: uint16(sym_input_value_definition),
	5626: uint16(aux_sym_input_fields_definition_repeat1),
	5627: uint16(8),
	5628: uint16(310),
	5629: uint16(1),
	5630: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5631: uint16(312),
	5632: uint16(1),
	5633: uint16(anon_sym_DQUOTE),
	5634: uint16(660),
	5635: uint16(1),
	5636: uint16(sym_name),
	5637: uint16(174),
	5638: uint16(1),
	5639: uint16(sym_enum_value),
	5640: uint16(319),
	5641: uint16(1),
	5642: uint16(sym_description),
	5643: uint16(361),
	5644: uint16(1),
	5645: uint16(sym_string_value),
	5646: uint16(3),
	5647: uint16(2),
	5648: uint16(sym_comment),
	5649: uint16(sym_comma),
	5650: uint16(162),
	5651: uint16(2),
	5652: uint16(sym_enum_value_definition),
	5653: uint16(aux_sym_enum_values_definition_repeat1),
	5654: uint16(9),
	5655: uint16(711),
	5656: uint16(1),
	5657: uint16(anon_sym_schema),
	5658: uint16(713),
	5659: uint16(1),
	5660: uint16(anon_sym_scalar),
	5661: uint16(715),
	5662: uint16(1),
	5663: uint16(anon_sym_type),
	5664: uint16(717),
	5665: uint16(1),
	5666: uint16(anon_sym_interface),
	5667: uint16(719),
	5668: uint16(1),
	5669: uint16(anon_sym_union),
	5670: uint16(721),
	5671: uint16(1),
	5672: uint16(anon_sym_enum),
	5673: uint16(723),
	5674: uint16(1),
	5675: uint16(anon_sym_input),
	5676: uint16(725),
	5677: uint16(1),
	5678: uint16(anon_sym_directive),
	5679: uint16(3),
	5680: uint16(2),
	5681: uint16(sym_comment),
	5682: uint16(sym_comma),
	5683: uint16(3),
	5684: uint16(296),
	5685: uint16(1),
	5686: uint16(anon_sym_DQUOTE),
	5687: uint16(3),
	5688: uint16(2),
	5689: uint16(sym_comment),
	5690: uint16(sym_comma),
	5691: uint16(294),
	5692: uint16(7),
	5693: uint16(anon_sym_LBRACE),
	5694: uint16(anon_sym_RBRACE),
	5695: uint16(anon_sym_RPAREN),
	5696: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5697: uint16(anon_sym_DOT_DOT_DOT),
	5698: uint16(anon_sym_AT),
	5699: uint16(sym_name),
	5700: uint16(8),
	5701: uint16(3),
	5702: uint16(1),
	5703: uint16(sym_comment),
	5704: uint16(729),
	5705: uint16(1),
	5706: uint16(anon_sym_EQ),
	5707: uint16(731),
	5708: uint16(1),
	5709: uint16(anon_sym_AT),
	5710: uint16(733),
	5711: uint16(1),
	5712: uint16(sym_comma),
	5713: uint16(207),
	5714: uint16(1),
	5715: uint16(sym_default_value),
	5716: uint16(301),
	5717: uint16(1),
	5718: uint16(sym_directives),
	5719: uint16(727),
	5720: uint16(2),
	5721: uint16(anon_sym_RPAREN),
	5722: uint16(anon_sym_DOLLAR),
	5723: uint16(238),
	5724: uint16(2),
	5725: uint16(sym_directive),
	5726: uint16(aux_sym_directives_repeat1),
	5727: uint16(8),
	5728: uint16(735),
	5729: uint16(1),
	5730: uint16(anon_sym_RBRACE),
	5731: uint16(737),
	5732: uint16(1),
	5733: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5734: uint16(740),
	5735: uint16(1),
	5736: uint16(anon_sym_DQUOTE),
	5737: uint16(743),
	5738: uint16(1),
	5739: uint16(sym_name),
	5740: uint16(329),
	5741: uint16(1),
	5742: uint16(sym_description),
	5743: uint16(361),
	5744: uint16(1),
	5745: uint16(sym_string_value),
	5746: uint16(3),
	5747: uint16(2),
	5748: uint16(sym_comment),
	5749: uint16(sym_comma),
	5750: uint16(185),
	5751: uint16(2),
	5752: uint16(sym_field_definition),
	5753: uint16(aux_sym_fields_definition_repeat1),
	5754: uint16(6),
	5755: uint16(596),
	5756: uint16(1),
	5757: uint16(anon_sym_AT),
	5758: uint16(748),
	5759: uint16(1),
	5760: uint16(anon_sym_DQUOTE),
	5761: uint16(278),
	5762: uint16(1),
	5763: uint16(sym_directives),
	5764: uint16(3),
	5765: uint16(2),
	5766: uint16(sym_comment),
	5767: uint16(sym_comma),
	5768: uint16(152),
	5769: uint16(2),
	5770: uint16(sym_directive),
	5771: uint16(aux_sym_directives_repeat1),
	5772: uint16(746),
	5773: uint16(3),
	5774: uint16(anon_sym_RBRACE),
	5775: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5776: uint16(sym_name),
	5777: uint16(3),
	5778: uint16(300),
	5779: uint16(1),
	5780: uint16(anon_sym_DQUOTE),
	5781: uint16(3),
	5782: uint16(2),
	5783: uint16(sym_comment),
	5784: uint16(sym_comma),
	5785: uint16(298),
	5786: uint16(7),
	5787: uint16(anon_sym_LBRACE),
	5788: uint16(anon_sym_RBRACE),
	5789: uint16(anon_sym_RPAREN),
	5790: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5791: uint16(anon_sym_DOT_DOT_DOT),
	5792: uint16(anon_sym_AT),
	5793: uint16(sym_name),
	5794: uint16(3),
	5795: uint16(580),
	5796: uint16(1),
	5797: uint16(anon_sym_DQUOTE),
	5798: uint16(3),
	5799: uint16(2),
	5800: uint16(sym_comment),
	5801: uint16(sym_comma),
	5802: uint16(578),
	5803: uint16(6),
	5804: uint16(anon_sym_RBRACE),
	5805: uint16(anon_sym_COLON),
	5806: uint16(anon_sym_RPAREN),
	5807: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5808: uint16(anon_sym_AT),
	5809: uint16(sym_name),
	5810: uint16(7),
	5811: uint16(310),
	5812: uint16(1),
	5813: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5814: uint16(312),
	5815: uint16(1),
	5816: uint16(anon_sym_DQUOTE),
	5817: uint16(687),
	5818: uint16(1),
	5819: uint16(sym_name),
	5820: uint16(329),
	5821: uint16(1),
	5822: uint16(sym_description),
	5823: uint16(361),
	5824: uint16(1),
	5825: uint16(sym_string_value),
	5826: uint16(3),
	5827: uint16(2),
	5828: uint16(sym_comment),
	5829: uint16(sym_comma),
	5830: uint16(170),
	5831: uint16(2),
	5832: uint16(sym_field_definition),
	5833: uint16(aux_sym_fields_definition_repeat1),
	5834: uint16(5),
	5835: uint16(750),
	5836: uint16(1),
	5837: uint16(anon_sym_RBRACE),
	5838: uint16(338),
	5839: uint16(1),
	5840: uint16(sym_operation_type),
	5841: uint16(3),
	5842: uint16(2),
	5843: uint16(sym_comment),
	5844: uint16(sym_comma),
	5845: uint16(190),
	5846: uint16(2),
	5847: uint16(sym_root_operation_type_definition),
	5848: uint16(aux_sym_schema_definition_repeat1),
	5849: uint16(752),
	5850: uint16(3),
	5851: uint16(anon_sym_query),
	5852: uint16(anon_sym_mutation),
	5853: uint16(anon_sym_subscription),
	5854: uint16(7),
	5855: uint16(310),
	5856: uint16(1),
	5857: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5858: uint16(312),
	5859: uint16(1),
	5860: uint16(anon_sym_DQUOTE),
	5861: uint16(683),
	5862: uint16(1),
	5863: uint16(sym_name),
	5864: uint16(345),
	5865: uint16(1),
	5866: uint16(sym_description),
	5867: uint16(361),
	5868: uint16(1),
	5869: uint16(sym_string_value),
	5870: uint16(3),
	5871: uint16(2),
	5872: uint16(sym_comment),
	5873: uint16(sym_comma),
	5874: uint16(169),
	5875: uint16(2),
	5876: uint16(sym_input_value_definition),
	5877: uint16(aux_sym_input_fields_definition_repeat1),
	5878: uint16(5),
	5879: uint16(755),
	5880: uint16(1),
	5881: uint16(anon_sym_RBRACE),
	5882: uint16(338),
	5883: uint16(1),
	5884: uint16(sym_operation_type),
	5885: uint16(3),
	5886: uint16(2),
	5887: uint16(sym_comment),
	5888: uint16(sym_comma),
	5889: uint16(190),
	5890: uint16(2),
	5891: uint16(sym_root_operation_type_definition),
	5892: uint16(aux_sym_schema_definition_repeat1),
	5893: uint16(23),
	5894: uint16(3),
	5895: uint16(anon_sym_query),
	5896: uint16(anon_sym_mutation),
	5897: uint16(anon_sym_subscription),
	5898: uint16(5),
	5899: uint16(757),
	5900: uint16(1),
	5901: uint16(anon_sym_RBRACE),
	5902: uint16(338),
	5903: uint16(1),
	5904: uint16(sym_operation_type),
	5905: uint16(3),
	5906: uint16(2),
	5907: uint16(sym_comment),
	5908: uint16(sym_comma),
	5909: uint16(190),
	5910: uint16(2),
	5911: uint16(sym_root_operation_type_definition),
	5912: uint16(aux_sym_schema_definition_repeat1),
	5913: uint16(23),
	5914: uint16(3),
	5915: uint16(anon_sym_query),
	5916: uint16(anon_sym_mutation),
	5917: uint16(anon_sym_subscription),
	5918: uint16(7),
	5919: uint16(310),
	5920: uint16(1),
	5921: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	5922: uint16(312),
	5923: uint16(1),
	5924: uint16(anon_sym_DQUOTE),
	5925: uint16(683),
	5926: uint16(1),
	5927: uint16(sym_name),
	5928: uint16(345),
	5929: uint16(1),
	5930: uint16(sym_description),
	5931: uint16(361),
	5932: uint16(1),
	5933: uint16(sym_string_value),
	5934: uint16(3),
	5935: uint16(2),
	5936: uint16(sym_comment),
	5937: uint16(sym_comma),
	5938: uint16(180),
	5939: uint16(2),
	5940: uint16(sym_input_value_definition),
	5941: uint16(aux_sym_input_fields_definition_repeat1),
	5942: uint16(6),
	5943: uint16(596),
	5944: uint16(1),
	5945: uint16(anon_sym_AT),
	5946: uint16(602),
	5947: uint16(1),
	5948: uint16(anon_sym_LBRACE),
	5949: uint16(247),
	5950: uint16(1),
	5951: uint16(sym_directive),
	5952: uint16(291),
	5953: uint16(1),
	5954: uint16(sym_selection_set),
	5955: uint16(3),
	5956: uint16(2),
	5957: uint16(sym_comment),
	5958: uint16(sym_comma),
	5959: uint16(759),
	5960: uint16(3),
	5961: uint16(anon_sym_RBRACE),
	5962: uint16(anon_sym_DOT_DOT_DOT),
	5963: uint16(sym_name),
	5964: uint16(5),
	5965: uint16(761),
	5966: uint16(1),
	5967: uint16(anon_sym_RBRACE),
	5968: uint16(338),
	5969: uint16(1),
	5970: uint16(sym_operation_type),
	5971: uint16(3),
	5972: uint16(2),
	5973: uint16(sym_comment),
	5974: uint16(sym_comma),
	5975: uint16(190),
	5976: uint16(2),
	5977: uint16(sym_root_operation_type_definition),
	5978: uint16(aux_sym_schema_definition_repeat1),
	5979: uint16(23),
	5980: uint16(3),
	5981: uint16(anon_sym_query),
	5982: uint16(anon_sym_mutation),
	5983: uint16(anon_sym_subscription),
	5984: uint16(5),
	5985: uint16(596),
	5986: uint16(1),
	5987: uint16(anon_sym_AT),
	5988: uint16(290),
	5989: uint16(1),
	5990: uint16(sym_directives),
	5991: uint16(3),
	5992: uint16(2),
	5993: uint16(sym_comment),
	5994: uint16(sym_comma),
	5995: uint16(152),
	5996: uint16(2),
	5997: uint16(sym_directive),
	5998: uint16(aux_sym_directives_repeat1),
	5999: uint16(763),
	6000: uint16(3),
	6001: uint16(anon_sym_RBRACE),
	6002: uint16(anon_sym_DOT_DOT_DOT),
	6003: uint16(sym_name),
	6004: uint16(6),
	6005: uint16(596),
	6006: uint16(1),
	6007: uint16(anon_sym_AT),
	6008: uint16(602),
	6009: uint16(1),
	6010: uint16(anon_sym_LBRACE),
	6011: uint16(250),
	6012: uint16(1),
	6013: uint16(sym_directive),
	6014: uint16(288),
	6015: uint16(1),
	6016: uint16(sym_selection_set),
	6017: uint16(3),
	6018: uint16(2),
	6019: uint16(sym_comment),
	6020: uint16(sym_comma),
	6021: uint16(640),
	6022: uint16(3),
	6023: uint16(anon_sym_RBRACE),
	6024: uint16(anon_sym_DOT_DOT_DOT),
	6025: uint16(sym_name),
	6026: uint16(8),
	6027: uint16(765),
	6028: uint16(1),
	6029: uint16(anon_sym_schema),
	6030: uint16(767),
	6031: uint16(1),
	6032: uint16(anon_sym_scalar),
	6033: uint16(769),
	6034: uint16(1),
	6035: uint16(anon_sym_type),
	6036: uint16(771),
	6037: uint16(1),
	6038: uint16(anon_sym_interface),
	6039: uint16(773),
	6040: uint16(1),
	6041: uint16(anon_sym_union),
	6042: uint16(775),
	6043: uint16(1),
	6044: uint16(anon_sym_enum),
	6045: uint16(777),
	6046: uint16(1),
	6047: uint16(anon_sym_input),
	6048: uint16(3),
	6049: uint16(2),
	6050: uint16(sym_comment),
	6051: uint16(sym_comma),
	6052: uint16(4),
	6053: uint16(338),
	6054: uint16(1),
	6055: uint16(sym_operation_type),
	6056: uint16(3),
	6057: uint16(2),
	6058: uint16(sym_comment),
	6059: uint16(sym_comma),
	6060: uint16(192),
	6061: uint16(2),
	6062: uint16(sym_root_operation_type_definition),
	6063: uint16(aux_sym_schema_definition_repeat1),
	6064: uint16(23),
	6065: uint16(3),
	6066: uint16(anon_sym_query),
	6067: uint16(anon_sym_mutation),
	6068: uint16(anon_sym_subscription),
	6069: uint16(3),
	6070: uint16(568),
	6071: uint16(1),
	6072: uint16(anon_sym_DQUOTE),
	6073: uint16(3),
	6074: uint16(2),
	6075: uint16(sym_comment),
	6076: uint16(sym_comma),
	6077: uint16(566),
	6078: uint16(5),
	6079: uint16(anon_sym_RBRACE),
	6080: uint16(anon_sym_RPAREN),
	6081: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6082: uint16(anon_sym_AT),
	6083: uint16(sym_name),
	6084: uint16(6),
	6085: uint16(7),
	6086: uint16(1),
	6087: uint16(anon_sym_LBRACE),
	6088: uint16(596),
	6089: uint16(1),
	6090: uint16(anon_sym_AT),
	6091: uint16(127),
	6092: uint16(1),
	6093: uint16(sym_selection_set),
	6094: uint16(314),
	6095: uint16(1),
	6096: uint16(sym_directives),
	6097: uint16(3),
	6098: uint16(2),
	6099: uint16(sym_comment),
	6100: uint16(sym_comma),
	6101: uint16(256),
	6102: uint16(2),
	6103: uint16(sym_directive),
	6104: uint16(aux_sym_directives_repeat1),
	6105: uint16(4),
	6106: uint16(338),
	6107: uint16(1),
	6108: uint16(sym_operation_type),
	6109: uint16(3),
	6110: uint16(2),
	6111: uint16(sym_comment),
	6112: uint16(sym_comma),
	6113: uint16(193),
	6114: uint16(2),
	6115: uint16(sym_root_operation_type_definition),
	6116: uint16(aux_sym_schema_definition_repeat1),
	6117: uint16(23),
	6118: uint16(3),
	6119: uint16(anon_sym_query),
	6120: uint16(anon_sym_mutation),
	6121: uint16(anon_sym_subscription),
	6122: uint16(6),
	6123: uint16(779),
	6124: uint16(1),
	6125: uint16(anon_sym_LBRACK),
	6126: uint16(781),
	6127: uint16(1),
	6128: uint16(sym_name),
	6129: uint16(178),
	6130: uint16(1),
	6131: uint16(sym_non_null_type),
	6132: uint16(372),
	6133: uint16(1),
	6134: uint16(sym_type),
	6135: uint16(3),
	6136: uint16(2),
	6137: uint16(sym_comment),
	6138: uint16(sym_comma),
	6139: uint16(161),
	6140: uint16(2),
	6141: uint16(sym_named_type),
	6142: uint16(sym_list_type),
	6143: uint16(3),
	6144: uint16(584),
	6145: uint16(1),
	6146: uint16(anon_sym_DQUOTE),
	6147: uint16(3),
	6148: uint16(2),
	6149: uint16(sym_comment),
	6150: uint16(sym_comma),
	6151: uint16(582),
	6152: uint16(5),
	6153: uint16(anon_sym_RBRACE),
	6154: uint16(anon_sym_RPAREN),
	6155: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6156: uint16(anon_sym_AT),
	6157: uint16(sym_name),
	6158: uint16(6),
	6159: uint16(779),
	6160: uint16(1),
	6161: uint16(anon_sym_LBRACK),
	6162: uint16(781),
	6163: uint16(1),
	6164: uint16(sym_name),
	6165: uint16(178),
	6166: uint16(1),
	6167: uint16(sym_non_null_type),
	6168: uint16(383),
	6169: uint16(1),
	6170: uint16(sym_type),
	6171: uint16(3),
	6172: uint16(2),
	6173: uint16(sym_comment),
	6174: uint16(sym_comma),
	6175: uint16(161),
	6176: uint16(2),
	6177: uint16(sym_named_type),
	6178: uint16(sym_list_type),
	6179: uint16(6),
	6180: uint16(3),
	6181: uint16(1),
	6182: uint16(sym_comment),
	6183: uint16(731),
	6184: uint16(1),
	6185: uint16(anon_sym_AT),
	6186: uint16(785),
	6187: uint16(1),
	6188: uint16(sym_comma),
	6189: uint16(308),
	6190: uint16(1),
	6191: uint16(sym_directives),
	6192: uint16(783),
	6193: uint16(2),
	6194: uint16(anon_sym_RPAREN),
	6195: uint16(anon_sym_DOLLAR),
	6196: uint16(238),
	6197: uint16(2),
	6198: uint16(sym_directive),
	6199: uint16(aux_sym_directives_repeat1),
	6200: uint16(6),
	6201: uint16(7),
	6202: uint16(1),
	6203: uint16(anon_sym_LBRACE),
	6204: uint16(596),
	6205: uint16(1),
	6206: uint16(anon_sym_AT),
	6207: uint16(132),
	6208: uint16(1),
	6209: uint16(sym_selection_set),
	6210: uint16(313),
	6211: uint16(1),
	6212: uint16(sym_directives),
	6213: uint16(3),
	6214: uint16(2),
	6215: uint16(sym_comment),
	6216: uint16(sym_comma),
	6217: uint16(256),
	6218: uint16(2),
	6219: uint16(sym_directive),
	6220: uint16(aux_sym_directives_repeat1),
	6221: uint16(3),
	6222: uint16(588),
	6223: uint16(1),
	6224: uint16(anon_sym_DQUOTE),
	6225: uint16(3),
	6226: uint16(2),
	6227: uint16(sym_comment),
	6228: uint16(sym_comma),
	6229: uint16(586),
	6230: uint16(5),
	6231: uint16(anon_sym_RBRACE),
	6232: uint16(anon_sym_RPAREN),
	6233: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6234: uint16(anon_sym_AT),
	6235: uint16(sym_name),
	6236: uint16(3),
	6237: uint16(560),
	6238: uint16(1),
	6239: uint16(anon_sym_DQUOTE),
	6240: uint16(3),
	6241: uint16(2),
	6242: uint16(sym_comment),
	6243: uint16(sym_comma),
	6244: uint16(558),
	6245: uint16(5),
	6246: uint16(anon_sym_RBRACE),
	6247: uint16(anon_sym_RPAREN),
	6248: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6249: uint16(anon_sym_AT),
	6250: uint16(sym_name),
	6251: uint16(6),
	6252: uint16(779),
	6253: uint16(1),
	6254: uint16(anon_sym_LBRACK),
	6255: uint16(787),
	6256: uint16(1),
	6257: uint16(sym_name),
	6258: uint16(178),
	6259: uint16(1),
	6260: uint16(sym_non_null_type),
	6261: uint16(179),
	6262: uint16(1),
	6263: uint16(sym_type),
	6264: uint16(3),
	6265: uint16(2),
	6266: uint16(sym_comment),
	6267: uint16(sym_comma),
	6268: uint16(161),
	6269: uint16(2),
	6270: uint16(sym_named_type),
	6271: uint16(sym_list_type),
	6272: uint16(4),
	6273: uint16(338),
	6274: uint16(1),
	6275: uint16(sym_operation_type),
	6276: uint16(3),
	6277: uint16(2),
	6278: uint16(sym_comment),
	6279: uint16(sym_comma),
	6280: uint16(196),
	6281: uint16(2),
	6282: uint16(sym_root_operation_type_definition),
	6283: uint16(aux_sym_schema_definition_repeat1),
	6284: uint16(23),
	6285: uint16(3),
	6286: uint16(anon_sym_query),
	6287: uint16(anon_sym_mutation),
	6288: uint16(anon_sym_subscription),
	6289: uint16(3),
	6290: uint16(572),
	6291: uint16(1),
	6292: uint16(anon_sym_DQUOTE),
	6293: uint16(3),
	6294: uint16(2),
	6295: uint16(sym_comment),
	6296: uint16(sym_comma),
	6297: uint16(570),
	6298: uint16(5),
	6299: uint16(anon_sym_RBRACE),
	6300: uint16(anon_sym_RPAREN),
	6301: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6302: uint16(anon_sym_AT),
	6303: uint16(sym_name),
	6304: uint16(3),
	6305: uint16(564),
	6306: uint16(1),
	6307: uint16(anon_sym_DQUOTE),
	6308: uint16(3),
	6309: uint16(2),
	6310: uint16(sym_comment),
	6311: uint16(sym_comma),
	6312: uint16(562),
	6313: uint16(5),
	6314: uint16(anon_sym_RBRACE),
	6315: uint16(anon_sym_RPAREN),
	6316: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6317: uint16(anon_sym_AT),
	6318: uint16(sym_name),
	6319: uint16(6),
	6320: uint16(596),
	6321: uint16(1),
	6322: uint16(anon_sym_AT),
	6323: uint16(602),
	6324: uint16(1),
	6325: uint16(anon_sym_LBRACE),
	6326: uint16(289),
	6327: uint16(1),
	6328: uint16(sym_selection_set),
	6329: uint16(315),
	6330: uint16(1),
	6331: uint16(sym_directives),
	6332: uint16(3),
	6333: uint16(2),
	6334: uint16(sym_comment),
	6335: uint16(sym_comma),
	6336: uint16(256),
	6337: uint16(2),
	6338: uint16(sym_directive),
	6339: uint16(aux_sym_directives_repeat1),
	6340: uint16(6),
	6341: uint16(779),
	6342: uint16(1),
	6343: uint16(anon_sym_LBRACK),
	6344: uint16(787),
	6345: uint16(1),
	6346: uint16(sym_name),
	6347: uint16(148),
	6348: uint16(1),
	6349: uint16(sym_type),
	6350: uint16(178),
	6351: uint16(1),
	6352: uint16(sym_non_null_type),
	6353: uint16(3),
	6354: uint16(2),
	6355: uint16(sym_comment),
	6356: uint16(sym_comma),
	6357: uint16(161),
	6358: uint16(2),
	6359: uint16(sym_named_type),
	6360: uint16(sym_list_type),
	6361: uint16(6),
	6362: uint16(7),
	6363: uint16(1),
	6364: uint16(anon_sym_LBRACE),
	6365: uint16(596),
	6366: uint16(1),
	6367: uint16(anon_sym_AT),
	6368: uint16(124),
	6369: uint16(1),
	6370: uint16(sym_selection_set),
	6371: uint16(304),
	6372: uint16(1),
	6373: uint16(sym_directives),
	6374: uint16(3),
	6375: uint16(2),
	6376: uint16(sym_comment),
	6377: uint16(sym_comma),
	6378: uint16(256),
	6379: uint16(2),
	6380: uint16(sym_directive),
	6381: uint16(aux_sym_directives_repeat1),
	6382: uint16(6),
	6383: uint16(779),
	6384: uint16(1),
	6385: uint16(anon_sym_LBRACK),
	6386: uint16(787),
	6387: uint16(1),
	6388: uint16(sym_name),
	6389: uint16(147),
	6390: uint16(1),
	6391: uint16(sym_type),
	6392: uint16(178),
	6393: uint16(1),
	6394: uint16(sym_non_null_type),
	6395: uint16(3),
	6396: uint16(2),
	6397: uint16(sym_comment),
	6398: uint16(sym_comma),
	6399: uint16(161),
	6400: uint16(2),
	6401: uint16(sym_named_type),
	6402: uint16(sym_list_type),
	6403: uint16(3),
	6404: uint16(576),
	6405: uint16(1),
	6406: uint16(anon_sym_DQUOTE),
	6407: uint16(3),
	6408: uint16(2),
	6409: uint16(sym_comment),
	6410: uint16(sym_comma),
	6411: uint16(574),
	6412: uint16(5),
	6413: uint16(anon_sym_RBRACE),
	6414: uint16(anon_sym_RPAREN),
	6415: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6416: uint16(anon_sym_AT),
	6417: uint16(sym_name),
	6418: uint16(6),
	6419: uint16(779),
	6420: uint16(1),
	6421: uint16(anon_sym_LBRACK),
	6422: uint16(787),
	6423: uint16(1),
	6424: uint16(sym_name),
	6425: uint16(178),
	6426: uint16(1),
	6427: uint16(sym_non_null_type),
	6428: uint16(186),
	6429: uint16(1),
	6430: uint16(sym_type),
	6431: uint16(3),
	6432: uint16(2),
	6433: uint16(sym_comment),
	6434: uint16(sym_comma),
	6435: uint16(161),
	6436: uint16(2),
	6437: uint16(sym_named_type),
	6438: uint16(sym_list_type),
	6439: uint16(6),
	6440: uint16(779),
	6441: uint16(1),
	6442: uint16(anon_sym_LBRACK),
	6443: uint16(787),
	6444: uint16(1),
	6445: uint16(sym_name),
	6446: uint16(171),
	6447: uint16(1),
	6448: uint16(sym_type),
	6449: uint16(178),
	6450: uint16(1),
	6451: uint16(sym_non_null_type),
	6452: uint16(3),
	6453: uint16(2),
	6454: uint16(sym_comment),
	6455: uint16(sym_comma),
	6456: uint16(161),
	6457: uint16(2),
	6458: uint16(sym_named_type),
	6459: uint16(sym_list_type),
	6460: uint16(6),
	6461: uint16(789),
	6462: uint16(1),
	6463: uint16(anon_sym_LBRACK),
	6464: uint16(791),
	6465: uint16(1),
	6466: uint16(sym_name),
	6467: uint16(184),
	6468: uint16(1),
	6469: uint16(sym_type),
	6470: uint16(260),
	6471: uint16(1),
	6472: uint16(sym_non_null_type),
	6473: uint16(3),
	6474: uint16(2),
	6475: uint16(sym_comment),
	6476: uint16(sym_comma),
	6477: uint16(249),
	6478: uint16(2),
	6479: uint16(sym_named_type),
	6480: uint16(sym_list_type),
	6481: uint16(3),
	6482: uint16(795),
	6483: uint16(1),
	6484: uint16(anon_sym_DQUOTE),
	6485: uint16(3),
	6486: uint16(2),
	6487: uint16(sym_comment),
	6488: uint16(sym_comma),
	6489: uint16(793),
	6490: uint16(5),
	6491: uint16(anon_sym_RBRACE),
	6492: uint16(anon_sym_RPAREN),
	6493: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6494: uint16(anon_sym_AT),
	6495: uint16(sym_name),
	6496: uint16(3),
	6497: uint16(556),
	6498: uint16(1),
	6499: uint16(anon_sym_DQUOTE),
	6500: uint16(3),
	6501: uint16(2),
	6502: uint16(sym_comment),
	6503: uint16(sym_comma),
	6504: uint16(554),
	6505: uint16(5),
	6506: uint16(anon_sym_RBRACE),
	6507: uint16(anon_sym_RPAREN),
	6508: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6509: uint16(anon_sym_AT),
	6510: uint16(sym_name),
	6511: uint16(2),
	6512: uint16(3),
	6513: uint16(1),
	6514: uint16(sym_comment),
	6515: uint16(562),
	6516: uint16(6),
	6517: uint16(anon_sym_RBRACE),
	6518: uint16(anon_sym_RPAREN),
	6519: uint16(anon_sym_DOLLAR),
	6520: uint16(anon_sym_AT),
	6521: uint16(sym_name),
	6522: uint16(sym_comma),
	6523: uint16(4),
	6524: uint16(3),
	6525: uint16(1),
	6526: uint16(sym_comment),
	6527: uint16(797),
	6528: uint16(1),
	6529: uint16(anon_sym_AT),
	6530: uint16(226),
	6531: uint16(2),
	6532: uint16(sym_directive),
	6533: uint16(aux_sym_directives_repeat1),
	6534: uint16(193),
	6535: uint16(3),
	6536: uint16(anon_sym_RPAREN),
	6537: uint16(anon_sym_DOLLAR),
	6538: uint16(sym_comma),
	6539: uint16(5),
	6540: uint16(596),
	6541: uint16(1),
	6542: uint16(anon_sym_AT),
	6543: uint16(800),
	6544: uint16(1),
	6545: uint16(anon_sym_LBRACE),
	6546: uint16(360),
	6547: uint16(1),
	6548: uint16(sym_directives),
	6549: uint16(3),
	6550: uint16(2),
	6551: uint16(sym_comment),
	6552: uint16(sym_comma),
	6553: uint16(256),
	6554: uint16(2),
	6555: uint16(sym_directive),
	6556: uint16(aux_sym_directives_repeat1),
	6557: uint16(5),
	6558: uint16(308),
	6559: uint16(1),
	6560: uint16(anon_sym_DOLLAR),
	6561: uint16(802),
	6562: uint16(1),
	6563: uint16(anon_sym_RPAREN),
	6564: uint16(353),
	6565: uint16(1),
	6566: uint16(sym_variable),
	6567: uint16(3),
	6568: uint16(2),
	6569: uint16(sym_comment),
	6570: uint16(sym_comma),
	6571: uint16(241),
	6572: uint16(2),
	6573: uint16(sym_variable_definition),
	6574: uint16(aux_sym_variable_definitions_repeat1),
	6575: uint16(4),
	6576: uint16(602),
	6577: uint16(1),
	6578: uint16(anon_sym_LBRACE),
	6579: uint16(288),
	6580: uint16(1),
	6581: uint16(sym_selection_set),
	6582: uint16(3),
	6583: uint16(2),
	6584: uint16(sym_comment),
	6585: uint16(sym_comma),
	6586: uint16(640),
	6587: uint16(3),
	6588: uint16(anon_sym_RBRACE),
	6589: uint16(anon_sym_DOT_DOT_DOT),
	6590: uint16(sym_name),
	6591: uint16(3),
	6592: uint16(664),
	6593: uint16(1),
	6594: uint16(anon_sym_DQUOTE),
	6595: uint16(3),
	6596: uint16(2),
	6597: uint16(sym_comment),
	6598: uint16(sym_comma),
	6599: uint16(662),
	6600: uint16(4),
	6601: uint16(anon_sym_RBRACE),
	6602: uint16(anon_sym_RPAREN),
	6603: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6604: uint16(sym_name),
	6605: uint16(4),
	6606: uint16(338),
	6607: uint16(1),
	6608: uint16(sym_operation_type),
	6609: uint16(339),
	6610: uint16(1),
	6611: uint16(sym_root_operation_type_definition),
	6612: uint16(3),
	6613: uint16(2),
	6614: uint16(sym_comment),
	6615: uint16(sym_comma),
	6616: uint16(23),
	6617: uint16(3),
	6618: uint16(anon_sym_query),
	6619: uint16(anon_sym_mutation),
	6620: uint16(anon_sym_subscription),
	6621: uint16(2),
	6622: uint16(3),
	6623: uint16(1),
	6624: uint16(sym_comment),
	6625: uint16(554),
	6626: uint16(6),
	6627: uint16(anon_sym_RBRACE),
	6628: uint16(anon_sym_RPAREN),
	6629: uint16(anon_sym_DOLLAR),
	6630: uint16(anon_sym_AT),
	6631: uint16(sym_name),
	6632: uint16(sym_comma),
	6633: uint16(2),
	6634: uint16(3),
	6635: uint16(1),
	6636: uint16(sym_comment),
	6637: uint16(558),
	6638: uint16(6),
	6639: uint16(anon_sym_RBRACE),
	6640: uint16(anon_sym_RPAREN),
	6641: uint16(anon_sym_DOLLAR),
	6642: uint16(anon_sym_AT),
	6643: uint16(sym_name),
	6644: uint16(sym_comma),
	6645: uint16(2),
	6646: uint16(3),
	6647: uint16(1),
	6648: uint16(sym_comment),
	6649: uint16(566),
	6650: uint16(6),
	6651: uint16(anon_sym_RBRACE),
	6652: uint16(anon_sym_RPAREN),
	6653: uint16(anon_sym_DOLLAR),
	6654: uint16(anon_sym_AT),
	6655: uint16(sym_name),
	6656: uint16(sym_comma),
	6657: uint16(2),
	6658: uint16(3),
	6659: uint16(1),
	6660: uint16(sym_comment),
	6661: uint16(570),
	6662: uint16(6),
	6663: uint16(anon_sym_RBRACE),
	6664: uint16(anon_sym_RPAREN),
	6665: uint16(anon_sym_DOLLAR),
	6666: uint16(anon_sym_AT),
	6667: uint16(sym_name),
	6668: uint16(sym_comma),
	6669: uint16(3),
	6670: uint16(806),
	6671: uint16(1),
	6672: uint16(anon_sym_DQUOTE),
	6673: uint16(3),
	6674: uint16(2),
	6675: uint16(sym_comment),
	6676: uint16(sym_comma),
	6677: uint16(804),
	6678: uint16(4),
	6679: uint16(anon_sym_RBRACE),
	6680: uint16(anon_sym_RPAREN),
	6681: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6682: uint16(sym_name),
	6683: uint16(5),
	6684: uint16(596),
	6685: uint16(1),
	6686: uint16(anon_sym_AT),
	6687: uint16(808),
	6688: uint16(1),
	6689: uint16(anon_sym_LBRACE),
	6690: uint16(351),
	6691: uint16(1),
	6692: uint16(sym_directives),
	6693: uint16(3),
	6694: uint16(2),
	6695: uint16(sym_comment),
	6696: uint16(sym_comma),
	6697: uint16(256),
	6698: uint16(2),
	6699: uint16(sym_directive),
	6700: uint16(aux_sym_directives_repeat1),
	6701: uint16(4),
	6702: uint16(3),
	6703: uint16(1),
	6704: uint16(sym_comment),
	6705: uint16(731),
	6706: uint16(1),
	6707: uint16(anon_sym_AT),
	6708: uint16(226),
	6709: uint16(2),
	6710: uint16(sym_directive),
	6711: uint16(aux_sym_directives_repeat1),
	6712: uint16(204),
	6713: uint16(3),
	6714: uint16(anon_sym_RPAREN),
	6715: uint16(anon_sym_DOLLAR),
	6716: uint16(sym_comma),
	6717: uint16(4),
	6718: uint16(3),
	6719: uint16(1),
	6720: uint16(sym_comment),
	6721: uint16(810),
	6722: uint16(1),
	6723: uint16(anon_sym_LPAREN),
	6724: uint16(292),
	6725: uint16(1),
	6726: uint16(sym_arguments),
	6727: uint16(183),
	6728: uint16(4),
	6729: uint16(anon_sym_RPAREN),
	6730: uint16(anon_sym_DOLLAR),
	6731: uint16(anon_sym_AT),
	6732: uint16(sym_comma),
	6733: uint16(2),
	6734: uint16(3),
	6735: uint16(1),
	6736: uint16(sym_comment),
	6737: uint16(648),
	6738: uint16(6),
	6739: uint16(anon_sym_RPAREN),
	6740: uint16(anon_sym_EQ),
	6741: uint16(anon_sym_DOLLAR),
	6742: uint16(anon_sym_AT),
	6743: uint16(anon_sym_BANG),
	6744: uint16(sym_comma),
	6745: uint16(5),
	6746: uint16(812),
	6747: uint16(1),
	6748: uint16(anon_sym_RPAREN),
	6749: uint16(814),
	6750: uint16(1),
	6751: uint16(anon_sym_DOLLAR),
	6752: uint16(353),
	6753: uint16(1),
	6754: uint16(sym_variable),
	6755: uint16(3),
	6756: uint16(2),
	6757: uint16(sym_comment),
	6758: uint16(sym_comma),
	6759: uint16(241),
	6760: uint16(2),
	6761: uint16(sym_variable_definition),
	6762: uint16(aux_sym_variable_definitions_repeat1),
	6763: uint16(3),
	6764: uint16(594),
	6765: uint16(1),
	6766: uint16(anon_sym_DQUOTE),
	6767: uint16(3),
	6768: uint16(2),
	6769: uint16(sym_comment),
	6770: uint16(sym_comma),
	6771: uint16(590),
	6772: uint16(4),
	6773: uint16(anon_sym_RBRACE),
	6774: uint16(anon_sym_RPAREN),
	6775: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6776: uint16(sym_name),
	6777: uint16(2),
	6778: uint16(3),
	6779: uint16(1),
	6780: uint16(sym_comment),
	6781: uint16(125),
	6782: uint16(6),
	6783: uint16(anon_sym_RPAREN),
	6784: uint16(anon_sym_EQ),
	6785: uint16(anon_sym_DOLLAR),
	6786: uint16(anon_sym_AT),
	6787: uint16(anon_sym_BANG),
	6788: uint16(sym_comma),
	6789: uint16(2),
	6790: uint16(3),
	6791: uint16(1),
	6792: uint16(sym_comment),
	6793: uint16(578),
	6794: uint16(6),
	6795: uint16(anon_sym_RBRACE),
	6796: uint16(anon_sym_RPAREN),
	6797: uint16(anon_sym_DOLLAR),
	6798: uint16(anon_sym_AT),
	6799: uint16(sym_name),
	6800: uint16(sym_comma),
	6801: uint16(2),
	6802: uint16(3),
	6803: uint16(1),
	6804: uint16(sym_comment),
	6805: uint16(586),
	6806: uint16(6),
	6807: uint16(anon_sym_RBRACE),
	6808: uint16(anon_sym_RPAREN),
	6809: uint16(anon_sym_DOLLAR),
	6810: uint16(anon_sym_AT),
	6811: uint16(sym_name),
	6812: uint16(sym_comma),
	6813: uint16(5),
	6814: uint16(596),
	6815: uint16(1),
	6816: uint16(anon_sym_AT),
	6817: uint16(817),
	6818: uint16(1),
	6819: uint16(anon_sym_LBRACE),
	6820: uint16(328),
	6821: uint16(1),
	6822: uint16(sym_directives),
	6823: uint16(3),
	6824: uint16(2),
	6825: uint16(sym_comment),
	6826: uint16(sym_comma),
	6827: uint16(256),
	6828: uint16(2),
	6829: uint16(sym_directive),
	6830: uint16(aux_sym_directives_repeat1),
	6831: uint16(4),
	6832: uint16(602),
	6833: uint16(1),
	6834: uint16(anon_sym_LBRACE),
	6835: uint16(293),
	6836: uint16(1),
	6837: uint16(sym_selection_set),
	6838: uint16(3),
	6839: uint16(2),
	6840: uint16(sym_comment),
	6841: uint16(sym_comma),
	6842: uint16(819),
	6843: uint16(3),
	6844: uint16(anon_sym_RBRACE),
	6845: uint16(anon_sym_DOT_DOT_DOT),
	6846: uint16(sym_name),
	6847: uint16(4),
	6848: uint16(338),
	6849: uint16(1),
	6850: uint16(sym_operation_type),
	6851: uint16(379),
	6852: uint16(1),
	6853: uint16(sym_root_operation_type_definition),
	6854: uint16(3),
	6855: uint16(2),
	6856: uint16(sym_comment),
	6857: uint16(sym_comma),
	6858: uint16(23),
	6859: uint16(3),
	6860: uint16(anon_sym_query),
	6861: uint16(anon_sym_mutation),
	6862: uint16(anon_sym_subscription),
	6863: uint16(3),
	6864: uint16(3),
	6865: uint16(1),
	6866: uint16(sym_comment),
	6867: uint16(821),
	6868: uint16(1),
	6869: uint16(anon_sym_BANG),
	6870: uint16(652),
	6871: uint16(5),
	6872: uint16(anon_sym_RPAREN),
	6873: uint16(anon_sym_EQ),
	6874: uint16(anon_sym_DOLLAR),
	6875: uint16(anon_sym_AT),
	6876: uint16(sym_comma),
	6877: uint16(4),
	6878: uint16(602),
	6879: uint16(1),
	6880: uint16(anon_sym_LBRACE),
	6881: uint16(291),
	6882: uint16(1),
	6883: uint16(sym_selection_set),
	6884: uint16(3),
	6885: uint16(2),
	6886: uint16(sym_comment),
	6887: uint16(sym_comma),
	6888: uint16(759),
	6889: uint16(3),
	6890: uint16(anon_sym_RBRACE),
	6891: uint16(anon_sym_DOT_DOT_DOT),
	6892: uint16(sym_name),
	6893: uint16(2),
	6894: uint16(3),
	6895: uint16(1),
	6896: uint16(sym_comment),
	6897: uint16(582),
	6898: uint16(6),
	6899: uint16(anon_sym_RBRACE),
	6900: uint16(anon_sym_RPAREN),
	6901: uint16(anon_sym_DOLLAR),
	6902: uint16(anon_sym_AT),
	6903: uint16(sym_name),
	6904: uint16(sym_comma),
	6905: uint16(2),
	6906: uint16(3),
	6907: uint16(2),
	6908: uint16(sym_comment),
	6909: uint16(sym_comma),
	6910: uint16(823),
	6911: uint16(5),
	6912: uint16(anon_sym_LBRACE),
	6913: uint16(anon_sym_COLON),
	6914: uint16(anon_sym_LPAREN),
	6915: uint16(anon_sym_AT),
	6916: uint16(sym_name),
	6917: uint16(2),
	6918: uint16(3),
	6919: uint16(1),
	6920: uint16(sym_comment),
	6921: uint16(574),
	6922: uint16(6),
	6923: uint16(anon_sym_RBRACE),
	6924: uint16(anon_sym_RPAREN),
	6925: uint16(anon_sym_DOLLAR),
	6926: uint16(anon_sym_AT),
	6927: uint16(sym_name),
	6928: uint16(sym_comma),
	6929: uint16(3),
	6930: uint16(827),
	6931: uint16(1),
	6932: uint16(anon_sym_DQUOTE),
	6933: uint16(3),
	6934: uint16(2),
	6935: uint16(sym_comment),
	6936: uint16(sym_comma),
	6937: uint16(825),
	6938: uint16(3),
	6939: uint16(anon_sym_RBRACE),
	6940: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	6941: uint16(sym_name),
	6942: uint16(2),
	6943: uint16(3),
	6944: uint16(2),
	6945: uint16(sym_comment),
	6946: uint16(sym_comma),
	6947: uint16(829),
	6948: uint16(4),
	6949: uint16(anon_sym_RBRACE),
	6950: uint16(anon_sym_query),
	6951: uint16(anon_sym_mutation),
	6952: uint16(anon_sym_subscription),
	6953: uint16(4),
	6954: uint16(204),
	6955: uint16(1),
	6956: uint16(anon_sym_LBRACE),
	6957: uint16(596),
	6958: uint16(1),
	6959: uint16(anon_sym_AT),
	6960: uint16(3),
	6961: uint16(2),
	6962: uint16(sym_comment),
	6963: uint16(sym_comma),
	6964: uint16(269),
	6965: uint16(2),
	6966: uint16(sym_directive),
	6967: uint16(aux_sym_directives_repeat1),
	6968: uint16(4),
	6969: uint16(831),
	6970: uint16(1),
	6971: uint16(anon_sym_RBRACE),
	6972: uint16(833),
	6973: uint16(1),
	6974: uint16(sym_name),
	6975: uint16(3),
	6976: uint16(2),
	6977: uint16(sym_comment),
	6978: uint16(sym_comma),
	6979: uint16(277),
	6980: uint16(2),
	6981: uint16(sym_object_field),
	6982: uint16(aux_sym_object_value_repeat1),
	6983: uint16(4),
	6984: uint16(833),
	6985: uint16(1),
	6986: uint16(sym_name),
	6987: uint16(835),
	6988: uint16(1),
	6989: uint16(anon_sym_RBRACE),
	6990: uint16(3),
	6991: uint16(2),
	6992: uint16(sym_comment),
	6993: uint16(sym_comma),
	6994: uint16(257),
	6995: uint16(2),
	6996: uint16(sym_object_field),
	6997: uint16(aux_sym_object_value_repeat1),
	6998: uint16(4),
	6999: uint16(837),
	7000: uint16(1),
	7001: uint16(anon_sym_RPAREN),
	7002: uint16(839),
	7003: uint16(1),
	7004: uint16(sym_name),
	7005: uint16(3),
	7006: uint16(2),
	7007: uint16(sym_comment),
	7008: uint16(sym_comma),
	7009: uint16(276),
	7010: uint16(2),
	7011: uint16(sym_argument),
	7012: uint16(aux_sym_arguments_repeat1),
	7013: uint16(2),
	7014: uint16(3),
	7015: uint16(1),
	7016: uint16(sym_comment),
	7017: uint16(652),
	7018: uint16(5),
	7019: uint16(anon_sym_RPAREN),
	7020: uint16(anon_sym_EQ),
	7021: uint16(anon_sym_DOLLAR),
	7022: uint16(anon_sym_AT),
	7023: uint16(sym_comma),
	7024: uint16(4),
	7025: uint16(833),
	7026: uint16(1),
	7027: uint16(sym_name),
	7028: uint16(841),
	7029: uint16(1),
	7030: uint16(anon_sym_RBRACE),
	7031: uint16(3),
	7032: uint16(2),
	7033: uint16(sym_comment),
	7034: uint16(sym_comma),
	7035: uint16(277),
	7036: uint16(2),
	7037: uint16(sym_object_field),
	7038: uint16(aux_sym_object_value_repeat1),
	7039: uint16(4),
	7040: uint16(833),
	7041: uint16(1),
	7042: uint16(sym_name),
	7043: uint16(843),
	7044: uint16(1),
	7045: uint16(anon_sym_RBRACE),
	7046: uint16(3),
	7047: uint16(2),
	7048: uint16(sym_comment),
	7049: uint16(sym_comma),
	7050: uint16(273),
	7051: uint16(2),
	7052: uint16(sym_object_field),
	7053: uint16(aux_sym_object_value_repeat1),
	7054: uint16(4),
	7055: uint16(833),
	7056: uint16(1),
	7057: uint16(sym_name),
	7058: uint16(845),
	7059: uint16(1),
	7060: uint16(anon_sym_RBRACE),
	7061: uint16(3),
	7062: uint16(2),
	7063: uint16(sym_comment),
	7064: uint16(sym_comma),
	7065: uint16(261),
	7066: uint16(2),
	7067: uint16(sym_object_field),
	7068: uint16(aux_sym_object_value_repeat1),
	7069: uint16(4),
	7070: uint16(839),
	7071: uint16(1),
	7072: uint16(sym_name),
	7073: uint16(847),
	7074: uint16(1),
	7075: uint16(anon_sym_RPAREN),
	7076: uint16(3),
	7077: uint16(2),
	7078: uint16(sym_comment),
	7079: uint16(sym_comma),
	7080: uint16(276),
	7081: uint16(2),
	7082: uint16(sym_argument),
	7083: uint16(aux_sym_arguments_repeat1),
	7084: uint16(4),
	7085: uint16(87),
	7086: uint16(1),
	7087: uint16(anon_sym_AT),
	7088: uint16(120),
	7089: uint16(1),
	7090: uint16(sym_directives),
	7091: uint16(3),
	7092: uint16(2),
	7093: uint16(sym_comment),
	7094: uint16(sym_comma),
	7095: uint16(36),
	7096: uint16(2),
	7097: uint16(sym_directive),
	7098: uint16(aux_sym_directives_repeat1),
	7099: uint16(5),
	7100: uint16(849),
	7101: uint16(1),
	7102: uint16(anon_sym_LPAREN),
	7103: uint16(851),
	7104: uint16(1),
	7105: uint16(anon_sym_on),
	7106: uint16(853),
	7107: uint16(1),
	7108: uint16(anon_sym_repeatable),
	7109: uint16(305),
	7110: uint16(1),
	7111: uint16(sym_arguments_definition),
	7112: uint16(3),
	7113: uint16(2),
	7114: uint16(sym_comment),
	7115: uint16(sym_comma),
	7116: uint16(4),
	7117: uint16(839),
	7118: uint16(1),
	7119: uint16(sym_name),
	7120: uint16(855),
	7121: uint16(1),
	7122: uint16(anon_sym_RPAREN),
	7123: uint16(3),
	7124: uint16(2),
	7125: uint16(sym_comment),
	7126: uint16(sym_comma),
	7127: uint16(276),
	7128: uint16(2),
	7129: uint16(sym_argument),
	7130: uint16(aux_sym_arguments_repeat1),
	7131: uint16(3),
	7132: uint16(695),
	7133: uint16(1),
	7134: uint16(anon_sym_DQUOTE),
	7135: uint16(3),
	7136: uint16(2),
	7137: uint16(sym_comment),
	7138: uint16(sym_comma),
	7139: uint16(693),
	7140: uint16(3),
	7141: uint16(anon_sym_RBRACE),
	7142: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7143: uint16(sym_name),
	7144: uint16(4),
	7145: uint16(193),
	7146: uint16(1),
	7147: uint16(anon_sym_LBRACE),
	7148: uint16(625),
	7149: uint16(1),
	7150: uint16(anon_sym_AT),
	7151: uint16(3),
	7152: uint16(2),
	7153: uint16(sym_comment),
	7154: uint16(sym_comma),
	7155: uint16(269),
	7156: uint16(2),
	7157: uint16(sym_directive),
	7158: uint16(aux_sym_directives_repeat1),
	7159: uint16(2),
	7160: uint16(3),
	7161: uint16(2),
	7162: uint16(sym_comment),
	7163: uint16(sym_comma),
	7164: uint16(857),
	7165: uint16(4),
	7166: uint16(anon_sym_RBRACE),
	7167: uint16(anon_sym_DOT_DOT_DOT),
	7168: uint16(anon_sym_AT),
	7169: uint16(sym_name),
	7170: uint16(3),
	7171: uint16(861),
	7172: uint16(1),
	7173: uint16(anon_sym_DQUOTE),
	7174: uint16(3),
	7175: uint16(2),
	7176: uint16(sym_comment),
	7177: uint16(sym_comma),
	7178: uint16(859),
	7179: uint16(3),
	7180: uint16(anon_sym_RBRACE),
	7181: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7182: uint16(sym_name),
	7183: uint16(3),
	7184: uint16(707),
	7185: uint16(1),
	7186: uint16(anon_sym_DQUOTE),
	7187: uint16(3),
	7188: uint16(2),
	7189: uint16(sym_comment),
	7190: uint16(sym_comma),
	7191: uint16(705),
	7192: uint16(3),
	7193: uint16(anon_sym_RBRACE),
	7194: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7195: uint16(sym_name),
	7196: uint16(4),
	7197: uint16(833),
	7198: uint16(1),
	7199: uint16(sym_name),
	7200: uint16(863),
	7201: uint16(1),
	7202: uint16(anon_sym_RBRACE),
	7203: uint16(3),
	7204: uint16(2),
	7205: uint16(sym_comment),
	7206: uint16(sym_comma),
	7207: uint16(277),
	7208: uint16(2),
	7209: uint16(sym_object_field),
	7210: uint16(aux_sym_object_value_repeat1),
	7211: uint16(5),
	7212: uint16(849),
	7213: uint16(1),
	7214: uint16(anon_sym_LPAREN),
	7215: uint16(865),
	7216: uint16(1),
	7217: uint16(anon_sym_on),
	7218: uint16(867),
	7219: uint16(1),
	7220: uint16(anon_sym_repeatable),
	7221: uint16(322),
	7222: uint16(1),
	7223: uint16(sym_arguments_definition),
	7224: uint16(3),
	7225: uint16(2),
	7226: uint16(sym_comment),
	7227: uint16(sym_comma),
	7228: uint16(2),
	7229: uint16(3),
	7230: uint16(1),
	7231: uint16(sym_comment),
	7232: uint16(701),
	7233: uint16(5),
	7234: uint16(anon_sym_RPAREN),
	7235: uint16(anon_sym_EQ),
	7236: uint16(anon_sym_DOLLAR),
	7237: uint16(anon_sym_AT),
	7238: uint16(sym_comma),
	7239: uint16(4),
	7240: uint16(869),
	7241: uint16(1),
	7242: uint16(anon_sym_RPAREN),
	7243: uint16(871),
	7244: uint16(1),
	7245: uint16(sym_name),
	7246: uint16(3),
	7247: uint16(2),
	7248: uint16(sym_comment),
	7249: uint16(sym_comma),
	7250: uint16(276),
	7251: uint16(2),
	7252: uint16(sym_argument),
	7253: uint16(aux_sym_arguments_repeat1),
	7254: uint16(4),
	7255: uint16(874),
	7256: uint16(1),
	7257: uint16(anon_sym_RBRACE),
	7258: uint16(876),
	7259: uint16(1),
	7260: uint16(sym_name),
	7261: uint16(3),
	7262: uint16(2),
	7263: uint16(sym_comment),
	7264: uint16(sym_comma),
	7265: uint16(277),
	7266: uint16(2),
	7267: uint16(sym_object_field),
	7268: uint16(aux_sym_object_value_repeat1),
	7269: uint16(3),
	7270: uint16(691),
	7271: uint16(1),
	7272: uint16(anon_sym_DQUOTE),
	7273: uint16(3),
	7274: uint16(2),
	7275: uint16(sym_comment),
	7276: uint16(sym_comma),
	7277: uint16(689),
	7278: uint16(3),
	7279: uint16(anon_sym_RBRACE),
	7280: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	7281: uint16(sym_name),
	7282: uint16(4),
	7283: uint16(308),
	7284: uint16(1),
	7285: uint16(anon_sym_DOLLAR),
	7286: uint16(353),
	7287: uint16(1),
	7288: uint16(sym_variable),
	7289: uint16(3),
	7290: uint16(2),
	7291: uint16(sym_comment),
	7292: uint16(sym_comma),
	7293: uint16(228),
	7294: uint16(2),
	7295: uint16(sym_variable_definition),
	7296: uint16(aux_sym_variable_definitions_repeat1),
	7297: uint16(2),
	7298: uint16(3),
	7299: uint16(2),
	7300: uint16(sym_comment),
	7301: uint16(sym_comma),
	7302: uint16(640),
	7303: uint16(3),
	7304: uint16(anon_sym_RBRACE),
	7305: uint16(anon_sym_DOT_DOT_DOT),
	7306: uint16(sym_name),
	7307: uint16(2),
	7308: uint16(3),
	7309: uint16(1),
	7310: uint16(sym_comment),
	7311: uint16(294),
	7312: uint16(4),
	7313: uint16(anon_sym_RPAREN),
	7314: uint16(anon_sym_DOLLAR),
	7315: uint16(anon_sym_AT),
	7316: uint16(sym_comma),
	7317: uint16(3),
	7318: uint16(839),
	7319: uint16(1),
	7320: uint16(sym_name),
	7321: uint16(3),
	7322: uint16(2),
	7323: uint16(sym_comment),
	7324: uint16(sym_comma),
	7325: uint16(259),
	7326: uint16(2),
	7327: uint16(sym_argument),
	7328: uint16(aux_sym_arguments_repeat1),
	7329: uint16(2),
	7330: uint16(3),
	7331: uint16(2),
	7332: uint16(sym_comment),
	7333: uint16(sym_comma),
	7334: uint16(879),
	7335: uint16(3),
	7336: uint16(anon_sym_COLON),
	7337: uint16(anon_sym_on),
	7338: uint16(anon_sym_repeatable),
	7339: uint16(4),
	7340: uint16(849),
	7341: uint16(1),
	7342: uint16(anon_sym_LPAREN),
	7343: uint16(881),
	7344: uint16(1),
	7345: uint16(anon_sym_COLON),
	7346: uint16(337),
	7347: uint16(1),
	7348: uint16(sym_arguments_definition),
	7349: uint16(3),
	7350: uint16(2),
	7351: uint16(sym_comment),
	7352: uint16(sym_comma),
	7353: uint16(3),
	7354: uint16(839),
	7355: uint16(1),
	7356: uint16(sym_name),
	7357: uint16(3),
	7358: uint16(2),
	7359: uint16(sym_comment),
	7360: uint16(sym_comma),
	7361: uint16(264),
	7362: uint16(2),
	7363: uint16(sym_argument),
	7364: uint16(aux_sym_arguments_repeat1),
	7365: uint16(2),
	7366: uint16(3),
	7367: uint16(2),
	7368: uint16(sym_comment),
	7369: uint16(sym_comma),
	7370: uint16(883),
	7371: uint16(3),
	7372: uint16(anon_sym_RBRACE),
	7373: uint16(anon_sym_DOT_DOT_DOT),
	7374: uint16(sym_name),
	7375: uint16(2),
	7376: uint16(3),
	7377: uint16(2),
	7378: uint16(sym_comment),
	7379: uint16(sym_comma),
	7380: uint16(466),
	7381: uint16(3),
	7382: uint16(anon_sym_RBRACE),
	7383: uint16(anon_sym_DOT_DOT_DOT),
	7384: uint16(sym_name),
	7385: uint16(2),
	7386: uint16(3),
	7387: uint16(2),
	7388: uint16(sym_comment),
	7389: uint16(sym_comma),
	7390: uint16(759),
	7391: uint16(3),
	7392: uint16(anon_sym_RBRACE),
	7393: uint16(anon_sym_DOT_DOT_DOT),
	7394: uint16(sym_name),
	7395: uint16(2),
	7396: uint16(3),
	7397: uint16(2),
	7398: uint16(sym_comment),
	7399: uint16(sym_comma),
	7400: uint16(885),
	7401: uint16(3),
	7402: uint16(anon_sym_RBRACE),
	7403: uint16(anon_sym_DOT_DOT_DOT),
	7404: uint16(sym_name),
	7405: uint16(2),
	7406: uint16(3),
	7407: uint16(2),
	7408: uint16(sym_comment),
	7409: uint16(sym_comma),
	7410: uint16(887),
	7411: uint16(3),
	7412: uint16(anon_sym_RBRACE),
	7413: uint16(anon_sym_DOT_DOT_DOT),
	7414: uint16(sym_name),
	7415: uint16(2),
	7416: uint16(3),
	7417: uint16(2),
	7418: uint16(sym_comment),
	7419: uint16(sym_comma),
	7420: uint16(819),
	7421: uint16(3),
	7422: uint16(anon_sym_RBRACE),
	7423: uint16(anon_sym_DOT_DOT_DOT),
	7424: uint16(sym_name),
	7425: uint16(2),
	7426: uint16(3),
	7427: uint16(1),
	7428: uint16(sym_comment),
	7429: uint16(298),
	7430: uint16(4),
	7431: uint16(anon_sym_RPAREN),
	7432: uint16(anon_sym_DOLLAR),
	7433: uint16(anon_sym_AT),
	7434: uint16(sym_comma),
	7435: uint16(2),
	7436: uint16(3),
	7437: uint16(2),
	7438: uint16(sym_comment),
	7439: uint16(sym_comma),
	7440: uint16(889),
	7441: uint16(3),
	7442: uint16(anon_sym_RBRACE),
	7443: uint16(anon_sym_DOT_DOT_DOT),
	7444: uint16(sym_name),
	7445: uint16(2),
	7446: uint16(3),
	7447: uint16(2),
	7448: uint16(sym_comment),
	7449: uint16(sym_comma),
	7450: uint16(891),
	7451: uint16(3),
	7452: uint16(anon_sym_RBRACE),
	7453: uint16(anon_sym_DOT_DOT_DOT),
	7454: uint16(sym_name),
	7455: uint16(4),
	7456: uint16(849),
	7457: uint16(1),
	7458: uint16(anon_sym_LPAREN),
	7459: uint16(893),
	7460: uint16(1),
	7461: uint16(anon_sym_COLON),
	7462: uint16(341),
	7463: uint16(1),
	7464: uint16(sym_arguments_definition),
	7465: uint16(3),
	7466: uint16(2),
	7467: uint16(sym_comment),
	7468: uint16(sym_comma),
	7469: uint16(3),
	7470: uint16(839),
	7471: uint16(1),
	7472: uint16(sym_name),
	7473: uint16(3),
	7474: uint16(2),
	7475: uint16(sym_comment),
	7476: uint16(sym_comma),
	7477: uint16(267),
	7478: uint16(2),
	7479: uint16(sym_argument),
	7480: uint16(aux_sym_arguments_repeat1),
	7481: uint16(2),
	7482: uint16(3),
	7483: uint16(2),
	7484: uint16(sym_comment),
	7485: uint16(sym_comma),
	7486: uint16(895),
	7487: uint16(3),
	7488: uint16(anon_sym_RBRACE),
	7489: uint16(anon_sym_DOT_DOT_DOT),
	7490: uint16(sym_name),
	7491: uint16(4),
	7492: uint16(131),
	7493: uint16(1),
	7494: uint16(anon_sym_AMP),
	7495: uint16(781),
	7496: uint16(1),
	7497: uint16(sym_name),
	7498: uint16(51),
	7499: uint16(1),
	7500: uint16(sym_named_type),
	7501: uint16(3),
	7502: uint16(2),
	7503: uint16(sym_comment),
	7504: uint16(sym_comma),
	7505: uint16(4),
	7506: uint16(324),
	7507: uint16(1),
	7508: uint16(anon_sym_PIPE),
	7509: uint16(781),
	7510: uint16(1),
	7511: uint16(sym_name),
	7512: uint16(57),
	7513: uint16(1),
	7514: uint16(sym_named_type),
	7515: uint16(3),
	7516: uint16(2),
	7517: uint16(sym_comment),
	7518: uint16(sym_comma),
	7519: uint16(2),
	7520: uint16(3),
	7521: uint16(1),
	7522: uint16(sym_comment),
	7523: uint16(793),
	7524: uint16(4),
	7525: uint16(anon_sym_RPAREN),
	7526: uint16(anon_sym_DOLLAR),
	7527: uint16(anon_sym_AT),
	7528: uint16(sym_comma),
	7529: uint16(3),
	7530: uint16(3),
	7531: uint16(1),
	7532: uint16(sym_comment),
	7533: uint16(785),
	7534: uint16(1),
	7535: uint16(sym_comma),
	7536: uint16(783),
	7537: uint16(2),
	7538: uint16(anon_sym_RPAREN),
	7539: uint16(anon_sym_DOLLAR),
	7540: uint16(2),
	7541: uint16(3),
	7542: uint16(2),
	7543: uint16(sym_comment),
	7544: uint16(sym_comma),
	7545: uint16(783),
	7546: uint16(2),
	7547: uint16(anon_sym_RPAREN),
	7548: uint16(anon_sym_DOLLAR),
	7549: uint16(2),
	7550: uint16(3),
	7551: uint16(2),
	7552: uint16(sym_comment),
	7553: uint16(sym_comma),
	7554: uint16(897),
	7555: uint16(2),
	7556: uint16(anon_sym_LBRACE),
	7557: uint16(anon_sym_AT),
	7558: uint16(3),
	7559: uint16(7),
	7560: uint16(1),
	7561: uint16(anon_sym_LBRACE),
	7562: uint16(127),
	7563: uint16(1),
	7564: uint16(sym_selection_set),
	7565: uint16(3),
	7566: uint16(2),
	7567: uint16(sym_comment),
	7568: uint16(sym_comma),
	7569: uint16(3),
	7570: uint16(865),
	7571: uint16(1),
	7572: uint16(anon_sym_on),
	7573: uint16(867),
	7574: uint16(1),
	7575: uint16(anon_sym_repeatable),
	7576: uint16(3),
	7577: uint16(2),
	7578: uint16(sym_comment),
	7579: uint16(sym_comma),
	7580: uint16(3),
	7581: uint16(781),
	7582: uint16(1),
	7583: uint16(sym_name),
	7584: uint16(310),
	7585: uint16(1),
	7586: uint16(sym_named_type),
	7587: uint16(3),
	7588: uint16(2),
	7589: uint16(sym_comment),
	7590: uint16(sym_comma),
	7591: uint16(2),
	7592: uint16(3),
	7593: uint16(2),
	7594: uint16(sym_comment),
	7595: uint16(sym_comma),
	7596: uint16(899),
	7597: uint16(2),
	7598: uint16(anon_sym_RPAREN),
	7599: uint16(anon_sym_DOLLAR),
	7600: uint16(3),
	7601: uint16(3),
	7602: uint16(1),
	7603: uint16(sym_comment),
	7604: uint16(901),
	7605: uint16(1),
	7606: uint16(sym_comma),
	7607: uint16(899),
	7608: uint16(2),
	7609: uint16(anon_sym_RPAREN),
	7610: uint16(anon_sym_DOLLAR),
	7611: uint16(3),
	7612: uint16(781),
	7613: uint16(1),
	7614: uint16(sym_name),
	7615: uint16(255),
	7616: uint16(1),
	7617: uint16(sym_named_type),
	7618: uint16(3),
	7619: uint16(2),
	7620: uint16(sym_comment),
	7621: uint16(sym_comma),
	7622: uint16(2),
	7623: uint16(3),
	7624: uint16(2),
	7625: uint16(sym_comment),
	7626: uint16(sym_comma),
	7627: uint16(903),
	7628: uint16(2),
	7629: uint16(anon_sym_LBRACE),
	7630: uint16(anon_sym_AT),
	7631: uint16(3),
	7632: uint16(905),
	7633: uint16(1),
	7634: uint16(anon_sym_on),
	7635: uint16(208),
	7636: uint16(1),
	7637: uint16(sym_type_condition),
	7638: uint16(3),
	7639: uint16(2),
	7640: uint16(sym_comment),
	7641: uint16(sym_comma),
	7642: uint16(3),
	7643: uint16(907),
	7644: uint16(1),
	7645: uint16(sym_name),
	7646: uint16(311),
	7647: uint16(1),
	7648: uint16(sym_fragment_name),
	7649: uint16(3),
	7650: uint16(2),
	7651: uint16(sym_comment),
	7652: uint16(sym_comma),
	7653: uint16(3),
	7654: uint16(7),
	7655: uint16(1),
	7656: uint16(anon_sym_LBRACE),
	7657: uint16(96),
	7658: uint16(1),
	7659: uint16(sym_selection_set),
	7660: uint16(3),
	7661: uint16(2),
	7662: uint16(sym_comment),
	7663: uint16(sym_comma),
	7664: uint16(3),
	7665: uint16(7),
	7666: uint16(1),
	7667: uint16(anon_sym_LBRACE),
	7668: uint16(122),
	7669: uint16(1),
	7670: uint16(sym_selection_set),
	7671: uint16(3),
	7672: uint16(2),
	7673: uint16(sym_comment),
	7674: uint16(sym_comma),
	7675: uint16(3),
	7676: uint16(602),
	7677: uint16(1),
	7678: uint16(anon_sym_LBRACE),
	7679: uint16(294),
	7680: uint16(1),
	7681: uint16(sym_selection_set),
	7682: uint16(3),
	7683: uint16(2),
	7684: uint16(sym_comment),
	7685: uint16(sym_comma),
	7686: uint16(2),
	7687: uint16(3),
	7688: uint16(2),
	7689: uint16(sym_comment),
	7690: uint16(sym_comma),
	7691: uint16(909),
	7692: uint16(2),
	7693: uint16(anon_sym_RPAREN),
	7694: uint16(anon_sym_DOLLAR),
	7695: uint16(2),
	7696: uint16(3),
	7697: uint16(2),
	7698: uint16(sym_comment),
	7699: uint16(sym_comma),
	7700: uint16(911),
	7701: uint16(2),
	7702: uint16(anon_sym_RPAREN),
	7703: uint16(sym_name),
	7704: uint16(3),
	7705: uint16(3),
	7706: uint16(1),
	7707: uint16(sym_comment),
	7708: uint16(915),
	7709: uint16(1),
	7710: uint16(sym_comma),
	7711: uint16(913),
	7712: uint16(2),
	7713: uint16(anon_sym_RBRACE),
	7714: uint16(sym_name),
	7715: uint16(3),
	7716: uint16(660),
	7717: uint16(1),
	7718: uint16(sym_name),
	7719: uint16(173),
	7720: uint16(1),
	7721: uint16(sym_enum_value),
	7722: uint16(3),
	7723: uint16(2),
	7724: uint16(sym_comment),
	7725: uint16(sym_comma),
	7726: uint16(3),
	7727: uint16(781),
	7728: uint16(1),
	7729: uint16(sym_name),
	7730: uint16(70),
	7731: uint16(1),
	7732: uint16(sym_named_type),
	7733: uint16(3),
	7734: uint16(2),
	7735: uint16(sym_comment),
	7736: uint16(sym_comma),
	7737: uint16(2),
	7738: uint16(3),
	7739: uint16(2),
	7740: uint16(sym_comment),
	7741: uint16(sym_comma),
	7742: uint16(917),
	7743: uint16(2),
	7744: uint16(anon_sym_RBRACE),
	7745: uint16(sym_name),
	7746: uint16(3),
	7747: uint16(919),
	7748: uint16(1),
	7749: uint16(anon_sym_on),
	7750: uint16(921),
	7751: uint16(1),
	7752: uint16(anon_sym_repeatable),
	7753: uint16(3),
	7754: uint16(2),
	7755: uint16(sym_comment),
	7756: uint16(sym_comma),
	7757: uint16(3),
	7758: uint16(7),
	7759: uint16(1),
	7760: uint16(anon_sym_LBRACE),
	7761: uint16(124),
	7762: uint16(1),
	7763: uint16(sym_selection_set),
	7764: uint16(3),
	7765: uint16(2),
	7766: uint16(sym_comment),
	7767: uint16(sym_comma),
	7768: uint16(3),
	7769: uint16(602),
	7770: uint16(1),
	7771: uint16(anon_sym_LBRACE),
	7772: uint16(289),
	7773: uint16(1),
	7774: uint16(sym_selection_set),
	7775: uint16(3),
	7776: uint16(2),
	7777: uint16(sym_comment),
	7778: uint16(sym_comma),
	7779: uint16(3),
	7780: uint16(781),
	7781: uint16(1),
	7782: uint16(sym_name),
	7783: uint16(46),
	7784: uint16(1),
	7785: uint16(sym_named_type),
	7786: uint16(3),
	7787: uint16(2),
	7788: uint16(sym_comment),
	7789: uint16(sym_comma),
	7790: uint16(2),
	7791: uint16(923),
	7792: uint16(1),
	7793: uint16(anon_sym_AT),
	7794: uint16(3),
	7795: uint16(2),
	7796: uint16(sym_comment),
	7797: uint16(sym_comma),
	7798: uint16(2),
	7799: uint16(925),
	7800: uint16(1),
	7801: uint16(anon_sym_DQUOTE),
	7802: uint16(3),
	7803: uint16(2),
	7804: uint16(sym_comment),
	7805: uint16(sym_comma),
	7806: uint16(2),
	7807: uint16(808),
	7808: uint16(1),
	7809: uint16(anon_sym_LBRACE),
	7810: uint16(3),
	7811: uint16(2),
	7812: uint16(sym_comment),
	7813: uint16(sym_comma),
	7814: uint16(2),
	7815: uint16(927),
	7816: uint16(1),
	7817: uint16(sym_name),
	7818: uint16(3),
	7819: uint16(2),
	7820: uint16(sym_comment),
	7821: uint16(sym_comma),
	7822: uint16(2),
	7823: uint16(929),
	7824: uint16(1),
	7825: uint16(sym_name),
	7826: uint16(3),
	7827: uint16(2),
	7828: uint16(sym_comment),
	7829: uint16(sym_comma),
	7830: uint16(2),
	7831: uint16(931),
	7832: uint16(1),
	7833: uint16(sym_name),
	7834: uint16(3),
	7835: uint16(2),
	7836: uint16(sym_comment),
	7837: uint16(sym_comma),
	7838: uint16(2),
	7839: uint16(933),
	7840: uint16(1),
	7841: uint16(sym_name),
	7842: uint16(3),
	7843: uint16(2),
	7844: uint16(sym_comment),
	7845: uint16(sym_comma),
	7846: uint16(2),
	7847: uint16(935),
	7848: uint16(1),
	7849: uint16(aux_sym_string_value_token1),
	7850: uint16(937),
	7851: uint16(2),
	7852: uint16(sym_comment),
	7853: uint16(sym_comma),
	7854: uint16(2),
	7855: uint16(939),
	7856: uint16(1),
	7857: uint16(sym_name),
	7858: uint16(3),
	7859: uint16(2),
	7860: uint16(sym_comment),
	7861: uint16(sym_comma),
	7862: uint16(2),
	7863: uint16(941),
	7864: uint16(1),
	7865: uint16(sym_name),
	7866: uint16(3),
	7867: uint16(2),
	7868: uint16(sym_comment),
	7869: uint16(sym_comma),
	7870: uint16(2),
	7871: uint16(943),
	7872: uint16(1),
	7873: uint16(sym_name),
	7874: uint16(3),
	7875: uint16(2),
	7876: uint16(sym_comment),
	7877: uint16(sym_comma),
	7878: uint16(2),
	7879: uint16(893),
	7880: uint16(1),
	7881: uint16(anon_sym_COLON),
	7882: uint16(3),
	7883: uint16(2),
	7884: uint16(sym_comment),
	7885: uint16(sym_comma),
	7886: uint16(2),
	7887: uint16(945),
	7888: uint16(1),
	7889: uint16(anon_sym_COLON),
	7890: uint16(3),
	7891: uint16(2),
	7892: uint16(sym_comment),
	7893: uint16(sym_comma),
	7894: uint16(2),
	7895: uint16(947),
	7896: uint16(1),
	7897: uint16(anon_sym_RBRACE),
	7898: uint16(3),
	7899: uint16(2),
	7900: uint16(sym_comment),
	7901: uint16(sym_comma),
	7902: uint16(2),
	7903: uint16(949),
	7904: uint16(1),
	7905: uint16(sym_name),
	7906: uint16(3),
	7907: uint16(2),
	7908: uint16(sym_comment),
	7909: uint16(sym_comma),
	7910: uint16(2),
	7911: uint16(951),
	7912: uint16(1),
	7913: uint16(anon_sym_COLON),
	7914: uint16(3),
	7915: uint16(2),
	7916: uint16(sym_comment),
	7917: uint16(sym_comma),
	7918: uint16(2),
	7919: uint16(953),
	7920: uint16(1),
	7921: uint16(anon_sym_COLON),
	7922: uint16(3),
	7923: uint16(2),
	7924: uint16(sym_comment),
	7925: uint16(sym_comma),
	7926: uint16(2),
	7927: uint16(955),
	7928: uint16(1),
	7929: uint16(aux_sym_string_value_token2),
	7930: uint16(937),
	7931: uint16(2),
	7932: uint16(sym_comment),
	7933: uint16(sym_comma),
	7934: uint16(2),
	7935: uint16(957),
	7936: uint16(1),
	7937: uint16(anon_sym_COLON),
	7938: uint16(3),
	7939: uint16(2),
	7940: uint16(sym_comment),
	7941: uint16(sym_comma),
	7942: uint16(2),
	7943: uint16(959),
	7944: uint16(1),
	7945: uint16(sym_name),
	7946: uint16(3),
	7947: uint16(2),
	7948: uint16(sym_comment),
	7949: uint16(sym_comma),
	7950: uint16(2),
	7951: uint16(865),
	7952: uint16(1),
	7953: uint16(anon_sym_on),
	7954: uint16(3),
	7955: uint16(2),
	7956: uint16(sym_comment),
	7957: uint16(sym_comma),
	7958: uint16(2),
	7959: uint16(961),
	7960: uint16(1),
	7961: uint16(sym_name),
	7962: uint16(3),
	7963: uint16(2),
	7964: uint16(sym_comment),
	7965: uint16(sym_comma),
	7966: uint16(2),
	7967: uint16(963),
	7968: uint16(1),
	7969: uint16(anon_sym_AT),
	7970: uint16(3),
	7971: uint16(2),
	7972: uint16(sym_comment),
	7973: uint16(sym_comma),
	7974: uint16(2),
	7975: uint16(965),
	7976: uint16(1),
	7977: uint16(sym_name),
	7978: uint16(3),
	7979: uint16(2),
	7980: uint16(sym_comment),
	7981: uint16(sym_comma),
	7982: uint16(2),
	7983: uint16(967),
	7984: uint16(1),
	7985: uint16(sym_name),
	7986: uint16(3),
	7987: uint16(2),
	7988: uint16(sym_comment),
	7989: uint16(sym_comma),
	7990: uint16(2),
	7991: uint16(969),
	7992: uint16(1),
	7993: uint16(anon_sym_LBRACE),
	7994: uint16(3),
	7995: uint16(2),
	7996: uint16(sym_comment),
	7997: uint16(sym_comma),
	7998: uint16(2),
	7999: uint16(971),
	8000: uint16(1),
	8002: uint16(3),
	8003: uint16(2),
	8004: uint16(sym_comment),
	8005: uint16(sym_comma),
	8006: uint16(2),
	8007: uint16(973),
	8008: uint16(1),
	8009: uint16(anon_sym_COLON),
	8010: uint16(3),
	8011: uint16(2),
	8012: uint16(sym_comment),
	8013: uint16(sym_comma),
	8014: uint16(2),
	8015: uint16(975),
	8016: uint16(1),
	8018: uint16(3),
	8019: uint16(2),
	8020: uint16(sym_comment),
	8021: uint16(sym_comma),
	8022: uint16(2),
	8023: uint16(977),
	8024: uint16(1),
	8025: uint16(sym_name),
	8026: uint16(3),
	8027: uint16(2),
	8028: uint16(sym_comment),
	8029: uint16(sym_comma),
	8030: uint16(2),
	8031: uint16(979),
	8032: uint16(1),
	8033: uint16(anon_sym_COLON),
	8034: uint16(3),
	8035: uint16(2),
	8036: uint16(sym_comment),
	8037: uint16(sym_comma),
	8038: uint16(2),
	8039: uint16(981),
	8040: uint16(1),
	8041: uint16(sym_name),
	8042: uint16(3),
	8043: uint16(2),
	8044: uint16(sym_comment),
	8045: uint16(sym_comma),
	8046: uint16(2),
	8047: uint16(857),
	8048: uint16(1),
	8049: uint16(anon_sym_on),
	8050: uint16(3),
	8051: uint16(2),
	8052: uint16(sym_comment),
	8053: uint16(sym_comma),
	8054: uint16(2),
	8055: uint16(983),
	8056: uint16(1),
	8057: uint16(sym_name),
	8058: uint16(3),
	8059: uint16(2),
	8060: uint16(sym_comment),
	8061: uint16(sym_comma),
	8062: uint16(2),
	8063: uint16(985),
	8064: uint16(1),
	8065: uint16(anon_sym_LBRACE),
	8066: uint16(3),
	8067: uint16(2),
	8068: uint16(sym_comment),
	8069: uint16(sym_comma),
	8070: uint16(2),
	8071: uint16(679),
	8072: uint16(1),
	8073: uint16(sym_name),
	8074: uint16(3),
	8075: uint16(2),
	8076: uint16(sym_comment),
	8077: uint16(sym_comma),
	8078: uint16(2),
	8079: uint16(987),
	8080: uint16(1),
	8081: uint16(sym_name),
	8082: uint16(3),
	8083: uint16(2),
	8084: uint16(sym_comment),
	8085: uint16(sym_comma),
	8086: uint16(2),
	8087: uint16(989),
	8088: uint16(1),
	8089: uint16(sym_name),
	8090: uint16(3),
	8091: uint16(2),
	8092: uint16(sym_comment),
	8093: uint16(sym_comma),
	8094: uint16(2),
	8095: uint16(925),
	8096: uint16(1),
	8097: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	8098: uint16(3),
	8099: uint16(2),
	8100: uint16(sym_comment),
	8101: uint16(sym_comma),
	8102: uint16(2),
	8103: uint16(991),
	8104: uint16(1),
	8105: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	8106: uint16(3),
	8107: uint16(2),
	8108: uint16(sym_comment),
	8109: uint16(sym_comma),
	8110: uint16(2),
	8111: uint16(991),
	8112: uint16(1),
	8113: uint16(anon_sym_DQUOTE),
	8114: uint16(3),
	8115: uint16(2),
	8116: uint16(sym_comment),
	8117: uint16(sym_comma),
	8118: uint16(2),
	8119: uint16(993),
	8120: uint16(1),
	8121: uint16(sym_name),
	8122: uint16(3),
	8123: uint16(2),
	8124: uint16(sym_comment),
	8125: uint16(sym_comma),
	8126: uint16(2),
	8127: uint16(995),
	8128: uint16(1),
	8129: uint16(sym_name),
	8130: uint16(3),
	8131: uint16(2),
	8132: uint16(sym_comment),
	8133: uint16(sym_comma),
	8134: uint16(2),
	8135: uint16(997),
	8136: uint16(1),
	8137: uint16(anon_sym_COLON),
	8138: uint16(3),
	8139: uint16(2),
	8140: uint16(sym_comment),
	8141: uint16(sym_comma),
	8142: uint16(2),
	8143: uint16(999),
	8144: uint16(1),
	8145: uint16(anon_sym_on),
	8146: uint16(3),
	8147: uint16(2),
	8148: uint16(sym_comment),
	8149: uint16(sym_comma),
	8150: uint16(2),
	8151: uint16(1001),
	8152: uint16(1),
	8153: uint16(sym_name),
	8154: uint16(3),
	8155: uint16(2),
	8156: uint16(sym_comment),
	8157: uint16(sym_comma),
	8158: uint16(2),
	8159: uint16(1003),
	8160: uint16(1),
	8161: uint16(anon_sym_RBRACK),
	8162: uint16(3),
	8163: uint16(2),
	8164: uint16(sym_comment),
	8165: uint16(sym_comma),
	8166: uint16(2),
	8167: uint16(1005),
	8168: uint16(1),
	8169: uint16(sym_name),
	8170: uint16(3),
	8171: uint16(2),
	8172: uint16(sym_comment),
	8173: uint16(sym_comma),
	8174: uint16(2),
	8175: uint16(1007),
	8176: uint16(1),
	8177: uint16(sym_name),
	8178: uint16(3),
	8179: uint16(2),
	8180: uint16(sym_comment),
	8181: uint16(sym_comma),
	8182: uint16(2),
	8183: uint16(1009),
	8184: uint16(1),
	8185: uint16(sym_name),
	8186: uint16(3),
	8187: uint16(2),
	8188: uint16(sym_comment),
	8189: uint16(sym_comma),
	8190: uint16(2),
	8191: uint16(1011),
	8192: uint16(1),
	8193: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	8194: uint16(3),
	8195: uint16(2),
	8196: uint16(sym_comment),
	8197: uint16(sym_comma),
	8198: uint16(2),
	8199: uint16(1011),
	8200: uint16(1),
	8201: uint16(anon_sym_DQUOTE),
	8202: uint16(3),
	8203: uint16(2),
	8204: uint16(sym_comment),
	8205: uint16(sym_comma),
	8206: uint16(2),
	8207: uint16(1013),
	8208: uint16(1),
	8209: uint16(sym_name),
	8210: uint16(3),
	8211: uint16(2),
	8212: uint16(sym_comment),
	8213: uint16(sym_comma),
	8214: uint16(2),
	8215: uint16(1015),
	8216: uint16(1),
	8217: uint16(anon_sym_RBRACE),
	8218: uint16(3),
	8219: uint16(2),
	8220: uint16(sym_comment),
	8221: uint16(sym_comma),
	8222: uint16(2),
	8223: uint16(919),
	8224: uint16(1),
	8225: uint16(anon_sym_on),
	8226: uint16(3),
	8227: uint16(2),
	8228: uint16(sym_comment),
	8229: uint16(sym_comma),
	8230: uint16(2),
	8231: uint16(1017),
	8232: uint16(1),
	8233: uint16(sym_name),
	8234: uint16(3),
	8235: uint16(2),
	8236: uint16(sym_comment),
	8237: uint16(sym_comma),
	8238: uint16(2),
	8239: uint16(1019),
	8240: uint16(1),
	8241: uint16(sym_name),
	8242: uint16(3),
	8243: uint16(2),
	8244: uint16(sym_comment),
	8245: uint16(sym_comma),
	8246: uint16(2),
	8247: uint16(1021),
	8248: uint16(1),
	8249: uint16(anon_sym_RBRACK),
	8250: uint16(3),
	8251: uint16(2),
	8252: uint16(sym_comment),
	8253: uint16(sym_comma),
	8254: uint16(2),
	8255: uint16(1023),
	8256: uint16(1),
	8257: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	8258: uint16(3),
	8259: uint16(2),
	8260: uint16(sym_comment),
	8261: uint16(sym_comma),
	8262: uint16(2),
	8263: uint16(1023),
	8264: uint16(1),
	8265: uint16(anon_sym_DQUOTE),
	8266: uint16(3),
	8267: uint16(2),
	8268: uint16(sym_comment),
	8269: uint16(sym_comma),
	8270: uint16(2),
	8271: uint16(1025),
	8272: uint16(1),
	8273: uint16(sym_name),
	8274: uint16(3),
	8275: uint16(2),
	8276: uint16(sym_comment),
	8277: uint16(sym_comma),
	8278: uint16(2),
	8279: uint16(1027),
	8280: uint16(1),
	8281: uint16(aux_sym_string_value_token1),
	8282: uint16(937),
	8283: uint16(2),
	8284: uint16(sym_comment),
	8285: uint16(sym_comma),
	8286: uint16(2),
	8287: uint16(1029),
	8288: uint16(1),
	8289: uint16(aux_sym_string_value_token2),
	8290: uint16(937),
	8291: uint16(2),
	8292: uint16(sym_comment),
	8293: uint16(sym_comma),
	8294: uint16(2),
	8295: uint16(1031),
	8296: uint16(1),
	8297: uint16(sym_name),
	8298: uint16(3),
	8299: uint16(2),
	8300: uint16(sym_comment),
	8301: uint16(sym_comma),
	8302: uint16(2),
	8303: uint16(1033),
	8304: uint16(1),
	8305: uint16(sym_name),
	8306: uint16(3),
	8307: uint16(2),
	8308: uint16(sym_comment),
	8309: uint16(sym_comma),
	8310: uint16(2),
	8311: uint16(1035),
	8312: uint16(1),
	8313: uint16(aux_sym_string_value_token1),
	8314: uint16(937),
	8315: uint16(2),
	8316: uint16(sym_comment),
	8317: uint16(sym_comma),
	8318: uint16(2),
	8319: uint16(1037),
	8320: uint16(1),
	8321: uint16(aux_sym_string_value_token2),
	8322: uint16(937),
	8323: uint16(2),
	8324: uint16(sym_comment),
	8325: uint16(sym_comma),
	8326: uint16(2),
	8327: uint16(1039),
	8328: uint16(1),
	8329: uint16(sym_name),
	8330: uint16(3),
	8331: uint16(2),
	8332: uint16(sym_comment),
	8333: uint16(sym_comma),
	8334: uint16(2),
	8335: uint16(1041),
	8336: uint16(1),
	8337: uint16(aux_sym_string_value_token1),
	8338: uint16(937),
	8339: uint16(2),
	8340: uint16(sym_comment),
	8341: uint16(sym_comma),
	8342: uint16(2),
	8343: uint16(1043),
	8344: uint16(1),
	8345: uint16(aux_sym_string_value_token2),
	8346: uint16(937),
	8347: uint16(2),
	8348: uint16(sym_comment),
	8349: uint16(sym_comma),
}

var ts_small_parse_table_map = [394]uint32_t{
	1:   uint32(102),
	2:   uint32(204),
	3:   uint32(251),
	4:   uint32(296),
	5:   uint32(343),
	6:   uint32(388),
	7:   uint32(435),
	8:   uint32(482),
	9:   uint32(527),
	10:  uint32(574),
	11:  uint32(619),
	12:  uint32(666),
	13:  uint32(708),
	14:  uint32(740),
	15:  uint32(784),
	16:  uint32(828),
	17:  uint32(872),
	18:  uint32(914),
	19:  uint32(958),
	20:  uint32(1002),
	21:  uint32(1044),
	22:  uint32(1086),
	23:  uint32(1130),
	24:  uint32(1171),
	25:  uint32(1212),
	26:  uint32(1251),
	27:  uint32(1292),
	28:  uint32(1331),
	29:  uint32(1372),
	30:  uint32(1413),
	31:  uint32(1447),
	32:  uint32(1483),
	33:  uint32(1517),
	34:  uint32(1553),
	35:  uint32(1587),
	36:  uint32(1636),
	37:  uint32(1685),
	38:  uint32(1734),
	39:  uint32(1783),
	40:  uint32(1832),
	41:  uint32(1881),
	42:  uint32(1930),
	43:  uint32(1962),
	44:  uint32(1994),
	45:  uint32(2022),
	46:  uint32(2054),
	47:  uint32(2086),
	48:  uint32(2118),
	49:  uint32(2150),
	50:  uint32(2178),
	51:  uint32(2206),
	52:  uint32(2234),
	53:  uint32(2266),
	54:  uint32(2311),
	55:  uint32(2340),
	56:  uint32(2367),
	57:  uint32(2398),
	58:  uint32(2429),
	59:  uint32(2460),
	60:  uint32(2505),
	61:  uint32(2536),
	62:  uint32(2567),
	63:  uint32(2598),
	64:  uint32(2627),
	65:  uint32(2658),
	66:  uint32(2689),
	67:  uint32(2720),
	68:  uint32(2749),
	69:  uint32(2776),
	70:  uint32(2803),
	71:  uint32(2830),
	72:  uint32(2859),
	73:  uint32(2886),
	74:  uint32(2913),
	75:  uint32(2944),
	76:  uint32(2975),
	77:  uint32(3006),
	78:  uint32(3035),
	79:  uint32(3062),
	80:  uint32(3093),
	81:  uint32(3122),
	82:  uint32(3167),
	83:  uint32(3198),
	84:  uint32(3225),
	85:  uint32(3254),
	86:  uint32(3283),
	87:  uint32(3328),
	88:  uint32(3357),
	89:  uint32(3388),
	90:  uint32(3414),
	91:  uint32(3440),
	92:  uint32(3466),
	93:  uint32(3492),
	94:  uint32(3518),
	95:  uint32(3544),
	96:  uint32(3570),
	97:  uint32(3596),
	98:  uint32(3622),
	99:  uint32(3648),
	100: uint32(3674),
	101: uint32(3700),
	102: uint32(3726),
	103: uint32(3752),
	104: uint32(3778),
	105: uint32(3804),
	106: uint32(3830),
	107: uint32(3856),
	108: uint32(3882),
	109: uint32(3908),
	110: uint32(3934),
	111: uint32(3960),
	112: uint32(3986),
	113: uint32(4012),
	114: uint32(4038),
	115: uint32(4064),
	116: uint32(4090),
	117: uint32(4116),
	118: uint32(4142),
	119: uint32(4168),
	120: uint32(4194),
	121: uint32(4220),
	122: uint32(4246),
	123: uint32(4272),
	124: uint32(4298),
	125: uint32(4324),
	126: uint32(4350),
	127: uint32(4376),
	128: uint32(4402),
	129: uint32(4428),
	130: uint32(4454),
	131: uint32(4480),
	132: uint32(4506),
	133: uint32(4532),
	134: uint32(4558),
	135: uint32(4584),
	136: uint32(4610),
	137: uint32(4631),
	138: uint32(4652),
	139: uint32(4673),
	140: uint32(4694),
	141: uint32(4715),
	142: uint32(4736),
	143: uint32(4757),
	144: uint32(4778),
	145: uint32(4799),
	146: uint32(4829),
	147: uint32(4859),
	148: uint32(4892),
	149: uint32(4923),
	150: uint32(4946),
	151: uint32(4968),
	152: uint32(4996),
	153: uint32(5020),
	154: uint32(5042),
	155: uint32(5072),
	156: uint32(5098),
	157: uint32(5126),
	158: uint32(5152),
	159: uint32(5170),
	160: uint32(5190),
	161: uint32(5220),
	162: uint32(5244),
	163: uint32(5274),
	164: uint32(5300),
	165: uint32(5315),
	166: uint32(5338),
	167: uint32(5355),
	168: uint32(5382),
	169: uint32(5409),
	170: uint32(5432),
	171: uint32(5459),
	172: uint32(5482),
	173: uint32(5505),
	174: uint32(5522),
	175: uint32(5537),
	176: uint32(5560),
	177: uint32(5577),
	178: uint32(5600),
	179: uint32(5627),
	180: uint32(5654),
	181: uint32(5683),
	182: uint32(5700),
	183: uint32(5727),
	184: uint32(5754),
	185: uint32(5777),
	186: uint32(5794),
	187: uint32(5810),
	188: uint32(5834),
	189: uint32(5854),
	190: uint32(5878),
	191: uint32(5898),
	192: uint32(5918),
	193: uint32(5942),
	194: uint32(5964),
	195: uint32(5984),
	196: uint32(6004),
	197: uint32(6026),
	198: uint32(6052),
	199: uint32(6069),
	200: uint32(6084),
	201: uint32(6105),
	202: uint32(6122),
	203: uint32(6143),
	204: uint32(6158),
	205: uint32(6179),
	206: uint32(6200),
	207: uint32(6221),
	208: uint32(6236),
	209: uint32(6251),
	210: uint32(6272),
	211: uint32(6289),
	212: uint32(6304),
	213: uint32(6319),
	214: uint32(6340),
	215: uint32(6361),
	216: uint32(6382),
	217: uint32(6403),
	218: uint32(6418),
	219: uint32(6439),
	220: uint32(6460),
	221: uint32(6481),
	222: uint32(6496),
	223: uint32(6511),
	224: uint32(6523),
	225: uint32(6539),
	226: uint32(6557),
	227: uint32(6575),
	228: uint32(6591),
	229: uint32(6605),
	230: uint32(6621),
	231: uint32(6633),
	232: uint32(6645),
	233: uint32(6657),
	234: uint32(6669),
	235: uint32(6683),
	236: uint32(6701),
	237: uint32(6717),
	238: uint32(6733),
	239: uint32(6745),
	240: uint32(6763),
	241: uint32(6777),
	242: uint32(6789),
	243: uint32(6801),
	244: uint32(6813),
	245: uint32(6831),
	246: uint32(6847),
	247: uint32(6863),
	248: uint32(6877),
	249: uint32(6893),
	250: uint32(6905),
	251: uint32(6917),
	252: uint32(6929),
	253: uint32(6942),
	254: uint32(6953),
	255: uint32(6968),
	256: uint32(6983),
	257: uint32(6998),
	258: uint32(7013),
	259: uint32(7024),
	260: uint32(7039),
	261: uint32(7054),
	262: uint32(7069),
	263: uint32(7084),
	264: uint32(7099),
	265: uint32(7116),
	266: uint32(7131),
	267: uint32(7144),
	268: uint32(7159),
	269: uint32(7170),
	270: uint32(7183),
	271: uint32(7196),
	272: uint32(7211),
	273: uint32(7228),
	274: uint32(7239),
	275: uint32(7254),
	276: uint32(7269),
	277: uint32(7282),
	278: uint32(7297),
	279: uint32(7307),
	280: uint32(7317),
	281: uint32(7329),
	282: uint32(7339),
	283: uint32(7353),
	284: uint32(7365),
	285: uint32(7375),
	286: uint32(7385),
	287: uint32(7395),
	288: uint32(7405),
	289: uint32(7415),
	290: uint32(7425),
	291: uint32(7435),
	292: uint32(7445),
	293: uint32(7455),
	294: uint32(7469),
	295: uint32(7481),
	296: uint32(7491),
	297: uint32(7505),
	298: uint32(7519),
	299: uint32(7529),
	300: uint32(7540),
	301: uint32(7549),
	302: uint32(7558),
	303: uint32(7569),
	304: uint32(7580),
	305: uint32(7591),
	306: uint32(7600),
	307: uint32(7611),
	308: uint32(7622),
	309: uint32(7631),
	310: uint32(7642),
	311: uint32(7653),
	312: uint32(7664),
	313: uint32(7675),
	314: uint32(7686),
	315: uint32(7695),
	316: uint32(7704),
	317: uint32(7715),
	318: uint32(7726),
	319: uint32(7737),
	320: uint32(7746),
	321: uint32(7757),
	322: uint32(7768),
	323: uint32(7779),
	324: uint32(7790),
	325: uint32(7798),
	326: uint32(7806),
	327: uint32(7814),
	328: uint32(7822),
	329: uint32(7830),
	330: uint32(7838),
	331: uint32(7846),
	332: uint32(7854),
	333: uint32(7862),
	334: uint32(7870),
	335: uint32(7878),
	336: uint32(7886),
	337: uint32(7894),
	338: uint32(7902),
	339: uint32(7910),
	340: uint32(7918),
	341: uint32(7926),
	342: uint32(7934),
	343: uint32(7942),
	344: uint32(7950),
	345: uint32(7958),
	346: uint32(7966),
	347: uint32(7974),
	348: uint32(7982),
	349: uint32(7990),
	350: uint32(7998),
	351: uint32(8006),
	352: uint32(8014),
	353: uint32(8022),
	354: uint32(8030),
	355: uint32(8038),
	356: uint32(8046),
	357: uint32(8054),
	358: uint32(8062),
	359: uint32(8070),
	360: uint32(8078),
	361: uint32(8086),
	362: uint32(8094),
	363: uint32(8102),
	364: uint32(8110),
	365: uint32(8118),
	366: uint32(8126),
	367: uint32(8134),
	368: uint32(8142),
	369: uint32(8150),
	370: uint32(8158),
	371: uint32(8166),
	372: uint32(8174),
	373: uint32(8182),
	374: uint32(8190),
	375: uint32(8198),
	376: uint32(8206),
	377: uint32(8214),
	378: uint32(8222),
	379: uint32(8230),
	380: uint32(8238),
	381: uint32(8246),
	382: uint32(8254),
	383: uint32(8262),
	384: uint32(8270),
	385: uint32(8278),
	386: uint32(8286),
	387: uint32(8294),
	388: uint32(8302),
	389: uint32(8310),
	390: uint32(8318),
	391: uint32(8326),
	392: uint32(8334),
	393: uint32(8342),
}

var ts_parse_actions = [1045]TSParseActionEntry{
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(246),
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
		Fstate: uint16(177),
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
		Fstate: uint16(199),
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
		Fstate: uint16(374),
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
		Fstate: uint16(371),
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
		Fstate: uint16(367),
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
		Fstate: uint16(363),
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
		Fstate: uint16(350),
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
		Fstate: uint16(347),
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
		Fstate: uint16(252),
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
		Fstate: uint16(333),
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
		Fcount: uint8(1),
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
		Fstate: uint16(343),
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
		Fstate: uint16(312),
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
		Fstate: uint16(348),
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
		Fstate:      uint16(246),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate:      uint16(177),
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
		Fstate:      uint16(199),
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
		Fstate:      uint16(374),
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
		Fcount:    uint8(2),
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
	51: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(371),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	52: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	53: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(367),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate:      uint16(363),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate:      uint16(350),
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
		Fstate:      uint16(347),
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
		Fstate:      uint16(252),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(333),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(343),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(312),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      uint16(348),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_object_type_definition),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	82: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(189),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(298),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_object_type_definition),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(362),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(29),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(71),
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
		Fcount: uint8(1),
	}})),
	94: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(71),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(72),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(72),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_interface_type_extension),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_interface_type_extension),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_object_type_definition),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_object_type_definition),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_object_type_extension),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_object_type_extension),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_interface_type_definition),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_interface_type_definition),
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
		Fsymbol:      uint16(sym_interface_type_definition),
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
		Fsymbol:      uint16(sym_interface_type_definition),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_union_type_definition),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(299),
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
		Fcount: uint8(1),
	}})),
	124: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_union_type_definition),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_named_type),
	})))),
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
		Fcount: uint8(1),
	}})),
	128: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_named_type),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_object_type_definition),
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
		Fstate: uint16(325),
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
		Fcount: uint8(1),
	}})),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_object_type_definition),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_object_type_extension),
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
		Fcount: uint8(1),
	}})),
	138: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_object_type_extension),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	140: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_union_type_definition),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_union_type_definition),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	144: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_interface_type_extension),
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
		Fcount: uint8(1),
	}})),
	146: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_interface_type_extension),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	148: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_input_object_type_extension),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(194),
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
		Fcount: uint8(1),
	}})),
	152: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_input_object_type_extension),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_union_type_extension),
	})))),
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
		Fcount: uint8(1),
	}})),
	156: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_union_type_extension),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_interface_type_definition),
	})))),
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
		Fcount: uint8(1),
	}})),
	160: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_interface_type_definition),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum_type_definition),
	})))),
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
		Fstate: uint16(181),
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
		Fcount: uint8(1),
	}})),
	166: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum_type_definition),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_input_object_type_definition),
	})))),
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
		Fcount: uint8(1),
	}})),
	170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_input_object_type_definition),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_enum_type_definition),
	})))),
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
		Fcount: uint8(1),
	}})),
	174: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_enum_type_definition),
	})))),
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
		Fsymbol:      uint16(sym_input_object_type_definition),
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
		Fcount: uint8(1),
	}})),
	178: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_input_object_type_definition),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum_type_extension),
	})))),
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
		Fcount: uint8(1),
	}})),
	182: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum_type_extension),
	})))),
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
		Fsymbol:      uint16(sym_directive),
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
		Fstate: uint16(285),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_directive),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_scalar_type_definition),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_scalar_type_definition),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_directives_repeat1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_directives_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	198: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_directives_repeat1),
	})))),
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
		Fstate:      uint16(362),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_scalar_type_definition),
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
		Fcount: uint8(1),
	}})),
	203: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_scalar_type_definition),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_directives),
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
		Fcount: uint8(1),
	}})),
	207: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_directives),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	209: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_value_repeat1),
	})))),
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
		Fstate:      uint16(263),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_value_repeat1),
	})))),
	213: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(368),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	215: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_value_repeat1),
	})))),
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
		Fstate:      uint16(391),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	218: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_value_repeat1),
	})))),
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
		Fstate:      uint16(392),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	221: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_value_repeat1),
	})))),
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
		Fstate:      uint16(145),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	224: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_value_repeat1),
	})))),
	225: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(145),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	227: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_value_repeat1),
	})))),
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
		Fstate:      uint16(146),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	230: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_value_repeat1),
	})))),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(41),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_value_repeat1),
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
		Fcount: uint8(2),
	}})),
	235: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_value_repeat1),
	})))),
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
		Fstate:      uint16(143),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(263),
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
		Fstate: uint16(368),
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
		Fstate: uint16(391),
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
		Fstate: uint16(392),
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
		Fcount: uint8(1),
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
		Fstate: uint16(145),
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
		Fstate: uint16(145),
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
		Fcount: uint8(1),
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
		Fstate: uint16(146),
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
		Fstate: uint16(41),
	}})))),
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
		Fstate: uint16(234),
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
		Fstate: uint16(143),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(210),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	260: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(201),
	}})))),
	261: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	262: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(139),
	}})))),
	263: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(141),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	266: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(233),
	}})))),
	267: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	268: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_input_object_type_extension_repeat1),
	})))),
	269: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	270: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_input_object_type_extension_repeat1),
	})))),
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
		Fstate:      uint16(194),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	273: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_input_object_type_extension_repeat1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_input_object_type_extension),
	})))),
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
		Fcount: uint8(1),
	}})),
	277: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_input_object_type_extension),
	})))),
	278: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	279: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_implements_interfaces),
	})))),
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
		Fcount: uint8(1),
	}})),
	281: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_implements_interfaces),
	})))),
	282: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	283: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_union_type_extension),
	})))),
	284: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	285: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_union_type_extension),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_input_object_type_extension),
	})))),
	288: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	289: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_input_object_type_extension),
	})))),
	290: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	291: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_implements_interfaces),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_implements_interfaces),
	})))),
	294: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	295: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_arguments),
	})))),
	296: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	297: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_arguments),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_directive),
	})))),
	300: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	301: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_directive),
	})))),
	302: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	303: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_union_type_definition),
	})))),
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
		Fcount: uint8(1),
	}})),
	305: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_union_type_definition),
	})))),
	306: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	307: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(262),
	}})))),
	308: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(355),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	311: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(387),
	}})))),
	312: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	313: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(388),
	}})))),
	314: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
		Fstate: uint16(205),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	317: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(205),
	}})))),
	318: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	319: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(209),
	}})))),
	320: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(39),
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
		Fcount: uint8(1),
	}})),
	323: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(219),
	}})))),
	324: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	325: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(320),
	}})))),
	326: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	327: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_union_member_types),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_union_member_types),
	})))),
	330: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	331: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_enum_type_definition),
	})))),
	332: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	333: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_enum_type_definition),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_input_object_type_definition),
	})))),
	336: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	337: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_input_object_type_definition),
	})))),
	338: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	339: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_enum_type_extension),
	})))),
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
		Fcount: uint8(1),
	}})),
	341: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_enum_type_extension),
	})))),
	342: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	343: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_object_type_extension),
	})))),
	344: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	345: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_object_type_extension),
	})))),
	346: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	347: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_interface_type_extension),
	})))),
	348: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	349: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_interface_type_extension),
	})))),
	350: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	351: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_union_type_extension),
	})))),
	352: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	353: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_union_type_extension),
	})))),
	354: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	355: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_union_member_types),
	})))),
	356: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	357: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_union_member_types),
	})))),
	358: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	359: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_executable_directive_location),
	})))),
	360: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	361: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_executable_directive_location),
	})))),
	362: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	363: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type_system_directive_location),
	})))),
	364: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	365: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type_system_directive_location),
	})))),
	366: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	367: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_directive_definition),
	})))),
	368: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	369: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(27),
	}})))),
	370: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	371: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_directive_definition),
	})))),
	372: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	373: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_directive_locations),
	})))),
	374: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	375: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_directive_locations),
	})))),
	376: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	377: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_directive_location),
	})))),
	378: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	379: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_directive_location),
	})))),
	380: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	381: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_object_type_definition),
	})))),
	382: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	383: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_object_type_definition),
	})))),
	384: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	385: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_interface_type_definition),
	})))),
	386: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	387: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_interface_type_definition),
	})))),
	388: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	389: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_union_type_definition),
	})))),
	390: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	391: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_union_type_definition),
	})))),
	392: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	393: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_directive_locations),
	})))),
	394: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	395: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_directive_locations),
	})))),
	396: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	397: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_directive_definition),
	})))),
	398: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	399: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_directive_definition),
	})))),
	400: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	401: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(258),
	}})))),
	402: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	403: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(378),
	}})))),
	404: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	405: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(394),
	}})))),
	406: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	407: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(395),
	}})))),
	408: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	409: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(251),
	}})))),
	410: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	411: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(251),
	}})))),
	412: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	413: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(245),
	}})))),
	414: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	415: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(43),
	}})))),
	416: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	417: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(253),
	}})))),
	418: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	419: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_directive_locations),
	})))),
	420: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	421: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_directive_locations),
	})))),
	422: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	423: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_directive_definition),
	})))),
	424: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	425: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_directive_definition),
	})))),
	426: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	427: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_directive_definition),
	})))),
	428: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	429: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_directive_definition),
	})))),
	430: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	431: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_interface_type_definition),
	})))),
	432: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	433: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_interface_type_definition),
	})))),
	434: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	435: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_input_object_type_definition),
	})))),
	436: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	437: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_input_object_type_definition),
	})))),
	438: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	439: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_enum_type_definition),
	})))),
	440: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	441: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_enum_type_definition),
	})))),
	442: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	443: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_schema_extension),
	})))),
	444: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	445: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_schema_extension),
	})))),
	446: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	447: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_fragment_definition),
	})))),
	448: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	449: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_fragment_definition),
	})))),
	450: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	451: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_input_fields_definition),
	})))),
	452: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	453: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_input_fields_definition),
	})))),
	454: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	455: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_object_type_extension),
	})))),
	456: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	457: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_object_type_extension),
	})))),
	458: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	459: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum_values_definition),
	})))),
	460: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	461: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum_values_definition),
	})))),
	462: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	463: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_interface_type_extension),
	})))),
	464: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	465: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_interface_type_extension),
	})))),
	466: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	467: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_selection_set),
	})))),
	468: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	469: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_selection_set),
	})))),
	470: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	471: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_fields_definition),
	})))),
	472: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	473: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_fields_definition),
	})))),
	474: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	475: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_enum_type_extension),
	})))),
	476: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	477: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_enum_type_extension),
	})))),
	478: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	479: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type_definition),
	})))),
	480: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	481: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type_definition),
	})))),
	482: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	483: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_executable_definition),
	})))),
	484: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	485: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_executable_definition),
	})))),
	486: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	487: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type_extension),
	})))),
	488: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	489: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type_extension),
	})))),
	490: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	491: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_schema_extension),
	})))),
	492: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	493: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_schema_extension),
	})))),
	494: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	495: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type_system_extension),
	})))),
	496: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	497: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type_system_extension),
	})))),
	498: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	499: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_schema_definition),
	})))),
	500: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	501: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_schema_definition),
	})))),
	502: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	503: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_object_type_definition),
	})))),
	504: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	505: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_object_type_definition),
	})))),
	506: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	507: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_scalar_type_extension),
	})))),
	508: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	509: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_scalar_type_extension),
	})))),
	510: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	511: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_schema_definition),
	})))),
	512: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	513: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_schema_definition),
	})))),
	514: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	515: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_operation_definition),
	})))),
	516: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	517: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_operation_definition),
	})))),
	518: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	519: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type_system_definition),
	})))),
	520: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	521: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type_system_definition),
	})))),
	522: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	523: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_operation_definition),
	})))),
	524: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	525: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_operation_definition),
	})))),
	526: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	527: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_scalar_type_definition),
	})))),
	528: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	529: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_scalar_type_definition),
	})))),
	530: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	531: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_operation_definition),
	})))),
	532: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	533: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_operation_definition),
	})))),
	534: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	535: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_operation_definition),
	})))),
	536: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	537: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_operation_definition),
	})))),
	538: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	539: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_schema_definition),
	})))),
	540: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	541: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_schema_definition),
	})))),
	542: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	543: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_definition),
	})))),
	544: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	545: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_definition),
	})))),
	546: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	547: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_fragment_definition),
	})))),
	548: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	549: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_fragment_definition),
	})))),
	550: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	551: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_operation_definition),
	})))),
	552: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	553: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_operation_definition),
	})))),
	554: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	555: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_object_value),
	})))),
	556: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	557: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_object_value),
	})))),
	558: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	559: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_list_value),
	})))),
	560: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	561: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_list_value),
	})))),
	562: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	563: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_object_value),
	})))),
	564: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	565: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_object_value),
	})))),
	566: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	567: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list_value),
	})))),
	568: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	569: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list_value),
	})))),
	570: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	571: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_string_value),
	})))),
	572: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	573: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_string_value),
	})))),
	574: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	575: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_enum_value),
	})))),
	576: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	577: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_enum_value),
	})))),
	578: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	579: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_variable),
	})))),
	580: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	581: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_variable),
	})))),
	582: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	583: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_value),
	})))),
	584: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	585: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_value),
	})))),
	586: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	587: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_boolean_value),
	})))),
	588: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	589: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_boolean_value),
	})))),
	590: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	591: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_input_value_definition),
	})))),
	592: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	593: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(61),
	}})))),
	594: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	595: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_input_value_definition),
	})))),
	596: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	597: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(382),
	}})))),
	598: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	599: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_input_value_definition),
	})))),
	600: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	601: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_input_value_definition),
	})))),
	602: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	603: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(167),
	}})))),
	604: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	605: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(306),
	}})))),
	606: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	607: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(270),
	}})))),
	608: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	609: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_field),
	})))),
	610: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	611: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(381),
	}})))),
	612: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	613: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(296),
	}})))),
	614: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	615: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_input_fields_definition_repeat1),
	})))),
	616: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	617: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_input_fields_definition_repeat1),
	})))),
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
		Fstate:      uint16(387),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	620: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_input_fields_definition_repeat1),
	})))),
	621: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(388),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	622: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	623: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_input_fields_definition_repeat1),
	})))),
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
		Fstate:      uint16(344),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	626: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_directives_repeat1),
	})))),
	627: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(382),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	628: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	629: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(279),
	}})))),
	630: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	631: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(172),
	}})))),
	632: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	633: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_selection_set_repeat1),
	})))),
	634: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	635: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_selection_set_repeat1),
	})))),
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
		Fstate:      uint16(149),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	638: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_selection_set_repeat1),
	})))),
	639: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(150),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	640: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	641: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_field),
	})))),
	642: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	643: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(103),
	}})))),
	644: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	645: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(149),
	}})))),
	646: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	647: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(150),
	}})))),
	648: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	649: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list_type),
	})))),
	650: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	651: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list_type),
	})))),
	652: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	653: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type),
	})))),
	654: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	655: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type),
	})))),
	656: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	657: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(175),
	}})))),
	658: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	659: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(99),
	}})))),
	660: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	661: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(219),
	}})))),
	662: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	663: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_input_value_definition),
	})))),
	664: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	665: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_input_value_definition),
	})))),
	666: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	667: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_enum_values_definition_repeat1),
	})))),
	668: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	669: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_enum_values_definition_repeat1),
	})))),
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
		Fstate:      uint16(387),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	672: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_enum_values_definition_repeat1),
	})))),
	673: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(388),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	674: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	675: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_enum_values_definition_repeat1),
	})))),
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
		Fstate:      uint16(219),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(287),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	680: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_description),
	})))),
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
		Fstate: uint16(283),
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
		Fstate: uint16(344),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(104),
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
		Fstate: uint16(284),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_field_definition),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_field_definition),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_enum_value_definition),
	})))),
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
		Fcount: uint8(1),
	}})),
	696: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_enum_value_definition),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	698: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_enum_value_definition),
	})))),
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
		Fcount: uint8(1),
	}})),
	700: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_enum_value_definition),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_non_null_type),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_non_null_type),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_field_definition),
	})))),
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
		Fcount: uint8(1),
	}})),
	708: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_field_definition),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(97),
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
		Fstate: uint16(237),
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
		Fstate: uint16(330),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(331),
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
		Fstate: uint16(332),
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
		Fstate: uint16(334),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(335),
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
		Fstate: uint16(336),
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
		Fstate: uint16(326),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	728: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_variable_definition),
	})))),
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
		Fstate: uint16(83),
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
		Fstate: uint16(375),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(302),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fields_definition_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	738: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fields_definition_repeat1),
	})))),
	739: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(387),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	740: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	741: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fields_definition_repeat1),
	})))),
	742: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(388),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_fields_definition_repeat1),
	})))),
	745: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(284),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	746: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	747: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_field_definition),
	})))),
	748: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	749: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_field_definition),
	})))),
	750: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	751: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_schema_definition_repeat1),
	})))),
	752: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	753: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_schema_definition_repeat1),
	})))),
	754: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(252),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	756: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(121),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(130),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	760: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_field),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(118),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	764: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_fragment_spread),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(227),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	768: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(340),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(393),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	772: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(390),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(389),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	776: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(386),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(373),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	780: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(204),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(15),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	784: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_variable_definition),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(307),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	788: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(168),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(206),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	792: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(243),
	}})))),
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
		Fsymbol:      uint16(sym_default_value),
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
		Fsymbol:      uint16(sym_default_value),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_directives_repeat1),
	})))),
	799: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(375),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	800: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	801: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(231),
	}})))),
	802: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	803: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(303),
	}})))),
	804: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	805: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_input_value_definition),
	})))),
	806: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	807: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_input_value_definition),
	})))),
	808: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	809: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(200),
	}})))),
	810: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	811: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(282),
	}})))),
	812: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	813: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_variable_definitions_repeat1),
	})))),
	814: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	815: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_variable_definitions_repeat1),
	})))),
	816: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(355),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(203),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	820: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_field),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(275),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	824: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_operation_type),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_field_definition),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_field_definition),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_root_operation_type_definition),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	832: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(232),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(369),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	836: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(225),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(281),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	840: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(342),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(138),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	844: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(214),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(140),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	848: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(52),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(191),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	852: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(5),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(346),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	856: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(183),
	}})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_fragment_name),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_enum_value_definition),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_enum_value_definition),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	864: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(224),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(7),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	868: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(380),
	}})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	872: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
	})))),
	873: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(342),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	874: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	875: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_object_value_repeat1),
	})))),
	876: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	877: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_object_value_repeat1),
	})))),
	878: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(369),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_arguments_definition),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(220),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	884: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_inline_fragment),
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
		Fsymbol:      uint16(sym_inline_fragment),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_fragment_spread),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_field),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	892: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_inline_fragment),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(221),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	896: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_selection),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_variable_definitions),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	900: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_variable_definition),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	902: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(316),
	}})))),
	903: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	904: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_type_condition),
	})))),
	905: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(306),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	908: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(358),
	}})))),
	909: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	910: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_variable_definition),
	})))),
	911: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	912: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_argument),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	914: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_object_field),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(321),
	}})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_object_field),
	})))),
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
		Fstate: uint16(10),
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
		Fstate: uint16(370),
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
		Fstate: uint16(349),
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
		Fstate: uint16(176),
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
		Fstate: uint16(295),
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
		Fstate: uint16(33),
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
		Fstate: uint16(8),
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
		Fstate: uint16(11),
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
		Fcount: uint8(1),
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
		Fstate: uint16(364),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(14),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(25),
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
		Fstate: uint16(26),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(309),
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
		Fstate: uint16(115),
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
		Fstate: uint16(265),
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
		Fstate: uint16(211),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	954: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(55),
	}})))),
	955: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	956: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(327),
	}})))),
	957: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(216),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(356),
	}})))),
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
		Fstate: uint16(30),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(357),
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
		Fstate: uint16(274),
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
		Fstate: uint16(28),
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
		Fstate: uint16(212),
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
	972: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(222),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_source_file),
	})))),
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
		Fstate: uint16(188),
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
		Fstate: uint16(218),
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
		Fstate: uint16(266),
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
		Fstate: uint16(158),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(248),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	988: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(32),
	}})))),
	989: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	990: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(19),
	}})))),
	991: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	992: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(213),
	}})))),
	993: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	994: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(13),
	}})))),
	995: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	996: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(144),
	}})))),
	997: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	998: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(88),
	}})))),
	999: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1000: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(12),
	}})))),
	1001: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1002: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(4),
	}})))),
	1003: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1004: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(160),
	}})))),
	1005: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1006: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(22),
	}})))),
	1007: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1008: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(35),
	}})))),
	1009: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1010: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(239),
	}})))),
	1011: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1012: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(142),
	}})))),
	1013: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1014: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(244),
	}})))),
	1015: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1016: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(95),
	}})))),
	1017: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1018: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_alias),
	})))),
	1019: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1020: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(151),
	}})))),
	1021: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1022: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(240),
	}})))),
	1023: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1024: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(235),
	}})))),
	1025: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1026: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(31),
	}})))),
	1027: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1028: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(365),
	}})))),
	1029: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1030: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(366),
	}})))),
	1031: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1032: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(23),
	}})))),
	1033: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1034: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(6),
	}})))),
	1035: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1036: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(376),
	}})))),
	1037: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1038: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(377),
	}})))),
	1039: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1040: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(9),
	}})))),
	1041: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1042: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(384),
	}})))),
	1043: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1044: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(385),
	}})))),
}

func tree_sitter_graphql(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fversion:                   uint32(LANGUAGE_VERSION),
	Fsymbol_count:              uint32(SYMBOL_COUNT),
	Ftoken_count:               uint32(TOKEN_COUNT),
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
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00schema\x00{\x00}\x00extend\x00scalar\x00type\x00interface\x00union\x00enum\x00input\x00&\x00implements\x00:\x00(\x00)\x00=\x00|\x00query\x00mutation\x00subscription\x00$\x00\"\"\"\x00string_value_token1\x00\"\x00string_value_token2\x00int_value\x00float_value\x00true\x00false\x00null_value\x00[\x00]\x00...\x00fragment\x00on\x00@\x00directive\x00repeatable\x00QUERY\x00MUTATION\x00SUBSCRIPTION\x00FIELD\x00FRAGMENT_DEFINITION\x00FRAGMENT_SPREAD\x00INLINE_FRAGMENT\x00VARIABLE_DEFINITION\x00SCHEMA\x00SCALAR\x00OBJECT\x00FIELD_DEFINITION\x00ARGUMENT_DEFINITION\x00INTERFACE\x00UNION\x00ENUM\x00ENUM_VALUE\x00INPUT_OBJECT\x00INPUT_FIELD_DEFINITION\x00!\x00name\x00comment\x00comma\x00source_file\x00document\x00definition\x00executable_definition\x00type_system_definition\x00type_system_extension\x00schema_definition\x00schema_extension\x00type_extension\x00scalar_type_extension\x00object_type_extension\x00interface_type_extension\x00union_type_extension\x00enum_type_extension\x00input_object_type_extension\x00input_fields_definition\x00enum_values_definition\x00enum_value_definition\x00implements_interfaces\x00fields_definition\x00field_definition\x00arguments_definition\x00input_value_definition\x00default_value\x00union_member_types\x00root_operation_type_definition\x00operation_definition\x00operation_type\x00type_definition\x00scalar_type_definition\x00object_type_definition\x00interface_type_definition\x00union_type_definition\x00enum_type_definition\x00input_object_type_definition\x00variable_definitions\x00variable_definition\x00selection_set\x00selection\x00field\x00alias\x00arguments\x00argument\x00value\x00variable\x00string_value\x00boolean_value\x00enum_value\x00list_value\x00object_value\x00object_field\x00fragment_spread\x00fragment_definition\x00fragment_name\x00inline_fragment\x00type_condition\x00directives\x00directive_definition\x00directive_locations\x00directive_location\x00executable_directive_location\x00type_system_directive_location\x00named_type\x00list_type\x00non_null_type\x00description\x00document_repeat1\x00schema_definition_repeat1\x00input_object_type_extension_repeat1\x00input_fields_definition_repeat1\x00enum_values_definition_repeat1\x00fields_definition_repeat1\x00variable_definitions_repeat1\x00selection_set_repeat1\x00arguments_repeat1\x00list_value_repeat1\x00object_value_repeat1\x00directives_repeat1\x00"
