// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-proto/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-proto -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-proto/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_proto

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
const LANGUAGE_VERSION = 13
const LARGE_STATE_COUNT = 2
const MAX_ALIAS_SEQUENCE_LENGTH = 14
const PRODUCTION_ID_COUNT = 3
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 288
const SYMBOL_COUNT = 119
const TOKEN_COUNT = 64
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

const anon_sym_SEMI = 1
const anon_sym_syntax = 2
const anon_sym_EQ = 3
const anon_sym_DQUOTEproto3_DQUOTE = 4
const anon_sym_import = 5
const anon_sym_weak = 6
const anon_sym_public = 7
const anon_sym_package = 8
const anon_sym_option = 9
const anon_sym_LPAREN = 10
const anon_sym_RPAREN = 11
const anon_sym_DOT = 12
const anon_sym_enum = 13
const anon_sym_LBRACE = 14
const anon_sym_RBRACE = 15
const anon_sym_DASH = 16
const anon_sym_LBRACK = 17
const anon_sym_COMMA = 18
const anon_sym_RBRACK = 19
const anon_sym_message = 20
const anon_sym_optional = 21
const anon_sym_repeated = 22
const anon_sym_oneof = 23
const anon_sym_map = 24
const anon_sym_LT = 25
const anon_sym_GT = 26
const anon_sym_int32 = 27
const anon_sym_int64 = 28
const anon_sym_uint32 = 29
const anon_sym_uint64 = 30
const anon_sym_sint32 = 31
const anon_sym_sint64 = 32
const anon_sym_fixed32 = 33
const anon_sym_fixed64 = 34
const anon_sym_sfixed32 = 35
const anon_sym_sfixed64 = 36
const anon_sym_bool = 37
const anon_sym_string = 38
const anon_sym_double = 39
const anon_sym_float = 40
const anon_sym_bytes = 41
const anon_sym_reserved = 42
const anon_sym_to = 43
const anon_sym_max = 44
const anon_sym_service = 45
const anon_sym_rpc = 46
const anon_sym_stream = 47
const anon_sym_returns = 48
const anon_sym_PLUS = 49
const anon_sym_COLON = 50
const sym_identifier = 51
const sym_true = 52
const sym_false = 53
const sym_decimal_lit = 54
const sym_octal_lit = 55
const sym_hex_lit = 56
const sym_float_lit = 57
const anon_sym_DQUOTE = 58
const aux_sym_string_token1 = 59
const anon_sym_SQUOTE = 60
const aux_sym_string_token2 = 61
const sym_escape_sequence = 62
const sym_comment = 63
const sym_source_file = 64
const sym_empty_statement = 65
const sym_syntax = 66
const sym_import = 67
const sym_package = 68
const sym_option = 69
const sym__option_name = 70
const sym_enum = 71
const sym_enum_name = 72
const sym_enum_body = 73
const sym_enum_field = 74
const sym_enum_value_option = 75
const sym_message = 76
const sym_message_body = 77
const sym_message_name = 78
const sym_field = 79
const sym_field_options = 80
const sym_field_option = 81
const sym_oneof = 82
const sym_oneof_field = 83
const sym_map_field = 84
const sym_key_type = 85
const sym_type = 86
const sym_reserved = 87
const sym_ranges = 88
const sym_range = 89
const sym_field_names = 90
const sym_message_or_enum_type = 91
const sym_field_number = 92
const sym_service = 93
const sym_service_name = 94
const sym_rpc = 95
const sym_rpc_name = 96
const sym_constant = 97
const sym_block_lit = 98
const sym_full_ident = 99
const sym_bool = 100
const sym_int_lit = 101
const sym_string = 102
const aux_sym_source_file_repeat1 = 103
const aux_sym__option_name_repeat1 = 104
const aux_sym_enum_body_repeat1 = 105
const aux_sym_enum_field_repeat1 = 106
const aux_sym_message_body_repeat1 = 107
const aux_sym_field_options_repeat1 = 108
const aux_sym_oneof_repeat1 = 109
const aux_sym_ranges_repeat1 = 110
const aux_sym_field_names_repeat1 = 111
const aux_sym_message_or_enum_type_repeat1 = 112
const aux_sym_service_repeat1 = 113
const aux_sym_rpc_repeat1 = 114
const aux_sym_block_lit_repeat1 = 115
const aux_sym_block_lit_repeat2 = 116
const aux_sym_string_repeat1 = 117
const aux_sym_string_repeat2 = 118

var ts_symbol_names = [119]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 6,
	3:   __ccgo_ts + 13,
	4:   __ccgo_ts + 15,
	5:   __ccgo_ts + 24,
	6:   __ccgo_ts + 31,
	7:   __ccgo_ts + 36,
	8:   __ccgo_ts + 43,
	9:   __ccgo_ts + 51,
	10:  __ccgo_ts + 58,
	11:  __ccgo_ts + 60,
	12:  __ccgo_ts + 62,
	13:  __ccgo_ts + 64,
	14:  __ccgo_ts + 69,
	15:  __ccgo_ts + 71,
	16:  __ccgo_ts + 73,
	17:  __ccgo_ts + 75,
	18:  __ccgo_ts + 77,
	19:  __ccgo_ts + 79,
	20:  __ccgo_ts + 81,
	21:  __ccgo_ts + 89,
	22:  __ccgo_ts + 98,
	23:  __ccgo_ts + 107,
	24:  __ccgo_ts + 113,
	25:  __ccgo_ts + 117,
	26:  __ccgo_ts + 119,
	27:  __ccgo_ts + 121,
	28:  __ccgo_ts + 127,
	29:  __ccgo_ts + 133,
	30:  __ccgo_ts + 140,
	31:  __ccgo_ts + 147,
	32:  __ccgo_ts + 154,
	33:  __ccgo_ts + 161,
	34:  __ccgo_ts + 169,
	35:  __ccgo_ts + 177,
	36:  __ccgo_ts + 186,
	37:  __ccgo_ts + 195,
	38:  __ccgo_ts + 200,
	39:  __ccgo_ts + 207,
	40:  __ccgo_ts + 214,
	41:  __ccgo_ts + 220,
	42:  __ccgo_ts + 226,
	43:  __ccgo_ts + 235,
	44:  __ccgo_ts + 238,
	45:  __ccgo_ts + 242,
	46:  __ccgo_ts + 250,
	47:  __ccgo_ts + 254,
	48:  __ccgo_ts + 261,
	49:  __ccgo_ts + 269,
	50:  __ccgo_ts + 271,
	51:  __ccgo_ts + 273,
	52:  __ccgo_ts + 284,
	53:  __ccgo_ts + 289,
	54:  __ccgo_ts + 295,
	55:  __ccgo_ts + 307,
	56:  __ccgo_ts + 317,
	57:  __ccgo_ts + 325,
	58:  __ccgo_ts + 335,
	59:  __ccgo_ts + 337,
	60:  __ccgo_ts + 351,
	61:  __ccgo_ts + 353,
	62:  __ccgo_ts + 367,
	63:  __ccgo_ts + 383,
	64:  __ccgo_ts + 391,
	65:  __ccgo_ts + 403,
	66:  __ccgo_ts + 6,
	67:  __ccgo_ts + 24,
	68:  __ccgo_ts + 43,
	69:  __ccgo_ts + 51,
	70:  __ccgo_ts + 419,
	71:  __ccgo_ts + 64,
	72:  __ccgo_ts + 432,
	73:  __ccgo_ts + 442,
	74:  __ccgo_ts + 452,
	75:  __ccgo_ts + 463,
	76:  __ccgo_ts + 81,
	77:  __ccgo_ts + 481,
	78:  __ccgo_ts + 494,
	79:  __ccgo_ts + 507,
	80:  __ccgo_ts + 513,
	81:  __ccgo_ts + 527,
	82:  __ccgo_ts + 107,
	83:  __ccgo_ts + 540,
	84:  __ccgo_ts + 552,
	85:  __ccgo_ts + 562,
	86:  __ccgo_ts + 571,
	87:  __ccgo_ts + 226,
	88:  __ccgo_ts + 576,
	89:  __ccgo_ts + 583,
	90:  __ccgo_ts + 589,
	91:  __ccgo_ts + 601,
	92:  __ccgo_ts + 622,
	93:  __ccgo_ts + 242,
	94:  __ccgo_ts + 635,
	95:  __ccgo_ts + 250,
	96:  __ccgo_ts + 648,
	97:  __ccgo_ts + 657,
	98:  __ccgo_ts + 666,
	99:  __ccgo_ts + 676,
	100: __ccgo_ts + 195,
	101: __ccgo_ts + 687,
	102: __ccgo_ts + 200,
	103: __ccgo_ts + 695,
	104: __ccgo_ts + 715,
	105: __ccgo_ts + 736,
	106: __ccgo_ts + 754,
	107: __ccgo_ts + 773,
	108: __ccgo_ts + 794,
	109: __ccgo_ts + 816,
	110: __ccgo_ts + 830,
	111: __ccgo_ts + 845,
	112: __ccgo_ts + 865,
	113: __ccgo_ts + 894,
	114: __ccgo_ts + 910,
	115: __ccgo_ts + 922,
	116: __ccgo_ts + 940,
	117: __ccgo_ts + 958,
	118: __ccgo_ts + 973,
}

var ts_symbol_map = [119]TSSymbol{
	1:   uint16(anon_sym_SEMI),
	2:   uint16(anon_sym_syntax),
	3:   uint16(anon_sym_EQ),
	4:   uint16(anon_sym_DQUOTEproto3_DQUOTE),
	5:   uint16(anon_sym_import),
	6:   uint16(anon_sym_weak),
	7:   uint16(anon_sym_public),
	8:   uint16(anon_sym_package),
	9:   uint16(anon_sym_option),
	10:  uint16(anon_sym_LPAREN),
	11:  uint16(anon_sym_RPAREN),
	12:  uint16(anon_sym_DOT),
	13:  uint16(anon_sym_enum),
	14:  uint16(anon_sym_LBRACE),
	15:  uint16(anon_sym_RBRACE),
	16:  uint16(anon_sym_DASH),
	17:  uint16(anon_sym_LBRACK),
	18:  uint16(anon_sym_COMMA),
	19:  uint16(anon_sym_RBRACK),
	20:  uint16(anon_sym_message),
	21:  uint16(anon_sym_optional),
	22:  uint16(anon_sym_repeated),
	23:  uint16(anon_sym_oneof),
	24:  uint16(anon_sym_map),
	25:  uint16(anon_sym_LT),
	26:  uint16(anon_sym_GT),
	27:  uint16(anon_sym_int32),
	28:  uint16(anon_sym_int64),
	29:  uint16(anon_sym_uint32),
	30:  uint16(anon_sym_uint64),
	31:  uint16(anon_sym_sint32),
	32:  uint16(anon_sym_sint64),
	33:  uint16(anon_sym_fixed32),
	34:  uint16(anon_sym_fixed64),
	35:  uint16(anon_sym_sfixed32),
	36:  uint16(anon_sym_sfixed64),
	37:  uint16(anon_sym_bool),
	38:  uint16(anon_sym_string),
	39:  uint16(anon_sym_double),
	40:  uint16(anon_sym_float),
	41:  uint16(anon_sym_bytes),
	42:  uint16(anon_sym_reserved),
	43:  uint16(anon_sym_to),
	44:  uint16(anon_sym_max),
	45:  uint16(anon_sym_service),
	46:  uint16(anon_sym_rpc),
	47:  uint16(anon_sym_stream),
	48:  uint16(anon_sym_returns),
	49:  uint16(anon_sym_PLUS),
	50:  uint16(anon_sym_COLON),
	51:  uint16(sym_identifier),
	52:  uint16(sym_true),
	53:  uint16(sym_false),
	54:  uint16(sym_decimal_lit),
	55:  uint16(sym_octal_lit),
	56:  uint16(sym_hex_lit),
	57:  uint16(sym_float_lit),
	58:  uint16(anon_sym_DQUOTE),
	59:  uint16(aux_sym_string_token1),
	60:  uint16(anon_sym_SQUOTE),
	61:  uint16(aux_sym_string_token2),
	62:  uint16(sym_escape_sequence),
	63:  uint16(sym_comment),
	64:  uint16(sym_source_file),
	65:  uint16(sym_empty_statement),
	66:  uint16(sym_syntax),
	67:  uint16(sym_import),
	68:  uint16(sym_package),
	69:  uint16(sym_option),
	70:  uint16(sym__option_name),
	71:  uint16(sym_enum),
	72:  uint16(sym_enum_name),
	73:  uint16(sym_enum_body),
	74:  uint16(sym_enum_field),
	75:  uint16(sym_enum_value_option),
	76:  uint16(sym_message),
	77:  uint16(sym_message_body),
	78:  uint16(sym_message_name),
	79:  uint16(sym_field),
	80:  uint16(sym_field_options),
	81:  uint16(sym_field_option),
	82:  uint16(sym_oneof),
	83:  uint16(sym_oneof_field),
	84:  uint16(sym_map_field),
	85:  uint16(sym_key_type),
	86:  uint16(sym_type),
	87:  uint16(sym_reserved),
	88:  uint16(sym_ranges),
	89:  uint16(sym_range),
	90:  uint16(sym_field_names),
	91:  uint16(sym_message_or_enum_type),
	92:  uint16(sym_field_number),
	93:  uint16(sym_service),
	94:  uint16(sym_service_name),
	95:  uint16(sym_rpc),
	96:  uint16(sym_rpc_name),
	97:  uint16(sym_constant),
	98:  uint16(sym_block_lit),
	99:  uint16(sym_full_ident),
	100: uint16(sym_bool),
	101: uint16(sym_int_lit),
	102: uint16(sym_string),
	103: uint16(aux_sym_source_file_repeat1),
	104: uint16(aux_sym__option_name_repeat1),
	105: uint16(aux_sym_enum_body_repeat1),
	106: uint16(aux_sym_enum_field_repeat1),
	107: uint16(aux_sym_message_body_repeat1),
	108: uint16(aux_sym_field_options_repeat1),
	109: uint16(aux_sym_oneof_repeat1),
	110: uint16(aux_sym_ranges_repeat1),
	111: uint16(aux_sym_field_names_repeat1),
	112: uint16(aux_sym_message_or_enum_type_repeat1),
	113: uint16(aux_sym_service_repeat1),
	114: uint16(aux_sym_rpc_repeat1),
	115: uint16(aux_sym_block_lit_repeat1),
	116: uint16(aux_sym_block_lit_repeat2),
	117: uint16(aux_sym_string_repeat1),
	118: uint16(aux_sym_string_repeat2),
}

var ts_symbol_metadata = [119]TSSymbolMetadata{
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	},
	59: {},
	60: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	61: {},
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
}

const field_path = 1

var ts_field_names = [2]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 988,
}

var ts_field_map_slices = [3]TSFieldMapSlice{
	1: {
		Flength: uint16(1),
	},
	2: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
}

var ts_field_map_entries = [2]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_path),
		Fchild_index: uint8(1),
	},
	1: {
		Ffield_id:    uint16(field_path),
		Fchild_index: uint8(2),
	},
}

var ts_alias_sequences = [3][14]TSSymbol{}

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
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(368)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(375)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(185)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(186)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(248)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(195)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(193)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(188)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(360)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(249)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(173)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(207)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(175)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(208)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(194)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(196)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(114)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(110)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(111)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(191)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(192)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(170)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(358)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('"') {
			state = uint16(368)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(375)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(248)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(193)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(159)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(360)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(249)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(194)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(350)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(311)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(351)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(325)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(191)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(358)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('"') {
			state = uint16(368)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(370)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			state = uint16(373)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(374)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('"') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('"') {
			state = uint16(124)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(193)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(187)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(362)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(173)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(175)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(269)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(315)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(310)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(267)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(309)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(266)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(268)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(276)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(298)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(192)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(359)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('\'') {
			state = uint16(375)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(377)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			state = uint16(380)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(381)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('(') {
			state = uint16(185)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(186)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(195)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(187)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(362)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(173)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(175)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(208)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(196)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(192)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(359)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('*') {
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(386)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('*') {
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(385)
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
		if lookahead == int32('.') {
			state = uint16(366)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(158)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('.') {
			state = uint16(187)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(173)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(194)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(269)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(315)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(267)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(309)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(323)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(298)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(192)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('.') {
			state = uint16(187)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(269)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(315)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(267)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(309)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(284)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(298)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('.') {
			state = uint16(187)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(269)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(315)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(267)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(309)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(298)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(13)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('.') {
			state = uint16(187)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(336)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(14)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(362)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(359)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(173)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(323)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(192)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(16)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('2') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('2') {
			state = uint16(217)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('2') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('2') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('2') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('3') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('3') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('3') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('3') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('3') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('3') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('4') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('4') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('4') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('4') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('4') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('U') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(165)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(163)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(384)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(382)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('a') {
			state = uint16(93)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(156)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('a') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('a') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('a') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('a') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('a') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('a') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('a') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('a') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('a') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('a') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('a') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('a') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('b') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('b') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('c') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('c') {
			state = uint16(244)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('c') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('c') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('d') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('d') {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('d') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('d') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('e') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('e') {
			state = uint16(125)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('e') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('e') {
			state = uint16(127)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(87)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(106)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(128)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('e') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('e') {
			state = uint16(354)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('e') {
			state = uint16(356)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('e') {
			state = uint16(233)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('e') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('e') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('e') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('e') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('e') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('e') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('e') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('e') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('e') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('e') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('e') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('e') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('f') {
			state = uint16(365)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('f') {
			state = uint16(365)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead == int32('f') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead == int32('g') {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead == int32('g') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead == int32('g') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('i') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('i') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('i') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('i') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead == int32('i') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead == int32('i') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead == int32('k') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('k') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('l') {
			state = uint16(229)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('l') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('l') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead == int32('l') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead == int32('l') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead == int32('m') {
			state = uint16(123)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead == int32('m') {
			state = uint16(123)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead == int32('m') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead == int32('m') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead == int32('n') {
			state = uint16(365)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('n') {
			state = uint16(182)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead == int32('n') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead == int32('n') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead == int32('n') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead == int32('n') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead == int32('n') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead == int32('n') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead == int32('n') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('n') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead == int32('o') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead == int32('o') {
			state = uint16(241)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(112):
		if lookahead == int32('o') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead == int32('o') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('o') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('o') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead == int32('o') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead == int32('o') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead == int32('o') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('o') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('o') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('p') {
			state = uint16(205)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('p') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead == int32('p') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(124):
		if lookahead == int32('p') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead == int32('p') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(75)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead == int32('p') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead == int32('r') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead == int32('r') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(129):
		if lookahead == int32('r') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(130):
		if lookahead == int32('r') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(131):
		if lookahead == int32('r') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(132):
		if lookahead == int32('r') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(133):
		if lookahead == int32('s') {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(134):
		if lookahead == int32('s') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead == int32('s') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(136):
		if lookahead == int32('s') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(137):
		if lookahead == int32('s') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(138):
		if lookahead == int32('t') {
			state = uint16(235)
			goto next_state
		}
		return result
	case int32(139):
		if lookahead == int32('t') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(140):
		if lookahead == int32('t') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead == int32('t') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(142):
		if lookahead == int32('t') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(143):
		if lookahead == int32('t') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(144):
		if lookahead == int32('t') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(145):
		if lookahead == int32('t') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(146):
		if lookahead == int32('t') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(147):
		if lookahead == int32('t') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(148):
		if lookahead == int32('u') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(149):
		if lookahead == int32('u') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(150):
		if lookahead == int32('u') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(151):
		if lookahead == int32('u') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(152):
		if lookahead == int32('v') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead == int32('v') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(154):
		if lookahead == int32('x') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(155):
		if lookahead == int32('x') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(156):
		if lookahead == int32('x') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(157):
		if lookahead == int32('x') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(158):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(160)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(367)
			goto next_state
		}
		return result
	case int32(159):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(366)
			goto next_state
		}
		return result
	case int32(160):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(367)
			goto next_state
		}
		return result
	case int32(161):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(382)
			goto next_state
		}
		return result
	case int32(162):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(364)
			goto next_state
		}
		return result
	case int32(163):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(164):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(163)
			goto next_state
		}
		return result
	case int32(165):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(166):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(167):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(166)
			goto next_state
		}
		return result
	case int32(168):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(169):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(170):
		if eof != 0 {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(368)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(375)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(185)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(186)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(248)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(195)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(193)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(188)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(360)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(249)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(173)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(207)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(175)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(208)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(194)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(196)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(114)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(110)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(111)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(191)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(192)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(170)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(358)
			goto next_state
		}
		return result
	case int32(171):
		if eof != 0 {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(159)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(360)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(173)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(126)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(122)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(59)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(192)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(171)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(358)
			goto next_state
		}
		return result
	case int32(172):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(173):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(174):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_syntax)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(175):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(176):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTEproto3_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(177):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_import)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(178):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_weak)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(179):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_public)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(180):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_package)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(181):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_option)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(182):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_option)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(183):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_option)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(184):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_option)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(185):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(186):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(187):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(188):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(366)
			goto next_state
		}
		return result
	case int32(189):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(190):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(191):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(192):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(193):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(194):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(195):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(196):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(197):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_message)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(198):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_message)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(199):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_optional)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(200):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_optional)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(201):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_repeated)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(202):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_repeated)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(203):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_oneof)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(204):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_oneof)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(205):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_map)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(206):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_map)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(207):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(208):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(209):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_int32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(210):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_int32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(211):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_int64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(212):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_int64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(213):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_uint32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(214):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_uint32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(215):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_uint64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(216):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_uint64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(217):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sint32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(218):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sint32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(219):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sint64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(220):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sint64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(221):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_fixed32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(222):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_fixed32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(223):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_fixed64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(224):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_fixed64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(225):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sfixed32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(226):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sfixed32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(227):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sfixed64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(228):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sfixed64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(229):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bool)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(230):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bool)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(231):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(232):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(233):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_double)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(234):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_double)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(235):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(236):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(237):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bytes)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(238):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bytes)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(239):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_reserved)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(240):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_reserved)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(241):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_to)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(242):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_max)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(243):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_service)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(244):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rpc)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(245):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_stream)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(246):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_stream)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(247):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_returns)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(248):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(249):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(250):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(210)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(251):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(218)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(252):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(214)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(253):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(222)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(254):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(226)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(255):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('3') {
			state = uint16(250)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(260)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(256):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('3') {
			state = uint16(251)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(261)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(257):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('3') {
			state = uint16(252)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(262)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(258):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('3') {
			state = uint16(253)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(263)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(259):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('3') {
			state = uint16(254)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(264)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(260):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('4') {
			state = uint16(212)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(261):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('4') {
			state = uint16(220)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(262):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('4') {
			state = uint16(216)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(263):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('4') {
			state = uint16(224)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(264):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('4') {
			state = uint16(228)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(265):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('e') || lookahead == int32('g') || lookahead == int32('h') || int32('j') <= lookahead && lookahead <= int32('s') || int32('u') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(299)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(313)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(327)
			goto next_state
		}
		return result
	case int32(266):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(321)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(329)
			goto next_state
		}
		return result
	case int32(267):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('h') || lookahead == int32('j') || lookahead == int32('k') || int32('m') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(317)
			goto next_state
		}
		return result
	case int32(268):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('m') || lookahead == int32('o') || int32('q') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(288)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(335)
			goto next_state
		}
		return result
	case int32(269):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('n') || int32('p') <= lookahead && lookahead <= int32('x') || lookahead == int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(316)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(334)
			goto next_state
		}
		return result
	case int32(270):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('o') || lookahead == int32('q') || lookahead == int32('r') || int32('t') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(286)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(282)
			goto next_state
		}
		return result
	case int32(271):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || lookahead == int32('a') || int32('c') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(272):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('c') || int32('e') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(273):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('c') || int32('e') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(240)
			goto next_state
		}
		return result
	case int32(274):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('c') || int32('e') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(275):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('c') || int32('e') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(259)
			goto next_state
		}
		return result
	case int32(276):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(277):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(274)
			goto next_state
		}
		return result
	case int32(278):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(279):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(280):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(281):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(357)
			goto next_state
		}
		return result
	case int32(282):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(324)
			goto next_state
		}
		return result
	case int32(283):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(272)
			goto next_state
		}
		return result
	case int32(284):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(322)
			goto next_state
		}
		return result
	case int32(285):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(328)
			goto next_state
		}
		return result
	case int32(286):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(352)
			goto next_state
		}
		return result
	case int32(287):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(288):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(318)
			goto next_state
		}
		return result
	case int32(289):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(348)
			goto next_state
		}
		return result
	case int32(290):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('d') || int32('f') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(291):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('e') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(292):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('e') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(204)
			goto next_state
		}
		return result
	case int32(293):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('f') || int32('h') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('g') {
			state = uint16(232)
			goto next_state
		}
		return result
	case int32(294):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('f') || int32('h') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('g') {
			state = uint16(279)
			goto next_state
		}
		return result
	case int32(295):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('h') || int32('j') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(312)
			goto next_state
		}
		return result
	case int32(296):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('h') || int32('j') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(319)
			goto next_state
		}
		return result
	case int32(297):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('h') || int32('j') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(320)
			goto next_state
		}
		return result
	case int32(298):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('h') || int32('j') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(314)
			goto next_state
		}
		return result
	case int32(299):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('h') || int32('j') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(300):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('k') || int32('m') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(230)
			goto next_state
		}
		return result
	case int32(301):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('k') || int32('m') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(302):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('k') || int32('m') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(331)
			goto next_state
		}
		return result
	case int32(303):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('k') || int32('m') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(304):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('l') || int32('n') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(305):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('l') || int32('n') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(306):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('m') || int32('o') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(307):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('m') || int32('o') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(308):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('m') || int32('o') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(309):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('m') || int32('o') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(332)
			goto next_state
		}
		return result
	case int32(310):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('m') || int32('o') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(311):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('m') || int32('o') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(291)
			goto next_state
		}
		return result
	case int32(312):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('m') || int32('o') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(293)
			goto next_state
		}
		return result
	case int32(313):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('m') || int32('o') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(338)
			goto next_state
		}
		return result
	case int32(314):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('m') || int32('o') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(340)
			goto next_state
		}
		return result
	case int32(315):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('n') || int32('p') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(341)
			goto next_state
		}
		return result
	case int32(316):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('n') || int32('p') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(300)
			goto next_state
		}
		return result
	case int32(317):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('n') || int32('p') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(349)
			goto next_state
		}
		return result
	case int32(318):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('n') || int32('p') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(319):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('n') || int32('p') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(307)
			goto next_state
		}
		return result
	case int32(320):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('n') || int32('p') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(308)
			goto next_state
		}
		return result
	case int32(321):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('o') || int32('q') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(322):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('o') || int32('q') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(286)
			goto next_state
		}
		return result
	case int32(323):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('o') || int32('q') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(339)
			goto next_state
		}
		return result
	case int32(324):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('q') || int32('s') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(344)
			goto next_state
		}
		return result
	case int32(325):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('q') || int32('s') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(343)
			goto next_state
		}
		return result
	case int32(326):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('q') || int32('s') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(289)
			goto next_state
		}
		return result
	case int32(327):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('q') || int32('s') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(295)
			goto next_state
		}
		return result
	case int32(328):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('r') || int32('t') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(329):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('r') || int32('t') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(330)
			goto next_state
		}
		return result
	case int32(330):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('r') || int32('t') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(347)
			goto next_state
		}
		return result
	case int32(331):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('r') || int32('t') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(281)
			goto next_state
		}
		return result
	case int32(332):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('s') || int32('u') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(333):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('s') || int32('u') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(236)
			goto next_state
		}
		return result
	case int32(334):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('s') || int32('u') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(285)
			goto next_state
		}
		return result
	case int32(335):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('s') || int32('u') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(336):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('s') || int32('u') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(326)
			goto next_state
		}
		return result
	case int32(337):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('s') || int32('u') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(283)
			goto next_state
		}
		return result
	case int32(338):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('s') || int32('u') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(339):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('s') || int32('u') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(297)
			goto next_state
		}
		return result
	case int32(340):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('s') || int32('u') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(341):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('t') || int32('v') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(342):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('t') || int32('v') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(304)
			goto next_state
		}
		return result
	case int32(343):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('t') || int32('v') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(280)
			goto next_state
		}
		return result
	case int32(344):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('u') || int32('w') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('v') {
			state = uint16(287)
			goto next_state
		}
		return result
	case int32(345):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('w') || lookahead == int32('y') || lookahead == int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(346):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('w') || lookahead == int32('y') || lookahead == int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(290)
			goto next_state
		}
		return result
	case int32(347):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(294)
			goto next_state
		}
		return result
	case int32(348):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(305)
			goto next_state
		}
		return result
	case int32(349):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(333)
			goto next_state
		}
		return result
	case int32(350):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(302)
			goto next_state
		}
		return result
	case int32(351):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(306)
			goto next_state
		}
		return result
	case int32(352):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(337)
			goto next_state
		}
		return result
	case int32(353):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(354):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(355):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(356):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(357):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(358):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_decimal_lit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(366)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(158)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(358)
			goto next_state
		}
		return result
	case int32(359):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_decimal_lit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(359)
			goto next_state
		}
		return result
	case int32(360):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_octal_lit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(366)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(158)
			goto next_state
		}
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('8') || lookahead == int32('9') {
			state = uint16(10)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(361)
			goto next_state
		}
		return result
	case int32(361):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_octal_lit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(366)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(158)
			goto next_state
		}
		if lookahead == int32('8') || lookahead == int32('9') {
			state = uint16(10)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(361)
			goto next_state
		}
		return result
	case int32(362):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_octal_lit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(162)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(363)
			goto next_state
		}
		return result
	case int32(363):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_octal_lit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(363)
			goto next_state
		}
		return result
	case int32(364):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_hex_lit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(364)
			goto next_state
		}
		return result
	case int32(365):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float_lit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(366):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float_lit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(158)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(366)
			goto next_state
		}
		return result
	case int32(367):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float_lit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(367)
			goto next_state
		}
		return result
	case int32(368):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(369):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(374)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(369)
			goto next_state
		}
		return result
	case int32(370):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(372)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(369)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(374)
			goto next_state
		}
		return result
	case int32(371):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(371)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(374)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(372)
			goto next_state
		}
		return result
	case int32(372):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(371)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(372)
			goto next_state
		}
		return result
	case int32(373):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(370)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			state = uint16(373)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(374)
			goto next_state
		}
		return result
	case int32(374):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(374)
			goto next_state
		}
		return result
	case int32(375):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(376):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(381)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(376)
			goto next_state
		}
		return result
	case int32(377):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(379)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(376)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(381)
			goto next_state
		}
		return result
	case int32(378):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(378)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(381)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(379)
			goto next_state
		}
		return result
	case int32(379):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(378)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(379)
			goto next_state
		}
		return result
	case int32(380):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(377)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			state = uint16(380)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(381)
			goto next_state
		}
		return result
	case int32(381):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(381)
			goto next_state
		}
		return result
	case int32(382):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(383):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(382)
			goto next_state
		}
		return result
	case int32(384):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(383)
			goto next_state
		}
		return result
	case int32(385):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(386):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(386)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [288]TSLexMode{
	0: {},
	1: {},
	2: {
		Flex_state: uint16(4),
	},
	3: {
		Flex_state: uint16(4),
	},
	4: {
		Flex_state: uint16(4),
	},
	5: {
		Flex_state: uint16(4),
	},
	6: {
		Flex_state: uint16(4),
	},
	7: {
		Flex_state: uint16(4),
	},
	8: {
		Flex_state: uint16(4),
	},
	9: {
		Flex_state: uint16(11),
	},
	10: {
		Flex_state: uint16(4),
	},
	11: {
		Flex_state: uint16(4),
	},
	12: {
		Flex_state: uint16(4),
	},
	13: {
		Flex_state: uint16(4),
	},
	14: {
		Flex_state: uint16(4),
	},
	15: {
		Flex_state: uint16(4),
	},
	16: {
		Flex_state: uint16(4),
	},
	17: {
		Flex_state: uint16(4),
	},
	18: {
		Flex_state: uint16(4),
	},
	19: {
		Flex_state: uint16(4),
	},
	20: {
		Flex_state: uint16(4),
	},
	21: {
		Flex_state: uint16(4),
	},
	22: {
		Flex_state: uint16(4),
	},
	23: {
		Flex_state: uint16(4),
	},
	24: {
		Flex_state: uint16(11),
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
		Flex_state: uint16(11),
	},
	29: {
		Flex_state: uint16(11),
	},
	30: {
		Flex_state: uint16(12),
	},
	31: {
		Flex_state: uint16(11),
	},
	32: {
		Flex_state: uint16(11),
	},
	33: {
		Flex_state: uint16(1),
	},
	34: {
		Flex_state: uint16(13),
	},
	35: {
		Flex_state: uint16(11),
	},
	36: {
		Flex_state: uint16(11),
	},
	37: {
		Flex_state: uint16(13),
	},
	38: {
		Flex_state: uint16(11),
	},
	39: {
		Flex_state: uint16(13),
	},
	40: {
		Flex_state: uint16(1),
	},
	41: {
		Flex_state: uint16(1),
	},
	42: {
		Flex_state: uint16(1),
	},
	43: {
		Flex_state: uint16(1),
	},
	44: {
		Flex_state: uint16(1),
	},
	45: {
		Flex_state: uint16(1),
	},
	46: {
		Flex_state: uint16(1),
	},
	47: {
		Flex_state: uint16(1),
	},
	48: {
		Flex_state: uint16(1),
	},
	49: {
		Flex_state: uint16(1),
	},
	50: {
		Flex_state: uint16(171),
	},
	51: {
		Flex_state: uint16(171),
	},
	52: {
		Flex_state: uint16(171),
	},
	53: {},
	54: {
		Flex_state: uint16(171),
	},
	55: {
		Flex_state: uint16(171),
	},
	56: {
		Flex_state: uint16(6),
	},
	57: {
		Flex_state: uint16(6),
	},
	58: {
		Flex_state: uint16(16),
	},
	59: {
		Flex_state: uint16(171),
	},
	60: {
		Flex_state: uint16(171),
	},
	61: {
		Flex_state: uint16(16),
	},
	62: {
		Flex_state: uint16(171),
	},
	63: {
		Flex_state: uint16(171),
	},
	64: {
		Flex_state: uint16(171),
	},
	65: {
		Flex_state: uint16(171),
	},
	66: {
		Flex_state: uint16(171),
	},
	67: {
		Flex_state: uint16(6),
	},
	68: {
		Flex_state: uint16(171),
	},
	69: {
		Flex_state: uint16(171),
	},
	70: {
		Flex_state: uint16(171),
	},
	71: {
		Flex_state: uint16(171),
	},
	72: {
		Flex_state: uint16(171),
	},
	73: {
		Flex_state: uint16(171),
	},
	74: {
		Flex_state: uint16(6),
	},
	75: {
		Flex_state: uint16(171),
	},
	76: {
		Flex_state: uint16(171),
	},
	77: {
		Flex_state: uint16(16),
	},
	78: {
		Flex_state: uint16(16),
	},
	79: {
		Flex_state: uint16(6),
	},
	80: {
		Flex_state: uint16(16),
	},
	81: {
		Flex_state: uint16(171),
	},
	82: {
		Flex_state: uint16(171),
	},
	83: {
		Flex_state: uint16(171),
	},
	84: {
		Flex_state: uint16(171),
	},
	85: {
		Flex_state: uint16(171),
	},
	86: {
		Flex_state: uint16(171),
	},
	87: {
		Flex_state: uint16(171),
	},
	88: {
		Flex_state: uint16(171),
	},
	89: {
		Flex_state: uint16(6),
	},
	90: {
		Flex_state: uint16(6),
	},
	91: {
		Flex_state: uint16(6),
	},
	92: {
		Flex_state: uint16(6),
	},
	93: {
		Flex_state: uint16(4),
	},
	94: {
		Flex_state: uint16(6),
	},
	95: {
		Flex_state: uint16(4),
	},
	96: {
		Flex_state: uint16(6),
	},
	97: {
		Flex_state: uint16(4),
	},
	98: {
		Flex_state: uint16(6),
	},
	99: {
		Flex_state: uint16(4),
	},
	100: {
		Flex_state: uint16(14),
	},
	101: {},
	102: {},
	103: {
		Flex_state: uint16(4),
	},
	104: {
		Flex_state: uint16(6),
	},
	105: {
		Flex_state: uint16(4),
	},
	106: {
		Flex_state: uint16(14),
	},
	107: {
		Flex_state: uint16(6),
	},
	108: {
		Flex_state: uint16(6),
	},
	109: {
		Flex_state: uint16(6),
	},
	110: {
		Flex_state: uint16(6),
	},
	111: {
		Flex_state: uint16(171),
	},
	112: {
		Flex_state: uint16(4),
	},
	113: {
		Flex_state: uint16(14),
	},
	114: {
		Flex_state: uint16(15),
	},
	115: {
		Flex_state: uint16(6),
	},
	116: {
		Flex_state: uint16(6),
	},
	117: {
		Flex_state: uint16(2),
	},
	118: {
		Flex_state: uint16(4),
	},
	119: {
		Flex_state: uint16(2),
	},
	120: {
		Flex_state: uint16(6),
	},
	121: {
		Flex_state: uint16(171),
	},
	122: {
		Flex_state: uint16(6),
	},
	123: {
		Flex_state: uint16(171),
	},
	124: {
		Flex_state: uint16(5),
	},
	125: {
		Flex_state: uint16(6),
	},
	126: {
		Flex_state: uint16(6),
	},
	127: {
		Flex_state: uint16(16),
	},
	128: {
		Flex_state: uint16(6),
	},
	129: {
		Flex_state: uint16(171),
	},
	130: {
		Flex_state: uint16(6),
	},
	131: {
		Flex_state: uint16(6),
	},
	132: {
		Flex_state: uint16(16),
	},
	133: {
		Flex_state: uint16(16),
	},
	134: {
		Flex_state: uint16(16),
	},
	135: {
		Flex_state: uint16(6),
	},
	136: {
		Flex_state: uint16(5),
	},
	137: {
		Flex_state: uint16(6),
	},
	138: {
		Flex_state: uint16(5),
	},
	139: {
		Flex_state: uint16(16),
	},
	140: {
		Flex_state: uint16(6),
	},
	141: {
		Flex_state: uint16(2),
	},
	142: {
		Flex_state: uint16(6),
	},
	143: {
		Flex_state: uint16(6),
	},
	144: {
		Flex_state: uint16(171),
	},
	145: {
		Flex_state: uint16(6),
	},
	146: {
		Flex_state: uint16(16),
	},
	147: {
		Flex_state: uint16(16),
	},
	148: {
		Flex_state: uint16(171),
	},
	149: {
		Flex_state: uint16(6),
	},
	150: {
		Flex_state: uint16(6),
	},
	151: {},
	152: {},
	153: {},
	154: {},
	155: {},
	156: {
		Flex_state: uint16(4),
	},
	157: {},
	158: {},
	159: {},
	160: {
		Flex_state: uint16(4),
	},
	161: {},
	162: {},
	163: {
		Flex_state: uint16(6),
	},
	164: {
		Flex_state: uint16(4),
	},
	165: {},
	166: {},
	167: {},
	168: {},
	169: {},
	170: {},
	171: {},
	172: {},
	173: {},
	174: {
		Flex_state: uint16(4),
	},
	175: {
		Flex_state: uint16(6),
	},
	176: {},
	177: {
		Flex_state: uint16(6),
	},
	178: {
		Flex_state: uint16(6),
	},
	179: {},
	180: {
		Flex_state: uint16(6),
	},
	181: {
		Flex_state: uint16(6),
	},
	182: {},
	183: {
		Flex_state: uint16(6),
	},
	184: {},
	185: {},
	186: {},
	187: {
		Flex_state: uint16(6),
	},
	188: {},
	189: {},
	190: {
		Flex_state: uint16(6),
	},
	191: {
		Flex_state: uint16(6),
	},
	192: {},
	193: {
		Flex_state: uint16(6),
	},
	194: {},
	195: {},
	196: {},
	197: {},
	198: {
		Flex_state: uint16(6),
	},
	199: {
		Flex_state: uint16(6),
	},
	200: {},
	201: {},
	202: {},
	203: {
		Flex_state: uint16(6),
	},
	204: {},
	205: {},
	206: {
		Flex_state: uint16(6),
	},
	207: {
		Flex_state: uint16(6),
	},
	208: {
		Flex_state: uint16(6),
	},
	209: {},
	210: {},
	211: {
		Flex_state: uint16(6),
	},
	212: {
		Flex_state: uint16(6),
	},
	213: {
		Flex_state: uint16(6),
	},
	214: {},
	215: {},
	216: {
		Flex_state: uint16(6),
	},
	217: {},
	218: {
		Flex_state: uint16(6),
	},
	219: {
		Flex_state: uint16(6),
	},
	220: {
		Flex_state: uint16(6),
	},
	221: {},
	222: {},
	223: {},
	224: {},
	225: {
		Flex_state: uint16(6),
	},
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
	236: {
		Flex_state: uint16(6),
	},
	237: {},
	238: {},
	239: {},
	240: {
		Flex_state: uint16(6),
	},
	241: {},
	242: {
		Flex_state: uint16(6),
	},
	243: {},
	244: {},
	245: {},
	246: {},
	247: {},
	248: {},
	249: {},
	250: {},
	251: {},
	252: {
		Flex_state: uint16(6),
	},
	253: {},
	254: {},
	255: {},
	256: {},
	257: {},
	258: {},
	259: {
		Flex_state: uint16(6),
	},
	260: {},
	261: {},
	262: {},
	263: {},
	264: {},
	265: {},
	266: {},
	267: {},
	268: {
		Flex_state: uint16(6),
	},
	269: {},
	270: {},
	271: {
		Flex_state: uint16(6),
	},
	272: {},
	273: {},
	274: {},
	275: {},
	276: {},
	277: {},
	278: {
		Flex_state: uint16(4),
	},
	279: {
		Flex_state: uint16(6),
	},
	280: {
		Flex_state: uint16(4),
	},
	281: {},
	282: {},
	283: {},
	284: {},
	285: {},
	286: {},
	287: {},
}

