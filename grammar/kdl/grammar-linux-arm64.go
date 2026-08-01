// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-kdl/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-kdl -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_kdl

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 4
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXTERNAL_TOKEN_COUNT = 3
const FIELD_COUNT = 1
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
const MAX_ALIAS_SEQUENCE_LENGTH = 10
const PRODUCTION_ID_COUNT = 17
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 304
const SYMBOL_COUNT = 127
const TOKEN_COUNT = 84
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

type wint_t = uint32

type wctype_t = uint64

type locale_t = uintptr

type wctrans_t = uintptr

const _EOF = 0
const MULTI_LINE_COMMENT = 1
const _RAW_STRING = 2

func tree_sitter_kdl_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_kdl_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
}

func tree_sitter_kdl_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_kdl_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
}

func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
}

func tree_sitter_kdl_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var after_star uint8
	var c int32_t
	var closing_hashes, nesting_depth, num_hashes uint32
	_, _, _, _, _ = after_star, c, closing_hashes, nesting_depth, num_hashes
	// check for End-of-file
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(_EOF))) != 0 && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == 0 {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(_EOF)
		advance(tls, lexer)
		return libc.BoolUint8(true1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(_RAW_STRING))) != 0 && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('r') {
		advance(tls, lexer)
		num_hashes = uint32(0)
		for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('#') {
			num_hashes = num_hashes + uint32(1)
			advance(tls, lexer)
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('"') {
			return libc.BoolUint8(false1 != 0)
		}
		advance(tls, lexer)
		for {
			if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
				// Unclosed raw string caused by EOF.
				return libc.BoolUint8(false1 != 0)
			}
			c = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
			advance(tls, lexer)
			if c != int32('"') {
				goto _1
			}
			// Try to match `num_hashes` closing hashes.
			closing_hashes = uint32(0)
			for {
				if closing_hashes == num_hashes {
					goto success
				}
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('#') {
					break
				}
				advance(tls, lexer)
				closing_hashes = closing_hashes + uint32(1)
				goto _2
			_2:
			}
			goto _1
		_1:
		}
		goto success
	success:
		;
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(_RAW_STRING)
		return libc.BoolUint8(true1 != 0)
	}
	// multi-line-comment := '/*' commented-block
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
		advance(tls, lexer)
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('*') {
			return libc.BoolUint8(false1 != 0)
		}
		advance(tls, lexer)
		after_star = libc.BoolUint8(false1 != 0)
		nesting_depth = uint32(1)
		// commented-block := '*/' | (multi-line-comment | '*' | '/' | [^*/]+)
		// commented-block
		for {
			switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
			case int32('\000'):
				return libc.BoolUint8(false1 != 0)
			case int32('*'):
				advance(tls, lexer)
				after_star = libc.BoolUint8(true1 != 0)
			case int32('/'):
				if after_star != 0 {
					advance(tls, lexer)
					after_star = libc.BoolUint8(false1 != 0)
					nesting_depth = nesting_depth - 1
					if nesting_depth == uint32(0) {
						(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(MULTI_LINE_COMMENT)
						return libc.BoolUint8(true1 != 0)
					}
				} else {
					advance(tls, lexer)
					after_star = libc.BoolUint8(false1 != 0)
					if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*') {
						nesting_depth = nesting_depth + 1
						advance(tls, lexer)
					}
				}
			default:
				advance(tls, lexer)
				after_star = libc.BoolUint8(false1 != 0)
				break
			}
			goto _3
		_3:
		}
	}
	return libc.BoolUint8(false1 != 0)
}

type ts_symbol_identifiers = int32

const sym__normal_bare_identifier = 1
const anon_sym_SLASH_DASH = 2
const anon_sym_LBRACE = 3
const anon_sym_RBRACE = 4
const anon_sym_SEMI = 5
const sym__identifier_char = 6
const sym___identifier_char_no_digit = 7
const sym___identifier_char_no_digit_sign = 8
const anon_sym_null = 9
const anon_sym_i8 = 10
const anon_sym_i16 = 11
const anon_sym_i32 = 12
const anon_sym_i64 = 13
const anon_sym_u8 = 14
const anon_sym_u16 = 15
const anon_sym_u32 = 16
const anon_sym_u64 = 17
const anon_sym_isize = 18
const anon_sym_usize = 19
const anon_sym_f32 = 20
const anon_sym_f64 = 21
const anon_sym_decimal64 = 22
const anon_sym_decimal128 = 23
const anon_sym_date_DASHtime = 24
const anon_sym_time = 25
const anon_sym_date = 26
const anon_sym_duration = 27
const anon_sym_decimal = 28
const anon_sym_currency = 29
const anon_sym_country_DASH2 = 30
const anon_sym_country_DASH3 = 31
const anon_sym_country_DASHsubdivision = 32
const anon_sym_email = 33
const anon_sym_idn_DASHemail = 34
const anon_sym_hostname = 35
const anon_sym_idn_DASHhostname = 36
const anon_sym_ipv4 = 37
const anon_sym_ipv6 = 38
const anon_sym_url = 39
const anon_sym_url_DASHreference = 40
const anon_sym_irl = 41
const anon_sym_iri_DASHreference = 42
const anon_sym_url_DASHtemplate = 43
const anon_sym_uuid = 44
const anon_sym_regex = 45
const anon_sym_base64 = 46
const anon_sym_EQ = 47
const anon_sym_LPAREN = 48
const anon_sym_RPAREN = 49
const anon_sym_DQUOTE = 50
const aux_sym__escaped_string_token1 = 51
const sym_escape = 52
const sym__hex_digit = 53
const anon_sym_DOT = 54
const anon_sym_e = 55
const anon_sym_E = 56
const anon_sym__ = 57
const sym__digit = 58
const anon_sym_PLUS = 59
const anon_sym_DASH = 60
const anon_sym_0x = 61
const anon_sym_0o = 62
const aux_sym__octal_token1 = 63
const anon_sym_0b = 64
const anon_sym_0 = 65
const anon_sym_1 = 66
const anon_sym_true = 67
const anon_sym_false = 68
const anon_sym_BSLASH = 69
const aux_sym__newline_token1 = 70
const aux_sym__newline_token2 = 71
const aux_sym__newline_token3 = 72
const aux_sym__newline_token4 = 73
const aux_sym__newline_token5 = 74
const aux_sym__newline_token6 = 75
const aux_sym__newline_token7 = 76
const sym__bom = 77
const sym__unicode_space = 78
const anon_sym_SLASH_SLASH = 79
const aux_sym_single_line_comment_token1 = 80
const sym__eof = 81
const sym_multi_line_comment = 82
const sym__raw_string = 83
const sym_document = 84
const sym_node = 85
const sym_node_field = 86
const sym__node_field_comment = 87
const sym__node_field = 88
const sym_node_children = 89
const sym__node_space = 90
const sym__node_terminator = 91
const sym_identifier = 92
const sym__bare_identifier = 93
const sym_keyword = 94
const sym_annotation_type = 95
const sym_prop = 96
const sym_value = 97
const sym_type = 98
const sym_string = 99
const sym__escaped_string = 100
const sym_number = 101
const sym__decimal = 102
const sym__exponent = 103
const sym__integer = 104
const sym__sign = 105
const sym__hex = 106
const sym__octal = 107
const sym__binary = 108
const sym_boolean = 109
const sym__escline = 110
const sym__linespace = 111
const sym__newline = 112
const sym__ws = 113
const sym_single_line_comment = 114
const aux_sym_document_repeat1 = 115
const aux_sym_document_repeat2 = 116
const aux_sym_node_repeat1 = 117
const aux_sym_node_repeat2 = 118
const aux_sym_node_repeat3 = 119
const aux_sym__bare_identifier_repeat1 = 120
const aux_sym__escaped_string_repeat1 = 121
const aux_sym__integer_repeat1 = 122
const aux_sym__hex_repeat1 = 123
const aux_sym__octal_repeat1 = 124
const aux_sym__binary_repeat1 = 125
const aux_sym_single_line_comment_repeat1 = 126
const alias_sym_decimal = 127
const alias_sym_node_children_comment = 128
const alias_sym_node_field_comment = 129
const alias_sym_string_fragment = 130

var ts_symbol_names = [131]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 28,
	3:   __ccgo_ts + 41,
	4:   __ccgo_ts + 43,
	5:   __ccgo_ts + 45,
	6:   __ccgo_ts + 47,
	7:   __ccgo_ts + 64,
	8:   __ccgo_ts + 91,
	9:   __ccgo_ts + 123,
	10:  __ccgo_ts + 128,
	11:  __ccgo_ts + 131,
	12:  __ccgo_ts + 135,
	13:  __ccgo_ts + 139,
	14:  __ccgo_ts + 143,
	15:  __ccgo_ts + 146,
	16:  __ccgo_ts + 150,
	17:  __ccgo_ts + 154,
	18:  __ccgo_ts + 158,
	19:  __ccgo_ts + 164,
	20:  __ccgo_ts + 170,
	21:  __ccgo_ts + 174,
	22:  __ccgo_ts + 178,
	23:  __ccgo_ts + 188,
	24:  __ccgo_ts + 199,
	25:  __ccgo_ts + 209,
	26:  __ccgo_ts + 214,
	27:  __ccgo_ts + 219,
	28:  __ccgo_ts + 228,
	29:  __ccgo_ts + 236,
	30:  __ccgo_ts + 245,
	31:  __ccgo_ts + 255,
	32:  __ccgo_ts + 265,
	33:  __ccgo_ts + 285,
	34:  __ccgo_ts + 291,
	35:  __ccgo_ts + 301,
	36:  __ccgo_ts + 310,
	37:  __ccgo_ts + 323,
	38:  __ccgo_ts + 328,
	39:  __ccgo_ts + 333,
	40:  __ccgo_ts + 337,
	41:  __ccgo_ts + 351,
	42:  __ccgo_ts + 355,
	43:  __ccgo_ts + 369,
	44:  __ccgo_ts + 382,
	45:  __ccgo_ts + 387,
	46:  __ccgo_ts + 393,
	47:  __ccgo_ts + 400,
	48:  __ccgo_ts + 402,
	49:  __ccgo_ts + 404,
	50:  __ccgo_ts + 406,
	51:  __ccgo_ts + 408,
	52:  __ccgo_ts + 431,
	53:  __ccgo_ts + 438,
	54:  __ccgo_ts + 449,
	55:  __ccgo_ts + 451,
	56:  __ccgo_ts + 453,
	57:  __ccgo_ts + 455,
	58:  __ccgo_ts + 457,
	59:  __ccgo_ts + 464,
	60:  __ccgo_ts + 466,
	61:  __ccgo_ts + 468,
	62:  __ccgo_ts + 471,
	63:  __ccgo_ts + 474,
	64:  __ccgo_ts + 488,
	65:  __ccgo_ts + 491,
	66:  __ccgo_ts + 493,
	67:  __ccgo_ts + 495,
	68:  __ccgo_ts + 500,
	69:  __ccgo_ts + 506,
	70:  __ccgo_ts + 508,
	71:  __ccgo_ts + 524,
	72:  __ccgo_ts + 540,
	73:  __ccgo_ts + 556,
	74:  __ccgo_ts + 572,
	75:  __ccgo_ts + 588,
	76:  __ccgo_ts + 604,
	77:  __ccgo_ts + 620,
	78:  __ccgo_ts + 625,
	79:  __ccgo_ts + 640,
	80:  __ccgo_ts + 643,
	81:  __ccgo_ts + 670,
	82:  __ccgo_ts + 675,
	83:  __ccgo_ts + 694,
	84:  __ccgo_ts + 706,
	85:  __ccgo_ts + 715,
	86:  __ccgo_ts + 720,
	87:  __ccgo_ts + 731,
	88:  __ccgo_ts + 751,
	89:  __ccgo_ts + 763,
	90:  __ccgo_ts + 777,
	91:  __ccgo_ts + 789,
	92:  __ccgo_ts + 806,
	93:  __ccgo_ts + 817,
	94:  __ccgo_ts + 834,
	95:  __ccgo_ts + 842,
	96:  __ccgo_ts + 858,
	97:  __ccgo_ts + 863,
	98:  __ccgo_ts + 869,
	99:  __ccgo_ts + 874,
	100: __ccgo_ts + 881,
	101: __ccgo_ts + 897,
	102: __ccgo_ts + 904,
	103: __ccgo_ts + 913,
	104: __ccgo_ts + 922,
	105: __ccgo_ts + 931,
	106: __ccgo_ts + 937,
	107: __ccgo_ts + 942,
	108: __ccgo_ts + 949,
	109: __ccgo_ts + 957,
	110: __ccgo_ts + 965,
	111: __ccgo_ts + 974,
	112: __ccgo_ts + 985,
	113: __ccgo_ts + 994,
	114: __ccgo_ts + 998,
	115: __ccgo_ts + 1018,
	116: __ccgo_ts + 1035,
	117: __ccgo_ts + 1052,
	118: __ccgo_ts + 1065,
	119: __ccgo_ts + 1078,
	120: __ccgo_ts + 1091,
	121: __ccgo_ts + 1116,
	122: __ccgo_ts + 1140,
	123: __ccgo_ts + 1157,
	124: __ccgo_ts + 1170,
	125: __ccgo_ts + 1185,
	126: __ccgo_ts + 1201,
	127: __ccgo_ts + 228,
	128: __ccgo_ts + 1229,
	129: __ccgo_ts + 1251,
	130: __ccgo_ts + 1270,
}

var ts_symbol_map = [131]TSSymbol{
	1:   uint16(sym__normal_bare_identifier),
	2:   uint16(anon_sym_SLASH_DASH),
	3:   uint16(anon_sym_LBRACE),
	4:   uint16(anon_sym_RBRACE),
	5:   uint16(anon_sym_SEMI),
	6:   uint16(sym__identifier_char),
	7:   uint16(sym___identifier_char_no_digit),
	8:   uint16(sym___identifier_char_no_digit_sign),
	9:   uint16(anon_sym_null),
	10:  uint16(anon_sym_i8),
	11:  uint16(anon_sym_i16),
	12:  uint16(anon_sym_i32),
	13:  uint16(anon_sym_i64),
	14:  uint16(anon_sym_u8),
	15:  uint16(anon_sym_u16),
	16:  uint16(anon_sym_u32),
	17:  uint16(anon_sym_u64),
	18:  uint16(anon_sym_isize),
	19:  uint16(anon_sym_usize),
	20:  uint16(anon_sym_f32),
	21:  uint16(anon_sym_f64),
	22:  uint16(anon_sym_decimal64),
	23:  uint16(anon_sym_decimal128),
	24:  uint16(anon_sym_date_DASHtime),
	25:  uint16(anon_sym_time),
	26:  uint16(anon_sym_date),
	27:  uint16(anon_sym_duration),
	28:  uint16(anon_sym_decimal),
	29:  uint16(anon_sym_currency),
	30:  uint16(anon_sym_country_DASH2),
	31:  uint16(anon_sym_country_DASH3),
	32:  uint16(anon_sym_country_DASHsubdivision),
	33:  uint16(anon_sym_email),
	34:  uint16(anon_sym_idn_DASHemail),
	35:  uint16(anon_sym_hostname),
	36:  uint16(anon_sym_idn_DASHhostname),
	37:  uint16(anon_sym_ipv4),
	38:  uint16(anon_sym_ipv6),
	39:  uint16(anon_sym_url),
	40:  uint16(anon_sym_url_DASHreference),
	41:  uint16(anon_sym_irl),
	42:  uint16(anon_sym_iri_DASHreference),
	43:  uint16(anon_sym_url_DASHtemplate),
	44:  uint16(anon_sym_uuid),
	45:  uint16(anon_sym_regex),
	46:  uint16(anon_sym_base64),
	47:  uint16(anon_sym_EQ),
	48:  uint16(anon_sym_LPAREN),
	49:  uint16(anon_sym_RPAREN),
	50:  uint16(anon_sym_DQUOTE),
	51:  uint16(aux_sym__escaped_string_token1),
	52:  uint16(sym_escape),
	53:  uint16(sym__hex_digit),
	54:  uint16(anon_sym_DOT),
	55:  uint16(anon_sym_e),
	56:  uint16(anon_sym_E),
	57:  uint16(anon_sym__),
	58:  uint16(sym__digit),
	59:  uint16(anon_sym_PLUS),
	60:  uint16(anon_sym_DASH),
	61:  uint16(anon_sym_0x),
	62:  uint16(anon_sym_0o),
	63:  uint16(aux_sym__octal_token1),
	64:  uint16(anon_sym_0b),
	65:  uint16(anon_sym_0),
	66:  uint16(anon_sym_1),
	67:  uint16(anon_sym_true),
	68:  uint16(anon_sym_false),
	69:  uint16(anon_sym_BSLASH),
	70:  uint16(aux_sym__newline_token1),
	71:  uint16(aux_sym__newline_token2),
	72:  uint16(aux_sym__newline_token3),
	73:  uint16(aux_sym__newline_token4),
	74:  uint16(aux_sym__newline_token5),
	75:  uint16(aux_sym__newline_token6),
	76:  uint16(aux_sym__newline_token7),
	77:  uint16(sym__bom),
	78:  uint16(sym__unicode_space),
	79:  uint16(anon_sym_SLASH_SLASH),
	80:  uint16(aux_sym_single_line_comment_token1),
	81:  uint16(sym__eof),
	82:  uint16(sym_multi_line_comment),
	83:  uint16(sym__raw_string),
	84:  uint16(sym_document),
	85:  uint16(sym_node),
	86:  uint16(sym_node_field),
	87:  uint16(sym__node_field_comment),
	88:  uint16(sym__node_field),
	89:  uint16(sym_node_children),
	90:  uint16(sym__node_space),
	91:  uint16(sym__node_terminator),
	92:  uint16(sym_identifier),
	93:  uint16(sym__bare_identifier),
	94:  uint16(sym_keyword),
	95:  uint16(sym_annotation_type),
	96:  uint16(sym_prop),
	97:  uint16(sym_value),
	98:  uint16(sym_type),
	99:  uint16(sym_string),
	100: uint16(sym__escaped_string),
	101: uint16(sym_number),
	102: uint16(sym__decimal),
	103: uint16(sym__exponent),
	104: uint16(sym__integer),
	105: uint16(sym__sign),
	106: uint16(sym__hex),
	107: uint16(sym__octal),
	108: uint16(sym__binary),
	109: uint16(sym_boolean),
	110: uint16(sym__escline),
	111: uint16(sym__linespace),
	112: uint16(sym__newline),
	113: uint16(sym__ws),
	114: uint16(sym_single_line_comment),
	115: uint16(aux_sym_document_repeat1),
	116: uint16(aux_sym_document_repeat2),
	117: uint16(aux_sym_node_repeat1),
	118: uint16(aux_sym_node_repeat2),
	119: uint16(aux_sym_node_repeat3),
	120: uint16(aux_sym__bare_identifier_repeat1),
	121: uint16(aux_sym__escaped_string_repeat1),
	122: uint16(aux_sym__integer_repeat1),
	123: uint16(aux_sym__hex_repeat1),
	124: uint16(aux_sym__octal_repeat1),
	125: uint16(aux_sym__binary_repeat1),
	126: uint16(aux_sym_single_line_comment_repeat1),
	127: uint16(alias_sym_decimal),
	128: uint16(alias_sym_node_children_comment),
	129: uint16(alias_sym_node_field_comment),
	130: uint16(alias_sym_string_fragment),
}

var ts_symbol_metadata = [131]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	2: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	7: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	8: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
	51: {},
	52: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	53: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
	63: {},
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
	70: {},
	71: {},
	72: {},
	73: {},
	74: {},
	75: {},
	76: {},
	77: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	78: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	79: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	80: {},
	81: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	82: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	83: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	88: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	89: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	90: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	91: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	92: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	93: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	101: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	102: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	103: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	104: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	105: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	106: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	111: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	112: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	113: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	114: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
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
	130: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
}

type ts_field_identifiers = int32

const field_children = 1

var ts_field_names = [2]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 1286,
}

var ts_field_map_slices = [17]TSFieldMapSlice{
	2: {
		Flength: uint16(1),
	},
	3: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	7: {
		Findex:  uint16(2),
		Flength: uint16(1),
	},
	8: {
		Findex:  uint16(2),
		Flength: uint16(1),
	},
	11: {
		Findex:  uint16(3),
		Flength: uint16(1),
	},
	12: {
		Findex:  uint16(3),
		Flength: uint16(1),
	},
	14: {
		Findex:  uint16(4),
		Flength: uint16(1),
	},
	15: {
		Findex:  uint16(4),
		Flength: uint16(1),
	},
	16: {
		Findex:  uint16(5),
		Flength: uint16(1),
	},
}

var ts_field_map_entries = [6]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_children),
		Fchild_index: uint8(1),
	},
	1: {
		Ffield_id:    uint16(field_children),
		Fchild_index: uint8(2),
	},
	2: {
		Ffield_id:    uint16(field_children),
		Fchild_index: uint8(3),
	},
	3: {
		Ffield_id:    uint16(field_children),
		Fchild_index: uint8(4),
	},
	4: {
		Ffield_id:    uint16(field_children),
		Fchild_index: uint8(5),
	},
	5: {
		Ffield_id:    uint16(field_children),
		Fchild_index: uint8(6),
	},
}

var ts_alias_sequences = [17][10]TSSymbol{
	0: {},
	1: {
		1: uint16(alias_sym_string_fragment),
	},
	4: {
		1: uint16(anon_sym_SLASH_DASH),
	},
	5: {
		0: uint16(alias_sym_node_children_comment),
	},
	6: {
		0: uint16(alias_sym_node_field_comment),
		1: uint16(alias_sym_node_field_comment),
	},
	8: {
		1: uint16(anon_sym_SLASH_DASH),
	},
	9: {
		0: uint16(alias_sym_node_field_comment),
		1: uint16(alias_sym_node_field_comment),
		2: uint16(alias_sym_node_field_comment),
	},
	10: {
		2: uint16(alias_sym_decimal),
	},
	12: {
		1: uint16(anon_sym_SLASH_DASH),
	},
	13: {
		3: uint16(alias_sym_decimal),
	},
	15: {
		1: uint16(anon_sym_SLASH_DASH),
	},
	16: {
		1: uint16(anon_sym_SLASH_DASH),
	},
}

var ts_non_terminal_alias_map = [18]uint16_t{
	0:  uint16(sym__node_field),
	1:  uint16(2),
	2:  uint16(sym__node_field),
	3:  uint16(alias_sym_node_field_comment),
	4:  uint16(sym__integer),
	5:  uint16(2),
	6:  uint16(sym__integer),
	7:  uint16(alias_sym_decimal),
	8:  uint16(aux_sym_node_repeat1),
	9:  uint16(3),
	10: uint16(aux_sym_node_repeat1),
	11: uint16(alias_sym_node_field_comment),
	12: uint16(anon_sym_SLASH_DASH),
	13: uint16(aux_sym__escaped_string_repeat1),
	14: uint16(2),
	15: uint16(aux_sym__escaped_string_repeat1),
	16: uint16(alias_sym_string_fragment),
}

var ts_primary_state_ids = [304]TSStateId{
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
	59:  uint16(55),
	60:  uint16(60),
	61:  uint16(60),
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
	74:  uint16(15),
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
	100: uint16(27),
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
	121: uint16(39),
	122: uint16(53),
	123: uint16(38),
	124: uint16(124),
	125: uint16(34),
	126: uint16(49),
	127: uint16(127),
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
	179: uint16(57),
	180: uint16(180),
	181: uint16(181),
	182: uint16(182),
	183: uint16(56),
	184: uint16(184),
	185: uint16(185),
	186: uint16(58),
	187: uint16(15),
	188: uint16(188),
	189: uint16(189),
	190: uint16(190),
	191: uint16(191),
	192: uint16(192),
	193: uint16(27),
	194: uint16(194),
	195: uint16(195),
	196: uint16(196),
	197: uint16(197),
	198: uint16(198),
	199: uint16(199),
	200: uint16(200),
	201: uint16(38),
	202: uint16(202),
	203: uint16(203),
	204: uint16(204),
	205: uint16(49),
	206: uint16(39),
	207: uint16(53),
	208: uint16(208),
	209: uint16(209),
	210: uint16(210),
	211: uint16(211),
	212: uint16(34),
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
	224: uint16(223),
	225: uint16(223),
	226: uint16(223),
	227: uint16(227),
	228: uint16(227),
	229: uint16(227),
	230: uint16(49),
	231: uint16(227),
	232: uint16(232),
	233: uint16(233),
	234: uint16(58),
	235: uint16(235),
	236: uint16(57),
	237: uint16(55),
	238: uint16(238),
	239: uint16(239),
	240: uint16(240),
	241: uint16(241),
	242: uint16(242),
	243: uint16(243),
	244: uint16(60),
	245: uint16(245),
	246: uint16(56),
	247: uint16(247),
	248: uint16(248),
	249: uint16(249),
	250: uint16(250),
	251: uint16(250),
	252: uint16(249),
	253: uint16(249),
	254: uint16(250),
	255: uint16(250),
	256: uint16(249),
	257: uint16(257),
	258: uint16(258),
	259: uint16(259),
	260: uint16(260),
	261: uint16(261),
	262: uint16(15),
	263: uint16(263),
	264: uint16(264),
	265: uint16(27),
	266: uint16(38),
	267: uint16(39),
	268: uint16(53),
	269: uint16(34),
	270: uint16(270),
	271: uint16(271),
	272: uint16(55),
	273: uint16(273),
	274: uint16(60),
	275: uint16(58),
	276: uint16(56),
	277: uint16(277),
	278: uint16(189),
	279: uint16(188),
	280: uint16(273),
	281: uint16(57),
	282: uint16(277),
	283: uint16(283),
	284: uint16(190),
	285: uint16(285),
	286: uint16(286),
	287: uint16(287),
	288: uint16(288),
	289: uint16(289),
	290: uint16(218),
	291: uint16(291),
	292: uint16(192),
	293: uint16(197),
	294: uint16(196),
	295: uint16(295),
	296: uint16(194),
	297: uint16(297),
	298: uint16(298),
	299: uint16(299),
	300: uint16(300),
	301: uint16(301),
	302: uint16(302),
	303: uint16(303),
}

var sym__normal_bare_identifier_character_set_2 = [642]TSCharacterRange{
	0: {
		Fstart: int32('!'),
		Fend:   int32('!'),
	},
	1: {
		Fstart: int32('#'),
		Fend:   int32('\''),
	},
	2: {
		Fstart: int32('*'),
		Fend:   int32('+'),
	},
	3: {
		Fstart: int32('-'),
		Fend:   int32('.'),
	},
	4: {
		Fstart: int32('0'),
		Fend:   int32(':'),
	},
	5: {
		Fstart: int32('?'),
		Fend:   int32('Z'),
	},
	6: {
		Fstart: int32('^'),
		Fend:   int32('_'),
	},
	7: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	8: {
		Fstart: int32('|'),
		Fend:   int32('|'),
	},
	9: {
		Fstart: int32('~'),
		Fend:   int32('~'),
	},
	10: {
		Fstart: int32(0xa9),
		Fend:   int32(0xaa),
	},
	11: {
		Fstart: int32(0xae),
		Fend:   int32(0xae),
	},
	12: {
		Fstart: int32(0xb2),
		Fend:   int32(0xb3),
	},
	13: {
		Fstart: int32(0xb5),
		Fend:   int32(0xb5),
	},
	14: {
		Fstart: int32(0xb9),
		Fend:   int32(0xba),
	},
	15: {
		Fstart: int32(0xbc),
		Fend:   int32(0xbe),
	},
	16: {
		Fstart: int32(0xc0),
		Fend:   int32(0xd6),
	},
	17: {
		Fstart: int32(0xd8),
		Fend:   int32(0xf6),
	},
	18: {
		Fstart: int32(0xf8),
		Fend:   int32(0x2c1),
	},
	19: {
		Fstart: int32(0x2c6),
		Fend:   int32(0x2d1),
	},
	20: {
		Fstart: int32(0x2e0),
		Fend:   int32(0x2e4),
	},
	21: {
		Fstart: int32(0x2ec),
		Fend:   int32(0x2ec),
	},
	22: {
		Fstart: int32(0x2ee),
		Fend:   int32(0x2ee),
	},
	23: {
		Fstart: int32(0x300),
		Fend:   int32(0x374),
	},
	24: {
		Fstart: int32(0x376),
		Fend:   int32(0x377),
	},
	25: {
		Fstart: int32(0x37a),
		Fend:   int32(0x37d),
	},
	26: {
		Fstart: int32(0x37f),
		Fend:   int32(0x37f),
	},
	27: {
		Fstart: int32(0x386),
		Fend:   int32(0x386),
	},
	28: {
		Fstart: int32(0x388),
		Fend:   int32(0x38a),
	},
	29: {
		Fstart: int32(0x38c),
		Fend:   int32(0x38c),
	},
	30: {
		Fstart: int32(0x38e),
		Fend:   int32(0x3a1),
	},
	31: {
		Fstart: int32(0x3a3),
		Fend:   int32(0x3f5),
	},
	32: {
		Fstart: int32(0x3f7),
		Fend:   int32(0x481),
	},
	33: {
		Fstart: int32(0x483),
		Fend:   int32(0x52f),
	},
	34: {
		Fstart: int32(0x531),
		Fend:   int32(0x556),
	},
	35: {
		Fstart: int32(0x559),
		Fend:   int32(0x559),
	},
	36: {
		Fstart: int32(0x560),
		Fend:   int32(0x588),
	},
	37: {
		Fstart: int32(0x591),
		Fend:   int32(0x5bd),
	},
	38: {
		Fstart: int32(0x5bf),
		Fend:   int32(0x5bf),
	},
	39: {
		Fstart: int32(0x5c1),
		Fend:   int32(0x5c2),
	},
	40: {
		Fstart: int32(0x5c4),
		Fend:   int32(0x5c5),
	},
	41: {
		Fstart: int32(0x5c7),
		Fend:   int32(0x5c7),
	},
	42: {
		Fstart: int32(0x5d0),
		Fend:   int32(0x5ea),
	},
	43: {
		Fstart: int32(0x5ef),
		Fend:   int32(0x5f2),
	},
	44: {
		Fstart: int32(0x610),
		Fend:   int32(0x61a),
	},
	45: {
		Fstart: int32(0x620),
		Fend:   int32(0x669),
	},
	46: {
		Fstart: int32(0x66e),
		Fend:   int32(0x6d3),
	},
	47: {
		Fstart: int32(0x6d5),
		Fend:   int32(0x6dc),
	},
	48: {
		Fstart: int32(0x6df),
		Fend:   int32(0x6e8),
	},
	49: {
		Fstart: int32(0x6ea),
		Fend:   int32(0x6fc),
	},
	50: {
		Fstart: int32(0x6ff),
		Fend:   int32(0x6ff),
	},
	51: {
		Fstart: int32(0x710),
		Fend:   int32(0x74a),
	},
	52: {
		Fstart: int32(0x74d),
		Fend:   int32(0x7b1),
	},
	53: {
		Fstart: int32(0x7c0),
		Fend:   int32(0x7f5),
	},
	54: {
		Fstart: int32(0x7fa),
		Fend:   int32(0x7fa),
	},
	55: {
		Fstart: int32(0x7fd),
		Fend:   int32(0x7fd),
	},
	56: {
		Fstart: int32(0x800),
		Fend:   int32(0x82d),
	},
	57: {
		Fstart: int32(0x840),
		Fend:   int32(0x85b),
	},
	58: {
		Fstart: int32(0x860),
		Fend:   int32(0x86a),
	},
	59: {
		Fstart: int32(0x870),
		Fend:   int32(0x887),
	},
	60: {
		Fstart: int32(0x889),
		Fend:   int32(0x88e),
	},
	61: {
		Fstart: int32(0x898),
		Fend:   int32(0x8e1),
	},
	62: {
		Fstart: int32(0x8e3),
		Fend:   int32(0x963),
	},
	63: {
		Fstart: int32(0x966),
		Fend:   int32(0x96f),
	},
	64: {
		Fstart: int32(0x971),
		Fend:   int32(0x983),
	},
	65: {
		Fstart: int32(0x985),
		Fend:   int32(0x98c),
	},
	66: {
		Fstart: int32(0x98f),
		Fend:   int32(0x990),
	},
	67: {
		Fstart: int32(0x993),
		Fend:   int32(0x9a8),
	},
	68: {
		Fstart: int32(0x9aa),
		Fend:   int32(0x9b0),
	},
	69: {
		Fstart: int32(0x9b2),
		Fend:   int32(0x9b2),
	},
	70: {
		Fstart: int32(0x9b6),
		Fend:   int32(0x9b9),
	},
	71: {
		Fstart: int32(0x9bc),
		Fend:   int32(0x9c4),
	},
	72: {
		Fstart: int32(0x9c7),
		Fend:   int32(0x9c8),
	},
	73: {
		Fstart: int32(0x9cb),
		Fend:   int32(0x9ce),
	},
	74: {
		Fstart: int32(0x9d7),
		Fend:   int32(0x9d7),
	},
	75: {
		Fstart: int32(0x9dc),
		Fend:   int32(0x9dd),
	},
	76: {
		Fstart: int32(0x9df),
		Fend:   int32(0x9e3),
	},
	77: {
		Fstart: int32(0x9e6),
		Fend:   int32(0x9f1),
	},
	78: {
		Fstart: int32(0x9f4),
		Fend:   int32(0x9f9),
	},
	79: {
		Fstart: int32(0x9fc),
		Fend:   int32(0x9fc),
	},
	80: {
		Fstart: int32(0x9fe),
		Fend:   int32(0x9fe),
	},
	81: {
		Fstart: int32(0xa01),
		Fend:   int32(0xa03),
	},
	82: {
		Fstart: int32(0xa05),
		Fend:   int32(0xa0a),
	},
	83: {
		Fstart: int32(0xa0f),
		Fend:   int32(0xa10),
	},
	84: {
		Fstart: int32(0xa13),
		Fend:   int32(0xa28),
	},
	85: {
		Fstart: int32(0xa2a),
		Fend:   int32(0xa30),
	},
	86: {
		Fstart: int32(0xa32),
		Fend:   int32(0xa33),
	},
	87: {
		Fstart: int32(0xa35),
		Fend:   int32(0xa36),
	},
	88: {
		Fstart: int32(0xa38),
		Fend:   int32(0xa39),
	},
	89: {
		Fstart: int32(0xa3c),
		Fend:   int32(0xa3c),
	},
	90: {
		Fstart: int32(0xa3e),
		Fend:   int32(0xa42),
	},
	91: {
		Fstart: int32(0xa47),
		Fend:   int32(0xa48),
	},
	92: {
		Fstart: int32(0xa4b),
		Fend:   int32(0xa4d),
	},
	93: {
		Fstart: int32(0xa51),
		Fend:   int32(0xa51),
	},
	94: {
		Fstart: int32(0xa59),
		Fend:   int32(0xa5c),
	},
	95: {
		Fstart: int32(0xa5e),
		Fend:   int32(0xa5e),
	},
	96: {
		Fstart: int32(0xa66),
		Fend:   int32(0xa75),
	},
	97: {
		Fstart: int32(0xa81),
		Fend:   int32(0xa83),
	},
	98: {
		Fstart: int32(0xa85),
		Fend:   int32(0xa8d),
	},
	99: {
		Fstart: int32(0xa8f),
		Fend:   int32(0xa91),
	},
	100: {
		Fstart: int32(0xa93),
		Fend:   int32(0xaa8),
	},
	101: {
		Fstart: int32(0xaaa),
		Fend:   int32(0xab0),
	},
	102: {
		Fstart: int32(0xab2),
		Fend:   int32(0xab3),
	},
	103: {
		Fstart: int32(0xab5),
		Fend:   int32(0xab9),
	},
	104: {
		Fstart: int32(0xabc),
		Fend:   int32(0xac5),
	},
	105: {
		Fstart: int32(0xac7),
		Fend:   int32(0xac9),
	},
	106: {
		Fstart: int32(0xacb),
		Fend:   int32(0xacd),
	},
	107: {
		Fstart: int32(0xad0),
		Fend:   int32(0xad0),
	},
	108: {
		Fstart: int32(0xae0),
		Fend:   int32(0xae3),
	},
	109: {
		Fstart: int32(0xae6),
		Fend:   int32(0xaef),
	},
	110: {
		Fstart: int32(0xaf9),
		Fend:   int32(0xaff),
	},
	111: {
		Fstart: int32(0xb01),
		Fend:   int32(0xb03),
	},
	112: {
		Fstart: int32(0xb05),
		Fend:   int32(0xb0c),
	},
	113: {
		Fstart: int32(0xb0f),
		Fend:   int32(0xb10),
	},
	114: {
		Fstart: int32(0xb13),
		Fend:   int32(0xb28),
	},
	115: {
		Fstart: int32(0xb2a),
		Fend:   int32(0xb30),
	},
	116: {
		Fstart: int32(0xb32),
		Fend:   int32(0xb33),
	},
	117: {
		Fstart: int32(0xb35),
		Fend:   int32(0xb39),
	},
	118: {
		Fstart: int32(0xb3c),
		Fend:   int32(0xb44),
	},
	119: {
		Fstart: int32(0xb47),
		Fend:   int32(0xb48),
	},
	120: {
		Fstart: int32(0xb4b),
		Fend:   int32(0xb4d),
	},
	121: {
		Fstart: int32(0xb55),
		Fend:   int32(0xb57),
	},
	122: {
		Fstart: int32(0xb5c),
		Fend:   int32(0xb5d),
	},
	123: {
		Fstart: int32(0xb5f),
		Fend:   int32(0xb63),
	},
	124: {
		Fstart: int32(0xb66),
		Fend:   int32(0xb6f),
	},
	125: {
		Fstart: int32(0xb71),
		Fend:   int32(0xb77),
	},
	126: {
		Fstart: int32(0xb82),
		Fend:   int32(0xb83),
	},
	127: {
		Fstart: int32(0xb85),
		Fend:   int32(0xb8a),
	},
	128: {
		Fstart: int32(0xb8e),
		Fend:   int32(0xb90),
	},
	129: {
		Fstart: int32(0xb92),
		Fend:   int32(0xb95),
	},
	130: {
		Fstart: int32(0xb99),
		Fend:   int32(0xb9a),
	},
	131: {
		Fstart: int32(0xb9c),
		Fend:   int32(0xb9c),
	},
	132: {
		Fstart: int32(0xb9e),
		Fend:   int32(0xb9f),
	},
	133: {
		Fstart: int32(0xba3),
		Fend:   int32(0xba4),
	},
	134: {
		Fstart: int32(0xba8),
		Fend:   int32(0xbaa),
	},
	135: {
		Fstart: int32(0xbae),
		Fend:   int32(0xbb9),
	},
	136: {
		Fstart: int32(0xbbe),
		Fend:   int32(0xbc2),
	},
	137: {
		Fstart: int32(0xbc6),
		Fend:   int32(0xbc8),
	},
	138: {
		Fstart: int32(0xbca),
		Fend:   int32(0xbcd),
	},
	139: {
		Fstart: int32(0xbd0),
		Fend:   int32(0xbd0),
	},
	140: {
		Fstart: int32(0xbd7),
		Fend:   int32(0xbd7),
	},
	141: {
		Fstart: int32(0xbe6),
		Fend:   int32(0xbf2),
	},
	142: {
		Fstart: int32(0xc00),
		Fend:   int32(0xc0c),
	},
	143: {
		Fstart: int32(0xc0e),
		Fend:   int32(0xc10),
	},
	144: {
		Fstart: int32(0xc12),
		Fend:   int32(0xc28),
	},
	145: {
		Fstart: int32(0xc2a),
		Fend:   int32(0xc39),
	},
	146: {
		Fstart: int32(0xc3c),
		Fend:   int32(0xc44),
	},
	147: {
		Fstart: int32(0xc46),
		Fend:   int32(0xc48),
	},
	148: {
		Fstart: int32(0xc4a),
		Fend:   int32(0xc4d),
	},
	149: {
		Fstart: int32(0xc55),
		Fend:   int32(0xc56),
	},
	150: {
		Fstart: int32(0xc58),
		Fend:   int32(0xc5a),
	},
	151: {
		Fstart: int32(0xc5d),
		Fend:   int32(0xc5d),
	},
	152: {
		Fstart: int32(0xc60),
		Fend:   int32(0xc63),
	},
	153: {
		Fstart: int32(0xc66),
		Fend:   int32(0xc6f),
	},
	154: {
		Fstart: int32(0xc78),
		Fend:   int32(0xc7e),
	},
	155: {
		Fstart: int32(0xc80),
		Fend:   int32(0xc83),
	},
	156: {
		Fstart: int32(0xc85),
		Fend:   int32(0xc8c),
	},
	157: {
		Fstart: int32(0xc8e),
		Fend:   int32(0xc90),
	},
	158: {
		Fstart: int32(0xc92),
		Fend:   int32(0xca8),
	},
	159: {
		Fstart: int32(0xcaa),
		Fend:   int32(0xcb3),
	},
	160: {
		Fstart: int32(0xcb5),
		Fend:   int32(0xcb9),
	},
	161: {
		Fstart: int32(0xcbc),
		Fend:   int32(0xcc4),
	},
	162: {
		Fstart: int32(0xcc6),
		Fend:   int32(0xcc8),
	},
	163: {
		Fstart: int32(0xcca),
		Fend:   int32(0xccd),
	},
	164: {
		Fstart: int32(0xcd5),
		Fend:   int32(0xcd6),
	},
	165: {
		Fstart: int32(0xcdd),
		Fend:   int32(0xcde),
	},
	166: {
		Fstart: int32(0xce0),
		Fend:   int32(0xce3),
	},
	167: {
		Fstart: int32(0xce6),
		Fend:   int32(0xcef),
	},
	168: {
		Fstart: int32(0xcf1),
		Fend:   int32(0xcf3),
	},
	169: {
		Fstart: int32(0xd00),
		Fend:   int32(0xd0c),
	},
	170: {
		Fstart: int32(0xd0e),
		Fend:   int32(0xd10),
	},
	171: {
		Fstart: int32(0xd12),
		Fend:   int32(0xd44),
	},
	172: {
		Fstart: int32(0xd46),
		Fend:   int32(0xd48),
	},
	173: {
		Fstart: int32(0xd4a),
		Fend:   int32(0xd4e),
	},
	174: {
		Fstart: int32(0xd54),
		Fend:   int32(0xd63),
	},
	175: {
		Fstart: int32(0xd66),
		Fend:   int32(0xd78),
	},
	176: {
		Fstart: int32(0xd7a),
		Fend:   int32(0xd7f),
	},
	177: {
		Fstart: int32(0xd81),
		Fend:   int32(0xd83),
	},
	178: {
		Fstart: int32(0xd85),
		Fend:   int32(0xd96),
	},
	179: {
		Fstart: int32(0xd9a),
		Fend:   int32(0xdb1),
	},
	180: {
		Fstart: int32(0xdb3),
		Fend:   int32(0xdbb),
	},
	181: {
		Fstart: int32(0xdbd),
		Fend:   int32(0xdbd),
	},
	182: {
		Fstart: int32(0xdc0),
		Fend:   int32(0xdc6),
	},
	183: {
		Fstart: int32(0xdca),
		Fend:   int32(0xdca),
	},
	184: {
		Fstart: int32(0xdcf),
		Fend:   int32(0xdd4),
	},
	185: {
		Fstart: int32(0xdd6),
		Fend:   int32(0xdd6),
	},
	186: {
		Fstart: int32(0xdd8),
		Fend:   int32(0xddf),
	},
	187: {
		Fstart: int32(0xde6),
		Fend:   int32(0xdef),
	},
	188: {
		Fstart: int32(0xdf2),
		Fend:   int32(0xdf3),
	},
	189: {
		Fstart: int32(0xe01),
		Fend:   int32(0xe3a),
	},
	190: {
		Fstart: int32(0xe40),
		Fend:   int32(0xe4e),
	},
	191: {
		Fstart: int32(0xe50),
		Fend:   int32(0xe59),
	},
	192: {
		Fstart: int32(0xe81),
		Fend:   int32(0xe82),
	},
	193: {
		Fstart: int32(0xe84),
		Fend:   int32(0xe84),
	},
	194: {
		Fstart: int32(0xe86),
		Fend:   int32(0xe8a),
	},
	195: {
		Fstart: int32(0xe8c),
		Fend:   int32(0xea3),
	},
	196: {
		Fstart: int32(0xea5),
		Fend:   int32(0xea5),
	},
	197: {
		Fstart: int32(0xea7),
		Fend:   int32(0xebd),
	},
	198: {
		Fstart: int32(0xec0),
		Fend:   int32(0xec4),
	},
	199: {
		Fstart: int32(0xec6),
		Fend:   int32(0xec6),
	},
	200: {
		Fstart: int32(0xec8),
		Fend:   int32(0xece),
	},
	201: {
		Fstart: int32(0xed0),
		Fend:   int32(0xed9),
	},
	202: {
		Fstart: int32(0xedc),
		Fend:   int32(0xedf),
	},
	203: {
		Fstart: int32(0xf00),
		Fend:   int32(0xf00),
	},
	204: {
		Fstart: int32(0xf18),
		Fend:   int32(0xf19),
	},
	205: {
		Fstart: int32(0xf20),
		Fend:   int32(0xf33),
	},
	206: {
		Fstart: int32(0xf35),
		Fend:   int32(0xf35),
	},
	207: {
		Fstart: int32(0xf37),
		Fend:   int32(0xf37),
	},
	208: {
		Fstart: int32(0xf39),
		Fend:   int32(0xf39),
	},
	209: {
		Fstart: int32(0xf3e),
		Fend:   int32(0xf47),
	},
	210: {
		Fstart: int32(0xf49),
		Fend:   int32(0xf6c),
	},
	211: {
		Fstart: int32(0xf71),
		Fend:   int32(0xf84),
	},
	212: {
		Fstart: int32(0xf86),
		Fend:   int32(0xf97),
	},
	213: {
		Fstart: int32(0xf99),
		Fend:   int32(0xfbc),
	},
	214: {
		Fstart: int32(0xfc6),
		Fend:   int32(0xfc6),
	},
	215: {
		Fstart: int32(0x1000),
		Fend:   int32(0x1049),
	},
	216: {
		Fstart: int32(0x1050),
		Fend:   int32(0x109d),
	},
	217: {
		Fstart: int32(0x10a0),
		Fend:   int32(0x10c5),
	},
	218: {
		Fstart: int32(0x10c7),
		Fend:   int32(0x10c7),
	},
	219: {
		Fstart: int32(0x10cd),
		Fend:   int32(0x10cd),
	},
	220: {
		Fstart: int32(0x10d0),
		Fend:   int32(0x10fa),
	},
	221: {
		Fstart: int32(0x10fc),
		Fend:   int32(0x1248),
	},
	222: {
		Fstart: int32(0x124a),
		Fend:   int32(0x124d),
	},
	223: {
		Fstart: int32(0x1250),
		Fend:   int32(0x1256),
	},
	224: {
		Fstart: int32(0x1258),
		Fend:   int32(0x1258),
	},
	225: {
		Fstart: int32(0x125a),
		Fend:   int32(0x125d),
	},
	226: {
		Fstart: int32(0x1260),
		Fend:   int32(0x1288),
	},
	227: {
		Fstart: int32(0x128a),
		Fend:   int32(0x128d),
	},
	228: {
		Fstart: int32(0x1290),
		Fend:   int32(0x12b0),
	},
	229: {
		Fstart: int32(0x12b2),
		Fend:   int32(0x12b5),
	},
	230: {
		Fstart: int32(0x12b8),
		Fend:   int32(0x12be),
	},
	231: {
		Fstart: int32(0x12c0),
		Fend:   int32(0x12c0),
	},
	232: {
		Fstart: int32(0x12c2),
		Fend:   int32(0x12c5),
	},
	233: {
		Fstart: int32(0x12c8),
		Fend:   int32(0x12d6),
	},
	234: {
		Fstart: int32(0x12d8),
		Fend:   int32(0x1310),
	},
	235: {
		Fstart: int32(0x1312),
		Fend:   int32(0x1315),
	},
	236: {
		Fstart: int32(0x1318),
		Fend:   int32(0x135a),
	},
	237: {
		Fstart: int32(0x135d),
		Fend:   int32(0x135f),
	},
	238: {
		Fstart: int32(0x1369),
		Fend:   int32(0x137c),
	},
	239: {
		Fstart: int32(0x1380),
		Fend:   int32(0x138f),
	},
	240: {
		Fstart: int32(0x13a0),
		Fend:   int32(0x13f5),
	},
	241: {
		Fstart: int32(0x13f8),
		Fend:   int32(0x13fd),
	},
	242: {
		Fstart: int32(0x1401),
		Fend:   int32(0x166c),
	},
	243: {
		Fstart: int32(0x166f),
		Fend:   int32(0x167f),
	},
	244: {
		Fstart: int32(0x1681),
		Fend:   int32(0x169a),
	},
	245: {
		Fstart: int32(0x16a0),
		Fend:   int32(0x16ea),
	},
	246: {
		Fstart: int32(0x16ee),
		Fend:   int32(0x16f8),
	},
	247: {
		Fstart: int32(0x1700),
		Fend:   int32(0x1715),
	},
	248: {
		Fstart: int32(0x171f),
		Fend:   int32(0x1734),
	},
	249: {
		Fstart: int32(0x1740),
		Fend:   int32(0x1753),
	},
	250: {
		Fstart: int32(0x1760),
		Fend:   int32(0x176c),
	},
	251: {
		Fstart: int32(0x176e),
		Fend:   int32(0x1770),
	},
	252: {
		Fstart: int32(0x1772),
		Fend:   int32(0x1773),
	},
	253: {
		Fstart: int32(0x1780),
		Fend:   int32(0x17d3),
	},
	254: {
		Fstart: int32(0x17d7),
		Fend:   int32(0x17d7),
	},
	255: {
		Fstart: int32(0x17dc),
		Fend:   int32(0x17dd),
	},
	256: {
		Fstart: int32(0x17e0),
		Fend:   int32(0x17e9),
	},
	257: {
		Fstart: int32(0x17f0),
		Fend:   int32(0x17f9),
	},
	258: {
		Fstart: int32(0x180b),
		Fend:   int32(0x180d),
	},
	259: {
		Fstart: int32(0x180f),
		Fend:   int32(0x1819),
	},
	260: {
		Fstart: int32(0x1820),
		Fend:   int32(0x1878),
	},
	261: {
		Fstart: int32(0x1880),
		Fend:   int32(0x18aa),
	},
	262: {
		Fstart: int32(0x18b0),
		Fend:   int32(0x18f5),
	},
	263: {
		Fstart: int32(0x1900),
		Fend:   int32(0x191e),
	},
	264: {
		Fstart: int32(0x1920),
		Fend:   int32(0x192b),
	},
	265: {
		Fstart: int32(0x1930),
		Fend:   int32(0x193b),
	},
	266: {
		Fstart: int32(0x1946),
		Fend:   int32(0x196d),
	},
	267: {
		Fstart: int32(0x1970),
		Fend:   int32(0x1974),
	},
	268: {
		Fstart: int32(0x1980),
		Fend:   int32(0x19ab),
	},
	269: {
		Fstart: int32(0x19b0),
		Fend:   int32(0x19c9),
	},
	270: {
		Fstart: int32(0x19d0),
		Fend:   int32(0x19da),
	},
	271: {
		Fstart: int32(0x1a00),
		Fend:   int32(0x1a1b),
	},
	272: {
		Fstart: int32(0x1a20),
		Fend:   int32(0x1a5e),
	},
	273: {
		Fstart: int32(0x1a60),
		Fend:   int32(0x1a7c),
	},
	274: {
		Fstart: int32(0x1a7f),
		Fend:   int32(0x1a89),
	},
	275: {
		Fstart: int32(0x1a90),
		Fend:   int32(0x1a99),
	},
	276: {
		Fstart: int32(0x1aa7),
		Fend:   int32(0x1aa7),
	},
	277: {
		Fstart: int32(0x1ab0),
		Fend:   int32(0x1ace),
	},
	278: {
		Fstart: int32(0x1b00),
		Fend:   int32(0x1b4c),
	},
	279: {
		Fstart: int32(0x1b50),
		Fend:   int32(0x1b59),
	},
	280: {
		Fstart: int32(0x1b6b),
		Fend:   int32(0x1b73),
	},
	281: {
		Fstart: int32(0x1b80),
		Fend:   int32(0x1bf3),
	},
	282: {
		Fstart: int32(0x1c00),
		Fend:   int32(0x1c37),
	},
	283: {
		Fstart: int32(0x1c40),
		Fend:   int32(0x1c49),
	},
	284: {
		Fstart: int32(0x1c4d),
		Fend:   int32(0x1c7d),
	},
	285: {
		Fstart: int32(0x1c80),
		Fend:   int32(0x1c88),
	},
	286: {
		Fstart: int32(0x1c90),
		Fend:   int32(0x1cba),
	},
	287: {
		Fstart: int32(0x1cbd),
		Fend:   int32(0x1cbf),
	},
	288: {
		Fstart: int32(0x1cd0),
		Fend:   int32(0x1cd2),
	},
	289: {
		Fstart: int32(0x1cd4),
		Fend:   int32(0x1cfa),
	},
	290: {
		Fstart: int32(0x1d00),
		Fend:   int32(0x1f15),
	},
	291: {
		Fstart: int32(0x1f18),
		Fend:   int32(0x1f1d),
	},
	292: {
		Fstart: int32(0x1f20),
		Fend:   int32(0x1f45),
	},
	293: {
		Fstart: int32(0x1f48),
		Fend:   int32(0x1f4d),
	},
	294: {
		Fstart: int32(0x1f50),
		Fend:   int32(0x1f57),
	},
	295: {
		Fstart: int32(0x1f59),
		Fend:   int32(0x1f59),
	},
	296: {
		Fstart: int32(0x1f5b),
		Fend:   int32(0x1f5b),
	},
	297: {
		Fstart: int32(0x1f5d),
		Fend:   int32(0x1f5d),
	},
	298: {
		Fstart: int32(0x1f5f),
		Fend:   int32(0x1f7d),
	},
	299: {
		Fstart: int32(0x1f80),
		Fend:   int32(0x1fb4),
	},
	300: {
		Fstart: int32(0x1fb6),
		Fend:   int32(0x1fbc),
	},
	301: {
		Fstart: int32(0x1fbe),
		Fend:   int32(0x1fbe),
	},
	302: {
		Fstart: int32(0x1fc2),
		Fend:   int32(0x1fc4),
	},
	303: {
		Fstart: int32(0x1fc6),
		Fend:   int32(0x1fcc),
	},
	304: {
		Fstart: int32(0x1fd0),
		Fend:   int32(0x1fd3),
	},
	305: {
		Fstart: int32(0x1fd6),
		Fend:   int32(0x1fdb),
	},
	306: {
		Fstart: int32(0x1fe0),
		Fend:   int32(0x1fec),
	},
	307: {
		Fstart: int32(0x1ff2),
		Fend:   int32(0x1ff4),
	},
	308: {
		Fstart: int32(0x1ff6),
		Fend:   int32(0x1ffc),
	},
	309: {
		Fstart: int32(0x203c),
		Fend:   int32(0x203c),
	},
	310: {
		Fstart: int32(0x2049),
		Fend:   int32(0x2049),
	},
	311: {
		Fstart: int32(0x2070),
		Fend:   int32(0x2071),
	},
	312: {
		Fstart: int32(0x2074),
		Fend:   int32(0x2079),
	},
	313: {
		Fstart: int32(0x207f),
		Fend:   int32(0x2089),
	},
	314: {
		Fstart: int32(0x2090),
		Fend:   int32(0x209c),
	},
	315: {
		Fstart: int32(0x20d0),
		Fend:   int32(0x20f0),
	},
	316: {
		Fstart: int32(0x2102),
		Fend:   int32(0x2102),
	},
	317: {
		Fstart: int32(0x2107),
		Fend:   int32(0x2107),
	},
	318: {
		Fstart: int32(0x210a),
		Fend:   int32(0x2113),
	},
	319: {
		Fstart: int32(0x2115),
		Fend:   int32(0x2115),
	},
	320: {
		Fstart: int32(0x2119),
		Fend:   int32(0x211d),
	},
	321: {
		Fstart: int32(0x2122),
		Fend:   int32(0x2122),
	},
	322: {
		Fstart: int32(0x2124),
		Fend:   int32(0x2124),
	},
	323: {
		Fstart: int32(0x2126),
		Fend:   int32(0x2126),
	},
	324: {
		Fstart: int32(0x2128),
		Fend:   int32(0x2128),
	},
	325: {
		Fstart: int32(0x212a),
		Fend:   int32(0x212d),
	},
	326: {
		Fstart: int32(0x212f),
		Fend:   int32(0x2139),
	},
	327: {
		Fstart: int32(0x213c),
		Fend:   int32(0x213f),
	},
	328: {
		Fstart: int32(0x2145),
		Fend:   int32(0x2149),
	},
	329: {
		Fstart: int32(0x214e),
		Fend:   int32(0x214e),
	},
	330: {
		Fstart: int32(0x2150),
		Fend:   int32(0x2189),
	},
	331: {
		Fstart: int32(0x2194),
		Fend:   int32(0x2199),
	},
	332: {
		Fstart: int32(0x21a9),
		Fend:   int32(0x21aa),
	},
	333: {
		Fstart: int32(0x231a),
		Fend:   int32(0x231b),
	},
	334: {
		Fstart: int32(0x2328),
		Fend:   int32(0x2328),
	},
	335: {
		Fstart: int32(0x23cf),
		Fend:   int32(0x23cf),
	},
	336: {
		Fstart: int32(0x23e9),
		Fend:   int32(0x23f3),
	},
	337: {
		Fstart: int32(0x23f8),
		Fend:   int32(0x23fa),
	},
	338: {
		Fstart: int32(0x2460),
		Fend:   int32(0x249b),
	},
	339: {
		Fstart: int32(0x24c2),
		Fend:   int32(0x24c2),
	},
	340: {
		Fstart: int32(0x24ea),
		Fend:   int32(0x24ff),
	},
	341: {
		Fstart: int32(0x25aa),
		Fend:   int32(0x25ab),
	},
	342: {
		Fstart: int32(0x25b6),
		Fend:   int32(0x25b6),
	},
	343: {
		Fstart: int32(0x25c0),
		Fend:   int32(0x25c0),
	},
	344: {
		Fstart: int32(0x25fb),
		Fend:   int32(0x25fe),
	},
	345: {
		Fstart: int32(0x2600),
		Fend:   int32(0x2604),
	},
	346: {
		Fstart: int32(0x260e),
		Fend:   int32(0x260e),
	},
	347: {
		Fstart: int32(0x2611),
		Fend:   int32(0x2611),
	},
	348: {
		Fstart: int32(0x2614),
		Fend:   int32(0x2615),
	},
	349: {
		Fstart: int32(0x2618),
		Fend:   int32(0x2618),
	},
	350: {
		Fstart: int32(0x261d),
		Fend:   int32(0x261d),
	},
	351: {
		Fstart: int32(0x2620),
		Fend:   int32(0x2620),
	},
	352: {
		Fstart: int32(0x2622),
		Fend:   int32(0x2623),
	},
	353: {
		Fstart: int32(0x2626),
		Fend:   int32(0x2626),
	},
	354: {
		Fstart: int32(0x262a),
		Fend:   int32(0x262a),
	},
	355: {
		Fstart: int32(0x262e),
		Fend:   int32(0x262f),
	},
	356: {
		Fstart: int32(0x2638),
		Fend:   int32(0x263a),
	},
	357: {
		Fstart: int32(0x2640),
		Fend:   int32(0x2640),
	},
	358: {
		Fstart: int32(0x2642),
		Fend:   int32(0x2642),
	},
	359: {
		Fstart: int32(0x2648),
		Fend:   int32(0x2653),
	},
	360: {
		Fstart: int32(0x265f),
		Fend:   int32(0x2660),
	},
	361: {
		Fstart: int32(0x2663),
		Fend:   int32(0x2663),
	},
	362: {
		Fstart: int32(0x2665),
		Fend:   int32(0x2666),
	},
	363: {
		Fstart: int32(0x2668),
		Fend:   int32(0x2668),
	},
	364: {
		Fstart: int32(0x267b),
		Fend:   int32(0x267b),
	},
	365: {
		Fstart: int32(0x267e),
		Fend:   int32(0x267f),
	},
	366: {
		Fstart: int32(0x2692),
		Fend:   int32(0x2697),
	},
	367: {
		Fstart: int32(0x2699),
		Fend:   int32(0x2699),
	},
	368: {
		Fstart: int32(0x269b),
		Fend:   int32(0x269c),
	},
	369: {
		Fstart: int32(0x26a0),
		Fend:   int32(0x26a1),
	},
	370: {
		Fstart: int32(0x26a7),
		Fend:   int32(0x26a7),
	},
	371: {
		Fstart: int32(0x26aa),
		Fend:   int32(0x26ab),
	},
	372: {
		Fstart: int32(0x26b0),
		Fend:   int32(0x26b1),
	},
	373: {
		Fstart: int32(0x26bd),
		Fend:   int32(0x26be),
	},
	374: {
		Fstart: int32(0x26c4),
		Fend:   int32(0x26c5),
	},
	375: {
		Fstart: int32(0x26c8),
		Fend:   int32(0x26c8),
	},
	376: {
		Fstart: int32(0x26ce),
		Fend:   int32(0x26cf),
	},
	377: {
		Fstart: int32(0x26d1),
		Fend:   int32(0x26d1),
	},
	378: {
		Fstart: int32(0x26d3),
		Fend:   int32(0x26d4),
	},
	379: {
		Fstart: int32(0x26e9),
		Fend:   int32(0x26ea),
	},
	380: {
		Fstart: int32(0x26f0),
		Fend:   int32(0x26f5),
	},
	381: {
		Fstart: int32(0x26f7),
		Fend:   int32(0x26fa),
	},
	382: {
		Fstart: int32(0x26fd),
		Fend:   int32(0x26fd),
	},
	383: {
		Fstart: int32(0x2702),
		Fend:   int32(0x2702),
	},
	384: {
		Fstart: int32(0x2705),
		Fend:   int32(0x2705),
	},
	385: {
		Fstart: int32(0x2708),
		Fend:   int32(0x270d),
	},
	386: {
		Fstart: int32(0x270f),
		Fend:   int32(0x270f),
	},
	387: {
		Fstart: int32(0x2712),
		Fend:   int32(0x2712),
	},
	388: {
		Fstart: int32(0x2714),
		Fend:   int32(0x2714),
	},
	389: {
		Fstart: int32(0x2716),
		Fend:   int32(0x2716),
	},
	390: {
		Fstart: int32(0x271d),
		Fend:   int32(0x271d),
	},
	391: {
		Fstart: int32(0x2721),
		Fend:   int32(0x2721),
	},
	392: {
		Fstart: int32(0x2728),
		Fend:   int32(0x2728),
	},
	393: {
		Fstart: int32(0x2733),
		Fend:   int32(0x2734),
	},
	394: {
		Fstart: int32(0x2744),
		Fend:   int32(0x2744),
	},
	395: {
		Fstart: int32(0x2747),
		Fend:   int32(0x2747),
	},
	396: {
		Fstart: int32(0x274c),
		Fend:   int32(0x274c),
	},
	397: {
		Fstart: int32(0x274e),
		Fend:   int32(0x274e),
	},
	398: {
		Fstart: int32(0x2753),
		Fend:   int32(0x2755),
	},
	399: {
		Fstart: int32(0x2757),
		Fend:   int32(0x2757),
	},
	400: {
		Fstart: int32(0x2763),
		Fend:   int32(0x2764),
	},
	401: {
		Fstart: int32(0x2776),
		Fend:   int32(0x2793),
	},
	402: {
		Fstart: int32(0x2795),
		Fend:   int32(0x2797),
	},
	403: {
		Fstart: int32(0x27a1),
		Fend:   int32(0x27a1),
	},
	404: {
		Fstart: int32(0x27b0),
		Fend:   int32(0x27b0),
	},
	405: {
		Fstart: int32(0x27bf),
		Fend:   int32(0x27bf),
	},
	406: {
		Fstart: int32(0x2934),
		Fend:   int32(0x2935),
	},
	407: {
		Fstart: int32(0x2b05),
		Fend:   int32(0x2b07),
	},
	408: {
		Fstart: int32(0x2b1b),
		Fend:   int32(0x2b1c),
	},
	409: {
		Fstart: int32(0x2b50),
		Fend:   int32(0x2b50),
	},
	410: {
		Fstart: int32(0x2b55),
		Fend:   int32(0x2b55),
	},
	411: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2ce4),
	},
	412: {
		Fstart: int32(0x2ceb),
		Fend:   int32(0x2cf3),
	},
	413: {
		Fstart: int32(0x2cfd),
		Fend:   int32(0x2cfd),
	},
	414: {
		Fstart: int32(0x2d00),
		Fend:   int32(0x2d25),
	},
	415: {
		Fstart: int32(0x2d27),
		Fend:   int32(0x2d27),
	},
	416: {
		Fstart: int32(0x2d2d),
		Fend:   int32(0x2d2d),
	},
	417: {
		Fstart: int32(0x2d30),
		Fend:   int32(0x2d67),
	},
	418: {
		Fstart: int32(0x2d6f),
		Fend:   int32(0x2d6f),
	},
	419: {
		Fstart: int32(0x2d7f),
		Fend:   int32(0x2d96),
	},
	420: {
		Fstart: int32(0x2da0),
		Fend:   int32(0x2da6),
	},
	421: {
		Fstart: int32(0x2da8),
		Fend:   int32(0x2dae),
	},
	422: {
		Fstart: int32(0x2db0),
		Fend:   int32(0x2db6),
	},
	423: {
		Fstart: int32(0x2db8),
		Fend:   int32(0x2dbe),
	},
	424: {
		Fstart: int32(0x2dc0),
		Fend:   int32(0x2dc6),
	},
	425: {
		Fstart: int32(0x2dc8),
		Fend:   int32(0x2dce),
	},
	426: {
		Fstart: int32(0x2dd0),
		Fend:   int32(0x2dd6),
	},
	427: {
		Fstart: int32(0x2dd8),
		Fend:   int32(0x2dde),
	},
	428: {
		Fstart: int32(0x2de0),
		Fend:   int32(0x2dff),
	},
	429: {
		Fstart: int32(0x2e2f),
		Fend:   int32(0x2e2f),
	},
	430: {
		Fstart: int32(0x3005),
		Fend:   int32(0x3007),
	},
	431: {
		Fstart: int32(0x3021),
		Fend:   int32(0x3035),
	},
	432: {
		Fstart: int32(0x3038),
		Fend:   int32(0x303d),
	},
	433: {
		Fstart: int32(0x3041),
		Fend:   int32(0x3096),
	},
	434: {
		Fstart: int32(0x3099),
		Fend:   int32(0x309a),
	},
	435: {
		Fstart: int32(0x309d),
		Fend:   int32(0x309f),
	},
	436: {
		Fstart: int32(0x30a1),
		Fend:   int32(0x30fa),
	},
	437: {
		Fstart: int32(0x30fc),
		Fend:   int32(0x30ff),
	},
	438: {
		Fstart: int32(0x3105),
		Fend:   int32(0x312f),
	},
	439: {
		Fstart: int32(0x3131),
		Fend:   int32(0x318e),
	},
	440: {
		Fstart: int32(0x3192),
		Fend:   int32(0x3195),
	},
	441: {
		Fstart: int32(0x31a0),
		Fend:   int32(0x31bf),
	},
	442: {
		Fstart: int32(0x31f0),
		Fend:   int32(0x31ff),
	},
	443: {
		Fstart: int32(0x3220),
		Fend:   int32(0x3229),
	},
	444: {
		Fstart: int32(0x3248),
		Fend:   int32(0x324f),
	},
	445: {
		Fstart: int32(0x3251),
		Fend:   int32(0x325f),
	},
	446: {
		Fstart: int32(0x3280),
		Fend:   int32(0x3289),
	},
	447: {
		Fstart: int32(0x3297),
		Fend:   int32(0x3297),
	},
	448: {
		Fstart: int32(0x3299),
		Fend:   int32(0x3299),
	},
	449: {
		Fstart: int32(0x32b1),
		Fend:   int32(0x32bf),
	},
	450: {
		Fstart: int32(0x3400),
		Fend:   int32(0x3400),
	},
	451: {
		Fstart: int32(0x4dbf),
		Fend:   int32(0x4dbf),
	},
	452: {
		Fstart: int32(0x4e00),
		Fend:   int32(0xa48c),
	},
	453: {
		Fstart: int32(0xa4d0),
		Fend:   int32(0xa4fd),
	},
	454: {
		Fstart: int32(0xa500),
		Fend:   int32(0xa60c),
	},
	455: {
		Fstart: int32(0xa610),
		Fend:   int32(0xa62b),
	},
	456: {
		Fstart: int32(0xa640),
		Fend:   int32(0xa672),
	},
	457: {
		Fstart: int32(0xa674),
		Fend:   int32(0xa67d),
	},
	458: {
		Fstart: int32(0xa67f),
		Fend:   int32(0xa6f1),
	},
	459: {
		Fstart: int32(0xa717),
		Fend:   int32(0xa71f),
	},
	460: {
		Fstart: int32(0xa722),
		Fend:   int32(0xa788),
	},
	461: {
		Fstart: int32(0xa78b),
		Fend:   int32(0xa7ca),
	},
	462: {
		Fstart: int32(0xa7d0),
		Fend:   int32(0xa7d1),
	},
	463: {
		Fstart: int32(0xa7d3),
		Fend:   int32(0xa7d3),
	},
	464: {
		Fstart: int32(0xa7d5),
		Fend:   int32(0xa7d9),
	},
	465: {
		Fstart: int32(0xa7f2),
		Fend:   int32(0xa827),
	},
	466: {
		Fstart: int32(0xa82c),
		Fend:   int32(0xa82c),
	},
	467: {
		Fstart: int32(0xa830),
		Fend:   int32(0xa835),
	},
	468: {
		Fstart: int32(0xa840),
		Fend:   int32(0xa873),
	},
	469: {
		Fstart: int32(0xa880),
		Fend:   int32(0xa8c5),
	},
	470: {
		Fstart: int32(0xa8d0),
		Fend:   int32(0xa8d9),
	},
	471: {
		Fstart: int32(0xa8e0),
		Fend:   int32(0xa8f7),
	},
	472: {
		Fstart: int32(0xa8fb),
		Fend:   int32(0xa8fb),
	},
	473: {
		Fstart: int32(0xa8fd),
		Fend:   int32(0xa92d),
	},
	474: {
		Fstart: int32(0xa930),
		Fend:   int32(0xa953),
	},
	475: {
		Fstart: int32(0xa960),
		Fend:   int32(0xa97c),
	},
	476: {
		Fstart: int32(0xa980),
		Fend:   int32(0xa9c0),
	},
	477: {
		Fstart: int32(0xa9cf),
		Fend:   int32(0xa9d9),
	},
	478: {
		Fstart: int32(0xa9e0),
		Fend:   int32(0xa9fe),
	},
	479: {
		Fstart: int32(0xaa00),
		Fend:   int32(0xaa36),
	},
	480: {
		Fstart: int32(0xaa40),
		Fend:   int32(0xaa4d),
	},
	481: {
		Fstart: int32(0xaa50),
		Fend:   int32(0xaa59),
	},
	482: {
		Fstart: int32(0xaa60),
		Fend:   int32(0xaa76),
	},
	483: {
		Fstart: int32(0xaa7a),
		Fend:   int32(0xaac2),
	},
	484: {
		Fstart: int32(0xaadb),
		Fend:   int32(0xaadd),
	},
	485: {
		Fstart: int32(0xaae0),
		Fend:   int32(0xaaef),
	},
	486: {
		Fstart: int32(0xaaf2),
		Fend:   int32(0xaaf6),
	},
	487: {
		Fstart: int32(0xab01),
		Fend:   int32(0xab06),
	},
	488: {
		Fstart: int32(0xab09),
		Fend:   int32(0xab0e),
	},
	489: {
		Fstart: int32(0xab11),
		Fend:   int32(0xab16),
	},
	490: {
		Fstart: int32(0xab20),
		Fend:   int32(0xab26),
	},
	491: {
		Fstart: int32(0xab28),
		Fend:   int32(0xab2e),
	},
	492: {
		Fstart: int32(0xab30),
		Fend:   int32(0xab5a),
	},
	493: {
		Fstart: int32(0xab5c),
		Fend:   int32(0xab69),
	},
	494: {
		Fstart: int32(0xab70),
		Fend:   int32(0xabea),
	},
	495: {
		Fstart: int32(0xabec),
		Fend:   int32(0xabed),
	},
	496: {
		Fstart: int32(0xabf0),
		Fend:   int32(0xabf9),
	},
	497: {
		Fstart: int32(0xac00),
		Fend:   int32(0xac00),
	},
	498: {
		Fstart: int32(0xd7a3),
		Fend:   int32(0xd7a3),
	},
	499: {
		Fstart: int32(0xd7b0),
		Fend:   int32(0xd7c6),
	},
	500: {
		Fstart: int32(0xd7cb),
		Fend:   int32(0xd7fb),
	},
	501: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfa6d),
	},
	502: {
		Fstart: int32(0xfa70),
		Fend:   int32(0xfad9),
	},
	503: {
		Fstart: int32(0xfb00),
		Fend:   int32(0xfb06),
	},
	504: {
		Fstart: int32(0xfb13),
		Fend:   int32(0xfb17),
	},
	505: {
		Fstart: int32(0xfb1d),
		Fend:   int32(0xfb28),
	},
	506: {
		Fstart: int32(0xfb2a),
		Fend:   int32(0xfb36),
	},
	507: {
		Fstart: int32(0xfb38),
		Fend:   int32(0xfb3c),
	},
	508: {
		Fstart: int32(0xfb3e),
		Fend:   int32(0xfb3e),
	},
	509: {
		Fstart: int32(0xfb40),
		Fend:   int32(0xfb41),
	},
	510: {
		Fstart: int32(0xfb43),
		Fend:   int32(0xfb44),
	},
	511: {
		Fstart: int32(0xfb46),
		Fend:   int32(0xfbb1),
	},
	512: {
		Fstart: int32(0xfbd3),
		Fend:   int32(0xfd3d),
	},
	513: {
		Fstart: int32(0xfd50),
		Fend:   int32(0xfd8f),
	},
	514: {
		Fstart: int32(0xfd92),
		Fend:   int32(0xfdc7),
	},
	515: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfdfb),
	},
	516: {
		Fstart: int32(0xfe00),
		Fend:   int32(0xfe0f),
	},
	517: {
		Fstart: int32(0xfe20),
		Fend:   int32(0xfe2f),
	},
	518: {
		Fstart: int32(0xfe70),
		Fend:   int32(0xfe74),
	},
	519: {
		Fstart: int32(0xfe76),
		Fend:   int32(0xfefc),
	},
	520: {
		Fstart: int32(0xff10),
		Fend:   int32(0xff19),
	},
	521: {
		Fstart: int32(0xff21),
		Fend:   int32(0xff3a),
	},
	522: {
		Fstart: int32(0xff41),
		Fend:   int32(0xff5a),
	},
	523: {
		Fstart: int32(0xff66),
		Fend:   int32(0xffbe),
	},
	524: {
		Fstart: int32(0xffc2),
		Fend:   int32(0xffc7),
	},
	525: {
		Fstart: int32(0xffca),
		Fend:   int32(0xffcf),
	},
	526: {
		Fstart: int32(0xffd2),
		Fend:   int32(0xffd7),
	},
	527: {
		Fstart: int32(0xffda),
		Fend:   int32(0xffdc),
	},
	528: {
		Fstart: int32(0x10000),
		Fend:   int32(0x1000b),
	},
	529: {
		Fstart: int32(0x1000d),
		Fend:   int32(0x10026),
	},
	530: {
		Fstart: int32(0x10028),
		Fend:   int32(0x1003a),
	},
	531: {
		Fstart: int32(0x1003c),
		Fend:   int32(0x1003d),
	},
	532: {
		Fstart: int32(0x1003f),
		Fend:   int32(0x1004d),
	},
	533: {
		Fstart: int32(0x10050),
		Fend:   int32(0x1005d),
	},
	534: {
		Fstart: int32(0x10080),
		Fend:   int32(0x100fa),
	},
	535: {
		Fstart: int32(0x10107),
		Fend:   int32(0x10133),
	},
	536: {
		Fstart: int32(0x10140),
		Fend:   int32(0x10178),
	},
	537: {
		Fstart: int32(0x1018a),
		Fend:   int32(0x1018b),
	},
	538: {
		Fstart: int32(0x101fd),
		Fend:   int32(0x101fd),
	},
	539: {
		Fstart: int32(0x10280),
		Fend:   int32(0x1029c),
	},
	540: {
		Fstart: int32(0x102a0),
		Fend:   int32(0x102d0),
	},
	541: {
		Fstart: int32(0x102e0),
		Fend:   int32(0x102fb),
	},
	542: {
		Fstart: int32(0x10300),
		Fend:   int32(0x10323),
	},
	543: {
		Fstart: int32(0x1032d),
		Fend:   int32(0x1034a),
	},
	544: {
		Fstart: int32(0x10350),
		Fend:   int32(0x1037a),
	},
	545: {
		Fstart: int32(0x10380),
		Fend:   int32(0x1039d),
	},
	546: {
		Fstart: int32(0x103a0),
		Fend:   int32(0x103c3),
	},
	547: {
		Fstart: int32(0x103c8),
		Fend:   int32(0x103cf),
	},
	548: {
		Fstart: int32(0x103d1),
		Fend:   int32(0x103d5),
	},
	549: {
		Fstart: int32(0x10400),
		Fend:   int32(0x1049d),
	},
	550: {
		Fstart: int32(0x104a0),
		Fend:   int32(0x104a9),
	},
	551: {
		Fstart: int32(0x104b0),
		Fend:   int32(0x104d3),
	},
	552: {
		Fstart: int32(0x104d8),
		Fend:   int32(0x104fb),
	},
	553: {
		Fstart: int32(0x10500),
		Fend:   int32(0x10527),
	},
	554: {
		Fstart: int32(0x10530),
		Fend:   int32(0x10563),
	},
	555: {
		Fstart: int32(0x10570),
		Fend:   int32(0x1057a),
	},
	556: {
		Fstart: int32(0x1057c),
		Fend:   int32(0x1058a),
	},
	557: {
		Fstart: int32(0x1058c),
		Fend:   int32(0x10592),
	},
	558: {
		Fstart: int32(0x10594),
		Fend:   int32(0x10595),
	},
	559: {
		Fstart: int32(0x10597),
		Fend:   int32(0x105a1),
	},
	560: {
		Fstart: int32(0x105a3),
		Fend:   int32(0x105b1),
	},
	561: {
		Fstart: int32(0x105b3),
		Fend:   int32(0x105b9),
	},
	562: {
		Fstart: int32(0x105bb),
		Fend:   int32(0x105bc),
	},
	563: {
		Fstart: int32(0x10600),
		Fend:   int32(0x10736),
	},
	564: {
		Fstart: int32(0x10740),
		Fend:   int32(0x10755),
	},
	565: {
		Fstart: int32(0x10760),
		Fend:   int32(0x10767),
	},
	566: {
		Fstart: int32(0x10780),
		Fend:   int32(0x10785),
	},
	567: {
		Fstart: int32(0x10787),
		Fend:   int32(0x107b0),
	},
	568: {
		Fstart: int32(0x107b2),
		Fend:   int32(0x107ba),
	},
	569: {
		Fstart: int32(0x10800),
		Fend:   int32(0x10805),
	},
	570: {
		Fstart: int32(0x10808),
		Fend:   int32(0x10808),
	},
	571: {
		Fstart: int32(0x1080a),
		Fend:   int32(0x10835),
	},
	572: {
		Fstart: int32(0x10837),
		Fend:   int32(0x10838),
	},
	573: {
		Fstart: int32(0x1083c),
		Fend:   int32(0x1083c),
	},
	574: {
		Fstart: int32(0x1083f),
		Fend:   int32(0x10855),
	},
	575: {
		Fstart: int32(0x10858),
		Fend:   int32(0x10876),
	},
	576: {
		Fstart: int32(0x10879),
		Fend:   int32(0x1089e),
	},
	577: {
		Fstart: int32(0x108a7),
		Fend:   int32(0x108af),
	},
	578: {
		Fstart: int32(0x108e0),
		Fend:   int32(0x108f2),
	},
	579: {
		Fstart: int32(0x108f4),
		Fend:   int32(0x108f5),
	},
	580: {
		Fstart: int32(0x108fb),
		Fend:   int32(0x1091b),
	},
	581: {
		Fstart: int32(0x1f004),
		Fend:   int32(0x1f004),
	},
	582: {
		Fstart: int32(0x1f0cf),
		Fend:   int32(0x1f0cf),
	},
	583: {
		Fstart: int32(0x1f170),
		Fend:   int32(0x1f171),
	},
	584: {
		Fstart: int32(0x1f17e),
		Fend:   int32(0x1f17f),
	},
	585: {
		Fstart: int32(0x1f18e),
		Fend:   int32(0x1f18e),
	},
	586: {
		Fstart: int32(0x1f191),
		Fend:   int32(0x1f19a),
	},
	587: {
		Fstart: int32(0x1f1e6),
		Fend:   int32(0x1f1ff),
	},
	588: {
		Fstart: int32(0x1f201),
		Fend:   int32(0x1f202),
	},
	589: {
		Fstart: int32(0x1f21a),
		Fend:   int32(0x1f21a),
	},
	590: {
		Fstart: int32(0x1f22f),
		Fend:   int32(0x1f22f),
	},
	591: {
		Fstart: int32(0x1f232),
		Fend:   int32(0x1f23a),
	},
	592: {
		Fstart: int32(0x1f250),
		Fend:   int32(0x1f251),
	},
	593: {
		Fstart: int32(0x1f300),
		Fend:   int32(0x1f321),
	},
	594: {
		Fstart: int32(0x1f324),
		Fend:   int32(0x1f393),
	},
	595: {
		Fstart: int32(0x1f396),
		Fend:   int32(0x1f397),
	},
	596: {
		Fstart: int32(0x1f399),
		Fend:   int32(0x1f39b),
	},
	597: {
		Fstart: int32(0x1f39e),
		Fend:   int32(0x1f3f0),
	},
	598: {
		Fstart: int32(0x1f3f3),
		Fend:   int32(0x1f3f5),
	},
	599: {
		Fstart: int32(0x1f3f7),
		Fend:   int32(0x1f4fd),
	},
	600: {
		Fstart: int32(0x1f4ff),
		Fend:   int32(0x1f53d),
	},
	601: {
		Fstart: int32(0x1f549),
		Fend:   int32(0x1f54e),
	},
	602: {
		Fstart: int32(0x1f550),
		Fend:   int32(0x1f567),
	},
	603: {
		Fstart: int32(0x1f56f),
		Fend:   int32(0x1f570),
	},
	604: {
		Fstart: int32(0x1f573),
		Fend:   int32(0x1f57a),
	},
	605: {
		Fstart: int32(0x1f587),
		Fend:   int32(0x1f587),
	},
	606: {
		Fstart: int32(0x1f58a),
		Fend:   int32(0x1f58d),
	},
	607: {
		Fstart: int32(0x1f590),
		Fend:   int32(0x1f590),
	},
	608: {
		Fstart: int32(0x1f595),
		Fend:   int32(0x1f596),
	},
	609: {
		Fstart: int32(0x1f5a4),
		Fend:   int32(0x1f5a5),
	},
	610: {
		Fstart: int32(0x1f5a8),
		Fend:   int32(0x1f5a8),
	},
	611: {
		Fstart: int32(0x1f5b1),
		Fend:   int32(0x1f5b2),
	},
	612: {
		Fstart: int32(0x1f5bc),
		Fend:   int32(0x1f5bc),
	},
	613: {
		Fstart: int32(0x1f5c2),
		Fend:   int32(0x1f5c4),
	},
	614: {
		Fstart: int32(0x1f5d1),
		Fend:   int32(0x1f5d3),
	},
	615: {
		Fstart: int32(0x1f5dc),
		Fend:   int32(0x1f5de),
	},
	616: {
		Fstart: int32(0x1f5e1),
		Fend:   int32(0x1f5e1),
	},
	617: {
		Fstart: int32(0x1f5e3),
		Fend:   int32(0x1f5e3),
	},
	618: {
		Fstart: int32(0x1f5e8),
		Fend:   int32(0x1f5e8),
	},
	619: {
		Fstart: int32(0x1f5ef),
		Fend:   int32(0x1f5ef),
	},
	620: {
		Fstart: int32(0x1f5f3),
		Fend:   int32(0x1f5f3),
	},
	621: {
		Fstart: int32(0x1f5fa),
		Fend:   int32(0x1f64f),
	},
	622: {
		Fstart: int32(0x1f680),
		Fend:   int32(0x1f6c5),
	},
	623: {
		Fstart: int32(0x1f6cb),
		Fend:   int32(0x1f6d2),
	},
	624: {
		Fstart: int32(0x1f6d5),
		Fend:   int32(0x1f6d7),
	},
	625: {
		Fstart: int32(0x1f6dc),
		Fend:   int32(0x1f6e5),
	},
	626: {
		Fstart: int32(0x1f6e9),
		Fend:   int32(0x1f6e9),
	},
	627: {
		Fstart: int32(0x1f6eb),
		Fend:   int32(0x1f6ec),
	},
	628: {
		Fstart: int32(0x1f6f0),
		Fend:   int32(0x1f6f0),
	},
	629: {
		Fstart: int32(0x1f6f3),
		Fend:   int32(0x1f6fc),
	},
	630: {
		Fstart: int32(0x1f7e0),
		Fend:   int32(0x1f7eb),
	},
	631: {
		Fstart: int32(0x1f7f0),
		Fend:   int32(0x1f7f0),
	},
	632: {
		Fstart: int32(0x1f90c),
		Fend:   int32(0x1f93a),
	},
	633: {
		Fstart: int32(0x1f93c),
		Fend:   int32(0x1f945),
	},
	634: {
		Fstart: int32(0x1f947),
		Fend:   int32(0x1f9ff),
	},
	635: {
		Fstart: int32(0x1fa70),
		Fend:   int32(0x1fa7c),
	},
	636: {
		Fstart: int32(0x1fa80),
		Fend:   int32(0x1fa88),
	},
	637: {
		Fstart: int32(0x1fa90),
		Fend:   int32(0x1fabd),
	},
	638: {
		Fstart: int32(0x1fabf),
		Fend:   int32(0x1fac5),
	},
	639: {
		Fstart: int32(0x1face),
		Fend:   int32(0x1fadb),
	},
	640: {
		Fstart: int32(0x1fae0),
		Fend:   int32(0x1fae8),
	},
	641: {
		Fstart: int32(0x1faf0),
		Fend:   int32(0x1faf8),
	},
}

var sym__identifier_char_character_set_1 = [499]TSCharacterRange{
	0: {
		Fstart: int32('!'),
		Fend:   int32('!'),
	},
	1: {
		Fstart: int32('#'),
		Fend:   int32('\''),
	},
	2: {
		Fstart: int32('*'),
		Fend:   int32('+'),
	},
	3: {
		Fstart: int32('-'),
		Fend:   int32('.'),
	},
	4: {
		Fstart: int32('0'),
		Fend:   int32(':'),
	},
	5: {
		Fstart: int32('?'),
		Fend:   int32('Z'),
	},
	6: {
		Fstart: int32('^'),
		Fend:   int32('_'),
	},
	7: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	8: {
		Fstart: int32('|'),
		Fend:   int32('|'),
	},
	9: {
		Fstart: int32('~'),
		Fend:   int32('~'),
	},
	10: {
		Fstart: int32(0xaa),
		Fend:   int32(0xaa),
	},
	11: {
		Fstart: int32(0xb2),
		Fend:   int32(0xb3),
	},
	12: {
		Fstart: int32(0xb5),
		Fend:   int32(0xb5),
	},
	13: {
		Fstart: int32(0xb9),
		Fend:   int32(0xba),
	},
	14: {
		Fstart: int32(0xbc),
		Fend:   int32(0xbe),
	},
	15: {
		Fstart: int32(0xc0),
		Fend:   int32(0xd6),
	},
	16: {
		Fstart: int32(0xd8),
		Fend:   int32(0xf6),
	},
	17: {
		Fstart: int32(0xf8),
		Fend:   int32(0x2c1),
	},
	18: {
		Fstart: int32(0x2c6),
		Fend:   int32(0x2d1),
	},
	19: {
		Fstart: int32(0x2e0),
		Fend:   int32(0x2e4),
	},
	20: {
		Fstart: int32(0x2ec),
		Fend:   int32(0x2ec),
	},
	21: {
		Fstart: int32(0x2ee),
		Fend:   int32(0x2ee),
	},
	22: {
		Fstart: int32(0x300),
		Fend:   int32(0x374),
	},
	23: {
		Fstart: int32(0x376),
		Fend:   int32(0x377),
	},
	24: {
		Fstart: int32(0x37a),
		Fend:   int32(0x37d),
	},
	25: {
		Fstart: int32(0x37f),
		Fend:   int32(0x37f),
	},
	26: {
		Fstart: int32(0x386),
		Fend:   int32(0x386),
	},
	27: {
		Fstart: int32(0x388),
		Fend:   int32(0x38a),
	},
	28: {
		Fstart: int32(0x38c),
		Fend:   int32(0x38c),
	},
	29: {
		Fstart: int32(0x38e),
		Fend:   int32(0x3a1),
	},
	30: {
		Fstart: int32(0x3a3),
		Fend:   int32(0x3f5),
	},
	31: {
		Fstart: int32(0x3f7),
		Fend:   int32(0x481),
	},
	32: {
		Fstart: int32(0x483),
		Fend:   int32(0x52f),
	},
	33: {
		Fstart: int32(0x531),
		Fend:   int32(0x556),
	},
	34: {
		Fstart: int32(0x559),
		Fend:   int32(0x559),
	},
	35: {
		Fstart: int32(0x560),
		Fend:   int32(0x588),
	},
	36: {
		Fstart: int32(0x591),
		Fend:   int32(0x5bd),
	},
	37: {
		Fstart: int32(0x5bf),
		Fend:   int32(0x5bf),
	},
	38: {
		Fstart: int32(0x5c1),
		Fend:   int32(0x5c2),
	},
	39: {
		Fstart: int32(0x5c4),
		Fend:   int32(0x5c5),
	},
	40: {
		Fstart: int32(0x5c7),
		Fend:   int32(0x5c7),
	},
	41: {
		Fstart: int32(0x5d0),
		Fend:   int32(0x5ea),
	},
	42: {
		Fstart: int32(0x5ef),
		Fend:   int32(0x5f2),
	},
	43: {
		Fstart: int32(0x610),
		Fend:   int32(0x61a),
	},
	44: {
		Fstart: int32(0x620),
		Fend:   int32(0x669),
	},
	45: {
		Fstart: int32(0x66e),
		Fend:   int32(0x6d3),
	},
	46: {
		Fstart: int32(0x6d5),
		Fend:   int32(0x6dc),
	},
	47: {
		Fstart: int32(0x6df),
		Fend:   int32(0x6e8),
	},
	48: {
		Fstart: int32(0x6ea),
		Fend:   int32(0x6fc),
	},
	49: {
		Fstart: int32(0x6ff),
		Fend:   int32(0x6ff),
	},
	50: {
		Fstart: int32(0x710),
		Fend:   int32(0x74a),
	},
	51: {
		Fstart: int32(0x74d),
		Fend:   int32(0x7b1),
	},
	52: {
		Fstart: int32(0x7c0),
		Fend:   int32(0x7f5),
	},
	53: {
		Fstart: int32(0x7fa),
		Fend:   int32(0x7fa),
	},
	54: {
		Fstart: int32(0x7fd),
		Fend:   int32(0x7fd),
	},
	55: {
		Fstart: int32(0x800),
		Fend:   int32(0x82d),
	},
	56: {
		Fstart: int32(0x840),
		Fend:   int32(0x85b),
	},
	57: {
		Fstart: int32(0x860),
		Fend:   int32(0x86a),
	},
	58: {
		Fstart: int32(0x870),
		Fend:   int32(0x887),
	},
	59: {
		Fstart: int32(0x889),
		Fend:   int32(0x88e),
	},
	60: {
		Fstart: int32(0x898),
		Fend:   int32(0x8e1),
	},
	61: {
		Fstart: int32(0x8e3),
		Fend:   int32(0x963),
	},
	62: {
		Fstart: int32(0x966),
		Fend:   int32(0x96f),
	},
	63: {
		Fstart: int32(0x971),
		Fend:   int32(0x983),
	},
	64: {
		Fstart: int32(0x985),
		Fend:   int32(0x98c),
	},
	65: {
		Fstart: int32(0x98f),
		Fend:   int32(0x990),
	},
	66: {
		Fstart: int32(0x993),
		Fend:   int32(0x9a8),
	},
	67: {
		Fstart: int32(0x9aa),
		Fend:   int32(0x9b0),
	},
	68: {
		Fstart: int32(0x9b2),
		Fend:   int32(0x9b2),
	},
	69: {
		Fstart: int32(0x9b6),
		Fend:   int32(0x9b9),
	},
	70: {
		Fstart: int32(0x9bc),
		Fend:   int32(0x9c4),
	},
	71: {
		Fstart: int32(0x9c7),
		Fend:   int32(0x9c8),
	},
	72: {
		Fstart: int32(0x9cb),
		Fend:   int32(0x9ce),
	},
	73: {
		Fstart: int32(0x9d7),
		Fend:   int32(0x9d7),
	},
	74: {
		Fstart: int32(0x9dc),
		Fend:   int32(0x9dd),
	},
	75: {
		Fstart: int32(0x9df),
		Fend:   int32(0x9e3),
	},
	76: {
		Fstart: int32(0x9e6),
		Fend:   int32(0x9f1),
	},
	77: {
		Fstart: int32(0x9f4),
		Fend:   int32(0x9f9),
	},
	78: {
		Fstart: int32(0x9fc),
		Fend:   int32(0x9fc),
	},
	79: {
		Fstart: int32(0x9fe),
		Fend:   int32(0x9fe),
	},
	80: {
		Fstart: int32(0xa01),
		Fend:   int32(0xa03),
	},
	81: {
		Fstart: int32(0xa05),
		Fend:   int32(0xa0a),
	},
	82: {
		Fstart: int32(0xa0f),
		Fend:   int32(0xa10),
	},
	83: {
		Fstart: int32(0xa13),
		Fend:   int32(0xa28),
	},
	84: {
		Fstart: int32(0xa2a),
		Fend:   int32(0xa30),
	},
	85: {
		Fstart: int32(0xa32),
		Fend:   int32(0xa33),
	},
	86: {
		Fstart: int32(0xa35),
		Fend:   int32(0xa36),
	},
	87: {
		Fstart: int32(0xa38),
		Fend:   int32(0xa39),
	},
	88: {
		Fstart: int32(0xa3c),
		Fend:   int32(0xa3c),
	},
	89: {
		Fstart: int32(0xa3e),
		Fend:   int32(0xa42),
	},
	90: {
		Fstart: int32(0xa47),
		Fend:   int32(0xa48),
	},
	91: {
		Fstart: int32(0xa4b),
		Fend:   int32(0xa4d),
	},
	92: {
		Fstart: int32(0xa51),
		Fend:   int32(0xa51),
	},
	93: {
		Fstart: int32(0xa59),
		Fend:   int32(0xa5c),
	},
	94: {
		Fstart: int32(0xa5e),
		Fend:   int32(0xa5e),
	},
	95: {
		Fstart: int32(0xa66),
		Fend:   int32(0xa75),
	},
	96: {
		Fstart: int32(0xa81),
		Fend:   int32(0xa83),
	},
	97: {
		Fstart: int32(0xa85),
		Fend:   int32(0xa8d),
	},
	98: {
		Fstart: int32(0xa8f),
		Fend:   int32(0xa91),
	},
	99: {
		Fstart: int32(0xa93),
		Fend:   int32(0xaa8),
	},
	100: {
		Fstart: int32(0xaaa),
		Fend:   int32(0xab0),
	},
	101: {
		Fstart: int32(0xab2),
		Fend:   int32(0xab3),
	},
	102: {
		Fstart: int32(0xab5),
		Fend:   int32(0xab9),
	},
	103: {
		Fstart: int32(0xabc),
		Fend:   int32(0xac5),
	},
	104: {
		Fstart: int32(0xac7),
		Fend:   int32(0xac9),
	},
	105: {
		Fstart: int32(0xacb),
		Fend:   int32(0xacd),
	},
	106: {
		Fstart: int32(0xad0),
		Fend:   int32(0xad0),
	},
	107: {
		Fstart: int32(0xae0),
		Fend:   int32(0xae3),
	},
	108: {
		Fstart: int32(0xae6),
		Fend:   int32(0xaef),
	},
	109: {
		Fstart: int32(0xaf9),
		Fend:   int32(0xaff),
	},
	110: {
		Fstart: int32(0xb01),
		Fend:   int32(0xb03),
	},
	111: {
		Fstart: int32(0xb05),
		Fend:   int32(0xb0c),
	},
	112: {
		Fstart: int32(0xb0f),
		Fend:   int32(0xb10),
	},
	113: {
		Fstart: int32(0xb13),
		Fend:   int32(0xb28),
	},
	114: {
		Fstart: int32(0xb2a),
		Fend:   int32(0xb30),
	},
	115: {
		Fstart: int32(0xb32),
		Fend:   int32(0xb33),
	},
	116: {
		Fstart: int32(0xb35),
		Fend:   int32(0xb39),
	},
	117: {
		Fstart: int32(0xb3c),
		Fend:   int32(0xb44),
	},
	118: {
		Fstart: int32(0xb47),
		Fend:   int32(0xb48),
	},
	119: {
		Fstart: int32(0xb4b),
		Fend:   int32(0xb4d),
	},
	120: {
		Fstart: int32(0xb55),
		Fend:   int32(0xb57),
	},
	121: {
		Fstart: int32(0xb5c),
		Fend:   int32(0xb5d),
	},
	122: {
		Fstart: int32(0xb5f),
		Fend:   int32(0xb63),
	},
	123: {
		Fstart: int32(0xb66),
		Fend:   int32(0xb6f),
	},
	124: {
		Fstart: int32(0xb71),
		Fend:   int32(0xb77),
	},
	125: {
		Fstart: int32(0xb82),
		Fend:   int32(0xb83),
	},
	126: {
		Fstart: int32(0xb85),
		Fend:   int32(0xb8a),
	},
	127: {
		Fstart: int32(0xb8e),
		Fend:   int32(0xb90),
	},
	128: {
		Fstart: int32(0xb92),
		Fend:   int32(0xb95),
	},
	129: {
		Fstart: int32(0xb99),
		Fend:   int32(0xb9a),
	},
	130: {
		Fstart: int32(0xb9c),
		Fend:   int32(0xb9c),
	},
	131: {
		Fstart: int32(0xb9e),
		Fend:   int32(0xb9f),
	},
	132: {
		Fstart: int32(0xba3),
		Fend:   int32(0xba4),
	},
	133: {
		Fstart: int32(0xba8),
		Fend:   int32(0xbaa),
	},
	134: {
		Fstart: int32(0xbae),
		Fend:   int32(0xbb9),
	},
	135: {
		Fstart: int32(0xbbe),
		Fend:   int32(0xbc2),
	},
	136: {
		Fstart: int32(0xbc6),
		Fend:   int32(0xbc8),
	},
	137: {
		Fstart: int32(0xbca),
		Fend:   int32(0xbcd),
	},
	138: {
		Fstart: int32(0xbd0),
		Fend:   int32(0xbd0),
	},
	139: {
		Fstart: int32(0xbd7),
		Fend:   int32(0xbd7),
	},
	140: {
		Fstart: int32(0xbe6),
		Fend:   int32(0xbf2),
	},
	141: {
		Fstart: int32(0xc00),
		Fend:   int32(0xc0c),
	},
	142: {
		Fstart: int32(0xc0e),
		Fend:   int32(0xc10),
	},
	143: {
		Fstart: int32(0xc12),
		Fend:   int32(0xc28),
	},
	144: {
		Fstart: int32(0xc2a),
		Fend:   int32(0xc39),
	},
	145: {
		Fstart: int32(0xc3c),
		Fend:   int32(0xc44),
	},
	146: {
		Fstart: int32(0xc46),
		Fend:   int32(0xc48),
	},
	147: {
		Fstart: int32(0xc4a),
		Fend:   int32(0xc4d),
	},
	148: {
		Fstart: int32(0xc55),
		Fend:   int32(0xc56),
	},
	149: {
		Fstart: int32(0xc58),
		Fend:   int32(0xc5a),
	},
	150: {
		Fstart: int32(0xc5d),
		Fend:   int32(0xc5d),
	},
	151: {
		Fstart: int32(0xc60),
		Fend:   int32(0xc63),
	},
	152: {
		Fstart: int32(0xc66),
		Fend:   int32(0xc6f),
	},
	153: {
		Fstart: int32(0xc78),
		Fend:   int32(0xc7e),
	},
	154: {
		Fstart: int32(0xc80),
		Fend:   int32(0xc83),
	},
	155: {
		Fstart: int32(0xc85),
		Fend:   int32(0xc8c),
	},
	156: {
		Fstart: int32(0xc8e),
		Fend:   int32(0xc90),
	},
	157: {
		Fstart: int32(0xc92),
		Fend:   int32(0xca8),
	},
	158: {
		Fstart: int32(0xcaa),
		Fend:   int32(0xcb3),
	},
	159: {
		Fstart: int32(0xcb5),
		Fend:   int32(0xcb9),
	},
	160: {
		Fstart: int32(0xcbc),
		Fend:   int32(0xcc4),
	},
	161: {
		Fstart: int32(0xcc6),
		Fend:   int32(0xcc8),
	},
	162: {
		Fstart: int32(0xcca),
		Fend:   int32(0xccd),
	},
	163: {
		Fstart: int32(0xcd5),
		Fend:   int32(0xcd6),
	},
	164: {
		Fstart: int32(0xcdd),
		Fend:   int32(0xcde),
	},
	165: {
		Fstart: int32(0xce0),
		Fend:   int32(0xce3),
	},
	166: {
		Fstart: int32(0xce6),
		Fend:   int32(0xcef),
	},
	167: {
		Fstart: int32(0xcf1),
		Fend:   int32(0xcf3),
	},
	168: {
		Fstart: int32(0xd00),
		Fend:   int32(0xd0c),
	},
	169: {
		Fstart: int32(0xd0e),
		Fend:   int32(0xd10),
	},
	170: {
		Fstart: int32(0xd12),
		Fend:   int32(0xd44),
	},
	171: {
		Fstart: int32(0xd46),
		Fend:   int32(0xd48),
	},
	172: {
		Fstart: int32(0xd4a),
		Fend:   int32(0xd4e),
	},
	173: {
		Fstart: int32(0xd54),
		Fend:   int32(0xd63),
	},
	174: {
		Fstart: int32(0xd66),
		Fend:   int32(0xd78),
	},
	175: {
		Fstart: int32(0xd7a),
		Fend:   int32(0xd7f),
	},
	176: {
		Fstart: int32(0xd81),
		Fend:   int32(0xd83),
	},
	177: {
		Fstart: int32(0xd85),
		Fend:   int32(0xd96),
	},
	178: {
		Fstart: int32(0xd9a),
		Fend:   int32(0xdb1),
	},
	179: {
		Fstart: int32(0xdb3),
		Fend:   int32(0xdbb),
	},
	180: {
		Fstart: int32(0xdbd),
		Fend:   int32(0xdbd),
	},
	181: {
		Fstart: int32(0xdc0),
		Fend:   int32(0xdc6),
	},
	182: {
		Fstart: int32(0xdca),
		Fend:   int32(0xdca),
	},
	183: {
		Fstart: int32(0xdcf),
		Fend:   int32(0xdd4),
	},
	184: {
		Fstart: int32(0xdd6),
		Fend:   int32(0xdd6),
	},
	185: {
		Fstart: int32(0xdd8),
		Fend:   int32(0xddf),
	},
	186: {
		Fstart: int32(0xde6),
		Fend:   int32(0xdef),
	},
	187: {
		Fstart: int32(0xdf2),
		Fend:   int32(0xdf3),
	},
	188: {
		Fstart: int32(0xe01),
		Fend:   int32(0xe3a),
	},
	189: {
		Fstart: int32(0xe40),
		Fend:   int32(0xe4e),
	},
	190: {
		Fstart: int32(0xe50),
		Fend:   int32(0xe59),
	},
	191: {
		Fstart: int32(0xe81),
		Fend:   int32(0xe82),
	},
	192: {
		Fstart: int32(0xe84),
		Fend:   int32(0xe84),
	},
	193: {
		Fstart: int32(0xe86),
		Fend:   int32(0xe8a),
	},
	194: {
		Fstart: int32(0xe8c),
		Fend:   int32(0xea3),
	},
	195: {
		Fstart: int32(0xea5),
		Fend:   int32(0xea5),
	},
	196: {
		Fstart: int32(0xea7),
		Fend:   int32(0xebd),
	},
	197: {
		Fstart: int32(0xec0),
		Fend:   int32(0xec4),
	},
	198: {
		Fstart: int32(0xec6),
		Fend:   int32(0xec6),
	},
	199: {
		Fstart: int32(0xec8),
		Fend:   int32(0xece),
	},
	200: {
		Fstart: int32(0xed0),
		Fend:   int32(0xed9),
	},
	201: {
		Fstart: int32(0xedc),
		Fend:   int32(0xedf),
	},
	202: {
		Fstart: int32(0xf00),
		Fend:   int32(0xf00),
	},
	203: {
		Fstart: int32(0xf18),
		Fend:   int32(0xf19),
	},
	204: {
		Fstart: int32(0xf20),
		Fend:   int32(0xf33),
	},
	205: {
		Fstart: int32(0xf35),
		Fend:   int32(0xf35),
	},
	206: {
		Fstart: int32(0xf37),
		Fend:   int32(0xf37),
	},
	207: {
		Fstart: int32(0xf39),
		Fend:   int32(0xf39),
	},
	208: {
		Fstart: int32(0xf3e),
		Fend:   int32(0xf47),
	},
	209: {
		Fstart: int32(0xf49),
		Fend:   int32(0xf6c),
	},
	210: {
		Fstart: int32(0xf71),
		Fend:   int32(0xf84),
	},
	211: {
		Fstart: int32(0xf86),
		Fend:   int32(0xf97),
	},
	212: {
		Fstart: int32(0xf99),
		Fend:   int32(0xfbc),
	},
	213: {
		Fstart: int32(0xfc6),
		Fend:   int32(0xfc6),
	},
	214: {
		Fstart: int32(0x1000),
		Fend:   int32(0x1049),
	},
	215: {
		Fstart: int32(0x1050),
		Fend:   int32(0x109d),
	},
	216: {
		Fstart: int32(0x10a0),
		Fend:   int32(0x10c5),
	},
	217: {
		Fstart: int32(0x10c7),
		Fend:   int32(0x10c7),
	},
	218: {
		Fstart: int32(0x10cd),
		Fend:   int32(0x10cd),
	},
	219: {
		Fstart: int32(0x10d0),
		Fend:   int32(0x10fa),
	},
	220: {
		Fstart: int32(0x10fc),
		Fend:   int32(0x1248),
	},
	221: {
		Fstart: int32(0x124a),
		Fend:   int32(0x124d),
	},
	222: {
		Fstart: int32(0x1250),
		Fend:   int32(0x1256),
	},
	223: {
		Fstart: int32(0x1258),
		Fend:   int32(0x1258),
	},
	224: {
		Fstart: int32(0x125a),
		Fend:   int32(0x125d),
	},
	225: {
		Fstart: int32(0x1260),
		Fend:   int32(0x1288),
	},
	226: {
		Fstart: int32(0x128a),
		Fend:   int32(0x128d),
	},
	227: {
		Fstart: int32(0x1290),
		Fend:   int32(0x12b0),
	},
	228: {
		Fstart: int32(0x12b2),
		Fend:   int32(0x12b5),
	},
	229: {
		Fstart: int32(0x12b8),
		Fend:   int32(0x12be),
	},
	230: {
		Fstart: int32(0x12c0),
		Fend:   int32(0x12c0),
	},
	231: {
		Fstart: int32(0x12c2),
		Fend:   int32(0x12c5),
	},
	232: {
		Fstart: int32(0x12c8),
		Fend:   int32(0x12d6),
	},
	233: {
		Fstart: int32(0x12d8),
		Fend:   int32(0x1310),
	},
	234: {
		Fstart: int32(0x1312),
		Fend:   int32(0x1315),
	},
	235: {
		Fstart: int32(0x1318),
		Fend:   int32(0x135a),
	},
	236: {
		Fstart: int32(0x135d),
		Fend:   int32(0x135f),
	},
	237: {
		Fstart: int32(0x1369),
		Fend:   int32(0x137c),
	},
	238: {
		Fstart: int32(0x1380),
		Fend:   int32(0x138f),
	},
	239: {
		Fstart: int32(0x13a0),
		Fend:   int32(0x13f5),
	},
	240: {
		Fstart: int32(0x13f8),
		Fend:   int32(0x13fd),
	},
	241: {
		Fstart: int32(0x1401),
		Fend:   int32(0x166c),
	},
	242: {
		Fstart: int32(0x166f),
		Fend:   int32(0x167f),
	},
	243: {
		Fstart: int32(0x1681),
		Fend:   int32(0x169a),
	},
	244: {
		Fstart: int32(0x16a0),
		Fend:   int32(0x16ea),
	},
	245: {
		Fstart: int32(0x16ee),
		Fend:   int32(0x16f8),
	},
	246: {
		Fstart: int32(0x1700),
		Fend:   int32(0x1715),
	},
	247: {
		Fstart: int32(0x171f),
		Fend:   int32(0x1734),
	},
	248: {
		Fstart: int32(0x1740),
		Fend:   int32(0x1753),
	},
	249: {
		Fstart: int32(0x1760),
		Fend:   int32(0x176c),
	},
	250: {
		Fstart: int32(0x176e),
		Fend:   int32(0x1770),
	},
	251: {
		Fstart: int32(0x1772),
		Fend:   int32(0x1773),
	},
	252: {
		Fstart: int32(0x1780),
		Fend:   int32(0x17d3),
	},
	253: {
		Fstart: int32(0x17d7),
		Fend:   int32(0x17d7),
	},
	254: {
		Fstart: int32(0x17dc),
		Fend:   int32(0x17dd),
	},
	255: {
		Fstart: int32(0x17e0),
		Fend:   int32(0x17e9),
	},
	256: {
		Fstart: int32(0x17f0),
		Fend:   int32(0x17f9),
	},
	257: {
		Fstart: int32(0x180b),
		Fend:   int32(0x180d),
	},
	258: {
		Fstart: int32(0x180f),
		Fend:   int32(0x1819),
	},
	259: {
		Fstart: int32(0x1820),
		Fend:   int32(0x1878),
	},
	260: {
		Fstart: int32(0x1880),
		Fend:   int32(0x18aa),
	},
	261: {
		Fstart: int32(0x18b0),
		Fend:   int32(0x18f5),
	},
	262: {
		Fstart: int32(0x1900),
		Fend:   int32(0x191e),
	},
	263: {
		Fstart: int32(0x1920),
		Fend:   int32(0x192b),
	},
	264: {
		Fstart: int32(0x1930),
		Fend:   int32(0x193b),
	},
	265: {
		Fstart: int32(0x1946),
		Fend:   int32(0x196d),
	},
	266: {
		Fstart: int32(0x1970),
		Fend:   int32(0x1974),
	},
	267: {
		Fstart: int32(0x1980),
		Fend:   int32(0x19ab),
	},
	268: {
		Fstart: int32(0x19b0),
		Fend:   int32(0x19c9),
	},
	269: {
		Fstart: int32(0x19d0),
		Fend:   int32(0x19da),
	},
	270: {
		Fstart: int32(0x1a00),
		Fend:   int32(0x1a1b),
	},
	271: {
		Fstart: int32(0x1a20),
		Fend:   int32(0x1a5e),
	},
	272: {
		Fstart: int32(0x1a60),
		Fend:   int32(0x1a7c),
	},
	273: {
		Fstart: int32(0x1a7f),
		Fend:   int32(0x1a89),
	},
	274: {
		Fstart: int32(0x1a90),
		Fend:   int32(0x1a99),
	},
	275: {
		Fstart: int32(0x1aa7),
		Fend:   int32(0x1aa7),
	},
	276: {
		Fstart: int32(0x1ab0),
		Fend:   int32(0x1ace),
	},
	277: {
		Fstart: int32(0x1b00),
		Fend:   int32(0x1b4c),
	},
	278: {
		Fstart: int32(0x1b50),
		Fend:   int32(0x1b59),
	},
	279: {
		Fstart: int32(0x1b6b),
		Fend:   int32(0x1b73),
	},
	280: {
		Fstart: int32(0x1b80),
		Fend:   int32(0x1bf3),
	},
	281: {
		Fstart: int32(0x1c00),
		Fend:   int32(0x1c37),
	},
	282: {
		Fstart: int32(0x1c40),
		Fend:   int32(0x1c49),
	},
	283: {
		Fstart: int32(0x1c4d),
		Fend:   int32(0x1c7d),
	},
	284: {
		Fstart: int32(0x1c80),
		Fend:   int32(0x1c88),
	},
	285: {
		Fstart: int32(0x1c90),
		Fend:   int32(0x1cba),
	},
	286: {
		Fstart: int32(0x1cbd),
		Fend:   int32(0x1cbf),
	},
	287: {
		Fstart: int32(0x1cd0),
		Fend:   int32(0x1cd2),
	},
	288: {
		Fstart: int32(0x1cd4),
		Fend:   int32(0x1cfa),
	},
	289: {
		Fstart: int32(0x1d00),
		Fend:   int32(0x1f15),
	},
	290: {
		Fstart: int32(0x1f18),
		Fend:   int32(0x1f1d),
	},
	291: {
		Fstart: int32(0x1f20),
		Fend:   int32(0x1f45),
	},
	292: {
		Fstart: int32(0x1f48),
		Fend:   int32(0x1f4d),
	},
	293: {
		Fstart: int32(0x1f50),
		Fend:   int32(0x1f57),
	},
	294: {
		Fstart: int32(0x1f59),
		Fend:   int32(0x1f59),
	},
	295: {
		Fstart: int32(0x1f5b),
		Fend:   int32(0x1f5b),
	},
	296: {
		Fstart: int32(0x1f5d),
		Fend:   int32(0x1f5d),
	},
	297: {
		Fstart: int32(0x1f5f),
		Fend:   int32(0x1f7d),
	},
	298: {
		Fstart: int32(0x1f80),
		Fend:   int32(0x1fb4),
	},
	299: {
		Fstart: int32(0x1fb6),
		Fend:   int32(0x1fbc),
	},
	300: {
		Fstart: int32(0x1fbe),
		Fend:   int32(0x1fbe),
	},
	301: {
		Fstart: int32(0x1fc2),
		Fend:   int32(0x1fc4),
	},
	302: {
		Fstart: int32(0x1fc6),
		Fend:   int32(0x1fcc),
	},
	303: {
		Fstart: int32(0x1fd0),
		Fend:   int32(0x1fd3),
	},
	304: {
		Fstart: int32(0x1fd6),
		Fend:   int32(0x1fdb),
	},
	305: {
		Fstart: int32(0x1fe0),
		Fend:   int32(0x1fec),
	},
	306: {
		Fstart: int32(0x1ff2),
		Fend:   int32(0x1ff4),
	},
	307: {
		Fstart: int32(0x1ff6),
		Fend:   int32(0x1ffc),
	},
	308: {
		Fstart: int32(0x2070),
		Fend:   int32(0x2071),
	},
	309: {
		Fstart: int32(0x2074),
		Fend:   int32(0x2079),
	},
	310: {
		Fstart: int32(0x207f),
		Fend:   int32(0x2089),
	},
	311: {
		Fstart: int32(0x2090),
		Fend:   int32(0x209c),
	},
	312: {
		Fstart: int32(0x20d0),
		Fend:   int32(0x20f0),
	},
	313: {
		Fstart: int32(0x2102),
		Fend:   int32(0x2102),
	},
	314: {
		Fstart: int32(0x2107),
		Fend:   int32(0x2107),
	},
	315: {
		Fstart: int32(0x210a),
		Fend:   int32(0x2113),
	},
	316: {
		Fstart: int32(0x2115),
		Fend:   int32(0x2115),
	},
	317: {
		Fstart: int32(0x2119),
		Fend:   int32(0x211d),
	},
	318: {
		Fstart: int32(0x2124),
		Fend:   int32(0x2124),
	},
	319: {
		Fstart: int32(0x2126),
		Fend:   int32(0x2126),
	},
	320: {
		Fstart: int32(0x2128),
		Fend:   int32(0x2128),
	},
	321: {
		Fstart: int32(0x212a),
		Fend:   int32(0x212d),
	},
	322: {
		Fstart: int32(0x212f),
		Fend:   int32(0x2139),
	},
	323: {
		Fstart: int32(0x213c),
		Fend:   int32(0x213f),
	},
	324: {
		Fstart: int32(0x2145),
		Fend:   int32(0x2149),
	},
	325: {
		Fstart: int32(0x214e),
		Fend:   int32(0x214e),
	},
	326: {
		Fstart: int32(0x2150),
		Fend:   int32(0x2189),
	},
	327: {
		Fstart: int32(0x2460),
		Fend:   int32(0x249b),
	},
	328: {
		Fstart: int32(0x24ea),
		Fend:   int32(0x24ff),
	},
	329: {
		Fstart: int32(0x2776),
		Fend:   int32(0x2793),
	},
	330: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2ce4),
	},
	331: {
		Fstart: int32(0x2ceb),
		Fend:   int32(0x2cf3),
	},
	332: {
		Fstart: int32(0x2cfd),
		Fend:   int32(0x2cfd),
	},
	333: {
		Fstart: int32(0x2d00),
		Fend:   int32(0x2d25),
	},
	334: {
		Fstart: int32(0x2d27),
		Fend:   int32(0x2d27),
	},
	335: {
		Fstart: int32(0x2d2d),
		Fend:   int32(0x2d2d),
	},
	336: {
		Fstart: int32(0x2d30),
		Fend:   int32(0x2d67),
	},
	337: {
		Fstart: int32(0x2d6f),
		Fend:   int32(0x2d6f),
	},
	338: {
		Fstart: int32(0x2d7f),
		Fend:   int32(0x2d96),
	},
	339: {
		Fstart: int32(0x2da0),
		Fend:   int32(0x2da6),
	},
	340: {
		Fstart: int32(0x2da8),
		Fend:   int32(0x2dae),
	},
	341: {
		Fstart: int32(0x2db0),
		Fend:   int32(0x2db6),
	},
	342: {
		Fstart: int32(0x2db8),
		Fend:   int32(0x2dbe),
	},
	343: {
		Fstart: int32(0x2dc0),
		Fend:   int32(0x2dc6),
	},
	344: {
		Fstart: int32(0x2dc8),
		Fend:   int32(0x2dce),
	},
	345: {
		Fstart: int32(0x2dd0),
		Fend:   int32(0x2dd6),
	},
	346: {
		Fstart: int32(0x2dd8),
		Fend:   int32(0x2dde),
	},
	347: {
		Fstart: int32(0x2de0),
		Fend:   int32(0x2dff),
	},
	348: {
		Fstart: int32(0x2e2f),
		Fend:   int32(0x2e2f),
	},
	349: {
		Fstart: int32(0x3005),
		Fend:   int32(0x3007),
	},
	350: {
		Fstart: int32(0x3021),
		Fend:   int32(0x302f),
	},
	351: {
		Fstart: int32(0x3031),
		Fend:   int32(0x3035),
	},
	352: {
		Fstart: int32(0x3038),
		Fend:   int32(0x303c),
	},
	353: {
		Fstart: int32(0x3041),
		Fend:   int32(0x3096),
	},
	354: {
		Fstart: int32(0x3099),
		Fend:   int32(0x309a),
	},
	355: {
		Fstart: int32(0x309d),
		Fend:   int32(0x309f),
	},
	356: {
		Fstart: int32(0x30a1),
		Fend:   int32(0x30fa),
	},
	357: {
		Fstart: int32(0x30fc),
		Fend:   int32(0x30ff),
	},
	358: {
		Fstart: int32(0x3105),
		Fend:   int32(0x312f),
	},
	359: {
		Fstart: int32(0x3131),
		Fend:   int32(0x318e),
	},
	360: {
		Fstart: int32(0x3192),
		Fend:   int32(0x3195),
	},
	361: {
		Fstart: int32(0x31a0),
		Fend:   int32(0x31bf),
	},
	362: {
		Fstart: int32(0x31f0),
		Fend:   int32(0x31ff),
	},
	363: {
		Fstart: int32(0x3220),
		Fend:   int32(0x3229),
	},
	364: {
		Fstart: int32(0x3248),
		Fend:   int32(0x324f),
	},
	365: {
		Fstart: int32(0x3251),
		Fend:   int32(0x325f),
	},
	366: {
		Fstart: int32(0x3280),
		Fend:   int32(0x3289),
	},
	367: {
		Fstart: int32(0x32b1),
		Fend:   int32(0x32bf),
	},
	368: {
		Fstart: int32(0x3400),
		Fend:   int32(0x3400),
	},
	369: {
		Fstart: int32(0x4dbf),
		Fend:   int32(0x4dbf),
	},
	370: {
		Fstart: int32(0x4e00),
		Fend:   int32(0xa48c),
	},
	371: {
		Fstart: int32(0xa4d0),
		Fend:   int32(0xa4fd),
	},
	372: {
		Fstart: int32(0xa500),
		Fend:   int32(0xa60c),
	},
	373: {
		Fstart: int32(0xa610),
		Fend:   int32(0xa62b),
	},
	374: {
		Fstart: int32(0xa640),
		Fend:   int32(0xa672),
	},
	375: {
		Fstart: int32(0xa674),
		Fend:   int32(0xa67d),
	},
	376: {
		Fstart: int32(0xa67f),
		Fend:   int32(0xa6f1),
	},
	377: {
		Fstart: int32(0xa717),
		Fend:   int32(0xa71f),
	},
	378: {
		Fstart: int32(0xa722),
		Fend:   int32(0xa788),
	},
	379: {
		Fstart: int32(0xa78b),
		Fend:   int32(0xa7ca),
	},
	380: {
		Fstart: int32(0xa7d0),
		Fend:   int32(0xa7d1),
	},
	381: {
		Fstart: int32(0xa7d3),
		Fend:   int32(0xa7d3),
	},
	382: {
		Fstart: int32(0xa7d5),
		Fend:   int32(0xa7d9),
	},
	383: {
		Fstart: int32(0xa7f2),
		Fend:   int32(0xa827),
	},
	384: {
		Fstart: int32(0xa82c),
		Fend:   int32(0xa82c),
	},
	385: {
		Fstart: int32(0xa830),
		Fend:   int32(0xa835),
	},
	386: {
		Fstart: int32(0xa840),
		Fend:   int32(0xa873),
	},
	387: {
		Fstart: int32(0xa880),
		Fend:   int32(0xa8c5),
	},
	388: {
		Fstart: int32(0xa8d0),
		Fend:   int32(0xa8d9),
	},
	389: {
		Fstart: int32(0xa8e0),
		Fend:   int32(0xa8f7),
	},
	390: {
		Fstart: int32(0xa8fb),
		Fend:   int32(0xa8fb),
	},
	391: {
		Fstart: int32(0xa8fd),
		Fend:   int32(0xa92d),
	},
	392: {
		Fstart: int32(0xa930),
		Fend:   int32(0xa953),
	},
	393: {
		Fstart: int32(0xa960),
		Fend:   int32(0xa97c),
	},
	394: {
		Fstart: int32(0xa980),
		Fend:   int32(0xa9c0),
	},
	395: {
		Fstart: int32(0xa9cf),
		Fend:   int32(0xa9d9),
	},
	396: {
		Fstart: int32(0xa9e0),
		Fend:   int32(0xa9fe),
	},
	397: {
		Fstart: int32(0xaa00),
		Fend:   int32(0xaa36),
	},
	398: {
		Fstart: int32(0xaa40),
		Fend:   int32(0xaa4d),
	},
	399: {
		Fstart: int32(0xaa50),
		Fend:   int32(0xaa59),
	},
	400: {
		Fstart: int32(0xaa60),
		Fend:   int32(0xaa76),
	},
	401: {
		Fstart: int32(0xaa7a),
		Fend:   int32(0xaac2),
	},
	402: {
		Fstart: int32(0xaadb),
		Fend:   int32(0xaadd),
	},
	403: {
		Fstart: int32(0xaae0),
		Fend:   int32(0xaaef),
	},
	404: {
		Fstart: int32(0xaaf2),
		Fend:   int32(0xaaf6),
	},
	405: {
		Fstart: int32(0xab01),
		Fend:   int32(0xab06),
	},
	406: {
		Fstart: int32(0xab09),
		Fend:   int32(0xab0e),
	},
	407: {
		Fstart: int32(0xab11),
		Fend:   int32(0xab16),
	},
	408: {
		Fstart: int32(0xab20),
		Fend:   int32(0xab26),
	},
	409: {
		Fstart: int32(0xab28),
		Fend:   int32(0xab2e),
	},
	410: {
		Fstart: int32(0xab30),
		Fend:   int32(0xab5a),
	},
	411: {
		Fstart: int32(0xab5c),
		Fend:   int32(0xab69),
	},
	412: {
		Fstart: int32(0xab70),
		Fend:   int32(0xabea),
	},
	413: {
		Fstart: int32(0xabec),
		Fend:   int32(0xabed),
	},
	414: {
		Fstart: int32(0xabf0),
		Fend:   int32(0xabf9),
	},
	415: {
		Fstart: int32(0xac00),
		Fend:   int32(0xac00),
	},
	416: {
		Fstart: int32(0xd7a3),
		Fend:   int32(0xd7a3),
	},
	417: {
		Fstart: int32(0xd7b0),
		Fend:   int32(0xd7c6),
	},
	418: {
		Fstart: int32(0xd7cb),
		Fend:   int32(0xd7fb),
	},
	419: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfa6d),
	},
	420: {
		Fstart: int32(0xfa70),
		Fend:   int32(0xfad9),
	},
	421: {
		Fstart: int32(0xfb00),
		Fend:   int32(0xfb06),
	},
	422: {
		Fstart: int32(0xfb13),
		Fend:   int32(0xfb17),
	},
	423: {
		Fstart: int32(0xfb1d),
		Fend:   int32(0xfb28),
	},
	424: {
		Fstart: int32(0xfb2a),
		Fend:   int32(0xfb36),
	},
	425: {
		Fstart: int32(0xfb38),
		Fend:   int32(0xfb3c),
	},
	426: {
		Fstart: int32(0xfb3e),
		Fend:   int32(0xfb3e),
	},
	427: {
		Fstart: int32(0xfb40),
		Fend:   int32(0xfb41),
	},
	428: {
		Fstart: int32(0xfb43),
		Fend:   int32(0xfb44),
	},
	429: {
		Fstart: int32(0xfb46),
		Fend:   int32(0xfbb1),
	},
	430: {
		Fstart: int32(0xfbd3),
		Fend:   int32(0xfd3d),
	},
	431: {
		Fstart: int32(0xfd50),
		Fend:   int32(0xfd8f),
	},
	432: {
		Fstart: int32(0xfd92),
		Fend:   int32(0xfdc7),
	},
	433: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfdfb),
	},
	434: {
		Fstart: int32(0xfe00),
		Fend:   int32(0xfe0f),
	},
	435: {
		Fstart: int32(0xfe20),
		Fend:   int32(0xfe2f),
	},
	436: {
		Fstart: int32(0xfe70),
		Fend:   int32(0xfe74),
	},
	437: {
		Fstart: int32(0xfe76),
		Fend:   int32(0xfefc),
	},
	438: {
		Fstart: int32(0xff10),
		Fend:   int32(0xff19),
	},
	439: {
		Fstart: int32(0xff21),
		Fend:   int32(0xff3a),
	},
	440: {
		Fstart: int32(0xff41),
		Fend:   int32(0xff5a),
	},
	441: {
		Fstart: int32(0xff66),
		Fend:   int32(0xffbe),
	},
	442: {
		Fstart: int32(0xffc2),
		Fend:   int32(0xffc7),
	},
	443: {
		Fstart: int32(0xffca),
		Fend:   int32(0xffcf),
	},
	444: {
		Fstart: int32(0xffd2),
		Fend:   int32(0xffd7),
	},
	445: {
		Fstart: int32(0xffda),
		Fend:   int32(0xffdc),
	},
	446: {
		Fstart: int32(0x10000),
		Fend:   int32(0x1000b),
	},
	447: {
		Fstart: int32(0x1000d),
		Fend:   int32(0x10026),
	},
	448: {
		Fstart: int32(0x10028),
		Fend:   int32(0x1003a),
	},
	449: {
		Fstart: int32(0x1003c),
		Fend:   int32(0x1003d),
	},
	450: {
		Fstart: int32(0x1003f),
		Fend:   int32(0x1004d),
	},
	451: {
		Fstart: int32(0x10050),
		Fend:   int32(0x1005d),
	},
	452: {
		Fstart: int32(0x10080),
		Fend:   int32(0x100fa),
	},
	453: {
		Fstart: int32(0x10107),
		Fend:   int32(0x10133),
	},
	454: {
		Fstart: int32(0x10140),
		Fend:   int32(0x10178),
	},
	455: {
		Fstart: int32(0x1018a),
		Fend:   int32(0x1018b),
	},
	456: {
		Fstart: int32(0x101fd),
		Fend:   int32(0x101fd),
	},
	457: {
		Fstart: int32(0x10280),
		Fend:   int32(0x1029c),
	},
	458: {
		Fstart: int32(0x102a0),
		Fend:   int32(0x102d0),
	},
	459: {
		Fstart: int32(0x102e0),
		Fend:   int32(0x102fb),
	},
	460: {
		Fstart: int32(0x10300),
		Fend:   int32(0x10323),
	},
	461: {
		Fstart: int32(0x1032d),
		Fend:   int32(0x1034a),
	},
	462: {
		Fstart: int32(0x10350),
		Fend:   int32(0x1037a),
	},
	463: {
		Fstart: int32(0x10380),
		Fend:   int32(0x1039d),
	},
	464: {
		Fstart: int32(0x103a0),
		Fend:   int32(0x103c3),
	},
	465: {
		Fstart: int32(0x103c8),
		Fend:   int32(0x103cf),
	},
	466: {
		Fstart: int32(0x103d1),
		Fend:   int32(0x103d5),
	},
	467: {
		Fstart: int32(0x10400),
		Fend:   int32(0x1049d),
	},
	468: {
		Fstart: int32(0x104a0),
		Fend:   int32(0x104a9),
	},
	469: {
		Fstart: int32(0x104b0),
		Fend:   int32(0x104d3),
	},
	470: {
		Fstart: int32(0x104d8),
		Fend:   int32(0x104fb),
	},
	471: {
		Fstart: int32(0x10500),
		Fend:   int32(0x10527),
	},
	472: {
		Fstart: int32(0x10530),
		Fend:   int32(0x10563),
	},
	473: {
		Fstart: int32(0x10570),
		Fend:   int32(0x1057a),
	},
	474: {
		Fstart: int32(0x1057c),
		Fend:   int32(0x1058a),
	},
	475: {
		Fstart: int32(0x1058c),
		Fend:   int32(0x10592),
	},
	476: {
		Fstart: int32(0x10594),
		Fend:   int32(0x10595),
	},
	477: {
		Fstart: int32(0x10597),
		Fend:   int32(0x105a1),
	},
	478: {
		Fstart: int32(0x105a3),
		Fend:   int32(0x105b1),
	},
	479: {
		Fstart: int32(0x105b3),
		Fend:   int32(0x105b9),
	},
	480: {
		Fstart: int32(0x105bb),
		Fend:   int32(0x105bc),
	},
	481: {
		Fstart: int32(0x10600),
		Fend:   int32(0x10736),
	},
	482: {
		Fstart: int32(0x10740),
		Fend:   int32(0x10755),
	},
	483: {
		Fstart: int32(0x10760),
		Fend:   int32(0x10767),
	},
	484: {
		Fstart: int32(0x10780),
		Fend:   int32(0x10785),
	},
	485: {
		Fstart: int32(0x10787),
		Fend:   int32(0x107b0),
	},
	486: {
		Fstart: int32(0x107b2),
		Fend:   int32(0x107ba),
	},
	487: {
		Fstart: int32(0x10800),
		Fend:   int32(0x10805),
	},
	488: {
		Fstart: int32(0x10808),
		Fend:   int32(0x10808),
	},
	489: {
		Fstart: int32(0x1080a),
		Fend:   int32(0x10835),
	},
	490: {
		Fstart: int32(0x10837),
		Fend:   int32(0x10838),
	},
	491: {
		Fstart: int32(0x1083c),
		Fend:   int32(0x1083c),
	},
	492: {
		Fstart: int32(0x1083f),
		Fend:   int32(0x10855),
	},
	493: {
		Fstart: int32(0x10858),
		Fend:   int32(0x10876),
	},
	494: {
		Fstart: int32(0x10879),
		Fend:   int32(0x1089e),
	},
	495: {
		Fstart: int32(0x108a7),
		Fend:   int32(0x108af),
	},
	496: {
		Fstart: int32(0x108e0),
		Fend:   int32(0x108f2),
	},
	497: {
		Fstart: int32(0x108f4),
		Fend:   int32(0x108f5),
	},
	498: {
		Fstart: int32(0x108fb),
		Fend:   int32(0x1091b),
	},
}

var sym___identifier_char_no_digit_character_set_1 = [499]TSCharacterRange{
	0: {
		Fstart: int32('!'),
		Fend:   int32('!'),
	},
	1: {
		Fstart: int32('#'),
		Fend:   int32('\''),
	},
	2: {
		Fstart: int32('*'),
		Fend:   int32('+'),
	},
	3: {
		Fstart: int32('-'),
		Fend:   int32('.'),
	},
	4: {
		Fstart: int32(':'),
		Fend:   int32(':'),
	},
	5: {
		Fstart: int32('?'),
		Fend:   int32('Z'),
	},
	6: {
		Fstart: int32('^'),
		Fend:   int32('_'),
	},
	7: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	8: {
		Fstart: int32('|'),
		Fend:   int32('|'),
	},
	9: {
		Fstart: int32('~'),
		Fend:   int32('~'),
	},
	10: {
		Fstart: int32(0xaa),
		Fend:   int32(0xaa),
	},
	11: {
		Fstart: int32(0xb2),
		Fend:   int32(0xb3),
	},
	12: {
		Fstart: int32(0xb5),
		Fend:   int32(0xb5),
	},
	13: {
		Fstart: int32(0xb9),
		Fend:   int32(0xba),
	},
	14: {
		Fstart: int32(0xbc),
		Fend:   int32(0xbe),
	},
	15: {
		Fstart: int32(0xc0),
		Fend:   int32(0xd6),
	},
	16: {
		Fstart: int32(0xd8),
		Fend:   int32(0xf6),
	},
	17: {
		Fstart: int32(0xf8),
		Fend:   int32(0x2c1),
	},
	18: {
		Fstart: int32(0x2c6),
		Fend:   int32(0x2d1),
	},
	19: {
		Fstart: int32(0x2e0),
		Fend:   int32(0x2e4),
	},
	20: {
		Fstart: int32(0x2ec),
		Fend:   int32(0x2ec),
	},
	21: {
		Fstart: int32(0x2ee),
		Fend:   int32(0x2ee),
	},
	22: {
		Fstart: int32(0x300),
		Fend:   int32(0x374),
	},
	23: {
		Fstart: int32(0x376),
		Fend:   int32(0x377),
	},
	24: {
		Fstart: int32(0x37a),
		Fend:   int32(0x37d),
	},
	25: {
		Fstart: int32(0x37f),
		Fend:   int32(0x37f),
	},
	26: {
		Fstart: int32(0x386),
		Fend:   int32(0x386),
	},
	27: {
		Fstart: int32(0x388),
		Fend:   int32(0x38a),
	},
	28: {
		Fstart: int32(0x38c),
		Fend:   int32(0x38c),
	},
	29: {
		Fstart: int32(0x38e),
		Fend:   int32(0x3a1),
	},
	30: {
		Fstart: int32(0x3a3),
		Fend:   int32(0x3f5),
	},
	31: {
		Fstart: int32(0x3f7),
		Fend:   int32(0x481),
	},
	32: {
		Fstart: int32(0x483),
		Fend:   int32(0x52f),
	},
	33: {
		Fstart: int32(0x531),
		Fend:   int32(0x556),
	},
	34: {
		Fstart: int32(0x559),
		Fend:   int32(0x559),
	},
	35: {
		Fstart: int32(0x560),
		Fend:   int32(0x588),
	},
	36: {
		Fstart: int32(0x591),
		Fend:   int32(0x5bd),
	},
	37: {
		Fstart: int32(0x5bf),
		Fend:   int32(0x5bf),
	},
	38: {
		Fstart: int32(0x5c1),
		Fend:   int32(0x5c2),
	},
	39: {
		Fstart: int32(0x5c4),
		Fend:   int32(0x5c5),
	},
	40: {
		Fstart: int32(0x5c7),
		Fend:   int32(0x5c7),
	},
	41: {
		Fstart: int32(0x5d0),
		Fend:   int32(0x5ea),
	},
	42: {
		Fstart: int32(0x5ef),
		Fend:   int32(0x5f2),
	},
	43: {
		Fstart: int32(0x610),
		Fend:   int32(0x61a),
	},
	44: {
		Fstart: int32(0x620),
		Fend:   int32(0x669),
	},
	45: {
		Fstart: int32(0x66e),
		Fend:   int32(0x6d3),
	},
	46: {
		Fstart: int32(0x6d5),
		Fend:   int32(0x6dc),
	},
	47: {
		Fstart: int32(0x6df),
		Fend:   int32(0x6e8),
	},
	48: {
		Fstart: int32(0x6ea),
		Fend:   int32(0x6fc),
	},
	49: {
		Fstart: int32(0x6ff),
		Fend:   int32(0x6ff),
	},
	50: {
		Fstart: int32(0x710),
		Fend:   int32(0x74a),
	},
	51: {
		Fstart: int32(0x74d),
		Fend:   int32(0x7b1),
	},
	52: {
		Fstart: int32(0x7c0),
		Fend:   int32(0x7f5),
	},
	53: {
		Fstart: int32(0x7fa),
		Fend:   int32(0x7fa),
	},
	54: {
		Fstart: int32(0x7fd),
		Fend:   int32(0x7fd),
	},
	55: {
		Fstart: int32(0x800),
		Fend:   int32(0x82d),
	},
	56: {
		Fstart: int32(0x840),
		Fend:   int32(0x85b),
	},
	57: {
		Fstart: int32(0x860),
		Fend:   int32(0x86a),
	},
	58: {
		Fstart: int32(0x870),
		Fend:   int32(0x887),
	},
	59: {
		Fstart: int32(0x889),
		Fend:   int32(0x88e),
	},
	60: {
		Fstart: int32(0x898),
		Fend:   int32(0x8e1),
	},
	61: {
		Fstart: int32(0x8e3),
		Fend:   int32(0x963),
	},
	62: {
		Fstart: int32(0x966),
		Fend:   int32(0x96f),
	},
	63: {
		Fstart: int32(0x971),
		Fend:   int32(0x983),
	},
	64: {
		Fstart: int32(0x985),
		Fend:   int32(0x98c),
	},
	65: {
		Fstart: int32(0x98f),
		Fend:   int32(0x990),
	},
	66: {
		Fstart: int32(0x993),
		Fend:   int32(0x9a8),
	},
	67: {
		Fstart: int32(0x9aa),
		Fend:   int32(0x9b0),
	},
	68: {
		Fstart: int32(0x9b2),
		Fend:   int32(0x9b2),
	},
	69: {
		Fstart: int32(0x9b6),
		Fend:   int32(0x9b9),
	},
	70: {
		Fstart: int32(0x9bc),
		Fend:   int32(0x9c4),
	},
	71: {
		Fstart: int32(0x9c7),
		Fend:   int32(0x9c8),
	},
	72: {
		Fstart: int32(0x9cb),
		Fend:   int32(0x9ce),
	},
	73: {
		Fstart: int32(0x9d7),
		Fend:   int32(0x9d7),
	},
	74: {
		Fstart: int32(0x9dc),
		Fend:   int32(0x9dd),
	},
	75: {
		Fstart: int32(0x9df),
		Fend:   int32(0x9e3),
	},
	76: {
		Fstart: int32(0x9e6),
		Fend:   int32(0x9f1),
	},
	77: {
		Fstart: int32(0x9f4),
		Fend:   int32(0x9f9),
	},
	78: {
		Fstart: int32(0x9fc),
		Fend:   int32(0x9fc),
	},
	79: {
		Fstart: int32(0x9fe),
		Fend:   int32(0x9fe),
	},
	80: {
		Fstart: int32(0xa01),
		Fend:   int32(0xa03),
	},
	81: {
		Fstart: int32(0xa05),
		Fend:   int32(0xa0a),
	},
	82: {
		Fstart: int32(0xa0f),
		Fend:   int32(0xa10),
	},
	83: {
		Fstart: int32(0xa13),
		Fend:   int32(0xa28),
	},
	84: {
		Fstart: int32(0xa2a),
		Fend:   int32(0xa30),
	},
	85: {
		Fstart: int32(0xa32),
		Fend:   int32(0xa33),
	},
	86: {
		Fstart: int32(0xa35),
		Fend:   int32(0xa36),
	},
	87: {
		Fstart: int32(0xa38),
		Fend:   int32(0xa39),
	},
	88: {
		Fstart: int32(0xa3c),
		Fend:   int32(0xa3c),
	},
	89: {
		Fstart: int32(0xa3e),
		Fend:   int32(0xa42),
	},
	90: {
		Fstart: int32(0xa47),
		Fend:   int32(0xa48),
	},
	91: {
		Fstart: int32(0xa4b),
		Fend:   int32(0xa4d),
	},
	92: {
		Fstart: int32(0xa51),
		Fend:   int32(0xa51),
	},
	93: {
		Fstart: int32(0xa59),
		Fend:   int32(0xa5c),
	},
	94: {
		Fstart: int32(0xa5e),
		Fend:   int32(0xa5e),
	},
	95: {
		Fstart: int32(0xa66),
		Fend:   int32(0xa75),
	},
	96: {
		Fstart: int32(0xa81),
		Fend:   int32(0xa83),
	},
	97: {
		Fstart: int32(0xa85),
		Fend:   int32(0xa8d),
	},
	98: {
		Fstart: int32(0xa8f),
		Fend:   int32(0xa91),
	},
	99: {
		Fstart: int32(0xa93),
		Fend:   int32(0xaa8),
	},
	100: {
		Fstart: int32(0xaaa),
		Fend:   int32(0xab0),
	},
	101: {
		Fstart: int32(0xab2),
		Fend:   int32(0xab3),
	},
	102: {
		Fstart: int32(0xab5),
		Fend:   int32(0xab9),
	},
	103: {
		Fstart: int32(0xabc),
		Fend:   int32(0xac5),
	},
	104: {
		Fstart: int32(0xac7),
		Fend:   int32(0xac9),
	},
	105: {
		Fstart: int32(0xacb),
		Fend:   int32(0xacd),
	},
	106: {
		Fstart: int32(0xad0),
		Fend:   int32(0xad0),
	},
	107: {
		Fstart: int32(0xae0),
		Fend:   int32(0xae3),
	},
	108: {
		Fstart: int32(0xae6),
		Fend:   int32(0xaef),
	},
	109: {
		Fstart: int32(0xaf9),
		Fend:   int32(0xaff),
	},
	110: {
		Fstart: int32(0xb01),
		Fend:   int32(0xb03),
	},
	111: {
		Fstart: int32(0xb05),
		Fend:   int32(0xb0c),
	},
	112: {
		Fstart: int32(0xb0f),
		Fend:   int32(0xb10),
	},
	113: {
		Fstart: int32(0xb13),
		Fend:   int32(0xb28),
	},
	114: {
		Fstart: int32(0xb2a),
		Fend:   int32(0xb30),
	},
	115: {
		Fstart: int32(0xb32),
		Fend:   int32(0xb33),
	},
	116: {
		Fstart: int32(0xb35),
		Fend:   int32(0xb39),
	},
	117: {
		Fstart: int32(0xb3c),
		Fend:   int32(0xb44),
	},
	118: {
		Fstart: int32(0xb47),
		Fend:   int32(0xb48),
	},
	119: {
		Fstart: int32(0xb4b),
		Fend:   int32(0xb4d),
	},
	120: {
		Fstart: int32(0xb55),
		Fend:   int32(0xb57),
	},
	121: {
		Fstart: int32(0xb5c),
		Fend:   int32(0xb5d),
	},
	122: {
		Fstart: int32(0xb5f),
		Fend:   int32(0xb63),
	},
	123: {
		Fstart: int32(0xb66),
		Fend:   int32(0xb6f),
	},
	124: {
		Fstart: int32(0xb71),
		Fend:   int32(0xb77),
	},
	125: {
		Fstart: int32(0xb82),
		Fend:   int32(0xb83),
	},
	126: {
		Fstart: int32(0xb85),
		Fend:   int32(0xb8a),
	},
	127: {
		Fstart: int32(0xb8e),
		Fend:   int32(0xb90),
	},
	128: {
		Fstart: int32(0xb92),
		Fend:   int32(0xb95),
	},
	129: {
		Fstart: int32(0xb99),
		Fend:   int32(0xb9a),
	},
	130: {
		Fstart: int32(0xb9c),
		Fend:   int32(0xb9c),
	},
	131: {
		Fstart: int32(0xb9e),
		Fend:   int32(0xb9f),
	},
	132: {
		Fstart: int32(0xba3),
		Fend:   int32(0xba4),
	},
	133: {
		Fstart: int32(0xba8),
		Fend:   int32(0xbaa),
	},
	134: {
		Fstart: int32(0xbae),
		Fend:   int32(0xbb9),
	},
	135: {
		Fstart: int32(0xbbe),
		Fend:   int32(0xbc2),
	},
	136: {
		Fstart: int32(0xbc6),
		Fend:   int32(0xbc8),
	},
	137: {
		Fstart: int32(0xbca),
		Fend:   int32(0xbcd),
	},
	138: {
		Fstart: int32(0xbd0),
		Fend:   int32(0xbd0),
	},
	139: {
		Fstart: int32(0xbd7),
		Fend:   int32(0xbd7),
	},
	140: {
		Fstart: int32(0xbe6),
		Fend:   int32(0xbf2),
	},
	141: {
		Fstart: int32(0xc00),
		Fend:   int32(0xc0c),
	},
	142: {
		Fstart: int32(0xc0e),
		Fend:   int32(0xc10),
	},
	143: {
		Fstart: int32(0xc12),
		Fend:   int32(0xc28),
	},
	144: {
		Fstart: int32(0xc2a),
		Fend:   int32(0xc39),
	},
	145: {
		Fstart: int32(0xc3c),
		Fend:   int32(0xc44),
	},
	146: {
		Fstart: int32(0xc46),
		Fend:   int32(0xc48),
	},
	147: {
		Fstart: int32(0xc4a),
		Fend:   int32(0xc4d),
	},
	148: {
		Fstart: int32(0xc55),
		Fend:   int32(0xc56),
	},
	149: {
		Fstart: int32(0xc58),
		Fend:   int32(0xc5a),
	},
	150: {
		Fstart: int32(0xc5d),
		Fend:   int32(0xc5d),
	},
	151: {
		Fstart: int32(0xc60),
		Fend:   int32(0xc63),
	},
	152: {
		Fstart: int32(0xc66),
		Fend:   int32(0xc6f),
	},
	153: {
		Fstart: int32(0xc78),
		Fend:   int32(0xc7e),
	},
	154: {
		Fstart: int32(0xc80),
		Fend:   int32(0xc83),
	},
	155: {
		Fstart: int32(0xc85),
		Fend:   int32(0xc8c),
	},
	156: {
		Fstart: int32(0xc8e),
		Fend:   int32(0xc90),
	},
	157: {
		Fstart: int32(0xc92),
		Fend:   int32(0xca8),
	},
	158: {
		Fstart: int32(0xcaa),
		Fend:   int32(0xcb3),
	},
	159: {
		Fstart: int32(0xcb5),
		Fend:   int32(0xcb9),
	},
	160: {
		Fstart: int32(0xcbc),
		Fend:   int32(0xcc4),
	},
	161: {
		Fstart: int32(0xcc6),
		Fend:   int32(0xcc8),
	},
	162: {
		Fstart: int32(0xcca),
		Fend:   int32(0xccd),
	},
	163: {
		Fstart: int32(0xcd5),
		Fend:   int32(0xcd6),
	},
	164: {
		Fstart: int32(0xcdd),
		Fend:   int32(0xcde),
	},
	165: {
		Fstart: int32(0xce0),
		Fend:   int32(0xce3),
	},
	166: {
		Fstart: int32(0xce6),
		Fend:   int32(0xcef),
	},
	167: {
		Fstart: int32(0xcf1),
		Fend:   int32(0xcf3),
	},
	168: {
		Fstart: int32(0xd00),
		Fend:   int32(0xd0c),
	},
	169: {
		Fstart: int32(0xd0e),
		Fend:   int32(0xd10),
	},
	170: {
		Fstart: int32(0xd12),
		Fend:   int32(0xd44),
	},
	171: {
		Fstart: int32(0xd46),
		Fend:   int32(0xd48),
	},
	172: {
		Fstart: int32(0xd4a),
		Fend:   int32(0xd4e),
	},
	173: {
		Fstart: int32(0xd54),
		Fend:   int32(0xd63),
	},
	174: {
		Fstart: int32(0xd66),
		Fend:   int32(0xd78),
	},
	175: {
		Fstart: int32(0xd7a),
		Fend:   int32(0xd7f),
	},
	176: {
		Fstart: int32(0xd81),
		Fend:   int32(0xd83),
	},
	177: {
		Fstart: int32(0xd85),
		Fend:   int32(0xd96),
	},
	178: {
		Fstart: int32(0xd9a),
		Fend:   int32(0xdb1),
	},
	179: {
		Fstart: int32(0xdb3),
		Fend:   int32(0xdbb),
	},
	180: {
		Fstart: int32(0xdbd),
		Fend:   int32(0xdbd),
	},
	181: {
		Fstart: int32(0xdc0),
		Fend:   int32(0xdc6),
	},
	182: {
		Fstart: int32(0xdca),
		Fend:   int32(0xdca),
	},
	183: {
		Fstart: int32(0xdcf),
		Fend:   int32(0xdd4),
	},
	184: {
		Fstart: int32(0xdd6),
		Fend:   int32(0xdd6),
	},
	185: {
		Fstart: int32(0xdd8),
		Fend:   int32(0xddf),
	},
	186: {
		Fstart: int32(0xde6),
		Fend:   int32(0xdef),
	},
	187: {
		Fstart: int32(0xdf2),
		Fend:   int32(0xdf3),
	},
	188: {
		Fstart: int32(0xe01),
		Fend:   int32(0xe3a),
	},
	189: {
		Fstart: int32(0xe40),
		Fend:   int32(0xe4e),
	},
	190: {
		Fstart: int32(0xe50),
		Fend:   int32(0xe59),
	},
	191: {
		Fstart: int32(0xe81),
		Fend:   int32(0xe82),
	},
	192: {
		Fstart: int32(0xe84),
		Fend:   int32(0xe84),
	},
	193: {
		Fstart: int32(0xe86),
		Fend:   int32(0xe8a),
	},
	194: {
		Fstart: int32(0xe8c),
		Fend:   int32(0xea3),
	},
	195: {
		Fstart: int32(0xea5),
		Fend:   int32(0xea5),
	},
	196: {
		Fstart: int32(0xea7),
		Fend:   int32(0xebd),
	},
	197: {
		Fstart: int32(0xec0),
		Fend:   int32(0xec4),
	},
	198: {
		Fstart: int32(0xec6),
		Fend:   int32(0xec6),
	},
	199: {
		Fstart: int32(0xec8),
		Fend:   int32(0xece),
	},
	200: {
		Fstart: int32(0xed0),
		Fend:   int32(0xed9),
	},
	201: {
		Fstart: int32(0xedc),
		Fend:   int32(0xedf),
	},
	202: {
		Fstart: int32(0xf00),
		Fend:   int32(0xf00),
	},
	203: {
		Fstart: int32(0xf18),
		Fend:   int32(0xf19),
	},
	204: {
		Fstart: int32(0xf20),
		Fend:   int32(0xf33),
	},
	205: {
		Fstart: int32(0xf35),
		Fend:   int32(0xf35),
	},
	206: {
		Fstart: int32(0xf37),
		Fend:   int32(0xf37),
	},
	207: {
		Fstart: int32(0xf39),
		Fend:   int32(0xf39),
	},
	208: {
		Fstart: int32(0xf3e),
		Fend:   int32(0xf47),
	},
	209: {
		Fstart: int32(0xf49),
		Fend:   int32(0xf6c),
	},
	210: {
		Fstart: int32(0xf71),
		Fend:   int32(0xf84),
	},
	211: {
		Fstart: int32(0xf86),
		Fend:   int32(0xf97),
	},
	212: {
		Fstart: int32(0xf99),
		Fend:   int32(0xfbc),
	},
	213: {
		Fstart: int32(0xfc6),
		Fend:   int32(0xfc6),
	},
	214: {
		Fstart: int32(0x1000),
		Fend:   int32(0x1049),
	},
	215: {
		Fstart: int32(0x1050),
		Fend:   int32(0x109d),
	},
	216: {
		Fstart: int32(0x10a0),
		Fend:   int32(0x10c5),
	},
	217: {
		Fstart: int32(0x10c7),
		Fend:   int32(0x10c7),
	},
	218: {
		Fstart: int32(0x10cd),
		Fend:   int32(0x10cd),
	},
	219: {
		Fstart: int32(0x10d0),
		Fend:   int32(0x10fa),
	},
	220: {
		Fstart: int32(0x10fc),
		Fend:   int32(0x1248),
	},
	221: {
		Fstart: int32(0x124a),
		Fend:   int32(0x124d),
	},
	222: {
		Fstart: int32(0x1250),
		Fend:   int32(0x1256),
	},
	223: {
		Fstart: int32(0x1258),
		Fend:   int32(0x1258),
	},
	224: {
		Fstart: int32(0x125a),
		Fend:   int32(0x125d),
	},
	225: {
		Fstart: int32(0x1260),
		Fend:   int32(0x1288),
	},
	226: {
		Fstart: int32(0x128a),
		Fend:   int32(0x128d),
	},
	227: {
		Fstart: int32(0x1290),
		Fend:   int32(0x12b0),
	},
	228: {
		Fstart: int32(0x12b2),
		Fend:   int32(0x12b5),
	},
	229: {
		Fstart: int32(0x12b8),
		Fend:   int32(0x12be),
	},
	230: {
		Fstart: int32(0x12c0),
		Fend:   int32(0x12c0),
	},
	231: {
		Fstart: int32(0x12c2),
		Fend:   int32(0x12c5),
	},
	232: {
		Fstart: int32(0x12c8),
		Fend:   int32(0x12d6),
	},
	233: {
		Fstart: int32(0x12d8),
		Fend:   int32(0x1310),
	},
	234: {
		Fstart: int32(0x1312),
		Fend:   int32(0x1315),
	},
	235: {
		Fstart: int32(0x1318),
		Fend:   int32(0x135a),
	},
	236: {
		Fstart: int32(0x135d),
		Fend:   int32(0x135f),
	},
	237: {
		Fstart: int32(0x1369),
		Fend:   int32(0x137c),
	},
	238: {
		Fstart: int32(0x1380),
		Fend:   int32(0x138f),
	},
	239: {
		Fstart: int32(0x13a0),
		Fend:   int32(0x13f5),
	},
	240: {
		Fstart: int32(0x13f8),
		Fend:   int32(0x13fd),
	},
	241: {
		Fstart: int32(0x1401),
		Fend:   int32(0x166c),
	},
	242: {
		Fstart: int32(0x166f),
		Fend:   int32(0x167f),
	},
	243: {
		Fstart: int32(0x1681),
		Fend:   int32(0x169a),
	},
	244: {
		Fstart: int32(0x16a0),
		Fend:   int32(0x16ea),
	},
	245: {
		Fstart: int32(0x16ee),
		Fend:   int32(0x16f8),
	},
	246: {
		Fstart: int32(0x1700),
		Fend:   int32(0x1715),
	},
	247: {
		Fstart: int32(0x171f),
		Fend:   int32(0x1734),
	},
	248: {
		Fstart: int32(0x1740),
		Fend:   int32(0x1753),
	},
	249: {
		Fstart: int32(0x1760),
		Fend:   int32(0x176c),
	},
	250: {
		Fstart: int32(0x176e),
		Fend:   int32(0x1770),
	},
	251: {
		Fstart: int32(0x1772),
		Fend:   int32(0x1773),
	},
	252: {
		Fstart: int32(0x1780),
		Fend:   int32(0x17d3),
	},
	253: {
		Fstart: int32(0x17d7),
		Fend:   int32(0x17d7),
	},
	254: {
		Fstart: int32(0x17dc),
		Fend:   int32(0x17dd),
	},
	255: {
		Fstart: int32(0x17e0),
		Fend:   int32(0x17e9),
	},
	256: {
		Fstart: int32(0x17f0),
		Fend:   int32(0x17f9),
	},
	257: {
		Fstart: int32(0x180b),
		Fend:   int32(0x180d),
	},
	258: {
		Fstart: int32(0x180f),
		Fend:   int32(0x1819),
	},
	259: {
		Fstart: int32(0x1820),
		Fend:   int32(0x1878),
	},
	260: {
		Fstart: int32(0x1880),
		Fend:   int32(0x18aa),
	},
	261: {
		Fstart: int32(0x18b0),
		Fend:   int32(0x18f5),
	},
	262: {
		Fstart: int32(0x1900),
		Fend:   int32(0x191e),
	},
	263: {
		Fstart: int32(0x1920),
		Fend:   int32(0x192b),
	},
	264: {
		Fstart: int32(0x1930),
		Fend:   int32(0x193b),
	},
	265: {
		Fstart: int32(0x1946),
		Fend:   int32(0x196d),
	},
	266: {
		Fstart: int32(0x1970),
		Fend:   int32(0x1974),
	},
	267: {
		Fstart: int32(0x1980),
		Fend:   int32(0x19ab),
	},
	268: {
		Fstart: int32(0x19b0),
		Fend:   int32(0x19c9),
	},
	269: {
		Fstart: int32(0x19d0),
		Fend:   int32(0x19da),
	},
	270: {
		Fstart: int32(0x1a00),
		Fend:   int32(0x1a1b),
	},
	271: {
		Fstart: int32(0x1a20),
		Fend:   int32(0x1a5e),
	},
	272: {
		Fstart: int32(0x1a60),
		Fend:   int32(0x1a7c),
	},
	273: {
		Fstart: int32(0x1a7f),
		Fend:   int32(0x1a89),
	},
	274: {
		Fstart: int32(0x1a90),
		Fend:   int32(0x1a99),
	},
	275: {
		Fstart: int32(0x1aa7),
		Fend:   int32(0x1aa7),
	},
	276: {
		Fstart: int32(0x1ab0),
		Fend:   int32(0x1ace),
	},
	277: {
		Fstart: int32(0x1b00),
		Fend:   int32(0x1b4c),
	},
	278: {
		Fstart: int32(0x1b50),
		Fend:   int32(0x1b59),
	},
	279: {
		Fstart: int32(0x1b6b),
		Fend:   int32(0x1b73),
	},
	280: {
		Fstart: int32(0x1b80),
		Fend:   int32(0x1bf3),
	},
	281: {
		Fstart: int32(0x1c00),
		Fend:   int32(0x1c37),
	},
	282: {
		Fstart: int32(0x1c40),
		Fend:   int32(0x1c49),
	},
	283: {
		Fstart: int32(0x1c4d),
		Fend:   int32(0x1c7d),
	},
	284: {
		Fstart: int32(0x1c80),
		Fend:   int32(0x1c88),
	},
	285: {
		Fstart: int32(0x1c90),
		Fend:   int32(0x1cba),
	},
	286: {
		Fstart: int32(0x1cbd),
		Fend:   int32(0x1cbf),
	},
	287: {
		Fstart: int32(0x1cd0),
		Fend:   int32(0x1cd2),
	},
	288: {
		Fstart: int32(0x1cd4),
		Fend:   int32(0x1cfa),
	},
	289: {
		Fstart: int32(0x1d00),
		Fend:   int32(0x1f15),
	},
	290: {
		Fstart: int32(0x1f18),
		Fend:   int32(0x1f1d),
	},
	291: {
		Fstart: int32(0x1f20),
		Fend:   int32(0x1f45),
	},
	292: {
		Fstart: int32(0x1f48),
		Fend:   int32(0x1f4d),
	},
	293: {
		Fstart: int32(0x1f50),
		Fend:   int32(0x1f57),
	},
	294: {
		Fstart: int32(0x1f59),
		Fend:   int32(0x1f59),
	},
	295: {
		Fstart: int32(0x1f5b),
		Fend:   int32(0x1f5b),
	},
	296: {
		Fstart: int32(0x1f5d),
		Fend:   int32(0x1f5d),
	},
	297: {
		Fstart: int32(0x1f5f),
		Fend:   int32(0x1f7d),
	},
	298: {
		Fstart: int32(0x1f80),
		Fend:   int32(0x1fb4),
	},
	299: {
		Fstart: int32(0x1fb6),
		Fend:   int32(0x1fbc),
	},
	300: {
		Fstart: int32(0x1fbe),
		Fend:   int32(0x1fbe),
	},
	301: {
		Fstart: int32(0x1fc2),
		Fend:   int32(0x1fc4),
	},
	302: {
		Fstart: int32(0x1fc6),
		Fend:   int32(0x1fcc),
	},
	303: {
		Fstart: int32(0x1fd0),
		Fend:   int32(0x1fd3),
	},
	304: {
		Fstart: int32(0x1fd6),
		Fend:   int32(0x1fdb),
	},
	305: {
		Fstart: int32(0x1fe0),
		Fend:   int32(0x1fec),
	},
	306: {
		Fstart: int32(0x1ff2),
		Fend:   int32(0x1ff4),
	},
	307: {
		Fstart: int32(0x1ff6),
		Fend:   int32(0x1ffc),
	},
	308: {
		Fstart: int32(0x2070),
		Fend:   int32(0x2071),
	},
	309: {
		Fstart: int32(0x2074),
		Fend:   int32(0x2079),
	},
	310: {
		Fstart: int32(0x207f),
		Fend:   int32(0x2089),
	},
	311: {
		Fstart: int32(0x2090),
		Fend:   int32(0x209c),
	},
	312: {
		Fstart: int32(0x20d0),
		Fend:   int32(0x20f0),
	},
	313: {
		Fstart: int32(0x2102),
		Fend:   int32(0x2102),
	},
	314: {
		Fstart: int32(0x2107),
		Fend:   int32(0x2107),
	},
	315: {
		Fstart: int32(0x210a),
		Fend:   int32(0x2113),
	},
	316: {
		Fstart: int32(0x2115),
		Fend:   int32(0x2115),
	},
	317: {
		Fstart: int32(0x2119),
		Fend:   int32(0x211d),
	},
	318: {
		Fstart: int32(0x2124),
		Fend:   int32(0x2124),
	},
	319: {
		Fstart: int32(0x2126),
		Fend:   int32(0x2126),
	},
	320: {
		Fstart: int32(0x2128),
		Fend:   int32(0x2128),
	},
	321: {
		Fstart: int32(0x212a),
		Fend:   int32(0x212d),
	},
	322: {
		Fstart: int32(0x212f),
		Fend:   int32(0x2139),
	},
	323: {
		Fstart: int32(0x213c),
		Fend:   int32(0x213f),
	},
	324: {
		Fstart: int32(0x2145),
		Fend:   int32(0x2149),
	},
	325: {
		Fstart: int32(0x214e),
		Fend:   int32(0x214e),
	},
	326: {
		Fstart: int32(0x2150),
		Fend:   int32(0x2189),
	},
	327: {
		Fstart: int32(0x2460),
		Fend:   int32(0x249b),
	},
	328: {
		Fstart: int32(0x24ea),
		Fend:   int32(0x24ff),
	},
	329: {
		Fstart: int32(0x2776),
		Fend:   int32(0x2793),
	},
	330: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2ce4),
	},
	331: {
		Fstart: int32(0x2ceb),
		Fend:   int32(0x2cf3),
	},
	332: {
		Fstart: int32(0x2cfd),
		Fend:   int32(0x2cfd),
	},
	333: {
		Fstart: int32(0x2d00),
		Fend:   int32(0x2d25),
	},
	334: {
		Fstart: int32(0x2d27),
		Fend:   int32(0x2d27),
	},
	335: {
		Fstart: int32(0x2d2d),
		Fend:   int32(0x2d2d),
	},
	336: {
		Fstart: int32(0x2d30),
		Fend:   int32(0x2d67),
	},
	337: {
		Fstart: int32(0x2d6f),
		Fend:   int32(0x2d6f),
	},
	338: {
		Fstart: int32(0x2d7f),
		Fend:   int32(0x2d96),
	},
	339: {
		Fstart: int32(0x2da0),
		Fend:   int32(0x2da6),
	},
	340: {
		Fstart: int32(0x2da8),
		Fend:   int32(0x2dae),
	},
	341: {
		Fstart: int32(0x2db0),
		Fend:   int32(0x2db6),
	},
	342: {
		Fstart: int32(0x2db8),
		Fend:   int32(0x2dbe),
	},
	343: {
		Fstart: int32(0x2dc0),
		Fend:   int32(0x2dc6),
	},
	344: {
		Fstart: int32(0x2dc8),
		Fend:   int32(0x2dce),
	},
	345: {
		Fstart: int32(0x2dd0),
		Fend:   int32(0x2dd6),
	},
	346: {
		Fstart: int32(0x2dd8),
		Fend:   int32(0x2dde),
	},
	347: {
		Fstart: int32(0x2de0),
		Fend:   int32(0x2dff),
	},
	348: {
		Fstart: int32(0x2e2f),
		Fend:   int32(0x2e2f),
	},
	349: {
		Fstart: int32(0x3005),
		Fend:   int32(0x3007),
	},
	350: {
		Fstart: int32(0x3021),
		Fend:   int32(0x302f),
	},
	351: {
		Fstart: int32(0x3031),
		Fend:   int32(0x3035),
	},
	352: {
		Fstart: int32(0x3038),
		Fend:   int32(0x303c),
	},
	353: {
		Fstart: int32(0x3041),
		Fend:   int32(0x3096),
	},
	354: {
		Fstart: int32(0x3099),
		Fend:   int32(0x309a),
	},
	355: {
		Fstart: int32(0x309d),
		Fend:   int32(0x309f),
	},
	356: {
		Fstart: int32(0x30a1),
		Fend:   int32(0x30fa),
	},
	357: {
		Fstart: int32(0x30fc),
		Fend:   int32(0x30ff),
	},
	358: {
		Fstart: int32(0x3105),
		Fend:   int32(0x312f),
	},
	359: {
		Fstart: int32(0x3131),
		Fend:   int32(0x318e),
	},
	360: {
		Fstart: int32(0x3192),
		Fend:   int32(0x3195),
	},
	361: {
		Fstart: int32(0x31a0),
		Fend:   int32(0x31bf),
	},
	362: {
		Fstart: int32(0x31f0),
		Fend:   int32(0x31ff),
	},
	363: {
		Fstart: int32(0x3220),
		Fend:   int32(0x3229),
	},
	364: {
		Fstart: int32(0x3248),
		Fend:   int32(0x324f),
	},
	365: {
		Fstart: int32(0x3251),
		Fend:   int32(0x325f),
	},
	366: {
		Fstart: int32(0x3280),
		Fend:   int32(0x3289),
	},
	367: {
		Fstart: int32(0x32b1),
		Fend:   int32(0x32bf),
	},
	368: {
		Fstart: int32(0x3400),
		Fend:   int32(0x3400),
	},
	369: {
		Fstart: int32(0x4dbf),
		Fend:   int32(0x4dbf),
	},
	370: {
		Fstart: int32(0x4e00),
		Fend:   int32(0xa48c),
	},
	371: {
		Fstart: int32(0xa4d0),
		Fend:   int32(0xa4fd),
	},
	372: {
		Fstart: int32(0xa500),
		Fend:   int32(0xa60c),
	},
	373: {
		Fstart: int32(0xa610),
		Fend:   int32(0xa62b),
	},
	374: {
		Fstart: int32(0xa640),
		Fend:   int32(0xa672),
	},
	375: {
		Fstart: int32(0xa674),
		Fend:   int32(0xa67d),
	},
	376: {
		Fstart: int32(0xa67f),
		Fend:   int32(0xa6f1),
	},
	377: {
		Fstart: int32(0xa717),
		Fend:   int32(0xa71f),
	},
	378: {
		Fstart: int32(0xa722),
		Fend:   int32(0xa788),
	},
	379: {
		Fstart: int32(0xa78b),
		Fend:   int32(0xa7ca),
	},
	380: {
		Fstart: int32(0xa7d0),
		Fend:   int32(0xa7d1),
	},
	381: {
		Fstart: int32(0xa7d3),
		Fend:   int32(0xa7d3),
	},
	382: {
		Fstart: int32(0xa7d5),
		Fend:   int32(0xa7d9),
	},
	383: {
		Fstart: int32(0xa7f2),
		Fend:   int32(0xa827),
	},
	384: {
		Fstart: int32(0xa82c),
		Fend:   int32(0xa82c),
	},
	385: {
		Fstart: int32(0xa830),
		Fend:   int32(0xa835),
	},
	386: {
		Fstart: int32(0xa840),
		Fend:   int32(0xa873),
	},
	387: {
		Fstart: int32(0xa880),
		Fend:   int32(0xa8c5),
	},
	388: {
		Fstart: int32(0xa8d0),
		Fend:   int32(0xa8d9),
	},
	389: {
		Fstart: int32(0xa8e0),
		Fend:   int32(0xa8f7),
	},
	390: {
		Fstart: int32(0xa8fb),
		Fend:   int32(0xa8fb),
	},
	391: {
		Fstart: int32(0xa8fd),
		Fend:   int32(0xa92d),
	},
	392: {
		Fstart: int32(0xa930),
		Fend:   int32(0xa953),
	},
	393: {
		Fstart: int32(0xa960),
		Fend:   int32(0xa97c),
	},
	394: {
		Fstart: int32(0xa980),
		Fend:   int32(0xa9c0),
	},
	395: {
		Fstart: int32(0xa9cf),
		Fend:   int32(0xa9d9),
	},
	396: {
		Fstart: int32(0xa9e0),
		Fend:   int32(0xa9fe),
	},
	397: {
		Fstart: int32(0xaa00),
		Fend:   int32(0xaa36),
	},
	398: {
		Fstart: int32(0xaa40),
		Fend:   int32(0xaa4d),
	},
	399: {
		Fstart: int32(0xaa50),
		Fend:   int32(0xaa59),
	},
	400: {
		Fstart: int32(0xaa60),
		Fend:   int32(0xaa76),
	},
	401: {
		Fstart: int32(0xaa7a),
		Fend:   int32(0xaac2),
	},
	402: {
		Fstart: int32(0xaadb),
		Fend:   int32(0xaadd),
	},
	403: {
		Fstart: int32(0xaae0),
		Fend:   int32(0xaaef),
	},
	404: {
		Fstart: int32(0xaaf2),
		Fend:   int32(0xaaf6),
	},
	405: {
		Fstart: int32(0xab01),
		Fend:   int32(0xab06),
	},
	406: {
		Fstart: int32(0xab09),
		Fend:   int32(0xab0e),
	},
	407: {
		Fstart: int32(0xab11),
		Fend:   int32(0xab16),
	},
	408: {
		Fstart: int32(0xab20),
		Fend:   int32(0xab26),
	},
	409: {
		Fstart: int32(0xab28),
		Fend:   int32(0xab2e),
	},
	410: {
		Fstart: int32(0xab30),
		Fend:   int32(0xab5a),
	},
	411: {
		Fstart: int32(0xab5c),
		Fend:   int32(0xab69),
	},
	412: {
		Fstart: int32(0xab70),
		Fend:   int32(0xabea),
	},
	413: {
		Fstart: int32(0xabec),
		Fend:   int32(0xabed),
	},
	414: {
		Fstart: int32(0xabf0),
		Fend:   int32(0xabf9),
	},
	415: {
		Fstart: int32(0xac00),
		Fend:   int32(0xac00),
	},
	416: {
		Fstart: int32(0xd7a3),
		Fend:   int32(0xd7a3),
	},
	417: {
		Fstart: int32(0xd7b0),
		Fend:   int32(0xd7c6),
	},
	418: {
		Fstart: int32(0xd7cb),
		Fend:   int32(0xd7fb),
	},
	419: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfa6d),
	},
	420: {
		Fstart: int32(0xfa70),
		Fend:   int32(0xfad9),
	},
	421: {
		Fstart: int32(0xfb00),
		Fend:   int32(0xfb06),
	},
	422: {
		Fstart: int32(0xfb13),
		Fend:   int32(0xfb17),
	},
	423: {
		Fstart: int32(0xfb1d),
		Fend:   int32(0xfb28),
	},
	424: {
		Fstart: int32(0xfb2a),
		Fend:   int32(0xfb36),
	},
	425: {
		Fstart: int32(0xfb38),
		Fend:   int32(0xfb3c),
	},
	426: {
		Fstart: int32(0xfb3e),
		Fend:   int32(0xfb3e),
	},
	427: {
		Fstart: int32(0xfb40),
		Fend:   int32(0xfb41),
	},
	428: {
		Fstart: int32(0xfb43),
		Fend:   int32(0xfb44),
	},
	429: {
		Fstart: int32(0xfb46),
		Fend:   int32(0xfbb1),
	},
	430: {
		Fstart: int32(0xfbd3),
		Fend:   int32(0xfd3d),
	},
	431: {
		Fstart: int32(0xfd50),
		Fend:   int32(0xfd8f),
	},
	432: {
		Fstart: int32(0xfd92),
		Fend:   int32(0xfdc7),
	},
	433: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfdfb),
	},
	434: {
		Fstart: int32(0xfe00),
		Fend:   int32(0xfe0f),
	},
	435: {
		Fstart: int32(0xfe20),
		Fend:   int32(0xfe2f),
	},
	436: {
		Fstart: int32(0xfe70),
		Fend:   int32(0xfe74),
	},
	437: {
		Fstart: int32(0xfe76),
		Fend:   int32(0xfefc),
	},
	438: {
		Fstart: int32(0xff10),
		Fend:   int32(0xff19),
	},
	439: {
		Fstart: int32(0xff21),
		Fend:   int32(0xff3a),
	},
	440: {
		Fstart: int32(0xff41),
		Fend:   int32(0xff5a),
	},
	441: {
		Fstart: int32(0xff66),
		Fend:   int32(0xffbe),
	},
	442: {
		Fstart: int32(0xffc2),
		Fend:   int32(0xffc7),
	},
	443: {
		Fstart: int32(0xffca),
		Fend:   int32(0xffcf),
	},
	444: {
		Fstart: int32(0xffd2),
		Fend:   int32(0xffd7),
	},
	445: {
		Fstart: int32(0xffda),
		Fend:   int32(0xffdc),
	},
	446: {
		Fstart: int32(0x10000),
		Fend:   int32(0x1000b),
	},
	447: {
		Fstart: int32(0x1000d),
		Fend:   int32(0x10026),
	},
	448: {
		Fstart: int32(0x10028),
		Fend:   int32(0x1003a),
	},
	449: {
		Fstart: int32(0x1003c),
		Fend:   int32(0x1003d),
	},
	450: {
		Fstart: int32(0x1003f),
		Fend:   int32(0x1004d),
	},
	451: {
		Fstart: int32(0x10050),
		Fend:   int32(0x1005d),
	},
	452: {
		Fstart: int32(0x10080),
		Fend:   int32(0x100fa),
	},
	453: {
		Fstart: int32(0x10107),
		Fend:   int32(0x10133),
	},
	454: {
		Fstart: int32(0x10140),
		Fend:   int32(0x10178),
	},
	455: {
		Fstart: int32(0x1018a),
		Fend:   int32(0x1018b),
	},
	456: {
		Fstart: int32(0x101fd),
		Fend:   int32(0x101fd),
	},
	457: {
		Fstart: int32(0x10280),
		Fend:   int32(0x1029c),
	},
	458: {
		Fstart: int32(0x102a0),
		Fend:   int32(0x102d0),
	},
	459: {
		Fstart: int32(0x102e0),
		Fend:   int32(0x102fb),
	},
	460: {
		Fstart: int32(0x10300),
		Fend:   int32(0x10323),
	},
	461: {
		Fstart: int32(0x1032d),
		Fend:   int32(0x1034a),
	},
	462: {
		Fstart: int32(0x10350),
		Fend:   int32(0x1037a),
	},
	463: {
		Fstart: int32(0x10380),
		Fend:   int32(0x1039d),
	},
	464: {
		Fstart: int32(0x103a0),
		Fend:   int32(0x103c3),
	},
	465: {
		Fstart: int32(0x103c8),
		Fend:   int32(0x103cf),
	},
	466: {
		Fstart: int32(0x103d1),
		Fend:   int32(0x103d5),
	},
	467: {
		Fstart: int32(0x10400),
		Fend:   int32(0x1049d),
	},
	468: {
		Fstart: int32(0x104a0),
		Fend:   int32(0x104a9),
	},
	469: {
		Fstart: int32(0x104b0),
		Fend:   int32(0x104d3),
	},
	470: {
		Fstart: int32(0x104d8),
		Fend:   int32(0x104fb),
	},
	471: {
		Fstart: int32(0x10500),
		Fend:   int32(0x10527),
	},
	472: {
		Fstart: int32(0x10530),
		Fend:   int32(0x10563),
	},
	473: {
		Fstart: int32(0x10570),
		Fend:   int32(0x1057a),
	},
	474: {
		Fstart: int32(0x1057c),
		Fend:   int32(0x1058a),
	},
	475: {
		Fstart: int32(0x1058c),
		Fend:   int32(0x10592),
	},
	476: {
		Fstart: int32(0x10594),
		Fend:   int32(0x10595),
	},
	477: {
		Fstart: int32(0x10597),
		Fend:   int32(0x105a1),
	},
	478: {
		Fstart: int32(0x105a3),
		Fend:   int32(0x105b1),
	},
	479: {
		Fstart: int32(0x105b3),
		Fend:   int32(0x105b9),
	},
	480: {
		Fstart: int32(0x105bb),
		Fend:   int32(0x105bc),
	},
	481: {
		Fstart: int32(0x10600),
		Fend:   int32(0x10736),
	},
	482: {
		Fstart: int32(0x10740),
		Fend:   int32(0x10755),
	},
	483: {
		Fstart: int32(0x10760),
		Fend:   int32(0x10767),
	},
	484: {
		Fstart: int32(0x10780),
		Fend:   int32(0x10785),
	},
	485: {
		Fstart: int32(0x10787),
		Fend:   int32(0x107b0),
	},
	486: {
		Fstart: int32(0x107b2),
		Fend:   int32(0x107ba),
	},
	487: {
		Fstart: int32(0x10800),
		Fend:   int32(0x10805),
	},
	488: {
		Fstart: int32(0x10808),
		Fend:   int32(0x10808),
	},
	489: {
		Fstart: int32(0x1080a),
		Fend:   int32(0x10835),
	},
	490: {
		Fstart: int32(0x10837),
		Fend:   int32(0x10838),
	},
	491: {
		Fstart: int32(0x1083c),
		Fend:   int32(0x1083c),
	},
	492: {
		Fstart: int32(0x1083f),
		Fend:   int32(0x10855),
	},
	493: {
		Fstart: int32(0x10858),
		Fend:   int32(0x10876),
	},
	494: {
		Fstart: int32(0x10879),
		Fend:   int32(0x1089e),
	},
	495: {
		Fstart: int32(0x108a7),
		Fend:   int32(0x108af),
	},
	496: {
		Fstart: int32(0x108e0),
		Fend:   int32(0x108f2),
	},
	497: {
		Fstart: int32(0x108f4),
		Fend:   int32(0x108f5),
	},
	498: {
		Fstart: int32(0x108fb),
		Fend:   int32(0x1091b),
	},
}

var sym___identifier_char_no_digit_sign_character_set_1 = [499]TSCharacterRange{
	0: {
		Fstart: int32('!'),
		Fend:   int32('!'),
	},
	1: {
		Fstart: int32('#'),
		Fend:   int32('\''),
	},
	2: {
		Fstart: int32('*'),
		Fend:   int32('*'),
	},
	3: {
		Fstart: int32('.'),
		Fend:   int32('.'),
	},
	4: {
		Fstart: int32(':'),
		Fend:   int32(':'),
	},
	5: {
		Fstart: int32('?'),
		Fend:   int32('Z'),
	},
	6: {
		Fstart: int32('^'),
		Fend:   int32('_'),
	},
	7: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	8: {
		Fstart: int32('|'),
		Fend:   int32('|'),
	},
	9: {
		Fstart: int32('~'),
		Fend:   int32('~'),
	},
	10: {
		Fstart: int32(0xaa),
		Fend:   int32(0xaa),
	},
	11: {
		Fstart: int32(0xb2),
		Fend:   int32(0xb3),
	},
	12: {
		Fstart: int32(0xb5),
		Fend:   int32(0xb5),
	},
	13: {
		Fstart: int32(0xb9),
		Fend:   int32(0xba),
	},
	14: {
		Fstart: int32(0xbc),
		Fend:   int32(0xbe),
	},
	15: {
		Fstart: int32(0xc0),
		Fend:   int32(0xd6),
	},
	16: {
		Fstart: int32(0xd8),
		Fend:   int32(0xf6),
	},
	17: {
		Fstart: int32(0xf8),
		Fend:   int32(0x2c1),
	},
	18: {
		Fstart: int32(0x2c6),
		Fend:   int32(0x2d1),
	},
	19: {
		Fstart: int32(0x2e0),
		Fend:   int32(0x2e4),
	},
	20: {
		Fstart: int32(0x2ec),
		Fend:   int32(0x2ec),
	},
	21: {
		Fstart: int32(0x2ee),
		Fend:   int32(0x2ee),
	},
	22: {
		Fstart: int32(0x300),
		Fend:   int32(0x374),
	},
	23: {
		Fstart: int32(0x376),
		Fend:   int32(0x377),
	},
	24: {
		Fstart: int32(0x37a),
		Fend:   int32(0x37d),
	},
	25: {
		Fstart: int32(0x37f),
		Fend:   int32(0x37f),
	},
	26: {
		Fstart: int32(0x386),
		Fend:   int32(0x386),
	},
	27: {
		Fstart: int32(0x388),
		Fend:   int32(0x38a),
	},
	28: {
		Fstart: int32(0x38c),
		Fend:   int32(0x38c),
	},
	29: {
		Fstart: int32(0x38e),
		Fend:   int32(0x3a1),
	},
	30: {
		Fstart: int32(0x3a3),
		Fend:   int32(0x3f5),
	},
	31: {
		Fstart: int32(0x3f7),
		Fend:   int32(0x481),
	},
	32: {
		Fstart: int32(0x483),
		Fend:   int32(0x52f),
	},
	33: {
		Fstart: int32(0x531),
		Fend:   int32(0x556),
	},
	34: {
		Fstart: int32(0x559),
		Fend:   int32(0x559),
	},
	35: {
		Fstart: int32(0x560),
		Fend:   int32(0x588),
	},
	36: {
		Fstart: int32(0x591),
		Fend:   int32(0x5bd),
	},
	37: {
		Fstart: int32(0x5bf),
		Fend:   int32(0x5bf),
	},
	38: {
		Fstart: int32(0x5c1),
		Fend:   int32(0x5c2),
	},
	39: {
		Fstart: int32(0x5c4),
		Fend:   int32(0x5c5),
	},
	40: {
		Fstart: int32(0x5c7),
		Fend:   int32(0x5c7),
	},
	41: {
		Fstart: int32(0x5d0),
		Fend:   int32(0x5ea),
	},
	42: {
		Fstart: int32(0x5ef),
		Fend:   int32(0x5f2),
	},
	43: {
		Fstart: int32(0x610),
		Fend:   int32(0x61a),
	},
	44: {
		Fstart: int32(0x620),
		Fend:   int32(0x669),
	},
	45: {
		Fstart: int32(0x66e),
		Fend:   int32(0x6d3),
	},
	46: {
		Fstart: int32(0x6d5),
		Fend:   int32(0x6dc),
	},
	47: {
		Fstart: int32(0x6df),
		Fend:   int32(0x6e8),
	},
	48: {
		Fstart: int32(0x6ea),
		Fend:   int32(0x6fc),
	},
	49: {
		Fstart: int32(0x6ff),
		Fend:   int32(0x6ff),
	},
	50: {
		Fstart: int32(0x710),
		Fend:   int32(0x74a),
	},
	51: {
		Fstart: int32(0x74d),
		Fend:   int32(0x7b1),
	},
	52: {
		Fstart: int32(0x7c0),
		Fend:   int32(0x7f5),
	},
	53: {
		Fstart: int32(0x7fa),
		Fend:   int32(0x7fa),
	},
	54: {
		Fstart: int32(0x7fd),
		Fend:   int32(0x7fd),
	},
	55: {
		Fstart: int32(0x800),
		Fend:   int32(0x82d),
	},
	56: {
		Fstart: int32(0x840),
		Fend:   int32(0x85b),
	},
	57: {
		Fstart: int32(0x860),
		Fend:   int32(0x86a),
	},
	58: {
		Fstart: int32(0x870),
		Fend:   int32(0x887),
	},
	59: {
		Fstart: int32(0x889),
		Fend:   int32(0x88e),
	},
	60: {
		Fstart: int32(0x898),
		Fend:   int32(0x8e1),
	},
	61: {
		Fstart: int32(0x8e3),
		Fend:   int32(0x963),
	},
	62: {
		Fstart: int32(0x966),
		Fend:   int32(0x96f),
	},
	63: {
		Fstart: int32(0x971),
		Fend:   int32(0x983),
	},
	64: {
		Fstart: int32(0x985),
		Fend:   int32(0x98c),
	},
	65: {
		Fstart: int32(0x98f),
		Fend:   int32(0x990),
	},
	66: {
		Fstart: int32(0x993),
		Fend:   int32(0x9a8),
	},
	67: {
		Fstart: int32(0x9aa),
		Fend:   int32(0x9b0),
	},
	68: {
		Fstart: int32(0x9b2),
		Fend:   int32(0x9b2),
	},
	69: {
		Fstart: int32(0x9b6),
		Fend:   int32(0x9b9),
	},
	70: {
		Fstart: int32(0x9bc),
		Fend:   int32(0x9c4),
	},
	71: {
		Fstart: int32(0x9c7),
		Fend:   int32(0x9c8),
	},
	72: {
		Fstart: int32(0x9cb),
		Fend:   int32(0x9ce),
	},
	73: {
		Fstart: int32(0x9d7),
		Fend:   int32(0x9d7),
	},
	74: {
		Fstart: int32(0x9dc),
		Fend:   int32(0x9dd),
	},
	75: {
		Fstart: int32(0x9df),
		Fend:   int32(0x9e3),
	},
	76: {
		Fstart: int32(0x9e6),
		Fend:   int32(0x9f1),
	},
	77: {
		Fstart: int32(0x9f4),
		Fend:   int32(0x9f9),
	},
	78: {
		Fstart: int32(0x9fc),
		Fend:   int32(0x9fc),
	},
	79: {
		Fstart: int32(0x9fe),
		Fend:   int32(0x9fe),
	},
	80: {
		Fstart: int32(0xa01),
		Fend:   int32(0xa03),
	},
	81: {
		Fstart: int32(0xa05),
		Fend:   int32(0xa0a),
	},
	82: {
		Fstart: int32(0xa0f),
		Fend:   int32(0xa10),
	},
	83: {
		Fstart: int32(0xa13),
		Fend:   int32(0xa28),
	},
	84: {
		Fstart: int32(0xa2a),
		Fend:   int32(0xa30),
	},
	85: {
		Fstart: int32(0xa32),
		Fend:   int32(0xa33),
	},
	86: {
		Fstart: int32(0xa35),
		Fend:   int32(0xa36),
	},
	87: {
		Fstart: int32(0xa38),
		Fend:   int32(0xa39),
	},
	88: {
		Fstart: int32(0xa3c),
		Fend:   int32(0xa3c),
	},
	89: {
		Fstart: int32(0xa3e),
		Fend:   int32(0xa42),
	},
	90: {
		Fstart: int32(0xa47),
		Fend:   int32(0xa48),
	},
	91: {
		Fstart: int32(0xa4b),
		Fend:   int32(0xa4d),
	},
	92: {
		Fstart: int32(0xa51),
		Fend:   int32(0xa51),
	},
	93: {
		Fstart: int32(0xa59),
		Fend:   int32(0xa5c),
	},
	94: {
		Fstart: int32(0xa5e),
		Fend:   int32(0xa5e),
	},
	95: {
		Fstart: int32(0xa66),
		Fend:   int32(0xa75),
	},
	96: {
		Fstart: int32(0xa81),
		Fend:   int32(0xa83),
	},
	97: {
		Fstart: int32(0xa85),
		Fend:   int32(0xa8d),
	},
	98: {
		Fstart: int32(0xa8f),
		Fend:   int32(0xa91),
	},
	99: {
		Fstart: int32(0xa93),
		Fend:   int32(0xaa8),
	},
	100: {
		Fstart: int32(0xaaa),
		Fend:   int32(0xab0),
	},
	101: {
		Fstart: int32(0xab2),
		Fend:   int32(0xab3),
	},
	102: {
		Fstart: int32(0xab5),
		Fend:   int32(0xab9),
	},
	103: {
		Fstart: int32(0xabc),
		Fend:   int32(0xac5),
	},
	104: {
		Fstart: int32(0xac7),
		Fend:   int32(0xac9),
	},
	105: {
		Fstart: int32(0xacb),
		Fend:   int32(0xacd),
	},
	106: {
		Fstart: int32(0xad0),
		Fend:   int32(0xad0),
	},
	107: {
		Fstart: int32(0xae0),
		Fend:   int32(0xae3),
	},
	108: {
		Fstart: int32(0xae6),
		Fend:   int32(0xaef),
	},
	109: {
		Fstart: int32(0xaf9),
		Fend:   int32(0xaff),
	},
	110: {
		Fstart: int32(0xb01),
		Fend:   int32(0xb03),
	},
	111: {
		Fstart: int32(0xb05),
		Fend:   int32(0xb0c),
	},
	112: {
		Fstart: int32(0xb0f),
		Fend:   int32(0xb10),
	},
	113: {
		Fstart: int32(0xb13),
		Fend:   int32(0xb28),
	},
	114: {
		Fstart: int32(0xb2a),
		Fend:   int32(0xb30),
	},
	115: {
		Fstart: int32(0xb32),
		Fend:   int32(0xb33),
	},
	116: {
		Fstart: int32(0xb35),
		Fend:   int32(0xb39),
	},
	117: {
		Fstart: int32(0xb3c),
		Fend:   int32(0xb44),
	},
	118: {
		Fstart: int32(0xb47),
		Fend:   int32(0xb48),
	},
	119: {
		Fstart: int32(0xb4b),
		Fend:   int32(0xb4d),
	},
	120: {
		Fstart: int32(0xb55),
		Fend:   int32(0xb57),
	},
	121: {
		Fstart: int32(0xb5c),
		Fend:   int32(0xb5d),
	},
	122: {
		Fstart: int32(0xb5f),
		Fend:   int32(0xb63),
	},
	123: {
		Fstart: int32(0xb66),
		Fend:   int32(0xb6f),
	},
	124: {
		Fstart: int32(0xb71),
		Fend:   int32(0xb77),
	},
	125: {
		Fstart: int32(0xb82),
		Fend:   int32(0xb83),
	},
	126: {
		Fstart: int32(0xb85),
		Fend:   int32(0xb8a),
	},
	127: {
		Fstart: int32(0xb8e),
		Fend:   int32(0xb90),
	},
	128: {
		Fstart: int32(0xb92),
		Fend:   int32(0xb95),
	},
	129: {
		Fstart: int32(0xb99),
		Fend:   int32(0xb9a),
	},
	130: {
		Fstart: int32(0xb9c),
		Fend:   int32(0xb9c),
	},
	131: {
		Fstart: int32(0xb9e),
		Fend:   int32(0xb9f),
	},
	132: {
		Fstart: int32(0xba3),
		Fend:   int32(0xba4),
	},
	133: {
		Fstart: int32(0xba8),
		Fend:   int32(0xbaa),
	},
	134: {
		Fstart: int32(0xbae),
		Fend:   int32(0xbb9),
	},
	135: {
		Fstart: int32(0xbbe),
		Fend:   int32(0xbc2),
	},
	136: {
		Fstart: int32(0xbc6),
		Fend:   int32(0xbc8),
	},
	137: {
		Fstart: int32(0xbca),
		Fend:   int32(0xbcd),
	},
	138: {
		Fstart: int32(0xbd0),
		Fend:   int32(0xbd0),
	},
	139: {
		Fstart: int32(0xbd7),
		Fend:   int32(0xbd7),
	},
	140: {
		Fstart: int32(0xbe6),
		Fend:   int32(0xbf2),
	},
	141: {
		Fstart: int32(0xc00),
		Fend:   int32(0xc0c),
	},
	142: {
		Fstart: int32(0xc0e),
		Fend:   int32(0xc10),
	},
	143: {
		Fstart: int32(0xc12),
		Fend:   int32(0xc28),
	},
	144: {
		Fstart: int32(0xc2a),
		Fend:   int32(0xc39),
	},
	145: {
		Fstart: int32(0xc3c),
		Fend:   int32(0xc44),
	},
	146: {
		Fstart: int32(0xc46),
		Fend:   int32(0xc48),
	},
	147: {
		Fstart: int32(0xc4a),
		Fend:   int32(0xc4d),
	},
	148: {
		Fstart: int32(0xc55),
		Fend:   int32(0xc56),
	},
	149: {
		Fstart: int32(0xc58),
		Fend:   int32(0xc5a),
	},
	150: {
		Fstart: int32(0xc5d),
		Fend:   int32(0xc5d),
	},
	151: {
		Fstart: int32(0xc60),
		Fend:   int32(0xc63),
	},
	152: {
		Fstart: int32(0xc66),
		Fend:   int32(0xc6f),
	},
	153: {
		Fstart: int32(0xc78),
		Fend:   int32(0xc7e),
	},
	154: {
		Fstart: int32(0xc80),
		Fend:   int32(0xc83),
	},
	155: {
		Fstart: int32(0xc85),
		Fend:   int32(0xc8c),
	},
	156: {
		Fstart: int32(0xc8e),
		Fend:   int32(0xc90),
	},
	157: {
		Fstart: int32(0xc92),
		Fend:   int32(0xca8),
	},
	158: {
		Fstart: int32(0xcaa),
		Fend:   int32(0xcb3),
	},
	159: {
		Fstart: int32(0xcb5),
		Fend:   int32(0xcb9),
	},
	160: {
		Fstart: int32(0xcbc),
		Fend:   int32(0xcc4),
	},
	161: {
		Fstart: int32(0xcc6),
		Fend:   int32(0xcc8),
	},
	162: {
		Fstart: int32(0xcca),
		Fend:   int32(0xccd),
	},
	163: {
		Fstart: int32(0xcd5),
		Fend:   int32(0xcd6),
	},
	164: {
		Fstart: int32(0xcdd),
		Fend:   int32(0xcde),
	},
	165: {
		Fstart: int32(0xce0),
		Fend:   int32(0xce3),
	},
	166: {
		Fstart: int32(0xce6),
		Fend:   int32(0xcef),
	},
	167: {
		Fstart: int32(0xcf1),
		Fend:   int32(0xcf3),
	},
	168: {
		Fstart: int32(0xd00),
		Fend:   int32(0xd0c),
	},
	169: {
		Fstart: int32(0xd0e),
		Fend:   int32(0xd10),
	},
	170: {
		Fstart: int32(0xd12),
		Fend:   int32(0xd44),
	},
	171: {
		Fstart: int32(0xd46),
		Fend:   int32(0xd48),
	},
	172: {
		Fstart: int32(0xd4a),
		Fend:   int32(0xd4e),
	},
	173: {
		Fstart: int32(0xd54),
		Fend:   int32(0xd63),
	},
	174: {
		Fstart: int32(0xd66),
		Fend:   int32(0xd78),
	},
	175: {
		Fstart: int32(0xd7a),
		Fend:   int32(0xd7f),
	},
	176: {
		Fstart: int32(0xd81),
		Fend:   int32(0xd83),
	},
	177: {
		Fstart: int32(0xd85),
		Fend:   int32(0xd96),
	},
	178: {
		Fstart: int32(0xd9a),
		Fend:   int32(0xdb1),
	},
	179: {
		Fstart: int32(0xdb3),
		Fend:   int32(0xdbb),
	},
	180: {
		Fstart: int32(0xdbd),
		Fend:   int32(0xdbd),
	},
	181: {
		Fstart: int32(0xdc0),
		Fend:   int32(0xdc6),
	},
	182: {
		Fstart: int32(0xdca),
		Fend:   int32(0xdca),
	},
	183: {
		Fstart: int32(0xdcf),
		Fend:   int32(0xdd4),
	},
	184: {
		Fstart: int32(0xdd6),
		Fend:   int32(0xdd6),
	},
	185: {
		Fstart: int32(0xdd8),
		Fend:   int32(0xddf),
	},
	186: {
		Fstart: int32(0xde6),
		Fend:   int32(0xdef),
	},
	187: {
		Fstart: int32(0xdf2),
		Fend:   int32(0xdf3),
	},
	188: {
		Fstart: int32(0xe01),
		Fend:   int32(0xe3a),
	},
	189: {
		Fstart: int32(0xe40),
		Fend:   int32(0xe4e),
	},
	190: {
		Fstart: int32(0xe50),
		Fend:   int32(0xe59),
	},
	191: {
		Fstart: int32(0xe81),
		Fend:   int32(0xe82),
	},
	192: {
		Fstart: int32(0xe84),
		Fend:   int32(0xe84),
	},
	193: {
		Fstart: int32(0xe86),
		Fend:   int32(0xe8a),
	},
	194: {
		Fstart: int32(0xe8c),
		Fend:   int32(0xea3),
	},
	195: {
		Fstart: int32(0xea5),
		Fend:   int32(0xea5),
	},
	196: {
		Fstart: int32(0xea7),
		Fend:   int32(0xebd),
	},
	197: {
		Fstart: int32(0xec0),
		Fend:   int32(0xec4),
	},
	198: {
		Fstart: int32(0xec6),
		Fend:   int32(0xec6),
	},
	199: {
		Fstart: int32(0xec8),
		Fend:   int32(0xece),
	},
	200: {
		Fstart: int32(0xed0),
		Fend:   int32(0xed9),
	},
	201: {
		Fstart: int32(0xedc),
		Fend:   int32(0xedf),
	},
	202: {
		Fstart: int32(0xf00),
		Fend:   int32(0xf00),
	},
	203: {
		Fstart: int32(0xf18),
		Fend:   int32(0xf19),
	},
	204: {
		Fstart: int32(0xf20),
		Fend:   int32(0xf33),
	},
	205: {
		Fstart: int32(0xf35),
		Fend:   int32(0xf35),
	},
	206: {
		Fstart: int32(0xf37),
		Fend:   int32(0xf37),
	},
	207: {
		Fstart: int32(0xf39),
		Fend:   int32(0xf39),
	},
	208: {
		Fstart: int32(0xf3e),
		Fend:   int32(0xf47),
	},
	209: {
		Fstart: int32(0xf49),
		Fend:   int32(0xf6c),
	},
	210: {
		Fstart: int32(0xf71),
		Fend:   int32(0xf84),
	},
	211: {
		Fstart: int32(0xf86),
		Fend:   int32(0xf97),
	},
	212: {
		Fstart: int32(0xf99),
		Fend:   int32(0xfbc),
	},
	213: {
		Fstart: int32(0xfc6),
		Fend:   int32(0xfc6),
	},
	214: {
		Fstart: int32(0x1000),
		Fend:   int32(0x1049),
	},
	215: {
		Fstart: int32(0x1050),
		Fend:   int32(0x109d),
	},
	216: {
		Fstart: int32(0x10a0),
		Fend:   int32(0x10c5),
	},
	217: {
		Fstart: int32(0x10c7),
		Fend:   int32(0x10c7),
	},
	218: {
		Fstart: int32(0x10cd),
		Fend:   int32(0x10cd),
	},
	219: {
		Fstart: int32(0x10d0),
		Fend:   int32(0x10fa),
	},
	220: {
		Fstart: int32(0x10fc),
		Fend:   int32(0x1248),
	},
	221: {
		Fstart: int32(0x124a),
		Fend:   int32(0x124d),
	},
	222: {
		Fstart: int32(0x1250),
		Fend:   int32(0x1256),
	},
	223: {
		Fstart: int32(0x1258),
		Fend:   int32(0x1258),
	},
	224: {
		Fstart: int32(0x125a),
		Fend:   int32(0x125d),
	},
	225: {
		Fstart: int32(0x1260),
		Fend:   int32(0x1288),
	},
	226: {
		Fstart: int32(0x128a),
		Fend:   int32(0x128d),
	},
	227: {
		Fstart: int32(0x1290),
		Fend:   int32(0x12b0),
	},
	228: {
		Fstart: int32(0x12b2),
		Fend:   int32(0x12b5),
	},
	229: {
		Fstart: int32(0x12b8),
		Fend:   int32(0x12be),
	},
	230: {
		Fstart: int32(0x12c0),
		Fend:   int32(0x12c0),
	},
	231: {
		Fstart: int32(0x12c2),
		Fend:   int32(0x12c5),
	},
	232: {
		Fstart: int32(0x12c8),
		Fend:   int32(0x12d6),
	},
	233: {
		Fstart: int32(0x12d8),
		Fend:   int32(0x1310),
	},
	234: {
		Fstart: int32(0x1312),
		Fend:   int32(0x1315),
	},
	235: {
		Fstart: int32(0x1318),
		Fend:   int32(0x135a),
	},
	236: {
		Fstart: int32(0x135d),
		Fend:   int32(0x135f),
	},
	237: {
		Fstart: int32(0x1369),
		Fend:   int32(0x137c),
	},
	238: {
		Fstart: int32(0x1380),
		Fend:   int32(0x138f),
	},
	239: {
		Fstart: int32(0x13a0),
		Fend:   int32(0x13f5),
	},
	240: {
		Fstart: int32(0x13f8),
		Fend:   int32(0x13fd),
	},
	241: {
		Fstart: int32(0x1401),
		Fend:   int32(0x166c),
	},
	242: {
		Fstart: int32(0x166f),
		Fend:   int32(0x167f),
	},
	243: {
		Fstart: int32(0x1681),
		Fend:   int32(0x169a),
	},
	244: {
		Fstart: int32(0x16a0),
		Fend:   int32(0x16ea),
	},
	245: {
		Fstart: int32(0x16ee),
		Fend:   int32(0x16f8),
	},
	246: {
		Fstart: int32(0x1700),
		Fend:   int32(0x1715),
	},
	247: {
		Fstart: int32(0x171f),
		Fend:   int32(0x1734),
	},
	248: {
		Fstart: int32(0x1740),
		Fend:   int32(0x1753),
	},
	249: {
		Fstart: int32(0x1760),
		Fend:   int32(0x176c),
	},
	250: {
		Fstart: int32(0x176e),
		Fend:   int32(0x1770),
	},
	251: {
		Fstart: int32(0x1772),
		Fend:   int32(0x1773),
	},
	252: {
		Fstart: int32(0x1780),
		Fend:   int32(0x17d3),
	},
	253: {
		Fstart: int32(0x17d7),
		Fend:   int32(0x17d7),
	},
	254: {
		Fstart: int32(0x17dc),
		Fend:   int32(0x17dd),
	},
	255: {
		Fstart: int32(0x17e0),
		Fend:   int32(0x17e9),
	},
	256: {
		Fstart: int32(0x17f0),
		Fend:   int32(0x17f9),
	},
	257: {
		Fstart: int32(0x180b),
		Fend:   int32(0x180d),
	},
	258: {
		Fstart: int32(0x180f),
		Fend:   int32(0x1819),
	},
	259: {
		Fstart: int32(0x1820),
		Fend:   int32(0x1878),
	},
	260: {
		Fstart: int32(0x1880),
		Fend:   int32(0x18aa),
	},
	261: {
		Fstart: int32(0x18b0),
		Fend:   int32(0x18f5),
	},
	262: {
		Fstart: int32(0x1900),
		Fend:   int32(0x191e),
	},
	263: {
		Fstart: int32(0x1920),
		Fend:   int32(0x192b),
	},
	264: {
		Fstart: int32(0x1930),
		Fend:   int32(0x193b),
	},
	265: {
		Fstart: int32(0x1946),
		Fend:   int32(0x196d),
	},
	266: {
		Fstart: int32(0x1970),
		Fend:   int32(0x1974),
	},
	267: {
		Fstart: int32(0x1980),
		Fend:   int32(0x19ab),
	},
	268: {
		Fstart: int32(0x19b0),
		Fend:   int32(0x19c9),
	},
	269: {
		Fstart: int32(0x19d0),
		Fend:   int32(0x19da),
	},
	270: {
		Fstart: int32(0x1a00),
		Fend:   int32(0x1a1b),
	},
	271: {
		Fstart: int32(0x1a20),
		Fend:   int32(0x1a5e),
	},
	272: {
		Fstart: int32(0x1a60),
		Fend:   int32(0x1a7c),
	},
	273: {
		Fstart: int32(0x1a7f),
		Fend:   int32(0x1a89),
	},
	274: {
		Fstart: int32(0x1a90),
		Fend:   int32(0x1a99),
	},
	275: {
		Fstart: int32(0x1aa7),
		Fend:   int32(0x1aa7),
	},
	276: {
		Fstart: int32(0x1ab0),
		Fend:   int32(0x1ace),
	},
	277: {
		Fstart: int32(0x1b00),
		Fend:   int32(0x1b4c),
	},
	278: {
		Fstart: int32(0x1b50),
		Fend:   int32(0x1b59),
	},
	279: {
		Fstart: int32(0x1b6b),
		Fend:   int32(0x1b73),
	},
	280: {
		Fstart: int32(0x1b80),
		Fend:   int32(0x1bf3),
	},
	281: {
		Fstart: int32(0x1c00),
		Fend:   int32(0x1c37),
	},
	282: {
		Fstart: int32(0x1c40),
		Fend:   int32(0x1c49),
	},
	283: {
		Fstart: int32(0x1c4d),
		Fend:   int32(0x1c7d),
	},
	284: {
		Fstart: int32(0x1c80),
		Fend:   int32(0x1c88),
	},
	285: {
		Fstart: int32(0x1c90),
		Fend:   int32(0x1cba),
	},
	286: {
		Fstart: int32(0x1cbd),
		Fend:   int32(0x1cbf),
	},
	287: {
		Fstart: int32(0x1cd0),
		Fend:   int32(0x1cd2),
	},
	288: {
		Fstart: int32(0x1cd4),
		Fend:   int32(0x1cfa),
	},
	289: {
		Fstart: int32(0x1d00),
		Fend:   int32(0x1f15),
	},
	290: {
		Fstart: int32(0x1f18),
		Fend:   int32(0x1f1d),
	},
	291: {
		Fstart: int32(0x1f20),
		Fend:   int32(0x1f45),
	},
	292: {
		Fstart: int32(0x1f48),
		Fend:   int32(0x1f4d),
	},
	293: {
		Fstart: int32(0x1f50),
		Fend:   int32(0x1f57),
	},
	294: {
		Fstart: int32(0x1f59),
		Fend:   int32(0x1f59),
	},
	295: {
		Fstart: int32(0x1f5b),
		Fend:   int32(0x1f5b),
	},
	296: {
		Fstart: int32(0x1f5d),
		Fend:   int32(0x1f5d),
	},
	297: {
		Fstart: int32(0x1f5f),
		Fend:   int32(0x1f7d),
	},
	298: {
		Fstart: int32(0x1f80),
		Fend:   int32(0x1fb4),
	},
	299: {
		Fstart: int32(0x1fb6),
		Fend:   int32(0x1fbc),
	},
	300: {
		Fstart: int32(0x1fbe),
		Fend:   int32(0x1fbe),
	},
	301: {
		Fstart: int32(0x1fc2),
		Fend:   int32(0x1fc4),
	},
	302: {
		Fstart: int32(0x1fc6),
		Fend:   int32(0x1fcc),
	},
	303: {
		Fstart: int32(0x1fd0),
		Fend:   int32(0x1fd3),
	},
	304: {
		Fstart: int32(0x1fd6),
		Fend:   int32(0x1fdb),
	},
	305: {
		Fstart: int32(0x1fe0),
		Fend:   int32(0x1fec),
	},
	306: {
		Fstart: int32(0x1ff2),
		Fend:   int32(0x1ff4),
	},
	307: {
		Fstart: int32(0x1ff6),
		Fend:   int32(0x1ffc),
	},
	308: {
		Fstart: int32(0x2070),
		Fend:   int32(0x2071),
	},
	309: {
		Fstart: int32(0x2074),
		Fend:   int32(0x2079),
	},
	310: {
		Fstart: int32(0x207f),
		Fend:   int32(0x2089),
	},
	311: {
		Fstart: int32(0x2090),
		Fend:   int32(0x209c),
	},
	312: {
		Fstart: int32(0x20d0),
		Fend:   int32(0x20f0),
	},
	313: {
		Fstart: int32(0x2102),
		Fend:   int32(0x2102),
	},
	314: {
		Fstart: int32(0x2107),
		Fend:   int32(0x2107),
	},
	315: {
		Fstart: int32(0x210a),
		Fend:   int32(0x2113),
	},
	316: {
		Fstart: int32(0x2115),
		Fend:   int32(0x2115),
	},
	317: {
		Fstart: int32(0x2119),
		Fend:   int32(0x211d),
	},
	318: {
		Fstart: int32(0x2124),
		Fend:   int32(0x2124),
	},
	319: {
		Fstart: int32(0x2126),
		Fend:   int32(0x2126),
	},
	320: {
		Fstart: int32(0x2128),
		Fend:   int32(0x2128),
	},
	321: {
		Fstart: int32(0x212a),
		Fend:   int32(0x212d),
	},
	322: {
		Fstart: int32(0x212f),
		Fend:   int32(0x2139),
	},
	323: {
		Fstart: int32(0x213c),
		Fend:   int32(0x213f),
	},
	324: {
		Fstart: int32(0x2145),
		Fend:   int32(0x2149),
	},
	325: {
		Fstart: int32(0x214e),
		Fend:   int32(0x214e),
	},
	326: {
		Fstart: int32(0x2150),
		Fend:   int32(0x2189),
	},
	327: {
		Fstart: int32(0x2460),
		Fend:   int32(0x249b),
	},
	328: {
		Fstart: int32(0x24ea),
		Fend:   int32(0x24ff),
	},
	329: {
		Fstart: int32(0x2776),
		Fend:   int32(0x2793),
	},
	330: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2ce4),
	},
	331: {
		Fstart: int32(0x2ceb),
		Fend:   int32(0x2cf3),
	},
	332: {
		Fstart: int32(0x2cfd),
		Fend:   int32(0x2cfd),
	},
	333: {
		Fstart: int32(0x2d00),
		Fend:   int32(0x2d25),
	},
	334: {
		Fstart: int32(0x2d27),
		Fend:   int32(0x2d27),
	},
	335: {
		Fstart: int32(0x2d2d),
		Fend:   int32(0x2d2d),
	},
	336: {
		Fstart: int32(0x2d30),
		Fend:   int32(0x2d67),
	},
	337: {
		Fstart: int32(0x2d6f),
		Fend:   int32(0x2d6f),
	},
	338: {
		Fstart: int32(0x2d7f),
		Fend:   int32(0x2d96),
	},
	339: {
		Fstart: int32(0x2da0),
		Fend:   int32(0x2da6),
	},
	340: {
		Fstart: int32(0x2da8),
		Fend:   int32(0x2dae),
	},
	341: {
		Fstart: int32(0x2db0),
		Fend:   int32(0x2db6),
	},
	342: {
		Fstart: int32(0x2db8),
		Fend:   int32(0x2dbe),
	},
	343: {
		Fstart: int32(0x2dc0),
		Fend:   int32(0x2dc6),
	},
	344: {
		Fstart: int32(0x2dc8),
		Fend:   int32(0x2dce),
	},
	345: {
		Fstart: int32(0x2dd0),
		Fend:   int32(0x2dd6),
	},
	346: {
		Fstart: int32(0x2dd8),
		Fend:   int32(0x2dde),
	},
	347: {
		Fstart: int32(0x2de0),
		Fend:   int32(0x2dff),
	},
	348: {
		Fstart: int32(0x2e2f),
		Fend:   int32(0x2e2f),
	},
	349: {
		Fstart: int32(0x3005),
		Fend:   int32(0x3007),
	},
	350: {
		Fstart: int32(0x3021),
		Fend:   int32(0x302f),
	},
	351: {
		Fstart: int32(0x3031),
		Fend:   int32(0x3035),
	},
	352: {
		Fstart: int32(0x3038),
		Fend:   int32(0x303c),
	},
	353: {
		Fstart: int32(0x3041),
		Fend:   int32(0x3096),
	},
	354: {
		Fstart: int32(0x3099),
		Fend:   int32(0x309a),
	},
	355: {
		Fstart: int32(0x309d),
		Fend:   int32(0x309f),
	},
	356: {
		Fstart: int32(0x30a1),
		Fend:   int32(0x30fa),
	},
	357: {
		Fstart: int32(0x30fc),
		Fend:   int32(0x30ff),
	},
	358: {
		Fstart: int32(0x3105),
		Fend:   int32(0x312f),
	},
	359: {
		Fstart: int32(0x3131),
		Fend:   int32(0x318e),
	},
	360: {
		Fstart: int32(0x3192),
		Fend:   int32(0x3195),
	},
	361: {
		Fstart: int32(0x31a0),
		Fend:   int32(0x31bf),
	},
	362: {
		Fstart: int32(0x31f0),
		Fend:   int32(0x31ff),
	},
	363: {
		Fstart: int32(0x3220),
		Fend:   int32(0x3229),
	},
	364: {
		Fstart: int32(0x3248),
		Fend:   int32(0x324f),
	},
	365: {
		Fstart: int32(0x3251),
		Fend:   int32(0x325f),
	},
	366: {
		Fstart: int32(0x3280),
		Fend:   int32(0x3289),
	},
	367: {
		Fstart: int32(0x32b1),
		Fend:   int32(0x32bf),
	},
	368: {
		Fstart: int32(0x3400),
		Fend:   int32(0x3400),
	},
	369: {
		Fstart: int32(0x4dbf),
		Fend:   int32(0x4dbf),
	},
	370: {
		Fstart: int32(0x4e00),
		Fend:   int32(0xa48c),
	},
	371: {
		Fstart: int32(0xa4d0),
		Fend:   int32(0xa4fd),
	},
	372: {
		Fstart: int32(0xa500),
		Fend:   int32(0xa60c),
	},
	373: {
		Fstart: int32(0xa610),
		Fend:   int32(0xa62b),
	},
	374: {
		Fstart: int32(0xa640),
		Fend:   int32(0xa672),
	},
	375: {
		Fstart: int32(0xa674),
		Fend:   int32(0xa67d),
	},
	376: {
		Fstart: int32(0xa67f),
		Fend:   int32(0xa6f1),
	},
	377: {
		Fstart: int32(0xa717),
		Fend:   int32(0xa71f),
	},
	378: {
		Fstart: int32(0xa722),
		Fend:   int32(0xa788),
	},
	379: {
		Fstart: int32(0xa78b),
		Fend:   int32(0xa7ca),
	},
	380: {
		Fstart: int32(0xa7d0),
		Fend:   int32(0xa7d1),
	},
	381: {
		Fstart: int32(0xa7d3),
		Fend:   int32(0xa7d3),
	},
	382: {
		Fstart: int32(0xa7d5),
		Fend:   int32(0xa7d9),
	},
	383: {
		Fstart: int32(0xa7f2),
		Fend:   int32(0xa827),
	},
	384: {
		Fstart: int32(0xa82c),
		Fend:   int32(0xa82c),
	},
	385: {
		Fstart: int32(0xa830),
		Fend:   int32(0xa835),
	},
	386: {
		Fstart: int32(0xa840),
		Fend:   int32(0xa873),
	},
	387: {
		Fstart: int32(0xa880),
		Fend:   int32(0xa8c5),
	},
	388: {
		Fstart: int32(0xa8d0),
		Fend:   int32(0xa8d9),
	},
	389: {
		Fstart: int32(0xa8e0),
		Fend:   int32(0xa8f7),
	},
	390: {
		Fstart: int32(0xa8fb),
		Fend:   int32(0xa8fb),
	},
	391: {
		Fstart: int32(0xa8fd),
		Fend:   int32(0xa92d),
	},
	392: {
		Fstart: int32(0xa930),
		Fend:   int32(0xa953),
	},
	393: {
		Fstart: int32(0xa960),
		Fend:   int32(0xa97c),
	},
	394: {
		Fstart: int32(0xa980),
		Fend:   int32(0xa9c0),
	},
	395: {
		Fstart: int32(0xa9cf),
		Fend:   int32(0xa9d9),
	},
	396: {
		Fstart: int32(0xa9e0),
		Fend:   int32(0xa9fe),
	},
	397: {
		Fstart: int32(0xaa00),
		Fend:   int32(0xaa36),
	},
	398: {
		Fstart: int32(0xaa40),
		Fend:   int32(0xaa4d),
	},
	399: {
		Fstart: int32(0xaa50),
		Fend:   int32(0xaa59),
	},
	400: {
		Fstart: int32(0xaa60),
		Fend:   int32(0xaa76),
	},
	401: {
		Fstart: int32(0xaa7a),
		Fend:   int32(0xaac2),
	},
	402: {
		Fstart: int32(0xaadb),
		Fend:   int32(0xaadd),
	},
	403: {
		Fstart: int32(0xaae0),
		Fend:   int32(0xaaef),
	},
	404: {
		Fstart: int32(0xaaf2),
		Fend:   int32(0xaaf6),
	},
	405: {
		Fstart: int32(0xab01),
		Fend:   int32(0xab06),
	},
	406: {
		Fstart: int32(0xab09),
		Fend:   int32(0xab0e),
	},
	407: {
		Fstart: int32(0xab11),
		Fend:   int32(0xab16),
	},
	408: {
		Fstart: int32(0xab20),
		Fend:   int32(0xab26),
	},
	409: {
		Fstart: int32(0xab28),
		Fend:   int32(0xab2e),
	},
	410: {
		Fstart: int32(0xab30),
		Fend:   int32(0xab5a),
	},
	411: {
		Fstart: int32(0xab5c),
		Fend:   int32(0xab69),
	},
	412: {
		Fstart: int32(0xab70),
		Fend:   int32(0xabea),
	},
	413: {
		Fstart: int32(0xabec),
		Fend:   int32(0xabed),
	},
	414: {
		Fstart: int32(0xabf0),
		Fend:   int32(0xabf9),
	},
	415: {
		Fstart: int32(0xac00),
		Fend:   int32(0xac00),
	},
	416: {
		Fstart: int32(0xd7a3),
		Fend:   int32(0xd7a3),
	},
	417: {
		Fstart: int32(0xd7b0),
		Fend:   int32(0xd7c6),
	},
	418: {
		Fstart: int32(0xd7cb),
		Fend:   int32(0xd7fb),
	},
	419: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfa6d),
	},
	420: {
		Fstart: int32(0xfa70),
		Fend:   int32(0xfad9),
	},
	421: {
		Fstart: int32(0xfb00),
		Fend:   int32(0xfb06),
	},
	422: {
		Fstart: int32(0xfb13),
		Fend:   int32(0xfb17),
	},
	423: {
		Fstart: int32(0xfb1d),
		Fend:   int32(0xfb28),
	},
	424: {
		Fstart: int32(0xfb2a),
		Fend:   int32(0xfb36),
	},
	425: {
		Fstart: int32(0xfb38),
		Fend:   int32(0xfb3c),
	},
	426: {
		Fstart: int32(0xfb3e),
		Fend:   int32(0xfb3e),
	},
	427: {
		Fstart: int32(0xfb40),
		Fend:   int32(0xfb41),
	},
	428: {
		Fstart: int32(0xfb43),
		Fend:   int32(0xfb44),
	},
	429: {
		Fstart: int32(0xfb46),
		Fend:   int32(0xfbb1),
	},
	430: {
		Fstart: int32(0xfbd3),
		Fend:   int32(0xfd3d),
	},
	431: {
		Fstart: int32(0xfd50),
		Fend:   int32(0xfd8f),
	},
	432: {
		Fstart: int32(0xfd92),
		Fend:   int32(0xfdc7),
	},
	433: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfdfb),
	},
	434: {
		Fstart: int32(0xfe00),
		Fend:   int32(0xfe0f),
	},
	435: {
		Fstart: int32(0xfe20),
		Fend:   int32(0xfe2f),
	},
	436: {
		Fstart: int32(0xfe70),
		Fend:   int32(0xfe74),
	},
	437: {
		Fstart: int32(0xfe76),
		Fend:   int32(0xfefc),
	},
	438: {
		Fstart: int32(0xff10),
		Fend:   int32(0xff19),
	},
	439: {
		Fstart: int32(0xff21),
		Fend:   int32(0xff3a),
	},
	440: {
		Fstart: int32(0xff41),
		Fend:   int32(0xff5a),
	},
	441: {
		Fstart: int32(0xff66),
		Fend:   int32(0xffbe),
	},
	442: {
		Fstart: int32(0xffc2),
		Fend:   int32(0xffc7),
	},
	443: {
		Fstart: int32(0xffca),
		Fend:   int32(0xffcf),
	},
	444: {
		Fstart: int32(0xffd2),
		Fend:   int32(0xffd7),
	},
	445: {
		Fstart: int32(0xffda),
		Fend:   int32(0xffdc),
	},
	446: {
		Fstart: int32(0x10000),
		Fend:   int32(0x1000b),
	},
	447: {
		Fstart: int32(0x1000d),
		Fend:   int32(0x10026),
	},
	448: {
		Fstart: int32(0x10028),
		Fend:   int32(0x1003a),
	},
	449: {
		Fstart: int32(0x1003c),
		Fend:   int32(0x1003d),
	},
	450: {
		Fstart: int32(0x1003f),
		Fend:   int32(0x1004d),
	},
	451: {
		Fstart: int32(0x10050),
		Fend:   int32(0x1005d),
	},
	452: {
		Fstart: int32(0x10080),
		Fend:   int32(0x100fa),
	},
	453: {
		Fstart: int32(0x10107),
		Fend:   int32(0x10133),
	},
	454: {
		Fstart: int32(0x10140),
		Fend:   int32(0x10178),
	},
	455: {
		Fstart: int32(0x1018a),
		Fend:   int32(0x1018b),
	},
	456: {
		Fstart: int32(0x101fd),
		Fend:   int32(0x101fd),
	},
	457: {
		Fstart: int32(0x10280),
		Fend:   int32(0x1029c),
	},
	458: {
		Fstart: int32(0x102a0),
		Fend:   int32(0x102d0),
	},
	459: {
		Fstart: int32(0x102e0),
		Fend:   int32(0x102fb),
	},
	460: {
		Fstart: int32(0x10300),
		Fend:   int32(0x10323),
	},
	461: {
		Fstart: int32(0x1032d),
		Fend:   int32(0x1034a),
	},
	462: {
		Fstart: int32(0x10350),
		Fend:   int32(0x1037a),
	},
	463: {
		Fstart: int32(0x10380),
		Fend:   int32(0x1039d),
	},
	464: {
		Fstart: int32(0x103a0),
		Fend:   int32(0x103c3),
	},
	465: {
		Fstart: int32(0x103c8),
		Fend:   int32(0x103cf),
	},
	466: {
		Fstart: int32(0x103d1),
		Fend:   int32(0x103d5),
	},
	467: {
		Fstart: int32(0x10400),
		Fend:   int32(0x1049d),
	},
	468: {
		Fstart: int32(0x104a0),
		Fend:   int32(0x104a9),
	},
	469: {
		Fstart: int32(0x104b0),
		Fend:   int32(0x104d3),
	},
	470: {
		Fstart: int32(0x104d8),
		Fend:   int32(0x104fb),
	},
	471: {
		Fstart: int32(0x10500),
		Fend:   int32(0x10527),
	},
	472: {
		Fstart: int32(0x10530),
		Fend:   int32(0x10563),
	},
	473: {
		Fstart: int32(0x10570),
		Fend:   int32(0x1057a),
	},
	474: {
		Fstart: int32(0x1057c),
		Fend:   int32(0x1058a),
	},
	475: {
		Fstart: int32(0x1058c),
		Fend:   int32(0x10592),
	},
	476: {
		Fstart: int32(0x10594),
		Fend:   int32(0x10595),
	},
	477: {
		Fstart: int32(0x10597),
		Fend:   int32(0x105a1),
	},
	478: {
		Fstart: int32(0x105a3),
		Fend:   int32(0x105b1),
	},
	479: {
		Fstart: int32(0x105b3),
		Fend:   int32(0x105b9),
	},
	480: {
		Fstart: int32(0x105bb),
		Fend:   int32(0x105bc),
	},
	481: {
		Fstart: int32(0x10600),
		Fend:   int32(0x10736),
	},
	482: {
		Fstart: int32(0x10740),
		Fend:   int32(0x10755),
	},
	483: {
		Fstart: int32(0x10760),
		Fend:   int32(0x10767),
	},
	484: {
		Fstart: int32(0x10780),
		Fend:   int32(0x10785),
	},
	485: {
		Fstart: int32(0x10787),
		Fend:   int32(0x107b0),
	},
	486: {
		Fstart: int32(0x107b2),
		Fend:   int32(0x107ba),
	},
	487: {
		Fstart: int32(0x10800),
		Fend:   int32(0x10805),
	},
	488: {
		Fstart: int32(0x10808),
		Fend:   int32(0x10808),
	},
	489: {
		Fstart: int32(0x1080a),
		Fend:   int32(0x10835),
	},
	490: {
		Fstart: int32(0x10837),
		Fend:   int32(0x10838),
	},
	491: {
		Fstart: int32(0x1083c),
		Fend:   int32(0x1083c),
	},
	492: {
		Fstart: int32(0x1083f),
		Fend:   int32(0x10855),
	},
	493: {
		Fstart: int32(0x10858),
		Fend:   int32(0x10876),
	},
	494: {
		Fstart: int32(0x10879),
		Fend:   int32(0x1089e),
	},
	495: {
		Fstart: int32(0x108a7),
		Fend:   int32(0x108af),
	},
	496: {
		Fstart: int32(0x108e0),
		Fend:   int32(0x108f2),
	},
	497: {
		Fstart: int32(0x108f4),
		Fend:   int32(0x108f5),
	},
	498: {
		Fstart: int32(0x108fb),
		Fend:   int32(0x1091b),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v4 uint8
	var half_size, i, i1, i2, i3, i4, i5, i6, i7, i8, index, mid_index, size uint32_t
	var lookahead1, v3 int32_t
	var range_token, range_token1, v2 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i2, i3, i4, i5, i6, i7, i8, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v2, v3, v4
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
			state = uint16(20)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(100)/libc.Uint64FromInt64(2)) {
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
		if int32('2') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(26)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(25)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(32)
			goto next_state
		}
		if lookahead1 == int32(0xa9) || lookahead1 == int32(0xae) || lookahead1 == int32(0x203c) || lookahead1 == int32(0x2049) || lookahead1 == int32(0x2122) || int32(0x2194) <= lookahead1 && lookahead1 <= int32(0x2199) || lookahead1 == int32(0x21a9) || lookahead1 == int32(0x21aa) || lookahead1 == int32(0x231a) || lookahead1 == int32(0x231b) || lookahead1 == int32(0x2328) || lookahead1 == int32(0x23cf) || int32(0x23e9) <= lookahead1 && lookahead1 <= int32(0x23f3) || int32(0x23f8) <= lookahead1 && lookahead1 <= int32(0x23fa) || lookahead1 == int32(0x24c2) || lookahead1 == int32(0x25aa) || lookahead1 == int32(0x25ab) || lookahead1 == int32(0x25b6) || lookahead1 == int32(0x25c0) || int32(0x25fb) <= lookahead1 && lookahead1 <= int32(0x25fe) || int32(0x2600) <= lookahead1 && lookahead1 <= int32(0x2604) || lookahead1 == int32(0x260e) || lookahead1 == int32(0x2611) || lookahead1 == int32(0x2614) || lookahead1 == int32(0x2615) || lookahead1 == int32(0x2618) || lookahead1 == int32(0x261d) || lookahead1 == int32(0x2620) || lookahead1 == int32(0x2622) || lookahead1 == int32(0x2623) || lookahead1 == int32(0x2626) || lookahead1 == int32(0x262a) || lookahead1 == int32(0x262e) || lookahead1 == int32(0x262f) || int32(0x2638) <= lookahead1 && lookahead1 <= int32(0x263a) || lookahead1 == int32(0x2640) || lookahead1 == int32(0x2642) || int32(0x2648) <= lookahead1 && lookahead1 <= int32(0x2653) || lookahead1 == int32(0x265f) || lookahead1 == int32(0x2660) || lookahead1 == int32(0x2663) || lookahead1 == int32(0x2665) || lookahead1 == int32(0x2666) || lookahead1 == int32(0x2668) || lookahead1 == int32(0x267b) || lookahead1 == int32(0x267e) || lookahead1 == int32(0x267f) || int32(0x2692) <= lookahead1 && lookahead1 <= int32(0x2697) || lookahead1 == int32(0x2699) || lookahead1 == int32(0x269b) || lookahead1 == int32(0x269c) || lookahead1 == int32(0x26a0) || lookahead1 == int32(0x26a1) || lookahead1 == int32(0x26a7) || lookahead1 == int32(0x26aa) || lookahead1 == int32(0x26ab) || lookahead1 == int32(0x26b0) || lookahead1 == int32(0x26b1) || lookahead1 == int32(0x26bd) || lookahead1 == int32(0x26be) || lookahead1 == int32(0x26c4) || lookahead1 == int32(0x26c5) || lookahead1 == int32(0x26c8) || lookahead1 == int32(0x26ce) || lookahead1 == int32(0x26cf) || lookahead1 == int32(0x26d1) || lookahead1 == int32(0x26d3) || lookahead1 == int32(0x26d4) || lookahead1 == int32(0x26e9) || lookahead1 == int32(0x26ea) || int32(0x26f0) <= lookahead1 && lookahead1 <= int32(0x26f5) || int32(0x26f7) <= lookahead1 && lookahead1 <= int32(0x26fa) || lookahead1 == int32(0x26fd) || lookahead1 == int32(0x2702) || lookahead1 == int32(0x2705) || int32(0x2708) <= lookahead1 && lookahead1 <= int32(0x270d) || lookahead1 == int32(0x270f) || lookahead1 == int32(0x2712) || lookahead1 == int32(0x2714) || lookahead1 == int32(0x2716) || lookahead1 == int32(0x271d) || lookahead1 == int32(0x2721) || lookahead1 == int32(0x2728) || lookahead1 == int32(0x2733) || lookahead1 == int32(0x2734) || lookahead1 == int32(0x2744) || lookahead1 == int32(0x2747) || lookahead1 == int32(0x274c) || lookahead1 == int32(0x274e) || int32(0x2753) <= lookahead1 && lookahead1 <= int32(0x2755) || lookahead1 == int32(0x2757) || lookahead1 == int32(0x2763) || lookahead1 == int32(0x2764) || int32(0x2795) <= lookahead1 && lookahead1 <= int32(0x2797) || lookahead1 == int32(0x27a1) || lookahead1 == int32(0x27b0) || lookahead1 == int32(0x27bf) || lookahead1 == int32(0x2934) || lookahead1 == int32(0x2935) || int32(0x2b05) <= lookahead1 && lookahead1 <= int32(0x2b07) || lookahead1 == int32(0x2b1b) || lookahead1 == int32(0x2b1c) || lookahead1 == int32(0x2b50) || lookahead1 == int32(0x2b55) || lookahead1 == int32(0x3030) || lookahead1 == int32(0x303d) || lookahead1 == int32(0x3297) || lookahead1 == int32(0x3299) || lookahead1 == int32(0x1f004) || lookahead1 == int32(0x1f0cf) || lookahead1 == int32(0x1f170) || lookahead1 == int32(0x1f171) || lookahead1 == int32(0x1f17e) || lookahead1 == int32(0x1f17f) || lookahead1 == int32(0x1f18e) || int32(0x1f191) <= lookahead1 && lookahead1 <= int32(0x1f19a) || int32(0x1f1e6) <= lookahead1 && lookahead1 <= int32(0x1f1ff) || lookahead1 == int32(0x1f201) || lookahead1 == int32(0x1f202) || lookahead1 == int32(0x1f21a) || lookahead1 == int32(0x1f22f) || int32(0x1f232) <= lookahead1 && lookahead1 <= int32(0x1f23a) || lookahead1 == int32(0x1f250) || lookahead1 == int32(0x1f251) || int32(0x1f300) <= lookahead1 && lookahead1 <= int32(0x1f321) || int32(0x1f324) <= lookahead1 && lookahead1 <= int32(0x1f393) || lookahead1 == int32(0x1f396) || lookahead1 == int32(0x1f397) || int32(0x1f399) <= lookahead1 && lookahead1 <= int32(0x1f39b) || int32(0x1f39e) <= lookahead1 && lookahead1 <= int32(0x1f3f0) || int32(0x1f3f3) <= lookahead1 && lookahead1 <= int32(0x1f3f5) || int32(0x1f3f7) <= lookahead1 && lookahead1 <= int32(0x1f4fd) || int32(0x1f4ff) <= lookahead1 && lookahead1 <= int32(0x1f53d) || int32(0x1f549) <= lookahead1 && lookahead1 <= int32(0x1f54e) || int32(0x1f550) <= lookahead1 && lookahead1 <= int32(0x1f567) || lookahead1 == int32(0x1f56f) || lookahead1 == int32(0x1f570) || int32(0x1f573) <= lookahead1 && lookahead1 <= int32(0x1f57a) || lookahead1 == int32(0x1f587) || int32(0x1f58a) <= lookahead1 && lookahead1 <= int32(0x1f58d) || lookahead1 == int32(0x1f590) || lookahead1 == int32(0x1f595) || lookahead1 == int32(0x1f596) || lookahead1 == int32(0x1f5a4) || lookahead1 == int32(0x1f5a5) || lookahead1 == int32(0x1f5a8) || lookahead1 == int32(0x1f5b1) || lookahead1 == int32(0x1f5b2) || lookahead1 == int32(0x1f5bc) || int32(0x1f5c2) <= lookahead1 && lookahead1 <= int32(0x1f5c4) || int32(0x1f5d1) <= lookahead1 && lookahead1 <= int32(0x1f5d3) || int32(0x1f5dc) <= lookahead1 && lookahead1 <= int32(0x1f5de) || lookahead1 == int32(0x1f5e1) || lookahead1 == int32(0x1f5e3) || lookahead1 == int32(0x1f5e8) || lookahead1 == int32(0x1f5ef) || lookahead1 == int32(0x1f5f3) || int32(0x1f5fa) <= lookahead1 && lookahead1 <= int32(0x1f64f) || int32(0x1f680) <= lookahead1 && lookahead1 <= int32(0x1f6c5) || int32(0x1f6cb) <= lookahead1 && lookahead1 <= int32(0x1f6d2) || int32(0x1f6d5) <= lookahead1 && lookahead1 <= int32(0x1f6d7) || int32(0x1f6dc) <= lookahead1 && lookahead1 <= int32(0x1f6e5) || lookahead1 == int32(0x1f6e9) || lookahead1 == int32(0x1f6eb) || lookahead1 == int32(0x1f6ec) || lookahead1 == int32(0x1f6f0) || int32(0x1f6f3) <= lookahead1 && lookahead1 <= int32(0x1f6fc) || int32(0x1f7e0) <= lookahead1 && lookahead1 <= int32(0x1f7eb) || lookahead1 == int32(0x1f7f0) || int32(0x1f90c) <= lookahead1 && lookahead1 <= int32(0x1f93a) || int32(0x1f93c) <= lookahead1 && lookahead1 <= int32(0x1f945) || int32(0x1f947) <= lookahead1 && lookahead1 <= int32(0x1f9ff) || int32(0x1fa70) <= lookahead1 && lookahead1 <= int32(0x1fa7c) || int32(0x1fa80) <= lookahead1 && lookahead1 <= int32(0x1fa88) || int32(0x1fa90) <= lookahead1 && lookahead1 <= int32(0x1fabd) || int32(0x1fabf) <= lookahead1 && lookahead1 <= int32(0x1fac5) || int32(0x1face) <= lookahead1 && lookahead1 <= int32(0x1fadb) || int32(0x1fae0) <= lookahead1 && lookahead1 <= int32(0x1fae8) || int32(0x1faf0) <= lookahead1 && lookahead1 <= int32(0x1faf8) {
			state = uint16(25)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym___identifier_char_no_digit_sign_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(499) - index
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
			state = uint16(25)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(1):
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
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(44)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(63)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym___identifier_char_no_digit_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(499) - index
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
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(2):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
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
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(63)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym__identifier_char_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(499) - index
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
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(3):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
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
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(44)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(4):
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token4[i4]) == lookahead1 {
				state = map_token4[i4+uint32(1)]
				goto next_state
			}
			goto _17
		_17:
			;
			i4 = i4 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(5):
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token5[i5]) == lookahead1 {
				state = map_token5[i5+uint32(1)]
				goto next_state
			}
			goto _18
		_18:
			;
			i5 = i5 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(50)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(6):
		i6 = uint32(0)
		for {
			if !(uint64(i6) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token6[i6]) == lookahead1 {
				state = map_token6[i6+uint32(1)]
				goto next_state
			}
			goto _19
		_19:
			;
			i6 = i6 + uint32(2)
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(63)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead1 == int32('\n') {
			state = uint16(56)
			goto next_state
		}
		if lookahead1 == int32('\f') {
			state = uint16(59)
			goto next_state
		}
		if lookahead1 == int32('\r') {
			state = uint16(8)
			goto next_state
		}
		if lookahead1 == int32(0x85) {
			state = uint16(58)
			goto next_state
		}
		if lookahead1 == int32(0x2028) {
			state = uint16(60)
			goto next_state
		}
		if lookahead1 == int32(0x2029) {
			state = uint16(61)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead1 == int32('\n') {
			state = uint16(57)
			goto next_state
		}
		if lookahead1 == int32('\'') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead1 == int32('"') {
			state = uint16(31)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(33)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead1 == int32('-') {
			state = uint16(21)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead1 == int32('{') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead1 == int32('}') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead1 == int32('}') {
			state = uint16(34)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead1 == int32('}') {
			state = uint16(34)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead1 == int32('}') {
			state = uint16(34)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead1 == int32('}') {
			state = uint16(34)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead1 == int32('}') {
			state = uint16(34)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(18):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(19):
		if eof != 0 {
			state = uint16(20)
			goto next_state
		}
		i7 = uint32(0)
		for {
			if !(uint64(i7) < libc.Uint64FromInt64(76)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token7[i7]) == lookahead1 {
				state = map_token7[i7+uint32(1)]
				goto next_state
			}
			goto _20
		_20:
			;
			i7 = i7 + uint32(2)
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(44)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32(' ') || lookahead1 == int32(0xa0) || lookahead1 == int32(0x1680) || int32(0x2000) <= lookahead1 && lookahead1 <= int32(0x200a) || lookahead1 == int32(0x202f) || lookahead1 == int32(0x205f) || lookahead1 == int32(0x3000) {
			state = uint16(63)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym__normal_bare_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(642) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _24
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _24
	_24:
		if v4 != 0 {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(20):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(22):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(23):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(24):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__normal_bare_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym__normal_bare_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(642) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _28
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _28
	_28:
		if v4 != 0 {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_char)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym___identifier_char_no_digit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escaped_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escaped_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i8 = uint32(0)
		for {
			if !(uint64(i8) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token8[i8]) == lookahead1 {
				state = map_token8[i8+uint32(1)]
				goto next_state
			}
			goto _29
		_29:
			;
			i8 = i8 + uint32(2)
		}
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__hex_digit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym__normal_bare_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(642) - index
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
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_e)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_e)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym__normal_bare_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(642) - index
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
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_E)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_E)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym__normal_bare_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(642) - index
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
		if v4 != 0 {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym__)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym__)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym__normal_bare_identifier_character_set_2))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(642) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _45
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _45
	_45:
		if v4 != 0 {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__digit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__digit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('b') {
			state = uint16(51)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(49)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0x)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0o)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__octal_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0b)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__newline_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__newline_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__newline_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__newline_token4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__newline_token5)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__newline_token6)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__newline_token7)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__bom)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__unicode_space)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_single_line_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [50]uint16_t{
	0:  uint16('\n'),
	1:  uint16(32),
	2:  uint16('\f'),
	3:  uint16(32),
	4:  uint16('\r'),
	5:  uint16(32),
	6:  uint16('"'),
	7:  uint16(31),
	8:  uint16('('),
	9:  uint16(29),
	10: uint16(')'),
	11: uint16(30),
	12: uint16('+'),
	13: uint16(46),
	14: uint16('-'),
	15: uint16(47),
	16: uint16('.'),
	17: uint16(37),
	18: uint16('0'),
	19: uint16(52),
	20: uint16('1'),
	21: uint16(53),
	22: uint16(';'),
	23: uint16(24),
	24: uint16('='),
	25: uint16(28),
	26: uint16('E'),
	27: uint16(41),
	28: uint16('\\'),
	29: uint16(54),
	30: uint16('_'),
	31: uint16(43),
	32: uint16('e'),
	33: uint16(39),
	34: uint16('{'),
	35: uint16(22),
	36: uint16('}'),
	37: uint16(23),
	38: uint16(0x85),
	39: uint16(32),
	40: uint16(0x2028),
	41: uint16(32),
	42: uint16(0x2029),
	43: uint16(32),
	44: uint16(0xfeff),
	45: uint16(32),
	46: uint16('8'),
	47: uint16(26),
	48: uint16('9'),
	49: uint16(26),
}

var map_token1 = [28]uint16_t{
	0:  uint16('\n'),
	1:  uint16(56),
	2:  uint16('\f'),
	3:  uint16(59),
	4:  uint16('\r'),
	5:  uint16(8),
	6:  uint16(')'),
	7:  uint16(30),
	8:  uint16('/'),
	9:  uint16(10),
	10: uint16('0'),
	11: uint16(45),
	12: uint16(';'),
	13: uint16(24),
	14: uint16('='),
	15: uint16(28),
	16: uint16('\\'),
	17: uint16(54),
	18: uint16('{'),
	19: uint16(22),
	20: uint16(0x85),
	21: uint16(58),
	22: uint16(0x2028),
	23: uint16(60),
	24: uint16(0x2029),
	25: uint16(61),
	26: uint16(0xfeff),
	27: uint16(62),
}

var map_token2 = [26]uint16_t{
	0:  uint16('\n'),
	1:  uint16(56),
	2:  uint16('\f'),
	3:  uint16(59),
	4:  uint16('\r'),
	5:  uint16(8),
	6:  uint16(')'),
	7:  uint16(30),
	8:  uint16('/'),
	9:  uint16(10),
	10: uint16(';'),
	11: uint16(24),
	12: uint16('='),
	13: uint16(28),
	14: uint16('\\'),
	15: uint16(54),
	16: uint16('{'),
	17: uint16(22),
	18: uint16(0x85),
	19: uint16(58),
	20: uint16(0x2028),
	21: uint16(60),
	22: uint16(0x2029),
	23: uint16(61),
	24: uint16(0xfeff),
	25: uint16(62),
}

var map_token3 = [30]uint16_t{
	0:  uint16('\n'),
	1:  uint16(56),
	2:  uint16('\f'),
	3:  uint16(59),
	4:  uint16('\r'),
	5:  uint16(8),
	6:  uint16('.'),
	7:  uint16(36),
	8:  uint16('/'),
	9:  uint16(10),
	10: uint16(';'),
	11: uint16(24),
	12: uint16('E'),
	13: uint16(40),
	14: uint16('\\'),
	15: uint16(54),
	16: uint16('_'),
	17: uint16(42),
	18: uint16('e'),
	19: uint16(38),
	20: uint16('{'),
	21: uint16(22),
	22: uint16(0x85),
	23: uint16(58),
	24: uint16(0x2028),
	25: uint16(60),
	26: uint16(0x2029),
	27: uint16(61),
	28: uint16(0xfeff),
	29: uint16(62),
}

var map_token4 = [28]uint16_t{
	0:  uint16('\n'),
	1:  uint16(56),
	2:  uint16('\f'),
	3:  uint16(59),
	4:  uint16('\r'),
	5:  uint16(8),
	6:  uint16('/'),
	7:  uint16(10),
	8:  uint16('0'),
	9:  uint16(52),
	10: uint16('1'),
	11: uint16(53),
	12: uint16(';'),
	13: uint16(24),
	14: uint16('\\'),
	15: uint16(54),
	16: uint16('_'),
	17: uint16(42),
	18: uint16('{'),
	19: uint16(22),
	20: uint16(0x85),
	21: uint16(58),
	22: uint16(0x2028),
	23: uint16(60),
	24: uint16(0x2029),
	25: uint16(61),
	26: uint16(0xfeff),
	27: uint16(62),
}

var map_token5 = [24]uint16_t{
	0:  uint16('\n'),
	1:  uint16(56),
	2:  uint16('\f'),
	3:  uint16(59),
	4:  uint16('\r'),
	5:  uint16(8),
	6:  uint16('/'),
	7:  uint16(10),
	8:  uint16(';'),
	9:  uint16(24),
	10: uint16('\\'),
	11: uint16(54),
	12: uint16('_'),
	13: uint16(42),
	14: uint16('{'),
	15: uint16(22),
	16: uint16(0x85),
	17: uint16(58),
	18: uint16(0x2028),
	19: uint16(60),
	20: uint16(0x2029),
	21: uint16(61),
	22: uint16(0xfeff),
	23: uint16(62),
}

var map_token6 = [24]uint16_t{
	0:  uint16('\n'),
	1:  uint16(56),
	2:  uint16('\f'),
	3:  uint16(59),
	4:  uint16('\r'),
	5:  uint16(8),
	6:  uint16('/'),
	7:  uint16(10),
	8:  uint16(';'),
	9:  uint16(24),
	10: uint16('\\'),
	11: uint16(54),
	12: uint16('_'),
	13: uint16(42),
	14: uint16('{'),
	15: uint16(22),
	16: uint16(0x85),
	17: uint16(58),
	18: uint16(0x2028),
	19: uint16(60),
	20: uint16(0x2029),
	21: uint16(61),
	22: uint16(0xfeff),
	23: uint16(62),
}

var map_token7 = [38]uint16_t{
	0:  uint16('\n'),
	1:  uint16(56),
	2:  uint16('\f'),
	3:  uint16(59),
	4:  uint16('\r'),
	5:  uint16(8),
	6:  uint16('"'),
	7:  uint16(31),
	8:  uint16('('),
	9:  uint16(29),
	10: uint16(')'),
	11: uint16(30),
	12: uint16('+'),
	13: uint16(46),
	14: uint16('-'),
	15: uint16(47),
	16: uint16('/'),
	17: uint16(10),
	18: uint16('0'),
	19: uint16(45),
	20: uint16(';'),
	21: uint16(24),
	22: uint16('='),
	23: uint16(28),
	24: uint16('\\'),
	25: uint16(54),
	26: uint16('{'),
	27: uint16(22),
	28: uint16('}'),
	29: uint16(23),
	30: uint16(0x85),
	31: uint16(58),
	32: uint16(0x2028),
	33: uint16(60),
	34: uint16(0x2029),
	35: uint16(61),
	36: uint16(0xfeff),
	37: uint16(62),
}

var map_token8 = [18]uint16_t{
	0:  uint16('u'),
	1:  uint16(11),
	2:  uint16('"'),
	3:  uint16(34),
	4:  uint16('/'),
	5:  uint16(34),
	6:  uint16('\\'),
	7:  uint16(34),
	8:  uint16('b'),
	9:  uint16(34),
	10: uint16('f'),
	11: uint16(34),
	12: uint16('n'),
	13: uint16(34),
	14: uint16('r'),
	15: uint16(34),
	16: uint16('t'),
	17: uint16(34),
}

func ts_lex_keywords(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
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
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token9[i]) == lookahead {
				state = map_token9[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		return result
	case int32(1):
		if lookahead == int32('a') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('o') {
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('a') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('m') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('3') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('o') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(7):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token11[i1]) == lookahead {
				state = map_token11[i1+uint32(1)]
				goto next_state
			}
			goto _2
		_2:
			;
			i1 = i1 + uint32(2)
		}
		return result
	case int32(8):
		if lookahead == int32('u') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('e') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('i') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('1') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('s') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('u') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('r') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('t') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('c') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('r') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('a') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('2') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('4') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('l') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('s') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('6') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('2') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('4') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_i8)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		if lookahead == int32('n') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('v') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('i') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('i') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('l') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('g') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('m') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('u') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('6') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('2') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('4') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u8)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		if lookahead == int32('l') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('i') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('i') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('e') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('n') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('r') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('e') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('i') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('a') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('i') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_f32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_f64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(51):
		if lookahead == int32('s') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('t') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_i16)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_i32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_i64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(56):
		if lookahead == int32('-') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('4') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('-') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_irl)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(60):
		if lookahead == int32('z') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('l') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('e') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('e') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('e') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u16)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_url)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('z') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('d') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('6') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('t') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('e') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_date)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('m') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('t') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('l') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('e') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead == int32('n') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead == int32('e') {
			state = uint16(101)
			goto next_state
		}
		if lookahead == int32('h') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ipv4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ipv6)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		if lookahead == int32('r') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('e') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_null)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(86):
		if lookahead == int32('x') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_time)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		if lookahead == int32('r') {
			state = uint16(106)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('e') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_uuid)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(92):
		if lookahead == int32('4') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('r') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead == int32('n') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead == int32('t') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead == int32('a') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead == int32('i') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_email)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(100):
		if lookahead == int32('a') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('m') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead == int32('o') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead == int32('e') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_isize)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_regex)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(106):
		if lookahead == int32('e') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead == int32('e') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_usize)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_base64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(110):
		if lookahead == int32('y') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead == int32('c') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(112):
		if lookahead == int32('i') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead == int32('l') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('o') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('m') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead == int32('a') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead == int32('s') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead == int32('f') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('f') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('m') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('-') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('y') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead == int32('m') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_decimal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('1') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead == int32('n') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead == int32('e') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead == int32('i') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead == int32('t') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(129):
		if lookahead == int32('e') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(130):
		if lookahead == int32('e') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(131):
		if lookahead == int32('p') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(132):
		if lookahead == int32('2') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(145)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_currency)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(134):
		if lookahead == int32('e') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead == int32('2') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(136):
		if lookahead == int32('4') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_duration)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(138):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_hostname)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(139):
		if lookahead == int32('l') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(140):
		if lookahead == int32('n') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead == int32('r') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(142):
		if lookahead == int32('r') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(143):
		if lookahead == int32('l') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(144):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_country_DASH2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_country_DASH3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(146):
		if lookahead == int32('u') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(147):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_date_DASHtime)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(148):
		if lookahead == int32('8') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(149):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_decimal64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(150):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_idn_DASHemail)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(151):
		if lookahead == int32('a') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(152):
		if lookahead == int32('e') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead == int32('e') {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(154):
		if lookahead == int32('a') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(155):
		if lookahead == int32('b') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(156):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_decimal128)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(157):
		if lookahead == int32('m') {
			state = uint16(162)
			goto next_state
		}
		return result
	case int32(158):
		if lookahead == int32('n') {
			state = uint16(163)
			goto next_state
		}
		return result
	case int32(159):
		if lookahead == int32('n') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(160):
		if lookahead == int32('t') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(161):
		if lookahead == int32('d') {
			state = uint16(166)
			goto next_state
		}
		return result
	case int32(162):
		if lookahead == int32('e') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(163):
		if lookahead == int32('c') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(164):
		if lookahead == int32('c') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(165):
		if lookahead == int32('e') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(166):
		if lookahead == int32('i') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(167):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_idn_DASHhostname)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(168):
		if lookahead == int32('e') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(169):
		if lookahead == int32('e') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(170):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_url_DASHtemplate)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(171):
		if lookahead == int32('v') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(172):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_iri_DASHreference)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(173):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_url_DASHreference)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(174):
		if lookahead == int32('i') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(175):
		if lookahead == int32('s') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(176):
		if lookahead == int32('i') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(177):
		if lookahead == int32('o') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(178):
		if lookahead == int32('n') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(179):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_country_DASHsubdivision)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token9 = [22]uint16_t{
	0:  uint16('b'),
	1:  uint16(1),
	2:  uint16('c'),
	3:  uint16(2),
	4:  uint16('d'),
	5:  uint16(3),
	6:  uint16('e'),
	7:  uint16(4),
	8:  uint16('f'),
	9:  uint16(5),
	10: uint16('h'),
	11: uint16(6),
	12: uint16('i'),
	13: uint16(7),
	14: uint16('n'),
	15: uint16(8),
	16: uint16('r'),
	17: uint16(9),
	18: uint16('t'),
	19: uint16(10),
	20: uint16('u'),
	21: uint16(11),
}

var map_token11 = [16]uint16_t{
	0:  uint16('1'),
	1:  uint16(23),
	2:  uint16('3'),
	3:  uint16(24),
	4:  uint16('6'),
	5:  uint16(25),
	6:  uint16('8'),
	7:  uint16(26),
	8:  uint16('d'),
	9:  uint16(27),
	10: uint16('p'),
	11: uint16(28),
	12: uint16('r'),
	13: uint16(29),
	14: uint16('s'),
	15: uint16(30),
}

var ts_lex_modes = [304]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	2: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	3: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	4: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	5: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	6: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	7: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	8: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	9: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	10: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	11: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	12: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	13: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	14: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	15: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	16: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	17: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	18: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	19: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	20: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	21: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	22: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	23: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	24: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	25: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	26: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	27: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	28: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	29: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	30: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	31: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	32: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	33: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	34: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	35: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	36: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	37: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	38: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	39: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	40: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	41: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	42: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	43: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	44: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	45: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	46: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	47: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	48: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	49: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	50: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	51: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	52: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	53: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	54: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	55: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	56: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	57: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	58: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	59: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	60: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	61: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(1),
	},
	62: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	63: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	64: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	65: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	66: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	67: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	68: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	69: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	70: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	71: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	72: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	73: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	74: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	75: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	76: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	77: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	78: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	79: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	80: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	81: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(3),
	},
	82: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	83: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	84: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	85: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	86: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	87: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	88: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	89: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	90: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	91: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	92: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	93: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	94: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	95: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	96: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	97: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	98: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	99: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(3),
	},
	100: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	101: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	102: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	103: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	104: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	105: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	106: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	107: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	108: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	109: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	110: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	111: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	112: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	113: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	114: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	115: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	116: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	117: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	118: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	119: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(3),
	},
	120: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	121: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	122: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	123: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	124: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	125: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	126: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	127: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	128: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	129: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	130: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	131: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	132: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	133: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	134: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	135: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(3),
	},
	136: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	137: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	138: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	139: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	140: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	141: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	142: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	143: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	144: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	145: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	146: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	147: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	148: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	149: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	150: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	151: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	152: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	153: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	154: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(3),
	},
	155: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	156: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	157: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	158: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	159: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	160: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	161: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	162: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	163: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	164: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	165: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	166: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	167: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	168: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	169: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	170: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	171: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	172: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	173: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	174: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(3),
	},
	175: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	176: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	177: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	178: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	179: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	180: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(3),
	},
	181: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	182: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	183: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	184: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	185: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	186: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	187: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	188: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	189: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	190: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	191: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	192: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(3),
	},
	193: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	194: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	195: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	196: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	197: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	198: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	199: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	200: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	201: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	202: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	203: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	204: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	205: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	206: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	207: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	208: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	209: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	210: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	211: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	212: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	213: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	214: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	215: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	216: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	217: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	218: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	219: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	220: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	221: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	222: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	223: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	224: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	225: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	226: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	227: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	228: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	229: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	230: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	231: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	232: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	233: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	234: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	235: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	236: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	237: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	238: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	239: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	240: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	241: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	242: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	243: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	244: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	245: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	246: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	247: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	248: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	249: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(3),
	},
	250: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(3),
	},
	251: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(3),
	},
	252: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(3),
	},
	253: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(3),
	},
	254: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(3),
	},
	255: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(3),
	},
	256: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(3),
	},
	257: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	258: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	259: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(3),
	},
	260: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	261: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	262: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	263: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	264: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	265: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	266: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	267: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	268: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	269: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	270: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	271: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	272: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	273: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	274: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	275: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	276: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	277: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	278: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	279: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	280: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	281: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	282: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	283: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	284: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	285: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	286: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	287: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	288: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(4),
	},
	289: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(4),
	},
	290: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	291: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	292: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	293: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	294: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	295: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	296: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	297: {
		Fexternal_lex_state: uint16(4),
	},
	298: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	299: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	300: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	301: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
	302: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	303: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(4),
	},
}

var ts_parse_table = [2][127]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
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
		53: uint16(1),
		54: uint16(1),
		55: uint16(1),
		56: uint16(1),
		57: uint16(1),
		58: uint16(1),
		59: uint16(1),
		60: uint16(1),
		63: uint16(1),
		65: uint16(1),
		66: uint16(1),
		67: uint16(1),
		68: uint16(1),
		69: uint16(1),
		71: uint16(1),
		73: uint16(1),
		74: uint16(1),
		75: uint16(1),
		76: uint16(1),
		77: uint16(1),
		78: uint16(1),
		80: uint16(1),
		81: uint16(1),
		82: uint16(3),
		83: uint16(1),
	},
	1: {
		0:   uint16(5),
		1:   uint16(7),
		2:   uint16(9),
		48:  uint16(11),
		50:  uint16(13),
		59:  uint16(15),
		60:  uint16(15),
		70:  uint16(17),
		71:  uint16(17),
		72:  uint16(17),
		73:  uint16(17),
		74:  uint16(17),
		75:  uint16(17),
		76:  uint16(17),
		77:  uint16(17),
		78:  uint16(17),
		79:  uint16(19),
		82:  uint16(17),
		83:  uint16(21),
		84:  uint16(297),
		85:  uint16(17),
		92:  uint16(71),
		93:  uint16(218),
		98:  uint16(260),
		99:  uint16(218),
		100: uint16(194),
		105: uint16(192),
		111: uint16(45),
		112: uint16(45),
		113: uint16(45),
		114: uint16(45),
		115: uint16(45),
	},
}

var ts_small_parse_table = [10444]uint16_t{
	0:     uint16(36),
	1:     uint16(11),
	2:     uint16(1),
	3:     uint16(anon_sym_LPAREN),
	4:     uint16(13),
	5:     uint16(1),
	6:     uint16(anon_sym_DQUOTE),
	7:     uint16(19),
	8:     uint16(1),
	9:     uint16(anon_sym_SLASH_SLASH),
	10:    uint16(21),
	11:    uint16(1),
	12:    uint16(sym__raw_string),
	13:    uint16(23),
	14:    uint16(1),
	15:    uint16(sym__normal_bare_identifier),
	16:    uint16(25),
	17:    uint16(1),
	18:    uint16(anon_sym_SLASH_DASH),
	19:    uint16(27),
	20:    uint16(1),
	21:    uint16(anon_sym_LBRACE),
	22:    uint16(31),
	23:    uint16(1),
	24:    uint16(anon_sym_null),
	25:    uint16(33),
	26:    uint16(1),
	27:    uint16(sym__digit),
	28:    uint16(37),
	29:    uint16(1),
	30:    uint16(anon_sym_0x),
	31:    uint16(39),
	32:    uint16(1),
	33:    uint16(anon_sym_0o),
	34:    uint16(41),
	35:    uint16(1),
	36:    uint16(anon_sym_0b),
	37:    uint16(45),
	38:    uint16(1),
	39:    uint16(anon_sym_BSLASH),
	40:    uint16(15),
	41:    uint16(1),
	42:    uint16(aux_sym_node_repeat1),
	43:    uint16(34),
	44:    uint16(1),
	45:    uint16(sym__escline),
	46:    uint16(58),
	47:    uint16(1),
	48:    uint16(sym__node_space),
	49:    uint16(79),
	50:    uint16(1),
	51:    uint16(sym_type),
	52:    uint16(109),
	53:    uint16(1),
	54:    uint16(sym_node_children),
	55:    uint16(154),
	56:    uint16(1),
	57:    uint16(sym__integer),
	58:    uint16(191),
	59:    uint16(1),
	60:    uint16(sym_string),
	61:    uint16(194),
	62:    uint16(1),
	63:    uint16(sym__escaped_string),
	64:    uint16(198),
	65:    uint16(1),
	66:    uint16(sym_node_field),
	67:    uint16(217),
	68:    uint16(1),
	69:    uint16(sym_boolean),
	70:    uint16(264),
	71:    uint16(1),
	72:    uint16(sym__sign),
	73:    uint16(290),
	74:    uint16(1),
	75:    uint16(sym__bare_identifier),
	76:    uint16(295),
	77:    uint16(1),
	78:    uint16(sym_identifier),
	79:    uint16(35),
	80:    uint16(2),
	81:    uint16(anon_sym_PLUS),
	82:    uint16(anon_sym_DASH),
	83:    uint16(43),
	84:    uint16(2),
	85:    uint16(anon_sym_true),
	86:    uint16(anon_sym_false),
	87:    uint16(27),
	88:    uint16(2),
	89:    uint16(sym__ws),
	90:    uint16(aux_sym_node_repeat3),
	91:    uint16(199),
	92:    uint16(2),
	93:    uint16(sym_keyword),
	94:    uint16(sym_number),
	95:    uint16(211),
	96:    uint16(2),
	97:    uint16(sym__node_field_comment),
	98:    uint16(sym__node_field),
	99:    uint16(213),
	100:   uint16(2),
	101:   uint16(sym_prop),
	102:   uint16(sym_value),
	103:   uint16(47),
	104:   uint16(3),
	105:   uint16(sym_multi_line_comment),
	106:   uint16(sym__bom),
	107:   uint16(sym__unicode_space),
	108:   uint16(152),
	109:   uint16(3),
	110:   uint16(sym__node_terminator),
	111:   uint16(sym__newline),
	112:   uint16(sym_single_line_comment),
	113:   uint16(214),
	114:   uint16(4),
	115:   uint16(sym__decimal),
	116:   uint16(sym__hex),
	117:   uint16(sym__octal),
	118:   uint16(sym__binary),
	119:   uint16(29),
	120:   uint16(9),
	121:   uint16(sym__eof),
	122:   uint16(anon_sym_SEMI),
	123:   uint16(aux_sym__newline_token1),
	124:   uint16(aux_sym__newline_token2),
	125:   uint16(aux_sym__newline_token3),
	126:   uint16(aux_sym__newline_token4),
	127:   uint16(aux_sym__newline_token5),
	128:   uint16(aux_sym__newline_token6),
	129:   uint16(aux_sym__newline_token7),
	130:   uint16(36),
	131:   uint16(11),
	132:   uint16(1),
	133:   uint16(anon_sym_LPAREN),
	134:   uint16(13),
	135:   uint16(1),
	136:   uint16(anon_sym_DQUOTE),
	137:   uint16(19),
	138:   uint16(1),
	139:   uint16(anon_sym_SLASH_SLASH),
	140:   uint16(21),
	141:   uint16(1),
	142:   uint16(sym__raw_string),
	143:   uint16(23),
	144:   uint16(1),
	145:   uint16(sym__normal_bare_identifier),
	146:   uint16(25),
	147:   uint16(1),
	148:   uint16(anon_sym_SLASH_DASH),
	149:   uint16(27),
	150:   uint16(1),
	151:   uint16(anon_sym_LBRACE),
	152:   uint16(31),
	153:   uint16(1),
	154:   uint16(anon_sym_null),
	155:   uint16(33),
	156:   uint16(1),
	157:   uint16(sym__digit),
	158:   uint16(37),
	159:   uint16(1),
	160:   uint16(anon_sym_0x),
	161:   uint16(39),
	162:   uint16(1),
	163:   uint16(anon_sym_0o),
	164:   uint16(41),
	165:   uint16(1),
	166:   uint16(anon_sym_0b),
	167:   uint16(45),
	168:   uint16(1),
	169:   uint16(anon_sym_BSLASH),
	170:   uint16(15),
	171:   uint16(1),
	172:   uint16(aux_sym_node_repeat1),
	173:   uint16(34),
	174:   uint16(1),
	175:   uint16(sym__escline),
	176:   uint16(58),
	177:   uint16(1),
	178:   uint16(sym__node_space),
	179:   uint16(79),
	180:   uint16(1),
	181:   uint16(sym_type),
	182:   uint16(118),
	183:   uint16(1),
	184:   uint16(sym_node_children),
	185:   uint16(154),
	186:   uint16(1),
	187:   uint16(sym__integer),
	188:   uint16(191),
	189:   uint16(1),
	190:   uint16(sym_string),
	191:   uint16(194),
	192:   uint16(1),
	193:   uint16(sym__escaped_string),
	194:   uint16(198),
	195:   uint16(1),
	196:   uint16(sym_node_field),
	197:   uint16(217),
	198:   uint16(1),
	199:   uint16(sym_boolean),
	200:   uint16(264),
	201:   uint16(1),
	202:   uint16(sym__sign),
	203:   uint16(290),
	204:   uint16(1),
	205:   uint16(sym__bare_identifier),
	206:   uint16(295),
	207:   uint16(1),
	208:   uint16(sym_identifier),
	209:   uint16(35),
	210:   uint16(2),
	211:   uint16(anon_sym_PLUS),
	212:   uint16(anon_sym_DASH),
	213:   uint16(43),
	214:   uint16(2),
	215:   uint16(anon_sym_true),
	216:   uint16(anon_sym_false),
	217:   uint16(27),
	218:   uint16(2),
	219:   uint16(sym__ws),
	220:   uint16(aux_sym_node_repeat3),
	221:   uint16(199),
	222:   uint16(2),
	223:   uint16(sym_keyword),
	224:   uint16(sym_number),
	225:   uint16(211),
	226:   uint16(2),
	227:   uint16(sym__node_field_comment),
	228:   uint16(sym__node_field),
	229:   uint16(213),
	230:   uint16(2),
	231:   uint16(sym_prop),
	232:   uint16(sym_value),
	233:   uint16(47),
	234:   uint16(3),
	235:   uint16(sym_multi_line_comment),
	236:   uint16(sym__bom),
	237:   uint16(sym__unicode_space),
	238:   uint16(142),
	239:   uint16(3),
	240:   uint16(sym__node_terminator),
	241:   uint16(sym__newline),
	242:   uint16(sym_single_line_comment),
	243:   uint16(214),
	244:   uint16(4),
	245:   uint16(sym__decimal),
	246:   uint16(sym__hex),
	247:   uint16(sym__octal),
	248:   uint16(sym__binary),
	249:   uint16(49),
	250:   uint16(9),
	251:   uint16(sym__eof),
	252:   uint16(anon_sym_SEMI),
	253:   uint16(aux_sym__newline_token1),
	254:   uint16(aux_sym__newline_token2),
	255:   uint16(aux_sym__newline_token3),
	256:   uint16(aux_sym__newline_token4),
	257:   uint16(aux_sym__newline_token5),
	258:   uint16(aux_sym__newline_token6),
	259:   uint16(aux_sym__newline_token7),
	260:   uint16(36),
	261:   uint16(11),
	262:   uint16(1),
	263:   uint16(anon_sym_LPAREN),
	264:   uint16(13),
	265:   uint16(1),
	266:   uint16(anon_sym_DQUOTE),
	267:   uint16(19),
	268:   uint16(1),
	269:   uint16(anon_sym_SLASH_SLASH),
	270:   uint16(21),
	271:   uint16(1),
	272:   uint16(sym__raw_string),
	273:   uint16(23),
	274:   uint16(1),
	275:   uint16(sym__normal_bare_identifier),
	276:   uint16(25),
	277:   uint16(1),
	278:   uint16(anon_sym_SLASH_DASH),
	279:   uint16(27),
	280:   uint16(1),
	281:   uint16(anon_sym_LBRACE),
	282:   uint16(31),
	283:   uint16(1),
	284:   uint16(anon_sym_null),
	285:   uint16(33),
	286:   uint16(1),
	287:   uint16(sym__digit),
	288:   uint16(37),
	289:   uint16(1),
	290:   uint16(anon_sym_0x),
	291:   uint16(39),
	292:   uint16(1),
	293:   uint16(anon_sym_0o),
	294:   uint16(41),
	295:   uint16(1),
	296:   uint16(anon_sym_0b),
	297:   uint16(45),
	298:   uint16(1),
	299:   uint16(anon_sym_BSLASH),
	300:   uint16(15),
	301:   uint16(1),
	302:   uint16(aux_sym_node_repeat1),
	303:   uint16(34),
	304:   uint16(1),
	305:   uint16(sym__escline),
	306:   uint16(58),
	307:   uint16(1),
	308:   uint16(sym__node_space),
	309:   uint16(79),
	310:   uint16(1),
	311:   uint16(sym_type),
	312:   uint16(95),
	313:   uint16(1),
	314:   uint16(sym_node_children),
	315:   uint16(154),
	316:   uint16(1),
	317:   uint16(sym__integer),
	318:   uint16(191),
	319:   uint16(1),
	320:   uint16(sym_string),
	321:   uint16(194),
	322:   uint16(1),
	323:   uint16(sym__escaped_string),
	324:   uint16(198),
	325:   uint16(1),
	326:   uint16(sym_node_field),
	327:   uint16(217),
	328:   uint16(1),
	329:   uint16(sym_boolean),
	330:   uint16(264),
	331:   uint16(1),
	332:   uint16(sym__sign),
	333:   uint16(290),
	334:   uint16(1),
	335:   uint16(sym__bare_identifier),
	336:   uint16(295),
	337:   uint16(1),
	338:   uint16(sym_identifier),
	339:   uint16(35),
	340:   uint16(2),
	341:   uint16(anon_sym_PLUS),
	342:   uint16(anon_sym_DASH),
	343:   uint16(43),
	344:   uint16(2),
	345:   uint16(anon_sym_true),
	346:   uint16(anon_sym_false),
	347:   uint16(27),
	348:   uint16(2),
	349:   uint16(sym__ws),
	350:   uint16(aux_sym_node_repeat3),
	351:   uint16(199),
	352:   uint16(2),
	353:   uint16(sym_keyword),
	354:   uint16(sym_number),
	355:   uint16(211),
	356:   uint16(2),
	357:   uint16(sym__node_field_comment),
	358:   uint16(sym__node_field),
	359:   uint16(213),
	360:   uint16(2),
	361:   uint16(sym_prop),
	362:   uint16(sym_value),
	363:   uint16(47),
	364:   uint16(3),
	365:   uint16(sym_multi_line_comment),
	366:   uint16(sym__bom),
	367:   uint16(sym__unicode_space),
	368:   uint16(171),
	369:   uint16(3),
	370:   uint16(sym__node_terminator),
	371:   uint16(sym__newline),
	372:   uint16(sym_single_line_comment),
	373:   uint16(214),
	374:   uint16(4),
	375:   uint16(sym__decimal),
	376:   uint16(sym__hex),
	377:   uint16(sym__octal),
	378:   uint16(sym__binary),
	379:   uint16(51),
	380:   uint16(9),
	381:   uint16(sym__eof),
	382:   uint16(anon_sym_SEMI),
	383:   uint16(aux_sym__newline_token1),
	384:   uint16(aux_sym__newline_token2),
	385:   uint16(aux_sym__newline_token3),
	386:   uint16(aux_sym__newline_token4),
	387:   uint16(aux_sym__newline_token5),
	388:   uint16(aux_sym__newline_token6),
	389:   uint16(aux_sym__newline_token7),
	390:   uint16(36),
	391:   uint16(11),
	392:   uint16(1),
	393:   uint16(anon_sym_LPAREN),
	394:   uint16(13),
	395:   uint16(1),
	396:   uint16(anon_sym_DQUOTE),
	397:   uint16(19),
	398:   uint16(1),
	399:   uint16(anon_sym_SLASH_SLASH),
	400:   uint16(21),
	401:   uint16(1),
	402:   uint16(sym__raw_string),
	403:   uint16(23),
	404:   uint16(1),
	405:   uint16(sym__normal_bare_identifier),
	406:   uint16(25),
	407:   uint16(1),
	408:   uint16(anon_sym_SLASH_DASH),
	409:   uint16(27),
	410:   uint16(1),
	411:   uint16(anon_sym_LBRACE),
	412:   uint16(31),
	413:   uint16(1),
	414:   uint16(anon_sym_null),
	415:   uint16(33),
	416:   uint16(1),
	417:   uint16(sym__digit),
	418:   uint16(37),
	419:   uint16(1),
	420:   uint16(anon_sym_0x),
	421:   uint16(39),
	422:   uint16(1),
	423:   uint16(anon_sym_0o),
	424:   uint16(41),
	425:   uint16(1),
	426:   uint16(anon_sym_0b),
	427:   uint16(45),
	428:   uint16(1),
	429:   uint16(anon_sym_BSLASH),
	430:   uint16(15),
	431:   uint16(1),
	432:   uint16(aux_sym_node_repeat1),
	433:   uint16(34),
	434:   uint16(1),
	435:   uint16(sym__escline),
	436:   uint16(58),
	437:   uint16(1),
	438:   uint16(sym__node_space),
	439:   uint16(79),
	440:   uint16(1),
	441:   uint16(sym_type),
	442:   uint16(85),
	443:   uint16(1),
	444:   uint16(sym_node_children),
	445:   uint16(154),
	446:   uint16(1),
	447:   uint16(sym__integer),
	448:   uint16(191),
	449:   uint16(1),
	450:   uint16(sym_string),
	451:   uint16(194),
	452:   uint16(1),
	453:   uint16(sym__escaped_string),
	454:   uint16(198),
	455:   uint16(1),
	456:   uint16(sym_node_field),
	457:   uint16(217),
	458:   uint16(1),
	459:   uint16(sym_boolean),
	460:   uint16(264),
	461:   uint16(1),
	462:   uint16(sym__sign),
	463:   uint16(290),
	464:   uint16(1),
	465:   uint16(sym__bare_identifier),
	466:   uint16(295),
	467:   uint16(1),
	468:   uint16(sym_identifier),
	469:   uint16(35),
	470:   uint16(2),
	471:   uint16(anon_sym_PLUS),
	472:   uint16(anon_sym_DASH),
	473:   uint16(43),
	474:   uint16(2),
	475:   uint16(anon_sym_true),
	476:   uint16(anon_sym_false),
	477:   uint16(27),
	478:   uint16(2),
	479:   uint16(sym__ws),
	480:   uint16(aux_sym_node_repeat3),
	481:   uint16(199),
	482:   uint16(2),
	483:   uint16(sym_keyword),
	484:   uint16(sym_number),
	485:   uint16(211),
	486:   uint16(2),
	487:   uint16(sym__node_field_comment),
	488:   uint16(sym__node_field),
	489:   uint16(213),
	490:   uint16(2),
	491:   uint16(sym_prop),
	492:   uint16(sym_value),
	493:   uint16(47),
	494:   uint16(3),
	495:   uint16(sym_multi_line_comment),
	496:   uint16(sym__bom),
	497:   uint16(sym__unicode_space),
	498:   uint16(141),
	499:   uint16(3),
	500:   uint16(sym__node_terminator),
	501:   uint16(sym__newline),
	502:   uint16(sym_single_line_comment),
	503:   uint16(214),
	504:   uint16(4),
	505:   uint16(sym__decimal),
	506:   uint16(sym__hex),
	507:   uint16(sym__octal),
	508:   uint16(sym__binary),
	509:   uint16(53),
	510:   uint16(9),
	511:   uint16(sym__eof),
	512:   uint16(anon_sym_SEMI),
	513:   uint16(aux_sym__newline_token1),
	514:   uint16(aux_sym__newline_token2),
	515:   uint16(aux_sym__newline_token3),
	516:   uint16(aux_sym__newline_token4),
	517:   uint16(aux_sym__newline_token5),
	518:   uint16(aux_sym__newline_token6),
	519:   uint16(aux_sym__newline_token7),
	520:   uint16(36),
	521:   uint16(11),
	522:   uint16(1),
	523:   uint16(anon_sym_LPAREN),
	524:   uint16(13),
	525:   uint16(1),
	526:   uint16(anon_sym_DQUOTE),
	527:   uint16(19),
	528:   uint16(1),
	529:   uint16(anon_sym_SLASH_SLASH),
	530:   uint16(21),
	531:   uint16(1),
	532:   uint16(sym__raw_string),
	533:   uint16(23),
	534:   uint16(1),
	535:   uint16(sym__normal_bare_identifier),
	536:   uint16(25),
	537:   uint16(1),
	538:   uint16(anon_sym_SLASH_DASH),
	539:   uint16(27),
	540:   uint16(1),
	541:   uint16(anon_sym_LBRACE),
	542:   uint16(31),
	543:   uint16(1),
	544:   uint16(anon_sym_null),
	545:   uint16(33),
	546:   uint16(1),
	547:   uint16(sym__digit),
	548:   uint16(37),
	549:   uint16(1),
	550:   uint16(anon_sym_0x),
	551:   uint16(39),
	552:   uint16(1),
	553:   uint16(anon_sym_0o),
	554:   uint16(41),
	555:   uint16(1),
	556:   uint16(anon_sym_0b),
	557:   uint16(45),
	558:   uint16(1),
	559:   uint16(anon_sym_BSLASH),
	560:   uint16(15),
	561:   uint16(1),
	562:   uint16(aux_sym_node_repeat1),
	563:   uint16(34),
	564:   uint16(1),
	565:   uint16(sym__escline),
	566:   uint16(58),
	567:   uint16(1),
	568:   uint16(sym__node_space),
	569:   uint16(79),
	570:   uint16(1),
	571:   uint16(sym_type),
	572:   uint16(90),
	573:   uint16(1),
	574:   uint16(sym_node_children),
	575:   uint16(154),
	576:   uint16(1),
	577:   uint16(sym__integer),
	578:   uint16(191),
	579:   uint16(1),
	580:   uint16(sym_string),
	581:   uint16(194),
	582:   uint16(1),
	583:   uint16(sym__escaped_string),
	584:   uint16(198),
	585:   uint16(1),
	586:   uint16(sym_node_field),
	587:   uint16(217),
	588:   uint16(1),
	589:   uint16(sym_boolean),
	590:   uint16(264),
	591:   uint16(1),
	592:   uint16(sym__sign),
	593:   uint16(290),
	594:   uint16(1),
	595:   uint16(sym__bare_identifier),
	596:   uint16(295),
	597:   uint16(1),
	598:   uint16(sym_identifier),
	599:   uint16(35),
	600:   uint16(2),
	601:   uint16(anon_sym_PLUS),
	602:   uint16(anon_sym_DASH),
	603:   uint16(43),
	604:   uint16(2),
	605:   uint16(anon_sym_true),
	606:   uint16(anon_sym_false),
	607:   uint16(27),
	608:   uint16(2),
	609:   uint16(sym__ws),
	610:   uint16(aux_sym_node_repeat3),
	611:   uint16(199),
	612:   uint16(2),
	613:   uint16(sym_keyword),
	614:   uint16(sym_number),
	615:   uint16(211),
	616:   uint16(2),
	617:   uint16(sym__node_field_comment),
	618:   uint16(sym__node_field),
	619:   uint16(213),
	620:   uint16(2),
	621:   uint16(sym_prop),
	622:   uint16(sym_value),
	623:   uint16(47),
	624:   uint16(3),
	625:   uint16(sym_multi_line_comment),
	626:   uint16(sym__bom),
	627:   uint16(sym__unicode_space),
	628:   uint16(162),
	629:   uint16(3),
	630:   uint16(sym__node_terminator),
	631:   uint16(sym__newline),
	632:   uint16(sym_single_line_comment),
	633:   uint16(214),
	634:   uint16(4),
	635:   uint16(sym__decimal),
	636:   uint16(sym__hex),
	637:   uint16(sym__octal),
	638:   uint16(sym__binary),
	639:   uint16(55),
	640:   uint16(9),
	641:   uint16(sym__eof),
	642:   uint16(anon_sym_SEMI),
	643:   uint16(aux_sym__newline_token1),
	644:   uint16(aux_sym__newline_token2),
	645:   uint16(aux_sym__newline_token3),
	646:   uint16(aux_sym__newline_token4),
	647:   uint16(aux_sym__newline_token5),
	648:   uint16(aux_sym__newline_token6),
	649:   uint16(aux_sym__newline_token7),
	650:   uint16(36),
	651:   uint16(11),
	652:   uint16(1),
	653:   uint16(anon_sym_LPAREN),
	654:   uint16(13),
	655:   uint16(1),
	656:   uint16(anon_sym_DQUOTE),
	657:   uint16(19),
	658:   uint16(1),
	659:   uint16(anon_sym_SLASH_SLASH),
	660:   uint16(21),
	661:   uint16(1),
	662:   uint16(sym__raw_string),
	663:   uint16(23),
	664:   uint16(1),
	665:   uint16(sym__normal_bare_identifier),
	666:   uint16(25),
	667:   uint16(1),
	668:   uint16(anon_sym_SLASH_DASH),
	669:   uint16(27),
	670:   uint16(1),
	671:   uint16(anon_sym_LBRACE),
	672:   uint16(31),
	673:   uint16(1),
	674:   uint16(anon_sym_null),
	675:   uint16(33),
	676:   uint16(1),
	677:   uint16(sym__digit),
	678:   uint16(37),
	679:   uint16(1),
	680:   uint16(anon_sym_0x),
	681:   uint16(39),
	682:   uint16(1),
	683:   uint16(anon_sym_0o),
	684:   uint16(41),
	685:   uint16(1),
	686:   uint16(anon_sym_0b),
	687:   uint16(45),
	688:   uint16(1),
	689:   uint16(anon_sym_BSLASH),
	690:   uint16(15),
	691:   uint16(1),
	692:   uint16(aux_sym_node_repeat1),
	693:   uint16(34),
	694:   uint16(1),
	695:   uint16(sym__escline),
	696:   uint16(58),
	697:   uint16(1),
	698:   uint16(sym__node_space),
	699:   uint16(79),
	700:   uint16(1),
	701:   uint16(sym_type),
	702:   uint16(80),
	703:   uint16(1),
	704:   uint16(sym_node_children),
	705:   uint16(154),
	706:   uint16(1),
	707:   uint16(sym__integer),
	708:   uint16(191),
	709:   uint16(1),
	710:   uint16(sym_string),
	711:   uint16(194),
	712:   uint16(1),
	713:   uint16(sym__escaped_string),
	714:   uint16(198),
	715:   uint16(1),
	716:   uint16(sym_node_field),
	717:   uint16(217),
	718:   uint16(1),
	719:   uint16(sym_boolean),
	720:   uint16(264),
	721:   uint16(1),
	722:   uint16(sym__sign),
	723:   uint16(290),
	724:   uint16(1),
	725:   uint16(sym__bare_identifier),
	726:   uint16(295),
	727:   uint16(1),
	728:   uint16(sym_identifier),
	729:   uint16(35),
	730:   uint16(2),
	731:   uint16(anon_sym_PLUS),
	732:   uint16(anon_sym_DASH),
	733:   uint16(43),
	734:   uint16(2),
	735:   uint16(anon_sym_true),
	736:   uint16(anon_sym_false),
	737:   uint16(27),
	738:   uint16(2),
	739:   uint16(sym__ws),
	740:   uint16(aux_sym_node_repeat3),
	741:   uint16(199),
	742:   uint16(2),
	743:   uint16(sym_keyword),
	744:   uint16(sym_number),
	745:   uint16(211),
	746:   uint16(2),
	747:   uint16(sym__node_field_comment),
	748:   uint16(sym__node_field),
	749:   uint16(213),
	750:   uint16(2),
	751:   uint16(sym_prop),
	752:   uint16(sym_value),
	753:   uint16(47),
	754:   uint16(3),
	755:   uint16(sym_multi_line_comment),
	756:   uint16(sym__bom),
	757:   uint16(sym__unicode_space),
	758:   uint16(148),
	759:   uint16(3),
	760:   uint16(sym__node_terminator),
	761:   uint16(sym__newline),
	762:   uint16(sym_single_line_comment),
	763:   uint16(214),
	764:   uint16(4),
	765:   uint16(sym__decimal),
	766:   uint16(sym__hex),
	767:   uint16(sym__octal),
	768:   uint16(sym__binary),
	769:   uint16(57),
	770:   uint16(9),
	771:   uint16(sym__eof),
	772:   uint16(anon_sym_SEMI),
	773:   uint16(aux_sym__newline_token1),
	774:   uint16(aux_sym__newline_token2),
	775:   uint16(aux_sym__newline_token3),
	776:   uint16(aux_sym__newline_token4),
	777:   uint16(aux_sym__newline_token5),
	778:   uint16(aux_sym__newline_token6),
	779:   uint16(aux_sym__newline_token7),
	780:   uint16(36),
	781:   uint16(11),
	782:   uint16(1),
	783:   uint16(anon_sym_LPAREN),
	784:   uint16(13),
	785:   uint16(1),
	786:   uint16(anon_sym_DQUOTE),
	787:   uint16(19),
	788:   uint16(1),
	789:   uint16(anon_sym_SLASH_SLASH),
	790:   uint16(21),
	791:   uint16(1),
	792:   uint16(sym__raw_string),
	793:   uint16(23),
	794:   uint16(1),
	795:   uint16(sym__normal_bare_identifier),
	796:   uint16(25),
	797:   uint16(1),
	798:   uint16(anon_sym_SLASH_DASH),
	799:   uint16(27),
	800:   uint16(1),
	801:   uint16(anon_sym_LBRACE),
	802:   uint16(31),
	803:   uint16(1),
	804:   uint16(anon_sym_null),
	805:   uint16(33),
	806:   uint16(1),
	807:   uint16(sym__digit),
	808:   uint16(37),
	809:   uint16(1),
	810:   uint16(anon_sym_0x),
	811:   uint16(39),
	812:   uint16(1),
	813:   uint16(anon_sym_0o),
	814:   uint16(41),
	815:   uint16(1),
	816:   uint16(anon_sym_0b),
	817:   uint16(45),
	818:   uint16(1),
	819:   uint16(anon_sym_BSLASH),
	820:   uint16(15),
	821:   uint16(1),
	822:   uint16(aux_sym_node_repeat1),
	823:   uint16(34),
	824:   uint16(1),
	825:   uint16(sym__escline),
	826:   uint16(58),
	827:   uint16(1),
	828:   uint16(sym__node_space),
	829:   uint16(79),
	830:   uint16(1),
	831:   uint16(sym_type),
	832:   uint16(108),
	833:   uint16(1),
	834:   uint16(sym_node_children),
	835:   uint16(154),
	836:   uint16(1),
	837:   uint16(sym__integer),
	838:   uint16(191),
	839:   uint16(1),
	840:   uint16(sym_string),
	841:   uint16(194),
	842:   uint16(1),
	843:   uint16(sym__escaped_string),
	844:   uint16(198),
	845:   uint16(1),
	846:   uint16(sym_node_field),
	847:   uint16(217),
	848:   uint16(1),
	849:   uint16(sym_boolean),
	850:   uint16(264),
	851:   uint16(1),
	852:   uint16(sym__sign),
	853:   uint16(290),
	854:   uint16(1),
	855:   uint16(sym__bare_identifier),
	856:   uint16(295),
	857:   uint16(1),
	858:   uint16(sym_identifier),
	859:   uint16(35),
	860:   uint16(2),
	861:   uint16(anon_sym_PLUS),
	862:   uint16(anon_sym_DASH),
	863:   uint16(43),
	864:   uint16(2),
	865:   uint16(anon_sym_true),
	866:   uint16(anon_sym_false),
	867:   uint16(27),
	868:   uint16(2),
	869:   uint16(sym__ws),
	870:   uint16(aux_sym_node_repeat3),
	871:   uint16(199),
	872:   uint16(2),
	873:   uint16(sym_keyword),
	874:   uint16(sym_number),
	875:   uint16(211),
	876:   uint16(2),
	877:   uint16(sym__node_field_comment),
	878:   uint16(sym__node_field),
	879:   uint16(213),
	880:   uint16(2),
	881:   uint16(sym_prop),
	882:   uint16(sym_value),
	883:   uint16(47),
	884:   uint16(3),
	885:   uint16(sym_multi_line_comment),
	886:   uint16(sym__bom),
	887:   uint16(sym__unicode_space),
	888:   uint16(160),
	889:   uint16(3),
	890:   uint16(sym__node_terminator),
	891:   uint16(sym__newline),
	892:   uint16(sym_single_line_comment),
	893:   uint16(214),
	894:   uint16(4),
	895:   uint16(sym__decimal),
	896:   uint16(sym__hex),
	897:   uint16(sym__octal),
	898:   uint16(sym__binary),
	899:   uint16(59),
	900:   uint16(9),
	901:   uint16(sym__eof),
	902:   uint16(anon_sym_SEMI),
	903:   uint16(aux_sym__newline_token1),
	904:   uint16(aux_sym__newline_token2),
	905:   uint16(aux_sym__newline_token3),
	906:   uint16(aux_sym__newline_token4),
	907:   uint16(aux_sym__newline_token5),
	908:   uint16(aux_sym__newline_token6),
	909:   uint16(aux_sym__newline_token7),
	910:   uint16(10),
	911:   uint16(3),
	912:   uint16(1),
	913:   uint16(sym_multi_line_comment),
	914:   uint16(23),
	915:   uint16(1),
	916:   uint16(sym__normal_bare_identifier),
	917:   uint16(63),
	918:   uint16(1),
	919:   uint16(anon_sym_DQUOTE),
	920:   uint16(67),
	921:   uint16(1),
	922:   uint16(sym__raw_string),
	923:   uint16(292),
	924:   uint16(1),
	925:   uint16(sym__sign),
	926:   uint16(296),
	927:   uint16(1),
	928:   uint16(sym__escaped_string),
	929:   uint16(65),
	930:   uint16(2),
	931:   uint16(anon_sym_PLUS),
	932:   uint16(anon_sym_DASH),
	933:   uint16(290),
	934:   uint16(2),
	935:   uint16(sym__bare_identifier),
	936:   uint16(sym_string),
	937:   uint16(301),
	938:   uint16(2),
	939:   uint16(sym_identifier),
	940:   uint16(sym_annotation_type),
	941:   uint16(61),
	942:   uint16(37),
	943:   uint16(anon_sym_i8),
	944:   uint16(anon_sym_i16),
	945:   uint16(anon_sym_i32),
	946:   uint16(anon_sym_i64),
	947:   uint16(anon_sym_u8),
	948:   uint16(anon_sym_u16),
	949:   uint16(anon_sym_u32),
	950:   uint16(anon_sym_u64),
	951:   uint16(anon_sym_isize),
	952:   uint16(anon_sym_usize),
	953:   uint16(anon_sym_f32),
	954:   uint16(anon_sym_f64),
	955:   uint16(anon_sym_decimal64),
	956:   uint16(anon_sym_decimal128),
	957:   uint16(anon_sym_date_DASHtime),
	958:   uint16(anon_sym_time),
	959:   uint16(anon_sym_date),
	960:   uint16(anon_sym_duration),
	961:   uint16(anon_sym_decimal),
	962:   uint16(anon_sym_currency),
	963:   uint16(anon_sym_country_DASH2),
	964:   uint16(anon_sym_country_DASH3),
	965:   uint16(anon_sym_country_DASHsubdivision),
	966:   uint16(anon_sym_email),
	967:   uint16(anon_sym_idn_DASHemail),
	968:   uint16(anon_sym_hostname),
	969:   uint16(anon_sym_idn_DASHhostname),
	970:   uint16(anon_sym_ipv4),
	971:   uint16(anon_sym_ipv6),
	972:   uint16(anon_sym_url),
	973:   uint16(anon_sym_url_DASHreference),
	974:   uint16(anon_sym_irl),
	975:   uint16(anon_sym_iri_DASHreference),
	976:   uint16(anon_sym_url_DASHtemplate),
	977:   uint16(anon_sym_uuid),
	978:   uint16(anon_sym_regex),
	979:   uint16(anon_sym_base64),
	980:   uint16(31),
	981:   uint16(11),
	982:   uint16(1),
	983:   uint16(anon_sym_LPAREN),
	984:   uint16(13),
	985:   uint16(1),
	986:   uint16(anon_sym_DQUOTE),
	987:   uint16(21),
	988:   uint16(1),
	989:   uint16(sym__raw_string),
	990:   uint16(23),
	991:   uint16(1),
	992:   uint16(sym__normal_bare_identifier),
	993:   uint16(31),
	994:   uint16(1),
	995:   uint16(anon_sym_null),
	996:   uint16(33),
	997:   uint16(1),
	998:   uint16(sym__digit),
	999:   uint16(37),
	1000:  uint16(1),
	1001:  uint16(anon_sym_0x),
	1002:  uint16(39),
	1003:  uint16(1),
	1004:  uint16(anon_sym_0o),
	1005:  uint16(41),
	1006:  uint16(1),
	1007:  uint16(anon_sym_0b),
	1008:  uint16(69),
	1009:  uint16(1),
	1010:  uint16(anon_sym_SLASH_DASH),
	1011:  uint16(71),
	1012:  uint16(1),
	1013:  uint16(anon_sym_BSLASH),
	1014:  uint16(74),
	1015:  uint16(1),
	1016:  uint16(aux_sym_node_repeat1),
	1017:  uint16(79),
	1018:  uint16(1),
	1019:  uint16(sym_type),
	1020:  uint16(125),
	1021:  uint16(1),
	1022:  uint16(sym__escline),
	1023:  uint16(154),
	1024:  uint16(1),
	1025:  uint16(sym__integer),
	1026:  uint16(186),
	1027:  uint16(1),
	1028:  uint16(sym__node_space),
	1029:  uint16(191),
	1030:  uint16(1),
	1031:  uint16(sym_string),
	1032:  uint16(194),
	1033:  uint16(1),
	1034:  uint16(sym__escaped_string),
	1035:  uint16(198),
	1036:  uint16(1),
	1037:  uint16(sym_node_field),
	1038:  uint16(217),
	1039:  uint16(1),
	1040:  uint16(sym_boolean),
	1041:  uint16(264),
	1042:  uint16(1),
	1043:  uint16(sym__sign),
	1044:  uint16(290),
	1045:  uint16(1),
	1046:  uint16(sym__bare_identifier),
	1047:  uint16(295),
	1048:  uint16(1),
	1049:  uint16(sym_identifier),
	1050:  uint16(35),
	1051:  uint16(2),
	1052:  uint16(anon_sym_PLUS),
	1053:  uint16(anon_sym_DASH),
	1054:  uint16(43),
	1055:  uint16(2),
	1056:  uint16(anon_sym_true),
	1057:  uint16(anon_sym_false),
	1058:  uint16(100),
	1059:  uint16(2),
	1060:  uint16(sym__ws),
	1061:  uint16(aux_sym_node_repeat3),
	1062:  uint16(199),
	1063:  uint16(2),
	1064:  uint16(sym_keyword),
	1065:  uint16(sym_number),
	1066:  uint16(211),
	1067:  uint16(2),
	1068:  uint16(sym__node_field_comment),
	1069:  uint16(sym__node_field),
	1070:  uint16(213),
	1071:  uint16(2),
	1072:  uint16(sym_prop),
	1073:  uint16(sym_value),
	1074:  uint16(73),
	1075:  uint16(3),
	1076:  uint16(sym_multi_line_comment),
	1077:  uint16(sym__bom),
	1078:  uint16(sym__unicode_space),
	1079:  uint16(214),
	1080:  uint16(4),
	1081:  uint16(sym__decimal),
	1082:  uint16(sym__hex),
	1083:  uint16(sym__octal),
	1084:  uint16(sym__binary),
	1085:  uint16(30),
	1086:  uint16(11),
	1087:  uint16(1),
	1088:  uint16(anon_sym_LPAREN),
	1089:  uint16(13),
	1090:  uint16(1),
	1091:  uint16(anon_sym_DQUOTE),
	1092:  uint16(21),
	1093:  uint16(1),
	1094:  uint16(sym__raw_string),
	1095:  uint16(23),
	1096:  uint16(1),
	1097:  uint16(sym__normal_bare_identifier),
	1098:  uint16(31),
	1099:  uint16(1),
	1100:  uint16(anon_sym_null),
	1101:  uint16(33),
	1102:  uint16(1),
	1103:  uint16(sym__digit),
	1104:  uint16(37),
	1105:  uint16(1),
	1106:  uint16(anon_sym_0x),
	1107:  uint16(39),
	1108:  uint16(1),
	1109:  uint16(anon_sym_0o),
	1110:  uint16(41),
	1111:  uint16(1),
	1112:  uint16(anon_sym_0b),
	1113:  uint16(71),
	1114:  uint16(1),
	1115:  uint16(anon_sym_BSLASH),
	1116:  uint16(75),
	1117:  uint16(1),
	1118:  uint16(anon_sym_LBRACE),
	1119:  uint16(12),
	1120:  uint16(1),
	1121:  uint16(aux_sym_node_repeat1),
	1122:  uint16(79),
	1123:  uint16(1),
	1124:  uint16(sym_type),
	1125:  uint16(125),
	1126:  uint16(1),
	1127:  uint16(sym__escline),
	1128:  uint16(154),
	1129:  uint16(1),
	1130:  uint16(sym__integer),
	1131:  uint16(186),
	1132:  uint16(1),
	1133:  uint16(sym__node_space),
	1134:  uint16(191),
	1135:  uint16(1),
	1136:  uint16(sym_string),
	1137:  uint16(194),
	1138:  uint16(1),
	1139:  uint16(sym__escaped_string),
	1140:  uint16(217),
	1141:  uint16(1),
	1142:  uint16(sym_boolean),
	1143:  uint16(221),
	1144:  uint16(1),
	1145:  uint16(sym__node_field),
	1146:  uint16(264),
	1147:  uint16(1),
	1148:  uint16(sym__sign),
	1149:  uint16(290),
	1150:  uint16(1),
	1151:  uint16(sym__bare_identifier),
	1152:  uint16(295),
	1153:  uint16(1),
	1154:  uint16(sym_identifier),
	1155:  uint16(35),
	1156:  uint16(2),
	1157:  uint16(anon_sym_PLUS),
	1158:  uint16(anon_sym_DASH),
	1159:  uint16(43),
	1160:  uint16(2),
	1161:  uint16(anon_sym_true),
	1162:  uint16(anon_sym_false),
	1163:  uint16(100),
	1164:  uint16(2),
	1165:  uint16(sym__ws),
	1166:  uint16(aux_sym_node_repeat3),
	1167:  uint16(199),
	1168:  uint16(2),
	1169:  uint16(sym_keyword),
	1170:  uint16(sym_number),
	1171:  uint16(213),
	1172:  uint16(2),
	1173:  uint16(sym_prop),
	1174:  uint16(sym_value),
	1175:  uint16(73),
	1176:  uint16(3),
	1177:  uint16(sym_multi_line_comment),
	1178:  uint16(sym__bom),
	1179:  uint16(sym__unicode_space),
	1180:  uint16(214),
	1181:  uint16(4),
	1182:  uint16(sym__decimal),
	1183:  uint16(sym__hex),
	1184:  uint16(sym__octal),
	1185:  uint16(sym__binary),
	1186:  uint16(30),
	1187:  uint16(11),
	1188:  uint16(1),
	1189:  uint16(anon_sym_LPAREN),
	1190:  uint16(13),
	1191:  uint16(1),
	1192:  uint16(anon_sym_DQUOTE),
	1193:  uint16(21),
	1194:  uint16(1),
	1195:  uint16(sym__raw_string),
	1196:  uint16(23),
	1197:  uint16(1),
	1198:  uint16(sym__normal_bare_identifier),
	1199:  uint16(31),
	1200:  uint16(1),
	1201:  uint16(anon_sym_null),
	1202:  uint16(33),
	1203:  uint16(1),
	1204:  uint16(sym__digit),
	1205:  uint16(37),
	1206:  uint16(1),
	1207:  uint16(anon_sym_0x),
	1208:  uint16(39),
	1209:  uint16(1),
	1210:  uint16(anon_sym_0o),
	1211:  uint16(41),
	1212:  uint16(1),
	1213:  uint16(anon_sym_0b),
	1214:  uint16(71),
	1215:  uint16(1),
	1216:  uint16(anon_sym_BSLASH),
	1217:  uint16(77),
	1218:  uint16(1),
	1219:  uint16(anon_sym_LBRACE),
	1220:  uint16(74),
	1221:  uint16(1),
	1222:  uint16(aux_sym_node_repeat1),
	1223:  uint16(79),
	1224:  uint16(1),
	1225:  uint16(sym_type),
	1226:  uint16(125),
	1227:  uint16(1),
	1228:  uint16(sym__escline),
	1229:  uint16(154),
	1230:  uint16(1),
	1231:  uint16(sym__integer),
	1232:  uint16(186),
	1233:  uint16(1),
	1234:  uint16(sym__node_space),
	1235:  uint16(191),
	1236:  uint16(1),
	1237:  uint16(sym_string),
	1238:  uint16(194),
	1239:  uint16(1),
	1240:  uint16(sym__escaped_string),
	1241:  uint16(209),
	1242:  uint16(1),
	1243:  uint16(sym__node_field),
	1244:  uint16(217),
	1245:  uint16(1),
	1246:  uint16(sym_boolean),
	1247:  uint16(264),
	1248:  uint16(1),
	1249:  uint16(sym__sign),
	1250:  uint16(290),
	1251:  uint16(1),
	1252:  uint16(sym__bare_identifier),
	1253:  uint16(295),
	1254:  uint16(1),
	1255:  uint16(sym_identifier),
	1256:  uint16(35),
	1257:  uint16(2),
	1258:  uint16(anon_sym_PLUS),
	1259:  uint16(anon_sym_DASH),
	1260:  uint16(43),
	1261:  uint16(2),
	1262:  uint16(anon_sym_true),
	1263:  uint16(anon_sym_false),
	1264:  uint16(100),
	1265:  uint16(2),
	1266:  uint16(sym__ws),
	1267:  uint16(aux_sym_node_repeat3),
	1268:  uint16(199),
	1269:  uint16(2),
	1270:  uint16(sym_keyword),
	1271:  uint16(sym_number),
	1272:  uint16(213),
	1273:  uint16(2),
	1274:  uint16(sym_prop),
	1275:  uint16(sym_value),
	1276:  uint16(73),
	1277:  uint16(3),
	1278:  uint16(sym_multi_line_comment),
	1279:  uint16(sym__bom),
	1280:  uint16(sym__unicode_space),
	1281:  uint16(214),
	1282:  uint16(4),
	1283:  uint16(sym__decimal),
	1284:  uint16(sym__hex),
	1285:  uint16(sym__octal),
	1286:  uint16(sym__binary),
	1287:  uint16(29),
	1288:  uint16(11),
	1289:  uint16(1),
	1290:  uint16(anon_sym_LPAREN),
	1291:  uint16(13),
	1292:  uint16(1),
	1293:  uint16(anon_sym_DQUOTE),
	1294:  uint16(21),
	1295:  uint16(1),
	1296:  uint16(sym__raw_string),
	1297:  uint16(23),
	1298:  uint16(1),
	1299:  uint16(sym__normal_bare_identifier),
	1300:  uint16(31),
	1301:  uint16(1),
	1302:  uint16(anon_sym_null),
	1303:  uint16(33),
	1304:  uint16(1),
	1305:  uint16(sym__digit),
	1306:  uint16(37),
	1307:  uint16(1),
	1308:  uint16(anon_sym_0x),
	1309:  uint16(39),
	1310:  uint16(1),
	1311:  uint16(anon_sym_0o),
	1312:  uint16(41),
	1313:  uint16(1),
	1314:  uint16(anon_sym_0b),
	1315:  uint16(71),
	1316:  uint16(1),
	1317:  uint16(anon_sym_BSLASH),
	1318:  uint16(14),
	1319:  uint16(1),
	1320:  uint16(aux_sym_node_repeat1),
	1321:  uint16(79),
	1322:  uint16(1),
	1323:  uint16(sym_type),
	1324:  uint16(125),
	1325:  uint16(1),
	1326:  uint16(sym__escline),
	1327:  uint16(154),
	1328:  uint16(1),
	1329:  uint16(sym__integer),
	1330:  uint16(186),
	1331:  uint16(1),
	1332:  uint16(sym__node_space),
	1333:  uint16(191),
	1334:  uint16(1),
	1335:  uint16(sym_string),
	1336:  uint16(194),
	1337:  uint16(1),
	1338:  uint16(sym__escaped_string),
	1339:  uint16(217),
	1340:  uint16(1),
	1341:  uint16(sym_boolean),
	1342:  uint16(221),
	1343:  uint16(1),
	1344:  uint16(sym__node_field),
	1345:  uint16(264),
	1346:  uint16(1),
	1347:  uint16(sym__sign),
	1348:  uint16(290),
	1349:  uint16(1),
	1350:  uint16(sym__bare_identifier),
	1351:  uint16(295),
	1352:  uint16(1),
	1353:  uint16(sym_identifier),
	1354:  uint16(35),
	1355:  uint16(2),
	1356:  uint16(anon_sym_PLUS),
	1357:  uint16(anon_sym_DASH),
	1358:  uint16(43),
	1359:  uint16(2),
	1360:  uint16(anon_sym_true),
	1361:  uint16(anon_sym_false),
	1362:  uint16(100),
	1363:  uint16(2),
	1364:  uint16(sym__ws),
	1365:  uint16(aux_sym_node_repeat3),
	1366:  uint16(199),
	1367:  uint16(2),
	1368:  uint16(sym_keyword),
	1369:  uint16(sym_number),
	1370:  uint16(213),
	1371:  uint16(2),
	1372:  uint16(sym_prop),
	1373:  uint16(sym_value),
	1374:  uint16(73),
	1375:  uint16(3),
	1376:  uint16(sym_multi_line_comment),
	1377:  uint16(sym__bom),
	1378:  uint16(sym__unicode_space),
	1379:  uint16(214),
	1380:  uint16(4),
	1381:  uint16(sym__decimal),
	1382:  uint16(sym__hex),
	1383:  uint16(sym__octal),
	1384:  uint16(sym__binary),
	1385:  uint16(29),
	1386:  uint16(11),
	1387:  uint16(1),
	1388:  uint16(anon_sym_LPAREN),
	1389:  uint16(13),
	1390:  uint16(1),
	1391:  uint16(anon_sym_DQUOTE),
	1392:  uint16(21),
	1393:  uint16(1),
	1394:  uint16(sym__raw_string),
	1395:  uint16(23),
	1396:  uint16(1),
	1397:  uint16(sym__normal_bare_identifier),
	1398:  uint16(31),
	1399:  uint16(1),
	1400:  uint16(anon_sym_null),
	1401:  uint16(33),
	1402:  uint16(1),
	1403:  uint16(sym__digit),
	1404:  uint16(37),
	1405:  uint16(1),
	1406:  uint16(anon_sym_0x),
	1407:  uint16(39),
	1408:  uint16(1),
	1409:  uint16(anon_sym_0o),
	1410:  uint16(41),
	1411:  uint16(1),
	1412:  uint16(anon_sym_0b),
	1413:  uint16(71),
	1414:  uint16(1),
	1415:  uint16(anon_sym_BSLASH),
	1416:  uint16(74),
	1417:  uint16(1),
	1418:  uint16(aux_sym_node_repeat1),
	1419:  uint16(79),
	1420:  uint16(1),
	1421:  uint16(sym_type),
	1422:  uint16(125),
	1423:  uint16(1),
	1424:  uint16(sym__escline),
	1425:  uint16(154),
	1426:  uint16(1),
	1427:  uint16(sym__integer),
	1428:  uint16(186),
	1429:  uint16(1),
	1430:  uint16(sym__node_space),
	1431:  uint16(191),
	1432:  uint16(1),
	1433:  uint16(sym_string),
	1434:  uint16(194),
	1435:  uint16(1),
	1436:  uint16(sym__escaped_string),
	1437:  uint16(209),
	1438:  uint16(1),
	1439:  uint16(sym__node_field),
	1440:  uint16(217),
	1441:  uint16(1),
	1442:  uint16(sym_boolean),
	1443:  uint16(264),
	1444:  uint16(1),
	1445:  uint16(sym__sign),
	1446:  uint16(290),
	1447:  uint16(1),
	1448:  uint16(sym__bare_identifier),
	1449:  uint16(295),
	1450:  uint16(1),
	1451:  uint16(sym_identifier),
	1452:  uint16(35),
	1453:  uint16(2),
	1454:  uint16(anon_sym_PLUS),
	1455:  uint16(anon_sym_DASH),
	1456:  uint16(43),
	1457:  uint16(2),
	1458:  uint16(anon_sym_true),
	1459:  uint16(anon_sym_false),
	1460:  uint16(100),
	1461:  uint16(2),
	1462:  uint16(sym__ws),
	1463:  uint16(aux_sym_node_repeat3),
	1464:  uint16(199),
	1465:  uint16(2),
	1466:  uint16(sym_keyword),
	1467:  uint16(sym_number),
	1468:  uint16(213),
	1469:  uint16(2),
	1470:  uint16(sym_prop),
	1471:  uint16(sym_value),
	1472:  uint16(73),
	1473:  uint16(3),
	1474:  uint16(sym_multi_line_comment),
	1475:  uint16(sym__bom),
	1476:  uint16(sym__unicode_space),
	1477:  uint16(214),
	1478:  uint16(4),
	1479:  uint16(sym__decimal),
	1480:  uint16(sym__hex),
	1481:  uint16(sym__octal),
	1482:  uint16(sym__binary),
	1483:  uint16(8),
	1484:  uint16(83),
	1485:  uint16(1),
	1486:  uint16(anon_sym_BSLASH),
	1487:  uint16(15),
	1488:  uint16(1),
	1489:  uint16(aux_sym_node_repeat1),
	1490:  uint16(34),
	1491:  uint16(1),
	1492:  uint16(sym__escline),
	1493:  uint16(58),
	1494:  uint16(1),
	1495:  uint16(sym__node_space),
	1496:  uint16(27),
	1497:  uint16(2),
	1498:  uint16(sym__ws),
	1499:  uint16(aux_sym_node_repeat3),
	1500:  uint16(86),
	1501:  uint16(3),
	1502:  uint16(sym_multi_line_comment),
	1503:  uint16(sym__bom),
	1504:  uint16(sym__unicode_space),
	1505:  uint16(79),
	1506:  uint16(5),
	1507:  uint16(sym__normal_bare_identifier),
	1508:  uint16(anon_sym_null),
	1509:  uint16(sym__digit),
	1510:  uint16(anon_sym_true),
	1511:  uint16(anon_sym_false),
	1512:  uint16(81),
	1513:  uint16(20),
	1514:  uint16(sym__eof),
	1515:  uint16(sym__raw_string),
	1516:  uint16(anon_sym_SLASH_DASH),
	1517:  uint16(anon_sym_LBRACE),
	1518:  uint16(anon_sym_SEMI),
	1519:  uint16(anon_sym_LPAREN),
	1520:  uint16(anon_sym_DQUOTE),
	1521:  uint16(anon_sym_PLUS),
	1522:  uint16(anon_sym_DASH),
	1523:  uint16(anon_sym_0x),
	1524:  uint16(anon_sym_0o),
	1525:  uint16(anon_sym_0b),
	1526:  uint16(aux_sym__newline_token1),
	1527:  uint16(aux_sym__newline_token2),
	1528:  uint16(aux_sym__newline_token3),
	1529:  uint16(aux_sym__newline_token4),
	1530:  uint16(aux_sym__newline_token5),
	1531:  uint16(aux_sym__newline_token6),
	1532:  uint16(aux_sym__newline_token7),
	1533:  uint16(anon_sym_SLASH_SLASH),
	1534:  uint16(16),
	1535:  uint16(91),
	1536:  uint16(1),
	1537:  uint16(sym__normal_bare_identifier),
	1538:  uint16(94),
	1539:  uint16(1),
	1540:  uint16(anon_sym_SLASH_DASH),
	1541:  uint16(97),
	1542:  uint16(1),
	1543:  uint16(anon_sym_LPAREN),
	1544:  uint16(100),
	1545:  uint16(1),
	1546:  uint16(anon_sym_DQUOTE),
	1547:  uint16(109),
	1548:  uint16(1),
	1549:  uint16(anon_sym_SLASH_SLASH),
	1550:  uint16(112),
	1551:  uint16(1),
	1552:  uint16(sym__raw_string),
	1553:  uint16(71),
	1554:  uint16(1),
	1555:  uint16(sym_identifier),
	1556:  uint16(192),
	1557:  uint16(1),
	1558:  uint16(sym__sign),
	1559:  uint16(194),
	1560:  uint16(1),
	1561:  uint16(sym__escaped_string),
	1562:  uint16(260),
	1563:  uint16(1),
	1564:  uint16(sym_type),
	1565:  uint16(89),
	1566:  uint16(2),
	1568:  uint16(anon_sym_RBRACE),
	1569:  uint16(103),
	1570:  uint16(2),
	1571:  uint16(anon_sym_PLUS),
	1572:  uint16(anon_sym_DASH),
	1573:  uint16(16),
	1574:  uint16(2),
	1575:  uint16(sym_node),
	1576:  uint16(aux_sym_document_repeat2),
	1577:  uint16(218),
	1578:  uint16(2),
	1579:  uint16(sym__bare_identifier),
	1580:  uint16(sym_string),
	1581:  uint16(54),
	1582:  uint16(5),
	1583:  uint16(sym__linespace),
	1584:  uint16(sym__newline),
	1585:  uint16(sym__ws),
	1586:  uint16(sym_single_line_comment),
	1587:  uint16(aux_sym_document_repeat1),
	1588:  uint16(106),
	1589:  uint16(10),
	1590:  uint16(sym_multi_line_comment),
	1591:  uint16(aux_sym__newline_token1),
	1592:  uint16(aux_sym__newline_token2),
	1593:  uint16(aux_sym__newline_token3),
	1594:  uint16(aux_sym__newline_token4),
	1595:  uint16(aux_sym__newline_token5),
	1596:  uint16(aux_sym__newline_token6),
	1597:  uint16(aux_sym__newline_token7),
	1598:  uint16(sym__bom),
	1599:  uint16(sym__unicode_space),
	1600:  uint16(16),
	1601:  uint16(7),
	1602:  uint16(1),
	1603:  uint16(sym__normal_bare_identifier),
	1604:  uint16(9),
	1605:  uint16(1),
	1606:  uint16(anon_sym_SLASH_DASH),
	1607:  uint16(11),
	1608:  uint16(1),
	1609:  uint16(anon_sym_LPAREN),
	1610:  uint16(13),
	1611:  uint16(1),
	1612:  uint16(anon_sym_DQUOTE),
	1613:  uint16(19),
	1614:  uint16(1),
	1615:  uint16(anon_sym_SLASH_SLASH),
	1616:  uint16(21),
	1617:  uint16(1),
	1618:  uint16(sym__raw_string),
	1619:  uint16(115),
	1620:  uint16(1),
	1622:  uint16(71),
	1623:  uint16(1),
	1624:  uint16(sym_identifier),
	1625:  uint16(192),
	1626:  uint16(1),
	1627:  uint16(sym__sign),
	1628:  uint16(194),
	1629:  uint16(1),
	1630:  uint16(sym__escaped_string),
	1631:  uint16(260),
	1632:  uint16(1),
	1633:  uint16(sym_type),
	1634:  uint16(15),
	1635:  uint16(2),
	1636:  uint16(anon_sym_PLUS),
	1637:  uint16(anon_sym_DASH),
	1638:  uint16(29),
	1639:  uint16(2),
	1640:  uint16(sym_node),
	1641:  uint16(aux_sym_document_repeat2),
	1642:  uint16(218),
	1643:  uint16(2),
	1644:  uint16(sym__bare_identifier),
	1645:  uint16(sym_string),
	1646:  uint16(43),
	1647:  uint16(5),
	1648:  uint16(sym__linespace),
	1649:  uint16(sym__newline),
	1650:  uint16(sym__ws),
	1651:  uint16(sym_single_line_comment),
	1652:  uint16(aux_sym_document_repeat1),
	1653:  uint16(117),
	1654:  uint16(10),
	1655:  uint16(sym_multi_line_comment),
	1656:  uint16(aux_sym__newline_token1),
	1657:  uint16(aux_sym__newline_token2),
	1658:  uint16(aux_sym__newline_token3),
	1659:  uint16(aux_sym__newline_token4),
	1660:  uint16(aux_sym__newline_token5),
	1661:  uint16(aux_sym__newline_token6),
	1662:  uint16(aux_sym__newline_token7),
	1663:  uint16(sym__bom),
	1664:  uint16(sym__unicode_space),
	1665:  uint16(16),
	1666:  uint16(7),
	1667:  uint16(1),
	1668:  uint16(sym__normal_bare_identifier),
	1669:  uint16(9),
	1670:  uint16(1),
	1671:  uint16(anon_sym_SLASH_DASH),
	1672:  uint16(11),
	1673:  uint16(1),
	1674:  uint16(anon_sym_LPAREN),
	1675:  uint16(13),
	1676:  uint16(1),
	1677:  uint16(anon_sym_DQUOTE),
	1678:  uint16(19),
	1679:  uint16(1),
	1680:  uint16(anon_sym_SLASH_SLASH),
	1681:  uint16(21),
	1682:  uint16(1),
	1683:  uint16(sym__raw_string),
	1684:  uint16(119),
	1685:  uint16(1),
	1686:  uint16(anon_sym_RBRACE),
	1687:  uint16(71),
	1688:  uint16(1),
	1689:  uint16(sym_identifier),
	1690:  uint16(192),
	1691:  uint16(1),
	1692:  uint16(sym__sign),
	1693:  uint16(194),
	1694:  uint16(1),
	1695:  uint16(sym__escaped_string),
	1696:  uint16(260),
	1697:  uint16(1),
	1698:  uint16(sym_type),
	1699:  uint16(15),
	1700:  uint16(2),
	1701:  uint16(anon_sym_PLUS),
	1702:  uint16(anon_sym_DASH),
	1703:  uint16(20),
	1704:  uint16(2),
	1705:  uint16(sym_node),
	1706:  uint16(aux_sym_document_repeat2),
	1707:  uint16(218),
	1708:  uint16(2),
	1709:  uint16(sym__bare_identifier),
	1710:  uint16(sym_string),
	1711:  uint16(51),
	1712:  uint16(5),
	1713:  uint16(sym__linespace),
	1714:  uint16(sym__newline),
	1715:  uint16(sym__ws),
	1716:  uint16(sym_single_line_comment),
	1717:  uint16(aux_sym_document_repeat1),
	1718:  uint16(121),
	1719:  uint16(10),
	1720:  uint16(sym_multi_line_comment),
	1721:  uint16(aux_sym__newline_token1),
	1722:  uint16(aux_sym__newline_token2),
	1723:  uint16(aux_sym__newline_token3),
	1724:  uint16(aux_sym__newline_token4),
	1725:  uint16(aux_sym__newline_token5),
	1726:  uint16(aux_sym__newline_token6),
	1727:  uint16(aux_sym__newline_token7),
	1728:  uint16(sym__bom),
	1729:  uint16(sym__unicode_space),
	1730:  uint16(16),
	1731:  uint16(7),
	1732:  uint16(1),
	1733:  uint16(sym__normal_bare_identifier),
	1734:  uint16(9),
	1735:  uint16(1),
	1736:  uint16(anon_sym_SLASH_DASH),
	1737:  uint16(11),
	1738:  uint16(1),
	1739:  uint16(anon_sym_LPAREN),
	1740:  uint16(13),
	1741:  uint16(1),
	1742:  uint16(anon_sym_DQUOTE),
	1743:  uint16(19),
	1744:  uint16(1),
	1745:  uint16(anon_sym_SLASH_SLASH),
	1746:  uint16(21),
	1747:  uint16(1),
	1748:  uint16(sym__raw_string),
	1749:  uint16(123),
	1750:  uint16(1),
	1752:  uint16(71),
	1753:  uint16(1),
	1754:  uint16(sym_identifier),
	1755:  uint16(192),
	1756:  uint16(1),
	1757:  uint16(sym__sign),
	1758:  uint16(194),
	1759:  uint16(1),
	1760:  uint16(sym__escaped_string),
	1761:  uint16(260),
	1762:  uint16(1),
	1763:  uint16(sym_type),
	1764:  uint16(15),
	1765:  uint16(2),
	1766:  uint16(anon_sym_PLUS),
	1767:  uint16(anon_sym_DASH),
	1768:  uint16(16),
	1769:  uint16(2),
	1770:  uint16(sym_node),
	1771:  uint16(aux_sym_document_repeat2),
	1772:  uint16(218),
	1773:  uint16(2),
	1774:  uint16(sym__bare_identifier),
	1775:  uint16(sym_string),
	1776:  uint16(50),
	1777:  uint16(5),
	1778:  uint16(sym__linespace),
	1779:  uint16(sym__newline),
	1780:  uint16(sym__ws),
	1781:  uint16(sym_single_line_comment),
	1782:  uint16(aux_sym_document_repeat1),
	1783:  uint16(125),
	1784:  uint16(10),
	1785:  uint16(sym_multi_line_comment),
	1786:  uint16(aux_sym__newline_token1),
	1787:  uint16(aux_sym__newline_token2),
	1788:  uint16(aux_sym__newline_token3),
	1789:  uint16(aux_sym__newline_token4),
	1790:  uint16(aux_sym__newline_token5),
	1791:  uint16(aux_sym__newline_token6),
	1792:  uint16(aux_sym__newline_token7),
	1793:  uint16(sym__bom),
	1794:  uint16(sym__unicode_space),
	1795:  uint16(16),
	1796:  uint16(7),
	1797:  uint16(1),
	1798:  uint16(sym__normal_bare_identifier),
	1799:  uint16(9),
	1800:  uint16(1),
	1801:  uint16(anon_sym_SLASH_DASH),
	1802:  uint16(11),
	1803:  uint16(1),
	1804:  uint16(anon_sym_LPAREN),
	1805:  uint16(13),
	1806:  uint16(1),
	1807:  uint16(anon_sym_DQUOTE),
	1808:  uint16(19),
	1809:  uint16(1),
	1810:  uint16(anon_sym_SLASH_SLASH),
	1811:  uint16(21),
	1812:  uint16(1),
	1813:  uint16(sym__raw_string),
	1814:  uint16(127),
	1815:  uint16(1),
	1816:  uint16(anon_sym_RBRACE),
	1817:  uint16(71),
	1818:  uint16(1),
	1819:  uint16(sym_identifier),
	1820:  uint16(192),
	1821:  uint16(1),
	1822:  uint16(sym__sign),
	1823:  uint16(194),
	1824:  uint16(1),
	1825:  uint16(sym__escaped_string),
	1826:  uint16(260),
	1827:  uint16(1),
	1828:  uint16(sym_type),
	1829:  uint16(15),
	1830:  uint16(2),
	1831:  uint16(anon_sym_PLUS),
	1832:  uint16(anon_sym_DASH),
	1833:  uint16(16),
	1834:  uint16(2),
	1835:  uint16(sym_node),
	1836:  uint16(aux_sym_document_repeat2),
	1837:  uint16(218),
	1838:  uint16(2),
	1839:  uint16(sym__bare_identifier),
	1840:  uint16(sym_string),
	1841:  uint16(41),
	1842:  uint16(5),
	1843:  uint16(sym__linespace),
	1844:  uint16(sym__newline),
	1845:  uint16(sym__ws),
	1846:  uint16(sym_single_line_comment),
	1847:  uint16(aux_sym_document_repeat1),
	1848:  uint16(129),
	1849:  uint16(10),
	1850:  uint16(sym_multi_line_comment),
	1851:  uint16(aux_sym__newline_token1),
	1852:  uint16(aux_sym__newline_token2),
	1853:  uint16(aux_sym__newline_token3),
	1854:  uint16(aux_sym__newline_token4),
	1855:  uint16(aux_sym__newline_token5),
	1856:  uint16(aux_sym__newline_token6),
	1857:  uint16(aux_sym__newline_token7),
	1858:  uint16(sym__bom),
	1859:  uint16(sym__unicode_space),
	1860:  uint16(16),
	1861:  uint16(7),
	1862:  uint16(1),
	1863:  uint16(sym__normal_bare_identifier),
	1864:  uint16(9),
	1865:  uint16(1),
	1866:  uint16(anon_sym_SLASH_DASH),
	1867:  uint16(11),
	1868:  uint16(1),
	1869:  uint16(anon_sym_LPAREN),
	1870:  uint16(13),
	1871:  uint16(1),
	1872:  uint16(anon_sym_DQUOTE),
	1873:  uint16(19),
	1874:  uint16(1),
	1875:  uint16(anon_sym_SLASH_SLASH),
	1876:  uint16(21),
	1877:  uint16(1),
	1878:  uint16(sym__raw_string),
	1879:  uint16(119),
	1880:  uint16(1),
	1881:  uint16(anon_sym_RBRACE),
	1882:  uint16(71),
	1883:  uint16(1),
	1884:  uint16(sym_identifier),
	1885:  uint16(192),
	1886:  uint16(1),
	1887:  uint16(sym__sign),
	1888:  uint16(194),
	1889:  uint16(1),
	1890:  uint16(sym__escaped_string),
	1891:  uint16(260),
	1892:  uint16(1),
	1893:  uint16(sym_type),
	1894:  uint16(15),
	1895:  uint16(2),
	1896:  uint16(anon_sym_PLUS),
	1897:  uint16(anon_sym_DASH),
	1898:  uint16(16),
	1899:  uint16(2),
	1900:  uint16(sym_node),
	1901:  uint16(aux_sym_document_repeat2),
	1902:  uint16(218),
	1903:  uint16(2),
	1904:  uint16(sym__bare_identifier),
	1905:  uint16(sym_string),
	1906:  uint16(51),
	1907:  uint16(5),
	1908:  uint16(sym__linespace),
	1909:  uint16(sym__newline),
	1910:  uint16(sym__ws),
	1911:  uint16(sym_single_line_comment),
	1912:  uint16(aux_sym_document_repeat1),
	1913:  uint16(121),
	1914:  uint16(10),
	1915:  uint16(sym_multi_line_comment),
	1916:  uint16(aux_sym__newline_token1),
	1917:  uint16(aux_sym__newline_token2),
	1918:  uint16(aux_sym__newline_token3),
	1919:  uint16(aux_sym__newline_token4),
	1920:  uint16(aux_sym__newline_token5),
	1921:  uint16(aux_sym__newline_token6),
	1922:  uint16(aux_sym__newline_token7),
	1923:  uint16(sym__bom),
	1924:  uint16(sym__unicode_space),
	1925:  uint16(16),
	1926:  uint16(7),
	1927:  uint16(1),
	1928:  uint16(sym__normal_bare_identifier),
	1929:  uint16(9),
	1930:  uint16(1),
	1931:  uint16(anon_sym_SLASH_DASH),
	1932:  uint16(11),
	1933:  uint16(1),
	1934:  uint16(anon_sym_LPAREN),
	1935:  uint16(13),
	1936:  uint16(1),
	1937:  uint16(anon_sym_DQUOTE),
	1938:  uint16(19),
	1939:  uint16(1),
	1940:  uint16(anon_sym_SLASH_SLASH),
	1941:  uint16(21),
	1942:  uint16(1),
	1943:  uint16(sym__raw_string),
	1944:  uint16(131),
	1945:  uint16(1),
	1946:  uint16(anon_sym_RBRACE),
	1947:  uint16(71),
	1948:  uint16(1),
	1949:  uint16(sym_identifier),
	1950:  uint16(192),
	1951:  uint16(1),
	1952:  uint16(sym__sign),
	1953:  uint16(194),
	1954:  uint16(1),
	1955:  uint16(sym__escaped_string),
	1956:  uint16(260),
	1957:  uint16(1),
	1958:  uint16(sym_type),
	1959:  uint16(15),
	1960:  uint16(2),
	1961:  uint16(anon_sym_PLUS),
	1962:  uint16(anon_sym_DASH),
	1963:  uint16(21),
	1964:  uint16(2),
	1965:  uint16(sym_node),
	1966:  uint16(aux_sym_document_repeat2),
	1967:  uint16(218),
	1968:  uint16(2),
	1969:  uint16(sym__bare_identifier),
	1970:  uint16(sym_string),
	1971:  uint16(47),
	1972:  uint16(5),
	1973:  uint16(sym__linespace),
	1974:  uint16(sym__newline),
	1975:  uint16(sym__ws),
	1976:  uint16(sym_single_line_comment),
	1977:  uint16(aux_sym_document_repeat1),
	1978:  uint16(133),
	1979:  uint16(10),
	1980:  uint16(sym_multi_line_comment),
	1981:  uint16(aux_sym__newline_token1),
	1982:  uint16(aux_sym__newline_token2),
	1983:  uint16(aux_sym__newline_token3),
	1984:  uint16(aux_sym__newline_token4),
	1985:  uint16(aux_sym__newline_token5),
	1986:  uint16(aux_sym__newline_token6),
	1987:  uint16(aux_sym__newline_token7),
	1988:  uint16(sym__bom),
	1989:  uint16(sym__unicode_space),
	1990:  uint16(16),
	1991:  uint16(7),
	1992:  uint16(1),
	1993:  uint16(sym__normal_bare_identifier),
	1994:  uint16(9),
	1995:  uint16(1),
	1996:  uint16(anon_sym_SLASH_DASH),
	1997:  uint16(11),
	1998:  uint16(1),
	1999:  uint16(anon_sym_LPAREN),
	2000:  uint16(13),
	2001:  uint16(1),
	2002:  uint16(anon_sym_DQUOTE),
	2003:  uint16(19),
	2004:  uint16(1),
	2005:  uint16(anon_sym_SLASH_SLASH),
	2006:  uint16(21),
	2007:  uint16(1),
	2008:  uint16(sym__raw_string),
	2009:  uint16(135),
	2010:  uint16(1),
	2011:  uint16(anon_sym_RBRACE),
	2012:  uint16(71),
	2013:  uint16(1),
	2014:  uint16(sym_identifier),
	2015:  uint16(192),
	2016:  uint16(1),
	2017:  uint16(sym__sign),
	2018:  uint16(194),
	2019:  uint16(1),
	2020:  uint16(sym__escaped_string),
	2021:  uint16(260),
	2022:  uint16(1),
	2023:  uint16(sym_type),
	2024:  uint16(15),
	2025:  uint16(2),
	2026:  uint16(anon_sym_PLUS),
	2027:  uint16(anon_sym_DASH),
	2028:  uint16(31),
	2029:  uint16(2),
	2030:  uint16(sym_node),
	2031:  uint16(aux_sym_document_repeat2),
	2032:  uint16(218),
	2033:  uint16(2),
	2034:  uint16(sym__bare_identifier),
	2035:  uint16(sym_string),
	2036:  uint16(32),
	2037:  uint16(5),
	2038:  uint16(sym__linespace),
	2039:  uint16(sym__newline),
	2040:  uint16(sym__ws),
	2041:  uint16(sym_single_line_comment),
	2042:  uint16(aux_sym_document_repeat1),
	2043:  uint16(137),
	2044:  uint16(10),
	2045:  uint16(sym_multi_line_comment),
	2046:  uint16(aux_sym__newline_token1),
	2047:  uint16(aux_sym__newline_token2),
	2048:  uint16(aux_sym__newline_token3),
	2049:  uint16(aux_sym__newline_token4),
	2050:  uint16(aux_sym__newline_token5),
	2051:  uint16(aux_sym__newline_token6),
	2052:  uint16(aux_sym__newline_token7),
	2053:  uint16(sym__bom),
	2054:  uint16(sym__unicode_space),
	2055:  uint16(16),
	2056:  uint16(7),
	2057:  uint16(1),
	2058:  uint16(sym__normal_bare_identifier),
	2059:  uint16(9),
	2060:  uint16(1),
	2061:  uint16(anon_sym_SLASH_DASH),
	2062:  uint16(11),
	2063:  uint16(1),
	2064:  uint16(anon_sym_LPAREN),
	2065:  uint16(13),
	2066:  uint16(1),
	2067:  uint16(anon_sym_DQUOTE),
	2068:  uint16(19),
	2069:  uint16(1),
	2070:  uint16(anon_sym_SLASH_SLASH),
	2071:  uint16(21),
	2072:  uint16(1),
	2073:  uint16(sym__raw_string),
	2074:  uint16(139),
	2075:  uint16(1),
	2076:  uint16(anon_sym_RBRACE),
	2077:  uint16(71),
	2078:  uint16(1),
	2079:  uint16(sym_identifier),
	2080:  uint16(192),
	2081:  uint16(1),
	2082:  uint16(sym__sign),
	2083:  uint16(194),
	2084:  uint16(1),
	2085:  uint16(sym__escaped_string),
	2086:  uint16(260),
	2087:  uint16(1),
	2088:  uint16(sym_type),
	2089:  uint16(15),
	2090:  uint16(2),
	2091:  uint16(anon_sym_PLUS),
	2092:  uint16(anon_sym_DASH),
	2093:  uint16(25),
	2094:  uint16(2),
	2095:  uint16(sym_node),
	2096:  uint16(aux_sym_document_repeat2),
	2097:  uint16(218),
	2098:  uint16(2),
	2099:  uint16(sym__bare_identifier),
	2100:  uint16(sym_string),
	2101:  uint16(40),
	2102:  uint16(5),
	2103:  uint16(sym__linespace),
	2104:  uint16(sym__newline),
	2105:  uint16(sym__ws),
	2106:  uint16(sym_single_line_comment),
	2107:  uint16(aux_sym_document_repeat1),
	2108:  uint16(141),
	2109:  uint16(10),
	2110:  uint16(sym_multi_line_comment),
	2111:  uint16(aux_sym__newline_token1),
	2112:  uint16(aux_sym__newline_token2),
	2113:  uint16(aux_sym__newline_token3),
	2114:  uint16(aux_sym__newline_token4),
	2115:  uint16(aux_sym__newline_token5),
	2116:  uint16(aux_sym__newline_token6),
	2117:  uint16(aux_sym__newline_token7),
	2118:  uint16(sym__bom),
	2119:  uint16(sym__unicode_space),
	2120:  uint16(16),
	2121:  uint16(7),
	2122:  uint16(1),
	2123:  uint16(sym__normal_bare_identifier),
	2124:  uint16(9),
	2125:  uint16(1),
	2126:  uint16(anon_sym_SLASH_DASH),
	2127:  uint16(11),
	2128:  uint16(1),
	2129:  uint16(anon_sym_LPAREN),
	2130:  uint16(13),
	2131:  uint16(1),
	2132:  uint16(anon_sym_DQUOTE),
	2133:  uint16(19),
	2134:  uint16(1),
	2135:  uint16(anon_sym_SLASH_SLASH),
	2136:  uint16(21),
	2137:  uint16(1),
	2138:  uint16(sym__raw_string),
	2139:  uint16(143),
	2140:  uint16(1),
	2141:  uint16(anon_sym_RBRACE),
	2142:  uint16(71),
	2143:  uint16(1),
	2144:  uint16(sym_identifier),
	2145:  uint16(192),
	2146:  uint16(1),
	2147:  uint16(sym__sign),
	2148:  uint16(194),
	2149:  uint16(1),
	2150:  uint16(sym__escaped_string),
	2151:  uint16(260),
	2152:  uint16(1),
	2153:  uint16(sym_type),
	2154:  uint16(15),
	2155:  uint16(2),
	2156:  uint16(anon_sym_PLUS),
	2157:  uint16(anon_sym_DASH),
	2158:  uint16(16),
	2159:  uint16(2),
	2160:  uint16(sym_node),
	2161:  uint16(aux_sym_document_repeat2),
	2162:  uint16(218),
	2163:  uint16(2),
	2164:  uint16(sym__bare_identifier),
	2165:  uint16(sym_string),
	2166:  uint16(37),
	2167:  uint16(5),
	2168:  uint16(sym__linespace),
	2169:  uint16(sym__newline),
	2170:  uint16(sym__ws),
	2171:  uint16(sym_single_line_comment),
	2172:  uint16(aux_sym_document_repeat1),
	2173:  uint16(145),
	2174:  uint16(10),
	2175:  uint16(sym_multi_line_comment),
	2176:  uint16(aux_sym__newline_token1),
	2177:  uint16(aux_sym__newline_token2),
	2178:  uint16(aux_sym__newline_token3),
	2179:  uint16(aux_sym__newline_token4),
	2180:  uint16(aux_sym__newline_token5),
	2181:  uint16(aux_sym__newline_token6),
	2182:  uint16(aux_sym__newline_token7),
	2183:  uint16(sym__bom),
	2184:  uint16(sym__unicode_space),
	2185:  uint16(16),
	2186:  uint16(7),
	2187:  uint16(1),
	2188:  uint16(sym__normal_bare_identifier),
	2189:  uint16(9),
	2190:  uint16(1),
	2191:  uint16(anon_sym_SLASH_DASH),
	2192:  uint16(11),
	2193:  uint16(1),
	2194:  uint16(anon_sym_LPAREN),
	2195:  uint16(13),
	2196:  uint16(1),
	2197:  uint16(anon_sym_DQUOTE),
	2198:  uint16(19),
	2199:  uint16(1),
	2200:  uint16(anon_sym_SLASH_SLASH),
	2201:  uint16(21),
	2202:  uint16(1),
	2203:  uint16(sym__raw_string),
	2204:  uint16(143),
	2205:  uint16(1),
	2206:  uint16(anon_sym_RBRACE),
	2207:  uint16(71),
	2208:  uint16(1),
	2209:  uint16(sym_identifier),
	2210:  uint16(192),
	2211:  uint16(1),
	2212:  uint16(sym__sign),
	2213:  uint16(194),
	2214:  uint16(1),
	2215:  uint16(sym__escaped_string),
	2216:  uint16(260),
	2217:  uint16(1),
	2218:  uint16(sym_type),
	2219:  uint16(15),
	2220:  uint16(2),
	2221:  uint16(anon_sym_PLUS),
	2222:  uint16(anon_sym_DASH),
	2223:  uint16(30),
	2224:  uint16(2),
	2225:  uint16(sym_node),
	2226:  uint16(aux_sym_document_repeat2),
	2227:  uint16(218),
	2228:  uint16(2),
	2229:  uint16(sym__bare_identifier),
	2230:  uint16(sym_string),
	2231:  uint16(37),
	2232:  uint16(5),
	2233:  uint16(sym__linespace),
	2234:  uint16(sym__newline),
	2235:  uint16(sym__ws),
	2236:  uint16(sym_single_line_comment),
	2237:  uint16(aux_sym_document_repeat1),
	2238:  uint16(145),
	2239:  uint16(10),
	2240:  uint16(sym_multi_line_comment),
	2241:  uint16(aux_sym__newline_token1),
	2242:  uint16(aux_sym__newline_token2),
	2243:  uint16(aux_sym__newline_token3),
	2244:  uint16(aux_sym__newline_token4),
	2245:  uint16(aux_sym__newline_token5),
	2246:  uint16(aux_sym__newline_token6),
	2247:  uint16(aux_sym__newline_token7),
	2248:  uint16(sym__bom),
	2249:  uint16(sym__unicode_space),
	2250:  uint16(6),
	2251:  uint16(151),
	2252:  uint16(1),
	2253:  uint16(anon_sym_BSLASH),
	2254:  uint16(39),
	2255:  uint16(1),
	2256:  uint16(sym__escline),
	2257:  uint16(49),
	2258:  uint16(2),
	2259:  uint16(sym__ws),
	2260:  uint16(aux_sym_node_repeat3),
	2261:  uint16(154),
	2262:  uint16(3),
	2263:  uint16(sym_multi_line_comment),
	2264:  uint16(sym__bom),
	2265:  uint16(sym__unicode_space),
	2266:  uint16(147),
	2267:  uint16(5),
	2268:  uint16(sym__normal_bare_identifier),
	2269:  uint16(anon_sym_null),
	2270:  uint16(sym__digit),
	2271:  uint16(anon_sym_true),
	2272:  uint16(anon_sym_false),
	2273:  uint16(149),
	2274:  uint16(20),
	2275:  uint16(sym__eof),
	2276:  uint16(sym__raw_string),
	2277:  uint16(anon_sym_SLASH_DASH),
	2278:  uint16(anon_sym_LBRACE),
	2279:  uint16(anon_sym_SEMI),
	2280:  uint16(anon_sym_LPAREN),
	2281:  uint16(anon_sym_DQUOTE),
	2282:  uint16(anon_sym_PLUS),
	2283:  uint16(anon_sym_DASH),
	2284:  uint16(anon_sym_0x),
	2285:  uint16(anon_sym_0o),
	2286:  uint16(anon_sym_0b),
	2287:  uint16(aux_sym__newline_token1),
	2288:  uint16(aux_sym__newline_token2),
	2289:  uint16(aux_sym__newline_token3),
	2290:  uint16(aux_sym__newline_token4),
	2291:  uint16(aux_sym__newline_token5),
	2292:  uint16(aux_sym__newline_token6),
	2293:  uint16(aux_sym__newline_token7),
	2294:  uint16(anon_sym_SLASH_SLASH),
	2295:  uint16(16),
	2296:  uint16(7),
	2297:  uint16(1),
	2298:  uint16(sym__normal_bare_identifier),
	2299:  uint16(9),
	2300:  uint16(1),
	2301:  uint16(anon_sym_SLASH_DASH),
	2302:  uint16(11),
	2303:  uint16(1),
	2304:  uint16(anon_sym_LPAREN),
	2305:  uint16(13),
	2306:  uint16(1),
	2307:  uint16(anon_sym_DQUOTE),
	2308:  uint16(19),
	2309:  uint16(1),
	2310:  uint16(anon_sym_SLASH_SLASH),
	2311:  uint16(21),
	2312:  uint16(1),
	2313:  uint16(sym__raw_string),
	2314:  uint16(157),
	2315:  uint16(1),
	2317:  uint16(71),
	2318:  uint16(1),
	2319:  uint16(sym_identifier),
	2320:  uint16(192),
	2321:  uint16(1),
	2322:  uint16(sym__sign),
	2323:  uint16(194),
	2324:  uint16(1),
	2325:  uint16(sym__escaped_string),
	2326:  uint16(260),
	2327:  uint16(1),
	2328:  uint16(sym_type),
	2329:  uint16(15),
	2330:  uint16(2),
	2331:  uint16(anon_sym_PLUS),
	2332:  uint16(anon_sym_DASH),
	2333:  uint16(19),
	2334:  uint16(2),
	2335:  uint16(sym_node),
	2336:  uint16(aux_sym_document_repeat2),
	2337:  uint16(218),
	2338:  uint16(2),
	2339:  uint16(sym__bare_identifier),
	2340:  uint16(sym_string),
	2341:  uint16(52),
	2342:  uint16(5),
	2343:  uint16(sym__linespace),
	2344:  uint16(sym__newline),
	2345:  uint16(sym__ws),
	2346:  uint16(sym_single_line_comment),
	2347:  uint16(aux_sym_document_repeat1),
	2348:  uint16(159),
	2349:  uint16(10),
	2350:  uint16(sym_multi_line_comment),
	2351:  uint16(aux_sym__newline_token1),
	2352:  uint16(aux_sym__newline_token2),
	2353:  uint16(aux_sym__newline_token3),
	2354:  uint16(aux_sym__newline_token4),
	2355:  uint16(aux_sym__newline_token5),
	2356:  uint16(aux_sym__newline_token6),
	2357:  uint16(aux_sym__newline_token7),
	2358:  uint16(sym__bom),
	2359:  uint16(sym__unicode_space),
	2360:  uint16(16),
	2361:  uint16(7),
	2362:  uint16(1),
	2363:  uint16(sym__normal_bare_identifier),
	2364:  uint16(9),
	2365:  uint16(1),
	2366:  uint16(anon_sym_SLASH_DASH),
	2367:  uint16(11),
	2368:  uint16(1),
	2369:  uint16(anon_sym_LPAREN),
	2370:  uint16(13),
	2371:  uint16(1),
	2372:  uint16(anon_sym_DQUOTE),
	2373:  uint16(19),
	2374:  uint16(1),
	2375:  uint16(anon_sym_SLASH_SLASH),
	2376:  uint16(21),
	2377:  uint16(1),
	2378:  uint16(sym__raw_string),
	2379:  uint16(157),
	2380:  uint16(1),
	2382:  uint16(71),
	2383:  uint16(1),
	2384:  uint16(sym_identifier),
	2385:  uint16(192),
	2386:  uint16(1),
	2387:  uint16(sym__sign),
	2388:  uint16(194),
	2389:  uint16(1),
	2390:  uint16(sym__escaped_string),
	2391:  uint16(260),
	2392:  uint16(1),
	2393:  uint16(sym_type),
	2394:  uint16(15),
	2395:  uint16(2),
	2396:  uint16(anon_sym_PLUS),
	2397:  uint16(anon_sym_DASH),
	2398:  uint16(16),
	2399:  uint16(2),
	2400:  uint16(sym_node),
	2401:  uint16(aux_sym_document_repeat2),
	2402:  uint16(218),
	2403:  uint16(2),
	2404:  uint16(sym__bare_identifier),
	2405:  uint16(sym_string),
	2406:  uint16(52),
	2407:  uint16(5),
	2408:  uint16(sym__linespace),
	2409:  uint16(sym__newline),
	2410:  uint16(sym__ws),
	2411:  uint16(sym_single_line_comment),
	2412:  uint16(aux_sym_document_repeat1),
	2413:  uint16(159),
	2414:  uint16(10),
	2415:  uint16(sym_multi_line_comment),
	2416:  uint16(aux_sym__newline_token1),
	2417:  uint16(aux_sym__newline_token2),
	2418:  uint16(aux_sym__newline_token3),
	2419:  uint16(aux_sym__newline_token4),
	2420:  uint16(aux_sym__newline_token5),
	2421:  uint16(aux_sym__newline_token6),
	2422:  uint16(aux_sym__newline_token7),
	2423:  uint16(sym__bom),
	2424:  uint16(sym__unicode_space),
	2425:  uint16(16),
	2426:  uint16(7),
	2427:  uint16(1),
	2428:  uint16(sym__normal_bare_identifier),
	2429:  uint16(9),
	2430:  uint16(1),
	2431:  uint16(anon_sym_SLASH_DASH),
	2432:  uint16(11),
	2433:  uint16(1),
	2434:  uint16(anon_sym_LPAREN),
	2435:  uint16(13),
	2436:  uint16(1),
	2437:  uint16(anon_sym_DQUOTE),
	2438:  uint16(19),
	2439:  uint16(1),
	2440:  uint16(anon_sym_SLASH_SLASH),
	2441:  uint16(21),
	2442:  uint16(1),
	2443:  uint16(sym__raw_string),
	2444:  uint16(161),
	2445:  uint16(1),
	2446:  uint16(anon_sym_RBRACE),
	2447:  uint16(71),
	2448:  uint16(1),
	2449:  uint16(sym_identifier),
	2450:  uint16(192),
	2451:  uint16(1),
	2452:  uint16(sym__sign),
	2453:  uint16(194),
	2454:  uint16(1),
	2455:  uint16(sym__escaped_string),
	2456:  uint16(260),
	2457:  uint16(1),
	2458:  uint16(sym_type),
	2459:  uint16(15),
	2460:  uint16(2),
	2461:  uint16(anon_sym_PLUS),
	2462:  uint16(anon_sym_DASH),
	2463:  uint16(16),
	2464:  uint16(2),
	2465:  uint16(sym_node),
	2466:  uint16(aux_sym_document_repeat2),
	2467:  uint16(218),
	2468:  uint16(2),
	2469:  uint16(sym__bare_identifier),
	2470:  uint16(sym_string),
	2471:  uint16(46),
	2472:  uint16(5),
	2473:  uint16(sym__linespace),
	2474:  uint16(sym__newline),
	2475:  uint16(sym__ws),
	2476:  uint16(sym_single_line_comment),
	2477:  uint16(aux_sym_document_repeat1),
	2478:  uint16(163),
	2479:  uint16(10),
	2480:  uint16(sym_multi_line_comment),
	2481:  uint16(aux_sym__newline_token1),
	2482:  uint16(aux_sym__newline_token2),
	2483:  uint16(aux_sym__newline_token3),
	2484:  uint16(aux_sym__newline_token4),
	2485:  uint16(aux_sym__newline_token5),
	2486:  uint16(aux_sym__newline_token6),
	2487:  uint16(aux_sym__newline_token7),
	2488:  uint16(sym__bom),
	2489:  uint16(sym__unicode_space),
	2490:  uint16(16),
	2491:  uint16(7),
	2492:  uint16(1),
	2493:  uint16(sym__normal_bare_identifier),
	2494:  uint16(9),
	2495:  uint16(1),
	2496:  uint16(anon_sym_SLASH_DASH),
	2497:  uint16(11),
	2498:  uint16(1),
	2499:  uint16(anon_sym_LPAREN),
	2500:  uint16(13),
	2501:  uint16(1),
	2502:  uint16(anon_sym_DQUOTE),
	2503:  uint16(19),
	2504:  uint16(1),
	2505:  uint16(anon_sym_SLASH_SLASH),
	2506:  uint16(21),
	2507:  uint16(1),
	2508:  uint16(sym__raw_string),
	2509:  uint16(139),
	2510:  uint16(1),
	2511:  uint16(anon_sym_RBRACE),
	2512:  uint16(71),
	2513:  uint16(1),
	2514:  uint16(sym_identifier),
	2515:  uint16(192),
	2516:  uint16(1),
	2517:  uint16(sym__sign),
	2518:  uint16(194),
	2519:  uint16(1),
	2520:  uint16(sym__escaped_string),
	2521:  uint16(260),
	2522:  uint16(1),
	2523:  uint16(sym_type),
	2524:  uint16(15),
	2525:  uint16(2),
	2526:  uint16(anon_sym_PLUS),
	2527:  uint16(anon_sym_DASH),
	2528:  uint16(16),
	2529:  uint16(2),
	2530:  uint16(sym_node),
	2531:  uint16(aux_sym_document_repeat2),
	2532:  uint16(218),
	2533:  uint16(2),
	2534:  uint16(sym__bare_identifier),
	2535:  uint16(sym_string),
	2536:  uint16(40),
	2537:  uint16(5),
	2538:  uint16(sym__linespace),
	2539:  uint16(sym__newline),
	2540:  uint16(sym__ws),
	2541:  uint16(sym_single_line_comment),
	2542:  uint16(aux_sym_document_repeat1),
	2543:  uint16(141),
	2544:  uint16(10),
	2545:  uint16(sym_multi_line_comment),
	2546:  uint16(aux_sym__newline_token1),
	2547:  uint16(aux_sym__newline_token2),
	2548:  uint16(aux_sym__newline_token3),
	2549:  uint16(aux_sym__newline_token4),
	2550:  uint16(aux_sym__newline_token5),
	2551:  uint16(aux_sym__newline_token6),
	2552:  uint16(aux_sym__newline_token7),
	2553:  uint16(sym__bom),
	2554:  uint16(sym__unicode_space),
	2555:  uint16(16),
	2556:  uint16(7),
	2557:  uint16(1),
	2558:  uint16(sym__normal_bare_identifier),
	2559:  uint16(9),
	2560:  uint16(1),
	2561:  uint16(anon_sym_SLASH_DASH),
	2562:  uint16(11),
	2563:  uint16(1),
	2564:  uint16(anon_sym_LPAREN),
	2565:  uint16(13),
	2566:  uint16(1),
	2567:  uint16(anon_sym_DQUOTE),
	2568:  uint16(19),
	2569:  uint16(1),
	2570:  uint16(anon_sym_SLASH_SLASH),
	2571:  uint16(21),
	2572:  uint16(1),
	2573:  uint16(sym__raw_string),
	2574:  uint16(139),
	2575:  uint16(1),
	2576:  uint16(anon_sym_RBRACE),
	2577:  uint16(71),
	2578:  uint16(1),
	2579:  uint16(sym_identifier),
	2580:  uint16(153),
	2581:  uint16(1),
	2582:  uint16(sym_node),
	2583:  uint16(192),
	2584:  uint16(1),
	2585:  uint16(sym__sign),
	2586:  uint16(194),
	2587:  uint16(1),
	2588:  uint16(sym__escaped_string),
	2589:  uint16(260),
	2590:  uint16(1),
	2591:  uint16(sym_type),
	2592:  uint16(15),
	2593:  uint16(2),
	2594:  uint16(anon_sym_PLUS),
	2595:  uint16(anon_sym_DASH),
	2596:  uint16(218),
	2597:  uint16(2),
	2598:  uint16(sym__bare_identifier),
	2599:  uint16(sym_string),
	2600:  uint16(73),
	2601:  uint16(5),
	2602:  uint16(sym__linespace),
	2603:  uint16(sym__newline),
	2604:  uint16(sym__ws),
	2605:  uint16(sym_single_line_comment),
	2606:  uint16(aux_sym_document_repeat1),
	2607:  uint16(165),
	2608:  uint16(10),
	2609:  uint16(sym_multi_line_comment),
	2610:  uint16(aux_sym__newline_token1),
	2611:  uint16(aux_sym__newline_token2),
	2612:  uint16(aux_sym__newline_token3),
	2613:  uint16(aux_sym__newline_token4),
	2614:  uint16(aux_sym__newline_token5),
	2615:  uint16(aux_sym__newline_token6),
	2616:  uint16(aux_sym__newline_token7),
	2617:  uint16(sym__bom),
	2618:  uint16(sym__unicode_space),
	2619:  uint16(16),
	2620:  uint16(7),
	2621:  uint16(1),
	2622:  uint16(sym__normal_bare_identifier),
	2623:  uint16(9),
	2624:  uint16(1),
	2625:  uint16(anon_sym_SLASH_DASH),
	2626:  uint16(11),
	2627:  uint16(1),
	2628:  uint16(anon_sym_LPAREN),
	2629:  uint16(13),
	2630:  uint16(1),
	2631:  uint16(anon_sym_DQUOTE),
	2632:  uint16(19),
	2633:  uint16(1),
	2634:  uint16(anon_sym_SLASH_SLASH),
	2635:  uint16(21),
	2636:  uint16(1),
	2637:  uint16(sym__raw_string),
	2638:  uint16(135),
	2639:  uint16(1),
	2640:  uint16(anon_sym_RBRACE),
	2641:  uint16(24),
	2642:  uint16(1),
	2643:  uint16(sym_node),
	2644:  uint16(71),
	2645:  uint16(1),
	2646:  uint16(sym_identifier),
	2647:  uint16(192),
	2648:  uint16(1),
	2649:  uint16(sym__sign),
	2650:  uint16(194),
	2651:  uint16(1),
	2652:  uint16(sym__escaped_string),
	2653:  uint16(260),
	2654:  uint16(1),
	2655:  uint16(sym_type),
	2656:  uint16(15),
	2657:  uint16(2),
	2658:  uint16(anon_sym_PLUS),
	2659:  uint16(anon_sym_DASH),
	2660:  uint16(218),
	2661:  uint16(2),
	2662:  uint16(sym__bare_identifier),
	2663:  uint16(sym_string),
	2664:  uint16(36),
	2665:  uint16(5),
	2666:  uint16(sym__linespace),
	2667:  uint16(sym__newline),
	2668:  uint16(sym__ws),
	2669:  uint16(sym_single_line_comment),
	2670:  uint16(aux_sym_document_repeat1),
	2671:  uint16(167),
	2672:  uint16(10),
	2673:  uint16(sym_multi_line_comment),
	2674:  uint16(aux_sym__newline_token1),
	2675:  uint16(aux_sym__newline_token2),
	2676:  uint16(aux_sym__newline_token3),
	2677:  uint16(aux_sym__newline_token4),
	2678:  uint16(aux_sym__newline_token5),
	2679:  uint16(aux_sym__newline_token6),
	2680:  uint16(aux_sym__newline_token7),
	2681:  uint16(sym__bom),
	2682:  uint16(sym__unicode_space),
	2683:  uint16(4),
	2684:  uint16(53),
	2685:  uint16(2),
	2686:  uint16(sym__ws),
	2687:  uint16(aux_sym_node_repeat3),
	2688:  uint16(169),
	2689:  uint16(3),
	2690:  uint16(sym_multi_line_comment),
	2691:  uint16(sym__bom),
	2692:  uint16(sym__unicode_space),
	2693:  uint16(147),
	2694:  uint16(5),
	2695:  uint16(sym__normal_bare_identifier),
	2696:  uint16(anon_sym_null),
	2697:  uint16(sym__digit),
	2698:  uint16(anon_sym_true),
	2699:  uint16(anon_sym_false),
	2700:  uint16(149),
	2701:  uint16(21),
	2702:  uint16(sym__eof),
	2703:  uint16(sym__raw_string),
	2704:  uint16(anon_sym_SLASH_DASH),
	2705:  uint16(anon_sym_LBRACE),
	2706:  uint16(anon_sym_SEMI),
	2707:  uint16(anon_sym_LPAREN),
	2708:  uint16(anon_sym_DQUOTE),
	2709:  uint16(anon_sym_PLUS),
	2710:  uint16(anon_sym_DASH),
	2711:  uint16(anon_sym_0x),
	2712:  uint16(anon_sym_0o),
	2713:  uint16(anon_sym_0b),
	2714:  uint16(anon_sym_BSLASH),
	2715:  uint16(aux_sym__newline_token1),
	2716:  uint16(aux_sym__newline_token2),
	2717:  uint16(aux_sym__newline_token3),
	2718:  uint16(aux_sym__newline_token4),
	2719:  uint16(aux_sym__newline_token5),
	2720:  uint16(aux_sym__newline_token6),
	2721:  uint16(aux_sym__newline_token7),
	2722:  uint16(anon_sym_SLASH_SLASH),
	2723:  uint16(16),
	2724:  uint16(7),
	2725:  uint16(1),
	2726:  uint16(sym__normal_bare_identifier),
	2727:  uint16(9),
	2728:  uint16(1),
	2729:  uint16(anon_sym_SLASH_DASH),
	2730:  uint16(11),
	2731:  uint16(1),
	2732:  uint16(anon_sym_LPAREN),
	2733:  uint16(13),
	2734:  uint16(1),
	2735:  uint16(anon_sym_DQUOTE),
	2736:  uint16(19),
	2737:  uint16(1),
	2738:  uint16(anon_sym_SLASH_SLASH),
	2739:  uint16(21),
	2740:  uint16(1),
	2741:  uint16(sym__raw_string),
	2742:  uint16(172),
	2743:  uint16(1),
	2744:  uint16(anon_sym_RBRACE),
	2745:  uint16(22),
	2746:  uint16(1),
	2747:  uint16(sym_node),
	2748:  uint16(71),
	2749:  uint16(1),
	2750:  uint16(sym_identifier),
	2751:  uint16(192),
	2752:  uint16(1),
	2753:  uint16(sym__sign),
	2754:  uint16(194),
	2755:  uint16(1),
	2756:  uint16(sym__escaped_string),
	2757:  uint16(260),
	2758:  uint16(1),
	2759:  uint16(sym_type),
	2760:  uint16(15),
	2761:  uint16(2),
	2762:  uint16(anon_sym_PLUS),
	2763:  uint16(anon_sym_DASH),
	2764:  uint16(218),
	2765:  uint16(2),
	2766:  uint16(sym__bare_identifier),
	2767:  uint16(sym_string),
	2768:  uint16(44),
	2769:  uint16(5),
	2770:  uint16(sym__linespace),
	2771:  uint16(sym__newline),
	2772:  uint16(sym__ws),
	2773:  uint16(sym_single_line_comment),
	2774:  uint16(aux_sym_document_repeat1),
	2775:  uint16(174),
	2776:  uint16(10),
	2777:  uint16(sym_multi_line_comment),
	2778:  uint16(aux_sym__newline_token1),
	2779:  uint16(aux_sym__newline_token2),
	2780:  uint16(aux_sym__newline_token3),
	2781:  uint16(aux_sym__newline_token4),
	2782:  uint16(aux_sym__newline_token5),
	2783:  uint16(aux_sym__newline_token6),
	2784:  uint16(aux_sym__newline_token7),
	2785:  uint16(sym__bom),
	2786:  uint16(sym__unicode_space),
	2787:  uint16(16),
	2788:  uint16(7),
	2789:  uint16(1),
	2790:  uint16(sym__normal_bare_identifier),
	2791:  uint16(9),
	2792:  uint16(1),
	2793:  uint16(anon_sym_SLASH_DASH),
	2794:  uint16(11),
	2795:  uint16(1),
	2796:  uint16(anon_sym_LPAREN),
	2797:  uint16(13),
	2798:  uint16(1),
	2799:  uint16(anon_sym_DQUOTE),
	2800:  uint16(19),
	2801:  uint16(1),
	2802:  uint16(anon_sym_SLASH_SLASH),
	2803:  uint16(21),
	2804:  uint16(1),
	2805:  uint16(sym__raw_string),
	2806:  uint16(139),
	2807:  uint16(1),
	2808:  uint16(anon_sym_RBRACE),
	2809:  uint16(26),
	2810:  uint16(1),
	2811:  uint16(sym_node),
	2812:  uint16(71),
	2813:  uint16(1),
	2814:  uint16(sym_identifier),
	2815:  uint16(192),
	2816:  uint16(1),
	2817:  uint16(sym__sign),
	2818:  uint16(194),
	2819:  uint16(1),
	2820:  uint16(sym__escaped_string),
	2821:  uint16(260),
	2822:  uint16(1),
	2823:  uint16(sym_type),
	2824:  uint16(15),
	2825:  uint16(2),
	2826:  uint16(anon_sym_PLUS),
	2827:  uint16(anon_sym_DASH),
	2828:  uint16(218),
	2829:  uint16(2),
	2830:  uint16(sym__bare_identifier),
	2831:  uint16(sym_string),
	2832:  uint16(78),
	2833:  uint16(5),
	2834:  uint16(sym__linespace),
	2835:  uint16(sym__newline),
	2836:  uint16(sym__ws),
	2837:  uint16(sym_single_line_comment),
	2838:  uint16(aux_sym_document_repeat1),
	2839:  uint16(176),
	2840:  uint16(10),
	2841:  uint16(sym_multi_line_comment),
	2842:  uint16(aux_sym__newline_token1),
	2843:  uint16(aux_sym__newline_token2),
	2844:  uint16(aux_sym__newline_token3),
	2845:  uint16(aux_sym__newline_token4),
	2846:  uint16(aux_sym__newline_token5),
	2847:  uint16(aux_sym__newline_token6),
	2848:  uint16(aux_sym__newline_token7),
	2849:  uint16(sym__bom),
	2850:  uint16(sym__unicode_space),
	2851:  uint16(16),
	2852:  uint16(7),
	2853:  uint16(1),
	2854:  uint16(sym__normal_bare_identifier),
	2855:  uint16(9),
	2856:  uint16(1),
	2857:  uint16(anon_sym_SLASH_DASH),
	2858:  uint16(11),
	2859:  uint16(1),
	2860:  uint16(anon_sym_LPAREN),
	2861:  uint16(13),
	2862:  uint16(1),
	2863:  uint16(anon_sym_DQUOTE),
	2864:  uint16(19),
	2865:  uint16(1),
	2866:  uint16(anon_sym_SLASH_SLASH),
	2867:  uint16(21),
	2868:  uint16(1),
	2869:  uint16(sym__raw_string),
	2870:  uint16(161),
	2871:  uint16(1),
	2872:  uint16(anon_sym_RBRACE),
	2873:  uint16(71),
	2874:  uint16(1),
	2875:  uint16(sym_identifier),
	2876:  uint16(153),
	2877:  uint16(1),
	2878:  uint16(sym_node),
	2879:  uint16(192),
	2880:  uint16(1),
	2881:  uint16(sym__sign),
	2882:  uint16(194),
	2883:  uint16(1),
	2884:  uint16(sym__escaped_string),
	2885:  uint16(260),
	2886:  uint16(1),
	2887:  uint16(sym_type),
	2888:  uint16(15),
	2889:  uint16(2),
	2890:  uint16(anon_sym_PLUS),
	2891:  uint16(anon_sym_DASH),
	2892:  uint16(218),
	2893:  uint16(2),
	2894:  uint16(sym__bare_identifier),
	2895:  uint16(sym_string),
	2896:  uint16(73),
	2897:  uint16(5),
	2898:  uint16(sym__linespace),
	2899:  uint16(sym__newline),
	2900:  uint16(sym__ws),
	2901:  uint16(sym_single_line_comment),
	2902:  uint16(aux_sym_document_repeat1),
	2903:  uint16(165),
	2904:  uint16(10),
	2905:  uint16(sym_multi_line_comment),
	2906:  uint16(aux_sym__newline_token1),
	2907:  uint16(aux_sym__newline_token2),
	2908:  uint16(aux_sym__newline_token3),
	2909:  uint16(aux_sym__newline_token4),
	2910:  uint16(aux_sym__newline_token5),
	2911:  uint16(aux_sym__newline_token6),
	2912:  uint16(aux_sym__newline_token7),
	2913:  uint16(sym__bom),
	2914:  uint16(sym__unicode_space),
	2915:  uint16(4),
	2916:  uint16(49),
	2917:  uint16(2),
	2918:  uint16(sym__ws),
	2919:  uint16(aux_sym_node_repeat3),
	2920:  uint16(182),
	2921:  uint16(3),
	2922:  uint16(sym_multi_line_comment),
	2923:  uint16(sym__bom),
	2924:  uint16(sym__unicode_space),
	2925:  uint16(178),
	2926:  uint16(5),
	2927:  uint16(sym__normal_bare_identifier),
	2928:  uint16(anon_sym_null),
	2929:  uint16(sym__digit),
	2930:  uint16(anon_sym_true),
	2931:  uint16(anon_sym_false),
	2932:  uint16(180),
	2933:  uint16(21),
	2934:  uint16(sym__eof),
	2935:  uint16(sym__raw_string),
	2936:  uint16(anon_sym_SLASH_DASH),
	2937:  uint16(anon_sym_LBRACE),
	2938:  uint16(anon_sym_SEMI),
	2939:  uint16(anon_sym_LPAREN),
	2940:  uint16(anon_sym_DQUOTE),
	2941:  uint16(anon_sym_PLUS),
	2942:  uint16(anon_sym_DASH),
	2943:  uint16(anon_sym_0x),
	2944:  uint16(anon_sym_0o),
	2945:  uint16(anon_sym_0b),
	2946:  uint16(anon_sym_BSLASH),
	2947:  uint16(aux_sym__newline_token1),
	2948:  uint16(aux_sym__newline_token2),
	2949:  uint16(aux_sym__newline_token3),
	2950:  uint16(aux_sym__newline_token4),
	2951:  uint16(aux_sym__newline_token5),
	2952:  uint16(aux_sym__newline_token6),
	2953:  uint16(aux_sym__newline_token7),
	2954:  uint16(anon_sym_SLASH_SLASH),
	2955:  uint16(4),
	2956:  uint16(38),
	2957:  uint16(2),
	2958:  uint16(sym__ws),
	2959:  uint16(aux_sym_node_repeat3),
	2960:  uint16(189),
	2961:  uint16(3),
	2962:  uint16(sym_multi_line_comment),
	2963:  uint16(sym__bom),
	2964:  uint16(sym__unicode_space),
	2965:  uint16(185),
	2966:  uint16(5),
	2967:  uint16(sym__normal_bare_identifier),
	2968:  uint16(anon_sym_null),
	2969:  uint16(sym__digit),
	2970:  uint16(anon_sym_true),
	2971:  uint16(anon_sym_false),
	2972:  uint16(187),
	2973:  uint16(21),
	2974:  uint16(sym__eof),
	2975:  uint16(sym__raw_string),
	2976:  uint16(anon_sym_SLASH_DASH),
	2977:  uint16(anon_sym_LBRACE),
	2978:  uint16(anon_sym_SEMI),
	2979:  uint16(anon_sym_LPAREN),
	2980:  uint16(anon_sym_DQUOTE),
	2981:  uint16(anon_sym_PLUS),
	2982:  uint16(anon_sym_DASH),
	2983:  uint16(anon_sym_0x),
	2984:  uint16(anon_sym_0o),
	2985:  uint16(anon_sym_0b),
	2986:  uint16(anon_sym_BSLASH),
	2987:  uint16(aux_sym__newline_token1),
	2988:  uint16(aux_sym__newline_token2),
	2989:  uint16(aux_sym__newline_token3),
	2990:  uint16(aux_sym__newline_token4),
	2991:  uint16(aux_sym__newline_token5),
	2992:  uint16(aux_sym__newline_token6),
	2993:  uint16(aux_sym__newline_token7),
	2994:  uint16(anon_sym_SLASH_SLASH),
	2995:  uint16(16),
	2996:  uint16(7),
	2997:  uint16(1),
	2998:  uint16(sym__normal_bare_identifier),
	2999:  uint16(9),
	3000:  uint16(1),
	3001:  uint16(anon_sym_SLASH_DASH),
	3002:  uint16(11),
	3003:  uint16(1),
	3004:  uint16(anon_sym_LPAREN),
	3005:  uint16(13),
	3006:  uint16(1),
	3007:  uint16(anon_sym_DQUOTE),
	3008:  uint16(19),
	3009:  uint16(1),
	3010:  uint16(anon_sym_SLASH_SLASH),
	3011:  uint16(21),
	3012:  uint16(1),
	3013:  uint16(sym__raw_string),
	3014:  uint16(143),
	3015:  uint16(1),
	3016:  uint16(anon_sym_RBRACE),
	3017:  uint16(71),
	3018:  uint16(1),
	3019:  uint16(sym_identifier),
	3020:  uint16(153),
	3021:  uint16(1),
	3022:  uint16(sym_node),
	3023:  uint16(192),
	3024:  uint16(1),
	3025:  uint16(sym__sign),
	3026:  uint16(194),
	3027:  uint16(1),
	3028:  uint16(sym__escaped_string),
	3029:  uint16(260),
	3030:  uint16(1),
	3031:  uint16(sym_type),
	3032:  uint16(15),
	3033:  uint16(2),
	3034:  uint16(anon_sym_PLUS),
	3035:  uint16(anon_sym_DASH),
	3036:  uint16(218),
	3037:  uint16(2),
	3038:  uint16(sym__bare_identifier),
	3039:  uint16(sym_string),
	3040:  uint16(73),
	3041:  uint16(5),
	3042:  uint16(sym__linespace),
	3043:  uint16(sym__newline),
	3044:  uint16(sym__ws),
	3045:  uint16(sym_single_line_comment),
	3046:  uint16(aux_sym_document_repeat1),
	3047:  uint16(165),
	3048:  uint16(10),
	3049:  uint16(sym_multi_line_comment),
	3050:  uint16(aux_sym__newline_token1),
	3051:  uint16(aux_sym__newline_token2),
	3052:  uint16(aux_sym__newline_token3),
	3053:  uint16(aux_sym__newline_token4),
	3054:  uint16(aux_sym__newline_token5),
	3055:  uint16(aux_sym__newline_token6),
	3056:  uint16(aux_sym__newline_token7),
	3057:  uint16(sym__bom),
	3058:  uint16(sym__unicode_space),
	3059:  uint16(16),
	3060:  uint16(7),
	3061:  uint16(1),
	3062:  uint16(sym__normal_bare_identifier),
	3063:  uint16(9),
	3064:  uint16(1),
	3065:  uint16(anon_sym_SLASH_DASH),
	3066:  uint16(11),
	3067:  uint16(1),
	3068:  uint16(anon_sym_LPAREN),
	3069:  uint16(13),
	3070:  uint16(1),
	3071:  uint16(anon_sym_DQUOTE),
	3072:  uint16(19),
	3073:  uint16(1),
	3074:  uint16(anon_sym_SLASH_SLASH),
	3075:  uint16(21),
	3076:  uint16(1),
	3077:  uint16(sym__raw_string),
	3078:  uint16(192),
	3079:  uint16(1),
	3080:  uint16(anon_sym_RBRACE),
	3081:  uint16(71),
	3082:  uint16(1),
	3083:  uint16(sym_identifier),
	3084:  uint16(153),
	3085:  uint16(1),
	3086:  uint16(sym_node),
	3087:  uint16(192),
	3088:  uint16(1),
	3089:  uint16(sym__sign),
	3090:  uint16(194),
	3091:  uint16(1),
	3092:  uint16(sym__escaped_string),
	3093:  uint16(260),
	3094:  uint16(1),
	3095:  uint16(sym_type),
	3096:  uint16(15),
	3097:  uint16(2),
	3098:  uint16(anon_sym_PLUS),
	3099:  uint16(anon_sym_DASH),
	3100:  uint16(218),
	3101:  uint16(2),
	3102:  uint16(sym__bare_identifier),
	3103:  uint16(sym_string),
	3104:  uint16(73),
	3105:  uint16(5),
	3106:  uint16(sym__linespace),
	3107:  uint16(sym__newline),
	3108:  uint16(sym__ws),
	3109:  uint16(sym_single_line_comment),
	3110:  uint16(aux_sym_document_repeat1),
	3111:  uint16(165),
	3112:  uint16(10),
	3113:  uint16(sym_multi_line_comment),
	3114:  uint16(aux_sym__newline_token1),
	3115:  uint16(aux_sym__newline_token2),
	3116:  uint16(aux_sym__newline_token3),
	3117:  uint16(aux_sym__newline_token4),
	3118:  uint16(aux_sym__newline_token5),
	3119:  uint16(aux_sym__newline_token6),
	3120:  uint16(aux_sym__newline_token7),
	3121:  uint16(sym__bom),
	3122:  uint16(sym__unicode_space),
	3123:  uint16(16),
	3124:  uint16(7),
	3125:  uint16(1),
	3126:  uint16(sym__normal_bare_identifier),
	3127:  uint16(9),
	3128:  uint16(1),
	3129:  uint16(anon_sym_SLASH_DASH),
	3130:  uint16(11),
	3131:  uint16(1),
	3132:  uint16(anon_sym_LPAREN),
	3133:  uint16(13),
	3134:  uint16(1),
	3135:  uint16(anon_sym_DQUOTE),
	3136:  uint16(19),
	3137:  uint16(1),
	3138:  uint16(anon_sym_SLASH_SLASH),
	3139:  uint16(21),
	3140:  uint16(1),
	3141:  uint16(sym__raw_string),
	3142:  uint16(135),
	3143:  uint16(1),
	3144:  uint16(anon_sym_RBRACE),
	3145:  uint16(24),
	3146:  uint16(1),
	3147:  uint16(sym_node),
	3148:  uint16(71),
	3149:  uint16(1),
	3150:  uint16(sym_identifier),
	3151:  uint16(192),
	3152:  uint16(1),
	3153:  uint16(sym__sign),
	3154:  uint16(194),
	3155:  uint16(1),
	3156:  uint16(sym__escaped_string),
	3157:  uint16(260),
	3158:  uint16(1),
	3159:  uint16(sym_type),
	3160:  uint16(15),
	3161:  uint16(2),
	3162:  uint16(anon_sym_PLUS),
	3163:  uint16(anon_sym_DASH),
	3164:  uint16(218),
	3165:  uint16(2),
	3166:  uint16(sym__bare_identifier),
	3167:  uint16(sym_string),
	3168:  uint16(76),
	3169:  uint16(5),
	3170:  uint16(sym__linespace),
	3171:  uint16(sym__newline),
	3172:  uint16(sym__ws),
	3173:  uint16(sym_single_line_comment),
	3174:  uint16(aux_sym_document_repeat1),
	3175:  uint16(194),
	3176:  uint16(10),
	3177:  uint16(sym_multi_line_comment),
	3178:  uint16(aux_sym__newline_token1),
	3179:  uint16(aux_sym__newline_token2),
	3180:  uint16(aux_sym__newline_token3),
	3181:  uint16(aux_sym__newline_token4),
	3182:  uint16(aux_sym__newline_token5),
	3183:  uint16(aux_sym__newline_token6),
	3184:  uint16(aux_sym__newline_token7),
	3185:  uint16(sym__bom),
	3186:  uint16(sym__unicode_space),
	3187:  uint16(16),
	3188:  uint16(7),
	3189:  uint16(1),
	3190:  uint16(sym__normal_bare_identifier),
	3191:  uint16(9),
	3192:  uint16(1),
	3193:  uint16(anon_sym_SLASH_DASH),
	3194:  uint16(11),
	3195:  uint16(1),
	3196:  uint16(anon_sym_LPAREN),
	3197:  uint16(13),
	3198:  uint16(1),
	3199:  uint16(anon_sym_DQUOTE),
	3200:  uint16(19),
	3201:  uint16(1),
	3202:  uint16(anon_sym_SLASH_SLASH),
	3203:  uint16(21),
	3204:  uint16(1),
	3205:  uint16(sym__raw_string),
	3206:  uint16(157),
	3207:  uint16(1),
	3209:  uint16(71),
	3210:  uint16(1),
	3211:  uint16(sym_identifier),
	3212:  uint16(153),
	3213:  uint16(1),
	3214:  uint16(sym_node),
	3215:  uint16(192),
	3216:  uint16(1),
	3217:  uint16(sym__sign),
	3218:  uint16(194),
	3219:  uint16(1),
	3220:  uint16(sym__escaped_string),
	3221:  uint16(260),
	3222:  uint16(1),
	3223:  uint16(sym_type),
	3224:  uint16(15),
	3225:  uint16(2),
	3226:  uint16(anon_sym_PLUS),
	3227:  uint16(anon_sym_DASH),
	3228:  uint16(218),
	3229:  uint16(2),
	3230:  uint16(sym__bare_identifier),
	3231:  uint16(sym_string),
	3232:  uint16(73),
	3233:  uint16(5),
	3234:  uint16(sym__linespace),
	3235:  uint16(sym__newline),
	3236:  uint16(sym__ws),
	3237:  uint16(sym_single_line_comment),
	3238:  uint16(aux_sym_document_repeat1),
	3239:  uint16(165),
	3240:  uint16(10),
	3241:  uint16(sym_multi_line_comment),
	3242:  uint16(aux_sym__newline_token1),
	3243:  uint16(aux_sym__newline_token2),
	3244:  uint16(aux_sym__newline_token3),
	3245:  uint16(aux_sym__newline_token4),
	3246:  uint16(aux_sym__newline_token5),
	3247:  uint16(aux_sym__newline_token6),
	3248:  uint16(aux_sym__newline_token7),
	3249:  uint16(sym__bom),
	3250:  uint16(sym__unicode_space),
	3251:  uint16(16),
	3252:  uint16(7),
	3253:  uint16(1),
	3254:  uint16(sym__normal_bare_identifier),
	3255:  uint16(9),
	3256:  uint16(1),
	3257:  uint16(anon_sym_SLASH_DASH),
	3258:  uint16(11),
	3259:  uint16(1),
	3260:  uint16(anon_sym_LPAREN),
	3261:  uint16(13),
	3262:  uint16(1),
	3263:  uint16(anon_sym_DQUOTE),
	3264:  uint16(19),
	3265:  uint16(1),
	3266:  uint16(anon_sym_SLASH_SLASH),
	3267:  uint16(21),
	3268:  uint16(1),
	3269:  uint16(sym__raw_string),
	3270:  uint16(131),
	3271:  uint16(1),
	3272:  uint16(anon_sym_RBRACE),
	3273:  uint16(18),
	3274:  uint16(1),
	3275:  uint16(sym_node),
	3276:  uint16(71),
	3277:  uint16(1),
	3278:  uint16(sym_identifier),
	3279:  uint16(192),
	3280:  uint16(1),
	3281:  uint16(sym__sign),
	3282:  uint16(194),
	3283:  uint16(1),
	3284:  uint16(sym__escaped_string),
	3285:  uint16(260),
	3286:  uint16(1),
	3287:  uint16(sym_type),
	3288:  uint16(15),
	3289:  uint16(2),
	3290:  uint16(anon_sym_PLUS),
	3291:  uint16(anon_sym_DASH),
	3292:  uint16(218),
	3293:  uint16(2),
	3294:  uint16(sym__bare_identifier),
	3295:  uint16(sym_string),
	3296:  uint16(75),
	3297:  uint16(5),
	3298:  uint16(sym__linespace),
	3299:  uint16(sym__newline),
	3300:  uint16(sym__ws),
	3301:  uint16(sym_single_line_comment),
	3302:  uint16(aux_sym_document_repeat1),
	3303:  uint16(196),
	3304:  uint16(10),
	3305:  uint16(sym_multi_line_comment),
	3306:  uint16(aux_sym__newline_token1),
	3307:  uint16(aux_sym__newline_token2),
	3308:  uint16(aux_sym__newline_token3),
	3309:  uint16(aux_sym__newline_token4),
	3310:  uint16(aux_sym__newline_token5),
	3311:  uint16(aux_sym__newline_token6),
	3312:  uint16(aux_sym__newline_token7),
	3313:  uint16(sym__bom),
	3314:  uint16(sym__unicode_space),
	3315:  uint16(16),
	3316:  uint16(7),
	3317:  uint16(1),
	3318:  uint16(sym__normal_bare_identifier),
	3319:  uint16(9),
	3320:  uint16(1),
	3321:  uint16(anon_sym_SLASH_DASH),
	3322:  uint16(11),
	3323:  uint16(1),
	3324:  uint16(anon_sym_LPAREN),
	3325:  uint16(13),
	3326:  uint16(1),
	3327:  uint16(anon_sym_DQUOTE),
	3328:  uint16(19),
	3329:  uint16(1),
	3330:  uint16(anon_sym_SLASH_SLASH),
	3331:  uint16(21),
	3332:  uint16(1),
	3333:  uint16(sym__raw_string),
	3334:  uint16(115),
	3335:  uint16(1),
	3337:  uint16(28),
	3338:  uint16(1),
	3339:  uint16(sym_node),
	3340:  uint16(71),
	3341:  uint16(1),
	3342:  uint16(sym_identifier),
	3343:  uint16(192),
	3344:  uint16(1),
	3345:  uint16(sym__sign),
	3346:  uint16(194),
	3347:  uint16(1),
	3348:  uint16(sym__escaped_string),
	3349:  uint16(260),
	3350:  uint16(1),
	3351:  uint16(sym_type),
	3352:  uint16(15),
	3353:  uint16(2),
	3354:  uint16(anon_sym_PLUS),
	3355:  uint16(anon_sym_DASH),
	3356:  uint16(218),
	3357:  uint16(2),
	3358:  uint16(sym__bare_identifier),
	3359:  uint16(sym_string),
	3360:  uint16(77),
	3361:  uint16(5),
	3362:  uint16(sym__linespace),
	3363:  uint16(sym__newline),
	3364:  uint16(sym__ws),
	3365:  uint16(sym_single_line_comment),
	3366:  uint16(aux_sym_document_repeat1),
	3367:  uint16(198),
	3368:  uint16(10),
	3369:  uint16(sym_multi_line_comment),
	3370:  uint16(aux_sym__newline_token1),
	3371:  uint16(aux_sym__newline_token2),
	3372:  uint16(aux_sym__newline_token3),
	3373:  uint16(aux_sym__newline_token4),
	3374:  uint16(aux_sym__newline_token5),
	3375:  uint16(aux_sym__newline_token6),
	3376:  uint16(aux_sym__newline_token7),
	3377:  uint16(sym__bom),
	3378:  uint16(sym__unicode_space),
	3379:  uint16(16),
	3380:  uint16(7),
	3381:  uint16(1),
	3382:  uint16(sym__normal_bare_identifier),
	3383:  uint16(9),
	3384:  uint16(1),
	3385:  uint16(anon_sym_SLASH_DASH),
	3386:  uint16(11),
	3387:  uint16(1),
	3388:  uint16(anon_sym_LPAREN),
	3389:  uint16(13),
	3390:  uint16(1),
	3391:  uint16(anon_sym_DQUOTE),
	3392:  uint16(19),
	3393:  uint16(1),
	3394:  uint16(anon_sym_SLASH_SLASH),
	3395:  uint16(21),
	3396:  uint16(1),
	3397:  uint16(sym__raw_string),
	3398:  uint16(200),
	3399:  uint16(1),
	3400:  uint16(anon_sym_RBRACE),
	3401:  uint16(71),
	3402:  uint16(1),
	3403:  uint16(sym_identifier),
	3404:  uint16(153),
	3405:  uint16(1),
	3406:  uint16(sym_node),
	3407:  uint16(192),
	3408:  uint16(1),
	3409:  uint16(sym__sign),
	3410:  uint16(194),
	3411:  uint16(1),
	3412:  uint16(sym__escaped_string),
	3413:  uint16(260),
	3414:  uint16(1),
	3415:  uint16(sym_type),
	3416:  uint16(15),
	3417:  uint16(2),
	3418:  uint16(anon_sym_PLUS),
	3419:  uint16(anon_sym_DASH),
	3420:  uint16(218),
	3421:  uint16(2),
	3422:  uint16(sym__bare_identifier),
	3423:  uint16(sym_string),
	3424:  uint16(73),
	3425:  uint16(5),
	3426:  uint16(sym__linespace),
	3427:  uint16(sym__newline),
	3428:  uint16(sym__ws),
	3429:  uint16(sym_single_line_comment),
	3430:  uint16(aux_sym_document_repeat1),
	3431:  uint16(165),
	3432:  uint16(10),
	3433:  uint16(sym_multi_line_comment),
	3434:  uint16(aux_sym__newline_token1),
	3435:  uint16(aux_sym__newline_token2),
	3436:  uint16(aux_sym__newline_token3),
	3437:  uint16(aux_sym__newline_token4),
	3438:  uint16(aux_sym__newline_token5),
	3439:  uint16(aux_sym__newline_token6),
	3440:  uint16(aux_sym__newline_token7),
	3441:  uint16(sym__bom),
	3442:  uint16(sym__unicode_space),
	3443:  uint16(16),
	3444:  uint16(7),
	3445:  uint16(1),
	3446:  uint16(sym__normal_bare_identifier),
	3447:  uint16(9),
	3448:  uint16(1),
	3449:  uint16(anon_sym_SLASH_DASH),
	3450:  uint16(11),
	3451:  uint16(1),
	3452:  uint16(anon_sym_LPAREN),
	3453:  uint16(13),
	3454:  uint16(1),
	3455:  uint16(anon_sym_DQUOTE),
	3456:  uint16(19),
	3457:  uint16(1),
	3458:  uint16(anon_sym_SLASH_SLASH),
	3459:  uint16(21),
	3460:  uint16(1),
	3461:  uint16(sym__raw_string),
	3462:  uint16(119),
	3463:  uint16(1),
	3464:  uint16(anon_sym_RBRACE),
	3465:  uint16(71),
	3466:  uint16(1),
	3467:  uint16(sym_identifier),
	3468:  uint16(153),
	3469:  uint16(1),
	3470:  uint16(sym_node),
	3471:  uint16(192),
	3472:  uint16(1),
	3473:  uint16(sym__sign),
	3474:  uint16(194),
	3475:  uint16(1),
	3476:  uint16(sym__escaped_string),
	3477:  uint16(260),
	3478:  uint16(1),
	3479:  uint16(sym_type),
	3480:  uint16(15),
	3481:  uint16(2),
	3482:  uint16(anon_sym_PLUS),
	3483:  uint16(anon_sym_DASH),
	3484:  uint16(218),
	3485:  uint16(2),
	3486:  uint16(sym__bare_identifier),
	3487:  uint16(sym_string),
	3488:  uint16(73),
	3489:  uint16(5),
	3490:  uint16(sym__linespace),
	3491:  uint16(sym__newline),
	3492:  uint16(sym__ws),
	3493:  uint16(sym_single_line_comment),
	3494:  uint16(aux_sym_document_repeat1),
	3495:  uint16(165),
	3496:  uint16(10),
	3497:  uint16(sym_multi_line_comment),
	3498:  uint16(aux_sym__newline_token1),
	3499:  uint16(aux_sym__newline_token2),
	3500:  uint16(aux_sym__newline_token3),
	3501:  uint16(aux_sym__newline_token4),
	3502:  uint16(aux_sym__newline_token5),
	3503:  uint16(aux_sym__newline_token6),
	3504:  uint16(aux_sym__newline_token7),
	3505:  uint16(sym__bom),
	3506:  uint16(sym__unicode_space),
	3507:  uint16(16),
	3508:  uint16(7),
	3509:  uint16(1),
	3510:  uint16(sym__normal_bare_identifier),
	3511:  uint16(9),
	3512:  uint16(1),
	3513:  uint16(anon_sym_SLASH_DASH),
	3514:  uint16(11),
	3515:  uint16(1),
	3516:  uint16(anon_sym_LPAREN),
	3517:  uint16(13),
	3518:  uint16(1),
	3519:  uint16(anon_sym_DQUOTE),
	3520:  uint16(19),
	3521:  uint16(1),
	3522:  uint16(anon_sym_SLASH_SLASH),
	3523:  uint16(21),
	3524:  uint16(1),
	3525:  uint16(sym__raw_string),
	3526:  uint16(202),
	3527:  uint16(1),
	3528:  uint16(anon_sym_RBRACE),
	3529:  uint16(23),
	3530:  uint16(1),
	3531:  uint16(sym_node),
	3532:  uint16(71),
	3533:  uint16(1),
	3534:  uint16(sym_identifier),
	3535:  uint16(192),
	3536:  uint16(1),
	3537:  uint16(sym__sign),
	3538:  uint16(194),
	3539:  uint16(1),
	3540:  uint16(sym__escaped_string),
	3541:  uint16(260),
	3542:  uint16(1),
	3543:  uint16(sym_type),
	3544:  uint16(15),
	3545:  uint16(2),
	3546:  uint16(anon_sym_PLUS),
	3547:  uint16(anon_sym_DASH),
	3548:  uint16(218),
	3549:  uint16(2),
	3550:  uint16(sym__bare_identifier),
	3551:  uint16(sym_string),
	3552:  uint16(42),
	3553:  uint16(5),
	3554:  uint16(sym__linespace),
	3555:  uint16(sym__newline),
	3556:  uint16(sym__ws),
	3557:  uint16(sym_single_line_comment),
	3558:  uint16(aux_sym_document_repeat1),
	3559:  uint16(204),
	3560:  uint16(10),
	3561:  uint16(sym_multi_line_comment),
	3562:  uint16(aux_sym__newline_token1),
	3563:  uint16(aux_sym__newline_token2),
	3564:  uint16(aux_sym__newline_token3),
	3565:  uint16(aux_sym__newline_token4),
	3566:  uint16(aux_sym__newline_token5),
	3567:  uint16(aux_sym__newline_token6),
	3568:  uint16(aux_sym__newline_token7),
	3569:  uint16(sym__bom),
	3570:  uint16(sym__unicode_space),
	3571:  uint16(4),
	3572:  uint16(49),
	3573:  uint16(2),
	3574:  uint16(sym__ws),
	3575:  uint16(aux_sym_node_repeat3),
	3576:  uint16(210),
	3577:  uint16(3),
	3578:  uint16(sym_multi_line_comment),
	3579:  uint16(sym__bom),
	3580:  uint16(sym__unicode_space),
	3581:  uint16(206),
	3582:  uint16(5),
	3583:  uint16(sym__normal_bare_identifier),
	3584:  uint16(anon_sym_null),
	3585:  uint16(sym__digit),
	3586:  uint16(anon_sym_true),
	3587:  uint16(anon_sym_false),
	3588:  uint16(208),
	3589:  uint16(21),
	3590:  uint16(sym__eof),
	3591:  uint16(sym__raw_string),
	3592:  uint16(anon_sym_SLASH_DASH),
	3593:  uint16(anon_sym_LBRACE),
	3594:  uint16(anon_sym_SEMI),
	3595:  uint16(anon_sym_LPAREN),
	3596:  uint16(anon_sym_DQUOTE),
	3597:  uint16(anon_sym_PLUS),
	3598:  uint16(anon_sym_DASH),
	3599:  uint16(anon_sym_0x),
	3600:  uint16(anon_sym_0o),
	3601:  uint16(anon_sym_0b),
	3602:  uint16(anon_sym_BSLASH),
	3603:  uint16(aux_sym__newline_token1),
	3604:  uint16(aux_sym__newline_token2),
	3605:  uint16(aux_sym__newline_token3),
	3606:  uint16(aux_sym__newline_token4),
	3607:  uint16(aux_sym__newline_token5),
	3608:  uint16(aux_sym__newline_token6),
	3609:  uint16(aux_sym__newline_token7),
	3610:  uint16(anon_sym_SLASH_SLASH),
	3611:  uint16(16),
	3612:  uint16(7),
	3613:  uint16(1),
	3614:  uint16(sym__normal_bare_identifier),
	3615:  uint16(9),
	3616:  uint16(1),
	3617:  uint16(anon_sym_SLASH_DASH),
	3618:  uint16(11),
	3619:  uint16(1),
	3620:  uint16(anon_sym_LPAREN),
	3621:  uint16(13),
	3622:  uint16(1),
	3623:  uint16(anon_sym_DQUOTE),
	3624:  uint16(19),
	3625:  uint16(1),
	3626:  uint16(anon_sym_SLASH_SLASH),
	3627:  uint16(21),
	3628:  uint16(1),
	3629:  uint16(sym__raw_string),
	3630:  uint16(213),
	3631:  uint16(1),
	3633:  uint16(71),
	3634:  uint16(1),
	3635:  uint16(sym_identifier),
	3636:  uint16(153),
	3637:  uint16(1),
	3638:  uint16(sym_node),
	3639:  uint16(192),
	3640:  uint16(1),
	3641:  uint16(sym__sign),
	3642:  uint16(194),
	3643:  uint16(1),
	3644:  uint16(sym__escaped_string),
	3645:  uint16(260),
	3646:  uint16(1),
	3647:  uint16(sym_type),
	3648:  uint16(15),
	3649:  uint16(2),
	3650:  uint16(anon_sym_PLUS),
	3651:  uint16(anon_sym_DASH),
	3652:  uint16(218),
	3653:  uint16(2),
	3654:  uint16(sym__bare_identifier),
	3655:  uint16(sym_string),
	3656:  uint16(73),
	3657:  uint16(5),
	3658:  uint16(sym__linespace),
	3659:  uint16(sym__newline),
	3660:  uint16(sym__ws),
	3661:  uint16(sym_single_line_comment),
	3662:  uint16(aux_sym_document_repeat1),
	3663:  uint16(165),
	3664:  uint16(10),
	3665:  uint16(sym_multi_line_comment),
	3666:  uint16(aux_sym__newline_token1),
	3667:  uint16(aux_sym__newline_token2),
	3668:  uint16(aux_sym__newline_token3),
	3669:  uint16(aux_sym__newline_token4),
	3670:  uint16(aux_sym__newline_token5),
	3671:  uint16(aux_sym__newline_token6),
	3672:  uint16(aux_sym__newline_token7),
	3673:  uint16(sym__bom),
	3674:  uint16(sym__unicode_space),
	3675:  uint16(16),
	3676:  uint16(7),
	3677:  uint16(1),
	3678:  uint16(sym__normal_bare_identifier),
	3679:  uint16(9),
	3680:  uint16(1),
	3681:  uint16(anon_sym_SLASH_DASH),
	3682:  uint16(11),
	3683:  uint16(1),
	3684:  uint16(anon_sym_LPAREN),
	3685:  uint16(13),
	3686:  uint16(1),
	3687:  uint16(anon_sym_DQUOTE),
	3688:  uint16(19),
	3689:  uint16(1),
	3690:  uint16(anon_sym_SLASH_SLASH),
	3691:  uint16(21),
	3692:  uint16(1),
	3693:  uint16(sym__raw_string),
	3694:  uint16(127),
	3695:  uint16(1),
	3696:  uint16(anon_sym_RBRACE),
	3697:  uint16(71),
	3698:  uint16(1),
	3699:  uint16(sym_identifier),
	3700:  uint16(153),
	3701:  uint16(1),
	3702:  uint16(sym_node),
	3703:  uint16(192),
	3704:  uint16(1),
	3705:  uint16(sym__sign),
	3706:  uint16(194),
	3707:  uint16(1),
	3708:  uint16(sym__escaped_string),
	3709:  uint16(260),
	3710:  uint16(1),
	3711:  uint16(sym_type),
	3712:  uint16(15),
	3713:  uint16(2),
	3714:  uint16(anon_sym_PLUS),
	3715:  uint16(anon_sym_DASH),
	3716:  uint16(218),
	3717:  uint16(2),
	3718:  uint16(sym__bare_identifier),
	3719:  uint16(sym_string),
	3720:  uint16(73),
	3721:  uint16(5),
	3722:  uint16(sym__linespace),
	3723:  uint16(sym__newline),
	3724:  uint16(sym__ws),
	3725:  uint16(sym_single_line_comment),
	3726:  uint16(aux_sym_document_repeat1),
	3727:  uint16(165),
	3728:  uint16(10),
	3729:  uint16(sym_multi_line_comment),
	3730:  uint16(aux_sym__newline_token1),
	3731:  uint16(aux_sym__newline_token2),
	3732:  uint16(aux_sym__newline_token3),
	3733:  uint16(aux_sym__newline_token4),
	3734:  uint16(aux_sym__newline_token5),
	3735:  uint16(aux_sym__newline_token6),
	3736:  uint16(aux_sym__newline_token7),
	3737:  uint16(sym__bom),
	3738:  uint16(sym__unicode_space),
	3739:  uint16(16),
	3740:  uint16(7),
	3741:  uint16(1),
	3742:  uint16(sym__normal_bare_identifier),
	3743:  uint16(9),
	3744:  uint16(1),
	3745:  uint16(anon_sym_SLASH_DASH),
	3746:  uint16(11),
	3747:  uint16(1),
	3748:  uint16(anon_sym_LPAREN),
	3749:  uint16(13),
	3750:  uint16(1),
	3751:  uint16(anon_sym_DQUOTE),
	3752:  uint16(19),
	3753:  uint16(1),
	3754:  uint16(anon_sym_SLASH_SLASH),
	3755:  uint16(21),
	3756:  uint16(1),
	3757:  uint16(sym__raw_string),
	3758:  uint16(123),
	3759:  uint16(1),
	3761:  uint16(71),
	3762:  uint16(1),
	3763:  uint16(sym_identifier),
	3764:  uint16(153),
	3765:  uint16(1),
	3766:  uint16(sym_node),
	3767:  uint16(192),
	3768:  uint16(1),
	3769:  uint16(sym__sign),
	3770:  uint16(194),
	3771:  uint16(1),
	3772:  uint16(sym__escaped_string),
	3773:  uint16(260),
	3774:  uint16(1),
	3775:  uint16(sym_type),
	3776:  uint16(15),
	3777:  uint16(2),
	3778:  uint16(anon_sym_PLUS),
	3779:  uint16(anon_sym_DASH),
	3780:  uint16(218),
	3781:  uint16(2),
	3782:  uint16(sym__bare_identifier),
	3783:  uint16(sym_string),
	3784:  uint16(73),
	3785:  uint16(5),
	3786:  uint16(sym__linespace),
	3787:  uint16(sym__newline),
	3788:  uint16(sym__ws),
	3789:  uint16(sym_single_line_comment),
	3790:  uint16(aux_sym_document_repeat1),
	3791:  uint16(165),
	3792:  uint16(10),
	3793:  uint16(sym_multi_line_comment),
	3794:  uint16(aux_sym__newline_token1),
	3795:  uint16(aux_sym__newline_token2),
	3796:  uint16(aux_sym__newline_token3),
	3797:  uint16(aux_sym__newline_token4),
	3798:  uint16(aux_sym__newline_token5),
	3799:  uint16(aux_sym__newline_token6),
	3800:  uint16(aux_sym__newline_token7),
	3801:  uint16(sym__bom),
	3802:  uint16(sym__unicode_space),
	3803:  uint16(4),
	3804:  uint16(49),
	3805:  uint16(2),
	3806:  uint16(sym__ws),
	3807:  uint16(aux_sym_node_repeat3),
	3808:  uint16(215),
	3809:  uint16(3),
	3810:  uint16(sym_multi_line_comment),
	3811:  uint16(sym__bom),
	3812:  uint16(sym__unicode_space),
	3813:  uint16(185),
	3814:  uint16(5),
	3815:  uint16(sym__normal_bare_identifier),
	3816:  uint16(anon_sym_null),
	3817:  uint16(sym__digit),
	3818:  uint16(anon_sym_true),
	3819:  uint16(anon_sym_false),
	3820:  uint16(187),
	3821:  uint16(21),
	3822:  uint16(sym__eof),
	3823:  uint16(sym__raw_string),
	3824:  uint16(anon_sym_SLASH_DASH),
	3825:  uint16(anon_sym_LBRACE),
	3826:  uint16(anon_sym_SEMI),
	3827:  uint16(anon_sym_LPAREN),
	3828:  uint16(anon_sym_DQUOTE),
	3829:  uint16(anon_sym_PLUS),
	3830:  uint16(anon_sym_DASH),
	3831:  uint16(anon_sym_0x),
	3832:  uint16(anon_sym_0o),
	3833:  uint16(anon_sym_0b),
	3834:  uint16(anon_sym_BSLASH),
	3835:  uint16(aux_sym__newline_token1),
	3836:  uint16(aux_sym__newline_token2),
	3837:  uint16(aux_sym__newline_token3),
	3838:  uint16(aux_sym__newline_token4),
	3839:  uint16(aux_sym__newline_token5),
	3840:  uint16(aux_sym__newline_token6),
	3841:  uint16(aux_sym__newline_token7),
	3842:  uint16(anon_sym_SLASH_SLASH),
	3843:  uint16(15),
	3844:  uint16(7),
	3845:  uint16(1),
	3846:  uint16(sym__normal_bare_identifier),
	3847:  uint16(9),
	3848:  uint16(1),
	3849:  uint16(anon_sym_SLASH_DASH),
	3850:  uint16(11),
	3851:  uint16(1),
	3852:  uint16(anon_sym_LPAREN),
	3853:  uint16(13),
	3854:  uint16(1),
	3855:  uint16(anon_sym_DQUOTE),
	3856:  uint16(19),
	3857:  uint16(1),
	3858:  uint16(anon_sym_SLASH_SLASH),
	3859:  uint16(21),
	3860:  uint16(1),
	3861:  uint16(sym__raw_string),
	3862:  uint16(71),
	3863:  uint16(1),
	3864:  uint16(sym_identifier),
	3865:  uint16(153),
	3866:  uint16(1),
	3867:  uint16(sym_node),
	3868:  uint16(192),
	3869:  uint16(1),
	3870:  uint16(sym__sign),
	3871:  uint16(194),
	3872:  uint16(1),
	3873:  uint16(sym__escaped_string),
	3874:  uint16(260),
	3875:  uint16(1),
	3876:  uint16(sym_type),
	3877:  uint16(15),
	3878:  uint16(2),
	3879:  uint16(anon_sym_PLUS),
	3880:  uint16(anon_sym_DASH),
	3881:  uint16(218),
	3882:  uint16(2),
	3883:  uint16(sym__bare_identifier),
	3884:  uint16(sym_string),
	3885:  uint16(73),
	3886:  uint16(5),
	3887:  uint16(sym__linespace),
	3888:  uint16(sym__newline),
	3889:  uint16(sym__ws),
	3890:  uint16(sym_single_line_comment),
	3891:  uint16(aux_sym_document_repeat1),
	3892:  uint16(165),
	3893:  uint16(10),
	3894:  uint16(sym_multi_line_comment),
	3895:  uint16(aux_sym__newline_token1),
	3896:  uint16(aux_sym__newline_token2),
	3897:  uint16(aux_sym__newline_token3),
	3898:  uint16(aux_sym__newline_token4),
	3899:  uint16(aux_sym__newline_token5),
	3900:  uint16(aux_sym__newline_token6),
	3901:  uint16(aux_sym__newline_token7),
	3902:  uint16(sym__bom),
	3903:  uint16(sym__unicode_space),
	3904:  uint16(2),
	3905:  uint16(220),
	3906:  uint16(5),
	3907:  uint16(sym__normal_bare_identifier),
	3908:  uint16(anon_sym_null),
	3909:  uint16(sym__digit),
	3910:  uint16(anon_sym_true),
	3911:  uint16(anon_sym_false),
	3912:  uint16(218),
	3913:  uint16(24),
	3914:  uint16(sym_multi_line_comment),
	3915:  uint16(sym__raw_string),
	3917:  uint16(anon_sym_SLASH_DASH),
	3918:  uint16(anon_sym_LBRACE),
	3919:  uint16(anon_sym_RBRACE),
	3920:  uint16(anon_sym_LPAREN),
	3921:  uint16(anon_sym_DQUOTE),
	3922:  uint16(anon_sym_PLUS),
	3923:  uint16(anon_sym_DASH),
	3924:  uint16(anon_sym_0x),
	3925:  uint16(anon_sym_0o),
	3926:  uint16(anon_sym_0b),
	3927:  uint16(anon_sym_BSLASH),
	3928:  uint16(aux_sym__newline_token1),
	3929:  uint16(aux_sym__newline_token2),
	3930:  uint16(aux_sym__newline_token3),
	3931:  uint16(aux_sym__newline_token4),
	3932:  uint16(aux_sym__newline_token5),
	3933:  uint16(aux_sym__newline_token6),
	3934:  uint16(aux_sym__newline_token7),
	3935:  uint16(sym__bom),
	3936:  uint16(sym__unicode_space),
	3937:  uint16(anon_sym_SLASH_SLASH),
	3938:  uint16(2),
	3939:  uint16(222),
	3940:  uint16(5),
	3941:  uint16(sym__normal_bare_identifier),
	3942:  uint16(anon_sym_null),
	3943:  uint16(sym__digit),
	3944:  uint16(anon_sym_true),
	3945:  uint16(anon_sym_false),
	3946:  uint16(224),
	3947:  uint16(24),
	3948:  uint16(sym__eof),
	3949:  uint16(sym_multi_line_comment),
	3950:  uint16(sym__raw_string),
	3951:  uint16(anon_sym_SLASH_DASH),
	3952:  uint16(anon_sym_LBRACE),
	3953:  uint16(anon_sym_SEMI),
	3954:  uint16(anon_sym_LPAREN),
	3955:  uint16(anon_sym_DQUOTE),
	3956:  uint16(anon_sym_PLUS),
	3957:  uint16(anon_sym_DASH),
	3958:  uint16(anon_sym_0x),
	3959:  uint16(anon_sym_0o),
	3960:  uint16(anon_sym_0b),
	3961:  uint16(anon_sym_BSLASH),
	3962:  uint16(aux_sym__newline_token1),
	3963:  uint16(aux_sym__newline_token2),
	3964:  uint16(aux_sym__newline_token3),
	3965:  uint16(aux_sym__newline_token4),
	3966:  uint16(aux_sym__newline_token5),
	3967:  uint16(aux_sym__newline_token6),
	3968:  uint16(aux_sym__newline_token7),
	3969:  uint16(sym__bom),
	3970:  uint16(sym__unicode_space),
	3971:  uint16(anon_sym_SLASH_SLASH),
	3972:  uint16(2),
	3973:  uint16(226),
	3974:  uint16(5),
	3975:  uint16(sym__normal_bare_identifier),
	3976:  uint16(anon_sym_null),
	3977:  uint16(sym__digit),
	3978:  uint16(anon_sym_true),
	3979:  uint16(anon_sym_false),
	3980:  uint16(228),
	3981:  uint16(24),
	3982:  uint16(sym__eof),
	3983:  uint16(sym_multi_line_comment),
	3984:  uint16(sym__raw_string),
	3985:  uint16(anon_sym_SLASH_DASH),
	3986:  uint16(anon_sym_LBRACE),
	3987:  uint16(anon_sym_SEMI),
	3988:  uint16(anon_sym_LPAREN),
	3989:  uint16(anon_sym_DQUOTE),
	3990:  uint16(anon_sym_PLUS),
	3991:  uint16(anon_sym_DASH),
	3992:  uint16(anon_sym_0x),
	3993:  uint16(anon_sym_0o),
	3994:  uint16(anon_sym_0b),
	3995:  uint16(anon_sym_BSLASH),
	3996:  uint16(aux_sym__newline_token1),
	3997:  uint16(aux_sym__newline_token2),
	3998:  uint16(aux_sym__newline_token3),
	3999:  uint16(aux_sym__newline_token4),
	4000:  uint16(aux_sym__newline_token5),
	4001:  uint16(aux_sym__newline_token6),
	4002:  uint16(aux_sym__newline_token7),
	4003:  uint16(sym__bom),
	4004:  uint16(sym__unicode_space),
	4005:  uint16(anon_sym_SLASH_SLASH),
	4006:  uint16(2),
	4007:  uint16(230),
	4008:  uint16(5),
	4009:  uint16(sym__normal_bare_identifier),
	4010:  uint16(anon_sym_null),
	4011:  uint16(sym__digit),
	4012:  uint16(anon_sym_true),
	4013:  uint16(anon_sym_false),
	4014:  uint16(232),
	4015:  uint16(24),
	4016:  uint16(sym__eof),
	4017:  uint16(sym_multi_line_comment),
	4018:  uint16(sym__raw_string),
	4019:  uint16(anon_sym_SLASH_DASH),
	4020:  uint16(anon_sym_LBRACE),
	4021:  uint16(anon_sym_SEMI),
	4022:  uint16(anon_sym_LPAREN),
	4023:  uint16(anon_sym_DQUOTE),
	4024:  uint16(anon_sym_PLUS),
	4025:  uint16(anon_sym_DASH),
	4026:  uint16(anon_sym_0x),
	4027:  uint16(anon_sym_0o),
	4028:  uint16(anon_sym_0b),
	4029:  uint16(anon_sym_BSLASH),
	4030:  uint16(aux_sym__newline_token1),
	4031:  uint16(aux_sym__newline_token2),
	4032:  uint16(aux_sym__newline_token3),
	4033:  uint16(aux_sym__newline_token4),
	4034:  uint16(aux_sym__newline_token5),
	4035:  uint16(aux_sym__newline_token6),
	4036:  uint16(aux_sym__newline_token7),
	4037:  uint16(sym__bom),
	4038:  uint16(sym__unicode_space),
	4039:  uint16(anon_sym_SLASH_SLASH),
	4040:  uint16(2),
	4041:  uint16(220),
	4042:  uint16(5),
	4043:  uint16(sym__normal_bare_identifier),
	4044:  uint16(anon_sym_null),
	4045:  uint16(sym__digit),
	4046:  uint16(anon_sym_true),
	4047:  uint16(anon_sym_false),
	4048:  uint16(218),
	4049:  uint16(24),
	4050:  uint16(sym__eof),
	4051:  uint16(sym_multi_line_comment),
	4052:  uint16(sym__raw_string),
	4053:  uint16(anon_sym_SLASH_DASH),
	4054:  uint16(anon_sym_LBRACE),
	4055:  uint16(anon_sym_SEMI),
	4056:  uint16(anon_sym_LPAREN),
	4057:  uint16(anon_sym_DQUOTE),
	4058:  uint16(anon_sym_PLUS),
	4059:  uint16(anon_sym_DASH),
	4060:  uint16(anon_sym_0x),
	4061:  uint16(anon_sym_0o),
	4062:  uint16(anon_sym_0b),
	4063:  uint16(anon_sym_BSLASH),
	4064:  uint16(aux_sym__newline_token1),
	4065:  uint16(aux_sym__newline_token2),
	4066:  uint16(aux_sym__newline_token3),
	4067:  uint16(aux_sym__newline_token4),
	4068:  uint16(aux_sym__newline_token5),
	4069:  uint16(aux_sym__newline_token6),
	4070:  uint16(aux_sym__newline_token7),
	4071:  uint16(sym__bom),
	4072:  uint16(sym__unicode_space),
	4073:  uint16(anon_sym_SLASH_SLASH),
	4074:  uint16(2),
	4075:  uint16(236),
	4076:  uint16(5),
	4077:  uint16(sym__normal_bare_identifier),
	4078:  uint16(anon_sym_null),
	4079:  uint16(sym__digit),
	4080:  uint16(anon_sym_true),
	4081:  uint16(anon_sym_false),
	4082:  uint16(234),
	4083:  uint16(24),
	4084:  uint16(sym_multi_line_comment),
	4085:  uint16(sym__raw_string),
	4087:  uint16(anon_sym_SLASH_DASH),
	4088:  uint16(anon_sym_LBRACE),
	4089:  uint16(anon_sym_RBRACE),
	4090:  uint16(anon_sym_LPAREN),
	4091:  uint16(anon_sym_DQUOTE),
	4092:  uint16(anon_sym_PLUS),
	4093:  uint16(anon_sym_DASH),
	4094:  uint16(anon_sym_0x),
	4095:  uint16(anon_sym_0o),
	4096:  uint16(anon_sym_0b),
	4097:  uint16(anon_sym_BSLASH),
	4098:  uint16(aux_sym__newline_token1),
	4099:  uint16(aux_sym__newline_token2),
	4100:  uint16(aux_sym__newline_token3),
	4101:  uint16(aux_sym__newline_token4),
	4102:  uint16(aux_sym__newline_token5),
	4103:  uint16(aux_sym__newline_token6),
	4104:  uint16(aux_sym__newline_token7),
	4105:  uint16(sym__bom),
	4106:  uint16(sym__unicode_space),
	4107:  uint16(anon_sym_SLASH_SLASH),
	4108:  uint16(2),
	4109:  uint16(236),
	4110:  uint16(5),
	4111:  uint16(sym__normal_bare_identifier),
	4112:  uint16(anon_sym_null),
	4113:  uint16(sym__digit),
	4114:  uint16(anon_sym_true),
	4115:  uint16(anon_sym_false),
	4116:  uint16(234),
	4117:  uint16(24),
	4118:  uint16(sym__eof),
	4119:  uint16(sym_multi_line_comment),
	4120:  uint16(sym__raw_string),
	4121:  uint16(anon_sym_SLASH_DASH),
	4122:  uint16(anon_sym_LBRACE),
	4123:  uint16(anon_sym_SEMI),
	4124:  uint16(anon_sym_LPAREN),
	4125:  uint16(anon_sym_DQUOTE),
	4126:  uint16(anon_sym_PLUS),
	4127:  uint16(anon_sym_DASH),
	4128:  uint16(anon_sym_0x),
	4129:  uint16(anon_sym_0o),
	4130:  uint16(anon_sym_0b),
	4131:  uint16(anon_sym_BSLASH),
	4132:  uint16(aux_sym__newline_token1),
	4133:  uint16(aux_sym__newline_token2),
	4134:  uint16(aux_sym__newline_token3),
	4135:  uint16(aux_sym__newline_token4),
	4136:  uint16(aux_sym__newline_token5),
	4137:  uint16(aux_sym__newline_token6),
	4138:  uint16(aux_sym__newline_token7),
	4139:  uint16(sym__bom),
	4140:  uint16(sym__unicode_space),
	4141:  uint16(anon_sym_SLASH_SLASH),
	4142:  uint16(13),
	4143:  uint16(19),
	4144:  uint16(1),
	4145:  uint16(anon_sym_SLASH_SLASH),
	4146:  uint16(27),
	4147:  uint16(1),
	4148:  uint16(anon_sym_LBRACE),
	4149:  uint16(45),
	4150:  uint16(1),
	4151:  uint16(anon_sym_BSLASH),
	4152:  uint16(238),
	4153:  uint16(1),
	4154:  uint16(anon_sym_SLASH_DASH),
	4155:  uint16(8),
	4156:  uint16(1),
	4157:  uint16(aux_sym_node_repeat1),
	4158:  uint16(34),
	4159:  uint16(1),
	4160:  uint16(sym__escline),
	4161:  uint16(58),
	4162:  uint16(1),
	4163:  uint16(sym__node_space),
	4164:  uint16(85),
	4165:  uint16(1),
	4166:  uint16(sym_node_children),
	4167:  uint16(89),
	4168:  uint16(1),
	4169:  uint16(aux_sym_node_repeat2),
	4170:  uint16(27),
	4171:  uint16(2),
	4172:  uint16(sym__ws),
	4173:  uint16(aux_sym_node_repeat3),
	4174:  uint16(47),
	4175:  uint16(3),
	4176:  uint16(sym_multi_line_comment),
	4177:  uint16(sym__bom),
	4178:  uint16(sym__unicode_space),
	4179:  uint16(141),
	4180:  uint16(3),
	4181:  uint16(sym__node_terminator),
	4182:  uint16(sym__newline),
	4183:  uint16(sym_single_line_comment),
	4184:  uint16(53),
	4185:  uint16(9),
	4186:  uint16(sym__eof),
	4187:  uint16(anon_sym_SEMI),
	4188:  uint16(aux_sym__newline_token1),
	4189:  uint16(aux_sym__newline_token2),
	4190:  uint16(aux_sym__newline_token3),
	4191:  uint16(aux_sym__newline_token4),
	4192:  uint16(aux_sym__newline_token5),
	4193:  uint16(aux_sym__newline_token6),
	4194:  uint16(aux_sym__newline_token7),
	4195:  uint16(13),
	4196:  uint16(19),
	4197:  uint16(1),
	4198:  uint16(anon_sym_SLASH_SLASH),
	4199:  uint16(27),
	4200:  uint16(1),
	4201:  uint16(anon_sym_LBRACE),
	4202:  uint16(45),
	4203:  uint16(1),
	4204:  uint16(anon_sym_BSLASH),
	4205:  uint16(238),
	4206:  uint16(1),
	4207:  uint16(anon_sym_SLASH_DASH),
	4208:  uint16(7),
	4209:  uint16(1),
	4210:  uint16(aux_sym_node_repeat1),
	4211:  uint16(34),
	4212:  uint16(1),
	4213:  uint16(sym__escline),
	4214:  uint16(58),
	4215:  uint16(1),
	4216:  uint16(sym__node_space),
	4217:  uint16(69),
	4218:  uint16(1),
	4219:  uint16(aux_sym_node_repeat2),
	4220:  uint16(97),
	4221:  uint16(1),
	4222:  uint16(sym_node_children),
	4223:  uint16(27),
	4224:  uint16(2),
	4225:  uint16(sym__ws),
	4226:  uint16(aux_sym_node_repeat3),
	4227:  uint16(47),
	4228:  uint16(3),
	4229:  uint16(sym_multi_line_comment),
	4230:  uint16(sym__bom),
	4231:  uint16(sym__unicode_space),
	4232:  uint16(169),
	4233:  uint16(3),
	4234:  uint16(sym__node_terminator),
	4235:  uint16(sym__newline),
	4236:  uint16(sym_single_line_comment),
	4237:  uint16(240),
	4238:  uint16(9),
	4239:  uint16(sym__eof),
	4240:  uint16(anon_sym_SEMI),
	4241:  uint16(aux_sym__newline_token1),
	4242:  uint16(aux_sym__newline_token2),
	4243:  uint16(aux_sym__newline_token3),
	4244:  uint16(aux_sym__newline_token4),
	4245:  uint16(aux_sym__newline_token5),
	4246:  uint16(aux_sym__newline_token6),
	4247:  uint16(aux_sym__newline_token7),
	4248:  uint16(13),
	4249:  uint16(19),
	4250:  uint16(1),
	4251:  uint16(anon_sym_SLASH_SLASH),
	4252:  uint16(27),
	4253:  uint16(1),
	4254:  uint16(anon_sym_LBRACE),
	4255:  uint16(45),
	4256:  uint16(1),
	4257:  uint16(anon_sym_BSLASH),
	4258:  uint16(238),
	4259:  uint16(1),
	4260:  uint16(anon_sym_SLASH_DASH),
	4261:  uint16(3),
	4262:  uint16(1),
	4263:  uint16(aux_sym_node_repeat1),
	4264:  uint16(34),
	4265:  uint16(1),
	4266:  uint16(sym__escline),
	4267:  uint16(58),
	4268:  uint16(1),
	4269:  uint16(sym__node_space),
	4270:  uint16(70),
	4271:  uint16(1),
	4272:  uint16(aux_sym_node_repeat2),
	4273:  uint16(95),
	4274:  uint16(1),
	4275:  uint16(sym_node_children),
	4276:  uint16(27),
	4277:  uint16(2),
	4278:  uint16(sym__ws),
	4279:  uint16(aux_sym_node_repeat3),
	4280:  uint16(47),
	4281:  uint16(3),
	4282:  uint16(sym_multi_line_comment),
	4283:  uint16(sym__bom),
	4284:  uint16(sym__unicode_space),
	4285:  uint16(171),
	4286:  uint16(3),
	4287:  uint16(sym__node_terminator),
	4288:  uint16(sym__newline),
	4289:  uint16(sym_single_line_comment),
	4290:  uint16(51),
	4291:  uint16(9),
	4292:  uint16(sym__eof),
	4293:  uint16(anon_sym_SEMI),
	4294:  uint16(aux_sym__newline_token1),
	4295:  uint16(aux_sym__newline_token2),
	4296:  uint16(aux_sym__newline_token3),
	4297:  uint16(aux_sym__newline_token4),
	4298:  uint16(aux_sym__newline_token5),
	4299:  uint16(aux_sym__newline_token6),
	4300:  uint16(aux_sym__newline_token7),
	4301:  uint16(13),
	4302:  uint16(19),
	4303:  uint16(1),
	4304:  uint16(anon_sym_SLASH_SLASH),
	4305:  uint16(27),
	4306:  uint16(1),
	4307:  uint16(anon_sym_LBRACE),
	4308:  uint16(45),
	4309:  uint16(1),
	4310:  uint16(anon_sym_BSLASH),
	4311:  uint16(238),
	4312:  uint16(1),
	4313:  uint16(anon_sym_SLASH_DASH),
	4314:  uint16(3),
	4315:  uint16(1),
	4316:  uint16(aux_sym_node_repeat1),
	4317:  uint16(34),
	4318:  uint16(1),
	4319:  uint16(sym__escline),
	4320:  uint16(58),
	4321:  uint16(1),
	4322:  uint16(sym__node_space),
	4323:  uint16(89),
	4324:  uint16(1),
	4325:  uint16(aux_sym_node_repeat2),
	4326:  uint16(95),
	4327:  uint16(1),
	4328:  uint16(sym_node_children),
	4329:  uint16(27),
	4330:  uint16(2),
	4331:  uint16(sym__ws),
	4332:  uint16(aux_sym_node_repeat3),
	4333:  uint16(47),
	4334:  uint16(3),
	4335:  uint16(sym_multi_line_comment),
	4336:  uint16(sym__bom),
	4337:  uint16(sym__unicode_space),
	4338:  uint16(171),
	4339:  uint16(3),
	4340:  uint16(sym__node_terminator),
	4341:  uint16(sym__newline),
	4342:  uint16(sym_single_line_comment),
	4343:  uint16(51),
	4344:  uint16(9),
	4345:  uint16(sym__eof),
	4346:  uint16(anon_sym_SEMI),
	4347:  uint16(aux_sym__newline_token1),
	4348:  uint16(aux_sym__newline_token2),
	4349:  uint16(aux_sym__newline_token3),
	4350:  uint16(aux_sym__newline_token4),
	4351:  uint16(aux_sym__newline_token5),
	4352:  uint16(aux_sym__newline_token6),
	4353:  uint16(aux_sym__newline_token7),
	4354:  uint16(13),
	4355:  uint16(19),
	4356:  uint16(1),
	4357:  uint16(anon_sym_SLASH_SLASH),
	4358:  uint16(27),
	4359:  uint16(1),
	4360:  uint16(anon_sym_LBRACE),
	4361:  uint16(45),
	4362:  uint16(1),
	4363:  uint16(anon_sym_BSLASH),
	4364:  uint16(238),
	4365:  uint16(1),
	4366:  uint16(anon_sym_SLASH_DASH),
	4367:  uint16(4),
	4368:  uint16(1),
	4369:  uint16(aux_sym_node_repeat1),
	4370:  uint16(34),
	4371:  uint16(1),
	4372:  uint16(sym__escline),
	4373:  uint16(58),
	4374:  uint16(1),
	4375:  uint16(sym__node_space),
	4376:  uint16(65),
	4377:  uint16(1),
	4378:  uint16(aux_sym_node_repeat2),
	4379:  uint16(109),
	4380:  uint16(1),
	4381:  uint16(sym_node_children),
	4382:  uint16(27),
	4383:  uint16(2),
	4384:  uint16(sym__ws),
	4385:  uint16(aux_sym_node_repeat3),
	4386:  uint16(47),
	4387:  uint16(3),
	4388:  uint16(sym_multi_line_comment),
	4389:  uint16(sym__bom),
	4390:  uint16(sym__unicode_space),
	4391:  uint16(152),
	4392:  uint16(3),
	4393:  uint16(sym__node_terminator),
	4394:  uint16(sym__newline),
	4395:  uint16(sym_single_line_comment),
	4396:  uint16(29),
	4397:  uint16(9),
	4398:  uint16(sym__eof),
	4399:  uint16(anon_sym_SEMI),
	4400:  uint16(aux_sym__newline_token1),
	4401:  uint16(aux_sym__newline_token2),
	4402:  uint16(aux_sym__newline_token3),
	4403:  uint16(aux_sym__newline_token4),
	4404:  uint16(aux_sym__newline_token5),
	4405:  uint16(aux_sym__newline_token6),
	4406:  uint16(aux_sym__newline_token7),
	4407:  uint16(19),
	4408:  uint16(3),
	4409:  uint16(1),
	4410:  uint16(sym_multi_line_comment),
	4411:  uint16(11),
	4412:  uint16(1),
	4413:  uint16(anon_sym_LPAREN),
	4414:  uint16(13),
	4415:  uint16(1),
	4416:  uint16(anon_sym_DQUOTE),
	4417:  uint16(21),
	4418:  uint16(1),
	4419:  uint16(sym__raw_string),
	4420:  uint16(33),
	4421:  uint16(1),
	4422:  uint16(sym__digit),
	4423:  uint16(37),
	4424:  uint16(1),
	4425:  uint16(anon_sym_0x),
	4426:  uint16(39),
	4427:  uint16(1),
	4428:  uint16(anon_sym_0o),
	4429:  uint16(41),
	4430:  uint16(1),
	4431:  uint16(anon_sym_0b),
	4432:  uint16(242),
	4433:  uint16(1),
	4434:  uint16(anon_sym_null),
	4435:  uint16(79),
	4436:  uint16(1),
	4437:  uint16(sym_type),
	4438:  uint16(154),
	4439:  uint16(1),
	4440:  uint16(sym__integer),
	4441:  uint16(194),
	4442:  uint16(1),
	4443:  uint16(sym__escaped_string),
	4444:  uint16(208),
	4445:  uint16(1),
	4446:  uint16(sym_value),
	4447:  uint16(217),
	4448:  uint16(1),
	4449:  uint16(sym_boolean),
	4450:  uint16(271),
	4451:  uint16(1),
	4452:  uint16(sym__sign),
	4453:  uint16(244),
	4454:  uint16(2),
	4455:  uint16(anon_sym_PLUS),
	4456:  uint16(anon_sym_DASH),
	4457:  uint16(246),
	4458:  uint16(2),
	4459:  uint16(anon_sym_true),
	4460:  uint16(anon_sym_false),
	4461:  uint16(199),
	4462:  uint16(3),
	4463:  uint16(sym_keyword),
	4464:  uint16(sym_string),
	4465:  uint16(sym_number),
	4466:  uint16(214),
	4467:  uint16(4),
	4468:  uint16(sym__decimal),
	4469:  uint16(sym__hex),
	4470:  uint16(sym__octal),
	4471:  uint16(sym__binary),
	4472:  uint16(13),
	4473:  uint16(19),
	4474:  uint16(1),
	4475:  uint16(anon_sym_SLASH_SLASH),
	4476:  uint16(27),
	4477:  uint16(1),
	4478:  uint16(anon_sym_LBRACE),
	4479:  uint16(45),
	4480:  uint16(1),
	4481:  uint16(anon_sym_BSLASH),
	4482:  uint16(238),
	4483:  uint16(1),
	4484:  uint16(anon_sym_SLASH_DASH),
	4485:  uint16(5),
	4486:  uint16(1),
	4487:  uint16(aux_sym_node_repeat1),
	4488:  uint16(34),
	4489:  uint16(1),
	4490:  uint16(sym__escline),
	4491:  uint16(58),
	4492:  uint16(1),
	4493:  uint16(sym__node_space),
	4494:  uint16(62),
	4495:  uint16(1),
	4496:  uint16(aux_sym_node_repeat2),
	4497:  uint16(80),
	4498:  uint16(1),
	4499:  uint16(sym_node_children),
	4500:  uint16(27),
	4501:  uint16(2),
	4502:  uint16(sym__ws),
	4503:  uint16(aux_sym_node_repeat3),
	4504:  uint16(47),
	4505:  uint16(3),
	4506:  uint16(sym_multi_line_comment),
	4507:  uint16(sym__bom),
	4508:  uint16(sym__unicode_space),
	4509:  uint16(148),
	4510:  uint16(3),
	4511:  uint16(sym__node_terminator),
	4512:  uint16(sym__newline),
	4513:  uint16(sym_single_line_comment),
	4514:  uint16(57),
	4515:  uint16(9),
	4516:  uint16(sym__eof),
	4517:  uint16(anon_sym_SEMI),
	4518:  uint16(aux_sym__newline_token1),
	4519:  uint16(aux_sym__newline_token2),
	4520:  uint16(aux_sym__newline_token3),
	4521:  uint16(aux_sym__newline_token4),
	4522:  uint16(aux_sym__newline_token5),
	4523:  uint16(aux_sym__newline_token6),
	4524:  uint16(aux_sym__newline_token7),
	4525:  uint16(13),
	4526:  uint16(19),
	4527:  uint16(1),
	4528:  uint16(anon_sym_SLASH_SLASH),
	4529:  uint16(27),
	4530:  uint16(1),
	4531:  uint16(anon_sym_LBRACE),
	4532:  uint16(45),
	4533:  uint16(1),
	4534:  uint16(anon_sym_BSLASH),
	4535:  uint16(238),
	4536:  uint16(1),
	4537:  uint16(anon_sym_SLASH_DASH),
	4538:  uint16(5),
	4539:  uint16(1),
	4540:  uint16(aux_sym_node_repeat1),
	4541:  uint16(34),
	4542:  uint16(1),
	4543:  uint16(sym__escline),
	4544:  uint16(58),
	4545:  uint16(1),
	4546:  uint16(sym__node_space),
	4547:  uint16(80),
	4548:  uint16(1),
	4549:  uint16(sym_node_children),
	4550:  uint16(89),
	4551:  uint16(1),
	4552:  uint16(aux_sym_node_repeat2),
	4553:  uint16(27),
	4554:  uint16(2),
	4555:  uint16(sym__ws),
	4556:  uint16(aux_sym_node_repeat3),
	4557:  uint16(47),
	4558:  uint16(3),
	4559:  uint16(sym_multi_line_comment),
	4560:  uint16(sym__bom),
	4561:  uint16(sym__unicode_space),
	4562:  uint16(148),
	4563:  uint16(3),
	4564:  uint16(sym__node_terminator),
	4565:  uint16(sym__newline),
	4566:  uint16(sym_single_line_comment),
	4567:  uint16(57),
	4568:  uint16(9),
	4569:  uint16(sym__eof),
	4570:  uint16(anon_sym_SEMI),
	4571:  uint16(aux_sym__newline_token1),
	4572:  uint16(aux_sym__newline_token2),
	4573:  uint16(aux_sym__newline_token3),
	4574:  uint16(aux_sym__newline_token4),
	4575:  uint16(aux_sym__newline_token5),
	4576:  uint16(aux_sym__newline_token6),
	4577:  uint16(aux_sym__newline_token7),
	4578:  uint16(13),
	4579:  uint16(19),
	4580:  uint16(1),
	4581:  uint16(anon_sym_SLASH_SLASH),
	4582:  uint16(27),
	4583:  uint16(1),
	4584:  uint16(anon_sym_LBRACE),
	4585:  uint16(45),
	4586:  uint16(1),
	4587:  uint16(anon_sym_BSLASH),
	4588:  uint16(238),
	4589:  uint16(1),
	4590:  uint16(anon_sym_SLASH_DASH),
	4591:  uint16(6),
	4592:  uint16(1),
	4593:  uint16(aux_sym_node_repeat1),
	4594:  uint16(34),
	4595:  uint16(1),
	4596:  uint16(sym__escline),
	4597:  uint16(58),
	4598:  uint16(1),
	4599:  uint16(sym__node_space),
	4600:  uint16(89),
	4601:  uint16(1),
	4602:  uint16(aux_sym_node_repeat2),
	4603:  uint16(118),
	4604:  uint16(1),
	4605:  uint16(sym_node_children),
	4606:  uint16(27),
	4607:  uint16(2),
	4608:  uint16(sym__ws),
	4609:  uint16(aux_sym_node_repeat3),
	4610:  uint16(47),
	4611:  uint16(3),
	4612:  uint16(sym_multi_line_comment),
	4613:  uint16(sym__bom),
	4614:  uint16(sym__unicode_space),
	4615:  uint16(142),
	4616:  uint16(3),
	4617:  uint16(sym__node_terminator),
	4618:  uint16(sym__newline),
	4619:  uint16(sym_single_line_comment),
	4620:  uint16(49),
	4621:  uint16(9),
	4622:  uint16(sym__eof),
	4623:  uint16(anon_sym_SEMI),
	4624:  uint16(aux_sym__newline_token1),
	4625:  uint16(aux_sym__newline_token2),
	4626:  uint16(aux_sym__newline_token3),
	4627:  uint16(aux_sym__newline_token4),
	4628:  uint16(aux_sym__newline_token5),
	4629:  uint16(aux_sym__newline_token6),
	4630:  uint16(aux_sym__newline_token7),
	4631:  uint16(13),
	4632:  uint16(19),
	4633:  uint16(1),
	4634:  uint16(anon_sym_SLASH_SLASH),
	4635:  uint16(27),
	4636:  uint16(1),
	4637:  uint16(anon_sym_LBRACE),
	4638:  uint16(45),
	4639:  uint16(1),
	4640:  uint16(anon_sym_BSLASH),
	4641:  uint16(238),
	4642:  uint16(1),
	4643:  uint16(anon_sym_SLASH_DASH),
	4644:  uint16(2),
	4645:  uint16(1),
	4646:  uint16(aux_sym_node_repeat1),
	4647:  uint16(34),
	4648:  uint16(1),
	4649:  uint16(sym__escline),
	4650:  uint16(58),
	4651:  uint16(1),
	4652:  uint16(sym__node_space),
	4653:  uint16(72),
	4654:  uint16(1),
	4655:  uint16(aux_sym_node_repeat2),
	4656:  uint16(104),
	4657:  uint16(1),
	4658:  uint16(sym_node_children),
	4659:  uint16(27),
	4660:  uint16(2),
	4661:  uint16(sym__ws),
	4662:  uint16(aux_sym_node_repeat3),
	4663:  uint16(47),
	4664:  uint16(3),
	4665:  uint16(sym_multi_line_comment),
	4666:  uint16(sym__bom),
	4667:  uint16(sym__unicode_space),
	4668:  uint16(156),
	4669:  uint16(3),
	4670:  uint16(sym__node_terminator),
	4671:  uint16(sym__newline),
	4672:  uint16(sym_single_line_comment),
	4673:  uint16(248),
	4674:  uint16(9),
	4675:  uint16(sym__eof),
	4676:  uint16(anon_sym_SEMI),
	4677:  uint16(aux_sym__newline_token1),
	4678:  uint16(aux_sym__newline_token2),
	4679:  uint16(aux_sym__newline_token3),
	4680:  uint16(aux_sym__newline_token4),
	4681:  uint16(aux_sym__newline_token5),
	4682:  uint16(aux_sym__newline_token6),
	4683:  uint16(aux_sym__newline_token7),
	4684:  uint16(13),
	4685:  uint16(19),
	4686:  uint16(1),
	4687:  uint16(anon_sym_SLASH_SLASH),
	4688:  uint16(27),
	4689:  uint16(1),
	4690:  uint16(anon_sym_LBRACE),
	4691:  uint16(45),
	4692:  uint16(1),
	4693:  uint16(anon_sym_BSLASH),
	4694:  uint16(238),
	4695:  uint16(1),
	4696:  uint16(anon_sym_SLASH_DASH),
	4697:  uint16(4),
	4698:  uint16(1),
	4699:  uint16(aux_sym_node_repeat1),
	4700:  uint16(34),
	4701:  uint16(1),
	4702:  uint16(sym__escline),
	4703:  uint16(58),
	4704:  uint16(1),
	4705:  uint16(sym__node_space),
	4706:  uint16(89),
	4707:  uint16(1),
	4708:  uint16(aux_sym_node_repeat2),
	4709:  uint16(109),
	4710:  uint16(1),
	4711:  uint16(sym_node_children),
	4712:  uint16(27),
	4713:  uint16(2),
	4714:  uint16(sym__ws),
	4715:  uint16(aux_sym_node_repeat3),
	4716:  uint16(47),
	4717:  uint16(3),
	4718:  uint16(sym_multi_line_comment),
	4719:  uint16(sym__bom),
	4720:  uint16(sym__unicode_space),
	4721:  uint16(152),
	4722:  uint16(3),
	4723:  uint16(sym__node_terminator),
	4724:  uint16(sym__newline),
	4725:  uint16(sym_single_line_comment),
	4726:  uint16(29),
	4727:  uint16(9),
	4728:  uint16(sym__eof),
	4729:  uint16(anon_sym_SEMI),
	4730:  uint16(aux_sym__newline_token1),
	4731:  uint16(aux_sym__newline_token2),
	4732:  uint16(aux_sym__newline_token3),
	4733:  uint16(aux_sym__newline_token4),
	4734:  uint16(aux_sym__newline_token5),
	4735:  uint16(aux_sym__newline_token6),
	4736:  uint16(aux_sym__newline_token7),
	4737:  uint16(4),
	4738:  uint16(255),
	4739:  uint16(1),
	4740:  uint16(anon_sym_SLASH_SLASH),
	4741:  uint16(73),
	4742:  uint16(5),
	4743:  uint16(sym__linespace),
	4744:  uint16(sym__newline),
	4745:  uint16(sym__ws),
	4746:  uint16(sym_single_line_comment),
	4747:  uint16(aux_sym_document_repeat1),
	4748:  uint16(250),
	4749:  uint16(9),
	4750:  uint16(sym__raw_string),
	4752:  uint16(anon_sym_SLASH_DASH),
	4753:  uint16(anon_sym_RBRACE),
	4754:  uint16(sym__normal_bare_identifier),
	4755:  uint16(anon_sym_LPAREN),
	4756:  uint16(anon_sym_DQUOTE),
	4757:  uint16(anon_sym_PLUS),
	4758:  uint16(anon_sym_DASH),
	4759:  uint16(252),
	4760:  uint16(10),
	4761:  uint16(sym_multi_line_comment),
	4762:  uint16(aux_sym__newline_token1),
	4763:  uint16(aux_sym__newline_token2),
	4764:  uint16(aux_sym__newline_token3),
	4765:  uint16(aux_sym__newline_token4),
	4766:  uint16(aux_sym__newline_token5),
	4767:  uint16(aux_sym__newline_token6),
	4768:  uint16(aux_sym__newline_token7),
	4769:  uint16(sym__bom),
	4770:  uint16(sym__unicode_space),
	4771:  uint16(8),
	4772:  uint16(258),
	4773:  uint16(1),
	4774:  uint16(anon_sym_BSLASH),
	4775:  uint16(74),
	4776:  uint16(1),
	4777:  uint16(aux_sym_node_repeat1),
	4778:  uint16(125),
	4779:  uint16(1),
	4780:  uint16(sym__escline),
	4781:  uint16(186),
	4782:  uint16(1),
	4783:  uint16(sym__node_space),
	4784:  uint16(100),
	4785:  uint16(2),
	4786:  uint16(sym__ws),
	4787:  uint16(aux_sym_node_repeat3),
	4788:  uint16(261),
	4789:  uint16(3),
	4790:  uint16(sym_multi_line_comment),
	4791:  uint16(sym__bom),
	4792:  uint16(sym__unicode_space),
	4793:  uint16(79),
	4794:  uint16(5),
	4795:  uint16(sym__normal_bare_identifier),
	4796:  uint16(anon_sym_null),
	4797:  uint16(sym__digit),
	4798:  uint16(anon_sym_true),
	4799:  uint16(anon_sym_false),
	4800:  uint16(81),
	4801:  uint16(10),
	4802:  uint16(sym__raw_string),
	4803:  uint16(anon_sym_SLASH_DASH),
	4804:  uint16(anon_sym_LBRACE),
	4805:  uint16(anon_sym_LPAREN),
	4806:  uint16(anon_sym_DQUOTE),
	4807:  uint16(anon_sym_PLUS),
	4808:  uint16(anon_sym_DASH),
	4809:  uint16(anon_sym_0x),
	4810:  uint16(anon_sym_0o),
	4811:  uint16(anon_sym_0b),
	4812:  uint16(5),
	4813:  uint16(255),
	4814:  uint16(1),
	4815:  uint16(anon_sym_SLASH_SLASH),
	4816:  uint16(264),
	4817:  uint16(1),
	4818:  uint16(anon_sym_RBRACE),
	4819:  uint16(73),
	4820:  uint16(5),
	4821:  uint16(sym__linespace),
	4822:  uint16(sym__newline),
	4823:  uint16(sym__ws),
	4824:  uint16(sym_single_line_comment),
	4825:  uint16(aux_sym_document_repeat1),
	4826:  uint16(250),
	4827:  uint16(7),
	4828:  uint16(sym__raw_string),
	4829:  uint16(anon_sym_SLASH_DASH),
	4830:  uint16(sym__normal_bare_identifier),
	4831:  uint16(anon_sym_LPAREN),
	4832:  uint16(anon_sym_DQUOTE),
	4833:  uint16(anon_sym_PLUS),
	4834:  uint16(anon_sym_DASH),
	4835:  uint16(252),
	4836:  uint16(10),
	4837:  uint16(sym_multi_line_comment),
	4838:  uint16(aux_sym__newline_token1),
	4839:  uint16(aux_sym__newline_token2),
	4840:  uint16(aux_sym__newline_token3),
	4841:  uint16(aux_sym__newline_token4),
	4842:  uint16(aux_sym__newline_token5),
	4843:  uint16(aux_sym__newline_token6),
	4844:  uint16(aux_sym__newline_token7),
	4845:  uint16(sym__bom),
	4846:  uint16(sym__unicode_space),
	4847:  uint16(5),
	4848:  uint16(255),
	4849:  uint16(1),
	4850:  uint16(anon_sym_SLASH_SLASH),
	4851:  uint16(267),
	4852:  uint16(1),
	4853:  uint16(anon_sym_RBRACE),
	4854:  uint16(73),
	4855:  uint16(5),
	4856:  uint16(sym__linespace),
	4857:  uint16(sym__newline),
	4858:  uint16(sym__ws),
	4859:  uint16(sym_single_line_comment),
	4860:  uint16(aux_sym_document_repeat1),
	4861:  uint16(250),
	4862:  uint16(7),
	4863:  uint16(sym__raw_string),
	4864:  uint16(anon_sym_SLASH_DASH),
	4865:  uint16(sym__normal_bare_identifier),
	4866:  uint16(anon_sym_LPAREN),
	4867:  uint16(anon_sym_DQUOTE),
	4868:  uint16(anon_sym_PLUS),
	4869:  uint16(anon_sym_DASH),
	4870:  uint16(252),
	4871:  uint16(10),
	4872:  uint16(sym_multi_line_comment),
	4873:  uint16(aux_sym__newline_token1),
	4874:  uint16(aux_sym__newline_token2),
	4875:  uint16(aux_sym__newline_token3),
	4876:  uint16(aux_sym__newline_token4),
	4877:  uint16(aux_sym__newline_token5),
	4878:  uint16(aux_sym__newline_token6),
	4879:  uint16(aux_sym__newline_token7),
	4880:  uint16(sym__bom),
	4881:  uint16(sym__unicode_space),
	4882:  uint16(5),
	4883:  uint16(255),
	4884:  uint16(1),
	4885:  uint16(anon_sym_SLASH_SLASH),
	4886:  uint16(270),
	4887:  uint16(1),
	4889:  uint16(73),
	4890:  uint16(5),
	4891:  uint16(sym__linespace),
	4892:  uint16(sym__newline),
	4893:  uint16(sym__ws),
	4894:  uint16(sym_single_line_comment),
	4895:  uint16(aux_sym_document_repeat1),
	4896:  uint16(250),
	4897:  uint16(7),
	4898:  uint16(sym__raw_string),
	4899:  uint16(anon_sym_SLASH_DASH),
	4900:  uint16(sym__normal_bare_identifier),
	4901:  uint16(anon_sym_LPAREN),
	4902:  uint16(anon_sym_DQUOTE),
	4903:  uint16(anon_sym_PLUS),
	4904:  uint16(anon_sym_DASH),
	4905:  uint16(252),
	4906:  uint16(10),
	4907:  uint16(sym_multi_line_comment),
	4908:  uint16(aux_sym__newline_token1),
	4909:  uint16(aux_sym__newline_token2),
	4910:  uint16(aux_sym__newline_token3),
	4911:  uint16(aux_sym__newline_token4),
	4912:  uint16(aux_sym__newline_token5),
	4913:  uint16(aux_sym__newline_token6),
	4914:  uint16(aux_sym__newline_token7),
	4915:  uint16(sym__bom),
	4916:  uint16(sym__unicode_space),
	4917:  uint16(5),
	4918:  uint16(255),
	4919:  uint16(1),
	4920:  uint16(anon_sym_SLASH_SLASH),
	4921:  uint16(273),
	4922:  uint16(1),
	4923:  uint16(anon_sym_RBRACE),
	4924:  uint16(73),
	4925:  uint16(5),
	4926:  uint16(sym__linespace),
	4927:  uint16(sym__newline),
	4928:  uint16(sym__ws),
	4929:  uint16(sym_single_line_comment),
	4930:  uint16(aux_sym_document_repeat1),
	4931:  uint16(250),
	4932:  uint16(7),
	4933:  uint16(sym__raw_string),
	4934:  uint16(anon_sym_SLASH_DASH),
	4935:  uint16(sym__normal_bare_identifier),
	4936:  uint16(anon_sym_LPAREN),
	4937:  uint16(anon_sym_DQUOTE),
	4938:  uint16(anon_sym_PLUS),
	4939:  uint16(anon_sym_DASH),
	4940:  uint16(252),
	4941:  uint16(10),
	4942:  uint16(sym_multi_line_comment),
	4943:  uint16(aux_sym__newline_token1),
	4944:  uint16(aux_sym__newline_token2),
	4945:  uint16(aux_sym__newline_token3),
	4946:  uint16(aux_sym__newline_token4),
	4947:  uint16(aux_sym__newline_token5),
	4948:  uint16(aux_sym__newline_token6),
	4949:  uint16(aux_sym__newline_token7),
	4950:  uint16(sym__bom),
	4951:  uint16(sym__unicode_space),
	4952:  uint16(16),
	4953:  uint16(3),
	4954:  uint16(1),
	4955:  uint16(sym_multi_line_comment),
	4956:  uint16(13),
	4957:  uint16(1),
	4958:  uint16(anon_sym_DQUOTE),
	4959:  uint16(21),
	4960:  uint16(1),
	4961:  uint16(sym__raw_string),
	4962:  uint16(33),
	4963:  uint16(1),
	4964:  uint16(sym__digit),
	4965:  uint16(37),
	4966:  uint16(1),
	4967:  uint16(anon_sym_0x),
	4968:  uint16(39),
	4969:  uint16(1),
	4970:  uint16(anon_sym_0o),
	4971:  uint16(41),
	4972:  uint16(1),
	4973:  uint16(anon_sym_0b),
	4974:  uint16(242),
	4975:  uint16(1),
	4976:  uint16(anon_sym_null),
	4977:  uint16(154),
	4978:  uint16(1),
	4979:  uint16(sym__integer),
	4980:  uint16(194),
	4981:  uint16(1),
	4982:  uint16(sym__escaped_string),
	4983:  uint16(217),
	4984:  uint16(1),
	4985:  uint16(sym_boolean),
	4986:  uint16(271),
	4987:  uint16(1),
	4988:  uint16(sym__sign),
	4989:  uint16(244),
	4990:  uint16(2),
	4991:  uint16(anon_sym_PLUS),
	4992:  uint16(anon_sym_DASH),
	4993:  uint16(246),
	4994:  uint16(2),
	4995:  uint16(anon_sym_true),
	4996:  uint16(anon_sym_false),
	4997:  uint16(220),
	4998:  uint16(3),
	4999:  uint16(sym_keyword),
	5000:  uint16(sym_string),
	5001:  uint16(sym_number),
	5002:  uint16(214),
	5003:  uint16(4),
	5004:  uint16(sym__decimal),
	5005:  uint16(sym__hex),
	5006:  uint16(sym__octal),
	5007:  uint16(sym__binary),
	5008:  uint16(9),
	5009:  uint16(19),
	5010:  uint16(1),
	5011:  uint16(anon_sym_SLASH_SLASH),
	5012:  uint16(278),
	5013:  uint16(1),
	5014:  uint16(anon_sym_BSLASH),
	5015:  uint16(87),
	5016:  uint16(1),
	5017:  uint16(aux_sym_node_repeat1),
	5018:  uint16(212),
	5019:  uint16(1),
	5020:  uint16(sym__escline),
	5021:  uint16(234),
	5022:  uint16(1),
	5023:  uint16(sym__node_space),
	5024:  uint16(86),
	5025:  uint16(2),
	5026:  uint16(sym__ws),
	5027:  uint16(aux_sym_node_repeat3),
	5028:  uint16(280),
	5029:  uint16(3),
	5030:  uint16(sym_multi_line_comment),
	5031:  uint16(sym__bom),
	5032:  uint16(sym__unicode_space),
	5033:  uint16(155),
	5034:  uint16(3),
	5035:  uint16(sym__node_terminator),
	5036:  uint16(sym__newline),
	5037:  uint16(sym_single_line_comment),
	5038:  uint16(276),
	5039:  uint16(9),
	5040:  uint16(sym__eof),
	5041:  uint16(anon_sym_SEMI),
	5042:  uint16(aux_sym__newline_token1),
	5043:  uint16(aux_sym__newline_token2),
	5044:  uint16(aux_sym__newline_token3),
	5045:  uint16(aux_sym__newline_token4),
	5046:  uint16(aux_sym__newline_token5),
	5047:  uint16(aux_sym__newline_token6),
	5048:  uint16(aux_sym__newline_token7),
	5049:  uint16(3),
	5050:  uint16(119),
	5051:  uint16(1),
	5052:  uint16(aux_sym__integer_repeat1),
	5053:  uint16(284),
	5054:  uint16(2),
	5055:  uint16(anon_sym__),
	5056:  uint16(sym__digit),
	5057:  uint16(282),
	5058:  uint16(19),
	5059:  uint16(sym__eof),
	5060:  uint16(sym_multi_line_comment),
	5061:  uint16(anon_sym_SLASH_DASH),
	5062:  uint16(anon_sym_LBRACE),
	5063:  uint16(anon_sym_SEMI),
	5064:  uint16(anon_sym_DOT),
	5065:  uint16(anon_sym_e),
	5066:  uint16(anon_sym_E),
	5067:  uint16(anon_sym_BSLASH),
	5068:  uint16(aux_sym__newline_token1),
	5069:  uint16(aux_sym__newline_token2),
	5070:  uint16(aux_sym__newline_token3),
	5071:  uint16(aux_sym__newline_token4),
	5072:  uint16(aux_sym__newline_token5),
	5073:  uint16(aux_sym__newline_token6),
	5074:  uint16(aux_sym__newline_token7),
	5075:  uint16(sym__bom),
	5076:  uint16(sym__unicode_space),
	5077:  uint16(anon_sym_SLASH_SLASH),
	5078:  uint16(9),
	5079:  uint16(19),
	5080:  uint16(1),
	5081:  uint16(anon_sym_SLASH_SLASH),
	5082:  uint16(278),
	5083:  uint16(1),
	5084:  uint16(anon_sym_BSLASH),
	5085:  uint16(187),
	5086:  uint16(1),
	5087:  uint16(aux_sym_node_repeat1),
	5088:  uint16(212),
	5089:  uint16(1),
	5090:  uint16(sym__escline),
	5091:  uint16(234),
	5092:  uint16(1),
	5093:  uint16(sym__node_space),
	5094:  uint16(193),
	5095:  uint16(2),
	5096:  uint16(sym__ws),
	5097:  uint16(aux_sym_node_repeat3),
	5098:  uint16(288),
	5099:  uint16(3),
	5100:  uint16(sym_multi_line_comment),
	5101:  uint16(sym__bom),
	5102:  uint16(sym__unicode_space),
	5103:  uint16(129),
	5104:  uint16(3),
	5105:  uint16(sym__node_terminator),
	5106:  uint16(sym__newline),
	5107:  uint16(sym_single_line_comment),
	5108:  uint16(286),
	5109:  uint16(9),
	5110:  uint16(sym__eof),
	5111:  uint16(anon_sym_SEMI),
	5112:  uint16(aux_sym__newline_token1),
	5113:  uint16(aux_sym__newline_token2),
	5114:  uint16(aux_sym__newline_token3),
	5115:  uint16(aux_sym__newline_token4),
	5116:  uint16(aux_sym__newline_token5),
	5117:  uint16(aux_sym__newline_token6),
	5118:  uint16(aux_sym__newline_token7),
	5119:  uint16(9),
	5120:  uint16(19),
	5121:  uint16(1),
	5122:  uint16(anon_sym_SLASH_SLASH),
	5123:  uint16(278),
	5124:  uint16(1),
	5125:  uint16(anon_sym_BSLASH),
	5126:  uint16(187),
	5127:  uint16(1),
	5128:  uint16(aux_sym_node_repeat1),
	5129:  uint16(212),
	5130:  uint16(1),
	5131:  uint16(sym__escline),
	5132:  uint16(234),
	5133:  uint16(1),
	5134:  uint16(sym__node_space),
	5135:  uint16(193),
	5136:  uint16(2),
	5137:  uint16(sym__ws),
	5138:  uint16(aux_sym_node_repeat3),
	5139:  uint16(288),
	5140:  uint16(3),
	5141:  uint16(sym_multi_line_comment),
	5142:  uint16(sym__bom),
	5143:  uint16(sym__unicode_space),
	5144:  uint16(136),
	5145:  uint16(3),
	5146:  uint16(sym__node_terminator),
	5147:  uint16(sym__newline),
	5148:  uint16(sym_single_line_comment),
	5149:  uint16(290),
	5150:  uint16(9),
	5151:  uint16(sym__eof),
	5152:  uint16(anon_sym_SEMI),
	5153:  uint16(aux_sym__newline_token1),
	5154:  uint16(aux_sym__newline_token2),
	5155:  uint16(aux_sym__newline_token3),
	5156:  uint16(aux_sym__newline_token4),
	5157:  uint16(aux_sym__newline_token5),
	5158:  uint16(aux_sym__newline_token6),
	5159:  uint16(aux_sym__newline_token7),
	5160:  uint16(9),
	5161:  uint16(19),
	5162:  uint16(1),
	5163:  uint16(anon_sym_SLASH_SLASH),
	5164:  uint16(278),
	5165:  uint16(1),
	5166:  uint16(anon_sym_BSLASH),
	5167:  uint16(187),
	5168:  uint16(1),
	5169:  uint16(aux_sym_node_repeat1),
	5170:  uint16(212),
	5171:  uint16(1),
	5172:  uint16(sym__escline),
	5173:  uint16(234),
	5174:  uint16(1),
	5175:  uint16(sym__node_space),
	5176:  uint16(193),
	5177:  uint16(2),
	5178:  uint16(sym__ws),
	5179:  uint16(aux_sym_node_repeat3),
	5180:  uint16(288),
	5181:  uint16(3),
	5182:  uint16(sym_multi_line_comment),
	5183:  uint16(sym__bom),
	5184:  uint16(sym__unicode_space),
	5185:  uint16(128),
	5186:  uint16(3),
	5187:  uint16(sym__node_terminator),
	5188:  uint16(sym__newline),
	5189:  uint16(sym_single_line_comment),
	5190:  uint16(292),
	5191:  uint16(9),
	5192:  uint16(sym__eof),
	5193:  uint16(anon_sym_SEMI),
	5194:  uint16(aux_sym__newline_token1),
	5195:  uint16(aux_sym__newline_token2),
	5196:  uint16(aux_sym__newline_token3),
	5197:  uint16(aux_sym__newline_token4),
	5198:  uint16(aux_sym__newline_token5),
	5199:  uint16(aux_sym__newline_token6),
	5200:  uint16(aux_sym__newline_token7),
	5201:  uint16(9),
	5202:  uint16(19),
	5203:  uint16(1),
	5204:  uint16(anon_sym_SLASH_SLASH),
	5205:  uint16(278),
	5206:  uint16(1),
	5207:  uint16(anon_sym_BSLASH),
	5208:  uint16(117),
	5209:  uint16(1),
	5210:  uint16(aux_sym_node_repeat1),
	5211:  uint16(212),
	5212:  uint16(1),
	5213:  uint16(sym__escline),
	5214:  uint16(234),
	5215:  uint16(1),
	5216:  uint16(sym__node_space),
	5217:  uint16(114),
	5218:  uint16(2),
	5219:  uint16(sym__ws),
	5220:  uint16(aux_sym_node_repeat3),
	5221:  uint16(296),
	5222:  uint16(3),
	5223:  uint16(sym_multi_line_comment),
	5224:  uint16(sym__bom),
	5225:  uint16(sym__unicode_space),
	5226:  uint16(131),
	5227:  uint16(3),
	5228:  uint16(sym__node_terminator),
	5229:  uint16(sym__newline),
	5230:  uint16(sym_single_line_comment),
	5231:  uint16(294),
	5232:  uint16(9),
	5233:  uint16(sym__eof),
	5234:  uint16(anon_sym_SEMI),
	5235:  uint16(aux_sym__newline_token1),
	5236:  uint16(aux_sym__newline_token2),
	5237:  uint16(aux_sym__newline_token3),
	5238:  uint16(aux_sym__newline_token4),
	5239:  uint16(aux_sym__newline_token5),
	5240:  uint16(aux_sym__newline_token6),
	5241:  uint16(aux_sym__newline_token7),
	5242:  uint16(9),
	5243:  uint16(19),
	5244:  uint16(1),
	5245:  uint16(anon_sym_SLASH_SLASH),
	5246:  uint16(278),
	5247:  uint16(1),
	5248:  uint16(anon_sym_BSLASH),
	5249:  uint16(83),
	5250:  uint16(1),
	5251:  uint16(aux_sym_node_repeat1),
	5252:  uint16(222),
	5253:  uint16(1),
	5254:  uint16(sym__escline),
	5255:  uint16(234),
	5256:  uint16(1),
	5257:  uint16(sym__node_space),
	5258:  uint16(195),
	5259:  uint16(2),
	5260:  uint16(sym__ws),
	5261:  uint16(aux_sym_node_repeat3),
	5262:  uint16(300),
	5263:  uint16(3),
	5264:  uint16(sym_multi_line_comment),
	5265:  uint16(sym__bom),
	5266:  uint16(sym__unicode_space),
	5267:  uint16(132),
	5268:  uint16(3),
	5269:  uint16(sym__node_terminator),
	5270:  uint16(sym__newline),
	5271:  uint16(sym_single_line_comment),
	5272:  uint16(298),
	5273:  uint16(9),
	5274:  uint16(sym__eof),
	5275:  uint16(anon_sym_SEMI),
	5276:  uint16(aux_sym__newline_token1),
	5277:  uint16(aux_sym__newline_token2),
	5278:  uint16(aux_sym__newline_token3),
	5279:  uint16(aux_sym__newline_token4),
	5280:  uint16(aux_sym__newline_token5),
	5281:  uint16(aux_sym__newline_token6),
	5282:  uint16(aux_sym__newline_token7),
	5283:  uint16(9),
	5284:  uint16(19),
	5285:  uint16(1),
	5286:  uint16(anon_sym_SLASH_SLASH),
	5287:  uint16(278),
	5288:  uint16(1),
	5289:  uint16(anon_sym_BSLASH),
	5290:  uint16(187),
	5291:  uint16(1),
	5292:  uint16(aux_sym_node_repeat1),
	5293:  uint16(212),
	5294:  uint16(1),
	5295:  uint16(sym__escline),
	5296:  uint16(234),
	5297:  uint16(1),
	5298:  uint16(sym__node_space),
	5299:  uint16(193),
	5300:  uint16(2),
	5301:  uint16(sym__ws),
	5302:  uint16(aux_sym_node_repeat3),
	5303:  uint16(288),
	5304:  uint16(3),
	5305:  uint16(sym_multi_line_comment),
	5306:  uint16(sym__bom),
	5307:  uint16(sym__unicode_space),
	5308:  uint16(132),
	5309:  uint16(3),
	5310:  uint16(sym__node_terminator),
	5311:  uint16(sym__newline),
	5312:  uint16(sym_single_line_comment),
	5313:  uint16(298),
	5314:  uint16(9),
	5315:  uint16(sym__eof),
	5316:  uint16(anon_sym_SEMI),
	5317:  uint16(aux_sym__newline_token1),
	5318:  uint16(aux_sym__newline_token2),
	5319:  uint16(aux_sym__newline_token3),
	5320:  uint16(aux_sym__newline_token4),
	5321:  uint16(aux_sym__newline_token5),
	5322:  uint16(aux_sym__newline_token6),
	5323:  uint16(aux_sym__newline_token7),
	5324:  uint16(9),
	5325:  uint16(19),
	5326:  uint16(1),
	5327:  uint16(anon_sym_SLASH_SLASH),
	5328:  uint16(278),
	5329:  uint16(1),
	5330:  uint16(anon_sym_BSLASH),
	5331:  uint16(187),
	5332:  uint16(1),
	5333:  uint16(aux_sym_node_repeat1),
	5334:  uint16(212),
	5335:  uint16(1),
	5336:  uint16(sym__escline),
	5337:  uint16(234),
	5338:  uint16(1),
	5339:  uint16(sym__node_space),
	5340:  uint16(193),
	5341:  uint16(2),
	5342:  uint16(sym__ws),
	5343:  uint16(aux_sym_node_repeat3),
	5344:  uint16(288),
	5345:  uint16(3),
	5346:  uint16(sym_multi_line_comment),
	5347:  uint16(sym__bom),
	5348:  uint16(sym__unicode_space),
	5349:  uint16(133),
	5350:  uint16(3),
	5351:  uint16(sym__node_terminator),
	5352:  uint16(sym__newline),
	5353:  uint16(sym_single_line_comment),
	5354:  uint16(302),
	5355:  uint16(9),
	5356:  uint16(sym__eof),
	5357:  uint16(anon_sym_SEMI),
	5358:  uint16(aux_sym__newline_token1),
	5359:  uint16(aux_sym__newline_token2),
	5360:  uint16(aux_sym__newline_token3),
	5361:  uint16(aux_sym__newline_token4),
	5362:  uint16(aux_sym__newline_token5),
	5363:  uint16(aux_sym__newline_token6),
	5364:  uint16(aux_sym__newline_token7),
	5365:  uint16(8),
	5366:  uint16(306),
	5367:  uint16(1),
	5368:  uint16(anon_sym_BSLASH),
	5369:  uint16(10),
	5370:  uint16(1),
	5371:  uint16(aux_sym_node_repeat1),
	5372:  uint16(89),
	5373:  uint16(1),
	5374:  uint16(aux_sym_node_repeat2),
	5375:  uint16(125),
	5376:  uint16(1),
	5377:  uint16(sym__escline),
	5378:  uint16(186),
	5379:  uint16(1),
	5380:  uint16(sym__node_space),
	5381:  uint16(100),
	5382:  uint16(2),
	5383:  uint16(sym__ws),
	5384:  uint16(aux_sym_node_repeat3),
	5385:  uint16(309),
	5386:  uint16(3),
	5387:  uint16(sym_multi_line_comment),
	5388:  uint16(sym__bom),
	5389:  uint16(sym__unicode_space),
	5390:  uint16(304),
	5391:  uint16(12),
	5392:  uint16(sym__eof),
	5393:  uint16(anon_sym_SLASH_DASH),
	5394:  uint16(anon_sym_LBRACE),
	5395:  uint16(anon_sym_SEMI),
	5396:  uint16(aux_sym__newline_token1),
	5397:  uint16(aux_sym__newline_token2),
	5398:  uint16(aux_sym__newline_token3),
	5399:  uint16(aux_sym__newline_token4),
	5400:  uint16(aux_sym__newline_token5),
	5401:  uint16(aux_sym__newline_token6),
	5402:  uint16(aux_sym__newline_token7),
	5403:  uint16(anon_sym_SLASH_SLASH),
	5404:  uint16(9),
	5405:  uint16(19),
	5406:  uint16(1),
	5407:  uint16(anon_sym_SLASH_SLASH),
	5408:  uint16(278),
	5409:  uint16(1),
	5410:  uint16(anon_sym_BSLASH),
	5411:  uint16(111),
	5412:  uint16(1),
	5413:  uint16(aux_sym_node_repeat1),
	5414:  uint16(212),
	5415:  uint16(1),
	5416:  uint16(sym__escline),
	5417:  uint16(234),
	5418:  uint16(1),
	5419:  uint16(sym__node_space),
	5420:  uint16(103),
	5421:  uint16(2),
	5422:  uint16(sym__ws),
	5423:  uint16(aux_sym_node_repeat3),
	5424:  uint16(314),
	5425:  uint16(3),
	5426:  uint16(sym_multi_line_comment),
	5427:  uint16(sym__bom),
	5428:  uint16(sym__unicode_space),
	5429:  uint16(134),
	5430:  uint16(3),
	5431:  uint16(sym__node_terminator),
	5432:  uint16(sym__newline),
	5433:  uint16(sym_single_line_comment),
	5434:  uint16(312),
	5435:  uint16(9),
	5436:  uint16(sym__eof),
	5437:  uint16(anon_sym_SEMI),
	5438:  uint16(aux_sym__newline_token1),
	5439:  uint16(aux_sym__newline_token2),
	5440:  uint16(aux_sym__newline_token3),
	5441:  uint16(aux_sym__newline_token4),
	5442:  uint16(aux_sym__newline_token5),
	5443:  uint16(aux_sym__newline_token6),
	5444:  uint16(aux_sym__newline_token7),
	5445:  uint16(9),
	5446:  uint16(19),
	5447:  uint16(1),
	5448:  uint16(anon_sym_SLASH_SLASH),
	5449:  uint16(278),
	5450:  uint16(1),
	5451:  uint16(anon_sym_BSLASH),
	5452:  uint16(98),
	5453:  uint16(1),
	5454:  uint16(aux_sym_node_repeat1),
	5455:  uint16(222),
	5456:  uint16(1),
	5457:  uint16(sym__escline),
	5458:  uint16(234),
	5459:  uint16(1),
	5460:  uint16(sym__node_space),
	5461:  uint16(195),
	5462:  uint16(2),
	5463:  uint16(sym__ws),
	5464:  uint16(aux_sym_node_repeat3),
	5465:  uint16(300),
	5466:  uint16(3),
	5467:  uint16(sym_multi_line_comment),
	5468:  uint16(sym__bom),
	5469:  uint16(sym__unicode_space),
	5470:  uint16(151),
	5471:  uint16(3),
	5472:  uint16(sym__node_terminator),
	5473:  uint16(sym__newline),
	5474:  uint16(sym_single_line_comment),
	5475:  uint16(316),
	5476:  uint16(9),
	5477:  uint16(sym__eof),
	5478:  uint16(anon_sym_SEMI),
	5479:  uint16(aux_sym__newline_token1),
	5480:  uint16(aux_sym__newline_token2),
	5481:  uint16(aux_sym__newline_token3),
	5482:  uint16(aux_sym__newline_token4),
	5483:  uint16(aux_sym__newline_token5),
	5484:  uint16(aux_sym__newline_token6),
	5485:  uint16(aux_sym__newline_token7),
	5486:  uint16(9),
	5487:  uint16(19),
	5488:  uint16(1),
	5489:  uint16(anon_sym_SLASH_SLASH),
	5490:  uint16(278),
	5491:  uint16(1),
	5492:  uint16(anon_sym_BSLASH),
	5493:  uint16(187),
	5494:  uint16(1),
	5495:  uint16(aux_sym_node_repeat1),
	5496:  uint16(212),
	5497:  uint16(1),
	5498:  uint16(sym__escline),
	5499:  uint16(234),
	5500:  uint16(1),
	5501:  uint16(sym__node_space),
	5502:  uint16(193),
	5503:  uint16(2),
	5504:  uint16(sym__ws),
	5505:  uint16(aux_sym_node_repeat3),
	5506:  uint16(288),
	5507:  uint16(3),
	5508:  uint16(sym_multi_line_comment),
	5509:  uint16(sym__bom),
	5510:  uint16(sym__unicode_space),
	5511:  uint16(151),
	5512:  uint16(3),
	5513:  uint16(sym__node_terminator),
	5514:  uint16(sym__newline),
	5515:  uint16(sym_single_line_comment),
	5516:  uint16(316),
	5517:  uint16(9),
	5518:  uint16(sym__eof),
	5519:  uint16(anon_sym_SEMI),
	5520:  uint16(aux_sym__newline_token1),
	5521:  uint16(aux_sym__newline_token2),
	5522:  uint16(aux_sym__newline_token3),
	5523:  uint16(aux_sym__newline_token4),
	5524:  uint16(aux_sym__newline_token5),
	5525:  uint16(aux_sym__newline_token6),
	5526:  uint16(aux_sym__newline_token7),
	5527:  uint16(9),
	5528:  uint16(19),
	5529:  uint16(1),
	5530:  uint16(anon_sym_SLASH_SLASH),
	5531:  uint16(278),
	5532:  uint16(1),
	5533:  uint16(anon_sym_BSLASH),
	5534:  uint16(187),
	5535:  uint16(1),
	5536:  uint16(aux_sym_node_repeat1),
	5537:  uint16(212),
	5538:  uint16(1),
	5539:  uint16(sym__escline),
	5540:  uint16(234),
	5541:  uint16(1),
	5542:  uint16(sym__node_space),
	5543:  uint16(193),
	5544:  uint16(2),
	5545:  uint16(sym__ws),
	5546:  uint16(aux_sym_node_repeat3),
	5547:  uint16(288),
	5548:  uint16(3),
	5549:  uint16(sym_multi_line_comment),
	5550:  uint16(sym__bom),
	5551:  uint16(sym__unicode_space),
	5552:  uint16(137),
	5553:  uint16(3),
	5554:  uint16(sym__node_terminator),
	5555:  uint16(sym__newline),
	5556:  uint16(sym_single_line_comment),
	5557:  uint16(318),
	5558:  uint16(9),
	5559:  uint16(sym__eof),
	5560:  uint16(anon_sym_SEMI),
	5561:  uint16(aux_sym__newline_token1),
	5562:  uint16(aux_sym__newline_token2),
	5563:  uint16(aux_sym__newline_token3),
	5564:  uint16(aux_sym__newline_token4),
	5565:  uint16(aux_sym__newline_token5),
	5566:  uint16(aux_sym__newline_token6),
	5567:  uint16(aux_sym__newline_token7),
	5568:  uint16(9),
	5569:  uint16(19),
	5570:  uint16(1),
	5571:  uint16(anon_sym_SLASH_SLASH),
	5572:  uint16(278),
	5573:  uint16(1),
	5574:  uint16(anon_sym_BSLASH),
	5575:  uint16(102),
	5576:  uint16(1),
	5577:  uint16(aux_sym_node_repeat1),
	5578:  uint16(222),
	5579:  uint16(1),
	5580:  uint16(sym__escline),
	5581:  uint16(234),
	5582:  uint16(1),
	5583:  uint16(sym__node_space),
	5584:  uint16(195),
	5585:  uint16(2),
	5586:  uint16(sym__ws),
	5587:  uint16(aux_sym_node_repeat3),
	5588:  uint16(300),
	5589:  uint16(3),
	5590:  uint16(sym_multi_line_comment),
	5591:  uint16(sym__bom),
	5592:  uint16(sym__unicode_space),
	5593:  uint16(137),
	5594:  uint16(3),
	5595:  uint16(sym__node_terminator),
	5596:  uint16(sym__newline),
	5597:  uint16(sym_single_line_comment),
	5598:  uint16(318),
	5599:  uint16(9),
	5600:  uint16(sym__eof),
	5601:  uint16(anon_sym_SEMI),
	5602:  uint16(aux_sym__newline_token1),
	5603:  uint16(aux_sym__newline_token2),
	5604:  uint16(aux_sym__newline_token3),
	5605:  uint16(aux_sym__newline_token4),
	5606:  uint16(aux_sym__newline_token5),
	5607:  uint16(aux_sym__newline_token6),
	5608:  uint16(aux_sym__newline_token7),
	5609:  uint16(9),
	5610:  uint16(19),
	5611:  uint16(1),
	5612:  uint16(anon_sym_SLASH_SLASH),
	5613:  uint16(278),
	5614:  uint16(1),
	5615:  uint16(anon_sym_BSLASH),
	5616:  uint16(116),
	5617:  uint16(1),
	5618:  uint16(aux_sym_node_repeat1),
	5619:  uint16(212),
	5620:  uint16(1),
	5621:  uint16(sym__escline),
	5622:  uint16(234),
	5623:  uint16(1),
	5624:  uint16(sym__node_space),
	5625:  uint16(120),
	5626:  uint16(2),
	5627:  uint16(sym__ws),
	5628:  uint16(aux_sym_node_repeat3),
	5629:  uint16(322),
	5630:  uint16(3),
	5631:  uint16(sym_multi_line_comment),
	5632:  uint16(sym__bom),
	5633:  uint16(sym__unicode_space),
	5634:  uint16(138),
	5635:  uint16(3),
	5636:  uint16(sym__node_terminator),
	5637:  uint16(sym__newline),
	5638:  uint16(sym_single_line_comment),
	5639:  uint16(320),
	5640:  uint16(9),
	5641:  uint16(sym__eof),
	5642:  uint16(anon_sym_SEMI),
	5643:  uint16(aux_sym__newline_token1),
	5644:  uint16(aux_sym__newline_token2),
	5645:  uint16(aux_sym__newline_token3),
	5646:  uint16(aux_sym__newline_token4),
	5647:  uint16(aux_sym__newline_token5),
	5648:  uint16(aux_sym__newline_token6),
	5649:  uint16(aux_sym__newline_token7),
	5650:  uint16(9),
	5651:  uint16(19),
	5652:  uint16(1),
	5653:  uint16(anon_sym_SLASH_SLASH),
	5654:  uint16(278),
	5655:  uint16(1),
	5656:  uint16(anon_sym_BSLASH),
	5657:  uint16(187),
	5658:  uint16(1),
	5659:  uint16(aux_sym_node_repeat1),
	5660:  uint16(212),
	5661:  uint16(1),
	5662:  uint16(sym__escline),
	5663:  uint16(234),
	5664:  uint16(1),
	5665:  uint16(sym__node_space),
	5666:  uint16(193),
	5667:  uint16(2),
	5668:  uint16(sym__ws),
	5669:  uint16(aux_sym_node_repeat3),
	5670:  uint16(288),
	5671:  uint16(3),
	5672:  uint16(sym_multi_line_comment),
	5673:  uint16(sym__bom),
	5674:  uint16(sym__unicode_space),
	5675:  uint16(157),
	5676:  uint16(3),
	5677:  uint16(sym__node_terminator),
	5678:  uint16(sym__newline),
	5679:  uint16(sym_single_line_comment),
	5680:  uint16(324),
	5681:  uint16(9),
	5682:  uint16(sym__eof),
	5683:  uint16(anon_sym_SEMI),
	5684:  uint16(aux_sym__newline_token1),
	5685:  uint16(aux_sym__newline_token2),
	5686:  uint16(aux_sym__newline_token3),
	5687:  uint16(aux_sym__newline_token4),
	5688:  uint16(aux_sym__newline_token5),
	5689:  uint16(aux_sym__newline_token6),
	5690:  uint16(aux_sym__newline_token7),
	5691:  uint16(9),
	5692:  uint16(19),
	5693:  uint16(1),
	5694:  uint16(anon_sym_SLASH_SLASH),
	5695:  uint16(278),
	5696:  uint16(1),
	5697:  uint16(anon_sym_BSLASH),
	5698:  uint16(113),
	5699:  uint16(1),
	5700:  uint16(aux_sym_node_repeat1),
	5701:  uint16(212),
	5702:  uint16(1),
	5703:  uint16(sym__escline),
	5704:  uint16(234),
	5705:  uint16(1),
	5706:  uint16(sym__node_space),
	5707:  uint16(112),
	5708:  uint16(2),
	5709:  uint16(sym__ws),
	5710:  uint16(aux_sym_node_repeat3),
	5711:  uint16(328),
	5712:  uint16(3),
	5713:  uint16(sym_multi_line_comment),
	5714:  uint16(sym__bom),
	5715:  uint16(sym__unicode_space),
	5716:  uint16(144),
	5717:  uint16(3),
	5718:  uint16(sym__node_terminator),
	5719:  uint16(sym__newline),
	5720:  uint16(sym_single_line_comment),
	5721:  uint16(326),
	5722:  uint16(9),
	5723:  uint16(sym__eof),
	5724:  uint16(anon_sym_SEMI),
	5725:  uint16(aux_sym__newline_token1),
	5726:  uint16(aux_sym__newline_token2),
	5727:  uint16(aux_sym__newline_token3),
	5728:  uint16(aux_sym__newline_token4),
	5729:  uint16(aux_sym__newline_token5),
	5730:  uint16(aux_sym__newline_token6),
	5731:  uint16(aux_sym__newline_token7),
	5732:  uint16(9),
	5733:  uint16(19),
	5734:  uint16(1),
	5735:  uint16(anon_sym_SLASH_SLASH),
	5736:  uint16(278),
	5737:  uint16(1),
	5738:  uint16(anon_sym_BSLASH),
	5739:  uint16(187),
	5740:  uint16(1),
	5741:  uint16(aux_sym_node_repeat1),
	5742:  uint16(212),
	5743:  uint16(1),
	5744:  uint16(sym__escline),
	5745:  uint16(234),
	5746:  uint16(1),
	5747:  uint16(sym__node_space),
	5748:  uint16(193),
	5749:  uint16(2),
	5750:  uint16(sym__ws),
	5751:  uint16(aux_sym_node_repeat3),
	5752:  uint16(288),
	5753:  uint16(3),
	5754:  uint16(sym_multi_line_comment),
	5755:  uint16(sym__bom),
	5756:  uint16(sym__unicode_space),
	5757:  uint16(145),
	5758:  uint16(3),
	5759:  uint16(sym__node_terminator),
	5760:  uint16(sym__newline),
	5761:  uint16(sym_single_line_comment),
	5762:  uint16(330),
	5763:  uint16(9),
	5764:  uint16(sym__eof),
	5765:  uint16(anon_sym_SEMI),
	5766:  uint16(aux_sym__newline_token1),
	5767:  uint16(aux_sym__newline_token2),
	5768:  uint16(aux_sym__newline_token3),
	5769:  uint16(aux_sym__newline_token4),
	5770:  uint16(aux_sym__newline_token5),
	5771:  uint16(aux_sym__newline_token6),
	5772:  uint16(aux_sym__newline_token7),
	5773:  uint16(3),
	5774:  uint16(99),
	5775:  uint16(1),
	5776:  uint16(aux_sym__integer_repeat1),
	5777:  uint16(334),
	5778:  uint16(2),
	5779:  uint16(anon_sym__),
	5780:  uint16(sym__digit),
	5781:  uint16(332),
	5782:  uint16(19),
	5783:  uint16(sym__eof),
	5784:  uint16(sym_multi_line_comment),
	5785:  uint16(anon_sym_SLASH_DASH),
	5786:  uint16(anon_sym_LBRACE),
	5787:  uint16(anon_sym_SEMI),
	5788:  uint16(anon_sym_DOT),
	5789:  uint16(anon_sym_e),
	5790:  uint16(anon_sym_E),
	5791:  uint16(anon_sym_BSLASH),
	5792:  uint16(aux_sym__newline_token1),
	5793:  uint16(aux_sym__newline_token2),
	5794:  uint16(aux_sym__newline_token3),
	5795:  uint16(aux_sym__newline_token4),
	5796:  uint16(aux_sym__newline_token5),
	5797:  uint16(aux_sym__newline_token6),
	5798:  uint16(aux_sym__newline_token7),
	5799:  uint16(sym__bom),
	5800:  uint16(sym__unicode_space),
	5801:  uint16(anon_sym_SLASH_SLASH),
	5802:  uint16(6),
	5803:  uint16(337),
	5804:  uint16(1),
	5805:  uint16(anon_sym_BSLASH),
	5806:  uint16(121),
	5807:  uint16(1),
	5808:  uint16(sym__escline),
	5809:  uint16(126),
	5810:  uint16(2),
	5811:  uint16(sym__ws),
	5812:  uint16(aux_sym_node_repeat3),
	5813:  uint16(340),
	5814:  uint16(3),
	5815:  uint16(sym_multi_line_comment),
	5816:  uint16(sym__bom),
	5817:  uint16(sym__unicode_space),
	5818:  uint16(147),
	5819:  uint16(5),
	5820:  uint16(sym__normal_bare_identifier),
	5821:  uint16(anon_sym_null),
	5822:  uint16(sym__digit),
	5823:  uint16(anon_sym_true),
	5824:  uint16(anon_sym_false),
	5825:  uint16(149),
	5826:  uint16(10),
	5827:  uint16(sym__raw_string),
	5828:  uint16(anon_sym_SLASH_DASH),
	5829:  uint16(anon_sym_LBRACE),
	5830:  uint16(anon_sym_LPAREN),
	5831:  uint16(anon_sym_DQUOTE),
	5832:  uint16(anon_sym_PLUS),
	5833:  uint16(anon_sym_DASH),
	5834:  uint16(anon_sym_0x),
	5835:  uint16(anon_sym_0o),
	5836:  uint16(anon_sym_0b),
	5837:  uint16(9),
	5838:  uint16(19),
	5839:  uint16(1),
	5840:  uint16(anon_sym_SLASH_SLASH),
	5841:  uint16(278),
	5842:  uint16(1),
	5843:  uint16(anon_sym_BSLASH),
	5844:  uint16(187),
	5845:  uint16(1),
	5846:  uint16(aux_sym_node_repeat1),
	5847:  uint16(212),
	5848:  uint16(1),
	5849:  uint16(sym__escline),
	5850:  uint16(234),
	5851:  uint16(1),
	5852:  uint16(sym__node_space),
	5853:  uint16(193),
	5854:  uint16(2),
	5855:  uint16(sym__ws),
	5856:  uint16(aux_sym_node_repeat3),
	5857:  uint16(288),
	5858:  uint16(3),
	5859:  uint16(sym_multi_line_comment),
	5860:  uint16(sym__bom),
	5861:  uint16(sym__unicode_space),
	5862:  uint16(143),
	5863:  uint16(3),
	5864:  uint16(sym__node_terminator),
	5865:  uint16(sym__newline),
	5866:  uint16(sym_single_line_comment),
	5867:  uint16(343),
	5868:  uint16(9),
	5869:  uint16(sym__eof),
	5870:  uint16(anon_sym_SEMI),
	5871:  uint16(aux_sym__newline_token1),
	5872:  uint16(aux_sym__newline_token2),
	5873:  uint16(aux_sym__newline_token3),
	5874:  uint16(aux_sym__newline_token4),
	5875:  uint16(aux_sym__newline_token5),
	5876:  uint16(aux_sym__newline_token6),
	5877:  uint16(aux_sym__newline_token7),
	5878:  uint16(9),
	5879:  uint16(19),
	5880:  uint16(1),
	5881:  uint16(anon_sym_SLASH_SLASH),
	5882:  uint16(278),
	5883:  uint16(1),
	5884:  uint16(anon_sym_BSLASH),
	5885:  uint16(187),
	5886:  uint16(1),
	5887:  uint16(aux_sym_node_repeat1),
	5888:  uint16(212),
	5889:  uint16(1),
	5890:  uint16(sym__escline),
	5891:  uint16(234),
	5892:  uint16(1),
	5893:  uint16(sym__node_space),
	5894:  uint16(193),
	5895:  uint16(2),
	5896:  uint16(sym__ws),
	5897:  uint16(aux_sym_node_repeat3),
	5898:  uint16(288),
	5899:  uint16(3),
	5900:  uint16(sym_multi_line_comment),
	5901:  uint16(sym__bom),
	5902:  uint16(sym__unicode_space),
	5903:  uint16(167),
	5904:  uint16(3),
	5905:  uint16(sym__node_terminator),
	5906:  uint16(sym__newline),
	5907:  uint16(sym_single_line_comment),
	5908:  uint16(345),
	5909:  uint16(9),
	5910:  uint16(sym__eof),
	5911:  uint16(anon_sym_SEMI),
	5912:  uint16(aux_sym__newline_token1),
	5913:  uint16(aux_sym__newline_token2),
	5914:  uint16(aux_sym__newline_token3),
	5915:  uint16(aux_sym__newline_token4),
	5916:  uint16(aux_sym__newline_token5),
	5917:  uint16(aux_sym__newline_token6),
	5918:  uint16(aux_sym__newline_token7),
	5919:  uint16(9),
	5920:  uint16(19),
	5921:  uint16(1),
	5922:  uint16(anon_sym_SLASH_SLASH),
	5923:  uint16(278),
	5924:  uint16(1),
	5925:  uint16(anon_sym_BSLASH),
	5926:  uint16(84),
	5927:  uint16(1),
	5928:  uint16(aux_sym_node_repeat1),
	5929:  uint16(222),
	5930:  uint16(1),
	5931:  uint16(sym__escline),
	5932:  uint16(234),
	5933:  uint16(1),
	5934:  uint16(sym__node_space),
	5935:  uint16(195),
	5936:  uint16(2),
	5937:  uint16(sym__ws),
	5938:  uint16(aux_sym_node_repeat3),
	5939:  uint16(300),
	5940:  uint16(3),
	5941:  uint16(sym_multi_line_comment),
	5942:  uint16(sym__bom),
	5943:  uint16(sym__unicode_space),
	5944:  uint16(140),
	5945:  uint16(3),
	5946:  uint16(sym__node_terminator),
	5947:  uint16(sym__newline),
	5948:  uint16(sym_single_line_comment),
	5949:  uint16(347),
	5950:  uint16(9),
	5951:  uint16(sym__eof),
	5952:  uint16(anon_sym_SEMI),
	5953:  uint16(aux_sym__newline_token1),
	5954:  uint16(aux_sym__newline_token2),
	5955:  uint16(aux_sym__newline_token3),
	5956:  uint16(aux_sym__newline_token4),
	5957:  uint16(aux_sym__newline_token5),
	5958:  uint16(aux_sym__newline_token6),
	5959:  uint16(aux_sym__newline_token7),
	5960:  uint16(9),
	5961:  uint16(19),
	5962:  uint16(1),
	5963:  uint16(anon_sym_SLASH_SLASH),
	5964:  uint16(278),
	5965:  uint16(1),
	5966:  uint16(anon_sym_BSLASH),
	5967:  uint16(107),
	5968:  uint16(1),
	5969:  uint16(aux_sym_node_repeat1),
	5970:  uint16(212),
	5971:  uint16(1),
	5972:  uint16(sym__escline),
	5973:  uint16(234),
	5974:  uint16(1),
	5975:  uint16(sym__node_space),
	5976:  uint16(106),
	5977:  uint16(2),
	5978:  uint16(sym__ws),
	5979:  uint16(aux_sym_node_repeat3),
	5980:  uint16(351),
	5981:  uint16(3),
	5982:  uint16(sym_multi_line_comment),
	5983:  uint16(sym__bom),
	5984:  uint16(sym__unicode_space),
	5985:  uint16(158),
	5986:  uint16(3),
	5987:  uint16(sym__node_terminator),
	5988:  uint16(sym__newline),
	5989:  uint16(sym_single_line_comment),
	5990:  uint16(349),
	5991:  uint16(9),
	5992:  uint16(sym__eof),
	5993:  uint16(anon_sym_SEMI),
	5994:  uint16(aux_sym__newline_token1),
	5995:  uint16(aux_sym__newline_token2),
	5996:  uint16(aux_sym__newline_token3),
	5997:  uint16(aux_sym__newline_token4),
	5998:  uint16(aux_sym__newline_token5),
	5999:  uint16(aux_sym__newline_token6),
	6000:  uint16(aux_sym__newline_token7),
	6001:  uint16(9),
	6002:  uint16(19),
	6003:  uint16(1),
	6004:  uint16(anon_sym_SLASH_SLASH),
	6005:  uint16(278),
	6006:  uint16(1),
	6007:  uint16(anon_sym_BSLASH),
	6008:  uint16(82),
	6009:  uint16(1),
	6010:  uint16(aux_sym_node_repeat1),
	6011:  uint16(222),
	6012:  uint16(1),
	6013:  uint16(sym__escline),
	6014:  uint16(234),
	6015:  uint16(1),
	6016:  uint16(sym__node_space),
	6017:  uint16(195),
	6018:  uint16(2),
	6019:  uint16(sym__ws),
	6020:  uint16(aux_sym_node_repeat3),
	6021:  uint16(300),
	6022:  uint16(3),
	6023:  uint16(sym_multi_line_comment),
	6024:  uint16(sym__bom),
	6025:  uint16(sym__unicode_space),
	6026:  uint16(143),
	6027:  uint16(3),
	6028:  uint16(sym__node_terminator),
	6029:  uint16(sym__newline),
	6030:  uint16(sym_single_line_comment),
	6031:  uint16(343),
	6032:  uint16(9),
	6033:  uint16(sym__eof),
	6034:  uint16(anon_sym_SEMI),
	6035:  uint16(aux_sym__newline_token1),
	6036:  uint16(aux_sym__newline_token2),
	6037:  uint16(aux_sym__newline_token3),
	6038:  uint16(aux_sym__newline_token4),
	6039:  uint16(aux_sym__newline_token5),
	6040:  uint16(aux_sym__newline_token6),
	6041:  uint16(aux_sym__newline_token7),
	6042:  uint16(9),
	6043:  uint16(19),
	6044:  uint16(1),
	6045:  uint16(anon_sym_SLASH_SLASH),
	6046:  uint16(278),
	6047:  uint16(1),
	6048:  uint16(anon_sym_BSLASH),
	6049:  uint16(115),
	6050:  uint16(1),
	6051:  uint16(aux_sym_node_repeat1),
	6052:  uint16(222),
	6053:  uint16(1),
	6054:  uint16(sym__escline),
	6055:  uint16(234),
	6056:  uint16(1),
	6057:  uint16(sym__node_space),
	6058:  uint16(195),
	6059:  uint16(2),
	6060:  uint16(sym__ws),
	6061:  uint16(aux_sym_node_repeat3),
	6062:  uint16(300),
	6063:  uint16(3),
	6064:  uint16(sym_multi_line_comment),
	6065:  uint16(sym__bom),
	6066:  uint16(sym__unicode_space),
	6067:  uint16(146),
	6068:  uint16(3),
	6069:  uint16(sym__node_terminator),
	6070:  uint16(sym__newline),
	6071:  uint16(sym_single_line_comment),
	6072:  uint16(353),
	6073:  uint16(9),
	6074:  uint16(sym__eof),
	6075:  uint16(anon_sym_SEMI),
	6076:  uint16(aux_sym__newline_token1),
	6077:  uint16(aux_sym__newline_token2),
	6078:  uint16(aux_sym__newline_token3),
	6079:  uint16(aux_sym__newline_token4),
	6080:  uint16(aux_sym__newline_token5),
	6081:  uint16(aux_sym__newline_token6),
	6082:  uint16(aux_sym__newline_token7),
	6083:  uint16(9),
	6084:  uint16(19),
	6085:  uint16(1),
	6086:  uint16(anon_sym_SLASH_SLASH),
	6087:  uint16(278),
	6088:  uint16(1),
	6089:  uint16(anon_sym_BSLASH),
	6090:  uint16(187),
	6091:  uint16(1),
	6092:  uint16(aux_sym_node_repeat1),
	6093:  uint16(212),
	6094:  uint16(1),
	6095:  uint16(sym__escline),
	6096:  uint16(234),
	6097:  uint16(1),
	6098:  uint16(sym__node_space),
	6099:  uint16(193),
	6100:  uint16(2),
	6101:  uint16(sym__ws),
	6102:  uint16(aux_sym_node_repeat3),
	6103:  uint16(288),
	6104:  uint16(3),
	6105:  uint16(sym_multi_line_comment),
	6106:  uint16(sym__bom),
	6107:  uint16(sym__unicode_space),
	6108:  uint16(146),
	6109:  uint16(3),
	6110:  uint16(sym__node_terminator),
	6111:  uint16(sym__newline),
	6112:  uint16(sym_single_line_comment),
	6113:  uint16(353),
	6114:  uint16(9),
	6115:  uint16(sym__eof),
	6116:  uint16(anon_sym_SEMI),
	6117:  uint16(aux_sym__newline_token1),
	6118:  uint16(aux_sym__newline_token2),
	6119:  uint16(aux_sym__newline_token3),
	6120:  uint16(aux_sym__newline_token4),
	6121:  uint16(aux_sym__newline_token5),
	6122:  uint16(aux_sym__newline_token6),
	6123:  uint16(aux_sym__newline_token7),
	6124:  uint16(9),
	6125:  uint16(19),
	6126:  uint16(1),
	6127:  uint16(anon_sym_SLASH_SLASH),
	6128:  uint16(278),
	6129:  uint16(1),
	6130:  uint16(anon_sym_BSLASH),
	6131:  uint16(101),
	6132:  uint16(1),
	6133:  uint16(aux_sym_node_repeat1),
	6134:  uint16(212),
	6135:  uint16(1),
	6136:  uint16(sym__escline),
	6137:  uint16(234),
	6138:  uint16(1),
	6139:  uint16(sym__node_space),
	6140:  uint16(105),
	6141:  uint16(2),
	6142:  uint16(sym__ws),
	6143:  uint16(aux_sym_node_repeat3),
	6144:  uint16(357),
	6145:  uint16(3),
	6146:  uint16(sym_multi_line_comment),
	6147:  uint16(sym__bom),
	6148:  uint16(sym__unicode_space),
	6149:  uint16(163),
	6150:  uint16(3),
	6151:  uint16(sym__node_terminator),
	6152:  uint16(sym__newline),
	6153:  uint16(sym_single_line_comment),
	6154:  uint16(355),
	6155:  uint16(9),
	6156:  uint16(sym__eof),
	6157:  uint16(anon_sym_SEMI),
	6158:  uint16(aux_sym__newline_token1),
	6159:  uint16(aux_sym__newline_token2),
	6160:  uint16(aux_sym__newline_token3),
	6161:  uint16(aux_sym__newline_token4),
	6162:  uint16(aux_sym__newline_token5),
	6163:  uint16(aux_sym__newline_token6),
	6164:  uint16(aux_sym__newline_token7),
	6165:  uint16(9),
	6166:  uint16(19),
	6167:  uint16(1),
	6168:  uint16(anon_sym_SLASH_SLASH),
	6169:  uint16(278),
	6170:  uint16(1),
	6171:  uint16(anon_sym_BSLASH),
	6172:  uint16(93),
	6173:  uint16(1),
	6174:  uint16(aux_sym_node_repeat1),
	6175:  uint16(212),
	6176:  uint16(1),
	6177:  uint16(sym__escline),
	6178:  uint16(234),
	6179:  uint16(1),
	6180:  uint16(sym__node_space),
	6181:  uint16(94),
	6182:  uint16(2),
	6183:  uint16(sym__ws),
	6184:  uint16(aux_sym_node_repeat3),
	6185:  uint16(361),
	6186:  uint16(3),
	6187:  uint16(sym_multi_line_comment),
	6188:  uint16(sym__bom),
	6189:  uint16(sym__unicode_space),
	6190:  uint16(149),
	6191:  uint16(3),
	6192:  uint16(sym__node_terminator),
	6193:  uint16(sym__newline),
	6194:  uint16(sym_single_line_comment),
	6195:  uint16(359),
	6196:  uint16(9),
	6197:  uint16(sym__eof),
	6198:  uint16(anon_sym_SEMI),
	6199:  uint16(aux_sym__newline_token1),
	6200:  uint16(aux_sym__newline_token2),
	6201:  uint16(aux_sym__newline_token3),
	6202:  uint16(aux_sym__newline_token4),
	6203:  uint16(aux_sym__newline_token5),
	6204:  uint16(aux_sym__newline_token6),
	6205:  uint16(aux_sym__newline_token7),
	6206:  uint16(9),
	6207:  uint16(19),
	6208:  uint16(1),
	6209:  uint16(anon_sym_SLASH_SLASH),
	6210:  uint16(278),
	6211:  uint16(1),
	6212:  uint16(anon_sym_BSLASH),
	6213:  uint16(187),
	6214:  uint16(1),
	6215:  uint16(aux_sym_node_repeat1),
	6216:  uint16(212),
	6217:  uint16(1),
	6218:  uint16(sym__escline),
	6219:  uint16(234),
	6220:  uint16(1),
	6221:  uint16(sym__node_space),
	6222:  uint16(193),
	6223:  uint16(2),
	6224:  uint16(sym__ws),
	6225:  uint16(aux_sym_node_repeat3),
	6226:  uint16(288),
	6227:  uint16(3),
	6228:  uint16(sym_multi_line_comment),
	6229:  uint16(sym__bom),
	6230:  uint16(sym__unicode_space),
	6231:  uint16(150),
	6232:  uint16(3),
	6233:  uint16(sym__node_terminator),
	6234:  uint16(sym__newline),
	6235:  uint16(sym_single_line_comment),
	6236:  uint16(363),
	6237:  uint16(9),
	6238:  uint16(sym__eof),
	6239:  uint16(anon_sym_SEMI),
	6240:  uint16(aux_sym__newline_token1),
	6241:  uint16(aux_sym__newline_token2),
	6242:  uint16(aux_sym__newline_token3),
	6243:  uint16(aux_sym__newline_token4),
	6244:  uint16(aux_sym__newline_token5),
	6245:  uint16(aux_sym__newline_token6),
	6246:  uint16(aux_sym__newline_token7),
	6247:  uint16(9),
	6248:  uint16(19),
	6249:  uint16(1),
	6250:  uint16(anon_sym_SLASH_SLASH),
	6251:  uint16(278),
	6252:  uint16(1),
	6253:  uint16(anon_sym_BSLASH),
	6254:  uint16(187),
	6255:  uint16(1),
	6256:  uint16(aux_sym_node_repeat1),
	6257:  uint16(212),
	6258:  uint16(1),
	6259:  uint16(sym__escline),
	6260:  uint16(234),
	6261:  uint16(1),
	6262:  uint16(sym__node_space),
	6263:  uint16(193),
	6264:  uint16(2),
	6265:  uint16(sym__ws),
	6266:  uint16(aux_sym_node_repeat3),
	6267:  uint16(288),
	6268:  uint16(3),
	6269:  uint16(sym_multi_line_comment),
	6270:  uint16(sym__bom),
	6271:  uint16(sym__unicode_space),
	6272:  uint16(140),
	6273:  uint16(3),
	6274:  uint16(sym__node_terminator),
	6275:  uint16(sym__newline),
	6276:  uint16(sym_single_line_comment),
	6277:  uint16(347),
	6278:  uint16(9),
	6279:  uint16(sym__eof),
	6280:  uint16(anon_sym_SEMI),
	6281:  uint16(aux_sym__newline_token1),
	6282:  uint16(aux_sym__newline_token2),
	6283:  uint16(aux_sym__newline_token3),
	6284:  uint16(aux_sym__newline_token4),
	6285:  uint16(aux_sym__newline_token5),
	6286:  uint16(aux_sym__newline_token6),
	6287:  uint16(aux_sym__newline_token7),
	6288:  uint16(9),
	6289:  uint16(19),
	6290:  uint16(1),
	6291:  uint16(anon_sym_SLASH_SLASH),
	6292:  uint16(278),
	6293:  uint16(1),
	6294:  uint16(anon_sym_BSLASH),
	6295:  uint16(88),
	6296:  uint16(1),
	6297:  uint16(aux_sym_node_repeat1),
	6298:  uint16(222),
	6299:  uint16(1),
	6300:  uint16(sym__escline),
	6301:  uint16(234),
	6302:  uint16(1),
	6303:  uint16(sym__node_space),
	6304:  uint16(195),
	6305:  uint16(2),
	6306:  uint16(sym__ws),
	6307:  uint16(aux_sym_node_repeat3),
	6308:  uint16(300),
	6309:  uint16(3),
	6310:  uint16(sym_multi_line_comment),
	6311:  uint16(sym__bom),
	6312:  uint16(sym__unicode_space),
	6313:  uint16(159),
	6314:  uint16(3),
	6315:  uint16(sym__node_terminator),
	6316:  uint16(sym__newline),
	6317:  uint16(sym_single_line_comment),
	6318:  uint16(365),
	6319:  uint16(9),
	6320:  uint16(sym__eof),
	6321:  uint16(anon_sym_SEMI),
	6322:  uint16(aux_sym__newline_token1),
	6323:  uint16(aux_sym__newline_token2),
	6324:  uint16(aux_sym__newline_token3),
	6325:  uint16(aux_sym__newline_token4),
	6326:  uint16(aux_sym__newline_token5),
	6327:  uint16(aux_sym__newline_token6),
	6328:  uint16(aux_sym__newline_token7),
	6329:  uint16(9),
	6330:  uint16(19),
	6331:  uint16(1),
	6332:  uint16(anon_sym_SLASH_SLASH),
	6333:  uint16(278),
	6334:  uint16(1),
	6335:  uint16(anon_sym_BSLASH),
	6336:  uint16(187),
	6337:  uint16(1),
	6338:  uint16(aux_sym_node_repeat1),
	6339:  uint16(212),
	6340:  uint16(1),
	6341:  uint16(sym__escline),
	6342:  uint16(234),
	6343:  uint16(1),
	6344:  uint16(sym__node_space),
	6345:  uint16(193),
	6346:  uint16(2),
	6347:  uint16(sym__ws),
	6348:  uint16(aux_sym_node_repeat3),
	6349:  uint16(288),
	6350:  uint16(3),
	6351:  uint16(sym_multi_line_comment),
	6352:  uint16(sym__bom),
	6353:  uint16(sym__unicode_space),
	6354:  uint16(159),
	6355:  uint16(3),
	6356:  uint16(sym__node_terminator),
	6357:  uint16(sym__newline),
	6358:  uint16(sym_single_line_comment),
	6359:  uint16(365),
	6360:  uint16(9),
	6361:  uint16(sym__eof),
	6362:  uint16(anon_sym_SEMI),
	6363:  uint16(aux_sym__newline_token1),
	6364:  uint16(aux_sym__newline_token2),
	6365:  uint16(aux_sym__newline_token3),
	6366:  uint16(aux_sym__newline_token4),
	6367:  uint16(aux_sym__newline_token5),
	6368:  uint16(aux_sym__newline_token6),
	6369:  uint16(aux_sym__newline_token7),
	6370:  uint16(9),
	6371:  uint16(19),
	6372:  uint16(1),
	6373:  uint16(anon_sym_SLASH_SLASH),
	6374:  uint16(278),
	6375:  uint16(1),
	6376:  uint16(anon_sym_BSLASH),
	6377:  uint16(110),
	6378:  uint16(1),
	6379:  uint16(aux_sym_node_repeat1),
	6380:  uint16(222),
	6381:  uint16(1),
	6382:  uint16(sym__escline),
	6383:  uint16(234),
	6384:  uint16(1),
	6385:  uint16(sym__node_space),
	6386:  uint16(195),
	6387:  uint16(2),
	6388:  uint16(sym__ws),
	6389:  uint16(aux_sym_node_repeat3),
	6390:  uint16(300),
	6391:  uint16(3),
	6392:  uint16(sym_multi_line_comment),
	6393:  uint16(sym__bom),
	6394:  uint16(sym__unicode_space),
	6395:  uint16(139),
	6396:  uint16(3),
	6397:  uint16(sym__node_terminator),
	6398:  uint16(sym__newline),
	6399:  uint16(sym_single_line_comment),
	6400:  uint16(367),
	6401:  uint16(9),
	6402:  uint16(sym__eof),
	6403:  uint16(anon_sym_SEMI),
	6404:  uint16(aux_sym__newline_token1),
	6405:  uint16(aux_sym__newline_token2),
	6406:  uint16(aux_sym__newline_token3),
	6407:  uint16(aux_sym__newline_token4),
	6408:  uint16(aux_sym__newline_token5),
	6409:  uint16(aux_sym__newline_token6),
	6410:  uint16(aux_sym__newline_token7),
	6411:  uint16(9),
	6412:  uint16(19),
	6413:  uint16(1),
	6414:  uint16(anon_sym_SLASH_SLASH),
	6415:  uint16(278),
	6416:  uint16(1),
	6417:  uint16(anon_sym_BSLASH),
	6418:  uint16(187),
	6419:  uint16(1),
	6420:  uint16(aux_sym_node_repeat1),
	6421:  uint16(212),
	6422:  uint16(1),
	6423:  uint16(sym__escline),
	6424:  uint16(234),
	6425:  uint16(1),
	6426:  uint16(sym__node_space),
	6427:  uint16(193),
	6428:  uint16(2),
	6429:  uint16(sym__ws),
	6430:  uint16(aux_sym_node_repeat3),
	6431:  uint16(288),
	6432:  uint16(3),
	6433:  uint16(sym_multi_line_comment),
	6434:  uint16(sym__bom),
	6435:  uint16(sym__unicode_space),
	6436:  uint16(161),
	6437:  uint16(3),
	6438:  uint16(sym__node_terminator),
	6439:  uint16(sym__newline),
	6440:  uint16(sym_single_line_comment),
	6441:  uint16(369),
	6442:  uint16(9),
	6443:  uint16(sym__eof),
	6444:  uint16(anon_sym_SEMI),
	6445:  uint16(aux_sym__newline_token1),
	6446:  uint16(aux_sym__newline_token2),
	6447:  uint16(aux_sym__newline_token3),
	6448:  uint16(aux_sym__newline_token4),
	6449:  uint16(aux_sym__newline_token5),
	6450:  uint16(aux_sym__newline_token6),
	6451:  uint16(aux_sym__newline_token7),
	6452:  uint16(9),
	6453:  uint16(19),
	6454:  uint16(1),
	6455:  uint16(anon_sym_SLASH_SLASH),
	6456:  uint16(278),
	6457:  uint16(1),
	6458:  uint16(anon_sym_BSLASH),
	6459:  uint16(187),
	6460:  uint16(1),
	6461:  uint16(aux_sym_node_repeat1),
	6462:  uint16(212),
	6463:  uint16(1),
	6464:  uint16(sym__escline),
	6465:  uint16(234),
	6466:  uint16(1),
	6467:  uint16(sym__node_space),
	6468:  uint16(193),
	6469:  uint16(2),
	6470:  uint16(sym__ws),
	6471:  uint16(aux_sym_node_repeat3),
	6472:  uint16(288),
	6473:  uint16(3),
	6474:  uint16(sym_multi_line_comment),
	6475:  uint16(sym__bom),
	6476:  uint16(sym__unicode_space),
	6477:  uint16(168),
	6478:  uint16(3),
	6479:  uint16(sym__node_terminator),
	6480:  uint16(sym__newline),
	6481:  uint16(sym_single_line_comment),
	6482:  uint16(371),
	6483:  uint16(9),
	6484:  uint16(sym__eof),
	6485:  uint16(anon_sym_SEMI),
	6486:  uint16(aux_sym__newline_token1),
	6487:  uint16(aux_sym__newline_token2),
	6488:  uint16(aux_sym__newline_token3),
	6489:  uint16(aux_sym__newline_token4),
	6490:  uint16(aux_sym__newline_token5),
	6491:  uint16(aux_sym__newline_token6),
	6492:  uint16(aux_sym__newline_token7),
	6493:  uint16(9),
	6494:  uint16(19),
	6495:  uint16(1),
	6496:  uint16(anon_sym_SLASH_SLASH),
	6497:  uint16(278),
	6498:  uint16(1),
	6499:  uint16(anon_sym_BSLASH),
	6500:  uint16(187),
	6501:  uint16(1),
	6502:  uint16(aux_sym_node_repeat1),
	6503:  uint16(212),
	6504:  uint16(1),
	6505:  uint16(sym__escline),
	6506:  uint16(234),
	6507:  uint16(1),
	6508:  uint16(sym__node_space),
	6509:  uint16(193),
	6510:  uint16(2),
	6511:  uint16(sym__ws),
	6512:  uint16(aux_sym_node_repeat3),
	6513:  uint16(288),
	6514:  uint16(3),
	6515:  uint16(sym_multi_line_comment),
	6516:  uint16(sym__bom),
	6517:  uint16(sym__unicode_space),
	6518:  uint16(139),
	6519:  uint16(3),
	6520:  uint16(sym__node_terminator),
	6521:  uint16(sym__newline),
	6522:  uint16(sym_single_line_comment),
	6523:  uint16(367),
	6524:  uint16(9),
	6525:  uint16(sym__eof),
	6526:  uint16(anon_sym_SEMI),
	6527:  uint16(aux_sym__newline_token1),
	6528:  uint16(aux_sym__newline_token2),
	6529:  uint16(aux_sym__newline_token3),
	6530:  uint16(aux_sym__newline_token4),
	6531:  uint16(aux_sym__newline_token5),
	6532:  uint16(aux_sym__newline_token6),
	6533:  uint16(aux_sym__newline_token7),
	6534:  uint16(9),
	6535:  uint16(19),
	6536:  uint16(1),
	6537:  uint16(anon_sym_SLASH_SLASH),
	6538:  uint16(278),
	6539:  uint16(1),
	6540:  uint16(anon_sym_BSLASH),
	6541:  uint16(92),
	6542:  uint16(1),
	6543:  uint16(aux_sym_node_repeat1),
	6544:  uint16(212),
	6545:  uint16(1),
	6546:  uint16(sym__escline),
	6547:  uint16(234),
	6548:  uint16(1),
	6549:  uint16(sym__node_space),
	6550:  uint16(91),
	6551:  uint16(2),
	6552:  uint16(sym__ws),
	6553:  uint16(aux_sym_node_repeat3),
	6554:  uint16(375),
	6555:  uint16(3),
	6556:  uint16(sym_multi_line_comment),
	6557:  uint16(sym__bom),
	6558:  uint16(sym__unicode_space),
	6559:  uint16(170),
	6560:  uint16(3),
	6561:  uint16(sym__node_terminator),
	6562:  uint16(sym__newline),
	6563:  uint16(sym_single_line_comment),
	6564:  uint16(373),
	6565:  uint16(9),
	6566:  uint16(sym__eof),
	6567:  uint16(anon_sym_SEMI),
	6568:  uint16(aux_sym__newline_token1),
	6569:  uint16(aux_sym__newline_token2),
	6570:  uint16(aux_sym__newline_token3),
	6571:  uint16(aux_sym__newline_token4),
	6572:  uint16(aux_sym__newline_token5),
	6573:  uint16(aux_sym__newline_token6),
	6574:  uint16(aux_sym__newline_token7),
	6575:  uint16(3),
	6576:  uint16(99),
	6577:  uint16(1),
	6578:  uint16(aux_sym__integer_repeat1),
	6579:  uint16(379),
	6580:  uint16(2),
	6581:  uint16(anon_sym__),
	6582:  uint16(sym__digit),
	6583:  uint16(377),
	6584:  uint16(19),
	6585:  uint16(sym__eof),
	6586:  uint16(sym_multi_line_comment),
	6587:  uint16(anon_sym_SLASH_DASH),
	6588:  uint16(anon_sym_LBRACE),
	6589:  uint16(anon_sym_SEMI),
	6590:  uint16(anon_sym_DOT),
	6591:  uint16(anon_sym_e),
	6592:  uint16(anon_sym_E),
	6593:  uint16(anon_sym_BSLASH),
	6594:  uint16(aux_sym__newline_token1),
	6595:  uint16(aux_sym__newline_token2),
	6596:  uint16(aux_sym__newline_token3),
	6597:  uint16(aux_sym__newline_token4),
	6598:  uint16(aux_sym__newline_token5),
	6599:  uint16(aux_sym__newline_token6),
	6600:  uint16(aux_sym__newline_token7),
	6601:  uint16(sym__bom),
	6602:  uint16(sym__unicode_space),
	6603:  uint16(anon_sym_SLASH_SLASH),
	6604:  uint16(9),
	6605:  uint16(19),
	6606:  uint16(1),
	6607:  uint16(anon_sym_SLASH_SLASH),
	6608:  uint16(278),
	6609:  uint16(1),
	6610:  uint16(anon_sym_BSLASH),
	6611:  uint16(96),
	6612:  uint16(1),
	6613:  uint16(aux_sym_node_repeat1),
	6614:  uint16(222),
	6615:  uint16(1),
	6616:  uint16(sym__escline),
	6617:  uint16(234),
	6618:  uint16(1),
	6619:  uint16(sym__node_space),
	6620:  uint16(195),
	6621:  uint16(2),
	6622:  uint16(sym__ws),
	6623:  uint16(aux_sym_node_repeat3),
	6624:  uint16(300),
	6625:  uint16(3),
	6626:  uint16(sym_multi_line_comment),
	6627:  uint16(sym__bom),
	6628:  uint16(sym__unicode_space),
	6629:  uint16(168),
	6630:  uint16(3),
	6631:  uint16(sym__node_terminator),
	6632:  uint16(sym__newline),
	6633:  uint16(sym_single_line_comment),
	6634:  uint16(371),
	6635:  uint16(9),
	6636:  uint16(sym__eof),
	6637:  uint16(anon_sym_SEMI),
	6638:  uint16(aux_sym__newline_token1),
	6639:  uint16(aux_sym__newline_token2),
	6640:  uint16(aux_sym__newline_token3),
	6641:  uint16(aux_sym__newline_token4),
	6642:  uint16(aux_sym__newline_token5),
	6643:  uint16(aux_sym__newline_token6),
	6644:  uint16(aux_sym__newline_token7),
	6645:  uint16(4),
	6646:  uint16(123),
	6647:  uint16(2),
	6648:  uint16(sym__ws),
	6649:  uint16(aux_sym_node_repeat3),
	6650:  uint16(381),
	6651:  uint16(3),
	6652:  uint16(sym_multi_line_comment),
	6653:  uint16(sym__bom),
	6654:  uint16(sym__unicode_space),
	6655:  uint16(185),
	6656:  uint16(5),
	6657:  uint16(sym__normal_bare_identifier),
	6658:  uint16(anon_sym_null),
	6659:  uint16(sym__digit),
	6660:  uint16(anon_sym_true),
	6661:  uint16(anon_sym_false),
	6662:  uint16(187),
	6663:  uint16(11),
	6664:  uint16(sym__raw_string),
	6665:  uint16(anon_sym_SLASH_DASH),
	6666:  uint16(anon_sym_LBRACE),
	6667:  uint16(anon_sym_LPAREN),
	6668:  uint16(anon_sym_DQUOTE),
	6669:  uint16(anon_sym_PLUS),
	6670:  uint16(anon_sym_DASH),
	6671:  uint16(anon_sym_0x),
	6672:  uint16(anon_sym_0o),
	6673:  uint16(anon_sym_0b),
	6674:  uint16(anon_sym_BSLASH),
	6675:  uint16(4),
	6676:  uint16(126),
	6677:  uint16(2),
	6678:  uint16(sym__ws),
	6679:  uint16(aux_sym_node_repeat3),
	6680:  uint16(384),
	6681:  uint16(3),
	6682:  uint16(sym_multi_line_comment),
	6683:  uint16(sym__bom),
	6684:  uint16(sym__unicode_space),
	6685:  uint16(185),
	6686:  uint16(5),
	6687:  uint16(sym__normal_bare_identifier),
	6688:  uint16(anon_sym_null),
	6689:  uint16(sym__digit),
	6690:  uint16(anon_sym_true),
	6691:  uint16(anon_sym_false),
	6692:  uint16(187),
	6693:  uint16(11),
	6694:  uint16(sym__raw_string),
	6695:  uint16(anon_sym_SLASH_DASH),
	6696:  uint16(anon_sym_LBRACE),
	6697:  uint16(anon_sym_LPAREN),
	6698:  uint16(anon_sym_DQUOTE),
	6699:  uint16(anon_sym_PLUS),
	6700:  uint16(anon_sym_DASH),
	6701:  uint16(anon_sym_0x),
	6702:  uint16(anon_sym_0o),
	6703:  uint16(anon_sym_0b),
	6704:  uint16(anon_sym_BSLASH),
	6705:  uint16(4),
	6706:  uint16(126),
	6707:  uint16(2),
	6708:  uint16(sym__ws),
	6709:  uint16(aux_sym_node_repeat3),
	6710:  uint16(387),
	6711:  uint16(3),
	6712:  uint16(sym_multi_line_comment),
	6713:  uint16(sym__bom),
	6714:  uint16(sym__unicode_space),
	6715:  uint16(178),
	6716:  uint16(5),
	6717:  uint16(sym__normal_bare_identifier),
	6718:  uint16(anon_sym_null),
	6719:  uint16(sym__digit),
	6720:  uint16(anon_sym_true),
	6721:  uint16(anon_sym_false),
	6722:  uint16(180),
	6723:  uint16(11),
	6724:  uint16(sym__raw_string),
	6725:  uint16(anon_sym_SLASH_DASH),
	6726:  uint16(anon_sym_LBRACE),
	6727:  uint16(anon_sym_LPAREN),
	6728:  uint16(anon_sym_DQUOTE),
	6729:  uint16(anon_sym_PLUS),
	6730:  uint16(anon_sym_DASH),
	6731:  uint16(anon_sym_0x),
	6732:  uint16(anon_sym_0o),
	6733:  uint16(anon_sym_0b),
	6734:  uint16(anon_sym_BSLASH),
	6735:  uint16(16),
	6736:  uint16(7),
	6737:  uint16(1),
	6738:  uint16(sym__normal_bare_identifier),
	6739:  uint16(11),
	6740:  uint16(1),
	6741:  uint16(anon_sym_LPAREN),
	6742:  uint16(13),
	6743:  uint16(1),
	6744:  uint16(anon_sym_DQUOTE),
	6745:  uint16(21),
	6746:  uint16(1),
	6747:  uint16(sym__raw_string),
	6748:  uint16(71),
	6749:  uint16(1),
	6750:  uint16(anon_sym_BSLASH),
	6751:  uint16(63),
	6752:  uint16(1),
	6753:  uint16(sym_identifier),
	6754:  uint16(74),
	6755:  uint16(1),
	6756:  uint16(aux_sym_node_repeat1),
	6757:  uint16(125),
	6758:  uint16(1),
	6759:  uint16(sym__escline),
	6760:  uint16(186),
	6761:  uint16(1),
	6762:  uint16(sym__node_space),
	6763:  uint16(192),
	6764:  uint16(1),
	6765:  uint16(sym__sign),
	6766:  uint16(194),
	6767:  uint16(1),
	6768:  uint16(sym__escaped_string),
	6769:  uint16(257),
	6770:  uint16(1),
	6771:  uint16(sym_type),
	6772:  uint16(15),
	6773:  uint16(2),
	6774:  uint16(anon_sym_PLUS),
	6775:  uint16(anon_sym_DASH),
	6776:  uint16(100),
	6777:  uint16(2),
	6778:  uint16(sym__ws),
	6779:  uint16(aux_sym_node_repeat3),
	6780:  uint16(218),
	6781:  uint16(2),
	6782:  uint16(sym__bare_identifier),
	6783:  uint16(sym_string),
	6784:  uint16(73),
	6785:  uint16(3),
	6786:  uint16(sym_multi_line_comment),
	6787:  uint16(sym__bom),
	6788:  uint16(sym__unicode_space),
	6789:  uint16(4),
	6790:  uint16(122),
	6791:  uint16(2),
	6792:  uint16(sym__ws),
	6793:  uint16(aux_sym_node_repeat3),
	6794:  uint16(390),
	6795:  uint16(3),
	6796:  uint16(sym_multi_line_comment),
	6797:  uint16(sym__bom),
	6798:  uint16(sym__unicode_space),
	6799:  uint16(147),
	6800:  uint16(5),
	6801:  uint16(sym__normal_bare_identifier),
	6802:  uint16(anon_sym_null),
	6803:  uint16(sym__digit),
	6804:  uint16(anon_sym_true),
	6805:  uint16(anon_sym_false),
	6806:  uint16(149),
	6807:  uint16(11),
	6808:  uint16(sym__raw_string),
	6809:  uint16(anon_sym_SLASH_DASH),
	6810:  uint16(anon_sym_LBRACE),
	6811:  uint16(anon_sym_LPAREN),
	6812:  uint16(anon_sym_DQUOTE),
	6813:  uint16(anon_sym_PLUS),
	6814:  uint16(anon_sym_DASH),
	6815:  uint16(anon_sym_0x),
	6816:  uint16(anon_sym_0o),
	6817:  uint16(anon_sym_0b),
	6818:  uint16(anon_sym_BSLASH),
	6819:  uint16(4),
	6820:  uint16(126),
	6821:  uint16(2),
	6822:  uint16(sym__ws),
	6823:  uint16(aux_sym_node_repeat3),
	6824:  uint16(393),
	6825:  uint16(3),
	6826:  uint16(sym_multi_line_comment),
	6827:  uint16(sym__bom),
	6828:  uint16(sym__unicode_space),
	6829:  uint16(206),
	6830:  uint16(5),
	6831:  uint16(sym__normal_bare_identifier),
	6832:  uint16(anon_sym_null),
	6833:  uint16(sym__digit),
	6834:  uint16(anon_sym_true),
	6835:  uint16(anon_sym_false),
	6836:  uint16(208),
	6837:  uint16(11),
	6838:  uint16(sym__raw_string),
	6839:  uint16(anon_sym_SLASH_DASH),
	6840:  uint16(anon_sym_LBRACE),
	6841:  uint16(anon_sym_LPAREN),
	6842:  uint16(anon_sym_DQUOTE),
	6843:  uint16(anon_sym_PLUS),
	6844:  uint16(anon_sym_DASH),
	6845:  uint16(anon_sym_0x),
	6846:  uint16(anon_sym_0o),
	6847:  uint16(anon_sym_0b),
	6848:  uint16(anon_sym_BSLASH),
	6849:  uint16(16),
	6850:  uint16(7),
	6851:  uint16(1),
	6852:  uint16(sym__normal_bare_identifier),
	6853:  uint16(11),
	6854:  uint16(1),
	6855:  uint16(anon_sym_LPAREN),
	6856:  uint16(13),
	6857:  uint16(1),
	6858:  uint16(anon_sym_DQUOTE),
	6859:  uint16(21),
	6860:  uint16(1),
	6861:  uint16(sym__raw_string),
	6862:  uint16(71),
	6863:  uint16(1),
	6864:  uint16(anon_sym_BSLASH),
	6865:  uint16(66),
	6866:  uint16(1),
	6867:  uint16(sym_identifier),
	6868:  uint16(124),
	6869:  uint16(1),
	6870:  uint16(aux_sym_node_repeat1),
	6871:  uint16(125),
	6872:  uint16(1),
	6873:  uint16(sym__escline),
	6874:  uint16(186),
	6875:  uint16(1),
	6876:  uint16(sym__node_space),
	6877:  uint16(192),
	6878:  uint16(1),
	6879:  uint16(sym__sign),
	6880:  uint16(194),
	6881:  uint16(1),
	6882:  uint16(sym__escaped_string),
	6883:  uint16(258),
	6884:  uint16(1),
	6885:  uint16(sym_type),
	6886:  uint16(15),
	6887:  uint16(2),
	6888:  uint16(anon_sym_PLUS),
	6889:  uint16(anon_sym_DASH),
	6890:  uint16(100),
	6891:  uint16(2),
	6892:  uint16(sym__ws),
	6893:  uint16(aux_sym_node_repeat3),
	6894:  uint16(218),
	6895:  uint16(2),
	6896:  uint16(sym__bare_identifier),
	6897:  uint16(sym_string),
	6898:  uint16(73),
	6899:  uint16(3),
	6900:  uint16(sym_multi_line_comment),
	6901:  uint16(sym__bom),
	6902:  uint16(sym__unicode_space),
	6903:  uint16(1),
	6904:  uint16(396),
	6905:  uint16(20),
	6906:  uint16(sym_multi_line_comment),
	6907:  uint16(sym__raw_string),
	6909:  uint16(anon_sym_SLASH_DASH),
	6910:  uint16(anon_sym_RBRACE),
	6911:  uint16(sym__normal_bare_identifier),
	6912:  uint16(anon_sym_LPAREN),
	6913:  uint16(anon_sym_DQUOTE),
	6914:  uint16(anon_sym_PLUS),
	6915:  uint16(anon_sym_DASH),
	6916:  uint16(aux_sym__newline_token1),
	6917:  uint16(aux_sym__newline_token2),
	6918:  uint16(aux_sym__newline_token3),
	6919:  uint16(aux_sym__newline_token4),
	6920:  uint16(aux_sym__newline_token5),
	6921:  uint16(aux_sym__newline_token6),
	6922:  uint16(aux_sym__newline_token7),
	6923:  uint16(sym__bom),
	6924:  uint16(sym__unicode_space),
	6925:  uint16(anon_sym_SLASH_SLASH),
	6926:  uint16(1),
	6927:  uint16(398),
	6928:  uint16(20),
	6929:  uint16(sym_multi_line_comment),
	6930:  uint16(sym__raw_string),
	6932:  uint16(anon_sym_SLASH_DASH),
	6933:  uint16(anon_sym_RBRACE),
	6934:  uint16(sym__normal_bare_identifier),
	6935:  uint16(anon_sym_LPAREN),
	6936:  uint16(anon_sym_DQUOTE),
	6937:  uint16(anon_sym_PLUS),
	6938:  uint16(anon_sym_DASH),
	6939:  uint16(aux_sym__newline_token1),
	6940:  uint16(aux_sym__newline_token2),
	6941:  uint16(aux_sym__newline_token3),
	6942:  uint16(aux_sym__newline_token4),
	6943:  uint16(aux_sym__newline_token5),
	6944:  uint16(aux_sym__newline_token6),
	6945:  uint16(aux_sym__newline_token7),
	6946:  uint16(sym__bom),
	6947:  uint16(sym__unicode_space),
	6948:  uint16(anon_sym_SLASH_SLASH),
	6949:  uint16(3),
	6950:  uint16(130),
	6951:  uint16(1),
	6952:  uint16(aux_sym__binary_repeat1),
	6953:  uint16(402),
	6954:  uint16(3),
	6955:  uint16(anon_sym__),
	6956:  uint16(anon_sym_0),
	6957:  uint16(anon_sym_1),
	6958:  uint16(400),
	6959:  uint16(16),
	6960:  uint16(sym__eof),
	6961:  uint16(sym_multi_line_comment),
	6962:  uint16(anon_sym_SLASH_DASH),
	6963:  uint16(anon_sym_LBRACE),
	6964:  uint16(anon_sym_SEMI),
	6965:  uint16(anon_sym_BSLASH),
	6966:  uint16(aux_sym__newline_token1),
	6967:  uint16(aux_sym__newline_token2),
	6968:  uint16(aux_sym__newline_token3),
	6969:  uint16(aux_sym__newline_token4),
	6970:  uint16(aux_sym__newline_token5),
	6971:  uint16(aux_sym__newline_token6),
	6972:  uint16(aux_sym__newline_token7),
	6973:  uint16(sym__bom),
	6974:  uint16(sym__unicode_space),
	6975:  uint16(anon_sym_SLASH_SLASH),
	6976:  uint16(1),
	6977:  uint16(405),
	6978:  uint16(20),
	6979:  uint16(sym_multi_line_comment),
	6980:  uint16(sym__raw_string),
	6982:  uint16(anon_sym_SLASH_DASH),
	6983:  uint16(anon_sym_RBRACE),
	6984:  uint16(sym__normal_bare_identifier),
	6985:  uint16(anon_sym_LPAREN),
	6986:  uint16(anon_sym_DQUOTE),
	6987:  uint16(anon_sym_PLUS),
	6988:  uint16(anon_sym_DASH),
	6989:  uint16(aux_sym__newline_token1),
	6990:  uint16(aux_sym__newline_token2),
	6991:  uint16(aux_sym__newline_token3),
	6992:  uint16(aux_sym__newline_token4),
	6993:  uint16(aux_sym__newline_token5),
	6994:  uint16(aux_sym__newline_token6),
	6995:  uint16(aux_sym__newline_token7),
	6996:  uint16(sym__bom),
	6997:  uint16(sym__unicode_space),
	6998:  uint16(anon_sym_SLASH_SLASH),
	6999:  uint16(1),
	7000:  uint16(407),
	7001:  uint16(20),
	7002:  uint16(sym_multi_line_comment),
	7003:  uint16(sym__raw_string),
	7005:  uint16(anon_sym_SLASH_DASH),
	7006:  uint16(anon_sym_RBRACE),
	7007:  uint16(sym__normal_bare_identifier),
	7008:  uint16(anon_sym_LPAREN),
	7009:  uint16(anon_sym_DQUOTE),
	7010:  uint16(anon_sym_PLUS),
	7011:  uint16(anon_sym_DASH),
	7012:  uint16(aux_sym__newline_token1),
	7013:  uint16(aux_sym__newline_token2),
	7014:  uint16(aux_sym__newline_token3),
	7015:  uint16(aux_sym__newline_token4),
	7016:  uint16(aux_sym__newline_token5),
	7017:  uint16(aux_sym__newline_token6),
	7018:  uint16(aux_sym__newline_token7),
	7019:  uint16(sym__bom),
	7020:  uint16(sym__unicode_space),
	7021:  uint16(anon_sym_SLASH_SLASH),
	7022:  uint16(1),
	7023:  uint16(409),
	7024:  uint16(20),
	7025:  uint16(sym_multi_line_comment),
	7026:  uint16(sym__raw_string),
	7028:  uint16(anon_sym_SLASH_DASH),
	7029:  uint16(anon_sym_RBRACE),
	7030:  uint16(sym__normal_bare_identifier),
	7031:  uint16(anon_sym_LPAREN),
	7032:  uint16(anon_sym_DQUOTE),
	7033:  uint16(anon_sym_PLUS),
	7034:  uint16(anon_sym_DASH),
	7035:  uint16(aux_sym__newline_token1),
	7036:  uint16(aux_sym__newline_token2),
	7037:  uint16(aux_sym__newline_token3),
	7038:  uint16(aux_sym__newline_token4),
	7039:  uint16(aux_sym__newline_token5),
	7040:  uint16(aux_sym__newline_token6),
	7041:  uint16(aux_sym__newline_token7),
	7042:  uint16(sym__bom),
	7043:  uint16(sym__unicode_space),
	7044:  uint16(anon_sym_SLASH_SLASH),
	7045:  uint16(1),
	7046:  uint16(411),
	7047:  uint16(20),
	7048:  uint16(sym_multi_line_comment),
	7049:  uint16(sym__raw_string),
	7051:  uint16(anon_sym_SLASH_DASH),
	7052:  uint16(anon_sym_RBRACE),
	7053:  uint16(sym__normal_bare_identifier),
	7054:  uint16(anon_sym_LPAREN),
	7055:  uint16(anon_sym_DQUOTE),
	7056:  uint16(anon_sym_PLUS),
	7057:  uint16(anon_sym_DASH),
	7058:  uint16(aux_sym__newline_token1),
	7059:  uint16(aux_sym__newline_token2),
	7060:  uint16(aux_sym__newline_token3),
	7061:  uint16(aux_sym__newline_token4),
	7062:  uint16(aux_sym__newline_token5),
	7063:  uint16(aux_sym__newline_token6),
	7064:  uint16(aux_sym__newline_token7),
	7065:  uint16(sym__bom),
	7066:  uint16(sym__unicode_space),
	7067:  uint16(anon_sym_SLASH_SLASH),
	7068:  uint16(4),
	7069:  uint16(415),
	7070:  uint16(1),
	7071:  uint16(anon_sym_DOT),
	7072:  uint16(202),
	7073:  uint16(1),
	7074:  uint16(sym__exponent),
	7075:  uint16(417),
	7076:  uint16(2),
	7077:  uint16(anon_sym_e),
	7078:  uint16(anon_sym_E),
	7079:  uint16(413),
	7080:  uint16(16),
	7081:  uint16(sym__eof),
	7082:  uint16(sym_multi_line_comment),
	7083:  uint16(anon_sym_SLASH_DASH),
	7084:  uint16(anon_sym_LBRACE),
	7085:  uint16(anon_sym_SEMI),
	7086:  uint16(anon_sym_BSLASH),
	7087:  uint16(aux_sym__newline_token1),
	7088:  uint16(aux_sym__newline_token2),
	7089:  uint16(aux_sym__newline_token3),
	7090:  uint16(aux_sym__newline_token4),
	7091:  uint16(aux_sym__newline_token5),
	7092:  uint16(aux_sym__newline_token6),
	7093:  uint16(aux_sym__newline_token7),
	7094:  uint16(sym__bom),
	7095:  uint16(sym__unicode_space),
	7096:  uint16(anon_sym_SLASH_SLASH),
	7097:  uint16(1),
	7098:  uint16(419),
	7099:  uint16(20),
	7100:  uint16(sym_multi_line_comment),
	7101:  uint16(sym__raw_string),
	7103:  uint16(anon_sym_SLASH_DASH),
	7104:  uint16(anon_sym_RBRACE),
	7105:  uint16(sym__normal_bare_identifier),
	7106:  uint16(anon_sym_LPAREN),
	7107:  uint16(anon_sym_DQUOTE),
	7108:  uint16(anon_sym_PLUS),
	7109:  uint16(anon_sym_DASH),
	7110:  uint16(aux_sym__newline_token1),
	7111:  uint16(aux_sym__newline_token2),
	7112:  uint16(aux_sym__newline_token3),
	7113:  uint16(aux_sym__newline_token4),
	7114:  uint16(aux_sym__newline_token5),
	7115:  uint16(aux_sym__newline_token6),
	7116:  uint16(aux_sym__newline_token7),
	7117:  uint16(sym__bom),
	7118:  uint16(sym__unicode_space),
	7119:  uint16(anon_sym_SLASH_SLASH),
	7120:  uint16(1),
	7121:  uint16(421),
	7122:  uint16(20),
	7123:  uint16(sym_multi_line_comment),
	7124:  uint16(sym__raw_string),
	7126:  uint16(anon_sym_SLASH_DASH),
	7127:  uint16(anon_sym_RBRACE),
	7128:  uint16(sym__normal_bare_identifier),
	7129:  uint16(anon_sym_LPAREN),
	7130:  uint16(anon_sym_DQUOTE),
	7131:  uint16(anon_sym_PLUS),
	7132:  uint16(anon_sym_DASH),
	7133:  uint16(aux_sym__newline_token1),
	7134:  uint16(aux_sym__newline_token2),
	7135:  uint16(aux_sym__newline_token3),
	7136:  uint16(aux_sym__newline_token4),
	7137:  uint16(aux_sym__newline_token5),
	7138:  uint16(aux_sym__newline_token6),
	7139:  uint16(aux_sym__newline_token7),
	7140:  uint16(sym__bom),
	7141:  uint16(sym__unicode_space),
	7142:  uint16(anon_sym_SLASH_SLASH),
	7143:  uint16(1),
	7144:  uint16(423),
	7145:  uint16(20),
	7146:  uint16(sym_multi_line_comment),
	7147:  uint16(sym__raw_string),
	7149:  uint16(anon_sym_SLASH_DASH),
	7150:  uint16(anon_sym_RBRACE),
	7151:  uint16(sym__normal_bare_identifier),
	7152:  uint16(anon_sym_LPAREN),
	7153:  uint16(anon_sym_DQUOTE),
	7154:  uint16(anon_sym_PLUS),
	7155:  uint16(anon_sym_DASH),
	7156:  uint16(aux_sym__newline_token1),
	7157:  uint16(aux_sym__newline_token2),
	7158:  uint16(aux_sym__newline_token3),
	7159:  uint16(aux_sym__newline_token4),
	7160:  uint16(aux_sym__newline_token5),
	7161:  uint16(aux_sym__newline_token6),
	7162:  uint16(aux_sym__newline_token7),
	7163:  uint16(sym__bom),
	7164:  uint16(sym__unicode_space),
	7165:  uint16(anon_sym_SLASH_SLASH),
	7166:  uint16(1),
	7167:  uint16(425),
	7168:  uint16(20),
	7169:  uint16(sym_multi_line_comment),
	7170:  uint16(sym__raw_string),
	7172:  uint16(anon_sym_SLASH_DASH),
	7173:  uint16(anon_sym_RBRACE),
	7174:  uint16(sym__normal_bare_identifier),
	7175:  uint16(anon_sym_LPAREN),
	7176:  uint16(anon_sym_DQUOTE),
	7177:  uint16(anon_sym_PLUS),
	7178:  uint16(anon_sym_DASH),
	7179:  uint16(aux_sym__newline_token1),
	7180:  uint16(aux_sym__newline_token2),
	7181:  uint16(aux_sym__newline_token3),
	7182:  uint16(aux_sym__newline_token4),
	7183:  uint16(aux_sym__newline_token5),
	7184:  uint16(aux_sym__newline_token6),
	7185:  uint16(aux_sym__newline_token7),
	7186:  uint16(sym__bom),
	7187:  uint16(sym__unicode_space),
	7188:  uint16(anon_sym_SLASH_SLASH),
	7189:  uint16(1),
	7190:  uint16(427),
	7191:  uint16(20),
	7192:  uint16(sym_multi_line_comment),
	7193:  uint16(sym__raw_string),
	7195:  uint16(anon_sym_SLASH_DASH),
	7196:  uint16(anon_sym_RBRACE),
	7197:  uint16(sym__normal_bare_identifier),
	7198:  uint16(anon_sym_LPAREN),
	7199:  uint16(anon_sym_DQUOTE),
	7200:  uint16(anon_sym_PLUS),
	7201:  uint16(anon_sym_DASH),
	7202:  uint16(aux_sym__newline_token1),
	7203:  uint16(aux_sym__newline_token2),
	7204:  uint16(aux_sym__newline_token3),
	7205:  uint16(aux_sym__newline_token4),
	7206:  uint16(aux_sym__newline_token5),
	7207:  uint16(aux_sym__newline_token6),
	7208:  uint16(aux_sym__newline_token7),
	7209:  uint16(sym__bom),
	7210:  uint16(sym__unicode_space),
	7211:  uint16(anon_sym_SLASH_SLASH),
	7212:  uint16(1),
	7213:  uint16(429),
	7214:  uint16(20),
	7215:  uint16(sym_multi_line_comment),
	7216:  uint16(sym__raw_string),
	7218:  uint16(anon_sym_SLASH_DASH),
	7219:  uint16(anon_sym_RBRACE),
	7220:  uint16(sym__normal_bare_identifier),
	7221:  uint16(anon_sym_LPAREN),
	7222:  uint16(anon_sym_DQUOTE),
	7223:  uint16(anon_sym_PLUS),
	7224:  uint16(anon_sym_DASH),
	7225:  uint16(aux_sym__newline_token1),
	7226:  uint16(aux_sym__newline_token2),
	7227:  uint16(aux_sym__newline_token3),
	7228:  uint16(aux_sym__newline_token4),
	7229:  uint16(aux_sym__newline_token5),
	7230:  uint16(aux_sym__newline_token6),
	7231:  uint16(aux_sym__newline_token7),
	7232:  uint16(sym__bom),
	7233:  uint16(sym__unicode_space),
	7234:  uint16(anon_sym_SLASH_SLASH),
	7235:  uint16(1),
	7236:  uint16(431),
	7237:  uint16(20),
	7238:  uint16(sym_multi_line_comment),
	7239:  uint16(sym__raw_string),
	7241:  uint16(anon_sym_SLASH_DASH),
	7242:  uint16(anon_sym_RBRACE),
	7243:  uint16(sym__normal_bare_identifier),
	7244:  uint16(anon_sym_LPAREN),
	7245:  uint16(anon_sym_DQUOTE),
	7246:  uint16(anon_sym_PLUS),
	7247:  uint16(anon_sym_DASH),
	7248:  uint16(aux_sym__newline_token1),
	7249:  uint16(aux_sym__newline_token2),
	7250:  uint16(aux_sym__newline_token3),
	7251:  uint16(aux_sym__newline_token4),
	7252:  uint16(aux_sym__newline_token5),
	7253:  uint16(aux_sym__newline_token6),
	7254:  uint16(aux_sym__newline_token7),
	7255:  uint16(sym__bom),
	7256:  uint16(sym__unicode_space),
	7257:  uint16(anon_sym_SLASH_SLASH),
	7258:  uint16(1),
	7259:  uint16(433),
	7260:  uint16(20),
	7261:  uint16(sym_multi_line_comment),
	7262:  uint16(sym__raw_string),
	7264:  uint16(anon_sym_SLASH_DASH),
	7265:  uint16(anon_sym_RBRACE),
	7266:  uint16(sym__normal_bare_identifier),
	7267:  uint16(anon_sym_LPAREN),
	7268:  uint16(anon_sym_DQUOTE),
	7269:  uint16(anon_sym_PLUS),
	7270:  uint16(anon_sym_DASH),
	7271:  uint16(aux_sym__newline_token1),
	7272:  uint16(aux_sym__newline_token2),
	7273:  uint16(aux_sym__newline_token3),
	7274:  uint16(aux_sym__newline_token4),
	7275:  uint16(aux_sym__newline_token5),
	7276:  uint16(aux_sym__newline_token6),
	7277:  uint16(aux_sym__newline_token7),
	7278:  uint16(sym__bom),
	7279:  uint16(sym__unicode_space),
	7280:  uint16(anon_sym_SLASH_SLASH),
	7281:  uint16(1),
	7282:  uint16(435),
	7283:  uint16(20),
	7284:  uint16(sym_multi_line_comment),
	7285:  uint16(sym__raw_string),
	7287:  uint16(anon_sym_SLASH_DASH),
	7288:  uint16(anon_sym_RBRACE),
	7289:  uint16(sym__normal_bare_identifier),
	7290:  uint16(anon_sym_LPAREN),
	7291:  uint16(anon_sym_DQUOTE),
	7292:  uint16(anon_sym_PLUS),
	7293:  uint16(anon_sym_DASH),
	7294:  uint16(aux_sym__newline_token1),
	7295:  uint16(aux_sym__newline_token2),
	7296:  uint16(aux_sym__newline_token3),
	7297:  uint16(aux_sym__newline_token4),
	7298:  uint16(aux_sym__newline_token5),
	7299:  uint16(aux_sym__newline_token6),
	7300:  uint16(aux_sym__newline_token7),
	7301:  uint16(sym__bom),
	7302:  uint16(sym__unicode_space),
	7303:  uint16(anon_sym_SLASH_SLASH),
	7304:  uint16(1),
	7305:  uint16(437),
	7306:  uint16(20),
	7307:  uint16(sym_multi_line_comment),
	7308:  uint16(sym__raw_string),
	7310:  uint16(anon_sym_SLASH_DASH),
	7311:  uint16(anon_sym_RBRACE),
	7312:  uint16(sym__normal_bare_identifier),
	7313:  uint16(anon_sym_LPAREN),
	7314:  uint16(anon_sym_DQUOTE),
	7315:  uint16(anon_sym_PLUS),
	7316:  uint16(anon_sym_DASH),
	7317:  uint16(aux_sym__newline_token1),
	7318:  uint16(aux_sym__newline_token2),
	7319:  uint16(aux_sym__newline_token3),
	7320:  uint16(aux_sym__newline_token4),
	7321:  uint16(aux_sym__newline_token5),
	7322:  uint16(aux_sym__newline_token6),
	7323:  uint16(aux_sym__newline_token7),
	7324:  uint16(sym__bom),
	7325:  uint16(sym__unicode_space),
	7326:  uint16(anon_sym_SLASH_SLASH),
	7327:  uint16(1),
	7328:  uint16(439),
	7329:  uint16(20),
	7330:  uint16(sym_multi_line_comment),
	7331:  uint16(sym__raw_string),
	7333:  uint16(anon_sym_SLASH_DASH),
	7334:  uint16(anon_sym_RBRACE),
	7335:  uint16(sym__normal_bare_identifier),
	7336:  uint16(anon_sym_LPAREN),
	7337:  uint16(anon_sym_DQUOTE),
	7338:  uint16(anon_sym_PLUS),
	7339:  uint16(anon_sym_DASH),
	7340:  uint16(aux_sym__newline_token1),
	7341:  uint16(aux_sym__newline_token2),
	7342:  uint16(aux_sym__newline_token3),
	7343:  uint16(aux_sym__newline_token4),
	7344:  uint16(aux_sym__newline_token5),
	7345:  uint16(aux_sym__newline_token6),
	7346:  uint16(aux_sym__newline_token7),
	7347:  uint16(sym__bom),
	7348:  uint16(sym__unicode_space),
	7349:  uint16(anon_sym_SLASH_SLASH),
	7350:  uint16(3),
	7351:  uint16(164),
	7352:  uint16(1),
	7353:  uint16(aux_sym__binary_repeat1),
	7354:  uint16(443),
	7355:  uint16(3),
	7356:  uint16(anon_sym__),
	7357:  uint16(anon_sym_0),
	7358:  uint16(anon_sym_1),
	7359:  uint16(441),
	7360:  uint16(16),
	7361:  uint16(sym__eof),
	7362:  uint16(sym_multi_line_comment),
	7363:  uint16(anon_sym_SLASH_DASH),
	7364:  uint16(anon_sym_LBRACE),
	7365:  uint16(anon_sym_SEMI),
	7366:  uint16(anon_sym_BSLASH),
	7367:  uint16(aux_sym__newline_token1),
	7368:  uint16(aux_sym__newline_token2),
	7369:  uint16(aux_sym__newline_token3),
	7370:  uint16(aux_sym__newline_token4),
	7371:  uint16(aux_sym__newline_token5),
	7372:  uint16(aux_sym__newline_token6),
	7373:  uint16(aux_sym__newline_token7),
	7374:  uint16(sym__bom),
	7375:  uint16(sym__unicode_space),
	7376:  uint16(anon_sym_SLASH_SLASH),
	7377:  uint16(1),
	7378:  uint16(445),
	7379:  uint16(20),
	7380:  uint16(sym_multi_line_comment),
	7381:  uint16(sym__raw_string),
	7383:  uint16(anon_sym_SLASH_DASH),
	7384:  uint16(anon_sym_RBRACE),
	7385:  uint16(sym__normal_bare_identifier),
	7386:  uint16(anon_sym_LPAREN),
	7387:  uint16(anon_sym_DQUOTE),
	7388:  uint16(anon_sym_PLUS),
	7389:  uint16(anon_sym_DASH),
	7390:  uint16(aux_sym__newline_token1),
	7391:  uint16(aux_sym__newline_token2),
	7392:  uint16(aux_sym__newline_token3),
	7393:  uint16(aux_sym__newline_token4),
	7394:  uint16(aux_sym__newline_token5),
	7395:  uint16(aux_sym__newline_token6),
	7396:  uint16(aux_sym__newline_token7),
	7397:  uint16(sym__bom),
	7398:  uint16(sym__unicode_space),
	7399:  uint16(anon_sym_SLASH_SLASH),
	7400:  uint16(1),
	7401:  uint16(447),
	7402:  uint16(20),
	7403:  uint16(sym_multi_line_comment),
	7404:  uint16(sym__raw_string),
	7406:  uint16(anon_sym_SLASH_DASH),
	7407:  uint16(anon_sym_RBRACE),
	7408:  uint16(sym__normal_bare_identifier),
	7409:  uint16(anon_sym_LPAREN),
	7410:  uint16(anon_sym_DQUOTE),
	7411:  uint16(anon_sym_PLUS),
	7412:  uint16(anon_sym_DASH),
	7413:  uint16(aux_sym__newline_token1),
	7414:  uint16(aux_sym__newline_token2),
	7415:  uint16(aux_sym__newline_token3),
	7416:  uint16(aux_sym__newline_token4),
	7417:  uint16(aux_sym__newline_token5),
	7418:  uint16(aux_sym__newline_token6),
	7419:  uint16(aux_sym__newline_token7),
	7420:  uint16(sym__bom),
	7421:  uint16(sym__unicode_space),
	7422:  uint16(anon_sym_SLASH_SLASH),
	7423:  uint16(1),
	7424:  uint16(449),
	7425:  uint16(20),
	7426:  uint16(sym_multi_line_comment),
	7427:  uint16(sym__raw_string),
	7429:  uint16(anon_sym_SLASH_DASH),
	7430:  uint16(anon_sym_RBRACE),
	7431:  uint16(sym__normal_bare_identifier),
	7432:  uint16(anon_sym_LPAREN),
	7433:  uint16(anon_sym_DQUOTE),
	7434:  uint16(anon_sym_PLUS),
	7435:  uint16(anon_sym_DASH),
	7436:  uint16(aux_sym__newline_token1),
	7437:  uint16(aux_sym__newline_token2),
	7438:  uint16(aux_sym__newline_token3),
	7439:  uint16(aux_sym__newline_token4),
	7440:  uint16(aux_sym__newline_token5),
	7441:  uint16(aux_sym__newline_token6),
	7442:  uint16(aux_sym__newline_token7),
	7443:  uint16(sym__bom),
	7444:  uint16(sym__unicode_space),
	7445:  uint16(anon_sym_SLASH_SLASH),
	7446:  uint16(1),
	7447:  uint16(451),
	7448:  uint16(20),
	7449:  uint16(sym_multi_line_comment),
	7450:  uint16(sym__raw_string),
	7452:  uint16(anon_sym_SLASH_DASH),
	7453:  uint16(anon_sym_RBRACE),
	7454:  uint16(sym__normal_bare_identifier),
	7455:  uint16(anon_sym_LPAREN),
	7456:  uint16(anon_sym_DQUOTE),
	7457:  uint16(anon_sym_PLUS),
	7458:  uint16(anon_sym_DASH),
	7459:  uint16(aux_sym__newline_token1),
	7460:  uint16(aux_sym__newline_token2),
	7461:  uint16(aux_sym__newline_token3),
	7462:  uint16(aux_sym__newline_token4),
	7463:  uint16(aux_sym__newline_token5),
	7464:  uint16(aux_sym__newline_token6),
	7465:  uint16(aux_sym__newline_token7),
	7466:  uint16(sym__bom),
	7467:  uint16(sym__unicode_space),
	7468:  uint16(anon_sym_SLASH_SLASH),
	7469:  uint16(1),
	7470:  uint16(453),
	7471:  uint16(20),
	7472:  uint16(sym_multi_line_comment),
	7473:  uint16(sym__raw_string),
	7475:  uint16(anon_sym_SLASH_DASH),
	7476:  uint16(anon_sym_RBRACE),
	7477:  uint16(sym__normal_bare_identifier),
	7478:  uint16(anon_sym_LPAREN),
	7479:  uint16(anon_sym_DQUOTE),
	7480:  uint16(anon_sym_PLUS),
	7481:  uint16(anon_sym_DASH),
	7482:  uint16(aux_sym__newline_token1),
	7483:  uint16(aux_sym__newline_token2),
	7484:  uint16(aux_sym__newline_token3),
	7485:  uint16(aux_sym__newline_token4),
	7486:  uint16(aux_sym__newline_token5),
	7487:  uint16(aux_sym__newline_token6),
	7488:  uint16(aux_sym__newline_token7),
	7489:  uint16(sym__bom),
	7490:  uint16(sym__unicode_space),
	7491:  uint16(anon_sym_SLASH_SLASH),
	7492:  uint16(1),
	7493:  uint16(89),
	7494:  uint16(20),
	7495:  uint16(sym_multi_line_comment),
	7496:  uint16(sym__raw_string),
	7498:  uint16(anon_sym_SLASH_DASH),
	7499:  uint16(anon_sym_RBRACE),
	7500:  uint16(sym__normal_bare_identifier),
	7501:  uint16(anon_sym_LPAREN),
	7502:  uint16(anon_sym_DQUOTE),
	7503:  uint16(anon_sym_PLUS),
	7504:  uint16(anon_sym_DASH),
	7505:  uint16(aux_sym__newline_token1),
	7506:  uint16(aux_sym__newline_token2),
	7507:  uint16(aux_sym__newline_token3),
	7508:  uint16(aux_sym__newline_token4),
	7509:  uint16(aux_sym__newline_token5),
	7510:  uint16(aux_sym__newline_token6),
	7511:  uint16(aux_sym__newline_token7),
	7512:  uint16(sym__bom),
	7513:  uint16(sym__unicode_space),
	7514:  uint16(anon_sym_SLASH_SLASH),
	7515:  uint16(4),
	7516:  uint16(457),
	7517:  uint16(1),
	7518:  uint16(anon_sym_DOT),
	7519:  uint16(216),
	7520:  uint16(1),
	7521:  uint16(sym__exponent),
	7522:  uint16(417),
	7523:  uint16(2),
	7524:  uint16(anon_sym_e),
	7525:  uint16(anon_sym_E),
	7526:  uint16(455),
	7527:  uint16(16),
	7528:  uint16(sym__eof),
	7529:  uint16(sym_multi_line_comment),
	7530:  uint16(anon_sym_SLASH_DASH),
	7531:  uint16(anon_sym_LBRACE),
	7532:  uint16(anon_sym_SEMI),
	7533:  uint16(anon_sym_BSLASH),
	7534:  uint16(aux_sym__newline_token1),
	7535:  uint16(aux_sym__newline_token2),
	7536:  uint16(aux_sym__newline_token3),
	7537:  uint16(aux_sym__newline_token4),
	7538:  uint16(aux_sym__newline_token5),
	7539:  uint16(aux_sym__newline_token6),
	7540:  uint16(aux_sym__newline_token7),
	7541:  uint16(sym__bom),
	7542:  uint16(sym__unicode_space),
	7543:  uint16(anon_sym_SLASH_SLASH),
	7544:  uint16(1),
	7545:  uint16(459),
	7546:  uint16(20),
	7547:  uint16(sym_multi_line_comment),
	7548:  uint16(sym__raw_string),
	7550:  uint16(anon_sym_SLASH_DASH),
	7551:  uint16(anon_sym_RBRACE),
	7552:  uint16(sym__normal_bare_identifier),
	7553:  uint16(anon_sym_LPAREN),
	7554:  uint16(anon_sym_DQUOTE),
	7555:  uint16(anon_sym_PLUS),
	7556:  uint16(anon_sym_DASH),
	7557:  uint16(aux_sym__newline_token1),
	7558:  uint16(aux_sym__newline_token2),
	7559:  uint16(aux_sym__newline_token3),
	7560:  uint16(aux_sym__newline_token4),
	7561:  uint16(aux_sym__newline_token5),
	7562:  uint16(aux_sym__newline_token6),
	7563:  uint16(aux_sym__newline_token7),
	7564:  uint16(sym__bom),
	7565:  uint16(sym__unicode_space),
	7566:  uint16(anon_sym_SLASH_SLASH),
	7567:  uint16(1),
	7568:  uint16(461),
	7569:  uint16(20),
	7570:  uint16(sym_multi_line_comment),
	7571:  uint16(sym__raw_string),
	7573:  uint16(anon_sym_SLASH_DASH),
	7574:  uint16(anon_sym_RBRACE),
	7575:  uint16(sym__normal_bare_identifier),
	7576:  uint16(anon_sym_LPAREN),
	7577:  uint16(anon_sym_DQUOTE),
	7578:  uint16(anon_sym_PLUS),
	7579:  uint16(anon_sym_DASH),
	7580:  uint16(aux_sym__newline_token1),
	7581:  uint16(aux_sym__newline_token2),
	7582:  uint16(aux_sym__newline_token3),
	7583:  uint16(aux_sym__newline_token4),
	7584:  uint16(aux_sym__newline_token5),
	7585:  uint16(aux_sym__newline_token6),
	7586:  uint16(aux_sym__newline_token7),
	7587:  uint16(sym__bom),
	7588:  uint16(sym__unicode_space),
	7589:  uint16(anon_sym_SLASH_SLASH),
	7590:  uint16(1),
	7591:  uint16(463),
	7592:  uint16(20),
	7593:  uint16(sym_multi_line_comment),
	7594:  uint16(sym__raw_string),
	7596:  uint16(anon_sym_SLASH_DASH),
	7597:  uint16(anon_sym_RBRACE),
	7598:  uint16(sym__normal_bare_identifier),
	7599:  uint16(anon_sym_LPAREN),
	7600:  uint16(anon_sym_DQUOTE),
	7601:  uint16(anon_sym_PLUS),
	7602:  uint16(anon_sym_DASH),
	7603:  uint16(aux_sym__newline_token1),
	7604:  uint16(aux_sym__newline_token2),
	7605:  uint16(aux_sym__newline_token3),
	7606:  uint16(aux_sym__newline_token4),
	7607:  uint16(aux_sym__newline_token5),
	7608:  uint16(aux_sym__newline_token6),
	7609:  uint16(aux_sym__newline_token7),
	7610:  uint16(sym__bom),
	7611:  uint16(sym__unicode_space),
	7612:  uint16(anon_sym_SLASH_SLASH),
	7613:  uint16(1),
	7614:  uint16(465),
	7615:  uint16(20),
	7616:  uint16(sym_multi_line_comment),
	7617:  uint16(sym__raw_string),
	7619:  uint16(anon_sym_SLASH_DASH),
	7620:  uint16(anon_sym_RBRACE),
	7621:  uint16(sym__normal_bare_identifier),
	7622:  uint16(anon_sym_LPAREN),
	7623:  uint16(anon_sym_DQUOTE),
	7624:  uint16(anon_sym_PLUS),
	7625:  uint16(anon_sym_DASH),
	7626:  uint16(aux_sym__newline_token1),
	7627:  uint16(aux_sym__newline_token2),
	7628:  uint16(aux_sym__newline_token3),
	7629:  uint16(aux_sym__newline_token4),
	7630:  uint16(aux_sym__newline_token5),
	7631:  uint16(aux_sym__newline_token6),
	7632:  uint16(aux_sym__newline_token7),
	7633:  uint16(sym__bom),
	7634:  uint16(sym__unicode_space),
	7635:  uint16(anon_sym_SLASH_SLASH),
	7636:  uint16(1),
	7637:  uint16(467),
	7638:  uint16(20),
	7639:  uint16(sym_multi_line_comment),
	7640:  uint16(sym__raw_string),
	7642:  uint16(anon_sym_SLASH_DASH),
	7643:  uint16(anon_sym_RBRACE),
	7644:  uint16(sym__normal_bare_identifier),
	7645:  uint16(anon_sym_LPAREN),
	7646:  uint16(anon_sym_DQUOTE),
	7647:  uint16(anon_sym_PLUS),
	7648:  uint16(anon_sym_DASH),
	7649:  uint16(aux_sym__newline_token1),
	7650:  uint16(aux_sym__newline_token2),
	7651:  uint16(aux_sym__newline_token3),
	7652:  uint16(aux_sym__newline_token4),
	7653:  uint16(aux_sym__newline_token5),
	7654:  uint16(aux_sym__newline_token6),
	7655:  uint16(aux_sym__newline_token7),
	7656:  uint16(sym__bom),
	7657:  uint16(sym__unicode_space),
	7658:  uint16(anon_sym_SLASH_SLASH),
	7659:  uint16(1),
	7660:  uint16(469),
	7661:  uint16(20),
	7662:  uint16(sym_multi_line_comment),
	7663:  uint16(sym__raw_string),
	7665:  uint16(anon_sym_SLASH_DASH),
	7666:  uint16(anon_sym_RBRACE),
	7667:  uint16(sym__normal_bare_identifier),
	7668:  uint16(anon_sym_LPAREN),
	7669:  uint16(anon_sym_DQUOTE),
	7670:  uint16(anon_sym_PLUS),
	7671:  uint16(anon_sym_DASH),
	7672:  uint16(aux_sym__newline_token1),
	7673:  uint16(aux_sym__newline_token2),
	7674:  uint16(aux_sym__newline_token3),
	7675:  uint16(aux_sym__newline_token4),
	7676:  uint16(aux_sym__newline_token5),
	7677:  uint16(aux_sym__newline_token6),
	7678:  uint16(aux_sym__newline_token7),
	7679:  uint16(sym__bom),
	7680:  uint16(sym__unicode_space),
	7681:  uint16(anon_sym_SLASH_SLASH),
	7682:  uint16(1),
	7683:  uint16(471),
	7684:  uint16(20),
	7685:  uint16(sym_multi_line_comment),
	7686:  uint16(sym__raw_string),
	7688:  uint16(anon_sym_SLASH_DASH),
	7689:  uint16(anon_sym_RBRACE),
	7690:  uint16(sym__normal_bare_identifier),
	7691:  uint16(anon_sym_LPAREN),
	7692:  uint16(anon_sym_DQUOTE),
	7693:  uint16(anon_sym_PLUS),
	7694:  uint16(anon_sym_DASH),
	7695:  uint16(aux_sym__newline_token1),
	7696:  uint16(aux_sym__newline_token2),
	7697:  uint16(aux_sym__newline_token3),
	7698:  uint16(aux_sym__newline_token4),
	7699:  uint16(aux_sym__newline_token5),
	7700:  uint16(aux_sym__newline_token6),
	7701:  uint16(aux_sym__newline_token7),
	7702:  uint16(sym__bom),
	7703:  uint16(sym__unicode_space),
	7704:  uint16(anon_sym_SLASH_SLASH),
	7705:  uint16(1),
	7706:  uint16(473),
	7707:  uint16(20),
	7708:  uint16(sym_multi_line_comment),
	7709:  uint16(sym__raw_string),
	7711:  uint16(anon_sym_SLASH_DASH),
	7712:  uint16(anon_sym_RBRACE),
	7713:  uint16(sym__normal_bare_identifier),
	7714:  uint16(anon_sym_LPAREN),
	7715:  uint16(anon_sym_DQUOTE),
	7716:  uint16(anon_sym_PLUS),
	7717:  uint16(anon_sym_DASH),
	7718:  uint16(aux_sym__newline_token1),
	7719:  uint16(aux_sym__newline_token2),
	7720:  uint16(aux_sym__newline_token3),
	7721:  uint16(aux_sym__newline_token4),
	7722:  uint16(aux_sym__newline_token5),
	7723:  uint16(aux_sym__newline_token6),
	7724:  uint16(aux_sym__newline_token7),
	7725:  uint16(sym__bom),
	7726:  uint16(sym__unicode_space),
	7727:  uint16(anon_sym_SLASH_SLASH),
	7728:  uint16(1),
	7729:  uint16(475),
	7730:  uint16(20),
	7731:  uint16(sym_multi_line_comment),
	7732:  uint16(sym__raw_string),
	7734:  uint16(anon_sym_SLASH_DASH),
	7735:  uint16(anon_sym_RBRACE),
	7736:  uint16(sym__normal_bare_identifier),
	7737:  uint16(anon_sym_LPAREN),
	7738:  uint16(anon_sym_DQUOTE),
	7739:  uint16(anon_sym_PLUS),
	7740:  uint16(anon_sym_DASH),
	7741:  uint16(aux_sym__newline_token1),
	7742:  uint16(aux_sym__newline_token2),
	7743:  uint16(aux_sym__newline_token3),
	7744:  uint16(aux_sym__newline_token4),
	7745:  uint16(aux_sym__newline_token5),
	7746:  uint16(aux_sym__newline_token6),
	7747:  uint16(aux_sym__newline_token7),
	7748:  uint16(sym__bom),
	7749:  uint16(sym__unicode_space),
	7750:  uint16(anon_sym_SLASH_SLASH),
	7751:  uint16(3),
	7752:  uint16(130),
	7753:  uint16(1),
	7754:  uint16(aux_sym__binary_repeat1),
	7755:  uint16(479),
	7756:  uint16(3),
	7757:  uint16(anon_sym__),
	7758:  uint16(anon_sym_0),
	7759:  uint16(anon_sym_1),
	7760:  uint16(477),
	7761:  uint16(16),
	7762:  uint16(sym__eof),
	7763:  uint16(sym_multi_line_comment),
	7764:  uint16(anon_sym_SLASH_DASH),
	7765:  uint16(anon_sym_LBRACE),
	7766:  uint16(anon_sym_SEMI),
	7767:  uint16(anon_sym_BSLASH),
	7768:  uint16(aux_sym__newline_token1),
	7769:  uint16(aux_sym__newline_token2),
	7770:  uint16(aux_sym__newline_token3),
	7771:  uint16(aux_sym__newline_token4),
	7772:  uint16(aux_sym__newline_token5),
	7773:  uint16(aux_sym__newline_token6),
	7774:  uint16(aux_sym__newline_token7),
	7775:  uint16(sym__bom),
	7776:  uint16(sym__unicode_space),
	7777:  uint16(anon_sym_SLASH_SLASH),
	7778:  uint16(3),
	7779:  uint16(130),
	7780:  uint16(1),
	7781:  uint16(aux_sym__binary_repeat1),
	7782:  uint16(479),
	7783:  uint16(3),
	7784:  uint16(anon_sym__),
	7785:  uint16(anon_sym_0),
	7786:  uint16(anon_sym_1),
	7787:  uint16(481),
	7788:  uint16(16),
	7789:  uint16(sym__eof),
	7790:  uint16(sym_multi_line_comment),
	7791:  uint16(anon_sym_SLASH_DASH),
	7792:  uint16(anon_sym_LBRACE),
	7793:  uint16(anon_sym_SEMI),
	7794:  uint16(anon_sym_BSLASH),
	7795:  uint16(aux_sym__newline_token1),
	7796:  uint16(aux_sym__newline_token2),
	7797:  uint16(aux_sym__newline_token3),
	7798:  uint16(aux_sym__newline_token4),
	7799:  uint16(aux_sym__newline_token5),
	7800:  uint16(aux_sym__newline_token6),
	7801:  uint16(aux_sym__newline_token7),
	7802:  uint16(sym__bom),
	7803:  uint16(sym__unicode_space),
	7804:  uint16(anon_sym_SLASH_SLASH),
	7805:  uint16(3),
	7806:  uint16(165),
	7807:  uint16(1),
	7808:  uint16(aux_sym__binary_repeat1),
	7809:  uint16(483),
	7810:  uint16(3),
	7811:  uint16(anon_sym__),
	7812:  uint16(anon_sym_0),
	7813:  uint16(anon_sym_1),
	7814:  uint16(477),
	7815:  uint16(16),
	7816:  uint16(sym__eof),
	7817:  uint16(sym_multi_line_comment),
	7818:  uint16(anon_sym_SLASH_DASH),
	7819:  uint16(anon_sym_LBRACE),
	7820:  uint16(anon_sym_SEMI),
	7821:  uint16(anon_sym_BSLASH),
	7822:  uint16(aux_sym__newline_token1),
	7823:  uint16(aux_sym__newline_token2),
	7824:  uint16(aux_sym__newline_token3),
	7825:  uint16(aux_sym__newline_token4),
	7826:  uint16(aux_sym__newline_token5),
	7827:  uint16(aux_sym__newline_token6),
	7828:  uint16(aux_sym__newline_token7),
	7829:  uint16(sym__bom),
	7830:  uint16(sym__unicode_space),
	7831:  uint16(anon_sym_SLASH_SLASH),
	7832:  uint16(1),
	7833:  uint16(485),
	7834:  uint16(20),
	7835:  uint16(sym_multi_line_comment),
	7836:  uint16(sym__raw_string),
	7838:  uint16(anon_sym_SLASH_DASH),
	7839:  uint16(anon_sym_RBRACE),
	7840:  uint16(sym__normal_bare_identifier),
	7841:  uint16(anon_sym_LPAREN),
	7842:  uint16(anon_sym_DQUOTE),
	7843:  uint16(anon_sym_PLUS),
	7844:  uint16(anon_sym_DASH),
	7845:  uint16(aux_sym__newline_token1),
	7846:  uint16(aux_sym__newline_token2),
	7847:  uint16(aux_sym__newline_token3),
	7848:  uint16(aux_sym__newline_token4),
	7849:  uint16(aux_sym__newline_token5),
	7850:  uint16(aux_sym__newline_token6),
	7851:  uint16(aux_sym__newline_token7),
	7852:  uint16(sym__bom),
	7853:  uint16(sym__unicode_space),
	7854:  uint16(anon_sym_SLASH_SLASH),
	7855:  uint16(1),
	7856:  uint16(487),
	7857:  uint16(20),
	7858:  uint16(sym_multi_line_comment),
	7859:  uint16(sym__raw_string),
	7861:  uint16(anon_sym_SLASH_DASH),
	7862:  uint16(anon_sym_RBRACE),
	7863:  uint16(sym__normal_bare_identifier),
	7864:  uint16(anon_sym_LPAREN),
	7865:  uint16(anon_sym_DQUOTE),
	7866:  uint16(anon_sym_PLUS),
	7867:  uint16(anon_sym_DASH),
	7868:  uint16(aux_sym__newline_token1),
	7869:  uint16(aux_sym__newline_token2),
	7870:  uint16(aux_sym__newline_token3),
	7871:  uint16(aux_sym__newline_token4),
	7872:  uint16(aux_sym__newline_token5),
	7873:  uint16(aux_sym__newline_token6),
	7874:  uint16(aux_sym__newline_token7),
	7875:  uint16(sym__bom),
	7876:  uint16(sym__unicode_space),
	7877:  uint16(anon_sym_SLASH_SLASH),
	7878:  uint16(1),
	7879:  uint16(489),
	7880:  uint16(20),
	7881:  uint16(sym_multi_line_comment),
	7882:  uint16(sym__raw_string),
	7884:  uint16(anon_sym_SLASH_DASH),
	7885:  uint16(anon_sym_RBRACE),
	7886:  uint16(sym__normal_bare_identifier),
	7887:  uint16(anon_sym_LPAREN),
	7888:  uint16(anon_sym_DQUOTE),
	7889:  uint16(anon_sym_PLUS),
	7890:  uint16(anon_sym_DASH),
	7891:  uint16(aux_sym__newline_token1),
	7892:  uint16(aux_sym__newline_token2),
	7893:  uint16(aux_sym__newline_token3),
	7894:  uint16(aux_sym__newline_token4),
	7895:  uint16(aux_sym__newline_token5),
	7896:  uint16(aux_sym__newline_token6),
	7897:  uint16(aux_sym__newline_token7),
	7898:  uint16(sym__bom),
	7899:  uint16(sym__unicode_space),
	7900:  uint16(anon_sym_SLASH_SLASH),
	7901:  uint16(1),
	7902:  uint16(491),
	7903:  uint16(20),
	7904:  uint16(sym_multi_line_comment),
	7905:  uint16(sym__raw_string),
	7907:  uint16(anon_sym_SLASH_DASH),
	7908:  uint16(anon_sym_RBRACE),
	7909:  uint16(sym__normal_bare_identifier),
	7910:  uint16(anon_sym_LPAREN),
	7911:  uint16(anon_sym_DQUOTE),
	7912:  uint16(anon_sym_PLUS),
	7913:  uint16(anon_sym_DASH),
	7914:  uint16(aux_sym__newline_token1),
	7915:  uint16(aux_sym__newline_token2),
	7916:  uint16(aux_sym__newline_token3),
	7917:  uint16(aux_sym__newline_token4),
	7918:  uint16(aux_sym__newline_token5),
	7919:  uint16(aux_sym__newline_token6),
	7920:  uint16(aux_sym__newline_token7),
	7921:  uint16(sym__bom),
	7922:  uint16(sym__unicode_space),
	7923:  uint16(anon_sym_SLASH_SLASH),
	7924:  uint16(1),
	7925:  uint16(493),
	7926:  uint16(20),
	7927:  uint16(sym_multi_line_comment),
	7928:  uint16(sym__raw_string),
	7930:  uint16(anon_sym_SLASH_DASH),
	7931:  uint16(anon_sym_RBRACE),
	7932:  uint16(sym__normal_bare_identifier),
	7933:  uint16(anon_sym_LPAREN),
	7934:  uint16(anon_sym_DQUOTE),
	7935:  uint16(anon_sym_PLUS),
	7936:  uint16(anon_sym_DASH),
	7937:  uint16(aux_sym__newline_token1),
	7938:  uint16(aux_sym__newline_token2),
	7939:  uint16(aux_sym__newline_token3),
	7940:  uint16(aux_sym__newline_token4),
	7941:  uint16(aux_sym__newline_token5),
	7942:  uint16(aux_sym__newline_token6),
	7943:  uint16(aux_sym__newline_token7),
	7944:  uint16(sym__bom),
	7945:  uint16(sym__unicode_space),
	7946:  uint16(anon_sym_SLASH_SLASH),
	7947:  uint16(3),
	7948:  uint16(181),
	7949:  uint16(1),
	7950:  uint16(aux_sym__octal_repeat1),
	7951:  uint16(497),
	7952:  uint16(2),
	7953:  uint16(anon_sym__),
	7954:  uint16(aux_sym__octal_token1),
	7955:  uint16(495),
	7956:  uint16(16),
	7957:  uint16(sym__eof),
	7958:  uint16(sym_multi_line_comment),
	7959:  uint16(anon_sym_SLASH_DASH),
	7960:  uint16(anon_sym_LBRACE),
	7961:  uint16(anon_sym_SEMI),
	7962:  uint16(anon_sym_BSLASH),
	7963:  uint16(aux_sym__newline_token1),
	7964:  uint16(aux_sym__newline_token2),
	7965:  uint16(aux_sym__newline_token3),
	7966:  uint16(aux_sym__newline_token4),
	7967:  uint16(aux_sym__newline_token5),
	7968:  uint16(aux_sym__newline_token6),
	7969:  uint16(aux_sym__newline_token7),
	7970:  uint16(sym__bom),
	7971:  uint16(sym__unicode_space),
	7972:  uint16(anon_sym_SLASH_SLASH),
	7973:  uint16(3),
	7974:  uint16(182),
	7975:  uint16(1),
	7976:  uint16(aux_sym__octal_repeat1),
	7977:  uint16(501),
	7978:  uint16(2),
	7979:  uint16(anon_sym__),
	7980:  uint16(aux_sym__octal_token1),
	7981:  uint16(499),
	7982:  uint16(16),
	7983:  uint16(sym__eof),
	7984:  uint16(sym_multi_line_comment),
	7985:  uint16(anon_sym_SLASH_DASH),
	7986:  uint16(anon_sym_LBRACE),
	7987:  uint16(anon_sym_SEMI),
	7988:  uint16(anon_sym_BSLASH),
	7989:  uint16(aux_sym__newline_token1),
	7990:  uint16(aux_sym__newline_token2),
	7991:  uint16(aux_sym__newline_token3),
	7992:  uint16(aux_sym__newline_token4),
	7993:  uint16(aux_sym__newline_token5),
	7994:  uint16(aux_sym__newline_token6),
	7995:  uint16(aux_sym__newline_token7),
	7996:  uint16(sym__bom),
	7997:  uint16(sym__unicode_space),
	7998:  uint16(anon_sym_SLASH_SLASH),
	7999:  uint16(3),
	8000:  uint16(203),
	8001:  uint16(1),
	8002:  uint16(sym__exponent),
	8003:  uint16(417),
	8004:  uint16(2),
	8005:  uint16(anon_sym_e),
	8006:  uint16(anon_sym_E),
	8007:  uint16(503),
	8008:  uint16(16),
	8009:  uint16(sym__eof),
	8010:  uint16(sym_multi_line_comment),
	8011:  uint16(anon_sym_SLASH_DASH),
	8012:  uint16(anon_sym_LBRACE),
	8013:  uint16(anon_sym_SEMI),
	8014:  uint16(anon_sym_BSLASH),
	8015:  uint16(aux_sym__newline_token1),
	8016:  uint16(aux_sym__newline_token2),
	8017:  uint16(aux_sym__newline_token3),
	8018:  uint16(aux_sym__newline_token4),
	8019:  uint16(aux_sym__newline_token5),
	8020:  uint16(aux_sym__newline_token6),
	8021:  uint16(aux_sym__newline_token7),
	8022:  uint16(sym__bom),
	8023:  uint16(sym__unicode_space),
	8024:  uint16(anon_sym_SLASH_SLASH),
	8025:  uint16(3),
	8026:  uint16(178),
	8027:  uint16(1),
	8028:  uint16(aux_sym__hex_repeat1),
	8029:  uint16(507),
	8030:  uint16(2),
	8031:  uint16(sym__hex_digit),
	8032:  uint16(anon_sym__),
	8033:  uint16(505),
	8034:  uint16(16),
	8035:  uint16(sym__eof),
	8036:  uint16(sym_multi_line_comment),
	8037:  uint16(anon_sym_SLASH_DASH),
	8038:  uint16(anon_sym_LBRACE),
	8039:  uint16(anon_sym_SEMI),
	8040:  uint16(anon_sym_BSLASH),
	8041:  uint16(aux_sym__newline_token1),
	8042:  uint16(aux_sym__newline_token2),
	8043:  uint16(aux_sym__newline_token3),
	8044:  uint16(aux_sym__newline_token4),
	8045:  uint16(aux_sym__newline_token5),
	8046:  uint16(aux_sym__newline_token6),
	8047:  uint16(aux_sym__newline_token7),
	8048:  uint16(sym__bom),
	8049:  uint16(sym__unicode_space),
	8050:  uint16(anon_sym_SLASH_SLASH),
	8051:  uint16(3),
	8052:  uint16(172),
	8053:  uint16(1),
	8054:  uint16(aux_sym__octal_repeat1),
	8055:  uint16(511),
	8056:  uint16(2),
	8057:  uint16(anon_sym__),
	8058:  uint16(aux_sym__octal_token1),
	8059:  uint16(509),
	8060:  uint16(16),
	8061:  uint16(sym__eof),
	8062:  uint16(sym_multi_line_comment),
	8063:  uint16(anon_sym_SLASH_DASH),
	8064:  uint16(anon_sym_LBRACE),
	8065:  uint16(anon_sym_SEMI),
	8066:  uint16(anon_sym_BSLASH),
	8067:  uint16(aux_sym__newline_token1),
	8068:  uint16(aux_sym__newline_token2),
	8069:  uint16(aux_sym__newline_token3),
	8070:  uint16(aux_sym__newline_token4),
	8071:  uint16(aux_sym__newline_token5),
	8072:  uint16(aux_sym__newline_token6),
	8073:  uint16(aux_sym__newline_token7),
	8074:  uint16(sym__bom),
	8075:  uint16(sym__unicode_space),
	8076:  uint16(anon_sym_SLASH_SLASH),
	8077:  uint16(3),
	8078:  uint16(184),
	8079:  uint16(1),
	8080:  uint16(aux_sym__hex_repeat1),
	8081:  uint16(513),
	8082:  uint16(2),
	8083:  uint16(sym__hex_digit),
	8084:  uint16(anon_sym__),
	8085:  uint16(505),
	8086:  uint16(16),
	8087:  uint16(sym__eof),
	8088:  uint16(sym_multi_line_comment),
	8089:  uint16(anon_sym_SLASH_DASH),
	8090:  uint16(anon_sym_LBRACE),
	8091:  uint16(anon_sym_SEMI),
	8092:  uint16(anon_sym_BSLASH),
	8093:  uint16(aux_sym__newline_token1),
	8094:  uint16(aux_sym__newline_token2),
	8095:  uint16(aux_sym__newline_token3),
	8096:  uint16(aux_sym__newline_token4),
	8097:  uint16(aux_sym__newline_token5),
	8098:  uint16(aux_sym__newline_token6),
	8099:  uint16(aux_sym__newline_token7),
	8100:  uint16(sym__bom),
	8101:  uint16(sym__unicode_space),
	8102:  uint16(anon_sym_SLASH_SLASH),
	8103:  uint16(3),
	8104:  uint16(184),
	8105:  uint16(1),
	8106:  uint16(aux_sym__hex_repeat1),
	8107:  uint16(513),
	8108:  uint16(2),
	8109:  uint16(sym__hex_digit),
	8110:  uint16(anon_sym__),
	8111:  uint16(515),
	8112:  uint16(16),
	8113:  uint16(sym__eof),
	8114:  uint16(sym_multi_line_comment),
	8115:  uint16(anon_sym_SLASH_DASH),
	8116:  uint16(anon_sym_LBRACE),
	8117:  uint16(anon_sym_SEMI),
	8118:  uint16(anon_sym_BSLASH),
	8119:  uint16(aux_sym__newline_token1),
	8120:  uint16(aux_sym__newline_token2),
	8121:  uint16(aux_sym__newline_token3),
	8122:  uint16(aux_sym__newline_token4),
	8123:  uint16(aux_sym__newline_token5),
	8124:  uint16(aux_sym__newline_token6),
	8125:  uint16(aux_sym__newline_token7),
	8126:  uint16(sym__bom),
	8127:  uint16(sym__unicode_space),
	8128:  uint16(anon_sym_SLASH_SLASH),
	8129:  uint16(2),
	8130:  uint16(226),
	8131:  uint16(5),
	8132:  uint16(sym__normal_bare_identifier),
	8133:  uint16(anon_sym_null),
	8134:  uint16(sym__digit),
	8135:  uint16(anon_sym_true),
	8136:  uint16(anon_sym_false),
	8137:  uint16(228),
	8138:  uint16(14),
	8139:  uint16(sym_multi_line_comment),
	8140:  uint16(sym__raw_string),
	8141:  uint16(anon_sym_SLASH_DASH),
	8142:  uint16(anon_sym_LBRACE),
	8143:  uint16(anon_sym_LPAREN),
	8144:  uint16(anon_sym_DQUOTE),
	8145:  uint16(anon_sym_PLUS),
	8146:  uint16(anon_sym_DASH),
	8147:  uint16(anon_sym_0x),
	8148:  uint16(anon_sym_0o),
	8149:  uint16(anon_sym_0b),
	8150:  uint16(anon_sym_BSLASH),
	8151:  uint16(sym__bom),
	8152:  uint16(sym__unicode_space),
	8153:  uint16(3),
	8154:  uint16(215),
	8155:  uint16(1),
	8156:  uint16(sym__exponent),
	8157:  uint16(417),
	8158:  uint16(2),
	8159:  uint16(anon_sym_e),
	8160:  uint16(anon_sym_E),
	8161:  uint16(517),
	8162:  uint16(16),
	8163:  uint16(sym__eof),
	8164:  uint16(sym_multi_line_comment),
	8165:  uint16(anon_sym_SLASH_DASH),
	8166:  uint16(anon_sym_LBRACE),
	8167:  uint16(anon_sym_SEMI),
	8168:  uint16(anon_sym_BSLASH),
	8169:  uint16(aux_sym__newline_token1),
	8170:  uint16(aux_sym__newline_token2),
	8171:  uint16(aux_sym__newline_token3),
	8172:  uint16(aux_sym__newline_token4),
	8173:  uint16(aux_sym__newline_token5),
	8174:  uint16(aux_sym__newline_token6),
	8175:  uint16(aux_sym__newline_token7),
	8176:  uint16(sym__bom),
	8177:  uint16(sym__unicode_space),
	8178:  uint16(anon_sym_SLASH_SLASH),
	8179:  uint16(3),
	8180:  uint16(181),
	8181:  uint16(1),
	8182:  uint16(aux_sym__octal_repeat1),
	8183:  uint16(521),
	8184:  uint16(2),
	8185:  uint16(anon_sym__),
	8186:  uint16(aux_sym__octal_token1),
	8187:  uint16(519),
	8188:  uint16(16),
	8189:  uint16(sym__eof),
	8190:  uint16(sym_multi_line_comment),
	8191:  uint16(anon_sym_SLASH_DASH),
	8192:  uint16(anon_sym_LBRACE),
	8193:  uint16(anon_sym_SEMI),
	8194:  uint16(anon_sym_BSLASH),
	8195:  uint16(aux_sym__newline_token1),
	8196:  uint16(aux_sym__newline_token2),
	8197:  uint16(aux_sym__newline_token3),
	8198:  uint16(aux_sym__newline_token4),
	8199:  uint16(aux_sym__newline_token5),
	8200:  uint16(aux_sym__newline_token6),
	8201:  uint16(aux_sym__newline_token7),
	8202:  uint16(sym__bom),
	8203:  uint16(sym__unicode_space),
	8204:  uint16(anon_sym_SLASH_SLASH),
	8205:  uint16(3),
	8206:  uint16(181),
	8207:  uint16(1),
	8208:  uint16(aux_sym__octal_repeat1),
	8209:  uint16(497),
	8210:  uint16(2),
	8211:  uint16(anon_sym__),
	8212:  uint16(aux_sym__octal_token1),
	8213:  uint16(509),
	8214:  uint16(16),
	8215:  uint16(sym__eof),
	8216:  uint16(sym_multi_line_comment),
	8217:  uint16(anon_sym_SLASH_DASH),
	8218:  uint16(anon_sym_LBRACE),
	8219:  uint16(anon_sym_SEMI),
	8220:  uint16(anon_sym_BSLASH),
	8221:  uint16(aux_sym__newline_token1),
	8222:  uint16(aux_sym__newline_token2),
	8223:  uint16(aux_sym__newline_token3),
	8224:  uint16(aux_sym__newline_token4),
	8225:  uint16(aux_sym__newline_token5),
	8226:  uint16(aux_sym__newline_token6),
	8227:  uint16(aux_sym__newline_token7),
	8228:  uint16(sym__bom),
	8229:  uint16(sym__unicode_space),
	8230:  uint16(anon_sym_SLASH_SLASH),
	8231:  uint16(2),
	8232:  uint16(222),
	8233:  uint16(5),
	8234:  uint16(sym__normal_bare_identifier),
	8235:  uint16(anon_sym_null),
	8236:  uint16(sym__digit),
	8237:  uint16(anon_sym_true),
	8238:  uint16(anon_sym_false),
	8239:  uint16(224),
	8240:  uint16(14),
	8241:  uint16(sym_multi_line_comment),
	8242:  uint16(sym__raw_string),
	8243:  uint16(anon_sym_SLASH_DASH),
	8244:  uint16(anon_sym_LBRACE),
	8245:  uint16(anon_sym_LPAREN),
	8246:  uint16(anon_sym_DQUOTE),
	8247:  uint16(anon_sym_PLUS),
	8248:  uint16(anon_sym_DASH),
	8249:  uint16(anon_sym_0x),
	8250:  uint16(anon_sym_0o),
	8251:  uint16(anon_sym_0b),
	8252:  uint16(anon_sym_BSLASH),
	8253:  uint16(sym__bom),
	8254:  uint16(sym__unicode_space),
	8255:  uint16(3),
	8256:  uint16(184),
	8257:  uint16(1),
	8258:  uint16(aux_sym__hex_repeat1),
	8259:  uint16(526),
	8260:  uint16(2),
	8261:  uint16(sym__hex_digit),
	8262:  uint16(anon_sym__),
	8263:  uint16(524),
	8264:  uint16(16),
	8265:  uint16(sym__eof),
	8266:  uint16(sym_multi_line_comment),
	8267:  uint16(anon_sym_SLASH_DASH),
	8268:  uint16(anon_sym_LBRACE),
	8269:  uint16(anon_sym_SEMI),
	8270:  uint16(anon_sym_BSLASH),
	8271:  uint16(aux_sym__newline_token1),
	8272:  uint16(aux_sym__newline_token2),
	8273:  uint16(aux_sym__newline_token3),
	8274:  uint16(aux_sym__newline_token4),
	8275:  uint16(aux_sym__newline_token5),
	8276:  uint16(aux_sym__newline_token6),
	8277:  uint16(aux_sym__newline_token7),
	8278:  uint16(sym__bom),
	8279:  uint16(sym__unicode_space),
	8280:  uint16(anon_sym_SLASH_SLASH),
	8281:  uint16(3),
	8282:  uint16(177),
	8283:  uint16(1),
	8284:  uint16(aux_sym__hex_repeat1),
	8285:  uint16(531),
	8286:  uint16(2),
	8287:  uint16(sym__hex_digit),
	8288:  uint16(anon_sym__),
	8289:  uint16(529),
	8290:  uint16(16),
	8291:  uint16(sym__eof),
	8292:  uint16(sym_multi_line_comment),
	8293:  uint16(anon_sym_SLASH_DASH),
	8294:  uint16(anon_sym_LBRACE),
	8295:  uint16(anon_sym_SEMI),
	8296:  uint16(anon_sym_BSLASH),
	8297:  uint16(aux_sym__newline_token1),
	8298:  uint16(aux_sym__newline_token2),
	8299:  uint16(aux_sym__newline_token3),
	8300:  uint16(aux_sym__newline_token4),
	8301:  uint16(aux_sym__newline_token5),
	8302:  uint16(aux_sym__newline_token6),
	8303:  uint16(aux_sym__newline_token7),
	8304:  uint16(sym__bom),
	8305:  uint16(sym__unicode_space),
	8306:  uint16(anon_sym_SLASH_SLASH),
	8307:  uint16(2),
	8308:  uint16(230),
	8309:  uint16(5),
	8310:  uint16(sym__normal_bare_identifier),
	8311:  uint16(anon_sym_null),
	8312:  uint16(sym__digit),
	8313:  uint16(anon_sym_true),
	8314:  uint16(anon_sym_false),
	8315:  uint16(232),
	8316:  uint16(14),
	8317:  uint16(sym_multi_line_comment),
	8318:  uint16(sym__raw_string),
	8319:  uint16(anon_sym_SLASH_DASH),
	8320:  uint16(anon_sym_LBRACE),
	8321:  uint16(anon_sym_LPAREN),
	8322:  uint16(anon_sym_DQUOTE),
	8323:  uint16(anon_sym_PLUS),
	8324:  uint16(anon_sym_DASH),
	8325:  uint16(anon_sym_0x),
	8326:  uint16(anon_sym_0o),
	8327:  uint16(anon_sym_0b),
	8328:  uint16(anon_sym_BSLASH),
	8329:  uint16(sym__bom),
	8330:  uint16(sym__unicode_space),
	8331:  uint16(7),
	8332:  uint16(533),
	8333:  uint16(1),
	8334:  uint16(anon_sym_BSLASH),
	8335:  uint16(187),
	8336:  uint16(1),
	8337:  uint16(aux_sym_node_repeat1),
	8338:  uint16(212),
	8339:  uint16(1),
	8340:  uint16(sym__escline),
	8341:  uint16(234),
	8342:  uint16(1),
	8343:  uint16(sym__node_space),
	8344:  uint16(193),
	8345:  uint16(2),
	8346:  uint16(sym__ws),
	8347:  uint16(aux_sym_node_repeat3),
	8348:  uint16(536),
	8349:  uint16(3),
	8350:  uint16(sym_multi_line_comment),
	8351:  uint16(sym__bom),
	8352:  uint16(sym__unicode_space),
	8353:  uint16(81),
	8354:  uint16(10),
	8355:  uint16(sym__eof),
	8356:  uint16(anon_sym_SEMI),
	8357:  uint16(aux_sym__newline_token1),
	8358:  uint16(aux_sym__newline_token2),
	8359:  uint16(aux_sym__newline_token3),
	8360:  uint16(aux_sym__newline_token4),
	8361:  uint16(aux_sym__newline_token5),
	8362:  uint16(aux_sym__newline_token6),
	8363:  uint16(aux_sym__newline_token7),
	8364:  uint16(anon_sym_SLASH_SLASH),
	8365:  uint16(3),
	8366:  uint16(541),
	8367:  uint16(1),
	8368:  uint16(sym__identifier_char),
	8369:  uint16(188),
	8370:  uint16(1),
	8371:  uint16(aux_sym__bare_identifier_repeat1),
	8372:  uint16(539),
	8373:  uint16(16),
	8374:  uint16(sym__eof),
	8375:  uint16(sym_multi_line_comment),
	8376:  uint16(anon_sym_SLASH_DASH),
	8377:  uint16(anon_sym_LBRACE),
	8378:  uint16(anon_sym_SEMI),
	8379:  uint16(anon_sym_BSLASH),
	8380:  uint16(aux_sym__newline_token1),
	8381:  uint16(aux_sym__newline_token2),
	8382:  uint16(aux_sym__newline_token3),
	8383:  uint16(aux_sym__newline_token4),
	8384:  uint16(aux_sym__newline_token5),
	8385:  uint16(aux_sym__newline_token6),
	8386:  uint16(aux_sym__newline_token7),
	8387:  uint16(sym__bom),
	8388:  uint16(sym__unicode_space),
	8389:  uint16(anon_sym_SLASH_SLASH),
	8390:  uint16(3),
	8391:  uint16(546),
	8392:  uint16(1),
	8393:  uint16(sym__identifier_char),
	8394:  uint16(188),
	8395:  uint16(1),
	8396:  uint16(aux_sym__bare_identifier_repeat1),
	8397:  uint16(544),
	8398:  uint16(16),
	8399:  uint16(sym__eof),
	8400:  uint16(sym_multi_line_comment),
	8401:  uint16(anon_sym_SLASH_DASH),
	8402:  uint16(anon_sym_LBRACE),
	8403:  uint16(anon_sym_SEMI),
	8404:  uint16(anon_sym_BSLASH),
	8405:  uint16(aux_sym__newline_token1),
	8406:  uint16(aux_sym__newline_token2),
	8407:  uint16(aux_sym__newline_token3),
	8408:  uint16(aux_sym__newline_token4),
	8409:  uint16(aux_sym__newline_token5),
	8410:  uint16(aux_sym__newline_token6),
	8411:  uint16(aux_sym__newline_token7),
	8412:  uint16(sym__bom),
	8413:  uint16(sym__unicode_space),
	8414:  uint16(anon_sym_SLASH_SLASH),
	8415:  uint16(3),
	8416:  uint16(550),
	8417:  uint16(1),
	8418:  uint16(sym__identifier_char),
	8419:  uint16(189),
	8420:  uint16(1),
	8421:  uint16(aux_sym__bare_identifier_repeat1),
	8422:  uint16(548),
	8423:  uint16(16),
	8424:  uint16(sym__eof),
	8425:  uint16(sym_multi_line_comment),
	8426:  uint16(anon_sym_SLASH_DASH),
	8427:  uint16(anon_sym_LBRACE),
	8428:  uint16(anon_sym_SEMI),
	8429:  uint16(anon_sym_BSLASH),
	8430:  uint16(aux_sym__newline_token1),
	8431:  uint16(aux_sym__newline_token2),
	8432:  uint16(aux_sym__newline_token3),
	8433:  uint16(aux_sym__newline_token4),
	8434:  uint16(aux_sym__newline_token5),
	8435:  uint16(aux_sym__newline_token6),
	8436:  uint16(aux_sym__newline_token7),
	8437:  uint16(sym__bom),
	8438:  uint16(sym__unicode_space),
	8439:  uint16(anon_sym_SLASH_SLASH),
	8440:  uint16(2),
	8441:  uint16(554),
	8442:  uint16(1),
	8443:  uint16(anon_sym_EQ),
	8444:  uint16(552),
	8445:  uint16(16),
	8446:  uint16(sym__eof),
	8447:  uint16(sym_multi_line_comment),
	8448:  uint16(anon_sym_SLASH_DASH),
	8449:  uint16(anon_sym_LBRACE),
	8450:  uint16(anon_sym_SEMI),
	8451:  uint16(anon_sym_BSLASH),
	8452:  uint16(aux_sym__newline_token1),
	8453:  uint16(aux_sym__newline_token2),
	8454:  uint16(aux_sym__newline_token3),
	8455:  uint16(aux_sym__newline_token4),
	8456:  uint16(aux_sym__newline_token5),
	8457:  uint16(aux_sym__newline_token6),
	8458:  uint16(aux_sym__newline_token7),
	8459:  uint16(sym__bom),
	8460:  uint16(sym__unicode_space),
	8461:  uint16(anon_sym_SLASH_SLASH),
	8462:  uint16(2),
	8463:  uint16(558),
	8464:  uint16(1),
	8465:  uint16(sym___identifier_char_no_digit),
	8466:  uint16(556),
	8467:  uint16(16),
	8468:  uint16(sym__eof),
	8469:  uint16(sym_multi_line_comment),
	8470:  uint16(anon_sym_SLASH_DASH),
	8471:  uint16(anon_sym_LBRACE),
	8472:  uint16(anon_sym_SEMI),
	8473:  uint16(anon_sym_BSLASH),
	8474:  uint16(aux_sym__newline_token1),
	8475:  uint16(aux_sym__newline_token2),
	8476:  uint16(aux_sym__newline_token3),
	8477:  uint16(aux_sym__newline_token4),
	8478:  uint16(aux_sym__newline_token5),
	8479:  uint16(aux_sym__newline_token6),
	8480:  uint16(aux_sym__newline_token7),
	8481:  uint16(sym__bom),
	8482:  uint16(sym__unicode_space),
	8483:  uint16(anon_sym_SLASH_SLASH),
	8484:  uint16(5),
	8485:  uint16(560),
	8486:  uint16(1),
	8487:  uint16(anon_sym_BSLASH),
	8488:  uint16(206),
	8489:  uint16(1),
	8490:  uint16(sym__escline),
	8491:  uint16(205),
	8492:  uint16(2),
	8493:  uint16(sym__ws),
	8494:  uint16(aux_sym_node_repeat3),
	8495:  uint16(563),
	8496:  uint16(3),
	8497:  uint16(sym_multi_line_comment),
	8498:  uint16(sym__bom),
	8499:  uint16(sym__unicode_space),
	8500:  uint16(149),
	8501:  uint16(10),
	8502:  uint16(sym__eof),
	8503:  uint16(anon_sym_SEMI),
	8504:  uint16(aux_sym__newline_token1),
	8505:  uint16(aux_sym__newline_token2),
	8506:  uint16(aux_sym__newline_token3),
	8507:  uint16(aux_sym__newline_token4),
	8508:  uint16(aux_sym__newline_token5),
	8509:  uint16(aux_sym__newline_token6),
	8510:  uint16(aux_sym__newline_token7),
	8511:  uint16(anon_sym_SLASH_SLASH),
	8512:  uint16(1),
	8513:  uint16(566),
	8514:  uint16(17),
	8515:  uint16(sym__eof),
	8516:  uint16(sym_multi_line_comment),
	8517:  uint16(anon_sym_SLASH_DASH),
	8518:  uint16(anon_sym_LBRACE),
	8519:  uint16(anon_sym_SEMI),
	8520:  uint16(anon_sym_EQ),
	8521:  uint16(anon_sym_BSLASH),
	8522:  uint16(aux_sym__newline_token1),
	8523:  uint16(aux_sym__newline_token2),
	8524:  uint16(aux_sym__newline_token3),
	8525:  uint16(aux_sym__newline_token4),
	8526:  uint16(aux_sym__newline_token5),
	8527:  uint16(aux_sym__newline_token6),
	8528:  uint16(aux_sym__newline_token7),
	8529:  uint16(sym__bom),
	8530:  uint16(sym__unicode_space),
	8531:  uint16(anon_sym_SLASH_SLASH),
	8532:  uint16(5),
	8533:  uint16(571),
	8534:  uint16(1),
	8535:  uint16(anon_sym_BSLASH),
	8536:  uint16(206),
	8537:  uint16(1),
	8538:  uint16(sym__escline),
	8539:  uint16(205),
	8540:  uint16(2),
	8541:  uint16(sym__ws),
	8542:  uint16(aux_sym_node_repeat3),
	8543:  uint16(575),
	8544:  uint16(3),
	8545:  uint16(sym_multi_line_comment),
	8546:  uint16(sym__bom),
	8547:  uint16(sym__unicode_space),
	8548:  uint16(568),
	8549:  uint16(10),
	8550:  uint16(sym__eof),
	8551:  uint16(anon_sym_SEMI),
	8552:  uint16(aux_sym__newline_token1),
	8553:  uint16(aux_sym__newline_token2),
	8554:  uint16(aux_sym__newline_token3),
	8555:  uint16(aux_sym__newline_token4),
	8556:  uint16(aux_sym__newline_token5),
	8557:  uint16(aux_sym__newline_token6),
	8558:  uint16(aux_sym__newline_token7),
	8559:  uint16(anon_sym_SLASH_SLASH),
	8560:  uint16(1),
	8561:  uint16(579),
	8562:  uint16(17),
	8563:  uint16(sym__eof),
	8564:  uint16(sym_multi_line_comment),
	8565:  uint16(anon_sym_SLASH_DASH),
	8566:  uint16(anon_sym_LBRACE),
	8567:  uint16(anon_sym_SEMI),
	8568:  uint16(anon_sym_EQ),
	8569:  uint16(anon_sym_BSLASH),
	8570:  uint16(aux_sym__newline_token1),
	8571:  uint16(aux_sym__newline_token2),
	8572:  uint16(aux_sym__newline_token3),
	8573:  uint16(aux_sym__newline_token4),
	8574:  uint16(aux_sym__newline_token5),
	8575:  uint16(aux_sym__newline_token6),
	8576:  uint16(aux_sym__newline_token7),
	8577:  uint16(sym__bom),
	8578:  uint16(sym__unicode_space),
	8579:  uint16(anon_sym_SLASH_SLASH),
	8580:  uint16(1),
	8581:  uint16(581),
	8582:  uint16(17),
	8583:  uint16(sym__eof),
	8584:  uint16(sym_multi_line_comment),
	8585:  uint16(anon_sym_SLASH_DASH),
	8586:  uint16(anon_sym_LBRACE),
	8587:  uint16(anon_sym_SEMI),
	8588:  uint16(anon_sym_EQ),
	8589:  uint16(anon_sym_BSLASH),
	8590:  uint16(aux_sym__newline_token1),
	8591:  uint16(aux_sym__newline_token2),
	8592:  uint16(aux_sym__newline_token3),
	8593:  uint16(aux_sym__newline_token4),
	8594:  uint16(aux_sym__newline_token5),
	8595:  uint16(aux_sym__newline_token6),
	8596:  uint16(aux_sym__newline_token7),
	8597:  uint16(sym__bom),
	8598:  uint16(sym__unicode_space),
	8599:  uint16(anon_sym_SLASH_SLASH),
	8600:  uint16(1),
	8601:  uint16(304),
	8602:  uint16(16),
	8603:  uint16(sym__eof),
	8604:  uint16(sym_multi_line_comment),
	8605:  uint16(anon_sym_SLASH_DASH),
	8606:  uint16(anon_sym_LBRACE),
	8607:  uint16(anon_sym_SEMI),
	8608:  uint16(anon_sym_BSLASH),
	8609:  uint16(aux_sym__newline_token1),
	8610:  uint16(aux_sym__newline_token2),
	8611:  uint16(aux_sym__newline_token3),
	8612:  uint16(aux_sym__newline_token4),
	8613:  uint16(aux_sym__newline_token5),
	8614:  uint16(aux_sym__newline_token6),
	8615:  uint16(aux_sym__newline_token7),
	8616:  uint16(sym__bom),
	8617:  uint16(sym__unicode_space),
	8618:  uint16(anon_sym_SLASH_SLASH),
	8619:  uint16(1),
	8620:  uint16(552),
	8621:  uint16(16),
	8622:  uint16(sym__eof),
	8623:  uint16(sym_multi_line_comment),
	8624:  uint16(anon_sym_SLASH_DASH),
	8625:  uint16(anon_sym_LBRACE),
	8626:  uint16(anon_sym_SEMI),
	8627:  uint16(anon_sym_BSLASH),
	8628:  uint16(aux_sym__newline_token1),
	8629:  uint16(aux_sym__newline_token2),
	8630:  uint16(aux_sym__newline_token3),
	8631:  uint16(aux_sym__newline_token4),
	8632:  uint16(aux_sym__newline_token5),
	8633:  uint16(aux_sym__newline_token6),
	8634:  uint16(aux_sym__newline_token7),
	8635:  uint16(sym__bom),
	8636:  uint16(sym__unicode_space),
	8637:  uint16(anon_sym_SLASH_SLASH),
	8638:  uint16(1),
	8639:  uint16(583),
	8640:  uint16(16),
	8641:  uint16(sym__eof),
	8642:  uint16(sym_multi_line_comment),
	8643:  uint16(anon_sym_SLASH_DASH),
	8644:  uint16(anon_sym_LBRACE),
	8645:  uint16(anon_sym_SEMI),
	8646:  uint16(anon_sym_BSLASH),
	8647:  uint16(aux_sym__newline_token1),
	8648:  uint16(aux_sym__newline_token2),
	8649:  uint16(aux_sym__newline_token3),
	8650:  uint16(aux_sym__newline_token4),
	8651:  uint16(aux_sym__newline_token5),
	8652:  uint16(aux_sym__newline_token6),
	8653:  uint16(aux_sym__newline_token7),
	8654:  uint16(sym__bom),
	8655:  uint16(sym__unicode_space),
	8656:  uint16(anon_sym_SLASH_SLASH),
	8657:  uint16(3),
	8658:  uint16(205),
	8659:  uint16(2),
	8660:  uint16(sym__ws),
	8661:  uint16(aux_sym_node_repeat3),
	8662:  uint16(585),
	8663:  uint16(3),
	8664:  uint16(sym_multi_line_comment),
	8665:  uint16(sym__bom),
	8666:  uint16(sym__unicode_space),
	8667:  uint16(180),
	8668:  uint16(11),
	8669:  uint16(sym__eof),
	8670:  uint16(anon_sym_SEMI),
	8671:  uint16(anon_sym_BSLASH),
	8672:  uint16(aux_sym__newline_token1),
	8673:  uint16(aux_sym__newline_token2),
	8674:  uint16(aux_sym__newline_token3),
	8675:  uint16(aux_sym__newline_token4),
	8676:  uint16(aux_sym__newline_token5),
	8677:  uint16(aux_sym__newline_token6),
	8678:  uint16(aux_sym__newline_token7),
	8679:  uint16(anon_sym_SLASH_SLASH),
	8680:  uint16(1),
	8681:  uint16(588),
	8682:  uint16(16),
	8683:  uint16(sym__eof),
	8684:  uint16(sym_multi_line_comment),
	8685:  uint16(anon_sym_SLASH_DASH),
	8686:  uint16(anon_sym_LBRACE),
	8687:  uint16(anon_sym_SEMI),
	8688:  uint16(anon_sym_BSLASH),
	8689:  uint16(aux_sym__newline_token1),
	8690:  uint16(aux_sym__newline_token2),
	8691:  uint16(aux_sym__newline_token3),
	8692:  uint16(aux_sym__newline_token4),
	8693:  uint16(aux_sym__newline_token5),
	8694:  uint16(aux_sym__newline_token6),
	8695:  uint16(aux_sym__newline_token7),
	8696:  uint16(sym__bom),
	8697:  uint16(sym__unicode_space),
	8698:  uint16(anon_sym_SLASH_SLASH),
	8699:  uint16(1),
	8700:  uint16(590),
	8701:  uint16(16),
	8702:  uint16(sym__eof),
	8703:  uint16(sym_multi_line_comment),
	8704:  uint16(anon_sym_SLASH_DASH),
	8705:  uint16(anon_sym_LBRACE),
	8706:  uint16(anon_sym_SEMI),
	8707:  uint16(anon_sym_BSLASH),
	8708:  uint16(aux_sym__newline_token1),
	8709:  uint16(aux_sym__newline_token2),
	8710:  uint16(aux_sym__newline_token3),
	8711:  uint16(aux_sym__newline_token4),
	8712:  uint16(aux_sym__newline_token5),
	8713:  uint16(aux_sym__newline_token6),
	8714:  uint16(aux_sym__newline_token7),
	8715:  uint16(sym__bom),
	8716:  uint16(sym__unicode_space),
	8717:  uint16(anon_sym_SLASH_SLASH),
	8718:  uint16(1),
	8719:  uint16(592),
	8720:  uint16(16),
	8721:  uint16(sym__eof),
	8722:  uint16(sym_multi_line_comment),
	8723:  uint16(anon_sym_SLASH_DASH),
	8724:  uint16(anon_sym_LBRACE),
	8725:  uint16(anon_sym_SEMI),
	8726:  uint16(anon_sym_BSLASH),
	8727:  uint16(aux_sym__newline_token1),
	8728:  uint16(aux_sym__newline_token2),
	8729:  uint16(aux_sym__newline_token3),
	8730:  uint16(aux_sym__newline_token4),
	8731:  uint16(aux_sym__newline_token5),
	8732:  uint16(aux_sym__newline_token6),
	8733:  uint16(aux_sym__newline_token7),
	8734:  uint16(sym__bom),
	8735:  uint16(sym__unicode_space),
	8736:  uint16(anon_sym_SLASH_SLASH),
	8737:  uint16(3),
	8738:  uint16(205),
	8739:  uint16(2),
	8740:  uint16(sym__ws),
	8741:  uint16(aux_sym_node_repeat3),
	8742:  uint16(594),
	8743:  uint16(3),
	8744:  uint16(sym_multi_line_comment),
	8745:  uint16(sym__bom),
	8746:  uint16(sym__unicode_space),
	8747:  uint16(208),
	8748:  uint16(11),
	8749:  uint16(sym__eof),
	8750:  uint16(anon_sym_SEMI),
	8751:  uint16(anon_sym_BSLASH),
	8752:  uint16(aux_sym__newline_token1),
	8753:  uint16(aux_sym__newline_token2),
	8754:  uint16(aux_sym__newline_token3),
	8755:  uint16(aux_sym__newline_token4),
	8756:  uint16(aux_sym__newline_token5),
	8757:  uint16(aux_sym__newline_token6),
	8758:  uint16(aux_sym__newline_token7),
	8759:  uint16(anon_sym_SLASH_SLASH),
	8760:  uint16(3),
	8761:  uint16(201),
	8762:  uint16(2),
	8763:  uint16(sym__ws),
	8764:  uint16(aux_sym_node_repeat3),
	8765:  uint16(597),
	8766:  uint16(3),
	8767:  uint16(sym_multi_line_comment),
	8768:  uint16(sym__bom),
	8769:  uint16(sym__unicode_space),
	8770:  uint16(187),
	8771:  uint16(11),
	8772:  uint16(sym__eof),
	8773:  uint16(anon_sym_SEMI),
	8774:  uint16(anon_sym_BSLASH),
	8775:  uint16(aux_sym__newline_token1),
	8776:  uint16(aux_sym__newline_token2),
	8777:  uint16(aux_sym__newline_token3),
	8778:  uint16(aux_sym__newline_token4),
	8779:  uint16(aux_sym__newline_token5),
	8780:  uint16(aux_sym__newline_token6),
	8781:  uint16(aux_sym__newline_token7),
	8782:  uint16(anon_sym_SLASH_SLASH),
	8783:  uint16(3),
	8784:  uint16(205),
	8785:  uint16(2),
	8786:  uint16(sym__ws),
	8787:  uint16(aux_sym_node_repeat3),
	8788:  uint16(600),
	8789:  uint16(3),
	8790:  uint16(sym_multi_line_comment),
	8791:  uint16(sym__bom),
	8792:  uint16(sym__unicode_space),
	8793:  uint16(187),
	8794:  uint16(11),
	8795:  uint16(sym__eof),
	8796:  uint16(anon_sym_SEMI),
	8797:  uint16(anon_sym_BSLASH),
	8798:  uint16(aux_sym__newline_token1),
	8799:  uint16(aux_sym__newline_token2),
	8800:  uint16(aux_sym__newline_token3),
	8801:  uint16(aux_sym__newline_token4),
	8802:  uint16(aux_sym__newline_token5),
	8803:  uint16(aux_sym__newline_token6),
	8804:  uint16(aux_sym__newline_token7),
	8805:  uint16(anon_sym_SLASH_SLASH),
	8806:  uint16(1),
	8807:  uint16(603),
	8808:  uint16(16),
	8809:  uint16(sym__eof),
	8810:  uint16(sym_multi_line_comment),
	8811:  uint16(anon_sym_SLASH_DASH),
	8812:  uint16(anon_sym_LBRACE),
	8813:  uint16(anon_sym_SEMI),
	8814:  uint16(anon_sym_BSLASH),
	8815:  uint16(aux_sym__newline_token1),
	8816:  uint16(aux_sym__newline_token2),
	8817:  uint16(aux_sym__newline_token3),
	8818:  uint16(aux_sym__newline_token4),
	8819:  uint16(aux_sym__newline_token5),
	8820:  uint16(aux_sym__newline_token6),
	8821:  uint16(aux_sym__newline_token7),
	8822:  uint16(sym__bom),
	8823:  uint16(sym__unicode_space),
	8824:  uint16(anon_sym_SLASH_SLASH),
	8825:  uint16(1),
	8826:  uint16(605),
	8827:  uint16(16),
	8828:  uint16(sym__eof),
	8829:  uint16(sym_multi_line_comment),
	8830:  uint16(anon_sym_SLASH_DASH),
	8831:  uint16(anon_sym_LBRACE),
	8832:  uint16(anon_sym_SEMI),
	8833:  uint16(anon_sym_BSLASH),
	8834:  uint16(aux_sym__newline_token1),
	8835:  uint16(aux_sym__newline_token2),
	8836:  uint16(aux_sym__newline_token3),
	8837:  uint16(aux_sym__newline_token4),
	8838:  uint16(aux_sym__newline_token5),
	8839:  uint16(aux_sym__newline_token6),
	8840:  uint16(aux_sym__newline_token7),
	8841:  uint16(sym__bom),
	8842:  uint16(sym__unicode_space),
	8843:  uint16(anon_sym_SLASH_SLASH),
	8844:  uint16(3),
	8845:  uint16(205),
	8846:  uint16(2),
	8847:  uint16(sym__ws),
	8848:  uint16(aux_sym_node_repeat3),
	8849:  uint16(610),
	8850:  uint16(3),
	8851:  uint16(sym_multi_line_comment),
	8852:  uint16(sym__bom),
	8853:  uint16(sym__unicode_space),
	8854:  uint16(607),
	8855:  uint16(11),
	8856:  uint16(sym__eof),
	8857:  uint16(anon_sym_SEMI),
	8858:  uint16(anon_sym_BSLASH),
	8859:  uint16(aux_sym__newline_token1),
	8860:  uint16(aux_sym__newline_token2),
	8861:  uint16(aux_sym__newline_token3),
	8862:  uint16(aux_sym__newline_token4),
	8863:  uint16(aux_sym__newline_token5),
	8864:  uint16(aux_sym__newline_token6),
	8865:  uint16(aux_sym__newline_token7),
	8866:  uint16(anon_sym_SLASH_SLASH),
	8867:  uint16(1),
	8868:  uint16(614),
	8869:  uint16(16),
	8870:  uint16(sym__eof),
	8871:  uint16(sym_multi_line_comment),
	8872:  uint16(anon_sym_SLASH_DASH),
	8873:  uint16(anon_sym_LBRACE),
	8874:  uint16(anon_sym_SEMI),
	8875:  uint16(anon_sym_BSLASH),
	8876:  uint16(aux_sym__newline_token1),
	8877:  uint16(aux_sym__newline_token2),
	8878:  uint16(aux_sym__newline_token3),
	8879:  uint16(aux_sym__newline_token4),
	8880:  uint16(aux_sym__newline_token5),
	8881:  uint16(aux_sym__newline_token6),
	8882:  uint16(aux_sym__newline_token7),
	8883:  uint16(sym__bom),
	8884:  uint16(sym__unicode_space),
	8885:  uint16(anon_sym_SLASH_SLASH),
	8886:  uint16(3),
	8887:  uint16(207),
	8888:  uint16(2),
	8889:  uint16(sym__ws),
	8890:  uint16(aux_sym_node_repeat3),
	8891:  uint16(616),
	8892:  uint16(3),
	8893:  uint16(sym_multi_line_comment),
	8894:  uint16(sym__bom),
	8895:  uint16(sym__unicode_space),
	8896:  uint16(149),
	8897:  uint16(11),
	8898:  uint16(sym__eof),
	8899:  uint16(anon_sym_SEMI),
	8900:  uint16(anon_sym_BSLASH),
	8901:  uint16(aux_sym__newline_token1),
	8902:  uint16(aux_sym__newline_token2),
	8903:  uint16(aux_sym__newline_token3),
	8904:  uint16(aux_sym__newline_token4),
	8905:  uint16(aux_sym__newline_token5),
	8906:  uint16(aux_sym__newline_token6),
	8907:  uint16(aux_sym__newline_token7),
	8908:  uint16(anon_sym_SLASH_SLASH),
	8909:  uint16(1),
	8910:  uint16(619),
	8911:  uint16(16),
	8912:  uint16(sym__eof),
	8913:  uint16(sym_multi_line_comment),
	8914:  uint16(anon_sym_SLASH_DASH),
	8915:  uint16(anon_sym_LBRACE),
	8916:  uint16(anon_sym_SEMI),
	8917:  uint16(anon_sym_BSLASH),
	8918:  uint16(aux_sym__newline_token1),
	8919:  uint16(aux_sym__newline_token2),
	8920:  uint16(aux_sym__newline_token3),
	8921:  uint16(aux_sym__newline_token4),
	8922:  uint16(aux_sym__newline_token5),
	8923:  uint16(aux_sym__newline_token6),
	8924:  uint16(aux_sym__newline_token7),
	8925:  uint16(sym__bom),
	8926:  uint16(sym__unicode_space),
	8927:  uint16(anon_sym_SLASH_SLASH),
	8928:  uint16(1),
	8929:  uint16(621),
	8930:  uint16(16),
	8931:  uint16(sym__eof),
	8932:  uint16(sym_multi_line_comment),
	8933:  uint16(anon_sym_SLASH_DASH),
	8934:  uint16(anon_sym_LBRACE),
	8935:  uint16(anon_sym_SEMI),
	8936:  uint16(anon_sym_BSLASH),
	8937:  uint16(aux_sym__newline_token1),
	8938:  uint16(aux_sym__newline_token2),
	8939:  uint16(aux_sym__newline_token3),
	8940:  uint16(aux_sym__newline_token4),
	8941:  uint16(aux_sym__newline_token5),
	8942:  uint16(aux_sym__newline_token6),
	8943:  uint16(aux_sym__newline_token7),
	8944:  uint16(sym__bom),
	8945:  uint16(sym__unicode_space),
	8946:  uint16(anon_sym_SLASH_SLASH),
	8947:  uint16(1),
	8948:  uint16(623),
	8949:  uint16(16),
	8950:  uint16(sym__eof),
	8951:  uint16(sym_multi_line_comment),
	8952:  uint16(anon_sym_SLASH_DASH),
	8953:  uint16(anon_sym_LBRACE),
	8954:  uint16(anon_sym_SEMI),
	8955:  uint16(anon_sym_BSLASH),
	8956:  uint16(aux_sym__newline_token1),
	8957:  uint16(aux_sym__newline_token2),
	8958:  uint16(aux_sym__newline_token3),
	8959:  uint16(aux_sym__newline_token4),
	8960:  uint16(aux_sym__newline_token5),
	8961:  uint16(aux_sym__newline_token6),
	8962:  uint16(aux_sym__newline_token7),
	8963:  uint16(sym__bom),
	8964:  uint16(sym__unicode_space),
	8965:  uint16(anon_sym_SLASH_SLASH),
	8966:  uint16(1),
	8967:  uint16(413),
	8968:  uint16(16),
	8969:  uint16(sym__eof),
	8970:  uint16(sym_multi_line_comment),
	8971:  uint16(anon_sym_SLASH_DASH),
	8972:  uint16(anon_sym_LBRACE),
	8973:  uint16(anon_sym_SEMI),
	8974:  uint16(anon_sym_BSLASH),
	8975:  uint16(aux_sym__newline_token1),
	8976:  uint16(aux_sym__newline_token2),
	8977:  uint16(aux_sym__newline_token3),
	8978:  uint16(aux_sym__newline_token4),
	8979:  uint16(aux_sym__newline_token5),
	8980:  uint16(aux_sym__newline_token6),
	8981:  uint16(aux_sym__newline_token7),
	8982:  uint16(sym__bom),
	8983:  uint16(sym__unicode_space),
	8984:  uint16(anon_sym_SLASH_SLASH),
	8985:  uint16(1),
	8986:  uint16(625),
	8987:  uint16(16),
	8988:  uint16(sym__eof),
	8989:  uint16(sym_multi_line_comment),
	8990:  uint16(anon_sym_SLASH_DASH),
	8991:  uint16(anon_sym_LBRACE),
	8992:  uint16(anon_sym_SEMI),
	8993:  uint16(anon_sym_BSLASH),
	8994:  uint16(aux_sym__newline_token1),
	8995:  uint16(aux_sym__newline_token2),
	8996:  uint16(aux_sym__newline_token3),
	8997:  uint16(aux_sym__newline_token4),
	8998:  uint16(aux_sym__newline_token5),
	8999:  uint16(aux_sym__newline_token6),
	9000:  uint16(aux_sym__newline_token7),
	9001:  uint16(sym__bom),
	9002:  uint16(sym__unicode_space),
	9003:  uint16(anon_sym_SLASH_SLASH),
	9004:  uint16(1),
	9005:  uint16(554),
	9006:  uint16(16),
	9007:  uint16(sym__eof),
	9008:  uint16(sym_multi_line_comment),
	9009:  uint16(anon_sym_SLASH_DASH),
	9010:  uint16(anon_sym_LBRACE),
	9011:  uint16(anon_sym_SEMI),
	9012:  uint16(anon_sym_BSLASH),
	9013:  uint16(aux_sym__newline_token1),
	9014:  uint16(aux_sym__newline_token2),
	9015:  uint16(aux_sym__newline_token3),
	9016:  uint16(aux_sym__newline_token4),
	9017:  uint16(aux_sym__newline_token5),
	9018:  uint16(aux_sym__newline_token6),
	9019:  uint16(aux_sym__newline_token7),
	9020:  uint16(sym__bom),
	9021:  uint16(sym__unicode_space),
	9022:  uint16(anon_sym_SLASH_SLASH),
	9023:  uint16(1),
	9024:  uint16(627),
	9025:  uint16(16),
	9026:  uint16(sym__eof),
	9027:  uint16(sym_multi_line_comment),
	9028:  uint16(anon_sym_SLASH_DASH),
	9029:  uint16(anon_sym_LBRACE),
	9030:  uint16(anon_sym_SEMI),
	9031:  uint16(anon_sym_BSLASH),
	9032:  uint16(aux_sym__newline_token1),
	9033:  uint16(aux_sym__newline_token2),
	9034:  uint16(aux_sym__newline_token3),
	9035:  uint16(aux_sym__newline_token4),
	9036:  uint16(aux_sym__newline_token5),
	9037:  uint16(aux_sym__newline_token6),
	9038:  uint16(aux_sym__newline_token7),
	9039:  uint16(sym__bom),
	9040:  uint16(sym__unicode_space),
	9041:  uint16(anon_sym_SLASH_SLASH),
	9042:  uint16(1),
	9043:  uint16(629),
	9044:  uint16(16),
	9045:  uint16(sym__eof),
	9046:  uint16(sym_multi_line_comment),
	9047:  uint16(anon_sym_SLASH_DASH),
	9048:  uint16(anon_sym_LBRACE),
	9049:  uint16(anon_sym_SEMI),
	9050:  uint16(anon_sym_BSLASH),
	9051:  uint16(aux_sym__newline_token1),
	9052:  uint16(aux_sym__newline_token2),
	9053:  uint16(aux_sym__newline_token3),
	9054:  uint16(aux_sym__newline_token4),
	9055:  uint16(aux_sym__newline_token5),
	9056:  uint16(aux_sym__newline_token6),
	9057:  uint16(aux_sym__newline_token7),
	9058:  uint16(sym__bom),
	9059:  uint16(sym__unicode_space),
	9060:  uint16(anon_sym_SLASH_SLASH),
	9061:  uint16(1),
	9062:  uint16(631),
	9063:  uint16(16),
	9064:  uint16(sym__eof),
	9065:  uint16(sym_multi_line_comment),
	9066:  uint16(anon_sym_SLASH_DASH),
	9067:  uint16(anon_sym_LBRACE),
	9068:  uint16(anon_sym_SEMI),
	9069:  uint16(anon_sym_BSLASH),
	9070:  uint16(aux_sym__newline_token1),
	9071:  uint16(aux_sym__newline_token2),
	9072:  uint16(aux_sym__newline_token3),
	9073:  uint16(aux_sym__newline_token4),
	9074:  uint16(aux_sym__newline_token5),
	9075:  uint16(aux_sym__newline_token6),
	9076:  uint16(aux_sym__newline_token7),
	9077:  uint16(sym__bom),
	9078:  uint16(sym__unicode_space),
	9079:  uint16(anon_sym_SLASH_SLASH),
	9080:  uint16(3),
	9081:  uint16(210),
	9082:  uint16(2),
	9083:  uint16(sym__ws),
	9084:  uint16(aux_sym_node_repeat3),
	9085:  uint16(636),
	9086:  uint16(3),
	9087:  uint16(sym_multi_line_comment),
	9088:  uint16(sym__bom),
	9089:  uint16(sym__unicode_space),
	9090:  uint16(633),
	9091:  uint16(11),
	9092:  uint16(sym__eof),
	9093:  uint16(anon_sym_SEMI),
	9094:  uint16(anon_sym_BSLASH),
	9095:  uint16(aux_sym__newline_token1),
	9096:  uint16(aux_sym__newline_token2),
	9097:  uint16(aux_sym__newline_token3),
	9098:  uint16(aux_sym__newline_token4),
	9099:  uint16(aux_sym__newline_token5),
	9100:  uint16(aux_sym__newline_token6),
	9101:  uint16(aux_sym__newline_token7),
	9102:  uint16(anon_sym_SLASH_SLASH),
	9103:  uint16(5),
	9104:  uint16(644),
	9105:  uint16(1),
	9106:  uint16(anon_sym_SLASH_SLASH),
	9107:  uint16(230),
	9108:  uint16(2),
	9109:  uint16(sym__ws),
	9110:  uint16(aux_sym_node_repeat3),
	9111:  uint16(281),
	9112:  uint16(2),
	9113:  uint16(sym__newline),
	9114:  uint16(sym_single_line_comment),
	9115:  uint16(642),
	9116:  uint16(3),
	9117:  uint16(sym_multi_line_comment),
	9118:  uint16(sym__bom),
	9119:  uint16(sym__unicode_space),
	9120:  uint16(640),
	9121:  uint16(7),
	9122:  uint16(aux_sym__newline_token1),
	9123:  uint16(aux_sym__newline_token2),
	9124:  uint16(aux_sym__newline_token3),
	9125:  uint16(aux_sym__newline_token4),
	9126:  uint16(aux_sym__newline_token5),
	9127:  uint16(aux_sym__newline_token6),
	9128:  uint16(aux_sym__newline_token7),
	9129:  uint16(5),
	9130:  uint16(19),
	9131:  uint16(1),
	9132:  uint16(anon_sym_SLASH_SLASH),
	9133:  uint16(179),
	9134:  uint16(2),
	9135:  uint16(sym__newline),
	9136:  uint16(sym_single_line_comment),
	9137:  uint16(230),
	9138:  uint16(2),
	9139:  uint16(sym__ws),
	9140:  uint16(aux_sym_node_repeat3),
	9141:  uint16(642),
	9142:  uint16(3),
	9143:  uint16(sym_multi_line_comment),
	9144:  uint16(sym__bom),
	9145:  uint16(sym__unicode_space),
	9146:  uint16(646),
	9147:  uint16(7),
	9148:  uint16(aux_sym__newline_token1),
	9149:  uint16(aux_sym__newline_token2),
	9150:  uint16(aux_sym__newline_token3),
	9151:  uint16(aux_sym__newline_token4),
	9152:  uint16(aux_sym__newline_token5),
	9153:  uint16(aux_sym__newline_token6),
	9154:  uint16(aux_sym__newline_token7),
	9155:  uint16(5),
	9156:  uint16(650),
	9157:  uint16(1),
	9158:  uint16(anon_sym_SLASH_SLASH),
	9159:  uint16(57),
	9160:  uint16(2),
	9161:  uint16(sym__newline),
	9162:  uint16(sym_single_line_comment),
	9163:  uint16(230),
	9164:  uint16(2),
	9165:  uint16(sym__ws),
	9166:  uint16(aux_sym_node_repeat3),
	9167:  uint16(642),
	9168:  uint16(3),
	9169:  uint16(sym_multi_line_comment),
	9170:  uint16(sym__bom),
	9171:  uint16(sym__unicode_space),
	9172:  uint16(648),
	9173:  uint16(7),
	9174:  uint16(aux_sym__newline_token1),
	9175:  uint16(aux_sym__newline_token2),
	9176:  uint16(aux_sym__newline_token3),
	9177:  uint16(aux_sym__newline_token4),
	9178:  uint16(aux_sym__newline_token5),
	9179:  uint16(aux_sym__newline_token6),
	9180:  uint16(aux_sym__newline_token7),
	9181:  uint16(5),
	9182:  uint16(654),
	9183:  uint16(1),
	9184:  uint16(anon_sym_SLASH_SLASH),
	9185:  uint16(230),
	9186:  uint16(2),
	9187:  uint16(sym__ws),
	9188:  uint16(aux_sym_node_repeat3),
	9189:  uint16(236),
	9190:  uint16(2),
	9191:  uint16(sym__newline),
	9192:  uint16(sym_single_line_comment),
	9193:  uint16(642),
	9194:  uint16(3),
	9195:  uint16(sym_multi_line_comment),
	9196:  uint16(sym__bom),
	9197:  uint16(sym__unicode_space),
	9198:  uint16(652),
	9199:  uint16(7),
	9200:  uint16(aux_sym__newline_token1),
	9201:  uint16(aux_sym__newline_token2),
	9202:  uint16(aux_sym__newline_token3),
	9203:  uint16(aux_sym__newline_token4),
	9204:  uint16(aux_sym__newline_token5),
	9205:  uint16(aux_sym__newline_token6),
	9206:  uint16(aux_sym__newline_token7),
	9207:  uint16(5),
	9208:  uint16(19),
	9209:  uint16(1),
	9210:  uint16(anon_sym_SLASH_SLASH),
	9211:  uint16(183),
	9212:  uint16(2),
	9213:  uint16(sym__newline),
	9214:  uint16(sym_single_line_comment),
	9215:  uint16(224),
	9216:  uint16(2),
	9217:  uint16(sym__ws),
	9218:  uint16(aux_sym_node_repeat3),
	9219:  uint16(658),
	9220:  uint16(3),
	9221:  uint16(sym_multi_line_comment),
	9222:  uint16(sym__bom),
	9223:  uint16(sym__unicode_space),
	9224:  uint16(656),
	9225:  uint16(7),
	9226:  uint16(aux_sym__newline_token1),
	9227:  uint16(aux_sym__newline_token2),
	9228:  uint16(aux_sym__newline_token3),
	9229:  uint16(aux_sym__newline_token4),
	9230:  uint16(aux_sym__newline_token5),
	9231:  uint16(aux_sym__newline_token6),
	9232:  uint16(aux_sym__newline_token7),
	9233:  uint16(5),
	9234:  uint16(644),
	9235:  uint16(1),
	9236:  uint16(anon_sym_SLASH_SLASH),
	9237:  uint16(223),
	9238:  uint16(2),
	9239:  uint16(sym__ws),
	9240:  uint16(aux_sym_node_repeat3),
	9241:  uint16(276),
	9242:  uint16(2),
	9243:  uint16(sym__newline),
	9244:  uint16(sym_single_line_comment),
	9245:  uint16(662),
	9246:  uint16(3),
	9247:  uint16(sym_multi_line_comment),
	9248:  uint16(sym__bom),
	9249:  uint16(sym__unicode_space),
	9250:  uint16(660),
	9251:  uint16(7),
	9252:  uint16(aux_sym__newline_token1),
	9253:  uint16(aux_sym__newline_token2),
	9254:  uint16(aux_sym__newline_token3),
	9255:  uint16(aux_sym__newline_token4),
	9256:  uint16(aux_sym__newline_token5),
	9257:  uint16(aux_sym__newline_token6),
	9258:  uint16(aux_sym__newline_token7),
	9259:  uint16(5),
	9260:  uint16(650),
	9261:  uint16(1),
	9262:  uint16(anon_sym_SLASH_SLASH),
	9263:  uint16(56),
	9264:  uint16(2),
	9265:  uint16(sym__newline),
	9266:  uint16(sym_single_line_comment),
	9267:  uint16(225),
	9268:  uint16(2),
	9269:  uint16(sym__ws),
	9270:  uint16(aux_sym_node_repeat3),
	9271:  uint16(666),
	9272:  uint16(3),
	9273:  uint16(sym_multi_line_comment),
	9274:  uint16(sym__bom),
	9275:  uint16(sym__unicode_space),
	9276:  uint16(664),
	9277:  uint16(7),
	9278:  uint16(aux_sym__newline_token1),
	9279:  uint16(aux_sym__newline_token2),
	9280:  uint16(aux_sym__newline_token3),
	9281:  uint16(aux_sym__newline_token4),
	9282:  uint16(aux_sym__newline_token5),
	9283:  uint16(aux_sym__newline_token6),
	9284:  uint16(aux_sym__newline_token7),
	9285:  uint16(3),
	9286:  uint16(230),
	9287:  uint16(2),
	9288:  uint16(sym__ws),
	9289:  uint16(aux_sym_node_repeat3),
	9290:  uint16(668),
	9291:  uint16(3),
	9292:  uint16(sym_multi_line_comment),
	9293:  uint16(sym__bom),
	9294:  uint16(sym__unicode_space),
	9295:  uint16(208),
	9296:  uint16(10),
	9297:  uint16(anon_sym_LBRACE),
	9298:  uint16(anon_sym_BSLASH),
	9299:  uint16(aux_sym__newline_token1),
	9300:  uint16(aux_sym__newline_token2),
	9301:  uint16(aux_sym__newline_token3),
	9302:  uint16(aux_sym__newline_token4),
	9303:  uint16(aux_sym__newline_token5),
	9304:  uint16(aux_sym__newline_token6),
	9305:  uint16(aux_sym__newline_token7),
	9306:  uint16(anon_sym_SLASH_SLASH),
	9307:  uint16(5),
	9308:  uint16(654),
	9309:  uint16(1),
	9310:  uint16(anon_sym_SLASH_SLASH),
	9311:  uint16(226),
	9312:  uint16(2),
	9313:  uint16(sym__ws),
	9314:  uint16(aux_sym_node_repeat3),
	9315:  uint16(246),
	9316:  uint16(2),
	9317:  uint16(sym__newline),
	9318:  uint16(sym_single_line_comment),
	9319:  uint16(673),
	9320:  uint16(3),
	9321:  uint16(sym_multi_line_comment),
	9322:  uint16(sym__bom),
	9323:  uint16(sym__unicode_space),
	9324:  uint16(671),
	9325:  uint16(7),
	9326:  uint16(aux_sym__newline_token1),
	9327:  uint16(aux_sym__newline_token2),
	9328:  uint16(aux_sym__newline_token3),
	9329:  uint16(aux_sym__newline_token4),
	9330:  uint16(aux_sym__newline_token5),
	9331:  uint16(aux_sym__newline_token6),
	9332:  uint16(aux_sym__newline_token7),
	9333:  uint16(1),
	9334:  uint16(675),
	9335:  uint16(14),
	9336:  uint16(sym__eof),
	9337:  uint16(sym_multi_line_comment),
	9338:  uint16(anon_sym_SEMI),
	9339:  uint16(anon_sym_BSLASH),
	9340:  uint16(aux_sym__newline_token1),
	9341:  uint16(aux_sym__newline_token2),
	9342:  uint16(aux_sym__newline_token3),
	9343:  uint16(aux_sym__newline_token4),
	9344:  uint16(aux_sym__newline_token5),
	9345:  uint16(aux_sym__newline_token6),
	9346:  uint16(aux_sym__newline_token7),
	9347:  uint16(sym__bom),
	9348:  uint16(sym__unicode_space),
	9349:  uint16(anon_sym_SLASH_SLASH),
	9350:  uint16(1),
	9351:  uint16(677),
	9352:  uint16(14),
	9353:  uint16(sym__eof),
	9354:  uint16(sym_multi_line_comment),
	9355:  uint16(anon_sym_SEMI),
	9356:  uint16(anon_sym_BSLASH),
	9357:  uint16(aux_sym__newline_token1),
	9358:  uint16(aux_sym__newline_token2),
	9359:  uint16(aux_sym__newline_token3),
	9360:  uint16(aux_sym__newline_token4),
	9361:  uint16(aux_sym__newline_token5),
	9362:  uint16(aux_sym__newline_token6),
	9363:  uint16(aux_sym__newline_token7),
	9364:  uint16(sym__bom),
	9365:  uint16(sym__unicode_space),
	9366:  uint16(anon_sym_SLASH_SLASH),
	9367:  uint16(1),
	9368:  uint16(232),
	9369:  uint16(14),
	9370:  uint16(sym__eof),
	9371:  uint16(sym_multi_line_comment),
	9372:  uint16(anon_sym_SEMI),
	9373:  uint16(anon_sym_BSLASH),
	9374:  uint16(aux_sym__newline_token1),
	9375:  uint16(aux_sym__newline_token2),
	9376:  uint16(aux_sym__newline_token3),
	9377:  uint16(aux_sym__newline_token4),
	9378:  uint16(aux_sym__newline_token5),
	9379:  uint16(aux_sym__newline_token6),
	9380:  uint16(aux_sym__newline_token7),
	9381:  uint16(sym__bom),
	9382:  uint16(sym__unicode_space),
	9383:  uint16(anon_sym_SLASH_SLASH),
	9384:  uint16(1),
	9385:  uint16(679),
	9386:  uint16(14),
	9387:  uint16(sym__eof),
	9388:  uint16(sym_multi_line_comment),
	9389:  uint16(anon_sym_SEMI),
	9390:  uint16(anon_sym_BSLASH),
	9391:  uint16(aux_sym__newline_token1),
	9392:  uint16(aux_sym__newline_token2),
	9393:  uint16(aux_sym__newline_token3),
	9394:  uint16(aux_sym__newline_token4),
	9395:  uint16(aux_sym__newline_token5),
	9396:  uint16(aux_sym__newline_token6),
	9397:  uint16(aux_sym__newline_token7),
	9398:  uint16(sym__bom),
	9399:  uint16(sym__unicode_space),
	9400:  uint16(anon_sym_SLASH_SLASH),
	9401:  uint16(1),
	9402:  uint16(228),
	9403:  uint16(14),
	9404:  uint16(sym__eof),
	9405:  uint16(sym_multi_line_comment),
	9406:  uint16(anon_sym_SEMI),
	9407:  uint16(anon_sym_BSLASH),
	9408:  uint16(aux_sym__newline_token1),
	9409:  uint16(aux_sym__newline_token2),
	9410:  uint16(aux_sym__newline_token3),
	9411:  uint16(aux_sym__newline_token4),
	9412:  uint16(aux_sym__newline_token5),
	9413:  uint16(aux_sym__newline_token6),
	9414:  uint16(aux_sym__newline_token7),
	9415:  uint16(sym__bom),
	9416:  uint16(sym__unicode_space),
	9417:  uint16(anon_sym_SLASH_SLASH),
	9418:  uint16(1),
	9419:  uint16(218),
	9420:  uint16(14),
	9421:  uint16(sym__eof),
	9422:  uint16(sym_multi_line_comment),
	9423:  uint16(anon_sym_SEMI),
	9424:  uint16(anon_sym_BSLASH),
	9425:  uint16(aux_sym__newline_token1),
	9426:  uint16(aux_sym__newline_token2),
	9427:  uint16(aux_sym__newline_token3),
	9428:  uint16(aux_sym__newline_token4),
	9429:  uint16(aux_sym__newline_token5),
	9430:  uint16(aux_sym__newline_token6),
	9431:  uint16(aux_sym__newline_token7),
	9432:  uint16(sym__bom),
	9433:  uint16(sym__unicode_space),
	9434:  uint16(anon_sym_SLASH_SLASH),
	9435:  uint16(1),
	9436:  uint16(681),
	9437:  uint16(14),
	9438:  uint16(sym__eof),
	9439:  uint16(sym_multi_line_comment),
	9440:  uint16(anon_sym_SEMI),
	9441:  uint16(anon_sym_BSLASH),
	9442:  uint16(aux_sym__newline_token1),
	9443:  uint16(aux_sym__newline_token2),
	9444:  uint16(aux_sym__newline_token3),
	9445:  uint16(aux_sym__newline_token4),
	9446:  uint16(aux_sym__newline_token5),
	9447:  uint16(aux_sym__newline_token6),
	9448:  uint16(aux_sym__newline_token7),
	9449:  uint16(sym__bom),
	9450:  uint16(sym__unicode_space),
	9451:  uint16(anon_sym_SLASH_SLASH),
	9452:  uint16(1),
	9453:  uint16(683),
	9454:  uint16(14),
	9455:  uint16(sym__eof),
	9456:  uint16(sym_multi_line_comment),
	9457:  uint16(anon_sym_SEMI),
	9458:  uint16(anon_sym_BSLASH),
	9459:  uint16(aux_sym__newline_token1),
	9460:  uint16(aux_sym__newline_token2),
	9461:  uint16(aux_sym__newline_token3),
	9462:  uint16(aux_sym__newline_token4),
	9463:  uint16(aux_sym__newline_token5),
	9464:  uint16(aux_sym__newline_token6),
	9465:  uint16(aux_sym__newline_token7),
	9466:  uint16(sym__bom),
	9467:  uint16(sym__unicode_space),
	9468:  uint16(anon_sym_SLASH_SLASH),
	9469:  uint16(1),
	9470:  uint16(685),
	9471:  uint16(14),
	9472:  uint16(sym__eof),
	9473:  uint16(sym_multi_line_comment),
	9474:  uint16(anon_sym_SEMI),
	9475:  uint16(anon_sym_BSLASH),
	9476:  uint16(aux_sym__newline_token1),
	9477:  uint16(aux_sym__newline_token2),
	9478:  uint16(aux_sym__newline_token3),
	9479:  uint16(aux_sym__newline_token4),
	9480:  uint16(aux_sym__newline_token5),
	9481:  uint16(aux_sym__newline_token6),
	9482:  uint16(aux_sym__newline_token7),
	9483:  uint16(sym__bom),
	9484:  uint16(sym__unicode_space),
	9485:  uint16(anon_sym_SLASH_SLASH),
	9486:  uint16(1),
	9487:  uint16(687),
	9488:  uint16(14),
	9489:  uint16(sym__eof),
	9490:  uint16(sym_multi_line_comment),
	9491:  uint16(anon_sym_SEMI),
	9492:  uint16(anon_sym_BSLASH),
	9493:  uint16(aux_sym__newline_token1),
	9494:  uint16(aux_sym__newline_token2),
	9495:  uint16(aux_sym__newline_token3),
	9496:  uint16(aux_sym__newline_token4),
	9497:  uint16(aux_sym__newline_token5),
	9498:  uint16(aux_sym__newline_token6),
	9499:  uint16(aux_sym__newline_token7),
	9500:  uint16(sym__bom),
	9501:  uint16(sym__unicode_space),
	9502:  uint16(anon_sym_SLASH_SLASH),
	9503:  uint16(1),
	9504:  uint16(689),
	9505:  uint16(14),
	9506:  uint16(sym__eof),
	9507:  uint16(sym_multi_line_comment),
	9508:  uint16(anon_sym_SEMI),
	9509:  uint16(anon_sym_BSLASH),
	9510:  uint16(aux_sym__newline_token1),
	9511:  uint16(aux_sym__newline_token2),
	9512:  uint16(aux_sym__newline_token3),
	9513:  uint16(aux_sym__newline_token4),
	9514:  uint16(aux_sym__newline_token5),
	9515:  uint16(aux_sym__newline_token6),
	9516:  uint16(aux_sym__newline_token7),
	9517:  uint16(sym__bom),
	9518:  uint16(sym__unicode_space),
	9519:  uint16(anon_sym_SLASH_SLASH),
	9520:  uint16(1),
	9521:  uint16(691),
	9522:  uint16(14),
	9523:  uint16(sym__eof),
	9524:  uint16(sym_multi_line_comment),
	9525:  uint16(anon_sym_SEMI),
	9526:  uint16(anon_sym_BSLASH),
	9527:  uint16(aux_sym__newline_token1),
	9528:  uint16(aux_sym__newline_token2),
	9529:  uint16(aux_sym__newline_token3),
	9530:  uint16(aux_sym__newline_token4),
	9531:  uint16(aux_sym__newline_token5),
	9532:  uint16(aux_sym__newline_token6),
	9533:  uint16(aux_sym__newline_token7),
	9534:  uint16(sym__bom),
	9535:  uint16(sym__unicode_space),
	9536:  uint16(anon_sym_SLASH_SLASH),
	9537:  uint16(1),
	9538:  uint16(234),
	9539:  uint16(14),
	9540:  uint16(sym__eof),
	9541:  uint16(sym_multi_line_comment),
	9542:  uint16(anon_sym_SEMI),
	9543:  uint16(anon_sym_BSLASH),
	9544:  uint16(aux_sym__newline_token1),
	9545:  uint16(aux_sym__newline_token2),
	9546:  uint16(aux_sym__newline_token3),
	9547:  uint16(aux_sym__newline_token4),
	9548:  uint16(aux_sym__newline_token5),
	9549:  uint16(aux_sym__newline_token6),
	9550:  uint16(aux_sym__newline_token7),
	9551:  uint16(sym__bom),
	9552:  uint16(sym__unicode_space),
	9553:  uint16(anon_sym_SLASH_SLASH),
	9554:  uint16(1),
	9555:  uint16(693),
	9556:  uint16(14),
	9557:  uint16(sym__eof),
	9558:  uint16(sym_multi_line_comment),
	9559:  uint16(anon_sym_SEMI),
	9560:  uint16(anon_sym_BSLASH),
	9561:  uint16(aux_sym__newline_token1),
	9562:  uint16(aux_sym__newline_token2),
	9563:  uint16(aux_sym__newline_token3),
	9564:  uint16(aux_sym__newline_token4),
	9565:  uint16(aux_sym__newline_token5),
	9566:  uint16(aux_sym__newline_token6),
	9567:  uint16(aux_sym__newline_token7),
	9568:  uint16(sym__bom),
	9569:  uint16(sym__unicode_space),
	9570:  uint16(anon_sym_SLASH_SLASH),
	9571:  uint16(1),
	9572:  uint16(224),
	9573:  uint16(14),
	9574:  uint16(sym__eof),
	9575:  uint16(sym_multi_line_comment),
	9576:  uint16(anon_sym_SEMI),
	9577:  uint16(anon_sym_BSLASH),
	9578:  uint16(aux_sym__newline_token1),
	9579:  uint16(aux_sym__newline_token2),
	9580:  uint16(aux_sym__newline_token3),
	9581:  uint16(aux_sym__newline_token4),
	9582:  uint16(aux_sym__newline_token5),
	9583:  uint16(aux_sym__newline_token6),
	9584:  uint16(aux_sym__newline_token7),
	9585:  uint16(sym__bom),
	9586:  uint16(sym__unicode_space),
	9587:  uint16(anon_sym_SLASH_SLASH),
	9588:  uint16(1),
	9589:  uint16(695),
	9590:  uint16(14),
	9591:  uint16(sym__eof),
	9592:  uint16(sym_multi_line_comment),
	9593:  uint16(anon_sym_SEMI),
	9594:  uint16(anon_sym_BSLASH),
	9595:  uint16(aux_sym__newline_token1),
	9596:  uint16(aux_sym__newline_token2),
	9597:  uint16(aux_sym__newline_token3),
	9598:  uint16(aux_sym__newline_token4),
	9599:  uint16(aux_sym__newline_token5),
	9600:  uint16(aux_sym__newline_token6),
	9601:  uint16(aux_sym__newline_token7),
	9602:  uint16(sym__bom),
	9603:  uint16(sym__unicode_space),
	9604:  uint16(anon_sym_SLASH_SLASH),
	9605:  uint16(3),
	9606:  uint16(3),
	9607:  uint16(1),
	9608:  uint16(sym_multi_line_comment),
	9609:  uint16(697),
	9610:  uint16(5),
	9611:  uint16(sym__normal_bare_identifier),
	9612:  uint16(anon_sym_null),
	9613:  uint16(sym__digit),
	9614:  uint16(anon_sym_true),
	9615:  uint16(anon_sym_false),
	9616:  uint16(699),
	9617:  uint16(7),
	9618:  uint16(sym__raw_string),
	9619:  uint16(anon_sym_DQUOTE),
	9620:  uint16(anon_sym_PLUS),
	9621:  uint16(anon_sym_DASH),
	9622:  uint16(anon_sym_0x),
	9623:  uint16(anon_sym_0o),
	9624:  uint16(anon_sym_0b),
	9625:  uint16(5),
	9626:  uint16(3),
	9627:  uint16(1),
	9628:  uint16(sym_multi_line_comment),
	9629:  uint16(703),
	9630:  uint16(1),
	9631:  uint16(aux_sym_single_line_comment_token1),
	9632:  uint16(61),
	9633:  uint16(1),
	9634:  uint16(sym__newline),
	9635:  uint16(255),
	9636:  uint16(1),
	9637:  uint16(aux_sym_single_line_comment_repeat1),
	9638:  uint16(701),
	9639:  uint16(8),
	9640:  uint16(sym__eof),
	9641:  uint16(aux_sym__newline_token1),
	9642:  uint16(aux_sym__newline_token2),
	9643:  uint16(aux_sym__newline_token3),
	9644:  uint16(aux_sym__newline_token4),
	9645:  uint16(aux_sym__newline_token5),
	9646:  uint16(aux_sym__newline_token6),
	9647:  uint16(aux_sym__newline_token7),
	9648:  uint16(5),
	9649:  uint16(3),
	9650:  uint16(1),
	9651:  uint16(sym_multi_line_comment),
	9652:  uint16(707),
	9653:  uint16(1),
	9654:  uint16(aux_sym_single_line_comment_token1),
	9655:  uint16(237),
	9656:  uint16(1),
	9657:  uint16(sym__newline),
	9658:  uint16(259),
	9659:  uint16(1),
	9660:  uint16(aux_sym_single_line_comment_repeat1),
	9661:  uint16(705),
	9662:  uint16(8),
	9663:  uint16(sym__eof),
	9664:  uint16(aux_sym__newline_token1),
	9665:  uint16(aux_sym__newline_token2),
	9666:  uint16(aux_sym__newline_token3),
	9667:  uint16(aux_sym__newline_token4),
	9668:  uint16(aux_sym__newline_token5),
	9669:  uint16(aux_sym__newline_token6),
	9670:  uint16(aux_sym__newline_token7),
	9671:  uint16(5),
	9672:  uint16(3),
	9673:  uint16(1),
	9674:  uint16(sym_multi_line_comment),
	9675:  uint16(707),
	9676:  uint16(1),
	9677:  uint16(aux_sym_single_line_comment_token1),
	9678:  uint16(259),
	9679:  uint16(1),
	9680:  uint16(aux_sym_single_line_comment_repeat1),
	9681:  uint16(272),
	9682:  uint16(1),
	9683:  uint16(sym__newline),
	9684:  uint16(709),
	9685:  uint16(8),
	9686:  uint16(sym__eof),
	9687:  uint16(aux_sym__newline_token1),
	9688:  uint16(aux_sym__newline_token2),
	9689:  uint16(aux_sym__newline_token3),
	9690:  uint16(aux_sym__newline_token4),
	9691:  uint16(aux_sym__newline_token5),
	9692:  uint16(aux_sym__newline_token6),
	9693:  uint16(aux_sym__newline_token7),
	9694:  uint16(5),
	9695:  uint16(3),
	9696:  uint16(1),
	9697:  uint16(sym_multi_line_comment),
	9698:  uint16(713),
	9699:  uint16(1),
	9700:  uint16(aux_sym_single_line_comment_token1),
	9701:  uint16(60),
	9702:  uint16(1),
	9703:  uint16(sym__newline),
	9704:  uint16(254),
	9705:  uint16(1),
	9706:  uint16(aux_sym_single_line_comment_repeat1),
	9707:  uint16(711),
	9708:  uint16(8),
	9709:  uint16(sym__eof),
	9710:  uint16(aux_sym__newline_token1),
	9711:  uint16(aux_sym__newline_token2),
	9712:  uint16(aux_sym__newline_token3),
	9713:  uint16(aux_sym__newline_token4),
	9714:  uint16(aux_sym__newline_token5),
	9715:  uint16(aux_sym__newline_token6),
	9716:  uint16(aux_sym__newline_token7),
	9717:  uint16(5),
	9718:  uint16(3),
	9719:  uint16(1),
	9720:  uint16(sym_multi_line_comment),
	9721:  uint16(717),
	9722:  uint16(1),
	9723:  uint16(aux_sym_single_line_comment_token1),
	9724:  uint16(251),
	9725:  uint16(1),
	9726:  uint16(aux_sym_single_line_comment_repeat1),
	9727:  uint16(274),
	9728:  uint16(1),
	9729:  uint16(sym__newline),
	9730:  uint16(715),
	9731:  uint16(8),
	9732:  uint16(sym__eof),
	9733:  uint16(aux_sym__newline_token1),
	9734:  uint16(aux_sym__newline_token2),
	9735:  uint16(aux_sym__newline_token3),
	9736:  uint16(aux_sym__newline_token4),
	9737:  uint16(aux_sym__newline_token5),
	9738:  uint16(aux_sym__newline_token6),
	9739:  uint16(aux_sym__newline_token7),
	9740:  uint16(5),
	9741:  uint16(3),
	9742:  uint16(1),
	9743:  uint16(sym_multi_line_comment),
	9744:  uint16(707),
	9745:  uint16(1),
	9746:  uint16(aux_sym_single_line_comment_token1),
	9747:  uint16(55),
	9748:  uint16(1),
	9749:  uint16(sym__newline),
	9750:  uint16(259),
	9751:  uint16(1),
	9752:  uint16(aux_sym_single_line_comment_repeat1),
	9753:  uint16(719),
	9754:  uint16(8),
	9755:  uint16(sym__eof),
	9756:  uint16(aux_sym__newline_token1),
	9757:  uint16(aux_sym__newline_token2),
	9758:  uint16(aux_sym__newline_token3),
	9759:  uint16(aux_sym__newline_token4),
	9760:  uint16(aux_sym__newline_token5),
	9761:  uint16(aux_sym__newline_token6),
	9762:  uint16(aux_sym__newline_token7),
	9763:  uint16(5),
	9764:  uint16(3),
	9765:  uint16(1),
	9766:  uint16(sym_multi_line_comment),
	9767:  uint16(707),
	9768:  uint16(1),
	9769:  uint16(aux_sym_single_line_comment_token1),
	9770:  uint16(59),
	9771:  uint16(1),
	9772:  uint16(sym__newline),
	9773:  uint16(259),
	9774:  uint16(1),
	9775:  uint16(aux_sym_single_line_comment_repeat1),
	9776:  uint16(721),
	9777:  uint16(8),
	9778:  uint16(sym__eof),
	9779:  uint16(aux_sym__newline_token1),
	9780:  uint16(aux_sym__newline_token2),
	9781:  uint16(aux_sym__newline_token3),
	9782:  uint16(aux_sym__newline_token4),
	9783:  uint16(aux_sym__newline_token5),
	9784:  uint16(aux_sym__newline_token6),
	9785:  uint16(aux_sym__newline_token7),
	9786:  uint16(5),
	9787:  uint16(3),
	9788:  uint16(1),
	9789:  uint16(sym_multi_line_comment),
	9790:  uint16(725),
	9791:  uint16(1),
	9792:  uint16(aux_sym_single_line_comment_token1),
	9793:  uint16(244),
	9794:  uint16(1),
	9795:  uint16(sym__newline),
	9796:  uint16(250),
	9797:  uint16(1),
	9798:  uint16(aux_sym_single_line_comment_repeat1),
	9799:  uint16(723),
	9800:  uint16(8),
	9801:  uint16(sym__eof),
	9802:  uint16(aux_sym__newline_token1),
	9803:  uint16(aux_sym__newline_token2),
	9804:  uint16(aux_sym__newline_token3),
	9805:  uint16(aux_sym__newline_token4),
	9806:  uint16(aux_sym__newline_token5),
	9807:  uint16(aux_sym__newline_token6),
	9808:  uint16(aux_sym__newline_token7),
	9809:  uint16(9),
	9810:  uint16(3),
	9811:  uint16(1),
	9812:  uint16(sym_multi_line_comment),
	9813:  uint16(7),
	9814:  uint16(1),
	9815:  uint16(sym__normal_bare_identifier),
	9816:  uint16(13),
	9817:  uint16(1),
	9818:  uint16(anon_sym_DQUOTE),
	9819:  uint16(21),
	9820:  uint16(1),
	9821:  uint16(sym__raw_string),
	9822:  uint16(68),
	9823:  uint16(1),
	9824:  uint16(sym_identifier),
	9825:  uint16(192),
	9826:  uint16(1),
	9827:  uint16(sym__sign),
	9828:  uint16(194),
	9829:  uint16(1),
	9830:  uint16(sym__escaped_string),
	9831:  uint16(15),
	9832:  uint16(2),
	9833:  uint16(anon_sym_PLUS),
	9834:  uint16(anon_sym_DASH),
	9835:  uint16(218),
	9836:  uint16(2),
	9837:  uint16(sym__bare_identifier),
	9838:  uint16(sym_string),
	9839:  uint16(9),
	9840:  uint16(3),
	9841:  uint16(1),
	9842:  uint16(sym_multi_line_comment),
	9843:  uint16(7),
	9844:  uint16(1),
	9845:  uint16(sym__normal_bare_identifier),
	9846:  uint16(13),
	9847:  uint16(1),
	9848:  uint16(anon_sym_DQUOTE),
	9849:  uint16(21),
	9850:  uint16(1),
	9851:  uint16(sym__raw_string),
	9852:  uint16(64),
	9853:  uint16(1),
	9854:  uint16(sym_identifier),
	9855:  uint16(192),
	9856:  uint16(1),
	9857:  uint16(sym__sign),
	9858:  uint16(194),
	9859:  uint16(1),
	9860:  uint16(sym__escaped_string),
	9861:  uint16(15),
	9862:  uint16(2),
	9863:  uint16(anon_sym_PLUS),
	9864:  uint16(anon_sym_DASH),
	9865:  uint16(218),
	9866:  uint16(2),
	9867:  uint16(sym__bare_identifier),
	9868:  uint16(sym_string),
	9869:  uint16(4),
	9870:  uint16(3),
	9871:  uint16(1),
	9872:  uint16(sym_multi_line_comment),
	9873:  uint16(729),
	9874:  uint16(1),
	9875:  uint16(aux_sym_single_line_comment_token1),
	9876:  uint16(259),
	9877:  uint16(1),
	9878:  uint16(aux_sym_single_line_comment_repeat1),
	9879:  uint16(727),
	9880:  uint16(8),
	9881:  uint16(sym__eof),
	9882:  uint16(aux_sym__newline_token1),
	9883:  uint16(aux_sym__newline_token2),
	9884:  uint16(aux_sym__newline_token3),
	9885:  uint16(aux_sym__newline_token4),
	9886:  uint16(aux_sym__newline_token5),
	9887:  uint16(aux_sym__newline_token6),
	9888:  uint16(aux_sym__newline_token7),
	9889:  uint16(9),
	9890:  uint16(3),
	9891:  uint16(1),
	9892:  uint16(sym_multi_line_comment),
	9893:  uint16(7),
	9894:  uint16(1),
	9895:  uint16(sym__normal_bare_identifier),
	9896:  uint16(13),
	9897:  uint16(1),
	9898:  uint16(anon_sym_DQUOTE),
	9899:  uint16(21),
	9900:  uint16(1),
	9901:  uint16(sym__raw_string),
	9902:  uint16(66),
	9903:  uint16(1),
	9904:  uint16(sym_identifier),
	9905:  uint16(192),
	9906:  uint16(1),
	9907:  uint16(sym__sign),
	9908:  uint16(194),
	9909:  uint16(1),
	9910:  uint16(sym__escaped_string),
	9911:  uint16(15),
	9912:  uint16(2),
	9913:  uint16(anon_sym_PLUS),
	9914:  uint16(anon_sym_DASH),
	9915:  uint16(218),
	9916:  uint16(2),
	9917:  uint16(sym__bare_identifier),
	9918:  uint16(sym_string),
	9919:  uint16(7),
	9920:  uint16(75),
	9921:  uint16(1),
	9922:  uint16(anon_sym_LBRACE),
	9923:  uint16(732),
	9924:  uint16(1),
	9925:  uint16(anon_sym_BSLASH),
	9926:  uint16(263),
	9927:  uint16(1),
	9928:  uint16(aux_sym_node_repeat1),
	9929:  uint16(269),
	9930:  uint16(1),
	9931:  uint16(sym__escline),
	9932:  uint16(275),
	9933:  uint16(1),
	9934:  uint16(sym__node_space),
	9935:  uint16(265),
	9936:  uint16(2),
	9937:  uint16(sym__ws),
	9938:  uint16(aux_sym_node_repeat3),
	9939:  uint16(734),
	9940:  uint16(3),
	9941:  uint16(sym_multi_line_comment),
	9942:  uint16(sym__bom),
	9943:  uint16(sym__unicode_space),
	9944:  uint16(7),
	9945:  uint16(81),
	9946:  uint16(1),
	9947:  uint16(anon_sym_LBRACE),
	9948:  uint16(736),
	9949:  uint16(1),
	9950:  uint16(anon_sym_BSLASH),
	9951:  uint16(262),
	9952:  uint16(1),
	9953:  uint16(aux_sym_node_repeat1),
	9954:  uint16(269),
	9955:  uint16(1),
	9956:  uint16(sym__escline),
	9957:  uint16(275),
	9958:  uint16(1),
	9959:  uint16(sym__node_space),
	9960:  uint16(265),
	9961:  uint16(2),
	9962:  uint16(sym__ws),
	9963:  uint16(aux_sym_node_repeat3),
	9964:  uint16(739),
	9965:  uint16(3),
	9966:  uint16(sym_multi_line_comment),
	9967:  uint16(sym__bom),
	9968:  uint16(sym__unicode_space),
	9969:  uint16(7),
	9970:  uint16(77),
	9971:  uint16(1),
	9972:  uint16(anon_sym_LBRACE),
	9973:  uint16(732),
	9974:  uint16(1),
	9975:  uint16(anon_sym_BSLASH),
	9976:  uint16(262),
	9977:  uint16(1),
	9978:  uint16(aux_sym_node_repeat1),
	9979:  uint16(269),
	9980:  uint16(1),
	9981:  uint16(sym__escline),
	9982:  uint16(275),
	9983:  uint16(1),
	9984:  uint16(sym__node_space),
	9985:  uint16(265),
	9986:  uint16(2),
	9987:  uint16(sym__ws),
	9988:  uint16(aux_sym_node_repeat3),
	9989:  uint16(734),
	9990:  uint16(3),
	9991:  uint16(sym_multi_line_comment),
	9992:  uint16(sym__bom),
	9993:  uint16(sym__unicode_space),
	9994:  uint16(8),
	9995:  uint16(3),
	9996:  uint16(1),
	9997:  uint16(sym_multi_line_comment),
	9998:  uint16(33),
	9999:  uint16(1),
	10000: uint16(sym__digit),
	10001: uint16(556),
	10002: uint16(1),
	10003: uint16(anon_sym_EQ),
	10004: uint16(742),
	10005: uint16(1),
	10006: uint16(sym___identifier_char_no_digit),
	10007: uint16(744),
	10008: uint16(1),
	10009: uint16(anon_sym_0x),
	10010: uint16(746),
	10011: uint16(1),
	10012: uint16(anon_sym_0o),
	10013: uint16(748),
	10014: uint16(1),
	10015: uint16(anon_sym_0b),
	10016: uint16(135),
	10017: uint16(1),
	10018: uint16(sym__integer),
	10019: uint16(5),
	10020: uint16(149),
	10021: uint16(1),
	10022: uint16(anon_sym_LBRACE),
	10023: uint16(750),
	10024: uint16(1),
	10025: uint16(anon_sym_BSLASH),
	10026: uint16(267),
	10027: uint16(1),
	10028: uint16(sym__escline),
	10029: uint16(230),
	10030: uint16(2),
	10031: uint16(sym__ws),
	10032: uint16(aux_sym_node_repeat3),
	10033: uint16(753),
	10034: uint16(3),
	10035: uint16(sym_multi_line_comment),
	10036: uint16(sym__bom),
	10037: uint16(sym__unicode_space),
	10038: uint16(3),
	10039: uint16(180),
	10040: uint16(2),
	10041: uint16(anon_sym_LBRACE),
	10042: uint16(anon_sym_BSLASH),
	10043: uint16(230),
	10044: uint16(2),
	10045: uint16(sym__ws),
	10046: uint16(aux_sym_node_repeat3),
	10047: uint16(756),
	10048: uint16(3),
	10049: uint16(sym_multi_line_comment),
	10050: uint16(sym__bom),
	10051: uint16(sym__unicode_space),
	10052: uint16(3),
	10053: uint16(187),
	10054: uint16(2),
	10055: uint16(anon_sym_LBRACE),
	10056: uint16(anon_sym_BSLASH),
	10057: uint16(266),
	10058: uint16(2),
	10059: uint16(sym__ws),
	10060: uint16(aux_sym_node_repeat3),
	10061: uint16(759),
	10062: uint16(3),
	10063: uint16(sym_multi_line_comment),
	10064: uint16(sym__bom),
	10065: uint16(sym__unicode_space),
	10066: uint16(3),
	10067: uint16(187),
	10068: uint16(2),
	10069: uint16(anon_sym_LBRACE),
	10070: uint16(anon_sym_BSLASH),
	10071: uint16(230),
	10072: uint16(2),
	10073: uint16(sym__ws),
	10074: uint16(aux_sym_node_repeat3),
	10075: uint16(762),
	10076: uint16(3),
	10077: uint16(sym_multi_line_comment),
	10078: uint16(sym__bom),
	10079: uint16(sym__unicode_space),
	10080: uint16(3),
	10081: uint16(149),
	10082: uint16(2),
	10083: uint16(anon_sym_LBRACE),
	10084: uint16(anon_sym_BSLASH),
	10085: uint16(268),
	10086: uint16(2),
	10087: uint16(sym__ws),
	10088: uint16(aux_sym_node_repeat3),
	10089: uint16(765),
	10090: uint16(3),
	10091: uint16(sym_multi_line_comment),
	10092: uint16(sym__bom),
	10093: uint16(sym__unicode_space),
	10094: uint16(5),
	10095: uint16(3),
	10096: uint16(1),
	10097: uint16(sym_multi_line_comment),
	10098: uint16(768),
	10099: uint16(1),
	10100: uint16(sym__digit),
	10101: uint16(200),
	10102: uint16(1),
	10103: uint16(sym__integer),
	10104: uint16(287),
	10105: uint16(1),
	10106: uint16(sym__sign),
	10107: uint16(770),
	10108: uint16(2),
	10109: uint16(anon_sym_PLUS),
	10110: uint16(anon_sym_DASH),
	10111: uint16(6),
	10112: uint16(3),
	10113: uint16(1),
	10114: uint16(sym_multi_line_comment),
	10115: uint16(33),
	10116: uint16(1),
	10117: uint16(sym__digit),
	10118: uint16(744),
	10119: uint16(1),
	10120: uint16(anon_sym_0x),
	10121: uint16(746),
	10122: uint16(1),
	10123: uint16(anon_sym_0o),
	10124: uint16(748),
	10125: uint16(1),
	10126: uint16(anon_sym_0b),
	10127: uint16(135),
	10128: uint16(1),
	10129: uint16(sym__integer),
	10130: uint16(1),
	10131: uint16(218),
	10132: uint16(5),
	10133: uint16(sym_multi_line_comment),
	10134: uint16(anon_sym_LBRACE),
	10135: uint16(anon_sym_BSLASH),
	10136: uint16(sym__bom),
	10137: uint16(sym__unicode_space),
	10138: uint16(5),
	10139: uint16(3),
	10140: uint16(1),
	10141: uint16(sym_multi_line_comment),
	10142: uint16(772),
	10143: uint16(1),
	10144: uint16(anon_sym_DQUOTE),
	10145: uint16(774),
	10146: uint16(1),
	10147: uint16(aux_sym__escaped_string_token1),
	10148: uint16(776),
	10149: uint16(1),
	10150: uint16(sym_escape),
	10151: uint16(282),
	10152: uint16(1),
	10153: uint16(aux_sym__escaped_string_repeat1),
	10154: uint16(1),
	10155: uint16(234),
	10156: uint16(5),
	10157: uint16(sym_multi_line_comment),
	10158: uint16(anon_sym_LBRACE),
	10159: uint16(anon_sym_BSLASH),
	10160: uint16(sym__bom),
	10161: uint16(sym__unicode_space),
	10162: uint16(1),
	10163: uint16(232),
	10164: uint16(5),
	10165: uint16(sym_multi_line_comment),
	10166: uint16(anon_sym_LBRACE),
	10167: uint16(anon_sym_BSLASH),
	10168: uint16(sym__bom),
	10169: uint16(sym__unicode_space),
	10170: uint16(1),
	10171: uint16(224),
	10172: uint16(5),
	10173: uint16(sym_multi_line_comment),
	10174: uint16(anon_sym_LBRACE),
	10175: uint16(anon_sym_BSLASH),
	10176: uint16(sym__bom),
	10177: uint16(sym__unicode_space),
	10178: uint16(5),
	10179: uint16(3),
	10180: uint16(1),
	10181: uint16(sym_multi_line_comment),
	10182: uint16(774),
	10183: uint16(1),
	10184: uint16(aux_sym__escaped_string_token1),
	10185: uint16(776),
	10186: uint16(1),
	10187: uint16(sym_escape),
	10188: uint16(778),
	10189: uint16(1),
	10190: uint16(anon_sym_DQUOTE),
	10191: uint16(283),
	10192: uint16(1),
	10193: uint16(aux_sym__escaped_string_repeat1),
	10194: uint16(4),
	10195: uint16(3),
	10196: uint16(1),
	10197: uint16(sym_multi_line_comment),
	10198: uint16(780),
	10199: uint16(1),
	10200: uint16(sym__identifier_char),
	10201: uint16(279),
	10202: uint16(1),
	10203: uint16(aux_sym__bare_identifier_repeat1),
	10204: uint16(544),
	10205: uint16(2),
	10206: uint16(anon_sym_EQ),
	10207: uint16(anon_sym_RPAREN),
	10208: uint16(4),
	10209: uint16(3),
	10210: uint16(1),
	10211: uint16(sym_multi_line_comment),
	10212: uint16(782),
	10213: uint16(1),
	10214: uint16(sym__identifier_char),
	10215: uint16(279),
	10216: uint16(1),
	10217: uint16(aux_sym__bare_identifier_repeat1),
	10218: uint16(539),
	10219: uint16(2),
	10220: uint16(anon_sym_EQ),
	10221: uint16(anon_sym_RPAREN),
	10222: uint16(5),
	10223: uint16(3),
	10224: uint16(1),
	10225: uint16(sym_multi_line_comment),
	10226: uint16(774),
	10227: uint16(1),
	10228: uint16(aux_sym__escaped_string_token1),
	10229: uint16(776),
	10230: uint16(1),
	10231: uint16(sym_escape),
	10232: uint16(785),
	10233: uint16(1),
	10234: uint16(anon_sym_DQUOTE),
	10235: uint16(277),
	10236: uint16(1),
	10237: uint16(aux_sym__escaped_string_repeat1),
	10238: uint16(1),
	10239: uint16(228),
	10240: uint16(5),
	10241: uint16(sym_multi_line_comment),
	10242: uint16(anon_sym_LBRACE),
	10243: uint16(anon_sym_BSLASH),
	10244: uint16(sym__bom),
	10245: uint16(sym__unicode_space),
	10246: uint16(5),
	10247: uint16(3),
	10248: uint16(1),
	10249: uint16(sym_multi_line_comment),
	10250: uint16(774),
	10251: uint16(1),
	10252: uint16(aux_sym__escaped_string_token1),
	10253: uint16(776),
	10254: uint16(1),
	10255: uint16(sym_escape),
	10256: uint16(787),
	10257: uint16(1),
	10258: uint16(anon_sym_DQUOTE),
	10259: uint16(283),
	10260: uint16(1),
	10261: uint16(aux_sym__escaped_string_repeat1),
	10262: uint16(5),
	10263: uint16(3),
	10264: uint16(1),
	10265: uint16(sym_multi_line_comment),
	10266: uint16(789),
	10267: uint16(1),
	10268: uint16(anon_sym_DQUOTE),
	10269: uint16(791),
	10270: uint16(1),
	10271: uint16(aux_sym__escaped_string_token1),
	10272: uint16(794),
	10273: uint16(1),
	10274: uint16(sym_escape),
	10275: uint16(283),
	10276: uint16(1),
	10277: uint16(aux_sym__escaped_string_repeat1),
	10278: uint16(4),
	10279: uint16(3),
	10280: uint16(1),
	10281: uint16(sym_multi_line_comment),
	10282: uint16(797),
	10283: uint16(1),
	10284: uint16(sym__identifier_char),
	10285: uint16(278),
	10286: uint16(1),
	10287: uint16(aux_sym__bare_identifier_repeat1),
	10288: uint16(548),
	10289: uint16(2),
	10290: uint16(anon_sym_EQ),
	10291: uint16(anon_sym_RPAREN),
	10292: uint16(3),
	10293: uint16(3),
	10294: uint16(1),
	10295: uint16(sym_multi_line_comment),
	10296: uint16(801),
	10297: uint16(1),
	10298: uint16(aux_sym__escaped_string_token1),
	10299: uint16(799),
	10300: uint16(2),
	10301: uint16(anon_sym_DQUOTE),
	10302: uint16(sym_escape),
	10303: uint16(3),
	10304: uint16(3),
	10305: uint16(1),
	10306: uint16(sym_multi_line_comment),
	10307: uint16(768),
	10308: uint16(1),
	10309: uint16(sym__digit),
	10310: uint16(174),
	10311: uint16(1),
	10312: uint16(sym__integer),
	10313: uint16(3),
	10314: uint16(3),
	10315: uint16(1),
	10316: uint16(sym_multi_line_comment),
	10317: uint16(768),
	10318: uint16(1),
	10319: uint16(sym__digit),
	10320: uint16(204),
	10321: uint16(1),
	10322: uint16(sym__integer),
	10323: uint16(2),
	10324: uint16(3),
	10325: uint16(1),
	10326: uint16(sym_multi_line_comment),
	10327: uint16(803),
	10328: uint16(2),
	10329: uint16(anon_sym_0),
	10330: uint16(anon_sym_1),
	10331: uint16(2),
	10332: uint16(3),
	10333: uint16(1),
	10334: uint16(sym_multi_line_comment),
	10335: uint16(805),
	10336: uint16(2),
	10337: uint16(anon_sym_0),
	10338: uint16(anon_sym_1),
	10339: uint16(2),
	10340: uint16(3),
	10341: uint16(1),
	10342: uint16(sym_multi_line_comment),
	10343: uint16(554),
	10344: uint16(2),
	10345: uint16(anon_sym_EQ),
	10346: uint16(anon_sym_RPAREN),
	10347: uint16(3),
	10348: uint16(3),
	10349: uint16(1),
	10350: uint16(sym_multi_line_comment),
	10351: uint16(768),
	10352: uint16(1),
	10353: uint16(sym__digit),
	10354: uint16(180),
	10355: uint16(1),
	10356: uint16(sym__integer),
	10357: uint16(3),
	10358: uint16(3),
	10359: uint16(1),
	10360: uint16(sym_multi_line_comment),
	10361: uint16(556),
	10362: uint16(1),
	10363: uint16(anon_sym_RPAREN),
	10364: uint16(742),
	10365: uint16(1),
	10366: uint16(sym___identifier_char_no_digit),
	10367: uint16(2),
	10368: uint16(3),
	10369: uint16(1),
	10370: uint16(sym_multi_line_comment),
	10371: uint16(581),
	10372: uint16(1),
	10373: uint16(anon_sym_RPAREN),
	10374: uint16(2),
	10375: uint16(3),
	10376: uint16(1),
	10377: uint16(sym_multi_line_comment),
	10378: uint16(579),
	10379: uint16(1),
	10380: uint16(anon_sym_RPAREN),
	10381: uint16(2),
	10382: uint16(3),
	10383: uint16(1),
	10384: uint16(sym_multi_line_comment),
	10385: uint16(807),
	10386: uint16(1),
	10387: uint16(anon_sym_EQ),
	10388: uint16(2),
	10389: uint16(3),
	10390: uint16(1),
	10391: uint16(sym_multi_line_comment),
	10392: uint16(566),
	10393: uint16(1),
	10394: uint16(anon_sym_RPAREN),
	10395: uint16(2),
	10396: uint16(3),
	10397: uint16(1),
	10398: uint16(sym_multi_line_comment),
	10399: uint16(809),
	10400: uint16(1),
	10402: uint16(2),
	10403: uint16(3),
	10404: uint16(1),
	10405: uint16(sym_multi_line_comment),
	10406: uint16(811),
	10407: uint16(1),
	10408: uint16(sym__hex_digit),
	10409: uint16(2),
	10410: uint16(3),
	10411: uint16(1),
	10412: uint16(sym_multi_line_comment),
	10413: uint16(813),
	10414: uint16(1),
	10415: uint16(aux_sym__octal_token1),
	10416: uint16(2),
	10417: uint16(3),
	10418: uint16(1),
	10419: uint16(sym_multi_line_comment),
	10420: uint16(815),
	10421: uint16(1),
	10422: uint16(aux_sym__octal_token1),
	10423: uint16(2),
	10424: uint16(3),
	10425: uint16(1),
	10426: uint16(sym_multi_line_comment),
	10427: uint16(817),
	10428: uint16(1),
	10429: uint16(anon_sym_RPAREN),
	10430: uint16(2),
	10431: uint16(3),
	10432: uint16(1),
	10433: uint16(sym_multi_line_comment),
	10434: uint16(819),
	10435: uint16(1),
	10436: uint16(sym__hex_digit),
	10437: uint16(2),
	10438: uint16(3),
	10439: uint16(1),
	10440: uint16(sym_multi_line_comment),
	10441: uint16(821),
	10442: uint16(1),
	10443: uint16(anon_sym_RPAREN),
}

var ts_small_parse_table_map = [302]uint32_t{
	1:   uint32(130),
	2:   uint32(260),
	3:   uint32(390),
	4:   uint32(520),
	5:   uint32(650),
	6:   uint32(780),
	7:   uint32(910),
	8:   uint32(980),
	9:   uint32(1085),
	10:  uint32(1186),
	11:  uint32(1287),
	12:  uint32(1385),
	13:  uint32(1483),
	14:  uint32(1534),
	15:  uint32(1600),
	16:  uint32(1665),
	17:  uint32(1730),
	18:  uint32(1795),
	19:  uint32(1860),
	20:  uint32(1925),
	21:  uint32(1990),
	22:  uint32(2055),
	23:  uint32(2120),
	24:  uint32(2185),
	25:  uint32(2250),
	26:  uint32(2295),
	27:  uint32(2360),
	28:  uint32(2425),
	29:  uint32(2490),
	30:  uint32(2555),
	31:  uint32(2619),
	32:  uint32(2683),
	33:  uint32(2723),
	34:  uint32(2787),
	35:  uint32(2851),
	36:  uint32(2915),
	37:  uint32(2955),
	38:  uint32(2995),
	39:  uint32(3059),
	40:  uint32(3123),
	41:  uint32(3187),
	42:  uint32(3251),
	43:  uint32(3315),
	44:  uint32(3379),
	45:  uint32(3443),
	46:  uint32(3507),
	47:  uint32(3571),
	48:  uint32(3611),
	49:  uint32(3675),
	50:  uint32(3739),
	51:  uint32(3803),
	52:  uint32(3843),
	53:  uint32(3904),
	54:  uint32(3938),
	55:  uint32(3972),
	56:  uint32(4006),
	57:  uint32(4040),
	58:  uint32(4074),
	59:  uint32(4108),
	60:  uint32(4142),
	61:  uint32(4195),
	62:  uint32(4248),
	63:  uint32(4301),
	64:  uint32(4354),
	65:  uint32(4407),
	66:  uint32(4472),
	67:  uint32(4525),
	68:  uint32(4578),
	69:  uint32(4631),
	70:  uint32(4684),
	71:  uint32(4737),
	72:  uint32(4771),
	73:  uint32(4812),
	74:  uint32(4847),
	75:  uint32(4882),
	76:  uint32(4917),
	77:  uint32(4952),
	78:  uint32(5008),
	79:  uint32(5049),
	80:  uint32(5078),
	81:  uint32(5119),
	82:  uint32(5160),
	83:  uint32(5201),
	84:  uint32(5242),
	85:  uint32(5283),
	86:  uint32(5324),
	87:  uint32(5365),
	88:  uint32(5404),
	89:  uint32(5445),
	90:  uint32(5486),
	91:  uint32(5527),
	92:  uint32(5568),
	93:  uint32(5609),
	94:  uint32(5650),
	95:  uint32(5691),
	96:  uint32(5732),
	97:  uint32(5773),
	98:  uint32(5802),
	99:  uint32(5837),
	100: uint32(5878),
	101: uint32(5919),
	102: uint32(5960),
	103: uint32(6001),
	104: uint32(6042),
	105: uint32(6083),
	106: uint32(6124),
	107: uint32(6165),
	108: uint32(6206),
	109: uint32(6247),
	110: uint32(6288),
	111: uint32(6329),
	112: uint32(6370),
	113: uint32(6411),
	114: uint32(6452),
	115: uint32(6493),
	116: uint32(6534),
	117: uint32(6575),
	118: uint32(6604),
	119: uint32(6645),
	120: uint32(6675),
	121: uint32(6705),
	122: uint32(6735),
	123: uint32(6789),
	124: uint32(6819),
	125: uint32(6849),
	126: uint32(6903),
	127: uint32(6926),
	128: uint32(6949),
	129: uint32(6976),
	130: uint32(6999),
	131: uint32(7022),
	132: uint32(7045),
	133: uint32(7068),
	134: uint32(7097),
	135: uint32(7120),
	136: uint32(7143),
	137: uint32(7166),
	138: uint32(7189),
	139: uint32(7212),
	140: uint32(7235),
	141: uint32(7258),
	142: uint32(7281),
	143: uint32(7304),
	144: uint32(7327),
	145: uint32(7350),
	146: uint32(7377),
	147: uint32(7400),
	148: uint32(7423),
	149: uint32(7446),
	150: uint32(7469),
	151: uint32(7492),
	152: uint32(7515),
	153: uint32(7544),
	154: uint32(7567),
	155: uint32(7590),
	156: uint32(7613),
	157: uint32(7636),
	158: uint32(7659),
	159: uint32(7682),
	160: uint32(7705),
	161: uint32(7728),
	162: uint32(7751),
	163: uint32(7778),
	164: uint32(7805),
	165: uint32(7832),
	166: uint32(7855),
	167: uint32(7878),
	168: uint32(7901),
	169: uint32(7924),
	170: uint32(7947),
	171: uint32(7973),
	172: uint32(7999),
	173: uint32(8025),
	174: uint32(8051),
	175: uint32(8077),
	176: uint32(8103),
	177: uint32(8129),
	178: uint32(8153),
	179: uint32(8179),
	180: uint32(8205),
	181: uint32(8231),
	182: uint32(8255),
	183: uint32(8281),
	184: uint32(8307),
	185: uint32(8331),
	186: uint32(8365),
	187: uint32(8390),
	188: uint32(8415),
	189: uint32(8440),
	190: uint32(8462),
	191: uint32(8484),
	192: uint32(8512),
	193: uint32(8532),
	194: uint32(8560),
	195: uint32(8580),
	196: uint32(8600),
	197: uint32(8619),
	198: uint32(8638),
	199: uint32(8657),
	200: uint32(8680),
	201: uint32(8699),
	202: uint32(8718),
	203: uint32(8737),
	204: uint32(8760),
	205: uint32(8783),
	206: uint32(8806),
	207: uint32(8825),
	208: uint32(8844),
	209: uint32(8867),
	210: uint32(8886),
	211: uint32(8909),
	212: uint32(8928),
	213: uint32(8947),
	214: uint32(8966),
	215: uint32(8985),
	216: uint32(9004),
	217: uint32(9023),
	218: uint32(9042),
	219: uint32(9061),
	220: uint32(9080),
	221: uint32(9103),
	222: uint32(9129),
	223: uint32(9155),
	224: uint32(9181),
	225: uint32(9207),
	226: uint32(9233),
	227: uint32(9259),
	228: uint32(9285),
	229: uint32(9307),
	230: uint32(9333),
	231: uint32(9350),
	232: uint32(9367),
	233: uint32(9384),
	234: uint32(9401),
	235: uint32(9418),
	236: uint32(9435),
	237: uint32(9452),
	238: uint32(9469),
	239: uint32(9486),
	240: uint32(9503),
	241: uint32(9520),
	242: uint32(9537),
	243: uint32(9554),
	244: uint32(9571),
	245: uint32(9588),
	246: uint32(9605),
	247: uint32(9625),
	248: uint32(9648),
	249: uint32(9671),
	250: uint32(9694),
	251: uint32(9717),
	252: uint32(9740),
	253: uint32(9763),
	254: uint32(9786),
	255: uint32(9809),
	256: uint32(9839),
	257: uint32(9869),
	258: uint32(9889),
	259: uint32(9919),
	260: uint32(9944),
	261: uint32(9969),
	262: uint32(9994),
	263: uint32(10019),
	264: uint32(10038),
	265: uint32(10052),
	266: uint32(10066),
	267: uint32(10080),
	268: uint32(10094),
	269: uint32(10111),
	270: uint32(10130),
	271: uint32(10138),
	272: uint32(10154),
	273: uint32(10162),
	274: uint32(10170),
	275: uint32(10178),
	276: uint32(10194),
	277: uint32(10208),
	278: uint32(10222),
	279: uint32(10238),
	280: uint32(10246),
	281: uint32(10262),
	282: uint32(10278),
	283: uint32(10292),
	284: uint32(10303),
	285: uint32(10313),
	286: uint32(10323),
	287: uint32(10331),
	288: uint32(10339),
	289: uint32(10347),
	290: uint32(10357),
	291: uint32(10367),
	292: uint32(10374),
	293: uint32(10381),
	294: uint32(10388),
	295: uint32(10395),
	296: uint32(10402),
	297: uint32(10409),
	298: uint32(10416),
	299: uint32(10423),
	300: uint32(10430),
	301: uint32(10437),
}

var ts_parse_actions = [823]TSParseActionEntry{
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(218)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(127)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(273)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(192)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(252)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(194)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(290)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(35)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(152)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(217)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(264)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(298)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(299)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(288)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(219)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(229)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(27)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(142)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(171)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(141)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(162)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(148)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(160)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(303)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(280)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(296)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(227)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	76: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_node_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(229)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_node_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(27)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	90: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat2),
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
		Fsymbol:      uint16(aux_sym_document_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(218)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(127)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_document_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(9)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_document_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(273)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(192)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_document_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(54)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_document_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(252)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_document_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(194)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_document),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(51)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_document),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(242)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fsymbol:      uint16(sym__node_space),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(229)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_document),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(240)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(46)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(241)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
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
		Fcount: uint8(1),
	}})),
	179: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	183: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__node_space),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__node_space),
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
		Fsymbol:      uint16(sym__node_space),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__node_space),
	})))),
	191: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(235)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(232)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat3),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat3),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	211: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat3),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(49)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_document),
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
		Fsymbol:      uint16(sym__node_space),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_single_line_comment),
	})))),
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
		Fcount: uint8(1),
	}})),
	221: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_single_line_comment),
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
		Fcount: uint8(1),
	}})),
	223: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__escline),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__escline),
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
		Fcount: uint8(1),
	}})),
	227: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__escline),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__escline),
	})))),
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
		Fcount: uint8(1),
	}})),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_node_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_node_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_single_line_comment),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_single_line_comment),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(261)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(169)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(271)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(219)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(156)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	253: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(73)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	256: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	257: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	259: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(227)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(100)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	265: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(247)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(243)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	271: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_document),
	})))),
	272: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	274: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(245)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(155)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	281: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym__integer),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(119)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	287: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(114)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	299: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(aux_sym_node_repeat2),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_node_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(227)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	310: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(100)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(134)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(103)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(151)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(137)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(138)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(120)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(157)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	329: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	333: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__integer_repeat1),
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
		Fsymbol:      uint16(aux_sym__integer_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(99)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(227)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	341: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(126)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(143)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(167)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(140)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(158)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(106)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(146)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(163)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(105)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(149)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(150)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(159)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(139)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(161)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(168)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(170)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__integer),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym__node_space),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(123)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(126)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	388: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(126)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	391: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(122)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat3),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(126)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(14),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	399: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(16),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__binary_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	403: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__binary_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(130)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(15),
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
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(12),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(8),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(14),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__decimal),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(291)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(270)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(12),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(3),
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
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(7),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(15),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(14),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(4),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_node),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(16),
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
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(8),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(11),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__binary),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(3),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(15),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(11),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_node),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__decimal),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(286)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(12),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_node),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(7),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(8),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	470: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	474: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_node),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(16),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	478: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__binary),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__binary),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(3),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(7),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	490: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_node),
		Fproduction_id: uint16(11),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	494: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__octal),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(181)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__octal),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(182)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__decimal),
		Fproduction_id: uint16(10),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__hex),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(178)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__octal),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(184)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__hex),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__decimal),
		Fproduction_id: uint16(13),
	})))),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym__octal_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	522: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__octal_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(181)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__hex_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__hex_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(184)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__hex),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(177)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	534: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat1),
	})))),
	535: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(231)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_node_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(193)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__bare_identifier_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	542: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__bare_identifier_repeat1),
	})))),
	543: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(188)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	545: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__bare_identifier),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	549: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__bare_identifier),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_value),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_identifier),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	557: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__bare_identifier),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	561: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(231)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(205)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_string),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
	})))),
	570: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat3),
	})))),
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
		Fcount:    uint8(3),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
	})))),
	573: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat3),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(231)),
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
		Fcount:    uint8(3),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	576: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
	})))),
	577: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat3),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(205)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__escaped_string),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__escaped_string),
		Fproduction_id: uint16(1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__exponent),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	586: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__node_space),
	})))),
	587: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__decimal),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__decimal),
		Fproduction_id: uint16(10),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__exponent),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	595: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat3),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(205)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	598: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__node_space),
	})))),
	599: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	601: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(205)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_prop),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	606: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__node_field_comment),
		Fproduction_id: uint16(9),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	608: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__node_space),
	})))),
	609: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__node_space),
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
		Fcount:    uint8(3),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	611: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__node_space),
	})))),
	612: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(205)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_node_field),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(207)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_field),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_number),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	624: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym__decimal),
		Fproduction_id: uint16(13),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_keyword),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_boolean),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_value),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__node_field_comment),
		Fproduction_id: uint16(6),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
	})))),
	635: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fcount:    uint8(3),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
	})))),
	638: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__node_space),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(210)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(281)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(230)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(253)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(179)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(249)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	653: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	655: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(183)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(224)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(276)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(225)),
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
		Fsymbol:      uint16(aux_sym_node_repeat3),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(230)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(246)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(226)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_node_children),
		Fproduction_id: uint16(5),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_node_children),
		Fproduction_id: uint16(5),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_node_children),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_node_children),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_node_children),
		Fproduction_id: uint16(5),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_node_children),
		Fproduction_id: uint16(5),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_node_children),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_node_children),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	692: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_node_children),
		Fproduction_id: uint16(5),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_node_children),
		Fproduction_id: uint16(5),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	696: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_node_children),
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
		Fcount: uint8(1),
	}})),
	698: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(61)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(255)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(237)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(259)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(272)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(254)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(274)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(251)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(55)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(244)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(250)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_single_line_comment_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_single_line_comment_repeat1),
	})))),
	731: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(259)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(265)),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	737: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(228)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	740: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_node_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(265)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(284)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(302)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(300)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(289)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(228)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(230)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	757: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(230)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(266)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	763: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__node_space),
	})))),
	764: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_space),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(268)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(287)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(196)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(285)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(285)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(293)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(279)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	783: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__bare_identifier_repeat1),
	})))),
	784: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(279)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(294)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(197)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__escaped_string_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__escaped_string_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(285)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	795: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__escaped_string_repeat1),
	})))),
	796: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(285)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(278)),
	}})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__escaped_string_repeat1),
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
		Fcount: uint8(1),
	}})),
	802: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__escaped_string_repeat1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(147)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(166)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(67)),
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
	810: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(185)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(173)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(176)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(248)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(175)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_annotation_type),
	})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token__eof = 0
const ts_external_token_multi_line_comment = 1
const ts_external_token__raw_string = 2

var ts_external_scanner_symbol_map = [3]TSSymbol{
	0: uint16(sym__eof),
	1: uint16(sym_multi_line_comment),
	2: uint16(sym__raw_string),
}

var ts_external_scanner_states = [5][3]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
	},
	2: {
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
	},
	3: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
	},
	4: {
		1: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_kdl(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fversion:                   uint32(LANGUAGE_VERSION),
	Fsymbol_count:              uint32(SYMBOL_COUNT),
	Falias_count:               uint32(ALIAS_COUNT),
	Ftoken_count:               uint32(TOKEN_COUNT),
	Fexternal_token_count:      uint32(EXTERNAL_TOKEN_COUNT),
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
	Fkeyword_capture_token:     uint16(sym__normal_bare_identifier),
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
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_kdl_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_kdl_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_kdl_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_kdl_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_kdl_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00_normal_bare_identifier\x00node_comment\x00{\x00}\x00;\x00_identifier_char\x00__identifier_char_no_digit\x00__identifier_char_no_digit_sign\x00null\x00i8\x00i16\x00i32\x00i64\x00u8\x00u16\x00u32\x00u64\x00isize\x00usize\x00f32\x00f64\x00decimal64\x00decimal128\x00date-time\x00time\x00date\x00duration\x00decimal\x00currency\x00country-2\x00country-3\x00country-subdivision\x00email\x00idn-email\x00hostname\x00idn-hostname\x00ipv4\x00ipv6\x00url\x00url-reference\x00irl\x00iri-reference\x00url-template\x00uuid\x00regex\x00base64\x00=\x00(\x00)\x00\"\x00_escaped_string_token1\x00escape\x00_hex_digit\x00.\x00e\x00E\x00_\x00_digit\x00+\x00-\x000x\x000o\x00_octal_token1\x000b\x000\x001\x00true\x00false\x00\\\x00_newline_token1\x00_newline_token2\x00_newline_token3\x00_newline_token4\x00_newline_token5\x00_newline_token6\x00_newline_token7\x00_bom\x00_unicode_space\x00//\x00single_line_comment_token1\x00_eof\x00multi_line_comment\x00_raw_string\x00document\x00node\x00node_field\x00_node_field_comment\x00_node_field\x00node_children\x00_node_space\x00_node_terminator\x00identifier\x00_bare_identifier\x00keyword\x00annotation_type\x00prop\x00value\x00type\x00string\x00_escaped_string\x00number\x00_decimal\x00exponent\x00_integer\x00_sign\x00_hex\x00_octal\x00_binary\x00boolean\x00_escline\x00_linespace\x00_newline\x00_ws\x00single_line_comment\x00document_repeat1\x00document_repeat2\x00node_repeat1\x00node_repeat2\x00node_repeat3\x00_bare_identifier_repeat1\x00_escaped_string_repeat1\x00_integer_repeat1\x00_hex_repeat1\x00_octal_repeat1\x00_binary_repeat1\x00single_line_comment_repeat1\x00node_children_comment\x00node_field_comment\x00string_fragment\x00children\x00"
