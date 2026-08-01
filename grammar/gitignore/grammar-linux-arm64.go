// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-gitignore/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-gitignore -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-gitignore/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_gitignore

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
const LANGUAGE_VERSION = 13
const LARGE_STATE_COUNT = 2
const MAX_ALIAS_SEQUENCE_LENGTH = 5
const PRODUCTION_ID_COUNT = 19
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 62
const SYMBOL_COUNT = 51
const TOKEN_COUNT = 32
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

const sym_comment = 1
const anon_sym_BANG = 2
const sym_directory_separator = 3
const sym_directory_separator_escaped = 4
const sym_pattern_char = 5
const anon_sym_BSLASH = 6
const aux_sym_pattern_char_escaped_token1 = 7
const anon_sym_QMARK = 8
const anon_sym_STAR = 9
const anon_sym_STAR_STAR = 10
const anon_sym_LBRACK = 11
const anon_sym_CARET = 12
const anon_sym_RBRACK = 13
const anon_sym_DASH = 14
const sym_bracket_char = 15
const anon_sym_LBRACK_COLON = 16
const anon_sym_alnum = 17
const anon_sym_COLON_RBRACK = 18
const anon_sym_alpha = 19
const anon_sym_blank = 20
const anon_sym_cntrl = 21
const anon_sym_digit = 22
const anon_sym_graph = 23
const anon_sym_lower = 24
const anon_sym_print = 25
const anon_sym_punct = 26
const anon_sym_space = 27
const anon_sym_upper = 28
const anon_sym_xdigit = 29
const sym__trailing_spaces = 30
const sym__newline = 31
const sym_document = 32
const sym__line = 33
const sym_pattern = 34
const sym__directory_separator = 35
const aux_sym__pattern = 36
const sym_pattern_char_escaped = 37
const sym__wildcard = 38
const sym_bracket_expr = 39
const sym__bracket_pattern = 40
const sym__bracket_pattern_closing_bracket = 41
const sym__bracket_char_closing_bracket = 42
const sym__bracket_range_closing_bracket = 43
const sym__bracket_char = 44
const sym_bracket_char_escaped = 45
const sym_bracket_range = 46
const sym_bracket_char_class = 47
const aux_sym_document_repeat1 = 48
const aux_sym_pattern_repeat1 = 49
const aux_sym_bracket_expr_repeat1 = 50

var ts_symbol_names = [51]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 12,
	3:  __ccgo_ts + 21,
	4:  __ccgo_ts + 41,
	5:  __ccgo_ts + 69,
	6:  __ccgo_ts + 82,
	7:  __ccgo_ts + 84,
	8:  __ccgo_ts + 112,
	9:  __ccgo_ts + 133,
	10: __ccgo_ts + 148,
	11: __ccgo_ts + 175,
	12: __ccgo_ts + 177,
	13: __ccgo_ts + 194,
	14: __ccgo_ts + 196,
	15: __ccgo_ts + 198,
	16: __ccgo_ts + 211,
	17: __ccgo_ts + 214,
	18: __ccgo_ts + 220,
	19: __ccgo_ts + 223,
	20: __ccgo_ts + 229,
	21: __ccgo_ts + 235,
	22: __ccgo_ts + 241,
	23: __ccgo_ts + 247,
	24: __ccgo_ts + 253,
	25: __ccgo_ts + 259,
	26: __ccgo_ts + 265,
	27: __ccgo_ts + 271,
	28: __ccgo_ts + 277,
	29: __ccgo_ts + 283,
	30: __ccgo_ts + 290,
	31: __ccgo_ts + 307,
	32: __ccgo_ts + 316,
	33: __ccgo_ts + 325,
	34: __ccgo_ts + 331,
	35: __ccgo_ts + 339,
	36: __ccgo_ts + 360,
	37: __ccgo_ts + 369,
	38: __ccgo_ts + 390,
	39: __ccgo_ts + 400,
	40: __ccgo_ts + 413,
	41: __ccgo_ts + 430,
	42: __ccgo_ts + 463,
	43: __ccgo_ts + 493,
	44: __ccgo_ts + 507,
	45: __ccgo_ts + 521,
	46: __ccgo_ts + 493,
	47: __ccgo_ts + 542,
	48: __ccgo_ts + 561,
	49: __ccgo_ts + 578,
	50: __ccgo_ts + 594,
}

var ts_symbol_map = [51]TSSymbol{
	1:  uint16(sym_comment),
	2:  uint16(anon_sym_BANG),
	3:  uint16(sym_directory_separator),
	4:  uint16(sym_directory_separator_escaped),
	5:  uint16(sym_pattern_char),
	6:  uint16(anon_sym_BSLASH),
	7:  uint16(aux_sym_pattern_char_escaped_token1),
	8:  uint16(anon_sym_QMARK),
	9:  uint16(anon_sym_STAR),
	10: uint16(anon_sym_STAR_STAR),
	11: uint16(anon_sym_LBRACK),
	12: uint16(anon_sym_CARET),
	13: uint16(anon_sym_RBRACK),
	14: uint16(anon_sym_DASH),
	15: uint16(sym_bracket_char),
	16: uint16(anon_sym_LBRACK_COLON),
	17: uint16(anon_sym_alnum),
	18: uint16(anon_sym_COLON_RBRACK),
	19: uint16(anon_sym_alpha),
	20: uint16(anon_sym_blank),
	21: uint16(anon_sym_cntrl),
	22: uint16(anon_sym_digit),
	23: uint16(anon_sym_graph),
	24: uint16(anon_sym_lower),
	25: uint16(anon_sym_print),
	26: uint16(anon_sym_punct),
	27: uint16(anon_sym_space),
	28: uint16(anon_sym_upper),
	29: uint16(anon_sym_xdigit),
	30: uint16(sym__trailing_spaces),
	31: uint16(sym__newline),
	32: uint16(sym_document),
	33: uint16(sym__line),
	34: uint16(sym_pattern),
	35: uint16(sym__directory_separator),
	36: uint16(aux_sym__pattern),
	37: uint16(sym_pattern_char_escaped),
	38: uint16(sym__wildcard),
	39: uint16(sym_bracket_expr),
	40: uint16(sym__bracket_pattern),
	41: uint16(sym__bracket_pattern_closing_bracket),
	42: uint16(sym__bracket_char_closing_bracket),
	43: uint16(sym_bracket_range),
	44: uint16(sym__bracket_char),
	45: uint16(sym_bracket_char_escaped),
	46: uint16(sym_bracket_range),
	47: uint16(sym_bracket_char_class),
	48: uint16(aux_sym_document_repeat1),
	49: uint16(aux_sym_pattern_repeat1),
	50: uint16(aux_sym_bracket_expr_repeat1),
}

var ts_symbol_metadata = [51]TSSymbolMetadata{
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
	},
	7: {},
	8: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	9: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	10: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	11: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	12: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	31: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	32: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	33: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	35: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	36: {},
	37: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	38: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	39: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	40: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	41: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	42: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	43: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	44: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
	48: {},
	49: {},
	50: {},
}

const field_directory_flag = 1
const field_name = 2
const field_relative_flag = 3

var ts_field_names = [4]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 615,
	2: __ccgo_ts + 630,
	3: __ccgo_ts + 635,
}

var ts_field_map_slices = [19]TSFieldMapSlice{
	2: {
		Flength: uint16(1),
	},
	3: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	4: {
		Findex:  uint16(2),
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(3),
		Flength: uint16(1),
	},
	6: {
		Findex:  uint16(4),
		Flength: uint16(1),
	},
	7: {
		Findex:  uint16(5),
		Flength: uint16(1),
	},
	8: {
		Findex:  uint16(6),
		Flength: uint16(2),
	},
	9: {
		Findex:  uint16(8),
		Flength: uint16(2),
	},
	10: {
		Findex:  uint16(10),
		Flength: uint16(2),
	},
	11: {
		Findex:  uint16(12),
		Flength: uint16(2),
	},
	12: {
		Findex:  uint16(14),
		Flength: uint16(2),
	},
	13: {
		Findex:  uint16(16),
		Flength: uint16(2),
	},
	14: {
		Findex:  uint16(18),
		Flength: uint16(2),
	},
	16: {
		Findex:  uint16(20),
		Flength: uint16(1),
	},
	17: {
		Findex:  uint16(21),
		Flength: uint16(3),
	},
	18: {
		Findex:  uint16(24),
		Flength: uint16(3),
	},
}

