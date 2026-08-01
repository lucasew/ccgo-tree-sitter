// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-firrtl/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-firrtl -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_firrtl

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
const EXTERNAL_TOKEN_COUNT = 3
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
const LANGUAGE_VERSION = 14
const LARGE_STATE_COUNT = 95
const MAX_ALIAS_SEQUENCE_LENGTH = 11
const PRODUCTION_ID_COUNT = 3
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 443
const SYMBOL_COUNT = 194
const TOKEN_COUNT = 129
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
const static_assert = "_Static_assert"
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

const NEWLINE = 0
const INDENT = 1
const DEDENT = 2
const ERROR = 3

type indent_vec = struct {
	Flen1 uint32_t
	Fcap1 uint32_t
	Fdata uintptr
}

type Scanner = struct {
	Findents uintptr
}

func skip(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
}

func tree_sitter_firrtl_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var current_indent_length uint16_t
	var found_end_of_line uint8
	var indent_length, v5 uint32_t
	var scanner, tmp, v6 uintptr
	var v2 uint32
	var v3 bool
	_, _, _, _, _, _, _, _, _ = current_indent_length, found_end_of_line, indent_length, scanner, tmp, v2, v3, v5, v6
	scanner = payload
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
	found_end_of_line = libc.BoolUint8(false1 != 0)
	indent_length = uint32(0)
	for {
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') {
			found_end_of_line = libc.BoolUint8(true1 != 0)
			indent_length = uint32(0)
			skip(tls, lexer)
		} else {
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(' ') {
				indent_length = indent_length + 1
				skip(tls, lexer)
			} else {
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\r') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\f') {
					indent_length = uint32(0)
					skip(tls, lexer)
				} else {
					if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\t') {
						indent_length = indent_length + uint32(8)
						skip(tls, lexer)
					} else {
						if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('#') {
							for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\n') {
								skip(tls, lexer)
							}
							skip(tls, lexer)
							indent_length = uint32(0)
						} else {
							if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\\') {
								skip(tls, lexer)
								if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\r') {
									skip(tls, lexer)
								}
								if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') || (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
									skip(tls, lexer)
								} else {
									return libc.BoolUint8(false1 != 0)
								}
							} else {
								if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
									indent_length = uint32(0)
									found_end_of_line = libc.BoolUint8(true1 != 0)
									break
								} else {
									break
								}
							}
						}
					}
				}
			}
		}
		goto _1
	_1:
	}
	if found_end_of_line != 0 {
		if (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 > uint32(0) {
			current_indent_length = *(*uint16_t)(unsafe.Pointer((*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata + uintptr((*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1-uint32(1))*2))
			if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(INDENT))) != 0 && indent_length > uint32(current_indent_length) {
				if (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fcap1 == (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 {
					if libc.Uint32FromInt32(libc.Int32FromInt32(16)) > (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1*uint32(2) {
						v2 = libc.Uint32FromInt32(libc.Int32FromInt32(16))
					} else {
						v2 = (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 * uint32(2)
					}
					tmp = libc.Xrealloc(tls, (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata, uint64(v2)*uint64(2))
					if v3 = tmp != libc.UintptrFromInt32(0); !v3 {
						libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+12, int32(108), uintptr(unsafe.Pointer(&__func__)))
					}
					_ = v3 || libc.Bool(libc.Int32FromInt32(0) != 0)
					(*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata = tmp
					if libc.Uint32FromInt32(libc.Int32FromInt32(16)) > (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1*uint32(2) {
						v2 = libc.Uint32FromInt32(libc.Int32FromInt32(16))
					} else {
						v2 = (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 * uint32(2)
					}
					(*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fcap1 = v2
				}
				v6 = (*Scanner)(unsafe.Pointer(scanner)).Findents
				v5 = *(*uint32_t)(unsafe.Pointer(v6))
				*(*uint32_t)(unsafe.Pointer(v6)) = *(*uint32_t)(unsafe.Pointer(v6)) + 1
				*(*uint16_t)(unsafe.Pointer((*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata + uintptr(v5)*2)) = uint16(indent_length)
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(INDENT)
				return libc.BoolUint8(true1 != 0)
			}
			if (*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(DEDENT))) != 0 || !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(NEWLINE))) != 0)) && indent_length < uint32(current_indent_length) {
				(*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 = (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 - 1
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(DEDENT)
				return libc.BoolUint8(true1 != 0)
			}
		}
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(NEWLINE))) != 0 {
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(NEWLINE)
			return libc.BoolUint8(true1 != 0)
		}
	}
	return libc.BoolUint8(false1 != 0)
}

var __func__ = [41]uint8{'t', 'r', 'e', 'e', '_', 's', 'i', 't', 't', 'e', 'r', '_', 'f', 'i', 'r', 'r', 't', 'l', '_', 'e', 'x', 't', 'e', 'r', 'n', 'a', 'l', '_', 's', 'c', 'a', 'n', 'n', 'e', 'r', '_', 's', 'c', 'a', 'n'}

func tree_sitter_firrtl_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	var iter int32
	var scanner uintptr
	var size, v2 size_t
	_, _, _, _ = iter, scanner, size, v2
	scanner = payload
	size = uint64(0)
	iter = int32(1)
	for {
		if !(libc.Uint32FromInt32(iter) < (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 && size < uint64(TREE_SITTER_SERIALIZATION_BUFFER_SIZE)) {
			break
		}
		v2 = size
		size = size + 1
		*(*uint8)(unsafe.Pointer(buffer + uintptr(v2))) = uint8(*(*uint16_t)(unsafe.Pointer((*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata + uintptr(iter)*2)))
		goto _1
	_1:
		;
		iter = iter + 1
	}
	return uint32(size)
}

func tree_sitter_firrtl_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
	var scanner, tmp, tmp1, v5 uintptr
	var size size_t
	var v1 uint32
	var v4 uint32_t
	var v2 bool
	_, _, _, _, _, _, _, _ = scanner, size, tmp, tmp1, v1, v2, v4, v5
	scanner = payload
	(*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 = uint32(0)
	if (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fcap1 == (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 {
		if libc.Uint32FromInt32(libc.Int32FromInt32(16)) > (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1*uint32(2) {
			v1 = libc.Uint32FromInt32(libc.Int32FromInt32(16))
		} else {
			v1 = (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 * uint32(2)
		}
		tmp = libc.Xrealloc(tls, (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata, uint64(v1)*uint64(2))
		if v2 = tmp != libc.UintptrFromInt32(0); !v2 {
			libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+12, int32(149), uintptr(unsafe.Pointer(&__func__1)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		(*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata = tmp
		if libc.Uint32FromInt32(libc.Int32FromInt32(16)) > (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1*uint32(2) {
			v1 = libc.Uint32FromInt32(libc.Int32FromInt32(16))
		} else {
			v1 = (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 * uint32(2)
		}
		(*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fcap1 = v1
	}
	v5 = (*Scanner)(unsafe.Pointer(scanner)).Findents
	v4 = *(*uint32_t)(unsafe.Pointer(v5))
	*(*uint32_t)(unsafe.Pointer(v5)) = *(*uint32_t)(unsafe.Pointer(v5)) + 1
	*(*uint16_t)(unsafe.Pointer((*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata + uintptr(v4)*2)) = libc.Uint16FromInt32(libc.Int32FromInt32(0))
	if length > uint32(0) {
		size = uint64(0)
		for {
			if !(size < uint64(length)) {
				break
			}
			if (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fcap1 == (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 {
				if libc.Uint32FromInt32(libc.Int32FromInt32(16)) > (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1*uint32(2) {
					v1 = libc.Uint32FromInt32(libc.Int32FromInt32(16))
				} else {
					v1 = (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 * uint32(2)
				}
				tmp1 = libc.Xrealloc(tls, (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata, uint64(v1)*uint64(2))
				if v2 = tmp1 != libc.UintptrFromInt32(0); !v2 {
					libc.X__assert_fail(tls, __ccgo_ts, __ccgo_ts+12, int32(154), uintptr(unsafe.Pointer(&__func__1)))
				}
				_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
				(*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata = tmp1
				if libc.Uint32FromInt32(libc.Int32FromInt32(16)) > (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1*uint32(2) {
					v1 = libc.Uint32FromInt32(libc.Int32FromInt32(16))
				} else {
					v1 = (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Flen1 * uint32(2)
				}
				(*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fcap1 = v1
			}
			v5 = (*Scanner)(unsafe.Pointer(scanner)).Findents
			v4 = *(*uint32_t)(unsafe.Pointer(v5))
			*(*uint32_t)(unsafe.Pointer(v5)) = *(*uint32_t)(unsafe.Pointer(v5)) + 1
			*(*uint16_t)(unsafe.Pointer((*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata + uintptr(v4)*2)) = uint16(*(*uint8)(unsafe.Pointer(buffer + uintptr(size))))
			goto _6
		_6:
			;
			size = size + 1
		}
		if v2 = size == uint64(length); !v2 {
			libc.X__assert_fail(tls, __ccgo_ts+23, __ccgo_ts+12, int32(156), uintptr(unsafe.Pointer(&__func__1)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
	}
}

var __func__1 = [48]uint8{'t', 'r', 'e', 'e', '_', 's', 'i', 't', 't', 'e', 'r', '_', 'f', 'i', 'r', 'r', 't', 'l', '_', 'e', 'x', 't', 'e', 'r', 'n', 'a', 'l', '_', 's', 'c', 'a', 'n', 'n', 'e', 'r', '_', 'd', 'e', 's', 'e', 'r', 'i', 'a', 'l', 'i', 'z', 'e'}

func tree_sitter_firrtl_external_scanner_create(tls *libc.TLS) (r uintptr) {
	var scanner uintptr
	_ = scanner
	scanner = libc.Xcalloc(tls, uint64(1), uint64(8))
	(*Scanner)(unsafe.Pointer(scanner)).Findents = libc.Xcalloc(tls, uint64(1), uint64(16))
	tree_sitter_firrtl_external_scanner_deserialize(tls, scanner, libc.UintptrFromInt32(0), uint32(0))
	return scanner
}

func tree_sitter_firrtl_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	if (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata != libc.UintptrFromInt32(0) {
		libc.Xfree(tls, (*indent_vec)(unsafe.Pointer((*Scanner)(unsafe.Pointer(scanner)).Findents)).Fdata)
	}
	libc.Xfree(tls, (*Scanner)(unsafe.Pointer(scanner)).Findents)
	libc.Xfree(tls, scanner)
}

type ts_symbol_identifiers = int32

const sym_identifier = 1
const anon_sym_circuit = 2
const anon_sym_COLON = 3
const anon_sym_module = 4
const anon_sym_extmodule = 5
const anon_sym_input = 6
const anon_sym_output = 7
const anon_sym_const = 8
const anon_sym_UInt = 9
const anon_sym_SInt = 10
const anon_sym_Analog = 11
const anon_sym_LT = 12
const anon_sym_GT = 13
const anon_sym_Fixed = 14
const anon_sym_Clock = 15
const anon_sym_AsyncReset = 16
const anon_sym_Reset = 17
const anon_sym_LBRACE = 18
const anon_sym_COMMA = 19
const anon_sym_RBRACE = 20
const anon_sym_LBRACK = 21
const anon_sym_RBRACK = 22
const anon_sym_flip = 23
const anon_sym_defname = 24
const anon_sym_EQ = 25
const anon_sym_parameter = 26
const anon_sym_reset = 27
const anon_sym_EQ_GT = 28
const anon_sym_LPAREN = 29
const anon_sym_RPAREN = 30
const anon_sym_wire = 31
const anon_sym_cmem = 32
const anon_sym_smem = 33
const anon_sym_reg = 34
const anon_sym_with = 35
const anon_sym_mem = 36
const anon_sym_mport = 37
const anon_sym_inst = 38
const anon_sym_of = 39
const anon_sym_node = 40
const anon_sym_LT_EQ = 41
const anon_sym_LT_DASH = 42
const anon_sym_is = 43
const anon_sym_invalid = 44
const anon_sym_data_DASHtype = 45
const anon_sym_depth = 46
const anon_sym_read_DASHlatency = 47
const anon_sym_write_DASHlatency = 48
const anon_sym_read_DASHunder_DASHwrite = 49
const anon_sym_reader = 50
const anon_sym_writer = 51
const anon_sym_readwriter = 52
const anon_sym_when = 53
const anon_sym_else = 54
const anon_sym_stop = 55
const anon_sym_printf = 56
const anon_sym_assert = 57
const anon_sym_assume = 58
const anon_sym_cover = 59
const anon_sym_skip = 60
const anon_sym_attach = 61
const sym_info = 62
const anon_sym_infer = 63
const anon_sym_read = 64
const anon_sym_write = 65
const anon_sym_rdwr = 66
const anon_sym_old = 67
const anon_sym_new = 68
const anon_sym_undefined = 69
const anon_sym_DOT = 70
const anon_sym_mux = 71
const anon_sym_validif = 72
const anon_sym_add = 73
const anon_sym_sub = 74
const anon_sym_mul = 75
const anon_sym_div = 76
const anon_sym_rem = 77
const anon_sym_lt = 78
const anon_sym_leq = 79
const anon_sym_gt = 80
const anon_sym_geq = 81
const anon_sym_eq = 82
const anon_sym_neq = 83
const anon_sym_pad = 84
const anon_sym_asUInt = 85
const anon_sym_asAsyncReset = 86
const anon_sym_asSInt = 87
const anon_sym_asClock = 88
const anon_sym_shl = 89
const anon_sym_shr = 90
const anon_sym_dshl = 91
const anon_sym_dshlw = 92
const anon_sym_dshr = 93
const anon_sym_dshrw = 94
const anon_sym_cvt = 95
const anon_sym_neg = 96
const anon_sym_not = 97
const anon_sym_and = 98
const anon_sym_or = 99
const anon_sym_xor = 100
const anon_sym_andr = 101
const anon_sym_orr = 102
const anon_sym_xorr = 103
const anon_sym_cat = 104
const anon_sym_bits = 105
const anon_sym_head = 106
const anon_sym_tail = 107
const anon_sym_asFixedPoint = 108
const anon_sym_bpshl = 109
const anon_sym_bpshr = 110
const anon_sym_bpset = 111
const anon_sym_0 = 112
const aux_sym_uint_token1 = 113
const anon_sym_PLUS = 114
const anon_sym_DASH = 115
const sym_number_str = 116
const sym_double = 117
const anon_sym_DQUOTE = 118
const anon_sym_SQUOTE = 119
const sym_string_content = 120
const sym_raw_string_content = 121
const aux_sym__escape_sequence_token1 = 122
const sym_escape_sequence = 123
const sym_relaxed_identifier = 124
const sym_comment = 125
const sym__newline = 126
const sym__indent = 127
const sym__dedent = 128
const sym_source_file = 129
const sym_circuit = 130
const sym_module = 131
const sym_port = 132
const sym_dir = 133
const sym_qualifier = 134
const sym_type = 135
const sym_field = 136
const sym_defname = 137
const sym_parameter = 138
const sym__reset = 139
const sym_reset = 140
const sym_reset_block = 141
const sym_statement = 142
const sym_wire = 143
const sym_cmem = 144
const sym_smem = 145
const sym_register = 146
const sym_memory = 147
const sym_rdwr = 148
const sym_inst = 149
const sym_node = 150
const sym_connection = 151
const sym_partial_connection = 152
const sym_is_invalid = 153
const sym_memory_field = 154
const sym_suite = 155
const sym_when = 156
const sym_else = 157
const sym_stop = 158
const sym_printf = 159
const sym_verif = 160
const sym_skip = 161
const sym_attach = 162
const sym_mdir = 163
const sym_ruw = 164
const sym_litType = 165
const sym_expression = 166
const sym_literal = 167
const sym_sub_field = 168
const sym_sub_index = 169
const sym_sub_access = 170
const sym_mux = 171
const sym_conditionally_valid = 172
const sym_primitive_operation = 173
const sym_field_id = 174
const sym_primop = 175
const sym_uint = 176
const sym_sint = 177
const sym_number = 178
const sym_string = 179
const sym_raw_string = 180
const sym__escape_sequence = 181
const aux_sym_source_file_repeat1 = 182
const aux_sym_circuit_repeat1 = 183
const aux_sym_module_repeat1 = 184
const aux_sym_module_repeat2 = 185
const aux_sym_module_repeat3 = 186
const aux_sym_type_repeat1 = 187
const aux_sym_memory_repeat1 = 188
const aux_sym_memory_field_repeat1 = 189
const aux_sym_printf_repeat1 = 190
const aux_sym_primitive_operation_repeat1 = 191
const aux_sym_string_repeat1 = 192
const aux_sym_raw_string_repeat1 = 193

var ts_symbol_names = [194]uintptr{
	0:   __ccgo_ts + 38,
	1:   __ccgo_ts + 42,
	2:   __ccgo_ts + 53,
	3:   __ccgo_ts + 61,
	4:   __ccgo_ts + 63,
	5:   __ccgo_ts + 70,
	6:   __ccgo_ts + 80,
	7:   __ccgo_ts + 86,
	8:   __ccgo_ts + 93,
	9:   __ccgo_ts + 99,
	10:  __ccgo_ts + 104,
	11:  __ccgo_ts + 109,
	12:  __ccgo_ts + 116,
	13:  __ccgo_ts + 118,
	14:  __ccgo_ts + 120,
	15:  __ccgo_ts + 126,
	16:  __ccgo_ts + 132,
	17:  __ccgo_ts + 143,
	18:  __ccgo_ts + 149,
	19:  __ccgo_ts + 151,
	20:  __ccgo_ts + 153,
	21:  __ccgo_ts + 155,
	22:  __ccgo_ts + 157,
	23:  __ccgo_ts + 159,
	24:  __ccgo_ts + 164,
	25:  __ccgo_ts + 172,
	26:  __ccgo_ts + 174,
	27:  __ccgo_ts + 184,
	28:  __ccgo_ts + 190,
	29:  __ccgo_ts + 193,
	30:  __ccgo_ts + 195,
	31:  __ccgo_ts + 197,
	32:  __ccgo_ts + 202,
	33:  __ccgo_ts + 207,
	34:  __ccgo_ts + 212,
	35:  __ccgo_ts + 216,
	36:  __ccgo_ts + 221,
	37:  __ccgo_ts + 225,
	38:  __ccgo_ts + 231,
	39:  __ccgo_ts + 236,
	40:  __ccgo_ts + 239,
	41:  __ccgo_ts + 244,
	42:  __ccgo_ts + 247,
	43:  __ccgo_ts + 250,
	44:  __ccgo_ts + 253,
	45:  __ccgo_ts + 261,
	46:  __ccgo_ts + 271,
	47:  __ccgo_ts + 277,
	48:  __ccgo_ts + 290,
	49:  __ccgo_ts + 304,
	50:  __ccgo_ts + 321,
	51:  __ccgo_ts + 328,
	52:  __ccgo_ts + 335,
	53:  __ccgo_ts + 346,
	54:  __ccgo_ts + 351,
	55:  __ccgo_ts + 356,
	56:  __ccgo_ts + 361,
	57:  __ccgo_ts + 368,
	58:  __ccgo_ts + 375,
	59:  __ccgo_ts + 382,
	60:  __ccgo_ts + 388,
	61:  __ccgo_ts + 393,
	62:  __ccgo_ts + 400,
	63:  __ccgo_ts + 405,
	64:  __ccgo_ts + 411,
	65:  __ccgo_ts + 416,
	66:  __ccgo_ts + 422,
	67:  __ccgo_ts + 427,
	68:  __ccgo_ts + 431,
	69:  __ccgo_ts + 435,
	70:  __ccgo_ts + 445,
	71:  __ccgo_ts + 447,
	72:  __ccgo_ts + 451,
	73:  __ccgo_ts + 459,
	74:  __ccgo_ts + 463,
	75:  __ccgo_ts + 467,
	76:  __ccgo_ts + 471,
	77:  __ccgo_ts + 475,
	78:  __ccgo_ts + 479,
	79:  __ccgo_ts + 482,
	80:  __ccgo_ts + 486,
	81:  __ccgo_ts + 489,
	82:  __ccgo_ts + 493,
	83:  __ccgo_ts + 496,
	84:  __ccgo_ts + 500,
	85:  __ccgo_ts + 504,
	86:  __ccgo_ts + 511,
	87:  __ccgo_ts + 524,
	88:  __ccgo_ts + 531,
	89:  __ccgo_ts + 539,
	90:  __ccgo_ts + 543,
	91:  __ccgo_ts + 547,
	92:  __ccgo_ts + 552,
	93:  __ccgo_ts + 558,
	94:  __ccgo_ts + 563,
	95:  __ccgo_ts + 569,
	96:  __ccgo_ts + 573,
	97:  __ccgo_ts + 577,
	98:  __ccgo_ts + 581,
	99:  __ccgo_ts + 585,
	100: __ccgo_ts + 588,
	101: __ccgo_ts + 592,
	102: __ccgo_ts + 597,
	103: __ccgo_ts + 601,
	104: __ccgo_ts + 606,
	105: __ccgo_ts + 610,
	106: __ccgo_ts + 615,
	107: __ccgo_ts + 620,
	108: __ccgo_ts + 625,
	109: __ccgo_ts + 638,
	110: __ccgo_ts + 644,
	111: __ccgo_ts + 650,
	112: __ccgo_ts + 656,
	113: __ccgo_ts + 658,
	114: __ccgo_ts + 670,
	115: __ccgo_ts + 672,
	116: __ccgo_ts + 674,
	117: __ccgo_ts + 685,
	118: __ccgo_ts + 692,
	119: __ccgo_ts + 694,
	120: __ccgo_ts + 696,
	121: __ccgo_ts + 696,
	122: __ccgo_ts + 711,
	123: __ccgo_ts + 735,
	124: __ccgo_ts + 751,
	125: __ccgo_ts + 770,
	126: __ccgo_ts + 778,
	127: __ccgo_ts + 787,
	128: __ccgo_ts + 795,
	129: __ccgo_ts + 803,
	130: __ccgo_ts + 53,
	131: __ccgo_ts + 63,
	132: __ccgo_ts + 815,
	133: __ccgo_ts + 820,
	134: __ccgo_ts + 824,
	135: __ccgo_ts + 834,
	136: __ccgo_ts + 839,
	137: __ccgo_ts + 164,
	138: __ccgo_ts + 174,
	139: __ccgo_ts + 845,
	140: __ccgo_ts + 184,
	141: __ccgo_ts + 852,
	142: __ccgo_ts + 864,
	143: __ccgo_ts + 197,
	144: __ccgo_ts + 202,
	145: __ccgo_ts + 207,
	146: __ccgo_ts + 874,
	147: __ccgo_ts + 883,
	148: __ccgo_ts + 422,
	149: __ccgo_ts + 231,
	150: __ccgo_ts + 239,
	151: __ccgo_ts + 890,
	152: __ccgo_ts + 901,
	153: __ccgo_ts + 920,
	154: __ccgo_ts + 931,
	155: __ccgo_ts + 944,
	156: __ccgo_ts + 346,
	157: __ccgo_ts + 351,
	158: __ccgo_ts + 356,
	159: __ccgo_ts + 361,
	160: __ccgo_ts + 950,
	161: __ccgo_ts + 388,
	162: __ccgo_ts + 393,
	163: __ccgo_ts + 956,
	164: __ccgo_ts + 961,
	165: __ccgo_ts + 965,
	166: __ccgo_ts + 973,
	167: __ccgo_ts + 984,
	168: __ccgo_ts + 992,
	169: __ccgo_ts + 1002,
	170: __ccgo_ts + 1012,
	171: __ccgo_ts + 447,
	172: __ccgo_ts + 1023,
	173: __ccgo_ts + 1043,
	174: __ccgo_ts + 1063,
	175: __ccgo_ts + 1072,
	176: __ccgo_ts + 1079,
	177: __ccgo_ts + 1084,
	178: __ccgo_ts + 1084,
	179: __ccgo_ts + 1091,
	180: __ccgo_ts + 1098,
	181: __ccgo_ts + 1109,
	182: __ccgo_ts + 1126,
	183: __ccgo_ts + 1146,
	184: __ccgo_ts + 1162,
	185: __ccgo_ts + 1177,
	186: __ccgo_ts + 1192,
	187: __ccgo_ts + 1207,
	188: __ccgo_ts + 1220,
	189: __ccgo_ts + 1235,
	190: __ccgo_ts + 1256,
	191: __ccgo_ts + 1271,
	192: __ccgo_ts + 1299,
	193: __ccgo_ts + 1314,
}

var ts_symbol_map = [194]TSSymbol{
	1:   uint16(sym_identifier),
	2:   uint16(anon_sym_circuit),
	3:   uint16(anon_sym_COLON),
	4:   uint16(anon_sym_module),
	5:   uint16(anon_sym_extmodule),
	6:   uint16(anon_sym_input),
	7:   uint16(anon_sym_output),
	8:   uint16(anon_sym_const),
	9:   uint16(anon_sym_UInt),
	10:  uint16(anon_sym_SInt),
	11:  uint16(anon_sym_Analog),
	12:  uint16(anon_sym_LT),
	13:  uint16(anon_sym_GT),
	14:  uint16(anon_sym_Fixed),
	15:  uint16(anon_sym_Clock),
	16:  uint16(anon_sym_AsyncReset),
	17:  uint16(anon_sym_Reset),
	18:  uint16(anon_sym_LBRACE),
	19:  uint16(anon_sym_COMMA),
	20:  uint16(anon_sym_RBRACE),
	21:  uint16(anon_sym_LBRACK),
	22:  uint16(anon_sym_RBRACK),
	23:  uint16(anon_sym_flip),
	24:  uint16(anon_sym_defname),
	25:  uint16(anon_sym_EQ),
	26:  uint16(anon_sym_parameter),
	27:  uint16(anon_sym_reset),
	28:  uint16(anon_sym_EQ_GT),
	29:  uint16(anon_sym_LPAREN),
	30:  uint16(anon_sym_RPAREN),
	31:  uint16(anon_sym_wire),
	32:  uint16(anon_sym_cmem),
	33:  uint16(anon_sym_smem),
	34:  uint16(anon_sym_reg),
	35:  uint16(anon_sym_with),
	36:  uint16(anon_sym_mem),
	37:  uint16(anon_sym_mport),
	38:  uint16(anon_sym_inst),
	39:  uint16(anon_sym_of),
	40:  uint16(anon_sym_node),
	41:  uint16(anon_sym_LT_EQ),
	42:  uint16(anon_sym_LT_DASH),
	43:  uint16(anon_sym_is),
	44:  uint16(anon_sym_invalid),
	45:  uint16(anon_sym_data_DASHtype),
	46:  uint16(anon_sym_depth),
	47:  uint16(anon_sym_read_DASHlatency),
	48:  uint16(anon_sym_write_DASHlatency),
	49:  uint16(anon_sym_read_DASHunder_DASHwrite),
	50:  uint16(anon_sym_reader),
	51:  uint16(anon_sym_writer),
	52:  uint16(anon_sym_readwriter),
	53:  uint16(anon_sym_when),
	54:  uint16(anon_sym_else),
	55:  uint16(anon_sym_stop),
	56:  uint16(anon_sym_printf),
	57:  uint16(anon_sym_assert),
	58:  uint16(anon_sym_assume),
	59:  uint16(anon_sym_cover),
	60:  uint16(anon_sym_skip),
	61:  uint16(anon_sym_attach),
	62:  uint16(sym_info),
	63:  uint16(anon_sym_infer),
	64:  uint16(anon_sym_read),
	65:  uint16(anon_sym_write),
	66:  uint16(anon_sym_rdwr),
	67:  uint16(anon_sym_old),
	68:  uint16(anon_sym_new),
	69:  uint16(anon_sym_undefined),
	70:  uint16(anon_sym_DOT),
	71:  uint16(anon_sym_mux),
	72:  uint16(anon_sym_validif),
	73:  uint16(anon_sym_add),
	74:  uint16(anon_sym_sub),
	75:  uint16(anon_sym_mul),
	76:  uint16(anon_sym_div),
	77:  uint16(anon_sym_rem),
	78:  uint16(anon_sym_lt),
	79:  uint16(anon_sym_leq),
	80:  uint16(anon_sym_gt),
	81:  uint16(anon_sym_geq),
	82:  uint16(anon_sym_eq),
	83:  uint16(anon_sym_neq),
	84:  uint16(anon_sym_pad),
	85:  uint16(anon_sym_asUInt),
	86:  uint16(anon_sym_asAsyncReset),
	87:  uint16(anon_sym_asSInt),
	88:  uint16(anon_sym_asClock),
	89:  uint16(anon_sym_shl),
	90:  uint16(anon_sym_shr),
	91:  uint16(anon_sym_dshl),
	92:  uint16(anon_sym_dshlw),
	93:  uint16(anon_sym_dshr),
	94:  uint16(anon_sym_dshrw),
	95:  uint16(anon_sym_cvt),
	96:  uint16(anon_sym_neg),
	97:  uint16(anon_sym_not),
	98:  uint16(anon_sym_and),
	99:  uint16(anon_sym_or),
	100: uint16(anon_sym_xor),
	101: uint16(anon_sym_andr),
	102: uint16(anon_sym_orr),
	103: uint16(anon_sym_xorr),
	104: uint16(anon_sym_cat),
	105: uint16(anon_sym_bits),
	106: uint16(anon_sym_head),
	107: uint16(anon_sym_tail),
	108: uint16(anon_sym_asFixedPoint),
	109: uint16(anon_sym_bpshl),
	110: uint16(anon_sym_bpshr),
	111: uint16(anon_sym_bpset),
	112: uint16(anon_sym_0),
	113: uint16(aux_sym_uint_token1),
	114: uint16(anon_sym_PLUS),
	115: uint16(anon_sym_DASH),
	116: uint16(sym_number_str),
	117: uint16(sym_double),
	118: uint16(anon_sym_DQUOTE),
	119: uint16(anon_sym_SQUOTE),
	120: uint16(sym_string_content),
	121: uint16(sym_string_content),
	122: uint16(aux_sym__escape_sequence_token1),
	123: uint16(sym_escape_sequence),
	124: uint16(sym_relaxed_identifier),
	125: uint16(sym_comment),
	126: uint16(sym__newline),
	127: uint16(sym__indent),
	128: uint16(sym__dedent),
	129: uint16(sym_source_file),
	130: uint16(sym_circuit),
	131: uint16(sym_module),
	132: uint16(sym_port),
	133: uint16(sym_dir),
	134: uint16(sym_qualifier),
	135: uint16(sym_type),
	136: uint16(sym_field),
	137: uint16(sym_defname),
	138: uint16(sym_parameter),
	139: uint16(sym__reset),
	140: uint16(sym_reset),
	141: uint16(sym_reset_block),
	142: uint16(sym_statement),
	143: uint16(sym_wire),
	144: uint16(sym_cmem),
	145: uint16(sym_smem),
	146: uint16(sym_register),
	147: uint16(sym_memory),
	148: uint16(sym_rdwr),
	149: uint16(sym_inst),
	150: uint16(sym_node),
	151: uint16(sym_connection),
	152: uint16(sym_partial_connection),
	153: uint16(sym_is_invalid),
	154: uint16(sym_memory_field),
	155: uint16(sym_suite),
	156: uint16(sym_when),
	157: uint16(sym_else),
	158: uint16(sym_stop),
	159: uint16(sym_printf),
	160: uint16(sym_verif),
	161: uint16(sym_skip),
	162: uint16(sym_attach),
	163: uint16(sym_mdir),
	164: uint16(sym_ruw),
	165: uint16(sym_litType),
	166: uint16(sym_expression),
	167: uint16(sym_literal),
	168: uint16(sym_sub_field),
	169: uint16(sym_sub_index),
	170: uint16(sym_sub_access),
	171: uint16(sym_mux),
	172: uint16(sym_conditionally_valid),
	173: uint16(sym_primitive_operation),
	174: uint16(sym_field_id),
	175: uint16(sym_primop),
	176: uint16(sym_uint),
	177: uint16(sym_number),
	178: uint16(sym_number),
	179: uint16(sym_string),
	180: uint16(sym_raw_string),
	181: uint16(sym__escape_sequence),
	182: uint16(aux_sym_source_file_repeat1),
	183: uint16(aux_sym_circuit_repeat1),
	184: uint16(aux_sym_module_repeat1),
	185: uint16(aux_sym_module_repeat2),
	186: uint16(aux_sym_module_repeat3),
	187: uint16(aux_sym_type_repeat1),
	188: uint16(aux_sym_memory_repeat1),
	189: uint16(aux_sym_memory_field_repeat1),
	190: uint16(aux_sym_printf_repeat1),
	191: uint16(aux_sym_primitive_operation_repeat1),
	192: uint16(aux_sym_string_repeat1),
	193: uint16(aux_sym_raw_string_repeat1),
}

var ts_symbol_metadata = [194]TSSymbolMetadata{
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
	},
	60: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	61: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	62: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	},
	100: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	101: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	102: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	103: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	104: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	105: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	106: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	107: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	108: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	109: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	110: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	111: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	112: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	113: {},
	114: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	115: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
	},
	119: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	120: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	121: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	122: {},
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	127: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
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
	153: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	154: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	155: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	156: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	157: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	158: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	159: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	160: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	161: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	162: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	163: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	164: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	165: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	166: {
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
	},
	167: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	168: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	169: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	170: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	171: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	172: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	173: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	174: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	175: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	176: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	177: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	178: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	179: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	180: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	181: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	182: {},
	183: {},
	184: {},
	185: {},
	186: {},
	187: {},
	188: {},
	189: {},
	190: {},
	191: {},
	192: {},
	193: {},
}

var ts_alias_sequences = [3][11]TSSymbol{
	0: {},
	1: {
		0: uint16(sym_identifier),
	},
	2: {
		0: uint16(sym_number),
	},
}

var ts_non_terminal_alias_map = [9]uint16_t{
	0: uint16(sym_primop),
	1: uint16(2),
	2: uint16(sym_primop),
	3: uint16(sym_identifier),
	4: uint16(sym_uint),
	5: uint16(2),
	6: uint16(sym_uint),
	7: uint16(sym_number),
}

var ts_primary_state_ids = [443]TSStateId{
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
	73:  uint16(73),
	74:  uint16(74),
	75:  uint16(75),
	76:  uint16(76),
	77:  uint16(77),
	78:  uint16(78),
	79:  uint16(79),
	80:  uint16(80),
	81:  uint16(81),
	82:  uint16(82),
	83:  uint16(83),
	84:  uint16(84),
	85:  uint16(85),
	86:  uint16(86),
	87:  uint16(87),
	88:  uint16(88),
	89:  uint16(89),
	90:  uint16(90),
	91:  uint16(91),
	92:  uint16(92),
	93:  uint16(93),
	94:  uint16(94),
	95:  uint16(95),
	96:  uint16(95),
	97:  uint16(97),
	98:  uint16(98),
	99:  uint16(97),
	100: uint16(100),
	101: uint16(101),
	102: uint16(102),
	103: uint16(103),
	104: uint16(104),
	105: uint16(105),
	106: uint16(100),
	107: uint16(107),
	108: uint16(108),
	109: uint16(109),
	110: uint16(110),
	111: uint16(105),
	112: uint16(103),
	113: uint16(113),
	114: uint16(114),
	115: uint16(115),
	116: uint16(116),
	117: uint16(117),
	118: uint16(107),
	119: uint16(119),
	120: uint16(120),
	121: uint16(121),
	122: uint16(122),
	123: uint16(123),
	124: uint16(114),
	125: uint16(104),
	126: uint16(126),
	127: uint16(101),
	128: uint16(128),
	129: uint16(129),
	130: uint16(130),
	131: uint16(131),
	132: uint16(132),
	133: uint16(133),
	134: uint16(134),
	135: uint16(135),
	136: uint16(136),
	137: uint16(137),
	138: uint16(138),
	139: uint16(139),
	140: uint16(140),
	141: uint16(18),
	142: uint16(17),
	143: uint16(19),
	144: uint16(23),
	145: uint16(145),
	146: uint16(20),
	147: uint16(26),
	148: uint16(148),
	149: uint16(149),
	150: uint16(150),
	151: uint16(151),
	152: uint16(24),
	153: uint16(151),
	154: uint16(22),
	155: uint16(155),
	156: uint16(156),
	157: uint16(151),
	158: uint16(158),
	159: uint16(159),
	160: uint16(31),
	161: uint16(161),
	162: uint16(29),
	163: uint16(25),
	164: uint16(164),
	165: uint16(21),
	166: uint16(27),
	167: uint16(30),
	168: uint16(168),
	169: uint16(169),
	170: uint16(169),
	171: uint16(171),
	172: uint16(172),
	173: uint16(172),
	174: uint16(172),
	175: uint16(175),
	176: uint16(175),
	177: uint16(177),
	178: uint16(175),
	179: uint16(179),
	180: uint16(177),
	181: uint16(179),
	182: uint16(182),
	183: uint16(177),
	184: uint16(184),
	185: uint16(185),
	186: uint16(182),
	187: uint16(182),
	188: uint16(179),
	189: uint16(189),
	190: uint16(190),
	191: uint16(191),
	192: uint16(192),
	193: uint16(193),
	194: uint16(194),
	195: uint16(195),
	196: uint16(196),
	197: uint16(196),
	198: uint16(198),
	199: uint16(195),
	200: uint16(194),
	201: uint16(201),
	202: uint16(202),
	203: uint16(203),
	204: uint16(204),
	205: uint16(205),
	206: uint16(206),
	207: uint16(207),
	208: uint16(208),
	209: uint16(209),
	210: uint16(204),
	211: uint16(211),
	212: uint16(212),
	213: uint16(34),
	214: uint16(214),
	215: uint16(215),
	216: uint16(216),
	217: uint16(217),
	218: uint16(218),
	219: uint16(40),
	220: uint16(220),
	221: uint16(221),
	222: uint16(222),
	223: uint16(223),
	224: uint16(39),
	225: uint16(225),
	226: uint16(226),
	227: uint16(40),
	228: uint16(228),
	229: uint16(229),
	230: uint16(230),
	231: uint16(34),
	232: uint16(39),
	233: uint16(233),
	234: uint16(234),
	235: uint16(235),
	236: uint16(236),
	237: uint16(237),
	238: uint16(43),
	239: uint16(49),
	240: uint16(240),
	241: uint16(241),
	242: uint16(242),
	243: uint16(243),
	244: uint16(244),
	245: uint16(245),
	246: uint16(244),
	247: uint16(240),
	248: uint16(248),
	249: uint16(249),
	250: uint16(250),
	251: uint16(248),
	252: uint16(252),
	253: uint16(50),
	254: uint16(254),
	255: uint16(255),
	256: uint16(256),
	257: uint16(257),
	258: uint16(258),
	259: uint16(259),
	260: uint16(260),
	261: uint16(261),
	262: uint16(252),
	263: uint16(263),
	264: uint16(48),
	265: uint16(245),
	266: uint16(266),
	267: uint16(267),
	268: uint16(268),
	269: uint16(269),
	270: uint16(270),
	271: uint16(261),
	272: uint16(267),
	273: uint16(47),
	274: uint16(252),
	275: uint16(275),
	276: uint16(276),
	277: uint16(277),
	278: uint16(258),
	279: uint16(254),
	280: uint16(280),
	281: uint16(281),
	282: uint16(266),
	283: uint16(261),
	284: uint16(284),
	285: uint16(42),
	286: uint16(286),
	287: uint16(287),
	288: uint16(288),
	289: uint16(250),
	290: uint16(47),
	291: uint16(291),
	292: uint16(43),
	293: uint16(48),
	294: uint16(294),
	295: uint16(295),
	296: uint16(296),
	297: uint16(50),
	298: uint16(298),
	299: uint16(49),
	300: uint16(300),
	301: uint16(301),
	302: uint16(42),
	303: uint16(275),
	304: uint16(304),
	305: uint16(305),
	306: uint16(306),
	307: uint16(307),
	308: uint16(308),
	309: uint16(309),
	310: uint16(310),
	311: uint16(311),
	312: uint16(312),
	313: uint16(313),
	314: uint16(314),
	315: uint16(315),
	316: uint16(316),
	317: uint16(317),
	318: uint16(313),
	319: uint16(319),
	320: uint16(320),
	321: uint16(321),
	322: uint16(322),
	323: uint16(323),
	324: uint16(324),
	325: uint16(325),
	326: uint16(326),
	327: uint16(327),
	328: uint16(328),
	329: uint16(329),
	330: uint16(330),
	331: uint16(331),
	332: uint16(332),
	333: uint16(333),
	334: uint16(334),
	335: uint16(335),
	336: uint16(336),
	337: uint16(337),
	338: uint16(338),
	339: uint16(339),
	340: uint16(340),
	341: uint16(341),
	342: uint16(23),
	343: uint16(222),
	344: uint16(344),
	345: uint16(345),
	346: uint16(235),
	347: uint16(347),
	348: uint16(348),
	349: uint16(349),
	350: uint16(350),
	351: uint16(351),
	352: uint16(352),
	353: uint16(353),
	354: uint16(354),
	355: uint16(355),
	356: uint16(356),
	357: uint16(357),
	358: uint16(358),
	359: uint16(359),
	360: uint16(360),
	361: uint16(361),
	362: uint16(362),
	363: uint16(363),
	364: uint16(314),
	365: uint16(365),
	366: uint16(324),
	367: uint16(323),
	368: uint16(368),
	369: uint16(369),
	370: uint16(370),
	371: uint16(371),
	372: uint16(372),
	373: uint16(373),
	374: uint16(374),
	375: uint16(375),
	376: uint16(376),
	377: uint16(377),
	378: uint16(378),
	379: uint16(353),
	380: uint16(380),
	381: uint16(381),
	382: uint16(382),
	383: uint16(383),
	384: uint16(384),
	385: uint16(385),
	386: uint16(386),
	387: uint16(387),
	388: uint16(388),
	389: uint16(389),
	390: uint16(390),
	391: uint16(391),
	392: uint16(348),
	393: uint16(393),
	394: uint16(394),
	395: uint16(395),
	396: uint16(377),
	397: uint16(352),
	398: uint16(385),
	399: uint16(399),
	400: uint16(400),
	401: uint16(401),
	402: uint16(391),
	403: uint16(348),
	404: uint16(404),
	405: uint16(394),
	406: uint16(377),
	407: uint16(352),
	408: uint16(408),
	409: uint16(409),
	410: uint16(410),
	411: uint16(411),
	412: uint16(412),
	413: uint16(413),
	414: uint16(309),
	415: uint16(415),
	416: uint16(400),
	417: uint16(358),
	418: uint16(418),
	419: uint16(382),
	420: uint16(420),
	421: uint16(380),
	422: uint16(422),
	423: uint16(400),
	424: uint16(358),
	425: uint16(425),
	426: uint16(426),
	427: uint16(311),
	428: uint16(428),
	429: uint16(429),
	430: uint16(430),
	431: uint16(391),
	432: uint16(432),
	433: uint16(425),
	434: uint16(434),
	435: uint16(376),
	436: uint16(436),
	437: uint16(376),
	438: uint16(426),
	439: uint16(394),
	440: uint16(349),
	441: uint16(441),
	442: uint16(356),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2, i3 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _ = eof, i, i1, i2, i3, lookahead, result, skip
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
			state = uint16(60)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(92)/libc.Uint64FromInt64(2)) {
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
			state = uint16(58)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(86)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('"') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(37)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(98)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('"') {
			state = uint16(93)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('"') {
			state = uint16(93)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('"') {
			state = uint16(93)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(4)
			goto next_state
		}
		return result
	case int32(5):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
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
			state = uint16(5)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(88)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('\'') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(122)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(37)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(100)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('-') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('-') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('.') {
			state = uint16(51)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('>') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('[') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('a') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('a') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('b') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('h') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('c') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('c') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('d') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('e') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('e') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('e') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('e') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('e') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('i') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('l') {
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('l') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('n') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('n') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('n') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('p') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('r') {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('r') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('t') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('t') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('t') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('t') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('u') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(57)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\'') || lookahead == int32('?') || lookahead == int32('\\') || lookahead == int32('a') || lookahead == int32('b') || lookahead == int32('f') || lookahead == int32('n') || lookahead == int32('r') || int32('t') <= lookahead && lookahead <= int32('v') {
			state = uint16(104)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('w') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('y') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('y') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('y') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('{') {
			state = uint16(55)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('}') {
			state = uint16(104)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(44)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(50)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(52)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(56)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(4)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(50):
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(51):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(52):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(53):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(54):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(55):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(56):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(4)
			goto next_state
		}
		return result
	case int32(57):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(58):
		if eof != 0 {
			state = uint16(60)
			goto next_state
		}
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(88)/libc.Uint64FromInt64(2)) {
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
			state = uint16(58)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(86)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(59):
		if eof != 0 {
			state = uint16(60)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
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
			state = uint16(59)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(87)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('>') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_data_DASHtype)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_read_DASHlatency)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_write_DASHlatency)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_read_DASHunder_DASHwrite)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_info)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(51)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_uint_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(51)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_uint_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_uint_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('$') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(44)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(44)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number_str)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_double)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('E') {
			state = uint16(47)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_double)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(121)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(98)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\\') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_string_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_raw_string_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\'') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(122)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(100)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\\') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_raw_string_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(26)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(118)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(112)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(106)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(110)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(108)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(117)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(114)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(111)
			goto next_state
		}
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_relaxed_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') || lookahead == int32('\\') {
			state = uint16(123)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\'') || lookahead == int32('\\') {
			state = uint16(123)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(123)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [46]uint16_t{
	0:  uint16('"'),
	1:  uint16(96),
	2:  uint16('\''),
	3:  uint16(97),
	4:  uint16('('),
	5:  uint16(73),
	6:  uint16(')'),
	7:  uint16(74),
	8:  uint16('+'),
	9:  uint16(90),
	10: uint16(','),
	11: uint16(66),
	12: uint16('-'),
	13: uint16(92),
	14: uint16('.'),
	15: uint16(82),
	16: uint16('0'),
	17: uint16(84),
	18: uint16(':'),
	19: uint16(61),
	20: uint16(';'),
	21: uint16(123),
	22: uint16('<'),
	23: uint16(63),
	24: uint16('='),
	25: uint16(71),
	26: uint16('>'),
	27: uint16(64),
	28: uint16('@'),
	29: uint16(11),
	30: uint16('['),
	31: uint16(68),
	32: uint16('\\'),
	33: uint16(37),
	34: uint16(']'),
	35: uint16(69),
	36: uint16('d'),
	37: uint16(109),
	38: uint16('r'),
	39: uint16(113),
	40: uint16('w'),
	41: uint16(116),
	42: uint16('{'),
	43: uint16(65),
	44: uint16('}'),
	45: uint16(67),
}

var map_token1 = [28]uint16_t{
	0:  uint16('$'),
	1:  uint16(120),
	2:  uint16('('),
	3:  uint16(73),
	4:  uint16(')'),
	5:  uint16(74),
	6:  uint16(','),
	7:  uint16(66),
	8:  uint16('.'),
	9:  uint16(82),
	10: uint16('0'),
	11: uint16(85),
	12: uint16(':'),
	13: uint16(61),
	14: uint16(';'),
	15: uint16(123),
	16: uint16('<'),
	17: uint16(8),
	18: uint16('='),
	19: uint16(70),
	20: uint16('>'),
	21: uint16(64),
	22: uint16('['),
	23: uint16(68),
	24: uint16(']'),
	25: uint16(69),
	26: uint16('}'),
	27: uint16(67),
}

var map_token2 = [44]uint16_t{
	0:  uint16('"'),
	1:  uint16(96),
	2:  uint16('\''),
	3:  uint16(97),
	4:  uint16('('),
	5:  uint16(73),
	6:  uint16(')'),
	7:  uint16(74),
	8:  uint16('+'),
	9:  uint16(90),
	10: uint16(','),
	11: uint16(66),
	12: uint16('-'),
	13: uint16(92),
	14: uint16('.'),
	15: uint16(82),
	16: uint16('0'),
	17: uint16(84),
	18: uint16(':'),
	19: uint16(61),
	20: uint16(';'),
	21: uint16(123),
	22: uint16('<'),
	23: uint16(63),
	24: uint16('='),
	25: uint16(71),
	26: uint16('>'),
	27: uint16(64),
	28: uint16('@'),
	29: uint16(11),
	30: uint16('['),
	31: uint16(68),
	32: uint16(']'),
	33: uint16(69),
	34: uint16('d'),
	35: uint16(109),
	36: uint16('r'),
	37: uint16(113),
	38: uint16('w'),
	39: uint16(116),
	40: uint16('{'),
	41: uint16(65),
	42: uint16('}'),
	43: uint16(67),
}

var map_token3 = [30]uint16_t{
	0:  uint16('"'),
	1:  uint16(15),
	2:  uint16('('),
	3:  uint16(73),
	4:  uint16('+'),
	5:  uint16(89),
	6:  uint16(','),
	7:  uint16(66),
	8:  uint16('-'),
	9:  uint16(91),
	10: uint16('.'),
	11: uint16(82),
	12: uint16('0'),
	13: uint16(83),
	14: uint16(':'),
	15: uint16(61),
	16: uint16(';'),
	17: uint16(123),
	18: uint16('<'),
	19: uint16(62),
	20: uint16('='),
	21: uint16(10),
	22: uint16('@'),
	23: uint16(11),
	24: uint16('['),
	25: uint16(68),
	26: uint16('{'),
	27: uint16(65),
	28: uint16('}'),
	29: uint16(67),
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
			if !(uint64(i) < libc.Uint64FromInt64(108)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token4[i]) == lookahead {
				state = map_token4[i+uint32(1)]
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
		if lookahead == int32('n') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('l') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('i') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('e') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('I') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('I') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('d') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('i') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('a') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('v') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('e') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('l') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('q') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('l') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('e') {
			state = uint16(53)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('e') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('n') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('e') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('e') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(62)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('e') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('f') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('a') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('d') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('h') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(75)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(77)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('a') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('n') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('a') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('h') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(83)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('o') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('a') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('y') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('o') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('x') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('s') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('n') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('n') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('d') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('d') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('A') {
			state = uint16(95)
			goto next_state
		}
		if lookahead == int32('C') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32('F') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('S') {
			state = uint16(98)
			goto next_state
		}
		if lookahead == int32('U') {
			state = uint16(99)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('t') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('t') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('s') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('t') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('r') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('e') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('n') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('v') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('t') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('f') {
			state = uint16(110)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('v') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('h') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('s') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_eq)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(51):
		if lookahead == int32('t') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('i') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('q') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_gt)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		if lookahead == int32('a') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('f') {
			state = uint16(119)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(120)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('v') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_is)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(58):
		if lookahead == int32('q') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_lt)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(60):
		if lookahead == int32('m') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('d') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('o') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('l') {
			state = uint16(127)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('g') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('q') {
			state = uint16(130)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('d') {
			state = uint16(132)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_of)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(67):
		if lookahead == int32('d') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_or)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('t') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('d') {
			state = uint16(137)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('i') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('w') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('a') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('g') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('l') {
			state = uint16(145)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('i') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('e') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('o') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('b') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead == int32('i') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead == int32('d') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead == int32('l') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead == int32('e') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('r') {
			state = uint16(155)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('i') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('r') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('l') {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead == int32('n') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead == int32('c') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead == int32('e') {
			state = uint16(162)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('e') {
			state = uint16(163)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('t') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('t') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_add)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_and)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(166)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead == int32('s') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead == int32('l') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead == int32('i') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead == int32('I') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead == int32('I') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead == int32('e') {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('a') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead == int32('s') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead == int32('e') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('h') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_cat)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(105):
		if lookahead == int32('c') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead == int32('m') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead == int32('s') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead == int32('e') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_cvt)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(110):
		if lookahead == int32('n') {
			state = uint16(182)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead == int32('t') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_div)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(113):
		if lookahead == int32('l') {
			state = uint16(184)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('e') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('m') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead == int32('p') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_geq)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(118):
		if lookahead == int32('d') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('e') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('u') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('t') {
			state = uint16(192)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('a') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_leq)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mem)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(125):
		if lookahead == int32('u') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead == int32('r') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mul)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mux)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_neg)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_neq)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_new)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(132):
		if lookahead == int32('e') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_not)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_old)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_orr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(136):
		if lookahead == int32('p') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_pad)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(138):
		if lookahead == int32('a') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(139):
		if lookahead == int32('n') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(140):
		if lookahead == int32('r') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead == int32('d') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_reg)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(143):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rem)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(144):
		if lookahead == int32('e') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_shl)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(146):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_shr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(147):
		if lookahead == int32('p') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(148):
		if lookahead == int32('m') {
			state = uint16(204)
			goto next_state
		}
		return result
	case int32(149):
		if lookahead == int32('p') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(150):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sub)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(151):
		if lookahead == int32('l') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(152):
		if lookahead == int32('e') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead == int32('i') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(154):
		if lookahead == int32('n') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(155):
		if lookahead == int32('e') {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(156):
		if lookahead == int32('h') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(157):
		if lookahead == int32('t') {
			state = uint16(212)
			goto next_state
		}
		return result
	case int32(158):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_xor)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(159):
		if lookahead == int32('o') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(160):
		if lookahead == int32('c') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(161):
		if lookahead == int32('k') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(162):
		if lookahead == int32('d') {
			state = uint16(217)
			goto next_state
		}
		return result
	case int32(163):
		if lookahead == int32('t') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(164):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SInt)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(165):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_UInt)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(166):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_andr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(167):
		if lookahead == int32('y') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(168):
		if lookahead == int32('o') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(169):
		if lookahead == int32('x') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(170):
		if lookahead == int32('n') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(171):
		if lookahead == int32('n') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(172):
		if lookahead == int32('r') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(173):
		if lookahead == int32('m') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(174):
		if lookahead == int32('c') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(175):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bits)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(176):
		if lookahead == int32('t') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(177):
		if lookahead == int32('l') {
			state = uint16(228)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(229)
			goto next_state
		}
		return result
	case int32(178):
		if lookahead == int32('u') {
			state = uint16(230)
			goto next_state
		}
		return result
	case int32(179):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_cmem)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(180):
		if lookahead == int32('t') {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(181):
		if lookahead == int32('r') {
			state = uint16(232)
			goto next_state
		}
		return result
	case int32(182):
		if lookahead == int32('a') {
			state = uint16(233)
			goto next_state
		}
		return result
	case int32(183):
		if lookahead == int32('h') {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(184):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_dshl)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('w') {
			state = uint16(235)
			goto next_state
		}
		return result
	case int32(185):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_dshr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('w') {
			state = uint16(236)
			goto next_state
		}
		return result
	case int32(186):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_else)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(187):
		if lookahead == int32('o') {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(188):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_flip)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(189):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_head)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(190):
		if lookahead == int32('r') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(191):
		if lookahead == int32('t') {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(192):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_inst)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(193):
		if lookahead == int32('l') {
			state = uint16(240)
			goto next_state
		}
		return result
	case int32(194):
		if lookahead == int32('l') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(195):
		if lookahead == int32('t') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(196):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_node)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(197):
		if lookahead == int32('u') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(198):
		if lookahead == int32('m') {
			state = uint16(244)
			goto next_state
		}
		return result
	case int32(199):
		if lookahead == int32('t') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(200):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rdwr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(201):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_read)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(246)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(202):
		if lookahead == int32('t') {
			state = uint16(248)
			goto next_state
		}
		return result
	case int32(203):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_skip)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(204):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_smem)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(205):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_stop)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(206):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_tail)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(207):
		if lookahead == int32('f') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(208):
		if lookahead == int32('d') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(209):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_when)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(210):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_wire)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(211):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_with)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(212):
		if lookahead == int32('e') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(213):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_xorr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(214):
		if lookahead == int32('g') {
			state = uint16(252)
			goto next_state
		}
		return result
	case int32(215):
		if lookahead == int32('R') {
			state = uint16(253)
			goto next_state
		}
		return result
	case int32(216):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Clock)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(217):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Fixed)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(218):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Reset)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(219):
		if lookahead == int32('n') {
			state = uint16(254)
			goto next_state
		}
		return result
	case int32(220):
		if lookahead == int32('c') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(221):
		if lookahead == int32('e') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(222):
		if lookahead == int32('t') {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(223):
		if lookahead == int32('t') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(224):
		if lookahead == int32('t') {
			state = uint16(259)
			goto next_state
		}
		return result
	case int32(225):
		if lookahead == int32('e') {
			state = uint16(260)
			goto next_state
		}
		return result
	case int32(226):
		if lookahead == int32('h') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(227):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bpset)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(228):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bpshl)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(229):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bpshr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(230):
		if lookahead == int32('i') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(231):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_const)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(232):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_cover)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(233):
		if lookahead == int32('m') {
			state = uint16(263)
			goto next_state
		}
		return result
	case int32(234):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_depth)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(235):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_dshlw)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(236):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_dshrw)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(237):
		if lookahead == int32('d') {
			state = uint16(264)
			goto next_state
		}
		return result
	case int32(238):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_infer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(239):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_input)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(240):
		if lookahead == int32('i') {
			state = uint16(265)
			goto next_state
		}
		return result
	case int32(241):
		if lookahead == int32('e') {
			state = uint16(266)
			goto next_state
		}
		return result
	case int32(242):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mport)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(243):
		if lookahead == int32('t') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(244):
		if lookahead == int32('e') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(245):
		if lookahead == int32('f') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(246):
		if lookahead == int32('r') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(247):
		if lookahead == int32('r') {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(248):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_reset)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(249):
		if lookahead == int32('i') {
			state = uint16(272)
			goto next_state
		}
		return result
	case int32(250):
		if lookahead == int32('i') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(251):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_write)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(274)
			goto next_state
		}
		return result
	case int32(252):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Analog)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(253):
		if lookahead == int32('e') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(254):
		if lookahead == int32('c') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(255):
		if lookahead == int32('k') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(256):
		if lookahead == int32('d') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(257):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_asSInt)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(258):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_asUInt)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(259):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_assert)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(260):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_assume)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(261):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_attach)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(262):
		if lookahead == int32('t') {
			state = uint16(279)
			goto next_state
		}
		return result
	case int32(263):
		if lookahead == int32('e') {
			state = uint16(280)
			goto next_state
		}
		return result
	case int32(264):
		if lookahead == int32('u') {
			state = uint16(281)
			goto next_state
		}
		return result
	case int32(265):
		if lookahead == int32('d') {
			state = uint16(282)
			goto next_state
		}
		return result
	case int32(266):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_module)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(267):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_output)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(268):
		if lookahead == int32('t') {
			state = uint16(283)
			goto next_state
		}
		return result
	case int32(269):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_printf)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(270):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_reader)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(271):
		if lookahead == int32('i') {
			state = uint16(284)
			goto next_state
		}
		return result
	case int32(272):
		if lookahead == int32('n') {
			state = uint16(285)
			goto next_state
		}
		return result
	case int32(273):
		if lookahead == int32('f') {
			state = uint16(286)
			goto next_state
		}
		return result
	case int32(274):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_writer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(275):
		if lookahead == int32('s') {
			state = uint16(287)
			goto next_state
		}
		return result
	case int32(276):
		if lookahead == int32('R') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(277):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_asClock)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(278):
		if lookahead == int32('P') {
			state = uint16(289)
			goto next_state
		}
		return result
	case int32(279):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_circuit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(280):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_defname)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(281):
		if lookahead == int32('l') {
			state = uint16(290)
			goto next_state
		}
		return result
	case int32(282):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_invalid)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(283):
		if lookahead == int32('e') {
			state = uint16(291)
			goto next_state
		}
		return result
	case int32(284):
		if lookahead == int32('t') {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(285):
		if lookahead == int32('e') {
			state = uint16(293)
			goto next_state
		}
		return result
	case int32(286):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_validif)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(287):
		if lookahead == int32('e') {
			state = uint16(294)
			goto next_state
		}
		return result
	case int32(288):
		if lookahead == int32('e') {
			state = uint16(295)
			goto next_state
		}
		return result
	case int32(289):
		if lookahead == int32('o') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(290):
		if lookahead == int32('e') {
			state = uint16(297)
			goto next_state
		}
		return result
	case int32(291):
		if lookahead == int32('r') {
			state = uint16(298)
			goto next_state
		}
		return result
	case int32(292):
		if lookahead == int32('e') {
			state = uint16(299)
			goto next_state
		}
		return result
	case int32(293):
		if lookahead == int32('d') {
			state = uint16(300)
			goto next_state
		}
		return result
	case int32(294):
		if lookahead == int32('t') {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(295):
		if lookahead == int32('s') {
			state = uint16(302)
			goto next_state
		}
		return result
	case int32(296):
		if lookahead == int32('i') {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(297):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_extmodule)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(298):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_parameter)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(299):
		if lookahead == int32('r') {
			state = uint16(304)
			goto next_state
		}
		return result
	case int32(300):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_undefined)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(301):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AsyncReset)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(302):
		if lookahead == int32('e') {
			state = uint16(305)
			goto next_state
		}
		return result
	case int32(303):
		if lookahead == int32('n') {
			state = uint16(306)
			goto next_state
		}
		return result
	case int32(304):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_readwriter)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(305):
		if lookahead == int32('t') {
			state = uint16(307)
			goto next_state
		}
		return result
	case int32(306):
		if lookahead == int32('t') {
			state = uint16(308)
			goto next_state
		}
		return result
	case int32(307):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_asAsyncReset)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(308):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_asFixedPoint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token4 = [54]uint16_t{
	0:  uint16('A'),
	1:  uint16(1),
	2:  uint16('C'),
	3:  uint16(2),
	4:  uint16('F'),
	5:  uint16(3),
	6:  uint16('R'),
	7:  uint16(4),
	8:  uint16('S'),
	9:  uint16(5),
	10: uint16('U'),
	11: uint16(6),
	12: uint16('a'),
	13: uint16(7),
	14: uint16('b'),
	15: uint16(8),
	16: uint16('c'),
	17: uint16(9),
	18: uint16('d'),
	19: uint16(10),
	20: uint16('e'),
	21: uint16(11),
	22: uint16('f'),
	23: uint16(12),
	24: uint16('g'),
	25: uint16(13),
	26: uint16('h'),
	27: uint16(14),
	28: uint16('i'),
	29: uint16(15),
	30: uint16('l'),
	31: uint16(16),
	32: uint16('m'),
	33: uint16(17),
	34: uint16('n'),
	35: uint16(18),
	36: uint16('o'),
	37: uint16(19),
	38: uint16('p'),
	39: uint16(20),
	40: uint16('r'),
	41: uint16(21),
	42: uint16('s'),
	43: uint16(22),
	44: uint16('t'),
	45: uint16(23),
	46: uint16('u'),
	47: uint16(24),
	48: uint16('v'),
	49: uint16(25),
	50: uint16('w'),
	51: uint16(26),
	52: uint16('x'),
	53: uint16(27),
}

var ts_lex_modes = [443]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state: uint16(59),
	},
	2: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	3: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	4: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	5: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	6: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(3),
	},
	7: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(3),
	},
	8: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(4),
	},
	9: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	10: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	11: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(4),
	},
	12: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	13: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	14: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	15: {
		Flex_state: uint16(59),
	},
	16: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	17: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	18: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	19: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	20: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	21: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	22: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	23: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	24: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	25: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	26: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	27: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	28: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	29: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	30: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	31: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	32: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	33: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	34: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	35: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	36: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	37: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	38: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	39: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	40: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	41: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	42: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	43: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	44: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	45: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	46: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	47: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	48: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	49: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	50: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	51: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	52: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	53: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	54: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	55: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	56: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	57: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	58: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	59: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	60: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	61: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	62: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	63: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	64: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	65: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	66: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	67: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	68: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	69: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	70: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	71: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	72: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	73: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	74: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	75: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	76: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	77: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	78: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	79: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	80: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	81: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	82: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	83: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	84: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	85: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	86: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	87: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	88: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	89: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	90: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	91: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	92: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	93: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	94: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	95: {
		Flex_state: uint16(59),
	},
	96: {
		Flex_state: uint16(59),
	},
	97: {
		Flex_state: uint16(59),
	},
	98: {
		Flex_state: uint16(59),
	},
	99: {
		Flex_state: uint16(59),
	},
	100: {
		Flex_state: uint16(59),
	},
	101: {
		Flex_state: uint16(59),
	},
	102: {
		Flex_state: uint16(59),
	},
	103: {
		Flex_state: uint16(59),
	},
	104: {
		Flex_state: uint16(59),
	},
	105: {
		Flex_state: uint16(59),
	},
	106: {
		Flex_state: uint16(59),
	},
	107: {
		Flex_state: uint16(59),
	},
	108: {
		Flex_state: uint16(59),
	},
	109: {
		Flex_state: uint16(59),
	},
	110: {
		Flex_state: uint16(59),
	},
	111: {
		Flex_state: uint16(59),
	},
	112: {
		Flex_state: uint16(59),
	},
	113: {
		Flex_state: uint16(59),
	},
	114: {
		Flex_state: uint16(59),
	},
	115: {
		Flex_state: uint16(59),
	},
	116: {
		Flex_state: uint16(59),
	},
	117: {
		Flex_state: uint16(59),
	},
	118: {
		Flex_state: uint16(59),
	},
	119: {
		Flex_state: uint16(59),
	},
	120: {
		Flex_state: uint16(59),
	},
	121: {
		Flex_state: uint16(59),
	},
	122: {
		Flex_state: uint16(59),
	},
	123: {
		Flex_state: uint16(59),
	},
	124: {
		Flex_state: uint16(59),
	},
	125: {
		Flex_state: uint16(59),
	},
	126: {
		Flex_state: uint16(59),
	},
	127: {
		Flex_state: uint16(59),
	},
	128: {
		Flex_state: uint16(59),
	},
	129: {
		Flex_state: uint16(59),
	},
	130: {},
	131: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	132: {
		Fexternal_lex_state: uint16(2),
	},
	133: {
		Fexternal_lex_state: uint16(2),
	},
	134: {
		Fexternal_lex_state: uint16(2),
	},
	135: {
		Flex_state: uint16(59),
	},
	136: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	137: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	138: {
		Fexternal_lex_state: uint16(2),
	},
	139: {
		Fexternal_lex_state: uint16(2),
	},
	140: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	141: {
		Flex_state: uint16(5),
	},
	142: {
		Flex_state: uint16(5),
	},
	143: {
		Flex_state: uint16(5),
	},
	144: {
		Flex_state: uint16(5),
	},
	145: {
		Fexternal_lex_state: uint16(2),
	},
	146: {
		Flex_state: uint16(5),
	},
	147: {
		Flex_state: uint16(5),
	},
	148: {
		Flex_state: uint16(59),
	},
	149: {
		Flex_state: uint16(59),
	},
	150: {
		Flex_state: uint16(59),
	},
	151: {
		Flex_state: uint16(5),
	},
	152: {
		Flex_state: uint16(5),
	},
	153: {
		Flex_state: uint16(5),
	},
	154: {
		Flex_state: uint16(5),
	},
	155: {
		Fexternal_lex_state: uint16(2),
	},
	156: {
		Flex_state: uint16(59),
	},
	157: {
		Flex_state: uint16(5),
	},
	158: {
		Flex_state: uint16(59),
	},
	159: {
		Flex_state: uint16(59),
	},
	160: {
		Flex_state: uint16(5),
	},
	161: {
		Flex_state: uint16(59),
	},
	162: {
		Flex_state: uint16(5),
	},
	163: {
		Flex_state: uint16(5),
	},
	164: {
		Flex_state: uint16(59),
	},
	165: {
		Flex_state: uint16(5),
	},
	166: {
		Flex_state: uint16(5),
	},
	167: {
		Flex_state: uint16(5),
	},
	168: {
		Flex_state: uint16(5),
	},
	169: {
		Flex_state: uint16(59),
	},
	170: {
		Flex_state: uint16(59),
	},
	171: {
		Flex_state: uint16(59),
	},
	172: {
		Flex_state: uint16(59),
	},
	173: {
		Flex_state: uint16(59),
	},
	174: {
		Flex_state: uint16(59),
	},
	175: {
		Flex_state: uint16(59),
	},
	176: {
		Flex_state: uint16(59),
	},
	177: {
		Flex_state: uint16(59),
	},
	178: {
		Flex_state: uint16(59),
	},
	179: {
		Flex_state: uint16(59),
	},
	180: {
		Flex_state: uint16(59),
	},
	181: {
		Flex_state: uint16(59),
	},
	182: {
		Flex_state: uint16(59),
	},
	183: {
		Flex_state: uint16(59),
	},
	184: {
		Flex_state: uint16(59),
	},
	185: {
		Flex_state: uint16(59),
	},
	186: {
		Flex_state: uint16(59),
	},
	187: {
		Flex_state: uint16(59),
	},
	188: {
		Flex_state: uint16(59),
	},
	189: {
		Flex_state: uint16(59),
	},
	190: {
		Flex_state: uint16(6),
	},
	191: {
		Flex_state: uint16(6),
	},
	192: {
		Flex_state: uint16(1),
	},
	193: {
		Flex_state: uint16(6),
	},
	194: {
		Flex_state: uint16(1),
	},
	195: {
		Flex_state: uint16(1),
	},
	196: {
		Flex_state: uint16(5),
	},
	197: {
		Flex_state: uint16(5),
	},
	198: {
		Flex_state: uint16(5),
	},
	199: {
		Flex_state: uint16(1),
	},
	200: {
		Flex_state: uint16(1),
	},
	201: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	202: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	203: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	204: {},
	205: {
		Flex_state: uint16(5),
	},
	206: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(3),
	},
	207: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(3),
	},
	208: {},
	209: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	210: {},
	211: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	212: {
		Flex_state: uint16(59),
	},
	213: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(5),
	},
	214: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	215: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	216: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(3),
	},
	217: {},
	218: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(4),
	},
	219: {
		Flex_state: uint16(59),
	},
	220: {},
	221: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	222: {},
	223: {
		Flex_state: uint16(59),
	},
	224: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(5),
	},
	225: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	226: {
		Flex_state: uint16(59),
	},
	227: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(5),
	},
	228: {
		Flex_state: uint16(59),
	},
	229: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	230: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	231: {
		Flex_state: uint16(59),
	},
	232: {
		Flex_state: uint16(59),
	},
	233: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(3),
	},
	234: {
		Flex_state: uint16(59),
	},
	235: {},
	236: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	237: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	238: {},
	239: {
		Fexternal_lex_state: uint16(5),
	},
	240: {},
	241: {},
	242: {},
	243: {},
	244: {},
	245: {},
	246: {},
	247: {},
	248: {},
	249: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	250: {},
	251: {},
	252: {},
	253: {},
	254: {},
	255: {},
	256: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	257: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(4),
	},
	258: {},
	259: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	260: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	261: {},
	262: {},
	263: {
		Fexternal_lex_state: uint16(5),
	},
	264: {},
	265: {},
	266: {},
	267: {},
	268: {},
	269: {},
	270: {},
	271: {},
	272: {},
	273: {},
	274: {},
	275: {},
	276: {
		Fexternal_lex_state: uint16(5),
	},
	277: {
		Flex_state: uint16(59),
	},
	278: {},
	279: {},
	280: {},
	281: {},
	282: {},
	283: {},
	284: {},
	285: {
		Fexternal_lex_state: uint16(5),
	},
	286: {},
	287: {},
	288: {
		Fexternal_lex_state: uint16(4),
	},
	289: {},
	290: {
		Fexternal_lex_state: uint16(5),
	},
	291: {
		Fexternal_lex_state: uint16(5),
	},
	292: {
		Fexternal_lex_state: uint16(5),
	},
	293: {
		Fexternal_lex_state: uint16(5),
	},
	294: {},
	295: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	296: {},
	297: {
		Fexternal_lex_state: uint16(5),
	},
	298: {
		Fexternal_lex_state: uint16(5),
	},
	299: {},
	300: {},
	301: {},
	302: {},
	303: {},
	304: {},
	305: {
		Fexternal_lex_state: uint16(5),
	},
	306: {
		Fexternal_lex_state: uint16(5),
	},
	307: {},
	308: {
		Fexternal_lex_state: uint16(5),
	},
	309: {},
	310: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	311: {},
	312: {},
	313: {
		Flex_state: uint16(59),
	},
	314: {
		Fexternal_lex_state: uint16(5),
	},
	315: {
		Fexternal_lex_state: uint16(5),
	},
	316: {},
	317: {
		Flex_state: uint16(59),
	},
	318: {
		Flex_state: uint16(59),
	},
	319: {},
	320: {
		Flex_state: uint16(59),
	},
	321: {
		Flex_state:          uint16(59),
		Fexternal_lex_state: uint16(2),
	},
	322: {},
	323: {
		Fexternal_lex_state: uint16(5),
	},
	324: {
		Fexternal_lex_state: uint16(5),
	},
	325: {
		Flex_state: uint16(59),
	},
	326: {
		Fexternal_lex_state: uint16(4),
	},
	327: {
		Flex_state: uint16(59),
	},
	328: {},
	329: {},
	330: {
		Fexternal_lex_state: uint16(2),
	},
	331: {
		Fexternal_lex_state: uint16(5),
	},
	332: {},
	333: {
		Flex_state: uint16(59),
	},
	334: {
		Flex_state: uint16(59),
	},
	335: {
		Flex_state: uint16(59),
	},
	336: {
		Fexternal_lex_state: uint16(2),
	},
	337: {
		Flex_state: uint16(59),
	},
	338: {},
	339: {},
	340: {},
	341: {},
	342: {
		Fexternal_lex_state: uint16(5),
	},
	343: {
		Fexternal_lex_state: uint16(5),
	},
	344: {
		Flex_state: uint16(59),
	},
	345: {
		Flex_state: uint16(5),
	},
	346: {
		Fexternal_lex_state: uint16(5),
	},
	347: {},
	348: {},
	349: {},
	350: {
		Flex_state: uint16(59),
	},
	351: {
		Flex_state: uint16(59),
	},
	352: {},
	353: {
		Flex_state: uint16(59),
	},
	354: {
		Flex_state: uint16(59),
	},
	355: {},
	356: {
		Flex_state: uint16(59),
	},
	357: {
		Fexternal_lex_state: uint16(5),
	},
	358: {},
	359: {},
	360: {},
	361: {
		Flex_state: uint16(59),
	},
	362: {
		Flex_state: uint16(59),
	},
	363: {
		Flex_state: uint16(59),
	},
	364: {},
	365: {
		Flex_state: uint16(59),
	},
	366: {},
	367: {},
	368: {},
	369: {
		Flex_state: uint16(59),
	},
	370: {
		Flex_state: uint16(59),
	},
	371: {
		Flex_state: uint16(59),
	},
	372: {},
	373: {
		Flex_state: uint16(59),
	},
	374: {
		Flex_state: uint16(59),
	},
	375: {
		Flex_state: uint16(59),
	},
	376: {
		Flex_state: uint16(59),
	},
	377: {},
	378: {},
	379: {
		Flex_state: uint16(59),
	},
	380: {},
	381: {
		Flex_state: uint16(59),
	},
	382: {},
	383: {
		Flex_state: uint16(59),
	},
	384: {},
	385: {},
	386: {},
	387: {
		Flex_state: uint16(59),
	},
	388: {
		Fexternal_lex_state: uint16(5),
	},
	389: {
		Fexternal_lex_state: uint16(5),
	},
	390: {
		Fexternal_lex_state: uint16(4),
	},
	391: {},
	392: {},
	393: {},
	394: {},
	395: {},
	396: {},
	397: {},
	398: {},
	399: {},
	400: {},
	401: {
		Flex_state: uint16(59),
	},
	402: {},
	403: {},
	404: {
		Fexternal_lex_state: uint16(5),
	},
	405: {},
	406: {},
	407: {},
	408: {
		Flex_state: uint16(5),
	},
	409: {
		Flex_state: uint16(59),
	},
	410: {},
	411: {
		Flex_state: uint16(5),
	},
	412: {},
	413: {
		Fexternal_lex_state: uint16(5),
	},
	414: {
		Fexternal_lex_state: uint16(5),
	},
	415: {
		Fexternal_lex_state: uint16(5),
	},
	416: {},
	417: {},
	418: {
		Fexternal_lex_state: uint16(5),
	},
	419: {},
	420: {},
	421: {},
	422: {
		Flex_state: uint16(5),
	},
	423: {},
	424: {},
	425: {},
	426: {},
	427: {
		Fexternal_lex_state: uint16(5),
	},
	428: {
		Fexternal_lex_state: uint16(5),
	},
	429: {
		Flex_state: uint16(59),
	},
	430: {
		Flex_state: uint16(59),
	},
	431: {},
	432: {},
	433: {},
	434: {
		Flex_state: uint16(59),
	},
	435: {
		Flex_state: uint16(59),
	},
	436: {
		Flex_state: uint16(59),
	},
	437: {
		Flex_state: uint16(59),
	},
	438: {},
	439: {},
	440: {},
	441: {},
	442: {
		Flex_state: uint16(59),
	},
}

var ts_parse_table = [95][194]uint16_t{
	0: {
		0:   uint16(1),
		1:   uint16(1),
		2:   uint16(1),
		3:   uint16(1),
		4:   uint16(1),
		5:   uint16(1),
		6:   uint16(1),
		7:   uint16(1),
		8:   uint16(1),
		9:   uint16(1),
		10:  uint16(1),
		11:  uint16(1),
		12:  uint16(1),
		13:  uint16(1),
		14:  uint16(1),
		15:  uint16(1),
		16:  uint16(1),
		17:  uint16(1),
		18:  uint16(1),
		19:  uint16(1),
		20:  uint16(1),
		21:  uint16(1),
		22:  uint16(1),
		23:  uint16(1),
		24:  uint16(1),
		25:  uint16(1),
		26:  uint16(1),
		27:  uint16(1),
		28:  uint16(1),
		29:  uint16(1),
		30:  uint16(1),
		31:  uint16(1),
		32:  uint16(1),
		33:  uint16(1),
		34:  uint16(1),
		35:  uint16(1),
		36:  uint16(1),
		37:  uint16(1),
		38:  uint16(1),
		39:  uint16(1),
		40:  uint16(1),
		41:  uint16(1),
		43:  uint16(1),
		44:  uint16(1),
		45:  uint16(1),
		46:  uint16(1),
		47:  uint16(1),
		48:  uint16(1),
		49:  uint16(1),
		50:  uint16(1),
		51:  uint16(1),
		52:  uint16(1),
		53:  uint16(1),
		54:  uint16(1),
		55:  uint16(1),
		56:  uint16(1),
		57:  uint16(1),
		58:  uint16(1),
		59:  uint16(1),
		60:  uint16(1),
		61:  uint16(1),
		62:  uint16(1),
		63:  uint16(1),
		64:  uint16(1),
		65:  uint16(1),
		66:  uint16(1),
		67:  uint16(1),
		68:  uint16(1),
		69:  uint16(1),
		70:  uint16(1),
		71:  uint16(1),
		72:  uint16(1),
		73:  uint16(1),
		74:  uint16(1),
		75:  uint16(1),
		76:  uint16(1),
		77:  uint16(1),
		78:  uint16(1),
		79:  uint16(1),
		80:  uint16(1),
		81:  uint16(1),
		82:  uint16(1),
		83:  uint16(1),
		84:  uint16(1),
		85:  uint16(1),
		86:  uint16(1),
		87:  uint16(1),
		88:  uint16(1),
		89:  uint16(1),
		90:  uint16(1),
		91:  uint16(1),
		92:  uint16(1),
		93:  uint16(1),
		94:  uint16(1),
		95:  uint16(1),
		96:  uint16(1),
		97:  uint16(1),
		98:  uint16(1),
		99:  uint16(1),
		100: uint16(1),
		101: uint16(1),
		102: uint16(1),
		103: uint16(1),
		104: uint16(1),
		105: uint16(1),
		106: uint16(1),
		107: uint16(1),
		108: uint16(1),
		109: uint16(1),
		110: uint16(1),
		111: uint16(1),
		112: uint16(1),
		113: uint16(1),
		114: uint16(1),
		115: uint16(1),
		117: uint16(1),
		118: uint16(1),
		119: uint16(1),
		122: uint16(1),
		123: uint16(1),
		125: uint16(3),
		126: uint16(1),
		127: uint16(1),
		128: uint16(1),
	},
	1: {
		0:   uint16(5),
		2:   uint16(7),
		125: uint16(3),
		129: uint16(410),
		130: uint16(234),
		182: uint16(234),
	},
	2: {
		1:   uint16(9),
		6:   uint16(11),
		7:   uint16(11),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		128: uint16(49),
		132: uint16(16),
		133: uint16(429),
		142: uint16(9),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		156: uint16(65),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
		184: uint16(16),
		185: uint16(9),
	},
	3: {
		1:   uint16(9),
		6:   uint16(11),
		7:   uint16(11),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		128: uint16(51),
		132: uint16(16),
		133: uint16(429),
		142: uint16(13),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		156: uint16(65),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
		184: uint16(16),
		185: uint16(13),
	},
	4: {
		1:   uint16(9),
		6:   uint16(11),
		7:   uint16(11),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		128: uint16(49),
		132: uint16(3),
		133: uint16(429),
		142: uint16(9),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		156: uint16(65),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
		184: uint16(3),
		185: uint16(9),
	},
	5: {
		1:   uint16(9),
		6:   uint16(11),
		7:   uint16(11),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		128: uint16(53),
		132: uint16(2),
		133: uint16(429),
		142: uint16(12),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		156: uint16(65),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
		184: uint16(2),
		185: uint16(12),
	},
	6: {
		1:   uint16(9),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		54:  uint16(55),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		62:  uint16(57),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		127: uint16(59),
		128: uint16(61),
		142: uint16(79),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		155: uint16(63),
		156: uint16(65),
		157: uint16(77),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
	},
	7: {
		1:   uint16(9),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		54:  uint16(55),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		127: uint16(59),
		128: uint16(63),
		142: uint16(79),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		155: uint16(58),
		156: uint16(65),
		157: uint16(76),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
	},
	8: {
		1:   uint16(9),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		62:  uint16(65),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		127: uint16(59),
		142: uint16(79),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		155: uint16(86),
		156: uint16(65),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
	},
	9: {
		1:   uint16(9),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		128: uint16(51),
		142: uint16(10),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		156: uint16(65),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
		185: uint16(10),
	},
	10: {
		1:   uint16(67),
		9:   uint16(70),
		10:  uint16(70),
		31:  uint16(73),
		32:  uint16(76),
		33:  uint16(79),
		34:  uint16(82),
		36:  uint16(85),
		38:  uint16(88),
		40:  uint16(91),
		53:  uint16(94),
		55:  uint16(97),
		56:  uint16(100),
		57:  uint16(103),
		58:  uint16(103),
		59:  uint16(103),
		60:  uint16(106),
		61:  uint16(109),
		63:  uint16(112),
		64:  uint16(112),
		65:  uint16(112),
		66:  uint16(112),
		71:  uint16(115),
		72:  uint16(118),
		73:  uint16(121),
		74:  uint16(121),
		75:  uint16(121),
		76:  uint16(121),
		77:  uint16(121),
		78:  uint16(121),
		79:  uint16(121),
		80:  uint16(121),
		81:  uint16(121),
		82:  uint16(121),
		83:  uint16(121),
		84:  uint16(121),
		85:  uint16(121),
		86:  uint16(121),
		87:  uint16(121),
		88:  uint16(121),
		89:  uint16(121),
		90:  uint16(121),
		91:  uint16(121),
		92:  uint16(121),
		93:  uint16(121),
		94:  uint16(121),
		95:  uint16(121),
		96:  uint16(121),
		97:  uint16(121),
		98:  uint16(121),
		99:  uint16(121),
		100: uint16(121),
		101: uint16(121),
		102: uint16(121),
		103: uint16(121),
		104: uint16(121),
		105: uint16(121),
		106: uint16(121),
		107: uint16(121),
		108: uint16(121),
		109: uint16(121),
		110: uint16(121),
		111: uint16(121),
		125: uint16(3),
		128: uint16(124),
		142: uint16(10),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		156: uint16(65),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
		185: uint16(10),
	},
	11: {
		1:   uint16(9),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		127: uint16(59),
		142: uint16(79),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		155: uint16(92),
		156: uint16(65),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
	},
	12: {
		1:   uint16(9),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		128: uint16(49),
		142: uint16(10),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		156: uint16(65),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
		185: uint16(10),
	},
	13: {
		1:   uint16(9),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		128: uint16(126),
		142: uint16(10),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		156: uint16(65),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
		185: uint16(10),
	},
	14: {
		1:   uint16(9),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		128: uint16(128),
		142: uint16(10),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		156: uint16(65),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
		185: uint16(10),
	},
	15: {
		1:   uint16(9),
		9:   uint16(13),
		10:  uint16(13),
		31:  uint16(15),
		32:  uint16(17),
		33:  uint16(19),
		34:  uint16(21),
		36:  uint16(23),
		38:  uint16(25),
		40:  uint16(27),
		53:  uint16(29),
		55:  uint16(31),
		56:  uint16(33),
		57:  uint16(35),
		58:  uint16(35),
		59:  uint16(35),
		60:  uint16(37),
		61:  uint16(39),
		63:  uint16(41),
		64:  uint16(41),
		65:  uint16(41),
		66:  uint16(41),
		71:  uint16(43),
		72:  uint16(45),
		73:  uint16(47),
		74:  uint16(47),
		75:  uint16(47),
		76:  uint16(47),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		89:  uint16(47),
		90:  uint16(47),
		91:  uint16(47),
		92:  uint16(47),
		93:  uint16(47),
		94:  uint16(47),
		95:  uint16(47),
		96:  uint16(47),
		97:  uint16(47),
		98:  uint16(47),
		99:  uint16(47),
		100: uint16(47),
		101: uint16(47),
		102: uint16(47),
		103: uint16(47),
		104: uint16(47),
		105: uint16(47),
		106: uint16(47),
		107: uint16(47),
		108: uint16(47),
		109: uint16(47),
		110: uint16(47),
		111: uint16(47),
		125: uint16(3),
		142: uint16(14),
		143: uint16(65),
		144: uint16(65),
		145: uint16(65),
		146: uint16(65),
		147: uint16(65),
		148: uint16(65),
		149: uint16(65),
		150: uint16(65),
		151: uint16(65),
		152: uint16(65),
		153: uint16(65),
		156: uint16(65),
		158: uint16(65),
		159: uint16(65),
		160: uint16(65),
		161: uint16(65),
		162: uint16(65),
		163: uint16(430),
		165: uint16(433),
		166: uint16(205),
		167: uint16(160),
		168: uint16(160),
		169: uint16(160),
		170: uint16(160),
		171: uint16(160),
		172: uint16(160),
		173: uint16(160),
		175: uint16(143),
		185: uint16(14),
	},
	16: {
		1:   uint16(130),
		6:   uint16(132),
		7:   uint16(132),
		9:   uint16(130),
		10:  uint16(130),
		24:  uint16(130),
		26:  uint16(130),
		31:  uint16(130),
		32:  uint16(130),
		33:  uint16(130),
		34:  uint16(130),
		36:  uint16(130),
		38:  uint16(130),
		40:  uint16(130),
		53:  uint16(130),
		55:  uint16(130),
		56:  uint16(130),
		57:  uint16(130),
		58:  uint16(130),
		59:  uint16(130),
		60:  uint16(130),
		61:  uint16(130),
		63:  uint16(130),
		64:  uint16(130),
		65:  uint16(130),
		66:  uint16(130),
		71:  uint16(130),
		72:  uint16(130),
		73:  uint16(130),
		74:  uint16(130),
		75:  uint16(130),
		76:  uint16(130),
		77:  uint16(130),
		78:  uint16(130),
		79:  uint16(130),
		80:  uint16(130),
		81:  uint16(130),
		82:  uint16(130),
		83:  uint16(130),
		84:  uint16(130),
		85:  uint16(130),
		86:  uint16(130),
		87:  uint16(130),
		88:  uint16(130),
		89:  uint16(130),
		90:  uint16(130),
		91:  uint16(130),
		92:  uint16(130),
		93:  uint16(130),
		94:  uint16(130),
		95:  uint16(130),
		96:  uint16(130),
		97:  uint16(130),
		98:  uint16(130),
		99:  uint16(130),
		100: uint16(130),
		101: uint16(130),
		102: uint16(130),
		103: uint16(130),
		104: uint16(130),
		105: uint16(130),
		106: uint16(130),
		107: uint16(130),
		108: uint16(130),
		109: uint16(130),
		110: uint16(130),
		111: uint16(130),
		125: uint16(3),
		128: uint16(135),
		132: uint16(16),
		133: uint16(429),
		184: uint16(16),
	},
	17: {
		1:   uint16(137),
		9:   uint16(137),
		10:  uint16(137),
		21:  uint16(139),
		29:  uint16(139),
		31:  uint16(137),
		32:  uint16(137),
		33:  uint16(137),
		34:  uint16(137),
		35:  uint16(137),
		36:  uint16(137),
		38:  uint16(137),
		40:  uint16(137),
		53:  uint16(137),
		54:  uint16(137),
		55:  uint16(137),
		56:  uint16(137),
		57:  uint16(137),
		58:  uint16(137),
		59:  uint16(137),
		60:  uint16(137),
		61:  uint16(137),
		62:  uint16(139),
		63:  uint16(137),
		64:  uint16(137),
		65:  uint16(137),
		66:  uint16(137),
		70:  uint16(139),
		71:  uint16(137),
		72:  uint16(137),
		73:  uint16(137),
		74:  uint16(137),
		75:  uint16(137),
		76:  uint16(137),
		77:  uint16(137),
		78:  uint16(137),
		79:  uint16(137),
		80:  uint16(137),
		81:  uint16(137),
		82:  uint16(137),
		83:  uint16(137),
		84:  uint16(137),
		85:  uint16(137),
		86:  uint16(137),
		87:  uint16(137),
		88:  uint16(137),
		89:  uint16(137),
		90:  uint16(137),
		91:  uint16(137),
		92:  uint16(137),
		93:  uint16(137),
		94:  uint16(137),
		95:  uint16(137),
		96:  uint16(137),
		97:  uint16(137),
		98:  uint16(137),
		99:  uint16(137),
		100: uint16(137),
		101: uint16(137),
		102: uint16(137),
		103: uint16(137),
		104: uint16(137),
		105: uint16(137),
		106: uint16(137),
		107: uint16(137),
		108: uint16(137),
		109: uint16(137),
		110: uint16(137),
		111: uint16(137),
		125: uint16(3),
		128: uint16(139),
	},
	18: {
		1:   uint16(141),
		9:   uint16(141),
		10:  uint16(141),
		21:  uint16(143),
		29:  uint16(145),
		31:  uint16(141),
		32:  uint16(141),
		33:  uint16(141),
		34:  uint16(141),
		35:  uint16(141),
		36:  uint16(141),
		38:  uint16(141),
		40:  uint16(141),
		53:  uint16(141),
		54:  uint16(141),
		55:  uint16(141),
		56:  uint16(141),
		57:  uint16(141),
		58:  uint16(141),
		59:  uint16(141),
		60:  uint16(141),
		61:  uint16(141),
		62:  uint16(143),
		63:  uint16(141),
		64:  uint16(141),
		65:  uint16(141),
		66:  uint16(141),
		70:  uint16(143),
		71:  uint16(141),
		72:  uint16(141),
		73:  uint16(141),
		74:  uint16(141),
		75:  uint16(141),
		76:  uint16(141),
		77:  uint16(141),
		78:  uint16(141),
		79:  uint16(141),
		80:  uint16(141),
		81:  uint16(141),
		82:  uint16(141),
		83:  uint16(141),
		84:  uint16(141),
		85:  uint16(141),
		86:  uint16(141),
		87:  uint16(141),
		88:  uint16(141),
		89:  uint16(141),
		90:  uint16(141),
		91:  uint16(141),
		92:  uint16(141),
		93:  uint16(141),
		94:  uint16(141),
		95:  uint16(141),
		96:  uint16(141),
		97:  uint16(141),
		98:  uint16(141),
		99:  uint16(141),
		100: uint16(141),
		101: uint16(141),
		102: uint16(141),
		103: uint16(141),
		104: uint16(141),
		105: uint16(141),
		106: uint16(141),
		107: uint16(141),
		108: uint16(141),
		109: uint16(141),
		110: uint16(141),
		111: uint16(141),
		125: uint16(3),
		128: uint16(143),
	},
	19: {
		1:   uint16(141),
		9:   uint16(141),
		10:  uint16(141),
		21:  uint16(143),
		29:  uint16(147),
		31:  uint16(141),
		32:  uint16(141),
		33:  uint16(141),
		34:  uint16(141),
		35:  uint16(141),
		36:  uint16(141),
		38:  uint16(141),
		40:  uint16(141),
		53:  uint16(141),
		54:  uint16(141),
		55:  uint16(141),
		56:  uint16(141),
		57:  uint16(141),
		58:  uint16(141),
		59:  uint16(141),
		60:  uint16(141),
		61:  uint16(141),
		62:  uint16(143),
		63:  uint16(141),
		64:  uint16(141),
		65:  uint16(141),
		66:  uint16(141),
		70:  uint16(143),
		71:  uint16(141),
		72:  uint16(141),
		73:  uint16(141),
		74:  uint16(141),
		75:  uint16(141),
		76:  uint16(141),
		77:  uint16(141),
		78:  uint16(141),
		79:  uint16(141),
		80:  uint16(141),
		81:  uint16(141),
		82:  uint16(141),
		83:  uint16(141),
		84:  uint16(141),
		85:  uint16(141),
		86:  uint16(141),
		87:  uint16(141),
		88:  uint16(141),
		89:  uint16(141),
		90:  uint16(141),
		91:  uint16(141),
		92:  uint16(141),
		93:  uint16(141),
		94:  uint16(141),
		95:  uint16(141),
		96:  uint16(141),
		97:  uint16(141),
		98:  uint16(141),
		99:  uint16(141),
		100: uint16(141),
		101: uint16(141),
		102: uint16(141),
		103: uint16(141),
		104: uint16(141),
		105: uint16(141),
		106: uint16(141),
		107: uint16(141),
		108: uint16(141),
		109: uint16(141),
		110: uint16(141),
		111: uint16(141),
		125: uint16(3),
		128: uint16(143),
	},
	20: {
		1:   uint16(149),
		9:   uint16(149),
		10:  uint16(149),
		21:  uint16(151),
		31:  uint16(149),
		32:  uint16(149),
		33:  uint16(149),
		34:  uint16(149),
		35:  uint16(149),
		36:  uint16(149),
		38:  uint16(149),
		40:  uint16(149),
		53:  uint16(149),
		54:  uint16(149),
		55:  uint16(149),
		56:  uint16(149),
		57:  uint16(149),
		58:  uint16(149),
		59:  uint16(149),
		60:  uint16(149),
		61:  uint16(149),
		62:  uint16(151),
		63:  uint16(149),
		64:  uint16(149),
		65:  uint16(149),
		66:  uint16(149),
		70:  uint16(151),
		71:  uint16(149),
		72:  uint16(149),
		73:  uint16(149),
		74:  uint16(149),
		75:  uint16(149),
		76:  uint16(149),
		77:  uint16(149),
		78:  uint16(149),
		79:  uint16(149),
		80:  uint16(149),
		81:  uint16(149),
		82:  uint16(149),
		83:  uint16(149),
		84:  uint16(149),
		85:  uint16(149),
		86:  uint16(149),
		87:  uint16(149),
		88:  uint16(149),
		89:  uint16(149),
		90:  uint16(149),
		91:  uint16(149),
		92:  uint16(149),
		93:  uint16(149),
		94:  uint16(149),
		95:  uint16(149),
		96:  uint16(149),
		97:  uint16(149),
		98:  uint16(149),
		99:  uint16(149),
		100: uint16(149),
		101: uint16(149),
		102: uint16(149),
		103: uint16(149),
		104: uint16(149),
		105: uint16(149),
		106: uint16(149),
		107: uint16(149),
		108: uint16(149),
		109: uint16(149),
		110: uint16(149),
		111: uint16(149),
		125: uint16(3),
		128: uint16(151),
	},
	21: {
		1:   uint16(153),
		9:   uint16(153),
		10:  uint16(153),
		21:  uint16(155),
		31:  uint16(153),
		32:  uint16(153),
		33:  uint16(153),
		34:  uint16(153),
		35:  uint16(153),
		36:  uint16(153),
		38:  uint16(153),
		40:  uint16(153),
		53:  uint16(153),
		54:  uint16(153),
		55:  uint16(153),
		56:  uint16(153),
		57:  uint16(153),
		58:  uint16(153),
		59:  uint16(153),
		60:  uint16(153),
		61:  uint16(153),
		62:  uint16(155),
		63:  uint16(153),
		64:  uint16(153),
		65:  uint16(153),
		66:  uint16(153),
		70:  uint16(155),
		71:  uint16(153),
		72:  uint16(153),
		73:  uint16(153),
		74:  uint16(153),
		75:  uint16(153),
		76:  uint16(153),
		77:  uint16(153),
		78:  uint16(153),
		79:  uint16(153),
		80:  uint16(153),
		81:  uint16(153),
		82:  uint16(153),
		83:  uint16(153),
		84:  uint16(153),
		85:  uint16(153),
		86:  uint16(153),
		87:  uint16(153),
		88:  uint16(153),
		89:  uint16(153),
		90:  uint16(153),
		91:  uint16(153),
		92:  uint16(153),
		93:  uint16(153),
		94:  uint16(153),
		95:  uint16(153),
		96:  uint16(153),
		97:  uint16(153),
		98:  uint16(153),
		99:  uint16(153),
		100: uint16(153),
		101: uint16(153),
		102: uint16(153),
		103: uint16(153),
		104: uint16(153),
		105: uint16(153),
		106: uint16(153),
		107: uint16(153),
		108: uint16(153),
		109: uint16(153),
		110: uint16(153),
		111: uint16(153),
		125: uint16(3),
		128: uint16(155),
	},
	22: {
		1:   uint16(157),
		9:   uint16(157),
		10:  uint16(157),
		21:  uint16(159),
		31:  uint16(157),
		32:  uint16(157),
		33:  uint16(157),
		34:  uint16(157),
		35:  uint16(157),
		36:  uint16(157),
		38:  uint16(157),
		40:  uint16(157),
		53:  uint16(157),
		54:  uint16(157),
		55:  uint16(157),
		56:  uint16(157),
		57:  uint16(157),
		58:  uint16(157),
		59:  uint16(157),
		60:  uint16(157),
		61:  uint16(157),
		62:  uint16(159),
		63:  uint16(157),
		64:  uint16(157),
		65:  uint16(157),
		66:  uint16(157),
		70:  uint16(159),
		71:  uint16(157),
		72:  uint16(157),
		73:  uint16(157),
		74:  uint16(157),
		75:  uint16(157),
		76:  uint16(157),
		77:  uint16(157),
		78:  uint16(157),
		79:  uint16(157),
		80:  uint16(157),
		81:  uint16(157),
		82:  uint16(157),
		83:  uint16(157),
		84:  uint16(157),
		85:  uint16(157),
		86:  uint16(157),
		87:  uint16(157),
		88:  uint16(157),
		89:  uint16(157),
		90:  uint16(157),
		91:  uint16(157),
		92:  uint16(157),
		93:  uint16(157),
		94:  uint16(157),
		95:  uint16(157),
		96:  uint16(157),
		97:  uint16(157),
		98:  uint16(157),
		99:  uint16(157),
		100: uint16(157),
		101: uint16(157),
		102: uint16(157),
		103: uint16(157),
		104: uint16(157),
		105: uint16(157),
		106: uint16(157),
		107: uint16(157),
		108: uint16(157),
		109: uint16(157),
		110: uint16(157),
		111: uint16(157),
		125: uint16(3),
		128: uint16(159),
	},
	23: {
		1:   uint16(161),
		9:   uint16(161),
		10:  uint16(161),
		21:  uint16(163),
		31:  uint16(161),
		32:  uint16(161),
		33:  uint16(161),
		34:  uint16(161),
		35:  uint16(161),
		36:  uint16(161),
		38:  uint16(161),
		40:  uint16(161),
		53:  uint16(161),
		54:  uint16(161),
		55:  uint16(161),
		56:  uint16(161),
		57:  uint16(161),
		58:  uint16(161),
		59:  uint16(161),
		60:  uint16(161),
		61:  uint16(161),
		62:  uint16(163),
		63:  uint16(161),
		64:  uint16(161),
		65:  uint16(161),
		66:  uint16(161),
		70:  uint16(163),
		71:  uint16(161),
		72:  uint16(161),
		73:  uint16(161),
		74:  uint16(161),
		75:  uint16(161),
		76:  uint16(161),
		77:  uint16(161),
		78:  uint16(161),
		79:  uint16(161),
		80:  uint16(161),
		81:  uint16(161),
		82:  uint16(161),
		83:  uint16(161),
		84:  uint16(161),
		85:  uint16(161),
		86:  uint16(161),
		87:  uint16(161),
		88:  uint16(161),
		89:  uint16(161),
		90:  uint16(161),
		91:  uint16(161),
		92:  uint16(161),
		93:  uint16(161),
		94:  uint16(161),
		95:  uint16(161),
		96:  uint16(161),
		97:  uint16(161),
		98:  uint16(161),
		99:  uint16(161),
		100: uint16(161),
		101: uint16(161),
		102: uint16(161),
		103: uint16(161),
		104: uint16(161),
		105: uint16(161),
		106: uint16(161),
		107: uint16(161),
		108: uint16(161),
		109: uint16(161),
		110: uint16(161),
		111: uint16(161),
		125: uint16(3),
		128: uint16(163),
	},
	24: {
		1:   uint16(165),
		9:   uint16(165),
		10:  uint16(165),
		21:  uint16(167),
		31:  uint16(165),
		32:  uint16(165),
		33:  uint16(165),
		34:  uint16(165),
		35:  uint16(165),
		36:  uint16(165),
		38:  uint16(165),
		40:  uint16(165),
		53:  uint16(165),
		54:  uint16(165),
		55:  uint16(165),
		56:  uint16(165),
		57:  uint16(165),
		58:  uint16(165),
		59:  uint16(165),
		60:  uint16(165),
		61:  uint16(165),
		62:  uint16(167),
		63:  uint16(165),
		64:  uint16(165),
		65:  uint16(165),
		66:  uint16(165),
		70:  uint16(167),
		71:  uint16(165),
		72:  uint16(165),
		73:  uint16(165),
		74:  uint16(165),
		75:  uint16(165),
		76:  uint16(165),
		77:  uint16(165),
		78:  uint16(165),
		79:  uint16(165),
		80:  uint16(165),
		81:  uint16(165),
		82:  uint16(165),
		83:  uint16(165),
		84:  uint16(165),
		85:  uint16(165),
		86:  uint16(165),
		87:  uint16(165),
		88:  uint16(165),
		89:  uint16(165),
		90:  uint16(165),
		91:  uint16(165),
		92:  uint16(165),
		93:  uint16(165),
		94:  uint16(165),
		95:  uint16(165),
		96:  uint16(165),
		97:  uint16(165),
		98:  uint16(165),
		99:  uint16(165),
		100: uint16(165),
		101: uint16(165),
		102: uint16(165),
		103: uint16(165),
		104: uint16(165),
		105: uint16(165),
		106: uint16(165),
		107: uint16(165),
		108: uint16(165),
		109: uint16(165),
		110: uint16(165),
		111: uint16(165),
		125: uint16(3),
		128: uint16(167),
	},
	25: {
		1:   uint16(169),
		9:   uint16(169),
		10:  uint16(169),
		21:  uint16(171),
		31:  uint16(169),
		32:  uint16(169),
		33:  uint16(169),
		34:  uint16(169),
		35:  uint16(169),
		36:  uint16(169),
		38:  uint16(169),
		40:  uint16(169),
		53:  uint16(169),
		54:  uint16(169),
		55:  uint16(169),
		56:  uint16(169),
		57:  uint16(169),
		58:  uint16(169),
		59:  uint16(169),
		60:  uint16(169),
		61:  uint16(169),
		62:  uint16(171),
		63:  uint16(169),
		64:  uint16(169),
		65:  uint16(169),
		66:  uint16(169),
		70:  uint16(171),
		71:  uint16(169),
		72:  uint16(169),
		73:  uint16(169),
		74:  uint16(169),
		75:  uint16(169),
		76:  uint16(169),
		77:  uint16(169),
		78:  uint16(169),
		79:  uint16(169),
		80:  uint16(169),
		81:  uint16(169),
		82:  uint16(169),
		83:  uint16(169),
		84:  uint16(169),
		85:  uint16(169),
		86:  uint16(169),
		87:  uint16(169),
		88:  uint16(169),
		89:  uint16(169),
		90:  uint16(169),
		91:  uint16(169),
		92:  uint16(169),
		93:  uint16(169),
		94:  uint16(169),
		95:  uint16(169),
		96:  uint16(169),
		97:  uint16(169),
		98:  uint16(169),
		99:  uint16(169),
		100: uint16(169),
		101: uint16(169),
		102: uint16(169),
		103: uint16(169),
		104: uint16(169),
		105: uint16(169),
		106: uint16(169),
		107: uint16(169),
		108: uint16(169),
		109: uint16(169),
		110: uint16(169),
		111: uint16(169),
		125: uint16(3),
		128: uint16(171),
	},
	26: {
		1:   uint16(173),
		9:   uint16(173),
		10:  uint16(173),
		21:  uint16(175),
		31:  uint16(173),
		32:  uint16(173),
		33:  uint16(173),
		34:  uint16(173),
		35:  uint16(173),
		36:  uint16(173),
		38:  uint16(173),
		40:  uint16(173),
		53:  uint16(173),
		54:  uint16(173),
		55:  uint16(173),
		56:  uint16(173),
		57:  uint16(173),
		58:  uint16(173),
		59:  uint16(173),
		60:  uint16(173),
		61:  uint16(173),
		62:  uint16(175),
		63:  uint16(173),
		64:  uint16(173),
		65:  uint16(173),
		66:  uint16(173),
		70:  uint16(175),
		71:  uint16(173),
		72:  uint16(173),
		73:  uint16(173),
		74:  uint16(173),
		75:  uint16(173),
		76:  uint16(173),
		77:  uint16(173),
		78:  uint16(173),
		79:  uint16(173),
		80:  uint16(173),
		81:  uint16(173),
		82:  uint16(173),
		83:  uint16(173),
		84:  uint16(173),
		85:  uint16(173),
		86:  uint16(173),
		87:  uint16(173),
		88:  uint16(173),
		89:  uint16(173),
		90:  uint16(173),
		91:  uint16(173),
		92:  uint16(173),
		93:  uint16(173),
		94:  uint16(173),
		95:  uint16(173),
		96:  uint16(173),
		97:  uint16(173),
		98:  uint16(173),
		99:  uint16(173),
		100: uint16(173),
		101: uint16(173),
		102: uint16(173),
		103: uint16(173),
		104: uint16(173),
		105: uint16(173),
		106: uint16(173),
		107: uint16(173),
		108: uint16(173),
		109: uint16(173),
		110: uint16(173),
		111: uint16(173),
		125: uint16(3),
		128: uint16(175),
	},
	27: {
		1:   uint16(177),
		9:   uint16(177),
		10:  uint16(177),
		21:  uint16(179),
		31:  uint16(177),
		32:  uint16(177),
		33:  uint16(177),
		34:  uint16(177),
		35:  uint16(177),
		36:  uint16(177),
		38:  uint16(177),
		40:  uint16(177),
		53:  uint16(177),
		54:  uint16(177),
		55:  uint16(177),
		56:  uint16(177),
		57:  uint16(177),
		58:  uint16(177),
		59:  uint16(177),
		60:  uint16(177),
		61:  uint16(177),
		62:  uint16(179),
		63:  uint16(177),
		64:  uint16(177),
		65:  uint16(177),
		66:  uint16(177),
		70:  uint16(179),
		71:  uint16(177),
		72:  uint16(177),
		73:  uint16(177),
		74:  uint16(177),
		75:  uint16(177),
		76:  uint16(177),
		77:  uint16(177),
		78:  uint16(177),
		79:  uint16(177),
		80:  uint16(177),
		81:  uint16(177),
		82:  uint16(177),
		83:  uint16(177),
		84:  uint16(177),
		85:  uint16(177),
		86:  uint16(177),
		87:  uint16(177),
		88:  uint16(177),
		89:  uint16(177),
		90:  uint16(177),
		91:  uint16(177),
		92:  uint16(177),
		93:  uint16(177),
		94:  uint16(177),
		95:  uint16(177),
		96:  uint16(177),
		97:  uint16(177),
		98:  uint16(177),
		99:  uint16(177),
		100: uint16(177),
		101: uint16(177),
		102: uint16(177),
		103: uint16(177),
		104: uint16(177),
		105: uint16(177),
		106: uint16(177),
		107: uint16(177),
		108: uint16(177),
		109: uint16(177),
		110: uint16(177),
		111: uint16(177),
		125: uint16(3),
		128: uint16(179),
	},
	28: {
		1:   uint16(181),
		9:   uint16(181),
		10:  uint16(181),
		21:  uint16(183),
		31:  uint16(181),
		32:  uint16(181),
		33:  uint16(181),
		34:  uint16(181),
		35:  uint16(185),
		36:  uint16(181),
		38:  uint16(181),
		40:  uint16(181),
		53:  uint16(181),
		54:  uint16(181),
		55:  uint16(181),
		56:  uint16(181),
		57:  uint16(181),
		58:  uint16(181),
		59:  uint16(181),
		60:  uint16(181),
		61:  uint16(181),
		62:  uint16(187),
		63:  uint16(181),
		64:  uint16(181),
		65:  uint16(181),
		66:  uint16(181),
		70:  uint16(189),
		71:  uint16(181),
		72:  uint16(181),
		73:  uint16(181),
		74:  uint16(181),
		75:  uint16(181),
		76:  uint16(181),
		77:  uint16(181),
		78:  uint16(181),
		79:  uint16(181),
		80:  uint16(181),
		81:  uint16(181),
		82:  uint16(181),
		83:  uint16(181),
		84:  uint16(181),
		85:  uint16(181),
		86:  uint16(181),
		87:  uint16(181),
		88:  uint16(181),
		89:  uint16(181),
		90:  uint16(181),
		91:  uint16(181),
		92:  uint16(181),
		93:  uint16(181),
		94:  uint16(181),
		95:  uint16(181),
		96:  uint16(181),
		97:  uint16(181),
		98:  uint16(181),
		99:  uint16(181),
		100: uint16(181),
		101: uint16(181),
		102: uint16(181),
		103: uint16(181),
		104: uint16(181),
		105: uint16(181),
		106: uint16(181),
		107: uint16(181),
		108: uint16(181),
		109: uint16(181),
		110: uint16(181),
		111: uint16(181),
		125: uint16(3),
		128: uint16(191),
	},
	29: {
		1:   uint16(193),
		9:   uint16(193),
		10:  uint16(193),
		21:  uint16(195),
		31:  uint16(193),
		32:  uint16(193),
		33:  uint16(193),
		34:  uint16(193),
		35:  uint16(193),
		36:  uint16(193),
		38:  uint16(193),
		40:  uint16(193),
		53:  uint16(193),
		54:  uint16(193),
		55:  uint16(193),
		56:  uint16(193),
		57:  uint16(193),
		58:  uint16(193),
		59:  uint16(193),
		60:  uint16(193),
		61:  uint16(193),
		62:  uint16(195),
		63:  uint16(193),
		64:  uint16(193),
		65:  uint16(193),
		66:  uint16(193),
		70:  uint16(195),
		71:  uint16(193),
		72:  uint16(193),
		73:  uint16(193),
		74:  uint16(193),
		75:  uint16(193),
		76:  uint16(193),
		77:  uint16(193),
		78:  uint16(193),
		79:  uint16(193),
		80:  uint16(193),
		81:  uint16(193),
		82:  uint16(193),
		83:  uint16(193),
		84:  uint16(193),
		85:  uint16(193),
		86:  uint16(193),
		87:  uint16(193),
		88:  uint16(193),
		89:  uint16(193),
		90:  uint16(193),
		91:  uint16(193),
		92:  uint16(193),
		93:  uint16(193),
		94:  uint16(193),
		95:  uint16(193),
		96:  uint16(193),
		97:  uint16(193),
		98:  uint16(193),
		99:  uint16(193),
		100: uint16(193),
		101: uint16(193),
		102: uint16(193),
		103: uint16(193),
		104: uint16(193),
		105: uint16(193),
		106: uint16(193),
		107: uint16(193),
		108: uint16(193),
		109: uint16(193),
		110: uint16(193),
		111: uint16(193),
		125: uint16(3),
		128: uint16(195),
	},
	30: {
		1:   uint16(197),
		9:   uint16(197),
		10:  uint16(197),
		21:  uint16(199),
		31:  uint16(197),
		32:  uint16(197),
		33:  uint16(197),
		34:  uint16(197),
		35:  uint16(197),
		36:  uint16(197),
		38:  uint16(197),
		40:  uint16(197),
		53:  uint16(197),
		54:  uint16(197),
		55:  uint16(197),
		56:  uint16(197),
		57:  uint16(197),
		58:  uint16(197),
		59:  uint16(197),
		60:  uint16(197),
		61:  uint16(197),
		62:  uint16(199),
		63:  uint16(197),
		64:  uint16(197),
		65:  uint16(197),
		66:  uint16(197),
		70:  uint16(199),
		71:  uint16(197),
		72:  uint16(197),
		73:  uint16(197),
		74:  uint16(197),
		75:  uint16(197),
		76:  uint16(197),
		77:  uint16(197),
		78:  uint16(197),
		79:  uint16(197),
		80:  uint16(197),
		81:  uint16(197),
		82:  uint16(197),
		83:  uint16(197),
		84:  uint16(197),
		85:  uint16(197),
		86:  uint16(197),
		87:  uint16(197),
		88:  uint16(197),
		89:  uint16(197),
		90:  uint16(197),
		91:  uint16(197),
		92:  uint16(197),
		93:  uint16(197),
		94:  uint16(197),
		95:  uint16(197),
		96:  uint16(197),
		97:  uint16(197),
		98:  uint16(197),
		99:  uint16(197),
		100: uint16(197),
		101: uint16(197),
		102: uint16(197),
		103: uint16(197),
		104: uint16(197),
		105: uint16(197),
		106: uint16(197),
		107: uint16(197),
		108: uint16(197),
		109: uint16(197),
		110: uint16(197),
		111: uint16(197),
		125: uint16(3),
		128: uint16(199),
	},
	31: {
		1:   uint16(201),
		9:   uint16(201),
		10:  uint16(201),
		21:  uint16(203),
		31:  uint16(201),
		32:  uint16(201),
		33:  uint16(201),
		34:  uint16(201),
		35:  uint16(201),
		36:  uint16(201),
		38:  uint16(201),
		40:  uint16(201),
		53:  uint16(201),
		54:  uint16(201),
		55:  uint16(201),
		56:  uint16(201),
		57:  uint16(201),
		58:  uint16(201),
		59:  uint16(201),
		60:  uint16(201),
		61:  uint16(201),
		62:  uint16(203),
		63:  uint16(201),
		64:  uint16(201),
		65:  uint16(201),
		66:  uint16(201),
		70:  uint16(203),
		71:  uint16(201),
		72:  uint16(201),
		73:  uint16(201),
		74:  uint16(201),
		75:  uint16(201),
		76:  uint16(201),
		77:  uint16(201),
		78:  uint16(201),
		79:  uint16(201),
		80:  uint16(201),
		81:  uint16(201),
		82:  uint16(201),
		83:  uint16(201),
		84:  uint16(201),
		85:  uint16(201),
		86:  uint16(201),
		87:  uint16(201),
		88:  uint16(201),
		89:  uint16(201),
		90:  uint16(201),
		91:  uint16(201),
		92:  uint16(201),
		93:  uint16(201),
		94:  uint16(201),
		95:  uint16(201),
		96:  uint16(201),
		97:  uint16(201),
		98:  uint16(201),
		99:  uint16(201),
		100: uint16(201),
		101: uint16(201),
		102: uint16(201),
		103: uint16(201),
		104: uint16(201),
		105: uint16(201),
		106: uint16(201),
		107: uint16(201),
		108: uint16(201),
		109: uint16(201),
		110: uint16(201),
		111: uint16(201),
		125: uint16(3),
		128: uint16(203),
	},
	32: {
		1:   uint16(205),
		9:   uint16(205),
		10:  uint16(205),
		21:  uint16(183),
		31:  uint16(205),
		32:  uint16(205),
		33:  uint16(205),
		34:  uint16(205),
		36:  uint16(205),
		38:  uint16(205),
		40:  uint16(205),
		53:  uint16(205),
		54:  uint16(205),
		55:  uint16(205),
		56:  uint16(205),
		57:  uint16(205),
		58:  uint16(205),
		59:  uint16(205),
		60:  uint16(205),
		61:  uint16(205),
		62:  uint16(207),
		63:  uint16(205),
		64:  uint16(205),
		65:  uint16(205),
		66:  uint16(205),
		70:  uint16(189),
		71:  uint16(205),
		72:  uint16(205),
		73:  uint16(205),
		74:  uint16(205),
		75:  uint16(205),
		76:  uint16(205),
		77:  uint16(205),
		78:  uint16(205),
		79:  uint16(205),
		80:  uint16(205),
		81:  uint16(205),
		82:  uint16(205),
		83:  uint16(205),
		84:  uint16(205),
		85:  uint16(205),
		86:  uint16(205),
		87:  uint16(205),
		88:  uint16(205),
		89:  uint16(205),
		90:  uint16(205),
		91:  uint16(205),
		92:  uint16(205),
		93:  uint16(205),
		94:  uint16(205),
		95:  uint16(205),
		96:  uint16(205),
		97:  uint16(205),
		98:  uint16(205),
		99:  uint16(205),
		100: uint16(205),
		101: uint16(205),
		102: uint16(205),
		103: uint16(205),
		104: uint16(205),
		105: uint16(205),
		106: uint16(205),
		107: uint16(205),
		108: uint16(205),
		109: uint16(205),
		110: uint16(205),
		111: uint16(205),
		125: uint16(3),
		128: uint16(209),
	},
	33: {
		1:   uint16(211),
		6:   uint16(211),
		7:   uint16(211),
		9:   uint16(211),
		10:  uint16(211),
		24:  uint16(211),
		26:  uint16(211),
		31:  uint16(211),
		32:  uint16(211),
		33:  uint16(211),
		34:  uint16(211),
		36:  uint16(211),
		38:  uint16(211),
		40:  uint16(211),
		53:  uint16(211),
		55:  uint16(211),
		56:  uint16(211),
		57:  uint16(211),
		58:  uint16(211),
		59:  uint16(211),
		60:  uint16(211),
		61:  uint16(211),
		63:  uint16(211),
		64:  uint16(211),
		65:  uint16(211),
		66:  uint16(211),
		71:  uint16(211),
		72:  uint16(211),
		73:  uint16(211),
		74:  uint16(211),
		75:  uint16(211),
		76:  uint16(211),
		77:  uint16(211),
		78:  uint16(211),
		79:  uint16(211),
		80:  uint16(211),
		81:  uint16(211),
		82:  uint16(211),
		83:  uint16(211),
		84:  uint16(211),
		85:  uint16(211),
		86:  uint16(211),
		87:  uint16(211),
		88:  uint16(211),
		89:  uint16(211),
		90:  uint16(211),
		91:  uint16(211),
		92:  uint16(211),
		93:  uint16(211),
		94:  uint16(211),
		95:  uint16(211),
		96:  uint16(211),
		97:  uint16(211),
		98:  uint16(211),
		99:  uint16(211),
		100: uint16(211),
		101: uint16(211),
		102: uint16(211),
		103: uint16(211),
		104: uint16(211),
		105: uint16(211),
		106: uint16(211),
		107: uint16(211),
		108: uint16(211),
		109: uint16(211),
		110: uint16(211),
		111: uint16(211),
		125: uint16(3),
		128: uint16(213),
	},
	34: {
		1:   uint16(215),
		9:   uint16(215),
		10:  uint16(215),
		12:  uint16(217),
		21:  uint16(219),
		31:  uint16(215),
		32:  uint16(215),
		33:  uint16(215),
		34:  uint16(215),
		36:  uint16(215),
		38:  uint16(215),
		40:  uint16(215),
		53:  uint16(215),
		54:  uint16(215),
		55:  uint16(215),
		56:  uint16(215),
		57:  uint16(215),
		58:  uint16(215),
		59:  uint16(215),
		60:  uint16(215),
		61:  uint16(215),
		62:  uint16(219),
		63:  uint16(215),
		64:  uint16(215),
		65:  uint16(215),
		66:  uint16(215),
		71:  uint16(215),
		72:  uint16(215),
		73:  uint16(215),
		74:  uint16(215),
		75:  uint16(215),
		76:  uint16(215),
		77:  uint16(215),
		78:  uint16(215),
		79:  uint16(215),
		80:  uint16(215),
		81:  uint16(215),
		82:  uint16(215),
		83:  uint16(215),
		84:  uint16(215),
		85:  uint16(215),
		86:  uint16(215),
		87:  uint16(215),
		88:  uint16(215),
		89:  uint16(215),
		90:  uint16(215),
		91:  uint16(215),
		92:  uint16(215),
		93:  uint16(215),
		94:  uint16(215),
		95:  uint16(215),
		96:  uint16(215),
		97:  uint16(215),
		98:  uint16(215),
		99:  uint16(215),
		100: uint16(215),
		101: uint16(215),
		102: uint16(215),
		103: uint16(215),
		104: uint16(215),
		105: uint16(215),
		106: uint16(215),
		107: uint16(215),
		108: uint16(215),
		109: uint16(215),
		110: uint16(215),
		111: uint16(215),
		125: uint16(3),
		128: uint16(219),
	},
	35: {
		1:   uint16(221),
		9:   uint16(221),
		10:  uint16(221),
		21:  uint16(183),
		31:  uint16(221),
		32:  uint16(221),
		33:  uint16(221),
		34:  uint16(221),
		36:  uint16(221),
		38:  uint16(221),
		40:  uint16(221),
		53:  uint16(221),
		54:  uint16(221),
		55:  uint16(221),
		56:  uint16(221),
		57:  uint16(221),
		58:  uint16(221),
		59:  uint16(221),
		60:  uint16(221),
		61:  uint16(221),
		62:  uint16(223),
		63:  uint16(221),
		64:  uint16(221),
		65:  uint16(221),
		66:  uint16(221),
		70:  uint16(189),
		71:  uint16(221),
		72:  uint16(221),
		73:  uint16(221),
		74:  uint16(221),
		75:  uint16(221),
		76:  uint16(221),
		77:  uint16(221),
		78:  uint16(221),
		79:  uint16(221),
		80:  uint16(221),
		81:  uint16(221),
		82:  uint16(221),
		83:  uint16(221),
		84:  uint16(221),
		85:  uint16(221),
		86:  uint16(221),
		87:  uint16(221),
		88:  uint16(221),
		89:  uint16(221),
		90:  uint16(221),
		91:  uint16(221),
		92:  uint16(221),
		93:  uint16(221),
		94:  uint16(221),
		95:  uint16(221),
		96:  uint16(221),
		97:  uint16(221),
		98:  uint16(221),
		99:  uint16(221),
		100: uint16(221),
		101: uint16(221),
		102: uint16(221),
		103: uint16(221),
		104: uint16(221),
		105: uint16(221),
		106: uint16(221),
		107: uint16(221),
		108: uint16(221),
		109: uint16(221),
		110: uint16(221),
		111: uint16(221),
		125: uint16(3),
		128: uint16(225),
	},
	36: {
		1:   uint16(227),
		6:   uint16(227),
		7:   uint16(227),
		9:   uint16(227),
		10:  uint16(227),
		24:  uint16(227),
		26:  uint16(227),
		31:  uint16(227),
		32:  uint16(227),
		33:  uint16(227),
		34:  uint16(227),
		36:  uint16(227),
		38:  uint16(227),
		40:  uint16(227),
		53:  uint16(227),
		55:  uint16(227),
		56:  uint16(227),
		57:  uint16(227),
		58:  uint16(227),
		59:  uint16(227),
		60:  uint16(227),
		61:  uint16(227),
		63:  uint16(227),
		64:  uint16(227),
		65:  uint16(227),
		66:  uint16(227),
		71:  uint16(227),
		72:  uint16(227),
		73:  uint16(227),
		74:  uint16(227),
		75:  uint16(227),
		76:  uint16(227),
		77:  uint16(227),
		78:  uint16(227),
		79:  uint16(227),
		80:  uint16(227),
		81:  uint16(227),
		82:  uint16(227),
		83:  uint16(227),
		84:  uint16(227),
		85:  uint16(227),
		86:  uint16(227),
		87:  uint16(227),
		88:  uint16(227),
		89:  uint16(227),
		90:  uint16(227),
		91:  uint16(227),
		92:  uint16(227),
		93:  uint16(227),
		94:  uint16(227),
		95:  uint16(227),
		96:  uint16(227),
		97:  uint16(227),
		98:  uint16(227),
		99:  uint16(227),
		100: uint16(227),
		101: uint16(227),
		102: uint16(227),
		103: uint16(227),
		104: uint16(227),
		105: uint16(227),
		106: uint16(227),
		107: uint16(227),
		108: uint16(227),
		109: uint16(227),
		110: uint16(227),
		111: uint16(227),
		125: uint16(3),
		128: uint16(229),
	},
	37: {
		1:   uint16(231),
		9:   uint16(231),
		10:  uint16(231),
		21:  uint16(183),
		31:  uint16(231),
		32:  uint16(231),
		33:  uint16(231),
		34:  uint16(231),
		36:  uint16(231),
		38:  uint16(231),
		40:  uint16(231),
		53:  uint16(231),
		54:  uint16(231),
		55:  uint16(231),
		56:  uint16(231),
		57:  uint16(231),
		58:  uint16(231),
		59:  uint16(231),
		60:  uint16(231),
		61:  uint16(231),
		62:  uint16(233),
		63:  uint16(231),
		64:  uint16(231),
		65:  uint16(231),
		66:  uint16(231),
		70:  uint16(189),
		71:  uint16(231),
		72:  uint16(231),
		73:  uint16(231),
		74:  uint16(231),
		75:  uint16(231),
		76:  uint16(231),
		77:  uint16(231),
		78:  uint16(231),
		79:  uint16(231),
		80:  uint16(231),
		81:  uint16(231),
		82:  uint16(231),
		83:  uint16(231),
		84:  uint16(231),
		85:  uint16(231),
		86:  uint16(231),
		87:  uint16(231),
		88:  uint16(231),
		89:  uint16(231),
		90:  uint16(231),
		91:  uint16(231),
		92:  uint16(231),
		93:  uint16(231),
		94:  uint16(231),
		95:  uint16(231),
		96:  uint16(231),
		97:  uint16(231),
		98:  uint16(231),
		99:  uint16(231),
		100: uint16(231),
		101: uint16(231),
		102: uint16(231),
		103: uint16(231),
		104: uint16(231),
		105: uint16(231),
		106: uint16(231),
		107: uint16(231),
		108: uint16(231),
		109: uint16(231),
		110: uint16(231),
		111: uint16(231),
		125: uint16(3),
		128: uint16(235),
	},
	38: {
		1:   uint16(237),
		9:   uint16(237),
		10:  uint16(237),
		21:  uint16(183),
		31:  uint16(237),
		32:  uint16(237),
		33:  uint16(237),
		34:  uint16(237),
		36:  uint16(237),
		38:  uint16(237),
		40:  uint16(237),
		53:  uint16(237),
		54:  uint16(237),
		55:  uint16(237),
		56:  uint16(237),
		57:  uint16(237),
		58:  uint16(237),
		59:  uint16(237),
		60:  uint16(237),
		61:  uint16(237),
		62:  uint16(239),
		63:  uint16(237),
		64:  uint16(237),
		65:  uint16(237),
		66:  uint16(237),
		70:  uint16(189),
		71:  uint16(237),
		72:  uint16(237),
		73:  uint16(237),
		74:  uint16(237),
		75:  uint16(237),
		76:  uint16(237),
		77:  uint16(237),
		78:  uint16(237),
		79:  uint16(237),
		80:  uint16(237),
		81:  uint16(237),
		82:  uint16(237),
		83:  uint16(237),
		84:  uint16(237),
		85:  uint16(237),
		86:  uint16(237),
		87:  uint16(237),
		88:  uint16(237),
		89:  uint16(237),
		90:  uint16(237),
		91:  uint16(237),
		92:  uint16(237),
		93:  uint16(237),
		94:  uint16(237),
		95:  uint16(237),
		96:  uint16(237),
		97:  uint16(237),
		98:  uint16(237),
		99:  uint16(237),
		100: uint16(237),
		101: uint16(237),
		102: uint16(237),
		103: uint16(237),
		104: uint16(237),
		105: uint16(237),
		106: uint16(237),
		107: uint16(237),
		108: uint16(237),
		109: uint16(237),
		110: uint16(237),
		111: uint16(237),
		125: uint16(3),
		128: uint16(241),
	},
	39: {
		1:   uint16(215),
		9:   uint16(215),
		10:  uint16(215),
		12:  uint16(243),
		21:  uint16(219),
		31:  uint16(215),
		32:  uint16(215),
		33:  uint16(215),
		34:  uint16(215),
		36:  uint16(215),
		38:  uint16(215),
		40:  uint16(215),
		53:  uint16(215),
		54:  uint16(215),
		55:  uint16(215),
		56:  uint16(215),
		57:  uint16(215),
		58:  uint16(215),
		59:  uint16(215),
		60:  uint16(215),
		61:  uint16(215),
		62:  uint16(219),
		63:  uint16(215),
		64:  uint16(215),
		65:  uint16(215),
		66:  uint16(215),
		71:  uint16(215),
		72:  uint16(215),
		73:  uint16(215),
		74:  uint16(215),
		75:  uint16(215),
		76:  uint16(215),
		77:  uint16(215),
		78:  uint16(215),
		79:  uint16(215),
		80:  uint16(215),
		81:  uint16(215),
		82:  uint16(215),
		83:  uint16(215),
		84:  uint16(215),
		85:  uint16(215),
		86:  uint16(215),
		87:  uint16(215),
		88:  uint16(215),
		89:  uint16(215),
		90:  uint16(215),
		91:  uint16(215),
		92:  uint16(215),
		93:  uint16(215),
		94:  uint16(215),
		95:  uint16(215),
		96:  uint16(215),
		97:  uint16(215),
		98:  uint16(215),
		99:  uint16(215),
		100: uint16(215),
		101: uint16(215),
		102: uint16(215),
		103: uint16(215),
		104: uint16(215),
		105: uint16(215),
		106: uint16(215),
		107: uint16(215),
		108: uint16(215),
		109: uint16(215),
		110: uint16(215),
		111: uint16(215),
		125: uint16(3),
		128: uint16(219),
	},
	40: {
		1:   uint16(245),
		9:   uint16(245),
		10:  uint16(245),
		12:  uint16(247),
		21:  uint16(249),
		31:  uint16(245),
		32:  uint16(245),
		33:  uint16(245),
		34:  uint16(245),
		36:  uint16(245),
		38:  uint16(245),
		40:  uint16(245),
		53:  uint16(245),
		54:  uint16(245),
		55:  uint16(245),
		56:  uint16(245),
		57:  uint16(245),
		58:  uint16(245),
		59:  uint16(245),
		60:  uint16(245),
		61:  uint16(245),
		62:  uint16(249),
		63:  uint16(245),
		64:  uint16(245),
		65:  uint16(245),
		66:  uint16(245),
		71:  uint16(245),
		72:  uint16(245),
		73:  uint16(245),
		74:  uint16(245),
		75:  uint16(245),
		76:  uint16(245),
		77:  uint16(245),
		78:  uint16(245),
		79:  uint16(245),
		80:  uint16(245),
		81:  uint16(245),
		82:  uint16(245),
		83:  uint16(245),
		84:  uint16(245),
		85:  uint16(245),
		86:  uint16(245),
		87:  uint16(245),
		88:  uint16(245),
		89:  uint16(245),
		90:  uint16(245),
		91:  uint16(245),
		92:  uint16(245),
		93:  uint16(245),
		94:  uint16(245),
		95:  uint16(245),
		96:  uint16(245),
		97:  uint16(245),
		98:  uint16(245),
		99:  uint16(245),
		100: uint16(245),
		101: uint16(245),
		102: uint16(245),
		103: uint16(245),
		104: uint16(245),
		105: uint16(245),
		106: uint16(245),
		107: uint16(245),
		108: uint16(245),
		109: uint16(245),
		110: uint16(245),
		111: uint16(245),
		125: uint16(3),
		128: uint16(249),
	},
	41: {
		1:   uint16(251),
		6:   uint16(251),
		7:   uint16(251),
		9:   uint16(251),
		10:  uint16(251),
		24:  uint16(251),
		26:  uint16(251),
		31:  uint16(251),
		32:  uint16(251),
		33:  uint16(251),
		34:  uint16(251),
		36:  uint16(251),
		38:  uint16(251),
		40:  uint16(251),
		53:  uint16(251),
		55:  uint16(251),
		56:  uint16(251),
		57:  uint16(251),
		58:  uint16(251),
		59:  uint16(251),
		60:  uint16(251),
		61:  uint16(251),
		63:  uint16(251),
		64:  uint16(251),
		65:  uint16(251),
		66:  uint16(251),
		71:  uint16(251),
		72:  uint16(251),
		73:  uint16(251),
		74:  uint16(251),
		75:  uint16(251),
		76:  uint16(251),
		77:  uint16(251),
		78:  uint16(251),
		79:  uint16(251),
		80:  uint16(251),
		81:  uint16(251),
		82:  uint16(251),
		83:  uint16(251),
		84:  uint16(251),
		85:  uint16(251),
		86:  uint16(251),
		87:  uint16(251),
		88:  uint16(251),
		89:  uint16(251),
		90:  uint16(251),
		91:  uint16(251),
		92:  uint16(251),
		93:  uint16(251),
		94:  uint16(251),
		95:  uint16(251),
		96:  uint16(251),
		97:  uint16(251),
		98:  uint16(251),
		99:  uint16(251),
		100: uint16(251),
		101: uint16(251),
		102: uint16(251),
		103: uint16(251),
		104: uint16(251),
		105: uint16(251),
		106: uint16(251),
		107: uint16(251),
		108: uint16(251),
		109: uint16(251),
		110: uint16(251),
		111: uint16(251),
		125: uint16(3),
		128: uint16(253),
	},
	42: {
		1:   uint16(255),
		9:   uint16(255),
		10:  uint16(255),
		21:  uint16(257),
		31:  uint16(255),
		32:  uint16(255),
		33:  uint16(255),
		34:  uint16(255),
		36:  uint16(255),
		38:  uint16(255),
		40:  uint16(255),
		53:  uint16(255),
		54:  uint16(255),
		55:  uint16(255),
		56:  uint16(255),
		57:  uint16(255),
		58:  uint16(255),
		59:  uint16(255),
		60:  uint16(255),
		61:  uint16(255),
		62:  uint16(257),
		63:  uint16(255),
		64:  uint16(255),
		65:  uint16(255),
		66:  uint16(255),
		71:  uint16(255),
		72:  uint16(255),
		73:  uint16(255),
		74:  uint16(255),
		75:  uint16(255),
		76:  uint16(255),
		77:  uint16(255),
		78:  uint16(255),
		79:  uint16(255),
		80:  uint16(255),
		81:  uint16(255),
		82:  uint16(255),
		83:  uint16(255),
		84:  uint16(255),
		85:  uint16(255),
		86:  uint16(255),
		87:  uint16(255),
		88:  uint16(255),
		89:  uint16(255),
		90:  uint16(255),
		91:  uint16(255),
		92:  uint16(255),
		93:  uint16(255),
		94:  uint16(255),
		95:  uint16(255),
		96:  uint16(255),
		97:  uint16(255),
		98:  uint16(255),
		99:  uint16(255),
		100: uint16(255),
		101: uint16(255),
		102: uint16(255),
		103: uint16(255),
		104: uint16(255),
		105: uint16(255),
		106: uint16(255),
		107: uint16(255),
		108: uint16(255),
		109: uint16(255),
		110: uint16(255),
		111: uint16(255),
		125: uint16(3),
		128: uint16(257),
	},
	43: {
		1:   uint16(259),
		9:   uint16(259),
		10:  uint16(259),
		21:  uint16(261),
		31:  uint16(259),
		32:  uint16(259),
		33:  uint16(259),
		34:  uint16(259),
		36:  uint16(259),
		38:  uint16(259),
		40:  uint16(259),
		53:  uint16(259),
		54:  uint16(259),
		55:  uint16(259),
		56:  uint16(259),
		57:  uint16(259),
		58:  uint16(259),
		59:  uint16(259),
		60:  uint16(259),
		61:  uint16(259),
		62:  uint16(261),
		63:  uint16(259),
		64:  uint16(259),
		65:  uint16(259),
		66:  uint16(259),
		71:  uint16(259),
		72:  uint16(259),
		73:  uint16(259),
		74:  uint16(259),
		75:  uint16(259),
		76:  uint16(259),
		77:  uint16(259),
		78:  uint16(259),
		79:  uint16(259),
		80:  uint16(259),
		81:  uint16(259),
		82:  uint16(259),
		83:  uint16(259),
		84:  uint16(259),
		85:  uint16(259),
		86:  uint16(259),
		87:  uint16(259),
		88:  uint16(259),
		89:  uint16(259),
		90:  uint16(259),
		91:  uint16(259),
		92:  uint16(259),
		93:  uint16(259),
		94:  uint16(259),
		95:  uint16(259),
		96:  uint16(259),
		97:  uint16(259),
		98:  uint16(259),
		99:  uint16(259),
		100: uint16(259),
		101: uint16(259),
		102: uint16(259),
		103: uint16(259),
		104: uint16(259),
		105: uint16(259),
		106: uint16(259),
		107: uint16(259),
		108: uint16(259),
		109: uint16(259),
		110: uint16(259),
		111: uint16(259),
		125: uint16(3),
		128: uint16(261),
	},
	44: {
		1:   uint16(263),
		9:   uint16(263),
		10:  uint16(263),
		21:  uint16(265),
		31:  uint16(263),
		32:  uint16(263),
		33:  uint16(263),
		34:  uint16(263),
		36:  uint16(263),
		38:  uint16(263),
		40:  uint16(263),
		53:  uint16(263),
		54:  uint16(263),
		55:  uint16(263),
		56:  uint16(263),
		57:  uint16(263),
		58:  uint16(263),
		59:  uint16(263),
		60:  uint16(263),
		61:  uint16(263),
		62:  uint16(267),
		63:  uint16(263),
		64:  uint16(263),
		65:  uint16(263),
		66:  uint16(263),
		71:  uint16(263),
		72:  uint16(263),
		73:  uint16(263),
		74:  uint16(263),
		75:  uint16(263),
		76:  uint16(263),
		77:  uint16(263),
		78:  uint16(263),
		79:  uint16(263),
		80:  uint16(263),
		81:  uint16(263),
		82:  uint16(263),
		83:  uint16(263),
		84:  uint16(263),
		85:  uint16(263),
		86:  uint16(263),
		87:  uint16(263),
		88:  uint16(263),
		89:  uint16(263),
		90:  uint16(263),
		91:  uint16(263),
		92:  uint16(263),
		93:  uint16(263),
		94:  uint16(263),
		95:  uint16(263),
		96:  uint16(263),
		97:  uint16(263),
		98:  uint16(263),
		99:  uint16(263),
		100: uint16(263),
		101: uint16(263),
		102: uint16(263),
		103: uint16(263),
		104: uint16(263),
		105: uint16(263),
		106: uint16(263),
		107: uint16(263),
		108: uint16(263),
		109: uint16(263),
		110: uint16(263),
		111: uint16(263),
		125: uint16(3),
		128: uint16(269),
	},
	45: {
		1:   uint16(271),
		9:   uint16(271),
		10:  uint16(271),
		21:  uint16(265),
		31:  uint16(271),
		32:  uint16(271),
		33:  uint16(271),
		34:  uint16(271),
		36:  uint16(271),
		38:  uint16(271),
		40:  uint16(271),
		53:  uint16(271),
		54:  uint16(271),
		55:  uint16(271),
		56:  uint16(271),
		57:  uint16(271),
		58:  uint16(271),
		59:  uint16(271),
		60:  uint16(271),
		61:  uint16(271),
		62:  uint16(273),
		63:  uint16(271),
		64:  uint16(271),
		65:  uint16(271),
		66:  uint16(271),
		71:  uint16(271),
		72:  uint16(271),
		73:  uint16(271),
		74:  uint16(271),
		75:  uint16(271),
		76:  uint16(271),
		77:  uint16(271),
		78:  uint16(271),
		79:  uint16(271),
		80:  uint16(271),
		81:  uint16(271),
		82:  uint16(271),
		83:  uint16(271),
		84:  uint16(271),
		85:  uint16(271),
		86:  uint16(271),
		87:  uint16(271),
		88:  uint16(271),
		89:  uint16(271),
		90:  uint16(271),
		91:  uint16(271),
		92:  uint16(271),
		93:  uint16(271),
		94:  uint16(271),
		95:  uint16(271),
		96:  uint16(271),
		97:  uint16(271),
		98:  uint16(271),
		99:  uint16(271),
		100: uint16(271),
		101: uint16(271),
		102: uint16(271),
		103: uint16(271),
		104: uint16(271),
		105: uint16(271),
		106: uint16(271),
		107: uint16(271),
		108: uint16(271),
		109: uint16(271),
		110: uint16(271),
		111: uint16(271),
		125: uint16(3),
		128: uint16(275),
	},
	46: {
		1:   uint16(277),
		9:   uint16(277),
		10:  uint16(277),
		21:  uint16(265),
		31:  uint16(277),
		32:  uint16(277),
		33:  uint16(277),
		34:  uint16(277),
		36:  uint16(277),
		38:  uint16(277),
		40:  uint16(277),
		53:  uint16(277),
		54:  uint16(277),
		55:  uint16(277),
		56:  uint16(277),
		57:  uint16(277),
		58:  uint16(277),
		59:  uint16(277),
		60:  uint16(277),
		61:  uint16(277),
		62:  uint16(279),
		63:  uint16(277),
		64:  uint16(277),
		65:  uint16(277),
		66:  uint16(277),
		71:  uint16(277),
		72:  uint16(277),
		73:  uint16(277),
		74:  uint16(277),
		75:  uint16(277),
		76:  uint16(277),
		77:  uint16(277),
		78:  uint16(277),
		79:  uint16(277),
		80:  uint16(277),
		81:  uint16(277),
		82:  uint16(277),
		83:  uint16(277),
		84:  uint16(277),
		85:  uint16(277),
		86:  uint16(277),
		87:  uint16(277),
		88:  uint16(277),
		89:  uint16(277),
		90:  uint16(277),
		91:  uint16(277),
		92:  uint16(277),
		93:  uint16(277),
		94:  uint16(277),
		95:  uint16(277),
		96:  uint16(277),
		97:  uint16(277),
		98:  uint16(277),
		99:  uint16(277),
		100: uint16(277),
		101: uint16(277),
		102: uint16(277),
		103: uint16(277),
		104: uint16(277),
		105: uint16(277),
		106: uint16(277),
		107: uint16(277),
		108: uint16(277),
		109: uint16(277),
		110: uint16(277),
		111: uint16(277),
		125: uint16(3),
		128: uint16(281),
	},
	47: {
		1:   uint16(245),
		9:   uint16(245),
		10:  uint16(245),
		21:  uint16(249),
		31:  uint16(245),
		32:  uint16(245),
		33:  uint16(245),
		34:  uint16(245),
		36:  uint16(245),
		38:  uint16(245),
		40:  uint16(245),
		53:  uint16(245),
		54:  uint16(245),
		55:  uint16(245),
		56:  uint16(245),
		57:  uint16(245),
		58:  uint16(245),
		59:  uint16(245),
		60:  uint16(245),
		61:  uint16(245),
		62:  uint16(249),
		63:  uint16(245),
		64:  uint16(245),
		65:  uint16(245),
		66:  uint16(245),
		71:  uint16(245),
		72:  uint16(245),
		73:  uint16(245),
		74:  uint16(245),
		75:  uint16(245),
		76:  uint16(245),
		77:  uint16(245),
		78:  uint16(245),
		79:  uint16(245),
		80:  uint16(245),
		81:  uint16(245),
		82:  uint16(245),
		83:  uint16(245),
		84:  uint16(245),
		85:  uint16(245),
		86:  uint16(245),
		87:  uint16(245),
		88:  uint16(245),
		89:  uint16(245),
		90:  uint16(245),
		91:  uint16(245),
		92:  uint16(245),
		93:  uint16(245),
		94:  uint16(245),
		95:  uint16(245),
		96:  uint16(245),
		97:  uint16(245),
		98:  uint16(245),
		99:  uint16(245),
		100: uint16(245),
		101: uint16(245),
		102: uint16(245),
		103: uint16(245),
		104: uint16(245),
		105: uint16(245),
		106: uint16(245),
		107: uint16(245),
		108: uint16(245),
		109: uint16(245),
		110: uint16(245),
		111: uint16(245),
		125: uint16(3),
		128: uint16(249),
	},
	48: {
		1:   uint16(283),
		9:   uint16(283),
		10:  uint16(283),
		21:  uint16(285),
		31:  uint16(283),
		32:  uint16(283),
		33:  uint16(283),
		34:  uint16(283),
		36:  uint16(283),
		38:  uint16(283),
		40:  uint16(283),
		53:  uint16(283),
		54:  uint16(283),
		55:  uint16(283),
		56:  uint16(283),
		57:  uint16(283),
		58:  uint16(283),
		59:  uint16(283),
		60:  uint16(283),
		61:  uint16(283),
		62:  uint16(285),
		63:  uint16(283),
		64:  uint16(283),
		65:  uint16(283),
		66:  uint16(283),
		71:  uint16(283),
		72:  uint16(283),
		73:  uint16(283),
		74:  uint16(283),
		75:  uint16(283),
		76:  uint16(283),
		77:  uint16(283),
		78:  uint16(283),
		79:  uint16(283),
		80:  uint16(283),
		81:  uint16(283),
		82:  uint16(283),
		83:  uint16(283),
		84:  uint16(283),
		85:  uint16(283),
		86:  uint16(283),
		87:  uint16(283),
		88:  uint16(283),
		89:  uint16(283),
		90:  uint16(283),
		91:  uint16(283),
		92:  uint16(283),
		93:  uint16(283),
		94:  uint16(283),
		95:  uint16(283),
		96:  uint16(283),
		97:  uint16(283),
		98:  uint16(283),
		99:  uint16(283),
		100: uint16(283),
		101: uint16(283),
		102: uint16(283),
		103: uint16(283),
		104: uint16(283),
		105: uint16(283),
		106: uint16(283),
		107: uint16(283),
		108: uint16(283),
		109: uint16(283),
		110: uint16(283),
		111: uint16(283),
		125: uint16(3),
		128: uint16(285),
	},
	49: {
		1:   uint16(287),
		9:   uint16(287),
		10:  uint16(287),
		21:  uint16(289),
		31:  uint16(287),
		32:  uint16(287),
		33:  uint16(287),
		34:  uint16(287),
		36:  uint16(287),
		38:  uint16(287),
		40:  uint16(287),
		53:  uint16(287),
		54:  uint16(287),
		55:  uint16(287),
		56:  uint16(287),
		57:  uint16(287),
		58:  uint16(287),
		59:  uint16(287),
		60:  uint16(287),
		61:  uint16(287),
		62:  uint16(289),
		63:  uint16(287),
		64:  uint16(287),
		65:  uint16(287),
		66:  uint16(287),
		71:  uint16(287),
		72:  uint16(287),
		73:  uint16(287),
		74:  uint16(287),
		75:  uint16(287),
		76:  uint16(287),
		77:  uint16(287),
		78:  uint16(287),
		79:  uint16(287),
		80:  uint16(287),
		81:  uint16(287),
		82:  uint16(287),
		83:  uint16(287),
		84:  uint16(287),
		85:  uint16(287),
		86:  uint16(287),
		87:  uint16(287),
		88:  uint16(287),
		89:  uint16(287),
		90:  uint16(287),
		91:  uint16(287),
		92:  uint16(287),
		93:  uint16(287),
		94:  uint16(287),
		95:  uint16(287),
		96:  uint16(287),
		97:  uint16(287),
		98:  uint16(287),
		99:  uint16(287),
		100: uint16(287),
		101: uint16(287),
		102: uint16(287),
		103: uint16(287),
		104: uint16(287),
		105: uint16(287),
		106: uint16(287),
		107: uint16(287),
		108: uint16(287),
		109: uint16(287),
		110: uint16(287),
		111: uint16(287),
		125: uint16(3),
		128: uint16(289),
	},
	50: {
		1:   uint16(215),
		9:   uint16(215),
		10:  uint16(215),
		21:  uint16(219),
		31:  uint16(215),
		32:  uint16(215),
		33:  uint16(215),
		34:  uint16(215),
		36:  uint16(215),
		38:  uint16(215),
		40:  uint16(215),
		53:  uint16(215),
		54:  uint16(215),
		55:  uint16(215),
		56:  uint16(215),
		57:  uint16(215),
		58:  uint16(215),
		59:  uint16(215),
		60:  uint16(215),
		61:  uint16(215),
		62:  uint16(219),
		63:  uint16(215),
		64:  uint16(215),
		65:  uint16(215),
		66:  uint16(215),
		71:  uint16(215),
		72:  uint16(215),
		73:  uint16(215),
		74:  uint16(215),
		75:  uint16(215),
		76:  uint16(215),
		77:  uint16(215),
		78:  uint16(215),
		79:  uint16(215),
		80:  uint16(215),
		81:  uint16(215),
		82:  uint16(215),
		83:  uint16(215),
		84:  uint16(215),
		85:  uint16(215),
		86:  uint16(215),
		87:  uint16(215),
		88:  uint16(215),
		89:  uint16(215),
		90:  uint16(215),
		91:  uint16(215),
		92:  uint16(215),
		93:  uint16(215),
		94:  uint16(215),
		95:  uint16(215),
		96:  uint16(215),
		97:  uint16(215),
		98:  uint16(215),
		99:  uint16(215),
		100: uint16(215),
		101: uint16(215),
		102: uint16(215),
		103: uint16(215),
		104: uint16(215),
		105: uint16(215),
		106: uint16(215),
		107: uint16(215),
		108: uint16(215),
		109: uint16(215),
		110: uint16(215),
		111: uint16(215),
		125: uint16(3),
		128: uint16(219),
	},
	51: {
		1:   uint16(291),
		9:   uint16(291),
		10:  uint16(291),
		31:  uint16(291),
		32:  uint16(291),
		33:  uint16(291),
		34:  uint16(291),
		36:  uint16(291),
		38:  uint16(291),
		40:  uint16(291),
		53:  uint16(291),
		54:  uint16(291),
		55:  uint16(291),
		56:  uint16(291),
		57:  uint16(291),
		58:  uint16(291),
		59:  uint16(291),
		60:  uint16(291),
		61:  uint16(291),
		62:  uint16(293),
		63:  uint16(291),
		64:  uint16(291),
		65:  uint16(291),
		66:  uint16(291),
		71:  uint16(291),
		72:  uint16(291),
		73:  uint16(291),
		74:  uint16(291),
		75:  uint16(291),
		76:  uint16(291),
		77:  uint16(291),
		78:  uint16(291),
		79:  uint16(291),
		80:  uint16(291),
		81:  uint16(291),
		82:  uint16(291),
		83:  uint16(291),
		84:  uint16(291),
		85:  uint16(291),
		86:  uint16(291),
		87:  uint16(291),
		88:  uint16(291),
		89:  uint16(291),
		90:  uint16(291),
		91:  uint16(291),
		92:  uint16(291),
		93:  uint16(291),
		94:  uint16(291),
		95:  uint16(291),
		96:  uint16(291),
		97:  uint16(291),
		98:  uint16(291),
		99:  uint16(291),
		100: uint16(291),
		101: uint16(291),
		102: uint16(291),
		103: uint16(291),
		104: uint16(291),
		105: uint16(291),
		106: uint16(291),
		107: uint16(291),
		108: uint16(291),
		109: uint16(291),
		110: uint16(291),
		111: uint16(291),
		125: uint16(3),
		128: uint16(295),
	},
	52: {
		1:   uint16(297),
		9:   uint16(297),
		10:  uint16(297),
		31:  uint16(297),
		32:  uint16(297),
		33:  uint16(297),
		34:  uint16(297),
		36:  uint16(297),
		38:  uint16(297),
		40:  uint16(297),
		53:  uint16(297),
		54:  uint16(297),
		55:  uint16(297),
		56:  uint16(297),
		57:  uint16(297),
		58:  uint16(297),
		59:  uint16(297),
		60:  uint16(297),
		61:  uint16(297),
		62:  uint16(299),
		63:  uint16(297),
		64:  uint16(297),
		65:  uint16(297),
		66:  uint16(297),
		71:  uint16(297),
		72:  uint16(297),
		73:  uint16(297),
		74:  uint16(297),
		75:  uint16(297),
		76:  uint16(297),
		77:  uint16(297),
		78:  uint16(297),
		79:  uint16(297),
		80:  uint16(297),
		81:  uint16(297),
		82:  uint16(297),
		83:  uint16(297),
		84:  uint16(297),
		85:  uint16(297),
		86:  uint16(297),
		87:  uint16(297),
		88:  uint16(297),
		89:  uint16(297),
		90:  uint16(297),
		91:  uint16(297),
		92:  uint16(297),
		93:  uint16(297),
		94:  uint16(297),
		95:  uint16(297),
		96:  uint16(297),
		97:  uint16(297),
		98:  uint16(297),
		99:  uint16(297),
		100: uint16(297),
		101: uint16(297),
		102: uint16(297),
		103: uint16(297),
		104: uint16(297),
		105: uint16(297),
		106: uint16(297),
		107: uint16(297),
		108: uint16(297),
		109: uint16(297),
		110: uint16(297),
		111: uint16(297),
		125: uint16(3),
		128: uint16(301),
	},
	53: {
		1:   uint16(303),
		9:   uint16(303),
		10:  uint16(303),
		31:  uint16(303),
		32:  uint16(303),
		33:  uint16(303),
		34:  uint16(303),
		36:  uint16(303),
		38:  uint16(303),
		40:  uint16(303),
		53:  uint16(303),
		54:  uint16(303),
		55:  uint16(303),
		56:  uint16(303),
		57:  uint16(303),
		58:  uint16(303),
		59:  uint16(303),
		60:  uint16(303),
		61:  uint16(303),
		62:  uint16(305),
		63:  uint16(303),
		64:  uint16(303),
		65:  uint16(303),
		66:  uint16(303),
		71:  uint16(303),
		72:  uint16(303),
		73:  uint16(303),
		74:  uint16(303),
		75:  uint16(303),
		76:  uint16(303),
		77:  uint16(303),
		78:  uint16(303),
		79:  uint16(303),
		80:  uint16(303),
		81:  uint16(303),
		82:  uint16(303),
		83:  uint16(303),
		84:  uint16(303),
		85:  uint16(303),
		86:  uint16(303),
		87:  uint16(303),
		88:  uint16(303),
		89:  uint16(303),
		90:  uint16(303),
		91:  uint16(303),
		92:  uint16(303),
		93:  uint16(303),
		94:  uint16(303),
		95:  uint16(303),
		96:  uint16(303),
		97:  uint16(303),
		98:  uint16(303),
		99:  uint16(303),
		100: uint16(303),
		101: uint16(303),
		102: uint16(303),
		103: uint16(303),
		104: uint16(303),
		105: uint16(303),
		106: uint16(303),
		107: uint16(303),
		108: uint16(303),
		109: uint16(303),
		110: uint16(303),
		111: uint16(303),
		125: uint16(3),
		128: uint16(307),
	},
	54: {
		1:   uint16(309),
		9:   uint16(309),
		10:  uint16(309),
		31:  uint16(309),
		32:  uint16(309),
		33:  uint16(309),
		34:  uint16(309),
		36:  uint16(309),
		38:  uint16(309),
		40:  uint16(309),
		53:  uint16(309),
		54:  uint16(309),
		55:  uint16(309),
		56:  uint16(309),
		57:  uint16(309),
		58:  uint16(309),
		59:  uint16(309),
		60:  uint16(309),
		61:  uint16(309),
		62:  uint16(311),
		63:  uint16(309),
		64:  uint16(309),
		65:  uint16(309),
		66:  uint16(309),
		71:  uint16(309),
		72:  uint16(309),
		73:  uint16(309),
		74:  uint16(309),
		75:  uint16(309),
		76:  uint16(309),
		77:  uint16(309),
		78:  uint16(309),
		79:  uint16(309),
		80:  uint16(309),
		81:  uint16(309),
		82:  uint16(309),
		83:  uint16(309),
		84:  uint16(309),
		85:  uint16(309),
		86:  uint16(309),
		87:  uint16(309),
		88:  uint16(309),
		89:  uint16(309),
		90:  uint16(309),
		91:  uint16(309),
		92:  uint16(309),
		93:  uint16(309),
		94:  uint16(309),
		95:  uint16(309),
		96:  uint16(309),
		97:  uint16(309),
		98:  uint16(309),
		99:  uint16(309),
		100: uint16(309),
		101: uint16(309),
		102: uint16(309),
		103: uint16(309),
		104: uint16(309),
		105: uint16(309),
		106: uint16(309),
		107: uint16(309),
		108: uint16(309),
		109: uint16(309),
		110: uint16(309),
		111: uint16(309),
		125: uint16(3),
		128: uint16(313),
	},
	55: {
		1:   uint16(315),
		9:   uint16(315),
		10:  uint16(315),
		31:  uint16(315),
		32:  uint16(315),
		33:  uint16(315),
		34:  uint16(315),
		36:  uint16(315),
		38:  uint16(315),
		40:  uint16(315),
		53:  uint16(315),
		54:  uint16(315),
		55:  uint16(315),
		56:  uint16(315),
		57:  uint16(315),
		58:  uint16(315),
		59:  uint16(315),
		60:  uint16(315),
		61:  uint16(315),
		62:  uint16(317),
		63:  uint16(315),
		64:  uint16(315),
		65:  uint16(315),
		66:  uint16(315),
		71:  uint16(315),
		72:  uint16(315),
		73:  uint16(315),
		74:  uint16(315),
		75:  uint16(315),
		76:  uint16(315),
		77:  uint16(315),
		78:  uint16(315),
		79:  uint16(315),
		80:  uint16(315),
		81:  uint16(315),
		82:  uint16(315),
		83:  uint16(315),
		84:  uint16(315),
		85:  uint16(315),
		86:  uint16(315),
		87:  uint16(315),
		88:  uint16(315),
		89:  uint16(315),
		90:  uint16(315),
		91:  uint16(315),
		92:  uint16(315),
		93:  uint16(315),
		94:  uint16(315),
		95:  uint16(315),
		96:  uint16(315),
		97:  uint16(315),
		98:  uint16(315),
		99:  uint16(315),
		100: uint16(315),
		101: uint16(315),
		102: uint16(315),
		103: uint16(315),
		104: uint16(315),
		105: uint16(315),
		106: uint16(315),
		107: uint16(315),
		108: uint16(315),
		109: uint16(315),
		110: uint16(315),
		111: uint16(315),
		125: uint16(3),
		128: uint16(319),
	},
	56: {
		1:   uint16(321),
		9:   uint16(321),
		10:  uint16(321),
		31:  uint16(321),
		32:  uint16(321),
		33:  uint16(321),
		34:  uint16(321),
		36:  uint16(321),
		38:  uint16(321),
		40:  uint16(321),
		53:  uint16(321),
		54:  uint16(321),
		55:  uint16(321),
		56:  uint16(321),
		57:  uint16(321),
		58:  uint16(321),
		59:  uint16(321),
		60:  uint16(321),
		61:  uint16(321),
		62:  uint16(323),
		63:  uint16(321),
		64:  uint16(321),
		65:  uint16(321),
		66:  uint16(321),
		71:  uint16(321),
		72:  uint16(321),
		73:  uint16(321),
		74:  uint16(321),
		75:  uint16(321),
		76:  uint16(321),
		77:  uint16(321),
		78:  uint16(321),
		79:  uint16(321),
		80:  uint16(321),
		81:  uint16(321),
		82:  uint16(321),
		83:  uint16(321),
		84:  uint16(321),
		85:  uint16(321),
		86:  uint16(321),
		87:  uint16(321),
		88:  uint16(321),
		89:  uint16(321),
		90:  uint16(321),
		91:  uint16(321),
		92:  uint16(321),
		93:  uint16(321),
		94:  uint16(321),
		95:  uint16(321),
		96:  uint16(321),
		97:  uint16(321),
		98:  uint16(321),
		99:  uint16(321),
		100: uint16(321),
		101: uint16(321),
		102: uint16(321),
		103: uint16(321),
		104: uint16(321),
		105: uint16(321),
		106: uint16(321),
		107: uint16(321),
		108: uint16(321),
		109: uint16(321),
		110: uint16(321),
		111: uint16(321),
		125: uint16(3),
		128: uint16(325),
	},
	57: {
		1:   uint16(327),
		9:   uint16(327),
		10:  uint16(327),
		31:  uint16(327),
		32:  uint16(327),
		33:  uint16(327),
		34:  uint16(327),
		36:  uint16(327),
		38:  uint16(327),
		40:  uint16(327),
		53:  uint16(327),
		54:  uint16(327),
		55:  uint16(327),
		56:  uint16(327),
		57:  uint16(327),
		58:  uint16(327),
		59:  uint16(327),
		60:  uint16(327),
		61:  uint16(327),
		62:  uint16(329),
		63:  uint16(327),
		64:  uint16(327),
		65:  uint16(327),
		66:  uint16(327),
		71:  uint16(327),
		72:  uint16(327),
		73:  uint16(327),
		74:  uint16(327),
		75:  uint16(327),
		76:  uint16(327),
		77:  uint16(327),
		78:  uint16(327),
		79:  uint16(327),
		80:  uint16(327),
		81:  uint16(327),
		82:  uint16(327),
		83:  uint16(327),
		84:  uint16(327),
		85:  uint16(327),
		86:  uint16(327),
		87:  uint16(327),
		88:  uint16(327),
		89:  uint16(327),
		90:  uint16(327),
		91:  uint16(327),
		92:  uint16(327),
		93:  uint16(327),
		94:  uint16(327),
		95:  uint16(327),
		96:  uint16(327),
		97:  uint16(327),
		98:  uint16(327),
		99:  uint16(327),
		100: uint16(327),
		101: uint16(327),
		102: uint16(327),
		103: uint16(327),
		104: uint16(327),
		105: uint16(327),
		106: uint16(327),
		107: uint16(327),
		108: uint16(327),
		109: uint16(327),
		110: uint16(327),
		111: uint16(327),
		125: uint16(3),
		128: uint16(331),
	},
	58: {
		1:   uint16(333),
		9:   uint16(333),
		10:  uint16(333),
		31:  uint16(333),
		32:  uint16(333),
		33:  uint16(333),
		34:  uint16(333),
		36:  uint16(333),
		38:  uint16(333),
		40:  uint16(333),
		53:  uint16(333),
		54:  uint16(55),
		55:  uint16(333),
		56:  uint16(333),
		57:  uint16(333),
		58:  uint16(333),
		59:  uint16(333),
		60:  uint16(333),
		61:  uint16(333),
		63:  uint16(333),
		64:  uint16(333),
		65:  uint16(333),
		66:  uint16(333),
		71:  uint16(333),
		72:  uint16(333),
		73:  uint16(333),
		74:  uint16(333),
		75:  uint16(333),
		76:  uint16(333),
		77:  uint16(333),
		78:  uint16(333),
		79:  uint16(333),
		80:  uint16(333),
		81:  uint16(333),
		82:  uint16(333),
		83:  uint16(333),
		84:  uint16(333),
		85:  uint16(333),
		86:  uint16(333),
		87:  uint16(333),
		88:  uint16(333),
		89:  uint16(333),
		90:  uint16(333),
		91:  uint16(333),
		92:  uint16(333),
		93:  uint16(333),
		94:  uint16(333),
		95:  uint16(333),
		96:  uint16(333),
		97:  uint16(333),
		98:  uint16(333),
		99:  uint16(333),
		100: uint16(333),
		101: uint16(333),
		102: uint16(333),
		103: uint16(333),
		104: uint16(333),
		105: uint16(333),
		106: uint16(333),
		107: uint16(333),
		108: uint16(333),
		109: uint16(333),
		110: uint16(333),
		111: uint16(333),
		125: uint16(3),
		128: uint16(335),
		157: uint16(88),
	},
	59: {
		1:   uint16(337),
		9:   uint16(337),
		10:  uint16(337),
		31:  uint16(337),
		32:  uint16(337),
		33:  uint16(337),
		34:  uint16(337),
		36:  uint16(337),
		38:  uint16(337),
		40:  uint16(337),
		53:  uint16(337),
		54:  uint16(337),
		55:  uint16(337),
		56:  uint16(337),
		57:  uint16(337),
		58:  uint16(337),
		59:  uint16(337),
		60:  uint16(337),
		61:  uint16(337),
		62:  uint16(339),
		63:  uint16(337),
		64:  uint16(337),
		65:  uint16(337),
		66:  uint16(337),
		71:  uint16(337),
		72:  uint16(337),
		73:  uint16(337),
		74:  uint16(337),
		75:  uint16(337),
		76:  uint16(337),
		77:  uint16(337),
		78:  uint16(337),
		79:  uint16(337),
		80:  uint16(337),
		81:  uint16(337),
		82:  uint16(337),
		83:  uint16(337),
		84:  uint16(337),
		85:  uint16(337),
		86:  uint16(337),
		87:  uint16(337),
		88:  uint16(337),
		89:  uint16(337),
		90:  uint16(337),
		91:  uint16(337),
		92:  uint16(337),
		93:  uint16(337),
		94:  uint16(337),
		95:  uint16(337),
		96:  uint16(337),
		97:  uint16(337),
		98:  uint16(337),
		99:  uint16(337),
		100: uint16(337),
		101: uint16(337),
		102: uint16(337),
		103: uint16(337),
		104: uint16(337),
		105: uint16(337),
		106: uint16(337),
		107: uint16(337),
		108: uint16(337),
		109: uint16(337),
		110: uint16(337),
		111: uint16(337),
		125: uint16(3),
		128: uint16(339),
	},
	60: {
		1:   uint16(341),
		9:   uint16(341),
		10:  uint16(341),
		31:  uint16(341),
		32:  uint16(341),
		33:  uint16(341),
		34:  uint16(341),
		36:  uint16(341),
		38:  uint16(341),
		40:  uint16(341),
		53:  uint16(341),
		54:  uint16(341),
		55:  uint16(341),
		56:  uint16(341),
		57:  uint16(341),
		58:  uint16(341),
		59:  uint16(341),
		60:  uint16(341),
		61:  uint16(341),
		62:  uint16(343),
		63:  uint16(341),
		64:  uint16(341),
		65:  uint16(341),
		66:  uint16(341),
		71:  uint16(341),
		72:  uint16(341),
		73:  uint16(341),
		74:  uint16(341),
		75:  uint16(341),
		76:  uint16(341),
		77:  uint16(341),
		78:  uint16(341),
		79:  uint16(341),
		80:  uint16(341),
		81:  uint16(341),
		82:  uint16(341),
		83:  uint16(341),
		84:  uint16(341),
		85:  uint16(341),
		86:  uint16(341),
		87:  uint16(341),
		88:  uint16(341),
		89:  uint16(341),
		90:  uint16(341),
		91:  uint16(341),
		92:  uint16(341),
		93:  uint16(341),
		94:  uint16(341),
		95:  uint16(341),
		96:  uint16(341),
		97:  uint16(341),
		98:  uint16(341),
		99:  uint16(341),
		100: uint16(341),
		101: uint16(341),
		102: uint16(341),
		103: uint16(341),
		104: uint16(341),
		105: uint16(341),
		106: uint16(341),
		107: uint16(341),
		108: uint16(341),
		109: uint16(341),
		110: uint16(341),
		111: uint16(341),
		125: uint16(3),
		128: uint16(343),
	},
	61: {
		1:   uint16(345),
		9:   uint16(345),
		10:  uint16(345),
		31:  uint16(345),
		32:  uint16(345),
		33:  uint16(345),
		34:  uint16(345),
		36:  uint16(345),
		38:  uint16(345),
		40:  uint16(345),
		53:  uint16(345),
		54:  uint16(345),
		55:  uint16(345),
		56:  uint16(345),
		57:  uint16(345),
		58:  uint16(345),
		59:  uint16(345),
		60:  uint16(345),
		61:  uint16(345),
		62:  uint16(347),
		63:  uint16(345),
		64:  uint16(345),
		65:  uint16(345),
		66:  uint16(345),
		71:  uint16(345),
		72:  uint16(345),
		73:  uint16(345),
		74:  uint16(345),
		75:  uint16(345),
		76:  uint16(345),
		77:  uint16(345),
		78:  uint16(345),
		79:  uint16(345),
		80:  uint16(345),
		81:  uint16(345),
		82:  uint16(345),
		83:  uint16(345),
		84:  uint16(345),
		85:  uint16(345),
		86:  uint16(345),
		87:  uint16(345),
		88:  uint16(345),
		89:  uint16(345),
		90:  uint16(345),
		91:  uint16(345),
		92:  uint16(345),
		93:  uint16(345),
		94:  uint16(345),
		95:  uint16(345),
		96:  uint16(345),
		97:  uint16(345),
		98:  uint16(345),
		99:  uint16(345),
		100: uint16(345),
		101: uint16(345),
		102: uint16(345),
		103: uint16(345),
		104: uint16(345),
		105: uint16(345),
		106: uint16(345),
		107: uint16(345),
		108: uint16(345),
		109: uint16(345),
		110: uint16(345),
		111: uint16(345),
		125: uint16(3),
		128: uint16(349),
	},
	62: {
		1:   uint16(351),
		9:   uint16(351),
		10:  uint16(351),
		31:  uint16(351),
		32:  uint16(351),
		33:  uint16(351),
		34:  uint16(351),
		36:  uint16(351),
		38:  uint16(351),
		40:  uint16(351),
		53:  uint16(351),
		54:  uint16(351),
		55:  uint16(351),
		56:  uint16(351),
		57:  uint16(351),
		58:  uint16(351),
		59:  uint16(351),
		60:  uint16(351),
		61:  uint16(351),
		62:  uint16(353),
		63:  uint16(351),
		64:  uint16(351),
		65:  uint16(351),
		66:  uint16(351),
		71:  uint16(351),
		72:  uint16(351),
		73:  uint16(351),
		74:  uint16(351),
		75:  uint16(351),
		76:  uint16(351),
		77:  uint16(351),
		78:  uint16(351),
		79:  uint16(351),
		80:  uint16(351),
		81:  uint16(351),
		82:  uint16(351),
		83:  uint16(351),
		84:  uint16(351),
		85:  uint16(351),
		86:  uint16(351),
		87:  uint16(351),
		88:  uint16(351),
		89:  uint16(351),
		90:  uint16(351),
		91:  uint16(351),
		92:  uint16(351),
		93:  uint16(351),
		94:  uint16(351),
		95:  uint16(351),
		96:  uint16(351),
		97:  uint16(351),
		98:  uint16(351),
		99:  uint16(351),
		100: uint16(351),
		101: uint16(351),
		102: uint16(351),
		103: uint16(351),
		104: uint16(351),
		105: uint16(351),
		106: uint16(351),
		107: uint16(351),
		108: uint16(351),
		109: uint16(351),
		110: uint16(351),
		111: uint16(351),
		125: uint16(3),
		128: uint16(355),
	},
	63: {
		1:   uint16(357),
		9:   uint16(357),
		10:  uint16(357),
		31:  uint16(357),
		32:  uint16(357),
		33:  uint16(357),
		34:  uint16(357),
		36:  uint16(357),
		38:  uint16(357),
		40:  uint16(357),
		53:  uint16(357),
		54:  uint16(55),
		55:  uint16(357),
		56:  uint16(357),
		57:  uint16(357),
		58:  uint16(357),
		59:  uint16(357),
		60:  uint16(357),
		61:  uint16(357),
		63:  uint16(357),
		64:  uint16(357),
		65:  uint16(357),
		66:  uint16(357),
		71:  uint16(357),
		72:  uint16(357),
		73:  uint16(357),
		74:  uint16(357),
		75:  uint16(357),
		76:  uint16(357),
		77:  uint16(357),
		78:  uint16(357),
		79:  uint16(357),
		80:  uint16(357),
		81:  uint16(357),
		82:  uint16(357),
		83:  uint16(357),
		84:  uint16(357),
		85:  uint16(357),
		86:  uint16(357),
		87:  uint16(357),
		88:  uint16(357),
		89:  uint16(357),
		90:  uint16(357),
		91:  uint16(357),
		92:  uint16(357),
		93:  uint16(357),
		94:  uint16(357),
		95:  uint16(357),
		96:  uint16(357),
		97:  uint16(357),
		98:  uint16(357),
		99:  uint16(357),
		100: uint16(357),
		101: uint16(357),
		102: uint16(357),
		103: uint16(357),
		104: uint16(357),
		105: uint16(357),
		106: uint16(357),
		107: uint16(357),
		108: uint16(357),
		109: uint16(357),
		110: uint16(357),
		111: uint16(357),
		125: uint16(3),
		128: uint16(63),
		157: uint16(76),
	},
	64: {
		1:   uint16(359),
		9:   uint16(359),
		10:  uint16(359),
		31:  uint16(359),
		32:  uint16(359),
		33:  uint16(359),
		34:  uint16(359),
		36:  uint16(359),
		38:  uint16(359),
		40:  uint16(359),
		53:  uint16(359),
		54:  uint16(359),
		55:  uint16(359),
		56:  uint16(359),
		57:  uint16(359),
		58:  uint16(359),
		59:  uint16(359),
		60:  uint16(359),
		61:  uint16(359),
		62:  uint16(361),
		63:  uint16(359),
		64:  uint16(359),
		65:  uint16(359),
		66:  uint16(359),
		71:  uint16(359),
		72:  uint16(359),
		73:  uint16(359),
		74:  uint16(359),
		75:  uint16(359),
		76:  uint16(359),
		77:  uint16(359),
		78:  uint16(359),
		79:  uint16(359),
		80:  uint16(359),
		81:  uint16(359),
		82:  uint16(359),
		83:  uint16(359),
		84:  uint16(359),
		85:  uint16(359),
		86:  uint16(359),
		87:  uint16(359),
		88:  uint16(359),
		89:  uint16(359),
		90:  uint16(359),
		91:  uint16(359),
		92:  uint16(359),
		93:  uint16(359),
		94:  uint16(359),
		95:  uint16(359),
		96:  uint16(359),
		97:  uint16(359),
		98:  uint16(359),
		99:  uint16(359),
		100: uint16(359),
		101: uint16(359),
		102: uint16(359),
		103: uint16(359),
		104: uint16(359),
		105: uint16(359),
		106: uint16(359),
		107: uint16(359),
		108: uint16(359),
		109: uint16(359),
		110: uint16(359),
		111: uint16(359),
		125: uint16(3),
		128: uint16(361),
	},
	65: {
		1:   uint16(363),
		9:   uint16(363),
		10:  uint16(363),
		31:  uint16(363),
		32:  uint16(363),
		33:  uint16(363),
		34:  uint16(363),
		36:  uint16(363),
		38:  uint16(363),
		40:  uint16(363),
		53:  uint16(363),
		54:  uint16(363),
		55:  uint16(363),
		56:  uint16(363),
		57:  uint16(363),
		58:  uint16(363),
		59:  uint16(363),
		60:  uint16(363),
		61:  uint16(363),
		63:  uint16(363),
		64:  uint16(363),
		65:  uint16(363),
		66:  uint16(363),
		71:  uint16(363),
		72:  uint16(363),
		73:  uint16(363),
		74:  uint16(363),
		75:  uint16(363),
		76:  uint16(363),
		77:  uint16(363),
		78:  uint16(363),
		79:  uint16(363),
		80:  uint16(363),
		81:  uint16(363),
		82:  uint16(363),
		83:  uint16(363),
		84:  uint16(363),
		85:  uint16(363),
		86:  uint16(363),
		87:  uint16(363),
		88:  uint16(363),
		89:  uint16(363),
		90:  uint16(363),
		91:  uint16(363),
		92:  uint16(363),
		93:  uint16(363),
		94:  uint16(363),
		95:  uint16(363),
		96:  uint16(363),
		97:  uint16(363),
		98:  uint16(363),
		99:  uint16(363),
		100: uint16(363),
		101: uint16(363),
		102: uint16(363),
		103: uint16(363),
		104: uint16(363),
		105: uint16(363),
		106: uint16(363),
		107: uint16(363),
		108: uint16(363),
		109: uint16(363),
		110: uint16(363),
		111: uint16(363),
		125: uint16(3),
		128: uint16(365),
	},
	66: {
		1:   uint16(367),
		9:   uint16(367),
		10:  uint16(367),
		31:  uint16(367),
		32:  uint16(367),
		33:  uint16(367),
		34:  uint16(367),
		36:  uint16(367),
		38:  uint16(367),
		40:  uint16(367),
		53:  uint16(367),
		54:  uint16(367),
		55:  uint16(367),
		56:  uint16(367),
		57:  uint16(367),
		58:  uint16(367),
		59:  uint16(367),
		60:  uint16(367),
		61:  uint16(367),
		63:  uint16(367),
		64:  uint16(367),
		65:  uint16(367),
		66:  uint16(367),
		71:  uint16(367),
		72:  uint16(367),
		73:  uint16(367),
		74:  uint16(367),
		75:  uint16(367),
		76:  uint16(367),
		77:  uint16(367),
		78:  uint16(367),
		79:  uint16(367),
		80:  uint16(367),
		81:  uint16(367),
		82:  uint16(367),
		83:  uint16(367),
		84:  uint16(367),
		85:  uint16(367),
		86:  uint16(367),
		87:  uint16(367),
		88:  uint16(367),
		89:  uint16(367),
		90:  uint16(367),
		91:  uint16(367),
		92:  uint16(367),
		93:  uint16(367),
		94:  uint16(367),
		95:  uint16(367),
		96:  uint16(367),
		97:  uint16(367),
		98:  uint16(367),
		99:  uint16(367),
		100: uint16(367),
		101: uint16(367),
		102: uint16(367),
		103: uint16(367),
		104: uint16(367),
		105: uint16(367),
		106: uint16(367),
		107: uint16(367),
		108: uint16(367),
		109: uint16(367),
		110: uint16(367),
		111: uint16(367),
		125: uint16(3),
		128: uint16(369),
	},
	67: {
		1:   uint16(371),
		9:   uint16(371),
		10:  uint16(371),
		31:  uint16(371),
		32:  uint16(371),
		33:  uint16(371),
		34:  uint16(371),
		36:  uint16(371),
		38:  uint16(371),
		40:  uint16(371),
		53:  uint16(371),
		54:  uint16(371),
		55:  uint16(371),
		56:  uint16(371),
		57:  uint16(371),
		58:  uint16(371),
		59:  uint16(371),
		60:  uint16(371),
		61:  uint16(371),
		63:  uint16(371),
		64:  uint16(371),
		65:  uint16(371),
		66:  uint16(371),
		71:  uint16(371),
		72:  uint16(371),
		73:  uint16(371),
		74:  uint16(371),
		75:  uint16(371),
		76:  uint16(371),
		77:  uint16(371),
		78:  uint16(371),
		79:  uint16(371),
		80:  uint16(371),
		81:  uint16(371),
		82:  uint16(371),
		83:  uint16(371),
		84:  uint16(371),
		85:  uint16(371),
		86:  uint16(371),
		87:  uint16(371),
		88:  uint16(371),
		89:  uint16(371),
		90:  uint16(371),
		91:  uint16(371),
		92:  uint16(371),
		93:  uint16(371),
		94:  uint16(371),
		95:  uint16(371),
		96:  uint16(371),
		97:  uint16(371),
		98:  uint16(371),
		99:  uint16(371),
		100: uint16(371),
		101: uint16(371),
		102: uint16(371),
		103: uint16(371),
		104: uint16(371),
		105: uint16(371),
		106: uint16(371),
		107: uint16(371),
		108: uint16(371),
		109: uint16(371),
		110: uint16(371),
		111: uint16(371),
		125: uint16(3),
		128: uint16(373),
	},
	68: {
		1:   uint16(375),
		9:   uint16(375),
		10:  uint16(375),
		31:  uint16(375),
		32:  uint16(375),
		33:  uint16(375),
		34:  uint16(375),
		36:  uint16(375),
		38:  uint16(375),
		40:  uint16(375),
		53:  uint16(375),
		54:  uint16(375),
		55:  uint16(375),
		56:  uint16(375),
		57:  uint16(375),
		58:  uint16(375),
		59:  uint16(375),
		60:  uint16(375),
		61:  uint16(375),
		63:  uint16(375),
		64:  uint16(375),
		65:  uint16(375),
		66:  uint16(375),
		71:  uint16(375),
		72:  uint16(375),
		73:  uint16(375),
		74:  uint16(375),
		75:  uint16(375),
		76:  uint16(375),
		77:  uint16(375),
		78:  uint16(375),
		79:  uint16(375),
		80:  uint16(375),
		81:  uint16(375),
		82:  uint16(375),
		83:  uint16(375),
		84:  uint16(375),
		85:  uint16(375),
		86:  uint16(375),
		87:  uint16(375),
		88:  uint16(375),
		89:  uint16(375),
		90:  uint16(375),
		91:  uint16(375),
		92:  uint16(375),
		93:  uint16(375),
		94:  uint16(375),
		95:  uint16(375),
		96:  uint16(375),
		97:  uint16(375),
		98:  uint16(375),
		99:  uint16(375),
		100: uint16(375),
		101: uint16(375),
		102: uint16(375),
		103: uint16(375),
		104: uint16(375),
		105: uint16(375),
		106: uint16(375),
		107: uint16(375),
		108: uint16(375),
		109: uint16(375),
		110: uint16(375),
		111: uint16(375),
		125: uint16(3),
		128: uint16(377),
	},
	69: {
		1:   uint16(379),
		9:   uint16(379),
		10:  uint16(379),
		31:  uint16(379),
		32:  uint16(379),
		33:  uint16(379),
		34:  uint16(379),
		36:  uint16(379),
		38:  uint16(379),
		40:  uint16(379),
		53:  uint16(379),
		54:  uint16(379),
		55:  uint16(379),
		56:  uint16(379),
		57:  uint16(379),
		58:  uint16(379),
		59:  uint16(379),
		60:  uint16(379),
		61:  uint16(379),
		63:  uint16(379),
		64:  uint16(379),
		65:  uint16(379),
		66:  uint16(379),
		71:  uint16(379),
		72:  uint16(379),
		73:  uint16(379),
		74:  uint16(379),
		75:  uint16(379),
		76:  uint16(379),
		77:  uint16(379),
		78:  uint16(379),
		79:  uint16(379),
		80:  uint16(379),
		81:  uint16(379),
		82:  uint16(379),
		83:  uint16(379),
		84:  uint16(379),
		85:  uint16(379),
		86:  uint16(379),
		87:  uint16(379),
		88:  uint16(379),
		89:  uint16(379),
		90:  uint16(379),
		91:  uint16(379),
		92:  uint16(379),
		93:  uint16(379),
		94:  uint16(379),
		95:  uint16(379),
		96:  uint16(379),
		97:  uint16(379),
		98:  uint16(379),
		99:  uint16(379),
		100: uint16(379),
		101: uint16(379),
		102: uint16(379),
		103: uint16(379),
		104: uint16(379),
		105: uint16(379),
		106: uint16(379),
		107: uint16(379),
		108: uint16(379),
		109: uint16(379),
		110: uint16(379),
		111: uint16(379),
		125: uint16(3),
		128: uint16(381),
	},
	70: {
		1:   uint16(383),
		9:   uint16(383),
		10:  uint16(383),
		31:  uint16(383),
		32:  uint16(383),
		33:  uint16(383),
		34:  uint16(383),
		36:  uint16(383),
		38:  uint16(383),
		40:  uint16(383),
		53:  uint16(383),
		54:  uint16(383),
		55:  uint16(383),
		56:  uint16(383),
		57:  uint16(383),
		58:  uint16(383),
		59:  uint16(383),
		60:  uint16(383),
		61:  uint16(383),
		63:  uint16(383),
		64:  uint16(383),
		65:  uint16(383),
		66:  uint16(383),
		71:  uint16(383),
		72:  uint16(383),
		73:  uint16(383),
		74:  uint16(383),
		75:  uint16(383),
		76:  uint16(383),
		77:  uint16(383),
		78:  uint16(383),
		79:  uint16(383),
		80:  uint16(383),
		81:  uint16(383),
		82:  uint16(383),
		83:  uint16(383),
		84:  uint16(383),
		85:  uint16(383),
		86:  uint16(383),
		87:  uint16(383),
		88:  uint16(383),
		89:  uint16(383),
		90:  uint16(383),
		91:  uint16(383),
		92:  uint16(383),
		93:  uint16(383),
		94:  uint16(383),
		95:  uint16(383),
		96:  uint16(383),
		97:  uint16(383),
		98:  uint16(383),
		99:  uint16(383),
		100: uint16(383),
		101: uint16(383),
		102: uint16(383),
		103: uint16(383),
		104: uint16(383),
		105: uint16(383),
		106: uint16(383),
		107: uint16(383),
		108: uint16(383),
		109: uint16(383),
		110: uint16(383),
		111: uint16(383),
		125: uint16(3),
		128: uint16(385),
	},
	71: {
		1:   uint16(387),
		9:   uint16(387),
		10:  uint16(387),
		31:  uint16(387),
		32:  uint16(387),
		33:  uint16(387),
		34:  uint16(387),
		36:  uint16(387),
		38:  uint16(387),
		40:  uint16(387),
		53:  uint16(387),
		54:  uint16(387),
		55:  uint16(387),
		56:  uint16(387),
		57:  uint16(387),
		58:  uint16(387),
		59:  uint16(387),
		60:  uint16(387),
		61:  uint16(387),
		63:  uint16(387),
		64:  uint16(387),
		65:  uint16(387),
		66:  uint16(387),
		71:  uint16(387),
		72:  uint16(387),
		73:  uint16(387),
		74:  uint16(387),
		75:  uint16(387),
		76:  uint16(387),
		77:  uint16(387),
		78:  uint16(387),
		79:  uint16(387),
		80:  uint16(387),
		81:  uint16(387),
		82:  uint16(387),
		83:  uint16(387),
		84:  uint16(387),
		85:  uint16(387),
		86:  uint16(387),
		87:  uint16(387),
		88:  uint16(387),
		89:  uint16(387),
		90:  uint16(387),
		91:  uint16(387),
		92:  uint16(387),
		93:  uint16(387),
		94:  uint16(387),
		95:  uint16(387),
		96:  uint16(387),
		97:  uint16(387),
		98:  uint16(387),
		99:  uint16(387),
		100: uint16(387),
		101: uint16(387),
		102: uint16(387),
		103: uint16(387),
		104: uint16(387),
		105: uint16(387),
		106: uint16(387),
		107: uint16(387),
		108: uint16(387),
		109: uint16(387),
		110: uint16(387),
		111: uint16(387),
		125: uint16(3),
		128: uint16(389),
	},
	72: {
		1:   uint16(391),
		9:   uint16(391),
		10:  uint16(391),
		31:  uint16(391),
		32:  uint16(391),
		33:  uint16(391),
		34:  uint16(391),
		36:  uint16(391),
		38:  uint16(391),
		40:  uint16(391),
		53:  uint16(391),
		54:  uint16(391),
		55:  uint16(391),
		56:  uint16(391),
		57:  uint16(391),
		58:  uint16(391),
		59:  uint16(391),
		60:  uint16(391),
		61:  uint16(391),
		63:  uint16(391),
		64:  uint16(391),
		65:  uint16(391),
		66:  uint16(391),
		71:  uint16(391),
		72:  uint16(391),
		73:  uint16(391),
		74:  uint16(391),
		75:  uint16(391),
		76:  uint16(391),
		77:  uint16(391),
		78:  uint16(391),
		79:  uint16(391),
		80:  uint16(391),
		81:  uint16(391),
		82:  uint16(391),
		83:  uint16(391),
		84:  uint16(391),
		85:  uint16(391),
		86:  uint16(391),
		87:  uint16(391),
		88:  uint16(391),
		89:  uint16(391),
		90:  uint16(391),
		91:  uint16(391),
		92:  uint16(391),
		93:  uint16(391),
		94:  uint16(391),
		95:  uint16(391),
		96:  uint16(391),
		97:  uint16(391),
		98:  uint16(391),
		99:  uint16(391),
		100: uint16(391),
		101: uint16(391),
		102: uint16(391),
		103: uint16(391),
		104: uint16(391),
		105: uint16(391),
		106: uint16(391),
		107: uint16(391),
		108: uint16(391),
		109: uint16(391),
		110: uint16(391),
		111: uint16(391),
		125: uint16(3),
		128: uint16(393),
	},
	73: {
		1:   uint16(395),
		9:   uint16(395),
		10:  uint16(395),
		31:  uint16(395),
		32:  uint16(395),
		33:  uint16(395),
		34:  uint16(395),
		36:  uint16(395),
		38:  uint16(395),
		40:  uint16(395),
		53:  uint16(395),
		54:  uint16(395),
		55:  uint16(395),
		56:  uint16(395),
		57:  uint16(395),
		58:  uint16(395),
		59:  uint16(395),
		60:  uint16(395),
		61:  uint16(395),
		63:  uint16(395),
		64:  uint16(395),
		65:  uint16(395),
		66:  uint16(395),
		71:  uint16(395),
		72:  uint16(395),
		73:  uint16(395),
		74:  uint16(395),
		75:  uint16(395),
		76:  uint16(395),
		77:  uint16(395),
		78:  uint16(395),
		79:  uint16(395),
		80:  uint16(395),
		81:  uint16(395),
		82:  uint16(395),
		83:  uint16(395),
		84:  uint16(395),
		85:  uint16(395),
		86:  uint16(395),
		87:  uint16(395),
		88:  uint16(395),
		89:  uint16(395),
		90:  uint16(395),
		91:  uint16(395),
		92:  uint16(395),
		93:  uint16(395),
		94:  uint16(395),
		95:  uint16(395),
		96:  uint16(395),
		97:  uint16(395),
		98:  uint16(395),
		99:  uint16(395),
		100: uint16(395),
		101: uint16(395),
		102: uint16(395),
		103: uint16(395),
		104: uint16(395),
		105: uint16(395),
		106: uint16(395),
		107: uint16(395),
		108: uint16(395),
		109: uint16(395),
		110: uint16(395),
		111: uint16(395),
		125: uint16(3),
		128: uint16(397),
	},
	74: {
		1:   uint16(399),
		9:   uint16(399),
		10:  uint16(399),
		31:  uint16(399),
		32:  uint16(399),
		33:  uint16(399),
		34:  uint16(399),
		36:  uint16(399),
		38:  uint16(399),
		40:  uint16(399),
		53:  uint16(399),
		54:  uint16(399),
		55:  uint16(399),
		56:  uint16(399),
		57:  uint16(399),
		58:  uint16(399),
		59:  uint16(399),
		60:  uint16(399),
		61:  uint16(399),
		63:  uint16(399),
		64:  uint16(399),
		65:  uint16(399),
		66:  uint16(399),
		71:  uint16(399),
		72:  uint16(399),
		73:  uint16(399),
		74:  uint16(399),
		75:  uint16(399),
		76:  uint16(399),
		77:  uint16(399),
		78:  uint16(399),
		79:  uint16(399),
		80:  uint16(399),
		81:  uint16(399),
		82:  uint16(399),
		83:  uint16(399),
		84:  uint16(399),
		85:  uint16(399),
		86:  uint16(399),
		87:  uint16(399),
		88:  uint16(399),
		89:  uint16(399),
		90:  uint16(399),
		91:  uint16(399),
		92:  uint16(399),
		93:  uint16(399),
		94:  uint16(399),
		95:  uint16(399),
		96:  uint16(399),
		97:  uint16(399),
		98:  uint16(399),
		99:  uint16(399),
		100: uint16(399),
		101: uint16(399),
		102: uint16(399),
		103: uint16(399),
		104: uint16(399),
		105: uint16(399),
		106: uint16(399),
		107: uint16(399),
		108: uint16(399),
		109: uint16(399),
		110: uint16(399),
		111: uint16(399),
		125: uint16(3),
		128: uint16(401),
	},
	75: {
		1:   uint16(403),
		9:   uint16(403),
		10:  uint16(403),
		31:  uint16(403),
		32:  uint16(403),
		33:  uint16(403),
		34:  uint16(403),
		36:  uint16(403),
		38:  uint16(403),
		40:  uint16(403),
		53:  uint16(403),
		54:  uint16(403),
		55:  uint16(403),
		56:  uint16(403),
		57:  uint16(403),
		58:  uint16(403),
		59:  uint16(403),
		60:  uint16(403),
		61:  uint16(403),
		63:  uint16(403),
		64:  uint16(403),
		65:  uint16(403),
		66:  uint16(403),
		71:  uint16(403),
		72:  uint16(403),
		73:  uint16(403),
		74:  uint16(403),
		75:  uint16(403),
		76:  uint16(403),
		77:  uint16(403),
		78:  uint16(403),
		79:  uint16(403),
		80:  uint16(403),
		81:  uint16(403),
		82:  uint16(403),
		83:  uint16(403),
		84:  uint16(403),
		85:  uint16(403),
		86:  uint16(403),
		87:  uint16(403),
		88:  uint16(403),
		89:  uint16(403),
		90:  uint16(403),
		91:  uint16(403),
		92:  uint16(403),
		93:  uint16(403),
		94:  uint16(403),
		95:  uint16(403),
		96:  uint16(403),
		97:  uint16(403),
		98:  uint16(403),
		99:  uint16(403),
		100: uint16(403),
		101: uint16(403),
		102: uint16(403),
		103: uint16(403),
		104: uint16(403),
		105: uint16(403),
		106: uint16(403),
		107: uint16(403),
		108: uint16(403),
		109: uint16(403),
		110: uint16(403),
		111: uint16(403),
		125: uint16(3),
		128: uint16(405),
	},
	76: {
		1:   uint16(333),
		9:   uint16(333),
		10:  uint16(333),
		31:  uint16(333),
		32:  uint16(333),
		33:  uint16(333),
		34:  uint16(333),
		36:  uint16(333),
		38:  uint16(333),
		40:  uint16(333),
		53:  uint16(333),
		54:  uint16(333),
		55:  uint16(333),
		56:  uint16(333),
		57:  uint16(333),
		58:  uint16(333),
		59:  uint16(333),
		60:  uint16(333),
		61:  uint16(333),
		63:  uint16(333),
		64:  uint16(333),
		65:  uint16(333),
		66:  uint16(333),
		71:  uint16(333),
		72:  uint16(333),
		73:  uint16(333),
		74:  uint16(333),
		75:  uint16(333),
		76:  uint16(333),
		77:  uint16(333),
		78:  uint16(333),
		79:  uint16(333),
		80:  uint16(333),
		81:  uint16(333),
		82:  uint16(333),
		83:  uint16(333),
		84:  uint16(333),
		85:  uint16(333),
		86:  uint16(333),
		87:  uint16(333),
		88:  uint16(333),
		89:  uint16(333),
		90:  uint16(333),
		91:  uint16(333),
		92:  uint16(333),
		93:  uint16(333),
		94:  uint16(333),
		95:  uint16(333),
		96:  uint16(333),
		97:  uint16(333),
		98:  uint16(333),
		99:  uint16(333),
		100: uint16(333),
		101: uint16(333),
		102: uint16(333),
		103: uint16(333),
		104: uint16(333),
		105: uint16(333),
		106: uint16(333),
		107: uint16(333),
		108: uint16(333),
		109: uint16(333),
		110: uint16(333),
		111: uint16(333),
		125: uint16(3),
		128: uint16(335),
	},
	77: {
		1:   uint16(357),
		9:   uint16(357),
		10:  uint16(357),
		31:  uint16(357),
		32:  uint16(357),
		33:  uint16(357),
		34:  uint16(357),
		36:  uint16(357),
		38:  uint16(357),
		40:  uint16(357),
		53:  uint16(357),
		54:  uint16(357),
		55:  uint16(357),
		56:  uint16(357),
		57:  uint16(357),
		58:  uint16(357),
		59:  uint16(357),
		60:  uint16(357),
		61:  uint16(357),
		63:  uint16(357),
		64:  uint16(357),
		65:  uint16(357),
		66:  uint16(357),
		71:  uint16(357),
		72:  uint16(357),
		73:  uint16(357),
		74:  uint16(357),
		75:  uint16(357),
		76:  uint16(357),
		77:  uint16(357),
		78:  uint16(357),
		79:  uint16(357),
		80:  uint16(357),
		81:  uint16(357),
		82:  uint16(357),
		83:  uint16(357),
		84:  uint16(357),
		85:  uint16(357),
		86:  uint16(357),
		87:  uint16(357),
		88:  uint16(357),
		89:  uint16(357),
		90:  uint16(357),
		91:  uint16(357),
		92:  uint16(357),
		93:  uint16(357),
		94:  uint16(357),
		95:  uint16(357),
		96:  uint16(357),
		97:  uint16(357),
		98:  uint16(357),
		99:  uint16(357),
		100: uint16(357),
		101: uint16(357),
		102: uint16(357),
		103: uint16(357),
		104: uint16(357),
		105: uint16(357),
		106: uint16(357),
		107: uint16(357),
		108: uint16(357),
		109: uint16(357),
		110: uint16(357),
		111: uint16(357),
		125: uint16(3),
		128: uint16(63),
	},
	78: {
		1:   uint16(327),
		9:   uint16(327),
		10:  uint16(327),
		31:  uint16(327),
		32:  uint16(327),
		33:  uint16(327),
		34:  uint16(327),
		36:  uint16(327),
		38:  uint16(327),
		40:  uint16(327),
		53:  uint16(327),
		54:  uint16(327),
		55:  uint16(327),
		56:  uint16(327),
		57:  uint16(327),
		58:  uint16(327),
		59:  uint16(327),
		60:  uint16(327),
		61:  uint16(327),
		63:  uint16(327),
		64:  uint16(327),
		65:  uint16(327),
		66:  uint16(327),
		71:  uint16(327),
		72:  uint16(327),
		73:  uint16(327),
		74:  uint16(327),
		75:  uint16(327),
		76:  uint16(327),
		77:  uint16(327),
		78:  uint16(327),
		79:  uint16(327),
		80:  uint16(327),
		81:  uint16(327),
		82:  uint16(327),
		83:  uint16(327),
		84:  uint16(327),
		85:  uint16(327),
		86:  uint16(327),
		87:  uint16(327),
		88:  uint16(327),
		89:  uint16(327),
		90:  uint16(327),
		91:  uint16(327),
		92:  uint16(327),
		93:  uint16(327),
		94:  uint16(327),
		95:  uint16(327),
		96:  uint16(327),
		97:  uint16(327),
		98:  uint16(327),
		99:  uint16(327),
		100: uint16(327),
		101: uint16(327),
		102: uint16(327),
		103: uint16(327),
		104: uint16(327),
		105: uint16(327),
		106: uint16(327),
		107: uint16(327),
		108: uint16(327),
		109: uint16(327),
		110: uint16(327),
		111: uint16(327),
		125: uint16(3),
		128: uint16(331),
	},
	79: {
		1:   uint16(407),
		9:   uint16(407),
		10:  uint16(407),
		31:  uint16(407),
		32:  uint16(407),
		33:  uint16(407),
		34:  uint16(407),
		36:  uint16(407),
		38:  uint16(407),
		40:  uint16(407),
		53:  uint16(407),
		54:  uint16(407),
		55:  uint16(407),
		56:  uint16(407),
		57:  uint16(407),
		58:  uint16(407),
		59:  uint16(407),
		60:  uint16(407),
		61:  uint16(407),
		63:  uint16(407),
		64:  uint16(407),
		65:  uint16(407),
		66:  uint16(407),
		71:  uint16(407),
		72:  uint16(407),
		73:  uint16(407),
		74:  uint16(407),
		75:  uint16(407),
		76:  uint16(407),
		77:  uint16(407),
		78:  uint16(407),
		79:  uint16(407),
		80:  uint16(407),
		81:  uint16(407),
		82:  uint16(407),
		83:  uint16(407),
		84:  uint16(407),
		85:  uint16(407),
		86:  uint16(407),
		87:  uint16(407),
		88:  uint16(407),
		89:  uint16(407),
		90:  uint16(407),
		91:  uint16(407),
		92:  uint16(407),
		93:  uint16(407),
		94:  uint16(407),
		95:  uint16(407),
		96:  uint16(407),
		97:  uint16(407),
		98:  uint16(407),
		99:  uint16(407),
		100: uint16(407),
		101: uint16(407),
		102: uint16(407),
		103: uint16(407),
		104: uint16(407),
		105: uint16(407),
		106: uint16(407),
		107: uint16(407),
		108: uint16(407),
		109: uint16(407),
		110: uint16(407),
		111: uint16(407),
		125: uint16(3),
		128: uint16(409),
	},
	80: {
		1:   uint16(411),
		9:   uint16(411),
		10:  uint16(411),
		31:  uint16(411),
		32:  uint16(411),
		33:  uint16(411),
		34:  uint16(411),
		36:  uint16(411),
		38:  uint16(411),
		40:  uint16(411),
		53:  uint16(411),
		54:  uint16(411),
		55:  uint16(411),
		56:  uint16(411),
		57:  uint16(411),
		58:  uint16(411),
		59:  uint16(411),
		60:  uint16(411),
		61:  uint16(411),
		63:  uint16(411),
		64:  uint16(411),
		65:  uint16(411),
		66:  uint16(411),
		71:  uint16(411),
		72:  uint16(411),
		73:  uint16(411),
		74:  uint16(411),
		75:  uint16(411),
		76:  uint16(411),
		77:  uint16(411),
		78:  uint16(411),
		79:  uint16(411),
		80:  uint16(411),
		81:  uint16(411),
		82:  uint16(411),
		83:  uint16(411),
		84:  uint16(411),
		85:  uint16(411),
		86:  uint16(411),
		87:  uint16(411),
		88:  uint16(411),
		89:  uint16(411),
		90:  uint16(411),
		91:  uint16(411),
		92:  uint16(411),
		93:  uint16(411),
		94:  uint16(411),
		95:  uint16(411),
		96:  uint16(411),
		97:  uint16(411),
		98:  uint16(411),
		99:  uint16(411),
		100: uint16(411),
		101: uint16(411),
		102: uint16(411),
		103: uint16(411),
		104: uint16(411),
		105: uint16(411),
		106: uint16(411),
		107: uint16(411),
		108: uint16(411),
		109: uint16(411),
		110: uint16(411),
		111: uint16(411),
		125: uint16(3),
		128: uint16(413),
	},
	81: {
		1:   uint16(415),
		9:   uint16(415),
		10:  uint16(415),
		31:  uint16(415),
		32:  uint16(415),
		33:  uint16(415),
		34:  uint16(415),
		36:  uint16(415),
		38:  uint16(415),
		40:  uint16(415),
		53:  uint16(415),
		54:  uint16(415),
		55:  uint16(415),
		56:  uint16(415),
		57:  uint16(415),
		58:  uint16(415),
		59:  uint16(415),
		60:  uint16(415),
		61:  uint16(415),
		63:  uint16(415),
		64:  uint16(415),
		65:  uint16(415),
		66:  uint16(415),
		71:  uint16(415),
		72:  uint16(415),
		73:  uint16(415),
		74:  uint16(415),
		75:  uint16(415),
		76:  uint16(415),
		77:  uint16(415),
		78:  uint16(415),
		79:  uint16(415),
		80:  uint16(415),
		81:  uint16(415),
		82:  uint16(415),
		83:  uint16(415),
		84:  uint16(415),
		85:  uint16(415),
		86:  uint16(415),
		87:  uint16(415),
		88:  uint16(415),
		89:  uint16(415),
		90:  uint16(415),
		91:  uint16(415),
		92:  uint16(415),
		93:  uint16(415),
		94:  uint16(415),
		95:  uint16(415),
		96:  uint16(415),
		97:  uint16(415),
		98:  uint16(415),
		99:  uint16(415),
		100: uint16(415),
		101: uint16(415),
		102: uint16(415),
		103: uint16(415),
		104: uint16(415),
		105: uint16(415),
		106: uint16(415),
		107: uint16(415),
		108: uint16(415),
		109: uint16(415),
		110: uint16(415),
		111: uint16(415),
		125: uint16(3),
		128: uint16(417),
	},
	82: {
		1:   uint16(419),
		9:   uint16(419),
		10:  uint16(419),
		31:  uint16(419),
		32:  uint16(419),
		33:  uint16(419),
		34:  uint16(419),
		36:  uint16(419),
		38:  uint16(419),
		40:  uint16(419),
		53:  uint16(419),
		54:  uint16(419),
		55:  uint16(419),
		56:  uint16(419),
		57:  uint16(419),
		58:  uint16(419),
		59:  uint16(419),
		60:  uint16(419),
		61:  uint16(419),
		63:  uint16(419),
		64:  uint16(419),
		65:  uint16(419),
		66:  uint16(419),
		71:  uint16(419),
		72:  uint16(419),
		73:  uint16(419),
		74:  uint16(419),
		75:  uint16(419),
		76:  uint16(419),
		77:  uint16(419),
		78:  uint16(419),
		79:  uint16(419),
		80:  uint16(419),
		81:  uint16(419),
		82:  uint16(419),
		83:  uint16(419),
		84:  uint16(419),
		85:  uint16(419),
		86:  uint16(419),
		87:  uint16(419),
		88:  uint16(419),
		89:  uint16(419),
		90:  uint16(419),
		91:  uint16(419),
		92:  uint16(419),
		93:  uint16(419),
		94:  uint16(419),
		95:  uint16(419),
		96:  uint16(419),
		97:  uint16(419),
		98:  uint16(419),
		99:  uint16(419),
		100: uint16(419),
		101: uint16(419),
		102: uint16(419),
		103: uint16(419),
		104: uint16(419),
		105: uint16(419),
		106: uint16(419),
		107: uint16(419),
		108: uint16(419),
		109: uint16(419),
		110: uint16(419),
		111: uint16(419),
		125: uint16(3),
		128: uint16(421),
	},
	83: {
		1:   uint16(423),
		9:   uint16(423),
		10:  uint16(423),
		31:  uint16(423),
		32:  uint16(423),
		33:  uint16(423),
		34:  uint16(423),
		36:  uint16(423),
		38:  uint16(423),
		40:  uint16(423),
		53:  uint16(423),
		54:  uint16(423),
		55:  uint16(423),
		56:  uint16(423),
		57:  uint16(423),
		58:  uint16(423),
		59:  uint16(423),
		60:  uint16(423),
		61:  uint16(423),
		63:  uint16(423),
		64:  uint16(423),
		65:  uint16(423),
		66:  uint16(423),
		71:  uint16(423),
		72:  uint16(423),
		73:  uint16(423),
		74:  uint16(423),
		75:  uint16(423),
		76:  uint16(423),
		77:  uint16(423),
		78:  uint16(423),
		79:  uint16(423),
		80:  uint16(423),
		81:  uint16(423),
		82:  uint16(423),
		83:  uint16(423),
		84:  uint16(423),
		85:  uint16(423),
		86:  uint16(423),
		87:  uint16(423),
		88:  uint16(423),
		89:  uint16(423),
		90:  uint16(423),
		91:  uint16(423),
		92:  uint16(423),
		93:  uint16(423),
		94:  uint16(423),
		95:  uint16(423),
		96:  uint16(423),
		97:  uint16(423),
		98:  uint16(423),
		99:  uint16(423),
		100: uint16(423),
		101: uint16(423),
		102: uint16(423),
		103: uint16(423),
		104: uint16(423),
		105: uint16(423),
		106: uint16(423),
		107: uint16(423),
		108: uint16(423),
		109: uint16(423),
		110: uint16(423),
		111: uint16(423),
		125: uint16(3),
		128: uint16(425),
	},
	84: {
		1:   uint16(427),
		9:   uint16(427),
		10:  uint16(427),
		31:  uint16(427),
		32:  uint16(427),
		33:  uint16(427),
		34:  uint16(427),
		36:  uint16(427),
		38:  uint16(427),
		40:  uint16(427),
		53:  uint16(427),
		54:  uint16(427),
		55:  uint16(427),
		56:  uint16(427),
		57:  uint16(427),
		58:  uint16(427),
		59:  uint16(427),
		60:  uint16(427),
		61:  uint16(427),
		63:  uint16(427),
		64:  uint16(427),
		65:  uint16(427),
		66:  uint16(427),
		71:  uint16(427),
		72:  uint16(427),
		73:  uint16(427),
		74:  uint16(427),
		75:  uint16(427),
		76:  uint16(427),
		77:  uint16(427),
		78:  uint16(427),
		79:  uint16(427),
		80:  uint16(427),
		81:  uint16(427),
		82:  uint16(427),
		83:  uint16(427),
		84:  uint16(427),
		85:  uint16(427),
		86:  uint16(427),
		87:  uint16(427),
		88:  uint16(427),
		89:  uint16(427),
		90:  uint16(427),
		91:  uint16(427),
		92:  uint16(427),
		93:  uint16(427),
		94:  uint16(427),
		95:  uint16(427),
		96:  uint16(427),
		97:  uint16(427),
		98:  uint16(427),
		99:  uint16(427),
		100: uint16(427),
		101: uint16(427),
		102: uint16(427),
		103: uint16(427),
		104: uint16(427),
		105: uint16(427),
		106: uint16(427),
		107: uint16(427),
		108: uint16(427),
		109: uint16(427),
		110: uint16(427),
		111: uint16(427),
		125: uint16(3),
		128: uint16(429),
	},
	85: {
		1:   uint16(431),
		9:   uint16(431),
		10:  uint16(431),
		31:  uint16(431),
		32:  uint16(431),
		33:  uint16(431),
		34:  uint16(431),
		36:  uint16(431),
		38:  uint16(431),
		40:  uint16(431),
		53:  uint16(431),
		54:  uint16(431),
		55:  uint16(431),
		56:  uint16(431),
		57:  uint16(431),
		58:  uint16(431),
		59:  uint16(431),
		60:  uint16(431),
		61:  uint16(431),
		63:  uint16(431),
		64:  uint16(431),
		65:  uint16(431),
		66:  uint16(431),
		71:  uint16(431),
		72:  uint16(431),
		73:  uint16(431),
		74:  uint16(431),
		75:  uint16(431),
		76:  uint16(431),
		77:  uint16(431),
		78:  uint16(431),
		79:  uint16(431),
		80:  uint16(431),
		81:  uint16(431),
		82:  uint16(431),
		83:  uint16(431),
		84:  uint16(431),
		85:  uint16(431),
		86:  uint16(431),
		87:  uint16(431),
		88:  uint16(431),
		89:  uint16(431),
		90:  uint16(431),
		91:  uint16(431),
		92:  uint16(431),
		93:  uint16(431),
		94:  uint16(431),
		95:  uint16(431),
		96:  uint16(431),
		97:  uint16(431),
		98:  uint16(431),
		99:  uint16(431),
		100: uint16(431),
		101: uint16(431),
		102: uint16(431),
		103: uint16(431),
		104: uint16(431),
		105: uint16(431),
		106: uint16(431),
		107: uint16(431),
		108: uint16(431),
		109: uint16(431),
		110: uint16(431),
		111: uint16(431),
		125: uint16(3),
		128: uint16(433),
	},
	86: {
		1:   uint16(435),
		9:   uint16(435),
		10:  uint16(435),
		31:  uint16(435),
		32:  uint16(435),
		33:  uint16(435),
		34:  uint16(435),
		36:  uint16(435),
		38:  uint16(435),
		40:  uint16(435),
		53:  uint16(435),
		54:  uint16(435),
		55:  uint16(435),
		56:  uint16(435),
		57:  uint16(435),
		58:  uint16(435),
		59:  uint16(435),
		60:  uint16(435),
		61:  uint16(435),
		63:  uint16(435),
		64:  uint16(435),
		65:  uint16(435),
		66:  uint16(435),
		71:  uint16(435),
		72:  uint16(435),
		73:  uint16(435),
		74:  uint16(435),
		75:  uint16(435),
		76:  uint16(435),
		77:  uint16(435),
		78:  uint16(435),
		79:  uint16(435),
		80:  uint16(435),
		81:  uint16(435),
		82:  uint16(435),
		83:  uint16(435),
		84:  uint16(435),
		85:  uint16(435),
		86:  uint16(435),
		87:  uint16(435),
		88:  uint16(435),
		89:  uint16(435),
		90:  uint16(435),
		91:  uint16(435),
		92:  uint16(435),
		93:  uint16(435),
		94:  uint16(435),
		95:  uint16(435),
		96:  uint16(435),
		97:  uint16(435),
		98:  uint16(435),
		99:  uint16(435),
		100: uint16(435),
		101: uint16(435),
		102: uint16(435),
		103: uint16(435),
		104: uint16(435),
		105: uint16(435),
		106: uint16(435),
		107: uint16(435),
		108: uint16(435),
		109: uint16(435),
		110: uint16(435),
		111: uint16(435),
		125: uint16(3),
		128: uint16(437),
	},
	87: {
		1:   uint16(439),
		9:   uint16(439),
		10:  uint16(439),
		31:  uint16(439),
		32:  uint16(439),
		33:  uint16(439),
		34:  uint16(439),
		36:  uint16(439),
		38:  uint16(439),
		40:  uint16(439),
		53:  uint16(439),
		54:  uint16(439),
		55:  uint16(439),
		56:  uint16(439),
		57:  uint16(439),
		58:  uint16(439),
		59:  uint16(439),
		60:  uint16(439),
		61:  uint16(439),
		63:  uint16(439),
		64:  uint16(439),
		65:  uint16(439),
		66:  uint16(439),
		71:  uint16(439),
		72:  uint16(439),
		73:  uint16(439),
		74:  uint16(439),
		75:  uint16(439),
		76:  uint16(439),
		77:  uint16(439),
		78:  uint16(439),
		79:  uint16(439),
		80:  uint16(439),
		81:  uint16(439),
		82:  uint16(439),
		83:  uint16(439),
		84:  uint16(439),
		85:  uint16(439),
		86:  uint16(439),
		87:  uint16(439),
		88:  uint16(439),
		89:  uint16(439),
		90:  uint16(439),
		91:  uint16(439),
		92:  uint16(439),
		93:  uint16(439),
		94:  uint16(439),
		95:  uint16(439),
		96:  uint16(439),
		97:  uint16(439),
		98:  uint16(439),
		99:  uint16(439),
		100: uint16(439),
		101: uint16(439),
		102: uint16(439),
		103: uint16(439),
		104: uint16(439),
		105: uint16(439),
		106: uint16(439),
		107: uint16(439),
		108: uint16(439),
		109: uint16(439),
		110: uint16(439),
		111: uint16(439),
		125: uint16(3),
		128: uint16(441),
	},
	88: {
		1:   uint16(443),
		9:   uint16(443),
		10:  uint16(443),
		31:  uint16(443),
		32:  uint16(443),
		33:  uint16(443),
		34:  uint16(443),
		36:  uint16(443),
		38:  uint16(443),
		40:  uint16(443),
		53:  uint16(443),
		54:  uint16(443),
		55:  uint16(443),
		56:  uint16(443),
		57:  uint16(443),
		58:  uint16(443),
		59:  uint16(443),
		60:  uint16(443),
		61:  uint16(443),
		63:  uint16(443),
		64:  uint16(443),
		65:  uint16(443),
		66:  uint16(443),
		71:  uint16(443),
		72:  uint16(443),
		73:  uint16(443),
		74:  uint16(443),
		75:  uint16(443),
		76:  uint16(443),
		77:  uint16(443),
		78:  uint16(443),
		79:  uint16(443),
		80:  uint16(443),
		81:  uint16(443),
		82:  uint16(443),
		83:  uint16(443),
		84:  uint16(443),
		85:  uint16(443),
		86:  uint16(443),
		87:  uint16(443),
		88:  uint16(443),
		89:  uint16(443),
		90:  uint16(443),
		91:  uint16(443),
		92:  uint16(443),
		93:  uint16(443),
		94:  uint16(443),
		95:  uint16(443),
		96:  uint16(443),
		97:  uint16(443),
		98:  uint16(443),
		99:  uint16(443),
		100: uint16(443),
		101: uint16(443),
		102: uint16(443),
		103: uint16(443),
		104: uint16(443),
		105: uint16(443),
		106: uint16(443),
		107: uint16(443),
		108: uint16(443),
		109: uint16(443),
		110: uint16(443),
		111: uint16(443),
		125: uint16(3),
		128: uint16(445),
	},
	89: {
		1:   uint16(447),
		9:   uint16(447),
		10:  uint16(447),
		31:  uint16(447),
		32:  uint16(447),
		33:  uint16(447),
		34:  uint16(447),
		36:  uint16(447),
		38:  uint16(447),
		40:  uint16(447),
		53:  uint16(447),
		54:  uint16(447),
		55:  uint16(447),
		56:  uint16(447),
		57:  uint16(447),
		58:  uint16(447),
		59:  uint16(447),
		60:  uint16(447),
		61:  uint16(447),
		63:  uint16(447),
		64:  uint16(447),
		65:  uint16(447),
		66:  uint16(447),
		71:  uint16(447),
		72:  uint16(447),
		73:  uint16(447),
		74:  uint16(447),
		75:  uint16(447),
		76:  uint16(447),
		77:  uint16(447),
		78:  uint16(447),
		79:  uint16(447),
		80:  uint16(447),
		81:  uint16(447),
		82:  uint16(447),
		83:  uint16(447),
		84:  uint16(447),
		85:  uint16(447),
		86:  uint16(447),
		87:  uint16(447),
		88:  uint16(447),
		89:  uint16(447),
		90:  uint16(447),
		91:  uint16(447),
		92:  uint16(447),
		93:  uint16(447),
		94:  uint16(447),
		95:  uint16(447),
		96:  uint16(447),
		97:  uint16(447),
		98:  uint16(447),
		99:  uint16(447),
		100: uint16(447),
		101: uint16(447),
		102: uint16(447),
		103: uint16(447),
		104: uint16(447),
		105: uint16(447),
		106: uint16(447),
		107: uint16(447),
		108: uint16(447),
		109: uint16(447),
		110: uint16(447),
		111: uint16(447),
		125: uint16(3),
		128: uint16(449),
	},
	90: {
		1:   uint16(451),
		9:   uint16(451),
		10:  uint16(451),
		31:  uint16(451),
		32:  uint16(451),
		33:  uint16(451),
		34:  uint16(451),
		36:  uint16(451),
		38:  uint16(451),
		40:  uint16(451),
		53:  uint16(451),
		54:  uint16(451),
		55:  uint16(451),
		56:  uint16(451),
		57:  uint16(451),
		58:  uint16(451),
		59:  uint16(451),
		60:  uint16(451),
		61:  uint16(451),
		63:  uint16(451),
		64:  uint16(451),
		65:  uint16(451),
		66:  uint16(451),
		71:  uint16(451),
		72:  uint16(451),
		73:  uint16(451),
		74:  uint16(451),
		75:  uint16(451),
		76:  uint16(451),
		77:  uint16(451),
		78:  uint16(451),
		79:  uint16(451),
		80:  uint16(451),
		81:  uint16(451),
		82:  uint16(451),
		83:  uint16(451),
		84:  uint16(451),
		85:  uint16(451),
		86:  uint16(451),
		87:  uint16(451),
		88:  uint16(451),
		89:  uint16(451),
		90:  uint16(451),
		91:  uint16(451),
		92:  uint16(451),
		93:  uint16(451),
		94:  uint16(451),
		95:  uint16(451),
		96:  uint16(451),
		97:  uint16(451),
		98:  uint16(451),
		99:  uint16(451),
		100: uint16(451),
		101: uint16(451),
		102: uint16(451),
		103: uint16(451),
		104: uint16(451),
		105: uint16(451),
		106: uint16(451),
		107: uint16(451),
		108: uint16(451),
		109: uint16(451),
		110: uint16(451),
		111: uint16(451),
		125: uint16(3),
		128: uint16(453),
	},
	91: {
		1:   uint16(455),
		9:   uint16(455),
		10:  uint16(455),
		31:  uint16(455),
		32:  uint16(455),
		33:  uint16(455),
		34:  uint16(455),
		36:  uint16(455),
		38:  uint16(455),
		40:  uint16(455),
		53:  uint16(455),
		54:  uint16(455),
		55:  uint16(455),
		56:  uint16(455),
		57:  uint16(455),
		58:  uint16(455),
		59:  uint16(455),
		60:  uint16(455),
		61:  uint16(455),
		63:  uint16(455),
		64:  uint16(455),
		65:  uint16(455),
		66:  uint16(455),
		71:  uint16(455),
		72:  uint16(455),
		73:  uint16(455),
		74:  uint16(455),
		75:  uint16(455),
		76:  uint16(455),
		77:  uint16(455),
		78:  uint16(455),
		79:  uint16(455),
		80:  uint16(455),
		81:  uint16(455),
		82:  uint16(455),
		83:  uint16(455),
		84:  uint16(455),
		85:  uint16(455),
		86:  uint16(455),
		87:  uint16(455),
		88:  uint16(455),
		89:  uint16(455),
		90:  uint16(455),
		91:  uint16(455),
		92:  uint16(455),
		93:  uint16(455),
		94:  uint16(455),
		95:  uint16(455),
		96:  uint16(455),
		97:  uint16(455),
		98:  uint16(455),
		99:  uint16(455),
		100: uint16(455),
		101: uint16(455),
		102: uint16(455),
		103: uint16(455),
		104: uint16(455),
		105: uint16(455),
		106: uint16(455),
		107: uint16(455),
		108: uint16(455),
		109: uint16(455),
		110: uint16(455),
		111: uint16(455),
		125: uint16(3),
		128: uint16(457),
	},
	92: {
		1:   uint16(459),
		9:   uint16(459),
		10:  uint16(459),
		31:  uint16(459),
		32:  uint16(459),
		33:  uint16(459),
		34:  uint16(459),
		36:  uint16(459),
		38:  uint16(459),
		40:  uint16(459),
		53:  uint16(459),
		54:  uint16(459),
		55:  uint16(459),
		56:  uint16(459),
		57:  uint16(459),
		58:  uint16(459),
		59:  uint16(459),
		60:  uint16(459),
		61:  uint16(459),
		63:  uint16(459),
		64:  uint16(459),
		65:  uint16(459),
		66:  uint16(459),
		71:  uint16(459),
		72:  uint16(459),
		73:  uint16(459),
		74:  uint16(459),
		75:  uint16(459),
		76:  uint16(459),
		77:  uint16(459),
		78:  uint16(459),
		79:  uint16(459),
		80:  uint16(459),
		81:  uint16(459),
		82:  uint16(459),
		83:  uint16(459),
		84:  uint16(459),
		85:  uint16(459),
		86:  uint16(459),
		87:  uint16(459),
		88:  uint16(459),
		89:  uint16(459),
		90:  uint16(459),
		91:  uint16(459),
		92:  uint16(459),
		93:  uint16(459),
		94:  uint16(459),
		95:  uint16(459),
		96:  uint16(459),
		97:  uint16(459),
		98:  uint16(459),
		99:  uint16(459),
		100: uint16(459),
		101: uint16(459),
		102: uint16(459),
		103: uint16(459),
		104: uint16(459),
		105: uint16(459),
		106: uint16(459),
		107: uint16(459),
		108: uint16(459),
		109: uint16(459),
		110: uint16(459),
		111: uint16(459),
		125: uint16(3),
		128: uint16(461),
	},
	93: {
		1:   uint16(463),
		9:   uint16(463),
		10:  uint16(463),
		31:  uint16(463),
		32:  uint16(463),
		33:  uint16(463),
		34:  uint16(463),
		36:  uint16(463),
		38:  uint16(463),
		40:  uint16(463),
		53:  uint16(463),
		54:  uint16(463),
		55:  uint16(463),
		56:  uint16(463),
		57:  uint16(463),
		58:  uint16(463),
		59:  uint16(463),
		60:  uint16(463),
		61:  uint16(463),
		63:  uint16(463),
		64:  uint16(463),
		65:  uint16(463),
		66:  uint16(463),
		71:  uint16(463),
		72:  uint16(463),
		73:  uint16(463),
		74:  uint16(463),
		75:  uint16(463),
		76:  uint16(463),
		77:  uint16(463),
		78:  uint16(463),
		79:  uint16(463),
		80:  uint16(463),
		81:  uint16(463),
		82:  uint16(463),
		83:  uint16(463),
		84:  uint16(463),
		85:  uint16(463),
		86:  uint16(463),
		87:  uint16(463),
		88:  uint16(463),
		89:  uint16(463),
		90:  uint16(463),
		91:  uint16(463),
		92:  uint16(463),
		93:  uint16(463),
		94:  uint16(463),
		95:  uint16(463),
		96:  uint16(463),
		97:  uint16(463),
		98:  uint16(463),
		99:  uint16(463),
		100: uint16(463),
		101: uint16(463),
		102: uint16(463),
		103: uint16(463),
		104: uint16(463),
		105: uint16(463),
		106: uint16(463),
		107: uint16(463),
		108: uint16(463),
		109: uint16(463),
		110: uint16(463),
		111: uint16(463),
		125: uint16(3),
		128: uint16(465),
	},
	94: {
		1:   uint16(321),
		9:   uint16(321),
		10:  uint16(321),
		31:  uint16(321),
		32:  uint16(321),
		33:  uint16(321),
		34:  uint16(321),
		36:  uint16(321),
		38:  uint16(321),
		40:  uint16(321),
		53:  uint16(321),
		54:  uint16(321),
		55:  uint16(321),
		56:  uint16(321),
		57:  uint16(321),
		58:  uint16(321),
		59:  uint16(321),
		60:  uint16(321),
		61:  uint16(321),
		63:  uint16(321),
		64:  uint16(321),
		65:  uint16(321),
		66:  uint16(321),
		71:  uint16(321),
		72:  uint16(321),
		73:  uint16(321),
		74:  uint16(321),
		75:  uint16(321),
		76:  uint16(321),
		77:  uint16(321),
		78:  uint16(321),
		79:  uint16(321),
		80:  uint16(321),
		81:  uint16(321),
		82:  uint16(321),
		83:  uint16(321),
		84:  uint16(321),
		85:  uint16(321),
		86:  uint16(321),
		87:  uint16(321),
		88:  uint16(321),
		89:  uint16(321),
		90:  uint16(321),
		91:  uint16(321),
		92:  uint16(321),
		93:  uint16(321),
		94:  uint16(321),
		95:  uint16(321),
		96:  uint16(321),
		97:  uint16(321),
		98:  uint16(321),
		99:  uint16(321),
		100: uint16(321),
		101: uint16(321),
		102: uint16(321),
		103: uint16(321),
		104: uint16(321),
		105: uint16(321),
		106: uint16(321),
		107: uint16(321),
		108: uint16(321),
		109: uint16(321),
		110: uint16(321),
		111: uint16(321),
		125: uint16(3),
		128: uint16(325),
	},
}

var ts_small_parse_table = [6529]uint16_t{
	0:    uint16(14),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(9),
	5:    uint16(1),
	6:    uint16(sym_identifier),
	7:    uint16(43),
	8:    uint16(1),
	9:    uint16(anon_sym_mux),
	10:   uint16(45),
	11:   uint16(1),
	12:   uint16(anon_sym_validif),
	13:   uint16(143),
	14:   uint16(1),
	15:   uint16(sym_primop),
	16:   uint16(210),
	17:   uint16(1),
	18:   uint16(sym_expression),
	19:   uint16(258),
	20:   uint16(1),
	21:   uint16(sym_number),
	22:   uint16(433),
	23:   uint16(1),
	24:   uint16(sym_litType),
	25:   uint16(13),
	26:   uint16(2),
	27:   uint16(anon_sym_UInt),
	28:   uint16(anon_sym_SInt),
	29:   uint16(467),
	30:   uint16(2),
	31:   uint16(anon_sym_0),
	32:   uint16(aux_sym_uint_token1),
	33:   uint16(469),
	34:   uint16(2),
	35:   uint16(anon_sym_PLUS),
	36:   uint16(anon_sym_DASH),
	37:   uint16(222),
	38:   uint16(2),
	39:   uint16(sym_uint),
	40:   uint16(sym_sint),
	41:   uint16(160),
	42:   uint16(7),
	43:   uint16(sym_literal),
	44:   uint16(sym_sub_field),
	45:   uint16(sym_sub_index),
	46:   uint16(sym_sub_access),
	47:   uint16(sym_mux),
	48:   uint16(sym_conditionally_valid),
	49:   uint16(sym_primitive_operation),
	50:   uint16(47),
	51:   uint16(39),
	52:   uint16(anon_sym_add),
	53:   uint16(anon_sym_sub),
	54:   uint16(anon_sym_mul),
	55:   uint16(anon_sym_div),
	56:   uint16(anon_sym_rem),
	57:   uint16(anon_sym_lt),
	58:   uint16(anon_sym_leq),
	59:   uint16(anon_sym_gt),
	60:   uint16(anon_sym_geq),
	61:   uint16(anon_sym_eq),
	62:   uint16(anon_sym_neq),
	63:   uint16(anon_sym_pad),
	64:   uint16(anon_sym_asUInt),
	65:   uint16(anon_sym_asAsyncReset),
	66:   uint16(anon_sym_asSInt),
	67:   uint16(anon_sym_asClock),
	68:   uint16(anon_sym_shl),
	69:   uint16(anon_sym_shr),
	70:   uint16(anon_sym_dshl),
	71:   uint16(anon_sym_dshlw),
	72:   uint16(anon_sym_dshr),
	73:   uint16(anon_sym_dshrw),
	74:   uint16(anon_sym_cvt),
	75:   uint16(anon_sym_neg),
	76:   uint16(anon_sym_not),
	77:   uint16(anon_sym_and),
	78:   uint16(anon_sym_or),
	79:   uint16(anon_sym_xor),
	80:   uint16(anon_sym_andr),
	81:   uint16(anon_sym_orr),
	82:   uint16(anon_sym_xorr),
	83:   uint16(anon_sym_cat),
	84:   uint16(anon_sym_bits),
	85:   uint16(anon_sym_head),
	86:   uint16(anon_sym_tail),
	87:   uint16(anon_sym_asFixedPoint),
	88:   uint16(anon_sym_bpshl),
	89:   uint16(anon_sym_bpshr),
	90:   uint16(anon_sym_bpset),
	91:   uint16(14),
	92:   uint16(3),
	93:   uint16(1),
	94:   uint16(sym_comment),
	95:   uint16(9),
	96:   uint16(1),
	97:   uint16(sym_identifier),
	98:   uint16(43),
	99:   uint16(1),
	100:  uint16(anon_sym_mux),
	101:  uint16(45),
	102:  uint16(1),
	103:  uint16(anon_sym_validif),
	104:  uint16(143),
	105:  uint16(1),
	106:  uint16(sym_primop),
	107:  uint16(204),
	108:  uint16(1),
	109:  uint16(sym_expression),
	110:  uint16(278),
	111:  uint16(1),
	112:  uint16(sym_number),
	113:  uint16(433),
	114:  uint16(1),
	115:  uint16(sym_litType),
	116:  uint16(13),
	117:  uint16(2),
	118:  uint16(anon_sym_UInt),
	119:  uint16(anon_sym_SInt),
	120:  uint16(467),
	121:  uint16(2),
	122:  uint16(anon_sym_0),
	123:  uint16(aux_sym_uint_token1),
	124:  uint16(469),
	125:  uint16(2),
	126:  uint16(anon_sym_PLUS),
	127:  uint16(anon_sym_DASH),
	128:  uint16(222),
	129:  uint16(2),
	130:  uint16(sym_uint),
	131:  uint16(sym_sint),
	132:  uint16(160),
	133:  uint16(7),
	134:  uint16(sym_literal),
	135:  uint16(sym_sub_field),
	136:  uint16(sym_sub_index),
	137:  uint16(sym_sub_access),
	138:  uint16(sym_mux),
	139:  uint16(sym_conditionally_valid),
	140:  uint16(sym_primitive_operation),
	141:  uint16(47),
	142:  uint16(39),
	143:  uint16(anon_sym_add),
	144:  uint16(anon_sym_sub),
	145:  uint16(anon_sym_mul),
	146:  uint16(anon_sym_div),
	147:  uint16(anon_sym_rem),
	148:  uint16(anon_sym_lt),
	149:  uint16(anon_sym_leq),
	150:  uint16(anon_sym_gt),
	151:  uint16(anon_sym_geq),
	152:  uint16(anon_sym_eq),
	153:  uint16(anon_sym_neq),
	154:  uint16(anon_sym_pad),
	155:  uint16(anon_sym_asUInt),
	156:  uint16(anon_sym_asAsyncReset),
	157:  uint16(anon_sym_asSInt),
	158:  uint16(anon_sym_asClock),
	159:  uint16(anon_sym_shl),
	160:  uint16(anon_sym_shr),
	161:  uint16(anon_sym_dshl),
	162:  uint16(anon_sym_dshlw),
	163:  uint16(anon_sym_dshr),
	164:  uint16(anon_sym_dshrw),
	165:  uint16(anon_sym_cvt),
	166:  uint16(anon_sym_neg),
	167:  uint16(anon_sym_not),
	168:  uint16(anon_sym_and),
	169:  uint16(anon_sym_or),
	170:  uint16(anon_sym_xor),
	171:  uint16(anon_sym_andr),
	172:  uint16(anon_sym_orr),
	173:  uint16(anon_sym_xorr),
	174:  uint16(anon_sym_cat),
	175:  uint16(anon_sym_bits),
	176:  uint16(anon_sym_head),
	177:  uint16(anon_sym_tail),
	178:  uint16(anon_sym_asFixedPoint),
	179:  uint16(anon_sym_bpshl),
	180:  uint16(anon_sym_bpshr),
	181:  uint16(anon_sym_bpset),
	182:  uint16(14),
	183:  uint16(3),
	184:  uint16(1),
	185:  uint16(sym_comment),
	186:  uint16(9),
	187:  uint16(1),
	188:  uint16(sym_identifier),
	189:  uint16(43),
	190:  uint16(1),
	191:  uint16(anon_sym_mux),
	192:  uint16(45),
	193:  uint16(1),
	194:  uint16(anon_sym_validif),
	195:  uint16(143),
	196:  uint16(1),
	197:  uint16(sym_primop),
	198:  uint16(254),
	199:  uint16(1),
	200:  uint16(sym_expression),
	201:  uint16(419),
	202:  uint16(1),
	203:  uint16(sym_number),
	204:  uint16(433),
	205:  uint16(1),
	206:  uint16(sym_litType),
	207:  uint16(13),
	208:  uint16(2),
	209:  uint16(anon_sym_UInt),
	210:  uint16(anon_sym_SInt),
	211:  uint16(467),
	212:  uint16(2),
	213:  uint16(anon_sym_0),
	214:  uint16(aux_sym_uint_token1),
	215:  uint16(469),
	216:  uint16(2),
	217:  uint16(anon_sym_PLUS),
	218:  uint16(anon_sym_DASH),
	219:  uint16(222),
	220:  uint16(2),
	221:  uint16(sym_uint),
	222:  uint16(sym_sint),
	223:  uint16(160),
	224:  uint16(7),
	225:  uint16(sym_literal),
	226:  uint16(sym_sub_field),
	227:  uint16(sym_sub_index),
	228:  uint16(sym_sub_access),
	229:  uint16(sym_mux),
	230:  uint16(sym_conditionally_valid),
	231:  uint16(sym_primitive_operation),
	232:  uint16(47),
	233:  uint16(39),
	234:  uint16(anon_sym_add),
	235:  uint16(anon_sym_sub),
	236:  uint16(anon_sym_mul),
	237:  uint16(anon_sym_div),
	238:  uint16(anon_sym_rem),
	239:  uint16(anon_sym_lt),
	240:  uint16(anon_sym_leq),
	241:  uint16(anon_sym_gt),
	242:  uint16(anon_sym_geq),
	243:  uint16(anon_sym_eq),
	244:  uint16(anon_sym_neq),
	245:  uint16(anon_sym_pad),
	246:  uint16(anon_sym_asUInt),
	247:  uint16(anon_sym_asAsyncReset),
	248:  uint16(anon_sym_asSInt),
	249:  uint16(anon_sym_asClock),
	250:  uint16(anon_sym_shl),
	251:  uint16(anon_sym_shr),
	252:  uint16(anon_sym_dshl),
	253:  uint16(anon_sym_dshlw),
	254:  uint16(anon_sym_dshr),
	255:  uint16(anon_sym_dshrw),
	256:  uint16(anon_sym_cvt),
	257:  uint16(anon_sym_neg),
	258:  uint16(anon_sym_not),
	259:  uint16(anon_sym_and),
	260:  uint16(anon_sym_or),
	261:  uint16(anon_sym_xor),
	262:  uint16(anon_sym_andr),
	263:  uint16(anon_sym_orr),
	264:  uint16(anon_sym_xorr),
	265:  uint16(anon_sym_cat),
	266:  uint16(anon_sym_bits),
	267:  uint16(anon_sym_head),
	268:  uint16(anon_sym_tail),
	269:  uint16(anon_sym_asFixedPoint),
	270:  uint16(anon_sym_bpshl),
	271:  uint16(anon_sym_bpshr),
	272:  uint16(anon_sym_bpset),
	273:  uint16(14),
	274:  uint16(3),
	275:  uint16(1),
	276:  uint16(sym_comment),
	277:  uint16(9),
	278:  uint16(1),
	279:  uint16(sym_identifier),
	280:  uint16(43),
	281:  uint16(1),
	282:  uint16(anon_sym_mux),
	283:  uint16(45),
	284:  uint16(1),
	285:  uint16(anon_sym_validif),
	286:  uint16(143),
	287:  uint16(1),
	288:  uint16(sym_primop),
	289:  uint16(220),
	290:  uint16(1),
	291:  uint16(sym_expression),
	292:  uint16(319),
	293:  uint16(1),
	294:  uint16(sym_number),
	295:  uint16(433),
	296:  uint16(1),
	297:  uint16(sym_litType),
	298:  uint16(13),
	299:  uint16(2),
	300:  uint16(anon_sym_UInt),
	301:  uint16(anon_sym_SInt),
	302:  uint16(467),
	303:  uint16(2),
	304:  uint16(anon_sym_0),
	305:  uint16(aux_sym_uint_token1),
	306:  uint16(469),
	307:  uint16(2),
	308:  uint16(anon_sym_PLUS),
	309:  uint16(anon_sym_DASH),
	310:  uint16(222),
	311:  uint16(2),
	312:  uint16(sym_uint),
	313:  uint16(sym_sint),
	314:  uint16(160),
	315:  uint16(7),
	316:  uint16(sym_literal),
	317:  uint16(sym_sub_field),
	318:  uint16(sym_sub_index),
	319:  uint16(sym_sub_access),
	320:  uint16(sym_mux),
	321:  uint16(sym_conditionally_valid),
	322:  uint16(sym_primitive_operation),
	323:  uint16(47),
	324:  uint16(39),
	325:  uint16(anon_sym_add),
	326:  uint16(anon_sym_sub),
	327:  uint16(anon_sym_mul),
	328:  uint16(anon_sym_div),
	329:  uint16(anon_sym_rem),
	330:  uint16(anon_sym_lt),
	331:  uint16(anon_sym_leq),
	332:  uint16(anon_sym_gt),
	333:  uint16(anon_sym_geq),
	334:  uint16(anon_sym_eq),
	335:  uint16(anon_sym_neq),
	336:  uint16(anon_sym_pad),
	337:  uint16(anon_sym_asUInt),
	338:  uint16(anon_sym_asAsyncReset),
	339:  uint16(anon_sym_asSInt),
	340:  uint16(anon_sym_asClock),
	341:  uint16(anon_sym_shl),
	342:  uint16(anon_sym_shr),
	343:  uint16(anon_sym_dshl),
	344:  uint16(anon_sym_dshlw),
	345:  uint16(anon_sym_dshr),
	346:  uint16(anon_sym_dshrw),
	347:  uint16(anon_sym_cvt),
	348:  uint16(anon_sym_neg),
	349:  uint16(anon_sym_not),
	350:  uint16(anon_sym_and),
	351:  uint16(anon_sym_or),
	352:  uint16(anon_sym_xor),
	353:  uint16(anon_sym_andr),
	354:  uint16(anon_sym_orr),
	355:  uint16(anon_sym_xorr),
	356:  uint16(anon_sym_cat),
	357:  uint16(anon_sym_bits),
	358:  uint16(anon_sym_head),
	359:  uint16(anon_sym_tail),
	360:  uint16(anon_sym_asFixedPoint),
	361:  uint16(anon_sym_bpshl),
	362:  uint16(anon_sym_bpshr),
	363:  uint16(anon_sym_bpset),
	364:  uint16(14),
	365:  uint16(3),
	366:  uint16(1),
	367:  uint16(sym_comment),
	368:  uint16(9),
	369:  uint16(1),
	370:  uint16(sym_identifier),
	371:  uint16(43),
	372:  uint16(1),
	373:  uint16(anon_sym_mux),
	374:  uint16(45),
	375:  uint16(1),
	376:  uint16(anon_sym_validif),
	377:  uint16(143),
	378:  uint16(1),
	379:  uint16(sym_primop),
	380:  uint16(279),
	381:  uint16(1),
	382:  uint16(sym_expression),
	383:  uint16(382),
	384:  uint16(1),
	385:  uint16(sym_number),
	386:  uint16(433),
	387:  uint16(1),
	388:  uint16(sym_litType),
	389:  uint16(13),
	390:  uint16(2),
	391:  uint16(anon_sym_UInt),
	392:  uint16(anon_sym_SInt),
	393:  uint16(467),
	394:  uint16(2),
	395:  uint16(anon_sym_0),
	396:  uint16(aux_sym_uint_token1),
	397:  uint16(469),
	398:  uint16(2),
	399:  uint16(anon_sym_PLUS),
	400:  uint16(anon_sym_DASH),
	401:  uint16(222),
	402:  uint16(2),
	403:  uint16(sym_uint),
	404:  uint16(sym_sint),
	405:  uint16(160),
	406:  uint16(7),
	407:  uint16(sym_literal),
	408:  uint16(sym_sub_field),
	409:  uint16(sym_sub_index),
	410:  uint16(sym_sub_access),
	411:  uint16(sym_mux),
	412:  uint16(sym_conditionally_valid),
	413:  uint16(sym_primitive_operation),
	414:  uint16(47),
	415:  uint16(39),
	416:  uint16(anon_sym_add),
	417:  uint16(anon_sym_sub),
	418:  uint16(anon_sym_mul),
	419:  uint16(anon_sym_div),
	420:  uint16(anon_sym_rem),
	421:  uint16(anon_sym_lt),
	422:  uint16(anon_sym_leq),
	423:  uint16(anon_sym_gt),
	424:  uint16(anon_sym_geq),
	425:  uint16(anon_sym_eq),
	426:  uint16(anon_sym_neq),
	427:  uint16(anon_sym_pad),
	428:  uint16(anon_sym_asUInt),
	429:  uint16(anon_sym_asAsyncReset),
	430:  uint16(anon_sym_asSInt),
	431:  uint16(anon_sym_asClock),
	432:  uint16(anon_sym_shl),
	433:  uint16(anon_sym_shr),
	434:  uint16(anon_sym_dshl),
	435:  uint16(anon_sym_dshlw),
	436:  uint16(anon_sym_dshr),
	437:  uint16(anon_sym_dshrw),
	438:  uint16(anon_sym_cvt),
	439:  uint16(anon_sym_neg),
	440:  uint16(anon_sym_not),
	441:  uint16(anon_sym_and),
	442:  uint16(anon_sym_or),
	443:  uint16(anon_sym_xor),
	444:  uint16(anon_sym_andr),
	445:  uint16(anon_sym_orr),
	446:  uint16(anon_sym_xorr),
	447:  uint16(anon_sym_cat),
	448:  uint16(anon_sym_bits),
	449:  uint16(anon_sym_head),
	450:  uint16(anon_sym_tail),
	451:  uint16(anon_sym_asFixedPoint),
	452:  uint16(anon_sym_bpshl),
	453:  uint16(anon_sym_bpshr),
	454:  uint16(anon_sym_bpset),
	455:  uint16(10),
	456:  uint16(3),
	457:  uint16(1),
	458:  uint16(sym_comment),
	459:  uint16(9),
	460:  uint16(1),
	461:  uint16(sym_identifier),
	462:  uint16(43),
	463:  uint16(1),
	464:  uint16(anon_sym_mux),
	465:  uint16(45),
	466:  uint16(1),
	467:  uint16(anon_sym_validif),
	468:  uint16(143),
	469:  uint16(1),
	470:  uint16(sym_primop),
	471:  uint16(244),
	472:  uint16(1),
	473:  uint16(sym_expression),
	474:  uint16(433),
	475:  uint16(1),
	476:  uint16(sym_litType),
	477:  uint16(13),
	478:  uint16(2),
	479:  uint16(anon_sym_UInt),
	480:  uint16(anon_sym_SInt),
	481:  uint16(160),
	482:  uint16(7),
	483:  uint16(sym_literal),
	484:  uint16(sym_sub_field),
	485:  uint16(sym_sub_index),
	486:  uint16(sym_sub_access),
	487:  uint16(sym_mux),
	488:  uint16(sym_conditionally_valid),
	489:  uint16(sym_primitive_operation),
	490:  uint16(47),
	491:  uint16(39),
	492:  uint16(anon_sym_add),
	493:  uint16(anon_sym_sub),
	494:  uint16(anon_sym_mul),
	495:  uint16(anon_sym_div),
	496:  uint16(anon_sym_rem),
	497:  uint16(anon_sym_lt),
	498:  uint16(anon_sym_leq),
	499:  uint16(anon_sym_gt),
	500:  uint16(anon_sym_geq),
	501:  uint16(anon_sym_eq),
	502:  uint16(anon_sym_neq),
	503:  uint16(anon_sym_pad),
	504:  uint16(anon_sym_asUInt),
	505:  uint16(anon_sym_asAsyncReset),
	506:  uint16(anon_sym_asSInt),
	507:  uint16(anon_sym_asClock),
	508:  uint16(anon_sym_shl),
	509:  uint16(anon_sym_shr),
	510:  uint16(anon_sym_dshl),
	511:  uint16(anon_sym_dshlw),
	512:  uint16(anon_sym_dshr),
	513:  uint16(anon_sym_dshrw),
	514:  uint16(anon_sym_cvt),
	515:  uint16(anon_sym_neg),
	516:  uint16(anon_sym_not),
	517:  uint16(anon_sym_and),
	518:  uint16(anon_sym_or),
	519:  uint16(anon_sym_xor),
	520:  uint16(anon_sym_andr),
	521:  uint16(anon_sym_orr),
	522:  uint16(anon_sym_xorr),
	523:  uint16(anon_sym_cat),
	524:  uint16(anon_sym_bits),
	525:  uint16(anon_sym_head),
	526:  uint16(anon_sym_tail),
	527:  uint16(anon_sym_asFixedPoint),
	528:  uint16(anon_sym_bpshl),
	529:  uint16(anon_sym_bpshr),
	530:  uint16(anon_sym_bpset),
	531:  uint16(10),
	532:  uint16(3),
	533:  uint16(1),
	534:  uint16(sym_comment),
	535:  uint16(9),
	536:  uint16(1),
	537:  uint16(sym_identifier),
	538:  uint16(43),
	539:  uint16(1),
	540:  uint16(anon_sym_mux),
	541:  uint16(45),
	542:  uint16(1),
	543:  uint16(anon_sym_validif),
	544:  uint16(143),
	545:  uint16(1),
	546:  uint16(sym_primop),
	547:  uint16(265),
	548:  uint16(1),
	549:  uint16(sym_expression),
	550:  uint16(433),
	551:  uint16(1),
	552:  uint16(sym_litType),
	553:  uint16(13),
	554:  uint16(2),
	555:  uint16(anon_sym_UInt),
	556:  uint16(anon_sym_SInt),
	557:  uint16(160),
	558:  uint16(7),
	559:  uint16(sym_literal),
	560:  uint16(sym_sub_field),
	561:  uint16(sym_sub_index),
	562:  uint16(sym_sub_access),
	563:  uint16(sym_mux),
	564:  uint16(sym_conditionally_valid),
	565:  uint16(sym_primitive_operation),
	566:  uint16(47),
	567:  uint16(39),
	568:  uint16(anon_sym_add),
	569:  uint16(anon_sym_sub),
	570:  uint16(anon_sym_mul),
	571:  uint16(anon_sym_div),
	572:  uint16(anon_sym_rem),
	573:  uint16(anon_sym_lt),
	574:  uint16(anon_sym_leq),
	575:  uint16(anon_sym_gt),
	576:  uint16(anon_sym_geq),
	577:  uint16(anon_sym_eq),
	578:  uint16(anon_sym_neq),
	579:  uint16(anon_sym_pad),
	580:  uint16(anon_sym_asUInt),
	581:  uint16(anon_sym_asAsyncReset),
	582:  uint16(anon_sym_asSInt),
	583:  uint16(anon_sym_asClock),
	584:  uint16(anon_sym_shl),
	585:  uint16(anon_sym_shr),
	586:  uint16(anon_sym_dshl),
	587:  uint16(anon_sym_dshlw),
	588:  uint16(anon_sym_dshr),
	589:  uint16(anon_sym_dshrw),
	590:  uint16(anon_sym_cvt),
	591:  uint16(anon_sym_neg),
	592:  uint16(anon_sym_not),
	593:  uint16(anon_sym_and),
	594:  uint16(anon_sym_or),
	595:  uint16(anon_sym_xor),
	596:  uint16(anon_sym_andr),
	597:  uint16(anon_sym_orr),
	598:  uint16(anon_sym_xorr),
	599:  uint16(anon_sym_cat),
	600:  uint16(anon_sym_bits),
	601:  uint16(anon_sym_head),
	602:  uint16(anon_sym_tail),
	603:  uint16(anon_sym_asFixedPoint),
	604:  uint16(anon_sym_bpshl),
	605:  uint16(anon_sym_bpshr),
	606:  uint16(anon_sym_bpset),
	607:  uint16(10),
	608:  uint16(3),
	609:  uint16(1),
	610:  uint16(sym_comment),
	611:  uint16(9),
	612:  uint16(1),
	613:  uint16(sym_identifier),
	614:  uint16(43),
	615:  uint16(1),
	616:  uint16(anon_sym_mux),
	617:  uint16(45),
	618:  uint16(1),
	619:  uint16(anon_sym_validif),
	620:  uint16(143),
	621:  uint16(1),
	622:  uint16(sym_primop),
	623:  uint16(304),
	624:  uint16(1),
	625:  uint16(sym_expression),
	626:  uint16(433),
	627:  uint16(1),
	628:  uint16(sym_litType),
	629:  uint16(13),
	630:  uint16(2),
	631:  uint16(anon_sym_UInt),
	632:  uint16(anon_sym_SInt),
	633:  uint16(160),
	634:  uint16(7),
	635:  uint16(sym_literal),
	636:  uint16(sym_sub_field),
	637:  uint16(sym_sub_index),
	638:  uint16(sym_sub_access),
	639:  uint16(sym_mux),
	640:  uint16(sym_conditionally_valid),
	641:  uint16(sym_primitive_operation),
	642:  uint16(47),
	643:  uint16(39),
	644:  uint16(anon_sym_add),
	645:  uint16(anon_sym_sub),
	646:  uint16(anon_sym_mul),
	647:  uint16(anon_sym_div),
	648:  uint16(anon_sym_rem),
	649:  uint16(anon_sym_lt),
	650:  uint16(anon_sym_leq),
	651:  uint16(anon_sym_gt),
	652:  uint16(anon_sym_geq),
	653:  uint16(anon_sym_eq),
	654:  uint16(anon_sym_neq),
	655:  uint16(anon_sym_pad),
	656:  uint16(anon_sym_asUInt),
	657:  uint16(anon_sym_asAsyncReset),
	658:  uint16(anon_sym_asSInt),
	659:  uint16(anon_sym_asClock),
	660:  uint16(anon_sym_shl),
	661:  uint16(anon_sym_shr),
	662:  uint16(anon_sym_dshl),
	663:  uint16(anon_sym_dshlw),
	664:  uint16(anon_sym_dshr),
	665:  uint16(anon_sym_dshrw),
	666:  uint16(anon_sym_cvt),
	667:  uint16(anon_sym_neg),
	668:  uint16(anon_sym_not),
	669:  uint16(anon_sym_and),
	670:  uint16(anon_sym_or),
	671:  uint16(anon_sym_xor),
	672:  uint16(anon_sym_andr),
	673:  uint16(anon_sym_orr),
	674:  uint16(anon_sym_xorr),
	675:  uint16(anon_sym_cat),
	676:  uint16(anon_sym_bits),
	677:  uint16(anon_sym_head),
	678:  uint16(anon_sym_tail),
	679:  uint16(anon_sym_asFixedPoint),
	680:  uint16(anon_sym_bpshl),
	681:  uint16(anon_sym_bpshr),
	682:  uint16(anon_sym_bpset),
	683:  uint16(10),
	684:  uint16(3),
	685:  uint16(1),
	686:  uint16(sym_comment),
	687:  uint16(9),
	688:  uint16(1),
	689:  uint16(sym_identifier),
	690:  uint16(43),
	691:  uint16(1),
	692:  uint16(anon_sym_mux),
	693:  uint16(45),
	694:  uint16(1),
	695:  uint16(anon_sym_validif),
	696:  uint16(143),
	697:  uint16(1),
	698:  uint16(sym_primop),
	699:  uint16(248),
	700:  uint16(1),
	701:  uint16(sym_expression),
	702:  uint16(433),
	703:  uint16(1),
	704:  uint16(sym_litType),
	705:  uint16(13),
	706:  uint16(2),
	707:  uint16(anon_sym_UInt),
	708:  uint16(anon_sym_SInt),
	709:  uint16(160),
	710:  uint16(7),
	711:  uint16(sym_literal),
	712:  uint16(sym_sub_field),
	713:  uint16(sym_sub_index),
	714:  uint16(sym_sub_access),
	715:  uint16(sym_mux),
	716:  uint16(sym_conditionally_valid),
	717:  uint16(sym_primitive_operation),
	718:  uint16(47),
	719:  uint16(39),
	720:  uint16(anon_sym_add),
	721:  uint16(anon_sym_sub),
	722:  uint16(anon_sym_mul),
	723:  uint16(anon_sym_div),
	724:  uint16(anon_sym_rem),
	725:  uint16(anon_sym_lt),
	726:  uint16(anon_sym_leq),
	727:  uint16(anon_sym_gt),
	728:  uint16(anon_sym_geq),
	729:  uint16(anon_sym_eq),
	730:  uint16(anon_sym_neq),
	731:  uint16(anon_sym_pad),
	732:  uint16(anon_sym_asUInt),
	733:  uint16(anon_sym_asAsyncReset),
	734:  uint16(anon_sym_asSInt),
	735:  uint16(anon_sym_asClock),
	736:  uint16(anon_sym_shl),
	737:  uint16(anon_sym_shr),
	738:  uint16(anon_sym_dshl),
	739:  uint16(anon_sym_dshlw),
	740:  uint16(anon_sym_dshr),
	741:  uint16(anon_sym_dshrw),
	742:  uint16(anon_sym_cvt),
	743:  uint16(anon_sym_neg),
	744:  uint16(anon_sym_not),
	745:  uint16(anon_sym_and),
	746:  uint16(anon_sym_or),
	747:  uint16(anon_sym_xor),
	748:  uint16(anon_sym_andr),
	749:  uint16(anon_sym_orr),
	750:  uint16(anon_sym_xorr),
	751:  uint16(anon_sym_cat),
	752:  uint16(anon_sym_bits),
	753:  uint16(anon_sym_head),
	754:  uint16(anon_sym_tail),
	755:  uint16(anon_sym_asFixedPoint),
	756:  uint16(anon_sym_bpshl),
	757:  uint16(anon_sym_bpshr),
	758:  uint16(anon_sym_bpset),
	759:  uint16(10),
	760:  uint16(3),
	761:  uint16(1),
	762:  uint16(sym_comment),
	763:  uint16(9),
	764:  uint16(1),
	765:  uint16(sym_identifier),
	766:  uint16(43),
	767:  uint16(1),
	768:  uint16(anon_sym_mux),
	769:  uint16(45),
	770:  uint16(1),
	771:  uint16(anon_sym_validif),
	772:  uint16(143),
	773:  uint16(1),
	774:  uint16(sym_primop),
	775:  uint16(240),
	776:  uint16(1),
	777:  uint16(sym_expression),
	778:  uint16(433),
	779:  uint16(1),
	780:  uint16(sym_litType),
	781:  uint16(13),
	782:  uint16(2),
	783:  uint16(anon_sym_UInt),
	784:  uint16(anon_sym_SInt),
	785:  uint16(160),
	786:  uint16(7),
	787:  uint16(sym_literal),
	788:  uint16(sym_sub_field),
	789:  uint16(sym_sub_index),
	790:  uint16(sym_sub_access),
	791:  uint16(sym_mux),
	792:  uint16(sym_conditionally_valid),
	793:  uint16(sym_primitive_operation),
	794:  uint16(47),
	795:  uint16(39),
	796:  uint16(anon_sym_add),
	797:  uint16(anon_sym_sub),
	798:  uint16(anon_sym_mul),
	799:  uint16(anon_sym_div),
	800:  uint16(anon_sym_rem),
	801:  uint16(anon_sym_lt),
	802:  uint16(anon_sym_leq),
	803:  uint16(anon_sym_gt),
	804:  uint16(anon_sym_geq),
	805:  uint16(anon_sym_eq),
	806:  uint16(anon_sym_neq),
	807:  uint16(anon_sym_pad),
	808:  uint16(anon_sym_asUInt),
	809:  uint16(anon_sym_asAsyncReset),
	810:  uint16(anon_sym_asSInt),
	811:  uint16(anon_sym_asClock),
	812:  uint16(anon_sym_shl),
	813:  uint16(anon_sym_shr),
	814:  uint16(anon_sym_dshl),
	815:  uint16(anon_sym_dshlw),
	816:  uint16(anon_sym_dshr),
	817:  uint16(anon_sym_dshrw),
	818:  uint16(anon_sym_cvt),
	819:  uint16(anon_sym_neg),
	820:  uint16(anon_sym_not),
	821:  uint16(anon_sym_and),
	822:  uint16(anon_sym_or),
	823:  uint16(anon_sym_xor),
	824:  uint16(anon_sym_andr),
	825:  uint16(anon_sym_orr),
	826:  uint16(anon_sym_xorr),
	827:  uint16(anon_sym_cat),
	828:  uint16(anon_sym_bits),
	829:  uint16(anon_sym_head),
	830:  uint16(anon_sym_tail),
	831:  uint16(anon_sym_asFixedPoint),
	832:  uint16(anon_sym_bpshl),
	833:  uint16(anon_sym_bpshr),
	834:  uint16(anon_sym_bpset),
	835:  uint16(10),
	836:  uint16(3),
	837:  uint16(1),
	838:  uint16(sym_comment),
	839:  uint16(9),
	840:  uint16(1),
	841:  uint16(sym_identifier),
	842:  uint16(43),
	843:  uint16(1),
	844:  uint16(anon_sym_mux),
	845:  uint16(45),
	846:  uint16(1),
	847:  uint16(anon_sym_validif),
	848:  uint16(143),
	849:  uint16(1),
	850:  uint16(sym_primop),
	851:  uint16(289),
	852:  uint16(1),
	853:  uint16(sym_expression),
	854:  uint16(433),
	855:  uint16(1),
	856:  uint16(sym_litType),
	857:  uint16(13),
	858:  uint16(2),
	859:  uint16(anon_sym_UInt),
	860:  uint16(anon_sym_SInt),
	861:  uint16(160),
	862:  uint16(7),
	863:  uint16(sym_literal),
	864:  uint16(sym_sub_field),
	865:  uint16(sym_sub_index),
	866:  uint16(sym_sub_access),
	867:  uint16(sym_mux),
	868:  uint16(sym_conditionally_valid),
	869:  uint16(sym_primitive_operation),
	870:  uint16(47),
	871:  uint16(39),
	872:  uint16(anon_sym_add),
	873:  uint16(anon_sym_sub),
	874:  uint16(anon_sym_mul),
	875:  uint16(anon_sym_div),
	876:  uint16(anon_sym_rem),
	877:  uint16(anon_sym_lt),
	878:  uint16(anon_sym_leq),
	879:  uint16(anon_sym_gt),
	880:  uint16(anon_sym_geq),
	881:  uint16(anon_sym_eq),
	882:  uint16(anon_sym_neq),
	883:  uint16(anon_sym_pad),
	884:  uint16(anon_sym_asUInt),
	885:  uint16(anon_sym_asAsyncReset),
	886:  uint16(anon_sym_asSInt),
	887:  uint16(anon_sym_asClock),
	888:  uint16(anon_sym_shl),
	889:  uint16(anon_sym_shr),
	890:  uint16(anon_sym_dshl),
	891:  uint16(anon_sym_dshlw),
	892:  uint16(anon_sym_dshr),
	893:  uint16(anon_sym_dshrw),
	894:  uint16(anon_sym_cvt),
	895:  uint16(anon_sym_neg),
	896:  uint16(anon_sym_not),
	897:  uint16(anon_sym_and),
	898:  uint16(anon_sym_or),
	899:  uint16(anon_sym_xor),
	900:  uint16(anon_sym_andr),
	901:  uint16(anon_sym_orr),
	902:  uint16(anon_sym_xorr),
	903:  uint16(anon_sym_cat),
	904:  uint16(anon_sym_bits),
	905:  uint16(anon_sym_head),
	906:  uint16(anon_sym_tail),
	907:  uint16(anon_sym_asFixedPoint),
	908:  uint16(anon_sym_bpshl),
	909:  uint16(anon_sym_bpshr),
	910:  uint16(anon_sym_bpset),
	911:  uint16(10),
	912:  uint16(3),
	913:  uint16(1),
	914:  uint16(sym_comment),
	915:  uint16(9),
	916:  uint16(1),
	917:  uint16(sym_identifier),
	918:  uint16(43),
	919:  uint16(1),
	920:  uint16(anon_sym_mux),
	921:  uint16(45),
	922:  uint16(1),
	923:  uint16(anon_sym_validif),
	924:  uint16(143),
	925:  uint16(1),
	926:  uint16(sym_primop),
	927:  uint16(246),
	928:  uint16(1),
	929:  uint16(sym_expression),
	930:  uint16(433),
	931:  uint16(1),
	932:  uint16(sym_litType),
	933:  uint16(13),
	934:  uint16(2),
	935:  uint16(anon_sym_UInt),
	936:  uint16(anon_sym_SInt),
	937:  uint16(160),
	938:  uint16(7),
	939:  uint16(sym_literal),
	940:  uint16(sym_sub_field),
	941:  uint16(sym_sub_index),
	942:  uint16(sym_sub_access),
	943:  uint16(sym_mux),
	944:  uint16(sym_conditionally_valid),
	945:  uint16(sym_primitive_operation),
	946:  uint16(47),
	947:  uint16(39),
	948:  uint16(anon_sym_add),
	949:  uint16(anon_sym_sub),
	950:  uint16(anon_sym_mul),
	951:  uint16(anon_sym_div),
	952:  uint16(anon_sym_rem),
	953:  uint16(anon_sym_lt),
	954:  uint16(anon_sym_leq),
	955:  uint16(anon_sym_gt),
	956:  uint16(anon_sym_geq),
	957:  uint16(anon_sym_eq),
	958:  uint16(anon_sym_neq),
	959:  uint16(anon_sym_pad),
	960:  uint16(anon_sym_asUInt),
	961:  uint16(anon_sym_asAsyncReset),
	962:  uint16(anon_sym_asSInt),
	963:  uint16(anon_sym_asClock),
	964:  uint16(anon_sym_shl),
	965:  uint16(anon_sym_shr),
	966:  uint16(anon_sym_dshl),
	967:  uint16(anon_sym_dshlw),
	968:  uint16(anon_sym_dshr),
	969:  uint16(anon_sym_dshrw),
	970:  uint16(anon_sym_cvt),
	971:  uint16(anon_sym_neg),
	972:  uint16(anon_sym_not),
	973:  uint16(anon_sym_and),
	974:  uint16(anon_sym_or),
	975:  uint16(anon_sym_xor),
	976:  uint16(anon_sym_andr),
	977:  uint16(anon_sym_orr),
	978:  uint16(anon_sym_xorr),
	979:  uint16(anon_sym_cat),
	980:  uint16(anon_sym_bits),
	981:  uint16(anon_sym_head),
	982:  uint16(anon_sym_tail),
	983:  uint16(anon_sym_asFixedPoint),
	984:  uint16(anon_sym_bpshl),
	985:  uint16(anon_sym_bpshr),
	986:  uint16(anon_sym_bpset),
	987:  uint16(10),
	988:  uint16(3),
	989:  uint16(1),
	990:  uint16(sym_comment),
	991:  uint16(9),
	992:  uint16(1),
	993:  uint16(sym_identifier),
	994:  uint16(43),
	995:  uint16(1),
	996:  uint16(anon_sym_mux),
	997:  uint16(45),
	998:  uint16(1),
	999:  uint16(anon_sym_validif),
	1000: uint16(143),
	1001: uint16(1),
	1002: uint16(sym_primop),
	1003: uint16(267),
	1004: uint16(1),
	1005: uint16(sym_expression),
	1006: uint16(433),
	1007: uint16(1),
	1008: uint16(sym_litType),
	1009: uint16(13),
	1010: uint16(2),
	1011: uint16(anon_sym_UInt),
	1012: uint16(anon_sym_SInt),
	1013: uint16(160),
	1014: uint16(7),
	1015: uint16(sym_literal),
	1016: uint16(sym_sub_field),
	1017: uint16(sym_sub_index),
	1018: uint16(sym_sub_access),
	1019: uint16(sym_mux),
	1020: uint16(sym_conditionally_valid),
	1021: uint16(sym_primitive_operation),
	1022: uint16(47),
	1023: uint16(39),
	1024: uint16(anon_sym_add),
	1025: uint16(anon_sym_sub),
	1026: uint16(anon_sym_mul),
	1027: uint16(anon_sym_div),
	1028: uint16(anon_sym_rem),
	1029: uint16(anon_sym_lt),
	1030: uint16(anon_sym_leq),
	1031: uint16(anon_sym_gt),
	1032: uint16(anon_sym_geq),
	1033: uint16(anon_sym_eq),
	1034: uint16(anon_sym_neq),
	1035: uint16(anon_sym_pad),
	1036: uint16(anon_sym_asUInt),
	1037: uint16(anon_sym_asAsyncReset),
	1038: uint16(anon_sym_asSInt),
	1039: uint16(anon_sym_asClock),
	1040: uint16(anon_sym_shl),
	1041: uint16(anon_sym_shr),
	1042: uint16(anon_sym_dshl),
	1043: uint16(anon_sym_dshlw),
	1044: uint16(anon_sym_dshr),
	1045: uint16(anon_sym_dshrw),
	1046: uint16(anon_sym_cvt),
	1047: uint16(anon_sym_neg),
	1048: uint16(anon_sym_not),
	1049: uint16(anon_sym_and),
	1050: uint16(anon_sym_or),
	1051: uint16(anon_sym_xor),
	1052: uint16(anon_sym_andr),
	1053: uint16(anon_sym_orr),
	1054: uint16(anon_sym_xorr),
	1055: uint16(anon_sym_cat),
	1056: uint16(anon_sym_bits),
	1057: uint16(anon_sym_head),
	1058: uint16(anon_sym_tail),
	1059: uint16(anon_sym_asFixedPoint),
	1060: uint16(anon_sym_bpshl),
	1061: uint16(anon_sym_bpshr),
	1062: uint16(anon_sym_bpset),
	1063: uint16(10),
	1064: uint16(3),
	1065: uint16(1),
	1066: uint16(sym_comment),
	1067: uint16(9),
	1068: uint16(1),
	1069: uint16(sym_identifier),
	1070: uint16(43),
	1071: uint16(1),
	1072: uint16(anon_sym_mux),
	1073: uint16(45),
	1074: uint16(1),
	1075: uint16(anon_sym_validif),
	1076: uint16(143),
	1077: uint16(1),
	1078: uint16(sym_primop),
	1079: uint16(268),
	1080: uint16(1),
	1081: uint16(sym_expression),
	1082: uint16(433),
	1083: uint16(1),
	1084: uint16(sym_litType),
	1085: uint16(13),
	1086: uint16(2),
	1087: uint16(anon_sym_UInt),
	1088: uint16(anon_sym_SInt),
	1089: uint16(160),
	1090: uint16(7),
	1091: uint16(sym_literal),
	1092: uint16(sym_sub_field),
	1093: uint16(sym_sub_index),
	1094: uint16(sym_sub_access),
	1095: uint16(sym_mux),
	1096: uint16(sym_conditionally_valid),
	1097: uint16(sym_primitive_operation),
	1098: uint16(47),
	1099: uint16(39),
	1100: uint16(anon_sym_add),
	1101: uint16(anon_sym_sub),
	1102: uint16(anon_sym_mul),
	1103: uint16(anon_sym_div),
	1104: uint16(anon_sym_rem),
	1105: uint16(anon_sym_lt),
	1106: uint16(anon_sym_leq),
	1107: uint16(anon_sym_gt),
	1108: uint16(anon_sym_geq),
	1109: uint16(anon_sym_eq),
	1110: uint16(anon_sym_neq),
	1111: uint16(anon_sym_pad),
	1112: uint16(anon_sym_asUInt),
	1113: uint16(anon_sym_asAsyncReset),
	1114: uint16(anon_sym_asSInt),
	1115: uint16(anon_sym_asClock),
	1116: uint16(anon_sym_shl),
	1117: uint16(anon_sym_shr),
	1118: uint16(anon_sym_dshl),
	1119: uint16(anon_sym_dshlw),
	1120: uint16(anon_sym_dshr),
	1121: uint16(anon_sym_dshrw),
	1122: uint16(anon_sym_cvt),
	1123: uint16(anon_sym_neg),
	1124: uint16(anon_sym_not),
	1125: uint16(anon_sym_and),
	1126: uint16(anon_sym_or),
	1127: uint16(anon_sym_xor),
	1128: uint16(anon_sym_andr),
	1129: uint16(anon_sym_orr),
	1130: uint16(anon_sym_xorr),
	1131: uint16(anon_sym_cat),
	1132: uint16(anon_sym_bits),
	1133: uint16(anon_sym_head),
	1134: uint16(anon_sym_tail),
	1135: uint16(anon_sym_asFixedPoint),
	1136: uint16(anon_sym_bpshl),
	1137: uint16(anon_sym_bpshr),
	1138: uint16(anon_sym_bpset),
	1139: uint16(10),
	1140: uint16(3),
	1141: uint16(1),
	1142: uint16(sym_comment),
	1143: uint16(9),
	1144: uint16(1),
	1145: uint16(sym_identifier),
	1146: uint16(43),
	1147: uint16(1),
	1148: uint16(anon_sym_mux),
	1149: uint16(45),
	1150: uint16(1),
	1151: uint16(anon_sym_validif),
	1152: uint16(143),
	1153: uint16(1),
	1154: uint16(sym_primop),
	1155: uint16(217),
	1156: uint16(1),
	1157: uint16(sym_expression),
	1158: uint16(433),
	1159: uint16(1),
	1160: uint16(sym_litType),
	1161: uint16(13),
	1162: uint16(2),
	1163: uint16(anon_sym_UInt),
	1164: uint16(anon_sym_SInt),
	1165: uint16(160),
	1166: uint16(7),
	1167: uint16(sym_literal),
	1168: uint16(sym_sub_field),
	1169: uint16(sym_sub_index),
	1170: uint16(sym_sub_access),
	1171: uint16(sym_mux),
	1172: uint16(sym_conditionally_valid),
	1173: uint16(sym_primitive_operation),
	1174: uint16(47),
	1175: uint16(39),
	1176: uint16(anon_sym_add),
	1177: uint16(anon_sym_sub),
	1178: uint16(anon_sym_mul),
	1179: uint16(anon_sym_div),
	1180: uint16(anon_sym_rem),
	1181: uint16(anon_sym_lt),
	1182: uint16(anon_sym_leq),
	1183: uint16(anon_sym_gt),
	1184: uint16(anon_sym_geq),
	1185: uint16(anon_sym_eq),
	1186: uint16(anon_sym_neq),
	1187: uint16(anon_sym_pad),
	1188: uint16(anon_sym_asUInt),
	1189: uint16(anon_sym_asAsyncReset),
	1190: uint16(anon_sym_asSInt),
	1191: uint16(anon_sym_asClock),
	1192: uint16(anon_sym_shl),
	1193: uint16(anon_sym_shr),
	1194: uint16(anon_sym_dshl),
	1195: uint16(anon_sym_dshlw),
	1196: uint16(anon_sym_dshr),
	1197: uint16(anon_sym_dshrw),
	1198: uint16(anon_sym_cvt),
	1199: uint16(anon_sym_neg),
	1200: uint16(anon_sym_not),
	1201: uint16(anon_sym_and),
	1202: uint16(anon_sym_or),
	1203: uint16(anon_sym_xor),
	1204: uint16(anon_sym_andr),
	1205: uint16(anon_sym_orr),
	1206: uint16(anon_sym_xorr),
	1207: uint16(anon_sym_cat),
	1208: uint16(anon_sym_bits),
	1209: uint16(anon_sym_head),
	1210: uint16(anon_sym_tail),
	1211: uint16(anon_sym_asFixedPoint),
	1212: uint16(anon_sym_bpshl),
	1213: uint16(anon_sym_bpshr),
	1214: uint16(anon_sym_bpset),
	1215: uint16(10),
	1216: uint16(3),
	1217: uint16(1),
	1218: uint16(sym_comment),
	1219: uint16(9),
	1220: uint16(1),
	1221: uint16(sym_identifier),
	1222: uint16(43),
	1223: uint16(1),
	1224: uint16(anon_sym_mux),
	1225: uint16(45),
	1226: uint16(1),
	1227: uint16(anon_sym_validif),
	1228: uint16(143),
	1229: uint16(1),
	1230: uint16(sym_primop),
	1231: uint16(281),
	1232: uint16(1),
	1233: uint16(sym_expression),
	1234: uint16(433),
	1235: uint16(1),
	1236: uint16(sym_litType),
	1237: uint16(13),
	1238: uint16(2),
	1239: uint16(anon_sym_UInt),
	1240: uint16(anon_sym_SInt),
	1241: uint16(160),
	1242: uint16(7),
	1243: uint16(sym_literal),
	1244: uint16(sym_sub_field),
	1245: uint16(sym_sub_index),
	1246: uint16(sym_sub_access),
	1247: uint16(sym_mux),
	1248: uint16(sym_conditionally_valid),
	1249: uint16(sym_primitive_operation),
	1250: uint16(47),
	1251: uint16(39),
	1252: uint16(anon_sym_add),
	1253: uint16(anon_sym_sub),
	1254: uint16(anon_sym_mul),
	1255: uint16(anon_sym_div),
	1256: uint16(anon_sym_rem),
	1257: uint16(anon_sym_lt),
	1258: uint16(anon_sym_leq),
	1259: uint16(anon_sym_gt),
	1260: uint16(anon_sym_geq),
	1261: uint16(anon_sym_eq),
	1262: uint16(anon_sym_neq),
	1263: uint16(anon_sym_pad),
	1264: uint16(anon_sym_asUInt),
	1265: uint16(anon_sym_asAsyncReset),
	1266: uint16(anon_sym_asSInt),
	1267: uint16(anon_sym_asClock),
	1268: uint16(anon_sym_shl),
	1269: uint16(anon_sym_shr),
	1270: uint16(anon_sym_dshl),
	1271: uint16(anon_sym_dshlw),
	1272: uint16(anon_sym_dshr),
	1273: uint16(anon_sym_dshrw),
	1274: uint16(anon_sym_cvt),
	1275: uint16(anon_sym_neg),
	1276: uint16(anon_sym_not),
	1277: uint16(anon_sym_and),
	1278: uint16(anon_sym_or),
	1279: uint16(anon_sym_xor),
	1280: uint16(anon_sym_andr),
	1281: uint16(anon_sym_orr),
	1282: uint16(anon_sym_xorr),
	1283: uint16(anon_sym_cat),
	1284: uint16(anon_sym_bits),
	1285: uint16(anon_sym_head),
	1286: uint16(anon_sym_tail),
	1287: uint16(anon_sym_asFixedPoint),
	1288: uint16(anon_sym_bpshl),
	1289: uint16(anon_sym_bpshr),
	1290: uint16(anon_sym_bpset),
	1291: uint16(10),
	1292: uint16(3),
	1293: uint16(1),
	1294: uint16(sym_comment),
	1295: uint16(9),
	1296: uint16(1),
	1297: uint16(sym_identifier),
	1298: uint16(43),
	1299: uint16(1),
	1300: uint16(anon_sym_mux),
	1301: uint16(45),
	1302: uint16(1),
	1303: uint16(anon_sym_validif),
	1304: uint16(143),
	1305: uint16(1),
	1306: uint16(sym_primop),
	1307: uint16(250),
	1308: uint16(1),
	1309: uint16(sym_expression),
	1310: uint16(433),
	1311: uint16(1),
	1312: uint16(sym_litType),
	1313: uint16(13),
	1314: uint16(2),
	1315: uint16(anon_sym_UInt),
	1316: uint16(anon_sym_SInt),
	1317: uint16(160),
	1318: uint16(7),
	1319: uint16(sym_literal),
	1320: uint16(sym_sub_field),
	1321: uint16(sym_sub_index),
	1322: uint16(sym_sub_access),
	1323: uint16(sym_mux),
	1324: uint16(sym_conditionally_valid),
	1325: uint16(sym_primitive_operation),
	1326: uint16(47),
	1327: uint16(39),
	1328: uint16(anon_sym_add),
	1329: uint16(anon_sym_sub),
	1330: uint16(anon_sym_mul),
	1331: uint16(anon_sym_div),
	1332: uint16(anon_sym_rem),
	1333: uint16(anon_sym_lt),
	1334: uint16(anon_sym_leq),
	1335: uint16(anon_sym_gt),
	1336: uint16(anon_sym_geq),
	1337: uint16(anon_sym_eq),
	1338: uint16(anon_sym_neq),
	1339: uint16(anon_sym_pad),
	1340: uint16(anon_sym_asUInt),
	1341: uint16(anon_sym_asAsyncReset),
	1342: uint16(anon_sym_asSInt),
	1343: uint16(anon_sym_asClock),
	1344: uint16(anon_sym_shl),
	1345: uint16(anon_sym_shr),
	1346: uint16(anon_sym_dshl),
	1347: uint16(anon_sym_dshlw),
	1348: uint16(anon_sym_dshr),
	1349: uint16(anon_sym_dshrw),
	1350: uint16(anon_sym_cvt),
	1351: uint16(anon_sym_neg),
	1352: uint16(anon_sym_not),
	1353: uint16(anon_sym_and),
	1354: uint16(anon_sym_or),
	1355: uint16(anon_sym_xor),
	1356: uint16(anon_sym_andr),
	1357: uint16(anon_sym_orr),
	1358: uint16(anon_sym_xorr),
	1359: uint16(anon_sym_cat),
	1360: uint16(anon_sym_bits),
	1361: uint16(anon_sym_head),
	1362: uint16(anon_sym_tail),
	1363: uint16(anon_sym_asFixedPoint),
	1364: uint16(anon_sym_bpshl),
	1365: uint16(anon_sym_bpshr),
	1366: uint16(anon_sym_bpset),
	1367: uint16(10),
	1368: uint16(3),
	1369: uint16(1),
	1370: uint16(sym_comment),
	1371: uint16(9),
	1372: uint16(1),
	1373: uint16(sym_identifier),
	1374: uint16(43),
	1375: uint16(1),
	1376: uint16(anon_sym_mux),
	1377: uint16(45),
	1378: uint16(1),
	1379: uint16(anon_sym_validif),
	1380: uint16(143),
	1381: uint16(1),
	1382: uint16(sym_primop),
	1383: uint16(251),
	1384: uint16(1),
	1385: uint16(sym_expression),
	1386: uint16(433),
	1387: uint16(1),
	1388: uint16(sym_litType),
	1389: uint16(13),
	1390: uint16(2),
	1391: uint16(anon_sym_UInt),
	1392: uint16(anon_sym_SInt),
	1393: uint16(160),
	1394: uint16(7),
	1395: uint16(sym_literal),
	1396: uint16(sym_sub_field),
	1397: uint16(sym_sub_index),
	1398: uint16(sym_sub_access),
	1399: uint16(sym_mux),
	1400: uint16(sym_conditionally_valid),
	1401: uint16(sym_primitive_operation),
	1402: uint16(47),
	1403: uint16(39),
	1404: uint16(anon_sym_add),
	1405: uint16(anon_sym_sub),
	1406: uint16(anon_sym_mul),
	1407: uint16(anon_sym_div),
	1408: uint16(anon_sym_rem),
	1409: uint16(anon_sym_lt),
	1410: uint16(anon_sym_leq),
	1411: uint16(anon_sym_gt),
	1412: uint16(anon_sym_geq),
	1413: uint16(anon_sym_eq),
	1414: uint16(anon_sym_neq),
	1415: uint16(anon_sym_pad),
	1416: uint16(anon_sym_asUInt),
	1417: uint16(anon_sym_asAsyncReset),
	1418: uint16(anon_sym_asSInt),
	1419: uint16(anon_sym_asClock),
	1420: uint16(anon_sym_shl),
	1421: uint16(anon_sym_shr),
	1422: uint16(anon_sym_dshl),
	1423: uint16(anon_sym_dshlw),
	1424: uint16(anon_sym_dshr),
	1425: uint16(anon_sym_dshrw),
	1426: uint16(anon_sym_cvt),
	1427: uint16(anon_sym_neg),
	1428: uint16(anon_sym_not),
	1429: uint16(anon_sym_and),
	1430: uint16(anon_sym_or),
	1431: uint16(anon_sym_xor),
	1432: uint16(anon_sym_andr),
	1433: uint16(anon_sym_orr),
	1434: uint16(anon_sym_xorr),
	1435: uint16(anon_sym_cat),
	1436: uint16(anon_sym_bits),
	1437: uint16(anon_sym_head),
	1438: uint16(anon_sym_tail),
	1439: uint16(anon_sym_asFixedPoint),
	1440: uint16(anon_sym_bpshl),
	1441: uint16(anon_sym_bpshr),
	1442: uint16(anon_sym_bpset),
	1443: uint16(10),
	1444: uint16(3),
	1445: uint16(1),
	1446: uint16(sym_comment),
	1447: uint16(471),
	1448: uint16(1),
	1449: uint16(sym_identifier),
	1450: uint16(473),
	1451: uint16(1),
	1452: uint16(anon_sym_mux),
	1453: uint16(475),
	1454: uint16(1),
	1455: uint16(anon_sym_validif),
	1456: uint16(19),
	1457: uint16(1),
	1458: uint16(sym_primop),
	1459: uint16(35),
	1460: uint16(1),
	1461: uint16(sym_expression),
	1462: uint16(425),
	1463: uint16(1),
	1464: uint16(sym_litType),
	1465: uint16(13),
	1466: uint16(2),
	1467: uint16(anon_sym_UInt),
	1468: uint16(anon_sym_SInt),
	1469: uint16(31),
	1470: uint16(7),
	1471: uint16(sym_literal),
	1472: uint16(sym_sub_field),
	1473: uint16(sym_sub_index),
	1474: uint16(sym_sub_access),
	1475: uint16(sym_mux),
	1476: uint16(sym_conditionally_valid),
	1477: uint16(sym_primitive_operation),
	1478: uint16(477),
	1479: uint16(39),
	1480: uint16(anon_sym_add),
	1481: uint16(anon_sym_sub),
	1482: uint16(anon_sym_mul),
	1483: uint16(anon_sym_div),
	1484: uint16(anon_sym_rem),
	1485: uint16(anon_sym_lt),
	1486: uint16(anon_sym_leq),
	1487: uint16(anon_sym_gt),
	1488: uint16(anon_sym_geq),
	1489: uint16(anon_sym_eq),
	1490: uint16(anon_sym_neq),
	1491: uint16(anon_sym_pad),
	1492: uint16(anon_sym_asUInt),
	1493: uint16(anon_sym_asAsyncReset),
	1494: uint16(anon_sym_asSInt),
	1495: uint16(anon_sym_asClock),
	1496: uint16(anon_sym_shl),
	1497: uint16(anon_sym_shr),
	1498: uint16(anon_sym_dshl),
	1499: uint16(anon_sym_dshlw),
	1500: uint16(anon_sym_dshr),
	1501: uint16(anon_sym_dshrw),
	1502: uint16(anon_sym_cvt),
	1503: uint16(anon_sym_neg),
	1504: uint16(anon_sym_not),
	1505: uint16(anon_sym_and),
	1506: uint16(anon_sym_or),
	1507: uint16(anon_sym_xor),
	1508: uint16(anon_sym_andr),
	1509: uint16(anon_sym_orr),
	1510: uint16(anon_sym_xorr),
	1511: uint16(anon_sym_cat),
	1512: uint16(anon_sym_bits),
	1513: uint16(anon_sym_head),
	1514: uint16(anon_sym_tail),
	1515: uint16(anon_sym_asFixedPoint),
	1516: uint16(anon_sym_bpshl),
	1517: uint16(anon_sym_bpshr),
	1518: uint16(anon_sym_bpset),
	1519: uint16(10),
	1520: uint16(3),
	1521: uint16(1),
	1522: uint16(sym_comment),
	1523: uint16(9),
	1524: uint16(1),
	1525: uint16(sym_identifier),
	1526: uint16(43),
	1527: uint16(1),
	1528: uint16(anon_sym_mux),
	1529: uint16(45),
	1530: uint16(1),
	1531: uint16(anon_sym_validif),
	1532: uint16(143),
	1533: uint16(1),
	1534: uint16(sym_primop),
	1535: uint16(282),
	1536: uint16(1),
	1537: uint16(sym_expression),
	1538: uint16(433),
	1539: uint16(1),
	1540: uint16(sym_litType),
	1541: uint16(13),
	1542: uint16(2),
	1543: uint16(anon_sym_UInt),
	1544: uint16(anon_sym_SInt),
	1545: uint16(160),
	1546: uint16(7),
	1547: uint16(sym_literal),
	1548: uint16(sym_sub_field),
	1549: uint16(sym_sub_index),
	1550: uint16(sym_sub_access),
	1551: uint16(sym_mux),
	1552: uint16(sym_conditionally_valid),
	1553: uint16(sym_primitive_operation),
	1554: uint16(47),
	1555: uint16(39),
	1556: uint16(anon_sym_add),
	1557: uint16(anon_sym_sub),
	1558: uint16(anon_sym_mul),
	1559: uint16(anon_sym_div),
	1560: uint16(anon_sym_rem),
	1561: uint16(anon_sym_lt),
	1562: uint16(anon_sym_leq),
	1563: uint16(anon_sym_gt),
	1564: uint16(anon_sym_geq),
	1565: uint16(anon_sym_eq),
	1566: uint16(anon_sym_neq),
	1567: uint16(anon_sym_pad),
	1568: uint16(anon_sym_asUInt),
	1569: uint16(anon_sym_asAsyncReset),
	1570: uint16(anon_sym_asSInt),
	1571: uint16(anon_sym_asClock),
	1572: uint16(anon_sym_shl),
	1573: uint16(anon_sym_shr),
	1574: uint16(anon_sym_dshl),
	1575: uint16(anon_sym_dshlw),
	1576: uint16(anon_sym_dshr),
	1577: uint16(anon_sym_dshrw),
	1578: uint16(anon_sym_cvt),
	1579: uint16(anon_sym_neg),
	1580: uint16(anon_sym_not),
	1581: uint16(anon_sym_and),
	1582: uint16(anon_sym_or),
	1583: uint16(anon_sym_xor),
	1584: uint16(anon_sym_andr),
	1585: uint16(anon_sym_orr),
	1586: uint16(anon_sym_xorr),
	1587: uint16(anon_sym_cat),
	1588: uint16(anon_sym_bits),
	1589: uint16(anon_sym_head),
	1590: uint16(anon_sym_tail),
	1591: uint16(anon_sym_asFixedPoint),
	1592: uint16(anon_sym_bpshl),
	1593: uint16(anon_sym_bpshr),
	1594: uint16(anon_sym_bpset),
	1595: uint16(10),
	1596: uint16(3),
	1597: uint16(1),
	1598: uint16(sym_comment),
	1599: uint16(471),
	1600: uint16(1),
	1601: uint16(sym_identifier),
	1602: uint16(473),
	1603: uint16(1),
	1604: uint16(anon_sym_mux),
	1605: uint16(475),
	1606: uint16(1),
	1607: uint16(anon_sym_validif),
	1608: uint16(19),
	1609: uint16(1),
	1610: uint16(sym_primop),
	1611: uint16(37),
	1612: uint16(1),
	1613: uint16(sym_expression),
	1614: uint16(425),
	1615: uint16(1),
	1616: uint16(sym_litType),
	1617: uint16(13),
	1618: uint16(2),
	1619: uint16(anon_sym_UInt),
	1620: uint16(anon_sym_SInt),
	1621: uint16(31),
	1622: uint16(7),
	1623: uint16(sym_literal),
	1624: uint16(sym_sub_field),
	1625: uint16(sym_sub_index),
	1626: uint16(sym_sub_access),
	1627: uint16(sym_mux),
	1628: uint16(sym_conditionally_valid),
	1629: uint16(sym_primitive_operation),
	1630: uint16(477),
	1631: uint16(39),
	1632: uint16(anon_sym_add),
	1633: uint16(anon_sym_sub),
	1634: uint16(anon_sym_mul),
	1635: uint16(anon_sym_div),
	1636: uint16(anon_sym_rem),
	1637: uint16(anon_sym_lt),
	1638: uint16(anon_sym_leq),
	1639: uint16(anon_sym_gt),
	1640: uint16(anon_sym_geq),
	1641: uint16(anon_sym_eq),
	1642: uint16(anon_sym_neq),
	1643: uint16(anon_sym_pad),
	1644: uint16(anon_sym_asUInt),
	1645: uint16(anon_sym_asAsyncReset),
	1646: uint16(anon_sym_asSInt),
	1647: uint16(anon_sym_asClock),
	1648: uint16(anon_sym_shl),
	1649: uint16(anon_sym_shr),
	1650: uint16(anon_sym_dshl),
	1651: uint16(anon_sym_dshlw),
	1652: uint16(anon_sym_dshr),
	1653: uint16(anon_sym_dshrw),
	1654: uint16(anon_sym_cvt),
	1655: uint16(anon_sym_neg),
	1656: uint16(anon_sym_not),
	1657: uint16(anon_sym_and),
	1658: uint16(anon_sym_or),
	1659: uint16(anon_sym_xor),
	1660: uint16(anon_sym_andr),
	1661: uint16(anon_sym_orr),
	1662: uint16(anon_sym_xorr),
	1663: uint16(anon_sym_cat),
	1664: uint16(anon_sym_bits),
	1665: uint16(anon_sym_head),
	1666: uint16(anon_sym_tail),
	1667: uint16(anon_sym_asFixedPoint),
	1668: uint16(anon_sym_bpshl),
	1669: uint16(anon_sym_bpshr),
	1670: uint16(anon_sym_bpset),
	1671: uint16(10),
	1672: uint16(3),
	1673: uint16(1),
	1674: uint16(sym_comment),
	1675: uint16(9),
	1676: uint16(1),
	1677: uint16(sym_identifier),
	1678: uint16(43),
	1679: uint16(1),
	1680: uint16(anon_sym_mux),
	1681: uint16(45),
	1682: uint16(1),
	1683: uint16(anon_sym_validif),
	1684: uint16(143),
	1685: uint16(1),
	1686: uint16(sym_primop),
	1687: uint16(284),
	1688: uint16(1),
	1689: uint16(sym_expression),
	1690: uint16(433),
	1691: uint16(1),
	1692: uint16(sym_litType),
	1693: uint16(13),
	1694: uint16(2),
	1695: uint16(anon_sym_UInt),
	1696: uint16(anon_sym_SInt),
	1697: uint16(160),
	1698: uint16(7),
	1699: uint16(sym_literal),
	1700: uint16(sym_sub_field),
	1701: uint16(sym_sub_index),
	1702: uint16(sym_sub_access),
	1703: uint16(sym_mux),
	1704: uint16(sym_conditionally_valid),
	1705: uint16(sym_primitive_operation),
	1706: uint16(47),
	1707: uint16(39),
	1708: uint16(anon_sym_add),
	1709: uint16(anon_sym_sub),
	1710: uint16(anon_sym_mul),
	1711: uint16(anon_sym_div),
	1712: uint16(anon_sym_rem),
	1713: uint16(anon_sym_lt),
	1714: uint16(anon_sym_leq),
	1715: uint16(anon_sym_gt),
	1716: uint16(anon_sym_geq),
	1717: uint16(anon_sym_eq),
	1718: uint16(anon_sym_neq),
	1719: uint16(anon_sym_pad),
	1720: uint16(anon_sym_asUInt),
	1721: uint16(anon_sym_asAsyncReset),
	1722: uint16(anon_sym_asSInt),
	1723: uint16(anon_sym_asClock),
	1724: uint16(anon_sym_shl),
	1725: uint16(anon_sym_shr),
	1726: uint16(anon_sym_dshl),
	1727: uint16(anon_sym_dshlw),
	1728: uint16(anon_sym_dshr),
	1729: uint16(anon_sym_dshrw),
	1730: uint16(anon_sym_cvt),
	1731: uint16(anon_sym_neg),
	1732: uint16(anon_sym_not),
	1733: uint16(anon_sym_and),
	1734: uint16(anon_sym_or),
	1735: uint16(anon_sym_xor),
	1736: uint16(anon_sym_andr),
	1737: uint16(anon_sym_orr),
	1738: uint16(anon_sym_xorr),
	1739: uint16(anon_sym_cat),
	1740: uint16(anon_sym_bits),
	1741: uint16(anon_sym_head),
	1742: uint16(anon_sym_tail),
	1743: uint16(anon_sym_asFixedPoint),
	1744: uint16(anon_sym_bpshl),
	1745: uint16(anon_sym_bpshr),
	1746: uint16(anon_sym_bpset),
	1747: uint16(10),
	1748: uint16(3),
	1749: uint16(1),
	1750: uint16(sym_comment),
	1751: uint16(471),
	1752: uint16(1),
	1753: uint16(sym_identifier),
	1754: uint16(473),
	1755: uint16(1),
	1756: uint16(anon_sym_mux),
	1757: uint16(475),
	1758: uint16(1),
	1759: uint16(anon_sym_validif),
	1760: uint16(19),
	1761: uint16(1),
	1762: uint16(sym_primop),
	1763: uint16(32),
	1764: uint16(1),
	1765: uint16(sym_expression),
	1766: uint16(425),
	1767: uint16(1),
	1768: uint16(sym_litType),
	1769: uint16(13),
	1770: uint16(2),
	1771: uint16(anon_sym_UInt),
	1772: uint16(anon_sym_SInt),
	1773: uint16(31),
	1774: uint16(7),
	1775: uint16(sym_literal),
	1776: uint16(sym_sub_field),
	1777: uint16(sym_sub_index),
	1778: uint16(sym_sub_access),
	1779: uint16(sym_mux),
	1780: uint16(sym_conditionally_valid),
	1781: uint16(sym_primitive_operation),
	1782: uint16(477),
	1783: uint16(39),
	1784: uint16(anon_sym_add),
	1785: uint16(anon_sym_sub),
	1786: uint16(anon_sym_mul),
	1787: uint16(anon_sym_div),
	1788: uint16(anon_sym_rem),
	1789: uint16(anon_sym_lt),
	1790: uint16(anon_sym_leq),
	1791: uint16(anon_sym_gt),
	1792: uint16(anon_sym_geq),
	1793: uint16(anon_sym_eq),
	1794: uint16(anon_sym_neq),
	1795: uint16(anon_sym_pad),
	1796: uint16(anon_sym_asUInt),
	1797: uint16(anon_sym_asAsyncReset),
	1798: uint16(anon_sym_asSInt),
	1799: uint16(anon_sym_asClock),
	1800: uint16(anon_sym_shl),
	1801: uint16(anon_sym_shr),
	1802: uint16(anon_sym_dshl),
	1803: uint16(anon_sym_dshlw),
	1804: uint16(anon_sym_dshr),
	1805: uint16(anon_sym_dshrw),
	1806: uint16(anon_sym_cvt),
	1807: uint16(anon_sym_neg),
	1808: uint16(anon_sym_not),
	1809: uint16(anon_sym_and),
	1810: uint16(anon_sym_or),
	1811: uint16(anon_sym_xor),
	1812: uint16(anon_sym_andr),
	1813: uint16(anon_sym_orr),
	1814: uint16(anon_sym_xorr),
	1815: uint16(anon_sym_cat),
	1816: uint16(anon_sym_bits),
	1817: uint16(anon_sym_head),
	1818: uint16(anon_sym_tail),
	1819: uint16(anon_sym_asFixedPoint),
	1820: uint16(anon_sym_bpshl),
	1821: uint16(anon_sym_bpshr),
	1822: uint16(anon_sym_bpset),
	1823: uint16(10),
	1824: uint16(3),
	1825: uint16(1),
	1826: uint16(sym_comment),
	1827: uint16(9),
	1828: uint16(1),
	1829: uint16(sym_identifier),
	1830: uint16(43),
	1831: uint16(1),
	1832: uint16(anon_sym_mux),
	1833: uint16(45),
	1834: uint16(1),
	1835: uint16(anon_sym_validif),
	1836: uint16(143),
	1837: uint16(1),
	1838: uint16(sym_primop),
	1839: uint16(272),
	1840: uint16(1),
	1841: uint16(sym_expression),
	1842: uint16(433),
	1843: uint16(1),
	1844: uint16(sym_litType),
	1845: uint16(13),
	1846: uint16(2),
	1847: uint16(anon_sym_UInt),
	1848: uint16(anon_sym_SInt),
	1849: uint16(160),
	1850: uint16(7),
	1851: uint16(sym_literal),
	1852: uint16(sym_sub_field),
	1853: uint16(sym_sub_index),
	1854: uint16(sym_sub_access),
	1855: uint16(sym_mux),
	1856: uint16(sym_conditionally_valid),
	1857: uint16(sym_primitive_operation),
	1858: uint16(47),
	1859: uint16(39),
	1860: uint16(anon_sym_add),
	1861: uint16(anon_sym_sub),
	1862: uint16(anon_sym_mul),
	1863: uint16(anon_sym_div),
	1864: uint16(anon_sym_rem),
	1865: uint16(anon_sym_lt),
	1866: uint16(anon_sym_leq),
	1867: uint16(anon_sym_gt),
	1868: uint16(anon_sym_geq),
	1869: uint16(anon_sym_eq),
	1870: uint16(anon_sym_neq),
	1871: uint16(anon_sym_pad),
	1872: uint16(anon_sym_asUInt),
	1873: uint16(anon_sym_asAsyncReset),
	1874: uint16(anon_sym_asSInt),
	1875: uint16(anon_sym_asClock),
	1876: uint16(anon_sym_shl),
	1877: uint16(anon_sym_shr),
	1878: uint16(anon_sym_dshl),
	1879: uint16(anon_sym_dshlw),
	1880: uint16(anon_sym_dshr),
	1881: uint16(anon_sym_dshrw),
	1882: uint16(anon_sym_cvt),
	1883: uint16(anon_sym_neg),
	1884: uint16(anon_sym_not),
	1885: uint16(anon_sym_and),
	1886: uint16(anon_sym_or),
	1887: uint16(anon_sym_xor),
	1888: uint16(anon_sym_andr),
	1889: uint16(anon_sym_orr),
	1890: uint16(anon_sym_xorr),
	1891: uint16(anon_sym_cat),
	1892: uint16(anon_sym_bits),
	1893: uint16(anon_sym_head),
	1894: uint16(anon_sym_tail),
	1895: uint16(anon_sym_asFixedPoint),
	1896: uint16(anon_sym_bpshl),
	1897: uint16(anon_sym_bpshr),
	1898: uint16(anon_sym_bpset),
	1899: uint16(10),
	1900: uint16(3),
	1901: uint16(1),
	1902: uint16(sym_comment),
	1903: uint16(471),
	1904: uint16(1),
	1905: uint16(sym_identifier),
	1906: uint16(473),
	1907: uint16(1),
	1908: uint16(anon_sym_mux),
	1909: uint16(475),
	1910: uint16(1),
	1911: uint16(anon_sym_validif),
	1912: uint16(19),
	1913: uint16(1),
	1914: uint16(sym_primop),
	1915: uint16(38),
	1916: uint16(1),
	1917: uint16(sym_expression),
	1918: uint16(425),
	1919: uint16(1),
	1920: uint16(sym_litType),
	1921: uint16(13),
	1922: uint16(2),
	1923: uint16(anon_sym_UInt),
	1924: uint16(anon_sym_SInt),
	1925: uint16(31),
	1926: uint16(7),
	1927: uint16(sym_literal),
	1928: uint16(sym_sub_field),
	1929: uint16(sym_sub_index),
	1930: uint16(sym_sub_access),
	1931: uint16(sym_mux),
	1932: uint16(sym_conditionally_valid),
	1933: uint16(sym_primitive_operation),
	1934: uint16(477),
	1935: uint16(39),
	1936: uint16(anon_sym_add),
	1937: uint16(anon_sym_sub),
	1938: uint16(anon_sym_mul),
	1939: uint16(anon_sym_div),
	1940: uint16(anon_sym_rem),
	1941: uint16(anon_sym_lt),
	1942: uint16(anon_sym_leq),
	1943: uint16(anon_sym_gt),
	1944: uint16(anon_sym_geq),
	1945: uint16(anon_sym_eq),
	1946: uint16(anon_sym_neq),
	1947: uint16(anon_sym_pad),
	1948: uint16(anon_sym_asUInt),
	1949: uint16(anon_sym_asAsyncReset),
	1950: uint16(anon_sym_asSInt),
	1951: uint16(anon_sym_asClock),
	1952: uint16(anon_sym_shl),
	1953: uint16(anon_sym_shr),
	1954: uint16(anon_sym_dshl),
	1955: uint16(anon_sym_dshlw),
	1956: uint16(anon_sym_dshr),
	1957: uint16(anon_sym_dshrw),
	1958: uint16(anon_sym_cvt),
	1959: uint16(anon_sym_neg),
	1960: uint16(anon_sym_not),
	1961: uint16(anon_sym_and),
	1962: uint16(anon_sym_or),
	1963: uint16(anon_sym_xor),
	1964: uint16(anon_sym_andr),
	1965: uint16(anon_sym_orr),
	1966: uint16(anon_sym_xorr),
	1967: uint16(anon_sym_cat),
	1968: uint16(anon_sym_bits),
	1969: uint16(anon_sym_head),
	1970: uint16(anon_sym_tail),
	1971: uint16(anon_sym_asFixedPoint),
	1972: uint16(anon_sym_bpshl),
	1973: uint16(anon_sym_bpshr),
	1974: uint16(anon_sym_bpset),
	1975: uint16(10),
	1976: uint16(3),
	1977: uint16(1),
	1978: uint16(sym_comment),
	1979: uint16(9),
	1980: uint16(1),
	1981: uint16(sym_identifier),
	1982: uint16(43),
	1983: uint16(1),
	1984: uint16(anon_sym_mux),
	1985: uint16(45),
	1986: uint16(1),
	1987: uint16(anon_sym_validif),
	1988: uint16(143),
	1989: uint16(1),
	1990: uint16(sym_primop),
	1991: uint16(300),
	1992: uint16(1),
	1993: uint16(sym_expression),
	1994: uint16(433),
	1995: uint16(1),
	1996: uint16(sym_litType),
	1997: uint16(13),
	1998: uint16(2),
	1999: uint16(anon_sym_UInt),
	2000: uint16(anon_sym_SInt),
	2001: uint16(160),
	2002: uint16(7),
	2003: uint16(sym_literal),
	2004: uint16(sym_sub_field),
	2005: uint16(sym_sub_index),
	2006: uint16(sym_sub_access),
	2007: uint16(sym_mux),
	2008: uint16(sym_conditionally_valid),
	2009: uint16(sym_primitive_operation),
	2010: uint16(47),
	2011: uint16(39),
	2012: uint16(anon_sym_add),
	2013: uint16(anon_sym_sub),
	2014: uint16(anon_sym_mul),
	2015: uint16(anon_sym_div),
	2016: uint16(anon_sym_rem),
	2017: uint16(anon_sym_lt),
	2018: uint16(anon_sym_leq),
	2019: uint16(anon_sym_gt),
	2020: uint16(anon_sym_geq),
	2021: uint16(anon_sym_eq),
	2022: uint16(anon_sym_neq),
	2023: uint16(anon_sym_pad),
	2024: uint16(anon_sym_asUInt),
	2025: uint16(anon_sym_asAsyncReset),
	2026: uint16(anon_sym_asSInt),
	2027: uint16(anon_sym_asClock),
	2028: uint16(anon_sym_shl),
	2029: uint16(anon_sym_shr),
	2030: uint16(anon_sym_dshl),
	2031: uint16(anon_sym_dshlw),
	2032: uint16(anon_sym_dshr),
	2033: uint16(anon_sym_dshrw),
	2034: uint16(anon_sym_cvt),
	2035: uint16(anon_sym_neg),
	2036: uint16(anon_sym_not),
	2037: uint16(anon_sym_and),
	2038: uint16(anon_sym_or),
	2039: uint16(anon_sym_xor),
	2040: uint16(anon_sym_andr),
	2041: uint16(anon_sym_orr),
	2042: uint16(anon_sym_xorr),
	2043: uint16(anon_sym_cat),
	2044: uint16(anon_sym_bits),
	2045: uint16(anon_sym_head),
	2046: uint16(anon_sym_tail),
	2047: uint16(anon_sym_asFixedPoint),
	2048: uint16(anon_sym_bpshl),
	2049: uint16(anon_sym_bpshr),
	2050: uint16(anon_sym_bpset),
	2051: uint16(10),
	2052: uint16(3),
	2053: uint16(1),
	2054: uint16(sym_comment),
	2055: uint16(9),
	2056: uint16(1),
	2057: uint16(sym_identifier),
	2058: uint16(43),
	2059: uint16(1),
	2060: uint16(anon_sym_mux),
	2061: uint16(45),
	2062: uint16(1),
	2063: uint16(anon_sym_validif),
	2064: uint16(143),
	2065: uint16(1),
	2066: uint16(sym_primop),
	2067: uint16(301),
	2068: uint16(1),
	2069: uint16(sym_expression),
	2070: uint16(433),
	2071: uint16(1),
	2072: uint16(sym_litType),
	2073: uint16(13),
	2074: uint16(2),
	2075: uint16(anon_sym_UInt),
	2076: uint16(anon_sym_SInt),
	2077: uint16(160),
	2078: uint16(7),
	2079: uint16(sym_literal),
	2080: uint16(sym_sub_field),
	2081: uint16(sym_sub_index),
	2082: uint16(sym_sub_access),
	2083: uint16(sym_mux),
	2084: uint16(sym_conditionally_valid),
	2085: uint16(sym_primitive_operation),
	2086: uint16(47),
	2087: uint16(39),
	2088: uint16(anon_sym_add),
	2089: uint16(anon_sym_sub),
	2090: uint16(anon_sym_mul),
	2091: uint16(anon_sym_div),
	2092: uint16(anon_sym_rem),
	2093: uint16(anon_sym_lt),
	2094: uint16(anon_sym_leq),
	2095: uint16(anon_sym_gt),
	2096: uint16(anon_sym_geq),
	2097: uint16(anon_sym_eq),
	2098: uint16(anon_sym_neq),
	2099: uint16(anon_sym_pad),
	2100: uint16(anon_sym_asUInt),
	2101: uint16(anon_sym_asAsyncReset),
	2102: uint16(anon_sym_asSInt),
	2103: uint16(anon_sym_asClock),
	2104: uint16(anon_sym_shl),
	2105: uint16(anon_sym_shr),
	2106: uint16(anon_sym_dshl),
	2107: uint16(anon_sym_dshlw),
	2108: uint16(anon_sym_dshr),
	2109: uint16(anon_sym_dshrw),
	2110: uint16(anon_sym_cvt),
	2111: uint16(anon_sym_neg),
	2112: uint16(anon_sym_not),
	2113: uint16(anon_sym_and),
	2114: uint16(anon_sym_or),
	2115: uint16(anon_sym_xor),
	2116: uint16(anon_sym_andr),
	2117: uint16(anon_sym_orr),
	2118: uint16(anon_sym_xorr),
	2119: uint16(anon_sym_cat),
	2120: uint16(anon_sym_bits),
	2121: uint16(anon_sym_head),
	2122: uint16(anon_sym_tail),
	2123: uint16(anon_sym_asFixedPoint),
	2124: uint16(anon_sym_bpshl),
	2125: uint16(anon_sym_bpshr),
	2126: uint16(anon_sym_bpset),
	2127: uint16(10),
	2128: uint16(3),
	2129: uint16(1),
	2130: uint16(sym_comment),
	2131: uint16(9),
	2132: uint16(1),
	2133: uint16(sym_identifier),
	2134: uint16(43),
	2135: uint16(1),
	2136: uint16(anon_sym_mux),
	2137: uint16(45),
	2138: uint16(1),
	2139: uint16(anon_sym_validif),
	2140: uint16(143),
	2141: uint16(1),
	2142: uint16(sym_primop),
	2143: uint16(208),
	2144: uint16(1),
	2145: uint16(sym_expression),
	2146: uint16(433),
	2147: uint16(1),
	2148: uint16(sym_litType),
	2149: uint16(13),
	2150: uint16(2),
	2151: uint16(anon_sym_UInt),
	2152: uint16(anon_sym_SInt),
	2153: uint16(160),
	2154: uint16(7),
	2155: uint16(sym_literal),
	2156: uint16(sym_sub_field),
	2157: uint16(sym_sub_index),
	2158: uint16(sym_sub_access),
	2159: uint16(sym_mux),
	2160: uint16(sym_conditionally_valid),
	2161: uint16(sym_primitive_operation),
	2162: uint16(47),
	2163: uint16(39),
	2164: uint16(anon_sym_add),
	2165: uint16(anon_sym_sub),
	2166: uint16(anon_sym_mul),
	2167: uint16(anon_sym_div),
	2168: uint16(anon_sym_rem),
	2169: uint16(anon_sym_lt),
	2170: uint16(anon_sym_leq),
	2171: uint16(anon_sym_gt),
	2172: uint16(anon_sym_geq),
	2173: uint16(anon_sym_eq),
	2174: uint16(anon_sym_neq),
	2175: uint16(anon_sym_pad),
	2176: uint16(anon_sym_asUInt),
	2177: uint16(anon_sym_asAsyncReset),
	2178: uint16(anon_sym_asSInt),
	2179: uint16(anon_sym_asClock),
	2180: uint16(anon_sym_shl),
	2181: uint16(anon_sym_shr),
	2182: uint16(anon_sym_dshl),
	2183: uint16(anon_sym_dshlw),
	2184: uint16(anon_sym_dshr),
	2185: uint16(anon_sym_dshrw),
	2186: uint16(anon_sym_cvt),
	2187: uint16(anon_sym_neg),
	2188: uint16(anon_sym_not),
	2189: uint16(anon_sym_and),
	2190: uint16(anon_sym_or),
	2191: uint16(anon_sym_xor),
	2192: uint16(anon_sym_andr),
	2193: uint16(anon_sym_orr),
	2194: uint16(anon_sym_xorr),
	2195: uint16(anon_sym_cat),
	2196: uint16(anon_sym_bits),
	2197: uint16(anon_sym_head),
	2198: uint16(anon_sym_tail),
	2199: uint16(anon_sym_asFixedPoint),
	2200: uint16(anon_sym_bpshl),
	2201: uint16(anon_sym_bpshr),
	2202: uint16(anon_sym_bpset),
	2203: uint16(10),
	2204: uint16(3),
	2205: uint16(1),
	2206: uint16(sym_comment),
	2207: uint16(9),
	2208: uint16(1),
	2209: uint16(sym_identifier),
	2210: uint16(43),
	2211: uint16(1),
	2212: uint16(anon_sym_mux),
	2213: uint16(45),
	2214: uint16(1),
	2215: uint16(anon_sym_validif),
	2216: uint16(143),
	2217: uint16(1),
	2218: uint16(sym_primop),
	2219: uint16(243),
	2220: uint16(1),
	2221: uint16(sym_expression),
	2222: uint16(433),
	2223: uint16(1),
	2224: uint16(sym_litType),
	2225: uint16(13),
	2226: uint16(2),
	2227: uint16(anon_sym_UInt),
	2228: uint16(anon_sym_SInt),
	2229: uint16(160),
	2230: uint16(7),
	2231: uint16(sym_literal),
	2232: uint16(sym_sub_field),
	2233: uint16(sym_sub_index),
	2234: uint16(sym_sub_access),
	2235: uint16(sym_mux),
	2236: uint16(sym_conditionally_valid),
	2237: uint16(sym_primitive_operation),
	2238: uint16(47),
	2239: uint16(39),
	2240: uint16(anon_sym_add),
	2241: uint16(anon_sym_sub),
	2242: uint16(anon_sym_mul),
	2243: uint16(anon_sym_div),
	2244: uint16(anon_sym_rem),
	2245: uint16(anon_sym_lt),
	2246: uint16(anon_sym_leq),
	2247: uint16(anon_sym_gt),
	2248: uint16(anon_sym_geq),
	2249: uint16(anon_sym_eq),
	2250: uint16(anon_sym_neq),
	2251: uint16(anon_sym_pad),
	2252: uint16(anon_sym_asUInt),
	2253: uint16(anon_sym_asAsyncReset),
	2254: uint16(anon_sym_asSInt),
	2255: uint16(anon_sym_asClock),
	2256: uint16(anon_sym_shl),
	2257: uint16(anon_sym_shr),
	2258: uint16(anon_sym_dshl),
	2259: uint16(anon_sym_dshlw),
	2260: uint16(anon_sym_dshr),
	2261: uint16(anon_sym_dshrw),
	2262: uint16(anon_sym_cvt),
	2263: uint16(anon_sym_neg),
	2264: uint16(anon_sym_not),
	2265: uint16(anon_sym_and),
	2266: uint16(anon_sym_or),
	2267: uint16(anon_sym_xor),
	2268: uint16(anon_sym_andr),
	2269: uint16(anon_sym_orr),
	2270: uint16(anon_sym_xorr),
	2271: uint16(anon_sym_cat),
	2272: uint16(anon_sym_bits),
	2273: uint16(anon_sym_head),
	2274: uint16(anon_sym_tail),
	2275: uint16(anon_sym_asFixedPoint),
	2276: uint16(anon_sym_bpshl),
	2277: uint16(anon_sym_bpshr),
	2278: uint16(anon_sym_bpset),
	2279: uint16(10),
	2280: uint16(3),
	2281: uint16(1),
	2282: uint16(sym_comment),
	2283: uint16(9),
	2284: uint16(1),
	2285: uint16(sym_identifier),
	2286: uint16(43),
	2287: uint16(1),
	2288: uint16(anon_sym_mux),
	2289: uint16(45),
	2290: uint16(1),
	2291: uint16(anon_sym_validif),
	2292: uint16(143),
	2293: uint16(1),
	2294: uint16(sym_primop),
	2295: uint16(266),
	2296: uint16(1),
	2297: uint16(sym_expression),
	2298: uint16(433),
	2299: uint16(1),
	2300: uint16(sym_litType),
	2301: uint16(13),
	2302: uint16(2),
	2303: uint16(anon_sym_UInt),
	2304: uint16(anon_sym_SInt),
	2305: uint16(160),
	2306: uint16(7),
	2307: uint16(sym_literal),
	2308: uint16(sym_sub_field),
	2309: uint16(sym_sub_index),
	2310: uint16(sym_sub_access),
	2311: uint16(sym_mux),
	2312: uint16(sym_conditionally_valid),
	2313: uint16(sym_primitive_operation),
	2314: uint16(47),
	2315: uint16(39),
	2316: uint16(anon_sym_add),
	2317: uint16(anon_sym_sub),
	2318: uint16(anon_sym_mul),
	2319: uint16(anon_sym_div),
	2320: uint16(anon_sym_rem),
	2321: uint16(anon_sym_lt),
	2322: uint16(anon_sym_leq),
	2323: uint16(anon_sym_gt),
	2324: uint16(anon_sym_geq),
	2325: uint16(anon_sym_eq),
	2326: uint16(anon_sym_neq),
	2327: uint16(anon_sym_pad),
	2328: uint16(anon_sym_asUInt),
	2329: uint16(anon_sym_asAsyncReset),
	2330: uint16(anon_sym_asSInt),
	2331: uint16(anon_sym_asClock),
	2332: uint16(anon_sym_shl),
	2333: uint16(anon_sym_shr),
	2334: uint16(anon_sym_dshl),
	2335: uint16(anon_sym_dshlw),
	2336: uint16(anon_sym_dshr),
	2337: uint16(anon_sym_dshrw),
	2338: uint16(anon_sym_cvt),
	2339: uint16(anon_sym_neg),
	2340: uint16(anon_sym_not),
	2341: uint16(anon_sym_and),
	2342: uint16(anon_sym_or),
	2343: uint16(anon_sym_xor),
	2344: uint16(anon_sym_andr),
	2345: uint16(anon_sym_orr),
	2346: uint16(anon_sym_xorr),
	2347: uint16(anon_sym_cat),
	2348: uint16(anon_sym_bits),
	2349: uint16(anon_sym_head),
	2350: uint16(anon_sym_tail),
	2351: uint16(anon_sym_asFixedPoint),
	2352: uint16(anon_sym_bpshl),
	2353: uint16(anon_sym_bpshr),
	2354: uint16(anon_sym_bpset),
	2355: uint16(10),
	2356: uint16(3),
	2357: uint16(1),
	2358: uint16(sym_comment),
	2359: uint16(9),
	2360: uint16(1),
	2361: uint16(sym_identifier),
	2362: uint16(43),
	2363: uint16(1),
	2364: uint16(anon_sym_mux),
	2365: uint16(45),
	2366: uint16(1),
	2367: uint16(anon_sym_validif),
	2368: uint16(143),
	2369: uint16(1),
	2370: uint16(sym_primop),
	2371: uint16(247),
	2372: uint16(1),
	2373: uint16(sym_expression),
	2374: uint16(433),
	2375: uint16(1),
	2376: uint16(sym_litType),
	2377: uint16(13),
	2378: uint16(2),
	2379: uint16(anon_sym_UInt),
	2380: uint16(anon_sym_SInt),
	2381: uint16(160),
	2382: uint16(7),
	2383: uint16(sym_literal),
	2384: uint16(sym_sub_field),
	2385: uint16(sym_sub_index),
	2386: uint16(sym_sub_access),
	2387: uint16(sym_mux),
	2388: uint16(sym_conditionally_valid),
	2389: uint16(sym_primitive_operation),
	2390: uint16(47),
	2391: uint16(39),
	2392: uint16(anon_sym_add),
	2393: uint16(anon_sym_sub),
	2394: uint16(anon_sym_mul),
	2395: uint16(anon_sym_div),
	2396: uint16(anon_sym_rem),
	2397: uint16(anon_sym_lt),
	2398: uint16(anon_sym_leq),
	2399: uint16(anon_sym_gt),
	2400: uint16(anon_sym_geq),
	2401: uint16(anon_sym_eq),
	2402: uint16(anon_sym_neq),
	2403: uint16(anon_sym_pad),
	2404: uint16(anon_sym_asUInt),
	2405: uint16(anon_sym_asAsyncReset),
	2406: uint16(anon_sym_asSInt),
	2407: uint16(anon_sym_asClock),
	2408: uint16(anon_sym_shl),
	2409: uint16(anon_sym_shr),
	2410: uint16(anon_sym_dshl),
	2411: uint16(anon_sym_dshlw),
	2412: uint16(anon_sym_dshr),
	2413: uint16(anon_sym_dshrw),
	2414: uint16(anon_sym_cvt),
	2415: uint16(anon_sym_neg),
	2416: uint16(anon_sym_not),
	2417: uint16(anon_sym_and),
	2418: uint16(anon_sym_or),
	2419: uint16(anon_sym_xor),
	2420: uint16(anon_sym_andr),
	2421: uint16(anon_sym_orr),
	2422: uint16(anon_sym_xorr),
	2423: uint16(anon_sym_cat),
	2424: uint16(anon_sym_bits),
	2425: uint16(anon_sym_head),
	2426: uint16(anon_sym_tail),
	2427: uint16(anon_sym_asFixedPoint),
	2428: uint16(anon_sym_bpshl),
	2429: uint16(anon_sym_bpshr),
	2430: uint16(anon_sym_bpset),
	2431: uint16(10),
	2432: uint16(3),
	2433: uint16(1),
	2434: uint16(sym_comment),
	2435: uint16(471),
	2436: uint16(1),
	2437: uint16(sym_identifier),
	2438: uint16(473),
	2439: uint16(1),
	2440: uint16(anon_sym_mux),
	2441: uint16(475),
	2442: uint16(1),
	2443: uint16(anon_sym_validif),
	2444: uint16(19),
	2445: uint16(1),
	2446: uint16(sym_primop),
	2447: uint16(28),
	2448: uint16(1),
	2449: uint16(sym_expression),
	2450: uint16(425),
	2451: uint16(1),
	2452: uint16(sym_litType),
	2453: uint16(13),
	2454: uint16(2),
	2455: uint16(anon_sym_UInt),
	2456: uint16(anon_sym_SInt),
	2457: uint16(31),
	2458: uint16(7),
	2459: uint16(sym_literal),
	2460: uint16(sym_sub_field),
	2461: uint16(sym_sub_index),
	2462: uint16(sym_sub_access),
	2463: uint16(sym_mux),
	2464: uint16(sym_conditionally_valid),
	2465: uint16(sym_primitive_operation),
	2466: uint16(477),
	2467: uint16(39),
	2468: uint16(anon_sym_add),
	2469: uint16(anon_sym_sub),
	2470: uint16(anon_sym_mul),
	2471: uint16(anon_sym_div),
	2472: uint16(anon_sym_rem),
	2473: uint16(anon_sym_lt),
	2474: uint16(anon_sym_leq),
	2475: uint16(anon_sym_gt),
	2476: uint16(anon_sym_geq),
	2477: uint16(anon_sym_eq),
	2478: uint16(anon_sym_neq),
	2479: uint16(anon_sym_pad),
	2480: uint16(anon_sym_asUInt),
	2481: uint16(anon_sym_asAsyncReset),
	2482: uint16(anon_sym_asSInt),
	2483: uint16(anon_sym_asClock),
	2484: uint16(anon_sym_shl),
	2485: uint16(anon_sym_shr),
	2486: uint16(anon_sym_dshl),
	2487: uint16(anon_sym_dshlw),
	2488: uint16(anon_sym_dshr),
	2489: uint16(anon_sym_dshrw),
	2490: uint16(anon_sym_cvt),
	2491: uint16(anon_sym_neg),
	2492: uint16(anon_sym_not),
	2493: uint16(anon_sym_and),
	2494: uint16(anon_sym_or),
	2495: uint16(anon_sym_xor),
	2496: uint16(anon_sym_andr),
	2497: uint16(anon_sym_orr),
	2498: uint16(anon_sym_xorr),
	2499: uint16(anon_sym_cat),
	2500: uint16(anon_sym_bits),
	2501: uint16(anon_sym_head),
	2502: uint16(anon_sym_tail),
	2503: uint16(anon_sym_asFixedPoint),
	2504: uint16(anon_sym_bpshl),
	2505: uint16(anon_sym_bpshr),
	2506: uint16(anon_sym_bpset),
	2507: uint16(10),
	2508: uint16(3),
	2509: uint16(1),
	2510: uint16(sym_comment),
	2511: uint16(9),
	2512: uint16(1),
	2513: uint16(sym_identifier),
	2514: uint16(43),
	2515: uint16(1),
	2516: uint16(anon_sym_mux),
	2517: uint16(45),
	2518: uint16(1),
	2519: uint16(anon_sym_validif),
	2520: uint16(143),
	2521: uint16(1),
	2522: uint16(sym_primop),
	2523: uint16(245),
	2524: uint16(1),
	2525: uint16(sym_expression),
	2526: uint16(433),
	2527: uint16(1),
	2528: uint16(sym_litType),
	2529: uint16(13),
	2530: uint16(2),
	2531: uint16(anon_sym_UInt),
	2532: uint16(anon_sym_SInt),
	2533: uint16(160),
	2534: uint16(7),
	2535: uint16(sym_literal),
	2536: uint16(sym_sub_field),
	2537: uint16(sym_sub_index),
	2538: uint16(sym_sub_access),
	2539: uint16(sym_mux),
	2540: uint16(sym_conditionally_valid),
	2541: uint16(sym_primitive_operation),
	2542: uint16(47),
	2543: uint16(39),
	2544: uint16(anon_sym_add),
	2545: uint16(anon_sym_sub),
	2546: uint16(anon_sym_mul),
	2547: uint16(anon_sym_div),
	2548: uint16(anon_sym_rem),
	2549: uint16(anon_sym_lt),
	2550: uint16(anon_sym_leq),
	2551: uint16(anon_sym_gt),
	2552: uint16(anon_sym_geq),
	2553: uint16(anon_sym_eq),
	2554: uint16(anon_sym_neq),
	2555: uint16(anon_sym_pad),
	2556: uint16(anon_sym_asUInt),
	2557: uint16(anon_sym_asAsyncReset),
	2558: uint16(anon_sym_asSInt),
	2559: uint16(anon_sym_asClock),
	2560: uint16(anon_sym_shl),
	2561: uint16(anon_sym_shr),
	2562: uint16(anon_sym_dshl),
	2563: uint16(anon_sym_dshlw),
	2564: uint16(anon_sym_dshr),
	2565: uint16(anon_sym_dshrw),
	2566: uint16(anon_sym_cvt),
	2567: uint16(anon_sym_neg),
	2568: uint16(anon_sym_not),
	2569: uint16(anon_sym_and),
	2570: uint16(anon_sym_or),
	2571: uint16(anon_sym_xor),
	2572: uint16(anon_sym_andr),
	2573: uint16(anon_sym_orr),
	2574: uint16(anon_sym_xorr),
	2575: uint16(anon_sym_cat),
	2576: uint16(anon_sym_bits),
	2577: uint16(anon_sym_head),
	2578: uint16(anon_sym_tail),
	2579: uint16(anon_sym_asFixedPoint),
	2580: uint16(anon_sym_bpshl),
	2581: uint16(anon_sym_bpshr),
	2582: uint16(anon_sym_bpset),
	2583: uint16(10),
	2584: uint16(3),
	2585: uint16(1),
	2586: uint16(sym_comment),
	2587: uint16(9),
	2588: uint16(1),
	2589: uint16(sym_identifier),
	2590: uint16(43),
	2591: uint16(1),
	2592: uint16(anon_sym_mux),
	2593: uint16(45),
	2594: uint16(1),
	2595: uint16(anon_sym_validif),
	2596: uint16(143),
	2597: uint16(1),
	2598: uint16(sym_primop),
	2599: uint16(241),
	2600: uint16(1),
	2601: uint16(sym_expression),
	2602: uint16(433),
	2603: uint16(1),
	2604: uint16(sym_litType),
	2605: uint16(13),
	2606: uint16(2),
	2607: uint16(anon_sym_UInt),
	2608: uint16(anon_sym_SInt),
	2609: uint16(160),
	2610: uint16(7),
	2611: uint16(sym_literal),
	2612: uint16(sym_sub_field),
	2613: uint16(sym_sub_index),
	2614: uint16(sym_sub_access),
	2615: uint16(sym_mux),
	2616: uint16(sym_conditionally_valid),
	2617: uint16(sym_primitive_operation),
	2618: uint16(47),
	2619: uint16(39),
	2620: uint16(anon_sym_add),
	2621: uint16(anon_sym_sub),
	2622: uint16(anon_sym_mul),
	2623: uint16(anon_sym_div),
	2624: uint16(anon_sym_rem),
	2625: uint16(anon_sym_lt),
	2626: uint16(anon_sym_leq),
	2627: uint16(anon_sym_gt),
	2628: uint16(anon_sym_geq),
	2629: uint16(anon_sym_eq),
	2630: uint16(anon_sym_neq),
	2631: uint16(anon_sym_pad),
	2632: uint16(anon_sym_asUInt),
	2633: uint16(anon_sym_asAsyncReset),
	2634: uint16(anon_sym_asSInt),
	2635: uint16(anon_sym_asClock),
	2636: uint16(anon_sym_shl),
	2637: uint16(anon_sym_shr),
	2638: uint16(anon_sym_dshl),
	2639: uint16(anon_sym_dshlw),
	2640: uint16(anon_sym_dshr),
	2641: uint16(anon_sym_dshrw),
	2642: uint16(anon_sym_cvt),
	2643: uint16(anon_sym_neg),
	2644: uint16(anon_sym_not),
	2645: uint16(anon_sym_and),
	2646: uint16(anon_sym_or),
	2647: uint16(anon_sym_xor),
	2648: uint16(anon_sym_andr),
	2649: uint16(anon_sym_orr),
	2650: uint16(anon_sym_xorr),
	2651: uint16(anon_sym_cat),
	2652: uint16(anon_sym_bits),
	2653: uint16(anon_sym_head),
	2654: uint16(anon_sym_tail),
	2655: uint16(anon_sym_asFixedPoint),
	2656: uint16(anon_sym_bpshl),
	2657: uint16(anon_sym_bpshr),
	2658: uint16(anon_sym_bpset),
	2659: uint16(10),
	2660: uint16(3),
	2661: uint16(1),
	2662: uint16(sym_comment),
	2663: uint16(9),
	2664: uint16(1),
	2665: uint16(sym_identifier),
	2666: uint16(43),
	2667: uint16(1),
	2668: uint16(anon_sym_mux),
	2669: uint16(45),
	2670: uint16(1),
	2671: uint16(anon_sym_validif),
	2672: uint16(143),
	2673: uint16(1),
	2674: uint16(sym_primop),
	2675: uint16(242),
	2676: uint16(1),
	2677: uint16(sym_expression),
	2678: uint16(433),
	2679: uint16(1),
	2680: uint16(sym_litType),
	2681: uint16(13),
	2682: uint16(2),
	2683: uint16(anon_sym_UInt),
	2684: uint16(anon_sym_SInt),
	2685: uint16(160),
	2686: uint16(7),
	2687: uint16(sym_literal),
	2688: uint16(sym_sub_field),
	2689: uint16(sym_sub_index),
	2690: uint16(sym_sub_access),
	2691: uint16(sym_mux),
	2692: uint16(sym_conditionally_valid),
	2693: uint16(sym_primitive_operation),
	2694: uint16(47),
	2695: uint16(39),
	2696: uint16(anon_sym_add),
	2697: uint16(anon_sym_sub),
	2698: uint16(anon_sym_mul),
	2699: uint16(anon_sym_div),
	2700: uint16(anon_sym_rem),
	2701: uint16(anon_sym_lt),
	2702: uint16(anon_sym_leq),
	2703: uint16(anon_sym_gt),
	2704: uint16(anon_sym_geq),
	2705: uint16(anon_sym_eq),
	2706: uint16(anon_sym_neq),
	2707: uint16(anon_sym_pad),
	2708: uint16(anon_sym_asUInt),
	2709: uint16(anon_sym_asAsyncReset),
	2710: uint16(anon_sym_asSInt),
	2711: uint16(anon_sym_asClock),
	2712: uint16(anon_sym_shl),
	2713: uint16(anon_sym_shr),
	2714: uint16(anon_sym_dshl),
	2715: uint16(anon_sym_dshlw),
	2716: uint16(anon_sym_dshr),
	2717: uint16(anon_sym_dshrw),
	2718: uint16(anon_sym_cvt),
	2719: uint16(anon_sym_neg),
	2720: uint16(anon_sym_not),
	2721: uint16(anon_sym_and),
	2722: uint16(anon_sym_or),
	2723: uint16(anon_sym_xor),
	2724: uint16(anon_sym_andr),
	2725: uint16(anon_sym_orr),
	2726: uint16(anon_sym_xorr),
	2727: uint16(anon_sym_cat),
	2728: uint16(anon_sym_bits),
	2729: uint16(anon_sym_head),
	2730: uint16(anon_sym_tail),
	2731: uint16(anon_sym_asFixedPoint),
	2732: uint16(anon_sym_bpshl),
	2733: uint16(anon_sym_bpshr),
	2734: uint16(anon_sym_bpset),
	2735: uint16(8),
	2736: uint16(3),
	2737: uint16(1),
	2738: uint16(sym_comment),
	2739: uint16(483),
	2740: uint16(1),
	2741: uint16(sym_double),
	2742: uint16(485),
	2743: uint16(1),
	2744: uint16(anon_sym_DQUOTE),
	2745: uint16(487),
	2746: uint16(1),
	2747: uint16(anon_sym_SQUOTE),
	2748: uint16(479),
	2749: uint16(2),
	2750: uint16(anon_sym_0),
	2751: uint16(aux_sym_uint_token1),
	2752: uint16(481),
	2753: uint16(2),
	2754: uint16(anon_sym_PLUS),
	2755: uint16(anon_sym_DASH),
	2756: uint16(343),
	2757: uint16(2),
	2758: uint16(sym_uint),
	2759: uint16(sym_sint),
	2760: uint16(357),
	2761: uint16(3),
	2762: uint16(sym_number),
	2763: uint16(sym_string),
	2764: uint16(sym_raw_string),
	2765: uint16(9),
	2766: uint16(3),
	2767: uint16(1),
	2768: uint16(sym_comment),
	2769: uint16(53),
	2770: uint16(1),
	2771: uint16(sym__dedent),
	2772: uint16(491),
	2773: uint16(1),
	2774: uint16(anon_sym_defname),
	2775: uint16(493),
	2776: uint16(1),
	2777: uint16(anon_sym_parameter),
	2778: uint16(225),
	2779: uint16(1),
	2780: uint16(sym_defname),
	2781: uint16(429),
	2782: uint16(1),
	2783: uint16(sym_dir),
	2784: uint16(489),
	2785: uint16(2),
	2786: uint16(anon_sym_input),
	2787: uint16(anon_sym_output),
	2788: uint16(137),
	2789: uint16(2),
	2790: uint16(sym_port),
	2791: uint16(aux_sym_module_repeat1),
	2792: uint16(230),
	2793: uint16(2),
	2794: uint16(sym_parameter),
	2795: uint16(aux_sym_module_repeat3),
	2796: uint16(7),
	2797: uint16(3),
	2798: uint16(1),
	2799: uint16(sym_comment),
	2800: uint16(495),
	2801: uint16(1),
	2802: uint16(anon_sym_data_DASHtype),
	2803: uint16(499),
	2804: uint16(1),
	2805: uint16(anon_sym_read_DASHunder_DASHwrite),
	2806: uint16(503),
	2807: uint16(1),
	2808: uint16(sym__dedent),
	2809: uint16(133),
	2810: uint16(2),
	2811: uint16(sym_memory_field),
	2812: uint16(aux_sym_memory_repeat1),
	2813: uint16(497),
	2814: uint16(3),
	2815: uint16(anon_sym_depth),
	2816: uint16(anon_sym_read_DASHlatency),
	2817: uint16(anon_sym_write_DASHlatency),
	2818: uint16(501),
	2819: uint16(3),
	2820: uint16(anon_sym_reader),
	2821: uint16(anon_sym_writer),
	2822: uint16(anon_sym_readwriter),
	2823: uint16(7),
	2824: uint16(3),
	2825: uint16(1),
	2826: uint16(sym_comment),
	2827: uint16(495),
	2828: uint16(1),
	2829: uint16(anon_sym_data_DASHtype),
	2830: uint16(499),
	2831: uint16(1),
	2832: uint16(anon_sym_read_DASHunder_DASHwrite),
	2833: uint16(505),
	2834: uint16(1),
	2835: uint16(sym__dedent),
	2836: uint16(139),
	2837: uint16(2),
	2838: uint16(sym_memory_field),
	2839: uint16(aux_sym_memory_repeat1),
	2840: uint16(497),
	2841: uint16(3),
	2842: uint16(anon_sym_depth),
	2843: uint16(anon_sym_read_DASHlatency),
	2844: uint16(anon_sym_write_DASHlatency),
	2845: uint16(501),
	2846: uint16(3),
	2847: uint16(anon_sym_reader),
	2848: uint16(anon_sym_writer),
	2849: uint16(anon_sym_readwriter),
	2850: uint16(7),
	2851: uint16(3),
	2852: uint16(1),
	2853: uint16(sym_comment),
	2854: uint16(495),
	2855: uint16(1),
	2856: uint16(anon_sym_data_DASHtype),
	2857: uint16(499),
	2858: uint16(1),
	2859: uint16(anon_sym_read_DASHunder_DASHwrite),
	2860: uint16(505),
	2861: uint16(1),
	2862: uint16(sym__dedent),
	2863: uint16(138),
	2864: uint16(2),
	2865: uint16(sym_memory_field),
	2866: uint16(aux_sym_memory_repeat1),
	2867: uint16(497),
	2868: uint16(3),
	2869: uint16(anon_sym_depth),
	2870: uint16(anon_sym_read_DASHlatency),
	2871: uint16(anon_sym_write_DASHlatency),
	2872: uint16(501),
	2873: uint16(3),
	2874: uint16(anon_sym_reader),
	2875: uint16(anon_sym_writer),
	2876: uint16(anon_sym_readwriter),
	2877: uint16(8),
	2878: uint16(3),
	2879: uint16(1),
	2880: uint16(sym_comment),
	2881: uint16(507),
	2882: uint16(1),
	2883: uint16(anon_sym_const),
	2884: uint16(511),
	2885: uint16(1),
	2886: uint16(anon_sym_Fixed),
	2887: uint16(515),
	2888: uint16(1),
	2889: uint16(anon_sym_LBRACE),
	2890: uint16(164),
	2891: uint16(1),
	2892: uint16(sym_qualifier),
	2893: uint16(298),
	2894: uint16(1),
	2895: uint16(sym_type),
	2896: uint16(509),
	2897: uint16(3),
	2898: uint16(anon_sym_UInt),
	2899: uint16(anon_sym_SInt),
	2900: uint16(anon_sym_Analog),
	2901: uint16(513),
	2902: uint16(3),
	2903: uint16(anon_sym_Clock),
	2904: uint16(anon_sym_AsyncReset),
	2905: uint16(anon_sym_Reset),
	2906: uint16(9),
	2907: uint16(3),
	2908: uint16(1),
	2909: uint16(sym_comment),
	2910: uint16(49),
	2911: uint16(1),
	2912: uint16(sym__dedent),
	2913: uint16(491),
	2914: uint16(1),
	2915: uint16(anon_sym_defname),
	2916: uint16(493),
	2917: uint16(1),
	2918: uint16(anon_sym_parameter),
	2919: uint16(237),
	2920: uint16(1),
	2921: uint16(sym_defname),
	2922: uint16(429),
	2923: uint16(1),
	2924: uint16(sym_dir),
	2925: uint16(489),
	2926: uint16(2),
	2927: uint16(anon_sym_input),
	2928: uint16(anon_sym_output),
	2929: uint16(140),
	2930: uint16(2),
	2931: uint16(sym_port),
	2932: uint16(aux_sym_module_repeat1),
	2933: uint16(236),
	2934: uint16(2),
	2935: uint16(sym_parameter),
	2936: uint16(aux_sym_module_repeat3),
	2937: uint16(9),
	2938: uint16(3),
	2939: uint16(1),
	2940: uint16(sym_comment),
	2941: uint16(49),
	2942: uint16(1),
	2943: uint16(sym__dedent),
	2944: uint16(491),
	2945: uint16(1),
	2946: uint16(anon_sym_defname),
	2947: uint16(493),
	2948: uint16(1),
	2949: uint16(anon_sym_parameter),
	2950: uint16(237),
	2951: uint16(1),
	2952: uint16(sym_defname),
	2953: uint16(429),
	2954: uint16(1),
	2955: uint16(sym_dir),
	2956: uint16(489),
	2957: uint16(2),
	2958: uint16(anon_sym_input),
	2959: uint16(anon_sym_output),
	2960: uint16(16),
	2961: uint16(2),
	2962: uint16(sym_port),
	2963: uint16(aux_sym_module_repeat1),
	2964: uint16(236),
	2965: uint16(2),
	2966: uint16(sym_parameter),
	2967: uint16(aux_sym_module_repeat3),
	2968: uint16(7),
	2969: uint16(3),
	2970: uint16(1),
	2971: uint16(sym_comment),
	2972: uint16(495),
	2973: uint16(1),
	2974: uint16(anon_sym_data_DASHtype),
	2975: uint16(499),
	2976: uint16(1),
	2977: uint16(anon_sym_read_DASHunder_DASHwrite),
	2978: uint16(517),
	2979: uint16(1),
	2980: uint16(sym__dedent),
	2981: uint16(139),
	2982: uint16(2),
	2983: uint16(sym_memory_field),
	2984: uint16(aux_sym_memory_repeat1),
	2985: uint16(497),
	2986: uint16(3),
	2987: uint16(anon_sym_depth),
	2988: uint16(anon_sym_read_DASHlatency),
	2989: uint16(anon_sym_write_DASHlatency),
	2990: uint16(501),
	2991: uint16(3),
	2992: uint16(anon_sym_reader),
	2993: uint16(anon_sym_writer),
	2994: uint16(anon_sym_readwriter),
	2995: uint16(7),
	2996: uint16(3),
	2997: uint16(1),
	2998: uint16(sym_comment),
	2999: uint16(519),
	3000: uint16(1),
	3001: uint16(anon_sym_data_DASHtype),
	3002: uint16(525),
	3003: uint16(1),
	3004: uint16(anon_sym_read_DASHunder_DASHwrite),
	3005: uint16(531),
	3006: uint16(1),
	3007: uint16(sym__dedent),
	3008: uint16(139),
	3009: uint16(2),
	3010: uint16(sym_memory_field),
	3011: uint16(aux_sym_memory_repeat1),
	3012: uint16(522),
	3013: uint16(3),
	3014: uint16(anon_sym_depth),
	3015: uint16(anon_sym_read_DASHlatency),
	3016: uint16(anon_sym_write_DASHlatency),
	3017: uint16(528),
	3018: uint16(3),
	3019: uint16(anon_sym_reader),
	3020: uint16(anon_sym_writer),
	3021: uint16(anon_sym_readwriter),
	3022: uint16(9),
	3023: uint16(3),
	3024: uint16(1),
	3025: uint16(sym_comment),
	3026: uint16(51),
	3027: uint16(1),
	3028: uint16(sym__dedent),
	3029: uint16(491),
	3030: uint16(1),
	3031: uint16(anon_sym_defname),
	3032: uint16(493),
	3033: uint16(1),
	3034: uint16(anon_sym_parameter),
	3035: uint16(215),
	3036: uint16(1),
	3037: uint16(sym_defname),
	3038: uint16(429),
	3039: uint16(1),
	3040: uint16(sym_dir),
	3041: uint16(489),
	3042: uint16(2),
	3043: uint16(anon_sym_input),
	3044: uint16(anon_sym_output),
	3045: uint16(16),
	3046: uint16(2),
	3047: uint16(sym_port),
	3048: uint16(aux_sym_module_repeat1),
	3049: uint16(221),
	3050: uint16(2),
	3051: uint16(sym_parameter),
	3052: uint16(aux_sym_module_repeat3),
	3053: uint16(3),
	3054: uint16(3),
	3055: uint16(1),
	3056: uint16(sym_comment),
	3057: uint16(533),
	3058: uint16(1),
	3059: uint16(anon_sym_LPAREN),
	3060: uint16(143),
	3061: uint16(9),
	3062: uint16(anon_sym_COLON),
	3063: uint16(anon_sym_COMMA),
	3064: uint16(anon_sym_LBRACK),
	3065: uint16(anon_sym_RBRACK),
	3066: uint16(anon_sym_RPAREN),
	3067: uint16(anon_sym_LT_EQ),
	3068: uint16(anon_sym_LT_DASH),
	3069: uint16(anon_sym_is),
	3070: uint16(anon_sym_DOT),
	3071: uint16(2),
	3072: uint16(3),
	3073: uint16(1),
	3074: uint16(sym_comment),
	3075: uint16(139),
	3076: uint16(10),
	3077: uint16(anon_sym_COLON),
	3078: uint16(anon_sym_COMMA),
	3079: uint16(anon_sym_LBRACK),
	3080: uint16(anon_sym_RBRACK),
	3081: uint16(anon_sym_LPAREN),
	3082: uint16(anon_sym_RPAREN),
	3083: uint16(anon_sym_LT_EQ),
	3084: uint16(anon_sym_LT_DASH),
	3085: uint16(anon_sym_is),
	3086: uint16(anon_sym_DOT),
	3087: uint16(3),
	3088: uint16(3),
	3089: uint16(1),
	3090: uint16(sym_comment),
	3091: uint16(535),
	3092: uint16(1),
	3093: uint16(anon_sym_LPAREN),
	3094: uint16(143),
	3095: uint16(9),
	3096: uint16(anon_sym_COLON),
	3097: uint16(anon_sym_COMMA),
	3098: uint16(anon_sym_LBRACK),
	3099: uint16(anon_sym_RBRACK),
	3100: uint16(anon_sym_RPAREN),
	3101: uint16(anon_sym_LT_EQ),
	3102: uint16(anon_sym_LT_DASH),
	3103: uint16(anon_sym_is),
	3104: uint16(anon_sym_DOT),
	3105: uint16(2),
	3106: uint16(3),
	3107: uint16(1),
	3108: uint16(sym_comment),
	3109: uint16(163),
	3110: uint16(10),
	3111: uint16(anon_sym_COLON),
	3112: uint16(anon_sym_GT),
	3113: uint16(anon_sym_COMMA),
	3114: uint16(anon_sym_LBRACK),
	3115: uint16(anon_sym_RBRACK),
	3116: uint16(anon_sym_RPAREN),
	3117: uint16(anon_sym_LT_EQ),
	3118: uint16(anon_sym_LT_DASH),
	3119: uint16(anon_sym_is),
	3120: uint16(anon_sym_DOT),
	3121: uint16(2),
	3122: uint16(3),
	3123: uint16(1),
	3124: uint16(sym_comment),
	3125: uint16(537),
	3126: uint16(9),
	3127: uint16(sym__dedent),
	3128: uint16(anon_sym_data_DASHtype),
	3129: uint16(anon_sym_depth),
	3130: uint16(anon_sym_read_DASHlatency),
	3131: uint16(anon_sym_write_DASHlatency),
	3132: uint16(anon_sym_read_DASHunder_DASHwrite),
	3133: uint16(anon_sym_reader),
	3134: uint16(anon_sym_writer),
	3135: uint16(anon_sym_readwriter),
	3136: uint16(2),
	3137: uint16(3),
	3138: uint16(1),
	3139: uint16(sym_comment),
	3140: uint16(151),
	3141: uint16(9),
	3142: uint16(anon_sym_COLON),
	3143: uint16(anon_sym_COMMA),
	3144: uint16(anon_sym_LBRACK),
	3145: uint16(anon_sym_RBRACK),
	3146: uint16(anon_sym_RPAREN),
	3147: uint16(anon_sym_LT_EQ),
	3148: uint16(anon_sym_LT_DASH),
	3149: uint16(anon_sym_is),
	3150: uint16(anon_sym_DOT),
	3151: uint16(2),
	3152: uint16(3),
	3153: uint16(1),
	3154: uint16(sym_comment),
	3155: uint16(175),
	3156: uint16(9),
	3157: uint16(anon_sym_COLON),
	3158: uint16(anon_sym_COMMA),
	3159: uint16(anon_sym_LBRACK),
	3160: uint16(anon_sym_RBRACK),
	3161: uint16(anon_sym_RPAREN),
	3162: uint16(anon_sym_LT_EQ),
	3163: uint16(anon_sym_LT_DASH),
	3164: uint16(anon_sym_is),
	3165: uint16(anon_sym_DOT),
	3166: uint16(6),
	3167: uint16(3),
	3168: uint16(1),
	3169: uint16(sym_comment),
	3170: uint16(511),
	3171: uint16(1),
	3172: uint16(anon_sym_Fixed),
	3173: uint16(515),
	3174: uint16(1),
	3175: uint16(anon_sym_LBRACE),
	3176: uint16(308),
	3177: uint16(1),
	3178: uint16(sym_type),
	3179: uint16(509),
	3180: uint16(3),
	3181: uint16(anon_sym_UInt),
	3182: uint16(anon_sym_SInt),
	3183: uint16(anon_sym_Analog),
	3184: uint16(513),
	3185: uint16(3),
	3186: uint16(anon_sym_Clock),
	3187: uint16(anon_sym_AsyncReset),
	3188: uint16(anon_sym_Reset),
	3189: uint16(6),
	3190: uint16(3),
	3191: uint16(1),
	3192: uint16(sym_comment),
	3193: uint16(541),
	3194: uint16(1),
	3195: uint16(anon_sym_Fixed),
	3196: uint16(545),
	3197: uint16(1),
	3198: uint16(anon_sym_LBRACE),
	3199: uint16(270),
	3200: uint16(1),
	3201: uint16(sym_type),
	3202: uint16(539),
	3203: uint16(3),
	3204: uint16(anon_sym_UInt),
	3205: uint16(anon_sym_SInt),
	3206: uint16(anon_sym_Analog),
	3207: uint16(543),
	3208: uint16(3),
	3209: uint16(anon_sym_Clock),
	3210: uint16(anon_sym_AsyncReset),
	3211: uint16(anon_sym_Reset),
	3212: uint16(6),
	3213: uint16(3),
	3214: uint16(1),
	3215: uint16(sym_comment),
	3216: uint16(549),
	3217: uint16(1),
	3218: uint16(anon_sym_Fixed),
	3219: uint16(553),
	3220: uint16(1),
	3221: uint16(anon_sym_LBRACE),
	3222: uint16(44),
	3223: uint16(1),
	3224: uint16(sym_type),
	3225: uint16(547),
	3226: uint16(3),
	3227: uint16(anon_sym_UInt),
	3228: uint16(anon_sym_SInt),
	3229: uint16(anon_sym_Analog),
	3230: uint16(551),
	3231: uint16(3),
	3232: uint16(anon_sym_Clock),
	3233: uint16(anon_sym_AsyncReset),
	3234: uint16(anon_sym_Reset),
	3235: uint16(8),
	3236: uint16(3),
	3237: uint16(1),
	3238: uint16(sym_comment),
	3239: uint16(557),
	3240: uint16(1),
	3241: uint16(anon_sym_RBRACE),
	3242: uint16(559),
	3243: uint16(1),
	3244: uint16(anon_sym_flip),
	3245: uint16(152),
	3246: uint16(1),
	3247: uint16(sym_uint),
	3248: uint16(252),
	3249: uint16(1),
	3250: uint16(sym_field),
	3251: uint16(347),
	3252: uint16(1),
	3253: uint16(sym_field_id),
	3254: uint16(555),
	3255: uint16(2),
	3256: uint16(sym_identifier),
	3257: uint16(sym_relaxed_identifier),
	3258: uint16(561),
	3259: uint16(2),
	3260: uint16(anon_sym_0),
	3261: uint16(aux_sym_uint_token1),
	3262: uint16(2),
	3263: uint16(3),
	3264: uint16(1),
	3265: uint16(sym_comment),
	3266: uint16(167),
	3267: uint16(9),
	3268: uint16(anon_sym_COLON),
	3269: uint16(anon_sym_COMMA),
	3270: uint16(anon_sym_LBRACK),
	3271: uint16(anon_sym_RBRACK),
	3272: uint16(anon_sym_RPAREN),
	3273: uint16(anon_sym_LT_EQ),
	3274: uint16(anon_sym_LT_DASH),
	3275: uint16(anon_sym_is),
	3276: uint16(anon_sym_DOT),
	3277: uint16(8),
	3278: uint16(3),
	3279: uint16(1),
	3280: uint16(sym_comment),
	3281: uint16(559),
	3282: uint16(1),
	3283: uint16(anon_sym_flip),
	3284: uint16(563),
	3285: uint16(1),
	3286: uint16(anon_sym_RBRACE),
	3287: uint16(152),
	3288: uint16(1),
	3289: uint16(sym_uint),
	3290: uint16(262),
	3291: uint16(1),
	3292: uint16(sym_field),
	3293: uint16(347),
	3294: uint16(1),
	3295: uint16(sym_field_id),
	3296: uint16(555),
	3297: uint16(2),
	3298: uint16(sym_identifier),
	3299: uint16(sym_relaxed_identifier),
	3300: uint16(561),
	3301: uint16(2),
	3302: uint16(anon_sym_0),
	3303: uint16(aux_sym_uint_token1),
	3304: uint16(2),
	3305: uint16(3),
	3306: uint16(1),
	3307: uint16(sym_comment),
	3308: uint16(159),
	3309: uint16(9),
	3310: uint16(anon_sym_COLON),
	3311: uint16(anon_sym_COMMA),
	3312: uint16(anon_sym_LBRACK),
	3313: uint16(anon_sym_RBRACK),
	3314: uint16(anon_sym_RPAREN),
	3315: uint16(anon_sym_LT_EQ),
	3316: uint16(anon_sym_LT_DASH),
	3317: uint16(anon_sym_is),
	3318: uint16(anon_sym_DOT),
	3319: uint16(2),
	3320: uint16(3),
	3321: uint16(1),
	3322: uint16(sym_comment),
	3323: uint16(565),
	3324: uint16(9),
	3325: uint16(sym__dedent),
	3326: uint16(anon_sym_data_DASHtype),
	3327: uint16(anon_sym_depth),
	3328: uint16(anon_sym_read_DASHlatency),
	3329: uint16(anon_sym_write_DASHlatency),
	3330: uint16(anon_sym_read_DASHunder_DASHwrite),
	3331: uint16(anon_sym_reader),
	3332: uint16(anon_sym_writer),
	3333: uint16(anon_sym_readwriter),
	3334: uint16(6),
	3335: uint16(3),
	3336: uint16(1),
	3337: uint16(sym_comment),
	3338: uint16(549),
	3339: uint16(1),
	3340: uint16(anon_sym_Fixed),
	3341: uint16(553),
	3342: uint16(1),
	3343: uint16(anon_sym_LBRACE),
	3344: uint16(45),
	3345: uint16(1),
	3346: uint16(sym_type),
	3347: uint16(547),
	3348: uint16(3),
	3349: uint16(anon_sym_UInt),
	3350: uint16(anon_sym_SInt),
	3351: uint16(anon_sym_Analog),
	3352: uint16(551),
	3353: uint16(3),
	3354: uint16(anon_sym_Clock),
	3355: uint16(anon_sym_AsyncReset),
	3356: uint16(anon_sym_Reset),
	3357: uint16(8),
	3358: uint16(3),
	3359: uint16(1),
	3360: uint16(sym_comment),
	3361: uint16(559),
	3362: uint16(1),
	3363: uint16(anon_sym_flip),
	3364: uint16(567),
	3365: uint16(1),
	3366: uint16(anon_sym_RBRACE),
	3367: uint16(152),
	3368: uint16(1),
	3369: uint16(sym_uint),
	3370: uint16(274),
	3371: uint16(1),
	3372: uint16(sym_field),
	3373: uint16(347),
	3374: uint16(1),
	3375: uint16(sym_field_id),
	3376: uint16(555),
	3377: uint16(2),
	3378: uint16(sym_identifier),
	3379: uint16(sym_relaxed_identifier),
	3380: uint16(561),
	3381: uint16(2),
	3382: uint16(anon_sym_0),
	3383: uint16(aux_sym_uint_token1),
	3384: uint16(6),
	3385: uint16(3),
	3386: uint16(1),
	3387: uint16(sym_comment),
	3388: uint16(549),
	3389: uint16(1),
	3390: uint16(anon_sym_Fixed),
	3391: uint16(553),
	3392: uint16(1),
	3393: uint16(anon_sym_LBRACE),
	3394: uint16(46),
	3395: uint16(1),
	3396: uint16(sym_type),
	3397: uint16(547),
	3398: uint16(3),
	3399: uint16(anon_sym_UInt),
	3400: uint16(anon_sym_SInt),
	3401: uint16(anon_sym_Analog),
	3402: uint16(551),
	3403: uint16(3),
	3404: uint16(anon_sym_Clock),
	3405: uint16(anon_sym_AsyncReset),
	3406: uint16(anon_sym_Reset),
	3407: uint16(6),
	3408: uint16(3),
	3409: uint16(1),
	3410: uint16(sym_comment),
	3411: uint16(541),
	3412: uint16(1),
	3413: uint16(anon_sym_Fixed),
	3414: uint16(545),
	3415: uint16(1),
	3416: uint16(anon_sym_LBRACE),
	3417: uint16(287),
	3418: uint16(1),
	3419: uint16(sym_type),
	3420: uint16(539),
	3421: uint16(3),
	3422: uint16(anon_sym_UInt),
	3423: uint16(anon_sym_SInt),
	3424: uint16(anon_sym_Analog),
	3425: uint16(543),
	3426: uint16(3),
	3427: uint16(anon_sym_Clock),
	3428: uint16(anon_sym_AsyncReset),
	3429: uint16(anon_sym_Reset),
	3430: uint16(2),
	3431: uint16(3),
	3432: uint16(1),
	3433: uint16(sym_comment),
	3434: uint16(203),
	3435: uint16(9),
	3436: uint16(anon_sym_COLON),
	3437: uint16(anon_sym_COMMA),
	3438: uint16(anon_sym_LBRACK),
	3439: uint16(anon_sym_RBRACK),
	3440: uint16(anon_sym_RPAREN),
	3441: uint16(anon_sym_LT_EQ),
	3442: uint16(anon_sym_LT_DASH),
	3443: uint16(anon_sym_is),
	3444: uint16(anon_sym_DOT),
	3445: uint16(6),
	3446: uint16(3),
	3447: uint16(1),
	3448: uint16(sym_comment),
	3449: uint16(541),
	3450: uint16(1),
	3451: uint16(anon_sym_Fixed),
	3452: uint16(545),
	3453: uint16(1),
	3454: uint16(anon_sym_LBRACE),
	3455: uint16(316),
	3456: uint16(1),
	3457: uint16(sym_type),
	3458: uint16(539),
	3459: uint16(3),
	3460: uint16(anon_sym_UInt),
	3461: uint16(anon_sym_SInt),
	3462: uint16(anon_sym_Analog),
	3463: uint16(543),
	3464: uint16(3),
	3465: uint16(anon_sym_Clock),
	3466: uint16(anon_sym_AsyncReset),
	3467: uint16(anon_sym_Reset),
	3468: uint16(2),
	3469: uint16(3),
	3470: uint16(1),
	3471: uint16(sym_comment),
	3472: uint16(195),
	3473: uint16(9),
	3474: uint16(anon_sym_COLON),
	3475: uint16(anon_sym_COMMA),
	3476: uint16(anon_sym_LBRACK),
	3477: uint16(anon_sym_RBRACK),
	3478: uint16(anon_sym_RPAREN),
	3479: uint16(anon_sym_LT_EQ),
	3480: uint16(anon_sym_LT_DASH),
	3481: uint16(anon_sym_is),
	3482: uint16(anon_sym_DOT),
	3483: uint16(2),
	3484: uint16(3),
	3485: uint16(1),
	3486: uint16(sym_comment),
	3487: uint16(171),
	3488: uint16(9),
	3489: uint16(anon_sym_COLON),
	3490: uint16(anon_sym_COMMA),
	3491: uint16(anon_sym_LBRACK),
	3492: uint16(anon_sym_RBRACK),
	3493: uint16(anon_sym_RPAREN),
	3494: uint16(anon_sym_LT_EQ),
	3495: uint16(anon_sym_LT_DASH),
	3496: uint16(anon_sym_is),
	3497: uint16(anon_sym_DOT),
	3498: uint16(6),
	3499: uint16(3),
	3500: uint16(1),
	3501: uint16(sym_comment),
	3502: uint16(511),
	3503: uint16(1),
	3504: uint16(anon_sym_Fixed),
	3505: uint16(515),
	3506: uint16(1),
	3507: uint16(anon_sym_LBRACE),
	3508: uint16(263),
	3509: uint16(1),
	3510: uint16(sym_type),
	3511: uint16(509),
	3512: uint16(3),
	3513: uint16(anon_sym_UInt),
	3514: uint16(anon_sym_SInt),
	3515: uint16(anon_sym_Analog),
	3516: uint16(513),
	3517: uint16(3),
	3518: uint16(anon_sym_Clock),
	3519: uint16(anon_sym_AsyncReset),
	3520: uint16(anon_sym_Reset),
	3521: uint16(2),
	3522: uint16(3),
	3523: uint16(1),
	3524: uint16(sym_comment),
	3525: uint16(155),
	3526: uint16(9),
	3527: uint16(anon_sym_COLON),
	3528: uint16(anon_sym_COMMA),
	3529: uint16(anon_sym_LBRACK),
	3530: uint16(anon_sym_RBRACK),
	3531: uint16(anon_sym_RPAREN),
	3532: uint16(anon_sym_LT_EQ),
	3533: uint16(anon_sym_LT_DASH),
	3534: uint16(anon_sym_is),
	3535: uint16(anon_sym_DOT),
	3536: uint16(2),
	3537: uint16(3),
	3538: uint16(1),
	3539: uint16(sym_comment),
	3540: uint16(179),
	3541: uint16(9),
	3542: uint16(anon_sym_COLON),
	3543: uint16(anon_sym_COMMA),
	3544: uint16(anon_sym_LBRACK),
	3545: uint16(anon_sym_RBRACK),
	3546: uint16(anon_sym_RPAREN),
	3547: uint16(anon_sym_LT_EQ),
	3548: uint16(anon_sym_LT_DASH),
	3549: uint16(anon_sym_is),
	3550: uint16(anon_sym_DOT),
	3551: uint16(2),
	3552: uint16(3),
	3553: uint16(1),
	3554: uint16(sym_comment),
	3555: uint16(199),
	3556: uint16(9),
	3557: uint16(anon_sym_COLON),
	3558: uint16(anon_sym_COMMA),
	3559: uint16(anon_sym_LBRACK),
	3560: uint16(anon_sym_RBRACK),
	3561: uint16(anon_sym_RPAREN),
	3562: uint16(anon_sym_LT_EQ),
	3563: uint16(anon_sym_LT_DASH),
	3564: uint16(anon_sym_is),
	3565: uint16(anon_sym_DOT),
	3566: uint16(7),
	3567: uint16(3),
	3568: uint16(1),
	3569: uint16(sym_comment),
	3570: uint16(559),
	3571: uint16(1),
	3572: uint16(anon_sym_flip),
	3573: uint16(152),
	3574: uint16(1),
	3575: uint16(sym_uint),
	3576: uint16(312),
	3577: uint16(1),
	3578: uint16(sym_field),
	3579: uint16(347),
	3580: uint16(1),
	3581: uint16(sym_field_id),
	3582: uint16(555),
	3583: uint16(2),
	3584: uint16(sym_identifier),
	3585: uint16(sym_relaxed_identifier),
	3586: uint16(561),
	3587: uint16(2),
	3588: uint16(anon_sym_0),
	3589: uint16(aux_sym_uint_token1),
	3590: uint16(6),
	3591: uint16(3),
	3592: uint16(1),
	3593: uint16(sym_comment),
	3594: uint16(569),
	3595: uint16(1),
	3596: uint16(sym_number_str),
	3597: uint16(421),
	3598: uint16(1),
	3599: uint16(sym_number),
	3600: uint16(467),
	3601: uint16(2),
	3602: uint16(anon_sym_0),
	3603: uint16(aux_sym_uint_token1),
	3604: uint16(469),
	3605: uint16(2),
	3606: uint16(anon_sym_PLUS),
	3607: uint16(anon_sym_DASH),
	3608: uint16(222),
	3609: uint16(2),
	3610: uint16(sym_uint),
	3611: uint16(sym_sint),
	3612: uint16(6),
	3613: uint16(3),
	3614: uint16(1),
	3615: uint16(sym_comment),
	3616: uint16(571),
	3617: uint16(1),
	3618: uint16(sym_number_str),
	3619: uint16(380),
	3620: uint16(1),
	3621: uint16(sym_number),
	3622: uint16(467),
	3623: uint16(2),
	3624: uint16(anon_sym_0),
	3625: uint16(aux_sym_uint_token1),
	3626: uint16(469),
	3627: uint16(2),
	3628: uint16(anon_sym_PLUS),
	3629: uint16(anon_sym_DASH),
	3630: uint16(222),
	3631: uint16(2),
	3632: uint16(sym_uint),
	3633: uint16(sym_sint),
	3634: uint16(2),
	3635: uint16(3),
	3636: uint16(1),
	3637: uint16(sym_comment),
	3638: uint16(573),
	3639: uint16(8),
	3640: uint16(anon_sym_UInt),
	3641: uint16(anon_sym_SInt),
	3642: uint16(anon_sym_Analog),
	3643: uint16(anon_sym_Fixed),
	3644: uint16(anon_sym_Clock),
	3645: uint16(anon_sym_AsyncReset),
	3646: uint16(anon_sym_Reset),
	3647: uint16(anon_sym_LBRACE),
	3648: uint16(6),
	3649: uint16(3),
	3650: uint16(1),
	3651: uint16(sym_comment),
	3652: uint16(575),
	3653: uint16(1),
	3654: uint16(anon_sym_LT),
	3655: uint16(403),
	3656: uint16(1),
	3657: uint16(sym_number),
	3658: uint16(467),
	3659: uint16(2),
	3660: uint16(anon_sym_0),
	3661: uint16(aux_sym_uint_token1),
	3662: uint16(469),
	3663: uint16(2),
	3664: uint16(anon_sym_PLUS),
	3665: uint16(anon_sym_DASH),
	3666: uint16(222),
	3667: uint16(2),
	3668: uint16(sym_uint),
	3669: uint16(sym_sint),
	3670: uint16(6),
	3671: uint16(3),
	3672: uint16(1),
	3673: uint16(sym_comment),
	3674: uint16(577),
	3675: uint16(1),
	3676: uint16(anon_sym_LT),
	3677: uint16(392),
	3678: uint16(1),
	3679: uint16(sym_number),
	3680: uint16(467),
	3681: uint16(2),
	3682: uint16(anon_sym_0),
	3683: uint16(aux_sym_uint_token1),
	3684: uint16(469),
	3685: uint16(2),
	3686: uint16(anon_sym_PLUS),
	3687: uint16(anon_sym_DASH),
	3688: uint16(222),
	3689: uint16(2),
	3690: uint16(sym_uint),
	3691: uint16(sym_sint),
	3692: uint16(6),
	3693: uint16(3),
	3694: uint16(1),
	3695: uint16(sym_comment),
	3696: uint16(579),
	3697: uint16(1),
	3698: uint16(anon_sym_LT),
	3699: uint16(348),
	3700: uint16(1),
	3701: uint16(sym_number),
	3702: uint16(467),
	3703: uint16(2),
	3704: uint16(anon_sym_0),
	3705: uint16(aux_sym_uint_token1),
	3706: uint16(469),
	3707: uint16(2),
	3708: uint16(anon_sym_PLUS),
	3709: uint16(anon_sym_DASH),
	3710: uint16(222),
	3711: uint16(2),
	3712: uint16(sym_uint),
	3713: uint16(sym_sint),
	3714: uint16(5),
	3715: uint16(3),
	3716: uint16(1),
	3717: uint16(sym_comment),
	3718: uint16(358),
	3719: uint16(1),
	3720: uint16(sym_number),
	3721: uint16(467),
	3722: uint16(2),
	3723: uint16(anon_sym_0),
	3724: uint16(aux_sym_uint_token1),
	3725: uint16(469),
	3726: uint16(2),
	3727: uint16(anon_sym_PLUS),
	3728: uint16(anon_sym_DASH),
	3729: uint16(222),
	3730: uint16(2),
	3731: uint16(sym_uint),
	3732: uint16(sym_sint),
	3733: uint16(5),
	3734: uint16(3),
	3735: uint16(1),
	3736: uint16(sym_comment),
	3737: uint16(417),
	3738: uint16(1),
	3739: uint16(sym_number),
	3740: uint16(467),
	3741: uint16(2),
	3742: uint16(anon_sym_0),
	3743: uint16(aux_sym_uint_token1),
	3744: uint16(469),
	3745: uint16(2),
	3746: uint16(anon_sym_PLUS),
	3747: uint16(anon_sym_DASH),
	3748: uint16(222),
	3749: uint16(2),
	3750: uint16(sym_uint),
	3751: uint16(sym_sint),
	3752: uint16(5),
	3753: uint16(3),
	3754: uint16(1),
	3755: uint16(sym_comment),
	3756: uint16(416),
	3757: uint16(1),
	3758: uint16(sym_number),
	3759: uint16(467),
	3760: uint16(2),
	3761: uint16(anon_sym_0),
	3762: uint16(aux_sym_uint_token1),
	3763: uint16(469),
	3764: uint16(2),
	3765: uint16(anon_sym_PLUS),
	3766: uint16(anon_sym_DASH),
	3767: uint16(222),
	3768: uint16(2),
	3769: uint16(sym_uint),
	3770: uint16(sym_sint),
	3771: uint16(5),
	3772: uint16(3),
	3773: uint16(1),
	3774: uint16(sym_comment),
	3775: uint16(424),
	3776: uint16(1),
	3777: uint16(sym_number),
	3778: uint16(467),
	3779: uint16(2),
	3780: uint16(anon_sym_0),
	3781: uint16(aux_sym_uint_token1),
	3782: uint16(469),
	3783: uint16(2),
	3784: uint16(anon_sym_PLUS),
	3785: uint16(anon_sym_DASH),
	3786: uint16(222),
	3787: uint16(2),
	3788: uint16(sym_uint),
	3789: uint16(sym_sint),
	3790: uint16(5),
	3791: uint16(3),
	3792: uint16(1),
	3793: uint16(sym_comment),
	3794: uint16(394),
	3795: uint16(1),
	3796: uint16(sym_number),
	3797: uint16(467),
	3798: uint16(2),
	3799: uint16(anon_sym_0),
	3800: uint16(aux_sym_uint_token1),
	3801: uint16(469),
	3802: uint16(2),
	3803: uint16(anon_sym_PLUS),
	3804: uint16(anon_sym_DASH),
	3805: uint16(222),
	3806: uint16(2),
	3807: uint16(sym_uint),
	3808: uint16(sym_sint),
	3809: uint16(5),
	3810: uint16(3),
	3811: uint16(1),
	3812: uint16(sym_comment),
	3813: uint16(400),
	3814: uint16(1),
	3815: uint16(sym_number),
	3816: uint16(467),
	3817: uint16(2),
	3818: uint16(anon_sym_0),
	3819: uint16(aux_sym_uint_token1),
	3820: uint16(469),
	3821: uint16(2),
	3822: uint16(anon_sym_PLUS),
	3823: uint16(anon_sym_DASH),
	3824: uint16(222),
	3825: uint16(2),
	3826: uint16(sym_uint),
	3827: uint16(sym_sint),
	3828: uint16(5),
	3829: uint16(3),
	3830: uint16(1),
	3831: uint16(sym_comment),
	3832: uint16(405),
	3833: uint16(1),
	3834: uint16(sym_number),
	3835: uint16(467),
	3836: uint16(2),
	3837: uint16(anon_sym_0),
	3838: uint16(aux_sym_uint_token1),
	3839: uint16(469),
	3840: uint16(2),
	3841: uint16(anon_sym_PLUS),
	3842: uint16(anon_sym_DASH),
	3843: uint16(222),
	3844: uint16(2),
	3845: uint16(sym_uint),
	3846: uint16(sym_sint),
	3847: uint16(5),
	3848: uint16(3),
	3849: uint16(1),
	3850: uint16(sym_comment),
	3851: uint16(431),
	3852: uint16(1),
	3853: uint16(sym_number),
	3854: uint16(467),
	3855: uint16(2),
	3856: uint16(anon_sym_0),
	3857: uint16(aux_sym_uint_token1),
	3858: uint16(469),
	3859: uint16(2),
	3860: uint16(anon_sym_PLUS),
	3861: uint16(anon_sym_DASH),
	3862: uint16(222),
	3863: uint16(2),
	3864: uint16(sym_uint),
	3865: uint16(sym_sint),
	3866: uint16(5),
	3867: uint16(3),
	3868: uint16(1),
	3869: uint16(sym_comment),
	3870: uint16(423),
	3871: uint16(1),
	3872: uint16(sym_number),
	3873: uint16(467),
	3874: uint16(2),
	3875: uint16(anon_sym_0),
	3876: uint16(aux_sym_uint_token1),
	3877: uint16(469),
	3878: uint16(2),
	3879: uint16(anon_sym_PLUS),
	3880: uint16(anon_sym_DASH),
	3881: uint16(222),
	3882: uint16(2),
	3883: uint16(sym_uint),
	3884: uint16(sym_sint),
	3885: uint16(5),
	3886: uint16(3),
	3887: uint16(1),
	3888: uint16(sym_comment),
	3889: uint16(441),
	3890: uint16(1),
	3891: uint16(sym_number),
	3892: uint16(467),
	3893: uint16(2),
	3894: uint16(anon_sym_0),
	3895: uint16(aux_sym_uint_token1),
	3896: uint16(469),
	3897: uint16(2),
	3898: uint16(anon_sym_PLUS),
	3899: uint16(anon_sym_DASH),
	3900: uint16(222),
	3901: uint16(2),
	3902: uint16(sym_uint),
	3903: uint16(sym_sint),
	3904: uint16(5),
	3905: uint16(3),
	3906: uint16(1),
	3907: uint16(sym_comment),
	3908: uint16(384),
	3909: uint16(1),
	3910: uint16(sym_number),
	3911: uint16(467),
	3912: uint16(2),
	3913: uint16(anon_sym_0),
	3914: uint16(aux_sym_uint_token1),
	3915: uint16(469),
	3916: uint16(2),
	3917: uint16(anon_sym_PLUS),
	3918: uint16(anon_sym_DASH),
	3919: uint16(222),
	3920: uint16(2),
	3921: uint16(sym_uint),
	3922: uint16(sym_sint),
	3923: uint16(5),
	3924: uint16(3),
	3925: uint16(1),
	3926: uint16(sym_comment),
	3927: uint16(402),
	3928: uint16(1),
	3929: uint16(sym_number),
	3930: uint16(467),
	3931: uint16(2),
	3932: uint16(anon_sym_0),
	3933: uint16(aux_sym_uint_token1),
	3934: uint16(469),
	3935: uint16(2),
	3936: uint16(anon_sym_PLUS),
	3937: uint16(anon_sym_DASH),
	3938: uint16(222),
	3939: uint16(2),
	3940: uint16(sym_uint),
	3941: uint16(sym_sint),
	3942: uint16(5),
	3943: uint16(3),
	3944: uint16(1),
	3945: uint16(sym_comment),
	3946: uint16(391),
	3947: uint16(1),
	3948: uint16(sym_number),
	3949: uint16(467),
	3950: uint16(2),
	3951: uint16(anon_sym_0),
	3952: uint16(aux_sym_uint_token1),
	3953: uint16(469),
	3954: uint16(2),
	3955: uint16(anon_sym_PLUS),
	3956: uint16(anon_sym_DASH),
	3957: uint16(222),
	3958: uint16(2),
	3959: uint16(sym_uint),
	3960: uint16(sym_sint),
	3961: uint16(5),
	3962: uint16(3),
	3963: uint16(1),
	3964: uint16(sym_comment),
	3965: uint16(439),
	3966: uint16(1),
	3967: uint16(sym_number),
	3968: uint16(467),
	3969: uint16(2),
	3970: uint16(anon_sym_0),
	3971: uint16(aux_sym_uint_token1),
	3972: uint16(469),
	3973: uint16(2),
	3974: uint16(anon_sym_PLUS),
	3975: uint16(anon_sym_DASH),
	3976: uint16(222),
	3977: uint16(2),
	3978: uint16(sym_uint),
	3979: uint16(sym_sint),
	3980: uint16(5),
	3981: uint16(3),
	3982: uint16(1),
	3983: uint16(sym_comment),
	3984: uint16(389),
	3985: uint16(1),
	3986: uint16(sym_number),
	3987: uint16(581),
	3988: uint16(2),
	3989: uint16(anon_sym_0),
	3990: uint16(aux_sym_uint_token1),
	3991: uint16(583),
	3992: uint16(2),
	3993: uint16(anon_sym_PLUS),
	3994: uint16(anon_sym_DASH),
	3995: uint16(343),
	3996: uint16(2),
	3997: uint16(sym_uint),
	3998: uint16(sym_sint),
	3999: uint16(4),
	4000: uint16(585),
	4001: uint16(1),
	4002: uint16(anon_sym_SQUOTE),
	4003: uint16(590),
	4004: uint16(1),
	4005: uint16(sym_comment),
	4006: uint16(190),
	4007: uint16(2),
	4008: uint16(sym__escape_sequence),
	4009: uint16(aux_sym_raw_string_repeat1),
	4010: uint16(587),
	4011: uint16(3),
	4012: uint16(sym_raw_string_content),
	4013: uint16(aux_sym__escape_sequence_token1),
	4014: uint16(sym_escape_sequence),
	4015: uint16(4),
	4016: uint16(590),
	4017: uint16(1),
	4018: uint16(sym_comment),
	4019: uint16(592),
	4020: uint16(1),
	4021: uint16(anon_sym_SQUOTE),
	4022: uint16(190),
	4023: uint16(2),
	4024: uint16(sym__escape_sequence),
	4025: uint16(aux_sym_raw_string_repeat1),
	4026: uint16(594),
	4027: uint16(3),
	4028: uint16(sym_raw_string_content),
	4029: uint16(aux_sym__escape_sequence_token1),
	4030: uint16(sym_escape_sequence),
	4031: uint16(4),
	4032: uint16(590),
	4033: uint16(1),
	4034: uint16(sym_comment),
	4035: uint16(596),
	4036: uint16(1),
	4037: uint16(anon_sym_DQUOTE),
	4038: uint16(192),
	4039: uint16(2),
	4040: uint16(sym__escape_sequence),
	4041: uint16(aux_sym_string_repeat1),
	4042: uint16(598),
	4043: uint16(3),
	4044: uint16(sym_string_content),
	4045: uint16(aux_sym__escape_sequence_token1),
	4046: uint16(sym_escape_sequence),
	4047: uint16(4),
	4048: uint16(590),
	4049: uint16(1),
	4050: uint16(sym_comment),
	4051: uint16(601),
	4052: uint16(1),
	4053: uint16(anon_sym_SQUOTE),
	4054: uint16(191),
	4055: uint16(2),
	4056: uint16(sym__escape_sequence),
	4057: uint16(aux_sym_raw_string_repeat1),
	4058: uint16(603),
	4059: uint16(3),
	4060: uint16(sym_raw_string_content),
	4061: uint16(aux_sym__escape_sequence_token1),
	4062: uint16(sym_escape_sequence),
	4063: uint16(4),
	4064: uint16(590),
	4065: uint16(1),
	4066: uint16(sym_comment),
	4067: uint16(605),
	4068: uint16(1),
	4069: uint16(anon_sym_DQUOTE),
	4070: uint16(192),
	4071: uint16(2),
	4072: uint16(sym__escape_sequence),
	4073: uint16(aux_sym_string_repeat1),
	4074: uint16(607),
	4075: uint16(3),
	4076: uint16(sym_string_content),
	4077: uint16(aux_sym__escape_sequence_token1),
	4078: uint16(sym_escape_sequence),
	4079: uint16(4),
	4080: uint16(590),
	4081: uint16(1),
	4082: uint16(sym_comment),
	4083: uint16(609),
	4084: uint16(1),
	4085: uint16(anon_sym_DQUOTE),
	4086: uint16(194),
	4087: uint16(2),
	4088: uint16(sym__escape_sequence),
	4089: uint16(aux_sym_string_repeat1),
	4090: uint16(611),
	4091: uint16(3),
	4092: uint16(sym_string_content),
	4093: uint16(aux_sym__escape_sequence_token1),
	4094: uint16(sym_escape_sequence),
	4095: uint16(6),
	4096: uint16(3),
	4097: uint16(1),
	4098: uint16(sym_comment),
	4099: uint16(613),
	4100: uint16(1),
	4101: uint16(sym_identifier),
	4102: uint16(617),
	4103: uint16(1),
	4104: uint16(sym_relaxed_identifier),
	4105: uint16(24),
	4106: uint16(1),
	4107: uint16(sym_uint),
	4108: uint16(26),
	4109: uint16(1),
	4110: uint16(sym_field_id),
	4111: uint16(615),
	4112: uint16(2),
	4113: uint16(anon_sym_0),
	4114: uint16(aux_sym_uint_token1),
	4115: uint16(6),
	4116: uint16(3),
	4117: uint16(1),
	4118: uint16(sym_comment),
	4119: uint16(555),
	4120: uint16(1),
	4121: uint16(sym_relaxed_identifier),
	4122: uint16(619),
	4123: uint16(1),
	4124: uint16(sym_identifier),
	4125: uint16(147),
	4126: uint16(1),
	4127: uint16(sym_field_id),
	4128: uint16(152),
	4129: uint16(1),
	4130: uint16(sym_uint),
	4131: uint16(561),
	4132: uint16(2),
	4133: uint16(anon_sym_0),
	4134: uint16(aux_sym_uint_token1),
	4135: uint16(6),
	4136: uint16(3),
	4137: uint16(1),
	4138: uint16(sym_comment),
	4139: uint16(555),
	4140: uint16(1),
	4141: uint16(sym_relaxed_identifier),
	4142: uint16(619),
	4143: uint16(1),
	4144: uint16(sym_identifier),
	4145: uint16(152),
	4146: uint16(1),
	4147: uint16(sym_uint),
	4148: uint16(341),
	4149: uint16(1),
	4150: uint16(sym_field_id),
	4151: uint16(561),
	4152: uint16(2),
	4153: uint16(anon_sym_0),
	4154: uint16(aux_sym_uint_token1),
	4155: uint16(4),
	4156: uint16(590),
	4157: uint16(1),
	4158: uint16(sym_comment),
	4159: uint16(621),
	4160: uint16(1),
	4161: uint16(anon_sym_DQUOTE),
	4162: uint16(200),
	4163: uint16(2),
	4164: uint16(sym__escape_sequence),
	4165: uint16(aux_sym_string_repeat1),
	4166: uint16(623),
	4167: uint16(3),
	4168: uint16(sym_string_content),
	4169: uint16(aux_sym__escape_sequence_token1),
	4170: uint16(sym_escape_sequence),
	4171: uint16(4),
	4172: uint16(590),
	4173: uint16(1),
	4174: uint16(sym_comment),
	4175: uint16(625),
	4176: uint16(1),
	4177: uint16(anon_sym_DQUOTE),
	4178: uint16(192),
	4179: uint16(2),
	4180: uint16(sym__escape_sequence),
	4181: uint16(aux_sym_string_repeat1),
	4182: uint16(607),
	4183: uint16(3),
	4184: uint16(sym_string_content),
	4185: uint16(aux_sym__escape_sequence_token1),
	4186: uint16(sym_escape_sequence),
	4187: uint16(5),
	4188: uint16(3),
	4189: uint16(1),
	4190: uint16(sym_comment),
	4191: uint16(627),
	4192: uint16(1),
	4193: uint16(anon_sym_module),
	4194: uint16(629),
	4195: uint16(1),
	4196: uint16(anon_sym_extmodule),
	4197: uint16(631),
	4198: uint16(1),
	4199: uint16(sym__dedent),
	4200: uint16(202),
	4201: uint16(2),
	4202: uint16(sym_module),
	4203: uint16(aux_sym_circuit_repeat1),
	4204: uint16(5),
	4205: uint16(3),
	4206: uint16(1),
	4207: uint16(sym_comment),
	4208: uint16(633),
	4209: uint16(1),
	4210: uint16(anon_sym_module),
	4211: uint16(636),
	4212: uint16(1),
	4213: uint16(anon_sym_extmodule),
	4214: uint16(639),
	4215: uint16(1),
	4216: uint16(sym__dedent),
	4217: uint16(202),
	4218: uint16(2),
	4219: uint16(sym_module),
	4220: uint16(aux_sym_circuit_repeat1),
	4221: uint16(5),
	4222: uint16(3),
	4223: uint16(1),
	4224: uint16(sym_comment),
	4225: uint16(627),
	4226: uint16(1),
	4227: uint16(anon_sym_module),
	4228: uint16(629),
	4229: uint16(1),
	4230: uint16(anon_sym_extmodule),
	4231: uint16(631),
	4232: uint16(1),
	4233: uint16(sym__dedent),
	4234: uint16(209),
	4235: uint16(2),
	4236: uint16(sym_module),
	4237: uint16(aux_sym_circuit_repeat1),
	4238: uint16(6),
	4239: uint16(3),
	4240: uint16(1),
	4241: uint16(sym_comment),
	4242: uint16(641),
	4243: uint16(1),
	4244: uint16(anon_sym_COMMA),
	4245: uint16(643),
	4246: uint16(1),
	4247: uint16(anon_sym_LBRACK),
	4248: uint16(645),
	4249: uint16(1),
	4250: uint16(anon_sym_RPAREN),
	4251: uint16(647),
	4252: uint16(1),
	4253: uint16(anon_sym_DOT),
	4254: uint16(275),
	4255: uint16(1),
	4256: uint16(aux_sym_primitive_operation_repeat1),
	4257: uint16(6),
	4258: uint16(3),
	4259: uint16(1),
	4260: uint16(sym_comment),
	4261: uint16(643),
	4262: uint16(1),
	4263: uint16(anon_sym_LBRACK),
	4264: uint16(647),
	4265: uint16(1),
	4266: uint16(anon_sym_DOT),
	4267: uint16(649),
	4268: uint16(1),
	4269: uint16(anon_sym_LT_EQ),
	4270: uint16(651),
	4271: uint16(1),
	4272: uint16(anon_sym_LT_DASH),
	4273: uint16(653),
	4274: uint16(1),
	4275: uint16(anon_sym_is),
	4276: uint16(4),
	4277: uint16(3),
	4278: uint16(1),
	4279: uint16(sym_comment),
	4280: uint16(657),
	4281: uint16(1),
	4282: uint16(sym_info),
	4283: uint16(659),
	4284: uint16(1),
	4285: uint16(sym__indent),
	4286: uint16(655),
	4287: uint16(3),
	4288: uint16(sym__dedent),
	4289: uint16(anon_sym_module),
	4290: uint16(anon_sym_extmodule),
	4291: uint16(4),
	4292: uint16(3),
	4293: uint16(1),
	4294: uint16(sym_comment),
	4295: uint16(661),
	4296: uint16(1),
	4297: uint16(sym_info),
	4298: uint16(663),
	4299: uint16(1),
	4300: uint16(sym__indent),
	4301: uint16(655),
	4302: uint16(3),
	4303: uint16(sym__dedent),
	4304: uint16(anon_sym_module),
	4305: uint16(anon_sym_extmodule),
	4306: uint16(6),
	4307: uint16(3),
	4308: uint16(1),
	4309: uint16(sym_comment),
	4310: uint16(643),
	4311: uint16(1),
	4312: uint16(anon_sym_LBRACK),
	4313: uint16(647),
	4314: uint16(1),
	4315: uint16(anon_sym_DOT),
	4316: uint16(665),
	4317: uint16(1),
	4318: uint16(anon_sym_COMMA),
	4319: uint16(667),
	4320: uint16(1),
	4321: uint16(anon_sym_RPAREN),
	4322: uint16(286),
	4323: uint16(1),
	4324: uint16(aux_sym_printf_repeat1),
	4325: uint16(5),
	4326: uint16(3),
	4327: uint16(1),
	4328: uint16(sym_comment),
	4329: uint16(627),
	4330: uint16(1),
	4331: uint16(anon_sym_module),
	4332: uint16(629),
	4333: uint16(1),
	4334: uint16(anon_sym_extmodule),
	4335: uint16(669),
	4336: uint16(1),
	4337: uint16(sym__dedent),
	4338: uint16(202),
	4339: uint16(2),
	4340: uint16(sym_module),
	4341: uint16(aux_sym_circuit_repeat1),
	4342: uint16(6),
	4343: uint16(3),
	4344: uint16(1),
	4345: uint16(sym_comment),
	4346: uint16(641),
	4347: uint16(1),
	4348: uint16(anon_sym_COMMA),
	4349: uint16(643),
	4350: uint16(1),
	4351: uint16(anon_sym_LBRACK),
	4352: uint16(647),
	4353: uint16(1),
	4354: uint16(anon_sym_DOT),
	4355: uint16(671),
	4356: uint16(1),
	4357: uint16(anon_sym_RPAREN),
	4358: uint16(303),
	4359: uint16(1),
	4360: uint16(aux_sym_primitive_operation_repeat1),
	4361: uint16(5),
	4362: uint16(3),
	4363: uint16(1),
	4364: uint16(sym_comment),
	4365: uint16(627),
	4366: uint16(1),
	4367: uint16(anon_sym_module),
	4368: uint16(629),
	4369: uint16(1),
	4370: uint16(anon_sym_extmodule),
	4371: uint16(673),
	4372: uint16(1),
	4373: uint16(sym__dedent),
	4374: uint16(201),
	4375: uint16(2),
	4376: uint16(sym_module),
	4377: uint16(aux_sym_circuit_repeat1),
	4378: uint16(3),
	4379: uint16(3),
	4380: uint16(1),
	4381: uint16(sym_comment),
	4382: uint16(389),
	4383: uint16(1),
	4384: uint16(sym_ruw),
	4385: uint16(675),
	4386: uint16(3),
	4387: uint16(anon_sym_old),
	4388: uint16(anon_sym_new),
	4389: uint16(anon_sym_undefined),
	4390: uint16(3),
	4391: uint16(3),
	4392: uint16(1),
	4393: uint16(sym_comment),
	4394: uint16(677),
	4395: uint16(1),
	4396: uint16(anon_sym_LT),
	4397: uint16(219),
	4398: uint16(3),
	4399: uint16(sym__newline),
	4400: uint16(anon_sym_LBRACK),
	4401: uint16(sym_info),
	4402: uint16(4),
	4403: uint16(3),
	4404: uint16(1),
	4405: uint16(sym_comment),
	4406: uint16(679),
	4407: uint16(1),
	4408: uint16(anon_sym_parameter),
	4409: uint16(682),
	4410: uint16(1),
	4411: uint16(sym__dedent),
	4412: uint16(214),
	4413: uint16(2),
	4414: uint16(sym_parameter),
	4415: uint16(aux_sym_module_repeat3),
	4416: uint16(4),
	4417: uint16(3),
	4418: uint16(1),
	4419: uint16(sym_comment),
	4420: uint16(126),
	4421: uint16(1),
	4422: uint16(sym__dedent),
	4423: uint16(493),
	4424: uint16(1),
	4425: uint16(anon_sym_parameter),
	4426: uint16(229),
	4427: uint16(2),
	4428: uint16(sym_parameter),
	4429: uint16(aux_sym_module_repeat3),
	4430: uint16(3),
	4431: uint16(3),
	4432: uint16(1),
	4433: uint16(sym_comment),
	4434: uint16(686),
	4435: uint16(1),
	4436: uint16(sym__indent),
	4437: uint16(684),
	4438: uint16(3),
	4439: uint16(sym__dedent),
	4440: uint16(anon_sym_module),
	4441: uint16(anon_sym_extmodule),
	4442: uint16(4),
	4443: uint16(3),
	4444: uint16(1),
	4445: uint16(sym_comment),
	4446: uint16(643),
	4447: uint16(1),
	4448: uint16(anon_sym_LBRACK),
	4449: uint16(647),
	4450: uint16(1),
	4451: uint16(anon_sym_DOT),
	4452: uint16(688),
	4453: uint16(2),
	4454: uint16(anon_sym_COMMA),
	4455: uint16(anon_sym_RPAREN),
	4456: uint16(4),
	4457: uint16(3),
	4458: uint16(1),
	4459: uint16(sym_comment),
	4460: uint16(692),
	4461: uint16(1),
	4462: uint16(sym_info),
	4463: uint16(694),
	4464: uint16(1),
	4465: uint16(sym__indent),
	4466: uint16(690),
	4467: uint16(2),
	4469: uint16(anon_sym_circuit),
	4470: uint16(3),
	4471: uint16(3),
	4472: uint16(1),
	4473: uint16(sym_comment),
	4474: uint16(696),
	4475: uint16(1),
	4476: uint16(anon_sym_LT),
	4477: uint16(249),
	4478: uint16(3),
	4479: uint16(anon_sym_COMMA),
	4480: uint16(anon_sym_RBRACE),
	4481: uint16(anon_sym_LBRACK),
	4482: uint16(4),
	4483: uint16(3),
	4484: uint16(1),
	4485: uint16(sym_comment),
	4486: uint16(643),
	4487: uint16(1),
	4488: uint16(anon_sym_LBRACK),
	4489: uint16(647),
	4490: uint16(1),
	4491: uint16(anon_sym_DOT),
	4492: uint16(698),
	4493: uint16(2),
	4494: uint16(anon_sym_COMMA),
	4495: uint16(anon_sym_RPAREN),
	4496: uint16(4),
	4497: uint16(3),
	4498: uint16(1),
	4499: uint16(sym_comment),
	4500: uint16(126),
	4501: uint16(1),
	4502: uint16(sym__dedent),
	4503: uint16(493),
	4504: uint16(1),
	4505: uint16(anon_sym_parameter),
	4506: uint16(214),
	4507: uint16(2),
	4508: uint16(sym_parameter),
	4509: uint16(aux_sym_module_repeat3),
	4510: uint16(2),
	4511: uint16(3),
	4512: uint16(1),
	4513: uint16(sym_comment),
	4514: uint16(700),
	4515: uint16(4),
	4516: uint16(anon_sym_GT),
	4517: uint16(anon_sym_COMMA),
	4518: uint16(anon_sym_RBRACK),
	4519: uint16(anon_sym_RPAREN),
	4520: uint16(4),
	4521: uint16(3),
	4522: uint16(1),
	4523: uint16(sym_comment),
	4524: uint16(702),
	4525: uint16(1),
	4527: uint16(704),
	4528: uint16(1),
	4529: uint16(anon_sym_circuit),
	4530: uint16(223),
	4531: uint16(2),
	4532: uint16(sym_circuit),
	4533: uint16(aux_sym_source_file_repeat1),
	4534: uint16(3),
	4535: uint16(3),
	4536: uint16(1),
	4537: uint16(sym_comment),
	4538: uint16(707),
	4539: uint16(1),
	4540: uint16(anon_sym_LT),
	4541: uint16(219),
	4542: uint16(3),
	4543: uint16(sym__newline),
	4544: uint16(anon_sym_LBRACK),
	4545: uint16(sym_info),
	4546: uint16(4),
	4547: uint16(3),
	4548: uint16(1),
	4549: uint16(sym_comment),
	4550: uint16(49),
	4551: uint16(1),
	4552: uint16(sym__dedent),
	4553: uint16(493),
	4554: uint16(1),
	4555: uint16(anon_sym_parameter),
	4556: uint16(236),
	4557: uint16(2),
	4558: uint16(sym_parameter),
	4559: uint16(aux_sym_module_repeat3),
	4560: uint16(5),
	4561: uint16(3),
	4562: uint16(1),
	4563: uint16(sym_comment),
	4564: uint16(709),
	4565: uint16(1),
	4566: uint16(anon_sym_reset),
	4567: uint16(711),
	4568: uint16(1),
	4569: uint16(anon_sym_LPAREN),
	4570: uint16(314),
	4571: uint16(1),
	4572: uint16(sym__reset),
	4573: uint16(315),
	4574: uint16(1),
	4575: uint16(sym_reset),
	4576: uint16(3),
	4577: uint16(3),
	4578: uint16(1),
	4579: uint16(sym_comment),
	4580: uint16(713),
	4581: uint16(1),
	4582: uint16(anon_sym_LT),
	4583: uint16(249),
	4584: uint16(3),
	4585: uint16(sym__newline),
	4586: uint16(anon_sym_LBRACK),
	4587: uint16(sym_info),
	4588: uint16(5),
	4589: uint16(3),
	4590: uint16(1),
	4591: uint16(sym_comment),
	4592: uint16(715),
	4593: uint16(1),
	4594: uint16(anon_sym_reset),
	4595: uint16(717),
	4596: uint16(1),
	4597: uint16(anon_sym_LPAREN),
	4598: uint16(355),
	4599: uint16(1),
	4600: uint16(sym_reset),
	4601: uint16(364),
	4602: uint16(1),
	4603: uint16(sym__reset),
	4604: uint16(4),
	4605: uint16(3),
	4606: uint16(1),
	4607: uint16(sym_comment),
	4608: uint16(493),
	4609: uint16(1),
	4610: uint16(anon_sym_parameter),
	4611: uint16(719),
	4612: uint16(1),
	4613: uint16(sym__dedent),
	4614: uint16(214),
	4615: uint16(2),
	4616: uint16(sym_parameter),
	4617: uint16(aux_sym_module_repeat3),
	4618: uint16(4),
	4619: uint16(3),
	4620: uint16(1),
	4621: uint16(sym_comment),
	4622: uint16(49),
	4623: uint16(1),
	4624: uint16(sym__dedent),
	4625: uint16(493),
	4626: uint16(1),
	4627: uint16(anon_sym_parameter),
	4628: uint16(214),
	4629: uint16(2),
	4630: uint16(sym_parameter),
	4631: uint16(aux_sym_module_repeat3),
	4632: uint16(3),
	4633: uint16(3),
	4634: uint16(1),
	4635: uint16(sym_comment),
	4636: uint16(721),
	4637: uint16(1),
	4638: uint16(anon_sym_LT),
	4639: uint16(219),
	4640: uint16(3),
	4641: uint16(anon_sym_COMMA),
	4642: uint16(anon_sym_RBRACE),
	4643: uint16(anon_sym_LBRACK),
	4644: uint16(3),
	4645: uint16(3),
	4646: uint16(1),
	4647: uint16(sym_comment),
	4648: uint16(723),
	4649: uint16(1),
	4650: uint16(anon_sym_LT),
	4651: uint16(219),
	4652: uint16(3),
	4653: uint16(anon_sym_COMMA),
	4654: uint16(anon_sym_RBRACE),
	4655: uint16(anon_sym_LBRACK),
	4656: uint16(3),
	4657: uint16(3),
	4658: uint16(1),
	4659: uint16(sym_comment),
	4660: uint16(725),
	4661: uint16(1),
	4662: uint16(sym__indent),
	4663: uint16(684),
	4664: uint16(3),
	4665: uint16(sym__dedent),
	4666: uint16(anon_sym_module),
	4667: uint16(anon_sym_extmodule),
	4668: uint16(4),
	4669: uint16(3),
	4670: uint16(1),
	4671: uint16(sym_comment),
	4672: uint16(7),
	4673: uint16(1),
	4674: uint16(anon_sym_circuit),
	4675: uint16(727),
	4676: uint16(1),
	4678: uint16(223),
	4679: uint16(2),
	4680: uint16(sym_circuit),
	4681: uint16(aux_sym_source_file_repeat1),
	4682: uint16(2),
	4683: uint16(3),
	4684: uint16(1),
	4685: uint16(sym_comment),
	4686: uint16(729),
	4687: uint16(4),
	4688: uint16(anon_sym_GT),
	4689: uint16(anon_sym_COMMA),
	4690: uint16(anon_sym_RBRACK),
	4691: uint16(anon_sym_RPAREN),
	4692: uint16(4),
	4693: uint16(3),
	4694: uint16(1),
	4695: uint16(sym_comment),
	4696: uint16(51),
	4697: uint16(1),
	4698: uint16(sym__dedent),
	4699: uint16(493),
	4700: uint16(1),
	4701: uint16(anon_sym_parameter),
	4702: uint16(214),
	4703: uint16(2),
	4704: uint16(sym_parameter),
	4705: uint16(aux_sym_module_repeat3),
	4706: uint16(4),
	4707: uint16(3),
	4708: uint16(1),
	4709: uint16(sym_comment),
	4710: uint16(51),
	4711: uint16(1),
	4712: uint16(sym__dedent),
	4713: uint16(493),
	4714: uint16(1),
	4715: uint16(anon_sym_parameter),
	4716: uint16(221),
	4717: uint16(2),
	4718: uint16(sym_parameter),
	4719: uint16(aux_sym_module_repeat3),
	4720: uint16(2),
	4721: uint16(3),
	4722: uint16(1),
	4723: uint16(sym_comment),
	4724: uint16(261),
	4725: uint16(3),
	4726: uint16(anon_sym_COMMA),
	4727: uint16(anon_sym_RBRACE),
	4728: uint16(anon_sym_LBRACK),
	4729: uint16(2),
	4730: uint16(3),
	4731: uint16(1),
	4732: uint16(sym_comment),
	4733: uint16(289),
	4734: uint16(3),
	4735: uint16(sym__newline),
	4736: uint16(anon_sym_LBRACK),
	4737: uint16(sym_info),
	4738: uint16(4),
	4739: uint16(3),
	4740: uint16(1),
	4741: uint16(sym_comment),
	4742: uint16(643),
	4743: uint16(1),
	4744: uint16(anon_sym_LBRACK),
	4745: uint16(647),
	4746: uint16(1),
	4747: uint16(anon_sym_DOT),
	4748: uint16(731),
	4749: uint16(1),
	4750: uint16(anon_sym_COMMA),
	4751: uint16(4),
	4752: uint16(3),
	4753: uint16(1),
	4754: uint16(sym_comment),
	4755: uint16(643),
	4756: uint16(1),
	4757: uint16(anon_sym_LBRACK),
	4758: uint16(647),
	4759: uint16(1),
	4760: uint16(anon_sym_DOT),
	4761: uint16(733),
	4762: uint16(1),
	4763: uint16(anon_sym_COMMA),
	4764: uint16(4),
	4765: uint16(3),
	4766: uint16(1),
	4767: uint16(sym_comment),
	4768: uint16(643),
	4769: uint16(1),
	4770: uint16(anon_sym_LBRACK),
	4771: uint16(647),
	4772: uint16(1),
	4773: uint16(anon_sym_DOT),
	4774: uint16(735),
	4775: uint16(1),
	4776: uint16(anon_sym_COMMA),
	4777: uint16(4),
	4778: uint16(3),
	4779: uint16(1),
	4780: uint16(sym_comment),
	4781: uint16(643),
	4782: uint16(1),
	4783: uint16(anon_sym_LBRACK),
	4784: uint16(647),
	4785: uint16(1),
	4786: uint16(anon_sym_DOT),
	4787: uint16(737),
	4788: uint16(1),
	4789: uint16(anon_sym_COMMA),
	4790: uint16(4),
	4791: uint16(3),
	4792: uint16(1),
	4793: uint16(sym_comment),
	4794: uint16(643),
	4795: uint16(1),
	4796: uint16(anon_sym_LBRACK),
	4797: uint16(647),
	4798: uint16(1),
	4799: uint16(anon_sym_DOT),
	4800: uint16(739),
	4801: uint16(1),
	4802: uint16(anon_sym_COMMA),
	4803: uint16(4),
	4804: uint16(3),
	4805: uint16(1),
	4806: uint16(sym_comment),
	4807: uint16(643),
	4808: uint16(1),
	4809: uint16(anon_sym_LBRACK),
	4810: uint16(647),
	4811: uint16(1),
	4812: uint16(anon_sym_DOT),
	4813: uint16(741),
	4814: uint16(1),
	4815: uint16(anon_sym_RPAREN),
	4816: uint16(4),
	4817: uint16(3),
	4818: uint16(1),
	4819: uint16(sym_comment),
	4820: uint16(643),
	4821: uint16(1),
	4822: uint16(anon_sym_LBRACK),
	4823: uint16(647),
	4824: uint16(1),
	4825: uint16(anon_sym_DOT),
	4826: uint16(743),
	4827: uint16(1),
	4828: uint16(anon_sym_COMMA),
	4829: uint16(4),
	4830: uint16(3),
	4831: uint16(1),
	4832: uint16(sym_comment),
	4833: uint16(643),
	4834: uint16(1),
	4835: uint16(anon_sym_LBRACK),
	4836: uint16(647),
	4837: uint16(1),
	4838: uint16(anon_sym_DOT),
	4839: uint16(745),
	4840: uint16(1),
	4841: uint16(anon_sym_COMMA),
	4842: uint16(4),
	4843: uint16(3),
	4844: uint16(1),
	4845: uint16(sym_comment),
	4846: uint16(643),
	4847: uint16(1),
	4848: uint16(anon_sym_LBRACK),
	4849: uint16(647),
	4850: uint16(1),
	4851: uint16(anon_sym_DOT),
	4852: uint16(747),
	4853: uint16(1),
	4854: uint16(anon_sym_COMMA),
	4855: uint16(2),
	4856: uint16(3),
	4857: uint16(1),
	4858: uint16(sym_comment),
	4859: uint16(749),
	4860: uint16(3),
	4861: uint16(sym__dedent),
	4862: uint16(anon_sym_module),
	4863: uint16(anon_sym_extmodule),
	4864: uint16(4),
	4865: uint16(3),
	4866: uint16(1),
	4867: uint16(sym_comment),
	4868: uint16(643),
	4869: uint16(1),
	4870: uint16(anon_sym_LBRACK),
	4871: uint16(647),
	4872: uint16(1),
	4873: uint16(anon_sym_DOT),
	4874: uint16(751),
	4875: uint16(1),
	4876: uint16(anon_sym_COMMA),
	4877: uint16(4),
	4878: uint16(3),
	4879: uint16(1),
	4880: uint16(sym_comment),
	4881: uint16(643),
	4882: uint16(1),
	4883: uint16(anon_sym_LBRACK),
	4884: uint16(647),
	4885: uint16(1),
	4886: uint16(anon_sym_DOT),
	4887: uint16(753),
	4888: uint16(1),
	4889: uint16(anon_sym_COMMA),
	4890: uint16(4),
	4891: uint16(3),
	4892: uint16(1),
	4893: uint16(sym_comment),
	4894: uint16(755),
	4895: uint16(1),
	4896: uint16(anon_sym_COMMA),
	4897: uint16(757),
	4898: uint16(1),
	4899: uint16(anon_sym_RBRACE),
	4900: uint16(283),
	4901: uint16(1),
	4902: uint16(aux_sym_type_repeat1),
	4903: uint16(2),
	4904: uint16(3),
	4905: uint16(1),
	4906: uint16(sym_comment),
	4907: uint16(219),
	4908: uint16(3),
	4909: uint16(anon_sym_COMMA),
	4910: uint16(anon_sym_RBRACE),
	4911: uint16(anon_sym_LBRACK),
	4912: uint16(4),
	4913: uint16(3),
	4914: uint16(1),
	4915: uint16(sym_comment),
	4916: uint16(643),
	4917: uint16(1),
	4918: uint16(anon_sym_LBRACK),
	4919: uint16(647),
	4920: uint16(1),
	4921: uint16(anon_sym_DOT),
	4922: uint16(759),
	4923: uint16(1),
	4924: uint16(anon_sym_RBRACK),
	4925: uint16(4),
	4926: uint16(3),
	4927: uint16(1),
	4928: uint16(sym_comment),
	4929: uint16(698),
	4930: uint16(1),
	4931: uint16(anon_sym_RPAREN),
	4932: uint16(761),
	4933: uint16(1),
	4934: uint16(anon_sym_COMMA),
	4935: uint16(255),
	4936: uint16(1),
	4937: uint16(aux_sym_primitive_operation_repeat1),
	4938: uint16(2),
	4939: uint16(3),
	4940: uint16(1),
	4941: uint16(sym_comment),
	4942: uint16(764),
	4943: uint16(3),
	4944: uint16(sym__dedent),
	4945: uint16(anon_sym_module),
	4946: uint16(anon_sym_extmodule),
	4947: uint16(3),
	4948: uint16(3),
	4949: uint16(1),
	4950: uint16(sym_comment),
	4951: uint16(768),
	4952: uint16(1),
	4953: uint16(sym__indent),
	4954: uint16(766),
	4955: uint16(2),
	4957: uint16(anon_sym_circuit),
	4958: uint16(4),
	4959: uint16(3),
	4960: uint16(1),
	4961: uint16(sym_comment),
	4962: uint16(641),
	4963: uint16(1),
	4964: uint16(anon_sym_COMMA),
	4965: uint16(671),
	4966: uint16(1),
	4967: uint16(anon_sym_RPAREN),
	4968: uint16(303),
	4969: uint16(1),
	4970: uint16(aux_sym_primitive_operation_repeat1),
	4971: uint16(2),
	4972: uint16(3),
	4973: uint16(1),
	4974: uint16(sym_comment),
	4975: uint16(770),
	4976: uint16(3),
	4977: uint16(sym__dedent),
	4978: uint16(anon_sym_module),
	4979: uint16(anon_sym_extmodule),
	4980: uint16(2),
	4981: uint16(3),
	4982: uint16(1),
	4983: uint16(sym_comment),
	4984: uint16(772),
	4985: uint16(3),
	4986: uint16(sym__dedent),
	4987: uint16(anon_sym_module),
	4988: uint16(anon_sym_extmodule),
	4989: uint16(4),
	4990: uint16(3),
	4991: uint16(1),
	4992: uint16(sym_comment),
	4993: uint16(755),
	4994: uint16(1),
	4995: uint16(anon_sym_COMMA),
	4996: uint16(774),
	4997: uint16(1),
	4998: uint16(anon_sym_RBRACE),
	4999: uint16(269),
	5000: uint16(1),
	5001: uint16(aux_sym_type_repeat1),
	5002: uint16(4),
	5003: uint16(3),
	5004: uint16(1),
	5005: uint16(sym_comment),
	5006: uint16(755),
	5007: uint16(1),
	5008: uint16(anon_sym_COMMA),
	5009: uint16(776),
	5010: uint16(1),
	5011: uint16(anon_sym_RBRACE),
	5012: uint16(261),
	5013: uint16(1),
	5014: uint16(aux_sym_type_repeat1),
	5015: uint16(4),
	5016: uint16(3),
	5017: uint16(1),
	5018: uint16(sym_comment),
	5019: uint16(778),
	5020: uint16(1),
	5021: uint16(anon_sym_LBRACK),
	5022: uint16(780),
	5023: uint16(1),
	5024: uint16(sym_info),
	5025: uint16(782),
	5026: uint16(1),
	5027: uint16(sym__newline),
	5028: uint16(2),
	5029: uint16(3),
	5030: uint16(1),
	5031: uint16(sym_comment),
	5032: uint16(285),
	5033: uint16(3),
	5034: uint16(anon_sym_COMMA),
	5035: uint16(anon_sym_RBRACE),
	5036: uint16(anon_sym_LBRACK),
	5037: uint16(4),
	5038: uint16(3),
	5039: uint16(1),
	5040: uint16(sym_comment),
	5041: uint16(643),
	5042: uint16(1),
	5043: uint16(anon_sym_LBRACK),
	5044: uint16(647),
	5045: uint16(1),
	5046: uint16(anon_sym_DOT),
	5047: uint16(784),
	5048: uint16(1),
	5049: uint16(anon_sym_RPAREN),
	5050: uint16(4),
	5051: uint16(3),
	5052: uint16(1),
	5053: uint16(sym_comment),
	5054: uint16(643),
	5055: uint16(1),
	5056: uint16(anon_sym_LBRACK),
	5057: uint16(647),
	5058: uint16(1),
	5059: uint16(anon_sym_DOT),
	5060: uint16(786),
	5061: uint16(1),
	5062: uint16(anon_sym_RPAREN),
	5063: uint16(4),
	5064: uint16(3),
	5065: uint16(1),
	5066: uint16(sym_comment),
	5067: uint16(643),
	5068: uint16(1),
	5069: uint16(anon_sym_LBRACK),
	5070: uint16(647),
	5071: uint16(1),
	5072: uint16(anon_sym_DOT),
	5073: uint16(788),
	5074: uint16(1),
	5075: uint16(anon_sym_RPAREN),
	5076: uint16(4),
	5077: uint16(3),
	5078: uint16(1),
	5079: uint16(sym_comment),
	5080: uint16(643),
	5081: uint16(1),
	5082: uint16(anon_sym_LBRACK),
	5083: uint16(647),
	5084: uint16(1),
	5085: uint16(anon_sym_DOT),
	5086: uint16(790),
	5087: uint16(1),
	5088: uint16(anon_sym_COLON),
	5089: uint16(4),
	5090: uint16(3),
	5091: uint16(1),
	5092: uint16(sym_comment),
	5093: uint16(792),
	5094: uint16(1),
	5095: uint16(anon_sym_COMMA),
	5096: uint16(795),
	5097: uint16(1),
	5098: uint16(anon_sym_RBRACE),
	5099: uint16(269),
	5100: uint16(1),
	5101: uint16(aux_sym_type_repeat1),
	5102: uint16(3),
	5103: uint16(3),
	5104: uint16(1),
	5105: uint16(sym_comment),
	5106: uint16(799),
	5107: uint16(1),
	5108: uint16(anon_sym_LBRACK),
	5109: uint16(797),
	5110: uint16(2),
	5111: uint16(anon_sym_COMMA),
	5112: uint16(anon_sym_RBRACE),
	5113: uint16(4),
	5114: uint16(3),
	5115: uint16(1),
	5116: uint16(sym_comment),
	5117: uint16(755),
	5118: uint16(1),
	5119: uint16(anon_sym_COMMA),
	5120: uint16(801),
	5121: uint16(1),
	5122: uint16(anon_sym_RBRACE),
	5123: uint16(269),
	5124: uint16(1),
	5125: uint16(aux_sym_type_repeat1),
	5126: uint16(4),
	5127: uint16(3),
	5128: uint16(1),
	5129: uint16(sym_comment),
	5130: uint16(643),
	5131: uint16(1),
	5132: uint16(anon_sym_LBRACK),
	5133: uint16(647),
	5134: uint16(1),
	5135: uint16(anon_sym_DOT),
	5136: uint16(803),
	5137: uint16(1),
	5138: uint16(anon_sym_RPAREN),
	5139: uint16(2),
	5140: uint16(3),
	5141: uint16(1),
	5142: uint16(sym_comment),
	5143: uint16(249),
	5144: uint16(3),
	5145: uint16(anon_sym_COMMA),
	5146: uint16(anon_sym_RBRACE),
	5147: uint16(anon_sym_LBRACK),
	5148: uint16(4),
	5149: uint16(3),
	5150: uint16(1),
	5151: uint16(sym_comment),
	5152: uint16(755),
	5153: uint16(1),
	5154: uint16(anon_sym_COMMA),
	5155: uint16(805),
	5156: uint16(1),
	5157: uint16(anon_sym_RBRACE),
	5158: uint16(271),
	5159: uint16(1),
	5160: uint16(aux_sym_type_repeat1),
	5161: uint16(4),
	5162: uint16(3),
	5163: uint16(1),
	5164: uint16(sym_comment),
	5165: uint16(641),
	5166: uint16(1),
	5167: uint16(anon_sym_COMMA),
	5168: uint16(807),
	5169: uint16(1),
	5170: uint16(anon_sym_RPAREN),
	5171: uint16(255),
	5172: uint16(1),
	5173: uint16(aux_sym_primitive_operation_repeat1),
	5174: uint16(4),
	5175: uint16(3),
	5176: uint16(1),
	5177: uint16(sym_comment),
	5178: uint16(809),
	5179: uint16(1),
	5180: uint16(anon_sym_COMMA),
	5181: uint16(811),
	5182: uint16(1),
	5183: uint16(sym__newline),
	5184: uint16(291),
	5185: uint16(1),
	5186: uint16(aux_sym_memory_field_repeat1),
	5187: uint16(4),
	5188: uint16(3),
	5189: uint16(1),
	5190: uint16(sym_comment),
	5191: uint16(813),
	5192: uint16(1),
	5193: uint16(anon_sym_COLON),
	5194: uint16(815),
	5195: uint16(1),
	5196: uint16(anon_sym_when),
	5197: uint16(70),
	5198: uint16(1),
	5199: uint16(sym_when),
	5200: uint16(4),
	5201: uint16(3),
	5202: uint16(1),
	5203: uint16(sym_comment),
	5204: uint16(641),
	5205: uint16(1),
	5206: uint16(anon_sym_COMMA),
	5207: uint16(645),
	5208: uint16(1),
	5209: uint16(anon_sym_RPAREN),
	5210: uint16(275),
	5211: uint16(1),
	5212: uint16(aux_sym_primitive_operation_repeat1),
	5213: uint16(4),
	5214: uint16(3),
	5215: uint16(1),
	5216: uint16(sym_comment),
	5217: uint16(643),
	5218: uint16(1),
	5219: uint16(anon_sym_LBRACK),
	5220: uint16(647),
	5221: uint16(1),
	5222: uint16(anon_sym_DOT),
	5223: uint16(817),
	5224: uint16(1),
	5225: uint16(anon_sym_RBRACK),
	5226: uint16(4),
	5227: uint16(3),
	5228: uint16(1),
	5229: uint16(sym_comment),
	5230: uint16(665),
	5231: uint16(1),
	5232: uint16(anon_sym_COMMA),
	5233: uint16(819),
	5234: uint16(1),
	5235: uint16(anon_sym_RPAREN),
	5236: uint16(294),
	5237: uint16(1),
	5238: uint16(aux_sym_printf_repeat1),
	5239: uint16(4),
	5240: uint16(3),
	5241: uint16(1),
	5242: uint16(sym_comment),
	5243: uint16(643),
	5244: uint16(1),
	5245: uint16(anon_sym_LBRACK),
	5246: uint16(647),
	5247: uint16(1),
	5248: uint16(anon_sym_DOT),
	5249: uint16(821),
	5250: uint16(1),
	5251: uint16(anon_sym_COMMA),
	5252: uint16(4),
	5253: uint16(3),
	5254: uint16(1),
	5255: uint16(sym_comment),
	5256: uint16(643),
	5257: uint16(1),
	5258: uint16(anon_sym_LBRACK),
	5259: uint16(647),
	5260: uint16(1),
	5261: uint16(anon_sym_DOT),
	5262: uint16(823),
	5263: uint16(1),
	5264: uint16(anon_sym_RPAREN),
	5265: uint16(4),
	5266: uint16(3),
	5267: uint16(1),
	5268: uint16(sym_comment),
	5269: uint16(755),
	5270: uint16(1),
	5271: uint16(anon_sym_COMMA),
	5272: uint16(825),
	5273: uint16(1),
	5274: uint16(anon_sym_RBRACE),
	5275: uint16(269),
	5276: uint16(1),
	5277: uint16(aux_sym_type_repeat1),
	5278: uint16(4),
	5279: uint16(3),
	5280: uint16(1),
	5281: uint16(sym_comment),
	5282: uint16(643),
	5283: uint16(1),
	5284: uint16(anon_sym_LBRACK),
	5285: uint16(647),
	5286: uint16(1),
	5287: uint16(anon_sym_DOT),
	5288: uint16(827),
	5289: uint16(1),
	5290: uint16(anon_sym_RBRACK),
	5291: uint16(2),
	5292: uint16(3),
	5293: uint16(1),
	5294: uint16(sym_comment),
	5295: uint16(257),
	5296: uint16(3),
	5297: uint16(sym__newline),
	5298: uint16(anon_sym_LBRACK),
	5299: uint16(sym_info),
	5300: uint16(4),
	5301: uint16(3),
	5302: uint16(1),
	5303: uint16(sym_comment),
	5304: uint16(665),
	5305: uint16(1),
	5306: uint16(anon_sym_COMMA),
	5307: uint16(829),
	5308: uint16(1),
	5309: uint16(anon_sym_RPAREN),
	5310: uint16(296),
	5311: uint16(1),
	5312: uint16(aux_sym_printf_repeat1),
	5313: uint16(3),
	5314: uint16(3),
	5315: uint16(1),
	5316: uint16(sym_comment),
	5317: uint16(799),
	5318: uint16(1),
	5319: uint16(anon_sym_LBRACK),
	5320: uint16(831),
	5321: uint16(2),
	5322: uint16(anon_sym_COMMA),
	5323: uint16(anon_sym_RBRACE),
	5324: uint16(4),
	5325: uint16(3),
	5326: uint16(1),
	5327: uint16(sym_comment),
	5328: uint16(833),
	5329: uint16(1),
	5330: uint16(anon_sym_LPAREN),
	5331: uint16(835),
	5332: uint16(1),
	5333: uint16(sym__indent),
	5334: uint16(54),
	5335: uint16(1),
	5336: uint16(sym_reset_block),
	5337: uint16(4),
	5338: uint16(3),
	5339: uint16(1),
	5340: uint16(sym_comment),
	5341: uint16(643),
	5342: uint16(1),
	5343: uint16(anon_sym_LBRACK),
	5344: uint16(647),
	5345: uint16(1),
	5346: uint16(anon_sym_DOT),
	5347: uint16(837),
	5348: uint16(1),
	5349: uint16(anon_sym_COMMA),
	5350: uint16(2),
	5351: uint16(3),
	5352: uint16(1),
	5353: uint16(sym_comment),
	5354: uint16(249),
	5355: uint16(3),
	5356: uint16(sym__newline),
	5357: uint16(anon_sym_LBRACK),
	5358: uint16(sym_info),
	5359: uint16(4),
	5360: uint16(3),
	5361: uint16(1),
	5362: uint16(sym_comment),
	5363: uint16(809),
	5364: uint16(1),
	5365: uint16(anon_sym_COMMA),
	5366: uint16(839),
	5367: uint16(1),
	5368: uint16(sym__newline),
	5369: uint16(305),
	5370: uint16(1),
	5371: uint16(aux_sym_memory_field_repeat1),
	5372: uint16(2),
	5373: uint16(3),
	5374: uint16(1),
	5375: uint16(sym_comment),
	5376: uint16(261),
	5377: uint16(3),
	5378: uint16(sym__newline),
	5379: uint16(anon_sym_LBRACK),
	5380: uint16(sym_info),
	5381: uint16(2),
	5382: uint16(3),
	5383: uint16(1),
	5384: uint16(sym_comment),
	5385: uint16(285),
	5386: uint16(3),
	5387: uint16(sym__newline),
	5388: uint16(anon_sym_LBRACK),
	5389: uint16(sym_info),
	5390: uint16(4),
	5391: uint16(3),
	5392: uint16(1),
	5393: uint16(sym_comment),
	5394: uint16(665),
	5395: uint16(1),
	5396: uint16(anon_sym_COMMA),
	5397: uint16(841),
	5398: uint16(1),
	5399: uint16(anon_sym_RPAREN),
	5400: uint16(296),
	5401: uint16(1),
	5402: uint16(aux_sym_printf_repeat1),
	5403: uint16(2),
	5404: uint16(3),
	5405: uint16(1),
	5406: uint16(sym_comment),
	5407: uint16(843),
	5408: uint16(3),
	5409: uint16(sym__dedent),
	5410: uint16(anon_sym_module),
	5411: uint16(anon_sym_extmodule),
	5412: uint16(4),
	5413: uint16(3),
	5414: uint16(1),
	5415: uint16(sym_comment),
	5416: uint16(688),
	5417: uint16(1),
	5418: uint16(anon_sym_RPAREN),
	5419: uint16(845),
	5420: uint16(1),
	5421: uint16(anon_sym_COMMA),
	5422: uint16(296),
	5423: uint16(1),
	5424: uint16(aux_sym_printf_repeat1),
	5425: uint16(2),
	5426: uint16(3),
	5427: uint16(1),
	5428: uint16(sym_comment),
	5429: uint16(219),
	5430: uint16(3),
	5431: uint16(sym__newline),
	5432: uint16(anon_sym_LBRACK),
	5433: uint16(sym_info),
	5434: uint16(4),
	5435: uint16(3),
	5436: uint16(1),
	5437: uint16(sym_comment),
	5438: uint16(778),
	5439: uint16(1),
	5440: uint16(anon_sym_LBRACK),
	5441: uint16(848),
	5442: uint16(1),
	5443: uint16(sym_info),
	5444: uint16(850),
	5445: uint16(1),
	5446: uint16(sym__newline),
	5447: uint16(2),
	5448: uint16(3),
	5449: uint16(1),
	5450: uint16(sym_comment),
	5451: uint16(289),
	5452: uint16(3),
	5453: uint16(anon_sym_COMMA),
	5454: uint16(anon_sym_RBRACE),
	5455: uint16(anon_sym_LBRACK),
	5456: uint16(4),
	5457: uint16(3),
	5458: uint16(1),
	5459: uint16(sym_comment),
	5460: uint16(643),
	5461: uint16(1),
	5462: uint16(anon_sym_LBRACK),
	5463: uint16(647),
	5464: uint16(1),
	5465: uint16(anon_sym_DOT),
	5466: uint16(852),
	5467: uint16(1),
	5468: uint16(anon_sym_COMMA),
	5469: uint16(4),
	5470: uint16(3),
	5471: uint16(1),
	5472: uint16(sym_comment),
	5473: uint16(643),
	5474: uint16(1),
	5475: uint16(anon_sym_LBRACK),
	5476: uint16(647),
	5477: uint16(1),
	5478: uint16(anon_sym_DOT),
	5479: uint16(854),
	5480: uint16(1),
	5481: uint16(anon_sym_COMMA),
	5482: uint16(2),
	5483: uint16(3),
	5484: uint16(1),
	5485: uint16(sym_comment),
	5486: uint16(257),
	5487: uint16(3),
	5488: uint16(anon_sym_COMMA),
	5489: uint16(anon_sym_RBRACE),
	5490: uint16(anon_sym_LBRACK),
	5491: uint16(4),
	5492: uint16(3),
	5493: uint16(1),
	5494: uint16(sym_comment),
	5495: uint16(641),
	5496: uint16(1),
	5497: uint16(anon_sym_COMMA),
	5498: uint16(856),
	5499: uint16(1),
	5500: uint16(anon_sym_RPAREN),
	5501: uint16(255),
	5502: uint16(1),
	5503: uint16(aux_sym_primitive_operation_repeat1),
	5504: uint16(4),
	5505: uint16(3),
	5506: uint16(1),
	5507: uint16(sym_comment),
	5508: uint16(643),
	5509: uint16(1),
	5510: uint16(anon_sym_LBRACK),
	5511: uint16(647),
	5512: uint16(1),
	5513: uint16(anon_sym_DOT),
	5514: uint16(858),
	5515: uint16(1),
	5516: uint16(anon_sym_COMMA),
	5517: uint16(4),
	5518: uint16(3),
	5519: uint16(1),
	5520: uint16(sym_comment),
	5521: uint16(860),
	5522: uint16(1),
	5523: uint16(anon_sym_COMMA),
	5524: uint16(863),
	5525: uint16(1),
	5526: uint16(sym__newline),
	5527: uint16(305),
	5528: uint16(1),
	5529: uint16(aux_sym_memory_field_repeat1),
	5530: uint16(2),
	5531: uint16(3),
	5532: uint16(1),
	5533: uint16(sym_comment),
	5534: uint16(863),
	5535: uint16(2),
	5536: uint16(sym__newline),
	5537: uint16(anon_sym_COMMA),
	5538: uint16(3),
	5539: uint16(3),
	5540: uint16(1),
	5541: uint16(sym_comment),
	5542: uint16(865),
	5543: uint16(1),
	5544: uint16(anon_sym_DQUOTE),
	5545: uint16(360),
	5546: uint16(1),
	5547: uint16(sym_string),
	5548: uint16(3),
	5549: uint16(3),
	5550: uint16(1),
	5551: uint16(sym_comment),
	5552: uint16(778),
	5553: uint16(1),
	5554: uint16(anon_sym_LBRACK),
	5555: uint16(811),
	5556: uint16(1),
	5557: uint16(sym__newline),
	5558: uint16(2),
	5559: uint16(3),
	5560: uint16(1),
	5561: uint16(sym_comment),
	5562: uint16(867),
	5563: uint16(2),
	5564: uint16(anon_sym_COMMA),
	5565: uint16(anon_sym_RPAREN),
	5566: uint16(2),
	5567: uint16(3),
	5568: uint16(1),
	5569: uint16(sym_comment),
	5570: uint16(869),
	5571: uint16(2),
	5572: uint16(sym__dedent),
	5573: uint16(anon_sym_parameter),
	5574: uint16(2),
	5575: uint16(3),
	5576: uint16(1),
	5577: uint16(sym_comment),
	5578: uint16(871),
	5579: uint16(2),
	5580: uint16(anon_sym_COMMA),
	5581: uint16(anon_sym_RPAREN),
	5582: uint16(2),
	5583: uint16(3),
	5584: uint16(1),
	5585: uint16(sym_comment),
	5586: uint16(795),
	5587: uint16(2),
	5588: uint16(anon_sym_COMMA),
	5589: uint16(anon_sym_RBRACE),
	5590: uint16(3),
	5591: uint16(3),
	5592: uint16(1),
	5593: uint16(sym_comment),
	5594: uint16(715),
	5595: uint16(1),
	5596: uint16(anon_sym_reset),
	5597: uint16(385),
	5598: uint16(1),
	5599: uint16(sym__reset),
	5600: uint16(2),
	5601: uint16(3),
	5602: uint16(1),
	5603: uint16(sym_comment),
	5604: uint16(873),
	5605: uint16(2),
	5606: uint16(sym__newline),
	5607: uint16(sym_info),
	5608: uint16(3),
	5609: uint16(3),
	5610: uint16(1),
	5611: uint16(sym_comment),
	5612: uint16(875),
	5613: uint16(1),
	5614: uint16(sym_info),
	5615: uint16(877),
	5616: uint16(1),
	5617: uint16(sym__newline),
	5618: uint16(3),
	5619: uint16(3),
	5620: uint16(1),
	5621: uint16(sym_comment),
	5622: uint16(799),
	5623: uint16(1),
	5624: uint16(anon_sym_LBRACK),
	5625: uint16(879),
	5626: uint16(1),
	5627: uint16(anon_sym_COMMA),
	5628: uint16(3),
	5629: uint16(3),
	5630: uint16(1),
	5631: uint16(sym_comment),
	5632: uint16(881),
	5633: uint16(1),
	5634: uint16(anon_sym_LT),
	5635: uint16(883),
	5636: uint16(1),
	5637: uint16(anon_sym_LPAREN),
	5638: uint16(3),
	5639: uint16(3),
	5640: uint16(1),
	5641: uint16(sym_comment),
	5642: uint16(715),
	5643: uint16(1),
	5644: uint16(anon_sym_reset),
	5645: uint16(398),
	5646: uint16(1),
	5647: uint16(sym__reset),
	5648: uint16(2),
	5649: uint16(3),
	5650: uint16(1),
	5651: uint16(sym_comment),
	5652: uint16(698),
	5653: uint16(2),
	5654: uint16(anon_sym_COMMA),
	5655: uint16(anon_sym_RPAREN),
	5656: uint16(2),
	5657: uint16(3),
	5658: uint16(1),
	5659: uint16(sym_comment),
	5660: uint16(885),
	5661: uint16(2),
	5663: uint16(anon_sym_circuit),
	5664: uint16(2),
	5665: uint16(3),
	5666: uint16(1),
	5667: uint16(sym_comment),
	5668: uint16(887),
	5669: uint16(2),
	5670: uint16(sym__dedent),
	5671: uint16(anon_sym_parameter),
	5672: uint16(3),
	5673: uint16(3),
	5674: uint16(1),
	5675: uint16(sym_comment),
	5676: uint16(865),
	5677: uint16(1),
	5678: uint16(anon_sym_DQUOTE),
	5679: uint16(280),
	5680: uint16(1),
	5681: uint16(sym_string),
	5682: uint16(2),
	5683: uint16(3),
	5684: uint16(1),
	5685: uint16(sym_comment),
	5686: uint16(889),
	5687: uint16(2),
	5688: uint16(sym__newline),
	5689: uint16(sym_info),
	5690: uint16(2),
	5691: uint16(3),
	5692: uint16(1),
	5693: uint16(sym_comment),
	5694: uint16(891),
	5695: uint16(2),
	5696: uint16(sym__newline),
	5697: uint16(sym_info),
	5698: uint16(2),
	5699: uint16(3),
	5700: uint16(1),
	5701: uint16(sym_comment),
	5702: uint16(893),
	5703: uint16(2),
	5705: uint16(anon_sym_circuit),
	5706: uint16(3),
	5707: uint16(3),
	5708: uint16(1),
	5709: uint16(sym_comment),
	5710: uint16(895),
	5711: uint16(1),
	5712: uint16(sym_info),
	5713: uint16(897),
	5714: uint16(1),
	5715: uint16(sym__indent),
	5716: uint16(2),
	5717: uint16(3),
	5718: uint16(1),
	5719: uint16(sym_comment),
	5720: uint16(899),
	5721: uint16(2),
	5723: uint16(anon_sym_circuit),
	5724: uint16(2),
	5725: uint16(3),
	5726: uint16(1),
	5727: uint16(sym_comment),
	5728: uint16(901),
	5729: uint16(1),
	5730: uint16(anon_sym_COLON),
	5731: uint16(2),
	5732: uint16(3),
	5733: uint16(1),
	5734: uint16(sym_comment),
	5735: uint16(903),
	5736: uint16(1),
	5737: uint16(anon_sym_LPAREN),
	5738: uint16(2),
	5739: uint16(3),
	5740: uint16(1),
	5741: uint16(sym_comment),
	5742: uint16(905),
	5743: uint16(1),
	5744: uint16(sym__dedent),
	5745: uint16(2),
	5746: uint16(3),
	5747: uint16(1),
	5748: uint16(sym_comment),
	5749: uint16(907),
	5750: uint16(1),
	5751: uint16(sym__newline),
	5752: uint16(2),
	5753: uint16(3),
	5754: uint16(1),
	5755: uint16(sym_comment),
	5756: uint16(909),
	5757: uint16(1),
	5758: uint16(anon_sym_COLON),
	5759: uint16(2),
	5760: uint16(3),
	5761: uint16(1),
	5762: uint16(sym_comment),
	5763: uint16(911),
	5764: uint16(1),
	5765: uint16(anon_sym_EQ_GT),
	5766: uint16(2),
	5767: uint16(3),
	5768: uint16(1),
	5769: uint16(sym_comment),
	5770: uint16(913),
	5771: uint16(1),
	5772: uint16(anon_sym_EQ_GT),
	5773: uint16(2),
	5774: uint16(3),
	5775: uint16(1),
	5776: uint16(sym_comment),
	5777: uint16(915),
	5778: uint16(1),
	5779: uint16(anon_sym_EQ_GT),
	5780: uint16(2),
	5781: uint16(3),
	5782: uint16(1),
	5783: uint16(sym_comment),
	5784: uint16(917),
	5785: uint16(1),
	5786: uint16(sym__dedent),
	5787: uint16(2),
	5788: uint16(3),
	5789: uint16(1),
	5790: uint16(sym_comment),
	5791: uint16(919),
	5792: uint16(1),
	5793: uint16(anon_sym_EQ_GT),
	5794: uint16(2),
	5795: uint16(3),
	5796: uint16(1),
	5797: uint16(sym_comment),
	5798: uint16(921),
	5799: uint16(1),
	5800: uint16(anon_sym_COLON),
	5801: uint16(2),
	5802: uint16(3),
	5803: uint16(1),
	5804: uint16(sym_comment),
	5805: uint16(923),
	5806: uint16(1),
	5807: uint16(anon_sym_COLON),
	5808: uint16(2),
	5809: uint16(3),
	5810: uint16(1),
	5811: uint16(sym_comment),
	5812: uint16(925),
	5813: uint16(1),
	5814: uint16(anon_sym_COLON),
	5815: uint16(2),
	5816: uint16(3),
	5817: uint16(1),
	5818: uint16(sym_comment),
	5819: uint16(927),
	5820: uint16(1),
	5821: uint16(anon_sym_COLON),
	5822: uint16(2),
	5823: uint16(3),
	5824: uint16(1),
	5825: uint16(sym_comment),
	5826: uint16(163),
	5827: uint16(1),
	5828: uint16(sym__newline),
	5829: uint16(2),
	5830: uint16(3),
	5831: uint16(1),
	5832: uint16(sym_comment),
	5833: uint16(700),
	5834: uint16(1),
	5835: uint16(sym__newline),
	5836: uint16(2),
	5837: uint16(3),
	5838: uint16(1),
	5839: uint16(sym_comment),
	5840: uint16(929),
	5841: uint16(1),
	5842: uint16(anon_sym_of),
	5843: uint16(2),
	5844: uint16(3),
	5845: uint16(1),
	5846: uint16(sym_comment),
	5847: uint16(931),
	5848: uint16(1),
	5849: uint16(anon_sym_EQ),
	5850: uint16(2),
	5851: uint16(3),
	5852: uint16(1),
	5853: uint16(sym_comment),
	5854: uint16(729),
	5855: uint16(1),
	5856: uint16(sym__newline),
	5857: uint16(2),
	5858: uint16(3),
	5859: uint16(1),
	5860: uint16(sym_comment),
	5861: uint16(933),
	5862: uint16(1),
	5863: uint16(anon_sym_COLON),
	5864: uint16(2),
	5865: uint16(3),
	5866: uint16(1),
	5867: uint16(sym_comment),
	5868: uint16(935),
	5869: uint16(1),
	5870: uint16(anon_sym_GT),
	5871: uint16(2),
	5872: uint16(3),
	5873: uint16(1),
	5874: uint16(sym_comment),
	5875: uint16(937),
	5876: uint16(1),
	5877: uint16(anon_sym_LPAREN),
	5878: uint16(2),
	5879: uint16(3),
	5880: uint16(1),
	5881: uint16(sym_comment),
	5882: uint16(939),
	5883: uint16(1),
	5884: uint16(sym_identifier),
	5885: uint16(2),
	5886: uint16(3),
	5887: uint16(1),
	5888: uint16(sym_comment),
	5889: uint16(941),
	5890: uint16(1),
	5891: uint16(sym_identifier),
	5892: uint16(2),
	5893: uint16(3),
	5894: uint16(1),
	5895: uint16(sym_comment),
	5896: uint16(943),
	5897: uint16(1),
	5898: uint16(anon_sym_GT),
	5899: uint16(2),
	5900: uint16(3),
	5901: uint16(1),
	5902: uint16(sym_comment),
	5903: uint16(945),
	5904: uint16(1),
	5905: uint16(aux_sym_uint_token1),
	5906: uint16(2),
	5907: uint16(3),
	5908: uint16(1),
	5909: uint16(sym_comment),
	5910: uint16(947),
	5911: uint16(1),
	5912: uint16(sym_identifier),
	5913: uint16(2),
	5914: uint16(3),
	5915: uint16(1),
	5916: uint16(sym_comment),
	5917: uint16(949),
	5918: uint16(1),
	5919: uint16(anon_sym_RPAREN),
	5920: uint16(2),
	5921: uint16(3),
	5922: uint16(1),
	5923: uint16(sym_comment),
	5924: uint16(951),
	5925: uint16(1),
	5926: uint16(anon_sym_EQ_GT),
	5927: uint16(2),
	5928: uint16(3),
	5929: uint16(1),
	5930: uint16(sym_comment),
	5931: uint16(953),
	5932: uint16(1),
	5933: uint16(sym__newline),
	5934: uint16(2),
	5935: uint16(3),
	5936: uint16(1),
	5937: uint16(sym_comment),
	5938: uint16(955),
	5939: uint16(1),
	5940: uint16(anon_sym_GT),
	5941: uint16(2),
	5942: uint16(3),
	5943: uint16(1),
	5944: uint16(sym_comment),
	5945: uint16(957),
	5946: uint16(1),
	5947: uint16(anon_sym_COLON),
	5948: uint16(2),
	5949: uint16(3),
	5950: uint16(1),
	5951: uint16(sym_comment),
	5952: uint16(959),
	5953: uint16(1),
	5954: uint16(anon_sym_RPAREN),
	5955: uint16(2),
	5956: uint16(3),
	5957: uint16(1),
	5958: uint16(sym_comment),
	5959: uint16(961),
	5960: uint16(1),
	5961: uint16(sym_identifier),
	5962: uint16(2),
	5963: uint16(3),
	5964: uint16(1),
	5965: uint16(sym_comment),
	5966: uint16(963),
	5967: uint16(1),
	5968: uint16(sym_identifier),
	5969: uint16(2),
	5970: uint16(3),
	5971: uint16(1),
	5972: uint16(sym_comment),
	5973: uint16(965),
	5974: uint16(1),
	5975: uint16(sym_identifier),
	5976: uint16(2),
	5977: uint16(3),
	5978: uint16(1),
	5979: uint16(sym_comment),
	5980: uint16(873),
	5981: uint16(1),
	5982: uint16(anon_sym_RPAREN),
	5983: uint16(2),
	5984: uint16(3),
	5985: uint16(1),
	5986: uint16(sym_comment),
	5987: uint16(967),
	5988: uint16(1),
	5989: uint16(sym_identifier),
	5990: uint16(2),
	5991: uint16(3),
	5992: uint16(1),
	5993: uint16(sym_comment),
	5994: uint16(891),
	5995: uint16(1),
	5996: uint16(anon_sym_RPAREN),
	5997: uint16(2),
	5998: uint16(3),
	5999: uint16(1),
	6000: uint16(sym_comment),
	6001: uint16(889),
	6002: uint16(1),
	6003: uint16(anon_sym_RPAREN),
	6004: uint16(2),
	6005: uint16(3),
	6006: uint16(1),
	6007: uint16(sym_comment),
	6008: uint16(969),
	6009: uint16(1),
	6010: uint16(anon_sym_COLON),
	6011: uint16(2),
	6012: uint16(3),
	6013: uint16(1),
	6014: uint16(sym_comment),
	6015: uint16(971),
	6016: uint16(1),
	6017: uint16(sym_identifier),
	6018: uint16(2),
	6019: uint16(3),
	6020: uint16(1),
	6021: uint16(sym_comment),
	6022: uint16(973),
	6023: uint16(1),
	6024: uint16(sym_identifier),
	6025: uint16(2),
	6026: uint16(3),
	6027: uint16(1),
	6028: uint16(sym_comment),
	6029: uint16(975),
	6030: uint16(1),
	6031: uint16(sym_identifier),
	6032: uint16(2),
	6033: uint16(3),
	6034: uint16(1),
	6035: uint16(sym_comment),
	6036: uint16(977),
	6037: uint16(1),
	6038: uint16(anon_sym_COMMA),
	6039: uint16(2),
	6040: uint16(3),
	6041: uint16(1),
	6042: uint16(sym_comment),
	6043: uint16(979),
	6044: uint16(1),
	6045: uint16(sym_identifier),
	6046: uint16(2),
	6047: uint16(3),
	6048: uint16(1),
	6049: uint16(sym_comment),
	6050: uint16(981),
	6051: uint16(1),
	6052: uint16(sym_identifier),
	6053: uint16(2),
	6054: uint16(3),
	6055: uint16(1),
	6056: uint16(sym_comment),
	6057: uint16(983),
	6058: uint16(1),
	6059: uint16(sym_identifier),
	6060: uint16(2),
	6061: uint16(3),
	6062: uint16(1),
	6063: uint16(sym_comment),
	6064: uint16(985),
	6065: uint16(1),
	6066: uint16(anon_sym_LT),
	6067: uint16(2),
	6068: uint16(3),
	6069: uint16(1),
	6070: uint16(sym_comment),
	6071: uint16(987),
	6072: uint16(1),
	6073: uint16(anon_sym_GT),
	6074: uint16(2),
	6075: uint16(3),
	6076: uint16(1),
	6077: uint16(sym_comment),
	6078: uint16(989),
	6079: uint16(1),
	6080: uint16(anon_sym_COLON),
	6081: uint16(2),
	6082: uint16(3),
	6083: uint16(1),
	6084: uint16(sym_comment),
	6085: uint16(991),
	6086: uint16(1),
	6087: uint16(aux_sym_uint_token1),
	6088: uint16(2),
	6089: uint16(3),
	6090: uint16(1),
	6091: uint16(sym_comment),
	6092: uint16(993),
	6093: uint16(1),
	6094: uint16(anon_sym_RPAREN),
	6095: uint16(2),
	6096: uint16(3),
	6097: uint16(1),
	6098: uint16(sym_comment),
	6099: uint16(995),
	6100: uint16(1),
	6101: uint16(sym_identifier),
	6102: uint16(2),
	6103: uint16(3),
	6104: uint16(1),
	6105: uint16(sym_comment),
	6106: uint16(997),
	6107: uint16(1),
	6108: uint16(anon_sym_RBRACK),
	6109: uint16(2),
	6110: uint16(3),
	6111: uint16(1),
	6112: uint16(sym_comment),
	6113: uint16(999),
	6114: uint16(1),
	6115: uint16(sym_identifier),
	6116: uint16(2),
	6117: uint16(3),
	6118: uint16(1),
	6119: uint16(sym_comment),
	6120: uint16(1001),
	6121: uint16(1),
	6122: uint16(anon_sym_RPAREN),
	6123: uint16(2),
	6124: uint16(3),
	6125: uint16(1),
	6126: uint16(sym_comment),
	6127: uint16(1003),
	6128: uint16(1),
	6129: uint16(anon_sym_RPAREN),
	6130: uint16(2),
	6131: uint16(3),
	6132: uint16(1),
	6133: uint16(sym_comment),
	6134: uint16(1005),
	6135: uint16(1),
	6136: uint16(anon_sym_LPAREN),
	6137: uint16(2),
	6138: uint16(3),
	6139: uint16(1),
	6140: uint16(sym_comment),
	6141: uint16(1007),
	6142: uint16(1),
	6143: uint16(anon_sym_invalid),
	6144: uint16(2),
	6145: uint16(3),
	6146: uint16(1),
	6147: uint16(sym_comment),
	6148: uint16(1009),
	6149: uint16(1),
	6150: uint16(sym__newline),
	6151: uint16(2),
	6152: uint16(3),
	6153: uint16(1),
	6154: uint16(sym_comment),
	6155: uint16(811),
	6156: uint16(1),
	6157: uint16(sym__newline),
	6158: uint16(2),
	6159: uint16(3),
	6160: uint16(1),
	6161: uint16(sym_comment),
	6162: uint16(1011),
	6163: uint16(1),
	6164: uint16(sym__indent),
	6165: uint16(2),
	6166: uint16(3),
	6167: uint16(1),
	6168: uint16(sym_comment),
	6169: uint16(801),
	6170: uint16(1),
	6171: uint16(anon_sym_GT),
	6172: uint16(2),
	6173: uint16(3),
	6174: uint16(1),
	6175: uint16(sym_comment),
	6176: uint16(1013),
	6177: uint16(1),
	6178: uint16(anon_sym_GT),
	6179: uint16(2),
	6180: uint16(3),
	6181: uint16(1),
	6182: uint16(sym_comment),
	6183: uint16(1015),
	6184: uint16(1),
	6185: uint16(anon_sym_LPAREN),
	6186: uint16(2),
	6187: uint16(3),
	6188: uint16(1),
	6189: uint16(sym_comment),
	6190: uint16(801),
	6191: uint16(1),
	6192: uint16(anon_sym_RBRACK),
	6193: uint16(2),
	6194: uint16(3),
	6195: uint16(1),
	6196: uint16(sym_comment),
	6197: uint16(1017),
	6198: uint16(1),
	6199: uint16(anon_sym_COLON),
	6200: uint16(2),
	6201: uint16(3),
	6202: uint16(1),
	6203: uint16(sym_comment),
	6204: uint16(1019),
	6205: uint16(1),
	6206: uint16(anon_sym_GT),
	6207: uint16(2),
	6208: uint16(3),
	6209: uint16(1),
	6210: uint16(sym_comment),
	6211: uint16(1021),
	6212: uint16(1),
	6213: uint16(anon_sym_GT),
	6214: uint16(2),
	6215: uint16(3),
	6216: uint16(1),
	6217: uint16(sym_comment),
	6218: uint16(1023),
	6219: uint16(1),
	6220: uint16(anon_sym_RPAREN),
	6221: uint16(2),
	6222: uint16(3),
	6223: uint16(1),
	6224: uint16(sym_comment),
	6225: uint16(1025),
	6226: uint16(1),
	6227: uint16(anon_sym_LPAREN),
	6228: uint16(2),
	6229: uint16(3),
	6230: uint16(1),
	6231: uint16(sym_comment),
	6232: uint16(1027),
	6233: uint16(1),
	6234: uint16(anon_sym_GT),
	6235: uint16(2),
	6236: uint16(3),
	6237: uint16(1),
	6238: uint16(sym_comment),
	6239: uint16(1029),
	6240: uint16(1),
	6241: uint16(anon_sym_mport),
	6242: uint16(2),
	6243: uint16(3),
	6244: uint16(1),
	6245: uint16(sym_comment),
	6246: uint16(774),
	6247: uint16(1),
	6248: uint16(anon_sym_GT),
	6249: uint16(2),
	6250: uint16(3),
	6251: uint16(1),
	6252: uint16(sym_comment),
	6253: uint16(1031),
	6254: uint16(1),
	6255: uint16(anon_sym_GT),
	6256: uint16(2),
	6257: uint16(3),
	6258: uint16(1),
	6259: uint16(sym_comment),
	6260: uint16(1033),
	6261: uint16(1),
	6262: uint16(sym__newline),
	6263: uint16(2),
	6264: uint16(3),
	6265: uint16(1),
	6266: uint16(sym_comment),
	6267: uint16(774),
	6268: uint16(1),
	6269: uint16(anon_sym_RBRACK),
	6270: uint16(2),
	6271: uint16(3),
	6272: uint16(1),
	6273: uint16(sym_comment),
	6274: uint16(1035),
	6275: uint16(1),
	6276: uint16(anon_sym_GT),
	6277: uint16(2),
	6278: uint16(3),
	6279: uint16(1),
	6280: uint16(sym_comment),
	6281: uint16(1037),
	6282: uint16(1),
	6283: uint16(anon_sym_GT),
	6284: uint16(2),
	6285: uint16(3),
	6286: uint16(1),
	6287: uint16(sym_comment),
	6288: uint16(1039),
	6289: uint16(1),
	6290: uint16(anon_sym_EQ),
	6291: uint16(2),
	6292: uint16(3),
	6293: uint16(1),
	6294: uint16(sym_comment),
	6295: uint16(1041),
	6296: uint16(1),
	6297: uint16(sym_identifier),
	6298: uint16(2),
	6299: uint16(3),
	6300: uint16(1),
	6301: uint16(sym_comment),
	6302: uint16(1043),
	6303: uint16(1),
	6305: uint16(2),
	6306: uint16(3),
	6307: uint16(1),
	6308: uint16(sym_comment),
	6309: uint16(1045),
	6310: uint16(1),
	6311: uint16(anon_sym_EQ),
	6312: uint16(2),
	6313: uint16(3),
	6314: uint16(1),
	6315: uint16(sym_comment),
	6316: uint16(1047),
	6317: uint16(1),
	6318: uint16(anon_sym_LPAREN),
	6319: uint16(2),
	6320: uint16(3),
	6321: uint16(1),
	6322: uint16(sym_comment),
	6323: uint16(1049),
	6324: uint16(1),
	6325: uint16(sym__newline),
	6326: uint16(2),
	6327: uint16(3),
	6328: uint16(1),
	6329: uint16(sym_comment),
	6330: uint16(867),
	6331: uint16(1),
	6332: uint16(sym__newline),
	6333: uint16(2),
	6334: uint16(3),
	6335: uint16(1),
	6336: uint16(sym_comment),
	6337: uint16(782),
	6338: uint16(1),
	6339: uint16(sym__newline),
	6340: uint16(2),
	6341: uint16(3),
	6342: uint16(1),
	6343: uint16(sym_comment),
	6344: uint16(1051),
	6345: uint16(1),
	6346: uint16(anon_sym_GT),
	6347: uint16(2),
	6348: uint16(3),
	6349: uint16(1),
	6350: uint16(sym_comment),
	6351: uint16(1053),
	6352: uint16(1),
	6353: uint16(anon_sym_GT),
	6354: uint16(2),
	6355: uint16(3),
	6356: uint16(1),
	6357: uint16(sym_comment),
	6358: uint16(1055),
	6359: uint16(1),
	6360: uint16(sym__newline),
	6361: uint16(2),
	6362: uint16(3),
	6363: uint16(1),
	6364: uint16(sym_comment),
	6365: uint16(1057),
	6366: uint16(1),
	6367: uint16(anon_sym_RBRACK),
	6368: uint16(2),
	6369: uint16(3),
	6370: uint16(1),
	6371: uint16(sym_comment),
	6372: uint16(1059),
	6373: uint16(1),
	6374: uint16(anon_sym_LBRACK),
	6375: uint16(2),
	6376: uint16(3),
	6377: uint16(1),
	6378: uint16(sym_comment),
	6379: uint16(1061),
	6380: uint16(1),
	6381: uint16(anon_sym_RPAREN),
	6382: uint16(2),
	6383: uint16(3),
	6384: uint16(1),
	6385: uint16(sym_comment),
	6386: uint16(1063),
	6387: uint16(1),
	6388: uint16(anon_sym_EQ),
	6389: uint16(2),
	6390: uint16(3),
	6391: uint16(1),
	6392: uint16(sym_comment),
	6393: uint16(1065),
	6394: uint16(1),
	6395: uint16(anon_sym_GT),
	6396: uint16(2),
	6397: uint16(3),
	6398: uint16(1),
	6399: uint16(sym_comment),
	6400: uint16(1067),
	6401: uint16(1),
	6402: uint16(anon_sym_GT),
	6403: uint16(2),
	6404: uint16(3),
	6405: uint16(1),
	6406: uint16(sym_comment),
	6407: uint16(1069),
	6408: uint16(1),
	6409: uint16(anon_sym_LPAREN),
	6410: uint16(2),
	6411: uint16(3),
	6412: uint16(1),
	6413: uint16(sym_comment),
	6414: uint16(1071),
	6415: uint16(1),
	6416: uint16(anon_sym_LPAREN),
	6417: uint16(2),
	6418: uint16(3),
	6419: uint16(1),
	6420: uint16(sym_comment),
	6421: uint16(871),
	6422: uint16(1),
	6423: uint16(sym__newline),
	6424: uint16(2),
	6425: uint16(3),
	6426: uint16(1),
	6427: uint16(sym_comment),
	6428: uint16(1073),
	6429: uint16(1),
	6430: uint16(sym__newline),
	6431: uint16(2),
	6432: uint16(3),
	6433: uint16(1),
	6434: uint16(sym_comment),
	6435: uint16(1075),
	6436: uint16(1),
	6437: uint16(sym_identifier),
	6438: uint16(2),
	6439: uint16(3),
	6440: uint16(1),
	6441: uint16(sym_comment),
	6442: uint16(1077),
	6443: uint16(1),
	6444: uint16(anon_sym_mport),
	6445: uint16(2),
	6446: uint16(3),
	6447: uint16(1),
	6448: uint16(sym_comment),
	6449: uint16(825),
	6450: uint16(1),
	6451: uint16(anon_sym_GT),
	6452: uint16(2),
	6453: uint16(3),
	6454: uint16(1),
	6455: uint16(sym_comment),
	6456: uint16(1079),
	6457: uint16(1),
	6458: uint16(anon_sym_COLON),
	6459: uint16(2),
	6460: uint16(3),
	6461: uint16(1),
	6462: uint16(sym_comment),
	6463: uint16(1081),
	6464: uint16(1),
	6465: uint16(anon_sym_LPAREN),
	6466: uint16(2),
	6467: uint16(3),
	6468: uint16(1),
	6469: uint16(sym_comment),
	6470: uint16(1083),
	6471: uint16(1),
	6472: uint16(sym_identifier),
	6473: uint16(2),
	6474: uint16(3),
	6475: uint16(1),
	6476: uint16(sym_comment),
	6477: uint16(1085),
	6478: uint16(1),
	6479: uint16(anon_sym_LT),
	6480: uint16(2),
	6481: uint16(3),
	6482: uint16(1),
	6483: uint16(sym_comment),
	6484: uint16(1087),
	6485: uint16(1),
	6486: uint16(sym_identifier),
	6487: uint16(2),
	6488: uint16(3),
	6489: uint16(1),
	6490: uint16(sym_comment),
	6491: uint16(1089),
	6492: uint16(1),
	6493: uint16(anon_sym_LT),
	6494: uint16(2),
	6495: uint16(3),
	6496: uint16(1),
	6497: uint16(sym_comment),
	6498: uint16(1091),
	6499: uint16(1),
	6500: uint16(anon_sym_LPAREN),
	6501: uint16(2),
	6502: uint16(3),
	6503: uint16(1),
	6504: uint16(sym_comment),
	6505: uint16(825),
	6506: uint16(1),
	6507: uint16(anon_sym_RBRACK),
	6508: uint16(2),
	6509: uint16(3),
	6510: uint16(1),
	6511: uint16(sym_comment),
	6512: uint16(1093),
	6513: uint16(1),
	6514: uint16(anon_sym_LPAREN),
	6515: uint16(2),
	6516: uint16(3),
	6517: uint16(1),
	6518: uint16(sym_comment),
	6519: uint16(1095),
	6520: uint16(1),
	6521: uint16(anon_sym_GT),
	6522: uint16(2),
	6523: uint16(3),
	6524: uint16(1),
	6525: uint16(sym_comment),
	6526: uint16(1097),
	6527: uint16(1),
	6528: uint16(anon_sym_EQ_GT),
}

var ts_small_parse_table_map = [348]uint32_t{
	1:   uint32(91),
	2:   uint32(182),
	3:   uint32(273),
	4:   uint32(364),
	5:   uint32(455),
	6:   uint32(531),
	7:   uint32(607),
	8:   uint32(683),
	9:   uint32(759),
	10:  uint32(835),
	11:  uint32(911),
	12:  uint32(987),
	13:  uint32(1063),
	14:  uint32(1139),
	15:  uint32(1215),
	16:  uint32(1291),
	17:  uint32(1367),
	18:  uint32(1443),
	19:  uint32(1519),
	20:  uint32(1595),
	21:  uint32(1671),
	22:  uint32(1747),
	23:  uint32(1823),
	24:  uint32(1899),
	25:  uint32(1975),
	26:  uint32(2051),
	27:  uint32(2127),
	28:  uint32(2203),
	29:  uint32(2279),
	30:  uint32(2355),
	31:  uint32(2431),
	32:  uint32(2507),
	33:  uint32(2583),
	34:  uint32(2659),
	35:  uint32(2735),
	36:  uint32(2765),
	37:  uint32(2796),
	38:  uint32(2823),
	39:  uint32(2850),
	40:  uint32(2877),
	41:  uint32(2906),
	42:  uint32(2937),
	43:  uint32(2968),
	44:  uint32(2995),
	45:  uint32(3022),
	46:  uint32(3053),
	47:  uint32(3071),
	48:  uint32(3087),
	49:  uint32(3105),
	50:  uint32(3121),
	51:  uint32(3136),
	52:  uint32(3151),
	53:  uint32(3166),
	54:  uint32(3189),
	55:  uint32(3212),
	56:  uint32(3235),
	57:  uint32(3262),
	58:  uint32(3277),
	59:  uint32(3304),
	60:  uint32(3319),
	61:  uint32(3334),
	62:  uint32(3357),
	63:  uint32(3384),
	64:  uint32(3407),
	65:  uint32(3430),
	66:  uint32(3445),
	67:  uint32(3468),
	68:  uint32(3483),
	69:  uint32(3498),
	70:  uint32(3521),
	71:  uint32(3536),
	72:  uint32(3551),
	73:  uint32(3566),
	74:  uint32(3590),
	75:  uint32(3612),
	76:  uint32(3634),
	77:  uint32(3648),
	78:  uint32(3670),
	79:  uint32(3692),
	80:  uint32(3714),
	81:  uint32(3733),
	82:  uint32(3752),
	83:  uint32(3771),
	84:  uint32(3790),
	85:  uint32(3809),
	86:  uint32(3828),
	87:  uint32(3847),
	88:  uint32(3866),
	89:  uint32(3885),
	90:  uint32(3904),
	91:  uint32(3923),
	92:  uint32(3942),
	93:  uint32(3961),
	94:  uint32(3980),
	95:  uint32(3999),
	96:  uint32(4015),
	97:  uint32(4031),
	98:  uint32(4047),
	99:  uint32(4063),
	100: uint32(4079),
	101: uint32(4095),
	102: uint32(4115),
	103: uint32(4135),
	104: uint32(4155),
	105: uint32(4171),
	106: uint32(4187),
	107: uint32(4204),
	108: uint32(4221),
	109: uint32(4238),
	110: uint32(4257),
	111: uint32(4276),
	112: uint32(4291),
	113: uint32(4306),
	114: uint32(4325),
	115: uint32(4342),
	116: uint32(4361),
	117: uint32(4378),
	118: uint32(4390),
	119: uint32(4402),
	120: uint32(4416),
	121: uint32(4430),
	122: uint32(4442),
	123: uint32(4456),
	124: uint32(4470),
	125: uint32(4482),
	126: uint32(4496),
	127: uint32(4510),
	128: uint32(4520),
	129: uint32(4534),
	130: uint32(4546),
	131: uint32(4560),
	132: uint32(4576),
	133: uint32(4588),
	134: uint32(4604),
	135: uint32(4618),
	136: uint32(4632),
	137: uint32(4644),
	138: uint32(4656),
	139: uint32(4668),
	140: uint32(4682),
	141: uint32(4692),
	142: uint32(4706),
	143: uint32(4720),
	144: uint32(4729),
	145: uint32(4738),
	146: uint32(4751),
	147: uint32(4764),
	148: uint32(4777),
	149: uint32(4790),
	150: uint32(4803),
	151: uint32(4816),
	152: uint32(4829),
	153: uint32(4842),
	154: uint32(4855),
	155: uint32(4864),
	156: uint32(4877),
	157: uint32(4890),
	158: uint32(4903),
	159: uint32(4912),
	160: uint32(4925),
	161: uint32(4938),
	162: uint32(4947),
	163: uint32(4958),
	164: uint32(4971),
	165: uint32(4980),
	166: uint32(4989),
	167: uint32(5002),
	168: uint32(5015),
	169: uint32(5028),
	170: uint32(5037),
	171: uint32(5050),
	172: uint32(5063),
	173: uint32(5076),
	174: uint32(5089),
	175: uint32(5102),
	176: uint32(5113),
	177: uint32(5126),
	178: uint32(5139),
	179: uint32(5148),
	180: uint32(5161),
	181: uint32(5174),
	182: uint32(5187),
	183: uint32(5200),
	184: uint32(5213),
	185: uint32(5226),
	186: uint32(5239),
	187: uint32(5252),
	188: uint32(5265),
	189: uint32(5278),
	190: uint32(5291),
	191: uint32(5300),
	192: uint32(5313),
	193: uint32(5324),
	194: uint32(5337),
	195: uint32(5350),
	196: uint32(5359),
	197: uint32(5372),
	198: uint32(5381),
	199: uint32(5390),
	200: uint32(5403),
	201: uint32(5412),
	202: uint32(5425),
	203: uint32(5434),
	204: uint32(5447),
	205: uint32(5456),
	206: uint32(5469),
	207: uint32(5482),
	208: uint32(5491),
	209: uint32(5504),
	210: uint32(5517),
	211: uint32(5530),
	212: uint32(5538),
	213: uint32(5548),
	214: uint32(5558),
	215: uint32(5566),
	216: uint32(5574),
	217: uint32(5582),
	218: uint32(5590),
	219: uint32(5600),
	220: uint32(5608),
	221: uint32(5618),
	222: uint32(5628),
	223: uint32(5638),
	224: uint32(5648),
	225: uint32(5656),
	226: uint32(5664),
	227: uint32(5672),
	228: uint32(5682),
	229: uint32(5690),
	230: uint32(5698),
	231: uint32(5706),
	232: uint32(5716),
	233: uint32(5724),
	234: uint32(5731),
	235: uint32(5738),
	236: uint32(5745),
	237: uint32(5752),
	238: uint32(5759),
	239: uint32(5766),
	240: uint32(5773),
	241: uint32(5780),
	242: uint32(5787),
	243: uint32(5794),
	244: uint32(5801),
	245: uint32(5808),
	246: uint32(5815),
	247: uint32(5822),
	248: uint32(5829),
	249: uint32(5836),
	250: uint32(5843),
	251: uint32(5850),
	252: uint32(5857),
	253: uint32(5864),
	254: uint32(5871),
	255: uint32(5878),
	256: uint32(5885),
	257: uint32(5892),
	258: uint32(5899),
	259: uint32(5906),
	260: uint32(5913),
	261: uint32(5920),
	262: uint32(5927),
	263: uint32(5934),
	264: uint32(5941),
	265: uint32(5948),
	266: uint32(5955),
	267: uint32(5962),
	268: uint32(5969),
	269: uint32(5976),
	270: uint32(5983),
	271: uint32(5990),
	272: uint32(5997),
	273: uint32(6004),
	274: uint32(6011),
	275: uint32(6018),
	276: uint32(6025),
	277: uint32(6032),
	278: uint32(6039),
	279: uint32(6046),
	280: uint32(6053),
	281: uint32(6060),
	282: uint32(6067),
	283: uint32(6074),
	284: uint32(6081),
	285: uint32(6088),
	286: uint32(6095),
	287: uint32(6102),
	288: uint32(6109),
	289: uint32(6116),
	290: uint32(6123),
	291: uint32(6130),
	292: uint32(6137),
	293: uint32(6144),
	294: uint32(6151),
	295: uint32(6158),
	296: uint32(6165),
	297: uint32(6172),
	298: uint32(6179),
	299: uint32(6186),
	300: uint32(6193),
	301: uint32(6200),
	302: uint32(6207),
	303: uint32(6214),
	304: uint32(6221),
	305: uint32(6228),
	306: uint32(6235),
	307: uint32(6242),
	308: uint32(6249),
	309: uint32(6256),
	310: uint32(6263),
	311: uint32(6270),
	312: uint32(6277),
	313: uint32(6284),
	314: uint32(6291),
	315: uint32(6298),
	316: uint32(6305),
	317: uint32(6312),
	318: uint32(6319),
	319: uint32(6326),
	320: uint32(6333),
	321: uint32(6340),
	322: uint32(6347),
	323: uint32(6354),
	324: uint32(6361),
	325: uint32(6368),
	326: uint32(6375),
	327: uint32(6382),
	328: uint32(6389),
	329: uint32(6396),
	330: uint32(6403),
	331: uint32(6410),
	332: uint32(6417),
	333: uint32(6424),
	334: uint32(6431),
	335: uint32(6438),
	336: uint32(6445),
	337: uint32(6452),
	338: uint32(6459),
	339: uint32(6466),
	340: uint32(6473),
	341: uint32(6480),
	342: uint32(6487),
	343: uint32(6494),
	344: uint32(6501),
	345: uint32(6508),
	346: uint32(6515),
	347: uint32(6522),
}

var ts_parse_actions = [1099]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_source_file),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(381)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(160)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(354)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(317)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(361)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(362)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(363)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(365)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(371)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(373)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(374)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(108)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(329)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(386)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(393)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(399)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(401)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(141)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(426)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(142)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(260)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(259)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(256)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(277)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_when),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_when),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_module_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(160)),
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
		Fsymbol:      uint16(aux_sym_module_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(317)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_module_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(361)),
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
		Fsymbol:      uint16(aux_sym_module_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(362)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_module_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(363)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_module_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(365)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_module_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(371)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_module_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(373)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_module_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(374)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_module_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(108)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	98: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_repeat2),
	})))),
	99: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(329)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	100: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	101: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_repeat2),
	})))),
	102: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(386)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	104: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_repeat2),
	})))),
	105: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(393)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	107: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_repeat2),
	})))),
	108: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(53)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_module_repeat2),
	})))),
	111: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(399)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	112: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	113: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_repeat2),
	})))),
	114: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(401)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_module_repeat2),
	})))),
	117: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(141)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	118: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	119: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_repeat2),
	})))),
	120: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(426)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	122: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_repeat2),
	})))),
	123: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(142)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	124: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	125: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_repeat2),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(295)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
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
		Fcount: uint8(1),
	}})),
	131: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_repeat1),
	})))),
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
		Fsymbol:      uint16(aux_sym_module_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(354)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_repeat1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_primop),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_primop),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_expression),
		Fproduction_id: uint16(1),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_expression),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_sub_access),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_sub_access),
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
		Fcount: uint8(1),
	}})),
	154: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_sub_index),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_sub_index),
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
		Fcount: uint8(1),
	}})),
	158: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_mux),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	160: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_mux),
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
		Fcount: uint8(1),
	}})),
	162: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_uint),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_uint),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_field_id),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_field_id),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_conditionally_valid),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_conditionally_valid),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_sub_field),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_sub_field),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_primitive_operation),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_primitive_operation),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_register),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(395)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(66)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(196)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_register),
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
		Fcount: uint8(1),
	}})),
	194: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_primitive_operation),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_primitive_operation),
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
		Fcount: uint8(1),
	}})),
	198: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_literal),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_literal),
	})))),
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
		Fcount: uint8(1),
	}})),
	202: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_expression),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_expression),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_partial_connection),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_partial_connection),
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
		Fcount: uint8(1),
	}})),
	212: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_port),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_port),
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
		Fcount: uint8(1),
	}})),
	216: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(174)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_node),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_node),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_port),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_port),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_connection),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_connection),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_rdwr),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(71)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_rdwr),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(182)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_type),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(376)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_type),
	})))),
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
		Fcount: uint8(1),
	}})),
	252: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_port),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_port),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_type),
	})))),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_type),
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
		Fcount: uint8(1),
	}})),
	260: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_type),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_type),
	})))),
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
		Fcount: uint8(1),
	}})),
	264: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_wire),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(188)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_wire),
	})))),
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
		Fcount: uint8(1),
	}})),
	272: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_cmem),
	})))),
	273: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	274: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	275: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	276: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_cmem),
	})))),
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
		Fcount: uint8(1),
	}})),
	278: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_smem),
	})))),
	279: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	280: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	281: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	282: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_smem),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_type),
	})))),
	285: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	286: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_type),
	})))),
	287: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	288: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_type),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_type),
	})))),
	291: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	292: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_printf),
	})))),
	293: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_printf),
	})))),
	297: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	298: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_stop),
	})))),
	299: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(93)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_stop),
	})))),
	303: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	304: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_skip),
	})))),
	305: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_skip),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_register),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_register),
	})))),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_inst),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_inst),
	})))),
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
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_printf),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_printf),
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
		Fcount: uint8(1),
	}})),
	328: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_attach),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(89)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_attach),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_when),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_when),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_reset_block),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_reset_block),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_reset_block),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_reset_block),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_is_invalid),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_is_invalid),
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
		Fcount: uint8(1),
	}})),
	352: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_attach),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_attach),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_when),
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
		Fcount: uint8(1),
	}})),
	360: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_reset_block),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	362: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_reset_block),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_statement),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_statement),
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
		Fcount: uint8(1),
	}})),
	368: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_register),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	370: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_register),
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
		Fcount: uint8(1),
	}})),
	372: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_smem),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	374: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_smem),
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
		Fcount: uint8(1),
	}})),
	376: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_cmem),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	378: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_cmem),
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
		Fcount: uint8(1),
	}})),
	380: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_wire),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	382: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_wire),
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
		Fcount: uint8(1),
	}})),
	384: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_else),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	386: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_else),
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
		Fcount: uint8(1),
	}})),
	388: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(11),
		Fsymbol:      uint16(sym_rdwr),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	390: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(11),
		Fsymbol:      uint16(sym_rdwr),
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
		Fcount: uint8(1),
	}})),
	392: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_is_invalid),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_is_invalid),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_partial_connection),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	398: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_partial_connection),
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
		Fcount: uint8(1),
	}})),
	400: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_connection),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	402: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_connection),
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
		Fcount: uint8(1),
	}})),
	404: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_skip),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	406: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_skip),
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
		Fcount: uint8(1),
	}})),
	408: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_suite),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	410: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_suite),
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
		Fcount: uint8(1),
	}})),
	412: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_node),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	414: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_node),
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
		Fcount: uint8(1),
	}})),
	416: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_inst),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	418: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_inst),
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
		Fcount: uint8(1),
	}})),
	420: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_verif),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	422: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_verif),
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
		Fcount: uint8(1),
	}})),
	424: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_printf),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	426: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_printf),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_memory),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	430: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_memory),
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
		Fcount: uint8(1),
	}})),
	432: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_suite),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	434: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_suite),
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
		Fcount: uint8(1),
	}})),
	436: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_else),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	438: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_else),
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
		Fcount: uint8(1),
	}})),
	440: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_register),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	442: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_register),
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
		Fcount: uint8(1),
	}})),
	444: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_when),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	446: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_when),
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
		Fcount: uint8(1),
	}})),
	448: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_attach),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	450: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_attach),
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
		Fcount: uint8(1),
	}})),
	452: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_memory),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	454: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_memory),
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
		Fcount: uint8(1),
	}})),
	456: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_memory),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	458: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_memory),
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
		Fcount: uint8(1),
	}})),
	460: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_else),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	462: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_else),
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
		Fcount: uint8(1),
	}})),
	464: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_stop),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	466: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_stop),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(353)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(438)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(17)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(342)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(379)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(357)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(195)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(193)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(354)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(408)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(370)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(337)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(335)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(334)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(333)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(171)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(224)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(213)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(297)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(153)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
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
		Fsymbol:      uint16(aux_sym_memory_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(337)),
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
		Fsymbol:      uint16(aux_sym_memory_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(335)),
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
		Fsymbol:      uint16(aux_sym_memory_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(334)),
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
		Fsymbol:      uint16(aux_sym_memory_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(333)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_memory_repeat1),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(125)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_memory_field),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(232)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(231)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(253)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(157)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(34)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(151)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(152)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(198)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(144)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(293)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	566: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_memory_field),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(264)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(421)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(380)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_qualifier),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(183)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(177)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(180)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(342)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(379)),
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
		Fcount: uint8(1),
	}})),
	586: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_raw_string_repeat1),
	})))),
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
		Fcount: uint8(2),
	}})),
	588: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_raw_string_repeat1),
	})))),
	589: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(190)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(1),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(404)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	597: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat1),
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
		Fcount: uint8(2),
	}})),
	599: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(192)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(428)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(191)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(414)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(192)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(427)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(194)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(152)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(311)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(200)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(309)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(351)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(350)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(325)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	634: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_circuit_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(351)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	637: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_circuit_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(350)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_circuit_repeat1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(98)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(97)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(197)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(115)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(117)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(387)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_module),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(233)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(131)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(216)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(109)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(62)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(327)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(162)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(320)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(388)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(172)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_repeat3),
	})))),
	681: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(370)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	682: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	683: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_repeat3),
	})))),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_module),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_printf_repeat1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_circuit),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(257)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(435)),
	}})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_primitive_operation_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_number),
		Fproduction_id: uint16(2),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	705: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(381)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(186)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(356)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(313)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(437)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(442)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(318)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(249)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(173)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(187)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(136)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_source_file),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_sint),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(111)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(102)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(121)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(120)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(127)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(323)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	744: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	748: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_module),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	752: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(168)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_primitive_operation_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(98)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_module),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_circuit),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(203)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_module),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_module),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(290)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(292)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(181)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(418)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(41)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(367)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(163)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	793: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_type_repeat1),
	})))),
	794: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(168)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_type_repeat1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_field),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	800: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(273)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	804: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	808: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(375)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	812: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(108)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(307)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	828: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(372)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(228)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(226)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(114)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(145)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(56)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_module),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_printf_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(109)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(415)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(322)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(166)),
	}})))),
	858: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	859: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	860: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	861: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_memory_field_repeat1),
	})))),
	862: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(375)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_memory_field_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_string),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_defname),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_string),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_reset),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	876: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(331)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(330)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(184)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_litType),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_circuit),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_parameter),
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
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym__reset),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_reset),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_circuit),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	896: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(390)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_circuit),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(207)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(64)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(336)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(158)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(436)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	914: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(189)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(148)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(161)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(326)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(206)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(159)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(434)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(149)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(340)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(328)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(235)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_dir),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(349)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(321)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(352)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(150)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(359)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(432)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(332)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(338)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(218)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(420)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(411)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(339)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(119)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(344)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(345)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(306)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(175)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(135)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(346)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(368)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(422)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(324)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(129)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(61)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_ruw),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(134)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(219)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(123)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(288)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(302)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(299)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(366)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(122)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(377)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1030: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_mdir),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(227)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_raw_string),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(285)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(239)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(409)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(413)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1044: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	1045: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1046: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1047: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1048: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_litType),
	})))),
	1049: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1050: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(310)),
	}})))),
	1051: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1052: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(396)),
	}})))),
	1053: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1054: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(397)),
	}})))),
	1055: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1056: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1057: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1058: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1059: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1060: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1061: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1062: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1063: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1064: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(369)),
	}})))),
	1065: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1066: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(406)),
	}})))),
	1067: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1068: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(407)),
	}})))),
	1069: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1070: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1071: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1072: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(103)),
	}})))),
	1073: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1074: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_raw_string),
	})))),
	1075: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1076: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(378)),
	}})))),
	1077: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1078: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(383)),
	}})))),
	1079: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1080: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1081: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1082: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1083: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1084: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1085: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1086: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(176)),
	}})))),
	1087: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1088: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(276)),
	}})))),
	1089: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1090: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1091: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1092: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1093: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1094: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1095: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1096: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(412)),
	}})))),
	1097: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1098: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(440)),
	}})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token__newline = 0
const ts_external_token__indent = 1
const ts_external_token__dedent = 2

var ts_external_scanner_symbol_map = [3]TSSymbol{
	0: uint16(sym__newline),
	1: uint16(sym__indent),
	2: uint16(sym__dedent),
}

var ts_external_scanner_states = [6][3]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
	},
	2: {
		2: libc.BoolUint8(true1 != 0),
	},
	3: {
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
	},
	4: {
		1: libc.BoolUint8(true1 != 0),
	},
	5: {
		0: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_firrtl(tls *libc.TLS) (r uintptr) {
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
	Fkeyword_capture_token:     uint16(sym_identifier),
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
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(ts_lex_keywords)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_firrtl_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_firrtl_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_firrtl_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_firrtl_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_firrtl_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "tmp != NULL\x00combined.c\x00size == length\x00end\x00identifier\x00circuit\x00:\x00module\x00extmodule\x00input\x00output\x00const\x00UInt\x00SInt\x00Analog\x00<\x00>\x00Fixed\x00Clock\x00AsyncReset\x00Reset\x00{\x00,\x00}\x00[\x00]\x00flip\x00defname\x00=\x00parameter\x00reset\x00=>\x00(\x00)\x00wire\x00cmem\x00smem\x00reg\x00with\x00mem\x00mport\x00inst\x00of\x00node\x00<=\x00<-\x00is\x00invalid\x00data-type\x00depth\x00read-latency\x00write-latency\x00read-under-write\x00reader\x00writer\x00readwriter\x00when\x00else\x00stop\x00printf\x00assert\x00assume\x00cover\x00skip\x00attach\x00info\x00infer\x00read\x00write\x00rdwr\x00old\x00new\x00undefined\x00.\x00mux\x00validif\x00add\x00sub\x00mul\x00div\x00rem\x00lt\x00leq\x00gt\x00geq\x00eq\x00neq\x00pad\x00asUInt\x00asAsyncReset\x00asSInt\x00asClock\x00shl\x00shr\x00dshl\x00dshlw\x00dshr\x00dshrw\x00cvt\x00neg\x00not\x00and\x00or\x00xor\x00andr\x00orr\x00xorr\x00cat\x00bits\x00head\x00tail\x00asFixedPoint\x00bpshl\x00bpshr\x00bpset\x000\x00uint_token1\x00+\x00-\x00number_str\x00double\x00\"\x00'\x00string_content\x00_escape_sequence_token1\x00escape_sequence\x00relaxed_identifier\x00comment\x00_newline\x00_indent\x00_dedent\x00source_file\x00port\x00dir\x00qualifier\x00type\x00field\x00_reset\x00reset_block\x00statement\x00register\x00memory\x00connection\x00partial_connection\x00is_invalid\x00memory_field\x00suite\x00verif\x00mdir\x00ruw\x00litType\x00expression\x00literal\x00sub_field\x00sub_index\x00sub_access\x00conditionally_valid\x00primitive_operation\x00field_id\x00primop\x00uint\x00number\x00string\x00raw_string\x00_escape_sequence\x00source_file_repeat1\x00circuit_repeat1\x00module_repeat1\x00module_repeat2\x00module_repeat3\x00type_repeat1\x00memory_repeat1\x00memory_field_repeat1\x00printf_repeat1\x00primitive_operation_repeat1\x00string_repeat1\x00raw_string_repeat1\x00"
