// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-turtle/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-turtle -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-turtle/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_turtle

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
const FIELD_COUNT = 3
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
const MAX_ALIAS_SEQUENCE_LENGTH = 5
const PRODUCTION_ID_COUNT = 4
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 118
const SYMBOL_COUNT = 89
const TOKEN_COUNT = 46
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

const sym_pn_prefix = 1
const anon_sym_LBRACE = 2
const anon_sym_RBRACE = 3
const anon_sym_GRAPH = 4
const sym_comment = 5
const anon_sym_DOT = 6
const anon_sym_ATprefix = 7
const anon_sym_ATbase = 8
const aux_sym_sparql_base_token1 = 9
const aux_sym_sparql_prefix_token1 = 10
const anon_sym_SEMI = 11
const anon_sym_COMMA = 12
const anon_sym_a = 13
const anon_sym_LBRACK = 14
const anon_sym_RBRACK = 15
const anon_sym_LPAREN = 16
const anon_sym_RPAREN = 17
const anon_sym_LT = 18
const anon_sym_POUND = 19
const aux_sym_iri_reference_token1 = 20
const anon_sym_GT = 21
const sym_integer = 22
const sym_decimal = 23
const sym_double = 24
const anon_sym_DQUOTE = 25
const aux_sym__string_literal_quote_token1 = 26
const aux_sym__string_literal_quote_token2 = 27
const anon_sym_SQUOTE = 28
const aux_sym__string_literal_single_quote_token1 = 29
const anon_sym_SQUOTE_SQUOTE_SQUOTE = 30
const anon_sym_SQUOTE_SQUOTE = 31
const aux_sym__string_literal_long_single_quote_token1 = 32
const anon_sym_DQUOTE_DQUOTE_DQUOTE = 33
const anon_sym_DQUOTE_DQUOTE = 34
const aux_sym__string_literal_long_quote_token1 = 35
const anon_sym_CARET_CARET = 36
const anon_sym_true = 37
const anon_sym_false = 38
const anon_sym_COLON = 39
const anon_sym__COLON = 40
const aux_sym_blank_node_label_token1 = 41
const sym_lang_tag = 42
const sym_echar = 43
const sym_anon = 44
const sym_pn_local = 45
const sym_document = 46
const sym_graph = 47
const sym__label = 48
const sym_triple = 49
const sym_directive = 50
const sym_prefix_id = 51
const sym_base = 52
const sym_sparql_base = 53
const sym_sparql_prefix = 54
const sym__triples = 55
const sym_property_list = 56
const sym_property = 57
const sym_object_list = 58
const sym_predicate = 59
const sym_subject = 60
const sym__object = 61
const sym__literal = 62
const sym_blank_node_property_list = 63
const sym_collection = 64
const sym_object_collection = 65
const sym__numeric_literal = 66
const sym_string = 67
const sym_iri_reference = 68
const sym__string_literal_quote = 69
const sym__string_literal_single_quote = 70
const sym__string_literal_long_single_quote = 71
const sym__string_literal_long_quote = 72
const sym_rdf_literal = 73
const sym_boolean_literal = 74
const sym__iri = 75
const sym_prefixed_name = 76
const sym__blank_node = 77
const sym_namespace = 78
const sym_blank_node_label = 79
const aux_sym_document_repeat1 = 80
const aux_sym_graph_repeat1 = 81
const aux_sym_property_list_repeat1 = 82
const aux_sym_object_list_repeat1 = 83
const aux_sym_object_collection_repeat1 = 84
const aux_sym__string_literal_quote_repeat1 = 85
const aux_sym__string_literal_single_quote_repeat1 = 86
const aux_sym__string_literal_long_single_quote_repeat1 = 87
const aux_sym__string_literal_long_quote_repeat1 = 88

var ts_symbol_names = [89]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 14,
	3:  __ccgo_ts + 16,
	4:  __ccgo_ts + 18,
	5:  __ccgo_ts + 24,
	6:  __ccgo_ts + 32,
	7:  __ccgo_ts + 34,
	8:  __ccgo_ts + 42,
	9:  __ccgo_ts + 48,
	10: __ccgo_ts + 53,
	11: __ccgo_ts + 60,
	12: __ccgo_ts + 62,
	13: __ccgo_ts + 64,
	14: __ccgo_ts + 66,
	15: __ccgo_ts + 68,
	16: __ccgo_ts + 70,
	17: __ccgo_ts + 72,
	18: __ccgo_ts + 74,
	19: __ccgo_ts + 76,
	20: __ccgo_ts + 78,
	21: __ccgo_ts + 99,
	22: __ccgo_ts + 101,
	23: __ccgo_ts + 109,
	24: __ccgo_ts + 117,
	25: __ccgo_ts + 124,
	26: __ccgo_ts + 126,
	27: __ccgo_ts + 155,
	28: __ccgo_ts + 184,
	29: __ccgo_ts + 186,
	30: __ccgo_ts + 222,
	31: __ccgo_ts + 226,
	32: __ccgo_ts + 229,
	33: __ccgo_ts + 270,
	34: __ccgo_ts + 274,
	35: __ccgo_ts + 277,
	36: __ccgo_ts + 311,
	37: __ccgo_ts + 314,
	38: __ccgo_ts + 319,
	39: __ccgo_ts + 325,
	40: __ccgo_ts + 327,
	41: __ccgo_ts + 330,
	42: __ccgo_ts + 354,
	43: __ccgo_ts + 363,
	44: __ccgo_ts + 369,
	45: __ccgo_ts + 374,
	46: __ccgo_ts + 383,
	47: __ccgo_ts + 392,
	48: __ccgo_ts + 398,
	49: __ccgo_ts + 405,
	50: __ccgo_ts + 412,
	51: __ccgo_ts + 422,
	52: __ccgo_ts + 432,
	53: __ccgo_ts + 437,
	54: __ccgo_ts + 449,
	55: __ccgo_ts + 463,
	56: __ccgo_ts + 472,
	57: __ccgo_ts + 486,
	58: __ccgo_ts + 495,
	59: __ccgo_ts + 507,
	60: __ccgo_ts + 517,
	61: __ccgo_ts + 525,
	62: __ccgo_ts + 533,
	63: __ccgo_ts + 542,
	64: __ccgo_ts + 567,
	65: __ccgo_ts + 578,
	66: __ccgo_ts + 596,
	67: __ccgo_ts + 613,
	68: __ccgo_ts + 620,
	69: __ccgo_ts + 634,
	70: __ccgo_ts + 656,
	71: __ccgo_ts + 685,
	72: __ccgo_ts + 719,
	73: __ccgo_ts + 746,
	74: __ccgo_ts + 758,
	75: __ccgo_ts + 774,
	76: __ccgo_ts + 779,
	77: __ccgo_ts + 793,
	78: __ccgo_ts + 805,
	79: __ccgo_ts + 815,
	80: __ccgo_ts + 832,
	81: __ccgo_ts + 849,
	82: __ccgo_ts + 863,
	83: __ccgo_ts + 885,
	84: __ccgo_ts + 905,
	85: __ccgo_ts + 931,
	86: __ccgo_ts + 961,
	87: __ccgo_ts + 998,
	88: __ccgo_ts + 1040,
}

var ts_symbol_map = [89]TSSymbol{
	1:  uint16(sym_pn_prefix),
	2:  uint16(anon_sym_LBRACE),
	3:  uint16(anon_sym_RBRACE),
	4:  uint16(anon_sym_GRAPH),
	5:  uint16(sym_comment),
	6:  uint16(anon_sym_DOT),
	7:  uint16(anon_sym_ATprefix),
	8:  uint16(anon_sym_ATbase),
	9:  uint16(aux_sym_sparql_base_token1),
	10: uint16(aux_sym_sparql_prefix_token1),
	11: uint16(anon_sym_SEMI),
	12: uint16(anon_sym_COMMA),
	13: uint16(anon_sym_a),
	14: uint16(anon_sym_LBRACK),
	15: uint16(anon_sym_RBRACK),
	16: uint16(anon_sym_LPAREN),
	17: uint16(anon_sym_RPAREN),
	18: uint16(anon_sym_LT),
	19: uint16(anon_sym_POUND),
	20: uint16(aux_sym_iri_reference_token1),
	21: uint16(anon_sym_GT),
	22: uint16(sym_integer),
	23: uint16(sym_decimal),
	24: uint16(sym_double),
	25: uint16(anon_sym_DQUOTE),
	26: uint16(aux_sym__string_literal_quote_token1),
	27: uint16(aux_sym__string_literal_quote_token2),
	28: uint16(anon_sym_SQUOTE),
	29: uint16(aux_sym__string_literal_single_quote_token1),
	30: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	31: uint16(anon_sym_SQUOTE_SQUOTE),
	32: uint16(aux_sym__string_literal_long_single_quote_token1),
	33: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	34: uint16(anon_sym_DQUOTE_DQUOTE),
	35: uint16(aux_sym__string_literal_long_quote_token1),
	36: uint16(anon_sym_CARET_CARET),
	37: uint16(anon_sym_true),
	38: uint16(anon_sym_false),
	39: uint16(anon_sym_COLON),
	40: uint16(anon_sym__COLON),
	41: uint16(aux_sym_blank_node_label_token1),
	42: uint16(sym_lang_tag),
	43: uint16(sym_echar),
	44: uint16(sym_anon),
	45: uint16(sym_pn_local),
	46: uint16(sym_document),
	47: uint16(sym_graph),
	48: uint16(sym__label),
	49: uint16(sym_triple),
	50: uint16(sym_directive),
	51: uint16(sym_prefix_id),
	52: uint16(sym_base),
	53: uint16(sym_sparql_base),
	54: uint16(sym_sparql_prefix),
	55: uint16(sym__triples),
	56: uint16(sym_property_list),
	57: uint16(sym_property),
	58: uint16(sym_object_list),
	59: uint16(sym_predicate),
	60: uint16(sym_subject),
	61: uint16(sym__object),
	62: uint16(sym__literal),
	63: uint16(sym_blank_node_property_list),
	64: uint16(sym_collection),
	65: uint16(sym_object_collection),
	66: uint16(sym__numeric_literal),
	67: uint16(sym_string),
	68: uint16(sym_iri_reference),
	69: uint16(sym__string_literal_quote),
	70: uint16(sym__string_literal_single_quote),
	71: uint16(sym__string_literal_long_single_quote),
	72: uint16(sym__string_literal_long_quote),
	73: uint16(sym_rdf_literal),
	74: uint16(sym_boolean_literal),
	75: uint16(sym__iri),
	76: uint16(sym_prefixed_name),
	77: uint16(sym__blank_node),
	78: uint16(sym_namespace),
	79: uint16(sym_blank_node_label),
	80: uint16(aux_sym_document_repeat1),
	81: uint16(aux_sym_graph_repeat1),
	82: uint16(aux_sym_property_list_repeat1),
	83: uint16(aux_sym_object_list_repeat1),
	84: uint16(aux_sym_object_collection_repeat1),
	85: uint16(aux_sym__string_literal_quote_repeat1),
	86: uint16(aux_sym__string_literal_single_quote_repeat1),
	87: uint16(aux_sym__string_literal_long_single_quote_repeat1),
	88: uint16(aux_sym__string_literal_long_quote_repeat1),
}

var ts_symbol_metadata = [89]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	20: {},
	21: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	22: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	},
	26: {},
	27: {},
	28: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	29: {},
	30: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	31: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	32: {},
	33: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	35: {},
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
	41: {},
	42: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	43: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	44: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	45: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	46: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	51: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	61: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	62: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	70: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	71: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	72: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	76: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	77: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	78: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	79: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	80: {},
	81: {},
	82: {},
	83: {},
	84: {},
	85: {},
	86: {},
	87: {},
	88: {},
}

type ts_field_identifiers = int32

const field_datatype = 1
const field_label = 2
const field_value = 3

var ts_field_names = [4]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 1075,
	2: __ccgo_ts + 1084,
	3: __ccgo_ts + 1090,
}

var ts_field_map_slices = [4]TSFieldMapSlice{
	1: {
		Flength: uint16(1),
	},
	2: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	3: {
		Findex:  uint16(2),
		Flength: uint16(3),
	},
}

var ts_field_map_entries = [5]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_value),
	},
	1: {
		Ffield_id: uint16(field_label),
	},
	2: {
		Ffield_id:    uint16(field_datatype),
		Fchild_index: uint8(1),
	},
	3: {
		Ffield_id:    uint16(field_datatype),
		Fchild_index: uint8(2),
	},
	4: {
		Ffield_id: uint16(field_value),
	},
}

var ts_alias_sequences = [4][5]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [118]TSStateId{
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
	57:  uint16(33),
	58:  uint16(58),
	59:  uint16(59),
	60:  uint16(56),
	61:  uint16(61),
	62:  uint16(34),
	63:  uint16(63),
	64:  uint16(36),
	65:  uint16(65),
	66:  uint16(35),
	67:  uint16(35),
	68:  uint16(68),
	69:  uint16(69),
	70:  uint16(34),
	71:  uint16(71),
	72:  uint16(36),
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
	114: uint16(110),
	115: uint16(115),
	116: uint16(110),
	117: uint16(117),
}

var aux_sym_blank_node_label_token1_character_set_1 = [16]TSCharacterRange{
	0: {
		Fstart: int32('0'),
		Fend:   int32('9'),
	},
	1: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	2: {
		Fstart: int32('_'),
		Fend:   int32('_'),
	},
	3: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	4: {
		Fstart: int32(0xc0),
		Fend:   int32(0xd6),
	},
	5: {
		Fstart: int32(0xd8),
		Fend:   int32(0xf6),
	},
	6: {
		Fstart: int32(0xf8),
		Fend:   int32(0x2ff),
	},
	7: {
		Fstart: int32(0x370),
		Fend:   int32(0x37d),
	},
	8: {
		Fstart: int32(0x37f),
		Fend:   int32(0x1fff),
	},
	9: {
		Fstart: int32(0x200c),
		Fend:   int32(0x200d),
	},
	10: {
		Fstart: int32(0x2070),
		Fend:   int32(0x218f),
	},
	11: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2fef),
	},
	12: {
		Fstart: int32(0x3001),
		Fend:   int32(0xd7ff),
	},
	13: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfdcf),
	},
	14: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfffd),
	},
	15: {
		Fstart: int32(0x10000),
		Fend:   int32(0xeffff),
	},
}

var aux_sym_blank_node_label_token1_character_set_2 = [18]TSCharacterRange{
	0: {
		Fstart: int32('-'),
		Fend:   int32('.'),
	},
	1: {
		Fstart: int32('0'),
		Fend:   int32('9'),
	},
	2: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	3: {
		Fstart: int32('_'),
		Fend:   int32('_'),
	},
	4: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	5: {
		Fstart: int32(0xb7),
		Fend:   int32(0xb7),
	},
	6: {
		Fstart: int32(0xc0),
		Fend:   int32(0xd6),
	},
	7: {
		Fstart: int32(0xd8),
		Fend:   int32(0xf6),
	},
	8: {
		Fstart: int32(0xf8),
		Fend:   int32(0x37d),
	},
	9: {
		Fstart: int32(0x37f),
		Fend:   int32(0x1fff),
	},
	10: {
		Fstart: int32(0x200c),
		Fend:   int32(0x200d),
	},
	11: {
		Fstart: int32(0x203f),
		Fend:   int32(0x2040),
	},
	12: {
		Fstart: int32(0x2070),
		Fend:   int32(0x218f),
	},
	13: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2fef),
	},
	14: {
		Fstart: int32(0x3001),
		Fend:   int32(0xd7ff),
	},
	15: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfdcf),
	},
	16: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfffd),
	},
	17: {
		Fstart: int32(0x10000),
		Fend:   int32(0xeffff),
	},
}

var sym_pn_prefix_character_set_1 = [14]TSCharacterRange{
	0: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	1: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	2: {
		Fstart: int32(0xc0),
		Fend:   int32(0xd6),
	},
	3: {
		Fstart: int32(0xd8),
		Fend:   int32(0xf6),
	},
	4: {
		Fstart: int32(0xf8),
		Fend:   int32(0x2ff),
	},
	5: {
		Fstart: int32(0x370),
		Fend:   int32(0x37d),
	},
	6: {
		Fstart: int32(0x37f),
		Fend:   int32(0x1fff),
	},
	7: {
		Fstart: int32(0x200c),
		Fend:   int32(0x200d),
	},
	8: {
		Fstart: int32(0x2070),
		Fend:   int32(0x218f),
	},
	9: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2fef),
	},
	10: {
		Fstart: int32(0x3001),
		Fend:   int32(0xd7ff),
	},
	11: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfdcf),
	},
	12: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfffd),
	},
	13: {
		Fstart: int32(0x10000),
		Fend:   int32(0xeffff),
	},
}

var sym_pn_local_character_set_1 = [18]TSCharacterRange{
	0: {
		Fstart: int32('%'),
		Fend:   int32('%'),
	},
	1: {
		Fstart: int32('0'),
		Fend:   int32(':'),
	},
	2: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	3: {
		Fstart: int32('\\'),
		Fend:   int32('\\'),
	},
	4: {
		Fstart: int32('_'),
		Fend:   int32('_'),
	},
	5: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	6: {
		Fstart: int32(0xc0),
		Fend:   int32(0xd6),
	},
	7: {
		Fstart: int32(0xd8),
		Fend:   int32(0xf6),
	},
	8: {
		Fstart: int32(0xf8),
		Fend:   int32(0x2ff),
	},
	9: {
		Fstart: int32(0x370),
		Fend:   int32(0x37d),
	},
	10: {
		Fstart: int32(0x37f),
		Fend:   int32(0x1fff),
	},
	11: {
		Fstart: int32(0x200c),
		Fend:   int32(0x200d),
	},
	12: {
		Fstart: int32(0x2070),
		Fend:   int32(0x218f),
	},
	13: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2fef),
	},
	14: {
		Fstart: int32(0x3001),
		Fend:   int32(0xd7ff),
	},
	15: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfdcf),
	},
	16: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfffd),
	},
	17: {
		Fstart: int32(0x10000),
		Fend:   int32(0xeffff),
	},
}

