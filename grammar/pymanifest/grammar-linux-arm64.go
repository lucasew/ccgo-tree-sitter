// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-pymanifest/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-pymanifest -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-pymanifest/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_pymanifest

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
const LANGUAGE_VERSION = 14
const LARGE_STATE_COUNT = 4
const MAX_ALIAS_SEQUENCE_LENGTH = 5
const MAX_RESERVED_WORD_SET_SIZE = 0
const PRODUCTION_ID_COUNT = 5
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 114
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 45
const TOKEN_COUNT = 23
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

const anon_sym_include = 1
const anon_sym_exclude = 2
const anon_sym_recursive_DASHinclude = 3
const anon_sym_recursive_DASHexclude = 4
const anon_sym_global_DASHinclude = 5
const anon_sym_global_DASHexclude = 6
const anon_sym_graft = 7
const anon_sym_prune = 8
const sym__space = 9
const aux_sym__end_of_line_token1 = 10
const anon_sym_BSLASH = 11
const aux_sym__pattern_token1 = 12
const sym_glob = 13
const sym_dir_sep = 14
const aux_sym_escaped_char_token1 = 15
const anon_sym_LBRACK = 16
const anon_sym_BANG = 17
const anon_sym_DASH = 18
const anon_sym_RBRACK = 19
const aux_sym__seq_char_token1 = 20
const aux_sym__seq_char_token2 = 21
const sym_comment = 22
const sym_manifest = 23
const sym_command = 24
const sym__include = 25
const sym__exclude = 26
const sym__recursive_include = 27
const sym__recursive_exclude = 28
const sym__global_include = 29
const sym__global_exclude = 30
const sym__graft = 31
const sym__prune = 32
const sym__end_of_line = 33
const sym_linebreak = 34
const sym_pattern = 35
const aux_sym__pattern = 36
const sym_escaped_char = 37
const sym_char_sequence = 38
const sym_char_range = 39
const sym__seq_char = 40
const aux_sym_manifest_repeat1 = 41
const aux_sym__include_repeat1 = 42
const aux_sym_pattern_repeat1 = 43
const aux_sym_char_sequence_repeat1 = 44

var ts_symbol_names = [45]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 4,
	3:  __ccgo_ts + 4,
	4:  __ccgo_ts + 4,
	5:  __ccgo_ts + 4,
	6:  __ccgo_ts + 4,
	7:  __ccgo_ts + 4,
	8:  __ccgo_ts + 4,
	9:  __ccgo_ts + 12,
	10: __ccgo_ts + 19,
	11: __ccgo_ts + 39,
	12: __ccgo_ts + 41,
	13: __ccgo_ts + 57,
	14: __ccgo_ts + 62,
	15: __ccgo_ts + 70,
	16: __ccgo_ts + 90,
	17: __ccgo_ts + 92,
	18: __ccgo_ts + 94,
	19: __ccgo_ts + 96,
	20: __ccgo_ts + 98,
	21: __ccgo_ts + 115,
	22: __ccgo_ts + 132,
	23: __ccgo_ts + 140,
	24: __ccgo_ts + 149,
	25: __ccgo_ts + 157,
	26: __ccgo_ts + 166,
	27: __ccgo_ts + 175,
	28: __ccgo_ts + 194,
	29: __ccgo_ts + 213,
	30: __ccgo_ts + 229,
	31: __ccgo_ts + 245,
	32: __ccgo_ts + 252,
	33: __ccgo_ts + 259,
	34: __ccgo_ts + 272,
	35: __ccgo_ts + 282,
	36: __ccgo_ts + 290,
	37: __ccgo_ts + 299,
	38: __ccgo_ts + 312,
	39: __ccgo_ts + 326,
	40: __ccgo_ts + 337,
	41: __ccgo_ts + 347,
	42: __ccgo_ts + 364,
	43: __ccgo_ts + 381,
	44: __ccgo_ts + 397,
}

var ts_symbol_map = [45]TSSymbol{
	1:  uint16(anon_sym_include),
	2:  uint16(anon_sym_include),
	3:  uint16(anon_sym_include),
	4:  uint16(anon_sym_include),
	5:  uint16(anon_sym_include),
	6:  uint16(anon_sym_include),
	7:  uint16(anon_sym_include),
	8:  uint16(anon_sym_include),
	9:  uint16(sym__space),
	10: uint16(aux_sym__end_of_line_token1),
	11: uint16(anon_sym_BSLASH),
	12: uint16(aux_sym__pattern_token1),
	13: uint16(sym_glob),
	14: uint16(sym_dir_sep),
	15: uint16(aux_sym_escaped_char_token1),
	16: uint16(anon_sym_LBRACK),
	17: uint16(anon_sym_BANG),
	18: uint16(anon_sym_DASH),
	19: uint16(anon_sym_RBRACK),
	20: uint16(aux_sym__seq_char_token1),
	21: uint16(aux_sym__seq_char_token2),
	22: uint16(sym_comment),
	23: uint16(sym_manifest),
	24: uint16(sym_command),
	25: uint16(sym__include),
	26: uint16(sym__exclude),
	27: uint16(sym__recursive_include),
	28: uint16(sym__recursive_exclude),
	29: uint16(sym__global_include),
	30: uint16(sym__global_exclude),
	31: uint16(sym__graft),
	32: uint16(sym__prune),
	33: uint16(sym__end_of_line),
	34: uint16(sym_linebreak),
	35: uint16(sym_pattern),
	36: uint16(aux_sym__pattern),
	37: uint16(sym_escaped_char),
	38: uint16(sym_char_sequence),
	39: uint16(sym_char_range),
	40: uint16(sym__seq_char),
	41: uint16(aux_sym_manifest_repeat1),
	42: uint16(aux_sym__include_repeat1),
	43: uint16(aux_sym_pattern_repeat1),
	44: uint16(aux_sym_char_sequence_repeat1),
}

var ts_symbol_metadata = [45]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	2: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	3: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	4: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	5: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	6: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	7: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	8: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	9: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	10: {},
	11: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	12: {},
	13: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	14: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	15: {},
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
	21: {},
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	26: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	27: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	28: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	29: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	30: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	31: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	32: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	33: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	35: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	36: {},
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	41: {},
	42: {},
	43: {},
	44: {},
}

type ts_field_identifiers = int32

const field_dir_pattern = 1

var ts_field_names = [2]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 419,
}

var ts_field_map_slices = [5]TSMapSlice{
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
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [5]TSFieldMapEntry{
	0: {
		Ffield_id:  uint16(field_dir_pattern),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	1: {
		Ffield_id:    uint16(field_dir_pattern),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	2: {
		Ffield_id:    uint16(field_dir_pattern),
		Fchild_index: uint8(1),
	},
	3: {
		Ffield_id:    uint16(field_dir_pattern),
		Fchild_index: uint8(1),
	},
	4: {
		Ffield_id:    uint16(field_dir_pattern),
		Fchild_index: uint8(2),
	},
}

var ts_alias_sequences = [5][5]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [114]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(3),
	4:   uint16(4),
	5:   uint16(5),
	6:   uint16(5),
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
	39:  uint16(28),
	40:  uint16(33),
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
	51:  uint16(13),
	52:  uint16(42),
	53:  uint16(45),
	54:  uint16(46),
	55:  uint16(27),
	56:  uint16(27),
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
	68:  uint16(49),
	69:  uint16(69),
	70:  uint16(50),
	71:  uint16(60),
	72:  uint16(59),
	73:  uint16(13),
	74:  uint16(74),
	75:  uint16(63),
	76:  uint16(62),
	77:  uint16(77),
	78:  uint16(77),
	79:  uint16(79),
	80:  uint16(80),
	81:  uint16(81),
	82:  uint16(79),
	83:  uint16(80),
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
	95:  uint16(90),
	96:  uint16(96),
	97:  uint16(96),
	98:  uint16(98),
	99:  uint16(99),
	100: uint16(100),
	101: uint16(101),
	102: uint16(102),
	103: uint16(103),
	104: uint16(104),
	105: uint16(104),
	106: uint16(104),
	107: uint16(107),
	108: uint16(108),
	109: uint16(107),
	110: uint16(110),
	111: uint16(110),
	112: uint16(107),
	113: uint16(110),
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
			state = uint16(68)
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
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(2):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
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
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(3):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
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
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('!') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(66)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('!') && lookahead != int32('#') && lookahead != int32('\\') && lookahead != int32(']') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('-') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(90)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('#') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('-') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('-') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('a') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('a') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('b') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('c') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('c') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('c') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('c') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('c') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('c') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('c') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('d') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('d') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('d') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('d') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('d') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('d') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('e') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('e') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('e') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('e') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('e') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('e') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('e') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('e') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('e') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('e') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('e') {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('f') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('i') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('l') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('l') {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('l') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('l') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('l') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('l') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('l') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('l') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('n') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('n') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('n') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('n') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('o') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('r') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('r') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('s') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('t') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('u') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('u') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('u') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('u') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('u') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('u') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('u') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('u') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('v') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('x') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('x') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('x') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('!') || lookahead == int32('#') || lookahead == int32('-') || int32('[') <= lookahead && lookahead <= int32(']') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(67):
		if eof != 0 {
			state = uint16(68)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
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
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_include)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_exclude)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_recursive_DASHinclude)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_recursive_DASHexclude)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_global_DASHinclude)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_global_DASHexclude)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_graft)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_prune)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__space)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__end_of_line_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') || lookahead == int32('[') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__pattern_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__pattern_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') || lookahead == int32('[') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_glob)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_glob)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_dir_sep)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_escaped_char_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__seq_char_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__seq_char_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(93)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [26]uint16_t{
	0:  uint16('\n'),
	1:  uint16(78),
	2:  uint16('\r'),
	3:  uint16(1),
	4:  uint16('!'),
	5:  uint16(88),
	6:  uint16('#'),
	7:  uint16(93),
	8:  uint16('*'),
	9:  uint16(84),
	10: uint16('-'),
	11: uint16(89),
	12: uint16('/'),
	13: uint16(85),
	14: uint16('?'),
	15: uint16(83),
	16: uint16('['),
	17: uint16(87),
	18: uint16('\\'),
	19: uint16(80),
	20: uint16(']'),
	21: uint16(90),
	22: uint16('\t'),
	23: uint16(77),
	24: uint16(' '),
	25: uint16(77),
}

var map_token1 = [20]uint16_t{
	0:  uint16('\n'),
	1:  uint16(78),
	2:  uint16('\r'),
	3:  uint16(1),
	4:  uint16('#'),
	5:  uint16(93),
	6:  uint16('*'),
	7:  uint16(84),
	8:  uint16('/'),
	9:  uint16(85),
	10: uint16('?'),
	11: uint16(83),
	12: uint16('['),
	13: uint16(87),
	14: uint16('\\'),
	15: uint16(80),
	16: uint16('\t'),
	17: uint16(77),
	18: uint16(' '),
	19: uint16(77),
}

var map_token2 = [20]uint16_t{
	0:  uint16('\n'),
	1:  uint16(78),
	2:  uint16('\r'),
	3:  uint16(1),
	4:  uint16('#'),
	5:  uint16(93),
	6:  uint16('*'),
	7:  uint16(84),
	8:  uint16('/'),
	9:  uint16(85),
	10: uint16('?'),
	11: uint16(83),
	12: uint16('['),
	13: uint16(87),
	14: uint16('\\'),
	15: uint16(82),
	16: uint16('\t'),
	17: uint16(77),
	18: uint16(' '),
	19: uint16(77),
}