var ts_parse_table = [2][119]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		3:  uint16(1),
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
		52: uint16(1),
		53: uint16(1),
		54: uint16(1),
		55: uint16(1),
		56: uint16(1),
		57: uint16(1),
		58: uint16(1),
		60: uint16(1),
		62: uint16(1),
		63: uint16(3),
	},
	1: {
		2:  uint16(5),
		63: uint16(3),
		64: uint16(281),
		66: uint16(50),
	},
}

var ts_small_parse_table = [4868]uint16_t{
	0:    uint16(18),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(7),
	5:    uint16(1),
	6:    uint16(anon_sym_SEMI),
	7:    uint16(9),
	8:    uint16(1),
	9:    uint16(anon_sym_option),
	10:   uint16(11),
	11:   uint16(1),
	12:   uint16(anon_sym_DOT),
	13:   uint16(13),
	14:   uint16(1),
	15:   uint16(anon_sym_enum),
	16:   uint16(15),
	17:   uint16(1),
	18:   uint16(anon_sym_RBRACE),
	19:   uint16(17),
	20:   uint16(1),
	21:   uint16(anon_sym_message),
	22:   uint16(19),
	23:   uint16(1),
	24:   uint16(anon_sym_optional),
	25:   uint16(21),
	26:   uint16(1),
	27:   uint16(anon_sym_repeated),
	28:   uint16(23),
	29:   uint16(1),
	30:   uint16(anon_sym_oneof),
	31:   uint16(25),
	32:   uint16(1),
	33:   uint16(anon_sym_map),
	34:   uint16(29),
	35:   uint16(1),
	36:   uint16(anon_sym_reserved),
	37:   uint16(31),
	38:   uint16(1),
	39:   uint16(sym_identifier),
	40:   uint16(199),
	41:   uint16(1),
	42:   uint16(aux_sym_message_or_enum_type_repeat1),
	43:   uint16(203),
	44:   uint16(1),
	45:   uint16(sym_message_or_enum_type),
	46:   uint16(240),
	47:   uint16(1),
	48:   uint16(sym_type),
	49:   uint16(6),
	50:   uint16(9),
	51:   uint16(sym_empty_statement),
	52:   uint16(sym_option),
	53:   uint16(sym_enum),
	54:   uint16(sym_message),
	55:   uint16(sym_field),
	56:   uint16(sym_oneof),
	57:   uint16(sym_map_field),
	58:   uint16(sym_reserved),
	59:   uint16(aux_sym_message_body_repeat1),
	60:   uint16(27),
	61:   uint16(15),
	62:   uint16(anon_sym_int32),
	63:   uint16(anon_sym_int64),
	64:   uint16(anon_sym_uint32),
	65:   uint16(anon_sym_uint64),
	66:   uint16(anon_sym_sint32),
	67:   uint16(anon_sym_sint64),
	68:   uint16(anon_sym_fixed32),
	69:   uint16(anon_sym_fixed64),
	70:   uint16(anon_sym_sfixed32),
	71:   uint16(anon_sym_sfixed64),
	72:   uint16(anon_sym_bool),
	73:   uint16(anon_sym_string),
	74:   uint16(anon_sym_double),
	75:   uint16(anon_sym_float),
	76:   uint16(anon_sym_bytes),
	77:   uint16(18),
	78:   uint16(3),
	79:   uint16(1),
	80:   uint16(sym_comment),
	81:   uint16(7),
	82:   uint16(1),
	83:   uint16(anon_sym_SEMI),
	84:   uint16(9),
	85:   uint16(1),
	86:   uint16(anon_sym_option),
	87:   uint16(11),
	88:   uint16(1),
	89:   uint16(anon_sym_DOT),
	90:   uint16(13),
	91:   uint16(1),
	92:   uint16(anon_sym_enum),
	93:   uint16(17),
	94:   uint16(1),
	95:   uint16(anon_sym_message),
	96:   uint16(19),
	97:   uint16(1),
	98:   uint16(anon_sym_optional),
	99:   uint16(21),
	100:  uint16(1),
	101:  uint16(anon_sym_repeated),
	102:  uint16(23),
	103:  uint16(1),
	104:  uint16(anon_sym_oneof),
	105:  uint16(25),
	106:  uint16(1),
	107:  uint16(anon_sym_map),
	108:  uint16(29),
	109:  uint16(1),
	110:  uint16(anon_sym_reserved),
	111:  uint16(31),
	112:  uint16(1),
	113:  uint16(sym_identifier),
	114:  uint16(33),
	115:  uint16(1),
	116:  uint16(anon_sym_RBRACE),
	117:  uint16(199),
	118:  uint16(1),
	119:  uint16(aux_sym_message_or_enum_type_repeat1),
	120:  uint16(203),
	121:  uint16(1),
	122:  uint16(sym_message_or_enum_type),
	123:  uint16(240),
	124:  uint16(1),
	125:  uint16(sym_type),
	126:  uint16(5),
	127:  uint16(9),
	128:  uint16(sym_empty_statement),
	129:  uint16(sym_option),
	130:  uint16(sym_enum),
	131:  uint16(sym_message),
	132:  uint16(sym_field),
	133:  uint16(sym_oneof),
	134:  uint16(sym_map_field),
	135:  uint16(sym_reserved),
	136:  uint16(aux_sym_message_body_repeat1),
	137:  uint16(27),
	138:  uint16(15),
	139:  uint16(anon_sym_int32),
	140:  uint16(anon_sym_int64),
	141:  uint16(anon_sym_uint32),
	142:  uint16(anon_sym_uint64),
	143:  uint16(anon_sym_sint32),
	144:  uint16(anon_sym_sint64),
	145:  uint16(anon_sym_fixed32),
	146:  uint16(anon_sym_fixed64),
	147:  uint16(anon_sym_sfixed32),
	148:  uint16(anon_sym_sfixed64),
	149:  uint16(anon_sym_bool),
	150:  uint16(anon_sym_string),
	151:  uint16(anon_sym_double),
	152:  uint16(anon_sym_float),
	153:  uint16(anon_sym_bytes),
	154:  uint16(18),
	155:  uint16(3),
	156:  uint16(1),
	157:  uint16(sym_comment),
	158:  uint16(35),
	159:  uint16(1),
	160:  uint16(anon_sym_SEMI),
	161:  uint16(38),
	162:  uint16(1),
	163:  uint16(anon_sym_option),
	164:  uint16(41),
	165:  uint16(1),
	166:  uint16(anon_sym_DOT),
	167:  uint16(44),
	168:  uint16(1),
	169:  uint16(anon_sym_enum),
	170:  uint16(47),
	171:  uint16(1),
	172:  uint16(anon_sym_RBRACE),
	173:  uint16(49),
	174:  uint16(1),
	175:  uint16(anon_sym_message),
	176:  uint16(52),
	177:  uint16(1),
	178:  uint16(anon_sym_optional),
	179:  uint16(55),
	180:  uint16(1),
	181:  uint16(anon_sym_repeated),
	182:  uint16(58),
	183:  uint16(1),
	184:  uint16(anon_sym_oneof),
	185:  uint16(61),
	186:  uint16(1),
	187:  uint16(anon_sym_map),
	188:  uint16(67),
	189:  uint16(1),
	190:  uint16(anon_sym_reserved),
	191:  uint16(70),
	192:  uint16(1),
	193:  uint16(sym_identifier),
	194:  uint16(199),
	195:  uint16(1),
	196:  uint16(aux_sym_message_or_enum_type_repeat1),
	197:  uint16(203),
	198:  uint16(1),
	199:  uint16(sym_message_or_enum_type),
	200:  uint16(240),
	201:  uint16(1),
	202:  uint16(sym_type),
	203:  uint16(4),
	204:  uint16(9),
	205:  uint16(sym_empty_statement),
	206:  uint16(sym_option),
	207:  uint16(sym_enum),
	208:  uint16(sym_message),
	209:  uint16(sym_field),
	210:  uint16(sym_oneof),
	211:  uint16(sym_map_field),
	212:  uint16(sym_reserved),
	213:  uint16(aux_sym_message_body_repeat1),
	214:  uint16(64),
	215:  uint16(15),
	216:  uint16(anon_sym_int32),
	217:  uint16(anon_sym_int64),
	218:  uint16(anon_sym_uint32),
	219:  uint16(anon_sym_uint64),
	220:  uint16(anon_sym_sint32),
	221:  uint16(anon_sym_sint64),
	222:  uint16(anon_sym_fixed32),
	223:  uint16(anon_sym_fixed64),
	224:  uint16(anon_sym_sfixed32),
	225:  uint16(anon_sym_sfixed64),
	226:  uint16(anon_sym_bool),
	227:  uint16(anon_sym_string),
	228:  uint16(anon_sym_double),
	229:  uint16(anon_sym_float),
	230:  uint16(anon_sym_bytes),
	231:  uint16(18),
	232:  uint16(3),
	233:  uint16(1),
	234:  uint16(sym_comment),
	235:  uint16(7),
	236:  uint16(1),
	237:  uint16(anon_sym_SEMI),
	238:  uint16(9),
	239:  uint16(1),
	240:  uint16(anon_sym_option),
	241:  uint16(11),
	242:  uint16(1),
	243:  uint16(anon_sym_DOT),
	244:  uint16(13),
	245:  uint16(1),
	246:  uint16(anon_sym_enum),
	247:  uint16(17),
	248:  uint16(1),
	249:  uint16(anon_sym_message),
	250:  uint16(19),
	251:  uint16(1),
	252:  uint16(anon_sym_optional),
	253:  uint16(21),
	254:  uint16(1),
	255:  uint16(anon_sym_repeated),
	256:  uint16(23),
	257:  uint16(1),
	258:  uint16(anon_sym_oneof),
	259:  uint16(25),
	260:  uint16(1),
	261:  uint16(anon_sym_map),
	262:  uint16(29),
	263:  uint16(1),
	264:  uint16(anon_sym_reserved),
	265:  uint16(31),
	266:  uint16(1),
	267:  uint16(sym_identifier),
	268:  uint16(73),
	269:  uint16(1),
	270:  uint16(anon_sym_RBRACE),
	271:  uint16(199),
	272:  uint16(1),
	273:  uint16(aux_sym_message_or_enum_type_repeat1),
	274:  uint16(203),
	275:  uint16(1),
	276:  uint16(sym_message_or_enum_type),
	277:  uint16(240),
	278:  uint16(1),
	279:  uint16(sym_type),
	280:  uint16(4),
	281:  uint16(9),
	282:  uint16(sym_empty_statement),
	283:  uint16(sym_option),
	284:  uint16(sym_enum),
	285:  uint16(sym_message),
	286:  uint16(sym_field),
	287:  uint16(sym_oneof),
	288:  uint16(sym_map_field),
	289:  uint16(sym_reserved),
	290:  uint16(aux_sym_message_body_repeat1),
	291:  uint16(27),
	292:  uint16(15),
	293:  uint16(anon_sym_int32),
	294:  uint16(anon_sym_int64),
	295:  uint16(anon_sym_uint32),
	296:  uint16(anon_sym_uint64),
	297:  uint16(anon_sym_sint32),
	298:  uint16(anon_sym_sint64),
	299:  uint16(anon_sym_fixed32),
	300:  uint16(anon_sym_fixed64),
	301:  uint16(anon_sym_sfixed32),
	302:  uint16(anon_sym_sfixed64),
	303:  uint16(anon_sym_bool),
	304:  uint16(anon_sym_string),
	305:  uint16(anon_sym_double),
	306:  uint16(anon_sym_float),
	307:  uint16(anon_sym_bytes),
	308:  uint16(18),
	309:  uint16(3),
	310:  uint16(1),
	311:  uint16(sym_comment),
	312:  uint16(7),
	313:  uint16(1),
	314:  uint16(anon_sym_SEMI),
	315:  uint16(9),
	316:  uint16(1),
	317:  uint16(anon_sym_option),
	318:  uint16(11),
	319:  uint16(1),
	320:  uint16(anon_sym_DOT),
	321:  uint16(13),
	322:  uint16(1),
	323:  uint16(anon_sym_enum),
	324:  uint16(17),
	325:  uint16(1),
	326:  uint16(anon_sym_message),
	327:  uint16(19),
	328:  uint16(1),
	329:  uint16(anon_sym_optional),
	330:  uint16(21),
	331:  uint16(1),
	332:  uint16(anon_sym_repeated),
	333:  uint16(23),
	334:  uint16(1),
	335:  uint16(anon_sym_oneof),
	336:  uint16(25),
	337:  uint16(1),
	338:  uint16(anon_sym_map),
	339:  uint16(29),
	340:  uint16(1),
	341:  uint16(anon_sym_reserved),
	342:  uint16(31),
	343:  uint16(1),
	344:  uint16(sym_identifier),
	345:  uint16(75),
	346:  uint16(1),
	347:  uint16(anon_sym_RBRACE),
	348:  uint16(199),
	349:  uint16(1),
	350:  uint16(aux_sym_message_or_enum_type_repeat1),
	351:  uint16(203),
	352:  uint16(1),
	353:  uint16(sym_message_or_enum_type),
	354:  uint16(240),
	355:  uint16(1),
	356:  uint16(sym_type),
	357:  uint16(4),
	358:  uint16(9),
	359:  uint16(sym_empty_statement),
	360:  uint16(sym_option),
	361:  uint16(sym_enum),
	362:  uint16(sym_message),
	363:  uint16(sym_field),
	364:  uint16(sym_oneof),
	365:  uint16(sym_map_field),
	366:  uint16(sym_reserved),
	367:  uint16(aux_sym_message_body_repeat1),
	368:  uint16(27),
	369:  uint16(15),
	370:  uint16(anon_sym_int32),
	371:  uint16(anon_sym_int64),
	372:  uint16(anon_sym_uint32),
	373:  uint16(anon_sym_uint64),
	374:  uint16(anon_sym_sint32),
	375:  uint16(anon_sym_sint64),
	376:  uint16(anon_sym_fixed32),
	377:  uint16(anon_sym_fixed64),
	378:  uint16(anon_sym_sfixed32),
	379:  uint16(anon_sym_sfixed64),
	380:  uint16(anon_sym_bool),
	381:  uint16(anon_sym_string),
	382:  uint16(anon_sym_double),
	383:  uint16(anon_sym_float),
	384:  uint16(anon_sym_bytes),
	385:  uint16(3),
	386:  uint16(3),
	387:  uint16(1),
	388:  uint16(sym_comment),
	389:  uint16(77),
	390:  uint16(3),
	391:  uint16(anon_sym_SEMI),
	392:  uint16(anon_sym_DOT),
	393:  uint16(anon_sym_RBRACE),
	394:  uint16(79),
	395:  uint16(24),
	396:  uint16(anon_sym_option),
	397:  uint16(anon_sym_enum),
	398:  uint16(anon_sym_message),
	399:  uint16(anon_sym_optional),
	400:  uint16(anon_sym_repeated),
	401:  uint16(anon_sym_oneof),
	402:  uint16(anon_sym_map),
	403:  uint16(anon_sym_int32),
	404:  uint16(anon_sym_int64),
	405:  uint16(anon_sym_uint32),
	406:  uint16(anon_sym_uint64),
	407:  uint16(anon_sym_sint32),
	408:  uint16(anon_sym_sint64),
	409:  uint16(anon_sym_fixed32),
	410:  uint16(anon_sym_fixed64),
	411:  uint16(anon_sym_sfixed32),
	412:  uint16(anon_sym_sfixed64),
	413:  uint16(anon_sym_bool),
	414:  uint16(anon_sym_string),
	415:  uint16(anon_sym_double),
	416:  uint16(anon_sym_float),
	417:  uint16(anon_sym_bytes),
	418:  uint16(anon_sym_reserved),
	419:  uint16(sym_identifier),
	420:  uint16(3),
	421:  uint16(3),
	422:  uint16(1),
	423:  uint16(sym_comment),
	424:  uint16(81),
	425:  uint16(3),
	426:  uint16(anon_sym_SEMI),
	427:  uint16(anon_sym_DOT),
	428:  uint16(anon_sym_RBRACE),
	429:  uint16(83),
	430:  uint16(24),
	431:  uint16(anon_sym_option),
	432:  uint16(anon_sym_enum),
	433:  uint16(anon_sym_message),
	434:  uint16(anon_sym_optional),
	435:  uint16(anon_sym_repeated),
	436:  uint16(anon_sym_oneof),
	437:  uint16(anon_sym_map),
	438:  uint16(anon_sym_int32),
	439:  uint16(anon_sym_int64),
	440:  uint16(anon_sym_uint32),
	441:  uint16(anon_sym_uint64),
	442:  uint16(anon_sym_sint32),
	443:  uint16(anon_sym_sint64),
	444:  uint16(anon_sym_fixed32),
	445:  uint16(anon_sym_fixed64),
	446:  uint16(anon_sym_sfixed32),
	447:  uint16(anon_sym_sfixed64),
	448:  uint16(anon_sym_bool),
	449:  uint16(anon_sym_string),
	450:  uint16(anon_sym_double),
	451:  uint16(anon_sym_float),
	452:  uint16(anon_sym_bytes),
	453:  uint16(anon_sym_reserved),
	454:  uint16(sym_identifier),
	455:  uint16(11),
	456:  uint16(3),
	457:  uint16(1),
	458:  uint16(sym_comment),
	459:  uint16(11),
	460:  uint16(1),
	461:  uint16(anon_sym_DOT),
	462:  uint16(31),
	463:  uint16(1),
	464:  uint16(sym_identifier),
	465:  uint16(85),
	466:  uint16(1),
	467:  uint16(anon_sym_SEMI),
	468:  uint16(87),
	469:  uint16(1),
	470:  uint16(anon_sym_option),
	471:  uint16(89),
	472:  uint16(1),
	473:  uint16(anon_sym_RBRACE),
	474:  uint16(199),
	475:  uint16(1),
	476:  uint16(aux_sym_message_or_enum_type_repeat1),
	477:  uint16(203),
	478:  uint16(1),
	479:  uint16(sym_message_or_enum_type),
	480:  uint16(271),
	481:  uint16(1),
	482:  uint16(sym_type),
	483:  uint16(28),
	484:  uint16(4),
	485:  uint16(sym_empty_statement),
	486:  uint16(sym_option),
	487:  uint16(sym_oneof_field),
	488:  uint16(aux_sym_oneof_repeat1),
	489:  uint16(27),
	490:  uint16(15),
	491:  uint16(anon_sym_int32),
	492:  uint16(anon_sym_int64),
	493:  uint16(anon_sym_uint32),
	494:  uint16(anon_sym_uint64),
	495:  uint16(anon_sym_sint32),
	496:  uint16(anon_sym_sint64),
	497:  uint16(anon_sym_fixed32),
	498:  uint16(anon_sym_fixed64),
	499:  uint16(anon_sym_sfixed32),
	500:  uint16(anon_sym_sfixed64),
	501:  uint16(anon_sym_bool),
	502:  uint16(anon_sym_string),
	503:  uint16(anon_sym_double),
	504:  uint16(anon_sym_float),
	505:  uint16(anon_sym_bytes),
	506:  uint16(3),
	507:  uint16(3),
	508:  uint16(1),
	509:  uint16(sym_comment),
	510:  uint16(91),
	511:  uint16(3),
	512:  uint16(anon_sym_SEMI),
	513:  uint16(anon_sym_DOT),
	514:  uint16(anon_sym_RBRACE),
	515:  uint16(93),
	516:  uint16(24),
	517:  uint16(anon_sym_option),
	518:  uint16(anon_sym_enum),
	519:  uint16(anon_sym_message),
	520:  uint16(anon_sym_optional),
	521:  uint16(anon_sym_repeated),
	522:  uint16(anon_sym_oneof),
	523:  uint16(anon_sym_map),
	524:  uint16(anon_sym_int32),
	525:  uint16(anon_sym_int64),
	526:  uint16(anon_sym_uint32),
	527:  uint16(anon_sym_uint64),
	528:  uint16(anon_sym_sint32),
	529:  uint16(anon_sym_sint64),
	530:  uint16(anon_sym_fixed32),
	531:  uint16(anon_sym_fixed64),
	532:  uint16(anon_sym_sfixed32),
	533:  uint16(anon_sym_sfixed64),
	534:  uint16(anon_sym_bool),
	535:  uint16(anon_sym_string),
	536:  uint16(anon_sym_double),
	537:  uint16(anon_sym_float),
	538:  uint16(anon_sym_bytes),
	539:  uint16(anon_sym_reserved),
	540:  uint16(sym_identifier),
	541:  uint16(3),
	542:  uint16(3),
	543:  uint16(1),
	544:  uint16(sym_comment),
	545:  uint16(95),
	546:  uint16(3),
	547:  uint16(anon_sym_SEMI),
	548:  uint16(anon_sym_DOT),
	549:  uint16(anon_sym_RBRACE),
	550:  uint16(97),
	551:  uint16(24),
	552:  uint16(anon_sym_option),
	553:  uint16(anon_sym_enum),
	554:  uint16(anon_sym_message),
	555:  uint16(anon_sym_optional),
	556:  uint16(anon_sym_repeated),
	557:  uint16(anon_sym_oneof),
	558:  uint16(anon_sym_map),
	559:  uint16(anon_sym_int32),
	560:  uint16(anon_sym_int64),
	561:  uint16(anon_sym_uint32),
	562:  uint16(anon_sym_uint64),
	563:  uint16(anon_sym_sint32),
	564:  uint16(anon_sym_sint64),
	565:  uint16(anon_sym_fixed32),
	566:  uint16(anon_sym_fixed64),
	567:  uint16(anon_sym_sfixed32),
	568:  uint16(anon_sym_sfixed64),
	569:  uint16(anon_sym_bool),
	570:  uint16(anon_sym_string),
	571:  uint16(anon_sym_double),
	572:  uint16(anon_sym_float),
	573:  uint16(anon_sym_bytes),
	574:  uint16(anon_sym_reserved),
	575:  uint16(sym_identifier),
	576:  uint16(3),
	577:  uint16(3),
	578:  uint16(1),
	579:  uint16(sym_comment),
	580:  uint16(99),
	581:  uint16(3),
	582:  uint16(anon_sym_SEMI),
	583:  uint16(anon_sym_DOT),
	584:  uint16(anon_sym_RBRACE),
	585:  uint16(101),
	586:  uint16(24),
	587:  uint16(anon_sym_option),
	588:  uint16(anon_sym_enum),
	589:  uint16(anon_sym_message),
	590:  uint16(anon_sym_optional),
	591:  uint16(anon_sym_repeated),
	592:  uint16(anon_sym_oneof),
	593:  uint16(anon_sym_map),
	594:  uint16(anon_sym_int32),
	595:  uint16(anon_sym_int64),
	596:  uint16(anon_sym_uint32),
	597:  uint16(anon_sym_uint64),
	598:  uint16(anon_sym_sint32),
	599:  uint16(anon_sym_sint64),
	600:  uint16(anon_sym_fixed32),
	601:  uint16(anon_sym_fixed64),
	602:  uint16(anon_sym_sfixed32),
	603:  uint16(anon_sym_sfixed64),
	604:  uint16(anon_sym_bool),
	605:  uint16(anon_sym_string),
	606:  uint16(anon_sym_double),
	607:  uint16(anon_sym_float),
	608:  uint16(anon_sym_bytes),
	609:  uint16(anon_sym_reserved),
	610:  uint16(sym_identifier),
	611:  uint16(3),
	612:  uint16(3),
	613:  uint16(1),
	614:  uint16(sym_comment),
	615:  uint16(103),
	616:  uint16(3),
	617:  uint16(anon_sym_SEMI),
	618:  uint16(anon_sym_DOT),
	619:  uint16(anon_sym_RBRACE),
	620:  uint16(105),
	621:  uint16(24),
	622:  uint16(anon_sym_option),
	623:  uint16(anon_sym_enum),
	624:  uint16(anon_sym_message),
	625:  uint16(anon_sym_optional),
	626:  uint16(anon_sym_repeated),
	627:  uint16(anon_sym_oneof),
	628:  uint16(anon_sym_map),
	629:  uint16(anon_sym_int32),
	630:  uint16(anon_sym_int64),
	631:  uint16(anon_sym_uint32),
	632:  uint16(anon_sym_uint64),
	633:  uint16(anon_sym_sint32),
	634:  uint16(anon_sym_sint64),
	635:  uint16(anon_sym_fixed32),
	636:  uint16(anon_sym_fixed64),
	637:  uint16(anon_sym_sfixed32),
	638:  uint16(anon_sym_sfixed64),
	639:  uint16(anon_sym_bool),
	640:  uint16(anon_sym_string),
	641:  uint16(anon_sym_double),
	642:  uint16(anon_sym_float),
	643:  uint16(anon_sym_bytes),
	644:  uint16(anon_sym_reserved),
	645:  uint16(sym_identifier),
	646:  uint16(3),
	647:  uint16(3),
	648:  uint16(1),
	649:  uint16(sym_comment),
	650:  uint16(107),
	651:  uint16(3),
	652:  uint16(anon_sym_SEMI),
	653:  uint16(anon_sym_DOT),
	654:  uint16(anon_sym_RBRACE),
	655:  uint16(109),
	656:  uint16(24),
	657:  uint16(anon_sym_option),
	658:  uint16(anon_sym_enum),
	659:  uint16(anon_sym_message),
	660:  uint16(anon_sym_optional),
	661:  uint16(anon_sym_repeated),
	662:  uint16(anon_sym_oneof),
	663:  uint16(anon_sym_map),
	664:  uint16(anon_sym_int32),
	665:  uint16(anon_sym_int64),
	666:  uint16(anon_sym_uint32),
	667:  uint16(anon_sym_uint64),
	668:  uint16(anon_sym_sint32),
	669:  uint16(anon_sym_sint64),
	670:  uint16(anon_sym_fixed32),
	671:  uint16(anon_sym_fixed64),
	672:  uint16(anon_sym_sfixed32),
	673:  uint16(anon_sym_sfixed64),
	674:  uint16(anon_sym_bool),
	675:  uint16(anon_sym_string),
	676:  uint16(anon_sym_double),
	677:  uint16(anon_sym_float),
	678:  uint16(anon_sym_bytes),
	679:  uint16(anon_sym_reserved),
	680:  uint16(sym_identifier),
	681:  uint16(3),
	682:  uint16(3),
	683:  uint16(1),
	684:  uint16(sym_comment),
	685:  uint16(111),
	686:  uint16(3),
	687:  uint16(anon_sym_SEMI),
	688:  uint16(anon_sym_DOT),
	689:  uint16(anon_sym_RBRACE),
	690:  uint16(113),
	691:  uint16(24),
	692:  uint16(anon_sym_option),
	693:  uint16(anon_sym_enum),
	694:  uint16(anon_sym_message),
	695:  uint16(anon_sym_optional),
	696:  uint16(anon_sym_repeated),
	697:  uint16(anon_sym_oneof),
	698:  uint16(anon_sym_map),
	699:  uint16(anon_sym_int32),
	700:  uint16(anon_sym_int64),
	701:  uint16(anon_sym_uint32),
	702:  uint16(anon_sym_uint64),
	703:  uint16(anon_sym_sint32),
	704:  uint16(anon_sym_sint64),
	705:  uint16(anon_sym_fixed32),
	706:  uint16(anon_sym_fixed64),
	707:  uint16(anon_sym_sfixed32),
	708:  uint16(anon_sym_sfixed64),
	709:  uint16(anon_sym_bool),
	710:  uint16(anon_sym_string),
	711:  uint16(anon_sym_double),
	712:  uint16(anon_sym_float),
	713:  uint16(anon_sym_bytes),
	714:  uint16(anon_sym_reserved),
	715:  uint16(sym_identifier),
	716:  uint16(3),
	717:  uint16(3),
	718:  uint16(1),
	719:  uint16(sym_comment),
	720:  uint16(115),
	721:  uint16(3),
	722:  uint16(anon_sym_SEMI),
	723:  uint16(anon_sym_DOT),
	724:  uint16(anon_sym_RBRACE),
	725:  uint16(117),
	726:  uint16(24),
	727:  uint16(anon_sym_option),
	728:  uint16(anon_sym_enum),
	729:  uint16(anon_sym_message),
	730:  uint16(anon_sym_optional),
	731:  uint16(anon_sym_repeated),
	732:  uint16(anon_sym_oneof),
	733:  uint16(anon_sym_map),
	734:  uint16(anon_sym_int32),
	735:  uint16(anon_sym_int64),
	736:  uint16(anon_sym_uint32),
	737:  uint16(anon_sym_uint64),
	738:  uint16(anon_sym_sint32),
	739:  uint16(anon_sym_sint64),
	740:  uint16(anon_sym_fixed32),
	741:  uint16(anon_sym_fixed64),
	742:  uint16(anon_sym_sfixed32),
	743:  uint16(anon_sym_sfixed64),
	744:  uint16(anon_sym_bool),
	745:  uint16(anon_sym_string),
	746:  uint16(anon_sym_double),
	747:  uint16(anon_sym_float),
	748:  uint16(anon_sym_bytes),
	749:  uint16(anon_sym_reserved),
	750:  uint16(sym_identifier),
	751:  uint16(3),
	752:  uint16(3),
	753:  uint16(1),
	754:  uint16(sym_comment),
	755:  uint16(119),
	756:  uint16(3),
	757:  uint16(anon_sym_SEMI),
	758:  uint16(anon_sym_DOT),
	759:  uint16(anon_sym_RBRACE),
	760:  uint16(121),
	761:  uint16(24),
	762:  uint16(anon_sym_option),
	763:  uint16(anon_sym_enum),
	764:  uint16(anon_sym_message),
	765:  uint16(anon_sym_optional),
	766:  uint16(anon_sym_repeated),
	767:  uint16(anon_sym_oneof),
	768:  uint16(anon_sym_map),
	769:  uint16(anon_sym_int32),
	770:  uint16(anon_sym_int64),
	771:  uint16(anon_sym_uint32),
	772:  uint16(anon_sym_uint64),
	773:  uint16(anon_sym_sint32),
	774:  uint16(anon_sym_sint64),
	775:  uint16(anon_sym_fixed32),
	776:  uint16(anon_sym_fixed64),
	777:  uint16(anon_sym_sfixed32),
	778:  uint16(anon_sym_sfixed64),
	779:  uint16(anon_sym_bool),
	780:  uint16(anon_sym_string),
	781:  uint16(anon_sym_double),
	782:  uint16(anon_sym_float),
	783:  uint16(anon_sym_bytes),
	784:  uint16(anon_sym_reserved),
	785:  uint16(sym_identifier),
	786:  uint16(3),
	787:  uint16(3),
	788:  uint16(1),
	789:  uint16(sym_comment),
	790:  uint16(123),
	791:  uint16(3),
	792:  uint16(anon_sym_SEMI),
	793:  uint16(anon_sym_DOT),
	794:  uint16(anon_sym_RBRACE),
	795:  uint16(125),
	796:  uint16(24),
	797:  uint16(anon_sym_option),
	798:  uint16(anon_sym_enum),
	799:  uint16(anon_sym_message),
	800:  uint16(anon_sym_optional),
	801:  uint16(anon_sym_repeated),
	802:  uint16(anon_sym_oneof),
	803:  uint16(anon_sym_map),
	804:  uint16(anon_sym_int32),
	805:  uint16(anon_sym_int64),
	806:  uint16(anon_sym_uint32),
	807:  uint16(anon_sym_uint64),
	808:  uint16(anon_sym_sint32),
	809:  uint16(anon_sym_sint64),
	810:  uint16(anon_sym_fixed32),
	811:  uint16(anon_sym_fixed64),
	812:  uint16(anon_sym_sfixed32),
	813:  uint16(anon_sym_sfixed64),
	814:  uint16(anon_sym_bool),
	815:  uint16(anon_sym_string),
	816:  uint16(anon_sym_double),
	817:  uint16(anon_sym_float),
	818:  uint16(anon_sym_bytes),
	819:  uint16(anon_sym_reserved),
	820:  uint16(sym_identifier),
	821:  uint16(3),
	822:  uint16(3),
	823:  uint16(1),
	824:  uint16(sym_comment),
	825:  uint16(127),
	826:  uint16(3),
	827:  uint16(anon_sym_SEMI),
	828:  uint16(anon_sym_DOT),
	829:  uint16(anon_sym_RBRACE),
	830:  uint16(129),
	831:  uint16(24),
	832:  uint16(anon_sym_option),
	833:  uint16(anon_sym_enum),
	834:  uint16(anon_sym_message),
	835:  uint16(anon_sym_optional),
	836:  uint16(anon_sym_repeated),
	837:  uint16(anon_sym_oneof),
	838:  uint16(anon_sym_map),
	839:  uint16(anon_sym_int32),
	840:  uint16(anon_sym_int64),
	841:  uint16(anon_sym_uint32),
	842:  uint16(anon_sym_uint64),
	843:  uint16(anon_sym_sint32),
	844:  uint16(anon_sym_sint64),
	845:  uint16(anon_sym_fixed32),
	846:  uint16(anon_sym_fixed64),
	847:  uint16(anon_sym_sfixed32),
	848:  uint16(anon_sym_sfixed64),
	849:  uint16(anon_sym_bool),
	850:  uint16(anon_sym_string),
	851:  uint16(anon_sym_double),
	852:  uint16(anon_sym_float),
	853:  uint16(anon_sym_bytes),
	854:  uint16(anon_sym_reserved),
	855:  uint16(sym_identifier),
	856:  uint16(3),
	857:  uint16(3),
	858:  uint16(1),
	859:  uint16(sym_comment),
	860:  uint16(131),
	861:  uint16(3),
	862:  uint16(anon_sym_SEMI),
	863:  uint16(anon_sym_DOT),
	864:  uint16(anon_sym_RBRACE),
	865:  uint16(133),
	866:  uint16(24),
	867:  uint16(anon_sym_option),
	868:  uint16(anon_sym_enum),
	869:  uint16(anon_sym_message),
	870:  uint16(anon_sym_optional),
	871:  uint16(anon_sym_repeated),
	872:  uint16(anon_sym_oneof),
	873:  uint16(anon_sym_map),
	874:  uint16(anon_sym_int32),
	875:  uint16(anon_sym_int64),
	876:  uint16(anon_sym_uint32),
	877:  uint16(anon_sym_uint64),
	878:  uint16(anon_sym_sint32),
	879:  uint16(anon_sym_sint64),
	880:  uint16(anon_sym_fixed32),
	881:  uint16(anon_sym_fixed64),
	882:  uint16(anon_sym_sfixed32),
	883:  uint16(anon_sym_sfixed64),
	884:  uint16(anon_sym_bool),
	885:  uint16(anon_sym_string),
	886:  uint16(anon_sym_double),
	887:  uint16(anon_sym_float),
	888:  uint16(anon_sym_bytes),
	889:  uint16(anon_sym_reserved),
	890:  uint16(sym_identifier),
	891:  uint16(3),
	892:  uint16(3),
	893:  uint16(1),
	894:  uint16(sym_comment),
	895:  uint16(135),
	896:  uint16(3),
	897:  uint16(anon_sym_SEMI),
	898:  uint16(anon_sym_DOT),
	899:  uint16(anon_sym_RBRACE),
	900:  uint16(137),
	901:  uint16(24),
	902:  uint16(anon_sym_option),
	903:  uint16(anon_sym_enum),
	904:  uint16(anon_sym_message),
	905:  uint16(anon_sym_optional),
	906:  uint16(anon_sym_repeated),
	907:  uint16(anon_sym_oneof),
	908:  uint16(anon_sym_map),
	909:  uint16(anon_sym_int32),
	910:  uint16(anon_sym_int64),
	911:  uint16(anon_sym_uint32),
	912:  uint16(anon_sym_uint64),
	913:  uint16(anon_sym_sint32),
	914:  uint16(anon_sym_sint64),
	915:  uint16(anon_sym_fixed32),
	916:  uint16(anon_sym_fixed64),
	917:  uint16(anon_sym_sfixed32),
	918:  uint16(anon_sym_sfixed64),
	919:  uint16(anon_sym_bool),
	920:  uint16(anon_sym_string),
	921:  uint16(anon_sym_double),
	922:  uint16(anon_sym_float),
	923:  uint16(anon_sym_bytes),
	924:  uint16(anon_sym_reserved),
	925:  uint16(sym_identifier),
	926:  uint16(3),
	927:  uint16(3),
	928:  uint16(1),
	929:  uint16(sym_comment),
	930:  uint16(139),
	931:  uint16(3),
	932:  uint16(anon_sym_SEMI),
	933:  uint16(anon_sym_DOT),
	934:  uint16(anon_sym_RBRACE),
	935:  uint16(141),
	936:  uint16(24),
	937:  uint16(anon_sym_option),
	938:  uint16(anon_sym_enum),
	939:  uint16(anon_sym_message),
	940:  uint16(anon_sym_optional),
	941:  uint16(anon_sym_repeated),
	942:  uint16(anon_sym_oneof),
	943:  uint16(anon_sym_map),
	944:  uint16(anon_sym_int32),
	945:  uint16(anon_sym_int64),
	946:  uint16(anon_sym_uint32),
	947:  uint16(anon_sym_uint64),
	948:  uint16(anon_sym_sint32),
	949:  uint16(anon_sym_sint64),
	950:  uint16(anon_sym_fixed32),
	951:  uint16(anon_sym_fixed64),
	952:  uint16(anon_sym_sfixed32),
	953:  uint16(anon_sym_sfixed64),
	954:  uint16(anon_sym_bool),
	955:  uint16(anon_sym_string),
	956:  uint16(anon_sym_double),
	957:  uint16(anon_sym_float),
	958:  uint16(anon_sym_bytes),
	959:  uint16(anon_sym_reserved),
	960:  uint16(sym_identifier),
	961:  uint16(3),
	962:  uint16(3),
	963:  uint16(1),
	964:  uint16(sym_comment),
	965:  uint16(143),
	966:  uint16(3),
	967:  uint16(anon_sym_SEMI),
	968:  uint16(anon_sym_DOT),
	969:  uint16(anon_sym_RBRACE),
	970:  uint16(145),
	971:  uint16(24),
	972:  uint16(anon_sym_option),
	973:  uint16(anon_sym_enum),
	974:  uint16(anon_sym_message),
	975:  uint16(anon_sym_optional),
	976:  uint16(anon_sym_repeated),
	977:  uint16(anon_sym_oneof),
	978:  uint16(anon_sym_map),
	979:  uint16(anon_sym_int32),
	980:  uint16(anon_sym_int64),
	981:  uint16(anon_sym_uint32),
	982:  uint16(anon_sym_uint64),
	983:  uint16(anon_sym_sint32),
	984:  uint16(anon_sym_sint64),
	985:  uint16(anon_sym_fixed32),
	986:  uint16(anon_sym_fixed64),
	987:  uint16(anon_sym_sfixed32),
	988:  uint16(anon_sym_sfixed64),
	989:  uint16(anon_sym_bool),
	990:  uint16(anon_sym_string),
	991:  uint16(anon_sym_double),
	992:  uint16(anon_sym_float),
	993:  uint16(anon_sym_bytes),
	994:  uint16(anon_sym_reserved),
	995:  uint16(sym_identifier),
	996:  uint16(11),
	997:  uint16(3),
	998:  uint16(1),
	999:  uint16(sym_comment),
	1000: uint16(147),
	1001: uint16(1),
	1002: uint16(anon_sym_SEMI),
	1003: uint16(150),
	1004: uint16(1),
	1005: uint16(anon_sym_option),
	1006: uint16(153),
	1007: uint16(1),
	1008: uint16(anon_sym_DOT),
	1009: uint16(156),
	1010: uint16(1),
	1011: uint16(anon_sym_RBRACE),
	1012: uint16(161),
	1013: uint16(1),
	1014: uint16(sym_identifier),
	1015: uint16(199),
	1016: uint16(1),
	1017: uint16(aux_sym_message_or_enum_type_repeat1),
	1018: uint16(203),
	1019: uint16(1),
	1020: uint16(sym_message_or_enum_type),
	1021: uint16(271),
	1022: uint16(1),
	1023: uint16(sym_type),
	1024: uint16(24),
	1025: uint16(4),
	1026: uint16(sym_empty_statement),
	1027: uint16(sym_option),
	1028: uint16(sym_oneof_field),
	1029: uint16(aux_sym_oneof_repeat1),
	1030: uint16(158),
	1031: uint16(15),
	1032: uint16(anon_sym_int32),
	1033: uint16(anon_sym_int64),
	1034: uint16(anon_sym_uint32),
	1035: uint16(anon_sym_uint64),
	1036: uint16(anon_sym_sint32),
	1037: uint16(anon_sym_sint64),
	1038: uint16(anon_sym_fixed32),
	1039: uint16(anon_sym_fixed64),
	1040: uint16(anon_sym_sfixed32),
	1041: uint16(anon_sym_sfixed64),
	1042: uint16(anon_sym_bool),
	1043: uint16(anon_sym_string),
	1044: uint16(anon_sym_double),
	1045: uint16(anon_sym_float),
	1046: uint16(anon_sym_bytes),
	1047: uint16(3),
	1048: uint16(3),
	1049: uint16(1),
	1050: uint16(sym_comment),
	1051: uint16(164),
	1052: uint16(3),
	1053: uint16(anon_sym_SEMI),
	1054: uint16(anon_sym_DOT),
	1055: uint16(anon_sym_RBRACE),
	1056: uint16(166),
	1057: uint16(24),
	1058: uint16(anon_sym_option),
	1059: uint16(anon_sym_enum),
	1060: uint16(anon_sym_message),
	1061: uint16(anon_sym_optional),
	1062: uint16(anon_sym_repeated),
	1063: uint16(anon_sym_oneof),
	1064: uint16(anon_sym_map),
	1065: uint16(anon_sym_int32),
	1066: uint16(anon_sym_int64),
	1067: uint16(anon_sym_uint32),
	1068: uint16(anon_sym_uint64),
	1069: uint16(anon_sym_sint32),
	1070: uint16(anon_sym_sint64),
	1071: uint16(anon_sym_fixed32),
	1072: uint16(anon_sym_fixed64),
	1073: uint16(anon_sym_sfixed32),
	1074: uint16(anon_sym_sfixed64),
	1075: uint16(anon_sym_bool),
	1076: uint16(anon_sym_string),
	1077: uint16(anon_sym_double),
	1078: uint16(anon_sym_float),
	1079: uint16(anon_sym_bytes),
	1080: uint16(anon_sym_reserved),
	1081: uint16(sym_identifier),
	1082: uint16(3),
	1083: uint16(3),
	1084: uint16(1),
	1085: uint16(sym_comment),
	1086: uint16(168),
	1087: uint16(3),
	1088: uint16(anon_sym_SEMI),
	1089: uint16(anon_sym_DOT),
	1090: uint16(anon_sym_RBRACE),
	1091: uint16(170),
	1092: uint16(24),
	1093: uint16(anon_sym_option),
	1094: uint16(anon_sym_enum),
	1095: uint16(anon_sym_message),
	1096: uint16(anon_sym_optional),
	1097: uint16(anon_sym_repeated),
	1098: uint16(anon_sym_oneof),
	1099: uint16(anon_sym_map),
	1100: uint16(anon_sym_int32),
	1101: uint16(anon_sym_int64),
	1102: uint16(anon_sym_uint32),
	1103: uint16(anon_sym_uint64),
	1104: uint16(anon_sym_sint32),
	1105: uint16(anon_sym_sint64),
	1106: uint16(anon_sym_fixed32),
	1107: uint16(anon_sym_fixed64),
	1108: uint16(anon_sym_sfixed32),
	1109: uint16(anon_sym_sfixed64),
	1110: uint16(anon_sym_bool),
	1111: uint16(anon_sym_string),
	1112: uint16(anon_sym_double),
	1113: uint16(anon_sym_float),
	1114: uint16(anon_sym_bytes),
	1115: uint16(anon_sym_reserved),
	1116: uint16(sym_identifier),
	1117: uint16(3),
	1118: uint16(3),
	1119: uint16(1),
	1120: uint16(sym_comment),
	1121: uint16(172),
	1122: uint16(3),
	1123: uint16(anon_sym_SEMI),
	1124: uint16(anon_sym_DOT),
	1125: uint16(anon_sym_RBRACE),
	1126: uint16(174),
	1127: uint16(24),
	1128: uint16(anon_sym_option),
	1129: uint16(anon_sym_enum),
	1130: uint16(anon_sym_message),
	1131: uint16(anon_sym_optional),
	1132: uint16(anon_sym_repeated),
	1133: uint16(anon_sym_oneof),
	1134: uint16(anon_sym_map),
	1135: uint16(anon_sym_int32),
	1136: uint16(anon_sym_int64),
	1137: uint16(anon_sym_uint32),
	1138: uint16(anon_sym_uint64),
	1139: uint16(anon_sym_sint32),
	1140: uint16(anon_sym_sint64),
	1141: uint16(anon_sym_fixed32),
	1142: uint16(anon_sym_fixed64),
	1143: uint16(anon_sym_sfixed32),
	1144: uint16(anon_sym_sfixed64),
	1145: uint16(anon_sym_bool),
	1146: uint16(anon_sym_string),
	1147: uint16(anon_sym_double),
	1148: uint16(anon_sym_float),
	1149: uint16(anon_sym_bytes),
	1150: uint16(anon_sym_reserved),
	1151: uint16(sym_identifier),
	1152: uint16(11),
	1153: uint16(3),
	1154: uint16(1),
	1155: uint16(sym_comment),
	1156: uint16(11),
	1157: uint16(1),
	1158: uint16(anon_sym_DOT),
	1159: uint16(31),
	1160: uint16(1),
	1161: uint16(sym_identifier),
	1162: uint16(85),
	1163: uint16(1),
	1164: uint16(anon_sym_SEMI),
	1165: uint16(87),
	1166: uint16(1),
	1167: uint16(anon_sym_option),
	1168: uint16(176),
	1169: uint16(1),
	1170: uint16(anon_sym_RBRACE),
	1171: uint16(199),
	1172: uint16(1),
	1173: uint16(aux_sym_message_or_enum_type_repeat1),
	1174: uint16(203),
	1175: uint16(1),
	1176: uint16(sym_message_or_enum_type),
	1177: uint16(271),
	1178: uint16(1),
	1179: uint16(sym_type),
	1180: uint16(24),
	1181: uint16(4),
	1182: uint16(sym_empty_statement),
	1183: uint16(sym_option),
	1184: uint16(sym_oneof_field),
	1185: uint16(aux_sym_oneof_repeat1),
	1186: uint16(27),
	1187: uint16(15),
	1188: uint16(anon_sym_int32),
	1189: uint16(anon_sym_int64),
	1190: uint16(anon_sym_uint32),
	1191: uint16(anon_sym_uint64),
	1192: uint16(anon_sym_sint32),
	1193: uint16(anon_sym_sint64),
	1194: uint16(anon_sym_fixed32),
	1195: uint16(anon_sym_fixed64),
	1196: uint16(anon_sym_sfixed32),
	1197: uint16(anon_sym_sfixed64),
	1198: uint16(anon_sym_bool),
	1199: uint16(anon_sym_string),
	1200: uint16(anon_sym_double),
	1201: uint16(anon_sym_float),
	1202: uint16(anon_sym_bytes),
	1203: uint16(4),
	1204: uint16(3),
	1205: uint16(1),
	1206: uint16(sym_comment),
	1207: uint16(182),
	1208: uint16(1),
	1209: uint16(anon_sym_LBRACK),
	1210: uint16(178),
	1211: uint16(3),
	1212: uint16(anon_sym_SEMI),
	1213: uint16(anon_sym_DOT),
	1214: uint16(anon_sym_RBRACE),
	1215: uint16(180),
	1216: uint16(17),
	1217: uint16(anon_sym_option),
	1218: uint16(anon_sym_int32),
	1219: uint16(anon_sym_int64),
	1220: uint16(anon_sym_uint32),
	1221: uint16(anon_sym_uint64),
	1222: uint16(anon_sym_sint32),
	1223: uint16(anon_sym_sint64),
	1224: uint16(anon_sym_fixed32),
	1225: uint16(anon_sym_fixed64),
	1226: uint16(anon_sym_sfixed32),
	1227: uint16(anon_sym_sfixed64),
	1228: uint16(anon_sym_bool),
	1229: uint16(anon_sym_string),
	1230: uint16(anon_sym_double),
	1231: uint16(anon_sym_float),
	1232: uint16(anon_sym_bytes),
	1233: uint16(sym_identifier),
	1234: uint16(8),
	1235: uint16(3),
	1236: uint16(1),
	1237: uint16(sym_comment),
	1238: uint16(11),
	1239: uint16(1),
	1240: uint16(anon_sym_DOT),
	1241: uint16(31),
	1242: uint16(1),
	1243: uint16(sym_identifier),
	1244: uint16(184),
	1245: uint16(1),
	1246: uint16(anon_sym_repeated),
	1247: uint16(199),
	1248: uint16(1),
	1249: uint16(aux_sym_message_or_enum_type_repeat1),
	1250: uint16(203),
	1251: uint16(1),
	1252: uint16(sym_message_or_enum_type),
	1253: uint16(259),
	1254: uint16(1),
	1255: uint16(sym_type),
	1256: uint16(27),
	1257: uint16(15),
	1258: uint16(anon_sym_int32),
	1259: uint16(anon_sym_int64),
	1260: uint16(anon_sym_uint32),
	1261: uint16(anon_sym_uint64),
	1262: uint16(anon_sym_sint32),
	1263: uint16(anon_sym_sint64),
	1264: uint16(anon_sym_fixed32),
	1265: uint16(anon_sym_fixed64),
	1266: uint16(anon_sym_sfixed32),
	1267: uint16(anon_sym_sfixed64),
	1268: uint16(anon_sym_bool),
	1269: uint16(anon_sym_string),
	1270: uint16(anon_sym_double),
	1271: uint16(anon_sym_float),
	1272: uint16(anon_sym_bytes),
	1273: uint16(3),
	1274: uint16(3),
	1275: uint16(1),
	1276: uint16(sym_comment),
	1277: uint16(186),
	1278: uint16(4),
	1279: uint16(anon_sym_SEMI),
	1280: uint16(anon_sym_DOT),
	1281: uint16(anon_sym_RBRACE),
	1282: uint16(anon_sym_LBRACK),
	1283: uint16(188),
	1284: uint16(17),
	1285: uint16(anon_sym_option),
	1286: uint16(anon_sym_int32),
	1287: uint16(anon_sym_int64),
	1288: uint16(anon_sym_uint32),
	1289: uint16(anon_sym_uint64),
	1290: uint16(anon_sym_sint32),
	1291: uint16(anon_sym_sint64),
	1292: uint16(anon_sym_fixed32),
	1293: uint16(anon_sym_fixed64),
	1294: uint16(anon_sym_sfixed32),
	1295: uint16(anon_sym_sfixed64),
	1296: uint16(anon_sym_bool),
	1297: uint16(anon_sym_string),
	1298: uint16(anon_sym_double),
	1299: uint16(anon_sym_float),
	1300: uint16(anon_sym_bytes),
	1301: uint16(sym_identifier),
	1302: uint16(3),
	1303: uint16(3),
	1304: uint16(1),
	1305: uint16(sym_comment),
	1306: uint16(190),
	1307: uint16(4),
	1308: uint16(anon_sym_SEMI),
	1309: uint16(anon_sym_DOT),
	1310: uint16(anon_sym_RBRACE),
	1311: uint16(anon_sym_LBRACK),
	1312: uint16(192),
	1313: uint16(17),
	1314: uint16(anon_sym_option),
	1315: uint16(anon_sym_int32),
	1316: uint16(anon_sym_int64),
	1317: uint16(anon_sym_uint32),
	1318: uint16(anon_sym_uint64),
	1319: uint16(anon_sym_sint32),
	1320: uint16(anon_sym_sint64),
	1321: uint16(anon_sym_fixed32),
	1322: uint16(anon_sym_fixed64),
	1323: uint16(anon_sym_sfixed32),
	1324: uint16(anon_sym_sfixed64),
	1325: uint16(anon_sym_bool),
	1326: uint16(anon_sym_string),
	1327: uint16(anon_sym_double),
	1328: uint16(anon_sym_float),
	1329: uint16(anon_sym_bytes),
	1330: uint16(sym_identifier),
	1331: uint16(14),
	1332: uint16(3),
	1333: uint16(1),
	1334: uint16(sym_comment),
	1335: uint16(194),
	1336: uint16(1),
	1337: uint16(anon_sym_LBRACE),
	1338: uint16(198),
	1339: uint16(1),
	1340: uint16(anon_sym_LBRACK),
	1341: uint16(200),
	1342: uint16(1),
	1343: uint16(anon_sym_COLON),
	1344: uint16(202),
	1345: uint16(1),
	1346: uint16(sym_identifier),
	1347: uint16(208),
	1348: uint16(1),
	1349: uint16(sym_hex_lit),
	1350: uint16(210),
	1351: uint16(1),
	1352: uint16(sym_float_lit),
	1353: uint16(212),
	1354: uint16(1),
	1355: uint16(anon_sym_DQUOTE),
	1356: uint16(214),
	1357: uint16(1),
	1358: uint16(anon_sym_SQUOTE),
	1359: uint16(143),
	1360: uint16(1),
	1361: uint16(sym_constant),
	1362: uint16(196),
	1363: uint16(2),
	1364: uint16(anon_sym_DASH),
	1365: uint16(anon_sym_PLUS),
	1366: uint16(204),
	1367: uint16(2),
	1368: uint16(sym_true),
	1369: uint16(sym_false),
	1370: uint16(206),
	1371: uint16(2),
	1372: uint16(sym_decimal_lit),
	1373: uint16(sym_octal_lit),
	1374: uint16(96),
	1375: uint16(5),
	1376: uint16(sym_block_lit),
	1377: uint16(sym_full_ident),
	1378: uint16(sym_bool),
	1379: uint16(sym_int_lit),
	1380: uint16(sym_string),
	1381: uint16(7),
	1382: uint16(3),
	1383: uint16(1),
	1384: uint16(sym_comment),
	1385: uint16(11),
	1386: uint16(1),
	1387: uint16(anon_sym_DOT),
	1388: uint16(31),
	1389: uint16(1),
	1390: uint16(sym_identifier),
	1391: uint16(199),
	1392: uint16(1),
	1393: uint16(aux_sym_message_or_enum_type_repeat1),
	1394: uint16(203),
	1395: uint16(1),
	1396: uint16(sym_message_or_enum_type),
	1397: uint16(259),
	1398: uint16(1),
	1399: uint16(sym_type),
	1400: uint16(27),
	1401: uint16(15),
	1402: uint16(anon_sym_int32),
	1403: uint16(anon_sym_int64),
	1404: uint16(anon_sym_uint32),
	1405: uint16(anon_sym_uint64),
	1406: uint16(anon_sym_sint32),
	1407: uint16(anon_sym_sint64),
	1408: uint16(anon_sym_fixed32),
	1409: uint16(anon_sym_fixed64),
	1410: uint16(anon_sym_sfixed32),
	1411: uint16(anon_sym_sfixed64),
	1412: uint16(anon_sym_bool),
	1413: uint16(anon_sym_string),
	1414: uint16(anon_sym_double),
	1415: uint16(anon_sym_float),
	1416: uint16(anon_sym_bytes),
	1417: uint16(3),
	1418: uint16(3),
	1419: uint16(1),
	1420: uint16(sym_comment),
	1421: uint16(123),
	1422: uint16(3),
	1423: uint16(anon_sym_SEMI),
	1424: uint16(anon_sym_DOT),
	1425: uint16(anon_sym_RBRACE),
	1426: uint16(125),
	1427: uint16(17),
	1428: uint16(anon_sym_option),
	1429: uint16(anon_sym_int32),
	1430: uint16(anon_sym_int64),
	1431: uint16(anon_sym_uint32),
	1432: uint16(anon_sym_uint64),
	1433: uint16(anon_sym_sint32),
	1434: uint16(anon_sym_sint64),
	1435: uint16(anon_sym_fixed32),
	1436: uint16(anon_sym_fixed64),
	1437: uint16(anon_sym_sfixed32),
	1438: uint16(anon_sym_sfixed64),
	1439: uint16(anon_sym_bool),
	1440: uint16(anon_sym_string),
	1441: uint16(anon_sym_double),
	1442: uint16(anon_sym_float),
	1443: uint16(anon_sym_bytes),
	1444: uint16(sym_identifier),
	1445: uint16(3),
	1446: uint16(3),
	1447: uint16(1),
	1448: uint16(sym_comment),
	1449: uint16(115),
	1450: uint16(3),
	1451: uint16(anon_sym_SEMI),
	1452: uint16(anon_sym_DOT),
	1453: uint16(anon_sym_RBRACE),
	1454: uint16(117),
	1455: uint16(17),
	1456: uint16(anon_sym_option),
	1457: uint16(anon_sym_int32),
	1458: uint16(anon_sym_int64),
	1459: uint16(anon_sym_uint32),
	1460: uint16(anon_sym_uint64),
	1461: uint16(anon_sym_sint32),
	1462: uint16(anon_sym_sint64),
	1463: uint16(anon_sym_fixed32),
	1464: uint16(anon_sym_fixed64),
	1465: uint16(anon_sym_sfixed32),
	1466: uint16(anon_sym_sfixed64),
	1467: uint16(anon_sym_bool),
	1468: uint16(anon_sym_string),
	1469: uint16(anon_sym_double),
	1470: uint16(anon_sym_float),
	1471: uint16(anon_sym_bytes),
	1472: uint16(sym_identifier),
	1473: uint16(7),
	1474: uint16(3),
	1475: uint16(1),
	1476: uint16(sym_comment),
	1477: uint16(11),
	1478: uint16(1),
	1479: uint16(anon_sym_DOT),
	1480: uint16(31),
	1481: uint16(1),
	1482: uint16(sym_identifier),
	1483: uint16(199),
	1484: uint16(1),
	1485: uint16(aux_sym_message_or_enum_type_repeat1),
	1486: uint16(203),
	1487: uint16(1),
	1488: uint16(sym_message_or_enum_type),
	1489: uint16(236),
	1490: uint16(1),
	1491: uint16(sym_type),
	1492: uint16(27),
	1493: uint16(15),
	1494: uint16(anon_sym_int32),
	1495: uint16(anon_sym_int64),
	1496: uint16(anon_sym_uint32),
	1497: uint16(anon_sym_uint64),
	1498: uint16(anon_sym_sint32),
	1499: uint16(anon_sym_sint64),
	1500: uint16(anon_sym_fixed32),
	1501: uint16(anon_sym_fixed64),
	1502: uint16(anon_sym_sfixed32),
	1503: uint16(anon_sym_sfixed64),
	1504: uint16(anon_sym_bool),
	1505: uint16(anon_sym_string),
	1506: uint16(anon_sym_double),
	1507: uint16(anon_sym_float),
	1508: uint16(anon_sym_bytes),
	1509: uint16(3),
	1510: uint16(3),
	1511: uint16(1),
	1512: uint16(sym_comment),
	1513: uint16(216),
	1514: uint16(3),
	1515: uint16(anon_sym_SEMI),
	1516: uint16(anon_sym_DOT),
	1517: uint16(anon_sym_RBRACE),
	1518: uint16(218),
	1519: uint16(17),
	1520: uint16(anon_sym_option),
	1521: uint16(anon_sym_int32),
	1522: uint16(anon_sym_int64),
	1523: uint16(anon_sym_uint32),
	1524: uint16(anon_sym_uint64),
	1525: uint16(anon_sym_sint32),
	1526: uint16(anon_sym_sint64),
	1527: uint16(anon_sym_fixed32),
	1528: uint16(anon_sym_fixed64),
	1529: uint16(anon_sym_sfixed32),
	1530: uint16(anon_sym_sfixed64),
	1531: uint16(anon_sym_bool),
	1532: uint16(anon_sym_string),
	1533: uint16(anon_sym_double),
	1534: uint16(anon_sym_float),
	1535: uint16(anon_sym_bytes),
	1536: uint16(sym_identifier),
	1537: uint16(7),
	1538: uint16(3),
	1539: uint16(1),
	1540: uint16(sym_comment),
	1541: uint16(11),
	1542: uint16(1),
	1543: uint16(anon_sym_DOT),
	1544: uint16(31),
	1545: uint16(1),
	1546: uint16(sym_identifier),
	1547: uint16(199),
	1548: uint16(1),
	1549: uint16(aux_sym_message_or_enum_type_repeat1),
	1550: uint16(203),
	1551: uint16(1),
	1552: uint16(sym_message_or_enum_type),
	1553: uint16(261),
	1554: uint16(1),
	1555: uint16(sym_type),
	1556: uint16(27),
	1557: uint16(15),
	1558: uint16(anon_sym_int32),
	1559: uint16(anon_sym_int64),
	1560: uint16(anon_sym_uint32),
	1561: uint16(anon_sym_uint64),
	1562: uint16(anon_sym_sint32),
	1563: uint16(anon_sym_sint64),
	1564: uint16(anon_sym_fixed32),
	1565: uint16(anon_sym_fixed64),
	1566: uint16(anon_sym_sfixed32),
	1567: uint16(anon_sym_sfixed64),
	1568: uint16(anon_sym_bool),
	1569: uint16(anon_sym_string),
	1570: uint16(anon_sym_double),
	1571: uint16(anon_sym_float),
	1572: uint16(anon_sym_bytes),
	1573: uint16(13),
	1574: uint16(3),
	1575: uint16(1),
	1576: uint16(sym_comment),
	1577: uint16(194),
	1578: uint16(1),
	1579: uint16(anon_sym_LBRACE),
	1580: uint16(202),
	1581: uint16(1),
	1582: uint16(sym_identifier),
	1583: uint16(208),
	1584: uint16(1),
	1585: uint16(sym_hex_lit),
	1586: uint16(210),
	1587: uint16(1),
	1588: uint16(sym_float_lit),
	1589: uint16(212),
	1590: uint16(1),
	1591: uint16(anon_sym_DQUOTE),
	1592: uint16(214),
	1593: uint16(1),
	1594: uint16(anon_sym_SQUOTE),
	1595: uint16(220),
	1596: uint16(1),
	1597: uint16(anon_sym_LBRACK),
	1598: uint16(131),
	1599: uint16(1),
	1600: uint16(sym_constant),
	1601: uint16(196),
	1602: uint16(2),
	1603: uint16(anon_sym_DASH),
	1604: uint16(anon_sym_PLUS),
	1605: uint16(204),
	1606: uint16(2),
	1607: uint16(sym_true),
	1608: uint16(sym_false),
	1609: uint16(206),
	1610: uint16(2),
	1611: uint16(sym_decimal_lit),
	1612: uint16(sym_octal_lit),
	1613: uint16(96),
	1614: uint16(5),
	1615: uint16(sym_block_lit),
	1616: uint16(sym_full_ident),
	1617: uint16(sym_bool),
	1618: uint16(sym_int_lit),
	1619: uint16(sym_string),
	1620: uint16(12),
	1621: uint16(3),
	1622: uint16(1),
	1623: uint16(sym_comment),
	1624: uint16(194),
	1625: uint16(1),
	1626: uint16(anon_sym_LBRACE),
	1627: uint16(202),
	1628: uint16(1),
	1629: uint16(sym_identifier),
	1630: uint16(210),
	1631: uint16(1),
	1632: uint16(sym_float_lit),
	1633: uint16(212),
	1634: uint16(1),
	1635: uint16(anon_sym_DQUOTE),
	1636: uint16(214),
	1637: uint16(1),
	1638: uint16(anon_sym_SQUOTE),
	1639: uint16(226),
	1640: uint16(1),
	1641: uint16(sym_hex_lit),
	1642: uint16(209),
	1643: uint16(1),
	1644: uint16(sym_constant),
	1645: uint16(204),
	1646: uint16(2),
	1647: uint16(sym_true),
	1648: uint16(sym_false),
	1649: uint16(222),
	1650: uint16(2),
	1651: uint16(anon_sym_DASH),
	1652: uint16(anon_sym_PLUS),
	1653: uint16(224),
	1654: uint16(2),
	1655: uint16(sym_decimal_lit),
	1656: uint16(sym_octal_lit),
	1657: uint16(96),
	1658: uint16(5),
	1659: uint16(sym_block_lit),
	1660: uint16(sym_full_ident),
	1661: uint16(sym_bool),
	1662: uint16(sym_int_lit),
	1663: uint16(sym_string),
	1664: uint16(12),
	1665: uint16(3),
	1666: uint16(1),
	1667: uint16(sym_comment),
	1668: uint16(194),
	1669: uint16(1),
	1670: uint16(anon_sym_LBRACE),
	1671: uint16(202),
	1672: uint16(1),
	1673: uint16(sym_identifier),
	1674: uint16(210),
	1675: uint16(1),
	1676: uint16(sym_float_lit),
	1677: uint16(212),
	1678: uint16(1),
	1679: uint16(anon_sym_DQUOTE),
	1680: uint16(214),
	1681: uint16(1),
	1682: uint16(anon_sym_SQUOTE),
	1683: uint16(226),
	1684: uint16(1),
	1685: uint16(sym_hex_lit),
	1686: uint16(154),
	1687: uint16(1),
	1688: uint16(sym_constant),
	1689: uint16(204),
	1690: uint16(2),
	1691: uint16(sym_true),
	1692: uint16(sym_false),
	1693: uint16(222),
	1694: uint16(2),
	1695: uint16(anon_sym_DASH),
	1696: uint16(anon_sym_PLUS),
	1697: uint16(224),
	1698: uint16(2),
	1699: uint16(sym_decimal_lit),
	1700: uint16(sym_octal_lit),
	1701: uint16(96),
	1702: uint16(5),
	1703: uint16(sym_block_lit),
	1704: uint16(sym_full_ident),
	1705: uint16(sym_bool),
	1706: uint16(sym_int_lit),
	1707: uint16(sym_string),
	1708: uint16(12),
	1709: uint16(3),
	1710: uint16(1),
	1711: uint16(sym_comment),
	1712: uint16(194),
	1713: uint16(1),
	1714: uint16(anon_sym_LBRACE),
	1715: uint16(202),
	1716: uint16(1),
	1717: uint16(sym_identifier),
	1718: uint16(210),
	1719: uint16(1),
	1720: uint16(sym_float_lit),
	1721: uint16(212),
	1722: uint16(1),
	1723: uint16(anon_sym_DQUOTE),
	1724: uint16(214),
	1725: uint16(1),
	1726: uint16(anon_sym_SQUOTE),
	1727: uint16(226),
	1728: uint16(1),
	1729: uint16(sym_hex_lit),
	1730: uint16(275),
	1731: uint16(1),
	1732: uint16(sym_constant),
	1733: uint16(204),
	1734: uint16(2),
	1735: uint16(sym_true),
	1736: uint16(sym_false),
	1737: uint16(222),
	1738: uint16(2),
	1739: uint16(anon_sym_DASH),
	1740: uint16(anon_sym_PLUS),
	1741: uint16(224),
	1742: uint16(2),
	1743: uint16(sym_decimal_lit),
	1744: uint16(sym_octal_lit),
	1745: uint16(96),
	1746: uint16(5),
	1747: uint16(sym_block_lit),
	1748: uint16(sym_full_ident),
	1749: uint16(sym_bool),
	1750: uint16(sym_int_lit),
	1751: uint16(sym_string),
	1752: uint16(12),
	1753: uint16(3),
	1754: uint16(1),
	1755: uint16(sym_comment),
	1756: uint16(194),
	1757: uint16(1),
	1758: uint16(anon_sym_LBRACE),
	1759: uint16(202),
	1760: uint16(1),
	1761: uint16(sym_identifier),
	1762: uint16(210),
	1763: uint16(1),
	1764: uint16(sym_float_lit),
	1765: uint16(212),
	1766: uint16(1),
	1767: uint16(anon_sym_DQUOTE),
	1768: uint16(214),
	1769: uint16(1),
	1770: uint16(anon_sym_SQUOTE),
	1771: uint16(226),
	1772: uint16(1),
	1773: uint16(sym_hex_lit),
	1774: uint16(276),
	1775: uint16(1),
	1776: uint16(sym_constant),
	1777: uint16(204),
	1778: uint16(2),
	1779: uint16(sym_true),
	1780: uint16(sym_false),
	1781: uint16(222),
	1782: uint16(2),
	1783: uint16(anon_sym_DASH),
	1784: uint16(anon_sym_PLUS),
	1785: uint16(224),
	1786: uint16(2),
	1787: uint16(sym_decimal_lit),
	1788: uint16(sym_octal_lit),
	1789: uint16(96),
	1790: uint16(5),
	1791: uint16(sym_block_lit),
	1792: uint16(sym_full_ident),
	1793: uint16(sym_bool),
	1794: uint16(sym_int_lit),
	1795: uint16(sym_string),
	1796: uint16(12),
	1797: uint16(3),
	1798: uint16(1),
	1799: uint16(sym_comment),
	1800: uint16(194),
	1801: uint16(1),
	1802: uint16(anon_sym_LBRACE),
	1803: uint16(202),
	1804: uint16(1),
	1805: uint16(sym_identifier),
	1806: uint16(210),
	1807: uint16(1),
	1808: uint16(sym_float_lit),
	1809: uint16(212),
	1810: uint16(1),
	1811: uint16(anon_sym_DQUOTE),
	1812: uint16(214),
	1813: uint16(1),
	1814: uint16(anon_sym_SQUOTE),
	1815: uint16(226),
	1816: uint16(1),
	1817: uint16(sym_hex_lit),
	1818: uint16(272),
	1819: uint16(1),
	1820: uint16(sym_constant),
	1821: uint16(204),
	1822: uint16(2),
	1823: uint16(sym_true),
	1824: uint16(sym_false),
	1825: uint16(222),
	1826: uint16(2),
	1827: uint16(anon_sym_DASH),
	1828: uint16(anon_sym_PLUS),
	1829: uint16(224),
	1830: uint16(2),
	1831: uint16(sym_decimal_lit),
	1832: uint16(sym_octal_lit),
	1833: uint16(96),
	1834: uint16(5),
	1835: uint16(sym_block_lit),
	1836: uint16(sym_full_ident),
	1837: uint16(sym_bool),
	1838: uint16(sym_int_lit),
	1839: uint16(sym_string),
	1840: uint16(12),
	1841: uint16(3),
	1842: uint16(1),
	1843: uint16(sym_comment),
	1844: uint16(194),
	1845: uint16(1),
	1846: uint16(anon_sym_LBRACE),
	1847: uint16(202),
	1848: uint16(1),
	1849: uint16(sym_identifier),
	1850: uint16(210),
	1851: uint16(1),
	1852: uint16(sym_float_lit),
	1853: uint16(212),
	1854: uint16(1),
	1855: uint16(anon_sym_DQUOTE),
	1856: uint16(214),
	1857: uint16(1),
	1858: uint16(anon_sym_SQUOTE),
	1859: uint16(226),
	1860: uint16(1),
	1861: uint16(sym_hex_lit),
	1862: uint16(186),
	1863: uint16(1),
	1864: uint16(sym_constant),
	1865: uint16(204),
	1866: uint16(2),
	1867: uint16(sym_true),
	1868: uint16(sym_false),
	1869: uint16(222),
	1870: uint16(2),
	1871: uint16(anon_sym_DASH),
	1872: uint16(anon_sym_PLUS),
	1873: uint16(224),
	1874: uint16(2),
	1875: uint16(sym_decimal_lit),
	1876: uint16(sym_octal_lit),
	1877: uint16(96),
	1878: uint16(5),
	1879: uint16(sym_block_lit),
	1880: uint16(sym_full_ident),
	1881: uint16(sym_bool),
	1882: uint16(sym_int_lit),
	1883: uint16(sym_string),
	1884: uint16(12),
	1885: uint16(3),
	1886: uint16(1),
	1887: uint16(sym_comment),
	1888: uint16(194),
	1889: uint16(1),
	1890: uint16(anon_sym_LBRACE),
	1891: uint16(202),
	1892: uint16(1),
	1893: uint16(sym_identifier),
	1894: uint16(210),
	1895: uint16(1),
	1896: uint16(sym_float_lit),
	1897: uint16(212),
	1898: uint16(1),
	1899: uint16(anon_sym_DQUOTE),
	1900: uint16(214),
	1901: uint16(1),
	1902: uint16(anon_sym_SQUOTE),
	1903: uint16(226),
	1904: uint16(1),
	1905: uint16(sym_hex_lit),
	1906: uint16(168),
	1907: uint16(1),
	1908: uint16(sym_constant),
	1909: uint16(204),
	1910: uint16(2),
	1911: uint16(sym_true),
	1912: uint16(sym_false),
	1913: uint16(222),
	1914: uint16(2),
	1915: uint16(anon_sym_DASH),
	1916: uint16(anon_sym_PLUS),
	1917: uint16(224),
	1918: uint16(2),
	1919: uint16(sym_decimal_lit),
	1920: uint16(sym_octal_lit),
	1921: uint16(96),
	1922: uint16(5),
	1923: uint16(sym_block_lit),
	1924: uint16(sym_full_ident),
	1925: uint16(sym_bool),
	1926: uint16(sym_int_lit),
	1927: uint16(sym_string),
	1928: uint16(12),
	1929: uint16(3),
	1930: uint16(1),
	1931: uint16(sym_comment),
	1932: uint16(194),
	1933: uint16(1),
	1934: uint16(anon_sym_LBRACE),
	1935: uint16(202),
	1936: uint16(1),
	1937: uint16(sym_identifier),
	1938: uint16(210),
	1939: uint16(1),
	1940: uint16(sym_float_lit),
	1941: uint16(212),
	1942: uint16(1),
	1943: uint16(anon_sym_DQUOTE),
	1944: uint16(214),
	1945: uint16(1),
	1946: uint16(anon_sym_SQUOTE),
	1947: uint16(226),
	1948: uint16(1),
	1949: uint16(sym_hex_lit),
	1950: uint16(194),
	1951: uint16(1),
	1952: uint16(sym_constant),
	1953: uint16(204),
	1954: uint16(2),
	1955: uint16(sym_true),
	1956: uint16(sym_false),
	1957: uint16(222),
	1958: uint16(2),
	1959: uint16(anon_sym_DASH),
	1960: uint16(anon_sym_PLUS),
	1961: uint16(224),
	1962: uint16(2),
	1963: uint16(sym_decimal_lit),
	1964: uint16(sym_octal_lit),
	1965: uint16(96),
	1966: uint16(5),
	1967: uint16(sym_block_lit),
	1968: uint16(sym_full_ident),
	1969: uint16(sym_bool),
	1970: uint16(sym_int_lit),
	1971: uint16(sym_string),
	1972: uint16(12),
	1973: uint16(3),
	1974: uint16(1),
	1975: uint16(sym_comment),
	1976: uint16(194),
	1977: uint16(1),
	1978: uint16(anon_sym_LBRACE),
	1979: uint16(202),
	1980: uint16(1),
	1981: uint16(sym_identifier),
	1982: uint16(210),
	1983: uint16(1),
	1984: uint16(sym_float_lit),
	1985: uint16(212),
	1986: uint16(1),
	1987: uint16(anon_sym_DQUOTE),
	1988: uint16(214),
	1989: uint16(1),
	1990: uint16(anon_sym_SQUOTE),
	1991: uint16(226),
	1992: uint16(1),
	1993: uint16(sym_hex_lit),
	1994: uint16(233),
	1995: uint16(1),
	1996: uint16(sym_constant),
	1997: uint16(204),
	1998: uint16(2),
	1999: uint16(sym_true),
	2000: uint16(sym_false),
	2001: uint16(222),
	2002: uint16(2),
	2003: uint16(anon_sym_DASH),
	2004: uint16(anon_sym_PLUS),
	2005: uint16(224),
	2006: uint16(2),
	2007: uint16(sym_decimal_lit),
	2008: uint16(sym_octal_lit),
	2009: uint16(96),
	2010: uint16(5),
	2011: uint16(sym_block_lit),
	2012: uint16(sym_full_ident),
	2013: uint16(sym_bool),
	2014: uint16(sym_int_lit),
	2015: uint16(sym_string),
	2016: uint16(10),
	2017: uint16(3),
	2018: uint16(1),
	2019: uint16(sym_comment),
	2020: uint16(228),
	2021: uint16(1),
	2023: uint16(230),
	2024: uint16(1),
	2025: uint16(anon_sym_SEMI),
	2026: uint16(232),
	2027: uint16(1),
	2028: uint16(anon_sym_import),
	2029: uint16(234),
	2030: uint16(1),
	2031: uint16(anon_sym_package),
	2032: uint16(236),
	2033: uint16(1),
	2034: uint16(anon_sym_option),
	2035: uint16(238),
	2036: uint16(1),
	2037: uint16(anon_sym_enum),
	2038: uint16(240),
	2039: uint16(1),
	2040: uint16(anon_sym_message),
	2041: uint16(242),
	2042: uint16(1),
	2043: uint16(anon_sym_service),
	2044: uint16(51),
	2045: uint16(8),
	2046: uint16(sym_empty_statement),
	2047: uint16(sym_import),
	2048: uint16(sym_package),
	2049: uint16(sym_option),
	2050: uint16(sym_enum),
	2051: uint16(sym_message),
	2052: uint16(sym_service),
	2053: uint16(aux_sym_source_file_repeat1),
	2054: uint16(10),
	2055: uint16(3),
	2056: uint16(1),
	2057: uint16(sym_comment),
	2058: uint16(230),
	2059: uint16(1),
	2060: uint16(anon_sym_SEMI),
	2061: uint16(232),
	2062: uint16(1),
	2063: uint16(anon_sym_import),
	2064: uint16(234),
	2065: uint16(1),
	2066: uint16(anon_sym_package),
	2067: uint16(236),
	2068: uint16(1),
	2069: uint16(anon_sym_option),
	2070: uint16(238),
	2071: uint16(1),
	2072: uint16(anon_sym_enum),
	2073: uint16(240),
	2074: uint16(1),
	2075: uint16(anon_sym_message),
	2076: uint16(242),
	2077: uint16(1),
	2078: uint16(anon_sym_service),
	2079: uint16(244),
	2080: uint16(1),
	2082: uint16(52),
	2083: uint16(8),
	2084: uint16(sym_empty_statement),
	2085: uint16(sym_import),
	2086: uint16(sym_package),
	2087: uint16(sym_option),
	2088: uint16(sym_enum),
	2089: uint16(sym_message),
	2090: uint16(sym_service),
	2091: uint16(aux_sym_source_file_repeat1),
	2092: uint16(10),
	2093: uint16(3),
	2094: uint16(1),
	2095: uint16(sym_comment),
	2096: uint16(246),
	2097: uint16(1),
	2099: uint16(248),
	2100: uint16(1),
	2101: uint16(anon_sym_SEMI),
	2102: uint16(251),
	2103: uint16(1),
	2104: uint16(anon_sym_import),
	2105: uint16(254),
	2106: uint16(1),
	2107: uint16(anon_sym_package),
	2108: uint16(257),
	2109: uint16(1),
	2110: uint16(anon_sym_option),
	2111: uint16(260),
	2112: uint16(1),
	2113: uint16(anon_sym_enum),
	2114: uint16(263),
	2115: uint16(1),
	2116: uint16(anon_sym_message),
	2117: uint16(266),
	2118: uint16(1),
	2119: uint16(anon_sym_service),
	2120: uint16(52),
	2121: uint16(8),
	2122: uint16(sym_empty_statement),
	2123: uint16(sym_import),
	2124: uint16(sym_package),
	2125: uint16(sym_option),
	2126: uint16(sym_enum),
	2127: uint16(sym_message),
	2128: uint16(sym_service),
	2129: uint16(aux_sym_source_file_repeat1),
	2130: uint16(3),
	2131: uint16(3),
	2132: uint16(1),
	2133: uint16(sym_comment),
	2134: uint16(285),
	2135: uint16(1),
	2136: uint16(sym_key_type),
	2137: uint16(269),
	2138: uint16(12),
	2139: uint16(anon_sym_int32),
	2140: uint16(anon_sym_int64),
	2141: uint16(anon_sym_uint32),
	2142: uint16(anon_sym_uint64),
	2143: uint16(anon_sym_sint32),
	2144: uint16(anon_sym_sint64),
	2145: uint16(anon_sym_fixed32),
	2146: uint16(anon_sym_fixed64),
	2147: uint16(anon_sym_sfixed32),
	2148: uint16(anon_sym_sfixed64),
	2149: uint16(anon_sym_bool),
	2150: uint16(anon_sym_string),
	2151: uint16(2),
	2152: uint16(3),
	2153: uint16(1),
	2154: uint16(sym_comment),
	2155: uint16(123),
	2156: uint16(10),
	2158: uint16(anon_sym_SEMI),
	2159: uint16(anon_sym_import),
	2160: uint16(anon_sym_package),
	2161: uint16(anon_sym_option),
	2162: uint16(anon_sym_enum),
	2163: uint16(anon_sym_RBRACE),
	2164: uint16(anon_sym_message),
	2165: uint16(anon_sym_service),
	2166: uint16(anon_sym_rpc),
	2167: uint16(2),
	2168: uint16(3),
	2169: uint16(1),
	2170: uint16(sym_comment),
	2171: uint16(115),
	2172: uint16(10),
	2174: uint16(anon_sym_SEMI),
	2175: uint16(anon_sym_import),
	2176: uint16(anon_sym_package),
	2177: uint16(anon_sym_option),
	2178: uint16(anon_sym_enum),
	2179: uint16(anon_sym_RBRACE),
	2180: uint16(anon_sym_message),
	2181: uint16(anon_sym_service),
	2182: uint16(anon_sym_rpc),
	2183: uint16(4),
	2184: uint16(3),
	2185: uint16(1),
	2186: uint16(sym_comment),
	2187: uint16(273),
	2188: uint16(1),
	2189: uint16(anon_sym_DOT),
	2190: uint16(56),
	2191: uint16(1),
	2192: uint16(aux_sym__option_name_repeat1),
	2193: uint16(271),
	2194: uint16(7),
	2195: uint16(anon_sym_SEMI),
	2196: uint16(anon_sym_EQ),
	2197: uint16(anon_sym_RPAREN),
	2198: uint16(anon_sym_RBRACE),
	2199: uint16(anon_sym_COMMA),
	2200: uint16(anon_sym_RBRACK),
	2201: uint16(sym_identifier),
	2202: uint16(4),
	2203: uint16(3),
	2204: uint16(1),
	2205: uint16(sym_comment),
	2206: uint16(278),
	2207: uint16(1),
	2208: uint16(anon_sym_DOT),
	2209: uint16(56),
	2210: uint16(1),
	2211: uint16(aux_sym__option_name_repeat1),
	2212: uint16(276),
	2213: uint16(6),
	2214: uint16(anon_sym_SEMI),
	2215: uint16(anon_sym_RPAREN),
	2216: uint16(anon_sym_RBRACE),
	2217: uint16(anon_sym_COMMA),
	2218: uint16(anon_sym_RBRACK),
	2219: uint16(sym_identifier),
	2220: uint16(6),
	2221: uint16(3),
	2222: uint16(1),
	2223: uint16(sym_comment),
	2224: uint16(280),
	2225: uint16(1),
	2226: uint16(anon_sym_SEMI),
	2227: uint16(282),
	2228: uint16(1),
	2229: uint16(anon_sym_option),
	2230: uint16(284),
	2231: uint16(1),
	2232: uint16(anon_sym_RBRACE),
	2233: uint16(286),
	2234: uint16(1),
	2235: uint16(sym_identifier),
	2236: uint16(61),
	2237: uint16(4),
	2238: uint16(sym_empty_statement),
	2239: uint16(sym_option),
	2240: uint16(sym_enum_field),
	2241: uint16(aux_sym_enum_body_repeat1),
	2242: uint16(2),
	2243: uint16(3),
	2244: uint16(1),
	2245: uint16(sym_comment),
	2246: uint16(288),
	2247: uint16(8),
	2249: uint16(anon_sym_SEMI),
	2250: uint16(anon_sym_import),
	2251: uint16(anon_sym_package),
	2252: uint16(anon_sym_option),
	2253: uint16(anon_sym_enum),
	2254: uint16(anon_sym_message),
	2255: uint16(anon_sym_service),
	2256: uint16(2),
	2257: uint16(3),
	2258: uint16(1),
	2259: uint16(sym_comment),
	2260: uint16(143),
	2261: uint16(8),
	2263: uint16(anon_sym_SEMI),
	2264: uint16(anon_sym_import),
	2265: uint16(anon_sym_package),
	2266: uint16(anon_sym_option),
	2267: uint16(anon_sym_enum),
	2268: uint16(anon_sym_message),
	2269: uint16(anon_sym_service),
	2270: uint16(6),
	2271: uint16(3),
	2272: uint16(1),
	2273: uint16(sym_comment),
	2274: uint16(280),
	2275: uint16(1),
	2276: uint16(anon_sym_SEMI),
	2277: uint16(282),
	2278: uint16(1),
	2279: uint16(anon_sym_option),
	2280: uint16(286),
	2281: uint16(1),
	2282: uint16(sym_identifier),
	2283: uint16(290),
	2284: uint16(1),
	2285: uint16(anon_sym_RBRACE),
	2286: uint16(77),
	2287: uint16(4),
	2288: uint16(sym_empty_statement),
	2289: uint16(sym_option),
	2290: uint16(sym_enum_field),
	2291: uint16(aux_sym_enum_body_repeat1),
	2292: uint16(2),
	2293: uint16(3),
	2294: uint16(1),
	2295: uint16(sym_comment),
	2296: uint16(139),
	2297: uint16(8),
	2299: uint16(anon_sym_SEMI),
	2300: uint16(anon_sym_import),
	2301: uint16(anon_sym_package),
	2302: uint16(anon_sym_option),
	2303: uint16(anon_sym_enum),
	2304: uint16(anon_sym_message),
	2305: uint16(anon_sym_service),
	2306: uint16(2),
	2307: uint16(3),
	2308: uint16(1),
	2309: uint16(sym_comment),
	2310: uint16(81),
	2311: uint16(8),
	2313: uint16(anon_sym_SEMI),
	2314: uint16(anon_sym_import),
	2315: uint16(anon_sym_package),
	2316: uint16(anon_sym_option),
	2317: uint16(anon_sym_enum),
	2318: uint16(anon_sym_message),
	2319: uint16(anon_sym_service),
	2320: uint16(6),
	2321: uint16(3),
	2322: uint16(1),
	2323: uint16(sym_comment),
	2324: uint16(230),
	2325: uint16(1),
	2326: uint16(anon_sym_SEMI),
	2327: uint16(236),
	2328: uint16(1),
	2329: uint16(anon_sym_option),
	2330: uint16(292),
	2331: uint16(1),
	2332: uint16(anon_sym_RBRACE),
	2333: uint16(294),
	2334: uint16(1),
	2335: uint16(anon_sym_rpc),
	2336: uint16(68),
	2337: uint16(4),
	2338: uint16(sym_empty_statement),
	2339: uint16(sym_option),
	2340: uint16(sym_rpc),
	2341: uint16(aux_sym_service_repeat1),
	2342: uint16(2),
	2343: uint16(3),
	2344: uint16(1),
	2345: uint16(sym_comment),
	2346: uint16(77),
	2347: uint16(8),
	2349: uint16(anon_sym_SEMI),
	2350: uint16(anon_sym_import),
	2351: uint16(anon_sym_package),
	2352: uint16(anon_sym_option),
	2353: uint16(anon_sym_enum),
	2354: uint16(anon_sym_message),
	2355: uint16(anon_sym_service),
	2356: uint16(2),
	2357: uint16(3),
	2358: uint16(1),
	2359: uint16(sym_comment),
	2360: uint16(296),
	2361: uint16(8),
	2363: uint16(anon_sym_SEMI),
	2364: uint16(anon_sym_import),
	2365: uint16(anon_sym_package),
	2366: uint16(anon_sym_option),
	2367: uint16(anon_sym_enum),
	2368: uint16(anon_sym_message),
	2369: uint16(anon_sym_service),
	2370: uint16(7),
	2371: uint16(3),
	2372: uint16(1),
	2373: uint16(sym_comment),
	2374: uint16(224),
	2375: uint16(1),
	2376: uint16(sym_octal_lit),
	2377: uint16(298),
	2378: uint16(1),
	2379: uint16(sym_identifier),
	2380: uint16(165),
	2381: uint16(1),
	2382: uint16(sym_int_lit),
	2383: uint16(167),
	2384: uint16(1),
	2385: uint16(sym_range),
	2386: uint16(226),
	2387: uint16(2),
	2388: uint16(sym_decimal_lit),
	2389: uint16(sym_hex_lit),
	2390: uint16(264),
	2391: uint16(2),
	2392: uint16(sym_ranges),
	2393: uint16(sym_field_names),
	2394: uint16(6),
	2395: uint16(3),
	2396: uint16(1),
	2397: uint16(sym_comment),
	2398: uint16(230),
	2399: uint16(1),
	2400: uint16(anon_sym_SEMI),
	2401: uint16(236),
	2402: uint16(1),
	2403: uint16(anon_sym_option),
	2404: uint16(294),
	2405: uint16(1),
	2406: uint16(anon_sym_rpc),
	2407: uint16(300),
	2408: uint16(1),
	2409: uint16(anon_sym_RBRACE),
	2410: uint16(70),
	2411: uint16(4),
	2412: uint16(sym_empty_statement),
	2413: uint16(sym_option),
	2414: uint16(sym_rpc),
	2415: uint16(aux_sym_service_repeat1),
	2416: uint16(2),
	2417: uint16(3),
	2418: uint16(1),
	2419: uint16(sym_comment),
	2420: uint16(302),
	2421: uint16(8),
	2423: uint16(anon_sym_SEMI),
	2424: uint16(anon_sym_import),
	2425: uint16(anon_sym_package),
	2426: uint16(anon_sym_option),
	2427: uint16(anon_sym_enum),
	2428: uint16(anon_sym_message),
	2429: uint16(anon_sym_service),
	2430: uint16(6),
	2431: uint16(3),
	2432: uint16(1),
	2433: uint16(sym_comment),
	2434: uint16(304),
	2435: uint16(1),
	2436: uint16(anon_sym_SEMI),
	2437: uint16(307),
	2438: uint16(1),
	2439: uint16(anon_sym_option),
	2440: uint16(310),
	2441: uint16(1),
	2442: uint16(anon_sym_RBRACE),
	2443: uint16(312),
	2444: uint16(1),
	2445: uint16(anon_sym_rpc),
	2446: uint16(70),
	2447: uint16(4),
	2448: uint16(sym_empty_statement),
	2449: uint16(sym_option),
	2450: uint16(sym_rpc),
	2451: uint16(aux_sym_service_repeat1),
	2452: uint16(2),
	2453: uint16(3),
	2454: uint16(1),
	2455: uint16(sym_comment),
	2456: uint16(164),
	2457: uint16(8),
	2459: uint16(anon_sym_SEMI),
	2460: uint16(anon_sym_import),
	2461: uint16(anon_sym_package),
	2462: uint16(anon_sym_option),
	2463: uint16(anon_sym_enum),
	2464: uint16(anon_sym_message),
	2465: uint16(anon_sym_service),
	2466: uint16(2),
	2467: uint16(3),
	2468: uint16(1),
	2469: uint16(sym_comment),
	2470: uint16(315),
	2471: uint16(8),
	2473: uint16(anon_sym_SEMI),
	2474: uint16(anon_sym_import),
	2475: uint16(anon_sym_package),
	2476: uint16(anon_sym_option),
	2477: uint16(anon_sym_enum),
	2478: uint16(anon_sym_message),
	2479: uint16(anon_sym_service),
	2480: uint16(2),
	2481: uint16(3),
	2482: uint16(1),
	2483: uint16(sym_comment),
	2484: uint16(317),
	2485: uint16(8),
	2487: uint16(anon_sym_SEMI),
	2488: uint16(anon_sym_import),
	2489: uint16(anon_sym_package),
	2490: uint16(anon_sym_option),
	2491: uint16(anon_sym_enum),
	2492: uint16(anon_sym_message),
	2493: uint16(anon_sym_service),
	2494: uint16(4),
	2495: uint16(3),
	2496: uint16(1),
	2497: uint16(sym_comment),
	2498: uint16(278),
	2499: uint16(1),
	2500: uint16(anon_sym_DOT),
	2501: uint16(57),
	2502: uint16(1),
	2503: uint16(aux_sym__option_name_repeat1),
	2504: uint16(319),
	2505: uint16(6),
	2506: uint16(anon_sym_SEMI),
	2507: uint16(anon_sym_RPAREN),
	2508: uint16(anon_sym_RBRACE),
	2509: uint16(anon_sym_COMMA),
	2510: uint16(anon_sym_RBRACK),
	2511: uint16(sym_identifier),
	2512: uint16(2),
	2513: uint16(3),
	2514: uint16(1),
	2515: uint16(sym_comment),
	2516: uint16(321),
	2517: uint16(8),
	2519: uint16(anon_sym_SEMI),
	2520: uint16(anon_sym_import),
	2521: uint16(anon_sym_package),
	2522: uint16(anon_sym_option),
	2523: uint16(anon_sym_enum),
	2524: uint16(anon_sym_message),
	2525: uint16(anon_sym_service),
	2526: uint16(2),
	2527: uint16(3),
	2528: uint16(1),
	2529: uint16(sym_comment),
	2530: uint16(135),
	2531: uint16(8),
	2533: uint16(anon_sym_SEMI),
	2534: uint16(anon_sym_import),
	2535: uint16(anon_sym_package),
	2536: uint16(anon_sym_option),
	2537: uint16(anon_sym_enum),
	2538: uint16(anon_sym_message),
	2539: uint16(anon_sym_service),
	2540: uint16(6),
	2541: uint16(3),
	2542: uint16(1),
	2543: uint16(sym_comment),
	2544: uint16(323),
	2545: uint16(1),
	2546: uint16(anon_sym_SEMI),
	2547: uint16(326),
	2548: uint16(1),
	2549: uint16(anon_sym_option),
	2550: uint16(329),
	2551: uint16(1),
	2552: uint16(anon_sym_RBRACE),
	2553: uint16(331),
	2554: uint16(1),
	2555: uint16(sym_identifier),
	2556: uint16(77),
	2557: uint16(4),
	2558: uint16(sym_empty_statement),
	2559: uint16(sym_option),
	2560: uint16(sym_enum_field),
	2561: uint16(aux_sym_enum_body_repeat1),
	2562: uint16(6),
	2563: uint16(3),
	2564: uint16(1),
	2565: uint16(sym_comment),
	2566: uint16(280),
	2567: uint16(1),
	2568: uint16(anon_sym_SEMI),
	2569: uint16(282),
	2570: uint16(1),
	2571: uint16(anon_sym_option),
	2572: uint16(286),
	2573: uint16(1),
	2574: uint16(sym_identifier),
	2575: uint16(334),
	2576: uint16(1),
	2577: uint16(anon_sym_RBRACE),
	2578: uint16(77),
	2579: uint16(4),
	2580: uint16(sym_empty_statement),
	2581: uint16(sym_option),
	2582: uint16(sym_enum_field),
	2583: uint16(aux_sym_enum_body_repeat1),
	2584: uint16(2),
	2585: uint16(3),
	2586: uint16(1),
	2587: uint16(sym_comment),
	2588: uint16(271),
	2589: uint16(8),
	2590: uint16(anon_sym_SEMI),
	2591: uint16(anon_sym_EQ),
	2592: uint16(anon_sym_RPAREN),
	2593: uint16(anon_sym_DOT),
	2594: uint16(anon_sym_RBRACE),
	2595: uint16(anon_sym_COMMA),
	2596: uint16(anon_sym_RBRACK),
	2597: uint16(sym_identifier),
	2598: uint16(6),
	2599: uint16(3),
	2600: uint16(1),
	2601: uint16(sym_comment),
	2602: uint16(280),
	2603: uint16(1),
	2604: uint16(anon_sym_SEMI),
	2605: uint16(282),
	2606: uint16(1),
	2607: uint16(anon_sym_option),
	2608: uint16(286),
	2609: uint16(1),
	2610: uint16(sym_identifier),
	2611: uint16(336),
	2612: uint16(1),
	2613: uint16(anon_sym_RBRACE),
	2614: uint16(78),
	2615: uint16(4),
	2616: uint16(sym_empty_statement),
	2617: uint16(sym_option),
	2618: uint16(sym_enum_field),
	2619: uint16(aux_sym_enum_body_repeat1),
	2620: uint16(5),
	2621: uint16(3),
	2622: uint16(1),
	2623: uint16(sym_comment),
	2624: uint16(230),
	2625: uint16(1),
	2626: uint16(anon_sym_SEMI),
	2627: uint16(236),
	2628: uint16(1),
	2629: uint16(anon_sym_option),
	2630: uint16(338),
	2631: uint16(1),
	2632: uint16(anon_sym_RBRACE),
	2633: uint16(86),
	2634: uint16(3),
	2635: uint16(sym_empty_statement),
	2636: uint16(sym_option),
	2637: uint16(aux_sym_rpc_repeat1),
	2638: uint16(5),
	2639: uint16(3),
	2640: uint16(1),
	2641: uint16(sym_comment),
	2642: uint16(340),
	2643: uint16(1),
	2644: uint16(anon_sym_SEMI),
	2645: uint16(343),
	2646: uint16(1),
	2647: uint16(anon_sym_option),
	2648: uint16(346),
	2649: uint16(1),
	2650: uint16(anon_sym_RBRACE),
	2651: uint16(82),
	2652: uint16(3),
	2653: uint16(sym_empty_statement),
	2654: uint16(sym_option),
	2655: uint16(aux_sym_rpc_repeat1),
	2656: uint16(5),
	2657: uint16(3),
	2658: uint16(1),
	2659: uint16(sym_comment),
	2660: uint16(230),
	2661: uint16(1),
	2662: uint16(anon_sym_SEMI),
	2663: uint16(236),
	2664: uint16(1),
	2665: uint16(anon_sym_option),
	2666: uint16(348),
	2667: uint16(1),
	2668: uint16(anon_sym_RBRACE),
	2669: uint16(82),
	2670: uint16(3),
	2671: uint16(sym_empty_statement),
	2672: uint16(sym_option),
	2673: uint16(aux_sym_rpc_repeat1),
	2674: uint16(5),
	2675: uint16(3),
	2676: uint16(1),
	2677: uint16(sym_comment),
	2678: uint16(230),
	2679: uint16(1),
	2680: uint16(anon_sym_SEMI),
	2681: uint16(236),
	2682: uint16(1),
	2683: uint16(anon_sym_option),
	2684: uint16(350),
	2685: uint16(1),
	2686: uint16(anon_sym_RBRACE),
	2687: uint16(82),
	2688: uint16(3),
	2689: uint16(sym_empty_statement),
	2690: uint16(sym_option),
	2691: uint16(aux_sym_rpc_repeat1),
	2692: uint16(5),
	2693: uint16(3),
	2694: uint16(1),
	2695: uint16(sym_comment),
	2696: uint16(230),
	2697: uint16(1),
	2698: uint16(anon_sym_SEMI),
	2699: uint16(236),
	2700: uint16(1),
	2701: uint16(anon_sym_option),
	2702: uint16(350),
	2703: uint16(1),
	2704: uint16(anon_sym_RBRACE),
	2705: uint16(83),
	2706: uint16(3),
	2707: uint16(sym_empty_statement),
	2708: uint16(sym_option),
	2709: uint16(aux_sym_rpc_repeat1),
	2710: uint16(5),
	2711: uint16(3),
	2712: uint16(1),
	2713: uint16(sym_comment),
	2714: uint16(230),
	2715: uint16(1),
	2716: uint16(anon_sym_SEMI),
	2717: uint16(236),
	2718: uint16(1),
	2719: uint16(anon_sym_option),
	2720: uint16(352),
	2721: uint16(1),
	2722: uint16(anon_sym_RBRACE),
	2723: uint16(82),
	2724: uint16(3),
	2725: uint16(sym_empty_statement),
	2726: uint16(sym_option),
	2727: uint16(aux_sym_rpc_repeat1),
	2728: uint16(5),
	2729: uint16(3),
	2730: uint16(1),
	2731: uint16(sym_comment),
	2732: uint16(230),
	2733: uint16(1),
	2734: uint16(anon_sym_SEMI),
	2735: uint16(236),
	2736: uint16(1),
	2737: uint16(anon_sym_option),
	2738: uint16(352),
	2739: uint16(1),
	2740: uint16(anon_sym_RBRACE),
	2741: uint16(84),
	2742: uint16(3),
	2743: uint16(sym_empty_statement),
	2744: uint16(sym_option),
	2745: uint16(aux_sym_rpc_repeat1),
	2746: uint16(5),
	2747: uint16(3),
	2748: uint16(1),
	2749: uint16(sym_comment),
	2750: uint16(226),
	2751: uint16(1),
	2752: uint16(sym_hex_lit),
	2753: uint16(354),
	2754: uint16(1),
	2755: uint16(sym_float_lit),
	2756: uint16(94),
	2757: uint16(1),
	2758: uint16(sym_int_lit),
	2759: uint16(224),
	2760: uint16(2),
	2761: uint16(sym_decimal_lit),
	2762: uint16(sym_octal_lit),
	2763: uint16(6),
	2764: uint16(3),
	2765: uint16(1),
	2766: uint16(sym_comment),
	2767: uint16(356),
	2768: uint16(1),
	2769: uint16(anon_sym_LPAREN),
	2770: uint16(358),
	2771: uint16(1),
	2772: uint16(sym_identifier),
	2773: uint16(152),
	2774: uint16(1),
	2775: uint16(sym_field_option),
	2776: uint16(250),
	2777: uint16(1),
	2778: uint16(sym_field_options),
	2779: uint16(251),
	2780: uint16(1),
	2781: uint16(sym__option_name),
	2782: uint16(6),
	2783: uint16(3),
	2784: uint16(1),
	2785: uint16(sym_comment),
	2786: uint16(356),
	2787: uint16(1),
	2788: uint16(anon_sym_LPAREN),
	2789: uint16(358),
	2790: uint16(1),
	2791: uint16(sym_identifier),
	2792: uint16(152),
	2793: uint16(1),
	2794: uint16(sym_field_option),
	2795: uint16(241),
	2796: uint16(1),
	2797: uint16(sym_field_options),
	2798: uint16(251),
	2799: uint16(1),
	2800: uint16(sym__option_name),
	2801: uint16(2),
	2802: uint16(3),
	2803: uint16(1),
	2804: uint16(sym_comment),
	2805: uint16(360),
	2806: uint16(5),
	2807: uint16(anon_sym_SEMI),
	2808: uint16(anon_sym_RBRACE),
	2809: uint16(anon_sym_COMMA),
	2810: uint16(anon_sym_RBRACK),
	2811: uint16(sym_identifier),
	2812: uint16(6),
	2813: uint16(3),
	2814: uint16(1),
	2815: uint16(sym_comment),
	2816: uint16(356),
	2817: uint16(1),
	2818: uint16(anon_sym_LPAREN),
	2819: uint16(358),
	2820: uint16(1),
	2821: uint16(sym_identifier),
	2822: uint16(152),
	2823: uint16(1),
	2824: uint16(sym_field_option),
	2825: uint16(234),
	2826: uint16(1),
	2827: uint16(sym_field_options),
	2828: uint16(251),
	2829: uint16(1),
	2830: uint16(sym__option_name),
	2831: uint16(5),
	2832: uint16(3),
	2833: uint16(1),
	2834: uint16(sym_comment),
	2835: uint16(224),
	2836: uint16(1),
	2837: uint16(sym_octal_lit),
	2838: uint16(32),
	2839: uint16(1),
	2840: uint16(sym_int_lit),
	2841: uint16(215),
	2842: uint16(1),
	2843: uint16(sym_field_number),
	2844: uint16(226),
	2845: uint16(2),
	2846: uint16(sym_decimal_lit),
	2847: uint16(sym_hex_lit),
	2848: uint16(2),
	2849: uint16(3),
	2850: uint16(1),
	2851: uint16(sym_comment),
	2852: uint16(362),
	2853: uint16(5),
	2854: uint16(anon_sym_SEMI),
	2855: uint16(anon_sym_RBRACE),
	2856: uint16(anon_sym_COMMA),
	2857: uint16(anon_sym_RBRACK),
	2858: uint16(sym_identifier),
	2859: uint16(5),
	2860: uint16(3),
	2861: uint16(1),
	2862: uint16(sym_comment),
	2863: uint16(224),
	2864: uint16(1),
	2865: uint16(sym_octal_lit),
	2866: uint16(32),
	2867: uint16(1),
	2868: uint16(sym_int_lit),
	2869: uint16(185),
	2870: uint16(1),
	2871: uint16(sym_field_number),
	2872: uint16(226),
	2873: uint16(2),
	2874: uint16(sym_decimal_lit),
	2875: uint16(sym_hex_lit),
	2876: uint16(2),
	2877: uint16(3),
	2878: uint16(1),
	2879: uint16(sym_comment),
	2880: uint16(364),
	2881: uint16(5),
	2882: uint16(anon_sym_SEMI),
	2883: uint16(anon_sym_RBRACE),
	2884: uint16(anon_sym_COMMA),
	2885: uint16(anon_sym_RBRACK),
	2886: uint16(sym_identifier),
	2887: uint16(5),
	2888: uint16(3),
	2889: uint16(1),
	2890: uint16(sym_comment),
	2891: uint16(224),
	2892: uint16(1),
	2893: uint16(sym_octal_lit),
	2894: uint16(366),
	2895: uint16(1),
	2896: uint16(anon_sym_DASH),
	2897: uint16(188),
	2898: uint16(1),
	2899: uint16(sym_int_lit),
	2900: uint16(226),
	2901: uint16(2),
	2902: uint16(sym_decimal_lit),
	2903: uint16(sym_hex_lit),
	2904: uint16(6),
	2905: uint16(3),
	2906: uint16(1),
	2907: uint16(sym_comment),
	2908: uint16(356),
	2909: uint16(1),
	2910: uint16(anon_sym_LPAREN),
	2911: uint16(358),
	2912: uint16(1),
	2913: uint16(sym_identifier),
	2914: uint16(152),
	2915: uint16(1),
	2916: uint16(sym_field_option),
	2917: uint16(226),
	2918: uint16(1),
	2919: uint16(sym_field_options),
	2920: uint16(251),
	2921: uint16(1),
	2922: uint16(sym__option_name),
	2923: uint16(5),
	2924: uint16(3),
	2925: uint16(1),
	2926: uint16(sym_comment),
	2927: uint16(370),
	2928: uint16(1),
	2929: uint16(sym_octal_lit),
	2930: uint16(29),
	2931: uint16(1),
	2932: uint16(sym_field_number),
	2933: uint16(32),
	2934: uint16(1),
	2935: uint16(sym_int_lit),
	2936: uint16(368),
	2937: uint16(2),
	2938: uint16(sym_decimal_lit),
	2939: uint16(sym_hex_lit),
	2940: uint16(6),
	2941: uint16(3),
	2942: uint16(1),
	2943: uint16(sym_comment),
	2944: uint16(11),
	2945: uint16(1),
	2946: uint16(anon_sym_DOT),
	2947: uint16(31),
	2948: uint16(1),
	2949: uint16(sym_identifier),
	2950: uint16(372),
	2951: uint16(1),
	2952: uint16(anon_sym_stream),
	2953: uint16(199),
	2954: uint16(1),
	2955: uint16(aux_sym_message_or_enum_type_repeat1),
	2956: uint16(222),
	2957: uint16(1),
	2958: uint16(sym_message_or_enum_type),
	2959: uint16(5),
	2960: uint16(3),
	2961: uint16(1),
	2962: uint16(sym_comment),
	2963: uint16(212),
	2964: uint16(1),
	2965: uint16(anon_sym_DQUOTE),
	2966: uint16(214),
	2967: uint16(1),
	2968: uint16(anon_sym_SQUOTE),
	2969: uint16(263),
	2970: uint16(1),
	2971: uint16(sym_string),
	2972: uint16(374),
	2973: uint16(2),
	2974: uint16(anon_sym_weak),
	2975: uint16(anon_sym_public),
	2976: uint16(2),
	2977: uint16(3),
	2978: uint16(1),
	2979: uint16(sym_comment),
	2980: uint16(186),
	2981: uint16(5),
	2982: uint16(anon_sym_SEMI),
	2983: uint16(anon_sym_LBRACK),
	2984: uint16(anon_sym_COMMA),
	2985: uint16(anon_sym_RBRACK),
	2986: uint16(anon_sym_to),
	2987: uint16(5),
	2988: uint16(3),
	2989: uint16(1),
	2990: uint16(sym_comment),
	2991: uint16(224),
	2992: uint16(1),
	2993: uint16(sym_octal_lit),
	2994: uint16(32),
	2995: uint16(1),
	2996: uint16(sym_int_lit),
	2997: uint16(192),
	2998: uint16(1),
	2999: uint16(sym_field_number),
	3000: uint16(226),
	3001: uint16(2),
	3002: uint16(sym_decimal_lit),
	3003: uint16(sym_hex_lit),
	3004: uint16(2),
	3005: uint16(3),
	3006: uint16(1),
	3007: uint16(sym_comment),
	3008: uint16(376),
	3009: uint16(5),
	3010: uint16(anon_sym_SEMI),
	3011: uint16(anon_sym_RBRACE),
	3012: uint16(anon_sym_COMMA),
	3013: uint16(anon_sym_RBRACK),
	3014: uint16(sym_identifier),
	3015: uint16(5),
	3016: uint16(3),
	3017: uint16(1),
	3018: uint16(sym_comment),
	3019: uint16(224),
	3020: uint16(1),
	3021: uint16(sym_octal_lit),
	3022: uint16(32),
	3023: uint16(1),
	3024: uint16(sym_int_lit),
	3025: uint16(197),
	3026: uint16(1),
	3027: uint16(sym_field_number),
	3028: uint16(226),
	3029: uint16(2),
	3030: uint16(sym_decimal_lit),
	3031: uint16(sym_hex_lit),
	3032: uint16(6),
	3033: uint16(3),
	3034: uint16(1),
	3035: uint16(sym_comment),
	3036: uint16(11),
	3037: uint16(1),
	3038: uint16(anon_sym_DOT),
	3039: uint16(31),
	3040: uint16(1),
	3041: uint16(sym_identifier),
	3042: uint16(378),
	3043: uint16(1),
	3044: uint16(anon_sym_stream),
	3045: uint16(199),
	3046: uint16(1),
	3047: uint16(aux_sym_message_or_enum_type_repeat1),
	3048: uint16(229),
	3049: uint16(1),
	3050: uint16(sym_message_or_enum_type),
	3051: uint16(2),
	3052: uint16(3),
	3053: uint16(1),
	3054: uint16(sym_comment),
	3055: uint16(380),
	3056: uint16(5),
	3057: uint16(anon_sym_SEMI),
	3058: uint16(anon_sym_RBRACE),
	3059: uint16(anon_sym_COMMA),
	3060: uint16(anon_sym_RBRACK),
	3061: uint16(sym_identifier),
	3062: uint16(6),
	3063: uint16(3),
	3064: uint16(1),
	3065: uint16(sym_comment),
	3066: uint16(356),
	3067: uint16(1),
	3068: uint16(anon_sym_LPAREN),
	3069: uint16(358),
	3070: uint16(1),
	3071: uint16(sym_identifier),
	3072: uint16(152),
	3073: uint16(1),
	3074: uint16(sym_field_option),
	3075: uint16(239),
	3076: uint16(1),
	3077: uint16(sym_field_options),
	3078: uint16(251),
	3079: uint16(1),
	3080: uint16(sym__option_name),
	3081: uint16(2),
	3082: uint16(3),
	3083: uint16(1),
	3084: uint16(sym_comment),
	3085: uint16(382),
	3086: uint16(5),
	3087: uint16(anon_sym_SEMI),
	3088: uint16(anon_sym_RBRACE),
	3089: uint16(anon_sym_COMMA),
	3090: uint16(anon_sym_RBRACK),
	3091: uint16(sym_identifier),
	3092: uint16(2),
	3093: uint16(3),
	3094: uint16(1),
	3095: uint16(sym_comment),
	3096: uint16(384),
	3097: uint16(5),
	3098: uint16(anon_sym_SEMI),
	3099: uint16(anon_sym_RBRACE),
	3100: uint16(anon_sym_COMMA),
	3101: uint16(anon_sym_RBRACK),
	3102: uint16(sym_identifier),
	3103: uint16(5),
	3104: uint16(3),
	3105: uint16(1),
	3106: uint16(sym_comment),
	3107: uint16(208),
	3108: uint16(1),
	3109: uint16(sym_hex_lit),
	3110: uint16(354),
	3111: uint16(1),
	3112: uint16(sym_float_lit),
	3113: uint16(94),
	3114: uint16(1),
	3115: uint16(sym_int_lit),
	3116: uint16(206),
	3117: uint16(2),
	3118: uint16(sym_decimal_lit),
	3119: uint16(sym_octal_lit),
	3120: uint16(5),
	3121: uint16(3),
	3122: uint16(1),
	3123: uint16(sym_comment),
	3124: uint16(224),
	3125: uint16(1),
	3126: uint16(sym_octal_lit),
	3127: uint16(165),
	3128: uint16(1),
	3129: uint16(sym_int_lit),
	3130: uint16(204),
	3131: uint16(1),
	3132: uint16(sym_range),
	3133: uint16(226),
	3134: uint16(2),
	3135: uint16(sym_decimal_lit),
	3136: uint16(sym_hex_lit),
	3137: uint16(6),
	3138: uint16(3),
	3139: uint16(1),
	3140: uint16(sym_comment),
	3141: uint16(11),
	3142: uint16(1),
	3143: uint16(anon_sym_DOT),
	3144: uint16(31),
	3145: uint16(1),
	3146: uint16(sym_identifier),
	3147: uint16(386),
	3148: uint16(1),
	3149: uint16(anon_sym_stream),
	3150: uint16(199),
	3151: uint16(1),
	3152: uint16(aux_sym_message_or_enum_type_repeat1),
	3153: uint16(267),
	3154: uint16(1),
	3155: uint16(sym_message_or_enum_type),
	3156: uint16(5),
	3157: uint16(3),
	3158: uint16(1),
	3159: uint16(sym_comment),
	3160: uint16(224),
	3161: uint16(1),
	3162: uint16(sym_octal_lit),
	3163: uint16(388),
	3164: uint16(1),
	3165: uint16(anon_sym_max),
	3166: uint16(201),
	3167: uint16(1),
	3168: uint16(sym_int_lit),
	3169: uint16(226),
	3170: uint16(2),
	3171: uint16(sym_decimal_lit),
	3172: uint16(sym_hex_lit),
	3173: uint16(3),
	3174: uint16(3),
	3175: uint16(1),
	3176: uint16(sym_comment),
	3177: uint16(392),
	3178: uint16(1),
	3179: uint16(anon_sym_DOT),
	3180: uint16(390),
	3181: uint16(3),
	3182: uint16(anon_sym_RPAREN),
	3183: uint16(anon_sym_GT),
	3184: uint16(sym_identifier),
	3185: uint16(5),
	3186: uint16(3),
	3187: uint16(1),
	3188: uint16(sym_comment),
	3189: uint16(11),
	3190: uint16(1),
	3191: uint16(anon_sym_DOT),
	3192: uint16(394),
	3193: uint16(1),
	3194: uint16(sym_identifier),
	3195: uint16(199),
	3196: uint16(1),
	3197: uint16(aux_sym_message_or_enum_type_repeat1),
	3198: uint16(256),
	3199: uint16(1),
	3200: uint16(sym_message_or_enum_type),
	3201: uint16(4),
	3202: uint16(396),
	3203: uint16(1),
	3204: uint16(anon_sym_DQUOTE),
	3205: uint16(400),
	3206: uint16(1),
	3207: uint16(sym_comment),
	3208: uint16(119),
	3209: uint16(1),
	3210: uint16(aux_sym_string_repeat1),
	3211: uint16(398),
	3212: uint16(2),
	3213: uint16(aux_sym_string_token1),
	3214: uint16(sym_escape_sequence),
	3215: uint16(4),
	3216: uint16(3),
	3217: uint16(1),
	3218: uint16(sym_comment),
	3219: uint16(224),
	3220: uint16(1),
	3221: uint16(sym_octal_lit),
	3222: uint16(214),
	3223: uint16(1),
	3224: uint16(sym_int_lit),
	3225: uint16(226),
	3226: uint16(2),
	3227: uint16(sym_decimal_lit),
	3228: uint16(sym_hex_lit),
	3229: uint16(4),
	3230: uint16(400),
	3231: uint16(1),
	3232: uint16(sym_comment),
	3233: uint16(402),
	3234: uint16(1),
	3235: uint16(anon_sym_DQUOTE),
	3236: uint16(119),
	3237: uint16(1),
	3238: uint16(aux_sym_string_repeat1),
	3239: uint16(404),
	3240: uint16(2),
	3241: uint16(aux_sym_string_token1),
	3242: uint16(sym_escape_sequence),
	3243: uint16(3),
	3244: uint16(3),
	3245: uint16(1),
	3246: uint16(sym_comment),
	3247: uint16(392),
	3248: uint16(1),
	3249: uint16(anon_sym_DOT),
	3250: uint16(407),
	3251: uint16(3),
	3252: uint16(anon_sym_RPAREN),
	3253: uint16(anon_sym_GT),
	3254: uint16(sym_identifier),
	3255: uint16(2),
	3256: uint16(3),
	3257: uint16(1),
	3258: uint16(sym_comment),
	3259: uint16(409),
	3260: uint16(4),
	3261: uint16(anon_sym_SEMI),
	3262: uint16(anon_sym_option),
	3263: uint16(anon_sym_RBRACE),
	3264: uint16(anon_sym_rpc),
	3265: uint16(3),
	3266: uint16(3),
	3267: uint16(1),
	3268: uint16(sym_comment),
	3269: uint16(392),
	3270: uint16(1),
	3271: uint16(anon_sym_DOT),
	3272: uint16(411),
	3273: uint16(3),
	3274: uint16(anon_sym_RPAREN),
	3275: uint16(anon_sym_GT),
	3276: uint16(sym_identifier),
	3277: uint16(2),
	3278: uint16(3),
	3279: uint16(1),
	3280: uint16(sym_comment),
	3281: uint16(413),
	3282: uint16(4),
	3283: uint16(anon_sym_SEMI),
	3284: uint16(anon_sym_option),
	3285: uint16(anon_sym_RBRACE),
	3286: uint16(anon_sym_rpc),
	3287: uint16(4),
	3288: uint16(400),
	3289: uint16(1),
	3290: uint16(sym_comment),
	3291: uint16(415),
	3292: uint16(1),
	3293: uint16(anon_sym_SQUOTE),
	3294: uint16(124),
	3295: uint16(1),
	3296: uint16(aux_sym_string_repeat2),
	3297: uint16(417),
	3298: uint16(2),
	3299: uint16(aux_sym_string_token2),
	3300: uint16(sym_escape_sequence),
	3301: uint16(5),
	3302: uint16(3),
	3303: uint16(1),
	3304: uint16(sym_comment),
	3305: uint16(356),
	3306: uint16(1),
	3307: uint16(anon_sym_LPAREN),
	3308: uint16(358),
	3309: uint16(1),
	3310: uint16(sym_identifier),
	3311: uint16(195),
	3312: uint16(1),
	3313: uint16(sym_enum_value_option),
	3314: uint16(221),
	3315: uint16(1),
	3316: uint16(sym__option_name),
	3317: uint16(3),
	3318: uint16(3),
	3319: uint16(1),
	3320: uint16(sym_comment),
	3321: uint16(420),
	3322: uint16(2),
	3323: uint16(anon_sym_SEMI),
	3324: uint16(anon_sym_COMMA),
	3325: uint16(422),
	3326: uint16(2),
	3327: uint16(anon_sym_RBRACE),
	3328: uint16(sym_identifier),
	3329: uint16(3),
	3330: uint16(3),
	3331: uint16(1),
	3332: uint16(sym_comment),
	3333: uint16(123),
	3334: uint16(2),
	3335: uint16(anon_sym_SEMI),
	3336: uint16(anon_sym_RBRACE),
	3337: uint16(125),
	3338: uint16(2),
	3339: uint16(anon_sym_option),
	3340: uint16(sym_identifier),
	3341: uint16(5),
	3342: uint16(3),
	3343: uint16(1),
	3344: uint16(sym_comment),
	3345: uint16(11),
	3346: uint16(1),
	3347: uint16(anon_sym_DOT),
	3348: uint16(394),
	3349: uint16(1),
	3350: uint16(sym_identifier),
	3351: uint16(199),
	3352: uint16(1),
	3353: uint16(aux_sym_message_or_enum_type_repeat1),
	3354: uint16(235),
	3355: uint16(1),
	3356: uint16(sym_message_or_enum_type),
	3357: uint16(2),
	3358: uint16(3),
	3359: uint16(1),
	3360: uint16(sym_comment),
	3361: uint16(424),
	3362: uint16(4),
	3363: uint16(anon_sym_SEMI),
	3364: uint16(anon_sym_option),
	3365: uint16(anon_sym_RBRACE),
	3366: uint16(anon_sym_rpc),
	3367: uint16(2),
	3368: uint16(3),
	3369: uint16(1),
	3370: uint16(sym_comment),
	3371: uint16(186),
	3372: uint16(4),
	3373: uint16(anon_sym_SEMI),
	3374: uint16(anon_sym_RBRACE),
	3375: uint16(anon_sym_COMMA),
	3376: uint16(sym_identifier),
	3377: uint16(3),
	3378: uint16(3),
	3379: uint16(1),
	3380: uint16(sym_comment),
	3381: uint16(426),
	3382: uint16(2),
	3383: uint16(anon_sym_SEMI),
	3384: uint16(anon_sym_COMMA),
	3385: uint16(428),
	3386: uint16(2),
	3387: uint16(anon_sym_RBRACE),
	3388: uint16(sym_identifier),
	3389: uint16(3),
	3390: uint16(3),
	3391: uint16(1),
	3392: uint16(sym_comment),
	3393: uint16(115),
	3394: uint16(2),
	3395: uint16(anon_sym_SEMI),
	3396: uint16(anon_sym_RBRACE),
	3397: uint16(117),
	3398: uint16(2),
	3399: uint16(anon_sym_option),
	3400: uint16(sym_identifier),
	3401: uint16(3),
	3402: uint16(3),
	3403: uint16(1),
	3404: uint16(sym_comment),
	3405: uint16(430),
	3406: uint16(2),
	3407: uint16(anon_sym_SEMI),
	3408: uint16(anon_sym_RBRACE),
	3409: uint16(432),
	3410: uint16(2),
	3411: uint16(anon_sym_option),
	3412: uint16(sym_identifier),
	3413: uint16(3),
	3414: uint16(3),
	3415: uint16(1),
	3416: uint16(sym_comment),
	3417: uint16(434),
	3418: uint16(2),
	3419: uint16(anon_sym_SEMI),
	3420: uint16(anon_sym_RBRACE),
	3421: uint16(436),
	3422: uint16(2),
	3423: uint16(anon_sym_option),
	3424: uint16(sym_identifier),
	3425: uint16(5),
	3426: uint16(3),
	3427: uint16(1),
	3428: uint16(sym_comment),
	3429: uint16(356),
	3430: uint16(1),
	3431: uint16(anon_sym_LPAREN),
	3432: uint16(358),
	3433: uint16(1),
	3434: uint16(sym_identifier),
	3435: uint16(159),
	3436: uint16(1),
	3437: uint16(sym_enum_value_option),
	3438: uint16(221),
	3439: uint16(1),
	3440: uint16(sym__option_name),
	3441: uint16(4),
	3442: uint16(396),
	3443: uint16(1),
	3444: uint16(anon_sym_SQUOTE),
	3445: uint16(400),
	3446: uint16(1),
	3447: uint16(sym_comment),
	3448: uint16(124),
	3449: uint16(1),
	3450: uint16(aux_sym_string_repeat2),
	3451: uint16(438),
	3452: uint16(2),
	3453: uint16(aux_sym_string_token2),
	3454: uint16(sym_escape_sequence),
	3455: uint16(5),
	3456: uint16(3),
	3457: uint16(1),
	3458: uint16(sym_comment),
	3459: uint16(11),
	3460: uint16(1),
	3461: uint16(anon_sym_DOT),
	3462: uint16(394),
	3463: uint16(1),
	3464: uint16(sym_identifier),
	3465: uint16(199),
	3466: uint16(1),
	3467: uint16(aux_sym_message_or_enum_type_repeat1),
	3468: uint16(229),
	3469: uint16(1),
	3470: uint16(sym_message_or_enum_type),
	3471: uint16(4),
	3472: uint16(400),
	3473: uint16(1),
	3474: uint16(sym_comment),
	3475: uint16(440),
	3476: uint16(1),
	3477: uint16(anon_sym_SQUOTE),
	3478: uint16(136),
	3479: uint16(1),
	3480: uint16(aux_sym_string_repeat2),
	3481: uint16(442),
	3482: uint16(2),
	3483: uint16(aux_sym_string_token2),
	3484: uint16(sym_escape_sequence),
	3485: uint16(3),
	3486: uint16(3),
	3487: uint16(1),
	3488: uint16(sym_comment),
	3489: uint16(444),
	3490: uint16(2),
	3491: uint16(anon_sym_SEMI),
	3492: uint16(anon_sym_RBRACE),
	3493: uint16(446),
	3494: uint16(2),
	3495: uint16(anon_sym_option),
	3496: uint16(sym_identifier),
	3497: uint16(3),
	3498: uint16(3),
	3499: uint16(1),
	3500: uint16(sym_comment),
	3501: uint16(448),
	3502: uint16(2),
	3503: uint16(anon_sym_SEMI),
	3504: uint16(anon_sym_COMMA),
	3505: uint16(450),
	3506: uint16(2),
	3507: uint16(anon_sym_RBRACE),
	3508: uint16(sym_identifier),
	3509: uint16(4),
	3510: uint16(400),
	3511: uint16(1),
	3512: uint16(sym_comment),
	3513: uint16(440),
	3514: uint16(1),
	3515: uint16(anon_sym_DQUOTE),
	3516: uint16(117),
	3517: uint16(1),
	3518: uint16(aux_sym_string_repeat1),
	3519: uint16(452),
	3520: uint16(2),
	3521: uint16(aux_sym_string_token1),
	3522: uint16(sym_escape_sequence),
	3523: uint16(5),
	3524: uint16(3),
	3525: uint16(1),
	3526: uint16(sym_comment),
	3527: uint16(356),
	3528: uint16(1),
	3529: uint16(anon_sym_LPAREN),
	3530: uint16(358),
	3531: uint16(1),
	3532: uint16(sym_identifier),
	3533: uint16(210),
	3534: uint16(1),
	3535: uint16(sym_field_option),
	3536: uint16(251),
	3537: uint16(1),
	3538: uint16(sym__option_name),
	3539: uint16(3),
	3540: uint16(3),
	3541: uint16(1),
	3542: uint16(sym_comment),
	3543: uint16(454),
	3544: uint16(2),
	3545: uint16(anon_sym_SEMI),
	3546: uint16(anon_sym_COMMA),
	3547: uint16(456),
	3548: uint16(2),
	3549: uint16(anon_sym_RBRACE),
	3550: uint16(sym_identifier),
	3551: uint16(2),
	3552: uint16(3),
	3553: uint16(1),
	3554: uint16(sym_comment),
	3555: uint16(458),
	3556: uint16(4),
	3557: uint16(anon_sym_SEMI),
	3558: uint16(anon_sym_option),
	3559: uint16(anon_sym_RBRACE),
	3560: uint16(anon_sym_rpc),
	3561: uint16(5),
	3562: uint16(3),
	3563: uint16(1),
	3564: uint16(sym_comment),
	3565: uint16(356),
	3566: uint16(1),
	3567: uint16(anon_sym_LPAREN),
	3568: uint16(358),
	3569: uint16(1),
	3570: uint16(sym_identifier),
	3571: uint16(173),
	3572: uint16(1),
	3573: uint16(sym_enum_value_option),
	3574: uint16(221),
	3575: uint16(1),
	3576: uint16(sym__option_name),
	3577: uint16(3),
	3578: uint16(3),
	3579: uint16(1),
	3580: uint16(sym_comment),
	3581: uint16(460),
	3582: uint16(2),
	3583: uint16(anon_sym_SEMI),
	3584: uint16(anon_sym_RBRACE),
	3585: uint16(462),
	3586: uint16(2),
	3587: uint16(anon_sym_option),
	3588: uint16(sym_identifier),
	3589: uint16(3),
	3590: uint16(3),
	3591: uint16(1),
	3592: uint16(sym_comment),
	3593: uint16(464),
	3594: uint16(2),
	3595: uint16(anon_sym_SEMI),
	3596: uint16(anon_sym_RBRACE),
	3597: uint16(466),
	3598: uint16(2),
	3599: uint16(anon_sym_option),
	3600: uint16(sym_identifier),
	3601: uint16(2),
	3602: uint16(3),
	3603: uint16(1),
	3604: uint16(sym_comment),
	3605: uint16(468),
	3606: uint16(4),
	3607: uint16(anon_sym_SEMI),
	3608: uint16(anon_sym_option),
	3609: uint16(anon_sym_RBRACE),
	3610: uint16(anon_sym_rpc),
	3611: uint16(3),
	3612: uint16(3),
	3613: uint16(1),
	3614: uint16(sym_comment),
	3615: uint16(470),
	3616: uint16(2),
	3617: uint16(anon_sym_SEMI),
	3618: uint16(anon_sym_COMMA),
	3619: uint16(472),
	3620: uint16(2),
	3621: uint16(anon_sym_RBRACE),
	3622: uint16(sym_identifier),
	3623: uint16(4),
	3624: uint16(3),
	3625: uint16(1),
	3626: uint16(sym_comment),
	3627: uint16(474),
	3628: uint16(1),
	3629: uint16(anon_sym_RBRACE),
	3630: uint16(476),
	3631: uint16(1),
	3632: uint16(sym_identifier),
	3633: uint16(178),
	3634: uint16(1),
	3635: uint16(aux_sym_block_lit_repeat2),
	3636: uint16(4),
	3637: uint16(3),
	3638: uint16(1),
	3639: uint16(sym_comment),
	3640: uint16(478),
	3641: uint16(1),
	3642: uint16(anon_sym_COMMA),
	3643: uint16(481),
	3644: uint16(1),
	3645: uint16(anon_sym_RBRACK),
	3646: uint16(151),
	3647: uint16(1),
	3648: uint16(aux_sym_enum_field_repeat1),
	3649: uint16(4),
	3650: uint16(3),
	3651: uint16(1),
	3652: uint16(sym_comment),
	3653: uint16(483),
	3654: uint16(1),
	3655: uint16(anon_sym_COMMA),
	3656: uint16(485),
	3657: uint16(1),
	3658: uint16(anon_sym_RBRACK),
	3659: uint16(161),
	3660: uint16(1),
	3661: uint16(aux_sym_field_options_repeat1),
	3662: uint16(4),
	3663: uint16(3),
	3664: uint16(1),
	3665: uint16(sym_comment),
	3666: uint16(487),
	3667: uint16(1),
	3668: uint16(anon_sym_COMMA),
	3669: uint16(489),
	3670: uint16(1),
	3671: uint16(anon_sym_RBRACK),
	3672: uint16(171),
	3673: uint16(1),
	3674: uint16(aux_sym_block_lit_repeat1),
	3675: uint16(4),
	3676: uint16(3),
	3677: uint16(1),
	3678: uint16(sym_comment),
	3679: uint16(487),
	3680: uint16(1),
	3681: uint16(anon_sym_COMMA),
	3682: uint16(489),
	3683: uint16(1),
	3684: uint16(anon_sym_RBRACK),
	3685: uint16(172),
	3686: uint16(1),
	3687: uint16(aux_sym_block_lit_repeat1),
	3688: uint16(4),
	3689: uint16(3),
	3690: uint16(1),
	3691: uint16(sym_comment),
	3692: uint16(491),
	3693: uint16(1),
	3694: uint16(anon_sym_COMMA),
	3695: uint16(493),
	3696: uint16(1),
	3697: uint16(anon_sym_RBRACK),
	3698: uint16(151),
	3699: uint16(1),
	3700: uint16(aux_sym_enum_field_repeat1),
	3701: uint16(4),
	3702: uint16(3),
	3703: uint16(1),
	3704: uint16(sym_comment),
	3705: uint16(278),
	3706: uint16(1),
	3707: uint16(anon_sym_DOT),
	3708: uint16(495),
	3709: uint16(1),
	3710: uint16(anon_sym_EQ),
	3711: uint16(160),
	3712: uint16(1),
	3713: uint16(aux_sym__option_name_repeat1),
	3714: uint16(4),
	3715: uint16(3),
	3716: uint16(1),
	3717: uint16(sym_comment),
	3718: uint16(497),
	3719: uint16(1),
	3720: uint16(anon_sym_SEMI),
	3721: uint16(499),
	3722: uint16(1),
	3723: uint16(anon_sym_COMMA),
	3724: uint16(157),
	3725: uint16(1),
	3726: uint16(aux_sym_ranges_repeat1),
	3727: uint16(4),
	3728: uint16(3),
	3729: uint16(1),
	3730: uint16(sym_comment),
	3731: uint16(502),
	3732: uint16(1),
	3733: uint16(anon_sym_SEMI),
	3734: uint16(504),
	3735: uint16(1),
	3736: uint16(anon_sym_COMMA),
	3737: uint16(158),
	3738: uint16(1),
	3739: uint16(aux_sym_field_names_repeat1),
	3740: uint16(4),
	3741: uint16(3),
	3742: uint16(1),
	3743: uint16(sym_comment),
	3744: uint16(491),
	3745: uint16(1),
	3746: uint16(anon_sym_COMMA),
	3747: uint16(507),
	3748: uint16(1),
	3749: uint16(anon_sym_RBRACK),
	3750: uint16(176),
	3751: uint16(1),
	3752: uint16(aux_sym_enum_field_repeat1),
	3753: uint16(4),
	3754: uint16(3),
	3755: uint16(1),
	3756: uint16(sym_comment),
	3757: uint16(278),
	3758: uint16(1),
	3759: uint16(anon_sym_DOT),
	3760: uint16(509),
	3761: uint16(1),
	3762: uint16(anon_sym_EQ),
	3763: uint16(56),
	3764: uint16(1),
	3765: uint16(aux_sym__option_name_repeat1),
	3766: uint16(4),
	3767: uint16(3),
	3768: uint16(1),
	3769: uint16(sym_comment),
	3770: uint16(483),
	3771: uint16(1),
	3772: uint16(anon_sym_COMMA),
	3773: uint16(511),
	3774: uint16(1),
	3775: uint16(anon_sym_RBRACK),
	3776: uint16(162),
	3777: uint16(1),
	3778: uint16(aux_sym_field_options_repeat1),
	3779: uint16(4),
	3780: uint16(3),
	3781: uint16(1),
	3782: uint16(sym_comment),
	3783: uint16(513),
	3784: uint16(1),
	3785: uint16(anon_sym_COMMA),
	3786: uint16(516),
	3787: uint16(1),
	3788: uint16(anon_sym_RBRACK),
	3789: uint16(162),
	3790: uint16(1),
	3791: uint16(aux_sym_field_options_repeat1),
	3792: uint16(4),
	3793: uint16(3),
	3794: uint16(1),
	3795: uint16(sym_comment),
	3796: uint16(356),
	3797: uint16(1),
	3798: uint16(anon_sym_LPAREN),
	3799: uint16(358),
	3800: uint16(1),
	3801: uint16(sym_identifier),
	3802: uint16(258),
	3803: uint16(1),
	3804: uint16(sym__option_name),
	3805: uint16(4),
	3806: uint16(3),
	3807: uint16(1),
	3808: uint16(sym_comment),
	3809: uint16(278),
	3810: uint16(1),
	3811: uint16(anon_sym_DOT),
	3812: uint16(518),
	3813: uint16(1),
	3814: uint16(anon_sym_EQ),
	3815: uint16(56),
	3816: uint16(1),
	3817: uint16(aux_sym__option_name_repeat1),
	3818: uint16(3),
	3819: uint16(3),
	3820: uint16(1),
	3821: uint16(sym_comment),
	3822: uint16(522),
	3823: uint16(1),
	3824: uint16(anon_sym_to),
	3825: uint16(520),
	3826: uint16(2),
	3827: uint16(anon_sym_SEMI),
	3828: uint16(anon_sym_COMMA),
	3829: uint16(4),
	3830: uint16(3),
	3831: uint16(1),
	3832: uint16(sym_comment),
	3833: uint16(524),
	3834: uint16(1),
	3835: uint16(anon_sym_SEMI),
	3836: uint16(526),
	3837: uint16(1),
	3838: uint16(anon_sym_COMMA),
	3839: uint16(170),
	3840: uint16(1),
	3841: uint16(aux_sym_field_names_repeat1),
	3842: uint16(4),
	3843: uint16(3),
	3844: uint16(1),
	3845: uint16(sym_comment),
	3846: uint16(528),
	3847: uint16(1),
	3848: uint16(anon_sym_SEMI),
	3849: uint16(530),
	3850: uint16(1),
	3851: uint16(anon_sym_COMMA),
	3852: uint16(169),
	3853: uint16(1),
	3854: uint16(aux_sym_ranges_repeat1),
	3855: uint16(4),
	3856: uint16(3),
	3857: uint16(1),
	3858: uint16(sym_comment),
	3859: uint16(487),
	3860: uint16(1),
	3861: uint16(anon_sym_COMMA),
	3862: uint16(532),
	3863: uint16(1),
	3864: uint16(anon_sym_RBRACK),
	3865: uint16(153),
	3866: uint16(1),
	3867: uint16(aux_sym_block_lit_repeat1),
	3868: uint16(4),
	3869: uint16(3),
	3870: uint16(1),
	3871: uint16(sym_comment),
	3872: uint16(530),
	3873: uint16(1),
	3874: uint16(anon_sym_COMMA),
	3875: uint16(534),
	3876: uint16(1),
	3877: uint16(anon_sym_SEMI),
	3878: uint16(157),
	3879: uint16(1),
	3880: uint16(aux_sym_ranges_repeat1),
	3881: uint16(4),
	3882: uint16(3),
	3883: uint16(1),
	3884: uint16(sym_comment),
	3885: uint16(526),
	3886: uint16(1),
	3887: uint16(anon_sym_COMMA),
	3888: uint16(536),
	3889: uint16(1),
	3890: uint16(anon_sym_SEMI),
	3891: uint16(158),
	3892: uint16(1),
	3893: uint16(aux_sym_field_names_repeat1),
	3894: uint16(4),
	3895: uint16(3),
	3896: uint16(1),
	3897: uint16(sym_comment),
	3898: uint16(538),
	3899: uint16(1),
	3900: uint16(anon_sym_COMMA),
	3901: uint16(541),
	3902: uint16(1),
	3903: uint16(anon_sym_RBRACK),
	3904: uint16(171),
	3905: uint16(1),
	3906: uint16(aux_sym_block_lit_repeat1),
	3907: uint16(4),
	3908: uint16(3),
	3909: uint16(1),
	3910: uint16(sym_comment),
	3911: uint16(487),
	3912: uint16(1),
	3913: uint16(anon_sym_COMMA),
	3914: uint16(543),
	3915: uint16(1),
	3916: uint16(anon_sym_RBRACK),
	3917: uint16(171),
	3918: uint16(1),
	3919: uint16(aux_sym_block_lit_repeat1),
	3920: uint16(4),
	3921: uint16(3),
	3922: uint16(1),
	3923: uint16(sym_comment),
	3924: uint16(491),
	3925: uint16(1),
	3926: uint16(anon_sym_COMMA),
	3927: uint16(545),
	3928: uint16(1),
	3929: uint16(anon_sym_RBRACK),
	3930: uint16(155),
	3931: uint16(1),
	3932: uint16(aux_sym_enum_field_repeat1),
	3933: uint16(4),
	3934: uint16(3),
	3935: uint16(1),
	3936: uint16(sym_comment),
	3937: uint16(278),
	3938: uint16(1),
	3939: uint16(anon_sym_DOT),
	3940: uint16(547),
	3941: uint16(1),
	3942: uint16(anon_sym_EQ),
	3943: uint16(164),
	3944: uint16(1),
	3945: uint16(aux_sym__option_name_repeat1),
	3946: uint16(4),
	3947: uint16(3),
	3948: uint16(1),
	3949: uint16(sym_comment),
	3950: uint16(356),
	3951: uint16(1),
	3952: uint16(anon_sym_LPAREN),
	3953: uint16(358),
	3954: uint16(1),
	3955: uint16(sym_identifier),
	3956: uint16(282),
	3957: uint16(1),
	3958: uint16(sym__option_name),
	3959: uint16(4),
	3960: uint16(3),
	3961: uint16(1),
	3962: uint16(sym_comment),
	3963: uint16(491),
	3964: uint16(1),
	3965: uint16(anon_sym_COMMA),
	3966: uint16(545),
	3967: uint16(1),
	3968: uint16(anon_sym_RBRACK),
	3969: uint16(151),
	3970: uint16(1),
	3971: uint16(aux_sym_enum_field_repeat1),
	3972: uint16(4),
	3973: uint16(3),
	3974: uint16(1),
	3975: uint16(sym_comment),
	3976: uint16(356),
	3977: uint16(1),
	3978: uint16(anon_sym_LPAREN),
	3979: uint16(358),
	3980: uint16(1),
	3981: uint16(sym_identifier),
	3982: uint16(283),
	3983: uint16(1),
	3984: uint16(sym__option_name),
	3985: uint16(4),
	3986: uint16(3),
	3987: uint16(1),
	3988: uint16(sym_comment),
	3989: uint16(476),
	3990: uint16(1),
	3991: uint16(sym_identifier),
	3992: uint16(549),
	3993: uint16(1),
	3994: uint16(anon_sym_RBRACE),
	3995: uint16(181),
	3996: uint16(1),
	3997: uint16(aux_sym_block_lit_repeat2),
	3998: uint16(4),
	3999: uint16(3),
	4000: uint16(1),
	4001: uint16(sym_comment),
	4002: uint16(212),
	4003: uint16(1),
	4004: uint16(anon_sym_DQUOTE),
	4005: uint16(214),
	4006: uint16(1),
	4007: uint16(anon_sym_SQUOTE),
	4008: uint16(247),
	4009: uint16(1),
	4010: uint16(sym_string),
	4011: uint16(4),
	4012: uint16(3),
	4013: uint16(1),
	4014: uint16(sym_comment),
	4015: uint16(356),
	4016: uint16(1),
	4017: uint16(anon_sym_LPAREN),
	4018: uint16(358),
	4019: uint16(1),
	4020: uint16(sym_identifier),
	4021: uint16(284),
	4022: uint16(1),
	4023: uint16(sym__option_name),
	4024: uint16(4),
	4025: uint16(3),
	4026: uint16(1),
	4027: uint16(sym_comment),
	4028: uint16(456),
	4029: uint16(1),
	4030: uint16(anon_sym_RBRACE),
	4031: uint16(551),
	4032: uint16(1),
	4033: uint16(sym_identifier),
	4034: uint16(181),
	4035: uint16(1),
	4036: uint16(aux_sym_block_lit_repeat2),
	4037: uint16(3),
	4038: uint16(3),
	4039: uint16(1),
	4040: uint16(sym_comment),
	4041: uint16(554),
	4042: uint16(1),
	4043: uint16(anon_sym_LBRACE),
	4044: uint16(8),
	4045: uint16(1),
	4046: uint16(sym_message_body),
	4047: uint16(3),
	4048: uint16(3),
	4049: uint16(1),
	4050: uint16(sym_comment),
	4051: uint16(556),
	4052: uint16(1),
	4053: uint16(sym_identifier),
	4054: uint16(207),
	4055: uint16(1),
	4056: uint16(aux_sym_message_or_enum_type_repeat1),
	4057: uint16(3),
	4058: uint16(3),
	4059: uint16(1),
	4060: uint16(sym_comment),
	4061: uint16(558),
	4062: uint16(1),
	4063: uint16(anon_sym_LBRACE),
	4064: uint16(63),
	4065: uint16(1),
	4066: uint16(sym_message_body),
	4067: uint16(3),
	4068: uint16(3),
	4069: uint16(1),
	4070: uint16(sym_comment),
	4071: uint16(560),
	4072: uint16(1),
	4073: uint16(anon_sym_SEMI),
	4074: uint16(562),
	4075: uint16(1),
	4076: uint16(anon_sym_LBRACK),
	4077: uint16(2),
	4078: uint16(3),
	4079: uint16(1),
	4080: uint16(sym_comment),
	4081: uint16(541),
	4082: uint16(2),
	4083: uint16(anon_sym_COMMA),
	4084: uint16(anon_sym_RBRACK),
	4085: uint16(2),
	4086: uint16(3),
	4087: uint16(1),
	4088: uint16(sym_comment),
	4089: uint16(422),
	4090: uint16(2),
	4091: uint16(anon_sym_RBRACE),
	4092: uint16(sym_identifier),
	4093: uint16(3),
	4094: uint16(3),
	4095: uint16(1),
	4096: uint16(sym_comment),
	4097: uint16(564),
	4098: uint16(1),
	4099: uint16(anon_sym_SEMI),
	4100: uint16(566),
	4101: uint16(1),
	4102: uint16(anon_sym_LBRACK),
	4103: uint16(3),
	4104: uint16(3),
	4105: uint16(1),
	4106: uint16(sym_comment),
	4107: uint16(568),
	4108: uint16(1),
	4109: uint16(anon_sym_LBRACE),
	4110: uint16(71),
	4111: uint16(1),
	4112: uint16(sym_enum_body),
	4113: uint16(2),
	4114: uint16(3),
	4115: uint16(1),
	4116: uint16(sym_comment),
	4117: uint16(450),
	4118: uint16(2),
	4119: uint16(anon_sym_RBRACE),
	4120: uint16(sym_identifier),
	4121: uint16(3),
	4122: uint16(3),
	4123: uint16(1),
	4124: uint16(sym_comment),
	4125: uint16(570),
	4126: uint16(1),
	4127: uint16(sym_identifier),
	4128: uint16(238),
	4129: uint16(1),
	4130: uint16(sym_full_ident),
	4131: uint16(3),
	4132: uint16(3),
	4133: uint16(1),
	4134: uint16(sym_comment),
	4135: uint16(572),
	4136: uint16(1),
	4137: uint16(anon_sym_SEMI),
	4138: uint16(574),
	4139: uint16(1),
	4140: uint16(anon_sym_LBRACK),
	4141: uint16(2),
	4142: uint16(3),
	4143: uint16(1),
	4144: uint16(sym_comment),
	4145: uint16(472),
	4146: uint16(2),
	4147: uint16(anon_sym_RBRACE),
	4148: uint16(sym_identifier),
	4149: uint16(2),
	4150: uint16(3),
	4151: uint16(1),
	4152: uint16(sym_comment),
	4153: uint16(576),
	4154: uint16(2),
	4155: uint16(anon_sym_COMMA),
	4156: uint16(anon_sym_RBRACK),
	4157: uint16(2),
	4158: uint16(3),
	4159: uint16(1),
	4160: uint16(sym_comment),
	4161: uint16(481),
	4162: uint16(2),
	4163: uint16(anon_sym_COMMA),
	4164: uint16(anon_sym_RBRACK),
	4165: uint16(3),
	4166: uint16(3),
	4167: uint16(1),
	4168: uint16(sym_comment),
	4169: uint16(578),
	4170: uint16(1),
	4171: uint16(anon_sym_LBRACE),
	4172: uint16(25),
	4173: uint16(1),
	4174: uint16(sym_enum_body),
	4175: uint16(3),
	4176: uint16(3),
	4177: uint16(1),
	4178: uint16(sym_comment),
	4179: uint16(580),
	4180: uint16(1),
	4181: uint16(anon_sym_SEMI),
	4182: uint16(582),
	4183: uint16(1),
	4184: uint16(anon_sym_LBRACK),
	4185: uint16(3),
	4186: uint16(3),
	4187: uint16(1),
	4188: uint16(sym_comment),
	4189: uint16(584),
	4190: uint16(1),
	4191: uint16(sym_identifier),
	4192: uint16(274),
	4193: uint16(1),
	4194: uint16(sym_rpc_name),
	4195: uint16(3),
	4196: uint16(3),
	4197: uint16(1),
	4198: uint16(sym_comment),
	4199: uint16(586),
	4200: uint16(1),
	4201: uint16(sym_identifier),
	4202: uint16(207),
	4203: uint16(1),
	4204: uint16(aux_sym_message_or_enum_type_repeat1),
	4205: uint16(3),
	4206: uint16(3),
	4207: uint16(1),
	4208: uint16(sym_comment),
	4209: uint16(352),
	4210: uint16(1),
	4211: uint16(anon_sym_SEMI),
	4212: uint16(588),
	4213: uint16(1),
	4214: uint16(anon_sym_LBRACE),
	4215: uint16(2),
	4216: uint16(3),
	4217: uint16(1),
	4218: uint16(sym_comment),
	4219: uint16(590),
	4220: uint16(2),
	4221: uint16(anon_sym_SEMI),
	4222: uint16(anon_sym_COMMA),
	4223: uint16(3),
	4224: uint16(3),
	4225: uint16(1),
	4226: uint16(sym_comment),
	4227: uint16(338),
	4228: uint16(1),
	4229: uint16(anon_sym_SEMI),
	4230: uint16(592),
	4231: uint16(1),
	4232: uint16(anon_sym_LBRACE),
	4233: uint16(2),
	4234: uint16(3),
	4235: uint16(1),
	4236: uint16(sym_comment),
	4237: uint16(594),
	4238: uint16(2),
	4239: uint16(anon_sym_GT),
	4240: uint16(sym_identifier),
	4241: uint16(2),
	4242: uint16(3),
	4243: uint16(1),
	4244: uint16(sym_comment),
	4245: uint16(497),
	4246: uint16(2),
	4247: uint16(anon_sym_SEMI),
	4248: uint16(anon_sym_COMMA),
	4249: uint16(2),
	4250: uint16(3),
	4251: uint16(1),
	4252: uint16(sym_comment),
	4253: uint16(502),
	4254: uint16(2),
	4255: uint16(anon_sym_SEMI),
	4256: uint16(anon_sym_COMMA),
	4257: uint16(2),
	4258: uint16(3),
	4259: uint16(1),
	4260: uint16(sym_comment),
	4261: uint16(596),
	4262: uint16(2),
	4263: uint16(anon_sym_RBRACE),
	4264: uint16(sym_identifier),
	4265: uint16(3),
	4266: uint16(3),
	4267: uint16(1),
	4268: uint16(sym_comment),
	4269: uint16(598),
	4270: uint16(1),
	4271: uint16(sym_identifier),
	4272: uint16(207),
	4273: uint16(1),
	4274: uint16(aux_sym_message_or_enum_type_repeat1),
	4275: uint16(3),
	4276: uint16(3),
	4277: uint16(1),
	4278: uint16(sym_comment),
	4279: uint16(586),
	4280: uint16(1),
	4281: uint16(sym_identifier),
	4282: uint16(183),
	4283: uint16(1),
	4284: uint16(aux_sym_message_or_enum_type_repeat1),
	4285: uint16(2),
	4286: uint16(3),
	4287: uint16(1),
	4288: uint16(sym_comment),
	4289: uint16(601),
	4290: uint16(2),
	4291: uint16(anon_sym_COMMA),
	4292: uint16(anon_sym_RBRACK),
	4293: uint16(2),
	4294: uint16(3),
	4295: uint16(1),
	4296: uint16(sym_comment),
	4297: uint16(516),
	4298: uint16(2),
	4299: uint16(anon_sym_COMMA),
	4300: uint16(anon_sym_RBRACK),
	4301: uint16(3),
	4302: uint16(3),
	4303: uint16(1),
	4304: uint16(sym_comment),
	4305: uint16(603),
	4306: uint16(1),
	4307: uint16(sym_identifier),
	4308: uint16(248),
	4309: uint16(1),
	4310: uint16(sym_service_name),
	4311: uint16(3),
	4312: uint16(3),
	4313: uint16(1),
	4314: uint16(sym_comment),
	4315: uint16(605),
	4316: uint16(1),
	4317: uint16(sym_identifier),
	4318: uint16(184),
	4319: uint16(1),
	4320: uint16(sym_message_name),
	4321: uint16(3),
	4322: uint16(3),
	4323: uint16(1),
	4324: uint16(sym_comment),
	4325: uint16(607),
	4326: uint16(1),
	4327: uint16(sym_identifier),
	4328: uint16(189),
	4329: uint16(1),
	4330: uint16(sym_enum_name),
	4331: uint16(3),
	4332: uint16(3),
	4333: uint16(1),
	4334: uint16(sym_comment),
	4335: uint16(609),
	4336: uint16(1),
	4337: uint16(anon_sym_SEMI),
	4338: uint16(611),
	4339: uint16(1),
	4340: uint16(anon_sym_LBRACK),
	4341: uint16(3),
	4342: uint16(3),
	4343: uint16(1),
	4344: uint16(sym_comment),
	4345: uint16(613),
	4346: uint16(1),
	4347: uint16(anon_sym_SEMI),
	4348: uint16(615),
	4349: uint16(1),
	4350: uint16(anon_sym_LBRACK),
	4351: uint16(2),
	4352: uint16(3),
	4353: uint16(1),
	4354: uint16(sym_comment),
	4355: uint16(428),
	4356: uint16(2),
	4357: uint16(anon_sym_RBRACE),
	4358: uint16(sym_identifier),
	4359: uint16(3),
	4360: uint16(3),
	4361: uint16(1),
	4362: uint16(sym_comment),
	4363: uint16(617),
	4364: uint16(1),
	4365: uint16(anon_sym_SEMI),
	4366: uint16(619),
	4367: uint16(1),
	4368: uint16(anon_sym_LBRACE),
	4369: uint16(3),
	4370: uint16(3),
	4371: uint16(1),
	4372: uint16(sym_comment),
	4373: uint16(607),
	4374: uint16(1),
	4375: uint16(sym_identifier),
	4376: uint16(196),
	4377: uint16(1),
	4378: uint16(sym_enum_name),
	4379: uint16(3),
	4380: uint16(3),
	4381: uint16(1),
	4382: uint16(sym_comment),
	4383: uint16(605),
	4384: uint16(1),
	4385: uint16(sym_identifier),
	4386: uint16(182),
	4387: uint16(1),
	4388: uint16(sym_message_name),
	4389: uint16(3),
	4390: uint16(3),
	4391: uint16(1),
	4392: uint16(sym_comment),
	4393: uint16(570),
	4394: uint16(1),
	4395: uint16(sym_identifier),
	4396: uint16(262),
	4397: uint16(1),
	4398: uint16(sym_full_ident),
	4399: uint16(2),
	4400: uint16(3),
	4401: uint16(1),
	4402: uint16(sym_comment),
	4403: uint16(621),
	4404: uint16(1),
	4405: uint16(anon_sym_EQ),
	4406: uint16(2),
	4407: uint16(3),
	4408: uint16(1),
	4409: uint16(sym_comment),
	4410: uint16(623),
	4411: uint16(1),
	4412: uint16(anon_sym_RPAREN),
	4413: uint16(2),
	4414: uint16(3),
	4415: uint16(1),
	4416: uint16(sym_comment),
	4417: uint16(625),
	4418: uint16(1),
	4419: uint16(anon_sym_SEMI),
	4420: uint16(2),
	4421: uint16(3),
	4422: uint16(1),
	4423: uint16(sym_comment),
	4424: uint16(627),
	4425: uint16(1),
	4426: uint16(anon_sym_SEMI),
	4427: uint16(2),
	4428: uint16(3),
	4429: uint16(1),
	4430: uint16(sym_comment),
	4431: uint16(629),
	4432: uint16(1),
	4433: uint16(sym_identifier),
	4434: uint16(2),
	4435: uint16(3),
	4436: uint16(1),
	4437: uint16(sym_comment),
	4438: uint16(631),
	4439: uint16(1),
	4440: uint16(anon_sym_RBRACK),
	4441: uint16(2),
	4442: uint16(3),
	4443: uint16(1),
	4444: uint16(sym_comment),
	4445: uint16(633),
	4446: uint16(1),
	4447: uint16(anon_sym_LT),
	4448: uint16(2),
	4449: uint16(3),
	4450: uint16(1),
	4451: uint16(sym_comment),
	4452: uint16(635),
	4453: uint16(1),
	4454: uint16(anon_sym_SEMI),
	4455: uint16(2),
	4456: uint16(3),
	4457: uint16(1),
	4458: uint16(sym_comment),
	4459: uint16(637),
	4460: uint16(1),
	4461: uint16(anon_sym_RPAREN),
	4462: uint16(2),
	4463: uint16(3),
	4464: uint16(1),
	4465: uint16(sym_comment),
	4466: uint16(639),
	4467: uint16(1),
	4468: uint16(anon_sym_EQ),
	4469: uint16(2),
	4470: uint16(3),
	4471: uint16(1),
	4472: uint16(sym_comment),
	4473: uint16(641),
	4474: uint16(1),
	4475: uint16(anon_sym_LPAREN),
	4476: uint16(2),
	4477: uint16(3),
	4478: uint16(1),
	4479: uint16(sym_comment),
	4480: uint16(643),
	4481: uint16(1),
	4482: uint16(anon_sym_SEMI),
	4483: uint16(2),
	4484: uint16(3),
	4485: uint16(1),
	4486: uint16(sym_comment),
	4487: uint16(645),
	4488: uint16(1),
	4489: uint16(anon_sym_SEMI),
	4490: uint16(2),
	4491: uint16(3),
	4492: uint16(1),
	4493: uint16(sym_comment),
	4494: uint16(647),
	4495: uint16(1),
	4496: uint16(anon_sym_RBRACK),
	4497: uint16(2),
	4498: uint16(3),
	4499: uint16(1),
	4500: uint16(sym_comment),
	4501: uint16(649),
	4502: uint16(1),
	4503: uint16(anon_sym_RPAREN),
	4504: uint16(2),
	4505: uint16(3),
	4506: uint16(1),
	4507: uint16(sym_comment),
	4508: uint16(651),
	4509: uint16(1),
	4510: uint16(sym_identifier),
	4511: uint16(2),
	4512: uint16(3),
	4513: uint16(1),
	4514: uint16(sym_comment),
	4515: uint16(653),
	4516: uint16(1),
	4517: uint16(anon_sym_EQ),
	4518: uint16(2),
	4519: uint16(3),
	4520: uint16(1),
	4521: uint16(sym_comment),
	4522: uint16(655),
	4523: uint16(1),
	4524: uint16(anon_sym_RPAREN),
	4525: uint16(2),
	4526: uint16(3),
	4527: uint16(1),
	4528: uint16(sym_comment),
	4529: uint16(657),
	4530: uint16(1),
	4531: uint16(anon_sym_RBRACK),
	4532: uint16(2),
	4533: uint16(3),
	4534: uint16(1),
	4535: uint16(sym_comment),
	4536: uint16(659),
	4537: uint16(1),
	4538: uint16(sym_identifier),
	4539: uint16(2),
	4540: uint16(3),
	4541: uint16(1),
	4542: uint16(sym_comment),
	4543: uint16(661),
	4544: uint16(1),
	4545: uint16(anon_sym_RBRACK),
	4546: uint16(2),
	4547: uint16(3),
	4548: uint16(1),
	4549: uint16(sym_comment),
	4550: uint16(663),
	4551: uint16(1),
	4552: uint16(sym_identifier),
	4553: uint16(2),
	4554: uint16(3),
	4555: uint16(1),
	4556: uint16(sym_comment),
	4557: uint16(665),
	4558: uint16(1),
	4559: uint16(anon_sym_SEMI),
	4560: uint16(2),
	4561: uint16(3),
	4562: uint16(1),
	4563: uint16(sym_comment),
	4564: uint16(667),
	4565: uint16(1),
	4566: uint16(anon_sym_SEMI),
	4567: uint16(2),
	4568: uint16(3),
	4569: uint16(1),
	4570: uint16(sym_comment),
	4571: uint16(669),
	4572: uint16(1),
	4573: uint16(anon_sym_LPAREN),
	4574: uint16(2),
	4575: uint16(3),
	4576: uint16(1),
	4577: uint16(sym_comment),
	4578: uint16(671),
	4579: uint16(1),
	4580: uint16(anon_sym_returns),
	4581: uint16(2),
	4582: uint16(3),
	4583: uint16(1),
	4584: uint16(sym_comment),
	4585: uint16(673),
	4586: uint16(1),
	4587: uint16(anon_sym_SEMI),
	4588: uint16(2),
	4589: uint16(3),
	4590: uint16(1),
	4591: uint16(sym_comment),
	4592: uint16(675),
	4593: uint16(1),
	4594: uint16(anon_sym_LBRACE),
	4595: uint16(2),
	4596: uint16(3),
	4597: uint16(1),
	4598: uint16(sym_comment),
	4599: uint16(677),
	4600: uint16(1),
	4601: uint16(anon_sym_LBRACE),
	4602: uint16(2),
	4603: uint16(3),
	4604: uint16(1),
	4605: uint16(sym_comment),
	4606: uint16(679),
	4607: uint16(1),
	4608: uint16(anon_sym_RBRACK),
	4609: uint16(2),
	4610: uint16(3),
	4611: uint16(1),
	4612: uint16(sym_comment),
	4613: uint16(681),
	4614: uint16(1),
	4615: uint16(anon_sym_EQ),
	4616: uint16(2),
	4617: uint16(3),
	4618: uint16(1),
	4619: uint16(sym_comment),
	4620: uint16(683),
	4621: uint16(1),
	4622: uint16(sym_identifier),
	4623: uint16(2),
	4624: uint16(3),
	4625: uint16(1),
	4626: uint16(sym_comment),
	4627: uint16(685),
	4628: uint16(1),
	4629: uint16(anon_sym_SEMI),
	4630: uint16(2),
	4631: uint16(3),
	4632: uint16(1),
	4633: uint16(sym_comment),
	4634: uint16(687),
	4635: uint16(1),
	4636: uint16(anon_sym_LBRACE),
	4637: uint16(2),
	4638: uint16(3),
	4639: uint16(1),
	4640: uint16(sym_comment),
	4641: uint16(689),
	4642: uint16(1),
	4643: uint16(anon_sym_returns),
	4644: uint16(2),
	4645: uint16(3),
	4646: uint16(1),
	4647: uint16(sym_comment),
	4648: uint16(691),
	4649: uint16(1),
	4650: uint16(anon_sym_RPAREN),
	4651: uint16(2),
	4652: uint16(3),
	4653: uint16(1),
	4654: uint16(sym_comment),
	4655: uint16(693),
	4656: uint16(1),
	4657: uint16(anon_sym_LBRACE),
	4658: uint16(2),
	4659: uint16(3),
	4660: uint16(1),
	4661: uint16(sym_comment),
	4662: uint16(695),
	4663: uint16(1),
	4664: uint16(anon_sym_EQ),
	4665: uint16(2),
	4666: uint16(3),
	4667: uint16(1),
	4668: uint16(sym_comment),
	4669: uint16(697),
	4670: uint16(1),
	4671: uint16(sym_identifier),
	4672: uint16(2),
	4673: uint16(3),
	4674: uint16(1),
	4675: uint16(sym_comment),
	4676: uint16(699),
	4677: uint16(1),
	4678: uint16(anon_sym_LBRACE),
	4679: uint16(2),
	4680: uint16(3),
	4681: uint16(1),
	4682: uint16(sym_comment),
	4683: uint16(701),
	4684: uint16(1),
	4685: uint16(anon_sym_GT),
	4686: uint16(2),
	4687: uint16(3),
	4688: uint16(1),
	4689: uint16(sym_comment),
	4690: uint16(703),
	4691: uint16(1),
	4692: uint16(anon_sym_SEMI),
	4693: uint16(2),
	4694: uint16(3),
	4695: uint16(1),
	4696: uint16(sym_comment),
	4697: uint16(705),
	4698: uint16(1),
	4699: uint16(anon_sym_SEMI),
	4700: uint16(2),
	4701: uint16(3),
	4702: uint16(1),
	4703: uint16(sym_comment),
	4704: uint16(707),
	4705: uint16(1),
	4706: uint16(anon_sym_SEMI),
	4707: uint16(2),
	4708: uint16(3),
	4709: uint16(1),
	4710: uint16(sym_comment),
	4711: uint16(709),
	4712: uint16(1),
	4713: uint16(anon_sym_EQ),
	4714: uint16(2),
	4715: uint16(3),
	4716: uint16(1),
	4717: uint16(sym_comment),
	4718: uint16(711),
	4719: uint16(1),
	4720: uint16(anon_sym_EQ),
	4721: uint16(2),
	4722: uint16(3),
	4723: uint16(1),
	4724: uint16(sym_comment),
	4725: uint16(713),
	4726: uint16(1),
	4727: uint16(anon_sym_RPAREN),
	4728: uint16(2),
	4729: uint16(3),
	4730: uint16(1),
	4731: uint16(sym_comment),
	4732: uint16(715),
	4733: uint16(1),
	4734: uint16(sym_identifier),
	4735: uint16(2),
	4736: uint16(3),
	4737: uint16(1),
	4738: uint16(sym_comment),
	4739: uint16(717),
	4740: uint16(1),
	4741: uint16(anon_sym_SEMI),
	4742: uint16(2),
	4743: uint16(3),
	4744: uint16(1),
	4745: uint16(sym_comment),
	4746: uint16(719),
	4747: uint16(1),
	4748: uint16(anon_sym_EQ),
	4749: uint16(2),
	4750: uint16(3),
	4751: uint16(1),
	4752: uint16(sym_comment),
	4753: uint16(721),
	4754: uint16(1),
	4755: uint16(sym_identifier),
	4756: uint16(2),
	4757: uint16(3),
	4758: uint16(1),
	4759: uint16(sym_comment),
	4760: uint16(723),
	4761: uint16(1),
	4762: uint16(anon_sym_SEMI),
	4763: uint16(2),
	4764: uint16(3),
	4765: uint16(1),
	4766: uint16(sym_comment),
	4767: uint16(725),
	4768: uint16(1),
	4769: uint16(anon_sym_LPAREN),
	4770: uint16(2),
	4771: uint16(3),
	4772: uint16(1),
	4773: uint16(sym_comment),
	4774: uint16(727),
	4775: uint16(1),
	4776: uint16(anon_sym_LPAREN),
	4777: uint16(2),
	4778: uint16(3),
	4779: uint16(1),
	4780: uint16(sym_comment),
	4781: uint16(729),
	4782: uint16(1),
	4783: uint16(anon_sym_SEMI),
	4784: uint16(2),
	4785: uint16(3),
	4786: uint16(1),
	4787: uint16(sym_comment),
	4788: uint16(731),
	4789: uint16(1),
	4790: uint16(anon_sym_SEMI),
	4791: uint16(2),
	4792: uint16(3),
	4793: uint16(1),
	4794: uint16(sym_comment),
	4795: uint16(733),
	4796: uint16(1),
	4797: uint16(anon_sym_EQ),
	4798: uint16(2),
	4799: uint16(3),
	4800: uint16(1),
	4801: uint16(sym_comment),
	4802: uint16(392),
	4803: uint16(1),
	4804: uint16(anon_sym_DOT),
	4805: uint16(2),
	4806: uint16(3),
	4807: uint16(1),
	4808: uint16(sym_comment),
	4809: uint16(735),
	4810: uint16(1),
	4811: uint16(sym_identifier),
	4812: uint16(2),
	4813: uint16(3),
	4814: uint16(1),
	4815: uint16(sym_comment),
	4816: uint16(737),
	4817: uint16(1),
	4818: uint16(anon_sym_DQUOTEproto3_DQUOTE),
	4819: uint16(2),
	4820: uint16(3),
	4821: uint16(1),
	4822: uint16(sym_comment),
	4823: uint16(739),
	4824: uint16(1),
	4826: uint16(2),
	4827: uint16(3),
	4828: uint16(1),
	4829: uint16(sym_comment),
	4830: uint16(741),
	4831: uint16(1),
	4832: uint16(anon_sym_EQ),
	4833: uint16(2),
	4834: uint16(3),
	4835: uint16(1),
	4836: uint16(sym_comment),
	4837: uint16(743),
	4838: uint16(1),
	4839: uint16(anon_sym_EQ),
	4840: uint16(2),
	4841: uint16(3),
	4842: uint16(1),
	4843: uint16(sym_comment),
	4844: uint16(745),
	4845: uint16(1),
	4846: uint16(anon_sym_EQ),
	4847: uint16(2),
	4848: uint16(3),
	4849: uint16(1),
	4850: uint16(sym_comment),
	4851: uint16(747),
	4852: uint16(1),
	4853: uint16(anon_sym_COMMA),
	4854: uint16(2),
	4855: uint16(3),
	4856: uint16(1),
	4857: uint16(sym_comment),
	4858: uint16(749),
	4859: uint16(1),
	4860: uint16(anon_sym_COMMA),
	4861: uint16(2),
	4862: uint16(3),
	4863: uint16(1),
	4864: uint16(sym_comment),
	4865: uint16(751),
	4866: uint16(1),
	4867: uint16(anon_sym_EQ),
}

