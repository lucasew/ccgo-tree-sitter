// Code generated for linux/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-ron/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-ron -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && amd64

package grammar_ron

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
const EXTERNAL_TOKEN_COUNT = 4
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
const LARGE_STATE_COUNT = 15
const MAX_ALIAS_SEQUENCE_LENGTH = 5
const PRODUCTION_ID_COUNT = 4
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 81
const SYMBOL_COUNT = 51
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
const _LP64 = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BIGGEST_ALIGNMENT__ = 16
const __BIG_ENDIAN = 4321
const __BITINT_MAXWIDTH__ = 8388608
const __BOOL_WIDTH__ = 8
const __BYTE_ORDER = 1234
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CHAR_BIT__ = 8
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
const __FLOAT128__ = 1
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
const __FUNCTION__ = "__func__"
const __FXSR__ = 1
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
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 4.9406564584124654e-324
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 2.2204460492503131e-16
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 1.7976931348623157e+308
const __LDBL_MIN__ = 2.2250738585072014e-308
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
const __MMX__ = 1
const __NO_INLINE__ = 1
const __NO_MATH_INLINES = 1
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
const __SEG_FS = 1
const __SEG_GS = 1
const __SHRT_MAX__ = 32767
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 2147483647
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT128__ = 16
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
const __SSE2_MATH__ = 1
const __SSE2__ = 1
const __SSE_MATH__ = 1
const __SSE__ = 1
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
const __WCHAR_MAX__ = 2147483647
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 4294967295
const __WINT_UNSIGNED__ = 1
const __WINT_WIDTH__ = 32
const __amd64 = 1
const __amd64__ = 1
const __bool_true_false_are_defined = 1
const __clang__ = 1
const __clang_literal_encoding__ = "UTF-8"
const __clang_major__ = 18
const __clang_minor__ = 1
const __clang_patchlevel__ = 3
const __clang_version__ = "18.1.3 (1ubuntu1)"
const __clang_wide_literal_encoding__ = "UTF-32"
const __code_model_small__ = 1
const __gnu_linux__ = 1
const __inline = "inline"
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __llvm__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict = "restrict"
const __restrict_arr = "restrict"
const __tune_k8__ = 1
const __unix = 1
const __unix__ = 1
const __x86_64 = 1
const __x86_64__ = 1
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

type __predefined_wchar_t = int32

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

type wchar_t = int32

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

type TokenType = int32

const STRING_CONTENT = 0
const RAW_STRING = 1
const FLOAT = 2
const BLOCK_COMMENT = 3

func tree_sitter_ron_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_ron_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
}

func tree_sitter_ron_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_ron_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
}

func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
}

func is_num_char(tls *libc.TLS, c int32_t) (r uint8) {
	return libc.BoolUint8(c == int32('_') || libc.BoolInt32(libc.Uint32FromInt32(c)-uint32('0') < uint32(10)) != 0)
}

func tree_sitter_ron_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var after_star, has_content, has_exponent, has_fraction uint8
	var hash_count, nesting_depth, opening_hash_count uint32
	_, _, _, _, _, _, _ = after_star, has_content, has_exponent, has_fraction, hash_count, nesting_depth, opening_hash_count
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(STRING_CONTENT))) != 0 && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(FLOAT))) != 0) {
		has_content = libc.BoolUint8(false1 != 0)
		for {
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('"') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\\') {
				break
			}
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == 0 {
				return libc.BoolUint8(false1 != 0)
			}
			has_content = libc.BoolUint8(true1 != 0)
			advance(tls, lexer)
			goto _1
		_1:
		}
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(STRING_CONTENT)
		return has_content
	}
	for libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(RAW_STRING))) != 0 && ((*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('r') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('b')) {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(RAW_STRING)
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('b') {
			advance(tls, lexer)
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('r') {
			return libc.BoolUint8(false1 != 0)
		}
		advance(tls, lexer)
		opening_hash_count = uint32(0)
		for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('#') {
			advance(tls, lexer)
			opening_hash_count = opening_hash_count + 1
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('"') {
			return libc.BoolUint8(false1 != 0)
		}
		advance(tls, lexer)
		for {
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == 0 {
				return libc.BoolUint8(false1 != 0)
			}
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('"') {
				advance(tls, lexer)
				hash_count = uint32(0)
				for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('#') && hash_count < opening_hash_count {
					advance(tls, lexer)
					hash_count = hash_count + 1
				}
				if hash_count == opening_hash_count {
					return libc.BoolUint8(true1 != 0)
				}
			} else {
				advance(tls, lexer)
			}
			goto _2
		_2:
		}
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(FLOAT))) != 0 && libc.BoolInt32(libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)-uint32('0') < uint32(10)) != 0 {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(FLOAT)
		advance(tls, lexer)
		for is_num_char(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0 {
			advance(tls, lexer)
		}
		has_fraction = libc.BoolUint8(false1 != 0)
		has_exponent = libc.BoolUint8(false1 != 0)
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('.') {
			has_fraction = libc.BoolUint8(true1 != 0)
			advance(tls, lexer)
			if libc.Xiswalpha(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
				// The dot is followed by a letter: 1.max(2) => not a float but an
				// integer
				return libc.BoolUint8(false1 != 0)
			}
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('.') {
				return libc.BoolUint8(false1 != 0)
			}
			for is_num_char(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0 {
				advance(tls, lexer)
			}
		}
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('e') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('E') {
			has_exponent = libc.BoolUint8(true1 != 0)
			advance(tls, lexer)
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('+') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('-') {
				advance(tls, lexer)
			}
			if !(is_num_char(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0) {
				return libc.BoolUint8(true1 != 0)
			}
			advance(tls, lexer)
			for is_num_char(tls, (*TSLexer)(unsafe.Pointer(lexer)).Flookahead) != 0 {
				advance(tls, lexer)
			}
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		}
		if !(has_exponent != 0) && !(has_fraction != 0) {
			return libc.BoolUint8(false1 != 0)
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('u') && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('i') && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('f') {
			return libc.BoolUint8(true1 != 0)
		}
		advance(tls, lexer)
		if !(libc.BoolInt32(libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)-libc.Uint32FromUint8('0') < libc.Uint32FromInt32(10)) != 0) {
			return libc.BoolUint8(true1 != 0)
		}
		for libc.BoolInt32(libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)-uint32('0') < uint32(10)) != 0 {
			advance(tls, lexer)
		}
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return libc.BoolUint8(true1 != 0)
	}
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
		advance(tls, lexer)
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('*') {
			return libc.BoolUint8(false1 != 0)
		}
		advance(tls, lexer)
		after_star = libc.BoolUint8(false1 != 0)
		nesting_depth = uint32(1)
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
						(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BLOCK_COMMENT)
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

const anon_sym_LBRACK = 1
const anon_sym_COMMA = 2
const anon_sym_RBRACK = 3
const anon_sym_LBRACE = 4
const anon_sym_RBRACE = 5
const anon_sym_LPAREN_RPAREN = 6
const anon_sym_LPAREN = 7
const anon_sym_RPAREN = 8
const anon_sym_COLON = 9
const sym_integer = 10
const anon_sym_DASH = 11
const aux_sym_string_token1 = 12
const anon_sym_DQUOTE = 13
const anon_sym_b = 14
const anon_sym_SQUOTE = 15
const aux_sym_char_token1 = 16
const aux_sym__escape_sequence_token1 = 17
const sym_escape_sequence = 18
const anon_sym_true = 19
const anon_sym_false = 20
const sym_identifier = 21
const sym_line_comment = 22
const sym__string_content = 23
const sym_raw_string = 24
const sym_float = 25
const sym_block_comment = 26
const sym_source_file = 27
const sym__value = 28
const sym_enum_variant = 29
const sym_array = 30
const sym_map = 31
const sym_struct = 32
const sym_unit_struct = 33
const sym_struct_name = 34
const sym__tuple_struct = 35
const sym__named_struct = 36
const sym__struct_body = 37
const sym_tuple = 38
const sym_map_entry = 39
const sym_struct_entry = 40
const sym__literal = 41
const sym_negative = 42
const sym_string = 43
const sym_char = 44
const sym__escape_sequence = 45
const sym_boolean = 46
const aux_sym_array_repeat1 = 47
const aux_sym_map_repeat1 = 48
const aux_sym__struct_body_repeat1 = 49
const aux_sym_string_repeat1 = 50

var ts_symbol_names = [51]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 6,
	3:  __ccgo_ts + 8,
	4:  __ccgo_ts + 10,
	5:  __ccgo_ts + 12,
	6:  __ccgo_ts + 14,
	7:  __ccgo_ts + 17,
	8:  __ccgo_ts + 19,
	9:  __ccgo_ts + 21,
	10: __ccgo_ts + 23,
	11: __ccgo_ts + 31,
	12: __ccgo_ts + 33,
	13: __ccgo_ts + 33,
	14: __ccgo_ts + 35,
	15: __ccgo_ts + 37,
	16: __ccgo_ts + 39,
	17: __ccgo_ts + 51,
	18: __ccgo_ts + 75,
	19: __ccgo_ts + 91,
	20: __ccgo_ts + 96,
	21: __ccgo_ts + 102,
	22: __ccgo_ts + 113,
	23: __ccgo_ts + 126,
	24: __ccgo_ts + 142,
	25: __ccgo_ts + 153,
	26: __ccgo_ts + 159,
	27: __ccgo_ts + 173,
	28: __ccgo_ts + 185,
	29: __ccgo_ts + 192,
	30: __ccgo_ts + 205,
	31: __ccgo_ts + 211,
	32: __ccgo_ts + 215,
	33: __ccgo_ts + 222,
	34: __ccgo_ts + 234,
	35: __ccgo_ts + 246,
	36: __ccgo_ts + 260,
	37: __ccgo_ts + 274,
	38: __ccgo_ts + 287,
	39: __ccgo_ts + 293,
	40: __ccgo_ts + 303,
	41: __ccgo_ts + 316,
	42: __ccgo_ts + 325,
	43: __ccgo_ts + 334,
	44: __ccgo_ts + 341,
	45: __ccgo_ts + 346,
	46: __ccgo_ts + 363,
	47: __ccgo_ts + 371,
	48: __ccgo_ts + 385,
	49: __ccgo_ts + 397,
	50: __ccgo_ts + 418,
}

var ts_symbol_map = [51]TSSymbol{
	1:  uint16(anon_sym_LBRACK),
	2:  uint16(anon_sym_COMMA),
	3:  uint16(anon_sym_RBRACK),
	4:  uint16(anon_sym_LBRACE),
	5:  uint16(anon_sym_RBRACE),
	6:  uint16(anon_sym_LPAREN_RPAREN),
	7:  uint16(anon_sym_LPAREN),
	8:  uint16(anon_sym_RPAREN),
	9:  uint16(anon_sym_COLON),
	10: uint16(sym_integer),
	11: uint16(anon_sym_DASH),
	12: uint16(anon_sym_DQUOTE),
	13: uint16(anon_sym_DQUOTE),
	14: uint16(anon_sym_b),
	15: uint16(anon_sym_SQUOTE),
	16: uint16(aux_sym_char_token1),
	17: uint16(aux_sym__escape_sequence_token1),
	18: uint16(sym_escape_sequence),
	19: uint16(anon_sym_true),
	20: uint16(anon_sym_false),
	21: uint16(sym_identifier),
	22: uint16(sym_line_comment),
	23: uint16(sym__string_content),
	24: uint16(sym_raw_string),
	25: uint16(sym_float),
	26: uint16(sym_block_comment),
	27: uint16(sym_source_file),
	28: uint16(sym__value),
	29: uint16(sym_enum_variant),
	30: uint16(sym_array),
	31: uint16(sym_map),
	32: uint16(sym_struct),
	33: uint16(sym_unit_struct),
	34: uint16(sym_struct_name),
	35: uint16(sym__tuple_struct),
	36: uint16(sym__named_struct),
	37: uint16(sym__struct_body),
	38: uint16(sym_tuple),
	39: uint16(sym_map_entry),
	40: uint16(sym_struct_entry),
	41: uint16(sym__literal),
	42: uint16(sym_negative),
	43: uint16(sym_string),
	44: uint16(sym_char),
	45: uint16(sym__escape_sequence),
	46: uint16(sym_boolean),
	47: uint16(aux_sym_array_repeat1),
	48: uint16(aux_sym_map_repeat1),
	49: uint16(aux_sym__struct_body_repeat1),
	50: uint16(aux_sym_string_repeat1),
}

var ts_symbol_metadata = [51]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	16: {},
	17: {},
	18: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	27: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	28: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	35: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	36: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	37: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	41: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	46: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	47: {},
	48: {},
	49: {},
	50: {},
}

type ts_field_identifiers = int32

const field_body = 1

var ts_field_names = [2]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 433,
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
		Flength: uint16(1),
	},
}

var ts_field_map_entries = [3]TSFieldMapEntry{
	0: {
		Ffield_id:  uint16(field_body),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	1: {
		Ffield_id: uint16(field_body),
	},
	2: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(1),
	},
}

var ts_alias_sequences = [4][5]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [81]TSStateId{
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
	30: uint16(30),
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
	42: uint16(42),
	43: uint16(43),
	44: uint16(44),
	45: uint16(45),
	46: uint16(46),
	47: uint16(47),
	48: uint16(48),
	49: uint16(49),
	50: uint16(50),
	51: uint16(51),
	52: uint16(52),
	53: uint16(53),
	54: uint16(54),
	55: uint16(55),
	56: uint16(56),
	57: uint16(57),
	58: uint16(58),
	59: uint16(59),
	60: uint16(60),
	61: uint16(61),
	62: uint16(62),
	63: uint16(63),
	64: uint16(64),
	65: uint16(65),
	66: uint16(66),
	67: uint16(67),
	68: uint16(68),
	69: uint16(69),
	70: uint16(70),
	71: uint16(71),
	72: uint16(72),
	73: uint16(73),
	74: uint16(74),
	75: uint16(75),
	76: uint16(76),
	77: uint16(77),
	78: uint16(78),
	79: uint16(79),
	80: uint16(80),
}