var sym_pn_local_character_set_2 = [20]TSCharacterRange{
	0: {
		Fstart: int32('%'),
		Fend:   int32('%'),
	},
	1: {
		Fstart: int32('-'),
		Fend:   int32('.'),
	},
	2: {
		Fstart: int32('0'),
		Fend:   int32(':'),
	},
	3: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	4: {
		Fstart: int32('\\'),
		Fend:   int32('\\'),
	},
	5: {
		Fstart: int32('_'),
		Fend:   int32('_'),
	},
	6: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	7: {
		Fstart: int32(0xb7),
		Fend:   int32(0xb7),
	},
	8: {
		Fstart: int32(0xc0),
		Fend:   int32(0xd6),
	},
	9: {
		Fstart: int32(0xd8),
		Fend:   int32(0xf6),
	},
	10: {
		Fstart: int32(0xf8),
		Fend:   int32(0x37d),
	},
	11: {
		Fstart: int32(0x37f),
		Fend:   int32(0x1fff),
	},
	12: {
		Fstart: int32(0x200c),
		Fend:   int32(0x200d),
	},
	13: {
		Fstart: int32(0x203f),
		Fend:   int32(0x2040),
	},
	14: {
		Fstart: int32(0x2070),
		Fend:   int32(0x218f),
	},
	15: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2fef),
	},
	16: {
		Fstart: int32(0x3001),
		Fend:   int32(0xd7ff),
	},
	17: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfdcf),
	},
	18: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfffd),
	},
	19: {
		Fstart: int32(0x10000),
		Fend:   int32(0xeffff),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v4 uint8
	var half_size, i, i1, i10, i11, i2, i3, i4, i5, i6, i7, i8, i9, index, mid_index, size uint32_t
	var lookahead1, v3 int32_t
	var range_token, range_token1, v2 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i10, i11, i2, i3, i4, i5, i6, i7, i8, i9, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v2, v3, v4
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
			state = uint16(67)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(84)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(63)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_prefix_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(14) - index
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
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead1 == int32('"') {
			state = uint16(101)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(117)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(25)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(118)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead1 == int32('"') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead1 == int32('"') {
			state = uint16(100)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(103)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(25)
			goto next_state
		}
		if lookahead1 == int32('\n') || lookahead1 == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(104)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(4):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token1[i1]) == lookahead1 {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _6
		_6:
			;
			i1 = i1 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(135)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(18) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _10
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _10
	_10:
		if v4 != 0 {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(5):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(72)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _11
		_11:
			;
			i2 = i2 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_prefix_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(14) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _15
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _15
	_15:
		if v4 != 0 {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(6):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token3[i3]) == lookahead1 {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _16
		_16:
			;
			i3 = i3 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_prefix_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(14) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _20
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _20
	_20:
		if v4 != 0 {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(7):
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_prefix_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(14) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _25
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _25
	_25:
		if v4 != 0 {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(8):
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token5[i5]) == lookahead1 {
				state = map_token5[i5+uint32(1)]
				goto next_state
			}
			goto _26
		_26:
			;
			i5 = i5 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(18) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _30
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _30
	_30:
		if v4 != 0 {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead1 == int32('#') {
			state = uint16(79)
			goto next_state
		}
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32(':') {
			state = uint16(121)
			goto next_state
		}
		if lookahead1 == int32('<') {
			state = uint16(90)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(68)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || lookahead1 == int32('_') {
			state = uint16(134)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(18) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _34
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _34
	_34:
		if v4 != 0 {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(10):
		i6 = uint32(0)
		for {
			if !(uint64(i6) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token6[i6]) == lookahead1 {
				state = map_token6[i6+uint32(1)]
				goto next_state
			}
			goto _35
		_35:
			;
			i6 = i6 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_prefix_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(14) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _39
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _39
	_39:
		if v4 != 0 {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead1 == int32('#') {
			state = uint16(79)
			goto next_state
		}
		if lookahead1 == int32(',') {
			state = uint16(85)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(80)
			goto next_state
		}
		if lookahead1 == int32(';') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 == int32(']') {
			state = uint16(87)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(68)
			goto next_state
		}
		if lookahead1 == int32('}') {
			state = uint16(69)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead1 == int32('#') {
			state = uint16(79)
			goto next_state
		}
		if lookahead1 == int32(':') {
			state = uint16(120)
			goto next_state
		}
		if lookahead1 == int32('<') {
			state = uint16(90)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(68)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_prefix_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(14) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _43
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _43
	_43:
		if v4 != 0 {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead1 == int32('#') {
			state = uint16(79)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead1 == int32('#') {
			state = uint16(79)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(13)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_blank_node_label_token1_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(16) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _47
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _47
	_47:
		if v4 != 0 {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead1 == int32('#') {
			state = uint16(113)
			goto next_state
		}
		if lookahead1 == int32('\'') {
			state = uint16(107)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(25)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead1 == int32('#') {
			state = uint16(109)
			goto next_state
		}
		if lookahead1 == int32('\'') {
			state = uint16(106)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(25)
			goto next_state
		}
		if lookahead1 == int32('\n') || lookahead1 == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(16)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\f') || lookahead1 == int32(' ') {
			state = uint16(110)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(130)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(132)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(20) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _51
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _51
	_51:
		if v4 != 0 {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(20) - index
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
		if v4 != 0 {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(19)
			goto next_state
		}
		if lookahead1 == int32(':') {
			state = uint16(134)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(20) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _59
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _59
	_59:
		if v4 != 0 {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead1 == int32('\'') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead1 == int32('.') {
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead1 == int32('.') {
			state = uint16(22)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_blank_node_label_token1_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(18) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _63
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _63
	_63:
		if v4 != 0 {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead1 == int32('.') {
			state = uint16(23)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_blank_node_label_token1_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(18) - index
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
		if v4 != 0 {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead1 == int32(':') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(25):
		i7 = uint32(0)
		for {
			if !(uint64(i7) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token7[i7]) == lookahead1 {
				state = map_token7[i7+uint32(1)]
				goto next_state
			}
			goto _68
		_68:
			;
			i7 = i7 + uint32(2)
		}
		return result
	case int32(26):
		if lookahead1 == int32('U') {
			state = uint16(60)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead1 == int32(']') {
			state = uint16(128)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32('\n') || lookahead1 == int32('\r') || lookahead1 == int32(' ') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead1 == int32('^') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead1 == int32('a') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead1 == int32('b') {
			state = uint16(29)
			goto next_state
		}
		if lookahead1 == int32('p') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead1 == int32('e') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead1 == int32('e') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead1 == int32('f') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead1 == int32('i') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead1 == int32('r') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead1 == int32('s') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead1 == int32('x') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead1 == int32('+') || lookahead1 == int32('-') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(38)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(40):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(41):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead1 == int32('!') || int32('#') <= lookahead1 && lookahead1 <= int32('/') || lookahead1 == int32(';') || lookahead1 == int32('=') || lookahead1 == int32('?') || lookahead1 == int32('@') || lookahead1 == int32('_') || lookahead1 == int32('~') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(43):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(44):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(45):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(46):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(47):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(48):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(49):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(50):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(51):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(52):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(53):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(54):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(55):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(56):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(57):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(58):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(59):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(60):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(61):
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(62):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(63):
		if eof != 0 {
			state = uint16(67)
			goto next_state
		}
		i8 = uint32(0)
		for {
			if !(uint64(i8) < libc.Uint64FromInt64(80)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token8[i8]) == lookahead1 {
				state = map_token8[i8+uint32(1)]
				goto next_state
			}
			goto _69
		_69:
			;
			i8 = i8 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(63)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_prefix_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(14) - index
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
		if v4 != 0 {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(64):
		if eof != 0 {
			state = uint16(67)
			goto next_state
		}
		i9 = uint32(0)
		for {
			if !(uint64(i9) < libc.Uint64FromInt64(72)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token9[i9]) == lookahead1 {
				state = map_token9[i9+uint32(1)]
				goto next_state
			}
			goto _74
		_74:
			;
			i9 = i9 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(64)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_prefix_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(14) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _78
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _78
	_78:
		if v4 != 0 {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(65):
		if eof != 0 {
			state = uint16(67)
			goto next_state
		}
		i10 = uint32(0)
		for {
			if !(uint64(i10) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token10[i10]) == lookahead1 {
				state = map_token10[i10+uint32(1)]
				goto next_state
			}
			goto _79
		_79:
			;
			i10 = i10 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(66)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_prefix_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(14) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _83
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _83
	_83:
		if v4 != 0 {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(66):
		if eof != 0 {
			state = uint16(67)
			goto next_state
		}
		i11 = uint32(0)
		for {
			if !(uint64(i11) < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token11[i11]) == lookahead1 {
				state = map_token11[i11+uint32(1)]
				goto next_state
			}
			goto _84
		_84:
			;
			i11 = i11 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(66)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_prefix_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(14) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _88
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _88
	_88:
		if v4 != 0 {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('U') {
			state = uint16(78)
			goto next_state
		}
		if lookahead1 == int32('u') {
			state = uint16(74)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(95)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(71)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(72)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(73)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(74)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(75)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(76)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(77)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATprefix)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATbase)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(']') {
			state = uint16(128)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32('\n') || lookahead1 == int32('\r') || lookahead1 == int32(' ') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_iri_reference_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(91)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(26)
			goto next_state
		}
		if lookahead1 > int32(' ') && lookahead1 != int32('"') && lookahead1 != int32('#') && lookahead1 != int32('<') && lookahead1 != int32('>') && lookahead1 != int32('^') && lookahead1 != int32('`') && (lookahead1 < int32('{') || int32('}') < lookahead1) {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_iri_reference_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(95)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(26)
			goto next_state
		}
		if lookahead1 > int32(' ') && lookahead1 != int32('"') && lookahead1 != int32('#') && lookahead1 != int32('<') && lookahead1 != int32('>') && lookahead1 != int32('^') && lookahead1 != int32('`') && (lookahead1 < int32('{') || int32('}') < lookahead1) {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_iri_reference_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(26)
			goto next_state
		}
		if lookahead1 > int32(' ') && lookahead1 != int32('"') && lookahead1 != int32('<') && lookahead1 != int32('>') && lookahead1 != int32('^') && lookahead1 != int32('`') && (lookahead1 < int32('{') || int32('}') < lookahead1) {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_iri_reference_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\\') {
			state = uint16(70)
			goto next_state
		}
		if lookahead1 > int32(' ') && lookahead1 != int32('"') && lookahead1 != int32('<') && lookahead1 != int32('>') && lookahead1 != int32('^') && lookahead1 != int32('`') && (lookahead1 < int32('{') || int32('}') < lookahead1) {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(39)
			goto next_state
		}
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(38)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_decimal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(38)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_double)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('"') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('"') {
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_literal_quote_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_literal_quote_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(103)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(0x0b) || lookahead1 == int32('\f') || lookahead1 == int32(' ') {
			state = uint16(104)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32('"') && lookahead1 != int32('#') && lookahead1 != int32('\\') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_literal_quote_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\'') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\'') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_literal_single_quote_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_literal_single_quote_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(109)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(0x0b) || lookahead1 == int32('\f') || lookahead1 == int32(' ') {
			state = uint16(110)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32('\'') && lookahead1 != int32('\\') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\'') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_literal_long_single_quote_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_literal_long_single_quote_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(113)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\'') && lookahead1 != int32('\\') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('"') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_literal_long_quote_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__string_literal_long_quote_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(117)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(118)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('"') && lookahead1 != int32('#') && lookahead1 != int32('\\') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET_CARET)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(20) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _92
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _92
	_92:
		if v4 != 0 {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym__COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym__COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(20) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _96
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _96
	_96:
		if v4 != 0 {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_blank_node_label_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(23)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_blank_node_label_token1_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(18) - index
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
		if v4 != 0 {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_lang_tag)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') {
			state = uint16(62)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(126):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_lang_tag)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') {
			state = uint16(62)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_echar)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_anon)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_pn_prefix)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(22)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&aux_sym_blank_node_label_token1_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(18) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _104
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _104
	_104:
		if v4 != 0 {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_pn_local)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('+') {
			state = uint16(41)
			goto next_state
		}
		if lookahead1 == int32('-') {
			state = uint16(133)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(133)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(20) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _108
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _108
	_108:
		if v4 != 0 {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_pn_local)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32(':') {
			state = uint16(123)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(20) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _112
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _112
	_112:
		if v4 != 0 {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(132):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_pn_local)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(130)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(132)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(20) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _116
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _116
	_116:
		if v4 != 0 {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_pn_local)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(133)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(20) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _120
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _120
	_120:
		if v4 != 0 {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_pn_local)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(18)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(20) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _124
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _124
	_124:
		if v4 != 0 {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_pn_local)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(17)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		if lookahead1 == int32('E') || lookahead1 == int32('e') {
			state = uint16(130)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(135)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(20) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _128
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _128
	_128:
		if v4 != 0 {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_pn_local)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(19)
			goto next_state
		}
		if lookahead1 == int32(':') {
			state = uint16(134)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(42)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_pn_local_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(20) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _132
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _132
	_132:
		if v4 != 0 {
			state = uint16(136)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [42]uint16_t{
	0:  uint16('"'),
	1:  uint16(101),
	2:  uint16('#'),
	3:  uint16(91),
	4:  uint16('\''),
	5:  uint16(107),
	6:  uint16('('),
	7:  uint16(88),
	8:  uint16(')'),
	9:  uint16(89),
	10: uint16(','),
	11: uint16(85),
	12: uint16('.'),
	13: uint16(81),
	14: uint16(':'),
	15: uint16(120),
	16: uint16(';'),
	17: uint16(84),
	18: uint16('<'),
	19: uint16(90),
	20: uint16('>'),
	21: uint16(96),
	22: uint16('@'),
	23: uint16(30),
	24: uint16('['),
	25: uint16(86),
	26: uint16('\\'),
	27: uint16(25),
	28: uint16(']'),
	29: uint16(87),
	30: uint16('^'),
	31: uint16(28),
	32: uint16('_'),
	33: uint16(24),
	34: uint16('{'),
	35: uint16(68),
	36: uint16('}'),
	37: uint16(69),
	38: uint16('+'),
	39: uint16(21),
	40: uint16('-'),
	41: uint16(21),
}

var map_token1 = [28]uint16_t{
	0:  uint16('"'),
	1:  uint16(102),
	2:  uint16('#'),
	3:  uint16(79),
	4:  uint16('%'),
	5:  uint16(47),
	6:  uint16('\''),
	7:  uint16(108),
	8:  uint16('('),
	9:  uint16(88),
	10: uint16(')'),
	11: uint16(89),
	12: uint16('.'),
	13: uint16(40),
	14: uint16(':'),
	15: uint16(121),
	16: uint16('<'),
	17: uint16(90),
	18: uint16('['),
	19: uint16(86),
	20: uint16('\\'),
	21: uint16(42),
	22: uint16('_'),
	23: uint16(131),
	24: uint16('+'),
	25: uint16(21),
	26: uint16('-'),
	27: uint16(21),
}

var map_token2 = [36]uint16_t{
	0:  uint16('"'),
	1:  uint16(102),
	2:  uint16('#'),
	3:  uint16(79),
	4:  uint16('\''),
	5:  uint16(108),
	6:  uint16('('),
	7:  uint16(88),
	8:  uint16(')'),
	9:  uint16(89),
	10: uint16(','),
	11: uint16(85),
	12: uint16('.'),
	13: uint16(81),
	14: uint16(':'),
	15: uint16(120),
	16: uint16(';'),
	17: uint16(84),
	18: uint16('<'),
	19: uint16(90),
	20: uint16('@'),
	21: uint16(61),
	22: uint16('['),
	23: uint16(86),
	24: uint16(']'),
	25: uint16(87),
	26: uint16('^'),
	27: uint16(28),
	28: uint16('_'),
	29: uint16(24),
	30: uint16('}'),
	31: uint16(69),
	32: uint16('+'),
	33: uint16(21),
	34: uint16('-'),
	35: uint16(21),
}

var map_token3 = [28]uint16_t{
	0:  uint16('"'),
	1:  uint16(102),
	2:  uint16('#'),
	3:  uint16(79),
	4:  uint16('\''),
	5:  uint16(108),
	6:  uint16('('),
	7:  uint16(88),
	8:  uint16(')'),
	9:  uint16(89),
	10: uint16('.'),
	11: uint16(40),
	12: uint16(':'),
	13: uint16(120),
	14: uint16('<'),
	15: uint16(90),
	16: uint16('@'),
	17: uint16(61),
	18: uint16('['),
	19: uint16(86),
	20: uint16('^'),
	21: uint16(28),
	22: uint16('_'),
	23: uint16(24),
	24: uint16('+'),
	25: uint16(21),
	26: uint16('-'),
	27: uint16(21),
}

var map_token4 = [24]uint16_t{
	0:  uint16('"'),
	1:  uint16(102),
	2:  uint16('#'),
	3:  uint16(79),
	4:  uint16('\''),
	5:  uint16(108),
	6:  uint16('('),
	7:  uint16(88),
	8:  uint16(')'),
	9:  uint16(89),
	10: uint16('.'),
	11: uint16(40),
	12: uint16(':'),
	13: uint16(120),
	14: uint16('<'),
	15: uint16(90),
	16: uint16('['),
	17: uint16(86),
	18: uint16('_'),
	19: uint16(24),
	20: uint16('+'),
	21: uint16(21),
	22: uint16('-'),
	23: uint16(21),
}

var map_token5 = [18]uint16_t{
	0:  uint16('#'),
	1:  uint16(79),
	2:  uint16('%'),
	3:  uint16(47),
	4:  uint16(','),
	5:  uint16(85),
	6:  uint16('.'),
	7:  uint16(80),
	8:  uint16(';'),
	9:  uint16(84),
	10: uint16('\\'),
	11: uint16(42),
	12: uint16(']'),
	13: uint16(87),
	14: uint16('{'),
	15: uint16(68),
	16: uint16('}'),
	17: uint16(69),
}

var map_token6 = [20]uint16_t{
	0:  uint16('#'),
	1:  uint16(79),
	2:  uint16(','),
	3:  uint16(85),
	4:  uint16('.'),
	5:  uint16(80),
	6:  uint16(':'),
	7:  uint16(120),
	8:  uint16(';'),
	9:  uint16(84),
	10: uint16('<'),
	11: uint16(90),
	12: uint16('@'),
	13: uint16(61),
	14: uint16(']'),
	15: uint16(87),
	16: uint16('^'),
	17: uint16(28),
	18: uint16('}'),
	19: uint16(69),
}

var map_token7 = [20]uint16_t{
	0:  uint16('U'),
	1:  uint16(59),
	2:  uint16('u'),
	3:  uint16(51),
	4:  uint16('"'),
	5:  uint16(127),
	6:  uint16('\''),
	7:  uint16(127),
	8:  uint16('\\'),
	9:  uint16(127),
	10: uint16('b'),
	11: uint16(127),
	12: uint16('f'),
	13: uint16(127),
	14: uint16('n'),
	15: uint16(127),
	16: uint16('r'),
	17: uint16(127),
	18: uint16('t'),
	19: uint16(127),
}

var map_token8 = [40]uint16_t{
	0:  uint16('"'),
	1:  uint16(101),
	2:  uint16('#'),
	3:  uint16(91),
	4:  uint16('\''),
	5:  uint16(107),
	6:  uint16('('),
	7:  uint16(88),
	8:  uint16(')'),
	9:  uint16(89),
	10: uint16(','),
	11: uint16(85),
	12: uint16('.'),
	13: uint16(81),
	14: uint16(':'),
	15: uint16(120),
	16: uint16(';'),
	17: uint16(84),
	18: uint16('<'),
	19: uint16(90),
	20: uint16('@'),
	21: uint16(30),
	22: uint16('['),
	23: uint16(86),
	24: uint16('\\'),
	25: uint16(25),
	26: uint16(']'),
	27: uint16(87),
	28: uint16('^'),
	29: uint16(28),
	30: uint16('_'),
	31: uint16(24),
	32: uint16('{'),
	33: uint16(68),
	34: uint16('}'),
	35: uint16(69),
	36: uint16('+'),
	37: uint16(21),
	38: uint16('-'),
	39: uint16(21),
}

var map_token9 = [36]uint16_t{
	0:  uint16('"'),
	1:  uint16(102),
	2:  uint16('#'),
	3:  uint16(79),
	4:  uint16('\''),
	5:  uint16(108),
	6:  uint16('('),
	7:  uint16(88),
	8:  uint16(')'),
	9:  uint16(89),
	10: uint16(','),
	11: uint16(85),
	12: uint16('.'),
	13: uint16(81),
	14: uint16(':'),
	15: uint16(120),
	16: uint16(';'),
	17: uint16(84),
	18: uint16('<'),
	19: uint16(90),
	20: uint16('@'),
	21: uint16(30),
	22: uint16('['),
	23: uint16(86),
	24: uint16(']'),
	25: uint16(87),
	26: uint16('_'),
	27: uint16(24),
	28: uint16('{'),
	29: uint16(68),
	30: uint16('}'),
	31: uint16(69),
	32: uint16('+'),
	33: uint16(21),
	34: uint16('-'),
	35: uint16(21),
}

var map_token10 = [34]uint16_t{
	0:  uint16('"'),
	1:  uint16(102),
	2:  uint16('#'),
	3:  uint16(79),
	4:  uint16('\''),
	5:  uint16(108),
	6:  uint16('('),
	7:  uint16(88),
	8:  uint16(')'),
	9:  uint16(89),
	10: uint16('.'),
	11: uint16(40),
	12: uint16(':'),
	13: uint16(120),
	14: uint16('<'),
	15: uint16(90),
	16: uint16('>'),
	17: uint16(96),
	18: uint16('@'),
	19: uint16(30),
	20: uint16('['),
	21: uint16(86),
	22: uint16(']'),
	23: uint16(87),
	24: uint16('_'),
	25: uint16(24),
	26: uint16('{'),
	27: uint16(68),
	28: uint16('}'),
	29: uint16(69),
	30: uint16('+'),
	31: uint16(21),
	32: uint16('-'),
	33: uint16(21),
}

var map_token11 = [32]uint16_t{
	0:  uint16('"'),
	1:  uint16(102),
	2:  uint16('#'),
	3:  uint16(79),
	4:  uint16('\''),
	5:  uint16(108),
	6:  uint16('('),
	7:  uint16(88),
	8:  uint16(')'),
	9:  uint16(89),
	10: uint16('.'),
	11: uint16(40),
	12: uint16(':'),
	13: uint16(120),
	14: uint16('<'),
	15: uint16(90),
	16: uint16('@'),
	17: uint16(30),
	18: uint16('['),
	19: uint16(86),
	20: uint16(']'),
	21: uint16(87),
	22: uint16('_'),
	23: uint16(24),
	24: uint16('{'),
	25: uint16(68),
	26: uint16('}'),
	27: uint16(69),
	28: uint16('+'),
	29: uint16(21),
	30: uint16('-'),
	31: uint16(21),
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
			if !(uint64(i) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token12[i]) == lookahead {
				state = map_token12[i+uint32(1)]
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
		if lookahead == int32('R') {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(2):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_a)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(3):
		if lookahead == int32('a') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('r') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('A') || lookahead == int32('a') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('R') || lookahead == int32('r') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('A') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('l') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('u') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('S') || lookahead == int32('s') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('P') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('s') {
			state = uint16(18)
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
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('F') || lookahead == int32('f') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('H') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('e') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(19):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(20):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sparql_base_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(21):
		if lookahead == int32('I') || lookahead == int32('i') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(22):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GRAPH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(23):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(24):
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_sparql_prefix_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token12 = [16]uint16_t{
	0:  uint16('G'),
	1:  uint16(1),
	2:  uint16('a'),
	3:  uint16(2),
	4:  uint16('f'),
	5:  uint16(3),
	6:  uint16('t'),
	7:  uint16(4),
	8:  uint16('B'),
	9:  uint16(5),
	10: uint16('b'),
	11: uint16(5),
	12: uint16('P'),
	13: uint16(6),
	14: uint16('p'),
	15: uint16(6),
}

var ts_lex_modes = [118]TSLexMode{
	0: {},
	1: {
		Flex_state: uint16(65),
	},
	2: {
		Flex_state: uint16(65),
	},
	3: {
		Flex_state: uint16(65),
	},
	4: {
		Flex_state: uint16(65),
	},
	5: {
		Flex_state: uint16(65),
	},
	6: {
		Flex_state: uint16(65),
	},
	7: {
		Flex_state: uint16(65),
	},
	8: {
		Flex_state: uint16(65),
	},
	9: {
		Flex_state: uint16(64),
	},
	10: {
		Flex_state: uint16(64),
	},
	11: {
		Flex_state: uint16(64),
	},
	12: {
		Flex_state: uint16(5),
	},
	13: {
		Flex_state: uint16(5),
	},
	14: {
		Flex_state: uint16(5),
	},
	15: {
		Flex_state: uint16(5),
	},
	16: {
		Flex_state: uint16(5),
	},
	17: {
		Flex_state: uint16(64),
	},
	18: {
		Flex_state: uint16(5),
	},
	19: {
		Flex_state: uint16(5),
	},
	20: {
		Flex_state: uint16(5),
	},
	21: {
		Flex_state: uint16(5),
	},
	22: {
		Flex_state: uint16(64),
	},
	23: {
		Flex_state: uint16(64),
	},
	24: {
		Flex_state: uint16(64),
	},
	25: {
		Flex_state: uint16(64),
	},
	26: {
		Flex_state: uint16(64),
	},
	27: {
		Flex_state: uint16(64),
	},
	28: {
		Flex_state: uint16(65),
	},
	29: {
		Flex_state: uint16(65),
	},
	30: {
		Flex_state: uint16(65),
	},
	31: {
		Flex_state: uint16(65),
	},
	32: {
		Flex_state: uint16(65),
	},
	33: {
		Flex_state: uint16(6),
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
		Flex_state: uint16(65),
	},
	38: {
		Flex_state: uint16(65),
	},
	39: {
		Flex_state: uint16(10),
	},
	40: {
		Flex_state: uint16(65),
	},
	41: {
		Flex_state: uint16(65),
	},
	42: {
		Flex_state: uint16(65),
	},
	43: {
		Flex_state: uint16(65),
	},
	44: {
		Flex_state: uint16(65),
	},
	45: {
		Flex_state: uint16(65),
	},
	46: {
		Flex_state: uint16(65),
	},
	47: {
		Flex_state: uint16(65),
	},
	48: {
		Flex_state: uint16(65),
	},
	49: {
		Flex_state: uint16(65),
	},
	50: {
		Flex_state: uint16(65),
	},
	51: {
		Flex_state: uint16(10),
	},
	52: {
		Flex_state: uint16(65),
	},
	53: {
		Flex_state: uint16(65),
	},
	54: {
		Flex_state: uint16(65),
	},
	55: {
		Flex_state: uint16(1),
	},
	56: {
		Flex_state: uint16(65),
	},
	57: {
		Flex_state: uint16(10),
	},
	58: {
		Flex_state: uint16(15),
	},
	59: {
		Flex_state: uint16(1),
	},
	60: {
		Flex_state: uint16(65),
	},
	61: {
		Flex_state: uint16(1),
	},
	62: {
		Flex_state: uint16(8),
	},
	63: {
		Flex_state: uint16(15),
	},
	64: {
		Flex_state: uint16(8),
	},
	65: {
		Flex_state: uint16(15),
	},
	66: {
		Flex_state: uint16(8),
	},
	67: {
		Flex_state: uint16(9),
	},
	68: {
		Flex_state: uint16(15),
	},
	69: {
		Flex_state: uint16(10),
	},
	70: {
		Flex_state: uint16(9),
	},
	71: {
		Flex_state: uint16(10),
	},
	72: {
		Flex_state: uint16(9),
	},
	73: {
		Flex_state: uint16(10),
	},
	74: {
		Flex_state: uint16(1),
	},
	75: {
		Flex_state: uint16(10),
	},
	76: {
		Flex_state: uint16(3),
	},
	77: {
		Flex_state: uint16(10),
	},
	78: {
		Flex_state: uint16(3),
	},
	79: {
		Flex_state: uint16(65),
	},
	80: {
		Flex_state: uint16(10),
	},
	81: {
		Flex_state: uint16(10),
	},
	82: {
		Flex_state: uint16(16),
	},
	83: {
		Flex_state: uint16(3),
	},
	84: {
		Flex_state: uint16(16),
	},
	85: {
		Flex_state: uint16(16),
	},
	86: {
		Flex_state: uint16(65),
	},
	87: {
		Flex_state: uint16(10),
	},
	88: {
		Flex_state: uint16(10),
	},
	89: {
		Flex_state: uint16(65),
	},
	90: {
		Flex_state: uint16(65),
	},
	91: {
		Flex_state: uint16(15),
	},
	92: {
		Flex_state: uint16(1),
	},
	93: {
		Flex_state: uint16(10),
	},
	94: {
		Flex_state: uint16(10),
	},
	95: {
		Flex_state: uint16(65),
	},
	96: {
		Flex_state: uint16(65),
	},
	97: {
		Flex_state: uint16(10),
	},
	98: {
		Flex_state: uint16(92),
	},
	99: {
		Flex_state: uint16(10),
	},
	100: {
		Flex_state: uint16(65),
	},
	101: {
		Flex_state: uint16(65),
	},
	102: {
		Flex_state: uint16(10),
	},
	103: {
		Flex_state: uint16(65),
	},
	104: {
		Flex_state: uint16(65),
	},
	105: {
		Flex_state: uint16(65),
	},
	106: {
		Flex_state: uint16(10),
	},
	107: {
		Flex_state: uint16(10),
	},
	108: {
		Flex_state: uint16(65),
	},
	109: {
		Flex_state: uint16(65),
	},
	110: {
		Flex_state: uint16(65),
	},
	111: {
		Flex_state: uint16(14),
	},
	112: {
		Flex_state: uint16(65),
	},
	113: {
		Flex_state: uint16(10),
	},
	114: {
		Flex_state: uint16(65),
	},
	115: {
		Flex_state: uint16(65),
	},
	116: {
		Flex_state: uint16(65),
	},
	117: {
		Flex_state: uint16(93),
	},
}

var ts_parse_table = [2][89]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		3:  uint16(1),
		4:  uint16(1),
		5:  uint16(3),
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
		21: uint16(1),
		22: uint16(1),
		23: uint16(1),
		24: uint16(1),
		25: uint16(1),
		27: uint16(1),
		28: uint16(1),
		30: uint16(1),
		31: uint16(1),
		33: uint16(1),
		34: uint16(1),
		36: uint16(1),
		37: uint16(1),
		38: uint16(1),
		39: uint16(1),
		40: uint16(1),
		43: uint16(1),
		44: uint16(1),
	},
	1: {
		0:  uint16(5),
		1:  uint16(7),
		2:  uint16(9),
		4:  uint16(11),
		5:  uint16(13),
		7:  uint16(15),
		8:  uint16(17),
		9:  uint16(19),
		10: uint16(21),
		14: uint16(23),
		16: uint16(25),
		18: uint16(27),
		39: uint16(29),
		40: uint16(31),
		44: uint16(33),
		46: uint16(109),
		47: uint16(7),
		48: uint16(108),
		49: uint16(7),
		50: uint16(7),
		51: uint16(40),
		52: uint16(40),
		53: uint16(40),
		54: uint16(40),
		55: uint16(107),
		60: uint16(53),
		63: uint16(51),
		64: uint16(86),
		68: uint16(79),
		75: uint16(79),
		76: uint16(79),
		77: uint16(79),
		78: uint16(67),
		79: uint16(79),
		80: uint16(7),
	},
}

var ts_small_parse_table = [2785]uint16_t{
	0:    uint16(20),
	1:    uint16(13),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(23),
	5:    uint16(1),
	6:    uint16(anon_sym_LBRACK),
	7:    uint16(25),
	8:    uint16(1),
	9:    uint16(anon_sym_LPAREN),
	10:   uint16(27),
	11:   uint16(1),
	12:   uint16(anon_sym_LT),
	13:   uint16(31),
	14:   uint16(1),
	15:   uint16(anon_sym__COLON),
	16:   uint16(35),
	17:   uint16(1),
	18:   uint16(sym_pn_prefix),
	19:   uint16(37),
	20:   uint16(1),
	21:   uint16(anon_sym_RPAREN),
	22:   uint16(43),
	23:   uint16(1),
	24:   uint16(anon_sym_DQUOTE),
	25:   uint16(45),
	26:   uint16(1),
	27:   uint16(anon_sym_SQUOTE),
	28:   uint16(47),
	29:   uint16(1),
	30:   uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	31:   uint16(49),
	32:   uint16(1),
	33:   uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	34:   uint16(53),
	35:   uint16(1),
	36:   uint16(anon_sym_COLON),
	37:   uint16(33),
	38:   uint16(1),
	39:   uint16(sym_string),
	40:   uint16(35),
	41:   uint16(1),
	42:   uint16(sym_namespace),
	43:   uint16(115),
	44:   uint16(1),
	45:   uint16(sym_object_collection),
	46:   uint16(39),
	47:   uint16(2),
	48:   uint16(sym_integer),
	49:   uint16(sym_decimal),
	50:   uint16(41),
	51:   uint16(2),
	52:   uint16(sym_double),
	53:   uint16(sym_anon),
	54:   uint16(51),
	55:   uint16(2),
	56:   uint16(anon_sym_true),
	57:   uint16(anon_sym_false),
	58:   uint16(18),
	59:   uint16(4),
	60:   uint16(sym__string_literal_quote),
	61:   uint16(sym__string_literal_single_quote),
	62:   uint16(sym__string_literal_long_single_quote),
	63:   uint16(sym__string_literal_long_quote),
	64:   uint16(4),
	65:   uint16(13),
	66:   uint16(sym__object),
	67:   uint16(sym__literal),
	68:   uint16(sym_blank_node_property_list),
	69:   uint16(sym_collection),
	70:   uint16(sym__numeric_literal),
	71:   uint16(sym_iri_reference),
	72:   uint16(sym_rdf_literal),
	73:   uint16(sym_boolean_literal),
	74:   uint16(sym__iri),
	75:   uint16(sym_prefixed_name),
	76:   uint16(sym__blank_node),
	77:   uint16(sym_blank_node_label),
	78:   uint16(aux_sym_object_collection_repeat1),
	79:   uint16(19),
	80:   uint16(13),
	81:   uint16(1),
	82:   uint16(sym_comment),
	83:   uint16(55),
	84:   uint16(1),
	85:   uint16(sym_pn_prefix),
	86:   uint16(58),
	87:   uint16(1),
	88:   uint16(anon_sym_LBRACK),
	89:   uint16(61),
	90:   uint16(1),
	91:   uint16(anon_sym_LPAREN),
	92:   uint16(64),
	93:   uint16(1),
	94:   uint16(anon_sym_RPAREN),
	95:   uint16(66),
	96:   uint16(1),
	97:   uint16(anon_sym_LT),
	98:   uint16(75),
	99:   uint16(1),
	100:  uint16(anon_sym_DQUOTE),
	101:  uint16(78),
	102:  uint16(1),
	103:  uint16(anon_sym_SQUOTE),
	104:  uint16(81),
	105:  uint16(1),
	106:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	107:  uint16(84),
	108:  uint16(1),
	109:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	110:  uint16(90),
	111:  uint16(1),
	112:  uint16(anon_sym_COLON),
	113:  uint16(93),
	114:  uint16(1),
	115:  uint16(anon_sym__COLON),
	116:  uint16(33),
	117:  uint16(1),
	118:  uint16(sym_string),
	119:  uint16(35),
	120:  uint16(1),
	121:  uint16(sym_namespace),
	122:  uint16(69),
	123:  uint16(2),
	124:  uint16(sym_integer),
	125:  uint16(sym_decimal),
	126:  uint16(72),
	127:  uint16(2),
	128:  uint16(sym_double),
	129:  uint16(sym_anon),
	130:  uint16(87),
	131:  uint16(2),
	132:  uint16(anon_sym_true),
	133:  uint16(anon_sym_false),
	134:  uint16(18),
	135:  uint16(4),
	136:  uint16(sym__string_literal_quote),
	137:  uint16(sym__string_literal_single_quote),
	138:  uint16(sym__string_literal_long_single_quote),
	139:  uint16(sym__string_literal_long_quote),
	140:  uint16(3),
	141:  uint16(13),
	142:  uint16(sym__object),
	143:  uint16(sym__literal),
	144:  uint16(sym_blank_node_property_list),
	145:  uint16(sym_collection),
	146:  uint16(sym__numeric_literal),
	147:  uint16(sym_iri_reference),
	148:  uint16(sym_rdf_literal),
	149:  uint16(sym_boolean_literal),
	150:  uint16(sym__iri),
	151:  uint16(sym_prefixed_name),
	152:  uint16(sym__blank_node),
	153:  uint16(sym_blank_node_label),
	154:  uint16(aux_sym_object_collection_repeat1),
	155:  uint16(19),
	156:  uint16(13),
	157:  uint16(1),
	158:  uint16(sym_comment),
	159:  uint16(23),
	160:  uint16(1),
	161:  uint16(anon_sym_LBRACK),
	162:  uint16(25),
	163:  uint16(1),
	164:  uint16(anon_sym_LPAREN),
	165:  uint16(27),
	166:  uint16(1),
	167:  uint16(anon_sym_LT),
	168:  uint16(31),
	169:  uint16(1),
	170:  uint16(anon_sym__COLON),
	171:  uint16(35),
	172:  uint16(1),
	173:  uint16(sym_pn_prefix),
	174:  uint16(43),
	175:  uint16(1),
	176:  uint16(anon_sym_DQUOTE),
	177:  uint16(45),
	178:  uint16(1),
	179:  uint16(anon_sym_SQUOTE),
	180:  uint16(47),
	181:  uint16(1),
	182:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	183:  uint16(49),
	184:  uint16(1),
	185:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	186:  uint16(53),
	187:  uint16(1),
	188:  uint16(anon_sym_COLON),
	189:  uint16(96),
	190:  uint16(1),
	191:  uint16(anon_sym_RPAREN),
	192:  uint16(33),
	193:  uint16(1),
	194:  uint16(sym_string),
	195:  uint16(35),
	196:  uint16(1),
	197:  uint16(sym_namespace),
	198:  uint16(51),
	199:  uint16(2),
	200:  uint16(anon_sym_true),
	201:  uint16(anon_sym_false),
	202:  uint16(98),
	203:  uint16(2),
	204:  uint16(sym_integer),
	205:  uint16(sym_decimal),
	206:  uint16(100),
	207:  uint16(2),
	208:  uint16(sym_double),
	209:  uint16(sym_anon),
	210:  uint16(18),
	211:  uint16(4),
	212:  uint16(sym__string_literal_quote),
	213:  uint16(sym__string_literal_single_quote),
	214:  uint16(sym__string_literal_long_single_quote),
	215:  uint16(sym__string_literal_long_quote),
	216:  uint16(3),
	217:  uint16(13),
	218:  uint16(sym__object),
	219:  uint16(sym__literal),
	220:  uint16(sym_blank_node_property_list),
	221:  uint16(sym_collection),
	222:  uint16(sym__numeric_literal),
	223:  uint16(sym_iri_reference),
	224:  uint16(sym_rdf_literal),
	225:  uint16(sym_boolean_literal),
	226:  uint16(sym__iri),
	227:  uint16(sym_prefixed_name),
	228:  uint16(sym__blank_node),
	229:  uint16(sym_blank_node_label),
	230:  uint16(aux_sym_object_collection_repeat1),
	231:  uint16(19),
	232:  uint16(13),
	233:  uint16(1),
	234:  uint16(sym_comment),
	235:  uint16(23),
	236:  uint16(1),
	237:  uint16(anon_sym_LBRACK),
	238:  uint16(25),
	239:  uint16(1),
	240:  uint16(anon_sym_LPAREN),
	241:  uint16(27),
	242:  uint16(1),
	243:  uint16(anon_sym_LT),
	244:  uint16(31),
	245:  uint16(1),
	246:  uint16(anon_sym__COLON),
	247:  uint16(43),
	248:  uint16(1),
	249:  uint16(anon_sym_DQUOTE),
	250:  uint16(45),
	251:  uint16(1),
	252:  uint16(anon_sym_SQUOTE),
	253:  uint16(47),
	254:  uint16(1),
	255:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	256:  uint16(49),
	257:  uint16(1),
	258:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	259:  uint16(102),
	260:  uint16(1),
	261:  uint16(sym_pn_prefix),
	262:  uint16(108),
	263:  uint16(1),
	264:  uint16(anon_sym_COLON),
	265:  uint16(57),
	266:  uint16(1),
	267:  uint16(sym_string),
	268:  uint16(66),
	269:  uint16(1),
	270:  uint16(sym_namespace),
	271:  uint16(87),
	272:  uint16(1),
	273:  uint16(sym_object_list),
	274:  uint16(51),
	275:  uint16(2),
	276:  uint16(anon_sym_true),
	277:  uint16(anon_sym_false),
	278:  uint16(104),
	279:  uint16(2),
	280:  uint16(sym_integer),
	281:  uint16(sym_decimal),
	282:  uint16(106),
	283:  uint16(2),
	284:  uint16(sym_double),
	285:  uint16(sym_anon),
	286:  uint16(18),
	287:  uint16(4),
	288:  uint16(sym__string_literal_quote),
	289:  uint16(sym__string_literal_single_quote),
	290:  uint16(sym__string_literal_long_single_quote),
	291:  uint16(sym__string_literal_long_quote),
	292:  uint16(69),
	293:  uint16(12),
	294:  uint16(sym__object),
	295:  uint16(sym__literal),
	296:  uint16(sym_blank_node_property_list),
	297:  uint16(sym_collection),
	298:  uint16(sym__numeric_literal),
	299:  uint16(sym_iri_reference),
	300:  uint16(sym_rdf_literal),
	301:  uint16(sym_boolean_literal),
	302:  uint16(sym__iri),
	303:  uint16(sym_prefixed_name),
	304:  uint16(sym__blank_node),
	305:  uint16(sym_blank_node_label),
	306:  uint16(18),
	307:  uint16(13),
	308:  uint16(1),
	309:  uint16(sym_comment),
	310:  uint16(23),
	311:  uint16(1),
	312:  uint16(anon_sym_LBRACK),
	313:  uint16(25),
	314:  uint16(1),
	315:  uint16(anon_sym_LPAREN),
	316:  uint16(27),
	317:  uint16(1),
	318:  uint16(anon_sym_LT),
	319:  uint16(31),
	320:  uint16(1),
	321:  uint16(anon_sym__COLON),
	322:  uint16(43),
	323:  uint16(1),
	324:  uint16(anon_sym_DQUOTE),
	325:  uint16(45),
	326:  uint16(1),
	327:  uint16(anon_sym_SQUOTE),
	328:  uint16(47),
	329:  uint16(1),
	330:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	331:  uint16(49),
	332:  uint16(1),
	333:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	334:  uint16(102),
	335:  uint16(1),
	336:  uint16(sym_pn_prefix),
	337:  uint16(108),
	338:  uint16(1),
	339:  uint16(anon_sym_COLON),
	340:  uint16(57),
	341:  uint16(1),
	342:  uint16(sym_string),
	343:  uint16(66),
	344:  uint16(1),
	345:  uint16(sym_namespace),
	346:  uint16(51),
	347:  uint16(2),
	348:  uint16(anon_sym_true),
	349:  uint16(anon_sym_false),
	350:  uint16(110),
	351:  uint16(2),
	352:  uint16(sym_integer),
	353:  uint16(sym_decimal),
	354:  uint16(112),
	355:  uint16(2),
	356:  uint16(sym_double),
	357:  uint16(sym_anon),
	358:  uint16(18),
	359:  uint16(4),
	360:  uint16(sym__string_literal_quote),
	361:  uint16(sym__string_literal_single_quote),
	362:  uint16(sym__string_literal_long_single_quote),
	363:  uint16(sym__string_literal_long_quote),
	364:  uint16(77),
	365:  uint16(12),
	366:  uint16(sym__object),
	367:  uint16(sym__literal),
	368:  uint16(sym_blank_node_property_list),
	369:  uint16(sym_collection),
	370:  uint16(sym__numeric_literal),
	371:  uint16(sym_iri_reference),
	372:  uint16(sym_rdf_literal),
	373:  uint16(sym_boolean_literal),
	374:  uint16(sym__iri),
	375:  uint16(sym_prefixed_name),
	376:  uint16(sym__blank_node),
	377:  uint16(sym_blank_node_label),
	378:  uint16(24),
	379:  uint16(7),
	380:  uint16(1),
	381:  uint16(sym_pn_prefix),
	382:  uint16(9),
	383:  uint16(1),
	384:  uint16(anon_sym_LBRACE),
	385:  uint16(11),
	386:  uint16(1),
	387:  uint16(anon_sym_GRAPH),
	388:  uint16(13),
	389:  uint16(1),
	390:  uint16(sym_comment),
	391:  uint16(15),
	392:  uint16(1),
	393:  uint16(anon_sym_ATprefix),
	394:  uint16(17),
	395:  uint16(1),
	396:  uint16(anon_sym_ATbase),
	397:  uint16(19),
	398:  uint16(1),
	399:  uint16(aux_sym_sparql_base_token1),
	400:  uint16(21),
	401:  uint16(1),
	402:  uint16(aux_sym_sparql_prefix_token1),
	403:  uint16(23),
	404:  uint16(1),
	405:  uint16(anon_sym_LBRACK),
	406:  uint16(25),
	407:  uint16(1),
	408:  uint16(anon_sym_LPAREN),
	409:  uint16(27),
	410:  uint16(1),
	411:  uint16(anon_sym_LT),
	412:  uint16(29),
	413:  uint16(1),
	414:  uint16(anon_sym_COLON),
	415:  uint16(31),
	416:  uint16(1),
	417:  uint16(anon_sym__COLON),
	418:  uint16(33),
	419:  uint16(1),
	420:  uint16(sym_anon),
	421:  uint16(114),
	422:  uint16(1),
	424:  uint16(51),
	425:  uint16(1),
	426:  uint16(sym_blank_node_property_list),
	427:  uint16(53),
	428:  uint16(1),
	429:  uint16(sym_subject),
	430:  uint16(67),
	431:  uint16(1),
	432:  uint16(sym_namespace),
	433:  uint16(86),
	434:  uint16(1),
	435:  uint16(sym_collection),
	436:  uint16(107),
	437:  uint16(1),
	438:  uint16(sym__triples),
	439:  uint16(108),
	440:  uint16(1),
	441:  uint16(sym__label),
	442:  uint16(8),
	443:  uint16(4),
	444:  uint16(sym_graph),
	445:  uint16(sym_triple),
	446:  uint16(sym_directive),
	447:  uint16(aux_sym_document_repeat1),
	448:  uint16(40),
	449:  uint16(4),
	450:  uint16(sym_prefix_id),
	451:  uint16(sym_base),
	452:  uint16(sym_sparql_base),
	453:  uint16(sym_sparql_prefix),
	454:  uint16(79),
	455:  uint16(5),
	456:  uint16(sym_iri_reference),
	457:  uint16(sym__iri),
	458:  uint16(sym_prefixed_name),
	459:  uint16(sym__blank_node),
	460:  uint16(sym_blank_node_label),
	461:  uint16(24),
	462:  uint16(13),
	463:  uint16(1),
	464:  uint16(sym_comment),
	465:  uint16(116),
	466:  uint16(1),
	468:  uint16(118),
	469:  uint16(1),
	470:  uint16(sym_pn_prefix),
	471:  uint16(121),
	472:  uint16(1),
	473:  uint16(anon_sym_LBRACE),
	474:  uint16(124),
	475:  uint16(1),
	476:  uint16(anon_sym_GRAPH),
	477:  uint16(127),
	478:  uint16(1),
	479:  uint16(anon_sym_ATprefix),
	480:  uint16(130),
	481:  uint16(1),
	482:  uint16(anon_sym_ATbase),
	483:  uint16(133),
	484:  uint16(1),
	485:  uint16(aux_sym_sparql_base_token1),
	486:  uint16(136),
	487:  uint16(1),
	488:  uint16(aux_sym_sparql_prefix_token1),
	489:  uint16(139),
	490:  uint16(1),
	491:  uint16(anon_sym_LBRACK),
	492:  uint16(142),
	493:  uint16(1),
	494:  uint16(anon_sym_LPAREN),
	495:  uint16(145),
	496:  uint16(1),
	497:  uint16(anon_sym_LT),
	498:  uint16(148),
	499:  uint16(1),
	500:  uint16(anon_sym_COLON),
	501:  uint16(151),
	502:  uint16(1),
	503:  uint16(anon_sym__COLON),
	504:  uint16(154),
	505:  uint16(1),
	506:  uint16(sym_anon),
	507:  uint16(51),
	508:  uint16(1),
	509:  uint16(sym_blank_node_property_list),
	510:  uint16(53),
	511:  uint16(1),
	512:  uint16(sym_subject),
	513:  uint16(67),
	514:  uint16(1),
	515:  uint16(sym_namespace),
	516:  uint16(86),
	517:  uint16(1),
	518:  uint16(sym_collection),
	519:  uint16(107),
	520:  uint16(1),
	521:  uint16(sym__triples),
	522:  uint16(108),
	523:  uint16(1),
	524:  uint16(sym__label),
	525:  uint16(8),
	526:  uint16(4),
	527:  uint16(sym_graph),
	528:  uint16(sym_triple),
	529:  uint16(sym_directive),
	530:  uint16(aux_sym_document_repeat1),
	531:  uint16(40),
	532:  uint16(4),
	533:  uint16(sym_prefix_id),
	534:  uint16(sym_base),
	535:  uint16(sym_sparql_base),
	536:  uint16(sym_sparql_prefix),
	537:  uint16(79),
	538:  uint16(5),
	539:  uint16(sym_iri_reference),
	540:  uint16(sym__iri),
	541:  uint16(sym_prefixed_name),
	542:  uint16(sym__blank_node),
	543:  uint16(sym_blank_node_label),
	544:  uint16(3),
	545:  uint16(13),
	546:  uint16(1),
	547:  uint16(sym_comment),
	548:  uint16(159),
	549:  uint16(13),
	550:  uint16(anon_sym_GRAPH),
	551:  uint16(anon_sym_DOT),
	552:  uint16(aux_sym_sparql_base_token1),
	553:  uint16(aux_sym_sparql_prefix_token1),
	554:  uint16(anon_sym_a),
	555:  uint16(anon_sym_LBRACK),
	556:  uint16(sym_integer),
	557:  uint16(sym_decimal),
	558:  uint16(anon_sym_DQUOTE),
	559:  uint16(anon_sym_SQUOTE),
	560:  uint16(anon_sym_true),
	561:  uint16(anon_sym_false),
	562:  uint16(sym_pn_prefix),
	563:  uint16(157),
	564:  uint16(17),
	566:  uint16(anon_sym_LBRACE),
	567:  uint16(anon_sym_RBRACE),
	568:  uint16(anon_sym_ATprefix),
	569:  uint16(anon_sym_ATbase),
	570:  uint16(anon_sym_SEMI),
	571:  uint16(anon_sym_COMMA),
	572:  uint16(anon_sym_RBRACK),
	573:  uint16(anon_sym_LPAREN),
	574:  uint16(anon_sym_RPAREN),
	575:  uint16(anon_sym_LT),
	576:  uint16(sym_double),
	577:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	578:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	579:  uint16(anon_sym_COLON),
	580:  uint16(anon_sym__COLON),
	581:  uint16(sym_anon),
	582:  uint16(3),
	583:  uint16(13),
	584:  uint16(1),
	585:  uint16(sym_comment),
	586:  uint16(163),
	587:  uint16(13),
	588:  uint16(anon_sym_GRAPH),
	589:  uint16(anon_sym_DOT),
	590:  uint16(aux_sym_sparql_base_token1),
	591:  uint16(aux_sym_sparql_prefix_token1),
	592:  uint16(anon_sym_a),
	593:  uint16(anon_sym_LBRACK),
	594:  uint16(sym_integer),
	595:  uint16(sym_decimal),
	596:  uint16(anon_sym_DQUOTE),
	597:  uint16(anon_sym_SQUOTE),
	598:  uint16(anon_sym_true),
	599:  uint16(anon_sym_false),
	600:  uint16(sym_pn_prefix),
	601:  uint16(161),
	602:  uint16(17),
	604:  uint16(anon_sym_LBRACE),
	605:  uint16(anon_sym_RBRACE),
	606:  uint16(anon_sym_ATprefix),
	607:  uint16(anon_sym_ATbase),
	608:  uint16(anon_sym_SEMI),
	609:  uint16(anon_sym_COMMA),
	610:  uint16(anon_sym_RBRACK),
	611:  uint16(anon_sym_LPAREN),
	612:  uint16(anon_sym_RPAREN),
	613:  uint16(anon_sym_LT),
	614:  uint16(sym_double),
	615:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	616:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	617:  uint16(anon_sym_COLON),
	618:  uint16(anon_sym__COLON),
	619:  uint16(sym_anon),
	620:  uint16(3),
	621:  uint16(13),
	622:  uint16(1),
	623:  uint16(sym_comment),
	624:  uint16(165),
	625:  uint16(10),
	626:  uint16(anon_sym_DOT),
	627:  uint16(anon_sym_a),
	628:  uint16(anon_sym_LBRACK),
	629:  uint16(sym_integer),
	630:  uint16(sym_decimal),
	631:  uint16(anon_sym_DQUOTE),
	632:  uint16(anon_sym_SQUOTE),
	633:  uint16(anon_sym_true),
	634:  uint16(anon_sym_false),
	635:  uint16(sym_pn_prefix),
	636:  uint16(167),
	637:  uint16(14),
	638:  uint16(anon_sym_LBRACE),
	639:  uint16(anon_sym_RBRACE),
	640:  uint16(anon_sym_SEMI),
	641:  uint16(anon_sym_COMMA),
	642:  uint16(anon_sym_RBRACK),
	643:  uint16(anon_sym_LPAREN),
	644:  uint16(anon_sym_RPAREN),
	645:  uint16(anon_sym_LT),
	646:  uint16(sym_double),
	647:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	648:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	649:  uint16(anon_sym_COLON),
	650:  uint16(anon_sym__COLON),
	651:  uint16(sym_anon),
	652:  uint16(3),
	653:  uint16(13),
	654:  uint16(1),
	655:  uint16(sym_comment),
	656:  uint16(169),
	657:  uint16(9),
	658:  uint16(anon_sym_DOT),
	659:  uint16(anon_sym_LBRACK),
	660:  uint16(sym_integer),
	661:  uint16(sym_decimal),
	662:  uint16(anon_sym_DQUOTE),
	663:  uint16(anon_sym_SQUOTE),
	664:  uint16(anon_sym_true),
	665:  uint16(anon_sym_false),
	666:  uint16(sym_pn_prefix),
	667:  uint16(171),
	668:  uint16(15),
	669:  uint16(anon_sym_RBRACE),
	670:  uint16(anon_sym_SEMI),
	671:  uint16(anon_sym_COMMA),
	672:  uint16(anon_sym_RBRACK),
	673:  uint16(anon_sym_LPAREN),
	674:  uint16(anon_sym_RPAREN),
	675:  uint16(anon_sym_LT),
	676:  uint16(sym_double),
	677:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	678:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	679:  uint16(anon_sym_CARET_CARET),
	680:  uint16(anon_sym_COLON),
	681:  uint16(anon_sym__COLON),
	682:  uint16(sym_lang_tag),
	683:  uint16(sym_anon),
	684:  uint16(3),
	685:  uint16(13),
	686:  uint16(1),
	687:  uint16(sym_comment),
	688:  uint16(173),
	689:  uint16(9),
	690:  uint16(anon_sym_DOT),
	691:  uint16(anon_sym_LBRACK),
	692:  uint16(sym_integer),
	693:  uint16(sym_decimal),
	694:  uint16(anon_sym_DQUOTE),
	695:  uint16(anon_sym_SQUOTE),
	696:  uint16(anon_sym_true),
	697:  uint16(anon_sym_false),
	698:  uint16(sym_pn_prefix),
	699:  uint16(175),
	700:  uint16(15),
	701:  uint16(anon_sym_RBRACE),
	702:  uint16(anon_sym_SEMI),
	703:  uint16(anon_sym_COMMA),
	704:  uint16(anon_sym_RBRACK),
	705:  uint16(anon_sym_LPAREN),
	706:  uint16(anon_sym_RPAREN),
	707:  uint16(anon_sym_LT),
	708:  uint16(sym_double),
	709:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	710:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	711:  uint16(anon_sym_CARET_CARET),
	712:  uint16(anon_sym_COLON),
	713:  uint16(anon_sym__COLON),
	714:  uint16(sym_lang_tag),
	715:  uint16(sym_anon),
	716:  uint16(3),
	717:  uint16(13),
	718:  uint16(1),
	719:  uint16(sym_comment),
	720:  uint16(177),
	721:  uint16(9),
	722:  uint16(anon_sym_DOT),
	723:  uint16(anon_sym_LBRACK),
	724:  uint16(sym_integer),
	725:  uint16(sym_decimal),
	726:  uint16(anon_sym_DQUOTE),
	727:  uint16(anon_sym_SQUOTE),
	728:  uint16(anon_sym_true),
	729:  uint16(anon_sym_false),
	730:  uint16(sym_pn_prefix),
	731:  uint16(179),
	732:  uint16(15),
	733:  uint16(anon_sym_RBRACE),
	734:  uint16(anon_sym_SEMI),
	735:  uint16(anon_sym_COMMA),
	736:  uint16(anon_sym_RBRACK),
	737:  uint16(anon_sym_LPAREN),
	738:  uint16(anon_sym_RPAREN),
	739:  uint16(anon_sym_LT),
	740:  uint16(sym_double),
	741:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	742:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	743:  uint16(anon_sym_CARET_CARET),
	744:  uint16(anon_sym_COLON),
	745:  uint16(anon_sym__COLON),
	746:  uint16(sym_lang_tag),
	747:  uint16(sym_anon),
	748:  uint16(3),
	749:  uint16(13),
	750:  uint16(1),
	751:  uint16(sym_comment),
	752:  uint16(181),
	753:  uint16(9),
	754:  uint16(anon_sym_DOT),
	755:  uint16(anon_sym_LBRACK),
	756:  uint16(sym_integer),
	757:  uint16(sym_decimal),
	758:  uint16(anon_sym_DQUOTE),
	759:  uint16(anon_sym_SQUOTE),
	760:  uint16(anon_sym_true),
	761:  uint16(anon_sym_false),
	762:  uint16(sym_pn_prefix),
	763:  uint16(183),
	764:  uint16(15),
	765:  uint16(anon_sym_RBRACE),
	766:  uint16(anon_sym_SEMI),
	767:  uint16(anon_sym_COMMA),
	768:  uint16(anon_sym_RBRACK),
	769:  uint16(anon_sym_LPAREN),
	770:  uint16(anon_sym_RPAREN),
	771:  uint16(anon_sym_LT),
	772:  uint16(sym_double),
	773:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	774:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	775:  uint16(anon_sym_CARET_CARET),
	776:  uint16(anon_sym_COLON),
	777:  uint16(anon_sym__COLON),
	778:  uint16(sym_lang_tag),
	779:  uint16(sym_anon),
	780:  uint16(3),
	781:  uint16(13),
	782:  uint16(1),
	783:  uint16(sym_comment),
	784:  uint16(185),
	785:  uint16(9),
	786:  uint16(anon_sym_DOT),
	787:  uint16(anon_sym_LBRACK),
	788:  uint16(sym_integer),
	789:  uint16(sym_decimal),
	790:  uint16(anon_sym_DQUOTE),
	791:  uint16(anon_sym_SQUOTE),
	792:  uint16(anon_sym_true),
	793:  uint16(anon_sym_false),
	794:  uint16(sym_pn_prefix),
	795:  uint16(187),
	796:  uint16(15),
	797:  uint16(anon_sym_RBRACE),
	798:  uint16(anon_sym_SEMI),
	799:  uint16(anon_sym_COMMA),
	800:  uint16(anon_sym_RBRACK),
	801:  uint16(anon_sym_LPAREN),
	802:  uint16(anon_sym_RPAREN),
	803:  uint16(anon_sym_LT),
	804:  uint16(sym_double),
	805:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	806:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	807:  uint16(anon_sym_CARET_CARET),
	808:  uint16(anon_sym_COLON),
	809:  uint16(anon_sym__COLON),
	810:  uint16(sym_lang_tag),
	811:  uint16(sym_anon),
	812:  uint16(3),
	813:  uint16(13),
	814:  uint16(1),
	815:  uint16(sym_comment),
	816:  uint16(189),
	817:  uint16(10),
	818:  uint16(anon_sym_DOT),
	819:  uint16(anon_sym_a),
	820:  uint16(anon_sym_LBRACK),
	821:  uint16(sym_integer),
	822:  uint16(sym_decimal),
	823:  uint16(anon_sym_DQUOTE),
	824:  uint16(anon_sym_SQUOTE),
	825:  uint16(anon_sym_true),
	826:  uint16(anon_sym_false),
	827:  uint16(sym_pn_prefix),
	828:  uint16(191),
	829:  uint16(14),
	830:  uint16(anon_sym_LBRACE),
	831:  uint16(anon_sym_RBRACE),
	832:  uint16(anon_sym_SEMI),
	833:  uint16(anon_sym_COMMA),
	834:  uint16(anon_sym_RBRACK),
	835:  uint16(anon_sym_LPAREN),
	836:  uint16(anon_sym_RPAREN),
	837:  uint16(anon_sym_LT),
	838:  uint16(sym_double),
	839:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	840:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	841:  uint16(anon_sym_COLON),
	842:  uint16(anon_sym__COLON),
	843:  uint16(sym_anon),
	844:  uint16(3),
	845:  uint16(13),
	846:  uint16(1),
	847:  uint16(sym_comment),
	848:  uint16(193),
	849:  uint16(9),
	850:  uint16(anon_sym_DOT),
	851:  uint16(anon_sym_LBRACK),
	852:  uint16(sym_integer),
	853:  uint16(sym_decimal),
	854:  uint16(anon_sym_DQUOTE),
	855:  uint16(anon_sym_SQUOTE),
	856:  uint16(anon_sym_true),
	857:  uint16(anon_sym_false),
	858:  uint16(sym_pn_prefix),
	859:  uint16(195),
	860:  uint16(15),
	861:  uint16(anon_sym_RBRACE),
	862:  uint16(anon_sym_SEMI),
	863:  uint16(anon_sym_COMMA),
	864:  uint16(anon_sym_RBRACK),
	865:  uint16(anon_sym_LPAREN),
	866:  uint16(anon_sym_RPAREN),
	867:  uint16(anon_sym_LT),
	868:  uint16(sym_double),
	869:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	870:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	871:  uint16(anon_sym_CARET_CARET),
	872:  uint16(anon_sym_COLON),
	873:  uint16(anon_sym__COLON),
	874:  uint16(sym_lang_tag),
	875:  uint16(sym_anon),
	876:  uint16(3),
	877:  uint16(13),
	878:  uint16(1),
	879:  uint16(sym_comment),
	880:  uint16(197),
	881:  uint16(9),
	882:  uint16(anon_sym_DOT),
	883:  uint16(anon_sym_LBRACK),
	884:  uint16(sym_integer),
	885:  uint16(sym_decimal),
	886:  uint16(anon_sym_DQUOTE),
	887:  uint16(anon_sym_SQUOTE),
	888:  uint16(anon_sym_true),
	889:  uint16(anon_sym_false),
	890:  uint16(sym_pn_prefix),
	891:  uint16(199),
	892:  uint16(15),
	893:  uint16(anon_sym_RBRACE),
	894:  uint16(anon_sym_SEMI),
	895:  uint16(anon_sym_COMMA),
	896:  uint16(anon_sym_RBRACK),
	897:  uint16(anon_sym_LPAREN),
	898:  uint16(anon_sym_RPAREN),
	899:  uint16(anon_sym_LT),
	900:  uint16(sym_double),
	901:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	902:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	903:  uint16(anon_sym_CARET_CARET),
	904:  uint16(anon_sym_COLON),
	905:  uint16(anon_sym__COLON),
	906:  uint16(sym_lang_tag),
	907:  uint16(sym_anon),
	908:  uint16(3),
	909:  uint16(13),
	910:  uint16(1),
	911:  uint16(sym_comment),
	912:  uint16(201),
	913:  uint16(9),
	914:  uint16(anon_sym_DOT),
	915:  uint16(anon_sym_LBRACK),
	916:  uint16(sym_integer),
	917:  uint16(sym_decimal),
	918:  uint16(anon_sym_DQUOTE),
	919:  uint16(anon_sym_SQUOTE),
	920:  uint16(anon_sym_true),
	921:  uint16(anon_sym_false),
	922:  uint16(sym_pn_prefix),
	923:  uint16(203),
	924:  uint16(15),
	925:  uint16(anon_sym_RBRACE),
	926:  uint16(anon_sym_SEMI),
	927:  uint16(anon_sym_COMMA),
	928:  uint16(anon_sym_RBRACK),
	929:  uint16(anon_sym_LPAREN),
	930:  uint16(anon_sym_RPAREN),
	931:  uint16(anon_sym_LT),
	932:  uint16(sym_double),
	933:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	934:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	935:  uint16(anon_sym_CARET_CARET),
	936:  uint16(anon_sym_COLON),
	937:  uint16(anon_sym__COLON),
	938:  uint16(sym_lang_tag),
	939:  uint16(sym_anon),
	940:  uint16(3),
	941:  uint16(13),
	942:  uint16(1),
	943:  uint16(sym_comment),
	944:  uint16(205),
	945:  uint16(9),
	946:  uint16(anon_sym_DOT),
	947:  uint16(anon_sym_LBRACK),
	948:  uint16(sym_integer),
	949:  uint16(sym_decimal),
	950:  uint16(anon_sym_DQUOTE),
	951:  uint16(anon_sym_SQUOTE),
	952:  uint16(anon_sym_true),
	953:  uint16(anon_sym_false),
	954:  uint16(sym_pn_prefix),
	955:  uint16(207),
	956:  uint16(15),
	957:  uint16(anon_sym_RBRACE),
	958:  uint16(anon_sym_SEMI),
	959:  uint16(anon_sym_COMMA),
	960:  uint16(anon_sym_RBRACK),
	961:  uint16(anon_sym_LPAREN),
	962:  uint16(anon_sym_RPAREN),
	963:  uint16(anon_sym_LT),
	964:  uint16(sym_double),
	965:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	966:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	967:  uint16(anon_sym_CARET_CARET),
	968:  uint16(anon_sym_COLON),
	969:  uint16(anon_sym__COLON),
	970:  uint16(sym_lang_tag),
	971:  uint16(sym_anon),
	972:  uint16(3),
	973:  uint16(13),
	974:  uint16(1),
	975:  uint16(sym_comment),
	976:  uint16(209),
	977:  uint16(10),
	978:  uint16(anon_sym_DOT),
	979:  uint16(anon_sym_a),
	980:  uint16(anon_sym_LBRACK),
	981:  uint16(sym_integer),
	982:  uint16(sym_decimal),
	983:  uint16(anon_sym_DQUOTE),
	984:  uint16(anon_sym_SQUOTE),
	985:  uint16(anon_sym_true),
	986:  uint16(anon_sym_false),
	987:  uint16(sym_pn_prefix),
	988:  uint16(211),
	989:  uint16(13),
	990:  uint16(anon_sym_RBRACE),
	991:  uint16(anon_sym_SEMI),
	992:  uint16(anon_sym_COMMA),
	993:  uint16(anon_sym_RBRACK),
	994:  uint16(anon_sym_LPAREN),
	995:  uint16(anon_sym_RPAREN),
	996:  uint16(anon_sym_LT),
	997:  uint16(sym_double),
	998:  uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	999:  uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1000: uint16(anon_sym_COLON),
	1001: uint16(anon_sym__COLON),
	1002: uint16(sym_anon),
	1003: uint16(3),
	1004: uint16(13),
	1005: uint16(1),
	1006: uint16(sym_comment),
	1007: uint16(213),
	1008: uint16(10),
	1009: uint16(anon_sym_DOT),
	1010: uint16(anon_sym_a),
	1011: uint16(anon_sym_LBRACK),
	1012: uint16(sym_integer),
	1013: uint16(sym_decimal),
	1014: uint16(anon_sym_DQUOTE),
	1015: uint16(anon_sym_SQUOTE),
	1016: uint16(anon_sym_true),
	1017: uint16(anon_sym_false),
	1018: uint16(sym_pn_prefix),
	1019: uint16(215),
	1020: uint16(13),
	1021: uint16(anon_sym_RBRACE),
	1022: uint16(anon_sym_SEMI),
	1023: uint16(anon_sym_COMMA),
	1024: uint16(anon_sym_RBRACK),
	1025: uint16(anon_sym_LPAREN),
	1026: uint16(anon_sym_RPAREN),
	1027: uint16(anon_sym_LT),
	1028: uint16(sym_double),
	1029: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1030: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1031: uint16(anon_sym_COLON),
	1032: uint16(anon_sym__COLON),
	1033: uint16(sym_anon),
	1034: uint16(3),
	1035: uint16(13),
	1036: uint16(1),
	1037: uint16(sym_comment),
	1038: uint16(217),
	1039: uint16(10),
	1040: uint16(anon_sym_DOT),
	1041: uint16(anon_sym_a),
	1042: uint16(anon_sym_LBRACK),
	1043: uint16(sym_integer),
	1044: uint16(sym_decimal),
	1045: uint16(anon_sym_DQUOTE),
	1046: uint16(anon_sym_SQUOTE),
	1047: uint16(anon_sym_true),
	1048: uint16(anon_sym_false),
	1049: uint16(sym_pn_prefix),
	1050: uint16(219),
	1051: uint16(13),
	1052: uint16(anon_sym_RBRACE),
	1053: uint16(anon_sym_SEMI),
	1054: uint16(anon_sym_COMMA),
	1055: uint16(anon_sym_RBRACK),
	1056: uint16(anon_sym_LPAREN),
	1057: uint16(anon_sym_RPAREN),
	1058: uint16(anon_sym_LT),
	1059: uint16(sym_double),
	1060: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1061: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1062: uint16(anon_sym_COLON),
	1063: uint16(anon_sym__COLON),
	1064: uint16(sym_anon),
	1065: uint16(3),
	1066: uint16(13),
	1067: uint16(1),
	1068: uint16(sym_comment),
	1069: uint16(221),
	1070: uint16(9),
	1071: uint16(anon_sym_DOT),
	1072: uint16(anon_sym_LBRACK),
	1073: uint16(sym_integer),
	1074: uint16(sym_decimal),
	1075: uint16(anon_sym_DQUOTE),
	1076: uint16(anon_sym_SQUOTE),
	1077: uint16(anon_sym_true),
	1078: uint16(anon_sym_false),
	1079: uint16(sym_pn_prefix),
	1080: uint16(223),
	1081: uint16(13),
	1082: uint16(anon_sym_RBRACE),
	1083: uint16(anon_sym_SEMI),
	1084: uint16(anon_sym_COMMA),
	1085: uint16(anon_sym_RBRACK),
	1086: uint16(anon_sym_LPAREN),
	1087: uint16(anon_sym_RPAREN),
	1088: uint16(anon_sym_LT),
	1089: uint16(sym_double),
	1090: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1091: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1092: uint16(anon_sym_COLON),
	1093: uint16(anon_sym__COLON),
	1094: uint16(sym_anon),
	1095: uint16(3),
	1096: uint16(13),
	1097: uint16(1),
	1098: uint16(sym_comment),
	1099: uint16(225),
	1100: uint16(9),
	1101: uint16(anon_sym_DOT),
	1102: uint16(anon_sym_LBRACK),
	1103: uint16(sym_integer),
	1104: uint16(sym_decimal),
	1105: uint16(anon_sym_DQUOTE),
	1106: uint16(anon_sym_SQUOTE),
	1107: uint16(anon_sym_true),
	1108: uint16(anon_sym_false),
	1109: uint16(sym_pn_prefix),
	1110: uint16(227),
	1111: uint16(13),
	1112: uint16(anon_sym_RBRACE),
	1113: uint16(anon_sym_SEMI),
	1114: uint16(anon_sym_COMMA),
	1115: uint16(anon_sym_RBRACK),
	1116: uint16(anon_sym_LPAREN),
	1117: uint16(anon_sym_RPAREN),
	1118: uint16(anon_sym_LT),
	1119: uint16(sym_double),
	1120: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1121: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1122: uint16(anon_sym_COLON),
	1123: uint16(anon_sym__COLON),
	1124: uint16(sym_anon),
	1125: uint16(3),
	1126: uint16(13),
	1127: uint16(1),
	1128: uint16(sym_comment),
	1129: uint16(229),
	1130: uint16(9),
	1131: uint16(anon_sym_DOT),
	1132: uint16(anon_sym_LBRACK),
	1133: uint16(sym_integer),
	1134: uint16(sym_decimal),
	1135: uint16(anon_sym_DQUOTE),
	1136: uint16(anon_sym_SQUOTE),
	1137: uint16(anon_sym_true),
	1138: uint16(anon_sym_false),
	1139: uint16(sym_pn_prefix),
	1140: uint16(231),
	1141: uint16(13),
	1142: uint16(anon_sym_RBRACE),
	1143: uint16(anon_sym_SEMI),
	1144: uint16(anon_sym_COMMA),
	1145: uint16(anon_sym_RBRACK),
	1146: uint16(anon_sym_LPAREN),
	1147: uint16(anon_sym_RPAREN),
	1148: uint16(anon_sym_LT),
	1149: uint16(sym_double),
	1150: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1151: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1152: uint16(anon_sym_COLON),
	1153: uint16(anon_sym__COLON),
	1154: uint16(sym_anon),
	1155: uint16(15),
	1156: uint16(13),
	1157: uint16(1),
	1158: uint16(sym_comment),
	1159: uint16(23),
	1160: uint16(1),
	1161: uint16(anon_sym_LBRACK),
	1162: uint16(25),
	1163: uint16(1),
	1164: uint16(anon_sym_LPAREN),
	1165: uint16(27),
	1166: uint16(1),
	1167: uint16(anon_sym_LT),
	1168: uint16(29),
	1169: uint16(1),
	1170: uint16(anon_sym_COLON),
	1171: uint16(31),
	1172: uint16(1),
	1173: uint16(anon_sym__COLON),
	1174: uint16(233),
	1175: uint16(1),
	1176: uint16(sym_pn_prefix),
	1177: uint16(235),
	1178: uint16(1),
	1179: uint16(anon_sym_RBRACE),
	1180: uint16(237),
	1181: uint16(1),
	1182: uint16(sym_anon),
	1183: uint16(51),
	1184: uint16(1),
	1185: uint16(sym_blank_node_property_list),
	1186: uint16(53),
	1187: uint16(1),
	1188: uint16(sym_subject),
	1189: uint16(67),
	1190: uint16(1),
	1191: uint16(sym_namespace),
	1192: uint16(93),
	1193: uint16(1),
	1194: uint16(sym__triples),
	1195: uint16(29),
	1196: uint16(2),
	1197: uint16(sym_triple),
	1198: uint16(aux_sym_graph_repeat1),
	1199: uint16(86),
	1200: uint16(6),
	1201: uint16(sym_collection),
	1202: uint16(sym_iri_reference),
	1203: uint16(sym__iri),
	1204: uint16(sym_prefixed_name),
	1205: uint16(sym__blank_node),
	1206: uint16(sym_blank_node_label),
	1207: uint16(15),
	1208: uint16(13),
	1209: uint16(1),
	1210: uint16(sym_comment),
	1211: uint16(239),
	1212: uint16(1),
	1213: uint16(sym_pn_prefix),
	1214: uint16(242),
	1215: uint16(1),
	1216: uint16(anon_sym_RBRACE),
	1217: uint16(244),
	1218: uint16(1),
	1219: uint16(anon_sym_LBRACK),
	1220: uint16(247),
	1221: uint16(1),
	1222: uint16(anon_sym_LPAREN),
	1223: uint16(250),
	1224: uint16(1),
	1225: uint16(anon_sym_LT),
	1226: uint16(253),
	1227: uint16(1),
	1228: uint16(anon_sym_COLON),
	1229: uint16(256),
	1230: uint16(1),
	1231: uint16(anon_sym__COLON),
	1232: uint16(259),
	1233: uint16(1),
	1234: uint16(sym_anon),
	1235: uint16(51),
	1236: uint16(1),
	1237: uint16(sym_blank_node_property_list),
	1238: uint16(53),
	1239: uint16(1),
	1240: uint16(sym_subject),
	1241: uint16(67),
	1242: uint16(1),
	1243: uint16(sym_namespace),
	1244: uint16(107),
	1245: uint16(1),
	1246: uint16(sym__triples),
	1247: uint16(29),
	1248: uint16(2),
	1249: uint16(sym_triple),
	1250: uint16(aux_sym_graph_repeat1),
	1251: uint16(86),
	1252: uint16(6),
	1253: uint16(sym_collection),
	1254: uint16(sym_iri_reference),
	1255: uint16(sym__iri),
	1256: uint16(sym_prefixed_name),
	1257: uint16(sym__blank_node),
	1258: uint16(sym_blank_node_label),
	1259: uint16(15),
	1260: uint16(13),
	1261: uint16(1),
	1262: uint16(sym_comment),
	1263: uint16(23),
	1264: uint16(1),
	1265: uint16(anon_sym_LBRACK),
	1266: uint16(25),
	1267: uint16(1),
	1268: uint16(anon_sym_LPAREN),
	1269: uint16(27),
	1270: uint16(1),
	1271: uint16(anon_sym_LT),
	1272: uint16(29),
	1273: uint16(1),
	1274: uint16(anon_sym_COLON),
	1275: uint16(31),
	1276: uint16(1),
	1277: uint16(anon_sym__COLON),
	1278: uint16(233),
	1279: uint16(1),
	1280: uint16(sym_pn_prefix),
	1281: uint16(237),
	1282: uint16(1),
	1283: uint16(sym_anon),
	1284: uint16(262),
	1285: uint16(1),
	1286: uint16(anon_sym_RBRACE),
	1287: uint16(51),
	1288: uint16(1),
	1289: uint16(sym_blank_node_property_list),
	1290: uint16(53),
	1291: uint16(1),
	1292: uint16(sym_subject),
	1293: uint16(67),
	1294: uint16(1),
	1295: uint16(sym_namespace),
	1296: uint16(99),
	1297: uint16(1),
	1298: uint16(sym__triples),
	1299: uint16(28),
	1300: uint16(2),
	1301: uint16(sym_triple),
	1302: uint16(aux_sym_graph_repeat1),
	1303: uint16(86),
	1304: uint16(6),
	1305: uint16(sym_collection),
	1306: uint16(sym_iri_reference),
	1307: uint16(sym__iri),
	1308: uint16(sym_prefixed_name),
	1309: uint16(sym__blank_node),
	1310: uint16(sym_blank_node_label),
	1311: uint16(15),
	1312: uint16(13),
	1313: uint16(1),
	1314: uint16(sym_comment),
	1315: uint16(23),
	1316: uint16(1),
	1317: uint16(anon_sym_LBRACK),
	1318: uint16(25),
	1319: uint16(1),
	1320: uint16(anon_sym_LPAREN),
	1321: uint16(27),
	1322: uint16(1),
	1323: uint16(anon_sym_LT),
	1324: uint16(29),
	1325: uint16(1),
	1326: uint16(anon_sym_COLON),
	1327: uint16(31),
	1328: uint16(1),
	1329: uint16(anon_sym__COLON),
	1330: uint16(233),
	1331: uint16(1),
	1332: uint16(sym_pn_prefix),
	1333: uint16(237),
	1334: uint16(1),
	1335: uint16(sym_anon),
	1336: uint16(264),
	1337: uint16(1),
	1338: uint16(anon_sym_RBRACE),
	1339: uint16(51),
	1340: uint16(1),
	1341: uint16(sym_blank_node_property_list),
	1342: uint16(53),
	1343: uint16(1),
	1344: uint16(sym_subject),
	1345: uint16(67),
	1346: uint16(1),
	1347: uint16(sym_namespace),
	1348: uint16(97),
	1349: uint16(1),
	1350: uint16(sym__triples),
	1351: uint16(29),
	1352: uint16(2),
	1353: uint16(sym_triple),
	1354: uint16(aux_sym_graph_repeat1),
	1355: uint16(86),
	1356: uint16(6),
	1357: uint16(sym_collection),
	1358: uint16(sym_iri_reference),
	1359: uint16(sym__iri),
	1360: uint16(sym_prefixed_name),
	1361: uint16(sym__blank_node),
	1362: uint16(sym_blank_node_label),
	1363: uint16(15),
	1364: uint16(13),
	1365: uint16(1),
	1366: uint16(sym_comment),
	1367: uint16(23),
	1368: uint16(1),
	1369: uint16(anon_sym_LBRACK),
	1370: uint16(25),
	1371: uint16(1),
	1372: uint16(anon_sym_LPAREN),
	1373: uint16(27),
	1374: uint16(1),
	1375: uint16(anon_sym_LT),
	1376: uint16(29),
	1377: uint16(1),
	1378: uint16(anon_sym_COLON),
	1379: uint16(31),
	1380: uint16(1),
	1381: uint16(anon_sym__COLON),
	1382: uint16(233),
	1383: uint16(1),
	1384: uint16(sym_pn_prefix),
	1385: uint16(237),
	1386: uint16(1),
	1387: uint16(sym_anon),
	1388: uint16(266),
	1389: uint16(1),
	1390: uint16(anon_sym_RBRACE),
	1391: uint16(51),
	1392: uint16(1),
	1393: uint16(sym_blank_node_property_list),
	1394: uint16(53),
	1395: uint16(1),
	1396: uint16(sym_subject),
	1397: uint16(67),
	1398: uint16(1),
	1399: uint16(sym_namespace),
	1400: uint16(102),
	1401: uint16(1),
	1402: uint16(sym__triples),
	1403: uint16(31),
	1404: uint16(2),
	1405: uint16(sym_triple),
	1406: uint16(aux_sym_graph_repeat1),
	1407: uint16(86),
	1408: uint16(6),
	1409: uint16(sym_collection),
	1410: uint16(sym_iri_reference),
	1411: uint16(sym__iri),
	1412: uint16(sym_prefixed_name),
	1413: uint16(sym__blank_node),
	1414: uint16(sym_blank_node_label),
	1415: uint16(5),
	1416: uint16(13),
	1417: uint16(1),
	1418: uint16(sym_comment),
	1419: uint16(272),
	1420: uint16(1),
	1421: uint16(anon_sym_CARET_CARET),
	1422: uint16(274),
	1423: uint16(1),
	1424: uint16(sym_lang_tag),
	1425: uint16(268),
	1426: uint16(8),
	1427: uint16(anon_sym_LBRACK),
	1428: uint16(sym_integer),
	1429: uint16(sym_decimal),
	1430: uint16(anon_sym_DQUOTE),
	1431: uint16(anon_sym_SQUOTE),
	1432: uint16(anon_sym_true),
	1433: uint16(anon_sym_false),
	1434: uint16(sym_pn_prefix),
	1435: uint16(270),
	1436: uint16(9),
	1437: uint16(anon_sym_LPAREN),
	1438: uint16(anon_sym_RPAREN),
	1439: uint16(anon_sym_LT),
	1440: uint16(sym_double),
	1441: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1442: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1443: uint16(anon_sym_COLON),
	1444: uint16(anon_sym__COLON),
	1445: uint16(sym_anon),
	1446: uint16(3),
	1447: uint16(13),
	1448: uint16(1),
	1449: uint16(sym_comment),
	1450: uint16(278),
	1451: uint16(6),
	1452: uint16(anon_sym_LPAREN),
	1453: uint16(anon_sym_RPAREN),
	1454: uint16(anon_sym_LT),
	1455: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1456: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1457: uint16(sym_anon),
	1458: uint16(276),
	1459: uint16(12),
	1460: uint16(anon_sym_LBRACK),
	1461: uint16(sym_integer),
	1462: uint16(sym_decimal),
	1463: uint16(sym_double),
	1464: uint16(anon_sym_DQUOTE),
	1465: uint16(anon_sym_SQUOTE),
	1466: uint16(anon_sym_true),
	1467: uint16(anon_sym_false),
	1468: uint16(anon_sym_COLON),
	1469: uint16(anon_sym__COLON),
	1470: uint16(sym_pn_prefix),
	1471: uint16(sym_pn_local),
	1472: uint16(4),
	1473: uint16(13),
	1474: uint16(1),
	1475: uint16(sym_comment),
	1476: uint16(284),
	1477: uint16(1),
	1478: uint16(sym_pn_local),
	1479: uint16(282),
	1480: uint16(6),
	1481: uint16(anon_sym_LPAREN),
	1482: uint16(anon_sym_RPAREN),
	1483: uint16(anon_sym_LT),
	1484: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1485: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1486: uint16(sym_anon),
	1487: uint16(280),
	1488: uint16(11),
	1489: uint16(anon_sym_LBRACK),
	1490: uint16(sym_integer),
	1491: uint16(sym_decimal),
	1492: uint16(sym_double),
	1493: uint16(anon_sym_DQUOTE),
	1494: uint16(anon_sym_SQUOTE),
	1495: uint16(anon_sym_true),
	1496: uint16(anon_sym_false),
	1497: uint16(anon_sym_COLON),
	1498: uint16(anon_sym__COLON),
	1499: uint16(sym_pn_prefix),
	1500: uint16(3),
	1501: uint16(13),
	1502: uint16(1),
	1503: uint16(sym_comment),
	1504: uint16(288),
	1505: uint16(6),
	1506: uint16(anon_sym_LPAREN),
	1507: uint16(anon_sym_RPAREN),
	1508: uint16(anon_sym_LT),
	1509: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1510: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1511: uint16(sym_anon),
	1512: uint16(286),
	1513: uint16(12),
	1514: uint16(anon_sym_LBRACK),
	1515: uint16(sym_integer),
	1516: uint16(sym_decimal),
	1517: uint16(sym_double),
	1518: uint16(anon_sym_DQUOTE),
	1519: uint16(anon_sym_SQUOTE),
	1520: uint16(anon_sym_true),
	1521: uint16(anon_sym_false),
	1522: uint16(anon_sym_COLON),
	1523: uint16(anon_sym__COLON),
	1524: uint16(sym_pn_prefix),
	1525: uint16(sym_pn_local),
	1526: uint16(3),
	1527: uint16(13),
	1528: uint16(1),
	1529: uint16(sym_comment),
	1530: uint16(290),
	1531: uint16(8),
	1532: uint16(anon_sym_LBRACK),
	1533: uint16(sym_integer),
	1534: uint16(sym_decimal),
	1535: uint16(anon_sym_DQUOTE),
	1536: uint16(anon_sym_SQUOTE),
	1537: uint16(anon_sym_true),
	1538: uint16(anon_sym_false),
	1539: uint16(sym_pn_prefix),
	1540: uint16(292),
	1541: uint16(8),
	1542: uint16(anon_sym_LPAREN),
	1543: uint16(anon_sym_LT),
	1544: uint16(sym_double),
	1545: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	1546: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1547: uint16(anon_sym_COLON),
	1548: uint16(anon_sym__COLON),
	1549: uint16(sym_anon),
	1550: uint16(3),
	1551: uint16(13),
	1552: uint16(1),
	1553: uint16(sym_comment),
	1554: uint16(296),
	1555: uint16(5),
	1556: uint16(anon_sym_GRAPH),
	1557: uint16(aux_sym_sparql_base_token1),
	1558: uint16(aux_sym_sparql_prefix_token1),
	1559: uint16(anon_sym_LBRACK),
	1560: uint16(sym_pn_prefix),
	1561: uint16(294),
	1562: uint16(10),
	1564: uint16(anon_sym_LBRACE),
	1565: uint16(anon_sym_RBRACE),
	1566: uint16(anon_sym_ATprefix),
	1567: uint16(anon_sym_ATbase),
	1568: uint16(anon_sym_LPAREN),
	1569: uint16(anon_sym_LT),
	1570: uint16(anon_sym_COLON),
	1571: uint16(anon_sym__COLON),
	1572: uint16(sym_anon),
	1573: uint16(10),
	1574: uint16(13),
	1575: uint16(1),
	1576: uint16(sym_comment),
	1577: uint16(27),
	1578: uint16(1),
	1579: uint16(anon_sym_LT),
	1580: uint16(35),
	1581: uint16(1),
	1582: uint16(sym_pn_prefix),
	1583: uint16(53),
	1584: uint16(1),
	1585: uint16(anon_sym_COLON),
	1586: uint16(300),
	1587: uint16(1),
	1588: uint16(anon_sym_a),
	1589: uint16(5),
	1590: uint16(1),
	1591: uint16(sym_predicate),
	1592: uint16(35),
	1593: uint16(1),
	1594: uint16(sym_namespace),
	1595: uint16(88),
	1596: uint16(1),
	1597: uint16(sym_property),
	1598: uint16(37),
	1599: uint16(3),
	1600: uint16(sym_iri_reference),
	1601: uint16(sym__iri),
	1602: uint16(sym_prefixed_name),
	1603: uint16(298),
	1604: uint16(4),
	1605: uint16(anon_sym_RBRACE),
	1606: uint16(anon_sym_DOT),
	1607: uint16(anon_sym_SEMI),
	1608: uint16(anon_sym_RBRACK),
	1609: uint16(3),
	1610: uint16(13),
	1611: uint16(1),
	1612: uint16(sym_comment),
	1613: uint16(304),
	1614: uint16(5),
	1615: uint16(anon_sym_GRAPH),
	1616: uint16(aux_sym_sparql_base_token1),
	1617: uint16(aux_sym_sparql_prefix_token1),
	1618: uint16(anon_sym_LBRACK),
	1619: uint16(sym_pn_prefix),
	1620: uint16(302),
	1621: uint16(9),
	1623: uint16(anon_sym_LBRACE),
	1624: uint16(anon_sym_ATprefix),
	1625: uint16(anon_sym_ATbase),
	1626: uint16(anon_sym_LPAREN),
	1627: uint16(anon_sym_LT),
	1628: uint16(anon_sym_COLON),
	1629: uint16(anon_sym__COLON),
	1630: uint16(sym_anon),
	1631: uint16(3),
	1632: uint16(13),
	1633: uint16(1),
	1634: uint16(sym_comment),
	1635: uint16(308),
	1636: uint16(5),
	1637: uint16(anon_sym_GRAPH),
	1638: uint16(aux_sym_sparql_base_token1),
	1639: uint16(aux_sym_sparql_prefix_token1),
	1640: uint16(anon_sym_LBRACK),
	1641: uint16(sym_pn_prefix),
	1642: uint16(306),
	1643: uint16(9),
	1645: uint16(anon_sym_LBRACE),
	1646: uint16(anon_sym_ATprefix),
	1647: uint16(anon_sym_ATbase),
	1648: uint16(anon_sym_LPAREN),
	1649: uint16(anon_sym_LT),
	1650: uint16(anon_sym_COLON),
	1651: uint16(anon_sym__COLON),
	1652: uint16(sym_anon),
	1653: uint16(3),
	1654: uint16(13),
	1655: uint16(1),
	1656: uint16(sym_comment),
	1657: uint16(312),
	1658: uint16(5),
	1659: uint16(anon_sym_GRAPH),
	1660: uint16(aux_sym_sparql_base_token1),
	1661: uint16(aux_sym_sparql_prefix_token1),
	1662: uint16(anon_sym_LBRACK),
	1663: uint16(sym_pn_prefix),
	1664: uint16(310),
	1665: uint16(9),
	1667: uint16(anon_sym_LBRACE),
	1668: uint16(anon_sym_ATprefix),
	1669: uint16(anon_sym_ATbase),
	1670: uint16(anon_sym_LPAREN),
	1671: uint16(anon_sym_LT),
	1672: uint16(anon_sym_COLON),
	1673: uint16(anon_sym__COLON),
	1674: uint16(sym_anon),
	1675: uint16(3),
	1676: uint16(13),
	1677: uint16(1),
	1678: uint16(sym_comment),
	1679: uint16(316),
	1680: uint16(5),
	1681: uint16(anon_sym_GRAPH),
	1682: uint16(aux_sym_sparql_base_token1),
	1683: uint16(aux_sym_sparql_prefix_token1),
	1684: uint16(anon_sym_LBRACK),
	1685: uint16(sym_pn_prefix),
	1686: uint16(314),
	1687: uint16(9),
	1689: uint16(anon_sym_LBRACE),
	1690: uint16(anon_sym_ATprefix),
	1691: uint16(anon_sym_ATbase),
	1692: uint16(anon_sym_LPAREN),
	1693: uint16(anon_sym_LT),
	1694: uint16(anon_sym_COLON),
	1695: uint16(anon_sym__COLON),
	1696: uint16(sym_anon),
	1697: uint16(3),
	1698: uint16(13),
	1699: uint16(1),
	1700: uint16(sym_comment),
	1701: uint16(320),
	1702: uint16(5),
	1703: uint16(anon_sym_GRAPH),
	1704: uint16(aux_sym_sparql_base_token1),
	1705: uint16(aux_sym_sparql_prefix_token1),
	1706: uint16(anon_sym_LBRACK),
	1707: uint16(sym_pn_prefix),
	1708: uint16(318),
	1709: uint16(9),
	1711: uint16(anon_sym_LBRACE),
	1712: uint16(anon_sym_ATprefix),
	1713: uint16(anon_sym_ATbase),
	1714: uint16(anon_sym_LPAREN),
	1715: uint16(anon_sym_LT),
	1716: uint16(anon_sym_COLON),
	1717: uint16(anon_sym__COLON),
	1718: uint16(sym_anon),
	1719: uint16(3),
	1720: uint16(13),
	1721: uint16(1),
	1722: uint16(sym_comment),
	1723: uint16(324),
	1724: uint16(5),
	1725: uint16(anon_sym_GRAPH),
	1726: uint16(aux_sym_sparql_base_token1),
	1727: uint16(aux_sym_sparql_prefix_token1),
	1728: uint16(anon_sym_LBRACK),
	1729: uint16(sym_pn_prefix),
	1730: uint16(322),
	1731: uint16(9),
	1733: uint16(anon_sym_LBRACE),
	1734: uint16(anon_sym_ATprefix),
	1735: uint16(anon_sym_ATbase),
	1736: uint16(anon_sym_LPAREN),
	1737: uint16(anon_sym_LT),
	1738: uint16(anon_sym_COLON),
	1739: uint16(anon_sym__COLON),
	1740: uint16(sym_anon),
	1741: uint16(3),
	1742: uint16(13),
	1743: uint16(1),
	1744: uint16(sym_comment),
	1745: uint16(328),
	1746: uint16(5),
	1747: uint16(anon_sym_GRAPH),
	1748: uint16(aux_sym_sparql_base_token1),
	1749: uint16(aux_sym_sparql_prefix_token1),
	1750: uint16(anon_sym_LBRACK),
	1751: uint16(sym_pn_prefix),
	1752: uint16(326),
	1753: uint16(9),
	1755: uint16(anon_sym_LBRACE),
	1756: uint16(anon_sym_ATprefix),
	1757: uint16(anon_sym_ATbase),
	1758: uint16(anon_sym_LPAREN),
	1759: uint16(anon_sym_LT),
	1760: uint16(anon_sym_COLON),
	1761: uint16(anon_sym__COLON),
	1762: uint16(sym_anon),
	1763: uint16(3),
	1764: uint16(13),
	1765: uint16(1),
	1766: uint16(sym_comment),
	1767: uint16(332),
	1768: uint16(5),
	1769: uint16(anon_sym_GRAPH),
	1770: uint16(aux_sym_sparql_base_token1),
	1771: uint16(aux_sym_sparql_prefix_token1),
	1772: uint16(anon_sym_LBRACK),
	1773: uint16(sym_pn_prefix),
	1774: uint16(330),
	1775: uint16(9),
	1777: uint16(anon_sym_LBRACE),
	1778: uint16(anon_sym_ATprefix),
	1779: uint16(anon_sym_ATbase),
	1780: uint16(anon_sym_LPAREN),
	1781: uint16(anon_sym_LT),
	1782: uint16(anon_sym_COLON),
	1783: uint16(anon_sym__COLON),
	1784: uint16(sym_anon),
	1785: uint16(3),
	1786: uint16(13),
	1787: uint16(1),
	1788: uint16(sym_comment),
	1789: uint16(336),
	1790: uint16(5),
	1791: uint16(anon_sym_GRAPH),
	1792: uint16(aux_sym_sparql_base_token1),
	1793: uint16(aux_sym_sparql_prefix_token1),
	1794: uint16(anon_sym_LBRACK),
	1795: uint16(sym_pn_prefix),
	1796: uint16(334),
	1797: uint16(9),
	1799: uint16(anon_sym_LBRACE),
	1800: uint16(anon_sym_ATprefix),
	1801: uint16(anon_sym_ATbase),
	1802: uint16(anon_sym_LPAREN),
	1803: uint16(anon_sym_LT),
	1804: uint16(anon_sym_COLON),
	1805: uint16(anon_sym__COLON),
	1806: uint16(sym_anon),
	1807: uint16(3),
	1808: uint16(13),
	1809: uint16(1),
	1810: uint16(sym_comment),
	1811: uint16(340),
	1812: uint16(5),
	1813: uint16(anon_sym_GRAPH),
	1814: uint16(aux_sym_sparql_base_token1),
	1815: uint16(aux_sym_sparql_prefix_token1),
	1816: uint16(anon_sym_LBRACK),
	1817: uint16(sym_pn_prefix),
	1818: uint16(338),
	1819: uint16(9),
	1821: uint16(anon_sym_LBRACE),
	1822: uint16(anon_sym_ATprefix),
	1823: uint16(anon_sym_ATbase),
	1824: uint16(anon_sym_LPAREN),
	1825: uint16(anon_sym_LT),
	1826: uint16(anon_sym_COLON),
	1827: uint16(anon_sym__COLON),
	1828: uint16(sym_anon),
	1829: uint16(3),
	1830: uint16(13),
	1831: uint16(1),
	1832: uint16(sym_comment),
	1833: uint16(344),
	1834: uint16(5),
	1835: uint16(anon_sym_GRAPH),
	1836: uint16(aux_sym_sparql_base_token1),
	1837: uint16(aux_sym_sparql_prefix_token1),
	1838: uint16(anon_sym_LBRACK),
	1839: uint16(sym_pn_prefix),
	1840: uint16(342),
	1841: uint16(9),
	1843: uint16(anon_sym_LBRACE),
	1844: uint16(anon_sym_ATprefix),
	1845: uint16(anon_sym_ATbase),
	1846: uint16(anon_sym_LPAREN),
	1847: uint16(anon_sym_LT),
	1848: uint16(anon_sym_COLON),
	1849: uint16(anon_sym__COLON),
	1850: uint16(sym_anon),
	1851: uint16(11),
	1852: uint16(13),
	1853: uint16(1),
	1854: uint16(sym_comment),
	1855: uint16(27),
	1856: uint16(1),
	1857: uint16(anon_sym_LT),
	1858: uint16(35),
	1859: uint16(1),
	1860: uint16(sym_pn_prefix),
	1861: uint16(53),
	1862: uint16(1),
	1863: uint16(anon_sym_COLON),
	1864: uint16(300),
	1865: uint16(1),
	1866: uint16(anon_sym_a),
	1867: uint16(5),
	1868: uint16(1),
	1869: uint16(sym_predicate),
	1870: uint16(35),
	1871: uint16(1),
	1872: uint16(sym_namespace),
	1873: uint16(80),
	1874: uint16(1),
	1875: uint16(sym_property),
	1876: uint16(94),
	1877: uint16(1),
	1878: uint16(sym_property_list),
	1879: uint16(346),
	1880: uint16(2),
	1881: uint16(anon_sym_RBRACE),
	1882: uint16(anon_sym_DOT),
	1883: uint16(37),
	1884: uint16(3),
	1885: uint16(sym_iri_reference),
	1886: uint16(sym__iri),
	1887: uint16(sym_prefixed_name),
	1888: uint16(8),
	1889: uint16(13),
	1890: uint16(1),
	1891: uint16(sym_comment),
	1892: uint16(27),
	1893: uint16(1),
	1894: uint16(anon_sym_LT),
	1895: uint16(31),
	1896: uint16(1),
	1897: uint16(anon_sym__COLON),
	1898: uint16(108),
	1899: uint16(1),
	1900: uint16(anon_sym_COLON),
	1901: uint16(348),
	1902: uint16(1),
	1903: uint16(sym_pn_prefix),
	1904: uint16(350),
	1905: uint16(1),
	1906: uint16(sym_anon),
	1907: uint16(66),
	1908: uint16(1),
	1909: uint16(sym_namespace),
	1910: uint16(103),
	1911: uint16(5),
	1912: uint16(sym_iri_reference),
	1913: uint16(sym__iri),
	1914: uint16(sym_prefixed_name),
	1915: uint16(sym__blank_node),
	1916: uint16(sym_blank_node_label),
	1917: uint16(10),
	1918: uint16(13),
	1919: uint16(1),
	1920: uint16(sym_comment),
	1921: uint16(27),
	1922: uint16(1),
	1923: uint16(anon_sym_LT),
	1924: uint16(35),
	1925: uint16(1),
	1926: uint16(sym_pn_prefix),
	1927: uint16(53),
	1928: uint16(1),
	1929: uint16(anon_sym_COLON),
	1930: uint16(300),
	1931: uint16(1),
	1932: uint16(anon_sym_a),
	1933: uint16(5),
	1934: uint16(1),
	1935: uint16(sym_predicate),
	1936: uint16(35),
	1937: uint16(1),
	1938: uint16(sym_namespace),
	1939: uint16(80),
	1940: uint16(1),
	1941: uint16(sym_property),
	1942: uint16(94),
	1943: uint16(1),
	1944: uint16(sym_property_list),
	1945: uint16(37),
	1946: uint16(3),
	1947: uint16(sym_iri_reference),
	1948: uint16(sym__iri),
	1949: uint16(sym_prefixed_name),
	1950: uint16(10),
	1951: uint16(13),
	1952: uint16(1),
	1953: uint16(sym_comment),
	1954: uint16(27),
	1955: uint16(1),
	1956: uint16(anon_sym_LT),
	1957: uint16(35),
	1958: uint16(1),
	1959: uint16(sym_pn_prefix),
	1960: uint16(53),
	1961: uint16(1),
	1962: uint16(anon_sym_COLON),
	1963: uint16(300),
	1964: uint16(1),
	1965: uint16(anon_sym_a),
	1966: uint16(5),
	1967: uint16(1),
	1968: uint16(sym_predicate),
	1969: uint16(35),
	1970: uint16(1),
	1971: uint16(sym_namespace),
	1972: uint16(80),
	1973: uint16(1),
	1974: uint16(sym_property),
	1975: uint16(112),
	1976: uint16(1),
	1977: uint16(sym_property_list),
	1978: uint16(37),
	1979: uint16(3),
	1980: uint16(sym_iri_reference),
	1981: uint16(sym__iri),
	1982: uint16(sym_prefixed_name),
	1983: uint16(6),
	1984: uint16(3),
	1985: uint16(1),
	1986: uint16(sym_comment),
	1987: uint16(356),
	1988: uint16(1),
	1989: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	1990: uint16(358),
	1991: uint16(1),
	1992: uint16(aux_sym__string_literal_long_quote_token1),
	1993: uint16(59),
	1994: uint16(1),
	1995: uint16(aux_sym__string_literal_long_quote_repeat1),
	1996: uint16(352),
	1997: uint16(2),
	1998: uint16(anon_sym_DQUOTE),
	1999: uint16(anon_sym_DQUOTE_DQUOTE),
	2000: uint16(354),
	2001: uint16(2),
	2002: uint16(aux_sym__string_literal_quote_token2),
	2003: uint16(sym_echar),
	2004: uint16(6),
	2005: uint16(13),
	2006: uint16(1),
	2007: uint16(sym_comment),
	2008: uint16(27),
	2009: uint16(1),
	2010: uint16(anon_sym_LT),
	2011: uint16(108),
	2012: uint16(1),
	2013: uint16(anon_sym_COLON),
	2014: uint16(348),
	2015: uint16(1),
	2016: uint16(sym_pn_prefix),
	2017: uint16(66),
	2018: uint16(1),
	2019: uint16(sym_namespace),
	2020: uint16(27),
	2021: uint16(3),
	2022: uint16(sym_iri_reference),
	2023: uint16(sym__iri),
	2024: uint16(sym_prefixed_name),
	2025: uint16(4),
	2026: uint16(13),
	2027: uint16(1),
	2028: uint16(sym_comment),
	2029: uint16(274),
	2030: uint16(1),
	2031: uint16(sym_lang_tag),
	2032: uint16(360),
	2033: uint16(1),
	2034: uint16(anon_sym_CARET_CARET),
	2035: uint16(270),
	2036: uint16(5),
	2037: uint16(anon_sym_RBRACE),
	2038: uint16(anon_sym_DOT),
	2039: uint16(anon_sym_SEMI),
	2040: uint16(anon_sym_COMMA),
	2041: uint16(anon_sym_RBRACK),
	2042: uint16(6),
	2043: uint16(3),
	2044: uint16(1),
	2045: uint16(sym_comment),
	2046: uint16(368),
	2047: uint16(1),
	2048: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2049: uint16(370),
	2050: uint16(1),
	2051: uint16(aux_sym__string_literal_long_single_quote_token1),
	2052: uint16(58),
	2053: uint16(1),
	2054: uint16(aux_sym__string_literal_long_single_quote_repeat1),
	2055: uint16(362),
	2056: uint16(2),
	2057: uint16(aux_sym__string_literal_quote_token2),
	2058: uint16(sym_echar),
	2059: uint16(365),
	2060: uint16(2),
	2061: uint16(anon_sym_SQUOTE),
	2062: uint16(anon_sym_SQUOTE_SQUOTE),
	2063: uint16(6),
	2064: uint16(3),
	2065: uint16(1),
	2066: uint16(sym_comment),
	2067: uint16(379),
	2068: uint16(1),
	2069: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2070: uint16(381),
	2071: uint16(1),
	2072: uint16(aux_sym__string_literal_long_quote_token1),
	2073: uint16(59),
	2074: uint16(1),
	2075: uint16(aux_sym__string_literal_long_quote_repeat1),
	2076: uint16(373),
	2077: uint16(2),
	2078: uint16(anon_sym_DQUOTE),
	2079: uint16(anon_sym_DQUOTE_DQUOTE),
	2080: uint16(376),
	2081: uint16(2),
	2082: uint16(aux_sym__string_literal_quote_token2),
	2083: uint16(sym_echar),
	2084: uint16(6),
	2085: uint16(13),
	2086: uint16(1),
	2087: uint16(sym_comment),
	2088: uint16(27),
	2089: uint16(1),
	2090: uint16(anon_sym_LT),
	2091: uint16(53),
	2092: uint16(1),
	2093: uint16(anon_sym_COLON),
	2094: uint16(384),
	2095: uint16(1),
	2096: uint16(sym_pn_prefix),
	2097: uint16(35),
	2098: uint16(1),
	2099: uint16(sym_namespace),
	2100: uint16(27),
	2101: uint16(3),
	2102: uint16(sym_iri_reference),
	2103: uint16(sym__iri),
	2104: uint16(sym_prefixed_name),
	2105: uint16(6),
	2106: uint16(3),
	2107: uint16(1),
	2108: uint16(sym_comment),
	2109: uint16(388),
	2110: uint16(1),
	2111: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2112: uint16(390),
	2113: uint16(1),
	2114: uint16(aux_sym__string_literal_long_quote_token1),
	2115: uint16(55),
	2116: uint16(1),
	2117: uint16(aux_sym__string_literal_long_quote_repeat1),
	2118: uint16(352),
	2119: uint16(2),
	2120: uint16(anon_sym_DQUOTE),
	2121: uint16(anon_sym_DQUOTE_DQUOTE),
	2122: uint16(386),
	2123: uint16(2),
	2124: uint16(aux_sym__string_literal_quote_token2),
	2125: uint16(sym_echar),
	2126: uint16(2),
	2127: uint16(13),
	2128: uint16(1),
	2129: uint16(sym_comment),
	2130: uint16(278),
	2131: uint16(7),
	2132: uint16(anon_sym_LBRACE),
	2133: uint16(anon_sym_RBRACE),
	2134: uint16(anon_sym_DOT),
	2135: uint16(anon_sym_SEMI),
	2136: uint16(anon_sym_COMMA),
	2137: uint16(anon_sym_RBRACK),
	2138: uint16(sym_pn_local),
	2139: uint16(6),
	2140: uint16(3),
	2141: uint16(1),
	2142: uint16(sym_comment),
	2143: uint16(396),
	2144: uint16(1),
	2145: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2146: uint16(398),
	2147: uint16(1),
	2148: uint16(aux_sym__string_literal_long_single_quote_token1),
	2149: uint16(65),
	2150: uint16(1),
	2151: uint16(aux_sym__string_literal_long_single_quote_repeat1),
	2152: uint16(392),
	2153: uint16(2),
	2154: uint16(aux_sym__string_literal_quote_token2),
	2155: uint16(sym_echar),
	2156: uint16(394),
	2157: uint16(2),
	2158: uint16(anon_sym_SQUOTE),
	2159: uint16(anon_sym_SQUOTE_SQUOTE),
	2160: uint16(2),
	2161: uint16(13),
	2162: uint16(1),
	2163: uint16(sym_comment),
	2164: uint16(288),
	2165: uint16(7),
	2166: uint16(anon_sym_LBRACE),
	2167: uint16(anon_sym_RBRACE),
	2168: uint16(anon_sym_DOT),
	2169: uint16(anon_sym_SEMI),
	2170: uint16(anon_sym_COMMA),
	2171: uint16(anon_sym_RBRACK),
	2172: uint16(sym_pn_local),
	2173: uint16(6),
	2174: uint16(3),
	2175: uint16(1),
	2176: uint16(sym_comment),
	2177: uint16(402),
	2178: uint16(1),
	2179: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2180: uint16(404),
	2181: uint16(1),
	2182: uint16(aux_sym__string_literal_long_single_quote_token1),
	2183: uint16(58),
	2184: uint16(1),
	2185: uint16(aux_sym__string_literal_long_single_quote_repeat1),
	2186: uint16(394),
	2187: uint16(2),
	2188: uint16(anon_sym_SQUOTE),
	2189: uint16(anon_sym_SQUOTE_SQUOTE),
	2190: uint16(400),
	2191: uint16(2),
	2192: uint16(aux_sym__string_literal_quote_token2),
	2193: uint16(sym_echar),
	2194: uint16(3),
	2195: uint16(13),
	2196: uint16(1),
	2197: uint16(sym_comment),
	2198: uint16(406),
	2199: uint16(1),
	2200: uint16(sym_pn_local),
	2201: uint16(282),
	2202: uint16(6),
	2203: uint16(anon_sym_LBRACE),
	2204: uint16(anon_sym_RBRACE),
	2205: uint16(anon_sym_DOT),
	2206: uint16(anon_sym_SEMI),
	2207: uint16(anon_sym_COMMA),
	2208: uint16(anon_sym_RBRACK),
	2209: uint16(4),
	2210: uint16(13),
	2211: uint16(1),
	2212: uint16(sym_comment),
	2213: uint16(284),
	2214: uint16(1),
	2215: uint16(sym_pn_local),
	2216: uint16(282),
	2217: uint16(2),
	2218: uint16(anon_sym_LBRACE),
	2219: uint16(anon_sym_LT),
	2220: uint16(280),
	2221: uint16(3),
	2222: uint16(anon_sym_a),
	2223: uint16(anon_sym_COLON),
	2224: uint16(sym_pn_prefix),
	2225: uint16(3),
	2226: uint16(3),
	2227: uint16(1),
	2228: uint16(sym_comment),
	2229: uint16(408),
	2230: uint16(1),
	2231: uint16(aux_sym__string_literal_long_single_quote_token1),
	2232: uint16(368),
	2233: uint16(5),
	2234: uint16(aux_sym__string_literal_quote_token2),
	2235: uint16(anon_sym_SQUOTE),
	2236: uint16(anon_sym_SQUOTE_SQUOTE_SQUOTE),
	2237: uint16(anon_sym_SQUOTE_SQUOTE),
	2238: uint16(sym_echar),
	2239: uint16(4),
	2240: uint16(13),
	2241: uint16(1),
	2242: uint16(sym_comment),
	2243: uint16(412),
	2244: uint16(1),
	2245: uint16(anon_sym_COMMA),
	2246: uint16(73),
	2247: uint16(1),
	2248: uint16(aux_sym_object_list_repeat1),
	2249: uint16(410),
	2250: uint16(4),
	2251: uint16(anon_sym_RBRACE),
	2252: uint16(anon_sym_DOT),
	2253: uint16(anon_sym_SEMI),
	2254: uint16(anon_sym_RBRACK),
	2255: uint16(3),
	2256: uint16(13),
	2257: uint16(1),
	2258: uint16(sym_comment),
	2259: uint16(278),
	2260: uint16(2),
	2261: uint16(anon_sym_LBRACE),
	2262: uint16(anon_sym_LT),
	2263: uint16(276),
	2264: uint16(4),
	2265: uint16(anon_sym_a),
	2266: uint16(anon_sym_COLON),
	2267: uint16(sym_pn_prefix),
	2268: uint16(sym_pn_local),
	2269: uint16(4),
	2270: uint16(13),
	2271: uint16(1),
	2272: uint16(sym_comment),
	2273: uint16(416),
	2274: uint16(1),
	2275: uint16(anon_sym_COMMA),
	2276: uint16(71),
	2277: uint16(1),
	2278: uint16(aux_sym_object_list_repeat1),
	2279: uint16(414),
	2280: uint16(4),
	2281: uint16(anon_sym_RBRACE),
	2282: uint16(anon_sym_DOT),
	2283: uint16(anon_sym_SEMI),
	2284: uint16(anon_sym_RBRACK),
	2285: uint16(3),
	2286: uint16(13),
	2287: uint16(1),
	2288: uint16(sym_comment),
	2289: uint16(288),
	2290: uint16(2),
	2291: uint16(anon_sym_LBRACE),
	2292: uint16(anon_sym_LT),
	2293: uint16(286),
	2294: uint16(4),
	2295: uint16(anon_sym_a),
	2296: uint16(anon_sym_COLON),
	2297: uint16(sym_pn_prefix),
	2298: uint16(sym_pn_local),
	2299: uint16(4),
	2300: uint16(13),
	2301: uint16(1),
	2302: uint16(sym_comment),
	2303: uint16(412),
	2304: uint16(1),
	2305: uint16(anon_sym_COMMA),
	2306: uint16(71),
	2307: uint16(1),
	2308: uint16(aux_sym_object_list_repeat1),
	2309: uint16(419),
	2310: uint16(4),
	2311: uint16(anon_sym_RBRACE),
	2312: uint16(anon_sym_DOT),
	2313: uint16(anon_sym_SEMI),
	2314: uint16(anon_sym_RBRACK),
	2315: uint16(3),
	2316: uint16(3),
	2317: uint16(1),
	2318: uint16(sym_comment),
	2319: uint16(421),
	2320: uint16(1),
	2321: uint16(aux_sym__string_literal_long_quote_token1),
	2322: uint16(379),
	2323: uint16(5),
	2324: uint16(anon_sym_DQUOTE),
	2325: uint16(aux_sym__string_literal_quote_token2),
	2326: uint16(anon_sym_DQUOTE_DQUOTE_DQUOTE),
	2327: uint16(anon_sym_DQUOTE_DQUOTE),
	2328: uint16(sym_echar),
	2329: uint16(4),
	2330: uint16(13),
	2331: uint16(1),
	2332: uint16(sym_comment),
	2333: uint16(425),
	2334: uint16(1),
	2335: uint16(anon_sym_SEMI),
	2336: uint16(75),
	2337: uint16(1),
	2338: uint16(aux_sym_property_list_repeat1),
	2339: uint16(423),
	2340: uint16(3),
	2341: uint16(anon_sym_RBRACE),
	2342: uint16(anon_sym_DOT),
	2343: uint16(anon_sym_RBRACK),
	2344: uint16(5),
	2345: uint16(3),
	2346: uint16(1),
	2347: uint16(sym_comment),
	2348: uint16(428),
	2349: uint16(1),
	2350: uint16(anon_sym_DQUOTE),
	2351: uint16(430),
	2352: uint16(1),
	2353: uint16(aux_sym__string_literal_quote_token1),
	2354: uint16(76),
	2355: uint16(1),
	2356: uint16(aux_sym__string_literal_quote_repeat1),
	2357: uint16(433),
	2358: uint16(2),
	2359: uint16(aux_sym__string_literal_quote_token2),
	2360: uint16(sym_echar),
	2361: uint16(2),
	2362: uint16(13),
	2363: uint16(1),
	2364: uint16(sym_comment),
	2365: uint16(414),
	2366: uint16(5),
	2367: uint16(anon_sym_RBRACE),
	2368: uint16(anon_sym_DOT),
	2369: uint16(anon_sym_SEMI),
	2370: uint16(anon_sym_COMMA),
	2371: uint16(anon_sym_RBRACK),
	2372: uint16(5),
	2373: uint16(3),
	2374: uint16(1),
	2375: uint16(sym_comment),
	2376: uint16(436),
	2377: uint16(1),
	2378: uint16(anon_sym_DQUOTE),
	2379: uint16(438),
	2380: uint16(1),
	2381: uint16(aux_sym__string_literal_quote_token1),
	2382: uint16(76),
	2383: uint16(1),
	2384: uint16(aux_sym__string_literal_quote_repeat1),
	2385: uint16(440),
	2386: uint16(2),
	2387: uint16(aux_sym__string_literal_quote_token2),
	2388: uint16(sym_echar),
	2389: uint16(4),
	2390: uint16(13),
	2391: uint16(1),
	2392: uint16(sym_comment),
	2393: uint16(444),
	2394: uint16(1),
	2395: uint16(anon_sym_LBRACE),
	2396: uint16(442),
	2397: uint16(2),
	2398: uint16(anon_sym_a),
	2399: uint16(sym_pn_prefix),
	2400: uint16(446),
	2401: uint16(2),
	2402: uint16(anon_sym_LT),
	2403: uint16(anon_sym_COLON),
	2404: uint16(4),
	2405: uint16(13),
	2406: uint16(1),
	2407: uint16(sym_comment),
	2408: uint16(450),
	2409: uint16(1),
	2410: uint16(anon_sym_SEMI),
	2411: uint16(81),
	2412: uint16(1),
	2413: uint16(aux_sym_property_list_repeat1),
	2414: uint16(448),
	2415: uint16(3),
	2416: uint16(anon_sym_RBRACE),
	2417: uint16(anon_sym_DOT),
	2418: uint16(anon_sym_RBRACK),
	2419: uint16(4),
	2420: uint16(13),
	2421: uint16(1),
	2422: uint16(sym_comment),
	2423: uint16(450),
	2424: uint16(1),
	2425: uint16(anon_sym_SEMI),
	2426: uint16(75),
	2427: uint16(1),
	2428: uint16(aux_sym_property_list_repeat1),
	2429: uint16(452),
	2430: uint16(3),
	2431: uint16(anon_sym_RBRACE),
	2432: uint16(anon_sym_DOT),
	2433: uint16(anon_sym_RBRACK),
	2434: uint16(5),
	2435: uint16(3),
	2436: uint16(1),
	2437: uint16(sym_comment),
	2438: uint16(456),
	2439: uint16(1),
	2440: uint16(anon_sym_SQUOTE),
	2441: uint16(458),
	2442: uint16(1),
	2443: uint16(aux_sym__string_literal_single_quote_token1),
	2444: uint16(84),
	2445: uint16(1),
	2446: uint16(aux_sym__string_literal_single_quote_repeat1),
	2447: uint16(454),
	2448: uint16(2),
	2449: uint16(aux_sym__string_literal_quote_token2),
	2450: uint16(sym_echar),
	2451: uint16(5),
	2452: uint16(3),
	2453: uint16(1),
	2454: uint16(sym_comment),
	2455: uint16(460),
	2456: uint16(1),
	2457: uint16(anon_sym_DQUOTE),
	2458: uint16(462),
	2459: uint16(1),
	2460: uint16(aux_sym__string_literal_quote_token1),
	2461: uint16(78),
	2462: uint16(1),
	2463: uint16(aux_sym__string_literal_quote_repeat1),
	2464: uint16(464),
	2465: uint16(2),
	2466: uint16(aux_sym__string_literal_quote_token2),
	2467: uint16(sym_echar),
	2468: uint16(5),
	2469: uint16(3),
	2470: uint16(1),
	2471: uint16(sym_comment),
	2472: uint16(469),
	2473: uint16(1),
	2474: uint16(anon_sym_SQUOTE),
	2475: uint16(471),
	2476: uint16(1),
	2477: uint16(aux_sym__string_literal_single_quote_token1),
	2478: uint16(84),
	2479: uint16(1),
	2480: uint16(aux_sym__string_literal_single_quote_repeat1),
	2481: uint16(466),
	2482: uint16(2),
	2483: uint16(aux_sym__string_literal_quote_token2),
	2484: uint16(sym_echar),
	2485: uint16(5),
	2486: uint16(3),
	2487: uint16(1),
	2488: uint16(sym_comment),
	2489: uint16(476),
	2490: uint16(1),
	2491: uint16(anon_sym_SQUOTE),
	2492: uint16(478),
	2493: uint16(1),
	2494: uint16(aux_sym__string_literal_single_quote_token1),
	2495: uint16(82),
	2496: uint16(1),
	2497: uint16(aux_sym__string_literal_single_quote_repeat1),
	2498: uint16(474),
	2499: uint16(2),
	2500: uint16(aux_sym__string_literal_quote_token2),
	2501: uint16(sym_echar),
	2502: uint16(3),
	2503: uint16(13),
	2504: uint16(1),
	2505: uint16(sym_comment),
	2506: uint16(442),
	2507: uint16(2),
	2508: uint16(anon_sym_a),
	2509: uint16(sym_pn_prefix),
	2510: uint16(446),
	2511: uint16(2),
	2512: uint16(anon_sym_LT),
	2513: uint16(anon_sym_COLON),
	2514: uint16(2),
	2515: uint16(13),
	2516: uint16(1),
	2517: uint16(sym_comment),
	2518: uint16(480),
	2519: uint16(4),
	2520: uint16(anon_sym_RBRACE),
	2521: uint16(anon_sym_DOT),
	2522: uint16(anon_sym_SEMI),
	2523: uint16(anon_sym_RBRACK),
	2524: uint16(2),
	2525: uint16(13),
	2526: uint16(1),
	2527: uint16(sym_comment),
	2528: uint16(423),
	2529: uint16(4),
	2530: uint16(anon_sym_RBRACE),
	2531: uint16(anon_sym_DOT),
	2532: uint16(anon_sym_SEMI),
	2533: uint16(anon_sym_RBRACK),
	2534: uint16(4),
	2535: uint16(13),
	2536: uint16(1),
	2537: uint16(sym_comment),
	2538: uint16(29),
	2539: uint16(1),
	2540: uint16(anon_sym_COLON),
	2541: uint16(233),
	2542: uint16(1),
	2543: uint16(sym_pn_prefix),
	2544: uint16(100),
	2545: uint16(1),
	2546: uint16(sym_namespace),
	2547: uint16(4),
	2548: uint16(13),
	2549: uint16(1),
	2550: uint16(sym_comment),
	2551: uint16(29),
	2552: uint16(1),
	2553: uint16(anon_sym_COLON),
	2554: uint16(233),
	2555: uint16(1),
	2556: uint16(sym_pn_prefix),
	2557: uint16(101),
	2558: uint16(1),
	2559: uint16(sym_namespace),
	2560: uint16(3),
	2561: uint16(3),
	2562: uint16(1),
	2563: uint16(sym_comment),
	2564: uint16(484),
	2565: uint16(1),
	2566: uint16(aux_sym__string_literal_long_single_quote_token1),
	2567: uint16(482),
	2568: uint16(2),
	2569: uint16(aux_sym__string_literal_quote_token2),
	2570: uint16(sym_echar),
	2571: uint16(3),
	2572: uint16(3),
	2573: uint16(1),
	2574: uint16(sym_comment),
	2575: uint16(488),
	2576: uint16(1),
	2577: uint16(aux_sym__string_literal_long_quote_token1),
	2578: uint16(486),
	2579: uint16(2),
	2580: uint16(aux_sym__string_literal_quote_token2),
	2581: uint16(sym_echar),
	2582: uint16(3),
	2583: uint16(13),
	2584: uint16(1),
	2585: uint16(sym_comment),
	2586: uint16(490),
	2587: uint16(1),
	2588: uint16(anon_sym_RBRACE),
	2589: uint16(492),
	2590: uint16(1),
	2591: uint16(anon_sym_DOT),
	2592: uint16(2),
	2593: uint16(13),
	2594: uint16(1),
	2595: uint16(sym_comment),
	2596: uint16(494),
	2597: uint16(2),
	2598: uint16(anon_sym_RBRACE),
	2599: uint16(anon_sym_DOT),
	2600: uint16(3),
	2601: uint16(13),
	2602: uint16(1),
	2603: uint16(sym_comment),
	2604: uint16(27),
	2605: uint16(1),
	2606: uint16(anon_sym_LT),
	2607: uint16(106),
	2608: uint16(1),
	2609: uint16(sym_iri_reference),
	2610: uint16(3),
	2611: uint16(13),
	2612: uint16(1),
	2613: uint16(sym_comment),
	2614: uint16(27),
	2615: uint16(1),
	2616: uint16(anon_sym_LT),
	2617: uint16(46),
	2618: uint16(1),
	2619: uint16(sym_iri_reference),
	2620: uint16(3),
	2621: uint16(13),
	2622: uint16(1),
	2623: uint16(sym_comment),
	2624: uint16(492),
	2625: uint16(1),
	2626: uint16(anon_sym_DOT),
	2627: uint16(496),
	2628: uint16(1),
	2629: uint16(anon_sym_RBRACE),
	2630: uint16(3),
	2631: uint16(3),
	2632: uint16(1),
	2633: uint16(sym_comment),
	2634: uint16(498),
	2635: uint16(1),
	2636: uint16(anon_sym_POUND),
	2637: uint16(500),
	2638: uint16(1),
	2639: uint16(aux_sym_iri_reference_token1),
	2640: uint16(3),
	2641: uint16(13),
	2642: uint16(1),
	2643: uint16(sym_comment),
	2644: uint16(235),
	2645: uint16(1),
	2646: uint16(anon_sym_RBRACE),
	2647: uint16(492),
	2648: uint16(1),
	2649: uint16(anon_sym_DOT),
	2650: uint16(3),
	2651: uint16(13),
	2652: uint16(1),
	2653: uint16(sym_comment),
	2654: uint16(27),
	2655: uint16(1),
	2656: uint16(anon_sym_LT),
	2657: uint16(42),
	2658: uint16(1),
	2659: uint16(sym_iri_reference),
	2660: uint16(3),
	2661: uint16(13),
	2662: uint16(1),
	2663: uint16(sym_comment),
	2664: uint16(27),
	2665: uint16(1),
	2666: uint16(anon_sym_LT),
	2667: uint16(113),
	2668: uint16(1),
	2669: uint16(sym_iri_reference),
	2670: uint16(3),
	2671: uint16(13),
	2672: uint16(1),
	2673: uint16(sym_comment),
	2674: uint16(264),
	2675: uint16(1),
	2676: uint16(anon_sym_RBRACE),
	2677: uint16(492),
	2678: uint16(1),
	2679: uint16(anon_sym_DOT),
	2680: uint16(2),
	2681: uint16(13),
	2682: uint16(1),
	2683: uint16(sym_comment),
	2684: uint16(502),
	2685: uint16(1),
	2686: uint16(anon_sym_LBRACE),
	2687: uint16(2),
	2688: uint16(13),
	2689: uint16(1),
	2690: uint16(sym_comment),
	2691: uint16(504),
	2692: uint16(1),
	2693: uint16(anon_sym_GT),
	2694: uint16(2),
	2695: uint16(13),
	2696: uint16(1),
	2697: uint16(sym_comment),
	2698: uint16(506),
	2699: uint16(1),
	2700: uint16(anon_sym_GT),
	2701: uint16(2),
	2702: uint16(13),
	2703: uint16(1),
	2704: uint16(sym_comment),
	2705: uint16(508),
	2706: uint16(1),
	2707: uint16(anon_sym_DOT),
	2708: uint16(2),
	2709: uint16(13),
	2710: uint16(1),
	2711: uint16(sym_comment),
	2712: uint16(492),
	2713: uint16(1),
	2714: uint16(anon_sym_DOT),
	2715: uint16(2),
	2716: uint16(13),
	2717: uint16(1),
	2718: uint16(sym_comment),
	2719: uint16(510),
	2720: uint16(1),
	2721: uint16(anon_sym_LBRACE),
	2722: uint16(2),
	2723: uint16(13),
	2724: uint16(1),
	2725: uint16(sym_comment),
	2726: uint16(512),
	2727: uint16(1),
	2729: uint16(2),
	2730: uint16(13),
	2731: uint16(1),
	2732: uint16(sym_comment),
	2733: uint16(514),
	2734: uint16(1),
	2735: uint16(anon_sym_COLON),
	2736: uint16(2),
	2737: uint16(13),
	2738: uint16(1),
	2739: uint16(sym_comment),
	2740: uint16(516),
	2741: uint16(1),
	2742: uint16(aux_sym_blank_node_label_token1),
	2743: uint16(2),
	2744: uint16(13),
	2745: uint16(1),
	2746: uint16(sym_comment),
	2747: uint16(518),
	2748: uint16(1),
	2749: uint16(anon_sym_RBRACK),
	2750: uint16(2),
	2751: uint16(13),
	2752: uint16(1),
	2753: uint16(sym_comment),
	2754: uint16(520),
	2755: uint16(1),
	2756: uint16(anon_sym_DOT),
	2757: uint16(2),
	2758: uint16(13),
	2759: uint16(1),
	2760: uint16(sym_comment),
	2761: uint16(522),
	2762: uint16(1),
	2763: uint16(anon_sym_COLON),
	2764: uint16(2),
	2765: uint16(13),
	2766: uint16(1),
	2767: uint16(sym_comment),
	2768: uint16(524),
	2769: uint16(1),
	2770: uint16(anon_sym_RPAREN),
	2771: uint16(2),
	2772: uint16(13),
	2773: uint16(1),
	2774: uint16(sym_comment),
	2775: uint16(526),
	2776: uint16(1),
	2777: uint16(anon_sym_COLON),
	2778: uint16(2),
	2779: uint16(3),
	2780: uint16(1),
	2781: uint16(sym_comment),
	2782: uint16(528),
	2783: uint16(1),
	2784: uint16(aux_sym_iri_reference_token1),
}

var ts_small_parse_table_map = [116]uint32_t{
	1:   uint32(79),
	2:   uint32(155),
	3:   uint32(231),
	4:   uint32(306),
	5:   uint32(378),
	6:   uint32(461),
	7:   uint32(544),
	8:   uint32(582),
	9:   uint32(620),
	10:  uint32(652),
	11:  uint32(684),
	12:  uint32(716),
	13:  uint32(748),
	14:  uint32(780),
	15:  uint32(812),
	16:  uint32(844),
	17:  uint32(876),
	18:  uint32(908),
	19:  uint32(940),
	20:  uint32(972),
	21:  uint32(1003),
	22:  uint32(1034),
	23:  uint32(1065),
	24:  uint32(1095),
	25:  uint32(1125),
	26:  uint32(1155),
	27:  uint32(1207),
	28:  uint32(1259),
	29:  uint32(1311),
	30:  uint32(1363),
	31:  uint32(1415),
	32:  uint32(1446),
	33:  uint32(1472),
	34:  uint32(1500),
	35:  uint32(1526),
	36:  uint32(1550),
	37:  uint32(1573),
	38:  uint32(1609),
	39:  uint32(1631),
	40:  uint32(1653),
	41:  uint32(1675),
	42:  uint32(1697),
	43:  uint32(1719),
	44:  uint32(1741),
	45:  uint32(1763),
	46:  uint32(1785),
	47:  uint32(1807),
	48:  uint32(1829),
	49:  uint32(1851),
	50:  uint32(1888),
	51:  uint32(1917),
	52:  uint32(1950),
	53:  uint32(1983),
	54:  uint32(2004),
	55:  uint32(2025),
	56:  uint32(2042),
	57:  uint32(2063),
	58:  uint32(2084),
	59:  uint32(2105),
	60:  uint32(2126),
	61:  uint32(2139),
	62:  uint32(2160),
	63:  uint32(2173),
	64:  uint32(2194),
	65:  uint32(2209),
	66:  uint32(2225),
	67:  uint32(2239),
	68:  uint32(2255),
	69:  uint32(2269),
	70:  uint32(2285),
	71:  uint32(2299),
	72:  uint32(2315),
	73:  uint32(2329),
	74:  uint32(2344),
	75:  uint32(2361),
	76:  uint32(2372),
	77:  uint32(2389),
	78:  uint32(2404),
	79:  uint32(2419),
	80:  uint32(2434),
	81:  uint32(2451),
	82:  uint32(2468),
	83:  uint32(2485),
	84:  uint32(2502),
	85:  uint32(2514),
	86:  uint32(2524),
	87:  uint32(2534),
	88:  uint32(2547),
	89:  uint32(2560),
	90:  uint32(2571),
	91:  uint32(2582),
	92:  uint32(2592),
	93:  uint32(2600),
	94:  uint32(2610),
	95:  uint32(2620),
	96:  uint32(2630),
	97:  uint32(2640),
	98:  uint32(2650),
	99:  uint32(2660),
	100: uint32(2670),
	101: uint32(2680),
	102: uint32(2687),
	103: uint32(2694),
	104: uint32(2701),
	105: uint32(2708),
	106: uint32(2715),
	107: uint32(2722),
	108: uint32(2729),
	109: uint32(2736),
	110: uint32(2743),
	111: uint32(2750),
	112: uint32(2757),
	113: uint32(2764),
	114: uint32(2771),
	115: uint32(2778),
}

var ts_parse_actions = [530]TSParseActionEntry{
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
		Fcount: uint8(1),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(89)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(54)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(98)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(111)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(79)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(116)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(61)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(116)),
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
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(54)),
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
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
	})))),
	66: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	67: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(98)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	73: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(3)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(83)),
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
		Fcount: uint8(2),
	}})),
	79: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(85)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	85: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
	})))),
	86: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(61)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	88: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
	})))),
	89: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(26)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(36)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_object_collection_repeat1),
	})))),
	95: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_object_collection),
	})))),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(114)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(64)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_document),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	117: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(52)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	128: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(90)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(95)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(96)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	137: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	138: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(54)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	143: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	144: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	147: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	148: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	149: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(72)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(111)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(79)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_iri_reference),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_iri_reference),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_iri_reference),
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
		Fcount: uint8(1),
	}})),
	164: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_iri_reference),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_prefixed_name),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_prefixed_name),
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
		Fsymbol:      uint16(sym__string_literal_long_single_quote),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__string_literal_long_single_quote),
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
		Fsymbol:      uint16(sym__string_literal_long_quote),
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
		Fsymbol:      uint16(sym__string_literal_long_quote),
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
		Fsymbol:      uint16(sym__string_literal_long_single_quote),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__string_literal_long_single_quote),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__string_literal_single_quote),
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
		Fsymbol:      uint16(sym__string_literal_single_quote),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__string_literal_quote),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__string_literal_quote),
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
		Fcount: uint8(1),
	}})),
	190: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_blank_node_label),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_blank_node_label),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_string),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_string),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__string_literal_long_quote),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__string_literal_long_quote),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__string_literal_single_quote),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__string_literal_single_quote),
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
		Fsymbol:      uint16(sym__string_literal_quote),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__string_literal_quote),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_collection),
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
		Fsymbol:      uint16(sym_collection),
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
		Fcount: uint8(1),
	}})),
	214: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_blank_node_property_list),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_blank_node_property_list),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_collection),
	})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_collection),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_rdf_literal),
		Fproduction_id: uint16(1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_rdf_literal),
		Fproduction_id: uint16(1),
	})))),
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
		Fcount: uint8(1),
	}})),
	226: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_boolean_literal),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_boolean_literal),
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
		Fcount: uint8(1),
	}})),
	230: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_rdf_literal),
		Fproduction_id: uint16(3),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	232: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_rdf_literal),
		Fproduction_id: uint16(3),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	240: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_graph_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(110)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
	243: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_graph_repeat1),
	})))),
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
		Fcount: uint8(2),
	}})),
	245: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_graph_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(54)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	248: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_graph_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(2)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	251: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_graph_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(98)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_graph_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(72)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_graph_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(111)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_graph_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(86)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
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
		Fcount: uint8(1),
	}})),
	269: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_rdf_literal),
		Fproduction_id: uint16(1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_rdf_literal),
		Fproduction_id: uint16(1),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
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
		Fsymbol:      uint16(sym_namespace),
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
		Fsymbol:      uint16(sym_namespace),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_prefixed_name),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_prefixed_name),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	287: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_namespace),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	289: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_namespace),
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
		Fcount: uint8(1),
	}})),
	291: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_predicate),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_predicate),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_triple),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_triple),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_property_list_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_directive),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_directive),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_graph),
		Fproduction_id: uint16(2),
	})))),
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
		Fcount: uint8(1),
	}})),
	309: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_graph),
		Fproduction_id: uint16(2),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_sparql_prefix),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_sparql_prefix),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	315: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_base),
	})))),
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
		Fcount: uint8(1),
	}})),
	317: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_base),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	319: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_graph),
		Fproduction_id: uint16(2),
	})))),
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
		Fcount: uint8(1),
	}})),
	321: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_graph),
		Fproduction_id: uint16(2),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	323: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	325: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fsymbol:      uint16(sym_sparql_base),
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
		Fsymbol:      uint16(sym_sparql_base),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_graph),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_graph),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_graph),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_graph),
		Fproduction_id: uint16(2),
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
		Fsymbol:      uint16(sym_prefix_id),
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
		Fsymbol:      uint16(sym_prefix_id),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_graph),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_graph),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__triples),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	349: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	355: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	361: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	363: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_long_single_quote_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(58)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	366: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_long_single_quote_repeat1),
	})))),
	367: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	369: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_long_single_quote_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	371: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_long_single_quote_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(58)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	374: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_long_quote_repeat1),
	})))),
	375: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(92)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	377: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_long_quote_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(59)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_long_quote_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_long_quote_repeat1),
	})))),
	383: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(59)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	389: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	391: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	393: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	397: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	399: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	409: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_long_single_quote_repeat1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_object_list),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_object_list_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	417: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_object_list_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(6)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_object_list),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_long_quote_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_property_list_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_property_list_repeat1),
	})))),
	427: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_quote_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_quote_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(76)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	434: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_quote_repeat1),
	})))),
	435: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(76)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	443: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_subject),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	445: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__label),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_subject),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	449: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_property_list),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	453: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_property_list),
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
		Fcount: uint8(1),
	}})),
	455: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	467: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_single_quote_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(84)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_single_quote_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__string_literal_single_quote_repeat1),
	})))),
	473: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	475: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	481: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_property),
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
		Fcount: uint8(1),
	}})),
	483: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	485: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	487: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	489: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	493: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__triples),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	497: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	499: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__label),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	505: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	509: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	513: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	517: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	523: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	525: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	529: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