var ts_field_map_entries = [27]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_relative_flag),
	},
	1: {
		Ffield_id:    uint16(field_directory_flag),
		Fchild_index: uint8(1),
	},
	2: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	3: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(1),
	},
	4: {
		Ffield_id:    uint16(field_directory_flag),
		Fchild_index: uint8(2),
	},
	5: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	6: {
		Ffield_id:    uint16(field_directory_flag),
		Fchild_index: uint8(2),
	},
	7: {
		Ffield_id: uint16(field_relative_flag),
	},
	8: {
		Ffield_id: uint16(field_relative_flag),
	},
	9: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	10: {
		Ffield_id:    uint16(field_directory_flag),
		Fchild_index: uint8(2),
	},
	11: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	12: {
		Ffield_id:  uint16(field_relative_flag),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	13: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	14: {
		Ffield_id:    uint16(field_directory_flag),
		Fchild_index: uint8(3),
	},
	15: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(1),
	},
	16: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(1),
	},
	17: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	18: {
		Ffield_id:    uint16(field_directory_flag),
		Fchild_index: uint8(3),
	},
	19: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	20: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	21: {
		Ffield_id:    uint16(field_directory_flag),
		Fchild_index: uint8(3),
	},
	22: {
		Ffield_id: uint16(field_relative_flag),
	},
	23: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	24: {
		Ffield_id:    uint16(field_directory_flag),
		Fchild_index: uint8(4),
	},
	25: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(1),
	},
	26: {
		Ffield_id:    uint16(field_relative_flag),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
}

var ts_alias_sequences = [19][5]TSSymbol{
	0: {},
	1: {
		0: uint16(sym_bracket_char),
	},
	15: {
		1: uint16(anon_sym_CARET),
	},
}

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
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(63)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(62)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(69)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(78)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(77)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('\n') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32(' ') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(70)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('\n') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32(' ') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(69)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('/') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('\n') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32(' ') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('g') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('!') {
			state = uint16(63)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(69)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(78)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(77)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('/') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('*') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(70)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('*') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(69)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('/') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('-') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(69)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(78)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('/') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('/') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('[') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(69)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(78)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('/') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('\\') {
			state = uint16(69)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('/') && lookahead != int32(']') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32(']') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('a') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('a') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('a') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('a') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('c') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('c') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('d') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('e') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('e') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('e') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('g') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('g') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('h') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('h') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('i') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('i') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('i') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('i') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('i') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('k') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('l') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('l') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('l') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('m') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('n') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('n') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('n') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('n') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('n') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('o') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('p') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('p') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('p') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('p') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('r') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('r') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('r') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('r') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('r') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('t') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('t') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('t') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('t') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('t') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('u') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('w') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('/') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(60):
		if eof != 0 {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('\n') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32(' ') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(63)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(62)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(70)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_directory_separator)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_directory_separator_escaped)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_pattern_char)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_pattern_char)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_pattern_char)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(' ') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_pattern_char_escaped_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bracket_char)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_bracket_char)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(':') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_alnum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_alpha)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_blank)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_cntrl)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_digit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_graph)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_lower)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_print)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_punct)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_space)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_upper)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_xdigit)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__trailing_spaces)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(' ') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__newline)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [62]TSLexMode{
	0: {},
	1: {
		Flex_state: uint16(60),
	},
	2: {
		Flex_state: uint16(60),
	},
	3: {
		Flex_state: uint16(60),
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
		Flex_state: uint16(5),
	},
	9: {
		Flex_state: uint16(2),
	},
	10: {
		Flex_state: uint16(2),
	},
	11: {
		Flex_state: uint16(60),
	},
	12: {
		Flex_state: uint16(60),
	},
	13: {
		Flex_state: uint16(10),
	},
	14: {
		Flex_state: uint16(6),
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
		Flex_state: uint16(4),
	},
	21: {
		Flex_state: uint16(3),
	},
	22: {
		Flex_state: uint16(3),
	},
	23: {
		Flex_state: uint16(3),
	},
	24: {
		Flex_state: uint16(10),
	},
	25: {
		Flex_state: uint16(10),
	},
	26: {
		Flex_state: uint16(2),
	},
	27: {
		Flex_state: uint16(7),
	},
	28: {
		Flex_state: uint16(2),
	},
	29: {
		Flex_state: uint16(7),
	},
	30: {
		Flex_state: uint16(2),
	},
	31: {
		Flex_state: uint16(10),
	},
	32: {
		Flex_state: uint16(2),
	},
	33: {
		Flex_state: uint16(10),
	},
	34: {
		Flex_state: uint16(10),
	},
	35: {
		Flex_state: uint16(10),
	},
	36: {
		Flex_state: uint16(7),
	},
	37: {
		Flex_state: uint16(2),
	},
	38: {
		Flex_state: uint16(10),
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
		Flex_state: uint16(8),
	},
	45: {
		Flex_state: uint16(8),
	},
	46: {
		Flex_state: uint16(8),
	},
	47: {
		Flex_state: uint16(8),
	},
	48: {
		Flex_state: uint16(11),
	},
	49: {
		Flex_state: uint16(10),
	},
	50: {
		Flex_state: uint16(10),
	},
	51: {
		Flex_state: uint16(10),
	},
	52: {
		Flex_state: uint16(11),
	},
	53: {
		Flex_state: uint16(10),
	},
	54: {
		Flex_state: uint16(4),
	},
	55: {
		Flex_state: uint16(60),
	},
	56: {
		Flex_state: uint16(4),
	},
	57: {
		Flex_state: uint16(59),
	},
	58: {
		Flex_state: uint16(59),
	},
	59: {
		Flex_state: uint16(60),
	},
	60: {},
	61: {
		Flex_state: uint16(59),
	},
}

var ts_parse_table = [2][51]uint16_t{
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
		11: uint16(1),
		12: uint16(1),
		13: uint16(1),
		14: uint16(1),
		15: uint16(1),
	},
	1: {
		0:  uint16(3),
		1:  uint16(5),
		2:  uint16(7),
		3:  uint16(9),
		4:  uint16(9),
		5:  uint16(11),
		6:  uint16(13),
		8:  uint16(15),
		9:  uint16(11),
		10: uint16(15),
		11: uint16(17),
		30: uint16(19),
		31: uint16(21),
		32: uint16(60),
		33: uint16(3),
		34: uint16(54),
		35: uint16(27),
		36: uint16(7),
		37: uint16(7),
		38: uint16(7),
		39: uint16(7),
		48: uint16(3),
	},
}