var map_token3 = [24]uint16_t{
	0:  uint16('\n'),
	1:  uint16(78),
	2:  uint16('\r'),
	3:  uint16(1),
	4:  uint16('#'),
	5:  uint16(93),
	6:  uint16('/'),
	7:  uint16(85),
	8:  uint16('\\'),
	9:  uint16(79),
	10: uint16('e'),
	11: uint16(63),
	12: uint16('g'),
	13: uint16(37),
	14: uint16('i'),
	15: uint16(46),
	16: uint16('p'),
	17: uint16(50),
	18: uint16('r'),
	19: uint16(31),
	20: uint16('\t'),
	21: uint16(77),
	22: uint16(' '),
	23: uint16(77),
}

var ts_lex_modes = [114]TSLexMode{
	0: {},
	1: {
		Flex_state: uint16(67),
	},
	2: {
		Flex_state: uint16(67),
	},
	3: {
		Flex_state: uint16(67),
	},
	4: {
		Flex_state: uint16(67),
	},
	5: {
		Flex_state: uint16(2),
	},
	6: {
		Flex_state: uint16(3),
	},
	7: {
		Flex_state: uint16(67),
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
		Flex_state: uint16(3),
	},
	12: {
		Flex_state: uint16(3),
	},
	13: {
		Flex_state: uint16(67),
	},
	14: {
		Flex_state: uint16(67),
	},
	15: {
		Flex_state: uint16(67),
	},
	16: {
		Flex_state: uint16(67),
	},
	17: {
		Flex_state: uint16(67),
	},
	18: {
		Flex_state: uint16(67),
	},
	19: {
		Flex_state: uint16(67),
	},
	20: {
		Flex_state: uint16(67),
	},
	21: {
		Flex_state: uint16(67),
	},
	22: {
		Flex_state: uint16(67),
	},
	23: {
		Flex_state: uint16(67),
	},
	24: {
		Flex_state: uint16(67),
	},
	25: {
		Flex_state: uint16(67),
	},
	26: {
		Flex_state: uint16(67),
	},
	27: {
		Flex_state: uint16(67),
	},
	28: {
		Flex_state: uint16(2),
	},
	29: {
		Flex_state: uint16(67),
	},
	30: {
		Flex_state: uint16(67),
	},
	31: {
		Flex_state: uint16(67),
	},
	32: {
		Flex_state: uint16(67),
	},
	33: {
		Flex_state: uint16(2),
	},
	34: {
		Flex_state: uint16(67),
	},
	35: {
		Flex_state: uint16(67),
	},
	36: {
		Flex_state: uint16(67),
	},
	37: {
		Flex_state: uint16(67),
	},
	38: {
		Flex_state: uint16(2),
	},
	39: {
		Flex_state: uint16(3),
	},
	40: {
		Flex_state: uint16(3),
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
		Flex_state: uint16(3),
	},
	53: {
		Flex_state: uint16(3),
	},
	54: {
		Flex_state: uint16(3),
	},
	55: {
		Flex_state: uint16(2),
	},
	56: {
		Flex_state: uint16(3),
	},
	57: {
		Flex_state: uint16(67),
	},
	58: {
		Flex_state: uint16(5),
	},
	59: {
		Flex_state: uint16(5),
	},
	60: {
		Flex_state: uint16(3),
	},
	61: {
		Flex_state: uint16(67),
	},
	62: {
		Flex_state: uint16(4),
	},
	63: {
		Flex_state: uint16(5),
	},
	64: {
		Flex_state: uint16(67),
	},
	65: {
		Flex_state: uint16(67),
	},
	66: {
		Flex_state: uint16(67),
	},
	67: {
		Flex_state: uint16(67),
	},
	68: {
		Flex_state: uint16(3),
	},
	69: {
		Flex_state: uint16(67),
	},
	70: {
		Flex_state: uint16(3),
	},
	71: {
		Flex_state: uint16(3),
	},
	72: {
		Flex_state: uint16(5),
	},
	73: {
		Flex_state: uint16(3),
	},
	74: {
		Flex_state: uint16(67),
	},
	75: {
		Flex_state: uint16(5),
	},
	76: {
		Flex_state: uint16(4),
	},
	77: {
		Flex_state: uint16(5),
	},
	78: {
		Flex_state: uint16(5),
	},
	79: {
		Flex_state: uint16(67),
	},
	80: {
		Flex_state: uint16(67),
	},
	81: {
		Flex_state: uint16(67),
	},
	82: {
		Flex_state: uint16(67),
	},
	83: {
		Flex_state: uint16(67),
	},
	84: {
		Flex_state: uint16(67),
	},
	85: {
		Flex_state: uint16(67),
	},
	86: {
		Flex_state: uint16(5),
	},
	87: {
		Flex_state: uint16(67),
	},
	88: {},
	89: {
		Flex_state: uint16(5),
	},
	90: {},
	91: {
		Flex_state: uint16(67),
	},
	92: {
		Flex_state: uint16(67),
	},
	93: {
		Flex_state: uint16(67),
	},
	94: {},
	95: {},
	96: {},
	97: {},
	98: {
		Flex_state: uint16(67),
	},
	99: {
		Flex_state: uint16(67),
	},
	100: {
		Flex_state: uint16(67),
	},
	101: {
		Flex_state: uint16(67),
	},
	102: {
		Flex_state: uint16(5),
	},
	103: {
		Flex_state: uint16(67),
	},
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
}

var ts_parse_table = [4][45]uint16_t{
	0: {
		0:  uint16(1),
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
		22: uint16(1),
	},
	1: {
		0:  uint16(3),
		1:  uint16(5),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		5:  uint16(13),
		6:  uint16(15),
		7:  uint16(17),
		8:  uint16(19),
		9:  uint16(21),
		10: uint16(23),
		22: uint16(25),
		23: uint16(108),
		24: uint16(2),
		25: uint16(23),
		26: uint16(23),
		27: uint16(19),
		28: uint16(20),
		29: uint16(23),
		30: uint16(23),
		31: uint16(36),
		32: uint16(7),
		33: uint16(2),
		41: uint16(2),
	},
	2: {
		0:  uint16(27),
		1:  uint16(5),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(11),
		5:  uint16(13),
		6:  uint16(15),
		7:  uint16(17),
		8:  uint16(19),
		9:  uint16(21),
		10: uint16(29),
		22: uint16(25),
		24: uint16(3),
		25: uint16(23),
		26: uint16(23),
		27: uint16(19),
		28: uint16(20),
		29: uint16(23),
		30: uint16(23),
		31: uint16(36),
		32: uint16(7),
		33: uint16(3),
		41: uint16(3),
	},
	3: {
		0:  uint16(31),
		1:  uint16(33),
		2:  uint16(36),
		3:  uint16(39),
		4:  uint16(42),
		5:  uint16(45),
		6:  uint16(48),
		7:  uint16(51),
		8:  uint16(54),
		9:  uint16(57),
		10: uint16(60),
		22: uint16(63),
		24: uint16(3),
		25: uint16(23),
		26: uint16(23),
		27: uint16(19),
		28: uint16(20),
		29: uint16(23),
		30: uint16(23),
		31: uint16(36),
		32: uint16(7),
		33: uint16(3),
		41: uint16(3),
	},
}

