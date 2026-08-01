// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-dockerfile/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-dockerfile -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_dockerfile

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 0
const DEL_SPACE = 512
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXTERNAL_TOKEN_COUNT = 5
const FIELD_COUNT = 9
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
const MAX_HEREDOCS = 10
const PRODUCTION_ID_COUNT = 16
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 325
const SYMBOL_COUNT = 164
const TOKEN_COUNT = 87
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

type locale_t = uintptr

type wint_t = uint32

type wctype_t = uint64

type wctrans_t = uintptr

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

type scanner_state = struct {
	Fin_heredoc        uint8
	Fstripping_heredoc uint8
	Fheredoc_count     uint32
	Fheredocs          [10]uintptr
}

type TokenType = int32

const HEREDOC_MARKER = 0
const HEREDOC_LINE = 1
const HEREDOC_END = 2
const HEREDOC_NL = 3
const ERROR_SENTINEL = 4

func tree_sitter_dockerfile_external_scanner_create(tls *libc.TLS) (r uintptr) {
	var state uintptr
	_ = state
	state = libc.Xmalloc(tls, uint64(88))
	libc.Xmemset(tls, state, 0, uint64(88))
	return state
}

func tree_sitter_dockerfile_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
	var i uint32
	var state uintptr
	_, _ = i, state
	if !(payload != 0) {
		return
	}
	state = payload
	i = uint32(0)
	for {
		if !(i < uint32(MAX_HEREDOCS)) {
			break
		}
		if *(*uintptr)(unsafe.Pointer(state + 8 + uintptr(i)*8)) != 0 {
			libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(state + 8 + uintptr(i)*8)))
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	libc.Xfree(tls, state)
}

func tree_sitter_dockerfile_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	var i, len1, pos, v1 uint32
	var state uintptr
	_, _, _, _, _ = i, len1, pos, state, v1
	state = payload
	pos = uint32(0)
	v1 = pos
	pos = pos + 1
	*(*uint8)(unsafe.Pointer(buffer + uintptr(v1))) = (*scanner_state)(unsafe.Pointer(state)).Fin_heredoc
	v1 = pos
	pos = pos + 1
	*(*uint8)(unsafe.Pointer(buffer + uintptr(v1))) = (*scanner_state)(unsafe.Pointer(state)).Fstripping_heredoc
	i = uint32(0)
	for {
		if !(i < (*scanner_state)(unsafe.Pointer(state)).Fheredoc_count) {
			break
		}
		// Add the ending null byte to the length since we'll have to copy it as
		// well.
		len1 = uint32(libc.Xstrlen(tls, *(*uintptr)(unsafe.Pointer(state + 8 + uintptr(i)*8))) + uint64(1))
		// If we run out of space, just drop the heredocs that don't fit.
		// We need at least len + 1 bytes space since we'll copy len bytes below
		// and later add a null byte at the end.
		if pos+len1+uint32(1) > uint32(TREE_SITTER_SERIALIZATION_BUFFER_SIZE) {
			break
		}
		libc.Xmemcpy(tls, buffer+uintptr(pos), *(*uintptr)(unsafe.Pointer(state + 8 + uintptr(i)*8)), uint64(len1))
		pos = pos + len1
		goto _3
	_3:
		;
		i = i + 1
	}
	// Add a null byte at the end to make it easy to detect.
	v1 = pos
	pos = pos + 1
	*(*uint8)(unsafe.Pointer(buffer + uintptr(v1))) = uint8(0)
	return pos
}

func tree_sitter_dockerfile_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
	var heredoc, state uintptr
	var heredoc_count, i, i1, len1, pos, v2 uint32
	_, _, _, _, _, _, _, _ = heredoc, heredoc_count, i, i1, len1, pos, state, v2
	state = payload
	// Free all current heredocs to avoid leaking memory when we overwrite the
	// array later.
	i = uint32(0)
	for {
		if !(i < (*scanner_state)(unsafe.Pointer(state)).Fheredoc_count) {
			break
		}
		libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(state + 8 + uintptr(i)*8)))
		*(*uintptr)(unsafe.Pointer(state + 8 + uintptr(i)*8)) = libc.UintptrFromInt32(0)
		goto _1
	_1:
		;
		i = i + 1
	}
	if length == uint32(0) {
		(*scanner_state)(unsafe.Pointer(state)).Fin_heredoc = libc.BoolUint8(false1 != 0)
		(*scanner_state)(unsafe.Pointer(state)).Fstripping_heredoc = libc.BoolUint8(false1 != 0)
		(*scanner_state)(unsafe.Pointer(state)).Fheredoc_count = uint32(0)
	} else {
		pos = uint32(0)
		v2 = pos
		pos = pos + 1
		(*scanner_state)(unsafe.Pointer(state)).Fin_heredoc = uint8(libc.BoolUint8(*(*uint8)(unsafe.Pointer(buffer + uintptr(v2))) != 0))
		v2 = pos
		pos = pos + 1
		(*scanner_state)(unsafe.Pointer(state)).Fstripping_heredoc = uint8(libc.BoolUint8(*(*uint8)(unsafe.Pointer(buffer + uintptr(v2))) != 0))
		heredoc_count = uint32(0)
		i1 = uint32(0)
		for {
			if !(i1 < uint32(MAX_HEREDOCS)) {
				break
			}
			len1 = uint32(libc.Xstrlen(tls, buffer+uintptr(pos)))
			// We found the ending null byte which means that we're done.
			if len1 == uint32(0) {
				break
			}
			// Account for the ending null byte in strings (again).
			len1 = len1 + 1
			heredoc = libc.Xmalloc(tls, uint64(len1))
			libc.Xmemcpy(tls, heredoc, buffer+uintptr(pos), uint64(len1))
			*(*uintptr)(unsafe.Pointer(state + 8 + uintptr(i1)*8)) = heredoc
			heredoc_count = heredoc_count + 1
			pos = pos + len1
			goto _4
		_4:
			;
			i1 = i1 + 1
		}
		(*scanner_state)(unsafe.Pointer(state)).Fheredoc_count = heredoc_count
	}
}

func skip_whitespace(tls *libc.TLS, lexer uintptr) {
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\000') && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\n') && libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
	}
}

func scan_marker(tls *libc.TLS, state uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(512)
	defer tls.Free(512)
	var del_copy, v6 uintptr
	var del_idx, v3 uint32
	var quote int32_t
	var stripping uint8
	var v1 int32
	var v2 bool
	var _ /* delimiter at bp+0 */ [512]uint8
	_, _, _, _, _, _, _, _ = del_copy, del_idx, quote, stripping, v1, v2, v3, v6
	skip_whitespace(tls, lexer)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('<') {
		return libc.BoolUint8(false1 != 0)
	}
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('<') {
		return libc.BoolUint8(false1 != 0)
	}
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	stripping = libc.BoolUint8(false1 != 0)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('-') {
		stripping = libc.BoolUint8(true1 != 0)
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	}
	quote = 0
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('"') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\'') {
		quote = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	}
	// We start recording the actual string at position 1 since we store whether
	// it's a stripping heredoc in the first position (with either a dash or a
	// space).
	del_idx = uint32(1)
	for {
		if v2 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\000'); v2 {
			if quote != 0 {
				v1 = libc.BoolInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead != quote)
			} else {
				v1 = libc.BoolInt32(!(libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0))
			}
		}
		if !(v2 && v1 != 0) {
			break
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\\') {
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\000') {
				return libc.BoolUint8(false1 != 0)
			}
		}
		if del_idx > uint32(0) {
			v3 = del_idx
			del_idx = del_idx + 1
			(*(*[512]uint8)(unsafe.Pointer(bp)))[v3] = libc.Uint8FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)
		}
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
		// If we run out of space, stop recording the delimiter but keep
		// advancing the lexer to ensure that we at least parse the marker
		// correctly. Reserve two bytes: one for the strip indicator and
		// one for the terminating null byte.
		if del_idx >= libc.Uint32FromInt32(libc.Int32FromInt32(DEL_SPACE)-libc.Int32FromInt32(2)) {
			del_idx = uint32(0)
		}
	}
	if quote != 0 {
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != quote {
			return libc.BoolUint8(false1 != 0)
		}
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	}
	if del_idx == uint32(0) {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(HEREDOC_MARKER)
		return libc.BoolUint8(true1 != 0)
	}
	if stripping != 0 {
		v1 = int32('-')
	} else {
		v1 = int32(' ')
	}
	(*(*[512]uint8)(unsafe.Pointer(bp)))[0] = libc.Uint8FromInt32(v1)
	(*(*[512]uint8)(unsafe.Pointer(bp)))[del_idx] = uint8('\000')
	// We copy the delimiter string to the heap here since we can't store our
	// stack-allocated string in our state (which is stored on the heap).
	del_copy = libc.Xmalloc(tls, uint64(del_idx+uint32(1)))
	libc.Xmemcpy(tls, del_copy, bp, uint64(del_idx+uint32(1)))
	if (*scanner_state)(unsafe.Pointer(state)).Fheredoc_count == uint32(0) {
		(*scanner_state)(unsafe.Pointer(state)).Fheredoc_count = uint32(1)
		*(*uintptr)(unsafe.Pointer(state + 8)) = del_copy
		(*scanner_state)(unsafe.Pointer(state)).Fstripping_heredoc = stripping
	} else {
		if (*scanner_state)(unsafe.Pointer(state)).Fheredoc_count >= uint32(MAX_HEREDOCS) {
			libc.Xfree(tls, del_copy)
		} else {
			v6 = state + 4
			v3 = *(*uint32)(unsafe.Pointer(v6))
			*(*uint32)(unsafe.Pointer(v6)) = *(*uint32)(unsafe.Pointer(v6)) + 1
			*(*uintptr)(unsafe.Pointer(state + 8 + uintptr(v3)*8)) = del_copy
		}
	}
	(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(HEREDOC_MARKER)
	return libc.BoolUint8(true1 != 0)
}

func scan_content(tls *libc.TLS, state uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var delim_idx, i uint32
	_, _ = delim_idx, i
	if (*scanner_state)(unsafe.Pointer(state)).Fheredoc_count == uint32(0) {
		(*scanner_state)(unsafe.Pointer(state)).Fin_heredoc = libc.BoolUint8(false1 != 0)
		return libc.BoolUint8(false1 != 0)
	}
	(*scanner_state)(unsafe.Pointer(state)).Fin_heredoc = libc.BoolUint8(true1 != 0)
	if (*scanner_state)(unsafe.Pointer(state)).Fstripping_heredoc != 0 {
		skip_whitespace(tls, lexer)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(HEREDOC_END))) != 0 {
		delim_idx = uint32(1)
		// Look for the current heredoc delimiter.
		for libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(state + 8)) + uintptr(delim_idx)))) != int32('\000') && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\000') && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(state + 8)) + uintptr(delim_idx)))) {
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
			delim_idx = delim_idx + 1
		}
		// Check if the entire string matched.
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(state + 8)) + uintptr(delim_idx)))) == int32('\000') {
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(HEREDOC_END)
			// Shift the first heredoc off the list.
			libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(state + 8)))
			i = uint32(1)
			for {
				if !(i < (*scanner_state)(unsafe.Pointer(state)).Fheredoc_count) {
					break
				}
				*(*uintptr)(unsafe.Pointer(state + 8 + uintptr(i-uint32(1))*8)) = *(*uintptr)(unsafe.Pointer(state + 8 + uintptr(i)*8))
				goto _1
			_1:
				;
				i = i + 1
			}
			*(*uintptr)(unsafe.Pointer(state + 8 + uintptr((*scanner_state)(unsafe.Pointer(state)).Fheredoc_count-uint32(1))*8)) = libc.UintptrFromInt32(0)
			(*scanner_state)(unsafe.Pointer(state)).Fheredoc_count = (*scanner_state)(unsafe.Pointer(state)).Fheredoc_count - 1
			if (*scanner_state)(unsafe.Pointer(state)).Fheredoc_count > uint32(0) {
				(*scanner_state)(unsafe.Pointer(state)).Fstripping_heredoc = libc.BoolUint8(libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(state + 8))))) == int32('-'))
			} else {
				(*scanner_state)(unsafe.Pointer(state)).Fin_heredoc = libc.BoolUint8(false1 != 0)
			}
			return libc.BoolUint8(true1 != 0)
		}
	}
	if !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(HEREDOC_LINE))) != 0) {
		return libc.BoolUint8(false1 != 0)
	}
	(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(HEREDOC_LINE)
	for {
		switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
		case int32('\000'):
			if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
				(*scanner_state)(unsafe.Pointer(state)).Fin_heredoc = libc.BoolUint8(false1 != 0)
				return libc.BoolUint8(true1 != 0)
			}
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
		case int32('\n'):
			return libc.BoolUint8(true1 != 0)
		default:
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
		}
		goto _2
	_2:
	}
	return r
}

func tree_sitter_dockerfile_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var state uintptr
	_ = state
	state = payload
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(ERROR_SENTINEL))) != 0 {
		if (*scanner_state)(unsafe.Pointer(state)).Fin_heredoc != 0 {
			return scan_content(tls, state, lexer, valid_symbols)
		} else {
			return scan_marker(tls, state, lexer)
		}
	}
	// HEREDOC_NL only matches a linebreak if there are open heredocs. This is
	// necessary to avoid a conflict in the grammar since a normal line break
	// could either be the start of a heredoc or the end of an instruction.
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(HEREDOC_NL))) != 0 {
		if (*scanner_state)(unsafe.Pointer(state)).Fheredoc_count > uint32(0) && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') {
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(HEREDOC_NL)
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
			return libc.BoolUint8(true1 != 0)
		}
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(HEREDOC_MARKER))) != 0 {
		return scan_marker(tls, state, lexer)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(HEREDOC_LINE))) != 0 || *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(HEREDOC_END))) != 0 {
		return scan_content(tls, state, lexer, valid_symbols)
	}
	return libc.BoolUint8(false1 != 0)
}

type ts_symbol_identifiers = int32

const anon_sym_LF = 1
const aux_sym_from_instruction_token1 = 2
const aux_sym_from_instruction_token2 = 3
const aux_sym_run_instruction_token1 = 4
const aux_sym_cmd_instruction_token1 = 5
const aux_sym_label_instruction_token1 = 6
const aux_sym_expose_instruction_token1 = 7
const aux_sym_env_instruction_token1 = 8
const aux_sym_add_instruction_token1 = 9
const aux_sym_copy_instruction_token1 = 10
const aux_sym_entrypoint_instruction_token1 = 11
const aux_sym_volume_instruction_token1 = 12
const aux_sym_user_instruction_token1 = 13
const anon_sym_COLON = 14
const aux_sym__user_name_or_group_token1 = 15
const aux_sym__immediate_user_name_or_group_fragment_token1 = 16
const aux_sym_workdir_instruction_token1 = 17
const aux_sym_arg_instruction_token1 = 18
const aux_sym_arg_instruction_token2 = 19
const anon_sym_EQ = 20
const aux_sym_onbuild_instruction_token1 = 21
const aux_sym_stopsignal_instruction_token1 = 22
const aux_sym__stopsignal_value_token1 = 23
const aux_sym__stopsignal_value_token2 = 24
const aux_sym_healthcheck_instruction_token1 = 25
const anon_sym_NONE = 26
const aux_sym_shell_instruction_token1 = 27
const aux_sym_maintainer_instruction_token1 = 28
const aux_sym_maintainer_instruction_token2 = 29
const aux_sym_cross_build_instruction_token1 = 30
const aux_sym_path_token1 = 31
const aux_sym_path_token2 = 32
const aux_sym_path_token3 = 33
const aux_sym_path_with_heredoc_token1 = 34
const anon_sym_DOLLAR = 35
const anon_sym_DOLLAR2 = 36
const anon_sym_LBRACE = 37
const aux_sym__expansion_body_token1 = 38
const anon_sym_RBRACE = 39
const sym_variable = 40
const aux_sym__spaced_env_pair_token1 = 41
const aux_sym__env_key_token1 = 42
const aux_sym_expose_port_token1 = 43
const anon_sym_SLASHtcp = 44
const anon_sym_SLASHudp = 45
const aux_sym_label_pair_token1 = 46
const aux_sym_image_name_token1 = 47
const aux_sym_image_name_token2 = 48
const aux_sym_image_tag_token1 = 49
const anon_sym_AT = 50
const aux_sym_image_digest_token1 = 51
const anon_sym_DASH_DASH = 52
const aux_sym_param_token1 = 53
const aux_sym_param_token2 = 54
const anon_sym_mount = 55
const anon_sym_COMMA = 56
const aux_sym_mount_param_param_token1 = 57
const aux_sym_image_alias_token1 = 58
const aux_sym_image_alias_token2 = 59
const aux_sym_shell_fragment_token1 = 60
const aux_sym_shell_fragment_token2 = 61
const aux_sym_shell_fragment_token3 = 62
const aux_sym_shell_fragment_token4 = 63
const sym_line_continuation = 64
const sym_required_line_continuation = 65
const anon_sym_LBRACK = 66
const anon_sym_COMMA2 = 67
const anon_sym_RBRACK = 68
const anon_sym_DQUOTE = 69
const aux_sym_json_string_token1 = 70
const sym_json_escape_sequence = 71
const aux_sym_double_quoted_string_token1 = 72
const anon_sym_BSLASH = 73
const anon_sym_SQUOTE = 74
const aux_sym_single_quoted_string_token1 = 75
const aux_sym_unquoted_string_token1 = 76
const anon_sym_BSLASH2 = 77
const sym_double_quoted_escape_sequence = 78
const sym_single_quoted_escape_sequence = 79
const sym__non_newline_whitespace = 80
const sym_comment = 81
const sym_heredoc_marker = 82
const sym_heredoc_line = 83
const sym_heredoc_end = 84
const sym_heredoc_nl = 85
const sym_error_sentinel = 86
const sym_source_file = 87
const sym__instruction = 88
const sym_from_instruction = 89
const sym_run_instruction = 90
const sym_cmd_instruction = 91
const sym_label_instruction = 92
const sym_expose_instruction = 93
const sym_env_instruction = 94
const sym_add_instruction = 95
const sym_copy_instruction = 96
const sym_entrypoint_instruction = 97
const sym_volume_instruction = 98
const sym_user_instruction = 99
const sym__user_name_or_group = 100
const aux_sym__immediate_user_name_or_group = 101
const sym__immediate_user_name_or_group_fragment = 102
const sym_workdir_instruction = 103
const sym_arg_instruction = 104
const sym_onbuild_instruction = 105
const sym_stopsignal_instruction = 106
const sym__stopsignal_value = 107
const sym_healthcheck_instruction = 108
const sym_shell_instruction = 109
const sym_maintainer_instruction = 110
const sym_cross_build_instruction = 111
const sym_heredoc_block = 112
const sym_path = 113
const sym_path_with_heredoc = 114
const sym_expansion = 115
const sym__immediate_expansion = 116
const sym__imm_expansion = 117
const sym__expansion_body = 118
const sym_env_pair = 119
const sym__spaced_env_pair = 120
const sym__env_key = 121
const sym_expose_port = 122
const sym_label_pair = 123
const sym_image_spec = 124
const sym_image_name = 125
const sym_image_tag = 126
const sym_image_digest = 127
const sym_param = 128
const sym_mount_param = 129
const sym_mount_param_param = 130
const sym_image_alias = 131
const sym_shell_command = 132
const sym_shell_fragment = 133
const sym_json_string_array = 134
const sym_json_string = 135
const sym_double_quoted_string = 136
const sym_single_quoted_string = 137
const sym_unquoted_string = 138
const aux_sym_source_file_repeat1 = 139
const aux_sym_run_instruction_repeat1 = 140
const aux_sym_run_instruction_repeat2 = 141
const aux_sym_label_instruction_repeat1 = 142
const aux_sym_expose_instruction_repeat1 = 143
const aux_sym_env_instruction_repeat1 = 144
const aux_sym_add_instruction_repeat1 = 145
const aux_sym_add_instruction_repeat2 = 146
const aux_sym_volume_instruction_repeat1 = 147
const aux_sym__user_name_or_group_repeat1 = 148
const aux_sym__stopsignal_value_repeat1 = 149
const aux_sym_heredoc_block_repeat1 = 150
const aux_sym_path_repeat1 = 151
const aux_sym_image_name_repeat1 = 152
const aux_sym_image_tag_repeat1 = 153
const aux_sym_image_digest_repeat1 = 154
const aux_sym_mount_param_repeat1 = 155
const aux_sym_image_alias_repeat1 = 156
const aux_sym_shell_command_repeat1 = 157
const aux_sym_shell_fragment_repeat1 = 158
const aux_sym_json_string_array_repeat1 = 159
const aux_sym_json_string_repeat1 = 160
const aux_sym_double_quoted_string_repeat1 = 161
const aux_sym_single_quoted_string_repeat1 = 162
const aux_sym_unquoted_string_repeat1 = 163

var ts_symbol_names = [164]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 6,
	3:   __ccgo_ts + 11,
	4:   __ccgo_ts + 14,
	5:   __ccgo_ts + 18,
	6:   __ccgo_ts + 22,
	7:   __ccgo_ts + 28,
	8:   __ccgo_ts + 35,
	9:   __ccgo_ts + 39,
	10:  __ccgo_ts + 43,
	11:  __ccgo_ts + 48,
	12:  __ccgo_ts + 59,
	13:  __ccgo_ts + 66,
	14:  __ccgo_ts + 71,
	15:  __ccgo_ts + 73,
	16:  __ccgo_ts + 100,
	17:  __ccgo_ts + 146,
	18:  __ccgo_ts + 154,
	19:  __ccgo_ts + 158,
	20:  __ccgo_ts + 174,
	21:  __ccgo_ts + 176,
	22:  __ccgo_ts + 184,
	23:  __ccgo_ts + 195,
	24:  __ccgo_ts + 220,
	25:  __ccgo_ts + 245,
	26:  __ccgo_ts + 257,
	27:  __ccgo_ts + 262,
	28:  __ccgo_ts + 268,
	29:  __ccgo_ts + 279,
	30:  __ccgo_ts + 309,
	31:  __ccgo_ts + 321,
	32:  __ccgo_ts + 333,
	33:  __ccgo_ts + 345,
	34:  __ccgo_ts + 357,
	35:  __ccgo_ts + 382,
	36:  __ccgo_ts + 382,
	37:  __ccgo_ts + 384,
	38:  __ccgo_ts + 386,
	39:  __ccgo_ts + 395,
	40:  __ccgo_ts + 386,
	41:  __ccgo_ts + 397,
	42:  __ccgo_ts + 158,
	43:  __ccgo_ts + 421,
	44:  __ccgo_ts + 440,
	45:  __ccgo_ts + 445,
	46:  __ccgo_ts + 158,
	47:  __ccgo_ts + 450,
	48:  __ccgo_ts + 468,
	49:  __ccgo_ts + 486,
	50:  __ccgo_ts + 503,
	51:  __ccgo_ts + 505,
	52:  __ccgo_ts + 525,
	53:  __ccgo_ts + 528,
	54:  __ccgo_ts + 541,
	55:  __ccgo_ts + 554,
	56:  __ccgo_ts + 560,
	57:  __ccgo_ts + 562,
	58:  __ccgo_ts + 587,
	59:  __ccgo_ts + 606,
	60:  __ccgo_ts + 625,
	61:  __ccgo_ts + 647,
	62:  __ccgo_ts + 669,
	63:  __ccgo_ts + 691,
	64:  __ccgo_ts + 713,
	65:  __ccgo_ts + 713,
	66:  __ccgo_ts + 731,
	67:  __ccgo_ts + 560,
	68:  __ccgo_ts + 733,
	69:  __ccgo_ts + 735,
	70:  __ccgo_ts + 737,
	71:  __ccgo_ts + 756,
	72:  __ccgo_ts + 772,
	73:  __ccgo_ts + 800,
	74:  __ccgo_ts + 802,
	75:  __ccgo_ts + 804,
	76:  __ccgo_ts + 832,
	77:  __ccgo_ts + 855,
	78:  __ccgo_ts + 756,
	79:  __ccgo_ts + 756,
	80:  __ccgo_ts + 858,
	81:  __ccgo_ts + 882,
	82:  __ccgo_ts + 890,
	83:  __ccgo_ts + 905,
	84:  __ccgo_ts + 918,
	85:  __ccgo_ts + 930,
	86:  __ccgo_ts + 942,
	87:  __ccgo_ts + 957,
	88:  __ccgo_ts + 969,
	89:  __ccgo_ts + 982,
	90:  __ccgo_ts + 999,
	91:  __ccgo_ts + 1015,
	92:  __ccgo_ts + 1031,
	93:  __ccgo_ts + 1049,
	94:  __ccgo_ts + 1068,
	95:  __ccgo_ts + 1084,
	96:  __ccgo_ts + 1100,
	97:  __ccgo_ts + 1117,
	98:  __ccgo_ts + 1140,
	99:  __ccgo_ts + 1159,
	100: __ccgo_ts + 158,
	101: __ccgo_ts + 1176,
	102: __ccgo_ts + 1206,
	103: __ccgo_ts + 1245,
	104: __ccgo_ts + 1265,
	105: __ccgo_ts + 1281,
	106: __ccgo_ts + 1301,
	107: __ccgo_ts + 1324,
	108: __ccgo_ts + 1342,
	109: __ccgo_ts + 1366,
	110: __ccgo_ts + 1384,
	111: __ccgo_ts + 1407,
	112: __ccgo_ts + 1431,
	113: __ccgo_ts + 1445,
	114: __ccgo_ts + 1445,
	115: __ccgo_ts + 1450,
	116: __ccgo_ts + 1460,
	117: __ccgo_ts + 1450,
	118: __ccgo_ts + 1481,
	119: __ccgo_ts + 1497,
	120: __ccgo_ts + 1497,
	121: __ccgo_ts + 1506,
	122: __ccgo_ts + 1515,
	123: __ccgo_ts + 1527,
	124: __ccgo_ts + 1538,
	125: __ccgo_ts + 1549,
	126: __ccgo_ts + 1560,
	127: __ccgo_ts + 1570,
	128: __ccgo_ts + 1583,
	129: __ccgo_ts + 1589,
	130: __ccgo_ts + 1601,
	131: __ccgo_ts + 1619,
	132: __ccgo_ts + 1631,
	133: __ccgo_ts + 1645,
	134: __ccgo_ts + 1660,
	135: __ccgo_ts + 1678,
	136: __ccgo_ts + 1690,
	137: __ccgo_ts + 1711,
	138: __ccgo_ts + 158,
	139: __ccgo_ts + 1732,
	140: __ccgo_ts + 1752,
	141: __ccgo_ts + 1776,
	142: __ccgo_ts + 1800,
	143: __ccgo_ts + 1826,
	144: __ccgo_ts + 1853,
	145: __ccgo_ts + 1877,
	146: __ccgo_ts + 1901,
	147: __ccgo_ts + 1925,
	148: __ccgo_ts + 1952,
	149: __ccgo_ts + 1980,
	150: __ccgo_ts + 2006,
	151: __ccgo_ts + 2028,
	152: __ccgo_ts + 2041,
	153: __ccgo_ts + 2060,
	154: __ccgo_ts + 2078,
	155: __ccgo_ts + 2099,
	156: __ccgo_ts + 2119,
	157: __ccgo_ts + 2139,
	158: __ccgo_ts + 2161,
	159: __ccgo_ts + 2184,
	160: __ccgo_ts + 2210,
	161: __ccgo_ts + 2230,
	162: __ccgo_ts + 2259,
	163: __ccgo_ts + 2288,
}

var ts_symbol_map = [164]TSSymbol{
	1:   uint16(anon_sym_LF),
	2:   uint16(aux_sym_from_instruction_token1),
	3:   uint16(aux_sym_from_instruction_token2),
	4:   uint16(aux_sym_run_instruction_token1),
	5:   uint16(aux_sym_cmd_instruction_token1),
	6:   uint16(aux_sym_label_instruction_token1),
	7:   uint16(aux_sym_expose_instruction_token1),
	8:   uint16(aux_sym_env_instruction_token1),
	9:   uint16(aux_sym_add_instruction_token1),
	10:  uint16(aux_sym_copy_instruction_token1),
	11:  uint16(aux_sym_entrypoint_instruction_token1),
	12:  uint16(aux_sym_volume_instruction_token1),
	13:  uint16(aux_sym_user_instruction_token1),
	14:  uint16(anon_sym_COLON),
	15:  uint16(aux_sym__user_name_or_group_token1),
	16:  uint16(aux_sym__immediate_user_name_or_group_fragment_token1),
	17:  uint16(aux_sym_workdir_instruction_token1),
	18:  uint16(aux_sym_arg_instruction_token1),
	19:  uint16(sym_unquoted_string),
	20:  uint16(anon_sym_EQ),
	21:  uint16(aux_sym_onbuild_instruction_token1),
	22:  uint16(aux_sym_stopsignal_instruction_token1),
	23:  uint16(aux_sym__stopsignal_value_token1),
	24:  uint16(aux_sym__stopsignal_value_token2),
	25:  uint16(aux_sym_healthcheck_instruction_token1),
	26:  uint16(anon_sym_NONE),
	27:  uint16(aux_sym_shell_instruction_token1),
	28:  uint16(aux_sym_maintainer_instruction_token1),
	29:  uint16(aux_sym_maintainer_instruction_token2),
	30:  uint16(aux_sym_cross_build_instruction_token1),
	31:  uint16(aux_sym_path_token1),
	32:  uint16(aux_sym_path_token2),
	33:  uint16(aux_sym_path_token3),
	34:  uint16(aux_sym_path_with_heredoc_token1),
	35:  uint16(anon_sym_DOLLAR),
	36:  uint16(anon_sym_DOLLAR),
	37:  uint16(anon_sym_LBRACE),
	38:  uint16(sym_variable),
	39:  uint16(anon_sym_RBRACE),
	40:  uint16(sym_variable),
	41:  uint16(aux_sym__spaced_env_pair_token1),
	42:  uint16(sym_unquoted_string),
	43:  uint16(aux_sym_expose_port_token1),
	44:  uint16(anon_sym_SLASHtcp),
	45:  uint16(anon_sym_SLASHudp),
	46:  uint16(sym_unquoted_string),
	47:  uint16(aux_sym_image_name_token1),
	48:  uint16(aux_sym_image_name_token2),
	49:  uint16(aux_sym_image_tag_token1),
	50:  uint16(anon_sym_AT),
	51:  uint16(aux_sym_image_digest_token1),
	52:  uint16(anon_sym_DASH_DASH),
	53:  uint16(aux_sym_param_token1),
	54:  uint16(aux_sym_param_token2),
	55:  uint16(anon_sym_mount),
	56:  uint16(anon_sym_COMMA),
	57:  uint16(aux_sym_mount_param_param_token1),
	58:  uint16(aux_sym_image_alias_token1),
	59:  uint16(aux_sym_image_alias_token2),
	60:  uint16(aux_sym_shell_fragment_token1),
	61:  uint16(aux_sym_shell_fragment_token2),
	62:  uint16(aux_sym_shell_fragment_token3),
	63:  uint16(aux_sym_shell_fragment_token4),
	64:  uint16(sym_line_continuation),
	65:  uint16(sym_line_continuation),
	66:  uint16(anon_sym_LBRACK),
	67:  uint16(anon_sym_COMMA),
	68:  uint16(anon_sym_RBRACK),
	69:  uint16(anon_sym_DQUOTE),
	70:  uint16(aux_sym_json_string_token1),
	71:  uint16(sym_json_escape_sequence),
	72:  uint16(aux_sym_double_quoted_string_token1),
	73:  uint16(anon_sym_BSLASH),
	74:  uint16(anon_sym_SQUOTE),
	75:  uint16(aux_sym_single_quoted_string_token1),
	76:  uint16(aux_sym_unquoted_string_token1),
	77:  uint16(anon_sym_BSLASH2),
	78:  uint16(sym_json_escape_sequence),
	79:  uint16(sym_json_escape_sequence),
	80:  uint16(sym__non_newline_whitespace),
	81:  uint16(sym_comment),
	82:  uint16(sym_heredoc_marker),
	83:  uint16(sym_heredoc_line),
	84:  uint16(sym_heredoc_end),
	85:  uint16(sym_heredoc_nl),
	86:  uint16(sym_error_sentinel),
	87:  uint16(sym_source_file),
	88:  uint16(sym__instruction),
	89:  uint16(sym_from_instruction),
	90:  uint16(sym_run_instruction),
	91:  uint16(sym_cmd_instruction),
	92:  uint16(sym_label_instruction),
	93:  uint16(sym_expose_instruction),
	94:  uint16(sym_env_instruction),
	95:  uint16(sym_add_instruction),
	96:  uint16(sym_copy_instruction),
	97:  uint16(sym_entrypoint_instruction),
	98:  uint16(sym_volume_instruction),
	99:  uint16(sym_user_instruction),
	100: uint16(sym_unquoted_string),
	101: uint16(aux_sym__immediate_user_name_or_group),
	102: uint16(sym__immediate_user_name_or_group_fragment),
	103: uint16(sym_workdir_instruction),
	104: uint16(sym_arg_instruction),
	105: uint16(sym_onbuild_instruction),
	106: uint16(sym_stopsignal_instruction),
	107: uint16(sym__stopsignal_value),
	108: uint16(sym_healthcheck_instruction),
	109: uint16(sym_shell_instruction),
	110: uint16(sym_maintainer_instruction),
	111: uint16(sym_cross_build_instruction),
	112: uint16(sym_heredoc_block),
	113: uint16(sym_path),
	114: uint16(sym_path),
	115: uint16(sym_expansion),
	116: uint16(sym__immediate_expansion),
	117: uint16(sym_expansion),
	118: uint16(sym__expansion_body),
	119: uint16(sym_env_pair),
	120: uint16(sym_env_pair),
	121: uint16(sym__env_key),
	122: uint16(sym_expose_port),
	123: uint16(sym_label_pair),
	124: uint16(sym_image_spec),
	125: uint16(sym_image_name),
	126: uint16(sym_image_tag),
	127: uint16(sym_image_digest),
	128: uint16(sym_param),
	129: uint16(sym_mount_param),
	130: uint16(sym_mount_param_param),
	131: uint16(sym_image_alias),
	132: uint16(sym_shell_command),
	133: uint16(sym_shell_fragment),
	134: uint16(sym_json_string_array),
	135: uint16(sym_json_string),
	136: uint16(sym_double_quoted_string),
	137: uint16(sym_single_quoted_string),
	138: uint16(sym_unquoted_string),
	139: uint16(aux_sym_source_file_repeat1),
	140: uint16(aux_sym_run_instruction_repeat1),
	141: uint16(aux_sym_run_instruction_repeat2),
	142: uint16(aux_sym_label_instruction_repeat1),
	143: uint16(aux_sym_expose_instruction_repeat1),
	144: uint16(aux_sym_env_instruction_repeat1),
	145: uint16(aux_sym_add_instruction_repeat1),
	146: uint16(aux_sym_add_instruction_repeat2),
	147: uint16(aux_sym_volume_instruction_repeat1),
	148: uint16(aux_sym__user_name_or_group_repeat1),
	149: uint16(aux_sym__stopsignal_value_repeat1),
	150: uint16(aux_sym_heredoc_block_repeat1),
	151: uint16(aux_sym_path_repeat1),
	152: uint16(aux_sym_image_name_repeat1),
	153: uint16(aux_sym_image_tag_repeat1),
	154: uint16(aux_sym_image_digest_repeat1),
	155: uint16(aux_sym_mount_param_repeat1),
	156: uint16(aux_sym_image_alias_repeat1),
	157: uint16(aux_sym_shell_command_repeat1),
	158: uint16(aux_sym_shell_fragment_repeat1),
	159: uint16(aux_sym_json_string_array_repeat1),
	160: uint16(aux_sym_json_string_repeat1),
	161: uint16(aux_sym_double_quoted_string_repeat1),
	162: uint16(aux_sym_single_quoted_string_repeat1),
	163: uint16(aux_sym_unquoted_string_repeat1),
}

var ts_symbol_metadata = [164]TSSymbolMetadata{
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
	15: {},
	16: {},
	17: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	18: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	19: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	24: {},
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
	29: {},
	30: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	31: {},
	32: {},
	33: {},
	34: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	39: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	40: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	41: {},
	42: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	43: {},
	44: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	45: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	46: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	47: {},
	48: {},
	49: {},
	50: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	51: {},
	52: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	53: {},
	54: {},
	55: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	57: {},
	58: {},
	59: {},
	60: {},
	61: {},
	62: {},
	63: {},
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
	71: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	72: {},
	73: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	74: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	75: {},
	76: {},
	77: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
	101: {},
	102: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
	139: {},
	140: {},
	141: {},
	142: {},
	143: {},
	144: {},
	145: {},
	146: {},
	147: {},
	148: {},
	149: {},
	150: {},
	151: {},
	152: {},
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
	163: {},
}