var ts_small_parse_table_map = [286]uint32_t{
	1:   uint32(77),
	2:   uint32(154),
	3:   uint32(231),
	4:   uint32(308),
	5:   uint32(385),
	6:   uint32(420),
	7:   uint32(455),
	8:   uint32(506),
	9:   uint32(541),
	10:  uint32(576),
	11:  uint32(611),
	12:  uint32(646),
	13:  uint32(681),
	14:  uint32(716),
	15:  uint32(751),
	16:  uint32(786),
	17:  uint32(821),
	18:  uint32(856),
	19:  uint32(891),
	20:  uint32(926),
	21:  uint32(961),
	22:  uint32(996),
	23:  uint32(1047),
	24:  uint32(1082),
	25:  uint32(1117),
	26:  uint32(1152),
	27:  uint32(1203),
	28:  uint32(1234),
	29:  uint32(1273),
	30:  uint32(1302),
	31:  uint32(1331),
	32:  uint32(1381),
	33:  uint32(1417),
	34:  uint32(1445),
	35:  uint32(1473),
	36:  uint32(1509),
	37:  uint32(1537),
	38:  uint32(1573),
	39:  uint32(1620),
	40:  uint32(1664),
	41:  uint32(1708),
	42:  uint32(1752),
	43:  uint32(1796),
	44:  uint32(1840),
	45:  uint32(1884),
	46:  uint32(1928),
	47:  uint32(1972),
	48:  uint32(2016),
	49:  uint32(2054),
	50:  uint32(2092),
	51:  uint32(2130),
	52:  uint32(2151),
	53:  uint32(2167),
	54:  uint32(2183),
	55:  uint32(2202),
	56:  uint32(2220),
	57:  uint32(2242),
	58:  uint32(2256),
	59:  uint32(2270),
	60:  uint32(2292),
	61:  uint32(2306),
	62:  uint32(2320),
	63:  uint32(2342),
	64:  uint32(2356),
	65:  uint32(2370),
	66:  uint32(2394),
	67:  uint32(2416),
	68:  uint32(2430),
	69:  uint32(2452),
	70:  uint32(2466),
	71:  uint32(2480),
	72:  uint32(2494),
	73:  uint32(2512),
	74:  uint32(2526),
	75:  uint32(2540),
	76:  uint32(2562),
	77:  uint32(2584),
	78:  uint32(2598),
	79:  uint32(2620),
	80:  uint32(2638),
	81:  uint32(2656),
	82:  uint32(2674),
	83:  uint32(2692),
	84:  uint32(2710),
	85:  uint32(2728),
	86:  uint32(2746),
	87:  uint32(2763),
	88:  uint32(2782),
	89:  uint32(2801),
	90:  uint32(2812),
	91:  uint32(2831),
	92:  uint32(2848),
	93:  uint32(2859),
	94:  uint32(2876),
	95:  uint32(2887),
	96:  uint32(2904),
	97:  uint32(2923),
	98:  uint32(2940),
	99:  uint32(2959),
	100: uint32(2976),
	101: uint32(2987),
	102: uint32(3004),
	103: uint32(3015),
	104: uint32(3032),
	105: uint32(3051),
	106: uint32(3062),
	107: uint32(3081),
	108: uint32(3092),
	109: uint32(3103),
	110: uint32(3120),
	111: uint32(3137),
	112: uint32(3156),
	113: uint32(3173),
	114: uint32(3185),
	115: uint32(3201),
	116: uint32(3215),
	117: uint32(3229),
	118: uint32(3243),
	119: uint32(3255),
	120: uint32(3265),
	121: uint32(3277),
	122: uint32(3287),
	123: uint32(3301),
	124: uint32(3317),
	125: uint32(3329),
	126: uint32(3341),
	127: uint32(3357),
	128: uint32(3367),
	129: uint32(3377),
	130: uint32(3389),
	131: uint32(3401),
	132: uint32(3413),
	133: uint32(3425),
	134: uint32(3441),
	135: uint32(3455),
	136: uint32(3471),
	137: uint32(3485),
	138: uint32(3497),
	139: uint32(3509),
	140: uint32(3523),
	141: uint32(3539),
	142: uint32(3551),
	143: uint32(3561),
	144: uint32(3577),
	145: uint32(3589),
	146: uint32(3601),
	147: uint32(3611),
	148: uint32(3623),
	149: uint32(3636),
	150: uint32(3649),
	151: uint32(3662),
	152: uint32(3675),
	153: uint32(3688),
	154: uint32(3701),
	155: uint32(3714),
	156: uint32(3727),
	157: uint32(3740),
	158: uint32(3753),
	159: uint32(3766),
	160: uint32(3779),
	161: uint32(3792),
	162: uint32(3805),
	163: uint32(3818),
	164: uint32(3829),
	165: uint32(3842),
	166: uint32(3855),
	167: uint32(3868),
	168: uint32(3881),
	169: uint32(3894),
	170: uint32(3907),
	171: uint32(3920),
	172: uint32(3933),
	173: uint32(3946),
	174: uint32(3959),
	175: uint32(3972),
	176: uint32(3985),
	177: uint32(3998),
	178: uint32(4011),
	179: uint32(4024),
	180: uint32(4037),
	181: uint32(4047),
	182: uint32(4057),
	183: uint32(4067),
	184: uint32(4077),
	185: uint32(4085),
	186: uint32(4093),
	187: uint32(4103),
	188: uint32(4113),
	189: uint32(4121),
	190: uint32(4131),
	191: uint32(4141),
	192: uint32(4149),
	193: uint32(4157),
	194: uint32(4165),
	195: uint32(4175),
	196: uint32(4185),
	197: uint32(4195),
	198: uint32(4205),
	199: uint32(4215),
	200: uint32(4223),
	201: uint32(4233),
	202: uint32(4241),
	203: uint32(4249),
	204: uint32(4257),
	205: uint32(4265),
	206: uint32(4275),
	207: uint32(4285),
	208: uint32(4293),
	209: uint32(4301),
	210: uint32(4311),
	211: uint32(4321),
	212: uint32(4331),
	213: uint32(4341),
	214: uint32(4351),
	215: uint32(4359),
	216: uint32(4369),
	217: uint32(4379),
	218: uint32(4389),
	219: uint32(4399),
	220: uint32(4406),
	221: uint32(4413),
	222: uint32(4420),
	223: uint32(4427),
	224: uint32(4434),
	225: uint32(4441),
	226: uint32(4448),
	227: uint32(4455),
	228: uint32(4462),
	229: uint32(4469),
	230: uint32(4476),
	231: uint32(4483),
	232: uint32(4490),
	233: uint32(4497),
	234: uint32(4504),
	235: uint32(4511),
	236: uint32(4518),
	237: uint32(4525),
	238: uint32(4532),
	239: uint32(4539),
	240: uint32(4546),
	241: uint32(4553),
	242: uint32(4560),
	243: uint32(4567),
	244: uint32(4574),
	245: uint32(4581),
	246: uint32(4588),
	247: uint32(4595),
	248: uint32(4602),
	249: uint32(4609),
	250: uint32(4616),
	251: uint32(4623),
	252: uint32(4630),
	253: uint32(4637),
	254: uint32(4644),
	255: uint32(4651),
	256: uint32(4658),
	257: uint32(4665),
	258: uint32(4672),
	259: uint32(4679),
	260: uint32(4686),
	261: uint32(4693),
	262: uint32(4700),
	263: uint32(4707),
	264: uint32(4714),
	265: uint32(4721),
	266: uint32(4728),
	267: uint32(4735),
	268: uint32(4742),
	269: uint32(4749),
	270: uint32(4756),
	271: uint32(4763),
	272: uint32(4770),
	273: uint32(4777),
	274: uint32(4784),
	275: uint32(4791),
	276: uint32(4798),
	277: uint32(4805),
	278: uint32(4812),
	279: uint32(4819),
	280: uint32(4826),
	281: uint32(4833),
	282: uint32(4840),
	283: uint32(4847),
	284: uint32(4854),
	285: uint32(4861),
}