var sym_identifier_character_set_1 = [656]TSCharacterRange{
	0: {
		Fstart: int32('A'),
		Fend:   int32('Z'),
	},
	1: {
		Fstart: int32('_'),
		Fend:   int32('_'),
	},
	2: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
	3: {
		Fstart: int32(0xaa),
		Fend:   int32(0xaa),
	},
	4: {
		Fstart: int32(0xb5),
		Fend:   int32(0xb5),
	},
	5: {
		Fstart: int32(0xba),
		Fend:   int32(0xba),
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
		Fend:   int32(0x2c1),
	},
	9: {
		Fstart: int32(0x2c6),
		Fend:   int32(0x2d1),
	},
	10: {
		Fstart: int32(0x2e0),
		Fend:   int32(0x2e4),
	},
	11: {
		Fstart: int32(0x2ec),
		Fend:   int32(0x2ec),
	},
	12: {
		Fstart: int32(0x2ee),
		Fend:   int32(0x2ee),
	},
	13: {
		Fstart: int32(0x370),
		Fend:   int32(0x374),
	},
	14: {
		Fstart: int32(0x376),
		Fend:   int32(0x377),
	},
	15: {
		Fstart: int32(0x37b),
		Fend:   int32(0x37d),
	},
	16: {
		Fstart: int32(0x37f),
		Fend:   int32(0x37f),
	},
	17: {
		Fstart: int32(0x386),
		Fend:   int32(0x386),
	},
	18: {
		Fstart: int32(0x388),
		Fend:   int32(0x38a),
	},
	19: {
		Fstart: int32(0x38c),
		Fend:   int32(0x38c),
	},
	20: {
		Fstart: int32(0x38e),
		Fend:   int32(0x3a1),
	},
	21: {
		Fstart: int32(0x3a3),
		Fend:   int32(0x3f5),
	},
	22: {
		Fstart: int32(0x3f7),
		Fend:   int32(0x481),
	},
	23: {
		Fstart: int32(0x48a),
		Fend:   int32(0x52f),
	},
	24: {
		Fstart: int32(0x531),
		Fend:   int32(0x556),
	},
	25: {
		Fstart: int32(0x559),
		Fend:   int32(0x559),
	},
	26: {
		Fstart: int32(0x560),
		Fend:   int32(0x588),
	},
	27: {
		Fstart: int32(0x5d0),
		Fend:   int32(0x5ea),
	},
	28: {
		Fstart: int32(0x5ef),
		Fend:   int32(0x5f2),
	},
	29: {
		Fstart: int32(0x620),
		Fend:   int32(0x64a),
	},
	30: {
		Fstart: int32(0x66e),
		Fend:   int32(0x66f),
	},
	31: {
		Fstart: int32(0x671),
		Fend:   int32(0x6d3),
	},
	32: {
		Fstart: int32(0x6d5),
		Fend:   int32(0x6d5),
	},
	33: {
		Fstart: int32(0x6e5),
		Fend:   int32(0x6e6),
	},
	34: {
		Fstart: int32(0x6ee),
		Fend:   int32(0x6ef),
	},
	35: {
		Fstart: int32(0x6fa),
		Fend:   int32(0x6fc),
	},
	36: {
		Fstart: int32(0x6ff),
		Fend:   int32(0x6ff),
	},
	37: {
		Fstart: int32(0x710),
		Fend:   int32(0x710),
	},
	38: {
		Fstart: int32(0x712),
		Fend:   int32(0x72f),
	},
	39: {
		Fstart: int32(0x74d),
		Fend:   int32(0x7a5),
	},
	40: {
		Fstart: int32(0x7b1),
		Fend:   int32(0x7b1),
	},
	41: {
		Fstart: int32(0x7ca),
		Fend:   int32(0x7ea),
	},
	42: {
		Fstart: int32(0x7f4),
		Fend:   int32(0x7f5),
	},
	43: {
		Fstart: int32(0x7fa),
		Fend:   int32(0x7fa),
	},
	44: {
		Fstart: int32(0x800),
		Fend:   int32(0x815),
	},
	45: {
		Fstart: int32(0x81a),
		Fend:   int32(0x81a),
	},
	46: {
		Fstart: int32(0x824),
		Fend:   int32(0x824),
	},
	47: {
		Fstart: int32(0x828),
		Fend:   int32(0x828),
	},
	48: {
		Fstart: int32(0x840),
		Fend:   int32(0x858),
	},
	49: {
		Fstart: int32(0x860),
		Fend:   int32(0x86a),
	},
	50: {
		Fstart: int32(0x870),
		Fend:   int32(0x887),
	},
	51: {
		Fstart: int32(0x889),
		Fend:   int32(0x88e),
	},
	52: {
		Fstart: int32(0x8a0),
		Fend:   int32(0x8c9),
	},
	53: {
		Fstart: int32(0x904),
		Fend:   int32(0x939),
	},
	54: {
		Fstart: int32(0x93d),
		Fend:   int32(0x93d),
	},
	55: {
		Fstart: int32(0x950),
		Fend:   int32(0x950),
	},
	56: {
		Fstart: int32(0x958),
		Fend:   int32(0x961),
	},
	57: {
		Fstart: int32(0x971),
		Fend:   int32(0x980),
	},
	58: {
		Fstart: int32(0x985),
		Fend:   int32(0x98c),
	},
	59: {
		Fstart: int32(0x98f),
		Fend:   int32(0x990),
	},
	60: {
		Fstart: int32(0x993),
		Fend:   int32(0x9a8),
	},
	61: {
		Fstart: int32(0x9aa),
		Fend:   int32(0x9b0),
	},
	62: {
		Fstart: int32(0x9b2),
		Fend:   int32(0x9b2),
	},
	63: {
		Fstart: int32(0x9b6),
		Fend:   int32(0x9b9),
	},
	64: {
		Fstart: int32(0x9bd),
		Fend:   int32(0x9bd),
	},
	65: {
		Fstart: int32(0x9ce),
		Fend:   int32(0x9ce),
	},
	66: {
		Fstart: int32(0x9dc),
		Fend:   int32(0x9dd),
	},
	67: {
		Fstart: int32(0x9df),
		Fend:   int32(0x9e1),
	},
	68: {
		Fstart: int32(0x9f0),
		Fend:   int32(0x9f1),
	},
	69: {
		Fstart: int32(0x9fc),
		Fend:   int32(0x9fc),
	},
	70: {
		Fstart: int32(0xa05),
		Fend:   int32(0xa0a),
	},
	71: {
		Fstart: int32(0xa0f),
		Fend:   int32(0xa10),
	},
	72: {
		Fstart: int32(0xa13),
		Fend:   int32(0xa28),
	},
	73: {
		Fstart: int32(0xa2a),
		Fend:   int32(0xa30),
	},
	74: {
		Fstart: int32(0xa32),
		Fend:   int32(0xa33),
	},
	75: {
		Fstart: int32(0xa35),
		Fend:   int32(0xa36),
	},
	76: {
		Fstart: int32(0xa38),
		Fend:   int32(0xa39),
	},
	77: {
		Fstart: int32(0xa59),
		Fend:   int32(0xa5c),
	},
	78: {
		Fstart: int32(0xa5e),
		Fend:   int32(0xa5e),
	},
	79: {
		Fstart: int32(0xa72),
		Fend:   int32(0xa74),
	},
	80: {
		Fstart: int32(0xa85),
		Fend:   int32(0xa8d),
	},
	81: {
		Fstart: int32(0xa8f),
		Fend:   int32(0xa91),
	},
	82: {
		Fstart: int32(0xa93),
		Fend:   int32(0xaa8),
	},
	83: {
		Fstart: int32(0xaaa),
		Fend:   int32(0xab0),
	},
	84: {
		Fstart: int32(0xab2),
		Fend:   int32(0xab3),
	},
	85: {
		Fstart: int32(0xab5),
		Fend:   int32(0xab9),
	},
	86: {
		Fstart: int32(0xabd),
		Fend:   int32(0xabd),
	},
	87: {
		Fstart: int32(0xad0),
		Fend:   int32(0xad0),
	},
	88: {
		Fstart: int32(0xae0),
		Fend:   int32(0xae1),
	},
	89: {
		Fstart: int32(0xaf9),
		Fend:   int32(0xaf9),
	},
	90: {
		Fstart: int32(0xb05),
		Fend:   int32(0xb0c),
	},
	91: {
		Fstart: int32(0xb0f),
		Fend:   int32(0xb10),
	},
	92: {
		Fstart: int32(0xb13),
		Fend:   int32(0xb28),
	},
	93: {
		Fstart: int32(0xb2a),
		Fend:   int32(0xb30),
	},
	94: {
		Fstart: int32(0xb32),
		Fend:   int32(0xb33),
	},
	95: {
		Fstart: int32(0xb35),
		Fend:   int32(0xb39),
	},
	96: {
		Fstart: int32(0xb3d),
		Fend:   int32(0xb3d),
	},
	97: {
		Fstart: int32(0xb5c),
		Fend:   int32(0xb5d),
	},
	98: {
		Fstart: int32(0xb5f),
		Fend:   int32(0xb61),
	},
	99: {
		Fstart: int32(0xb71),
		Fend:   int32(0xb71),
	},
	100: {
		Fstart: int32(0xb83),
		Fend:   int32(0xb83),
	},
	101: {
		Fstart: int32(0xb85),
		Fend:   int32(0xb8a),
	},
	102: {
		Fstart: int32(0xb8e),
		Fend:   int32(0xb90),
	},
	103: {
		Fstart: int32(0xb92),
		Fend:   int32(0xb95),
	},
	104: {
		Fstart: int32(0xb99),
		Fend:   int32(0xb9a),
	},
	105: {
		Fstart: int32(0xb9c),
		Fend:   int32(0xb9c),
	},
	106: {
		Fstart: int32(0xb9e),
		Fend:   int32(0xb9f),
	},
	107: {
		Fstart: int32(0xba3),
		Fend:   int32(0xba4),
	},
	108: {
		Fstart: int32(0xba8),
		Fend:   int32(0xbaa),
	},
	109: {
		Fstart: int32(0xbae),
		Fend:   int32(0xbb9),
	},
	110: {
		Fstart: int32(0xbd0),
		Fend:   int32(0xbd0),
	},
	111: {
		Fstart: int32(0xc05),
		Fend:   int32(0xc0c),
	},
	112: {
		Fstart: int32(0xc0e),
		Fend:   int32(0xc10),
	},
	113: {
		Fstart: int32(0xc12),
		Fend:   int32(0xc28),
	},
	114: {
		Fstart: int32(0xc2a),
		Fend:   int32(0xc39),
	},
	115: {
		Fstart: int32(0xc3d),
		Fend:   int32(0xc3d),
	},
	116: {
		Fstart: int32(0xc58),
		Fend:   int32(0xc5a),
	},
	117: {
		Fstart: int32(0xc5d),
		Fend:   int32(0xc5d),
	},
	118: {
		Fstart: int32(0xc60),
		Fend:   int32(0xc61),
	},
	119: {
		Fstart: int32(0xc80),
		Fend:   int32(0xc80),
	},
	120: {
		Fstart: int32(0xc85),
		Fend:   int32(0xc8c),
	},
	121: {
		Fstart: int32(0xc8e),
		Fend:   int32(0xc90),
	},
	122: {
		Fstart: int32(0xc92),
		Fend:   int32(0xca8),
	},
	123: {
		Fstart: int32(0xcaa),
		Fend:   int32(0xcb3),
	},
	124: {
		Fstart: int32(0xcb5),
		Fend:   int32(0xcb9),
	},
	125: {
		Fstart: int32(0xcbd),
		Fend:   int32(0xcbd),
	},
	126: {
		Fstart: int32(0xcdd),
		Fend:   int32(0xcde),
	},
	127: {
		Fstart: int32(0xce0),
		Fend:   int32(0xce1),
	},
	128: {
		Fstart: int32(0xcf1),
		Fend:   int32(0xcf2),
	},
	129: {
		Fstart: int32(0xd04),
		Fend:   int32(0xd0c),
	},
	130: {
		Fstart: int32(0xd0e),
		Fend:   int32(0xd10),
	},
	131: {
		Fstart: int32(0xd12),
		Fend:   int32(0xd3a),
	},
	132: {
		Fstart: int32(0xd3d),
		Fend:   int32(0xd3d),
	},
	133: {
		Fstart: int32(0xd4e),
		Fend:   int32(0xd4e),
	},
	134: {
		Fstart: int32(0xd54),
		Fend:   int32(0xd56),
	},
	135: {
		Fstart: int32(0xd5f),
		Fend:   int32(0xd61),
	},
	136: {
		Fstart: int32(0xd7a),
		Fend:   int32(0xd7f),
	},
	137: {
		Fstart: int32(0xd85),
		Fend:   int32(0xd96),
	},
	138: {
		Fstart: int32(0xd9a),
		Fend:   int32(0xdb1),
	},
	139: {
		Fstart: int32(0xdb3),
		Fend:   int32(0xdbb),
	},
	140: {
		Fstart: int32(0xdbd),
		Fend:   int32(0xdbd),
	},
	141: {
		Fstart: int32(0xdc0),
		Fend:   int32(0xdc6),
	},
	142: {
		Fstart: int32(0xe01),
		Fend:   int32(0xe30),
	},
	143: {
		Fstart: int32(0xe32),
		Fend:   int32(0xe32),
	},
	144: {
		Fstart: int32(0xe40),
		Fend:   int32(0xe46),
	},
	145: {
		Fstart: int32(0xe81),
		Fend:   int32(0xe82),
	},
	146: {
		Fstart: int32(0xe84),
		Fend:   int32(0xe84),
	},
	147: {
		Fstart: int32(0xe86),
		Fend:   int32(0xe8a),
	},
	148: {
		Fstart: int32(0xe8c),
		Fend:   int32(0xea3),
	},
	149: {
		Fstart: int32(0xea5),
		Fend:   int32(0xea5),
	},
	150: {
		Fstart: int32(0xea7),
		Fend:   int32(0xeb0),
	},
	151: {
		Fstart: int32(0xeb2),
		Fend:   int32(0xeb2),
	},
	152: {
		Fstart: int32(0xebd),
		Fend:   int32(0xebd),
	},
	153: {
		Fstart: int32(0xec0),
		Fend:   int32(0xec4),
	},
	154: {
		Fstart: int32(0xec6),
		Fend:   int32(0xec6),
	},
	155: {
		Fstart: int32(0xedc),
		Fend:   int32(0xedf),
	},
	156: {
		Fstart: int32(0xf00),
		Fend:   int32(0xf00),
	},
	157: {
		Fstart: int32(0xf40),
		Fend:   int32(0xf47),
	},
	158: {
		Fstart: int32(0xf49),
		Fend:   int32(0xf6c),
	},
	159: {
		Fstart: int32(0xf88),
		Fend:   int32(0xf8c),
	},
	160: {
		Fstart: int32(0x1000),
		Fend:   int32(0x102a),
	},
	161: {
		Fstart: int32(0x103f),
		Fend:   int32(0x103f),
	},
	162: {
		Fstart: int32(0x1050),
		Fend:   int32(0x1055),
	},
	163: {
		Fstart: int32(0x105a),
		Fend:   int32(0x105d),
	},
	164: {
		Fstart: int32(0x1061),
		Fend:   int32(0x1061),
	},
	165: {
		Fstart: int32(0x1065),
		Fend:   int32(0x1066),
	},
	166: {
		Fstart: int32(0x106e),
		Fend:   int32(0x1070),
	},
	167: {
		Fstart: int32(0x1075),
		Fend:   int32(0x1081),
	},
	168: {
		Fstart: int32(0x108e),
		Fend:   int32(0x108e),
	},
	169: {
		Fstart: int32(0x10a0),
		Fend:   int32(0x10c5),
	},
	170: {
		Fstart: int32(0x10c7),
		Fend:   int32(0x10c7),
	},
	171: {
		Fstart: int32(0x10cd),
		Fend:   int32(0x10cd),
	},
	172: {
		Fstart: int32(0x10d0),
		Fend:   int32(0x10fa),
	},
	173: {
		Fstart: int32(0x10fc),
		Fend:   int32(0x1248),
	},
	174: {
		Fstart: int32(0x124a),
		Fend:   int32(0x124d),
	},
	175: {
		Fstart: int32(0x1250),
		Fend:   int32(0x1256),
	},
	176: {
		Fstart: int32(0x1258),
		Fend:   int32(0x1258),
	},
	177: {
		Fstart: int32(0x125a),
		Fend:   int32(0x125d),
	},
	178: {
		Fstart: int32(0x1260),
		Fend:   int32(0x1288),
	},
	179: {
		Fstart: int32(0x128a),
		Fend:   int32(0x128d),
	},
	180: {
		Fstart: int32(0x1290),
		Fend:   int32(0x12b0),
	},
	181: {
		Fstart: int32(0x12b2),
		Fend:   int32(0x12b5),
	},
	182: {
		Fstart: int32(0x12b8),
		Fend:   int32(0x12be),
	},
	183: {
		Fstart: int32(0x12c0),
		Fend:   int32(0x12c0),
	},
	184: {
		Fstart: int32(0x12c2),
		Fend:   int32(0x12c5),
	},
	185: {
		Fstart: int32(0x12c8),
		Fend:   int32(0x12d6),
	},
	186: {
		Fstart: int32(0x12d8),
		Fend:   int32(0x1310),
	},
	187: {
		Fstart: int32(0x1312),
		Fend:   int32(0x1315),
	},
	188: {
		Fstart: int32(0x1318),
		Fend:   int32(0x135a),
	},
	189: {
		Fstart: int32(0x1380),
		Fend:   int32(0x138f),
	},
	190: {
		Fstart: int32(0x13a0),
		Fend:   int32(0x13f5),
	},
	191: {
		Fstart: int32(0x13f8),
		Fend:   int32(0x13fd),
	},
	192: {
		Fstart: int32(0x1401),
		Fend:   int32(0x166c),
	},
	193: {
		Fstart: int32(0x166f),
		Fend:   int32(0x167f),
	},
	194: {
		Fstart: int32(0x1681),
		Fend:   int32(0x169a),
	},
	195: {
		Fstart: int32(0x16a0),
		Fend:   int32(0x16ea),
	},
	196: {
		Fstart: int32(0x16ee),
		Fend:   int32(0x16f8),
	},
	197: {
		Fstart: int32(0x1700),
		Fend:   int32(0x1711),
	},
	198: {
		Fstart: int32(0x171f),
		Fend:   int32(0x1731),
	},
	199: {
		Fstart: int32(0x1740),
		Fend:   int32(0x1751),
	},
	200: {
		Fstart: int32(0x1760),
		Fend:   int32(0x176c),
	},
	201: {
		Fstart: int32(0x176e),
		Fend:   int32(0x1770),
	},
	202: {
		Fstart: int32(0x1780),
		Fend:   int32(0x17b3),
	},
	203: {
		Fstart: int32(0x17d7),
		Fend:   int32(0x17d7),
	},
	204: {
		Fstart: int32(0x17dc),
		Fend:   int32(0x17dc),
	},
	205: {
		Fstart: int32(0x1820),
		Fend:   int32(0x1878),
	},
	206: {
		Fstart: int32(0x1880),
		Fend:   int32(0x18a8),
	},
	207: {
		Fstart: int32(0x18aa),
		Fend:   int32(0x18aa),
	},
	208: {
		Fstart: int32(0x18b0),
		Fend:   int32(0x18f5),
	},
	209: {
		Fstart: int32(0x1900),
		Fend:   int32(0x191e),
	},
	210: {
		Fstart: int32(0x1950),
		Fend:   int32(0x196d),
	},
	211: {
		Fstart: int32(0x1970),
		Fend:   int32(0x1974),
	},
	212: {
		Fstart: int32(0x1980),
		Fend:   int32(0x19ab),
	},
	213: {
		Fstart: int32(0x19b0),
		Fend:   int32(0x19c9),
	},
	214: {
		Fstart: int32(0x1a00),
		Fend:   int32(0x1a16),
	},
	215: {
		Fstart: int32(0x1a20),
		Fend:   int32(0x1a54),
	},
	216: {
		Fstart: int32(0x1aa7),
		Fend:   int32(0x1aa7),
	},
	217: {
		Fstart: int32(0x1b05),
		Fend:   int32(0x1b33),
	},
	218: {
		Fstart: int32(0x1b45),
		Fend:   int32(0x1b4c),
	},
	219: {
		Fstart: int32(0x1b83),
		Fend:   int32(0x1ba0),
	},
	220: {
		Fstart: int32(0x1bae),
		Fend:   int32(0x1baf),
	},
	221: {
		Fstart: int32(0x1bba),
		Fend:   int32(0x1be5),
	},
	222: {
		Fstart: int32(0x1c00),
		Fend:   int32(0x1c23),
	},
	223: {
		Fstart: int32(0x1c4d),
		Fend:   int32(0x1c4f),
	},
	224: {
		Fstart: int32(0x1c5a),
		Fend:   int32(0x1c7d),
	},
	225: {
		Fstart: int32(0x1c80),
		Fend:   int32(0x1c88),
	},
	226: {
		Fstart: int32(0x1c90),
		Fend:   int32(0x1cba),
	},
	227: {
		Fstart: int32(0x1cbd),
		Fend:   int32(0x1cbf),
	},
	228: {
		Fstart: int32(0x1ce9),
		Fend:   int32(0x1cec),
	},
	229: {
		Fstart: int32(0x1cee),
		Fend:   int32(0x1cf3),
	},
	230: {
		Fstart: int32(0x1cf5),
		Fend:   int32(0x1cf6),
	},
	231: {
		Fstart: int32(0x1cfa),
		Fend:   int32(0x1cfa),
	},
	232: {
		Fstart: int32(0x1d00),
		Fend:   int32(0x1dbf),
	},
	233: {
		Fstart: int32(0x1e00),
		Fend:   int32(0x1f15),
	},
	234: {
		Fstart: int32(0x1f18),
		Fend:   int32(0x1f1d),
	},
	235: {
		Fstart: int32(0x1f20),
		Fend:   int32(0x1f45),
	},
	236: {
		Fstart: int32(0x1f48),
		Fend:   int32(0x1f4d),
	},
	237: {
		Fstart: int32(0x1f50),
		Fend:   int32(0x1f57),
	},
	238: {
		Fstart: int32(0x1f59),
		Fend:   int32(0x1f59),
	},
	239: {
		Fstart: int32(0x1f5b),
		Fend:   int32(0x1f5b),
	},
	240: {
		Fstart: int32(0x1f5d),
		Fend:   int32(0x1f5d),
	},
	241: {
		Fstart: int32(0x1f5f),
		Fend:   int32(0x1f7d),
	},
	242: {
		Fstart: int32(0x1f80),
		Fend:   int32(0x1fb4),
	},
	243: {
		Fstart: int32(0x1fb6),
		Fend:   int32(0x1fbc),
	},
	244: {
		Fstart: int32(0x1fbe),
		Fend:   int32(0x1fbe),
	},
	245: {
		Fstart: int32(0x1fc2),
		Fend:   int32(0x1fc4),
	},
	246: {
		Fstart: int32(0x1fc6),
		Fend:   int32(0x1fcc),
	},
	247: {
		Fstart: int32(0x1fd0),
		Fend:   int32(0x1fd3),
	},
	248: {
		Fstart: int32(0x1fd6),
		Fend:   int32(0x1fdb),
	},
	249: {
		Fstart: int32(0x1fe0),
		Fend:   int32(0x1fec),
	},
	250: {
		Fstart: int32(0x1ff2),
		Fend:   int32(0x1ff4),
	},
	251: {
		Fstart: int32(0x1ff6),
		Fend:   int32(0x1ffc),
	},
	252: {
		Fstart: int32(0x2071),
		Fend:   int32(0x2071),
	},
	253: {
		Fstart: int32(0x207f),
		Fend:   int32(0x207f),
	},
	254: {
		Fstart: int32(0x2090),
		Fend:   int32(0x209c),
	},
	255: {
		Fstart: int32(0x2102),
		Fend:   int32(0x2102),
	},
	256: {
		Fstart: int32(0x2107),
		Fend:   int32(0x2107),
	},
	257: {
		Fstart: int32(0x210a),
		Fend:   int32(0x2113),
	},
	258: {
		Fstart: int32(0x2115),
		Fend:   int32(0x2115),
	},
	259: {
		Fstart: int32(0x2118),
		Fend:   int32(0x211d),
	},
	260: {
		Fstart: int32(0x2124),
		Fend:   int32(0x2124),
	},
	261: {
		Fstart: int32(0x2126),
		Fend:   int32(0x2126),
	},
	262: {
		Fstart: int32(0x2128),
		Fend:   int32(0x2128),
	},
	263: {
		Fstart: int32(0x212a),
		Fend:   int32(0x2139),
	},
	264: {
		Fstart: int32(0x213c),
		Fend:   int32(0x213f),
	},
	265: {
		Fstart: int32(0x2145),
		Fend:   int32(0x2149),
	},
	266: {
		Fstart: int32(0x214e),
		Fend:   int32(0x214e),
	},
	267: {
		Fstart: int32(0x2160),
		Fend:   int32(0x2188),
	},
	268: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2ce4),
	},
	269: {
		Fstart: int32(0x2ceb),
		Fend:   int32(0x2cee),
	},
	270: {
		Fstart: int32(0x2cf2),
		Fend:   int32(0x2cf3),
	},
	271: {
		Fstart: int32(0x2d00),
		Fend:   int32(0x2d25),
	},
	272: {
		Fstart: int32(0x2d27),
		Fend:   int32(0x2d27),
	},
	273: {
		Fstart: int32(0x2d2d),
		Fend:   int32(0x2d2d),
	},
	274: {
		Fstart: int32(0x2d30),
		Fend:   int32(0x2d67),
	},
	275: {
		Fstart: int32(0x2d6f),
		Fend:   int32(0x2d6f),
	},
	276: {
		Fstart: int32(0x2d80),
		Fend:   int32(0x2d96),
	},
	277: {
		Fstart: int32(0x2da0),
		Fend:   int32(0x2da6),
	},
	278: {
		Fstart: int32(0x2da8),
		Fend:   int32(0x2dae),
	},
	279: {
		Fstart: int32(0x2db0),
		Fend:   int32(0x2db6),
	},
	280: {
		Fstart: int32(0x2db8),
		Fend:   int32(0x2dbe),
	},
	281: {
		Fstart: int32(0x2dc0),
		Fend:   int32(0x2dc6),
	},
	282: {
		Fstart: int32(0x2dc8),
		Fend:   int32(0x2dce),
	},
	283: {
		Fstart: int32(0x2dd0),
		Fend:   int32(0x2dd6),
	},
	284: {
		Fstart: int32(0x2dd8),
		Fend:   int32(0x2dde),
	},
	285: {
		Fstart: int32(0x3005),
		Fend:   int32(0x3007),
	},
	286: {
		Fstart: int32(0x3021),
		Fend:   int32(0x3029),
	},
	287: {
		Fstart: int32(0x3031),
		Fend:   int32(0x3035),
	},
	288: {
		Fstart: int32(0x3038),
		Fend:   int32(0x303c),
	},
	289: {
		Fstart: int32(0x3041),
		Fend:   int32(0x3096),
	},
	290: {
		Fstart: int32(0x309d),
		Fend:   int32(0x309f),
	},
	291: {
		Fstart: int32(0x30a1),
		Fend:   int32(0x30fa),
	},
	292: {
		Fstart: int32(0x30fc),
		Fend:   int32(0x30ff),
	},
	293: {
		Fstart: int32(0x3105),
		Fend:   int32(0x312f),
	},
	294: {
		Fstart: int32(0x3131),
		Fend:   int32(0x318e),
	},
	295: {
		Fstart: int32(0x31a0),
		Fend:   int32(0x31bf),
	},
	296: {
		Fstart: int32(0x31f0),
		Fend:   int32(0x31ff),
	},
	297: {
		Fstart: int32(0x3400),
		Fend:   int32(0x4dbf),
	},
	298: {
		Fstart: int32(0x4e00),
		Fend:   int32(0xa48c),
	},
	299: {
		Fstart: int32(0xa4d0),
		Fend:   int32(0xa4fd),
	},
	300: {
		Fstart: int32(0xa500),
		Fend:   int32(0xa60c),
	},
	301: {
		Fstart: int32(0xa610),
		Fend:   int32(0xa61f),
	},
	302: {
		Fstart: int32(0xa62a),
		Fend:   int32(0xa62b),
	},
	303: {
		Fstart: int32(0xa640),
		Fend:   int32(0xa66e),
	},
	304: {
		Fstart: int32(0xa67f),
		Fend:   int32(0xa69d),
	},
	305: {
		Fstart: int32(0xa6a0),
		Fend:   int32(0xa6ef),
	},
	306: {
		Fstart: int32(0xa717),
		Fend:   int32(0xa71f),
	},
	307: {
		Fstart: int32(0xa722),
		Fend:   int32(0xa788),
	},
	308: {
		Fstart: int32(0xa78b),
		Fend:   int32(0xa7ca),
	},
	309: {
		Fstart: int32(0xa7d0),
		Fend:   int32(0xa7d1),
	},
	310: {
		Fstart: int32(0xa7d3),
		Fend:   int32(0xa7d3),
	},
	311: {
		Fstart: int32(0xa7d5),
		Fend:   int32(0xa7d9),
	},
	312: {
		Fstart: int32(0xa7f2),
		Fend:   int32(0xa801),
	},
	313: {
		Fstart: int32(0xa803),
		Fend:   int32(0xa805),
	},
	314: {
		Fstart: int32(0xa807),
		Fend:   int32(0xa80a),
	},
	315: {
		Fstart: int32(0xa80c),
		Fend:   int32(0xa822),
	},
	316: {
		Fstart: int32(0xa840),
		Fend:   int32(0xa873),
	},
	317: {
		Fstart: int32(0xa882),
		Fend:   int32(0xa8b3),
	},
	318: {
		Fstart: int32(0xa8f2),
		Fend:   int32(0xa8f7),
	},
	319: {
		Fstart: int32(0xa8fb),
		Fend:   int32(0xa8fb),
	},
	320: {
		Fstart: int32(0xa8fd),
		Fend:   int32(0xa8fe),
	},
	321: {
		Fstart: int32(0xa90a),
		Fend:   int32(0xa925),
	},
	322: {
		Fstart: int32(0xa930),
		Fend:   int32(0xa946),
	},
	323: {
		Fstart: int32(0xa960),
		Fend:   int32(0xa97c),
	},
	324: {
		Fstart: int32(0xa984),
		Fend:   int32(0xa9b2),
	},
	325: {
		Fstart: int32(0xa9cf),
		Fend:   int32(0xa9cf),
	},
	326: {
		Fstart: int32(0xa9e0),
		Fend:   int32(0xa9e4),
	},
	327: {
		Fstart: int32(0xa9e6),
		Fend:   int32(0xa9ef),
	},
	328: {
		Fstart: int32(0xa9fa),
		Fend:   int32(0xa9fe),
	},
	329: {
		Fstart: int32(0xaa00),
		Fend:   int32(0xaa28),
	},
	330: {
		Fstart: int32(0xaa40),
		Fend:   int32(0xaa42),
	},
	331: {
		Fstart: int32(0xaa44),
		Fend:   int32(0xaa4b),
	},
	332: {
		Fstart: int32(0xaa60),
		Fend:   int32(0xaa76),
	},
	333: {
		Fstart: int32(0xaa7a),
		Fend:   int32(0xaa7a),
	},
	334: {
		Fstart: int32(0xaa7e),
		Fend:   int32(0xaaaf),
	},
	335: {
		Fstart: int32(0xaab1),
		Fend:   int32(0xaab1),
	},
	336: {
		Fstart: int32(0xaab5),
		Fend:   int32(0xaab6),
	},
	337: {
		Fstart: int32(0xaab9),
		Fend:   int32(0xaabd),
	},
	338: {
		Fstart: int32(0xaac0),
		Fend:   int32(0xaac0),
	},
	339: {
		Fstart: int32(0xaac2),
		Fend:   int32(0xaac2),
	},
	340: {
		Fstart: int32(0xaadb),
		Fend:   int32(0xaadd),
	},
	341: {
		Fstart: int32(0xaae0),
		Fend:   int32(0xaaea),
	},
	342: {
		Fstart: int32(0xaaf2),
		Fend:   int32(0xaaf4),
	},
	343: {
		Fstart: int32(0xab01),
		Fend:   int32(0xab06),
	},
	344: {
		Fstart: int32(0xab09),
		Fend:   int32(0xab0e),
	},
	345: {
		Fstart: int32(0xab11),
		Fend:   int32(0xab16),
	},
	346: {
		Fstart: int32(0xab20),
		Fend:   int32(0xab26),
	},
	347: {
		Fstart: int32(0xab28),
		Fend:   int32(0xab2e),
	},
	348: {
		Fstart: int32(0xab30),
		Fend:   int32(0xab5a),
	},
	349: {
		Fstart: int32(0xab5c),
		Fend:   int32(0xab69),
	},
	350: {
		Fstart: int32(0xab70),
		Fend:   int32(0xabe2),
	},
	351: {
		Fstart: int32(0xac00),
		Fend:   int32(0xd7a3),
	},
	352: {
		Fstart: int32(0xd7b0),
		Fend:   int32(0xd7c6),
	},
	353: {
		Fstart: int32(0xd7cb),
		Fend:   int32(0xd7fb),
	},
	354: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfa6d),
	},
	355: {
		Fstart: int32(0xfa70),
		Fend:   int32(0xfad9),
	},
	356: {
		Fstart: int32(0xfb00),
		Fend:   int32(0xfb06),
	},
	357: {
		Fstart: int32(0xfb13),
		Fend:   int32(0xfb17),
	},
	358: {
		Fstart: int32(0xfb1d),
		Fend:   int32(0xfb1d),
	},
	359: {
		Fstart: int32(0xfb1f),
		Fend:   int32(0xfb28),
	},
	360: {
		Fstart: int32(0xfb2a),
		Fend:   int32(0xfb36),
	},
	361: {
		Fstart: int32(0xfb38),
		Fend:   int32(0xfb3c),
	},
	362: {
		Fstart: int32(0xfb3e),
		Fend:   int32(0xfb3e),
	},
	363: {
		Fstart: int32(0xfb40),
		Fend:   int32(0xfb41),
	},
	364: {
		Fstart: int32(0xfb43),
		Fend:   int32(0xfb44),
	},
	365: {
		Fstart: int32(0xfb46),
		Fend:   int32(0xfbb1),
	},
	366: {
		Fstart: int32(0xfbd3),
		Fend:   int32(0xfc5d),
	},
	367: {
		Fstart: int32(0xfc64),
		Fend:   int32(0xfd3d),
	},
	368: {
		Fstart: int32(0xfd50),
		Fend:   int32(0xfd8f),
	},
	369: {
		Fstart: int32(0xfd92),
		Fend:   int32(0xfdc7),
	},
	370: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfdf9),
	},
	371: {
		Fstart: int32(0xfe71),
		Fend:   int32(0xfe71),
	},
	372: {
		Fstart: int32(0xfe73),
		Fend:   int32(0xfe73),
	},
	373: {
		Fstart: int32(0xfe77),
		Fend:   int32(0xfe77),
	},
	374: {
		Fstart: int32(0xfe79),
		Fend:   int32(0xfe79),
	},
	375: {
		Fstart: int32(0xfe7b),
		Fend:   int32(0xfe7b),
	},
	376: {
		Fstart: int32(0xfe7d),
		Fend:   int32(0xfe7d),
	},
	377: {
		Fstart: int32(0xfe7f),
		Fend:   int32(0xfefc),
	},
	378: {
		Fstart: int32(0xff21),
		Fend:   int32(0xff3a),
	},
	379: {
		Fstart: int32(0xff41),
		Fend:   int32(0xff5a),
	},
	380: {
		Fstart: int32(0xff66),
		Fend:   int32(0xff9d),
	},
	381: {
		Fstart: int32(0xffa0),
		Fend:   int32(0xffbe),
	},
	382: {
		Fstart: int32(0xffc2),
		Fend:   int32(0xffc7),
	},
	383: {
		Fstart: int32(0xffca),
		Fend:   int32(0xffcf),
	},
	384: {
		Fstart: int32(0xffd2),
		Fend:   int32(0xffd7),
	},
	385: {
		Fstart: int32(0xffda),
		Fend:   int32(0xffdc),
	},
	386: {
		Fstart: int32(0x10000),
		Fend:   int32(0x1000b),
	},
	387: {
		Fstart: int32(0x1000d),
		Fend:   int32(0x10026),
	},
	388: {
		Fstart: int32(0x10028),
		Fend:   int32(0x1003a),
	},
	389: {
		Fstart: int32(0x1003c),
		Fend:   int32(0x1003d),
	},
	390: {
		Fstart: int32(0x1003f),
		Fend:   int32(0x1004d),
	},
	391: {
		Fstart: int32(0x10050),
		Fend:   int32(0x1005d),
	},
	392: {
		Fstart: int32(0x10080),
		Fend:   int32(0x100fa),
	},
	393: {
		Fstart: int32(0x10140),
		Fend:   int32(0x10174),
	},
	394: {
		Fstart: int32(0x10280),
		Fend:   int32(0x1029c),
	},
	395: {
		Fstart: int32(0x102a0),
		Fend:   int32(0x102d0),
	},
	396: {
		Fstart: int32(0x10300),
		Fend:   int32(0x1031f),
	},
	397: {
		Fstart: int32(0x1032d),
		Fend:   int32(0x1034a),
	},
	398: {
		Fstart: int32(0x10350),
		Fend:   int32(0x10375),
	},
	399: {
		Fstart: int32(0x10380),
		Fend:   int32(0x1039d),
	},
	400: {
		Fstart: int32(0x103a0),
		Fend:   int32(0x103c3),
	},
	401: {
		Fstart: int32(0x103c8),
		Fend:   int32(0x103cf),
	},
	402: {
		Fstart: int32(0x103d1),
		Fend:   int32(0x103d5),
	},
	403: {
		Fstart: int32(0x10400),
		Fend:   int32(0x1049d),
	},
	404: {
		Fstart: int32(0x104b0),
		Fend:   int32(0x104d3),
	},
	405: {
		Fstart: int32(0x104d8),
		Fend:   int32(0x104fb),
	},
	406: {
		Fstart: int32(0x10500),
		Fend:   int32(0x10527),
	},
	407: {
		Fstart: int32(0x10530),
		Fend:   int32(0x10563),
	},
	408: {
		Fstart: int32(0x10570),
		Fend:   int32(0x1057a),
	},
	409: {
		Fstart: int32(0x1057c),
		Fend:   int32(0x1058a),
	},
	410: {
		Fstart: int32(0x1058c),
		Fend:   int32(0x10592),
	},
	411: {
		Fstart: int32(0x10594),
		Fend:   int32(0x10595),
	},
	412: {
		Fstart: int32(0x10597),
		Fend:   int32(0x105a1),
	},
	413: {
		Fstart: int32(0x105a3),
		Fend:   int32(0x105b1),
	},
	414: {
		Fstart: int32(0x105b3),
		Fend:   int32(0x105b9),
	},
	415: {
		Fstart: int32(0x105bb),
		Fend:   int32(0x105bc),
	},
	416: {
		Fstart: int32(0x10600),
		Fend:   int32(0x10736),
	},
	417: {
		Fstart: int32(0x10740),
		Fend:   int32(0x10755),
	},
	418: {
		Fstart: int32(0x10760),
		Fend:   int32(0x10767),
	},
	419: {
		Fstart: int32(0x10780),
		Fend:   int32(0x10785),
	},
	420: {
		Fstart: int32(0x10787),
		Fend:   int32(0x107b0),
	},
	421: {
		Fstart: int32(0x107b2),
		Fend:   int32(0x107ba),
	},
	422: {
		Fstart: int32(0x10800),
		Fend:   int32(0x10805),
	},
	423: {
		Fstart: int32(0x10808),
		Fend:   int32(0x10808),
	},
	424: {
		Fstart: int32(0x1080a),
		Fend:   int32(0x10835),
	},
	425: {
		Fstart: int32(0x10837),
		Fend:   int32(0x10838),
	},
	426: {
		Fstart: int32(0x1083c),
		Fend:   int32(0x1083c),
	},
	427: {
		Fstart: int32(0x1083f),
		Fend:   int32(0x10855),
	},
	428: {
		Fstart: int32(0x10860),
		Fend:   int32(0x10876),
	},
	429: {
		Fstart: int32(0x10880),
		Fend:   int32(0x1089e),
	},
	430: {
		Fstart: int32(0x108e0),
		Fend:   int32(0x108f2),
	},
	431: {
		Fstart: int32(0x108f4),
		Fend:   int32(0x108f5),
	},
	432: {
		Fstart: int32(0x10900),
		Fend:   int32(0x10915),
	},
	433: {
		Fstart: int32(0x10920),
		Fend:   int32(0x10939),
	},
	434: {
		Fstart: int32(0x10980),
		Fend:   int32(0x109b7),
	},
	435: {
		Fstart: int32(0x109be),
		Fend:   int32(0x109bf),
	},
	436: {
		Fstart: int32(0x10a00),
		Fend:   int32(0x10a00),
	},
	437: {
		Fstart: int32(0x10a10),
		Fend:   int32(0x10a13),
	},
	438: {
		Fstart: int32(0x10a15),
		Fend:   int32(0x10a17),
	},
	439: {
		Fstart: int32(0x10a19),
		Fend:   int32(0x10a35),
	},
	440: {
		Fstart: int32(0x10a60),
		Fend:   int32(0x10a7c),
	},
	441: {
		Fstart: int32(0x10a80),
		Fend:   int32(0x10a9c),
	},
	442: {
		Fstart: int32(0x10ac0),
		Fend:   int32(0x10ac7),
	},
	443: {
		Fstart: int32(0x10ac9),
		Fend:   int32(0x10ae4),
	},
	444: {
		Fstart: int32(0x10b00),
		Fend:   int32(0x10b35),
	},
	445: {
		Fstart: int32(0x10b40),
		Fend:   int32(0x10b55),
	},
	446: {
		Fstart: int32(0x10b60),
		Fend:   int32(0x10b72),
	},
	447: {
		Fstart: int32(0x10b80),
		Fend:   int32(0x10b91),
	},
	448: {
		Fstart: int32(0x10c00),
		Fend:   int32(0x10c48),
	},
	449: {
		Fstart: int32(0x10c80),
		Fend:   int32(0x10cb2),
	},
	450: {
		Fstart: int32(0x10cc0),
		Fend:   int32(0x10cf2),
	},
	451: {
		Fstart: int32(0x10d00),
		Fend:   int32(0x10d23),
	},
	452: {
		Fstart: int32(0x10e80),
		Fend:   int32(0x10ea9),
	},
	453: {
		Fstart: int32(0x10eb0),
		Fend:   int32(0x10eb1),
	},
	454: {
		Fstart: int32(0x10f00),
		Fend:   int32(0x10f1c),
	},
	455: {
		Fstart: int32(0x10f27),
		Fend:   int32(0x10f27),
	},
	456: {
		Fstart: int32(0x10f30),
		Fend:   int32(0x10f45),
	},
	457: {
		Fstart: int32(0x10f70),
		Fend:   int32(0x10f81),
	},
	458: {
		Fstart: int32(0x10fb0),
		Fend:   int32(0x10fc4),
	},
	459: {
		Fstart: int32(0x10fe0),
		Fend:   int32(0x10ff6),
	},
	460: {
		Fstart: int32(0x11003),
		Fend:   int32(0x11037),
	},
	461: {
		Fstart: int32(0x11071),
		Fend:   int32(0x11072),
	},
	462: {
		Fstart: int32(0x11075),
		Fend:   int32(0x11075),
	},
	463: {
		Fstart: int32(0x11083),
		Fend:   int32(0x110af),
	},
	464: {
		Fstart: int32(0x110d0),
		Fend:   int32(0x110e8),
	},
	465: {
		Fstart: int32(0x11103),
		Fend:   int32(0x11126),
	},
	466: {
		Fstart: int32(0x11144),
		Fend:   int32(0x11144),
	},
	467: {
		Fstart: int32(0x11147),
		Fend:   int32(0x11147),
	},
	468: {
		Fstart: int32(0x11150),
		Fend:   int32(0x11172),
	},
	469: {
		Fstart: int32(0x11176),
		Fend:   int32(0x11176),
	},
	470: {
		Fstart: int32(0x11183),
		Fend:   int32(0x111b2),
	},
	471: {
		Fstart: int32(0x111c1),
		Fend:   int32(0x111c4),
	},
	472: {
		Fstart: int32(0x111da),
		Fend:   int32(0x111da),
	},
	473: {
		Fstart: int32(0x111dc),
		Fend:   int32(0x111dc),
	},
	474: {
		Fstart: int32(0x11200),
		Fend:   int32(0x11211),
	},
	475: {
		Fstart: int32(0x11213),
		Fend:   int32(0x1122b),
	},
	476: {
		Fstart: int32(0x11280),
		Fend:   int32(0x11286),
	},
	477: {
		Fstart: int32(0x11288),
		Fend:   int32(0x11288),
	},
	478: {
		Fstart: int32(0x1128a),
		Fend:   int32(0x1128d),
	},
	479: {
		Fstart: int32(0x1128f),
		Fend:   int32(0x1129d),
	},
	480: {
		Fstart: int32(0x1129f),
		Fend:   int32(0x112a8),
	},
	481: {
		Fstart: int32(0x112b0),
		Fend:   int32(0x112de),
	},
	482: {
		Fstart: int32(0x11305),
		Fend:   int32(0x1130c),
	},
	483: {
		Fstart: int32(0x1130f),
		Fend:   int32(0x11310),
	},
	484: {
		Fstart: int32(0x11313),
		Fend:   int32(0x11328),
	},
	485: {
		Fstart: int32(0x1132a),
		Fend:   int32(0x11330),
	},
	486: {
		Fstart: int32(0x11332),
		Fend:   int32(0x11333),
	},
	487: {
		Fstart: int32(0x11335),
		Fend:   int32(0x11339),
	},
	488: {
		Fstart: int32(0x1133d),
		Fend:   int32(0x1133d),
	},
	489: {
		Fstart: int32(0x11350),
		Fend:   int32(0x11350),
	},
	490: {
		Fstart: int32(0x1135d),
		Fend:   int32(0x11361),
	},
	491: {
		Fstart: int32(0x11400),
		Fend:   int32(0x11434),
	},
	492: {
		Fstart: int32(0x11447),
		Fend:   int32(0x1144a),
	},
	493: {
		Fstart: int32(0x1145f),
		Fend:   int32(0x11461),
	},
	494: {
		Fstart: int32(0x11480),
		Fend:   int32(0x114af),
	},
	495: {
		Fstart: int32(0x114c4),
		Fend:   int32(0x114c5),
	},
	496: {
		Fstart: int32(0x114c7),
		Fend:   int32(0x114c7),
	},
	497: {
		Fstart: int32(0x11580),
		Fend:   int32(0x115ae),
	},
	498: {
		Fstart: int32(0x115d8),
		Fend:   int32(0x115db),
	},
	499: {
		Fstart: int32(0x11600),
		Fend:   int32(0x1162f),
	},
	500: {
		Fstart: int32(0x11644),
		Fend:   int32(0x11644),
	},
	501: {
		Fstart: int32(0x11680),
		Fend:   int32(0x116aa),
	},
	502: {
		Fstart: int32(0x116b8),
		Fend:   int32(0x116b8),
	},
	503: {
		Fstart: int32(0x11700),
		Fend:   int32(0x1171a),
	},
	504: {
		Fstart: int32(0x11740),
		Fend:   int32(0x11746),
	},
	505: {
		Fstart: int32(0x11800),
		Fend:   int32(0x1182b),
	},
	506: {
		Fstart: int32(0x118a0),
		Fend:   int32(0x118df),
	},
	507: {
		Fstart: int32(0x118ff),
		Fend:   int32(0x11906),
	},
	508: {
		Fstart: int32(0x11909),
		Fend:   int32(0x11909),
	},
	509: {
		Fstart: int32(0x1190c),
		Fend:   int32(0x11913),
	},
	510: {
		Fstart: int32(0x11915),
		Fend:   int32(0x11916),
	},
	511: {
		Fstart: int32(0x11918),
		Fend:   int32(0x1192f),
	},
	512: {
		Fstart: int32(0x1193f),
		Fend:   int32(0x1193f),
	},
	513: {
		Fstart: int32(0x11941),
		Fend:   int32(0x11941),
	},
	514: {
		Fstart: int32(0x119a0),
		Fend:   int32(0x119a7),
	},
	515: {
		Fstart: int32(0x119aa),
		Fend:   int32(0x119d0),
	},
	516: {
		Fstart: int32(0x119e1),
		Fend:   int32(0x119e1),
	},
	517: {
		Fstart: int32(0x119e3),
		Fend:   int32(0x119e3),
	},
	518: {
		Fstart: int32(0x11a00),
		Fend:   int32(0x11a00),
	},
	519: {
		Fstart: int32(0x11a0b),
		Fend:   int32(0x11a32),
	},
	520: {
		Fstart: int32(0x11a3a),
		Fend:   int32(0x11a3a),
	},
	521: {
		Fstart: int32(0x11a50),
		Fend:   int32(0x11a50),
	},
	522: {
		Fstart: int32(0x11a5c),
		Fend:   int32(0x11a89),
	},
	523: {
		Fstart: int32(0x11a9d),
		Fend:   int32(0x11a9d),
	},
	524: {
		Fstart: int32(0x11ab0),
		Fend:   int32(0x11af8),
	},
	525: {
		Fstart: int32(0x11c00),
		Fend:   int32(0x11c08),
	},
	526: {
		Fstart: int32(0x11c0a),
		Fend:   int32(0x11c2e),
	},
	527: {
		Fstart: int32(0x11c40),
		Fend:   int32(0x11c40),
	},
	528: {
		Fstart: int32(0x11c72),
		Fend:   int32(0x11c8f),
	},
	529: {
		Fstart: int32(0x11d00),
		Fend:   int32(0x11d06),
	},
	530: {
		Fstart: int32(0x11d08),
		Fend:   int32(0x11d09),
	},
	531: {
		Fstart: int32(0x11d0b),
		Fend:   int32(0x11d30),
	},
	532: {
		Fstart: int32(0x11d46),
		Fend:   int32(0x11d46),
	},
	533: {
		Fstart: int32(0x11d60),
		Fend:   int32(0x11d65),
	},
	534: {
		Fstart: int32(0x11d67),
		Fend:   int32(0x11d68),
	},
	535: {
		Fstart: int32(0x11d6a),
		Fend:   int32(0x11d89),
	},
	536: {
		Fstart: int32(0x11d98),
		Fend:   int32(0x11d98),
	},
	537: {
		Fstart: int32(0x11ee0),
		Fend:   int32(0x11ef2),
	},
	538: {
		Fstart: int32(0x11fb0),
		Fend:   int32(0x11fb0),
	},
	539: {
		Fstart: int32(0x12000),
		Fend:   int32(0x12399),
	},
	540: {
		Fstart: int32(0x12400),
		Fend:   int32(0x1246e),
	},
	541: {
		Fstart: int32(0x12480),
		Fend:   int32(0x12543),
	},
	542: {
		Fstart: int32(0x12f90),
		Fend:   int32(0x12ff0),
	},
	543: {
		Fstart: int32(0x13000),
		Fend:   int32(0x1342e),
	},
	544: {
		Fstart: int32(0x14400),
		Fend:   int32(0x14646),
	},
	545: {
		Fstart: int32(0x16800),
		Fend:   int32(0x16a38),
	},
	546: {
		Fstart: int32(0x16a40),
		Fend:   int32(0x16a5e),
	},
	547: {
		Fstart: int32(0x16a70),
		Fend:   int32(0x16abe),
	},
	548: {
		Fstart: int32(0x16ad0),
		Fend:   int32(0x16aed),
	},
	549: {
		Fstart: int32(0x16b00),
		Fend:   int32(0x16b2f),
	},
	550: {
		Fstart: int32(0x16b40),
		Fend:   int32(0x16b43),
	},
	551: {
		Fstart: int32(0x16b63),
		Fend:   int32(0x16b77),
	},
	552: {
		Fstart: int32(0x16b7d),
		Fend:   int32(0x16b8f),
	},
	553: {
		Fstart: int32(0x16e40),
		Fend:   int32(0x16e7f),
	},
	554: {
		Fstart: int32(0x16f00),
		Fend:   int32(0x16f4a),
	},
	555: {
		Fstart: int32(0x16f50),
		Fend:   int32(0x16f50),
	},
	556: {
		Fstart: int32(0x16f93),
		Fend:   int32(0x16f9f),
	},
	557: {
		Fstart: int32(0x16fe0),
		Fend:   int32(0x16fe1),
	},
	558: {
		Fstart: int32(0x16fe3),
		Fend:   int32(0x16fe3),
	},
	559: {
		Fstart: int32(0x17000),
		Fend:   int32(0x187f7),
	},
	560: {
		Fstart: int32(0x18800),
		Fend:   int32(0x18cd5),
	},
	561: {
		Fstart: int32(0x18d00),
		Fend:   int32(0x18d08),
	},
	562: {
		Fstart: int32(0x1aff0),
		Fend:   int32(0x1aff3),
	},
	563: {
		Fstart: int32(0x1aff5),
		Fend:   int32(0x1affb),
	},
	564: {
		Fstart: int32(0x1affd),
		Fend:   int32(0x1affe),
	},
	565: {
		Fstart: int32(0x1b000),
		Fend:   int32(0x1b122),
	},
	566: {
		Fstart: int32(0x1b150),
		Fend:   int32(0x1b152),
	},
	567: {
		Fstart: int32(0x1b164),
		Fend:   int32(0x1b167),
	},
	568: {
		Fstart: int32(0x1b170),
		Fend:   int32(0x1b2fb),
	},
	569: {
		Fstart: int32(0x1bc00),
		Fend:   int32(0x1bc6a),
	},
	570: {
		Fstart: int32(0x1bc70),
		Fend:   int32(0x1bc7c),
	},
	571: {
		Fstart: int32(0x1bc80),
		Fend:   int32(0x1bc88),
	},
	572: {
		Fstart: int32(0x1bc90),
		Fend:   int32(0x1bc99),
	},
	573: {
		Fstart: int32(0x1d400),
		Fend:   int32(0x1d454),
	},
	574: {
		Fstart: int32(0x1d456),
		Fend:   int32(0x1d49c),
	},
	575: {
		Fstart: int32(0x1d49e),
		Fend:   int32(0x1d49f),
	},
	576: {
		Fstart: int32(0x1d4a2),
		Fend:   int32(0x1d4a2),
	},
	577: {
		Fstart: int32(0x1d4a5),
		Fend:   int32(0x1d4a6),
	},
	578: {
		Fstart: int32(0x1d4a9),
		Fend:   int32(0x1d4ac),
	},
	579: {
		Fstart: int32(0x1d4ae),
		Fend:   int32(0x1d4b9),
	},
	580: {
		Fstart: int32(0x1d4bb),
		Fend:   int32(0x1d4bb),
	},
	581: {
		Fstart: int32(0x1d4bd),
		Fend:   int32(0x1d4c3),
	},
	582: {
		Fstart: int32(0x1d4c5),
		Fend:   int32(0x1d505),
	},
	583: {
		Fstart: int32(0x1d507),
		Fend:   int32(0x1d50a),
	},
	584: {
		Fstart: int32(0x1d50d),
		Fend:   int32(0x1d514),
	},
	585: {
		Fstart: int32(0x1d516),
		Fend:   int32(0x1d51c),
	},
	586: {
		Fstart: int32(0x1d51e),
		Fend:   int32(0x1d539),
	},
	587: {
		Fstart: int32(0x1d53b),
		Fend:   int32(0x1d53e),
	},
	588: {
		Fstart: int32(0x1d540),
		Fend:   int32(0x1d544),
	},
	589: {
		Fstart: int32(0x1d546),
		Fend:   int32(0x1d546),
	},
	590: {
		Fstart: int32(0x1d54a),
		Fend:   int32(0x1d550),
	},
	591: {
		Fstart: int32(0x1d552),
		Fend:   int32(0x1d6a5),
	},
	592: {
		Fstart: int32(0x1d6a8),
		Fend:   int32(0x1d6c0),
	},
	593: {
		Fstart: int32(0x1d6c2),
		Fend:   int32(0x1d6da),
	},
	594: {
		Fstart: int32(0x1d6dc),
		Fend:   int32(0x1d6fa),
	},
	595: {
		Fstart: int32(0x1d6fc),
		Fend:   int32(0x1d714),
	},
	596: {
		Fstart: int32(0x1d716),
		Fend:   int32(0x1d734),
	},
	597: {
		Fstart: int32(0x1d736),
		Fend:   int32(0x1d74e),
	},
	598: {
		Fstart: int32(0x1d750),
		Fend:   int32(0x1d76e),
	},
	599: {
		Fstart: int32(0x1d770),
		Fend:   int32(0x1d788),
	},
	600: {
		Fstart: int32(0x1d78a),
		Fend:   int32(0x1d7a8),
	},
	601: {
		Fstart: int32(0x1d7aa),
		Fend:   int32(0x1d7c2),
	},
	602: {
		Fstart: int32(0x1d7c4),
		Fend:   int32(0x1d7cb),
	},
	603: {
		Fstart: int32(0x1df00),
		Fend:   int32(0x1df1e),
	},
	604: {
		Fstart: int32(0x1e100),
		Fend:   int32(0x1e12c),
	},
	605: {
		Fstart: int32(0x1e137),
		Fend:   int32(0x1e13d),
	},
	606: {
		Fstart: int32(0x1e14e),
		Fend:   int32(0x1e14e),
	},
	607: {
		Fstart: int32(0x1e290),
		Fend:   int32(0x1e2ad),
	},
	608: {
		Fstart: int32(0x1e2c0),
		Fend:   int32(0x1e2eb),
	},
	609: {
		Fstart: int32(0x1e7e0),
		Fend:   int32(0x1e7e6),
	},
	610: {
		Fstart: int32(0x1e7e8),
		Fend:   int32(0x1e7eb),
	},
	611: {
		Fstart: int32(0x1e7ed),
		Fend:   int32(0x1e7ee),
	},
	612: {
		Fstart: int32(0x1e7f0),
		Fend:   int32(0x1e7fe),
	},
	613: {
		Fstart: int32(0x1e800),
		Fend:   int32(0x1e8c4),
	},
	614: {
		Fstart: int32(0x1e900),
		Fend:   int32(0x1e943),
	},
	615: {
		Fstart: int32(0x1e94b),
		Fend:   int32(0x1e94b),
	},
	616: {
		Fstart: int32(0x1ee00),
		Fend:   int32(0x1ee03),
	},
	617: {
		Fstart: int32(0x1ee05),
		Fend:   int32(0x1ee1f),
	},
	618: {
		Fstart: int32(0x1ee21),
		Fend:   int32(0x1ee22),
	},
	619: {
		Fstart: int32(0x1ee24),
		Fend:   int32(0x1ee24),
	},
	620: {
		Fstart: int32(0x1ee27),
		Fend:   int32(0x1ee27),
	},
	621: {
		Fstart: int32(0x1ee29),
		Fend:   int32(0x1ee32),
	},
	622: {
		Fstart: int32(0x1ee34),
		Fend:   int32(0x1ee37),
	},
	623: {
		Fstart: int32(0x1ee39),
		Fend:   int32(0x1ee39),
	},
	624: {
		Fstart: int32(0x1ee3b),
		Fend:   int32(0x1ee3b),
	},
	625: {
		Fstart: int32(0x1ee42),
		Fend:   int32(0x1ee42),
	},
	626: {
		Fstart: int32(0x1ee47),
		Fend:   int32(0x1ee47),
	},
	627: {
		Fstart: int32(0x1ee49),
		Fend:   int32(0x1ee49),
	},
	628: {
		Fstart: int32(0x1ee4b),
		Fend:   int32(0x1ee4b),
	},
	629: {
		Fstart: int32(0x1ee4d),
		Fend:   int32(0x1ee4f),
	},
	630: {
		Fstart: int32(0x1ee51),
		Fend:   int32(0x1ee52),
	},
	631: {
		Fstart: int32(0x1ee54),
		Fend:   int32(0x1ee54),
	},
	632: {
		Fstart: int32(0x1ee57),
		Fend:   int32(0x1ee57),
	},
	633: {
		Fstart: int32(0x1ee59),
		Fend:   int32(0x1ee59),
	},
	634: {
		Fstart: int32(0x1ee5b),
		Fend:   int32(0x1ee5b),
	},
	635: {
		Fstart: int32(0x1ee5d),
		Fend:   int32(0x1ee5d),
	},
	636: {
		Fstart: int32(0x1ee5f),
		Fend:   int32(0x1ee5f),
	},
	637: {
		Fstart: int32(0x1ee61),
		Fend:   int32(0x1ee62),
	},
	638: {
		Fstart: int32(0x1ee64),
		Fend:   int32(0x1ee64),
	},
	639: {
		Fstart: int32(0x1ee67),
		Fend:   int32(0x1ee6a),
	},
	640: {
		Fstart: int32(0x1ee6c),
		Fend:   int32(0x1ee72),
	},
	641: {
		Fstart: int32(0x1ee74),
		Fend:   int32(0x1ee77),
	},
	642: {
		Fstart: int32(0x1ee79),
		Fend:   int32(0x1ee7c),
	},
	643: {
		Fstart: int32(0x1ee7e),
		Fend:   int32(0x1ee7e),
	},
	644: {
		Fstart: int32(0x1ee80),
		Fend:   int32(0x1ee89),
	},
	645: {
		Fstart: int32(0x1ee8b),
		Fend:   int32(0x1ee9b),
	},
	646: {
		Fstart: int32(0x1eea1),
		Fend:   int32(0x1eea3),
	},
	647: {
		Fstart: int32(0x1eea5),
		Fend:   int32(0x1eea9),
	},
	648: {
		Fstart: int32(0x1eeab),
		Fend:   int32(0x1eebb),
	},
	649: {
		Fstart: int32(0x20000),
		Fend:   int32(0x2a6df),
	},
	650: {
		Fstart: int32(0x2a700),
		Fend:   int32(0x2b738),
	},
	651: {
		Fstart: int32(0x2b740),
		Fend:   int32(0x2b81d),
	},
	652: {
		Fstart: int32(0x2b820),
		Fend:   int32(0x2cea1),
	},
	653: {
		Fstart: int32(0x2ceb0),
		Fend:   int32(0x2ebe0),
	},
	654: {
		Fstart: int32(0x2f800),
		Fend:   int32(0x2fa1d),
	},
	655: {
		Fstart: int32(0x30000),
		Fend:   int32(0x3134a),
	},
}