type ts_field_identifiers = int32

const field_as = 1
const field_default = 2
const field_digest = 3
const field_group = 4
const field_key = 5
const field_name = 6
const field_tag = 7
const field_user = 8
const field_value = 9

var ts_field_names = [10]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 2312,
	2: __ccgo_ts + 2315,
	3: __ccgo_ts + 2323,
	4: __ccgo_ts + 2330,
	5: __ccgo_ts + 2336,
	6: __ccgo_ts + 2340,
	7: __ccgo_ts + 2345,
	8: __ccgo_ts + 2349,
	9: __ccgo_ts + 2354,
}

var ts_field_map_slices = [16]TSFieldMapSlice{
	1: {
		Flength: uint16(1),
	},
	2: {
		Findex:  uint16(1),
		Flength: uint16(2),
	},
	3: {
		Findex:  uint16(3),
		Flength: uint16(1),
	},
	4: {
		Findex:  uint16(4),
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(5),
		Flength: uint16(2),
	},
	6: {
		Findex:  uint16(7),
		Flength: uint16(2),
	},
	7: {
		Findex:  uint16(9),
		Flength: uint16(1),
	},
	8: {
		Findex:  uint16(10),
		Flength: uint16(3),
	},
	9: {
		Findex:  uint16(13),
		Flength: uint16(2),
	},
	10: {
		Findex:  uint16(15),
		Flength: uint16(2),
	},
	11: {
		Findex:  uint16(17),
		Flength: uint16(2),
	},
	12: {
		Findex:  uint16(19),
		Flength: uint16(2),
	},
	13: {
		Findex:  uint16(21),
		Flength: uint16(2),
	},
	14: {
		Findex:  uint16(23),
		Flength: uint16(1),
	},
	15: {
		Findex:  uint16(24),
		Flength: uint16(3),
	},
}

var ts_field_map_entries = [27]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_name),
	},
	1: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	2: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	3: {
		Ffield_id:    uint16(field_user),
		Fchild_index: uint8(1),
	},
	4: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	5: {
		Ffield_id: uint16(field_name),
	},
	6: {
		Ffield_id:    uint16(field_tag),
		Fchild_index: uint8(1),
	},
	7: {
		Ffield_id:    uint16(field_digest),
		Fchild_index: uint8(1),
	},
	8: {
		Ffield_id: uint16(field_name),
	},
	9: {
		Ffield_id:    uint16(field_as),
		Fchild_index: uint8(3),
	},
	10: {
		Ffield_id:    uint16(field_digest),
		Fchild_index: uint8(2),
	},
	11: {
		Ffield_id: uint16(field_name),
	},
	12: {
		Ffield_id:    uint16(field_tag),
		Fchild_index: uint8(1),
	},
	13: {
		Ffield_id: uint16(field_key),
	},
	14: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	15: {
		Ffield_id: uint16(field_name),
	},
	16: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	17: {
		Ffield_id:    uint16(field_group),
		Fchild_index: uint8(3),
	},
	18: {
		Ffield_id:    uint16(field_user),
		Fchild_index: uint8(1),
	},
	19: {
		Ffield_id:    uint16(field_default),
		Fchild_index: uint8(3),
	},
	20: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	21: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	22: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(3),
	},
	23: {
		Ffield_id:    uint16(field_as),
		Fchild_index: uint8(4),
	},
	24: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	25: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(3),
	},
	26: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(4),
	},
}

var ts_alias_sequences = [16][5]TSSymbol{
	0: {},
	11: {
		3: uint16(sym_unquoted_string),
	},
}

var ts_non_terminal_alias_map = [5]uint16_t{
	0: uint16(aux_sym__immediate_user_name_or_group),
	1: uint16(2),
	2: uint16(aux_sym__immediate_user_name_or_group),
	3: uint16(sym_unquoted_string),
}

var ts_primary_state_ids = [325]TSStateId{
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
	33:  uint16(30),
	34:  uint16(34),
	35:  uint16(35),
	36:  uint16(36),
	37:  uint16(37),
	38:  uint16(38),
	39:  uint16(39),
	40:  uint16(40),
	41:  uint16(41),
	42:  uint16(42),
	43:  uint16(22),
	44:  uint16(44),
	45:  uint16(23),
	46:  uint16(46),
	47:  uint16(18),
	48:  uint16(20),
	49:  uint16(49),
	50:  uint16(36),
	51:  uint16(30),
	52:  uint16(36),
	53:  uint16(53),
	54:  uint16(54),
	55:  uint16(55),
	56:  uint16(56),
	57:  uint16(57),
	58:  uint16(58),
	59:  uint16(59),
	60:  uint16(60),
	61:  uint16(61),
	62:  uint16(32),
	63:  uint16(58),
	64:  uint16(64),
	65:  uint16(65),
	66:  uint16(66),
	67:  uint16(49),
	68:  uint16(68),
	69:  uint16(69),
	70:  uint16(70),
	71:  uint16(71),
	72:  uint16(72),
	73:  uint16(73),
	74:  uint16(74),
	75:  uint16(72),
	76:  uint16(60),
	77:  uint16(77),
	78:  uint16(73),
	79:  uint16(79),
	80:  uint16(80),
	81:  uint16(70),
	82:  uint16(49),
	83:  uint16(83),
	84:  uint16(44),
	85:  uint16(85),
	86:  uint16(49),
	87:  uint16(87),
	88:  uint16(88),
	89:  uint16(89),
	90:  uint16(90),
	91:  uint16(91),
	92:  uint16(92),
	93:  uint16(46),
	94:  uint16(94),
	95:  uint16(72),
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
	109: uint16(72),
	110: uint16(110),
	111: uint16(68),
	112: uint16(112),
	113: uint16(80),
	114: uint16(73),
	115: uint16(73),
	116: uint16(116),
	117: uint16(72),
	118: uint16(73),
	119: uint16(72),
	120: uint16(73),
	121: uint16(121),
	122: uint16(112),
	123: uint16(106),
	124: uint16(124),
	125: uint16(112),
	126: uint16(106),
	127: uint16(127),
	128: uint16(128),
	129: uint16(129),
	130: uint16(130),
	131: uint16(131),
	132: uint16(132),
	133: uint16(80),
	134: uint16(72),
	135: uint16(135),
	136: uint16(73),
	137: uint16(80),
	138: uint16(72),
	139: uint16(73),
	140: uint16(140),
	141: uint16(141),
	142: uint16(142),
	143: uint16(143),
	144: uint16(144),
	145: uint16(145),
	146: uint16(146),
	147: uint16(147),
	148: uint16(148),
	149: uint16(73),
	150: uint16(150),
	151: uint16(151),
	152: uint16(152),
	153: uint16(64),
	154: uint16(154),
	155: uint16(155),
	156: uint16(156),
	157: uint16(157),
	158: uint16(158),
	159: uint16(159),
	160: uint16(160),
	161: uint16(72),
	162: uint16(162),
	163: uint16(163),
	164: uint16(164),
	165: uint16(165),
	166: uint16(166),
	167: uint16(150),
	168: uint16(168),
	169: uint16(169),
	170: uint16(80),
	171: uint16(171),
	172: uint16(171),
	173: uint16(152),
	174: uint16(129),
	175: uint16(73),
	176: uint16(80),
	177: uint16(72),
	178: uint16(73),
	179: uint16(179),
	180: uint16(180),
	181: uint16(80),
	182: uint16(72),
	183: uint16(73),
	184: uint16(80),
	185: uint16(72),
	186: uint16(171),
	187: uint16(187),
	188: uint16(80),
	189: uint16(189),
	190: uint16(190),
	191: uint16(191),
	192: uint16(73),
	193: uint16(72),
	194: uint16(162),
	195: uint16(191),
	196: uint16(189),
	197: uint16(162),
	198: uint16(198),
	199: uint16(179),
	200: uint16(171),
	201: uint16(162),
	202: uint16(202),
	203: uint16(171),
	204: uint16(171),
	205: uint16(162),
	206: uint16(171),
	207: uint16(162),
	208: uint16(171),
	209: uint16(162),
	210: uint16(171),
	211: uint16(162),
	212: uint16(171),
	213: uint16(162),
	214: uint16(171),
	215: uint16(171),
	216: uint16(171),
	217: uint16(73),
	218: uint16(218),
	219: uint16(219),
	220: uint16(220),
	221: uint16(68),
	222: uint16(148),
	223: uint16(151),
	224: uint16(224),
	225: uint16(159),
	226: uint16(160),
	227: uint16(227),
	228: uint16(228),
	229: uint16(229),
	230: uint16(230),
	231: uint16(231),
	232: uint16(68),
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
	243: uint16(187),
	244: uint16(244),
	245: uint16(245),
	246: uint16(246),
	247: uint16(247),
	248: uint16(248),
	249: uint16(249),
	250: uint16(250),
	251: uint16(248),
	252: uint16(159),
	253: uint16(253),
	254: uint16(160),
	255: uint16(255),
	256: uint16(256),
	257: uint16(257),
	258: uint16(258),
	259: uint16(259),
	260: uint16(260),
	261: uint16(261),
	262: uint16(262),
	263: uint16(263),
	264: uint16(264),
	265: uint16(265),
	266: uint16(266),
	267: uint16(267),
	268: uint16(268),
	269: uint16(236),
	270: uint16(270),
	271: uint16(271),
	272: uint16(148),
	273: uint16(273),
	274: uint16(274),
	275: uint16(275),
	276: uint16(267),
	277: uint16(268),
	278: uint16(278),
	279: uint16(151),
	280: uint16(267),
	281: uint16(268),
	282: uint16(282),
	283: uint16(267),
	284: uint16(267),
	285: uint16(285),
	286: uint16(260),
	287: uint16(267),
	288: uint16(218),
	289: uint16(231),
	290: uint16(267),
	291: uint16(291),
	292: uint16(292),
	293: uint16(267),
	294: uint16(168),
	295: uint16(295),
	296: uint16(267),
	297: uint16(297),
	298: uint16(267),
	299: uint16(299),
	300: uint16(267),
	301: uint16(301),
	302: uint16(267),
	303: uint16(303),
	304: uint16(267),
	305: uint16(267),
	306: uint16(248),
	307: uint16(274),
	308: uint16(248),
	309: uint16(274),
	310: uint16(248),
	311: uint16(274),
	312: uint16(248),
	313: uint16(313),
	314: uint16(248),
	315: uint16(248),
	316: uint16(248),
	317: uint16(248),
	318: uint16(248),
	319: uint16(248),
	320: uint16(248),
	321: uint16(248),
	322: uint16(303),
	323: uint16(303),
	324: uint16(268),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2, i3, i4, i5, i6 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _, _, _, _ = eof, i, i1, i2, i3, i4, i5, i6, lookahead, result, skip
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
			state = uint16(167)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(2)) {
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
			state = uint16(164)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(185)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(184)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\t') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32(' ') {
			state = uint16(289)
			goto next_state
		}
		return result
	case int32(2):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
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
		return result
	case int32(3):
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(264)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32(',') && lookahead != int32('-') && lookahead != int32('=') {
			state = uint16(263)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(285)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(1)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(286)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(285)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(1)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(287)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(181)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(190)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(185)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('A') || lookahead == int32('a') {
			state = uint16(240)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32(':') || int32('B') <= lookahead && lookahead <= int32('Z') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(13):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
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
			state = uint16(14)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(212)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('A') || lookahead == int32('a') {
			state = uint16(142)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(14)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(261)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(24)
			goto next_state
		}
		if lookahead == int32(',') || lookahead == int32('-') || lookahead == int32('=') {
			state = uint16(259)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('[') && lookahead != int32('\\') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(292)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('A') || lookahead == int32('a') {
			state = uint16(142)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(19)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(20):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
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
			state = uint16(17)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(237)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(239)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(235)
			goto next_state
		}
		if lookahead == int32('A') || lookahead == int32('a') {
			state = uint16(236)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(209)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(208)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(292)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(209)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(208)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('\n') {
			state = uint16(266)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(264)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32(',') && lookahead != int32('-') && lookahead != int32('=') {
			state = uint16(263)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(276)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(279)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(275)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(283)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(280)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(282)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(284)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('"') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('"') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('"') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(278)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('"') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(285)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(1)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(27)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('"') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(271)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(2)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(272)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('#') {
			state = uint16(205)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(212)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(204)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(32)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('#') {
			state = uint16(205)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(212)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(267)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(204)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(33)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('-') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('#') {
			state = uint16(205)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(212)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(204)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(34)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('-') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(48)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(185)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(48)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(212)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(244)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(38)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(247)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(212)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(38)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(212)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(212)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(40)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(182)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(278)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(252)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(260)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(261)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(259)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(267)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(4)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(43)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(260)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(261)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(267)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32(',') || lookahead == int32('=') {
			state = uint16(259)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(43)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(261)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(267)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32(',') || lookahead == int32('-') || lookahead == int32('=') {
			state = uint16(259)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(44)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(190)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(214)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(48)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(292)
			goto next_state
		}
		if int32('\n') <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(48)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(50)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('#') {
			state = uint16(293)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(3)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(51)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('#') {
			state = uint16(237)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(235)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(48)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('@') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('#') {
			state = uint16(209)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(208)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(292)
			goto next_state
		}
		if int32('\n') <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(48)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('#') {
			state = uint16(230)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(212)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(229)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(54)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32(':') && lookahead != int32('@') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('#') {
			state = uint16(254)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(253)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(48)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32(',') && lookahead != int32('=') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('#') {
			state = uint16(215)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(216)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(217)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('}') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('#') {
			state = uint16(249)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(248)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(48)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('-') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('E') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('N') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('O') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('_') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('c') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('d') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('p') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('p') {
			state = uint16(226)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('t') {
			state = uint16(63)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('A') || lookahead == int32('a') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('A') || lookahead == int32('a') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('A') || lookahead == int32('a') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('A') || lookahead == int32('a') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('A') || lookahead == int32('a') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('B') || lookahead == int32('b') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('B') || lookahead == int32('b') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('B') || lookahead == int32('b') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('C') || lookahead == int32('c') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('C') || lookahead == int32('c') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('D') || lookahead == int32('d') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('R') || lookahead == int32('r') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead == int32('D') || lookahead == int32('d') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead == int32('D') || lookahead == int32('d') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead == int32('D') || lookahead == int32('d') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead == int32('D') || lookahead == int32('d') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('D') || lookahead == int32('d') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('G') || lookahead == int32('g') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('G') || lookahead == int32('g') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead == int32('H') || lookahead == int32('h') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead == int32('H') || lookahead == int32('h') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('T') || lookahead == int32('t') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead == int32('H') || lookahead == int32('h') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead == int32('I') || lookahead == int32('i') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead == int32('I') || lookahead == int32('i') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead == int32('I') || lookahead == int32('i') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead == int32('I') || lookahead == int32('i') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('I') || lookahead == int32('i') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead == int32('I') || lookahead == int32('i') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead == int32('I') || lookahead == int32('i') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead == int32('K') || lookahead == int32('k') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead == int32('K') || lookahead == int32('k') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead == int32('L') || lookahead == int32('l') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead == int32('L') || lookahead == int32('l') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead == int32('L') || lookahead == int32('l') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('L') || lookahead == int32('l') {
			state = uint16(192)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead == int32('L') || lookahead == int32('l') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead == int32('L') || lookahead == int32('l') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(112):
		if lookahead == int32('L') || lookahead == int32('l') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead == int32('L') || lookahead == int32('l') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('M') || lookahead == int32('m') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('M') || lookahead == int32('m') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(131)
			goto next_state
		}
		if lookahead == int32('R') || lookahead == int32('r') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead == int32('M') || lookahead == int32('m') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead == int32('N') || lookahead == int32('n') {
			state = uint16(149)
			goto next_state
		}
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead == int32('N') || lookahead == int32('n') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('N') || lookahead == int32('n') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('N') || lookahead == int32('n') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('N') || lookahead == int32('n') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('N') || lookahead == int32('n') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead == int32('N') || lookahead == int32('n') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(124):
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(129):
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(130):
		if lookahead == int32('O') || lookahead == int32('o') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(131):
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(132):
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(133):
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(134):
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead == int32('R') || lookahead == int32('r') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(136):
		if lookahead == int32('R') || lookahead == int32('r') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(137):
		if lookahead == int32('R') || lookahead == int32('r') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(138):
		if lookahead == int32('R') || lookahead == int32('r') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(139):
		if lookahead == int32('R') || lookahead == int32('r') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(140):
		if lookahead == int32('R') || lookahead == int32('r') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead == int32('S') || lookahead == int32('s') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(142):
		if lookahead == int32('S') || lookahead == int32('s') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(143):
		if lookahead == int32('S') || lookahead == int32('s') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(144):
		if lookahead == int32('S') || lookahead == int32('s') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(145):
		if lookahead == int32('S') || lookahead == int32('s') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(146):
		if lookahead == int32('S') || lookahead == int32('s') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(147):
		if lookahead == int32('T') || lookahead == int32('t') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(148):
		if lookahead == int32('T') || lookahead == int32('t') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(149):
		if lookahead == int32('T') || lookahead == int32('t') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('V') || lookahead == int32('v') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(150):
		if lookahead == int32('T') || lookahead == int32('t') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(151):
		if lookahead == int32('U') || lookahead == int32('u') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(152):
		if lookahead == int32('U') || lookahead == int32('u') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead == int32('U') || lookahead == int32('u') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(154):
		if lookahead == int32('U') || lookahead == int32('u') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(155):
		if lookahead == int32('Y') || lookahead == int32('y') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(156):
		if lookahead == int32('Y') || lookahead == int32('y') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(157):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(158):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(274)
			goto next_state
		}
		return result
	case int32(159):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(160):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(161):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(162):
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('$') && lookahead != int32('-') && lookahead != int32('<') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(163):
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(164):
		if eof != 0 {
			state = uint16(167)
			goto next_state
		}
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
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
			state = uint16(164)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(165):
		if eof != 0 {
			state = uint16(167)
			goto next_state
		}
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(148)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token5[i5]) == lookahead {
				state = map_token5[i5+uint32(1)]
				goto next_state
			}
			goto _6
		_6:
			;
			i5 = i5 + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(166)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(166):
		if eof != 0 {
			state = uint16(167)
			goto next_state
		}
		i6 = uint32(0)
		for {
			if !(uint64(i6) < libc.Uint64FromInt64(140)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token6[i6]) == lookahead {
				state = map_token6[i6+uint32(1)]
				goto next_state
			}
			goto _7
		_7:
			;
			i6 = i6 + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(166)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(167):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(168):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LF)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(169):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_from_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(170):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_from_instruction_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(171):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_run_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(172):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_cmd_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(173):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_label_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(174):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_expose_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(175):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_env_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(176):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_add_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(177):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_copy_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(178):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_entrypoint_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(179):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_volume_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(180):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_user_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(181):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(182):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__user_name_or_group_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(182)
			goto next_state
		}
		return result
	case int32(183):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__user_name_or_group_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(184):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__immediate_user_name_or_group_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(186)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(185):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__immediate_user_name_or_group_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(186):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__immediate_user_name_or_group_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(187):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_workdir_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(188):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_arg_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(189):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_arg_instruction_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(190):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(191):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_onbuild_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(192):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_stopsignal_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(193):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__stopsignal_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(194):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__stopsignal_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(195):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_healthcheck_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(196):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NONE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(197):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_shell_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(198):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_maintainer_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(199):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_maintainer_instruction_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(199)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(200):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_maintainer_instruction_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(201)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(199)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(200)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(201):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_maintainer_instruction_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(202):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_cross_build_instruction_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(203):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_path_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(204):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_path_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(205):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_path_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(293)
			goto next_state
		}
		return result
	case int32(206):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_path_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(207):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_path_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('<') && lookahead != int32('\\') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(208):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_path_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32('$') {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(209):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_path_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32('$') {
			state = uint16(293)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(210):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_path_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('$') {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(211):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_path_with_heredoc_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(212):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(213):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(214):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(215):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__expansion_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(218)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(293)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(216):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__expansion_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(218)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(216)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('}') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(217):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__expansion_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(215)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(216)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(217)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('}') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(218):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__expansion_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('}') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(219):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(220):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_variable)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(221):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__spaced_env_pair_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(222):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__env_key_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(223):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_expose_port_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(157)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(224):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_expose_port_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(225):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASHtcp)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(226):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASHudp)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(227):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_label_pair_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(228):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_name_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(229):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_name_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(230):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_name_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(293)
			goto next_state
		}
		return result
	case int32(231):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_name_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32('$') && lookahead != int32(':') && lookahead != int32('@') {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(232):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_name_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('S') || lookahead == int32('s') {
			state = uint16(234)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('$') && lookahead != int32(':') && lookahead != int32('@') {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(233):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_name_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32('$') || lookahead == int32(':') || lookahead == int32('@') {
			state = uint16(293)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(233)
			goto next_state
		}
		return result
	case int32(234):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_name_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('$') && lookahead != int32(':') && lookahead != int32('@') {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(235):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_tag_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32('$') && lookahead != int32('@') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(236):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_tag_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('S') || lookahead == int32('s') {
			state = uint16(238)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('$') && lookahead != int32('@') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(237):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_tag_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32('$') || lookahead == int32('@') {
			state = uint16(293)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(238):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_tag_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('$') && lookahead != int32('@') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(239):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(240):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_digest_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('S') || lookahead == int32('s') {
			state = uint16(241)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(241):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_digest_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(242):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(243):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_param_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(245)
			goto next_state
		}
		if lookahead == int32('-') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(244):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_param_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(246)
			goto next_state
		}
		if lookahead == int32('-') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(245):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_param_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(251)
			goto next_state
		}
		if lookahead == int32('-') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(246):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_param_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(243)
			goto next_state
		}
		if lookahead == int32('-') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(247):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_param_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(248):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_param_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(249):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_param_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(293)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(250):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_param_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(251):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mount)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(252):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(253):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_mount_param_param_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(',') && lookahead != int32('=') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(254):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_mount_param_param_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(',') || lookahead == int32('=') {
			state = uint16(293)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(254)
			goto next_state
		}
		return result
	case int32(255):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_mount_param_param_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32(',') && lookahead != int32('=') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(256):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_alias_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(257):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_image_alias_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(258):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_shell_fragment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(259):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_shell_fragment_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(260):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_shell_fragment_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(261):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_shell_fragment_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') || lookahead == int32('\\') {
			state = uint16(206)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(262):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_shell_fragment_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('<') && lookahead != int32('\\') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(263):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_shell_fragment_token4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(264):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_shell_fragment_token4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(265):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_line_continuation)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(266):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_required_line_continuation)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(267):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(268):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(269):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(270):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(271):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_json_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(273)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\\') {
			state = uint16(293)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(272):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_json_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(271)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(272)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('#') && lookahead != int32('\\') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(273):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_json_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(274):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_json_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(275):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_double_quoted_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(276)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(275)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && (lookahead < int32('"') || int32('$') < lookahead) && lookahead != int32('\\') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(276):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_double_quoted_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('"') || lookahead == int32('$') || lookahead == int32('\\') {
			state = uint16(293)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(277):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_double_quoted_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('"') && lookahead != int32('$') && lookahead != int32('\\') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(278):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(279):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\\') {
			state = uint16(290)
			goto next_state
		}
		return result
	case int32(280):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('\'') || lookahead == int32('\\') {
			state = uint16(291)
			goto next_state
		}
		return result
	case int32(281):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(282):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_single_quoted_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(283)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(282)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(284)
			goto next_state
		}
		return result
	case int32(283):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_single_quoted_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\'') || lookahead == int32('\\') {
			state = uint16(293)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(283)
			goto next_state
		}
		return result
	case int32(284):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_single_quoted_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(284)
			goto next_state
		}
		return result
	case int32(285):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_unquoted_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32('"') || lookahead == int32('$') || lookahead == int32('\'') || lookahead == int32('\\') {
			state = uint16(293)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(285)
			goto next_state
		}
		return result
	case int32(286):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_unquoted_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(286)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('$') && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(287):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_unquoted_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(287)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('$') && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(288):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_unquoted_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') && lookahead != int32('"') && lookahead != int32('$') && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(289):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(290):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_double_quoted_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(291):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_single_quoted_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(292):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__non_newline_whitespace)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(293):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(293)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [32]uint16_t{
	0:  uint16('"'),
	1:  uint16(270),
	2:  uint16('#'),
	3:  uint16(205),
	4:  uint16('$'),
	5:  uint16(213),
	6:  uint16('\''),
	7:  uint16(281),
	8:  uint16(','),
	9:  uint16(252),
	10: uint16('-'),
	11: uint16(260),
	12: uint16(':'),
	13: uint16(181),
	14: uint16('<'),
	15: uint16(228),
	16: uint16('='),
	17: uint16(190),
	18: uint16('@'),
	19: uint16(239),
	20: uint16('['),
	21: uint16(267),
	22: uint16('\\'),
	23: uint16(278),
	24: uint16(']'),
	25: uint16(269),
	26: uint16('_'),
	27: uint16(220),
	28: uint16('{'),
	29: uint16(214),
	30: uint16('}'),
	31: uint16(219),
}

var map_token1 = [24]uint16_t{
	0:  uint16('\n'),
	1:  uint16(265),
	2:  uint16('u'),
	3:  uint16(161),
	4:  uint16('\t'),
	5:  uint16(3),
	6:  uint16(' '),
	7:  uint16(3),
	8:  uint16('"'),
	9:  uint16(274),
	10: uint16('/'),
	11: uint16(274),
	12: uint16('\\'),
	13: uint16(274),
	14: uint16('b'),
	15: uint16(274),
	16: uint16('f'),
	17: uint16(274),
	18: uint16('n'),
	19: uint16(274),
	20: uint16('r'),
	21: uint16(274),
	22: uint16('t'),
	23: uint16(274),
}

var map_token2 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(168),
	2:  uint16('#'),
	3:  uint16(293),
	4:  uint16('$'),
	5:  uint16(212),
	6:  uint16('/'),
	7:  uint16(67),
	8:  uint16(':'),
	9:  uint16(181),
	10: uint16('@'),
	11: uint16(239),
	12: uint16('\\'),
	13: uint16(3),
	14: uint16('A'),
	15: uint16(142),
	16: uint16('a'),
	17: uint16(142),
}

var map_token3 = [16]uint16_t{
	0:  uint16('\n'),
	1:  uint16(168),
	2:  uint16('#'),
	3:  uint16(233),
	4:  uint16('$'),
	5:  uint16(213),
	6:  uint16(':'),
	7:  uint16(181),
	8:  uint16('@'),
	9:  uint16(239),
	10: uint16('\\'),
	11: uint16(231),
	12: uint16('A'),
	13: uint16(232),
	14: uint16('a'),
	15: uint16(232),
}

var map_token4 = [26]uint16_t{
	0:  uint16('"'),
	1:  uint16(270),
	2:  uint16('#'),
	3:  uint16(205),
	4:  uint16('$'),
	5:  uint16(212),
	6:  uint16('\''),
	7:  uint16(281),
	8:  uint16(','),
	9:  uint16(268),
	10: uint16('-'),
	11: uint16(260),
	12: uint16('<'),
	13: uint16(228),
	14: uint16('='),
	15: uint16(203),
	16: uint16('['),
	17: uint16(267),
	18: uint16('\\'),
	19: uint16(278),
	20: uint16(']'),
	21: uint16(269),
	22: uint16(':'),
	23: uint16(203),
	24: uint16('@'),
	25: uint16(203),
}

var map_token5 = [74]uint16_t{
	0:  uint16('"'),
	1:  uint16(270),
	2:  uint16('#'),
	3:  uint16(293),
	4:  uint16('$'),
	5:  uint16(212),
	6:  uint16(','),
	7:  uint16(268),
	8:  uint16('-'),
	9:  uint16(58),
	10: uint16('='),
	11: uint16(190),
	12: uint16('N'),
	13: uint16(61),
	14: uint16('['),
	15: uint16(267),
	16: uint16('\\'),
	17: uint16(3),
	18: uint16(']'),
	19: uint16(269),
	20: uint16('}'),
	21: uint16(219),
	22: uint16('A'),
	23: uint16(78),
	24: uint16('a'),
	25: uint16(78),
	26: uint16('C'),
	27: uint16(115),
	28: uint16('c'),
	29: uint16(115),
	30: uint16('E'),
	31: uint16(117),
	32: uint16('e'),
	33: uint16(117),
	34: uint16('F'),
	35: uint16(140),
	36: uint16('f'),
	37: uint16(140),
	38: uint16('H'),
	39: uint16(87),
	40: uint16('h'),
	41: uint16(87),
	42: uint16('L'),
	43: uint16(68),
	44: uint16('l'),
	45: uint16(68),
	46: uint16('M'),
	47: uint16(69),
	48: uint16('m'),
	49: uint16(69),
	50: uint16('O'),
	51: uint16(119),
	52: uint16('o'),
	53: uint16(119),
	54: uint16('R'),
	55: uint16(151),
	56: uint16('r'),
	57: uint16(151),
	58: uint16('S'),
	59: uint16(95),
	60: uint16('s'),
	61: uint16(95),
	62: uint16('U'),
	63: uint16(143),
	64: uint16('u'),
	65: uint16(143),
	66: uint16('V'),
	67: uint16(124),
	68: uint16('v'),
	69: uint16(124),
	70: uint16('W'),
	71: uint16(126),
	72: uint16('w'),
	73: uint16(126),
}

var map_token6 = [70]uint16_t{
	0:  uint16('"'),
	1:  uint16(270),
	2:  uint16('#'),
	3:  uint16(293),
	4:  uint16('$'),
	5:  uint16(212),
	6:  uint16(','),
	7:  uint16(268),
	8:  uint16('-'),
	9:  uint16(58),
	10: uint16('N'),
	11: uint16(61),
	12: uint16('['),
	13: uint16(267),
	14: uint16('\\'),
	15: uint16(3),
	16: uint16(']'),
	17: uint16(269),
	18: uint16('A'),
	19: uint16(78),
	20: uint16('a'),
	21: uint16(78),
	22: uint16('C'),
	23: uint16(115),
	24: uint16('c'),
	25: uint16(115),
	26: uint16('E'),
	27: uint16(117),
	28: uint16('e'),
	29: uint16(117),
	30: uint16('F'),
	31: uint16(140),
	32: uint16('f'),
	33: uint16(140),
	34: uint16('H'),
	35: uint16(87),
	36: uint16('h'),
	37: uint16(87),
	38: uint16('L'),
	39: uint16(68),
	40: uint16('l'),
	41: uint16(68),
	42: uint16('M'),
	43: uint16(69),
	44: uint16('m'),
	45: uint16(69),
	46: uint16('O'),
	47: uint16(119),
	48: uint16('o'),
	49: uint16(119),
	50: uint16('R'),
	51: uint16(151),
	52: uint16('r'),
	53: uint16(151),
	54: uint16('S'),
	55: uint16(95),
	56: uint16('s'),
	57: uint16(95),
	58: uint16('U'),
	59: uint16(143),
	60: uint16('u'),
	61: uint16(143),
	62: uint16('V'),
	63: uint16(124),
	64: uint16('v'),
	65: uint16(124),
	66: uint16('W'),
	67: uint16(126),
	68: uint16('w'),
	69: uint16(126),
}

var ts_lex_modes = [325]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state: uint16(165),
	},
	2: {
		Flex_state: uint16(165),
	},
	3: {
		Flex_state: uint16(165),
	},
	4: {
		Flex_state: uint16(165),
	},
	5: {
		Flex_state: uint16(165),
	},
	6: {
		Flex_state:          uint16(43),
		Fexternal_lex_state: uint16(2),
	},
	7: {
		Flex_state:          uint16(43),
		Fexternal_lex_state: uint16(2),
	},
	8: {
		Flex_state: uint16(7),
	},
	9: {
		Flex_state: uint16(30),
	},
	10: {
		Flex_state: uint16(30),
	},
	11: {
		Flex_state: uint16(30),
	},
	12: {
		Flex_state:          uint16(44),
		Fexternal_lex_state: uint16(2),
	},
	13: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(2),
	},
	14: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(2),
	},
	15: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(2),
	},
	16: {
		Flex_state:          uint16(44),
		Fexternal_lex_state: uint16(2),
	},
	17: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(2),
	},
	18: {
		Flex_state: uint16(8),
	},
	19: {
		Flex_state:          uint16(43),
		Fexternal_lex_state: uint16(2),
	},
	20: {
		Flex_state: uint16(8),
	},
	21: {
		Flex_state: uint16(20),
	},
	22: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(3),
	},
	23: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(3),
	},
	24: {
		Flex_state: uint16(20),
	},
	25: {
		Flex_state: uint16(20),
	},
	26: {
		Flex_state:          uint16(42),
		Fexternal_lex_state: uint16(2),
	},
	27: {
		Flex_state:          uint16(42),
		Fexternal_lex_state: uint16(2),
	},
	28: {
		Flex_state:          uint16(42),
		Fexternal_lex_state: uint16(2),
	},
	29: {
		Flex_state: uint16(21),
	},
	30: {
		Flex_state: uint16(25),
	},
	31: {
		Flex_state: uint16(9),
	},
	32: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(3),
	},
	33: {
		Flex_state: uint16(25),
	},
	34: {
		Flex_state: uint16(6),
	},
	35: {
		Flex_state: uint16(9),
	},
	36: {
		Flex_state: uint16(25),
	},
	37: {
		Flex_state: uint16(25),
	},
	38: {
		Flex_state: uint16(9),
	},
	39: {
		Flex_state: uint16(21),
	},
	40: {
		Flex_state: uint16(6),
	},
	41: {
		Flex_state:          uint16(42),
		Fexternal_lex_state: uint16(2),
	},
	42: {
		Flex_state:          uint16(42),
		Fexternal_lex_state: uint16(2),
	},
	43: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(2),
	},
	44: {
		Flex_state:          uint16(22),
		Fexternal_lex_state: uint16(4),
	},
	45: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(2),
	},
	46: {
		Flex_state:          uint16(22),
		Fexternal_lex_state: uint16(4),
	},
	47: {
		Flex_state: uint16(7),
	},
	48: {
		Flex_state: uint16(7),
	},
	49: {
		Flex_state:          uint16(22),
		Fexternal_lex_state: uint16(4),
	},
	50: {
		Flex_state: uint16(25),
	},
	51: {
		Flex_state: uint16(25),
	},
	52: {
		Flex_state: uint16(25),
	},
	53: {
		Flex_state: uint16(9),
	},
	54: {
		Flex_state: uint16(33),
	},
	55: {
		Flex_state: uint16(10),
	},
	56: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(2),
	},
	57: {
		Flex_state: uint16(9),
	},
	58: {
		Flex_state:          uint16(44),
		Fexternal_lex_state: uint16(2),
	},
	59: {
		Flex_state: uint16(10),
	},
	60: {
		Flex_state: uint16(22),
	},
	61: {
		Flex_state: uint16(28),
	},
	62: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(2),
	},
	63: {
		Flex_state:          uint16(44),
		Fexternal_lex_state: uint16(2),
	},
	64: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(2),
	},
	65: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(2),
	},
	66: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(2),
	},
	67: {
		Flex_state: uint16(22),
	},
	68: {
		Flex_state:          uint16(43),
		Fexternal_lex_state: uint16(2),
	},
	69: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(2),
	},
	70: {
		Flex_state: uint16(22),
	},
	71: {
		Flex_state: uint16(54),
	},
	72: {
		Flex_state: uint16(8),
	},
	73: {
		Flex_state: uint16(8),
	},
	74: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(2),
	},
	75: {
		Flex_state: uint16(20),
	},
	76: {
		Flex_state: uint16(23),
	},
	77: {
		Flex_state: uint16(11),
	},
	78: {
		Flex_state: uint16(20),
	},
	79: {
		Flex_state: uint16(12),
	},
	80: {
		Flex_state: uint16(20),
	},
	81: {
		Flex_state: uint16(23),
	},
	82: {
		Flex_state: uint16(53),
	},
	83: {
		Flex_state: uint16(13),
	},
	84: {
		Flex_state: uint16(53),
	},
	85: {
		Flex_state: uint16(12),
	},
	86: {
		Flex_state: uint16(23),
	},
	87: {
		Flex_state: uint16(13),
	},
	88: {
		Flex_state: uint16(165),
	},
	89: {
		Flex_state: uint16(12),
	},
	90: {
		Flex_state: uint16(13),
	},
	91: {
		Flex_state: uint16(11),
	},
	92: {
		Flex_state: uint16(35),
	},
	93: {
		Flex_state: uint16(53),
	},
	94: {
		Flex_state: uint16(11),
	},
	95: {
		Flex_state: uint16(21),
	},
	96: {
		Flex_state: uint16(19),
	},
	97: {
		Flex_state: uint16(165),
	},
	98: {
		Flex_state:          uint16(16),
		Fexternal_lex_state: uint16(4),
	},
	99: {
		Flex_state: uint16(13),
	},
	100: {
		Flex_state:          uint16(16),
		Fexternal_lex_state: uint16(4),
	},
	101: {
		Flex_state: uint16(34),
	},
	102: {
		Flex_state: uint16(50),
	},
	103: {
		Flex_state: uint16(19),
	},
	104: {
		Flex_state: uint16(165),
	},
	105: {
		Flex_state: uint16(54),
	},
	106: {
		Flex_state: uint16(26),
	},
	107: {
		Flex_state: uint16(52),
	},
	108: {
		Flex_state: uint16(36),
	},
	109: {
		Flex_state: uint16(25),
	},
	110: {
		Flex_state: uint16(34),
	},
	111: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(2),
	},
	112: {
		Flex_state: uint16(26),
	},
	113: {
		Flex_state:          uint16(22),
		Fexternal_lex_state: uint16(4),
	},
	114: {
		Flex_state: uint16(25),
	},
	115: {
		Flex_state:          uint16(22),
		Fexternal_lex_state: uint16(4),
	},
	116: {
		Flex_state: uint16(26),
	},
	117: {
		Flex_state: uint16(7),
	},
	118: {
		Flex_state: uint16(21),
	},
	119: {
		Flex_state:          uint16(22),
		Fexternal_lex_state: uint16(4),
	},
	120: {
		Flex_state: uint16(7),
	},
	121: {
		Flex_state:          uint16(16),
		Fexternal_lex_state: uint16(4),
	},
	122: {
		Flex_state: uint16(26),
	},
	123: {
		Flex_state: uint16(26),
	},
	124: {
		Flex_state:          uint16(16),
		Fexternal_lex_state: uint16(4),
	},
	125: {
		Flex_state: uint16(26),
	},
	126: {
		Flex_state: uint16(26),
	},
	127: {
		Flex_state: uint16(6),
	},
	128: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	129: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(4),
	},
	130: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	131: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	132: {
		Flex_state: uint16(37),
	},
	133: {
		Flex_state: uint16(22),
	},
	134: {
		Flex_state: uint16(22),
	},
	135: {
		Flex_state: uint16(31),
	},
	136: {
		Flex_state: uint16(22),
	},
	137: {
		Flex_state: uint16(9),
	},
	138: {
		Flex_state: uint16(9),
	},
	139: {
		Flex_state: uint16(9),
	},
	140: {
		Flex_state: uint16(31),
	},
	141: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	142: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	143: {
		Flex_state: uint16(13),
	},
	144: {
		Flex_state: uint16(39),
	},
	145: {
		Flex_state: uint16(40),
	},
	146: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	147: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(2),
	},
	148: {
		Flex_state: uint16(6),
	},
	149: {
		Flex_state: uint16(10),
	},
	150: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(4),
	},
	151: {
		Flex_state: uint16(6),
	},
	152: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(4),
	},
	153: {
		Flex_state: uint16(165),
	},
	154: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	155: {
		Flex_state: uint16(39),
	},
	156: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	157: {
		Flex_state: uint16(31),
	},
	158: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	159: {
		Flex_state: uint16(6),
	},
	160: {
		Flex_state: uint16(6),
	},
	161: {
		Flex_state: uint16(10),
	},
	162: {
		Flex_state: uint16(46),
	},
	163: {
		Flex_state: uint16(13),
	},
	164: {
		Flex_state: uint16(9),
	},
	165: {
		Flex_state: uint16(165),
	},
	166: {
		Flex_state:          uint16(165),
		Fexternal_lex_state: uint16(5),
	},
	167: {
		Flex_state: uint16(15),
	},
	168: {
		Flex_state:          uint16(16),
		Fexternal_lex_state: uint16(4),
	},
	169: {
		Flex_state: uint16(16),
	},
	170: {
		Flex_state: uint16(13),
	},
	171: {
		Flex_state: uint16(46),
	},
	172: {
		Flex_state: uint16(46),
	},
	173: {
		Flex_state: uint16(15),
	},
	174: {
		Flex_state: uint16(15),
	},
	175: {
		Flex_state: uint16(13),
	},
	176: {
		Flex_state: uint16(53),
	},
	177: {
		Flex_state: uint16(53),
	},
	178: {
		Flex_state: uint16(53),
	},
	179: {
		Flex_state: uint16(165),
	},
	180: {
		Flex_state:          uint16(165),
		Fexternal_lex_state: uint16(5),
	},
	181: {
		Flex_state: uint16(23),
	},
	182: {
		Flex_state: uint16(23),
	},
	183: {
		Flex_state: uint16(23),
	},
	184: {
		Flex_state: uint16(11),
	},
	185: {
		Flex_state: uint16(11),
	},
	186: {
		Flex_state: uint16(46),
	},
	187: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(4),
	},
	188: {
		Flex_state: uint16(12),
	},
	189: {
		Flex_state: uint16(165),
	},
	190: {
		Flex_state:          uint16(165),
		Fexternal_lex_state: uint16(5),
	},
	191: {
		Flex_state: uint16(165),
	},
	192: {
		Flex_state: uint16(12),
	},
	193: {
		Flex_state: uint16(12),
	},
	194: {
		Flex_state: uint16(46),
	},
	195: {
		Flex_state: uint16(165),
	},
	196: {
		Flex_state: uint16(165),
	},
	197: {
		Flex_state: uint16(46),
	},
	198: {
		Flex_state: uint16(16),
	},
	199: {
		Flex_state: uint16(165),
	},
	200: {
		Flex_state: uint16(46),
	},
	201: {
		Flex_state: uint16(46),
	},
	202: {
		Flex_state: uint16(16),
	},
	203: {
		Flex_state: uint16(46),
	},
	204: {
		Flex_state: uint16(46),
	},
	205: {
		Flex_state: uint16(46),
	},
	206: {
		Flex_state: uint16(46),
	},
	207: {
		Flex_state: uint16(46),
	},
	208: {
		Flex_state: uint16(46),
	},
	209: {
		Flex_state: uint16(46),
	},
	210: {
		Flex_state: uint16(46),
	},
	211: {
		Flex_state: uint16(46),
	},
	212: {
		Flex_state: uint16(46),
	},
	213: {
		Flex_state: uint16(46),
	},
	214: {
		Flex_state: uint16(46),
	},
	215: {
		Flex_state: uint16(46),
	},
	216: {
		Flex_state: uint16(46),
	},
	217: {
		Flex_state: uint16(11),
	},
	218: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	219: {
		Flex_state: uint16(13),
	},
	220: {
		Flex_state: uint16(37),
	},
	221: {
		Flex_state: uint16(54),
	},
	222: {
		Flex_state: uint16(19),
	},
	223: {
		Flex_state: uint16(19),
	},
	224: {
		Flex_state:          uint16(165),
		Fexternal_lex_state: uint16(5),
	},
	225: {
		Flex_state: uint16(19),
	},
	226: {
		Flex_state: uint16(19),
	},
	227: {
		Flex_state: uint16(55),
	},
	228: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	229: {
		Flex_state: uint16(165),
	},
	230: {
		Flex_state: uint16(165),
	},
	231: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	232: {
		Flex_state: uint16(165),
	},
	233: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	234: {
		Flex_state: uint16(165),
	},
	235: {
		Flex_state: uint16(55),
	},
	236: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(4),
	},
	237: {
		Flex_state: uint16(13),
	},
	238: {
		Flex_state: uint16(9),
	},
	239: {
		Flex_state: uint16(19),
	},
	240: {
		Flex_state: uint16(165),
	},
	241: {
		Flex_state: uint16(16),
	},
	242: {
		Flex_state: uint16(45),
	},
	243: {
		Flex_state: uint16(15),
	},
	244: {
		Flex_state: uint16(165),
	},
	245: {
		Flex_state: uint16(13),
	},
	246: {
		Flex_state: uint16(13),
	},
	247: {
		Flex_state: uint16(9),
	},
	248: {
		Flex_state: uint16(56),
	},
	249: {
		Flex_state: uint16(165),
	},
	250: {
		Flex_state: uint16(9),
	},
	251: {
		Flex_state: uint16(56),
	},
	252: {
		Flex_state: uint16(165),
	},
	253: {
		Flex_state: uint16(165),
	},
	254: {
		Flex_state: uint16(165),
	},
	255: {
		Flex_state: uint16(9),
	},
	256: {
		Flex_state: uint16(9),
	},
	257: {
		Flex_state: uint16(9),
	},
	258: {
		Flex_state: uint16(9),
	},
	259: {
		Flex_state: uint16(200),
	},
	260: {
		Flex_state: uint16(258),
	},
	261: {
		Flex_state: uint16(9),
	},
	262: {
		Flex_state: uint16(51),
	},
	263: {
		Flex_state: uint16(9),
	},
	264: {
		Flex_state: uint16(165),
	},
	265: {
		Flex_state: uint16(9),
	},
	266: {
		Flex_state: uint16(165),
	},
	267: {
		Flex_state: uint16(165),
	},
	268: {
		Flex_state: uint16(57),
	},
	269: {
		Flex_state: uint16(9),
	},
	270: {
		Flex_state: uint16(200),
	},
	271: {
		Flex_state: uint16(165),
	},
	272: {
		Flex_state: uint16(165),
	},
	273: {
		Flex_state: uint16(9),
	},
	274: {
		Flex_state: uint16(165),
	},
	275: {
		Flex_state: uint16(9),
	},
	276: {
		Flex_state: uint16(165),
	},
	277: {
		Flex_state: uint16(57),
	},
	278: {
		Flex_state: uint16(9),
	},
	279: {
		Flex_state: uint16(165),
	},
	280: {
		Flex_state: uint16(165),
	},
	281: {
		Flex_state: uint16(57),
	},
	282: {
		Flex_state: uint16(9),
	},
	283: {
		Flex_state: uint16(165),
	},
	284: {
		Flex_state: uint16(165),
	},
	285: {
		Flex_state: uint16(55),
	},
	286: {
		Flex_state: uint16(258),
	},
	287: {
		Flex_state: uint16(165),
	},
	288: {
		Flex_state: uint16(9),
	},
	289: {
		Flex_state: uint16(9),
	},
	290: {
		Flex_state: uint16(165),
	},
	291: {
		Flex_state: uint16(9),
	},
	292: {
		Flex_state: uint16(9),
	},
	293: {
		Flex_state: uint16(165),
	},
	294: {
		Flex_state: uint16(47),
	},
	295: {
		Flex_state: uint16(9),
	},
	296: {
		Flex_state: uint16(165),
	},
	297: {
		Flex_state: uint16(9),
	},
	298: {
		Flex_state: uint16(165),
	},
	299: {
		Flex_state: uint16(47),
	},
	300: {
		Flex_state: uint16(165),
	},
	301: {
		Flex_state: uint16(9),
	},
	302: {
		Flex_state: uint16(165),
	},
	303: {
		Flex_state: uint16(49),
	},
	304: {
		Flex_state: uint16(165),
	},
	305: {
		Flex_state: uint16(165),
	},
	306: {
		Flex_state: uint16(56),
	},
	307: {
		Flex_state: uint16(165),
	},
	308: {
		Flex_state: uint16(56),
	},
	309: {
		Flex_state: uint16(165),
	},
	310: {
		Flex_state: uint16(56),
	},
	311: {
		Flex_state: uint16(165),
	},
	312: {
		Flex_state: uint16(56),
	},
	313: {
		Flex_state: uint16(9),
	},
	314: {
		Flex_state: uint16(56),
	},
	315: {
		Flex_state: uint16(56),
	},
	316: {
		Flex_state: uint16(56),
	},
	317: {
		Flex_state: uint16(56),
	},
	318: {
		Flex_state: uint16(56),
	},
	319: {
		Flex_state: uint16(56),
	},
	320: {
		Flex_state: uint16(56),
	},
	321: {
		Flex_state: uint16(56),
	},
	322: {
		Flex_state: uint16(49),
	},
	323: {
		Flex_state: uint16(49),
	},
	324: {
		Flex_state: uint16(57),
	},
}

