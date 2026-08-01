// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-godot-resource/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-godot-resource -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_godot_resource

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
const EXTERNAL_TOKEN_COUNT = 1
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
const LARGE_STATE_COUNT = 2
const MAX_ALIAS_SEQUENCE_LENGTH = 5
const PRODUCTION_ID_COUNT = 2
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 93
const SYMBOL_COUNT = 40
const TOKEN_COUNT = 20
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

type wint_t = uint32

type wctype_t = uint64

type locale_t = uintptr

type wctrans_t = uintptr

// using wctype::iswspace

type TokenType = int32

const STRING = 0

func tree_sitter_godot_resource_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var last_char uint32_t
	_ = last_char
	if !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(STRING))) != 0) {
		return libc.BoolUint8(false1 != 0)
	}
	for libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
	}
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('"') {
		return libc.BoolUint8(false1 != 0)
	}
	last_char = uint32('"')
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		if last_char != uint32('\\') && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('"') {
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(STRING)
			return libc.BoolUint8(true1 != 0)
		}
		last_char = libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	}
	return libc.BoolUint8(false1 != 0)
}

func tree_sitter_godot_resource_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_godot_resource_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_godot_resource_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
}

func tree_sitter_godot_resource_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
}

type ts_symbol_identifiers = int32

const sym__identifier = 1
const sym_comment = 2
const sym_true = 3
const sym_false = 4
const sym_null = 5
const sym_float = 6
const sym_integer = 7
const anon_sym_AMP = 8
const anon_sym_LBRACK = 9
const anon_sym_RBRACK = 10
const anon_sym_EQ = 11
const sym_path = 12
const anon_sym_COLON = 13
const anon_sym_LBRACE = 14
const anon_sym_COMMA = 15
const anon_sym_RBRACE = 16
const anon_sym_LPAREN = 17
const anon_sym_RPAREN = 18
const sym_string = 19
const sym_resource = 20
const sym_identifier = 21
const sym_string_name = 22
const sym__value = 23
const sym_section = 24
const aux_sym__attributes = 25
const sym_attribute = 26
const aux_sym__properties = 27
const sym_property = 28
const sym_pair = 29
const sym_dictionary = 30
const sym_array = 31
const sym_arguments = 32
const sym__type_args = 33
const sym_constructor = 34
const aux_sym_resource_repeat1 = 35
const aux_sym_resource_repeat2 = 36
const aux_sym_dictionary_repeat1 = 37
const aux_sym_array_repeat1 = 38
const aux_sym_arguments_repeat1 = 39
const anon_alias_sym_string = 40

var ts_symbol_names = [41]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 16,
	3:  __ccgo_ts + 24,
	4:  __ccgo_ts + 29,
	5:  __ccgo_ts + 35,
	6:  __ccgo_ts + 40,
	7:  __ccgo_ts + 46,
	8:  __ccgo_ts + 54,
	9:  __ccgo_ts + 56,
	10: __ccgo_ts + 58,
	11: __ccgo_ts + 60,
	12: __ccgo_ts + 62,
	13: __ccgo_ts + 67,
	14: __ccgo_ts + 69,
	15: __ccgo_ts + 71,
	16: __ccgo_ts + 73,
	17: __ccgo_ts + 75,
	18: __ccgo_ts + 77,
	19: __ccgo_ts + 79,
	20: __ccgo_ts + 86,
	21: __ccgo_ts + 95,
	22: __ccgo_ts + 106,
	23: __ccgo_ts + 118,
	24: __ccgo_ts + 125,
	25: __ccgo_ts + 133,
	26: __ccgo_ts + 145,
	27: __ccgo_ts + 155,
	28: __ccgo_ts + 167,
	29: __ccgo_ts + 176,
	30: __ccgo_ts + 181,
	31: __ccgo_ts + 192,
	32: __ccgo_ts + 198,
	33: __ccgo_ts + 208,
	34: __ccgo_ts + 219,
	35: __ccgo_ts + 231,
	36: __ccgo_ts + 248,
	37: __ccgo_ts + 265,
	38: __ccgo_ts + 284,
	39: __ccgo_ts + 298,
	40: __ccgo_ts + 79,
}

var ts_symbol_map = [41]TSSymbol{
	1:  uint16(sym__identifier),
	2:  uint16(sym_comment),
	3:  uint16(sym_true),
	4:  uint16(sym_false),
	5:  uint16(sym_null),
	6:  uint16(sym_float),
	7:  uint16(sym_integer),
	8:  uint16(anon_sym_AMP),
	9:  uint16(anon_sym_LBRACK),
	10: uint16(anon_sym_RBRACK),
	11: uint16(anon_sym_EQ),
	12: uint16(sym_path),
	13: uint16(anon_sym_COLON),
	14: uint16(anon_sym_LBRACE),
	15: uint16(anon_sym_COMMA),
	16: uint16(anon_sym_RBRACE),
	17: uint16(anon_sym_LPAREN),
	18: uint16(anon_sym_RPAREN),
	19: uint16(sym_string),
	20: uint16(sym_resource),
	21: uint16(sym_identifier),
	22: uint16(sym_string_name),
	23: uint16(sym__value),
	24: uint16(sym_section),
	25: uint16(aux_sym__attributes),
	26: uint16(sym_attribute),
	27: uint16(aux_sym__properties),
	28: uint16(sym_property),
	29: uint16(sym_pair),
	30: uint16(sym_dictionary),
	31: uint16(sym_array),
	32: uint16(sym_arguments),
	33: uint16(sym__type_args),
	34: uint16(sym_constructor),
	35: uint16(aux_sym_resource_repeat1),
	36: uint16(aux_sym_resource_repeat2),
	37: uint16(aux_sym_dictionary_repeat1),
	38: uint16(aux_sym_array_repeat1),
	39: uint16(aux_sym_arguments_repeat1),
	40: uint16(anon_alias_sym_string),
}

var ts_symbol_metadata = [41]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	20: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	24: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	25: {},
	26: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	27: {},
	28: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	29: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	30: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	31: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	35: {},
	36: {},
	37: {},
	38: {},
	39: {},
	40: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
}