var ts_parse_actions = [753]TSParseActionEntry{
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
		Fstate: uint16(266),
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
		Fstate: uint16(18),
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
		Fstate: uint16(177),
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
		Fstate: uint16(208),
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
		Fstate: uint16(218),
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
		Fstate: uint16(22),
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
		Fstate: uint16(219),
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
		Fstate: uint16(30),
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
		Fstate: uint16(34),
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
		Fstate: uint16(225),
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
		Fstate: uint16(227),
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
		Fstate: uint16(203),
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
		Fstate: uint16(67),
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
		Fstate: uint16(115),
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
		Fstate: uint16(62),
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
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
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
		Fstate:      uint16(18),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
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
		Fstate:      uint16(177),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
	})))),
	43: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(208),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	44: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	45: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
	})))),
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
		Fstate:      uint16(218),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
	})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
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
		Fstate:      uint16(219),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
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
		Fstate:      uint16(30),
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
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
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
		Fstate:      uint16(34),
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
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
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
		Fstate:      uint16(225),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
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
		Fstate:      uint16(227),
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
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
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
		Fstate:      uint16(203),
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
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
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
		Fstate:      uint16(67),
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
		Fsymbol:      uint16(aux_sym_message_body_repeat1),
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
		Fstate:      uint16(115),
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
		Fstate: uint16(65),
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
		Fstate: uint16(7),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_message_body),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_message_body),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_message),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_message),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(35),
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
		Fcount: uint8(1),
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
		Fstate: uint16(180),
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
		Fstate: uint16(15),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_field),
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
		Fsymbol:      uint16(sym_reserved),
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
		Fsymbol:      uint16(sym_reserved),
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
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_field),
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
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_field),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_field),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_field),
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
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_field),
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
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_field),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_oneof),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_oneof),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_option),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_option),
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
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_field),
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
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_field),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_empty_statement),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_empty_statement),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_oneof),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_oneof),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_map_field),
	})))),
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
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_map_field),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum_body),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum_body),
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
		Fsymbol:      uint16(sym_message_body),
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
		Fsymbol:      uint16(sym_message_body),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_enum_body),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_enum_body),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_oneof_repeat1),
	})))),
	149: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(35),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	150: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	151: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_oneof_repeat1),
	})))),
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
		Fstate:      uint16(180),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_oneof_repeat1),
	})))),
	155: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(208),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_oneof_repeat1),
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
		Fcount: uint8(2),
	}})),
	159: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_oneof_repeat1),
	})))),
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
		Fstate:      uint16(203),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	162: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_oneof_repeat1),
	})))),
	163: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(115),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum),
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
		Fcount: uint8(1),
	}})),
	167: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_field),
	})))),
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
		Fcount: uint8(1),
	}})),
	171: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_field),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(13),
		Fsymbol:      uint16(sym_map_field),
	})))),
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
		Fcount: uint8(1),
	}})),
	175: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(13),
		Fsymbol:      uint16(sym_map_field),
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
		Fstate: uint16(19),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_oneof_field),
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
		Fcount: uint8(1),
	}})),
	181: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_oneof_field),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	183: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(98),
	}})))),
	184: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	185: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(37),
	}})))),
	186: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	187: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_int_lit),
	})))),
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
		Fcount: uint8(1),
	}})),
	189: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_int_lit),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_field_number),
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
		Fcount: uint8(1),
	}})),
	193: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_field_number),
	})))),
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
		Fstate: uint16(150),
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
		Fstate: uint16(111),
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
		Fstate: uint16(47),
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
		Fstate: uint16(40),
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
		Fcount: uint8(1),
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
		Fstate: uint16(74),
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
		Fcount: uint8(1),
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
		Fstate: uint16(104),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(96),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_oneof_field),
	})))),
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
		Fcount: uint8(1),
	}})),
	219: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_oneof_field),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	221: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(42),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(102),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	227: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(102),
	}})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_source_file),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(54),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(101),
	}})))),
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
		Fstate: uint16(220),
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
		Fstate: uint16(163),
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
		Fstate: uint16(213),
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
		Fstate: uint16(212),
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
		Fstate: uint16(211),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_source_file),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	249: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(54),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
	})))),
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
		Fstate:      uint16(101),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	255: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(220),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
	})))),
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
		Fstate:      uint16(163),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	261: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(213),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	264: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(212),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
	})))),
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
		Fstate:      uint16(211),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(286),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym__option_name_repeat1),
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
		Fsymbol:      uint16(aux_sym__option_name_repeat1),
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
		Fstate:      uint16(242),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_full_ident),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(242),
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
		Fstate: uint16(127),
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
		Fcount: uint8(1),
	}})),
	283: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(60),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(230),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_service),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(76),
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
		Fstate: uint16(59),
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
		Fstate: uint16(198),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_syntax),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(166),
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
		Fstate: uint16(75),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_import),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_service_repeat1),
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
		Fstate:      uint16(54),
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
		Fsymbol:      uint16(aux_sym_service_repeat1),
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
		Fstate:      uint16(163),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_service_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	313: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_service_repeat1),
	})))),
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
		Fstate:      uint16(198),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	316: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_import),
		Fproduction_id: uint16(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_package),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_full_ident),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	322: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_service),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_enum_body_repeat1),
	})))),
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
		Fstate:      uint16(127),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_enum_body_repeat1),
	})))),
	328: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(175),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_enum_body_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_enum_body_repeat1),
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
		Fstate:      uint16(230),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	335: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(21),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	337: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(148),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rpc_repeat1),
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
		Fstate:      uint16(54),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rpc_repeat1),
	})))),
	345: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(163),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rpc_repeat1),
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
		Fstate: uint16(123),
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
		Fstate: uint16(121),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(129),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(94),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(191),
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
		Fstate: uint16(156),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_block_lit),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_constant),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_constant),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(31),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	373: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(137),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(179),
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
		Fsymbol:      uint16(sym_bool),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(128),
	}})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_string),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	383: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_block_lit),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(116),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(201),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_message_or_enum_type),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(268),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(115),
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
		Fstate: uint16(107),
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
		Fstate: uint16(119),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	405: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(119),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_message_or_enum_type),
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
		Fchild_count: uint8(13),
		Fsymbol:      uint16(sym_rpc),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_message_or_enum_type),
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
		Fchild_count: uint8(14),
		Fsymbol:      uint16(sym_rpc),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat2),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat2),
	})))),
	419: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(124),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	421: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(190),
	}})))),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(aux_sym_block_lit_repeat2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	425: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(12),
		Fsymbol:      uint16(sym_rpc),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(193),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	429: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(aux_sym_block_lit_repeat2),
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
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_enum_field),
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
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_enum_field),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_enum_field),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_enum_field),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(124),
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
		Fstate: uint16(109),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	443: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(136),
	}})))),
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
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_enum_field),
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
		Fcount: uint8(1),
	}})),
	447: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_enum_field),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(aux_sym_block_lit_repeat2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	453: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(117),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	457: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_lit_repeat2),
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
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_rpc),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	461: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_enum_field),
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
		Fcount: uint8(1),
	}})),
	463: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_enum_field),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	465: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_enum_field),
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
		Fcount: uint8(1),
	}})),
	467: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_enum_field),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	469: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(11),
		Fsymbol:      uint16(sym_rpc),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(187),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	473: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(aux_sym_block_lit_repeat2),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(91),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(33),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_enum_field_repeat1),
	})))),
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
		Fstate:      uint16(125),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_enum_field_repeat1),
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
		Fstate: uint16(142),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_field_options),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(46),
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
		Fstate: uint16(126),
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
		Fstate: uint16(125),
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
		Fstate: uint16(232),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__option_name),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_ranges_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_ranges_repeat1),
	})))),
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
		Fstate:      uint16(112),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_field_names_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	505: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_field_names_repeat1),
	})))),
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
		Fstate:      uint16(279),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(253),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__option_name),
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
		Fsymbol:      uint16(sym_field_options),
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
		Fsymbol:      uint16(aux_sym_field_options_repeat1),
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
		Fstate:      uint16(142),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym_field_options_repeat1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__option_name),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_range),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(114),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_field_names),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_ranges),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(112),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	533: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_ranges),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_field_names),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_lit_repeat1),
	})))),
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
		Fstate:      uint16(46),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_lit_repeat1),
	})))),
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
		Fstate: uint16(140),
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
		Fstate: uint16(243),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__option_name),
	})))),
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
		Fstate: uint16(110),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	552: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_lit_repeat2),
	})))),
	553: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(33),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(2),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(120),
	}})))),
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
		Fstate: uint16(3),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	561: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(92),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(134),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(135),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	569: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(58),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(74),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(13),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(90),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	577: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enum_value_option),
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
		Fstate: uint16(80),
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
		Fstate: uint16(10),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(89),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(273),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(122),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(85),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_range),
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
		Fstate: uint16(87),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(aux_sym_block_lit_repeat2),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_message_or_enum_type_repeat1),
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
		Fstate:      uint16(278),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	602: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_field_option),
	})))),
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
		Fstate: uint16(249),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(254),
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
		Fstate: uint16(257),
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
		Fstate: uint16(146),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(145),
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
		Fstate: uint16(20),
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
		Fstate: uint16(108),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(144),
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
		Fstate: uint16(81),
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
		Fstate: uint16(48),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(217),
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
		Fstate: uint16(14),
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
		Fstate: uint16(12),
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
		Fstate: uint16(260),
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
		Fstate: uint16(38),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(53),
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
		Fstate: uint16(17),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(202),
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
		Fstate: uint16(97),
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
		Fstate: uint16(106),
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
		Fstate: uint16(133),
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
		Fstate: uint16(55),
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
		Fstate: uint16(224),
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
		Fstate: uint16(200),
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
		Fstate: uint16(277),
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
		Fstate: uint16(93),
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
		Fstate: uint16(174),
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
		Fstate: uint16(244),
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
		Fstate: uint16(270),
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
		Fstate: uint16(223),
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
		Fstate: uint16(79),
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
		Fstate: uint16(139),
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
		Fstate: uint16(27),
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
		Fstate: uint16(100),
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
		Fstate: uint16(231),
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
		Fstate: uint16(69),
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
		Fstate: uint16(64),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_service_name),
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
		Fstate: uint16(228),
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
		Fstate: uint16(41),
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
		Fstate: uint16(237),
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
		Fstate: uint16(147),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_message_name),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_enum_name),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(49),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(287),
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
		Fstate: uint16(9),
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
		Fstate: uint16(252),
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
		Fstate: uint16(73),
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
		Fstate: uint16(72),
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
		Fstate: uint16(11),
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
		Fstate: uint16(99),
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
		Fstate: uint16(280),
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
		Fstate: uint16(255),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_message_or_enum_type_repeat1),
	})))),
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
		Fstate: uint16(66),
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
		Fstate: uint16(105),
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
		Fstate: uint16(265),
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
		Fstate: uint16(132),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_rpc_name),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(113),
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
		Fstate: uint16(16),
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
		Fstate: uint16(36),
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
		Fstate: uint16(95),
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
		Fstate: uint16(205),
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
		Fstate: uint16(269),
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
	740: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: uint16(45),
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
		Fstate: uint16(43),
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
		Fstate: uint16(44),
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
		Fstate: uint16(39),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_key_type),
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
		Fstate: uint16(103),
	}})))),
}