var ts_parse_table = [2][164]uint16_t{
	0: {
		0:  uint16(1),
		14: uint16(1),
		16: uint16(1),
		20: uint16(1),
		31: uint16(1),
		35: uint16(1),
		36: uint16(1),
		37: uint16(1),
		39: uint16(1),
		40: uint16(1),
		47: uint16(1),
		50: uint16(1),
		52: uint16(1),
		56: uint16(1),
		61: uint16(1),
		64: uint16(3),
		66: uint16(1),
		67: uint16(1),
		68: uint16(1),
		69: uint16(1),
		73: uint16(1),
		74: uint16(1),
		81: uint16(5),
		82: uint16(1),
		83: uint16(1),
		84: uint16(1),
		85: uint16(1),
		86: uint16(1),
	},
	1: {
		0:   uint16(7),
		2:   uint16(9),
		4:   uint16(11),
		5:   uint16(13),
		6:   uint16(15),
		7:   uint16(17),
		8:   uint16(19),
		9:   uint16(21),
		10:  uint16(23),
		11:  uint16(25),
		12:  uint16(27),
		13:  uint16(29),
		17:  uint16(31),
		18:  uint16(33),
		21:  uint16(35),
		22:  uint16(37),
		25:  uint16(39),
		27:  uint16(41),
		28:  uint16(43),
		30:  uint16(45),
		64:  uint16(3),
		81:  uint16(3),
		87:  uint16(253),
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
		103: uint16(255),
		104: uint16(255),
		105: uint16(255),
		106: uint16(255),
		108: uint16(255),
		109: uint16(255),
		110: uint16(255),
		111: uint16(255),
		139: uint16(2),
	},
}