var ts_small_parse_table = [1176]uint16_t{
	0:    uint16(14),
	1:    uint16(23),
	2:    uint16(1),
	4:    uint16(25),
	5:    uint16(1),
	6:    uint16(sym_comment),
	7:    uint16(28),
	8:    uint16(1),
	9:    uint16(anon_sym_BANG),
	10:   uint16(37),
	11:   uint16(1),
	12:   uint16(anon_sym_BSLASH),
	13:   uint16(43),
	14:   uint16(1),
	15:   uint16(anon_sym_LBRACK),
	16:   uint16(46),
	17:   uint16(1),
	18:   uint16(sym__trailing_spaces),
	19:   uint16(49),
	20:   uint16(1),
	21:   uint16(sym__newline),
	22:   uint16(27),
	23:   uint16(1),
	24:   uint16(sym__directory_separator),
	25:   uint16(54),
	26:   uint16(1),
	27:   uint16(sym_pattern),
	28:   uint16(31),
	29:   uint16(2),
	30:   uint16(sym_directory_separator),
	31:   uint16(sym_directory_separator_escaped),
	32:   uint16(34),
	33:   uint16(2),
	34:   uint16(sym_pattern_char),
	35:   uint16(anon_sym_STAR),
	36:   uint16(40),
	37:   uint16(2),
	38:   uint16(anon_sym_QMARK),
	39:   uint16(anon_sym_STAR_STAR),
	40:   uint16(2),
	41:   uint16(2),
	42:   uint16(sym__line),
	43:   uint16(aux_sym_document_repeat1),
	44:   uint16(7),
	45:   uint16(4),
	46:   uint16(aux_sym__pattern),
	47:   uint16(sym_pattern_char_escaped),
	48:   uint16(sym__wildcard),
	49:   uint16(sym_bracket_expr),
	50:   uint16(14),
	51:   uint16(5),
	52:   uint16(1),
	53:   uint16(sym_comment),
	54:   uint16(7),
	55:   uint16(1),
	56:   uint16(anon_sym_BANG),
	57:   uint16(13),
	58:   uint16(1),
	59:   uint16(anon_sym_BSLASH),
	60:   uint16(17),
	61:   uint16(1),
	62:   uint16(anon_sym_LBRACK),
	63:   uint16(19),
	64:   uint16(1),
	65:   uint16(sym__trailing_spaces),
	66:   uint16(52),
	67:   uint16(1),
	69:   uint16(54),
	70:   uint16(1),
	71:   uint16(sym__newline),
	72:   uint16(27),
	73:   uint16(1),
	74:   uint16(sym__directory_separator),
	75:   uint16(54),
	76:   uint16(1),
	77:   uint16(sym_pattern),
	78:   uint16(9),
	79:   uint16(2),
	80:   uint16(sym_directory_separator),
	81:   uint16(sym_directory_separator_escaped),
	82:   uint16(11),
	83:   uint16(2),
	84:   uint16(sym_pattern_char),
	85:   uint16(anon_sym_STAR),
	86:   uint16(15),
	87:   uint16(2),
	88:   uint16(anon_sym_QMARK),
	89:   uint16(anon_sym_STAR_STAR),
	90:   uint16(2),
	91:   uint16(2),
	92:   uint16(sym__line),
	93:   uint16(aux_sym_document_repeat1),
	94:   uint16(7),
	95:   uint16(4),
	96:   uint16(aux_sym__pattern),
	97:   uint16(sym_pattern_char_escaped),
	98:   uint16(sym__wildcard),
	99:   uint16(sym_bracket_expr),
	100:  uint16(10),
	101:  uint16(13),
	102:  uint16(1),
	103:  uint16(anon_sym_BSLASH),
	104:  uint16(17),
	105:  uint16(1),
	106:  uint16(anon_sym_LBRACK),
	107:  uint16(62),
	108:  uint16(1),
	109:  uint16(sym__trailing_spaces),
	110:  uint16(64),
	111:  uint16(1),
	112:  uint16(sym__newline),
	113:  uint16(15),
	114:  uint16(1),
	115:  uint16(sym__directory_separator),
	116:  uint16(42),
	117:  uint16(1),
	118:  uint16(aux_sym_pattern_repeat1),
	119:  uint16(56),
	120:  uint16(2),
	121:  uint16(sym_directory_separator),
	122:  uint16(sym_directory_separator_escaped),
	123:  uint16(58),
	124:  uint16(2),
	125:  uint16(sym_pattern_char),
	126:  uint16(anon_sym_STAR),
	127:  uint16(60),
	128:  uint16(2),
	129:  uint16(anon_sym_QMARK),
	130:  uint16(anon_sym_STAR_STAR),
	131:  uint16(10),
	132:  uint16(4),
	133:  uint16(aux_sym__pattern),
	134:  uint16(sym_pattern_char_escaped),
	135:  uint16(sym__wildcard),
	136:  uint16(sym_bracket_expr),
	137:  uint16(10),
	138:  uint16(13),
	139:  uint16(1),
	140:  uint16(anon_sym_BSLASH),
	141:  uint16(17),
	142:  uint16(1),
	143:  uint16(anon_sym_LBRACK),
	144:  uint16(68),
	145:  uint16(1),
	146:  uint16(sym__trailing_spaces),
	147:  uint16(70),
	148:  uint16(1),
	149:  uint16(sym__newline),
	150:  uint16(19),
	151:  uint16(1),
	152:  uint16(sym__directory_separator),
	153:  uint16(43),
	154:  uint16(1),
	155:  uint16(aux_sym_pattern_repeat1),
	156:  uint16(58),
	157:  uint16(2),
	158:  uint16(sym_pattern_char),
	159:  uint16(anon_sym_STAR),
	160:  uint16(60),
	161:  uint16(2),
	162:  uint16(anon_sym_QMARK),
	163:  uint16(anon_sym_STAR_STAR),
	164:  uint16(66),
	165:  uint16(2),
	166:  uint16(sym_directory_separator),
	167:  uint16(sym_directory_separator_escaped),
	168:  uint16(10),
	169:  uint16(4),
	170:  uint16(aux_sym__pattern),
	171:  uint16(sym_pattern_char_escaped),
	172:  uint16(sym__wildcard),
	173:  uint16(sym_bracket_expr),
	174:  uint16(10),
	175:  uint16(13),
	176:  uint16(1),
	177:  uint16(anon_sym_BSLASH),
	178:  uint16(17),
	179:  uint16(1),
	180:  uint16(anon_sym_LBRACK),
	181:  uint16(74),
	182:  uint16(1),
	183:  uint16(sym__trailing_spaces),
	184:  uint16(76),
	185:  uint16(1),
	186:  uint16(sym__newline),
	187:  uint16(22),
	188:  uint16(1),
	189:  uint16(sym__directory_separator),
	190:  uint16(39),
	191:  uint16(1),
	192:  uint16(aux_sym_pattern_repeat1),
	193:  uint16(58),
	194:  uint16(2),
	195:  uint16(sym_pattern_char),
	196:  uint16(anon_sym_STAR),
	197:  uint16(60),
	198:  uint16(2),
	199:  uint16(anon_sym_QMARK),
	200:  uint16(anon_sym_STAR_STAR),
	201:  uint16(72),
	202:  uint16(2),
	203:  uint16(sym_directory_separator),
	204:  uint16(sym_directory_separator_escaped),
	205:  uint16(10),
	206:  uint16(4),
	207:  uint16(aux_sym__pattern),
	208:  uint16(sym_pattern_char_escaped),
	209:  uint16(sym__wildcard),
	210:  uint16(sym_bracket_expr),
	211:  uint16(10),
	212:  uint16(13),
	213:  uint16(1),
	214:  uint16(anon_sym_BSLASH),
	215:  uint16(17),
	216:  uint16(1),
	217:  uint16(anon_sym_LBRACK),
	218:  uint16(80),
	219:  uint16(1),
	220:  uint16(sym__trailing_spaces),
	221:  uint16(82),
	222:  uint16(1),
	223:  uint16(sym__newline),
	224:  uint16(23),
	225:  uint16(1),
	226:  uint16(sym__directory_separator),
	227:  uint16(41),
	228:  uint16(1),
	229:  uint16(aux_sym_pattern_repeat1),
	230:  uint16(58),
	231:  uint16(2),
	232:  uint16(sym_pattern_char),
	233:  uint16(anon_sym_STAR),
	234:  uint16(60),
	235:  uint16(2),
	236:  uint16(anon_sym_QMARK),
	237:  uint16(anon_sym_STAR_STAR),
	238:  uint16(78),
	239:  uint16(2),
	240:  uint16(sym_directory_separator),
	241:  uint16(sym_directory_separator_escaped),
	242:  uint16(10),
	243:  uint16(4),
	244:  uint16(aux_sym__pattern),
	245:  uint16(sym_pattern_char_escaped),
	246:  uint16(sym__wildcard),
	247:  uint16(sym_bracket_expr),
	248:  uint16(9),
	249:  uint16(86),
	250:  uint16(1),
	251:  uint16(anon_sym_BSLASH),
	252:  uint16(88),
	253:  uint16(1),
	254:  uint16(anon_sym_RBRACK),
	255:  uint16(90),
	256:  uint16(1),
	257:  uint16(sym_bracket_char),
	258:  uint16(92),
	259:  uint16(1),
	260:  uint16(anon_sym_LBRACK_COLON),
	261:  uint16(46),
	262:  uint16(1),
	263:  uint16(sym__bracket_char_closing_bracket),
	264:  uint16(84),
	265:  uint16(2),
	266:  uint16(anon_sym_BANG),
	267:  uint16(anon_sym_CARET),
	268:  uint16(35),
	269:  uint16(2),
	270:  uint16(sym__bracket_pattern_closing_bracket),
	271:  uint16(sym__bracket_range_closing_bracket),
	272:  uint16(45),
	273:  uint16(2),
	274:  uint16(sym__bracket_char),
	275:  uint16(sym_bracket_char_escaped),
	276:  uint16(24),
	277:  uint16(4),
	278:  uint16(sym__bracket_pattern),
	279:  uint16(sym_bracket_range),
	280:  uint16(sym_bracket_char_class),
	281:  uint16(aux_sym_bracket_expr_repeat1),
	282:  uint16(7),
	283:  uint16(13),
	284:  uint16(1),
	285:  uint16(anon_sym_BSLASH),
	286:  uint16(17),
	287:  uint16(1),
	288:  uint16(anon_sym_LBRACK),
	289:  uint16(96),
	290:  uint16(1),
	291:  uint16(sym__trailing_spaces),
	292:  uint16(58),
	293:  uint16(2),
	294:  uint16(sym_pattern_char),
	295:  uint16(anon_sym_STAR),
	296:  uint16(60),
	297:  uint16(2),
	298:  uint16(anon_sym_QMARK),
	299:  uint16(anon_sym_STAR_STAR),
	300:  uint16(94),
	301:  uint16(3),
	302:  uint16(sym_directory_separator),
	303:  uint16(sym_directory_separator_escaped),
	304:  uint16(sym__newline),
	305:  uint16(10),
	306:  uint16(4),
	307:  uint16(aux_sym__pattern),
	308:  uint16(sym_pattern_char_escaped),
	309:  uint16(sym__wildcard),
	310:  uint16(sym_bracket_expr),
	311:  uint16(7),
	312:  uint16(103),
	313:  uint16(1),
	314:  uint16(anon_sym_BSLASH),
	315:  uint16(109),
	316:  uint16(1),
	317:  uint16(anon_sym_LBRACK),
	318:  uint16(112),
	319:  uint16(1),
	320:  uint16(sym__trailing_spaces),
	321:  uint16(100),
	322:  uint16(2),
	323:  uint16(sym_pattern_char),
	324:  uint16(anon_sym_STAR),
	325:  uint16(106),
	326:  uint16(2),
	327:  uint16(anon_sym_QMARK),
	328:  uint16(anon_sym_STAR_STAR),
	329:  uint16(98),
	330:  uint16(3),
	331:  uint16(sym_directory_separator),
	332:  uint16(sym_directory_separator_escaped),
	333:  uint16(sym__newline),
	334:  uint16(10),
	335:  uint16(4),
	336:  uint16(aux_sym__pattern),
	337:  uint16(sym_pattern_char_escaped),
	338:  uint16(sym__wildcard),
	339:  uint16(sym_bracket_expr),
	340:  uint16(2),
	341:  uint16(116),
	342:  uint16(4),
	343:  uint16(sym_pattern_char),
	344:  uint16(anon_sym_BSLASH),
	345:  uint16(anon_sym_STAR),
	346:  uint16(sym__trailing_spaces),
	347:  uint16(114),
	348:  uint16(9),
	350:  uint16(sym_comment),
	351:  uint16(anon_sym_BANG),
	352:  uint16(sym_directory_separator),
	353:  uint16(sym_directory_separator_escaped),
	354:  uint16(anon_sym_QMARK),
	355:  uint16(anon_sym_STAR_STAR),
	356:  uint16(anon_sym_LBRACK),
	357:  uint16(sym__newline),
	358:  uint16(2),
	359:  uint16(120),
	360:  uint16(4),
	361:  uint16(sym_pattern_char),
	362:  uint16(anon_sym_BSLASH),
	363:  uint16(anon_sym_STAR),
	364:  uint16(sym__trailing_spaces),
	365:  uint16(118),
	366:  uint16(9),
	368:  uint16(sym_comment),
	369:  uint16(anon_sym_BANG),
	370:  uint16(sym_directory_separator),
	371:  uint16(sym_directory_separator_escaped),
	372:  uint16(anon_sym_QMARK),
	373:  uint16(anon_sym_STAR_STAR),
	374:  uint16(anon_sym_LBRACK),
	375:  uint16(sym__newline),
	376:  uint16(8),
	377:  uint16(86),
	378:  uint16(1),
	379:  uint16(anon_sym_BSLASH),
	380:  uint16(88),
	381:  uint16(1),
	382:  uint16(anon_sym_RBRACK),
	383:  uint16(90),
	384:  uint16(1),
	385:  uint16(sym_bracket_char),
	386:  uint16(92),
	387:  uint16(1),
	388:  uint16(anon_sym_LBRACK_COLON),
	389:  uint16(46),
	390:  uint16(1),
	391:  uint16(sym__bracket_char_closing_bracket),
	392:  uint16(33),
	393:  uint16(2),
	394:  uint16(sym__bracket_pattern_closing_bracket),
	395:  uint16(sym__bracket_range_closing_bracket),
	396:  uint16(45),
	397:  uint16(2),
	398:  uint16(sym__bracket_char),
	399:  uint16(sym_bracket_char_escaped),
	400:  uint16(34),
	401:  uint16(4),
	402:  uint16(sym__bracket_pattern),
	403:  uint16(sym_bracket_range),
	404:  uint16(sym_bracket_char_class),
	405:  uint16(aux_sym_bracket_expr_repeat1),
	406:  uint16(7),
	407:  uint16(13),
	408:  uint16(1),
	409:  uint16(anon_sym_BSLASH),
	410:  uint16(17),
	411:  uint16(1),
	412:  uint16(anon_sym_LBRACK),
	413:  uint16(36),
	414:  uint16(1),
	415:  uint16(sym__directory_separator),
	416:  uint16(122),
	417:  uint16(2),
	418:  uint16(sym_directory_separator),
	419:  uint16(sym_directory_separator_escaped),
	420:  uint16(124),
	421:  uint16(2),
	422:  uint16(sym_pattern_char),
	423:  uint16(anon_sym_STAR),
	424:  uint16(126),
	425:  uint16(2),
	426:  uint16(anon_sym_QMARK),
	427:  uint16(anon_sym_STAR_STAR),
	428:  uint16(4),
	429:  uint16(4),
	430:  uint16(aux_sym__pattern),
	431:  uint16(sym_pattern_char_escaped),
	432:  uint16(sym__wildcard),
	433:  uint16(sym_bracket_expr),
	434:  uint16(7),
	435:  uint16(17),
	436:  uint16(1),
	437:  uint16(anon_sym_LBRACK),
	438:  uint16(130),
	439:  uint16(1),
	440:  uint16(anon_sym_BSLASH),
	441:  uint16(134),
	442:  uint16(1),
	443:  uint16(sym__trailing_spaces),
	444:  uint16(136),
	445:  uint16(1),
	446:  uint16(sym__newline),
	447:  uint16(128),
	448:  uint16(2),
	449:  uint16(sym_pattern_char),
	450:  uint16(anon_sym_STAR),
	451:  uint16(132),
	452:  uint16(2),
	453:  uint16(anon_sym_QMARK),
	454:  uint16(anon_sym_STAR_STAR),
	455:  uint16(9),
	456:  uint16(4),
	457:  uint16(aux_sym__pattern),
	458:  uint16(sym_pattern_char_escaped),
	459:  uint16(sym__wildcard),
	460:  uint16(sym_bracket_expr),
	461:  uint16(7),
	462:  uint16(17),
	463:  uint16(1),
	464:  uint16(anon_sym_LBRACK),
	465:  uint16(130),
	466:  uint16(1),
	467:  uint16(anon_sym_BSLASH),
	468:  uint16(138),
	469:  uint16(1),
	470:  uint16(sym__trailing_spaces),
	471:  uint16(140),
	472:  uint16(1),
	473:  uint16(sym__newline),
	474:  uint16(128),
	475:  uint16(2),
	476:  uint16(sym_pattern_char),
	477:  uint16(anon_sym_STAR),
	478:  uint16(132),
	479:  uint16(2),
	480:  uint16(anon_sym_QMARK),
	481:  uint16(anon_sym_STAR_STAR),
	482:  uint16(9),
	483:  uint16(4),
	484:  uint16(aux_sym__pattern),
	485:  uint16(sym_pattern_char_escaped),
	486:  uint16(sym__wildcard),
	487:  uint16(sym_bracket_expr),
	488:  uint16(7),
	489:  uint16(17),
	490:  uint16(1),
	491:  uint16(anon_sym_LBRACK),
	492:  uint16(130),
	493:  uint16(1),
	494:  uint16(anon_sym_BSLASH),
	495:  uint16(142),
	496:  uint16(1),
	497:  uint16(sym__trailing_spaces),
	498:  uint16(144),
	499:  uint16(1),
	500:  uint16(sym__newline),
	501:  uint16(128),
	502:  uint16(2),
	503:  uint16(sym_pattern_char),
	504:  uint16(anon_sym_STAR),
	505:  uint16(132),
	506:  uint16(2),
	507:  uint16(anon_sym_QMARK),
	508:  uint16(anon_sym_STAR_STAR),
	509:  uint16(9),
	510:  uint16(4),
	511:  uint16(aux_sym__pattern),
	512:  uint16(sym_pattern_char_escaped),
	513:  uint16(sym__wildcard),
	514:  uint16(sym_bracket_expr),
	515:  uint16(7),
	516:  uint16(17),
	517:  uint16(1),
	518:  uint16(anon_sym_LBRACK),
	519:  uint16(130),
	520:  uint16(1),
	521:  uint16(anon_sym_BSLASH),
	522:  uint16(146),
	523:  uint16(1),
	524:  uint16(sym__trailing_spaces),
	525:  uint16(148),
	526:  uint16(1),
	527:  uint16(sym__newline),
	528:  uint16(128),
	529:  uint16(2),
	530:  uint16(sym_pattern_char),
	531:  uint16(anon_sym_STAR),
	532:  uint16(132),
	533:  uint16(2),
	534:  uint16(anon_sym_QMARK),
	535:  uint16(anon_sym_STAR_STAR),
	536:  uint16(9),
	537:  uint16(4),
	538:  uint16(aux_sym__pattern),
	539:  uint16(sym_pattern_char_escaped),
	540:  uint16(sym__wildcard),
	541:  uint16(sym_bracket_expr),
	542:  uint16(7),
	543:  uint16(17),
	544:  uint16(1),
	545:  uint16(anon_sym_LBRACK),
	546:  uint16(130),
	547:  uint16(1),
	548:  uint16(anon_sym_BSLASH),
	549:  uint16(150),
	550:  uint16(1),
	551:  uint16(sym__trailing_spaces),
	552:  uint16(152),
	553:  uint16(1),
	554:  uint16(sym__newline),
	555:  uint16(128),
	556:  uint16(2),
	557:  uint16(sym_pattern_char),
	558:  uint16(anon_sym_STAR),
	559:  uint16(132),
	560:  uint16(2),
	561:  uint16(anon_sym_QMARK),
	562:  uint16(anon_sym_STAR_STAR),
	563:  uint16(9),
	564:  uint16(4),
	565:  uint16(aux_sym__pattern),
	566:  uint16(sym_pattern_char_escaped),
	567:  uint16(sym__wildcard),
	568:  uint16(sym_bracket_expr),
	569:  uint16(1),
	570:  uint16(154),
	571:  uint16(12),
	572:  uint16(anon_sym_alnum),
	573:  uint16(anon_sym_alpha),
	574:  uint16(anon_sym_blank),
	575:  uint16(anon_sym_cntrl),
	576:  uint16(anon_sym_digit),
	577:  uint16(anon_sym_graph),
	578:  uint16(anon_sym_lower),
	579:  uint16(anon_sym_print),
	580:  uint16(anon_sym_punct),
	581:  uint16(anon_sym_space),
	582:  uint16(anon_sym_upper),
	583:  uint16(anon_sym_xdigit),
	584:  uint16(7),
	585:  uint16(17),
	586:  uint16(1),
	587:  uint16(anon_sym_LBRACK),
	588:  uint16(130),
	589:  uint16(1),
	590:  uint16(anon_sym_BSLASH),
	591:  uint16(156),
	592:  uint16(1),
	593:  uint16(sym__trailing_spaces),
	594:  uint16(158),
	595:  uint16(1),
	596:  uint16(sym__newline),
	597:  uint16(128),
	598:  uint16(2),
	599:  uint16(sym_pattern_char),
	600:  uint16(anon_sym_STAR),
	601:  uint16(132),
	602:  uint16(2),
	603:  uint16(anon_sym_QMARK),
	604:  uint16(anon_sym_STAR_STAR),
	605:  uint16(9),
	606:  uint16(4),
	607:  uint16(aux_sym__pattern),
	608:  uint16(sym_pattern_char_escaped),
	609:  uint16(sym__wildcard),
	610:  uint16(sym_bracket_expr),
	611:  uint16(7),
	612:  uint16(17),
	613:  uint16(1),
	614:  uint16(anon_sym_LBRACK),
	615:  uint16(130),
	616:  uint16(1),
	617:  uint16(anon_sym_BSLASH),
	618:  uint16(160),
	619:  uint16(1),
	620:  uint16(sym__trailing_spaces),
	621:  uint16(162),
	622:  uint16(1),
	623:  uint16(sym__newline),
	624:  uint16(128),
	625:  uint16(2),
	626:  uint16(sym_pattern_char),
	627:  uint16(anon_sym_STAR),
	628:  uint16(132),
	629:  uint16(2),
	630:  uint16(anon_sym_QMARK),
	631:  uint16(anon_sym_STAR_STAR),
	632:  uint16(9),
	633:  uint16(4),
	634:  uint16(aux_sym__pattern),
	635:  uint16(sym_pattern_char_escaped),
	636:  uint16(sym__wildcard),
	637:  uint16(sym_bracket_expr),
	638:  uint16(7),
	639:  uint16(17),
	640:  uint16(1),
	641:  uint16(anon_sym_LBRACK),
	642:  uint16(130),
	643:  uint16(1),
	644:  uint16(anon_sym_BSLASH),
	645:  uint16(164),
	646:  uint16(1),
	647:  uint16(sym__trailing_spaces),
	648:  uint16(166),
	649:  uint16(1),
	650:  uint16(sym__newline),
	651:  uint16(128),
	652:  uint16(2),
	653:  uint16(sym_pattern_char),
	654:  uint16(anon_sym_STAR),
	655:  uint16(132),
	656:  uint16(2),
	657:  uint16(anon_sym_QMARK),
	658:  uint16(anon_sym_STAR_STAR),
	659:  uint16(9),
	660:  uint16(4),
	661:  uint16(aux_sym__pattern),
	662:  uint16(sym_pattern_char_escaped),
	663:  uint16(sym__wildcard),
	664:  uint16(sym_bracket_expr),
	665:  uint16(6),
	666:  uint16(86),
	667:  uint16(1),
	668:  uint16(anon_sym_BSLASH),
	669:  uint16(90),
	670:  uint16(1),
	671:  uint16(sym_bracket_char),
	672:  uint16(92),
	673:  uint16(1),
	674:  uint16(anon_sym_LBRACK_COLON),
	675:  uint16(168),
	676:  uint16(1),
	677:  uint16(anon_sym_RBRACK),
	678:  uint16(45),
	679:  uint16(2),
	680:  uint16(sym__bracket_char),
	681:  uint16(sym_bracket_char_escaped),
	682:  uint16(25),
	683:  uint16(4),
	684:  uint16(sym__bracket_pattern),
	685:  uint16(sym_bracket_range),
	686:  uint16(sym_bracket_char_class),
	687:  uint16(aux_sym_bracket_expr_repeat1),
	688:  uint16(6),
	689:  uint16(170),
	690:  uint16(1),
	691:  uint16(anon_sym_BSLASH),
	692:  uint16(173),
	693:  uint16(1),
	694:  uint16(anon_sym_RBRACK),
	695:  uint16(175),
	696:  uint16(1),
	697:  uint16(sym_bracket_char),
	698:  uint16(178),
	699:  uint16(1),
	700:  uint16(anon_sym_LBRACK_COLON),
	701:  uint16(45),
	702:  uint16(2),
	703:  uint16(sym__bracket_char),
	704:  uint16(sym_bracket_char_escaped),
	705:  uint16(25),
	706:  uint16(4),
	707:  uint16(sym__bracket_pattern),
	708:  uint16(sym_bracket_range),
	709:  uint16(sym_bracket_char_class),
	710:  uint16(aux_sym_bracket_expr_repeat1),
	711:  uint16(2),
	712:  uint16(183),
	713:  uint16(4),
	714:  uint16(sym_pattern_char),
	715:  uint16(anon_sym_BSLASH),
	716:  uint16(anon_sym_STAR),
	717:  uint16(sym__trailing_spaces),
	718:  uint16(181),
	719:  uint16(6),
	720:  uint16(sym_directory_separator),
	721:  uint16(sym_directory_separator_escaped),
	722:  uint16(anon_sym_QMARK),
	723:  uint16(anon_sym_STAR_STAR),
	724:  uint16(anon_sym_LBRACK),
	725:  uint16(sym__newline),
	726:  uint16(5),
	727:  uint16(17),
	728:  uint16(1),
	729:  uint16(anon_sym_LBRACK),
	730:  uint16(130),
	731:  uint16(1),
	732:  uint16(anon_sym_BSLASH),
	733:  uint16(185),
	734:  uint16(2),
	735:  uint16(sym_pattern_char),
	736:  uint16(anon_sym_STAR),
	737:  uint16(187),
	738:  uint16(2),
	739:  uint16(anon_sym_QMARK),
	740:  uint16(anon_sym_STAR_STAR),
	741:  uint16(6),
	742:  uint16(4),
	743:  uint16(aux_sym__pattern),
	744:  uint16(sym_pattern_char_escaped),
	745:  uint16(sym__wildcard),
	746:  uint16(sym_bracket_expr),
	747:  uint16(2),
	748:  uint16(191),
	749:  uint16(4),
	750:  uint16(sym_pattern_char),
	751:  uint16(anon_sym_BSLASH),
	752:  uint16(anon_sym_STAR),
	753:  uint16(sym__trailing_spaces),
	754:  uint16(189),
	755:  uint16(6),
	756:  uint16(sym_directory_separator),
	757:  uint16(sym_directory_separator_escaped),
	758:  uint16(anon_sym_QMARK),
	759:  uint16(anon_sym_STAR_STAR),
	760:  uint16(anon_sym_LBRACK),
	761:  uint16(sym__newline),
	762:  uint16(5),
	763:  uint16(17),
	764:  uint16(1),
	765:  uint16(anon_sym_LBRACK),
	766:  uint16(130),
	767:  uint16(1),
	768:  uint16(anon_sym_BSLASH),
	769:  uint16(128),
	770:  uint16(2),
	771:  uint16(sym_pattern_char),
	772:  uint16(anon_sym_STAR),
	773:  uint16(132),
	774:  uint16(2),
	775:  uint16(anon_sym_QMARK),
	776:  uint16(anon_sym_STAR_STAR),
	777:  uint16(9),
	778:  uint16(4),
	779:  uint16(aux_sym__pattern),
	780:  uint16(sym_pattern_char_escaped),
	781:  uint16(sym__wildcard),
	782:  uint16(sym_bracket_expr),
	783:  uint16(2),
	784:  uint16(195),
	785:  uint16(4),
	786:  uint16(sym_pattern_char),
	787:  uint16(anon_sym_BSLASH),
	788:  uint16(anon_sym_STAR),
	789:  uint16(sym__trailing_spaces),
	790:  uint16(193),
	791:  uint16(6),
	792:  uint16(sym_directory_separator),
	793:  uint16(sym_directory_separator_escaped),
	794:  uint16(anon_sym_QMARK),
	795:  uint16(anon_sym_STAR_STAR),
	796:  uint16(anon_sym_LBRACK),
	797:  uint16(sym__newline),
	798:  uint16(6),
	799:  uint16(86),
	800:  uint16(1),
	801:  uint16(anon_sym_BSLASH),
	802:  uint16(90),
	803:  uint16(1),
	804:  uint16(sym_bracket_char),
	805:  uint16(92),
	806:  uint16(1),
	807:  uint16(anon_sym_LBRACK_COLON),
	808:  uint16(197),
	809:  uint16(1),
	810:  uint16(anon_sym_RBRACK),
	811:  uint16(45),
	812:  uint16(2),
	813:  uint16(sym__bracket_char),
	814:  uint16(sym_bracket_char_escaped),
	815:  uint16(25),
	816:  uint16(4),
	817:  uint16(sym__bracket_pattern),
	818:  uint16(sym_bracket_range),
	819:  uint16(sym_bracket_char_class),
	820:  uint16(aux_sym_bracket_expr_repeat1),
	821:  uint16(2),
	822:  uint16(201),
	823:  uint16(4),
	824:  uint16(sym_pattern_char),
	825:  uint16(anon_sym_BSLASH),
	826:  uint16(anon_sym_STAR),
	827:  uint16(sym__trailing_spaces),
	828:  uint16(199),
	829:  uint16(6),
	830:  uint16(sym_directory_separator),
	831:  uint16(sym_directory_separator_escaped),
	832:  uint16(anon_sym_QMARK),
	833:  uint16(anon_sym_STAR_STAR),
	834:  uint16(anon_sym_LBRACK),
	835:  uint16(sym__newline),
	836:  uint16(6),
	837:  uint16(86),
	838:  uint16(1),
	839:  uint16(anon_sym_BSLASH),
	840:  uint16(90),
	841:  uint16(1),
	842:  uint16(sym_bracket_char),
	843:  uint16(92),
	844:  uint16(1),
	845:  uint16(anon_sym_LBRACK_COLON),
	846:  uint16(203),
	847:  uint16(1),
	848:  uint16(anon_sym_RBRACK),
	849:  uint16(45),
	850:  uint16(2),
	851:  uint16(sym__bracket_char),
	852:  uint16(sym_bracket_char_escaped),
	853:  uint16(31),
	854:  uint16(4),
	855:  uint16(sym__bracket_pattern),
	856:  uint16(sym_bracket_range),
	857:  uint16(sym_bracket_char_class),
	858:  uint16(aux_sym_bracket_expr_repeat1),
	859:  uint16(6),
	860:  uint16(86),
	861:  uint16(1),
	862:  uint16(anon_sym_BSLASH),
	863:  uint16(90),
	864:  uint16(1),
	865:  uint16(sym_bracket_char),
	866:  uint16(92),
	867:  uint16(1),
	868:  uint16(anon_sym_LBRACK_COLON),
	869:  uint16(203),
	870:  uint16(1),
	871:  uint16(anon_sym_RBRACK),
	872:  uint16(45),
	873:  uint16(2),
	874:  uint16(sym__bracket_char),
	875:  uint16(sym_bracket_char_escaped),
	876:  uint16(25),
	877:  uint16(4),
	878:  uint16(sym__bracket_pattern),
	879:  uint16(sym_bracket_range),
	880:  uint16(sym_bracket_char_class),
	881:  uint16(aux_sym_bracket_expr_repeat1),
	882:  uint16(6),
	883:  uint16(86),
	884:  uint16(1),
	885:  uint16(anon_sym_BSLASH),
	886:  uint16(90),
	887:  uint16(1),
	888:  uint16(sym_bracket_char),
	889:  uint16(92),
	890:  uint16(1),
	891:  uint16(anon_sym_LBRACK_COLON),
	892:  uint16(168),
	893:  uint16(1),
	894:  uint16(anon_sym_RBRACK),
	895:  uint16(45),
	896:  uint16(2),
	897:  uint16(sym__bracket_char),
	898:  uint16(sym_bracket_char_escaped),
	899:  uint16(38),
	900:  uint16(4),
	901:  uint16(sym__bracket_pattern),
	902:  uint16(sym_bracket_range),
	903:  uint16(sym_bracket_char_class),
	904:  uint16(aux_sym_bracket_expr_repeat1),
	905:  uint16(5),
	906:  uint16(17),
	907:  uint16(1),
	908:  uint16(anon_sym_LBRACK),
	909:  uint16(130),
	910:  uint16(1),
	911:  uint16(anon_sym_BSLASH),
	912:  uint16(205),
	913:  uint16(2),
	914:  uint16(sym_pattern_char),
	915:  uint16(anon_sym_STAR),
	916:  uint16(207),
	917:  uint16(2),
	918:  uint16(anon_sym_QMARK),
	919:  uint16(anon_sym_STAR_STAR),
	920:  uint16(5),
	921:  uint16(4),
	922:  uint16(aux_sym__pattern),
	923:  uint16(sym_pattern_char_escaped),
	924:  uint16(sym__wildcard),
	925:  uint16(sym_bracket_expr),
	926:  uint16(2),
	927:  uint16(211),
	928:  uint16(4),
	929:  uint16(sym_pattern_char),
	930:  uint16(anon_sym_BSLASH),
	931:  uint16(anon_sym_STAR),
	932:  uint16(sym__trailing_spaces),
	933:  uint16(209),
	934:  uint16(6),
	935:  uint16(sym_directory_separator),
	936:  uint16(sym_directory_separator_escaped),
	937:  uint16(anon_sym_QMARK),
	938:  uint16(anon_sym_STAR_STAR),
	939:  uint16(anon_sym_LBRACK),
	940:  uint16(sym__newline),
	941:  uint16(6),
	942:  uint16(86),
	943:  uint16(1),
	944:  uint16(anon_sym_BSLASH),
	945:  uint16(90),
	946:  uint16(1),
	947:  uint16(sym_bracket_char),
	948:  uint16(92),
	949:  uint16(1),
	950:  uint16(anon_sym_LBRACK_COLON),
	951:  uint16(213),
	952:  uint16(1),
	953:  uint16(anon_sym_RBRACK),
	954:  uint16(45),
	955:  uint16(2),
	956:  uint16(sym__bracket_char),
	957:  uint16(sym_bracket_char_escaped),
	958:  uint16(25),
	959:  uint16(4),
	960:  uint16(sym__bracket_pattern),
	961:  uint16(sym_bracket_range),
	962:  uint16(sym_bracket_char_class),
	963:  uint16(aux_sym_bracket_expr_repeat1),
	964:  uint16(4),
	965:  uint16(17),
	966:  uint16(1),
	967:  uint16(sym__directory_separator),
	968:  uint16(40),
	969:  uint16(1),
	970:  uint16(aux_sym_pattern_repeat1),
	971:  uint16(215),
	972:  uint16(2),
	973:  uint16(sym_directory_separator),
	974:  uint16(sym_directory_separator_escaped),
	975:  uint16(217),
	976:  uint16(2),
	977:  uint16(sym__trailing_spaces),
	978:  uint16(sym__newline),
	979:  uint16(4),
	980:  uint16(29),
	981:  uint16(1),
	982:  uint16(sym__directory_separator),
	983:  uint16(40),
	984:  uint16(1),
	985:  uint16(aux_sym_pattern_repeat1),
	986:  uint16(219),
	987:  uint16(2),
	988:  uint16(sym_directory_separator),
	989:  uint16(sym_directory_separator_escaped),
	990:  uint16(222),
	991:  uint16(2),
	992:  uint16(sym__trailing_spaces),
	993:  uint16(sym__newline),
	994:  uint16(4),
	995:  uint16(21),
	996:  uint16(1),
	997:  uint16(sym__directory_separator),
	998:  uint16(40),
	999:  uint16(1),
	1000: uint16(aux_sym_pattern_repeat1),
	1001: uint16(224),
	1002: uint16(2),
	1003: uint16(sym_directory_separator),
	1004: uint16(sym_directory_separator_escaped),
	1005: uint16(226),
	1006: uint16(2),
	1007: uint16(sym__trailing_spaces),
	1008: uint16(sym__newline),
	1009: uint16(4),
	1010: uint16(18),
	1011: uint16(1),
	1012: uint16(sym__directory_separator),
	1013: uint16(40),
	1014: uint16(1),
	1015: uint16(aux_sym_pattern_repeat1),
	1016: uint16(228),
	1017: uint16(2),
	1018: uint16(sym_directory_separator),
	1019: uint16(sym_directory_separator_escaped),
	1020: uint16(230),
	1021: uint16(2),
	1022: uint16(sym__trailing_spaces),
	1023: uint16(sym__newline),
	1024: uint16(4),
	1025: uint16(16),
	1026: uint16(1),
	1027: uint16(sym__directory_separator),
	1028: uint16(40),
	1029: uint16(1),
	1030: uint16(aux_sym_pattern_repeat1),
	1031: uint16(232),
	1032: uint16(2),
	1033: uint16(sym_directory_separator),
	1034: uint16(sym_directory_separator_escaped),
	1035: uint16(234),
	1036: uint16(2),
	1037: uint16(sym__trailing_spaces),
	1038: uint16(sym__newline),
	1039: uint16(2),
	1040: uint16(238),
	1041: uint16(1),
	1042: uint16(sym_bracket_char),
	1043: uint16(236),
	1044: uint16(4),
	1045: uint16(anon_sym_BSLASH),
	1046: uint16(anon_sym_RBRACK),
	1047: uint16(anon_sym_DASH),
	1048: uint16(anon_sym_LBRACK_COLON),
	1049: uint16(3),
	1050: uint16(242),
	1051: uint16(1),
	1052: uint16(anon_sym_DASH),
	1053: uint16(244),
	1054: uint16(1),
	1055: uint16(sym_bracket_char),
	1056: uint16(240),
	1057: uint16(3),
	1058: uint16(anon_sym_BSLASH),
	1059: uint16(anon_sym_RBRACK),
	1060: uint16(anon_sym_LBRACK_COLON),
	1061: uint16(3),
	1062: uint16(248),
	1063: uint16(1),
	1064: uint16(anon_sym_DASH),
	1065: uint16(250),
	1066: uint16(1),
	1067: uint16(sym_bracket_char),
	1068: uint16(246),
	1069: uint16(3),
	1070: uint16(anon_sym_BSLASH),
	1071: uint16(anon_sym_RBRACK),
	1072: uint16(anon_sym_LBRACK_COLON),
	1073: uint16(2),
	1074: uint16(254),
	1075: uint16(1),
	1076: uint16(sym_bracket_char),
	1077: uint16(252),
	1078: uint16(4),
	1079: uint16(anon_sym_BSLASH),
	1080: uint16(anon_sym_RBRACK),
	1081: uint16(anon_sym_DASH),
	1082: uint16(anon_sym_LBRACK_COLON),
	1083: uint16(3),
	1084: uint16(256),
	1085: uint16(1),
	1086: uint16(anon_sym_BSLASH),
	1087: uint16(258),
	1088: uint16(1),
	1089: uint16(sym_bracket_char),
	1090: uint16(50),
	1091: uint16(2),
	1092: uint16(sym__bracket_char),
	1093: uint16(sym_bracket_char_escaped),
	1094: uint16(2),
	1095: uint16(262),
	1096: uint16(1),
	1097: uint16(sym_bracket_char),
	1098: uint16(260),
	1099: uint16(3),
	1100: uint16(anon_sym_BSLASH),
	1101: uint16(anon_sym_RBRACK),
	1102: uint16(anon_sym_LBRACK_COLON),
	1103: uint16(2),
	1104: uint16(266),
	1105: uint16(1),
	1106: uint16(sym_bracket_char),
	1107: uint16(264),
	1108: uint16(3),
	1109: uint16(anon_sym_BSLASH),
	1110: uint16(anon_sym_RBRACK),
	1111: uint16(anon_sym_LBRACK_COLON),
	1112: uint16(2),
	1113: uint16(270),
	1114: uint16(1),
	1115: uint16(sym_bracket_char),
	1116: uint16(268),
	1117: uint16(3),
	1118: uint16(anon_sym_BSLASH),
	1119: uint16(anon_sym_RBRACK),
	1120: uint16(anon_sym_LBRACK_COLON),
	1121: uint16(3),
	1122: uint16(256),
	1123: uint16(1),
	1124: uint16(anon_sym_BSLASH),
	1125: uint16(272),
	1126: uint16(1),
	1127: uint16(sym_bracket_char),
	1128: uint16(51),
	1129: uint16(2),
	1130: uint16(sym__bracket_char),
	1131: uint16(sym_bracket_char_escaped),
	1132: uint16(2),
	1133: uint16(238),
	1134: uint16(1),
	1135: uint16(sym_bracket_char),
	1136: uint16(236),
	1137: uint16(3),
	1138: uint16(anon_sym_BSLASH),
	1139: uint16(anon_sym_RBRACK),
	1140: uint16(anon_sym_LBRACK_COLON),
	1141: uint16(2),
	1142: uint16(274),
	1143: uint16(1),
	1144: uint16(sym__trailing_spaces),
	1145: uint16(276),
	1146: uint16(1),
	1147: uint16(sym__newline),
	1148: uint16(1),
	1149: uint16(278),
	1150: uint16(1),
	1151: uint16(sym__newline),
	1152: uint16(1),
	1153: uint16(280),
	1154: uint16(1),
	1155: uint16(anon_sym_COLON_RBRACK),
	1156: uint16(1),
	1157: uint16(282),
	1158: uint16(1),
	1159: uint16(aux_sym_pattern_char_escaped_token1),
	1160: uint16(1),
	1161: uint16(284),
	1162: uint16(1),
	1163: uint16(aux_sym_pattern_char_escaped_token1),
	1164: uint16(1),
	1165: uint16(276),
	1166: uint16(1),
	1167: uint16(sym__newline),
	1168: uint16(1),
	1169: uint16(286),
	1170: uint16(1),
	1172: uint16(1),
	1173: uint16(288),
	1174: uint16(1),
	1175: uint16(aux_sym_pattern_char_escaped_token1),
}