var sym_identifier_character_set_3 = [763]TSCharacterRange{
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
		Fstart: int32(0xaa),
		Fend:   int32(0xaa),
	},
	5: {
		Fstart: int32(0xb5),
		Fend:   int32(0xb5),
	},
	6: {
		Fstart: int32(0xb7),
		Fend:   int32(0xb7),
	},
	7: {
		Fstart: int32(0xba),
		Fend:   int32(0xba),
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
		Fend:   int32(0x2c1),
	},
	11: {
		Fstart: int32(0x2c6),
		Fend:   int32(0x2d1),
	},
	12: {
		Fstart: int32(0x2e0),
		Fend:   int32(0x2e4),
	},
	13: {
		Fstart: int32(0x2ec),
		Fend:   int32(0x2ec),
	},
	14: {
		Fstart: int32(0x2ee),
		Fend:   int32(0x2ee),
	},
	15: {
		Fstart: int32(0x300),
		Fend:   int32(0x374),
	},
	16: {
		Fstart: int32(0x376),
		Fend:   int32(0x377),
	},
	17: {
		Fstart: int32(0x37b),
		Fend:   int32(0x37d),
	},
	18: {
		Fstart: int32(0x37f),
		Fend:   int32(0x37f),
	},
	19: {
		Fstart: int32(0x386),
		Fend:   int32(0x38a),
	},
	20: {
		Fstart: int32(0x38c),
		Fend:   int32(0x38c),
	},
	21: {
		Fstart: int32(0x38e),
		Fend:   int32(0x3a1),
	},
	22: {
		Fstart: int32(0x3a3),
		Fend:   int32(0x3f5),
	},
	23: {
		Fstart: int32(0x3f7),
		Fend:   int32(0x481),
	},
	24: {
		Fstart: int32(0x483),
		Fend:   int32(0x487),
	},
	25: {
		Fstart: int32(0x48a),
		Fend:   int32(0x52f),
	},
	26: {
		Fstart: int32(0x531),
		Fend:   int32(0x556),
	},
	27: {
		Fstart: int32(0x559),
		Fend:   int32(0x559),
	},
	28: {
		Fstart: int32(0x560),
		Fend:   int32(0x588),
	},
	29: {
		Fstart: int32(0x591),
		Fend:   int32(0x5bd),
	},
	30: {
		Fstart: int32(0x5bf),
		Fend:   int32(0x5bf),
	},
	31: {
		Fstart: int32(0x5c1),
		Fend:   int32(0x5c2),
	},
	32: {
		Fstart: int32(0x5c4),
		Fend:   int32(0x5c5),
	},
	33: {
		Fstart: int32(0x5c7),
		Fend:   int32(0x5c7),
	},
	34: {
		Fstart: int32(0x5d0),
		Fend:   int32(0x5ea),
	},
	35: {
		Fstart: int32(0x5ef),
		Fend:   int32(0x5f2),
	},
	36: {
		Fstart: int32(0x610),
		Fend:   int32(0x61a),
	},
	37: {
		Fstart: int32(0x620),
		Fend:   int32(0x669),
	},
	38: {
		Fstart: int32(0x66e),
		Fend:   int32(0x6d3),
	},
	39: {
		Fstart: int32(0x6d5),
		Fend:   int32(0x6dc),
	},
	40: {
		Fstart: int32(0x6df),
		Fend:   int32(0x6e8),
	},
	41: {
		Fstart: int32(0x6ea),
		Fend:   int32(0x6fc),
	},
	42: {
		Fstart: int32(0x6ff),
		Fend:   int32(0x6ff),
	},
	43: {
		Fstart: int32(0x710),
		Fend:   int32(0x74a),
	},
	44: {
		Fstart: int32(0x74d),
		Fend:   int32(0x7b1),
	},
	45: {
		Fstart: int32(0x7c0),
		Fend:   int32(0x7f5),
	},
	46: {
		Fstart: int32(0x7fa),
		Fend:   int32(0x7fa),
	},
	47: {
		Fstart: int32(0x7fd),
		Fend:   int32(0x7fd),
	},
	48: {
		Fstart: int32(0x800),
		Fend:   int32(0x82d),
	},
	49: {
		Fstart: int32(0x840),
		Fend:   int32(0x85b),
	},
	50: {
		Fstart: int32(0x860),
		Fend:   int32(0x86a),
	},
	51: {
		Fstart: int32(0x870),
		Fend:   int32(0x887),
	},
	52: {
		Fstart: int32(0x889),
		Fend:   int32(0x88e),
	},
	53: {
		Fstart: int32(0x898),
		Fend:   int32(0x8e1),
	},
	54: {
		Fstart: int32(0x8e3),
		Fend:   int32(0x963),
	},
	55: {
		Fstart: int32(0x966),
		Fend:   int32(0x96f),
	},
	56: {
		Fstart: int32(0x971),
		Fend:   int32(0x983),
	},
	57: {
		Fstart: int32(0x985),
		Fend:   int32(0x98c),
	},
	58: {
		Fstart: int32(0x98f),
		Fend:   int32(0x990),
	},
	59: {
		Fstart: int32(0x993),
		Fend:   int32(0x9a8),
	},
	60: {
		Fstart: int32(0x9aa),
		Fend:   int32(0x9b0),
	},
	61: {
		Fstart: int32(0x9b2),
		Fend:   int32(0x9b2),
	},
	62: {
		Fstart: int32(0x9b6),
		Fend:   int32(0x9b9),
	},
	63: {
		Fstart: int32(0x9bc),
		Fend:   int32(0x9c4),
	},
	64: {
		Fstart: int32(0x9c7),
		Fend:   int32(0x9c8),
	},
	65: {
		Fstart: int32(0x9cb),
		Fend:   int32(0x9ce),
	},
	66: {
		Fstart: int32(0x9d7),
		Fend:   int32(0x9d7),
	},
	67: {
		Fstart: int32(0x9dc),
		Fend:   int32(0x9dd),
	},
	68: {
		Fstart: int32(0x9df),
		Fend:   int32(0x9e3),
	},
	69: {
		Fstart: int32(0x9e6),
		Fend:   int32(0x9f1),
	},
	70: {
		Fstart: int32(0x9fc),
		Fend:   int32(0x9fc),
	},
	71: {
		Fstart: int32(0x9fe),
		Fend:   int32(0x9fe),
	},
	72: {
		Fstart: int32(0xa01),
		Fend:   int32(0xa03),
	},
	73: {
		Fstart: int32(0xa05),
		Fend:   int32(0xa0a),
	},
	74: {
		Fstart: int32(0xa0f),
		Fend:   int32(0xa10),
	},
	75: {
		Fstart: int32(0xa13),
		Fend:   int32(0xa28),
	},
	76: {
		Fstart: int32(0xa2a),
		Fend:   int32(0xa30),
	},
	77: {
		Fstart: int32(0xa32),
		Fend:   int32(0xa33),
	},
	78: {
		Fstart: int32(0xa35),
		Fend:   int32(0xa36),
	},
	79: {
		Fstart: int32(0xa38),
		Fend:   int32(0xa39),
	},
	80: {
		Fstart: int32(0xa3c),
		Fend:   int32(0xa3c),
	},
	81: {
		Fstart: int32(0xa3e),
		Fend:   int32(0xa42),
	},
	82: {
		Fstart: int32(0xa47),
		Fend:   int32(0xa48),
	},
	83: {
		Fstart: int32(0xa4b),
		Fend:   int32(0xa4d),
	},
	84: {
		Fstart: int32(0xa51),
		Fend:   int32(0xa51),
	},
	85: {
		Fstart: int32(0xa59),
		Fend:   int32(0xa5c),
	},
	86: {
		Fstart: int32(0xa5e),
		Fend:   int32(0xa5e),
	},
	87: {
		Fstart: int32(0xa66),
		Fend:   int32(0xa75),
	},
	88: {
		Fstart: int32(0xa81),
		Fend:   int32(0xa83),
	},
	89: {
		Fstart: int32(0xa85),
		Fend:   int32(0xa8d),
	},
	90: {
		Fstart: int32(0xa8f),
		Fend:   int32(0xa91),
	},
	91: {
		Fstart: int32(0xa93),
		Fend:   int32(0xaa8),
	},
	92: {
		Fstart: int32(0xaaa),
		Fend:   int32(0xab0),
	},
	93: {
		Fstart: int32(0xab2),
		Fend:   int32(0xab3),
	},
	94: {
		Fstart: int32(0xab5),
		Fend:   int32(0xab9),
	},
	95: {
		Fstart: int32(0xabc),
		Fend:   int32(0xac5),
	},
	96: {
		Fstart: int32(0xac7),
		Fend:   int32(0xac9),
	},
	97: {
		Fstart: int32(0xacb),
		Fend:   int32(0xacd),
	},
	98: {
		Fstart: int32(0xad0),
		Fend:   int32(0xad0),
	},
	99: {
		Fstart: int32(0xae0),
		Fend:   int32(0xae3),
	},
	100: {
		Fstart: int32(0xae6),
		Fend:   int32(0xaef),
	},
	101: {
		Fstart: int32(0xaf9),
		Fend:   int32(0xaff),
	},
	102: {
		Fstart: int32(0xb01),
		Fend:   int32(0xb03),
	},
	103: {
		Fstart: int32(0xb05),
		Fend:   int32(0xb0c),
	},
	104: {
		Fstart: int32(0xb0f),
		Fend:   int32(0xb10),
	},
	105: {
		Fstart: int32(0xb13),
		Fend:   int32(0xb28),
	},
	106: {
		Fstart: int32(0xb2a),
		Fend:   int32(0xb30),
	},
	107: {
		Fstart: int32(0xb32),
		Fend:   int32(0xb33),
	},
	108: {
		Fstart: int32(0xb35),
		Fend:   int32(0xb39),
	},
	109: {
		Fstart: int32(0xb3c),
		Fend:   int32(0xb44),
	},
	110: {
		Fstart: int32(0xb47),
		Fend:   int32(0xb48),
	},
	111: {
		Fstart: int32(0xb4b),
		Fend:   int32(0xb4d),
	},
	112: {
		Fstart: int32(0xb55),
		Fend:   int32(0xb57),
	},
	113: {
		Fstart: int32(0xb5c),
		Fend:   int32(0xb5d),
	},
	114: {
		Fstart: int32(0xb5f),
		Fend:   int32(0xb63),
	},
	115: {
		Fstart: int32(0xb66),
		Fend:   int32(0xb6f),
	},
	116: {
		Fstart: int32(0xb71),
		Fend:   int32(0xb71),
	},
	117: {
		Fstart: int32(0xb82),
		Fend:   int32(0xb83),
	},
	118: {
		Fstart: int32(0xb85),
		Fend:   int32(0xb8a),
	},
	119: {
		Fstart: int32(0xb8e),
		Fend:   int32(0xb90),
	},
	120: {
		Fstart: int32(0xb92),
		Fend:   int32(0xb95),
	},
	121: {
		Fstart: int32(0xb99),
		Fend:   int32(0xb9a),
	},
	122: {
		Fstart: int32(0xb9c),
		Fend:   int32(0xb9c),
	},
	123: {
		Fstart: int32(0xb9e),
		Fend:   int32(0xb9f),
	},
	124: {
		Fstart: int32(0xba3),
		Fend:   int32(0xba4),
	},
	125: {
		Fstart: int32(0xba8),
		Fend:   int32(0xbaa),
	},
	126: {
		Fstart: int32(0xbae),
		Fend:   int32(0xbb9),
	},
	127: {
		Fstart: int32(0xbbe),
		Fend:   int32(0xbc2),
	},
	128: {
		Fstart: int32(0xbc6),
		Fend:   int32(0xbc8),
	},
	129: {
		Fstart: int32(0xbca),
		Fend:   int32(0xbcd),
	},
	130: {
		Fstart: int32(0xbd0),
		Fend:   int32(0xbd0),
	},
	131: {
		Fstart: int32(0xbd7),
		Fend:   int32(0xbd7),
	},
	132: {
		Fstart: int32(0xbe6),
		Fend:   int32(0xbef),
	},
	133: {
		Fstart: int32(0xc00),
		Fend:   int32(0xc0c),
	},
	134: {
		Fstart: int32(0xc0e),
		Fend:   int32(0xc10),
	},
	135: {
		Fstart: int32(0xc12),
		Fend:   int32(0xc28),
	},
	136: {
		Fstart: int32(0xc2a),
		Fend:   int32(0xc39),
	},
	137: {
		Fstart: int32(0xc3c),
		Fend:   int32(0xc44),
	},
	138: {
		Fstart: int32(0xc46),
		Fend:   int32(0xc48),
	},
	139: {
		Fstart: int32(0xc4a),
		Fend:   int32(0xc4d),
	},
	140: {
		Fstart: int32(0xc55),
		Fend:   int32(0xc56),
	},
	141: {
		Fstart: int32(0xc58),
		Fend:   int32(0xc5a),
	},
	142: {
		Fstart: int32(0xc5d),
		Fend:   int32(0xc5d),
	},
	143: {
		Fstart: int32(0xc60),
		Fend:   int32(0xc63),
	},
	144: {
		Fstart: int32(0xc66),
		Fend:   int32(0xc6f),
	},
	145: {
		Fstart: int32(0xc80),
		Fend:   int32(0xc83),
	},
	146: {
		Fstart: int32(0xc85),
		Fend:   int32(0xc8c),
	},
	147: {
		Fstart: int32(0xc8e),
		Fend:   int32(0xc90),
	},
	148: {
		Fstart: int32(0xc92),
		Fend:   int32(0xca8),
	},
	149: {
		Fstart: int32(0xcaa),
		Fend:   int32(0xcb3),
	},
	150: {
		Fstart: int32(0xcb5),
		Fend:   int32(0xcb9),
	},
	151: {
		Fstart: int32(0xcbc),
		Fend:   int32(0xcc4),
	},
	152: {
		Fstart: int32(0xcc6),
		Fend:   int32(0xcc8),
	},
	153: {
		Fstart: int32(0xcca),
		Fend:   int32(0xccd),
	},
	154: {
		Fstart: int32(0xcd5),
		Fend:   int32(0xcd6),
	},
	155: {
		Fstart: int32(0xcdd),
		Fend:   int32(0xcde),
	},
	156: {
		Fstart: int32(0xce0),
		Fend:   int32(0xce3),
	},
	157: {
		Fstart: int32(0xce6),
		Fend:   int32(0xcef),
	},
	158: {
		Fstart: int32(0xcf1),
		Fend:   int32(0xcf2),
	},
	159: {
		Fstart: int32(0xd00),
		Fend:   int32(0xd0c),
	},
	160: {
		Fstart: int32(0xd0e),
		Fend:   int32(0xd10),
	},
	161: {
		Fstart: int32(0xd12),
		Fend:   int32(0xd44),
	},
	162: {
		Fstart: int32(0xd46),
		Fend:   int32(0xd48),
	},
	163: {
		Fstart: int32(0xd4a),
		Fend:   int32(0xd4e),
	},
	164: {
		Fstart: int32(0xd54),
		Fend:   int32(0xd57),
	},
	165: {
		Fstart: int32(0xd5f),
		Fend:   int32(0xd63),
	},
	166: {
		Fstart: int32(0xd66),
		Fend:   int32(0xd6f),
	},
	167: {
		Fstart: int32(0xd7a),
		Fend:   int32(0xd7f),
	},
	168: {
		Fstart: int32(0xd81),
		Fend:   int32(0xd83),
	},
	169: {
		Fstart: int32(0xd85),
		Fend:   int32(0xd96),
	},
	170: {
		Fstart: int32(0xd9a),
		Fend:   int32(0xdb1),
	},
	171: {
		Fstart: int32(0xdb3),
		Fend:   int32(0xdbb),
	},
	172: {
		Fstart: int32(0xdbd),
		Fend:   int32(0xdbd),
	},
	173: {
		Fstart: int32(0xdc0),
		Fend:   int32(0xdc6),
	},
	174: {
		Fstart: int32(0xdca),
		Fend:   int32(0xdca),
	},
	175: {
		Fstart: int32(0xdcf),
		Fend:   int32(0xdd4),
	},
	176: {
		Fstart: int32(0xdd6),
		Fend:   int32(0xdd6),
	},
	177: {
		Fstart: int32(0xdd8),
		Fend:   int32(0xddf),
	},
	178: {
		Fstart: int32(0xde6),
		Fend:   int32(0xdef),
	},
	179: {
		Fstart: int32(0xdf2),
		Fend:   int32(0xdf3),
	},
	180: {
		Fstart: int32(0xe01),
		Fend:   int32(0xe3a),
	},
	181: {
		Fstart: int32(0xe40),
		Fend:   int32(0xe4e),
	},
	182: {
		Fstart: int32(0xe50),
		Fend:   int32(0xe59),
	},
	183: {
		Fstart: int32(0xe81),
		Fend:   int32(0xe82),
	},
	184: {
		Fstart: int32(0xe84),
		Fend:   int32(0xe84),
	},
	185: {
		Fstart: int32(0xe86),
		Fend:   int32(0xe8a),
	},
	186: {
		Fstart: int32(0xe8c),
		Fend:   int32(0xea3),
	},
	187: {
		Fstart: int32(0xea5),
		Fend:   int32(0xea5),
	},
	188: {
		Fstart: int32(0xea7),
		Fend:   int32(0xebd),
	},
	189: {
		Fstart: int32(0xec0),
		Fend:   int32(0xec4),
	},
	190: {
		Fstart: int32(0xec6),
		Fend:   int32(0xec6),
	},
	191: {
		Fstart: int32(0xec8),
		Fend:   int32(0xecd),
	},
	192: {
		Fstart: int32(0xed0),
		Fend:   int32(0xed9),
	},
	193: {
		Fstart: int32(0xedc),
		Fend:   int32(0xedf),
	},
	194: {
		Fstart: int32(0xf00),
		Fend:   int32(0xf00),
	},
	195: {
		Fstart: int32(0xf18),
		Fend:   int32(0xf19),
	},
	196: {
		Fstart: int32(0xf20),
		Fend:   int32(0xf29),
	},
	197: {
		Fstart: int32(0xf35),
		Fend:   int32(0xf35),
	},
	198: {
		Fstart: int32(0xf37),
		Fend:   int32(0xf37),
	},
	199: {
		Fstart: int32(0xf39),
		Fend:   int32(0xf39),
	},
	200: {
		Fstart: int32(0xf3e),
		Fend:   int32(0xf47),
	},
	201: {
		Fstart: int32(0xf49),
		Fend:   int32(0xf6c),
	},
	202: {
		Fstart: int32(0xf71),
		Fend:   int32(0xf84),
	},
	203: {
		Fstart: int32(0xf86),
		Fend:   int32(0xf97),
	},
	204: {
		Fstart: int32(0xf99),
		Fend:   int32(0xfbc),
	},
	205: {
		Fstart: int32(0xfc6),
		Fend:   int32(0xfc6),
	},
	206: {
		Fstart: int32(0x1000),
		Fend:   int32(0x1049),
	},
	207: {
		Fstart: int32(0x1050),
		Fend:   int32(0x109d),
	},
	208: {
		Fstart: int32(0x10a0),
		Fend:   int32(0x10c5),
	},
	209: {
		Fstart: int32(0x10c7),
		Fend:   int32(0x10c7),
	},
	210: {
		Fstart: int32(0x10cd),
		Fend:   int32(0x10cd),
	},
	211: {
		Fstart: int32(0x10d0),
		Fend:   int32(0x10fa),
	},
	212: {
		Fstart: int32(0x10fc),
		Fend:   int32(0x1248),
	},
	213: {
		Fstart: int32(0x124a),
		Fend:   int32(0x124d),
	},
	214: {
		Fstart: int32(0x1250),
		Fend:   int32(0x1256),
	},
	215: {
		Fstart: int32(0x1258),
		Fend:   int32(0x1258),
	},
	216: {
		Fstart: int32(0x125a),
		Fend:   int32(0x125d),
	},
	217: {
		Fstart: int32(0x1260),
		Fend:   int32(0x1288),
	},
	218: {
		Fstart: int32(0x128a),
		Fend:   int32(0x128d),
	},
	219: {
		Fstart: int32(0x1290),
		Fend:   int32(0x12b0),
	},
	220: {
		Fstart: int32(0x12b2),
		Fend:   int32(0x12b5),
	},
	221: {
		Fstart: int32(0x12b8),
		Fend:   int32(0x12be),
	},
	222: {
		Fstart: int32(0x12c0),
		Fend:   int32(0x12c0),
	},
	223: {
		Fstart: int32(0x12c2),
		Fend:   int32(0x12c5),
	},
	224: {
		Fstart: int32(0x12c8),
		Fend:   int32(0x12d6),
	},
	225: {
		Fstart: int32(0x12d8),
		Fend:   int32(0x1310),
	},
	226: {
		Fstart: int32(0x1312),
		Fend:   int32(0x1315),
	},
	227: {
		Fstart: int32(0x1318),
		Fend:   int32(0x135a),
	},
	228: {
		Fstart: int32(0x135d),
		Fend:   int32(0x135f),
	},
	229: {
		Fstart: int32(0x1369),
		Fend:   int32(0x1371),
	},
	230: {
		Fstart: int32(0x1380),
		Fend:   int32(0x138f),
	},
	231: {
		Fstart: int32(0x13a0),
		Fend:   int32(0x13f5),
	},
	232: {
		Fstart: int32(0x13f8),
		Fend:   int32(0x13fd),
	},
	233: {
		Fstart: int32(0x1401),
		Fend:   int32(0x166c),
	},
	234: {
		Fstart: int32(0x166f),
		Fend:   int32(0x167f),
	},
	235: {
		Fstart: int32(0x1681),
		Fend:   int32(0x169a),
	},
	236: {
		Fstart: int32(0x16a0),
		Fend:   int32(0x16ea),
	},
	237: {
		Fstart: int32(0x16ee),
		Fend:   int32(0x16f8),
	},
	238: {
		Fstart: int32(0x1700),
		Fend:   int32(0x1715),
	},
	239: {
		Fstart: int32(0x171f),
		Fend:   int32(0x1734),
	},
	240: {
		Fstart: int32(0x1740),
		Fend:   int32(0x1753),
	},
	241: {
		Fstart: int32(0x1760),
		Fend:   int32(0x176c),
	},
	242: {
		Fstart: int32(0x176e),
		Fend:   int32(0x1770),
	},
	243: {
		Fstart: int32(0x1772),
		Fend:   int32(0x1773),
	},
	244: {
		Fstart: int32(0x1780),
		Fend:   int32(0x17d3),
	},
	245: {
		Fstart: int32(0x17d7),
		Fend:   int32(0x17d7),
	},
	246: {
		Fstart: int32(0x17dc),
		Fend:   int32(0x17dd),
	},
	247: {
		Fstart: int32(0x17e0),
		Fend:   int32(0x17e9),
	},
	248: {
		Fstart: int32(0x180b),
		Fend:   int32(0x180d),
	},
	249: {
		Fstart: int32(0x180f),
		Fend:   int32(0x1819),
	},
	250: {
		Fstart: int32(0x1820),
		Fend:   int32(0x1878),
	},
	251: {
		Fstart: int32(0x1880),
		Fend:   int32(0x18aa),
	},
	252: {
		Fstart: int32(0x18b0),
		Fend:   int32(0x18f5),
	},
	253: {
		Fstart: int32(0x1900),
		Fend:   int32(0x191e),
	},
	254: {
		Fstart: int32(0x1920),
		Fend:   int32(0x192b),
	},
	255: {
		Fstart: int32(0x1930),
		Fend:   int32(0x193b),
	},
	256: {
		Fstart: int32(0x1946),
		Fend:   int32(0x196d),
	},
	257: {
		Fstart: int32(0x1970),
		Fend:   int32(0x1974),
	},
	258: {
		Fstart: int32(0x1980),
		Fend:   int32(0x19ab),
	},
	259: {
		Fstart: int32(0x19b0),
		Fend:   int32(0x19c9),
	},
	260: {
		Fstart: int32(0x19d0),
		Fend:   int32(0x19da),
	},
	261: {
		Fstart: int32(0x1a00),
		Fend:   int32(0x1a1b),
	},
	262: {
		Fstart: int32(0x1a20),
		Fend:   int32(0x1a5e),
	},
	263: {
		Fstart: int32(0x1a60),
		Fend:   int32(0x1a7c),
	},
	264: {
		Fstart: int32(0x1a7f),
		Fend:   int32(0x1a89),
	},
	265: {
		Fstart: int32(0x1a90),
		Fend:   int32(0x1a99),
	},
	266: {
		Fstart: int32(0x1aa7),
		Fend:   int32(0x1aa7),
	},
	267: {
		Fstart: int32(0x1ab0),
		Fend:   int32(0x1abd),
	},
	268: {
		Fstart: int32(0x1abf),
		Fend:   int32(0x1ace),
	},
	269: {
		Fstart: int32(0x1b00),
		Fend:   int32(0x1b4c),
	},
	270: {
		Fstart: int32(0x1b50),
		Fend:   int32(0x1b59),
	},
	271: {
		Fstart: int32(0x1b6b),
		Fend:   int32(0x1b73),
	},
	272: {
		Fstart: int32(0x1b80),
		Fend:   int32(0x1bf3),
	},
	273: {
		Fstart: int32(0x1c00),
		Fend:   int32(0x1c37),
	},
	274: {
		Fstart: int32(0x1c40),
		Fend:   int32(0x1c49),
	},
	275: {
		Fstart: int32(0x1c4d),
		Fend:   int32(0x1c7d),
	},
	276: {
		Fstart: int32(0x1c80),
		Fend:   int32(0x1c88),
	},
	277: {
		Fstart: int32(0x1c90),
		Fend:   int32(0x1cba),
	},
	278: {
		Fstart: int32(0x1cbd),
		Fend:   int32(0x1cbf),
	},
	279: {
		Fstart: int32(0x1cd0),
		Fend:   int32(0x1cd2),
	},
	280: {
		Fstart: int32(0x1cd4),
		Fend:   int32(0x1cfa),
	},
	281: {
		Fstart: int32(0x1d00),
		Fend:   int32(0x1f15),
	},
	282: {
		Fstart: int32(0x1f18),
		Fend:   int32(0x1f1d),
	},
	283: {
		Fstart: int32(0x1f20),
		Fend:   int32(0x1f45),
	},
	284: {
		Fstart: int32(0x1f48),
		Fend:   int32(0x1f4d),
	},
	285: {
		Fstart: int32(0x1f50),
		Fend:   int32(0x1f57),
	},
	286: {
		Fstart: int32(0x1f59),
		Fend:   int32(0x1f59),
	},
	287: {
		Fstart: int32(0x1f5b),
		Fend:   int32(0x1f5b),
	},
	288: {
		Fstart: int32(0x1f5d),
		Fend:   int32(0x1f5d),
	},
	289: {
		Fstart: int32(0x1f5f),
		Fend:   int32(0x1f7d),
	},
	290: {
		Fstart: int32(0x1f80),
		Fend:   int32(0x1fb4),
	},
	291: {
		Fstart: int32(0x1fb6),
		Fend:   int32(0x1fbc),
	},
	292: {
		Fstart: int32(0x1fbe),
		Fend:   int32(0x1fbe),
	},
	293: {
		Fstart: int32(0x1fc2),
		Fend:   int32(0x1fc4),
	},
	294: {
		Fstart: int32(0x1fc6),
		Fend:   int32(0x1fcc),
	},
	295: {
		Fstart: int32(0x1fd0),
		Fend:   int32(0x1fd3),
	},
	296: {
		Fstart: int32(0x1fd6),
		Fend:   int32(0x1fdb),
	},
	297: {
		Fstart: int32(0x1fe0),
		Fend:   int32(0x1fec),
	},
	298: {
		Fstart: int32(0x1ff2),
		Fend:   int32(0x1ff4),
	},
	299: {
		Fstart: int32(0x1ff6),
		Fend:   int32(0x1ffc),
	},
	300: {
		Fstart: int32(0x203f),
		Fend:   int32(0x2040),
	},
	301: {
		Fstart: int32(0x2054),
		Fend:   int32(0x2054),
	},
	302: {
		Fstart: int32(0x2071),
		Fend:   int32(0x2071),
	},
	303: {
		Fstart: int32(0x207f),
		Fend:   int32(0x207f),
	},
	304: {
		Fstart: int32(0x2090),
		Fend:   int32(0x209c),
	},
	305: {
		Fstart: int32(0x20d0),
		Fend:   int32(0x20dc),
	},
	306: {
		Fstart: int32(0x20e1),
		Fend:   int32(0x20e1),
	},
	307: {
		Fstart: int32(0x20e5),
		Fend:   int32(0x20f0),
	},
	308: {
		Fstart: int32(0x2102),
		Fend:   int32(0x2102),
	},
	309: {
		Fstart: int32(0x2107),
		Fend:   int32(0x2107),
	},
	310: {
		Fstart: int32(0x210a),
		Fend:   int32(0x2113),
	},
	311: {
		Fstart: int32(0x2115),
		Fend:   int32(0x2115),
	},
	312: {
		Fstart: int32(0x2118),
		Fend:   int32(0x211d),
	},
	313: {
		Fstart: int32(0x2124),
		Fend:   int32(0x2124),
	},
	314: {
		Fstart: int32(0x2126),
		Fend:   int32(0x2126),
	},
	315: {
		Fstart: int32(0x2128),
		Fend:   int32(0x2128),
	},
	316: {
		Fstart: int32(0x212a),
		Fend:   int32(0x2139),
	},
	317: {
		Fstart: int32(0x213c),
		Fend:   int32(0x213f),
	},
	318: {
		Fstart: int32(0x2145),
		Fend:   int32(0x2149),
	},
	319: {
		Fstart: int32(0x214e),
		Fend:   int32(0x214e),
	},
	320: {
		Fstart: int32(0x2160),
		Fend:   int32(0x2188),
	},
	321: {
		Fstart: int32(0x2c00),
		Fend:   int32(0x2ce4),
	},
	322: {
		Fstart: int32(0x2ceb),
		Fend:   int32(0x2cf3),
	},
	323: {
		Fstart: int32(0x2d00),
		Fend:   int32(0x2d25),
	},
	324: {
		Fstart: int32(0x2d27),
		Fend:   int32(0x2d27),
	},
	325: {
		Fstart: int32(0x2d2d),
		Fend:   int32(0x2d2d),
	},
	326: {
		Fstart: int32(0x2d30),
		Fend:   int32(0x2d67),
	},
	327: {
		Fstart: int32(0x2d6f),
		Fend:   int32(0x2d6f),
	},
	328: {
		Fstart: int32(0x2d7f),
		Fend:   int32(0x2d96),
	},
	329: {
		Fstart: int32(0x2da0),
		Fend:   int32(0x2da6),
	},
	330: {
		Fstart: int32(0x2da8),
		Fend:   int32(0x2dae),
	},
	331: {
		Fstart: int32(0x2db0),
		Fend:   int32(0x2db6),
	},
	332: {
		Fstart: int32(0x2db8),
		Fend:   int32(0x2dbe),
	},
	333: {
		Fstart: int32(0x2dc0),
		Fend:   int32(0x2dc6),
	},
	334: {
		Fstart: int32(0x2dc8),
		Fend:   int32(0x2dce),
	},
	335: {
		Fstart: int32(0x2dd0),
		Fend:   int32(0x2dd6),
	},
	336: {
		Fstart: int32(0x2dd8),
		Fend:   int32(0x2dde),
	},
	337: {
		Fstart: int32(0x2de0),
		Fend:   int32(0x2dff),
	},
	338: {
		Fstart: int32(0x3005),
		Fend:   int32(0x3007),
	},
	339: {
		Fstart: int32(0x3021),
		Fend:   int32(0x302f),
	},
	340: {
		Fstart: int32(0x3031),
		Fend:   int32(0x3035),
	},
	341: {
		Fstart: int32(0x3038),
		Fend:   int32(0x303c),
	},
	342: {
		Fstart: int32(0x3041),
		Fend:   int32(0x3096),
	},
	343: {
		Fstart: int32(0x3099),
		Fend:   int32(0x309a),
	},
	344: {
		Fstart: int32(0x309d),
		Fend:   int32(0x309f),
	},
	345: {
		Fstart: int32(0x30a1),
		Fend:   int32(0x30fa),
	},
	346: {
		Fstart: int32(0x30fc),
		Fend:   int32(0x30ff),
	},
	347: {
		Fstart: int32(0x3105),
		Fend:   int32(0x312f),
	},
	348: {
		Fstart: int32(0x3131),
		Fend:   int32(0x318e),
	},
	349: {
		Fstart: int32(0x31a0),
		Fend:   int32(0x31bf),
	},
	350: {
		Fstart: int32(0x31f0),
		Fend:   int32(0x31ff),
	},
	351: {
		Fstart: int32(0x3400),
		Fend:   int32(0x4dbf),
	},
	352: {
		Fstart: int32(0x4e00),
		Fend:   int32(0xa48c),
	},
	353: {
		Fstart: int32(0xa4d0),
		Fend:   int32(0xa4fd),
	},
	354: {
		Fstart: int32(0xa500),
		Fend:   int32(0xa60c),
	},
	355: {
		Fstart: int32(0xa610),
		Fend:   int32(0xa62b),
	},
	356: {
		Fstart: int32(0xa640),
		Fend:   int32(0xa66f),
	},
	357: {
		Fstart: int32(0xa674),
		Fend:   int32(0xa67d),
	},
	358: {
		Fstart: int32(0xa67f),
		Fend:   int32(0xa6f1),
	},
	359: {
		Fstart: int32(0xa717),
		Fend:   int32(0xa71f),
	},
	360: {
		Fstart: int32(0xa722),
		Fend:   int32(0xa788),
	},
	361: {
		Fstart: int32(0xa78b),
		Fend:   int32(0xa7ca),
	},
	362: {
		Fstart: int32(0xa7d0),
		Fend:   int32(0xa7d1),
	},
	363: {
		Fstart: int32(0xa7d3),
		Fend:   int32(0xa7d3),
	},
	364: {
		Fstart: int32(0xa7d5),
		Fend:   int32(0xa7d9),
	},
	365: {
		Fstart: int32(0xa7f2),
		Fend:   int32(0xa827),
	},
	366: {
		Fstart: int32(0xa82c),
		Fend:   int32(0xa82c),
	},
	367: {
		Fstart: int32(0xa840),
		Fend:   int32(0xa873),
	},
	368: {
		Fstart: int32(0xa880),
		Fend:   int32(0xa8c5),
	},
	369: {
		Fstart: int32(0xa8d0),
		Fend:   int32(0xa8d9),
	},
	370: {
		Fstart: int32(0xa8e0),
		Fend:   int32(0xa8f7),
	},
	371: {
		Fstart: int32(0xa8fb),
		Fend:   int32(0xa8fb),
	},
	372: {
		Fstart: int32(0xa8fd),
		Fend:   int32(0xa92d),
	},
	373: {
		Fstart: int32(0xa930),
		Fend:   int32(0xa953),
	},
	374: {
		Fstart: int32(0xa960),
		Fend:   int32(0xa97c),
	},
	375: {
		Fstart: int32(0xa980),
		Fend:   int32(0xa9c0),
	},
	376: {
		Fstart: int32(0xa9cf),
		Fend:   int32(0xa9d9),
	},
	377: {
		Fstart: int32(0xa9e0),
		Fend:   int32(0xa9fe),
	},
	378: {
		Fstart: int32(0xaa00),
		Fend:   int32(0xaa36),
	},
	379: {
		Fstart: int32(0xaa40),
		Fend:   int32(0xaa4d),
	},
	380: {
		Fstart: int32(0xaa50),
		Fend:   int32(0xaa59),
	},
	381: {
		Fstart: int32(0xaa60),
		Fend:   int32(0xaa76),
	},
	382: {
		Fstart: int32(0xaa7a),
		Fend:   int32(0xaac2),
	},
	383: {
		Fstart: int32(0xaadb),
		Fend:   int32(0xaadd),
	},
	384: {
		Fstart: int32(0xaae0),
		Fend:   int32(0xaaef),
	},
	385: {
		Fstart: int32(0xaaf2),
		Fend:   int32(0xaaf6),
	},
	386: {
		Fstart: int32(0xab01),
		Fend:   int32(0xab06),
	},
	387: {
		Fstart: int32(0xab09),
		Fend:   int32(0xab0e),
	},
	388: {
		Fstart: int32(0xab11),
		Fend:   int32(0xab16),
	},
	389: {
		Fstart: int32(0xab20),
		Fend:   int32(0xab26),
	},
	390: {
		Fstart: int32(0xab28),
		Fend:   int32(0xab2e),
	},
	391: {
		Fstart: int32(0xab30),
		Fend:   int32(0xab5a),
	},
	392: {
		Fstart: int32(0xab5c),
		Fend:   int32(0xab69),
	},
	393: {
		Fstart: int32(0xab70),
		Fend:   int32(0xabea),
	},
	394: {
		Fstart: int32(0xabec),
		Fend:   int32(0xabed),
	},
	395: {
		Fstart: int32(0xabf0),
		Fend:   int32(0xabf9),
	},
	396: {
		Fstart: int32(0xac00),
		Fend:   int32(0xd7a3),
	},
	397: {
		Fstart: int32(0xd7b0),
		Fend:   int32(0xd7c6),
	},
	398: {
		Fstart: int32(0xd7cb),
		Fend:   int32(0xd7fb),
	},
	399: {
		Fstart: int32(0xf900),
		Fend:   int32(0xfa6d),
	},
	400: {
		Fstart: int32(0xfa70),
		Fend:   int32(0xfad9),
	},
	401: {
		Fstart: int32(0xfb00),
		Fend:   int32(0xfb06),
	},
	402: {
		Fstart: int32(0xfb13),
		Fend:   int32(0xfb17),
	},
	403: {
		Fstart: int32(0xfb1d),
		Fend:   int32(0xfb28),
	},
	404: {
		Fstart: int32(0xfb2a),
		Fend:   int32(0xfb36),
	},
	405: {
		Fstart: int32(0xfb38),
		Fend:   int32(0xfb3c),
	},
	406: {
		Fstart: int32(0xfb3e),
		Fend:   int32(0xfb3e),
	},
	407: {
		Fstart: int32(0xfb40),
		Fend:   int32(0xfb41),
	},
	408: {
		Fstart: int32(0xfb43),
		Fend:   int32(0xfb44),
	},
	409: {
		Fstart: int32(0xfb46),
		Fend:   int32(0xfbb1),
	},
	410: {
		Fstart: int32(0xfbd3),
		Fend:   int32(0xfc5d),
	},
	411: {
		Fstart: int32(0xfc64),
		Fend:   int32(0xfd3d),
	},
	412: {
		Fstart: int32(0xfd50),
		Fend:   int32(0xfd8f),
	},
	413: {
		Fstart: int32(0xfd92),
		Fend:   int32(0xfdc7),
	},
	414: {
		Fstart: int32(0xfdf0),
		Fend:   int32(0xfdf9),
	},
	415: {
		Fstart: int32(0xfe00),
		Fend:   int32(0xfe0f),
	},
	416: {
		Fstart: int32(0xfe20),
		Fend:   int32(0xfe2f),
	},
	417: {
		Fstart: int32(0xfe33),
		Fend:   int32(0xfe34),
	},
	418: {
		Fstart: int32(0xfe4d),
		Fend:   int32(0xfe4f),
	},
	419: {
		Fstart: int32(0xfe71),
		Fend:   int32(0xfe71),
	},
	420: {
		Fstart: int32(0xfe73),
		Fend:   int32(0xfe73),
	},
	421: {
		Fstart: int32(0xfe77),
		Fend:   int32(0xfe77),
	},
	422: {
		Fstart: int32(0xfe79),
		Fend:   int32(0xfe79),
	},
	423: {
		Fstart: int32(0xfe7b),
		Fend:   int32(0xfe7b),
	},
	424: {
		Fstart: int32(0xfe7d),
		Fend:   int32(0xfe7d),
	},
	425: {
		Fstart: int32(0xfe7f),
		Fend:   int32(0xfefc),
	},
	426: {
		Fstart: int32(0xff10),
		Fend:   int32(0xff19),
	},
	427: {
		Fstart: int32(0xff21),
		Fend:   int32(0xff3a),
	},
	428: {
		Fstart: int32(0xff3f),
		Fend:   int32(0xff3f),
	},
	429: {
		Fstart: int32(0xff41),
		Fend:   int32(0xff5a),
	},
	430: {
		Fstart: int32(0xff66),
		Fend:   int32(0xffbe),
	},
	431: {
		Fstart: int32(0xffc2),
		Fend:   int32(0xffc7),
	},
	432: {
		Fstart: int32(0xffca),
		Fend:   int32(0xffcf),
	},
	433: {
		Fstart: int32(0xffd2),
		Fend:   int32(0xffd7),
	},
	434: {
		Fstart: int32(0xffda),
		Fend:   int32(0xffdc),
	},
	435: {
		Fstart: int32(0x10000),
		Fend:   int32(0x1000b),
	},
	436: {
		Fstart: int32(0x1000d),
		Fend:   int32(0x10026),
	},
	437: {
		Fstart: int32(0x10028),
		Fend:   int32(0x1003a),
	},
	438: {
		Fstart: int32(0x1003c),
		Fend:   int32(0x1003d),
	},
	439: {
		Fstart: int32(0x1003f),
		Fend:   int32(0x1004d),
	},
	440: {
		Fstart: int32(0x10050),
		Fend:   int32(0x1005d),
	},
	441: {
		Fstart: int32(0x10080),
		Fend:   int32(0x100fa),
	},
	442: {
		Fstart: int32(0x10140),
		Fend:   int32(0x10174),
	},
	443: {
		Fstart: int32(0x101fd),
		Fend:   int32(0x101fd),
	},
	444: {
		Fstart: int32(0x10280),
		Fend:   int32(0x1029c),
	},
	445: {
		Fstart: int32(0x102a0),
		Fend:   int32(0x102d0),
	},
	446: {
		Fstart: int32(0x102e0),
		Fend:   int32(0x102e0),
	},
	447: {
		Fstart: int32(0x10300),
		Fend:   int32(0x1031f),
	},
	448: {
		Fstart: int32(0x1032d),
		Fend:   int32(0x1034a),
	},
	449: {
		Fstart: int32(0x10350),
		Fend:   int32(0x1037a),
	},
	450: {
		Fstart: int32(0x10380),
		Fend:   int32(0x1039d),
	},
	451: {
		Fstart: int32(0x103a0),
		Fend:   int32(0x103c3),
	},
	452: {
		Fstart: int32(0x103c8),
		Fend:   int32(0x103cf),
	},
	453: {
		Fstart: int32(0x103d1),
		Fend:   int32(0x103d5),
	},
	454: {
		Fstart: int32(0x10400),
		Fend:   int32(0x1049d),
	},
	455: {
		Fstart: int32(0x104a0),
		Fend:   int32(0x104a9),
	},
	456: {
		Fstart: int32(0x104b0),
		Fend:   int32(0x104d3),
	},
	457: {
		Fstart: int32(0x104d8),
		Fend:   int32(0x104fb),
	},
	458: {
		Fstart: int32(0x10500),
		Fend:   int32(0x10527),
	},
	459: {
		Fstart: int32(0x10530),
		Fend:   int32(0x10563),
	},
	460: {
		Fstart: int32(0x10570),
		Fend:   int32(0x1057a),
	},
	461: {
		Fstart: int32(0x1057c),
		Fend:   int32(0x1058a),
	},
	462: {
		Fstart: int32(0x1058c),
		Fend:   int32(0x10592),
	},
	463: {
		Fstart: int32(0x10594),
		Fend:   int32(0x10595),
	},
	464: {
		Fstart: int32(0x10597),
		Fend:   int32(0x105a1),
	},
	465: {
		Fstart: int32(0x105a3),
		Fend:   int32(0x105b1),
	},
	466: {
		Fstart: int32(0x105b3),
		Fend:   int32(0x105b9),
	},
	467: {
		Fstart: int32(0x105bb),
		Fend:   int32(0x105bc),
	},
	468: {
		Fstart: int32(0x10600),
		Fend:   int32(0x10736),
	},
	469: {
		Fstart: int32(0x10740),
		Fend:   int32(0x10755),
	},
	470: {
		Fstart: int32(0x10760),
		Fend:   int32(0x10767),
	},
	471: {
		Fstart: int32(0x10780),
		Fend:   int32(0x10785),
	},
	472: {
		Fstart: int32(0x10787),
		Fend:   int32(0x107b0),
	},
	473: {
		Fstart: int32(0x107b2),
		Fend:   int32(0x107ba),
	},
	474: {
		Fstart: int32(0x10800),
		Fend:   int32(0x10805),
	},
	475: {
		Fstart: int32(0x10808),
		Fend:   int32(0x10808),
	},
	476: {
		Fstart: int32(0x1080a),
		Fend:   int32(0x10835),
	},
	477: {
		Fstart: int32(0x10837),
		Fend:   int32(0x10838),
	},
	478: {
		Fstart: int32(0x1083c),
		Fend:   int32(0x1083c),
	},
	479: {
		Fstart: int32(0x1083f),
		Fend:   int32(0x10855),
	},
	480: {
		Fstart: int32(0x10860),
		Fend:   int32(0x10876),
	},
	481: {
		Fstart: int32(0x10880),
		Fend:   int32(0x1089e),
	},
	482: {
		Fstart: int32(0x108e0),
		Fend:   int32(0x108f2),
	},
	483: {
		Fstart: int32(0x108f4),
		Fend:   int32(0x108f5),
	},
	484: {
		Fstart: int32(0x10900),
		Fend:   int32(0x10915),
	},
	485: {
		Fstart: int32(0x10920),
		Fend:   int32(0x10939),
	},
	486: {
		Fstart: int32(0x10980),
		Fend:   int32(0x109b7),
	},
	487: {
		Fstart: int32(0x109be),
		Fend:   int32(0x109bf),
	},
	488: {
		Fstart: int32(0x10a00),
		Fend:   int32(0x10a03),
	},
	489: {
		Fstart: int32(0x10a05),
		Fend:   int32(0x10a06),
	},
	490: {
		Fstart: int32(0x10a0c),
		Fend:   int32(0x10a13),
	},
	491: {
		Fstart: int32(0x10a15),
		Fend:   int32(0x10a17),
	},
	492: {
		Fstart: int32(0x10a19),
		Fend:   int32(0x10a35),
	},
	493: {
		Fstart: int32(0x10a38),
		Fend:   int32(0x10a3a),
	},
	494: {
		Fstart: int32(0x10a3f),
		Fend:   int32(0x10a3f),
	},
	495: {
		Fstart: int32(0x10a60),
		Fend:   int32(0x10a7c),
	},
	496: {
		Fstart: int32(0x10a80),
		Fend:   int32(0x10a9c),
	},
	497: {
		Fstart: int32(0x10ac0),
		Fend:   int32(0x10ac7),
	},
	498: {
		Fstart: int32(0x10ac9),
		Fend:   int32(0x10ae6),
	},
	499: {
		Fstart: int32(0x10b00),
		Fend:   int32(0x10b35),
	},
	500: {
		Fstart: int32(0x10b40),
		Fend:   int32(0x10b55),
	},
	501: {
		Fstart: int32(0x10b60),
		Fend:   int32(0x10b72),
	},
	502: {
		Fstart: int32(0x10b80),
		Fend:   int32(0x10b91),
	},
	503: {
		Fstart: int32(0x10c00),
		Fend:   int32(0x10c48),
	},
	504: {
		Fstart: int32(0x10c80),
		Fend:   int32(0x10cb2),
	},
	505: {
		Fstart: int32(0x10cc0),
		Fend:   int32(0x10cf2),
	},
	506: {
		Fstart: int32(0x10d00),
		Fend:   int32(0x10d27),
	},
	507: {
		Fstart: int32(0x10d30),
		Fend:   int32(0x10d39),
	},
	508: {
		Fstart: int32(0x10e80),
		Fend:   int32(0x10ea9),
	},
	509: {
		Fstart: int32(0x10eab),
		Fend:   int32(0x10eac),
	},
	510: {
		Fstart: int32(0x10eb0),
		Fend:   int32(0x10eb1),
	},
	511: {
		Fstart: int32(0x10f00),
		Fend:   int32(0x10f1c),
	},
	512: {
		Fstart: int32(0x10f27),
		Fend:   int32(0x10f27),
	},
	513: {
		Fstart: int32(0x10f30),
		Fend:   int32(0x10f50),
	},
	514: {
		Fstart: int32(0x10f70),
		Fend:   int32(0x10f85),
	},
	515: {
		Fstart: int32(0x10fb0),
		Fend:   int32(0x10fc4),
	},
	516: {
		Fstart: int32(0x10fe0),
		Fend:   int32(0x10ff6),
	},
	517: {
		Fstart: int32(0x11000),
		Fend:   int32(0x11046),
	},
	518: {
		Fstart: int32(0x11066),
		Fend:   int32(0x11075),
	},
	519: {
		Fstart: int32(0x1107f),
		Fend:   int32(0x110ba),
	},
	520: {
		Fstart: int32(0x110c2),
		Fend:   int32(0x110c2),
	},
	521: {
		Fstart: int32(0x110d0),
		Fend:   int32(0x110e8),
	},
	522: {
		Fstart: int32(0x110f0),
		Fend:   int32(0x110f9),
	},
	523: {
		Fstart: int32(0x11100),
		Fend:   int32(0x11134),
	},
	524: {
		Fstart: int32(0x11136),
		Fend:   int32(0x1113f),
	},
	525: {
		Fstart: int32(0x11144),
		Fend:   int32(0x11147),
	},
	526: {
		Fstart: int32(0x11150),
		Fend:   int32(0x11173),
	},
	527: {
		Fstart: int32(0x11176),
		Fend:   int32(0x11176),
	},
	528: {
		Fstart: int32(0x11180),
		Fend:   int32(0x111c4),
	},
	529: {
		Fstart: int32(0x111c9),
		Fend:   int32(0x111cc),
	},
	530: {
		Fstart: int32(0x111ce),
		Fend:   int32(0x111da),
	},
	531: {
		Fstart: int32(0x111dc),
		Fend:   int32(0x111dc),
	},
	532: {
		Fstart: int32(0x11200),
		Fend:   int32(0x11211),
	},
	533: {
		Fstart: int32(0x11213),
		Fend:   int32(0x11237),
	},
	534: {
		Fstart: int32(0x1123e),
		Fend:   int32(0x1123e),
	},
	535: {
		Fstart: int32(0x11280),
		Fend:   int32(0x11286),
	},
	536: {
		Fstart: int32(0x11288),
		Fend:   int32(0x11288),
	},
	537: {
		Fstart: int32(0x1128a),
		Fend:   int32(0x1128d),
	},
	538: {
		Fstart: int32(0x1128f),
		Fend:   int32(0x1129d),
	},
	539: {
		Fstart: int32(0x1129f),
		Fend:   int32(0x112a8),
	},
	540: {
		Fstart: int32(0x112b0),
		Fend:   int32(0x112ea),
	},
	541: {
		Fstart: int32(0x112f0),
		Fend:   int32(0x112f9),
	},
	542: {
		Fstart: int32(0x11300),
		Fend:   int32(0x11303),
	},
	543: {
		Fstart: int32(0x11305),
		Fend:   int32(0x1130c),
	},
	544: {
		Fstart: int32(0x1130f),
		Fend:   int32(0x11310),
	},
	545: {
		Fstart: int32(0x11313),
		Fend:   int32(0x11328),
	},
	546: {
		Fstart: int32(0x1132a),
		Fend:   int32(0x11330),
	},
	547: {
		Fstart: int32(0x11332),
		Fend:   int32(0x11333),
	},
	548: {
		Fstart: int32(0x11335),
		Fend:   int32(0x11339),
	},
	549: {
		Fstart: int32(0x1133b),
		Fend:   int32(0x11344),
	},
	550: {
		Fstart: int32(0x11347),
		Fend:   int32(0x11348),
	},
	551: {
		Fstart: int32(0x1134b),
		Fend:   int32(0x1134d),
	},
	552: {
		Fstart: int32(0x11350),
		Fend:   int32(0x11350),
	},
	553: {
		Fstart: int32(0x11357),
		Fend:   int32(0x11357),
	},
	554: {
		Fstart: int32(0x1135d),
		Fend:   int32(0x11363),
	},
	555: {
		Fstart: int32(0x11366),
		Fend:   int32(0x1136c),
	},
	556: {
		Fstart: int32(0x11370),
		Fend:   int32(0x11374),
	},
	557: {
		Fstart: int32(0x11400),
		Fend:   int32(0x1144a),
	},
	558: {
		Fstart: int32(0x11450),
		Fend:   int32(0x11459),
	},
	559: {
		Fstart: int32(0x1145e),
		Fend:   int32(0x11461),
	},
	560: {
		Fstart: int32(0x11480),
		Fend:   int32(0x114c5),
	},
	561: {
		Fstart: int32(0x114c7),
		Fend:   int32(0x114c7),
	},
	562: {
		Fstart: int32(0x114d0),
		Fend:   int32(0x114d9),
	},
	563: {
		Fstart: int32(0x11580),
		Fend:   int32(0x115b5),
	},
	564: {
		Fstart: int32(0x115b8),
		Fend:   int32(0x115c0),
	},
	565: {
		Fstart: int32(0x115d8),
		Fend:   int32(0x115dd),
	},
	566: {
		Fstart: int32(0x11600),
		Fend:   int32(0x11640),
	},
	567: {
		Fstart: int32(0x11644),
		Fend:   int32(0x11644),
	},
	568: {
		Fstart: int32(0x11650),
		Fend:   int32(0x11659),
	},
	569: {
		Fstart: int32(0x11680),
		Fend:   int32(0x116b8),
	},
	570: {
		Fstart: int32(0x116c0),
		Fend:   int32(0x116c9),
	},
	571: {
		Fstart: int32(0x11700),
		Fend:   int32(0x1171a),
	},
	572: {
		Fstart: int32(0x1171d),
		Fend:   int32(0x1172b),
	},
	573: {
		Fstart: int32(0x11730),
		Fend:   int32(0x11739),
	},
	574: {
		Fstart: int32(0x11740),
		Fend:   int32(0x11746),
	},
	575: {
		Fstart: int32(0x11800),
		Fend:   int32(0x1183a),
	},
	576: {
		Fstart: int32(0x118a0),
		Fend:   int32(0x118e9),
	},
	577: {
		Fstart: int32(0x118ff),
		Fend:   int32(0x11906),
	},
	578: {
		Fstart: int32(0x11909),
		Fend:   int32(0x11909),
	},
	579: {
		Fstart: int32(0x1190c),
		Fend:   int32(0x11913),
	},
	580: {
		Fstart: int32(0x11915),
		Fend:   int32(0x11916),
	},
	581: {
		Fstart: int32(0x11918),
		Fend:   int32(0x11935),
	},
	582: {
		Fstart: int32(0x11937),
		Fend:   int32(0x11938),
	},
	583: {
		Fstart: int32(0x1193b),
		Fend:   int32(0x11943),
	},
	584: {
		Fstart: int32(0x11950),
		Fend:   int32(0x11959),
	},
	585: {
		Fstart: int32(0x119a0),
		Fend:   int32(0x119a7),
	},
	586: {
		Fstart: int32(0x119aa),
		Fend:   int32(0x119d7),
	},
	587: {
		Fstart: int32(0x119da),
		Fend:   int32(0x119e1),
	},
	588: {
		Fstart: int32(0x119e3),
		Fend:   int32(0x119e4),
	},
	589: {
		Fstart: int32(0x11a00),
		Fend:   int32(0x11a3e),
	},
	590: {
		Fstart: int32(0x11a47),
		Fend:   int32(0x11a47),
	},
	591: {
		Fstart: int32(0x11a50),
		Fend:   int32(0x11a99),
	},
	592: {
		Fstart: int32(0x11a9d),
		Fend:   int32(0x11a9d),
	},
	593: {
		Fstart: int32(0x11ab0),
		Fend:   int32(0x11af8),
	},
	594: {
		Fstart: int32(0x11c00),
		Fend:   int32(0x11c08),
	},
	595: {
		Fstart: int32(0x11c0a),
		Fend:   int32(0x11c36),
	},
	596: {
		Fstart: int32(0x11c38),
		Fend:   int32(0x11c40),
	},
	597: {
		Fstart: int32(0x11c50),
		Fend:   int32(0x11c59),
	},
	598: {
		Fstart: int32(0x11c72),
		Fend:   int32(0x11c8f),
	},
	599: {
		Fstart: int32(0x11c92),
		Fend:   int32(0x11ca7),
	},
	600: {
		Fstart: int32(0x11ca9),
		Fend:   int32(0x11cb6),
	},
	601: {
		Fstart: int32(0x11d00),
		Fend:   int32(0x11d06),
	},
	602: {
		Fstart: int32(0x11d08),
		Fend:   int32(0x11d09),
	},
	603: {
		Fstart: int32(0x11d0b),
		Fend:   int32(0x11d36),
	},
	604: {
		Fstart: int32(0x11d3a),
		Fend:   int32(0x11d3a),
	},
	605: {
		Fstart: int32(0x11d3c),
		Fend:   int32(0x11d3d),
	},
	606: {
		Fstart: int32(0x11d3f),
		Fend:   int32(0x11d47),
	},
	607: {
		Fstart: int32(0x11d50),
		Fend:   int32(0x11d59),
	},
	608: {
		Fstart: int32(0x11d60),
		Fend:   int32(0x11d65),
	},
	609: {
		Fstart: int32(0x11d67),
		Fend:   int32(0x11d68),
	},
	610: {
		Fstart: int32(0x11d6a),
		Fend:   int32(0x11d8e),
	},
	611: {
		Fstart: int32(0x11d90),
		Fend:   int32(0x11d91),
	},
	612: {
		Fstart: int32(0x11d93),
		Fend:   int32(0x11d98),
	},
	613: {
		Fstart: int32(0x11da0),
		Fend:   int32(0x11da9),
	},
	614: {
		Fstart: int32(0x11ee0),
		Fend:   int32(0x11ef6),
	},
	615: {
		Fstart: int32(0x11fb0),
		Fend:   int32(0x11fb0),
	},
	616: {
		Fstart: int32(0x12000),
		Fend:   int32(0x12399),
	},
	617: {
		Fstart: int32(0x12400),
		Fend:   int32(0x1246e),
	},
	618: {
		Fstart: int32(0x12480),
		Fend:   int32(0x12543),
	},
	619: {
		Fstart: int32(0x12f90),
		Fend:   int32(0x12ff0),
	},
	620: {
		Fstart: int32(0x13000),
		Fend:   int32(0x1342e),
	},
	621: {
		Fstart: int32(0x14400),
		Fend:   int32(0x14646),
	},
	622: {
		Fstart: int32(0x16800),
		Fend:   int32(0x16a38),
	},
	623: {
		Fstart: int32(0x16a40),
		Fend:   int32(0x16a5e),
	},
	624: {
		Fstart: int32(0x16a60),
		Fend:   int32(0x16a69),
	},
	625: {
		Fstart: int32(0x16a70),
		Fend:   int32(0x16abe),
	},
	626: {
		Fstart: int32(0x16ac0),
		Fend:   int32(0x16ac9),
	},
	627: {
		Fstart: int32(0x16ad0),
		Fend:   int32(0x16aed),
	},
	628: {
		Fstart: int32(0x16af0),
		Fend:   int32(0x16af4),
	},
	629: {
		Fstart: int32(0x16b00),
		Fend:   int32(0x16b36),
	},
	630: {
		Fstart: int32(0x16b40),
		Fend:   int32(0x16b43),
	},
	631: {
		Fstart: int32(0x16b50),
		Fend:   int32(0x16b59),
	},
	632: {
		Fstart: int32(0x16b63),
		Fend:   int32(0x16b77),
	},
	633: {
		Fstart: int32(0x16b7d),
		Fend:   int32(0x16b8f),
	},
	634: {
		Fstart: int32(0x16e40),
		Fend:   int32(0x16e7f),
	},
	635: {
		Fstart: int32(0x16f00),
		Fend:   int32(0x16f4a),
	},
	636: {
		Fstart: int32(0x16f4f),
		Fend:   int32(0x16f87),
	},
	637: {
		Fstart: int32(0x16f8f),
		Fend:   int32(0x16f9f),
	},
	638: {
		Fstart: int32(0x16fe0),
		Fend:   int32(0x16fe1),
	},
	639: {
		Fstart: int32(0x16fe3),
		Fend:   int32(0x16fe4),
	},
	640: {
		Fstart: int32(0x16ff0),
		Fend:   int32(0x16ff1),
	},
	641: {
		Fstart: int32(0x17000),
		Fend:   int32(0x187f7),
	},
	642: {
		Fstart: int32(0x18800),
		Fend:   int32(0x18cd5),
	},
	643: {
		Fstart: int32(0x18d00),
		Fend:   int32(0x18d08),
	},
	644: {
		Fstart: int32(0x1aff0),
		Fend:   int32(0x1aff3),
	},
	645: {
		Fstart: int32(0x1aff5),
		Fend:   int32(0x1affb),
	},
	646: {
		Fstart: int32(0x1affd),
		Fend:   int32(0x1affe),
	},
	647: {
		Fstart: int32(0x1b000),
		Fend:   int32(0x1b122),
	},
	648: {
		Fstart: int32(0x1b150),
		Fend:   int32(0x1b152),
	},
	649: {
		Fstart: int32(0x1b164),
		Fend:   int32(0x1b167),
	},
	650: {
		Fstart: int32(0x1b170),
		Fend:   int32(0x1b2fb),
	},
	651: {
		Fstart: int32(0x1bc00),
		Fend:   int32(0x1bc6a),
	},
	652: {
		Fstart: int32(0x1bc70),
		Fend:   int32(0x1bc7c),
	},
	653: {
		Fstart: int32(0x1bc80),
		Fend:   int32(0x1bc88),
	},
	654: {
		Fstart: int32(0x1bc90),
		Fend:   int32(0x1bc99),
	},
	655: {
		Fstart: int32(0x1bc9d),
		Fend:   int32(0x1bc9e),
	},
	656: {
		Fstart: int32(0x1cf00),
		Fend:   int32(0x1cf2d),
	},
	657: {
		Fstart: int32(0x1cf30),
		Fend:   int32(0x1cf46),
	},
	658: {
		Fstart: int32(0x1d165),
		Fend:   int32(0x1d169),
	},
	659: {
		Fstart: int32(0x1d16d),
		Fend:   int32(0x1d172),
	},
	660: {
		Fstart: int32(0x1d17b),
		Fend:   int32(0x1d182),
	},
	661: {
		Fstart: int32(0x1d185),
		Fend:   int32(0x1d18b),
	},
	662: {
		Fstart: int32(0x1d1aa),
		Fend:   int32(0x1d1ad),
	},
	663: {
		Fstart: int32(0x1d242),
		Fend:   int32(0x1d244),
	},
	664: {
		Fstart: int32(0x1d400),
		Fend:   int32(0x1d454),
	},
	665: {
		Fstart: int32(0x1d456),
		Fend:   int32(0x1d49c),
	},
	666: {
		Fstart: int32(0x1d49e),
		Fend:   int32(0x1d49f),
	},
	667: {
		Fstart: int32(0x1d4a2),
		Fend:   int32(0x1d4a2),
	},
	668: {
		Fstart: int32(0x1d4a5),
		Fend:   int32(0x1d4a6),
	},
	669: {
		Fstart: int32(0x1d4a9),
		Fend:   int32(0x1d4ac),
	},
	670: {
		Fstart: int32(0x1d4ae),
		Fend:   int32(0x1d4b9),
	},
	671: {
		Fstart: int32(0x1d4bb),
		Fend:   int32(0x1d4bb),
	},
	672: {
		Fstart: int32(0x1d4bd),
		Fend:   int32(0x1d4c3),
	},
	673: {
		Fstart: int32(0x1d4c5),
		Fend:   int32(0x1d505),
	},
	674: {
		Fstart: int32(0x1d507),
		Fend:   int32(0x1d50a),
	},
	675: {
		Fstart: int32(0x1d50d),
		Fend:   int32(0x1d514),
	},
	676: {
		Fstart: int32(0x1d516),
		Fend:   int32(0x1d51c),
	},
	677: {
		Fstart: int32(0x1d51e),
		Fend:   int32(0x1d539),
	},
	678: {
		Fstart: int32(0x1d53b),
		Fend:   int32(0x1d53e),
	},
	679: {
		Fstart: int32(0x1d540),
		Fend:   int32(0x1d544),
	},
	680: {
		Fstart: int32(0x1d546),
		Fend:   int32(0x1d546),
	},
	681: {
		Fstart: int32(0x1d54a),
		Fend:   int32(0x1d550),
	},
	682: {
		Fstart: int32(0x1d552),
		Fend:   int32(0x1d6a5),
	},
	683: {
		Fstart: int32(0x1d6a8),
		Fend:   int32(0x1d6c0),
	},
	684: {
		Fstart: int32(0x1d6c2),
		Fend:   int32(0x1d6da),
	},
	685: {
		Fstart: int32(0x1d6dc),
		Fend:   int32(0x1d6fa),
	},
	686: {
		Fstart: int32(0x1d6fc),
		Fend:   int32(0x1d714),
	},
	687: {
		Fstart: int32(0x1d716),
		Fend:   int32(0x1d734),
	},
	688: {
		Fstart: int32(0x1d736),
		Fend:   int32(0x1d74e),
	},
	689: {
		Fstart: int32(0x1d750),
		Fend:   int32(0x1d76e),
	},
	690: {
		Fstart: int32(0x1d770),
		Fend:   int32(0x1d788),
	},
	691: {
		Fstart: int32(0x1d78a),
		Fend:   int32(0x1d7a8),
	},
	692: {
		Fstart: int32(0x1d7aa),
		Fend:   int32(0x1d7c2),
	},
	693: {
		Fstart: int32(0x1d7c4),
		Fend:   int32(0x1d7cb),
	},
	694: {
		Fstart: int32(0x1d7ce),
		Fend:   int32(0x1d7ff),
	},
	695: {
		Fstart: int32(0x1da00),
		Fend:   int32(0x1da36),
	},
	696: {
		Fstart: int32(0x1da3b),
		Fend:   int32(0x1da6c),
	},
	697: {
		Fstart: int32(0x1da75),
		Fend:   int32(0x1da75),
	},
	698: {
		Fstart: int32(0x1da84),
		Fend:   int32(0x1da84),
	},
	699: {
		Fstart: int32(0x1da9b),
		Fend:   int32(0x1da9f),
	},
	700: {
		Fstart: int32(0x1daa1),
		Fend:   int32(0x1daaf),
	},
	701: {
		Fstart: int32(0x1df00),
		Fend:   int32(0x1df1e),
	},
	702: {
		Fstart: int32(0x1e000),
		Fend:   int32(0x1e006),
	},
	703: {
		Fstart: int32(0x1e008),
		Fend:   int32(0x1e018),
	},
	704: {
		Fstart: int32(0x1e01b),
		Fend:   int32(0x1e021),
	},
	705: {
		Fstart: int32(0x1e023),
		Fend:   int32(0x1e024),
	},
	706: {
		Fstart: int32(0x1e026),
		Fend:   int32(0x1e02a),
	},
	707: {
		Fstart: int32(0x1e100),
		Fend:   int32(0x1e12c),
	},
	708: {
		Fstart: int32(0x1e130),
		Fend:   int32(0x1e13d),
	},
	709: {
		Fstart: int32(0x1e140),
		Fend:   int32(0x1e149),
	},
	710: {
		Fstart: int32(0x1e14e),
		Fend:   int32(0x1e14e),
	},
	711: {
		Fstart: int32(0x1e290),
		Fend:   int32(0x1e2ae),
	},
	712: {
		Fstart: int32(0x1e2c0),
		Fend:   int32(0x1e2f9),
	},
	713: {
		Fstart: int32(0x1e7e0),
		Fend:   int32(0x1e7e6),
	},
	714: {
		Fstart: int32(0x1e7e8),
		Fend:   int32(0x1e7eb),
	},
	715: {
		Fstart: int32(0x1e7ed),
		Fend:   int32(0x1e7ee),
	},
	716: {
		Fstart: int32(0x1e7f0),
		Fend:   int32(0x1e7fe),
	},
	717: {
		Fstart: int32(0x1e800),
		Fend:   int32(0x1e8c4),
	},
	718: {
		Fstart: int32(0x1e8d0),
		Fend:   int32(0x1e8d6),
	},
	719: {
		Fstart: int32(0x1e900),
		Fend:   int32(0x1e94b),
	},
	720: {
		Fstart: int32(0x1e950),
		Fend:   int32(0x1e959),
	},
	721: {
		Fstart: int32(0x1ee00),
		Fend:   int32(0x1ee03),
	},
	722: {
		Fstart: int32(0x1ee05),
		Fend:   int32(0x1ee1f),
	},
	723: {
		Fstart: int32(0x1ee21),
		Fend:   int32(0x1ee22),
	},
	724: {
		Fstart: int32(0x1ee24),
		Fend:   int32(0x1ee24),
	},
	725: {
		Fstart: int32(0x1ee27),
		Fend:   int32(0x1ee27),
	},
	726: {
		Fstart: int32(0x1ee29),
		Fend:   int32(0x1ee32),
	},
	727: {
		Fstart: int32(0x1ee34),
		Fend:   int32(0x1ee37),
	},
	728: {
		Fstart: int32(0x1ee39),
		Fend:   int32(0x1ee39),
	},
	729: {
		Fstart: int32(0x1ee3b),
		Fend:   int32(0x1ee3b),
	},
	730: {
		Fstart: int32(0x1ee42),
		Fend:   int32(0x1ee42),
	},
	731: {
		Fstart: int32(0x1ee47),
		Fend:   int32(0x1ee47),
	},
	732: {
		Fstart: int32(0x1ee49),
		Fend:   int32(0x1ee49),
	},
	733: {
		Fstart: int32(0x1ee4b),
		Fend:   int32(0x1ee4b),
	},
	734: {
		Fstart: int32(0x1ee4d),
		Fend:   int32(0x1ee4f),
	},
	735: {
		Fstart: int32(0x1ee51),
		Fend:   int32(0x1ee52),
	},
	736: {
		Fstart: int32(0x1ee54),
		Fend:   int32(0x1ee54),
	},
	737: {
		Fstart: int32(0x1ee57),
		Fend:   int32(0x1ee57),
	},
	738: {
		Fstart: int32(0x1ee59),
		Fend:   int32(0x1ee59),
	},
	739: {
		Fstart: int32(0x1ee5b),
		Fend:   int32(0x1ee5b),
	},
	740: {
		Fstart: int32(0x1ee5d),
		Fend:   int32(0x1ee5d),
	},
	741: {
		Fstart: int32(0x1ee5f),
		Fend:   int32(0x1ee5f),
	},
	742: {
		Fstart: int32(0x1ee61),
		Fend:   int32(0x1ee62),
	},
	743: {
		Fstart: int32(0x1ee64),
		Fend:   int32(0x1ee64),
	},
	744: {
		Fstart: int32(0x1ee67),
		Fend:   int32(0x1ee6a),
	},
	745: {
		Fstart: int32(0x1ee6c),
		Fend:   int32(0x1ee72),
	},
	746: {
		Fstart: int32(0x1ee74),
		Fend:   int32(0x1ee77),
	},
	747: {
		Fstart: int32(0x1ee79),
		Fend:   int32(0x1ee7c),
	},
	748: {
		Fstart: int32(0x1ee7e),
		Fend:   int32(0x1ee7e),
	},
	749: {
		Fstart: int32(0x1ee80),
		Fend:   int32(0x1ee89),
	},
	750: {
		Fstart: int32(0x1ee8b),
		Fend:   int32(0x1ee9b),
	},
	751: {
		Fstart: int32(0x1eea1),
		Fend:   int32(0x1eea3),
	},
	752: {
		Fstart: int32(0x1eea5),
		Fend:   int32(0x1eea9),
	},
	753: {
		Fstart: int32(0x1eeab),
		Fend:   int32(0x1eebb),
	},
	754: {
		Fstart: int32(0x1fbf0),
		Fend:   int32(0x1fbf9),
	},
	755: {
		Fstart: int32(0x20000),
		Fend:   int32(0x2a6df),
	},
	756: {
		Fstart: int32(0x2a700),
		Fend:   int32(0x2b738),
	},
	757: {
		Fstart: int32(0x2b740),
		Fend:   int32(0x2b81d),
	},
	758: {
		Fstart: int32(0x2b820),
		Fend:   int32(0x2cea1),
	},
	759: {
		Fstart: int32(0x2ceb0),
		Fend:   int32(0x2ebe0),
	},
	760: {
		Fstart: int32(0x2f800),
		Fend:   int32(0x2fa1d),
	},
	761: {
		Fstart: int32(0x30000),
		Fend:   int32(0x3134a),
	},
	762: {
		Fstart: int32(0xe0100),
		Fend:   int32(0xe01ef),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v4 uint8
	var half_size, i, i1, i2, i3, i4, index, mid_index, size uint32_t
	var lookahead1, v3 int32_t
	var range_token, range_token1, v2 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i2, i3, i4, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v2, v3, v4
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
			state = uint16(18)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(72)/libc.Uint64FromInt64(2)) {
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
			state = uint16(16)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(32)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(656) - index
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
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(1):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(2)) {
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
			state = uint16(1)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(32)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(656) - index
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
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead1 == int32('\'') {
			state = uint16(38)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(40)
			goto next_state
		}
		if lookahead1 == int32('\\') {
			state = uint16(4)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(41)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead1 == int32('/') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead1 == int32('u') {
			state = uint16(5)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(12)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(43)
			goto next_state
		}
		if lookahead1 == int32('"') || lookahead1 == int32('\'') || lookahead1 == int32('?') || lookahead1 == int32('\\') || lookahead1 == int32('a') || lookahead1 == int32('b') || lookahead1 == int32('f') || lookahead1 == int32('n') || lookahead1 == int32('r') || int32('t') <= lookahead1 && lookahead1 <= int32('v') {
			state = uint16(44)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead1 == int32('{') {
			state = uint16(11)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(9)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead1 == int32('}') {
			state = uint16(44)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead1 == int32('0') || lookahead1 == int32('1') || lookahead1 == int32('_') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(8):
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') || lookahead1 == int32('_') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(9):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(10):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(11):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(12):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(13):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(14):
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(656) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _14
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _14
	_14:
		if v4 != 0 {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(15):
		if eof != 0 {
			state = uint16(18)
			goto next_state
		}
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _15
		_15:
			;
			i2 = i2 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(656) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _19
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _19
	_19:
		if v4 != 0 {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(16):
		if eof != 0 {
			state = uint16(18)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token3[i3]) == lookahead1 {
				state = map_token3[i3+uint32(1)]
				goto next_state
			}
			goto _20
		_20:
			;
			i3 = i3 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(16)
			goto next_state
		}
		if int32('1') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(32)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(656) - index
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
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(17):
		if eof != 0 {
			state = uint16(18)
			goto next_state
		}
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token4[i4]) == lookahead1 {
				state = map_token4[i4+uint32(1)]
				goto next_state
			}
			goto _25
		_25:
			;
			i4 = i4 + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_1))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(656) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _29
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _29
	_29:
		if v4 != 0 {
			state = uint16(56)
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
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(20):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
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
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(')') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('b') {
			state = uint16(7)
			goto next_state
		}
		if lookahead1 == int32('o') {
			state = uint16(8)
			goto next_state
		}
		if lookahead1 == int32('x') {
			state = uint16(13)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || lookahead1 == int32('_') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('0') || lookahead1 == int32('1') || lookahead1 == int32('_') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') || lookahead1 == int32('_') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || lookahead1 == int32('_') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_b)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('"') {
			state = uint16(35)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_3))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(763) - index
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
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_char_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_char_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('/') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_char_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('/') {
			state = uint16(40)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(41)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\'') && lookahead1 != int32('\\') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('7') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_3))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(763) - index
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
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_3))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(763) - index
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
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(14)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_3))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(763) - index
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
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(52)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_3))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(763) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _49
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _49
	_49:
		if v4 != 0 {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(46)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_3))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(763) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _53
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _53
	_53:
		if v4 != 0 {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(47)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_3))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(763) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _57
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _57
	_57:
		if v4 != 0 {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(54)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_3))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(763) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _61
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _61
	_61:
		if v4 != 0 {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('r') {
			state = uint16(55)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_3))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(763) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _65
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _65
	_65:
		if v4 != 0 {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('s') {
			state = uint16(51)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_3))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(763) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _69
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _69
	_69:
		if v4 != 0 {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('u') {
			state = uint16(50)
			goto next_state
		}
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_3))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(763) - index
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
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v2 = uintptr(unsafe.Pointer(&sym_identifier_character_set_3))
		v3 = lookahead1
		index = uint32(0)
		size = uint32(763) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v2 + uintptr(mid_index)*8
			if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v4 = libc.BoolUint8(true1 != 0)
				goto _77
			} else {
				if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v2 + uintptr(index)*8
		v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _77
	_77:
		if v4 != 0 {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_line_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\n') {
			state = uint16(57)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [36]uint16_t{
	0:  uint16('"'),
	1:  uint16(36),
	2:  uint16('\''),
	3:  uint16(38),
	4:  uint16('('),
	5:  uint16(26),
	6:  uint16(')'),
	7:  uint16(27),
	8:  uint16(','),
	9:  uint16(20),
	10: uint16('-'),
	11: uint16(34),
	12: uint16('/'),
	13: uint16(3),
	14: uint16('0'),
	15: uint16(29),
	16: uint16(':'),
	17: uint16(28),
	18: uint16('['),
	19: uint16(19),
	20: uint16('\\'),
	21: uint16(4),
	22: uint16(']'),
	23: uint16(21),
	24: uint16('b'),
	25: uint16(37),
	26: uint16('f'),
	27: uint16(49),
	28: uint16('r'),
	29: uint16(48),
	30: uint16('t'),
	31: uint16(53),
	32: uint16('{'),
	33: uint16(22),
	34: uint16('}'),
	35: uint16(23),
}

var map_token1 = [32]uint16_t{
	0:  uint16('"'),
	1:  uint16(35),
	2:  uint16('\''),
	3:  uint16(38),
	4:  uint16('('),
	5:  uint16(26),
	6:  uint16(')'),
	7:  uint16(27),
	8:  uint16(','),
	9:  uint16(20),
	10: uint16('-'),
	11: uint16(34),
	12: uint16('/'),
	13: uint16(3),
	14: uint16('0'),
	15: uint16(29),
	16: uint16('['),
	17: uint16(19),
	18: uint16(']'),
	19: uint16(21),
	20: uint16('b'),
	21: uint16(37),
	22: uint16('f'),
	23: uint16(49),
	24: uint16('r'),
	25: uint16(48),
	26: uint16('t'),
	27: uint16(53),
	28: uint16('{'),
	29: uint16(22),
	30: uint16('}'),
	31: uint16(23),
}

var map_token2 = [20]uint16_t{
	0:  uint16('"'),
	1:  uint16(36),
	2:  uint16('('),
	3:  uint16(25),
	4:  uint16(')'),
	5:  uint16(27),
	6:  uint16(','),
	7:  uint16(20),
	8:  uint16('/'),
	9:  uint16(3),
	10: uint16(':'),
	11: uint16(28),
	12: uint16('\\'),
	13: uint16(4),
	14: uint16(']'),
	15: uint16(21),
	16: uint16('r'),
	17: uint16(48),
	18: uint16('}'),
	19: uint16(23),
}

var map_token3 = [34]uint16_t{
	0:  uint16('"'),
	1:  uint16(35),
	2:  uint16('\''),
	3:  uint16(38),
	4:  uint16('('),
	5:  uint16(26),
	6:  uint16(')'),
	7:  uint16(27),
	8:  uint16(','),
	9:  uint16(20),
	10: uint16('-'),
	11: uint16(34),
	12: uint16('/'),
	13: uint16(3),
	14: uint16('0'),
	15: uint16(29),
	16: uint16(':'),
	17: uint16(28),
	18: uint16('['),
	19: uint16(19),
	20: uint16(']'),
	21: uint16(21),
	22: uint16('b'),
	23: uint16(37),
	24: uint16('f'),
	25: uint16(49),
	26: uint16('r'),
	27: uint16(48),
	28: uint16('t'),
	29: uint16(53),
	30: uint16('{'),
	31: uint16(22),
	32: uint16('}'),
	33: uint16(23),
}

var map_token4 = [16]uint16_t{
	0:  uint16('('),
	1:  uint16(25),
	2:  uint16(')'),
	3:  uint16(27),
	4:  uint16(','),
	5:  uint16(20),
	6:  uint16('/'),
	7:  uint16(3),
	8:  uint16(':'),
	9:  uint16(28),
	10: uint16(']'),
	11: uint16(21),
	12: uint16('r'),
	13: uint16(48),
	14: uint16('}'),
	15: uint16(23),
}

var ts_lex_modes = [81]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	2: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	3: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	4: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	5: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	6: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	7: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	8: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	9: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	10: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	11: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	12: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	13: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	14: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	15: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(3),
	},
	16: {
		Fexternal_lex_state: uint16(3),
	},
	17: {
		Fexternal_lex_state: uint16(3),
	},
	18: {
		Fexternal_lex_state: uint16(3),
	},
	19: {
		Fexternal_lex_state: uint16(3),
	},
	20: {
		Fexternal_lex_state: uint16(3),
	},
	21: {
		Fexternal_lex_state: uint16(3),
	},
	22: {
		Fexternal_lex_state: uint16(3),
	},
	23: {
		Fexternal_lex_state: uint16(3),
	},
	24: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(4),
	},
	25: {
		Fexternal_lex_state: uint16(3),
	},
	26: {
		Fexternal_lex_state: uint16(3),
	},
	27: {
		Fexternal_lex_state: uint16(3),
	},
	28: {
		Fexternal_lex_state: uint16(3),
	},
	29: {
		Fexternal_lex_state: uint16(3),
	},
	30: {
		Fexternal_lex_state: uint16(3),
	},
	31: {
		Fexternal_lex_state: uint16(3),
	},
	32: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(4),
	},
	33: {
		Fexternal_lex_state: uint16(3),
	},
	34: {
		Fexternal_lex_state: uint16(3),
	},
	35: {
		Fexternal_lex_state: uint16(3),
	},
	36: {
		Fexternal_lex_state: uint16(3),
	},
	37: {
		Fexternal_lex_state: uint16(3),
	},
	38: {
		Fexternal_lex_state: uint16(3),
	},
	39: {
		Fexternal_lex_state: uint16(3),
	},
	40: {
		Fexternal_lex_state: uint16(3),
	},
	41: {
		Fexternal_lex_state: uint16(3),
	},
	42: {
		Fexternal_lex_state: uint16(3),
	},
	43: {
		Fexternal_lex_state: uint16(3),
	},
	44: {
		Fexternal_lex_state: uint16(3),
	},
	45: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(4),
	},
	46: {
		Fexternal_lex_state: uint16(3),
	},
	47: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	48: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	49: {
		Fexternal_lex_state: uint16(3),
	},
	50: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(3),
	},
	51: {
		Fexternal_lex_state: uint16(3),
	},
	52: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(3),
	},
	53: {
		Fexternal_lex_state: uint16(3),
	},
	54: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(3),
	},
	55: {
		Fexternal_lex_state: uint16(3),
	},
	56: {
		Fexternal_lex_state: uint16(3),
	},
	57: {
		Fexternal_lex_state: uint16(3),
	},
	58: {
		Fexternal_lex_state: uint16(3),
	},
	59: {
		Fexternal_lex_state: uint16(3),
	},
	60: {
		Fexternal_lex_state: uint16(3),
	},
	61: {
		Fexternal_lex_state: uint16(3),
	},
	62: {
		Fexternal_lex_state: uint16(3),
	},
	63: {
		Fexternal_lex_state: uint16(3),
	},
	64: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(3),
	},
	65: {
		Fexternal_lex_state: uint16(3),
	},
	66: {
		Fexternal_lex_state: uint16(3),
	},
	67: {
		Fexternal_lex_state: uint16(3),
	},
	68: {
		Fexternal_lex_state: uint16(3),
	},
	69: {
		Fexternal_lex_state: uint16(2),
	},
	70: {
		Flex_state:          uint16(15),
		Fexternal_lex_state: uint16(3),
	},
	71: {
		Fexternal_lex_state: uint16(3),
	},
	72: {
		Fexternal_lex_state: uint16(3),
	},
	73: {
		Fexternal_lex_state: uint16(3),
	},
	74: {
		Fexternal_lex_state: uint16(3),
	},
	75: {
		Fexternal_lex_state: uint16(3),
	},
	76: {
		Fexternal_lex_state: uint16(3),
	},
	77: {
		Fexternal_lex_state: uint16(3),
	},
	78: {
		Fexternal_lex_state: uint16(3),
	},
	79: {
		Fexternal_lex_state: uint16(3),
	},
	80: {
		Fexternal_lex_state: uint16(3),
	},
}