var ts_small_parse_table = [5079]uint16_t{
	0:    uint16(23),
	1:    uint16(9),
	2:    uint16(1),
	3:    uint16(aux_sym_from_instruction_token1),
	4:    uint16(11),
	5:    uint16(1),
	6:    uint16(aux_sym_run_instruction_token1),
	7:    uint16(13),
	8:    uint16(1),
	9:    uint16(aux_sym_cmd_instruction_token1),
	10:   uint16(15),
	11:   uint16(1),
	12:   uint16(aux_sym_label_instruction_token1),
	13:   uint16(17),
	14:   uint16(1),
	15:   uint16(aux_sym_expose_instruction_token1),
	16:   uint16(19),
	17:   uint16(1),
	18:   uint16(aux_sym_env_instruction_token1),
	19:   uint16(21),
	20:   uint16(1),
	21:   uint16(aux_sym_add_instruction_token1),
	22:   uint16(23),
	23:   uint16(1),
	24:   uint16(aux_sym_copy_instruction_token1),
	25:   uint16(25),
	26:   uint16(1),
	27:   uint16(aux_sym_entrypoint_instruction_token1),
	28:   uint16(27),
	29:   uint16(1),
	30:   uint16(aux_sym_volume_instruction_token1),
	31:   uint16(29),
	32:   uint16(1),
	33:   uint16(aux_sym_user_instruction_token1),
	34:   uint16(31),
	35:   uint16(1),
	36:   uint16(aux_sym_workdir_instruction_token1),
	37:   uint16(33),
	38:   uint16(1),
	39:   uint16(aux_sym_arg_instruction_token1),
	40:   uint16(35),
	41:   uint16(1),
	42:   uint16(aux_sym_onbuild_instruction_token1),
	43:   uint16(37),
	44:   uint16(1),
	45:   uint16(aux_sym_stopsignal_instruction_token1),
	46:   uint16(39),
	47:   uint16(1),
	48:   uint16(aux_sym_healthcheck_instruction_token1),
	49:   uint16(41),
	50:   uint16(1),
	51:   uint16(aux_sym_shell_instruction_token1),
	52:   uint16(43),
	53:   uint16(1),
	54:   uint16(aux_sym_maintainer_instruction_token1),
	55:   uint16(45),
	56:   uint16(1),
	57:   uint16(aux_sym_cross_build_instruction_token1),
	58:   uint16(47),
	59:   uint16(1),
	61:   uint16(3),
	62:   uint16(1),
	63:   uint16(aux_sym_source_file_repeat1),
	64:   uint16(3),
	65:   uint16(2),
	66:   uint16(sym_line_continuation),
	67:   uint16(sym_comment),
	68:   uint16(255),
	69:   uint16(20),
	70:   uint16(sym__instruction),
	71:   uint16(sym_from_instruction),
	72:   uint16(sym_run_instruction),
	73:   uint16(sym_cmd_instruction),
	74:   uint16(sym_label_instruction),
	75:   uint16(sym_expose_instruction),
	76:   uint16(sym_env_instruction),
	77:   uint16(sym_add_instruction),
	78:   uint16(sym_copy_instruction),
	79:   uint16(sym_entrypoint_instruction),
	80:   uint16(sym_volume_instruction),
	81:   uint16(sym_user_instruction),
	82:   uint16(sym_workdir_instruction),
	83:   uint16(sym_arg_instruction),
	84:   uint16(sym_onbuild_instruction),
	85:   uint16(sym_stopsignal_instruction),
	86:   uint16(sym_healthcheck_instruction),
	87:   uint16(sym_shell_instruction),
	88:   uint16(sym_maintainer_instruction),
	89:   uint16(sym_cross_build_instruction),
	90:   uint16(23),
	91:   uint16(49),
	92:   uint16(1),
	94:   uint16(51),
	95:   uint16(1),
	96:   uint16(aux_sym_from_instruction_token1),
	97:   uint16(54),
	98:   uint16(1),
	99:   uint16(aux_sym_run_instruction_token1),
	100:  uint16(57),
	101:  uint16(1),
	102:  uint16(aux_sym_cmd_instruction_token1),
	103:  uint16(60),
	104:  uint16(1),
	105:  uint16(aux_sym_label_instruction_token1),
	106:  uint16(63),
	107:  uint16(1),
	108:  uint16(aux_sym_expose_instruction_token1),
	109:  uint16(66),
	110:  uint16(1),
	111:  uint16(aux_sym_env_instruction_token1),
	112:  uint16(69),
	113:  uint16(1),
	114:  uint16(aux_sym_add_instruction_token1),
	115:  uint16(72),
	116:  uint16(1),
	117:  uint16(aux_sym_copy_instruction_token1),
	118:  uint16(75),
	119:  uint16(1),
	120:  uint16(aux_sym_entrypoint_instruction_token1),
	121:  uint16(78),
	122:  uint16(1),
	123:  uint16(aux_sym_volume_instruction_token1),
	124:  uint16(81),
	125:  uint16(1),
	126:  uint16(aux_sym_user_instruction_token1),
	127:  uint16(84),
	128:  uint16(1),
	129:  uint16(aux_sym_workdir_instruction_token1),
	130:  uint16(87),
	131:  uint16(1),
	132:  uint16(aux_sym_arg_instruction_token1),
	133:  uint16(90),
	134:  uint16(1),
	135:  uint16(aux_sym_onbuild_instruction_token1),
	136:  uint16(93),
	137:  uint16(1),
	138:  uint16(aux_sym_stopsignal_instruction_token1),
	139:  uint16(96),
	140:  uint16(1),
	141:  uint16(aux_sym_healthcheck_instruction_token1),
	142:  uint16(99),
	143:  uint16(1),
	144:  uint16(aux_sym_shell_instruction_token1),
	145:  uint16(102),
	146:  uint16(1),
	147:  uint16(aux_sym_maintainer_instruction_token1),
	148:  uint16(105),
	149:  uint16(1),
	150:  uint16(aux_sym_cross_build_instruction_token1),
	151:  uint16(3),
	152:  uint16(1),
	153:  uint16(aux_sym_source_file_repeat1),
	154:  uint16(3),
	155:  uint16(2),
	156:  uint16(sym_line_continuation),
	157:  uint16(sym_comment),
	158:  uint16(255),
	159:  uint16(20),
	160:  uint16(sym__instruction),
	161:  uint16(sym_from_instruction),
	162:  uint16(sym_run_instruction),
	163:  uint16(sym_cmd_instruction),
	164:  uint16(sym_label_instruction),
	165:  uint16(sym_expose_instruction),
	166:  uint16(sym_env_instruction),
	167:  uint16(sym_add_instruction),
	168:  uint16(sym_copy_instruction),
	169:  uint16(sym_entrypoint_instruction),
	170:  uint16(sym_volume_instruction),
	171:  uint16(sym_user_instruction),
	172:  uint16(sym_workdir_instruction),
	173:  uint16(sym_arg_instruction),
	174:  uint16(sym_onbuild_instruction),
	175:  uint16(sym_stopsignal_instruction),
	176:  uint16(sym_healthcheck_instruction),
	177:  uint16(sym_shell_instruction),
	178:  uint16(sym_maintainer_instruction),
	179:  uint16(sym_cross_build_instruction),
	180:  uint16(21),
	181:  uint16(9),
	182:  uint16(1),
	183:  uint16(aux_sym_from_instruction_token1),
	184:  uint16(11),
	185:  uint16(1),
	186:  uint16(aux_sym_run_instruction_token1),
	187:  uint16(13),
	188:  uint16(1),
	189:  uint16(aux_sym_cmd_instruction_token1),
	190:  uint16(15),
	191:  uint16(1),
	192:  uint16(aux_sym_label_instruction_token1),
	193:  uint16(17),
	194:  uint16(1),
	195:  uint16(aux_sym_expose_instruction_token1),
	196:  uint16(19),
	197:  uint16(1),
	198:  uint16(aux_sym_env_instruction_token1),
	199:  uint16(21),
	200:  uint16(1),
	201:  uint16(aux_sym_add_instruction_token1),
	202:  uint16(23),
	203:  uint16(1),
	204:  uint16(aux_sym_copy_instruction_token1),
	205:  uint16(25),
	206:  uint16(1),
	207:  uint16(aux_sym_entrypoint_instruction_token1),
	208:  uint16(27),
	209:  uint16(1),
	210:  uint16(aux_sym_volume_instruction_token1),
	211:  uint16(29),
	212:  uint16(1),
	213:  uint16(aux_sym_user_instruction_token1),
	214:  uint16(31),
	215:  uint16(1),
	216:  uint16(aux_sym_workdir_instruction_token1),
	217:  uint16(33),
	218:  uint16(1),
	219:  uint16(aux_sym_arg_instruction_token1),
	220:  uint16(35),
	221:  uint16(1),
	222:  uint16(aux_sym_onbuild_instruction_token1),
	223:  uint16(37),
	224:  uint16(1),
	225:  uint16(aux_sym_stopsignal_instruction_token1),
	226:  uint16(39),
	227:  uint16(1),
	228:  uint16(aux_sym_healthcheck_instruction_token1),
	229:  uint16(41),
	230:  uint16(1),
	231:  uint16(aux_sym_shell_instruction_token1),
	232:  uint16(43),
	233:  uint16(1),
	234:  uint16(aux_sym_maintainer_instruction_token1),
	235:  uint16(45),
	236:  uint16(1),
	237:  uint16(aux_sym_cross_build_instruction_token1),
	238:  uint16(3),
	239:  uint16(2),
	240:  uint16(sym_line_continuation),
	241:  uint16(sym_comment),
	242:  uint16(265),
	243:  uint16(20),
	244:  uint16(sym__instruction),
	245:  uint16(sym_from_instruction),
	246:  uint16(sym_run_instruction),
	247:  uint16(sym_cmd_instruction),
	248:  uint16(sym_label_instruction),
	249:  uint16(sym_expose_instruction),
	250:  uint16(sym_env_instruction),
	251:  uint16(sym_add_instruction),
	252:  uint16(sym_copy_instruction),
	253:  uint16(sym_entrypoint_instruction),
	254:  uint16(sym_volume_instruction),
	255:  uint16(sym_user_instruction),
	256:  uint16(sym_workdir_instruction),
	257:  uint16(sym_arg_instruction),
	258:  uint16(sym_onbuild_instruction),
	259:  uint16(sym_stopsignal_instruction),
	260:  uint16(sym_healthcheck_instruction),
	261:  uint16(sym_shell_instruction),
	262:  uint16(sym_maintainer_instruction),
	263:  uint16(sym_cross_build_instruction),
	264:  uint16(2),
	265:  uint16(3),
	266:  uint16(2),
	267:  uint16(sym_line_continuation),
	268:  uint16(sym_comment),
	269:  uint16(49),
	270:  uint16(20),
	272:  uint16(aux_sym_from_instruction_token1),
	273:  uint16(aux_sym_run_instruction_token1),
	274:  uint16(aux_sym_cmd_instruction_token1),
	275:  uint16(aux_sym_label_instruction_token1),
	276:  uint16(aux_sym_expose_instruction_token1),
	277:  uint16(aux_sym_env_instruction_token1),
	278:  uint16(aux_sym_add_instruction_token1),
	279:  uint16(aux_sym_copy_instruction_token1),
	280:  uint16(aux_sym_entrypoint_instruction_token1),
	281:  uint16(aux_sym_volume_instruction_token1),
	282:  uint16(aux_sym_user_instruction_token1),
	283:  uint16(aux_sym_workdir_instruction_token1),
	284:  uint16(aux_sym_arg_instruction_token1),
	285:  uint16(aux_sym_onbuild_instruction_token1),
	286:  uint16(aux_sym_stopsignal_instruction_token1),
	287:  uint16(aux_sym_healthcheck_instruction_token1),
	288:  uint16(aux_sym_shell_instruction_token1),
	289:  uint16(aux_sym_maintainer_instruction_token1),
	290:  uint16(aux_sym_cross_build_instruction_token1),
	291:  uint16(9),
	292:  uint16(110),
	293:  uint16(1),
	294:  uint16(anon_sym_DASH_DASH),
	295:  uint16(112),
	296:  uint16(1),
	297:  uint16(anon_sym_LBRACK),
	298:  uint16(114),
	299:  uint16(1),
	300:  uint16(sym_heredoc_marker),
	301:  uint16(22),
	302:  uint16(1),
	303:  uint16(aux_sym_shell_fragment_repeat1),
	304:  uint16(150),
	305:  uint16(1),
	306:  uint16(sym_shell_fragment),
	307:  uint16(3),
	308:  uint16(2),
	309:  uint16(sym_line_continuation),
	310:  uint16(sym_comment),
	311:  uint16(131),
	312:  uint16(2),
	313:  uint16(sym_shell_command),
	314:  uint16(sym_json_string_array),
	315:  uint16(19),
	316:  uint16(3),
	317:  uint16(sym_param),
	318:  uint16(sym_mount_param),
	319:  uint16(aux_sym_run_instruction_repeat1),
	320:  uint16(108),
	321:  uint16(4),
	322:  uint16(aux_sym_path_token2),
	323:  uint16(aux_sym_shell_fragment_token2),
	324:  uint16(aux_sym_shell_fragment_token3),
	325:  uint16(aux_sym_shell_fragment_token4),
	326:  uint16(9),
	327:  uint16(110),
	328:  uint16(1),
	329:  uint16(anon_sym_DASH_DASH),
	330:  uint16(112),
	331:  uint16(1),
	332:  uint16(anon_sym_LBRACK),
	333:  uint16(114),
	334:  uint16(1),
	335:  uint16(sym_heredoc_marker),
	336:  uint16(22),
	337:  uint16(1),
	338:  uint16(aux_sym_shell_fragment_repeat1),
	339:  uint16(150),
	340:  uint16(1),
	341:  uint16(sym_shell_fragment),
	342:  uint16(3),
	343:  uint16(2),
	344:  uint16(sym_line_continuation),
	345:  uint16(sym_comment),
	346:  uint16(141),
	347:  uint16(2),
	348:  uint16(sym_shell_command),
	349:  uint16(sym_json_string_array),
	350:  uint16(6),
	351:  uint16(3),
	352:  uint16(sym_param),
	353:  uint16(sym_mount_param),
	354:  uint16(aux_sym_run_instruction_repeat1),
	355:  uint16(108),
	356:  uint16(4),
	357:  uint16(aux_sym_path_token2),
	358:  uint16(aux_sym_shell_fragment_token2),
	359:  uint16(aux_sym_shell_fragment_token3),
	360:  uint16(aux_sym_shell_fragment_token4),
	361:  uint16(9),
	362:  uint16(116),
	363:  uint16(1),
	364:  uint16(anon_sym_LF),
	365:  uint16(118),
	366:  uint16(1),
	367:  uint16(anon_sym_DOLLAR2),
	368:  uint16(120),
	369:  uint16(1),
	370:  uint16(aux_sym__env_key_token1),
	371:  uint16(122),
	372:  uint16(1),
	373:  uint16(anon_sym_DQUOTE),
	374:  uint16(124),
	375:  uint16(1),
	376:  uint16(anon_sym_SQUOTE),
	377:  uint16(5),
	378:  uint16(2),
	379:  uint16(sym_line_continuation),
	380:  uint16(sym_comment),
	381:  uint16(126),
	382:  uint16(2),
	383:  uint16(aux_sym_unquoted_string_token1),
	384:  uint16(anon_sym_BSLASH2),
	385:  uint16(47),
	386:  uint16(3),
	387:  uint16(sym__immediate_expansion),
	388:  uint16(sym__imm_expansion),
	389:  uint16(aux_sym_unquoted_string_repeat1),
	390:  uint16(239),
	391:  uint16(3),
	392:  uint16(sym_double_quoted_string),
	393:  uint16(sym_single_quoted_string),
	394:  uint16(sym_unquoted_string),
	395:  uint16(8),
	396:  uint16(3),
	397:  uint16(1),
	398:  uint16(sym_line_continuation),
	399:  uint16(5),
	400:  uint16(1),
	401:  uint16(sym_comment),
	402:  uint16(128),
	403:  uint16(1),
	404:  uint16(anon_sym_DOLLAR2),
	405:  uint16(130),
	406:  uint16(1),
	407:  uint16(anon_sym_DQUOTE),
	408:  uint16(132),
	409:  uint16(1),
	410:  uint16(anon_sym_SQUOTE),
	411:  uint16(134),
	412:  uint16(2),
	413:  uint16(aux_sym_unquoted_string_token1),
	414:  uint16(anon_sym_BSLASH2),
	415:  uint16(18),
	416:  uint16(3),
	417:  uint16(sym__immediate_expansion),
	418:  uint16(sym__imm_expansion),
	419:  uint16(aux_sym_unquoted_string_repeat1),
	420:  uint16(313),
	421:  uint16(3),
	422:  uint16(sym_double_quoted_string),
	423:  uint16(sym_single_quoted_string),
	424:  uint16(sym_unquoted_string),
	425:  uint16(8),
	426:  uint16(3),
	427:  uint16(1),
	428:  uint16(sym_line_continuation),
	429:  uint16(5),
	430:  uint16(1),
	431:  uint16(sym_comment),
	432:  uint16(128),
	433:  uint16(1),
	434:  uint16(anon_sym_DOLLAR2),
	435:  uint16(130),
	436:  uint16(1),
	437:  uint16(anon_sym_DQUOTE),
	438:  uint16(132),
	439:  uint16(1),
	440:  uint16(anon_sym_SQUOTE),
	441:  uint16(134),
	442:  uint16(2),
	443:  uint16(aux_sym_unquoted_string_token1),
	444:  uint16(anon_sym_BSLASH2),
	445:  uint16(18),
	446:  uint16(3),
	447:  uint16(sym__immediate_expansion),
	448:  uint16(sym__imm_expansion),
	449:  uint16(aux_sym_unquoted_string_repeat1),
	450:  uint16(127),
	451:  uint16(3),
	452:  uint16(sym_double_quoted_string),
	453:  uint16(sym_single_quoted_string),
	454:  uint16(sym_unquoted_string),
	455:  uint16(8),
	456:  uint16(3),
	457:  uint16(1),
	458:  uint16(sym_line_continuation),
	459:  uint16(5),
	460:  uint16(1),
	461:  uint16(sym_comment),
	462:  uint16(128),
	463:  uint16(1),
	464:  uint16(anon_sym_DOLLAR2),
	465:  uint16(130),
	466:  uint16(1),
	467:  uint16(anon_sym_DQUOTE),
	468:  uint16(132),
	469:  uint16(1),
	470:  uint16(anon_sym_SQUOTE),
	471:  uint16(134),
	472:  uint16(2),
	473:  uint16(aux_sym_unquoted_string_token1),
	474:  uint16(anon_sym_BSLASH2),
	475:  uint16(18),
	476:  uint16(3),
	477:  uint16(sym__immediate_expansion),
	478:  uint16(sym__imm_expansion),
	479:  uint16(aux_sym_unquoted_string_repeat1),
	480:  uint16(258),
	481:  uint16(3),
	482:  uint16(sym_double_quoted_string),
	483:  uint16(sym_single_quoted_string),
	484:  uint16(sym_unquoted_string),
	485:  uint16(8),
	486:  uint16(138),
	487:  uint16(1),
	488:  uint16(aux_sym_shell_fragment_token2),
	489:  uint16(140),
	490:  uint16(1),
	491:  uint16(anon_sym_LBRACK),
	492:  uint16(142),
	493:  uint16(1),
	494:  uint16(sym_heredoc_marker),
	495:  uint16(43),
	496:  uint16(1),
	497:  uint16(aux_sym_shell_fragment_repeat1),
	498:  uint16(167),
	499:  uint16(1),
	500:  uint16(sym_shell_fragment),
	501:  uint16(3),
	502:  uint16(2),
	503:  uint16(sym_line_continuation),
	504:  uint16(sym_comment),
	505:  uint16(263),
	506:  uint16(2),
	507:  uint16(sym_shell_command),
	508:  uint16(sym_json_string_array),
	509:  uint16(136),
	510:  uint16(3),
	511:  uint16(aux_sym_path_token2),
	512:  uint16(aux_sym_shell_fragment_token3),
	513:  uint16(aux_sym_shell_fragment_token4),
	514:  uint16(11),
	515:  uint16(3),
	516:  uint16(1),
	517:  uint16(sym_line_continuation),
	518:  uint16(5),
	519:  uint16(1),
	520:  uint16(sym_comment),
	521:  uint16(144),
	522:  uint16(1),
	523:  uint16(aux_sym_path_token1),
	524:  uint16(146),
	525:  uint16(1),
	526:  uint16(aux_sym_path_with_heredoc_token1),
	527:  uint16(148),
	528:  uint16(1),
	529:  uint16(anon_sym_DOLLAR),
	530:  uint16(150),
	531:  uint16(1),
	532:  uint16(anon_sym_DASH_DASH),
	533:  uint16(152),
	534:  uint16(1),
	535:  uint16(sym_heredoc_marker),
	536:  uint16(56),
	537:  uint16(1),
	538:  uint16(aux_sym_add_instruction_repeat2),
	539:  uint16(84),
	540:  uint16(1),
	541:  uint16(sym_expansion),
	542:  uint16(299),
	543:  uint16(1),
	544:  uint16(sym_path_with_heredoc),
	545:  uint16(17),
	546:  uint16(2),
	547:  uint16(sym_param),
	548:  uint16(aux_sym_add_instruction_repeat1),
	549:  uint16(11),
	550:  uint16(3),
	551:  uint16(1),
	552:  uint16(sym_line_continuation),
	553:  uint16(5),
	554:  uint16(1),
	555:  uint16(sym_comment),
	556:  uint16(144),
	557:  uint16(1),
	558:  uint16(aux_sym_path_token1),
	559:  uint16(146),
	560:  uint16(1),
	561:  uint16(aux_sym_path_with_heredoc_token1),
	562:  uint16(148),
	563:  uint16(1),
	564:  uint16(anon_sym_DOLLAR),
	565:  uint16(150),
	566:  uint16(1),
	567:  uint16(anon_sym_DASH_DASH),
	568:  uint16(152),
	569:  uint16(1),
	570:  uint16(sym_heredoc_marker),
	571:  uint16(74),
	572:  uint16(1),
	573:  uint16(aux_sym_add_instruction_repeat2),
	574:  uint16(84),
	575:  uint16(1),
	576:  uint16(sym_expansion),
	577:  uint16(299),
	578:  uint16(1),
	579:  uint16(sym_path_with_heredoc),
	580:  uint16(15),
	581:  uint16(2),
	582:  uint16(sym_param),
	583:  uint16(aux_sym_add_instruction_repeat1),
	584:  uint16(11),
	585:  uint16(3),
	586:  uint16(1),
	587:  uint16(sym_line_continuation),
	588:  uint16(5),
	589:  uint16(1),
	590:  uint16(sym_comment),
	591:  uint16(144),
	592:  uint16(1),
	593:  uint16(aux_sym_path_token1),
	594:  uint16(146),
	595:  uint16(1),
	596:  uint16(aux_sym_path_with_heredoc_token1),
	597:  uint16(148),
	598:  uint16(1),
	599:  uint16(anon_sym_DOLLAR),
	600:  uint16(150),
	601:  uint16(1),
	602:  uint16(anon_sym_DASH_DASH),
	603:  uint16(152),
	604:  uint16(1),
	605:  uint16(sym_heredoc_marker),
	606:  uint16(65),
	607:  uint16(1),
	608:  uint16(aux_sym_add_instruction_repeat2),
	609:  uint16(84),
	610:  uint16(1),
	611:  uint16(sym_expansion),
	612:  uint16(299),
	613:  uint16(1),
	614:  uint16(sym_path_with_heredoc),
	615:  uint16(64),
	616:  uint16(2),
	617:  uint16(sym_param),
	618:  uint16(aux_sym_add_instruction_repeat1),
	619:  uint16(8),
	620:  uint16(138),
	621:  uint16(1),
	622:  uint16(aux_sym_shell_fragment_token2),
	623:  uint16(140),
	624:  uint16(1),
	625:  uint16(anon_sym_LBRACK),
	626:  uint16(142),
	627:  uint16(1),
	628:  uint16(sym_heredoc_marker),
	629:  uint16(43),
	630:  uint16(1),
	631:  uint16(aux_sym_shell_fragment_repeat1),
	632:  uint16(167),
	633:  uint16(1),
	634:  uint16(sym_shell_fragment),
	635:  uint16(3),
	636:  uint16(2),
	637:  uint16(sym_line_continuation),
	638:  uint16(sym_comment),
	639:  uint16(292),
	640:  uint16(2),
	641:  uint16(sym_shell_command),
	642:  uint16(sym_json_string_array),
	643:  uint16(136),
	644:  uint16(3),
	645:  uint16(aux_sym_path_token2),
	646:  uint16(aux_sym_shell_fragment_token3),
	647:  uint16(aux_sym_shell_fragment_token4),
	648:  uint16(11),
	649:  uint16(3),
	650:  uint16(1),
	651:  uint16(sym_line_continuation),
	652:  uint16(5),
	653:  uint16(1),
	654:  uint16(sym_comment),
	655:  uint16(144),
	656:  uint16(1),
	657:  uint16(aux_sym_path_token1),
	658:  uint16(146),
	659:  uint16(1),
	660:  uint16(aux_sym_path_with_heredoc_token1),
	661:  uint16(148),
	662:  uint16(1),
	663:  uint16(anon_sym_DOLLAR),
	664:  uint16(150),
	665:  uint16(1),
	666:  uint16(anon_sym_DASH_DASH),
	667:  uint16(152),
	668:  uint16(1),
	669:  uint16(sym_heredoc_marker),
	670:  uint16(69),
	671:  uint16(1),
	672:  uint16(aux_sym_add_instruction_repeat2),
	673:  uint16(84),
	674:  uint16(1),
	675:  uint16(sym_expansion),
	676:  uint16(299),
	677:  uint16(1),
	678:  uint16(sym_path_with_heredoc),
	679:  uint16(64),
	680:  uint16(2),
	681:  uint16(sym_param),
	682:  uint16(aux_sym_add_instruction_repeat1),
	683:  uint16(6),
	684:  uint16(128),
	685:  uint16(1),
	686:  uint16(anon_sym_DOLLAR2),
	687:  uint16(154),
	688:  uint16(1),
	689:  uint16(anon_sym_LF),
	690:  uint16(5),
	691:  uint16(2),
	692:  uint16(sym_line_continuation),
	693:  uint16(sym_comment),
	694:  uint16(158),
	695:  uint16(2),
	696:  uint16(aux_sym_unquoted_string_token1),
	697:  uint16(anon_sym_BSLASH2),
	698:  uint16(156),
	699:  uint16(3),
	700:  uint16(aux_sym_label_pair_token1),
	701:  uint16(anon_sym_DQUOTE),
	702:  uint16(anon_sym_SQUOTE),
	703:  uint16(20),
	704:  uint16(3),
	705:  uint16(sym__immediate_expansion),
	706:  uint16(sym__imm_expansion),
	707:  uint16(aux_sym_unquoted_string_repeat1),
	708:  uint16(5),
	709:  uint16(162),
	710:  uint16(1),
	711:  uint16(anon_sym_DASH_DASH),
	712:  uint16(3),
	713:  uint16(2),
	714:  uint16(sym_line_continuation),
	715:  uint16(sym_comment),
	716:  uint16(165),
	717:  uint16(2),
	718:  uint16(sym_heredoc_marker),
	719:  uint16(anon_sym_LBRACK),
	720:  uint16(19),
	721:  uint16(3),
	722:  uint16(sym_param),
	723:  uint16(sym_mount_param),
	724:  uint16(aux_sym_run_instruction_repeat1),
	725:  uint16(160),
	726:  uint16(4),
	727:  uint16(aux_sym_path_token2),
	728:  uint16(aux_sym_shell_fragment_token2),
	729:  uint16(aux_sym_shell_fragment_token3),
	730:  uint16(aux_sym_shell_fragment_token4),
	731:  uint16(6),
	732:  uint16(167),
	733:  uint16(1),
	734:  uint16(anon_sym_LF),
	735:  uint16(169),
	736:  uint16(1),
	737:  uint16(anon_sym_DOLLAR2),
	738:  uint16(5),
	739:  uint16(2),
	740:  uint16(sym_line_continuation),
	741:  uint16(sym_comment),
	742:  uint16(174),
	743:  uint16(2),
	744:  uint16(aux_sym_unquoted_string_token1),
	745:  uint16(anon_sym_BSLASH2),
	746:  uint16(172),
	747:  uint16(3),
	748:  uint16(aux_sym_label_pair_token1),
	749:  uint16(anon_sym_DQUOTE),
	750:  uint16(anon_sym_SQUOTE),
	751:  uint16(20),
	752:  uint16(3),
	753:  uint16(sym__immediate_expansion),
	754:  uint16(sym__imm_expansion),
	755:  uint16(aux_sym_unquoted_string_repeat1),
	756:  uint16(6),
	757:  uint16(179),
	758:  uint16(1),
	759:  uint16(aux_sym_from_instruction_token2),
	760:  uint16(181),
	761:  uint16(1),
	762:  uint16(anon_sym_DOLLAR2),
	763:  uint16(183),
	764:  uint16(1),
	765:  uint16(aux_sym_image_name_token2),
	766:  uint16(5),
	767:  uint16(2),
	768:  uint16(sym_line_continuation),
	769:  uint16(sym_comment),
	770:  uint16(177),
	771:  uint16(3),
	772:  uint16(anon_sym_LF),
	773:  uint16(anon_sym_COLON),
	774:  uint16(anon_sym_AT),
	775:  uint16(24),
	776:  uint16(3),
	777:  uint16(sym__immediate_expansion),
	778:  uint16(sym__imm_expansion),
	779:  uint16(aux_sym_image_name_repeat1),
	780:  uint16(6),
	781:  uint16(114),
	782:  uint16(1),
	783:  uint16(sym_heredoc_marker),
	784:  uint16(189),
	785:  uint16(1),
	786:  uint16(sym_required_line_continuation),
	787:  uint16(23),
	788:  uint16(1),
	789:  uint16(aux_sym_shell_fragment_repeat1),
	790:  uint16(5),
	791:  uint16(2),
	792:  uint16(sym_line_continuation),
	793:  uint16(sym_comment),
	794:  uint16(185),
	795:  uint16(2),
	796:  uint16(sym_heredoc_nl),
	797:  uint16(anon_sym_LF),
	798:  uint16(187),
	799:  uint16(4),
	800:  uint16(aux_sym_path_token2),
	801:  uint16(aux_sym_shell_fragment_token2),
	802:  uint16(aux_sym_shell_fragment_token3),
	803:  uint16(aux_sym_shell_fragment_token4),
	804:  uint16(6),
	805:  uint16(196),
	806:  uint16(1),
	807:  uint16(sym_required_line_continuation),
	808:  uint16(198),
	809:  uint16(1),
	810:  uint16(sym_heredoc_marker),
	811:  uint16(23),
	812:  uint16(1),
	813:  uint16(aux_sym_shell_fragment_repeat1),
	814:  uint16(5),
	815:  uint16(2),
	816:  uint16(sym_line_continuation),
	817:  uint16(sym_comment),
	818:  uint16(191),
	819:  uint16(2),
	820:  uint16(sym_heredoc_nl),
	821:  uint16(anon_sym_LF),
	822:  uint16(193),
	823:  uint16(4),
	824:  uint16(aux_sym_path_token2),
	825:  uint16(aux_sym_shell_fragment_token2),
	826:  uint16(aux_sym_shell_fragment_token3),
	827:  uint16(aux_sym_shell_fragment_token4),
	828:  uint16(6),
	829:  uint16(181),
	830:  uint16(1),
	831:  uint16(anon_sym_DOLLAR2),
	832:  uint16(203),
	833:  uint16(1),
	834:  uint16(aux_sym_from_instruction_token2),
	835:  uint16(205),
	836:  uint16(1),
	837:  uint16(aux_sym_image_name_token2),
	838:  uint16(5),
	839:  uint16(2),
	840:  uint16(sym_line_continuation),
	841:  uint16(sym_comment),
	842:  uint16(201),
	843:  uint16(3),
	844:  uint16(anon_sym_LF),
	845:  uint16(anon_sym_COLON),
	846:  uint16(anon_sym_AT),
	847:  uint16(25),
	848:  uint16(3),
	849:  uint16(sym__immediate_expansion),
	850:  uint16(sym__imm_expansion),
	851:  uint16(aux_sym_image_name_repeat1),
	852:  uint16(6),
	853:  uint16(209),
	854:  uint16(1),
	855:  uint16(aux_sym_from_instruction_token2),
	856:  uint16(211),
	857:  uint16(1),
	858:  uint16(anon_sym_DOLLAR2),
	859:  uint16(214),
	860:  uint16(1),
	861:  uint16(aux_sym_image_name_token2),
	862:  uint16(5),
	863:  uint16(2),
	864:  uint16(sym_line_continuation),
	865:  uint16(sym_comment),
	866:  uint16(207),
	867:  uint16(3),
	868:  uint16(anon_sym_LF),
	869:  uint16(anon_sym_COLON),
	870:  uint16(anon_sym_AT),
	871:  uint16(25),
	872:  uint16(3),
	873:  uint16(sym__immediate_expansion),
	874:  uint16(sym__imm_expansion),
	875:  uint16(aux_sym_image_name_repeat1),
	876:  uint16(5),
	877:  uint16(221),
	878:  uint16(1),
	879:  uint16(anon_sym_COMMA),
	880:  uint16(27),
	881:  uint16(1),
	882:  uint16(aux_sym_mount_param_repeat1),
	883:  uint16(3),
	884:  uint16(2),
	885:  uint16(sym_line_continuation),
	886:  uint16(sym_comment),
	887:  uint16(219),
	888:  uint16(3),
	889:  uint16(sym_heredoc_marker),
	890:  uint16(anon_sym_DASH_DASH),
	891:  uint16(anon_sym_LBRACK),
	892:  uint16(217),
	893:  uint16(4),
	894:  uint16(aux_sym_path_token2),
	895:  uint16(aux_sym_shell_fragment_token2),
	896:  uint16(aux_sym_shell_fragment_token3),
	897:  uint16(aux_sym_shell_fragment_token4),
	898:  uint16(5),
	899:  uint16(221),
	900:  uint16(1),
	901:  uint16(anon_sym_COMMA),
	902:  uint16(28),
	903:  uint16(1),
	904:  uint16(aux_sym_mount_param_repeat1),
	905:  uint16(3),
	906:  uint16(2),
	907:  uint16(sym_line_continuation),
	908:  uint16(sym_comment),
	909:  uint16(225),
	910:  uint16(3),
	911:  uint16(sym_heredoc_marker),
	912:  uint16(anon_sym_DASH_DASH),
	913:  uint16(anon_sym_LBRACK),
	914:  uint16(223),
	915:  uint16(4),
	916:  uint16(aux_sym_path_token2),
	917:  uint16(aux_sym_shell_fragment_token2),
	918:  uint16(aux_sym_shell_fragment_token3),
	919:  uint16(aux_sym_shell_fragment_token4),
	920:  uint16(5),
	921:  uint16(231),
	922:  uint16(1),
	923:  uint16(anon_sym_COMMA),
	924:  uint16(28),
	925:  uint16(1),
	926:  uint16(aux_sym_mount_param_repeat1),
	927:  uint16(3),
	928:  uint16(2),
	929:  uint16(sym_line_continuation),
	930:  uint16(sym_comment),
	931:  uint16(229),
	932:  uint16(3),
	933:  uint16(sym_heredoc_marker),
	934:  uint16(anon_sym_DASH_DASH),
	935:  uint16(anon_sym_LBRACK),
	936:  uint16(227),
	937:  uint16(4),
	938:  uint16(aux_sym_path_token2),
	939:  uint16(aux_sym_shell_fragment_token2),
	940:  uint16(aux_sym_shell_fragment_token3),
	941:  uint16(aux_sym_shell_fragment_token4),
	942:  uint16(6),
	943:  uint16(236),
	944:  uint16(1),
	945:  uint16(aux_sym_from_instruction_token2),
	946:  uint16(238),
	947:  uint16(1),
	948:  uint16(anon_sym_DOLLAR2),
	949:  uint16(240),
	950:  uint16(1),
	951:  uint16(aux_sym_image_tag_token1),
	952:  uint16(5),
	953:  uint16(2),
	954:  uint16(sym_line_continuation),
	955:  uint16(sym_comment),
	956:  uint16(234),
	957:  uint16(2),
	958:  uint16(anon_sym_LF),
	959:  uint16(anon_sym_AT),
	960:  uint16(39),
	961:  uint16(3),
	962:  uint16(sym__immediate_expansion),
	963:  uint16(sym__imm_expansion),
	964:  uint16(aux_sym_image_tag_repeat1),
	965:  uint16(6),
	966:  uint16(242),
	967:  uint16(1),
	968:  uint16(anon_sym_DOLLAR2),
	969:  uint16(244),
	970:  uint16(1),
	971:  uint16(anon_sym_DQUOTE),
	972:  uint16(248),
	973:  uint16(1),
	974:  uint16(sym_double_quoted_escape_sequence),
	975:  uint16(5),
	976:  uint16(2),
	977:  uint16(sym_line_continuation),
	978:  uint16(sym_comment),
	979:  uint16(246),
	980:  uint16(2),
	981:  uint16(aux_sym_double_quoted_string_token1),
	982:  uint16(anon_sym_BSLASH),
	983:  uint16(37),
	984:  uint16(3),
	985:  uint16(sym__immediate_expansion),
	986:  uint16(sym__imm_expansion),
	987:  uint16(aux_sym_double_quoted_string_repeat1),
	988:  uint16(5),
	989:  uint16(252),
	990:  uint16(1),
	991:  uint16(aux_sym__immediate_user_name_or_group_fragment_token1),
	992:  uint16(254),
	993:  uint16(1),
	994:  uint16(anon_sym_DOLLAR2),
	995:  uint16(5),
	996:  uint16(2),
	997:  uint16(sym_line_continuation),
	998:  uint16(sym_comment),
	999:  uint16(250),
	1000: uint16(2),
	1001: uint16(anon_sym_LF),
	1002: uint16(anon_sym_COLON),
	1003: uint16(35),
	1004: uint16(4),
	1005: uint16(sym__immediate_user_name_or_group_fragment),
	1006: uint16(sym__immediate_expansion),
	1007: uint16(sym__imm_expansion),
	1008: uint16(aux_sym__user_name_or_group_repeat1),
	1009: uint16(3),
	1010: uint16(5),
	1011: uint16(2),
	1012: uint16(sym_line_continuation),
	1013: uint16(sym_comment),
	1014: uint16(191),
	1015: uint16(3),
	1016: uint16(sym_heredoc_marker),
	1017: uint16(sym_heredoc_nl),
	1018: uint16(anon_sym_LF),
	1019: uint16(196),
	1020: uint16(5),
	1021: uint16(aux_sym_path_token2),
	1022: uint16(aux_sym_shell_fragment_token2),
	1023: uint16(aux_sym_shell_fragment_token3),
	1024: uint16(aux_sym_shell_fragment_token4),
	1025: uint16(sym_required_line_continuation),
	1026: uint16(6),
	1027: uint16(242),
	1028: uint16(1),
	1029: uint16(anon_sym_DOLLAR2),
	1030: uint16(248),
	1031: uint16(1),
	1032: uint16(sym_double_quoted_escape_sequence),
	1033: uint16(256),
	1034: uint16(1),
	1035: uint16(anon_sym_DQUOTE),
	1036: uint16(5),
	1037: uint16(2),
	1038: uint16(sym_line_continuation),
	1039: uint16(sym_comment),
	1040: uint16(246),
	1041: uint16(2),
	1042: uint16(aux_sym_double_quoted_string_token1),
	1043: uint16(anon_sym_BSLASH),
	1044: uint16(37),
	1045: uint16(3),
	1046: uint16(sym__immediate_expansion),
	1047: uint16(sym__imm_expansion),
	1048: uint16(aux_sym_double_quoted_string_repeat1),
	1049: uint16(7),
	1050: uint16(258),
	1051: uint16(1),
	1052: uint16(anon_sym_LF),
	1053: uint16(260),
	1054: uint16(1),
	1055: uint16(aux_sym_label_pair_token1),
	1056: uint16(263),
	1057: uint16(1),
	1058: uint16(anon_sym_DQUOTE),
	1059: uint16(266),
	1060: uint16(1),
	1061: uint16(anon_sym_SQUOTE),
	1062: uint16(5),
	1063: uint16(2),
	1064: uint16(sym_line_continuation),
	1065: uint16(sym_comment),
	1066: uint16(34),
	1067: uint16(2),
	1068: uint16(sym_label_pair),
	1069: uint16(aux_sym_label_instruction_repeat1),
	1070: uint16(249),
	1071: uint16(2),
	1072: uint16(sym_double_quoted_string),
	1073: uint16(sym_single_quoted_string),
	1074: uint16(5),
	1075: uint16(254),
	1076: uint16(1),
	1077: uint16(anon_sym_DOLLAR2),
	1078: uint16(271),
	1079: uint16(1),
	1080: uint16(aux_sym__immediate_user_name_or_group_fragment_token1),
	1081: uint16(5),
	1082: uint16(2),
	1083: uint16(sym_line_continuation),
	1084: uint16(sym_comment),
	1085: uint16(269),
	1086: uint16(2),
	1087: uint16(anon_sym_LF),
	1088: uint16(anon_sym_COLON),
	1089: uint16(38),
	1090: uint16(4),
	1091: uint16(sym__immediate_user_name_or_group_fragment),
	1092: uint16(sym__immediate_expansion),
	1093: uint16(sym__imm_expansion),
	1094: uint16(aux_sym__user_name_or_group_repeat1),
	1095: uint16(6),
	1096: uint16(242),
	1097: uint16(1),
	1098: uint16(anon_sym_DOLLAR2),
	1099: uint16(273),
	1100: uint16(1),
	1101: uint16(anon_sym_DQUOTE),
	1102: uint16(277),
	1103: uint16(1),
	1104: uint16(sym_double_quoted_escape_sequence),
	1105: uint16(5),
	1106: uint16(2),
	1107: uint16(sym_line_continuation),
	1108: uint16(sym_comment),
	1109: uint16(275),
	1110: uint16(2),
	1111: uint16(aux_sym_double_quoted_string_token1),
	1112: uint16(anon_sym_BSLASH),
	1113: uint16(33),
	1114: uint16(3),
	1115: uint16(sym__immediate_expansion),
	1116: uint16(sym__imm_expansion),
	1117: uint16(aux_sym_double_quoted_string_repeat1),
	1118: uint16(6),
	1119: uint16(279),
	1120: uint16(1),
	1121: uint16(anon_sym_DOLLAR2),
	1122: uint16(282),
	1123: uint16(1),
	1124: uint16(anon_sym_DQUOTE),
	1125: uint16(287),
	1126: uint16(1),
	1127: uint16(sym_double_quoted_escape_sequence),
	1128: uint16(5),
	1129: uint16(2),
	1130: uint16(sym_line_continuation),
	1131: uint16(sym_comment),
	1132: uint16(284),
	1133: uint16(2),
	1134: uint16(aux_sym_double_quoted_string_token1),
	1135: uint16(anon_sym_BSLASH),
	1136: uint16(37),
	1137: uint16(3),
	1138: uint16(sym__immediate_expansion),
	1139: uint16(sym__imm_expansion),
	1140: uint16(aux_sym_double_quoted_string_repeat1),
	1141: uint16(5),
	1142: uint16(292),
	1143: uint16(1),
	1144: uint16(aux_sym__immediate_user_name_or_group_fragment_token1),
	1145: uint16(295),
	1146: uint16(1),
	1147: uint16(anon_sym_DOLLAR2),
	1148: uint16(5),
	1149: uint16(2),
	1150: uint16(sym_line_continuation),
	1151: uint16(sym_comment),
	1152: uint16(290),
	1153: uint16(2),
	1154: uint16(anon_sym_LF),
	1155: uint16(anon_sym_COLON),
	1156: uint16(38),
	1157: uint16(4),
	1158: uint16(sym__immediate_user_name_or_group_fragment),
	1159: uint16(sym__immediate_expansion),
	1160: uint16(sym__imm_expansion),
	1161: uint16(aux_sym__user_name_or_group_repeat1),
	1162: uint16(6),
	1163: uint16(300),
	1164: uint16(1),
	1165: uint16(aux_sym_from_instruction_token2),
	1166: uint16(302),
	1167: uint16(1),
	1168: uint16(anon_sym_DOLLAR2),
	1169: uint16(305),
	1170: uint16(1),
	1171: uint16(aux_sym_image_tag_token1),
	1172: uint16(5),
	1173: uint16(2),
	1174: uint16(sym_line_continuation),
	1175: uint16(sym_comment),
	1176: uint16(298),
	1177: uint16(2),
	1178: uint16(anon_sym_LF),
	1179: uint16(anon_sym_AT),
	1180: uint16(39),
	1181: uint16(3),
	1182: uint16(sym__immediate_expansion),
	1183: uint16(sym__imm_expansion),
	1184: uint16(aux_sym_image_tag_repeat1),
	1185: uint16(7),
	1186: uint16(308),
	1187: uint16(1),
	1188: uint16(anon_sym_LF),
	1189: uint16(310),
	1190: uint16(1),
	1191: uint16(aux_sym_label_pair_token1),
	1192: uint16(312),
	1193: uint16(1),
	1194: uint16(anon_sym_DQUOTE),
	1195: uint16(314),
	1196: uint16(1),
	1197: uint16(anon_sym_SQUOTE),
	1198: uint16(5),
	1199: uint16(2),
	1200: uint16(sym_line_continuation),
	1201: uint16(sym_comment),
	1202: uint16(34),
	1203: uint16(2),
	1204: uint16(sym_label_pair),
	1205: uint16(aux_sym_label_instruction_repeat1),
	1206: uint16(249),
	1207: uint16(2),
	1208: uint16(sym_double_quoted_string),
	1209: uint16(sym_single_quoted_string),
	1210: uint16(3),
	1211: uint16(3),
	1212: uint16(2),
	1213: uint16(sym_line_continuation),
	1214: uint16(sym_comment),
	1215: uint16(316),
	1216: uint16(4),
	1217: uint16(aux_sym_path_token2),
	1218: uint16(aux_sym_shell_fragment_token2),
	1219: uint16(aux_sym_shell_fragment_token3),
	1220: uint16(aux_sym_shell_fragment_token4),
	1221: uint16(318),
	1222: uint16(4),
	1223: uint16(sym_heredoc_marker),
	1224: uint16(anon_sym_DASH_DASH),
	1225: uint16(anon_sym_COMMA),
	1226: uint16(anon_sym_LBRACK),
	1227: uint16(3),
	1228: uint16(3),
	1229: uint16(2),
	1230: uint16(sym_line_continuation),
	1231: uint16(sym_comment),
	1232: uint16(227),
	1233: uint16(4),
	1234: uint16(aux_sym_path_token2),
	1235: uint16(aux_sym_shell_fragment_token2),
	1236: uint16(aux_sym_shell_fragment_token3),
	1237: uint16(aux_sym_shell_fragment_token4),
	1238: uint16(229),
	1239: uint16(4),
	1240: uint16(sym_heredoc_marker),
	1241: uint16(anon_sym_DASH_DASH),
	1242: uint16(anon_sym_COMMA),
	1243: uint16(anon_sym_LBRACK),
	1244: uint16(6),
	1245: uint16(142),
	1246: uint16(1),
	1247: uint16(sym_heredoc_marker),
	1248: uint16(185),
	1249: uint16(1),
	1250: uint16(anon_sym_LF),
	1251: uint16(189),
	1252: uint16(1),
	1253: uint16(sym_required_line_continuation),
	1254: uint16(45),
	1255: uint16(1),
	1256: uint16(aux_sym_shell_fragment_repeat1),
	1257: uint16(5),
	1258: uint16(2),
	1259: uint16(sym_line_continuation),
	1260: uint16(sym_comment),
	1261: uint16(320),
	1262: uint16(4),
	1263: uint16(aux_sym_path_token2),
	1264: uint16(aux_sym_shell_fragment_token2),
	1265: uint16(aux_sym_shell_fragment_token3),
	1266: uint16(aux_sym_shell_fragment_token4),
	1267: uint16(6),
	1268: uint16(322),
	1269: uint16(1),
	1270: uint16(anon_sym_LF),
	1271: uint16(324),
	1272: uint16(1),
	1273: uint16(aux_sym_path_token3),
	1274: uint16(326),
	1275: uint16(1),
	1276: uint16(anon_sym_DOLLAR2),
	1277: uint16(5),
	1278: uint16(2),
	1279: uint16(sym_line_continuation),
	1280: uint16(sym_comment),
	1281: uint16(328),
	1282: uint16(2),
	1283: uint16(sym_heredoc_nl),
	1284: uint16(sym__non_newline_whitespace),
	1285: uint16(46),
	1286: uint16(3),
	1287: uint16(sym__immediate_expansion),
	1288: uint16(sym__imm_expansion),
	1289: uint16(aux_sym_path_repeat1),
	1290: uint16(6),
	1291: uint16(191),
	1292: uint16(1),
	1293: uint16(anon_sym_LF),
	1294: uint16(196),
	1295: uint16(1),
	1296: uint16(sym_required_line_continuation),
	1297: uint16(333),
	1298: uint16(1),
	1299: uint16(sym_heredoc_marker),
	1300: uint16(45),
	1301: uint16(1),
	1302: uint16(aux_sym_shell_fragment_repeat1),
	1303: uint16(5),
	1304: uint16(2),
	1305: uint16(sym_line_continuation),
	1306: uint16(sym_comment),
	1307: uint16(330),
	1308: uint16(4),
	1309: uint16(aux_sym_path_token2),
	1310: uint16(aux_sym_shell_fragment_token2),
	1311: uint16(aux_sym_shell_fragment_token3),
	1312: uint16(aux_sym_shell_fragment_token4),
	1313: uint16(6),
	1314: uint16(326),
	1315: uint16(1),
	1316: uint16(anon_sym_DOLLAR2),
	1317: uint16(336),
	1318: uint16(1),
	1319: uint16(anon_sym_LF),
	1320: uint16(338),
	1321: uint16(1),
	1322: uint16(aux_sym_path_token3),
	1323: uint16(5),
	1324: uint16(2),
	1325: uint16(sym_line_continuation),
	1326: uint16(sym_comment),
	1327: uint16(340),
	1328: uint16(2),
	1329: uint16(sym_heredoc_nl),
	1330: uint16(sym__non_newline_whitespace),
	1331: uint16(49),
	1332: uint16(3),
	1333: uint16(sym__immediate_expansion),
	1334: uint16(sym__imm_expansion),
	1335: uint16(aux_sym_path_repeat1),
	1336: uint16(6),
	1337: uint16(118),
	1338: uint16(1),
	1339: uint16(anon_sym_DOLLAR2),
	1340: uint16(154),
	1341: uint16(1),
	1342: uint16(anon_sym_LF),
	1343: uint16(156),
	1344: uint16(1),
	1345: uint16(aux_sym__env_key_token1),
	1346: uint16(5),
	1347: uint16(2),
	1348: uint16(sym_line_continuation),
	1349: uint16(sym_comment),
	1350: uint16(342),
	1351: uint16(2),
	1352: uint16(aux_sym_unquoted_string_token1),
	1353: uint16(anon_sym_BSLASH2),
	1354: uint16(48),
	1355: uint16(3),
	1356: uint16(sym__immediate_expansion),
	1357: uint16(sym__imm_expansion),
	1358: uint16(aux_sym_unquoted_string_repeat1),
	1359: uint16(6),
	1360: uint16(167),
	1361: uint16(1),
	1362: uint16(anon_sym_LF),
	1363: uint16(172),
	1364: uint16(1),
	1365: uint16(aux_sym__env_key_token1),
	1366: uint16(344),
	1367: uint16(1),
	1368: uint16(anon_sym_DOLLAR2),
	1369: uint16(5),
	1370: uint16(2),
	1371: uint16(sym_line_continuation),
	1372: uint16(sym_comment),
	1373: uint16(347),
	1374: uint16(2),
	1375: uint16(aux_sym_unquoted_string_token1),
	1376: uint16(anon_sym_BSLASH2),
	1377: uint16(48),
	1378: uint16(3),
	1379: uint16(sym__immediate_expansion),
	1380: uint16(sym__imm_expansion),
	1381: uint16(aux_sym_unquoted_string_repeat1),
	1382: uint16(6),
	1383: uint16(350),
	1384: uint16(1),
	1385: uint16(anon_sym_LF),
	1386: uint16(352),
	1387: uint16(1),
	1388: uint16(aux_sym_path_token3),
	1389: uint16(355),
	1390: uint16(1),
	1391: uint16(anon_sym_DOLLAR2),
	1392: uint16(5),
	1393: uint16(2),
	1394: uint16(sym_line_continuation),
	1395: uint16(sym_comment),
	1396: uint16(358),
	1397: uint16(2),
	1398: uint16(sym_heredoc_nl),
	1399: uint16(sym__non_newline_whitespace),
	1400: uint16(49),
	1401: uint16(3),
	1402: uint16(sym__immediate_expansion),
	1403: uint16(sym__imm_expansion),
	1404: uint16(aux_sym_path_repeat1),
	1405: uint16(6),
	1406: uint16(242),
	1407: uint16(1),
	1408: uint16(anon_sym_DOLLAR2),
	1409: uint16(360),
	1410: uint16(1),
	1411: uint16(anon_sym_DQUOTE),
	1412: uint16(364),
	1413: uint16(1),
	1414: uint16(sym_double_quoted_escape_sequence),
	1415: uint16(5),
	1416: uint16(2),
	1417: uint16(sym_line_continuation),
	1418: uint16(sym_comment),
	1419: uint16(362),
	1420: uint16(2),
	1421: uint16(aux_sym_double_quoted_string_token1),
	1422: uint16(anon_sym_BSLASH),
	1423: uint16(51),
	1424: uint16(3),
	1425: uint16(sym__immediate_expansion),
	1426: uint16(sym__imm_expansion),
	1427: uint16(aux_sym_double_quoted_string_repeat1),
	1428: uint16(6),
	1429: uint16(242),
	1430: uint16(1),
	1431: uint16(anon_sym_DOLLAR2),
	1432: uint16(248),
	1433: uint16(1),
	1434: uint16(sym_double_quoted_escape_sequence),
	1435: uint16(366),
	1436: uint16(1),
	1437: uint16(anon_sym_DQUOTE),
	1438: uint16(5),
	1439: uint16(2),
	1440: uint16(sym_line_continuation),
	1441: uint16(sym_comment),
	1442: uint16(246),
	1443: uint16(2),
	1444: uint16(aux_sym_double_quoted_string_token1),
	1445: uint16(anon_sym_BSLASH),
	1446: uint16(37),
	1447: uint16(3),
	1448: uint16(sym__immediate_expansion),
	1449: uint16(sym__imm_expansion),
	1450: uint16(aux_sym_double_quoted_string_repeat1),
	1451: uint16(6),
	1452: uint16(242),
	1453: uint16(1),
	1454: uint16(anon_sym_DOLLAR2),
	1455: uint16(368),
	1456: uint16(1),
	1457: uint16(anon_sym_DQUOTE),
	1458: uint16(372),
	1459: uint16(1),
	1460: uint16(sym_double_quoted_escape_sequence),
	1461: uint16(5),
	1462: uint16(2),
	1463: uint16(sym_line_continuation),
	1464: uint16(sym_comment),
	1465: uint16(370),
	1466: uint16(2),
	1467: uint16(aux_sym_double_quoted_string_token1),
	1468: uint16(anon_sym_BSLASH),
	1469: uint16(30),
	1470: uint16(3),
	1471: uint16(sym__immediate_expansion),
	1472: uint16(sym__imm_expansion),
	1473: uint16(aux_sym_double_quoted_string_repeat1),
	1474: uint16(6),
	1475: uint16(374),
	1476: uint16(1),
	1477: uint16(anon_sym_LF),
	1478: uint16(376),
	1479: uint16(1),
	1480: uint16(aux_sym__immediate_user_name_or_group_fragment_token1),
	1481: uint16(379),
	1482: uint16(1),
	1483: uint16(anon_sym_DOLLAR2),
	1484: uint16(53),
	1485: uint16(1),
	1486: uint16(aux_sym__immediate_user_name_or_group),
	1487: uint16(5),
	1488: uint16(2),
	1489: uint16(sym_line_continuation),
	1490: uint16(sym_comment),
	1491: uint16(164),
	1492: uint16(3),
	1493: uint16(sym__immediate_user_name_or_group_fragment),
	1494: uint16(sym__immediate_expansion),
	1495: uint16(sym__imm_expansion),
	1496: uint16(9),
	1497: uint16(3),
	1498: uint16(1),
	1499: uint16(sym_line_continuation),
	1500: uint16(5),
	1501: uint16(1),
	1502: uint16(sym_comment),
	1503: uint16(140),
	1504: uint16(1),
	1505: uint16(anon_sym_LBRACK),
	1506: uint16(382),
	1507: uint16(1),
	1508: uint16(aux_sym_path_token1),
	1509: uint16(384),
	1510: uint16(1),
	1511: uint16(aux_sym_path_token2),
	1512: uint16(386),
	1513: uint16(1),
	1514: uint16(anon_sym_DOLLAR),
	1515: uint16(60),
	1516: uint16(1),
	1517: uint16(sym_expansion),
	1518: uint16(198),
	1519: uint16(1),
	1520: uint16(sym_path),
	1521: uint16(275),
	1522: uint16(1),
	1523: uint16(sym_json_string_array),
	1524: uint16(6),
	1525: uint16(388),
	1526: uint16(1),
	1527: uint16(anon_sym_LF),
	1528: uint16(390),
	1529: uint16(1),
	1530: uint16(aux_sym_from_instruction_token2),
	1531: uint16(392),
	1532: uint16(1),
	1533: uint16(anon_sym_DOLLAR2),
	1534: uint16(394),
	1535: uint16(1),
	1536: uint16(aux_sym_image_digest_token1),
	1537: uint16(5),
	1538: uint16(2),
	1539: uint16(sym_line_continuation),
	1540: uint16(sym_comment),
	1541: uint16(59),
	1542: uint16(3),
	1543: uint16(sym__immediate_expansion),
	1544: uint16(sym__imm_expansion),
	1545: uint16(aux_sym_image_digest_repeat1),
	1546: uint16(9),
	1547: uint16(3),
	1548: uint16(1),
	1549: uint16(sym_line_continuation),
	1550: uint16(5),
	1551: uint16(1),
	1552: uint16(sym_comment),
	1553: uint16(396),
	1554: uint16(1),
	1555: uint16(aux_sym_path_token1),
	1556: uint16(398),
	1557: uint16(1),
	1558: uint16(aux_sym_path_with_heredoc_token1),
	1559: uint16(400),
	1560: uint16(1),
	1561: uint16(anon_sym_DOLLAR),
	1562: uint16(402),
	1563: uint16(1),
	1564: uint16(sym_heredoc_marker),
	1565: uint16(44),
	1566: uint16(1),
	1567: uint16(sym_expansion),
	1568: uint16(66),
	1569: uint16(1),
	1570: uint16(aux_sym_add_instruction_repeat2),
	1571: uint16(100),
	1572: uint16(1),
	1573: uint16(sym_path_with_heredoc),
	1574: uint16(6),
	1575: uint16(254),
	1576: uint16(1),
	1577: uint16(anon_sym_DOLLAR2),
	1578: uint16(404),
	1579: uint16(1),
	1580: uint16(anon_sym_LF),
	1581: uint16(406),
	1582: uint16(1),
	1583: uint16(aux_sym__immediate_user_name_or_group_fragment_token1),
	1584: uint16(53),
	1585: uint16(1),
	1586: uint16(aux_sym__immediate_user_name_or_group),
	1587: uint16(5),
	1588: uint16(2),
	1589: uint16(sym_line_continuation),
	1590: uint16(sym_comment),
	1591: uint16(164),
	1592: uint16(3),
	1593: uint16(sym__immediate_user_name_or_group_fragment),
	1594: uint16(sym__immediate_expansion),
	1595: uint16(sym__imm_expansion),
	1596: uint16(6),
	1597: uint16(114),
	1598: uint16(1),
	1599: uint16(sym_heredoc_marker),
	1600: uint16(408),
	1601: uint16(1),
	1602: uint16(aux_sym_shell_fragment_token2),
	1603: uint16(22),
	1604: uint16(1),
	1605: uint16(aux_sym_shell_fragment_repeat1),
	1606: uint16(187),
	1607: uint16(1),
	1608: uint16(sym_shell_fragment),
	1609: uint16(3),
	1610: uint16(2),
	1611: uint16(sym_line_continuation),
	1612: uint16(sym_comment),
	1613: uint16(108),
	1614: uint16(3),
	1615: uint16(aux_sym_path_token2),
	1616: uint16(aux_sym_shell_fragment_token3),
	1617: uint16(aux_sym_shell_fragment_token4),
	1618: uint16(6),
	1619: uint16(410),
	1620: uint16(1),
	1621: uint16(anon_sym_LF),
	1622: uint16(412),
	1623: uint16(1),
	1624: uint16(aux_sym_from_instruction_token2),
	1625: uint16(414),
	1626: uint16(1),
	1627: uint16(anon_sym_DOLLAR2),
	1628: uint16(417),
	1629: uint16(1),
	1630: uint16(aux_sym_image_digest_token1),
	1631: uint16(5),
	1632: uint16(2),
	1633: uint16(sym_line_continuation),
	1634: uint16(sym_comment),
	1635: uint16(59),
	1636: uint16(3),
	1637: uint16(sym__immediate_expansion),
	1638: uint16(sym__imm_expansion),
	1639: uint16(aux_sym_image_digest_repeat1),
	1640: uint16(6),
	1641: uint16(420),
	1642: uint16(1),
	1643: uint16(anon_sym_LF),
	1644: uint16(422),
	1645: uint16(1),
	1646: uint16(aux_sym_path_token3),
	1647: uint16(424),
	1648: uint16(1),
	1649: uint16(anon_sym_DOLLAR2),
	1650: uint16(426),
	1651: uint16(1),
	1652: uint16(sym__non_newline_whitespace),
	1653: uint16(5),
	1654: uint16(2),
	1655: uint16(sym_line_continuation),
	1656: uint16(sym_comment),
	1657: uint16(70),
	1658: uint16(3),
	1659: uint16(sym__immediate_expansion),
	1660: uint16(sym__imm_expansion),
	1661: uint16(aux_sym_path_repeat1),
	1662: uint16(6),
	1663: uint16(428),
	1664: uint16(1),
	1665: uint16(aux_sym_label_pair_token1),
	1666: uint16(430),
	1667: uint16(1),
	1668: uint16(anon_sym_DQUOTE),
	1669: uint16(432),
	1670: uint16(1),
	1671: uint16(anon_sym_SQUOTE),
	1672: uint16(3),
	1673: uint16(2),
	1674: uint16(sym_line_continuation),
	1675: uint16(sym_comment),
	1676: uint16(40),
	1677: uint16(2),
	1678: uint16(sym_label_pair),
	1679: uint16(aux_sym_label_instruction_repeat1),
	1680: uint16(249),
	1681: uint16(2),
	1682: uint16(sym_double_quoted_string),
	1683: uint16(sym_single_quoted_string),
	1684: uint16(3),
	1685: uint16(5),
	1686: uint16(2),
	1687: uint16(sym_line_continuation),
	1688: uint16(sym_comment),
	1689: uint16(191),
	1690: uint16(2),
	1691: uint16(sym_heredoc_marker),
	1692: uint16(anon_sym_LF),
	1693: uint16(196),
	1694: uint16(5),
	1695: uint16(aux_sym_path_token2),
	1696: uint16(aux_sym_shell_fragment_token2),
	1697: uint16(aux_sym_shell_fragment_token3),
	1698: uint16(aux_sym_shell_fragment_token4),
	1699: uint16(sym_required_line_continuation),
	1700: uint16(6),
	1701: uint16(138),
	1702: uint16(1),
	1703: uint16(aux_sym_shell_fragment_token2),
	1704: uint16(142),
	1705: uint16(1),
	1706: uint16(sym_heredoc_marker),
	1707: uint16(43),
	1708: uint16(1),
	1709: uint16(aux_sym_shell_fragment_repeat1),
	1710: uint16(243),
	1711: uint16(1),
	1712: uint16(sym_shell_fragment),
	1713: uint16(3),
	1714: uint16(2),
	1715: uint16(sym_line_continuation),
	1716: uint16(sym_comment),
	1717: uint16(136),
	1718: uint16(3),
	1719: uint16(aux_sym_path_token2),
	1720: uint16(aux_sym_shell_fragment_token3),
	1721: uint16(aux_sym_shell_fragment_token4),
	1722: uint16(6),
	1723: uint16(3),
	1724: uint16(1),
	1725: uint16(sym_line_continuation),
	1726: uint16(5),
	1727: uint16(1),
	1728: uint16(sym_comment),
	1729: uint16(434),
	1730: uint16(1),
	1731: uint16(aux_sym_path_token1),
	1732: uint16(438),
	1733: uint16(1),
	1734: uint16(anon_sym_DASH_DASH),
	1735: uint16(64),
	1736: uint16(2),
	1737: uint16(sym_param),
	1738: uint16(aux_sym_add_instruction_repeat1),
	1739: uint16(436),
	1740: uint16(3),
	1741: uint16(sym_heredoc_marker),
	1742: uint16(aux_sym_path_with_heredoc_token1),
	1743: uint16(anon_sym_DOLLAR),
	1744: uint16(9),
	1745: uint16(3),
	1746: uint16(1),
	1747: uint16(sym_line_continuation),
	1748: uint16(5),
	1749: uint16(1),
	1750: uint16(sym_comment),
	1751: uint16(396),
	1752: uint16(1),
	1753: uint16(aux_sym_path_token1),
	1754: uint16(398),
	1755: uint16(1),
	1756: uint16(aux_sym_path_with_heredoc_token1),
	1757: uint16(400),
	1758: uint16(1),
	1759: uint16(anon_sym_DOLLAR),
	1760: uint16(402),
	1761: uint16(1),
	1762: uint16(sym_heredoc_marker),
	1763: uint16(44),
	1764: uint16(1),
	1765: uint16(sym_expansion),
	1766: uint16(66),
	1767: uint16(1),
	1768: uint16(aux_sym_add_instruction_repeat2),
	1769: uint16(121),
	1770: uint16(1),
	1771: uint16(sym_path_with_heredoc),
	1772: uint16(9),
	1773: uint16(3),
	1774: uint16(1),
	1775: uint16(sym_line_continuation),
	1776: uint16(5),
	1777: uint16(1),
	1778: uint16(sym_comment),
	1779: uint16(441),
	1780: uint16(1),
	1781: uint16(aux_sym_path_token1),
	1782: uint16(444),
	1783: uint16(1),
	1784: uint16(aux_sym_path_with_heredoc_token1),
	1785: uint16(447),
	1786: uint16(1),
	1787: uint16(anon_sym_DOLLAR),
	1788: uint16(450),
	1789: uint16(1),
	1790: uint16(sym_heredoc_marker),
	1791: uint16(66),
	1792: uint16(1),
	1793: uint16(aux_sym_add_instruction_repeat2),
	1794: uint16(84),
	1795: uint16(1),
	1796: uint16(sym_expansion),
	1797: uint16(299),
	1798: uint16(1),
	1799: uint16(sym_path_with_heredoc),
	1800: uint16(6),
	1801: uint16(350),
	1802: uint16(1),
	1803: uint16(anon_sym_LF),
	1804: uint16(358),
	1805: uint16(1),
	1806: uint16(sym__non_newline_whitespace),
	1807: uint16(453),
	1808: uint16(1),
	1809: uint16(aux_sym_path_token3),
	1810: uint16(456),
	1811: uint16(1),
	1812: uint16(anon_sym_DOLLAR2),
	1813: uint16(5),
	1814: uint16(2),
	1815: uint16(sym_line_continuation),
	1816: uint16(sym_comment),
	1817: uint16(67),
	1818: uint16(3),
	1819: uint16(sym__immediate_expansion),
	1820: uint16(sym__imm_expansion),
	1821: uint16(aux_sym_path_repeat1),
	1822: uint16(3),
	1823: uint16(3),
	1824: uint16(2),
	1825: uint16(sym_line_continuation),
	1826: uint16(sym_comment),
	1827: uint16(461),
	1828: uint16(3),
	1829: uint16(sym_heredoc_marker),
	1830: uint16(anon_sym_DASH_DASH),
	1831: uint16(anon_sym_LBRACK),
	1832: uint16(459),
	1833: uint16(4),
	1834: uint16(aux_sym_path_token2),
	1835: uint16(aux_sym_shell_fragment_token2),
	1836: uint16(aux_sym_shell_fragment_token3),
	1837: uint16(aux_sym_shell_fragment_token4),
	1838: uint16(9),
	1839: uint16(3),
	1840: uint16(1),
	1841: uint16(sym_line_continuation),
	1842: uint16(5),
	1843: uint16(1),
	1844: uint16(sym_comment),
	1845: uint16(396),
	1846: uint16(1),
	1847: uint16(aux_sym_path_token1),
	1848: uint16(398),
	1849: uint16(1),
	1850: uint16(aux_sym_path_with_heredoc_token1),
	1851: uint16(400),
	1852: uint16(1),
	1853: uint16(anon_sym_DOLLAR),
	1854: uint16(402),
	1855: uint16(1),
	1856: uint16(sym_heredoc_marker),
	1857: uint16(44),
	1858: uint16(1),
	1859: uint16(sym_expansion),
	1860: uint16(66),
	1861: uint16(1),
	1862: uint16(aux_sym_add_instruction_repeat2),
	1863: uint16(124),
	1864: uint16(1),
	1865: uint16(sym_path_with_heredoc),
	1866: uint16(6),
	1867: uint16(424),
	1868: uint16(1),
	1869: uint16(anon_sym_DOLLAR2),
	1870: uint16(463),
	1871: uint16(1),
	1872: uint16(anon_sym_LF),
	1873: uint16(465),
	1874: uint16(1),
	1875: uint16(aux_sym_path_token3),
	1876: uint16(467),
	1877: uint16(1),
	1878: uint16(sym__non_newline_whitespace),
	1879: uint16(5),
	1880: uint16(2),
	1881: uint16(sym_line_continuation),
	1882: uint16(sym_comment),
	1883: uint16(67),
	1884: uint16(3),
	1885: uint16(sym__immediate_expansion),
	1886: uint16(sym__imm_expansion),
	1887: uint16(aux_sym_path_repeat1),
	1888: uint16(9),
	1889: uint16(3),
	1890: uint16(1),
	1891: uint16(sym_line_continuation),
	1892: uint16(5),
	1893: uint16(1),
	1894: uint16(sym_comment),
	1895: uint16(469),
	1896: uint16(1),
	1897: uint16(anon_sym_DOLLAR),
	1898: uint16(471),
	1899: uint16(1),
	1900: uint16(aux_sym_image_name_token1),
	1901: uint16(473),
	1902: uint16(1),
	1903: uint16(anon_sym_DASH_DASH),
	1904: uint16(21),
	1905: uint16(1),
	1906: uint16(sym_expansion),
	1907: uint16(83),
	1908: uint16(1),
	1909: uint16(sym_image_name),
	1910: uint16(105),
	1911: uint16(1),
	1912: uint16(sym_param),
	1913: uint16(237),
	1914: uint16(1),
	1915: uint16(sym_image_spec),
	1916: uint16(3),
	1917: uint16(5),
	1918: uint16(2),
	1919: uint16(sym_line_continuation),
	1920: uint16(sym_comment),
	1921: uint16(475),
	1922: uint16(2),
	1923: uint16(anon_sym_LF),
	1924: uint16(anon_sym_DOLLAR2),
	1925: uint16(477),
	1926: uint16(5),
	1927: uint16(aux_sym_label_pair_token1),
	1928: uint16(anon_sym_DQUOTE),
	1929: uint16(anon_sym_SQUOTE),
	1930: uint16(aux_sym_unquoted_string_token1),
	1931: uint16(anon_sym_BSLASH2),
	1932: uint16(3),
	1933: uint16(5),
	1934: uint16(2),
	1935: uint16(sym_line_continuation),
	1936: uint16(sym_comment),
	1937: uint16(479),
	1938: uint16(2),
	1939: uint16(anon_sym_LF),
	1940: uint16(anon_sym_DOLLAR2),
	1941: uint16(481),
	1942: uint16(5),
	1943: uint16(aux_sym_label_pair_token1),
	1944: uint16(anon_sym_DQUOTE),
	1945: uint16(anon_sym_SQUOTE),
	1946: uint16(aux_sym_unquoted_string_token1),
	1947: uint16(anon_sym_BSLASH2),
	1948: uint16(9),
	1949: uint16(3),
	1950: uint16(1),
	1951: uint16(sym_line_continuation),
	1952: uint16(5),
	1953: uint16(1),
	1954: uint16(sym_comment),
	1955: uint16(396),
	1956: uint16(1),
	1957: uint16(aux_sym_path_token1),
	1958: uint16(398),
	1959: uint16(1),
	1960: uint16(aux_sym_path_with_heredoc_token1),
	1961: uint16(400),
	1962: uint16(1),
	1963: uint16(anon_sym_DOLLAR),
	1964: uint16(402),
	1965: uint16(1),
	1966: uint16(sym_heredoc_marker),
	1967: uint16(44),
	1968: uint16(1),
	1969: uint16(sym_expansion),
	1970: uint16(66),
	1971: uint16(1),
	1972: uint16(aux_sym_add_instruction_repeat2),
	1973: uint16(98),
	1974: uint16(1),
	1975: uint16(sym_path_with_heredoc),
	1976: uint16(3),
	1977: uint16(5),
	1978: uint16(2),
	1979: uint16(sym_line_continuation),
	1980: uint16(sym_comment),
	1981: uint16(477),
	1982: uint16(2),
	1983: uint16(aux_sym_from_instruction_token2),
	1984: uint16(aux_sym_image_name_token2),
	1985: uint16(475),
	1986: uint16(4),
	1987: uint16(anon_sym_LF),
	1988: uint16(anon_sym_COLON),
	1989: uint16(anon_sym_DOLLAR2),
	1990: uint16(anon_sym_AT),
	1991: uint16(5),
	1992: uint16(426),
	1993: uint16(1),
	1994: uint16(anon_sym_LF),
	1995: uint16(483),
	1996: uint16(1),
	1997: uint16(aux_sym_path_token3),
	1998: uint16(485),
	1999: uint16(1),
	2000: uint16(anon_sym_DOLLAR2),
	2001: uint16(5),
	2002: uint16(2),
	2003: uint16(sym_line_continuation),
	2004: uint16(sym_comment),
	2005: uint16(81),
	2006: uint16(3),
	2007: uint16(sym__immediate_expansion),
	2008: uint16(sym__imm_expansion),
	2009: uint16(aux_sym_path_repeat1),
	2010: uint16(5),
	2011: uint16(487),
	2012: uint16(1),
	2013: uint16(anon_sym_LF),
	2014: uint16(489),
	2015: uint16(1),
	2016: uint16(aux_sym__stopsignal_value_token2),
	2017: uint16(492),
	2018: uint16(1),
	2019: uint16(anon_sym_DOLLAR2),
	2020: uint16(5),
	2021: uint16(2),
	2022: uint16(sym_line_continuation),
	2023: uint16(sym_comment),
	2024: uint16(77),
	2025: uint16(3),
	2026: uint16(sym__immediate_expansion),
	2027: uint16(sym__imm_expansion),
	2028: uint16(aux_sym__stopsignal_value_repeat1),
	2029: uint16(3),
	2030: uint16(5),
	2031: uint16(2),
	2032: uint16(sym_line_continuation),
	2033: uint16(sym_comment),
	2034: uint16(481),
	2035: uint16(2),
	2036: uint16(aux_sym_from_instruction_token2),
	2037: uint16(aux_sym_image_name_token2),
	2038: uint16(479),
	2039: uint16(4),
	2040: uint16(anon_sym_LF),
	2041: uint16(anon_sym_COLON),
	2042: uint16(anon_sym_DOLLAR2),
	2043: uint16(anon_sym_AT),
	2044: uint16(5),
	2045: uint16(495),
	2046: uint16(1),
	2047: uint16(anon_sym_LF),
	2048: uint16(497),
	2049: uint16(1),
	2050: uint16(anon_sym_DOLLAR2),
	2051: uint16(499),
	2052: uint16(1),
	2053: uint16(aux_sym_image_alias_token2),
	2054: uint16(5),
	2055: uint16(2),
	2056: uint16(sym_line_continuation),
	2057: uint16(sym_comment),
	2058: uint16(85),
	2059: uint16(3),
	2060: uint16(sym__immediate_expansion),
	2061: uint16(sym__imm_expansion),
	2062: uint16(aux_sym_image_alias_repeat1),
	2063: uint16(3),
	2064: uint16(5),
	2065: uint16(2),
	2066: uint16(sym_line_continuation),
	2067: uint16(sym_comment),
	2068: uint16(503),
	2069: uint16(2),
	2070: uint16(aux_sym_from_instruction_token2),
	2071: uint16(aux_sym_image_name_token2),
	2072: uint16(501),
	2073: uint16(4),
	2074: uint16(anon_sym_LF),
	2075: uint16(anon_sym_COLON),
	2076: uint16(anon_sym_DOLLAR2),
	2077: uint16(anon_sym_AT),
	2078: uint16(5),
	2079: uint16(467),
	2080: uint16(1),
	2081: uint16(anon_sym_LF),
	2082: uint16(485),
	2083: uint16(1),
	2084: uint16(anon_sym_DOLLAR2),
	2085: uint16(505),
	2086: uint16(1),
	2087: uint16(aux_sym_path_token3),
	2088: uint16(5),
	2089: uint16(2),
	2090: uint16(sym_line_continuation),
	2091: uint16(sym_comment),
	2092: uint16(86),
	2093: uint16(3),
	2094: uint16(sym__immediate_expansion),
	2095: uint16(sym__imm_expansion),
	2096: uint16(aux_sym_path_repeat1),
	2097: uint16(5),
	2098: uint16(358),
	2099: uint16(1),
	2100: uint16(sym__non_newline_whitespace),
	2101: uint16(507),
	2102: uint16(1),
	2103: uint16(aux_sym_path_token3),
	2104: uint16(510),
	2105: uint16(1),
	2106: uint16(anon_sym_DOLLAR2),
	2107: uint16(5),
	2108: uint16(2),
	2109: uint16(sym_line_continuation),
	2110: uint16(sym_comment),
	2111: uint16(82),
	2112: uint16(3),
	2113: uint16(sym__immediate_expansion),
	2114: uint16(sym__imm_expansion),
	2115: uint16(aux_sym_path_repeat1),
	2116: uint16(7),
	2117: uint16(513),
	2118: uint16(1),
	2119: uint16(anon_sym_LF),
	2120: uint16(515),
	2121: uint16(1),
	2122: uint16(aux_sym_from_instruction_token2),
	2123: uint16(517),
	2124: uint16(1),
	2125: uint16(anon_sym_COLON),
	2126: uint16(519),
	2127: uint16(1),
	2128: uint16(anon_sym_AT),
	2129: uint16(143),
	2130: uint16(1),
	2131: uint16(sym_image_tag),
	2132: uint16(245),
	2133: uint16(1),
	2134: uint16(sym_image_digest),
	2135: uint16(5),
	2136: uint16(2),
	2137: uint16(sym_line_continuation),
	2138: uint16(sym_comment),
	2139: uint16(5),
	2140: uint16(328),
	2141: uint16(1),
	2142: uint16(sym__non_newline_whitespace),
	2143: uint16(521),
	2144: uint16(1),
	2145: uint16(aux_sym_path_token3),
	2146: uint16(523),
	2147: uint16(1),
	2148: uint16(anon_sym_DOLLAR2),
	2149: uint16(5),
	2150: uint16(2),
	2151: uint16(sym_line_continuation),
	2152: uint16(sym_comment),
	2153: uint16(93),
	2154: uint16(3),
	2155: uint16(sym__immediate_expansion),
	2156: uint16(sym__imm_expansion),
	2157: uint16(aux_sym_path_repeat1),
	2158: uint16(5),
	2159: uint16(525),
	2160: uint16(1),
	2161: uint16(anon_sym_LF),
	2162: uint16(527),
	2163: uint16(1),
	2164: uint16(anon_sym_DOLLAR2),
	2165: uint16(530),
	2166: uint16(1),
	2167: uint16(aux_sym_image_alias_token2),
	2168: uint16(5),
	2169: uint16(2),
	2170: uint16(sym_line_continuation),
	2171: uint16(sym_comment),
	2172: uint16(85),
	2173: uint16(3),
	2174: uint16(sym__immediate_expansion),
	2175: uint16(sym__imm_expansion),
	2176: uint16(aux_sym_image_alias_repeat1),
	2177: uint16(5),
	2178: uint16(358),
	2179: uint16(1),
	2180: uint16(anon_sym_LF),
	2181: uint16(533),
	2182: uint16(1),
	2183: uint16(aux_sym_path_token3),
	2184: uint16(536),
	2185: uint16(1),
	2186: uint16(anon_sym_DOLLAR2),
	2187: uint16(5),
	2188: uint16(2),
	2189: uint16(sym_line_continuation),
	2190: uint16(sym_comment),
	2191: uint16(86),
	2192: uint16(3),
	2193: uint16(sym__immediate_expansion),
	2194: uint16(sym__imm_expansion),
	2195: uint16(aux_sym_path_repeat1),
	2196: uint16(5),
	2197: uint16(539),
	2198: uint16(1),
	2199: uint16(anon_sym_LF),
	2200: uint16(541),
	2201: uint16(1),
	2202: uint16(anon_sym_DOLLAR),
	2203: uint16(543),
	2204: uint16(1),
	2205: uint16(aux_sym_expose_port_token1),
	2206: uint16(5),
	2207: uint16(2),
	2208: uint16(sym_line_continuation),
	2209: uint16(sym_comment),
	2210: uint16(90),
	2211: uint16(3),
	2212: uint16(sym_expansion),
	2213: uint16(sym_expose_port),
	2214: uint16(aux_sym_expose_instruction_repeat1),
	2215: uint16(6),
	2216: uint16(13),
	2217: uint16(1),
	2218: uint16(aux_sym_cmd_instruction_token1),
	2219: uint16(545),
	2220: uint16(1),
	2221: uint16(anon_sym_NONE),
	2222: uint16(547),
	2223: uint16(1),
	2224: uint16(anon_sym_DASH_DASH),
	2225: uint16(257),
	2226: uint16(1),
	2227: uint16(sym_cmd_instruction),
	2228: uint16(3),
	2229: uint16(2),
	2230: uint16(sym_line_continuation),
	2231: uint16(sym_comment),
	2232: uint16(97),
	2233: uint16(2),
	2234: uint16(sym_param),
	2235: uint16(aux_sym_add_instruction_repeat1),
	2236: uint16(5),
	2237: uint16(497),
	2238: uint16(1),
	2239: uint16(anon_sym_DOLLAR2),
	2240: uint16(549),
	2241: uint16(1),
	2242: uint16(anon_sym_LF),
	2243: uint16(551),
	2244: uint16(1),
	2245: uint16(aux_sym_image_alias_token2),
	2246: uint16(5),
	2247: uint16(2),
	2248: uint16(sym_line_continuation),
	2249: uint16(sym_comment),
	2250: uint16(79),
	2251: uint16(3),
	2252: uint16(sym__immediate_expansion),
	2253: uint16(sym__imm_expansion),
	2254: uint16(aux_sym_image_alias_repeat1),
	2255: uint16(5),
	2256: uint16(553),
	2257: uint16(1),
	2258: uint16(anon_sym_LF),
	2259: uint16(555),
	2260: uint16(1),
	2261: uint16(anon_sym_DOLLAR),
	2262: uint16(558),
	2263: uint16(1),
	2264: uint16(aux_sym_expose_port_token1),
	2265: uint16(5),
	2266: uint16(2),
	2267: uint16(sym_line_continuation),
	2268: uint16(sym_comment),
	2269: uint16(90),
	2270: uint16(3),
	2271: uint16(sym_expansion),
	2272: uint16(sym_expose_port),
	2273: uint16(aux_sym_expose_instruction_repeat1),
	2274: uint16(5),
	2275: uint16(561),
	2276: uint16(1),
	2277: uint16(anon_sym_LF),
	2278: uint16(563),
	2279: uint16(1),
	2280: uint16(aux_sym__stopsignal_value_token2),
	2281: uint16(565),
	2282: uint16(1),
	2283: uint16(anon_sym_DOLLAR2),
	2284: uint16(5),
	2285: uint16(2),
	2286: uint16(sym_line_continuation),
	2287: uint16(sym_comment),
	2288: uint16(94),
	2289: uint16(3),
	2290: uint16(sym__immediate_expansion),
	2291: uint16(sym__imm_expansion),
	2292: uint16(aux_sym__stopsignal_value_repeat1),
	2293: uint16(5),
	2294: uint16(254),
	2295: uint16(1),
	2296: uint16(anon_sym_DOLLAR2),
	2297: uint16(406),
	2298: uint16(1),
	2299: uint16(aux_sym__immediate_user_name_or_group_fragment_token1),
	2300: uint16(57),
	2301: uint16(1),
	2302: uint16(aux_sym__immediate_user_name_or_group),
	2303: uint16(3),
	2304: uint16(2),
	2305: uint16(sym_line_continuation),
	2306: uint16(sym_comment),
	2307: uint16(164),
	2308: uint16(3),
	2309: uint16(sym__immediate_user_name_or_group_fragment),
	2310: uint16(sym__immediate_expansion),
	2311: uint16(sym__imm_expansion),
	2312: uint16(5),
	2313: uint16(340),
	2314: uint16(1),
	2315: uint16(sym__non_newline_whitespace),
	2316: uint16(523),
	2317: uint16(1),
	2318: uint16(anon_sym_DOLLAR2),
	2319: uint16(567),
	2320: uint16(1),
	2321: uint16(aux_sym_path_token3),
	2322: uint16(5),
	2323: uint16(2),
	2324: uint16(sym_line_continuation),
	2325: uint16(sym_comment),
	2326: uint16(82),
	2327: uint16(3),
	2328: uint16(sym__immediate_expansion),
	2329: uint16(sym__imm_expansion),
	2330: uint16(aux_sym_path_repeat1),
	2331: uint16(5),
	2332: uint16(565),
	2333: uint16(1),
	2334: uint16(anon_sym_DOLLAR2),
	2335: uint16(569),
	2336: uint16(1),
	2337: uint16(anon_sym_LF),
	2338: uint16(571),
	2339: uint16(1),
	2340: uint16(aux_sym__stopsignal_value_token2),
	2341: uint16(5),
	2342: uint16(2),
	2343: uint16(sym_line_continuation),
	2344: uint16(sym_comment),
	2345: uint16(77),
	2346: uint16(3),
	2347: uint16(sym__immediate_expansion),
	2348: uint16(sym__imm_expansion),
	2349: uint16(aux_sym__stopsignal_value_repeat1),
	2350: uint16(3),
	2351: uint16(5),
	2352: uint16(2),
	2353: uint16(sym_line_continuation),
	2354: uint16(sym_comment),
	2355: uint16(477),
	2356: uint16(2),
	2357: uint16(aux_sym_from_instruction_token2),
	2358: uint16(aux_sym_image_tag_token1),
	2359: uint16(475),
	2360: uint16(3),
	2361: uint16(anon_sym_LF),
	2362: uint16(anon_sym_DOLLAR2),
	2363: uint16(anon_sym_AT),
	2364: uint16(5),
	2365: uint16(573),
	2366: uint16(1),
	2367: uint16(anon_sym_LF),
	2368: uint16(575),
	2369: uint16(1),
	2370: uint16(aux_sym__env_key_token1),
	2371: uint16(271),
	2372: uint16(1),
	2373: uint16(sym__env_key),
	2374: uint16(5),
	2375: uint16(2),
	2376: uint16(sym_line_continuation),
	2377: uint16(sym_comment),
	2378: uint16(96),
	2379: uint16(2),
	2380: uint16(sym_env_pair),
	2381: uint16(aux_sym_env_instruction_repeat1),
	2382: uint16(5),
	2383: uint16(13),
	2384: uint16(1),
	2385: uint16(aux_sym_cmd_instruction_token1),
	2386: uint16(547),
	2387: uint16(1),
	2388: uint16(anon_sym_DASH_DASH),
	2389: uint16(282),
	2390: uint16(1),
	2391: uint16(sym_cmd_instruction),
	2392: uint16(3),
	2393: uint16(2),
	2394: uint16(sym_line_continuation),
	2395: uint16(sym_comment),
	2396: uint16(153),
	2397: uint16(2),
	2398: uint16(sym_param),
	2399: uint16(aux_sym_add_instruction_repeat1),
	2400: uint16(5),
	2401: uint16(578),
	2402: uint16(1),
	2403: uint16(anon_sym_LF),
	2404: uint16(580),
	2405: uint16(1),
	2406: uint16(sym__non_newline_whitespace),
	2407: uint16(582),
	2408: uint16(1),
	2409: uint16(sym_heredoc_nl),
	2410: uint16(5),
	2411: uint16(2),
	2412: uint16(sym_line_continuation),
	2413: uint16(sym_comment),
	2414: uint16(154),
	2415: uint16(2),
	2416: uint16(sym_heredoc_block),
	2417: uint16(aux_sym_run_instruction_repeat2),
	2418: uint16(4),
	2419: uint16(584),
	2420: uint16(1),
	2421: uint16(anon_sym_LF),
	2422: uint16(5),
	2423: uint16(2),
	2424: uint16(sym_line_continuation),
	2425: uint16(sym_comment),
	2426: uint16(586),
	2427: uint16(2),
	2428: uint16(anon_sym_DOLLAR),
	2429: uint16(aux_sym_expose_port_token1),
	2430: uint16(588),
	2431: uint16(2),
	2432: uint16(anon_sym_SLASHtcp),
	2433: uint16(anon_sym_SLASHudp),
	2434: uint16(5),
	2435: uint16(580),
	2436: uint16(1),
	2437: uint16(sym__non_newline_whitespace),
	2438: uint16(582),
	2439: uint16(1),
	2440: uint16(sym_heredoc_nl),
	2441: uint16(590),
	2442: uint16(1),
	2443: uint16(anon_sym_LF),
	2444: uint16(5),
	2445: uint16(2),
	2446: uint16(sym_line_continuation),
	2447: uint16(sym_comment),
	2448: uint16(158),
	2449: uint16(2),
	2450: uint16(sym_heredoc_block),
	2451: uint16(aux_sym_run_instruction_repeat2),
	2452: uint16(7),
	2453: uint16(3),
	2454: uint16(1),
	2455: uint16(sym_line_continuation),
	2456: uint16(5),
	2457: uint16(1),
	2458: uint16(sym_comment),
	2459: uint16(382),
	2460: uint16(1),
	2461: uint16(aux_sym_path_token1),
	2462: uint16(384),
	2463: uint16(1),
	2464: uint16(aux_sym_path_token2),
	2465: uint16(386),
	2466: uint16(1),
	2467: uint16(anon_sym_DOLLAR),
	2468: uint16(60),
	2469: uint16(1),
	2470: uint16(sym_expansion),
	2471: uint16(241),
	2472: uint16(1),
	2473: uint16(sym_path),
	2474: uint16(5),
	2475: uint16(592),
	2476: uint16(1),
	2477: uint16(aux_sym__env_key_token1),
	2478: uint16(242),
	2479: uint16(1),
	2480: uint16(sym__env_key),
	2481: uint16(278),
	2482: uint16(1),
	2483: uint16(sym__spaced_env_pair),
	2484: uint16(3),
	2485: uint16(2),
	2486: uint16(sym_line_continuation),
	2487: uint16(sym_comment),
	2488: uint16(103),
	2489: uint16(2),
	2490: uint16(sym_env_pair),
	2491: uint16(aux_sym_env_instruction_repeat1),
	2492: uint16(5),
	2493: uint16(594),
	2494: uint16(1),
	2495: uint16(anon_sym_LF),
	2496: uint16(596),
	2497: uint16(1),
	2498: uint16(aux_sym__env_key_token1),
	2499: uint16(271),
	2500: uint16(1),
	2501: uint16(sym__env_key),
	2502: uint16(5),
	2503: uint16(2),
	2504: uint16(sym_line_continuation),
	2505: uint16(sym_comment),
	2506: uint16(96),
	2507: uint16(2),
	2508: uint16(sym_env_pair),
	2509: uint16(aux_sym_env_instruction_repeat1),
	2510: uint16(4),
	2511: uint16(598),
	2512: uint16(1),
	2513: uint16(anon_sym_DOLLAR),
	2514: uint16(600),
	2515: uint16(1),
	2516: uint16(aux_sym_expose_port_token1),
	2517: uint16(3),
	2518: uint16(2),
	2519: uint16(sym_line_continuation),
	2520: uint16(sym_comment),
	2521: uint16(87),
	2522: uint16(3),
	2523: uint16(sym_expansion),
	2524: uint16(sym_expose_port),
	2525: uint16(aux_sym_expose_instruction_repeat1),
	2526: uint16(7),
	2527: uint16(3),
	2528: uint16(1),
	2529: uint16(sym_line_continuation),
	2530: uint16(5),
	2531: uint16(1),
	2532: uint16(sym_comment),
	2533: uint16(469),
	2534: uint16(1),
	2535: uint16(anon_sym_DOLLAR),
	2536: uint16(471),
	2537: uint16(1),
	2538: uint16(aux_sym_image_name_token1),
	2539: uint16(21),
	2540: uint16(1),
	2541: uint16(sym_expansion),
	2542: uint16(83),
	2543: uint16(1),
	2544: uint16(sym_image_name),
	2545: uint16(246),
	2546: uint16(1),
	2547: uint16(sym_image_spec),
	2548: uint16(5),
	2549: uint16(604),
	2550: uint16(1),
	2551: uint16(anon_sym_SQUOTE),
	2552: uint16(606),
	2553: uint16(1),
	2554: uint16(sym_single_quoted_escape_sequence),
	2555: uint16(116),
	2556: uint16(1),
	2557: uint16(aux_sym_single_quoted_string_repeat1),
	2558: uint16(5),
	2559: uint16(2),
	2560: uint16(sym_line_continuation),
	2561: uint16(sym_comment),
	2562: uint16(602),
	2563: uint16(2),
	2564: uint16(anon_sym_BSLASH),
	2565: uint16(aux_sym_single_quoted_string_token1),
	2566: uint16(5),
	2567: uint16(3),
	2568: uint16(1),
	2569: uint16(sym_line_continuation),
	2570: uint16(5),
	2571: uint16(1),
	2572: uint16(sym_comment),
	2573: uint16(238),
	2574: uint16(1),
	2575: uint16(anon_sym_DOLLAR2),
	2576: uint16(608),
	2577: uint16(1),
	2578: uint16(aux_sym_image_tag_token1),
	2579: uint16(29),
	2580: uint16(3),
	2581: uint16(sym__immediate_expansion),
	2582: uint16(sym__imm_expansion),
	2583: uint16(aux_sym_image_tag_repeat1),
	2584: uint16(4),
	2585: uint16(392),
	2586: uint16(1),
	2587: uint16(anon_sym_DOLLAR2),
	2588: uint16(610),
	2589: uint16(1),
	2590: uint16(aux_sym_image_digest_token1),
	2591: uint16(3),
	2592: uint16(2),
	2593: uint16(sym_line_continuation),
	2594: uint16(sym_comment),
	2595: uint16(55),
	2596: uint16(3),
	2597: uint16(sym__immediate_expansion),
	2598: uint16(sym__imm_expansion),
	2599: uint16(aux_sym_image_digest_repeat1),
	2600: uint16(3),
	2601: uint16(5),
	2602: uint16(2),
	2603: uint16(sym_line_continuation),
	2604: uint16(sym_comment),
	2605: uint16(475),
	2606: uint16(2),
	2607: uint16(anon_sym_DOLLAR2),
	2608: uint16(sym_double_quoted_escape_sequence),
	2609: uint16(477),
	2610: uint16(3),
	2611: uint16(anon_sym_DQUOTE),
	2612: uint16(aux_sym_double_quoted_string_token1),
	2613: uint16(anon_sym_BSLASH),
	2614: uint16(7),
	2615: uint16(3),
	2616: uint16(1),
	2617: uint16(sym_line_continuation),
	2618: uint16(5),
	2619: uint16(1),
	2620: uint16(sym_comment),
	2621: uint16(612),
	2622: uint16(1),
	2623: uint16(aux_sym_path_token1),
	2624: uint16(614),
	2625: uint16(1),
	2626: uint16(aux_sym_path_token2),
	2627: uint16(616),
	2628: uint16(1),
	2629: uint16(anon_sym_DOLLAR),
	2630: uint16(76),
	2631: uint16(1),
	2632: uint16(sym_expansion),
	2633: uint16(250),
	2634: uint16(1),
	2635: uint16(sym_path),
	2636: uint16(4),
	2637: uint16(3),
	2638: uint16(1),
	2639: uint16(sym_line_continuation),
	2640: uint16(5),
	2641: uint16(1),
	2642: uint16(sym_comment),
	2643: uint16(459),
	2644: uint16(1),
	2645: uint16(aux_sym_path_token1),
	2646: uint16(461),
	2647: uint16(4),
	2648: uint16(sym_heredoc_marker),
	2649: uint16(aux_sym_path_with_heredoc_token1),
	2650: uint16(anon_sym_DOLLAR),
	2651: uint16(anon_sym_DASH_DASH),
	2652: uint16(5),
	2653: uint16(620),
	2654: uint16(1),
	2655: uint16(anon_sym_SQUOTE),
	2656: uint16(622),
	2657: uint16(1),
	2658: uint16(sym_single_quoted_escape_sequence),
	2659: uint16(106),
	2660: uint16(1),
	2661: uint16(aux_sym_single_quoted_string_repeat1),
	2662: uint16(5),
	2663: uint16(2),
	2664: uint16(sym_line_continuation),
	2665: uint16(sym_comment),
	2666: uint16(618),
	2667: uint16(2),
	2668: uint16(anon_sym_BSLASH),
	2669: uint16(aux_sym_single_quoted_string_token1),
	2670: uint16(3),
	2671: uint16(5),
	2672: uint16(2),
	2673: uint16(sym_line_continuation),
	2674: uint16(sym_comment),
	2675: uint16(503),
	2676: uint16(2),
	2677: uint16(anon_sym_LF),
	2678: uint16(aux_sym_path_token3),
	2679: uint16(501),
	2680: uint16(3),
	2681: uint16(sym_heredoc_nl),
	2682: uint16(anon_sym_DOLLAR2),
	2683: uint16(sym__non_newline_whitespace),
	2684: uint16(3),
	2685: uint16(5),
	2686: uint16(2),
	2687: uint16(sym_line_continuation),
	2688: uint16(sym_comment),
	2689: uint16(479),
	2690: uint16(2),
	2691: uint16(anon_sym_DOLLAR2),
	2692: uint16(sym_double_quoted_escape_sequence),
	2693: uint16(481),
	2694: uint16(3),
	2695: uint16(anon_sym_DQUOTE),
	2696: uint16(aux_sym_double_quoted_string_token1),
	2697: uint16(anon_sym_BSLASH),
	2698: uint16(3),
	2699: uint16(5),
	2700: uint16(2),
	2701: uint16(sym_line_continuation),
	2702: uint16(sym_comment),
	2703: uint16(481),
	2704: uint16(2),
	2705: uint16(anon_sym_LF),
	2706: uint16(aux_sym_path_token3),
	2707: uint16(479),
	2708: uint16(3),
	2709: uint16(sym_heredoc_nl),
	2710: uint16(anon_sym_DOLLAR2),
	2711: uint16(sym__non_newline_whitespace),
	2712: uint16(5),
	2713: uint16(627),
	2714: uint16(1),
	2715: uint16(anon_sym_SQUOTE),
	2716: uint16(629),
	2717: uint16(1),
	2718: uint16(sym_single_quoted_escape_sequence),
	2719: uint16(116),
	2720: uint16(1),
	2721: uint16(aux_sym_single_quoted_string_repeat1),
	2722: uint16(5),
	2723: uint16(2),
	2724: uint16(sym_line_continuation),
	2725: uint16(sym_comment),
	2726: uint16(624),
	2727: uint16(2),
	2728: uint16(anon_sym_BSLASH),
	2729: uint16(aux_sym_single_quoted_string_token1),
	2730: uint16(3),
	2731: uint16(5),
	2732: uint16(2),
	2733: uint16(sym_line_continuation),
	2734: uint16(sym_comment),
	2735: uint16(475),
	2736: uint16(2),
	2737: uint16(anon_sym_LF),
	2738: uint16(anon_sym_DOLLAR2),
	2739: uint16(477),
	2740: uint16(3),
	2741: uint16(aux_sym__env_key_token1),
	2742: uint16(aux_sym_unquoted_string_token1),
	2743: uint16(anon_sym_BSLASH2),
	2744: uint16(3),
	2745: uint16(5),
	2746: uint16(2),
	2747: uint16(sym_line_continuation),
	2748: uint16(sym_comment),
	2749: uint16(481),
	2750: uint16(2),
	2751: uint16(aux_sym_from_instruction_token2),
	2752: uint16(aux_sym_image_tag_token1),
	2753: uint16(479),
	2754: uint16(3),
	2755: uint16(anon_sym_LF),
	2756: uint16(anon_sym_DOLLAR2),
	2757: uint16(anon_sym_AT),
	2758: uint16(3),
	2759: uint16(5),
	2760: uint16(2),
	2761: uint16(sym_line_continuation),
	2762: uint16(sym_comment),
	2763: uint16(477),
	2764: uint16(2),
	2765: uint16(anon_sym_LF),
	2766: uint16(aux_sym_path_token3),
	2767: uint16(475),
	2768: uint16(3),
	2769: uint16(sym_heredoc_nl),
	2770: uint16(anon_sym_DOLLAR2),
	2771: uint16(sym__non_newline_whitespace),
	2772: uint16(3),
	2773: uint16(5),
	2774: uint16(2),
	2775: uint16(sym_line_continuation),
	2776: uint16(sym_comment),
	2777: uint16(479),
	2778: uint16(2),
	2779: uint16(anon_sym_LF),
	2780: uint16(anon_sym_DOLLAR2),
	2781: uint16(481),
	2782: uint16(3),
	2783: uint16(aux_sym__env_key_token1),
	2784: uint16(aux_sym_unquoted_string_token1),
	2785: uint16(anon_sym_BSLASH2),
	2786: uint16(5),
	2787: uint16(580),
	2788: uint16(1),
	2789: uint16(sym__non_newline_whitespace),
	2790: uint16(582),
	2791: uint16(1),
	2792: uint16(sym_heredoc_nl),
	2793: uint16(632),
	2794: uint16(1),
	2795: uint16(anon_sym_LF),
	2796: uint16(5),
	2797: uint16(2),
	2798: uint16(sym_line_continuation),
	2799: uint16(sym_comment),
	2800: uint16(142),
	2801: uint16(2),
	2802: uint16(sym_heredoc_block),
	2803: uint16(aux_sym_run_instruction_repeat2),
	2804: uint16(5),
	2805: uint16(636),
	2806: uint16(1),
	2807: uint16(anon_sym_SQUOTE),
	2808: uint16(638),
	2809: uint16(1),
	2810: uint16(sym_single_quoted_escape_sequence),
	2811: uint16(123),
	2812: uint16(1),
	2813: uint16(aux_sym_single_quoted_string_repeat1),
	2814: uint16(5),
	2815: uint16(2),
	2816: uint16(sym_line_continuation),
	2817: uint16(sym_comment),
	2818: uint16(634),
	2819: uint16(2),
	2820: uint16(anon_sym_BSLASH),
	2821: uint16(aux_sym_single_quoted_string_token1),
	2822: uint16(5),
	2823: uint16(606),
	2824: uint16(1),
	2825: uint16(sym_single_quoted_escape_sequence),
	2826: uint16(640),
	2827: uint16(1),
	2828: uint16(anon_sym_SQUOTE),
	2829: uint16(116),
	2830: uint16(1),
	2831: uint16(aux_sym_single_quoted_string_repeat1),
	2832: uint16(5),
	2833: uint16(2),
	2834: uint16(sym_line_continuation),
	2835: uint16(sym_comment),
	2836: uint16(602),
	2837: uint16(2),
	2838: uint16(anon_sym_BSLASH),
	2839: uint16(aux_sym_single_quoted_string_token1),
	2840: uint16(5),
	2841: uint16(580),
	2842: uint16(1),
	2843: uint16(sym__non_newline_whitespace),
	2844: uint16(582),
	2845: uint16(1),
	2846: uint16(sym_heredoc_nl),
	2847: uint16(642),
	2848: uint16(1),
	2849: uint16(anon_sym_LF),
	2850: uint16(5),
	2851: uint16(2),
	2852: uint16(sym_line_continuation),
	2853: uint16(sym_comment),
	2854: uint16(146),
	2855: uint16(2),
	2856: uint16(sym_heredoc_block),
	2857: uint16(aux_sym_run_instruction_repeat2),
	2858: uint16(5),
	2859: uint16(646),
	2860: uint16(1),
	2861: uint16(anon_sym_SQUOTE),
	2862: uint16(648),
	2863: uint16(1),
	2864: uint16(sym_single_quoted_escape_sequence),
	2865: uint16(126),
	2866: uint16(1),
	2867: uint16(aux_sym_single_quoted_string_repeat1),
	2868: uint16(5),
	2869: uint16(2),
	2870: uint16(sym_line_continuation),
	2871: uint16(sym_comment),
	2872: uint16(644),
	2873: uint16(2),
	2874: uint16(anon_sym_BSLASH),
	2875: uint16(aux_sym_single_quoted_string_token1),
	2876: uint16(5),
	2877: uint16(606),
	2878: uint16(1),
	2879: uint16(sym_single_quoted_escape_sequence),
	2880: uint16(650),
	2881: uint16(1),
	2882: uint16(anon_sym_SQUOTE),
	2883: uint16(116),
	2884: uint16(1),
	2885: uint16(aux_sym_single_quoted_string_repeat1),
	2886: uint16(5),
	2887: uint16(2),
	2888: uint16(sym_line_continuation),
	2889: uint16(sym_comment),
	2890: uint16(602),
	2891: uint16(2),
	2892: uint16(anon_sym_BSLASH),
	2893: uint16(aux_sym_single_quoted_string_token1),
	2894: uint16(3),
	2895: uint16(652),
	2896: uint16(1),
	2897: uint16(anon_sym_LF),
	2898: uint16(5),
	2899: uint16(2),
	2900: uint16(sym_line_continuation),
	2901: uint16(sym_comment),
	2902: uint16(654),
	2903: uint16(3),
	2904: uint16(aux_sym_label_pair_token1),
	2905: uint16(anon_sym_DQUOTE),
	2906: uint16(anon_sym_SQUOTE),
	2907: uint16(4),
	2908: uint16(656),
	2909: uint16(1),
	2910: uint16(anon_sym_LF),
	2911: uint16(658),
	2912: uint16(1),
	2913: uint16(sym_heredoc_nl),
	2914: uint16(5),
	2915: uint16(2),
	2916: uint16(sym_line_continuation),
	2917: uint16(sym_comment),
	2918: uint16(128),
	2919: uint16(2),
	2920: uint16(sym_heredoc_block),
	2921: uint16(aux_sym_run_instruction_repeat2),
	2922: uint16(4),
	2923: uint16(663),
	2924: uint16(1),
	2925: uint16(sym_required_line_continuation),
	2926: uint16(129),
	2927: uint16(1),
	2928: uint16(aux_sym_shell_command_repeat1),
	2929: uint16(5),
	2930: uint16(2),
	2931: uint16(sym_line_continuation),
	2932: uint16(sym_comment),
	2933: uint16(661),
	2934: uint16(2),
	2935: uint16(sym_heredoc_nl),
	2936: uint16(anon_sym_LF),
	2937: uint16(4),
	2938: uint16(582),
	2939: uint16(1),
	2940: uint16(sym_heredoc_nl),
	2941: uint16(666),
	2942: uint16(1),
	2943: uint16(anon_sym_LF),
	2944: uint16(5),
	2945: uint16(2),
	2946: uint16(sym_line_continuation),
	2947: uint16(sym_comment),
	2948: uint16(128),
	2949: uint16(2),
	2950: uint16(sym_heredoc_block),
	2951: uint16(aux_sym_run_instruction_repeat2),
	2952: uint16(4),
	2953: uint16(582),
	2954: uint16(1),
	2955: uint16(sym_heredoc_nl),
	2956: uint16(668),
	2957: uint16(1),
	2958: uint16(anon_sym_LF),
	2959: uint16(5),
	2960: uint16(2),
	2961: uint16(sym_line_continuation),
	2962: uint16(sym_comment),
	2963: uint16(130),
	2964: uint16(2),
	2965: uint16(sym_heredoc_block),
	2966: uint16(aux_sym_run_instruction_repeat2),
	2967: uint16(5),
	2968: uint16(670),
	2969: uint16(1),
	2970: uint16(aux_sym__stopsignal_value_token1),
	2971: uint16(672),
	2972: uint16(1),
	2973: uint16(anon_sym_DOLLAR),
	2974: uint16(91),
	2975: uint16(1),
	2976: uint16(sym_expansion),
	2977: uint16(256),
	2978: uint16(1),
	2979: uint16(sym__stopsignal_value),
	2980: uint16(3),
	2981: uint16(2),
	2982: uint16(sym_line_continuation),
	2983: uint16(sym_comment),
	2984: uint16(3),
	2985: uint16(5),
	2986: uint16(2),
	2987: uint16(sym_line_continuation),
	2988: uint16(sym_comment),
	2989: uint16(501),
	2990: uint16(2),
	2991: uint16(anon_sym_DOLLAR2),
	2992: uint16(sym__non_newline_whitespace),
	2993: uint16(503),
	2994: uint16(2),
	2995: uint16(anon_sym_LF),
	2996: uint16(aux_sym_path_token3),
	2997: uint16(3),
	2998: uint16(5),
	2999: uint16(2),
	3000: uint16(sym_line_continuation),
	3001: uint16(sym_comment),
	3002: uint16(475),
	3003: uint16(2),
	3004: uint16(anon_sym_DOLLAR2),
	3005: uint16(sym__non_newline_whitespace),
	3006: uint16(477),
	3007: uint16(2),
	3008: uint16(anon_sym_LF),
	3009: uint16(aux_sym_path_token3),
	3010: uint16(5),
	3011: uint16(674),
	3012: uint16(1),
	3013: uint16(anon_sym_DQUOTE),
	3014: uint16(676),
	3015: uint16(1),
	3016: uint16(aux_sym_json_string_token1),
	3017: uint16(679),
	3018: uint16(1),
	3019: uint16(sym_json_escape_sequence),
	3020: uint16(135),
	3021: uint16(1),
	3022: uint16(aux_sym_json_string_repeat1),
	3023: uint16(5),
	3024: uint16(2),
	3025: uint16(sym_line_continuation),
	3026: uint16(sym_comment),
	3027: uint16(3),
	3028: uint16(5),
	3029: uint16(2),
	3030: uint16(sym_line_continuation),
	3031: uint16(sym_comment),
	3032: uint16(479),
	3033: uint16(2),
	3034: uint16(anon_sym_DOLLAR2),
	3035: uint16(sym__non_newline_whitespace),
	3036: uint16(481),
	3037: uint16(2),
	3038: uint16(anon_sym_LF),
	3039: uint16(aux_sym_path_token3),
	3040: uint16(2),
	3041: uint16(5),
	3042: uint16(2),
	3043: uint16(sym_line_continuation),
	3044: uint16(sym_comment),
	3045: uint16(501),
	3046: uint16(4),
	3047: uint16(anon_sym_LF),
	3048: uint16(anon_sym_COLON),
	3049: uint16(aux_sym__immediate_user_name_or_group_fragment_token1),
	3050: uint16(anon_sym_DOLLAR2),
	3051: uint16(2),
	3052: uint16(5),
	3053: uint16(2),
	3054: uint16(sym_line_continuation),
	3055: uint16(sym_comment),
	3056: uint16(475),
	3057: uint16(4),
	3058: uint16(anon_sym_LF),
	3059: uint16(anon_sym_COLON),
	3060: uint16(aux_sym__immediate_user_name_or_group_fragment_token1),
	3061: uint16(anon_sym_DOLLAR2),
	3062: uint16(2),
	3063: uint16(5),
	3064: uint16(2),
	3065: uint16(sym_line_continuation),
	3066: uint16(sym_comment),
	3067: uint16(479),
	3068: uint16(4),
	3069: uint16(anon_sym_LF),
	3070: uint16(anon_sym_COLON),
	3071: uint16(aux_sym__immediate_user_name_or_group_fragment_token1),
	3072: uint16(anon_sym_DOLLAR2),
	3073: uint16(5),
	3074: uint16(682),
	3075: uint16(1),
	3076: uint16(anon_sym_DQUOTE),
	3077: uint16(684),
	3078: uint16(1),
	3079: uint16(aux_sym_json_string_token1),
	3080: uint16(686),
	3081: uint16(1),
	3082: uint16(sym_json_escape_sequence),
	3083: uint16(157),
	3084: uint16(1),
	3085: uint16(aux_sym_json_string_repeat1),
	3086: uint16(5),
	3087: uint16(2),
	3088: uint16(sym_line_continuation),
	3089: uint16(sym_comment),
	3090: uint16(4),
	3091: uint16(582),
	3092: uint16(1),
	3093: uint16(sym_heredoc_nl),
	3094: uint16(688),
	3095: uint16(1),
	3096: uint16(anon_sym_LF),
	3097: uint16(5),
	3098: uint16(2),
	3099: uint16(sym_line_continuation),
	3100: uint16(sym_comment),
	3101: uint16(156),
	3102: uint16(2),
	3103: uint16(sym_heredoc_block),
	3104: uint16(aux_sym_run_instruction_repeat2),
	3105: uint16(4),
	3106: uint16(582),
	3107: uint16(1),
	3108: uint16(sym_heredoc_nl),
	3109: uint16(690),
	3110: uint16(1),
	3111: uint16(anon_sym_LF),
	3112: uint16(5),
	3113: uint16(2),
	3114: uint16(sym_line_continuation),
	3115: uint16(sym_comment),
	3116: uint16(128),
	3117: uint16(2),
	3118: uint16(sym_heredoc_block),
	3119: uint16(aux_sym_run_instruction_repeat2),
	3120: uint16(5),
	3121: uint16(519),
	3122: uint16(1),
	3123: uint16(anon_sym_AT),
	3124: uint16(692),
	3125: uint16(1),
	3126: uint16(anon_sym_LF),
	3127: uint16(694),
	3128: uint16(1),
	3129: uint16(aux_sym_from_instruction_token2),
	3130: uint16(219),
	3131: uint16(1),
	3132: uint16(sym_image_digest),
	3133: uint16(5),
	3134: uint16(2),
	3135: uint16(sym_line_continuation),
	3136: uint16(sym_comment),
	3137: uint16(5),
	3138: uint16(696),
	3139: uint16(1),
	3140: uint16(anon_sym_DOLLAR),
	3141: uint16(698),
	3142: uint16(1),
	3143: uint16(aux_sym_image_alias_token1),
	3144: uint16(89),
	3145: uint16(1),
	3146: uint16(sym_expansion),
	3147: uint16(295),
	3148: uint16(1),
	3149: uint16(sym_image_alias),
	3150: uint16(3),
	3151: uint16(2),
	3152: uint16(sym_line_continuation),
	3153: uint16(sym_comment),
	3154: uint16(5),
	3155: uint16(700),
	3156: uint16(1),
	3157: uint16(aux_sym__user_name_or_group_token1),
	3158: uint16(702),
	3159: uint16(1),
	3160: uint16(anon_sym_DOLLAR),
	3161: uint16(31),
	3162: uint16(1),
	3163: uint16(sym_expansion),
	3164: uint16(247),
	3165: uint16(1),
	3166: uint16(sym__user_name_or_group),
	3167: uint16(3),
	3168: uint16(2),
	3169: uint16(sym_line_continuation),
	3170: uint16(sym_comment),
	3171: uint16(4),
	3172: uint16(582),
	3173: uint16(1),
	3174: uint16(sym_heredoc_nl),
	3175: uint16(704),
	3176: uint16(1),
	3177: uint16(anon_sym_LF),
	3178: uint16(5),
	3179: uint16(2),
	3180: uint16(sym_line_continuation),
	3181: uint16(sym_comment),
	3182: uint16(128),
	3183: uint16(2),
	3184: uint16(sym_heredoc_block),
	3185: uint16(aux_sym_run_instruction_repeat2),
	3186: uint16(4),
	3187: uint16(3),
	3188: uint16(1),
	3189: uint16(sym_line_continuation),
	3190: uint16(5),
	3191: uint16(1),
	3192: uint16(sym_comment),
	3193: uint16(706),
	3194: uint16(1),
	3195: uint16(aux_sym_path_token1),
	3196: uint16(708),
	3197: uint16(3),
	3198: uint16(sym_heredoc_marker),
	3199: uint16(aux_sym_path_with_heredoc_token1),
	3200: uint16(anon_sym_DOLLAR),
	3201: uint16(3),
	3202: uint16(710),
	3203: uint16(1),
	3204: uint16(anon_sym_LF),
	3205: uint16(5),
	3206: uint16(2),
	3207: uint16(sym_line_continuation),
	3208: uint16(sym_comment),
	3209: uint16(712),
	3210: uint16(3),
	3211: uint16(aux_sym_label_pair_token1),
	3212: uint16(anon_sym_DQUOTE),
	3213: uint16(anon_sym_SQUOTE),
	3214: uint16(3),
	3215: uint16(481),
	3216: uint16(1),
	3217: uint16(aux_sym_from_instruction_token2),
	3218: uint16(5),
	3219: uint16(2),
	3220: uint16(sym_line_continuation),
	3221: uint16(sym_comment),
	3222: uint16(479),
	3223: uint16(3),
	3224: uint16(anon_sym_LF),
	3225: uint16(anon_sym_DOLLAR2),
	3226: uint16(aux_sym_image_digest_token1),
	3227: uint16(4),
	3228: uint16(716),
	3229: uint16(1),
	3230: uint16(sym_required_line_continuation),
	3231: uint16(152),
	3232: uint16(1),
	3233: uint16(aux_sym_shell_command_repeat1),
	3234: uint16(5),
	3235: uint16(2),
	3236: uint16(sym_line_continuation),
	3237: uint16(sym_comment),
	3238: uint16(714),
	3239: uint16(2),
	3240: uint16(sym_heredoc_nl),
	3241: uint16(anon_sym_LF),
	3242: uint16(3),
	3243: uint16(718),
	3244: uint16(1),
	3245: uint16(anon_sym_LF),
	3246: uint16(5),
	3247: uint16(2),
	3248: uint16(sym_line_continuation),
	3249: uint16(sym_comment),
	3250: uint16(720),
	3251: uint16(3),
	3252: uint16(aux_sym_label_pair_token1),
	3253: uint16(anon_sym_DQUOTE),
	3254: uint16(anon_sym_SQUOTE),
	3255: uint16(4),
	3256: uint16(716),
	3257: uint16(1),
	3258: uint16(sym_required_line_continuation),
	3259: uint16(129),
	3260: uint16(1),
	3261: uint16(aux_sym_shell_command_repeat1),
	3262: uint16(5),
	3263: uint16(2),
	3264: uint16(sym_line_continuation),
	3265: uint16(sym_comment),
	3266: uint16(722),
	3267: uint16(2),
	3268: uint16(sym_heredoc_nl),
	3269: uint16(anon_sym_LF),
	3270: uint16(4),
	3271: uint16(436),
	3272: uint16(1),
	3273: uint16(aux_sym_cmd_instruction_token1),
	3274: uint16(724),
	3275: uint16(1),
	3276: uint16(anon_sym_DASH_DASH),
	3277: uint16(3),
	3278: uint16(2),
	3279: uint16(sym_line_continuation),
	3280: uint16(sym_comment),
	3281: uint16(153),
	3282: uint16(2),
	3283: uint16(sym_param),
	3284: uint16(aux_sym_add_instruction_repeat1),
	3285: uint16(4),
	3286: uint16(582),
	3287: uint16(1),
	3288: uint16(sym_heredoc_nl),
	3289: uint16(727),
	3290: uint16(1),
	3291: uint16(anon_sym_LF),
	3292: uint16(5),
	3293: uint16(2),
	3294: uint16(sym_line_continuation),
	3295: uint16(sym_comment),
	3296: uint16(128),
	3297: uint16(2),
	3298: uint16(sym_heredoc_block),
	3299: uint16(aux_sym_run_instruction_repeat2),
	3300: uint16(5),
	3301: uint16(696),
	3302: uint16(1),
	3303: uint16(anon_sym_DOLLAR),
	3304: uint16(698),
	3305: uint16(1),
	3306: uint16(aux_sym_image_alias_token1),
	3307: uint16(89),
	3308: uint16(1),
	3309: uint16(sym_expansion),
	3310: uint16(261),
	3311: uint16(1),
	3312: uint16(sym_image_alias),
	3313: uint16(3),
	3314: uint16(2),
	3315: uint16(sym_line_continuation),
	3316: uint16(sym_comment),
	3317: uint16(4),
	3318: uint16(582),
	3319: uint16(1),
	3320: uint16(sym_heredoc_nl),
	3321: uint16(668),
	3322: uint16(1),
	3323: uint16(anon_sym_LF),
	3324: uint16(5),
	3325: uint16(2),
	3326: uint16(sym_line_continuation),
	3327: uint16(sym_comment),
	3328: uint16(128),
	3329: uint16(2),
	3330: uint16(sym_heredoc_block),
	3331: uint16(aux_sym_run_instruction_repeat2),
	3332: uint16(5),
	3333: uint16(729),
	3334: uint16(1),
	3335: uint16(anon_sym_DQUOTE),
	3336: uint16(731),
	3337: uint16(1),
	3338: uint16(aux_sym_json_string_token1),
	3339: uint16(733),
	3340: uint16(1),
	3341: uint16(sym_json_escape_sequence),
	3342: uint16(135),
	3343: uint16(1),
	3344: uint16(aux_sym_json_string_repeat1),
	3345: uint16(5),
	3346: uint16(2),
	3347: uint16(sym_line_continuation),
	3348: uint16(sym_comment),
	3349: uint16(4),
	3350: uint16(582),
	3351: uint16(1),
	3352: uint16(sym_heredoc_nl),
	3353: uint16(735),
	3354: uint16(1),
	3355: uint16(anon_sym_LF),
	3356: uint16(5),
	3357: uint16(2),
	3358: uint16(sym_line_continuation),
	3359: uint16(sym_comment),
	3360: uint16(128),
	3361: uint16(2),
	3362: uint16(sym_heredoc_block),
	3363: uint16(aux_sym_run_instruction_repeat2),
	3364: uint16(3),
	3365: uint16(737),
	3366: uint16(1),
	3367: uint16(anon_sym_LF),
	3368: uint16(5),
	3369: uint16(2),
	3370: uint16(sym_line_continuation),
	3371: uint16(sym_comment),
	3372: uint16(739),
	3373: uint16(3),
	3374: uint16(aux_sym_label_pair_token1),
	3375: uint16(anon_sym_DQUOTE),
	3376: uint16(anon_sym_SQUOTE),
	3377: uint16(3),
	3378: uint16(741),
	3379: uint16(1),
	3380: uint16(anon_sym_LF),
	3381: uint16(5),
	3382: uint16(2),
	3383: uint16(sym_line_continuation),
	3384: uint16(sym_comment),
	3385: uint16(743),
	3386: uint16(3),
	3387: uint16(aux_sym_label_pair_token1),
	3388: uint16(anon_sym_DQUOTE),
	3389: uint16(anon_sym_SQUOTE),
	3390: uint16(3),
	3391: uint16(477),
	3392: uint16(1),
	3393: uint16(aux_sym_from_instruction_token2),
	3394: uint16(5),
	3395: uint16(2),
	3396: uint16(sym_line_continuation),
	3397: uint16(sym_comment),
	3398: uint16(475),
	3399: uint16(3),
	3400: uint16(anon_sym_LF),
	3401: uint16(anon_sym_DOLLAR2),
	3402: uint16(aux_sym_image_digest_token1),
	3403: uint16(4),
	3404: uint16(745),
	3405: uint16(1),
	3406: uint16(anon_sym_LBRACE),
	3407: uint16(747),
	3408: uint16(1),
	3409: uint16(sym_variable),
	3410: uint16(133),
	3411: uint16(1),
	3412: uint16(sym__expansion_body),
	3413: uint16(3),
	3414: uint16(2),
	3415: uint16(sym_line_continuation),
	3416: uint16(sym_comment),
	3417: uint16(3),
	3418: uint16(749),
	3419: uint16(1),
	3420: uint16(anon_sym_LF),
	3421: uint16(5),
	3422: uint16(2),
	3423: uint16(sym_line_continuation),
	3424: uint16(sym_comment),
	3425: uint16(751),
	3426: uint16(2),
	3427: uint16(anon_sym_DOLLAR),
	3428: uint16(aux_sym_expose_port_token1),
	3429: uint16(2),
	3430: uint16(5),
	3431: uint16(2),
	3432: uint16(sym_line_continuation),
	3433: uint16(sym_comment),
	3434: uint16(753),
	3435: uint16(3),
	3436: uint16(anon_sym_LF),
	3437: uint16(aux_sym__immediate_user_name_or_group_fragment_token1),
	3438: uint16(anon_sym_DOLLAR2),
	3439: uint16(4),
	3440: uint16(755),
	3441: uint16(1),
	3442: uint16(anon_sym_COMMA2),
	3443: uint16(758),
	3444: uint16(1),
	3445: uint16(anon_sym_RBRACK),
	3446: uint16(165),
	3447: uint16(1),
	3448: uint16(aux_sym_json_string_array_repeat1),
	3449: uint16(3),
	3450: uint16(2),
	3451: uint16(sym_line_continuation),
	3452: uint16(sym_comment),
	3453: uint16(4),
	3454: uint16(760),
	3455: uint16(1),
	3456: uint16(sym_heredoc_line),
	3457: uint16(763),
	3458: uint16(1),
	3459: uint16(sym_heredoc_end),
	3460: uint16(166),
	3461: uint16(1),
	3462: uint16(aux_sym_heredoc_block_repeat1),
	3463: uint16(3),
	3464: uint16(2),
	3465: uint16(sym_line_continuation),
	3466: uint16(sym_comment),
	3467: uint16(4),
	3468: uint16(714),
	3469: uint16(1),
	3470: uint16(anon_sym_LF),
	3471: uint16(765),
	3472: uint16(1),
	3473: uint16(sym_required_line_continuation),
	3474: uint16(173),
	3475: uint16(1),
	3476: uint16(aux_sym_shell_command_repeat1),
	3477: uint16(5),
	3478: uint16(2),
	3479: uint16(sym_line_continuation),
	3480: uint16(sym_comment),
	3481: uint16(3),
	3482: uint16(322),
	3483: uint16(1),
	3484: uint16(anon_sym_LF),
	3485: uint16(5),
	3486: uint16(2),
	3487: uint16(sym_line_continuation),
	3488: uint16(sym_comment),
	3489: uint16(328),
	3490: uint16(2),
	3491: uint16(sym_heredoc_nl),
	3492: uint16(sym__non_newline_whitespace),
	3493: uint16(4),
	3494: uint16(767),
	3495: uint16(1),
	3496: uint16(anon_sym_LF),
	3497: uint16(769),
	3498: uint16(1),
	3499: uint16(sym__non_newline_whitespace),
	3500: uint16(202),
	3501: uint16(1),
	3502: uint16(aux_sym_volume_instruction_repeat1),
	3503: uint16(5),
	3504: uint16(2),
	3505: uint16(sym_line_continuation),
	3506: uint16(sym_comment),
	3507: uint16(3),
	3508: uint16(501),
	3509: uint16(1),
	3510: uint16(anon_sym_LF),
	3511: uint16(5),
	3512: uint16(2),
	3513: uint16(sym_line_continuation),
	3514: uint16(sym_comment),
	3515: uint16(503),
	3516: uint16(2),
	3517: uint16(anon_sym_DOLLAR),
	3518: uint16(aux_sym_expose_port_token1),
	3519: uint16(4),
	3520: uint16(771),
	3521: uint16(1),
	3522: uint16(anon_sym_LBRACE),
	3523: uint16(773),
	3524: uint16(1),
	3525: uint16(sym_variable),
	3526: uint16(109),
	3527: uint16(1),
	3528: uint16(sym__expansion_body),
	3529: uint16(3),
	3530: uint16(2),
	3531: uint16(sym_line_continuation),
	3532: uint16(sym_comment),
	3533: uint16(4),
	3534: uint16(775),
	3535: uint16(1),
	3536: uint16(anon_sym_LBRACE),
	3537: uint16(777),
	3538: uint16(1),
	3539: uint16(sym_variable),
	3540: uint16(75),
	3541: uint16(1),
	3542: uint16(sym__expansion_body),
	3543: uint16(3),
	3544: uint16(2),
	3545: uint16(sym_line_continuation),
	3546: uint16(sym_comment),
	3547: uint16(4),
	3548: uint16(722),
	3549: uint16(1),
	3550: uint16(anon_sym_LF),
	3551: uint16(765),
	3552: uint16(1),
	3553: uint16(sym_required_line_continuation),
	3554: uint16(174),
	3555: uint16(1),
	3556: uint16(aux_sym_shell_command_repeat1),
	3557: uint16(5),
	3558: uint16(2),
	3559: uint16(sym_line_continuation),
	3560: uint16(sym_comment),
	3561: uint16(4),
	3562: uint16(661),
	3563: uint16(1),
	3564: uint16(anon_sym_LF),
	3565: uint16(779),
	3566: uint16(1),
	3567: uint16(sym_required_line_continuation),
	3568: uint16(174),
	3569: uint16(1),
	3570: uint16(aux_sym_shell_command_repeat1),
	3571: uint16(5),
	3572: uint16(2),
	3573: uint16(sym_line_continuation),
	3574: uint16(sym_comment),
	3575: uint16(3),
	3576: uint16(479),
	3577: uint16(1),
	3578: uint16(anon_sym_LF),
	3579: uint16(5),
	3580: uint16(2),
	3581: uint16(sym_line_continuation),
	3582: uint16(sym_comment),
	3583: uint16(481),
	3584: uint16(2),
	3585: uint16(anon_sym_DOLLAR),
	3586: uint16(aux_sym_expose_port_token1),
	3587: uint16(3),
	3588: uint16(503),
	3589: uint16(1),
	3590: uint16(aux_sym_path_token3),
	3591: uint16(5),
	3592: uint16(2),
	3593: uint16(sym_line_continuation),
	3594: uint16(sym_comment),
	3595: uint16(501),
	3596: uint16(2),
	3597: uint16(anon_sym_DOLLAR2),
	3598: uint16(sym__non_newline_whitespace),
	3599: uint16(3),
	3600: uint16(477),
	3601: uint16(1),
	3602: uint16(aux_sym_path_token3),
	3603: uint16(5),
	3604: uint16(2),
	3605: uint16(sym_line_continuation),
	3606: uint16(sym_comment),
	3607: uint16(475),
	3608: uint16(2),
	3609: uint16(anon_sym_DOLLAR2),
	3610: uint16(sym__non_newline_whitespace),
	3611: uint16(3),
	3612: uint16(481),
	3613: uint16(1),
	3614: uint16(aux_sym_path_token3),
	3615: uint16(5),
	3616: uint16(2),
	3617: uint16(sym_line_continuation),
	3618: uint16(sym_comment),
	3619: uint16(479),
	3620: uint16(2),
	3621: uint16(anon_sym_DOLLAR2),
	3622: uint16(sym__non_newline_whitespace),
	3623: uint16(4),
	3624: uint16(782),
	3625: uint16(1),
	3626: uint16(anon_sym_COMMA2),
	3627: uint16(784),
	3628: uint16(1),
	3629: uint16(anon_sym_RBRACK),
	3630: uint16(165),
	3631: uint16(1),
	3632: uint16(aux_sym_json_string_array_repeat1),
	3633: uint16(3),
	3634: uint16(2),
	3635: uint16(sym_line_continuation),
	3636: uint16(sym_comment),
	3637: uint16(4),
	3638: uint16(786),
	3639: uint16(1),
	3640: uint16(sym_heredoc_line),
	3641: uint16(788),
	3642: uint16(1),
	3643: uint16(sym_heredoc_end),
	3644: uint16(166),
	3645: uint16(1),
	3646: uint16(aux_sym_heredoc_block_repeat1),
	3647: uint16(3),
	3648: uint16(2),
	3649: uint16(sym_line_continuation),
	3650: uint16(sym_comment),
	3651: uint16(3),
	3652: uint16(503),
	3653: uint16(1),
	3654: uint16(aux_sym_path_token3),
	3655: uint16(5),
	3656: uint16(2),
	3657: uint16(sym_line_continuation),
	3658: uint16(sym_comment),
	3659: uint16(501),
	3660: uint16(2),
	3661: uint16(anon_sym_LF),
	3662: uint16(anon_sym_DOLLAR2),
	3663: uint16(3),
	3664: uint16(477),
	3665: uint16(1),
	3666: uint16(aux_sym_path_token3),
	3667: uint16(5),
	3668: uint16(2),
	3669: uint16(sym_line_continuation),
	3670: uint16(sym_comment),
	3671: uint16(475),
	3672: uint16(2),
	3673: uint16(anon_sym_LF),
	3674: uint16(anon_sym_DOLLAR2),
	3675: uint16(3),
	3676: uint16(481),
	3677: uint16(1),
	3678: uint16(aux_sym_path_token3),
	3679: uint16(5),
	3680: uint16(2),
	3681: uint16(sym_line_continuation),
	3682: uint16(sym_comment),
	3683: uint16(479),
	3684: uint16(2),
	3685: uint16(anon_sym_LF),
	3686: uint16(anon_sym_DOLLAR2),
	3687: uint16(2),
	3688: uint16(5),
	3689: uint16(2),
	3690: uint16(sym_line_continuation),
	3691: uint16(sym_comment),
	3692: uint16(501),
	3693: uint16(3),
	3694: uint16(anon_sym_LF),
	3695: uint16(aux_sym__stopsignal_value_token2),
	3696: uint16(anon_sym_DOLLAR2),
	3697: uint16(2),
	3698: uint16(5),
	3699: uint16(2),
	3700: uint16(sym_line_continuation),
	3701: uint16(sym_comment),
	3702: uint16(475),
	3703: uint16(3),
	3704: uint16(anon_sym_LF),
	3705: uint16(aux_sym__stopsignal_value_token2),
	3706: uint16(anon_sym_DOLLAR2),
	3707: uint16(4),
	3708: uint16(790),
	3709: uint16(1),
	3710: uint16(anon_sym_LBRACE),
	3711: uint16(792),
	3712: uint16(1),
	3713: uint16(sym_variable),
	3714: uint16(193),
	3715: uint16(1),
	3716: uint16(sym__expansion_body),
	3717: uint16(3),
	3718: uint16(2),
	3719: uint16(sym_line_continuation),
	3720: uint16(sym_comment),
	3721: uint16(3),
	3722: uint16(794),
	3723: uint16(1),
	3724: uint16(sym_required_line_continuation),
	3725: uint16(5),
	3726: uint16(2),
	3727: uint16(sym_line_continuation),
	3728: uint16(sym_comment),
	3729: uint16(661),
	3730: uint16(2),
	3731: uint16(sym_heredoc_nl),
	3732: uint16(anon_sym_LF),
	3733: uint16(2),
	3734: uint16(5),
	3735: uint16(2),
	3736: uint16(sym_line_continuation),
	3737: uint16(sym_comment),
	3738: uint16(501),
	3739: uint16(3),
	3740: uint16(anon_sym_LF),
	3741: uint16(anon_sym_DOLLAR2),
	3742: uint16(aux_sym_image_alias_token2),
	3743: uint16(4),
	3744: uint16(782),
	3745: uint16(1),
	3746: uint16(anon_sym_COMMA2),
	3747: uint16(796),
	3748: uint16(1),
	3749: uint16(anon_sym_RBRACK),
	3750: uint16(179),
	3751: uint16(1),
	3752: uint16(aux_sym_json_string_array_repeat1),
	3753: uint16(3),
	3754: uint16(2),
	3755: uint16(sym_line_continuation),
	3756: uint16(sym_comment),
	3757: uint16(4),
	3758: uint16(786),
	3759: uint16(1),
	3760: uint16(sym_heredoc_line),
	3761: uint16(798),
	3762: uint16(1),
	3763: uint16(sym_heredoc_end),
	3764: uint16(180),
	3765: uint16(1),
	3766: uint16(aux_sym_heredoc_block_repeat1),
	3767: uint16(3),
	3768: uint16(2),
	3769: uint16(sym_line_continuation),
	3770: uint16(sym_comment),
	3771: uint16(4),
	3772: uint16(800),
	3773: uint16(1),
	3774: uint16(anon_sym_RBRACK),
	3775: uint16(802),
	3776: uint16(1),
	3777: uint16(anon_sym_DQUOTE),
	3778: uint16(189),
	3779: uint16(1),
	3780: uint16(sym_json_string),
	3781: uint16(3),
	3782: uint16(2),
	3783: uint16(sym_line_continuation),
	3784: uint16(sym_comment),
	3785: uint16(2),
	3786: uint16(5),
	3787: uint16(2),
	3788: uint16(sym_line_continuation),
	3789: uint16(sym_comment),
	3790: uint16(479),
	3791: uint16(3),
	3792: uint16(anon_sym_LF),
	3793: uint16(anon_sym_DOLLAR2),
	3794: uint16(aux_sym_image_alias_token2),
	3795: uint16(2),
	3796: uint16(5),
	3797: uint16(2),
	3798: uint16(sym_line_continuation),
	3799: uint16(sym_comment),
	3800: uint16(475),
	3801: uint16(3),
	3802: uint16(anon_sym_LF),
	3803: uint16(anon_sym_DOLLAR2),
	3804: uint16(aux_sym_image_alias_token2),
	3805: uint16(4),
	3806: uint16(804),
	3807: uint16(1),
	3808: uint16(anon_sym_LBRACE),
	3809: uint16(806),
	3810: uint16(1),
	3811: uint16(sym_variable),
	3812: uint16(170),
	3813: uint16(1),
	3814: uint16(sym__expansion_body),
	3815: uint16(3),
	3816: uint16(2),
	3817: uint16(sym_line_continuation),
	3818: uint16(sym_comment),
	3819: uint16(4),
	3820: uint16(802),
	3821: uint16(1),
	3822: uint16(anon_sym_DQUOTE),
	3823: uint16(808),
	3824: uint16(1),
	3825: uint16(anon_sym_RBRACK),
	3826: uint16(196),
	3827: uint16(1),
	3828: uint16(sym_json_string),
	3829: uint16(3),
	3830: uint16(2),
	3831: uint16(sym_line_continuation),
	3832: uint16(sym_comment),
	3833: uint16(4),
	3834: uint16(782),
	3835: uint16(1),
	3836: uint16(anon_sym_COMMA2),
	3837: uint16(810),
	3838: uint16(1),
	3839: uint16(anon_sym_RBRACK),
	3840: uint16(199),
	3841: uint16(1),
	3842: uint16(aux_sym_json_string_array_repeat1),
	3843: uint16(3),
	3844: uint16(2),
	3845: uint16(sym_line_continuation),
	3846: uint16(sym_comment),
	3847: uint16(4),
	3848: uint16(775),
	3849: uint16(1),
	3850: uint16(anon_sym_LBRACE),
	3851: uint16(812),
	3852: uint16(1),
	3853: uint16(sym_variable),
	3854: uint16(80),
	3855: uint16(1),
	3856: uint16(sym__expansion_body),
	3857: uint16(3),
	3858: uint16(2),
	3859: uint16(sym_line_continuation),
	3860: uint16(sym_comment),
	3861: uint16(4),
	3862: uint16(769),
	3863: uint16(1),
	3864: uint16(sym__non_newline_whitespace),
	3865: uint16(814),
	3866: uint16(1),
	3867: uint16(anon_sym_LF),
	3868: uint16(169),
	3869: uint16(1),
	3870: uint16(aux_sym_volume_instruction_repeat1),
	3871: uint16(5),
	3872: uint16(2),
	3873: uint16(sym_line_continuation),
	3874: uint16(sym_comment),
	3875: uint16(4),
	3876: uint16(782),
	3877: uint16(1),
	3878: uint16(anon_sym_COMMA2),
	3879: uint16(816),
	3880: uint16(1),
	3881: uint16(anon_sym_RBRACK),
	3882: uint16(165),
	3883: uint16(1),
	3884: uint16(aux_sym_json_string_array_repeat1),
	3885: uint16(3),
	3886: uint16(2),
	3887: uint16(sym_line_continuation),
	3888: uint16(sym_comment),
	3889: uint16(4),
	3890: uint16(818),
	3891: uint16(1),
	3892: uint16(anon_sym_LBRACE),
	3893: uint16(820),
	3894: uint16(1),
	3895: uint16(sym_variable),
	3896: uint16(177),
	3897: uint16(1),
	3898: uint16(sym__expansion_body),
	3899: uint16(3),
	3900: uint16(2),
	3901: uint16(sym_line_continuation),
	3902: uint16(sym_comment),
	3903: uint16(4),
	3904: uint16(818),
	3905: uint16(1),
	3906: uint16(anon_sym_LBRACE),
	3907: uint16(822),
	3908: uint16(1),
	3909: uint16(sym_variable),
	3910: uint16(176),
	3911: uint16(1),
	3912: uint16(sym__expansion_body),
	3913: uint16(3),
	3914: uint16(2),
	3915: uint16(sym_line_continuation),
	3916: uint16(sym_comment),
	3917: uint16(4),
	3918: uint16(824),
	3919: uint16(1),
	3920: uint16(anon_sym_LF),
	3921: uint16(826),
	3922: uint16(1),
	3923: uint16(sym__non_newline_whitespace),
	3924: uint16(202),
	3925: uint16(1),
	3926: uint16(aux_sym_volume_instruction_repeat1),
	3927: uint16(5),
	3928: uint16(2),
	3929: uint16(sym_line_continuation),
	3930: uint16(sym_comment),
	3931: uint16(4),
	3932: uint16(745),
	3933: uint16(1),
	3934: uint16(anon_sym_LBRACE),
	3935: uint16(829),
	3936: uint16(1),
	3937: uint16(sym_variable),
	3938: uint16(134),
	3939: uint16(1),
	3940: uint16(sym__expansion_body),
	3941: uint16(3),
	3942: uint16(2),
	3943: uint16(sym_line_continuation),
	3944: uint16(sym_comment),
	3945: uint16(4),
	3946: uint16(831),
	3947: uint16(1),
	3948: uint16(anon_sym_LBRACE),
	3949: uint16(833),
	3950: uint16(1),
	3951: uint16(sym_variable),
	3952: uint16(138),
	3953: uint16(1),
	3954: uint16(sym__expansion_body),
	3955: uint16(3),
	3956: uint16(2),
	3957: uint16(sym_line_continuation),
	3958: uint16(sym_comment),
	3959: uint16(4),
	3960: uint16(831),
	3961: uint16(1),
	3962: uint16(anon_sym_LBRACE),
	3963: uint16(835),
	3964: uint16(1),
	3965: uint16(sym_variable),
	3966: uint16(137),
	3967: uint16(1),
	3968: uint16(sym__expansion_body),
	3969: uint16(3),
	3970: uint16(2),
	3971: uint16(sym_line_continuation),
	3972: uint16(sym_comment),
	3973: uint16(4),
	3974: uint16(837),
	3975: uint16(1),
	3976: uint16(anon_sym_LBRACE),
	3977: uint16(839),
	3978: uint16(1),
	3979: uint16(sym_variable),
	3980: uint16(182),
	3981: uint16(1),
	3982: uint16(sym__expansion_body),
	3983: uint16(3),
	3984: uint16(2),
	3985: uint16(sym_line_continuation),
	3986: uint16(sym_comment),
	3987: uint16(4),
	3988: uint16(837),
	3989: uint16(1),
	3990: uint16(anon_sym_LBRACE),
	3991: uint16(841),
	3992: uint16(1),
	3993: uint16(sym_variable),
	3994: uint16(181),
	3995: uint16(1),
	3996: uint16(sym__expansion_body),
	3997: uint16(3),
	3998: uint16(2),
	3999: uint16(sym_line_continuation),
	4000: uint16(sym_comment),
	4001: uint16(4),
	4002: uint16(843),
	4003: uint16(1),
	4004: uint16(anon_sym_LBRACE),
	4005: uint16(845),
	4006: uint16(1),
	4007: uint16(sym_variable),
	4008: uint16(185),
	4009: uint16(1),
	4010: uint16(sym__expansion_body),
	4011: uint16(3),
	4012: uint16(2),
	4013: uint16(sym_line_continuation),
	4014: uint16(sym_comment),
	4015: uint16(4),
	4016: uint16(843),
	4017: uint16(1),
	4018: uint16(anon_sym_LBRACE),
	4019: uint16(847),
	4020: uint16(1),
	4021: uint16(sym_variable),
	4022: uint16(184),
	4023: uint16(1),
	4024: uint16(sym__expansion_body),
	4025: uint16(3),
	4026: uint16(2),
	4027: uint16(sym_line_continuation),
	4028: uint16(sym_comment),
	4029: uint16(4),
	4030: uint16(849),
	4031: uint16(1),
	4032: uint16(anon_sym_LBRACE),
	4033: uint16(851),
	4034: uint16(1),
	4035: uint16(sym_variable),
	4036: uint16(95),
	4037: uint16(1),
	4038: uint16(sym__expansion_body),
	4039: uint16(3),
	4040: uint16(2),
	4041: uint16(sym_line_continuation),
	4042: uint16(sym_comment),
	4043: uint16(4),
	4044: uint16(853),
	4045: uint16(1),
	4046: uint16(anon_sym_LBRACE),
	4047: uint16(855),
	4048: uint16(1),
	4049: uint16(sym_variable),
	4050: uint16(113),
	4051: uint16(1),
	4052: uint16(sym__expansion_body),
	4053: uint16(3),
	4054: uint16(2),
	4055: uint16(sym_line_continuation),
	4056: uint16(sym_comment),
	4057: uint16(4),
	4058: uint16(857),
	4059: uint16(1),
	4060: uint16(anon_sym_LBRACE),
	4061: uint16(859),
	4062: uint16(1),
	4063: uint16(sym_variable),
	4064: uint16(161),
	4065: uint16(1),
	4066: uint16(sym__expansion_body),
	4067: uint16(3),
	4068: uint16(2),
	4069: uint16(sym_line_continuation),
	4070: uint16(sym_comment),
	4071: uint16(4),
	4072: uint16(790),
	4073: uint16(1),
	4074: uint16(anon_sym_LBRACE),
	4075: uint16(861),
	4076: uint16(1),
	4077: uint16(sym_variable),
	4078: uint16(188),
	4079: uint16(1),
	4080: uint16(sym__expansion_body),
	4081: uint16(3),
	4082: uint16(2),
	4083: uint16(sym_line_continuation),
	4084: uint16(sym_comment),
	4085: uint16(4),
	4086: uint16(863),
	4087: uint16(1),
	4088: uint16(anon_sym_LBRACE),
	4089: uint16(865),
	4090: uint16(1),
	4091: uint16(sym_variable),
	4092: uint16(72),
	4093: uint16(1),
	4094: uint16(sym__expansion_body),
	4095: uint16(3),
	4096: uint16(2),
	4097: uint16(sym_line_continuation),
	4098: uint16(sym_comment),
	4099: uint16(4),
	4100: uint16(867),
	4101: uint16(1),
	4102: uint16(anon_sym_LBRACE),
	4103: uint16(869),
	4104: uint16(1),
	4105: uint16(sym_variable),
	4106: uint16(117),
	4107: uint16(1),
	4108: uint16(sym__expansion_body),
	4109: uint16(3),
	4110: uint16(2),
	4111: uint16(sym_line_continuation),
	4112: uint16(sym_comment),
	4113: uint16(4),
	4114: uint16(853),
	4115: uint16(1),
	4116: uint16(anon_sym_LBRACE),
	4117: uint16(871),
	4118: uint16(1),
	4119: uint16(sym_variable),
	4120: uint16(119),
	4121: uint16(1),
	4122: uint16(sym__expansion_body),
	4123: uint16(3),
	4124: uint16(2),
	4125: uint16(sym_line_continuation),
	4126: uint16(sym_comment),
	4127: uint16(2),
	4128: uint16(5),
	4129: uint16(2),
	4130: uint16(sym_line_continuation),
	4131: uint16(sym_comment),
	4132: uint16(479),
	4133: uint16(3),
	4134: uint16(anon_sym_LF),
	4135: uint16(aux_sym__stopsignal_value_token2),
	4136: uint16(anon_sym_DOLLAR2),
	4137: uint16(2),
	4138: uint16(5),
	4139: uint16(2),
	4140: uint16(sym_line_continuation),
	4141: uint16(sym_comment),
	4142: uint16(873),
	4143: uint16(2),
	4144: uint16(sym_heredoc_nl),
	4145: uint16(anon_sym_LF),
	4146: uint16(3),
	4147: uint16(875),
	4148: uint16(1),
	4149: uint16(anon_sym_LF),
	4150: uint16(877),
	4151: uint16(1),
	4152: uint16(aux_sym_from_instruction_token2),
	4153: uint16(5),
	4154: uint16(2),
	4155: uint16(sym_line_continuation),
	4156: uint16(sym_comment),
	4157: uint16(3),
	4158: uint16(879),
	4159: uint16(1),
	4160: uint16(aux_sym_param_token1),
	4161: uint16(881),
	4162: uint16(1),
	4163: uint16(anon_sym_mount),
	4164: uint16(3),
	4165: uint16(2),
	4166: uint16(sym_line_continuation),
	4167: uint16(sym_comment),
	4168: uint16(4),
	4169: uint16(3),
	4170: uint16(1),
	4171: uint16(sym_line_continuation),
	4172: uint16(5),
	4173: uint16(1),
	4174: uint16(sym_comment),
	4175: uint16(459),
	4176: uint16(1),
	4177: uint16(aux_sym_image_name_token1),
	4178: uint16(461),
	4179: uint16(1),
	4180: uint16(anon_sym_DOLLAR),
	4181: uint16(3),
	4182: uint16(710),
	4183: uint16(1),
	4184: uint16(anon_sym_LF),
	4185: uint16(712),
	4186: uint16(1),
	4187: uint16(aux_sym__env_key_token1),
	4188: uint16(5),
	4189: uint16(2),
	4190: uint16(sym_line_continuation),
	4191: uint16(sym_comment),
	4192: uint16(3),
	4193: uint16(718),
	4194: uint16(1),
	4195: uint16(anon_sym_LF),
	4196: uint16(720),
	4197: uint16(1),
	4198: uint16(aux_sym__env_key_token1),
	4199: uint16(5),
	4200: uint16(2),
	4201: uint16(sym_line_continuation),
	4202: uint16(sym_comment),
	4203: uint16(2),
	4204: uint16(3),
	4205: uint16(2),
	4206: uint16(sym_line_continuation),
	4207: uint16(sym_comment),
	4208: uint16(763),
	4209: uint16(2),
	4210: uint16(sym_heredoc_line),
	4211: uint16(sym_heredoc_end),
	4212: uint16(3),
	4213: uint16(737),
	4214: uint16(1),
	4215: uint16(anon_sym_LF),
	4216: uint16(739),
	4217: uint16(1),
	4218: uint16(aux_sym__env_key_token1),
	4219: uint16(5),
	4220: uint16(2),
	4221: uint16(sym_line_continuation),
	4222: uint16(sym_comment),
	4223: uint16(3),
	4224: uint16(741),
	4225: uint16(1),
	4226: uint16(anon_sym_LF),
	4227: uint16(743),
	4228: uint16(1),
	4229: uint16(aux_sym__env_key_token1),
	4230: uint16(5),
	4231: uint16(2),
	4232: uint16(sym_line_continuation),
	4233: uint16(sym_comment),
	4234: uint16(4),
	4235: uint16(3),
	4236: uint16(1),
	4237: uint16(sym_line_continuation),
	4238: uint16(5),
	4239: uint16(1),
	4240: uint16(sym_comment),
	4241: uint16(883),
	4242: uint16(1),
	4243: uint16(aux_sym_mount_param_param_token1),
	4244: uint16(26),
	4245: uint16(1),
	4246: uint16(sym_mount_param_param),
	4247: uint16(2),
	4248: uint16(5),
	4249: uint16(2),
	4250: uint16(sym_line_continuation),
	4251: uint16(sym_comment),
	4252: uint16(885),
	4253: uint16(2),
	4254: uint16(sym_heredoc_nl),
	4255: uint16(anon_sym_LF),
	4256: uint16(2),
	4257: uint16(3),
	4258: uint16(2),
	4259: uint16(sym_line_continuation),
	4260: uint16(sym_comment),
	4261: uint16(887),
	4262: uint16(2),
	4263: uint16(anon_sym_COMMA2),
	4264: uint16(anon_sym_RBRACK),
	4265: uint16(3),
	4266: uint16(802),
	4267: uint16(1),
	4268: uint16(anon_sym_DQUOTE),
	4269: uint16(244),
	4270: uint16(1),
	4271: uint16(sym_json_string),
	4272: uint16(3),
	4273: uint16(2),
	4274: uint16(sym_line_continuation),
	4275: uint16(sym_comment),
	4276: uint16(2),
	4277: uint16(5),
	4278: uint16(2),
	4279: uint16(sym_line_continuation),
	4280: uint16(sym_comment),
	4281: uint16(889),
	4282: uint16(2),
	4283: uint16(sym_heredoc_nl),
	4284: uint16(anon_sym_LF),
	4285: uint16(2),
	4286: uint16(3),
	4287: uint16(2),
	4288: uint16(sym_line_continuation),
	4289: uint16(sym_comment),
	4290: uint16(461),
	4291: uint16(2),
	4292: uint16(aux_sym_cmd_instruction_token1),
	4293: uint16(anon_sym_DASH_DASH),
	4294: uint16(2),
	4295: uint16(5),
	4296: uint16(2),
	4297: uint16(sym_line_continuation),
	4298: uint16(sym_comment),
	4299: uint16(891),
	4300: uint16(2),
	4301: uint16(sym_heredoc_nl),
	4302: uint16(anon_sym_LF),
	4303: uint16(3),
	4304: uint16(140),
	4305: uint16(1),
	4306: uint16(anon_sym_LBRACK),
	4307: uint16(273),
	4308: uint16(1),
	4309: uint16(sym_json_string_array),
	4310: uint16(3),
	4311: uint16(2),
	4312: uint16(sym_line_continuation),
	4313: uint16(sym_comment),
	4314: uint16(4),
	4315: uint16(3),
	4316: uint16(1),
	4317: uint16(sym_line_continuation),
	4318: uint16(5),
	4319: uint16(1),
	4320: uint16(sym_comment),
	4321: uint16(883),
	4322: uint16(1),
	4323: uint16(aux_sym_mount_param_param_token1),
	4324: uint16(42),
	4325: uint16(1),
	4326: uint16(sym_mount_param_param),
	4327: uint16(2),
	4328: uint16(5),
	4329: uint16(2),
	4330: uint16(sym_line_continuation),
	4331: uint16(sym_comment),
	4332: uint16(893),
	4333: uint16(2),
	4334: uint16(sym_heredoc_nl),
	4335: uint16(anon_sym_LF),
	4336: uint16(3),
	4337: uint16(895),
	4338: uint16(1),
	4339: uint16(anon_sym_LF),
	4340: uint16(897),
	4341: uint16(1),
	4342: uint16(aux_sym_from_instruction_token2),
	4343: uint16(5),
	4344: uint16(2),
	4345: uint16(sym_line_continuation),
	4346: uint16(sym_comment),
	4347: uint16(3),
	4348: uint16(899),
	4349: uint16(1),
	4350: uint16(anon_sym_LF),
	4351: uint16(901),
	4352: uint16(1),
	4353: uint16(anon_sym_EQ),
	4354: uint16(5),
	4355: uint16(2),
	4356: uint16(sym_line_continuation),
	4357: uint16(sym_comment),
	4358: uint16(3),
	4359: uint16(903),
	4360: uint16(1),
	4361: uint16(anon_sym_LF),
	4362: uint16(905),
	4363: uint16(1),
	4364: uint16(aux_sym__env_key_token1),
	4365: uint16(5),
	4366: uint16(2),
	4367: uint16(sym_line_continuation),
	4368: uint16(sym_comment),
	4369: uint16(2),
	4370: uint16(3),
	4371: uint16(2),
	4372: uint16(sym_line_continuation),
	4373: uint16(sym_comment),
	4374: uint16(907),
	4375: uint16(2),
	4376: uint16(anon_sym_COMMA2),
	4377: uint16(anon_sym_RBRACK),
	4378: uint16(3),
	4379: uint16(824),
	4380: uint16(1),
	4381: uint16(anon_sym_LF),
	4382: uint16(909),
	4383: uint16(1),
	4384: uint16(sym__non_newline_whitespace),
	4385: uint16(5),
	4386: uint16(2),
	4387: uint16(sym_line_continuation),
	4388: uint16(sym_comment),
	4389: uint16(3),
	4390: uint16(911),
	4391: uint16(1),
	4392: uint16(anon_sym_EQ),
	4393: uint16(913),
	4394: uint16(1),
	4395: uint16(aux_sym__spaced_env_pair_token1),
	4396: uint16(5),
	4397: uint16(2),
	4398: uint16(sym_line_continuation),
	4399: uint16(sym_comment),
	4400: uint16(3),
	4401: uint16(661),
	4402: uint16(1),
	4403: uint16(anon_sym_LF),
	4404: uint16(794),
	4405: uint16(1),
	4406: uint16(sym_required_line_continuation),
	4407: uint16(5),
	4408: uint16(2),
	4409: uint16(sym_line_continuation),
	4410: uint16(sym_comment),
	4411: uint16(2),
	4412: uint16(3),
	4413: uint16(2),
	4414: uint16(sym_line_continuation),
	4415: uint16(sym_comment),
	4416: uint16(758),
	4417: uint16(2),
	4418: uint16(anon_sym_COMMA2),
	4419: uint16(anon_sym_RBRACK),
	4420: uint16(3),
	4421: uint16(915),
	4422: uint16(1),
	4423: uint16(anon_sym_LF),
	4424: uint16(917),
	4425: uint16(1),
	4426: uint16(aux_sym_from_instruction_token2),
	4427: uint16(5),
	4428: uint16(2),
	4429: uint16(sym_line_continuation),
	4430: uint16(sym_comment),
	4431: uint16(3),
	4432: uint16(919),
	4433: uint16(1),
	4434: uint16(anon_sym_LF),
	4435: uint16(921),
	4436: uint16(1),
	4437: uint16(aux_sym_from_instruction_token2),
	4438: uint16(5),
	4439: uint16(2),
	4440: uint16(sym_line_continuation),
	4441: uint16(sym_comment),
	4442: uint16(3),
	4443: uint16(923),
	4444: uint16(1),
	4445: uint16(anon_sym_LF),
	4446: uint16(925),
	4447: uint16(1),
	4448: uint16(anon_sym_COLON),
	4449: uint16(5),
	4450: uint16(2),
	4451: uint16(sym_line_continuation),
	4452: uint16(sym_comment),
	4453: uint16(2),
	4454: uint16(927),
	4455: uint16(1),
	4456: uint16(aux_sym__expansion_body_token1),
	4457: uint16(5),
	4458: uint16(2),
	4459: uint16(sym_line_continuation),
	4460: uint16(sym_comment),
	4461: uint16(2),
	4462: uint16(929),
	4463: uint16(1),
	4464: uint16(anon_sym_EQ),
	4465: uint16(3),
	4466: uint16(2),
	4467: uint16(sym_line_continuation),
	4468: uint16(sym_comment),
	4469: uint16(2),
	4470: uint16(931),
	4471: uint16(1),
	4472: uint16(anon_sym_LF),
	4473: uint16(5),
	4474: uint16(2),
	4475: uint16(sym_line_continuation),
	4476: uint16(sym_comment),
	4477: uint16(2),
	4478: uint16(933),
	4479: uint16(1),
	4480: uint16(aux_sym__expansion_body_token1),
	4481: uint16(5),
	4482: uint16(2),
	4483: uint16(sym_line_continuation),
	4484: uint16(sym_comment),
	4485: uint16(2),
	4486: uint16(737),
	4487: uint16(1),
	4488: uint16(anon_sym_EQ),
	4489: uint16(3),
	4490: uint16(2),
	4491: uint16(sym_line_continuation),
	4492: uint16(sym_comment),
	4493: uint16(2),
	4494: uint16(935),
	4495: uint16(1),
	4497: uint16(3),
	4498: uint16(2),
	4499: uint16(sym_line_continuation),
	4500: uint16(sym_comment),
	4501: uint16(2),
	4502: uint16(741),
	4503: uint16(1),
	4504: uint16(anon_sym_EQ),
	4505: uint16(3),
	4506: uint16(2),
	4507: uint16(sym_line_continuation),
	4508: uint16(sym_comment),
	4509: uint16(2),
	4510: uint16(937),
	4511: uint16(1),
	4512: uint16(anon_sym_LF),
	4513: uint16(5),
	4514: uint16(2),
	4515: uint16(sym_line_continuation),
	4516: uint16(sym_comment),
	4517: uint16(2),
	4518: uint16(939),
	4519: uint16(1),
	4520: uint16(anon_sym_LF),
	4521: uint16(5),
	4522: uint16(2),
	4523: uint16(sym_line_continuation),
	4524: uint16(sym_comment),
	4525: uint16(2),
	4526: uint16(941),
	4527: uint16(1),
	4528: uint16(anon_sym_LF),
	4529: uint16(5),
	4530: uint16(2),
	4531: uint16(sym_line_continuation),
	4532: uint16(sym_comment),
	4533: uint16(2),
	4534: uint16(943),
	4535: uint16(1),
	4536: uint16(anon_sym_LF),
	4537: uint16(5),
	4538: uint16(2),
	4539: uint16(sym_line_continuation),
	4540: uint16(sym_comment),
	4541: uint16(2),
	4542: uint16(945),
	4543: uint16(1),
	4544: uint16(aux_sym_maintainer_instruction_token2),
	4545: uint16(5),
	4546: uint16(2),
	4547: uint16(sym_line_continuation),
	4548: uint16(sym_comment),
	4549: uint16(2),
	4550: uint16(947),
	4551: uint16(1),
	4552: uint16(aux_sym_shell_fragment_token1),
	4553: uint16(5),
	4554: uint16(2),
	4555: uint16(sym_line_continuation),
	4556: uint16(sym_comment),
	4557: uint16(2),
	4558: uint16(949),
	4559: uint16(1),
	4560: uint16(anon_sym_LF),
	4561: uint16(5),
	4562: uint16(2),
	4563: uint16(sym_line_continuation),
	4564: uint16(sym_comment),
	4565: uint16(2),
	4566: uint16(951),
	4567: uint16(1),
	4568: uint16(aux_sym_arg_instruction_token2),
	4569: uint16(3),
	4570: uint16(2),
	4571: uint16(sym_line_continuation),
	4572: uint16(sym_comment),
	4573: uint16(2),
	4574: uint16(953),
	4575: uint16(1),
	4576: uint16(anon_sym_LF),
	4577: uint16(5),
	4578: uint16(2),
	4579: uint16(sym_line_continuation),
	4580: uint16(sym_comment),
	4581: uint16(2),
	4582: uint16(955),
	4583: uint16(1),
	4584: uint16(anon_sym_EQ),
	4585: uint16(3),
	4586: uint16(2),
	4587: uint16(sym_line_continuation),
	4588: uint16(sym_comment),
	4589: uint16(2),
	4590: uint16(957),
	4591: uint16(1),
	4592: uint16(anon_sym_LF),
	4593: uint16(5),
	4594: uint16(2),
	4595: uint16(sym_line_continuation),
	4596: uint16(sym_comment),
	4597: uint16(2),
	4598: uint16(959),
	4599: uint16(1),
	4600: uint16(anon_sym_EQ),
	4601: uint16(3),
	4602: uint16(2),
	4603: uint16(sym_line_continuation),
	4604: uint16(sym_comment),
	4605: uint16(2),
	4606: uint16(961),
	4607: uint16(1),
	4608: uint16(anon_sym_RBRACE),
	4609: uint16(3),
	4610: uint16(2),
	4611: uint16(sym_line_continuation),
	4612: uint16(sym_comment),
	4613: uint16(3),
	4614: uint16(3),
	4615: uint16(1),
	4616: uint16(sym_line_continuation),
	4617: uint16(5),
	4618: uint16(1),
	4619: uint16(sym_comment),
	4620: uint16(963),
	4621: uint16(1),
	4622: uint16(aux_sym_param_token2),
	4623: uint16(2),
	4624: uint16(893),
	4625: uint16(1),
	4626: uint16(anon_sym_LF),
	4627: uint16(5),
	4628: uint16(2),
	4629: uint16(sym_line_continuation),
	4630: uint16(sym_comment),
	4631: uint16(2),
	4632: uint16(965),
	4633: uint16(1),
	4634: uint16(aux_sym_maintainer_instruction_token2),
	4635: uint16(5),
	4636: uint16(2),
	4637: uint16(sym_line_continuation),
	4638: uint16(sym_comment),
	4639: uint16(2),
	4640: uint16(911),
	4641: uint16(1),
	4642: uint16(anon_sym_EQ),
	4643: uint16(3),
	4644: uint16(2),
	4645: uint16(sym_line_continuation),
	4646: uint16(sym_comment),
	4647: uint16(2),
	4648: uint16(710),
	4649: uint16(1),
	4650: uint16(anon_sym_EQ),
	4651: uint16(3),
	4652: uint16(2),
	4653: uint16(sym_line_continuation),
	4654: uint16(sym_comment),
	4655: uint16(2),
	4656: uint16(967),
	4657: uint16(1),
	4658: uint16(anon_sym_LF),
	4659: uint16(5),
	4660: uint16(2),
	4661: uint16(sym_line_continuation),
	4662: uint16(sym_comment),
	4663: uint16(2),
	4664: uint16(969),
	4665: uint16(1),
	4666: uint16(anon_sym_EQ),
	4667: uint16(3),
	4668: uint16(2),
	4669: uint16(sym_line_continuation),
	4670: uint16(sym_comment),
	4671: uint16(2),
	4672: uint16(971),
	4673: uint16(1),
	4674: uint16(anon_sym_LF),
	4675: uint16(5),
	4676: uint16(2),
	4677: uint16(sym_line_continuation),
	4678: uint16(sym_comment),
	4679: uint16(2),
	4680: uint16(973),
	4681: uint16(1),
	4682: uint16(anon_sym_RBRACE),
	4683: uint16(3),
	4684: uint16(2),
	4685: uint16(sym_line_continuation),
	4686: uint16(sym_comment),
	4687: uint16(3),
	4688: uint16(3),
	4689: uint16(1),
	4690: uint16(sym_line_continuation),
	4691: uint16(5),
	4692: uint16(1),
	4693: uint16(sym_comment),
	4694: uint16(975),
	4695: uint16(1),
	4696: uint16(aux_sym_param_token2),
	4697: uint16(2),
	4698: uint16(977),
	4699: uint16(1),
	4700: uint16(anon_sym_LF),
	4701: uint16(5),
	4702: uint16(2),
	4703: uint16(sym_line_continuation),
	4704: uint16(sym_comment),
	4705: uint16(2),
	4706: uint16(718),
	4707: uint16(1),
	4708: uint16(anon_sym_EQ),
	4709: uint16(3),
	4710: uint16(2),
	4711: uint16(sym_line_continuation),
	4712: uint16(sym_comment),
	4713: uint16(2),
	4714: uint16(979),
	4715: uint16(1),
	4716: uint16(anon_sym_RBRACE),
	4717: uint16(3),
	4718: uint16(2),
	4719: uint16(sym_line_continuation),
	4720: uint16(sym_comment),
	4721: uint16(3),
	4722: uint16(3),
	4723: uint16(1),
	4724: uint16(sym_line_continuation),
	4725: uint16(5),
	4726: uint16(1),
	4727: uint16(sym_comment),
	4728: uint16(981),
	4729: uint16(1),
	4730: uint16(aux_sym_param_token2),
	4731: uint16(2),
	4732: uint16(983),
	4733: uint16(1),
	4734: uint16(anon_sym_LF),
	4735: uint16(5),
	4736: uint16(2),
	4737: uint16(sym_line_continuation),
	4738: uint16(sym_comment),
	4739: uint16(2),
	4740: uint16(985),
	4741: uint16(1),
	4742: uint16(anon_sym_RBRACE),
	4743: uint16(3),
	4744: uint16(2),
	4745: uint16(sym_line_continuation),
	4746: uint16(sym_comment),
	4747: uint16(2),
	4748: uint16(987),
	4749: uint16(1),
	4750: uint16(anon_sym_RBRACE),
	4751: uint16(3),
	4752: uint16(2),
	4753: uint16(sym_line_continuation),
	4754: uint16(sym_comment),
	4755: uint16(3),
	4756: uint16(3),
	4757: uint16(1),
	4758: uint16(sym_line_continuation),
	4759: uint16(5),
	4760: uint16(1),
	4761: uint16(sym_comment),
	4762: uint16(989),
	4763: uint16(1),
	4764: uint16(aux_sym_mount_param_param_token1),
	4765: uint16(2),
	4766: uint16(991),
	4767: uint16(1),
	4768: uint16(aux_sym_shell_fragment_token1),
	4769: uint16(5),
	4770: uint16(2),
	4771: uint16(sym_line_continuation),
	4772: uint16(sym_comment),
	4773: uint16(2),
	4774: uint16(993),
	4775: uint16(1),
	4776: uint16(anon_sym_RBRACE),
	4777: uint16(3),
	4778: uint16(2),
	4779: uint16(sym_line_continuation),
	4780: uint16(sym_comment),
	4781: uint16(2),
	4782: uint16(873),
	4783: uint16(1),
	4784: uint16(anon_sym_LF),
	4785: uint16(5),
	4786: uint16(2),
	4787: uint16(sym_line_continuation),
	4788: uint16(sym_comment),
	4789: uint16(2),
	4790: uint16(889),
	4791: uint16(1),
	4792: uint16(anon_sym_LF),
	4793: uint16(5),
	4794: uint16(2),
	4795: uint16(sym_line_continuation),
	4796: uint16(sym_comment),
	4797: uint16(2),
	4798: uint16(995),
	4799: uint16(1),
	4800: uint16(anon_sym_RBRACE),
	4801: uint16(3),
	4802: uint16(2),
	4803: uint16(sym_line_continuation),
	4804: uint16(sym_comment),
	4805: uint16(2),
	4806: uint16(997),
	4807: uint16(1),
	4808: uint16(anon_sym_LF),
	4809: uint16(5),
	4810: uint16(2),
	4811: uint16(sym_line_continuation),
	4812: uint16(sym_comment),
	4813: uint16(2),
	4814: uint16(999),
	4815: uint16(1),
	4816: uint16(anon_sym_LF),
	4817: uint16(5),
	4818: uint16(2),
	4819: uint16(sym_line_continuation),
	4820: uint16(sym_comment),
	4821: uint16(2),
	4822: uint16(1001),
	4823: uint16(1),
	4824: uint16(anon_sym_RBRACE),
	4825: uint16(3),
	4826: uint16(2),
	4827: uint16(sym_line_continuation),
	4828: uint16(sym_comment),
	4829: uint16(2),
	4830: uint16(328),
	4831: uint16(1),
	4832: uint16(sym__non_newline_whitespace),
	4833: uint16(5),
	4834: uint16(2),
	4835: uint16(sym_line_continuation),
	4836: uint16(sym_comment),
	4837: uint16(2),
	4838: uint16(1003),
	4839: uint16(1),
	4840: uint16(anon_sym_LF),
	4841: uint16(5),
	4842: uint16(2),
	4843: uint16(sym_line_continuation),
	4844: uint16(sym_comment),
	4845: uint16(2),
	4846: uint16(1005),
	4847: uint16(1),
	4848: uint16(anon_sym_RBRACE),
	4849: uint16(3),
	4850: uint16(2),
	4851: uint16(sym_line_continuation),
	4852: uint16(sym_comment),
	4853: uint16(2),
	4854: uint16(1007),
	4855: uint16(1),
	4856: uint16(anon_sym_LF),
	4857: uint16(5),
	4858: uint16(2),
	4859: uint16(sym_line_continuation),
	4860: uint16(sym_comment),
	4861: uint16(2),
	4862: uint16(1009),
	4863: uint16(1),
	4864: uint16(anon_sym_RBRACE),
	4865: uint16(3),
	4866: uint16(2),
	4867: uint16(sym_line_continuation),
	4868: uint16(sym_comment),
	4869: uint16(2),
	4870: uint16(580),
	4871: uint16(1),
	4872: uint16(sym__non_newline_whitespace),
	4873: uint16(5),
	4874: uint16(2),
	4875: uint16(sym_line_continuation),
	4876: uint16(sym_comment),
	4877: uint16(2),
	4878: uint16(1011),
	4879: uint16(1),
	4880: uint16(anon_sym_RBRACE),
	4881: uint16(3),
	4882: uint16(2),
	4883: uint16(sym_line_continuation),
	4884: uint16(sym_comment),
	4885: uint16(2),
	4886: uint16(1013),
	4887: uint16(1),
	4888: uint16(anon_sym_LF),
	4889: uint16(5),
	4890: uint16(2),
	4891: uint16(sym_line_continuation),
	4892: uint16(sym_comment),
	4893: uint16(2),
	4894: uint16(1015),
	4895: uint16(1),
	4896: uint16(anon_sym_RBRACE),
	4897: uint16(3),
	4898: uint16(2),
	4899: uint16(sym_line_continuation),
	4900: uint16(sym_comment),
	4901: uint16(2),
	4902: uint16(1017),
	4903: uint16(1),
	4904: uint16(aux_sym_param_token1),
	4905: uint16(3),
	4906: uint16(2),
	4907: uint16(sym_line_continuation),
	4908: uint16(sym_comment),
	4909: uint16(2),
	4910: uint16(1019),
	4911: uint16(1),
	4912: uint16(anon_sym_RBRACE),
	4913: uint16(3),
	4914: uint16(2),
	4915: uint16(sym_line_continuation),
	4916: uint16(sym_comment),
	4917: uint16(2),
	4918: uint16(1021),
	4919: uint16(1),
	4920: uint16(anon_sym_RBRACE),
	4921: uint16(3),
	4922: uint16(2),
	4923: uint16(sym_line_continuation),
	4924: uint16(sym_comment),
	4925: uint16(2),
	4926: uint16(1023),
	4927: uint16(1),
	4928: uint16(aux_sym__expansion_body_token1),
	4929: uint16(5),
	4930: uint16(2),
	4931: uint16(sym_line_continuation),
	4932: uint16(sym_comment),
	4933: uint16(2),
	4934: uint16(1025),
	4935: uint16(1),
	4936: uint16(anon_sym_EQ),
	4937: uint16(3),
	4938: uint16(2),
	4939: uint16(sym_line_continuation),
	4940: uint16(sym_comment),
	4941: uint16(2),
	4942: uint16(1027),
	4943: uint16(1),
	4944: uint16(aux_sym__expansion_body_token1),
	4945: uint16(5),
	4946: uint16(2),
	4947: uint16(sym_line_continuation),
	4948: uint16(sym_comment),
	4949: uint16(2),
	4950: uint16(1029),
	4951: uint16(1),
	4952: uint16(anon_sym_EQ),
	4953: uint16(3),
	4954: uint16(2),
	4955: uint16(sym_line_continuation),
	4956: uint16(sym_comment),
	4957: uint16(2),
	4958: uint16(1031),
	4959: uint16(1),
	4960: uint16(aux_sym__expansion_body_token1),
	4961: uint16(5),
	4962: uint16(2),
	4963: uint16(sym_line_continuation),
	4964: uint16(sym_comment),
	4965: uint16(2),
	4966: uint16(1033),
	4967: uint16(1),
	4968: uint16(anon_sym_EQ),
	4969: uint16(3),
	4970: uint16(2),
	4971: uint16(sym_line_continuation),
	4972: uint16(sym_comment),
	4973: uint16(2),
	4974: uint16(1035),
	4975: uint16(1),
	4976: uint16(aux_sym__expansion_body_token1),
	4977: uint16(5),
	4978: uint16(2),
	4979: uint16(sym_line_continuation),
	4980: uint16(sym_comment),
	4981: uint16(2),
	4982: uint16(1037),
	4983: uint16(1),
	4984: uint16(anon_sym_LF),
	4985: uint16(5),
	4986: uint16(2),
	4987: uint16(sym_line_continuation),
	4988: uint16(sym_comment),
	4989: uint16(2),
	4990: uint16(1039),
	4991: uint16(1),
	4992: uint16(aux_sym__expansion_body_token1),
	4993: uint16(5),
	4994: uint16(2),
	4995: uint16(sym_line_continuation),
	4996: uint16(sym_comment),
	4997: uint16(2),
	4998: uint16(1041),
	4999: uint16(1),
	5000: uint16(aux_sym__expansion_body_token1),
	5001: uint16(5),
	5002: uint16(2),
	5003: uint16(sym_line_continuation),
	5004: uint16(sym_comment),
	5005: uint16(2),
	5006: uint16(1043),
	5007: uint16(1),
	5008: uint16(aux_sym__expansion_body_token1),
	5009: uint16(5),
	5010: uint16(2),
	5011: uint16(sym_line_continuation),
	5012: uint16(sym_comment),
	5013: uint16(2),
	5014: uint16(1045),
	5015: uint16(1),
	5016: uint16(aux_sym__expansion_body_token1),
	5017: uint16(5),
	5018: uint16(2),
	5019: uint16(sym_line_continuation),
	5020: uint16(sym_comment),
	5021: uint16(2),
	5022: uint16(1047),
	5023: uint16(1),
	5024: uint16(aux_sym__expansion_body_token1),
	5025: uint16(5),
	5026: uint16(2),
	5027: uint16(sym_line_continuation),
	5028: uint16(sym_comment),
	5029: uint16(2),
	5030: uint16(1049),
	5031: uint16(1),
	5032: uint16(aux_sym__expansion_body_token1),
	5033: uint16(5),
	5034: uint16(2),
	5035: uint16(sym_line_continuation),
	5036: uint16(sym_comment),
	5037: uint16(2),
	5038: uint16(1051),
	5039: uint16(1),
	5040: uint16(aux_sym__expansion_body_token1),
	5041: uint16(5),
	5042: uint16(2),
	5043: uint16(sym_line_continuation),
	5044: uint16(sym_comment),
	5045: uint16(2),
	5046: uint16(1053),
	5047: uint16(1),
	5048: uint16(aux_sym__expansion_body_token1),
	5049: uint16(5),
	5050: uint16(2),
	5051: uint16(sym_line_continuation),
	5052: uint16(sym_comment),
	5053: uint16(2),
	5054: uint16(1055),
	5055: uint16(1),
	5056: uint16(aux_sym_param_token1),
	5057: uint16(3),
	5058: uint16(2),
	5059: uint16(sym_line_continuation),
	5060: uint16(sym_comment),
	5061: uint16(2),
	5062: uint16(1057),
	5063: uint16(1),
	5064: uint16(aux_sym_param_token1),
	5065: uint16(3),
	5066: uint16(2),
	5067: uint16(sym_line_continuation),
	5068: uint16(sym_comment),
	5069: uint16(3),
	5070: uint16(3),
	5071: uint16(1),
	5072: uint16(sym_line_continuation),
	5073: uint16(5),
	5074: uint16(1),
	5075: uint16(sym_comment),
	5076: uint16(1059),
	5077: uint16(1),
	5078: uint16(aux_sym_param_token2),
}