var ts_small_parse_table_map = [60]uint32_t{
	1:  uint32(50),
	2:  uint32(100),
	3:  uint32(137),
	4:  uint32(174),
	5:  uint32(211),
	6:  uint32(248),
	7:  uint32(282),
	8:  uint32(311),
	9:  uint32(340),
	10: uint32(358),
	11: uint32(376),
	12: uint32(406),
	13: uint32(434),
	14: uint32(461),
	15: uint32(488),
	16: uint32(515),
	17: uint32(542),
	18: uint32(569),
	19: uint32(584),
	20: uint32(611),
	21: uint32(638),
	22: uint32(665),
	23: uint32(688),
	24: uint32(711),
	25: uint32(726),
	26: uint32(747),
	27: uint32(762),
	28: uint32(783),
	29: uint32(798),
	30: uint32(821),
	31: uint32(836),
	32: uint32(859),
	33: uint32(882),
	34: uint32(905),
	35: uint32(926),
	36: uint32(941),
	37: uint32(964),
	38: uint32(979),
	39: uint32(994),
	40: uint32(1009),
	41: uint32(1024),
	42: uint32(1039),
	43: uint32(1049),
	44: uint32(1061),
	45: uint32(1073),
	46: uint32(1083),
	47: uint32(1094),
	48: uint32(1103),
	49: uint32(1112),
	50: uint32(1121),
	51: uint32(1132),
	52: uint32(1141),
	53: uint32(1148),
	54: uint32(1152),
	55: uint32(1156),
	56: uint32(1160),
	57: uint32(1164),
	58: uint32(1168),
	59: uint32(1172),
}