var ts_parse_table = [15][51]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
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
		17: uint16(1),
		18: uint16(1),
		19: uint16(1),
		20: uint16(1),
		21: uint16(1),
		22: uint16(3),
		23: uint16(1),
		24: uint16(1),
		25: uint16(1),
		26: uint16(3),
	},
	1: {
		1:  uint16(5),
		4:  uint16(7),
		6:  uint16(9),
		7:  uint16(11),
		10: uint16(13),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(13),
		26: uint16(3),
		27: uint16(79),
		28: uint16(78),
		29: uint16(78),
		30: uint16(78),
		31: uint16(78),
		32: uint16(78),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(78),
		41: uint16(78),
		42: uint16(78),
		43: uint16(78),
		44: uint16(78),
		46: uint16(78),
	},
	2: {
		1:  uint16(5),
		2:  uint16(27),
		4:  uint16(7),
		5:  uint16(29),
		6:  uint16(9),
		7:  uint16(11),
		10: uint16(31),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(31),
		26: uint16(3),
		28: uint16(72),
		29: uint16(72),
		30: uint16(72),
		31: uint16(72),
		32: uint16(72),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(72),
		39: uint16(58),
		41: uint16(72),
		42: uint16(72),
		43: uint16(72),
		44: uint16(72),
		46: uint16(72),
	},
	3: {
		1:  uint16(5),
		2:  uint16(33),
		4:  uint16(7),
		6:  uint16(9),
		7:  uint16(11),
		8:  uint16(35),
		10: uint16(37),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(39),
		22: uint16(3),
		25: uint16(37),
		26: uint16(3),
		28: uint16(61),
		29: uint16(61),
		30: uint16(61),
		31: uint16(61),
		32: uint16(61),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(61),
		40: uint16(59),
		41: uint16(61),
		42: uint16(61),
		43: uint16(61),
		44: uint16(61),
		46: uint16(61),
	},
	4: {
		1:  uint16(5),
		2:  uint16(41),
		3:  uint16(43),
		4:  uint16(7),
		6:  uint16(9),
		7:  uint16(11),
		10: uint16(45),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(45),
		26: uint16(3),
		28: uint16(55),
		29: uint16(55),
		30: uint16(55),
		31: uint16(55),
		32: uint16(55),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(55),
		41: uint16(55),
		42: uint16(55),
		43: uint16(55),
		44: uint16(55),
		46: uint16(55),
	},
	5: {
		1:  uint16(5),
		4:  uint16(7),
		5:  uint16(47),
		6:  uint16(9),
		7:  uint16(11),
		10: uint16(31),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(31),
		26: uint16(3),
		28: uint16(72),
		29: uint16(72),
		30: uint16(72),
		31: uint16(72),
		32: uint16(72),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(72),
		39: uint16(65),
		41: uint16(72),
		42: uint16(72),
		43: uint16(72),
		44: uint16(72),
		46: uint16(72),
	},
	6: {
		1:  uint16(5),
		4:  uint16(7),
		5:  uint16(49),
		6:  uint16(9),
		7:  uint16(11),
		10: uint16(31),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(31),
		26: uint16(3),
		28: uint16(72),
		29: uint16(72),
		30: uint16(72),
		31: uint16(72),
		32: uint16(72),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(72),
		39: uint16(65),
		41: uint16(72),
		42: uint16(72),
		43: uint16(72),
		44: uint16(72),
		46: uint16(72),
	},
	7: {
		1:  uint16(5),
		4:  uint16(7),
		6:  uint16(9),
		7:  uint16(11),
		8:  uint16(51),
		10: uint16(53),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(53),
		26: uint16(3),
		28: uint16(57),
		29: uint16(57),
		30: uint16(57),
		31: uint16(57),
		32: uint16(57),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(57),
		41: uint16(57),
		42: uint16(57),
		43: uint16(57),
		44: uint16(57),
		46: uint16(57),
	},
	8: {
		1:  uint16(5),
		4:  uint16(7),
		6:  uint16(9),
		7:  uint16(11),
		8:  uint16(55),
		10: uint16(53),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(53),
		26: uint16(3),
		28: uint16(57),
		29: uint16(57),
		30: uint16(57),
		31: uint16(57),
		32: uint16(57),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(57),
		41: uint16(57),
		42: uint16(57),
		43: uint16(57),
		44: uint16(57),
		46: uint16(57),
	},
	9: {
		1:  uint16(5),
		3:  uint16(57),
		4:  uint16(7),
		6:  uint16(9),
		7:  uint16(11),
		10: uint16(53),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(53),
		26: uint16(3),
		28: uint16(57),
		29: uint16(57),
		30: uint16(57),
		31: uint16(57),
		32: uint16(57),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(57),
		41: uint16(57),
		42: uint16(57),
		43: uint16(57),
		44: uint16(57),
		46: uint16(57),
	},
	10: {
		1:  uint16(5),
		3:  uint16(59),
		4:  uint16(7),
		6:  uint16(9),
		7:  uint16(11),
		10: uint16(53),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(53),
		26: uint16(3),
		28: uint16(57),
		29: uint16(57),
		30: uint16(57),
		31: uint16(57),
		32: uint16(57),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(57),
		41: uint16(57),
		42: uint16(57),
		43: uint16(57),
		44: uint16(57),
		46: uint16(57),
	},
	11: {
		1:  uint16(5),
		4:  uint16(7),
		6:  uint16(9),
		7:  uint16(11),
		10: uint16(31),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(31),
		26: uint16(3),
		28: uint16(72),
		29: uint16(72),
		30: uint16(72),
		31: uint16(72),
		32: uint16(72),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(72),
		39: uint16(65),
		41: uint16(72),
		42: uint16(72),
		43: uint16(72),
		44: uint16(72),
		46: uint16(72),
	},
	12: {
		1:  uint16(5),
		4:  uint16(7),
		6:  uint16(9),
		7:  uint16(11),
		10: uint16(61),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(61),
		26: uint16(3),
		28: uint16(66),
		29: uint16(66),
		30: uint16(66),
		31: uint16(66),
		32: uint16(66),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(66),
		41: uint16(66),
		42: uint16(66),
		43: uint16(66),
		44: uint16(66),
		46: uint16(66),
	},
	13: {
		1:  uint16(5),
		4:  uint16(7),
		6:  uint16(9),
		7:  uint16(11),
		10: uint16(63),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(63),
		26: uint16(3),
		28: uint16(68),
		29: uint16(68),
		30: uint16(68),
		31: uint16(68),
		32: uint16(68),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(68),
		41: uint16(68),
		42: uint16(68),
		43: uint16(68),
		44: uint16(68),
		46: uint16(68),
	},
	14: {
		1:  uint16(5),
		4:  uint16(7),
		6:  uint16(9),
		7:  uint16(11),
		10: uint16(53),
		11: uint16(15),
		12: uint16(17),
		14: uint16(19),
		15: uint16(21),
		19: uint16(23),
		20: uint16(23),
		21: uint16(25),
		22: uint16(3),
		25: uint16(53),
		26: uint16(3),
		28: uint16(57),
		29: uint16(57),
		30: uint16(57),
		31: uint16(57),
		32: uint16(57),
		33: uint16(28),
		34: uint16(54),
		35: uint16(28),
		36: uint16(41),
		37: uint16(18),
		38: uint16(57),
		41: uint16(57),
		42: uint16(57),
		43: uint16(57),
		44: uint16(57),
		46: uint16(57),
	},
}