var ts_small_parse_table_map = [323]uint32_t{
	1:   uint32(90),
	2:   uint32(180),
	3:   uint32(264),
	4:   uint32(291),
	5:   uint32(326),
	6:   uint32(361),
	7:   uint32(395),
	8:   uint32(425),
	9:   uint32(455),
	10:  uint32(485),
	11:  uint32(514),
	12:  uint32(549),
	13:  uint32(584),
	14:  uint32(619),
	15:  uint32(648),
	16:  uint32(683),
	17:  uint32(708),
	18:  uint32(731),
	19:  uint32(756),
	20:  uint32(780),
	21:  uint32(804),
	22:  uint32(828),
	23:  uint32(852),
	24:  uint32(876),
	25:  uint32(898),
	26:  uint32(920),
	27:  uint32(942),
	28:  uint32(965),
	29:  uint32(988),
	30:  uint32(1009),
	31:  uint32(1026),
	32:  uint32(1049),
	33:  uint32(1074),
	34:  uint32(1095),
	35:  uint32(1118),
	36:  uint32(1141),
	37:  uint32(1162),
	38:  uint32(1185),
	39:  uint32(1210),
	40:  uint32(1227),
	41:  uint32(1244),
	42:  uint32(1267),
	43:  uint32(1290),
	44:  uint32(1313),
	45:  uint32(1336),
	46:  uint32(1359),
	47:  uint32(1382),
	48:  uint32(1405),
	49:  uint32(1428),
	50:  uint32(1451),
	51:  uint32(1474),
	52:  uint32(1496),
	53:  uint32(1524),
	54:  uint32(1546),
	55:  uint32(1574),
	56:  uint32(1596),
	57:  uint32(1618),
	58:  uint32(1640),
	59:  uint32(1662),
	60:  uint32(1684),
	61:  uint32(1700),
	62:  uint32(1722),
	63:  uint32(1744),
	64:  uint32(1772),
	65:  uint32(1800),
	66:  uint32(1822),
	67:  uint32(1838),
	68:  uint32(1866),
	69:  uint32(1888),
	70:  uint32(1916),
	71:  uint32(1932),
	72:  uint32(1948),
	73:  uint32(1976),
	74:  uint32(1991),
	75:  uint32(2010),
	76:  uint32(2029),
	77:  uint32(2044),
	78:  uint32(2063),
	79:  uint32(2078),
	80:  uint32(2097),
	81:  uint32(2116),
	82:  uint32(2139),
	83:  uint32(2158),
	84:  uint32(2177),
	85:  uint32(2196),
	86:  uint32(2215),
	87:  uint32(2236),
	88:  uint32(2255),
	89:  uint32(2274),
	90:  uint32(2293),
	91:  uint32(2312),
	92:  uint32(2331),
	93:  uint32(2350),
	94:  uint32(2364),
	95:  uint32(2382),
	96:  uint32(2400),
	97:  uint32(2418),
	98:  uint32(2434),
	99:  uint32(2452),
	100: uint32(2474),
	101: uint32(2492),
	102: uint32(2510),
	103: uint32(2526),
	104: uint32(2548),
	105: uint32(2566),
	106: uint32(2584),
	107: uint32(2600),
	108: uint32(2614),
	109: uint32(2636),
	110: uint32(2652),
	111: uint32(2670),
	112: uint32(2684),
	113: uint32(2698),
	114: uint32(2712),
	115: uint32(2730),
	116: uint32(2744),
	117: uint32(2758),
	118: uint32(2772),
	119: uint32(2786),
	120: uint32(2804),
	121: uint32(2822),
	122: uint32(2840),
	123: uint32(2858),
	124: uint32(2876),
	125: uint32(2894),
	126: uint32(2907),
	127: uint32(2922),
	128: uint32(2937),
	129: uint32(2952),
	130: uint32(2967),
	131: uint32(2984),
	132: uint32(2997),
	133: uint32(3010),
	134: uint32(3027),
	135: uint32(3040),
	136: uint32(3051),
	137: uint32(3062),
	138: uint32(3073),
	139: uint32(3090),
	140: uint32(3105),
	141: uint32(3120),
	142: uint32(3137),
	143: uint32(3154),
	144: uint32(3171),
	145: uint32(3186),
	146: uint32(3201),
	147: uint32(3214),
	148: uint32(3227),
	149: uint32(3242),
	150: uint32(3255),
	151: uint32(3270),
	152: uint32(3285),
	153: uint32(3300),
	154: uint32(3317),
	155: uint32(3332),
	156: uint32(3349),
	157: uint32(3364),
	158: uint32(3377),
	159: uint32(3390),
	160: uint32(3403),
	161: uint32(3417),
	162: uint32(3429),
	163: uint32(3439),
	164: uint32(3453),
	165: uint32(3467),
	166: uint32(3481),
	167: uint32(3493),
	168: uint32(3507),
	169: uint32(3519),
	170: uint32(3533),
	171: uint32(3547),
	172: uint32(3561),
	173: uint32(3575),
	174: uint32(3587),
	175: uint32(3599),
	176: uint32(3611),
	177: uint32(3623),
	178: uint32(3637),
	179: uint32(3651),
	180: uint32(3663),
	181: uint32(3675),
	182: uint32(3687),
	183: uint32(3697),
	184: uint32(3707),
	185: uint32(3721),
	186: uint32(3733),
	187: uint32(3743),
	188: uint32(3757),
	189: uint32(3771),
	190: uint32(3785),
	191: uint32(3795),
	192: uint32(3805),
	193: uint32(3819),
	194: uint32(3833),
	195: uint32(3847),
	196: uint32(3861),
	197: uint32(3875),
	198: uint32(3889),
	199: uint32(3903),
	200: uint32(3917),
	201: uint32(3931),
	202: uint32(3945),
	203: uint32(3959),
	204: uint32(3973),
	205: uint32(3987),
	206: uint32(4001),
	207: uint32(4015),
	208: uint32(4029),
	209: uint32(4043),
	210: uint32(4057),
	211: uint32(4071),
	212: uint32(4085),
	213: uint32(4099),
	214: uint32(4113),
	215: uint32(4127),
	216: uint32(4137),
	217: uint32(4146),
	218: uint32(4157),
	219: uint32(4168),
	220: uint32(4181),
	221: uint32(4192),
	222: uint32(4203),
	223: uint32(4212),
	224: uint32(4223),
	225: uint32(4234),
	226: uint32(4247),
	227: uint32(4256),
	228: uint32(4265),
	229: uint32(4276),
	230: uint32(4285),
	231: uint32(4294),
	232: uint32(4303),
	233: uint32(4314),
	234: uint32(4327),
	235: uint32(4336),
	236: uint32(4347),
	237: uint32(4358),
	238: uint32(4369),
	239: uint32(4378),
	240: uint32(4389),
	241: uint32(4400),
	242: uint32(4411),
	243: uint32(4420),
	244: uint32(4431),
	245: uint32(4442),
	246: uint32(4453),
	247: uint32(4461),
	248: uint32(4469),
	249: uint32(4477),
	250: uint32(4485),
	251: uint32(4493),
	252: uint32(4501),
	253: uint32(4509),
	254: uint32(4517),
	255: uint32(4525),
	256: uint32(4533),
	257: uint32(4541),
	258: uint32(4549),
	259: uint32(4557),
	260: uint32(4565),
	261: uint32(4573),
	262: uint32(4581),
	263: uint32(4589),
	264: uint32(4597),
	265: uint32(4605),
	266: uint32(4613),
	267: uint32(4623),
	268: uint32(4631),
	269: uint32(4639),
	270: uint32(4647),
	271: uint32(4655),
	272: uint32(4663),
	273: uint32(4671),
	274: uint32(4679),
	275: uint32(4687),
	276: uint32(4697),
	277: uint32(4705),
	278: uint32(4713),
	279: uint32(4721),
	280: uint32(4731),
	281: uint32(4739),
	282: uint32(4747),
	283: uint32(4755),
	284: uint32(4765),
	285: uint32(4773),
	286: uint32(4781),
	287: uint32(4789),
	288: uint32(4797),
	289: uint32(4805),
	290: uint32(4813),
	291: uint32(4821),
	292: uint32(4829),
	293: uint32(4837),
	294: uint32(4845),
	295: uint32(4853),
	296: uint32(4861),
	297: uint32(4869),
	298: uint32(4877),
	299: uint32(4885),
	300: uint32(4893),
	301: uint32(4901),
	302: uint32(4909),
	303: uint32(4917),
	304: uint32(4925),
	305: uint32(4933),
	306: uint32(4941),
	307: uint32(4949),
	308: uint32(4957),
	309: uint32(4965),
	310: uint32(4973),
	311: uint32(4981),
	312: uint32(4989),
	313: uint32(4997),
	314: uint32(5005),
	315: uint32(5013),
	316: uint32(5021),
	317: uint32(5029),
	318: uint32(5037),
	319: uint32(5045),
	320: uint32(5053),
	321: uint32(5061),
	322: uint32(5069),
}