var ts_parse_actions = [290]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_document),
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
		Fstate: uint16(54),
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
		Fstate: uint16(14),
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
		Fstate: uint16(27),
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
		Fstate: uint16(7),
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
		Fstate: uint16(57),
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
		Fstate: uint16(7),
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
		Fstate: uint16(8),
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
		Fstate: uint16(59),
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
		Fstate: uint16(3),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	26: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	27: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	28: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	29: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(14),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
	33: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(27),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	34: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	35: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(7),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fstate:      uint16(57),
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
		Fstate:      uint16(7),
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
		Fstate:      uint16(8),
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
		Fcount: uint8(2),
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
		Fstate:      uint16(59),
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
		Fstate:      uint16(2),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_document),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	55: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	56: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(15),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(10),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	61: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	62: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	63: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fsymbol:      uint16(sym_pattern),
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
		Fstate: uint16(19),
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
		Fcount: uint8(1),
	}})),
	69: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(5),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(5),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	73: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(2),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(2),
	})))),
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
		Fstate: uint16(23),
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
		Fcount: uint8(1),
	}})),
	81: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: uint16(13),
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
		Fstate: uint16(58),
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
		Fstate: uint16(47),
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
		Fcount: uint8(1),
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
		Fstate: uint16(45),
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
		Fstate: uint16(20),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_pattern_repeat1),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	97: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_pattern_repeat1),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	99: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fsymbol:      uint16(aux_sym__pattern),
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
		Fstate:      uint16(10),
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
		Fsymbol:      uint16(aux_sym__pattern),
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
		Fstate:      uint16(57),
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
		Fsymbol:      uint16(aux_sym__pattern),
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
		Fstate:      uint16(10),
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
		Fsymbol:      uint16(aux_sym__pattern),
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
		Fstate:      uint16(8),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym__pattern),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__line),
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
		Fcount: uint8(1),
	}})),
	117: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__line),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__line),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__line),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(4),
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
		Fstate: uint16(4),
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
		Fcount: uint8(1),
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
		Fstate: uint16(9),
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
		Fstate: uint16(57),
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
		Fstate: uint16(9),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(6),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(6),
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
		Fcount: uint8(1),
	}})),
	139: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(18),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	141: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(18),
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
		Fcount: uint8(1),
	}})),
	143: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(17),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	145: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(17),
	})))),
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
		Fcount: uint8(1),
	}})),
	147: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(14),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(14),
	})))),
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
		Fcount: uint8(1),
	}})),
	151: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(12),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(12),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(56),
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
		Fcount: uint8(1),
	}})),
	157: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(10),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(10),
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
		Fcount: uint8(1),
	}})),
	161: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(8),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(8),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(3),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(3),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(aux_sym_bracket_expr_repeat1),
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
		Fstate:      uint16(58),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_bracket_expr_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_bracket_expr_repeat1),
	})))),
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
		Fstate:      uint16(45),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_bracket_expr_repeat1),
	})))),
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
		Fstate:      uint16(20),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_bracket_expr),
		Fproduction_id: uint16(15),
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
		Fcount: uint8(1),
	}})),
	184: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_bracket_expr),
		Fproduction_id: uint16(15),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(6),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_pattern_char_escaped),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_pattern_char_escaped),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_bracket_expr),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_bracket_expr),
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
		Fstate: uint16(26),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_bracket_expr),
		Fproduction_id: uint16(15),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_bracket_expr),
		Fproduction_id: uint16(15),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(5),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_bracket_expr),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_bracket_expr),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(17),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	218: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(9),
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
		Fcount:    uint8(2),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_pattern_repeat1),
		Fproduction_id: uint16(11),
	})))),
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
		Fstate:      uint16(29),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_pattern_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(4),
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
		Fstate: uint16(18),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(7),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_pattern),
		Fproduction_id: uint16(13),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_bracket_char_escaped),
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
		Fcount: uint8(1),
	}})),
	239: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_bracket_char_escaped),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__bracket_pattern),
	})))),
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
		Fstate: uint16(52),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__bracket_pattern),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__bracket_pattern_closing_bracket),
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
		Fstate: uint16(48),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__bracket_pattern_closing_bracket),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__bracket_char_closing_bracket),
		Fproduction_id: uint16(1),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__bracket_char_closing_bracket),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(50),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_bracket_char_class),
		Fproduction_id: uint16(16),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_bracket_char_class),
		Fproduction_id: uint16(16),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__bracket_range_closing_bracket),
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
		Fcount: uint8(1),
	}})),
	267: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__bracket_range_closing_bracket),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_bracket_range),
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
		Fcount: uint8(1),
	}})),
	271: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_bracket_range),
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
		Fcount: uint8(1),
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
		Fstate: uint16(51),
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
		Fstate: uint16(55),
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
		Fstate: uint16(12),
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
		Fstate: uint16(11),
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
		Fstate: uint16(49),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(44),
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
	287: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: uint16(53),
	}})))),
}

func tree_sitter_gitignore(tls *libc.TLS) (r uintptr) {
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

var __ccgo_ts1 = "end\x00comment\x00negation\x00directory_separator\x00directory_separator_escaped\x00pattern_char\x00\\\x00pattern_char_escaped_token1\x00wildcard_char_single\x00wildcard_chars\x00wildcard_chars_allow_slash\x00[\x00bracket_negation\x00]\x00-\x00bracket_char\x00[:\x00alnum\x00:]\x00alpha\x00blank\x00cntrl\x00digit\x00graph\x00lower\x00print\x00punct\x00space\x00upper\x00xdigit\x00_trailing_spaces\x00_newline\x00document\x00_line\x00pattern\x00_directory_separator\x00_pattern\x00pattern_char_escaped\x00_wildcard\x00bracket_expr\x00_bracket_pattern\x00_bracket_pattern_closing_bracket\x00_bracket_char_closing_bracket\x00bracket_range\x00_bracket_char\x00bracket_char_escaped\x00bracket_char_class\x00document_repeat1\x00pattern_repeat1\x00bracket_expr_repeat1\x00directory_flag\x00name\x00relative_flag\x00"