func tree_sitter_proto(tls *libc.TLS) (r uintptr) {
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

var __ccgo_ts1 = "end\x00;\x00syntax\x00=\x00\"proto3\"\x00import\x00weak\x00public\x00package\x00option\x00(\x00)\x00.\x00enum\x00{\x00}\x00-\x00[\x00,\x00]\x00message\x00optional\x00repeated\x00oneof\x00map\x00<\x00>\x00int32\x00int64\x00uint32\x00uint64\x00sint32\x00sint64\x00fixed32\x00fixed64\x00sfixed32\x00sfixed64\x00bool\x00string\x00double\x00float\x00bytes\x00reserved\x00to\x00max\x00service\x00rpc\x00stream\x00returns\x00+\x00:\x00identifier\x00true\x00false\x00decimal_lit\x00octal_lit\x00hex_lit\x00float_lit\x00\"\x00string_token1\x00'\x00string_token2\x00escape_sequence\x00comment\x00source_file\x00empty_statement\x00_option_name\x00enum_name\x00enum_body\x00enum_field\x00enum_value_option\x00message_body\x00message_name\x00field\x00field_options\x00field_option\x00oneof_field\x00map_field\x00key_type\x00type\x00ranges\x00range\x00field_names\x00message_or_enum_type\x00field_number\x00service_name\x00rpc_name\x00constant\x00block_lit\x00full_ident\x00int_lit\x00source_file_repeat1\x00_option_name_repeat1\x00enum_body_repeat1\x00enum_field_repeat1\x00message_body_repeat1\x00field_options_repeat1\x00oneof_repeat1\x00ranges_repeat1\x00field_names_repeat1\x00message_or_enum_type_repeat1\x00service_repeat1\x00rpc_repeat1\x00block_lit_repeat1\x00block_lit_repeat2\x00string_repeat1\x00string_repeat2\x00path\x00"
