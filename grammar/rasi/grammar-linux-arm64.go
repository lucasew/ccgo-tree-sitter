// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-rasi/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-rasi -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-rasi/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_rasi

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 1
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXTERNAL_TOKEN_COUNT = 0
const FIELD_COUNT = 17
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
const LANGUAGE_VERSION = 15
const LARGE_STATE_COUNT = 12
const MAX_ALIAS_SEQUENCE_LENGTH = 12
const MAX_RESERVED_WORD_SET_SIZE = 0
const PRODUCTION_ID_COUNT = 35
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 251
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 163
const TOKEN_COUNT = 100
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
	Fabi_version               uint32_t
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
	Fprimary_state_ids          uintptr
	Fname                       uintptr
	Freserved_words             uintptr
	Fmax_reserved_word_set_size uint16_t
	Fsupertype_count            uint32_t
	Fsupertype_symbols          uintptr
	Fsupertype_map_slices       uintptr
	Fsupertype_map_entries      uintptr
	Fmetadata                   TSLanguageMetadata
}

type TSLanguageMetadata = struct {
	Fmajor_version uint8_t
	Fminor_version uint8_t
	Fpatch_version uint8_t
}

type TSFieldMapEntry = struct {
	Ffield_id    TSFieldId
	Fchild_index uint8_t
	Finherited   uint8
}

type TSMapSlice = struct {
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

type TSLexerMode = struct {
	Flex_state            uint16_t
	Fexternal_lex_state   uint16_t
	Freserved_word_set_id uint16_t
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

const sym_identifier = 1
const anon_sym_ATimport = 2
const anon_sym_ATtheme = 3
const anon_sym_ATmedia = 4
const anon_sym_COMMA = 5
const anon_sym_LBRACE = 6
const anon_sym_RBRACE = 7
const anon_sym_STAR = 8
const anon_sym_POUND = 9
const anon_sym_DOT = 10
const anon_sym_normal = 11
const anon_sym_selected = 12
const anon_sym_alternate = 13
const anon_sym_urgent = 14
const anon_sym_active = 15
const anon_sym_COLON = 16
const anon_sym_SEMI = 17
const anon_sym_LPAREN = 18
const anon_sym_RPAREN = 19
const anon_sym_min_DASHwidth = 20
const anon_sym_max_DASHwidth = 21
const anon_sym_min_DASHheight = 22
const anon_sym_max_DASHheight = 23
const anon_sym_min_DASHaspect_DASHratio = 24
const anon_sym_max_DASHaspect_DASHratio = 25
const anon_sym_monitor_DASHid = 26
const anon_sym_enabled = 27
const anon_sym_inherit = 28
const anon_sym_DMENU = 29
const anon_sym_DQUOTE = 30
const aux_sym_string_value_token1 = 31
const sym_integer_value = 32
const sym_float_value = 33
const anon_sym_true = 34
const anon_sym_false = 35
const anon_sym_url = 36
const anon_sym_none = 37
const anon_sym_both = 38
const anon_sym_width = 39
const anon_sym_height = 40
const anon_sym_linear_DASHgradient = 41
const anon_sym_to = 42
const anon_sym_top = 43
const anon_sym_left = 44
const anon_sym_right = 45
const anon_sym_bottom = 46
const anon_sym_deg = 47
const anon_sym_rad = 48
const anon_sym_grad = 49
const anon_sym_turn = 50
const aux_sym_hex_color_token1 = 51
const anon_sym_0 = 52
const anon_sym_PERCENT = 53
const anon_sym_rgb = 54
const anon_sym_rgba = 55
const anon_sym_hsl = 56
const anon_sym_hsla = 57
const anon_sym_hwb = 58
const anon_sym_hwba = 59
const anon_sym_cmyk = 60
const anon_sym_SLASH = 61
const anon_sym_bold = 62
const anon_sym_italic = 63
const anon_sym_underline = 64
const anon_sym_strikethrough = 65
const anon_sym_dash = 66
const anon_sym_solid = 67
const anon_sym_calc = 68
const anon_sym_PLUS = 69
const anon_sym_DASH = 70
const anon_sym_modulo = 71
const anon_sym_min = 72
const anon_sym_max = 73
const anon_sym_floor = 74
const anon_sym_ceil = 75
const anon_sym_round = 76
const anon_sym_px = 77
const anon_sym_mm = 78
const anon_sym_cm = 79
const anon_sym_ph = 80
const anon_sym_em = 81
const anon_sym_center = 82
const anon_sym_north = 83
const anon_sym_east = 84
const anon_sym_south = 85
const anon_sym_west = 86
const anon_sym_AT = 87
const anon_sym_var = 88
const anon_sym_horizontal = 89
const anon_sym_vertical = 90
const anon_sym_default = 91
const anon_sym_pointer = 92
const anon_sym_text = 93
const anon_sym_LBRACK = 94
const anon_sym_RBRACK = 95
const anon_sym_DOLLAR = 96
const anon_sym_LBRACE2 = 97
const anon_sym_env = 98
const sym_comment = 99
const sym_stylesheet = 100
const sym_import_statement = 101
const sym_theme_statement = 102
const sym_media_statement = 103
const sym_rule_set = 104
const sym_selectors = 105
const sym_block = 106
const sym__block_item = 107
const sym__selector = 108
const sym_global_selector = 109
const sym_id_selector = 110
const sym_id_selector_view = 111
const sym_id_selector_state = 112
const sym_declaration = 113
const sym__query = 114
const sym_feature_query = 115
const sym_parenthesized_query = 116
const sym_feature_name = 117
const sym__value = 118
const sym_string_value = 119
const sym_boolean_value = 120
const sym_image_value = 121
const sym_url_image = 122
const sym_url_image_scale = 123
const sym_gradient_image = 124
const sym_gradient_image_dir = 125
const sym_angle = 126
const sym_angle_unit = 127
const sym__color_value = 128
const sym_hex_color = 129
const sym_percentage = 130
const sym_rgb_color = 131
const sym_hsl_color = 132
const sym_hwb_color = 133
const sym_cmyk_color = 134
const sym_named_color = 135
const sym_text_style_value = 136
const sym_line_style_value = 137
const sym_distance_value = 138
const sym_distance_calc = 139
const sym_distance_op = 140
const sym_integer_distance_unit = 141
const sym_float_distance_unit = 142
const sym_padding_value = 143
const sym_border_value = 144
const sym_first_border_style = 145
const sym_border_style = 146
const sym_position_value = 147
const sym_reference_value = 148
const sym_orientation_value = 149
const sym_cursor_value = 150
const sym_list_value = 151
const sym_environ_value = 152
const aux_sym_stylesheet_repeat1 = 153
const aux_sym_media_statement_repeat1 = 154
const aux_sym_selectors_repeat1 = 155
const aux_sym_block_repeat1 = 156
const aux_sym_declaration_repeat1 = 157
const aux_sym_feature_query_repeat1 = 158
const aux_sym_gradient_image_repeat1 = 159
const aux_sym_distance_calc_repeat1 = 160
const aux_sym_position_value_repeat1 = 161
const aux_sym_list_value_repeat1 = 162
const alias_sym_property_name = 163

var ts_symbol_names = [164]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 15,
	3:   __ccgo_ts + 23,
	4:   __ccgo_ts + 30,
	5:   __ccgo_ts + 37,
	6:   __ccgo_ts + 39,
	7:   __ccgo_ts + 41,
	8:   __ccgo_ts + 43,
	9:   __ccgo_ts + 45,
	10:  __ccgo_ts + 47,
	11:  __ccgo_ts + 49,
	12:  __ccgo_ts + 56,
	13:  __ccgo_ts + 65,
	14:  __ccgo_ts + 75,
	15:  __ccgo_ts + 82,
	16:  __ccgo_ts + 89,
	17:  __ccgo_ts + 91,
	18:  __ccgo_ts + 93,
	19:  __ccgo_ts + 95,
	20:  __ccgo_ts + 97,
	21:  __ccgo_ts + 107,
	22:  __ccgo_ts + 117,
	23:  __ccgo_ts + 128,
	24:  __ccgo_ts + 139,
	25:  __ccgo_ts + 156,
	26:  __ccgo_ts + 173,
	27:  __ccgo_ts + 184,
	28:  __ccgo_ts + 192,
	29:  __ccgo_ts + 200,
	30:  __ccgo_ts + 206,
	31:  __ccgo_ts + 208,
	32:  __ccgo_ts + 228,
	33:  __ccgo_ts + 242,
	34:  __ccgo_ts + 254,
	35:  __ccgo_ts + 259,
	36:  __ccgo_ts + 265,
	37:  __ccgo_ts + 269,
	38:  __ccgo_ts + 274,
	39:  __ccgo_ts + 279,
	40:  __ccgo_ts + 285,
	41:  __ccgo_ts + 292,
	42:  __ccgo_ts + 308,
	43:  __ccgo_ts + 311,
	44:  __ccgo_ts + 315,
	45:  __ccgo_ts + 320,
	46:  __ccgo_ts + 326,
	47:  __ccgo_ts + 333,
	48:  __ccgo_ts + 337,
	49:  __ccgo_ts + 341,
	50:  __ccgo_ts + 346,
	51:  __ccgo_ts + 351,
	52:  __ccgo_ts + 368,
	53:  __ccgo_ts + 370,
	54:  __ccgo_ts + 372,
	55:  __ccgo_ts + 376,
	56:  __ccgo_ts + 381,
	57:  __ccgo_ts + 385,
	58:  __ccgo_ts + 390,
	59:  __ccgo_ts + 394,
	60:  __ccgo_ts + 399,
	61:  __ccgo_ts + 404,
	62:  __ccgo_ts + 406,
	63:  __ccgo_ts + 411,
	64:  __ccgo_ts + 418,
	65:  __ccgo_ts + 428,
	66:  __ccgo_ts + 442,
	67:  __ccgo_ts + 447,
	68:  __ccgo_ts + 453,
	69:  __ccgo_ts + 458,
	70:  __ccgo_ts + 460,
	71:  __ccgo_ts + 462,
	72:  __ccgo_ts + 469,
	73:  __ccgo_ts + 473,
	74:  __ccgo_ts + 477,
	75:  __ccgo_ts + 483,
	76:  __ccgo_ts + 488,
	77:  __ccgo_ts + 494,
	78:  __ccgo_ts + 497,
	79:  __ccgo_ts + 500,
	80:  __ccgo_ts + 503,
	81:  __ccgo_ts + 506,
	82:  __ccgo_ts + 509,
	83:  __ccgo_ts + 516,
	84:  __ccgo_ts + 522,
	85:  __ccgo_ts + 527,
	86:  __ccgo_ts + 533,
	87:  __ccgo_ts + 538,
	88:  __ccgo_ts + 540,
	89:  __ccgo_ts + 544,
	90:  __ccgo_ts + 555,
	91:  __ccgo_ts + 564,
	92:  __ccgo_ts + 572,
	93:  __ccgo_ts + 580,
	94:  __ccgo_ts + 585,
	95:  __ccgo_ts + 587,
	96:  __ccgo_ts + 589,
	97:  __ccgo_ts + 39,
	98:  __ccgo_ts + 591,
	99:  __ccgo_ts + 595,
	100: __ccgo_ts + 603,
	101: __ccgo_ts + 614,
	102: __ccgo_ts + 631,
	103: __ccgo_ts + 647,
	104: __ccgo_ts + 663,
	105: __ccgo_ts + 672,
	106: __ccgo_ts + 682,
	107: __ccgo_ts + 688,
	108: __ccgo_ts + 700,
	109: __ccgo_ts + 710,
	110: __ccgo_ts + 726,
	111: __ccgo_ts + 738,
	112: __ccgo_ts + 755,
	113: __ccgo_ts + 773,
	114: __ccgo_ts + 785,
	115: __ccgo_ts + 792,
	116: __ccgo_ts + 806,
	117: __ccgo_ts + 826,
	118: __ccgo_ts + 839,
	119: __ccgo_ts + 846,
	120: __ccgo_ts + 859,
	121: __ccgo_ts + 873,
	122: __ccgo_ts + 885,
	123: __ccgo_ts + 895,
	124: __ccgo_ts + 911,
	125: __ccgo_ts + 926,
	126: __ccgo_ts + 936,
	127: __ccgo_ts + 942,
	128: __ccgo_ts + 953,
	129: __ccgo_ts + 966,
	130: __ccgo_ts + 976,
	131: __ccgo_ts + 987,
	132: __ccgo_ts + 997,
	133: __ccgo_ts + 1007,
	134: __ccgo_ts + 1017,
	135: __ccgo_ts + 1028,
	136: __ccgo_ts + 1040,
	137: __ccgo_ts + 1057,
	138: __ccgo_ts + 1074,
	139: __ccgo_ts + 1089,
	140: __ccgo_ts + 1103,
	141: __ccgo_ts + 1115,
	142: __ccgo_ts + 1137,
	143: __ccgo_ts + 1157,
	144: __ccgo_ts + 1171,
	145: __ccgo_ts + 1184,
	146: __ccgo_ts + 1203,
	147: __ccgo_ts + 1216,
	148: __ccgo_ts + 1231,
	149: __ccgo_ts + 1247,
	150: __ccgo_ts + 1265,
	151: __ccgo_ts + 1278,
	152: __ccgo_ts + 1289,
	153: __ccgo_ts + 1303,
	154: __ccgo_ts + 1322,
	155: __ccgo_ts + 1346,
	156: __ccgo_ts + 1364,
	157: __ccgo_ts + 1378,
	158: __ccgo_ts + 1398,
	159: __ccgo_ts + 1420,
	160: __ccgo_ts + 1443,
	161: __ccgo_ts + 1465,
	162: __ccgo_ts + 1488,
	163: __ccgo_ts + 1507,
}

var ts_symbol_map = [164]TSSymbol{
	1:   uint16(sym_identifier),
	2:   uint16(anon_sym_ATimport),
	3:   uint16(anon_sym_ATtheme),
	4:   uint16(anon_sym_ATmedia),
	5:   uint16(anon_sym_COMMA),
	6:   uint16(anon_sym_LBRACE),
	7:   uint16(anon_sym_RBRACE),
	8:   uint16(anon_sym_STAR),
	9:   uint16(anon_sym_POUND),
	10:  uint16(anon_sym_DOT),
	11:  uint16(anon_sym_normal),
	12:  uint16(anon_sym_selected),
	13:  uint16(anon_sym_alternate),
	14:  uint16(anon_sym_urgent),
	15:  uint16(anon_sym_active),
	16:  uint16(anon_sym_COLON),
	17:  uint16(anon_sym_SEMI),
	18:  uint16(anon_sym_LPAREN),
	19:  uint16(anon_sym_RPAREN),
	20:  uint16(anon_sym_min_DASHwidth),
	21:  uint16(anon_sym_max_DASHwidth),
	22:  uint16(anon_sym_min_DASHheight),
	23:  uint16(anon_sym_max_DASHheight),
	24:  uint16(anon_sym_min_DASHaspect_DASHratio),
	25:  uint16(anon_sym_max_DASHaspect_DASHratio),
	26:  uint16(anon_sym_monitor_DASHid),
	27:  uint16(anon_sym_enabled),
	28:  uint16(anon_sym_inherit),
	29:  uint16(anon_sym_DMENU),
	30:  uint16(anon_sym_DQUOTE),
	31:  uint16(aux_sym_string_value_token1),
	32:  uint16(sym_integer_value),
	33:  uint16(sym_float_value),
	34:  uint16(anon_sym_true),
	35:  uint16(anon_sym_false),
	36:  uint16(anon_sym_url),
	37:  uint16(anon_sym_none),
	38:  uint16(anon_sym_both),
	39:  uint16(anon_sym_width),
	40:  uint16(anon_sym_height),
	41:  uint16(anon_sym_linear_DASHgradient),
	42:  uint16(anon_sym_to),
	43:  uint16(anon_sym_top),
	44:  uint16(anon_sym_left),
	45:  uint16(anon_sym_right),
	46:  uint16(anon_sym_bottom),
	47:  uint16(anon_sym_deg),
	48:  uint16(anon_sym_rad),
	49:  uint16(anon_sym_grad),
	50:  uint16(anon_sym_turn),
	51:  uint16(aux_sym_hex_color_token1),
	52:  uint16(anon_sym_0),
	53:  uint16(anon_sym_PERCENT),
	54:  uint16(anon_sym_rgb),
	55:  uint16(anon_sym_rgba),
	56:  uint16(anon_sym_hsl),
	57:  uint16(anon_sym_hsla),
	58:  uint16(anon_sym_hwb),
	59:  uint16(anon_sym_hwba),
	60:  uint16(anon_sym_cmyk),
	61:  uint16(anon_sym_SLASH),
	62:  uint16(anon_sym_bold),
	63:  uint16(anon_sym_italic),
	64:  uint16(anon_sym_underline),
	65:  uint16(anon_sym_strikethrough),
	66:  uint16(anon_sym_dash),
	67:  uint16(anon_sym_solid),
	68:  uint16(anon_sym_calc),
	69:  uint16(anon_sym_PLUS),
	70:  uint16(anon_sym_DASH),
	71:  uint16(anon_sym_modulo),
	72:  uint16(anon_sym_min),
	73:  uint16(anon_sym_max),
	74:  uint16(anon_sym_floor),
	75:  uint16(anon_sym_ceil),
	76:  uint16(anon_sym_round),
	77:  uint16(anon_sym_px),
	78:  uint16(anon_sym_mm),
	79:  uint16(anon_sym_cm),
	80:  uint16(anon_sym_ph),
	81:  uint16(anon_sym_em),
	82:  uint16(anon_sym_center),
	83:  uint16(anon_sym_north),
	84:  uint16(anon_sym_east),
	85:  uint16(anon_sym_south),
	86:  uint16(anon_sym_west),
	87:  uint16(anon_sym_AT),
	88:  uint16(anon_sym_var),
	89:  uint16(anon_sym_horizontal),
	90:  uint16(anon_sym_vertical),
	91:  uint16(anon_sym_default),
	92:  uint16(anon_sym_pointer),
	93:  uint16(anon_sym_text),
	94:  uint16(anon_sym_LBRACK),
	95:  uint16(anon_sym_RBRACK),
	96:  uint16(anon_sym_DOLLAR),
	97:  uint16(anon_sym_LBRACE),
	98:  uint16(anon_sym_env),
	99:  uint16(sym_comment),
	100: uint16(sym_stylesheet),
	101: uint16(sym_import_statement),
	102: uint16(sym_theme_statement),
	103: uint16(sym_media_statement),
	104: uint16(sym_rule_set),
	105: uint16(sym_selectors),
	106: uint16(sym_block),
	107: uint16(sym__block_item),
	108: uint16(sym__selector),
	109: uint16(sym_global_selector),
	110: uint16(sym_id_selector),
	111: uint16(sym_id_selector_view),
	112: uint16(sym_id_selector_state),
	113: uint16(sym_declaration),
	114: uint16(sym__query),
	115: uint16(sym_feature_query),
	116: uint16(sym_parenthesized_query),
	117: uint16(sym_feature_name),
	118: uint16(sym__value),
	119: uint16(sym_string_value),
	120: uint16(sym_boolean_value),
	121: uint16(sym_image_value),
	122: uint16(sym_url_image),
	123: uint16(sym_url_image_scale),
	124: uint16(sym_gradient_image),
	125: uint16(sym_gradient_image_dir),
	126: uint16(sym_angle),
	127: uint16(sym_angle_unit),
	128: uint16(sym__color_value),
	129: uint16(sym_hex_color),
	130: uint16(sym_percentage),
	131: uint16(sym_rgb_color),
	132: uint16(sym_hsl_color),
	133: uint16(sym_hwb_color),
	134: uint16(sym_cmyk_color),
	135: uint16(sym_named_color),
	136: uint16(sym_text_style_value),
	137: uint16(sym_line_style_value),
	138: uint16(sym_distance_value),
	139: uint16(sym_distance_calc),
	140: uint16(sym_distance_op),
	141: uint16(sym_integer_distance_unit),
	142: uint16(sym_float_distance_unit),
	143: uint16(sym_padding_value),
	144: uint16(sym_border_value),
	145: uint16(sym_first_border_style),
	146: uint16(sym_border_style),
	147: uint16(sym_position_value),
	148: uint16(sym_reference_value),
	149: uint16(sym_orientation_value),
	150: uint16(sym_cursor_value),
	151: uint16(sym_list_value),
	152: uint16(sym_environ_value),
	153: uint16(aux_sym_stylesheet_repeat1),
	154: uint16(aux_sym_media_statement_repeat1),
	155: uint16(aux_sym_selectors_repeat1),
	156: uint16(aux_sym_block_repeat1),
	157: uint16(aux_sym_declaration_repeat1),
	158: uint16(aux_sym_feature_query_repeat1),
	159: uint16(aux_sym_gradient_image_repeat1),
	160: uint16(aux_sym_distance_calc_repeat1),
	161: uint16(aux_sym_position_value_repeat1),
	162: uint16(aux_sym_list_value_repeat1),
	163: uint16(alias_sym_property_name),
}

var ts_symbol_metadata = [164]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	31: {},
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
	51: {},
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
	},
	60: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	61: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	62: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	63: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	64: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	65: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	66: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	67: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	68: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	69: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	70: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	71: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	72: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	73: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	74: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	75: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	76: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	77: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	78: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	79: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	80: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	81: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	82: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	83: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	84: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	85: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	86: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	87: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	88: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	89: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	90: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	91: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	92: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	93: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	94: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	95: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	96: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	97: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	98: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	108: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	129: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	130: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	131: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	132: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	133: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	134: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	135: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	136: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	137: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	138: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	139: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	140: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	141: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	142: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	143: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	144: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	145: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	146: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	147: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	148: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	149: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	150: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	151: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	152: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	153: {},
	154: {},
	155: {},
	156: {},
	157: {},
	158: {},
	159: {},
	160: {},
	161: {},
	162: {},
	163: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
}

type ts_field_identifiers = int32

const field_body = 1
const field_conditions = 2
const field_file = 3
const field_filename = 4
const field_key = 5
const field_name = 6
const field_op = 7
const field_opacity = 8
const field_scale = 9
const field_selectors = 10
const field_size = 11
const field_state = 12
const field_style = 13
const field_unit = 14
const field_value = 15
const field_view = 16
const field_widget = 17

var ts_field_names = [18]uintptr{
	0:  libc.UintptrFromInt32(0),
	1:  __ccgo_ts + 1521,
	2:  __ccgo_ts + 1526,
	3:  __ccgo_ts + 1537,
	4:  __ccgo_ts + 1542,
	5:  __ccgo_ts + 1551,
	6:  __ccgo_ts + 1555,
	7:  __ccgo_ts + 1560,
	8:  __ccgo_ts + 1563,
	9:  __ccgo_ts + 1571,
	10: __ccgo_ts + 672,
	11: __ccgo_ts + 1577,
	12: __ccgo_ts + 1582,
	13: __ccgo_ts + 1588,
	14: __ccgo_ts + 1594,
	15: __ccgo_ts + 1599,
	16: __ccgo_ts + 1605,
	17: __ccgo_ts + 1610,
}

var ts_field_map_slices = [35]TSMapSlice{
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
		Flength: uint16(2),
	},
	6: {
		Findex:  uint16(6),
		Flength: uint16(2),
	},
	7: {
		Findex:  uint16(8),
		Flength: uint16(2),
	},
	8: {
		Findex:  uint16(10),
		Flength: uint16(2),
	},
	9: {
		Findex:  uint16(12),
		Flength: uint16(1),
	},
	10: {
		Findex:  uint16(13),
		Flength: uint16(1),
	},
	11: {
		Findex:  uint16(14),
		Flength: uint16(2),
	},
	12: {
		Findex:  uint16(16),
		Flength: uint16(3),
	},
	13: {
		Findex:  uint16(19),
		Flength: uint16(2),
	},
	14: {
		Findex:  uint16(21),
		Flength: uint16(1),
	},
	15: {
		Findex:  uint16(22),
		Flength: uint16(1),
	},
	16: {
		Findex:  uint16(23),
		Flength: uint16(2),
	},
	17: {
		Findex:  uint16(25),
		Flength: uint16(1),
	},
	18: {
		Findex:  uint16(26),
		Flength: uint16(2),
	},
	19: {
		Findex:  uint16(28),
		Flength: uint16(1),
	},
	20: {
		Findex:  uint16(29),
		Flength: uint16(2),
	},
	21: {
		Findex:  uint16(31),
		Flength: uint16(1),
	},
	22: {
		Findex:  uint16(32),
		Flength: uint16(1),
	},
	23: {
		Findex:  uint16(33),
		Flength: uint16(1),
	},
	24: {
		Findex:  uint16(34),
		Flength: uint16(3),
	},
	25: {
		Findex:  uint16(37),
		Flength: uint16(2),
	},
	26: {
		Findex:  uint16(39),
		Flength: uint16(2),
	},
	27: {
		Findex:  uint16(41),
		Flength: uint16(1),
	},
	28: {
		Findex:  uint16(42),
		Flength: uint16(1),
	},
	29: {
		Findex:  uint16(43),
		Flength: uint16(2),
	},
	30: {
		Findex:  uint16(45),
		Flength: uint16(1),
	},
	31: {
		Findex:  uint16(46),
		Flength: uint16(1),
	},
	32: {
		Findex:  uint16(47),
		Flength: uint16(2),
	},
	33: {
		Findex:  uint16(49),
		Flength: uint16(2),
	},
	34: {
		Findex:  uint16(51),
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [53]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_widget),
	},
	1: {
		Ffield_id:    uint16(field_file),
		Fchild_index: uint8(1),
	},
	2: {
		Ffield_id:    uint16(field_widget),
		Fchild_index: uint8(1),
	},
	3: {
		Ffield_id: uint16(field_view),
	},
	4: {
		Ffield_id:    uint16(field_view),
		Fchild_index: uint8(1),
	},
	5: {
		Ffield_id: uint16(field_widget),
	},
	6: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(1),
	},
	7: {
		Ffield_id: uint16(field_selectors),
	},
	8: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(2),
	},
	9: {
		Ffield_id:    uint16(field_conditions),
		Fchild_index: uint8(1),
	},
	10: {
		Ffield_id:    uint16(field_view),
		Fchild_index: uint8(2),
	},
	11: {
		Ffield_id:    uint16(field_widget),
		Fchild_index: uint8(1),
	},
	12: {
		Ffield_id:    uint16(field_view),
		Fchild_index: uint8(1),
	},
	13: {
		Ffield_id: uint16(field_state),
	},
	14: {
		Ffield_id:    uint16(field_state),
		Fchild_index: uint8(1),
	},
	15: {
		Ffield_id: uint16(field_view),
	},
	16: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(3),
	},
	17: {
		Ffield_id:    uint16(field_conditions),
		Fchild_index: uint8(1),
	},
	18: {
		Ffield_id:    uint16(field_conditions),
		Fchild_index: uint8(2),
	},
	19: {
		Ffield_id:    uint16(field_state),
		Fchild_index: uint8(2),
	},
	20: {
		Ffield_id:    uint16(field_view),
		Fchild_index: uint8(1),
	},
	21: {
		Ffield_id:    uint16(field_state),
		Fchild_index: uint8(1),
	},
	22: {
		Ffield_id: uint16(field_name),
	},
	23: {
		Ffield_id:    uint16(field_unit),
		Fchild_index: uint8(1),
	},
	24: {
		Ffield_id: uint16(field_value),
	},
	25: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	26: {
		Ffield_id: uint16(field_size),
	},
	27: {
		Ffield_id:    uint16(field_style),
		Fchild_index: uint8(1),
	},
	28: {
		Ffield_id: uint16(field_size),
	},
	29: {
		Ffield_id:    uint16(field_key),
		Fchild_index: uint8(1),
	},
	30: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(3),
	},
	31: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	32: {
		Ffield_id: uint16(field_value),
	},
	33: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
	},
	34: {
		Ffield_id: uint16(field_name),
	},
	35: {
		Ffield_id:    uint16(field_opacity),
		Fchild_index: uint8(1),
	},
	36: {
		Ffield_id:    uint16(field_opacity),
		Fchild_index: uint8(2),
	},
	37: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	38: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	39: {
		Ffield_id:  uint16(field_value),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	40: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	41: {
		Ffield_id:    uint16(field_filename),
		Fchild_index: uint8(2),
	},
	42: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	43: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
	},
	44: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	45: {
		Ffield_id: uint16(field_op),
	},
	46: {
		Ffield_id:    uint16(field_op),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	47: {
		Ffield_id:  uint16(field_op),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	48: {
		Ffield_id:    uint16(field_op),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	49: {
		Ffield_id:    uint16(field_filename),
		Fchild_index: uint8(2),
	},
	50: {
		Ffield_id:    uint16(field_scale),
		Fchild_index: uint8(4),
	},
	51: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	52: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(4),
	},
}

var ts_alias_sequences = [35][12]TSSymbol{
	0: {},
	21: {
		0: uint16(alias_sym_property_name),
	},
	25: {
		0: uint16(alias_sym_property_name),
	},
}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [251]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(3),
	4:   uint16(4),
	5:   uint16(5),
	6:   uint16(6),
	7:   uint16(7),
	8:   uint16(8),
	9:   uint16(9),
	10:  uint16(10),
	11:  uint16(11),
	12:  uint16(12),
	13:  uint16(13),
	14:  uint16(14),
	15:  uint16(15),
	16:  uint16(16),
	17:  uint16(17),
	18:  uint16(18),
	19:  uint16(19),
	20:  uint16(20),
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
	39:  uint16(39),
	40:  uint16(40),
	41:  uint16(41),
	42:  uint16(42),
	43:  uint16(43),
	44:  uint16(44),
	45:  uint16(45),
	46:  uint16(46),
	47:  uint16(47),
	48:  uint16(48),
	49:  uint16(49),
	50:  uint16(50),
	51:  uint16(51),
	52:  uint16(52),
	53:  uint16(53),
	54:  uint16(54),
	55:  uint16(55),
	56:  uint16(56),
	57:  uint16(57),
	58:  uint16(58),
	59:  uint16(59),
	60:  uint16(60),
	61:  uint16(61),
	62:  uint16(62),
	63:  uint16(63),
	64:  uint16(64),
	65:  uint16(65),
	66:  uint16(66),
	67:  uint16(67),
	68:  uint16(68),
	69:  uint16(69),
	70:  uint16(70),
	71:  uint16(71),
	72:  uint16(72),
	73:  uint16(13),
	74:  uint16(74),
	75:  uint16(75),
	76:  uint16(76),
	77:  uint16(77),
	78:  uint16(78),
	79:  uint16(79),
	80:  uint16(80),
	81:  uint16(81),
	82:  uint16(13),
	83:  uint16(83),
	84:  uint16(84),
	85:  uint16(85),
	86:  uint16(86),
	87:  uint16(83),
	88:  uint16(86),
	89:  uint16(89),
	90:  uint16(90),
	91:  uint16(91),
	92:  uint16(51),
	93:  uint16(37),
	94:  uint16(94),
	95:  uint16(27),
	96:  uint16(35),
	97:  uint16(43),
	98:  uint16(38),
	99:  uint16(45),
	100: uint16(15),
	101: uint16(68),
	102: uint16(12),
	103: uint16(16),
	104: uint16(104),
	105: uint16(105),
	106: uint16(106),
	107: uint16(107),
	108: uint16(106),
	109: uint16(17),
	110: uint16(18),
	111: uint16(111),
	112: uint16(112),
	113: uint16(113),
	114: uint16(114),
	115: uint16(115),
	116: uint16(116),
	117: uint16(117),
	118: uint16(118),
	119: uint16(19),
	120: uint16(120),
	121: uint16(20),
	122: uint16(122),
	123: uint16(123),
	124: uint16(14),
	125: uint16(125),
	126: uint16(125),
	127: uint16(127),
	128: uint16(128),
	129: uint16(129),
	130: uint16(130),
	131: uint16(130),
	132: uint16(132),
	133: uint16(133),
	134: uint16(134),
	135: uint16(135),
	136: uint16(136),
	137: uint16(137),
	138: uint16(138),
	139: uint16(139),
	140: uint16(140),
	141: uint16(141),
	142: uint16(142),
	143: uint16(143),
	144: uint16(144),
	145: uint16(145),
	146: uint16(146),
	147: uint16(147),
	148: uint16(148),
	149: uint16(149),
	150: uint16(150),
	151: uint16(151),
	152: uint16(152),
	153: uint16(153),
	154: uint16(154),
	155: uint16(155),
	156: uint16(156),
	157: uint16(157),
	158: uint16(158),
	159: uint16(159),
	160: uint16(160),
	161: uint16(161),
	162: uint16(162),
	163: uint16(163),
	164: uint16(164),
	165: uint16(165),
	166: uint16(166),
	167: uint16(167),
	168: uint16(168),
	169: uint16(169),
	170: uint16(170),
	171: uint16(171),
	172: uint16(172),
	173: uint16(173),
	174: uint16(174),
	175: uint16(175),
	176: uint16(176),
	177: uint16(177),
	178: uint16(178),
	179: uint16(179),
	180: uint16(180),
	181: uint16(181),
	182: uint16(182),
	183: uint16(183),
	184: uint16(184),
	185: uint16(185),
	186: uint16(186),
	187: uint16(187),
	188: uint16(188),
	189: uint16(189),
	190: uint16(190),
	191: uint16(191),
	192: uint16(192),
	193: uint16(193),
	194: uint16(194),
	195: uint16(195),
	196: uint16(196),
	197: uint16(197),
	198: uint16(198),
	199: uint16(199),
	200: uint16(200),
	201: uint16(201),
	202: uint16(202),
	203: uint16(203),
	204: uint16(204),
	205: uint16(205),
	206: uint16(206),
	207: uint16(207),
	208: uint16(208),
	209: uint16(209),
	210: uint16(210),
	211: uint16(211),
	212: uint16(212),
	213: uint16(213),
	214: uint16(214),
	215: uint16(215),
	216: uint16(216),
	217: uint16(217),
	218: uint16(218),
	219: uint16(219),
	220: uint16(220),
	221: uint16(221),
	222: uint16(222),
	223: uint16(223),
	224: uint16(224),
	225: uint16(225),
	226: uint16(226),
	227: uint16(227),
	228: uint16(228),
	229: uint16(229),
	230: uint16(230),
	231: uint16(231),
	232: uint16(232),
	233: uint16(233),
	234: uint16(234),
	235: uint16(235),
	236: uint16(236),
	237: uint16(237),
	238: uint16(238),
	239: uint16(239),
	240: uint16(240),
	241: uint16(241),
	242: uint16(242),
	243: uint16(243),
	244: uint16(224),
	245: uint16(245),
	246: uint16(246),
	247: uint16(242),
	248: uint16(248),
	249: uint16(205),
	250: uint16(250),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2, i3, i4 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _, _ = eof, i, i1, i2, i3, i4, lookahead, result, skip
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
			state = uint16(32)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(80)/libc.Uint64FromInt64(2)) {
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
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\r') {
			state = uint16(79)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(2):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(58)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(3):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token2[i2]) == lookahead {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _3
		_3:
			;
			i2 = i2 + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(58)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('*') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('*') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(77)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('*') {
			state = uint16(5)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('.') {
			state = uint16(25)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('/') {
			state = uint16(4)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('/') {
			state = uint16(4)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('a') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('d') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('e') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('e') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('e') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('h') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('i') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('i') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('m') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('m') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('o') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('p') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('r') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('t') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(24):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(25):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(26):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(27):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(28):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(29):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(30):
		if eof != 0 {
			state = uint16(32)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(80)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token3[i3]) == lookahead {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _4
		_4:
			;
			i3 = i3 + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(30)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(31):
		if eof != 0 {
			state = uint16(32)
			goto next_state
		}
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token4[i4]) == lookahead {
				state = map_token4[i4+uint32(1)]
				goto next_state
			}
			goto _5
		_5:
			;
			i4 = i4 + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(31)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(59)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATimport)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATtheme)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATmedia)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(48)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\r') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(47)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') {
			state = uint16(78)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(48)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(56)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('"') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(53)
			goto next_state
		}
		if lookahead == int32('\n') || lookahead == int32('"') {
			state = uint16(6)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(53)
			goto next_state
		}
		if lookahead == int32('\n') || lookahead == int32('"') {
			state = uint16(6)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(53)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(54)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32('"') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\\') {
			state = uint16(56)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('"') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\\') {
			state = uint16(56)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\\') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('\n') || lookahead == int32('"') {
			state = uint16(78)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(25)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_color_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_color_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_color_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_color_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(25)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\\') {
			state = uint16(1)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\\') {
			state = uint16(1)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(78)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [40]uint16_t{
	0:  uint16('"'),
	1:  uint16(46),
	2:  uint16('#'),
	3:  uint16(40),
	4:  uint16('$'),
	5:  uint16(74),
	6:  uint16('%'),
	7:  uint16(67),
	8:  uint16('('),
	9:  uint16(44),
	10: uint16(')'),
	11: uint16(45),
	12: uint16('*'),
	13: uint16(39),
	14: uint16('+'),
	15: uint16(69),
	16: uint16(','),
	17: uint16(36),
	18: uint16('-'),
	19: uint16(70),
	20: uint16('.'),
	21: uint16(41),
	22: uint16('/'),
	23: uint16(68),
	24: uint16('0'),
	25: uint16(65),
	26: uint16(':'),
	27: uint16(42),
	28: uint16(';'),
	29: uint16(43),
	30: uint16('@'),
	31: uint16(71),
	32: uint16('['),
	33: uint16(72),
	34: uint16(']'),
	35: uint16(73),
	36: uint16('{'),
	37: uint16(75),
	38: uint16('}'),
	39: uint16(38),
}

var map_token1 = [30]uint16_t{
	0:  uint16('"'),
	1:  uint16(46),
	2:  uint16('#'),
	3:  uint16(40),
	4:  uint16('$'),
	5:  uint16(74),
	6:  uint16('%'),
	7:  uint16(67),
	8:  uint16(')'),
	9:  uint16(45),
	10: uint16(','),
	11: uint16(36),
	12: uint16('.'),
	13: uint16(25),
	14: uint16('/'),
	15: uint16(68),
	16: uint16('0'),
	17: uint16(66),
	18: uint16(';'),
	19: uint16(43),
	20: uint16('@'),
	21: uint16(71),
	22: uint16('['),
	23: uint16(72),
	24: uint16('{'),
	25: uint16(75),
	26: uint16('+'),
	27: uint16(7),
	28: uint16('-'),
	29: uint16(7),
}

var map_token2 = [28]uint16_t{
	0:  uint16('"'),
	1:  uint16(46),
	2:  uint16('#'),
	3:  uint16(40),
	4:  uint16('$'),
	5:  uint16(74),
	6:  uint16('%'),
	7:  uint16(67),
	8:  uint16(')'),
	9:  uint16(45),
	10: uint16(','),
	11: uint16(36),
	12: uint16('.'),
	13: uint16(25),
	14: uint16('/'),
	15: uint16(68),
	16: uint16('0'),
	17: uint16(66),
	18: uint16(';'),
	19: uint16(43),
	20: uint16('@'),
	21: uint16(71),
	22: uint16('['),
	23: uint16(72),
	24: uint16('+'),
	25: uint16(7),
	26: uint16('-'),
	27: uint16(7),
}

var map_token3 = [40]uint16_t{
	0:  uint16('"'),
	1:  uint16(46),
	2:  uint16('#'),
	3:  uint16(40),
	4:  uint16('$'),
	5:  uint16(74),
	6:  uint16('%'),
	7:  uint16(67),
	8:  uint16('('),
	9:  uint16(44),
	10: uint16(')'),
	11: uint16(45),
	12: uint16('*'),
	13: uint16(39),
	14: uint16('+'),
	15: uint16(69),
	16: uint16(','),
	17: uint16(36),
	18: uint16('-'),
	19: uint16(70),
	20: uint16('.'),
	21: uint16(41),
	22: uint16('/'),
	23: uint16(68),
	24: uint16('0'),
	25: uint16(65),
	26: uint16(':'),
	27: uint16(42),
	28: uint16(';'),
	29: uint16(43),
	30: uint16('@'),
	31: uint16(71),
	32: uint16('['),
	33: uint16(72),
	34: uint16(']'),
	35: uint16(73),
	36: uint16('{'),
	37: uint16(37),
	38: uint16('}'),
	39: uint16(38),
}

var map_token4 = [24]uint16_t{
	0:  uint16('#'),
	1:  uint16(40),
	2:  uint16(')'),
	3:  uint16(45),
	4:  uint16('*'),
	5:  uint16(39),
	6:  uint16(','),
	7:  uint16(36),
	8:  uint16('.'),
	9:  uint16(41),
	10: uint16('/'),
	11: uint16(4),
	12: uint16(':'),
	13: uint16(42),
	14: uint16('@'),
	15: uint16(16),
	16: uint16('{'),
	17: uint16(37),
	18: uint16('}'),
	19: uint16(38),
	20: uint16('+'),
	21: uint16(24),
	22: uint16('-'),
	23: uint16(24),
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
			if !(uint64(i) < libc.Uint64FromInt64(80)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token5[i]) == lookahead {
				state = map_token5[i+uint32(1)]
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
		if lookahead == int32('M') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('c') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('o') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('a') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(26)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('a') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('a') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('a') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('r') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('e') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('n') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('e') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('a') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('o') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('h') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('a') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('g') {
			state = uint16(53)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('e') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('e') {
			state = uint16(59)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('n') {
			state = uint16(63)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('a') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('e') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('E') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('t') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('t') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('l') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('l') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('i') {
			state = uint16(75)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_cm)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('y') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('s') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('f') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('g') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('s') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_em)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		if lookahead == int32('a') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('v') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('l') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('o') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('a') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('i') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('r') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('l') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('b') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('h') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('a') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('f') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('n') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('x') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('n') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mm)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		if lookahead == int32('d') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('n') {
			state = uint16(99)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ph)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		if lookahead == int32('i') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_px)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(52):
		if lookahead == int32('d') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('b') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('g') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('u') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('l') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('l') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('r') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('x') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_to)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('u') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('r') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('d') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('g') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('r') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('r') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('s') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('d') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('N') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('i') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('e') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('d') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('h') {
			state = uint16(125)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('c') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('l') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('t') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('k') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('h') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead == int32('a') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_deg)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		if lookahead == int32('t') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead == int32('b') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_env)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		if lookahead == int32('s') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('o') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('d') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead == int32('g') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead == int32('i') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_hsl)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_hwb)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('e') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('l') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('t') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead == int32('e') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_max)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_min)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead == int32('u') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead == int32('i') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead == int32('e') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead == int32('m') {
			state = uint16(151)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('n') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rad)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgb)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead == int32('h') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead == int32('n') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead == int32('e') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead == int32('i') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead == int32('t') {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('i') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead == int32('t') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_top)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(112):
		if lookahead == int32('e') {
			state = uint16(162)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead == int32('n') {
			state = uint16(163)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('e') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('e') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_url)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_var)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(118):
		if lookahead == int32('t') {
			state = uint16(166)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('t') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('t') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('U') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('v') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead == int32('r') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bold)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_both)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(126):
		if lookahead == int32('o') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_calc)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ceil)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(129):
		if lookahead == int32('e') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_cmyk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_dash)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(132):
		if lookahead == int32('u') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_east)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(134):
		if lookahead == int32('l') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead == int32('e') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(136):
		if lookahead == int32('r') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_grad)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(138):
		if lookahead == int32('h') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(139):
		if lookahead == int32('z') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(140):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_hsla)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(141):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_hwba)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(142):
		if lookahead == int32('r') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(143):
		if lookahead == int32('i') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(144):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_left)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(145):
		if lookahead == int32('a') {
			state = uint16(182)
			goto next_state
		}
		return result
	case int32(146):
		if lookahead == int32('a') {
			state = uint16(183)
			goto next_state
		}
		if lookahead == int32('h') {
			state = uint16(184)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(147):
		if lookahead == int32('a') {
			state = uint16(186)
			goto next_state
		}
		if lookahead == int32('h') {
			state = uint16(187)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(148):
		if lookahead == int32('l') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(149):
		if lookahead == int32('t') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(150):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_none)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(151):
		if lookahead == int32('a') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(152):
		if lookahead == int32('h') {
			state = uint16(192)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead == int32('t') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(154):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(155):
		if lookahead == int32('t') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(156):
		if lookahead == int32('d') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(157):
		if lookahead == int32('c') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(158):
		if lookahead == int32('d') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(159):
		if lookahead == int32('h') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(160):
		if lookahead == int32('k') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(161):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(162):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(163):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_turn)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(164):
		if lookahead == int32('r') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(165):
		if lookahead == int32('n') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(166):
		if lookahead == int32('i') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(167):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_west)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(168):
		if lookahead == int32('h') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(169):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DMENU)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(170):
		if lookahead == int32('e') {
			state = uint16(204)
			goto next_state
		}
		return result
	case int32(171):
		if lookahead == int32('n') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(172):
		if lookahead == int32('m') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(173):
		if lookahead == int32('r') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(174):
		if lookahead == int32('l') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(175):
		if lookahead == int32('e') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(176):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(177):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_floor)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(178):
		if lookahead == int32('t') {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(179):
		if lookahead == int32('o') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(180):
		if lookahead == int32('i') {
			state = uint16(212)
			goto next_state
		}
		return result
	case int32(181):
		if lookahead == int32('c') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(182):
		if lookahead == int32('r') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(183):
		if lookahead == int32('s') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(184):
		if lookahead == int32('e') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(185):
		if lookahead == int32('i') {
			state = uint16(217)
			goto next_state
		}
		return result
	case int32(186):
		if lookahead == int32('s') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(187):
		if lookahead == int32('e') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(188):
		if lookahead == int32('i') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(189):
		if lookahead == int32('o') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(190):
		if lookahead == int32('o') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(191):
		if lookahead == int32('l') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(192):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_north)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(193):
		if lookahead == int32('e') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(194):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_right)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(195):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_round)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(196):
		if lookahead == int32('t') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(197):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_solid)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(198):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_south)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(199):
		if lookahead == int32('e') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(200):
		if lookahead == int32('l') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(201):
		if lookahead == int32('t') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(202):
		if lookahead == int32('c') {
			state = uint16(229)
			goto next_state
		}
		return result
	case int32(203):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_width)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(204):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_active)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(205):
		if lookahead == int32('a') {
			state = uint16(230)
			goto next_state
		}
		return result
	case int32(206):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bottom)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(207):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_center)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(208):
		if lookahead == int32('t') {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(209):
		if lookahead == int32('d') {
			state = uint16(232)
			goto next_state
		}
		return result
	case int32(210):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_height)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(211):
		if lookahead == int32('n') {
			state = uint16(233)
			goto next_state
		}
		return result
	case int32(212):
		if lookahead == int32('t') {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(213):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_italic)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(214):
		if lookahead == int32('-') {
			state = uint16(235)
			goto next_state
		}
		return result
	case int32(215):
		if lookahead == int32('p') {
			state = uint16(236)
			goto next_state
		}
		return result
	case int32(216):
		if lookahead == int32('i') {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(217):
		if lookahead == int32('d') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(218):
		if lookahead == int32('p') {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(219):
		if lookahead == int32('i') {
			state = uint16(240)
			goto next_state
		}
		return result
	case int32(220):
		if lookahead == int32('d') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(221):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_modulo)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(222):
		if lookahead == int32('r') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(223):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_normal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(224):
		if lookahead == int32('r') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(225):
		if lookahead == int32('e') {
			state = uint16(244)
			goto next_state
		}
		return result
	case int32(226):
		if lookahead == int32('t') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(227):
		if lookahead == int32('i') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(228):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_urgent)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(229):
		if lookahead == int32('a') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(230):
		if lookahead == int32('t') {
			state = uint16(248)
			goto next_state
		}
		return result
	case int32(231):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_default)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(232):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enabled)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(233):
		if lookahead == int32('t') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(234):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_inherit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(235):
		if lookahead == int32('g') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(236):
		if lookahead == int32('e') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(237):
		if lookahead == int32('g') {
			state = uint16(252)
			goto next_state
		}
		return result
	case int32(238):
		if lookahead == int32('t') {
			state = uint16(253)
			goto next_state
		}
		return result
	case int32(239):
		if lookahead == int32('e') {
			state = uint16(254)
			goto next_state
		}
		return result
	case int32(240):
		if lookahead == int32('g') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(241):
		if lookahead == int32('t') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(242):
		if lookahead == int32('-') {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(243):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_pointer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(244):
		if lookahead == int32('d') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(245):
		if lookahead == int32('h') {
			state = uint16(259)
			goto next_state
		}
		return result
	case int32(246):
		if lookahead == int32('n') {
			state = uint16(260)
			goto next_state
		}
		return result
	case int32(247):
		if lookahead == int32('l') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(248):
		if lookahead == int32('e') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(249):
		if lookahead == int32('a') {
			state = uint16(263)
			goto next_state
		}
		return result
	case int32(250):
		if lookahead == int32('r') {
			state = uint16(264)
			goto next_state
		}
		return result
	case int32(251):
		if lookahead == int32('c') {
			state = uint16(265)
			goto next_state
		}
		return result
	case int32(252):
		if lookahead == int32('h') {
			state = uint16(266)
			goto next_state
		}
		return result
	case int32(253):
		if lookahead == int32('h') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(254):
		if lookahead == int32('c') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(255):
		if lookahead == int32('h') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(256):
		if lookahead == int32('h') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(257):
		if lookahead == int32('i') {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(258):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_selected)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(259):
		if lookahead == int32('r') {
			state = uint16(272)
			goto next_state
		}
		return result
	case int32(260):
		if lookahead == int32('e') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(261):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_vertical)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(262):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_alternate)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(263):
		if lookahead == int32('l') {
			state = uint16(274)
			goto next_state
		}
		return result
	case int32(264):
		if lookahead == int32('a') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(265):
		if lookahead == int32('t') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(266):
		if lookahead == int32('t') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(267):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_max_DASHwidth)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(268):
		if lookahead == int32('t') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(269):
		if lookahead == int32('t') {
			state = uint16(279)
			goto next_state
		}
		return result
	case int32(270):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_min_DASHwidth)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(271):
		if lookahead == int32('d') {
			state = uint16(280)
			goto next_state
		}
		return result
	case int32(272):
		if lookahead == int32('o') {
			state = uint16(281)
			goto next_state
		}
		return result
	case int32(273):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_underline)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(274):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_horizontal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(275):
		if lookahead == int32('d') {
			state = uint16(282)
			goto next_state
		}
		return result
	case int32(276):
		if lookahead == int32('-') {
			state = uint16(283)
			goto next_state
		}
		return result
	case int32(277):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_max_DASHheight)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(278):
		if lookahead == int32('-') {
			state = uint16(284)
			goto next_state
		}
		return result
	case int32(279):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_min_DASHheight)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(280):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_monitor_DASHid)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(281):
		if lookahead == int32('u') {
			state = uint16(285)
			goto next_state
		}
		return result
	case int32(282):
		if lookahead == int32('i') {
			state = uint16(286)
			goto next_state
		}
		return result
	case int32(283):
		if lookahead == int32('r') {
			state = uint16(287)
			goto next_state
		}
		return result
	case int32(284):
		if lookahead == int32('r') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(285):
		if lookahead == int32('g') {
			state = uint16(289)
			goto next_state
		}
		return result
	case int32(286):
		if lookahead == int32('e') {
			state = uint16(290)
			goto next_state
		}
		return result
	case int32(287):
		if lookahead == int32('a') {
			state = uint16(291)
			goto next_state
		}
		return result
	case int32(288):
		if lookahead == int32('a') {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(289):
		if lookahead == int32('h') {
			state = uint16(293)
			goto next_state
		}
		return result
	case int32(290):
		if lookahead == int32('n') {
			state = uint16(294)
			goto next_state
		}
		return result
	case int32(291):
		if lookahead == int32('t') {
			state = uint16(295)
			goto next_state
		}
		return result
	case int32(292):
		if lookahead == int32('t') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(293):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_strikethrough)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(294):
		if lookahead == int32('t') {
			state = uint16(297)
			goto next_state
		}
		return result
	case int32(295):
		if lookahead == int32('i') {
			state = uint16(298)
			goto next_state
		}
		return result
	case int32(296):
		if lookahead == int32('i') {
			state = uint16(299)
			goto next_state
		}
		return result
	case int32(297):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_linear_DASHgradient)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(298):
		if lookahead == int32('o') {
			state = uint16(300)
			goto next_state
		}
		return result
	case int32(299):
		if lookahead == int32('o') {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(300):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_max_DASHaspect_DASHratio)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(301):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_min_DASHaspect_DASHratio)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token5 = [40]uint16_t{
	0:  uint16('D'),
	1:  uint16(1),
	2:  uint16('a'),
	3:  uint16(2),
	4:  uint16('b'),
	5:  uint16(3),
	6:  uint16('c'),
	7:  uint16(4),
	8:  uint16('d'),
	9:  uint16(5),
	10: uint16('e'),
	11: uint16(6),
	12: uint16('f'),
	13: uint16(7),
	14: uint16('g'),
	15: uint16(8),
	16: uint16('h'),
	17: uint16(9),
	18: uint16('i'),
	19: uint16(10),
	20: uint16('l'),
	21: uint16(11),
	22: uint16('m'),
	23: uint16(12),
	24: uint16('n'),
	25: uint16(13),
	26: uint16('p'),
	27: uint16(14),
	28: uint16('r'),
	29: uint16(15),
	30: uint16('s'),
	31: uint16(16),
	32: uint16('t'),
	33: uint16(17),
	34: uint16('u'),
	35: uint16(18),
	36: uint16('v'),
	37: uint16(19),
	38: uint16('w'),
	39: uint16(20),
}

var ts_lex_modes = [251]TSLexerMode{
	0: {},
	1: {
		Flex_state: uint16(31),
	},
	2: {
		Flex_state: uint16(2),
	},
	3: {
		Flex_state: uint16(2),
	},
	4: {
		Flex_state: uint16(2),
	},
	5: {
		Flex_state: uint16(2),
	},
	6: {
		Flex_state: uint16(2),
	},
	7: {
		Flex_state: uint16(2),
	},
	8: {
		Flex_state: uint16(2),
	},
	9: {
		Flex_state: uint16(2),
	},
	10: {
		Flex_state: uint16(2),
	},
	11: {
		Flex_state: uint16(2),
	},
	12: {
		Flex_state: uint16(2),
	},
	13: {
		Flex_state: uint16(2),
	},
	14: {
		Flex_state: uint16(2),
	},
	15: {
		Flex_state: uint16(2),
	},
	16: {
		Flex_state: uint16(2),
	},
	17: {
		Flex_state: uint16(2),
	},
	18: {
		Flex_state: uint16(2),
	},
	19: {
		Flex_state: uint16(2),
	},
	20: {
		Flex_state: uint16(2),
	},
	21: {
		Flex_state: uint16(2),
	},
	22: {
		Flex_state: uint16(2),
	},
	23: {
		Flex_state: uint16(2),
	},
	24: {
		Flex_state: uint16(2),
	},
	25: {
		Flex_state: uint16(2),
	},
	26: {
		Flex_state: uint16(2),
	},
	27: {
		Flex_state: uint16(2),
	},
	28: {
		Flex_state: uint16(2),
	},
	29: {
		Flex_state: uint16(2),
	},
	30: {
		Flex_state: uint16(2),
	},
	31: {
		Flex_state: uint16(2),
	},
	32: {
		Flex_state: uint16(2),
	},
	33: {
		Flex_state: uint16(2),
	},
	34: {
		Flex_state: uint16(2),
	},
	35: {
		Flex_state: uint16(2),
	},
	36: {
		Flex_state: uint16(2),
	},
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
	44: {
		Flex_state: uint16(2),
	},
	45: {
		Flex_state: uint16(2),
	},
	46: {
		Flex_state: uint16(2),
	},
	47: {
		Flex_state: uint16(2),
	},
	48: {
		Flex_state: uint16(2),
	},
	49: {
		Flex_state: uint16(2),
	},
	50: {
		Flex_state: uint16(2),
	},
	51: {
		Flex_state: uint16(2),
	},
	52: {
		Flex_state: uint16(2),
	},
	53: {
		Flex_state: uint16(2),
	},
	54: {
		Flex_state: uint16(2),
	},
	55: {
		Flex_state: uint16(2),
	},
	56: {
		Flex_state: uint16(2),
	},
	57: {
		Flex_state: uint16(2),
	},
	58: {
		Flex_state: uint16(2),
	},
	59: {
		Flex_state: uint16(2),
	},
	60: {
		Flex_state: uint16(2),
	},
	61: {
		Flex_state: uint16(2),
	},
	62: {
		Flex_state: uint16(2),
	},
	63: {
		Flex_state: uint16(2),
	},
	64: {
		Flex_state: uint16(2),
	},
	65: {
		Flex_state: uint16(2),
	},
	66: {
		Flex_state: uint16(2),
	},
	67: {
		Flex_state: uint16(2),
	},
	68: {
		Flex_state: uint16(2),
	},
	69: {
		Flex_state: uint16(2),
	},
	70: {
		Flex_state: uint16(2),
	},
	71: {
		Flex_state: uint16(2),
	},
	72: {
		Flex_state: uint16(31),
	},
	73: {},
	74: {},
	75: {},
	76: {
		Flex_state: uint16(31),
	},
	77: {},
	78: {},
	79: {},
	80: {},
	81: {
		Flex_state: uint16(31),
	},
	82: {
		Flex_state: uint16(2),
	},
	83: {},
	84: {},
	85: {},
	86: {},
	87: {},
	88: {},
	89: {},
	90: {},
	91: {},
	92: {},
	93: {},
	94: {},
	95: {},
	96: {},
	97: {},
	98: {},
	99: {},
	100: {
		Flex_state: uint16(2),
	},
	101: {
		Flex_state: uint16(31),
	},
	102: {},
	103: {
		Flex_state: uint16(2),
	},
	104: {
		Flex_state: uint16(31),
	},
	105: {
		Flex_state: uint16(31),
	},
	106: {},
	107: {
		Flex_state: uint16(31),
	},
	108: {},
	109: {
		Flex_state: uint16(2),
	},
	110: {
		Flex_state: uint16(2),
	},
	111: {
		Flex_state: uint16(31),
	},
	112: {
		Flex_state: uint16(31),
	},
	113: {
		Flex_state: uint16(31),
	},
	114: {
		Flex_state: uint16(31),
	},
	115: {
		Flex_state: uint16(31),
	},
	116: {
		Flex_state: uint16(31),
	},
	117: {
		Flex_state: uint16(31),
	},
	118: {
		Flex_state: uint16(31),
	},
	119: {
		Flex_state: uint16(2),
	},
	120: {
		Flex_state: uint16(31),
	},
	121: {
		Flex_state: uint16(2),
	},
	122: {
		Flex_state: uint16(2),
	},
	123: {},
	124: {},
	125: {
		Flex_state: uint16(2),
	},
	126: {
		Flex_state: uint16(2),
	},
	127: {},
	128: {},
	129: {},
	130: {},
	131: {},
	132: {
		Flex_state: uint16(2),
	},
	133: {},
	134: {
		Flex_state: uint16(2),
	},
	135: {
		Flex_state: uint16(2),
	},
	136: {
		Flex_state: uint16(2),
	},
	137: {
		Flex_state: uint16(2),
	},
	138: {
		Flex_state: uint16(31),
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
	142: {},
	143: {
		Flex_state: uint16(2),
	},
	144: {
		Flex_state: uint16(31),
	},
	145: {
		Flex_state: uint16(2),
	},
	146: {},
	147: {
		Flex_state: uint16(2),
	},
	148: {
		Flex_state: uint16(2),
	},
	149: {
		Flex_state: uint16(2),
	},
	150: {},
	151: {
		Flex_state: uint16(2),
	},
	152: {},
	153: {
		Flex_state: uint16(31),
	},
	154: {
		Flex_state: uint16(31),
	},
	155: {},
	156: {},
	157: {},
	158: {},
	159: {},
	160: {},
	161: {},
	162: {
		Flex_state: uint16(31),
	},
	163: {},
	164: {},
	165: {},
	166: {},
	167: {
		Flex_state: uint16(31),
	},
	168: {
		Flex_state: uint16(31),
	},
	169: {
		Flex_state: uint16(31),
	},
	170: {
		Flex_state: uint16(31),
	},
	171: {},
	172: {
		Flex_state: uint16(31),
	},
	173: {},
	174: {},
	175: {},
	176: {},
	177: {},
	178: {},
	179: {},
	180: {},
	181: {},
	182: {},
	183: {
		Flex_state: uint16(31),
	},
	184: {
		Flex_state: uint16(31),
	},
	185: {
		Flex_state: uint16(31),
	},
	186: {
		Flex_state: uint16(31),
	},
	187: {
		Flex_state: uint16(31),
	},
	188: {
		Flex_state: uint16(31),
	},
	189: {
		Flex_state: uint16(31),
	},
	190: {
		Flex_state: uint16(31),
	},
	191: {
		Flex_state: uint16(31),
	},
	192: {},
	193: {
		Flex_state: uint16(31),
	},
	194: {},
	195: {},
	196: {},
	197: {},
	198: {},
	199: {},
	200: {},
	201: {},
	202: {},
	203: {},
	204: {},
	205: {},
	206: {},
	207: {},
	208: {},
	209: {},
	210: {},
	211: {},
	212: {},
	213: {},
	214: {
		Flex_state: uint16(9),
	},
	215: {},
	216: {},
	217: {},
	218: {},
	219: {},
	220: {},
	221: {},
	222: {},
	223: {
		Flex_state: uint16(31),
	},
	224: {},
	225: {},
	226: {},
	227: {},
	228: {},
	229: {},
	230: {},
	231: {},
	232: {},
	233: {},
	234: {},
	235: {},
	236: {},
	237: {},
	238: {
		Flex_state: uint16(31),
	},
	239: {
		Flex_state: uint16(31),
	},
	240: {},
	241: {},
	242: {
		Flex_state: uint16(54),
	},
	243: {
		Flex_state: uint16(2),
	},
	244: {},
	245: {},
	246: {},
	247: {
		Flex_state: uint16(54),
	},
	248: {},
	249: {},
	250: {},
}

var ts_parse_table = [12][163]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
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
		27: uint16(1),
		28: uint16(1),
		29: uint16(1),
		30: uint16(1),
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
		52: uint16(1),
		53: uint16(1),
		54: uint16(1),
		55: uint16(1),
		56: uint16(1),
		57: uint16(1),
		58: uint16(1),
		59: uint16(1),
		60: uint16(1),
		61: uint16(1),
		62: uint16(1),
		63: uint16(1),
		64: uint16(1),
		65: uint16(1),
		66: uint16(1),
		67: uint16(1),
		68: uint16(1),
		69: uint16(1),
		70: uint16(1),
		71: uint16(1),
		72: uint16(1),
		73: uint16(1),
		74: uint16(1),
		75: uint16(1),
		76: uint16(1),
		77: uint16(1),
		78: uint16(1),
		79: uint16(1),
		80: uint16(1),
		81: uint16(1),
		82: uint16(1),
		83: uint16(1),
		84: uint16(1),
		85: uint16(1),
		86: uint16(1),
		87: uint16(1),
		88: uint16(1),
		89: uint16(1),
		90: uint16(1),
		91: uint16(1),
		92: uint16(1),
		93: uint16(1),
		94: uint16(1),
		95: uint16(1),
		96: uint16(1),
		97: uint16(1),
		98: uint16(1),
		99: uint16(3),
	},
	1: {
		0:   uint16(5),
		1:   uint16(7),
		2:   uint16(9),
		3:   uint16(11),
		4:   uint16(13),
		8:   uint16(15),
		9:   uint16(17),
		99:  uint16(3),
		100: uint16(200),
		101: uint16(76),
		102: uint16(76),
		103: uint16(76),
		104: uint16(76),
		105: uint16(186),
		108: uint16(154),
		109: uint16(154),
		110: uint16(154),
		153: uint16(76),
	},
	2: {
		1:   uint16(19),
		5:   uint16(21),
		9:   uint16(23),
		17:  uint16(25),
		28:  uint16(27),
		29:  uint16(27),
		30:  uint16(29),
		32:  uint16(31),
		33:  uint16(33),
		34:  uint16(35),
		35:  uint16(35),
		36:  uint16(37),
		41:  uint16(39),
		52:  uint16(41),
		54:  uint16(43),
		55:  uint16(43),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(47),
		60:  uint16(49),
		62:  uint16(51),
		63:  uint16(51),
		64:  uint16(51),
		65:  uint16(51),
		66:  uint16(53),
		67:  uint16(53),
		68:  uint16(55),
		82:  uint16(57),
		83:  uint16(57),
		84:  uint16(57),
		85:  uint16(57),
		86:  uint16(57),
		87:  uint16(59),
		88:  uint16(61),
		89:  uint16(63),
		90:  uint16(63),
		91:  uint16(65),
		92:  uint16(65),
		93:  uint16(65),
		94:  uint16(67),
		96:  uint16(69),
		98:  uint16(71),
		99:  uint16(3),
		118: uint16(71),
		119: uint16(71),
		120: uint16(71),
		121: uint16(71),
		122: uint16(41),
		124: uint16(41),
		128: uint16(71),
		129: uint16(71),
		131: uint16(71),
		132: uint16(71),
		133: uint16(71),
		134: uint16(71),
		135: uint16(71),
		136: uint16(71),
		137: uint16(71),
		138: uint16(15),
		139: uint16(43),
		143: uint16(71),
		144: uint16(71),
		145: uint16(16),
		147: uint16(71),
		148: uint16(71),
		149: uint16(71),
		150: uint16(71),
		151: uint16(71),
		152: uint16(71),
		157: uint16(4),
		161: uint16(21),
	},
	3: {
		1:   uint16(19),
		5:   uint16(21),
		9:   uint16(23),
		17:  uint16(73),
		28:  uint16(27),
		29:  uint16(27),
		30:  uint16(29),
		32:  uint16(31),
		33:  uint16(33),
		34:  uint16(35),
		35:  uint16(35),
		36:  uint16(37),
		41:  uint16(39),
		52:  uint16(41),
		54:  uint16(43),
		55:  uint16(43),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(47),
		60:  uint16(49),
		62:  uint16(51),
		63:  uint16(51),
		64:  uint16(51),
		65:  uint16(51),
		66:  uint16(53),
		67:  uint16(53),
		68:  uint16(55),
		82:  uint16(57),
		83:  uint16(57),
		84:  uint16(57),
		85:  uint16(57),
		86:  uint16(57),
		87:  uint16(59),
		88:  uint16(61),
		89:  uint16(63),
		90:  uint16(63),
		91:  uint16(65),
		92:  uint16(65),
		93:  uint16(65),
		94:  uint16(67),
		96:  uint16(69),
		98:  uint16(71),
		99:  uint16(3),
		118: uint16(71),
		119: uint16(71),
		120: uint16(71),
		121: uint16(71),
		122: uint16(41),
		124: uint16(41),
		128: uint16(71),
		129: uint16(71),
		131: uint16(71),
		132: uint16(71),
		133: uint16(71),
		134: uint16(71),
		135: uint16(71),
		136: uint16(71),
		137: uint16(71),
		138: uint16(15),
		139: uint16(43),
		143: uint16(71),
		144: uint16(71),
		145: uint16(16),
		147: uint16(71),
		148: uint16(71),
		149: uint16(71),
		150: uint16(71),
		151: uint16(71),
		152: uint16(71),
		157: uint16(2),
		161: uint16(21),
	},
	4: {
		1:   uint16(75),
		5:   uint16(78),
		9:   uint16(81),
		17:  uint16(84),
		28:  uint16(86),
		29:  uint16(86),
		30:  uint16(89),
		32:  uint16(92),
		33:  uint16(95),
		34:  uint16(98),
		35:  uint16(98),
		36:  uint16(101),
		41:  uint16(104),
		52:  uint16(107),
		54:  uint16(110),
		55:  uint16(110),
		56:  uint16(113),
		57:  uint16(113),
		58:  uint16(116),
		59:  uint16(116),
		60:  uint16(119),
		62:  uint16(122),
		63:  uint16(122),
		64:  uint16(122),
		65:  uint16(122),
		66:  uint16(125),
		67:  uint16(125),
		68:  uint16(128),
		82:  uint16(131),
		83:  uint16(131),
		84:  uint16(131),
		85:  uint16(131),
		86:  uint16(131),
		87:  uint16(134),
		88:  uint16(137),
		89:  uint16(140),
		90:  uint16(140),
		91:  uint16(143),
		92:  uint16(143),
		93:  uint16(143),
		94:  uint16(146),
		96:  uint16(149),
		98:  uint16(152),
		99:  uint16(3),
		118: uint16(71),
		119: uint16(71),
		120: uint16(71),
		121: uint16(71),
		122: uint16(41),
		124: uint16(41),
		128: uint16(71),
		129: uint16(71),
		131: uint16(71),
		132: uint16(71),
		133: uint16(71),
		134: uint16(71),
		135: uint16(71),
		136: uint16(71),
		137: uint16(71),
		138: uint16(15),
		139: uint16(43),
		143: uint16(71),
		144: uint16(71),
		145: uint16(16),
		147: uint16(71),
		148: uint16(71),
		149: uint16(71),
		150: uint16(71),
		151: uint16(71),
		152: uint16(71),
		157: uint16(4),
		161: uint16(21),
	},
	5: {
		1:   uint16(19),
		9:   uint16(23),
		19:  uint16(155),
		28:  uint16(157),
		29:  uint16(157),
		30:  uint16(29),
		32:  uint16(31),
		33:  uint16(33),
		34:  uint16(35),
		35:  uint16(35),
		36:  uint16(37),
		41:  uint16(39),
		52:  uint16(41),
		54:  uint16(43),
		55:  uint16(43),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(47),
		60:  uint16(49),
		62:  uint16(51),
		63:  uint16(51),
		64:  uint16(51),
		65:  uint16(51),
		66:  uint16(53),
		67:  uint16(53),
		68:  uint16(55),
		82:  uint16(57),
		83:  uint16(57),
		84:  uint16(57),
		85:  uint16(57),
		86:  uint16(57),
		87:  uint16(59),
		88:  uint16(61),
		89:  uint16(63),
		90:  uint16(63),
		91:  uint16(65),
		92:  uint16(65),
		93:  uint16(65),
		94:  uint16(67),
		96:  uint16(69),
		98:  uint16(71),
		99:  uint16(3),
		118: uint16(6),
		119: uint16(6),
		120: uint16(6),
		121: uint16(6),
		122: uint16(41),
		124: uint16(41),
		128: uint16(6),
		129: uint16(6),
		131: uint16(6),
		132: uint16(6),
		133: uint16(6),
		134: uint16(6),
		135: uint16(6),
		136: uint16(6),
		137: uint16(6),
		138: uint16(15),
		139: uint16(43),
		143: uint16(6),
		144: uint16(6),
		145: uint16(16),
		147: uint16(6),
		148: uint16(6),
		149: uint16(6),
		150: uint16(6),
		151: uint16(6),
		152: uint16(6),
		158: uint16(6),
		161: uint16(21),
	},
	6: {
		1:   uint16(159),
		9:   uint16(162),
		19:  uint16(165),
		28:  uint16(167),
		29:  uint16(167),
		30:  uint16(170),
		32:  uint16(173),
		33:  uint16(176),
		34:  uint16(179),
		35:  uint16(179),
		36:  uint16(182),
		41:  uint16(185),
		52:  uint16(188),
		54:  uint16(191),
		55:  uint16(191),
		56:  uint16(194),
		57:  uint16(194),
		58:  uint16(197),
		59:  uint16(197),
		60:  uint16(200),
		62:  uint16(203),
		63:  uint16(203),
		64:  uint16(203),
		65:  uint16(203),
		66:  uint16(206),
		67:  uint16(206),
		68:  uint16(209),
		82:  uint16(212),
		83:  uint16(212),
		84:  uint16(212),
		85:  uint16(212),
		86:  uint16(212),
		87:  uint16(215),
		88:  uint16(218),
		89:  uint16(221),
		90:  uint16(221),
		91:  uint16(224),
		92:  uint16(224),
		93:  uint16(224),
		94:  uint16(227),
		96:  uint16(230),
		98:  uint16(233),
		99:  uint16(3),
		118: uint16(6),
		119: uint16(6),
		120: uint16(6),
		121: uint16(6),
		122: uint16(41),
		124: uint16(41),
		128: uint16(6),
		129: uint16(6),
		131: uint16(6),
		132: uint16(6),
		133: uint16(6),
		134: uint16(6),
		135: uint16(6),
		136: uint16(6),
		137: uint16(6),
		138: uint16(15),
		139: uint16(43),
		143: uint16(6),
		144: uint16(6),
		145: uint16(16),
		147: uint16(6),
		148: uint16(6),
		149: uint16(6),
		150: uint16(6),
		151: uint16(6),
		152: uint16(6),
		158: uint16(6),
		161: uint16(21),
	},
	7: {
		1:   uint16(19),
		9:   uint16(23),
		28:  uint16(236),
		29:  uint16(236),
		30:  uint16(29),
		32:  uint16(31),
		33:  uint16(33),
		34:  uint16(35),
		35:  uint16(35),
		36:  uint16(37),
		41:  uint16(39),
		52:  uint16(41),
		54:  uint16(43),
		55:  uint16(43),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(47),
		60:  uint16(49),
		62:  uint16(51),
		63:  uint16(51),
		64:  uint16(51),
		65:  uint16(51),
		66:  uint16(53),
		67:  uint16(53),
		68:  uint16(55),
		82:  uint16(57),
		83:  uint16(57),
		84:  uint16(57),
		85:  uint16(57),
		86:  uint16(57),
		87:  uint16(59),
		88:  uint16(61),
		89:  uint16(63),
		90:  uint16(63),
		91:  uint16(65),
		92:  uint16(65),
		93:  uint16(65),
		94:  uint16(67),
		96:  uint16(69),
		98:  uint16(71),
		99:  uint16(3),
		118: uint16(5),
		119: uint16(5),
		120: uint16(5),
		121: uint16(5),
		122: uint16(41),
		124: uint16(41),
		128: uint16(5),
		129: uint16(5),
		131: uint16(5),
		132: uint16(5),
		133: uint16(5),
		134: uint16(5),
		135: uint16(5),
		136: uint16(5),
		137: uint16(5),
		138: uint16(15),
		139: uint16(43),
		143: uint16(5),
		144: uint16(5),
		145: uint16(16),
		147: uint16(5),
		148: uint16(5),
		149: uint16(5),
		150: uint16(5),
		151: uint16(5),
		152: uint16(5),
		158: uint16(5),
		161: uint16(21),
	},
	8: {
		1:   uint16(19),
		9:   uint16(23),
		28:  uint16(238),
		29:  uint16(238),
		30:  uint16(240),
		32:  uint16(242),
		33:  uint16(244),
		34:  uint16(35),
		35:  uint16(35),
		36:  uint16(37),
		41:  uint16(39),
		52:  uint16(246),
		54:  uint16(43),
		55:  uint16(43),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(47),
		60:  uint16(49),
		62:  uint16(51),
		63:  uint16(51),
		64:  uint16(51),
		65:  uint16(51),
		66:  uint16(53),
		67:  uint16(53),
		68:  uint16(55),
		82:  uint16(57),
		83:  uint16(57),
		84:  uint16(57),
		85:  uint16(57),
		86:  uint16(57),
		87:  uint16(59),
		88:  uint16(61),
		89:  uint16(63),
		90:  uint16(63),
		91:  uint16(65),
		92:  uint16(65),
		93:  uint16(65),
		94:  uint16(67),
		96:  uint16(69),
		98:  uint16(71),
		99:  uint16(3),
		118: uint16(197),
		119: uint16(197),
		120: uint16(197),
		121: uint16(197),
		122: uint16(41),
		124: uint16(41),
		128: uint16(197),
		129: uint16(197),
		131: uint16(197),
		132: uint16(197),
		133: uint16(197),
		134: uint16(197),
		135: uint16(197),
		136: uint16(197),
		137: uint16(197),
		138: uint16(100),
		139: uint16(43),
		143: uint16(197),
		144: uint16(197),
		145: uint16(103),
		147: uint16(197),
		148: uint16(197),
		149: uint16(197),
		150: uint16(197),
		151: uint16(197),
		152: uint16(197),
		161: uint16(21),
	},
	9: {
		1:   uint16(19),
		9:   uint16(23),
		28:  uint16(248),
		29:  uint16(248),
		30:  uint16(29),
		32:  uint16(31),
		33:  uint16(33),
		34:  uint16(35),
		35:  uint16(35),
		36:  uint16(37),
		41:  uint16(39),
		52:  uint16(41),
		54:  uint16(43),
		55:  uint16(43),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(47),
		60:  uint16(49),
		62:  uint16(51),
		63:  uint16(51),
		64:  uint16(51),
		65:  uint16(51),
		66:  uint16(53),
		67:  uint16(53),
		68:  uint16(55),
		82:  uint16(57),
		83:  uint16(57),
		84:  uint16(57),
		85:  uint16(57),
		86:  uint16(57),
		87:  uint16(59),
		88:  uint16(61),
		89:  uint16(63),
		90:  uint16(63),
		91:  uint16(65),
		92:  uint16(65),
		93:  uint16(65),
		94:  uint16(67),
		96:  uint16(69),
		98:  uint16(71),
		99:  uint16(3),
		118: uint16(3),
		119: uint16(3),
		120: uint16(3),
		121: uint16(3),
		122: uint16(41),
		124: uint16(41),
		128: uint16(3),
		129: uint16(3),
		131: uint16(3),
		132: uint16(3),
		133: uint16(3),
		134: uint16(3),
		135: uint16(3),
		136: uint16(3),
		137: uint16(3),
		138: uint16(15),
		139: uint16(43),
		143: uint16(3),
		144: uint16(3),
		145: uint16(16),
		147: uint16(3),
		148: uint16(3),
		149: uint16(3),
		150: uint16(3),
		151: uint16(3),
		152: uint16(3),
		161: uint16(21),
	},
	10: {
		1:   uint16(19),
		9:   uint16(23),
		28:  uint16(250),
		29:  uint16(250),
		30:  uint16(29),
		32:  uint16(31),
		33:  uint16(33),
		34:  uint16(35),
		35:  uint16(35),
		36:  uint16(37),
		41:  uint16(39),
		52:  uint16(41),
		54:  uint16(43),
		55:  uint16(43),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(47),
		60:  uint16(49),
		62:  uint16(51),
		63:  uint16(51),
		64:  uint16(51),
		65:  uint16(51),
		66:  uint16(53),
		67:  uint16(53),
		68:  uint16(55),
		82:  uint16(57),
		83:  uint16(57),
		84:  uint16(57),
		85:  uint16(57),
		86:  uint16(57),
		87:  uint16(59),
		88:  uint16(61),
		89:  uint16(63),
		90:  uint16(63),
		91:  uint16(65),
		92:  uint16(65),
		93:  uint16(65),
		94:  uint16(67),
		96:  uint16(69),
		98:  uint16(71),
		99:  uint16(3),
		118: uint16(70),
		119: uint16(70),
		120: uint16(70),
		121: uint16(70),
		122: uint16(41),
		124: uint16(41),
		128: uint16(70),
		129: uint16(70),
		131: uint16(70),
		132: uint16(70),
		133: uint16(70),
		134: uint16(70),
		135: uint16(70),
		136: uint16(70),
		137: uint16(70),
		138: uint16(15),
		139: uint16(43),
		143: uint16(70),
		144: uint16(70),
		145: uint16(16),
		147: uint16(70),
		148: uint16(70),
		149: uint16(70),
		150: uint16(70),
		151: uint16(70),
		152: uint16(70),
		161: uint16(21),
	},
	11: {
		1:   uint16(19),
		9:   uint16(23),
		28:  uint16(252),
		29:  uint16(252),
		30:  uint16(240),
		32:  uint16(242),
		33:  uint16(244),
		34:  uint16(35),
		35:  uint16(35),
		36:  uint16(37),
		41:  uint16(39),
		52:  uint16(246),
		54:  uint16(43),
		55:  uint16(43),
		56:  uint16(45),
		57:  uint16(45),
		58:  uint16(47),
		59:  uint16(47),
		60:  uint16(49),
		62:  uint16(51),
		63:  uint16(51),
		64:  uint16(51),
		65:  uint16(51),
		66:  uint16(53),
		67:  uint16(53),
		68:  uint16(55),
		82:  uint16(57),
		83:  uint16(57),
		84:  uint16(57),
		85:  uint16(57),
		86:  uint16(57),
		87:  uint16(59),
		88:  uint16(61),
		89:  uint16(63),
		90:  uint16(63),
		91:  uint16(65),
		92:  uint16(65),
		93:  uint16(65),
		94:  uint16(67),
		96:  uint16(69),
		98:  uint16(71),
		99:  uint16(3),
		118: uint16(196),
		119: uint16(196),
		120: uint16(196),
		121: uint16(196),
		122: uint16(41),
		124: uint16(41),
		128: uint16(196),
		129: uint16(196),
		131: uint16(196),
		132: uint16(196),
		133: uint16(196),
		134: uint16(196),
		135: uint16(196),
		136: uint16(196),
		137: uint16(196),
		138: uint16(100),
		139: uint16(43),
		143: uint16(196),
		144: uint16(196),
		145: uint16(103),
		147: uint16(196),
		148: uint16(196),
		149: uint16(196),
		150: uint16(196),
		151: uint16(196),
		152: uint16(196),
		161: uint16(21),
	},
}

var ts_small_parse_table = [5800]uint16_t{
	0:    uint16(7),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(258),
	5:    uint16(1),
	6:    uint16(anon_sym_PERCENT),
	7:    uint16(260),
	8:    uint16(2),
	9:    uint16(anon_sym_px),
	10:   uint16(anon_sym_mm),
	11:   uint16(38),
	12:   uint16(2),
	13:   uint16(sym_integer_distance_unit),
	14:   uint16(sym_float_distance_unit),
	15:   uint16(262),
	16:   uint16(3),
	17:   uint16(anon_sym_cm),
	18:   uint16(anon_sym_ph),
	19:   uint16(anon_sym_em),
	20:   uint16(256),
	21:   uint16(9),
	22:   uint16(anon_sym_COMMA),
	23:   uint16(anon_sym_POUND),
	24:   uint16(anon_sym_SEMI),
	25:   uint16(anon_sym_RPAREN),
	26:   uint16(anon_sym_DQUOTE),
	27:   uint16(sym_float_value),
	28:   uint16(anon_sym_AT),
	29:   uint16(anon_sym_LBRACK),
	30:   uint16(anon_sym_DOLLAR),
	31:   uint16(254),
	32:   uint16(35),
	33:   uint16(anon_sym_inherit),
	34:   uint16(anon_sym_DMENU),
	35:   uint16(sym_integer_value),
	36:   uint16(anon_sym_true),
	37:   uint16(anon_sym_false),
	38:   uint16(anon_sym_url),
	39:   uint16(anon_sym_linear_DASHgradient),
	40:   uint16(anon_sym_0),
	41:   uint16(anon_sym_rgb),
	42:   uint16(anon_sym_rgba),
	43:   uint16(anon_sym_hsl),
	44:   uint16(anon_sym_hsla),
	45:   uint16(anon_sym_hwb),
	46:   uint16(anon_sym_hwba),
	47:   uint16(anon_sym_cmyk),
	48:   uint16(anon_sym_bold),
	49:   uint16(anon_sym_italic),
	50:   uint16(anon_sym_underline),
	51:   uint16(anon_sym_strikethrough),
	52:   uint16(anon_sym_dash),
	53:   uint16(anon_sym_solid),
	54:   uint16(anon_sym_calc),
	55:   uint16(anon_sym_center),
	56:   uint16(anon_sym_north),
	57:   uint16(anon_sym_east),
	58:   uint16(anon_sym_south),
	59:   uint16(anon_sym_west),
	60:   uint16(anon_sym_var),
	61:   uint16(anon_sym_horizontal),
	62:   uint16(anon_sym_vertical),
	63:   uint16(anon_sym_default),
	64:   uint16(anon_sym_pointer),
	65:   uint16(anon_sym_text),
	66:   uint16(anon_sym_env),
	67:   uint16(sym_identifier),
	68:   uint16(7),
	69:   uint16(3),
	70:   uint16(1),
	71:   uint16(sym_comment),
	72:   uint16(258),
	73:   uint16(1),
	74:   uint16(anon_sym_PERCENT),
	75:   uint16(260),
	76:   uint16(2),
	77:   uint16(anon_sym_px),
	78:   uint16(anon_sym_mm),
	79:   uint16(27),
	80:   uint16(2),
	81:   uint16(sym_integer_distance_unit),
	82:   uint16(sym_float_distance_unit),
	83:   uint16(262),
	84:   uint16(3),
	85:   uint16(anon_sym_cm),
	86:   uint16(anon_sym_ph),
	87:   uint16(anon_sym_em),
	88:   uint16(266),
	89:   uint16(9),
	90:   uint16(anon_sym_COMMA),
	91:   uint16(anon_sym_POUND),
	92:   uint16(anon_sym_SEMI),
	93:   uint16(anon_sym_RPAREN),
	94:   uint16(anon_sym_DQUOTE),
	95:   uint16(sym_float_value),
	96:   uint16(anon_sym_AT),
	97:   uint16(anon_sym_LBRACK),
	98:   uint16(anon_sym_DOLLAR),
	99:   uint16(264),
	100:  uint16(35),
	101:  uint16(anon_sym_inherit),
	102:  uint16(anon_sym_DMENU),
	103:  uint16(sym_integer_value),
	104:  uint16(anon_sym_true),
	105:  uint16(anon_sym_false),
	106:  uint16(anon_sym_url),
	107:  uint16(anon_sym_linear_DASHgradient),
	108:  uint16(anon_sym_0),
	109:  uint16(anon_sym_rgb),
	110:  uint16(anon_sym_rgba),
	111:  uint16(anon_sym_hsl),
	112:  uint16(anon_sym_hsla),
	113:  uint16(anon_sym_hwb),
	114:  uint16(anon_sym_hwba),
	115:  uint16(anon_sym_cmyk),
	116:  uint16(anon_sym_bold),
	117:  uint16(anon_sym_italic),
	118:  uint16(anon_sym_underline),
	119:  uint16(anon_sym_strikethrough),
	120:  uint16(anon_sym_dash),
	121:  uint16(anon_sym_solid),
	122:  uint16(anon_sym_calc),
	123:  uint16(anon_sym_center),
	124:  uint16(anon_sym_north),
	125:  uint16(anon_sym_east),
	126:  uint16(anon_sym_south),
	127:  uint16(anon_sym_west),
	128:  uint16(anon_sym_var),
	129:  uint16(anon_sym_horizontal),
	130:  uint16(anon_sym_vertical),
	131:  uint16(anon_sym_default),
	132:  uint16(anon_sym_pointer),
	133:  uint16(anon_sym_text),
	134:  uint16(anon_sym_env),
	135:  uint16(sym_identifier),
	136:  uint16(6),
	137:  uint16(3),
	138:  uint16(1),
	139:  uint16(sym_comment),
	140:  uint16(258),
	141:  uint16(1),
	142:  uint16(anon_sym_PERCENT),
	143:  uint16(38),
	144:  uint16(1),
	145:  uint16(sym_float_distance_unit),
	146:  uint16(262),
	147:  uint16(3),
	148:  uint16(anon_sym_cm),
	149:  uint16(anon_sym_ph),
	150:  uint16(anon_sym_em),
	151:  uint16(256),
	152:  uint16(9),
	153:  uint16(anon_sym_COMMA),
	154:  uint16(anon_sym_POUND),
	155:  uint16(anon_sym_SEMI),
	156:  uint16(anon_sym_RPAREN),
	157:  uint16(anon_sym_DQUOTE),
	158:  uint16(sym_float_value),
	159:  uint16(anon_sym_AT),
	160:  uint16(anon_sym_LBRACK),
	161:  uint16(anon_sym_DOLLAR),
	162:  uint16(254),
	163:  uint16(35),
	164:  uint16(anon_sym_inherit),
	165:  uint16(anon_sym_DMENU),
	166:  uint16(sym_integer_value),
	167:  uint16(anon_sym_true),
	168:  uint16(anon_sym_false),
	169:  uint16(anon_sym_url),
	170:  uint16(anon_sym_linear_DASHgradient),
	171:  uint16(anon_sym_0),
	172:  uint16(anon_sym_rgb),
	173:  uint16(anon_sym_rgba),
	174:  uint16(anon_sym_hsl),
	175:  uint16(anon_sym_hsla),
	176:  uint16(anon_sym_hwb),
	177:  uint16(anon_sym_hwba),
	178:  uint16(anon_sym_cmyk),
	179:  uint16(anon_sym_bold),
	180:  uint16(anon_sym_italic),
	181:  uint16(anon_sym_underline),
	182:  uint16(anon_sym_strikethrough),
	183:  uint16(anon_sym_dash),
	184:  uint16(anon_sym_solid),
	185:  uint16(anon_sym_calc),
	186:  uint16(anon_sym_center),
	187:  uint16(anon_sym_north),
	188:  uint16(anon_sym_east),
	189:  uint16(anon_sym_south),
	190:  uint16(anon_sym_west),
	191:  uint16(anon_sym_var),
	192:  uint16(anon_sym_horizontal),
	193:  uint16(anon_sym_vertical),
	194:  uint16(anon_sym_default),
	195:  uint16(anon_sym_pointer),
	196:  uint16(anon_sym_text),
	197:  uint16(anon_sym_env),
	198:  uint16(sym_identifier),
	199:  uint16(11),
	200:  uint16(3),
	201:  uint16(1),
	202:  uint16(sym_comment),
	203:  uint16(41),
	204:  uint16(1),
	205:  uint16(anon_sym_0),
	206:  uint16(55),
	207:  uint16(1),
	208:  uint16(anon_sym_calc),
	209:  uint16(268),
	210:  uint16(1),
	211:  uint16(sym_integer_value),
	212:  uint16(270),
	213:  uint16(1),
	214:  uint16(sym_float_value),
	215:  uint16(20),
	216:  uint16(1),
	217:  uint16(sym_distance_value),
	218:  uint16(30),
	219:  uint16(1),
	220:  uint16(sym_line_style_value),
	221:  uint16(43),
	222:  uint16(1),
	223:  uint16(sym_distance_calc),
	224:  uint16(53),
	225:  uint16(2),
	226:  uint16(anon_sym_dash),
	227:  uint16(anon_sym_solid),
	228:  uint16(256),
	229:  uint16(8),
	230:  uint16(anon_sym_COMMA),
	231:  uint16(anon_sym_POUND),
	232:  uint16(anon_sym_SEMI),
	233:  uint16(anon_sym_RPAREN),
	234:  uint16(anon_sym_DQUOTE),
	235:  uint16(anon_sym_AT),
	236:  uint16(anon_sym_LBRACK),
	237:  uint16(anon_sym_DOLLAR),
	238:  uint16(254),
	239:  uint16(30),
	240:  uint16(anon_sym_inherit),
	241:  uint16(anon_sym_DMENU),
	242:  uint16(anon_sym_true),
	243:  uint16(anon_sym_false),
	244:  uint16(anon_sym_url),
	245:  uint16(anon_sym_linear_DASHgradient),
	246:  uint16(anon_sym_rgb),
	247:  uint16(anon_sym_rgba),
	248:  uint16(anon_sym_hsl),
	249:  uint16(anon_sym_hsla),
	250:  uint16(anon_sym_hwb),
	251:  uint16(anon_sym_hwba),
	252:  uint16(anon_sym_cmyk),
	253:  uint16(anon_sym_bold),
	254:  uint16(anon_sym_italic),
	255:  uint16(anon_sym_underline),
	256:  uint16(anon_sym_strikethrough),
	257:  uint16(anon_sym_center),
	258:  uint16(anon_sym_north),
	259:  uint16(anon_sym_east),
	260:  uint16(anon_sym_south),
	261:  uint16(anon_sym_west),
	262:  uint16(anon_sym_var),
	263:  uint16(anon_sym_horizontal),
	264:  uint16(anon_sym_vertical),
	265:  uint16(anon_sym_default),
	266:  uint16(anon_sym_pointer),
	267:  uint16(anon_sym_text),
	268:  uint16(anon_sym_env),
	269:  uint16(sym_identifier),
	270:  uint16(10),
	271:  uint16(3),
	272:  uint16(1),
	273:  uint16(sym_comment),
	274:  uint16(41),
	275:  uint16(1),
	276:  uint16(anon_sym_0),
	277:  uint16(55),
	278:  uint16(1),
	279:  uint16(anon_sym_calc),
	280:  uint16(268),
	281:  uint16(1),
	282:  uint16(sym_integer_value),
	283:  uint16(270),
	284:  uint16(1),
	285:  uint16(sym_float_value),
	286:  uint16(17),
	287:  uint16(1),
	288:  uint16(sym_border_style),
	289:  uint16(22),
	290:  uint16(1),
	291:  uint16(sym_distance_value),
	292:  uint16(43),
	293:  uint16(1),
	294:  uint16(sym_distance_calc),
	295:  uint16(274),
	296:  uint16(8),
	297:  uint16(anon_sym_COMMA),
	298:  uint16(anon_sym_POUND),
	299:  uint16(anon_sym_SEMI),
	300:  uint16(anon_sym_RPAREN),
	301:  uint16(anon_sym_DQUOTE),
	302:  uint16(anon_sym_AT),
	303:  uint16(anon_sym_LBRACK),
	304:  uint16(anon_sym_DOLLAR),
	305:  uint16(272),
	306:  uint16(32),
	307:  uint16(anon_sym_inherit),
	308:  uint16(anon_sym_DMENU),
	309:  uint16(anon_sym_true),
	310:  uint16(anon_sym_false),
	311:  uint16(anon_sym_url),
	312:  uint16(anon_sym_linear_DASHgradient),
	313:  uint16(anon_sym_rgb),
	314:  uint16(anon_sym_rgba),
	315:  uint16(anon_sym_hsl),
	316:  uint16(anon_sym_hsla),
	317:  uint16(anon_sym_hwb),
	318:  uint16(anon_sym_hwba),
	319:  uint16(anon_sym_cmyk),
	320:  uint16(anon_sym_bold),
	321:  uint16(anon_sym_italic),
	322:  uint16(anon_sym_underline),
	323:  uint16(anon_sym_strikethrough),
	324:  uint16(anon_sym_dash),
	325:  uint16(anon_sym_solid),
	326:  uint16(anon_sym_center),
	327:  uint16(anon_sym_north),
	328:  uint16(anon_sym_east),
	329:  uint16(anon_sym_south),
	330:  uint16(anon_sym_west),
	331:  uint16(anon_sym_var),
	332:  uint16(anon_sym_horizontal),
	333:  uint16(anon_sym_vertical),
	334:  uint16(anon_sym_default),
	335:  uint16(anon_sym_pointer),
	336:  uint16(anon_sym_text),
	337:  uint16(anon_sym_env),
	338:  uint16(sym_identifier),
	339:  uint16(10),
	340:  uint16(3),
	341:  uint16(1),
	342:  uint16(sym_comment),
	343:  uint16(41),
	344:  uint16(1),
	345:  uint16(anon_sym_0),
	346:  uint16(55),
	347:  uint16(1),
	348:  uint16(anon_sym_calc),
	349:  uint16(268),
	350:  uint16(1),
	351:  uint16(sym_integer_value),
	352:  uint16(270),
	353:  uint16(1),
	354:  uint16(sym_float_value),
	355:  uint16(18),
	356:  uint16(1),
	357:  uint16(sym_border_style),
	358:  uint16(22),
	359:  uint16(1),
	360:  uint16(sym_distance_value),
	361:  uint16(43),
	362:  uint16(1),
	363:  uint16(sym_distance_calc),
	364:  uint16(278),
	365:  uint16(8),
	366:  uint16(anon_sym_COMMA),
	367:  uint16(anon_sym_POUND),
	368:  uint16(anon_sym_SEMI),
	369:  uint16(anon_sym_RPAREN),
	370:  uint16(anon_sym_DQUOTE),
	371:  uint16(anon_sym_AT),
	372:  uint16(anon_sym_LBRACK),
	373:  uint16(anon_sym_DOLLAR),
	374:  uint16(276),
	375:  uint16(32),
	376:  uint16(anon_sym_inherit),
	377:  uint16(anon_sym_DMENU),
	378:  uint16(anon_sym_true),
	379:  uint16(anon_sym_false),
	380:  uint16(anon_sym_url),
	381:  uint16(anon_sym_linear_DASHgradient),
	382:  uint16(anon_sym_rgb),
	383:  uint16(anon_sym_rgba),
	384:  uint16(anon_sym_hsl),
	385:  uint16(anon_sym_hsla),
	386:  uint16(anon_sym_hwb),
	387:  uint16(anon_sym_hwba),
	388:  uint16(anon_sym_cmyk),
	389:  uint16(anon_sym_bold),
	390:  uint16(anon_sym_italic),
	391:  uint16(anon_sym_underline),
	392:  uint16(anon_sym_strikethrough),
	393:  uint16(anon_sym_dash),
	394:  uint16(anon_sym_solid),
	395:  uint16(anon_sym_center),
	396:  uint16(anon_sym_north),
	397:  uint16(anon_sym_east),
	398:  uint16(anon_sym_south),
	399:  uint16(anon_sym_west),
	400:  uint16(anon_sym_var),
	401:  uint16(anon_sym_horizontal),
	402:  uint16(anon_sym_vertical),
	403:  uint16(anon_sym_default),
	404:  uint16(anon_sym_pointer),
	405:  uint16(anon_sym_text),
	406:  uint16(anon_sym_env),
	407:  uint16(sym_identifier),
	408:  uint16(10),
	409:  uint16(3),
	410:  uint16(1),
	411:  uint16(sym_comment),
	412:  uint16(41),
	413:  uint16(1),
	414:  uint16(anon_sym_0),
	415:  uint16(55),
	416:  uint16(1),
	417:  uint16(anon_sym_calc),
	418:  uint16(268),
	419:  uint16(1),
	420:  uint16(sym_integer_value),
	421:  uint16(270),
	422:  uint16(1),
	423:  uint16(sym_float_value),
	424:  uint16(22),
	425:  uint16(1),
	426:  uint16(sym_distance_value),
	427:  uint16(43),
	428:  uint16(1),
	429:  uint16(sym_distance_calc),
	430:  uint16(50),
	431:  uint16(1),
	432:  uint16(sym_border_style),
	433:  uint16(282),
	434:  uint16(8),
	435:  uint16(anon_sym_COMMA),
	436:  uint16(anon_sym_POUND),
	437:  uint16(anon_sym_SEMI),
	438:  uint16(anon_sym_RPAREN),
	439:  uint16(anon_sym_DQUOTE),
	440:  uint16(anon_sym_AT),
	441:  uint16(anon_sym_LBRACK),
	442:  uint16(anon_sym_DOLLAR),
	443:  uint16(280),
	444:  uint16(32),
	445:  uint16(anon_sym_inherit),
	446:  uint16(anon_sym_DMENU),
	447:  uint16(anon_sym_true),
	448:  uint16(anon_sym_false),
	449:  uint16(anon_sym_url),
	450:  uint16(anon_sym_linear_DASHgradient),
	451:  uint16(anon_sym_rgb),
	452:  uint16(anon_sym_rgba),
	453:  uint16(anon_sym_hsl),
	454:  uint16(anon_sym_hsla),
	455:  uint16(anon_sym_hwb),
	456:  uint16(anon_sym_hwba),
	457:  uint16(anon_sym_cmyk),
	458:  uint16(anon_sym_bold),
	459:  uint16(anon_sym_italic),
	460:  uint16(anon_sym_underline),
	461:  uint16(anon_sym_strikethrough),
	462:  uint16(anon_sym_dash),
	463:  uint16(anon_sym_solid),
	464:  uint16(anon_sym_center),
	465:  uint16(anon_sym_north),
	466:  uint16(anon_sym_east),
	467:  uint16(anon_sym_south),
	468:  uint16(anon_sym_west),
	469:  uint16(anon_sym_var),
	470:  uint16(anon_sym_horizontal),
	471:  uint16(anon_sym_vertical),
	472:  uint16(anon_sym_default),
	473:  uint16(anon_sym_pointer),
	474:  uint16(anon_sym_text),
	475:  uint16(anon_sym_env),
	476:  uint16(sym_identifier),
	477:  uint16(9),
	478:  uint16(3),
	479:  uint16(1),
	480:  uint16(sym_comment),
	481:  uint16(41),
	482:  uint16(1),
	483:  uint16(anon_sym_0),
	484:  uint16(55),
	485:  uint16(1),
	486:  uint16(anon_sym_calc),
	487:  uint16(268),
	488:  uint16(1),
	489:  uint16(sym_integer_value),
	490:  uint16(270),
	491:  uint16(1),
	492:  uint16(sym_float_value),
	493:  uint16(43),
	494:  uint16(1),
	495:  uint16(sym_distance_calc),
	496:  uint16(49),
	497:  uint16(1),
	498:  uint16(sym_distance_value),
	499:  uint16(286),
	500:  uint16(8),
	501:  uint16(anon_sym_COMMA),
	502:  uint16(anon_sym_POUND),
	503:  uint16(anon_sym_SEMI),
	504:  uint16(anon_sym_RPAREN),
	505:  uint16(anon_sym_DQUOTE),
	506:  uint16(anon_sym_AT),
	507:  uint16(anon_sym_LBRACK),
	508:  uint16(anon_sym_DOLLAR),
	509:  uint16(284),
	510:  uint16(32),
	511:  uint16(anon_sym_inherit),
	512:  uint16(anon_sym_DMENU),
	513:  uint16(anon_sym_true),
	514:  uint16(anon_sym_false),
	515:  uint16(anon_sym_url),
	516:  uint16(anon_sym_linear_DASHgradient),
	517:  uint16(anon_sym_rgb),
	518:  uint16(anon_sym_rgba),
	519:  uint16(anon_sym_hsl),
	520:  uint16(anon_sym_hsla),
	521:  uint16(anon_sym_hwb),
	522:  uint16(anon_sym_hwba),
	523:  uint16(anon_sym_cmyk),
	524:  uint16(anon_sym_bold),
	525:  uint16(anon_sym_italic),
	526:  uint16(anon_sym_underline),
	527:  uint16(anon_sym_strikethrough),
	528:  uint16(anon_sym_dash),
	529:  uint16(anon_sym_solid),
	530:  uint16(anon_sym_center),
	531:  uint16(anon_sym_north),
	532:  uint16(anon_sym_east),
	533:  uint16(anon_sym_south),
	534:  uint16(anon_sym_west),
	535:  uint16(anon_sym_var),
	536:  uint16(anon_sym_horizontal),
	537:  uint16(anon_sym_vertical),
	538:  uint16(anon_sym_default),
	539:  uint16(anon_sym_pointer),
	540:  uint16(anon_sym_text),
	541:  uint16(anon_sym_env),
	542:  uint16(sym_identifier),
	543:  uint16(9),
	544:  uint16(3),
	545:  uint16(1),
	546:  uint16(sym_comment),
	547:  uint16(41),
	548:  uint16(1),
	549:  uint16(anon_sym_0),
	550:  uint16(55),
	551:  uint16(1),
	552:  uint16(anon_sym_calc),
	553:  uint16(268),
	554:  uint16(1),
	555:  uint16(sym_integer_value),
	556:  uint16(270),
	557:  uint16(1),
	558:  uint16(sym_float_value),
	559:  uint16(19),
	560:  uint16(1),
	561:  uint16(sym_distance_value),
	562:  uint16(43),
	563:  uint16(1),
	564:  uint16(sym_distance_calc),
	565:  uint16(290),
	566:  uint16(8),
	567:  uint16(anon_sym_COMMA),
	568:  uint16(anon_sym_POUND),
	569:  uint16(anon_sym_SEMI),
	570:  uint16(anon_sym_RPAREN),
	571:  uint16(anon_sym_DQUOTE),
	572:  uint16(anon_sym_AT),
	573:  uint16(anon_sym_LBRACK),
	574:  uint16(anon_sym_DOLLAR),
	575:  uint16(288),
	576:  uint16(32),
	577:  uint16(anon_sym_inherit),
	578:  uint16(anon_sym_DMENU),
	579:  uint16(anon_sym_true),
	580:  uint16(anon_sym_false),
	581:  uint16(anon_sym_url),
	582:  uint16(anon_sym_linear_DASHgradient),
	583:  uint16(anon_sym_rgb),
	584:  uint16(anon_sym_rgba),
	585:  uint16(anon_sym_hsl),
	586:  uint16(anon_sym_hsla),
	587:  uint16(anon_sym_hwb),
	588:  uint16(anon_sym_hwba),
	589:  uint16(anon_sym_cmyk),
	590:  uint16(anon_sym_bold),
	591:  uint16(anon_sym_italic),
	592:  uint16(anon_sym_underline),
	593:  uint16(anon_sym_strikethrough),
	594:  uint16(anon_sym_dash),
	595:  uint16(anon_sym_solid),
	596:  uint16(anon_sym_center),
	597:  uint16(anon_sym_north),
	598:  uint16(anon_sym_east),
	599:  uint16(anon_sym_south),
	600:  uint16(anon_sym_west),
	601:  uint16(anon_sym_var),
	602:  uint16(anon_sym_horizontal),
	603:  uint16(anon_sym_vertical),
	604:  uint16(anon_sym_default),
	605:  uint16(anon_sym_pointer),
	606:  uint16(anon_sym_text),
	607:  uint16(anon_sym_env),
	608:  uint16(sym_identifier),
	609:  uint16(5),
	610:  uint16(3),
	611:  uint16(1),
	612:  uint16(sym_comment),
	613:  uint16(23),
	614:  uint16(1),
	615:  uint16(aux_sym_position_value_repeat1),
	616:  uint16(296),
	617:  uint16(5),
	618:  uint16(anon_sym_center),
	619:  uint16(anon_sym_north),
	620:  uint16(anon_sym_east),
	621:  uint16(anon_sym_south),
	622:  uint16(anon_sym_west),
	623:  uint16(294),
	624:  uint16(9),
	625:  uint16(anon_sym_COMMA),
	626:  uint16(anon_sym_POUND),
	627:  uint16(anon_sym_SEMI),
	628:  uint16(anon_sym_RPAREN),
	629:  uint16(anon_sym_DQUOTE),
	630:  uint16(sym_float_value),
	631:  uint16(anon_sym_AT),
	632:  uint16(anon_sym_LBRACK),
	633:  uint16(anon_sym_DOLLAR),
	634:  uint16(292),
	635:  uint16(30),
	636:  uint16(anon_sym_inherit),
	637:  uint16(anon_sym_DMENU),
	638:  uint16(sym_integer_value),
	639:  uint16(anon_sym_true),
	640:  uint16(anon_sym_false),
	641:  uint16(anon_sym_url),
	642:  uint16(anon_sym_linear_DASHgradient),
	643:  uint16(anon_sym_0),
	644:  uint16(anon_sym_rgb),
	645:  uint16(anon_sym_rgba),
	646:  uint16(anon_sym_hsl),
	647:  uint16(anon_sym_hsla),
	648:  uint16(anon_sym_hwb),
	649:  uint16(anon_sym_hwba),
	650:  uint16(anon_sym_cmyk),
	651:  uint16(anon_sym_bold),
	652:  uint16(anon_sym_italic),
	653:  uint16(anon_sym_underline),
	654:  uint16(anon_sym_strikethrough),
	655:  uint16(anon_sym_dash),
	656:  uint16(anon_sym_solid),
	657:  uint16(anon_sym_calc),
	658:  uint16(anon_sym_var),
	659:  uint16(anon_sym_horizontal),
	660:  uint16(anon_sym_vertical),
	661:  uint16(anon_sym_default),
	662:  uint16(anon_sym_pointer),
	663:  uint16(anon_sym_text),
	664:  uint16(anon_sym_env),
	665:  uint16(sym_identifier),
	666:  uint16(5),
	667:  uint16(3),
	668:  uint16(1),
	669:  uint16(sym_comment),
	670:  uint16(42),
	671:  uint16(1),
	672:  uint16(sym_line_style_value),
	673:  uint16(53),
	674:  uint16(2),
	675:  uint16(anon_sym_dash),
	676:  uint16(anon_sym_solid),
	677:  uint16(300),
	678:  uint16(9),
	679:  uint16(anon_sym_COMMA),
	680:  uint16(anon_sym_POUND),
	681:  uint16(anon_sym_SEMI),
	682:  uint16(anon_sym_RPAREN),
	683:  uint16(anon_sym_DQUOTE),
	684:  uint16(sym_float_value),
	685:  uint16(anon_sym_AT),
	686:  uint16(anon_sym_LBRACK),
	687:  uint16(anon_sym_DOLLAR),
	688:  uint16(298),
	689:  uint16(33),
	690:  uint16(anon_sym_inherit),
	691:  uint16(anon_sym_DMENU),
	692:  uint16(sym_integer_value),
	693:  uint16(anon_sym_true),
	694:  uint16(anon_sym_false),
	695:  uint16(anon_sym_url),
	696:  uint16(anon_sym_linear_DASHgradient),
	697:  uint16(anon_sym_0),
	698:  uint16(anon_sym_rgb),
	699:  uint16(anon_sym_rgba),
	700:  uint16(anon_sym_hsl),
	701:  uint16(anon_sym_hsla),
	702:  uint16(anon_sym_hwb),
	703:  uint16(anon_sym_hwba),
	704:  uint16(anon_sym_cmyk),
	705:  uint16(anon_sym_bold),
	706:  uint16(anon_sym_italic),
	707:  uint16(anon_sym_underline),
	708:  uint16(anon_sym_strikethrough),
	709:  uint16(anon_sym_calc),
	710:  uint16(anon_sym_center),
	711:  uint16(anon_sym_north),
	712:  uint16(anon_sym_east),
	713:  uint16(anon_sym_south),
	714:  uint16(anon_sym_west),
	715:  uint16(anon_sym_var),
	716:  uint16(anon_sym_horizontal),
	717:  uint16(anon_sym_vertical),
	718:  uint16(anon_sym_default),
	719:  uint16(anon_sym_pointer),
	720:  uint16(anon_sym_text),
	721:  uint16(anon_sym_env),
	722:  uint16(sym_identifier),
	723:  uint16(5),
	724:  uint16(3),
	725:  uint16(1),
	726:  uint16(sym_comment),
	727:  uint16(23),
	728:  uint16(1),
	729:  uint16(aux_sym_position_value_repeat1),
	730:  uint16(306),
	731:  uint16(5),
	732:  uint16(anon_sym_center),
	733:  uint16(anon_sym_north),
	734:  uint16(anon_sym_east),
	735:  uint16(anon_sym_south),
	736:  uint16(anon_sym_west),
	737:  uint16(304),
	738:  uint16(9),
	739:  uint16(anon_sym_COMMA),
	740:  uint16(anon_sym_POUND),
	741:  uint16(anon_sym_SEMI),
	742:  uint16(anon_sym_RPAREN),
	743:  uint16(anon_sym_DQUOTE),
	744:  uint16(sym_float_value),
	745:  uint16(anon_sym_AT),
	746:  uint16(anon_sym_LBRACK),
	747:  uint16(anon_sym_DOLLAR),
	748:  uint16(302),
	749:  uint16(30),
	750:  uint16(anon_sym_inherit),
	751:  uint16(anon_sym_DMENU),
	752:  uint16(sym_integer_value),
	753:  uint16(anon_sym_true),
	754:  uint16(anon_sym_false),
	755:  uint16(anon_sym_url),
	756:  uint16(anon_sym_linear_DASHgradient),
	757:  uint16(anon_sym_0),
	758:  uint16(anon_sym_rgb),
	759:  uint16(anon_sym_rgba),
	760:  uint16(anon_sym_hsl),
	761:  uint16(anon_sym_hsla),
	762:  uint16(anon_sym_hwb),
	763:  uint16(anon_sym_hwba),
	764:  uint16(anon_sym_cmyk),
	765:  uint16(anon_sym_bold),
	766:  uint16(anon_sym_italic),
	767:  uint16(anon_sym_underline),
	768:  uint16(anon_sym_strikethrough),
	769:  uint16(anon_sym_dash),
	770:  uint16(anon_sym_solid),
	771:  uint16(anon_sym_calc),
	772:  uint16(anon_sym_var),
	773:  uint16(anon_sym_horizontal),
	774:  uint16(anon_sym_vertical),
	775:  uint16(anon_sym_default),
	776:  uint16(anon_sym_pointer),
	777:  uint16(anon_sym_text),
	778:  uint16(anon_sym_env),
	779:  uint16(sym_identifier),
	780:  uint16(4),
	781:  uint16(3),
	782:  uint16(1),
	783:  uint16(sym_comment),
	784:  uint16(313),
	785:  uint16(1),
	786:  uint16(anon_sym_SLASH),
	787:  uint16(311),
	788:  uint16(9),
	789:  uint16(anon_sym_COMMA),
	790:  uint16(anon_sym_POUND),
	791:  uint16(anon_sym_SEMI),
	792:  uint16(anon_sym_RPAREN),
	793:  uint16(anon_sym_DQUOTE),
	794:  uint16(sym_float_value),
	795:  uint16(anon_sym_AT),
	796:  uint16(anon_sym_LBRACK),
	797:  uint16(anon_sym_DOLLAR),
	798:  uint16(309),
	799:  uint16(35),
	800:  uint16(anon_sym_inherit),
	801:  uint16(anon_sym_DMENU),
	802:  uint16(sym_integer_value),
	803:  uint16(anon_sym_true),
	804:  uint16(anon_sym_false),
	805:  uint16(anon_sym_url),
	806:  uint16(anon_sym_linear_DASHgradient),
	807:  uint16(anon_sym_0),
	808:  uint16(anon_sym_rgb),
	809:  uint16(anon_sym_rgba),
	810:  uint16(anon_sym_hsl),
	811:  uint16(anon_sym_hsla),
	812:  uint16(anon_sym_hwb),
	813:  uint16(anon_sym_hwba),
	814:  uint16(anon_sym_cmyk),
	815:  uint16(anon_sym_bold),
	816:  uint16(anon_sym_italic),
	817:  uint16(anon_sym_underline),
	818:  uint16(anon_sym_strikethrough),
	819:  uint16(anon_sym_dash),
	820:  uint16(anon_sym_solid),
	821:  uint16(anon_sym_calc),
	822:  uint16(anon_sym_center),
	823:  uint16(anon_sym_north),
	824:  uint16(anon_sym_east),
	825:  uint16(anon_sym_south),
	826:  uint16(anon_sym_west),
	827:  uint16(anon_sym_var),
	828:  uint16(anon_sym_horizontal),
	829:  uint16(anon_sym_vertical),
	830:  uint16(anon_sym_default),
	831:  uint16(anon_sym_pointer),
	832:  uint16(anon_sym_text),
	833:  uint16(anon_sym_env),
	834:  uint16(sym_identifier),
	835:  uint16(4),
	836:  uint16(3),
	837:  uint16(1),
	838:  uint16(sym_comment),
	839:  uint16(319),
	840:  uint16(1),
	841:  uint16(anon_sym_PERCENT),
	842:  uint16(317),
	843:  uint16(9),
	844:  uint16(anon_sym_COMMA),
	845:  uint16(anon_sym_POUND),
	846:  uint16(anon_sym_SEMI),
	847:  uint16(anon_sym_RPAREN),
	848:  uint16(anon_sym_DQUOTE),
	849:  uint16(sym_float_value),
	850:  uint16(anon_sym_AT),
	851:  uint16(anon_sym_LBRACK),
	852:  uint16(anon_sym_DOLLAR),
	853:  uint16(315),
	854:  uint16(35),
	855:  uint16(anon_sym_inherit),
	856:  uint16(anon_sym_DMENU),
	857:  uint16(sym_integer_value),
	858:  uint16(anon_sym_true),
	859:  uint16(anon_sym_false),
	860:  uint16(anon_sym_url),
	861:  uint16(anon_sym_linear_DASHgradient),
	862:  uint16(anon_sym_0),
	863:  uint16(anon_sym_rgb),
	864:  uint16(anon_sym_rgba),
	865:  uint16(anon_sym_hsl),
	866:  uint16(anon_sym_hsla),
	867:  uint16(anon_sym_hwb),
	868:  uint16(anon_sym_hwba),
	869:  uint16(anon_sym_cmyk),
	870:  uint16(anon_sym_bold),
	871:  uint16(anon_sym_italic),
	872:  uint16(anon_sym_underline),
	873:  uint16(anon_sym_strikethrough),
	874:  uint16(anon_sym_dash),
	875:  uint16(anon_sym_solid),
	876:  uint16(anon_sym_calc),
	877:  uint16(anon_sym_center),
	878:  uint16(anon_sym_north),
	879:  uint16(anon_sym_east),
	880:  uint16(anon_sym_south),
	881:  uint16(anon_sym_west),
	882:  uint16(anon_sym_var),
	883:  uint16(anon_sym_horizontal),
	884:  uint16(anon_sym_vertical),
	885:  uint16(anon_sym_default),
	886:  uint16(anon_sym_pointer),
	887:  uint16(anon_sym_text),
	888:  uint16(anon_sym_env),
	889:  uint16(sym_identifier),
	890:  uint16(3),
	891:  uint16(3),
	892:  uint16(1),
	893:  uint16(sym_comment),
	894:  uint16(323),
	895:  uint16(9),
	896:  uint16(anon_sym_COMMA),
	897:  uint16(anon_sym_POUND),
	898:  uint16(anon_sym_SEMI),
	899:  uint16(anon_sym_RPAREN),
	900:  uint16(anon_sym_DQUOTE),
	901:  uint16(sym_float_value),
	902:  uint16(anon_sym_AT),
	903:  uint16(anon_sym_LBRACK),
	904:  uint16(anon_sym_DOLLAR),
	905:  uint16(321),
	906:  uint16(35),
	907:  uint16(anon_sym_inherit),
	908:  uint16(anon_sym_DMENU),
	909:  uint16(sym_integer_value),
	910:  uint16(anon_sym_true),
	911:  uint16(anon_sym_false),
	912:  uint16(anon_sym_url),
	913:  uint16(anon_sym_linear_DASHgradient),
	914:  uint16(anon_sym_0),
	915:  uint16(anon_sym_rgb),
	916:  uint16(anon_sym_rgba),
	917:  uint16(anon_sym_hsl),
	918:  uint16(anon_sym_hsla),
	919:  uint16(anon_sym_hwb),
	920:  uint16(anon_sym_hwba),
	921:  uint16(anon_sym_cmyk),
	922:  uint16(anon_sym_bold),
	923:  uint16(anon_sym_italic),
	924:  uint16(anon_sym_underline),
	925:  uint16(anon_sym_strikethrough),
	926:  uint16(anon_sym_dash),
	927:  uint16(anon_sym_solid),
	928:  uint16(anon_sym_calc),
	929:  uint16(anon_sym_center),
	930:  uint16(anon_sym_north),
	931:  uint16(anon_sym_east),
	932:  uint16(anon_sym_south),
	933:  uint16(anon_sym_west),
	934:  uint16(anon_sym_var),
	935:  uint16(anon_sym_horizontal),
	936:  uint16(anon_sym_vertical),
	937:  uint16(anon_sym_default),
	938:  uint16(anon_sym_pointer),
	939:  uint16(anon_sym_text),
	940:  uint16(anon_sym_env),
	941:  uint16(sym_identifier),
	942:  uint16(3),
	943:  uint16(3),
	944:  uint16(1),
	945:  uint16(sym_comment),
	946:  uint16(327),
	947:  uint16(9),
	948:  uint16(anon_sym_COMMA),
	949:  uint16(anon_sym_POUND),
	950:  uint16(anon_sym_SEMI),
	951:  uint16(anon_sym_RPAREN),
	952:  uint16(anon_sym_DQUOTE),
	953:  uint16(sym_float_value),
	954:  uint16(anon_sym_AT),
	955:  uint16(anon_sym_LBRACK),
	956:  uint16(anon_sym_DOLLAR),
	957:  uint16(325),
	958:  uint16(35),
	959:  uint16(anon_sym_inherit),
	960:  uint16(anon_sym_DMENU),
	961:  uint16(sym_integer_value),
	962:  uint16(anon_sym_true),
	963:  uint16(anon_sym_false),
	964:  uint16(anon_sym_url),
	965:  uint16(anon_sym_linear_DASHgradient),
	966:  uint16(anon_sym_0),
	967:  uint16(anon_sym_rgb),
	968:  uint16(anon_sym_rgba),
	969:  uint16(anon_sym_hsl),
	970:  uint16(anon_sym_hsla),
	971:  uint16(anon_sym_hwb),
	972:  uint16(anon_sym_hwba),
	973:  uint16(anon_sym_cmyk),
	974:  uint16(anon_sym_bold),
	975:  uint16(anon_sym_italic),
	976:  uint16(anon_sym_underline),
	977:  uint16(anon_sym_strikethrough),
	978:  uint16(anon_sym_dash),
	979:  uint16(anon_sym_solid),
	980:  uint16(anon_sym_calc),
	981:  uint16(anon_sym_center),
	982:  uint16(anon_sym_north),
	983:  uint16(anon_sym_east),
	984:  uint16(anon_sym_south),
	985:  uint16(anon_sym_west),
	986:  uint16(anon_sym_var),
	987:  uint16(anon_sym_horizontal),
	988:  uint16(anon_sym_vertical),
	989:  uint16(anon_sym_default),
	990:  uint16(anon_sym_pointer),
	991:  uint16(anon_sym_text),
	992:  uint16(anon_sym_env),
	993:  uint16(sym_identifier),
	994:  uint16(3),
	995:  uint16(3),
	996:  uint16(1),
	997:  uint16(sym_comment),
	998:  uint16(331),
	999:  uint16(9),
	1000: uint16(anon_sym_COMMA),
	1001: uint16(anon_sym_POUND),
	1002: uint16(anon_sym_SEMI),
	1003: uint16(anon_sym_RPAREN),
	1004: uint16(anon_sym_DQUOTE),
	1005: uint16(sym_float_value),
	1006: uint16(anon_sym_AT),
	1007: uint16(anon_sym_LBRACK),
	1008: uint16(anon_sym_DOLLAR),
	1009: uint16(329),
	1010: uint16(35),
	1011: uint16(anon_sym_inherit),
	1012: uint16(anon_sym_DMENU),
	1013: uint16(sym_integer_value),
	1014: uint16(anon_sym_true),
	1015: uint16(anon_sym_false),
	1016: uint16(anon_sym_url),
	1017: uint16(anon_sym_linear_DASHgradient),
	1018: uint16(anon_sym_0),
	1019: uint16(anon_sym_rgb),
	1020: uint16(anon_sym_rgba),
	1021: uint16(anon_sym_hsl),
	1022: uint16(anon_sym_hsla),
	1023: uint16(anon_sym_hwb),
	1024: uint16(anon_sym_hwba),
	1025: uint16(anon_sym_cmyk),
	1026: uint16(anon_sym_bold),
	1027: uint16(anon_sym_italic),
	1028: uint16(anon_sym_underline),
	1029: uint16(anon_sym_strikethrough),
	1030: uint16(anon_sym_dash),
	1031: uint16(anon_sym_solid),
	1032: uint16(anon_sym_calc),
	1033: uint16(anon_sym_center),
	1034: uint16(anon_sym_north),
	1035: uint16(anon_sym_east),
	1036: uint16(anon_sym_south),
	1037: uint16(anon_sym_west),
	1038: uint16(anon_sym_var),
	1039: uint16(anon_sym_horizontal),
	1040: uint16(anon_sym_vertical),
	1041: uint16(anon_sym_default),
	1042: uint16(anon_sym_pointer),
	1043: uint16(anon_sym_text),
	1044: uint16(anon_sym_env),
	1045: uint16(sym_identifier),
	1046: uint16(3),
	1047: uint16(3),
	1048: uint16(1),
	1049: uint16(sym_comment),
	1050: uint16(335),
	1051: uint16(9),
	1052: uint16(anon_sym_COMMA),
	1053: uint16(anon_sym_POUND),
	1054: uint16(anon_sym_SEMI),
	1055: uint16(anon_sym_RPAREN),
	1056: uint16(anon_sym_DQUOTE),
	1057: uint16(sym_float_value),
	1058: uint16(anon_sym_AT),
	1059: uint16(anon_sym_LBRACK),
	1060: uint16(anon_sym_DOLLAR),
	1061: uint16(333),
	1062: uint16(35),
	1063: uint16(anon_sym_inherit),
	1064: uint16(anon_sym_DMENU),
	1065: uint16(sym_integer_value),
	1066: uint16(anon_sym_true),
	1067: uint16(anon_sym_false),
	1068: uint16(anon_sym_url),
	1069: uint16(anon_sym_linear_DASHgradient),
	1070: uint16(anon_sym_0),
	1071: uint16(anon_sym_rgb),
	1072: uint16(anon_sym_rgba),
	1073: uint16(anon_sym_hsl),
	1074: uint16(anon_sym_hsla),
	1075: uint16(anon_sym_hwb),
	1076: uint16(anon_sym_hwba),
	1077: uint16(anon_sym_cmyk),
	1078: uint16(anon_sym_bold),
	1079: uint16(anon_sym_italic),
	1080: uint16(anon_sym_underline),
	1081: uint16(anon_sym_strikethrough),
	1082: uint16(anon_sym_dash),
	1083: uint16(anon_sym_solid),
	1084: uint16(anon_sym_calc),
	1085: uint16(anon_sym_center),
	1086: uint16(anon_sym_north),
	1087: uint16(anon_sym_east),
	1088: uint16(anon_sym_south),
	1089: uint16(anon_sym_west),
	1090: uint16(anon_sym_var),
	1091: uint16(anon_sym_horizontal),
	1092: uint16(anon_sym_vertical),
	1093: uint16(anon_sym_default),
	1094: uint16(anon_sym_pointer),
	1095: uint16(anon_sym_text),
	1096: uint16(anon_sym_env),
	1097: uint16(sym_identifier),
	1098: uint16(3),
	1099: uint16(3),
	1100: uint16(1),
	1101: uint16(sym_comment),
	1102: uint16(339),
	1103: uint16(9),
	1104: uint16(anon_sym_COMMA),
	1105: uint16(anon_sym_POUND),
	1106: uint16(anon_sym_SEMI),
	1107: uint16(anon_sym_RPAREN),
	1108: uint16(anon_sym_DQUOTE),
	1109: uint16(sym_float_value),
	1110: uint16(anon_sym_AT),
	1111: uint16(anon_sym_LBRACK),
	1112: uint16(anon_sym_DOLLAR),
	1113: uint16(337),
	1114: uint16(35),
	1115: uint16(anon_sym_inherit),
	1116: uint16(anon_sym_DMENU),
	1117: uint16(sym_integer_value),
	1118: uint16(anon_sym_true),
	1119: uint16(anon_sym_false),
	1120: uint16(anon_sym_url),
	1121: uint16(anon_sym_linear_DASHgradient),
	1122: uint16(anon_sym_0),
	1123: uint16(anon_sym_rgb),
	1124: uint16(anon_sym_rgba),
	1125: uint16(anon_sym_hsl),
	1126: uint16(anon_sym_hsla),
	1127: uint16(anon_sym_hwb),
	1128: uint16(anon_sym_hwba),
	1129: uint16(anon_sym_cmyk),
	1130: uint16(anon_sym_bold),
	1131: uint16(anon_sym_italic),
	1132: uint16(anon_sym_underline),
	1133: uint16(anon_sym_strikethrough),
	1134: uint16(anon_sym_dash),
	1135: uint16(anon_sym_solid),
	1136: uint16(anon_sym_calc),
	1137: uint16(anon_sym_center),
	1138: uint16(anon_sym_north),
	1139: uint16(anon_sym_east),
	1140: uint16(anon_sym_south),
	1141: uint16(anon_sym_west),
	1142: uint16(anon_sym_var),
	1143: uint16(anon_sym_horizontal),
	1144: uint16(anon_sym_vertical),
	1145: uint16(anon_sym_default),
	1146: uint16(anon_sym_pointer),
	1147: uint16(anon_sym_text),
	1148: uint16(anon_sym_env),
	1149: uint16(sym_identifier),
	1150: uint16(3),
	1151: uint16(3),
	1152: uint16(1),
	1153: uint16(sym_comment),
	1154: uint16(343),
	1155: uint16(9),
	1156: uint16(anon_sym_COMMA),
	1157: uint16(anon_sym_POUND),
	1158: uint16(anon_sym_SEMI),
	1159: uint16(anon_sym_RPAREN),
	1160: uint16(anon_sym_DQUOTE),
	1161: uint16(sym_float_value),
	1162: uint16(anon_sym_AT),
	1163: uint16(anon_sym_LBRACK),
	1164: uint16(anon_sym_DOLLAR),
	1165: uint16(341),
	1166: uint16(35),
	1167: uint16(anon_sym_inherit),
	1168: uint16(anon_sym_DMENU),
	1169: uint16(sym_integer_value),
	1170: uint16(anon_sym_true),
	1171: uint16(anon_sym_false),
	1172: uint16(anon_sym_url),
	1173: uint16(anon_sym_linear_DASHgradient),
	1174: uint16(anon_sym_0),
	1175: uint16(anon_sym_rgb),
	1176: uint16(anon_sym_rgba),
	1177: uint16(anon_sym_hsl),
	1178: uint16(anon_sym_hsla),
	1179: uint16(anon_sym_hwb),
	1180: uint16(anon_sym_hwba),
	1181: uint16(anon_sym_cmyk),
	1182: uint16(anon_sym_bold),
	1183: uint16(anon_sym_italic),
	1184: uint16(anon_sym_underline),
	1185: uint16(anon_sym_strikethrough),
	1186: uint16(anon_sym_dash),
	1187: uint16(anon_sym_solid),
	1188: uint16(anon_sym_calc),
	1189: uint16(anon_sym_center),
	1190: uint16(anon_sym_north),
	1191: uint16(anon_sym_east),
	1192: uint16(anon_sym_south),
	1193: uint16(anon_sym_west),
	1194: uint16(anon_sym_var),
	1195: uint16(anon_sym_horizontal),
	1196: uint16(anon_sym_vertical),
	1197: uint16(anon_sym_default),
	1198: uint16(anon_sym_pointer),
	1199: uint16(anon_sym_text),
	1200: uint16(anon_sym_env),
	1201: uint16(sym_identifier),
	1202: uint16(3),
	1203: uint16(3),
	1204: uint16(1),
	1205: uint16(sym_comment),
	1206: uint16(347),
	1207: uint16(9),
	1208: uint16(anon_sym_COMMA),
	1209: uint16(anon_sym_POUND),
	1210: uint16(anon_sym_SEMI),
	1211: uint16(anon_sym_RPAREN),
	1212: uint16(anon_sym_DQUOTE),
	1213: uint16(sym_float_value),
	1214: uint16(anon_sym_AT),
	1215: uint16(anon_sym_LBRACK),
	1216: uint16(anon_sym_DOLLAR),
	1217: uint16(345),
	1218: uint16(35),
	1219: uint16(anon_sym_inherit),
	1220: uint16(anon_sym_DMENU),
	1221: uint16(sym_integer_value),
	1222: uint16(anon_sym_true),
	1223: uint16(anon_sym_false),
	1224: uint16(anon_sym_url),
	1225: uint16(anon_sym_linear_DASHgradient),
	1226: uint16(anon_sym_0),
	1227: uint16(anon_sym_rgb),
	1228: uint16(anon_sym_rgba),
	1229: uint16(anon_sym_hsl),
	1230: uint16(anon_sym_hsla),
	1231: uint16(anon_sym_hwb),
	1232: uint16(anon_sym_hwba),
	1233: uint16(anon_sym_cmyk),
	1234: uint16(anon_sym_bold),
	1235: uint16(anon_sym_italic),
	1236: uint16(anon_sym_underline),
	1237: uint16(anon_sym_strikethrough),
	1238: uint16(anon_sym_dash),
	1239: uint16(anon_sym_solid),
	1240: uint16(anon_sym_calc),
	1241: uint16(anon_sym_center),
	1242: uint16(anon_sym_north),
	1243: uint16(anon_sym_east),
	1244: uint16(anon_sym_south),
	1245: uint16(anon_sym_west),
	1246: uint16(anon_sym_var),
	1247: uint16(anon_sym_horizontal),
	1248: uint16(anon_sym_vertical),
	1249: uint16(anon_sym_default),
	1250: uint16(anon_sym_pointer),
	1251: uint16(anon_sym_text),
	1252: uint16(anon_sym_env),
	1253: uint16(sym_identifier),
	1254: uint16(3),
	1255: uint16(3),
	1256: uint16(1),
	1257: uint16(sym_comment),
	1258: uint16(351),
	1259: uint16(9),
	1260: uint16(anon_sym_COMMA),
	1261: uint16(anon_sym_POUND),
	1262: uint16(anon_sym_SEMI),
	1263: uint16(anon_sym_RPAREN),
	1264: uint16(anon_sym_DQUOTE),
	1265: uint16(sym_float_value),
	1266: uint16(anon_sym_AT),
	1267: uint16(anon_sym_LBRACK),
	1268: uint16(anon_sym_DOLLAR),
	1269: uint16(349),
	1270: uint16(35),
	1271: uint16(anon_sym_inherit),
	1272: uint16(anon_sym_DMENU),
	1273: uint16(sym_integer_value),
	1274: uint16(anon_sym_true),
	1275: uint16(anon_sym_false),
	1276: uint16(anon_sym_url),
	1277: uint16(anon_sym_linear_DASHgradient),
	1278: uint16(anon_sym_0),
	1279: uint16(anon_sym_rgb),
	1280: uint16(anon_sym_rgba),
	1281: uint16(anon_sym_hsl),
	1282: uint16(anon_sym_hsla),
	1283: uint16(anon_sym_hwb),
	1284: uint16(anon_sym_hwba),
	1285: uint16(anon_sym_cmyk),
	1286: uint16(anon_sym_bold),
	1287: uint16(anon_sym_italic),
	1288: uint16(anon_sym_underline),
	1289: uint16(anon_sym_strikethrough),
	1290: uint16(anon_sym_dash),
	1291: uint16(anon_sym_solid),
	1292: uint16(anon_sym_calc),
	1293: uint16(anon_sym_center),
	1294: uint16(anon_sym_north),
	1295: uint16(anon_sym_east),
	1296: uint16(anon_sym_south),
	1297: uint16(anon_sym_west),
	1298: uint16(anon_sym_var),
	1299: uint16(anon_sym_horizontal),
	1300: uint16(anon_sym_vertical),
	1301: uint16(anon_sym_default),
	1302: uint16(anon_sym_pointer),
	1303: uint16(anon_sym_text),
	1304: uint16(anon_sym_env),
	1305: uint16(sym_identifier),
	1306: uint16(3),
	1307: uint16(3),
	1308: uint16(1),
	1309: uint16(sym_comment),
	1310: uint16(355),
	1311: uint16(9),
	1312: uint16(anon_sym_COMMA),
	1313: uint16(anon_sym_POUND),
	1314: uint16(anon_sym_SEMI),
	1315: uint16(anon_sym_RPAREN),
	1316: uint16(anon_sym_DQUOTE),
	1317: uint16(sym_float_value),
	1318: uint16(anon_sym_AT),
	1319: uint16(anon_sym_LBRACK),
	1320: uint16(anon_sym_DOLLAR),
	1321: uint16(353),
	1322: uint16(35),
	1323: uint16(anon_sym_inherit),
	1324: uint16(anon_sym_DMENU),
	1325: uint16(sym_integer_value),
	1326: uint16(anon_sym_true),
	1327: uint16(anon_sym_false),
	1328: uint16(anon_sym_url),
	1329: uint16(anon_sym_linear_DASHgradient),
	1330: uint16(anon_sym_0),
	1331: uint16(anon_sym_rgb),
	1332: uint16(anon_sym_rgba),
	1333: uint16(anon_sym_hsl),
	1334: uint16(anon_sym_hsla),
	1335: uint16(anon_sym_hwb),
	1336: uint16(anon_sym_hwba),
	1337: uint16(anon_sym_cmyk),
	1338: uint16(anon_sym_bold),
	1339: uint16(anon_sym_italic),
	1340: uint16(anon_sym_underline),
	1341: uint16(anon_sym_strikethrough),
	1342: uint16(anon_sym_dash),
	1343: uint16(anon_sym_solid),
	1344: uint16(anon_sym_calc),
	1345: uint16(anon_sym_center),
	1346: uint16(anon_sym_north),
	1347: uint16(anon_sym_east),
	1348: uint16(anon_sym_south),
	1349: uint16(anon_sym_west),
	1350: uint16(anon_sym_var),
	1351: uint16(anon_sym_horizontal),
	1352: uint16(anon_sym_vertical),
	1353: uint16(anon_sym_default),
	1354: uint16(anon_sym_pointer),
	1355: uint16(anon_sym_text),
	1356: uint16(anon_sym_env),
	1357: uint16(sym_identifier),
	1358: uint16(3),
	1359: uint16(3),
	1360: uint16(1),
	1361: uint16(sym_comment),
	1362: uint16(359),
	1363: uint16(9),
	1364: uint16(anon_sym_COMMA),
	1365: uint16(anon_sym_POUND),
	1366: uint16(anon_sym_SEMI),
	1367: uint16(anon_sym_RPAREN),
	1368: uint16(anon_sym_DQUOTE),
	1369: uint16(sym_float_value),
	1370: uint16(anon_sym_AT),
	1371: uint16(anon_sym_LBRACK),
	1372: uint16(anon_sym_DOLLAR),
	1373: uint16(357),
	1374: uint16(35),
	1375: uint16(anon_sym_inherit),
	1376: uint16(anon_sym_DMENU),
	1377: uint16(sym_integer_value),
	1378: uint16(anon_sym_true),
	1379: uint16(anon_sym_false),
	1380: uint16(anon_sym_url),
	1381: uint16(anon_sym_linear_DASHgradient),
	1382: uint16(anon_sym_0),
	1383: uint16(anon_sym_rgb),
	1384: uint16(anon_sym_rgba),
	1385: uint16(anon_sym_hsl),
	1386: uint16(anon_sym_hsla),
	1387: uint16(anon_sym_hwb),
	1388: uint16(anon_sym_hwba),
	1389: uint16(anon_sym_cmyk),
	1390: uint16(anon_sym_bold),
	1391: uint16(anon_sym_italic),
	1392: uint16(anon_sym_underline),
	1393: uint16(anon_sym_strikethrough),
	1394: uint16(anon_sym_dash),
	1395: uint16(anon_sym_solid),
	1396: uint16(anon_sym_calc),
	1397: uint16(anon_sym_center),
	1398: uint16(anon_sym_north),
	1399: uint16(anon_sym_east),
	1400: uint16(anon_sym_south),
	1401: uint16(anon_sym_west),
	1402: uint16(anon_sym_var),
	1403: uint16(anon_sym_horizontal),
	1404: uint16(anon_sym_vertical),
	1405: uint16(anon_sym_default),
	1406: uint16(anon_sym_pointer),
	1407: uint16(anon_sym_text),
	1408: uint16(anon_sym_env),
	1409: uint16(sym_identifier),
	1410: uint16(3),
	1411: uint16(3),
	1412: uint16(1),
	1413: uint16(sym_comment),
	1414: uint16(363),
	1415: uint16(9),
	1416: uint16(anon_sym_COMMA),
	1417: uint16(anon_sym_POUND),
	1418: uint16(anon_sym_SEMI),
	1419: uint16(anon_sym_RPAREN),
	1420: uint16(anon_sym_DQUOTE),
	1421: uint16(sym_float_value),
	1422: uint16(anon_sym_AT),
	1423: uint16(anon_sym_LBRACK),
	1424: uint16(anon_sym_DOLLAR),
	1425: uint16(361),
	1426: uint16(35),
	1427: uint16(anon_sym_inherit),
	1428: uint16(anon_sym_DMENU),
	1429: uint16(sym_integer_value),
	1430: uint16(anon_sym_true),
	1431: uint16(anon_sym_false),
	1432: uint16(anon_sym_url),
	1433: uint16(anon_sym_linear_DASHgradient),
	1434: uint16(anon_sym_0),
	1435: uint16(anon_sym_rgb),
	1436: uint16(anon_sym_rgba),
	1437: uint16(anon_sym_hsl),
	1438: uint16(anon_sym_hsla),
	1439: uint16(anon_sym_hwb),
	1440: uint16(anon_sym_hwba),
	1441: uint16(anon_sym_cmyk),
	1442: uint16(anon_sym_bold),
	1443: uint16(anon_sym_italic),
	1444: uint16(anon_sym_underline),
	1445: uint16(anon_sym_strikethrough),
	1446: uint16(anon_sym_dash),
	1447: uint16(anon_sym_solid),
	1448: uint16(anon_sym_calc),
	1449: uint16(anon_sym_center),
	1450: uint16(anon_sym_north),
	1451: uint16(anon_sym_east),
	1452: uint16(anon_sym_south),
	1453: uint16(anon_sym_west),
	1454: uint16(anon_sym_var),
	1455: uint16(anon_sym_horizontal),
	1456: uint16(anon_sym_vertical),
	1457: uint16(anon_sym_default),
	1458: uint16(anon_sym_pointer),
	1459: uint16(anon_sym_text),
	1460: uint16(anon_sym_env),
	1461: uint16(sym_identifier),
	1462: uint16(3),
	1463: uint16(3),
	1464: uint16(1),
	1465: uint16(sym_comment),
	1466: uint16(367),
	1467: uint16(9),
	1468: uint16(anon_sym_COMMA),
	1469: uint16(anon_sym_POUND),
	1470: uint16(anon_sym_SEMI),
	1471: uint16(anon_sym_RPAREN),
	1472: uint16(anon_sym_DQUOTE),
	1473: uint16(sym_float_value),
	1474: uint16(anon_sym_AT),
	1475: uint16(anon_sym_LBRACK),
	1476: uint16(anon_sym_DOLLAR),
	1477: uint16(365),
	1478: uint16(35),
	1479: uint16(anon_sym_inherit),
	1480: uint16(anon_sym_DMENU),
	1481: uint16(sym_integer_value),
	1482: uint16(anon_sym_true),
	1483: uint16(anon_sym_false),
	1484: uint16(anon_sym_url),
	1485: uint16(anon_sym_linear_DASHgradient),
	1486: uint16(anon_sym_0),
	1487: uint16(anon_sym_rgb),
	1488: uint16(anon_sym_rgba),
	1489: uint16(anon_sym_hsl),
	1490: uint16(anon_sym_hsla),
	1491: uint16(anon_sym_hwb),
	1492: uint16(anon_sym_hwba),
	1493: uint16(anon_sym_cmyk),
	1494: uint16(anon_sym_bold),
	1495: uint16(anon_sym_italic),
	1496: uint16(anon_sym_underline),
	1497: uint16(anon_sym_strikethrough),
	1498: uint16(anon_sym_dash),
	1499: uint16(anon_sym_solid),
	1500: uint16(anon_sym_calc),
	1501: uint16(anon_sym_center),
	1502: uint16(anon_sym_north),
	1503: uint16(anon_sym_east),
	1504: uint16(anon_sym_south),
	1505: uint16(anon_sym_west),
	1506: uint16(anon_sym_var),
	1507: uint16(anon_sym_horizontal),
	1508: uint16(anon_sym_vertical),
	1509: uint16(anon_sym_default),
	1510: uint16(anon_sym_pointer),
	1511: uint16(anon_sym_text),
	1512: uint16(anon_sym_env),
	1513: uint16(sym_identifier),
	1514: uint16(3),
	1515: uint16(3),
	1516: uint16(1),
	1517: uint16(sym_comment),
	1518: uint16(371),
	1519: uint16(9),
	1520: uint16(anon_sym_COMMA),
	1521: uint16(anon_sym_POUND),
	1522: uint16(anon_sym_SEMI),
	1523: uint16(anon_sym_RPAREN),
	1524: uint16(anon_sym_DQUOTE),
	1525: uint16(sym_float_value),
	1526: uint16(anon_sym_AT),
	1527: uint16(anon_sym_LBRACK),
	1528: uint16(anon_sym_DOLLAR),
	1529: uint16(369),
	1530: uint16(35),
	1531: uint16(anon_sym_inherit),
	1532: uint16(anon_sym_DMENU),
	1533: uint16(sym_integer_value),
	1534: uint16(anon_sym_true),
	1535: uint16(anon_sym_false),
	1536: uint16(anon_sym_url),
	1537: uint16(anon_sym_linear_DASHgradient),
	1538: uint16(anon_sym_0),
	1539: uint16(anon_sym_rgb),
	1540: uint16(anon_sym_rgba),
	1541: uint16(anon_sym_hsl),
	1542: uint16(anon_sym_hsla),
	1543: uint16(anon_sym_hwb),
	1544: uint16(anon_sym_hwba),
	1545: uint16(anon_sym_cmyk),
	1546: uint16(anon_sym_bold),
	1547: uint16(anon_sym_italic),
	1548: uint16(anon_sym_underline),
	1549: uint16(anon_sym_strikethrough),
	1550: uint16(anon_sym_dash),
	1551: uint16(anon_sym_solid),
	1552: uint16(anon_sym_calc),
	1553: uint16(anon_sym_center),
	1554: uint16(anon_sym_north),
	1555: uint16(anon_sym_east),
	1556: uint16(anon_sym_south),
	1557: uint16(anon_sym_west),
	1558: uint16(anon_sym_var),
	1559: uint16(anon_sym_horizontal),
	1560: uint16(anon_sym_vertical),
	1561: uint16(anon_sym_default),
	1562: uint16(anon_sym_pointer),
	1563: uint16(anon_sym_text),
	1564: uint16(anon_sym_env),
	1565: uint16(sym_identifier),
	1566: uint16(3),
	1567: uint16(3),
	1568: uint16(1),
	1569: uint16(sym_comment),
	1570: uint16(375),
	1571: uint16(9),
	1572: uint16(anon_sym_COMMA),
	1573: uint16(anon_sym_POUND),
	1574: uint16(anon_sym_SEMI),
	1575: uint16(anon_sym_RPAREN),
	1576: uint16(anon_sym_DQUOTE),
	1577: uint16(sym_float_value),
	1578: uint16(anon_sym_AT),
	1579: uint16(anon_sym_LBRACK),
	1580: uint16(anon_sym_DOLLAR),
	1581: uint16(373),
	1582: uint16(35),
	1583: uint16(anon_sym_inherit),
	1584: uint16(anon_sym_DMENU),
	1585: uint16(sym_integer_value),
	1586: uint16(anon_sym_true),
	1587: uint16(anon_sym_false),
	1588: uint16(anon_sym_url),
	1589: uint16(anon_sym_linear_DASHgradient),
	1590: uint16(anon_sym_0),
	1591: uint16(anon_sym_rgb),
	1592: uint16(anon_sym_rgba),
	1593: uint16(anon_sym_hsl),
	1594: uint16(anon_sym_hsla),
	1595: uint16(anon_sym_hwb),
	1596: uint16(anon_sym_hwba),
	1597: uint16(anon_sym_cmyk),
	1598: uint16(anon_sym_bold),
	1599: uint16(anon_sym_italic),
	1600: uint16(anon_sym_underline),
	1601: uint16(anon_sym_strikethrough),
	1602: uint16(anon_sym_dash),
	1603: uint16(anon_sym_solid),
	1604: uint16(anon_sym_calc),
	1605: uint16(anon_sym_center),
	1606: uint16(anon_sym_north),
	1607: uint16(anon_sym_east),
	1608: uint16(anon_sym_south),
	1609: uint16(anon_sym_west),
	1610: uint16(anon_sym_var),
	1611: uint16(anon_sym_horizontal),
	1612: uint16(anon_sym_vertical),
	1613: uint16(anon_sym_default),
	1614: uint16(anon_sym_pointer),
	1615: uint16(anon_sym_text),
	1616: uint16(anon_sym_env),
	1617: uint16(sym_identifier),
	1618: uint16(3),
	1619: uint16(3),
	1620: uint16(1),
	1621: uint16(sym_comment),
	1622: uint16(379),
	1623: uint16(9),
	1624: uint16(anon_sym_COMMA),
	1625: uint16(anon_sym_POUND),
	1626: uint16(anon_sym_SEMI),
	1627: uint16(anon_sym_RPAREN),
	1628: uint16(anon_sym_DQUOTE),
	1629: uint16(sym_float_value),
	1630: uint16(anon_sym_AT),
	1631: uint16(anon_sym_LBRACK),
	1632: uint16(anon_sym_DOLLAR),
	1633: uint16(377),
	1634: uint16(35),
	1635: uint16(anon_sym_inherit),
	1636: uint16(anon_sym_DMENU),
	1637: uint16(sym_integer_value),
	1638: uint16(anon_sym_true),
	1639: uint16(anon_sym_false),
	1640: uint16(anon_sym_url),
	1641: uint16(anon_sym_linear_DASHgradient),
	1642: uint16(anon_sym_0),
	1643: uint16(anon_sym_rgb),
	1644: uint16(anon_sym_rgba),
	1645: uint16(anon_sym_hsl),
	1646: uint16(anon_sym_hsla),
	1647: uint16(anon_sym_hwb),
	1648: uint16(anon_sym_hwba),
	1649: uint16(anon_sym_cmyk),
	1650: uint16(anon_sym_bold),
	1651: uint16(anon_sym_italic),
	1652: uint16(anon_sym_underline),
	1653: uint16(anon_sym_strikethrough),
	1654: uint16(anon_sym_dash),
	1655: uint16(anon_sym_solid),
	1656: uint16(anon_sym_calc),
	1657: uint16(anon_sym_center),
	1658: uint16(anon_sym_north),
	1659: uint16(anon_sym_east),
	1660: uint16(anon_sym_south),
	1661: uint16(anon_sym_west),
	1662: uint16(anon_sym_var),
	1663: uint16(anon_sym_horizontal),
	1664: uint16(anon_sym_vertical),
	1665: uint16(anon_sym_default),
	1666: uint16(anon_sym_pointer),
	1667: uint16(anon_sym_text),
	1668: uint16(anon_sym_env),
	1669: uint16(sym_identifier),
	1670: uint16(3),
	1671: uint16(3),
	1672: uint16(1),
	1673: uint16(sym_comment),
	1674: uint16(383),
	1675: uint16(9),
	1676: uint16(anon_sym_COMMA),
	1677: uint16(anon_sym_POUND),
	1678: uint16(anon_sym_SEMI),
	1679: uint16(anon_sym_RPAREN),
	1680: uint16(anon_sym_DQUOTE),
	1681: uint16(sym_float_value),
	1682: uint16(anon_sym_AT),
	1683: uint16(anon_sym_LBRACK),
	1684: uint16(anon_sym_DOLLAR),
	1685: uint16(381),
	1686: uint16(35),
	1687: uint16(anon_sym_inherit),
	1688: uint16(anon_sym_DMENU),
	1689: uint16(sym_integer_value),
	1690: uint16(anon_sym_true),
	1691: uint16(anon_sym_false),
	1692: uint16(anon_sym_url),
	1693: uint16(anon_sym_linear_DASHgradient),
	1694: uint16(anon_sym_0),
	1695: uint16(anon_sym_rgb),
	1696: uint16(anon_sym_rgba),
	1697: uint16(anon_sym_hsl),
	1698: uint16(anon_sym_hsla),
	1699: uint16(anon_sym_hwb),
	1700: uint16(anon_sym_hwba),
	1701: uint16(anon_sym_cmyk),
	1702: uint16(anon_sym_bold),
	1703: uint16(anon_sym_italic),
	1704: uint16(anon_sym_underline),
	1705: uint16(anon_sym_strikethrough),
	1706: uint16(anon_sym_dash),
	1707: uint16(anon_sym_solid),
	1708: uint16(anon_sym_calc),
	1709: uint16(anon_sym_center),
	1710: uint16(anon_sym_north),
	1711: uint16(anon_sym_east),
	1712: uint16(anon_sym_south),
	1713: uint16(anon_sym_west),
	1714: uint16(anon_sym_var),
	1715: uint16(anon_sym_horizontal),
	1716: uint16(anon_sym_vertical),
	1717: uint16(anon_sym_default),
	1718: uint16(anon_sym_pointer),
	1719: uint16(anon_sym_text),
	1720: uint16(anon_sym_env),
	1721: uint16(sym_identifier),
	1722: uint16(3),
	1723: uint16(3),
	1724: uint16(1),
	1725: uint16(sym_comment),
	1726: uint16(387),
	1727: uint16(9),
	1728: uint16(anon_sym_COMMA),
	1729: uint16(anon_sym_POUND),
	1730: uint16(anon_sym_SEMI),
	1731: uint16(anon_sym_RPAREN),
	1732: uint16(anon_sym_DQUOTE),
	1733: uint16(sym_float_value),
	1734: uint16(anon_sym_AT),
	1735: uint16(anon_sym_LBRACK),
	1736: uint16(anon_sym_DOLLAR),
	1737: uint16(385),
	1738: uint16(35),
	1739: uint16(anon_sym_inherit),
	1740: uint16(anon_sym_DMENU),
	1741: uint16(sym_integer_value),
	1742: uint16(anon_sym_true),
	1743: uint16(anon_sym_false),
	1744: uint16(anon_sym_url),
	1745: uint16(anon_sym_linear_DASHgradient),
	1746: uint16(anon_sym_0),
	1747: uint16(anon_sym_rgb),
	1748: uint16(anon_sym_rgba),
	1749: uint16(anon_sym_hsl),
	1750: uint16(anon_sym_hsla),
	1751: uint16(anon_sym_hwb),
	1752: uint16(anon_sym_hwba),
	1753: uint16(anon_sym_cmyk),
	1754: uint16(anon_sym_bold),
	1755: uint16(anon_sym_italic),
	1756: uint16(anon_sym_underline),
	1757: uint16(anon_sym_strikethrough),
	1758: uint16(anon_sym_dash),
	1759: uint16(anon_sym_solid),
	1760: uint16(anon_sym_calc),
	1761: uint16(anon_sym_center),
	1762: uint16(anon_sym_north),
	1763: uint16(anon_sym_east),
	1764: uint16(anon_sym_south),
	1765: uint16(anon_sym_west),
	1766: uint16(anon_sym_var),
	1767: uint16(anon_sym_horizontal),
	1768: uint16(anon_sym_vertical),
	1769: uint16(anon_sym_default),
	1770: uint16(anon_sym_pointer),
	1771: uint16(anon_sym_text),
	1772: uint16(anon_sym_env),
	1773: uint16(sym_identifier),
	1774: uint16(3),
	1775: uint16(3),
	1776: uint16(1),
	1777: uint16(sym_comment),
	1778: uint16(266),
	1779: uint16(9),
	1780: uint16(anon_sym_COMMA),
	1781: uint16(anon_sym_POUND),
	1782: uint16(anon_sym_SEMI),
	1783: uint16(anon_sym_RPAREN),
	1784: uint16(anon_sym_DQUOTE),
	1785: uint16(sym_float_value),
	1786: uint16(anon_sym_AT),
	1787: uint16(anon_sym_LBRACK),
	1788: uint16(anon_sym_DOLLAR),
	1789: uint16(264),
	1790: uint16(35),
	1791: uint16(anon_sym_inherit),
	1792: uint16(anon_sym_DMENU),
	1793: uint16(sym_integer_value),
	1794: uint16(anon_sym_true),
	1795: uint16(anon_sym_false),
	1796: uint16(anon_sym_url),
	1797: uint16(anon_sym_linear_DASHgradient),
	1798: uint16(anon_sym_0),
	1799: uint16(anon_sym_rgb),
	1800: uint16(anon_sym_rgba),
	1801: uint16(anon_sym_hsl),
	1802: uint16(anon_sym_hsla),
	1803: uint16(anon_sym_hwb),
	1804: uint16(anon_sym_hwba),
	1805: uint16(anon_sym_cmyk),
	1806: uint16(anon_sym_bold),
	1807: uint16(anon_sym_italic),
	1808: uint16(anon_sym_underline),
	1809: uint16(anon_sym_strikethrough),
	1810: uint16(anon_sym_dash),
	1811: uint16(anon_sym_solid),
	1812: uint16(anon_sym_calc),
	1813: uint16(anon_sym_center),
	1814: uint16(anon_sym_north),
	1815: uint16(anon_sym_east),
	1816: uint16(anon_sym_south),
	1817: uint16(anon_sym_west),
	1818: uint16(anon_sym_var),
	1819: uint16(anon_sym_horizontal),
	1820: uint16(anon_sym_vertical),
	1821: uint16(anon_sym_default),
	1822: uint16(anon_sym_pointer),
	1823: uint16(anon_sym_text),
	1824: uint16(anon_sym_env),
	1825: uint16(sym_identifier),
	1826: uint16(3),
	1827: uint16(3),
	1828: uint16(1),
	1829: uint16(sym_comment),
	1830: uint16(391),
	1831: uint16(9),
	1832: uint16(anon_sym_COMMA),
	1833: uint16(anon_sym_POUND),
	1834: uint16(anon_sym_SEMI),
	1835: uint16(anon_sym_RPAREN),
	1836: uint16(anon_sym_DQUOTE),
	1837: uint16(sym_float_value),
	1838: uint16(anon_sym_AT),
	1839: uint16(anon_sym_LBRACK),
	1840: uint16(anon_sym_DOLLAR),
	1841: uint16(389),
	1842: uint16(35),
	1843: uint16(anon_sym_inherit),
	1844: uint16(anon_sym_DMENU),
	1845: uint16(sym_integer_value),
	1846: uint16(anon_sym_true),
	1847: uint16(anon_sym_false),
	1848: uint16(anon_sym_url),
	1849: uint16(anon_sym_linear_DASHgradient),
	1850: uint16(anon_sym_0),
	1851: uint16(anon_sym_rgb),
	1852: uint16(anon_sym_rgba),
	1853: uint16(anon_sym_hsl),
	1854: uint16(anon_sym_hsla),
	1855: uint16(anon_sym_hwb),
	1856: uint16(anon_sym_hwba),
	1857: uint16(anon_sym_cmyk),
	1858: uint16(anon_sym_bold),
	1859: uint16(anon_sym_italic),
	1860: uint16(anon_sym_underline),
	1861: uint16(anon_sym_strikethrough),
	1862: uint16(anon_sym_dash),
	1863: uint16(anon_sym_solid),
	1864: uint16(anon_sym_calc),
	1865: uint16(anon_sym_center),
	1866: uint16(anon_sym_north),
	1867: uint16(anon_sym_east),
	1868: uint16(anon_sym_south),
	1869: uint16(anon_sym_west),
	1870: uint16(anon_sym_var),
	1871: uint16(anon_sym_horizontal),
	1872: uint16(anon_sym_vertical),
	1873: uint16(anon_sym_default),
	1874: uint16(anon_sym_pointer),
	1875: uint16(anon_sym_text),
	1876: uint16(anon_sym_env),
	1877: uint16(sym_identifier),
	1878: uint16(3),
	1879: uint16(3),
	1880: uint16(1),
	1881: uint16(sym_comment),
	1882: uint16(395),
	1883: uint16(9),
	1884: uint16(anon_sym_COMMA),
	1885: uint16(anon_sym_POUND),
	1886: uint16(anon_sym_SEMI),
	1887: uint16(anon_sym_RPAREN),
	1888: uint16(anon_sym_DQUOTE),
	1889: uint16(sym_float_value),
	1890: uint16(anon_sym_AT),
	1891: uint16(anon_sym_LBRACK),
	1892: uint16(anon_sym_DOLLAR),
	1893: uint16(393),
	1894: uint16(35),
	1895: uint16(anon_sym_inherit),
	1896: uint16(anon_sym_DMENU),
	1897: uint16(sym_integer_value),
	1898: uint16(anon_sym_true),
	1899: uint16(anon_sym_false),
	1900: uint16(anon_sym_url),
	1901: uint16(anon_sym_linear_DASHgradient),
	1902: uint16(anon_sym_0),
	1903: uint16(anon_sym_rgb),
	1904: uint16(anon_sym_rgba),
	1905: uint16(anon_sym_hsl),
	1906: uint16(anon_sym_hsla),
	1907: uint16(anon_sym_hwb),
	1908: uint16(anon_sym_hwba),
	1909: uint16(anon_sym_cmyk),
	1910: uint16(anon_sym_bold),
	1911: uint16(anon_sym_italic),
	1912: uint16(anon_sym_underline),
	1913: uint16(anon_sym_strikethrough),
	1914: uint16(anon_sym_dash),
	1915: uint16(anon_sym_solid),
	1916: uint16(anon_sym_calc),
	1917: uint16(anon_sym_center),
	1918: uint16(anon_sym_north),
	1919: uint16(anon_sym_east),
	1920: uint16(anon_sym_south),
	1921: uint16(anon_sym_west),
	1922: uint16(anon_sym_var),
	1923: uint16(anon_sym_horizontal),
	1924: uint16(anon_sym_vertical),
	1925: uint16(anon_sym_default),
	1926: uint16(anon_sym_pointer),
	1927: uint16(anon_sym_text),
	1928: uint16(anon_sym_env),
	1929: uint16(sym_identifier),
	1930: uint16(3),
	1931: uint16(3),
	1932: uint16(1),
	1933: uint16(sym_comment),
	1934: uint16(399),
	1935: uint16(9),
	1936: uint16(anon_sym_COMMA),
	1937: uint16(anon_sym_POUND),
	1938: uint16(anon_sym_SEMI),
	1939: uint16(anon_sym_RPAREN),
	1940: uint16(anon_sym_DQUOTE),
	1941: uint16(sym_float_value),
	1942: uint16(anon_sym_AT),
	1943: uint16(anon_sym_LBRACK),
	1944: uint16(anon_sym_DOLLAR),
	1945: uint16(397),
	1946: uint16(35),
	1947: uint16(anon_sym_inherit),
	1948: uint16(anon_sym_DMENU),
	1949: uint16(sym_integer_value),
	1950: uint16(anon_sym_true),
	1951: uint16(anon_sym_false),
	1952: uint16(anon_sym_url),
	1953: uint16(anon_sym_linear_DASHgradient),
	1954: uint16(anon_sym_0),
	1955: uint16(anon_sym_rgb),
	1956: uint16(anon_sym_rgba),
	1957: uint16(anon_sym_hsl),
	1958: uint16(anon_sym_hsla),
	1959: uint16(anon_sym_hwb),
	1960: uint16(anon_sym_hwba),
	1961: uint16(anon_sym_cmyk),
	1962: uint16(anon_sym_bold),
	1963: uint16(anon_sym_italic),
	1964: uint16(anon_sym_underline),
	1965: uint16(anon_sym_strikethrough),
	1966: uint16(anon_sym_dash),
	1967: uint16(anon_sym_solid),
	1968: uint16(anon_sym_calc),
	1969: uint16(anon_sym_center),
	1970: uint16(anon_sym_north),
	1971: uint16(anon_sym_east),
	1972: uint16(anon_sym_south),
	1973: uint16(anon_sym_west),
	1974: uint16(anon_sym_var),
	1975: uint16(anon_sym_horizontal),
	1976: uint16(anon_sym_vertical),
	1977: uint16(anon_sym_default),
	1978: uint16(anon_sym_pointer),
	1979: uint16(anon_sym_text),
	1980: uint16(anon_sym_env),
	1981: uint16(sym_identifier),
	1982: uint16(3),
	1983: uint16(3),
	1984: uint16(1),
	1985: uint16(sym_comment),
	1986: uint16(403),
	1987: uint16(9),
	1988: uint16(anon_sym_COMMA),
	1989: uint16(anon_sym_POUND),
	1990: uint16(anon_sym_SEMI),
	1991: uint16(anon_sym_RPAREN),
	1992: uint16(anon_sym_DQUOTE),
	1993: uint16(sym_float_value),
	1994: uint16(anon_sym_AT),
	1995: uint16(anon_sym_LBRACK),
	1996: uint16(anon_sym_DOLLAR),
	1997: uint16(401),
	1998: uint16(35),
	1999: uint16(anon_sym_inherit),
	2000: uint16(anon_sym_DMENU),
	2001: uint16(sym_integer_value),
	2002: uint16(anon_sym_true),
	2003: uint16(anon_sym_false),
	2004: uint16(anon_sym_url),
	2005: uint16(anon_sym_linear_DASHgradient),
	2006: uint16(anon_sym_0),
	2007: uint16(anon_sym_rgb),
	2008: uint16(anon_sym_rgba),
	2009: uint16(anon_sym_hsl),
	2010: uint16(anon_sym_hsla),
	2011: uint16(anon_sym_hwb),
	2012: uint16(anon_sym_hwba),
	2013: uint16(anon_sym_cmyk),
	2014: uint16(anon_sym_bold),
	2015: uint16(anon_sym_italic),
	2016: uint16(anon_sym_underline),
	2017: uint16(anon_sym_strikethrough),
	2018: uint16(anon_sym_dash),
	2019: uint16(anon_sym_solid),
	2020: uint16(anon_sym_calc),
	2021: uint16(anon_sym_center),
	2022: uint16(anon_sym_north),
	2023: uint16(anon_sym_east),
	2024: uint16(anon_sym_south),
	2025: uint16(anon_sym_west),
	2026: uint16(anon_sym_var),
	2027: uint16(anon_sym_horizontal),
	2028: uint16(anon_sym_vertical),
	2029: uint16(anon_sym_default),
	2030: uint16(anon_sym_pointer),
	2031: uint16(anon_sym_text),
	2032: uint16(anon_sym_env),
	2033: uint16(sym_identifier),
	2034: uint16(3),
	2035: uint16(3),
	2036: uint16(1),
	2037: uint16(sym_comment),
	2038: uint16(407),
	2039: uint16(9),
	2040: uint16(anon_sym_COMMA),
	2041: uint16(anon_sym_POUND),
	2042: uint16(anon_sym_SEMI),
	2043: uint16(anon_sym_RPAREN),
	2044: uint16(anon_sym_DQUOTE),
	2045: uint16(sym_float_value),
	2046: uint16(anon_sym_AT),
	2047: uint16(anon_sym_LBRACK),
	2048: uint16(anon_sym_DOLLAR),
	2049: uint16(405),
	2050: uint16(35),
	2051: uint16(anon_sym_inherit),
	2052: uint16(anon_sym_DMENU),
	2053: uint16(sym_integer_value),
	2054: uint16(anon_sym_true),
	2055: uint16(anon_sym_false),
	2056: uint16(anon_sym_url),
	2057: uint16(anon_sym_linear_DASHgradient),
	2058: uint16(anon_sym_0),
	2059: uint16(anon_sym_rgb),
	2060: uint16(anon_sym_rgba),
	2061: uint16(anon_sym_hsl),
	2062: uint16(anon_sym_hsla),
	2063: uint16(anon_sym_hwb),
	2064: uint16(anon_sym_hwba),
	2065: uint16(anon_sym_cmyk),
	2066: uint16(anon_sym_bold),
	2067: uint16(anon_sym_italic),
	2068: uint16(anon_sym_underline),
	2069: uint16(anon_sym_strikethrough),
	2070: uint16(anon_sym_dash),
	2071: uint16(anon_sym_solid),
	2072: uint16(anon_sym_calc),
	2073: uint16(anon_sym_center),
	2074: uint16(anon_sym_north),
	2075: uint16(anon_sym_east),
	2076: uint16(anon_sym_south),
	2077: uint16(anon_sym_west),
	2078: uint16(anon_sym_var),
	2079: uint16(anon_sym_horizontal),
	2080: uint16(anon_sym_vertical),
	2081: uint16(anon_sym_default),
	2082: uint16(anon_sym_pointer),
	2083: uint16(anon_sym_text),
	2084: uint16(anon_sym_env),
	2085: uint16(sym_identifier),
	2086: uint16(3),
	2087: uint16(3),
	2088: uint16(1),
	2089: uint16(sym_comment),
	2090: uint16(411),
	2091: uint16(9),
	2092: uint16(anon_sym_COMMA),
	2093: uint16(anon_sym_POUND),
	2094: uint16(anon_sym_SEMI),
	2095: uint16(anon_sym_RPAREN),
	2096: uint16(anon_sym_DQUOTE),
	2097: uint16(sym_float_value),
	2098: uint16(anon_sym_AT),
	2099: uint16(anon_sym_LBRACK),
	2100: uint16(anon_sym_DOLLAR),
	2101: uint16(409),
	2102: uint16(35),
	2103: uint16(anon_sym_inherit),
	2104: uint16(anon_sym_DMENU),
	2105: uint16(sym_integer_value),
	2106: uint16(anon_sym_true),
	2107: uint16(anon_sym_false),
	2108: uint16(anon_sym_url),
	2109: uint16(anon_sym_linear_DASHgradient),
	2110: uint16(anon_sym_0),
	2111: uint16(anon_sym_rgb),
	2112: uint16(anon_sym_rgba),
	2113: uint16(anon_sym_hsl),
	2114: uint16(anon_sym_hsla),
	2115: uint16(anon_sym_hwb),
	2116: uint16(anon_sym_hwba),
	2117: uint16(anon_sym_cmyk),
	2118: uint16(anon_sym_bold),
	2119: uint16(anon_sym_italic),
	2120: uint16(anon_sym_underline),
	2121: uint16(anon_sym_strikethrough),
	2122: uint16(anon_sym_dash),
	2123: uint16(anon_sym_solid),
	2124: uint16(anon_sym_calc),
	2125: uint16(anon_sym_center),
	2126: uint16(anon_sym_north),
	2127: uint16(anon_sym_east),
	2128: uint16(anon_sym_south),
	2129: uint16(anon_sym_west),
	2130: uint16(anon_sym_var),
	2131: uint16(anon_sym_horizontal),
	2132: uint16(anon_sym_vertical),
	2133: uint16(anon_sym_default),
	2134: uint16(anon_sym_pointer),
	2135: uint16(anon_sym_text),
	2136: uint16(anon_sym_env),
	2137: uint16(sym_identifier),
	2138: uint16(3),
	2139: uint16(3),
	2140: uint16(1),
	2141: uint16(sym_comment),
	2142: uint16(415),
	2143: uint16(9),
	2144: uint16(anon_sym_COMMA),
	2145: uint16(anon_sym_POUND),
	2146: uint16(anon_sym_SEMI),
	2147: uint16(anon_sym_RPAREN),
	2148: uint16(anon_sym_DQUOTE),
	2149: uint16(sym_float_value),
	2150: uint16(anon_sym_AT),
	2151: uint16(anon_sym_LBRACK),
	2152: uint16(anon_sym_DOLLAR),
	2153: uint16(413),
	2154: uint16(35),
	2155: uint16(anon_sym_inherit),
	2156: uint16(anon_sym_DMENU),
	2157: uint16(sym_integer_value),
	2158: uint16(anon_sym_true),
	2159: uint16(anon_sym_false),
	2160: uint16(anon_sym_url),
	2161: uint16(anon_sym_linear_DASHgradient),
	2162: uint16(anon_sym_0),
	2163: uint16(anon_sym_rgb),
	2164: uint16(anon_sym_rgba),
	2165: uint16(anon_sym_hsl),
	2166: uint16(anon_sym_hsla),
	2167: uint16(anon_sym_hwb),
	2168: uint16(anon_sym_hwba),
	2169: uint16(anon_sym_cmyk),
	2170: uint16(anon_sym_bold),
	2171: uint16(anon_sym_italic),
	2172: uint16(anon_sym_underline),
	2173: uint16(anon_sym_strikethrough),
	2174: uint16(anon_sym_dash),
	2175: uint16(anon_sym_solid),
	2176: uint16(anon_sym_calc),
	2177: uint16(anon_sym_center),
	2178: uint16(anon_sym_north),
	2179: uint16(anon_sym_east),
	2180: uint16(anon_sym_south),
	2181: uint16(anon_sym_west),
	2182: uint16(anon_sym_var),
	2183: uint16(anon_sym_horizontal),
	2184: uint16(anon_sym_vertical),
	2185: uint16(anon_sym_default),
	2186: uint16(anon_sym_pointer),
	2187: uint16(anon_sym_text),
	2188: uint16(anon_sym_env),
	2189: uint16(sym_identifier),
	2190: uint16(3),
	2191: uint16(3),
	2192: uint16(1),
	2193: uint16(sym_comment),
	2194: uint16(419),
	2195: uint16(9),
	2196: uint16(anon_sym_COMMA),
	2197: uint16(anon_sym_POUND),
	2198: uint16(anon_sym_SEMI),
	2199: uint16(anon_sym_RPAREN),
	2200: uint16(anon_sym_DQUOTE),
	2201: uint16(sym_float_value),
	2202: uint16(anon_sym_AT),
	2203: uint16(anon_sym_LBRACK),
	2204: uint16(anon_sym_DOLLAR),
	2205: uint16(417),
	2206: uint16(35),
	2207: uint16(anon_sym_inherit),
	2208: uint16(anon_sym_DMENU),
	2209: uint16(sym_integer_value),
	2210: uint16(anon_sym_true),
	2211: uint16(anon_sym_false),
	2212: uint16(anon_sym_url),
	2213: uint16(anon_sym_linear_DASHgradient),
	2214: uint16(anon_sym_0),
	2215: uint16(anon_sym_rgb),
	2216: uint16(anon_sym_rgba),
	2217: uint16(anon_sym_hsl),
	2218: uint16(anon_sym_hsla),
	2219: uint16(anon_sym_hwb),
	2220: uint16(anon_sym_hwba),
	2221: uint16(anon_sym_cmyk),
	2222: uint16(anon_sym_bold),
	2223: uint16(anon_sym_italic),
	2224: uint16(anon_sym_underline),
	2225: uint16(anon_sym_strikethrough),
	2226: uint16(anon_sym_dash),
	2227: uint16(anon_sym_solid),
	2228: uint16(anon_sym_calc),
	2229: uint16(anon_sym_center),
	2230: uint16(anon_sym_north),
	2231: uint16(anon_sym_east),
	2232: uint16(anon_sym_south),
	2233: uint16(anon_sym_west),
	2234: uint16(anon_sym_var),
	2235: uint16(anon_sym_horizontal),
	2236: uint16(anon_sym_vertical),
	2237: uint16(anon_sym_default),
	2238: uint16(anon_sym_pointer),
	2239: uint16(anon_sym_text),
	2240: uint16(anon_sym_env),
	2241: uint16(sym_identifier),
	2242: uint16(3),
	2243: uint16(3),
	2244: uint16(1),
	2245: uint16(sym_comment),
	2246: uint16(423),
	2247: uint16(9),
	2248: uint16(anon_sym_COMMA),
	2249: uint16(anon_sym_POUND),
	2250: uint16(anon_sym_SEMI),
	2251: uint16(anon_sym_RPAREN),
	2252: uint16(anon_sym_DQUOTE),
	2253: uint16(sym_float_value),
	2254: uint16(anon_sym_AT),
	2255: uint16(anon_sym_LBRACK),
	2256: uint16(anon_sym_DOLLAR),
	2257: uint16(421),
	2258: uint16(35),
	2259: uint16(anon_sym_inherit),
	2260: uint16(anon_sym_DMENU),
	2261: uint16(sym_integer_value),
	2262: uint16(anon_sym_true),
	2263: uint16(anon_sym_false),
	2264: uint16(anon_sym_url),
	2265: uint16(anon_sym_linear_DASHgradient),
	2266: uint16(anon_sym_0),
	2267: uint16(anon_sym_rgb),
	2268: uint16(anon_sym_rgba),
	2269: uint16(anon_sym_hsl),
	2270: uint16(anon_sym_hsla),
	2271: uint16(anon_sym_hwb),
	2272: uint16(anon_sym_hwba),
	2273: uint16(anon_sym_cmyk),
	2274: uint16(anon_sym_bold),
	2275: uint16(anon_sym_italic),
	2276: uint16(anon_sym_underline),
	2277: uint16(anon_sym_strikethrough),
	2278: uint16(anon_sym_dash),
	2279: uint16(anon_sym_solid),
	2280: uint16(anon_sym_calc),
	2281: uint16(anon_sym_center),
	2282: uint16(anon_sym_north),
	2283: uint16(anon_sym_east),
	2284: uint16(anon_sym_south),
	2285: uint16(anon_sym_west),
	2286: uint16(anon_sym_var),
	2287: uint16(anon_sym_horizontal),
	2288: uint16(anon_sym_vertical),
	2289: uint16(anon_sym_default),
	2290: uint16(anon_sym_pointer),
	2291: uint16(anon_sym_text),
	2292: uint16(anon_sym_env),
	2293: uint16(sym_identifier),
	2294: uint16(3),
	2295: uint16(3),
	2296: uint16(1),
	2297: uint16(sym_comment),
	2298: uint16(427),
	2299: uint16(9),
	2300: uint16(anon_sym_COMMA),
	2301: uint16(anon_sym_POUND),
	2302: uint16(anon_sym_SEMI),
	2303: uint16(anon_sym_RPAREN),
	2304: uint16(anon_sym_DQUOTE),
	2305: uint16(sym_float_value),
	2306: uint16(anon_sym_AT),
	2307: uint16(anon_sym_LBRACK),
	2308: uint16(anon_sym_DOLLAR),
	2309: uint16(425),
	2310: uint16(35),
	2311: uint16(anon_sym_inherit),
	2312: uint16(anon_sym_DMENU),
	2313: uint16(sym_integer_value),
	2314: uint16(anon_sym_true),
	2315: uint16(anon_sym_false),
	2316: uint16(anon_sym_url),
	2317: uint16(anon_sym_linear_DASHgradient),
	2318: uint16(anon_sym_0),
	2319: uint16(anon_sym_rgb),
	2320: uint16(anon_sym_rgba),
	2321: uint16(anon_sym_hsl),
	2322: uint16(anon_sym_hsla),
	2323: uint16(anon_sym_hwb),
	2324: uint16(anon_sym_hwba),
	2325: uint16(anon_sym_cmyk),
	2326: uint16(anon_sym_bold),
	2327: uint16(anon_sym_italic),
	2328: uint16(anon_sym_underline),
	2329: uint16(anon_sym_strikethrough),
	2330: uint16(anon_sym_dash),
	2331: uint16(anon_sym_solid),
	2332: uint16(anon_sym_calc),
	2333: uint16(anon_sym_center),
	2334: uint16(anon_sym_north),
	2335: uint16(anon_sym_east),
	2336: uint16(anon_sym_south),
	2337: uint16(anon_sym_west),
	2338: uint16(anon_sym_var),
	2339: uint16(anon_sym_horizontal),
	2340: uint16(anon_sym_vertical),
	2341: uint16(anon_sym_default),
	2342: uint16(anon_sym_pointer),
	2343: uint16(anon_sym_text),
	2344: uint16(anon_sym_env),
	2345: uint16(sym_identifier),
	2346: uint16(3),
	2347: uint16(3),
	2348: uint16(1),
	2349: uint16(sym_comment),
	2350: uint16(431),
	2351: uint16(9),
	2352: uint16(anon_sym_COMMA),
	2353: uint16(anon_sym_POUND),
	2354: uint16(anon_sym_SEMI),
	2355: uint16(anon_sym_RPAREN),
	2356: uint16(anon_sym_DQUOTE),
	2357: uint16(sym_float_value),
	2358: uint16(anon_sym_AT),
	2359: uint16(anon_sym_LBRACK),
	2360: uint16(anon_sym_DOLLAR),
	2361: uint16(429),
	2362: uint16(35),
	2363: uint16(anon_sym_inherit),
	2364: uint16(anon_sym_DMENU),
	2365: uint16(sym_integer_value),
	2366: uint16(anon_sym_true),
	2367: uint16(anon_sym_false),
	2368: uint16(anon_sym_url),
	2369: uint16(anon_sym_linear_DASHgradient),
	2370: uint16(anon_sym_0),
	2371: uint16(anon_sym_rgb),
	2372: uint16(anon_sym_rgba),
	2373: uint16(anon_sym_hsl),
	2374: uint16(anon_sym_hsla),
	2375: uint16(anon_sym_hwb),
	2376: uint16(anon_sym_hwba),
	2377: uint16(anon_sym_cmyk),
	2378: uint16(anon_sym_bold),
	2379: uint16(anon_sym_italic),
	2380: uint16(anon_sym_underline),
	2381: uint16(anon_sym_strikethrough),
	2382: uint16(anon_sym_dash),
	2383: uint16(anon_sym_solid),
	2384: uint16(anon_sym_calc),
	2385: uint16(anon_sym_center),
	2386: uint16(anon_sym_north),
	2387: uint16(anon_sym_east),
	2388: uint16(anon_sym_south),
	2389: uint16(anon_sym_west),
	2390: uint16(anon_sym_var),
	2391: uint16(anon_sym_horizontal),
	2392: uint16(anon_sym_vertical),
	2393: uint16(anon_sym_default),
	2394: uint16(anon_sym_pointer),
	2395: uint16(anon_sym_text),
	2396: uint16(anon_sym_env),
	2397: uint16(sym_identifier),
	2398: uint16(3),
	2399: uint16(3),
	2400: uint16(1),
	2401: uint16(sym_comment),
	2402: uint16(435),
	2403: uint16(9),
	2404: uint16(anon_sym_COMMA),
	2405: uint16(anon_sym_POUND),
	2406: uint16(anon_sym_SEMI),
	2407: uint16(anon_sym_RPAREN),
	2408: uint16(anon_sym_DQUOTE),
	2409: uint16(sym_float_value),
	2410: uint16(anon_sym_AT),
	2411: uint16(anon_sym_LBRACK),
	2412: uint16(anon_sym_DOLLAR),
	2413: uint16(433),
	2414: uint16(35),
	2415: uint16(anon_sym_inherit),
	2416: uint16(anon_sym_DMENU),
	2417: uint16(sym_integer_value),
	2418: uint16(anon_sym_true),
	2419: uint16(anon_sym_false),
	2420: uint16(anon_sym_url),
	2421: uint16(anon_sym_linear_DASHgradient),
	2422: uint16(anon_sym_0),
	2423: uint16(anon_sym_rgb),
	2424: uint16(anon_sym_rgba),
	2425: uint16(anon_sym_hsl),
	2426: uint16(anon_sym_hsla),
	2427: uint16(anon_sym_hwb),
	2428: uint16(anon_sym_hwba),
	2429: uint16(anon_sym_cmyk),
	2430: uint16(anon_sym_bold),
	2431: uint16(anon_sym_italic),
	2432: uint16(anon_sym_underline),
	2433: uint16(anon_sym_strikethrough),
	2434: uint16(anon_sym_dash),
	2435: uint16(anon_sym_solid),
	2436: uint16(anon_sym_calc),
	2437: uint16(anon_sym_center),
	2438: uint16(anon_sym_north),
	2439: uint16(anon_sym_east),
	2440: uint16(anon_sym_south),
	2441: uint16(anon_sym_west),
	2442: uint16(anon_sym_var),
	2443: uint16(anon_sym_horizontal),
	2444: uint16(anon_sym_vertical),
	2445: uint16(anon_sym_default),
	2446: uint16(anon_sym_pointer),
	2447: uint16(anon_sym_text),
	2448: uint16(anon_sym_env),
	2449: uint16(sym_identifier),
	2450: uint16(3),
	2451: uint16(3),
	2452: uint16(1),
	2453: uint16(sym_comment),
	2454: uint16(439),
	2455: uint16(9),
	2456: uint16(anon_sym_COMMA),
	2457: uint16(anon_sym_POUND),
	2458: uint16(anon_sym_SEMI),
	2459: uint16(anon_sym_RPAREN),
	2460: uint16(anon_sym_DQUOTE),
	2461: uint16(sym_float_value),
	2462: uint16(anon_sym_AT),
	2463: uint16(anon_sym_LBRACK),
	2464: uint16(anon_sym_DOLLAR),
	2465: uint16(437),
	2466: uint16(35),
	2467: uint16(anon_sym_inherit),
	2468: uint16(anon_sym_DMENU),
	2469: uint16(sym_integer_value),
	2470: uint16(anon_sym_true),
	2471: uint16(anon_sym_false),
	2472: uint16(anon_sym_url),
	2473: uint16(anon_sym_linear_DASHgradient),
	2474: uint16(anon_sym_0),
	2475: uint16(anon_sym_rgb),
	2476: uint16(anon_sym_rgba),
	2477: uint16(anon_sym_hsl),
	2478: uint16(anon_sym_hsla),
	2479: uint16(anon_sym_hwb),
	2480: uint16(anon_sym_hwba),
	2481: uint16(anon_sym_cmyk),
	2482: uint16(anon_sym_bold),
	2483: uint16(anon_sym_italic),
	2484: uint16(anon_sym_underline),
	2485: uint16(anon_sym_strikethrough),
	2486: uint16(anon_sym_dash),
	2487: uint16(anon_sym_solid),
	2488: uint16(anon_sym_calc),
	2489: uint16(anon_sym_center),
	2490: uint16(anon_sym_north),
	2491: uint16(anon_sym_east),
	2492: uint16(anon_sym_south),
	2493: uint16(anon_sym_west),
	2494: uint16(anon_sym_var),
	2495: uint16(anon_sym_horizontal),
	2496: uint16(anon_sym_vertical),
	2497: uint16(anon_sym_default),
	2498: uint16(anon_sym_pointer),
	2499: uint16(anon_sym_text),
	2500: uint16(anon_sym_env),
	2501: uint16(sym_identifier),
	2502: uint16(3),
	2503: uint16(3),
	2504: uint16(1),
	2505: uint16(sym_comment),
	2506: uint16(443),
	2507: uint16(9),
	2508: uint16(anon_sym_COMMA),
	2509: uint16(anon_sym_POUND),
	2510: uint16(anon_sym_SEMI),
	2511: uint16(anon_sym_RPAREN),
	2512: uint16(anon_sym_DQUOTE),
	2513: uint16(sym_float_value),
	2514: uint16(anon_sym_AT),
	2515: uint16(anon_sym_LBRACK),
	2516: uint16(anon_sym_DOLLAR),
	2517: uint16(441),
	2518: uint16(35),
	2519: uint16(anon_sym_inherit),
	2520: uint16(anon_sym_DMENU),
	2521: uint16(sym_integer_value),
	2522: uint16(anon_sym_true),
	2523: uint16(anon_sym_false),
	2524: uint16(anon_sym_url),
	2525: uint16(anon_sym_linear_DASHgradient),
	2526: uint16(anon_sym_0),
	2527: uint16(anon_sym_rgb),
	2528: uint16(anon_sym_rgba),
	2529: uint16(anon_sym_hsl),
	2530: uint16(anon_sym_hsla),
	2531: uint16(anon_sym_hwb),
	2532: uint16(anon_sym_hwba),
	2533: uint16(anon_sym_cmyk),
	2534: uint16(anon_sym_bold),
	2535: uint16(anon_sym_italic),
	2536: uint16(anon_sym_underline),
	2537: uint16(anon_sym_strikethrough),
	2538: uint16(anon_sym_dash),
	2539: uint16(anon_sym_solid),
	2540: uint16(anon_sym_calc),
	2541: uint16(anon_sym_center),
	2542: uint16(anon_sym_north),
	2543: uint16(anon_sym_east),
	2544: uint16(anon_sym_south),
	2545: uint16(anon_sym_west),
	2546: uint16(anon_sym_var),
	2547: uint16(anon_sym_horizontal),
	2548: uint16(anon_sym_vertical),
	2549: uint16(anon_sym_default),
	2550: uint16(anon_sym_pointer),
	2551: uint16(anon_sym_text),
	2552: uint16(anon_sym_env),
	2553: uint16(sym_identifier),
	2554: uint16(3),
	2555: uint16(3),
	2556: uint16(1),
	2557: uint16(sym_comment),
	2558: uint16(447),
	2559: uint16(9),
	2560: uint16(anon_sym_COMMA),
	2561: uint16(anon_sym_POUND),
	2562: uint16(anon_sym_SEMI),
	2563: uint16(anon_sym_RPAREN),
	2564: uint16(anon_sym_DQUOTE),
	2565: uint16(sym_float_value),
	2566: uint16(anon_sym_AT),
	2567: uint16(anon_sym_LBRACK),
	2568: uint16(anon_sym_DOLLAR),
	2569: uint16(445),
	2570: uint16(35),
	2571: uint16(anon_sym_inherit),
	2572: uint16(anon_sym_DMENU),
	2573: uint16(sym_integer_value),
	2574: uint16(anon_sym_true),
	2575: uint16(anon_sym_false),
	2576: uint16(anon_sym_url),
	2577: uint16(anon_sym_linear_DASHgradient),
	2578: uint16(anon_sym_0),
	2579: uint16(anon_sym_rgb),
	2580: uint16(anon_sym_rgba),
	2581: uint16(anon_sym_hsl),
	2582: uint16(anon_sym_hsla),
	2583: uint16(anon_sym_hwb),
	2584: uint16(anon_sym_hwba),
	2585: uint16(anon_sym_cmyk),
	2586: uint16(anon_sym_bold),
	2587: uint16(anon_sym_italic),
	2588: uint16(anon_sym_underline),
	2589: uint16(anon_sym_strikethrough),
	2590: uint16(anon_sym_dash),
	2591: uint16(anon_sym_solid),
	2592: uint16(anon_sym_calc),
	2593: uint16(anon_sym_center),
	2594: uint16(anon_sym_north),
	2595: uint16(anon_sym_east),
	2596: uint16(anon_sym_south),
	2597: uint16(anon_sym_west),
	2598: uint16(anon_sym_var),
	2599: uint16(anon_sym_horizontal),
	2600: uint16(anon_sym_vertical),
	2601: uint16(anon_sym_default),
	2602: uint16(anon_sym_pointer),
	2603: uint16(anon_sym_text),
	2604: uint16(anon_sym_env),
	2605: uint16(sym_identifier),
	2606: uint16(3),
	2607: uint16(3),
	2608: uint16(1),
	2609: uint16(sym_comment),
	2610: uint16(451),
	2611: uint16(9),
	2612: uint16(anon_sym_COMMA),
	2613: uint16(anon_sym_POUND),
	2614: uint16(anon_sym_SEMI),
	2615: uint16(anon_sym_RPAREN),
	2616: uint16(anon_sym_DQUOTE),
	2617: uint16(sym_float_value),
	2618: uint16(anon_sym_AT),
	2619: uint16(anon_sym_LBRACK),
	2620: uint16(anon_sym_DOLLAR),
	2621: uint16(449),
	2622: uint16(35),
	2623: uint16(anon_sym_inherit),
	2624: uint16(anon_sym_DMENU),
	2625: uint16(sym_integer_value),
	2626: uint16(anon_sym_true),
	2627: uint16(anon_sym_false),
	2628: uint16(anon_sym_url),
	2629: uint16(anon_sym_linear_DASHgradient),
	2630: uint16(anon_sym_0),
	2631: uint16(anon_sym_rgb),
	2632: uint16(anon_sym_rgba),
	2633: uint16(anon_sym_hsl),
	2634: uint16(anon_sym_hsla),
	2635: uint16(anon_sym_hwb),
	2636: uint16(anon_sym_hwba),
	2637: uint16(anon_sym_cmyk),
	2638: uint16(anon_sym_bold),
	2639: uint16(anon_sym_italic),
	2640: uint16(anon_sym_underline),
	2641: uint16(anon_sym_strikethrough),
	2642: uint16(anon_sym_dash),
	2643: uint16(anon_sym_solid),
	2644: uint16(anon_sym_calc),
	2645: uint16(anon_sym_center),
	2646: uint16(anon_sym_north),
	2647: uint16(anon_sym_east),
	2648: uint16(anon_sym_south),
	2649: uint16(anon_sym_west),
	2650: uint16(anon_sym_var),
	2651: uint16(anon_sym_horizontal),
	2652: uint16(anon_sym_vertical),
	2653: uint16(anon_sym_default),
	2654: uint16(anon_sym_pointer),
	2655: uint16(anon_sym_text),
	2656: uint16(anon_sym_env),
	2657: uint16(sym_identifier),
	2658: uint16(3),
	2659: uint16(3),
	2660: uint16(1),
	2661: uint16(sym_comment),
	2662: uint16(455),
	2663: uint16(9),
	2664: uint16(anon_sym_COMMA),
	2665: uint16(anon_sym_POUND),
	2666: uint16(anon_sym_SEMI),
	2667: uint16(anon_sym_RPAREN),
	2668: uint16(anon_sym_DQUOTE),
	2669: uint16(sym_float_value),
	2670: uint16(anon_sym_AT),
	2671: uint16(anon_sym_LBRACK),
	2672: uint16(anon_sym_DOLLAR),
	2673: uint16(453),
	2674: uint16(35),
	2675: uint16(anon_sym_inherit),
	2676: uint16(anon_sym_DMENU),
	2677: uint16(sym_integer_value),
	2678: uint16(anon_sym_true),
	2679: uint16(anon_sym_false),
	2680: uint16(anon_sym_url),
	2681: uint16(anon_sym_linear_DASHgradient),
	2682: uint16(anon_sym_0),
	2683: uint16(anon_sym_rgb),
	2684: uint16(anon_sym_rgba),
	2685: uint16(anon_sym_hsl),
	2686: uint16(anon_sym_hsla),
	2687: uint16(anon_sym_hwb),
	2688: uint16(anon_sym_hwba),
	2689: uint16(anon_sym_cmyk),
	2690: uint16(anon_sym_bold),
	2691: uint16(anon_sym_italic),
	2692: uint16(anon_sym_underline),
	2693: uint16(anon_sym_strikethrough),
	2694: uint16(anon_sym_dash),
	2695: uint16(anon_sym_solid),
	2696: uint16(anon_sym_calc),
	2697: uint16(anon_sym_center),
	2698: uint16(anon_sym_north),
	2699: uint16(anon_sym_east),
	2700: uint16(anon_sym_south),
	2701: uint16(anon_sym_west),
	2702: uint16(anon_sym_var),
	2703: uint16(anon_sym_horizontal),
	2704: uint16(anon_sym_vertical),
	2705: uint16(anon_sym_default),
	2706: uint16(anon_sym_pointer),
	2707: uint16(anon_sym_text),
	2708: uint16(anon_sym_env),
	2709: uint16(sym_identifier),
	2710: uint16(3),
	2711: uint16(3),
	2712: uint16(1),
	2713: uint16(sym_comment),
	2714: uint16(459),
	2715: uint16(9),
	2716: uint16(anon_sym_COMMA),
	2717: uint16(anon_sym_POUND),
	2718: uint16(anon_sym_SEMI),
	2719: uint16(anon_sym_RPAREN),
	2720: uint16(anon_sym_DQUOTE),
	2721: uint16(sym_float_value),
	2722: uint16(anon_sym_AT),
	2723: uint16(anon_sym_LBRACK),
	2724: uint16(anon_sym_DOLLAR),
	2725: uint16(457),
	2726: uint16(35),
	2727: uint16(anon_sym_inherit),
	2728: uint16(anon_sym_DMENU),
	2729: uint16(sym_integer_value),
	2730: uint16(anon_sym_true),
	2731: uint16(anon_sym_false),
	2732: uint16(anon_sym_url),
	2733: uint16(anon_sym_linear_DASHgradient),
	2734: uint16(anon_sym_0),
	2735: uint16(anon_sym_rgb),
	2736: uint16(anon_sym_rgba),
	2737: uint16(anon_sym_hsl),
	2738: uint16(anon_sym_hsla),
	2739: uint16(anon_sym_hwb),
	2740: uint16(anon_sym_hwba),
	2741: uint16(anon_sym_cmyk),
	2742: uint16(anon_sym_bold),
	2743: uint16(anon_sym_italic),
	2744: uint16(anon_sym_underline),
	2745: uint16(anon_sym_strikethrough),
	2746: uint16(anon_sym_dash),
	2747: uint16(anon_sym_solid),
	2748: uint16(anon_sym_calc),
	2749: uint16(anon_sym_center),
	2750: uint16(anon_sym_north),
	2751: uint16(anon_sym_east),
	2752: uint16(anon_sym_south),
	2753: uint16(anon_sym_west),
	2754: uint16(anon_sym_var),
	2755: uint16(anon_sym_horizontal),
	2756: uint16(anon_sym_vertical),
	2757: uint16(anon_sym_default),
	2758: uint16(anon_sym_pointer),
	2759: uint16(anon_sym_text),
	2760: uint16(anon_sym_env),
	2761: uint16(sym_identifier),
	2762: uint16(3),
	2763: uint16(3),
	2764: uint16(1),
	2765: uint16(sym_comment),
	2766: uint16(463),
	2767: uint16(9),
	2768: uint16(anon_sym_COMMA),
	2769: uint16(anon_sym_POUND),
	2770: uint16(anon_sym_SEMI),
	2771: uint16(anon_sym_RPAREN),
	2772: uint16(anon_sym_DQUOTE),
	2773: uint16(sym_float_value),
	2774: uint16(anon_sym_AT),
	2775: uint16(anon_sym_LBRACK),
	2776: uint16(anon_sym_DOLLAR),
	2777: uint16(461),
	2778: uint16(35),
	2779: uint16(anon_sym_inherit),
	2780: uint16(anon_sym_DMENU),
	2781: uint16(sym_integer_value),
	2782: uint16(anon_sym_true),
	2783: uint16(anon_sym_false),
	2784: uint16(anon_sym_url),
	2785: uint16(anon_sym_linear_DASHgradient),
	2786: uint16(anon_sym_0),
	2787: uint16(anon_sym_rgb),
	2788: uint16(anon_sym_rgba),
	2789: uint16(anon_sym_hsl),
	2790: uint16(anon_sym_hsla),
	2791: uint16(anon_sym_hwb),
	2792: uint16(anon_sym_hwba),
	2793: uint16(anon_sym_cmyk),
	2794: uint16(anon_sym_bold),
	2795: uint16(anon_sym_italic),
	2796: uint16(anon_sym_underline),
	2797: uint16(anon_sym_strikethrough),
	2798: uint16(anon_sym_dash),
	2799: uint16(anon_sym_solid),
	2800: uint16(anon_sym_calc),
	2801: uint16(anon_sym_center),
	2802: uint16(anon_sym_north),
	2803: uint16(anon_sym_east),
	2804: uint16(anon_sym_south),
	2805: uint16(anon_sym_west),
	2806: uint16(anon_sym_var),
	2807: uint16(anon_sym_horizontal),
	2808: uint16(anon_sym_vertical),
	2809: uint16(anon_sym_default),
	2810: uint16(anon_sym_pointer),
	2811: uint16(anon_sym_text),
	2812: uint16(anon_sym_env),
	2813: uint16(sym_identifier),
	2814: uint16(3),
	2815: uint16(3),
	2816: uint16(1),
	2817: uint16(sym_comment),
	2818: uint16(467),
	2819: uint16(9),
	2820: uint16(anon_sym_COMMA),
	2821: uint16(anon_sym_POUND),
	2822: uint16(anon_sym_SEMI),
	2823: uint16(anon_sym_RPAREN),
	2824: uint16(anon_sym_DQUOTE),
	2825: uint16(sym_float_value),
	2826: uint16(anon_sym_AT),
	2827: uint16(anon_sym_LBRACK),
	2828: uint16(anon_sym_DOLLAR),
	2829: uint16(465),
	2830: uint16(35),
	2831: uint16(anon_sym_inherit),
	2832: uint16(anon_sym_DMENU),
	2833: uint16(sym_integer_value),
	2834: uint16(anon_sym_true),
	2835: uint16(anon_sym_false),
	2836: uint16(anon_sym_url),
	2837: uint16(anon_sym_linear_DASHgradient),
	2838: uint16(anon_sym_0),
	2839: uint16(anon_sym_rgb),
	2840: uint16(anon_sym_rgba),
	2841: uint16(anon_sym_hsl),
	2842: uint16(anon_sym_hsla),
	2843: uint16(anon_sym_hwb),
	2844: uint16(anon_sym_hwba),
	2845: uint16(anon_sym_cmyk),
	2846: uint16(anon_sym_bold),
	2847: uint16(anon_sym_italic),
	2848: uint16(anon_sym_underline),
	2849: uint16(anon_sym_strikethrough),
	2850: uint16(anon_sym_dash),
	2851: uint16(anon_sym_solid),
	2852: uint16(anon_sym_calc),
	2853: uint16(anon_sym_center),
	2854: uint16(anon_sym_north),
	2855: uint16(anon_sym_east),
	2856: uint16(anon_sym_south),
	2857: uint16(anon_sym_west),
	2858: uint16(anon_sym_var),
	2859: uint16(anon_sym_horizontal),
	2860: uint16(anon_sym_vertical),
	2861: uint16(anon_sym_default),
	2862: uint16(anon_sym_pointer),
	2863: uint16(anon_sym_text),
	2864: uint16(anon_sym_env),
	2865: uint16(sym_identifier),
	2866: uint16(3),
	2867: uint16(3),
	2868: uint16(1),
	2869: uint16(sym_comment),
	2870: uint16(471),
	2871: uint16(9),
	2872: uint16(anon_sym_COMMA),
	2873: uint16(anon_sym_POUND),
	2874: uint16(anon_sym_SEMI),
	2875: uint16(anon_sym_RPAREN),
	2876: uint16(anon_sym_DQUOTE),
	2877: uint16(sym_float_value),
	2878: uint16(anon_sym_AT),
	2879: uint16(anon_sym_LBRACK),
	2880: uint16(anon_sym_DOLLAR),
	2881: uint16(469),
	2882: uint16(35),
	2883: uint16(anon_sym_inherit),
	2884: uint16(anon_sym_DMENU),
	2885: uint16(sym_integer_value),
	2886: uint16(anon_sym_true),
	2887: uint16(anon_sym_false),
	2888: uint16(anon_sym_url),
	2889: uint16(anon_sym_linear_DASHgradient),
	2890: uint16(anon_sym_0),
	2891: uint16(anon_sym_rgb),
	2892: uint16(anon_sym_rgba),
	2893: uint16(anon_sym_hsl),
	2894: uint16(anon_sym_hsla),
	2895: uint16(anon_sym_hwb),
	2896: uint16(anon_sym_hwba),
	2897: uint16(anon_sym_cmyk),
	2898: uint16(anon_sym_bold),
	2899: uint16(anon_sym_italic),
	2900: uint16(anon_sym_underline),
	2901: uint16(anon_sym_strikethrough),
	2902: uint16(anon_sym_dash),
	2903: uint16(anon_sym_solid),
	2904: uint16(anon_sym_calc),
	2905: uint16(anon_sym_center),
	2906: uint16(anon_sym_north),
	2907: uint16(anon_sym_east),
	2908: uint16(anon_sym_south),
	2909: uint16(anon_sym_west),
	2910: uint16(anon_sym_var),
	2911: uint16(anon_sym_horizontal),
	2912: uint16(anon_sym_vertical),
	2913: uint16(anon_sym_default),
	2914: uint16(anon_sym_pointer),
	2915: uint16(anon_sym_text),
	2916: uint16(anon_sym_env),
	2917: uint16(sym_identifier),
	2918: uint16(3),
	2919: uint16(3),
	2920: uint16(1),
	2921: uint16(sym_comment),
	2922: uint16(475),
	2923: uint16(9),
	2924: uint16(anon_sym_COMMA),
	2925: uint16(anon_sym_POUND),
	2926: uint16(anon_sym_SEMI),
	2927: uint16(anon_sym_RPAREN),
	2928: uint16(anon_sym_DQUOTE),
	2929: uint16(sym_float_value),
	2930: uint16(anon_sym_AT),
	2931: uint16(anon_sym_LBRACK),
	2932: uint16(anon_sym_DOLLAR),
	2933: uint16(473),
	2934: uint16(35),
	2935: uint16(anon_sym_inherit),
	2936: uint16(anon_sym_DMENU),
	2937: uint16(sym_integer_value),
	2938: uint16(anon_sym_true),
	2939: uint16(anon_sym_false),
	2940: uint16(anon_sym_url),
	2941: uint16(anon_sym_linear_DASHgradient),
	2942: uint16(anon_sym_0),
	2943: uint16(anon_sym_rgb),
	2944: uint16(anon_sym_rgba),
	2945: uint16(anon_sym_hsl),
	2946: uint16(anon_sym_hsla),
	2947: uint16(anon_sym_hwb),
	2948: uint16(anon_sym_hwba),
	2949: uint16(anon_sym_cmyk),
	2950: uint16(anon_sym_bold),
	2951: uint16(anon_sym_italic),
	2952: uint16(anon_sym_underline),
	2953: uint16(anon_sym_strikethrough),
	2954: uint16(anon_sym_dash),
	2955: uint16(anon_sym_solid),
	2956: uint16(anon_sym_calc),
	2957: uint16(anon_sym_center),
	2958: uint16(anon_sym_north),
	2959: uint16(anon_sym_east),
	2960: uint16(anon_sym_south),
	2961: uint16(anon_sym_west),
	2962: uint16(anon_sym_var),
	2963: uint16(anon_sym_horizontal),
	2964: uint16(anon_sym_vertical),
	2965: uint16(anon_sym_default),
	2966: uint16(anon_sym_pointer),
	2967: uint16(anon_sym_text),
	2968: uint16(anon_sym_env),
	2969: uint16(sym_identifier),
	2970: uint16(3),
	2971: uint16(3),
	2972: uint16(1),
	2973: uint16(sym_comment),
	2974: uint16(479),
	2975: uint16(9),
	2976: uint16(anon_sym_COMMA),
	2977: uint16(anon_sym_POUND),
	2978: uint16(anon_sym_SEMI),
	2979: uint16(anon_sym_RPAREN),
	2980: uint16(anon_sym_DQUOTE),
	2981: uint16(sym_float_value),
	2982: uint16(anon_sym_AT),
	2983: uint16(anon_sym_LBRACK),
	2984: uint16(anon_sym_DOLLAR),
	2985: uint16(477),
	2986: uint16(35),
	2987: uint16(anon_sym_inherit),
	2988: uint16(anon_sym_DMENU),
	2989: uint16(sym_integer_value),
	2990: uint16(anon_sym_true),
	2991: uint16(anon_sym_false),
	2992: uint16(anon_sym_url),
	2993: uint16(anon_sym_linear_DASHgradient),
	2994: uint16(anon_sym_0),
	2995: uint16(anon_sym_rgb),
	2996: uint16(anon_sym_rgba),
	2997: uint16(anon_sym_hsl),
	2998: uint16(anon_sym_hsla),
	2999: uint16(anon_sym_hwb),
	3000: uint16(anon_sym_hwba),
	3001: uint16(anon_sym_cmyk),
	3002: uint16(anon_sym_bold),
	3003: uint16(anon_sym_italic),
	3004: uint16(anon_sym_underline),
	3005: uint16(anon_sym_strikethrough),
	3006: uint16(anon_sym_dash),
	3007: uint16(anon_sym_solid),
	3008: uint16(anon_sym_calc),
	3009: uint16(anon_sym_center),
	3010: uint16(anon_sym_north),
	3011: uint16(anon_sym_east),
	3012: uint16(anon_sym_south),
	3013: uint16(anon_sym_west),
	3014: uint16(anon_sym_var),
	3015: uint16(anon_sym_horizontal),
	3016: uint16(anon_sym_vertical),
	3017: uint16(anon_sym_default),
	3018: uint16(anon_sym_pointer),
	3019: uint16(anon_sym_text),
	3020: uint16(anon_sym_env),
	3021: uint16(sym_identifier),
	3022: uint16(3),
	3023: uint16(3),
	3024: uint16(1),
	3025: uint16(sym_comment),
	3026: uint16(483),
	3027: uint16(9),
	3028: uint16(anon_sym_COMMA),
	3029: uint16(anon_sym_POUND),
	3030: uint16(anon_sym_SEMI),
	3031: uint16(anon_sym_RPAREN),
	3032: uint16(anon_sym_DQUOTE),
	3033: uint16(sym_float_value),
	3034: uint16(anon_sym_AT),
	3035: uint16(anon_sym_LBRACK),
	3036: uint16(anon_sym_DOLLAR),
	3037: uint16(481),
	3038: uint16(35),
	3039: uint16(anon_sym_inherit),
	3040: uint16(anon_sym_DMENU),
	3041: uint16(sym_integer_value),
	3042: uint16(anon_sym_true),
	3043: uint16(anon_sym_false),
	3044: uint16(anon_sym_url),
	3045: uint16(anon_sym_linear_DASHgradient),
	3046: uint16(anon_sym_0),
	3047: uint16(anon_sym_rgb),
	3048: uint16(anon_sym_rgba),
	3049: uint16(anon_sym_hsl),
	3050: uint16(anon_sym_hsla),
	3051: uint16(anon_sym_hwb),
	3052: uint16(anon_sym_hwba),
	3053: uint16(anon_sym_cmyk),
	3054: uint16(anon_sym_bold),
	3055: uint16(anon_sym_italic),
	3056: uint16(anon_sym_underline),
	3057: uint16(anon_sym_strikethrough),
	3058: uint16(anon_sym_dash),
	3059: uint16(anon_sym_solid),
	3060: uint16(anon_sym_calc),
	3061: uint16(anon_sym_center),
	3062: uint16(anon_sym_north),
	3063: uint16(anon_sym_east),
	3064: uint16(anon_sym_south),
	3065: uint16(anon_sym_west),
	3066: uint16(anon_sym_var),
	3067: uint16(anon_sym_horizontal),
	3068: uint16(anon_sym_vertical),
	3069: uint16(anon_sym_default),
	3070: uint16(anon_sym_pointer),
	3071: uint16(anon_sym_text),
	3072: uint16(anon_sym_env),
	3073: uint16(sym_identifier),
	3074: uint16(3),
	3075: uint16(3),
	3076: uint16(1),
	3077: uint16(sym_comment),
	3078: uint16(487),
	3079: uint16(9),
	3080: uint16(anon_sym_COMMA),
	3081: uint16(anon_sym_POUND),
	3082: uint16(anon_sym_SEMI),
	3083: uint16(anon_sym_RPAREN),
	3084: uint16(anon_sym_DQUOTE),
	3085: uint16(sym_float_value),
	3086: uint16(anon_sym_AT),
	3087: uint16(anon_sym_LBRACK),
	3088: uint16(anon_sym_DOLLAR),
	3089: uint16(485),
	3090: uint16(35),
	3091: uint16(anon_sym_inherit),
	3092: uint16(anon_sym_DMENU),
	3093: uint16(sym_integer_value),
	3094: uint16(anon_sym_true),
	3095: uint16(anon_sym_false),
	3096: uint16(anon_sym_url),
	3097: uint16(anon_sym_linear_DASHgradient),
	3098: uint16(anon_sym_0),
	3099: uint16(anon_sym_rgb),
	3100: uint16(anon_sym_rgba),
	3101: uint16(anon_sym_hsl),
	3102: uint16(anon_sym_hsla),
	3103: uint16(anon_sym_hwb),
	3104: uint16(anon_sym_hwba),
	3105: uint16(anon_sym_cmyk),
	3106: uint16(anon_sym_bold),
	3107: uint16(anon_sym_italic),
	3108: uint16(anon_sym_underline),
	3109: uint16(anon_sym_strikethrough),
	3110: uint16(anon_sym_dash),
	3111: uint16(anon_sym_solid),
	3112: uint16(anon_sym_calc),
	3113: uint16(anon_sym_center),
	3114: uint16(anon_sym_north),
	3115: uint16(anon_sym_east),
	3116: uint16(anon_sym_south),
	3117: uint16(anon_sym_west),
	3118: uint16(anon_sym_var),
	3119: uint16(anon_sym_horizontal),
	3120: uint16(anon_sym_vertical),
	3121: uint16(anon_sym_default),
	3122: uint16(anon_sym_pointer),
	3123: uint16(anon_sym_text),
	3124: uint16(anon_sym_env),
	3125: uint16(sym_identifier),
	3126: uint16(3),
	3127: uint16(3),
	3128: uint16(1),
	3129: uint16(sym_comment),
	3130: uint16(491),
	3131: uint16(9),
	3132: uint16(anon_sym_COMMA),
	3133: uint16(anon_sym_POUND),
	3134: uint16(anon_sym_SEMI),
	3135: uint16(anon_sym_RPAREN),
	3136: uint16(anon_sym_DQUOTE),
	3137: uint16(sym_float_value),
	3138: uint16(anon_sym_AT),
	3139: uint16(anon_sym_LBRACK),
	3140: uint16(anon_sym_DOLLAR),
	3141: uint16(489),
	3142: uint16(35),
	3143: uint16(anon_sym_inherit),
	3144: uint16(anon_sym_DMENU),
	3145: uint16(sym_integer_value),
	3146: uint16(anon_sym_true),
	3147: uint16(anon_sym_false),
	3148: uint16(anon_sym_url),
	3149: uint16(anon_sym_linear_DASHgradient),
	3150: uint16(anon_sym_0),
	3151: uint16(anon_sym_rgb),
	3152: uint16(anon_sym_rgba),
	3153: uint16(anon_sym_hsl),
	3154: uint16(anon_sym_hsla),
	3155: uint16(anon_sym_hwb),
	3156: uint16(anon_sym_hwba),
	3157: uint16(anon_sym_cmyk),
	3158: uint16(anon_sym_bold),
	3159: uint16(anon_sym_italic),
	3160: uint16(anon_sym_underline),
	3161: uint16(anon_sym_strikethrough),
	3162: uint16(anon_sym_dash),
	3163: uint16(anon_sym_solid),
	3164: uint16(anon_sym_calc),
	3165: uint16(anon_sym_center),
	3166: uint16(anon_sym_north),
	3167: uint16(anon_sym_east),
	3168: uint16(anon_sym_south),
	3169: uint16(anon_sym_west),
	3170: uint16(anon_sym_var),
	3171: uint16(anon_sym_horizontal),
	3172: uint16(anon_sym_vertical),
	3173: uint16(anon_sym_default),
	3174: uint16(anon_sym_pointer),
	3175: uint16(anon_sym_text),
	3176: uint16(anon_sym_env),
	3177: uint16(sym_identifier),
	3178: uint16(3),
	3179: uint16(3),
	3180: uint16(1),
	3181: uint16(sym_comment),
	3182: uint16(495),
	3183: uint16(8),
	3184: uint16(anon_sym_COMMA),
	3185: uint16(anon_sym_POUND),
	3186: uint16(anon_sym_SEMI),
	3187: uint16(anon_sym_DQUOTE),
	3188: uint16(sym_float_value),
	3189: uint16(anon_sym_AT),
	3190: uint16(anon_sym_LBRACK),
	3191: uint16(anon_sym_DOLLAR),
	3192: uint16(493),
	3193: uint16(35),
	3194: uint16(anon_sym_inherit),
	3195: uint16(anon_sym_DMENU),
	3196: uint16(sym_integer_value),
	3197: uint16(anon_sym_true),
	3198: uint16(anon_sym_false),
	3199: uint16(anon_sym_url),
	3200: uint16(anon_sym_linear_DASHgradient),
	3201: uint16(anon_sym_0),
	3202: uint16(anon_sym_rgb),
	3203: uint16(anon_sym_rgba),
	3204: uint16(anon_sym_hsl),
	3205: uint16(anon_sym_hsla),
	3206: uint16(anon_sym_hwb),
	3207: uint16(anon_sym_hwba),
	3208: uint16(anon_sym_cmyk),
	3209: uint16(anon_sym_bold),
	3210: uint16(anon_sym_italic),
	3211: uint16(anon_sym_underline),
	3212: uint16(anon_sym_strikethrough),
	3213: uint16(anon_sym_dash),
	3214: uint16(anon_sym_solid),
	3215: uint16(anon_sym_calc),
	3216: uint16(anon_sym_center),
	3217: uint16(anon_sym_north),
	3218: uint16(anon_sym_east),
	3219: uint16(anon_sym_south),
	3220: uint16(anon_sym_west),
	3221: uint16(anon_sym_var),
	3222: uint16(anon_sym_horizontal),
	3223: uint16(anon_sym_vertical),
	3224: uint16(anon_sym_default),
	3225: uint16(anon_sym_pointer),
	3226: uint16(anon_sym_text),
	3227: uint16(anon_sym_env),
	3228: uint16(sym_identifier),
	3229: uint16(3),
	3230: uint16(3),
	3231: uint16(1),
	3232: uint16(sym_comment),
	3233: uint16(499),
	3234: uint16(8),
	3235: uint16(anon_sym_COMMA),
	3236: uint16(anon_sym_POUND),
	3237: uint16(anon_sym_SEMI),
	3238: uint16(anon_sym_DQUOTE),
	3239: uint16(sym_float_value),
	3240: uint16(anon_sym_AT),
	3241: uint16(anon_sym_LBRACK),
	3242: uint16(anon_sym_DOLLAR),
	3243: uint16(497),
	3244: uint16(35),
	3245: uint16(anon_sym_inherit),
	3246: uint16(anon_sym_DMENU),
	3247: uint16(sym_integer_value),
	3248: uint16(anon_sym_true),
	3249: uint16(anon_sym_false),
	3250: uint16(anon_sym_url),
	3251: uint16(anon_sym_linear_DASHgradient),
	3252: uint16(anon_sym_0),
	3253: uint16(anon_sym_rgb),
	3254: uint16(anon_sym_rgba),
	3255: uint16(anon_sym_hsl),
	3256: uint16(anon_sym_hsla),
	3257: uint16(anon_sym_hwb),
	3258: uint16(anon_sym_hwba),
	3259: uint16(anon_sym_cmyk),
	3260: uint16(anon_sym_bold),
	3261: uint16(anon_sym_italic),
	3262: uint16(anon_sym_underline),
	3263: uint16(anon_sym_strikethrough),
	3264: uint16(anon_sym_dash),
	3265: uint16(anon_sym_solid),
	3266: uint16(anon_sym_calc),
	3267: uint16(anon_sym_center),
	3268: uint16(anon_sym_north),
	3269: uint16(anon_sym_east),
	3270: uint16(anon_sym_south),
	3271: uint16(anon_sym_west),
	3272: uint16(anon_sym_var),
	3273: uint16(anon_sym_horizontal),
	3274: uint16(anon_sym_vertical),
	3275: uint16(anon_sym_default),
	3276: uint16(anon_sym_pointer),
	3277: uint16(anon_sym_text),
	3278: uint16(anon_sym_env),
	3279: uint16(sym_identifier),
	3280: uint16(11),
	3281: uint16(3),
	3282: uint16(1),
	3283: uint16(sym_comment),
	3284: uint16(19),
	3285: uint16(1),
	3286: uint16(sym_identifier),
	3287: uint16(23),
	3288: uint16(1),
	3289: uint16(anon_sym_POUND),
	3290: uint16(49),
	3291: uint16(1),
	3292: uint16(anon_sym_cmyk),
	3293: uint16(501),
	3294: uint16(1),
	3295: uint16(sym_integer_value),
	3296: uint16(503),
	3297: uint16(1),
	3298: uint16(anon_sym_to),
	3299: uint16(220),
	3300: uint16(1),
	3301: uint16(sym_angle),
	3302: uint16(43),
	3303: uint16(2),
	3304: uint16(anon_sym_rgb),
	3305: uint16(anon_sym_rgba),
	3306: uint16(45),
	3307: uint16(2),
	3308: uint16(anon_sym_hsl),
	3309: uint16(anon_sym_hsla),
	3310: uint16(47),
	3311: uint16(2),
	3312: uint16(anon_sym_hwb),
	3313: uint16(anon_sym_hwba),
	3314: uint16(237),
	3315: uint16(7),
	3316: uint16(sym__color_value),
	3317: uint16(sym_hex_color),
	3318: uint16(sym_rgb_color),
	3319: uint16(sym_hsl_color),
	3320: uint16(sym_hwb_color),
	3321: uint16(sym_cmyk_color),
	3322: uint16(sym_named_color),
	3323: uint16(6),
	3324: uint16(3),
	3325: uint16(1),
	3326: uint16(sym_comment),
	3327: uint16(264),
	3328: uint16(1),
	3329: uint16(anon_sym_SLASH),
	3330: uint16(507),
	3331: uint16(2),
	3332: uint16(anon_sym_px),
	3333: uint16(anon_sym_mm),
	3334: uint16(95),
	3335: uint16(2),
	3336: uint16(sym_integer_distance_unit),
	3337: uint16(sym_float_distance_unit),
	3338: uint16(505),
	3339: uint16(4),
	3340: uint16(anon_sym_PERCENT),
	3341: uint16(anon_sym_cm),
	3342: uint16(anon_sym_ph),
	3343: uint16(anon_sym_em),
	3344: uint16(266),
	3345: uint16(10),
	3346: uint16(anon_sym_STAR),
	3347: uint16(anon_sym_RPAREN),
	3348: uint16(anon_sym_PLUS),
	3349: uint16(anon_sym_DASH),
	3350: uint16(anon_sym_modulo),
	3351: uint16(anon_sym_min),
	3352: uint16(anon_sym_max),
	3353: uint16(anon_sym_floor),
	3354: uint16(anon_sym_ceil),
	3355: uint16(anon_sym_round),
	3356: uint16(8),
	3357: uint16(3),
	3358: uint16(1),
	3359: uint16(sym_comment),
	3360: uint16(19),
	3361: uint16(1),
	3362: uint16(sym_identifier),
	3363: uint16(23),
	3364: uint16(1),
	3365: uint16(anon_sym_POUND),
	3366: uint16(49),
	3367: uint16(1),
	3368: uint16(anon_sym_cmyk),
	3369: uint16(43),
	3370: uint16(2),
	3371: uint16(anon_sym_rgb),
	3372: uint16(anon_sym_rgba),
	3373: uint16(45),
	3374: uint16(2),
	3375: uint16(anon_sym_hsl),
	3376: uint16(anon_sym_hsla),
	3377: uint16(47),
	3378: uint16(2),
	3379: uint16(anon_sym_hwb),
	3380: uint16(anon_sym_hwba),
	3381: uint16(160),
	3382: uint16(7),
	3383: uint16(sym__color_value),
	3384: uint16(sym_hex_color),
	3385: uint16(sym_rgb_color),
	3386: uint16(sym_hsl_color),
	3387: uint16(sym_hwb_color),
	3388: uint16(sym_cmyk_color),
	3389: uint16(sym_named_color),
	3390: uint16(8),
	3391: uint16(3),
	3392: uint16(1),
	3393: uint16(sym_comment),
	3394: uint16(19),
	3395: uint16(1),
	3396: uint16(sym_identifier),
	3397: uint16(23),
	3398: uint16(1),
	3399: uint16(anon_sym_POUND),
	3400: uint16(49),
	3401: uint16(1),
	3402: uint16(anon_sym_cmyk),
	3403: uint16(43),
	3404: uint16(2),
	3405: uint16(anon_sym_rgb),
	3406: uint16(anon_sym_rgba),
	3407: uint16(45),
	3408: uint16(2),
	3409: uint16(anon_sym_hsl),
	3410: uint16(anon_sym_hsla),
	3411: uint16(47),
	3412: uint16(2),
	3413: uint16(anon_sym_hwb),
	3414: uint16(anon_sym_hwba),
	3415: uint16(156),
	3416: uint16(7),
	3417: uint16(sym__color_value),
	3418: uint16(sym_hex_color),
	3419: uint16(sym_rgb_color),
	3420: uint16(sym_hsl_color),
	3421: uint16(sym_hwb_color),
	3422: uint16(sym_cmyk_color),
	3423: uint16(sym_named_color),
	3424: uint16(11),
	3425: uint16(3),
	3426: uint16(1),
	3427: uint16(sym_comment),
	3428: uint16(7),
	3429: uint16(1),
	3430: uint16(sym_identifier),
	3431: uint16(9),
	3432: uint16(1),
	3433: uint16(anon_sym_ATimport),
	3434: uint16(11),
	3435: uint16(1),
	3436: uint16(anon_sym_ATtheme),
	3437: uint16(13),
	3438: uint16(1),
	3439: uint16(anon_sym_ATmedia),
	3440: uint16(15),
	3441: uint16(1),
	3442: uint16(anon_sym_STAR),
	3443: uint16(17),
	3444: uint16(1),
	3445: uint16(anon_sym_POUND),
	3446: uint16(509),
	3447: uint16(1),
	3449: uint16(186),
	3450: uint16(1),
	3451: uint16(sym_selectors),
	3452: uint16(154),
	3453: uint16(3),
	3454: uint16(sym__selector),
	3455: uint16(sym_global_selector),
	3456: uint16(sym_id_selector),
	3457: uint16(81),
	3458: uint16(5),
	3459: uint16(sym_import_statement),
	3460: uint16(sym_theme_statement),
	3461: uint16(sym_media_statement),
	3462: uint16(sym_rule_set),
	3463: uint16(aux_sym_stylesheet_repeat1),
	3464: uint16(8),
	3465: uint16(3),
	3466: uint16(1),
	3467: uint16(sym_comment),
	3468: uint16(19),
	3469: uint16(1),
	3470: uint16(sym_identifier),
	3471: uint16(23),
	3472: uint16(1),
	3473: uint16(anon_sym_POUND),
	3474: uint16(49),
	3475: uint16(1),
	3476: uint16(anon_sym_cmyk),
	3477: uint16(43),
	3478: uint16(2),
	3479: uint16(anon_sym_rgb),
	3480: uint16(anon_sym_rgba),
	3481: uint16(45),
	3482: uint16(2),
	3483: uint16(anon_sym_hsl),
	3484: uint16(anon_sym_hsla),
	3485: uint16(47),
	3486: uint16(2),
	3487: uint16(anon_sym_hwb),
	3488: uint16(anon_sym_hwba),
	3489: uint16(221),
	3490: uint16(7),
	3491: uint16(sym__color_value),
	3492: uint16(sym_hex_color),
	3493: uint16(sym_rgb_color),
	3494: uint16(sym_hsl_color),
	3495: uint16(sym_hwb_color),
	3496: uint16(sym_cmyk_color),
	3497: uint16(sym_named_color),
	3498: uint16(8),
	3499: uint16(3),
	3500: uint16(1),
	3501: uint16(sym_comment),
	3502: uint16(19),
	3503: uint16(1),
	3504: uint16(sym_identifier),
	3505: uint16(23),
	3506: uint16(1),
	3507: uint16(anon_sym_POUND),
	3508: uint16(49),
	3509: uint16(1),
	3510: uint16(anon_sym_cmyk),
	3511: uint16(43),
	3512: uint16(2),
	3513: uint16(anon_sym_rgb),
	3514: uint16(anon_sym_rgba),
	3515: uint16(45),
	3516: uint16(2),
	3517: uint16(anon_sym_hsl),
	3518: uint16(anon_sym_hsla),
	3519: uint16(47),
	3520: uint16(2),
	3521: uint16(anon_sym_hwb),
	3522: uint16(anon_sym_hwba),
	3523: uint16(204),
	3524: uint16(7),
	3525: uint16(sym__color_value),
	3526: uint16(sym_hex_color),
	3527: uint16(sym_rgb_color),
	3528: uint16(sym_hsl_color),
	3529: uint16(sym_hwb_color),
	3530: uint16(sym_cmyk_color),
	3531: uint16(sym_named_color),
	3532: uint16(8),
	3533: uint16(3),
	3534: uint16(1),
	3535: uint16(sym_comment),
	3536: uint16(19),
	3537: uint16(1),
	3538: uint16(sym_identifier),
	3539: uint16(23),
	3540: uint16(1),
	3541: uint16(anon_sym_POUND),
	3542: uint16(49),
	3543: uint16(1),
	3544: uint16(anon_sym_cmyk),
	3545: uint16(43),
	3546: uint16(2),
	3547: uint16(anon_sym_rgb),
	3548: uint16(anon_sym_rgba),
	3549: uint16(45),
	3550: uint16(2),
	3551: uint16(anon_sym_hsl),
	3552: uint16(anon_sym_hsla),
	3553: uint16(47),
	3554: uint16(2),
	3555: uint16(anon_sym_hwb),
	3556: uint16(anon_sym_hwba),
	3557: uint16(176),
	3558: uint16(7),
	3559: uint16(sym__color_value),
	3560: uint16(sym_hex_color),
	3561: uint16(sym_rgb_color),
	3562: uint16(sym_hsl_color),
	3563: uint16(sym_hwb_color),
	3564: uint16(sym_cmyk_color),
	3565: uint16(sym_named_color),
	3566: uint16(8),
	3567: uint16(3),
	3568: uint16(1),
	3569: uint16(sym_comment),
	3570: uint16(19),
	3571: uint16(1),
	3572: uint16(sym_identifier),
	3573: uint16(23),
	3574: uint16(1),
	3575: uint16(anon_sym_POUND),
	3576: uint16(49),
	3577: uint16(1),
	3578: uint16(anon_sym_cmyk),
	3579: uint16(43),
	3580: uint16(2),
	3581: uint16(anon_sym_rgb),
	3582: uint16(anon_sym_rgba),
	3583: uint16(45),
	3584: uint16(2),
	3585: uint16(anon_sym_hsl),
	3586: uint16(anon_sym_hsla),
	3587: uint16(47),
	3588: uint16(2),
	3589: uint16(anon_sym_hwb),
	3590: uint16(anon_sym_hwba),
	3591: uint16(158),
	3592: uint16(7),
	3593: uint16(sym__color_value),
	3594: uint16(sym_hex_color),
	3595: uint16(sym_rgb_color),
	3596: uint16(sym_hsl_color),
	3597: uint16(sym_hwb_color),
	3598: uint16(sym_cmyk_color),
	3599: uint16(sym_named_color),
	3600: uint16(11),
	3601: uint16(3),
	3602: uint16(1),
	3603: uint16(sym_comment),
	3604: uint16(511),
	3605: uint16(1),
	3607: uint16(513),
	3608: uint16(1),
	3609: uint16(sym_identifier),
	3610: uint16(516),
	3611: uint16(1),
	3612: uint16(anon_sym_ATimport),
	3613: uint16(519),
	3614: uint16(1),
	3615: uint16(anon_sym_ATtheme),
	3616: uint16(522),
	3617: uint16(1),
	3618: uint16(anon_sym_ATmedia),
	3619: uint16(525),
	3620: uint16(1),
	3621: uint16(anon_sym_STAR),
	3622: uint16(528),
	3623: uint16(1),
	3624: uint16(anon_sym_POUND),
	3625: uint16(186),
	3626: uint16(1),
	3627: uint16(sym_selectors),
	3628: uint16(154),
	3629: uint16(3),
	3630: uint16(sym__selector),
	3631: uint16(sym_global_selector),
	3632: uint16(sym_id_selector),
	3633: uint16(81),
	3634: uint16(5),
	3635: uint16(sym_import_statement),
	3636: uint16(sym_theme_statement),
	3637: uint16(sym_media_statement),
	3638: uint16(sym_rule_set),
	3639: uint16(aux_sym_stylesheet_repeat1),
	3640: uint16(6),
	3641: uint16(3),
	3642: uint16(1),
	3643: uint16(sym_comment),
	3644: uint16(264),
	3645: uint16(2),
	3646: uint16(sym_integer_value),
	3647: uint16(anon_sym_0),
	3648: uint16(531),
	3649: uint16(2),
	3650: uint16(anon_sym_px),
	3651: uint16(anon_sym_mm),
	3652: uint16(27),
	3653: uint16(2),
	3654: uint16(sym_integer_distance_unit),
	3655: uint16(sym_float_distance_unit),
	3656: uint16(258),
	3657: uint16(4),
	3658: uint16(anon_sym_PERCENT),
	3659: uint16(anon_sym_cm),
	3660: uint16(anon_sym_ph),
	3661: uint16(anon_sym_em),
	3662: uint16(266),
	3663: uint16(5),
	3664: uint16(anon_sym_RPAREN),
	3665: uint16(sym_float_value),
	3666: uint16(anon_sym_dash),
	3667: uint16(anon_sym_solid),
	3668: uint16(anon_sym_calc),
	3669: uint16(6),
	3670: uint16(3),
	3671: uint16(1),
	3672: uint16(sym_comment),
	3673: uint16(535),
	3674: uint16(1),
	3675: uint16(anon_sym_RPAREN),
	3676: uint16(537),
	3677: uint16(1),
	3678: uint16(anon_sym_SLASH),
	3679: uint16(86),
	3680: uint16(1),
	3681: uint16(aux_sym_distance_calc_repeat1),
	3682: uint16(122),
	3683: uint16(1),
	3684: uint16(sym_distance_op),
	3685: uint16(533),
	3686: uint16(9),
	3687: uint16(anon_sym_STAR),
	3688: uint16(anon_sym_PLUS),
	3689: uint16(anon_sym_DASH),
	3690: uint16(anon_sym_modulo),
	3691: uint16(anon_sym_min),
	3692: uint16(anon_sym_max),
	3693: uint16(anon_sym_floor),
	3694: uint16(anon_sym_ceil),
	3695: uint16(anon_sym_round),
	3696: uint16(5),
	3697: uint16(3),
	3698: uint16(1),
	3699: uint16(sym_comment),
	3700: uint16(539),
	3701: uint16(1),
	3702: uint16(anon_sym_LPAREN),
	3703: uint16(246),
	3704: uint16(1),
	3705: uint16(sym_feature_name),
	3706: uint16(245),
	3707: uint16(3),
	3708: uint16(sym__query),
	3709: uint16(sym_feature_query),
	3710: uint16(sym_parenthesized_query),
	3711: uint16(541),
	3712: uint16(8),
	3713: uint16(anon_sym_min_DASHwidth),
	3714: uint16(anon_sym_max_DASHwidth),
	3715: uint16(anon_sym_min_DASHheight),
	3716: uint16(anon_sym_max_DASHheight),
	3717: uint16(anon_sym_min_DASHaspect_DASHratio),
	3718: uint16(anon_sym_max_DASHaspect_DASHratio),
	3719: uint16(anon_sym_monitor_DASHid),
	3720: uint16(anon_sym_enabled),
	3721: uint16(6),
	3722: uint16(3),
	3723: uint16(1),
	3724: uint16(sym_comment),
	3725: uint16(546),
	3726: uint16(1),
	3727: uint16(anon_sym_RPAREN),
	3728: uint16(548),
	3729: uint16(1),
	3730: uint16(anon_sym_SLASH),
	3731: uint16(85),
	3732: uint16(1),
	3733: uint16(aux_sym_distance_calc_repeat1),
	3734: uint16(122),
	3735: uint16(1),
	3736: uint16(sym_distance_op),
	3737: uint16(543),
	3738: uint16(9),
	3739: uint16(anon_sym_STAR),
	3740: uint16(anon_sym_PLUS),
	3741: uint16(anon_sym_DASH),
	3742: uint16(anon_sym_modulo),
	3743: uint16(anon_sym_min),
	3744: uint16(anon_sym_max),
	3745: uint16(anon_sym_floor),
	3746: uint16(anon_sym_ceil),
	3747: uint16(anon_sym_round),
	3748: uint16(6),
	3749: uint16(3),
	3750: uint16(1),
	3751: uint16(sym_comment),
	3752: uint16(537),
	3753: uint16(1),
	3754: uint16(anon_sym_SLASH),
	3755: uint16(551),
	3756: uint16(1),
	3757: uint16(anon_sym_RPAREN),
	3758: uint16(85),
	3759: uint16(1),
	3760: uint16(aux_sym_distance_calc_repeat1),
	3761: uint16(122),
	3762: uint16(1),
	3763: uint16(sym_distance_op),
	3764: uint16(533),
	3765: uint16(9),
	3766: uint16(anon_sym_STAR),
	3767: uint16(anon_sym_PLUS),
	3768: uint16(anon_sym_DASH),
	3769: uint16(anon_sym_modulo),
	3770: uint16(anon_sym_min),
	3771: uint16(anon_sym_max),
	3772: uint16(anon_sym_floor),
	3773: uint16(anon_sym_ceil),
	3774: uint16(anon_sym_round),
	3775: uint16(6),
	3776: uint16(3),
	3777: uint16(1),
	3778: uint16(sym_comment),
	3779: uint16(537),
	3780: uint16(1),
	3781: uint16(anon_sym_SLASH),
	3782: uint16(553),
	3783: uint16(1),
	3784: uint16(anon_sym_RPAREN),
	3785: uint16(88),
	3786: uint16(1),
	3787: uint16(aux_sym_distance_calc_repeat1),
	3788: uint16(122),
	3789: uint16(1),
	3790: uint16(sym_distance_op),
	3791: uint16(533),
	3792: uint16(9),
	3793: uint16(anon_sym_STAR),
	3794: uint16(anon_sym_PLUS),
	3795: uint16(anon_sym_DASH),
	3796: uint16(anon_sym_modulo),
	3797: uint16(anon_sym_min),
	3798: uint16(anon_sym_max),
	3799: uint16(anon_sym_floor),
	3800: uint16(anon_sym_ceil),
	3801: uint16(anon_sym_round),
	3802: uint16(6),
	3803: uint16(3),
	3804: uint16(1),
	3805: uint16(sym_comment),
	3806: uint16(537),
	3807: uint16(1),
	3808: uint16(anon_sym_SLASH),
	3809: uint16(555),
	3810: uint16(1),
	3811: uint16(anon_sym_RPAREN),
	3812: uint16(85),
	3813: uint16(1),
	3814: uint16(aux_sym_distance_calc_repeat1),
	3815: uint16(122),
	3816: uint16(1),
	3817: uint16(sym_distance_op),
	3818: uint16(533),
	3819: uint16(9),
	3820: uint16(anon_sym_STAR),
	3821: uint16(anon_sym_PLUS),
	3822: uint16(anon_sym_DASH),
	3823: uint16(anon_sym_modulo),
	3824: uint16(anon_sym_min),
	3825: uint16(anon_sym_max),
	3826: uint16(anon_sym_floor),
	3827: uint16(anon_sym_ceil),
	3828: uint16(anon_sym_round),
	3829: uint16(8),
	3830: uint16(3),
	3831: uint16(1),
	3832: uint16(sym_comment),
	3833: uint16(15),
	3834: uint16(1),
	3835: uint16(anon_sym_STAR),
	3836: uint16(17),
	3837: uint16(1),
	3838: uint16(anon_sym_POUND),
	3839: uint16(557),
	3840: uint16(1),
	3841: uint16(sym_identifier),
	3842: uint16(559),
	3843: uint16(1),
	3844: uint16(anon_sym_RBRACE),
	3845: uint16(186),
	3846: uint16(1),
	3847: uint16(sym_selectors),
	3848: uint16(154),
	3849: uint16(3),
	3850: uint16(sym__selector),
	3851: uint16(sym_global_selector),
	3852: uint16(sym_id_selector),
	3853: uint16(90),
	3854: uint16(4),
	3855: uint16(sym_rule_set),
	3856: uint16(sym__block_item),
	3857: uint16(sym_declaration),
	3858: uint16(aux_sym_block_repeat1),
	3859: uint16(8),
	3860: uint16(3),
	3861: uint16(1),
	3862: uint16(sym_comment),
	3863: uint16(15),
	3864: uint16(1),
	3865: uint16(anon_sym_STAR),
	3866: uint16(17),
	3867: uint16(1),
	3868: uint16(anon_sym_POUND),
	3869: uint16(557),
	3870: uint16(1),
	3871: uint16(sym_identifier),
	3872: uint16(561),
	3873: uint16(1),
	3874: uint16(anon_sym_RBRACE),
	3875: uint16(186),
	3876: uint16(1),
	3877: uint16(sym_selectors),
	3878: uint16(154),
	3879: uint16(3),
	3880: uint16(sym__selector),
	3881: uint16(sym_global_selector),
	3882: uint16(sym_id_selector),
	3883: uint16(91),
	3884: uint16(4),
	3885: uint16(sym_rule_set),
	3886: uint16(sym__block_item),
	3887: uint16(sym_declaration),
	3888: uint16(aux_sym_block_repeat1),
	3889: uint16(8),
	3890: uint16(3),
	3891: uint16(1),
	3892: uint16(sym_comment),
	3893: uint16(563),
	3894: uint16(1),
	3895: uint16(sym_identifier),
	3896: uint16(566),
	3897: uint16(1),
	3898: uint16(anon_sym_RBRACE),
	3899: uint16(568),
	3900: uint16(1),
	3901: uint16(anon_sym_STAR),
	3902: uint16(571),
	3903: uint16(1),
	3904: uint16(anon_sym_POUND),
	3905: uint16(186),
	3906: uint16(1),
	3907: uint16(sym_selectors),
	3908: uint16(154),
	3909: uint16(3),
	3910: uint16(sym__selector),
	3911: uint16(sym_global_selector),
	3912: uint16(sym_id_selector),
	3913: uint16(91),
	3914: uint16(4),
	3915: uint16(sym_rule_set),
	3916: uint16(sym__block_item),
	3917: uint16(sym_declaration),
	3918: uint16(aux_sym_block_repeat1),
	3919: uint16(3),
	3920: uint16(3),
	3921: uint16(1),
	3922: uint16(sym_comment),
	3923: uint16(417),
	3924: uint16(1),
	3925: uint16(anon_sym_SLASH),
	3926: uint16(419),
	3927: uint16(10),
	3928: uint16(anon_sym_STAR),
	3929: uint16(anon_sym_RPAREN),
	3930: uint16(anon_sym_PLUS),
	3931: uint16(anon_sym_DASH),
	3932: uint16(anon_sym_modulo),
	3933: uint16(anon_sym_min),
	3934: uint16(anon_sym_max),
	3935: uint16(anon_sym_floor),
	3936: uint16(anon_sym_ceil),
	3937: uint16(anon_sym_round),
	3938: uint16(3),
	3939: uint16(3),
	3940: uint16(1),
	3941: uint16(sym_comment),
	3942: uint16(365),
	3943: uint16(1),
	3944: uint16(anon_sym_SLASH),
	3945: uint16(367),
	3946: uint16(10),
	3947: uint16(anon_sym_STAR),
	3948: uint16(anon_sym_RPAREN),
	3949: uint16(anon_sym_PLUS),
	3950: uint16(anon_sym_DASH),
	3951: uint16(anon_sym_modulo),
	3952: uint16(anon_sym_min),
	3953: uint16(anon_sym_max),
	3954: uint16(anon_sym_floor),
	3955: uint16(anon_sym_ceil),
	3956: uint16(anon_sym_round),
	3957: uint16(3),
	3958: uint16(3),
	3959: uint16(1),
	3960: uint16(sym_comment),
	3961: uint16(576),
	3962: uint16(1),
	3963: uint16(anon_sym_SLASH),
	3964: uint16(574),
	3965: uint16(10),
	3966: uint16(anon_sym_STAR),
	3967: uint16(anon_sym_RPAREN),
	3968: uint16(anon_sym_PLUS),
	3969: uint16(anon_sym_DASH),
	3970: uint16(anon_sym_modulo),
	3971: uint16(anon_sym_min),
	3972: uint16(anon_sym_max),
	3973: uint16(anon_sym_floor),
	3974: uint16(anon_sym_ceil),
	3975: uint16(anon_sym_round),
	3976: uint16(3),
	3977: uint16(3),
	3978: uint16(1),
	3979: uint16(sym_comment),
	3980: uint16(325),
	3981: uint16(1),
	3982: uint16(anon_sym_SLASH),
	3983: uint16(327),
	3984: uint16(10),
	3985: uint16(anon_sym_STAR),
	3986: uint16(anon_sym_RPAREN),
	3987: uint16(anon_sym_PLUS),
	3988: uint16(anon_sym_DASH),
	3989: uint16(anon_sym_modulo),
	3990: uint16(anon_sym_min),
	3991: uint16(anon_sym_max),
	3992: uint16(anon_sym_floor),
	3993: uint16(anon_sym_ceil),
	3994: uint16(anon_sym_round),
	3995: uint16(3),
	3996: uint16(3),
	3997: uint16(1),
	3998: uint16(sym_comment),
	3999: uint16(357),
	4000: uint16(1),
	4001: uint16(anon_sym_SLASH),
	4002: uint16(359),
	4003: uint16(10),
	4004: uint16(anon_sym_STAR),
	4005: uint16(anon_sym_RPAREN),
	4006: uint16(anon_sym_PLUS),
	4007: uint16(anon_sym_DASH),
	4008: uint16(anon_sym_modulo),
	4009: uint16(anon_sym_min),
	4010: uint16(anon_sym_max),
	4011: uint16(anon_sym_floor),
	4012: uint16(anon_sym_ceil),
	4013: uint16(anon_sym_round),
	4014: uint16(3),
	4015: uint16(3),
	4016: uint16(1),
	4017: uint16(sym_comment),
	4018: uint16(264),
	4019: uint16(1),
	4020: uint16(anon_sym_SLASH),
	4021: uint16(266),
	4022: uint16(10),
	4023: uint16(anon_sym_STAR),
	4024: uint16(anon_sym_RPAREN),
	4025: uint16(anon_sym_PLUS),
	4026: uint16(anon_sym_DASH),
	4027: uint16(anon_sym_modulo),
	4028: uint16(anon_sym_min),
	4029: uint16(anon_sym_max),
	4030: uint16(anon_sym_floor),
	4031: uint16(anon_sym_ceil),
	4032: uint16(anon_sym_round),
	4033: uint16(3),
	4034: uint16(3),
	4035: uint16(1),
	4036: uint16(sym_comment),
	4037: uint16(369),
	4038: uint16(1),
	4039: uint16(anon_sym_SLASH),
	4040: uint16(371),
	4041: uint16(10),
	4042: uint16(anon_sym_STAR),
	4043: uint16(anon_sym_RPAREN),
	4044: uint16(anon_sym_PLUS),
	4045: uint16(anon_sym_DASH),
	4046: uint16(anon_sym_modulo),
	4047: uint16(anon_sym_min),
	4048: uint16(anon_sym_max),
	4049: uint16(anon_sym_floor),
	4050: uint16(anon_sym_ceil),
	4051: uint16(anon_sym_round),
	4052: uint16(3),
	4053: uint16(3),
	4054: uint16(1),
	4055: uint16(sym_comment),
	4056: uint16(393),
	4057: uint16(1),
	4058: uint16(anon_sym_SLASH),
	4059: uint16(395),
	4060: uint16(10),
	4061: uint16(anon_sym_STAR),
	4062: uint16(anon_sym_RPAREN),
	4063: uint16(anon_sym_PLUS),
	4064: uint16(anon_sym_DASH),
	4065: uint16(anon_sym_modulo),
	4066: uint16(anon_sym_min),
	4067: uint16(anon_sym_max),
	4068: uint16(anon_sym_floor),
	4069: uint16(anon_sym_ceil),
	4070: uint16(anon_sym_round),
	4071: uint16(10),
	4072: uint16(3),
	4073: uint16(1),
	4074: uint16(sym_comment),
	4075: uint16(246),
	4076: uint16(1),
	4077: uint16(anon_sym_0),
	4078: uint16(256),
	4079: uint16(1),
	4080: uint16(anon_sym_RPAREN),
	4081: uint16(268),
	4082: uint16(1),
	4083: uint16(sym_integer_value),
	4084: uint16(270),
	4085: uint16(1),
	4086: uint16(sym_float_value),
	4087: uint16(580),
	4088: uint16(1),
	4089: uint16(anon_sym_calc),
	4090: uint16(30),
	4091: uint16(1),
	4092: uint16(sym_line_style_value),
	4093: uint16(43),
	4094: uint16(1),
	4095: uint16(sym_distance_calc),
	4096: uint16(121),
	4097: uint16(1),
	4098: uint16(sym_distance_value),
	4099: uint16(578),
	4100: uint16(2),
	4101: uint16(anon_sym_dash),
	4102: uint16(anon_sym_solid),
	4103: uint16(2),
	4104: uint16(3),
	4105: uint16(1),
	4106: uint16(sym_comment),
	4107: uint16(487),
	4108: uint16(9),
	4110: uint16(anon_sym_ATimport),
	4111: uint16(anon_sym_ATtheme),
	4112: uint16(anon_sym_ATmedia),
	4113: uint16(anon_sym_COMMA),
	4114: uint16(anon_sym_STAR),
	4115: uint16(anon_sym_POUND),
	4116: uint16(anon_sym_RPAREN),
	4117: uint16(sym_identifier),
	4118: uint16(5),
	4119: uint16(3),
	4120: uint16(1),
	4121: uint16(sym_comment),
	4122: uint16(256),
	4123: uint16(1),
	4124: uint16(anon_sym_RPAREN),
	4125: uint16(531),
	4126: uint16(2),
	4127: uint16(anon_sym_px),
	4128: uint16(anon_sym_mm),
	4129: uint16(38),
	4130: uint16(2),
	4131: uint16(sym_integer_distance_unit),
	4132: uint16(sym_float_distance_unit),
	4133: uint16(258),
	4134: uint16(4),
	4135: uint16(anon_sym_PERCENT),
	4136: uint16(anon_sym_cm),
	4137: uint16(anon_sym_ph),
	4138: uint16(anon_sym_em),
	4139: uint16(9),
	4140: uint16(3),
	4141: uint16(1),
	4142: uint16(sym_comment),
	4143: uint16(246),
	4144: uint16(1),
	4145: uint16(anon_sym_0),
	4146: uint16(268),
	4147: uint16(1),
	4148: uint16(sym_integer_value),
	4149: uint16(270),
	4150: uint16(1),
	4151: uint16(sym_float_value),
	4152: uint16(274),
	4153: uint16(1),
	4154: uint16(anon_sym_RPAREN),
	4155: uint16(580),
	4156: uint16(1),
	4157: uint16(anon_sym_calc),
	4158: uint16(22),
	4159: uint16(1),
	4160: uint16(sym_distance_value),
	4161: uint16(43),
	4162: uint16(1),
	4163: uint16(sym_distance_calc),
	4164: uint16(109),
	4165: uint16(1),
	4166: uint16(sym_border_style),
	4167: uint16(2),
	4168: uint16(3),
	4169: uint16(1),
	4170: uint16(sym_comment),
	4171: uint16(582),
	4172: uint16(8),
	4174: uint16(anon_sym_ATimport),
	4175: uint16(anon_sym_ATtheme),
	4176: uint16(anon_sym_ATmedia),
	4177: uint16(anon_sym_RBRACE),
	4178: uint16(anon_sym_STAR),
	4179: uint16(anon_sym_POUND),
	4180: uint16(sym_identifier),
	4181: uint16(2),
	4182: uint16(3),
	4183: uint16(1),
	4184: uint16(sym_comment),
	4185: uint16(584),
	4186: uint16(8),
	4188: uint16(anon_sym_ATimport),
	4189: uint16(anon_sym_ATtheme),
	4190: uint16(anon_sym_ATmedia),
	4191: uint16(anon_sym_RBRACE),
	4192: uint16(anon_sym_STAR),
	4193: uint16(anon_sym_POUND),
	4194: uint16(sym_identifier),
	4195: uint16(4),
	4196: uint16(3),
	4197: uint16(1),
	4198: uint16(sym_comment),
	4199: uint16(531),
	4200: uint16(2),
	4201: uint16(anon_sym_px),
	4202: uint16(anon_sym_mm),
	4203: uint16(38),
	4204: uint16(2),
	4205: uint16(sym_integer_distance_unit),
	4206: uint16(sym_float_distance_unit),
	4207: uint16(258),
	4208: uint16(4),
	4209: uint16(anon_sym_PERCENT),
	4210: uint16(anon_sym_cm),
	4211: uint16(anon_sym_ph),
	4212: uint16(anon_sym_em),
	4213: uint16(2),
	4214: uint16(3),
	4215: uint16(1),
	4216: uint16(sym_comment),
	4217: uint16(586),
	4218: uint16(8),
	4220: uint16(anon_sym_ATimport),
	4221: uint16(anon_sym_ATtheme),
	4222: uint16(anon_sym_ATmedia),
	4223: uint16(anon_sym_RBRACE),
	4224: uint16(anon_sym_STAR),
	4225: uint16(anon_sym_POUND),
	4226: uint16(sym_identifier),
	4227: uint16(4),
	4228: uint16(3),
	4229: uint16(1),
	4230: uint16(sym_comment),
	4231: uint16(507),
	4232: uint16(2),
	4233: uint16(anon_sym_px),
	4234: uint16(anon_sym_mm),
	4235: uint16(98),
	4236: uint16(2),
	4237: uint16(sym_integer_distance_unit),
	4238: uint16(sym_float_distance_unit),
	4239: uint16(505),
	4240: uint16(4),
	4241: uint16(anon_sym_PERCENT),
	4242: uint16(anon_sym_cm),
	4243: uint16(anon_sym_ph),
	4244: uint16(anon_sym_em),
	4245: uint16(9),
	4246: uint16(3),
	4247: uint16(1),
	4248: uint16(sym_comment),
	4249: uint16(246),
	4250: uint16(1),
	4251: uint16(anon_sym_0),
	4252: uint16(268),
	4253: uint16(1),
	4254: uint16(sym_integer_value),
	4255: uint16(270),
	4256: uint16(1),
	4257: uint16(sym_float_value),
	4258: uint16(278),
	4259: uint16(1),
	4260: uint16(anon_sym_RPAREN),
	4261: uint16(580),
	4262: uint16(1),
	4263: uint16(anon_sym_calc),
	4264: uint16(22),
	4265: uint16(1),
	4266: uint16(sym_distance_value),
	4267: uint16(43),
	4268: uint16(1),
	4269: uint16(sym_distance_calc),
	4270: uint16(110),
	4271: uint16(1),
	4272: uint16(sym_border_style),
	4273: uint16(9),
	4274: uint16(3),
	4275: uint16(1),
	4276: uint16(sym_comment),
	4277: uint16(246),
	4278: uint16(1),
	4279: uint16(anon_sym_0),
	4280: uint16(268),
	4281: uint16(1),
	4282: uint16(sym_integer_value),
	4283: uint16(270),
	4284: uint16(1),
	4285: uint16(sym_float_value),
	4286: uint16(282),
	4287: uint16(1),
	4288: uint16(anon_sym_RPAREN),
	4289: uint16(580),
	4290: uint16(1),
	4291: uint16(anon_sym_calc),
	4292: uint16(22),
	4293: uint16(1),
	4294: uint16(sym_distance_value),
	4295: uint16(43),
	4296: uint16(1),
	4297: uint16(sym_distance_calc),
	4298: uint16(50),
	4299: uint16(1),
	4300: uint16(sym_border_style),
	4301: uint16(6),
	4302: uint16(3),
	4303: uint16(1),
	4304: uint16(sym_comment),
	4305: uint16(590),
	4306: uint16(1),
	4307: uint16(anon_sym_DOT),
	4308: uint16(594),
	4309: uint16(1),
	4310: uint16(anon_sym_COLON),
	4311: uint16(183),
	4312: uint16(1),
	4313: uint16(sym_id_selector_view),
	4314: uint16(588),
	4315: uint16(2),
	4316: uint16(anon_sym_COMMA),
	4317: uint16(anon_sym_LBRACE),
	4318: uint16(592),
	4319: uint16(3),
	4320: uint16(anon_sym_normal),
	4321: uint16(anon_sym_selected),
	4322: uint16(anon_sym_alternate),
	4323: uint16(2),
	4324: uint16(3),
	4325: uint16(1),
	4326: uint16(sym_comment),
	4327: uint16(596),
	4328: uint16(7),
	4330: uint16(anon_sym_ATimport),
	4331: uint16(anon_sym_ATtheme),
	4332: uint16(anon_sym_ATmedia),
	4333: uint16(anon_sym_STAR),
	4334: uint16(anon_sym_POUND),
	4335: uint16(sym_identifier),
	4336: uint16(5),
	4337: uint16(3),
	4338: uint16(1),
	4339: uint16(sym_comment),
	4340: uint16(590),
	4341: uint16(1),
	4342: uint16(anon_sym_DOT),
	4343: uint16(183),
	4344: uint16(1),
	4345: uint16(sym_id_selector_view),
	4346: uint16(588),
	4347: uint16(2),
	4348: uint16(anon_sym_COMMA),
	4349: uint16(anon_sym_LBRACE),
	4350: uint16(592),
	4351: uint16(3),
	4352: uint16(anon_sym_normal),
	4353: uint16(anon_sym_selected),
	4354: uint16(anon_sym_alternate),
	4355: uint16(2),
	4356: uint16(3),
	4357: uint16(1),
	4358: uint16(sym_comment),
	4359: uint16(598),
	4360: uint16(7),
	4362: uint16(anon_sym_ATimport),
	4363: uint16(anon_sym_ATtheme),
	4364: uint16(anon_sym_ATmedia),
	4365: uint16(anon_sym_STAR),
	4366: uint16(anon_sym_POUND),
	4367: uint16(sym_identifier),
	4368: uint16(5),
	4369: uint16(3),
	4370: uint16(1),
	4371: uint16(sym_comment),
	4372: uint16(590),
	4373: uint16(1),
	4374: uint16(anon_sym_DOT),
	4375: uint16(193),
	4376: uint16(1),
	4377: uint16(sym_id_selector_view),
	4378: uint16(600),
	4379: uint16(2),
	4380: uint16(anon_sym_COMMA),
	4381: uint16(anon_sym_LBRACE),
	4382: uint16(592),
	4383: uint16(3),
	4384: uint16(anon_sym_normal),
	4385: uint16(anon_sym_selected),
	4386: uint16(anon_sym_alternate),
	4387: uint16(2),
	4388: uint16(3),
	4389: uint16(1),
	4390: uint16(sym_comment),
	4391: uint16(602),
	4392: uint16(7),
	4394: uint16(anon_sym_ATimport),
	4395: uint16(anon_sym_ATtheme),
	4396: uint16(anon_sym_ATmedia),
	4397: uint16(anon_sym_STAR),
	4398: uint16(anon_sym_POUND),
	4399: uint16(sym_identifier),
	4400: uint16(2),
	4401: uint16(3),
	4402: uint16(1),
	4403: uint16(sym_comment),
	4404: uint16(604),
	4405: uint16(7),
	4407: uint16(anon_sym_ATimport),
	4408: uint16(anon_sym_ATtheme),
	4409: uint16(anon_sym_ATmedia),
	4410: uint16(anon_sym_STAR),
	4411: uint16(anon_sym_POUND),
	4412: uint16(sym_identifier),
	4413: uint16(5),
	4414: uint16(3),
	4415: uint16(1),
	4416: uint16(sym_comment),
	4417: uint16(608),
	4418: uint16(1),
	4419: uint16(anon_sym_DOT),
	4420: uint16(189),
	4421: uint16(1),
	4422: uint16(sym_id_selector_state),
	4423: uint16(606),
	4424: uint16(2),
	4425: uint16(anon_sym_COMMA),
	4426: uint16(anon_sym_LBRACE),
	4427: uint16(610),
	4428: uint16(3),
	4429: uint16(anon_sym_normal),
	4430: uint16(anon_sym_urgent),
	4431: uint16(anon_sym_active),
	4432: uint16(8),
	4433: uint16(3),
	4434: uint16(1),
	4435: uint16(sym_comment),
	4436: uint16(246),
	4437: uint16(1),
	4438: uint16(anon_sym_0),
	4439: uint16(268),
	4440: uint16(1),
	4441: uint16(sym_integer_value),
	4442: uint16(270),
	4443: uint16(1),
	4444: uint16(sym_float_value),
	4445: uint16(286),
	4446: uint16(1),
	4447: uint16(anon_sym_RPAREN),
	4448: uint16(580),
	4449: uint16(1),
	4450: uint16(anon_sym_calc),
	4451: uint16(43),
	4452: uint16(1),
	4453: uint16(sym_distance_calc),
	4454: uint16(49),
	4455: uint16(1),
	4456: uint16(sym_distance_value),
	4457: uint16(5),
	4458: uint16(3),
	4459: uint16(1),
	4460: uint16(sym_comment),
	4461: uint16(608),
	4462: uint16(1),
	4463: uint16(anon_sym_DOT),
	4464: uint16(188),
	4465: uint16(1),
	4466: uint16(sym_id_selector_state),
	4467: uint16(612),
	4468: uint16(2),
	4469: uint16(anon_sym_COMMA),
	4470: uint16(anon_sym_LBRACE),
	4471: uint16(610),
	4472: uint16(3),
	4473: uint16(anon_sym_normal),
	4474: uint16(anon_sym_urgent),
	4475: uint16(anon_sym_active),
	4476: uint16(8),
	4477: uint16(3),
	4478: uint16(1),
	4479: uint16(sym_comment),
	4480: uint16(246),
	4481: uint16(1),
	4482: uint16(anon_sym_0),
	4483: uint16(268),
	4484: uint16(1),
	4485: uint16(sym_integer_value),
	4486: uint16(270),
	4487: uint16(1),
	4488: uint16(sym_float_value),
	4489: uint16(290),
	4490: uint16(1),
	4491: uint16(anon_sym_RPAREN),
	4492: uint16(580),
	4493: uint16(1),
	4494: uint16(anon_sym_calc),
	4495: uint16(43),
	4496: uint16(1),
	4497: uint16(sym_distance_calc),
	4498: uint16(119),
	4499: uint16(1),
	4500: uint16(sym_distance_value),
	4501: uint16(7),
	4502: uint16(3),
	4503: uint16(1),
	4504: uint16(sym_comment),
	4505: uint16(614),
	4506: uint16(1),
	4507: uint16(sym_integer_value),
	4508: uint16(616),
	4509: uint16(1),
	4510: uint16(sym_float_value),
	4511: uint16(618),
	4512: uint16(1),
	4513: uint16(anon_sym_0),
	4514: uint16(620),
	4515: uint16(1),
	4516: uint16(anon_sym_calc),
	4517: uint16(94),
	4518: uint16(1),
	4519: uint16(sym_distance_value),
	4520: uint16(97),
	4521: uint16(1),
	4522: uint16(sym_distance_calc),
	4523: uint16(5),
	4524: uint16(3),
	4525: uint16(1),
	4526: uint16(sym_comment),
	4527: uint16(7),
	4528: uint16(1),
	4529: uint16(sym_identifier),
	4530: uint16(15),
	4531: uint16(1),
	4532: uint16(anon_sym_STAR),
	4533: uint16(17),
	4534: uint16(1),
	4535: uint16(anon_sym_POUND),
	4536: uint16(172),
	4537: uint16(3),
	4538: uint16(sym__selector),
	4539: uint16(sym_global_selector),
	4540: uint16(sym_id_selector),
	4541: uint16(4),
	4542: uint16(3),
	4543: uint16(1),
	4544: uint16(sym_comment),
	4545: uint16(256),
	4546: uint16(1),
	4547: uint16(anon_sym_RPAREN),
	4548: uint16(38),
	4549: uint16(1),
	4550: uint16(sym_float_distance_unit),
	4551: uint16(258),
	4552: uint16(4),
	4553: uint16(anon_sym_PERCENT),
	4554: uint16(anon_sym_cm),
	4555: uint16(anon_sym_ph),
	4556: uint16(anon_sym_em),
	4557: uint16(7),
	4558: uint16(3),
	4559: uint16(1),
	4560: uint16(sym_comment),
	4561: uint16(614),
	4562: uint16(1),
	4563: uint16(sym_integer_value),
	4564: uint16(616),
	4565: uint16(1),
	4566: uint16(sym_float_value),
	4567: uint16(618),
	4568: uint16(1),
	4569: uint16(anon_sym_0),
	4570: uint16(620),
	4571: uint16(1),
	4572: uint16(anon_sym_calc),
	4573: uint16(83),
	4574: uint16(1),
	4575: uint16(sym_distance_value),
	4576: uint16(97),
	4577: uint16(1),
	4578: uint16(sym_distance_calc),
	4579: uint16(7),
	4580: uint16(3),
	4581: uint16(1),
	4582: uint16(sym_comment),
	4583: uint16(614),
	4584: uint16(1),
	4585: uint16(sym_integer_value),
	4586: uint16(616),
	4587: uint16(1),
	4588: uint16(sym_float_value),
	4589: uint16(618),
	4590: uint16(1),
	4591: uint16(anon_sym_0),
	4592: uint16(620),
	4593: uint16(1),
	4594: uint16(anon_sym_calc),
	4595: uint16(87),
	4596: uint16(1),
	4597: uint16(sym_distance_value),
	4598: uint16(97),
	4599: uint16(1),
	4600: uint16(sym_distance_calc),
	4601: uint16(4),
	4602: uint16(3),
	4603: uint16(1),
	4604: uint16(sym_comment),
	4605: uint16(622),
	4606: uint16(1),
	4607: uint16(anon_sym_COMMA),
	4608: uint16(233),
	4609: uint16(1),
	4610: uint16(sym_angle_unit),
	4611: uint16(624),
	4612: uint16(4),
	4613: uint16(anon_sym_deg),
	4614: uint16(anon_sym_rad),
	4615: uint16(anon_sym_grad),
	4616: uint16(anon_sym_turn),
	4617: uint16(3),
	4618: uint16(3),
	4619: uint16(1),
	4620: uint16(sym_comment),
	4621: uint16(211),
	4622: uint16(1),
	4623: uint16(sym_url_image_scale),
	4624: uint16(626),
	4625: uint16(4),
	4626: uint16(anon_sym_none),
	4627: uint16(anon_sym_both),
	4628: uint16(anon_sym_width),
	4629: uint16(anon_sym_height),
	4630: uint16(3),
	4631: uint16(3),
	4632: uint16(1),
	4633: uint16(sym_comment),
	4634: uint16(235),
	4635: uint16(1),
	4636: uint16(sym_gradient_image_dir),
	4637: uint16(628),
	4638: uint16(4),
	4639: uint16(anon_sym_top),
	4640: uint16(anon_sym_left),
	4641: uint16(anon_sym_right),
	4642: uint16(anon_sym_bottom),
	4643: uint16(3),
	4644: uint16(3),
	4645: uint16(1),
	4646: uint16(sym_comment),
	4647: uint16(98),
	4648: uint16(1),
	4649: uint16(sym_float_distance_unit),
	4650: uint16(505),
	4651: uint16(4),
	4652: uint16(anon_sym_PERCENT),
	4653: uint16(anon_sym_cm),
	4654: uint16(anon_sym_ph),
	4655: uint16(anon_sym_em),
	4656: uint16(3),
	4657: uint16(3),
	4658: uint16(1),
	4659: uint16(sym_comment),
	4660: uint16(38),
	4661: uint16(1),
	4662: uint16(sym_float_distance_unit),
	4663: uint16(258),
	4664: uint16(4),
	4665: uint16(anon_sym_PERCENT),
	4666: uint16(anon_sym_cm),
	4667: uint16(anon_sym_ph),
	4668: uint16(anon_sym_em),
	4669: uint16(5),
	4670: uint16(3),
	4671: uint16(1),
	4672: uint16(sym_comment),
	4673: uint16(630),
	4674: uint16(1),
	4675: uint16(sym_integer_value),
	4676: uint16(632),
	4677: uint16(1),
	4678: uint16(sym_float_value),
	4679: uint16(634),
	4680: uint16(1),
	4681: uint16(anon_sym_0),
	4682: uint16(194),
	4683: uint16(1),
	4684: uint16(sym_percentage),
	4685: uint16(3),
	4686: uint16(3),
	4687: uint16(1),
	4688: uint16(sym_comment),
	4689: uint16(539),
	4690: uint16(1),
	4691: uint16(anon_sym_LPAREN),
	4692: uint16(184),
	4693: uint16(3),
	4694: uint16(sym__query),
	4695: uint16(sym_feature_query),
	4696: uint16(sym_parenthesized_query),
	4697: uint16(5),
	4698: uint16(3),
	4699: uint16(1),
	4700: uint16(sym_comment),
	4701: uint16(630),
	4702: uint16(1),
	4703: uint16(sym_integer_value),
	4704: uint16(632),
	4705: uint16(1),
	4706: uint16(sym_float_value),
	4707: uint16(634),
	4708: uint16(1),
	4709: uint16(anon_sym_0),
	4710: uint16(250),
	4711: uint16(1),
	4712: uint16(sym_percentage),
	4713: uint16(5),
	4714: uint16(3),
	4715: uint16(1),
	4716: uint16(sym_comment),
	4717: uint16(630),
	4718: uint16(1),
	4719: uint16(sym_integer_value),
	4720: uint16(632),
	4721: uint16(1),
	4722: uint16(sym_float_value),
	4723: uint16(634),
	4724: uint16(1),
	4725: uint16(anon_sym_0),
	4726: uint16(40),
	4727: uint16(1),
	4728: uint16(sym_percentage),
	4729: uint16(5),
	4730: uint16(3),
	4731: uint16(1),
	4732: uint16(sym_comment),
	4733: uint16(630),
	4734: uint16(1),
	4735: uint16(sym_integer_value),
	4736: uint16(632),
	4737: uint16(1),
	4738: uint16(sym_float_value),
	4739: uint16(634),
	4740: uint16(1),
	4741: uint16(anon_sym_0),
	4742: uint16(222),
	4743: uint16(1),
	4744: uint16(sym_percentage),
	4745: uint16(5),
	4746: uint16(3),
	4747: uint16(1),
	4748: uint16(sym_comment),
	4749: uint16(630),
	4750: uint16(1),
	4751: uint16(sym_integer_value),
	4752: uint16(632),
	4753: uint16(1),
	4754: uint16(sym_float_value),
	4755: uint16(634),
	4756: uint16(1),
	4757: uint16(anon_sym_0),
	4758: uint16(178),
	4759: uint16(1),
	4760: uint16(sym_percentage),
	4761: uint16(5),
	4762: uint16(3),
	4763: uint16(1),
	4764: uint16(sym_comment),
	4765: uint16(636),
	4766: uint16(1),
	4767: uint16(anon_sym_COMMA),
	4768: uint16(638),
	4769: uint16(1),
	4770: uint16(anon_sym_LBRACE),
	4771: uint16(112),
	4772: uint16(1),
	4773: uint16(sym_block),
	4774: uint16(153),
	4775: uint16(1),
	4776: uint16(aux_sym_media_statement_repeat1),
	4777: uint16(5),
	4778: uint16(3),
	4779: uint16(1),
	4780: uint16(sym_comment),
	4781: uint16(630),
	4782: uint16(1),
	4783: uint16(sym_integer_value),
	4784: uint16(632),
	4785: uint16(1),
	4786: uint16(sym_float_value),
	4787: uint16(634),
	4788: uint16(1),
	4789: uint16(anon_sym_0),
	4790: uint16(179),
	4791: uint16(1),
	4792: uint16(sym_percentage),
	4793: uint16(5),
	4794: uint16(3),
	4795: uint16(1),
	4796: uint16(sym_comment),
	4797: uint16(630),
	4798: uint16(1),
	4799: uint16(sym_integer_value),
	4800: uint16(632),
	4801: uint16(1),
	4802: uint16(sym_float_value),
	4803: uint16(634),
	4804: uint16(1),
	4805: uint16(anon_sym_0),
	4806: uint16(199),
	4807: uint16(1),
	4808: uint16(sym_percentage),
	4809: uint16(5),
	4810: uint16(3),
	4811: uint16(1),
	4812: uint16(sym_comment),
	4813: uint16(630),
	4814: uint16(1),
	4815: uint16(sym_integer_value),
	4816: uint16(632),
	4817: uint16(1),
	4818: uint16(sym_float_value),
	4819: uint16(634),
	4820: uint16(1),
	4821: uint16(anon_sym_0),
	4822: uint16(230),
	4823: uint16(1),
	4824: uint16(sym_percentage),
	4825: uint16(2),
	4826: uint16(3),
	4827: uint16(1),
	4828: uint16(sym_comment),
	4829: uint16(640),
	4830: uint16(4),
	4831: uint16(anon_sym_RBRACE),
	4832: uint16(anon_sym_STAR),
	4833: uint16(anon_sym_POUND),
	4834: uint16(sym_identifier),
	4835: uint16(5),
	4836: uint16(3),
	4837: uint16(1),
	4838: uint16(sym_comment),
	4839: uint16(630),
	4840: uint16(1),
	4841: uint16(sym_integer_value),
	4842: uint16(632),
	4843: uint16(1),
	4844: uint16(sym_float_value),
	4845: uint16(634),
	4846: uint16(1),
	4847: uint16(anon_sym_0),
	4848: uint16(212),
	4849: uint16(1),
	4850: uint16(sym_percentage),
	4851: uint16(5),
	4852: uint16(3),
	4853: uint16(1),
	4854: uint16(sym_comment),
	4855: uint16(636),
	4856: uint16(1),
	4857: uint16(anon_sym_COMMA),
	4858: uint16(638),
	4859: uint16(1),
	4860: uint16(anon_sym_LBRACE),
	4861: uint16(116),
	4862: uint16(1),
	4863: uint16(sym_block),
	4864: uint16(138),
	4865: uint16(1),
	4866: uint16(aux_sym_media_statement_repeat1),
	4867: uint16(5),
	4868: uint16(3),
	4869: uint16(1),
	4870: uint16(sym_comment),
	4871: uint16(630),
	4872: uint16(1),
	4873: uint16(sym_integer_value),
	4874: uint16(632),
	4875: uint16(1),
	4876: uint16(sym_float_value),
	4877: uint16(634),
	4878: uint16(1),
	4879: uint16(anon_sym_0),
	4880: uint16(213),
	4881: uint16(1),
	4882: uint16(sym_percentage),
	4883: uint16(2),
	4884: uint16(3),
	4885: uint16(1),
	4886: uint16(sym_comment),
	4887: uint16(642),
	4888: uint16(4),
	4889: uint16(anon_sym_RBRACE),
	4890: uint16(anon_sym_STAR),
	4891: uint16(anon_sym_POUND),
	4892: uint16(sym_identifier),
	4893: uint16(5),
	4894: uint16(3),
	4895: uint16(1),
	4896: uint16(sym_comment),
	4897: uint16(630),
	4898: uint16(1),
	4899: uint16(sym_integer_value),
	4900: uint16(632),
	4901: uint16(1),
	4902: uint16(sym_float_value),
	4903: uint16(634),
	4904: uint16(1),
	4905: uint16(anon_sym_0),
	4906: uint16(203),
	4907: uint16(1),
	4908: uint16(sym_percentage),
	4909: uint16(5),
	4910: uint16(3),
	4911: uint16(1),
	4912: uint16(sym_comment),
	4913: uint16(630),
	4914: uint16(1),
	4915: uint16(sym_integer_value),
	4916: uint16(632),
	4917: uint16(1),
	4918: uint16(sym_float_value),
	4919: uint16(634),
	4920: uint16(1),
	4921: uint16(anon_sym_0),
	4922: uint16(182),
	4923: uint16(1),
	4924: uint16(sym_percentage),
	4925: uint16(5),
	4926: uint16(3),
	4927: uint16(1),
	4928: uint16(sym_comment),
	4929: uint16(630),
	4930: uint16(1),
	4931: uint16(sym_integer_value),
	4932: uint16(632),
	4933: uint16(1),
	4934: uint16(sym_float_value),
	4935: uint16(634),
	4936: uint16(1),
	4937: uint16(anon_sym_0),
	4938: uint16(231),
	4939: uint16(1),
	4940: uint16(sym_percentage),
	4941: uint16(3),
	4942: uint16(3),
	4943: uint16(1),
	4944: uint16(sym_comment),
	4945: uint16(539),
	4946: uint16(1),
	4947: uint16(anon_sym_LPAREN),
	4948: uint16(144),
	4949: uint16(3),
	4950: uint16(sym__query),
	4951: uint16(sym_feature_query),
	4952: uint16(sym_parenthesized_query),
	4953: uint16(3),
	4954: uint16(3),
	4955: uint16(1),
	4956: uint16(sym_comment),
	4957: uint16(644),
	4958: uint16(2),
	4959: uint16(sym_integer_value),
	4960: uint16(anon_sym_0),
	4961: uint16(646),
	4962: uint16(2),
	4963: uint16(sym_float_value),
	4964: uint16(anon_sym_calc),
	4965: uint16(4),
	4966: uint16(3),
	4967: uint16(1),
	4968: uint16(sym_comment),
	4969: uint16(648),
	4970: uint16(1),
	4971: uint16(anon_sym_COMMA),
	4972: uint16(650),
	4973: uint16(1),
	4974: uint16(anon_sym_RBRACK),
	4975: uint16(165),
	4976: uint16(1),
	4977: uint16(aux_sym_list_value_repeat1),
	4978: uint16(4),
	4979: uint16(3),
	4980: uint16(1),
	4981: uint16(sym_comment),
	4982: uint16(652),
	4983: uint16(1),
	4984: uint16(anon_sym_COMMA),
	4985: uint16(655),
	4986: uint16(1),
	4987: uint16(anon_sym_LBRACE),
	4988: uint16(153),
	4989: uint16(1),
	4990: uint16(aux_sym_media_statement_repeat1),
	4991: uint16(4),
	4992: uint16(3),
	4993: uint16(1),
	4994: uint16(sym_comment),
	4995: uint16(657),
	4996: uint16(1),
	4997: uint16(anon_sym_COMMA),
	4998: uint16(659),
	4999: uint16(1),
	5000: uint16(anon_sym_LBRACE),
	5001: uint16(168),
	5002: uint16(1),
	5003: uint16(aux_sym_selectors_repeat1),
	5004: uint16(4),
	5005: uint16(3),
	5006: uint16(1),
	5007: uint16(sym_comment),
	5008: uint16(661),
	5009: uint16(1),
	5010: uint16(anon_sym_COMMA),
	5011: uint16(664),
	5012: uint16(1),
	5013: uint16(anon_sym_RBRACK),
	5014: uint16(155),
	5015: uint16(1),
	5016: uint16(aux_sym_list_value_repeat1),
	5017: uint16(4),
	5018: uint16(3),
	5019: uint16(1),
	5020: uint16(sym_comment),
	5021: uint16(666),
	5022: uint16(1),
	5023: uint16(anon_sym_COMMA),
	5024: uint16(668),
	5025: uint16(1),
	5026: uint16(anon_sym_RPAREN),
	5027: uint16(157),
	5028: uint16(1),
	5029: uint16(aux_sym_gradient_image_repeat1),
	5030: uint16(4),
	5031: uint16(3),
	5032: uint16(1),
	5033: uint16(sym_comment),
	5034: uint16(666),
	5035: uint16(1),
	5036: uint16(anon_sym_COMMA),
	5037: uint16(670),
	5038: uint16(1),
	5039: uint16(anon_sym_RPAREN),
	5040: uint16(159),
	5041: uint16(1),
	5042: uint16(aux_sym_gradient_image_repeat1),
	5043: uint16(4),
	5044: uint16(3),
	5045: uint16(1),
	5046: uint16(sym_comment),
	5047: uint16(666),
	5048: uint16(1),
	5049: uint16(anon_sym_COMMA),
	5050: uint16(672),
	5051: uint16(1),
	5052: uint16(anon_sym_RPAREN),
	5053: uint16(161),
	5054: uint16(1),
	5055: uint16(aux_sym_gradient_image_repeat1),
	5056: uint16(4),
	5057: uint16(3),
	5058: uint16(1),
	5059: uint16(sym_comment),
	5060: uint16(674),
	5061: uint16(1),
	5062: uint16(anon_sym_COMMA),
	5063: uint16(677),
	5064: uint16(1),
	5065: uint16(anon_sym_RPAREN),
	5066: uint16(159),
	5067: uint16(1),
	5068: uint16(aux_sym_gradient_image_repeat1),
	5069: uint16(4),
	5070: uint16(3),
	5071: uint16(1),
	5072: uint16(sym_comment),
	5073: uint16(666),
	5074: uint16(1),
	5075: uint16(anon_sym_COMMA),
	5076: uint16(679),
	5077: uint16(1),
	5078: uint16(anon_sym_RPAREN),
	5079: uint16(166),
	5080: uint16(1),
	5081: uint16(aux_sym_gradient_image_repeat1),
	5082: uint16(4),
	5083: uint16(3),
	5084: uint16(1),
	5085: uint16(sym_comment),
	5086: uint16(666),
	5087: uint16(1),
	5088: uint16(anon_sym_COMMA),
	5089: uint16(679),
	5090: uint16(1),
	5091: uint16(anon_sym_RPAREN),
	5092: uint16(159),
	5093: uint16(1),
	5094: uint16(aux_sym_gradient_image_repeat1),
	5095: uint16(4),
	5096: uint16(3),
	5097: uint16(1),
	5098: uint16(sym_comment),
	5099: uint16(681),
	5100: uint16(1),
	5101: uint16(anon_sym_COMMA),
	5102: uint16(684),
	5103: uint16(1),
	5104: uint16(anon_sym_LBRACE),
	5105: uint16(162),
	5106: uint16(1),
	5107: uint16(aux_sym_selectors_repeat1),
	5108: uint16(2),
	5109: uint16(3),
	5110: uint16(1),
	5111: uint16(sym_comment),
	5112: uint16(686),
	5113: uint16(3),
	5114: uint16(anon_sym_normal),
	5115: uint16(anon_sym_urgent),
	5116: uint16(anon_sym_active),
	5117: uint16(2),
	5118: uint16(3),
	5119: uint16(1),
	5120: uint16(sym_comment),
	5121: uint16(688),
	5122: uint16(3),
	5123: uint16(anon_sym_normal),
	5124: uint16(anon_sym_selected),
	5125: uint16(anon_sym_alternate),
	5126: uint16(4),
	5127: uint16(3),
	5128: uint16(1),
	5129: uint16(sym_comment),
	5130: uint16(648),
	5131: uint16(1),
	5132: uint16(anon_sym_COMMA),
	5133: uint16(690),
	5134: uint16(1),
	5135: uint16(anon_sym_RBRACK),
	5136: uint16(155),
	5137: uint16(1),
	5138: uint16(aux_sym_list_value_repeat1),
	5139: uint16(4),
	5140: uint16(3),
	5141: uint16(1),
	5142: uint16(sym_comment),
	5143: uint16(666),
	5144: uint16(1),
	5145: uint16(anon_sym_COMMA),
	5146: uint16(692),
	5147: uint16(1),
	5148: uint16(anon_sym_RPAREN),
	5149: uint16(159),
	5150: uint16(1),
	5151: uint16(aux_sym_gradient_image_repeat1),
	5152: uint16(2),
	5153: uint16(3),
	5154: uint16(1),
	5155: uint16(sym_comment),
	5156: uint16(694),
	5157: uint16(3),
	5158: uint16(anon_sym_COMMA),
	5159: uint16(anon_sym_LBRACE),
	5160: uint16(anon_sym_RPAREN),
	5161: uint16(4),
	5162: uint16(3),
	5163: uint16(1),
	5164: uint16(sym_comment),
	5165: uint16(657),
	5166: uint16(1),
	5167: uint16(anon_sym_COMMA),
	5168: uint16(696),
	5169: uint16(1),
	5170: uint16(anon_sym_LBRACE),
	5171: uint16(162),
	5172: uint16(1),
	5173: uint16(aux_sym_selectors_repeat1),
	5174: uint16(2),
	5175: uint16(3),
	5176: uint16(1),
	5177: uint16(sym_comment),
	5178: uint16(698),
	5179: uint16(3),
	5180: uint16(anon_sym_COMMA),
	5181: uint16(anon_sym_LBRACE),
	5182: uint16(anon_sym_RPAREN),
	5183: uint16(3),
	5184: uint16(3),
	5185: uint16(1),
	5186: uint16(sym_comment),
	5187: uint16(501),
	5188: uint16(1),
	5189: uint16(sym_integer_value),
	5190: uint16(225),
	5191: uint16(1),
	5192: uint16(sym_angle),
	5193: uint16(3),
	5194: uint16(3),
	5195: uint16(1),
	5196: uint16(sym_comment),
	5197: uint16(240),
	5198: uint16(1),
	5199: uint16(anon_sym_DQUOTE),
	5200: uint16(192),
	5201: uint16(1),
	5202: uint16(sym_string_value),
	5203: uint16(2),
	5204: uint16(3),
	5205: uint16(1),
	5206: uint16(sym_comment),
	5207: uint16(684),
	5208: uint16(2),
	5209: uint16(anon_sym_COMMA),
	5210: uint16(anon_sym_LBRACE),
	5211: uint16(3),
	5212: uint16(3),
	5213: uint16(1),
	5214: uint16(sym_comment),
	5215: uint16(240),
	5216: uint16(1),
	5217: uint16(anon_sym_DQUOTE),
	5218: uint16(114),
	5219: uint16(1),
	5220: uint16(sym_string_value),
	5221: uint16(3),
	5222: uint16(3),
	5223: uint16(1),
	5224: uint16(sym_comment),
	5225: uint16(700),
	5226: uint16(1),
	5227: uint16(sym_identifier),
	5228: uint16(702),
	5229: uint16(1),
	5230: uint16(anon_sym_RBRACK),
	5231: uint16(2),
	5232: uint16(3),
	5233: uint16(1),
	5234: uint16(sym_comment),
	5235: uint16(704),
	5236: uint16(2),
	5237: uint16(anon_sym_COMMA),
	5238: uint16(anon_sym_RBRACK),
	5239: uint16(2),
	5240: uint16(3),
	5241: uint16(1),
	5242: uint16(sym_comment),
	5243: uint16(677),
	5244: uint16(2),
	5245: uint16(anon_sym_COMMA),
	5246: uint16(anon_sym_RPAREN),
	5247: uint16(3),
	5248: uint16(3),
	5249: uint16(1),
	5250: uint16(sym_comment),
	5251: uint16(706),
	5252: uint16(1),
	5253: uint16(anon_sym_COMMA),
	5254: uint16(708),
	5255: uint16(1),
	5256: uint16(anon_sym_RPAREN),
	5257: uint16(3),
	5258: uint16(3),
	5259: uint16(1),
	5260: uint16(sym_comment),
	5261: uint16(710),
	5262: uint16(1),
	5263: uint16(anon_sym_COMMA),
	5264: uint16(712),
	5265: uint16(1),
	5266: uint16(anon_sym_RPAREN),
	5267: uint16(3),
	5268: uint16(3),
	5269: uint16(1),
	5270: uint16(sym_comment),
	5271: uint16(714),
	5272: uint16(1),
	5273: uint16(anon_sym_COMMA),
	5274: uint16(716),
	5275: uint16(1),
	5276: uint16(anon_sym_RPAREN),
	5277: uint16(3),
	5278: uint16(3),
	5279: uint16(1),
	5280: uint16(sym_comment),
	5281: uint16(240),
	5282: uint16(1),
	5283: uint16(anon_sym_DQUOTE),
	5284: uint16(117),
	5285: uint16(1),
	5286: uint16(sym_string_value),
	5287: uint16(3),
	5288: uint16(3),
	5289: uint16(1),
	5290: uint16(sym_comment),
	5291: uint16(718),
	5292: uint16(1),
	5293: uint16(anon_sym_COMMA),
	5294: uint16(720),
	5295: uint16(1),
	5296: uint16(anon_sym_RPAREN),
	5297: uint16(3),
	5298: uint16(3),
	5299: uint16(1),
	5300: uint16(sym_comment),
	5301: uint16(722),
	5302: uint16(1),
	5303: uint16(anon_sym_COMMA),
	5304: uint16(724),
	5305: uint16(1),
	5306: uint16(anon_sym_RPAREN),
	5307: uint16(2),
	5308: uint16(3),
	5309: uint16(1),
	5310: uint16(sym_comment),
	5311: uint16(726),
	5312: uint16(2),
	5313: uint16(anon_sym_COMMA),
	5314: uint16(anon_sym_LBRACE),
	5315: uint16(2),
	5316: uint16(3),
	5317: uint16(1),
	5318: uint16(sym_comment),
	5319: uint16(655),
	5320: uint16(2),
	5321: uint16(anon_sym_COMMA),
	5322: uint16(anon_sym_LBRACE),
	5323: uint16(2),
	5324: uint16(3),
	5325: uint16(1),
	5326: uint16(sym_comment),
	5327: uint16(728),
	5328: uint16(2),
	5329: uint16(anon_sym_COMMA),
	5330: uint16(anon_sym_LBRACE),
	5331: uint16(3),
	5332: uint16(3),
	5333: uint16(1),
	5334: uint16(sym_comment),
	5335: uint16(638),
	5336: uint16(1),
	5337: uint16(anon_sym_LBRACE),
	5338: uint16(105),
	5339: uint16(1),
	5340: uint16(sym_block),
	5341: uint16(2),
	5342: uint16(3),
	5343: uint16(1),
	5344: uint16(sym_comment),
	5345: uint16(730),
	5346: uint16(2),
	5347: uint16(anon_sym_COMMA),
	5348: uint16(anon_sym_LBRACE),
	5349: uint16(2),
	5350: uint16(3),
	5351: uint16(1),
	5352: uint16(sym_comment),
	5353: uint16(732),
	5354: uint16(2),
	5355: uint16(anon_sym_COMMA),
	5356: uint16(anon_sym_LBRACE),
	5357: uint16(2),
	5358: uint16(3),
	5359: uint16(1),
	5360: uint16(sym_comment),
	5361: uint16(734),
	5362: uint16(2),
	5363: uint16(anon_sym_COMMA),
	5364: uint16(anon_sym_LBRACE),
	5365: uint16(2),
	5366: uint16(3),
	5367: uint16(1),
	5368: uint16(sym_comment),
	5369: uint16(736),
	5370: uint16(2),
	5371: uint16(anon_sym_COMMA),
	5372: uint16(anon_sym_LBRACE),
	5373: uint16(3),
	5374: uint16(3),
	5375: uint16(1),
	5376: uint16(sym_comment),
	5377: uint16(501),
	5378: uint16(1),
	5379: uint16(sym_integer_value),
	5380: uint16(228),
	5381: uint16(1),
	5382: uint16(sym_angle),
	5383: uint16(3),
	5384: uint16(3),
	5385: uint16(1),
	5386: uint16(sym_comment),
	5387: uint16(738),
	5388: uint16(1),
	5389: uint16(anon_sym_COMMA),
	5390: uint16(740),
	5391: uint16(1),
	5392: uint16(anon_sym_RPAREN),
	5393: uint16(2),
	5394: uint16(3),
	5395: uint16(1),
	5396: uint16(sym_comment),
	5397: uint16(742),
	5398: uint16(2),
	5399: uint16(anon_sym_COMMA),
	5400: uint16(anon_sym_LBRACE),
	5401: uint16(2),
	5402: uint16(3),
	5403: uint16(1),
	5404: uint16(sym_comment),
	5405: uint16(744),
	5406: uint16(1),
	5407: uint16(anon_sym_RPAREN),
	5408: uint16(2),
	5409: uint16(3),
	5410: uint16(1),
	5411: uint16(sym_comment),
	5412: uint16(746),
	5413: uint16(1),
	5414: uint16(anon_sym_LPAREN),
	5415: uint16(2),
	5416: uint16(3),
	5417: uint16(1),
	5418: uint16(sym_comment),
	5419: uint16(748),
	5420: uint16(1),
	5421: uint16(anon_sym_RPAREN),
	5422: uint16(2),
	5423: uint16(3),
	5424: uint16(1),
	5425: uint16(sym_comment),
	5426: uint16(750),
	5427: uint16(1),
	5428: uint16(anon_sym_RPAREN),
	5429: uint16(2),
	5430: uint16(3),
	5431: uint16(1),
	5432: uint16(sym_comment),
	5433: uint16(752),
	5434: uint16(1),
	5435: uint16(anon_sym_COMMA),
	5436: uint16(2),
	5437: uint16(3),
	5438: uint16(1),
	5439: uint16(sym_comment),
	5440: uint16(754),
	5441: uint16(1),
	5442: uint16(anon_sym_COMMA),
	5443: uint16(2),
	5444: uint16(3),
	5445: uint16(1),
	5446: uint16(sym_comment),
	5447: uint16(756),
	5448: uint16(1),
	5450: uint16(2),
	5451: uint16(3),
	5452: uint16(1),
	5453: uint16(sym_comment),
	5454: uint16(758),
	5455: uint16(1),
	5456: uint16(anon_sym_LPAREN),
	5457: uint16(2),
	5458: uint16(3),
	5459: uint16(1),
	5460: uint16(sym_comment),
	5461: uint16(760),
	5462: uint16(1),
	5463: uint16(anon_sym_LPAREN),
	5464: uint16(2),
	5465: uint16(3),
	5466: uint16(1),
	5467: uint16(sym_comment),
	5468: uint16(762),
	5469: uint16(1),
	5470: uint16(anon_sym_COMMA),
	5471: uint16(2),
	5472: uint16(3),
	5473: uint16(1),
	5474: uint16(sym_comment),
	5475: uint16(764),
	5476: uint16(1),
	5477: uint16(anon_sym_COMMA),
	5478: uint16(2),
	5479: uint16(3),
	5480: uint16(1),
	5481: uint16(sym_comment),
	5482: uint16(766),
	5483: uint16(1),
	5484: uint16(anon_sym_LPAREN),
	5485: uint16(2),
	5486: uint16(3),
	5487: uint16(1),
	5488: uint16(sym_comment),
	5489: uint16(768),
	5490: uint16(1),
	5491: uint16(sym_identifier),
	5492: uint16(2),
	5493: uint16(3),
	5494: uint16(1),
	5495: uint16(sym_comment),
	5496: uint16(770),
	5497: uint16(1),
	5498: uint16(anon_sym_LPAREN),
	5499: uint16(2),
	5500: uint16(3),
	5501: uint16(1),
	5502: uint16(sym_comment),
	5503: uint16(772),
	5504: uint16(1),
	5505: uint16(anon_sym_LPAREN),
	5506: uint16(2),
	5507: uint16(3),
	5508: uint16(1),
	5509: uint16(sym_comment),
	5510: uint16(774),
	5511: uint16(1),
	5512: uint16(anon_sym_RPAREN),
	5513: uint16(2),
	5514: uint16(3),
	5515: uint16(1),
	5516: uint16(sym_comment),
	5517: uint16(776),
	5518: uint16(1),
	5519: uint16(anon_sym_RBRACE),
	5520: uint16(2),
	5521: uint16(3),
	5522: uint16(1),
	5523: uint16(sym_comment),
	5524: uint16(778),
	5525: uint16(1),
	5526: uint16(anon_sym_RPAREN),
	5527: uint16(2),
	5528: uint16(3),
	5529: uint16(1),
	5530: uint16(sym_comment),
	5531: uint16(780),
	5532: uint16(1),
	5533: uint16(anon_sym_RPAREN),
	5534: uint16(2),
	5535: uint16(3),
	5536: uint16(1),
	5537: uint16(sym_comment),
	5538: uint16(782),
	5539: uint16(1),
	5540: uint16(anon_sym_RPAREN),
	5541: uint16(2),
	5542: uint16(3),
	5543: uint16(1),
	5544: uint16(sym_comment),
	5545: uint16(784),
	5546: uint16(1),
	5547: uint16(aux_sym_hex_color_token1),
	5548: uint16(2),
	5549: uint16(3),
	5550: uint16(1),
	5551: uint16(sym_comment),
	5552: uint16(786),
	5553: uint16(1),
	5554: uint16(sym_identifier),
	5555: uint16(2),
	5556: uint16(3),
	5557: uint16(1),
	5558: uint16(sym_comment),
	5559: uint16(788),
	5560: uint16(1),
	5561: uint16(anon_sym_LPAREN),
	5562: uint16(2),
	5563: uint16(3),
	5564: uint16(1),
	5565: uint16(sym_comment),
	5566: uint16(790),
	5567: uint16(1),
	5568: uint16(anon_sym_COMMA),
	5569: uint16(2),
	5570: uint16(3),
	5571: uint16(1),
	5572: uint16(sym_comment),
	5573: uint16(792),
	5574: uint16(1),
	5575: uint16(anon_sym_LPAREN),
	5576: uint16(2),
	5577: uint16(3),
	5578: uint16(1),
	5579: uint16(sym_comment),
	5580: uint16(794),
	5581: uint16(1),
	5582: uint16(sym_identifier),
	5583: uint16(2),
	5584: uint16(3),
	5585: uint16(1),
	5586: uint16(sym_comment),
	5587: uint16(796),
	5588: uint16(1),
	5589: uint16(anon_sym_COMMA),
	5590: uint16(2),
	5591: uint16(3),
	5592: uint16(1),
	5593: uint16(sym_comment),
	5594: uint16(798),
	5595: uint16(1),
	5596: uint16(anon_sym_COMMA),
	5597: uint16(2),
	5598: uint16(3),
	5599: uint16(1),
	5600: uint16(sym_comment),
	5601: uint16(800),
	5602: uint16(1),
	5603: uint16(anon_sym_COMMA),
	5604: uint16(2),
	5605: uint16(3),
	5606: uint16(1),
	5607: uint16(sym_comment),
	5608: uint16(802),
	5609: uint16(1),
	5610: uint16(sym_integer_value),
	5611: uint16(2),
	5612: uint16(3),
	5613: uint16(1),
	5614: uint16(sym_comment),
	5615: uint16(804),
	5616: uint16(1),
	5617: uint16(anon_sym_DQUOTE),
	5618: uint16(2),
	5619: uint16(3),
	5620: uint16(1),
	5621: uint16(sym_comment),
	5622: uint16(806),
	5623: uint16(1),
	5624: uint16(anon_sym_COMMA),
	5625: uint16(2),
	5626: uint16(3),
	5627: uint16(1),
	5628: uint16(sym_comment),
	5629: uint16(808),
	5630: uint16(1),
	5631: uint16(anon_sym_COMMA),
	5632: uint16(2),
	5633: uint16(3),
	5634: uint16(1),
	5635: uint16(sym_comment),
	5636: uint16(810),
	5637: uint16(1),
	5638: uint16(anon_sym_COLON),
	5639: uint16(2),
	5640: uint16(3),
	5641: uint16(1),
	5642: uint16(sym_comment),
	5643: uint16(812),
	5644: uint16(1),
	5645: uint16(anon_sym_COMMA),
	5646: uint16(2),
	5647: uint16(3),
	5648: uint16(1),
	5649: uint16(sym_comment),
	5650: uint16(319),
	5651: uint16(1),
	5652: uint16(anon_sym_PERCENT),
	5653: uint16(2),
	5654: uint16(3),
	5655: uint16(1),
	5656: uint16(sym_comment),
	5657: uint16(814),
	5658: uint16(1),
	5659: uint16(anon_sym_COMMA),
	5660: uint16(2),
	5661: uint16(3),
	5662: uint16(1),
	5663: uint16(sym_comment),
	5664: uint16(816),
	5665: uint16(1),
	5666: uint16(anon_sym_COMMA),
	5667: uint16(2),
	5668: uint16(3),
	5669: uint16(1),
	5670: uint16(sym_comment),
	5671: uint16(818),
	5672: uint16(1),
	5673: uint16(anon_sym_COMMA),
	5674: uint16(2),
	5675: uint16(3),
	5676: uint16(1),
	5677: uint16(sym_comment),
	5678: uint16(820),
	5679: uint16(1),
	5680: uint16(anon_sym_COMMA),
	5681: uint16(2),
	5682: uint16(3),
	5683: uint16(1),
	5684: uint16(sym_comment),
	5685: uint16(822),
	5686: uint16(1),
	5687: uint16(anon_sym_COMMA),
	5688: uint16(2),
	5689: uint16(3),
	5690: uint16(1),
	5691: uint16(sym_comment),
	5692: uint16(824),
	5693: uint16(1),
	5694: uint16(anon_sym_COMMA),
	5695: uint16(2),
	5696: uint16(3),
	5697: uint16(1),
	5698: uint16(sym_comment),
	5699: uint16(826),
	5700: uint16(1),
	5701: uint16(sym_identifier),
	5702: uint16(2),
	5703: uint16(3),
	5704: uint16(1),
	5705: uint16(sym_comment),
	5706: uint16(828),
	5707: uint16(1),
	5708: uint16(anon_sym_COMMA),
	5709: uint16(2),
	5710: uint16(3),
	5711: uint16(1),
	5712: uint16(sym_comment),
	5713: uint16(830),
	5714: uint16(1),
	5715: uint16(sym_integer_value),
	5716: uint16(2),
	5717: uint16(3),
	5718: uint16(1),
	5719: uint16(sym_comment),
	5720: uint16(832),
	5721: uint16(1),
	5722: uint16(sym_integer_value),
	5723: uint16(2),
	5724: uint16(3),
	5725: uint16(1),
	5726: uint16(sym_comment),
	5727: uint16(834),
	5728: uint16(1),
	5729: uint16(sym_identifier),
	5730: uint16(2),
	5731: uint16(3),
	5732: uint16(1),
	5733: uint16(sym_comment),
	5734: uint16(836),
	5735: uint16(1),
	5736: uint16(anon_sym_LPAREN),
	5737: uint16(2),
	5738: uint16(838),
	5739: uint16(1),
	5740: uint16(aux_sym_string_value_token1),
	5741: uint16(840),
	5742: uint16(1),
	5743: uint16(sym_comment),
	5744: uint16(2),
	5745: uint16(3),
	5746: uint16(1),
	5747: uint16(sym_comment),
	5748: uint16(842),
	5749: uint16(1),
	5750: uint16(anon_sym_LBRACE2),
	5751: uint16(2),
	5752: uint16(3),
	5753: uint16(1),
	5754: uint16(sym_comment),
	5755: uint16(844),
	5756: uint16(1),
	5757: uint16(anon_sym_DQUOTE),
	5758: uint16(2),
	5759: uint16(3),
	5760: uint16(1),
	5761: uint16(sym_comment),
	5762: uint16(846),
	5763: uint16(1),
	5764: uint16(anon_sym_RPAREN),
	5765: uint16(2),
	5766: uint16(3),
	5767: uint16(1),
	5768: uint16(sym_comment),
	5769: uint16(848),
	5770: uint16(1),
	5771: uint16(anon_sym_COLON),
	5772: uint16(2),
	5773: uint16(840),
	5774: uint16(1),
	5775: uint16(sym_comment),
	5776: uint16(850),
	5777: uint16(1),
	5778: uint16(aux_sym_string_value_token1),
	5779: uint16(2),
	5780: uint16(3),
	5781: uint16(1),
	5782: uint16(sym_comment),
	5783: uint16(852),
	5784: uint16(1),
	5785: uint16(sym_identifier),
	5786: uint16(2),
	5787: uint16(3),
	5788: uint16(1),
	5789: uint16(sym_comment),
	5790: uint16(854),
	5791: uint16(1),
	5792: uint16(anon_sym_LPAREN),
	5793: uint16(2),
	5794: uint16(3),
	5795: uint16(1),
	5796: uint16(sym_comment),
	5797: uint16(856),
	5798: uint16(1),
	5799: uint16(anon_sym_RPAREN),
}

var ts_small_parse_table_map = [239]uint32_t{
	1:   uint32(68),
	2:   uint32(136),
	3:   uint32(199),
	4:   uint32(270),
	5:   uint32(339),
	6:   uint32(408),
	7:   uint32(477),
	8:   uint32(543),
	9:   uint32(609),
	10:  uint32(666),
	11:  uint32(723),
	12:  uint32(780),
	13:  uint32(835),
	14:  uint32(890),
	15:  uint32(942),
	16:  uint32(994),
	17:  uint32(1046),
	18:  uint32(1098),
	19:  uint32(1150),
	20:  uint32(1202),
	21:  uint32(1254),
	22:  uint32(1306),
	23:  uint32(1358),
	24:  uint32(1410),
	25:  uint32(1462),
	26:  uint32(1514),
	27:  uint32(1566),
	28:  uint32(1618),
	29:  uint32(1670),
	30:  uint32(1722),
	31:  uint32(1774),
	32:  uint32(1826),
	33:  uint32(1878),
	34:  uint32(1930),
	35:  uint32(1982),
	36:  uint32(2034),
	37:  uint32(2086),
	38:  uint32(2138),
	39:  uint32(2190),
	40:  uint32(2242),
	41:  uint32(2294),
	42:  uint32(2346),
	43:  uint32(2398),
	44:  uint32(2450),
	45:  uint32(2502),
	46:  uint32(2554),
	47:  uint32(2606),
	48:  uint32(2658),
	49:  uint32(2710),
	50:  uint32(2762),
	51:  uint32(2814),
	52:  uint32(2866),
	53:  uint32(2918),
	54:  uint32(2970),
	55:  uint32(3022),
	56:  uint32(3074),
	57:  uint32(3126),
	58:  uint32(3178),
	59:  uint32(3229),
	60:  uint32(3280),
	61:  uint32(3323),
	62:  uint32(3356),
	63:  uint32(3390),
	64:  uint32(3424),
	65:  uint32(3464),
	66:  uint32(3498),
	67:  uint32(3532),
	68:  uint32(3566),
	69:  uint32(3600),
	70:  uint32(3640),
	71:  uint32(3669),
	72:  uint32(3696),
	73:  uint32(3721),
	74:  uint32(3748),
	75:  uint32(3775),
	76:  uint32(3802),
	77:  uint32(3829),
	78:  uint32(3859),
	79:  uint32(3889),
	80:  uint32(3919),
	81:  uint32(3938),
	82:  uint32(3957),
	83:  uint32(3976),
	84:  uint32(3995),
	85:  uint32(4014),
	86:  uint32(4033),
	87:  uint32(4052),
	88:  uint32(4071),
	89:  uint32(4103),
	90:  uint32(4118),
	91:  uint32(4139),
	92:  uint32(4167),
	93:  uint32(4181),
	94:  uint32(4195),
	95:  uint32(4213),
	96:  uint32(4227),
	97:  uint32(4245),
	98:  uint32(4273),
	99:  uint32(4301),
	100: uint32(4323),
	101: uint32(4336),
	102: uint32(4355),
	103: uint32(4368),
	104: uint32(4387),
	105: uint32(4400),
	106: uint32(4413),
	107: uint32(4432),
	108: uint32(4457),
	109: uint32(4476),
	110: uint32(4501),
	111: uint32(4523),
	112: uint32(4541),
	113: uint32(4557),
	114: uint32(4579),
	115: uint32(4601),
	116: uint32(4617),
	117: uint32(4630),
	118: uint32(4643),
	119: uint32(4656),
	120: uint32(4669),
	121: uint32(4685),
	122: uint32(4697),
	123: uint32(4713),
	124: uint32(4729),
	125: uint32(4745),
	126: uint32(4761),
	127: uint32(4777),
	128: uint32(4793),
	129: uint32(4809),
	130: uint32(4825),
	131: uint32(4835),
	132: uint32(4851),
	133: uint32(4867),
	134: uint32(4883),
	135: uint32(4893),
	136: uint32(4909),
	137: uint32(4925),
	138: uint32(4941),
	139: uint32(4953),
	140: uint32(4965),
	141: uint32(4978),
	142: uint32(4991),
	143: uint32(5004),
	144: uint32(5017),
	145: uint32(5030),
	146: uint32(5043),
	147: uint32(5056),
	148: uint32(5069),
	149: uint32(5082),
	150: uint32(5095),
	151: uint32(5108),
	152: uint32(5117),
	153: uint32(5126),
	154: uint32(5139),
	155: uint32(5152),
	156: uint32(5161),
	157: uint32(5174),
	158: uint32(5183),
	159: uint32(5193),
	160: uint32(5203),
	161: uint32(5211),
	162: uint32(5221),
	163: uint32(5231),
	164: uint32(5239),
	165: uint32(5247),
	166: uint32(5257),
	167: uint32(5267),
	168: uint32(5277),
	169: uint32(5287),
	170: uint32(5297),
	171: uint32(5307),
	172: uint32(5315),
	173: uint32(5323),
	174: uint32(5331),
	175: uint32(5341),
	176: uint32(5349),
	177: uint32(5357),
	178: uint32(5365),
	179: uint32(5373),
	180: uint32(5383),
	181: uint32(5393),
	182: uint32(5401),
	183: uint32(5408),
	184: uint32(5415),
	185: uint32(5422),
	186: uint32(5429),
	187: uint32(5436),
	188: uint32(5443),
	189: uint32(5450),
	190: uint32(5457),
	191: uint32(5464),
	192: uint32(5471),
	193: uint32(5478),
	194: uint32(5485),
	195: uint32(5492),
	196: uint32(5499),
	197: uint32(5506),
	198: uint32(5513),
	199: uint32(5520),
	200: uint32(5527),
	201: uint32(5534),
	202: uint32(5541),
	203: uint32(5548),
	204: uint32(5555),
	205: uint32(5562),
	206: uint32(5569),
	207: uint32(5576),
	208: uint32(5583),
	209: uint32(5590),
	210: uint32(5597),
	211: uint32(5604),
	212: uint32(5611),
	213: uint32(5618),
	214: uint32(5625),
	215: uint32(5632),
	216: uint32(5639),
	217: uint32(5646),
	218: uint32(5653),
	219: uint32(5660),
	220: uint32(5667),
	221: uint32(5674),
	222: uint32(5681),
	223: uint32(5688),
	224: uint32(5695),
	225: uint32(5702),
	226: uint32(5709),
	227: uint32(5716),
	228: uint32(5723),
	229: uint32(5730),
	230: uint32(5737),
	231: uint32(5744),
	232: uint32(5751),
	233: uint32(5758),
	234: uint32(5765),
	235: uint32(5772),
	236: uint32(5779),
	237: uint32(5786),
	238: uint32(5793),
}

var ts_parse_actions = [858]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_stylesheet),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(180)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(173)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(150)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(185)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(219)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(214)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(146)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(71)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(247)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(195)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(202)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(207)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(208)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(218)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(241)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(205)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(34)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(174)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	70: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(201)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	74: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	76: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(24)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	79: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	80: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	82: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	83: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
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
		Fcount: uint8(2),
	}})),
	87: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(71)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	90: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	91: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(247)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	92: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	93: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(12)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	96: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	97: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(14)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	98: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	99: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	100: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(69)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	102: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	103: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(195)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	104: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	105: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(202)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	108: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	109: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(13)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	110: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	111: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(207)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	114: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	115: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(208)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	116: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	117: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(218)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	120: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	121: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(241)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	122: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	123: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(31)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	126: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fcount: uint8(2),
	}})),
	129: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	130: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(205)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	132: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(21)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	135: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	136: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(215)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	138: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(216)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	141: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	142: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(34)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	144: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(36)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(174)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	150: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
	151: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(243)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	152: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	153: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(26),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(201)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(169)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Fcount: uint8(2),
	}})),
	160: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(24)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(214)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
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
		Fcount: uint8(2),
	}})),
	168: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(6)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	171: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(247)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
	175: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(12)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	177: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(14)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	180: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
	181: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(69)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	182: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	183: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(195)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	186: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
	187: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(202)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	188: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	189: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(13)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	192: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(207)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	195: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(208)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(218)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(241)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	204: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
	205: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(31)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	207: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(33)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	210: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(205)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	213: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(21)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	216: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(215)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	219: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(216)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	222: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
	223: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(34)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	225: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(36)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	228: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(174)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(243)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	234: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_feature_query_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(201)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(197)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(242)),
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
		Fcount: uint8(1),
	}})),
	243: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(124)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(70)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(196)),
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
		Fcount: uint8(1),
	}})),
	255: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__value),
	})))),
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
		Fsymbol:      uint16(sym__value),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(35)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(35)),
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
		Fcount: uint8(1),
	}})),
	265: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_distance_value),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_distance_value),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(106)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(131)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_border_value),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_border_value),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_border_value),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_border_value),
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
		Fsymbol:      uint16(sym_border_value),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_border_value),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_padding_value),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_padding_value),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_padding_value),
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
		Fsymbol:      uint16(sym_padding_value),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_position_value),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_position_value),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	299: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_border_style),
		Fproduction_id: uint16(19),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	301: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_border_style),
		Fproduction_id: uint16(19),
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
		Fcount: uint8(1),
	}})),
	303: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_position_value_repeat1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_position_value_repeat1),
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
		Fcount: uint8(2),
	}})),
	307: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_position_value_repeat1),
	})))),
	308: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(23)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	309: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	310: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_named_color),
		Fproduction_id: uint16(15),
	})))),
	311: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	312: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_named_color),
		Fproduction_id: uint16(15),
	})))),
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
		Fcount: uint8(1),
	}})),
	314: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	315: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	316: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_percentage),
	})))),
	317: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	318: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_percentage),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	320: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	321: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	322: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_percentage),
	})))),
	323: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	324: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_percentage),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_distance_value),
	})))),
	327: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	328: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_distance_value),
	})))),
	329: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	330: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_reference_value),
		Fproduction_id: uint16(17),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	332: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_reference_value),
		Fproduction_id: uint16(17),
	})))),
	333: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	334: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	335: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	336: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	338: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_first_border_style),
		Fproduction_id: uint16(18),
	})))),
	339: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	340: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_first_border_style),
		Fproduction_id: uint16(18),
	})))),
	341: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	342: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_text_style_value),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_text_style_value),
	})))),
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
		Fcount: uint8(1),
	}})),
	346: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_hex_color),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_hex_color),
	})))),
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
		Fcount: uint8(1),
	}})),
	350: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_line_style_value),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	352: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_line_style_value),
	})))),
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
		Fcount: uint8(1),
	}})),
	354: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_orientation_value),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_orientation_value),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_float_distance_unit),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_float_distance_unit),
	})))),
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
		Fcount: uint8(1),
	}})),
	362: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_cursor_value),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	364: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_cursor_value),
	})))),
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
		Fcount: uint8(1),
	}})),
	366: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_integer_distance_unit),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	368: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_integer_distance_unit),
	})))),
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
		Fcount: uint8(1),
	}})),
	370: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_distance_value),
		Fproduction_id: uint16(16),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	372: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_distance_value),
		Fproduction_id: uint16(16),
	})))),
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
		Fcount: uint8(1),
	}})),
	374: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_list_value),
		Fproduction_id: uint16(23),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	376: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_list_value),
		Fproduction_id: uint16(23),
	})))),
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
		Fcount: uint8(1),
	}})),
	378: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_named_color),
		Fproduction_id: uint16(24),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	380: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_named_color),
		Fproduction_id: uint16(24),
	})))),
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
		Fcount: uint8(1),
	}})),
	382: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_image_value),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	384: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_image_value),
	})))),
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
		Fcount: uint8(1),
	}})),
	386: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_border_style),
		Fproduction_id: uint16(18),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	388: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_border_style),
		Fproduction_id: uint16(18),
	})))),
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
		Fcount: uint8(1),
	}})),
	390: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_url_image),
		Fproduction_id: uint16(27),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	392: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_url_image),
		Fproduction_id: uint16(27),
	})))),
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
		Fcount: uint8(1),
	}})),
	394: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_distance_calc),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	396: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_distance_calc),
	})))),
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
		Fcount: uint8(1),
	}})),
	398: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_reference_value),
		Fproduction_id: uint16(28),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	400: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_reference_value),
		Fproduction_id: uint16(28),
	})))),
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
		Fcount: uint8(1),
	}})),
	402: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_list_value),
		Fproduction_id: uint16(29),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	404: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_list_value),
		Fproduction_id: uint16(29),
	})))),
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
		Fcount: uint8(1),
	}})),
	406: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_environ_value),
		Fproduction_id: uint16(28),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_environ_value),
		Fproduction_id: uint16(28),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_padding_value),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_padding_value),
	})))),
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
		Fcount: uint8(1),
	}})),
	414: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_border_value),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	416: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_border_value),
	})))),
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
		Fcount: uint8(1),
	}})),
	418: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_distance_calc),
		Fproduction_id: uint16(31),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	420: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_distance_calc),
		Fproduction_id: uint16(31),
	})))),
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
		Fcount: uint8(1),
	}})),
	422: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_url_image),
		Fproduction_id: uint16(33),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_url_image),
		Fproduction_id: uint16(33),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_gradient_image),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_gradient_image),
	})))),
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
		Fcount: uint8(1),
	}})),
	430: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_reference_value),
		Fproduction_id: uint16(34),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	432: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_reference_value),
		Fproduction_id: uint16(34),
	})))),
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
		Fcount: uint8(1),
	}})),
	434: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_environ_value),
		Fproduction_id: uint16(34),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_environ_value),
		Fproduction_id: uint16(34),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_gradient_image),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_gradient_image),
	})))),
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
		Fcount: uint8(1),
	}})),
	442: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_gradient_image),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	444: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_gradient_image),
	})))),
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
		Fcount: uint8(1),
	}})),
	446: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_rgb_color),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_rgb_color),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_hsl_color),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_hsl_color),
	})))),
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
		Fcount: uint8(1),
	}})),
	454: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_hwb_color),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	456: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_hwb_color),
	})))),
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
		Fcount: uint8(1),
	}})),
	458: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_gradient_image),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_gradient_image),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_gradient_image),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_gradient_image),
	})))),
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
		Fcount: uint8(1),
	}})),
	466: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_rgb_color),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	468: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_rgb_color),
	})))),
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
		Fcount: uint8(1),
	}})),
	470: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_hsl_color),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_hsl_color),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_hwb_color),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_hwb_color),
	})))),
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
		Fcount: uint8(1),
	}})),
	478: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_cmyk_color),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	480: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_cmyk_color),
	})))),
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
		Fcount: uint8(1),
	}})),
	482: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(12),
		Fsymbol:      uint16(sym_cmyk_color),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	484: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(12),
		Fsymbol:      uint16(sym_cmyk_color),
	})))),
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
		Fcount: uint8(1),
	}})),
	490: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	494: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(23),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(23),
	})))),
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
		Fcount: uint8(1),
	}})),
	498: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(22),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_declaration_repeat1),
		Fproduction_id: uint16(22),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(127)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(129)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(93)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_stylesheet),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_stylesheet_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	514: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_stylesheet_repeat1),
	})))),
	515: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(113)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	517: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_stylesheet_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(180)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	520: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_stylesheet_repeat1),
	})))),
	521: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(173)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_stylesheet_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(150)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	526: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_stylesheet_repeat1),
	})))),
	527: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(185)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	529: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_stylesheet_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(219)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(151)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(151)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(227)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	544: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_distance_calc_repeat1),
		Fproduction_id: uint16(32),
	})))),
	545: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(151)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_distance_calc_repeat1),
		Fproduction_id: uint16(32),
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
		Fcount: uint8(2),
	}})),
	549: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_distance_calc_repeat1),
		Fproduction_id: uint16(32),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(151)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(51)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(99)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(92)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(111)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(107)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(104)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	564: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
	565: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(111)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	569: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(185)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	572: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
	573: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(219)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_distance_calc_repeat1),
		Fproduction_id: uint16(30),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_distance_calc_repeat1),
		Fproduction_id: uint16(30),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	581: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(205)),
	}})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_block),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	585: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_rule_set),
		Fproduction_id: uint16(6),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_block),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	589: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_id_selector),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(118)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	595: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_media_statement),
		Fproduction_id: uint16(12),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_theme_statement),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	601: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_id_selector),
		Fproduction_id: uint16(3),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_media_statement),
		Fproduction_id: uint16(7),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	605: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_import_statement),
		Fproduction_id: uint16(2),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	607: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_id_selector_view),
		Fproduction_id: uint16(4),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(187)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_id_selector_view),
		Fproduction_id: uint16(9),
	})))),
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
		Fcount: uint8(1),
	}})),
	615: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(108)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	617: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	618: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	619: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	620: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(249)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	623: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_angle),
		Fproduction_id: uint16(22),
	})))),
	624: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	625: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	626: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(209)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(234)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(229)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	635: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	636: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	637: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	638: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(89)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(21),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(25),
	})))),
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
		Fcount: uint8(1),
	}})),
	645: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_distance_op),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_distance_op),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	651: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_media_statement_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(133)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_media_statement_repeat1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(123)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_selectors),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	662: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_list_value_repeat1),
		Fproduction_id: uint16(26),
	})))),
	663: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(206)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	665: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_list_value_repeat1),
		Fproduction_id: uint16(26),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	669: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	670: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	671: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	672: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(57)),
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
		Fsymbol:      uint16(aux_sym_gradient_image_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(79)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_gradient_image_repeat1),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	682: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_selectors_repeat1),
	})))),
	683: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(123)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	684: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	685: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_selectors_repeat1),
	})))),
	686: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	687: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	688: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	689: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(120)),
	}})))),
	690: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	691: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	692: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	693: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	694: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	695: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_parenthesized_query),
	})))),
	696: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	697: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_selectors),
	})))),
	698: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	699: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_feature_query),
		Fproduction_id: uint16(20),
	})))),
	700: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	701: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	702: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	703: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	704: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	705: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_list_value_repeat1),
		Fproduction_id: uint16(23),
	})))),
	706: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	707: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(143)),
	}})))),
	708: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	709: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	710: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	711: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	712: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	713: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	714: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	715: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	716: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	717: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	718: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	719: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	720: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	721: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	722: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	723: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(134)),
	}})))),
	724: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	725: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	726: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	727: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_id_selector),
		Fproduction_id: uint16(5),
	})))),
	728: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	729: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_global_selector),
	})))),
	730: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	731: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_id_selector_state),
		Fproduction_id: uint16(10),
	})))),
	732: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	733: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_id_selector_view),
		Fproduction_id: uint16(13),
	})))),
	734: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	735: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_id_selector_view),
		Fproduction_id: uint16(11),
	})))),
	736: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	737: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_id_selector_state),
		Fproduction_id: uint16(14),
	})))),
	738: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(128)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	741: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	742: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	743: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_id_selector),
		Fproduction_id: uint16(8),
	})))),
	744: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(65)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	749: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	753: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	754: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	755: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	756: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	757: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	758: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	759: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	760: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	761: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	762: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	763: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	764: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	765: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	766: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	767: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	768: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	769: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(175)),
	}})))),
	770: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	771: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	772: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	773: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	774: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	775: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_url_image_scale),
	})))),
	776: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	777: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	778: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	779: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	780: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	781: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	782: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	783: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	784: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	785: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	786: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	787: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	788: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	789: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	790: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	791: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	792: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	793: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	794: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	795: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	796: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	797: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	798: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(137)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(177)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(101)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	807: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(223)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_feature_name),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	815: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	816: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	817: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	818: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	819: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_angle_unit),
	})))),
	820: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	821: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_angle),
		Fproduction_id: uint16(16),
	})))),
	822: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	823: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_gradient_image_dir),
	})))),
	824: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	825: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	826: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	827: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	828: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	829: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	830: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	831: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	832: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	833: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	834: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	835: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	836: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	837: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(147)),
	}})))),
	838: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	839: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	840: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	841: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	842: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	843: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	844: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	845: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	846: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	847: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(167)),
	}})))),
	848: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	849: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	850: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	851: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	852: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	853: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	854: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	855: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	856: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	857: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

func tree_sitter_rasi(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(LANGUAGE_VERSION),
	Fsymbol_count:              uint32(SYMBOL_COUNT),
	Falias_count:               uint32(ALIAS_COUNT),
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
	Fkeyword_capture_token:     uint16(sym_identifier),
	Fprimary_state_ids:         uintptr(unsafe.Pointer(&ts_primary_state_ids)),
	Fname:                      __ccgo_ts + 1617,
	Fmetadata: TSLanguageMetadata{
		Fminor_version: uint8(1),
	},
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

var __ccgo_ts1 = "end\x00identifier\x00@import\x00@theme\x00@media\x00,\x00{\x00}\x00*\x00#\x00.\x00normal\x00selected\x00alternate\x00urgent\x00active\x00:\x00;\x00(\x00)\x00min-width\x00max-width\x00min-height\x00max-height\x00min-aspect-ratio\x00max-aspect-ratio\x00monitor-id\x00enabled\x00inherit\x00DMENU\x00\"\x00string_value_token1\x00integer_value\x00float_value\x00true\x00false\x00url\x00none\x00both\x00width\x00height\x00linear-gradient\x00to\x00top\x00left\x00right\x00bottom\x00deg\x00rad\x00grad\x00turn\x00hex_color_token1\x000\x00%\x00rgb\x00rgba\x00hsl\x00hsla\x00hwb\x00hwba\x00cmyk\x00/\x00bold\x00italic\x00underline\x00strikethrough\x00dash\x00solid\x00calc\x00+\x00-\x00modulo\x00min\x00max\x00floor\x00ceil\x00round\x00px\x00mm\x00cm\x00ph\x00em\x00center\x00north\x00east\x00south\x00west\x00@\x00var\x00horizontal\x00vertical\x00default\x00pointer\x00text\x00[\x00]\x00$\x00env\x00comment\x00stylesheet\x00import_statement\x00theme_statement\x00media_statement\x00rule_set\x00selectors\x00block\x00_block_item\x00_selector\x00global_selector\x00id_selector\x00id_selector_view\x00id_selector_state\x00declaration\x00_query\x00feature_query\x00parenthesized_query\x00feature_name\x00_value\x00string_value\x00boolean_value\x00image_value\x00url_image\x00url_image_scale\x00gradient_image\x00direction\x00angle\x00angle_unit\x00_color_value\x00hex_color\x00percentage\x00rgb_color\x00hsl_color\x00hwb_color\x00cmyk_color\x00named_color\x00text_style_value\x00line_style_value\x00distance_value\x00distance_calc\x00distance_op\x00integer_distance_unit\x00float_distance_unit\x00padding_value\x00border_value\x00first_border_style\x00border_style\x00position_value\x00reference_value\x00orientation_value\x00cursor_value\x00list_value\x00environ_value\x00stylesheet_repeat1\x00media_statement_repeat1\x00selectors_repeat1\x00block_repeat1\x00declaration_repeat1\x00feature_query_repeat1\x00gradient_image_repeat1\x00distance_calc_repeat1\x00position_value_repeat1\x00list_value_repeat1\x00property_name\x00body\x00conditions\x00file\x00filename\x00key\x00name\x00op\x00opacity\x00scale\x00size\x00state\x00style\x00unit\x00value\x00view\x00widget\x00rasi\x00"