var ts_small_parse_table = [831]uint16_t{
	0:   uint16(3),
	1:   uint16(67),
	2:   uint16(1),
	3:   uint16(anon_sym_LPAREN),
	4:   uint16(3),
	5:   uint16(2),
	6:   uint16(sym_block_comment),
	7:   uint16(sym_line_comment),
	8:   uint16(65),
	9:   uint16(6),
	11:  uint16(anon_sym_COMMA),
	12:  uint16(anon_sym_RBRACK),
	13:  uint16(anon_sym_RBRACE),
	14:  uint16(anon_sym_RPAREN),
	15:  uint16(anon_sym_COLON),
	16:  uint16(2),
	17:  uint16(3),
	18:  uint16(2),
	19:  uint16(sym_block_comment),
	20:  uint16(sym_line_comment),
	21:  uint16(69),
	22:  uint16(6),
	24:  uint16(anon_sym_COMMA),
	25:  uint16(anon_sym_RBRACK),
	26:  uint16(anon_sym_RBRACE),
	27:  uint16(anon_sym_RPAREN),
	28:  uint16(anon_sym_COLON),
	29:  uint16(2),
	30:  uint16(3),
	31:  uint16(2),
	32:  uint16(sym_block_comment),
	33:  uint16(sym_line_comment),
	34:  uint16(71),
	35:  uint16(6),
	37:  uint16(anon_sym_COMMA),
	38:  uint16(anon_sym_RBRACK),
	39:  uint16(anon_sym_RBRACE),
	40:  uint16(anon_sym_RPAREN),
	41:  uint16(anon_sym_COLON),
	42:  uint16(2),
	43:  uint16(3),
	44:  uint16(2),
	45:  uint16(sym_block_comment),
	46:  uint16(sym_line_comment),
	47:  uint16(73),
	48:  uint16(6),
	50:  uint16(anon_sym_COMMA),
	51:  uint16(anon_sym_RBRACK),
	52:  uint16(anon_sym_RBRACE),
	53:  uint16(anon_sym_RPAREN),
	54:  uint16(anon_sym_COLON),
	55:  uint16(2),
	56:  uint16(3),
	57:  uint16(2),
	58:  uint16(sym_block_comment),
	59:  uint16(sym_line_comment),
	60:  uint16(75),
	61:  uint16(6),
	63:  uint16(anon_sym_COMMA),
	64:  uint16(anon_sym_RBRACK),
	65:  uint16(anon_sym_RBRACE),
	66:  uint16(anon_sym_RPAREN),
	67:  uint16(anon_sym_COLON),
	68:  uint16(2),
	69:  uint16(3),
	70:  uint16(2),
	71:  uint16(sym_block_comment),
	72:  uint16(sym_line_comment),
	73:  uint16(77),
	74:  uint16(6),
	76:  uint16(anon_sym_COMMA),
	77:  uint16(anon_sym_RBRACK),
	78:  uint16(anon_sym_RBRACE),
	79:  uint16(anon_sym_RPAREN),
	80:  uint16(anon_sym_COLON),
	81:  uint16(2),
	82:  uint16(3),
	83:  uint16(2),
	84:  uint16(sym_block_comment),
	85:  uint16(sym_line_comment),
	86:  uint16(79),
	87:  uint16(6),
	89:  uint16(anon_sym_COMMA),
	90:  uint16(anon_sym_RBRACK),
	91:  uint16(anon_sym_RBRACE),
	92:  uint16(anon_sym_RPAREN),
	93:  uint16(anon_sym_COLON),
	94:  uint16(2),
	95:  uint16(3),
	96:  uint16(2),
	97:  uint16(sym_block_comment),
	98:  uint16(sym_line_comment),
	99:  uint16(81),
	100: uint16(6),
	102: uint16(anon_sym_COMMA),
	103: uint16(anon_sym_RBRACK),
	104: uint16(anon_sym_RBRACE),
	105: uint16(anon_sym_RPAREN),
	106: uint16(anon_sym_COLON),
	107: uint16(2),
	108: uint16(3),
	109: uint16(2),
	110: uint16(sym_block_comment),
	111: uint16(sym_line_comment),
	112: uint16(83),
	113: uint16(6),
	115: uint16(anon_sym_COMMA),
	116: uint16(anon_sym_RBRACK),
	117: uint16(anon_sym_RBRACE),
	118: uint16(anon_sym_RPAREN),
	119: uint16(anon_sym_COLON),
	120: uint16(5),
	121: uint16(85),
	122: uint16(1),
	123: uint16(anon_sym_DQUOTE),
	124: uint16(90),
	125: uint16(1),
	126: uint16(sym__string_content),
	127: uint16(3),
	128: uint16(2),
	129: uint16(sym_block_comment),
	130: uint16(sym_line_comment),
	131: uint16(87),
	132: uint16(2),
	133: uint16(aux_sym__escape_sequence_token1),
	134: uint16(sym_escape_sequence),
	135: uint16(24),
	136: uint16(2),
	137: uint16(sym__escape_sequence),
	138: uint16(aux_sym_string_repeat1),
	139: uint16(2),
	140: uint16(3),
	141: uint16(2),
	142: uint16(sym_block_comment),
	143: uint16(sym_line_comment),
	144: uint16(93),
	145: uint16(6),
	147: uint16(anon_sym_COMMA),
	148: uint16(anon_sym_RBRACK),
	149: uint16(anon_sym_RBRACE),
	150: uint16(anon_sym_RPAREN),
	151: uint16(anon_sym_COLON),
	152: uint16(2),
	153: uint16(3),
	154: uint16(2),
	155: uint16(sym_block_comment),
	156: uint16(sym_line_comment),
	157: uint16(95),
	158: uint16(6),
	160: uint16(anon_sym_COMMA),
	161: uint16(anon_sym_RBRACK),
	162: uint16(anon_sym_RBRACE),
	163: uint16(anon_sym_RPAREN),
	164: uint16(anon_sym_COLON),
	165: uint16(2),
	166: uint16(3),
	167: uint16(2),
	168: uint16(sym_block_comment),
	169: uint16(sym_line_comment),
	170: uint16(97),
	171: uint16(6),
	173: uint16(anon_sym_COMMA),
	174: uint16(anon_sym_RBRACK),
	175: uint16(anon_sym_RBRACE),
	176: uint16(anon_sym_RPAREN),
	177: uint16(anon_sym_COLON),
	178: uint16(2),
	179: uint16(3),
	180: uint16(2),
	181: uint16(sym_block_comment),
	182: uint16(sym_line_comment),
	183: uint16(99),
	184: uint16(6),
	186: uint16(anon_sym_COMMA),
	187: uint16(anon_sym_RBRACK),
	188: uint16(anon_sym_RBRACE),
	189: uint16(anon_sym_RPAREN),
	190: uint16(anon_sym_COLON),
	191: uint16(2),
	192: uint16(3),
	193: uint16(2),
	194: uint16(sym_block_comment),
	195: uint16(sym_line_comment),
	196: uint16(101),
	197: uint16(6),
	199: uint16(anon_sym_COMMA),
	200: uint16(anon_sym_RBRACK),
	201: uint16(anon_sym_RBRACE),
	202: uint16(anon_sym_RPAREN),
	203: uint16(anon_sym_COLON),
	204: uint16(2),
	205: uint16(3),
	206: uint16(2),
	207: uint16(sym_block_comment),
	208: uint16(sym_line_comment),
	209: uint16(103),
	210: uint16(6),
	212: uint16(anon_sym_COMMA),
	213: uint16(anon_sym_RBRACK),
	214: uint16(anon_sym_RBRACE),
	215: uint16(anon_sym_RPAREN),
	216: uint16(anon_sym_COLON),
	217: uint16(2),
	218: uint16(3),
	219: uint16(2),
	220: uint16(sym_block_comment),
	221: uint16(sym_line_comment),
	222: uint16(105),
	223: uint16(6),
	225: uint16(anon_sym_COMMA),
	226: uint16(anon_sym_RBRACK),
	227: uint16(anon_sym_RBRACE),
	228: uint16(anon_sym_RPAREN),
	229: uint16(anon_sym_COLON),
	230: uint16(5),
	231: uint16(107),
	232: uint16(1),
	233: uint16(anon_sym_DQUOTE),
	234: uint16(111),
	235: uint16(1),
	236: uint16(sym__string_content),
	237: uint16(3),
	238: uint16(2),
	239: uint16(sym_block_comment),
	240: uint16(sym_line_comment),
	241: uint16(109),
	242: uint16(2),
	243: uint16(aux_sym__escape_sequence_token1),
	244: uint16(sym_escape_sequence),
	245: uint16(24),
	246: uint16(2),
	247: uint16(sym__escape_sequence),
	248: uint16(aux_sym_string_repeat1),
	249: uint16(2),
	250: uint16(3),
	251: uint16(2),
	252: uint16(sym_block_comment),
	253: uint16(sym_line_comment),
	254: uint16(113),
	255: uint16(6),
	257: uint16(anon_sym_COMMA),
	258: uint16(anon_sym_RBRACK),
	259: uint16(anon_sym_RBRACE),
	260: uint16(anon_sym_RPAREN),
	261: uint16(anon_sym_COLON),
	262: uint16(2),
	263: uint16(3),
	264: uint16(2),
	265: uint16(sym_block_comment),
	266: uint16(sym_line_comment),
	267: uint16(115),
	268: uint16(6),
	270: uint16(anon_sym_COMMA),
	271: uint16(anon_sym_RBRACK),
	272: uint16(anon_sym_RBRACE),
	273: uint16(anon_sym_RPAREN),
	274: uint16(anon_sym_COLON),
	275: uint16(2),
	276: uint16(3),
	277: uint16(2),
	278: uint16(sym_block_comment),
	279: uint16(sym_line_comment),
	280: uint16(117),
	281: uint16(6),
	283: uint16(anon_sym_COMMA),
	284: uint16(anon_sym_RBRACK),
	285: uint16(anon_sym_RBRACE),
	286: uint16(anon_sym_RPAREN),
	287: uint16(anon_sym_COLON),
	288: uint16(2),
	289: uint16(3),
	290: uint16(2),
	291: uint16(sym_block_comment),
	292: uint16(sym_line_comment),
	293: uint16(119),
	294: uint16(6),
	296: uint16(anon_sym_COMMA),
	297: uint16(anon_sym_RBRACK),
	298: uint16(anon_sym_RBRACE),
	299: uint16(anon_sym_RPAREN),
	300: uint16(anon_sym_COLON),
	301: uint16(2),
	302: uint16(3),
	303: uint16(2),
	304: uint16(sym_block_comment),
	305: uint16(sym_line_comment),
	306: uint16(121),
	307: uint16(6),
	309: uint16(anon_sym_COMMA),
	310: uint16(anon_sym_RBRACK),
	311: uint16(anon_sym_RBRACE),
	312: uint16(anon_sym_RPAREN),
	313: uint16(anon_sym_COLON),
	314: uint16(2),
	315: uint16(3),
	316: uint16(2),
	317: uint16(sym_block_comment),
	318: uint16(sym_line_comment),
	319: uint16(123),
	320: uint16(6),
	322: uint16(anon_sym_COMMA),
	323: uint16(anon_sym_RBRACK),
	324: uint16(anon_sym_RBRACE),
	325: uint16(anon_sym_RPAREN),
	326: uint16(anon_sym_COLON),
	327: uint16(2),
	328: uint16(3),
	329: uint16(2),
	330: uint16(sym_block_comment),
	331: uint16(sym_line_comment),
	332: uint16(125),
	333: uint16(6),
	335: uint16(anon_sym_COMMA),
	336: uint16(anon_sym_RBRACK),
	337: uint16(anon_sym_RBRACE),
	338: uint16(anon_sym_RPAREN),
	339: uint16(anon_sym_COLON),
	340: uint16(2),
	341: uint16(3),
	342: uint16(2),
	343: uint16(sym_block_comment),
	344: uint16(sym_line_comment),
	345: uint16(127),
	346: uint16(6),
	348: uint16(anon_sym_COMMA),
	349: uint16(anon_sym_RBRACK),
	350: uint16(anon_sym_RBRACE),
	351: uint16(anon_sym_RPAREN),
	352: uint16(anon_sym_COLON),
	353: uint16(2),
	354: uint16(3),
	355: uint16(2),
	356: uint16(sym_block_comment),
	357: uint16(sym_line_comment),
	358: uint16(129),
	359: uint16(6),
	361: uint16(anon_sym_COMMA),
	362: uint16(anon_sym_RBRACK),
	363: uint16(anon_sym_RBRACE),
	364: uint16(anon_sym_RPAREN),
	365: uint16(anon_sym_COLON),
	366: uint16(2),
	367: uint16(3),
	368: uint16(2),
	369: uint16(sym_block_comment),
	370: uint16(sym_line_comment),
	371: uint16(131),
	372: uint16(6),
	374: uint16(anon_sym_COMMA),
	375: uint16(anon_sym_RBRACK),
	376: uint16(anon_sym_RBRACE),
	377: uint16(anon_sym_RPAREN),
	378: uint16(anon_sym_COLON),
	379: uint16(2),
	380: uint16(3),
	381: uint16(2),
	382: uint16(sym_block_comment),
	383: uint16(sym_line_comment),
	384: uint16(133),
	385: uint16(6),
	387: uint16(anon_sym_COMMA),
	388: uint16(anon_sym_RBRACK),
	389: uint16(anon_sym_RBRACE),
	390: uint16(anon_sym_RPAREN),
	391: uint16(anon_sym_COLON),
	392: uint16(2),
	393: uint16(3),
	394: uint16(2),
	395: uint16(sym_block_comment),
	396: uint16(sym_line_comment),
	397: uint16(135),
	398: uint16(6),
	400: uint16(anon_sym_COMMA),
	401: uint16(anon_sym_RBRACK),
	402: uint16(anon_sym_RBRACE),
	403: uint16(anon_sym_RPAREN),
	404: uint16(anon_sym_COLON),
	405: uint16(5),
	406: uint16(137),
	407: uint16(1),
	408: uint16(anon_sym_DQUOTE),
	409: uint16(141),
	410: uint16(1),
	411: uint16(sym__string_content),
	412: uint16(3),
	413: uint16(2),
	414: uint16(sym_block_comment),
	415: uint16(sym_line_comment),
	416: uint16(139),
	417: uint16(2),
	418: uint16(aux_sym__escape_sequence_token1),
	419: uint16(sym_escape_sequence),
	420: uint16(32),
	421: uint16(2),
	422: uint16(sym__escape_sequence),
	423: uint16(aux_sym_string_repeat1),
	424: uint16(2),
	425: uint16(3),
	426: uint16(2),
	427: uint16(sym_block_comment),
	428: uint16(sym_line_comment),
	429: uint16(143),
	430: uint16(6),
	432: uint16(anon_sym_COMMA),
	433: uint16(anon_sym_RBRACK),
	434: uint16(anon_sym_RBRACE),
	435: uint16(anon_sym_RPAREN),
	436: uint16(anon_sym_COLON),
	437: uint16(5),
	438: uint16(3),
	439: uint16(1),
	440: uint16(sym_block_comment),
	441: uint16(145),
	442: uint16(1),
	443: uint16(anon_sym_SQUOTE),
	444: uint16(149),
	445: uint16(1),
	446: uint16(sym_line_comment),
	447: uint16(75),
	448: uint16(1),
	449: uint16(sym__escape_sequence),
	450: uint16(147),
	451: uint16(3),
	452: uint16(aux_sym_char_token1),
	453: uint16(aux_sym__escape_sequence_token1),
	454: uint16(sym_escape_sequence),
	455: uint16(5),
	456: uint16(3),
	457: uint16(1),
	458: uint16(sym_block_comment),
	459: uint16(149),
	460: uint16(1),
	461: uint16(sym_line_comment),
	462: uint16(151),
	463: uint16(1),
	464: uint16(anon_sym_SQUOTE),
	465: uint16(76),
	466: uint16(1),
	467: uint16(sym__escape_sequence),
	468: uint16(153),
	469: uint16(3),
	470: uint16(aux_sym_char_token1),
	471: uint16(aux_sym__escape_sequence_token1),
	472: uint16(sym_escape_sequence),
	473: uint16(4),
	474: uint16(155),
	475: uint16(1),
	476: uint16(anon_sym_COMMA),
	477: uint16(49),
	478: uint16(1),
	479: uint16(aux_sym_array_repeat1),
	480: uint16(3),
	481: uint16(2),
	482: uint16(sym_block_comment),
	483: uint16(sym_line_comment),
	484: uint16(158),
	485: uint16(2),
	486: uint16(anon_sym_RBRACK),
	487: uint16(anon_sym_RPAREN),
	488: uint16(4),
	489: uint16(67),
	490: uint16(1),
	491: uint16(anon_sym_LPAREN),
	492: uint16(160),
	493: uint16(1),
	494: uint16(anon_sym_COLON),
	495: uint16(3),
	496: uint16(2),
	497: uint16(sym_block_comment),
	498: uint16(sym_line_comment),
	499: uint16(65),
	500: uint16(2),
	501: uint16(anon_sym_COMMA),
	502: uint16(anon_sym_RPAREN),
	503: uint16(4),
	504: uint16(162),
	505: uint16(1),
	506: uint16(anon_sym_COMMA),
	507: uint16(165),
	508: uint16(1),
	509: uint16(anon_sym_RBRACE),
	510: uint16(51),
	511: uint16(1),
	512: uint16(aux_sym_map_repeat1),
	513: uint16(3),
	514: uint16(2),
	515: uint16(sym_block_comment),
	516: uint16(sym_line_comment),
	517: uint16(4),
	518: uint16(167),
	519: uint16(1),
	520: uint16(anon_sym_RPAREN),
	521: uint16(169),
	522: uint16(1),
	523: uint16(sym_identifier),
	524: uint16(67),
	525: uint16(1),
	526: uint16(sym_struct_entry),
	527: uint16(3),
	528: uint16(2),
	529: uint16(sym_block_comment),
	530: uint16(sym_line_comment),
	531: uint16(4),
	532: uint16(55),
	533: uint16(1),
	534: uint16(anon_sym_RPAREN),
	535: uint16(171),
	536: uint16(1),
	537: uint16(anon_sym_COMMA),
	538: uint16(49),
	539: uint16(1),
	540: uint16(aux_sym_array_repeat1),
	541: uint16(3),
	542: uint16(2),
	543: uint16(sym_block_comment),
	544: uint16(sym_line_comment),
	545: uint16(4),
	546: uint16(173),
	547: uint16(1),
	548: uint16(anon_sym_LPAREN),
	549: uint16(36),
	550: uint16(1),
	551: uint16(sym__struct_body),
	552: uint16(37),
	553: uint16(1),
	554: uint16(sym_tuple),
	555: uint16(3),
	556: uint16(2),
	557: uint16(sym_block_comment),
	558: uint16(sym_line_comment),
	559: uint16(4),
	560: uint16(175),
	561: uint16(1),
	562: uint16(anon_sym_COMMA),
	563: uint16(177),
	564: uint16(1),
	565: uint16(anon_sym_RBRACK),
	566: uint16(56),
	567: uint16(1),
	568: uint16(aux_sym_array_repeat1),
	569: uint16(3),
	570: uint16(2),
	571: uint16(sym_block_comment),
	572: uint16(sym_line_comment),
	573: uint16(4),
	574: uint16(57),
	575: uint16(1),
	576: uint16(anon_sym_RBRACK),
	577: uint16(179),
	578: uint16(1),
	579: uint16(anon_sym_COMMA),
	580: uint16(49),
	581: uint16(1),
	582: uint16(aux_sym_array_repeat1),
	583: uint16(3),
	584: uint16(2),
	585: uint16(sym_block_comment),
	586: uint16(sym_line_comment),
	587: uint16(2),
	588: uint16(3),
	589: uint16(2),
	590: uint16(sym_block_comment),
	591: uint16(sym_line_comment),
	592: uint16(158),
	593: uint16(3),
	594: uint16(anon_sym_COMMA),
	595: uint16(anon_sym_RBRACK),
	596: uint16(anon_sym_RPAREN),
	597: uint16(4),
	598: uint16(181),
	599: uint16(1),
	600: uint16(anon_sym_COMMA),
	601: uint16(183),
	602: uint16(1),
	603: uint16(anon_sym_RBRACE),
	604: uint16(62),
	605: uint16(1),
	606: uint16(aux_sym_map_repeat1),
	607: uint16(3),
	608: uint16(2),
	609: uint16(sym_block_comment),
	610: uint16(sym_line_comment),
	611: uint16(4),
	612: uint16(185),
	613: uint16(1),
	614: uint16(anon_sym_COMMA),
	615: uint16(187),
	616: uint16(1),
	617: uint16(anon_sym_RPAREN),
	618: uint16(63),
	619: uint16(1),
	620: uint16(aux_sym__struct_body_repeat1),
	621: uint16(3),
	622: uint16(2),
	623: uint16(sym_block_comment),
	624: uint16(sym_line_comment),
	625: uint16(4),
	626: uint16(189),
	627: uint16(1),
	628: uint16(anon_sym_COMMA),
	629: uint16(192),
	630: uint16(1),
	631: uint16(anon_sym_RPAREN),
	632: uint16(60),
	633: uint16(1),
	634: uint16(aux_sym__struct_body_repeat1),
	635: uint16(3),
	636: uint16(2),
	637: uint16(sym_block_comment),
	638: uint16(sym_line_comment),
	639: uint16(4),
	640: uint16(194),
	641: uint16(1),
	642: uint16(anon_sym_COMMA),
	643: uint16(196),
	644: uint16(1),
	645: uint16(anon_sym_RPAREN),
	646: uint16(53),
	647: uint16(1),
	648: uint16(aux_sym_array_repeat1),
	649: uint16(3),
	650: uint16(2),
	651: uint16(sym_block_comment),
	652: uint16(sym_line_comment),
	653: uint16(4),
	654: uint16(47),
	655: uint16(1),
	656: uint16(anon_sym_RBRACE),
	657: uint16(198),
	658: uint16(1),
	659: uint16(anon_sym_COMMA),
	660: uint16(51),
	661: uint16(1),
	662: uint16(aux_sym_map_repeat1),
	663: uint16(3),
	664: uint16(2),
	665: uint16(sym_block_comment),
	666: uint16(sym_line_comment),
	667: uint16(4),
	668: uint16(200),
	669: uint16(1),
	670: uint16(anon_sym_COMMA),
	671: uint16(202),
	672: uint16(1),
	673: uint16(anon_sym_RPAREN),
	674: uint16(60),
	675: uint16(1),
	676: uint16(aux_sym__struct_body_repeat1),
	677: uint16(3),
	678: uint16(2),
	679: uint16(sym_block_comment),
	680: uint16(sym_line_comment),
	681: uint16(4),
	682: uint16(169),
	683: uint16(1),
	684: uint16(sym_identifier),
	685: uint16(202),
	686: uint16(1),
	687: uint16(anon_sym_RPAREN),
	688: uint16(67),
	689: uint16(1),
	690: uint16(sym_struct_entry),
	691: uint16(3),
	692: uint16(2),
	693: uint16(sym_block_comment),
	694: uint16(sym_line_comment),
	695: uint16(2),
	696: uint16(3),
	697: uint16(2),
	698: uint16(sym_block_comment),
	699: uint16(sym_line_comment),
	700: uint16(165),
	701: uint16(2),
	702: uint16(anon_sym_COMMA),
	703: uint16(anon_sym_RBRACE),
	704: uint16(2),
	705: uint16(3),
	706: uint16(2),
	707: uint16(sym_block_comment),
	708: uint16(sym_line_comment),
	709: uint16(204),
	710: uint16(2),
	711: uint16(anon_sym_COMMA),
	712: uint16(anon_sym_RPAREN),
	713: uint16(2),
	714: uint16(3),
	715: uint16(2),
	716: uint16(sym_block_comment),
	717: uint16(sym_line_comment),
	718: uint16(192),
	719: uint16(2),
	720: uint16(anon_sym_COMMA),
	721: uint16(anon_sym_RPAREN),
	722: uint16(2),
	723: uint16(3),
	724: uint16(2),
	725: uint16(sym_block_comment),
	726: uint16(sym_line_comment),
	727: uint16(206),
	728: uint16(2),
	729: uint16(anon_sym_COMMA),
	730: uint16(anon_sym_RBRACE),
	731: uint16(2),
	732: uint16(3),
	733: uint16(2),
	734: uint16(sym_block_comment),
	735: uint16(sym_line_comment),
	736: uint16(208),
	737: uint16(2),
	738: uint16(sym_float),
	739: uint16(sym_integer),
	740: uint16(3),
	741: uint16(169),
	742: uint16(1),
	743: uint16(sym_identifier),
	744: uint16(67),
	745: uint16(1),
	746: uint16(sym_struct_entry),
	747: uint16(3),
	748: uint16(2),
	749: uint16(sym_block_comment),
	750: uint16(sym_line_comment),
	751: uint16(2),
	752: uint16(183),
	753: uint16(1),
	754: uint16(anon_sym_RBRACE),
	755: uint16(3),
	756: uint16(2),
	757: uint16(sym_block_comment),
	758: uint16(sym_line_comment),
	759: uint16(2),
	760: uint16(210),
	761: uint16(1),
	762: uint16(anon_sym_COLON),
	763: uint16(3),
	764: uint16(2),
	765: uint16(sym_block_comment),
	766: uint16(sym_line_comment),
	767: uint16(2),
	768: uint16(160),
	769: uint16(1),
	770: uint16(anon_sym_COLON),
	771: uint16(3),
	772: uint16(2),
	773: uint16(sym_block_comment),
	774: uint16(sym_line_comment),
	775: uint16(2),
	776: uint16(187),
	777: uint16(1),
	778: uint16(anon_sym_RPAREN),
	779: uint16(3),
	780: uint16(2),
	781: uint16(sym_block_comment),
	782: uint16(sym_line_comment),
	783: uint16(2),
	784: uint16(212),
	785: uint16(1),
	786: uint16(anon_sym_SQUOTE),
	787: uint16(3),
	788: uint16(2),
	789: uint16(sym_block_comment),
	790: uint16(sym_line_comment),
	791: uint16(2),
	792: uint16(214),
	793: uint16(1),
	794: uint16(anon_sym_SQUOTE),
	795: uint16(3),
	796: uint16(2),
	797: uint16(sym_block_comment),
	798: uint16(sym_line_comment),
	799: uint16(2),
	800: uint16(177),
	801: uint16(1),
	802: uint16(anon_sym_RBRACK),
	803: uint16(3),
	804: uint16(2),
	805: uint16(sym_block_comment),
	806: uint16(sym_line_comment),
	807: uint16(2),
	808: uint16(216),
	809: uint16(1),
	811: uint16(3),
	812: uint16(2),
	813: uint16(sym_block_comment),
	814: uint16(sym_line_comment),
	815: uint16(2),
	816: uint16(218),
	817: uint16(1),
	819: uint16(3),
	820: uint16(2),
	821: uint16(sym_block_comment),
	822: uint16(sym_line_comment),
	823: uint16(2),
	824: uint16(220),
	825: uint16(1),
	826: uint16(anon_sym_SQUOTE),
	827: uint16(3),
	828: uint16(2),
	829: uint16(sym_block_comment),
	830: uint16(sym_line_comment),
}

