// Code generated for linux/arm64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-comment/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-comment -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && arm64

package grammar_comment

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 0
const BUFSIZ = 1024
const CHAR_CARRIAGE_RETURN = 13
const CHAR_EOF = 0
const CHAR_FORM_FEED = 12
const CHAR_NEWLINE = 10
const CHAR_SPACE = 32
const CHAR_TAB = 9
const CHAR_VERTICAL_TAB = 11
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXTERNAL_TOKEN_COUNT = 2
const FIELD_COUNT = 0
const FILENAME_MAX = 4096
const FOPEN_MAX = 1000
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
const LARGE_STATE_COUNT = 9
const L_ctermid = 20
const L_cuserid = 20
const L_tmpnam = 20
const MAX_ALIAS_SEQUENCE_LENGTH = 3
const MAX_RESERVED_WORD_SET_SIZE = 0
const PRODUCTION_ID_COUNT = 1
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const P_tmpdir = "/tmp"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 16
const SUPERTYPE_COUNT = 0
const SYMBOL_COUNT = 35
const TMP_MAX = 10000
const TOKEN_COUNT = 27
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
const _IOFBF = 0
const _IOLBF = 1
const _IONBF = 2
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

func is_upper(tls *libc.TLS, c int32_t) (r uint8) {
	var lower, upper int32_t
	_, _ = lower, upper
	upper = int32(65)
	lower = int32(90)
	return libc.BoolUint8(c >= upper && c <= lower)
}

func is_digit(tls *libc.TLS, c int32_t) (r uint8) {
	var lower, upper int32_t
	_, _ = lower, upper
	upper = int32(48)
	lower = int32(57)
	return libc.BoolUint8(c >= upper && c <= lower)
}