var ts_parse_actions = [1061]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_source_file),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(71)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(61)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(104)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(102)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(54)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(145)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(262)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(132)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(88)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(234)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(259)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(270)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_source_file),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(71)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(61)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(104)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(102)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(54)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(145)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(262)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(132)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(88)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(234)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(259)),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(270)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(220)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(191)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(286)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_env_pair),
		Fproduction_id: uint16(1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(215)),
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
		Fcount: uint8(1),
	}})),
	121: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_env_pair),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(125)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(214)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(122)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(195)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(260)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(201)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(322)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(294)),
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
		Fsymbol:      uint16(sym_unquoted_string),
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
		Fcount: uint8(1),
	}})),
	157: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_unquoted_string),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		Fcount: uint8(1),
	}})),
	161: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_run_instruction_repeat1),
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
		Fsymbol:      uint16(aux_sym_run_instruction_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(220)),
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
		Fsymbol:      uint16(aux_sym_run_instruction_repeat1),
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
		Fsymbol:      uint16(aux_sym_unquoted_string_repeat1),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_unquoted_string_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(214)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	173: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_unquoted_string_repeat1),
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
		Fsymbol:      uint16(aux_sym_unquoted_string_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_image_name),
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
		Fcount: uint8(1),
	}})),
	180: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_image_name),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(172)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_shell_fragment),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_shell_fragment),
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
		Fsymbol:      uint16(aux_sym_shell_fragment_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_shell_fragment_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(23)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	197: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shell_fragment_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	199: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shell_fragment_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(286)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_image_name),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_image_name),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_image_name_repeat1),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_image_name_repeat1),
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
		Fsymbol:      uint16(aux_sym_image_name_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(172)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_image_name_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(25)),
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
		Fcount: uint8(1),
	}})),
	218: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_mount_param),
		Fproduction_id: uint16(13),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_mount_param),
		Fproduction_id: uint16(13),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(235)),
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
		Fcount: uint8(1),
	}})),
	224: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_mount_param),
		Fproduction_id: uint16(15),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_mount_param),
		Fproduction_id: uint16(15),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_mount_param_repeat1),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_mount_param_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	232: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_mount_param_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(235)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_image_tag),
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
		Fsymbol:      uint16(sym_image_tag),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(210)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(171)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(225)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__user_name_or_group),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(35)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(204)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(252)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_label_instruction_repeat1),
	})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_label_instruction_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(249)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_label_instruction_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(36)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_label_instruction_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(112)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__user_name_or_group),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	272: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(272)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fsymbol:      uint16(aux_sym_double_quoted_string_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(171)),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_double_quoted_string_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_double_quoted_string_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fsymbol:      uint16(aux_sym_double_quoted_string_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__user_name_or_group_repeat1),
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
		Fsymbol:      uint16(aux_sym__user_name_or_group_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(38)),
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
		Fsymbol:      uint16(aux_sym__user_name_or_group_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(204)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_image_tag_repeat1),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_image_tag_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_image_tag_repeat1),
	})))),
	304: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(210)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	306: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_image_tag_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(39)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_label_instruction),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(249)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
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
		Fsymbol:      uint16(sym_mount_param_param),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_mount_param_param),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_path_with_heredoc),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(216)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_path_with_heredoc),
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
		Fcount: uint8(2),
	}})),
	331: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shell_fragment_repeat1),
	})))),
	332: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_shell_fragment_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(260)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_path_with_heredoc),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_path_with_heredoc),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_unquoted_string_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(215)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_unquoted_string_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(48)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	351: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
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
		Fcount: uint8(2),
	}})),
	353: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(49)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(216)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	363: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(159)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__immediate_user_name_or_group),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__immediate_user_name_or_group),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(164)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym__immediate_user_name_or_group),
	})))),
	381: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(204)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(162)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_image_digest),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_image_digest),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(212)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(211)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(168)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_user_instruction),
		Fproduction_id: uint16(11),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(164)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_image_digest_repeat1),
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
		Fcount: uint8(1),
	}})),
	413: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_image_digest_repeat1),
	})))),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_image_digest_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(212)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_image_digest_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(59)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_path),
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
		Fcount: uint8(1),
	}})),
	423: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_path),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	429: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	433: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	435: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_add_instruction_repeat1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	437: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_add_instruction_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_add_instruction_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(322)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_add_instruction_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(84)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_add_instruction_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(84)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_add_instruction_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(201)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_add_instruction_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(294)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	454: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(67)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_path_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(203)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	460: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_param),
		Fproduction_id: uint16(13),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_param),
		Fproduction_id: uint16(13),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_path),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_path),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(303)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__imm_expansion),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__imm_expansion),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__expansion_body),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__expansion_body),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(206)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__stopsignal_value_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	490: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__stopsignal_value_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(77)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	493: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__stopsignal_value_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(208)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_image_alias),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(186)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_expansion),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_expansion),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
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
		Fcount: uint8(2),
	}})),
	508: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(82)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_path_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(200)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_image_spec),
		Fproduction_id: uint16(1),
	})))),
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
		Fcount: uint8(1),
	}})),
	516: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_image_spec),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(93)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(200)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_image_alias_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	528: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_image_alias_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(186)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_image_alias_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(85)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_path_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(86)),
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
		Fsymbol:      uint16(aux_sym_path_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(206)),
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
		Fsymbol:      uint16(sym_expose_instruction),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(194)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(99)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(257)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(323)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_image_alias),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(79)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_expose_instruction_repeat1),
	})))),
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
		Fcount: uint8(2),
	}})),
	556: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_expose_instruction_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(194)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_expose_instruction_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(99)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__stopsignal_value),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__stopsignal_value),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_env_instruction_repeat1),
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
		Fcount: uint8(2),
	}})),
	576: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_env_instruction_repeat1),
	})))),
	577: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(271)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(1),
	}})),
	579: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_add_instruction),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(147)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(190)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_expose_port),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_expose_port),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_copy_instruction),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(242)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_env_instruction),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(271)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(116)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(254)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(116)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(55)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(207)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(106)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(279)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	625: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_single_quoted_string_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(116)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	628: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_single_quoted_string_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_single_quoted_string_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(116)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	633: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_add_instruction),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(123)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(151)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(123)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(160)),
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
		Fcount: uint8(1),
	}})),
	643: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_copy_instruction),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(223)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(126)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(226)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_label_pair),
		Fproduction_id: uint16(9),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_label_pair),
		Fproduction_id: uint16(9),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_run_instruction_repeat2),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	659: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_run_instruction_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(190)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shell_command_repeat1),
	})))),
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
		Fcount: uint8(2),
	}})),
	664: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_shell_command_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(58)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_run_instruction),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_run_instruction),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(209)),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_json_string_repeat1),
	})))),
	676: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	677: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_json_string_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(135)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_json_string_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(135)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(229)),
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
		Fcount: uint8(1),
	}})),
	685: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(157)),
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
		Fsymbol:      uint16(sym_run_instruction),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_add_instruction),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_image_spec),
		Fproduction_id: uint16(5),
	})))),
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
		Fcount: uint8(1),
	}})),
	695: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_image_spec),
		Fproduction_id: uint16(5),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(205)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_copy_instruction),
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
		Fcount: uint8(1),
	}})),
	707: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_add_instruction_repeat2),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_add_instruction_repeat2),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_double_quoted_string),
	})))),
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
		Fcount: uint8(1),
	}})),
	713: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_double_quoted_string),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_shell_command),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_single_quoted_string),
	})))),
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
		Fcount: uint8(1),
	}})),
	721: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_single_quoted_string),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_shell_command),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	725: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_add_instruction_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(323)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_add_instruction),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(240)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(135)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(135)),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_copy_instruction),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_double_quoted_string),
	})))),
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
		Fcount: uint8(1),
	}})),
	740: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_double_quoted_string),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_single_quoted_string),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_single_quoted_string),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(133)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_expose_port),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_expose_port),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__immediate_user_name_or_group),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_json_string_array_repeat1),
	})))),
	757: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_json_string_array_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	761: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_heredoc_block_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(291)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_heredoc_block_repeat1),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_volume_instruction),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(315)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(109)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(251)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_shell_command_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(63)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(230)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(218)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(291)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(228)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(317)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(193)),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_shell_command_repeat1),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(231)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(233)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(236)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(140)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(306)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(170)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(269)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(289)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_volume_instruction),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(288)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(308)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	825: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_volume_instruction_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	827: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_volume_instruction_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(101)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(134)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(312)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(138)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(137)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(248)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(182)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(181)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(314)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(185)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(184)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(318)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(316)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(319)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(320)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(321)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_json_string_array),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_image_spec),
		Fproduction_id: uint16(8),
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
		Fcount: uint8(1),
	}})),
	878: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_image_spec),
		Fproduction_id: uint16(8),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(266)),
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
		Fcount: uint8(1),
	}})),
	884: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_heredoc_block),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_json_string),
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
		Fsymbol:      uint16(sym_json_string_array),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_heredoc_block),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_json_string_array),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_from_instruction),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(144)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_arg_instruction),
		Fproduction_id: uint16(4),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_env_pair),
		Fproduction_id: uint16(10),
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
		Fcount: uint8(1),
	}})),
	906: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_env_pair),
		Fproduction_id: uint16(10),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_json_string),
	})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_volume_instruction_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_image_spec),
		Fproduction_id: uint16(6),
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
		Fcount: uint8(1),
	}})),
	918: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_image_spec),
		Fproduction_id: uint16(6),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_from_instruction),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(155)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_user_instruction),
		Fproduction_id: uint16(3),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(92)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(287)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_workdir_instruction),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(283)),
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
	936: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_stopsignal_instruction),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_healthcheck_instruction),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__spaced_env_pair),
		Fproduction_id: uint16(10),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(297)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(62)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_from_instruction),
		Fproduction_id: uint16(14),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(238)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_entrypoint_instruction),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(285)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_onbuild_instruction),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(227)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(175)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(68)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(301)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_shell_instruction),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(324)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_volume_instruction),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(178)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(111)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_env_instruction),
		Fproduction_id: uint16(2),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(136)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(232)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_healthcheck_instruction),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(139)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(41)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(183)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(217)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(224)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_cmd_instruction),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(114)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_from_instruction),
		Fproduction_id: uint16(7),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(115)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_maintainer_instruction),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(192)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(118)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_cross_build_instruction),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(149)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(274)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(120)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(267)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(268)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(276)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(280)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(281)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(284)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_arg_instruction),
		Fproduction_id: uint16(12),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(290)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(293)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(296)),
	}})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(298)),
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
		Fcount: uint8(1),
	}})),
	1048: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(302)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(304)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(305)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(309)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(311)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(221)),
	}})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token_heredoc_marker = 0