var ts_small_parse_table_map = [66]uint32_t{
	1:  uint32(16),
	2:  uint32(29),
	3:  uint32(42),
	4:  uint32(55),
	5:  uint32(68),
	6:  uint32(81),
	7:  uint32(94),
	8:  uint32(107),
	9:  uint32(120),
	10: uint32(139),
	11: uint32(152),
	12: uint32(165),
	13: uint32(178),
	14: uint32(191),
	15: uint32(204),
	16: uint32(217),
	17: uint32(230),
	18: uint32(249),
	19: uint32(262),
	20: uint32(275),
	21: uint32(288),
	22: uint32(301),
	23: uint32(314),
	24: uint32(327),
	25: uint32(340),
	26: uint32(353),
	27: uint32(366),
	28: uint32(379),
	29: uint32(392),
	30: uint32(405),
	31: uint32(424),
	32: uint32(437),
	33: uint32(455),
	34: uint32(473),
	35: uint32(488),
	36: uint32(503),
	37: uint32(517),
	38: uint32(531),
	39: uint32(545),
	40: uint32(559),
	41: uint32(573),
	42: uint32(587),
	43: uint32(597),
	44: uint32(611),
	45: uint32(625),
	46: uint32(639),
	47: uint32(653),
	48: uint32(667),
	49: uint32(681),
	50: uint32(695),
	51: uint32(704),
	52: uint32(713),
	53: uint32(722),
	54: uint32(731),
	55: uint32(740),
	56: uint32(751),
	57: uint32(759),
	58: uint32(767),
	59: uint32(775),
	60: uint32(783),
	61: uint32(791),
	62: uint32(799),
	63: uint32(807),
	64: uint32(815),
	65: uint32(823),
}