var ts_small_parse_table = [1779]uint16_t{
	0:    uint16(15),
	1:    uint16(5),
	2:    uint16(1),
	3:    uint16(anon_sym_include),
	4:    uint16(7),
	5:    uint16(1),
	6:    uint16(anon_sym_exclude),
	7:    uint16(9),
	8:    uint16(1),
	9:    uint16(anon_sym_recursive_DASHinclude),
	10:   uint16(11),
	11:   uint16(1),
	12:   uint16(anon_sym_recursive_DASHexclude),
	13:   uint16(13),
	14:   uint16(1),
	15:   uint16(anon_sym_global_DASHinclude),
	16:   uint16(15),
	17:   uint16(1),
	18:   uint16(anon_sym_global_DASHexclude),
	19:   uint16(17),
	20:   uint16(1),
	21:   uint16(anon_sym_graft),
	22:   uint16(19),
	23:   uint16(1),
	24:   uint16(anon_sym_prune),
	25:   uint16(66),
	26:   uint16(1),
	27:   uint16(aux_sym__end_of_line_token1),
	28:   uint16(68),
	29:   uint16(1),
	30:   uint16(sym_comment),
	31:   uint16(15),
	32:   uint16(1),
	33:   uint16(sym__recursive_include),
	34:   uint16(16),
	35:   uint16(1),
	36:   uint16(sym__recursive_exclude),
	37:   uint16(17),
	38:   uint16(1),
	39:   uint16(sym__graft),
	40:   uint16(18),
	41:   uint16(1),
	42:   uint16(sym__prune),
	43:   uint16(14),
	44:   uint16(4),
	45:   uint16(sym__include),
	46:   uint16(sym__exclude),
	47:   uint16(sym__global_include),
	48:   uint16(sym__global_exclude),
	49:   uint16(9),
	50:   uint16(72),
	51:   uint16(1),
	52:   uint16(anon_sym_BSLASH),
	53:   uint16(74),
	54:   uint16(1),
	55:   uint16(aux_sym__pattern_token1),
	56:   uint16(76),
	57:   uint16(1),
	58:   uint16(sym_glob),
	59:   uint16(78),
	60:   uint16(1),
	61:   uint16(sym_dir_sep),
	62:   uint16(80),
	63:   uint16(1),
	64:   uint16(aux_sym_escaped_char_token1),
	65:   uint16(82),
	66:   uint16(1),
	67:   uint16(anon_sym_LBRACK),
	68:   uint16(79),
	69:   uint16(1),
	70:   uint16(aux_sym_pattern_repeat1),
	71:   uint16(70),
	72:   uint16(3),
	73:   uint16(sym__space),
	74:   uint16(aux_sym__end_of_line_token1),
	75:   uint16(sym_comment),
	76:   uint16(28),
	77:   uint16(3),
	78:   uint16(aux_sym__pattern),
	79:   uint16(sym_escaped_char),
	80:   uint16(sym_char_sequence),
	81:   uint16(8),
	82:   uint16(84),
	83:   uint16(1),
	84:   uint16(aux_sym__pattern_token1),
	85:   uint16(86),
	86:   uint16(1),
	87:   uint16(sym_glob),
	88:   uint16(88),
	89:   uint16(1),
	90:   uint16(sym_dir_sep),
	91:   uint16(90),
	92:   uint16(1),
	93:   uint16(aux_sym_escaped_char_token1),
	94:   uint16(92),
	95:   uint16(1),
	96:   uint16(anon_sym_LBRACK),
	97:   uint16(82),
	98:   uint16(1),
	99:   uint16(aux_sym_pattern_repeat1),
	100:  uint16(70),
	101:  uint16(3),
	102:  uint16(sym__space),
	103:  uint16(aux_sym__end_of_line_token1),
	104:  uint16(sym_comment),
	105:  uint16(39),
	106:  uint16(3),
	107:  uint16(aux_sym__pattern),
	108:  uint16(sym_escaped_char),
	109:  uint16(sym_char_sequence),
	110:  uint16(1),
	111:  uint16(94),
	112:  uint16(12),
	114:  uint16(anon_sym_include),
	115:  uint16(anon_sym_exclude),
	116:  uint16(anon_sym_recursive_DASHinclude),
	117:  uint16(anon_sym_recursive_DASHexclude),
	118:  uint16(anon_sym_global_DASHinclude),
	119:  uint16(anon_sym_global_DASHexclude),
	120:  uint16(anon_sym_graft),
	121:  uint16(anon_sym_prune),
	122:  uint16(sym__space),
	123:  uint16(aux_sym__end_of_line_token1),
	124:  uint16(sym_comment),
	125:  uint16(8),
	126:  uint16(80),
	127:  uint16(1),
	128:  uint16(aux_sym_escaped_char_token1),
	129:  uint16(82),
	130:  uint16(1),
	131:  uint16(anon_sym_LBRACK),
	132:  uint16(98),
	133:  uint16(1),
	134:  uint16(anon_sym_BSLASH),
	135:  uint16(100),
	136:  uint16(1),
	137:  uint16(aux_sym__pattern_token1),
	138:  uint16(102),
	139:  uint16(1),
	140:  uint16(sym_glob),
	141:  uint16(98),
	142:  uint16(1),
	143:  uint16(sym_pattern),
	144:  uint16(96),
	145:  uint16(3),
	146:  uint16(sym__space),
	147:  uint16(aux_sym__end_of_line_token1),
	148:  uint16(sym_comment),
	149:  uint16(5),
	150:  uint16(3),
	151:  uint16(aux_sym__pattern),
	152:  uint16(sym_escaped_char),
	153:  uint16(sym_char_sequence),
	154:  uint16(10),
	155:  uint16(80),
	156:  uint16(1),
	157:  uint16(aux_sym_escaped_char_token1),
	158:  uint16(82),
	159:  uint16(1),
	160:  uint16(anon_sym_LBRACK),
	161:  uint16(100),
	162:  uint16(1),
	163:  uint16(aux_sym__pattern_token1),
	164:  uint16(102),
	165:  uint16(1),
	166:  uint16(sym_glob),
	167:  uint16(104),
	168:  uint16(1),
	169:  uint16(sym__space),
	170:  uint16(106),
	171:  uint16(1),
	172:  uint16(anon_sym_BSLASH),
	173:  uint16(8),
	174:  uint16(1),
	175:  uint16(sym_linebreak),
	176:  uint16(69),
	177:  uint16(1),
	178:  uint16(aux_sym__include_repeat1),
	179:  uint16(85),
	180:  uint16(1),
	181:  uint16(sym_pattern),
	182:  uint16(5),
	183:  uint16(3),
	184:  uint16(aux_sym__pattern),
	185:  uint16(sym_escaped_char),
	186:  uint16(sym_char_sequence),
	187:  uint16(10),
	188:  uint16(80),
	189:  uint16(1),
	190:  uint16(aux_sym_escaped_char_token1),
	191:  uint16(82),
	192:  uint16(1),
	193:  uint16(anon_sym_LBRACK),
	194:  uint16(100),
	195:  uint16(1),
	196:  uint16(aux_sym__pattern_token1),
	197:  uint16(102),
	198:  uint16(1),
	199:  uint16(sym_glob),
	200:  uint16(104),
	201:  uint16(1),
	202:  uint16(sym__space),
	203:  uint16(106),
	204:  uint16(1),
	205:  uint16(anon_sym_BSLASH),
	206:  uint16(8),
	207:  uint16(1),
	208:  uint16(sym_linebreak),
	209:  uint16(74),
	210:  uint16(1),
	211:  uint16(aux_sym__include_repeat1),
	212:  uint16(93),
	213:  uint16(1),
	214:  uint16(sym_pattern),
	215:  uint16(5),
	216:  uint16(3),
	217:  uint16(aux_sym__pattern),
	218:  uint16(sym_escaped_char),
	219:  uint16(sym_char_sequence),
	220:  uint16(10),
	221:  uint16(25),
	222:  uint16(1),
	223:  uint16(sym_comment),
	224:  uint16(90),
	225:  uint16(1),
	226:  uint16(aux_sym_escaped_char_token1),
	227:  uint16(92),
	228:  uint16(1),
	229:  uint16(anon_sym_LBRACK),
	230:  uint16(108),
	231:  uint16(1),
	232:  uint16(sym__space),
	233:  uint16(110),
	234:  uint16(1),
	235:  uint16(aux_sym__end_of_line_token1),
	236:  uint16(112),
	237:  uint16(1),
	238:  uint16(aux_sym__pattern_token1),
	239:  uint16(114),
	240:  uint16(1),
	241:  uint16(sym_glob),
	242:  uint16(25),
	243:  uint16(1),
	244:  uint16(sym__end_of_line),
	245:  uint16(88),
	246:  uint16(1),
	247:  uint16(sym_pattern),
	248:  uint16(6),
	249:  uint16(3),
	250:  uint16(aux_sym__pattern),
	251:  uint16(sym_escaped_char),
	252:  uint16(sym_char_sequence),
	253:  uint16(10),
	254:  uint16(25),
	255:  uint16(1),
	256:  uint16(sym_comment),
	257:  uint16(90),
	258:  uint16(1),
	259:  uint16(aux_sym_escaped_char_token1),
	260:  uint16(92),
	261:  uint16(1),
	262:  uint16(anon_sym_LBRACK),
	263:  uint16(108),
	264:  uint16(1),
	265:  uint16(sym__space),
	266:  uint16(112),
	267:  uint16(1),
	268:  uint16(aux_sym__pattern_token1),
	269:  uint16(114),
	270:  uint16(1),
	271:  uint16(sym_glob),
	272:  uint16(116),
	273:  uint16(1),
	274:  uint16(aux_sym__end_of_line_token1),
	275:  uint16(26),
	276:  uint16(1),
	277:  uint16(sym__end_of_line),
	278:  uint16(94),
	279:  uint16(1),
	280:  uint16(sym_pattern),
	281:  uint16(6),
	282:  uint16(3),
	283:  uint16(aux_sym__pattern),
	284:  uint16(sym_escaped_char),
	285:  uint16(sym_char_sequence),
	286:  uint16(1),
	287:  uint16(118),
	288:  uint16(12),
	290:  uint16(anon_sym_include),
	291:  uint16(anon_sym_exclude),
	292:  uint16(anon_sym_recursive_DASHinclude),
	293:  uint16(anon_sym_recursive_DASHexclude),
	294:  uint16(anon_sym_global_DASHinclude),
	295:  uint16(anon_sym_global_DASHexclude),
	296:  uint16(anon_sym_graft),
	297:  uint16(anon_sym_prune),
	298:  uint16(sym__space),
	299:  uint16(aux_sym__end_of_line_token1),
	300:  uint16(sym_comment),
	301:  uint16(1),
	302:  uint16(120),
	303:  uint16(12),
	305:  uint16(anon_sym_include),
	306:  uint16(anon_sym_exclude),
	307:  uint16(anon_sym_recursive_DASHinclude),
	308:  uint16(anon_sym_recursive_DASHexclude),
	309:  uint16(anon_sym_global_DASHinclude),
	310:  uint16(anon_sym_global_DASHexclude),
	311:  uint16(anon_sym_graft),
	312:  uint16(anon_sym_prune),
	313:  uint16(sym__space),
	314:  uint16(aux_sym__end_of_line_token1),
	315:  uint16(sym_comment),
	316:  uint16(1),
	317:  uint16(122),
	318:  uint16(12),
	320:  uint16(anon_sym_include),
	321:  uint16(anon_sym_exclude),
	322:  uint16(anon_sym_recursive_DASHinclude),
	323:  uint16(anon_sym_recursive_DASHexclude),
	324:  uint16(anon_sym_global_DASHinclude),
	325:  uint16(anon_sym_global_DASHexclude),
	326:  uint16(anon_sym_graft),
	327:  uint16(anon_sym_prune),
	328:  uint16(sym__space),
	329:  uint16(aux_sym__end_of_line_token1),
	330:  uint16(sym_comment),
	331:  uint16(1),
	332:  uint16(122),
	333:  uint16(12),
	335:  uint16(anon_sym_include),
	336:  uint16(anon_sym_exclude),
	337:  uint16(anon_sym_recursive_DASHinclude),
	338:  uint16(anon_sym_recursive_DASHexclude),
	339:  uint16(anon_sym_global_DASHinclude),
	340:  uint16(anon_sym_global_DASHexclude),
	341:  uint16(anon_sym_graft),
	342:  uint16(anon_sym_prune),
	343:  uint16(sym__space),
	344:  uint16(aux_sym__end_of_line_token1),
	345:  uint16(sym_comment),
	346:  uint16(1),
	347:  uint16(122),
	348:  uint16(12),
	350:  uint16(anon_sym_include),
	351:  uint16(anon_sym_exclude),
	352:  uint16(anon_sym_recursive_DASHinclude),
	353:  uint16(anon_sym_recursive_DASHexclude),
	354:  uint16(anon_sym_global_DASHinclude),
	355:  uint16(anon_sym_global_DASHexclude),
	356:  uint16(anon_sym_graft),
	357:  uint16(anon_sym_prune),
	358:  uint16(sym__space),
	359:  uint16(aux_sym__end_of_line_token1),
	360:  uint16(sym_comment),
	361:  uint16(1),
	362:  uint16(122),
	363:  uint16(12),
	365:  uint16(anon_sym_include),
	366:  uint16(anon_sym_exclude),
	367:  uint16(anon_sym_recursive_DASHinclude),
	368:  uint16(anon_sym_recursive_DASHexclude),
	369:  uint16(anon_sym_global_DASHinclude),
	370:  uint16(anon_sym_global_DASHexclude),
	371:  uint16(anon_sym_graft),
	372:  uint16(anon_sym_prune),
	373:  uint16(sym__space),
	374:  uint16(aux_sym__end_of_line_token1),
	375:  uint16(sym_comment),
	376:  uint16(1),
	377:  uint16(94),
	378:  uint16(12),
	380:  uint16(anon_sym_include),
	381:  uint16(anon_sym_exclude),
	382:  uint16(anon_sym_recursive_DASHinclude),
	383:  uint16(anon_sym_recursive_DASHexclude),
	384:  uint16(anon_sym_global_DASHinclude),
	385:  uint16(anon_sym_global_DASHexclude),
	386:  uint16(anon_sym_graft),
	387:  uint16(anon_sym_prune),
	388:  uint16(sym__space),
	389:  uint16(aux_sym__end_of_line_token1),
	390:  uint16(sym_comment),
	391:  uint16(1),
	392:  uint16(94),
	393:  uint16(12),
	395:  uint16(anon_sym_include),
	396:  uint16(anon_sym_exclude),
	397:  uint16(anon_sym_recursive_DASHinclude),
	398:  uint16(anon_sym_recursive_DASHexclude),
	399:  uint16(anon_sym_global_DASHinclude),
	400:  uint16(anon_sym_global_DASHexclude),
	401:  uint16(anon_sym_graft),
	402:  uint16(anon_sym_prune),
	403:  uint16(sym__space),
	404:  uint16(aux_sym__end_of_line_token1),
	405:  uint16(sym_comment),
	406:  uint16(1),
	407:  uint16(124),
	408:  uint16(12),
	410:  uint16(anon_sym_include),
	411:  uint16(anon_sym_exclude),
	412:  uint16(anon_sym_recursive_DASHinclude),
	413:  uint16(anon_sym_recursive_DASHexclude),
	414:  uint16(anon_sym_global_DASHinclude),
	415:  uint16(anon_sym_global_DASHexclude),
	416:  uint16(anon_sym_graft),
	417:  uint16(anon_sym_prune),
	418:  uint16(sym__space),
	419:  uint16(aux_sym__end_of_line_token1),
	420:  uint16(sym_comment),
	421:  uint16(1),
	422:  uint16(126),
	423:  uint16(12),
	425:  uint16(anon_sym_include),
	426:  uint16(anon_sym_exclude),
	427:  uint16(anon_sym_recursive_DASHinclude),
	428:  uint16(anon_sym_recursive_DASHexclude),
	429:  uint16(anon_sym_global_DASHinclude),
	430:  uint16(anon_sym_global_DASHexclude),
	431:  uint16(anon_sym_graft),
	432:  uint16(anon_sym_prune),
	433:  uint16(sym__space),
	434:  uint16(aux_sym__end_of_line_token1),
	435:  uint16(sym_comment),
	436:  uint16(1),
	437:  uint16(128),
	438:  uint16(12),
	440:  uint16(anon_sym_include),
	441:  uint16(anon_sym_exclude),
	442:  uint16(anon_sym_recursive_DASHinclude),
	443:  uint16(anon_sym_recursive_DASHexclude),
	444:  uint16(anon_sym_global_DASHinclude),
	445:  uint16(anon_sym_global_DASHexclude),
	446:  uint16(anon_sym_graft),
	447:  uint16(anon_sym_prune),
	448:  uint16(sym__space),
	449:  uint16(aux_sym__end_of_line_token1),
	450:  uint16(sym_comment),
	451:  uint16(1),
	452:  uint16(130),
	453:  uint16(12),
	455:  uint16(anon_sym_include),
	456:  uint16(anon_sym_exclude),
	457:  uint16(anon_sym_recursive_DASHinclude),
	458:  uint16(anon_sym_recursive_DASHexclude),
	459:  uint16(anon_sym_global_DASHinclude),
	460:  uint16(anon_sym_global_DASHexclude),
	461:  uint16(anon_sym_graft),
	462:  uint16(anon_sym_prune),
	463:  uint16(sym__space),
	464:  uint16(aux_sym__end_of_line_token1),
	465:  uint16(sym_comment),
	466:  uint16(1),
	467:  uint16(132),
	468:  uint16(12),
	470:  uint16(anon_sym_include),
	471:  uint16(anon_sym_exclude),
	472:  uint16(anon_sym_recursive_DASHinclude),
	473:  uint16(anon_sym_recursive_DASHexclude),
	474:  uint16(anon_sym_global_DASHinclude),
	475:  uint16(anon_sym_global_DASHexclude),
	476:  uint16(anon_sym_graft),
	477:  uint16(anon_sym_prune),
	478:  uint16(sym__space),
	479:  uint16(aux_sym__end_of_line_token1),
	480:  uint16(sym_comment),
	481:  uint16(1),
	482:  uint16(134),
	483:  uint16(12),
	485:  uint16(anon_sym_include),
	486:  uint16(anon_sym_exclude),
	487:  uint16(anon_sym_recursive_DASHinclude),
	488:  uint16(anon_sym_recursive_DASHexclude),
	489:  uint16(anon_sym_global_DASHinclude),
	490:  uint16(anon_sym_global_DASHexclude),
	491:  uint16(anon_sym_graft),
	492:  uint16(anon_sym_prune),
	493:  uint16(sym__space),
	494:  uint16(aux_sym__end_of_line_token1),
	495:  uint16(sym_comment),
	496:  uint16(1),
	497:  uint16(136),
	498:  uint16(12),
	500:  uint16(anon_sym_include),
	501:  uint16(anon_sym_exclude),
	502:  uint16(anon_sym_recursive_DASHinclude),
	503:  uint16(anon_sym_recursive_DASHexclude),
	504:  uint16(anon_sym_global_DASHinclude),
	505:  uint16(anon_sym_global_DASHexclude),
	506:  uint16(anon_sym_graft),
	507:  uint16(anon_sym_prune),
	508:  uint16(sym__space),
	509:  uint16(aux_sym__end_of_line_token1),
	510:  uint16(sym_comment),
	511:  uint16(7),
	512:  uint16(140),
	513:  uint16(1),
	514:  uint16(anon_sym_BSLASH),
	515:  uint16(142),
	516:  uint16(1),
	517:  uint16(aux_sym__pattern_token1),
	518:  uint16(145),
	519:  uint16(1),
	520:  uint16(sym_glob),
	521:  uint16(148),
	522:  uint16(1),
	523:  uint16(aux_sym_escaped_char_token1),
	524:  uint16(151),
	525:  uint16(1),
	526:  uint16(anon_sym_LBRACK),
	527:  uint16(28),
	528:  uint16(3),
	529:  uint16(aux_sym__pattern),
	530:  uint16(sym_escaped_char),
	531:  uint16(sym_char_sequence),
	532:  uint16(138),
	533:  uint16(4),
	534:  uint16(sym__space),
	535:  uint16(aux_sym__end_of_line_token1),
	536:  uint16(sym_dir_sep),
	537:  uint16(sym_comment),
	538:  uint16(1),
	539:  uint16(154),
	540:  uint16(12),
	542:  uint16(anon_sym_include),
	543:  uint16(anon_sym_exclude),
	544:  uint16(anon_sym_recursive_DASHinclude),
	545:  uint16(anon_sym_recursive_DASHexclude),
	546:  uint16(anon_sym_global_DASHinclude),
	547:  uint16(anon_sym_global_DASHexclude),
	548:  uint16(anon_sym_graft),
	549:  uint16(anon_sym_prune),
	550:  uint16(sym__space),
	551:  uint16(aux_sym__end_of_line_token1),
	552:  uint16(sym_comment),
	553:  uint16(1),
	554:  uint16(156),
	555:  uint16(12),
	557:  uint16(anon_sym_include),
	558:  uint16(anon_sym_exclude),
	559:  uint16(anon_sym_recursive_DASHinclude),
	560:  uint16(anon_sym_recursive_DASHexclude),
	561:  uint16(anon_sym_global_DASHinclude),
	562:  uint16(anon_sym_global_DASHexclude),
	563:  uint16(anon_sym_graft),
	564:  uint16(anon_sym_prune),
	565:  uint16(sym__space),
	566:  uint16(aux_sym__end_of_line_token1),
	567:  uint16(sym_comment),
	568:  uint16(1),
	569:  uint16(158),
	570:  uint16(12),
	572:  uint16(anon_sym_include),
	573:  uint16(anon_sym_exclude),
	574:  uint16(anon_sym_recursive_DASHinclude),
	575:  uint16(anon_sym_recursive_DASHexclude),
	576:  uint16(anon_sym_global_DASHinclude),
	577:  uint16(anon_sym_global_DASHexclude),
	578:  uint16(anon_sym_graft),
	579:  uint16(anon_sym_prune),
	580:  uint16(sym__space),
	581:  uint16(aux_sym__end_of_line_token1),
	582:  uint16(sym_comment),
	583:  uint16(1),
	584:  uint16(160),
	585:  uint16(12),
	587:  uint16(anon_sym_include),
	588:  uint16(anon_sym_exclude),
	589:  uint16(anon_sym_recursive_DASHinclude),
	590:  uint16(anon_sym_recursive_DASHexclude),
	591:  uint16(anon_sym_global_DASHinclude),
	592:  uint16(anon_sym_global_DASHexclude),
	593:  uint16(anon_sym_graft),
	594:  uint16(anon_sym_prune),
	595:  uint16(sym__space),
	596:  uint16(aux_sym__end_of_line_token1),
	597:  uint16(sym_comment),
	598:  uint16(7),
	599:  uint16(74),
	600:  uint16(1),
	601:  uint16(aux_sym__pattern_token1),
	602:  uint16(76),
	603:  uint16(1),
	604:  uint16(sym_glob),
	605:  uint16(80),
	606:  uint16(1),
	607:  uint16(aux_sym_escaped_char_token1),
	608:  uint16(82),
	609:  uint16(1),
	610:  uint16(anon_sym_LBRACK),
	611:  uint16(164),
	612:  uint16(1),
	613:  uint16(anon_sym_BSLASH),
	614:  uint16(28),
	615:  uint16(3),
	616:  uint16(aux_sym__pattern),
	617:  uint16(sym_escaped_char),
	618:  uint16(sym_char_sequence),
	619:  uint16(162),
	620:  uint16(4),
	621:  uint16(sym__space),
	622:  uint16(aux_sym__end_of_line_token1),
	623:  uint16(sym_dir_sep),
	624:  uint16(sym_comment),
	625:  uint16(1),
	626:  uint16(166),
	627:  uint16(12),
	629:  uint16(anon_sym_include),
	630:  uint16(anon_sym_exclude),
	631:  uint16(anon_sym_recursive_DASHinclude),
	632:  uint16(anon_sym_recursive_DASHexclude),
	633:  uint16(anon_sym_global_DASHinclude),
	634:  uint16(anon_sym_global_DASHexclude),
	635:  uint16(anon_sym_graft),
	636:  uint16(anon_sym_prune),
	637:  uint16(sym__space),
	638:  uint16(aux_sym__end_of_line_token1),
	639:  uint16(sym_comment),
	640:  uint16(1),
	641:  uint16(168),
	642:  uint16(12),
	644:  uint16(anon_sym_include),
	645:  uint16(anon_sym_exclude),
	646:  uint16(anon_sym_recursive_DASHinclude),
	647:  uint16(anon_sym_recursive_DASHexclude),
	648:  uint16(anon_sym_global_DASHinclude),
	649:  uint16(anon_sym_global_DASHexclude),
	650:  uint16(anon_sym_graft),
	651:  uint16(anon_sym_prune),
	652:  uint16(sym__space),
	653:  uint16(aux_sym__end_of_line_token1),
	654:  uint16(sym_comment),
	655:  uint16(1),
	656:  uint16(94),
	657:  uint16(12),
	659:  uint16(anon_sym_include),
	660:  uint16(anon_sym_exclude),
	661:  uint16(anon_sym_recursive_DASHinclude),
	662:  uint16(anon_sym_recursive_DASHexclude),
	663:  uint16(anon_sym_global_DASHinclude),
	664:  uint16(anon_sym_global_DASHexclude),
	665:  uint16(anon_sym_graft),
	666:  uint16(anon_sym_prune),
	667:  uint16(sym__space),
	668:  uint16(aux_sym__end_of_line_token1),
	669:  uint16(sym_comment),
	670:  uint16(1),
	671:  uint16(170),
	672:  uint16(12),
	674:  uint16(anon_sym_include),
	675:  uint16(anon_sym_exclude),
	676:  uint16(anon_sym_recursive_DASHinclude),
	677:  uint16(anon_sym_recursive_DASHexclude),
	678:  uint16(anon_sym_global_DASHinclude),
	679:  uint16(anon_sym_global_DASHexclude),
	680:  uint16(anon_sym_graft),
	681:  uint16(anon_sym_prune),
	682:  uint16(sym__space),
	683:  uint16(aux_sym__end_of_line_token1),
	684:  uint16(sym_comment),
	685:  uint16(9),
	686:  uint16(66),
	687:  uint16(1),
	688:  uint16(aux_sym__end_of_line_token1),
	689:  uint16(68),
	690:  uint16(1),
	691:  uint16(sym_comment),
	692:  uint16(80),
	693:  uint16(1),
	694:  uint16(aux_sym_escaped_char_token1),
	695:  uint16(82),
	696:  uint16(1),
	697:  uint16(anon_sym_LBRACK),
	698:  uint16(100),
	699:  uint16(1),
	700:  uint16(aux_sym__pattern_token1),
	701:  uint16(102),
	702:  uint16(1),
	703:  uint16(sym_glob),
	704:  uint16(172),
	705:  uint16(1),
	706:  uint16(anon_sym_BSLASH),
	707:  uint16(98),
	708:  uint16(1),
	709:  uint16(sym_pattern),
	710:  uint16(5),
	711:  uint16(3),
	712:  uint16(aux_sym__pattern),
	713:  uint16(sym_escaped_char),
	714:  uint16(sym_char_sequence),
	715:  uint16(6),
	716:  uint16(174),
	717:  uint16(1),
	718:  uint16(aux_sym__pattern_token1),
	719:  uint16(177),
	720:  uint16(1),
	721:  uint16(sym_glob),
	722:  uint16(180),
	723:  uint16(1),
	724:  uint16(aux_sym_escaped_char_token1),
	725:  uint16(183),
	726:  uint16(1),
	727:  uint16(anon_sym_LBRACK),
	728:  uint16(39),
	729:  uint16(3),
	730:  uint16(aux_sym__pattern),
	731:  uint16(sym_escaped_char),
	732:  uint16(sym_char_sequence),
	733:  uint16(138),
	734:  uint16(4),
	735:  uint16(sym__space),
	736:  uint16(aux_sym__end_of_line_token1),
	737:  uint16(sym_dir_sep),
	738:  uint16(sym_comment),
	739:  uint16(6),
	740:  uint16(84),
	741:  uint16(1),
	742:  uint16(aux_sym__pattern_token1),
	743:  uint16(86),
	744:  uint16(1),
	745:  uint16(sym_glob),
	746:  uint16(90),
	747:  uint16(1),
	748:  uint16(aux_sym_escaped_char_token1),
	749:  uint16(92),
	750:  uint16(1),
	751:  uint16(anon_sym_LBRACK),
	752:  uint16(39),
	753:  uint16(3),
	754:  uint16(aux_sym__pattern),
	755:  uint16(sym_escaped_char),
	756:  uint16(sym_char_sequence),
	757:  uint16(162),
	758:  uint16(4),
	759:  uint16(sym__space),
	760:  uint16(aux_sym__end_of_line_token1),
	761:  uint16(sym_dir_sep),
	762:  uint16(sym_comment),
	763:  uint16(7),
	764:  uint16(80),
	765:  uint16(1),
	766:  uint16(aux_sym_escaped_char_token1),
	767:  uint16(82),
	768:  uint16(1),
	769:  uint16(anon_sym_LBRACK),
	770:  uint16(100),
	771:  uint16(1),
	772:  uint16(aux_sym__pattern_token1),
	773:  uint16(102),
	774:  uint16(1),
	775:  uint16(sym_glob),
	776:  uint16(172),
	777:  uint16(1),
	778:  uint16(anon_sym_BSLASH),
	779:  uint16(98),
	780:  uint16(1),
	781:  uint16(sym_pattern),
	782:  uint16(5),
	783:  uint16(3),
	784:  uint16(aux_sym__pattern),
	785:  uint16(sym_escaped_char),
	786:  uint16(sym_char_sequence),
	787:  uint16(2),
	788:  uint16(188),
	789:  uint16(2),
	790:  uint16(anon_sym_BSLASH),
	791:  uint16(aux_sym__pattern_token1),
	792:  uint16(186),
	793:  uint16(7),
	794:  uint16(sym__space),
	795:  uint16(aux_sym__end_of_line_token1),
	796:  uint16(sym_glob),
	797:  uint16(sym_dir_sep),
	798:  uint16(aux_sym_escaped_char_token1),
	799:  uint16(anon_sym_LBRACK),
	800:  uint16(sym_comment),
	801:  uint16(7),
	802:  uint16(80),
	803:  uint16(1),
	804:  uint16(aux_sym_escaped_char_token1),
	805:  uint16(82),
	806:  uint16(1),
	807:  uint16(anon_sym_LBRACK),
	808:  uint16(100),
	809:  uint16(1),
	810:  uint16(aux_sym__pattern_token1),
	811:  uint16(102),
	812:  uint16(1),
	813:  uint16(sym_glob),
	814:  uint16(172),
	815:  uint16(1),
	816:  uint16(anon_sym_BSLASH),
	817:  uint16(85),
	818:  uint16(1),
	819:  uint16(sym_pattern),
	820:  uint16(5),
	821:  uint16(3),
	822:  uint16(aux_sym__pattern),
	823:  uint16(sym_escaped_char),
	824:  uint16(sym_char_sequence),
	825:  uint16(7),
	826:  uint16(90),
	827:  uint16(1),
	828:  uint16(aux_sym_escaped_char_token1),
	829:  uint16(92),
	830:  uint16(1),
	831:  uint16(anon_sym_LBRACK),
	832:  uint16(112),
	833:  uint16(1),
	834:  uint16(aux_sym__pattern_token1),
	835:  uint16(114),
	836:  uint16(1),
	837:  uint16(sym_glob),
	838:  uint16(190),
	839:  uint16(1),
	840:  uint16(anon_sym_BSLASH),
	841:  uint16(94),
	842:  uint16(1),
	843:  uint16(sym_pattern),
	844:  uint16(6),
	845:  uint16(3),
	846:  uint16(aux_sym__pattern),
	847:  uint16(sym_escaped_char),
	848:  uint16(sym_char_sequence),
	849:  uint16(2),
	850:  uint16(194),
	851:  uint16(2),
	852:  uint16(anon_sym_BSLASH),
	853:  uint16(aux_sym__pattern_token1),
	854:  uint16(192),
	855:  uint16(7),
	856:  uint16(sym__space),
	857:  uint16(aux_sym__end_of_line_token1),
	858:  uint16(sym_glob),
	859:  uint16(sym_dir_sep),
	860:  uint16(aux_sym_escaped_char_token1),
	861:  uint16(anon_sym_LBRACK),
	862:  uint16(sym_comment),
	863:  uint16(2),
	864:  uint16(198),
	865:  uint16(2),
	866:  uint16(anon_sym_BSLASH),
	867:  uint16(aux_sym__pattern_token1),
	868:  uint16(196),
	869:  uint16(7),
	870:  uint16(sym__space),
	871:  uint16(aux_sym__end_of_line_token1),
	872:  uint16(sym_glob),
	873:  uint16(sym_dir_sep),
	874:  uint16(aux_sym_escaped_char_token1),
	875:  uint16(anon_sym_LBRACK),
	876:  uint16(sym_comment),
	877:  uint16(7),
	878:  uint16(80),
	879:  uint16(1),
	880:  uint16(aux_sym_escaped_char_token1),
	881:  uint16(82),
	882:  uint16(1),
	883:  uint16(anon_sym_LBRACK),
	884:  uint16(100),
	885:  uint16(1),
	886:  uint16(aux_sym__pattern_token1),
	887:  uint16(102),
	888:  uint16(1),
	889:  uint16(sym_glob),
	890:  uint16(172),
	891:  uint16(1),
	892:  uint16(anon_sym_BSLASH),
	893:  uint16(93),
	894:  uint16(1),
	895:  uint16(sym_pattern),
	896:  uint16(5),
	897:  uint16(3),
	898:  uint16(aux_sym__pattern),
	899:  uint16(sym_escaped_char),
	900:  uint16(sym_char_sequence),
	901:  uint16(7),
	902:  uint16(90),
	903:  uint16(1),
	904:  uint16(aux_sym_escaped_char_token1),
	905:  uint16(92),
	906:  uint16(1),
	907:  uint16(anon_sym_LBRACK),
	908:  uint16(112),
	909:  uint16(1),
	910:  uint16(aux_sym__pattern_token1),
	911:  uint16(114),
	912:  uint16(1),
	913:  uint16(sym_glob),
	914:  uint16(190),
	915:  uint16(1),
	916:  uint16(anon_sym_BSLASH),
	917:  uint16(88),
	918:  uint16(1),
	919:  uint16(sym_pattern),
	920:  uint16(6),
	921:  uint16(3),
	922:  uint16(aux_sym__pattern),
	923:  uint16(sym_escaped_char),
	924:  uint16(sym_char_sequence),
	925:  uint16(2),
	926:  uint16(202),
	927:  uint16(2),
	928:  uint16(anon_sym_BSLASH),
	929:  uint16(aux_sym__pattern_token1),
	930:  uint16(200),
	931:  uint16(6),
	932:  uint16(sym__space),
	933:  uint16(aux_sym__end_of_line_token1),
	934:  uint16(sym_glob),
	935:  uint16(aux_sym_escaped_char_token1),
	936:  uint16(anon_sym_LBRACK),
	937:  uint16(sym_comment),
	938:  uint16(2),
	939:  uint16(206),
	940:  uint16(2),
	941:  uint16(anon_sym_BSLASH),
	942:  uint16(aux_sym__pattern_token1),
	943:  uint16(204),
	944:  uint16(6),
	945:  uint16(sym__space),
	946:  uint16(aux_sym__end_of_line_token1),
	947:  uint16(sym_glob),
	948:  uint16(aux_sym_escaped_char_token1),
	949:  uint16(anon_sym_LBRACK),
	950:  uint16(sym_comment),
	951:  uint16(2),
	952:  uint16(208),
	953:  uint16(2),
	954:  uint16(anon_sym_BSLASH),
	955:  uint16(aux_sym__pattern_token1),
	956:  uint16(118),
	957:  uint16(6),
	958:  uint16(sym__space),
	959:  uint16(aux_sym__end_of_line_token1),
	960:  uint16(sym_glob),
	961:  uint16(aux_sym_escaped_char_token1),
	962:  uint16(anon_sym_LBRACK),
	963:  uint16(sym_comment),
	964:  uint16(2),
	965:  uint16(188),
	966:  uint16(1),
	967:  uint16(aux_sym__pattern_token1),
	968:  uint16(186),
	969:  uint16(7),
	970:  uint16(sym__space),
	971:  uint16(aux_sym__end_of_line_token1),
	972:  uint16(sym_glob),
	973:  uint16(sym_dir_sep),
	974:  uint16(aux_sym_escaped_char_token1),
	975:  uint16(anon_sym_LBRACK),
	976:  uint16(sym_comment),
	977:  uint16(2),
	978:  uint16(194),
	979:  uint16(1),
	980:  uint16(aux_sym__pattern_token1),
	981:  uint16(192),
	982:  uint16(7),
	983:  uint16(sym__space),
	984:  uint16(aux_sym__end_of_line_token1),
	985:  uint16(sym_glob),
	986:  uint16(sym_dir_sep),
	987:  uint16(aux_sym_escaped_char_token1),
	988:  uint16(anon_sym_LBRACK),
	989:  uint16(sym_comment),
	990:  uint16(2),
	991:  uint16(198),
	992:  uint16(1),
	993:  uint16(aux_sym__pattern_token1),
	994:  uint16(196),
	995:  uint16(7),
	996:  uint16(sym__space),
	997:  uint16(aux_sym__end_of_line_token1),
	998:  uint16(sym_glob),
	999:  uint16(sym_dir_sep),
	1000: uint16(aux_sym_escaped_char_token1),
	1001: uint16(anon_sym_LBRACK),
	1002: uint16(sym_comment),
	1003: uint16(2),
	1004: uint16(210),
	1005: uint16(2),
	1006: uint16(anon_sym_BSLASH),
	1007: uint16(aux_sym__pattern_token1),
	1008: uint16(136),
	1009: uint16(6),
	1010: uint16(sym__space),
	1011: uint16(aux_sym__end_of_line_token1),
	1012: uint16(sym_glob),
	1013: uint16(aux_sym_escaped_char_token1),
	1014: uint16(anon_sym_LBRACK),
	1015: uint16(sym_comment),
	1016: uint16(2),
	1017: uint16(210),
	1018: uint16(1),
	1019: uint16(aux_sym__pattern_token1),
	1020: uint16(136),
	1021: uint16(6),
	1022: uint16(sym__space),
	1023: uint16(aux_sym__end_of_line_token1),
	1024: uint16(sym_glob),
	1025: uint16(aux_sym_escaped_char_token1),
	1026: uint16(anon_sym_LBRACK),
	1027: uint16(sym_comment),
	1028: uint16(7),
	1029: uint16(25),
	1030: uint16(1),
	1031: uint16(sym_comment),
	1032: uint16(212),
	1033: uint16(1),
	1034: uint16(sym__space),
	1035: uint16(214),
	1036: uint16(1),
	1037: uint16(aux_sym__end_of_line_token1),
	1038: uint16(216),
	1039: uint16(1),
	1040: uint16(anon_sym_BSLASH),
	1041: uint16(8),
	1042: uint16(1),
	1043: uint16(sym_linebreak),
	1044: uint16(21),
	1045: uint16(1),
	1046: uint16(sym__end_of_line),
	1047: uint16(81),
	1048: uint16(1),
	1049: uint16(aux_sym__include_repeat1),
	1050: uint16(5),
	1051: uint16(218),
	1052: uint16(1),
	1053: uint16(anon_sym_DASH),
	1054: uint16(221),
	1055: uint16(1),
	1056: uint16(anon_sym_RBRACK),
	1057: uint16(89),
	1058: uint16(1),
	1059: uint16(sym__seq_char),
	1060: uint16(223),
	1061: uint16(2),
	1062: uint16(aux_sym__seq_char_token1),
	1063: uint16(aux_sym__seq_char_token2),
	1064: uint16(58),
	1065: uint16(2),
	1066: uint16(sym_char_range),
	1067: uint16(aux_sym_char_sequence_repeat1),
	1068: uint16(5),
	1069: uint16(226),
	1070: uint16(1),
	1071: uint16(anon_sym_DASH),
	1072: uint16(228),
	1073: uint16(1),
	1074: uint16(anon_sym_RBRACK),
	1075: uint16(89),
	1076: uint16(1),
	1077: uint16(sym__seq_char),
	1078: uint16(230),
	1079: uint16(2),
	1080: uint16(aux_sym__seq_char_token1),
	1081: uint16(aux_sym__seq_char_token2),
	1082: uint16(58),
	1083: uint16(2),
	1084: uint16(sym_char_range),
	1085: uint16(aux_sym_char_sequence_repeat1),
	1086: uint16(5),
	1087: uint16(80),
	1088: uint16(1),
	1089: uint16(aux_sym_escaped_char_token1),
	1090: uint16(82),
	1091: uint16(1),
	1092: uint16(anon_sym_LBRACK),
	1093: uint16(232),
	1094: uint16(1),
	1095: uint16(aux_sym__pattern_token1),
	1096: uint16(234),
	1097: uint16(1),
	1098: uint16(sym_glob),
	1099: uint16(33),
	1100: uint16(3),
	1101: uint16(aux_sym__pattern),
	1102: uint16(sym_escaped_char),
	1103: uint16(sym_char_sequence),
	1104: uint16(7),
	1105: uint16(25),
	1106: uint16(1),
	1107: uint16(sym_comment),
	1108: uint16(212),
	1109: uint16(1),
	1110: uint16(sym__space),
	1111: uint16(216),
	1112: uint16(1),
	1113: uint16(anon_sym_BSLASH),
	1114: uint16(236),
	1115: uint16(1),
	1116: uint16(aux_sym__end_of_line_token1),
	1117: uint16(8),
	1118: uint16(1),
	1119: uint16(sym_linebreak),
	1120: uint16(34),
	1121: uint16(1),
	1122: uint16(sym__end_of_line),
	1123: uint16(81),
	1124: uint16(1),
	1125: uint16(aux_sym__include_repeat1),
	1126: uint16(6),
	1127: uint16(230),
	1128: uint16(1),
	1129: uint16(aux_sym__seq_char_token2),
	1130: uint16(238),
	1131: uint16(1),
	1132: uint16(anon_sym_BANG),
	1133: uint16(240),
	1134: uint16(1),
	1135: uint16(anon_sym_DASH),
	1136: uint16(242),
	1137: uint16(1),
	1138: uint16(aux_sym__seq_char_token1),
	1139: uint16(89),
	1140: uint16(1),
	1141: uint16(sym__seq_char),
	1142: uint16(59),
	1143: uint16(2),
	1144: uint16(sym_char_range),
	1145: uint16(aux_sym_char_sequence_repeat1),
	1146: uint16(5),
	1147: uint16(226),
	1148: uint16(1),
	1149: uint16(anon_sym_DASH),
	1150: uint16(244),
	1151: uint16(1),
	1152: uint16(anon_sym_RBRACK),
	1153: uint16(89),
	1154: uint16(1),
	1155: uint16(sym__seq_char),
	1156: uint16(230),
	1157: uint16(2),
	1158: uint16(aux_sym__seq_char_token1),
	1159: uint16(aux_sym__seq_char_token2),
	1160: uint16(58),
	1161: uint16(2),
	1162: uint16(sym_char_range),
	1163: uint16(aux_sym_char_sequence_repeat1),
	1164: uint16(7),
	1165: uint16(25),
	1166: uint16(1),
	1167: uint16(sym_comment),
	1168: uint16(212),
	1169: uint16(1),
	1170: uint16(sym__space),
	1171: uint16(216),
	1172: uint16(1),
	1173: uint16(anon_sym_BSLASH),
	1174: uint16(246),
	1175: uint16(1),
	1176: uint16(aux_sym__end_of_line_token1),
	1177: uint16(8),
	1178: uint16(1),
	1179: uint16(sym_linebreak),
	1180: uint16(22),
	1181: uint16(1),
	1182: uint16(sym__end_of_line),
	1183: uint16(81),
	1184: uint16(1),
	1185: uint16(aux_sym__include_repeat1),
	1186: uint16(7),
	1187: uint16(25),
	1188: uint16(1),
	1189: uint16(sym_comment),
	1190: uint16(212),
	1191: uint16(1),
	1192: uint16(sym__space),
	1193: uint16(216),
	1194: uint16(1),
	1195: uint16(anon_sym_BSLASH),
	1196: uint16(248),
	1197: uint16(1),
	1198: uint16(aux_sym__end_of_line_token1),
	1199: uint16(8),
	1200: uint16(1),
	1201: uint16(sym_linebreak),
	1202: uint16(35),
	1203: uint16(1),
	1204: uint16(sym__end_of_line),
	1205: uint16(81),
	1206: uint16(1),
	1207: uint16(aux_sym__include_repeat1),
	1208: uint16(7),
	1209: uint16(25),
	1210: uint16(1),
	1211: uint16(sym_comment),
	1212: uint16(212),
	1213: uint16(1),
	1214: uint16(sym__space),
	1215: uint16(216),
	1216: uint16(1),
	1217: uint16(anon_sym_BSLASH),
	1218: uint16(250),
	1219: uint16(1),
	1220: uint16(aux_sym__end_of_line_token1),
	1221: uint16(8),
	1222: uint16(1),
	1223: uint16(sym_linebreak),
	1224: uint16(37),
	1225: uint16(1),
	1226: uint16(sym__end_of_line),
	1227: uint16(81),
	1228: uint16(1),
	1229: uint16(aux_sym__include_repeat1),
	1230: uint16(7),
	1231: uint16(25),
	1232: uint16(1),
	1233: uint16(sym_comment),
	1234: uint16(212),
	1235: uint16(1),
	1236: uint16(sym__space),
	1237: uint16(216),
	1238: uint16(1),
	1239: uint16(anon_sym_BSLASH),
	1240: uint16(252),
	1241: uint16(1),
	1242: uint16(aux_sym__end_of_line_token1),
	1243: uint16(8),
	1244: uint16(1),
	1245: uint16(sym_linebreak),
	1246: uint16(24),
	1247: uint16(1),
	1248: uint16(sym__end_of_line),
	1249: uint16(81),
	1250: uint16(1),
	1251: uint16(aux_sym__include_repeat1),
	1252: uint16(2),
	1253: uint16(202),
	1254: uint16(1),
	1255: uint16(aux_sym__pattern_token1),
	1256: uint16(200),
	1257: uint16(6),
	1258: uint16(sym__space),
	1259: uint16(aux_sym__end_of_line_token1),
	1260: uint16(sym_glob),
	1261: uint16(aux_sym_escaped_char_token1),
	1262: uint16(anon_sym_LBRACK),
	1263: uint16(sym_comment),
	1264: uint16(7),
	1265: uint16(25),
	1266: uint16(1),
	1267: uint16(sym_comment),
	1268: uint16(212),
	1269: uint16(1),
	1270: uint16(sym__space),
	1271: uint16(216),
	1272: uint16(1),
	1273: uint16(anon_sym_BSLASH),
	1274: uint16(254),
	1275: uint16(1),
	1276: uint16(aux_sym__end_of_line_token1),
	1277: uint16(8),
	1278: uint16(1),
	1279: uint16(sym_linebreak),
	1280: uint16(29),
	1281: uint16(1),
	1282: uint16(sym__end_of_line),
	1283: uint16(81),
	1284: uint16(1),
	1285: uint16(aux_sym__include_repeat1),
	1286: uint16(2),
	1287: uint16(206),
	1288: uint16(1),
	1289: uint16(aux_sym__pattern_token1),
	1290: uint16(204),
	1291: uint16(6),
	1292: uint16(sym__space),
	1293: uint16(aux_sym__end_of_line_token1),
	1294: uint16(sym_glob),
	1295: uint16(aux_sym_escaped_char_token1),
	1296: uint16(anon_sym_LBRACK),
	1297: uint16(sym_comment),
	1298: uint16(5),
	1299: uint16(90),
	1300: uint16(1),
	1301: uint16(aux_sym_escaped_char_token1),
	1302: uint16(92),
	1303: uint16(1),
	1304: uint16(anon_sym_LBRACK),
	1305: uint16(256),
	1306: uint16(1),
	1307: uint16(aux_sym__pattern_token1),
	1308: uint16(258),
	1309: uint16(1),
	1310: uint16(sym_glob),
	1311: uint16(40),
	1312: uint16(3),
	1313: uint16(aux_sym__pattern),
	1314: uint16(sym_escaped_char),
	1315: uint16(sym_char_sequence),
	1316: uint16(5),
	1317: uint16(226),
	1318: uint16(1),
	1319: uint16(anon_sym_DASH),
	1320: uint16(260),
	1321: uint16(1),
	1322: uint16(anon_sym_RBRACK),
	1323: uint16(89),
	1324: uint16(1),
	1325: uint16(sym__seq_char),
	1326: uint16(230),
	1327: uint16(2),
	1328: uint16(aux_sym__seq_char_token1),
	1329: uint16(aux_sym__seq_char_token2),
	1330: uint16(58),
	1331: uint16(2),
	1332: uint16(sym_char_range),
	1333: uint16(aux_sym_char_sequence_repeat1),
	1334: uint16(2),
	1335: uint16(208),
	1336: uint16(1),
	1337: uint16(aux_sym__pattern_token1),
	1338: uint16(118),
	1339: uint16(6),
	1340: uint16(sym__space),
	1341: uint16(aux_sym__end_of_line_token1),
	1342: uint16(sym_glob),
	1343: uint16(aux_sym_escaped_char_token1),
	1344: uint16(anon_sym_LBRACK),
	1345: uint16(sym_comment),
	1346: uint16(7),
	1347: uint16(25),
	1348: uint16(1),
	1349: uint16(sym_comment),
	1350: uint16(212),
	1351: uint16(1),
	1352: uint16(sym__space),
	1353: uint16(216),
	1354: uint16(1),
	1355: uint16(anon_sym_BSLASH),
	1356: uint16(262),
	1357: uint16(1),
	1358: uint16(aux_sym__end_of_line_token1),
	1359: uint16(8),
	1360: uint16(1),
	1361: uint16(sym_linebreak),
	1362: uint16(30),
	1363: uint16(1),
	1364: uint16(sym__end_of_line),
	1365: uint16(81),
	1366: uint16(1),
	1367: uint16(aux_sym__include_repeat1),
	1368: uint16(5),
	1369: uint16(226),
	1370: uint16(1),
	1371: uint16(anon_sym_DASH),
	1372: uint16(264),
	1373: uint16(1),
	1374: uint16(anon_sym_RBRACK),
	1375: uint16(89),
	1376: uint16(1),
	1377: uint16(sym__seq_char),
	1378: uint16(230),
	1379: uint16(2),
	1380: uint16(aux_sym__seq_char_token1),
	1381: uint16(aux_sym__seq_char_token2),
	1382: uint16(58),
	1383: uint16(2),
	1384: uint16(sym_char_range),
	1385: uint16(aux_sym_char_sequence_repeat1),
	1386: uint16(6),
	1387: uint16(230),
	1388: uint16(1),
	1389: uint16(aux_sym__seq_char_token2),
	1390: uint16(242),
	1391: uint16(1),
	1392: uint16(aux_sym__seq_char_token1),
	1393: uint16(266),
	1394: uint16(1),
	1395: uint16(anon_sym_BANG),
	1396: uint16(268),
	1397: uint16(1),
	1398: uint16(anon_sym_DASH),
	1399: uint16(89),
	1400: uint16(1),
	1401: uint16(sym__seq_char),
	1402: uint16(72),
	1403: uint16(2),
	1404: uint16(sym_char_range),
	1405: uint16(aux_sym_char_sequence_repeat1),
	1406: uint16(4),
	1407: uint16(270),
	1408: uint16(1),
	1409: uint16(anon_sym_DASH),
	1410: uint16(89),
	1411: uint16(1),
	1412: uint16(sym__seq_char),
	1413: uint16(230),
	1414: uint16(2),
	1415: uint16(aux_sym__seq_char_token1),
	1416: uint16(aux_sym__seq_char_token2),
	1417: uint16(75),
	1418: uint16(2),
	1419: uint16(sym_char_range),
	1420: uint16(aux_sym_char_sequence_repeat1),
	1421: uint16(4),
	1422: uint16(272),
	1423: uint16(1),
	1424: uint16(anon_sym_DASH),
	1425: uint16(89),
	1426: uint16(1),
	1427: uint16(sym__seq_char),
	1428: uint16(230),
	1429: uint16(2),
	1430: uint16(aux_sym__seq_char_token1),
	1431: uint16(aux_sym__seq_char_token2),
	1432: uint16(63),
	1433: uint16(2),
	1434: uint16(sym_char_range),
	1435: uint16(aux_sym_char_sequence_repeat1),
	1436: uint16(3),
	1437: uint16(78),
	1438: uint16(1),
	1439: uint16(sym_dir_sep),
	1440: uint16(80),
	1441: uint16(1),
	1442: uint16(aux_sym_pattern_repeat1),
	1443: uint16(274),
	1444: uint16(4),
	1445: uint16(sym__space),
	1446: uint16(aux_sym__end_of_line_token1),
	1447: uint16(anon_sym_BSLASH),
	1448: uint16(sym_comment),
	1449: uint16(3),
	1450: uint16(276),
	1451: uint16(1),
	1452: uint16(sym_dir_sep),
	1453: uint16(80),
	1454: uint16(1),
	1455: uint16(aux_sym_pattern_repeat1),
	1456: uint16(162),
	1457: uint16(4),
	1458: uint16(sym__space),
	1459: uint16(aux_sym__end_of_line_token1),
	1460: uint16(anon_sym_BSLASH),
	1461: uint16(sym_comment),
	1462: uint16(5),
	1463: uint16(279),
	1464: uint16(1),
	1465: uint16(sym__space),
	1466: uint16(284),
	1467: uint16(1),
	1468: uint16(anon_sym_BSLASH),
	1469: uint16(8),
	1470: uint16(1),
	1471: uint16(sym_linebreak),
	1472: uint16(81),
	1473: uint16(1),
	1474: uint16(aux_sym__include_repeat1),
	1475: uint16(282),
	1476: uint16(2),
	1477: uint16(aux_sym__end_of_line_token1),
	1478: uint16(sym_comment),
	1479: uint16(3),
	1480: uint16(88),
	1481: uint16(1),
	1482: uint16(sym_dir_sep),
	1483: uint16(83),
	1484: uint16(1),
	1485: uint16(aux_sym_pattern_repeat1),
	1486: uint16(274),
	1487: uint16(3),
	1488: uint16(sym__space),
	1489: uint16(aux_sym__end_of_line_token1),
	1490: uint16(sym_comment),
	1491: uint16(3),
	1492: uint16(287),
	1493: uint16(1),
	1494: uint16(sym_dir_sep),
	1495: uint16(83),
	1496: uint16(1),
	1497: uint16(aux_sym_pattern_repeat1),
	1498: uint16(162),
	1499: uint16(3),
	1500: uint16(sym__space),
	1501: uint16(aux_sym__end_of_line_token1),
	1502: uint16(sym_comment),
	1503: uint16(4),
	1504: uint16(104),
	1505: uint16(1),
	1506: uint16(sym__space),
	1507: uint16(216),
	1508: uint16(1),
	1509: uint16(anon_sym_BSLASH),
	1510: uint16(8),
	1511: uint16(1),
	1512: uint16(sym_linebreak),
	1513: uint16(67),
	1514: uint16(1),
	1515: uint16(aux_sym__include_repeat1),
	1516: uint16(4),
	1517: uint16(104),
	1518: uint16(1),
	1519: uint16(sym__space),
	1520: uint16(216),
	1521: uint16(1),
	1522: uint16(anon_sym_BSLASH),
	1523: uint16(8),
	1524: uint16(1),
	1525: uint16(sym_linebreak),
	1526: uint16(61),
	1527: uint16(1),
	1528: uint16(aux_sym__include_repeat1),
	1529: uint16(1),
	1530: uint16(290),
	1531: uint16(4),
	1532: uint16(anon_sym_DASH),
	1533: uint16(anon_sym_RBRACK),
	1534: uint16(aux_sym__seq_char_token1),
	1535: uint16(aux_sym__seq_char_token2),
	1536: uint16(4),
	1537: uint16(104),
	1538: uint16(1),
	1539: uint16(sym__space),
	1540: uint16(216),
	1541: uint16(1),
	1542: uint16(anon_sym_BSLASH),
	1543: uint16(8),
	1544: uint16(1),
	1545: uint16(sym_linebreak),
	1546: uint16(57),
	1547: uint16(1),
	1548: uint16(aux_sym__include_repeat1),
	1549: uint16(4),
	1550: uint16(25),
	1551: uint16(1),
	1552: uint16(sym_comment),
	1553: uint16(108),
	1554: uint16(1),
	1555: uint16(sym__space),
	1556: uint16(292),
	1557: uint16(1),
	1558: uint16(aux_sym__end_of_line_token1),
	1559: uint16(31),
	1560: uint16(1),
	1561: uint16(sym__end_of_line),
	1562: uint16(2),
	1563: uint16(294),
	1564: uint16(1),
	1565: uint16(anon_sym_DASH),
	1566: uint16(296),
	1567: uint16(3),
	1568: uint16(anon_sym_RBRACK),
	1569: uint16(aux_sym__seq_char_token1),
	1570: uint16(aux_sym__seq_char_token2),
	1571: uint16(4),
	1572: uint16(298),
	1573: uint16(1),
	1574: uint16(sym__space),
	1575: uint16(300),
	1576: uint16(1),
	1577: uint16(aux_sym__end_of_line_token1),
	1578: uint16(302),
	1579: uint16(1),
	1580: uint16(sym_comment),
	1581: uint16(50),
	1582: uint16(1),
	1583: uint16(sym__end_of_line),
	1584: uint16(4),
	1585: uint16(104),
	1586: uint16(1),
	1587: uint16(sym__space),
	1588: uint16(216),
	1589: uint16(1),
	1590: uint16(anon_sym_BSLASH),
	1591: uint16(8),
	1592: uint16(1),
	1593: uint16(sym_linebreak),
	1594: uint16(66),
	1595: uint16(1),
	1596: uint16(aux_sym__include_repeat1),
	1597: uint16(4),
	1598: uint16(104),
	1599: uint16(1),
	1600: uint16(sym__space),
	1601: uint16(216),
	1602: uint16(1),
	1603: uint16(anon_sym_BSLASH),
	1604: uint16(8),
	1605: uint16(1),
	1606: uint16(sym_linebreak),
	1607: uint16(64),
	1608: uint16(1),
	1609: uint16(aux_sym__include_repeat1),
	1610: uint16(4),
	1611: uint16(104),
	1612: uint16(1),
	1613: uint16(sym__space),
	1614: uint16(216),
	1615: uint16(1),
	1616: uint16(anon_sym_BSLASH),
	1617: uint16(8),
	1618: uint16(1),
	1619: uint16(sym_linebreak),
	1620: uint16(65),
	1621: uint16(1),
	1622: uint16(aux_sym__include_repeat1),
	1623: uint16(4),
	1624: uint16(25),
	1625: uint16(1),
	1626: uint16(sym_comment),
	1627: uint16(108),
	1628: uint16(1),
	1629: uint16(sym__space),
	1630: uint16(304),
	1631: uint16(1),
	1632: uint16(aux_sym__end_of_line_token1),
	1633: uint16(32),
	1634: uint16(1),
	1635: uint16(sym__end_of_line),
	1636: uint16(4),
	1637: uint16(306),
	1638: uint16(1),
	1639: uint16(sym__space),
	1640: uint16(308),
	1641: uint16(1),
	1642: uint16(aux_sym__end_of_line_token1),
	1643: uint16(310),
	1644: uint16(1),
	1645: uint16(sym_comment),
	1646: uint16(70),
	1647: uint16(1),
	1648: uint16(sym__end_of_line),
	1649: uint16(4),
	1650: uint16(298),
	1651: uint16(1),
	1652: uint16(sym__space),
	1653: uint16(302),
	1654: uint16(1),
	1655: uint16(sym_comment),
	1656: uint16(312),
	1657: uint16(1),
	1658: uint16(aux_sym__end_of_line_token1),
	1659: uint16(49),
	1660: uint16(1),
	1661: uint16(sym__end_of_line),
	1662: uint16(4),
	1663: uint16(306),
	1664: uint16(1),
	1665: uint16(sym__space),
	1666: uint16(310),
	1667: uint16(1),
	1668: uint16(sym_comment),
	1669: uint16(314),
	1670: uint16(1),
	1671: uint16(aux_sym__end_of_line_token1),
	1672: uint16(68),
	1673: uint16(1),
	1674: uint16(sym__end_of_line),
	1675: uint16(1),
	1676: uint16(282),
	1677: uint16(4),
	1678: uint16(sym__space),
	1679: uint16(aux_sym__end_of_line_token1),
	1680: uint16(anon_sym_BSLASH),
	1681: uint16(sym_comment),
	1682: uint16(3),
	1683: uint16(316),
	1684: uint16(1),
	1685: uint16(sym__space),
	1686: uint16(318),
	1687: uint16(1),
	1688: uint16(anon_sym_BSLASH),
	1689: uint16(12),
	1690: uint16(1),
	1691: uint16(sym_linebreak),
	1692: uint16(3),
	1693: uint16(216),
	1694: uint16(1),
	1695: uint16(anon_sym_BSLASH),
	1696: uint16(320),
	1697: uint16(1),
	1698: uint16(sym__space),
	1699: uint16(10),
	1700: uint16(1),
	1701: uint16(sym_linebreak),
	1702: uint16(3),
	1703: uint16(318),
	1704: uint16(1),
	1705: uint16(anon_sym_BSLASH),
	1706: uint16(322),
	1707: uint16(1),
	1708: uint16(sym__space),
	1709: uint16(11),
	1710: uint16(1),
	1711: uint16(sym_linebreak),
	1712: uint16(2),
	1713: uint16(86),
	1714: uint16(1),
	1715: uint16(sym__seq_char),
	1716: uint16(324),
	1717: uint16(2),
	1718: uint16(aux_sym__seq_char_token1),
	1719: uint16(aux_sym__seq_char_token2),
	1720: uint16(3),
	1721: uint16(216),
	1722: uint16(1),
	1723: uint16(anon_sym_BSLASH),
	1724: uint16(326),
	1725: uint16(1),
	1726: uint16(sym__space),
	1727: uint16(9),
	1728: uint16(1),
	1729: uint16(sym_linebreak),
	1730: uint16(2),
	1731: uint16(66),
	1732: uint16(1),
	1733: uint16(aux_sym__end_of_line_token1),
	1734: uint16(68),
	1735: uint16(1),
	1736: uint16(sym_comment),
	1737: uint16(2),
	1738: uint16(328),
	1739: uint16(1),
	1740: uint16(aux_sym__end_of_line_token1),
	1741: uint16(330),
	1742: uint16(1),
	1743: uint16(sym_comment),
	1744: uint16(2),
	1745: uint16(332),
	1746: uint16(1),
	1747: uint16(aux_sym__end_of_line_token1),
	1748: uint16(334),
	1749: uint16(1),
	1750: uint16(sym_comment),
	1751: uint16(1),
	1752: uint16(66),
	1753: uint16(1),
	1754: uint16(aux_sym__end_of_line_token1),
	1755: uint16(1),
	1756: uint16(336),
	1757: uint16(1),
	1759: uint16(1),
	1760: uint16(332),
	1761: uint16(1),
	1762: uint16(aux_sym__end_of_line_token1),
	1763: uint16(1),
	1764: uint16(338),
	1765: uint16(1),
	1766: uint16(aux_sym__end_of_line_token1),
	1767: uint16(1),
	1768: uint16(340),
	1769: uint16(1),
	1770: uint16(aux_sym__end_of_line_token1),
	1771: uint16(1),
	1772: uint16(328),
	1773: uint16(1),
	1774: uint16(aux_sym__end_of_line_token1),
	1775: uint16(1),
	1776: uint16(342),
	1777: uint16(1),
	1778: uint16(aux_sym__end_of_line_token1),
}