func is_newline(tls *libc.TLS, c int32_t) (r uint8) {
	var i, length int32
	var newline_chars [3]int32_t
	_, _, _ = i, length, newline_chars
	newline_chars = [3]int32_t{
		1: int32(CHAR_NEWLINE),
		2: int32(CHAR_CARRIAGE_RETURN),
	}
	length = libc.Int32FromUint64(libc.Uint64FromInt64(12) / libc.Uint64FromInt64(4))
	i = 0
	for {
		if !(i < length) {
			break
		}
		if c == newline_chars[i] {
			return libc.BoolUint8(true1 != 0)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return libc.BoolUint8(false1 != 0)
}

func is_space(tls *libc.TLS, c int32_t) (r uint8) {
	var i, length int32
	var is_space_char uint8
	var space_chars [4]int32_t
	_, _, _, _ = i, is_space_char, length, space_chars
	space_chars = [4]int32_t{
		0: int32(' '),
		1: int32('\f'),
		2: int32('\t'),
		3: int32('\v'),
	}
	length = libc.Int32FromUint64(libc.Uint64FromInt64(16) / libc.Uint64FromInt64(4))
	is_space_char = libc.BoolUint8(false1 != 0)
	i = 0
	for {
		if !(i < length) {
			break
		}
		if c == space_chars[i] {
			is_space_char = libc.BoolUint8(true1 != 0)
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return libc.BoolUint8(is_space_char != 0 || is_newline(tls, c) != 0)
}

// C documentation
//
//	/// Check if the character is allowed inside the name.
func is_internal_char(tls *libc.TLS, c int32_t) (r uint8) {
	var i, length int32
	var valid_chars [2]int32_t
	_, _, _ = i, length, valid_chars
	valid_chars = [2]int32_t{
		0: int32('-'),
		1: int32('_'),
	}
	length = libc.Int32FromUint64(libc.Uint64FromInt64(8) / libc.Uint64FromInt64(4))
	i = 0
	for {
		if !(i < length) {
			break
		}
		if c == valid_chars[i] {
			return libc.BoolUint8(true1 != 0)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return libc.BoolUint8(false1 != 0)
}

type TokenType = int32

const T_TAGNAME = 0
const T_INVALID_TOKEN = 1

type ssize_t = int64

type off_t = int64

type va_list = uintptr

type __isoc_va_list = uintptr

type fpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]uint8
}

type _G_fpos64_t = fpos_t

type cookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type _IO_cookie_io_functions_t = cookie_io_functions_t

// C documentation
//
//	/// Parse the name of the tag.
//	///
//	/// They can be of the form:
//	/// - TODO:
//	/// - TODO: text
//	/// - TODO(stsewd):
//	/// - TODO(stsewd): text
//	/// - TODO (stsewd): text
func parse_tagname(tls *libc.TLS, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var previous int32_t
	var user_length int32
	_, _ = previous, user_length
	if !(is_upper(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0) || !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(T_TAGNAME))) != 0) {
		return libc.BoolUint8(false1 != 0)
	}
	previous = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	for is_upper(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0 || is_digit(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0 || is_internal_char(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0 {
		previous = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	}
	// The tag name ends here.
	// But we keep parsing to see if it's a valid tag name.
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
	// It can't end with an internal char.
	if is_internal_char(tls, previous) != 0 {
		return libc.BoolUint8(false1 != 0)
	}
	// For the user component this is `\s*(`.
	// We don't parse that part, we just need to be sure it ends with `:\s`.
	if is_space(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0 && !(is_newline(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0) || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('(') {
		// Skip white spaces.
		for is_space(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0 && !(is_newline(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0) {
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
		}
		// Checking aperture.
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('(') {
			return libc.BoolUint8(false1 != 0)
		}
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
		// Checking closure.
		user_length = 0
		for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32(')') {
			if is_newline(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0 {
				return libc.BoolUint8(false1 != 0)
			}
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
			user_length = user_length + 1
		}
		if user_length <= 0 {
			return libc.BoolUint8(false1 != 0)
		}
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	}
	// It should end with `:`...
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32(':') {
		return libc.BoolUint8(false1 != 0)
	}
	// ... and be followed by one space.
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
	if !(is_space(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0) {
		return libc.BoolUint8(false1 != 0)
	}
	(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(T_TAGNAME)
	return libc.BoolUint8(true1 != 0)
}

func parse(tls *libc.TLS, lexer uintptr, valid_symbols uintptr) (r uint8) {
	// If all valid symbols are true, tree-sitter is in correction mode.
	// We don't want to parse anything in that case.
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(T_INVALID_TOKEN))) != 0 {
		return libc.BoolUint8(false1 != 0)
	}
	if is_upper(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0 && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(T_TAGNAME))) != 0 {
		return parse_tagname(tls, lexer, valid_symbols)
	}
	return libc.BoolUint8(false1 != 0)
}

func tree_sitter_comment_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_comment_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
}

func tree_sitter_comment_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_comment_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
}

func tree_sitter_comment_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	return parse(tls, lexer, valid_symbols)
}

/* Automatically generated by tree-sitter v0.25.3 (2a835ee029dca1c325e6f1c01dbce40396f6123e) */

type ts_symbol_identifiers = int32

const anon_sym_COLON = 1
const anon_sym_LPAREN = 2
const aux_sym__user_token1 = 3
const anon_sym_RPAREN = 4
const aux_sym__full_uri_token1 = 5
const sym_uri = 6
const aux_sym__text_token1 = 7
const anon_sym_SLASH = 8
const anon_sym_SQUOTE = 9
const anon_sym_DQUOTE = 10
const anon_sym_BQUOTE = 11
const anon_sym_LT = 12
const anon_sym_LBRACK = 13
const anon_sym_LBRACE = 14
const anon_sym_DOT = 15
const anon_sym_COMMA = 16
const anon_sym_SEMI = 17
const anon_sym_BANG = 18
const anon_sym_QMARK = 19
const anon_sym_BSLASH = 20
const anon_sym_RBRACE = 21
const anon_sym_RBRACK = 22
const anon_sym_GT = 23
const anon_sym_DASH = 24
const sym_name = 25
const sym_invalid_token = 26
const sym_source = 27
const sym_tag = 28
const sym__user = 29
const sym__full_uri = 30
const sym__text = 31
const sym__stop_char = 32
const sym__end_char = 33
const aux_sym_source_repeat1 = 34

var ts_symbol_names = [35]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 6,
	3:  __ccgo_ts + 8,
	4:  __ccgo_ts + 13,
	5:  __ccgo_ts + 15,
	6:  __ccgo_ts + 32,
	7:  __ccgo_ts + 36,
	8:  __ccgo_ts + 49,
	9:  __ccgo_ts + 51,
	10: __ccgo_ts + 53,
	11: __ccgo_ts + 55,
	12: __ccgo_ts + 57,
	13: __ccgo_ts + 59,
	14: __ccgo_ts + 61,
	15: __ccgo_ts + 63,
	16: __ccgo_ts + 65,
	17: __ccgo_ts + 67,
	18: __ccgo_ts + 69,
	19: __ccgo_ts + 71,
	20: __ccgo_ts + 73,
	21: __ccgo_ts + 75,
	22: __ccgo_ts + 77,
	23: __ccgo_ts + 79,
	24: __ccgo_ts + 81,
	25: __ccgo_ts + 83,
	26: __ccgo_ts + 88,
	27: __ccgo_ts + 102,
	28: __ccgo_ts + 109,
	29: __ccgo_ts + 113,
	30: __ccgo_ts + 119,
	31: __ccgo_ts + 129,
	32: __ccgo_ts + 134,
	33: __ccgo_ts + 129,
	34: __ccgo_ts + 145,
}

var ts_symbol_map = [35]TSSymbol{
	1:  uint16(anon_sym_COLON),
	2:  uint16(anon_sym_LPAREN),
	3:  uint16(aux_sym__user_token1),
	4:  uint16(anon_sym_RPAREN),
	5:  uint16(aux_sym__full_uri_token1),
	6:  uint16(sym_uri),
	7:  uint16(aux_sym__text_token1),
	8:  uint16(anon_sym_SLASH),
	9:  uint16(anon_sym_SQUOTE),
	10: uint16(anon_sym_DQUOTE),
	11: uint16(anon_sym_BQUOTE),
	12: uint16(anon_sym_LT),
	13: uint16(anon_sym_LBRACK),
	14: uint16(anon_sym_LBRACE),
	15: uint16(anon_sym_DOT),
	16: uint16(anon_sym_COMMA),
	17: uint16(anon_sym_SEMI),
	18: uint16(anon_sym_BANG),
	19: uint16(anon_sym_QMARK),
	20: uint16(anon_sym_BSLASH),
	21: uint16(anon_sym_RBRACE),
	22: uint16(anon_sym_RBRACK),
	23: uint16(anon_sym_GT),
	24: uint16(anon_sym_DASH),
	25: uint16(sym_name),
	26: uint16(sym_invalid_token),
	27: uint16(sym_source),
	28: uint16(sym_tag),
	29: uint16(sym__user),
	30: uint16(sym__full_uri),
	31: uint16(sym__text),
	32: uint16(sym__stop_char),
	33: uint16(sym__text),
	34: uint16(aux_sym_source_repeat1),
}

var ts_symbol_metadata = [35]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	4: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	5: {},
	6: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	7: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	26: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	27: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	28: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	29: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	30: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	31: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	32: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	33: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	34: {},
}

var ts_alias_sequences = [1][3]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [16]TSStateId{
	1:  uint16(1),
	2:  uint16(2),
	3:  uint16(3),
	4:  uint16(4),
	5:  uint16(5),
	6:  uint16(6),
	7:  uint16(7),
	8:  uint16(8),
	9:  uint16(9),
	10: uint16(10),
	11: uint16(11),
	12: uint16(12),
	13: uint16(13),
	14: uint16(14),
	15: uint16(15),
}

var sym_uri_character_set_1 = [12]TSCharacterRange{
	0: {
		Fend: int32(0x08),
	},
	1: {
		Fstart: int32(0x0e),
		Fend:   int32(0x1f),
	},
	2: {
		Fstart: int32('#'),
		Fend:   int32('&'),
	},
	3: {
		Fstart: int32('('),
		Fend:   int32('('),
	},
	4: {
		Fstart: int32('*'),
		Fend:   int32('+'),
	},
	5: {
		Fstart: int32('-'),
		Fend:   int32('-'),
	},
	6: {
		Fstart: int32('/'),
		Fend:   int32('9'),
	},
	7: {
		Fstart: int32('<'),
		Fend:   int32('='),
	},
	8: {
		Fstart: int32('@'),
		Fend:   int32('['),
	},
	9: {
		Fstart: int32('^'),
		Fend:   int32('_'),
	},
	10: {
		Fstart: int32('a'),
		Fend:   int32('|'),
	},
	11: {
		Fstart: int32('~'),
		Fend:   int32(0x10ffff),
	},
}

var aux_sym__text_token1_character_set_1 = [11]TSCharacterRange{
	0: {
		Fend: int32(0x08),
	},
	1: {
		Fstart: int32(0x0e),
		Fend:   int32(0x1f),
	},
	2: {
		Fstart: int32('#'),
		Fend:   int32('&'),
	},
	3: {
		Fstart: int32('*'),
		Fend:   int32('+'),
	},
	4: {
		Fstart: int32('0'),
		Fend:   int32('9'),
	},
	5: {
		Fstart: int32('='),
		Fend:   int32('='),
	},
	6: {
		Fstart: int32('@'),
		Fend:   int32('Z'),
	},
	7: {
		Fstart: int32('^'),
		Fend:   int32('_'),
	},
	8: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	9: {
		Fstart: int32('|'),
		Fend:   int32('|'),
	},
	10: {
		Fstart: int32('~'),
		Fend:   int32(0x10ffff),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v5 uint8
	var half_size, i, i1, i2, index, mid_index, size uint32_t
	var lookahead1, v4 int32_t
	var range_token, range_token1, v3 uintptr
	var v7 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i2, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v3, v4, v5, v7
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
			state = uint16(6)
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
			state = uint16(11)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead1 == int32('/') {
			state = uint16(4)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead1 == int32('/') {
			state = uint16(1)
			goto next_state
		}
		return result
	case int32(3):
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(9)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('(') && lookahead1 != int32(')') {
			state = uint16(9)
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
			goto _2
		_2:
			;
			i1 = i1 + uint32(2)
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && (lookahead1 < int32(' ') || int32('"') < lookahead1) {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(5):
		if v7 = !(eof != 0); v7 {
			v3 = uintptr(unsafe.Pointer(&sym_uri_character_set_1))
			v4 = lookahead1
			index = uint32(0)
			size = uint32(12) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v3 + uintptr(mid_index)*8
				if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v5 = libc.BoolUint8(true1 != 0)
					goto _6
				} else {
					if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v3 + uintptr(index)*8
			v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _6
		_6:
		}
		if v7 && v5 != 0 {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(6):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(7):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(8):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(9):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__user_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('(') && lookahead1 != int32(')') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(10):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(11):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__full_uri_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(12):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_uri)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _8
		_8:
			;
			i2 = i2 + uint32(2)
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && (lookahead1 < int32(' ') || int32('"') < lookahead1) {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(13):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__text_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') {
			state = uint16(2)
			goto next_state
		}
		if lookahead1 == int32('s') {
			state = uint16(14)
			goto next_state
		}
		if v7 = !(eof != 0); v7 {
			v3 = uintptr(unsafe.Pointer(&aux_sym__text_token1_character_set_1))
			v4 = lookahead1
			index = uint32(0)
			size = uint32(11) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v3 + uintptr(mid_index)*8
				if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v5 = libc.BoolUint8(true1 != 0)
					goto _12
				} else {
					if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v3 + uintptr(index)*8
			v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _12
		_12:
		}
		if v7 && v5 != 0 {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(14):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__text_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') {
			state = uint16(2)
			goto next_state
		}
		if v7 = !(eof != 0); v7 {
			v3 = uintptr(unsafe.Pointer(&aux_sym__text_token1_character_set_1))
			v4 = lookahead1
			index = uint32(0)
			size = uint32(11) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v3 + uintptr(mid_index)*8
				if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v5 = libc.BoolUint8(true1 != 0)
					goto _17
				} else {
					if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v3 + uintptr(index)*8
			v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _17
		_17:
		}
		if v7 && v5 != 0 {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(15):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__text_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('p') {
			state = uint16(13)
			goto next_state
		}
		if v7 = !(eof != 0); v7 {
			v3 = uintptr(unsafe.Pointer(&aux_sym__text_token1_character_set_1))
			v4 = lookahead1
			index = uint32(0)
			size = uint32(11) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v3 + uintptr(mid_index)*8
				if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v5 = libc.BoolUint8(true1 != 0)
					goto _22
				} else {
					if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v3 + uintptr(index)*8
			v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _22
		_22:
		}
		if v7 && v5 != 0 {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(16):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__text_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('t') {
			state = uint16(15)
			goto next_state
		}
		if v7 = !(eof != 0); v7 {
			v3 = uintptr(unsafe.Pointer(&aux_sym__text_token1_character_set_1))
			v4 = lookahead1
			index = uint32(0)
			size = uint32(11) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v3 + uintptr(mid_index)*8
				if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v5 = libc.BoolUint8(true1 != 0)
					goto _27
				} else {
					if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v3 + uintptr(index)*8
			v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _27
		_27:
		}
		if v7 && v5 != 0 {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(17):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__text_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('t') {
			state = uint16(16)
			goto next_state
		}
		if v7 = !(eof != 0); v7 {
			v3 = uintptr(unsafe.Pointer(&aux_sym__text_token1_character_set_1))
			v4 = lookahead1
			index = uint32(0)
			size = uint32(11) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v3 + uintptr(mid_index)*8
				if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v5 = libc.BoolUint8(true1 != 0)
					goto _32
				} else {
					if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v3 + uintptr(index)*8
			v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _32
		_32:
		}
		if v7 && v5 != 0 {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(18):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__text_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if v7 = !(eof != 0); v7 {
			v3 = uintptr(unsafe.Pointer(&aux_sym__text_token1_character_set_1))
			v4 = lookahead1
			index = uint32(0)
			size = uint32(11) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v3 + uintptr(mid_index)*8
				if v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v5 = libc.BoolUint8(true1 != 0)
					goto _37
				} else {
					if v4 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v3 + uintptr(index)*8
			v5 = libc.BoolUint8(v4 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v4 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _37
		_37:
		}
		if v7 && v5 != 0 {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(19):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(20):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(22):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(23):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(24):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [42]uint16_t{
	0:  uint16('!'),
	1:  uint16(29),
	2:  uint16('"'),
	3:  uint16(21),
	4:  uint16('\''),
	5:  uint16(20),
	6:  uint16('('),
	7:  uint16(8),
	8:  uint16(')'),
	9:  uint16(10),
	10: uint16(','),
	11: uint16(27),
	12: uint16('-'),
	13: uint16(35),
	14: uint16('.'),
	15: uint16(26),
	16: uint16('/'),
	17: uint16(19),
	18: uint16(':'),
	19: uint16(7),
	20: uint16(';'),
	21: uint16(28),
	22: uint16('<'),
	23: uint16(23),
	24: uint16('>'),
	25: uint16(34),
	26: uint16('?'),
	27: uint16(30),
	28: uint16('['),
	29: uint16(24),
	30: uint16('\\'),
	31: uint16(31),
	32: uint16(']'),
	33: uint16(33),
	34: uint16('`'),
	35: uint16(22),
	36: uint16('h'),
	37: uint16(17),
	38: uint16('{'),
	39: uint16(25),
	40: uint16('}'),
	41: uint16(32),
}

var map_token1 = [28]uint16_t{
	0:  uint16('!'),
	1:  uint16(5),
	2:  uint16('"'),
	3:  uint16(5),
	4:  uint16('\''),
	5:  uint16(5),
	6:  uint16(')'),
	7:  uint16(5),
	8:  uint16(','),
	9:  uint16(5),
	10: uint16('.'),
	11: uint16(5),
	12: uint16(':'),
	13: uint16(5),
	14: uint16(';'),
	15: uint16(5),
	16: uint16('>'),
	17: uint16(5),
	18: uint16('?'),
	19: uint16(5),
	20: uint16('\\'),
	21: uint16(5),
	22: uint16(']'),
	23: uint16(5),
	24: uint16('`'),
	25: uint16(5),
	26: uint16('}'),
	27: uint16(5),
}

var map_token2 = [28]uint16_t{
	0:  uint16('!'),
	1:  uint16(5),
	2:  uint16('"'),
	3:  uint16(5),
	4:  uint16('\''),
	5:  uint16(5),
	6:  uint16(')'),
	7:  uint16(5),
	8:  uint16(','),
	9:  uint16(5),
	10: uint16('.'),
	11: uint16(5),
	12: uint16(':'),
	13: uint16(5),
	14: uint16(';'),
	15: uint16(5),
	16: uint16('>'),
	17: uint16(5),
	18: uint16('?'),
	19: uint16(5),
	20: uint16('\\'),
	21: uint16(5),
	22: uint16(']'),
	23: uint16(5),
	24: uint16('`'),
	25: uint16(5),
	26: uint16('}'),
	27: uint16(5),
}

var ts_lex_modes = [16]TSLexerMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Fexternal_lex_state: uint16(2),
	},
	2: {
		Fexternal_lex_state: uint16(2),
	},
	3: {
		Fexternal_lex_state: uint16(2),
	},
	4: {
		Fexternal_lex_state: uint16(2),
	},
	5: {
		Fexternal_lex_state: uint16(2),
	},
	6: {
		Fexternal_lex_state: uint16(2),
	},
	7: {
		Fexternal_lex_state: uint16(2),
	},
	8: {
		Fexternal_lex_state: uint16(2),
	},
	9:  {},
	10: {},
	11: {},
	12: {
		Flex_state: uint16(3),
	},
	13: {},
	14: {},
	15: {},
}

var ts_parse_table = [9][35]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
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
		20: uint16(1),
		21: uint16(1),
		22: uint16(1),
		23: uint16(1),
		24: uint16(1),
		25: uint16(1),
		26: uint16(1),
	},
	1: {
		0:  uint16(5),
		1:  uint16(7),
		2:  uint16(7),
		4:  uint16(7),
		5:  uint16(3),
		6:  uint16(9),
		7:  uint16(11),
		8:  uint16(7),
		9:  uint16(7),
		10: uint16(7),
		11: uint16(7),
		12: uint16(7),
		13: uint16(7),
		14: uint16(7),
		15: uint16(7),
		16: uint16(7),
		17: uint16(7),
		18: uint16(7),
		19: uint16(7),
		20: uint16(7),
		21: uint16(7),
		22: uint16(7),
		23: uint16(7),
		24: uint16(7),
		25: uint16(13),
		27: uint16(11),
		28: uint16(2),
		30: uint16(2),
		31: uint16(2),
		32: uint16(4),
		34: uint16(2),
	},
	2: {
		0:  uint16(15),
		1:  uint16(7),
		2:  uint16(7),
		4:  uint16(7),
		5:  uint16(3),
		6:  uint16(9),
		7:  uint16(11),
		8:  uint16(7),
		9:  uint16(7),
		10: uint16(7),
		11: uint16(7),
		12: uint16(7),
		13: uint16(7),
		14: uint16(7),
		15: uint16(7),
		16: uint16(7),
		17: uint16(7),
		18: uint16(7),
		19: uint16(7),
		20: uint16(7),
		21: uint16(7),
		22: uint16(7),
		23: uint16(7),
		24: uint16(7),
		25: uint16(13),
		28: uint16(3),
		30: uint16(3),
		31: uint16(3),
		32: uint16(4),
		34: uint16(3),
	},
	3: {
		0:  uint16(17),
		1:  uint16(19),
		2:  uint16(19),
		4:  uint16(19),
		5:  uint16(3),
		6:  uint16(22),
		7:  uint16(25),
		8:  uint16(19),
		9:  uint16(19),
		10: uint16(19),
		11: uint16(19),
		12: uint16(19),
		13: uint16(19),
		14: uint16(19),
		15: uint16(19),
		16: uint16(19),
		17: uint16(19),
		18: uint16(19),
		19: uint16(19),
		20: uint16(19),
		21: uint16(19),
		22: uint16(19),
		23: uint16(19),
		24: uint16(19),
		25: uint16(28),
		28: uint16(3),
		30: uint16(3),
		31: uint16(3),
		32: uint16(4),
		34: uint16(3),
	},
	4: {
		0:  uint16(31),
		1:  uint16(31),
		2:  uint16(31),
		4:  uint16(31),
		5:  uint16(3),
		6:  uint16(31),
		7:  uint16(33),
		8:  uint16(31),
		9:  uint16(31),
		10: uint16(31),
		11: uint16(31),
		12: uint16(31),
		13: uint16(31),
		14: uint16(31),
		15: uint16(31),
		16: uint16(31),
		17: uint16(31),
		18: uint16(31),
		19: uint16(31),
		20: uint16(31),
		21: uint16(31),
		22: uint16(31),
		23: uint16(31),
		24: uint16(31),
		25: uint16(31),
	},
	5: {
		0:  uint16(35),
		1:  uint16(35),
		2:  uint16(35),
		4:  uint16(35),
		5:  uint16(3),
		6:  uint16(35),
		7:  uint16(37),
		8:  uint16(35),
		9:  uint16(35),
		10: uint16(35),
		11: uint16(35),
		12: uint16(35),
		13: uint16(35),
		14: uint16(35),
		15: uint16(35),
		16: uint16(35),
		17: uint16(35),
		18: uint16(35),
		19: uint16(35),
		20: uint16(35),
		21: uint16(35),
		22: uint16(35),
		23: uint16(35),
		24: uint16(35),
		25: uint16(35),
	},
	6: {
		0:  uint16(39),
		1:  uint16(39),
		2:  uint16(39),
		4:  uint16(39),
		5:  uint16(3),
		6:  uint16(39),
		7:  uint16(41),
		8:  uint16(39),
		9:  uint16(39),
		10: uint16(39),
		11: uint16(39),
		12: uint16(39),
		13: uint16(39),
		14: uint16(39),
		15: uint16(39),
		16: uint16(39),
		17: uint16(39),
		18: uint16(39),
		19: uint16(39),
		20: uint16(39),
		21: uint16(39),
		22: uint16(39),
		23: uint16(39),
		24: uint16(39),
		25: uint16(39),
	},
	7: {
		0:  uint16(43),
		1:  uint16(43),
		2:  uint16(43),
		4:  uint16(43),
		5:  uint16(3),
		6:  uint16(43),
		7:  uint16(45),
		8:  uint16(43),
		9:  uint16(43),
		10: uint16(43),
		11: uint16(43),
		12: uint16(43),
		13: uint16(43),
		14: uint16(43),
		15: uint16(43),
		16: uint16(43),
		17: uint16(43),
		18: uint16(43),
		19: uint16(43),
		20: uint16(43),
		21: uint16(43),
		22: uint16(43),
		23: uint16(43),
		24: uint16(43),
		25: uint16(43),
	},
	8: {
		0:  uint16(47),
		1:  uint16(47),
		2:  uint16(47),
		4:  uint16(47),
		5:  uint16(3),
		6:  uint16(47),
		7:  uint16(49),
		8:  uint16(47),
		9:  uint16(47),
		10: uint16(47),
		11: uint16(47),
		12: uint16(47),
		13: uint16(47),
		14: uint16(47),
		15: uint16(47),
		16: uint16(47),
		17: uint16(47),
		18: uint16(47),
		19: uint16(47),
		20: uint16(47),
		21: uint16(47),
		22: uint16(47),
		23: uint16(47),
		24: uint16(47),
		25: uint16(47),
	},
}

var ts_small_parse_table = [71]uint16_t{
	0:  uint16(3),
	1:  uint16(53),
	2:  uint16(1),
	3:  uint16(aux_sym__full_uri_token1),
	4:  uint16(7),
	5:  uint16(1),
	6:  uint16(sym__end_char),
	7:  uint16(51),
	8:  uint16(14),
	9:  uint16(anon_sym_COLON),
	10: uint16(anon_sym_RPAREN),
	11: uint16(anon_sym_SQUOTE),
	12: uint16(anon_sym_DQUOTE),
	13: uint16(anon_sym_BQUOTE),
	14: uint16(anon_sym_DOT),
	15: uint16(anon_sym_COMMA),
	16: uint16(anon_sym_SEMI),
	17: uint16(anon_sym_BANG),
	18: uint16(anon_sym_QMARK),
	19: uint16(anon_sym_BSLASH),
	20: uint16(anon_sym_RBRACE),
	21: uint16(anon_sym_RBRACK),
	22: uint16(anon_sym_GT),
	23: uint16(4),
	24: uint16(3),
	25: uint16(1),
	26: uint16(aux_sym__full_uri_token1),
	27: uint16(55),
	28: uint16(1),
	29: uint16(anon_sym_COLON),
	30: uint16(57),
	31: uint16(1),
	32: uint16(anon_sym_LPAREN),
	33: uint16(13),
	34: uint16(1),
	35: uint16(sym__user),
	36: uint16(2),
	37: uint16(3),
	38: uint16(1),
	39: uint16(aux_sym__full_uri_token1),
	40: uint16(59),
	41: uint16(1),
	43: uint16(2),
	44: uint16(61),
	45: uint16(1),
	46: uint16(aux_sym__user_token1),
	47: uint16(63),
	48: uint16(1),
	49: uint16(aux_sym__full_uri_token1),
	50: uint16(2),
	51: uint16(3),
	52: uint16(1),
	53: uint16(aux_sym__full_uri_token1),
	54: uint16(65),
	55: uint16(1),
	56: uint16(anon_sym_COLON),
	57: uint16(2),
	58: uint16(3),
	59: uint16(1),
	60: uint16(aux_sym__full_uri_token1),
	61: uint16(67),
	62: uint16(1),
	63: uint16(anon_sym_RPAREN),
	64: uint16(2),
	65: uint16(3),
	66: uint16(1),
	67: uint16(aux_sym__full_uri_token1),
	68: uint16(69),
	69: uint16(1),
	70: uint16(anon_sym_COLON),
}

var ts_small_parse_table_map = [7]uint32_t{
	1: uint32(23),
	2: uint32(36),
	3: uint32(43),
	4: uint32(50),
	5: uint32(57),
	6: uint32(64),
}

var ts_parse_actions = [71]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_source),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_source),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	18: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	20: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
	21: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	22: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	23: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(9)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_source_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fsymbol:      uint16(aux_sym_source_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__text),
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
		Fcount: uint8(1),
	}})),
	34: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__text),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_tag),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_tag),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__end_char),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__end_char),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__full_uri),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__full_uri),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_tag),
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
		Fcount: uint8(1),
	}})),
	50: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_tag),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
	60: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__user),
	})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token_name = 0
const ts_external_token_invalid_token = 1

var ts_external_scanner_symbol_map = [2]TSSymbol{
	0: uint16(sym_name),
	1: uint16(sym_invalid_token),
}

var ts_external_scanner_states = [3][2]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
	},
	2: {
		0: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_comment(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(LANGUAGE_VERSION),
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
	Fname:              __ccgo_ts + 160,
	Fmetadata: TSLanguageMetadata{
		Fminor_version: uint8(2),
	},
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_comment_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_comment_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_comment_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_comment_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_comment_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00:\x00(\x00user\x00)\x00_full_uri_token1\x00uri\x00_text_token1\x00/\x00'\x00\"\x00`\x00<\x00[\x00{\x00.\x00,\x00;\x00!\x00?\x00\\\x00}\x00]\x00>\x00-\x00name\x00invalid_token\x00source\x00tag\x00_user\x00_full_uri\x00text\x00_stop_char\x00source_repeat1\x00comment\x00"