var ts_parse_actions = [222]TSParseActionEntry{
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(71)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(61)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(57)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(35)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(66)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(68)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_enum_variant),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_struct_name),
	})))),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_tuple),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_map),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__named_struct),
		Fproduction_id: uint16(2),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_array),
	})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_array),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_char),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__struct_body),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_map),
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
		Fsymbol:      uint16(aux_sym_string_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Fsymbol:      uint16(aux_sym_string_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_unit_struct),
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
		Fsymbol:      uint16(sym__struct_body),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_map),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_struct),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_map),
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
		Fsymbol:      uint16(sym_negative),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_string),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_string),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_char),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__named_struct),
		Fproduction_id: uint16(3),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__tuple_struct),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_array),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_char),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_struct),
		Fproduction_id: uint16(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_tuple),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_tuple),
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
		Fsymbol:      uint16(sym__struct_body),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym__struct_body),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_array_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(14)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_array_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(aux_sym_map_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(11)),
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
		Fsymbol:      uint16(aux_sym_map_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(17)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		Fsymbol:      uint16(aux_sym__struct_body_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(70)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__struct_body_repeat1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_struct_entry),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_map_entry),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
	219: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
	}})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token__string_content = 0
const ts_external_token_raw_string = 1
const ts_external_token_float = 2
const ts_external_token_block_comment = 3