var ts_small_parse_table_map = [110]uint32_t{
	1:   uint32(49),
	2:   uint32(81),
	3:   uint32(110),
	4:   uint32(125),
	5:   uint32(154),
	6:   uint32(187),
	7:   uint32(220),
	8:   uint32(253),
	9:   uint32(286),
	10:  uint32(301),
	11:  uint32(316),
	12:  uint32(331),
	13:  uint32(346),
	14:  uint32(361),
	15:  uint32(376),
	16:  uint32(391),
	17:  uint32(406),
	18:  uint32(421),
	19:  uint32(436),
	20:  uint32(451),
	21:  uint32(466),
	22:  uint32(481),
	23:  uint32(496),
	24:  uint32(511),
	25:  uint32(538),
	26:  uint32(553),
	27:  uint32(568),
	28:  uint32(583),
	29:  uint32(598),
	30:  uint32(625),
	31:  uint32(640),
	32:  uint32(655),
	33:  uint32(670),
	34:  uint32(685),
	35:  uint32(715),
	36:  uint32(739),
	37:  uint32(763),
	38:  uint32(787),
	39:  uint32(801),
	40:  uint32(825),
	41:  uint32(849),
	42:  uint32(863),
	43:  uint32(877),
	44:  uint32(901),
	45:  uint32(925),
	46:  uint32(938),
	47:  uint32(951),
	48:  uint32(964),
	49:  uint32(977),
	50:  uint32(990),
	51:  uint32(1003),
	52:  uint32(1016),
	53:  uint32(1028),
	54:  uint32(1050),
	55:  uint32(1068),
	56:  uint32(1086),
	57:  uint32(1104),
	58:  uint32(1126),
	59:  uint32(1146),
	60:  uint32(1164),
	61:  uint32(1186),
	62:  uint32(1208),
	63:  uint32(1230),
	64:  uint32(1252),
	65:  uint32(1264),
	66:  uint32(1286),
	67:  uint32(1298),
	68:  uint32(1316),
	69:  uint32(1334),
	70:  uint32(1346),
	71:  uint32(1368),
	72:  uint32(1386),
	73:  uint32(1406),
	74:  uint32(1421),
	75:  uint32(1436),
	76:  uint32(1449),
	77:  uint32(1462),
	78:  uint32(1479),
	79:  uint32(1491),
	80:  uint32(1503),
	81:  uint32(1516),
	82:  uint32(1529),
	83:  uint32(1536),
	84:  uint32(1549),
	85:  uint32(1562),
	86:  uint32(1571),
	87:  uint32(1584),
	88:  uint32(1597),
	89:  uint32(1610),
	90:  uint32(1623),
	91:  uint32(1636),
	92:  uint32(1649),
	93:  uint32(1662),
	94:  uint32(1675),
	95:  uint32(1682),
	96:  uint32(1692),
	97:  uint32(1702),
	98:  uint32(1712),
	99:  uint32(1720),
	100: uint32(1730),
	101: uint32(1737),
	102: uint32(1744),
	103: uint32(1751),
	104: uint32(1755),
	105: uint32(1759),
	106: uint32(1763),
	107: uint32(1767),
	108: uint32(1771),
	109: uint32(1775),
}

