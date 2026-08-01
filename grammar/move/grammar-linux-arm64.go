// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-move/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-move -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-move/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_move

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
const EXTERNAL_TOKEN_COUNT = 0
const FIELD_COUNT = 21
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
const LARGE_STATE_COUNT = 3
const MAX_ALIAS_SEQUENCE_LENGTH = 8
const PRODUCTION_ID_COUNT = 38
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 357
const SYMBOL_COUNT = 163
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

const sym_identifier = 1
const anon_sym_module = 2
const anon_sym_COLON_COLON = 3
const anon_sym_LBRACE = 4
const anon_sym_RBRACE = 5
const anon_sym_SEMI = 6
const anon_sym_POUND = 7
const anon_sym_LBRACK = 8
const anon_sym_RBRACK = 9
const anon_sym_EQ = 10
const anon_sym_fun = 11
const anon_sym_COLON = 12
const anon_sym_public = 13
const anon_sym_LT = 14
const anon_sym_COMMA = 15
const anon_sym_GT = 16
const anon_sym_LPAREN = 17
const anon_sym__ = 18
const anon_sym_RPAREN = 19
const anon_sym_use = 20
const anon_sym_u8 = 21
const anon_sym_u64 = 22
const anon_sym_u128 = 23
const anon_sym_bool = 24
const anon_sym_address = 25
const anon_sym_as = 26
const anon_sym_LT2 = 27
const anon_sym_AMP = 28
const sym_mutable_specifier = 29
const anon_sym_BANG = 30
const anon_sym_AMP_AMP = 31
const anon_sym_PIPE_PIPE = 32
const anon_sym_PIPE = 33
const anon_sym_CARET = 34
const anon_sym_EQ_EQ = 35
const anon_sym_BANG_EQ = 36
const anon_sym_LT_EQ = 37
const anon_sym_GT_EQ = 38
const anon_sym_LT_LT = 39
const anon_sym_GT_GT = 40
const anon_sym_PLUS = 41
const anon_sym_DASH = 42
const anon_sym_STAR = 43
const anon_sym_SLASH = 44
const anon_sym_PERCENT = 45
const anon_sym_return = 46
const anon_sym_if = 47
const anon_sym_let = 48
const anon_sym_else = 49
const anon_sym_while = 50
const anon_sym_loop = 51
const anon_sym_const = 52
const anon_sym_break = 53
const anon_sym_continue = 54
const anon_sym_DOT = 55
const anon_sym_DOLLAR = 56
const aux_sym__non_special_token_token1 = 57
const anon_sym_SQUOTE = 58
const anon_sym_async = 59
const anon_sym_await = 60
const anon_sym_default = 61
const anon_sym_enum = 62
const anon_sym_fn = 63
const anon_sym_for = 64
const anon_sym_impl = 65
const anon_sym_match = 66
const anon_sym_mod = 67
const anon_sym_pub = 68
const anon_sym_static = 69
const anon_sym_struct = 70
const anon_sym_trait = 71
const anon_sym_type = 72
const anon_sym_union = 73
const anon_sym_unsafe = 74
const anon_sym_where = 75
const sym_integer_literal = 76
const sym_float_literal = 77
const anon_sym_true = 78
const anon_sym_false = 79
const anon_sym_0x = 80
const anon_sym_0X = 81
const aux_sym_hex_address_token1 = 82
const sym_comment = 83
const sym_module = 84
const sym_module_body = 85
const sym__statement = 86
const sym_expression_statement = 87
const sym__declaration_statement = 88
const sym_attribute_item = 89
const sym_attribute = 90
const sym_function_item = 91
const sym_visibility_modifier = 92
const sym_type_parameters = 93
const sym_parameters = 94
const sym_parameter = 95
const sym_use_declaration = 96
const sym__use_clause = 97
const sym__type = 98
const sym_bracketed_type = 99
const sym_qualified_type = 100
const sym_unit_type = 101
const sym_generic_function = 102
const sym_generic_type = 103
const sym_type_arguments = 104
const sym_reference_type = 105
const sym__expression_except_range = 106
const sym__expression = 107
const sym__expression_ending_with_block = 108
const sym_scoped_type_identifier = 109
const sym_unary_expression = 110
const sym_reference_expression = 111
const sym_binary_expression = 112
const sym_assignment_expression = 113
const sym_type_cast_expression = 114
const sym_return_expression = 115
const sym_call_expression = 116
const sym_arguments = 117
const sym_parenthesized_expression = 118
const sym_tuple_expression = 119
const sym_unit_expression = 120
const sym_if_expression = 121
const sym_let_condition = 122
const sym__condition = 123
const sym_else_clause = 124
const sym_while_expression = 125
const sym_loop_expression = 126
const sym_const_block = 127
const sym_break_expression = 128
const sym_continue_expression = 129
const sym_field_expression = 130
const sym_block = 131
const sym__pattern = 132
const sym_tuple_pattern = 133
const sym_struct_pattern = 134
const sym_field_pattern = 135
const sym_mut_pattern = 136
const sym_reference_pattern = 137
const sym_or_pattern = 138
const sym_delim_token_tree = 139
const sym__delim_tokens = 140
const sym__non_delim_token = 141
const sym_scoped_identifier = 142
const sym__non_special_token = 143
const sym__literal = 144
const sym__literal_pattern = 145
const sym_negative_literal = 146
const sym_boolean_literal = 147
const sym_hex_address = 148
const sym__path = 149
const sym__type_identifier = 150
const sym__field_identifier = 151
const aux_sym_module_body_repeat1 = 152
const aux_sym_type_parameters_repeat1 = 153
const aux_sym_parameters_repeat1 = 154
const aux_sym_type_arguments_repeat1 = 155
const aux_sym_arguments_repeat1 = 156
const aux_sym_tuple_expression_repeat1 = 157
const aux_sym_tuple_expression_repeat2 = 158
const aux_sym_block_repeat1 = 159
const aux_sym_tuple_pattern_repeat1 = 160
const aux_sym_struct_pattern_repeat1 = 161
const aux_sym_delim_token_tree_repeat1 = 162
const alias_sym_field_identifier = 163
const alias_sym_primitive_type = 164
const alias_sym_shorthand_field_identifier = 165
const alias_sym_type_identifier = 166

var ts_symbol_names = [167]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 15,
	3:   __ccgo_ts + 22,
	4:   __ccgo_ts + 25,
	5:   __ccgo_ts + 27,
	6:   __ccgo_ts + 29,
	7:   __ccgo_ts + 31,
	8:   __ccgo_ts + 33,
	9:   __ccgo_ts + 35,
	10:  __ccgo_ts + 37,
	11:  __ccgo_ts + 39,
	12:  __ccgo_ts + 43,
	13:  __ccgo_ts + 45,
	14:  __ccgo_ts + 52,
	15:  __ccgo_ts + 54,
	16:  __ccgo_ts + 56,
	17:  __ccgo_ts + 58,
	18:  __ccgo_ts + 60,
	19:  __ccgo_ts + 62,
	20:  __ccgo_ts + 64,
	21:  __ccgo_ts + 4,
	22:  __ccgo_ts + 4,
	23:  __ccgo_ts + 4,
	24:  __ccgo_ts + 4,
	25:  __ccgo_ts + 4,
	26:  __ccgo_ts + 68,
	27:  __ccgo_ts + 52,
	28:  __ccgo_ts + 71,
	29:  __ccgo_ts + 73,
	30:  __ccgo_ts + 91,
	31:  __ccgo_ts + 93,
	32:  __ccgo_ts + 96,
	33:  __ccgo_ts + 99,
	34:  __ccgo_ts + 101,
	35:  __ccgo_ts + 103,
	36:  __ccgo_ts + 106,
	37:  __ccgo_ts + 109,
	38:  __ccgo_ts + 112,
	39:  __ccgo_ts + 115,
	40:  __ccgo_ts + 118,
	41:  __ccgo_ts + 121,
	42:  __ccgo_ts + 123,
	43:  __ccgo_ts + 125,
	44:  __ccgo_ts + 127,
	45:  __ccgo_ts + 129,
	46:  __ccgo_ts + 131,
	47:  __ccgo_ts + 138,
	48:  __ccgo_ts + 141,
	49:  __ccgo_ts + 145,
	50:  __ccgo_ts + 150,
	51:  __ccgo_ts + 156,
	52:  __ccgo_ts + 161,
	53:  __ccgo_ts + 167,
	54:  __ccgo_ts + 173,
	55:  __ccgo_ts + 182,
	56:  __ccgo_ts + 184,
	57:  __ccgo_ts + 186,
	58:  __ccgo_ts + 212,
	59:  __ccgo_ts + 214,
	60:  __ccgo_ts + 220,
	61:  __ccgo_ts + 226,
	62:  __ccgo_ts + 234,
	63:  __ccgo_ts + 239,
	64:  __ccgo_ts + 242,
	65:  __ccgo_ts + 246,
	66:  __ccgo_ts + 251,
	67:  __ccgo_ts + 257,
	68:  __ccgo_ts + 261,
	69:  __ccgo_ts + 265,
	70:  __ccgo_ts + 272,
	71:  __ccgo_ts + 279,
	72:  __ccgo_ts + 285,
	73:  __ccgo_ts + 290,
	74:  __ccgo_ts + 296,
	75:  __ccgo_ts + 303,
	76:  __ccgo_ts + 309,
	77:  __ccgo_ts + 325,
	78:  __ccgo_ts + 339,
	79:  __ccgo_ts + 344,
	80:  __ccgo_ts + 350,
	81:  __ccgo_ts + 353,
	82:  __ccgo_ts + 356,
	83:  __ccgo_ts + 375,
	84:  __ccgo_ts + 15,
	85:  __ccgo_ts + 383,
	86:  __ccgo_ts + 395,
	87:  __ccgo_ts + 406,
	88:  __ccgo_ts + 427,
	89:  __ccgo_ts + 450,
	90:  __ccgo_ts + 465,
	91:  __ccgo_ts + 475,
	92:  __ccgo_ts + 489,
	93:  __ccgo_ts + 509,
	94:  __ccgo_ts + 525,
	95:  __ccgo_ts + 536,
	96:  __ccgo_ts + 546,
	97:  __ccgo_ts + 562,
	98:  __ccgo_ts + 574,
	99:  __ccgo_ts + 580,
	100: __ccgo_ts + 595,
	101: __ccgo_ts + 610,
	102: __ccgo_ts + 620,
	103: __ccgo_ts + 637,
	104: __ccgo_ts + 650,
	105: __ccgo_ts + 665,
	106: __ccgo_ts + 680,
	107: __ccgo_ts + 705,
	108: __ccgo_ts + 717,
	109: __ccgo_ts + 747,
	110: __ccgo_ts + 770,
	111: __ccgo_ts + 787,
	112: __ccgo_ts + 808,
	113: __ccgo_ts + 826,
	114: __ccgo_ts + 848,
	115: __ccgo_ts + 869,
	116: __ccgo_ts + 887,
	117: __ccgo_ts + 903,
	118: __ccgo_ts + 913,
	119: __ccgo_ts + 938,
	120: __ccgo_ts + 955,
	121: __ccgo_ts + 971,
	122: __ccgo_ts + 985,
	123: __ccgo_ts + 999,
	124: __ccgo_ts + 1010,
	125: __ccgo_ts + 1022,
	126: __ccgo_ts + 1039,
	127: __ccgo_ts + 1055,
	128: __ccgo_ts + 1067,
	129: __ccgo_ts + 1084,
	130: __ccgo_ts + 1104,
	131: __ccgo_ts + 1121,
	132: __ccgo_ts + 1127,
	133: __ccgo_ts + 1136,
	134: __ccgo_ts + 1150,
	135: __ccgo_ts + 1165,
	136: __ccgo_ts + 1179,
	137: __ccgo_ts + 1191,
	138: __ccgo_ts + 1209,
	139: __ccgo_ts + 1220,
	140: __ccgo_ts + 1231,
	141: __ccgo_ts + 1245,
	142: __ccgo_ts + 1262,
	143: __ccgo_ts + 1280,
	144: __ccgo_ts + 1299,
	145: __ccgo_ts + 1308,
	146: __ccgo_ts + 1325,
	147: __ccgo_ts + 1342,
	148: __ccgo_ts + 1358,
	149: __ccgo_ts + 1370,
	150: __ccgo_ts + 1376,
	151: __ccgo_ts + 1393,
	152: __ccgo_ts + 1411,
	153: __ccgo_ts + 1431,
	154: __ccgo_ts + 1455,
	155: __ccgo_ts + 1474,
	156: __ccgo_ts + 1497,
	157: __ccgo_ts + 1515,
	158: __ccgo_ts + 1540,
	159: __ccgo_ts + 1565,
	160: __ccgo_ts + 1579,
	161: __ccgo_ts + 1601,
	162: __ccgo_ts + 1624,
	163: __ccgo_ts + 1649,
	164: __ccgo_ts + 1666,
	165: __ccgo_ts + 1681,
	166: __ccgo_ts + 1708,
}

var ts_symbol_map = [167]TSSymbol{
	1:   uint16(sym_identifier),
	2:   uint16(anon_sym_module),
	3:   uint16(anon_sym_COLON_COLON),
	4:   uint16(anon_sym_LBRACE),
	5:   uint16(anon_sym_RBRACE),
	6:   uint16(anon_sym_SEMI),
	7:   uint16(anon_sym_POUND),
	8:   uint16(anon_sym_LBRACK),
	9:   uint16(anon_sym_RBRACK),
	10:  uint16(anon_sym_EQ),
	11:  uint16(anon_sym_fun),
	12:  uint16(anon_sym_COLON),
	13:  uint16(anon_sym_public),
	14:  uint16(anon_sym_LT),
	15:  uint16(anon_sym_COMMA),
	16:  uint16(anon_sym_GT),
	17:  uint16(anon_sym_LPAREN),
	18:  uint16(anon_sym__),
	19:  uint16(anon_sym_RPAREN),
	20:  uint16(anon_sym_use),
	21:  uint16(sym_identifier),
	22:  uint16(sym_identifier),
	23:  uint16(sym_identifier),
	24:  uint16(sym_identifier),
	25:  uint16(sym_identifier),
	26:  uint16(anon_sym_as),
	27:  uint16(anon_sym_LT),
	28:  uint16(anon_sym_AMP),
	29:  uint16(sym_mutable_specifier),
	30:  uint16(anon_sym_BANG),
	31:  uint16(anon_sym_AMP_AMP),
	32:  uint16(anon_sym_PIPE_PIPE),
	33:  uint16(anon_sym_PIPE),
	34:  uint16(anon_sym_CARET),
	35:  uint16(anon_sym_EQ_EQ),
	36:  uint16(anon_sym_BANG_EQ),
	37:  uint16(anon_sym_LT_EQ),
	38:  uint16(anon_sym_GT_EQ),
	39:  uint16(anon_sym_LT_LT),
	40:  uint16(anon_sym_GT_GT),
	41:  uint16(anon_sym_PLUS),
	42:  uint16(anon_sym_DASH),
	43:  uint16(anon_sym_STAR),
	44:  uint16(anon_sym_SLASH),
	45:  uint16(anon_sym_PERCENT),
	46:  uint16(anon_sym_return),
	47:  uint16(anon_sym_if),
	48:  uint16(anon_sym_let),
	49:  uint16(anon_sym_else),
	50:  uint16(anon_sym_while),
	51:  uint16(anon_sym_loop),
	52:  uint16(anon_sym_const),
	53:  uint16(anon_sym_break),
	54:  uint16(anon_sym_continue),
	55:  uint16(anon_sym_DOT),
	56:  uint16(anon_sym_DOLLAR),
	57:  uint16(aux_sym__non_special_token_token1),
	58:  uint16(anon_sym_SQUOTE),
	59:  uint16(anon_sym_async),
	60:  uint16(anon_sym_await),
	61:  uint16(anon_sym_default),
	62:  uint16(anon_sym_enum),
	63:  uint16(anon_sym_fn),
	64:  uint16(anon_sym_for),
	65:  uint16(anon_sym_impl),
	66:  uint16(anon_sym_match),
	67:  uint16(anon_sym_mod),
	68:  uint16(anon_sym_pub),
	69:  uint16(anon_sym_static),
	70:  uint16(anon_sym_struct),
	71:  uint16(anon_sym_trait),
	72:  uint16(anon_sym_type),
	73:  uint16(anon_sym_union),
	74:  uint16(anon_sym_unsafe),
	75:  uint16(anon_sym_where),
	76:  uint16(sym_integer_literal),
	77:  uint16(sym_float_literal),
	78:  uint16(anon_sym_true),
	79:  uint16(anon_sym_false),
	80:  uint16(anon_sym_0x),
	81:  uint16(anon_sym_0X),
	82:  uint16(aux_sym_hex_address_token1),
	83:  uint16(sym_comment),
	84:  uint16(sym_module),
	85:  uint16(sym_module_body),
	86:  uint16(sym__statement),
	87:  uint16(sym_expression_statement),
	88:  uint16(sym__declaration_statement),
	89:  uint16(sym_attribute_item),
	90:  uint16(sym_attribute),
	91:  uint16(sym_function_item),
	92:  uint16(sym_visibility_modifier),
	93:  uint16(sym_type_parameters),
	94:  uint16(sym_parameters),
	95:  uint16(sym_parameter),
	96:  uint16(sym_use_declaration),
	97:  uint16(sym__use_clause),
	98:  uint16(sym__type),
	99:  uint16(sym_bracketed_type),
	100: uint16(sym_qualified_type),
	101: uint16(sym_unit_type),
	102: uint16(sym_generic_function),
	103: uint16(sym_generic_type),
	104: uint16(sym_type_arguments),
	105: uint16(sym_reference_type),
	106: uint16(sym__expression_except_range),
	107: uint16(sym__expression),
	108: uint16(sym__expression_ending_with_block),
	109: uint16(sym_scoped_type_identifier),
	110: uint16(sym_unary_expression),
	111: uint16(sym_reference_expression),
	112: uint16(sym_binary_expression),
	113: uint16(sym_assignment_expression),
	114: uint16(sym_type_cast_expression),
	115: uint16(sym_return_expression),
	116: uint16(sym_call_expression),
	117: uint16(sym_arguments),
	118: uint16(sym_parenthesized_expression),
	119: uint16(sym_tuple_expression),
	120: uint16(sym_unit_expression),
	121: uint16(sym_if_expression),
	122: uint16(sym_let_condition),
	123: uint16(sym__condition),
	124: uint16(sym_else_clause),
	125: uint16(sym_while_expression),
	126: uint16(sym_loop_expression),
	127: uint16(sym_const_block),
	128: uint16(sym_break_expression),
	129: uint16(sym_continue_expression),
	130: uint16(sym_field_expression),
	131: uint16(sym_block),
	132: uint16(sym__pattern),
	133: uint16(sym_tuple_pattern),
	134: uint16(sym_struct_pattern),
	135: uint16(sym_field_pattern),
	136: uint16(sym_mut_pattern),
	137: uint16(sym_reference_pattern),
	138: uint16(sym_or_pattern),
	139: uint16(sym_delim_token_tree),
	140: uint16(sym__delim_tokens),
	141: uint16(sym__non_delim_token),
	142: uint16(sym_scoped_identifier),
	143: uint16(sym__non_special_token),
	144: uint16(sym__literal),
	145: uint16(sym__literal_pattern),
	146: uint16(sym_negative_literal),
	147: uint16(sym_boolean_literal),
	148: uint16(sym_hex_address),
	149: uint16(sym__path),
	150: uint16(sym__type_identifier),
	151: uint16(sym__field_identifier),
	152: uint16(aux_sym_module_body_repeat1),
	153: uint16(aux_sym_type_parameters_repeat1),
	154: uint16(aux_sym_parameters_repeat1),
	155: uint16(aux_sym_type_arguments_repeat1),
	156: uint16(aux_sym_arguments_repeat1),
	157: uint16(aux_sym_tuple_expression_repeat1),
	158: uint16(aux_sym_tuple_expression_repeat2),
	159: uint16(aux_sym_block_repeat1),
	160: uint16(aux_sym_tuple_pattern_repeat1),
	161: uint16(aux_sym_struct_pattern_repeat1),
	162: uint16(aux_sym_delim_token_tree_repeat1),
	163: uint16(alias_sym_field_identifier),
	164: uint16(alias_sym_primitive_type),
	165: uint16(alias_sym_shorthand_field_identifier),
	166: uint16(alias_sym_type_identifier),
}

var ts_symbol_metadata = [167]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	57: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	77: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	82: {},
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	98: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	141: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	142: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	143: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	144: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	145: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	150: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	151: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
}

const field_alias = 1
const field_alternative = 2
const field_argument = 3
const field_arguments = 4
const field_body = 5
const field_condition = 6
const field_consequence = 7
const field_field = 8
const field_function = 9
const field_left = 10
const field_name = 11
const field_operator = 12
const field_parameters = 13
const field_path = 14
const field_pattern = 15
const field_return_type = 16
const field_right = 17
const field_type = 18
const field_type_arguments = 19
const field_type_parameters = 20
const field_value = 21

var ts_field_names = [22]uintptr{
	0:  libc.UintptrFromInt32(0),
	1:  __ccgo_ts + 1724,
	2:  __ccgo_ts + 1730,
	3:  __ccgo_ts + 1742,
	4:  __ccgo_ts + 903,
	5:  __ccgo_ts + 1751,
	6:  __ccgo_ts + 1756,
	7:  __ccgo_ts + 1766,
	8:  __ccgo_ts + 1778,
	9:  __ccgo_ts + 1784,
	10: __ccgo_ts + 1793,
	11: __ccgo_ts + 1798,
	12: __ccgo_ts + 1803,
	13: __ccgo_ts + 525,
	14: __ccgo_ts + 1812,
	15: __ccgo_ts + 1817,
	16: __ccgo_ts + 1825,
	17: __ccgo_ts + 1837,
	18: __ccgo_ts + 285,
	19: __ccgo_ts + 650,
	20: __ccgo_ts + 509,
	21: __ccgo_ts + 1843,
}

var ts_field_map_slices = [38]TSFieldMapSlice{
	1: {
		Flength: uint16(1),
	},
	4: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(2),
		Flength: uint16(3),
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
		Flength: uint16(4),
	},
	9: {
		Findex:  uint16(12),
		Flength: uint16(2),
	},
	10: {
		Findex:  uint16(14),
		Flength: uint16(3),
	},
	11: {
		Findex:  uint16(17),
		Flength: uint16(1),
	},
	12: {
		Findex:  uint16(18),
		Flength: uint16(1),
	},
	13: {
		Findex:  uint16(19),
		Flength: uint16(1),
	},
	14: {
		Findex:  uint16(20),
		Flength: uint16(2),
	},
	15: {
		Findex:  uint16(22),
		Flength: uint16(4),
	},
	16: {
		Findex:  uint16(26),
		Flength: uint16(4),
	},
	17: {
		Findex:  uint16(30),
		Flength: uint16(5),
	},
	18: {
		Findex:  uint16(35),
		Flength: uint16(1),
	},
	19: {
		Findex:  uint16(36),
		Flength: uint16(2),
	},
	20: {
		Findex:  uint16(38),
		Flength: uint16(2),
	},
	21: {
		Findex:  uint16(40),
		Flength: uint16(3),
	},
	22: {
		Findex:  uint16(43),
		Flength: uint16(2),
	},
	23: {
		Findex:  uint16(45),
		Flength: uint16(2),
	},
	25: {
		Findex:  uint16(47),
		Flength: uint16(4),
	},
	26: {
		Findex:  uint16(51),
		Flength: uint16(5),
	},
	27: {
		Findex:  uint16(56),
		Flength: uint16(1),
	},
	28: {
		Findex:  uint16(57),
		Flength: uint16(2),
	},
	29: {
		Findex:  uint16(59),
		Flength: uint16(2),
	},
	30: {
		Findex:  uint16(61),
		Flength: uint16(2),
	},
	31: {
		Findex:  uint16(63),
		Flength: uint16(1),
	},
	32: {
		Findex:  uint16(64),
		Flength: uint16(1),
	},
	33: {
		Findex:  uint16(65),
		Flength: uint16(2),
	},
	34: {
		Findex:  uint16(67),
		Flength: uint16(3),
	},
	35: {
		Flength: uint16(1),
	},
	36: {
		Findex:  uint16(70),
		Flength: uint16(2),
	},
	37: {
		Findex:  uint16(72),
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [74]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	1: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(1),
	},
	2: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(3),
	},
	3: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	4: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(2),
	},
	5: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(1),
	},
	6: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	7: {
		Ffield_id: uint16(field_path),
	},
	8: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(4),
	},
	9: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	10: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(3),
	},
	11: {
		Ffield_id:    uint16(field_type_parameters),
		Fchild_index: uint8(2),
	},
	12: {
		Ffield_id:    uint16(field_alias),
		Fchild_index: uint8(2),
	},
	13: {
		Ffield_id: uint16(field_type),
	},
	14: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(4),
	},
	15: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	16: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(3),
	},
	17: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(2),
	},
	18: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
	},
	19: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(1),
	},
	20: {
		Ffield_id:    uint16(field_arguments),
		Fchild_index: uint8(1),
	},
	21: {
		Ffield_id: uint16(field_function),
	},
	22: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(5),
	},
	23: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	24: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(2),
	},
	25: {
		Ffield_id:    uint16(field_return_type),
		Fchild_index: uint8(4),
	},
	26: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(5),
	},
	27: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	28: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(4),
	},
	29: {
		Ffield_id:    uint16(field_type_parameters),
		Fchild_index: uint8(3),
	},
	30: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(6),
	},
	31: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	32: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(3),
	},
	33: {
		Ffield_id:    uint16(field_return_type),
		Fchild_index: uint8(5),
	},
	34: {
		Ffield_id:    uint16(field_type_parameters),
		Fchild_index: uint8(2),
	},
	35: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	36: {
		Ffield_id: uint16(field_function),
	},
	37: {
		Ffield_id:    uint16(field_type_arguments),
		Fchild_index: uint8(2),
	},
	38: {
		Ffield_id: uint16(field_left),
	},
	39: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	40: {
		Ffield_id: uint16(field_left),
	},
	41: {
		Ffield_id:    uint16(field_operator),
		Fchild_index: uint8(1),
	},
	42: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	43: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(2),
	},
	44: {
		Ffield_id: uint16(field_value),
	},
	45: {
		Ffield_id:    uint16(field_field),
		Fchild_index: uint8(2),
	},
	46: {
		Ffield_id: uint16(field_value),
	},
	47: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(6),
	},
	48: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	49: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(3),
	},
	50: {
		Ffield_id:    uint16(field_return_type),
		Fchild_index: uint8(5),
	},
	51: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(7),
	},
	52: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	53: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(4),
	},
	54: {
		Ffield_id:    uint16(field_return_type),
		Fchild_index: uint8(6),
	},
	55: {
		Ffield_id:    uint16(field_type_parameters),
		Fchild_index: uint8(3),
	},
	56: {
		Ffield_id:    uint16(field_arguments),
		Fchild_index: uint8(1),
	},
	57: {
		Ffield_id: uint16(field_type),
	},
	58: {
		Ffield_id:    uint16(field_type_arguments),
		Fchild_index: uint8(1),
	},
	59: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(2),
	},
	60: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(4),
	},
	61: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(4),
	},
	62: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(2),
	},
	63: {
		Ffield_id: uint16(field_type),
	},
	64: {
		Ffield_id: uint16(field_name),
	},
	65: {
		Ffield_id:    uint16(field_pattern),
		Fchild_index: uint8(1),
	},
	66: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(3),
	},
	67: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(5),
	},
	68: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(2),
	},
	69: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(4),
	},
	70: {
		Ffield_id: uint16(field_name),
	},
	71: {
		Ffield_id:    uint16(field_pattern),
		Fchild_index: uint8(2),
	},
	72: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	73: {
		Ffield_id:    uint16(field_pattern),
		Fchild_index: uint8(3),
	},
}

var ts_alias_sequences = [38][8]TSSymbol{
	0: {},
	2: {
		0: uint16(alias_sym_primitive_type),
	},
	3: {
		0: uint16(alias_sym_type_identifier),
	},
	24: {
		0: uint16(alias_sym_field_identifier),
	},
	32: {
		0: uint16(alias_sym_shorthand_field_identifier),
	},
	35: {
		1: uint16(alias_sym_shorthand_field_identifier),
	},
}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [357]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(3),
	4:   uint16(4),
	5:   uint16(4),
	6:   uint16(6),
	7:   uint16(6),
	8:   uint16(4),
	9:   uint16(6),
	10:  uint16(10),
	11:  uint16(11),
	12:  uint16(12),
	13:  uint16(12),
	14:  uint16(11),
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
	54:  uint16(52),
	55:  uint16(51),
	56:  uint16(50),
	57:  uint16(57),
	58:  uint16(58),
	59:  uint16(53),
	60:  uint16(58),
	61:  uint16(57),
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
	76:  uint16(43),
	77:  uint16(42),
	78:  uint16(43),
	79:  uint16(42),
	80:  uint16(40),
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
	114: uint16(114),
	115: uint16(115),
	116: uint16(46),
	117: uint16(117),
	118: uint16(118),
	119: uint16(119),
	120: uint16(120),
	121: uint16(121),
	122: uint16(122),
	123: uint16(123),
	124: uint16(124),
	125: uint16(47),
	126: uint16(48),
	127: uint16(41),
	128: uint16(128),
	129: uint16(129),
	130: uint16(45),
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
	171: uint16(168),
	172: uint16(163),
	173: uint16(173),
	174: uint16(167),
	175: uint16(173),
	176: uint16(176),
	177: uint16(170),
	178: uint16(164),
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
	189: uint16(137),
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
	203: uint16(128),
	204: uint16(204),
	205: uint16(129),
	206: uint16(206),
	207: uint16(207),
	208: uint16(110),
	209: uint16(97),
	210: uint16(99),
	211: uint16(211),
	212: uint16(212),
	213: uint16(213),
	214: uint16(212),
	215: uint16(91),
	216: uint16(211),
	217: uint16(89),
	218: uint16(81),
	219: uint16(95),
	220: uint16(220),
	221: uint16(221),
	222: uint16(87),
	223: uint16(213),
	224: uint16(93),
	225: uint16(100),
	226: uint16(226),
	227: uint16(227),
	228: uint16(228),
	229: uint16(229),
	230: uint16(230),
	231: uint16(231),
	232: uint16(232),
	233: uint16(233),
	234: uint16(234),
	235: uint16(234),
	236: uint16(236),
	237: uint16(237),
	238: uint16(238),
	239: uint16(239),
	240: uint16(233),
	241: uint16(241),
	242: uint16(232),
	243: uint16(243),
	244: uint16(244),
	245: uint16(245),
	246: uint16(246),
	247: uint16(247),
	248: uint16(248),
	249: uint16(249),
	250: uint16(250),
	251: uint16(251),
	252: uint16(252),
	253: uint16(253),
	254: uint16(254),
	255: uint16(255),
	256: uint16(256),
	257: uint16(257),
	258: uint16(258),
	259: uint16(259),
	260: uint16(229),
	261: uint16(261),
	262: uint16(262),
	263: uint16(263),
	264: uint16(264),
	265: uint16(265),
	266: uint16(266),
	267: uint16(267),
	268: uint16(268),
	269: uint16(269),
	270: uint16(270),
	271: uint16(271),
	272: uint16(272),
	273: uint16(273),
	274: uint16(274),
	275: uint16(275),
	276: uint16(276),
	277: uint16(276),
	278: uint16(278),
	279: uint16(279),
	280: uint16(280),
	281: uint16(281),
	282: uint16(280),
	283: uint16(274),
	284: uint16(284),
	285: uint16(285),
	286: uint16(286),
	287: uint16(266),
	288: uint16(288),
	289: uint16(289),
	290: uint16(286),
	291: uint16(275),
	292: uint16(292),
	293: uint16(265),
	294: uint16(294),
	295: uint16(84),
	296: uint16(296),
	297: uint16(297),
	298: uint16(276),
	299: uint16(299),
	300: uint16(278),
	301: uint16(301),
	302: uint16(302),
	303: uint16(303),
	304: uint16(273),
	305: uint16(284),
	306: uint16(306),
	307: uint16(307),
	308: uint16(308),
	309: uint16(309),
	310: uint16(310),
	311: uint16(311),
	312: uint16(312),
	313: uint16(313),
	314: uint16(314),
	315: uint16(306),
	316: uint16(316),
	317: uint16(317),
	318: uint16(318),
	319: uint16(319),
	320: uint16(320),
	321: uint16(321),
	322: uint16(322),
	323: uint16(64),
	324: uint16(324),
	325: uint16(325),
	326: uint16(326),
	327: uint16(327),
	328: uint16(328),
	329: uint16(329),
	330: uint16(313),
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
	342: uint16(342),
	343: uint16(343),
	344: uint16(63),
	345: uint16(345),
	346: uint16(346),
	347: uint16(347),
	348: uint16(321),
	349: uint16(320),
	350: uint16(336),
	351: uint16(317),
	352: uint16(352),
	353: uint16(353),
	354: uint16(333),
	355: uint16(327),
	356: uint16(337),
}

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
			state = uint16(127)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(165)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(132)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(200)
			goto next_state
		}
		if lookahead == int32('%') {
			state = uint16(181)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(206)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(199)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(253)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(137)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(131)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(161)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(133)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(170)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(256)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(262)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(263)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(258)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(261)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(260)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(300)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(296)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(277)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(356)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(299)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(346)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(339)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(268)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(303)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(130)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') {
			state = uint16(266)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(254)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('!') {
			state = uint16(165)
			goto next_state
		}
		if lookahead == int32('%') {
			state = uint16(181)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(199)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(242)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(131)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(170)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(286)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(324)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(326)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(301)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(328)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(299)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(305)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(243)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('d') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('!') {
			state = uint16(165)
			goto next_state
		}
		if lookahead == int32('%') {
			state = uint16(181)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(199)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(242)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(131)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(170)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(286)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(324)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(326)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(301)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(328)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(299)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(269)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(305)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(130)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(2)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(243)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('d') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('!') {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(132)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(123)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(242)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(138)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(287)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(324)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(326)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(301)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(328)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(299)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(305)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(243)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('d') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('!') {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(123)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(242)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(138)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(287)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(324)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(326)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(301)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(328)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(299)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(269)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(305)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(130)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(243)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('d') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('!') {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(123)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(242)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(138)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(287)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(324)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(326)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(301)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(296)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(299)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(305)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(243)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('d') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('!') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('%') {
			state = uint16(181)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(198)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(131)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(170)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(102)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(130)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('!') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('%') {
			state = uint16(181)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(198)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(131)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(170)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(341)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(343)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(130)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('!') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('%') {
			state = uint16(181)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(179)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(198)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(131)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(170)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(341)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('$') {
			state = uint16(200)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(206)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(201)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(133)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(90)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(62)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(105)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(100)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(24)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(130)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('!') || int32('#') <= lookahead && lookahead <= int32('.') || int32(':') <= lookahead && lookahead <= int32('@') || lookahead == int32('^') || lookahead == int32('_') || int32('|') <= lookahead && lookahead <= int32('~') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('&') {
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(123)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(242)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(138)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(287)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(325)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(331)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(243)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('d') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('&') {
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(123)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(242)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(138)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(287)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(325)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(270)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(243)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('c') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('(') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(131)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(161)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(133)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(130)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32(')') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(161)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(102)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('*') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(364)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('*') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(363)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('*') {
			state = uint16(15)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('.') {
			state = uint16(250)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(18)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(254)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('.') {
			state = uint16(250)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('/') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(244)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(19)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(247)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('/') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(343)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(130)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(20)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('/') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(21)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('1') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('1') {
			state = uint16(26)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('1') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(147)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('2') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('2') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('2') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('4') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('4') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('4') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('8') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('8') {
			state = uint16(241)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('8') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32(':') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('=') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('X') {
			state = uint16(252)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('a') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('a') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('a') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('a') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('a') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('a') {
			state = uint16(114)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('a') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('b') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('c') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('c') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('c') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('c') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('d') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(159)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('d') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('d') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('e') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('e') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('e') {
			state = uint16(233)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('e') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('e') {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('e') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('e') {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('e') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('e') {
			state = uint16(106)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('e') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('e') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('e') {
			state = uint16(99)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('f') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('f') {
			state = uint16(184)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('f') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('f') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('h') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('h') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('i') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('i') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('i') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('i') {
			state = uint16(92)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('i') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('k') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('l') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('l') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('l') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead == int32('l') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead == int32('m') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead == int32('n') {
			state = uint16(215)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead == int32('n') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('n') {
			state = uint16(235)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('n') {
			state = uint16(182)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('n') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('n') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead == int32('n') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead == int32('o') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead == int32('o') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('o') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('o') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('o') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('p') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead == int32('p') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead == int32('p') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead == int32('r') {
			state = uint16(217)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead == int32('r') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead == int32('r') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead == int32('r') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead == int32('r') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('s') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead == int32('s') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead == int32('s') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead == int32('s') {
			state = uint16(109)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead == int32('t') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead == int32('t') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead == int32('t') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead == int32('t') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('t') {
			state = uint16(192)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead == int32('t') {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead == int32('t') {
			state = uint16(229)
			goto next_state
		}
		return result
	case int32(112):
		if lookahead == int32('t') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead == int32('t') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('t') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('u') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead == int32('u') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead == int32('u') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead == int32('u') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('u') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('u') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('0') || lookahead == int32('1') || lookahead == int32('_') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(122):
		if int32('0') <= lookahead && lookahead <= int32('7') || lookahead == int32('_') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(123):
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(124):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(125):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(248)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead != 0 && lookahead != int32('\r') {
			state = uint16(364)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(365)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(132):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(138):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(139):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('<') {
			state = uint16(175)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(140):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(141):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(142):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(174)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(143):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(144):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_use)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(146):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_use)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(147):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u8)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(148):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u8)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(149):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(150):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(151):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u128)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(152):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u128)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(153):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bool)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(154):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bool)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(155):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_address)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(156):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_address)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(157):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_as)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(158):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_as)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('y') {
			state = uint16(321)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(159):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_as)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('y') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(160):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_as)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(161):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(162):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(163):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('&') {
			state = uint16(166)
			goto next_state
		}
		return result
	case int32(164):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(165):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(166):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(167):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(168):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(169):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('|') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(170):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(171):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(172):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(173):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(174):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(175):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(176):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(177):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(178):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(179):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(180):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(364)
			goto next_state
		}
		return result
	case int32(181):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(182):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_return)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(183):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_return)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(184):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_if)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(185):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_if)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(186):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_let)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(187):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_let)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(188):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_while)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(189):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_while)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(190):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_loop)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(191):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_loop)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(192):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_const)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(193):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_const)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(194):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_break)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(195):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_break)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(196):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_continue)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(197):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_continue)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(198):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(199):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(200):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(201):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__non_special_token_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(203)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(204)
			goto next_state
		}
		if lookahead == int32('!') || lookahead == int32('#') || lookahead == int32('%') || lookahead == int32('&') || int32('+') <= lookahead && lookahead <= int32('.') || int32(':') <= lookahead && lookahead <= int32('@') || lookahead == int32('^') || lookahead == int32('_') || lookahead == int32('|') || lookahead == int32('~') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(202):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__non_special_token_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(202)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(205)
			goto next_state
		}
		if lookahead == int32('!') || lookahead == int32('#') || lookahead == int32('%') || lookahead == int32('&') || int32('+') <= lookahead && lookahead <= int32('.') || int32(':') <= lookahead && lookahead <= int32('@') || lookahead == int32('^') || lookahead == int32('_') || lookahead == int32('|') || lookahead == int32('~') {
			state = uint16(203)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(203):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__non_special_token_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(202)
			goto next_state
		}
		if lookahead == int32('!') || lookahead == int32('#') || lookahead == int32('%') || lookahead == int32('&') || int32('+') <= lookahead && lookahead <= int32('/') || int32(':') <= lookahead && lookahead <= int32('@') || lookahead == int32('^') || lookahead == int32('_') || lookahead == int32('|') || lookahead == int32('~') {
			state = uint16(203)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(204):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__non_special_token_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\\') {
			state = uint16(126)
			goto next_state
		}
		if lookahead == int32('!') || lookahead == int32('#') || lookahead == int32('%') || lookahead == int32('&') || int32('*') <= lookahead && lookahead <= int32('/') || int32(':') <= lookahead && lookahead <= int32('@') || lookahead == int32('^') || lookahead == int32('_') || lookahead == int32('|') || lookahead == int32('~') {
			state = uint16(204)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(364)
			goto next_state
		}
		return result
	case int32(205):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__non_special_token_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') || lookahead == int32('#') || lookahead == int32('%') || lookahead == int32('&') || int32('*') <= lookahead && lookahead <= int32('/') || int32(':') <= lookahead && lookahead <= int32('@') || lookahead == int32('^') || lookahead == int32('_') || lookahead == int32('|') || lookahead == int32('~') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(206):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(207):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_async)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(208):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_async)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(209):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_await)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(210):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_await)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(211):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_default)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(212):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_default)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(213):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(214):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(215):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_fn)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(216):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_fn)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(217):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_for)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(218):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_for)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(219):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_impl)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(220):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_impl)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(221):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_match)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(222):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_match)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(223):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mod)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(224):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mod)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(225):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_pub)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(226):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_pub)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(227):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_static)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(228):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_static)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(229):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_struct)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(230):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_struct)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(231):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_trait)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(232):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_trait)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(233):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(234):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(235):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_union)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(236):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_union)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(237):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_unsafe)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(238):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_unsafe)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(239):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_where)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(240):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_where)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(241):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(242):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(250)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(122)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(125)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(243):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(250)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(23)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(244):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('b') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(122)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(125)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(245):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') || lookahead == int32('_') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(246):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(23)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') || lookahead == int32('_') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(247):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(23)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(248):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(23)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(248)
			goto next_state
		}
		return result
	case int32(249):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(250):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(22)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || lookahead == int32('_') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(251):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0x)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(252):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0X)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(253):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(250)
			goto next_state
		}
		if lookahead == int32('X') {
			state = uint16(252)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(251)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(254)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(254):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(250)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(17)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(254)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(255):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(265)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('b') <= lookahead && lookahead <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(256):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(257)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(158)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(274)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(257):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(264)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(258):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(259)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(259):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(255)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('e') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(260):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(216)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(335)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(261):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(357)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(262):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(327)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(297)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(263):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(318)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(264):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(291)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(265):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(316)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(266):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(361)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(267):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_address_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(124)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(268):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('1') {
			state = uint16(271)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(272)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(306)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(289)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(269):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('1') {
			state = uint16(271)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(272)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(289)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(270):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('1') {
			state = uint16(271)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(272)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(148)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(271):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(273)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(272):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('4') {
			state = uint16(150)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(273):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('8') {
			state = uint16(152)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(274):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(310)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(275):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(312)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(276):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(302)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(277):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(348)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(285)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(278):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(355)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(359)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(279):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(311)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(280):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('b') {
			state = uint16(226)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(281):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(208)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(282):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(228)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(283):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(304)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(284):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(352)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(285):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(224)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(286):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(288)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(160)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(287):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(288)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(288):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(336)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(289):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(146)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(290):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(234)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(291):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(342)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(292):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(240)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(293):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(189)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(294):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(238)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(295):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(197)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(296):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(347)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(330)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(297):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(275)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(298):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(338)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(315)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(299):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(354)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(300):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(185)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(333)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(301):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(185)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(302):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(294)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(303):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('h') {
			state = uint16(298)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(304):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('h') {
			state = uint16(222)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(305):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('h') {
			state = uint16(307)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(306):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(329)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(276)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(307):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(315)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(308):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(323)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(309):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(282)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(310):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(349)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(311):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(351)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(312):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(195)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(313):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(154)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(314):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(220)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(315):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(293)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(316):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(353)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(317):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('m') {
			state = uint16(214)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(318):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(344)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(319):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(236)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(320):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(183)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(321):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(281)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(322):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(345)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(323):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(360)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(324):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(327)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(297)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(325):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(327)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(326):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(318)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(327):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(313)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(328):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(330)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(329):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(319)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(330):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(332)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(331):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(322)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(332):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(191)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(333):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(314)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(334):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(290)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(335):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(218)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(336):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(291)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(337):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(320)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(338):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(292)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(339):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(334)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(340):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(156)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(341):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(160)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(342):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(340)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(343):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(289)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(344):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(350)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(308)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(345):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(350)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(346):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(278)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(347):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(187)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(348):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(283)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(349):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(210)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(350):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(193)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(351):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(232)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(352):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(230)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(353):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(212)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(354):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(358)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(355):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(309)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(356):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(280)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(357):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(317)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(358):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(337)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(359):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(284)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(360):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(295)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(361):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(266)
			goto next_state
		}
		if int32('G') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(362):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(363):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(364):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\\') {
			state = uint16(126)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(364)
			goto next_state
		}
		return result
	case int32(365):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\\') {
			state = uint16(364)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(126)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

func ts_lex_keywords(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
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
		if lookahead == int32('_') {
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym__)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(2):
		if lookahead == int32('l') {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('a') {
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('o') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('u') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('r') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('s') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('l') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('n') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('d') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('t') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('b') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('u') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('e') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('s') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(16):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_fun)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(17):
		if lookahead == int32('u') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(18):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_mutable_specifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(19):
		if lookahead == int32('l') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('e') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_else)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(22):
		if lookahead == int32('e') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('l') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('i') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		if lookahead == int32('e') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('c') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_module)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_public)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [357]TSLexMode{
	0: {},
	1: {
		Flex_state: uint16(12),
	},
	2: {
		Flex_state: uint16(1),
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
		Flex_state: uint16(4),
	},
	10: {
		Flex_state: uint16(3),
	},
	11: {
		Flex_state: uint16(5),
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
		Flex_state: uint16(3),
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
		Flex_state: uint16(3),
	},
	25: {
		Flex_state: uint16(3),
	},
	26: {
		Flex_state: uint16(3),
	},
	27: {
		Flex_state: uint16(3),
	},
	28: {
		Flex_state: uint16(3),
	},
	29: {
		Flex_state: uint16(3),
	},
	30: {
		Flex_state: uint16(3),
	},
	31: {
		Flex_state: uint16(3),
	},
	32: {
		Flex_state: uint16(3),
	},
	33: {
		Flex_state: uint16(3),
	},
	34: {
		Flex_state: uint16(3),
	},
	35: {
		Flex_state: uint16(3),
	},
	36: {
		Flex_state: uint16(3),
	},
	37: {
		Flex_state: uint16(3),
	},
	38: {
		Flex_state: uint16(3),
	},
	39: {
		Flex_state: uint16(3),
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
		Flex_state: uint16(9),
	},
	50: {
		Flex_state: uint16(9),
	},
	51: {
		Flex_state: uint16(9),
	},
	52: {
		Flex_state: uint16(9),
	},
	53: {
		Flex_state: uint16(9),
	},
	54: {
		Flex_state: uint16(9),
	},
	55: {
		Flex_state: uint16(9),
	},
	56: {
		Flex_state: uint16(9),
	},
	57: {
		Flex_state: uint16(9),
	},
	58: {
		Flex_state: uint16(9),
	},
	59: {
		Flex_state: uint16(9),
	},
	60: {
		Flex_state: uint16(9),
	},
	61: {
		Flex_state: uint16(9),
	},
	62: {
		Flex_state: uint16(9),
	},
	63: {
		Flex_state: uint16(9),
	},
	64: {
		Flex_state: uint16(9),
	},
	65: {
		Flex_state: uint16(10),
	},
	66: {
		Flex_state: uint16(10),
	},
	67: {
		Flex_state: uint16(10),
	},
	68: {
		Flex_state: uint16(10),
	},
	69: {
		Flex_state: uint16(10),
	},
	70: {
		Flex_state: uint16(10),
	},
	71: {
		Flex_state: uint16(10),
	},
	72: {
		Flex_state: uint16(10),
	},
	73: {
		Flex_state: uint16(10),
	},
	74: {
		Flex_state: uint16(10),
	},
	75: {
		Flex_state: uint16(10),
	},
	76: {
		Flex_state: uint16(7),
	},
	77: {
		Flex_state: uint16(7),
	},
	78: {
		Flex_state: uint16(4),
	},
	79: {
		Flex_state: uint16(4),
	},
	80: {
		Flex_state: uint16(8),
	},
	81: {
		Flex_state: uint16(4),
	},
	82: {
		Flex_state: uint16(6),
	},
	83: {
		Flex_state: uint16(6),
	},
	84: {
		Flex_state: uint16(6),
	},
	85: {
		Flex_state: uint16(6),
	},
	86: {
		Flex_state: uint16(6),
	},
	87: {
		Flex_state: uint16(4),
	},
	88: {
		Flex_state: uint16(4),
	},
	89: {
		Flex_state: uint16(4),
	},
	90: {
		Flex_state: uint16(6),
	},
	91: {
		Flex_state: uint16(4),
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
		Flex_state: uint16(4),
	},
	101: {
		Flex_state: uint16(6),
	},
	102: {
		Flex_state: uint16(6),
	},
	103: {
		Flex_state: uint16(6),
	},
	104: {
		Flex_state: uint16(6),
	},
	105: {
		Flex_state: uint16(6),
	},
	106: {
		Flex_state: uint16(6),
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
		Flex_state: uint16(6),
	},
	112: {
		Flex_state: uint16(6),
	},
	113: {
		Flex_state: uint16(6),
	},
	114: {
		Flex_state: uint16(6),
	},
	115: {
		Flex_state: uint16(6),
	},
	116: {
		Flex_state: uint16(6),
	},
	117: {
		Flex_state: uint16(6),
	},
	118: {
		Flex_state: uint16(6),
	},
	119: {
		Flex_state: uint16(6),
	},
	120: {
		Flex_state: uint16(6),
	},
	121: {
		Flex_state: uint16(6),
	},
	122: {
		Flex_state: uint16(6),
	},
	123: {
		Flex_state: uint16(6),
	},
	124: {
		Flex_state: uint16(6),
	},
	125: {
		Flex_state: uint16(6),
	},
	126: {
		Flex_state: uint16(6),
	},
	127: {
		Flex_state: uint16(6),
	},
	128: {
		Flex_state: uint16(6),
	},
	129: {
		Flex_state: uint16(6),
	},
	130: {
		Flex_state: uint16(6),
	},
	131: {
		Flex_state: uint16(6),
	},
	132: {
		Flex_state: uint16(6),
	},
	133: {
		Flex_state: uint16(6),
	},
	134: {
		Flex_state: uint16(6),
	},
	135: {
		Flex_state: uint16(6),
	},
	136: {
		Flex_state: uint16(6),
	},
	137: {
		Flex_state: uint16(6),
	},
	138: {
		Flex_state: uint16(6),
	},
	139: {
		Flex_state: uint16(3),
	},
	140: {
		Flex_state: uint16(3),
	},
	141: {
		Flex_state: uint16(6),
	},
	142: {
		Flex_state: uint16(3),
	},
	143: {
		Flex_state: uint16(6),
	},
	144: {
		Flex_state: uint16(6),
	},
	145: {
		Flex_state: uint16(6),
	},
	146: {
		Flex_state: uint16(6),
	},
	147: {
		Flex_state: uint16(6),
	},
	148: {
		Flex_state: uint16(6),
	},
	149: {
		Flex_state: uint16(6),
	},
	150: {
		Flex_state: uint16(6),
	},
	151: {
		Flex_state: uint16(6),
	},
	152: {
		Flex_state: uint16(6),
	},
	153: {
		Flex_state: uint16(6),
	},
	154: {
		Flex_state: uint16(11),
	},
	155: {
		Flex_state: uint16(11),
	},
	156: {
		Flex_state: uint16(11),
	},
	157: {
		Flex_state: uint16(11),
	},
	158: {
		Flex_state: uint16(11),
	},
	159: {
		Flex_state: uint16(11),
	},
	160: {
		Flex_state: uint16(11),
	},
	161: {
		Flex_state: uint16(11),
	},
	162: {
		Flex_state: uint16(11),
	},
	163: {
		Flex_state: uint16(11),
	},
	164: {
		Flex_state: uint16(11),
	},
	165: {
		Flex_state: uint16(11),
	},
	166: {
		Flex_state: uint16(11),
	},
	167: {
		Flex_state: uint16(11),
	},
	168: {
		Flex_state: uint16(11),
	},
	169: {
		Flex_state: uint16(11),
	},
	170: {
		Flex_state: uint16(11),
	},
	171: {
		Flex_state: uint16(11),
	},
	172: {
		Flex_state: uint16(11),
	},
	173: {
		Flex_state: uint16(11),
	},
	174: {
		Flex_state: uint16(11),
	},
	175: {
		Flex_state: uint16(11),
	},
	176: {
		Flex_state: uint16(11),
	},
	177: {
		Flex_state: uint16(11),
	},
	178: {
		Flex_state: uint16(11),
	},
	179: {
		Flex_state: uint16(20),
	},
	180: {
		Flex_state: uint16(20),
	},
	181: {
		Flex_state: uint16(20),
	},
	182: {
		Flex_state: uint16(12),
	},
	183: {
		Flex_state: uint16(12),
	},
	184: {
		Flex_state: uint16(12),
	},
	185: {
		Flex_state: uint16(12),
	},
	186: {
		Flex_state: uint16(12),
	},
	187: {
		Flex_state: uint16(12),
	},
	188: {
		Flex_state: uint16(12),
	},
	189: {
		Flex_state: uint16(13),
	},
	190: {
		Flex_state: uint16(12),
	},
	191: {},
	192: {},
	193: {},
	194: {
		Flex_state: uint16(12),
	},
	195: {},
	196: {},
	197: {},
	198: {},
	199: {
		Flex_state: uint16(12),
	},
	200: {},
	201: {},
	202: {},
	203: {
		Flex_state: uint16(13),
	},
	204: {},
	205: {
		Flex_state: uint16(13),
	},
	206: {},
	207: {},
	208: {
		Flex_state: uint16(13),
	},
	209: {
		Flex_state: uint16(20),
	},
	210: {
		Flex_state: uint16(20),
	},
	211: {
		Flex_state: uint16(4),
	},
	212: {
		Flex_state: uint16(4),
	},
	213: {
		Flex_state: uint16(6),
	},
	214: {
		Flex_state: uint16(4),
	},
	215: {
		Flex_state: uint16(20),
	},
	216: {
		Flex_state: uint16(4),
	},
	217: {
		Flex_state: uint16(20),
	},
	218: {
		Flex_state: uint16(20),
	},
	219: {
		Flex_state: uint16(20),
	},
	220: {
		Flex_state: uint16(12),
	},
	221: {},
	222: {
		Flex_state: uint16(20),
	},
	223: {
		Flex_state: uint16(6),
	},
	224: {
		Flex_state: uint16(20),
	},
	225: {
		Flex_state: uint16(20),
	},
	226: {},
	227: {},
	228: {
		Flex_state: uint16(12),
	},
	229: {},
	230: {
		Flex_state: uint16(12),
	},
	231: {},
	232: {},
	233: {},
	234: {},
	235: {},
	236: {},
	237: {
		Flex_state: uint16(12),
	},
	238: {},
	239: {},
	240: {},
	241: {},
	242: {},
	243: {
		Flex_state: uint16(12),
	},
	244: {},
	245: {
		Flex_state: uint16(12),
	},
	246: {
		Flex_state: uint16(12),
	},
	247: {},
	248: {},
	249: {},
	250: {},
	251: {
		Flex_state: uint16(12),
	},
	252: {
		Flex_state: uint16(12),
	},
	253: {
		Flex_state: uint16(19),
	},
	254: {},
	255: {},
	256: {
		Flex_state: uint16(12),
	},
	257: {},
	258: {
		Flex_state: uint16(12),
	},
	259: {},
	260: {},
	261: {
		Flex_state: uint16(12),
	},
	262: {
		Flex_state: uint16(13),
	},
	263: {
		Flex_state: uint16(12),
	},
	264: {
		Flex_state: uint16(12),
	},
	265: {},
	266: {},
	267: {},
	268: {},
	269: {},
	270: {},
	271: {},
	272: {
		Flex_state: uint16(12),
	},
	273: {},
	274: {},
	275: {},
	276: {},
	277: {},
	278: {},
	279: {},
	280: {},
	281: {
		Flex_state: uint16(12),
	},
	282: {},
	283: {},
	284: {},
	285: {
		Flex_state: uint16(12),
	},
	286: {},
	287: {},
	288: {},
	289: {
		Flex_state: uint16(12),
	},
	290: {},
	291: {},
	292: {
		Flex_state: uint16(12),
	},
	293: {},
	294: {
		Flex_state: uint16(12),
	},
	295: {
		Flex_state: uint16(12),
	},
	296: {
		Flex_state: uint16(12),
	},
	297: {},
	298: {},
	299: {},
	300: {},
	301: {
		Flex_state: uint16(1),
	},
	302: {
		Flex_state: uint16(12),
	},
	303: {
		Flex_state: uint16(12),
	},
	304: {},
	305: {},
	306: {},
	307: {
		Flex_state: uint16(12),
	},
	308: {},
	309: {},
	310: {},
	311: {
		Flex_state: uint16(12),
	},
	312: {},
	313: {},
	314: {},
	315: {},
	316: {},
	317: {
		Flex_state: uint16(12),
	},
	318: {
		Flex_state: uint16(12),
	},
	319: {},
	320: {},
	321: {},
	322: {
		Flex_state: uint16(12),
	},
	323: {},
	324: {},
	325: {},
	326: {},
	327: {},
	328: {},
	329: {
		Flex_state: uint16(12),
	},
	330: {},
	331: {
		Flex_state: uint16(12),
	},
	332: {
		Flex_state: uint16(12),
	},
	333: {
		Flex_state: uint16(12),
	},
	334: {
		Flex_state: uint16(12),
	},
	335: {
		Flex_state: uint16(12),
	},
	336: {
		Flex_state: uint16(12),
	},
	337: {},
	338: {},
	339: {},
	340: {
		Flex_state: uint16(12),
	},
	341: {
		Flex_state: uint16(12),
	},
	342: {},
	343: {},
	344: {},
	345: {
		Flex_state: uint16(12),
	},
	346: {
		Flex_state: uint16(12),
	},
	347: {
		Flex_state: uint16(12),
	},
	348: {},
	349: {},
	350: {
		Flex_state: uint16(12),
	},
	351: {
		Flex_state: uint16(12),
	},
	352: {
		Flex_state: uint16(21),
	},
	353: {},
	354: {
		Flex_state: uint16(12),
	},
	355: {},
	356: {},
}

var ts_parse_table = [3][163]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
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
		52: uint16(1),
		53: uint16(1),
		54: uint16(1),
		55: uint16(1),
		56: uint16(1),
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
		77: uint16(1),
		78: uint16(1),
		79: uint16(1),
		80: uint16(1),
		81: uint16(1),
		82: uint16(1),
		83: uint16(3),
	},
	1: {
		2:  uint16(5),
		83: uint16(3),
		84: uint16(353),
	},
	2: {
		1:   uint16(7),
		3:   uint16(9),
		4:   uint16(11),
		6:   uint16(13),
		9:   uint16(13),
		10:  uint16(15),
		14:  uint16(17),
		15:  uint16(13),
		16:  uint16(15),
		17:  uint16(19),
		19:  uint16(13),
		21:  uint16(21),
		22:  uint16(21),
		23:  uint16(21),
		24:  uint16(21),
		25:  uint16(21),
		26:  uint16(15),
		28:  uint16(23),
		30:  uint16(25),
		31:  uint16(13),
		32:  uint16(13),
		33:  uint16(15),
		34:  uint16(13),
		35:  uint16(13),
		36:  uint16(13),
		37:  uint16(13),
		38:  uint16(13),
		39:  uint16(13),
		40:  uint16(13),
		41:  uint16(13),
		42:  uint16(13),
		43:  uint16(13),
		44:  uint16(15),
		45:  uint16(13),
		46:  uint16(27),
		47:  uint16(29),
		50:  uint16(31),
		51:  uint16(33),
		52:  uint16(35),
		53:  uint16(37),
		54:  uint16(39),
		55:  uint16(15),
		76:  uint16(41),
		77:  uint16(43),
		78:  uint16(45),
		79:  uint16(45),
		83:  uint16(3),
		99:  uint16(329),
		102: uint16(83),
		106: uint16(83),
		107: uint16(107),
		108: uint16(83),
		110: uint16(83),
		111: uint16(83),
		112: uint16(83),
		113: uint16(83),
		114: uint16(83),
		115: uint16(83),
		116: uint16(83),
		118: uint16(83),
		119: uint16(83),
		120: uint16(83),
		121: uint16(83),
		125: uint16(83),
		126: uint16(83),
		127: uint16(83),
		128: uint16(83),
		129: uint16(83),
		130: uint16(82),
		131: uint16(83),
		142: uint16(340),
		144: uint16(83),
		147: uint16(83),
		149: uint16(329),
	},
}

var ts_small_parse_table = [12065]uint16_t{
	0:     uint16(31),
	1:     uint16(3),
	2:     uint16(1),
	3:     uint16(sym_comment),
	4:     uint16(47),
	5:     uint16(1),
	6:     uint16(sym_identifier),
	7:     uint16(50),
	8:     uint16(1),
	9:     uint16(anon_sym_COLON_COLON),
	10:    uint16(53),
	11:    uint16(1),
	12:    uint16(anon_sym_LBRACE),
	13:    uint16(56),
	14:    uint16(1),
	15:    uint16(anon_sym_RBRACE),
	16:    uint16(58),
	17:    uint16(1),
	18:    uint16(anon_sym_fun),
	19:    uint16(61),
	20:    uint16(1),
	21:    uint16(anon_sym_public),
	22:    uint16(64),
	23:    uint16(1),
	24:    uint16(anon_sym_LT),
	25:    uint16(67),
	26:    uint16(1),
	27:    uint16(anon_sym_LPAREN),
	28:    uint16(70),
	29:    uint16(1),
	30:    uint16(anon_sym_use),
	31:    uint16(76),
	32:    uint16(1),
	33:    uint16(anon_sym_AMP),
	34:    uint16(79),
	35:    uint16(1),
	36:    uint16(anon_sym_BANG),
	37:    uint16(82),
	38:    uint16(1),
	39:    uint16(anon_sym_return),
	40:    uint16(85),
	41:    uint16(1),
	42:    uint16(anon_sym_if),
	43:    uint16(88),
	44:    uint16(1),
	45:    uint16(anon_sym_while),
	46:    uint16(91),
	47:    uint16(1),
	48:    uint16(anon_sym_loop),
	49:    uint16(94),
	50:    uint16(1),
	51:    uint16(anon_sym_const),
	52:    uint16(97),
	53:    uint16(1),
	54:    uint16(anon_sym_break),
	55:    uint16(100),
	56:    uint16(1),
	57:    uint16(anon_sym_continue),
	58:    uint16(103),
	59:    uint16(1),
	60:    uint16(sym_integer_literal),
	61:    uint16(106),
	62:    uint16(1),
	63:    uint16(sym_float_literal),
	64:    uint16(82),
	65:    uint16(1),
	66:    uint16(sym_field_expression),
	67:    uint16(153),
	68:    uint16(1),
	69:    uint16(sym__expression),
	70:    uint16(340),
	71:    uint16(1),
	72:    uint16(sym_scoped_identifier),
	73:    uint16(354),
	74:    uint16(1),
	75:    uint16(sym_visibility_modifier),
	76:    uint16(109),
	77:    uint16(2),
	78:    uint16(anon_sym_true),
	79:    uint16(anon_sym_false),
	80:    uint16(329),
	81:    uint16(2),
	82:    uint16(sym_bracketed_type),
	83:    uint16(sym__path),
	84:    uint16(73),
	85:    uint16(5),
	86:    uint16(anon_sym_u8),
	87:    uint16(anon_sym_u64),
	88:    uint16(anon_sym_u128),
	89:    uint16(anon_sym_bool),
	90:    uint16(anon_sym_address),
	91:    uint16(3),
	92:    uint16(6),
	93:    uint16(sym__statement),
	94:    uint16(sym_expression_statement),
	95:    uint16(sym__declaration_statement),
	96:    uint16(sym_function_item),
	97:    uint16(sym_use_declaration),
	98:    uint16(aux_sym_block_repeat1),
	99:    uint16(44),
	100:   uint16(6),
	101:   uint16(sym__expression_ending_with_block),
	102:   uint16(sym_if_expression),
	103:   uint16(sym_while_expression),
	104:   uint16(sym_loop_expression),
	105:   uint16(sym_const_block),
	106:   uint16(sym_block),
	107:   uint16(83),
	108:   uint16(16),
	109:   uint16(sym_generic_function),
	110:   uint16(sym__expression_except_range),
	111:   uint16(sym_unary_expression),
	112:   uint16(sym_reference_expression),
	113:   uint16(sym_binary_expression),
	114:   uint16(sym_assignment_expression),
	115:   uint16(sym_type_cast_expression),
	116:   uint16(sym_return_expression),
	117:   uint16(sym_call_expression),
	118:   uint16(sym_parenthesized_expression),
	119:   uint16(sym_tuple_expression),
	120:   uint16(sym_unit_expression),
	121:   uint16(sym_break_expression),
	122:   uint16(sym_continue_expression),
	123:   uint16(sym__literal),
	124:   uint16(sym_boolean_literal),
	125:   uint16(31),
	126:   uint16(3),
	127:   uint16(1),
	128:   uint16(sym_comment),
	129:   uint16(7),
	130:   uint16(1),
	131:   uint16(sym_identifier),
	132:   uint16(9),
	133:   uint16(1),
	134:   uint16(anon_sym_COLON_COLON),
	135:   uint16(19),
	136:   uint16(1),
	137:   uint16(anon_sym_LPAREN),
	138:   uint16(27),
	139:   uint16(1),
	140:   uint16(anon_sym_return),
	141:   uint16(37),
	142:   uint16(1),
	143:   uint16(anon_sym_break),
	144:   uint16(39),
	145:   uint16(1),
	146:   uint16(anon_sym_continue),
	147:   uint16(41),
	148:   uint16(1),
	149:   uint16(sym_integer_literal),
	150:   uint16(43),
	151:   uint16(1),
	152:   uint16(sym_float_literal),
	153:   uint16(112),
	154:   uint16(1),
	155:   uint16(anon_sym_LBRACE),
	156:   uint16(114),
	157:   uint16(1),
	158:   uint16(anon_sym_RBRACE),
	159:   uint16(116),
	160:   uint16(1),
	161:   uint16(anon_sym_fun),
	162:   uint16(118),
	163:   uint16(1),
	164:   uint16(anon_sym_public),
	165:   uint16(120),
	166:   uint16(1),
	167:   uint16(anon_sym_LT),
	168:   uint16(122),
	169:   uint16(1),
	170:   uint16(anon_sym_use),
	171:   uint16(124),
	172:   uint16(1),
	173:   uint16(anon_sym_AMP),
	174:   uint16(126),
	175:   uint16(1),
	176:   uint16(anon_sym_BANG),
	177:   uint16(128),
	178:   uint16(1),
	179:   uint16(anon_sym_if),
	180:   uint16(130),
	181:   uint16(1),
	182:   uint16(anon_sym_while),
	183:   uint16(132),
	184:   uint16(1),
	185:   uint16(anon_sym_loop),
	186:   uint16(134),
	187:   uint16(1),
	188:   uint16(anon_sym_const),
	189:   uint16(82),
	190:   uint16(1),
	191:   uint16(sym_field_expression),
	192:   uint16(153),
	193:   uint16(1),
	194:   uint16(sym__expression),
	195:   uint16(340),
	196:   uint16(1),
	197:   uint16(sym_scoped_identifier),
	198:   uint16(354),
	199:   uint16(1),
	200:   uint16(sym_visibility_modifier),
	201:   uint16(45),
	202:   uint16(2),
	203:   uint16(anon_sym_true),
	204:   uint16(anon_sym_false),
	205:   uint16(329),
	206:   uint16(2),
	207:   uint16(sym_bracketed_type),
	208:   uint16(sym__path),
	209:   uint16(21),
	210:   uint16(5),
	211:   uint16(anon_sym_u8),
	212:   uint16(anon_sym_u64),
	213:   uint16(anon_sym_u128),
	214:   uint16(anon_sym_bool),
	215:   uint16(anon_sym_address),
	216:   uint16(9),
	217:   uint16(6),
	218:   uint16(sym__statement),
	219:   uint16(sym_expression_statement),
	220:   uint16(sym__declaration_statement),
	221:   uint16(sym_function_item),
	222:   uint16(sym_use_declaration),
	223:   uint16(aux_sym_block_repeat1),
	224:   uint16(44),
	225:   uint16(6),
	226:   uint16(sym__expression_ending_with_block),
	227:   uint16(sym_if_expression),
	228:   uint16(sym_while_expression),
	229:   uint16(sym_loop_expression),
	230:   uint16(sym_const_block),
	231:   uint16(sym_block),
	232:   uint16(83),
	233:   uint16(16),
	234:   uint16(sym_generic_function),
	235:   uint16(sym__expression_except_range),
	236:   uint16(sym_unary_expression),
	237:   uint16(sym_reference_expression),
	238:   uint16(sym_binary_expression),
	239:   uint16(sym_assignment_expression),
	240:   uint16(sym_type_cast_expression),
	241:   uint16(sym_return_expression),
	242:   uint16(sym_call_expression),
	243:   uint16(sym_parenthesized_expression),
	244:   uint16(sym_tuple_expression),
	245:   uint16(sym_unit_expression),
	246:   uint16(sym_break_expression),
	247:   uint16(sym_continue_expression),
	248:   uint16(sym__literal),
	249:   uint16(sym_boolean_literal),
	250:   uint16(31),
	251:   uint16(3),
	252:   uint16(1),
	253:   uint16(sym_comment),
	254:   uint16(7),
	255:   uint16(1),
	256:   uint16(sym_identifier),
	257:   uint16(9),
	258:   uint16(1),
	259:   uint16(anon_sym_COLON_COLON),
	260:   uint16(19),
	261:   uint16(1),
	262:   uint16(anon_sym_LPAREN),
	263:   uint16(27),
	264:   uint16(1),
	265:   uint16(anon_sym_return),
	266:   uint16(37),
	267:   uint16(1),
	268:   uint16(anon_sym_break),
	269:   uint16(39),
	270:   uint16(1),
	271:   uint16(anon_sym_continue),
	272:   uint16(41),
	273:   uint16(1),
	274:   uint16(sym_integer_literal),
	275:   uint16(43),
	276:   uint16(1),
	277:   uint16(sym_float_literal),
	278:   uint16(112),
	279:   uint16(1),
	280:   uint16(anon_sym_LBRACE),
	281:   uint16(116),
	282:   uint16(1),
	283:   uint16(anon_sym_fun),
	284:   uint16(118),
	285:   uint16(1),
	286:   uint16(anon_sym_public),
	287:   uint16(120),
	288:   uint16(1),
	289:   uint16(anon_sym_LT),
	290:   uint16(122),
	291:   uint16(1),
	292:   uint16(anon_sym_use),
	293:   uint16(124),
	294:   uint16(1),
	295:   uint16(anon_sym_AMP),
	296:   uint16(126),
	297:   uint16(1),
	298:   uint16(anon_sym_BANG),
	299:   uint16(128),
	300:   uint16(1),
	301:   uint16(anon_sym_if),
	302:   uint16(130),
	303:   uint16(1),
	304:   uint16(anon_sym_while),
	305:   uint16(132),
	306:   uint16(1),
	307:   uint16(anon_sym_loop),
	308:   uint16(134),
	309:   uint16(1),
	310:   uint16(anon_sym_const),
	311:   uint16(136),
	312:   uint16(1),
	313:   uint16(anon_sym_RBRACE),
	314:   uint16(82),
	315:   uint16(1),
	316:   uint16(sym_field_expression),
	317:   uint16(153),
	318:   uint16(1),
	319:   uint16(sym__expression),
	320:   uint16(340),
	321:   uint16(1),
	322:   uint16(sym_scoped_identifier),
	323:   uint16(354),
	324:   uint16(1),
	325:   uint16(sym_visibility_modifier),
	326:   uint16(45),
	327:   uint16(2),
	328:   uint16(anon_sym_true),
	329:   uint16(anon_sym_false),
	330:   uint16(329),
	331:   uint16(2),
	332:   uint16(sym_bracketed_type),
	333:   uint16(sym__path),
	334:   uint16(21),
	335:   uint16(5),
	336:   uint16(anon_sym_u8),
	337:   uint16(anon_sym_u64),
	338:   uint16(anon_sym_u128),
	339:   uint16(anon_sym_bool),
	340:   uint16(anon_sym_address),
	341:   uint16(6),
	342:   uint16(6),
	343:   uint16(sym__statement),
	344:   uint16(sym_expression_statement),
	345:   uint16(sym__declaration_statement),
	346:   uint16(sym_function_item),
	347:   uint16(sym_use_declaration),
	348:   uint16(aux_sym_block_repeat1),
	349:   uint16(44),
	350:   uint16(6),
	351:   uint16(sym__expression_ending_with_block),
	352:   uint16(sym_if_expression),
	353:   uint16(sym_while_expression),
	354:   uint16(sym_loop_expression),
	355:   uint16(sym_const_block),
	356:   uint16(sym_block),
	357:   uint16(83),
	358:   uint16(16),
	359:   uint16(sym_generic_function),
	360:   uint16(sym__expression_except_range),
	361:   uint16(sym_unary_expression),
	362:   uint16(sym_reference_expression),
	363:   uint16(sym_binary_expression),
	364:   uint16(sym_assignment_expression),
	365:   uint16(sym_type_cast_expression),
	366:   uint16(sym_return_expression),
	367:   uint16(sym_call_expression),
	368:   uint16(sym_parenthesized_expression),
	369:   uint16(sym_tuple_expression),
	370:   uint16(sym_unit_expression),
	371:   uint16(sym_break_expression),
	372:   uint16(sym_continue_expression),
	373:   uint16(sym__literal),
	374:   uint16(sym_boolean_literal),
	375:   uint16(31),
	376:   uint16(3),
	377:   uint16(1),
	378:   uint16(sym_comment),
	379:   uint16(7),
	380:   uint16(1),
	381:   uint16(sym_identifier),
	382:   uint16(9),
	383:   uint16(1),
	384:   uint16(anon_sym_COLON_COLON),
	385:   uint16(19),
	386:   uint16(1),
	387:   uint16(anon_sym_LPAREN),
	388:   uint16(27),
	389:   uint16(1),
	390:   uint16(anon_sym_return),
	391:   uint16(37),
	392:   uint16(1),
	393:   uint16(anon_sym_break),
	394:   uint16(39),
	395:   uint16(1),
	396:   uint16(anon_sym_continue),
	397:   uint16(41),
	398:   uint16(1),
	399:   uint16(sym_integer_literal),
	400:   uint16(43),
	401:   uint16(1),
	402:   uint16(sym_float_literal),
	403:   uint16(112),
	404:   uint16(1),
	405:   uint16(anon_sym_LBRACE),
	406:   uint16(116),
	407:   uint16(1),
	408:   uint16(anon_sym_fun),
	409:   uint16(118),
	410:   uint16(1),
	411:   uint16(anon_sym_public),
	412:   uint16(120),
	413:   uint16(1),
	414:   uint16(anon_sym_LT),
	415:   uint16(122),
	416:   uint16(1),
	417:   uint16(anon_sym_use),
	418:   uint16(124),
	419:   uint16(1),
	420:   uint16(anon_sym_AMP),
	421:   uint16(126),
	422:   uint16(1),
	423:   uint16(anon_sym_BANG),
	424:   uint16(128),
	425:   uint16(1),
	426:   uint16(anon_sym_if),
	427:   uint16(130),
	428:   uint16(1),
	429:   uint16(anon_sym_while),
	430:   uint16(132),
	431:   uint16(1),
	432:   uint16(anon_sym_loop),
	433:   uint16(134),
	434:   uint16(1),
	435:   uint16(anon_sym_const),
	436:   uint16(138),
	437:   uint16(1),
	438:   uint16(anon_sym_RBRACE),
	439:   uint16(82),
	440:   uint16(1),
	441:   uint16(sym_field_expression),
	442:   uint16(153),
	443:   uint16(1),
	444:   uint16(sym__expression),
	445:   uint16(340),
	446:   uint16(1),
	447:   uint16(sym_scoped_identifier),
	448:   uint16(354),
	449:   uint16(1),
	450:   uint16(sym_visibility_modifier),
	451:   uint16(45),
	452:   uint16(2),
	453:   uint16(anon_sym_true),
	454:   uint16(anon_sym_false),
	455:   uint16(329),
	456:   uint16(2),
	457:   uint16(sym_bracketed_type),
	458:   uint16(sym__path),
	459:   uint16(21),
	460:   uint16(5),
	461:   uint16(anon_sym_u8),
	462:   uint16(anon_sym_u64),
	463:   uint16(anon_sym_u128),
	464:   uint16(anon_sym_bool),
	465:   uint16(anon_sym_address),
	466:   uint16(3),
	467:   uint16(6),
	468:   uint16(sym__statement),
	469:   uint16(sym_expression_statement),
	470:   uint16(sym__declaration_statement),
	471:   uint16(sym_function_item),
	472:   uint16(sym_use_declaration),
	473:   uint16(aux_sym_block_repeat1),
	474:   uint16(44),
	475:   uint16(6),
	476:   uint16(sym__expression_ending_with_block),
	477:   uint16(sym_if_expression),
	478:   uint16(sym_while_expression),
	479:   uint16(sym_loop_expression),
	480:   uint16(sym_const_block),
	481:   uint16(sym_block),
	482:   uint16(83),
	483:   uint16(16),
	484:   uint16(sym_generic_function),
	485:   uint16(sym__expression_except_range),
	486:   uint16(sym_unary_expression),
	487:   uint16(sym_reference_expression),
	488:   uint16(sym_binary_expression),
	489:   uint16(sym_assignment_expression),
	490:   uint16(sym_type_cast_expression),
	491:   uint16(sym_return_expression),
	492:   uint16(sym_call_expression),
	493:   uint16(sym_parenthesized_expression),
	494:   uint16(sym_tuple_expression),
	495:   uint16(sym_unit_expression),
	496:   uint16(sym_break_expression),
	497:   uint16(sym_continue_expression),
	498:   uint16(sym__literal),
	499:   uint16(sym_boolean_literal),
	500:   uint16(31),
	501:   uint16(3),
	502:   uint16(1),
	503:   uint16(sym_comment),
	504:   uint16(7),
	505:   uint16(1),
	506:   uint16(sym_identifier),
	507:   uint16(9),
	508:   uint16(1),
	509:   uint16(anon_sym_COLON_COLON),
	510:   uint16(19),
	511:   uint16(1),
	512:   uint16(anon_sym_LPAREN),
	513:   uint16(27),
	514:   uint16(1),
	515:   uint16(anon_sym_return),
	516:   uint16(37),
	517:   uint16(1),
	518:   uint16(anon_sym_break),
	519:   uint16(39),
	520:   uint16(1),
	521:   uint16(anon_sym_continue),
	522:   uint16(41),
	523:   uint16(1),
	524:   uint16(sym_integer_literal),
	525:   uint16(43),
	526:   uint16(1),
	527:   uint16(sym_float_literal),
	528:   uint16(112),
	529:   uint16(1),
	530:   uint16(anon_sym_LBRACE),
	531:   uint16(116),
	532:   uint16(1),
	533:   uint16(anon_sym_fun),
	534:   uint16(118),
	535:   uint16(1),
	536:   uint16(anon_sym_public),
	537:   uint16(120),
	538:   uint16(1),
	539:   uint16(anon_sym_LT),
	540:   uint16(122),
	541:   uint16(1),
	542:   uint16(anon_sym_use),
	543:   uint16(124),
	544:   uint16(1),
	545:   uint16(anon_sym_AMP),
	546:   uint16(126),
	547:   uint16(1),
	548:   uint16(anon_sym_BANG),
	549:   uint16(128),
	550:   uint16(1),
	551:   uint16(anon_sym_if),
	552:   uint16(130),
	553:   uint16(1),
	554:   uint16(anon_sym_while),
	555:   uint16(132),
	556:   uint16(1),
	557:   uint16(anon_sym_loop),
	558:   uint16(134),
	559:   uint16(1),
	560:   uint16(anon_sym_const),
	561:   uint16(140),
	562:   uint16(1),
	563:   uint16(anon_sym_RBRACE),
	564:   uint16(82),
	565:   uint16(1),
	566:   uint16(sym_field_expression),
	567:   uint16(153),
	568:   uint16(1),
	569:   uint16(sym__expression),
	570:   uint16(340),
	571:   uint16(1),
	572:   uint16(sym_scoped_identifier),
	573:   uint16(354),
	574:   uint16(1),
	575:   uint16(sym_visibility_modifier),
	576:   uint16(45),
	577:   uint16(2),
	578:   uint16(anon_sym_true),
	579:   uint16(anon_sym_false),
	580:   uint16(329),
	581:   uint16(2),
	582:   uint16(sym_bracketed_type),
	583:   uint16(sym__path),
	584:   uint16(21),
	585:   uint16(5),
	586:   uint16(anon_sym_u8),
	587:   uint16(anon_sym_u64),
	588:   uint16(anon_sym_u128),
	589:   uint16(anon_sym_bool),
	590:   uint16(anon_sym_address),
	591:   uint16(3),
	592:   uint16(6),
	593:   uint16(sym__statement),
	594:   uint16(sym_expression_statement),
	595:   uint16(sym__declaration_statement),
	596:   uint16(sym_function_item),
	597:   uint16(sym_use_declaration),
	598:   uint16(aux_sym_block_repeat1),
	599:   uint16(44),
	600:   uint16(6),
	601:   uint16(sym__expression_ending_with_block),
	602:   uint16(sym_if_expression),
	603:   uint16(sym_while_expression),
	604:   uint16(sym_loop_expression),
	605:   uint16(sym_const_block),
	606:   uint16(sym_block),
	607:   uint16(83),
	608:   uint16(16),
	609:   uint16(sym_generic_function),
	610:   uint16(sym__expression_except_range),
	611:   uint16(sym_unary_expression),
	612:   uint16(sym_reference_expression),
	613:   uint16(sym_binary_expression),
	614:   uint16(sym_assignment_expression),
	615:   uint16(sym_type_cast_expression),
	616:   uint16(sym_return_expression),
	617:   uint16(sym_call_expression),
	618:   uint16(sym_parenthesized_expression),
	619:   uint16(sym_tuple_expression),
	620:   uint16(sym_unit_expression),
	621:   uint16(sym_break_expression),
	622:   uint16(sym_continue_expression),
	623:   uint16(sym__literal),
	624:   uint16(sym_boolean_literal),
	625:   uint16(31),
	626:   uint16(3),
	627:   uint16(1),
	628:   uint16(sym_comment),
	629:   uint16(7),
	630:   uint16(1),
	631:   uint16(sym_identifier),
	632:   uint16(9),
	633:   uint16(1),
	634:   uint16(anon_sym_COLON_COLON),
	635:   uint16(19),
	636:   uint16(1),
	637:   uint16(anon_sym_LPAREN),
	638:   uint16(27),
	639:   uint16(1),
	640:   uint16(anon_sym_return),
	641:   uint16(37),
	642:   uint16(1),
	643:   uint16(anon_sym_break),
	644:   uint16(39),
	645:   uint16(1),
	646:   uint16(anon_sym_continue),
	647:   uint16(41),
	648:   uint16(1),
	649:   uint16(sym_integer_literal),
	650:   uint16(43),
	651:   uint16(1),
	652:   uint16(sym_float_literal),
	653:   uint16(112),
	654:   uint16(1),
	655:   uint16(anon_sym_LBRACE),
	656:   uint16(116),
	657:   uint16(1),
	658:   uint16(anon_sym_fun),
	659:   uint16(118),
	660:   uint16(1),
	661:   uint16(anon_sym_public),
	662:   uint16(120),
	663:   uint16(1),
	664:   uint16(anon_sym_LT),
	665:   uint16(122),
	666:   uint16(1),
	667:   uint16(anon_sym_use),
	668:   uint16(124),
	669:   uint16(1),
	670:   uint16(anon_sym_AMP),
	671:   uint16(126),
	672:   uint16(1),
	673:   uint16(anon_sym_BANG),
	674:   uint16(128),
	675:   uint16(1),
	676:   uint16(anon_sym_if),
	677:   uint16(130),
	678:   uint16(1),
	679:   uint16(anon_sym_while),
	680:   uint16(132),
	681:   uint16(1),
	682:   uint16(anon_sym_loop),
	683:   uint16(134),
	684:   uint16(1),
	685:   uint16(anon_sym_const),
	686:   uint16(142),
	687:   uint16(1),
	688:   uint16(anon_sym_RBRACE),
	689:   uint16(82),
	690:   uint16(1),
	691:   uint16(sym_field_expression),
	692:   uint16(153),
	693:   uint16(1),
	694:   uint16(sym__expression),
	695:   uint16(340),
	696:   uint16(1),
	697:   uint16(sym_scoped_identifier),
	698:   uint16(354),
	699:   uint16(1),
	700:   uint16(sym_visibility_modifier),
	701:   uint16(45),
	702:   uint16(2),
	703:   uint16(anon_sym_true),
	704:   uint16(anon_sym_false),
	705:   uint16(329),
	706:   uint16(2),
	707:   uint16(sym_bracketed_type),
	708:   uint16(sym__path),
	709:   uint16(21),
	710:   uint16(5),
	711:   uint16(anon_sym_u8),
	712:   uint16(anon_sym_u64),
	713:   uint16(anon_sym_u128),
	714:   uint16(anon_sym_bool),
	715:   uint16(anon_sym_address),
	716:   uint16(7),
	717:   uint16(6),
	718:   uint16(sym__statement),
	719:   uint16(sym_expression_statement),
	720:   uint16(sym__declaration_statement),
	721:   uint16(sym_function_item),
	722:   uint16(sym_use_declaration),
	723:   uint16(aux_sym_block_repeat1),
	724:   uint16(44),
	725:   uint16(6),
	726:   uint16(sym__expression_ending_with_block),
	727:   uint16(sym_if_expression),
	728:   uint16(sym_while_expression),
	729:   uint16(sym_loop_expression),
	730:   uint16(sym_const_block),
	731:   uint16(sym_block),
	732:   uint16(83),
	733:   uint16(16),
	734:   uint16(sym_generic_function),
	735:   uint16(sym__expression_except_range),
	736:   uint16(sym_unary_expression),
	737:   uint16(sym_reference_expression),
	738:   uint16(sym_binary_expression),
	739:   uint16(sym_assignment_expression),
	740:   uint16(sym_type_cast_expression),
	741:   uint16(sym_return_expression),
	742:   uint16(sym_call_expression),
	743:   uint16(sym_parenthesized_expression),
	744:   uint16(sym_tuple_expression),
	745:   uint16(sym_unit_expression),
	746:   uint16(sym_break_expression),
	747:   uint16(sym_continue_expression),
	748:   uint16(sym__literal),
	749:   uint16(sym_boolean_literal),
	750:   uint16(31),
	751:   uint16(3),
	752:   uint16(1),
	753:   uint16(sym_comment),
	754:   uint16(7),
	755:   uint16(1),
	756:   uint16(sym_identifier),
	757:   uint16(9),
	758:   uint16(1),
	759:   uint16(anon_sym_COLON_COLON),
	760:   uint16(19),
	761:   uint16(1),
	762:   uint16(anon_sym_LPAREN),
	763:   uint16(27),
	764:   uint16(1),
	765:   uint16(anon_sym_return),
	766:   uint16(37),
	767:   uint16(1),
	768:   uint16(anon_sym_break),
	769:   uint16(39),
	770:   uint16(1),
	771:   uint16(anon_sym_continue),
	772:   uint16(41),
	773:   uint16(1),
	774:   uint16(sym_integer_literal),
	775:   uint16(43),
	776:   uint16(1),
	777:   uint16(sym_float_literal),
	778:   uint16(112),
	779:   uint16(1),
	780:   uint16(anon_sym_LBRACE),
	781:   uint16(116),
	782:   uint16(1),
	783:   uint16(anon_sym_fun),
	784:   uint16(118),
	785:   uint16(1),
	786:   uint16(anon_sym_public),
	787:   uint16(120),
	788:   uint16(1),
	789:   uint16(anon_sym_LT),
	790:   uint16(122),
	791:   uint16(1),
	792:   uint16(anon_sym_use),
	793:   uint16(124),
	794:   uint16(1),
	795:   uint16(anon_sym_AMP),
	796:   uint16(126),
	797:   uint16(1),
	798:   uint16(anon_sym_BANG),
	799:   uint16(128),
	800:   uint16(1),
	801:   uint16(anon_sym_if),
	802:   uint16(130),
	803:   uint16(1),
	804:   uint16(anon_sym_while),
	805:   uint16(132),
	806:   uint16(1),
	807:   uint16(anon_sym_loop),
	808:   uint16(134),
	809:   uint16(1),
	810:   uint16(anon_sym_const),
	811:   uint16(144),
	812:   uint16(1),
	813:   uint16(anon_sym_RBRACE),
	814:   uint16(82),
	815:   uint16(1),
	816:   uint16(sym_field_expression),
	817:   uint16(153),
	818:   uint16(1),
	819:   uint16(sym__expression),
	820:   uint16(340),
	821:   uint16(1),
	822:   uint16(sym_scoped_identifier),
	823:   uint16(354),
	824:   uint16(1),
	825:   uint16(sym_visibility_modifier),
	826:   uint16(45),
	827:   uint16(2),
	828:   uint16(anon_sym_true),
	829:   uint16(anon_sym_false),
	830:   uint16(329),
	831:   uint16(2),
	832:   uint16(sym_bracketed_type),
	833:   uint16(sym__path),
	834:   uint16(21),
	835:   uint16(5),
	836:   uint16(anon_sym_u8),
	837:   uint16(anon_sym_u64),
	838:   uint16(anon_sym_u128),
	839:   uint16(anon_sym_bool),
	840:   uint16(anon_sym_address),
	841:   uint16(3),
	842:   uint16(6),
	843:   uint16(sym__statement),
	844:   uint16(sym_expression_statement),
	845:   uint16(sym__declaration_statement),
	846:   uint16(sym_function_item),
	847:   uint16(sym_use_declaration),
	848:   uint16(aux_sym_block_repeat1),
	849:   uint16(44),
	850:   uint16(6),
	851:   uint16(sym__expression_ending_with_block),
	852:   uint16(sym_if_expression),
	853:   uint16(sym_while_expression),
	854:   uint16(sym_loop_expression),
	855:   uint16(sym_const_block),
	856:   uint16(sym_block),
	857:   uint16(83),
	858:   uint16(16),
	859:   uint16(sym_generic_function),
	860:   uint16(sym__expression_except_range),
	861:   uint16(sym_unary_expression),
	862:   uint16(sym_reference_expression),
	863:   uint16(sym_binary_expression),
	864:   uint16(sym_assignment_expression),
	865:   uint16(sym_type_cast_expression),
	866:   uint16(sym_return_expression),
	867:   uint16(sym_call_expression),
	868:   uint16(sym_parenthesized_expression),
	869:   uint16(sym_tuple_expression),
	870:   uint16(sym_unit_expression),
	871:   uint16(sym_break_expression),
	872:   uint16(sym_continue_expression),
	873:   uint16(sym__literal),
	874:   uint16(sym_boolean_literal),
	875:   uint16(27),
	876:   uint16(3),
	877:   uint16(1),
	878:   uint16(sym_comment),
	879:   uint16(7),
	880:   uint16(1),
	881:   uint16(sym_identifier),
	882:   uint16(9),
	883:   uint16(1),
	884:   uint16(anon_sym_COLON_COLON),
	885:   uint16(11),
	886:   uint16(1),
	887:   uint16(anon_sym_LBRACE),
	888:   uint16(19),
	889:   uint16(1),
	890:   uint16(anon_sym_LPAREN),
	891:   uint16(27),
	892:   uint16(1),
	893:   uint16(anon_sym_return),
	894:   uint16(29),
	895:   uint16(1),
	896:   uint16(anon_sym_if),
	897:   uint16(31),
	898:   uint16(1),
	899:   uint16(anon_sym_while),
	900:   uint16(33),
	901:   uint16(1),
	902:   uint16(anon_sym_loop),
	903:   uint16(35),
	904:   uint16(1),
	905:   uint16(anon_sym_const),
	906:   uint16(37),
	907:   uint16(1),
	908:   uint16(anon_sym_break),
	909:   uint16(39),
	910:   uint16(1),
	911:   uint16(anon_sym_continue),
	912:   uint16(41),
	913:   uint16(1),
	914:   uint16(sym_integer_literal),
	915:   uint16(43),
	916:   uint16(1),
	917:   uint16(sym_float_literal),
	918:   uint16(120),
	919:   uint16(1),
	920:   uint16(anon_sym_LT),
	921:   uint16(124),
	922:   uint16(1),
	923:   uint16(anon_sym_AMP),
	924:   uint16(126),
	925:   uint16(1),
	926:   uint16(anon_sym_BANG),
	927:   uint16(146),
	928:   uint16(1),
	929:   uint16(anon_sym_POUND),
	930:   uint16(148),
	931:   uint16(1),
	932:   uint16(anon_sym_RPAREN),
	933:   uint16(82),
	934:   uint16(1),
	935:   uint16(sym_field_expression),
	936:   uint16(144),
	937:   uint16(1),
	938:   uint16(sym__expression),
	939:   uint16(340),
	940:   uint16(1),
	941:   uint16(sym_scoped_identifier),
	942:   uint16(45),
	943:   uint16(2),
	944:   uint16(anon_sym_true),
	945:   uint16(anon_sym_false),
	946:   uint16(15),
	947:   uint16(2),
	948:   uint16(sym_attribute_item),
	949:   uint16(aux_sym_tuple_expression_repeat1),
	950:   uint16(329),
	951:   uint16(2),
	952:   uint16(sym_bracketed_type),
	953:   uint16(sym__path),
	954:   uint16(21),
	955:   uint16(5),
	956:   uint16(anon_sym_u8),
	957:   uint16(anon_sym_u64),
	958:   uint16(anon_sym_u128),
	959:   uint16(anon_sym_bool),
	960:   uint16(anon_sym_address),
	961:   uint16(83),
	962:   uint16(22),
	963:   uint16(sym_generic_function),
	964:   uint16(sym__expression_except_range),
	965:   uint16(sym__expression_ending_with_block),
	966:   uint16(sym_unary_expression),
	967:   uint16(sym_reference_expression),
	968:   uint16(sym_binary_expression),
	969:   uint16(sym_assignment_expression),
	970:   uint16(sym_type_cast_expression),
	971:   uint16(sym_return_expression),
	972:   uint16(sym_call_expression),
	973:   uint16(sym_parenthesized_expression),
	974:   uint16(sym_tuple_expression),
	975:   uint16(sym_unit_expression),
	976:   uint16(sym_if_expression),
	977:   uint16(sym_while_expression),
	978:   uint16(sym_loop_expression),
	979:   uint16(sym_const_block),
	980:   uint16(sym_break_expression),
	981:   uint16(sym_continue_expression),
	982:   uint16(sym_block),
	983:   uint16(sym__literal),
	984:   uint16(sym_boolean_literal),
	985:   uint16(26),
	986:   uint16(3),
	987:   uint16(1),
	988:   uint16(sym_comment),
	989:   uint16(7),
	990:   uint16(1),
	991:   uint16(sym_identifier),
	992:   uint16(9),
	993:   uint16(1),
	994:   uint16(anon_sym_COLON_COLON),
	995:   uint16(11),
	996:   uint16(1),
	997:   uint16(anon_sym_LBRACE),
	998:   uint16(19),
	999:   uint16(1),
	1000:  uint16(anon_sym_LPAREN),
	1001:  uint16(27),
	1002:  uint16(1),
	1003:  uint16(anon_sym_return),
	1004:  uint16(29),
	1005:  uint16(1),
	1006:  uint16(anon_sym_if),
	1007:  uint16(31),
	1008:  uint16(1),
	1009:  uint16(anon_sym_while),
	1010:  uint16(33),
	1011:  uint16(1),
	1012:  uint16(anon_sym_loop),
	1013:  uint16(35),
	1014:  uint16(1),
	1015:  uint16(anon_sym_const),
	1016:  uint16(37),
	1017:  uint16(1),
	1018:  uint16(anon_sym_break),
	1019:  uint16(39),
	1020:  uint16(1),
	1021:  uint16(anon_sym_continue),
	1022:  uint16(41),
	1023:  uint16(1),
	1024:  uint16(sym_integer_literal),
	1025:  uint16(43),
	1026:  uint16(1),
	1027:  uint16(sym_float_literal),
	1028:  uint16(120),
	1029:  uint16(1),
	1030:  uint16(anon_sym_LT),
	1031:  uint16(124),
	1032:  uint16(1),
	1033:  uint16(anon_sym_AMP),
	1034:  uint16(126),
	1035:  uint16(1),
	1036:  uint16(anon_sym_BANG),
	1037:  uint16(150),
	1038:  uint16(1),
	1039:  uint16(anon_sym_let),
	1040:  uint16(82),
	1041:  uint16(1),
	1042:  uint16(sym_field_expression),
	1043:  uint16(151),
	1044:  uint16(1),
	1045:  uint16(sym__expression),
	1046:  uint16(340),
	1047:  uint16(1),
	1048:  uint16(sym_scoped_identifier),
	1049:  uint16(45),
	1050:  uint16(2),
	1051:  uint16(anon_sym_true),
	1052:  uint16(anon_sym_false),
	1053:  uint16(329),
	1054:  uint16(2),
	1055:  uint16(sym_bracketed_type),
	1056:  uint16(sym__path),
	1057:  uint16(349),
	1058:  uint16(2),
	1059:  uint16(sym_let_condition),
	1060:  uint16(sym__condition),
	1061:  uint16(21),
	1062:  uint16(5),
	1063:  uint16(anon_sym_u8),
	1064:  uint16(anon_sym_u64),
	1065:  uint16(anon_sym_u128),
	1066:  uint16(anon_sym_bool),
	1067:  uint16(anon_sym_address),
	1068:  uint16(83),
	1069:  uint16(22),
	1070:  uint16(sym_generic_function),
	1071:  uint16(sym__expression_except_range),
	1072:  uint16(sym__expression_ending_with_block),
	1073:  uint16(sym_unary_expression),
	1074:  uint16(sym_reference_expression),
	1075:  uint16(sym_binary_expression),
	1076:  uint16(sym_assignment_expression),
	1077:  uint16(sym_type_cast_expression),
	1078:  uint16(sym_return_expression),
	1079:  uint16(sym_call_expression),
	1080:  uint16(sym_parenthesized_expression),
	1081:  uint16(sym_tuple_expression),
	1082:  uint16(sym_unit_expression),
	1083:  uint16(sym_if_expression),
	1084:  uint16(sym_while_expression),
	1085:  uint16(sym_loop_expression),
	1086:  uint16(sym_const_block),
	1087:  uint16(sym_break_expression),
	1088:  uint16(sym_continue_expression),
	1089:  uint16(sym_block),
	1090:  uint16(sym__literal),
	1091:  uint16(sym_boolean_literal),
	1092:  uint16(26),
	1093:  uint16(3),
	1094:  uint16(1),
	1095:  uint16(sym_comment),
	1096:  uint16(7),
	1097:  uint16(1),
	1098:  uint16(sym_identifier),
	1099:  uint16(9),
	1100:  uint16(1),
	1101:  uint16(anon_sym_COLON_COLON),
	1102:  uint16(11),
	1103:  uint16(1),
	1104:  uint16(anon_sym_LBRACE),
	1105:  uint16(19),
	1106:  uint16(1),
	1107:  uint16(anon_sym_LPAREN),
	1108:  uint16(27),
	1109:  uint16(1),
	1110:  uint16(anon_sym_return),
	1111:  uint16(29),
	1112:  uint16(1),
	1113:  uint16(anon_sym_if),
	1114:  uint16(31),
	1115:  uint16(1),
	1116:  uint16(anon_sym_while),
	1117:  uint16(33),
	1118:  uint16(1),
	1119:  uint16(anon_sym_loop),
	1120:  uint16(35),
	1121:  uint16(1),
	1122:  uint16(anon_sym_const),
	1123:  uint16(37),
	1124:  uint16(1),
	1125:  uint16(anon_sym_break),
	1126:  uint16(39),
	1127:  uint16(1),
	1128:  uint16(anon_sym_continue),
	1129:  uint16(41),
	1130:  uint16(1),
	1131:  uint16(sym_integer_literal),
	1132:  uint16(43),
	1133:  uint16(1),
	1134:  uint16(sym_float_literal),
	1135:  uint16(120),
	1136:  uint16(1),
	1137:  uint16(anon_sym_LT),
	1138:  uint16(124),
	1139:  uint16(1),
	1140:  uint16(anon_sym_AMP),
	1141:  uint16(126),
	1142:  uint16(1),
	1143:  uint16(anon_sym_BANG),
	1144:  uint16(150),
	1145:  uint16(1),
	1146:  uint16(anon_sym_let),
	1147:  uint16(82),
	1148:  uint16(1),
	1149:  uint16(sym_field_expression),
	1150:  uint16(151),
	1151:  uint16(1),
	1152:  uint16(sym__expression),
	1153:  uint16(340),
	1154:  uint16(1),
	1155:  uint16(sym_scoped_identifier),
	1156:  uint16(45),
	1157:  uint16(2),
	1158:  uint16(anon_sym_true),
	1159:  uint16(anon_sym_false),
	1160:  uint16(329),
	1161:  uint16(2),
	1162:  uint16(sym_bracketed_type),
	1163:  uint16(sym__path),
	1164:  uint16(348),
	1165:  uint16(2),
	1166:  uint16(sym_let_condition),
	1167:  uint16(sym__condition),
	1168:  uint16(21),
	1169:  uint16(5),
	1170:  uint16(anon_sym_u8),
	1171:  uint16(anon_sym_u64),
	1172:  uint16(anon_sym_u128),
	1173:  uint16(anon_sym_bool),
	1174:  uint16(anon_sym_address),
	1175:  uint16(83),
	1176:  uint16(22),
	1177:  uint16(sym_generic_function),
	1178:  uint16(sym__expression_except_range),
	1179:  uint16(sym__expression_ending_with_block),
	1180:  uint16(sym_unary_expression),
	1181:  uint16(sym_reference_expression),
	1182:  uint16(sym_binary_expression),
	1183:  uint16(sym_assignment_expression),
	1184:  uint16(sym_type_cast_expression),
	1185:  uint16(sym_return_expression),
	1186:  uint16(sym_call_expression),
	1187:  uint16(sym_parenthesized_expression),
	1188:  uint16(sym_tuple_expression),
	1189:  uint16(sym_unit_expression),
	1190:  uint16(sym_if_expression),
	1191:  uint16(sym_while_expression),
	1192:  uint16(sym_loop_expression),
	1193:  uint16(sym_const_block),
	1194:  uint16(sym_break_expression),
	1195:  uint16(sym_continue_expression),
	1196:  uint16(sym_block),
	1197:  uint16(sym__literal),
	1198:  uint16(sym_boolean_literal),
	1199:  uint16(26),
	1200:  uint16(3),
	1201:  uint16(1),
	1202:  uint16(sym_comment),
	1203:  uint16(7),
	1204:  uint16(1),
	1205:  uint16(sym_identifier),
	1206:  uint16(9),
	1207:  uint16(1),
	1208:  uint16(anon_sym_COLON_COLON),
	1209:  uint16(11),
	1210:  uint16(1),
	1211:  uint16(anon_sym_LBRACE),
	1212:  uint16(19),
	1213:  uint16(1),
	1214:  uint16(anon_sym_LPAREN),
	1215:  uint16(27),
	1216:  uint16(1),
	1217:  uint16(anon_sym_return),
	1218:  uint16(29),
	1219:  uint16(1),
	1220:  uint16(anon_sym_if),
	1221:  uint16(31),
	1222:  uint16(1),
	1223:  uint16(anon_sym_while),
	1224:  uint16(33),
	1225:  uint16(1),
	1226:  uint16(anon_sym_loop),
	1227:  uint16(35),
	1228:  uint16(1),
	1229:  uint16(anon_sym_const),
	1230:  uint16(37),
	1231:  uint16(1),
	1232:  uint16(anon_sym_break),
	1233:  uint16(39),
	1234:  uint16(1),
	1235:  uint16(anon_sym_continue),
	1236:  uint16(41),
	1237:  uint16(1),
	1238:  uint16(sym_integer_literal),
	1239:  uint16(43),
	1240:  uint16(1),
	1241:  uint16(sym_float_literal),
	1242:  uint16(120),
	1243:  uint16(1),
	1244:  uint16(anon_sym_LT),
	1245:  uint16(124),
	1246:  uint16(1),
	1247:  uint16(anon_sym_AMP),
	1248:  uint16(126),
	1249:  uint16(1),
	1250:  uint16(anon_sym_BANG),
	1251:  uint16(150),
	1252:  uint16(1),
	1253:  uint16(anon_sym_let),
	1254:  uint16(82),
	1255:  uint16(1),
	1256:  uint16(sym_field_expression),
	1257:  uint16(151),
	1258:  uint16(1),
	1259:  uint16(sym__expression),
	1260:  uint16(340),
	1261:  uint16(1),
	1262:  uint16(sym_scoped_identifier),
	1263:  uint16(45),
	1264:  uint16(2),
	1265:  uint16(anon_sym_true),
	1266:  uint16(anon_sym_false),
	1267:  uint16(321),
	1268:  uint16(2),
	1269:  uint16(sym_let_condition),
	1270:  uint16(sym__condition),
	1271:  uint16(329),
	1272:  uint16(2),
	1273:  uint16(sym_bracketed_type),
	1274:  uint16(sym__path),
	1275:  uint16(21),
	1276:  uint16(5),
	1277:  uint16(anon_sym_u8),
	1278:  uint16(anon_sym_u64),
	1279:  uint16(anon_sym_u128),
	1280:  uint16(anon_sym_bool),
	1281:  uint16(anon_sym_address),
	1282:  uint16(83),
	1283:  uint16(22),
	1284:  uint16(sym_generic_function),
	1285:  uint16(sym__expression_except_range),
	1286:  uint16(sym__expression_ending_with_block),
	1287:  uint16(sym_unary_expression),
	1288:  uint16(sym_reference_expression),
	1289:  uint16(sym_binary_expression),
	1290:  uint16(sym_assignment_expression),
	1291:  uint16(sym_type_cast_expression),
	1292:  uint16(sym_return_expression),
	1293:  uint16(sym_call_expression),
	1294:  uint16(sym_parenthesized_expression),
	1295:  uint16(sym_tuple_expression),
	1296:  uint16(sym_unit_expression),
	1297:  uint16(sym_if_expression),
	1298:  uint16(sym_while_expression),
	1299:  uint16(sym_loop_expression),
	1300:  uint16(sym_const_block),
	1301:  uint16(sym_break_expression),
	1302:  uint16(sym_continue_expression),
	1303:  uint16(sym_block),
	1304:  uint16(sym__literal),
	1305:  uint16(sym_boolean_literal),
	1306:  uint16(26),
	1307:  uint16(3),
	1308:  uint16(1),
	1309:  uint16(sym_comment),
	1310:  uint16(7),
	1311:  uint16(1),
	1312:  uint16(sym_identifier),
	1313:  uint16(9),
	1314:  uint16(1),
	1315:  uint16(anon_sym_COLON_COLON),
	1316:  uint16(11),
	1317:  uint16(1),
	1318:  uint16(anon_sym_LBRACE),
	1319:  uint16(19),
	1320:  uint16(1),
	1321:  uint16(anon_sym_LPAREN),
	1322:  uint16(27),
	1323:  uint16(1),
	1324:  uint16(anon_sym_return),
	1325:  uint16(29),
	1326:  uint16(1),
	1327:  uint16(anon_sym_if),
	1328:  uint16(31),
	1329:  uint16(1),
	1330:  uint16(anon_sym_while),
	1331:  uint16(33),
	1332:  uint16(1),
	1333:  uint16(anon_sym_loop),
	1334:  uint16(35),
	1335:  uint16(1),
	1336:  uint16(anon_sym_const),
	1337:  uint16(37),
	1338:  uint16(1),
	1339:  uint16(anon_sym_break),
	1340:  uint16(39),
	1341:  uint16(1),
	1342:  uint16(anon_sym_continue),
	1343:  uint16(41),
	1344:  uint16(1),
	1345:  uint16(sym_integer_literal),
	1346:  uint16(43),
	1347:  uint16(1),
	1348:  uint16(sym_float_literal),
	1349:  uint16(120),
	1350:  uint16(1),
	1351:  uint16(anon_sym_LT),
	1352:  uint16(124),
	1353:  uint16(1),
	1354:  uint16(anon_sym_AMP),
	1355:  uint16(126),
	1356:  uint16(1),
	1357:  uint16(anon_sym_BANG),
	1358:  uint16(150),
	1359:  uint16(1),
	1360:  uint16(anon_sym_let),
	1361:  uint16(82),
	1362:  uint16(1),
	1363:  uint16(sym_field_expression),
	1364:  uint16(151),
	1365:  uint16(1),
	1366:  uint16(sym__expression),
	1367:  uint16(340),
	1368:  uint16(1),
	1369:  uint16(sym_scoped_identifier),
	1370:  uint16(45),
	1371:  uint16(2),
	1372:  uint16(anon_sym_true),
	1373:  uint16(anon_sym_false),
	1374:  uint16(320),
	1375:  uint16(2),
	1376:  uint16(sym_let_condition),
	1377:  uint16(sym__condition),
	1378:  uint16(329),
	1379:  uint16(2),
	1380:  uint16(sym_bracketed_type),
	1381:  uint16(sym__path),
	1382:  uint16(21),
	1383:  uint16(5),
	1384:  uint16(anon_sym_u8),
	1385:  uint16(anon_sym_u64),
	1386:  uint16(anon_sym_u128),
	1387:  uint16(anon_sym_bool),
	1388:  uint16(anon_sym_address),
	1389:  uint16(83),
	1390:  uint16(22),
	1391:  uint16(sym_generic_function),
	1392:  uint16(sym__expression_except_range),
	1393:  uint16(sym__expression_ending_with_block),
	1394:  uint16(sym_unary_expression),
	1395:  uint16(sym_reference_expression),
	1396:  uint16(sym_binary_expression),
	1397:  uint16(sym_assignment_expression),
	1398:  uint16(sym_type_cast_expression),
	1399:  uint16(sym_return_expression),
	1400:  uint16(sym_call_expression),
	1401:  uint16(sym_parenthesized_expression),
	1402:  uint16(sym_tuple_expression),
	1403:  uint16(sym_unit_expression),
	1404:  uint16(sym_if_expression),
	1405:  uint16(sym_while_expression),
	1406:  uint16(sym_loop_expression),
	1407:  uint16(sym_const_block),
	1408:  uint16(sym_break_expression),
	1409:  uint16(sym_continue_expression),
	1410:  uint16(sym_block),
	1411:  uint16(sym__literal),
	1412:  uint16(sym_boolean_literal),
	1413:  uint16(26),
	1414:  uint16(3),
	1415:  uint16(1),
	1416:  uint16(sym_comment),
	1417:  uint16(7),
	1418:  uint16(1),
	1419:  uint16(sym_identifier),
	1420:  uint16(9),
	1421:  uint16(1),
	1422:  uint16(anon_sym_COLON_COLON),
	1423:  uint16(11),
	1424:  uint16(1),
	1425:  uint16(anon_sym_LBRACE),
	1426:  uint16(19),
	1427:  uint16(1),
	1428:  uint16(anon_sym_LPAREN),
	1429:  uint16(27),
	1430:  uint16(1),
	1431:  uint16(anon_sym_return),
	1432:  uint16(29),
	1433:  uint16(1),
	1434:  uint16(anon_sym_if),
	1435:  uint16(31),
	1436:  uint16(1),
	1437:  uint16(anon_sym_while),
	1438:  uint16(33),
	1439:  uint16(1),
	1440:  uint16(anon_sym_loop),
	1441:  uint16(35),
	1442:  uint16(1),
	1443:  uint16(anon_sym_const),
	1444:  uint16(37),
	1445:  uint16(1),
	1446:  uint16(anon_sym_break),
	1447:  uint16(39),
	1448:  uint16(1),
	1449:  uint16(anon_sym_continue),
	1450:  uint16(41),
	1451:  uint16(1),
	1452:  uint16(sym_integer_literal),
	1453:  uint16(43),
	1454:  uint16(1),
	1455:  uint16(sym_float_literal),
	1456:  uint16(120),
	1457:  uint16(1),
	1458:  uint16(anon_sym_LT),
	1459:  uint16(124),
	1460:  uint16(1),
	1461:  uint16(anon_sym_AMP),
	1462:  uint16(126),
	1463:  uint16(1),
	1464:  uint16(anon_sym_BANG),
	1465:  uint16(146),
	1466:  uint16(1),
	1467:  uint16(anon_sym_POUND),
	1468:  uint16(82),
	1469:  uint16(1),
	1470:  uint16(sym_field_expression),
	1471:  uint16(149),
	1472:  uint16(1),
	1473:  uint16(sym__expression),
	1474:  uint16(340),
	1475:  uint16(1),
	1476:  uint16(sym_scoped_identifier),
	1477:  uint16(45),
	1478:  uint16(2),
	1479:  uint16(anon_sym_true),
	1480:  uint16(anon_sym_false),
	1481:  uint16(139),
	1482:  uint16(2),
	1483:  uint16(sym_attribute_item),
	1484:  uint16(aux_sym_tuple_expression_repeat1),
	1485:  uint16(329),
	1486:  uint16(2),
	1487:  uint16(sym_bracketed_type),
	1488:  uint16(sym__path),
	1489:  uint16(21),
	1490:  uint16(5),
	1491:  uint16(anon_sym_u8),
	1492:  uint16(anon_sym_u64),
	1493:  uint16(anon_sym_u128),
	1494:  uint16(anon_sym_bool),
	1495:  uint16(anon_sym_address),
	1496:  uint16(83),
	1497:  uint16(22),
	1498:  uint16(sym_generic_function),
	1499:  uint16(sym__expression_except_range),
	1500:  uint16(sym__expression_ending_with_block),
	1501:  uint16(sym_unary_expression),
	1502:  uint16(sym_reference_expression),
	1503:  uint16(sym_binary_expression),
	1504:  uint16(sym_assignment_expression),
	1505:  uint16(sym_type_cast_expression),
	1506:  uint16(sym_return_expression),
	1507:  uint16(sym_call_expression),
	1508:  uint16(sym_parenthesized_expression),
	1509:  uint16(sym_tuple_expression),
	1510:  uint16(sym_unit_expression),
	1511:  uint16(sym_if_expression),
	1512:  uint16(sym_while_expression),
	1513:  uint16(sym_loop_expression),
	1514:  uint16(sym_const_block),
	1515:  uint16(sym_break_expression),
	1516:  uint16(sym_continue_expression),
	1517:  uint16(sym_block),
	1518:  uint16(sym__literal),
	1519:  uint16(sym_boolean_literal),
	1520:  uint16(26),
	1521:  uint16(3),
	1522:  uint16(1),
	1523:  uint16(sym_comment),
	1524:  uint16(7),
	1525:  uint16(1),
	1526:  uint16(sym_identifier),
	1527:  uint16(9),
	1528:  uint16(1),
	1529:  uint16(anon_sym_COLON_COLON),
	1530:  uint16(11),
	1531:  uint16(1),
	1532:  uint16(anon_sym_LBRACE),
	1533:  uint16(19),
	1534:  uint16(1),
	1535:  uint16(anon_sym_LPAREN),
	1536:  uint16(27),
	1537:  uint16(1),
	1538:  uint16(anon_sym_return),
	1539:  uint16(29),
	1540:  uint16(1),
	1541:  uint16(anon_sym_if),
	1542:  uint16(31),
	1543:  uint16(1),
	1544:  uint16(anon_sym_while),
	1545:  uint16(33),
	1546:  uint16(1),
	1547:  uint16(anon_sym_loop),
	1548:  uint16(35),
	1549:  uint16(1),
	1550:  uint16(anon_sym_const),
	1551:  uint16(37),
	1552:  uint16(1),
	1553:  uint16(anon_sym_break),
	1554:  uint16(39),
	1555:  uint16(1),
	1556:  uint16(anon_sym_continue),
	1557:  uint16(41),
	1558:  uint16(1),
	1559:  uint16(sym_integer_literal),
	1560:  uint16(43),
	1561:  uint16(1),
	1562:  uint16(sym_float_literal),
	1563:  uint16(120),
	1564:  uint16(1),
	1565:  uint16(anon_sym_LT),
	1566:  uint16(124),
	1567:  uint16(1),
	1568:  uint16(anon_sym_AMP),
	1569:  uint16(126),
	1570:  uint16(1),
	1571:  uint16(anon_sym_BANG),
	1572:  uint16(152),
	1573:  uint16(1),
	1574:  uint16(anon_sym_RPAREN),
	1575:  uint16(19),
	1576:  uint16(1),
	1577:  uint16(aux_sym_tuple_expression_repeat2),
	1578:  uint16(82),
	1579:  uint16(1),
	1580:  uint16(sym_field_expression),
	1581:  uint16(145),
	1582:  uint16(1),
	1583:  uint16(sym__expression),
	1584:  uint16(340),
	1585:  uint16(1),
	1586:  uint16(sym_scoped_identifier),
	1587:  uint16(45),
	1588:  uint16(2),
	1589:  uint16(anon_sym_true),
	1590:  uint16(anon_sym_false),
	1591:  uint16(329),
	1592:  uint16(2),
	1593:  uint16(sym_bracketed_type),
	1594:  uint16(sym__path),
	1595:  uint16(21),
	1596:  uint16(5),
	1597:  uint16(anon_sym_u8),
	1598:  uint16(anon_sym_u64),
	1599:  uint16(anon_sym_u128),
	1600:  uint16(anon_sym_bool),
	1601:  uint16(anon_sym_address),
	1602:  uint16(83),
	1603:  uint16(22),
	1604:  uint16(sym_generic_function),
	1605:  uint16(sym__expression_except_range),
	1606:  uint16(sym__expression_ending_with_block),
	1607:  uint16(sym_unary_expression),
	1608:  uint16(sym_reference_expression),
	1609:  uint16(sym_binary_expression),
	1610:  uint16(sym_assignment_expression),
	1611:  uint16(sym_type_cast_expression),
	1612:  uint16(sym_return_expression),
	1613:  uint16(sym_call_expression),
	1614:  uint16(sym_parenthesized_expression),
	1615:  uint16(sym_tuple_expression),
	1616:  uint16(sym_unit_expression),
	1617:  uint16(sym_if_expression),
	1618:  uint16(sym_while_expression),
	1619:  uint16(sym_loop_expression),
	1620:  uint16(sym_const_block),
	1621:  uint16(sym_break_expression),
	1622:  uint16(sym_continue_expression),
	1623:  uint16(sym_block),
	1624:  uint16(sym__literal),
	1625:  uint16(sym_boolean_literal),
	1626:  uint16(26),
	1627:  uint16(3),
	1628:  uint16(1),
	1629:  uint16(sym_comment),
	1630:  uint16(7),
	1631:  uint16(1),
	1632:  uint16(sym_identifier),
	1633:  uint16(9),
	1634:  uint16(1),
	1635:  uint16(anon_sym_COLON_COLON),
	1636:  uint16(11),
	1637:  uint16(1),
	1638:  uint16(anon_sym_LBRACE),
	1639:  uint16(19),
	1640:  uint16(1),
	1641:  uint16(anon_sym_LPAREN),
	1642:  uint16(27),
	1643:  uint16(1),
	1644:  uint16(anon_sym_return),
	1645:  uint16(29),
	1646:  uint16(1),
	1647:  uint16(anon_sym_if),
	1648:  uint16(31),
	1649:  uint16(1),
	1650:  uint16(anon_sym_while),
	1651:  uint16(33),
	1652:  uint16(1),
	1653:  uint16(anon_sym_loop),
	1654:  uint16(35),
	1655:  uint16(1),
	1656:  uint16(anon_sym_const),
	1657:  uint16(37),
	1658:  uint16(1),
	1659:  uint16(anon_sym_break),
	1660:  uint16(39),
	1661:  uint16(1),
	1662:  uint16(anon_sym_continue),
	1663:  uint16(41),
	1664:  uint16(1),
	1665:  uint16(sym_integer_literal),
	1666:  uint16(43),
	1667:  uint16(1),
	1668:  uint16(sym_float_literal),
	1669:  uint16(120),
	1670:  uint16(1),
	1671:  uint16(anon_sym_LT),
	1672:  uint16(124),
	1673:  uint16(1),
	1674:  uint16(anon_sym_AMP),
	1675:  uint16(126),
	1676:  uint16(1),
	1677:  uint16(anon_sym_BANG),
	1678:  uint16(154),
	1679:  uint16(1),
	1680:  uint16(anon_sym_RPAREN),
	1681:  uint16(21),
	1682:  uint16(1),
	1683:  uint16(aux_sym_tuple_expression_repeat2),
	1684:  uint16(82),
	1685:  uint16(1),
	1686:  uint16(sym_field_expression),
	1687:  uint16(143),
	1688:  uint16(1),
	1689:  uint16(sym__expression),
	1690:  uint16(340),
	1691:  uint16(1),
	1692:  uint16(sym_scoped_identifier),
	1693:  uint16(45),
	1694:  uint16(2),
	1695:  uint16(anon_sym_true),
	1696:  uint16(anon_sym_false),
	1697:  uint16(329),
	1698:  uint16(2),
	1699:  uint16(sym_bracketed_type),
	1700:  uint16(sym__path),
	1701:  uint16(21),
	1702:  uint16(5),
	1703:  uint16(anon_sym_u8),
	1704:  uint16(anon_sym_u64),
	1705:  uint16(anon_sym_u128),
	1706:  uint16(anon_sym_bool),
	1707:  uint16(anon_sym_address),
	1708:  uint16(83),
	1709:  uint16(22),
	1710:  uint16(sym_generic_function),
	1711:  uint16(sym__expression_except_range),
	1712:  uint16(sym__expression_ending_with_block),
	1713:  uint16(sym_unary_expression),
	1714:  uint16(sym_reference_expression),
	1715:  uint16(sym_binary_expression),
	1716:  uint16(sym_assignment_expression),
	1717:  uint16(sym_type_cast_expression),
	1718:  uint16(sym_return_expression),
	1719:  uint16(sym_call_expression),
	1720:  uint16(sym_parenthesized_expression),
	1721:  uint16(sym_tuple_expression),
	1722:  uint16(sym_unit_expression),
	1723:  uint16(sym_if_expression),
	1724:  uint16(sym_while_expression),
	1725:  uint16(sym_loop_expression),
	1726:  uint16(sym_const_block),
	1727:  uint16(sym_break_expression),
	1728:  uint16(sym_continue_expression),
	1729:  uint16(sym_block),
	1730:  uint16(sym__literal),
	1731:  uint16(sym_boolean_literal),
	1732:  uint16(26),
	1733:  uint16(3),
	1734:  uint16(1),
	1735:  uint16(sym_comment),
	1736:  uint16(7),
	1737:  uint16(1),
	1738:  uint16(sym_identifier),
	1739:  uint16(9),
	1740:  uint16(1),
	1741:  uint16(anon_sym_COLON_COLON),
	1742:  uint16(11),
	1743:  uint16(1),
	1744:  uint16(anon_sym_LBRACE),
	1745:  uint16(19),
	1746:  uint16(1),
	1747:  uint16(anon_sym_LPAREN),
	1748:  uint16(27),
	1749:  uint16(1),
	1750:  uint16(anon_sym_return),
	1751:  uint16(29),
	1752:  uint16(1),
	1753:  uint16(anon_sym_if),
	1754:  uint16(31),
	1755:  uint16(1),
	1756:  uint16(anon_sym_while),
	1757:  uint16(33),
	1758:  uint16(1),
	1759:  uint16(anon_sym_loop),
	1760:  uint16(35),
	1761:  uint16(1),
	1762:  uint16(anon_sym_const),
	1763:  uint16(37),
	1764:  uint16(1),
	1765:  uint16(anon_sym_break),
	1766:  uint16(39),
	1767:  uint16(1),
	1768:  uint16(anon_sym_continue),
	1769:  uint16(41),
	1770:  uint16(1),
	1771:  uint16(sym_integer_literal),
	1772:  uint16(43),
	1773:  uint16(1),
	1774:  uint16(sym_float_literal),
	1775:  uint16(120),
	1776:  uint16(1),
	1777:  uint16(anon_sym_LT),
	1778:  uint16(124),
	1779:  uint16(1),
	1780:  uint16(anon_sym_AMP),
	1781:  uint16(126),
	1782:  uint16(1),
	1783:  uint16(anon_sym_BANG),
	1784:  uint16(156),
	1785:  uint16(1),
	1786:  uint16(anon_sym_RPAREN),
	1787:  uint16(16),
	1788:  uint16(1),
	1789:  uint16(aux_sym_tuple_expression_repeat2),
	1790:  uint16(82),
	1791:  uint16(1),
	1792:  uint16(sym_field_expression),
	1793:  uint16(147),
	1794:  uint16(1),
	1795:  uint16(sym__expression),
	1796:  uint16(340),
	1797:  uint16(1),
	1798:  uint16(sym_scoped_identifier),
	1799:  uint16(45),
	1800:  uint16(2),
	1801:  uint16(anon_sym_true),
	1802:  uint16(anon_sym_false),
	1803:  uint16(329),
	1804:  uint16(2),
	1805:  uint16(sym_bracketed_type),
	1806:  uint16(sym__path),
	1807:  uint16(21),
	1808:  uint16(5),
	1809:  uint16(anon_sym_u8),
	1810:  uint16(anon_sym_u64),
	1811:  uint16(anon_sym_u128),
	1812:  uint16(anon_sym_bool),
	1813:  uint16(anon_sym_address),
	1814:  uint16(83),
	1815:  uint16(22),
	1816:  uint16(sym_generic_function),
	1817:  uint16(sym__expression_except_range),
	1818:  uint16(sym__expression_ending_with_block),
	1819:  uint16(sym_unary_expression),
	1820:  uint16(sym_reference_expression),
	1821:  uint16(sym_binary_expression),
	1822:  uint16(sym_assignment_expression),
	1823:  uint16(sym_type_cast_expression),
	1824:  uint16(sym_return_expression),
	1825:  uint16(sym_call_expression),
	1826:  uint16(sym_parenthesized_expression),
	1827:  uint16(sym_tuple_expression),
	1828:  uint16(sym_unit_expression),
	1829:  uint16(sym_if_expression),
	1830:  uint16(sym_while_expression),
	1831:  uint16(sym_loop_expression),
	1832:  uint16(sym_const_block),
	1833:  uint16(sym_break_expression),
	1834:  uint16(sym_continue_expression),
	1835:  uint16(sym_block),
	1836:  uint16(sym__literal),
	1837:  uint16(sym_boolean_literal),
	1838:  uint16(26),
	1839:  uint16(3),
	1840:  uint16(1),
	1841:  uint16(sym_comment),
	1842:  uint16(158),
	1843:  uint16(1),
	1844:  uint16(sym_identifier),
	1845:  uint16(161),
	1846:  uint16(1),
	1847:  uint16(anon_sym_COLON_COLON),
	1848:  uint16(164),
	1849:  uint16(1),
	1850:  uint16(anon_sym_LBRACE),
	1851:  uint16(167),
	1852:  uint16(1),
	1853:  uint16(anon_sym_LT),
	1854:  uint16(170),
	1855:  uint16(1),
	1856:  uint16(anon_sym_LPAREN),
	1857:  uint16(173),
	1858:  uint16(1),
	1859:  uint16(anon_sym_RPAREN),
	1860:  uint16(178),
	1861:  uint16(1),
	1862:  uint16(anon_sym_AMP),
	1863:  uint16(181),
	1864:  uint16(1),
	1865:  uint16(anon_sym_BANG),
	1866:  uint16(184),
	1867:  uint16(1),
	1868:  uint16(anon_sym_return),
	1869:  uint16(187),
	1870:  uint16(1),
	1871:  uint16(anon_sym_if),
	1872:  uint16(190),
	1873:  uint16(1),
	1874:  uint16(anon_sym_while),
	1875:  uint16(193),
	1876:  uint16(1),
	1877:  uint16(anon_sym_loop),
	1878:  uint16(196),
	1879:  uint16(1),
	1880:  uint16(anon_sym_const),
	1881:  uint16(199),
	1882:  uint16(1),
	1883:  uint16(anon_sym_break),
	1884:  uint16(202),
	1885:  uint16(1),
	1886:  uint16(anon_sym_continue),
	1887:  uint16(205),
	1888:  uint16(1),
	1889:  uint16(sym_integer_literal),
	1890:  uint16(208),
	1891:  uint16(1),
	1892:  uint16(sym_float_literal),
	1893:  uint16(19),
	1894:  uint16(1),
	1895:  uint16(aux_sym_tuple_expression_repeat2),
	1896:  uint16(82),
	1897:  uint16(1),
	1898:  uint16(sym_field_expression),
	1899:  uint16(152),
	1900:  uint16(1),
	1901:  uint16(sym__expression),
	1902:  uint16(340),
	1903:  uint16(1),
	1904:  uint16(sym_scoped_identifier),
	1905:  uint16(211),
	1906:  uint16(2),
	1907:  uint16(anon_sym_true),
	1908:  uint16(anon_sym_false),
	1909:  uint16(329),
	1910:  uint16(2),
	1911:  uint16(sym_bracketed_type),
	1912:  uint16(sym__path),
	1913:  uint16(175),
	1914:  uint16(5),
	1915:  uint16(anon_sym_u8),
	1916:  uint16(anon_sym_u64),
	1917:  uint16(anon_sym_u128),
	1918:  uint16(anon_sym_bool),
	1919:  uint16(anon_sym_address),
	1920:  uint16(83),
	1921:  uint16(22),
	1922:  uint16(sym_generic_function),
	1923:  uint16(sym__expression_except_range),
	1924:  uint16(sym__expression_ending_with_block),
	1925:  uint16(sym_unary_expression),
	1926:  uint16(sym_reference_expression),
	1927:  uint16(sym_binary_expression),
	1928:  uint16(sym_assignment_expression),
	1929:  uint16(sym_type_cast_expression),
	1930:  uint16(sym_return_expression),
	1931:  uint16(sym_call_expression),
	1932:  uint16(sym_parenthesized_expression),
	1933:  uint16(sym_tuple_expression),
	1934:  uint16(sym_unit_expression),
	1935:  uint16(sym_if_expression),
	1936:  uint16(sym_while_expression),
	1937:  uint16(sym_loop_expression),
	1938:  uint16(sym_const_block),
	1939:  uint16(sym_break_expression),
	1940:  uint16(sym_continue_expression),
	1941:  uint16(sym_block),
	1942:  uint16(sym__literal),
	1943:  uint16(sym_boolean_literal),
	1944:  uint16(26),
	1945:  uint16(3),
	1946:  uint16(1),
	1947:  uint16(sym_comment),
	1948:  uint16(7),
	1949:  uint16(1),
	1950:  uint16(sym_identifier),
	1951:  uint16(9),
	1952:  uint16(1),
	1953:  uint16(anon_sym_COLON_COLON),
	1954:  uint16(11),
	1955:  uint16(1),
	1956:  uint16(anon_sym_LBRACE),
	1957:  uint16(19),
	1958:  uint16(1),
	1959:  uint16(anon_sym_LPAREN),
	1960:  uint16(27),
	1961:  uint16(1),
	1962:  uint16(anon_sym_return),
	1963:  uint16(29),
	1964:  uint16(1),
	1965:  uint16(anon_sym_if),
	1966:  uint16(31),
	1967:  uint16(1),
	1968:  uint16(anon_sym_while),
	1969:  uint16(33),
	1970:  uint16(1),
	1971:  uint16(anon_sym_loop),
	1972:  uint16(35),
	1973:  uint16(1),
	1974:  uint16(anon_sym_const),
	1975:  uint16(37),
	1976:  uint16(1),
	1977:  uint16(anon_sym_break),
	1978:  uint16(39),
	1979:  uint16(1),
	1980:  uint16(anon_sym_continue),
	1981:  uint16(41),
	1982:  uint16(1),
	1983:  uint16(sym_integer_literal),
	1984:  uint16(43),
	1985:  uint16(1),
	1986:  uint16(sym_float_literal),
	1987:  uint16(120),
	1988:  uint16(1),
	1989:  uint16(anon_sym_LT),
	1990:  uint16(124),
	1991:  uint16(1),
	1992:  uint16(anon_sym_AMP),
	1993:  uint16(126),
	1994:  uint16(1),
	1995:  uint16(anon_sym_BANG),
	1996:  uint16(214),
	1997:  uint16(1),
	1998:  uint16(anon_sym_COMMA),
	1999:  uint16(216),
	2000:  uint16(1),
	2001:  uint16(anon_sym_RPAREN),
	2002:  uint16(82),
	2003:  uint16(1),
	2004:  uint16(sym_field_expression),
	2005:  uint16(141),
	2006:  uint16(1),
	2007:  uint16(sym__expression),
	2008:  uint16(340),
	2009:  uint16(1),
	2010:  uint16(sym_scoped_identifier),
	2011:  uint16(45),
	2012:  uint16(2),
	2013:  uint16(anon_sym_true),
	2014:  uint16(anon_sym_false),
	2015:  uint16(329),
	2016:  uint16(2),
	2017:  uint16(sym_bracketed_type),
	2018:  uint16(sym__path),
	2019:  uint16(21),
	2020:  uint16(5),
	2021:  uint16(anon_sym_u8),
	2022:  uint16(anon_sym_u64),
	2023:  uint16(anon_sym_u128),
	2024:  uint16(anon_sym_bool),
	2025:  uint16(anon_sym_address),
	2026:  uint16(83),
	2027:  uint16(22),
	2028:  uint16(sym_generic_function),
	2029:  uint16(sym__expression_except_range),
	2030:  uint16(sym__expression_ending_with_block),
	2031:  uint16(sym_unary_expression),
	2032:  uint16(sym_reference_expression),
	2033:  uint16(sym_binary_expression),
	2034:  uint16(sym_assignment_expression),
	2035:  uint16(sym_type_cast_expression),
	2036:  uint16(sym_return_expression),
	2037:  uint16(sym_call_expression),
	2038:  uint16(sym_parenthesized_expression),
	2039:  uint16(sym_tuple_expression),
	2040:  uint16(sym_unit_expression),
	2041:  uint16(sym_if_expression),
	2042:  uint16(sym_while_expression),
	2043:  uint16(sym_loop_expression),
	2044:  uint16(sym_const_block),
	2045:  uint16(sym_break_expression),
	2046:  uint16(sym_continue_expression),
	2047:  uint16(sym_block),
	2048:  uint16(sym__literal),
	2049:  uint16(sym_boolean_literal),
	2050:  uint16(26),
	2051:  uint16(3),
	2052:  uint16(1),
	2053:  uint16(sym_comment),
	2054:  uint16(7),
	2055:  uint16(1),
	2056:  uint16(sym_identifier),
	2057:  uint16(9),
	2058:  uint16(1),
	2059:  uint16(anon_sym_COLON_COLON),
	2060:  uint16(11),
	2061:  uint16(1),
	2062:  uint16(anon_sym_LBRACE),
	2063:  uint16(19),
	2064:  uint16(1),
	2065:  uint16(anon_sym_LPAREN),
	2066:  uint16(27),
	2067:  uint16(1),
	2068:  uint16(anon_sym_return),
	2069:  uint16(29),
	2070:  uint16(1),
	2071:  uint16(anon_sym_if),
	2072:  uint16(31),
	2073:  uint16(1),
	2074:  uint16(anon_sym_while),
	2075:  uint16(33),
	2076:  uint16(1),
	2077:  uint16(anon_sym_loop),
	2078:  uint16(35),
	2079:  uint16(1),
	2080:  uint16(anon_sym_const),
	2081:  uint16(37),
	2082:  uint16(1),
	2083:  uint16(anon_sym_break),
	2084:  uint16(39),
	2085:  uint16(1),
	2086:  uint16(anon_sym_continue),
	2087:  uint16(41),
	2088:  uint16(1),
	2089:  uint16(sym_integer_literal),
	2090:  uint16(43),
	2091:  uint16(1),
	2092:  uint16(sym_float_literal),
	2093:  uint16(120),
	2094:  uint16(1),
	2095:  uint16(anon_sym_LT),
	2096:  uint16(124),
	2097:  uint16(1),
	2098:  uint16(anon_sym_AMP),
	2099:  uint16(126),
	2100:  uint16(1),
	2101:  uint16(anon_sym_BANG),
	2102:  uint16(156),
	2103:  uint16(1),
	2104:  uint16(anon_sym_RPAREN),
	2105:  uint16(19),
	2106:  uint16(1),
	2107:  uint16(aux_sym_tuple_expression_repeat2),
	2108:  uint16(82),
	2109:  uint16(1),
	2110:  uint16(sym_field_expression),
	2111:  uint16(147),
	2112:  uint16(1),
	2113:  uint16(sym__expression),
	2114:  uint16(340),
	2115:  uint16(1),
	2116:  uint16(sym_scoped_identifier),
	2117:  uint16(45),
	2118:  uint16(2),
	2119:  uint16(anon_sym_true),
	2120:  uint16(anon_sym_false),
	2121:  uint16(329),
	2122:  uint16(2),
	2123:  uint16(sym_bracketed_type),
	2124:  uint16(sym__path),
	2125:  uint16(21),
	2126:  uint16(5),
	2127:  uint16(anon_sym_u8),
	2128:  uint16(anon_sym_u64),
	2129:  uint16(anon_sym_u128),
	2130:  uint16(anon_sym_bool),
	2131:  uint16(anon_sym_address),
	2132:  uint16(83),
	2133:  uint16(22),
	2134:  uint16(sym_generic_function),
	2135:  uint16(sym__expression_except_range),
	2136:  uint16(sym__expression_ending_with_block),
	2137:  uint16(sym_unary_expression),
	2138:  uint16(sym_reference_expression),
	2139:  uint16(sym_binary_expression),
	2140:  uint16(sym_assignment_expression),
	2141:  uint16(sym_type_cast_expression),
	2142:  uint16(sym_return_expression),
	2143:  uint16(sym_call_expression),
	2144:  uint16(sym_parenthesized_expression),
	2145:  uint16(sym_tuple_expression),
	2146:  uint16(sym_unit_expression),
	2147:  uint16(sym_if_expression),
	2148:  uint16(sym_while_expression),
	2149:  uint16(sym_loop_expression),
	2150:  uint16(sym_const_block),
	2151:  uint16(sym_break_expression),
	2152:  uint16(sym_continue_expression),
	2153:  uint16(sym_block),
	2154:  uint16(sym__literal),
	2155:  uint16(sym_boolean_literal),
	2156:  uint16(25),
	2157:  uint16(3),
	2158:  uint16(1),
	2159:  uint16(sym_comment),
	2160:  uint16(7),
	2161:  uint16(1),
	2162:  uint16(sym_identifier),
	2163:  uint16(9),
	2164:  uint16(1),
	2165:  uint16(anon_sym_COLON_COLON),
	2166:  uint16(11),
	2167:  uint16(1),
	2168:  uint16(anon_sym_LBRACE),
	2169:  uint16(19),
	2170:  uint16(1),
	2171:  uint16(anon_sym_LPAREN),
	2172:  uint16(27),
	2173:  uint16(1),
	2174:  uint16(anon_sym_return),
	2175:  uint16(29),
	2176:  uint16(1),
	2177:  uint16(anon_sym_if),
	2178:  uint16(31),
	2179:  uint16(1),
	2180:  uint16(anon_sym_while),
	2181:  uint16(33),
	2182:  uint16(1),
	2183:  uint16(anon_sym_loop),
	2184:  uint16(35),
	2185:  uint16(1),
	2186:  uint16(anon_sym_const),
	2187:  uint16(37),
	2188:  uint16(1),
	2189:  uint16(anon_sym_break),
	2190:  uint16(39),
	2191:  uint16(1),
	2192:  uint16(anon_sym_continue),
	2193:  uint16(41),
	2194:  uint16(1),
	2195:  uint16(sym_integer_literal),
	2196:  uint16(43),
	2197:  uint16(1),
	2198:  uint16(sym_float_literal),
	2199:  uint16(120),
	2200:  uint16(1),
	2201:  uint16(anon_sym_LT),
	2202:  uint16(124),
	2203:  uint16(1),
	2204:  uint16(anon_sym_AMP),
	2205:  uint16(126),
	2206:  uint16(1),
	2207:  uint16(anon_sym_BANG),
	2208:  uint16(218),
	2209:  uint16(1),
	2210:  uint16(sym_mutable_specifier),
	2211:  uint16(82),
	2212:  uint16(1),
	2213:  uint16(sym_field_expression),
	2214:  uint16(105),
	2215:  uint16(1),
	2216:  uint16(sym__expression),
	2217:  uint16(340),
	2218:  uint16(1),
	2219:  uint16(sym_scoped_identifier),
	2220:  uint16(45),
	2221:  uint16(2),
	2222:  uint16(anon_sym_true),
	2223:  uint16(anon_sym_false),
	2224:  uint16(329),
	2225:  uint16(2),
	2226:  uint16(sym_bracketed_type),
	2227:  uint16(sym__path),
	2228:  uint16(21),
	2229:  uint16(5),
	2230:  uint16(anon_sym_u8),
	2231:  uint16(anon_sym_u64),
	2232:  uint16(anon_sym_u128),
	2233:  uint16(anon_sym_bool),
	2234:  uint16(anon_sym_address),
	2235:  uint16(83),
	2236:  uint16(22),
	2237:  uint16(sym_generic_function),
	2238:  uint16(sym__expression_except_range),
	2239:  uint16(sym__expression_ending_with_block),
	2240:  uint16(sym_unary_expression),
	2241:  uint16(sym_reference_expression),
	2242:  uint16(sym_binary_expression),
	2243:  uint16(sym_assignment_expression),
	2244:  uint16(sym_type_cast_expression),
	2245:  uint16(sym_return_expression),
	2246:  uint16(sym_call_expression),
	2247:  uint16(sym_parenthesized_expression),
	2248:  uint16(sym_tuple_expression),
	2249:  uint16(sym_unit_expression),
	2250:  uint16(sym_if_expression),
	2251:  uint16(sym_while_expression),
	2252:  uint16(sym_loop_expression),
	2253:  uint16(sym_const_block),
	2254:  uint16(sym_break_expression),
	2255:  uint16(sym_continue_expression),
	2256:  uint16(sym_block),
	2257:  uint16(sym__literal),
	2258:  uint16(sym_boolean_literal),
	2259:  uint16(25),
	2260:  uint16(3),
	2261:  uint16(1),
	2262:  uint16(sym_comment),
	2263:  uint16(7),
	2264:  uint16(1),
	2265:  uint16(sym_identifier),
	2266:  uint16(9),
	2267:  uint16(1),
	2268:  uint16(anon_sym_COLON_COLON),
	2269:  uint16(11),
	2270:  uint16(1),
	2271:  uint16(anon_sym_LBRACE),
	2272:  uint16(19),
	2273:  uint16(1),
	2274:  uint16(anon_sym_LPAREN),
	2275:  uint16(27),
	2276:  uint16(1),
	2277:  uint16(anon_sym_return),
	2278:  uint16(29),
	2279:  uint16(1),
	2280:  uint16(anon_sym_if),
	2281:  uint16(31),
	2282:  uint16(1),
	2283:  uint16(anon_sym_while),
	2284:  uint16(33),
	2285:  uint16(1),
	2286:  uint16(anon_sym_loop),
	2287:  uint16(35),
	2288:  uint16(1),
	2289:  uint16(anon_sym_const),
	2290:  uint16(37),
	2291:  uint16(1),
	2292:  uint16(anon_sym_break),
	2293:  uint16(39),
	2294:  uint16(1),
	2295:  uint16(anon_sym_continue),
	2296:  uint16(41),
	2297:  uint16(1),
	2298:  uint16(sym_integer_literal),
	2299:  uint16(43),
	2300:  uint16(1),
	2301:  uint16(sym_float_literal),
	2302:  uint16(120),
	2303:  uint16(1),
	2304:  uint16(anon_sym_LT),
	2305:  uint16(124),
	2306:  uint16(1),
	2307:  uint16(anon_sym_AMP),
	2308:  uint16(126),
	2309:  uint16(1),
	2310:  uint16(anon_sym_BANG),
	2311:  uint16(220),
	2312:  uint16(1),
	2313:  uint16(anon_sym_RPAREN),
	2314:  uint16(82),
	2315:  uint16(1),
	2316:  uint16(sym_field_expression),
	2317:  uint16(146),
	2318:  uint16(1),
	2319:  uint16(sym__expression),
	2320:  uint16(340),
	2321:  uint16(1),
	2322:  uint16(sym_scoped_identifier),
	2323:  uint16(45),
	2324:  uint16(2),
	2325:  uint16(anon_sym_true),
	2326:  uint16(anon_sym_false),
	2327:  uint16(329),
	2328:  uint16(2),
	2329:  uint16(sym_bracketed_type),
	2330:  uint16(sym__path),
	2331:  uint16(21),
	2332:  uint16(5),
	2333:  uint16(anon_sym_u8),
	2334:  uint16(anon_sym_u64),
	2335:  uint16(anon_sym_u128),
	2336:  uint16(anon_sym_bool),
	2337:  uint16(anon_sym_address),
	2338:  uint16(83),
	2339:  uint16(22),
	2340:  uint16(sym_generic_function),
	2341:  uint16(sym__expression_except_range),
	2342:  uint16(sym__expression_ending_with_block),
	2343:  uint16(sym_unary_expression),
	2344:  uint16(sym_reference_expression),
	2345:  uint16(sym_binary_expression),
	2346:  uint16(sym_assignment_expression),
	2347:  uint16(sym_type_cast_expression),
	2348:  uint16(sym_return_expression),
	2349:  uint16(sym_call_expression),
	2350:  uint16(sym_parenthesized_expression),
	2351:  uint16(sym_tuple_expression),
	2352:  uint16(sym_unit_expression),
	2353:  uint16(sym_if_expression),
	2354:  uint16(sym_while_expression),
	2355:  uint16(sym_loop_expression),
	2356:  uint16(sym_const_block),
	2357:  uint16(sym_break_expression),
	2358:  uint16(sym_continue_expression),
	2359:  uint16(sym_block),
	2360:  uint16(sym__literal),
	2361:  uint16(sym_boolean_literal),
	2362:  uint16(25),
	2363:  uint16(3),
	2364:  uint16(1),
	2365:  uint16(sym_comment),
	2366:  uint16(7),
	2367:  uint16(1),
	2368:  uint16(sym_identifier),
	2369:  uint16(9),
	2370:  uint16(1),
	2371:  uint16(anon_sym_COLON_COLON),
	2372:  uint16(11),
	2373:  uint16(1),
	2374:  uint16(anon_sym_LBRACE),
	2375:  uint16(19),
	2376:  uint16(1),
	2377:  uint16(anon_sym_LPAREN),
	2378:  uint16(27),
	2379:  uint16(1),
	2380:  uint16(anon_sym_return),
	2381:  uint16(29),
	2382:  uint16(1),
	2383:  uint16(anon_sym_if),
	2384:  uint16(31),
	2385:  uint16(1),
	2386:  uint16(anon_sym_while),
	2387:  uint16(33),
	2388:  uint16(1),
	2389:  uint16(anon_sym_loop),
	2390:  uint16(35),
	2391:  uint16(1),
	2392:  uint16(anon_sym_const),
	2393:  uint16(37),
	2394:  uint16(1),
	2395:  uint16(anon_sym_break),
	2396:  uint16(39),
	2397:  uint16(1),
	2398:  uint16(anon_sym_continue),
	2399:  uint16(41),
	2400:  uint16(1),
	2401:  uint16(sym_integer_literal),
	2402:  uint16(43),
	2403:  uint16(1),
	2404:  uint16(sym_float_literal),
	2405:  uint16(120),
	2406:  uint16(1),
	2407:  uint16(anon_sym_LT),
	2408:  uint16(124),
	2409:  uint16(1),
	2410:  uint16(anon_sym_AMP),
	2411:  uint16(126),
	2412:  uint16(1),
	2413:  uint16(anon_sym_BANG),
	2414:  uint16(222),
	2415:  uint16(1),
	2416:  uint16(anon_sym_RPAREN),
	2417:  uint16(82),
	2418:  uint16(1),
	2419:  uint16(sym_field_expression),
	2420:  uint16(146),
	2421:  uint16(1),
	2422:  uint16(sym__expression),
	2423:  uint16(340),
	2424:  uint16(1),
	2425:  uint16(sym_scoped_identifier),
	2426:  uint16(45),
	2427:  uint16(2),
	2428:  uint16(anon_sym_true),
	2429:  uint16(anon_sym_false),
	2430:  uint16(329),
	2431:  uint16(2),
	2432:  uint16(sym_bracketed_type),
	2433:  uint16(sym__path),
	2434:  uint16(21),
	2435:  uint16(5),
	2436:  uint16(anon_sym_u8),
	2437:  uint16(anon_sym_u64),
	2438:  uint16(anon_sym_u128),
	2439:  uint16(anon_sym_bool),
	2440:  uint16(anon_sym_address),
	2441:  uint16(83),
	2442:  uint16(22),
	2443:  uint16(sym_generic_function),
	2444:  uint16(sym__expression_except_range),
	2445:  uint16(sym__expression_ending_with_block),
	2446:  uint16(sym_unary_expression),
	2447:  uint16(sym_reference_expression),
	2448:  uint16(sym_binary_expression),
	2449:  uint16(sym_assignment_expression),
	2450:  uint16(sym_type_cast_expression),
	2451:  uint16(sym_return_expression),
	2452:  uint16(sym_call_expression),
	2453:  uint16(sym_parenthesized_expression),
	2454:  uint16(sym_tuple_expression),
	2455:  uint16(sym_unit_expression),
	2456:  uint16(sym_if_expression),
	2457:  uint16(sym_while_expression),
	2458:  uint16(sym_loop_expression),
	2459:  uint16(sym_const_block),
	2460:  uint16(sym_break_expression),
	2461:  uint16(sym_continue_expression),
	2462:  uint16(sym_block),
	2463:  uint16(sym__literal),
	2464:  uint16(sym_boolean_literal),
	2465:  uint16(24),
	2466:  uint16(3),
	2467:  uint16(1),
	2468:  uint16(sym_comment),
	2469:  uint16(7),
	2470:  uint16(1),
	2471:  uint16(sym_identifier),
	2472:  uint16(9),
	2473:  uint16(1),
	2474:  uint16(anon_sym_COLON_COLON),
	2475:  uint16(11),
	2476:  uint16(1),
	2477:  uint16(anon_sym_LBRACE),
	2478:  uint16(19),
	2479:  uint16(1),
	2480:  uint16(anon_sym_LPAREN),
	2481:  uint16(27),
	2482:  uint16(1),
	2483:  uint16(anon_sym_return),
	2484:  uint16(29),
	2485:  uint16(1),
	2486:  uint16(anon_sym_if),
	2487:  uint16(31),
	2488:  uint16(1),
	2489:  uint16(anon_sym_while),
	2490:  uint16(33),
	2491:  uint16(1),
	2492:  uint16(anon_sym_loop),
	2493:  uint16(35),
	2494:  uint16(1),
	2495:  uint16(anon_sym_const),
	2496:  uint16(37),
	2497:  uint16(1),
	2498:  uint16(anon_sym_break),
	2499:  uint16(39),
	2500:  uint16(1),
	2501:  uint16(anon_sym_continue),
	2502:  uint16(41),
	2503:  uint16(1),
	2504:  uint16(sym_integer_literal),
	2505:  uint16(43),
	2506:  uint16(1),
	2507:  uint16(sym_float_literal),
	2508:  uint16(120),
	2509:  uint16(1),
	2510:  uint16(anon_sym_LT),
	2511:  uint16(124),
	2512:  uint16(1),
	2513:  uint16(anon_sym_AMP),
	2514:  uint16(126),
	2515:  uint16(1),
	2516:  uint16(anon_sym_BANG),
	2517:  uint16(82),
	2518:  uint16(1),
	2519:  uint16(sym_field_expression),
	2520:  uint16(104),
	2521:  uint16(1),
	2522:  uint16(sym__expression),
	2523:  uint16(340),
	2524:  uint16(1),
	2525:  uint16(sym_scoped_identifier),
	2526:  uint16(45),
	2527:  uint16(2),
	2528:  uint16(anon_sym_true),
	2529:  uint16(anon_sym_false),
	2530:  uint16(329),
	2531:  uint16(2),
	2532:  uint16(sym_bracketed_type),
	2533:  uint16(sym__path),
	2534:  uint16(21),
	2535:  uint16(5),
	2536:  uint16(anon_sym_u8),
	2537:  uint16(anon_sym_u64),
	2538:  uint16(anon_sym_u128),
	2539:  uint16(anon_sym_bool),
	2540:  uint16(anon_sym_address),
	2541:  uint16(83),
	2542:  uint16(22),
	2543:  uint16(sym_generic_function),
	2544:  uint16(sym__expression_except_range),
	2545:  uint16(sym__expression_ending_with_block),
	2546:  uint16(sym_unary_expression),
	2547:  uint16(sym_reference_expression),
	2548:  uint16(sym_binary_expression),
	2549:  uint16(sym_assignment_expression),
	2550:  uint16(sym_type_cast_expression),
	2551:  uint16(sym_return_expression),
	2552:  uint16(sym_call_expression),
	2553:  uint16(sym_parenthesized_expression),
	2554:  uint16(sym_tuple_expression),
	2555:  uint16(sym_unit_expression),
	2556:  uint16(sym_if_expression),
	2557:  uint16(sym_while_expression),
	2558:  uint16(sym_loop_expression),
	2559:  uint16(sym_const_block),
	2560:  uint16(sym_break_expression),
	2561:  uint16(sym_continue_expression),
	2562:  uint16(sym_block),
	2563:  uint16(sym__literal),
	2564:  uint16(sym_boolean_literal),
	2565:  uint16(24),
	2566:  uint16(3),
	2567:  uint16(1),
	2568:  uint16(sym_comment),
	2569:  uint16(7),
	2570:  uint16(1),
	2571:  uint16(sym_identifier),
	2572:  uint16(9),
	2573:  uint16(1),
	2574:  uint16(anon_sym_COLON_COLON),
	2575:  uint16(11),
	2576:  uint16(1),
	2577:  uint16(anon_sym_LBRACE),
	2578:  uint16(19),
	2579:  uint16(1),
	2580:  uint16(anon_sym_LPAREN),
	2581:  uint16(27),
	2582:  uint16(1),
	2583:  uint16(anon_sym_return),
	2584:  uint16(29),
	2585:  uint16(1),
	2586:  uint16(anon_sym_if),
	2587:  uint16(31),
	2588:  uint16(1),
	2589:  uint16(anon_sym_while),
	2590:  uint16(33),
	2591:  uint16(1),
	2592:  uint16(anon_sym_loop),
	2593:  uint16(35),
	2594:  uint16(1),
	2595:  uint16(anon_sym_const),
	2596:  uint16(37),
	2597:  uint16(1),
	2598:  uint16(anon_sym_break),
	2599:  uint16(39),
	2600:  uint16(1),
	2601:  uint16(anon_sym_continue),
	2602:  uint16(41),
	2603:  uint16(1),
	2604:  uint16(sym_integer_literal),
	2605:  uint16(43),
	2606:  uint16(1),
	2607:  uint16(sym_float_literal),
	2608:  uint16(120),
	2609:  uint16(1),
	2610:  uint16(anon_sym_LT),
	2611:  uint16(124),
	2612:  uint16(1),
	2613:  uint16(anon_sym_AMP),
	2614:  uint16(126),
	2615:  uint16(1),
	2616:  uint16(anon_sym_BANG),
	2617:  uint16(82),
	2618:  uint16(1),
	2619:  uint16(sym_field_expression),
	2620:  uint16(120),
	2621:  uint16(1),
	2622:  uint16(sym__expression),
	2623:  uint16(340),
	2624:  uint16(1),
	2625:  uint16(sym_scoped_identifier),
	2626:  uint16(45),
	2627:  uint16(2),
	2628:  uint16(anon_sym_true),
	2629:  uint16(anon_sym_false),
	2630:  uint16(329),
	2631:  uint16(2),
	2632:  uint16(sym_bracketed_type),
	2633:  uint16(sym__path),
	2634:  uint16(21),
	2635:  uint16(5),
	2636:  uint16(anon_sym_u8),
	2637:  uint16(anon_sym_u64),
	2638:  uint16(anon_sym_u128),
	2639:  uint16(anon_sym_bool),
	2640:  uint16(anon_sym_address),
	2641:  uint16(83),
	2642:  uint16(22),
	2643:  uint16(sym_generic_function),
	2644:  uint16(sym__expression_except_range),
	2645:  uint16(sym__expression_ending_with_block),
	2646:  uint16(sym_unary_expression),
	2647:  uint16(sym_reference_expression),
	2648:  uint16(sym_binary_expression),
	2649:  uint16(sym_assignment_expression),
	2650:  uint16(sym_type_cast_expression),
	2651:  uint16(sym_return_expression),
	2652:  uint16(sym_call_expression),
	2653:  uint16(sym_parenthesized_expression),
	2654:  uint16(sym_tuple_expression),
	2655:  uint16(sym_unit_expression),
	2656:  uint16(sym_if_expression),
	2657:  uint16(sym_while_expression),
	2658:  uint16(sym_loop_expression),
	2659:  uint16(sym_const_block),
	2660:  uint16(sym_break_expression),
	2661:  uint16(sym_continue_expression),
	2662:  uint16(sym_block),
	2663:  uint16(sym__literal),
	2664:  uint16(sym_boolean_literal),
	2665:  uint16(24),
	2666:  uint16(3),
	2667:  uint16(1),
	2668:  uint16(sym_comment),
	2669:  uint16(7),
	2670:  uint16(1),
	2671:  uint16(sym_identifier),
	2672:  uint16(9),
	2673:  uint16(1),
	2674:  uint16(anon_sym_COLON_COLON),
	2675:  uint16(11),
	2676:  uint16(1),
	2677:  uint16(anon_sym_LBRACE),
	2678:  uint16(19),
	2679:  uint16(1),
	2680:  uint16(anon_sym_LPAREN),
	2681:  uint16(27),
	2682:  uint16(1),
	2683:  uint16(anon_sym_return),
	2684:  uint16(29),
	2685:  uint16(1),
	2686:  uint16(anon_sym_if),
	2687:  uint16(31),
	2688:  uint16(1),
	2689:  uint16(anon_sym_while),
	2690:  uint16(33),
	2691:  uint16(1),
	2692:  uint16(anon_sym_loop),
	2693:  uint16(35),
	2694:  uint16(1),
	2695:  uint16(anon_sym_const),
	2696:  uint16(37),
	2697:  uint16(1),
	2698:  uint16(anon_sym_break),
	2699:  uint16(39),
	2700:  uint16(1),
	2701:  uint16(anon_sym_continue),
	2702:  uint16(41),
	2703:  uint16(1),
	2704:  uint16(sym_integer_literal),
	2705:  uint16(43),
	2706:  uint16(1),
	2707:  uint16(sym_float_literal),
	2708:  uint16(120),
	2709:  uint16(1),
	2710:  uint16(anon_sym_LT),
	2711:  uint16(124),
	2712:  uint16(1),
	2713:  uint16(anon_sym_AMP),
	2714:  uint16(126),
	2715:  uint16(1),
	2716:  uint16(anon_sym_BANG),
	2717:  uint16(82),
	2718:  uint16(1),
	2719:  uint16(sym_field_expression),
	2720:  uint16(146),
	2721:  uint16(1),
	2722:  uint16(sym__expression),
	2723:  uint16(340),
	2724:  uint16(1),
	2725:  uint16(sym_scoped_identifier),
	2726:  uint16(45),
	2727:  uint16(2),
	2728:  uint16(anon_sym_true),
	2729:  uint16(anon_sym_false),
	2730:  uint16(329),
	2731:  uint16(2),
	2732:  uint16(sym_bracketed_type),
	2733:  uint16(sym__path),
	2734:  uint16(21),
	2735:  uint16(5),
	2736:  uint16(anon_sym_u8),
	2737:  uint16(anon_sym_u64),
	2738:  uint16(anon_sym_u128),
	2739:  uint16(anon_sym_bool),
	2740:  uint16(anon_sym_address),
	2741:  uint16(83),
	2742:  uint16(22),
	2743:  uint16(sym_generic_function),
	2744:  uint16(sym__expression_except_range),
	2745:  uint16(sym__expression_ending_with_block),
	2746:  uint16(sym_unary_expression),
	2747:  uint16(sym_reference_expression),
	2748:  uint16(sym_binary_expression),
	2749:  uint16(sym_assignment_expression),
	2750:  uint16(sym_type_cast_expression),
	2751:  uint16(sym_return_expression),
	2752:  uint16(sym_call_expression),
	2753:  uint16(sym_parenthesized_expression),
	2754:  uint16(sym_tuple_expression),
	2755:  uint16(sym_unit_expression),
	2756:  uint16(sym_if_expression),
	2757:  uint16(sym_while_expression),
	2758:  uint16(sym_loop_expression),
	2759:  uint16(sym_const_block),
	2760:  uint16(sym_break_expression),
	2761:  uint16(sym_continue_expression),
	2762:  uint16(sym_block),
	2763:  uint16(sym__literal),
	2764:  uint16(sym_boolean_literal),
	2765:  uint16(24),
	2766:  uint16(3),
	2767:  uint16(1),
	2768:  uint16(sym_comment),
	2769:  uint16(7),
	2770:  uint16(1),
	2771:  uint16(sym_identifier),
	2772:  uint16(9),
	2773:  uint16(1),
	2774:  uint16(anon_sym_COLON_COLON),
	2775:  uint16(11),
	2776:  uint16(1),
	2777:  uint16(anon_sym_LBRACE),
	2778:  uint16(19),
	2779:  uint16(1),
	2780:  uint16(anon_sym_LPAREN),
	2781:  uint16(27),
	2782:  uint16(1),
	2783:  uint16(anon_sym_return),
	2784:  uint16(29),
	2785:  uint16(1),
	2786:  uint16(anon_sym_if),
	2787:  uint16(31),
	2788:  uint16(1),
	2789:  uint16(anon_sym_while),
	2790:  uint16(33),
	2791:  uint16(1),
	2792:  uint16(anon_sym_loop),
	2793:  uint16(35),
	2794:  uint16(1),
	2795:  uint16(anon_sym_const),
	2796:  uint16(37),
	2797:  uint16(1),
	2798:  uint16(anon_sym_break),
	2799:  uint16(39),
	2800:  uint16(1),
	2801:  uint16(anon_sym_continue),
	2802:  uint16(41),
	2803:  uint16(1),
	2804:  uint16(sym_integer_literal),
	2805:  uint16(43),
	2806:  uint16(1),
	2807:  uint16(sym_float_literal),
	2808:  uint16(120),
	2809:  uint16(1),
	2810:  uint16(anon_sym_LT),
	2811:  uint16(124),
	2812:  uint16(1),
	2813:  uint16(anon_sym_AMP),
	2814:  uint16(126),
	2815:  uint16(1),
	2816:  uint16(anon_sym_BANG),
	2817:  uint16(82),
	2818:  uint16(1),
	2819:  uint16(sym_field_expression),
	2820:  uint16(148),
	2821:  uint16(1),
	2822:  uint16(sym__expression),
	2823:  uint16(340),
	2824:  uint16(1),
	2825:  uint16(sym_scoped_identifier),
	2826:  uint16(45),
	2827:  uint16(2),
	2828:  uint16(anon_sym_true),
	2829:  uint16(anon_sym_false),
	2830:  uint16(329),
	2831:  uint16(2),
	2832:  uint16(sym_bracketed_type),
	2833:  uint16(sym__path),
	2834:  uint16(21),
	2835:  uint16(5),
	2836:  uint16(anon_sym_u8),
	2837:  uint16(anon_sym_u64),
	2838:  uint16(anon_sym_u128),
	2839:  uint16(anon_sym_bool),
	2840:  uint16(anon_sym_address),
	2841:  uint16(83),
	2842:  uint16(22),
	2843:  uint16(sym_generic_function),
	2844:  uint16(sym__expression_except_range),
	2845:  uint16(sym__expression_ending_with_block),
	2846:  uint16(sym_unary_expression),
	2847:  uint16(sym_reference_expression),
	2848:  uint16(sym_binary_expression),
	2849:  uint16(sym_assignment_expression),
	2850:  uint16(sym_type_cast_expression),
	2851:  uint16(sym_return_expression),
	2852:  uint16(sym_call_expression),
	2853:  uint16(sym_parenthesized_expression),
	2854:  uint16(sym_tuple_expression),
	2855:  uint16(sym_unit_expression),
	2856:  uint16(sym_if_expression),
	2857:  uint16(sym_while_expression),
	2858:  uint16(sym_loop_expression),
	2859:  uint16(sym_const_block),
	2860:  uint16(sym_break_expression),
	2861:  uint16(sym_continue_expression),
	2862:  uint16(sym_block),
	2863:  uint16(sym__literal),
	2864:  uint16(sym_boolean_literal),
	2865:  uint16(24),
	2866:  uint16(3),
	2867:  uint16(1),
	2868:  uint16(sym_comment),
	2869:  uint16(7),
	2870:  uint16(1),
	2871:  uint16(sym_identifier),
	2872:  uint16(9),
	2873:  uint16(1),
	2874:  uint16(anon_sym_COLON_COLON),
	2875:  uint16(11),
	2876:  uint16(1),
	2877:  uint16(anon_sym_LBRACE),
	2878:  uint16(19),
	2879:  uint16(1),
	2880:  uint16(anon_sym_LPAREN),
	2881:  uint16(27),
	2882:  uint16(1),
	2883:  uint16(anon_sym_return),
	2884:  uint16(29),
	2885:  uint16(1),
	2886:  uint16(anon_sym_if),
	2887:  uint16(31),
	2888:  uint16(1),
	2889:  uint16(anon_sym_while),
	2890:  uint16(33),
	2891:  uint16(1),
	2892:  uint16(anon_sym_loop),
	2893:  uint16(35),
	2894:  uint16(1),
	2895:  uint16(anon_sym_const),
	2896:  uint16(37),
	2897:  uint16(1),
	2898:  uint16(anon_sym_break),
	2899:  uint16(39),
	2900:  uint16(1),
	2901:  uint16(anon_sym_continue),
	2902:  uint16(41),
	2903:  uint16(1),
	2904:  uint16(sym_integer_literal),
	2905:  uint16(43),
	2906:  uint16(1),
	2907:  uint16(sym_float_literal),
	2908:  uint16(120),
	2909:  uint16(1),
	2910:  uint16(anon_sym_LT),
	2911:  uint16(124),
	2912:  uint16(1),
	2913:  uint16(anon_sym_AMP),
	2914:  uint16(126),
	2915:  uint16(1),
	2916:  uint16(anon_sym_BANG),
	2917:  uint16(82),
	2918:  uint16(1),
	2919:  uint16(sym_field_expression),
	2920:  uint16(117),
	2921:  uint16(1),
	2922:  uint16(sym__expression),
	2923:  uint16(340),
	2924:  uint16(1),
	2925:  uint16(sym_scoped_identifier),
	2926:  uint16(45),
	2927:  uint16(2),
	2928:  uint16(anon_sym_true),
	2929:  uint16(anon_sym_false),
	2930:  uint16(329),
	2931:  uint16(2),
	2932:  uint16(sym_bracketed_type),
	2933:  uint16(sym__path),
	2934:  uint16(21),
	2935:  uint16(5),
	2936:  uint16(anon_sym_u8),
	2937:  uint16(anon_sym_u64),
	2938:  uint16(anon_sym_u128),
	2939:  uint16(anon_sym_bool),
	2940:  uint16(anon_sym_address),
	2941:  uint16(83),
	2942:  uint16(22),
	2943:  uint16(sym_generic_function),
	2944:  uint16(sym__expression_except_range),
	2945:  uint16(sym__expression_ending_with_block),
	2946:  uint16(sym_unary_expression),
	2947:  uint16(sym_reference_expression),
	2948:  uint16(sym_binary_expression),
	2949:  uint16(sym_assignment_expression),
	2950:  uint16(sym_type_cast_expression),
	2951:  uint16(sym_return_expression),
	2952:  uint16(sym_call_expression),
	2953:  uint16(sym_parenthesized_expression),
	2954:  uint16(sym_tuple_expression),
	2955:  uint16(sym_unit_expression),
	2956:  uint16(sym_if_expression),
	2957:  uint16(sym_while_expression),
	2958:  uint16(sym_loop_expression),
	2959:  uint16(sym_const_block),
	2960:  uint16(sym_break_expression),
	2961:  uint16(sym_continue_expression),
	2962:  uint16(sym_block),
	2963:  uint16(sym__literal),
	2964:  uint16(sym_boolean_literal),
	2965:  uint16(24),
	2966:  uint16(3),
	2967:  uint16(1),
	2968:  uint16(sym_comment),
	2969:  uint16(7),
	2970:  uint16(1),
	2971:  uint16(sym_identifier),
	2972:  uint16(9),
	2973:  uint16(1),
	2974:  uint16(anon_sym_COLON_COLON),
	2975:  uint16(11),
	2976:  uint16(1),
	2977:  uint16(anon_sym_LBRACE),
	2978:  uint16(19),
	2979:  uint16(1),
	2980:  uint16(anon_sym_LPAREN),
	2981:  uint16(27),
	2982:  uint16(1),
	2983:  uint16(anon_sym_return),
	2984:  uint16(29),
	2985:  uint16(1),
	2986:  uint16(anon_sym_if),
	2987:  uint16(31),
	2988:  uint16(1),
	2989:  uint16(anon_sym_while),
	2990:  uint16(33),
	2991:  uint16(1),
	2992:  uint16(anon_sym_loop),
	2993:  uint16(35),
	2994:  uint16(1),
	2995:  uint16(anon_sym_const),
	2996:  uint16(37),
	2997:  uint16(1),
	2998:  uint16(anon_sym_break),
	2999:  uint16(39),
	3000:  uint16(1),
	3001:  uint16(anon_sym_continue),
	3002:  uint16(41),
	3003:  uint16(1),
	3004:  uint16(sym_integer_literal),
	3005:  uint16(43),
	3006:  uint16(1),
	3007:  uint16(sym_float_literal),
	3008:  uint16(120),
	3009:  uint16(1),
	3010:  uint16(anon_sym_LT),
	3011:  uint16(124),
	3012:  uint16(1),
	3013:  uint16(anon_sym_AMP),
	3014:  uint16(126),
	3015:  uint16(1),
	3016:  uint16(anon_sym_BANG),
	3017:  uint16(82),
	3018:  uint16(1),
	3019:  uint16(sym_field_expression),
	3020:  uint16(113),
	3021:  uint16(1),
	3022:  uint16(sym__expression),
	3023:  uint16(340),
	3024:  uint16(1),
	3025:  uint16(sym_scoped_identifier),
	3026:  uint16(45),
	3027:  uint16(2),
	3028:  uint16(anon_sym_true),
	3029:  uint16(anon_sym_false),
	3030:  uint16(329),
	3031:  uint16(2),
	3032:  uint16(sym_bracketed_type),
	3033:  uint16(sym__path),
	3034:  uint16(21),
	3035:  uint16(5),
	3036:  uint16(anon_sym_u8),
	3037:  uint16(anon_sym_u64),
	3038:  uint16(anon_sym_u128),
	3039:  uint16(anon_sym_bool),
	3040:  uint16(anon_sym_address),
	3041:  uint16(83),
	3042:  uint16(22),
	3043:  uint16(sym_generic_function),
	3044:  uint16(sym__expression_except_range),
	3045:  uint16(sym__expression_ending_with_block),
	3046:  uint16(sym_unary_expression),
	3047:  uint16(sym_reference_expression),
	3048:  uint16(sym_binary_expression),
	3049:  uint16(sym_assignment_expression),
	3050:  uint16(sym_type_cast_expression),
	3051:  uint16(sym_return_expression),
	3052:  uint16(sym_call_expression),
	3053:  uint16(sym_parenthesized_expression),
	3054:  uint16(sym_tuple_expression),
	3055:  uint16(sym_unit_expression),
	3056:  uint16(sym_if_expression),
	3057:  uint16(sym_while_expression),
	3058:  uint16(sym_loop_expression),
	3059:  uint16(sym_const_block),
	3060:  uint16(sym_break_expression),
	3061:  uint16(sym_continue_expression),
	3062:  uint16(sym_block),
	3063:  uint16(sym__literal),
	3064:  uint16(sym_boolean_literal),
	3065:  uint16(24),
	3066:  uint16(3),
	3067:  uint16(1),
	3068:  uint16(sym_comment),
	3069:  uint16(7),
	3070:  uint16(1),
	3071:  uint16(sym_identifier),
	3072:  uint16(9),
	3073:  uint16(1),
	3074:  uint16(anon_sym_COLON_COLON),
	3075:  uint16(11),
	3076:  uint16(1),
	3077:  uint16(anon_sym_LBRACE),
	3078:  uint16(19),
	3079:  uint16(1),
	3080:  uint16(anon_sym_LPAREN),
	3081:  uint16(27),
	3082:  uint16(1),
	3083:  uint16(anon_sym_return),
	3084:  uint16(29),
	3085:  uint16(1),
	3086:  uint16(anon_sym_if),
	3087:  uint16(31),
	3088:  uint16(1),
	3089:  uint16(anon_sym_while),
	3090:  uint16(33),
	3091:  uint16(1),
	3092:  uint16(anon_sym_loop),
	3093:  uint16(35),
	3094:  uint16(1),
	3095:  uint16(anon_sym_const),
	3096:  uint16(37),
	3097:  uint16(1),
	3098:  uint16(anon_sym_break),
	3099:  uint16(39),
	3100:  uint16(1),
	3101:  uint16(anon_sym_continue),
	3102:  uint16(41),
	3103:  uint16(1),
	3104:  uint16(sym_integer_literal),
	3105:  uint16(43),
	3106:  uint16(1),
	3107:  uint16(sym_float_literal),
	3108:  uint16(120),
	3109:  uint16(1),
	3110:  uint16(anon_sym_LT),
	3111:  uint16(124),
	3112:  uint16(1),
	3113:  uint16(anon_sym_AMP),
	3114:  uint16(126),
	3115:  uint16(1),
	3116:  uint16(anon_sym_BANG),
	3117:  uint16(82),
	3118:  uint16(1),
	3119:  uint16(sym_field_expression),
	3120:  uint16(108),
	3121:  uint16(1),
	3122:  uint16(sym__expression),
	3123:  uint16(340),
	3124:  uint16(1),
	3125:  uint16(sym_scoped_identifier),
	3126:  uint16(45),
	3127:  uint16(2),
	3128:  uint16(anon_sym_true),
	3129:  uint16(anon_sym_false),
	3130:  uint16(329),
	3131:  uint16(2),
	3132:  uint16(sym_bracketed_type),
	3133:  uint16(sym__path),
	3134:  uint16(21),
	3135:  uint16(5),
	3136:  uint16(anon_sym_u8),
	3137:  uint16(anon_sym_u64),
	3138:  uint16(anon_sym_u128),
	3139:  uint16(anon_sym_bool),
	3140:  uint16(anon_sym_address),
	3141:  uint16(83),
	3142:  uint16(22),
	3143:  uint16(sym_generic_function),
	3144:  uint16(sym__expression_except_range),
	3145:  uint16(sym__expression_ending_with_block),
	3146:  uint16(sym_unary_expression),
	3147:  uint16(sym_reference_expression),
	3148:  uint16(sym_binary_expression),
	3149:  uint16(sym_assignment_expression),
	3150:  uint16(sym_type_cast_expression),
	3151:  uint16(sym_return_expression),
	3152:  uint16(sym_call_expression),
	3153:  uint16(sym_parenthesized_expression),
	3154:  uint16(sym_tuple_expression),
	3155:  uint16(sym_unit_expression),
	3156:  uint16(sym_if_expression),
	3157:  uint16(sym_while_expression),
	3158:  uint16(sym_loop_expression),
	3159:  uint16(sym_const_block),
	3160:  uint16(sym_break_expression),
	3161:  uint16(sym_continue_expression),
	3162:  uint16(sym_block),
	3163:  uint16(sym__literal),
	3164:  uint16(sym_boolean_literal),
	3165:  uint16(24),
	3166:  uint16(3),
	3167:  uint16(1),
	3168:  uint16(sym_comment),
	3169:  uint16(7),
	3170:  uint16(1),
	3171:  uint16(sym_identifier),
	3172:  uint16(9),
	3173:  uint16(1),
	3174:  uint16(anon_sym_COLON_COLON),
	3175:  uint16(11),
	3176:  uint16(1),
	3177:  uint16(anon_sym_LBRACE),
	3178:  uint16(19),
	3179:  uint16(1),
	3180:  uint16(anon_sym_LPAREN),
	3181:  uint16(27),
	3182:  uint16(1),
	3183:  uint16(anon_sym_return),
	3184:  uint16(29),
	3185:  uint16(1),
	3186:  uint16(anon_sym_if),
	3187:  uint16(31),
	3188:  uint16(1),
	3189:  uint16(anon_sym_while),
	3190:  uint16(33),
	3191:  uint16(1),
	3192:  uint16(anon_sym_loop),
	3193:  uint16(35),
	3194:  uint16(1),
	3195:  uint16(anon_sym_const),
	3196:  uint16(37),
	3197:  uint16(1),
	3198:  uint16(anon_sym_break),
	3199:  uint16(39),
	3200:  uint16(1),
	3201:  uint16(anon_sym_continue),
	3202:  uint16(41),
	3203:  uint16(1),
	3204:  uint16(sym_integer_literal),
	3205:  uint16(43),
	3206:  uint16(1),
	3207:  uint16(sym_float_literal),
	3208:  uint16(120),
	3209:  uint16(1),
	3210:  uint16(anon_sym_LT),
	3211:  uint16(124),
	3212:  uint16(1),
	3213:  uint16(anon_sym_AMP),
	3214:  uint16(126),
	3215:  uint16(1),
	3216:  uint16(anon_sym_BANG),
	3217:  uint16(82),
	3218:  uint16(1),
	3219:  uint16(sym_field_expression),
	3220:  uint16(150),
	3221:  uint16(1),
	3222:  uint16(sym__expression),
	3223:  uint16(340),
	3224:  uint16(1),
	3225:  uint16(sym_scoped_identifier),
	3226:  uint16(45),
	3227:  uint16(2),
	3228:  uint16(anon_sym_true),
	3229:  uint16(anon_sym_false),
	3230:  uint16(329),
	3231:  uint16(2),
	3232:  uint16(sym_bracketed_type),
	3233:  uint16(sym__path),
	3234:  uint16(21),
	3235:  uint16(5),
	3236:  uint16(anon_sym_u8),
	3237:  uint16(anon_sym_u64),
	3238:  uint16(anon_sym_u128),
	3239:  uint16(anon_sym_bool),
	3240:  uint16(anon_sym_address),
	3241:  uint16(83),
	3242:  uint16(22),
	3243:  uint16(sym_generic_function),
	3244:  uint16(sym__expression_except_range),
	3245:  uint16(sym__expression_ending_with_block),
	3246:  uint16(sym_unary_expression),
	3247:  uint16(sym_reference_expression),
	3248:  uint16(sym_binary_expression),
	3249:  uint16(sym_assignment_expression),
	3250:  uint16(sym_type_cast_expression),
	3251:  uint16(sym_return_expression),
	3252:  uint16(sym_call_expression),
	3253:  uint16(sym_parenthesized_expression),
	3254:  uint16(sym_tuple_expression),
	3255:  uint16(sym_unit_expression),
	3256:  uint16(sym_if_expression),
	3257:  uint16(sym_while_expression),
	3258:  uint16(sym_loop_expression),
	3259:  uint16(sym_const_block),
	3260:  uint16(sym_break_expression),
	3261:  uint16(sym_continue_expression),
	3262:  uint16(sym_block),
	3263:  uint16(sym__literal),
	3264:  uint16(sym_boolean_literal),
	3265:  uint16(24),
	3266:  uint16(3),
	3267:  uint16(1),
	3268:  uint16(sym_comment),
	3269:  uint16(7),
	3270:  uint16(1),
	3271:  uint16(sym_identifier),
	3272:  uint16(9),
	3273:  uint16(1),
	3274:  uint16(anon_sym_COLON_COLON),
	3275:  uint16(11),
	3276:  uint16(1),
	3277:  uint16(anon_sym_LBRACE),
	3278:  uint16(19),
	3279:  uint16(1),
	3280:  uint16(anon_sym_LPAREN),
	3281:  uint16(27),
	3282:  uint16(1),
	3283:  uint16(anon_sym_return),
	3284:  uint16(29),
	3285:  uint16(1),
	3286:  uint16(anon_sym_if),
	3287:  uint16(31),
	3288:  uint16(1),
	3289:  uint16(anon_sym_while),
	3290:  uint16(33),
	3291:  uint16(1),
	3292:  uint16(anon_sym_loop),
	3293:  uint16(35),
	3294:  uint16(1),
	3295:  uint16(anon_sym_const),
	3296:  uint16(37),
	3297:  uint16(1),
	3298:  uint16(anon_sym_break),
	3299:  uint16(39),
	3300:  uint16(1),
	3301:  uint16(anon_sym_continue),
	3302:  uint16(41),
	3303:  uint16(1),
	3304:  uint16(sym_integer_literal),
	3305:  uint16(43),
	3306:  uint16(1),
	3307:  uint16(sym_float_literal),
	3308:  uint16(120),
	3309:  uint16(1),
	3310:  uint16(anon_sym_LT),
	3311:  uint16(124),
	3312:  uint16(1),
	3313:  uint16(anon_sym_AMP),
	3314:  uint16(126),
	3315:  uint16(1),
	3316:  uint16(anon_sym_BANG),
	3317:  uint16(82),
	3318:  uint16(1),
	3319:  uint16(sym_field_expression),
	3320:  uint16(115),
	3321:  uint16(1),
	3322:  uint16(sym__expression),
	3323:  uint16(340),
	3324:  uint16(1),
	3325:  uint16(sym_scoped_identifier),
	3326:  uint16(45),
	3327:  uint16(2),
	3328:  uint16(anon_sym_true),
	3329:  uint16(anon_sym_false),
	3330:  uint16(329),
	3331:  uint16(2),
	3332:  uint16(sym_bracketed_type),
	3333:  uint16(sym__path),
	3334:  uint16(21),
	3335:  uint16(5),
	3336:  uint16(anon_sym_u8),
	3337:  uint16(anon_sym_u64),
	3338:  uint16(anon_sym_u128),
	3339:  uint16(anon_sym_bool),
	3340:  uint16(anon_sym_address),
	3341:  uint16(83),
	3342:  uint16(22),
	3343:  uint16(sym_generic_function),
	3344:  uint16(sym__expression_except_range),
	3345:  uint16(sym__expression_ending_with_block),
	3346:  uint16(sym_unary_expression),
	3347:  uint16(sym_reference_expression),
	3348:  uint16(sym_binary_expression),
	3349:  uint16(sym_assignment_expression),
	3350:  uint16(sym_type_cast_expression),
	3351:  uint16(sym_return_expression),
	3352:  uint16(sym_call_expression),
	3353:  uint16(sym_parenthesized_expression),
	3354:  uint16(sym_tuple_expression),
	3355:  uint16(sym_unit_expression),
	3356:  uint16(sym_if_expression),
	3357:  uint16(sym_while_expression),
	3358:  uint16(sym_loop_expression),
	3359:  uint16(sym_const_block),
	3360:  uint16(sym_break_expression),
	3361:  uint16(sym_continue_expression),
	3362:  uint16(sym_block),
	3363:  uint16(sym__literal),
	3364:  uint16(sym_boolean_literal),
	3365:  uint16(24),
	3366:  uint16(3),
	3367:  uint16(1),
	3368:  uint16(sym_comment),
	3369:  uint16(7),
	3370:  uint16(1),
	3371:  uint16(sym_identifier),
	3372:  uint16(9),
	3373:  uint16(1),
	3374:  uint16(anon_sym_COLON_COLON),
	3375:  uint16(11),
	3376:  uint16(1),
	3377:  uint16(anon_sym_LBRACE),
	3378:  uint16(19),
	3379:  uint16(1),
	3380:  uint16(anon_sym_LPAREN),
	3381:  uint16(27),
	3382:  uint16(1),
	3383:  uint16(anon_sym_return),
	3384:  uint16(29),
	3385:  uint16(1),
	3386:  uint16(anon_sym_if),
	3387:  uint16(31),
	3388:  uint16(1),
	3389:  uint16(anon_sym_while),
	3390:  uint16(33),
	3391:  uint16(1),
	3392:  uint16(anon_sym_loop),
	3393:  uint16(35),
	3394:  uint16(1),
	3395:  uint16(anon_sym_const),
	3396:  uint16(37),
	3397:  uint16(1),
	3398:  uint16(anon_sym_break),
	3399:  uint16(39),
	3400:  uint16(1),
	3401:  uint16(anon_sym_continue),
	3402:  uint16(41),
	3403:  uint16(1),
	3404:  uint16(sym_integer_literal),
	3405:  uint16(43),
	3406:  uint16(1),
	3407:  uint16(sym_float_literal),
	3408:  uint16(120),
	3409:  uint16(1),
	3410:  uint16(anon_sym_LT),
	3411:  uint16(124),
	3412:  uint16(1),
	3413:  uint16(anon_sym_AMP),
	3414:  uint16(126),
	3415:  uint16(1),
	3416:  uint16(anon_sym_BANG),
	3417:  uint16(82),
	3418:  uint16(1),
	3419:  uint16(sym_field_expression),
	3420:  uint16(136),
	3421:  uint16(1),
	3422:  uint16(sym__expression),
	3423:  uint16(340),
	3424:  uint16(1),
	3425:  uint16(sym_scoped_identifier),
	3426:  uint16(45),
	3427:  uint16(2),
	3428:  uint16(anon_sym_true),
	3429:  uint16(anon_sym_false),
	3430:  uint16(329),
	3431:  uint16(2),
	3432:  uint16(sym_bracketed_type),
	3433:  uint16(sym__path),
	3434:  uint16(21),
	3435:  uint16(5),
	3436:  uint16(anon_sym_u8),
	3437:  uint16(anon_sym_u64),
	3438:  uint16(anon_sym_u128),
	3439:  uint16(anon_sym_bool),
	3440:  uint16(anon_sym_address),
	3441:  uint16(83),
	3442:  uint16(22),
	3443:  uint16(sym_generic_function),
	3444:  uint16(sym__expression_except_range),
	3445:  uint16(sym__expression_ending_with_block),
	3446:  uint16(sym_unary_expression),
	3447:  uint16(sym_reference_expression),
	3448:  uint16(sym_binary_expression),
	3449:  uint16(sym_assignment_expression),
	3450:  uint16(sym_type_cast_expression),
	3451:  uint16(sym_return_expression),
	3452:  uint16(sym_call_expression),
	3453:  uint16(sym_parenthesized_expression),
	3454:  uint16(sym_tuple_expression),
	3455:  uint16(sym_unit_expression),
	3456:  uint16(sym_if_expression),
	3457:  uint16(sym_while_expression),
	3458:  uint16(sym_loop_expression),
	3459:  uint16(sym_const_block),
	3460:  uint16(sym_break_expression),
	3461:  uint16(sym_continue_expression),
	3462:  uint16(sym_block),
	3463:  uint16(sym__literal),
	3464:  uint16(sym_boolean_literal),
	3465:  uint16(24),
	3466:  uint16(3),
	3467:  uint16(1),
	3468:  uint16(sym_comment),
	3469:  uint16(7),
	3470:  uint16(1),
	3471:  uint16(sym_identifier),
	3472:  uint16(9),
	3473:  uint16(1),
	3474:  uint16(anon_sym_COLON_COLON),
	3475:  uint16(11),
	3476:  uint16(1),
	3477:  uint16(anon_sym_LBRACE),
	3478:  uint16(19),
	3479:  uint16(1),
	3480:  uint16(anon_sym_LPAREN),
	3481:  uint16(27),
	3482:  uint16(1),
	3483:  uint16(anon_sym_return),
	3484:  uint16(29),
	3485:  uint16(1),
	3486:  uint16(anon_sym_if),
	3487:  uint16(31),
	3488:  uint16(1),
	3489:  uint16(anon_sym_while),
	3490:  uint16(33),
	3491:  uint16(1),
	3492:  uint16(anon_sym_loop),
	3493:  uint16(35),
	3494:  uint16(1),
	3495:  uint16(anon_sym_const),
	3496:  uint16(37),
	3497:  uint16(1),
	3498:  uint16(anon_sym_break),
	3499:  uint16(39),
	3500:  uint16(1),
	3501:  uint16(anon_sym_continue),
	3502:  uint16(41),
	3503:  uint16(1),
	3504:  uint16(sym_integer_literal),
	3505:  uint16(43),
	3506:  uint16(1),
	3507:  uint16(sym_float_literal),
	3508:  uint16(120),
	3509:  uint16(1),
	3510:  uint16(anon_sym_LT),
	3511:  uint16(124),
	3512:  uint16(1),
	3513:  uint16(anon_sym_AMP),
	3514:  uint16(126),
	3515:  uint16(1),
	3516:  uint16(anon_sym_BANG),
	3517:  uint16(82),
	3518:  uint16(1),
	3519:  uint16(sym_field_expression),
	3520:  uint16(135),
	3521:  uint16(1),
	3522:  uint16(sym__expression),
	3523:  uint16(340),
	3524:  uint16(1),
	3525:  uint16(sym_scoped_identifier),
	3526:  uint16(45),
	3527:  uint16(2),
	3528:  uint16(anon_sym_true),
	3529:  uint16(anon_sym_false),
	3530:  uint16(329),
	3531:  uint16(2),
	3532:  uint16(sym_bracketed_type),
	3533:  uint16(sym__path),
	3534:  uint16(21),
	3535:  uint16(5),
	3536:  uint16(anon_sym_u8),
	3537:  uint16(anon_sym_u64),
	3538:  uint16(anon_sym_u128),
	3539:  uint16(anon_sym_bool),
	3540:  uint16(anon_sym_address),
	3541:  uint16(83),
	3542:  uint16(22),
	3543:  uint16(sym_generic_function),
	3544:  uint16(sym__expression_except_range),
	3545:  uint16(sym__expression_ending_with_block),
	3546:  uint16(sym_unary_expression),
	3547:  uint16(sym_reference_expression),
	3548:  uint16(sym_binary_expression),
	3549:  uint16(sym_assignment_expression),
	3550:  uint16(sym_type_cast_expression),
	3551:  uint16(sym_return_expression),
	3552:  uint16(sym_call_expression),
	3553:  uint16(sym_parenthesized_expression),
	3554:  uint16(sym_tuple_expression),
	3555:  uint16(sym_unit_expression),
	3556:  uint16(sym_if_expression),
	3557:  uint16(sym_while_expression),
	3558:  uint16(sym_loop_expression),
	3559:  uint16(sym_const_block),
	3560:  uint16(sym_break_expression),
	3561:  uint16(sym_continue_expression),
	3562:  uint16(sym_block),
	3563:  uint16(sym__literal),
	3564:  uint16(sym_boolean_literal),
	3565:  uint16(24),
	3566:  uint16(3),
	3567:  uint16(1),
	3568:  uint16(sym_comment),
	3569:  uint16(7),
	3570:  uint16(1),
	3571:  uint16(sym_identifier),
	3572:  uint16(9),
	3573:  uint16(1),
	3574:  uint16(anon_sym_COLON_COLON),
	3575:  uint16(11),
	3576:  uint16(1),
	3577:  uint16(anon_sym_LBRACE),
	3578:  uint16(19),
	3579:  uint16(1),
	3580:  uint16(anon_sym_LPAREN),
	3581:  uint16(27),
	3582:  uint16(1),
	3583:  uint16(anon_sym_return),
	3584:  uint16(29),
	3585:  uint16(1),
	3586:  uint16(anon_sym_if),
	3587:  uint16(31),
	3588:  uint16(1),
	3589:  uint16(anon_sym_while),
	3590:  uint16(33),
	3591:  uint16(1),
	3592:  uint16(anon_sym_loop),
	3593:  uint16(35),
	3594:  uint16(1),
	3595:  uint16(anon_sym_const),
	3596:  uint16(37),
	3597:  uint16(1),
	3598:  uint16(anon_sym_break),
	3599:  uint16(39),
	3600:  uint16(1),
	3601:  uint16(anon_sym_continue),
	3602:  uint16(41),
	3603:  uint16(1),
	3604:  uint16(sym_integer_literal),
	3605:  uint16(43),
	3606:  uint16(1),
	3607:  uint16(sym_float_literal),
	3608:  uint16(120),
	3609:  uint16(1),
	3610:  uint16(anon_sym_LT),
	3611:  uint16(124),
	3612:  uint16(1),
	3613:  uint16(anon_sym_AMP),
	3614:  uint16(126),
	3615:  uint16(1),
	3616:  uint16(anon_sym_BANG),
	3617:  uint16(82),
	3618:  uint16(1),
	3619:  uint16(sym_field_expression),
	3620:  uint16(106),
	3621:  uint16(1),
	3622:  uint16(sym__expression),
	3623:  uint16(340),
	3624:  uint16(1),
	3625:  uint16(sym_scoped_identifier),
	3626:  uint16(45),
	3627:  uint16(2),
	3628:  uint16(anon_sym_true),
	3629:  uint16(anon_sym_false),
	3630:  uint16(329),
	3631:  uint16(2),
	3632:  uint16(sym_bracketed_type),
	3633:  uint16(sym__path),
	3634:  uint16(21),
	3635:  uint16(5),
	3636:  uint16(anon_sym_u8),
	3637:  uint16(anon_sym_u64),
	3638:  uint16(anon_sym_u128),
	3639:  uint16(anon_sym_bool),
	3640:  uint16(anon_sym_address),
	3641:  uint16(83),
	3642:  uint16(22),
	3643:  uint16(sym_generic_function),
	3644:  uint16(sym__expression_except_range),
	3645:  uint16(sym__expression_ending_with_block),
	3646:  uint16(sym_unary_expression),
	3647:  uint16(sym_reference_expression),
	3648:  uint16(sym_binary_expression),
	3649:  uint16(sym_assignment_expression),
	3650:  uint16(sym_type_cast_expression),
	3651:  uint16(sym_return_expression),
	3652:  uint16(sym_call_expression),
	3653:  uint16(sym_parenthesized_expression),
	3654:  uint16(sym_tuple_expression),
	3655:  uint16(sym_unit_expression),
	3656:  uint16(sym_if_expression),
	3657:  uint16(sym_while_expression),
	3658:  uint16(sym_loop_expression),
	3659:  uint16(sym_const_block),
	3660:  uint16(sym_break_expression),
	3661:  uint16(sym_continue_expression),
	3662:  uint16(sym_block),
	3663:  uint16(sym__literal),
	3664:  uint16(sym_boolean_literal),
	3665:  uint16(24),
	3666:  uint16(3),
	3667:  uint16(1),
	3668:  uint16(sym_comment),
	3669:  uint16(7),
	3670:  uint16(1),
	3671:  uint16(sym_identifier),
	3672:  uint16(9),
	3673:  uint16(1),
	3674:  uint16(anon_sym_COLON_COLON),
	3675:  uint16(11),
	3676:  uint16(1),
	3677:  uint16(anon_sym_LBRACE),
	3678:  uint16(19),
	3679:  uint16(1),
	3680:  uint16(anon_sym_LPAREN),
	3681:  uint16(27),
	3682:  uint16(1),
	3683:  uint16(anon_sym_return),
	3684:  uint16(29),
	3685:  uint16(1),
	3686:  uint16(anon_sym_if),
	3687:  uint16(31),
	3688:  uint16(1),
	3689:  uint16(anon_sym_while),
	3690:  uint16(33),
	3691:  uint16(1),
	3692:  uint16(anon_sym_loop),
	3693:  uint16(35),
	3694:  uint16(1),
	3695:  uint16(anon_sym_const),
	3696:  uint16(37),
	3697:  uint16(1),
	3698:  uint16(anon_sym_break),
	3699:  uint16(39),
	3700:  uint16(1),
	3701:  uint16(anon_sym_continue),
	3702:  uint16(41),
	3703:  uint16(1),
	3704:  uint16(sym_integer_literal),
	3705:  uint16(43),
	3706:  uint16(1),
	3707:  uint16(sym_float_literal),
	3708:  uint16(120),
	3709:  uint16(1),
	3710:  uint16(anon_sym_LT),
	3711:  uint16(124),
	3712:  uint16(1),
	3713:  uint16(anon_sym_AMP),
	3714:  uint16(126),
	3715:  uint16(1),
	3716:  uint16(anon_sym_BANG),
	3717:  uint16(82),
	3718:  uint16(1),
	3719:  uint16(sym_field_expression),
	3720:  uint16(134),
	3721:  uint16(1),
	3722:  uint16(sym__expression),
	3723:  uint16(340),
	3724:  uint16(1),
	3725:  uint16(sym_scoped_identifier),
	3726:  uint16(45),
	3727:  uint16(2),
	3728:  uint16(anon_sym_true),
	3729:  uint16(anon_sym_false),
	3730:  uint16(329),
	3731:  uint16(2),
	3732:  uint16(sym_bracketed_type),
	3733:  uint16(sym__path),
	3734:  uint16(21),
	3735:  uint16(5),
	3736:  uint16(anon_sym_u8),
	3737:  uint16(anon_sym_u64),
	3738:  uint16(anon_sym_u128),
	3739:  uint16(anon_sym_bool),
	3740:  uint16(anon_sym_address),
	3741:  uint16(83),
	3742:  uint16(22),
	3743:  uint16(sym_generic_function),
	3744:  uint16(sym__expression_except_range),
	3745:  uint16(sym__expression_ending_with_block),
	3746:  uint16(sym_unary_expression),
	3747:  uint16(sym_reference_expression),
	3748:  uint16(sym_binary_expression),
	3749:  uint16(sym_assignment_expression),
	3750:  uint16(sym_type_cast_expression),
	3751:  uint16(sym_return_expression),
	3752:  uint16(sym_call_expression),
	3753:  uint16(sym_parenthesized_expression),
	3754:  uint16(sym_tuple_expression),
	3755:  uint16(sym_unit_expression),
	3756:  uint16(sym_if_expression),
	3757:  uint16(sym_while_expression),
	3758:  uint16(sym_loop_expression),
	3759:  uint16(sym_const_block),
	3760:  uint16(sym_break_expression),
	3761:  uint16(sym_continue_expression),
	3762:  uint16(sym_block),
	3763:  uint16(sym__literal),
	3764:  uint16(sym_boolean_literal),
	3765:  uint16(24),
	3766:  uint16(3),
	3767:  uint16(1),
	3768:  uint16(sym_comment),
	3769:  uint16(7),
	3770:  uint16(1),
	3771:  uint16(sym_identifier),
	3772:  uint16(9),
	3773:  uint16(1),
	3774:  uint16(anon_sym_COLON_COLON),
	3775:  uint16(11),
	3776:  uint16(1),
	3777:  uint16(anon_sym_LBRACE),
	3778:  uint16(19),
	3779:  uint16(1),
	3780:  uint16(anon_sym_LPAREN),
	3781:  uint16(27),
	3782:  uint16(1),
	3783:  uint16(anon_sym_return),
	3784:  uint16(29),
	3785:  uint16(1),
	3786:  uint16(anon_sym_if),
	3787:  uint16(31),
	3788:  uint16(1),
	3789:  uint16(anon_sym_while),
	3790:  uint16(33),
	3791:  uint16(1),
	3792:  uint16(anon_sym_loop),
	3793:  uint16(35),
	3794:  uint16(1),
	3795:  uint16(anon_sym_const),
	3796:  uint16(37),
	3797:  uint16(1),
	3798:  uint16(anon_sym_break),
	3799:  uint16(39),
	3800:  uint16(1),
	3801:  uint16(anon_sym_continue),
	3802:  uint16(41),
	3803:  uint16(1),
	3804:  uint16(sym_integer_literal),
	3805:  uint16(43),
	3806:  uint16(1),
	3807:  uint16(sym_float_literal),
	3808:  uint16(120),
	3809:  uint16(1),
	3810:  uint16(anon_sym_LT),
	3811:  uint16(124),
	3812:  uint16(1),
	3813:  uint16(anon_sym_AMP),
	3814:  uint16(126),
	3815:  uint16(1),
	3816:  uint16(anon_sym_BANG),
	3817:  uint16(82),
	3818:  uint16(1),
	3819:  uint16(sym_field_expression),
	3820:  uint16(133),
	3821:  uint16(1),
	3822:  uint16(sym__expression),
	3823:  uint16(340),
	3824:  uint16(1),
	3825:  uint16(sym_scoped_identifier),
	3826:  uint16(45),
	3827:  uint16(2),
	3828:  uint16(anon_sym_true),
	3829:  uint16(anon_sym_false),
	3830:  uint16(329),
	3831:  uint16(2),
	3832:  uint16(sym_bracketed_type),
	3833:  uint16(sym__path),
	3834:  uint16(21),
	3835:  uint16(5),
	3836:  uint16(anon_sym_u8),
	3837:  uint16(anon_sym_u64),
	3838:  uint16(anon_sym_u128),
	3839:  uint16(anon_sym_bool),
	3840:  uint16(anon_sym_address),
	3841:  uint16(83),
	3842:  uint16(22),
	3843:  uint16(sym_generic_function),
	3844:  uint16(sym__expression_except_range),
	3845:  uint16(sym__expression_ending_with_block),
	3846:  uint16(sym_unary_expression),
	3847:  uint16(sym_reference_expression),
	3848:  uint16(sym_binary_expression),
	3849:  uint16(sym_assignment_expression),
	3850:  uint16(sym_type_cast_expression),
	3851:  uint16(sym_return_expression),
	3852:  uint16(sym_call_expression),
	3853:  uint16(sym_parenthesized_expression),
	3854:  uint16(sym_tuple_expression),
	3855:  uint16(sym_unit_expression),
	3856:  uint16(sym_if_expression),
	3857:  uint16(sym_while_expression),
	3858:  uint16(sym_loop_expression),
	3859:  uint16(sym_const_block),
	3860:  uint16(sym_break_expression),
	3861:  uint16(sym_continue_expression),
	3862:  uint16(sym_block),
	3863:  uint16(sym__literal),
	3864:  uint16(sym_boolean_literal),
	3865:  uint16(24),
	3866:  uint16(3),
	3867:  uint16(1),
	3868:  uint16(sym_comment),
	3869:  uint16(7),
	3870:  uint16(1),
	3871:  uint16(sym_identifier),
	3872:  uint16(9),
	3873:  uint16(1),
	3874:  uint16(anon_sym_COLON_COLON),
	3875:  uint16(11),
	3876:  uint16(1),
	3877:  uint16(anon_sym_LBRACE),
	3878:  uint16(19),
	3879:  uint16(1),
	3880:  uint16(anon_sym_LPAREN),
	3881:  uint16(27),
	3882:  uint16(1),
	3883:  uint16(anon_sym_return),
	3884:  uint16(29),
	3885:  uint16(1),
	3886:  uint16(anon_sym_if),
	3887:  uint16(31),
	3888:  uint16(1),
	3889:  uint16(anon_sym_while),
	3890:  uint16(33),
	3891:  uint16(1),
	3892:  uint16(anon_sym_loop),
	3893:  uint16(35),
	3894:  uint16(1),
	3895:  uint16(anon_sym_const),
	3896:  uint16(37),
	3897:  uint16(1),
	3898:  uint16(anon_sym_break),
	3899:  uint16(39),
	3900:  uint16(1),
	3901:  uint16(anon_sym_continue),
	3902:  uint16(41),
	3903:  uint16(1),
	3904:  uint16(sym_integer_literal),
	3905:  uint16(43),
	3906:  uint16(1),
	3907:  uint16(sym_float_literal),
	3908:  uint16(120),
	3909:  uint16(1),
	3910:  uint16(anon_sym_LT),
	3911:  uint16(124),
	3912:  uint16(1),
	3913:  uint16(anon_sym_AMP),
	3914:  uint16(126),
	3915:  uint16(1),
	3916:  uint16(anon_sym_BANG),
	3917:  uint16(82),
	3918:  uint16(1),
	3919:  uint16(sym_field_expression),
	3920:  uint16(132),
	3921:  uint16(1),
	3922:  uint16(sym__expression),
	3923:  uint16(340),
	3924:  uint16(1),
	3925:  uint16(sym_scoped_identifier),
	3926:  uint16(45),
	3927:  uint16(2),
	3928:  uint16(anon_sym_true),
	3929:  uint16(anon_sym_false),
	3930:  uint16(329),
	3931:  uint16(2),
	3932:  uint16(sym_bracketed_type),
	3933:  uint16(sym__path),
	3934:  uint16(21),
	3935:  uint16(5),
	3936:  uint16(anon_sym_u8),
	3937:  uint16(anon_sym_u64),
	3938:  uint16(anon_sym_u128),
	3939:  uint16(anon_sym_bool),
	3940:  uint16(anon_sym_address),
	3941:  uint16(83),
	3942:  uint16(22),
	3943:  uint16(sym_generic_function),
	3944:  uint16(sym__expression_except_range),
	3945:  uint16(sym__expression_ending_with_block),
	3946:  uint16(sym_unary_expression),
	3947:  uint16(sym_reference_expression),
	3948:  uint16(sym_binary_expression),
	3949:  uint16(sym_assignment_expression),
	3950:  uint16(sym_type_cast_expression),
	3951:  uint16(sym_return_expression),
	3952:  uint16(sym_call_expression),
	3953:  uint16(sym_parenthesized_expression),
	3954:  uint16(sym_tuple_expression),
	3955:  uint16(sym_unit_expression),
	3956:  uint16(sym_if_expression),
	3957:  uint16(sym_while_expression),
	3958:  uint16(sym_loop_expression),
	3959:  uint16(sym_const_block),
	3960:  uint16(sym_break_expression),
	3961:  uint16(sym_continue_expression),
	3962:  uint16(sym_block),
	3963:  uint16(sym__literal),
	3964:  uint16(sym_boolean_literal),
	3965:  uint16(5),
	3966:  uint16(3),
	3967:  uint16(1),
	3968:  uint16(sym_comment),
	3969:  uint16(228),
	3970:  uint16(1),
	3971:  uint16(anon_sym_else),
	3972:  uint16(47),
	3973:  uint16(1),
	3974:  uint16(sym_else_clause),
	3975:  uint16(226),
	3976:  uint16(19),
	3977:  uint16(anon_sym_COLON_COLON),
	3978:  uint16(anon_sym_LBRACE),
	3979:  uint16(anon_sym_RBRACE),
	3980:  uint16(anon_sym_SEMI),
	3981:  uint16(anon_sym_LPAREN),
	3982:  uint16(anon_sym_AMP_AMP),
	3983:  uint16(anon_sym_PIPE_PIPE),
	3984:  uint16(anon_sym_CARET),
	3985:  uint16(anon_sym_EQ_EQ),
	3986:  uint16(anon_sym_BANG_EQ),
	3987:  uint16(anon_sym_LT_EQ),
	3988:  uint16(anon_sym_GT_EQ),
	3989:  uint16(anon_sym_LT_LT),
	3990:  uint16(anon_sym_GT_GT),
	3991:  uint16(anon_sym_PLUS),
	3992:  uint16(anon_sym_DASH),
	3993:  uint16(anon_sym_STAR),
	3994:  uint16(anon_sym_PERCENT),
	3995:  uint16(sym_float_literal),
	3996:  uint16(224),
	3997:  uint16(28),
	3998:  uint16(anon_sym_EQ),
	3999:  uint16(anon_sym_fun),
	4000:  uint16(anon_sym_public),
	4001:  uint16(anon_sym_LT),
	4002:  uint16(anon_sym_GT),
	4003:  uint16(anon_sym_use),
	4004:  uint16(anon_sym_u8),
	4005:  uint16(anon_sym_u64),
	4006:  uint16(anon_sym_u128),
	4007:  uint16(anon_sym_bool),
	4008:  uint16(anon_sym_address),
	4009:  uint16(anon_sym_as),
	4010:  uint16(anon_sym_AMP),
	4011:  uint16(anon_sym_BANG),
	4012:  uint16(anon_sym_PIPE),
	4013:  uint16(anon_sym_SLASH),
	4014:  uint16(anon_sym_return),
	4015:  uint16(anon_sym_if),
	4016:  uint16(anon_sym_while),
	4017:  uint16(anon_sym_loop),
	4018:  uint16(anon_sym_const),
	4019:  uint16(anon_sym_break),
	4020:  uint16(anon_sym_continue),
	4021:  uint16(anon_sym_DOT),
	4022:  uint16(sym_integer_literal),
	4023:  uint16(anon_sym_true),
	4024:  uint16(anon_sym_false),
	4025:  uint16(sym_identifier),
	4026:  uint16(3),
	4027:  uint16(3),
	4028:  uint16(1),
	4029:  uint16(sym_comment),
	4030:  uint16(232),
	4031:  uint16(21),
	4032:  uint16(anon_sym_COLON_COLON),
	4033:  uint16(anon_sym_LBRACE),
	4034:  uint16(anon_sym_RBRACE),
	4035:  uint16(anon_sym_SEMI),
	4036:  uint16(anon_sym_COMMA),
	4037:  uint16(anon_sym_LPAREN),
	4038:  uint16(anon_sym_RPAREN),
	4039:  uint16(anon_sym_AMP_AMP),
	4040:  uint16(anon_sym_PIPE_PIPE),
	4041:  uint16(anon_sym_CARET),
	4042:  uint16(anon_sym_EQ_EQ),
	4043:  uint16(anon_sym_BANG_EQ),
	4044:  uint16(anon_sym_LT_EQ),
	4045:  uint16(anon_sym_GT_EQ),
	4046:  uint16(anon_sym_LT_LT),
	4047:  uint16(anon_sym_GT_GT),
	4048:  uint16(anon_sym_PLUS),
	4049:  uint16(anon_sym_DASH),
	4050:  uint16(anon_sym_STAR),
	4051:  uint16(anon_sym_PERCENT),
	4052:  uint16(sym_float_literal),
	4053:  uint16(230),
	4054:  uint16(28),
	4055:  uint16(anon_sym_EQ),
	4056:  uint16(anon_sym_fun),
	4057:  uint16(anon_sym_public),
	4058:  uint16(anon_sym_LT),
	4059:  uint16(anon_sym_GT),
	4060:  uint16(anon_sym_use),
	4061:  uint16(anon_sym_u8),
	4062:  uint16(anon_sym_u64),
	4063:  uint16(anon_sym_u128),
	4064:  uint16(anon_sym_bool),
	4065:  uint16(anon_sym_address),
	4066:  uint16(anon_sym_as),
	4067:  uint16(anon_sym_AMP),
	4068:  uint16(anon_sym_BANG),
	4069:  uint16(anon_sym_PIPE),
	4070:  uint16(anon_sym_SLASH),
	4071:  uint16(anon_sym_return),
	4072:  uint16(anon_sym_if),
	4073:  uint16(anon_sym_while),
	4074:  uint16(anon_sym_loop),
	4075:  uint16(anon_sym_const),
	4076:  uint16(anon_sym_break),
	4077:  uint16(anon_sym_continue),
	4078:  uint16(anon_sym_DOT),
	4079:  uint16(sym_integer_literal),
	4080:  uint16(anon_sym_true),
	4081:  uint16(anon_sym_false),
	4082:  uint16(sym_identifier),
	4083:  uint16(3),
	4084:  uint16(3),
	4085:  uint16(1),
	4086:  uint16(sym_comment),
	4087:  uint16(236),
	4088:  uint16(19),
	4089:  uint16(anon_sym_COLON_COLON),
	4090:  uint16(anon_sym_LBRACE),
	4091:  uint16(anon_sym_RBRACE),
	4092:  uint16(anon_sym_SEMI),
	4093:  uint16(anon_sym_LPAREN),
	4094:  uint16(anon_sym_AMP_AMP),
	4095:  uint16(anon_sym_PIPE_PIPE),
	4096:  uint16(anon_sym_CARET),
	4097:  uint16(anon_sym_EQ_EQ),
	4098:  uint16(anon_sym_BANG_EQ),
	4099:  uint16(anon_sym_LT_EQ),
	4100:  uint16(anon_sym_GT_EQ),
	4101:  uint16(anon_sym_LT_LT),
	4102:  uint16(anon_sym_GT_GT),
	4103:  uint16(anon_sym_PLUS),
	4104:  uint16(anon_sym_DASH),
	4105:  uint16(anon_sym_STAR),
	4106:  uint16(anon_sym_PERCENT),
	4107:  uint16(sym_float_literal),
	4108:  uint16(234),
	4109:  uint16(29),
	4110:  uint16(anon_sym_EQ),
	4111:  uint16(anon_sym_fun),
	4112:  uint16(anon_sym_public),
	4113:  uint16(anon_sym_LT),
	4114:  uint16(anon_sym_GT),
	4115:  uint16(anon_sym_use),
	4116:  uint16(anon_sym_u8),
	4117:  uint16(anon_sym_u64),
	4118:  uint16(anon_sym_u128),
	4119:  uint16(anon_sym_bool),
	4120:  uint16(anon_sym_address),
	4121:  uint16(anon_sym_as),
	4122:  uint16(anon_sym_AMP),
	4123:  uint16(anon_sym_BANG),
	4124:  uint16(anon_sym_PIPE),
	4125:  uint16(anon_sym_SLASH),
	4126:  uint16(anon_sym_return),
	4127:  uint16(anon_sym_if),
	4128:  uint16(anon_sym_else),
	4129:  uint16(anon_sym_while),
	4130:  uint16(anon_sym_loop),
	4131:  uint16(anon_sym_const),
	4132:  uint16(anon_sym_break),
	4133:  uint16(anon_sym_continue),
	4134:  uint16(anon_sym_DOT),
	4135:  uint16(sym_integer_literal),
	4136:  uint16(anon_sym_true),
	4137:  uint16(anon_sym_false),
	4138:  uint16(sym_identifier),
	4139:  uint16(3),
	4140:  uint16(3),
	4141:  uint16(1),
	4142:  uint16(sym_comment),
	4143:  uint16(240),
	4144:  uint16(19),
	4145:  uint16(anon_sym_COLON_COLON),
	4146:  uint16(anon_sym_LBRACE),
	4147:  uint16(anon_sym_RBRACE),
	4148:  uint16(anon_sym_SEMI),
	4149:  uint16(anon_sym_LPAREN),
	4150:  uint16(anon_sym_AMP_AMP),
	4151:  uint16(anon_sym_PIPE_PIPE),
	4152:  uint16(anon_sym_CARET),
	4153:  uint16(anon_sym_EQ_EQ),
	4154:  uint16(anon_sym_BANG_EQ),
	4155:  uint16(anon_sym_LT_EQ),
	4156:  uint16(anon_sym_GT_EQ),
	4157:  uint16(anon_sym_LT_LT),
	4158:  uint16(anon_sym_GT_GT),
	4159:  uint16(anon_sym_PLUS),
	4160:  uint16(anon_sym_DASH),
	4161:  uint16(anon_sym_STAR),
	4162:  uint16(anon_sym_PERCENT),
	4163:  uint16(sym_float_literal),
	4164:  uint16(238),
	4165:  uint16(29),
	4166:  uint16(anon_sym_EQ),
	4167:  uint16(anon_sym_fun),
	4168:  uint16(anon_sym_public),
	4169:  uint16(anon_sym_LT),
	4170:  uint16(anon_sym_GT),
	4171:  uint16(anon_sym_use),
	4172:  uint16(anon_sym_u8),
	4173:  uint16(anon_sym_u64),
	4174:  uint16(anon_sym_u128),
	4175:  uint16(anon_sym_bool),
	4176:  uint16(anon_sym_address),
	4177:  uint16(anon_sym_as),
	4178:  uint16(anon_sym_AMP),
	4179:  uint16(anon_sym_BANG),
	4180:  uint16(anon_sym_PIPE),
	4181:  uint16(anon_sym_SLASH),
	4182:  uint16(anon_sym_return),
	4183:  uint16(anon_sym_if),
	4184:  uint16(anon_sym_else),
	4185:  uint16(anon_sym_while),
	4186:  uint16(anon_sym_loop),
	4187:  uint16(anon_sym_const),
	4188:  uint16(anon_sym_break),
	4189:  uint16(anon_sym_continue),
	4190:  uint16(anon_sym_DOT),
	4191:  uint16(sym_integer_literal),
	4192:  uint16(anon_sym_true),
	4193:  uint16(anon_sym_false),
	4194:  uint16(sym_identifier),
	4195:  uint16(5),
	4196:  uint16(3),
	4197:  uint16(1),
	4198:  uint16(sym_comment),
	4199:  uint16(244),
	4200:  uint16(5),
	4201:  uint16(anon_sym_COLON_COLON),
	4202:  uint16(anon_sym_LBRACE),
	4203:  uint16(anon_sym_RBRACE),
	4204:  uint16(anon_sym_LPAREN),
	4205:  uint16(sym_float_literal),
	4206:  uint16(248),
	4207:  uint16(6),
	4208:  uint16(anon_sym_EQ),
	4209:  uint16(anon_sym_GT),
	4210:  uint16(anon_sym_as),
	4211:  uint16(anon_sym_PIPE),
	4212:  uint16(anon_sym_SLASH),
	4213:  uint16(anon_sym_DOT),
	4214:  uint16(246),
	4215:  uint16(14),
	4216:  uint16(anon_sym_SEMI),
	4217:  uint16(anon_sym_AMP_AMP),
	4218:  uint16(anon_sym_PIPE_PIPE),
	4219:  uint16(anon_sym_CARET),
	4220:  uint16(anon_sym_EQ_EQ),
	4221:  uint16(anon_sym_BANG_EQ),
	4222:  uint16(anon_sym_LT_EQ),
	4223:  uint16(anon_sym_GT_EQ),
	4224:  uint16(anon_sym_LT_LT),
	4225:  uint16(anon_sym_GT_GT),
	4226:  uint16(anon_sym_PLUS),
	4227:  uint16(anon_sym_DASH),
	4228:  uint16(anon_sym_STAR),
	4229:  uint16(anon_sym_PERCENT),
	4230:  uint16(242),
	4231:  uint16(22),
	4232:  uint16(anon_sym_fun),
	4233:  uint16(anon_sym_public),
	4234:  uint16(anon_sym_LT),
	4235:  uint16(anon_sym_use),
	4236:  uint16(anon_sym_u8),
	4237:  uint16(anon_sym_u64),
	4238:  uint16(anon_sym_u128),
	4239:  uint16(anon_sym_bool),
	4240:  uint16(anon_sym_address),
	4241:  uint16(anon_sym_AMP),
	4242:  uint16(anon_sym_BANG),
	4243:  uint16(anon_sym_return),
	4244:  uint16(anon_sym_if),
	4245:  uint16(anon_sym_while),
	4246:  uint16(anon_sym_loop),
	4247:  uint16(anon_sym_const),
	4248:  uint16(anon_sym_break),
	4249:  uint16(anon_sym_continue),
	4250:  uint16(sym_integer_literal),
	4251:  uint16(anon_sym_true),
	4252:  uint16(anon_sym_false),
	4253:  uint16(sym_identifier),
	4254:  uint16(3),
	4255:  uint16(3),
	4256:  uint16(1),
	4257:  uint16(sym_comment),
	4258:  uint16(252),
	4259:  uint16(19),
	4260:  uint16(anon_sym_COLON_COLON),
	4261:  uint16(anon_sym_LBRACE),
	4262:  uint16(anon_sym_RBRACE),
	4263:  uint16(anon_sym_SEMI),
	4264:  uint16(anon_sym_LPAREN),
	4265:  uint16(anon_sym_AMP_AMP),
	4266:  uint16(anon_sym_PIPE_PIPE),
	4267:  uint16(anon_sym_CARET),
	4268:  uint16(anon_sym_EQ_EQ),
	4269:  uint16(anon_sym_BANG_EQ),
	4270:  uint16(anon_sym_LT_EQ),
	4271:  uint16(anon_sym_GT_EQ),
	4272:  uint16(anon_sym_LT_LT),
	4273:  uint16(anon_sym_GT_GT),
	4274:  uint16(anon_sym_PLUS),
	4275:  uint16(anon_sym_DASH),
	4276:  uint16(anon_sym_STAR),
	4277:  uint16(anon_sym_PERCENT),
	4278:  uint16(sym_float_literal),
	4279:  uint16(250),
	4280:  uint16(28),
	4281:  uint16(anon_sym_EQ),
	4282:  uint16(anon_sym_fun),
	4283:  uint16(anon_sym_public),
	4284:  uint16(anon_sym_LT),
	4285:  uint16(anon_sym_GT),
	4286:  uint16(anon_sym_use),
	4287:  uint16(anon_sym_u8),
	4288:  uint16(anon_sym_u64),
	4289:  uint16(anon_sym_u128),
	4290:  uint16(anon_sym_bool),
	4291:  uint16(anon_sym_address),
	4292:  uint16(anon_sym_as),
	4293:  uint16(anon_sym_AMP),
	4294:  uint16(anon_sym_BANG),
	4295:  uint16(anon_sym_PIPE),
	4296:  uint16(anon_sym_SLASH),
	4297:  uint16(anon_sym_return),
	4298:  uint16(anon_sym_if),
	4299:  uint16(anon_sym_while),
	4300:  uint16(anon_sym_loop),
	4301:  uint16(anon_sym_const),
	4302:  uint16(anon_sym_break),
	4303:  uint16(anon_sym_continue),
	4304:  uint16(anon_sym_DOT),
	4305:  uint16(sym_integer_literal),
	4306:  uint16(anon_sym_true),
	4307:  uint16(anon_sym_false),
	4308:  uint16(sym_identifier),
	4309:  uint16(3),
	4310:  uint16(3),
	4311:  uint16(1),
	4312:  uint16(sym_comment),
	4313:  uint16(256),
	4314:  uint16(19),
	4315:  uint16(anon_sym_COLON_COLON),
	4316:  uint16(anon_sym_LBRACE),
	4317:  uint16(anon_sym_RBRACE),
	4318:  uint16(anon_sym_SEMI),
	4319:  uint16(anon_sym_LPAREN),
	4320:  uint16(anon_sym_AMP_AMP),
	4321:  uint16(anon_sym_PIPE_PIPE),
	4322:  uint16(anon_sym_CARET),
	4323:  uint16(anon_sym_EQ_EQ),
	4324:  uint16(anon_sym_BANG_EQ),
	4325:  uint16(anon_sym_LT_EQ),
	4326:  uint16(anon_sym_GT_EQ),
	4327:  uint16(anon_sym_LT_LT),
	4328:  uint16(anon_sym_GT_GT),
	4329:  uint16(anon_sym_PLUS),
	4330:  uint16(anon_sym_DASH),
	4331:  uint16(anon_sym_STAR),
	4332:  uint16(anon_sym_PERCENT),
	4333:  uint16(sym_float_literal),
	4334:  uint16(254),
	4335:  uint16(28),
	4336:  uint16(anon_sym_EQ),
	4337:  uint16(anon_sym_fun),
	4338:  uint16(anon_sym_public),
	4339:  uint16(anon_sym_LT),
	4340:  uint16(anon_sym_GT),
	4341:  uint16(anon_sym_use),
	4342:  uint16(anon_sym_u8),
	4343:  uint16(anon_sym_u64),
	4344:  uint16(anon_sym_u128),
	4345:  uint16(anon_sym_bool),
	4346:  uint16(anon_sym_address),
	4347:  uint16(anon_sym_as),
	4348:  uint16(anon_sym_AMP),
	4349:  uint16(anon_sym_BANG),
	4350:  uint16(anon_sym_PIPE),
	4351:  uint16(anon_sym_SLASH),
	4352:  uint16(anon_sym_return),
	4353:  uint16(anon_sym_if),
	4354:  uint16(anon_sym_while),
	4355:  uint16(anon_sym_loop),
	4356:  uint16(anon_sym_const),
	4357:  uint16(anon_sym_break),
	4358:  uint16(anon_sym_continue),
	4359:  uint16(anon_sym_DOT),
	4360:  uint16(sym_integer_literal),
	4361:  uint16(anon_sym_true),
	4362:  uint16(anon_sym_false),
	4363:  uint16(sym_identifier),
	4364:  uint16(3),
	4365:  uint16(3),
	4366:  uint16(1),
	4367:  uint16(sym_comment),
	4368:  uint16(260),
	4369:  uint16(19),
	4370:  uint16(anon_sym_COLON_COLON),
	4371:  uint16(anon_sym_LBRACE),
	4372:  uint16(anon_sym_RBRACE),
	4373:  uint16(anon_sym_SEMI),
	4374:  uint16(anon_sym_LPAREN),
	4375:  uint16(anon_sym_AMP_AMP),
	4376:  uint16(anon_sym_PIPE_PIPE),
	4377:  uint16(anon_sym_CARET),
	4378:  uint16(anon_sym_EQ_EQ),
	4379:  uint16(anon_sym_BANG_EQ),
	4380:  uint16(anon_sym_LT_EQ),
	4381:  uint16(anon_sym_GT_EQ),
	4382:  uint16(anon_sym_LT_LT),
	4383:  uint16(anon_sym_GT_GT),
	4384:  uint16(anon_sym_PLUS),
	4385:  uint16(anon_sym_DASH),
	4386:  uint16(anon_sym_STAR),
	4387:  uint16(anon_sym_PERCENT),
	4388:  uint16(sym_float_literal),
	4389:  uint16(258),
	4390:  uint16(28),
	4391:  uint16(anon_sym_EQ),
	4392:  uint16(anon_sym_fun),
	4393:  uint16(anon_sym_public),
	4394:  uint16(anon_sym_LT),
	4395:  uint16(anon_sym_GT),
	4396:  uint16(anon_sym_use),
	4397:  uint16(anon_sym_u8),
	4398:  uint16(anon_sym_u64),
	4399:  uint16(anon_sym_u128),
	4400:  uint16(anon_sym_bool),
	4401:  uint16(anon_sym_address),
	4402:  uint16(anon_sym_as),
	4403:  uint16(anon_sym_AMP),
	4404:  uint16(anon_sym_BANG),
	4405:  uint16(anon_sym_PIPE),
	4406:  uint16(anon_sym_SLASH),
	4407:  uint16(anon_sym_return),
	4408:  uint16(anon_sym_if),
	4409:  uint16(anon_sym_while),
	4410:  uint16(anon_sym_loop),
	4411:  uint16(anon_sym_const),
	4412:  uint16(anon_sym_break),
	4413:  uint16(anon_sym_continue),
	4414:  uint16(anon_sym_DOT),
	4415:  uint16(sym_integer_literal),
	4416:  uint16(anon_sym_true),
	4417:  uint16(anon_sym_false),
	4418:  uint16(sym_identifier),
	4419:  uint16(3),
	4420:  uint16(3),
	4421:  uint16(1),
	4422:  uint16(sym_comment),
	4423:  uint16(264),
	4424:  uint16(19),
	4425:  uint16(anon_sym_COLON_COLON),
	4426:  uint16(anon_sym_LBRACE),
	4427:  uint16(anon_sym_RBRACE),
	4428:  uint16(anon_sym_SEMI),
	4429:  uint16(anon_sym_LPAREN),
	4430:  uint16(anon_sym_AMP_AMP),
	4431:  uint16(anon_sym_PIPE_PIPE),
	4432:  uint16(anon_sym_CARET),
	4433:  uint16(anon_sym_EQ_EQ),
	4434:  uint16(anon_sym_BANG_EQ),
	4435:  uint16(anon_sym_LT_EQ),
	4436:  uint16(anon_sym_GT_EQ),
	4437:  uint16(anon_sym_LT_LT),
	4438:  uint16(anon_sym_GT_GT),
	4439:  uint16(anon_sym_PLUS),
	4440:  uint16(anon_sym_DASH),
	4441:  uint16(anon_sym_STAR),
	4442:  uint16(anon_sym_PERCENT),
	4443:  uint16(sym_float_literal),
	4444:  uint16(262),
	4445:  uint16(28),
	4446:  uint16(anon_sym_EQ),
	4447:  uint16(anon_sym_fun),
	4448:  uint16(anon_sym_public),
	4449:  uint16(anon_sym_LT),
	4450:  uint16(anon_sym_GT),
	4451:  uint16(anon_sym_use),
	4452:  uint16(anon_sym_u8),
	4453:  uint16(anon_sym_u64),
	4454:  uint16(anon_sym_u128),
	4455:  uint16(anon_sym_bool),
	4456:  uint16(anon_sym_address),
	4457:  uint16(anon_sym_as),
	4458:  uint16(anon_sym_AMP),
	4459:  uint16(anon_sym_BANG),
	4460:  uint16(anon_sym_PIPE),
	4461:  uint16(anon_sym_SLASH),
	4462:  uint16(anon_sym_return),
	4463:  uint16(anon_sym_if),
	4464:  uint16(anon_sym_while),
	4465:  uint16(anon_sym_loop),
	4466:  uint16(anon_sym_const),
	4467:  uint16(anon_sym_break),
	4468:  uint16(anon_sym_continue),
	4469:  uint16(anon_sym_DOT),
	4470:  uint16(sym_integer_literal),
	4471:  uint16(anon_sym_true),
	4472:  uint16(anon_sym_false),
	4473:  uint16(sym_identifier),
	4474:  uint16(9),
	4475:  uint16(266),
	4476:  uint16(1),
	4477:  uint16(anon_sym_LBRACE),
	4478:  uint16(271),
	4479:  uint16(1),
	4480:  uint16(anon_sym_LBRACK),
	4481:  uint16(274),
	4482:  uint16(1),
	4483:  uint16(anon_sym_LPAREN),
	4484:  uint16(286),
	4485:  uint16(1),
	4486:  uint16(sym_comment),
	4487:  uint16(283),
	4488:  uint16(2),
	4489:  uint16(anon_sym_as),
	4490:  uint16(aux_sym__non_special_token_token1),
	4491:  uint16(269),
	4492:  uint16(3),
	4493:  uint16(anon_sym_RBRACE),
	4494:  uint16(anon_sym_RBRACK),
	4495:  uint16(anon_sym_RPAREN),
	4496:  uint16(280),
	4497:  uint16(5),
	4498:  uint16(anon_sym_u8),
	4499:  uint16(anon_sym_u64),
	4500:  uint16(anon_sym_u128),
	4501:  uint16(anon_sym_bool),
	4502:  uint16(anon_sym_address),
	4503:  uint16(49),
	4504:  uint16(5),
	4505:  uint16(sym_delim_token_tree),
	4506:  uint16(sym__delim_tokens),
	4507:  uint16(sym__non_delim_token),
	4508:  uint16(sym__non_special_token),
	4509:  uint16(aux_sym_delim_token_tree_repeat1),
	4510:  uint16(277),
	4511:  uint16(28),
	4512:  uint16(anon_sym_use),
	4513:  uint16(anon_sym_return),
	4514:  uint16(anon_sym_if),
	4515:  uint16(anon_sym_let),
	4516:  uint16(anon_sym_while),
	4517:  uint16(anon_sym_loop),
	4518:  uint16(anon_sym_const),
	4519:  uint16(anon_sym_break),
	4520:  uint16(anon_sym_continue),
	4521:  uint16(anon_sym_DOLLAR),
	4522:  uint16(anon_sym_SQUOTE),
	4523:  uint16(anon_sym_async),
	4524:  uint16(anon_sym_await),
	4525:  uint16(anon_sym_default),
	4526:  uint16(anon_sym_enum),
	4527:  uint16(anon_sym_fn),
	4528:  uint16(anon_sym_for),
	4529:  uint16(anon_sym_impl),
	4530:  uint16(anon_sym_match),
	4531:  uint16(anon_sym_mod),
	4532:  uint16(anon_sym_pub),
	4533:  uint16(anon_sym_static),
	4534:  uint16(anon_sym_struct),
	4535:  uint16(anon_sym_trait),
	4536:  uint16(anon_sym_type),
	4537:  uint16(anon_sym_union),
	4538:  uint16(anon_sym_unsafe),
	4539:  uint16(anon_sym_where),
	4540:  uint16(9),
	4541:  uint16(286),
	4542:  uint16(1),
	4543:  uint16(sym_comment),
	4544:  uint16(288),
	4545:  uint16(1),
	4546:  uint16(anon_sym_LBRACE),
	4547:  uint16(290),
	4548:  uint16(1),
	4549:  uint16(anon_sym_RBRACE),
	4550:  uint16(292),
	4551:  uint16(1),
	4552:  uint16(anon_sym_LBRACK),
	4553:  uint16(294),
	4554:  uint16(1),
	4555:  uint16(anon_sym_LPAREN),
	4556:  uint16(300),
	4557:  uint16(2),
	4558:  uint16(anon_sym_as),
	4559:  uint16(aux_sym__non_special_token_token1),
	4560:  uint16(298),
	4561:  uint16(5),
	4562:  uint16(anon_sym_u8),
	4563:  uint16(anon_sym_u64),
	4564:  uint16(anon_sym_u128),
	4565:  uint16(anon_sym_bool),
	4566:  uint16(anon_sym_address),
	4567:  uint16(49),
	4568:  uint16(5),
	4569:  uint16(sym_delim_token_tree),
	4570:  uint16(sym__delim_tokens),
	4571:  uint16(sym__non_delim_token),
	4572:  uint16(sym__non_special_token),
	4573:  uint16(aux_sym_delim_token_tree_repeat1),
	4574:  uint16(296),
	4575:  uint16(28),
	4576:  uint16(anon_sym_use),
	4577:  uint16(anon_sym_return),
	4578:  uint16(anon_sym_if),
	4579:  uint16(anon_sym_let),
	4580:  uint16(anon_sym_while),
	4581:  uint16(anon_sym_loop),
	4582:  uint16(anon_sym_const),
	4583:  uint16(anon_sym_break),
	4584:  uint16(anon_sym_continue),
	4585:  uint16(anon_sym_DOLLAR),
	4586:  uint16(anon_sym_SQUOTE),
	4587:  uint16(anon_sym_async),
	4588:  uint16(anon_sym_await),
	4589:  uint16(anon_sym_default),
	4590:  uint16(anon_sym_enum),
	4591:  uint16(anon_sym_fn),
	4592:  uint16(anon_sym_for),
	4593:  uint16(anon_sym_impl),
	4594:  uint16(anon_sym_match),
	4595:  uint16(anon_sym_mod),
	4596:  uint16(anon_sym_pub),
	4597:  uint16(anon_sym_static),
	4598:  uint16(anon_sym_struct),
	4599:  uint16(anon_sym_trait),
	4600:  uint16(anon_sym_type),
	4601:  uint16(anon_sym_union),
	4602:  uint16(anon_sym_unsafe),
	4603:  uint16(anon_sym_where),
	4604:  uint16(9),
	4605:  uint16(286),
	4606:  uint16(1),
	4607:  uint16(sym_comment),
	4608:  uint16(288),
	4609:  uint16(1),
	4610:  uint16(anon_sym_LBRACE),
	4611:  uint16(292),
	4612:  uint16(1),
	4613:  uint16(anon_sym_LBRACK),
	4614:  uint16(294),
	4615:  uint16(1),
	4616:  uint16(anon_sym_LPAREN),
	4617:  uint16(302),
	4618:  uint16(1),
	4619:  uint16(anon_sym_RBRACE),
	4620:  uint16(306),
	4621:  uint16(2),
	4622:  uint16(anon_sym_as),
	4623:  uint16(aux_sym__non_special_token_token1),
	4624:  uint16(298),
	4625:  uint16(5),
	4626:  uint16(anon_sym_u8),
	4627:  uint16(anon_sym_u64),
	4628:  uint16(anon_sym_u128),
	4629:  uint16(anon_sym_bool),
	4630:  uint16(anon_sym_address),
	4631:  uint16(56),
	4632:  uint16(5),
	4633:  uint16(sym_delim_token_tree),
	4634:  uint16(sym__delim_tokens),
	4635:  uint16(sym__non_delim_token),
	4636:  uint16(sym__non_special_token),
	4637:  uint16(aux_sym_delim_token_tree_repeat1),
	4638:  uint16(304),
	4639:  uint16(28),
	4640:  uint16(anon_sym_use),
	4641:  uint16(anon_sym_return),
	4642:  uint16(anon_sym_if),
	4643:  uint16(anon_sym_let),
	4644:  uint16(anon_sym_while),
	4645:  uint16(anon_sym_loop),
	4646:  uint16(anon_sym_const),
	4647:  uint16(anon_sym_break),
	4648:  uint16(anon_sym_continue),
	4649:  uint16(anon_sym_DOLLAR),
	4650:  uint16(anon_sym_SQUOTE),
	4651:  uint16(anon_sym_async),
	4652:  uint16(anon_sym_await),
	4653:  uint16(anon_sym_default),
	4654:  uint16(anon_sym_enum),
	4655:  uint16(anon_sym_fn),
	4656:  uint16(anon_sym_for),
	4657:  uint16(anon_sym_impl),
	4658:  uint16(anon_sym_match),
	4659:  uint16(anon_sym_mod),
	4660:  uint16(anon_sym_pub),
	4661:  uint16(anon_sym_static),
	4662:  uint16(anon_sym_struct),
	4663:  uint16(anon_sym_trait),
	4664:  uint16(anon_sym_type),
	4665:  uint16(anon_sym_union),
	4666:  uint16(anon_sym_unsafe),
	4667:  uint16(anon_sym_where),
	4668:  uint16(9),
	4669:  uint16(286),
	4670:  uint16(1),
	4671:  uint16(sym_comment),
	4672:  uint16(288),
	4673:  uint16(1),
	4674:  uint16(anon_sym_LBRACE),
	4675:  uint16(290),
	4676:  uint16(1),
	4677:  uint16(anon_sym_RBRACK),
	4678:  uint16(292),
	4679:  uint16(1),
	4680:  uint16(anon_sym_LBRACK),
	4681:  uint16(294),
	4682:  uint16(1),
	4683:  uint16(anon_sym_LPAREN),
	4684:  uint16(300),
	4685:  uint16(2),
	4686:  uint16(anon_sym_as),
	4687:  uint16(aux_sym__non_special_token_token1),
	4688:  uint16(298),
	4689:  uint16(5),
	4690:  uint16(anon_sym_u8),
	4691:  uint16(anon_sym_u64),
	4692:  uint16(anon_sym_u128),
	4693:  uint16(anon_sym_bool),
	4694:  uint16(anon_sym_address),
	4695:  uint16(49),
	4696:  uint16(5),
	4697:  uint16(sym_delim_token_tree),
	4698:  uint16(sym__delim_tokens),
	4699:  uint16(sym__non_delim_token),
	4700:  uint16(sym__non_special_token),
	4701:  uint16(aux_sym_delim_token_tree_repeat1),
	4702:  uint16(296),
	4703:  uint16(28),
	4704:  uint16(anon_sym_use),
	4705:  uint16(anon_sym_return),
	4706:  uint16(anon_sym_if),
	4707:  uint16(anon_sym_let),
	4708:  uint16(anon_sym_while),
	4709:  uint16(anon_sym_loop),
	4710:  uint16(anon_sym_const),
	4711:  uint16(anon_sym_break),
	4712:  uint16(anon_sym_continue),
	4713:  uint16(anon_sym_DOLLAR),
	4714:  uint16(anon_sym_SQUOTE),
	4715:  uint16(anon_sym_async),
	4716:  uint16(anon_sym_await),
	4717:  uint16(anon_sym_default),
	4718:  uint16(anon_sym_enum),
	4719:  uint16(anon_sym_fn),
	4720:  uint16(anon_sym_for),
	4721:  uint16(anon_sym_impl),
	4722:  uint16(anon_sym_match),
	4723:  uint16(anon_sym_mod),
	4724:  uint16(anon_sym_pub),
	4725:  uint16(anon_sym_static),
	4726:  uint16(anon_sym_struct),
	4727:  uint16(anon_sym_trait),
	4728:  uint16(anon_sym_type),
	4729:  uint16(anon_sym_union),
	4730:  uint16(anon_sym_unsafe),
	4731:  uint16(anon_sym_where),
	4732:  uint16(9),
	4733:  uint16(286),
	4734:  uint16(1),
	4735:  uint16(sym_comment),
	4736:  uint16(288),
	4737:  uint16(1),
	4738:  uint16(anon_sym_LBRACE),
	4739:  uint16(290),
	4740:  uint16(1),
	4741:  uint16(anon_sym_RPAREN),
	4742:  uint16(292),
	4743:  uint16(1),
	4744:  uint16(anon_sym_LBRACK),
	4745:  uint16(294),
	4746:  uint16(1),
	4747:  uint16(anon_sym_LPAREN),
	4748:  uint16(300),
	4749:  uint16(2),
	4750:  uint16(anon_sym_as),
	4751:  uint16(aux_sym__non_special_token_token1),
	4752:  uint16(298),
	4753:  uint16(5),
	4754:  uint16(anon_sym_u8),
	4755:  uint16(anon_sym_u64),
	4756:  uint16(anon_sym_u128),
	4757:  uint16(anon_sym_bool),
	4758:  uint16(anon_sym_address),
	4759:  uint16(49),
	4760:  uint16(5),
	4761:  uint16(sym_delim_token_tree),
	4762:  uint16(sym__delim_tokens),
	4763:  uint16(sym__non_delim_token),
	4764:  uint16(sym__non_special_token),
	4765:  uint16(aux_sym_delim_token_tree_repeat1),
	4766:  uint16(296),
	4767:  uint16(28),
	4768:  uint16(anon_sym_use),
	4769:  uint16(anon_sym_return),
	4770:  uint16(anon_sym_if),
	4771:  uint16(anon_sym_let),
	4772:  uint16(anon_sym_while),
	4773:  uint16(anon_sym_loop),
	4774:  uint16(anon_sym_const),
	4775:  uint16(anon_sym_break),
	4776:  uint16(anon_sym_continue),
	4777:  uint16(anon_sym_DOLLAR),
	4778:  uint16(anon_sym_SQUOTE),
	4779:  uint16(anon_sym_async),
	4780:  uint16(anon_sym_await),
	4781:  uint16(anon_sym_default),
	4782:  uint16(anon_sym_enum),
	4783:  uint16(anon_sym_fn),
	4784:  uint16(anon_sym_for),
	4785:  uint16(anon_sym_impl),
	4786:  uint16(anon_sym_match),
	4787:  uint16(anon_sym_mod),
	4788:  uint16(anon_sym_pub),
	4789:  uint16(anon_sym_static),
	4790:  uint16(anon_sym_struct),
	4791:  uint16(anon_sym_trait),
	4792:  uint16(anon_sym_type),
	4793:  uint16(anon_sym_union),
	4794:  uint16(anon_sym_unsafe),
	4795:  uint16(anon_sym_where),
	4796:  uint16(9),
	4797:  uint16(286),
	4798:  uint16(1),
	4799:  uint16(sym_comment),
	4800:  uint16(288),
	4801:  uint16(1),
	4802:  uint16(anon_sym_LBRACE),
	4803:  uint16(292),
	4804:  uint16(1),
	4805:  uint16(anon_sym_LBRACK),
	4806:  uint16(294),
	4807:  uint16(1),
	4808:  uint16(anon_sym_LPAREN),
	4809:  uint16(308),
	4810:  uint16(1),
	4811:  uint16(anon_sym_RBRACK),
	4812:  uint16(300),
	4813:  uint16(2),
	4814:  uint16(anon_sym_as),
	4815:  uint16(aux_sym__non_special_token_token1),
	4816:  uint16(298),
	4817:  uint16(5),
	4818:  uint16(anon_sym_u8),
	4819:  uint16(anon_sym_u64),
	4820:  uint16(anon_sym_u128),
	4821:  uint16(anon_sym_bool),
	4822:  uint16(anon_sym_address),
	4823:  uint16(49),
	4824:  uint16(5),
	4825:  uint16(sym_delim_token_tree),
	4826:  uint16(sym__delim_tokens),
	4827:  uint16(sym__non_delim_token),
	4828:  uint16(sym__non_special_token),
	4829:  uint16(aux_sym_delim_token_tree_repeat1),
	4830:  uint16(296),
	4831:  uint16(28),
	4832:  uint16(anon_sym_use),
	4833:  uint16(anon_sym_return),
	4834:  uint16(anon_sym_if),
	4835:  uint16(anon_sym_let),
	4836:  uint16(anon_sym_while),
	4837:  uint16(anon_sym_loop),
	4838:  uint16(anon_sym_const),
	4839:  uint16(anon_sym_break),
	4840:  uint16(anon_sym_continue),
	4841:  uint16(anon_sym_DOLLAR),
	4842:  uint16(anon_sym_SQUOTE),
	4843:  uint16(anon_sym_async),
	4844:  uint16(anon_sym_await),
	4845:  uint16(anon_sym_default),
	4846:  uint16(anon_sym_enum),
	4847:  uint16(anon_sym_fn),
	4848:  uint16(anon_sym_for),
	4849:  uint16(anon_sym_impl),
	4850:  uint16(anon_sym_match),
	4851:  uint16(anon_sym_mod),
	4852:  uint16(anon_sym_pub),
	4853:  uint16(anon_sym_static),
	4854:  uint16(anon_sym_struct),
	4855:  uint16(anon_sym_trait),
	4856:  uint16(anon_sym_type),
	4857:  uint16(anon_sym_union),
	4858:  uint16(anon_sym_unsafe),
	4859:  uint16(anon_sym_where),
	4860:  uint16(9),
	4861:  uint16(286),
	4862:  uint16(1),
	4863:  uint16(sym_comment),
	4864:  uint16(288),
	4865:  uint16(1),
	4866:  uint16(anon_sym_LBRACE),
	4867:  uint16(292),
	4868:  uint16(1),
	4869:  uint16(anon_sym_LBRACK),
	4870:  uint16(294),
	4871:  uint16(1),
	4872:  uint16(anon_sym_LPAREN),
	4873:  uint16(310),
	4874:  uint16(1),
	4875:  uint16(anon_sym_RBRACE),
	4876:  uint16(314),
	4877:  uint16(2),
	4878:  uint16(anon_sym_as),
	4879:  uint16(aux_sym__non_special_token_token1),
	4880:  uint16(298),
	4881:  uint16(5),
	4882:  uint16(anon_sym_u8),
	4883:  uint16(anon_sym_u64),
	4884:  uint16(anon_sym_u128),
	4885:  uint16(anon_sym_bool),
	4886:  uint16(anon_sym_address),
	4887:  uint16(50),
	4888:  uint16(5),
	4889:  uint16(sym_delim_token_tree),
	4890:  uint16(sym__delim_tokens),
	4891:  uint16(sym__non_delim_token),
	4892:  uint16(sym__non_special_token),
	4893:  uint16(aux_sym_delim_token_tree_repeat1),
	4894:  uint16(312),
	4895:  uint16(28),
	4896:  uint16(anon_sym_use),
	4897:  uint16(anon_sym_return),
	4898:  uint16(anon_sym_if),
	4899:  uint16(anon_sym_let),
	4900:  uint16(anon_sym_while),
	4901:  uint16(anon_sym_loop),
	4902:  uint16(anon_sym_const),
	4903:  uint16(anon_sym_break),
	4904:  uint16(anon_sym_continue),
	4905:  uint16(anon_sym_DOLLAR),
	4906:  uint16(anon_sym_SQUOTE),
	4907:  uint16(anon_sym_async),
	4908:  uint16(anon_sym_await),
	4909:  uint16(anon_sym_default),
	4910:  uint16(anon_sym_enum),
	4911:  uint16(anon_sym_fn),
	4912:  uint16(anon_sym_for),
	4913:  uint16(anon_sym_impl),
	4914:  uint16(anon_sym_match),
	4915:  uint16(anon_sym_mod),
	4916:  uint16(anon_sym_pub),
	4917:  uint16(anon_sym_static),
	4918:  uint16(anon_sym_struct),
	4919:  uint16(anon_sym_trait),
	4920:  uint16(anon_sym_type),
	4921:  uint16(anon_sym_union),
	4922:  uint16(anon_sym_unsafe),
	4923:  uint16(anon_sym_where),
	4924:  uint16(9),
	4925:  uint16(286),
	4926:  uint16(1),
	4927:  uint16(sym_comment),
	4928:  uint16(288),
	4929:  uint16(1),
	4930:  uint16(anon_sym_LBRACE),
	4931:  uint16(292),
	4932:  uint16(1),
	4933:  uint16(anon_sym_LBRACK),
	4934:  uint16(294),
	4935:  uint16(1),
	4936:  uint16(anon_sym_LPAREN),
	4937:  uint16(308),
	4938:  uint16(1),
	4939:  uint16(anon_sym_RBRACE),
	4940:  uint16(300),
	4941:  uint16(2),
	4942:  uint16(anon_sym_as),
	4943:  uint16(aux_sym__non_special_token_token1),
	4944:  uint16(298),
	4945:  uint16(5),
	4946:  uint16(anon_sym_u8),
	4947:  uint16(anon_sym_u64),
	4948:  uint16(anon_sym_u128),
	4949:  uint16(anon_sym_bool),
	4950:  uint16(anon_sym_address),
	4951:  uint16(49),
	4952:  uint16(5),
	4953:  uint16(sym_delim_token_tree),
	4954:  uint16(sym__delim_tokens),
	4955:  uint16(sym__non_delim_token),
	4956:  uint16(sym__non_special_token),
	4957:  uint16(aux_sym_delim_token_tree_repeat1),
	4958:  uint16(296),
	4959:  uint16(28),
	4960:  uint16(anon_sym_use),
	4961:  uint16(anon_sym_return),
	4962:  uint16(anon_sym_if),
	4963:  uint16(anon_sym_let),
	4964:  uint16(anon_sym_while),
	4965:  uint16(anon_sym_loop),
	4966:  uint16(anon_sym_const),
	4967:  uint16(anon_sym_break),
	4968:  uint16(anon_sym_continue),
	4969:  uint16(anon_sym_DOLLAR),
	4970:  uint16(anon_sym_SQUOTE),
	4971:  uint16(anon_sym_async),
	4972:  uint16(anon_sym_await),
	4973:  uint16(anon_sym_default),
	4974:  uint16(anon_sym_enum),
	4975:  uint16(anon_sym_fn),
	4976:  uint16(anon_sym_for),
	4977:  uint16(anon_sym_impl),
	4978:  uint16(anon_sym_match),
	4979:  uint16(anon_sym_mod),
	4980:  uint16(anon_sym_pub),
	4981:  uint16(anon_sym_static),
	4982:  uint16(anon_sym_struct),
	4983:  uint16(anon_sym_trait),
	4984:  uint16(anon_sym_type),
	4985:  uint16(anon_sym_union),
	4986:  uint16(anon_sym_unsafe),
	4987:  uint16(anon_sym_where),
	4988:  uint16(9),
	4989:  uint16(286),
	4990:  uint16(1),
	4991:  uint16(sym_comment),
	4992:  uint16(288),
	4993:  uint16(1),
	4994:  uint16(anon_sym_LBRACE),
	4995:  uint16(292),
	4996:  uint16(1),
	4997:  uint16(anon_sym_LBRACK),
	4998:  uint16(294),
	4999:  uint16(1),
	5000:  uint16(anon_sym_LPAREN),
	5001:  uint16(310),
	5002:  uint16(1),
	5003:  uint16(anon_sym_RBRACK),
	5004:  uint16(318),
	5005:  uint16(2),
	5006:  uint16(anon_sym_as),
	5007:  uint16(aux_sym__non_special_token_token1),
	5008:  uint16(298),
	5009:  uint16(5),
	5010:  uint16(anon_sym_u8),
	5011:  uint16(anon_sym_u64),
	5012:  uint16(anon_sym_u128),
	5013:  uint16(anon_sym_bool),
	5014:  uint16(anon_sym_address),
	5015:  uint16(52),
	5016:  uint16(5),
	5017:  uint16(sym_delim_token_tree),
	5018:  uint16(sym__delim_tokens),
	5019:  uint16(sym__non_delim_token),
	5020:  uint16(sym__non_special_token),
	5021:  uint16(aux_sym_delim_token_tree_repeat1),
	5022:  uint16(316),
	5023:  uint16(28),
	5024:  uint16(anon_sym_use),
	5025:  uint16(anon_sym_return),
	5026:  uint16(anon_sym_if),
	5027:  uint16(anon_sym_let),
	5028:  uint16(anon_sym_while),
	5029:  uint16(anon_sym_loop),
	5030:  uint16(anon_sym_const),
	5031:  uint16(anon_sym_break),
	5032:  uint16(anon_sym_continue),
	5033:  uint16(anon_sym_DOLLAR),
	5034:  uint16(anon_sym_SQUOTE),
	5035:  uint16(anon_sym_async),
	5036:  uint16(anon_sym_await),
	5037:  uint16(anon_sym_default),
	5038:  uint16(anon_sym_enum),
	5039:  uint16(anon_sym_fn),
	5040:  uint16(anon_sym_for),
	5041:  uint16(anon_sym_impl),
	5042:  uint16(anon_sym_match),
	5043:  uint16(anon_sym_mod),
	5044:  uint16(anon_sym_pub),
	5045:  uint16(anon_sym_static),
	5046:  uint16(anon_sym_struct),
	5047:  uint16(anon_sym_trait),
	5048:  uint16(anon_sym_type),
	5049:  uint16(anon_sym_union),
	5050:  uint16(anon_sym_unsafe),
	5051:  uint16(anon_sym_where),
	5052:  uint16(9),
	5053:  uint16(286),
	5054:  uint16(1),
	5055:  uint16(sym_comment),
	5056:  uint16(288),
	5057:  uint16(1),
	5058:  uint16(anon_sym_LBRACE),
	5059:  uint16(292),
	5060:  uint16(1),
	5061:  uint16(anon_sym_LBRACK),
	5062:  uint16(294),
	5063:  uint16(1),
	5064:  uint16(anon_sym_LPAREN),
	5065:  uint16(310),
	5066:  uint16(1),
	5067:  uint16(anon_sym_RPAREN),
	5068:  uint16(322),
	5069:  uint16(2),
	5070:  uint16(anon_sym_as),
	5071:  uint16(aux_sym__non_special_token_token1),
	5072:  uint16(298),
	5073:  uint16(5),
	5074:  uint16(anon_sym_u8),
	5075:  uint16(anon_sym_u64),
	5076:  uint16(anon_sym_u128),
	5077:  uint16(anon_sym_bool),
	5078:  uint16(anon_sym_address),
	5079:  uint16(53),
	5080:  uint16(5),
	5081:  uint16(sym_delim_token_tree),
	5082:  uint16(sym__delim_tokens),
	5083:  uint16(sym__non_delim_token),
	5084:  uint16(sym__non_special_token),
	5085:  uint16(aux_sym_delim_token_tree_repeat1),
	5086:  uint16(320),
	5087:  uint16(28),
	5088:  uint16(anon_sym_use),
	5089:  uint16(anon_sym_return),
	5090:  uint16(anon_sym_if),
	5091:  uint16(anon_sym_let),
	5092:  uint16(anon_sym_while),
	5093:  uint16(anon_sym_loop),
	5094:  uint16(anon_sym_const),
	5095:  uint16(anon_sym_break),
	5096:  uint16(anon_sym_continue),
	5097:  uint16(anon_sym_DOLLAR),
	5098:  uint16(anon_sym_SQUOTE),
	5099:  uint16(anon_sym_async),
	5100:  uint16(anon_sym_await),
	5101:  uint16(anon_sym_default),
	5102:  uint16(anon_sym_enum),
	5103:  uint16(anon_sym_fn),
	5104:  uint16(anon_sym_for),
	5105:  uint16(anon_sym_impl),
	5106:  uint16(anon_sym_match),
	5107:  uint16(anon_sym_mod),
	5108:  uint16(anon_sym_pub),
	5109:  uint16(anon_sym_static),
	5110:  uint16(anon_sym_struct),
	5111:  uint16(anon_sym_trait),
	5112:  uint16(anon_sym_type),
	5113:  uint16(anon_sym_union),
	5114:  uint16(anon_sym_unsafe),
	5115:  uint16(anon_sym_where),
	5116:  uint16(9),
	5117:  uint16(286),
	5118:  uint16(1),
	5119:  uint16(sym_comment),
	5120:  uint16(288),
	5121:  uint16(1),
	5122:  uint16(anon_sym_LBRACE),
	5123:  uint16(292),
	5124:  uint16(1),
	5125:  uint16(anon_sym_LBRACK),
	5126:  uint16(294),
	5127:  uint16(1),
	5128:  uint16(anon_sym_LPAREN),
	5129:  uint16(308),
	5130:  uint16(1),
	5131:  uint16(anon_sym_RPAREN),
	5132:  uint16(300),
	5133:  uint16(2),
	5134:  uint16(anon_sym_as),
	5135:  uint16(aux_sym__non_special_token_token1),
	5136:  uint16(298),
	5137:  uint16(5),
	5138:  uint16(anon_sym_u8),
	5139:  uint16(anon_sym_u64),
	5140:  uint16(anon_sym_u128),
	5141:  uint16(anon_sym_bool),
	5142:  uint16(anon_sym_address),
	5143:  uint16(49),
	5144:  uint16(5),
	5145:  uint16(sym_delim_token_tree),
	5146:  uint16(sym__delim_tokens),
	5147:  uint16(sym__non_delim_token),
	5148:  uint16(sym__non_special_token),
	5149:  uint16(aux_sym_delim_token_tree_repeat1),
	5150:  uint16(296),
	5151:  uint16(28),
	5152:  uint16(anon_sym_use),
	5153:  uint16(anon_sym_return),
	5154:  uint16(anon_sym_if),
	5155:  uint16(anon_sym_let),
	5156:  uint16(anon_sym_while),
	5157:  uint16(anon_sym_loop),
	5158:  uint16(anon_sym_const),
	5159:  uint16(anon_sym_break),
	5160:  uint16(anon_sym_continue),
	5161:  uint16(anon_sym_DOLLAR),
	5162:  uint16(anon_sym_SQUOTE),
	5163:  uint16(anon_sym_async),
	5164:  uint16(anon_sym_await),
	5165:  uint16(anon_sym_default),
	5166:  uint16(anon_sym_enum),
	5167:  uint16(anon_sym_fn),
	5168:  uint16(anon_sym_for),
	5169:  uint16(anon_sym_impl),
	5170:  uint16(anon_sym_match),
	5171:  uint16(anon_sym_mod),
	5172:  uint16(anon_sym_pub),
	5173:  uint16(anon_sym_static),
	5174:  uint16(anon_sym_struct),
	5175:  uint16(anon_sym_trait),
	5176:  uint16(anon_sym_type),
	5177:  uint16(anon_sym_union),
	5178:  uint16(anon_sym_unsafe),
	5179:  uint16(anon_sym_where),
	5180:  uint16(9),
	5181:  uint16(286),
	5182:  uint16(1),
	5183:  uint16(sym_comment),
	5184:  uint16(288),
	5185:  uint16(1),
	5186:  uint16(anon_sym_LBRACE),
	5187:  uint16(292),
	5188:  uint16(1),
	5189:  uint16(anon_sym_LBRACK),
	5190:  uint16(294),
	5191:  uint16(1),
	5192:  uint16(anon_sym_LPAREN),
	5193:  uint16(302),
	5194:  uint16(1),
	5195:  uint16(anon_sym_RPAREN),
	5196:  uint16(326),
	5197:  uint16(2),
	5198:  uint16(anon_sym_as),
	5199:  uint16(aux_sym__non_special_token_token1),
	5200:  uint16(298),
	5201:  uint16(5),
	5202:  uint16(anon_sym_u8),
	5203:  uint16(anon_sym_u64),
	5204:  uint16(anon_sym_u128),
	5205:  uint16(anon_sym_bool),
	5206:  uint16(anon_sym_address),
	5207:  uint16(59),
	5208:  uint16(5),
	5209:  uint16(sym_delim_token_tree),
	5210:  uint16(sym__delim_tokens),
	5211:  uint16(sym__non_delim_token),
	5212:  uint16(sym__non_special_token),
	5213:  uint16(aux_sym_delim_token_tree_repeat1),
	5214:  uint16(324),
	5215:  uint16(28),
	5216:  uint16(anon_sym_use),
	5217:  uint16(anon_sym_return),
	5218:  uint16(anon_sym_if),
	5219:  uint16(anon_sym_let),
	5220:  uint16(anon_sym_while),
	5221:  uint16(anon_sym_loop),
	5222:  uint16(anon_sym_const),
	5223:  uint16(anon_sym_break),
	5224:  uint16(anon_sym_continue),
	5225:  uint16(anon_sym_DOLLAR),
	5226:  uint16(anon_sym_SQUOTE),
	5227:  uint16(anon_sym_async),
	5228:  uint16(anon_sym_await),
	5229:  uint16(anon_sym_default),
	5230:  uint16(anon_sym_enum),
	5231:  uint16(anon_sym_fn),
	5232:  uint16(anon_sym_for),
	5233:  uint16(anon_sym_impl),
	5234:  uint16(anon_sym_match),
	5235:  uint16(anon_sym_mod),
	5236:  uint16(anon_sym_pub),
	5237:  uint16(anon_sym_static),
	5238:  uint16(anon_sym_struct),
	5239:  uint16(anon_sym_trait),
	5240:  uint16(anon_sym_type),
	5241:  uint16(anon_sym_union),
	5242:  uint16(anon_sym_unsafe),
	5243:  uint16(anon_sym_where),
	5244:  uint16(9),
	5245:  uint16(286),
	5246:  uint16(1),
	5247:  uint16(sym_comment),
	5248:  uint16(288),
	5249:  uint16(1),
	5250:  uint16(anon_sym_LBRACE),
	5251:  uint16(292),
	5252:  uint16(1),
	5253:  uint16(anon_sym_LBRACK),
	5254:  uint16(294),
	5255:  uint16(1),
	5256:  uint16(anon_sym_LPAREN),
	5257:  uint16(302),
	5258:  uint16(1),
	5259:  uint16(anon_sym_RBRACK),
	5260:  uint16(330),
	5261:  uint16(2),
	5262:  uint16(anon_sym_as),
	5263:  uint16(aux_sym__non_special_token_token1),
	5264:  uint16(298),
	5265:  uint16(5),
	5266:  uint16(anon_sym_u8),
	5267:  uint16(anon_sym_u64),
	5268:  uint16(anon_sym_u128),
	5269:  uint16(anon_sym_bool),
	5270:  uint16(anon_sym_address),
	5271:  uint16(54),
	5272:  uint16(5),
	5273:  uint16(sym_delim_token_tree),
	5274:  uint16(sym__delim_tokens),
	5275:  uint16(sym__non_delim_token),
	5276:  uint16(sym__non_special_token),
	5277:  uint16(aux_sym_delim_token_tree_repeat1),
	5278:  uint16(328),
	5279:  uint16(28),
	5280:  uint16(anon_sym_use),
	5281:  uint16(anon_sym_return),
	5282:  uint16(anon_sym_if),
	5283:  uint16(anon_sym_let),
	5284:  uint16(anon_sym_while),
	5285:  uint16(anon_sym_loop),
	5286:  uint16(anon_sym_const),
	5287:  uint16(anon_sym_break),
	5288:  uint16(anon_sym_continue),
	5289:  uint16(anon_sym_DOLLAR),
	5290:  uint16(anon_sym_SQUOTE),
	5291:  uint16(anon_sym_async),
	5292:  uint16(anon_sym_await),
	5293:  uint16(anon_sym_default),
	5294:  uint16(anon_sym_enum),
	5295:  uint16(anon_sym_fn),
	5296:  uint16(anon_sym_for),
	5297:  uint16(anon_sym_impl),
	5298:  uint16(anon_sym_match),
	5299:  uint16(anon_sym_mod),
	5300:  uint16(anon_sym_pub),
	5301:  uint16(anon_sym_static),
	5302:  uint16(anon_sym_struct),
	5303:  uint16(anon_sym_trait),
	5304:  uint16(anon_sym_type),
	5305:  uint16(anon_sym_union),
	5306:  uint16(anon_sym_unsafe),
	5307:  uint16(anon_sym_where),
	5308:  uint16(3),
	5309:  uint16(286),
	5310:  uint16(1),
	5311:  uint16(sym_comment),
	5312:  uint16(334),
	5313:  uint16(2),
	5314:  uint16(anon_sym_as),
	5315:  uint16(aux_sym__non_special_token_token1),
	5316:  uint16(332),
	5317:  uint16(39),
	5318:  uint16(anon_sym_LBRACE),
	5319:  uint16(anon_sym_RBRACE),
	5320:  uint16(anon_sym_LBRACK),
	5321:  uint16(anon_sym_RBRACK),
	5322:  uint16(anon_sym_LPAREN),
	5323:  uint16(anon_sym_RPAREN),
	5324:  uint16(anon_sym_use),
	5325:  uint16(anon_sym_u8),
	5326:  uint16(anon_sym_u64),
	5327:  uint16(anon_sym_u128),
	5328:  uint16(anon_sym_bool),
	5329:  uint16(anon_sym_address),
	5330:  uint16(anon_sym_return),
	5331:  uint16(anon_sym_if),
	5332:  uint16(anon_sym_let),
	5333:  uint16(anon_sym_while),
	5334:  uint16(anon_sym_loop),
	5335:  uint16(anon_sym_const),
	5336:  uint16(anon_sym_break),
	5337:  uint16(anon_sym_continue),
	5338:  uint16(anon_sym_DOLLAR),
	5339:  uint16(anon_sym_SQUOTE),
	5340:  uint16(anon_sym_async),
	5341:  uint16(anon_sym_await),
	5342:  uint16(anon_sym_default),
	5343:  uint16(anon_sym_enum),
	5344:  uint16(anon_sym_fn),
	5345:  uint16(anon_sym_for),
	5346:  uint16(anon_sym_impl),
	5347:  uint16(anon_sym_match),
	5348:  uint16(anon_sym_mod),
	5349:  uint16(anon_sym_pub),
	5350:  uint16(anon_sym_static),
	5351:  uint16(anon_sym_struct),
	5352:  uint16(anon_sym_trait),
	5353:  uint16(anon_sym_type),
	5354:  uint16(anon_sym_union),
	5355:  uint16(anon_sym_unsafe),
	5356:  uint16(anon_sym_where),
	5357:  uint16(3),
	5358:  uint16(286),
	5359:  uint16(1),
	5360:  uint16(sym_comment),
	5361:  uint16(338),
	5362:  uint16(2),
	5363:  uint16(anon_sym_as),
	5364:  uint16(aux_sym__non_special_token_token1),
	5365:  uint16(336),
	5366:  uint16(39),
	5367:  uint16(anon_sym_LBRACE),
	5368:  uint16(anon_sym_RBRACE),
	5369:  uint16(anon_sym_LBRACK),
	5370:  uint16(anon_sym_RBRACK),
	5371:  uint16(anon_sym_LPAREN),
	5372:  uint16(anon_sym_RPAREN),
	5373:  uint16(anon_sym_use),
	5374:  uint16(anon_sym_u8),
	5375:  uint16(anon_sym_u64),
	5376:  uint16(anon_sym_u128),
	5377:  uint16(anon_sym_bool),
	5378:  uint16(anon_sym_address),
	5379:  uint16(anon_sym_return),
	5380:  uint16(anon_sym_if),
	5381:  uint16(anon_sym_let),
	5382:  uint16(anon_sym_while),
	5383:  uint16(anon_sym_loop),
	5384:  uint16(anon_sym_const),
	5385:  uint16(anon_sym_break),
	5386:  uint16(anon_sym_continue),
	5387:  uint16(anon_sym_DOLLAR),
	5388:  uint16(anon_sym_SQUOTE),
	5389:  uint16(anon_sym_async),
	5390:  uint16(anon_sym_await),
	5391:  uint16(anon_sym_default),
	5392:  uint16(anon_sym_enum),
	5393:  uint16(anon_sym_fn),
	5394:  uint16(anon_sym_for),
	5395:  uint16(anon_sym_impl),
	5396:  uint16(anon_sym_match),
	5397:  uint16(anon_sym_mod),
	5398:  uint16(anon_sym_pub),
	5399:  uint16(anon_sym_static),
	5400:  uint16(anon_sym_struct),
	5401:  uint16(anon_sym_trait),
	5402:  uint16(anon_sym_type),
	5403:  uint16(anon_sym_union),
	5404:  uint16(anon_sym_unsafe),
	5405:  uint16(anon_sym_where),
	5406:  uint16(3),
	5407:  uint16(286),
	5408:  uint16(1),
	5409:  uint16(sym_comment),
	5410:  uint16(342),
	5411:  uint16(2),
	5412:  uint16(anon_sym_as),
	5413:  uint16(aux_sym__non_special_token_token1),
	5414:  uint16(340),
	5415:  uint16(39),
	5416:  uint16(anon_sym_LBRACE),
	5417:  uint16(anon_sym_RBRACE),
	5418:  uint16(anon_sym_LBRACK),
	5419:  uint16(anon_sym_RBRACK),
	5420:  uint16(anon_sym_LPAREN),
	5421:  uint16(anon_sym_RPAREN),
	5422:  uint16(anon_sym_use),
	5423:  uint16(anon_sym_u8),
	5424:  uint16(anon_sym_u64),
	5425:  uint16(anon_sym_u128),
	5426:  uint16(anon_sym_bool),
	5427:  uint16(anon_sym_address),
	5428:  uint16(anon_sym_return),
	5429:  uint16(anon_sym_if),
	5430:  uint16(anon_sym_let),
	5431:  uint16(anon_sym_while),
	5432:  uint16(anon_sym_loop),
	5433:  uint16(anon_sym_const),
	5434:  uint16(anon_sym_break),
	5435:  uint16(anon_sym_continue),
	5436:  uint16(anon_sym_DOLLAR),
	5437:  uint16(anon_sym_SQUOTE),
	5438:  uint16(anon_sym_async),
	5439:  uint16(anon_sym_await),
	5440:  uint16(anon_sym_default),
	5441:  uint16(anon_sym_enum),
	5442:  uint16(anon_sym_fn),
	5443:  uint16(anon_sym_for),
	5444:  uint16(anon_sym_impl),
	5445:  uint16(anon_sym_match),
	5446:  uint16(anon_sym_mod),
	5447:  uint16(anon_sym_pub),
	5448:  uint16(anon_sym_static),
	5449:  uint16(anon_sym_struct),
	5450:  uint16(anon_sym_trait),
	5451:  uint16(anon_sym_type),
	5452:  uint16(anon_sym_union),
	5453:  uint16(anon_sym_unsafe),
	5454:  uint16(anon_sym_where),
	5455:  uint16(20),
	5456:  uint16(3),
	5457:  uint16(1),
	5458:  uint16(sym_comment),
	5459:  uint16(120),
	5460:  uint16(1),
	5461:  uint16(anon_sym_LT),
	5462:  uint16(344),
	5463:  uint16(1),
	5464:  uint16(sym_identifier),
	5465:  uint16(346),
	5466:  uint16(1),
	5467:  uint16(anon_sym_COLON_COLON),
	5468:  uint16(348),
	5469:  uint16(1),
	5470:  uint16(anon_sym_COMMA),
	5471:  uint16(350),
	5472:  uint16(1),
	5473:  uint16(anon_sym_LPAREN),
	5474:  uint16(354),
	5475:  uint16(1),
	5476:  uint16(anon_sym_RPAREN),
	5477:  uint16(358),
	5478:  uint16(1),
	5479:  uint16(anon_sym_AMP),
	5480:  uint16(360),
	5481:  uint16(1),
	5482:  uint16(sym_mutable_specifier),
	5483:  uint16(362),
	5484:  uint16(1),
	5485:  uint16(anon_sym_DASH),
	5486:  uint16(364),
	5487:  uint16(1),
	5488:  uint16(anon_sym_const),
	5489:  uint16(366),
	5490:  uint16(1),
	5491:  uint16(sym_float_literal),
	5492:  uint16(190),
	5493:  uint16(1),
	5494:  uint16(sym_scoped_identifier),
	5495:  uint16(335),
	5496:  uint16(1),
	5497:  uint16(sym_generic_type),
	5498:  uint16(45),
	5499:  uint16(2),
	5500:  uint16(anon_sym_true),
	5501:  uint16(anon_sym_false),
	5502:  uint16(352),
	5503:  uint16(2),
	5504:  uint16(anon_sym__),
	5505:  uint16(sym_integer_literal),
	5506:  uint16(246),
	5507:  uint16(2),
	5508:  uint16(sym_scoped_type_identifier),
	5509:  uint16(sym__type_identifier),
	5510:  uint16(318),
	5511:  uint16(2),
	5512:  uint16(sym_bracketed_type),
	5513:  uint16(sym__path),
	5514:  uint16(356),
	5515:  uint16(5),
	5516:  uint16(anon_sym_u8),
	5517:  uint16(anon_sym_u64),
	5518:  uint16(anon_sym_u128),
	5519:  uint16(anon_sym_bool),
	5520:  uint16(anon_sym_address),
	5521:  uint16(221),
	5522:  uint16(10),
	5523:  uint16(sym_const_block),
	5524:  uint16(sym__pattern),
	5525:  uint16(sym_tuple_pattern),
	5526:  uint16(sym_struct_pattern),
	5527:  uint16(sym_mut_pattern),
	5528:  uint16(sym_reference_pattern),
	5529:  uint16(sym_or_pattern),
	5530:  uint16(sym__literal_pattern),
	5531:  uint16(sym_negative_literal),
	5532:  uint16(sym_boolean_literal),
	5533:  uint16(19),
	5534:  uint16(3),
	5535:  uint16(1),
	5536:  uint16(sym_comment),
	5537:  uint16(120),
	5538:  uint16(1),
	5539:  uint16(anon_sym_LT),
	5540:  uint16(344),
	5541:  uint16(1),
	5542:  uint16(sym_identifier),
	5543:  uint16(346),
	5544:  uint16(1),
	5545:  uint16(anon_sym_COLON_COLON),
	5546:  uint16(350),
	5547:  uint16(1),
	5548:  uint16(anon_sym_LPAREN),
	5549:  uint16(358),
	5550:  uint16(1),
	5551:  uint16(anon_sym_AMP),
	5552:  uint16(360),
	5553:  uint16(1),
	5554:  uint16(sym_mutable_specifier),
	5555:  uint16(362),
	5556:  uint16(1),
	5557:  uint16(anon_sym_DASH),
	5558:  uint16(364),
	5559:  uint16(1),
	5560:  uint16(anon_sym_const),
	5561:  uint16(370),
	5562:  uint16(1),
	5563:  uint16(anon_sym_RPAREN),
	5564:  uint16(372),
	5565:  uint16(1),
	5566:  uint16(sym_float_literal),
	5567:  uint16(190),
	5568:  uint16(1),
	5569:  uint16(sym_scoped_identifier),
	5570:  uint16(335),
	5571:  uint16(1),
	5572:  uint16(sym_generic_type),
	5573:  uint16(45),
	5574:  uint16(2),
	5575:  uint16(anon_sym_true),
	5576:  uint16(anon_sym_false),
	5577:  uint16(368),
	5578:  uint16(2),
	5579:  uint16(anon_sym__),
	5580:  uint16(sym_integer_literal),
	5581:  uint16(246),
	5582:  uint16(2),
	5583:  uint16(sym_scoped_type_identifier),
	5584:  uint16(sym__type_identifier),
	5585:  uint16(318),
	5586:  uint16(2),
	5587:  uint16(sym_bracketed_type),
	5588:  uint16(sym__path),
	5589:  uint16(356),
	5590:  uint16(5),
	5591:  uint16(anon_sym_u8),
	5592:  uint16(anon_sym_u64),
	5593:  uint16(anon_sym_u128),
	5594:  uint16(anon_sym_bool),
	5595:  uint16(anon_sym_address),
	5596:  uint16(254),
	5597:  uint16(10),
	5598:  uint16(sym_const_block),
	5599:  uint16(sym__pattern),
	5600:  uint16(sym_tuple_pattern),
	5601:  uint16(sym_struct_pattern),
	5602:  uint16(sym_mut_pattern),
	5603:  uint16(sym_reference_pattern),
	5604:  uint16(sym_or_pattern),
	5605:  uint16(sym__literal_pattern),
	5606:  uint16(sym_negative_literal),
	5607:  uint16(sym_boolean_literal),
	5608:  uint16(19),
	5609:  uint16(3),
	5610:  uint16(1),
	5611:  uint16(sym_comment),
	5612:  uint16(120),
	5613:  uint16(1),
	5614:  uint16(anon_sym_LT),
	5615:  uint16(344),
	5616:  uint16(1),
	5617:  uint16(sym_identifier),
	5618:  uint16(346),
	5619:  uint16(1),
	5620:  uint16(anon_sym_COLON_COLON),
	5621:  uint16(350),
	5622:  uint16(1),
	5623:  uint16(anon_sym_LPAREN),
	5624:  uint16(358),
	5625:  uint16(1),
	5626:  uint16(anon_sym_AMP),
	5627:  uint16(360),
	5628:  uint16(1),
	5629:  uint16(sym_mutable_specifier),
	5630:  uint16(362),
	5631:  uint16(1),
	5632:  uint16(anon_sym_DASH),
	5633:  uint16(364),
	5634:  uint16(1),
	5635:  uint16(anon_sym_const),
	5636:  uint16(372),
	5637:  uint16(1),
	5638:  uint16(sym_float_literal),
	5639:  uint16(374),
	5640:  uint16(1),
	5641:  uint16(anon_sym_RPAREN),
	5642:  uint16(190),
	5643:  uint16(1),
	5644:  uint16(sym_scoped_identifier),
	5645:  uint16(335),
	5646:  uint16(1),
	5647:  uint16(sym_generic_type),
	5648:  uint16(45),
	5649:  uint16(2),
	5650:  uint16(anon_sym_true),
	5651:  uint16(anon_sym_false),
	5652:  uint16(368),
	5653:  uint16(2),
	5654:  uint16(anon_sym__),
	5655:  uint16(sym_integer_literal),
	5656:  uint16(246),
	5657:  uint16(2),
	5658:  uint16(sym_scoped_type_identifier),
	5659:  uint16(sym__type_identifier),
	5660:  uint16(318),
	5661:  uint16(2),
	5662:  uint16(sym_bracketed_type),
	5663:  uint16(sym__path),
	5664:  uint16(356),
	5665:  uint16(5),
	5666:  uint16(anon_sym_u8),
	5667:  uint16(anon_sym_u64),
	5668:  uint16(anon_sym_u128),
	5669:  uint16(anon_sym_bool),
	5670:  uint16(anon_sym_address),
	5671:  uint16(254),
	5672:  uint16(10),
	5673:  uint16(sym_const_block),
	5674:  uint16(sym__pattern),
	5675:  uint16(sym_tuple_pattern),
	5676:  uint16(sym_struct_pattern),
	5677:  uint16(sym_mut_pattern),
	5678:  uint16(sym_reference_pattern),
	5679:  uint16(sym_or_pattern),
	5680:  uint16(sym__literal_pattern),
	5681:  uint16(sym_negative_literal),
	5682:  uint16(sym_boolean_literal),
	5683:  uint16(18),
	5684:  uint16(3),
	5685:  uint16(1),
	5686:  uint16(sym_comment),
	5687:  uint16(120),
	5688:  uint16(1),
	5689:  uint16(anon_sym_LT),
	5690:  uint16(344),
	5691:  uint16(1),
	5692:  uint16(sym_identifier),
	5693:  uint16(346),
	5694:  uint16(1),
	5695:  uint16(anon_sym_COLON_COLON),
	5696:  uint16(350),
	5697:  uint16(1),
	5698:  uint16(anon_sym_LPAREN),
	5699:  uint16(358),
	5700:  uint16(1),
	5701:  uint16(anon_sym_AMP),
	5702:  uint16(360),
	5703:  uint16(1),
	5704:  uint16(sym_mutable_specifier),
	5705:  uint16(362),
	5706:  uint16(1),
	5707:  uint16(anon_sym_DASH),
	5708:  uint16(364),
	5709:  uint16(1),
	5710:  uint16(anon_sym_const),
	5711:  uint16(378),
	5712:  uint16(1),
	5713:  uint16(sym_float_literal),
	5714:  uint16(190),
	5715:  uint16(1),
	5716:  uint16(sym_scoped_identifier),
	5717:  uint16(335),
	5718:  uint16(1),
	5719:  uint16(sym_generic_type),
	5720:  uint16(45),
	5721:  uint16(2),
	5722:  uint16(anon_sym_true),
	5723:  uint16(anon_sym_false),
	5724:  uint16(376),
	5725:  uint16(2),
	5726:  uint16(anon_sym__),
	5727:  uint16(sym_integer_literal),
	5728:  uint16(246),
	5729:  uint16(2),
	5730:  uint16(sym_scoped_type_identifier),
	5731:  uint16(sym__type_identifier),
	5732:  uint16(318),
	5733:  uint16(2),
	5734:  uint16(sym_bracketed_type),
	5735:  uint16(sym__path),
	5736:  uint16(356),
	5737:  uint16(5),
	5738:  uint16(anon_sym_u8),
	5739:  uint16(anon_sym_u64),
	5740:  uint16(anon_sym_u128),
	5741:  uint16(anon_sym_bool),
	5742:  uint16(anon_sym_address),
	5743:  uint16(226),
	5744:  uint16(10),
	5745:  uint16(sym_const_block),
	5746:  uint16(sym__pattern),
	5747:  uint16(sym_tuple_pattern),
	5748:  uint16(sym_struct_pattern),
	5749:  uint16(sym_mut_pattern),
	5750:  uint16(sym_reference_pattern),
	5751:  uint16(sym_or_pattern),
	5752:  uint16(sym__literal_pattern),
	5753:  uint16(sym_negative_literal),
	5754:  uint16(sym_boolean_literal),
	5755:  uint16(18),
	5756:  uint16(3),
	5757:  uint16(1),
	5758:  uint16(sym_comment),
	5759:  uint16(120),
	5760:  uint16(1),
	5761:  uint16(anon_sym_LT),
	5762:  uint16(344),
	5763:  uint16(1),
	5764:  uint16(sym_identifier),
	5765:  uint16(346),
	5766:  uint16(1),
	5767:  uint16(anon_sym_COLON_COLON),
	5768:  uint16(350),
	5769:  uint16(1),
	5770:  uint16(anon_sym_LPAREN),
	5771:  uint16(358),
	5772:  uint16(1),
	5773:  uint16(anon_sym_AMP),
	5774:  uint16(360),
	5775:  uint16(1),
	5776:  uint16(sym_mutable_specifier),
	5777:  uint16(362),
	5778:  uint16(1),
	5779:  uint16(anon_sym_DASH),
	5780:  uint16(364),
	5781:  uint16(1),
	5782:  uint16(anon_sym_const),
	5783:  uint16(382),
	5784:  uint16(1),
	5785:  uint16(sym_float_literal),
	5786:  uint16(190),
	5787:  uint16(1),
	5788:  uint16(sym_scoped_identifier),
	5789:  uint16(335),
	5790:  uint16(1),
	5791:  uint16(sym_generic_type),
	5792:  uint16(45),
	5793:  uint16(2),
	5794:  uint16(anon_sym_true),
	5795:  uint16(anon_sym_false),
	5796:  uint16(380),
	5797:  uint16(2),
	5798:  uint16(anon_sym__),
	5799:  uint16(sym_integer_literal),
	5800:  uint16(246),
	5801:  uint16(2),
	5802:  uint16(sym_scoped_type_identifier),
	5803:  uint16(sym__type_identifier),
	5804:  uint16(318),
	5805:  uint16(2),
	5806:  uint16(sym_bracketed_type),
	5807:  uint16(sym__path),
	5808:  uint16(356),
	5809:  uint16(5),
	5810:  uint16(anon_sym_u8),
	5811:  uint16(anon_sym_u64),
	5812:  uint16(anon_sym_u128),
	5813:  uint16(anon_sym_bool),
	5814:  uint16(anon_sym_address),
	5815:  uint16(271),
	5816:  uint16(10),
	5817:  uint16(sym_const_block),
	5818:  uint16(sym__pattern),
	5819:  uint16(sym_tuple_pattern),
	5820:  uint16(sym_struct_pattern),
	5821:  uint16(sym_mut_pattern),
	5822:  uint16(sym_reference_pattern),
	5823:  uint16(sym_or_pattern),
	5824:  uint16(sym__literal_pattern),
	5825:  uint16(sym_negative_literal),
	5826:  uint16(sym_boolean_literal),
	5827:  uint16(18),
	5828:  uint16(3),
	5829:  uint16(1),
	5830:  uint16(sym_comment),
	5831:  uint16(120),
	5832:  uint16(1),
	5833:  uint16(anon_sym_LT),
	5834:  uint16(344),
	5835:  uint16(1),
	5836:  uint16(sym_identifier),
	5837:  uint16(346),
	5838:  uint16(1),
	5839:  uint16(anon_sym_COLON_COLON),
	5840:  uint16(350),
	5841:  uint16(1),
	5842:  uint16(anon_sym_LPAREN),
	5843:  uint16(358),
	5844:  uint16(1),
	5845:  uint16(anon_sym_AMP),
	5846:  uint16(360),
	5847:  uint16(1),
	5848:  uint16(sym_mutable_specifier),
	5849:  uint16(362),
	5850:  uint16(1),
	5851:  uint16(anon_sym_DASH),
	5852:  uint16(364),
	5853:  uint16(1),
	5854:  uint16(anon_sym_const),
	5855:  uint16(386),
	5856:  uint16(1),
	5857:  uint16(sym_float_literal),
	5858:  uint16(190),
	5859:  uint16(1),
	5860:  uint16(sym_scoped_identifier),
	5861:  uint16(335),
	5862:  uint16(1),
	5863:  uint16(sym_generic_type),
	5864:  uint16(45),
	5865:  uint16(2),
	5866:  uint16(anon_sym_true),
	5867:  uint16(anon_sym_false),
	5868:  uint16(384),
	5869:  uint16(2),
	5870:  uint16(anon_sym__),
	5871:  uint16(sym_integer_literal),
	5872:  uint16(246),
	5873:  uint16(2),
	5874:  uint16(sym_scoped_type_identifier),
	5875:  uint16(sym__type_identifier),
	5876:  uint16(318),
	5877:  uint16(2),
	5878:  uint16(sym_bracketed_type),
	5879:  uint16(sym__path),
	5880:  uint16(356),
	5881:  uint16(5),
	5882:  uint16(anon_sym_u8),
	5883:  uint16(anon_sym_u64),
	5884:  uint16(anon_sym_u128),
	5885:  uint16(anon_sym_bool),
	5886:  uint16(anon_sym_address),
	5887:  uint16(204),
	5888:  uint16(10),
	5889:  uint16(sym_const_block),
	5890:  uint16(sym__pattern),
	5891:  uint16(sym_tuple_pattern),
	5892:  uint16(sym_struct_pattern),
	5893:  uint16(sym_mut_pattern),
	5894:  uint16(sym_reference_pattern),
	5895:  uint16(sym_or_pattern),
	5896:  uint16(sym__literal_pattern),
	5897:  uint16(sym_negative_literal),
	5898:  uint16(sym_boolean_literal),
	5899:  uint16(18),
	5900:  uint16(3),
	5901:  uint16(1),
	5902:  uint16(sym_comment),
	5903:  uint16(120),
	5904:  uint16(1),
	5905:  uint16(anon_sym_LT),
	5906:  uint16(344),
	5907:  uint16(1),
	5908:  uint16(sym_identifier),
	5909:  uint16(346),
	5910:  uint16(1),
	5911:  uint16(anon_sym_COLON_COLON),
	5912:  uint16(350),
	5913:  uint16(1),
	5914:  uint16(anon_sym_LPAREN),
	5915:  uint16(358),
	5916:  uint16(1),
	5917:  uint16(anon_sym_AMP),
	5918:  uint16(360),
	5919:  uint16(1),
	5920:  uint16(sym_mutable_specifier),
	5921:  uint16(362),
	5922:  uint16(1),
	5923:  uint16(anon_sym_DASH),
	5924:  uint16(364),
	5925:  uint16(1),
	5926:  uint16(anon_sym_const),
	5927:  uint16(390),
	5928:  uint16(1),
	5929:  uint16(sym_float_literal),
	5930:  uint16(190),
	5931:  uint16(1),
	5932:  uint16(sym_scoped_identifier),
	5933:  uint16(335),
	5934:  uint16(1),
	5935:  uint16(sym_generic_type),
	5936:  uint16(45),
	5937:  uint16(2),
	5938:  uint16(anon_sym_true),
	5939:  uint16(anon_sym_false),
	5940:  uint16(388),
	5941:  uint16(2),
	5942:  uint16(anon_sym__),
	5943:  uint16(sym_integer_literal),
	5944:  uint16(246),
	5945:  uint16(2),
	5946:  uint16(sym_scoped_type_identifier),
	5947:  uint16(sym__type_identifier),
	5948:  uint16(318),
	5949:  uint16(2),
	5950:  uint16(sym_bracketed_type),
	5951:  uint16(sym__path),
	5952:  uint16(356),
	5953:  uint16(5),
	5954:  uint16(anon_sym_u8),
	5955:  uint16(anon_sym_u64),
	5956:  uint16(anon_sym_u128),
	5957:  uint16(anon_sym_bool),
	5958:  uint16(anon_sym_address),
	5959:  uint16(247),
	5960:  uint16(10),
	5961:  uint16(sym_const_block),
	5962:  uint16(sym__pattern),
	5963:  uint16(sym_tuple_pattern),
	5964:  uint16(sym_struct_pattern),
	5965:  uint16(sym_mut_pattern),
	5966:  uint16(sym_reference_pattern),
	5967:  uint16(sym_or_pattern),
	5968:  uint16(sym__literal_pattern),
	5969:  uint16(sym_negative_literal),
	5970:  uint16(sym_boolean_literal),
	5971:  uint16(18),
	5972:  uint16(3),
	5973:  uint16(1),
	5974:  uint16(sym_comment),
	5975:  uint16(120),
	5976:  uint16(1),
	5977:  uint16(anon_sym_LT),
	5978:  uint16(344),
	5979:  uint16(1),
	5980:  uint16(sym_identifier),
	5981:  uint16(346),
	5982:  uint16(1),
	5983:  uint16(anon_sym_COLON_COLON),
	5984:  uint16(350),
	5985:  uint16(1),
	5986:  uint16(anon_sym_LPAREN),
	5987:  uint16(358),
	5988:  uint16(1),
	5989:  uint16(anon_sym_AMP),
	5990:  uint16(360),
	5991:  uint16(1),
	5992:  uint16(sym_mutable_specifier),
	5993:  uint16(362),
	5994:  uint16(1),
	5995:  uint16(anon_sym_DASH),
	5996:  uint16(364),
	5997:  uint16(1),
	5998:  uint16(anon_sym_const),
	5999:  uint16(372),
	6000:  uint16(1),
	6001:  uint16(sym_float_literal),
	6002:  uint16(190),
	6003:  uint16(1),
	6004:  uint16(sym_scoped_identifier),
	6005:  uint16(335),
	6006:  uint16(1),
	6007:  uint16(sym_generic_type),
	6008:  uint16(45),
	6009:  uint16(2),
	6010:  uint16(anon_sym_true),
	6011:  uint16(anon_sym_false),
	6012:  uint16(368),
	6013:  uint16(2),
	6014:  uint16(anon_sym__),
	6015:  uint16(sym_integer_literal),
	6016:  uint16(246),
	6017:  uint16(2),
	6018:  uint16(sym_scoped_type_identifier),
	6019:  uint16(sym__type_identifier),
	6020:  uint16(318),
	6021:  uint16(2),
	6022:  uint16(sym_bracketed_type),
	6023:  uint16(sym__path),
	6024:  uint16(356),
	6025:  uint16(5),
	6026:  uint16(anon_sym_u8),
	6027:  uint16(anon_sym_u64),
	6028:  uint16(anon_sym_u128),
	6029:  uint16(anon_sym_bool),
	6030:  uint16(anon_sym_address),
	6031:  uint16(254),
	6032:  uint16(10),
	6033:  uint16(sym_const_block),
	6034:  uint16(sym__pattern),
	6035:  uint16(sym_tuple_pattern),
	6036:  uint16(sym_struct_pattern),
	6037:  uint16(sym_mut_pattern),
	6038:  uint16(sym_reference_pattern),
	6039:  uint16(sym_or_pattern),
	6040:  uint16(sym__literal_pattern),
	6041:  uint16(sym_negative_literal),
	6042:  uint16(sym_boolean_literal),
	6043:  uint16(18),
	6044:  uint16(3),
	6045:  uint16(1),
	6046:  uint16(sym_comment),
	6047:  uint16(120),
	6048:  uint16(1),
	6049:  uint16(anon_sym_LT),
	6050:  uint16(344),
	6051:  uint16(1),
	6052:  uint16(sym_identifier),
	6053:  uint16(346),
	6054:  uint16(1),
	6055:  uint16(anon_sym_COLON_COLON),
	6056:  uint16(350),
	6057:  uint16(1),
	6058:  uint16(anon_sym_LPAREN),
	6059:  uint16(358),
	6060:  uint16(1),
	6061:  uint16(anon_sym_AMP),
	6062:  uint16(362),
	6063:  uint16(1),
	6064:  uint16(anon_sym_DASH),
	6065:  uint16(364),
	6066:  uint16(1),
	6067:  uint16(anon_sym_const),
	6068:  uint16(394),
	6069:  uint16(1),
	6070:  uint16(sym_mutable_specifier),
	6071:  uint16(396),
	6072:  uint16(1),
	6073:  uint16(sym_float_literal),
	6074:  uint16(190),
	6075:  uint16(1),
	6076:  uint16(sym_scoped_identifier),
	6077:  uint16(335),
	6078:  uint16(1),
	6079:  uint16(sym_generic_type),
	6080:  uint16(45),
	6081:  uint16(2),
	6082:  uint16(anon_sym_true),
	6083:  uint16(anon_sym_false),
	6084:  uint16(392),
	6085:  uint16(2),
	6086:  uint16(anon_sym__),
	6087:  uint16(sym_integer_literal),
	6088:  uint16(246),
	6089:  uint16(2),
	6090:  uint16(sym_scoped_type_identifier),
	6091:  uint16(sym__type_identifier),
	6092:  uint16(318),
	6093:  uint16(2),
	6094:  uint16(sym_bracketed_type),
	6095:  uint16(sym__path),
	6096:  uint16(356),
	6097:  uint16(5),
	6098:  uint16(anon_sym_u8),
	6099:  uint16(anon_sym_u64),
	6100:  uint16(anon_sym_u128),
	6101:  uint16(anon_sym_bool),
	6102:  uint16(anon_sym_address),
	6103:  uint16(202),
	6104:  uint16(10),
	6105:  uint16(sym_const_block),
	6106:  uint16(sym__pattern),
	6107:  uint16(sym_tuple_pattern),
	6108:  uint16(sym_struct_pattern),
	6109:  uint16(sym_mut_pattern),
	6110:  uint16(sym_reference_pattern),
	6111:  uint16(sym_or_pattern),
	6112:  uint16(sym__literal_pattern),
	6113:  uint16(sym_negative_literal),
	6114:  uint16(sym_boolean_literal),
	6115:  uint16(18),
	6116:  uint16(3),
	6117:  uint16(1),
	6118:  uint16(sym_comment),
	6119:  uint16(120),
	6120:  uint16(1),
	6121:  uint16(anon_sym_LT),
	6122:  uint16(344),
	6123:  uint16(1),
	6124:  uint16(sym_identifier),
	6125:  uint16(346),
	6126:  uint16(1),
	6127:  uint16(anon_sym_COLON_COLON),
	6128:  uint16(350),
	6129:  uint16(1),
	6130:  uint16(anon_sym_LPAREN),
	6131:  uint16(358),
	6132:  uint16(1),
	6133:  uint16(anon_sym_AMP),
	6134:  uint16(360),
	6135:  uint16(1),
	6136:  uint16(sym_mutable_specifier),
	6137:  uint16(362),
	6138:  uint16(1),
	6139:  uint16(anon_sym_DASH),
	6140:  uint16(364),
	6141:  uint16(1),
	6142:  uint16(anon_sym_const),
	6143:  uint16(400),
	6144:  uint16(1),
	6145:  uint16(sym_float_literal),
	6146:  uint16(190),
	6147:  uint16(1),
	6148:  uint16(sym_scoped_identifier),
	6149:  uint16(335),
	6150:  uint16(1),
	6151:  uint16(sym_generic_type),
	6152:  uint16(45),
	6153:  uint16(2),
	6154:  uint16(anon_sym_true),
	6155:  uint16(anon_sym_false),
	6156:  uint16(398),
	6157:  uint16(2),
	6158:  uint16(anon_sym__),
	6159:  uint16(sym_integer_literal),
	6160:  uint16(246),
	6161:  uint16(2),
	6162:  uint16(sym_scoped_type_identifier),
	6163:  uint16(sym__type_identifier),
	6164:  uint16(318),
	6165:  uint16(2),
	6166:  uint16(sym_bracketed_type),
	6167:  uint16(sym__path),
	6168:  uint16(356),
	6169:  uint16(5),
	6170:  uint16(anon_sym_u8),
	6171:  uint16(anon_sym_u64),
	6172:  uint16(anon_sym_u128),
	6173:  uint16(anon_sym_bool),
	6174:  uint16(anon_sym_address),
	6175:  uint16(198),
	6176:  uint16(10),
	6177:  uint16(sym_const_block),
	6178:  uint16(sym__pattern),
	6179:  uint16(sym_tuple_pattern),
	6180:  uint16(sym_struct_pattern),
	6181:  uint16(sym_mut_pattern),
	6182:  uint16(sym_reference_pattern),
	6183:  uint16(sym_or_pattern),
	6184:  uint16(sym__literal_pattern),
	6185:  uint16(sym_negative_literal),
	6186:  uint16(sym_boolean_literal),
	6187:  uint16(18),
	6188:  uint16(3),
	6189:  uint16(1),
	6190:  uint16(sym_comment),
	6191:  uint16(120),
	6192:  uint16(1),
	6193:  uint16(anon_sym_LT),
	6194:  uint16(344),
	6195:  uint16(1),
	6196:  uint16(sym_identifier),
	6197:  uint16(346),
	6198:  uint16(1),
	6199:  uint16(anon_sym_COLON_COLON),
	6200:  uint16(350),
	6201:  uint16(1),
	6202:  uint16(anon_sym_LPAREN),
	6203:  uint16(358),
	6204:  uint16(1),
	6205:  uint16(anon_sym_AMP),
	6206:  uint16(360),
	6207:  uint16(1),
	6208:  uint16(sym_mutable_specifier),
	6209:  uint16(362),
	6210:  uint16(1),
	6211:  uint16(anon_sym_DASH),
	6212:  uint16(364),
	6213:  uint16(1),
	6214:  uint16(anon_sym_const),
	6215:  uint16(404),
	6216:  uint16(1),
	6217:  uint16(sym_float_literal),
	6218:  uint16(190),
	6219:  uint16(1),
	6220:  uint16(sym_scoped_identifier),
	6221:  uint16(335),
	6222:  uint16(1),
	6223:  uint16(sym_generic_type),
	6224:  uint16(45),
	6225:  uint16(2),
	6226:  uint16(anon_sym_true),
	6227:  uint16(anon_sym_false),
	6228:  uint16(402),
	6229:  uint16(2),
	6230:  uint16(anon_sym__),
	6231:  uint16(sym_integer_literal),
	6232:  uint16(246),
	6233:  uint16(2),
	6234:  uint16(sym_scoped_type_identifier),
	6235:  uint16(sym__type_identifier),
	6236:  uint16(318),
	6237:  uint16(2),
	6238:  uint16(sym_bracketed_type),
	6239:  uint16(sym__path),
	6240:  uint16(356),
	6241:  uint16(5),
	6242:  uint16(anon_sym_u8),
	6243:  uint16(anon_sym_u64),
	6244:  uint16(anon_sym_u128),
	6245:  uint16(anon_sym_bool),
	6246:  uint16(anon_sym_address),
	6247:  uint16(206),
	6248:  uint16(10),
	6249:  uint16(sym_const_block),
	6250:  uint16(sym__pattern),
	6251:  uint16(sym_tuple_pattern),
	6252:  uint16(sym_struct_pattern),
	6253:  uint16(sym_mut_pattern),
	6254:  uint16(sym_reference_pattern),
	6255:  uint16(sym_or_pattern),
	6256:  uint16(sym__literal_pattern),
	6257:  uint16(sym_negative_literal),
	6258:  uint16(sym_boolean_literal),
	6259:  uint16(3),
	6260:  uint16(3),
	6261:  uint16(1),
	6262:  uint16(sym_comment),
	6263:  uint16(238),
	6264:  uint16(6),
	6265:  uint16(anon_sym_EQ),
	6266:  uint16(anon_sym_LT),
	6267:  uint16(anon_sym_GT),
	6268:  uint16(anon_sym_AMP),
	6269:  uint16(anon_sym_PIPE),
	6270:  uint16(anon_sym_SLASH),
	6271:  uint16(240),
	6272:  uint16(25),
	6273:  uint16(anon_sym_RBRACE),
	6274:  uint16(anon_sym_SEMI),
	6275:  uint16(anon_sym_RBRACK),
	6276:  uint16(anon_sym_fun),
	6277:  uint16(anon_sym_public),
	6278:  uint16(anon_sym_COMMA),
	6279:  uint16(anon_sym_LPAREN),
	6280:  uint16(anon_sym_RPAREN),
	6281:  uint16(anon_sym_use),
	6282:  uint16(anon_sym_as),
	6283:  uint16(anon_sym_AMP_AMP),
	6284:  uint16(anon_sym_PIPE_PIPE),
	6285:  uint16(anon_sym_CARET),
	6286:  uint16(anon_sym_EQ_EQ),
	6287:  uint16(anon_sym_BANG_EQ),
	6288:  uint16(anon_sym_LT_EQ),
	6289:  uint16(anon_sym_GT_EQ),
	6290:  uint16(anon_sym_LT_LT),
	6291:  uint16(anon_sym_GT_GT),
	6292:  uint16(anon_sym_PLUS),
	6293:  uint16(anon_sym_DASH),
	6294:  uint16(anon_sym_STAR),
	6295:  uint16(anon_sym_PERCENT),
	6296:  uint16(anon_sym_else),
	6297:  uint16(anon_sym_DOT),
	6298:  uint16(3),
	6299:  uint16(3),
	6300:  uint16(1),
	6301:  uint16(sym_comment),
	6302:  uint16(234),
	6303:  uint16(6),
	6304:  uint16(anon_sym_EQ),
	6305:  uint16(anon_sym_LT),
	6306:  uint16(anon_sym_GT),
	6307:  uint16(anon_sym_AMP),
	6308:  uint16(anon_sym_PIPE),
	6309:  uint16(anon_sym_SLASH),
	6310:  uint16(236),
	6311:  uint16(25),
	6312:  uint16(anon_sym_RBRACE),
	6313:  uint16(anon_sym_SEMI),
	6314:  uint16(anon_sym_RBRACK),
	6315:  uint16(anon_sym_fun),
	6316:  uint16(anon_sym_public),
	6317:  uint16(anon_sym_COMMA),
	6318:  uint16(anon_sym_LPAREN),
	6319:  uint16(anon_sym_RPAREN),
	6320:  uint16(anon_sym_use),
	6321:  uint16(anon_sym_as),
	6322:  uint16(anon_sym_AMP_AMP),
	6323:  uint16(anon_sym_PIPE_PIPE),
	6324:  uint16(anon_sym_CARET),
	6325:  uint16(anon_sym_EQ_EQ),
	6326:  uint16(anon_sym_BANG_EQ),
	6327:  uint16(anon_sym_LT_EQ),
	6328:  uint16(anon_sym_GT_EQ),
	6329:  uint16(anon_sym_LT_LT),
	6330:  uint16(anon_sym_GT_GT),
	6331:  uint16(anon_sym_PLUS),
	6332:  uint16(anon_sym_DASH),
	6333:  uint16(anon_sym_STAR),
	6334:  uint16(anon_sym_PERCENT),
	6335:  uint16(anon_sym_else),
	6336:  uint16(anon_sym_DOT),
	6337:  uint16(3),
	6338:  uint16(3),
	6339:  uint16(1),
	6340:  uint16(sym_comment),
	6341:  uint16(240),
	6342:  uint16(10),
	6343:  uint16(anon_sym_COLON_COLON),
	6344:  uint16(anon_sym_LBRACE),
	6345:  uint16(anon_sym_RBRACE),
	6346:  uint16(anon_sym_LT),
	6347:  uint16(anon_sym_COMMA),
	6348:  uint16(anon_sym_GT),
	6349:  uint16(anon_sym_LPAREN),
	6350:  uint16(anon_sym_AMP),
	6351:  uint16(anon_sym_BANG),
	6352:  uint16(sym_float_literal),
	6353:  uint16(238),
	6354:  uint16(19),
	6355:  uint16(anon_sym_fun),
	6356:  uint16(anon_sym_public),
	6357:  uint16(anon_sym_use),
	6358:  uint16(anon_sym_u8),
	6359:  uint16(anon_sym_u64),
	6360:  uint16(anon_sym_u128),
	6361:  uint16(anon_sym_bool),
	6362:  uint16(anon_sym_address),
	6363:  uint16(anon_sym_return),
	6364:  uint16(anon_sym_if),
	6365:  uint16(anon_sym_while),
	6366:  uint16(anon_sym_loop),
	6367:  uint16(anon_sym_const),
	6368:  uint16(anon_sym_break),
	6369:  uint16(anon_sym_continue),
	6370:  uint16(sym_integer_literal),
	6371:  uint16(anon_sym_true),
	6372:  uint16(anon_sym_false),
	6373:  uint16(sym_identifier),
	6374:  uint16(3),
	6375:  uint16(3),
	6376:  uint16(1),
	6377:  uint16(sym_comment),
	6378:  uint16(236),
	6379:  uint16(10),
	6380:  uint16(anon_sym_COLON_COLON),
	6381:  uint16(anon_sym_LBRACE),
	6382:  uint16(anon_sym_RBRACE),
	6383:  uint16(anon_sym_LT),
	6384:  uint16(anon_sym_COMMA),
	6385:  uint16(anon_sym_GT),
	6386:  uint16(anon_sym_LPAREN),
	6387:  uint16(anon_sym_AMP),
	6388:  uint16(anon_sym_BANG),
	6389:  uint16(sym_float_literal),
	6390:  uint16(234),
	6391:  uint16(19),
	6392:  uint16(anon_sym_fun),
	6393:  uint16(anon_sym_public),
	6394:  uint16(anon_sym_use),
	6395:  uint16(anon_sym_u8),
	6396:  uint16(anon_sym_u64),
	6397:  uint16(anon_sym_u128),
	6398:  uint16(anon_sym_bool),
	6399:  uint16(anon_sym_address),
	6400:  uint16(anon_sym_return),
	6401:  uint16(anon_sym_if),
	6402:  uint16(anon_sym_while),
	6403:  uint16(anon_sym_loop),
	6404:  uint16(anon_sym_const),
	6405:  uint16(anon_sym_break),
	6406:  uint16(anon_sym_continue),
	6407:  uint16(sym_integer_literal),
	6408:  uint16(anon_sym_true),
	6409:  uint16(anon_sym_false),
	6410:  uint16(sym_identifier),
	6411:  uint16(5),
	6412:  uint16(3),
	6413:  uint16(1),
	6414:  uint16(sym_comment),
	6415:  uint16(406),
	6416:  uint16(1),
	6417:  uint16(anon_sym_else),
	6418:  uint16(125),
	6419:  uint16(1),
	6420:  uint16(sym_else_clause),
	6421:  uint16(224),
	6422:  uint16(6),
	6423:  uint16(anon_sym_EQ),
	6424:  uint16(anon_sym_LT),
	6425:  uint16(anon_sym_GT),
	6426:  uint16(anon_sym_AMP),
	6427:  uint16(anon_sym_PIPE),
	6428:  uint16(anon_sym_SLASH),
	6429:  uint16(226),
	6430:  uint16(20),
	6431:  uint16(anon_sym_SEMI),
	6432:  uint16(anon_sym_RBRACK),
	6433:  uint16(anon_sym_COMMA),
	6434:  uint16(anon_sym_LPAREN),
	6435:  uint16(anon_sym_RPAREN),
	6436:  uint16(anon_sym_as),
	6437:  uint16(anon_sym_AMP_AMP),
	6438:  uint16(anon_sym_PIPE_PIPE),
	6439:  uint16(anon_sym_CARET),
	6440:  uint16(anon_sym_EQ_EQ),
	6441:  uint16(anon_sym_BANG_EQ),
	6442:  uint16(anon_sym_LT_EQ),
	6443:  uint16(anon_sym_GT_EQ),
	6444:  uint16(anon_sym_LT_LT),
	6445:  uint16(anon_sym_GT_GT),
	6446:  uint16(anon_sym_PLUS),
	6447:  uint16(anon_sym_DASH),
	6448:  uint16(anon_sym_STAR),
	6449:  uint16(anon_sym_PERCENT),
	6450:  uint16(anon_sym_DOT),
	6451:  uint16(3),
	6452:  uint16(3),
	6453:  uint16(1),
	6454:  uint16(sym_comment),
	6455:  uint16(410),
	6456:  uint16(8),
	6457:  uint16(anon_sym_COLON_COLON),
	6458:  uint16(anon_sym_LBRACE),
	6459:  uint16(anon_sym_RBRACE),
	6460:  uint16(anon_sym_LT),
	6461:  uint16(anon_sym_LPAREN),
	6462:  uint16(anon_sym_AMP),
	6463:  uint16(anon_sym_BANG),
	6464:  uint16(sym_float_literal),
	6465:  uint16(408),
	6466:  uint16(19),
	6467:  uint16(anon_sym_fun),
	6468:  uint16(anon_sym_public),
	6469:  uint16(anon_sym_use),
	6470:  uint16(anon_sym_u8),
	6471:  uint16(anon_sym_u64),
	6472:  uint16(anon_sym_u128),
	6473:  uint16(anon_sym_bool),
	6474:  uint16(anon_sym_address),
	6475:  uint16(anon_sym_return),
	6476:  uint16(anon_sym_if),
	6477:  uint16(anon_sym_while),
	6478:  uint16(anon_sym_loop),
	6479:  uint16(anon_sym_const),
	6480:  uint16(anon_sym_break),
	6481:  uint16(anon_sym_continue),
	6482:  uint16(sym_integer_literal),
	6483:  uint16(anon_sym_true),
	6484:  uint16(anon_sym_false),
	6485:  uint16(sym_identifier),
	6486:  uint16(4),
	6487:  uint16(3),
	6488:  uint16(1),
	6489:  uint16(sym_comment),
	6490:  uint16(412),
	6491:  uint16(1),
	6492:  uint16(anon_sym_COLON_COLON),
	6493:  uint16(248),
	6494:  uint16(6),
	6495:  uint16(anon_sym_EQ),
	6496:  uint16(anon_sym_LT),
	6497:  uint16(anon_sym_GT),
	6498:  uint16(anon_sym_AMP),
	6499:  uint16(anon_sym_PIPE),
	6500:  uint16(anon_sym_SLASH),
	6501:  uint16(246),
	6502:  uint16(20),
	6503:  uint16(anon_sym_SEMI),
	6504:  uint16(anon_sym_RBRACK),
	6505:  uint16(anon_sym_COMMA),
	6506:  uint16(anon_sym_LPAREN),
	6507:  uint16(anon_sym_RPAREN),
	6508:  uint16(anon_sym_as),
	6509:  uint16(anon_sym_AMP_AMP),
	6510:  uint16(anon_sym_PIPE_PIPE),
	6511:  uint16(anon_sym_CARET),
	6512:  uint16(anon_sym_EQ_EQ),
	6513:  uint16(anon_sym_BANG_EQ),
	6514:  uint16(anon_sym_LT_EQ),
	6515:  uint16(anon_sym_GT_EQ),
	6516:  uint16(anon_sym_LT_LT),
	6517:  uint16(anon_sym_GT_GT),
	6518:  uint16(anon_sym_PLUS),
	6519:  uint16(anon_sym_DASH),
	6520:  uint16(anon_sym_STAR),
	6521:  uint16(anon_sym_PERCENT),
	6522:  uint16(anon_sym_DOT),
	6523:  uint16(5),
	6524:  uint16(3),
	6525:  uint16(1),
	6526:  uint16(sym_comment),
	6527:  uint16(418),
	6528:  uint16(1),
	6529:  uint16(anon_sym_LPAREN),
	6530:  uint16(114),
	6531:  uint16(1),
	6532:  uint16(sym_arguments),
	6533:  uint16(416),
	6534:  uint16(6),
	6535:  uint16(anon_sym_EQ),
	6536:  uint16(anon_sym_LT),
	6537:  uint16(anon_sym_GT),
	6538:  uint16(anon_sym_AMP),
	6539:  uint16(anon_sym_PIPE),
	6540:  uint16(anon_sym_SLASH),
	6541:  uint16(414),
	6542:  uint16(19),
	6543:  uint16(anon_sym_SEMI),
	6544:  uint16(anon_sym_RBRACK),
	6545:  uint16(anon_sym_COMMA),
	6546:  uint16(anon_sym_RPAREN),
	6547:  uint16(anon_sym_as),
	6548:  uint16(anon_sym_AMP_AMP),
	6549:  uint16(anon_sym_PIPE_PIPE),
	6550:  uint16(anon_sym_CARET),
	6551:  uint16(anon_sym_EQ_EQ),
	6552:  uint16(anon_sym_BANG_EQ),
	6553:  uint16(anon_sym_LT_EQ),
	6554:  uint16(anon_sym_GT_EQ),
	6555:  uint16(anon_sym_LT_LT),
	6556:  uint16(anon_sym_GT_GT),
	6557:  uint16(anon_sym_PLUS),
	6558:  uint16(anon_sym_DASH),
	6559:  uint16(anon_sym_STAR),
	6560:  uint16(anon_sym_PERCENT),
	6561:  uint16(anon_sym_DOT),
	6562:  uint16(3),
	6563:  uint16(3),
	6564:  uint16(1),
	6565:  uint16(sym_comment),
	6566:  uint16(422),
	6567:  uint16(6),
	6568:  uint16(anon_sym_EQ),
	6569:  uint16(anon_sym_LT),
	6570:  uint16(anon_sym_GT),
	6571:  uint16(anon_sym_AMP),
	6572:  uint16(anon_sym_PIPE),
	6573:  uint16(anon_sym_SLASH),
	6574:  uint16(420),
	6575:  uint16(21),
	6576:  uint16(anon_sym_RBRACE),
	6577:  uint16(anon_sym_SEMI),
	6578:  uint16(anon_sym_RBRACK),
	6579:  uint16(anon_sym_COMMA),
	6580:  uint16(anon_sym_LPAREN),
	6581:  uint16(anon_sym_RPAREN),
	6582:  uint16(anon_sym_as),
	6583:  uint16(anon_sym_AMP_AMP),
	6584:  uint16(anon_sym_PIPE_PIPE),
	6585:  uint16(anon_sym_CARET),
	6586:  uint16(anon_sym_EQ_EQ),
	6587:  uint16(anon_sym_BANG_EQ),
	6588:  uint16(anon_sym_LT_EQ),
	6589:  uint16(anon_sym_GT_EQ),
	6590:  uint16(anon_sym_LT_LT),
	6591:  uint16(anon_sym_GT_GT),
	6592:  uint16(anon_sym_PLUS),
	6593:  uint16(anon_sym_DASH),
	6594:  uint16(anon_sym_STAR),
	6595:  uint16(anon_sym_PERCENT),
	6596:  uint16(anon_sym_DOT),
	6597:  uint16(4),
	6598:  uint16(3),
	6599:  uint16(1),
	6600:  uint16(sym_comment),
	6601:  uint16(412),
	6602:  uint16(1),
	6603:  uint16(anon_sym_COLON_COLON),
	6604:  uint16(248),
	6605:  uint16(6),
	6606:  uint16(anon_sym_EQ),
	6607:  uint16(anon_sym_LT),
	6608:  uint16(anon_sym_GT),
	6609:  uint16(anon_sym_AMP),
	6610:  uint16(anon_sym_PIPE),
	6611:  uint16(anon_sym_SLASH),
	6612:  uint16(246),
	6613:  uint16(20),
	6614:  uint16(anon_sym_SEMI),
	6615:  uint16(anon_sym_RBRACK),
	6616:  uint16(anon_sym_COMMA),
	6617:  uint16(anon_sym_LPAREN),
	6618:  uint16(anon_sym_RPAREN),
	6619:  uint16(anon_sym_as),
	6620:  uint16(anon_sym_AMP_AMP),
	6621:  uint16(anon_sym_PIPE_PIPE),
	6622:  uint16(anon_sym_CARET),
	6623:  uint16(anon_sym_EQ_EQ),
	6624:  uint16(anon_sym_BANG_EQ),
	6625:  uint16(anon_sym_LT_EQ),
	6626:  uint16(anon_sym_GT_EQ),
	6627:  uint16(anon_sym_LT_LT),
	6628:  uint16(anon_sym_GT_GT),
	6629:  uint16(anon_sym_PLUS),
	6630:  uint16(anon_sym_DASH),
	6631:  uint16(anon_sym_STAR),
	6632:  uint16(anon_sym_PERCENT),
	6633:  uint16(anon_sym_DOT),
	6634:  uint16(3),
	6635:  uint16(3),
	6636:  uint16(1),
	6637:  uint16(sym_comment),
	6638:  uint16(426),
	6639:  uint16(6),
	6640:  uint16(anon_sym_EQ),
	6641:  uint16(anon_sym_LT),
	6642:  uint16(anon_sym_GT),
	6643:  uint16(anon_sym_AMP),
	6644:  uint16(anon_sym_PIPE),
	6645:  uint16(anon_sym_SLASH),
	6646:  uint16(424),
	6647:  uint16(21),
	6648:  uint16(anon_sym_COLON_COLON),
	6649:  uint16(anon_sym_SEMI),
	6650:  uint16(anon_sym_RBRACK),
	6651:  uint16(anon_sym_COMMA),
	6652:  uint16(anon_sym_LPAREN),
	6653:  uint16(anon_sym_RPAREN),
	6654:  uint16(anon_sym_as),
	6655:  uint16(anon_sym_AMP_AMP),
	6656:  uint16(anon_sym_PIPE_PIPE),
	6657:  uint16(anon_sym_CARET),
	6658:  uint16(anon_sym_EQ_EQ),
	6659:  uint16(anon_sym_BANG_EQ),
	6660:  uint16(anon_sym_LT_EQ),
	6661:  uint16(anon_sym_GT_EQ),
	6662:  uint16(anon_sym_LT_LT),
	6663:  uint16(anon_sym_GT_GT),
	6664:  uint16(anon_sym_PLUS),
	6665:  uint16(anon_sym_DASH),
	6666:  uint16(anon_sym_STAR),
	6667:  uint16(anon_sym_PERCENT),
	6668:  uint16(anon_sym_DOT),
	6669:  uint16(3),
	6670:  uint16(3),
	6671:  uint16(1),
	6672:  uint16(sym_comment),
	6673:  uint16(430),
	6674:  uint16(8),
	6675:  uint16(anon_sym_COLON_COLON),
	6676:  uint16(anon_sym_LBRACE),
	6677:  uint16(anon_sym_RBRACE),
	6678:  uint16(anon_sym_LT),
	6679:  uint16(anon_sym_LPAREN),
	6680:  uint16(anon_sym_AMP),
	6681:  uint16(anon_sym_BANG),
	6682:  uint16(sym_float_literal),
	6683:  uint16(428),
	6684:  uint16(19),
	6685:  uint16(anon_sym_fun),
	6686:  uint16(anon_sym_public),
	6687:  uint16(anon_sym_use),
	6688:  uint16(anon_sym_u8),
	6689:  uint16(anon_sym_u64),
	6690:  uint16(anon_sym_u128),
	6691:  uint16(anon_sym_bool),
	6692:  uint16(anon_sym_address),
	6693:  uint16(anon_sym_return),
	6694:  uint16(anon_sym_if),
	6695:  uint16(anon_sym_while),
	6696:  uint16(anon_sym_loop),
	6697:  uint16(anon_sym_const),
	6698:  uint16(anon_sym_break),
	6699:  uint16(anon_sym_continue),
	6700:  uint16(sym_integer_literal),
	6701:  uint16(anon_sym_true),
	6702:  uint16(anon_sym_false),
	6703:  uint16(sym_identifier),
	6704:  uint16(3),
	6705:  uint16(3),
	6706:  uint16(1),
	6707:  uint16(sym_comment),
	6708:  uint16(434),
	6709:  uint16(8),
	6710:  uint16(anon_sym_COLON_COLON),
	6711:  uint16(anon_sym_LBRACE),
	6712:  uint16(anon_sym_RBRACE),
	6713:  uint16(anon_sym_LT),
	6714:  uint16(anon_sym_LPAREN),
	6715:  uint16(anon_sym_AMP),
	6716:  uint16(anon_sym_BANG),
	6717:  uint16(sym_float_literal),
	6718:  uint16(432),
	6719:  uint16(19),
	6720:  uint16(anon_sym_fun),
	6721:  uint16(anon_sym_public),
	6722:  uint16(anon_sym_use),
	6723:  uint16(anon_sym_u8),
	6724:  uint16(anon_sym_u64),
	6725:  uint16(anon_sym_u128),
	6726:  uint16(anon_sym_bool),
	6727:  uint16(anon_sym_address),
	6728:  uint16(anon_sym_return),
	6729:  uint16(anon_sym_if),
	6730:  uint16(anon_sym_while),
	6731:  uint16(anon_sym_loop),
	6732:  uint16(anon_sym_const),
	6733:  uint16(anon_sym_break),
	6734:  uint16(anon_sym_continue),
	6735:  uint16(sym_integer_literal),
	6736:  uint16(anon_sym_true),
	6737:  uint16(anon_sym_false),
	6738:  uint16(sym_identifier),
	6739:  uint16(3),
	6740:  uint16(3),
	6741:  uint16(1),
	6742:  uint16(sym_comment),
	6743:  uint16(438),
	6744:  uint16(8),
	6745:  uint16(anon_sym_COLON_COLON),
	6746:  uint16(anon_sym_LBRACE),
	6747:  uint16(anon_sym_RBRACE),
	6748:  uint16(anon_sym_LT),
	6749:  uint16(anon_sym_LPAREN),
	6750:  uint16(anon_sym_AMP),
	6751:  uint16(anon_sym_BANG),
	6752:  uint16(sym_float_literal),
	6753:  uint16(436),
	6754:  uint16(19),
	6755:  uint16(anon_sym_fun),
	6756:  uint16(anon_sym_public),
	6757:  uint16(anon_sym_use),
	6758:  uint16(anon_sym_u8),
	6759:  uint16(anon_sym_u64),
	6760:  uint16(anon_sym_u128),
	6761:  uint16(anon_sym_bool),
	6762:  uint16(anon_sym_address),
	6763:  uint16(anon_sym_return),
	6764:  uint16(anon_sym_if),
	6765:  uint16(anon_sym_while),
	6766:  uint16(anon_sym_loop),
	6767:  uint16(anon_sym_const),
	6768:  uint16(anon_sym_break),
	6769:  uint16(anon_sym_continue),
	6770:  uint16(sym_integer_literal),
	6771:  uint16(anon_sym_true),
	6772:  uint16(anon_sym_false),
	6773:  uint16(sym_identifier),
	6774:  uint16(4),
	6775:  uint16(3),
	6776:  uint16(1),
	6777:  uint16(sym_comment),
	6778:  uint16(440),
	6779:  uint16(1),
	6780:  uint16(anon_sym_COLON_COLON),
	6781:  uint16(248),
	6782:  uint16(6),
	6783:  uint16(anon_sym_EQ),
	6784:  uint16(anon_sym_LT),
	6785:  uint16(anon_sym_GT),
	6786:  uint16(anon_sym_AMP),
	6787:  uint16(anon_sym_PIPE),
	6788:  uint16(anon_sym_SLASH),
	6789:  uint16(246),
	6790:  uint16(20),
	6791:  uint16(anon_sym_SEMI),
	6792:  uint16(anon_sym_RBRACK),
	6793:  uint16(anon_sym_COMMA),
	6794:  uint16(anon_sym_LPAREN),
	6795:  uint16(anon_sym_RPAREN),
	6796:  uint16(anon_sym_as),
	6797:  uint16(anon_sym_AMP_AMP),
	6798:  uint16(anon_sym_PIPE_PIPE),
	6799:  uint16(anon_sym_CARET),
	6800:  uint16(anon_sym_EQ_EQ),
	6801:  uint16(anon_sym_BANG_EQ),
	6802:  uint16(anon_sym_LT_EQ),
	6803:  uint16(anon_sym_GT_EQ),
	6804:  uint16(anon_sym_LT_LT),
	6805:  uint16(anon_sym_GT_GT),
	6806:  uint16(anon_sym_PLUS),
	6807:  uint16(anon_sym_DASH),
	6808:  uint16(anon_sym_STAR),
	6809:  uint16(anon_sym_PERCENT),
	6810:  uint16(anon_sym_DOT),
	6811:  uint16(3),
	6812:  uint16(3),
	6813:  uint16(1),
	6814:  uint16(sym_comment),
	6815:  uint16(444),
	6816:  uint16(8),
	6817:  uint16(anon_sym_COLON_COLON),
	6818:  uint16(anon_sym_LBRACE),
	6819:  uint16(anon_sym_RBRACE),
	6820:  uint16(anon_sym_LT),
	6821:  uint16(anon_sym_LPAREN),
	6822:  uint16(anon_sym_AMP),
	6823:  uint16(anon_sym_BANG),
	6824:  uint16(sym_float_literal),
	6825:  uint16(442),
	6826:  uint16(19),
	6827:  uint16(anon_sym_fun),
	6828:  uint16(anon_sym_public),
	6829:  uint16(anon_sym_use),
	6830:  uint16(anon_sym_u8),
	6831:  uint16(anon_sym_u64),
	6832:  uint16(anon_sym_u128),
	6833:  uint16(anon_sym_bool),
	6834:  uint16(anon_sym_address),
	6835:  uint16(anon_sym_return),
	6836:  uint16(anon_sym_if),
	6837:  uint16(anon_sym_while),
	6838:  uint16(anon_sym_loop),
	6839:  uint16(anon_sym_const),
	6840:  uint16(anon_sym_break),
	6841:  uint16(anon_sym_continue),
	6842:  uint16(sym_integer_literal),
	6843:  uint16(anon_sym_true),
	6844:  uint16(anon_sym_false),
	6845:  uint16(sym_identifier),
	6846:  uint16(3),
	6847:  uint16(3),
	6848:  uint16(1),
	6849:  uint16(sym_comment),
	6850:  uint16(448),
	6851:  uint16(6),
	6852:  uint16(anon_sym_EQ),
	6853:  uint16(anon_sym_LT),
	6854:  uint16(anon_sym_GT),
	6855:  uint16(anon_sym_AMP),
	6856:  uint16(anon_sym_PIPE),
	6857:  uint16(anon_sym_SLASH),
	6858:  uint16(446),
	6859:  uint16(21),
	6860:  uint16(anon_sym_COLON_COLON),
	6861:  uint16(anon_sym_SEMI),
	6862:  uint16(anon_sym_RBRACK),
	6863:  uint16(anon_sym_COMMA),
	6864:  uint16(anon_sym_LPAREN),
	6865:  uint16(anon_sym_RPAREN),
	6866:  uint16(anon_sym_as),
	6867:  uint16(anon_sym_AMP_AMP),
	6868:  uint16(anon_sym_PIPE_PIPE),
	6869:  uint16(anon_sym_CARET),
	6870:  uint16(anon_sym_EQ_EQ),
	6871:  uint16(anon_sym_BANG_EQ),
	6872:  uint16(anon_sym_LT_EQ),
	6873:  uint16(anon_sym_GT_EQ),
	6874:  uint16(anon_sym_LT_LT),
	6875:  uint16(anon_sym_GT_GT),
	6876:  uint16(anon_sym_PLUS),
	6877:  uint16(anon_sym_DASH),
	6878:  uint16(anon_sym_STAR),
	6879:  uint16(anon_sym_PERCENT),
	6880:  uint16(anon_sym_DOT),
	6881:  uint16(3),
	6882:  uint16(3),
	6883:  uint16(1),
	6884:  uint16(sym_comment),
	6885:  uint16(452),
	6886:  uint16(8),
	6887:  uint16(anon_sym_COLON_COLON),
	6888:  uint16(anon_sym_LBRACE),
	6889:  uint16(anon_sym_RBRACE),
	6890:  uint16(anon_sym_LT),
	6891:  uint16(anon_sym_LPAREN),
	6892:  uint16(anon_sym_AMP),
	6893:  uint16(anon_sym_BANG),
	6894:  uint16(sym_float_literal),
	6895:  uint16(450),
	6896:  uint16(19),
	6897:  uint16(anon_sym_fun),
	6898:  uint16(anon_sym_public),
	6899:  uint16(anon_sym_use),
	6900:  uint16(anon_sym_u8),
	6901:  uint16(anon_sym_u64),
	6902:  uint16(anon_sym_u128),
	6903:  uint16(anon_sym_bool),
	6904:  uint16(anon_sym_address),
	6905:  uint16(anon_sym_return),
	6906:  uint16(anon_sym_if),
	6907:  uint16(anon_sym_while),
	6908:  uint16(anon_sym_loop),
	6909:  uint16(anon_sym_const),
	6910:  uint16(anon_sym_break),
	6911:  uint16(anon_sym_continue),
	6912:  uint16(sym_integer_literal),
	6913:  uint16(anon_sym_true),
	6914:  uint16(anon_sym_false),
	6915:  uint16(sym_identifier),
	6916:  uint16(3),
	6917:  uint16(3),
	6918:  uint16(1),
	6919:  uint16(sym_comment),
	6920:  uint16(456),
	6921:  uint16(6),
	6922:  uint16(anon_sym_EQ),
	6923:  uint16(anon_sym_LT),
	6924:  uint16(anon_sym_GT),
	6925:  uint16(anon_sym_AMP),
	6926:  uint16(anon_sym_PIPE),
	6927:  uint16(anon_sym_SLASH),
	6928:  uint16(454),
	6929:  uint16(21),
	6930:  uint16(anon_sym_COLON_COLON),
	6931:  uint16(anon_sym_SEMI),
	6932:  uint16(anon_sym_RBRACK),
	6933:  uint16(anon_sym_COMMA),
	6934:  uint16(anon_sym_LPAREN),
	6935:  uint16(anon_sym_RPAREN),
	6936:  uint16(anon_sym_as),
	6937:  uint16(anon_sym_AMP_AMP),
	6938:  uint16(anon_sym_PIPE_PIPE),
	6939:  uint16(anon_sym_CARET),
	6940:  uint16(anon_sym_EQ_EQ),
	6941:  uint16(anon_sym_BANG_EQ),
	6942:  uint16(anon_sym_LT_EQ),
	6943:  uint16(anon_sym_GT_EQ),
	6944:  uint16(anon_sym_LT_LT),
	6945:  uint16(anon_sym_GT_GT),
	6946:  uint16(anon_sym_PLUS),
	6947:  uint16(anon_sym_DASH),
	6948:  uint16(anon_sym_STAR),
	6949:  uint16(anon_sym_PERCENT),
	6950:  uint16(anon_sym_DOT),
	6951:  uint16(3),
	6952:  uint16(3),
	6953:  uint16(1),
	6954:  uint16(sym_comment),
	6955:  uint16(460),
	6956:  uint16(8),
	6957:  uint16(anon_sym_COLON_COLON),
	6958:  uint16(anon_sym_LBRACE),
	6959:  uint16(anon_sym_RBRACE),
	6960:  uint16(anon_sym_LT),
	6961:  uint16(anon_sym_LPAREN),
	6962:  uint16(anon_sym_AMP),
	6963:  uint16(anon_sym_BANG),
	6964:  uint16(sym_float_literal),
	6965:  uint16(458),
	6966:  uint16(19),
	6967:  uint16(anon_sym_fun),
	6968:  uint16(anon_sym_public),
	6969:  uint16(anon_sym_use),
	6970:  uint16(anon_sym_u8),
	6971:  uint16(anon_sym_u64),
	6972:  uint16(anon_sym_u128),
	6973:  uint16(anon_sym_bool),
	6974:  uint16(anon_sym_address),
	6975:  uint16(anon_sym_return),
	6976:  uint16(anon_sym_if),
	6977:  uint16(anon_sym_while),
	6978:  uint16(anon_sym_loop),
	6979:  uint16(anon_sym_const),
	6980:  uint16(anon_sym_break),
	6981:  uint16(anon_sym_continue),
	6982:  uint16(sym_integer_literal),
	6983:  uint16(anon_sym_true),
	6984:  uint16(anon_sym_false),
	6985:  uint16(sym_identifier),
	6986:  uint16(3),
	6987:  uint16(3),
	6988:  uint16(1),
	6989:  uint16(sym_comment),
	6990:  uint16(464),
	6991:  uint16(6),
	6992:  uint16(anon_sym_EQ),
	6993:  uint16(anon_sym_LT),
	6994:  uint16(anon_sym_GT),
	6995:  uint16(anon_sym_AMP),
	6996:  uint16(anon_sym_PIPE),
	6997:  uint16(anon_sym_SLASH),
	6998:  uint16(462),
	6999:  uint16(21),
	7000:  uint16(anon_sym_COLON_COLON),
	7001:  uint16(anon_sym_SEMI),
	7002:  uint16(anon_sym_RBRACK),
	7003:  uint16(anon_sym_COMMA),
	7004:  uint16(anon_sym_LPAREN),
	7005:  uint16(anon_sym_RPAREN),
	7006:  uint16(anon_sym_as),
	7007:  uint16(anon_sym_AMP_AMP),
	7008:  uint16(anon_sym_PIPE_PIPE),
	7009:  uint16(anon_sym_CARET),
	7010:  uint16(anon_sym_EQ_EQ),
	7011:  uint16(anon_sym_BANG_EQ),
	7012:  uint16(anon_sym_LT_EQ),
	7013:  uint16(anon_sym_GT_EQ),
	7014:  uint16(anon_sym_LT_LT),
	7015:  uint16(anon_sym_GT_GT),
	7016:  uint16(anon_sym_PLUS),
	7017:  uint16(anon_sym_DASH),
	7018:  uint16(anon_sym_STAR),
	7019:  uint16(anon_sym_PERCENT),
	7020:  uint16(anon_sym_DOT),
	7021:  uint16(3),
	7022:  uint16(3),
	7023:  uint16(1),
	7024:  uint16(sym_comment),
	7025:  uint16(468),
	7026:  uint16(8),
	7027:  uint16(anon_sym_COLON_COLON),
	7028:  uint16(anon_sym_LBRACE),
	7029:  uint16(anon_sym_RBRACE),
	7030:  uint16(anon_sym_LT),
	7031:  uint16(anon_sym_LPAREN),
	7032:  uint16(anon_sym_AMP),
	7033:  uint16(anon_sym_BANG),
	7034:  uint16(sym_float_literal),
	7035:  uint16(466),
	7036:  uint16(19),
	7037:  uint16(anon_sym_fun),
	7038:  uint16(anon_sym_public),
	7039:  uint16(anon_sym_use),
	7040:  uint16(anon_sym_u8),
	7041:  uint16(anon_sym_u64),
	7042:  uint16(anon_sym_u128),
	7043:  uint16(anon_sym_bool),
	7044:  uint16(anon_sym_address),
	7045:  uint16(anon_sym_return),
	7046:  uint16(anon_sym_if),
	7047:  uint16(anon_sym_while),
	7048:  uint16(anon_sym_loop),
	7049:  uint16(anon_sym_const),
	7050:  uint16(anon_sym_break),
	7051:  uint16(anon_sym_continue),
	7052:  uint16(sym_integer_literal),
	7053:  uint16(anon_sym_true),
	7054:  uint16(anon_sym_false),
	7055:  uint16(sym_identifier),
	7056:  uint16(3),
	7057:  uint16(3),
	7058:  uint16(1),
	7059:  uint16(sym_comment),
	7060:  uint16(472),
	7061:  uint16(6),
	7062:  uint16(anon_sym_EQ),
	7063:  uint16(anon_sym_LT),
	7064:  uint16(anon_sym_GT),
	7065:  uint16(anon_sym_AMP),
	7066:  uint16(anon_sym_PIPE),
	7067:  uint16(anon_sym_SLASH),
	7068:  uint16(470),
	7069:  uint16(21),
	7070:  uint16(anon_sym_COLON_COLON),
	7071:  uint16(anon_sym_SEMI),
	7072:  uint16(anon_sym_RBRACK),
	7073:  uint16(anon_sym_COMMA),
	7074:  uint16(anon_sym_LPAREN),
	7075:  uint16(anon_sym_RPAREN),
	7076:  uint16(anon_sym_as),
	7077:  uint16(anon_sym_AMP_AMP),
	7078:  uint16(anon_sym_PIPE_PIPE),
	7079:  uint16(anon_sym_CARET),
	7080:  uint16(anon_sym_EQ_EQ),
	7081:  uint16(anon_sym_BANG_EQ),
	7082:  uint16(anon_sym_LT_EQ),
	7083:  uint16(anon_sym_GT_EQ),
	7084:  uint16(anon_sym_LT_LT),
	7085:  uint16(anon_sym_GT_GT),
	7086:  uint16(anon_sym_PLUS),
	7087:  uint16(anon_sym_DASH),
	7088:  uint16(anon_sym_STAR),
	7089:  uint16(anon_sym_PERCENT),
	7090:  uint16(anon_sym_DOT),
	7091:  uint16(3),
	7092:  uint16(3),
	7093:  uint16(1),
	7094:  uint16(sym_comment),
	7095:  uint16(476),
	7096:  uint16(8),
	7097:  uint16(anon_sym_COLON_COLON),
	7098:  uint16(anon_sym_LBRACE),
	7099:  uint16(anon_sym_RBRACE),
	7100:  uint16(anon_sym_LT),
	7101:  uint16(anon_sym_LPAREN),
	7102:  uint16(anon_sym_AMP),
	7103:  uint16(anon_sym_BANG),
	7104:  uint16(sym_float_literal),
	7105:  uint16(474),
	7106:  uint16(19),
	7107:  uint16(anon_sym_fun),
	7108:  uint16(anon_sym_public),
	7109:  uint16(anon_sym_use),
	7110:  uint16(anon_sym_u8),
	7111:  uint16(anon_sym_u64),
	7112:  uint16(anon_sym_u128),
	7113:  uint16(anon_sym_bool),
	7114:  uint16(anon_sym_address),
	7115:  uint16(anon_sym_return),
	7116:  uint16(anon_sym_if),
	7117:  uint16(anon_sym_while),
	7118:  uint16(anon_sym_loop),
	7119:  uint16(anon_sym_const),
	7120:  uint16(anon_sym_break),
	7121:  uint16(anon_sym_continue),
	7122:  uint16(sym_integer_literal),
	7123:  uint16(anon_sym_true),
	7124:  uint16(anon_sym_false),
	7125:  uint16(sym_identifier),
	7126:  uint16(3),
	7127:  uint16(3),
	7128:  uint16(1),
	7129:  uint16(sym_comment),
	7130:  uint16(480),
	7131:  uint16(8),
	7132:  uint16(anon_sym_COLON_COLON),
	7133:  uint16(anon_sym_LBRACE),
	7134:  uint16(anon_sym_RBRACE),
	7135:  uint16(anon_sym_LT),
	7136:  uint16(anon_sym_LPAREN),
	7137:  uint16(anon_sym_AMP),
	7138:  uint16(anon_sym_BANG),
	7139:  uint16(sym_float_literal),
	7140:  uint16(478),
	7141:  uint16(19),
	7142:  uint16(anon_sym_fun),
	7143:  uint16(anon_sym_public),
	7144:  uint16(anon_sym_use),
	7145:  uint16(anon_sym_u8),
	7146:  uint16(anon_sym_u64),
	7147:  uint16(anon_sym_u128),
	7148:  uint16(anon_sym_bool),
	7149:  uint16(anon_sym_address),
	7150:  uint16(anon_sym_return),
	7151:  uint16(anon_sym_if),
	7152:  uint16(anon_sym_while),
	7153:  uint16(anon_sym_loop),
	7154:  uint16(anon_sym_const),
	7155:  uint16(anon_sym_break),
	7156:  uint16(anon_sym_continue),
	7157:  uint16(sym_integer_literal),
	7158:  uint16(anon_sym_true),
	7159:  uint16(anon_sym_false),
	7160:  uint16(sym_identifier),
	7161:  uint16(3),
	7162:  uint16(3),
	7163:  uint16(1),
	7164:  uint16(sym_comment),
	7165:  uint16(484),
	7166:  uint16(6),
	7167:  uint16(anon_sym_EQ),
	7168:  uint16(anon_sym_LT),
	7169:  uint16(anon_sym_GT),
	7170:  uint16(anon_sym_AMP),
	7171:  uint16(anon_sym_PIPE),
	7172:  uint16(anon_sym_SLASH),
	7173:  uint16(482),
	7174:  uint16(20),
	7175:  uint16(anon_sym_SEMI),
	7176:  uint16(anon_sym_RBRACK),
	7177:  uint16(anon_sym_COMMA),
	7178:  uint16(anon_sym_LPAREN),
	7179:  uint16(anon_sym_RPAREN),
	7180:  uint16(anon_sym_as),
	7181:  uint16(anon_sym_AMP_AMP),
	7182:  uint16(anon_sym_PIPE_PIPE),
	7183:  uint16(anon_sym_CARET),
	7184:  uint16(anon_sym_EQ_EQ),
	7185:  uint16(anon_sym_BANG_EQ),
	7186:  uint16(anon_sym_LT_EQ),
	7187:  uint16(anon_sym_GT_EQ),
	7188:  uint16(anon_sym_LT_LT),
	7189:  uint16(anon_sym_GT_GT),
	7190:  uint16(anon_sym_PLUS),
	7191:  uint16(anon_sym_DASH),
	7192:  uint16(anon_sym_STAR),
	7193:  uint16(anon_sym_PERCENT),
	7194:  uint16(anon_sym_DOT),
	7195:  uint16(3),
	7196:  uint16(3),
	7197:  uint16(1),
	7198:  uint16(sym_comment),
	7199:  uint16(488),
	7200:  uint16(6),
	7201:  uint16(anon_sym_EQ),
	7202:  uint16(anon_sym_LT),
	7203:  uint16(anon_sym_GT),
	7204:  uint16(anon_sym_AMP),
	7205:  uint16(anon_sym_PIPE),
	7206:  uint16(anon_sym_SLASH),
	7207:  uint16(486),
	7208:  uint16(20),
	7209:  uint16(anon_sym_SEMI),
	7210:  uint16(anon_sym_RBRACK),
	7211:  uint16(anon_sym_COMMA),
	7212:  uint16(anon_sym_LPAREN),
	7213:  uint16(anon_sym_RPAREN),
	7214:  uint16(anon_sym_as),
	7215:  uint16(anon_sym_AMP_AMP),
	7216:  uint16(anon_sym_PIPE_PIPE),
	7217:  uint16(anon_sym_CARET),
	7218:  uint16(anon_sym_EQ_EQ),
	7219:  uint16(anon_sym_BANG_EQ),
	7220:  uint16(anon_sym_LT_EQ),
	7221:  uint16(anon_sym_GT_EQ),
	7222:  uint16(anon_sym_LT_LT),
	7223:  uint16(anon_sym_GT_GT),
	7224:  uint16(anon_sym_PLUS),
	7225:  uint16(anon_sym_DASH),
	7226:  uint16(anon_sym_STAR),
	7227:  uint16(anon_sym_PERCENT),
	7228:  uint16(anon_sym_DOT),
	7229:  uint16(3),
	7230:  uint16(3),
	7231:  uint16(1),
	7232:  uint16(sym_comment),
	7233:  uint16(492),
	7234:  uint16(6),
	7235:  uint16(anon_sym_EQ),
	7236:  uint16(anon_sym_LT),
	7237:  uint16(anon_sym_GT),
	7238:  uint16(anon_sym_AMP),
	7239:  uint16(anon_sym_PIPE),
	7240:  uint16(anon_sym_SLASH),
	7241:  uint16(490),
	7242:  uint16(20),
	7243:  uint16(anon_sym_SEMI),
	7244:  uint16(anon_sym_RBRACK),
	7245:  uint16(anon_sym_COMMA),
	7246:  uint16(anon_sym_LPAREN),
	7247:  uint16(anon_sym_RPAREN),
	7248:  uint16(anon_sym_as),
	7249:  uint16(anon_sym_AMP_AMP),
	7250:  uint16(anon_sym_PIPE_PIPE),
	7251:  uint16(anon_sym_CARET),
	7252:  uint16(anon_sym_EQ_EQ),
	7253:  uint16(anon_sym_BANG_EQ),
	7254:  uint16(anon_sym_LT_EQ),
	7255:  uint16(anon_sym_GT_EQ),
	7256:  uint16(anon_sym_LT_LT),
	7257:  uint16(anon_sym_GT_GT),
	7258:  uint16(anon_sym_PLUS),
	7259:  uint16(anon_sym_DASH),
	7260:  uint16(anon_sym_STAR),
	7261:  uint16(anon_sym_PERCENT),
	7262:  uint16(anon_sym_DOT),
	7263:  uint16(14),
	7264:  uint16(3),
	7265:  uint16(1),
	7266:  uint16(sym_comment),
	7267:  uint16(496),
	7268:  uint16(1),
	7269:  uint16(anon_sym_EQ),
	7270:  uint16(500),
	7271:  uint16(1),
	7272:  uint16(anon_sym_as),
	7273:  uint16(502),
	7274:  uint16(1),
	7275:  uint16(anon_sym_AMP),
	7276:  uint16(504),
	7277:  uint16(1),
	7278:  uint16(anon_sym_PIPE),
	7279:  uint16(506),
	7280:  uint16(1),
	7281:  uint16(anon_sym_CARET),
	7282:  uint16(516),
	7283:  uint16(1),
	7284:  uint16(anon_sym_SLASH),
	7285:  uint16(518),
	7286:  uint16(1),
	7287:  uint16(anon_sym_DOT),
	7288:  uint16(498),
	7289:  uint16(2),
	7290:  uint16(anon_sym_LT),
	7291:  uint16(anon_sym_GT),
	7292:  uint16(510),
	7293:  uint16(2),
	7294:  uint16(anon_sym_LT_LT),
	7295:  uint16(anon_sym_GT_GT),
	7296:  uint16(512),
	7297:  uint16(2),
	7298:  uint16(anon_sym_PLUS),
	7299:  uint16(anon_sym_DASH),
	7300:  uint16(514),
	7301:  uint16(2),
	7302:  uint16(anon_sym_STAR),
	7303:  uint16(anon_sym_PERCENT),
	7304:  uint16(508),
	7305:  uint16(4),
	7306:  uint16(anon_sym_EQ_EQ),
	7307:  uint16(anon_sym_BANG_EQ),
	7308:  uint16(anon_sym_LT_EQ),
	7309:  uint16(anon_sym_GT_EQ),
	7310:  uint16(494),
	7311:  uint16(7),
	7312:  uint16(anon_sym_SEMI),
	7313:  uint16(anon_sym_RBRACK),
	7314:  uint16(anon_sym_COMMA),
	7315:  uint16(anon_sym_LPAREN),
	7316:  uint16(anon_sym_RPAREN),
	7317:  uint16(anon_sym_AMP_AMP),
	7318:  uint16(anon_sym_PIPE_PIPE),
	7319:  uint16(4),
	7320:  uint16(3),
	7321:  uint16(1),
	7322:  uint16(sym_comment),
	7323:  uint16(518),
	7324:  uint16(1),
	7325:  uint16(anon_sym_DOT),
	7326:  uint16(522),
	7327:  uint16(6),
	7328:  uint16(anon_sym_EQ),
	7329:  uint16(anon_sym_LT),
	7330:  uint16(anon_sym_GT),
	7331:  uint16(anon_sym_AMP),
	7332:  uint16(anon_sym_PIPE),
	7333:  uint16(anon_sym_SLASH),
	7334:  uint16(520),
	7335:  uint16(19),
	7336:  uint16(anon_sym_SEMI),
	7337:  uint16(anon_sym_RBRACK),
	7338:  uint16(anon_sym_COMMA),
	7339:  uint16(anon_sym_LPAREN),
	7340:  uint16(anon_sym_RPAREN),
	7341:  uint16(anon_sym_as),
	7342:  uint16(anon_sym_AMP_AMP),
	7343:  uint16(anon_sym_PIPE_PIPE),
	7344:  uint16(anon_sym_CARET),
	7345:  uint16(anon_sym_EQ_EQ),
	7346:  uint16(anon_sym_BANG_EQ),
	7347:  uint16(anon_sym_LT_EQ),
	7348:  uint16(anon_sym_GT_EQ),
	7349:  uint16(anon_sym_LT_LT),
	7350:  uint16(anon_sym_GT_GT),
	7351:  uint16(anon_sym_PLUS),
	7352:  uint16(anon_sym_DASH),
	7353:  uint16(anon_sym_STAR),
	7354:  uint16(anon_sym_PERCENT),
	7355:  uint16(4),
	7356:  uint16(3),
	7357:  uint16(1),
	7358:  uint16(sym_comment),
	7359:  uint16(518),
	7360:  uint16(1),
	7361:  uint16(anon_sym_DOT),
	7362:  uint16(526),
	7363:  uint16(6),
	7364:  uint16(anon_sym_EQ),
	7365:  uint16(anon_sym_LT),
	7366:  uint16(anon_sym_GT),
	7367:  uint16(anon_sym_AMP),
	7368:  uint16(anon_sym_PIPE),
	7369:  uint16(anon_sym_SLASH),
	7370:  uint16(524),
	7371:  uint16(19),
	7372:  uint16(anon_sym_SEMI),
	7373:  uint16(anon_sym_RBRACK),
	7374:  uint16(anon_sym_COMMA),
	7375:  uint16(anon_sym_LPAREN),
	7376:  uint16(anon_sym_RPAREN),
	7377:  uint16(anon_sym_as),
	7378:  uint16(anon_sym_AMP_AMP),
	7379:  uint16(anon_sym_PIPE_PIPE),
	7380:  uint16(anon_sym_CARET),
	7381:  uint16(anon_sym_EQ_EQ),
	7382:  uint16(anon_sym_BANG_EQ),
	7383:  uint16(anon_sym_LT_EQ),
	7384:  uint16(anon_sym_GT_EQ),
	7385:  uint16(anon_sym_LT_LT),
	7386:  uint16(anon_sym_GT_GT),
	7387:  uint16(anon_sym_PLUS),
	7388:  uint16(anon_sym_DASH),
	7389:  uint16(anon_sym_STAR),
	7390:  uint16(anon_sym_PERCENT),
	7391:  uint16(16),
	7392:  uint16(3),
	7393:  uint16(1),
	7394:  uint16(sym_comment),
	7395:  uint16(500),
	7396:  uint16(1),
	7397:  uint16(anon_sym_as),
	7398:  uint16(502),
	7399:  uint16(1),
	7400:  uint16(anon_sym_AMP),
	7401:  uint16(504),
	7402:  uint16(1),
	7403:  uint16(anon_sym_PIPE),
	7404:  uint16(506),
	7405:  uint16(1),
	7406:  uint16(anon_sym_CARET),
	7407:  uint16(516),
	7408:  uint16(1),
	7409:  uint16(anon_sym_SLASH),
	7410:  uint16(518),
	7411:  uint16(1),
	7412:  uint16(anon_sym_DOT),
	7413:  uint16(530),
	7414:  uint16(1),
	7415:  uint16(anon_sym_EQ),
	7416:  uint16(532),
	7417:  uint16(1),
	7418:  uint16(anon_sym_AMP_AMP),
	7419:  uint16(534),
	7420:  uint16(1),
	7421:  uint16(anon_sym_PIPE_PIPE),
	7422:  uint16(498),
	7423:  uint16(2),
	7424:  uint16(anon_sym_LT),
	7425:  uint16(anon_sym_GT),
	7426:  uint16(510),
	7427:  uint16(2),
	7428:  uint16(anon_sym_LT_LT),
	7429:  uint16(anon_sym_GT_GT),
	7430:  uint16(512),
	7431:  uint16(2),
	7432:  uint16(anon_sym_PLUS),
	7433:  uint16(anon_sym_DASH),
	7434:  uint16(514),
	7435:  uint16(2),
	7436:  uint16(anon_sym_STAR),
	7437:  uint16(anon_sym_PERCENT),
	7438:  uint16(508),
	7439:  uint16(4),
	7440:  uint16(anon_sym_EQ_EQ),
	7441:  uint16(anon_sym_BANG_EQ),
	7442:  uint16(anon_sym_LT_EQ),
	7443:  uint16(anon_sym_GT_EQ),
	7444:  uint16(528),
	7445:  uint16(5),
	7446:  uint16(anon_sym_SEMI),
	7447:  uint16(anon_sym_RBRACK),
	7448:  uint16(anon_sym_COMMA),
	7449:  uint16(anon_sym_LPAREN),
	7450:  uint16(anon_sym_RPAREN),
	7451:  uint16(9),
	7452:  uint16(3),
	7453:  uint16(1),
	7454:  uint16(sym_comment),
	7455:  uint16(500),
	7456:  uint16(1),
	7457:  uint16(anon_sym_as),
	7458:  uint16(516),
	7459:  uint16(1),
	7460:  uint16(anon_sym_SLASH),
	7461:  uint16(518),
	7462:  uint16(1),
	7463:  uint16(anon_sym_DOT),
	7464:  uint16(510),
	7465:  uint16(2),
	7466:  uint16(anon_sym_LT_LT),
	7467:  uint16(anon_sym_GT_GT),
	7468:  uint16(512),
	7469:  uint16(2),
	7470:  uint16(anon_sym_PLUS),
	7471:  uint16(anon_sym_DASH),
	7472:  uint16(514),
	7473:  uint16(2),
	7474:  uint16(anon_sym_STAR),
	7475:  uint16(anon_sym_PERCENT),
	7476:  uint16(496),
	7477:  uint16(5),
	7478:  uint16(anon_sym_EQ),
	7479:  uint16(anon_sym_LT),
	7480:  uint16(anon_sym_GT),
	7481:  uint16(anon_sym_AMP),
	7482:  uint16(anon_sym_PIPE),
	7483:  uint16(494),
	7484:  uint16(12),
	7485:  uint16(anon_sym_SEMI),
	7486:  uint16(anon_sym_RBRACK),
	7487:  uint16(anon_sym_COMMA),
	7488:  uint16(anon_sym_LPAREN),
	7489:  uint16(anon_sym_RPAREN),
	7490:  uint16(anon_sym_AMP_AMP),
	7491:  uint16(anon_sym_PIPE_PIPE),
	7492:  uint16(anon_sym_CARET),
	7493:  uint16(anon_sym_EQ_EQ),
	7494:  uint16(anon_sym_BANG_EQ),
	7495:  uint16(anon_sym_LT_EQ),
	7496:  uint16(anon_sym_GT_EQ),
	7497:  uint16(3),
	7498:  uint16(3),
	7499:  uint16(1),
	7500:  uint16(sym_comment),
	7501:  uint16(538),
	7502:  uint16(6),
	7503:  uint16(anon_sym_EQ),
	7504:  uint16(anon_sym_LT),
	7505:  uint16(anon_sym_GT),
	7506:  uint16(anon_sym_AMP),
	7507:  uint16(anon_sym_PIPE),
	7508:  uint16(anon_sym_SLASH),
	7509:  uint16(536),
	7510:  uint16(20),
	7511:  uint16(anon_sym_SEMI),
	7512:  uint16(anon_sym_RBRACK),
	7513:  uint16(anon_sym_COMMA),
	7514:  uint16(anon_sym_LPAREN),
	7515:  uint16(anon_sym_RPAREN),
	7516:  uint16(anon_sym_as),
	7517:  uint16(anon_sym_AMP_AMP),
	7518:  uint16(anon_sym_PIPE_PIPE),
	7519:  uint16(anon_sym_CARET),
	7520:  uint16(anon_sym_EQ_EQ),
	7521:  uint16(anon_sym_BANG_EQ),
	7522:  uint16(anon_sym_LT_EQ),
	7523:  uint16(anon_sym_GT_EQ),
	7524:  uint16(anon_sym_LT_LT),
	7525:  uint16(anon_sym_GT_GT),
	7526:  uint16(anon_sym_PLUS),
	7527:  uint16(anon_sym_DASH),
	7528:  uint16(anon_sym_STAR),
	7529:  uint16(anon_sym_PERCENT),
	7530:  uint16(anon_sym_DOT),
	7531:  uint16(3),
	7532:  uint16(3),
	7533:  uint16(1),
	7534:  uint16(sym_comment),
	7535:  uint16(542),
	7536:  uint16(6),
	7537:  uint16(anon_sym_EQ),
	7538:  uint16(anon_sym_LT),
	7539:  uint16(anon_sym_GT),
	7540:  uint16(anon_sym_AMP),
	7541:  uint16(anon_sym_PIPE),
	7542:  uint16(anon_sym_SLASH),
	7543:  uint16(540),
	7544:  uint16(20),
	7545:  uint16(anon_sym_SEMI),
	7546:  uint16(anon_sym_RBRACK),
	7547:  uint16(anon_sym_COMMA),
	7548:  uint16(anon_sym_LPAREN),
	7549:  uint16(anon_sym_RPAREN),
	7550:  uint16(anon_sym_as),
	7551:  uint16(anon_sym_AMP_AMP),
	7552:  uint16(anon_sym_PIPE_PIPE),
	7553:  uint16(anon_sym_CARET),
	7554:  uint16(anon_sym_EQ_EQ),
	7555:  uint16(anon_sym_BANG_EQ),
	7556:  uint16(anon_sym_LT_EQ),
	7557:  uint16(anon_sym_GT_EQ),
	7558:  uint16(anon_sym_LT_LT),
	7559:  uint16(anon_sym_GT_GT),
	7560:  uint16(anon_sym_PLUS),
	7561:  uint16(anon_sym_DASH),
	7562:  uint16(anon_sym_STAR),
	7563:  uint16(anon_sym_PERCENT),
	7564:  uint16(anon_sym_DOT),
	7565:  uint16(3),
	7566:  uint16(3),
	7567:  uint16(1),
	7568:  uint16(sym_comment),
	7569:  uint16(546),
	7570:  uint16(6),
	7571:  uint16(anon_sym_EQ),
	7572:  uint16(anon_sym_LT),
	7573:  uint16(anon_sym_GT),
	7574:  uint16(anon_sym_AMP),
	7575:  uint16(anon_sym_PIPE),
	7576:  uint16(anon_sym_SLASH),
	7577:  uint16(544),
	7578:  uint16(20),
	7579:  uint16(anon_sym_SEMI),
	7580:  uint16(anon_sym_RBRACK),
	7581:  uint16(anon_sym_COMMA),
	7582:  uint16(anon_sym_LPAREN),
	7583:  uint16(anon_sym_RPAREN),
	7584:  uint16(anon_sym_as),
	7585:  uint16(anon_sym_AMP_AMP),
	7586:  uint16(anon_sym_PIPE_PIPE),
	7587:  uint16(anon_sym_CARET),
	7588:  uint16(anon_sym_EQ_EQ),
	7589:  uint16(anon_sym_BANG_EQ),
	7590:  uint16(anon_sym_LT_EQ),
	7591:  uint16(anon_sym_GT_EQ),
	7592:  uint16(anon_sym_LT_LT),
	7593:  uint16(anon_sym_GT_GT),
	7594:  uint16(anon_sym_PLUS),
	7595:  uint16(anon_sym_DASH),
	7596:  uint16(anon_sym_STAR),
	7597:  uint16(anon_sym_PERCENT),
	7598:  uint16(anon_sym_DOT),
	7599:  uint16(3),
	7600:  uint16(3),
	7601:  uint16(1),
	7602:  uint16(sym_comment),
	7603:  uint16(550),
	7604:  uint16(6),
	7605:  uint16(anon_sym_EQ),
	7606:  uint16(anon_sym_LT),
	7607:  uint16(anon_sym_GT),
	7608:  uint16(anon_sym_AMP),
	7609:  uint16(anon_sym_PIPE),
	7610:  uint16(anon_sym_SLASH),
	7611:  uint16(548),
	7612:  uint16(20),
	7613:  uint16(anon_sym_SEMI),
	7614:  uint16(anon_sym_RBRACK),
	7615:  uint16(anon_sym_COMMA),
	7616:  uint16(anon_sym_LPAREN),
	7617:  uint16(anon_sym_RPAREN),
	7618:  uint16(anon_sym_as),
	7619:  uint16(anon_sym_AMP_AMP),
	7620:  uint16(anon_sym_PIPE_PIPE),
	7621:  uint16(anon_sym_CARET),
	7622:  uint16(anon_sym_EQ_EQ),
	7623:  uint16(anon_sym_BANG_EQ),
	7624:  uint16(anon_sym_LT_EQ),
	7625:  uint16(anon_sym_GT_EQ),
	7626:  uint16(anon_sym_LT_LT),
	7627:  uint16(anon_sym_GT_GT),
	7628:  uint16(anon_sym_PLUS),
	7629:  uint16(anon_sym_DASH),
	7630:  uint16(anon_sym_STAR),
	7631:  uint16(anon_sym_PERCENT),
	7632:  uint16(anon_sym_DOT),
	7633:  uint16(12),
	7634:  uint16(3),
	7635:  uint16(1),
	7636:  uint16(sym_comment),
	7637:  uint16(500),
	7638:  uint16(1),
	7639:  uint16(anon_sym_as),
	7640:  uint16(502),
	7641:  uint16(1),
	7642:  uint16(anon_sym_AMP),
	7643:  uint16(504),
	7644:  uint16(1),
	7645:  uint16(anon_sym_PIPE),
	7646:  uint16(506),
	7647:  uint16(1),
	7648:  uint16(anon_sym_CARET),
	7649:  uint16(516),
	7650:  uint16(1),
	7651:  uint16(anon_sym_SLASH),
	7652:  uint16(518),
	7653:  uint16(1),
	7654:  uint16(anon_sym_DOT),
	7655:  uint16(510),
	7656:  uint16(2),
	7657:  uint16(anon_sym_LT_LT),
	7658:  uint16(anon_sym_GT_GT),
	7659:  uint16(512),
	7660:  uint16(2),
	7661:  uint16(anon_sym_PLUS),
	7662:  uint16(anon_sym_DASH),
	7663:  uint16(514),
	7664:  uint16(2),
	7665:  uint16(anon_sym_STAR),
	7666:  uint16(anon_sym_PERCENT),
	7667:  uint16(496),
	7668:  uint16(3),
	7669:  uint16(anon_sym_EQ),
	7670:  uint16(anon_sym_LT),
	7671:  uint16(anon_sym_GT),
	7672:  uint16(494),
	7673:  uint16(11),
	7674:  uint16(anon_sym_SEMI),
	7675:  uint16(anon_sym_RBRACK),
	7676:  uint16(anon_sym_COMMA),
	7677:  uint16(anon_sym_LPAREN),
	7678:  uint16(anon_sym_RPAREN),
	7679:  uint16(anon_sym_AMP_AMP),
	7680:  uint16(anon_sym_PIPE_PIPE),
	7681:  uint16(anon_sym_EQ_EQ),
	7682:  uint16(anon_sym_BANG_EQ),
	7683:  uint16(anon_sym_LT_EQ),
	7684:  uint16(anon_sym_GT_EQ),
	7685:  uint16(3),
	7686:  uint16(3),
	7687:  uint16(1),
	7688:  uint16(sym_comment),
	7689:  uint16(554),
	7690:  uint16(6),
	7691:  uint16(anon_sym_EQ),
	7692:  uint16(anon_sym_LT),
	7693:  uint16(anon_sym_GT),
	7694:  uint16(anon_sym_AMP),
	7695:  uint16(anon_sym_PIPE),
	7696:  uint16(anon_sym_SLASH),
	7697:  uint16(552),
	7698:  uint16(20),
	7699:  uint16(anon_sym_SEMI),
	7700:  uint16(anon_sym_RBRACK),
	7701:  uint16(anon_sym_COMMA),
	7702:  uint16(anon_sym_LPAREN),
	7703:  uint16(anon_sym_RPAREN),
	7704:  uint16(anon_sym_as),
	7705:  uint16(anon_sym_AMP_AMP),
	7706:  uint16(anon_sym_PIPE_PIPE),
	7707:  uint16(anon_sym_CARET),
	7708:  uint16(anon_sym_EQ_EQ),
	7709:  uint16(anon_sym_BANG_EQ),
	7710:  uint16(anon_sym_LT_EQ),
	7711:  uint16(anon_sym_GT_EQ),
	7712:  uint16(anon_sym_LT_LT),
	7713:  uint16(anon_sym_GT_GT),
	7714:  uint16(anon_sym_PLUS),
	7715:  uint16(anon_sym_DASH),
	7716:  uint16(anon_sym_STAR),
	7717:  uint16(anon_sym_PERCENT),
	7718:  uint16(anon_sym_DOT),
	7719:  uint16(15),
	7720:  uint16(3),
	7721:  uint16(1),
	7722:  uint16(sym_comment),
	7723:  uint16(496),
	7724:  uint16(1),
	7725:  uint16(anon_sym_EQ),
	7726:  uint16(500),
	7727:  uint16(1),
	7728:  uint16(anon_sym_as),
	7729:  uint16(502),
	7730:  uint16(1),
	7731:  uint16(anon_sym_AMP),
	7732:  uint16(504),
	7733:  uint16(1),
	7734:  uint16(anon_sym_PIPE),
	7735:  uint16(506),
	7736:  uint16(1),
	7737:  uint16(anon_sym_CARET),
	7738:  uint16(516),
	7739:  uint16(1),
	7740:  uint16(anon_sym_SLASH),
	7741:  uint16(518),
	7742:  uint16(1),
	7743:  uint16(anon_sym_DOT),
	7744:  uint16(532),
	7745:  uint16(1),
	7746:  uint16(anon_sym_AMP_AMP),
	7747:  uint16(498),
	7748:  uint16(2),
	7749:  uint16(anon_sym_LT),
	7750:  uint16(anon_sym_GT),
	7751:  uint16(510),
	7752:  uint16(2),
	7753:  uint16(anon_sym_LT_LT),
	7754:  uint16(anon_sym_GT_GT),
	7755:  uint16(512),
	7756:  uint16(2),
	7757:  uint16(anon_sym_PLUS),
	7758:  uint16(anon_sym_DASH),
	7759:  uint16(514),
	7760:  uint16(2),
	7761:  uint16(anon_sym_STAR),
	7762:  uint16(anon_sym_PERCENT),
	7763:  uint16(508),
	7764:  uint16(4),
	7765:  uint16(anon_sym_EQ_EQ),
	7766:  uint16(anon_sym_BANG_EQ),
	7767:  uint16(anon_sym_LT_EQ),
	7768:  uint16(anon_sym_GT_EQ),
	7769:  uint16(494),
	7770:  uint16(6),
	7771:  uint16(anon_sym_SEMI),
	7772:  uint16(anon_sym_RBRACK),
	7773:  uint16(anon_sym_COMMA),
	7774:  uint16(anon_sym_LPAREN),
	7775:  uint16(anon_sym_RPAREN),
	7776:  uint16(anon_sym_PIPE_PIPE),
	7777:  uint16(3),
	7778:  uint16(3),
	7779:  uint16(1),
	7780:  uint16(sym_comment),
	7781:  uint16(254),
	7782:  uint16(6),
	7783:  uint16(anon_sym_EQ),
	7784:  uint16(anon_sym_LT),
	7785:  uint16(anon_sym_GT),
	7786:  uint16(anon_sym_AMP),
	7787:  uint16(anon_sym_PIPE),
	7788:  uint16(anon_sym_SLASH),
	7789:  uint16(256),
	7790:  uint16(20),
	7791:  uint16(anon_sym_SEMI),
	7792:  uint16(anon_sym_RBRACK),
	7793:  uint16(anon_sym_COMMA),
	7794:  uint16(anon_sym_LPAREN),
	7795:  uint16(anon_sym_RPAREN),
	7796:  uint16(anon_sym_as),
	7797:  uint16(anon_sym_AMP_AMP),
	7798:  uint16(anon_sym_PIPE_PIPE),
	7799:  uint16(anon_sym_CARET),
	7800:  uint16(anon_sym_EQ_EQ),
	7801:  uint16(anon_sym_BANG_EQ),
	7802:  uint16(anon_sym_LT_EQ),
	7803:  uint16(anon_sym_GT_EQ),
	7804:  uint16(anon_sym_LT_LT),
	7805:  uint16(anon_sym_GT_GT),
	7806:  uint16(anon_sym_PLUS),
	7807:  uint16(anon_sym_DASH),
	7808:  uint16(anon_sym_STAR),
	7809:  uint16(anon_sym_PERCENT),
	7810:  uint16(anon_sym_DOT),
	7811:  uint16(16),
	7812:  uint16(3),
	7813:  uint16(1),
	7814:  uint16(sym_comment),
	7815:  uint16(500),
	7816:  uint16(1),
	7817:  uint16(anon_sym_as),
	7818:  uint16(502),
	7819:  uint16(1),
	7820:  uint16(anon_sym_AMP),
	7821:  uint16(504),
	7822:  uint16(1),
	7823:  uint16(anon_sym_PIPE),
	7824:  uint16(506),
	7825:  uint16(1),
	7826:  uint16(anon_sym_CARET),
	7827:  uint16(516),
	7828:  uint16(1),
	7829:  uint16(anon_sym_SLASH),
	7830:  uint16(518),
	7831:  uint16(1),
	7832:  uint16(anon_sym_DOT),
	7833:  uint16(532),
	7834:  uint16(1),
	7835:  uint16(anon_sym_AMP_AMP),
	7836:  uint16(534),
	7837:  uint16(1),
	7838:  uint16(anon_sym_PIPE_PIPE),
	7839:  uint16(558),
	7840:  uint16(1),
	7841:  uint16(anon_sym_EQ),
	7842:  uint16(498),
	7843:  uint16(2),
	7844:  uint16(anon_sym_LT),
	7845:  uint16(anon_sym_GT),
	7846:  uint16(510),
	7847:  uint16(2),
	7848:  uint16(anon_sym_LT_LT),
	7849:  uint16(anon_sym_GT_GT),
	7850:  uint16(512),
	7851:  uint16(2),
	7852:  uint16(anon_sym_PLUS),
	7853:  uint16(anon_sym_DASH),
	7854:  uint16(514),
	7855:  uint16(2),
	7856:  uint16(anon_sym_STAR),
	7857:  uint16(anon_sym_PERCENT),
	7858:  uint16(508),
	7859:  uint16(4),
	7860:  uint16(anon_sym_EQ_EQ),
	7861:  uint16(anon_sym_BANG_EQ),
	7862:  uint16(anon_sym_LT_EQ),
	7863:  uint16(anon_sym_GT_EQ),
	7864:  uint16(556),
	7865:  uint16(5),
	7866:  uint16(anon_sym_SEMI),
	7867:  uint16(anon_sym_RBRACK),
	7868:  uint16(anon_sym_COMMA),
	7869:  uint16(anon_sym_LPAREN),
	7870:  uint16(anon_sym_RPAREN),
	7871:  uint16(3),
	7872:  uint16(3),
	7873:  uint16(1),
	7874:  uint16(sym_comment),
	7875:  uint16(562),
	7876:  uint16(6),
	7877:  uint16(anon_sym_EQ),
	7878:  uint16(anon_sym_LT),
	7879:  uint16(anon_sym_GT),
	7880:  uint16(anon_sym_AMP),
	7881:  uint16(anon_sym_PIPE),
	7882:  uint16(anon_sym_SLASH),
	7883:  uint16(560),
	7884:  uint16(20),
	7885:  uint16(anon_sym_SEMI),
	7886:  uint16(anon_sym_RBRACK),
	7887:  uint16(anon_sym_COMMA),
	7888:  uint16(anon_sym_LPAREN),
	7889:  uint16(anon_sym_RPAREN),
	7890:  uint16(anon_sym_as),
	7891:  uint16(anon_sym_AMP_AMP),
	7892:  uint16(anon_sym_PIPE_PIPE),
	7893:  uint16(anon_sym_CARET),
	7894:  uint16(anon_sym_EQ_EQ),
	7895:  uint16(anon_sym_BANG_EQ),
	7896:  uint16(anon_sym_LT_EQ),
	7897:  uint16(anon_sym_GT_EQ),
	7898:  uint16(anon_sym_LT_LT),
	7899:  uint16(anon_sym_GT_GT),
	7900:  uint16(anon_sym_PLUS),
	7901:  uint16(anon_sym_DASH),
	7902:  uint16(anon_sym_STAR),
	7903:  uint16(anon_sym_PERCENT),
	7904:  uint16(anon_sym_DOT),
	7905:  uint16(3),
	7906:  uint16(3),
	7907:  uint16(1),
	7908:  uint16(sym_comment),
	7909:  uint16(566),
	7910:  uint16(6),
	7911:  uint16(anon_sym_EQ),
	7912:  uint16(anon_sym_LT),
	7913:  uint16(anon_sym_GT),
	7914:  uint16(anon_sym_AMP),
	7915:  uint16(anon_sym_PIPE),
	7916:  uint16(anon_sym_SLASH),
	7917:  uint16(564),
	7918:  uint16(20),
	7919:  uint16(anon_sym_SEMI),
	7920:  uint16(anon_sym_RBRACK),
	7921:  uint16(anon_sym_COMMA),
	7922:  uint16(anon_sym_LPAREN),
	7923:  uint16(anon_sym_RPAREN),
	7924:  uint16(anon_sym_as),
	7925:  uint16(anon_sym_AMP_AMP),
	7926:  uint16(anon_sym_PIPE_PIPE),
	7927:  uint16(anon_sym_CARET),
	7928:  uint16(anon_sym_EQ_EQ),
	7929:  uint16(anon_sym_BANG_EQ),
	7930:  uint16(anon_sym_LT_EQ),
	7931:  uint16(anon_sym_GT_EQ),
	7932:  uint16(anon_sym_LT_LT),
	7933:  uint16(anon_sym_GT_GT),
	7934:  uint16(anon_sym_PLUS),
	7935:  uint16(anon_sym_DASH),
	7936:  uint16(anon_sym_STAR),
	7937:  uint16(anon_sym_PERCENT),
	7938:  uint16(anon_sym_DOT),
	7939:  uint16(4),
	7940:  uint16(3),
	7941:  uint16(1),
	7942:  uint16(sym_comment),
	7943:  uint16(518),
	7944:  uint16(1),
	7945:  uint16(anon_sym_DOT),
	7946:  uint16(570),
	7947:  uint16(6),
	7948:  uint16(anon_sym_EQ),
	7949:  uint16(anon_sym_LT),
	7950:  uint16(anon_sym_GT),
	7951:  uint16(anon_sym_AMP),
	7952:  uint16(anon_sym_PIPE),
	7953:  uint16(anon_sym_SLASH),
	7954:  uint16(568),
	7955:  uint16(19),
	7956:  uint16(anon_sym_SEMI),
	7957:  uint16(anon_sym_RBRACK),
	7958:  uint16(anon_sym_COMMA),
	7959:  uint16(anon_sym_LPAREN),
	7960:  uint16(anon_sym_RPAREN),
	7961:  uint16(anon_sym_as),
	7962:  uint16(anon_sym_AMP_AMP),
	7963:  uint16(anon_sym_PIPE_PIPE),
	7964:  uint16(anon_sym_CARET),
	7965:  uint16(anon_sym_EQ_EQ),
	7966:  uint16(anon_sym_BANG_EQ),
	7967:  uint16(anon_sym_LT_EQ),
	7968:  uint16(anon_sym_GT_EQ),
	7969:  uint16(anon_sym_LT_LT),
	7970:  uint16(anon_sym_GT_GT),
	7971:  uint16(anon_sym_PLUS),
	7972:  uint16(anon_sym_DASH),
	7973:  uint16(anon_sym_STAR),
	7974:  uint16(anon_sym_PERCENT),
	7975:  uint16(3),
	7976:  uint16(3),
	7977:  uint16(1),
	7978:  uint16(sym_comment),
	7979:  uint16(574),
	7980:  uint16(6),
	7981:  uint16(anon_sym_EQ),
	7982:  uint16(anon_sym_LT),
	7983:  uint16(anon_sym_GT),
	7984:  uint16(anon_sym_AMP),
	7985:  uint16(anon_sym_PIPE),
	7986:  uint16(anon_sym_SLASH),
	7987:  uint16(572),
	7988:  uint16(20),
	7989:  uint16(anon_sym_SEMI),
	7990:  uint16(anon_sym_RBRACK),
	7991:  uint16(anon_sym_COMMA),
	7992:  uint16(anon_sym_LPAREN),
	7993:  uint16(anon_sym_RPAREN),
	7994:  uint16(anon_sym_as),
	7995:  uint16(anon_sym_AMP_AMP),
	7996:  uint16(anon_sym_PIPE_PIPE),
	7997:  uint16(anon_sym_CARET),
	7998:  uint16(anon_sym_EQ_EQ),
	7999:  uint16(anon_sym_BANG_EQ),
	8000:  uint16(anon_sym_LT_EQ),
	8001:  uint16(anon_sym_GT_EQ),
	8002:  uint16(anon_sym_LT_LT),
	8003:  uint16(anon_sym_GT_GT),
	8004:  uint16(anon_sym_PLUS),
	8005:  uint16(anon_sym_DASH),
	8006:  uint16(anon_sym_STAR),
	8007:  uint16(anon_sym_PERCENT),
	8008:  uint16(anon_sym_DOT),
	8009:  uint16(3),
	8010:  uint16(3),
	8011:  uint16(1),
	8012:  uint16(sym_comment),
	8013:  uint16(578),
	8014:  uint16(6),
	8015:  uint16(anon_sym_EQ),
	8016:  uint16(anon_sym_LT),
	8017:  uint16(anon_sym_GT),
	8018:  uint16(anon_sym_AMP),
	8019:  uint16(anon_sym_PIPE),
	8020:  uint16(anon_sym_SLASH),
	8021:  uint16(576),
	8022:  uint16(20),
	8023:  uint16(anon_sym_SEMI),
	8024:  uint16(anon_sym_RBRACK),
	8025:  uint16(anon_sym_COMMA),
	8026:  uint16(anon_sym_LPAREN),
	8027:  uint16(anon_sym_RPAREN),
	8028:  uint16(anon_sym_as),
	8029:  uint16(anon_sym_AMP_AMP),
	8030:  uint16(anon_sym_PIPE_PIPE),
	8031:  uint16(anon_sym_CARET),
	8032:  uint16(anon_sym_EQ_EQ),
	8033:  uint16(anon_sym_BANG_EQ),
	8034:  uint16(anon_sym_LT_EQ),
	8035:  uint16(anon_sym_GT_EQ),
	8036:  uint16(anon_sym_LT_LT),
	8037:  uint16(anon_sym_GT_GT),
	8038:  uint16(anon_sym_PLUS),
	8039:  uint16(anon_sym_DASH),
	8040:  uint16(anon_sym_STAR),
	8041:  uint16(anon_sym_PERCENT),
	8042:  uint16(anon_sym_DOT),
	8043:  uint16(3),
	8044:  uint16(3),
	8045:  uint16(1),
	8046:  uint16(sym_comment),
	8047:  uint16(582),
	8048:  uint16(6),
	8049:  uint16(anon_sym_EQ),
	8050:  uint16(anon_sym_LT),
	8051:  uint16(anon_sym_GT),
	8052:  uint16(anon_sym_AMP),
	8053:  uint16(anon_sym_PIPE),
	8054:  uint16(anon_sym_SLASH),
	8055:  uint16(580),
	8056:  uint16(20),
	8057:  uint16(anon_sym_SEMI),
	8058:  uint16(anon_sym_RBRACK),
	8059:  uint16(anon_sym_COMMA),
	8060:  uint16(anon_sym_LPAREN),
	8061:  uint16(anon_sym_RPAREN),
	8062:  uint16(anon_sym_as),
	8063:  uint16(anon_sym_AMP_AMP),
	8064:  uint16(anon_sym_PIPE_PIPE),
	8065:  uint16(anon_sym_CARET),
	8066:  uint16(anon_sym_EQ_EQ),
	8067:  uint16(anon_sym_BANG_EQ),
	8068:  uint16(anon_sym_LT_EQ),
	8069:  uint16(anon_sym_GT_EQ),
	8070:  uint16(anon_sym_LT_LT),
	8071:  uint16(anon_sym_GT_GT),
	8072:  uint16(anon_sym_PLUS),
	8073:  uint16(anon_sym_DASH),
	8074:  uint16(anon_sym_STAR),
	8075:  uint16(anon_sym_PERCENT),
	8076:  uint16(anon_sym_DOT),
	8077:  uint16(3),
	8078:  uint16(3),
	8079:  uint16(1),
	8080:  uint16(sym_comment),
	8081:  uint16(586),
	8082:  uint16(6),
	8083:  uint16(anon_sym_EQ),
	8084:  uint16(anon_sym_LT),
	8085:  uint16(anon_sym_GT),
	8086:  uint16(anon_sym_AMP),
	8087:  uint16(anon_sym_PIPE),
	8088:  uint16(anon_sym_SLASH),
	8089:  uint16(584),
	8090:  uint16(20),
	8091:  uint16(anon_sym_SEMI),
	8092:  uint16(anon_sym_RBRACK),
	8093:  uint16(anon_sym_COMMA),
	8094:  uint16(anon_sym_LPAREN),
	8095:  uint16(anon_sym_RPAREN),
	8096:  uint16(anon_sym_as),
	8097:  uint16(anon_sym_AMP_AMP),
	8098:  uint16(anon_sym_PIPE_PIPE),
	8099:  uint16(anon_sym_CARET),
	8100:  uint16(anon_sym_EQ_EQ),
	8101:  uint16(anon_sym_BANG_EQ),
	8102:  uint16(anon_sym_LT_EQ),
	8103:  uint16(anon_sym_GT_EQ),
	8104:  uint16(anon_sym_LT_LT),
	8105:  uint16(anon_sym_GT_GT),
	8106:  uint16(anon_sym_PLUS),
	8107:  uint16(anon_sym_DASH),
	8108:  uint16(anon_sym_STAR),
	8109:  uint16(anon_sym_PERCENT),
	8110:  uint16(anon_sym_DOT),
	8111:  uint16(3),
	8112:  uint16(3),
	8113:  uint16(1),
	8114:  uint16(sym_comment),
	8115:  uint16(258),
	8116:  uint16(6),
	8117:  uint16(anon_sym_EQ),
	8118:  uint16(anon_sym_LT),
	8119:  uint16(anon_sym_GT),
	8120:  uint16(anon_sym_AMP),
	8121:  uint16(anon_sym_PIPE),
	8122:  uint16(anon_sym_SLASH),
	8123:  uint16(260),
	8124:  uint16(20),
	8125:  uint16(anon_sym_SEMI),
	8126:  uint16(anon_sym_RBRACK),
	8127:  uint16(anon_sym_COMMA),
	8128:  uint16(anon_sym_LPAREN),
	8129:  uint16(anon_sym_RPAREN),
	8130:  uint16(anon_sym_as),
	8131:  uint16(anon_sym_AMP_AMP),
	8132:  uint16(anon_sym_PIPE_PIPE),
	8133:  uint16(anon_sym_CARET),
	8134:  uint16(anon_sym_EQ_EQ),
	8135:  uint16(anon_sym_BANG_EQ),
	8136:  uint16(anon_sym_LT_EQ),
	8137:  uint16(anon_sym_GT_EQ),
	8138:  uint16(anon_sym_LT_LT),
	8139:  uint16(anon_sym_GT_GT),
	8140:  uint16(anon_sym_PLUS),
	8141:  uint16(anon_sym_DASH),
	8142:  uint16(anon_sym_STAR),
	8143:  uint16(anon_sym_PERCENT),
	8144:  uint16(anon_sym_DOT),
	8145:  uint16(3),
	8146:  uint16(3),
	8147:  uint16(1),
	8148:  uint16(sym_comment),
	8149:  uint16(262),
	8150:  uint16(6),
	8151:  uint16(anon_sym_EQ),
	8152:  uint16(anon_sym_LT),
	8153:  uint16(anon_sym_GT),
	8154:  uint16(anon_sym_AMP),
	8155:  uint16(anon_sym_PIPE),
	8156:  uint16(anon_sym_SLASH),
	8157:  uint16(264),
	8158:  uint16(20),
	8159:  uint16(anon_sym_SEMI),
	8160:  uint16(anon_sym_RBRACK),
	8161:  uint16(anon_sym_COMMA),
	8162:  uint16(anon_sym_LPAREN),
	8163:  uint16(anon_sym_RPAREN),
	8164:  uint16(anon_sym_as),
	8165:  uint16(anon_sym_AMP_AMP),
	8166:  uint16(anon_sym_PIPE_PIPE),
	8167:  uint16(anon_sym_CARET),
	8168:  uint16(anon_sym_EQ_EQ),
	8169:  uint16(anon_sym_BANG_EQ),
	8170:  uint16(anon_sym_LT_EQ),
	8171:  uint16(anon_sym_GT_EQ),
	8172:  uint16(anon_sym_LT_LT),
	8173:  uint16(anon_sym_GT_GT),
	8174:  uint16(anon_sym_PLUS),
	8175:  uint16(anon_sym_DASH),
	8176:  uint16(anon_sym_STAR),
	8177:  uint16(anon_sym_PERCENT),
	8178:  uint16(anon_sym_DOT),
	8179:  uint16(3),
	8180:  uint16(3),
	8181:  uint16(1),
	8182:  uint16(sym_comment),
	8183:  uint16(230),
	8184:  uint16(6),
	8185:  uint16(anon_sym_EQ),
	8186:  uint16(anon_sym_LT),
	8187:  uint16(anon_sym_GT),
	8188:  uint16(anon_sym_AMP),
	8189:  uint16(anon_sym_PIPE),
	8190:  uint16(anon_sym_SLASH),
	8191:  uint16(232),
	8192:  uint16(20),
	8193:  uint16(anon_sym_SEMI),
	8194:  uint16(anon_sym_RBRACK),
	8195:  uint16(anon_sym_COMMA),
	8196:  uint16(anon_sym_LPAREN),
	8197:  uint16(anon_sym_RPAREN),
	8198:  uint16(anon_sym_as),
	8199:  uint16(anon_sym_AMP_AMP),
	8200:  uint16(anon_sym_PIPE_PIPE),
	8201:  uint16(anon_sym_CARET),
	8202:  uint16(anon_sym_EQ_EQ),
	8203:  uint16(anon_sym_BANG_EQ),
	8204:  uint16(anon_sym_LT_EQ),
	8205:  uint16(anon_sym_GT_EQ),
	8206:  uint16(anon_sym_LT_LT),
	8207:  uint16(anon_sym_GT_GT),
	8208:  uint16(anon_sym_PLUS),
	8209:  uint16(anon_sym_DASH),
	8210:  uint16(anon_sym_STAR),
	8211:  uint16(anon_sym_PERCENT),
	8212:  uint16(anon_sym_DOT),
	8213:  uint16(3),
	8214:  uint16(3),
	8215:  uint16(1),
	8216:  uint16(sym_comment),
	8217:  uint16(590),
	8218:  uint16(6),
	8219:  uint16(anon_sym_EQ),
	8220:  uint16(anon_sym_LT),
	8221:  uint16(anon_sym_GT),
	8222:  uint16(anon_sym_AMP),
	8223:  uint16(anon_sym_PIPE),
	8224:  uint16(anon_sym_SLASH),
	8225:  uint16(588),
	8226:  uint16(20),
	8227:  uint16(anon_sym_SEMI),
	8228:  uint16(anon_sym_RBRACK),
	8229:  uint16(anon_sym_COMMA),
	8230:  uint16(anon_sym_LPAREN),
	8231:  uint16(anon_sym_RPAREN),
	8232:  uint16(anon_sym_as),
	8233:  uint16(anon_sym_AMP_AMP),
	8234:  uint16(anon_sym_PIPE_PIPE),
	8235:  uint16(anon_sym_CARET),
	8236:  uint16(anon_sym_EQ_EQ),
	8237:  uint16(anon_sym_BANG_EQ),
	8238:  uint16(anon_sym_LT_EQ),
	8239:  uint16(anon_sym_GT_EQ),
	8240:  uint16(anon_sym_LT_LT),
	8241:  uint16(anon_sym_GT_GT),
	8242:  uint16(anon_sym_PLUS),
	8243:  uint16(anon_sym_DASH),
	8244:  uint16(anon_sym_STAR),
	8245:  uint16(anon_sym_PERCENT),
	8246:  uint16(anon_sym_DOT),
	8247:  uint16(3),
	8248:  uint16(3),
	8249:  uint16(1),
	8250:  uint16(sym_comment),
	8251:  uint16(594),
	8252:  uint16(6),
	8253:  uint16(anon_sym_EQ),
	8254:  uint16(anon_sym_LT),
	8255:  uint16(anon_sym_GT),
	8256:  uint16(anon_sym_AMP),
	8257:  uint16(anon_sym_PIPE),
	8258:  uint16(anon_sym_SLASH),
	8259:  uint16(592),
	8260:  uint16(20),
	8261:  uint16(anon_sym_SEMI),
	8262:  uint16(anon_sym_RBRACK),
	8263:  uint16(anon_sym_COMMA),
	8264:  uint16(anon_sym_LPAREN),
	8265:  uint16(anon_sym_RPAREN),
	8266:  uint16(anon_sym_as),
	8267:  uint16(anon_sym_AMP_AMP),
	8268:  uint16(anon_sym_PIPE_PIPE),
	8269:  uint16(anon_sym_CARET),
	8270:  uint16(anon_sym_EQ_EQ),
	8271:  uint16(anon_sym_BANG_EQ),
	8272:  uint16(anon_sym_LT_EQ),
	8273:  uint16(anon_sym_GT_EQ),
	8274:  uint16(anon_sym_LT_LT),
	8275:  uint16(anon_sym_GT_GT),
	8276:  uint16(anon_sym_PLUS),
	8277:  uint16(anon_sym_DASH),
	8278:  uint16(anon_sym_STAR),
	8279:  uint16(anon_sym_PERCENT),
	8280:  uint16(anon_sym_DOT),
	8281:  uint16(3),
	8282:  uint16(3),
	8283:  uint16(1),
	8284:  uint16(sym_comment),
	8285:  uint16(250),
	8286:  uint16(6),
	8287:  uint16(anon_sym_EQ),
	8288:  uint16(anon_sym_LT),
	8289:  uint16(anon_sym_GT),
	8290:  uint16(anon_sym_AMP),
	8291:  uint16(anon_sym_PIPE),
	8292:  uint16(anon_sym_SLASH),
	8293:  uint16(252),
	8294:  uint16(20),
	8295:  uint16(anon_sym_SEMI),
	8296:  uint16(anon_sym_RBRACK),
	8297:  uint16(anon_sym_COMMA),
	8298:  uint16(anon_sym_LPAREN),
	8299:  uint16(anon_sym_RPAREN),
	8300:  uint16(anon_sym_as),
	8301:  uint16(anon_sym_AMP_AMP),
	8302:  uint16(anon_sym_PIPE_PIPE),
	8303:  uint16(anon_sym_CARET),
	8304:  uint16(anon_sym_EQ_EQ),
	8305:  uint16(anon_sym_BANG_EQ),
	8306:  uint16(anon_sym_LT_EQ),
	8307:  uint16(anon_sym_GT_EQ),
	8308:  uint16(anon_sym_LT_LT),
	8309:  uint16(anon_sym_GT_GT),
	8310:  uint16(anon_sym_PLUS),
	8311:  uint16(anon_sym_DASH),
	8312:  uint16(anon_sym_STAR),
	8313:  uint16(anon_sym_PERCENT),
	8314:  uint16(anon_sym_DOT),
	8315:  uint16(3),
	8316:  uint16(3),
	8317:  uint16(1),
	8318:  uint16(sym_comment),
	8319:  uint16(598),
	8320:  uint16(6),
	8321:  uint16(anon_sym_EQ),
	8322:  uint16(anon_sym_LT),
	8323:  uint16(anon_sym_GT),
	8324:  uint16(anon_sym_AMP),
	8325:  uint16(anon_sym_PIPE),
	8326:  uint16(anon_sym_SLASH),
	8327:  uint16(596),
	8328:  uint16(20),
	8329:  uint16(anon_sym_SEMI),
	8330:  uint16(anon_sym_RBRACK),
	8331:  uint16(anon_sym_COMMA),
	8332:  uint16(anon_sym_LPAREN),
	8333:  uint16(anon_sym_RPAREN),
	8334:  uint16(anon_sym_as),
	8335:  uint16(anon_sym_AMP_AMP),
	8336:  uint16(anon_sym_PIPE_PIPE),
	8337:  uint16(anon_sym_CARET),
	8338:  uint16(anon_sym_EQ_EQ),
	8339:  uint16(anon_sym_BANG_EQ),
	8340:  uint16(anon_sym_LT_EQ),
	8341:  uint16(anon_sym_GT_EQ),
	8342:  uint16(anon_sym_LT_LT),
	8343:  uint16(anon_sym_GT_GT),
	8344:  uint16(anon_sym_PLUS),
	8345:  uint16(anon_sym_DASH),
	8346:  uint16(anon_sym_STAR),
	8347:  uint16(anon_sym_PERCENT),
	8348:  uint16(anon_sym_DOT),
	8349:  uint16(5),
	8350:  uint16(3),
	8351:  uint16(1),
	8352:  uint16(sym_comment),
	8353:  uint16(500),
	8354:  uint16(1),
	8355:  uint16(anon_sym_as),
	8356:  uint16(518),
	8357:  uint16(1),
	8358:  uint16(anon_sym_DOT),
	8359:  uint16(496),
	8360:  uint16(6),
	8361:  uint16(anon_sym_EQ),
	8362:  uint16(anon_sym_LT),
	8363:  uint16(anon_sym_GT),
	8364:  uint16(anon_sym_AMP),
	8365:  uint16(anon_sym_PIPE),
	8366:  uint16(anon_sym_SLASH),
	8367:  uint16(494),
	8368:  uint16(18),
	8369:  uint16(anon_sym_SEMI),
	8370:  uint16(anon_sym_RBRACK),
	8371:  uint16(anon_sym_COMMA),
	8372:  uint16(anon_sym_LPAREN),
	8373:  uint16(anon_sym_RPAREN),
	8374:  uint16(anon_sym_AMP_AMP),
	8375:  uint16(anon_sym_PIPE_PIPE),
	8376:  uint16(anon_sym_CARET),
	8377:  uint16(anon_sym_EQ_EQ),
	8378:  uint16(anon_sym_BANG_EQ),
	8379:  uint16(anon_sym_LT_EQ),
	8380:  uint16(anon_sym_GT_EQ),
	8381:  uint16(anon_sym_LT_LT),
	8382:  uint16(anon_sym_GT_GT),
	8383:  uint16(anon_sym_PLUS),
	8384:  uint16(anon_sym_DASH),
	8385:  uint16(anon_sym_STAR),
	8386:  uint16(anon_sym_PERCENT),
	8387:  uint16(7),
	8388:  uint16(3),
	8389:  uint16(1),
	8390:  uint16(sym_comment),
	8391:  uint16(500),
	8392:  uint16(1),
	8393:  uint16(anon_sym_as),
	8394:  uint16(516),
	8395:  uint16(1),
	8396:  uint16(anon_sym_SLASH),
	8397:  uint16(518),
	8398:  uint16(1),
	8399:  uint16(anon_sym_DOT),
	8400:  uint16(514),
	8401:  uint16(2),
	8402:  uint16(anon_sym_STAR),
	8403:  uint16(anon_sym_PERCENT),
	8404:  uint16(496),
	8405:  uint16(5),
	8406:  uint16(anon_sym_EQ),
	8407:  uint16(anon_sym_LT),
	8408:  uint16(anon_sym_GT),
	8409:  uint16(anon_sym_AMP),
	8410:  uint16(anon_sym_PIPE),
	8411:  uint16(494),
	8412:  uint16(16),
	8413:  uint16(anon_sym_SEMI),
	8414:  uint16(anon_sym_RBRACK),
	8415:  uint16(anon_sym_COMMA),
	8416:  uint16(anon_sym_LPAREN),
	8417:  uint16(anon_sym_RPAREN),
	8418:  uint16(anon_sym_AMP_AMP),
	8419:  uint16(anon_sym_PIPE_PIPE),
	8420:  uint16(anon_sym_CARET),
	8421:  uint16(anon_sym_EQ_EQ),
	8422:  uint16(anon_sym_BANG_EQ),
	8423:  uint16(anon_sym_LT_EQ),
	8424:  uint16(anon_sym_GT_EQ),
	8425:  uint16(anon_sym_LT_LT),
	8426:  uint16(anon_sym_GT_GT),
	8427:  uint16(anon_sym_PLUS),
	8428:  uint16(anon_sym_DASH),
	8429:  uint16(8),
	8430:  uint16(3),
	8431:  uint16(1),
	8432:  uint16(sym_comment),
	8433:  uint16(500),
	8434:  uint16(1),
	8435:  uint16(anon_sym_as),
	8436:  uint16(516),
	8437:  uint16(1),
	8438:  uint16(anon_sym_SLASH),
	8439:  uint16(518),
	8440:  uint16(1),
	8441:  uint16(anon_sym_DOT),
	8442:  uint16(512),
	8443:  uint16(2),
	8444:  uint16(anon_sym_PLUS),
	8445:  uint16(anon_sym_DASH),
	8446:  uint16(514),
	8447:  uint16(2),
	8448:  uint16(anon_sym_STAR),
	8449:  uint16(anon_sym_PERCENT),
	8450:  uint16(496),
	8451:  uint16(5),
	8452:  uint16(anon_sym_EQ),
	8453:  uint16(anon_sym_LT),
	8454:  uint16(anon_sym_GT),
	8455:  uint16(anon_sym_AMP),
	8456:  uint16(anon_sym_PIPE),
	8457:  uint16(494),
	8458:  uint16(14),
	8459:  uint16(anon_sym_SEMI),
	8460:  uint16(anon_sym_RBRACK),
	8461:  uint16(anon_sym_COMMA),
	8462:  uint16(anon_sym_LPAREN),
	8463:  uint16(anon_sym_RPAREN),
	8464:  uint16(anon_sym_AMP_AMP),
	8465:  uint16(anon_sym_PIPE_PIPE),
	8466:  uint16(anon_sym_CARET),
	8467:  uint16(anon_sym_EQ_EQ),
	8468:  uint16(anon_sym_BANG_EQ),
	8469:  uint16(anon_sym_LT_EQ),
	8470:  uint16(anon_sym_GT_EQ),
	8471:  uint16(anon_sym_LT_LT),
	8472:  uint16(anon_sym_GT_GT),
	8473:  uint16(10),
	8474:  uint16(3),
	8475:  uint16(1),
	8476:  uint16(sym_comment),
	8477:  uint16(500),
	8478:  uint16(1),
	8479:  uint16(anon_sym_as),
	8480:  uint16(502),
	8481:  uint16(1),
	8482:  uint16(anon_sym_AMP),
	8483:  uint16(516),
	8484:  uint16(1),
	8485:  uint16(anon_sym_SLASH),
	8486:  uint16(518),
	8487:  uint16(1),
	8488:  uint16(anon_sym_DOT),
	8489:  uint16(510),
	8490:  uint16(2),
	8491:  uint16(anon_sym_LT_LT),
	8492:  uint16(anon_sym_GT_GT),
	8493:  uint16(512),
	8494:  uint16(2),
	8495:  uint16(anon_sym_PLUS),
	8496:  uint16(anon_sym_DASH),
	8497:  uint16(514),
	8498:  uint16(2),
	8499:  uint16(anon_sym_STAR),
	8500:  uint16(anon_sym_PERCENT),
	8501:  uint16(496),
	8502:  uint16(4),
	8503:  uint16(anon_sym_EQ),
	8504:  uint16(anon_sym_LT),
	8505:  uint16(anon_sym_GT),
	8506:  uint16(anon_sym_PIPE),
	8507:  uint16(494),
	8508:  uint16(12),
	8509:  uint16(anon_sym_SEMI),
	8510:  uint16(anon_sym_RBRACK),
	8511:  uint16(anon_sym_COMMA),
	8512:  uint16(anon_sym_LPAREN),
	8513:  uint16(anon_sym_RPAREN),
	8514:  uint16(anon_sym_AMP_AMP),
	8515:  uint16(anon_sym_PIPE_PIPE),
	8516:  uint16(anon_sym_CARET),
	8517:  uint16(anon_sym_EQ_EQ),
	8518:  uint16(anon_sym_BANG_EQ),
	8519:  uint16(anon_sym_LT_EQ),
	8520:  uint16(anon_sym_GT_EQ),
	8521:  uint16(11),
	8522:  uint16(3),
	8523:  uint16(1),
	8524:  uint16(sym_comment),
	8525:  uint16(500),
	8526:  uint16(1),
	8527:  uint16(anon_sym_as),
	8528:  uint16(502),
	8529:  uint16(1),
	8530:  uint16(anon_sym_AMP),
	8531:  uint16(506),
	8532:  uint16(1),
	8533:  uint16(anon_sym_CARET),
	8534:  uint16(516),
	8535:  uint16(1),
	8536:  uint16(anon_sym_SLASH),
	8537:  uint16(518),
	8538:  uint16(1),
	8539:  uint16(anon_sym_DOT),
	8540:  uint16(510),
	8541:  uint16(2),
	8542:  uint16(anon_sym_LT_LT),
	8543:  uint16(anon_sym_GT_GT),
	8544:  uint16(512),
	8545:  uint16(2),
	8546:  uint16(anon_sym_PLUS),
	8547:  uint16(anon_sym_DASH),
	8548:  uint16(514),
	8549:  uint16(2),
	8550:  uint16(anon_sym_STAR),
	8551:  uint16(anon_sym_PERCENT),
	8552:  uint16(496),
	8553:  uint16(4),
	8554:  uint16(anon_sym_EQ),
	8555:  uint16(anon_sym_LT),
	8556:  uint16(anon_sym_GT),
	8557:  uint16(anon_sym_PIPE),
	8558:  uint16(494),
	8559:  uint16(11),
	8560:  uint16(anon_sym_SEMI),
	8561:  uint16(anon_sym_RBRACK),
	8562:  uint16(anon_sym_COMMA),
	8563:  uint16(anon_sym_LPAREN),
	8564:  uint16(anon_sym_RPAREN),
	8565:  uint16(anon_sym_AMP_AMP),
	8566:  uint16(anon_sym_PIPE_PIPE),
	8567:  uint16(anon_sym_EQ_EQ),
	8568:  uint16(anon_sym_BANG_EQ),
	8569:  uint16(anon_sym_LT_EQ),
	8570:  uint16(anon_sym_GT_EQ),
	8571:  uint16(3),
	8572:  uint16(3),
	8573:  uint16(1),
	8574:  uint16(sym_comment),
	8575:  uint16(602),
	8576:  uint16(6),
	8577:  uint16(anon_sym_EQ),
	8578:  uint16(anon_sym_LT),
	8579:  uint16(anon_sym_GT),
	8580:  uint16(anon_sym_AMP),
	8581:  uint16(anon_sym_PIPE),
	8582:  uint16(anon_sym_SLASH),
	8583:  uint16(600),
	8584:  uint16(20),
	8585:  uint16(anon_sym_SEMI),
	8586:  uint16(anon_sym_RBRACK),
	8587:  uint16(anon_sym_COMMA),
	8588:  uint16(anon_sym_LPAREN),
	8589:  uint16(anon_sym_RPAREN),
	8590:  uint16(anon_sym_as),
	8591:  uint16(anon_sym_AMP_AMP),
	8592:  uint16(anon_sym_PIPE_PIPE),
	8593:  uint16(anon_sym_CARET),
	8594:  uint16(anon_sym_EQ_EQ),
	8595:  uint16(anon_sym_BANG_EQ),
	8596:  uint16(anon_sym_LT_EQ),
	8597:  uint16(anon_sym_GT_EQ),
	8598:  uint16(anon_sym_LT_LT),
	8599:  uint16(anon_sym_GT_GT),
	8600:  uint16(anon_sym_PLUS),
	8601:  uint16(anon_sym_DASH),
	8602:  uint16(anon_sym_STAR),
	8603:  uint16(anon_sym_PERCENT),
	8604:  uint16(anon_sym_DOT),
	8605:  uint16(3),
	8606:  uint16(3),
	8607:  uint16(1),
	8608:  uint16(sym_comment),
	8609:  uint16(606),
	8610:  uint16(6),
	8611:  uint16(anon_sym_EQ),
	8612:  uint16(anon_sym_LT),
	8613:  uint16(anon_sym_GT),
	8614:  uint16(anon_sym_AMP),
	8615:  uint16(anon_sym_PIPE),
	8616:  uint16(anon_sym_SLASH),
	8617:  uint16(604),
	8618:  uint16(20),
	8619:  uint16(anon_sym_SEMI),
	8620:  uint16(anon_sym_RBRACK),
	8621:  uint16(anon_sym_COMMA),
	8622:  uint16(anon_sym_LPAREN),
	8623:  uint16(anon_sym_RPAREN),
	8624:  uint16(anon_sym_as),
	8625:  uint16(anon_sym_AMP_AMP),
	8626:  uint16(anon_sym_PIPE_PIPE),
	8627:  uint16(anon_sym_CARET),
	8628:  uint16(anon_sym_EQ_EQ),
	8629:  uint16(anon_sym_BANG_EQ),
	8630:  uint16(anon_sym_LT_EQ),
	8631:  uint16(anon_sym_GT_EQ),
	8632:  uint16(anon_sym_LT_LT),
	8633:  uint16(anon_sym_GT_GT),
	8634:  uint16(anon_sym_PLUS),
	8635:  uint16(anon_sym_DASH),
	8636:  uint16(anon_sym_STAR),
	8637:  uint16(anon_sym_PERCENT),
	8638:  uint16(anon_sym_DOT),
	8639:  uint16(5),
	8640:  uint16(3),
	8641:  uint16(1),
	8642:  uint16(sym_comment),
	8643:  uint16(612),
	8644:  uint16(1),
	8645:  uint16(anon_sym_POUND),
	8646:  uint16(139),
	8647:  uint16(2),
	8648:  uint16(sym_attribute_item),
	8649:  uint16(aux_sym_tuple_expression_repeat1),
	8650:  uint16(610),
	8651:  uint16(7),
	8652:  uint16(anon_sym_COLON_COLON),
	8653:  uint16(anon_sym_LBRACE),
	8654:  uint16(anon_sym_LT),
	8655:  uint16(anon_sym_LPAREN),
	8656:  uint16(anon_sym_AMP),
	8657:  uint16(anon_sym_BANG),
	8658:  uint16(sym_float_literal),
	8659:  uint16(608),
	8660:  uint16(16),
	8661:  uint16(anon_sym_u8),
	8662:  uint16(anon_sym_u64),
	8663:  uint16(anon_sym_u128),
	8664:  uint16(anon_sym_bool),
	8665:  uint16(anon_sym_address),
	8666:  uint16(anon_sym_return),
	8667:  uint16(anon_sym_if),
	8668:  uint16(anon_sym_while),
	8669:  uint16(anon_sym_loop),
	8670:  uint16(anon_sym_const),
	8671:  uint16(anon_sym_break),
	8672:  uint16(anon_sym_continue),
	8673:  uint16(sym_integer_literal),
	8674:  uint16(anon_sym_true),
	8675:  uint16(anon_sym_false),
	8676:  uint16(sym_identifier),
	8677:  uint16(3),
	8678:  uint16(3),
	8679:  uint16(1),
	8680:  uint16(sym_comment),
	8681:  uint16(617),
	8682:  uint16(8),
	8683:  uint16(anon_sym_COLON_COLON),
	8684:  uint16(anon_sym_LBRACE),
	8685:  uint16(anon_sym_POUND),
	8686:  uint16(anon_sym_LT),
	8687:  uint16(anon_sym_LPAREN),
	8688:  uint16(anon_sym_AMP),
	8689:  uint16(anon_sym_BANG),
	8690:  uint16(sym_float_literal),
	8691:  uint16(615),
	8692:  uint16(16),
	8693:  uint16(anon_sym_u8),
	8694:  uint16(anon_sym_u64),
	8695:  uint16(anon_sym_u128),
	8696:  uint16(anon_sym_bool),
	8697:  uint16(anon_sym_address),
	8698:  uint16(anon_sym_return),
	8699:  uint16(anon_sym_if),
	8700:  uint16(anon_sym_while),
	8701:  uint16(anon_sym_loop),
	8702:  uint16(anon_sym_const),
	8703:  uint16(anon_sym_break),
	8704:  uint16(anon_sym_continue),
	8705:  uint16(sym_integer_literal),
	8706:  uint16(anon_sym_true),
	8707:  uint16(anon_sym_false),
	8708:  uint16(sym_identifier),
	8709:  uint16(18),
	8710:  uint16(3),
	8711:  uint16(1),
	8712:  uint16(sym_comment),
	8713:  uint16(500),
	8714:  uint16(1),
	8715:  uint16(anon_sym_as),
	8716:  uint16(502),
	8717:  uint16(1),
	8718:  uint16(anon_sym_AMP),
	8719:  uint16(504),
	8720:  uint16(1),
	8721:  uint16(anon_sym_PIPE),
	8722:  uint16(506),
	8723:  uint16(1),
	8724:  uint16(anon_sym_CARET),
	8725:  uint16(516),
	8726:  uint16(1),
	8727:  uint16(anon_sym_SLASH),
	8728:  uint16(518),
	8729:  uint16(1),
	8730:  uint16(anon_sym_DOT),
	8731:  uint16(532),
	8732:  uint16(1),
	8733:  uint16(anon_sym_AMP_AMP),
	8734:  uint16(534),
	8735:  uint16(1),
	8736:  uint16(anon_sym_PIPE_PIPE),
	8737:  uint16(619),
	8738:  uint16(1),
	8739:  uint16(anon_sym_EQ),
	8740:  uint16(621),
	8741:  uint16(1),
	8742:  uint16(anon_sym_COMMA),
	8743:  uint16(623),
	8744:  uint16(1),
	8745:  uint16(anon_sym_RPAREN),
	8746:  uint16(236),
	8747:  uint16(1),
	8748:  uint16(aux_sym_arguments_repeat1),
	8749:  uint16(498),
	8750:  uint16(2),
	8751:  uint16(anon_sym_LT),
	8752:  uint16(anon_sym_GT),
	8753:  uint16(510),
	8754:  uint16(2),
	8755:  uint16(anon_sym_LT_LT),
	8756:  uint16(anon_sym_GT_GT),
	8757:  uint16(512),
	8758:  uint16(2),
	8759:  uint16(anon_sym_PLUS),
	8760:  uint16(anon_sym_DASH),
	8761:  uint16(514),
	8762:  uint16(2),
	8763:  uint16(anon_sym_STAR),
	8764:  uint16(anon_sym_PERCENT),
	8765:  uint16(508),
	8766:  uint16(4),
	8767:  uint16(anon_sym_EQ_EQ),
	8768:  uint16(anon_sym_BANG_EQ),
	8769:  uint16(anon_sym_LT_EQ),
	8770:  uint16(anon_sym_GT_EQ),
	8771:  uint16(3),
	8772:  uint16(3),
	8773:  uint16(1),
	8774:  uint16(sym_comment),
	8775:  uint16(173),
	8776:  uint16(8),
	8777:  uint16(anon_sym_COLON_COLON),
	8778:  uint16(anon_sym_LBRACE),
	8779:  uint16(anon_sym_LT),
	8780:  uint16(anon_sym_LPAREN),
	8781:  uint16(anon_sym_RPAREN),
	8782:  uint16(anon_sym_AMP),
	8783:  uint16(anon_sym_BANG),
	8784:  uint16(sym_float_literal),
	8785:  uint16(625),
	8786:  uint16(16),
	8787:  uint16(anon_sym_u8),
	8788:  uint16(anon_sym_u64),
	8789:  uint16(anon_sym_u128),
	8790:  uint16(anon_sym_bool),
	8791:  uint16(anon_sym_address),
	8792:  uint16(anon_sym_return),
	8793:  uint16(anon_sym_if),
	8794:  uint16(anon_sym_while),
	8795:  uint16(anon_sym_loop),
	8796:  uint16(anon_sym_const),
	8797:  uint16(anon_sym_break),
	8798:  uint16(anon_sym_continue),
	8799:  uint16(sym_integer_literal),
	8800:  uint16(anon_sym_true),
	8801:  uint16(anon_sym_false),
	8802:  uint16(sym_identifier),
	8803:  uint16(17),
	8804:  uint16(3),
	8805:  uint16(1),
	8806:  uint16(sym_comment),
	8807:  uint16(156),
	8808:  uint16(1),
	8809:  uint16(anon_sym_RPAREN),
	8810:  uint16(500),
	8811:  uint16(1),
	8812:  uint16(anon_sym_as),
	8813:  uint16(502),
	8814:  uint16(1),
	8815:  uint16(anon_sym_AMP),
	8816:  uint16(504),
	8817:  uint16(1),
	8818:  uint16(anon_sym_PIPE),
	8819:  uint16(506),
	8820:  uint16(1),
	8821:  uint16(anon_sym_CARET),
	8822:  uint16(516),
	8823:  uint16(1),
	8824:  uint16(anon_sym_SLASH),
	8825:  uint16(518),
	8826:  uint16(1),
	8827:  uint16(anon_sym_DOT),
	8828:  uint16(532),
	8829:  uint16(1),
	8830:  uint16(anon_sym_AMP_AMP),
	8831:  uint16(534),
	8832:  uint16(1),
	8833:  uint16(anon_sym_PIPE_PIPE),
	8834:  uint16(619),
	8835:  uint16(1),
	8836:  uint16(anon_sym_EQ),
	8837:  uint16(627),
	8838:  uint16(1),
	8839:  uint16(anon_sym_COMMA),
	8840:  uint16(498),
	8841:  uint16(2),
	8842:  uint16(anon_sym_LT),
	8843:  uint16(anon_sym_GT),
	8844:  uint16(510),
	8845:  uint16(2),
	8846:  uint16(anon_sym_LT_LT),
	8847:  uint16(anon_sym_GT_GT),
	8848:  uint16(512),
	8849:  uint16(2),
	8850:  uint16(anon_sym_PLUS),
	8851:  uint16(anon_sym_DASH),
	8852:  uint16(514),
	8853:  uint16(2),
	8854:  uint16(anon_sym_STAR),
	8855:  uint16(anon_sym_PERCENT),
	8856:  uint16(508),
	8857:  uint16(4),
	8858:  uint16(anon_sym_EQ_EQ),
	8859:  uint16(anon_sym_BANG_EQ),
	8860:  uint16(anon_sym_LT_EQ),
	8861:  uint16(anon_sym_GT_EQ),
	8862:  uint16(17),
	8863:  uint16(3),
	8864:  uint16(1),
	8865:  uint16(sym_comment),
	8866:  uint16(500),
	8867:  uint16(1),
	8868:  uint16(anon_sym_as),
	8869:  uint16(502),
	8870:  uint16(1),
	8871:  uint16(anon_sym_AMP),
	8872:  uint16(504),
	8873:  uint16(1),
	8874:  uint16(anon_sym_PIPE),
	8875:  uint16(506),
	8876:  uint16(1),
	8877:  uint16(anon_sym_CARET),
	8878:  uint16(516),
	8879:  uint16(1),
	8880:  uint16(anon_sym_SLASH),
	8881:  uint16(518),
	8882:  uint16(1),
	8883:  uint16(anon_sym_DOT),
	8884:  uint16(532),
	8885:  uint16(1),
	8886:  uint16(anon_sym_AMP_AMP),
	8887:  uint16(534),
	8888:  uint16(1),
	8889:  uint16(anon_sym_PIPE_PIPE),
	8890:  uint16(619),
	8891:  uint16(1),
	8892:  uint16(anon_sym_EQ),
	8893:  uint16(629),
	8894:  uint16(1),
	8895:  uint16(anon_sym_COMMA),
	8896:  uint16(631),
	8897:  uint16(1),
	8898:  uint16(anon_sym_RPAREN),
	8899:  uint16(498),
	8900:  uint16(2),
	8901:  uint16(anon_sym_LT),
	8902:  uint16(anon_sym_GT),
	8903:  uint16(510),
	8904:  uint16(2),
	8905:  uint16(anon_sym_LT_LT),
	8906:  uint16(anon_sym_GT_GT),
	8907:  uint16(512),
	8908:  uint16(2),
	8909:  uint16(anon_sym_PLUS),
	8910:  uint16(anon_sym_DASH),
	8911:  uint16(514),
	8912:  uint16(2),
	8913:  uint16(anon_sym_STAR),
	8914:  uint16(anon_sym_PERCENT),
	8915:  uint16(508),
	8916:  uint16(4),
	8917:  uint16(anon_sym_EQ_EQ),
	8918:  uint16(anon_sym_BANG_EQ),
	8919:  uint16(anon_sym_LT_EQ),
	8920:  uint16(anon_sym_GT_EQ),
	8921:  uint16(17),
	8922:  uint16(3),
	8923:  uint16(1),
	8924:  uint16(sym_comment),
	8925:  uint16(500),
	8926:  uint16(1),
	8927:  uint16(anon_sym_as),
	8928:  uint16(502),
	8929:  uint16(1),
	8930:  uint16(anon_sym_AMP),
	8931:  uint16(504),
	8932:  uint16(1),
	8933:  uint16(anon_sym_PIPE),
	8934:  uint16(506),
	8935:  uint16(1),
	8936:  uint16(anon_sym_CARET),
	8937:  uint16(516),
	8938:  uint16(1),
	8939:  uint16(anon_sym_SLASH),
	8940:  uint16(518),
	8941:  uint16(1),
	8942:  uint16(anon_sym_DOT),
	8943:  uint16(532),
	8944:  uint16(1),
	8945:  uint16(anon_sym_AMP_AMP),
	8946:  uint16(534),
	8947:  uint16(1),
	8948:  uint16(anon_sym_PIPE_PIPE),
	8949:  uint16(619),
	8950:  uint16(1),
	8951:  uint16(anon_sym_EQ),
	8952:  uint16(627),
	8953:  uint16(1),
	8954:  uint16(anon_sym_COMMA),
	8955:  uint16(633),
	8956:  uint16(1),
	8957:  uint16(anon_sym_RPAREN),
	8958:  uint16(498),
	8959:  uint16(2),
	8960:  uint16(anon_sym_LT),
	8961:  uint16(anon_sym_GT),
	8962:  uint16(510),
	8963:  uint16(2),
	8964:  uint16(anon_sym_LT_LT),
	8965:  uint16(anon_sym_GT_GT),
	8966:  uint16(512),
	8967:  uint16(2),
	8968:  uint16(anon_sym_PLUS),
	8969:  uint16(anon_sym_DASH),
	8970:  uint16(514),
	8971:  uint16(2),
	8972:  uint16(anon_sym_STAR),
	8973:  uint16(anon_sym_PERCENT),
	8974:  uint16(508),
	8975:  uint16(4),
	8976:  uint16(anon_sym_EQ_EQ),
	8977:  uint16(anon_sym_BANG_EQ),
	8978:  uint16(anon_sym_LT_EQ),
	8979:  uint16(anon_sym_GT_EQ),
	8980:  uint16(16),
	8981:  uint16(3),
	8982:  uint16(1),
	8983:  uint16(sym_comment),
	8984:  uint16(500),
	8985:  uint16(1),
	8986:  uint16(anon_sym_as),
	8987:  uint16(502),
	8988:  uint16(1),
	8989:  uint16(anon_sym_AMP),
	8990:  uint16(504),
	8991:  uint16(1),
	8992:  uint16(anon_sym_PIPE),
	8993:  uint16(506),
	8994:  uint16(1),
	8995:  uint16(anon_sym_CARET),
	8996:  uint16(516),
	8997:  uint16(1),
	8998:  uint16(anon_sym_SLASH),
	8999:  uint16(518),
	9000:  uint16(1),
	9001:  uint16(anon_sym_DOT),
	9002:  uint16(532),
	9003:  uint16(1),
	9004:  uint16(anon_sym_AMP_AMP),
	9005:  uint16(534),
	9006:  uint16(1),
	9007:  uint16(anon_sym_PIPE_PIPE),
	9008:  uint16(619),
	9009:  uint16(1),
	9010:  uint16(anon_sym_EQ),
	9011:  uint16(498),
	9012:  uint16(2),
	9013:  uint16(anon_sym_LT),
	9014:  uint16(anon_sym_GT),
	9015:  uint16(510),
	9016:  uint16(2),
	9017:  uint16(anon_sym_LT_LT),
	9018:  uint16(anon_sym_GT_GT),
	9019:  uint16(512),
	9020:  uint16(2),
	9021:  uint16(anon_sym_PLUS),
	9022:  uint16(anon_sym_DASH),
	9023:  uint16(514),
	9024:  uint16(2),
	9025:  uint16(anon_sym_STAR),
	9026:  uint16(anon_sym_PERCENT),
	9027:  uint16(635),
	9028:  uint16(2),
	9029:  uint16(anon_sym_COMMA),
	9030:  uint16(anon_sym_RPAREN),
	9031:  uint16(508),
	9032:  uint16(4),
	9033:  uint16(anon_sym_EQ_EQ),
	9034:  uint16(anon_sym_BANG_EQ),
	9035:  uint16(anon_sym_LT_EQ),
	9036:  uint16(anon_sym_GT_EQ),
	9037:  uint16(17),
	9038:  uint16(3),
	9039:  uint16(1),
	9040:  uint16(sym_comment),
	9041:  uint16(152),
	9042:  uint16(1),
	9043:  uint16(anon_sym_RPAREN),
	9044:  uint16(500),
	9045:  uint16(1),
	9046:  uint16(anon_sym_as),
	9047:  uint16(502),
	9048:  uint16(1),
	9049:  uint16(anon_sym_AMP),
	9050:  uint16(504),
	9051:  uint16(1),
	9052:  uint16(anon_sym_PIPE),
	9053:  uint16(506),
	9054:  uint16(1),
	9055:  uint16(anon_sym_CARET),
	9056:  uint16(516),
	9057:  uint16(1),
	9058:  uint16(anon_sym_SLASH),
	9059:  uint16(518),
	9060:  uint16(1),
	9061:  uint16(anon_sym_DOT),
	9062:  uint16(532),
	9063:  uint16(1),
	9064:  uint16(anon_sym_AMP_AMP),
	9065:  uint16(534),
	9066:  uint16(1),
	9067:  uint16(anon_sym_PIPE_PIPE),
	9068:  uint16(619),
	9069:  uint16(1),
	9070:  uint16(anon_sym_EQ),
	9071:  uint16(627),
	9072:  uint16(1),
	9073:  uint16(anon_sym_COMMA),
	9074:  uint16(498),
	9075:  uint16(2),
	9076:  uint16(anon_sym_LT),
	9077:  uint16(anon_sym_GT),
	9078:  uint16(510),
	9079:  uint16(2),
	9080:  uint16(anon_sym_LT_LT),
	9081:  uint16(anon_sym_GT_GT),
	9082:  uint16(512),
	9083:  uint16(2),
	9084:  uint16(anon_sym_PLUS),
	9085:  uint16(anon_sym_DASH),
	9086:  uint16(514),
	9087:  uint16(2),
	9088:  uint16(anon_sym_STAR),
	9089:  uint16(anon_sym_PERCENT),
	9090:  uint16(508),
	9091:  uint16(4),
	9092:  uint16(anon_sym_EQ_EQ),
	9093:  uint16(anon_sym_BANG_EQ),
	9094:  uint16(anon_sym_LT_EQ),
	9095:  uint16(anon_sym_GT_EQ),
	9096:  uint16(16),
	9097:  uint16(3),
	9098:  uint16(1),
	9099:  uint16(sym_comment),
	9100:  uint16(500),
	9101:  uint16(1),
	9102:  uint16(anon_sym_as),
	9103:  uint16(502),
	9104:  uint16(1),
	9105:  uint16(anon_sym_AMP),
	9106:  uint16(504),
	9107:  uint16(1),
	9108:  uint16(anon_sym_PIPE),
	9109:  uint16(506),
	9110:  uint16(1),
	9111:  uint16(anon_sym_CARET),
	9112:  uint16(516),
	9113:  uint16(1),
	9114:  uint16(anon_sym_SLASH),
	9115:  uint16(518),
	9116:  uint16(1),
	9117:  uint16(anon_sym_DOT),
	9118:  uint16(532),
	9119:  uint16(1),
	9120:  uint16(anon_sym_AMP_AMP),
	9121:  uint16(534),
	9122:  uint16(1),
	9123:  uint16(anon_sym_PIPE_PIPE),
	9124:  uint16(619),
	9125:  uint16(1),
	9126:  uint16(anon_sym_EQ),
	9127:  uint16(637),
	9128:  uint16(1),
	9129:  uint16(anon_sym_RPAREN),
	9130:  uint16(498),
	9131:  uint16(2),
	9132:  uint16(anon_sym_LT),
	9133:  uint16(anon_sym_GT),
	9134:  uint16(510),
	9135:  uint16(2),
	9136:  uint16(anon_sym_LT_LT),
	9137:  uint16(anon_sym_GT_GT),
	9138:  uint16(512),
	9139:  uint16(2),
	9140:  uint16(anon_sym_PLUS),
	9141:  uint16(anon_sym_DASH),
	9142:  uint16(514),
	9143:  uint16(2),
	9144:  uint16(anon_sym_STAR),
	9145:  uint16(anon_sym_PERCENT),
	9146:  uint16(508),
	9147:  uint16(4),
	9148:  uint16(anon_sym_EQ_EQ),
	9149:  uint16(anon_sym_BANG_EQ),
	9150:  uint16(anon_sym_LT_EQ),
	9151:  uint16(anon_sym_GT_EQ),
	9152:  uint16(16),
	9153:  uint16(3),
	9154:  uint16(1),
	9155:  uint16(sym_comment),
	9156:  uint16(500),
	9157:  uint16(1),
	9158:  uint16(anon_sym_as),
	9159:  uint16(502),
	9160:  uint16(1),
	9161:  uint16(anon_sym_AMP),
	9162:  uint16(504),
	9163:  uint16(1),
	9164:  uint16(anon_sym_PIPE),
	9165:  uint16(506),
	9166:  uint16(1),
	9167:  uint16(anon_sym_CARET),
	9168:  uint16(516),
	9169:  uint16(1),
	9170:  uint16(anon_sym_SLASH),
	9171:  uint16(518),
	9172:  uint16(1),
	9173:  uint16(anon_sym_DOT),
	9174:  uint16(532),
	9175:  uint16(1),
	9176:  uint16(anon_sym_AMP_AMP),
	9177:  uint16(534),
	9178:  uint16(1),
	9179:  uint16(anon_sym_PIPE_PIPE),
	9180:  uint16(619),
	9181:  uint16(1),
	9182:  uint16(anon_sym_EQ),
	9183:  uint16(639),
	9184:  uint16(1),
	9185:  uint16(anon_sym_COMMA),
	9186:  uint16(498),
	9187:  uint16(2),
	9188:  uint16(anon_sym_LT),
	9189:  uint16(anon_sym_GT),
	9190:  uint16(510),
	9191:  uint16(2),
	9192:  uint16(anon_sym_LT_LT),
	9193:  uint16(anon_sym_GT_GT),
	9194:  uint16(512),
	9195:  uint16(2),
	9196:  uint16(anon_sym_PLUS),
	9197:  uint16(anon_sym_DASH),
	9198:  uint16(514),
	9199:  uint16(2),
	9200:  uint16(anon_sym_STAR),
	9201:  uint16(anon_sym_PERCENT),
	9202:  uint16(508),
	9203:  uint16(4),
	9204:  uint16(anon_sym_EQ_EQ),
	9205:  uint16(anon_sym_BANG_EQ),
	9206:  uint16(anon_sym_LT_EQ),
	9207:  uint16(anon_sym_GT_EQ),
	9208:  uint16(16),
	9209:  uint16(3),
	9210:  uint16(1),
	9211:  uint16(sym_comment),
	9212:  uint16(500),
	9213:  uint16(1),
	9214:  uint16(anon_sym_as),
	9215:  uint16(502),
	9216:  uint16(1),
	9217:  uint16(anon_sym_AMP),
	9218:  uint16(504),
	9219:  uint16(1),
	9220:  uint16(anon_sym_PIPE),
	9221:  uint16(506),
	9222:  uint16(1),
	9223:  uint16(anon_sym_CARET),
	9224:  uint16(516),
	9225:  uint16(1),
	9226:  uint16(anon_sym_SLASH),
	9227:  uint16(518),
	9228:  uint16(1),
	9229:  uint16(anon_sym_DOT),
	9230:  uint16(532),
	9231:  uint16(1),
	9232:  uint16(anon_sym_AMP_AMP),
	9233:  uint16(534),
	9234:  uint16(1),
	9235:  uint16(anon_sym_PIPE_PIPE),
	9236:  uint16(619),
	9237:  uint16(1),
	9238:  uint16(anon_sym_EQ),
	9239:  uint16(641),
	9240:  uint16(1),
	9241:  uint16(anon_sym_RBRACK),
	9242:  uint16(498),
	9243:  uint16(2),
	9244:  uint16(anon_sym_LT),
	9245:  uint16(anon_sym_GT),
	9246:  uint16(510),
	9247:  uint16(2),
	9248:  uint16(anon_sym_LT_LT),
	9249:  uint16(anon_sym_GT_GT),
	9250:  uint16(512),
	9251:  uint16(2),
	9252:  uint16(anon_sym_PLUS),
	9253:  uint16(anon_sym_DASH),
	9254:  uint16(514),
	9255:  uint16(2),
	9256:  uint16(anon_sym_STAR),
	9257:  uint16(anon_sym_PERCENT),
	9258:  uint16(508),
	9259:  uint16(4),
	9260:  uint16(anon_sym_EQ_EQ),
	9261:  uint16(anon_sym_BANG_EQ),
	9262:  uint16(anon_sym_LT_EQ),
	9263:  uint16(anon_sym_GT_EQ),
	9264:  uint16(16),
	9265:  uint16(3),
	9266:  uint16(1),
	9267:  uint16(sym_comment),
	9268:  uint16(500),
	9269:  uint16(1),
	9270:  uint16(anon_sym_as),
	9271:  uint16(502),
	9272:  uint16(1),
	9273:  uint16(anon_sym_AMP),
	9274:  uint16(504),
	9275:  uint16(1),
	9276:  uint16(anon_sym_PIPE),
	9277:  uint16(506),
	9278:  uint16(1),
	9279:  uint16(anon_sym_CARET),
	9280:  uint16(516),
	9281:  uint16(1),
	9282:  uint16(anon_sym_SLASH),
	9283:  uint16(518),
	9284:  uint16(1),
	9285:  uint16(anon_sym_DOT),
	9286:  uint16(532),
	9287:  uint16(1),
	9288:  uint16(anon_sym_AMP_AMP),
	9289:  uint16(534),
	9290:  uint16(1),
	9291:  uint16(anon_sym_PIPE_PIPE),
	9292:  uint16(619),
	9293:  uint16(1),
	9294:  uint16(anon_sym_EQ),
	9295:  uint16(643),
	9296:  uint16(1),
	9297:  uint16(anon_sym_RPAREN),
	9298:  uint16(498),
	9299:  uint16(2),
	9300:  uint16(anon_sym_LT),
	9301:  uint16(anon_sym_GT),
	9302:  uint16(510),
	9303:  uint16(2),
	9304:  uint16(anon_sym_LT_LT),
	9305:  uint16(anon_sym_GT_GT),
	9306:  uint16(512),
	9307:  uint16(2),
	9308:  uint16(anon_sym_PLUS),
	9309:  uint16(anon_sym_DASH),
	9310:  uint16(514),
	9311:  uint16(2),
	9312:  uint16(anon_sym_STAR),
	9313:  uint16(anon_sym_PERCENT),
	9314:  uint16(508),
	9315:  uint16(4),
	9316:  uint16(anon_sym_EQ_EQ),
	9317:  uint16(anon_sym_BANG_EQ),
	9318:  uint16(anon_sym_LT_EQ),
	9319:  uint16(anon_sym_GT_EQ),
	9320:  uint16(16),
	9321:  uint16(3),
	9322:  uint16(1),
	9323:  uint16(sym_comment),
	9324:  uint16(500),
	9325:  uint16(1),
	9326:  uint16(anon_sym_as),
	9327:  uint16(502),
	9328:  uint16(1),
	9329:  uint16(anon_sym_AMP),
	9330:  uint16(504),
	9331:  uint16(1),
	9332:  uint16(anon_sym_PIPE),
	9333:  uint16(506),
	9334:  uint16(1),
	9335:  uint16(anon_sym_CARET),
	9336:  uint16(516),
	9337:  uint16(1),
	9338:  uint16(anon_sym_SLASH),
	9339:  uint16(518),
	9340:  uint16(1),
	9341:  uint16(anon_sym_DOT),
	9342:  uint16(532),
	9343:  uint16(1),
	9344:  uint16(anon_sym_AMP_AMP),
	9345:  uint16(534),
	9346:  uint16(1),
	9347:  uint16(anon_sym_PIPE_PIPE),
	9348:  uint16(619),
	9349:  uint16(1),
	9350:  uint16(anon_sym_EQ),
	9351:  uint16(627),
	9352:  uint16(1),
	9353:  uint16(anon_sym_COMMA),
	9354:  uint16(498),
	9355:  uint16(2),
	9356:  uint16(anon_sym_LT),
	9357:  uint16(anon_sym_GT),
	9358:  uint16(510),
	9359:  uint16(2),
	9360:  uint16(anon_sym_LT_LT),
	9361:  uint16(anon_sym_GT_GT),
	9362:  uint16(512),
	9363:  uint16(2),
	9364:  uint16(anon_sym_PLUS),
	9365:  uint16(anon_sym_DASH),
	9366:  uint16(514),
	9367:  uint16(2),
	9368:  uint16(anon_sym_STAR),
	9369:  uint16(anon_sym_PERCENT),
	9370:  uint16(508),
	9371:  uint16(4),
	9372:  uint16(anon_sym_EQ_EQ),
	9373:  uint16(anon_sym_BANG_EQ),
	9374:  uint16(anon_sym_LT_EQ),
	9375:  uint16(anon_sym_GT_EQ),
	9376:  uint16(16),
	9377:  uint16(3),
	9378:  uint16(1),
	9379:  uint16(sym_comment),
	9380:  uint16(500),
	9381:  uint16(1),
	9382:  uint16(anon_sym_as),
	9383:  uint16(502),
	9384:  uint16(1),
	9385:  uint16(anon_sym_AMP),
	9386:  uint16(504),
	9387:  uint16(1),
	9388:  uint16(anon_sym_PIPE),
	9389:  uint16(506),
	9390:  uint16(1),
	9391:  uint16(anon_sym_CARET),
	9392:  uint16(516),
	9393:  uint16(1),
	9394:  uint16(anon_sym_SLASH),
	9395:  uint16(518),
	9396:  uint16(1),
	9397:  uint16(anon_sym_DOT),
	9398:  uint16(532),
	9399:  uint16(1),
	9400:  uint16(anon_sym_AMP_AMP),
	9401:  uint16(534),
	9402:  uint16(1),
	9403:  uint16(anon_sym_PIPE_PIPE),
	9404:  uint16(619),
	9405:  uint16(1),
	9406:  uint16(anon_sym_EQ),
	9407:  uint16(645),
	9408:  uint16(1),
	9409:  uint16(anon_sym_SEMI),
	9410:  uint16(498),
	9411:  uint16(2),
	9412:  uint16(anon_sym_LT),
	9413:  uint16(anon_sym_GT),
	9414:  uint16(510),
	9415:  uint16(2),
	9416:  uint16(anon_sym_LT_LT),
	9417:  uint16(anon_sym_GT_GT),
	9418:  uint16(512),
	9419:  uint16(2),
	9420:  uint16(anon_sym_PLUS),
	9421:  uint16(anon_sym_DASH),
	9422:  uint16(514),
	9423:  uint16(2),
	9424:  uint16(anon_sym_STAR),
	9425:  uint16(anon_sym_PERCENT),
	9426:  uint16(508),
	9427:  uint16(4),
	9428:  uint16(anon_sym_EQ_EQ),
	9429:  uint16(anon_sym_BANG_EQ),
	9430:  uint16(anon_sym_LT_EQ),
	9431:  uint16(anon_sym_GT_EQ),
	9432:  uint16(11),
	9433:  uint16(3),
	9434:  uint16(1),
	9435:  uint16(sym_comment),
	9436:  uint16(647),
	9437:  uint16(1),
	9438:  uint16(sym_identifier),
	9439:  uint16(649),
	9440:  uint16(1),
	9441:  uint16(anon_sym_LBRACE),
	9442:  uint16(651),
	9443:  uint16(1),
	9444:  uint16(anon_sym_GT),
	9445:  uint16(653),
	9446:  uint16(1),
	9447:  uint16(anon_sym_LPAREN),
	9448:  uint16(657),
	9449:  uint16(1),
	9450:  uint16(anon_sym_AMP),
	9451:  uint16(659),
	9452:  uint16(1),
	9453:  uint16(sym_integer_literal),
	9454:  uint16(661),
	9455:  uint16(1),
	9456:  uint16(sym_float_literal),
	9457:  uint16(663),
	9458:  uint16(2),
	9459:  uint16(anon_sym_true),
	9460:  uint16(anon_sym_false),
	9461:  uint16(655),
	9462:  uint16(5),
	9463:  uint16(anon_sym_u8),
	9464:  uint16(anon_sym_u64),
	9465:  uint16(anon_sym_u128),
	9466:  uint16(anon_sym_bool),
	9467:  uint16(anon_sym_address),
	9468:  uint16(263),
	9469:  uint16(7),
	9470:  uint16(sym__type),
	9471:  uint16(sym_unit_type),
	9472:  uint16(sym_reference_type),
	9473:  uint16(sym_block),
	9474:  uint16(sym__literal),
	9475:  uint16(sym_boolean_literal),
	9476:  uint16(sym__type_identifier),
	9477:  uint16(11),
	9478:  uint16(3),
	9479:  uint16(1),
	9480:  uint16(sym_comment),
	9481:  uint16(647),
	9482:  uint16(1),
	9483:  uint16(sym_identifier),
	9484:  uint16(649),
	9485:  uint16(1),
	9486:  uint16(anon_sym_LBRACE),
	9487:  uint16(653),
	9488:  uint16(1),
	9489:  uint16(anon_sym_LPAREN),
	9490:  uint16(657),
	9491:  uint16(1),
	9492:  uint16(anon_sym_AMP),
	9493:  uint16(659),
	9494:  uint16(1),
	9495:  uint16(sym_integer_literal),
	9496:  uint16(661),
	9497:  uint16(1),
	9498:  uint16(sym_float_literal),
	9499:  uint16(665),
	9500:  uint16(1),
	9501:  uint16(anon_sym_GT),
	9502:  uint16(663),
	9503:  uint16(2),
	9504:  uint16(anon_sym_true),
	9505:  uint16(anon_sym_false),
	9506:  uint16(655),
	9507:  uint16(5),
	9508:  uint16(anon_sym_u8),
	9509:  uint16(anon_sym_u64),
	9510:  uint16(anon_sym_u128),
	9511:  uint16(anon_sym_bool),
	9512:  uint16(anon_sym_address),
	9513:  uint16(263),
	9514:  uint16(7),
	9515:  uint16(sym__type),
	9516:  uint16(sym_unit_type),
	9517:  uint16(sym_reference_type),
	9518:  uint16(sym_block),
	9519:  uint16(sym__literal),
	9520:  uint16(sym_boolean_literal),
	9521:  uint16(sym__type_identifier),
	9522:  uint16(10),
	9523:  uint16(3),
	9524:  uint16(1),
	9525:  uint16(sym_comment),
	9526:  uint16(647),
	9527:  uint16(1),
	9528:  uint16(sym_identifier),
	9529:  uint16(649),
	9530:  uint16(1),
	9531:  uint16(anon_sym_LBRACE),
	9532:  uint16(653),
	9533:  uint16(1),
	9534:  uint16(anon_sym_LPAREN),
	9535:  uint16(657),
	9536:  uint16(1),
	9537:  uint16(anon_sym_AMP),
	9538:  uint16(667),
	9539:  uint16(1),
	9540:  uint16(sym_integer_literal),
	9541:  uint16(669),
	9542:  uint16(1),
	9543:  uint16(sym_float_literal),
	9544:  uint16(663),
	9545:  uint16(2),
	9546:  uint16(anon_sym_true),
	9547:  uint16(anon_sym_false),
	9548:  uint16(655),
	9549:  uint16(5),
	9550:  uint16(anon_sym_u8),
	9551:  uint16(anon_sym_u64),
	9552:  uint16(anon_sym_u128),
	9553:  uint16(anon_sym_bool),
	9554:  uint16(anon_sym_address),
	9555:  uint16(228),
	9556:  uint16(7),
	9557:  uint16(sym__type),
	9558:  uint16(sym_unit_type),
	9559:  uint16(sym_reference_type),
	9560:  uint16(sym_block),
	9561:  uint16(sym__literal),
	9562:  uint16(sym_boolean_literal),
	9563:  uint16(sym__type_identifier),
	9564:  uint16(10),
	9565:  uint16(3),
	9566:  uint16(1),
	9567:  uint16(sym_comment),
	9568:  uint16(647),
	9569:  uint16(1),
	9570:  uint16(sym_identifier),
	9571:  uint16(649),
	9572:  uint16(1),
	9573:  uint16(anon_sym_LBRACE),
	9574:  uint16(653),
	9575:  uint16(1),
	9576:  uint16(anon_sym_LPAREN),
	9577:  uint16(657),
	9578:  uint16(1),
	9579:  uint16(anon_sym_AMP),
	9580:  uint16(659),
	9581:  uint16(1),
	9582:  uint16(sym_integer_literal),
	9583:  uint16(661),
	9584:  uint16(1),
	9585:  uint16(sym_float_literal),
	9586:  uint16(663),
	9587:  uint16(2),
	9588:  uint16(anon_sym_true),
	9589:  uint16(anon_sym_false),
	9590:  uint16(655),
	9591:  uint16(5),
	9592:  uint16(anon_sym_u8),
	9593:  uint16(anon_sym_u64),
	9594:  uint16(anon_sym_u128),
	9595:  uint16(anon_sym_bool),
	9596:  uint16(anon_sym_address),
	9597:  uint16(263),
	9598:  uint16(7),
	9599:  uint16(sym__type),
	9600:  uint16(sym_unit_type),
	9601:  uint16(sym_reference_type),
	9602:  uint16(sym_block),
	9603:  uint16(sym__literal),
	9604:  uint16(sym_boolean_literal),
	9605:  uint16(sym__type_identifier),
	9606:  uint16(9),
	9607:  uint16(3),
	9608:  uint16(1),
	9609:  uint16(sym_comment),
	9610:  uint16(653),
	9611:  uint16(1),
	9612:  uint16(anon_sym_LPAREN),
	9613:  uint16(657),
	9614:  uint16(1),
	9615:  uint16(anon_sym_AMP),
	9616:  uint16(671),
	9617:  uint16(1),
	9618:  uint16(sym_identifier),
	9619:  uint16(673),
	9620:  uint16(1),
	9621:  uint16(anon_sym_COMMA),
	9622:  uint16(675),
	9623:  uint16(1),
	9624:  uint16(anon_sym__),
	9625:  uint16(677),
	9626:  uint16(1),
	9627:  uint16(anon_sym_RPAREN),
	9628:  uint16(655),
	9629:  uint16(5),
	9630:  uint16(anon_sym_u8),
	9631:  uint16(anon_sym_u64),
	9632:  uint16(anon_sym_u128),
	9633:  uint16(anon_sym_bool),
	9634:  uint16(anon_sym_address),
	9635:  uint16(238),
	9636:  uint16(5),
	9637:  uint16(sym_parameter),
	9638:  uint16(sym__type),
	9639:  uint16(sym_unit_type),
	9640:  uint16(sym_reference_type),
	9641:  uint16(sym__type_identifier),
	9642:  uint16(8),
	9643:  uint16(3),
	9644:  uint16(1),
	9645:  uint16(sym_comment),
	9646:  uint16(653),
	9647:  uint16(1),
	9648:  uint16(anon_sym_LPAREN),
	9649:  uint16(657),
	9650:  uint16(1),
	9651:  uint16(anon_sym_AMP),
	9652:  uint16(671),
	9653:  uint16(1),
	9654:  uint16(sym_identifier),
	9655:  uint16(679),
	9656:  uint16(1),
	9657:  uint16(anon_sym__),
	9658:  uint16(681),
	9659:  uint16(1),
	9660:  uint16(anon_sym_RPAREN),
	9661:  uint16(655),
	9662:  uint16(5),
	9663:  uint16(anon_sym_u8),
	9664:  uint16(anon_sym_u64),
	9665:  uint16(anon_sym_u128),
	9666:  uint16(anon_sym_bool),
	9667:  uint16(anon_sym_address),
	9668:  uint16(270),
	9669:  uint16(5),
	9670:  uint16(sym_parameter),
	9671:  uint16(sym__type),
	9672:  uint16(sym_unit_type),
	9673:  uint16(sym_reference_type),
	9674:  uint16(sym__type_identifier),
	9675:  uint16(8),
	9676:  uint16(3),
	9677:  uint16(1),
	9678:  uint16(sym_comment),
	9679:  uint16(653),
	9680:  uint16(1),
	9681:  uint16(anon_sym_LPAREN),
	9682:  uint16(657),
	9683:  uint16(1),
	9684:  uint16(anon_sym_AMP),
	9685:  uint16(671),
	9686:  uint16(1),
	9687:  uint16(sym_identifier),
	9688:  uint16(679),
	9689:  uint16(1),
	9690:  uint16(anon_sym__),
	9691:  uint16(683),
	9692:  uint16(1),
	9693:  uint16(anon_sym_RPAREN),
	9694:  uint16(655),
	9695:  uint16(5),
	9696:  uint16(anon_sym_u8),
	9697:  uint16(anon_sym_u64),
	9698:  uint16(anon_sym_u128),
	9699:  uint16(anon_sym_bool),
	9700:  uint16(anon_sym_address),
	9701:  uint16(270),
	9702:  uint16(5),
	9703:  uint16(sym_parameter),
	9704:  uint16(sym__type),
	9705:  uint16(sym_unit_type),
	9706:  uint16(sym_reference_type),
	9707:  uint16(sym__type_identifier),
	9708:  uint16(7),
	9709:  uint16(3),
	9710:  uint16(1),
	9711:  uint16(sym_comment),
	9712:  uint16(653),
	9713:  uint16(1),
	9714:  uint16(anon_sym_LPAREN),
	9715:  uint16(657),
	9716:  uint16(1),
	9717:  uint16(anon_sym_AMP),
	9718:  uint16(671),
	9719:  uint16(1),
	9720:  uint16(sym_identifier),
	9721:  uint16(679),
	9722:  uint16(1),
	9723:  uint16(anon_sym__),
	9724:  uint16(655),
	9725:  uint16(5),
	9726:  uint16(anon_sym_u8),
	9727:  uint16(anon_sym_u64),
	9728:  uint16(anon_sym_u128),
	9729:  uint16(anon_sym_bool),
	9730:  uint16(anon_sym_address),
	9731:  uint16(270),
	9732:  uint16(5),
	9733:  uint16(sym_parameter),
	9734:  uint16(sym__type),
	9735:  uint16(sym_unit_type),
	9736:  uint16(sym_reference_type),
	9737:  uint16(sym__type_identifier),
	9738:  uint16(7),
	9739:  uint16(3),
	9740:  uint16(1),
	9741:  uint16(sym_comment),
	9742:  uint16(647),
	9743:  uint16(1),
	9744:  uint16(sym_identifier),
	9745:  uint16(653),
	9746:  uint16(1),
	9747:  uint16(anon_sym_LPAREN),
	9748:  uint16(657),
	9749:  uint16(1),
	9750:  uint16(anon_sym_AMP),
	9751:  uint16(331),
	9752:  uint16(1),
	9753:  uint16(sym_qualified_type),
	9754:  uint16(262),
	9755:  uint16(4),
	9756:  uint16(sym__type),
	9757:  uint16(sym_unit_type),
	9758:  uint16(sym_reference_type),
	9759:  uint16(sym__type_identifier),
	9760:  uint16(655),
	9761:  uint16(5),
	9762:  uint16(anon_sym_u8),
	9763:  uint16(anon_sym_u64),
	9764:  uint16(anon_sym_u128),
	9765:  uint16(anon_sym_bool),
	9766:  uint16(anon_sym_address),
	9767:  uint16(6),
	9768:  uint16(3),
	9769:  uint16(1),
	9770:  uint16(sym_comment),
	9771:  uint16(647),
	9772:  uint16(1),
	9773:  uint16(sym_identifier),
	9774:  uint16(653),
	9775:  uint16(1),
	9776:  uint16(anon_sym_LPAREN),
	9777:  uint16(657),
	9778:  uint16(1),
	9779:  uint16(anon_sym_AMP),
	9780:  uint16(305),
	9781:  uint16(4),
	9782:  uint16(sym__type),
	9783:  uint16(sym_unit_type),
	9784:  uint16(sym_reference_type),
	9785:  uint16(sym__type_identifier),
	9786:  uint16(655),
	9787:  uint16(5),
	9788:  uint16(anon_sym_u8),
	9789:  uint16(anon_sym_u64),
	9790:  uint16(anon_sym_u128),
	9791:  uint16(anon_sym_bool),
	9792:  uint16(anon_sym_address),
	9793:  uint16(7),
	9794:  uint16(3),
	9795:  uint16(1),
	9796:  uint16(sym_comment),
	9797:  uint16(9),
	9798:  uint16(1),
	9799:  uint16(anon_sym_COLON_COLON),
	9800:  uint16(120),
	9801:  uint16(1),
	9802:  uint16(anon_sym_LT),
	9803:  uint16(329),
	9804:  uint16(1),
	9805:  uint16(sym_bracketed_type),
	9806:  uint16(330),
	9807:  uint16(1),
	9808:  uint16(sym__use_clause),
	9809:  uint16(294),
	9810:  uint16(2),
	9811:  uint16(sym_scoped_identifier),
	9812:  uint16(sym__path),
	9813:  uint16(685),
	9814:  uint16(6),
	9815:  uint16(anon_sym_u8),
	9816:  uint16(anon_sym_u64),
	9817:  uint16(anon_sym_u128),
	9818:  uint16(anon_sym_bool),
	9819:  uint16(anon_sym_address),
	9820:  uint16(sym_identifier),
	9821:  uint16(7),
	9822:  uint16(3),
	9823:  uint16(1),
	9824:  uint16(sym_comment),
	9825:  uint16(9),
	9826:  uint16(1),
	9827:  uint16(anon_sym_COLON_COLON),
	9828:  uint16(120),
	9829:  uint16(1),
	9830:  uint16(anon_sym_LT),
	9831:  uint16(309),
	9832:  uint16(1),
	9833:  uint16(sym_attribute),
	9834:  uint16(329),
	9835:  uint16(1),
	9836:  uint16(sym_bracketed_type),
	9837:  uint16(185),
	9838:  uint16(2),
	9839:  uint16(sym_scoped_identifier),
	9840:  uint16(sym__path),
	9841:  uint16(687),
	9842:  uint16(6),
	9843:  uint16(anon_sym_u8),
	9844:  uint16(anon_sym_u64),
	9845:  uint16(anon_sym_u128),
	9846:  uint16(anon_sym_bool),
	9847:  uint16(anon_sym_address),
	9848:  uint16(sym_identifier),
	9849:  uint16(6),
	9850:  uint16(3),
	9851:  uint16(1),
	9852:  uint16(sym_comment),
	9853:  uint16(647),
	9854:  uint16(1),
	9855:  uint16(sym_identifier),
	9856:  uint16(653),
	9857:  uint16(1),
	9858:  uint16(anon_sym_LPAREN),
	9859:  uint16(657),
	9860:  uint16(1),
	9861:  uint16(anon_sym_AMP),
	9862:  uint16(341),
	9863:  uint16(4),
	9864:  uint16(sym__type),
	9865:  uint16(sym_unit_type),
	9866:  uint16(sym_reference_type),
	9867:  uint16(sym__type_identifier),
	9868:  uint16(655),
	9869:  uint16(5),
	9870:  uint16(anon_sym_u8),
	9871:  uint16(anon_sym_u64),
	9872:  uint16(anon_sym_u128),
	9873:  uint16(anon_sym_bool),
	9874:  uint16(anon_sym_address),
	9875:  uint16(6),
	9876:  uint16(3),
	9877:  uint16(1),
	9878:  uint16(sym_comment),
	9879:  uint16(647),
	9880:  uint16(1),
	9881:  uint16(sym_identifier),
	9882:  uint16(653),
	9883:  uint16(1),
	9884:  uint16(anon_sym_LPAREN),
	9885:  uint16(657),
	9886:  uint16(1),
	9887:  uint16(anon_sym_AMP),
	9888:  uint16(286),
	9889:  uint16(4),
	9890:  uint16(sym__type),
	9891:  uint16(sym_unit_type),
	9892:  uint16(sym_reference_type),
	9893:  uint16(sym__type_identifier),
	9894:  uint16(655),
	9895:  uint16(5),
	9896:  uint16(anon_sym_u8),
	9897:  uint16(anon_sym_u64),
	9898:  uint16(anon_sym_u128),
	9899:  uint16(anon_sym_bool),
	9900:  uint16(anon_sym_address),
	9901:  uint16(6),
	9902:  uint16(3),
	9903:  uint16(1),
	9904:  uint16(sym_comment),
	9905:  uint16(647),
	9906:  uint16(1),
	9907:  uint16(sym_identifier),
	9908:  uint16(653),
	9909:  uint16(1),
	9910:  uint16(anon_sym_LPAREN),
	9911:  uint16(657),
	9912:  uint16(1),
	9913:  uint16(anon_sym_AMP),
	9914:  uint16(287),
	9915:  uint16(4),
	9916:  uint16(sym__type),
	9917:  uint16(sym_unit_type),
	9918:  uint16(sym_reference_type),
	9919:  uint16(sym__type_identifier),
	9920:  uint16(655),
	9921:  uint16(5),
	9922:  uint16(anon_sym_u8),
	9923:  uint16(anon_sym_u64),
	9924:  uint16(anon_sym_u128),
	9925:  uint16(anon_sym_bool),
	9926:  uint16(anon_sym_address),
	9927:  uint16(6),
	9928:  uint16(3),
	9929:  uint16(1),
	9930:  uint16(sym_comment),
	9931:  uint16(689),
	9932:  uint16(1),
	9933:  uint16(sym_identifier),
	9934:  uint16(691),
	9935:  uint16(1),
	9936:  uint16(anon_sym_LPAREN),
	9937:  uint16(695),
	9938:  uint16(1),
	9939:  uint16(anon_sym_AMP),
	9940:  uint16(101),
	9941:  uint16(4),
	9942:  uint16(sym__type),
	9943:  uint16(sym_unit_type),
	9944:  uint16(sym_reference_type),
	9945:  uint16(sym__type_identifier),
	9946:  uint16(693),
	9947:  uint16(5),
	9948:  uint16(anon_sym_u8),
	9949:  uint16(anon_sym_u64),
	9950:  uint16(anon_sym_u128),
	9951:  uint16(anon_sym_bool),
	9952:  uint16(anon_sym_address),
	9953:  uint16(6),
	9954:  uint16(3),
	9955:  uint16(1),
	9956:  uint16(sym_comment),
	9957:  uint16(689),
	9958:  uint16(1),
	9959:  uint16(sym_identifier),
	9960:  uint16(691),
	9961:  uint16(1),
	9962:  uint16(anon_sym_LPAREN),
	9963:  uint16(695),
	9964:  uint16(1),
	9965:  uint16(anon_sym_AMP),
	9966:  uint16(110),
	9967:  uint16(4),
	9968:  uint16(sym__type),
	9969:  uint16(sym_unit_type),
	9970:  uint16(sym_reference_type),
	9971:  uint16(sym__type_identifier),
	9972:  uint16(693),
	9973:  uint16(5),
	9974:  uint16(anon_sym_u8),
	9975:  uint16(anon_sym_u64),
	9976:  uint16(anon_sym_u128),
	9977:  uint16(anon_sym_bool),
	9978:  uint16(anon_sym_address),
	9979:  uint16(6),
	9980:  uint16(3),
	9981:  uint16(1),
	9982:  uint16(sym_comment),
	9983:  uint16(647),
	9984:  uint16(1),
	9985:  uint16(sym_identifier),
	9986:  uint16(653),
	9987:  uint16(1),
	9988:  uint16(anon_sym_LPAREN),
	9989:  uint16(657),
	9990:  uint16(1),
	9991:  uint16(anon_sym_AMP),
	9992:  uint16(266),
	9993:  uint16(4),
	9994:  uint16(sym__type),
	9995:  uint16(sym_unit_type),
	9996:  uint16(sym_reference_type),
	9997:  uint16(sym__type_identifier),
	9998:  uint16(655),
	9999:  uint16(5),
	10000: uint16(anon_sym_u8),
	10001: uint16(anon_sym_u64),
	10002: uint16(anon_sym_u128),
	10003: uint16(anon_sym_bool),
	10004: uint16(anon_sym_address),
	10005: uint16(6),
	10006: uint16(3),
	10007: uint16(1),
	10008: uint16(sym_comment),
	10009: uint16(647),
	10010: uint16(1),
	10011: uint16(sym_identifier),
	10012: uint16(653),
	10013: uint16(1),
	10014: uint16(anon_sym_LPAREN),
	10015: uint16(657),
	10016: uint16(1),
	10017: uint16(anon_sym_AMP),
	10018: uint16(284),
	10019: uint16(4),
	10020: uint16(sym__type),
	10021: uint16(sym_unit_type),
	10022: uint16(sym_reference_type),
	10023: uint16(sym__type_identifier),
	10024: uint16(655),
	10025: uint16(5),
	10026: uint16(anon_sym_u8),
	10027: uint16(anon_sym_u64),
	10028: uint16(anon_sym_u128),
	10029: uint16(anon_sym_bool),
	10030: uint16(anon_sym_address),
	10031: uint16(6),
	10032: uint16(3),
	10033: uint16(1),
	10034: uint16(sym_comment),
	10035: uint16(647),
	10036: uint16(1),
	10037: uint16(sym_identifier),
	10038: uint16(653),
	10039: uint16(1),
	10040: uint16(anon_sym_LPAREN),
	10041: uint16(657),
	10042: uint16(1),
	10043: uint16(anon_sym_AMP),
	10044: uint16(275),
	10045: uint16(4),
	10046: uint16(sym__type),
	10047: uint16(sym_unit_type),
	10048: uint16(sym_reference_type),
	10049: uint16(sym__type_identifier),
	10050: uint16(655),
	10051: uint16(5),
	10052: uint16(anon_sym_u8),
	10053: uint16(anon_sym_u64),
	10054: uint16(anon_sym_u128),
	10055: uint16(anon_sym_bool),
	10056: uint16(anon_sym_address),
	10057: uint16(6),
	10058: uint16(3),
	10059: uint16(1),
	10060: uint16(sym_comment),
	10061: uint16(647),
	10062: uint16(1),
	10063: uint16(sym_identifier),
	10064: uint16(653),
	10065: uint16(1),
	10066: uint16(anon_sym_LPAREN),
	10067: uint16(657),
	10068: uint16(1),
	10069: uint16(anon_sym_AMP),
	10070: uint16(290),
	10071: uint16(4),
	10072: uint16(sym__type),
	10073: uint16(sym_unit_type),
	10074: uint16(sym_reference_type),
	10075: uint16(sym__type_identifier),
	10076: uint16(655),
	10077: uint16(5),
	10078: uint16(anon_sym_u8),
	10079: uint16(anon_sym_u64),
	10080: uint16(anon_sym_u128),
	10081: uint16(anon_sym_bool),
	10082: uint16(anon_sym_address),
	10083: uint16(6),
	10084: uint16(3),
	10085: uint16(1),
	10086: uint16(sym_comment),
	10087: uint16(647),
	10088: uint16(1),
	10089: uint16(sym_identifier),
	10090: uint16(653),
	10091: uint16(1),
	10092: uint16(anon_sym_LPAREN),
	10093: uint16(657),
	10094: uint16(1),
	10095: uint16(anon_sym_AMP),
	10096: uint16(291),
	10097: uint16(4),
	10098: uint16(sym__type),
	10099: uint16(sym_unit_type),
	10100: uint16(sym_reference_type),
	10101: uint16(sym__type_identifier),
	10102: uint16(655),
	10103: uint16(5),
	10104: uint16(anon_sym_u8),
	10105: uint16(anon_sym_u64),
	10106: uint16(anon_sym_u128),
	10107: uint16(anon_sym_bool),
	10108: uint16(anon_sym_address),
	10109: uint16(6),
	10110: uint16(3),
	10111: uint16(1),
	10112: uint16(sym_comment),
	10113: uint16(647),
	10114: uint16(1),
	10115: uint16(sym_identifier),
	10116: uint16(653),
	10117: uint16(1),
	10118: uint16(anon_sym_LPAREN),
	10119: uint16(657),
	10120: uint16(1),
	10121: uint16(anon_sym_AMP),
	10122: uint16(268),
	10123: uint16(4),
	10124: uint16(sym__type),
	10125: uint16(sym_unit_type),
	10126: uint16(sym_reference_type),
	10127: uint16(sym__type_identifier),
	10128: uint16(655),
	10129: uint16(5),
	10130: uint16(anon_sym_u8),
	10131: uint16(anon_sym_u64),
	10132: uint16(anon_sym_u128),
	10133: uint16(anon_sym_bool),
	10134: uint16(anon_sym_address),
	10135: uint16(6),
	10136: uint16(3),
	10137: uint16(1),
	10138: uint16(sym_comment),
	10139: uint16(647),
	10140: uint16(1),
	10141: uint16(sym_identifier),
	10142: uint16(653),
	10143: uint16(1),
	10144: uint16(anon_sym_LPAREN),
	10145: uint16(657),
	10146: uint16(1),
	10147: uint16(anon_sym_AMP),
	10148: uint16(208),
	10149: uint16(4),
	10150: uint16(sym__type),
	10151: uint16(sym_unit_type),
	10152: uint16(sym_reference_type),
	10153: uint16(sym__type_identifier),
	10154: uint16(655),
	10155: uint16(5),
	10156: uint16(anon_sym_u8),
	10157: uint16(anon_sym_u64),
	10158: uint16(anon_sym_u128),
	10159: uint16(anon_sym_bool),
	10160: uint16(anon_sym_address),
	10161: uint16(7),
	10162: uint16(3),
	10163: uint16(1),
	10164: uint16(sym_comment),
	10165: uint16(9),
	10166: uint16(1),
	10167: uint16(anon_sym_COLON_COLON),
	10168: uint16(120),
	10169: uint16(1),
	10170: uint16(anon_sym_LT),
	10171: uint16(313),
	10172: uint16(1),
	10173: uint16(sym__use_clause),
	10174: uint16(329),
	10175: uint16(1),
	10176: uint16(sym_bracketed_type),
	10177: uint16(294),
	10178: uint16(2),
	10179: uint16(sym_scoped_identifier),
	10180: uint16(sym__path),
	10181: uint16(685),
	10182: uint16(6),
	10183: uint16(anon_sym_u8),
	10184: uint16(anon_sym_u64),
	10185: uint16(anon_sym_u128),
	10186: uint16(anon_sym_bool),
	10187: uint16(anon_sym_address),
	10188: uint16(sym_identifier),
	10189: uint16(7),
	10190: uint16(3),
	10191: uint16(1),
	10192: uint16(sym_comment),
	10193: uint16(697),
	10194: uint16(1),
	10195: uint16(anon_sym_RBRACE),
	10196: uint16(699),
	10197: uint16(1),
	10198: uint16(anon_sym_fun),
	10199: uint16(701),
	10200: uint16(1),
	10201: uint16(anon_sym_public),
	10202: uint16(703),
	10203: uint16(1),
	10204: uint16(anon_sym_use),
	10205: uint16(333),
	10206: uint16(1),
	10207: uint16(sym_visibility_modifier),
	10208: uint16(181),
	10209: uint16(4),
	10210: uint16(sym__declaration_statement),
	10211: uint16(sym_function_item),
	10212: uint16(sym_use_declaration),
	10213: uint16(aux_sym_module_body_repeat1),
	10214: uint16(7),
	10215: uint16(3),
	10216: uint16(1),
	10217: uint16(sym_comment),
	10218: uint16(699),
	10219: uint16(1),
	10220: uint16(anon_sym_fun),
	10221: uint16(701),
	10222: uint16(1),
	10223: uint16(anon_sym_public),
	10224: uint16(703),
	10225: uint16(1),
	10226: uint16(anon_sym_use),
	10227: uint16(705),
	10228: uint16(1),
	10229: uint16(anon_sym_RBRACE),
	10230: uint16(333),
	10231: uint16(1),
	10232: uint16(sym_visibility_modifier),
	10233: uint16(179),
	10234: uint16(4),
	10235: uint16(sym__declaration_statement),
	10236: uint16(sym_function_item),
	10237: uint16(sym_use_declaration),
	10238: uint16(aux_sym_module_body_repeat1),
	10239: uint16(7),
	10240: uint16(3),
	10241: uint16(1),
	10242: uint16(sym_comment),
	10243: uint16(707),
	10244: uint16(1),
	10245: uint16(anon_sym_RBRACE),
	10246: uint16(709),
	10247: uint16(1),
	10248: uint16(anon_sym_fun),
	10249: uint16(712),
	10250: uint16(1),
	10251: uint16(anon_sym_public),
	10252: uint16(715),
	10253: uint16(1),
	10254: uint16(anon_sym_use),
	10255: uint16(333),
	10256: uint16(1),
	10257: uint16(sym_visibility_modifier),
	10258: uint16(181),
	10259: uint16(4),
	10260: uint16(sym__declaration_statement),
	10261: uint16(sym_function_item),
	10262: uint16(sym_use_declaration),
	10263: uint16(aux_sym_module_body_repeat1),
	10264: uint16(4),
	10265: uint16(3),
	10266: uint16(1),
	10267: uint16(sym_comment),
	10268: uint16(440),
	10269: uint16(1),
	10270: uint16(anon_sym_COLON_COLON),
	10271: uint16(600),
	10272: uint16(2),
	10273: uint16(anon_sym_LBRACE),
	10274: uint16(anon_sym_LT2),
	10275: uint16(718),
	10276: uint16(5),
	10277: uint16(anon_sym_RBRACE),
	10278: uint16(anon_sym_EQ),
	10279: uint16(anon_sym_COMMA),
	10280: uint16(anon_sym_RPAREN),
	10281: uint16(anon_sym_PIPE),
	10282: uint16(3),
	10283: uint16(3),
	10284: uint16(1),
	10285: uint16(sym_comment),
	10286: uint16(600),
	10287: uint16(2),
	10288: uint16(anon_sym_LBRACE),
	10289: uint16(anon_sym_LT2),
	10290: uint16(720),
	10291: uint16(6),
	10292: uint16(anon_sym_COLON_COLON),
	10293: uint16(anon_sym_RBRACE),
	10294: uint16(anon_sym_EQ),
	10295: uint16(anon_sym_COMMA),
	10296: uint16(anon_sym_RPAREN),
	10297: uint16(anon_sym_PIPE),
	10298: uint16(3),
	10299: uint16(3),
	10300: uint16(1),
	10301: uint16(sym_comment),
	10302: uint16(600),
	10303: uint16(2),
	10304: uint16(anon_sym_LBRACE),
	10305: uint16(anon_sym_LT2),
	10306: uint16(722),
	10307: uint16(6),
	10308: uint16(anon_sym_COLON_COLON),
	10309: uint16(anon_sym_RBRACE),
	10310: uint16(anon_sym_EQ),
	10311: uint16(anon_sym_COMMA),
	10312: uint16(anon_sym_RPAREN),
	10313: uint16(anon_sym_PIPE),
	10314: uint16(8),
	10315: uint16(3),
	10316: uint16(1),
	10317: uint16(sym_comment),
	10318: uint16(724),
	10319: uint16(1),
	10320: uint16(anon_sym_COLON_COLON),
	10321: uint16(726),
	10322: uint16(1),
	10323: uint16(anon_sym_LBRACE),
	10324: uint16(728),
	10325: uint16(1),
	10326: uint16(anon_sym_LBRACK),
	10327: uint16(730),
	10328: uint16(1),
	10329: uint16(anon_sym_RBRACK),
	10330: uint16(732),
	10331: uint16(1),
	10332: uint16(anon_sym_EQ),
	10333: uint16(734),
	10334: uint16(1),
	10335: uint16(anon_sym_LPAREN),
	10336: uint16(314),
	10337: uint16(1),
	10338: uint16(sym_delim_token_tree),
	10339: uint16(2),
	10340: uint16(3),
	10341: uint16(1),
	10342: uint16(sym_comment),
	10343: uint16(720),
	10344: uint16(7),
	10345: uint16(anon_sym_COLON_COLON),
	10346: uint16(anon_sym_LBRACE),
	10347: uint16(anon_sym_SEMI),
	10348: uint16(anon_sym_LBRACK),
	10349: uint16(anon_sym_RBRACK),
	10350: uint16(anon_sym_EQ),
	10351: uint16(anon_sym_LPAREN),
	10352: uint16(2),
	10353: uint16(3),
	10354: uint16(1),
	10355: uint16(sym_comment),
	10356: uint16(722),
	10357: uint16(7),
	10358: uint16(anon_sym_COLON_COLON),
	10359: uint16(anon_sym_LBRACE),
	10360: uint16(anon_sym_SEMI),
	10361: uint16(anon_sym_LBRACK),
	10362: uint16(anon_sym_RBRACK),
	10363: uint16(anon_sym_EQ),
	10364: uint16(anon_sym_LPAREN),
	10365: uint16(7),
	10366: uint16(3),
	10367: uint16(1),
	10368: uint16(sym_comment),
	10369: uint16(736),
	10370: uint16(1),
	10371: uint16(sym_identifier),
	10372: uint16(738),
	10373: uint16(1),
	10374: uint16(anon_sym_RBRACE),
	10375: uint16(740),
	10376: uint16(1),
	10377: uint16(anon_sym_COMMA),
	10378: uint16(742),
	10379: uint16(1),
	10380: uint16(sym_mutable_specifier),
	10381: uint16(249),
	10382: uint16(1),
	10383: uint16(sym_field_pattern),
	10384: uint16(328),
	10385: uint16(1),
	10386: uint16(sym__field_identifier),
	10387: uint16(2),
	10388: uint16(3),
	10389: uint16(1),
	10390: uint16(sym_comment),
	10391: uint16(600),
	10392: uint16(6),
	10393: uint16(anon_sym_LBRACE),
	10394: uint16(anon_sym_COMMA),
	10395: uint16(anon_sym_GT),
	10396: uint16(anon_sym_RPAREN),
	10397: uint16(anon_sym_as),
	10398: uint16(anon_sym_LT2),
	10399: uint16(3),
	10400: uint16(3),
	10401: uint16(1),
	10402: uint16(sym_comment),
	10403: uint16(440),
	10404: uint16(1),
	10405: uint16(anon_sym_COLON_COLON),
	10406: uint16(718),
	10407: uint16(5),
	10408: uint16(anon_sym_RBRACE),
	10409: uint16(anon_sym_EQ),
	10410: uint16(anon_sym_COMMA),
	10411: uint16(anon_sym_RPAREN),
	10412: uint16(anon_sym_PIPE),
	10413: uint16(2),
	10414: uint16(3),
	10415: uint16(1),
	10416: uint16(sym_comment),
	10417: uint16(744),
	10418: uint16(5),
	10419: uint16(anon_sym_RBRACE),
	10420: uint16(anon_sym_EQ),
	10421: uint16(anon_sym_COMMA),
	10422: uint16(anon_sym_RPAREN),
	10423: uint16(anon_sym_PIPE),
	10424: uint16(2),
	10425: uint16(3),
	10426: uint16(1),
	10427: uint16(sym_comment),
	10428: uint16(746),
	10429: uint16(5),
	10430: uint16(anon_sym_RBRACE),
	10431: uint16(anon_sym_EQ),
	10432: uint16(anon_sym_COMMA),
	10433: uint16(anon_sym_RPAREN),
	10434: uint16(anon_sym_PIPE),
	10435: uint16(2),
	10436: uint16(3),
	10437: uint16(1),
	10438: uint16(sym_comment),
	10439: uint16(748),
	10440: uint16(5),
	10441: uint16(anon_sym_RBRACE),
	10442: uint16(anon_sym_EQ),
	10443: uint16(anon_sym_COMMA),
	10444: uint16(anon_sym_RPAREN),
	10445: uint16(anon_sym_PIPE),
	10446: uint16(6),
	10447: uint16(3),
	10448: uint16(1),
	10449: uint16(sym_comment),
	10450: uint16(736),
	10451: uint16(1),
	10452: uint16(sym_identifier),
	10453: uint16(742),
	10454: uint16(1),
	10455: uint16(sym_mutable_specifier),
	10456: uint16(750),
	10457: uint16(1),
	10458: uint16(anon_sym_RBRACE),
	10459: uint16(279),
	10460: uint16(1),
	10461: uint16(sym_field_pattern),
	10462: uint16(328),
	10463: uint16(1),
	10464: uint16(sym__field_identifier),
	10465: uint16(2),
	10466: uint16(3),
	10467: uint16(1),
	10468: uint16(sym_comment),
	10469: uint16(752),
	10470: uint16(5),
	10471: uint16(anon_sym_RBRACE),
	10472: uint16(anon_sym_EQ),
	10473: uint16(anon_sym_COMMA),
	10474: uint16(anon_sym_RPAREN),
	10475: uint16(anon_sym_PIPE),
	10476: uint16(2),
	10477: uint16(3),
	10478: uint16(1),
	10479: uint16(sym_comment),
	10480: uint16(754),
	10481: uint16(5),
	10482: uint16(anon_sym_RBRACE),
	10483: uint16(anon_sym_EQ),
	10484: uint16(anon_sym_COMMA),
	10485: uint16(anon_sym_RPAREN),
	10486: uint16(anon_sym_PIPE),
	10487: uint16(2),
	10488: uint16(3),
	10489: uint16(1),
	10490: uint16(sym_comment),
	10491: uint16(756),
	10492: uint16(5),
	10493: uint16(anon_sym_RBRACE),
	10494: uint16(anon_sym_EQ),
	10495: uint16(anon_sym_COMMA),
	10496: uint16(anon_sym_RPAREN),
	10497: uint16(anon_sym_PIPE),
	10498: uint16(2),
	10499: uint16(3),
	10500: uint16(1),
	10501: uint16(sym_comment),
	10502: uint16(758),
	10503: uint16(5),
	10504: uint16(anon_sym_RBRACE),
	10505: uint16(anon_sym_EQ),
	10506: uint16(anon_sym_COMMA),
	10507: uint16(anon_sym_RPAREN),
	10508: uint16(anon_sym_PIPE),
	10509: uint16(6),
	10510: uint16(3),
	10511: uint16(1),
	10512: uint16(sym_comment),
	10513: uint16(736),
	10514: uint16(1),
	10515: uint16(sym_identifier),
	10516: uint16(742),
	10517: uint16(1),
	10518: uint16(sym_mutable_specifier),
	10519: uint16(760),
	10520: uint16(1),
	10521: uint16(anon_sym_RBRACE),
	10522: uint16(279),
	10523: uint16(1),
	10524: uint16(sym_field_pattern),
	10525: uint16(328),
	10526: uint16(1),
	10527: uint16(sym__field_identifier),
	10528: uint16(2),
	10529: uint16(3),
	10530: uint16(1),
	10531: uint16(sym_comment),
	10532: uint16(762),
	10533: uint16(5),
	10534: uint16(anon_sym_RBRACE),
	10535: uint16(anon_sym_EQ),
	10536: uint16(anon_sym_COMMA),
	10537: uint16(anon_sym_RPAREN),
	10538: uint16(anon_sym_PIPE),
	10539: uint16(2),
	10540: uint16(3),
	10541: uint16(1),
	10542: uint16(sym_comment),
	10543: uint16(764),
	10544: uint16(5),
	10545: uint16(anon_sym_RBRACE),
	10546: uint16(anon_sym_EQ),
	10547: uint16(anon_sym_COMMA),
	10548: uint16(anon_sym_RPAREN),
	10549: uint16(anon_sym_PIPE),
	10550: uint16(2),
	10551: uint16(3),
	10552: uint16(1),
	10553: uint16(sym_comment),
	10554: uint16(766),
	10555: uint16(5),
	10556: uint16(anon_sym_RBRACE),
	10557: uint16(anon_sym_EQ),
	10558: uint16(anon_sym_COMMA),
	10559: uint16(anon_sym_RPAREN),
	10560: uint16(anon_sym_PIPE),
	10561: uint16(2),
	10562: uint16(3),
	10563: uint16(1),
	10564: uint16(sym_comment),
	10565: uint16(588),
	10566: uint16(5),
	10567: uint16(anon_sym_LBRACE),
	10568: uint16(anon_sym_COMMA),
	10569: uint16(anon_sym_GT),
	10570: uint16(anon_sym_RPAREN),
	10571: uint16(anon_sym_as),
	10572: uint16(2),
	10573: uint16(3),
	10574: uint16(1),
	10575: uint16(sym_comment),
	10576: uint16(768),
	10577: uint16(5),
	10578: uint16(anon_sym_RBRACE),
	10579: uint16(anon_sym_EQ),
	10580: uint16(anon_sym_COMMA),
	10581: uint16(anon_sym_RPAREN),
	10582: uint16(anon_sym_PIPE),
	10583: uint16(2),
	10584: uint16(3),
	10585: uint16(1),
	10586: uint16(sym_comment),
	10587: uint16(592),
	10588: uint16(5),
	10589: uint16(anon_sym_LBRACE),
	10590: uint16(anon_sym_COMMA),
	10591: uint16(anon_sym_GT),
	10592: uint16(anon_sym_RPAREN),
	10593: uint16(anon_sym_as),
	10594: uint16(2),
	10595: uint16(3),
	10596: uint16(1),
	10597: uint16(sym_comment),
	10598: uint16(770),
	10599: uint16(5),
	10600: uint16(anon_sym_RBRACE),
	10601: uint16(anon_sym_EQ),
	10602: uint16(anon_sym_COMMA),
	10603: uint16(anon_sym_RPAREN),
	10604: uint16(anon_sym_PIPE),
	10605: uint16(2),
	10606: uint16(3),
	10607: uint16(1),
	10608: uint16(sym_comment),
	10609: uint16(772),
	10610: uint16(5),
	10611: uint16(anon_sym_RBRACE),
	10612: uint16(anon_sym_EQ),
	10613: uint16(anon_sym_COMMA),
	10614: uint16(anon_sym_RPAREN),
	10615: uint16(anon_sym_PIPE),
	10616: uint16(2),
	10617: uint16(3),
	10618: uint16(1),
	10619: uint16(sym_comment),
	10620: uint16(540),
	10621: uint16(5),
	10622: uint16(anon_sym_LBRACE),
	10623: uint16(anon_sym_COMMA),
	10624: uint16(anon_sym_GT),
	10625: uint16(anon_sym_RPAREN),
	10626: uint16(anon_sym_as),
	10627: uint16(2),
	10628: uint16(3),
	10629: uint16(1),
	10630: uint16(sym_comment),
	10631: uint16(468),
	10632: uint16(4),
	10633: uint16(anon_sym_RBRACE),
	10634: uint16(anon_sym_fun),
	10635: uint16(anon_sym_public),
	10636: uint16(anon_sym_use),
	10637: uint16(2),
	10638: uint16(3),
	10639: uint16(1),
	10640: uint16(sym_comment),
	10641: uint16(476),
	10642: uint16(4),
	10643: uint16(anon_sym_RBRACE),
	10644: uint16(anon_sym_fun),
	10645: uint16(anon_sym_public),
	10646: uint16(anon_sym_use),
	10647: uint16(5),
	10648: uint16(3),
	10649: uint16(1),
	10650: uint16(sym_comment),
	10651: uint16(774),
	10652: uint16(1),
	10653: uint16(anon_sym_LT),
	10654: uint16(776),
	10655: uint16(1),
	10656: uint16(anon_sym_LPAREN),
	10657: uint16(232),
	10658: uint16(1),
	10659: uint16(sym_parameters),
	10660: uint16(265),
	10661: uint16(1),
	10662: uint16(sym_type_parameters),
	10663: uint16(5),
	10664: uint16(3),
	10665: uint16(1),
	10666: uint16(sym_comment),
	10667: uint16(774),
	10668: uint16(1),
	10669: uint16(anon_sym_LT),
	10670: uint16(776),
	10671: uint16(1),
	10672: uint16(anon_sym_LPAREN),
	10673: uint16(234),
	10674: uint16(1),
	10675: uint16(sym_parameters),
	10676: uint16(273),
	10677: uint16(1),
	10678: uint16(sym_type_parameters),
	10679: uint16(4),
	10680: uint16(3),
	10681: uint16(1),
	10682: uint16(sym_comment),
	10683: uint16(11),
	10684: uint16(1),
	10685: uint16(anon_sym_LBRACE),
	10686: uint16(778),
	10687: uint16(1),
	10688: uint16(anon_sym_if),
	10689: uint16(116),
	10690: uint16(2),
	10691: uint16(sym_if_expression),
	10692: uint16(sym_block),
	10693: uint16(5),
	10694: uint16(3),
	10695: uint16(1),
	10696: uint16(sym_comment),
	10697: uint16(774),
	10698: uint16(1),
	10699: uint16(anon_sym_LT),
	10700: uint16(776),
	10701: uint16(1),
	10702: uint16(anon_sym_LPAREN),
	10703: uint16(235),
	10704: uint16(1),
	10705: uint16(sym_parameters),
	10706: uint16(304),
	10707: uint16(1),
	10708: uint16(sym_type_parameters),
	10709: uint16(2),
	10710: uint16(3),
	10711: uint16(1),
	10712: uint16(sym_comment),
	10713: uint16(444),
	10714: uint16(4),
	10715: uint16(anon_sym_RBRACE),
	10716: uint16(anon_sym_fun),
	10717: uint16(anon_sym_public),
	10718: uint16(anon_sym_use),
	10719: uint16(5),
	10720: uint16(3),
	10721: uint16(1),
	10722: uint16(sym_comment),
	10723: uint16(774),
	10724: uint16(1),
	10725: uint16(anon_sym_LT),
	10726: uint16(776),
	10727: uint16(1),
	10728: uint16(anon_sym_LPAREN),
	10729: uint16(242),
	10730: uint16(1),
	10731: uint16(sym_parameters),
	10732: uint16(293),
	10733: uint16(1),
	10734: uint16(sym_type_parameters),
	10735: uint16(2),
	10736: uint16(3),
	10737: uint16(1),
	10738: uint16(sym_comment),
	10739: uint16(438),
	10740: uint16(4),
	10741: uint16(anon_sym_RBRACE),
	10742: uint16(anon_sym_fun),
	10743: uint16(anon_sym_public),
	10744: uint16(anon_sym_use),
	10745: uint16(2),
	10746: uint16(3),
	10747: uint16(1),
	10748: uint16(sym_comment),
	10749: uint16(410),
	10750: uint16(4),
	10751: uint16(anon_sym_RBRACE),
	10752: uint16(anon_sym_fun),
	10753: uint16(anon_sym_public),
	10754: uint16(anon_sym_use),
	10755: uint16(2),
	10756: uint16(3),
	10757: uint16(1),
	10758: uint16(sym_comment),
	10759: uint16(460),
	10760: uint16(4),
	10761: uint16(anon_sym_RBRACE),
	10762: uint16(anon_sym_fun),
	10763: uint16(anon_sym_public),
	10764: uint16(anon_sym_use),
	10765: uint16(5),
	10766: uint16(3),
	10767: uint16(1),
	10768: uint16(sym_comment),
	10769: uint16(736),
	10770: uint16(1),
	10771: uint16(sym_identifier),
	10772: uint16(742),
	10773: uint16(1),
	10774: uint16(sym_mutable_specifier),
	10775: uint16(279),
	10776: uint16(1),
	10777: uint16(sym_field_pattern),
	10778: uint16(328),
	10779: uint16(1),
	10780: uint16(sym__field_identifier),
	10781: uint16(5),
	10782: uint16(3),
	10783: uint16(1),
	10784: uint16(sym_comment),
	10785: uint16(780),
	10786: uint16(1),
	10787: uint16(anon_sym_COMMA),
	10788: uint16(782),
	10789: uint16(1),
	10790: uint16(anon_sym_RPAREN),
	10791: uint16(784),
	10792: uint16(1),
	10793: uint16(anon_sym_PIPE),
	10794: uint16(241),
	10795: uint16(1),
	10796: uint16(aux_sym_tuple_pattern_repeat1),
	10797: uint16(2),
	10798: uint16(3),
	10799: uint16(1),
	10800: uint16(sym_comment),
	10801: uint16(430),
	10802: uint16(4),
	10803: uint16(anon_sym_RBRACE),
	10804: uint16(anon_sym_fun),
	10805: uint16(anon_sym_public),
	10806: uint16(anon_sym_use),
	10807: uint16(4),
	10808: uint16(3),
	10809: uint16(1),
	10810: uint16(sym_comment),
	10811: uint16(112),
	10812: uint16(1),
	10813: uint16(anon_sym_LBRACE),
	10814: uint16(786),
	10815: uint16(1),
	10816: uint16(anon_sym_if),
	10817: uint16(46),
	10818: uint16(2),
	10819: uint16(sym_if_expression),
	10820: uint16(sym_block),
	10821: uint16(2),
	10822: uint16(3),
	10823: uint16(1),
	10824: uint16(sym_comment),
	10825: uint16(452),
	10826: uint16(4),
	10827: uint16(anon_sym_RBRACE),
	10828: uint16(anon_sym_fun),
	10829: uint16(anon_sym_public),
	10830: uint16(anon_sym_use),
	10831: uint16(2),
	10832: uint16(3),
	10833: uint16(1),
	10834: uint16(sym_comment),
	10835: uint16(480),
	10836: uint16(4),
	10837: uint16(anon_sym_RBRACE),
	10838: uint16(anon_sym_fun),
	10839: uint16(anon_sym_public),
	10840: uint16(anon_sym_use),
	10841: uint16(3),
	10842: uint16(3),
	10843: uint16(1),
	10844: uint16(sym_comment),
	10845: uint16(784),
	10846: uint16(1),
	10847: uint16(anon_sym_PIPE),
	10848: uint16(788),
	10849: uint16(2),
	10850: uint16(anon_sym_RBRACE),
	10851: uint16(anon_sym_COMMA),
	10852: uint16(4),
	10853: uint16(3),
	10854: uint16(1),
	10855: uint16(sym_comment),
	10856: uint16(790),
	10857: uint16(1),
	10858: uint16(anon_sym_COMMA),
	10859: uint16(793),
	10860: uint16(1),
	10861: uint16(anon_sym_RPAREN),
	10862: uint16(227),
	10863: uint16(1),
	10864: uint16(aux_sym_parameters_repeat1),
	10865: uint16(4),
	10866: uint16(3),
	10867: uint16(1),
	10868: uint16(sym_comment),
	10869: uint16(795),
	10870: uint16(1),
	10871: uint16(anon_sym_COMMA),
	10872: uint16(797),
	10873: uint16(1),
	10874: uint16(anon_sym_GT),
	10875: uint16(230),
	10876: uint16(1),
	10877: uint16(aux_sym_type_arguments_repeat1),
	10878: uint16(4),
	10879: uint16(3),
	10880: uint16(1),
	10881: uint16(sym_comment),
	10882: uint16(649),
	10883: uint16(1),
	10884: uint16(anon_sym_LBRACE),
	10885: uint16(799),
	10886: uint16(1),
	10887: uint16(anon_sym_COLON),
	10888: uint16(100),
	10889: uint16(1),
	10890: uint16(sym_block),
	10891: uint16(4),
	10892: uint16(3),
	10893: uint16(1),
	10894: uint16(sym_comment),
	10895: uint16(665),
	10896: uint16(1),
	10897: uint16(anon_sym_GT),
	10898: uint16(801),
	10899: uint16(1),
	10900: uint16(anon_sym_COMMA),
	10901: uint16(256),
	10902: uint16(1),
	10903: uint16(aux_sym_type_arguments_repeat1),
	10904: uint16(4),
	10905: uint16(3),
	10906: uint16(1),
	10907: uint16(sym_comment),
	10908: uint16(635),
	10909: uint16(1),
	10910: uint16(anon_sym_RPAREN),
	10911: uint16(803),
	10912: uint16(1),
	10913: uint16(anon_sym_COMMA),
	10914: uint16(231),
	10915: uint16(1),
	10916: uint16(aux_sym_arguments_repeat1),
	10917: uint16(4),
	10918: uint16(3),
	10919: uint16(1),
	10920: uint16(sym_comment),
	10921: uint16(649),
	10922: uint16(1),
	10923: uint16(anon_sym_LBRACE),
	10924: uint16(806),
	10925: uint16(1),
	10926: uint16(anon_sym_COLON),
	10927: uint16(97),
	10928: uint16(1),
	10929: uint16(sym_block),
	10930: uint16(4),
	10931: uint16(3),
	10932: uint16(1),
	10933: uint16(sym_comment),
	10934: uint16(649),
	10935: uint16(1),
	10936: uint16(anon_sym_LBRACE),
	10937: uint16(808),
	10938: uint16(1),
	10939: uint16(anon_sym_COLON),
	10940: uint16(95),
	10941: uint16(1),
	10942: uint16(sym_block),
	10943: uint16(4),
	10944: uint16(3),
	10945: uint16(1),
	10946: uint16(sym_comment),
	10947: uint16(649),
	10948: uint16(1),
	10949: uint16(anon_sym_LBRACE),
	10950: uint16(810),
	10951: uint16(1),
	10952: uint16(anon_sym_COLON),
	10953: uint16(81),
	10954: uint16(1),
	10955: uint16(sym_block),
	10956: uint16(4),
	10957: uint16(3),
	10958: uint16(1),
	10959: uint16(sym_comment),
	10960: uint16(11),
	10961: uint16(1),
	10962: uint16(anon_sym_LBRACE),
	10963: uint16(812),
	10964: uint16(1),
	10965: uint16(anon_sym_COLON),
	10966: uint16(218),
	10967: uint16(1),
	10968: uint16(sym_block),
	10969: uint16(4),
	10970: uint16(3),
	10971: uint16(1),
	10972: uint16(sym_comment),
	10973: uint16(220),
	10974: uint16(1),
	10975: uint16(anon_sym_RPAREN),
	10976: uint16(814),
	10977: uint16(1),
	10978: uint16(anon_sym_COMMA),
	10979: uint16(231),
	10980: uint16(1),
	10981: uint16(aux_sym_arguments_repeat1),
	10982: uint16(4),
	10983: uint16(3),
	10984: uint16(1),
	10985: uint16(sym_comment),
	10986: uint16(816),
	10987: uint16(1),
	10988: uint16(anon_sym_COMMA),
	10989: uint16(818),
	10990: uint16(1),
	10991: uint16(anon_sym_GT),
	10992: uint16(243),
	10993: uint16(1),
	10994: uint16(aux_sym_type_parameters_repeat1),
	10995: uint16(4),
	10996: uint16(3),
	10997: uint16(1),
	10998: uint16(sym_comment),
	10999: uint16(820),
	11000: uint16(1),
	11001: uint16(anon_sym_COMMA),
	11002: uint16(822),
	11003: uint16(1),
	11004: uint16(anon_sym_RPAREN),
	11005: uint16(244),
	11006: uint16(1),
	11007: uint16(aux_sym_parameters_repeat1),
	11008: uint16(3),
	11009: uint16(3),
	11010: uint16(1),
	11011: uint16(sym_comment),
	11012: uint16(824),
	11013: uint16(1),
	11014: uint16(anon_sym_COLON),
	11015: uint16(600),
	11016: uint16(2),
	11017: uint16(anon_sym_COMMA),
	11018: uint16(anon_sym_RPAREN),
	11019: uint16(4),
	11020: uint16(3),
	11021: uint16(1),
	11022: uint16(sym_comment),
	11023: uint16(11),
	11024: uint16(1),
	11025: uint16(anon_sym_LBRACE),
	11026: uint16(826),
	11027: uint16(1),
	11028: uint16(anon_sym_COLON),
	11029: uint16(219),
	11030: uint16(1),
	11031: uint16(sym_block),
	11032: uint16(4),
	11033: uint16(3),
	11034: uint16(1),
	11035: uint16(sym_comment),
	11036: uint16(370),
	11037: uint16(1),
	11038: uint16(anon_sym_RPAREN),
	11039: uint16(828),
	11040: uint16(1),
	11041: uint16(anon_sym_COMMA),
	11042: uint16(257),
	11043: uint16(1),
	11044: uint16(aux_sym_tuple_pattern_repeat1),
	11045: uint16(4),
	11046: uint16(3),
	11047: uint16(1),
	11048: uint16(sym_comment),
	11049: uint16(11),
	11050: uint16(1),
	11051: uint16(anon_sym_LBRACE),
	11052: uint16(830),
	11053: uint16(1),
	11054: uint16(anon_sym_COLON),
	11055: uint16(209),
	11056: uint16(1),
	11057: uint16(sym_block),
	11058: uint16(4),
	11059: uint16(3),
	11060: uint16(1),
	11061: uint16(sym_comment),
	11062: uint16(832),
	11063: uint16(1),
	11064: uint16(anon_sym_COMMA),
	11065: uint16(834),
	11066: uint16(1),
	11067: uint16(anon_sym_GT),
	11068: uint16(258),
	11069: uint16(1),
	11070: uint16(aux_sym_type_parameters_repeat1),
	11071: uint16(4),
	11072: uint16(3),
	11073: uint16(1),
	11074: uint16(sym_comment),
	11075: uint16(681),
	11076: uint16(1),
	11077: uint16(anon_sym_RPAREN),
	11078: uint16(836),
	11079: uint16(1),
	11080: uint16(anon_sym_COMMA),
	11081: uint16(227),
	11082: uint16(1),
	11083: uint16(aux_sym_parameters_repeat1),
	11084: uint16(4),
	11085: uint16(3),
	11086: uint16(1),
	11087: uint16(sym_comment),
	11088: uint16(834),
	11089: uint16(1),
	11090: uint16(anon_sym_GT),
	11091: uint16(838),
	11092: uint16(1),
	11093: uint16(sym_identifier),
	11094: uint16(272),
	11095: uint16(1),
	11096: uint16(sym__type_identifier),
	11097: uint16(4),
	11098: uint16(3),
	11099: uint16(1),
	11100: uint16(sym_comment),
	11101: uint16(840),
	11102: uint16(1),
	11103: uint16(anon_sym_LBRACE),
	11104: uint16(842),
	11105: uint16(1),
	11106: uint16(anon_sym_LT2),
	11107: uint16(332),
	11108: uint16(1),
	11109: uint16(sym_type_arguments),
	11110: uint16(3),
	11111: uint16(3),
	11112: uint16(1),
	11113: uint16(sym_comment),
	11114: uint16(784),
	11115: uint16(1),
	11116: uint16(anon_sym_PIPE),
	11117: uint16(844),
	11118: uint16(2),
	11119: uint16(anon_sym_RBRACE),
	11120: uint16(anon_sym_COMMA),
	11121: uint16(3),
	11122: uint16(3),
	11123: uint16(1),
	11124: uint16(sym_comment),
	11125: uint16(462),
	11126: uint16(1),
	11127: uint16(anon_sym_COLON),
	11128: uint16(846),
	11129: uint16(2),
	11130: uint16(anon_sym_RBRACE),
	11131: uint16(anon_sym_COMMA),
	11132: uint16(4),
	11133: uint16(3),
	11134: uint16(1),
	11135: uint16(sym_comment),
	11136: uint16(848),
	11137: uint16(1),
	11138: uint16(anon_sym_RBRACE),
	11139: uint16(850),
	11140: uint16(1),
	11141: uint16(anon_sym_COMMA),
	11142: uint16(255),
	11143: uint16(1),
	11144: uint16(aux_sym_struct_pattern_repeat1),
	11145: uint16(4),
	11146: uint16(3),
	11147: uint16(1),
	11148: uint16(sym_comment),
	11149: uint16(852),
	11150: uint16(1),
	11151: uint16(anon_sym_RBRACE),
	11152: uint16(854),
	11153: uint16(1),
	11154: uint16(anon_sym_COMMA),
	11155: uint16(250),
	11156: uint16(1),
	11157: uint16(aux_sym_struct_pattern_repeat1),
	11158: uint16(4),
	11159: uint16(3),
	11160: uint16(1),
	11161: uint16(sym_comment),
	11162: uint16(838),
	11163: uint16(1),
	11164: uint16(sym_identifier),
	11165: uint16(857),
	11166: uint16(1),
	11167: uint16(anon_sym_GT),
	11168: uint16(272),
	11169: uint16(1),
	11170: uint16(sym__type_identifier),
	11171: uint16(3),
	11172: uint16(3),
	11173: uint16(1),
	11174: uint16(sym_comment),
	11175: uint16(347),
	11176: uint16(1),
	11177: uint16(sym_hex_address),
	11178: uint16(859),
	11179: uint16(2),
	11180: uint16(anon_sym_0x),
	11181: uint16(anon_sym_0X),
	11182: uint16(4),
	11183: uint16(3),
	11184: uint16(1),
	11185: uint16(sym_comment),
	11186: uint16(861),
	11187: uint16(1),
	11188: uint16(sym_identifier),
	11189: uint16(863),
	11190: uint16(1),
	11191: uint16(sym_integer_literal),
	11192: uint16(94),
	11193: uint16(1),
	11194: uint16(sym__field_identifier),
	11195: uint16(3),
	11196: uint16(3),
	11197: uint16(1),
	11198: uint16(sym_comment),
	11199: uint16(784),
	11200: uint16(1),
	11201: uint16(anon_sym_PIPE),
	11202: uint16(865),
	11203: uint16(2),
	11204: uint16(anon_sym_COMMA),
	11205: uint16(anon_sym_RPAREN),
	11206: uint16(4),
	11207: uint16(3),
	11208: uint16(1),
	11209: uint16(sym_comment),
	11210: uint16(760),
	11211: uint16(1),
	11212: uint16(anon_sym_RBRACE),
	11213: uint16(867),
	11214: uint16(1),
	11215: uint16(anon_sym_COMMA),
	11216: uint16(250),
	11217: uint16(1),
	11218: uint16(aux_sym_struct_pattern_repeat1),
	11219: uint16(4),
	11220: uint16(3),
	11221: uint16(1),
	11222: uint16(sym_comment),
	11223: uint16(869),
	11224: uint16(1),
	11225: uint16(anon_sym_COMMA),
	11226: uint16(872),
	11227: uint16(1),
	11228: uint16(anon_sym_GT),
	11229: uint16(256),
	11230: uint16(1),
	11231: uint16(aux_sym_type_arguments_repeat1),
	11232: uint16(4),
	11233: uint16(3),
	11234: uint16(1),
	11235: uint16(sym_comment),
	11236: uint16(865),
	11237: uint16(1),
	11238: uint16(anon_sym_RPAREN),
	11239: uint16(874),
	11240: uint16(1),
	11241: uint16(anon_sym_COMMA),
	11242: uint16(257),
	11243: uint16(1),
	11244: uint16(aux_sym_tuple_pattern_repeat1),
	11245: uint16(4),
	11246: uint16(3),
	11247: uint16(1),
	11248: uint16(sym_comment),
	11249: uint16(877),
	11250: uint16(1),
	11251: uint16(anon_sym_COMMA),
	11252: uint16(880),
	11253: uint16(1),
	11254: uint16(anon_sym_GT),
	11255: uint16(258),
	11256: uint16(1),
	11257: uint16(aux_sym_type_parameters_repeat1),
	11258: uint16(3),
	11259: uint16(3),
	11260: uint16(1),
	11261: uint16(sym_comment),
	11262: uint16(462),
	11263: uint16(1),
	11264: uint16(anon_sym_COLON),
	11265: uint16(882),
	11266: uint16(2),
	11267: uint16(anon_sym_RBRACE),
	11268: uint16(anon_sym_COMMA),
	11269: uint16(4),
	11270: uint16(3),
	11271: uint16(1),
	11272: uint16(sym_comment),
	11273: uint16(11),
	11274: uint16(1),
	11275: uint16(anon_sym_LBRACE),
	11276: uint16(884),
	11277: uint16(1),
	11278: uint16(anon_sym_COLON),
	11279: uint16(225),
	11280: uint16(1),
	11281: uint16(sym_block),
	11282: uint16(3),
	11283: uint16(3),
	11284: uint16(1),
	11285: uint16(sym_comment),
	11286: uint16(886),
	11287: uint16(1),
	11288: uint16(sym_identifier),
	11289: uint16(289),
	11290: uint16(1),
	11291: uint16(sym__type_identifier),
	11292: uint16(3),
	11293: uint16(3),
	11294: uint16(1),
	11295: uint16(sym_comment),
	11296: uint16(888),
	11297: uint16(1),
	11298: uint16(anon_sym_GT),
	11299: uint16(890),
	11300: uint16(1),
	11301: uint16(anon_sym_as),
	11302: uint16(2),
	11303: uint16(3),
	11304: uint16(1),
	11305: uint16(sym_comment),
	11306: uint16(872),
	11307: uint16(2),
	11308: uint16(anon_sym_COMMA),
	11309: uint16(anon_sym_GT),
	11310: uint16(3),
	11311: uint16(3),
	11312: uint16(1),
	11313: uint16(sym_comment),
	11314: uint16(842),
	11315: uint16(1),
	11316: uint16(anon_sym_LT2),
	11317: uint16(124),
	11318: uint16(1),
	11319: uint16(sym_type_arguments),
	11320: uint16(3),
	11321: uint16(3),
	11322: uint16(1),
	11323: uint16(sym_comment),
	11324: uint16(776),
	11325: uint16(1),
	11326: uint16(anon_sym_LPAREN),
	11327: uint16(229),
	11328: uint16(1),
	11329: uint16(sym_parameters),
	11330: uint16(3),
	11331: uint16(3),
	11332: uint16(1),
	11333: uint16(sym_comment),
	11334: uint16(11),
	11335: uint16(1),
	11336: uint16(anon_sym_LBRACE),
	11337: uint16(222),
	11338: uint16(1),
	11339: uint16(sym_block),
	11340: uint16(3),
	11341: uint16(3),
	11342: uint16(1),
	11343: uint16(sym_comment),
	11344: uint16(892),
	11345: uint16(1),
	11346: uint16(anon_sym_LBRACE),
	11347: uint16(339),
	11348: uint16(1),
	11349: uint16(sym_module_body),
	11350: uint16(2),
	11351: uint16(3),
	11352: uint16(1),
	11353: uint16(sym_comment),
	11354: uint16(894),
	11355: uint16(2),
	11356: uint16(anon_sym_COMMA),
	11357: uint16(anon_sym_RPAREN),
	11358: uint16(2),
	11359: uint16(3),
	11360: uint16(1),
	11361: uint16(sym_comment),
	11362: uint16(896),
	11363: uint16(2),
	11364: uint16(anon_sym_LBRACE),
	11365: uint16(anon_sym_COLON),
	11366: uint16(2),
	11367: uint16(3),
	11368: uint16(1),
	11369: uint16(sym_comment),
	11370: uint16(793),
	11371: uint16(2),
	11372: uint16(anon_sym_COMMA),
	11373: uint16(anon_sym_RPAREN),
	11374: uint16(3),
	11375: uint16(3),
	11376: uint16(1),
	11377: uint16(sym_comment),
	11378: uint16(784),
	11379: uint16(1),
	11380: uint16(anon_sym_PIPE),
	11381: uint16(898),
	11382: uint16(1),
	11383: uint16(anon_sym_EQ),
	11384: uint16(2),
	11385: uint16(3),
	11386: uint16(1),
	11387: uint16(sym_comment),
	11388: uint16(880),
	11389: uint16(2),
	11390: uint16(anon_sym_COMMA),
	11391: uint16(anon_sym_GT),
	11392: uint16(3),
	11393: uint16(3),
	11394: uint16(1),
	11395: uint16(sym_comment),
	11396: uint16(776),
	11397: uint16(1),
	11398: uint16(anon_sym_LPAREN),
	11399: uint16(233),
	11400: uint16(1),
	11401: uint16(sym_parameters),
	11402: uint16(3),
	11403: uint16(3),
	11404: uint16(1),
	11405: uint16(sym_comment),
	11406: uint16(112),
	11407: uint16(1),
	11408: uint16(anon_sym_LBRACE),
	11409: uint16(40),
	11410: uint16(1),
	11411: uint16(sym_block),
	11412: uint16(3),
	11413: uint16(3),
	11414: uint16(1),
	11415: uint16(sym_comment),
	11416: uint16(11),
	11417: uint16(1),
	11418: uint16(anon_sym_LBRACE),
	11419: uint16(210),
	11420: uint16(1),
	11421: uint16(sym_block),
	11422: uint16(3),
	11423: uint16(3),
	11424: uint16(1),
	11425: uint16(sym_comment),
	11426: uint16(11),
	11427: uint16(1),
	11428: uint16(anon_sym_LBRACE),
	11429: uint16(41),
	11430: uint16(1),
	11431: uint16(sym_block),
	11432: uint16(3),
	11433: uint16(3),
	11434: uint16(1),
	11435: uint16(sym_comment),
	11436: uint16(112),
	11437: uint16(1),
	11438: uint16(anon_sym_LBRACE),
	11439: uint16(41),
	11440: uint16(1),
	11441: uint16(sym_block),
	11442: uint16(3),
	11443: uint16(3),
	11444: uint16(1),
	11445: uint16(sym_comment),
	11446: uint16(112),
	11447: uint16(1),
	11448: uint16(anon_sym_LBRACE),
	11449: uint16(45),
	11450: uint16(1),
	11451: uint16(sym_block),
	11452: uint16(2),
	11453: uint16(3),
	11454: uint16(1),
	11455: uint16(sym_comment),
	11456: uint16(852),
	11457: uint16(2),
	11458: uint16(anon_sym_RBRACE),
	11459: uint16(anon_sym_COMMA),
	11460: uint16(3),
	11461: uint16(3),
	11462: uint16(1),
	11463: uint16(sym_comment),
	11464: uint16(112),
	11465: uint16(1),
	11466: uint16(anon_sym_LBRACE),
	11467: uint16(48),
	11468: uint16(1),
	11469: uint16(sym_block),
	11470: uint16(3),
	11471: uint16(3),
	11472: uint16(1),
	11473: uint16(sym_comment),
	11474: uint16(838),
	11475: uint16(1),
	11476: uint16(sym_identifier),
	11477: uint16(289),
	11478: uint16(1),
	11479: uint16(sym__type_identifier),
	11480: uint16(3),
	11481: uint16(3),
	11482: uint16(1),
	11483: uint16(sym_comment),
	11484: uint16(11),
	11485: uint16(1),
	11486: uint16(anon_sym_LBRACE),
	11487: uint16(126),
	11488: uint16(1),
	11489: uint16(sym_block),
	11490: uint16(3),
	11491: uint16(3),
	11492: uint16(1),
	11493: uint16(sym_comment),
	11494: uint16(11),
	11495: uint16(1),
	11496: uint16(anon_sym_LBRACE),
	11497: uint16(80),
	11498: uint16(1),
	11499: uint16(sym_block),
	11500: uint16(3),
	11501: uint16(3),
	11502: uint16(1),
	11503: uint16(sym_comment),
	11504: uint16(649),
	11505: uint16(1),
	11506: uint16(anon_sym_LBRACE),
	11507: uint16(89),
	11508: uint16(1),
	11509: uint16(sym_block),
	11510: uint16(3),
	11511: uint16(3),
	11512: uint16(1),
	11513: uint16(sym_comment),
	11514: uint16(900),
	11515: uint16(1),
	11516: uint16(sym_identifier),
	11517: uint16(343),
	11518: uint16(1),
	11519: uint16(sym__field_identifier),
	11520: uint16(3),
	11521: uint16(3),
	11522: uint16(1),
	11523: uint16(sym_comment),
	11524: uint16(649),
	11525: uint16(1),
	11526: uint16(anon_sym_LBRACE),
	11527: uint16(93),
	11528: uint16(1),
	11529: uint16(sym_block),
	11530: uint16(3),
	11531: uint16(3),
	11532: uint16(1),
	11533: uint16(sym_comment),
	11534: uint16(649),
	11535: uint16(1),
	11536: uint16(anon_sym_LBRACE),
	11537: uint16(87),
	11538: uint16(1),
	11539: uint16(sym_block),
	11540: uint16(2),
	11541: uint16(3),
	11542: uint16(1),
	11543: uint16(sym_comment),
	11544: uint16(902),
	11545: uint16(2),
	11546: uint16(anon_sym_LBRACE),
	11547: uint16(anon_sym_COLON),
	11548: uint16(2),
	11549: uint16(3),
	11550: uint16(1),
	11551: uint16(sym_comment),
	11552: uint16(904),
	11553: uint16(2),
	11554: uint16(anon_sym_LBRACE),
	11555: uint16(anon_sym_LT2),
	11556: uint16(3),
	11557: uint16(3),
	11558: uint16(1),
	11559: uint16(sym_comment),
	11560: uint16(11),
	11561: uint16(1),
	11562: uint16(anon_sym_LBRACE),
	11563: uint16(224),
	11564: uint16(1),
	11565: uint16(sym_block),
	11566: uint16(3),
	11567: uint16(3),
	11568: uint16(1),
	11569: uint16(sym_comment),
	11570: uint16(649),
	11571: uint16(1),
	11572: uint16(anon_sym_LBRACE),
	11573: uint16(99),
	11574: uint16(1),
	11575: uint16(sym_block),
	11576: uint16(3),
	11577: uint16(3),
	11578: uint16(1),
	11579: uint16(sym_comment),
	11580: uint16(838),
	11581: uint16(1),
	11582: uint16(sym_identifier),
	11583: uint16(272),
	11584: uint16(1),
	11585: uint16(sym__type_identifier),
	11586: uint16(3),
	11587: uint16(3),
	11588: uint16(1),
	11589: uint16(sym_comment),
	11590: uint16(776),
	11591: uint16(1),
	11592: uint16(anon_sym_LPAREN),
	11593: uint16(260),
	11594: uint16(1),
	11595: uint16(sym_parameters),
	11596: uint16(3),
	11597: uint16(3),
	11598: uint16(1),
	11599: uint16(sym_comment),
	11600: uint16(724),
	11601: uint16(1),
	11602: uint16(anon_sym_COLON_COLON),
	11603: uint16(906),
	11604: uint16(1),
	11605: uint16(anon_sym_SEMI),
	11606: uint16(2),
	11607: uint16(3),
	11608: uint16(1),
	11609: uint16(sym_comment),
	11610: uint16(420),
	11611: uint16(2),
	11612: uint16(anon_sym_COMMA),
	11613: uint16(anon_sym_GT),
	11614: uint16(2),
	11615: uint16(3),
	11616: uint16(1),
	11617: uint16(sym_comment),
	11618: uint16(908),
	11619: uint16(2),
	11620: uint16(anon_sym_LBRACE),
	11621: uint16(anon_sym_LT2),
	11622: uint16(2),
	11623: uint16(3),
	11624: uint16(1),
	11625: uint16(sym_comment),
	11626: uint16(910),
	11627: uint16(2),
	11628: uint16(anon_sym_LBRACE),
	11629: uint16(anon_sym_COLON),
	11630: uint16(3),
	11631: uint16(3),
	11632: uint16(1),
	11633: uint16(sym_comment),
	11634: uint16(11),
	11635: uint16(1),
	11636: uint16(anon_sym_LBRACE),
	11637: uint16(127),
	11638: uint16(1),
	11639: uint16(sym_block),
	11640: uint16(2),
	11641: uint16(3),
	11642: uint16(1),
	11643: uint16(sym_comment),
	11644: uint16(912),
	11645: uint16(2),
	11646: uint16(anon_sym_LBRACE),
	11647: uint16(anon_sym_COLON),
	11648: uint16(3),
	11649: uint16(3),
	11650: uint16(1),
	11651: uint16(sym_comment),
	11652: uint16(11),
	11653: uint16(1),
	11654: uint16(anon_sym_LBRACE),
	11655: uint16(130),
	11656: uint16(1),
	11657: uint16(sym_block),
	11658: uint16(3),
	11659: uint16(3),
	11660: uint16(1),
	11661: uint16(sym_comment),
	11662: uint16(914),
	11663: uint16(1),
	11664: uint16(sym_integer_literal),
	11665: uint16(916),
	11666: uint16(1),
	11667: uint16(sym_float_literal),
	11668: uint16(3),
	11669: uint16(3),
	11670: uint16(1),
	11671: uint16(sym_comment),
	11672: uint16(918),
	11673: uint16(1),
	11674: uint16(sym_identifier),
	11675: uint16(296),
	11676: uint16(1),
	11677: uint16(sym__type_identifier),
	11678: uint16(3),
	11679: uint16(3),
	11680: uint16(1),
	11681: uint16(sym_comment),
	11682: uint16(838),
	11683: uint16(1),
	11684: uint16(sym_identifier),
	11685: uint16(237),
	11686: uint16(1),
	11687: uint16(sym__type_identifier),
	11688: uint16(3),
	11689: uint16(3),
	11690: uint16(1),
	11691: uint16(sym_comment),
	11692: uint16(776),
	11693: uint16(1),
	11694: uint16(anon_sym_LPAREN),
	11695: uint16(240),
	11696: uint16(1),
	11697: uint16(sym_parameters),
	11698: uint16(3),
	11699: uint16(3),
	11700: uint16(1),
	11701: uint16(sym_comment),
	11702: uint16(11),
	11703: uint16(1),
	11704: uint16(anon_sym_LBRACE),
	11705: uint16(217),
	11706: uint16(1),
	11707: uint16(sym_block),
	11708: uint16(2),
	11709: uint16(3),
	11710: uint16(1),
	11711: uint16(sym_comment),
	11712: uint16(920),
	11713: uint16(1),
	11714: uint16(anon_sym_RPAREN),
	11715: uint16(2),
	11716: uint16(3),
	11717: uint16(1),
	11718: uint16(sym_comment),
	11719: uint16(922),
	11720: uint16(1),
	11721: uint16(sym_identifier),
	11722: uint16(2),
	11723: uint16(3),
	11724: uint16(1),
	11725: uint16(sym_comment),
	11726: uint16(924),
	11727: uint16(1),
	11728: uint16(anon_sym_LPAREN),
	11729: uint16(2),
	11730: uint16(3),
	11731: uint16(1),
	11732: uint16(sym_comment),
	11733: uint16(926),
	11734: uint16(1),
	11735: uint16(anon_sym_RBRACK),
	11736: uint16(2),
	11737: uint16(3),
	11738: uint16(1),
	11739: uint16(sym_comment),
	11740: uint16(928),
	11741: uint16(1),
	11742: uint16(anon_sym_LBRACK),
	11743: uint16(2),
	11744: uint16(3),
	11745: uint16(1),
	11746: uint16(sym_comment),
	11747: uint16(930),
	11748: uint16(1),
	11749: uint16(sym_identifier),
	11750: uint16(2),
	11751: uint16(3),
	11752: uint16(1),
	11753: uint16(sym_comment),
	11754: uint16(932),
	11755: uint16(1),
	11757: uint16(2),
	11758: uint16(3),
	11759: uint16(1),
	11760: uint16(sym_comment),
	11761: uint16(934),
	11762: uint16(1),
	11763: uint16(anon_sym_SEMI),
	11764: uint16(2),
	11765: uint16(3),
	11766: uint16(1),
	11767: uint16(sym_comment),
	11768: uint16(936),
	11769: uint16(1),
	11770: uint16(anon_sym_RBRACK),
	11771: uint16(2),
	11772: uint16(3),
	11773: uint16(1),
	11774: uint16(sym_comment),
	11775: uint16(938),
	11776: uint16(1),
	11777: uint16(anon_sym_RPAREN),
	11778: uint16(2),
	11779: uint16(3),
	11780: uint16(1),
	11781: uint16(sym_comment),
	11782: uint16(623),
	11783: uint16(1),
	11784: uint16(anon_sym_RPAREN),
	11785: uint16(2),
	11786: uint16(3),
	11787: uint16(1),
	11788: uint16(sym_comment),
	11789: uint16(940),
	11790: uint16(1),
	11791: uint16(sym_identifier),
	11792: uint16(2),
	11793: uint16(3),
	11794: uint16(1),
	11795: uint16(sym_comment),
	11796: uint16(942),
	11797: uint16(1),
	11798: uint16(anon_sym_COLON_COLON),
	11799: uint16(2),
	11800: uint16(3),
	11801: uint16(1),
	11802: uint16(sym_comment),
	11803: uint16(822),
	11804: uint16(1),
	11805: uint16(anon_sym_RPAREN),
	11806: uint16(2),
	11807: uint16(3),
	11808: uint16(1),
	11809: uint16(sym_comment),
	11810: uint16(944),
	11811: uint16(1),
	11812: uint16(anon_sym_RPAREN),
	11813: uint16(2),
	11814: uint16(3),
	11815: uint16(1),
	11816: uint16(sym_comment),
	11817: uint16(946),
	11818: uint16(1),
	11819: uint16(anon_sym_RPAREN),
	11820: uint16(2),
	11821: uint16(3),
	11822: uint16(1),
	11823: uint16(sym_comment),
	11824: uint16(948),
	11825: uint16(1),
	11826: uint16(anon_sym_COLON_COLON),
	11827: uint16(2),
	11828: uint16(3),
	11829: uint16(1),
	11830: uint16(sym_comment),
	11831: uint16(340),
	11832: uint16(1),
	11833: uint16(anon_sym_RBRACK),
	11834: uint16(2),
	11835: uint16(3),
	11836: uint16(1),
	11837: uint16(sym_comment),
	11838: uint16(950),
	11839: uint16(1),
	11840: uint16(anon_sym_LPAREN),
	11841: uint16(2),
	11842: uint16(3),
	11843: uint16(1),
	11844: uint16(sym_comment),
	11845: uint16(848),
	11846: uint16(1),
	11847: uint16(anon_sym_RBRACE),
	11848: uint16(2),
	11849: uint16(3),
	11850: uint16(1),
	11851: uint16(sym_comment),
	11852: uint16(782),
	11853: uint16(1),
	11854: uint16(anon_sym_RPAREN),
	11855: uint16(2),
	11856: uint16(3),
	11857: uint16(1),
	11858: uint16(sym_comment),
	11859: uint16(952),
	11860: uint16(1),
	11861: uint16(anon_sym_LPAREN),
	11862: uint16(2),
	11863: uint16(3),
	11864: uint16(1),
	11865: uint16(sym_comment),
	11866: uint16(954),
	11867: uint16(1),
	11868: uint16(anon_sym_COLON),
	11869: uint16(2),
	11870: uint16(3),
	11871: uint16(1),
	11872: uint16(sym_comment),
	11873: uint16(724),
	11874: uint16(1),
	11875: uint16(anon_sym_COLON_COLON),
	11876: uint16(2),
	11877: uint16(3),
	11878: uint16(1),
	11879: uint16(sym_comment),
	11880: uint16(956),
	11881: uint16(1),
	11882: uint16(anon_sym_SEMI),
	11883: uint16(2),
	11884: uint16(3),
	11885: uint16(1),
	11886: uint16(sym_comment),
	11887: uint16(888),
	11888: uint16(1),
	11889: uint16(anon_sym_GT),
	11890: uint16(2),
	11891: uint16(3),
	11892: uint16(1),
	11893: uint16(sym_comment),
	11894: uint16(958),
	11895: uint16(1),
	11896: uint16(anon_sym_COLON_COLON),
	11897: uint16(2),
	11898: uint16(3),
	11899: uint16(1),
	11900: uint16(sym_comment),
	11901: uint16(960),
	11902: uint16(1),
	11903: uint16(anon_sym_fun),
	11904: uint16(2),
	11905: uint16(3),
	11906: uint16(1),
	11907: uint16(sym_comment),
	11908: uint16(962),
	11909: uint16(1),
	11910: uint16(anon_sym_fun),
	11911: uint16(2),
	11912: uint16(3),
	11913: uint16(1),
	11914: uint16(sym_comment),
	11915: uint16(964),
	11916: uint16(1),
	11917: uint16(anon_sym_COLON_COLON),
	11918: uint16(2),
	11919: uint16(3),
	11920: uint16(1),
	11921: uint16(sym_comment),
	11922: uint16(966),
	11923: uint16(1),
	11924: uint16(sym_identifier),
	11925: uint16(2),
	11926: uint16(3),
	11927: uint16(1),
	11928: uint16(sym_comment),
	11929: uint16(968),
	11930: uint16(1),
	11931: uint16(anon_sym_LPAREN),
	11932: uint16(2),
	11933: uint16(3),
	11934: uint16(1),
	11935: uint16(sym_comment),
	11936: uint16(970),
	11937: uint16(1),
	11939: uint16(2),
	11940: uint16(3),
	11941: uint16(1),
	11942: uint16(sym_comment),
	11943: uint16(972),
	11944: uint16(1),
	11946: uint16(2),
	11947: uint16(3),
	11948: uint16(1),
	11949: uint16(sym_comment),
	11950: uint16(412),
	11951: uint16(1),
	11952: uint16(anon_sym_COLON_COLON),
	11953: uint16(2),
	11954: uint16(3),
	11955: uint16(1),
	11956: uint16(sym_comment),
	11957: uint16(974),
	11958: uint16(1),
	11959: uint16(anon_sym_GT),
	11960: uint16(2),
	11961: uint16(3),
	11962: uint16(1),
	11963: uint16(sym_comment),
	11964: uint16(976),
	11965: uint16(1),
	11966: uint16(anon_sym_LPAREN),
	11967: uint16(2),
	11968: uint16(3),
	11969: uint16(1),
	11970: uint16(sym_comment),
	11971: uint16(978),
	11972: uint16(1),
	11973: uint16(anon_sym_COLON),
	11974: uint16(2),
	11975: uint16(3),
	11976: uint16(1),
	11977: uint16(sym_comment),
	11978: uint16(336),
	11979: uint16(1),
	11980: uint16(anon_sym_RBRACK),
	11981: uint16(2),
	11982: uint16(3),
	11983: uint16(1),
	11984: uint16(sym_comment),
	11985: uint16(980),
	11986: uint16(1),
	11987: uint16(sym_identifier),
	11988: uint16(2),
	11989: uint16(3),
	11990: uint16(1),
	11991: uint16(sym_comment),
	11992: uint16(982),
	11993: uint16(1),
	11994: uint16(anon_sym_COLON_COLON),
	11995: uint16(2),
	11996: uint16(3),
	11997: uint16(1),
	11998: uint16(sym_comment),
	11999: uint16(984),
	12000: uint16(1),
	12001: uint16(anon_sym_COLON_COLON),
	12002: uint16(2),
	12003: uint16(3),
	12004: uint16(1),
	12005: uint16(sym_comment),
	12006: uint16(986),
	12007: uint16(1),
	12008: uint16(anon_sym_RPAREN),
	12009: uint16(2),
	12010: uint16(3),
	12011: uint16(1),
	12012: uint16(sym_comment),
	12013: uint16(988),
	12014: uint16(1),
	12015: uint16(anon_sym_RPAREN),
	12016: uint16(2),
	12017: uint16(3),
	12018: uint16(1),
	12019: uint16(sym_comment),
	12020: uint16(990),
	12021: uint16(1),
	12022: uint16(sym_identifier),
	12023: uint16(2),
	12024: uint16(3),
	12025: uint16(1),
	12026: uint16(sym_comment),
	12027: uint16(992),
	12028: uint16(1),
	12029: uint16(sym_identifier),
	12030: uint16(2),
	12031: uint16(3),
	12032: uint16(1),
	12033: uint16(sym_comment),
	12034: uint16(994),
	12035: uint16(1),
	12036: uint16(aux_sym_hex_address_token1),
	12037: uint16(2),
	12038: uint16(3),
	12039: uint16(1),
	12040: uint16(sym_comment),
	12041: uint16(996),
	12042: uint16(1),
	12044: uint16(2),
	12045: uint16(3),
	12046: uint16(1),
	12047: uint16(sym_comment),
	12048: uint16(998),
	12049: uint16(1),
	12050: uint16(anon_sym_fun),
	12051: uint16(2),
	12052: uint16(3),
	12053: uint16(1),
	12054: uint16(sym_comment),
	12055: uint16(1000),
	12056: uint16(1),
	12057: uint16(anon_sym_LPAREN),
	12058: uint16(2),
	12059: uint16(3),
	12060: uint16(1),
	12061: uint16(sym_comment),
	12062: uint16(1002),
	12063: uint16(1),
	12064: uint16(anon_sym_LPAREN),
}

var ts_small_parse_table_map = [354]uint32_t{
	1:   uint32(125),
	2:   uint32(250),
	3:   uint32(375),
	4:   uint32(500),
	5:   uint32(625),
	6:   uint32(750),
	7:   uint32(875),
	8:   uint32(985),
	9:   uint32(1092),
	10:  uint32(1199),
	11:  uint32(1306),
	12:  uint32(1413),
	13:  uint32(1520),
	14:  uint32(1626),
	15:  uint32(1732),
	16:  uint32(1838),
	17:  uint32(1944),
	18:  uint32(2050),
	19:  uint32(2156),
	20:  uint32(2259),
	21:  uint32(2362),
	22:  uint32(2465),
	23:  uint32(2565),
	24:  uint32(2665),
	25:  uint32(2765),
	26:  uint32(2865),
	27:  uint32(2965),
	28:  uint32(3065),
	29:  uint32(3165),
	30:  uint32(3265),
	31:  uint32(3365),
	32:  uint32(3465),
	33:  uint32(3565),
	34:  uint32(3665),
	35:  uint32(3765),
	36:  uint32(3865),
	37:  uint32(3965),
	38:  uint32(4026),
	39:  uint32(4083),
	40:  uint32(4139),
	41:  uint32(4195),
	42:  uint32(4254),
	43:  uint32(4309),
	44:  uint32(4364),
	45:  uint32(4419),
	46:  uint32(4474),
	47:  uint32(4540),
	48:  uint32(4604),
	49:  uint32(4668),
	50:  uint32(4732),
	51:  uint32(4796),
	52:  uint32(4860),
	53:  uint32(4924),
	54:  uint32(4988),
	55:  uint32(5052),
	56:  uint32(5116),
	57:  uint32(5180),
	58:  uint32(5244),
	59:  uint32(5308),
	60:  uint32(5357),
	61:  uint32(5406),
	62:  uint32(5455),
	63:  uint32(5533),
	64:  uint32(5608),
	65:  uint32(5683),
	66:  uint32(5755),
	67:  uint32(5827),
	68:  uint32(5899),
	69:  uint32(5971),
	70:  uint32(6043),
	71:  uint32(6115),
	72:  uint32(6187),
	73:  uint32(6259),
	74:  uint32(6298),
	75:  uint32(6337),
	76:  uint32(6374),
	77:  uint32(6411),
	78:  uint32(6451),
	79:  uint32(6486),
	80:  uint32(6523),
	81:  uint32(6562),
	82:  uint32(6597),
	83:  uint32(6634),
	84:  uint32(6669),
	85:  uint32(6704),
	86:  uint32(6739),
	87:  uint32(6774),
	88:  uint32(6811),
	89:  uint32(6846),
	90:  uint32(6881),
	91:  uint32(6916),
	92:  uint32(6951),
	93:  uint32(6986),
	94:  uint32(7021),
	95:  uint32(7056),
	96:  uint32(7091),
	97:  uint32(7126),
	98:  uint32(7161),
	99:  uint32(7195),
	100: uint32(7229),
	101: uint32(7263),
	102: uint32(7319),
	103: uint32(7355),
	104: uint32(7391),
	105: uint32(7451),
	106: uint32(7497),
	107: uint32(7531),
	108: uint32(7565),
	109: uint32(7599),
	110: uint32(7633),
	111: uint32(7685),
	112: uint32(7719),
	113: uint32(7777),
	114: uint32(7811),
	115: uint32(7871),
	116: uint32(7905),
	117: uint32(7939),
	118: uint32(7975),
	119: uint32(8009),
	120: uint32(8043),
	121: uint32(8077),
	122: uint32(8111),
	123: uint32(8145),
	124: uint32(8179),
	125: uint32(8213),
	126: uint32(8247),
	127: uint32(8281),
	128: uint32(8315),
	129: uint32(8349),
	130: uint32(8387),
	131: uint32(8429),
	132: uint32(8473),
	133: uint32(8521),
	134: uint32(8571),
	135: uint32(8605),
	136: uint32(8639),
	137: uint32(8677),
	138: uint32(8709),
	139: uint32(8771),
	140: uint32(8803),
	141: uint32(8862),
	142: uint32(8921),
	143: uint32(8980),
	144: uint32(9037),
	145: uint32(9096),
	146: uint32(9152),
	147: uint32(9208),
	148: uint32(9264),
	149: uint32(9320),
	150: uint32(9376),
	151: uint32(9432),
	152: uint32(9477),
	153: uint32(9522),
	154: uint32(9564),
	155: uint32(9606),
	156: uint32(9642),
	157: uint32(9675),
	158: uint32(9708),
	159: uint32(9738),
	160: uint32(9767),
	161: uint32(9793),
	162: uint32(9821),
	163: uint32(9849),
	164: uint32(9875),
	165: uint32(9901),
	166: uint32(9927),
	167: uint32(9953),
	168: uint32(9979),
	169: uint32(10005),
	170: uint32(10031),
	171: uint32(10057),
	172: uint32(10083),
	173: uint32(10109),
	174: uint32(10135),
	175: uint32(10161),
	176: uint32(10189),
	177: uint32(10214),
	178: uint32(10239),
	179: uint32(10264),
	180: uint32(10282),
	181: uint32(10298),
	182: uint32(10314),
	183: uint32(10339),
	184: uint32(10352),
	185: uint32(10365),
	186: uint32(10387),
	187: uint32(10399),
	188: uint32(10413),
	189: uint32(10424),
	190: uint32(10435),
	191: uint32(10446),
	192: uint32(10465),
	193: uint32(10476),
	194: uint32(10487),
	195: uint32(10498),
	196: uint32(10509),
	197: uint32(10528),
	198: uint32(10539),
	199: uint32(10550),
	200: uint32(10561),
	201: uint32(10572),
	202: uint32(10583),
	203: uint32(10594),
	204: uint32(10605),
	205: uint32(10616),
	206: uint32(10627),
	207: uint32(10637),
	208: uint32(10647),
	209: uint32(10663),
	210: uint32(10679),
	211: uint32(10693),
	212: uint32(10709),
	213: uint32(10719),
	214: uint32(10735),
	215: uint32(10745),
	216: uint32(10755),
	217: uint32(10765),
	218: uint32(10781),
	219: uint32(10797),
	220: uint32(10807),
	221: uint32(10821),
	222: uint32(10831),
	223: uint32(10841),
	224: uint32(10852),
	225: uint32(10865),
	226: uint32(10878),
	227: uint32(10891),
	228: uint32(10904),
	229: uint32(10917),
	230: uint32(10930),
	231: uint32(10943),
	232: uint32(10956),
	233: uint32(10969),
	234: uint32(10982),
	235: uint32(10995),
	236: uint32(11008),
	237: uint32(11019),
	238: uint32(11032),
	239: uint32(11045),
	240: uint32(11058),
	241: uint32(11071),
	242: uint32(11084),
	243: uint32(11097),
	244: uint32(11110),
	245: uint32(11121),
	246: uint32(11132),
	247: uint32(11145),
	248: uint32(11158),
	249: uint32(11171),
	250: uint32(11182),
	251: uint32(11195),
	252: uint32(11206),
	253: uint32(11219),
	254: uint32(11232),
	255: uint32(11245),
	256: uint32(11258),
	257: uint32(11269),
	258: uint32(11282),
	259: uint32(11292),
	260: uint32(11302),
	261: uint32(11310),
	262: uint32(11320),
	263: uint32(11330),
	264: uint32(11340),
	265: uint32(11350),
	266: uint32(11358),
	267: uint32(11366),
	268: uint32(11374),
	269: uint32(11384),
	270: uint32(11392),
	271: uint32(11402),
	272: uint32(11412),
	273: uint32(11422),
	274: uint32(11432),
	275: uint32(11442),
	276: uint32(11452),
	277: uint32(11460),
	278: uint32(11470),
	279: uint32(11480),
	280: uint32(11490),
	281: uint32(11500),
	282: uint32(11510),
	283: uint32(11520),
	284: uint32(11530),
	285: uint32(11540),
	286: uint32(11548),
	287: uint32(11556),
	288: uint32(11566),
	289: uint32(11576),
	290: uint32(11586),
	291: uint32(11596),
	292: uint32(11606),
	293: uint32(11614),
	294: uint32(11622),
	295: uint32(11630),
	296: uint32(11640),
	297: uint32(11648),
	298: uint32(11658),
	299: uint32(11668),
	300: uint32(11678),
	301: uint32(11688),
	302: uint32(11698),
	303: uint32(11708),
	304: uint32(11715),
	305: uint32(11722),
	306: uint32(11729),
	307: uint32(11736),
	308: uint32(11743),
	309: uint32(11750),
	310: uint32(11757),
	311: uint32(11764),
	312: uint32(11771),
	313: uint32(11778),
	314: uint32(11785),
	315: uint32(11792),
	316: uint32(11799),
	317: uint32(11806),
	318: uint32(11813),
	319: uint32(11820),
	320: uint32(11827),
	321: uint32(11834),
	322: uint32(11841),
	323: uint32(11848),
	324: uint32(11855),
	325: uint32(11862),
	326: uint32(11869),
	327: uint32(11876),
	328: uint32(11883),
	329: uint32(11890),
	330: uint32(11897),
	331: uint32(11904),
	332: uint32(11911),
	333: uint32(11918),
	334: uint32(11925),
	335: uint32(11932),
	336: uint32(11939),
	337: uint32(11946),
	338: uint32(11953),
	339: uint32(11960),
	340: uint32(11967),
	341: uint32(11974),
	342: uint32(11981),
	343: uint32(11988),
	344: uint32(11995),
	345: uint32(12002),
	346: uint32(12009),
	347: uint32(12016),
	348: uint32(12023),
	349: uint32(12030),
	350: uint32(12037),
	351: uint32(12044),
	352: uint32(12051),
	353: uint32(12058),
}

var ts_parse_actions = [1004]TSParseActionEntry{
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
		Fstate: uint16(252),
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
		Fstate: uint16(85),
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
		Fstate: uint16(307),
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
		Fstate: uint16(4),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_return_expression),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_return_expression),
	})))),
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
		Fstate: uint16(162),
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
		Fstate: uint16(10),
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
		Fstate: uint16(90),
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
		Fstate: uint16(22),
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
		Fstate: uint16(36),
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
		Fstate: uint16(2),
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
		Fstate: uint16(355),
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
		Fstate: uint16(356),
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
		Fstate: uint16(300),
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
		Fstate: uint16(298),
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
		Fstate: uint16(122),
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
		Fstate: uint16(121),
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
		Fstate: uint16(83),
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
		Fstate: uint16(83),
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
		Fstate: uint16(84),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
	})))),
	49: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(85),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	50: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	51: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(307),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	54: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(5),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(334),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(162),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(10),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(178),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(90),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(22),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(36),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(2),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(327),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(337),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(278),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(277),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(122),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(121),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(83),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(83),
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
		Fsymbol:      uint16(aux_sym_block_repeat1),
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
		Fstate:      uint16(84),
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
		Fstate: uint16(5),
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
		Fstate: uint16(77),
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
		Fcount: uint8(1),
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
		Fstate: uint16(350),
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
		Fcount: uint8(1),
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
		Fstate: uint16(334),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(162),
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
		Fstate: uint16(178),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(36),
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
		Fstate: uint16(327),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(337),
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
		Fcount: uint8(1),
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
		Fstate: uint16(278),
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
		Fstate: uint16(277),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(42),
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
		Fstate: uint16(43),
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
		Fstate: uint16(78),
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
		Fstate: uint16(79),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(76),
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
		Fstate: uint16(310),
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
		Fstate: uint16(103),
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
		Fcount: uint8(1),
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
		Fstate: uint16(69),
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
		Fstate: uint16(102),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(131),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
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
		Fstate:      uint16(85),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
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
		Fstate:      uint16(307),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
	})))),
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
		Fstate:      uint16(4),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
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
		Fstate:      uint16(162),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
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
		Fstate:      uint16(10),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
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
		Fstate:      uint16(90),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
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
		Fstate:      uint16(22),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
	})))),
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
		Fstate:      uint16(36),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	185: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
	})))),
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
		Fstate:      uint16(2),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
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
		Fstate:      uint16(355),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
	})))),
	192: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(356),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
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
		Fstate:      uint16(300),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
	})))),
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
		Fstate:      uint16(298),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	200: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
	})))),
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
		Fstate:      uint16(122),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
	})))),
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
		Fstate:      uint16(121),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	206: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
	})))),
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
		Fstate:      uint16(83),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
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
		Fstate:      uint16(83),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
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
		Fstate:      uint16(84),
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
		Fstate: uint16(316),
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
		Fstate: uint16(119),
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
		Fcount: uint8(1),
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
		Fstate: uint16(26),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(109),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_expression),
		Fproduction_id: uint16(29),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_expression),
		Fproduction_id: uint16(29),
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
		Fcount: uint8(1),
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
		Fstate: uint16(223),
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
		Fcount: uint8(1),
	}})),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_const_block),
		Fproduction_id: uint16(13),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_const_block),
		Fproduction_id: uint16(13),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_block),
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
		Fsymbol:      uint16(sym_block),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_block),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_block),
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
		Fcount: uint8(1),
	}})),
	243: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_expression_statement),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_expression_statement),
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
		Fsymbol:      uint16(sym__expression_except_range),
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
		Fcount: uint8(1),
	}})),
	249: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__expression_except_range),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_loop_expression),
		Fproduction_id: uint16(13),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_loop_expression),
		Fproduction_id: uint16(13),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_else_clause),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_else_clause),
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
		Fcount: uint8(1),
	}})),
	259: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_if_expression),
		Fproduction_id: uint16(34),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_if_expression),
		Fproduction_id: uint16(34),
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
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_while_expression),
		Fproduction_id: uint16(30),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_while_expression),
		Fproduction_id: uint16(30),
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
		Fsymbol:      uint16(aux_sym_delim_token_tree_repeat1),
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
		Fstate:      uint16(51),
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
		Fsymbol:      uint16(aux_sym_delim_token_tree_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_delim_token_tree_repeat1),
	})))),
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
		Fstate:      uint16(61),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_delim_token_tree_repeat1),
	})))),
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
		Fstate:      uint16(60),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	278: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_delim_token_tree_repeat1),
	})))),
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
		Fstate:      uint16(49),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	281: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_delim_token_tree_repeat1),
	})))),
	282: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(62),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_delim_token_tree_repeat1),
	})))),
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
		Fstate:      uint16(49),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(51),
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
		Fstate: uint16(344),
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
		Fstate: uint16(61),
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
		Fstate: uint16(60),
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
		Fstate: uint16(49),
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
		Fstate: uint16(62),
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
		Fstate: uint16(49),
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
		Fstate: uint16(64),
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
		Fstate: uint16(56),
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
		Fcount: uint8(1),
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
		Fstate: uint16(56),
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
		Fstate: uint16(63),
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
		Fstate: uint16(323),
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
		Fstate: uint16(50),
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
		Fstate: uint16(50),
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
		Fstate: uint16(52),
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
		Fstate: uint16(52),
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
		Fstate: uint16(53),
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
		Fstate: uint16(53),
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
		Fstate: uint16(59),
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
		Fcount: uint8(1),
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
		Fstate: uint16(59),
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
		Fstate: uint16(54),
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
		Fcount: uint8(1),
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
		Fstate: uint16(54),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__non_special_token),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	335: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__non_special_token),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	337: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_delim_token_tree),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_delim_token_tree),
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
		Fsymbol:      uint16(sym_delim_token_tree),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_delim_token_tree),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(182),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(326),
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
		Fstate: uint16(65),
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
		Fstate: uint16(221),
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
		Fstate: uint16(207),
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
		Fstate: uint16(190),
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
		Fstate: uint16(73),
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
		Fstate: uint16(74),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(301),
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
		Fcount: uint8(1),
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
		Fstate: uint16(276),
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
		Fstate: uint16(221),
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
		Fstate: uint16(254),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(201),
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
		Fstate: uint16(254),
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
		Fstate: uint16(196),
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
		Fcount: uint8(1),
	}})),
	377: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(226),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(226),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(271),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(271),
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
		Fcount: uint8(1),
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
		Fstate: uint16(204),
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
		Fstate: uint16(204),
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
		Fstate: uint16(247),
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
		Fstate: uint16(247),
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
		Fstate: uint16(202),
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
		Fstate: uint16(75),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(198),
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
		Fstate: uint16(206),
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
		Fstate: uint16(206),
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
		Fstate: uint16(213),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(5),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(5),
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
		Fstate: uint16(264),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__expression),
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
		Fcount: uint8(1),
	}})),
	417: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__expression),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_type_arguments),
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
		Fcount: uint8(1),
	}})),
	427: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_type_arguments),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(17),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(17),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_expression_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_expression_statement),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(26),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(26),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	441: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__path),
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
		Fcount: uint8(1),
	}})),
	443: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_use_declaration),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_use_declaration),
		Fproduction_id: uint16(4),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_type_arguments),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_type_arguments),
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
		Fcount: uint8(1),
	}})),
	451: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(25),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(25),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_field_expression),
		Fproduction_id: uint16(23),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_field_expression),
		Fproduction_id: uint16(23),
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
		Fcount: uint8(1),
	}})),
	459: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(8),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(8),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__field_identifier),
		Fproduction_id: uint16(24),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__field_identifier),
		Fproduction_id: uint16(24),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(10),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(10),
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
		Fsymbol:      uint16(sym_type_arguments),
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
		Fsymbol:      uint16(sym_type_arguments),
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
		Fcount: uint8(1),
	}})),
	475: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(15),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	477: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(15),
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
		Fcount: uint8(1),
	}})),
	479: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(16),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_function_item),
		Fproduction_id: uint16(16),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_type_cast_expression),
		Fproduction_id: uint16(22),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_type_cast_expression),
		Fproduction_id: uint16(22),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_tuple_expression),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_tuple_expression),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unit_expression),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unit_expression),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(21),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(21),
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
		Fstate: uint16(30),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(169),
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
		Fcount: uint8(1),
	}})),
	503: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(35),
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
		Fstate: uint16(30),
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
		Fstate: uint16(37),
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
	513: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(39),
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
		Fcount: uint8(1),
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
		Fstate: uint16(39),
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
		Fstate: uint16(253),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_reference_expression),
		Fproduction_id: uint16(12),
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
		Fcount: uint8(1),
	}})),
	523: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_reference_expression),
		Fproduction_id: uint16(12),
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
		Fsymbol:      uint16(sym_unary_expression),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_unary_expression),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_return_expression),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_return_expression),
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
		Fstate: uint16(25),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_arguments),
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
		Fcount: uint8(1),
	}})),
	539: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_arguments),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	541: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_reference_type),
		Fproduction_id: uint16(6),
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
		Fcount: uint8(1),
	}})),
	543: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_reference_type),
		Fproduction_id: uint16(6),
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
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_tuple_expression),
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
		Fcount: uint8(1),
	}})),
	547: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_tuple_expression),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_tuple_expression),
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
		Fcount: uint8(1),
	}})),
	551: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_tuple_expression),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_call_expression),
		Fproduction_id: uint16(14),
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
		Fcount: uint8(1),
	}})),
	555: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_call_expression),
		Fproduction_id: uint16(14),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_assignment_expression),
		Fproduction_id: uint16(20),
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
		Fcount: uint8(1),
	}})),
	559: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_assignment_expression),
		Fproduction_id: uint16(20),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_arguments),
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
		Fcount: uint8(1),
	}})),
	563: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_arguments),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_arguments),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_arguments),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	569: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_reference_expression),
		Fproduction_id: uint16(18),
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
		Fcount: uint8(1),
	}})),
	571: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_reference_expression),
		Fproduction_id: uint16(18),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	573: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_continue_expression),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_continue_expression),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_break_expression),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_break_expression),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_arguments),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_generic_function),
		Fproduction_id: uint16(19),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_generic_function),
		Fproduction_id: uint16(19),
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
		Fsymbol:        uint16(sym__type),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	591: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__type),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unit_type),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unit_type),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_tuple_expression),
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
		Fcount: uint8(1),
	}})),
	599: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_tuple_expression),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__type_identifier),
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
		Fcount: uint8(1),
	}})),
	603: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__type_identifier),
		Fproduction_id: uint16(3),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_parenthesized_expression),
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
		Fcount: uint8(1),
	}})),
	607: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_parenthesized_expression),
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
		Fcount: uint8(1),
	}})),
	609: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	613: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat1),
	})))),
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
		Fstate:      uint16(310),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_attribute_item),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_attribute_item),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(29),
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
		Fstate: uint16(23),
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
		Fstate: uint16(123),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tuple_expression_repeat2),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(17),
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
		Fstate: uint16(138),
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
		Fstate: uint16(111),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_let_condition),
		Fproduction_id: uint16(33),
	})))),
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
		Fstate: uint16(18),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(18),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__condition),
	})))),
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
		Fstate: uint16(88),
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
		Fcount: uint8(1),
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
		Fstate: uint16(189),
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
		Fstate: uint16(8),
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
		Fstate: uint16(86),
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
		Fstate: uint16(306),
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
		Fcount: uint8(1),
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
		Fstate: uint16(203),
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
		Fstate: uint16(177),
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
		Fcount: uint8(1),
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
		Fstate: uint16(263),
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
		Fstate: uint16(263),
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
		Fcount: uint8(1),
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
		Fstate: uint16(295),
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
		Fstate: uint16(92),
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
		Fcount: uint8(1),
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
		Fstate: uint16(228),
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
		Fstate: uint16(228),
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
		Fcount: uint8(1),
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
		Fstate: uint16(239),
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
		Fstate: uint16(319),
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
		Fcount: uint8(1),
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
		Fstate: uint16(238),
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
		Fstate: uint16(299),
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
		Fcount: uint8(1),
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
		Fstate: uint16(270),
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
		Fstate: uint16(269),
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
		Fstate: uint16(297),
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
		Fcount: uint8(1),
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
		Fstate: uint16(294),
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
		Fcount: uint8(1),
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
		Fstate: uint16(185),
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
		Fcount: uint8(1),
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
		Fstate: uint16(137),
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
		Fstate: uint16(315),
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
		Fcount: uint8(1),
	}})),
	694: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(170),
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
		Fstate: uint16(312),
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
		Fstate: uint16(336),
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
		Fstate: uint16(334),
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
		Fstate: uint16(164),
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
		Fstate: uint16(338),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_body_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	710: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_module_body_repeat1),
	})))),
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
		Fstate:      uint16(336),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_module_body_repeat1),
	})))),
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
		Fstate:      uint16(334),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_module_body_repeat1),
	})))),
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
		Fstate:      uint16(164),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__pattern),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	721: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_scoped_identifier),
		Fproduction_id: uint16(1),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_scoped_identifier),
		Fproduction_id: uint16(7),
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
		Fstate: uint16(311),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_attribute),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(58),
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
		Fcount: uint8(1),
	}})),
	737: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(191),
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
		Fstate: uint16(325),
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
		Fcount: uint8(1),
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
		Fstate: uint16(285),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_struct_pattern),
		Fproduction_id: uint16(31),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_struct_pattern),
		Fproduction_id: uint16(31),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	749: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_negative_literal),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(192),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_struct_pattern),
		Fproduction_id: uint16(31),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_tuple_pattern),
	})))),
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
	757: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_tuple_pattern),
	})))),
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
		Fsymbol:      uint16(sym_mut_pattern),
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
		Fstate: uint16(195),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_struct_pattern),
		Fproduction_id: uint16(31),
	})))),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_tuple_pattern),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_reference_pattern),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_or_pattern),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_reference_pattern),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_tuple_pattern),
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
		Fstate: uint16(303),
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
		Fstate: uint16(158),
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
		Fstate: uint16(355),
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
		Fstate: uint16(66),
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
		Fstate: uint16(197),
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
		Fstate: uint16(70),
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
		Fstate: uint16(327),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_field_pattern),
		Fproduction_id: uint16(36),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	791: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_parameters_repeat1),
	})))),
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
		Fstate:      uint16(161),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_parameters_repeat1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(155),
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
		Fstate: uint16(98),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(154),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	804: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(27),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(167),
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
		Fstate: uint16(168),
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
		Fstate: uint16(175),
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
		Fstate: uint16(173),
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
		Fstate: uint16(24),
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
		Fstate: uint16(245),
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
		Fstate: uint16(308),
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
		Fstate: uint16(159),
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
		Fstate: uint16(288),
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
		Fstate: uint16(176),
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
		Fstate: uint16(171),
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
		Fstate: uint16(67),
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
		Fstate: uint16(174),
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
		Fstate: uint16(251),
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
		Fstate: uint16(342),
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
		Fstate: uint16(160),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(189),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(188),
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
		Fstate: uint16(156),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_field_pattern),
		Fproduction_id: uint16(37),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_field_pattern),
		Fproduction_id: uint16(32),
	})))),
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
		Fstate: uint16(200),
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
		Fstate: uint16(199),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_struct_pattern_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	855: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_struct_pattern_repeat1),
	})))),
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
		Fstate:      uint16(220),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(324),
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
		Fstate: uint16(352),
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
		Fstate: uint16(96),
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
		Fstate: uint16(94),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tuple_pattern_repeat1),
	})))),
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
		Fstate: uint16(194),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_type_arguments_repeat1),
	})))),
	871: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(157),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	872: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	873: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_type_arguments_repeat1),
	})))),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_tuple_pattern_repeat1),
	})))),
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
		Fstate:      uint16(72),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	878: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_type_parameters_repeat1),
	})))),
	879: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(292),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	880: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	881: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_type_parameters_repeat1),
	})))),
	882: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	883: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_field_pattern),
		Fproduction_id: uint16(35),
	})))),
	884: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	885: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	886: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	887: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(184),
	}})))),
	888: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	889: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(322),
	}})))),
	890: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	891: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	892: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	893: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	894: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	895: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_parameter),
		Fproduction_id: uint16(11),
	})))),
	896: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	897: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_parameters),
	})))),
	898: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	899: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	900: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	901: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(259),
	}})))),
	902: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	903: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_parameters),
	})))),
	904: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	905: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_scoped_type_identifier),
		Fproduction_id: uint16(7),
	})))),
	906: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	907: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__use_clause),
	})))),
	908: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	909: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_scoped_type_identifier),
		Fproduction_id: uint16(1),
	})))),
	910: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	911: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_parameters),
	})))),
	912: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	913: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_parameters),
	})))),
	914: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	915: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	916: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	917: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	918: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	919: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	920: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	921: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	922: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	923: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(186),
	}})))),
	924: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	925: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_type_parameters),
	})))),
	926: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	927: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	928: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	929: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(165),
	}})))),
	930: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	931: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	932: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	933: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_module_body),
	})))),
	934: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	935: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	936: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	937: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_attribute),
		Fproduction_id: uint16(27),
	})))),
	938: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	939: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	940: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	941: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	942: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	943: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(261),
	}})))),
	944: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	945: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	946: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	947: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	948: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	949: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_bracketed_type),
	})))),
	950: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	951: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_type_parameters),
	})))),
	952: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	953: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	954: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	955: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(68),
	}})))),
	956: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	957: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(215),
	}})))),
	958: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	959: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_generic_type),
		Fproduction_id: uint16(28),
	})))),
	960: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	961: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(317),
	}})))),
	962: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	963: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_visibility_modifier),
	})))),
	964: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	965: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	966: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	967: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	968: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	969: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	970: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	971: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_module_body),
	})))),
	972: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	973: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	974: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	975: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_qualified_type),
		Fproduction_id: uint16(9),
	})))),
	976: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	977: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_type_parameters),
	})))),
	978: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	979: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	980: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	981: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(267),
	}})))),
	982: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	983: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_hex_address),
	})))),
	984: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	985: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(345),
	}})))),
	986: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	987: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	988: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	989: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	990: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	991: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	992: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	993: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	994: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	995: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	996: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	997: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	998: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	999: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(351),
	}})))),
	1000: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1001: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1002: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1003: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

func tree_sitter_move(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fversion:                   uint32(LANGUAGE_VERSION),
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

var __ccgo_ts1 = "end\x00identifier\x00module\x00::\x00{\x00}\x00;\x00#\x00[\x00]\x00=\x00fun\x00:\x00public\x00<\x00,\x00>\x00(\x00_\x00)\x00use\x00as\x00&\x00mutable_specifier\x00!\x00&&\x00||\x00|\x00^\x00==\x00!=\x00<=\x00>=\x00<<\x00>>\x00+\x00-\x00*\x00/\x00%\x00return\x00if\x00let\x00else\x00while\x00loop\x00const\x00break\x00continue\x00.\x00$\x00_non_special_token_token1\x00'\x00async\x00await\x00default\x00enum\x00fn\x00for\x00impl\x00match\x00mod\x00pub\x00static\x00struct\x00trait\x00type\x00union\x00unsafe\x00where\x00integer_literal\x00float_literal\x00true\x00false\x000x\x000X\x00hex_address_token1\x00comment\x00module_body\x00_statement\x00expression_statement\x00_declaration_statement\x00attribute_item\x00attribute\x00function_item\x00visibility_modifier\x00type_parameters\x00parameters\x00parameter\x00use_declaration\x00_use_clause\x00_type\x00bracketed_type\x00qualified_type\x00unit_type\x00generic_function\x00generic_type\x00type_arguments\x00reference_type\x00_expression_except_range\x00_expression\x00_expression_ending_with_block\x00scoped_type_identifier\x00unary_expression\x00reference_expression\x00binary_expression\x00assignment_expression\x00type_cast_expression\x00return_expression\x00call_expression\x00arguments\x00parenthesized_expression\x00tuple_expression\x00unit_expression\x00if_expression\x00let_condition\x00_condition\x00else_clause\x00while_expression\x00loop_expression\x00const_block\x00break_expression\x00continue_expression\x00field_expression\x00block\x00_pattern\x00tuple_pattern\x00struct_pattern\x00field_pattern\x00mut_pattern\x00reference_pattern\x00or_pattern\x00token_tree\x00_delim_tokens\x00_non_delim_token\x00scoped_identifier\x00_non_special_token\x00_literal\x00_literal_pattern\x00negative_literal\x00boolean_literal\x00hex_address\x00_path\x00_type_identifier\x00_field_identifier\x00module_body_repeat1\x00type_parameters_repeat1\x00parameters_repeat1\x00type_arguments_repeat1\x00arguments_repeat1\x00tuple_expression_repeat1\x00tuple_expression_repeat2\x00block_repeat1\x00tuple_pattern_repeat1\x00struct_pattern_repeat1\x00delim_token_tree_repeat1\x00field_identifier\x00primitive_type\x00shorthand_field_identifier\x00type_identifier\x00alias\x00alternative\x00argument\x00body\x00condition\x00consequence\x00field\x00function\x00left\x00name\x00operator\x00path\x00pattern\x00return_type\x00right\x00value\x00"