var ts_external_scanner_symbol_map = [4]TSSymbol{
	0: uint16(sym__string_content),
	1: uint16(sym_raw_string),
	2: uint16(sym_float),
	3: uint16(sym_block_comment),
}

var ts_external_scanner_states = [5][4]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
		3: libc.BoolUint8(true1 != 0),
	},
	2: {
		2: libc.BoolUint8(true1 != 0),
		3: libc.BoolUint8(true1 != 0),
	},
	3: {
		3: libc.BoolUint8(true1 != 0),
	},
	4: {
		0: libc.BoolUint8(true1 != 0),
		3: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_ron(tls *libc.TLS) (r uintptr) {
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
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_ron_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_ron_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_ron_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_ron_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_ron_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00[\x00,\x00]\x00{\x00}\x00()\x00(\x00)\x00:\x00integer\x00-\x00\"\x00b\x00'\x00char_token1\x00_escape_sequence_token1\x00escape_sequence\x00true\x00false\x00identifier\x00line_comment\x00_string_content\x00raw_string\x00float\x00block_comment\x00source_file\x00_value\x00enum_variant\x00array\x00map\x00struct\x00unit_struct\x00struct_name\x00_tuple_struct\x00_named_struct\x00_struct_body\x00tuple\x00map_entry\x00struct_entry\x00_literal\x00negative\x00string\x00char\x00_escape_sequence\x00boolean\x00array_repeat1\x00map_repeat1\x00_struct_body_repeat1\x00string_repeat1\x00body\x00"