var ts_alias_sequences = [2][5]TSSymbol{
	0: {},
	1: {
		1: uint16(anon_alias_sym_string),
	},
}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [93]TSStateId{
	1:  uint16(1),
	2:  uint16(2),
	3:  uint16(3),
	4:  uint16(2),
	5:  uint16(3),
	6:  uint16(6),
	7:  uint16(7),
	8:  uint16(8),
	9:  uint16(6),
	10: uint16(10),
	11: uint16(11),
	12: uint16(12),
	13: uint16(13),
	14: uint16(14),
	15: uint16(15),
	16: uint16(16),
	17: uint16(17),
	18: uint16(18),
	19: uint16(19),
	20: uint16(20),
	21: uint16(21),
	22: uint16(22),
	23: uint16(23),
	24: uint16(24),
	25: uint16(25),
	26: uint16(26),
	27: uint16(27),
	28: uint16(28),
	29: uint16(29),
	30: uint16(15),
	31: uint16(31),
	32: uint16(32),
	33: uint16(33),
	34: uint16(34),
	35: uint16(35),
	36: uint16(36),
	37: uint16(37),
	38: uint16(38),
	39: uint16(39),
	40: uint16(40),
	41: uint16(41),
	42: uint16(16),
	43: uint16(43),
	44: uint16(44),
	45: uint16(44),
	46: uint16(46),
	47: uint16(47),
	48: uint16(48),
	49: uint16(49),
	50: uint16(50),
	51: uint16(51),
	52: uint16(52),
	53: uint16(53),
	54: uint16(47),
	55: uint16(55),
	56: uint16(56),
	57: uint16(57),
	58: uint16(58),
	59: uint16(59),
	60: uint16(53),
	61: uint16(58),
	62: uint16(62),
	63: uint16(56),
	64: uint16(59),
	65: uint16(46),
	66: uint16(27),
	67: uint16(67),
	68: uint16(21),
	69: uint16(18),
	70: uint16(23),
	71: uint16(26),
	72: uint16(17),
	73: uint16(28),
	74: uint16(25),
	75: uint16(75),
	76: uint16(76),
	77: uint16(77),
	78: uint16(78),
	79: uint16(79),
	80: uint16(75),
	81: uint16(24),
	82: uint16(19),
	83: uint16(20),
	84: uint16(22),
	85: uint16(85),
	86: uint16(86),
	87: uint16(87),
	88: uint16(87),
	89: uint16(89),
	90: uint16(90),
	91: uint16(91),
	92: uint16(92),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
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
		if eof != 0 {
			state = uint16(18)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(39)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(26)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('+') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(3)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('-') {
			state = uint16(10)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('-') {
			state = uint16(11)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('_') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('_') {
			state = uint16(8)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('_') {
			state = uint16(12)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(8):
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(9):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(10):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(11):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(12):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(13):
		if eof != 0 {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(14):
		if eof != 0 {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(15):
		if eof != 0 {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(16):
		if eof != 0 {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(17):
		if eof != 0 {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(39)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(18):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(19):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(20):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(1)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(22):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(1)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(23):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(24)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(24):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(11)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
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
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(1)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(1)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(8)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(12)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_path)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('/') <= lookahead && lookahead <= int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [22]uint16_t{
	0:  uint16('&'),
	1:  uint16(31),
	2:  uint16('('),
	3:  uint16(40),
	4:  uint16(')'),
	5:  uint16(41),
	6:  uint16(','),
	7:  uint16(38),
	8:  uint16('-'),
	9:  uint16(9),
	10: uint16('.'),
	11: uint16(2),
	12: uint16('0'),
	13: uint16(25),
	14: uint16(':'),
	15: uint16(36),
	16: uint16(';'),
	17: uint16(20),
	18: uint16('='),
	19: uint16(34),
	20: uint16('['),
	21: uint16(32),
}

var map_token1 = [22]uint16_t{
	0:  uint16('-'),
	1:  uint16(9),
	2:  uint16('.'),
	3:  uint16(22),
	4:  uint16('_'),
	5:  uint16(27),
	6:  uint16('B'),
	7:  uint16(4),
	8:  uint16('b'),
	9:  uint16(4),
	10: uint16('E'),
	11: uint16(1),
	12: uint16('e'),
	13: uint16(1),
	14: uint16('O'),
	15: uint16(5),
	16: uint16('o'),
	17: uint16(5),
	18: uint16('X'),
	19: uint16(6),
	20: uint16('x'),
	21: uint16(6),
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
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(4)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('a') {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('u') {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('r') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('l') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('l') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('u') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('s') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('l') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('e') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('e') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(13):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_null)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(14):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(15):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [93]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state: uint16(17),
	},
	2: {
		Fexternal_lex_state: uint16(1),
	},
	3: {
		Fexternal_lex_state: uint16(1),
	},
	4: {
		Fexternal_lex_state: uint16(1),
	},
	5: {
		Fexternal_lex_state: uint16(1),
	},
	6: {
		Fexternal_lex_state: uint16(1),
	},
	7: {
		Fexternal_lex_state: uint16(1),
	},
	8: {
		Fexternal_lex_state: uint16(1),
	},
	9: {
		Fexternal_lex_state: uint16(1),
	},
	10: {
		Fexternal_lex_state: uint16(1),
	},
	11: {
		Fexternal_lex_state: uint16(1),
	},
	12: {
		Fexternal_lex_state: uint16(1),
	},
	13: {
		Fexternal_lex_state: uint16(1),
	},
	14: {
		Fexternal_lex_state: uint16(1),
	},
	15: {
		Flex_state: uint16(17),
	},
	16: {},
	17: {
		Flex_state: uint16(17),
	},
	18: {
		Flex_state: uint16(17),
	},
	19: {
		Flex_state: uint16(17),
	},
	20: {
		Flex_state: uint16(17),
	},
	21: {
		Flex_state: uint16(17),
	},
	22: {
		Flex_state: uint16(17),
	},
	23: {
		Flex_state: uint16(17),
	},
	24: {
		Flex_state: uint16(17),
	},
	25: {
		Flex_state: uint16(17),
	},
	26: {
		Flex_state: uint16(17),
	},
	27: {
		Flex_state: uint16(17),
	},
	28: {
		Flex_state: uint16(17),
	},
	29: {
		Flex_state: uint16(17),
	},
	30: {},
	31: {
		Flex_state: uint16(17),
	},
	32: {
		Flex_state: uint16(17),
	},
	33: {},
	34: {},
	35: {
		Flex_state: uint16(17),
	},
	36: {
		Flex_state: uint16(17),
	},
	37: {},
	38: {
		Flex_state: uint16(17),
	},
	39: {
		Flex_state: uint16(17),
	},
	40: {},
	41: {},
	42: {
		Flex_state: uint16(17),
	},
	43: {},
	44: {},
	45: {},
	46: {},
	47: {},
	48: {},
	49: {},
	50: {
		Flex_state: uint16(17),
	},
	51: {},
	52: {},
	53: {},
	54: {},
	55: {},
	56: {},
	57: {},
	58: {},
	59: {},
	60: {},
	61: {},
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
	83: {},
	84: {},
	85: {},
	86: {},
	87: {
		Fexternal_lex_state: uint16(1),
	},
	88: {
		Fexternal_lex_state: uint16(1),
	},
	89: {},
	90: {},
	91: {},
	92: {},
}

var ts_parse_table = [2][40]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(3),
		3:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		6:  uint16(1),
		7:  uint16(1),
		8:  uint16(1),
		9:  uint16(1),
		10: uint16(1),
		11: uint16(1),
		13: uint16(1),
		14: uint16(1),
		15: uint16(1),
		16: uint16(1),
		17: uint16(1),
		18: uint16(1),
		19: uint16(1),
	},
	1: {
		0:  uint16(5),
		2:  uint16(3),
		9:  uint16(7),
		12: uint16(9),
		20: uint16(92),
		24: uint16(41),
		28: uint16(29),
		35: uint16(29),
		36: uint16(41),
	},
}

var ts_small_parse_table = [1442]uint16_t{
	0:    uint16(11),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(11),
	5:    uint16(1),
	6:    uint16(sym__identifier),
	7:    uint16(17),
	8:    uint16(1),
	9:    uint16(anon_sym_AMP),
	10:   uint16(19),
	11:   uint16(1),
	12:   uint16(anon_sym_LBRACK),
	13:   uint16(21),
	14:   uint16(1),
	15:   uint16(anon_sym_LBRACE),
	16:   uint16(23),
	17:   uint16(1),
	18:   uint16(anon_sym_RPAREN),
	19:   uint16(15),
	20:   uint16(1),
	21:   uint16(sym_identifier),
	22:   uint16(46),
	23:   uint16(1),
	24:   uint16(sym_pair),
	25:   uint16(15),
	26:   uint16(2),
	27:   uint16(sym_string),
	28:   uint16(sym_float),
	29:   uint16(13),
	30:   uint16(4),
	31:   uint16(sym_true),
	32:   uint16(sym_false),
	33:   uint16(sym_null),
	34:   uint16(sym_integer),
	35:   uint16(44),
	36:   uint16(5),
	37:   uint16(sym_string_name),
	38:   uint16(sym__value),
	39:   uint16(sym_dictionary),
	40:   uint16(sym_array),
	41:   uint16(sym_constructor),
	42:   uint16(11),
	43:   uint16(3),
	44:   uint16(1),
	45:   uint16(sym_comment),
	46:   uint16(11),
	47:   uint16(1),
	48:   uint16(sym__identifier),
	49:   uint16(17),
	50:   uint16(1),
	51:   uint16(anon_sym_AMP),
	52:   uint16(19),
	53:   uint16(1),
	54:   uint16(anon_sym_LBRACK),
	55:   uint16(21),
	56:   uint16(1),
	57:   uint16(anon_sym_LBRACE),
	58:   uint16(29),
	59:   uint16(1),
	60:   uint16(anon_sym_RBRACE),
	61:   uint16(15),
	62:   uint16(1),
	63:   uint16(sym_identifier),
	64:   uint16(58),
	65:   uint16(1),
	66:   uint16(sym_pair),
	67:   uint16(27),
	68:   uint16(2),
	69:   uint16(sym_string),
	70:   uint16(sym_float),
	71:   uint16(25),
	72:   uint16(4),
	73:   uint16(sym_true),
	74:   uint16(sym_false),
	75:   uint16(sym_null),
	76:   uint16(sym_integer),
	77:   uint16(85),
	78:   uint16(5),
	79:   uint16(sym_string_name),
	80:   uint16(sym__value),
	81:   uint16(sym_dictionary),
	82:   uint16(sym_array),
	83:   uint16(sym_constructor),
	84:   uint16(11),
	85:   uint16(3),
	86:   uint16(1),
	87:   uint16(sym_comment),
	88:   uint16(11),
	89:   uint16(1),
	90:   uint16(sym__identifier),
	91:   uint16(17),
	92:   uint16(1),
	93:   uint16(anon_sym_AMP),
	94:   uint16(19),
	95:   uint16(1),
	96:   uint16(anon_sym_LBRACK),
	97:   uint16(21),
	98:   uint16(1),
	99:   uint16(anon_sym_LBRACE),
	100:  uint16(35),
	101:  uint16(1),
	102:  uint16(anon_sym_RPAREN),
	103:  uint16(15),
	104:  uint16(1),
	105:  uint16(sym_identifier),
	106:  uint16(65),
	107:  uint16(1),
	108:  uint16(sym_pair),
	109:  uint16(33),
	110:  uint16(2),
	111:  uint16(sym_string),
	112:  uint16(sym_float),
	113:  uint16(31),
	114:  uint16(4),
	115:  uint16(sym_true),
	116:  uint16(sym_false),
	117:  uint16(sym_null),
	118:  uint16(sym_integer),
	119:  uint16(45),
	120:  uint16(5),
	121:  uint16(sym_string_name),
	122:  uint16(sym__value),
	123:  uint16(sym_dictionary),
	124:  uint16(sym_array),
	125:  uint16(sym_constructor),
	126:  uint16(11),
	127:  uint16(3),
	128:  uint16(1),
	129:  uint16(sym_comment),
	130:  uint16(11),
	131:  uint16(1),
	132:  uint16(sym__identifier),
	133:  uint16(17),
	134:  uint16(1),
	135:  uint16(anon_sym_AMP),
	136:  uint16(19),
	137:  uint16(1),
	138:  uint16(anon_sym_LBRACK),
	139:  uint16(21),
	140:  uint16(1),
	141:  uint16(anon_sym_LBRACE),
	142:  uint16(37),
	143:  uint16(1),
	144:  uint16(anon_sym_RBRACE),
	145:  uint16(15),
	146:  uint16(1),
	147:  uint16(sym_identifier),
	148:  uint16(61),
	149:  uint16(1),
	150:  uint16(sym_pair),
	151:  uint16(27),
	152:  uint16(2),
	153:  uint16(sym_string),
	154:  uint16(sym_float),
	155:  uint16(25),
	156:  uint16(4),
	157:  uint16(sym_true),
	158:  uint16(sym_false),
	159:  uint16(sym_null),
	160:  uint16(sym_integer),
	161:  uint16(85),
	162:  uint16(5),
	163:  uint16(sym_string_name),
	164:  uint16(sym__value),
	165:  uint16(sym_dictionary),
	166:  uint16(sym_array),
	167:  uint16(sym_constructor),
	168:  uint16(10),
	169:  uint16(3),
	170:  uint16(1),
	171:  uint16(sym_comment),
	172:  uint16(11),
	173:  uint16(1),
	174:  uint16(sym__identifier),
	175:  uint16(17),
	176:  uint16(1),
	177:  uint16(anon_sym_AMP),
	178:  uint16(19),
	179:  uint16(1),
	180:  uint16(anon_sym_LBRACK),
	181:  uint16(21),
	182:  uint16(1),
	183:  uint16(anon_sym_LBRACE),
	184:  uint16(43),
	185:  uint16(1),
	186:  uint16(anon_sym_RBRACK),
	187:  uint16(15),
	188:  uint16(1),
	189:  uint16(sym_identifier),
	190:  uint16(41),
	191:  uint16(2),
	192:  uint16(sym_string),
	193:  uint16(sym_float),
	194:  uint16(39),
	195:  uint16(4),
	196:  uint16(sym_true),
	197:  uint16(sym_false),
	198:  uint16(sym_null),
	199:  uint16(sym_integer),
	200:  uint16(53),
	201:  uint16(5),
	202:  uint16(sym_string_name),
	203:  uint16(sym__value),
	204:  uint16(sym_dictionary),
	205:  uint16(sym_array),
	206:  uint16(sym_constructor),
	207:  uint16(10),
	208:  uint16(3),
	209:  uint16(1),
	210:  uint16(sym_comment),
	211:  uint16(11),
	212:  uint16(1),
	213:  uint16(sym__identifier),
	214:  uint16(17),
	215:  uint16(1),
	216:  uint16(anon_sym_AMP),
	217:  uint16(19),
	218:  uint16(1),
	219:  uint16(anon_sym_LBRACK),
	220:  uint16(21),
	221:  uint16(1),
	222:  uint16(anon_sym_LBRACE),
	223:  uint16(15),
	224:  uint16(1),
	225:  uint16(sym_identifier),
	226:  uint16(79),
	227:  uint16(1),
	228:  uint16(sym_pair),
	229:  uint16(27),
	230:  uint16(2),
	231:  uint16(sym_string),
	232:  uint16(sym_float),
	233:  uint16(25),
	234:  uint16(4),
	235:  uint16(sym_true),
	236:  uint16(sym_false),
	237:  uint16(sym_null),
	238:  uint16(sym_integer),
	239:  uint16(85),
	240:  uint16(5),
	241:  uint16(sym_string_name),
	242:  uint16(sym__value),
	243:  uint16(sym_dictionary),
	244:  uint16(sym_array),
	245:  uint16(sym_constructor),
	246:  uint16(10),
	247:  uint16(3),
	248:  uint16(1),
	249:  uint16(sym_comment),
	250:  uint16(11),
	251:  uint16(1),
	252:  uint16(sym__identifier),
	253:  uint16(17),
	254:  uint16(1),
	255:  uint16(anon_sym_AMP),
	256:  uint16(19),
	257:  uint16(1),
	258:  uint16(anon_sym_LBRACK),
	259:  uint16(21),
	260:  uint16(1),
	261:  uint16(anon_sym_LBRACE),
	262:  uint16(15),
	263:  uint16(1),
	264:  uint16(sym_identifier),
	265:  uint16(76),
	266:  uint16(1),
	267:  uint16(sym_pair),
	268:  uint16(47),
	269:  uint16(2),
	270:  uint16(sym_string),
	271:  uint16(sym_float),
	272:  uint16(45),
	273:  uint16(4),
	274:  uint16(sym_true),
	275:  uint16(sym_false),
	276:  uint16(sym_null),
	277:  uint16(sym_integer),
	278:  uint16(55),
	279:  uint16(5),
	280:  uint16(sym_string_name),
	281:  uint16(sym__value),
	282:  uint16(sym_dictionary),
	283:  uint16(sym_array),
	284:  uint16(sym_constructor),
	285:  uint16(10),
	286:  uint16(3),
	287:  uint16(1),
	288:  uint16(sym_comment),
	289:  uint16(11),
	290:  uint16(1),
	291:  uint16(sym__identifier),
	292:  uint16(17),
	293:  uint16(1),
	294:  uint16(anon_sym_AMP),
	295:  uint16(19),
	296:  uint16(1),
	297:  uint16(anon_sym_LBRACK),
	298:  uint16(21),
	299:  uint16(1),
	300:  uint16(anon_sym_LBRACE),
	301:  uint16(53),
	302:  uint16(1),
	303:  uint16(anon_sym_RBRACK),
	304:  uint16(15),
	305:  uint16(1),
	306:  uint16(sym_identifier),
	307:  uint16(51),
	308:  uint16(2),
	309:  uint16(sym_string),
	310:  uint16(sym_float),
	311:  uint16(49),
	312:  uint16(4),
	313:  uint16(sym_true),
	314:  uint16(sym_false),
	315:  uint16(sym_null),
	316:  uint16(sym_integer),
	317:  uint16(60),
	318:  uint16(5),
	319:  uint16(sym_string_name),
	320:  uint16(sym__value),
	321:  uint16(sym_dictionary),
	322:  uint16(sym_array),
	323:  uint16(sym_constructor),
	324:  uint16(9),
	325:  uint16(3),
	326:  uint16(1),
	327:  uint16(sym_comment),
	328:  uint16(11),
	329:  uint16(1),
	330:  uint16(sym__identifier),
	331:  uint16(17),
	332:  uint16(1),
	333:  uint16(anon_sym_AMP),
	334:  uint16(19),
	335:  uint16(1),
	336:  uint16(anon_sym_LBRACK),
	337:  uint16(21),
	338:  uint16(1),
	339:  uint16(anon_sym_LBRACE),
	340:  uint16(15),
	341:  uint16(1),
	342:  uint16(sym_identifier),
	343:  uint16(57),
	344:  uint16(2),
	345:  uint16(sym_string),
	346:  uint16(sym_float),
	347:  uint16(55),
	348:  uint16(4),
	349:  uint16(sym_true),
	350:  uint16(sym_false),
	351:  uint16(sym_null),
	352:  uint16(sym_integer),
	353:  uint16(49),
	354:  uint16(5),
	355:  uint16(sym_string_name),
	356:  uint16(sym__value),
	357:  uint16(sym_dictionary),
	358:  uint16(sym_array),
	359:  uint16(sym_constructor),
	360:  uint16(9),
	361:  uint16(3),
	362:  uint16(1),
	363:  uint16(sym_comment),
	364:  uint16(11),
	365:  uint16(1),
	366:  uint16(sym__identifier),
	367:  uint16(17),
	368:  uint16(1),
	369:  uint16(anon_sym_AMP),
	370:  uint16(19),
	371:  uint16(1),
	372:  uint16(anon_sym_LBRACK),
	373:  uint16(21),
	374:  uint16(1),
	375:  uint16(anon_sym_LBRACE),
	376:  uint16(15),
	377:  uint16(1),
	378:  uint16(sym_identifier),
	379:  uint16(61),
	380:  uint16(2),
	381:  uint16(sym_string),
	382:  uint16(sym_float),
	383:  uint16(59),
	384:  uint16(4),
	385:  uint16(sym_true),
	386:  uint16(sym_false),
	387:  uint16(sym_null),
	388:  uint16(sym_integer),
	389:  uint16(62),
	390:  uint16(5),
	391:  uint16(sym_string_name),
	392:  uint16(sym__value),
	393:  uint16(sym_dictionary),
	394:  uint16(sym_array),
	395:  uint16(sym_constructor),
	396:  uint16(9),
	397:  uint16(3),
	398:  uint16(1),
	399:  uint16(sym_comment),
	400:  uint16(17),
	401:  uint16(1),
	402:  uint16(anon_sym_AMP),
	403:  uint16(19),
	404:  uint16(1),
	405:  uint16(anon_sym_LBRACK),
	406:  uint16(21),
	407:  uint16(1),
	408:  uint16(anon_sym_LBRACE),
	409:  uint16(63),
	410:  uint16(1),
	411:  uint16(sym__identifier),
	412:  uint16(15),
	413:  uint16(1),
	414:  uint16(sym_identifier),
	415:  uint16(67),
	416:  uint16(2),
	417:  uint16(sym_string),
	418:  uint16(sym_float),
	419:  uint16(65),
	420:  uint16(4),
	421:  uint16(sym_true),
	422:  uint16(sym_false),
	423:  uint16(sym_null),
	424:  uint16(sym_integer),
	425:  uint16(50),
	426:  uint16(5),
	427:  uint16(sym_string_name),
	428:  uint16(sym__value),
	429:  uint16(sym_dictionary),
	430:  uint16(sym_array),
	431:  uint16(sym_constructor),
	432:  uint16(9),
	433:  uint16(3),
	434:  uint16(1),
	435:  uint16(sym_comment),
	436:  uint16(11),
	437:  uint16(1),
	438:  uint16(sym__identifier),
	439:  uint16(17),
	440:  uint16(1),
	441:  uint16(anon_sym_AMP),
	442:  uint16(19),
	443:  uint16(1),
	444:  uint16(anon_sym_LBRACK),
	445:  uint16(21),
	446:  uint16(1),
	447:  uint16(anon_sym_LBRACE),
	448:  uint16(15),
	449:  uint16(1),
	450:  uint16(sym_identifier),
	451:  uint16(71),
	452:  uint16(2),
	453:  uint16(sym_string),
	454:  uint16(sym_float),
	455:  uint16(69),
	456:  uint16(4),
	457:  uint16(sym_true),
	458:  uint16(sym_false),
	459:  uint16(sym_null),
	460:  uint16(sym_integer),
	461:  uint16(67),
	462:  uint16(5),
	463:  uint16(sym_string_name),
	464:  uint16(sym__value),
	465:  uint16(sym_dictionary),
	466:  uint16(sym_array),
	467:  uint16(sym_constructor),
	468:  uint16(9),
	469:  uint16(3),
	470:  uint16(1),
	471:  uint16(sym_comment),
	472:  uint16(11),
	473:  uint16(1),
	474:  uint16(sym__identifier),
	475:  uint16(77),
	476:  uint16(1),
	477:  uint16(anon_sym_AMP),
	478:  uint16(79),
	479:  uint16(1),
	480:  uint16(anon_sym_LBRACK),
	481:  uint16(81),
	482:  uint16(1),
	483:  uint16(anon_sym_LBRACE),
	484:  uint16(30),
	485:  uint16(1),
	486:  uint16(sym_identifier),
	487:  uint16(75),
	488:  uint16(2),
	489:  uint16(sym_string),
	490:  uint16(sym_float),
	491:  uint16(73),
	492:  uint16(4),
	493:  uint16(sym_true),
	494:  uint16(sym_false),
	495:  uint16(sym_null),
	496:  uint16(sym_integer),
	497:  uint16(78),
	498:  uint16(5),
	499:  uint16(sym_string_name),
	500:  uint16(sym__value),
	501:  uint16(sym_dictionary),
	502:  uint16(sym_array),
	503:  uint16(sym_constructor),
	504:  uint16(6),
	505:  uint16(3),
	506:  uint16(1),
	507:  uint16(sym_comment),
	508:  uint16(85),
	509:  uint16(1),
	510:  uint16(anon_sym_LBRACK),
	511:  uint16(87),
	512:  uint16(1),
	513:  uint16(anon_sym_LPAREN),
	514:  uint16(22),
	515:  uint16(1),
	516:  uint16(sym_arguments),
	517:  uint16(75),
	518:  uint16(1),
	519:  uint16(sym__type_args),
	520:  uint16(83),
	521:  uint16(7),
	523:  uint16(anon_sym_RBRACK),
	524:  uint16(sym_path),
	525:  uint16(anon_sym_COLON),
	526:  uint16(anon_sym_COMMA),
	527:  uint16(anon_sym_RBRACE),
	528:  uint16(anon_sym_RPAREN),
	529:  uint16(2),
	530:  uint16(3),
	531:  uint16(1),
	532:  uint16(sym_comment),
	533:  uint16(89),
	534:  uint16(9),
	535:  uint16(sym__identifier),
	536:  uint16(anon_sym_LBRACK),
	537:  uint16(anon_sym_RBRACK),
	538:  uint16(anon_sym_EQ),
	539:  uint16(anon_sym_COLON),
	540:  uint16(anon_sym_COMMA),
	541:  uint16(anon_sym_RBRACE),
	542:  uint16(anon_sym_LPAREN),
	543:  uint16(anon_sym_RPAREN),
	544:  uint16(2),
	545:  uint16(3),
	546:  uint16(1),
	547:  uint16(sym_comment),
	548:  uint16(91),
	549:  uint16(8),
	551:  uint16(anon_sym_LBRACK),
	552:  uint16(anon_sym_RBRACK),
	553:  uint16(sym_path),
	554:  uint16(anon_sym_COLON),
	555:  uint16(anon_sym_COMMA),
	556:  uint16(anon_sym_RBRACE),
	557:  uint16(anon_sym_RPAREN),
	558:  uint16(2),
	559:  uint16(3),
	560:  uint16(1),
	561:  uint16(sym_comment),
	562:  uint16(93),
	563:  uint16(8),
	565:  uint16(anon_sym_LBRACK),
	566:  uint16(anon_sym_RBRACK),
	567:  uint16(sym_path),
	568:  uint16(anon_sym_COLON),
	569:  uint16(anon_sym_COMMA),
	570:  uint16(anon_sym_RBRACE),
	571:  uint16(anon_sym_RPAREN),
	572:  uint16(2),
	573:  uint16(3),
	574:  uint16(1),
	575:  uint16(sym_comment),
	576:  uint16(95),
	577:  uint16(8),
	579:  uint16(anon_sym_LBRACK),
	580:  uint16(anon_sym_RBRACK),
	581:  uint16(sym_path),
	582:  uint16(anon_sym_COLON),
	583:  uint16(anon_sym_COMMA),
	584:  uint16(anon_sym_RBRACE),
	585:  uint16(anon_sym_RPAREN),
	586:  uint16(2),
	587:  uint16(3),
	588:  uint16(1),
	589:  uint16(sym_comment),
	590:  uint16(97),
	591:  uint16(8),
	593:  uint16(anon_sym_LBRACK),
	594:  uint16(anon_sym_RBRACK),
	595:  uint16(sym_path),
	596:  uint16(anon_sym_COLON),
	597:  uint16(anon_sym_COMMA),
	598:  uint16(anon_sym_RBRACE),
	599:  uint16(anon_sym_RPAREN),
	600:  uint16(2),
	601:  uint16(3),
	602:  uint16(1),
	603:  uint16(sym_comment),
	604:  uint16(99),
	605:  uint16(8),
	607:  uint16(anon_sym_LBRACK),
	608:  uint16(anon_sym_RBRACK),
	609:  uint16(sym_path),
	610:  uint16(anon_sym_COLON),
	611:  uint16(anon_sym_COMMA),
	612:  uint16(anon_sym_RBRACE),
	613:  uint16(anon_sym_RPAREN),
	614:  uint16(2),
	615:  uint16(3),
	616:  uint16(1),
	617:  uint16(sym_comment),
	618:  uint16(101),
	619:  uint16(8),
	621:  uint16(anon_sym_LBRACK),
	622:  uint16(anon_sym_RBRACK),
	623:  uint16(sym_path),
	624:  uint16(anon_sym_COLON),
	625:  uint16(anon_sym_COMMA),
	626:  uint16(anon_sym_RBRACE),
	627:  uint16(anon_sym_RPAREN),
	628:  uint16(2),
	629:  uint16(3),
	630:  uint16(1),
	631:  uint16(sym_comment),
	632:  uint16(103),
	633:  uint16(8),
	635:  uint16(anon_sym_LBRACK),
	636:  uint16(anon_sym_RBRACK),
	637:  uint16(sym_path),
	638:  uint16(anon_sym_COLON),
	639:  uint16(anon_sym_COMMA),
	640:  uint16(anon_sym_RBRACE),
	641:  uint16(anon_sym_RPAREN),
	642:  uint16(2),
	643:  uint16(3),
	644:  uint16(1),
	645:  uint16(sym_comment),
	646:  uint16(105),
	647:  uint16(8),
	649:  uint16(anon_sym_LBRACK),
	650:  uint16(anon_sym_RBRACK),
	651:  uint16(sym_path),
	652:  uint16(anon_sym_COLON),
	653:  uint16(anon_sym_COMMA),
	654:  uint16(anon_sym_RBRACE),
	655:  uint16(anon_sym_RPAREN),
	656:  uint16(2),
	657:  uint16(3),
	658:  uint16(1),
	659:  uint16(sym_comment),
	660:  uint16(107),
	661:  uint16(8),
	663:  uint16(anon_sym_LBRACK),
	664:  uint16(anon_sym_RBRACK),
	665:  uint16(sym_path),
	666:  uint16(anon_sym_COLON),
	667:  uint16(anon_sym_COMMA),
	668:  uint16(anon_sym_RBRACE),
	669:  uint16(anon_sym_RPAREN),
	670:  uint16(2),
	671:  uint16(3),
	672:  uint16(1),
	673:  uint16(sym_comment),
	674:  uint16(109),
	675:  uint16(8),
	677:  uint16(anon_sym_LBRACK),
	678:  uint16(anon_sym_RBRACK),
	679:  uint16(sym_path),
	680:  uint16(anon_sym_COLON),
	681:  uint16(anon_sym_COMMA),
	682:  uint16(anon_sym_RBRACE),
	683:  uint16(anon_sym_RPAREN),
	684:  uint16(2),
	685:  uint16(3),
	686:  uint16(1),
	687:  uint16(sym_comment),
	688:  uint16(111),
	689:  uint16(8),
	691:  uint16(anon_sym_LBRACK),
	692:  uint16(anon_sym_RBRACK),
	693:  uint16(sym_path),
	694:  uint16(anon_sym_COLON),
	695:  uint16(anon_sym_COMMA),
	696:  uint16(anon_sym_RBRACE),
	697:  uint16(anon_sym_RPAREN),
	698:  uint16(2),
	699:  uint16(3),
	700:  uint16(1),
	701:  uint16(sym_comment),
	702:  uint16(113),
	703:  uint16(8),
	705:  uint16(anon_sym_LBRACK),
	706:  uint16(anon_sym_RBRACK),
	707:  uint16(sym_path),
	708:  uint16(anon_sym_COLON),
	709:  uint16(anon_sym_COMMA),
	710:  uint16(anon_sym_RBRACE),
	711:  uint16(anon_sym_RPAREN),
	712:  uint16(6),
	713:  uint16(3),
	714:  uint16(1),
	715:  uint16(sym_comment),
	716:  uint16(7),
	717:  uint16(1),
	718:  uint16(anon_sym_LBRACK),
	719:  uint16(9),
	720:  uint16(1),
	721:  uint16(sym_path),
	722:  uint16(115),
	723:  uint16(1),
	725:  uint16(36),
	726:  uint16(2),
	727:  uint16(sym_property),
	728:  uint16(aux_sym_resource_repeat1),
	729:  uint16(43),
	730:  uint16(2),
	731:  uint16(sym_section),
	732:  uint16(aux_sym_resource_repeat2),
	733:  uint16(6),
	734:  uint16(3),
	735:  uint16(1),
	736:  uint16(sym_comment),
	737:  uint16(85),
	738:  uint16(1),
	739:  uint16(anon_sym_LBRACK),
	740:  uint16(117),
	741:  uint16(1),
	742:  uint16(anon_sym_LPAREN),
	743:  uint16(80),
	744:  uint16(1),
	745:  uint16(sym__type_args),
	746:  uint16(84),
	747:  uint16(1),
	748:  uint16(sym_arguments),
	749:  uint16(83),
	750:  uint16(2),
	751:  uint16(sym__identifier),
	752:  uint16(anon_sym_RBRACK),
	753:  uint16(4),
	754:  uint16(3),
	755:  uint16(1),
	756:  uint16(sym_comment),
	757:  uint16(9),
	758:  uint16(1),
	759:  uint16(sym_path),
	760:  uint16(119),
	761:  uint16(2),
	763:  uint16(anon_sym_LBRACK),
	764:  uint16(39),
	765:  uint16(2),
	766:  uint16(aux_sym__properties),
	767:  uint16(sym_property),
	768:  uint16(4),
	769:  uint16(3),
	770:  uint16(1),
	771:  uint16(sym_comment),
	772:  uint16(9),
	773:  uint16(1),
	774:  uint16(sym_path),
	775:  uint16(121),
	776:  uint16(2),
	778:  uint16(anon_sym_LBRACK),
	779:  uint16(35),
	780:  uint16(2),
	781:  uint16(aux_sym__properties),
	782:  uint16(sym_property),
	783:  uint16(5),
	784:  uint16(3),
	785:  uint16(1),
	786:  uint16(sym_comment),
	787:  uint16(123),
	788:  uint16(1),
	789:  uint16(sym__identifier),
	790:  uint16(125),
	791:  uint16(1),
	792:  uint16(anon_sym_RBRACK),
	793:  uint16(91),
	794:  uint16(1),
	795:  uint16(sym_identifier),
	796:  uint16(37),
	797:  uint16(2),
	798:  uint16(aux_sym__attributes),
	799:  uint16(sym_attribute),
	800:  uint16(5),
	801:  uint16(3),
	802:  uint16(1),
	803:  uint16(sym_comment),
	804:  uint16(123),
	805:  uint16(1),
	806:  uint16(sym__identifier),
	807:  uint16(127),
	808:  uint16(1),
	809:  uint16(anon_sym_RBRACK),
	810:  uint16(91),
	811:  uint16(1),
	812:  uint16(sym_identifier),
	813:  uint16(33),
	814:  uint16(2),
	815:  uint16(aux_sym__attributes),
	816:  uint16(sym_attribute),
	817:  uint16(4),
	818:  uint16(3),
	819:  uint16(1),
	820:  uint16(sym_comment),
	821:  uint16(131),
	822:  uint16(1),
	823:  uint16(sym_path),
	824:  uint16(129),
	825:  uint16(2),
	827:  uint16(anon_sym_LBRACK),
	828:  uint16(35),
	829:  uint16(2),
	830:  uint16(aux_sym__properties),
	831:  uint16(sym_property),
	832:  uint16(4),
	833:  uint16(3),
	834:  uint16(1),
	835:  uint16(sym_comment),
	836:  uint16(136),
	837:  uint16(1),
	838:  uint16(sym_path),
	839:  uint16(134),
	840:  uint16(2),
	842:  uint16(anon_sym_LBRACK),
	843:  uint16(36),
	844:  uint16(2),
	845:  uint16(sym_property),
	846:  uint16(aux_sym_resource_repeat1),
	847:  uint16(5),
	848:  uint16(3),
	849:  uint16(1),
	850:  uint16(sym_comment),
	851:  uint16(139),
	852:  uint16(1),
	853:  uint16(sym__identifier),
	854:  uint16(142),
	855:  uint16(1),
	856:  uint16(anon_sym_RBRACK),
	857:  uint16(91),
	858:  uint16(1),
	859:  uint16(sym_identifier),
	860:  uint16(37),
	861:  uint16(2),
	862:  uint16(aux_sym__attributes),
	863:  uint16(sym_attribute),
	864:  uint16(4),
	865:  uint16(3),
	866:  uint16(1),
	867:  uint16(sym_comment),
	868:  uint16(9),
	869:  uint16(1),
	870:  uint16(sym_path),
	871:  uint16(144),
	872:  uint16(2),
	874:  uint16(anon_sym_LBRACK),
	875:  uint16(32),
	876:  uint16(2),
	877:  uint16(aux_sym__properties),
	878:  uint16(sym_property),
	879:  uint16(4),
	880:  uint16(3),
	881:  uint16(1),
	882:  uint16(sym_comment),
	883:  uint16(9),
	884:  uint16(1),
	885:  uint16(sym_path),
	886:  uint16(144),
	887:  uint16(2),
	889:  uint16(anon_sym_LBRACK),
	890:  uint16(35),
	891:  uint16(2),
	892:  uint16(aux_sym__properties),
	893:  uint16(sym_property),
	894:  uint16(4),
	895:  uint16(3),
	896:  uint16(1),
	897:  uint16(sym_comment),
	898:  uint16(146),
	899:  uint16(1),
	901:  uint16(148),
	902:  uint16(1),
	903:  uint16(anon_sym_LBRACK),
	904:  uint16(40),
	905:  uint16(2),
	906:  uint16(sym_section),
	907:  uint16(aux_sym_resource_repeat2),
	908:  uint16(4),
	909:  uint16(3),
	910:  uint16(1),
	911:  uint16(sym_comment),
	912:  uint16(7),
	913:  uint16(1),
	914:  uint16(anon_sym_LBRACK),
	915:  uint16(115),
	916:  uint16(1),
	918:  uint16(40),
	919:  uint16(2),
	920:  uint16(sym_section),
	921:  uint16(aux_sym_resource_repeat2),
	922:  uint16(2),
	923:  uint16(3),
	924:  uint16(1),
	925:  uint16(sym_comment),
	926:  uint16(89),
	927:  uint16(4),
	929:  uint16(anon_sym_LBRACK),
	930:  uint16(sym_path),
	931:  uint16(anon_sym_LPAREN),
	932:  uint16(4),
	933:  uint16(3),
	934:  uint16(1),
	935:  uint16(sym_comment),
	936:  uint16(7),
	937:  uint16(1),
	938:  uint16(anon_sym_LBRACK),
	939:  uint16(151),
	940:  uint16(1),
	942:  uint16(40),
	943:  uint16(2),
	944:  uint16(sym_section),
	945:  uint16(aux_sym_resource_repeat2),
	946:  uint16(5),
	947:  uint16(3),
	948:  uint16(1),
	949:  uint16(sym_comment),
	950:  uint16(153),
	951:  uint16(1),
	952:  uint16(anon_sym_COLON),
	953:  uint16(155),
	954:  uint16(1),
	955:  uint16(anon_sym_COMMA),
	956:  uint16(157),
	957:  uint16(1),
	958:  uint16(anon_sym_RPAREN),
	959:  uint16(54),
	960:  uint16(1),
	961:  uint16(aux_sym_arguments_repeat1),
	962:  uint16(5),
	963:  uint16(3),
	964:  uint16(1),
	965:  uint16(sym_comment),
	966:  uint16(153),
	967:  uint16(1),
	968:  uint16(anon_sym_COLON),
	969:  uint16(155),
	970:  uint16(1),
	971:  uint16(anon_sym_COMMA),
	972:  uint16(159),
	973:  uint16(1),
	974:  uint16(anon_sym_RPAREN),
	975:  uint16(47),
	976:  uint16(1),
	977:  uint16(aux_sym_arguments_repeat1),
	978:  uint16(4),
	979:  uint16(3),
	980:  uint16(1),
	981:  uint16(sym_comment),
	982:  uint16(155),
	983:  uint16(1),
	984:  uint16(anon_sym_COMMA),
	985:  uint16(157),
	986:  uint16(1),
	987:  uint16(anon_sym_RPAREN),
	988:  uint16(54),
	989:  uint16(1),
	990:  uint16(aux_sym_arguments_repeat1),
	991:  uint16(4),
	992:  uint16(3),
	993:  uint16(1),
	994:  uint16(sym_comment),
	995:  uint16(155),
	996:  uint16(1),
	997:  uint16(anon_sym_COMMA),
	998:  uint16(161),
	999:  uint16(1),
	1000: uint16(anon_sym_RPAREN),
	1001: uint16(57),
	1002: uint16(1),
	1003: uint16(aux_sym_arguments_repeat1),
	1004: uint16(4),
	1005: uint16(3),
	1006: uint16(1),
	1007: uint16(sym_comment),
	1008: uint16(163),
	1009: uint16(1),
	1010: uint16(anon_sym_RBRACK),
	1011: uint16(165),
	1012: uint16(1),
	1013: uint16(anon_sym_COMMA),
	1014: uint16(48),
	1015: uint16(1),
	1016: uint16(aux_sym_array_repeat1),
	1017: uint16(2),
	1018: uint16(3),
	1019: uint16(1),
	1020: uint16(sym_comment),
	1021: uint16(168),
	1022: uint16(3),
	1023: uint16(anon_sym_COMMA),
	1024: uint16(anon_sym_RBRACE),
	1025: uint16(anon_sym_RPAREN),
	1026: uint16(2),
	1027: uint16(3),
	1028: uint16(1),
	1029: uint16(sym_comment),
	1030: uint16(170),
	1031: uint16(3),
	1033: uint16(anon_sym_LBRACK),
	1034: uint16(sym_path),
	1035: uint16(4),
	1036: uint16(3),
	1037: uint16(1),
	1038: uint16(sym_comment),
	1039: uint16(172),
	1040: uint16(1),
	1041: uint16(anon_sym_COMMA),
	1042: uint16(175),
	1043: uint16(1),
	1044: uint16(anon_sym_RBRACE),
	1045: uint16(51),
	1046: uint16(1),
	1047: uint16(aux_sym_dictionary_repeat1),
	1048: uint16(4),
	1049: uint16(3),
	1050: uint16(1),
	1051: uint16(sym_comment),
	1052: uint16(177),
	1053: uint16(1),
	1054: uint16(anon_sym_RBRACK),
	1055: uint16(179),
	1056: uint16(1),
	1057: uint16(anon_sym_COMMA),
	1058: uint16(48),
	1059: uint16(1),
	1060: uint16(aux_sym_array_repeat1),
	1061: uint16(4),
	1062: uint16(3),
	1063: uint16(1),
	1064: uint16(sym_comment),
	1065: uint16(179),
	1066: uint16(1),
	1067: uint16(anon_sym_COMMA),
	1068: uint16(181),
	1069: uint16(1),
	1070: uint16(anon_sym_RBRACK),
	1071: uint16(56),
	1072: uint16(1),
	1073: uint16(aux_sym_array_repeat1),
	1074: uint16(4),
	1075: uint16(3),
	1076: uint16(1),
	1077: uint16(sym_comment),
	1078: uint16(155),
	1079: uint16(1),
	1080: uint16(anon_sym_COMMA),
	1081: uint16(183),
	1082: uint16(1),
	1083: uint16(anon_sym_RPAREN),
	1084: uint16(57),
	1085: uint16(1),
	1086: uint16(aux_sym_arguments_repeat1),
	1087: uint16(3),
	1088: uint16(3),
	1089: uint16(1),
	1090: uint16(sym_comment),
	1091: uint16(153),
	1092: uint16(1),
	1093: uint16(anon_sym_COLON),
	1094: uint16(185),
	1095: uint16(2),
	1096: uint16(anon_sym_COMMA),
	1097: uint16(anon_sym_RPAREN),
	1098: uint16(4),
	1099: uint16(3),
	1100: uint16(1),
	1101: uint16(sym_comment),
	1102: uint16(179),
	1103: uint16(1),
	1104: uint16(anon_sym_COMMA),
	1105: uint16(187),
	1106: uint16(1),
	1107: uint16(anon_sym_RBRACK),
	1108: uint16(48),
	1109: uint16(1),
	1110: uint16(aux_sym_array_repeat1),
	1111: uint16(4),
	1112: uint16(3),
	1113: uint16(1),
	1114: uint16(sym_comment),
	1115: uint16(185),
	1116: uint16(1),
	1117: uint16(anon_sym_RPAREN),
	1118: uint16(189),
	1119: uint16(1),
	1120: uint16(anon_sym_COMMA),
	1121: uint16(57),
	1122: uint16(1),
	1123: uint16(aux_sym_arguments_repeat1),
	1124: uint16(4),
	1125: uint16(3),
	1126: uint16(1),
	1127: uint16(sym_comment),
	1128: uint16(192),
	1129: uint16(1),
	1130: uint16(anon_sym_COMMA),
	1131: uint16(194),
	1132: uint16(1),
	1133: uint16(anon_sym_RBRACE),
	1134: uint16(59),
	1135: uint16(1),
	1136: uint16(aux_sym_dictionary_repeat1),
	1137: uint16(4),
	1138: uint16(3),
	1139: uint16(1),
	1140: uint16(sym_comment),
	1141: uint16(192),
	1142: uint16(1),
	1143: uint16(anon_sym_COMMA),
	1144: uint16(196),
	1145: uint16(1),
	1146: uint16(anon_sym_RBRACE),
	1147: uint16(51),
	1148: uint16(1),
	1149: uint16(aux_sym_dictionary_repeat1),
	1150: uint16(4),
	1151: uint16(3),
	1152: uint16(1),
	1153: uint16(sym_comment),
	1154: uint16(179),
	1155: uint16(1),
	1156: uint16(anon_sym_COMMA),
	1157: uint16(198),
	1158: uint16(1),
	1159: uint16(anon_sym_RBRACK),
	1160: uint16(63),
	1161: uint16(1),
	1162: uint16(aux_sym_array_repeat1),
	1163: uint16(4),
	1164: uint16(3),
	1165: uint16(1),
	1166: uint16(sym_comment),
	1167: uint16(192),
	1168: uint16(1),
	1169: uint16(anon_sym_COMMA),
	1170: uint16(200),
	1171: uint16(1),
	1172: uint16(anon_sym_RBRACE),
	1173: uint16(64),
	1174: uint16(1),
	1175: uint16(aux_sym_dictionary_repeat1),
	1176: uint16(4),
	1177: uint16(3),
	1178: uint16(1),
	1179: uint16(sym_comment),
	1180: uint16(179),
	1181: uint16(1),
	1182: uint16(anon_sym_COMMA),
	1183: uint16(202),
	1184: uint16(1),
	1185: uint16(anon_sym_RBRACK),
	1186: uint16(52),
	1187: uint16(1),
	1188: uint16(aux_sym_array_repeat1),
	1189: uint16(4),
	1190: uint16(3),
	1191: uint16(1),
	1192: uint16(sym_comment),
	1193: uint16(179),
	1194: uint16(1),
	1195: uint16(anon_sym_COMMA),
	1196: uint16(204),
	1197: uint16(1),
	1198: uint16(anon_sym_RBRACK),
	1199: uint16(48),
	1200: uint16(1),
	1201: uint16(aux_sym_array_repeat1),
	1202: uint16(4),
	1203: uint16(3),
	1204: uint16(1),
	1205: uint16(sym_comment),
	1206: uint16(192),
	1207: uint16(1),
	1208: uint16(anon_sym_COMMA),
	1209: uint16(206),
	1210: uint16(1),
	1211: uint16(anon_sym_RBRACE),
	1212: uint16(51),
	1213: uint16(1),
	1214: uint16(aux_sym_dictionary_repeat1),
	1215: uint16(4),
	1216: uint16(3),
	1217: uint16(1),
	1218: uint16(sym_comment),
	1219: uint16(155),
	1220: uint16(1),
	1221: uint16(anon_sym_COMMA),
	1222: uint16(159),
	1223: uint16(1),
	1224: uint16(anon_sym_RPAREN),
	1225: uint16(47),
	1226: uint16(1),
	1227: uint16(aux_sym_arguments_repeat1),
	1228: uint16(2),
	1229: uint16(3),
	1230: uint16(1),
	1231: uint16(sym_comment),
	1232: uint16(111),
	1233: uint16(2),
	1234: uint16(sym__identifier),
	1235: uint16(anon_sym_RBRACK),
	1236: uint16(2),
	1237: uint16(3),
	1238: uint16(1),
	1239: uint16(sym_comment),
	1240: uint16(163),
	1241: uint16(2),
	1242: uint16(anon_sym_RBRACK),
	1243: uint16(anon_sym_COMMA),
	1244: uint16(2),
	1245: uint16(3),
	1246: uint16(1),
	1247: uint16(sym_comment),
	1248: uint16(99),
	1249: uint16(2),
	1250: uint16(sym__identifier),
	1251: uint16(anon_sym_RBRACK),
	1252: uint16(2),
	1253: uint16(3),
	1254: uint16(1),
	1255: uint16(sym_comment),
	1256: uint16(93),
	1257: uint16(2),
	1258: uint16(sym__identifier),
	1259: uint16(anon_sym_RBRACK),
	1260: uint16(2),
	1261: uint16(3),
	1262: uint16(1),
	1263: uint16(sym_comment),
	1264: uint16(103),
	1265: uint16(2),
	1266: uint16(sym__identifier),
	1267: uint16(anon_sym_RBRACK),
	1268: uint16(2),
	1269: uint16(3),
	1270: uint16(1),
	1271: uint16(sym_comment),
	1272: uint16(109),
	1273: uint16(2),
	1274: uint16(sym__identifier),
	1275: uint16(anon_sym_RBRACK),
	1276: uint16(2),
	1277: uint16(3),
	1278: uint16(1),
	1279: uint16(sym_comment),
	1280: uint16(91),
	1281: uint16(2),
	1282: uint16(sym__identifier),
	1283: uint16(anon_sym_RBRACK),
	1284: uint16(2),
	1285: uint16(3),
	1286: uint16(1),
	1287: uint16(sym_comment),
	1288: uint16(113),
	1289: uint16(2),
	1290: uint16(sym__identifier),
	1291: uint16(anon_sym_RBRACK),
	1292: uint16(2),
	1293: uint16(3),
	1294: uint16(1),
	1295: uint16(sym_comment),
	1296: uint16(107),
	1297: uint16(2),
	1298: uint16(sym__identifier),
	1299: uint16(anon_sym_RBRACK),
	1300: uint16(3),
	1301: uint16(3),
	1302: uint16(1),
	1303: uint16(sym_comment),
	1304: uint16(87),
	1305: uint16(1),
	1306: uint16(anon_sym_LPAREN),
	1307: uint16(26),
	1308: uint16(1),
	1309: uint16(sym_arguments),
	1310: uint16(2),
	1311: uint16(3),
	1312: uint16(1),
	1313: uint16(sym_comment),
	1314: uint16(185),
	1315: uint16(2),
	1316: uint16(anon_sym_COMMA),
	1317: uint16(anon_sym_RPAREN),
	1318: uint16(3),
	1319: uint16(3),
	1320: uint16(1),
	1321: uint16(sym_comment),
	1322: uint16(123),
	1323: uint16(1),
	1324: uint16(sym__identifier),
	1325: uint16(34),
	1326: uint16(1),
	1327: uint16(sym_identifier),
	1328: uint16(2),
	1329: uint16(3),
	1330: uint16(1),
	1331: uint16(sym_comment),
	1332: uint16(208),
	1333: uint16(2),
	1334: uint16(sym__identifier),
	1335: uint16(anon_sym_RBRACK),
	1336: uint16(2),
	1337: uint16(3),
	1338: uint16(1),
	1339: uint16(sym_comment),
	1340: uint16(175),
	1341: uint16(2),
	1342: uint16(anon_sym_COMMA),
	1343: uint16(anon_sym_RBRACE),
	1344: uint16(3),
	1345: uint16(3),
	1346: uint16(1),
	1347: uint16(sym_comment),
	1348: uint16(117),
	1349: uint16(1),
	1350: uint16(anon_sym_LPAREN),
	1351: uint16(71),
	1352: uint16(1),
	1353: uint16(sym_arguments),
	1354: uint16(2),
	1355: uint16(3),
	1356: uint16(1),
	1357: uint16(sym_comment),
	1358: uint16(105),
	1359: uint16(2),
	1360: uint16(sym__identifier),
	1361: uint16(anon_sym_RBRACK),
	1362: uint16(2),
	1363: uint16(3),
	1364: uint16(1),
	1365: uint16(sym_comment),
	1366: uint16(95),
	1367: uint16(2),
	1368: uint16(sym__identifier),
	1369: uint16(anon_sym_RBRACK),
	1370: uint16(2),
	1371: uint16(3),
	1372: uint16(1),
	1373: uint16(sym_comment),
	1374: uint16(97),
	1375: uint16(2),
	1376: uint16(sym__identifier),
	1377: uint16(anon_sym_RBRACK),
	1378: uint16(2),
	1379: uint16(3),
	1380: uint16(1),
	1381: uint16(sym_comment),
	1382: uint16(101),
	1383: uint16(2),
	1384: uint16(sym__identifier),
	1385: uint16(anon_sym_RBRACK),
	1386: uint16(2),
	1387: uint16(3),
	1388: uint16(1),
	1389: uint16(sym_comment),
	1390: uint16(153),
	1391: uint16(1),
	1392: uint16(anon_sym_COLON),
	1393: uint16(2),
	1394: uint16(3),
	1395: uint16(1),
	1396: uint16(sym_comment),
	1397: uint16(210),
	1398: uint16(1),
	1399: uint16(anon_sym_LPAREN),
	1400: uint16(2),
	1401: uint16(3),
	1402: uint16(1),
	1403: uint16(sym_comment),
	1404: uint16(212),
	1405: uint16(1),
	1406: uint16(sym_string),
	1407: uint16(2),
	1408: uint16(3),
	1409: uint16(1),
	1410: uint16(sym_comment),
	1411: uint16(214),
	1412: uint16(1),
	1413: uint16(sym_string),
	1414: uint16(2),
	1415: uint16(3),
	1416: uint16(1),
	1417: uint16(sym_comment),
	1418: uint16(216),
	1419: uint16(1),
	1420: uint16(anon_sym_LPAREN),
	1421: uint16(2),
	1422: uint16(3),
	1423: uint16(1),
	1424: uint16(sym_comment),
	1425: uint16(218),
	1426: uint16(1),
	1427: uint16(anon_sym_EQ),
	1428: uint16(2),
	1429: uint16(3),
	1430: uint16(1),
	1431: uint16(sym_comment),
	1432: uint16(220),
	1433: uint16(1),
	1434: uint16(anon_sym_EQ),
	1435: uint16(2),
	1436: uint16(3),
	1437: uint16(1),
	1438: uint16(sym_comment),
	1439: uint16(222),
	1440: uint16(1),
}

var ts_small_parse_table_map = [91]uint32_t{
	1:  uint32(42),
	2:  uint32(84),
	3:  uint32(126),
	4:  uint32(168),
	5:  uint32(207),
	6:  uint32(246),
	7:  uint32(285),
	8:  uint32(324),
	9:  uint32(360),
	10: uint32(396),
	11: uint32(432),
	12: uint32(468),
	13: uint32(504),
	14: uint32(529),
	15: uint32(544),
	16: uint32(558),
	17: uint32(572),
	18: uint32(586),
	19: uint32(600),
	20: uint32(614),
	21: uint32(628),
	22: uint32(642),
	23: uint32(656),
	24: uint32(670),
	25: uint32(684),
	26: uint32(698),
	27: uint32(712),
	28: uint32(733),
	29: uint32(753),
	30: uint32(768),
	31: uint32(783),
	32: uint32(800),
	33: uint32(817),
	34: uint32(832),
	35: uint32(847),
	36: uint32(864),
	37: uint32(879),
	38: uint32(894),
	39: uint32(908),
	40: uint32(922),
	41: uint32(932),
	42: uint32(946),
	43: uint32(962),
	44: uint32(978),
	45: uint32(991),
	46: uint32(1004),
	47: uint32(1017),
	48: uint32(1026),
	49: uint32(1035),
	50: uint32(1048),
	51: uint32(1061),
	52: uint32(1074),
	53: uint32(1087),
	54: uint32(1098),
	55: uint32(1111),
	56: uint32(1124),
	57: uint32(1137),
	58: uint32(1150),
	59: uint32(1163),
	60: uint32(1176),
	61: uint32(1189),
	62: uint32(1202),
	63: uint32(1215),
	64: uint32(1228),
	65: uint32(1236),
	66: uint32(1244),
	67: uint32(1252),
	68: uint32(1260),
	69: uint32(1268),
	70: uint32(1276),
	71: uint32(1284),
	72: uint32(1292),
	73: uint32(1300),
	74: uint32(1310),
	75: uint32(1318),
	76: uint32(1328),
	77: uint32(1336),
	78: uint32(1344),
	79: uint32(1354),
	80: uint32(1362),
	81: uint32(1370),
	82: uint32(1378),
	83: uint32(1386),
	84: uint32(1393),
	85: uint32(1400),
	86: uint32(1407),
	87: uint32(1414),
	88: uint32(1421),
	89: uint32(1428),
	90: uint32(1435),
}

var ts_parse_actions = [224]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_resource),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(88)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(70)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(55)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(55)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(62)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(67)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(67)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_identifier),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_array),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_dictionary),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_array),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_dictionary),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_array),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	102: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_constructor),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_arguments),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_string_name),
		Fproduction_id: uint16(1),
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
		Fsymbol:      uint16(sym_arguments),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_constructor),
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
		Fsymbol:      uint16(sym_arguments),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_dictionary),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_resource),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_section),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_section),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__properties),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__properties),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_resource_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_resource_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym__attributes),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym__attributes),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_section),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	147: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_resource_repeat2),
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
		Fsymbol:      uint16(aux_sym_resource_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_resource),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(66)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_array_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_array_repeat1),
	})))),
	167: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_pair),
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
		Fsymbol:      uint16(sym_property),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_dictionary_repeat1),
	})))),
	174: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(aux_sym_dictionary_repeat1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(89)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(27)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(aux_sym_arguments_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(8)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(28)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(68)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_attribute),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__type_args),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__type_args),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
	223: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token_string = 0

var ts_external_scanner_symbol_map = [1]TSSymbol{
	0: uint16(sym_string),
}

var ts_external_scanner_states = [2][1]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_godot_resource(tls *libc.TLS) (r uintptr) {
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
	Fkeyword_capture_token:     uint16(sym__identifier),
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
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_godot_resource_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_godot_resource_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_godot_resource_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_godot_resource_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_godot_resource_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00_identifier\x00comment\x00true\x00false\x00null\x00float\x00integer\x00&\x00[\x00]\x00=\x00path\x00:\x00{\x00,\x00}\x00(\x00)\x00string\x00resource\x00identifier\x00string_name\x00_value\x00section\x00_attributes\x00attribute\x00_properties\x00property\x00pair\x00dictionary\x00array\x00arguments\x00_type_args\x00constructor\x00resource_repeat1\x00resource_repeat2\x00dictionary_repeat1\x00array_repeat1\x00arguments_repeat1\x00"