var ts_parse_actions = [344]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_manifest),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(92)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(103)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(101)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(99)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(107)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_manifest),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_manifest_repeat1),
	})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_manifest_repeat1),
	})))),
	35: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	36: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	37: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_manifest_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(92)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	40: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_manifest_repeat1),
	})))),
	41: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(103)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	42: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	43: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_manifest_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(100)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	46: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_manifest_repeat1),
	})))),
	47: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	48: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	49: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_manifest_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(84)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_manifest_repeat1),
	})))),
	53: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(101)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	54: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	55: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_manifest_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(99)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	58: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_manifest_repeat1),
	})))),
	59: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	60: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	61: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_manifest_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(3)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_manifest_repeat1),
	})))),
	65: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(107)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	67: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	68: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_pattern),
	})))),
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
		Fcount: uint8(1),
	}})),
	73: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_pattern),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(28)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(62)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(71)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_command),
		Fproduction_id: uint16(1),
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
		Fsymbol:      uint16(aux_sym__include_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__include_repeat1),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(41)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(104)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym__end_of_line),
	})))),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_command),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:        uint16(sym_command),
		Fproduction_id: uint16(2),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__include),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__exclude),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_command),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__global_exclude),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	133: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__graft),
		Fproduction_id: uint16(3),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__prune),
		Fproduction_id: uint16(3),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__end_of_line),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
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
		Fcount: uint8(1),
	}})),
	141: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__pattern),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(28)),
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
		Fsymbol:      uint16(aux_sym__pattern),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(28)),
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
		Fsymbol:      uint16(aux_sym__pattern),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		Fsymbol:      uint16(aux_sym__pattern),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(62)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__recursive_include),
		Fproduction_id: uint16(3),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__recursive_exclude),
		Fproduction_id: uint16(3),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__graft),
		Fproduction_id: uint16(4),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__prune),
		Fproduction_id: uint16(4),
	})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_pattern_repeat1),
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
		Fcount: uint8(1),
	}})),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_pattern_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym__recursive_include),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym__recursive_exclude),
		Fproduction_id: uint16(4),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__global_include),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Fcount: uint8(2),
	}})),
	175: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(39)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(39)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym__pattern),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(52)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(76)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_escaped_char),
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
		Fsymbol:      uint16(sym_escaped_char),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_char_sequence),
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
		Fcount: uint8(1),
	}})),
	195: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_char_sequence),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_char_sequence),
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
		Fcount: uint8(1),
	}})),
	199: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_char_sequence),
	})))),
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
		Fsymbol:      uint16(sym_linebreak),
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
		Fsymbol:      uint16(sym_linebreak),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_linebreak),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_linebreak),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__end_of_line),
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
		Fcount: uint8(1),
	}})),
	211: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__end_of_line),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_char_sequence_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(58)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_char_sequence_repeat1),
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
		Fsymbol:      uint16(aux_sym_char_sequence_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(89)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(89)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(34)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(89)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(46)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(35)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(54)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_pattern),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_pattern_repeat1),
	})))),
	278: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(60)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	280: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__include_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(41)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__include_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	285: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__include_repeat1),
	})))),
	286: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	288: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_pattern_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(71)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_char_range),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(102)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_char_sequence_repeat1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(106)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(109)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(105)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(70)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(68)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(97)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(111)),
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
	337: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(27)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	341: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

func tree_sitter_pymanifest(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(LANGUAGE_VERSION),
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
	Fprimary_state_ids:         uintptr(unsafe.Pointer(&ts_primary_state_ids)),
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

var __ccgo_ts1 = "end\x00keyword\x00_space\x00_end_of_line_token1\x00\\\x00_pattern_token1\x00glob\x00dir_sep\x00escaped_char_token1\x00[\x00!\x00-\x00]\x00_seq_char_token1\x00_seq_char_token2\x00comment\x00manifest\x00command\x00_include\x00_exclude\x00_recursive_include\x00_recursive_exclude\x00_global_include\x00_global_exclude\x00_graft\x00_prune\x00_end_of_line\x00linebreak\x00pattern\x00_pattern\x00escaped_char\x00char_sequence\x00char_range\x00_seq_char\x00manifest_repeat1\x00_include_repeat1\x00pattern_repeat1\x00char_sequence_repeat1\x00dir_pattern\x00"