const ts_external_token_heredoc_line = 1
const ts_external_token_heredoc_end = 2
const ts_external_token_heredoc_nl = 3
const ts_external_token_error_sentinel = 4

var ts_external_scanner_symbol_map = [5]TSSymbol{
	0: uint16(sym_heredoc_marker),
	1: uint16(sym_heredoc_line),
	2: uint16(sym_heredoc_end),
	3: uint16(sym_heredoc_nl),
	4: uint16(sym_error_sentinel),
}

var ts_external_scanner_states = [6][5]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
		3: libc.BoolUint8(true1 != 0),
		4: libc.BoolUint8(true1 != 0),
	},
	2: {
		0: libc.BoolUint8(true1 != 0),
	},
	3: {
		0: libc.BoolUint8(true1 != 0),
		3: libc.BoolUint8(true1 != 0),
	},
	4: {
		3: libc.BoolUint8(true1 != 0),
	},
	5: {
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_dockerfile(tls *libc.TLS) (r uintptr) {
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
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_dockerfile_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_dockerfile_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_dockerfile_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_dockerfile_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_dockerfile_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00\n\x00FROM\x00AS\x00RUN\x00CMD\x00LABEL\x00EXPOSE\x00ENV\x00ADD\x00COPY\x00ENTRYPOINT\x00VOLUME\x00USER\x00:\x00_user_name_or_group_token1\x00_immediate_user_name_or_group_fragment_token1\x00WORKDIR\x00ARG\x00unquoted_string\x00=\x00ONBUILD\x00STOPSIGNAL\x00_stopsignal_value_token1\x00_stopsignal_value_token2\x00HEALTHCHECK\x00NONE\x00SHELL\x00MAINTAINER\x00maintainer_instruction_token2\x00CROSS_BUILD\x00path_token1\x00path_token2\x00path_token3\x00path_with_heredoc_token1\x00$\x00{\x00variable\x00}\x00_spaced_env_pair_token1\x00expose_port_token1\x00/tcp\x00/udp\x00image_name_token1\x00image_name_token2\x00image_tag_token1\x00@\x00image_digest_token1\x00--\x00param_token1\x00param_token2\x00mount\x00,\x00mount_param_param_token1\x00image_alias_token1\x00image_alias_token2\x00shell_fragment_token1\x00shell_fragment_token2\x00shell_fragment_token3\x00shell_fragment_token4\x00line_continuation\x00[\x00]\x00\"\x00json_string_token1\x00escape_sequence\x00double_quoted_string_token1\x00\\\x00'\x00single_quoted_string_token1\x00unquoted_string_token1\x00\\ \x00_non_newline_whitespace\x00comment\x00heredoc_marker\x00heredoc_line\x00heredoc_end\x00_heredoc_nl\x00error_sentinel\x00source_file\x00_instruction\x00from_instruction\x00run_instruction\x00cmd_instruction\x00label_instruction\x00expose_instruction\x00env_instruction\x00add_instruction\x00copy_instruction\x00entrypoint_instruction\x00volume_instruction\x00user_instruction\x00_immediate_user_name_or_group\x00_immediate_user_name_or_group_fragment\x00workdir_instruction\x00arg_instruction\x00onbuild_instruction\x00stopsignal_instruction\x00_stopsignal_value\x00healthcheck_instruction\x00shell_instruction\x00maintainer_instruction\x00cross_build_instruction\x00heredoc_block\x00path\x00expansion\x00_immediate_expansion\x00_expansion_body\x00env_pair\x00_env_key\x00expose_port\x00label_pair\x00image_spec\x00image_name\x00image_tag\x00image_digest\x00param\x00mount_param\x00mount_param_param\x00image_alias\x00shell_command\x00shell_fragment\x00json_string_array\x00json_string\x00double_quoted_string\x00single_quoted_string\x00source_file_repeat1\x00run_instruction_repeat1\x00run_instruction_repeat2\x00label_instruction_repeat1\x00expose_instruction_repeat1\x00env_instruction_repeat1\x00add_instruction_repeat1\x00add_instruction_repeat2\x00volume_instruction_repeat1\x00_user_name_or_group_repeat1\x00_stopsignal_value_repeat1\x00heredoc_block_repeat1\x00path_repeat1\x00image_name_repeat1\x00image_tag_repeat1\x00image_digest_repeat1\x00mount_param_repeat1\x00image_alias_repeat1\x00shell_command_repeat1\x00shell_fragment_repeat1\x00json_string_array_repeat1\x00json_string_repeat1\x00double_quoted_string_repeat1\x00single_quoted_string_repeat1\x00unquoted_string_repeat1\x00as\x00default\x00digest\x00group\x00key\x00name\x00tag\x00user\x00value\x00"