func tree_sitter_turtle(tls *libc.TLS) (r uintptr) {
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
	Fkeyword_capture_token:     uint16(sym_pn_prefix),
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

var __ccgo_ts1 = "end\x00pn_prefix\x00{\x00}\x00GRAPH\x00comment\x00.\x00@prefix\x00@base\x00BASE\x00PREFIX\x00;\x00,\x00a\x00[\x00]\x00(\x00)\x00<\x00#\x00iri_reference_token1\x00>\x00integer\x00decimal\x00double\x00\"\x00_string_literal_quote_token1\x00_string_literal_quote_token2\x00'\x00_string_literal_single_quote_token1\x00'''\x00''\x00_string_literal_long_single_quote_token1\x00\"\"\"\x00\"\"\x00_string_literal_long_quote_token1\x00^^\x00true\x00false\x00:\x00_:\x00blank_node_label_token1\x00lang_tag\x00echar\x00anon\x00pn_local\x00document\x00graph\x00_label\x00triple\x00directive\x00prefix_id\x00base\x00sparql_base\x00sparql_prefix\x00_triples\x00property_list\x00property\x00object_list\x00predicate\x00subject\x00_object\x00_literal\x00blank_node_property_list\x00collection\x00object_collection\x00_numeric_literal\x00string\x00iri_reference\x00_string_literal_quote\x00_string_literal_single_quote\x00_string_literal_long_single_quote\x00_string_literal_long_quote\x00rdf_literal\x00boolean_literal\x00_iri\x00prefixed_name\x00_blank_node\x00namespace\x00blank_node_label\x00document_repeat1\x00graph_repeat1\x00property_list_repeat1\x00object_list_repeat1\x00object_collection_repeat1\x00_string_literal_quote_repeat1\x00_string_literal_single_quote_repeat1\x00_string_literal_long_single_quote_repeat1\x00_string_literal_long_quote_repeat1\x00datatype\x00label\x00value\x00"
