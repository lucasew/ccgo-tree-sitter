// Code generated for linux/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-wgsl/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-wgsl -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && amd64

package grammar_wgsl

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
const EXTERNAL_TOKEN_COUNT = 1
const FIELD_COUNT = 13
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
const LARGE_STATE_COUNT = 32
const MAX_ALIAS_SEQUENCE_LENGTH = 8
const PRODUCTION_ID_COUNT = 18
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 360
const SYMBOL_COUNT = 211
const TOKEN_COUNT = 136
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

type wint_t = uint32

type wctype_t = uint64

type locale_t = uintptr

type wctrans_t = uintptr

type TokenType = int32

const BLOCK_COMMENT = 0

func tree_sitter_wgsl_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_wgsl_external_scanner_destroy(tls *libc.TLS, p uintptr) {
}

func tree_sitter_wgsl_external_scanner_reset(tls *libc.TLS, p uintptr) {
}

func tree_sitter_wgsl_external_scanner_serialize(tls *libc.TLS, p uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_wgsl_external_scanner_deserialize(tls *libc.TLS, p uintptr, b uintptr, n uint32) {
}

func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
}

func at_eof(tls *libc.TLS, lexer uintptr) (r uint8) {
	return (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer)
}

// C documentation
//
//	// based on https://github.com/tree-sitter/tree-sitter-rust/blob/f7fb205c424b0962de59b26b931fe484e1262b35/src/scanner.c
func tree_sitter_wgsl_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var comment_depth uint32
	_ = comment_depth
	for libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
	}
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('/') {
		return libc.BoolUint8(false1 != 0)
	}
	advance(tls, lexer)
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('*') {
		return libc.BoolUint8(false1 != 0)
	}
	advance(tls, lexer)
	comment_depth = uint32(1)
	for int32(true1) != 0 {
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
			advance(tls, lexer)
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*') {
				advance(tls, lexer)
				comment_depth = comment_depth + uint32(1)
			}
		} else {
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*') {
				advance(tls, lexer)
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
					advance(tls, lexer)
					comment_depth = comment_depth - uint32(1)
					if comment_depth == uint32(0) {
						(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BLOCK_COMMENT)
						return libc.BoolUint8(true1 != 0)
					}
				}
			} else {
				if at_eof(tls, lexer) != 0 {
					return libc.BoolUint8(false1 != 0)
				} else {
					advance(tls, lexer)
				}
			}
		}
	}
	return r
}

const sym_identifier = 1
const sym_line_comment = 2
const anon_sym_SEMI = 3
const anon_sym_EQ = 4
const anon_sym_let = 5
const anon_sym_override = 6
const anon_sym_LPAREN = 7
const anon_sym_COMMA = 8
const anon_sym_RPAREN = 9
const anon_sym_type = 10
const anon_sym_fn = 11
const anon_sym_DASH_GT = 12
const anon_sym_struct = 13
const anon_sym_LBRACE = 14
const anon_sym_RBRACE = 15
const anon_sym_enable = 16
const anon_sym_AT = 17
const anon_sym__ = 18
const anon_sym_PLUS_EQ = 19
const anon_sym_DASH_EQ = 20
const anon_sym_STAR_EQ = 21
const anon_sym_SLASH_EQ = 22
const anon_sym_PERCENT_EQ = 23
const anon_sym_AMP_EQ = 24
const anon_sym_PIPE_EQ = 25
const anon_sym_CARET_EQ = 26
const anon_sym_if = 27
const anon_sym_else = 28
const anon_sym_switch = 29
const anon_sym_case = 30
const anon_sym_COLON = 31
const anon_sym_default = 32
const anon_sym_fallthrough = 33
const anon_sym_loop = 34
const anon_sym_for = 35
const anon_sym_while = 36
const anon_sym_break = 37
const anon_sym_continue = 38
const anon_sym_continuing = 39
const anon_sym_return = 40
const anon_sym_discard = 41
const anon_sym_var = 42
const anon_sym_LT = 43
const anon_sym_GT = 44
const anon_sym_PLUS_PLUS = 45
const anon_sym_DASH_DASH = 46
const sym_int_literal = 47
const aux_sym_float_literal_token1 = 48
const aux_sym_float_literal_token2 = 49
const anon_sym_true = 50
const anon_sym_false = 51
const anon_sym_bool = 52
const anon_sym_u32 = 53
const anon_sym_i32 = 54
const anon_sym_f32 = 55
const anon_sym_f16 = 56
const anon_sym_array = 57
const anon_sym_ptr = 58
const anon_sym_sampler = 59
const anon_sym_sampler_comparison = 60
const anon_sym_texture_depth_2d = 61
const anon_sym_texture_depth_2d_array = 62
const anon_sym_texture_depth_cube = 63
const anon_sym_texture_depth_cube_array = 64
const anon_sym_texture_depth_multisampled_2d = 65
const anon_sym_texture_1d = 66
const anon_sym_texture_2d = 67
const anon_sym_texture_2d_array = 68
const anon_sym_texture_3d = 69
const anon_sym_texture_cube = 70
const anon_sym_texture_cube_array = 71
const anon_sym_texture_multisampled_2d = 72
const anon_sym_texture_storage_1d = 73
const anon_sym_texture_storage_2d = 74
const anon_sym_texture_storage_2d_array = 75
const anon_sym_texture_storage_3d = 76
const anon_sym_vec2 = 77
const anon_sym_vec3 = 78
const anon_sym_vec4 = 79
const anon_sym_mat2x2 = 80
const anon_sym_mat2x3 = 81
const anon_sym_mat2x4 = 82
const anon_sym_mat3x2 = 83
const anon_sym_mat3x3 = 84
const anon_sym_mat3x4 = 85
const anon_sym_mat4x2 = 86
const anon_sym_mat4x3 = 87
const anon_sym_mat4x4 = 88
const anon_sym_rgba8unorm = 89
const anon_sym_rgba8snorm = 90
const anon_sym_rgba8uint = 91
const anon_sym_rgba8sint = 92
const anon_sym_rgba16uint = 93
const anon_sym_rgba16sint = 94
const anon_sym_rgba16float = 95
const anon_sym_r32uint = 96
const anon_sym_r32sint = 97
const anon_sym_r32float = 98
const anon_sym_rg32uint = 99
const anon_sym_rg32sint = 100
const anon_sym_rg32float = 101
const anon_sym_rgba32uint = 102
const anon_sym_rgba32sint = 103
const anon_sym_rgba32float = 104
const anon_sym_function = 105
const anon_sym_private = 106
const anon_sym_workgroup = 107
const anon_sym_uniform = 108
const anon_sym_storage = 109
const anon_sym_read = 110
const anon_sym_write = 111
const anon_sym_read_write = 112
const anon_sym_bitcast = 113
const anon_sym_PIPE_PIPE = 114
const anon_sym_AMP_AMP = 115
const anon_sym_PIPE = 116
const anon_sym_CARET = 117
const anon_sym_AMP = 118
const anon_sym_EQ_EQ = 119
const anon_sym_BANG_EQ = 120
const anon_sym_LT_EQ = 121
const anon_sym_GT_EQ = 122
const anon_sym_LT_LT = 123
const anon_sym_GT_GT = 124
const anon_sym_PLUS = 125
const anon_sym_DASH = 126
const anon_sym_STAR = 127
const anon_sym_SLASH = 128
const anon_sym_PERCENT = 129
const anon_sym_BANG = 130
const anon_sym_TILDE = 131
const anon_sym_LBRACK = 132
const anon_sym_RBRACK = 133
const anon_sym_DOT = 134
const sym_block_comment = 135
const sym_source_file = 136
const sym__declaration = 137
const sym_global_variable_declaration = 138
const sym_global_constant_declaration = 139
const sym_type_alias_declaration = 140
const sym_const_expression = 141
const sym_function_declaration = 142
const sym_function_return_type_declaration = 143
const sym_struct_declaration = 144
const sym_struct_member = 145
const sym_enable_directive = 146
const sym_attribute = 147
const sym__literal_or_identifier = 148
const sym_parameter_list = 149
const sym_parameter = 150
const sym__statement = 151
const sym_compound_statement = 152
const sym_assignment_statement = 153
const sym_compound_assignment_operator = 154
const sym_if_statement = 155
const sym_else_statement = 156
const sym_switch_statement = 157
const sym_switch_body = 158
const sym_case_selectors = 159
const sym_case_compound_statement = 160
const sym_fallthrough_statement = 161
const sym_loop_statement = 162
const sym_for_statement = 163
const sym_for_header = 164
const sym_while_statement = 165
const sym_break_statement = 166
const sym_break_if_statement = 167
const sym_continue_statement = 168
const sym_continuing_statement = 169
const sym_continuing_compound_statement = 170
const sym_return_statement = 171
const sym_discard_statement = 172
const sym_variable_statement = 173
const sym_variable_declaration = 174
const sym_variable_qualifier = 175
const sym_variable_identifier_declaration = 176
const sym_increment_statement = 177
const sym_decrement_statement = 178
const sym__expression = 179
const sym_const_literal = 180
const sym_float_literal = 181
const sym_bool_literal = 182
const sym_parenthesized_expression = 183
const sym_type_constructor_or_function_call_expression = 184
const sym_type_declaration = 185
const sym__vec_prefix = 186
const sym__mat_prefix = 187
const sym_texel_format = 188
const sym_address_space = 189
const sym_access_mode = 190
const sym_argument_list_expression = 191
const sym_bitcast_expression = 192
const sym_binary_expression = 193
const sym_unary_expression = 194
const sym_postfix_expression = 195
const sym_subscript_expression = 196
const sym_lhs_expression = 197
const sym_composite_value_decomposition_expression = 198
const aux_sym_source_file_repeat1 = 199
const aux_sym_source_file_repeat2 = 200
const aux_sym_global_variable_declaration_repeat1 = 201
const aux_sym_const_expresssion_repeat1 = 202
const aux_sym_struct_declaration_repeat1 = 203
const aux_sym_attribute_repeat1 = 204
const aux_sym_parameter_list_repeat1 = 205
const aux_sym_compound_statement_repeat1 = 206
const aux_sym_switch_statement_repeat1 = 207
const aux_sym_case_selectors_repeat1 = 208
const aux_sym_argument_list_expression_repeat1 = 209
const aux_sym_lhs_expression_repeat1 = 210

var ts_symbol_names = [211]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 15,
	3:   __ccgo_ts + 28,
	4:   __ccgo_ts + 30,
	5:   __ccgo_ts + 32,
	6:   __ccgo_ts + 36,
	7:   __ccgo_ts + 45,
	8:   __ccgo_ts + 47,
	9:   __ccgo_ts + 49,
	10:  __ccgo_ts + 51,
	11:  __ccgo_ts + 56,
	12:  __ccgo_ts + 59,
	13:  __ccgo_ts + 62,
	14:  __ccgo_ts + 69,
	15:  __ccgo_ts + 71,
	16:  __ccgo_ts + 73,
	17:  __ccgo_ts + 80,
	18:  __ccgo_ts + 82,
	19:  __ccgo_ts + 84,
	20:  __ccgo_ts + 87,
	21:  __ccgo_ts + 90,
	22:  __ccgo_ts + 93,
	23:  __ccgo_ts + 96,
	24:  __ccgo_ts + 99,
	25:  __ccgo_ts + 102,
	26:  __ccgo_ts + 105,
	27:  __ccgo_ts + 108,
	28:  __ccgo_ts + 111,
	29:  __ccgo_ts + 116,
	30:  __ccgo_ts + 123,
	31:  __ccgo_ts + 128,
	32:  __ccgo_ts + 130,
	33:  __ccgo_ts + 138,
	34:  __ccgo_ts + 150,
	35:  __ccgo_ts + 155,
	36:  __ccgo_ts + 159,
	37:  __ccgo_ts + 165,
	38:  __ccgo_ts + 171,
	39:  __ccgo_ts + 180,
	40:  __ccgo_ts + 191,
	41:  __ccgo_ts + 198,
	42:  __ccgo_ts + 206,
	43:  __ccgo_ts + 210,
	44:  __ccgo_ts + 212,
	45:  __ccgo_ts + 214,
	46:  __ccgo_ts + 217,
	47:  __ccgo_ts + 220,
	48:  __ccgo_ts + 232,
	49:  __ccgo_ts + 253,
	50:  __ccgo_ts + 274,
	51:  __ccgo_ts + 279,
	52:  __ccgo_ts + 285,
	53:  __ccgo_ts + 290,
	54:  __ccgo_ts + 294,
	55:  __ccgo_ts + 298,
	56:  __ccgo_ts + 302,
	57:  __ccgo_ts + 306,
	58:  __ccgo_ts + 312,
	59:  __ccgo_ts + 316,
	60:  __ccgo_ts + 324,
	61:  __ccgo_ts + 343,
	62:  __ccgo_ts + 360,
	63:  __ccgo_ts + 383,
	64:  __ccgo_ts + 402,
	65:  __ccgo_ts + 427,
	66:  __ccgo_ts + 457,
	67:  __ccgo_ts + 468,
	68:  __ccgo_ts + 479,
	69:  __ccgo_ts + 496,
	70:  __ccgo_ts + 507,
	71:  __ccgo_ts + 520,
	72:  __ccgo_ts + 539,
	73:  __ccgo_ts + 563,
	74:  __ccgo_ts + 582,
	75:  __ccgo_ts + 601,
	76:  __ccgo_ts + 626,
	77:  __ccgo_ts + 645,
	78:  __ccgo_ts + 650,
	79:  __ccgo_ts + 655,
	80:  __ccgo_ts + 660,
	81:  __ccgo_ts + 667,
	82:  __ccgo_ts + 674,
	83:  __ccgo_ts + 681,
	84:  __ccgo_ts + 688,
	85:  __ccgo_ts + 695,
	86:  __ccgo_ts + 702,
	87:  __ccgo_ts + 709,
	88:  __ccgo_ts + 716,
	89:  __ccgo_ts + 723,
	90:  __ccgo_ts + 734,
	91:  __ccgo_ts + 745,
	92:  __ccgo_ts + 755,
	93:  __ccgo_ts + 765,
	94:  __ccgo_ts + 776,
	95:  __ccgo_ts + 787,
	96:  __ccgo_ts + 799,
	97:  __ccgo_ts + 807,
	98:  __ccgo_ts + 815,
	99:  __ccgo_ts + 824,
	100: __ccgo_ts + 833,
	101: __ccgo_ts + 842,
	102: __ccgo_ts + 852,
	103: __ccgo_ts + 863,
	104: __ccgo_ts + 874,
	105: __ccgo_ts + 886,
	106: __ccgo_ts + 895,
	107: __ccgo_ts + 903,
	108: __ccgo_ts + 913,
	109: __ccgo_ts + 921,
	110: __ccgo_ts + 929,
	111: __ccgo_ts + 934,
	112: __ccgo_ts + 940,
	113: __ccgo_ts + 951,
	114: __ccgo_ts + 959,
	115: __ccgo_ts + 962,
	116: __ccgo_ts + 965,
	117: __ccgo_ts + 967,
	118: __ccgo_ts + 969,
	119: __ccgo_ts + 971,
	120: __ccgo_ts + 974,
	121: __ccgo_ts + 977,
	122: __ccgo_ts + 980,
	123: __ccgo_ts + 983,
	124: __ccgo_ts + 986,
	125: __ccgo_ts + 989,
	126: __ccgo_ts + 991,
	127: __ccgo_ts + 993,
	128: __ccgo_ts + 995,
	129: __ccgo_ts + 997,
	130: __ccgo_ts + 999,
	131: __ccgo_ts + 1001,
	132: __ccgo_ts + 1003,
	133: __ccgo_ts + 1005,
	134: __ccgo_ts + 1007,
	135: __ccgo_ts + 1009,
	136: __ccgo_ts + 1023,
	137: __ccgo_ts + 1035,
	138: __ccgo_ts + 1048,
	139: __ccgo_ts + 1076,
	140: __ccgo_ts + 1104,
	141: __ccgo_ts + 1127,
	142: __ccgo_ts + 1144,
	143: __ccgo_ts + 1165,
	144: __ccgo_ts + 1198,
	145: __ccgo_ts + 1217,
	146: __ccgo_ts + 1231,
	147: __ccgo_ts + 1248,
	148: __ccgo_ts + 1258,
	149: __ccgo_ts + 1281,
	150: __ccgo_ts + 1296,
	151: __ccgo_ts + 1306,
	152: __ccgo_ts + 1317,
	153: __ccgo_ts + 1336,
	154: __ccgo_ts + 1357,
	155: __ccgo_ts + 1386,
	156: __ccgo_ts + 1399,
	157: __ccgo_ts + 1414,
	158: __ccgo_ts + 1431,
	159: __ccgo_ts + 1443,
	160: __ccgo_ts + 1458,
	161: __ccgo_ts + 1482,
	162: __ccgo_ts + 1504,
	163: __ccgo_ts + 1519,
	164: __ccgo_ts + 1533,
	165: __ccgo_ts + 1544,
	166: __ccgo_ts + 1560,
	167: __ccgo_ts + 1576,
	168: __ccgo_ts + 1595,
	169: __ccgo_ts + 1614,
	170: __ccgo_ts + 1635,
	171: __ccgo_ts + 1665,
	172: __ccgo_ts + 1682,
	173: __ccgo_ts + 1700,
	174: __ccgo_ts + 1719,
	175: __ccgo_ts + 1740,
	176: __ccgo_ts + 1759,
	177: __ccgo_ts + 1791,
	178: __ccgo_ts + 1811,
	179: __ccgo_ts + 1831,
	180: __ccgo_ts + 1843,
	181: __ccgo_ts + 1857,
	182: __ccgo_ts + 1871,
	183: __ccgo_ts + 1884,
	184: __ccgo_ts + 1909,
	185: __ccgo_ts + 1954,
	186: __ccgo_ts + 1971,
	187: __ccgo_ts + 1983,
	188: __ccgo_ts + 1995,
	189: __ccgo_ts + 2008,
	190: __ccgo_ts + 2022,
	191: __ccgo_ts + 2034,
	192: __ccgo_ts + 2059,
	193: __ccgo_ts + 2078,
	194: __ccgo_ts + 2096,
	195: __ccgo_ts + 2113,
	196: __ccgo_ts + 2132,
	197: __ccgo_ts + 2153,
	198: __ccgo_ts + 2168,
	199: __ccgo_ts + 2209,
	200: __ccgo_ts + 2229,
	201: __ccgo_ts + 2249,
	202: __ccgo_ts + 2285,
	203: __ccgo_ts + 2311,
	204: __ccgo_ts + 2338,
	205: __ccgo_ts + 2356,
	206: __ccgo_ts + 2379,
	207: __ccgo_ts + 2406,
	208: __ccgo_ts + 2431,
	209: __ccgo_ts + 2454,
	210: __ccgo_ts + 2487,
}

var ts_symbol_map = [211]TSSymbol{
	1:   uint16(sym_identifier),
	2:   uint16(sym_line_comment),
	3:   uint16(anon_sym_SEMI),
	4:   uint16(anon_sym_EQ),
	5:   uint16(anon_sym_let),
	6:   uint16(anon_sym_override),
	7:   uint16(anon_sym_LPAREN),
	8:   uint16(anon_sym_COMMA),
	9:   uint16(anon_sym_RPAREN),
	10:  uint16(anon_sym_type),
	11:  uint16(anon_sym_fn),
	12:  uint16(anon_sym_DASH_GT),
	13:  uint16(anon_sym_struct),
	14:  uint16(anon_sym_LBRACE),
	15:  uint16(anon_sym_RBRACE),
	16:  uint16(anon_sym_enable),
	17:  uint16(anon_sym_AT),
	18:  uint16(anon_sym__),
	19:  uint16(anon_sym_PLUS_EQ),
	20:  uint16(anon_sym_DASH_EQ),
	21:  uint16(anon_sym_STAR_EQ),
	22:  uint16(anon_sym_SLASH_EQ),
	23:  uint16(anon_sym_PERCENT_EQ),
	24:  uint16(anon_sym_AMP_EQ),
	25:  uint16(anon_sym_PIPE_EQ),
	26:  uint16(anon_sym_CARET_EQ),
	27:  uint16(anon_sym_if),
	28:  uint16(anon_sym_else),
	29:  uint16(anon_sym_switch),
	30:  uint16(anon_sym_case),
	31:  uint16(anon_sym_COLON),
	32:  uint16(anon_sym_default),
	33:  uint16(anon_sym_fallthrough),
	34:  uint16(anon_sym_loop),
	35:  uint16(anon_sym_for),
	36:  uint16(anon_sym_while),
	37:  uint16(anon_sym_break),
	38:  uint16(anon_sym_continue),
	39:  uint16(anon_sym_continuing),
	40:  uint16(anon_sym_return),
	41:  uint16(anon_sym_discard),
	42:  uint16(anon_sym_var),
	43:  uint16(anon_sym_LT),
	44:  uint16(anon_sym_GT),
	45:  uint16(anon_sym_PLUS_PLUS),
	46:  uint16(anon_sym_DASH_DASH),
	47:  uint16(sym_int_literal),
	48:  uint16(aux_sym_float_literal_token1),
	49:  uint16(aux_sym_float_literal_token2),
	50:  uint16(anon_sym_true),
	51:  uint16(anon_sym_false),
	52:  uint16(anon_sym_bool),
	53:  uint16(anon_sym_u32),
	54:  uint16(anon_sym_i32),
	55:  uint16(anon_sym_f32),
	56:  uint16(anon_sym_f16),
	57:  uint16(anon_sym_array),
	58:  uint16(anon_sym_ptr),
	59:  uint16(anon_sym_sampler),
	60:  uint16(anon_sym_sampler_comparison),
	61:  uint16(anon_sym_texture_depth_2d),
	62:  uint16(anon_sym_texture_depth_2d_array),
	63:  uint16(anon_sym_texture_depth_cube),
	64:  uint16(anon_sym_texture_depth_cube_array),
	65:  uint16(anon_sym_texture_depth_multisampled_2d),
	66:  uint16(anon_sym_texture_1d),
	67:  uint16(anon_sym_texture_2d),
	68:  uint16(anon_sym_texture_2d_array),
	69:  uint16(anon_sym_texture_3d),
	70:  uint16(anon_sym_texture_cube),
	71:  uint16(anon_sym_texture_cube_array),
	72:  uint16(anon_sym_texture_multisampled_2d),
	73:  uint16(anon_sym_texture_storage_1d),
	74:  uint16(anon_sym_texture_storage_2d),
	75:  uint16(anon_sym_texture_storage_2d_array),
	76:  uint16(anon_sym_texture_storage_3d),
	77:  uint16(anon_sym_vec2),
	78:  uint16(anon_sym_vec3),
	79:  uint16(anon_sym_vec4),
	80:  uint16(anon_sym_mat2x2),
	81:  uint16(anon_sym_mat2x3),
	82:  uint16(anon_sym_mat2x4),
	83:  uint16(anon_sym_mat3x2),
	84:  uint16(anon_sym_mat3x3),
	85:  uint16(anon_sym_mat3x4),
	86:  uint16(anon_sym_mat4x2),
	87:  uint16(anon_sym_mat4x3),
	88:  uint16(anon_sym_mat4x4),
	89:  uint16(anon_sym_rgba8unorm),
	90:  uint16(anon_sym_rgba8snorm),
	91:  uint16(anon_sym_rgba8uint),
	92:  uint16(anon_sym_rgba8sint),
	93:  uint16(anon_sym_rgba16uint),
	94:  uint16(anon_sym_rgba16sint),
	95:  uint16(anon_sym_rgba16float),
	96:  uint16(anon_sym_r32uint),
	97:  uint16(anon_sym_r32sint),
	98:  uint16(anon_sym_r32float),
	99:  uint16(anon_sym_rg32uint),
	100: uint16(anon_sym_rg32sint),
	101: uint16(anon_sym_rg32float),
	102: uint16(anon_sym_rgba32uint),
	103: uint16(anon_sym_rgba32sint),
	104: uint16(anon_sym_rgba32float),
	105: uint16(anon_sym_function),
	106: uint16(anon_sym_private),
	107: uint16(anon_sym_workgroup),
	108: uint16(anon_sym_uniform),
	109: uint16(anon_sym_storage),
	110: uint16(anon_sym_read),
	111: uint16(anon_sym_write),
	112: uint16(anon_sym_read_write),
	113: uint16(anon_sym_bitcast),
	114: uint16(anon_sym_PIPE_PIPE),
	115: uint16(anon_sym_AMP_AMP),
	116: uint16(anon_sym_PIPE),
	117: uint16(anon_sym_CARET),
	118: uint16(anon_sym_AMP),
	119: uint16(anon_sym_EQ_EQ),
	120: uint16(anon_sym_BANG_EQ),
	121: uint16(anon_sym_LT_EQ),
	122: uint16(anon_sym_GT_EQ),
	123: uint16(anon_sym_LT_LT),
	124: uint16(anon_sym_GT_GT),
	125: uint16(anon_sym_PLUS),
	126: uint16(anon_sym_DASH),
	127: uint16(anon_sym_STAR),
	128: uint16(anon_sym_SLASH),
	129: uint16(anon_sym_PERCENT),
	130: uint16(anon_sym_BANG),
	131: uint16(anon_sym_TILDE),
	132: uint16(anon_sym_LBRACK),
	133: uint16(anon_sym_RBRACK),
	134: uint16(anon_sym_DOT),
	135: uint16(sym_block_comment),
	136: uint16(sym_source_file),
	137: uint16(sym__declaration),
	138: uint16(sym_global_variable_declaration),
	139: uint16(sym_global_constant_declaration),
	140: uint16(sym_type_alias_declaration),
	141: uint16(sym_const_expression),
	142: uint16(sym_function_declaration),
	143: uint16(sym_function_return_type_declaration),
	144: uint16(sym_struct_declaration),
	145: uint16(sym_struct_member),
	146: uint16(sym_enable_directive),
	147: uint16(sym_attribute),
	148: uint16(sym__literal_or_identifier),
	149: uint16(sym_parameter_list),
	150: uint16(sym_parameter),
	151: uint16(sym__statement),
	152: uint16(sym_compound_statement),
	153: uint16(sym_assignment_statement),
	154: uint16(sym_compound_assignment_operator),
	155: uint16(sym_if_statement),
	156: uint16(sym_else_statement),
	157: uint16(sym_switch_statement),
	158: uint16(sym_switch_body),
	159: uint16(sym_case_selectors),
	160: uint16(sym_case_compound_statement),
	161: uint16(sym_fallthrough_statement),
	162: uint16(sym_loop_statement),
	163: uint16(sym_for_statement),
	164: uint16(sym_for_header),
	165: uint16(sym_while_statement),
	166: uint16(sym_break_statement),
	167: uint16(sym_break_if_statement),
	168: uint16(sym_continue_statement),
	169: uint16(sym_continuing_statement),
	170: uint16(sym_continuing_compound_statement),
	171: uint16(sym_return_statement),
	172: uint16(sym_discard_statement),
	173: uint16(sym_variable_statement),
	174: uint16(sym_variable_declaration),
	175: uint16(sym_variable_qualifier),
	176: uint16(sym_variable_identifier_declaration),
	177: uint16(sym_increment_statement),
	178: uint16(sym_decrement_statement),
	179: uint16(sym__expression),
	180: uint16(sym_const_literal),
	181: uint16(sym_float_literal),
	182: uint16(sym_bool_literal),
	183: uint16(sym_parenthesized_expression),
	184: uint16(sym_type_constructor_or_function_call_expression),
	185: uint16(sym_type_declaration),
	186: uint16(sym__vec_prefix),
	187: uint16(sym__mat_prefix),
	188: uint16(sym_texel_format),
	189: uint16(sym_address_space),
	190: uint16(sym_access_mode),
	191: uint16(sym_argument_list_expression),
	192: uint16(sym_bitcast_expression),
	193: uint16(sym_binary_expression),
	194: uint16(sym_unary_expression),
	195: uint16(sym_postfix_expression),
	196: uint16(sym_subscript_expression),
	197: uint16(sym_lhs_expression),
	198: uint16(sym_composite_value_decomposition_expression),
	199: uint16(aux_sym_source_file_repeat1),
	200: uint16(aux_sym_source_file_repeat2),
	201: uint16(aux_sym_global_variable_declaration_repeat1),
	202: uint16(aux_sym_const_expresssion_repeat1),
	203: uint16(aux_sym_struct_declaration_repeat1),
	204: uint16(aux_sym_attribute_repeat1),
	205: uint16(aux_sym_parameter_list_repeat1),
	206: uint16(aux_sym_compound_statement_repeat1),
	207: uint16(aux_sym_switch_statement_repeat1),
	208: uint16(aux_sym_case_selectors_repeat1),
	209: uint16(aux_sym_argument_list_expression_repeat1),
	210: uint16(aux_sym_lhs_expression_repeat1),
}

var ts_symbol_metadata = [211]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	48: {},
	49: {},
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
	57: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
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
	},
	77: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
	82: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	83: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	84: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	85: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	86: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	87: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	88: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	89: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	90: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	91: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	92: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	93: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	94: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	95: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	96: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	97: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	98: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	99: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	100: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	101: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	102: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	103: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	104: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	105: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	106: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	107: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	108: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	109: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	110: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	111: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	112: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	113: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	114: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	115: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	116: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	117: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	118: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	119: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	120: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	121: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	122: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	123: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	124: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	125: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	126: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	127: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	128: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	129: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	130: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	131: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	132: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	133: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	134: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	141: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	142: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	143: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	144: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	145: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	149: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	150: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	151: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	152: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	153: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	154: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	155: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	156: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	157: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	158: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	159: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	160: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	161: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	162: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
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
	167: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	168: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	169: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	170: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	171: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	172: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	173: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	174: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	175: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	176: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	177: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	178: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	179: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	180: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	181: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	182: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	183: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	184: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	185: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	186: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	187: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	188: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	189: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	190: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	191: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	192: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	193: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	194: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	195: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	196: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	197: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	198: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	199: {},
	200: {},
	201: {},
	202: {},
	203: {},
	204: {},
	205: {},
	206: {},
	207: {},
	208: {},
	209: {},
	210: {},
}

const field_accessor = 1
const field_alternative = 2
const field_argument = 3
const field_body = 4
const field_condition = 5
const field_consequence = 6
const field_left = 7
const field_name = 8
const field_parameters = 9
const field_right = 10
const field_subscript = 11
const field_type = 12
const field_value = 13

var ts_field_names = [14]uintptr{
	0:  libc.UintptrFromInt32(0),
	1:  __ccgo_ts + 2510,
	2:  __ccgo_ts + 2519,
	3:  __ccgo_ts + 2531,
	4:  __ccgo_ts + 2540,
	5:  __ccgo_ts + 2545,
	6:  __ccgo_ts + 2555,
	7:  __ccgo_ts + 2567,
	8:  __ccgo_ts + 2572,
	9:  __ccgo_ts + 2577,
	10: __ccgo_ts + 2588,
	11: __ccgo_ts + 2594,
	12: __ccgo_ts + 51,
	13: __ccgo_ts + 2604,
}

var ts_field_map_slices = [18]TSFieldMapSlice{
	1: {
		Flength: uint16(2),
	},
	2: {
		Findex:  uint16(2),
		Flength: uint16(1),
	},
	3: {
		Findex:  uint16(3),
		Flength: uint16(2),
	},
	4: {
		Findex:  uint16(5),
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(6),
		Flength: uint16(2),
	},
	6: {
		Findex:  uint16(8),
		Flength: uint16(2),
	},
	7: {
		Findex:  uint16(10),
		Flength: uint16(3),
	},
	8: {
		Findex:  uint16(13),
		Flength: uint16(3),
	},
	9: {
		Findex:  uint16(16),
		Flength: uint16(2),
	},
	10: {
		Findex:  uint16(18),
		Flength: uint16(2),
	},
	11: {
		Findex:  uint16(20),
		Flength: uint16(4),
	},
	12: {
		Findex:  uint16(24),
		Flength: uint16(3),
	},
	13: {
		Findex:  uint16(27),
		Flength: uint16(3),
	},
	14: {
		Findex:  uint16(30),
		Flength: uint16(2),
	},
	15: {
		Findex:  uint16(32),
		Flength: uint16(1),
	},
	16: {
		Findex:  uint16(33),
		Flength: uint16(4),
	},
	17: {
		Findex:  uint16(37),
		Flength: uint16(3),
	},
}

var ts_field_map_entries = [40]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_name),
	},
	1: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(2),
	},
	2: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(1),
	},
	3: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(4),
	},
	4: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	5: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	6: {
		Ffield_id: uint16(field_left),
	},
	7: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	8: {
		Ffield_id:    uint16(field_accessor),
		Fchild_index: uint8(2),
	},
	9: {
		Ffield_id: uint16(field_value),
	},
	10: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(5),
	},
	11: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	12: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(4),
	},
	13: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(5),
	},
	14: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	15: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(3),
	},
	16: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(5),
	},
	17: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	18: {
		Ffield_id:    uint16(field_subscript),
		Fchild_index: uint8(2),
	},
	19: {
		Ffield_id: uint16(field_value),
	},
	20: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(6),
	},
	21: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	22: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(3),
	},
	23: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(5),
	},
	24: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(6),
	},
	25: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	26: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(5),
	},
	27: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(6),
	},
	28: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	29: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(4),
	},
	30: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	31: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(2),
	},
	32: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	33: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(7),
	},
	34: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	35: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(4),
	},
	36: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(6),
	},
	37: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(4),
	},
	38: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	39: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(2),
	},
}

var ts_alias_sequences = [18][8]TSSymbol{}

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
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(98)
			goto next_state
		}
		if lookahead == int32('%') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(92)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(87)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(94)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(100)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(101)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(99)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') || lookahead == int32(160) || lookahead == int32(8203) || lookahead == int32(8288) || lookahead == int32(65279) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(59)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('!') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('%') {
			state = uint16(95)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(78)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(91)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(102)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(93)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(100)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(101)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(75)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') || lookahead == int32(160) || lookahead == int32(8203) || lookahead == int32(8288) || lookahead == int32(65279) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('%') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('+') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(102)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(62)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(100)
			goto next_state
		}
		if lookahead == int32('^') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(26)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') || lookahead == int32(160) || lookahead == int32(8203) || lookahead == int32(8288) || lookahead == int32(65279) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(2)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(63)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('+') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('-') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(36)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('.') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(18)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('.') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(8)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('.') {
			state = uint16(66)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('.') {
			state = uint16(24)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('/') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('/') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('=') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('=') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('=') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('=') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('=') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('=') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('=') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(22)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(23)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(21):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(22):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(23):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(24):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(25):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(26):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(27):
		if eof != 0 {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(77)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(91)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(90)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(53)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(99)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') || lookahead == int32(160) || lookahead == int32(8203) || lookahead == int32(8288) || lookahead == int32(65279) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(27)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(59)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_line_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym__)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('<') {
			state = uint16(84)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(83)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(57)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(57)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(57)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(57)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(57)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_int_literal)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') || lookahead == int32('u') {
			state = uint16(57)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_literal_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_literal_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(18)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_literal_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(65)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_literal_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_literal_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(68)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_float_literal_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(19)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('|') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('&') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('&') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('+') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(36)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(6)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TILDE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(66)
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
		if lookahead == int32('a') {
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('v') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') || lookahead == int32(160) || lookahead == int32(8203) || lookahead == int32(8288) || lookahead == int32(65279) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('r') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('i') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('a') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('e') {
			state = uint16(24)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('l') {
			state = uint16(26)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('1') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('3') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('e') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(37)
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
		if lookahead == int32('v') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('r') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('3') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('g') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('a') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('e') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('y') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('3') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('a') {
			state = uint16(53)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('h') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('r') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('t') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('o') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('e') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('s') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('n') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('f') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('s') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('s') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('a') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('6') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('2') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('l') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_fn)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		if lookahead == int32('r') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('n') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('2') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_if)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		if lookahead == int32('t') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('o') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('t') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('e') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('i') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('r') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('2') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('a') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('3') {
			state = uint16(83)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('m') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('o') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('i') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('x') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('u') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('p') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('2') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('i') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('r') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('c') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('i') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('r') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('i') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('a') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('c') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('l') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('a') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('e') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('t') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('a') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('c') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('e') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('b') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_f16)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_f32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		if lookahead == int32('l') {
			state = uint16(109)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_for)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		if lookahead == int32('c') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_i32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_let)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		if lookahead == int32('p') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('2') {
			state = uint16(113)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(114)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('r') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('v') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ptr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		if lookahead == int32('f') {
			state = uint16(118)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(119)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead == int32('d') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead == int32('u') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('2') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('a') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('p') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('r') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead == int32('u') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead == int32('t') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead == int32('t') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('e') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('e') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_u32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(93):
		if lookahead == int32('f') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_var)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		if lookahead == int32('2') {
			state = uint16(133)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead == int32('l') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead == int32('k') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead == int32('t') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead == int32('y') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead == int32('a') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bool)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(102):
		if lookahead == int32('k') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_case)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(104):
		if lookahead == int32('i') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead == int32('u') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead == int32('a') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_else)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(108):
		if lookahead == int32('l') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('t') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead == int32('e') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead == int32('t') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_loop)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(113):
		if lookahead == int32('x') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('x') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('x') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead == int32('r') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead == int32('a') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead == int32('l') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('i') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('i') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_read)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('r') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead == int32('f') {
			state = uint16(159)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(160)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(124):
		if lookahead == int32('1') {
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead == int32('l') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead == int32('a') {
			state = uint16(166)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead == int32('c') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead == int32('c') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(129):
		if lookahead == int32('u') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(132):
		if lookahead == int32('o') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_vec2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_vec3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_vec4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(136):
		if lookahead == int32('e') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(137):
		if lookahead == int32('g') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(138):
		if lookahead == int32('e') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(139):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_array)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(140):
		if lookahead == int32('s') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(141):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_break)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(142):
		if lookahead == int32('n') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(143):
		if lookahead == int32('l') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(144):
		if lookahead == int32('r') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(145):
		if lookahead == int32('e') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(146):
		if lookahead == int32('h') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(147):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(148):
		if lookahead == int32('i') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(149):
		if lookahead == int32('2') {
			state = uint16(181)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(182)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(150):
		if lookahead == int32('2') {
			state = uint16(184)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(185)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(151):
		if lookahead == int32('2') {
			state = uint16(187)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(188)
			goto next_state
		}
		if lookahead == int32('4') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(152):
		if lookahead == int32('i') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(153):
		if lookahead == int32('t') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(154):
		if lookahead == int32('o') {
			state = uint16(192)
			goto next_state
		}
		return result
	case int32(155):
		if lookahead == int32('n') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(156):
		if lookahead == int32('n') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(157):
		if lookahead == int32('w') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(158):
		if lookahead == int32('n') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(159):
		if lookahead == int32('l') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(160):
		if lookahead == int32('i') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(161):
		if lookahead == int32('i') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(162):
		if lookahead == int32('6') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(163):
		if lookahead == int32('2') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(164):
		if lookahead == int32('s') {
			state = uint16(202)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(165):
		if lookahead == int32('e') {
			state = uint16(204)
			goto next_state
		}
		return result
	case int32(166):
		if lookahead == int32('g') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(167):
		if lookahead == int32('t') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(168):
		if lookahead == int32('h') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(169):
		if lookahead == int32('r') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(170):
		if lookahead == int32('r') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(171):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_while)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(172):
		if lookahead == int32('r') {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(173):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_write)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(174):
		if lookahead == int32('t') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(175):
		if lookahead == int32('u') {
			state = uint16(212)
			goto next_state
		}
		return result
	case int32(176):
		if lookahead == int32('t') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(177):
		if lookahead == int32('d') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(178):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enable)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(179):
		if lookahead == int32('r') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(180):
		if lookahead == int32('o') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(181):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat2x2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(182):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat2x3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(183):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat2x4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(184):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat3x2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(185):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat3x3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(186):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat3x4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(187):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat4x2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(188):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat4x3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(189):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mat4x4)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(190):
		if lookahead == int32('d') {
			state = uint16(217)
			goto next_state
		}
		return result
	case int32(191):
		if lookahead == int32('e') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(192):
		if lookahead == int32('a') {
			state = uint16(219)
			goto next_state
		}
		return result
	case int32(193):
		if lookahead == int32('t') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(194):
		if lookahead == int32('t') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(195):
		if lookahead == int32('r') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(196):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_return)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(197):
		if lookahead == int32('o') {
			state = uint16(223)
			goto next_state
		}
		return result
	case int32(198):
		if lookahead == int32('n') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(199):
		if lookahead == int32('n') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(200):
		if lookahead == int32('f') {
			state = uint16(226)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(227)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(201):
		if lookahead == int32('f') {
			state = uint16(229)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(230)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(202):
		if lookahead == int32('i') {
			state = uint16(232)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(233)
			goto next_state
		}
		return result
	case int32(203):
		if lookahead == int32('i') {
			state = uint16(234)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(235)
			goto next_state
		}
		return result
	case int32(204):
		if lookahead == int32('r') {
			state = uint16(236)
			goto next_state
		}
		return result
	case int32(205):
		if lookahead == int32('e') {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(206):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_struct)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(207):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_switch)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(208):
		if lookahead == int32('e') {
			state = uint16(238)
			goto next_state
		}
		return result
	case int32(209):
		if lookahead == int32('m') {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(210):
		if lookahead == int32('o') {
			state = uint16(240)
			goto next_state
		}
		return result
	case int32(211):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bitcast)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(212):
		if lookahead == int32('e') {
			state = uint16(241)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(213):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_default)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(214):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_discard)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(215):
		if lookahead == int32('o') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(216):
		if lookahead == int32('n') {
			state = uint16(244)
			goto next_state
		}
		return result
	case int32(217):
		if lookahead == int32('e') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(218):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_private)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(219):
		if lookahead == int32('t') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(220):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_r32sint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(221):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_r32uint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(222):
		if lookahead == int32('i') {
			state = uint16(247)
			goto next_state
		}
		return result
	case int32(223):
		if lookahead == int32('a') {
			state = uint16(248)
			goto next_state
		}
		return result
	case int32(224):
		if lookahead == int32('t') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(225):
		if lookahead == int32('t') {
			state = uint16(250)
			goto next_state
		}
		return result
	case int32(226):
		if lookahead == int32('l') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(227):
		if lookahead == int32('i') {
			state = uint16(252)
			goto next_state
		}
		return result
	case int32(228):
		if lookahead == int32('i') {
			state = uint16(253)
			goto next_state
		}
		return result
	case int32(229):
		if lookahead == int32('l') {
			state = uint16(254)
			goto next_state
		}
		return result
	case int32(230):
		if lookahead == int32('i') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(231):
		if lookahead == int32('i') {
			state = uint16(256)
			goto next_state
		}
		return result
	case int32(232):
		if lookahead == int32('n') {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(233):
		if lookahead == int32('o') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(234):
		if lookahead == int32('n') {
			state = uint16(259)
			goto next_state
		}
		return result
	case int32(235):
		if lookahead == int32('o') {
			state = uint16(260)
			goto next_state
		}
		return result
	case int32(236):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sampler)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(237):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_storage)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(238):
		if lookahead == int32('_') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(239):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_uniform)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(240):
		if lookahead == int32('u') {
			state = uint16(263)
			goto next_state
		}
		return result
	case int32(241):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_continue)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(242):
		if lookahead == int32('n') {
			state = uint16(264)
			goto next_state
		}
		return result
	case int32(243):
		if lookahead == int32('u') {
			state = uint16(265)
			goto next_state
		}
		return result
	case int32(244):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_function)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(245):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_override)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(246):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_r32float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(247):
		if lookahead == int32('t') {
			state = uint16(266)
			goto next_state
		}
		return result
	case int32(248):
		if lookahead == int32('t') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(249):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rg32sint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(250):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rg32uint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(251):
		if lookahead == int32('o') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(252):
		if lookahead == int32('n') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(253):
		if lookahead == int32('n') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(254):
		if lookahead == int32('o') {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(255):
		if lookahead == int32('n') {
			state = uint16(272)
			goto next_state
		}
		return result
	case int32(256):
		if lookahead == int32('n') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(257):
		if lookahead == int32('t') {
			state = uint16(274)
			goto next_state
		}
		return result
	case int32(258):
		if lookahead == int32('r') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(259):
		if lookahead == int32('t') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(260):
		if lookahead == int32('r') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(261):
		if lookahead == int32('c') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(262):
		if lookahead == int32('1') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(280)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(282)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(283)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(284)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(285)
			goto next_state
		}
		return result
	case int32(263):
		if lookahead == int32('p') {
			state = uint16(286)
			goto next_state
		}
		return result
	case int32(264):
		if lookahead == int32('g') {
			state = uint16(287)
			goto next_state
		}
		return result
	case int32(265):
		if lookahead == int32('g') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(266):
		if lookahead == int32('e') {
			state = uint16(289)
			goto next_state
		}
		return result
	case int32(267):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rg32float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(268):
		if lookahead == int32('a') {
			state = uint16(290)
			goto next_state
		}
		return result
	case int32(269):
		if lookahead == int32('t') {
			state = uint16(291)
			goto next_state
		}
		return result
	case int32(270):
		if lookahead == int32('t') {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(271):
		if lookahead == int32('a') {
			state = uint16(293)
			goto next_state
		}
		return result
	case int32(272):
		if lookahead == int32('t') {
			state = uint16(294)
			goto next_state
		}
		return result
	case int32(273):
		if lookahead == int32('t') {
			state = uint16(295)
			goto next_state
		}
		return result
	case int32(274):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba8sint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(275):
		if lookahead == int32('m') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(276):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba8uint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(277):
		if lookahead == int32('m') {
			state = uint16(297)
			goto next_state
		}
		return result
	case int32(278):
		if lookahead == int32('o') {
			state = uint16(298)
			goto next_state
		}
		return result
	case int32(279):
		if lookahead == int32('d') {
			state = uint16(299)
			goto next_state
		}
		return result
	case int32(280):
		if lookahead == int32('d') {
			state = uint16(300)
			goto next_state
		}
		return result
	case int32(281):
		if lookahead == int32('d') {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(282):
		if lookahead == int32('u') {
			state = uint16(302)
			goto next_state
		}
		return result
	case int32(283):
		if lookahead == int32('e') {
			state = uint16(303)
			goto next_state
		}
		return result
	case int32(284):
		if lookahead == int32('u') {
			state = uint16(304)
			goto next_state
		}
		return result
	case int32(285):
		if lookahead == int32('t') {
			state = uint16(305)
			goto next_state
		}
		return result
	case int32(286):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_workgroup)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(287):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_continuing)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(288):
		if lookahead == int32('h') {
			state = uint16(306)
			goto next_state
		}
		return result
	case int32(289):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_read_write)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(290):
		if lookahead == int32('t') {
			state = uint16(307)
			goto next_state
		}
		return result
	case int32(291):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba16sint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(292):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba16uint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(293):
		if lookahead == int32('t') {
			state = uint16(308)
			goto next_state
		}
		return result
	case int32(294):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba32sint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(295):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba32uint)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(296):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba8snorm)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(297):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba8unorm)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(298):
		if lookahead == int32('m') {
			state = uint16(309)
			goto next_state
		}
		return result
	case int32(299):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_1d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(300):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_2d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(310)
			goto next_state
		}
		return result
	case int32(301):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_3d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(302):
		if lookahead == int32('b') {
			state = uint16(311)
			goto next_state
		}
		return result
	case int32(303):
		if lookahead == int32('p') {
			state = uint16(312)
			goto next_state
		}
		return result
	case int32(304):
		if lookahead == int32('l') {
			state = uint16(313)
			goto next_state
		}
		return result
	case int32(305):
		if lookahead == int32('o') {
			state = uint16(314)
			goto next_state
		}
		return result
	case int32(306):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_fallthrough)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(307):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba16float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(308):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rgba32float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(309):
		if lookahead == int32('p') {
			state = uint16(315)
			goto next_state
		}
		return result
	case int32(310):
		if lookahead == int32('a') {
			state = uint16(316)
			goto next_state
		}
		return result
	case int32(311):
		if lookahead == int32('e') {
			state = uint16(317)
			goto next_state
		}
		return result
	case int32(312):
		if lookahead == int32('t') {
			state = uint16(318)
			goto next_state
		}
		return result
	case int32(313):
		if lookahead == int32('t') {
			state = uint16(319)
			goto next_state
		}
		return result
	case int32(314):
		if lookahead == int32('r') {
			state = uint16(320)
			goto next_state
		}
		return result
	case int32(315):
		if lookahead == int32('a') {
			state = uint16(321)
			goto next_state
		}
		return result
	case int32(316):
		if lookahead == int32('r') {
			state = uint16(322)
			goto next_state
		}
		return result
	case int32(317):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_cube)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(323)
			goto next_state
		}
		return result
	case int32(318):
		if lookahead == int32('h') {
			state = uint16(324)
			goto next_state
		}
		return result
	case int32(319):
		if lookahead == int32('i') {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(320):
		if lookahead == int32('a') {
			state = uint16(326)
			goto next_state
		}
		return result
	case int32(321):
		if lookahead == int32('r') {
			state = uint16(327)
			goto next_state
		}
		return result
	case int32(322):
		if lookahead == int32('r') {
			state = uint16(328)
			goto next_state
		}
		return result
	case int32(323):
		if lookahead == int32('a') {
			state = uint16(329)
			goto next_state
		}
		return result
	case int32(324):
		if lookahead == int32('_') {
			state = uint16(330)
			goto next_state
		}
		return result
	case int32(325):
		if lookahead == int32('s') {
			state = uint16(331)
			goto next_state
		}
		return result
	case int32(326):
		if lookahead == int32('g') {
			state = uint16(332)
			goto next_state
		}
		return result
	case int32(327):
		if lookahead == int32('i') {
			state = uint16(333)
			goto next_state
		}
		return result
	case int32(328):
		if lookahead == int32('a') {
			state = uint16(334)
			goto next_state
		}
		return result
	case int32(329):
		if lookahead == int32('r') {
			state = uint16(335)
			goto next_state
		}
		return result
	case int32(330):
		if lookahead == int32('2') {
			state = uint16(336)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(337)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(338)
			goto next_state
		}
		return result
	case int32(331):
		if lookahead == int32('a') {
			state = uint16(339)
			goto next_state
		}
		return result
	case int32(332):
		if lookahead == int32('e') {
			state = uint16(340)
			goto next_state
		}
		return result
	case int32(333):
		if lookahead == int32('s') {
			state = uint16(341)
			goto next_state
		}
		return result
	case int32(334):
		if lookahead == int32('y') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(335):
		if lookahead == int32('r') {
			state = uint16(343)
			goto next_state
		}
		return result
	case int32(336):
		if lookahead == int32('d') {
			state = uint16(344)
			goto next_state
		}
		return result
	case int32(337):
		if lookahead == int32('u') {
			state = uint16(345)
			goto next_state
		}
		return result
	case int32(338):
		if lookahead == int32('u') {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(339):
		if lookahead == int32('m') {
			state = uint16(347)
			goto next_state
		}
		return result
	case int32(340):
		if lookahead == int32('_') {
			state = uint16(348)
			goto next_state
		}
		return result
	case int32(341):
		if lookahead == int32('o') {
			state = uint16(349)
			goto next_state
		}
		return result
	case int32(342):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_2d_array)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(343):
		if lookahead == int32('a') {
			state = uint16(350)
			goto next_state
		}
		return result
	case int32(344):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_depth_2d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(351)
			goto next_state
		}
		return result
	case int32(345):
		if lookahead == int32('b') {
			state = uint16(352)
			goto next_state
		}
		return result
	case int32(346):
		if lookahead == int32('l') {
			state = uint16(353)
			goto next_state
		}
		return result
	case int32(347):
		if lookahead == int32('p') {
			state = uint16(354)
			goto next_state
		}
		return result
	case int32(348):
		if lookahead == int32('1') {
			state = uint16(355)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(356)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(357)
			goto next_state
		}
		return result
	case int32(349):
		if lookahead == int32('n') {
			state = uint16(358)
			goto next_state
		}
		return result
	case int32(350):
		if lookahead == int32('y') {
			state = uint16(359)
			goto next_state
		}
		return result
	case int32(351):
		if lookahead == int32('a') {
			state = uint16(360)
			goto next_state
		}
		return result
	case int32(352):
		if lookahead == int32('e') {
			state = uint16(361)
			goto next_state
		}
		return result
	case int32(353):
		if lookahead == int32('t') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(354):
		if lookahead == int32('l') {
			state = uint16(363)
			goto next_state
		}
		return result
	case int32(355):
		if lookahead == int32('d') {
			state = uint16(364)
			goto next_state
		}
		return result
	case int32(356):
		if lookahead == int32('d') {
			state = uint16(365)
			goto next_state
		}
		return result
	case int32(357):
		if lookahead == int32('d') {
			state = uint16(366)
			goto next_state
		}
		return result
	case int32(358):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_sampler_comparison)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(359):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_cube_array)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(360):
		if lookahead == int32('r') {
			state = uint16(367)
			goto next_state
		}
		return result
	case int32(361):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_depth_cube)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(368)
			goto next_state
		}
		return result
	case int32(362):
		if lookahead == int32('i') {
			state = uint16(369)
			goto next_state
		}
		return result
	case int32(363):
		if lookahead == int32('e') {
			state = uint16(370)
			goto next_state
		}
		return result
	case int32(364):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_storage_1d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(365):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_storage_2d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(366):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_storage_3d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(367):
		if lookahead == int32('r') {
			state = uint16(372)
			goto next_state
		}
		return result
	case int32(368):
		if lookahead == int32('a') {
			state = uint16(373)
			goto next_state
		}
		return result
	case int32(369):
		if lookahead == int32('s') {
			state = uint16(374)
			goto next_state
		}
		return result
	case int32(370):
		if lookahead == int32('d') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(371):
		if lookahead == int32('a') {
			state = uint16(376)
			goto next_state
		}
		return result
	case int32(372):
		if lookahead == int32('a') {
			state = uint16(377)
			goto next_state
		}
		return result
	case int32(373):
		if lookahead == int32('r') {
			state = uint16(378)
			goto next_state
		}
		return result
	case int32(374):
		if lookahead == int32('a') {
			state = uint16(379)
			goto next_state
		}
		return result
	case int32(375):
		if lookahead == int32('_') {
			state = uint16(380)
			goto next_state
		}
		return result
	case int32(376):
		if lookahead == int32('r') {
			state = uint16(381)
			goto next_state
		}
		return result
	case int32(377):
		if lookahead == int32('y') {
			state = uint16(382)
			goto next_state
		}
		return result
	case int32(378):
		if lookahead == int32('r') {
			state = uint16(383)
			goto next_state
		}
		return result
	case int32(379):
		if lookahead == int32('m') {
			state = uint16(384)
			goto next_state
		}
		return result
	case int32(380):
		if lookahead == int32('2') {
			state = uint16(385)
			goto next_state
		}
		return result
	case int32(381):
		if lookahead == int32('r') {
			state = uint16(386)
			goto next_state
		}
		return result
	case int32(382):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_depth_2d_array)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(383):
		if lookahead == int32('a') {
			state = uint16(387)
			goto next_state
		}
		return result
	case int32(384):
		if lookahead == int32('p') {
			state = uint16(388)
			goto next_state
		}
		return result
	case int32(385):
		if lookahead == int32('d') {
			state = uint16(389)
			goto next_state
		}
		return result
	case int32(386):
		if lookahead == int32('a') {
			state = uint16(390)
			goto next_state
		}
		return result
	case int32(387):
		if lookahead == int32('y') {
			state = uint16(391)
			goto next_state
		}
		return result
	case int32(388):
		if lookahead == int32('l') {
			state = uint16(392)
			goto next_state
		}
		return result
	case int32(389):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_multisampled_2d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(390):
		if lookahead == int32('y') {
			state = uint16(393)
			goto next_state
		}
		return result
	case int32(391):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_depth_cube_array)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(392):
		if lookahead == int32('e') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(393):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_storage_2d_array)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(394):
		if lookahead == int32('d') {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(395):
		if lookahead == int32('_') {
			state = uint16(396)
			goto next_state
		}
		return result
	case int32(396):
		if lookahead == int32('2') {
			state = uint16(397)
			goto next_state
		}
		return result
	case int32(397):
		if lookahead == int32('d') {
			state = uint16(398)
			goto next_state
		}
		return result
	case int32(398):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_texture_depth_multisampled_2d)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [360]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Fexternal_lex_state: uint16(1),
	},
	2: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	3: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	4: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	5: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	6: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	7: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	8: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	9: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	10: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	11: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	12: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	13: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	14: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	15: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	16: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	17: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	18: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	19: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	20: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	21: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	22: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	23: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	24: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	25: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	26: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	27: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	28: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	29: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	30: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	31: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	32: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	33: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	34: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	35: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	36: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	37: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	38: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	39: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	40: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	41: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	42: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	43: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	44: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	45: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	46: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	47: {
		Fexternal_lex_state: uint16(1),
	},
	48: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	49: {
		Fexternal_lex_state: uint16(1),
	},
	50: {
		Fexternal_lex_state: uint16(1),
	},
	51: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	52: {
		Fexternal_lex_state: uint16(1),
	},
	53: {
		Fexternal_lex_state: uint16(1),
	},
	54: {
		Fexternal_lex_state: uint16(1),
	},
	55: {
		Fexternal_lex_state: uint16(1),
	},
	56: {
		Fexternal_lex_state: uint16(1),
	},
	57: {
		Fexternal_lex_state: uint16(1),
	},
	58: {
		Fexternal_lex_state: uint16(1),
	},
	59: {
		Fexternal_lex_state: uint16(1),
	},
	60: {
		Fexternal_lex_state: uint16(1),
	},
	61: {
		Fexternal_lex_state: uint16(1),
	},
	62: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	63: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	64: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	65: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	66: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	67: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	68: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	69: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	70: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	71: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	72: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	73: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	74: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	75: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	76: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	77: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	78: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	79: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	80: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	81: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	82: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	83: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	84: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	85: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	86: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	87: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	88: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	89: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	90: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	91: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	92: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	93: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	94: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	95: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	96: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	97: {
		Fexternal_lex_state: uint16(1),
	},
	98: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	99: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	100: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	101: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	102: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	103: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	104: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	105: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	106: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	107: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	108: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	109: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	110: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	111: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	112: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	113: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	114: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	115: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	116: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	117: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	118: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	119: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	120: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	121: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	122: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	123: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	124: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	125: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	126: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	127: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	128: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	129: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	130: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	131: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	132: {
		Fexternal_lex_state: uint16(1),
	},
	133: {
		Fexternal_lex_state: uint16(1),
	},
	134: {
		Fexternal_lex_state: uint16(1),
	},
	135: {
		Fexternal_lex_state: uint16(1),
	},
	136: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	137: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	138: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	139: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	140: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	141: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	142: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	143: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	144: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	145: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	146: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	147: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	148: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	149: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	150: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	151: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	152: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	153: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	154: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	155: {
		Fexternal_lex_state: uint16(1),
	},
	156: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	157: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	158: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	159: {
		Fexternal_lex_state: uint16(1),
	},
	160: {
		Fexternal_lex_state: uint16(1),
	},
	161: {
		Fexternal_lex_state: uint16(1),
	},
	162: {
		Fexternal_lex_state: uint16(1),
	},
	163: {
		Fexternal_lex_state: uint16(1),
	},
	164: {
		Fexternal_lex_state: uint16(1),
	},
	165: {
		Fexternal_lex_state: uint16(1),
	},
	166: {
		Fexternal_lex_state: uint16(1),
	},
	167: {
		Fexternal_lex_state: uint16(1),
	},
	168: {
		Fexternal_lex_state: uint16(1),
	},
	169: {
		Fexternal_lex_state: uint16(1),
	},
	170: {
		Fexternal_lex_state: uint16(1),
	},
	171: {
		Fexternal_lex_state: uint16(1),
	},
	172: {
		Fexternal_lex_state: uint16(1),
	},
	173: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	174: {
		Fexternal_lex_state: uint16(1),
	},
	175: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	176: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	177: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	178: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	179: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	180: {
		Fexternal_lex_state: uint16(1),
	},
	181: {
		Fexternal_lex_state: uint16(1),
	},
	182: {
		Fexternal_lex_state: uint16(1),
	},
	183: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	184: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	185: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	186: {
		Fexternal_lex_state: uint16(1),
	},
	187: {
		Fexternal_lex_state: uint16(1),
	},
	188: {
		Fexternal_lex_state: uint16(1),
	},
	189: {
		Fexternal_lex_state: uint16(1),
	},
	190: {
		Fexternal_lex_state: uint16(1),
	},
	191: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	192: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	193: {
		Fexternal_lex_state: uint16(1),
	},
	194: {
		Fexternal_lex_state: uint16(1),
	},
	195: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	196: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	197: {
		Fexternal_lex_state: uint16(1),
	},
	198: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	199: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	200: {
		Fexternal_lex_state: uint16(1),
	},
	201: {
		Fexternal_lex_state: uint16(1),
	},
	202: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	203: {
		Fexternal_lex_state: uint16(1),
	},
	204: {
		Fexternal_lex_state: uint16(1),
	},
	205: {
		Fexternal_lex_state: uint16(1),
	},
	206: {
		Fexternal_lex_state: uint16(1),
	},
	207: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	208: {
		Fexternal_lex_state: uint16(1),
	},
	209: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	210: {
		Fexternal_lex_state: uint16(1),
	},
	211: {
		Fexternal_lex_state: uint16(1),
	},
	212: {
		Fexternal_lex_state: uint16(1),
	},
	213: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	214: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	215: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	216: {
		Fexternal_lex_state: uint16(1),
	},
	217: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	218: {
		Fexternal_lex_state: uint16(1),
	},
	219: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	220: {
		Fexternal_lex_state: uint16(1),
	},
	221: {
		Fexternal_lex_state: uint16(1),
	},
	222: {
		Fexternal_lex_state: uint16(1),
	},
	223: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	224: {
		Fexternal_lex_state: uint16(1),
	},
	225: {
		Fexternal_lex_state: uint16(1),
	},
	226: {
		Fexternal_lex_state: uint16(1),
	},
	227: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	228: {
		Fexternal_lex_state: uint16(1),
	},
	229: {
		Fexternal_lex_state: uint16(1),
	},
	230: {
		Fexternal_lex_state: uint16(1),
	},
	231: {
		Fexternal_lex_state: uint16(1),
	},
	232: {
		Fexternal_lex_state: uint16(1),
	},
	233: {
		Fexternal_lex_state: uint16(1),
	},
	234: {
		Fexternal_lex_state: uint16(1),
	},
	235: {
		Fexternal_lex_state: uint16(1),
	},
	236: {
		Fexternal_lex_state: uint16(1),
	},
	237: {
		Fexternal_lex_state: uint16(1),
	},
	238: {
		Fexternal_lex_state: uint16(1),
	},
	239: {
		Fexternal_lex_state: uint16(1),
	},
	240: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	241: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	242: {
		Fexternal_lex_state: uint16(1),
	},
	243: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	244: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	245: {
		Fexternal_lex_state: uint16(1),
	},
	246: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	247: {
		Fexternal_lex_state: uint16(1),
	},
	248: {
		Fexternal_lex_state: uint16(1),
	},
	249: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	250: {
		Fexternal_lex_state: uint16(1),
	},
	251: {
		Fexternal_lex_state: uint16(1),
	},
	252: {
		Fexternal_lex_state: uint16(1),
	},
	253: {
		Fexternal_lex_state: uint16(1),
	},
	254: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	255: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	256: {
		Fexternal_lex_state: uint16(1),
	},
	257: {
		Fexternal_lex_state: uint16(1),
	},
	258: {
		Fexternal_lex_state: uint16(1),
	},
	259: {
		Fexternal_lex_state: uint16(1),
	},
	260: {
		Fexternal_lex_state: uint16(1),
	},
	261: {
		Fexternal_lex_state: uint16(1),
	},
	262: {
		Fexternal_lex_state: uint16(1),
	},
	263: {
		Fexternal_lex_state: uint16(1),
	},
	264: {
		Fexternal_lex_state: uint16(1),
	},
	265: {
		Fexternal_lex_state: uint16(1),
	},
	266: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	267: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(1),
	},
	268: {
		Fexternal_lex_state: uint16(1),
	},
	269: {
		Fexternal_lex_state: uint16(1),
	},
	270: {
		Fexternal_lex_state: uint16(1),
	},
	271: {
		Fexternal_lex_state: uint16(1),
	},
	272: {
		Fexternal_lex_state: uint16(1),
	},
	273: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	274: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	275: {
		Fexternal_lex_state: uint16(1),
	},
	276: {
		Fexternal_lex_state: uint16(1),
	},
	277: {
		Fexternal_lex_state: uint16(1),
	},
	278: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	279: {
		Fexternal_lex_state: uint16(1),
	},
	280: {
		Fexternal_lex_state: uint16(1),
	},
	281: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	282: {
		Fexternal_lex_state: uint16(1),
	},
	283: {
		Fexternal_lex_state: uint16(1),
	},
	284: {
		Fexternal_lex_state: uint16(1),
	},
	285: {
		Fexternal_lex_state: uint16(1),
	},
	286: {
		Fexternal_lex_state: uint16(1),
	},
	287: {
		Fexternal_lex_state: uint16(1),
	},
	288: {
		Fexternal_lex_state: uint16(1),
	},
	289: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	290: {
		Fexternal_lex_state: uint16(1),
	},
	291: {
		Fexternal_lex_state: uint16(1),
	},
	292: {
		Fexternal_lex_state: uint16(1),
	},
	293: {
		Fexternal_lex_state: uint16(1),
	},
	294: {
		Fexternal_lex_state: uint16(1),
	},
	295: {
		Fexternal_lex_state: uint16(1),
	},
	296: {
		Fexternal_lex_state: uint16(1),
	},
	297: {
		Fexternal_lex_state: uint16(1),
	},
	298: {
		Fexternal_lex_state: uint16(1),
	},
	299: {
		Fexternal_lex_state: uint16(1),
	},
	300: {
		Fexternal_lex_state: uint16(1),
	},
	301: {
		Fexternal_lex_state: uint16(1),
	},
	302: {
		Fexternal_lex_state: uint16(1),
	},
	303: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	304: {
		Fexternal_lex_state: uint16(1),
	},
	305: {
		Fexternal_lex_state: uint16(1),
	},
	306: {
		Fexternal_lex_state: uint16(1),
	},
	307: {
		Fexternal_lex_state: uint16(1),
	},
	308: {
		Fexternal_lex_state: uint16(1),
	},
	309: {
		Fexternal_lex_state: uint16(1),
	},
	310: {
		Fexternal_lex_state: uint16(1),
	},
	311: {
		Fexternal_lex_state: uint16(1),
	},
	312: {
		Fexternal_lex_state: uint16(1),
	},
	313: {
		Fexternal_lex_state: uint16(1),
	},
	314: {
		Fexternal_lex_state: uint16(1),
	},
	315: {
		Fexternal_lex_state: uint16(1),
	},
	316: {
		Fexternal_lex_state: uint16(1),
	},
	317: {
		Fexternal_lex_state: uint16(1),
	},
	318: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	319: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	320: {
		Fexternal_lex_state: uint16(1),
	},
	321: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	322: {
		Fexternal_lex_state: uint16(1),
	},
	323: {
		Fexternal_lex_state: uint16(1),
	},
	324: {
		Fexternal_lex_state: uint16(1),
	},
	325: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	326: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	327: {
		Fexternal_lex_state: uint16(1),
	},
	328: {
		Fexternal_lex_state: uint16(1),
	},
	329: {
		Fexternal_lex_state: uint16(1),
	},
	330: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	331: {
		Fexternal_lex_state: uint16(1),
	},
	332: {
		Fexternal_lex_state: uint16(1),
	},
	333: {
		Fexternal_lex_state: uint16(1),
	},
	334: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
	335: {
		Fexternal_lex_state: uint16(1),
	},
	336: {
		Fexternal_lex_state: uint16(1),
	},
	337: {
		Fexternal_lex_state: uint16(1),
	},
	338: {
		Fexternal_lex_state: uint16(1),
	},
	339: {
		Fexternal_lex_state: uint16(1),
	},
	340: {
		Fexternal_lex_state: uint16(1),
	},
	341: {
		Fexternal_lex_state: uint16(1),
	},
	342: {
		Fexternal_lex_state: uint16(1),
	},
	343: {
		Fexternal_lex_state: uint16(1),
	},
	344: {
		Fexternal_lex_state: uint16(1),
	},
	345: {
		Fexternal_lex_state: uint16(1),
	},
	346: {
		Fexternal_lex_state: uint16(1),
	},
	347: {
		Fexternal_lex_state: uint16(1),
	},
	348: {
		Fexternal_lex_state: uint16(1),
	},
	349: {
		Fexternal_lex_state: uint16(1),
	},
	350: {
		Fexternal_lex_state: uint16(1),
	},
	351: {
		Fexternal_lex_state: uint16(1),
	},
	352: {
		Fexternal_lex_state: uint16(1),
	},
	353: {
		Fexternal_lex_state: uint16(1),
	},
	354: {
		Fexternal_lex_state: uint16(1),
	},
	355: {
		Fexternal_lex_state: uint16(1),
	},
	356: {
		Fexternal_lex_state: uint16(1),
	},
	357: {
		Fexternal_lex_state: uint16(1),
	},
	358: {
		Fexternal_lex_state: uint16(1),
	},
	359: {
		Flex_state:          uint16(27),
		Fexternal_lex_state: uint16(1),
	},
}

const ts_external_token_block_comment = 0

var ts_external_scanner_symbol_map = [1]TSSymbol{
	0: uint16(sym_block_comment),
}

var ts_external_scanner_states = [2][1]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
	},
}

var ts_parse_table = [32][211]uint16_t{
	0: {
		0:   uint16(1),
		1:   uint16(1),
		2:   uint16(3),
		3:   uint16(1),
		4:   uint16(1),
		5:   uint16(1),
		6:   uint16(1),
		7:   uint16(1),
		8:   uint16(1),
		9:   uint16(1),
		10:  uint16(1),
		11:  uint16(1),
		12:  uint16(1),
		13:  uint16(1),
		14:  uint16(1),
		15:  uint16(1),
		16:  uint16(1),
		17:  uint16(1),
		18:  uint16(1),
		19:  uint16(1),
		20:  uint16(1),
		21:  uint16(1),
		22:  uint16(1),
		23:  uint16(1),
		24:  uint16(1),
		25:  uint16(1),
		26:  uint16(1),
		27:  uint16(1),
		28:  uint16(1),
		29:  uint16(1),
		30:  uint16(1),
		31:  uint16(1),
		32:  uint16(1),
		33:  uint16(1),
		34:  uint16(1),
		35:  uint16(1),
		36:  uint16(1),
		37:  uint16(1),
		38:  uint16(1),
		39:  uint16(1),
		40:  uint16(1),
		41:  uint16(1),
		42:  uint16(1),
		43:  uint16(1),
		44:  uint16(1),
		45:  uint16(1),
		47:  uint16(1),
		48:  uint16(1),
		49:  uint16(1),
		50:  uint16(1),
		51:  uint16(1),
		52:  uint16(1),
		53:  uint16(1),
		54:  uint16(1),
		55:  uint16(1),
		56:  uint16(1),
		57:  uint16(1),
		58:  uint16(1),
		59:  uint16(1),
		60:  uint16(1),
		61:  uint16(1),
		62:  uint16(1),
		63:  uint16(1),
		64:  uint16(1),
		65:  uint16(1),
		66:  uint16(1),
		67:  uint16(1),
		68:  uint16(1),
		69:  uint16(1),
		70:  uint16(1),
		71:  uint16(1),
		72:  uint16(1),
		73:  uint16(1),
		74:  uint16(1),
		75:  uint16(1),
		76:  uint16(1),
		77:  uint16(1),
		78:  uint16(1),
		79:  uint16(1),
		80:  uint16(1),
		81:  uint16(1),
		82:  uint16(1),
		83:  uint16(1),
		84:  uint16(1),
		85:  uint16(1),
		86:  uint16(1),
		87:  uint16(1),
		88:  uint16(1),
		89:  uint16(1),
		90:  uint16(1),
		91:  uint16(1),
		92:  uint16(1),
		93:  uint16(1),
		94:  uint16(1),
		95:  uint16(1),
		96:  uint16(1),
		97:  uint16(1),
		98:  uint16(1),
		99:  uint16(1),
		100: uint16(1),
		101: uint16(1),
		102: uint16(1),
		103: uint16(1),
		104: uint16(1),
		105: uint16(1),
		106: uint16(1),
		107: uint16(1),
		108: uint16(1),
		109: uint16(1),
		110: uint16(1),
		111: uint16(1),
		112: uint16(1),
		113: uint16(1),
		114: uint16(1),
		115: uint16(1),
		116: uint16(1),
		117: uint16(1),
		118: uint16(1),
		119: uint16(1),
		120: uint16(1),
		121: uint16(1),
		122: uint16(1),
		123: uint16(1),
		124: uint16(1),
		125: uint16(1),
		126: uint16(1),
		127: uint16(1),
		128: uint16(1),
		129: uint16(1),
		130: uint16(1),
		131: uint16(1),
		132: uint16(1),
		133: uint16(1),
		134: uint16(1),
		135: uint16(3),
	},
	1: {
		0:   uint16(5),
		2:   uint16(3),
		3:   uint16(7),
		5:   uint16(9),
		6:   uint16(11),
		10:  uint16(13),
		11:  uint16(15),
		13:  uint16(17),
		16:  uint16(19),
		17:  uint16(21),
		42:  uint16(23),
		135: uint16(3),
		136: uint16(347),
		137: uint16(133),
		138: uint16(342),
		139: uint16(342),
		140: uint16(342),
		142: uint16(133),
		144: uint16(133),
		146: uint16(97),
		147: uint16(182),
		174: uint16(243),
		199: uint16(97),
		200: uint16(133),
		201: uint16(182),
	},
	2: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		9:   uint16(29),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(104),
		180: uint16(104),
		181: uint16(75),
		182: uint16(75),
		183: uint16(104),
		184: uint16(104),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(104),
		193: uint16(104),
		194: uint16(104),
		196: uint16(104),
		198: uint16(104),
		209: uint16(6),
	},
	3: {
		1:   uint16(25),
		2:   uint16(3),
		3:   uint16(55),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(116),
		180: uint16(116),
		181: uint16(75),
		182: uint16(75),
		183: uint16(116),
		184: uint16(116),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(116),
		193: uint16(116),
		194: uint16(116),
		196: uint16(116),
		198: uint16(116),
	},
	4: {
		1:   uint16(57),
		2:   uint16(3),
		7:   uint16(60),
		47:  uint16(63),
		48:  uint16(66),
		49:  uint16(66),
		50:  uint16(69),
		51:  uint16(69),
		52:  uint16(72),
		53:  uint16(72),
		54:  uint16(72),
		55:  uint16(72),
		56:  uint16(72),
		57:  uint16(75),
		58:  uint16(78),
		59:  uint16(72),
		60:  uint16(72),
		61:  uint16(72),
		62:  uint16(72),
		63:  uint16(72),
		64:  uint16(72),
		65:  uint16(72),
		66:  uint16(81),
		67:  uint16(81),
		68:  uint16(81),
		69:  uint16(81),
		70:  uint16(81),
		71:  uint16(81),
		72:  uint16(81),
		73:  uint16(84),
		74:  uint16(84),
		75:  uint16(84),
		76:  uint16(84),
		77:  uint16(87),
		78:  uint16(87),
		79:  uint16(87),
		80:  uint16(87),
		81:  uint16(87),
		82:  uint16(87),
		83:  uint16(87),
		84:  uint16(87),
		85:  uint16(87),
		86:  uint16(87),
		87:  uint16(87),
		88:  uint16(87),
		113: uint16(90),
		118: uint16(93),
		126: uint16(96),
		127: uint16(93),
		130: uint16(93),
		131: uint16(93),
		135: uint16(3),
		179: uint16(106),
		180: uint16(106),
		181: uint16(75),
		182: uint16(75),
		183: uint16(106),
		184: uint16(106),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(106),
		193: uint16(106),
		194: uint16(106),
		196: uint16(106),
		198: uint16(106),
		209: uint16(4),
	},
	5: {
		1:   uint16(25),
		2:   uint16(3),
		3:   uint16(99),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(110),
		180: uint16(110),
		181: uint16(75),
		182: uint16(75),
		183: uint16(110),
		184: uint16(110),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(110),
		193: uint16(110),
		194: uint16(110),
		196: uint16(110),
		198: uint16(110),
	},
	6: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(99),
		180: uint16(99),
		181: uint16(75),
		182: uint16(75),
		183: uint16(99),
		184: uint16(99),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(99),
		193: uint16(99),
		194: uint16(99),
		196: uint16(99),
		198: uint16(99),
		209: uint16(4),
	},
	7: {
		1:   uint16(25),
		2:   uint16(3),
		3:   uint16(101),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(117),
		180: uint16(117),
		181: uint16(75),
		182: uint16(75),
		183: uint16(117),
		184: uint16(117),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(117),
		193: uint16(117),
		194: uint16(117),
		196: uint16(117),
		198: uint16(117),
	},
	8: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(113),
		180: uint16(113),
		181: uint16(75),
		182: uint16(75),
		183: uint16(113),
		184: uint16(113),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(113),
		193: uint16(113),
		194: uint16(113),
		196: uint16(113),
		198: uint16(113),
	},
	9: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(92),
		180: uint16(92),
		181: uint16(75),
		182: uint16(75),
		183: uint16(92),
		184: uint16(92),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(92),
		193: uint16(92),
		194: uint16(92),
		196: uint16(92),
		198: uint16(92),
	},
	10: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(98),
		180: uint16(98),
		181: uint16(75),
		182: uint16(75),
		183: uint16(98),
		184: uint16(98),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(98),
		193: uint16(98),
		194: uint16(98),
		196: uint16(98),
		198: uint16(98),
	},
	11: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(118),
		180: uint16(118),
		181: uint16(75),
		182: uint16(75),
		183: uint16(118),
		184: uint16(118),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(118),
		193: uint16(118),
		194: uint16(118),
		196: uint16(118),
		198: uint16(118),
	},
	12: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(87),
		180: uint16(87),
		181: uint16(75),
		182: uint16(75),
		183: uint16(87),
		184: uint16(87),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(87),
		193: uint16(87),
		194: uint16(87),
		196: uint16(87),
		198: uint16(87),
	},
	13: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(100),
		180: uint16(100),
		181: uint16(75),
		182: uint16(75),
		183: uint16(100),
		184: uint16(100),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(100),
		193: uint16(100),
		194: uint16(100),
		196: uint16(100),
		198: uint16(100),
	},
	14: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(93),
		180: uint16(93),
		181: uint16(75),
		182: uint16(75),
		183: uint16(93),
		184: uint16(93),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(93),
		193: uint16(93),
		194: uint16(93),
		196: uint16(93),
		198: uint16(93),
	},
	15: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(84),
		180: uint16(84),
		181: uint16(75),
		182: uint16(75),
		183: uint16(84),
		184: uint16(84),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(84),
		193: uint16(84),
		194: uint16(84),
		196: uint16(84),
		198: uint16(84),
	},
	16: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(81),
		180: uint16(81),
		181: uint16(75),
		182: uint16(75),
		183: uint16(81),
		184: uint16(81),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(81),
		193: uint16(81),
		194: uint16(81),
		196: uint16(81),
		198: uint16(81),
	},
	17: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(112),
		180: uint16(112),
		181: uint16(75),
		182: uint16(75),
		183: uint16(112),
		184: uint16(112),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(112),
		193: uint16(112),
		194: uint16(112),
		196: uint16(112),
		198: uint16(112),
	},
	18: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(107),
		180: uint16(107),
		181: uint16(75),
		182: uint16(75),
		183: uint16(107),
		184: uint16(107),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(107),
		193: uint16(107),
		194: uint16(107),
		196: uint16(107),
		198: uint16(107),
	},
	19: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(108),
		180: uint16(108),
		181: uint16(75),
		182: uint16(75),
		183: uint16(108),
		184: uint16(108),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(108),
		193: uint16(108),
		194: uint16(108),
		196: uint16(108),
		198: uint16(108),
	},
	20: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(109),
		180: uint16(109),
		181: uint16(75),
		182: uint16(75),
		183: uint16(109),
		184: uint16(109),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(109),
		193: uint16(109),
		194: uint16(109),
		196: uint16(109),
		198: uint16(109),
	},
	21: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(105),
		180: uint16(105),
		181: uint16(75),
		182: uint16(75),
		183: uint16(105),
		184: uint16(105),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(105),
		193: uint16(105),
		194: uint16(105),
		196: uint16(105),
		198: uint16(105),
	},
	22: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(79),
		180: uint16(79),
		181: uint16(75),
		182: uint16(75),
		183: uint16(79),
		184: uint16(79),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(79),
		193: uint16(79),
		194: uint16(79),
		196: uint16(79),
		198: uint16(79),
	},
	23: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(94),
		180: uint16(94),
		181: uint16(75),
		182: uint16(75),
		183: uint16(94),
		184: uint16(94),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(94),
		193: uint16(94),
		194: uint16(94),
		196: uint16(94),
		198: uint16(94),
	},
	24: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(88),
		180: uint16(88),
		181: uint16(75),
		182: uint16(75),
		183: uint16(88),
		184: uint16(88),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(88),
		193: uint16(88),
		194: uint16(88),
		196: uint16(88),
		198: uint16(88),
	},
	25: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(102),
		180: uint16(102),
		181: uint16(75),
		182: uint16(75),
		183: uint16(102),
		184: uint16(102),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(102),
		193: uint16(102),
		194: uint16(102),
		196: uint16(102),
		198: uint16(102),
	},
	26: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(119),
		180: uint16(119),
		181: uint16(75),
		182: uint16(75),
		183: uint16(119),
		184: uint16(119),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(119),
		193: uint16(119),
		194: uint16(119),
		196: uint16(119),
		198: uint16(119),
	},
	27: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(91),
		180: uint16(91),
		181: uint16(75),
		182: uint16(75),
		183: uint16(91),
		184: uint16(91),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(91),
		193: uint16(91),
		194: uint16(91),
		196: uint16(91),
		198: uint16(91),
	},
	28: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(114),
		180: uint16(114),
		181: uint16(75),
		182: uint16(75),
		183: uint16(114),
		184: uint16(114),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(114),
		193: uint16(114),
		194: uint16(114),
		196: uint16(114),
		198: uint16(114),
	},
	29: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(115),
		180: uint16(115),
		181: uint16(75),
		182: uint16(75),
		183: uint16(115),
		184: uint16(115),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(115),
		193: uint16(115),
		194: uint16(115),
		196: uint16(115),
		198: uint16(115),
	},
	30: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(78),
		180: uint16(78),
		181: uint16(75),
		182: uint16(75),
		183: uint16(78),
		184: uint16(78),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(78),
		193: uint16(78),
		194: uint16(78),
		196: uint16(78),
		198: uint16(78),
	},
	31: {
		1:   uint16(25),
		2:   uint16(3),
		7:   uint16(27),
		47:  uint16(31),
		48:  uint16(33),
		49:  uint16(33),
		50:  uint16(35),
		51:  uint16(35),
		52:  uint16(37),
		53:  uint16(37),
		54:  uint16(37),
		55:  uint16(37),
		56:  uint16(37),
		57:  uint16(39),
		58:  uint16(41),
		59:  uint16(37),
		60:  uint16(37),
		61:  uint16(37),
		62:  uint16(37),
		63:  uint16(37),
		64:  uint16(37),
		65:  uint16(37),
		66:  uint16(43),
		67:  uint16(43),
		68:  uint16(43),
		69:  uint16(43),
		70:  uint16(43),
		71:  uint16(43),
		72:  uint16(43),
		73:  uint16(45),
		74:  uint16(45),
		75:  uint16(45),
		76:  uint16(45),
		77:  uint16(47),
		78:  uint16(47),
		79:  uint16(47),
		80:  uint16(47),
		81:  uint16(47),
		82:  uint16(47),
		83:  uint16(47),
		84:  uint16(47),
		85:  uint16(47),
		86:  uint16(47),
		87:  uint16(47),
		88:  uint16(47),
		113: uint16(49),
		118: uint16(51),
		126: uint16(53),
		127: uint16(51),
		130: uint16(51),
		131: uint16(51),
		135: uint16(3),
		179: uint16(83),
		180: uint16(83),
		181: uint16(75),
		182: uint16(75),
		183: uint16(83),
		184: uint16(83),
		185: uint16(280),
		186: uint16(218),
		187: uint16(218),
		192: uint16(83),
		193: uint16(83),
		194: uint16(83),
		196: uint16(83),
		198: uint16(83),
	},
}

var ts_small_parse_table = [8615]uint16_t{
	0:    uint16(21),
	1:    uint16(39),
	2:    uint16(1),
	3:    uint16(anon_sym_array),
	4:    uint16(41),
	5:    uint16(1),
	6:    uint16(anon_sym_ptr),
	7:    uint16(103),
	8:    uint16(1),
	9:    uint16(sym_identifier),
	10:   uint16(105),
	11:   uint16(1),
	12:   uint16(anon_sym_SEMI),
	13:   uint16(107),
	14:   uint16(1),
	15:   uint16(anon_sym_let),
	16:   uint16(109),
	17:   uint16(1),
	18:   uint16(anon_sym_LPAREN),
	19:   uint16(111),
	20:   uint16(1),
	21:   uint16(anon_sym__),
	22:   uint16(113),
	23:   uint16(1),
	24:   uint16(anon_sym_var),
	25:   uint16(154),
	26:   uint16(1),
	27:   uint16(sym_lhs_expression),
	28:   uint16(202),
	29:   uint16(1),
	30:   uint16(aux_sym_lhs_expression_repeat1),
	31:   uint16(246),
	32:   uint16(1),
	33:   uint16(sym_variable_declaration),
	34:   uint16(280),
	35:   uint16(1),
	36:   uint16(sym_type_declaration),
	37:   uint16(307),
	38:   uint16(1),
	39:   uint16(sym_for_header),
	40:   uint16(3),
	41:   uint16(2),
	42:   uint16(sym_block_comment),
	43:   uint16(sym_line_comment),
	44:   uint16(115),
	45:   uint16(2),
	46:   uint16(anon_sym_AMP),
	47:   uint16(anon_sym_STAR),
	48:   uint16(218),
	49:   uint16(2),
	50:   uint16(sym__vec_prefix),
	51:   uint16(sym__mat_prefix),
	52:   uint16(45),
	53:   uint16(4),
	54:   uint16(anon_sym_texture_storage_1d),
	55:   uint16(anon_sym_texture_storage_2d),
	56:   uint16(anon_sym_texture_storage_2d_array),
	57:   uint16(anon_sym_texture_storage_3d),
	58:   uint16(308),
	59:   uint16(5),
	60:   uint16(sym_assignment_statement),
	61:   uint16(sym_variable_statement),
	62:   uint16(sym_increment_statement),
	63:   uint16(sym_decrement_statement),
	64:   uint16(sym_type_constructor_or_function_call_expression),
	65:   uint16(43),
	66:   uint16(7),
	67:   uint16(anon_sym_texture_1d),
	68:   uint16(anon_sym_texture_2d),
	69:   uint16(anon_sym_texture_2d_array),
	70:   uint16(anon_sym_texture_3d),
	71:   uint16(anon_sym_texture_cube),
	72:   uint16(anon_sym_texture_cube_array),
	73:   uint16(anon_sym_texture_multisampled_2d),
	74:   uint16(37),
	75:   uint16(12),
	76:   uint16(anon_sym_bool),
	77:   uint16(anon_sym_u32),
	78:   uint16(anon_sym_i32),
	79:   uint16(anon_sym_f32),
	80:   uint16(anon_sym_f16),
	81:   uint16(anon_sym_sampler),
	82:   uint16(anon_sym_sampler_comparison),
	83:   uint16(anon_sym_texture_depth_2d),
	84:   uint16(anon_sym_texture_depth_2d_array),
	85:   uint16(anon_sym_texture_depth_cube),
	86:   uint16(anon_sym_texture_depth_cube_array),
	87:   uint16(anon_sym_texture_depth_multisampled_2d),
	88:   uint16(47),
	89:   uint16(12),
	90:   uint16(anon_sym_vec2),
	91:   uint16(anon_sym_vec3),
	92:   uint16(anon_sym_vec4),
	93:   uint16(anon_sym_mat2x2),
	94:   uint16(anon_sym_mat2x3),
	95:   uint16(anon_sym_mat2x4),
	96:   uint16(anon_sym_mat3x2),
	97:   uint16(anon_sym_mat3x3),
	98:   uint16(anon_sym_mat3x4),
	99:   uint16(anon_sym_mat4x2),
	100:  uint16(anon_sym_mat4x3),
	101:  uint16(anon_sym_mat4x4),
	102:  uint16(17),
	103:  uint16(39),
	104:  uint16(1),
	105:  uint16(anon_sym_array),
	106:  uint16(41),
	107:  uint16(1),
	108:  uint16(anon_sym_ptr),
	109:  uint16(103),
	110:  uint16(1),
	111:  uint16(sym_identifier),
	112:  uint16(109),
	113:  uint16(1),
	114:  uint16(anon_sym_LPAREN),
	115:  uint16(111),
	116:  uint16(1),
	117:  uint16(anon_sym__),
	118:  uint16(117),
	119:  uint16(1),
	120:  uint16(anon_sym_RPAREN),
	121:  uint16(154),
	122:  uint16(1),
	123:  uint16(sym_lhs_expression),
	124:  uint16(202),
	125:  uint16(1),
	126:  uint16(aux_sym_lhs_expression_repeat1),
	127:  uint16(280),
	128:  uint16(1),
	129:  uint16(sym_type_declaration),
	130:  uint16(3),
	131:  uint16(2),
	132:  uint16(sym_block_comment),
	133:  uint16(sym_line_comment),
	134:  uint16(115),
	135:  uint16(2),
	136:  uint16(anon_sym_AMP),
	137:  uint16(anon_sym_STAR),
	138:  uint16(218),
	139:  uint16(2),
	140:  uint16(sym__vec_prefix),
	141:  uint16(sym__mat_prefix),
	142:  uint16(45),
	143:  uint16(4),
	144:  uint16(anon_sym_texture_storage_1d),
	145:  uint16(anon_sym_texture_storage_2d),
	146:  uint16(anon_sym_texture_storage_2d_array),
	147:  uint16(anon_sym_texture_storage_3d),
	148:  uint16(338),
	149:  uint16(4),
	150:  uint16(sym_assignment_statement),
	151:  uint16(sym_increment_statement),
	152:  uint16(sym_decrement_statement),
	153:  uint16(sym_type_constructor_or_function_call_expression),
	154:  uint16(43),
	155:  uint16(7),
	156:  uint16(anon_sym_texture_1d),
	157:  uint16(anon_sym_texture_2d),
	158:  uint16(anon_sym_texture_2d_array),
	159:  uint16(anon_sym_texture_3d),
	160:  uint16(anon_sym_texture_cube),
	161:  uint16(anon_sym_texture_cube_array),
	162:  uint16(anon_sym_texture_multisampled_2d),
	163:  uint16(37),
	164:  uint16(12),
	165:  uint16(anon_sym_bool),
	166:  uint16(anon_sym_u32),
	167:  uint16(anon_sym_i32),
	168:  uint16(anon_sym_f32),
	169:  uint16(anon_sym_f16),
	170:  uint16(anon_sym_sampler),
	171:  uint16(anon_sym_sampler_comparison),
	172:  uint16(anon_sym_texture_depth_2d),
	173:  uint16(anon_sym_texture_depth_2d_array),
	174:  uint16(anon_sym_texture_depth_cube),
	175:  uint16(anon_sym_texture_depth_cube_array),
	176:  uint16(anon_sym_texture_depth_multisampled_2d),
	177:  uint16(47),
	178:  uint16(12),
	179:  uint16(anon_sym_vec2),
	180:  uint16(anon_sym_vec3),
	181:  uint16(anon_sym_vec4),
	182:  uint16(anon_sym_mat2x2),
	183:  uint16(anon_sym_mat2x3),
	184:  uint16(anon_sym_mat2x4),
	185:  uint16(anon_sym_mat3x2),
	186:  uint16(anon_sym_mat3x3),
	187:  uint16(anon_sym_mat3x4),
	188:  uint16(anon_sym_mat4x2),
	189:  uint16(anon_sym_mat4x3),
	190:  uint16(anon_sym_mat4x4),
	191:  uint16(17),
	192:  uint16(31),
	193:  uint16(1),
	194:  uint16(sym_int_literal),
	195:  uint16(39),
	196:  uint16(1),
	197:  uint16(anon_sym_array),
	198:  uint16(41),
	199:  uint16(1),
	200:  uint16(anon_sym_ptr),
	201:  uint16(119),
	202:  uint16(1),
	203:  uint16(anon_sym_RPAREN),
	204:  uint16(39),
	205:  uint16(1),
	206:  uint16(aux_sym_const_expresssion_repeat1),
	207:  uint16(224),
	208:  uint16(1),
	209:  uint16(sym_const_literal),
	210:  uint16(245),
	211:  uint16(1),
	212:  uint16(sym_const_expression),
	213:  uint16(295),
	214:  uint16(1),
	215:  uint16(sym_type_declaration),
	216:  uint16(3),
	217:  uint16(2),
	218:  uint16(sym_block_comment),
	219:  uint16(sym_line_comment),
	220:  uint16(33),
	221:  uint16(2),
	222:  uint16(aux_sym_float_literal_token1),
	223:  uint16(aux_sym_float_literal_token2),
	224:  uint16(35),
	225:  uint16(2),
	226:  uint16(anon_sym_true),
	227:  uint16(anon_sym_false),
	228:  uint16(75),
	229:  uint16(2),
	230:  uint16(sym_float_literal),
	231:  uint16(sym_bool_literal),
	232:  uint16(293),
	233:  uint16(2),
	234:  uint16(sym__vec_prefix),
	235:  uint16(sym__mat_prefix),
	236:  uint16(45),
	237:  uint16(4),
	238:  uint16(anon_sym_texture_storage_1d),
	239:  uint16(anon_sym_texture_storage_2d),
	240:  uint16(anon_sym_texture_storage_2d_array),
	241:  uint16(anon_sym_texture_storage_3d),
	242:  uint16(43),
	243:  uint16(7),
	244:  uint16(anon_sym_texture_1d),
	245:  uint16(anon_sym_texture_2d),
	246:  uint16(anon_sym_texture_2d_array),
	247:  uint16(anon_sym_texture_3d),
	248:  uint16(anon_sym_texture_cube),
	249:  uint16(anon_sym_texture_cube_array),
	250:  uint16(anon_sym_texture_multisampled_2d),
	251:  uint16(121),
	252:  uint16(12),
	253:  uint16(anon_sym_vec2),
	254:  uint16(anon_sym_vec3),
	255:  uint16(anon_sym_vec4),
	256:  uint16(anon_sym_mat2x2),
	257:  uint16(anon_sym_mat2x3),
	258:  uint16(anon_sym_mat2x4),
	259:  uint16(anon_sym_mat3x2),
	260:  uint16(anon_sym_mat3x3),
	261:  uint16(anon_sym_mat3x4),
	262:  uint16(anon_sym_mat4x2),
	263:  uint16(anon_sym_mat4x3),
	264:  uint16(anon_sym_mat4x4),
	265:  uint16(37),
	266:  uint16(13),
	267:  uint16(sym_identifier),
	268:  uint16(anon_sym_bool),
	269:  uint16(anon_sym_u32),
	270:  uint16(anon_sym_i32),
	271:  uint16(anon_sym_f32),
	272:  uint16(anon_sym_f16),
	273:  uint16(anon_sym_sampler),
	274:  uint16(anon_sym_sampler_comparison),
	275:  uint16(anon_sym_texture_depth_2d),
	276:  uint16(anon_sym_texture_depth_2d_array),
	277:  uint16(anon_sym_texture_depth_cube),
	278:  uint16(anon_sym_texture_depth_cube_array),
	279:  uint16(anon_sym_texture_depth_multisampled_2d),
	280:  uint16(17),
	281:  uint16(39),
	282:  uint16(1),
	283:  uint16(anon_sym_array),
	284:  uint16(41),
	285:  uint16(1),
	286:  uint16(anon_sym_ptr),
	287:  uint16(103),
	288:  uint16(1),
	289:  uint16(sym_identifier),
	290:  uint16(109),
	291:  uint16(1),
	292:  uint16(anon_sym_LPAREN),
	293:  uint16(111),
	294:  uint16(1),
	295:  uint16(anon_sym__),
	296:  uint16(123),
	297:  uint16(1),
	298:  uint16(anon_sym_RPAREN),
	299:  uint16(154),
	300:  uint16(1),
	301:  uint16(sym_lhs_expression),
	302:  uint16(202),
	303:  uint16(1),
	304:  uint16(aux_sym_lhs_expression_repeat1),
	305:  uint16(280),
	306:  uint16(1),
	307:  uint16(sym_type_declaration),
	308:  uint16(3),
	309:  uint16(2),
	310:  uint16(sym_block_comment),
	311:  uint16(sym_line_comment),
	312:  uint16(115),
	313:  uint16(2),
	314:  uint16(anon_sym_AMP),
	315:  uint16(anon_sym_STAR),
	316:  uint16(218),
	317:  uint16(2),
	318:  uint16(sym__vec_prefix),
	319:  uint16(sym__mat_prefix),
	320:  uint16(45),
	321:  uint16(4),
	322:  uint16(anon_sym_texture_storage_1d),
	323:  uint16(anon_sym_texture_storage_2d),
	324:  uint16(anon_sym_texture_storage_2d_array),
	325:  uint16(anon_sym_texture_storage_3d),
	326:  uint16(349),
	327:  uint16(4),
	328:  uint16(sym_assignment_statement),
	329:  uint16(sym_increment_statement),
	330:  uint16(sym_decrement_statement),
	331:  uint16(sym_type_constructor_or_function_call_expression),
	332:  uint16(43),
	333:  uint16(7),
	334:  uint16(anon_sym_texture_1d),
	335:  uint16(anon_sym_texture_2d),
	336:  uint16(anon_sym_texture_2d_array),
	337:  uint16(anon_sym_texture_3d),
	338:  uint16(anon_sym_texture_cube),
	339:  uint16(anon_sym_texture_cube_array),
	340:  uint16(anon_sym_texture_multisampled_2d),
	341:  uint16(37),
	342:  uint16(12),
	343:  uint16(anon_sym_bool),
	344:  uint16(anon_sym_u32),
	345:  uint16(anon_sym_i32),
	346:  uint16(anon_sym_f32),
	347:  uint16(anon_sym_f16),
	348:  uint16(anon_sym_sampler),
	349:  uint16(anon_sym_sampler_comparison),
	350:  uint16(anon_sym_texture_depth_2d),
	351:  uint16(anon_sym_texture_depth_2d_array),
	352:  uint16(anon_sym_texture_depth_cube),
	353:  uint16(anon_sym_texture_depth_cube_array),
	354:  uint16(anon_sym_texture_depth_multisampled_2d),
	355:  uint16(47),
	356:  uint16(12),
	357:  uint16(anon_sym_vec2),
	358:  uint16(anon_sym_vec3),
	359:  uint16(anon_sym_vec4),
	360:  uint16(anon_sym_mat2x2),
	361:  uint16(anon_sym_mat2x3),
	362:  uint16(anon_sym_mat2x4),
	363:  uint16(anon_sym_mat3x2),
	364:  uint16(anon_sym_mat3x3),
	365:  uint16(anon_sym_mat3x4),
	366:  uint16(anon_sym_mat4x2),
	367:  uint16(anon_sym_mat4x3),
	368:  uint16(anon_sym_mat4x4),
	369:  uint16(17),
	370:  uint16(39),
	371:  uint16(1),
	372:  uint16(anon_sym_array),
	373:  uint16(41),
	374:  uint16(1),
	375:  uint16(anon_sym_ptr),
	376:  uint16(103),
	377:  uint16(1),
	378:  uint16(sym_identifier),
	379:  uint16(109),
	380:  uint16(1),
	381:  uint16(anon_sym_LPAREN),
	382:  uint16(111),
	383:  uint16(1),
	384:  uint16(anon_sym__),
	385:  uint16(125),
	386:  uint16(1),
	387:  uint16(anon_sym_RPAREN),
	388:  uint16(154),
	389:  uint16(1),
	390:  uint16(sym_lhs_expression),
	391:  uint16(202),
	392:  uint16(1),
	393:  uint16(aux_sym_lhs_expression_repeat1),
	394:  uint16(280),
	395:  uint16(1),
	396:  uint16(sym_type_declaration),
	397:  uint16(3),
	398:  uint16(2),
	399:  uint16(sym_block_comment),
	400:  uint16(sym_line_comment),
	401:  uint16(115),
	402:  uint16(2),
	403:  uint16(anon_sym_AMP),
	404:  uint16(anon_sym_STAR),
	405:  uint16(218),
	406:  uint16(2),
	407:  uint16(sym__vec_prefix),
	408:  uint16(sym__mat_prefix),
	409:  uint16(45),
	410:  uint16(4),
	411:  uint16(anon_sym_texture_storage_1d),
	412:  uint16(anon_sym_texture_storage_2d),
	413:  uint16(anon_sym_texture_storage_2d_array),
	414:  uint16(anon_sym_texture_storage_3d),
	415:  uint16(320),
	416:  uint16(4),
	417:  uint16(sym_assignment_statement),
	418:  uint16(sym_increment_statement),
	419:  uint16(sym_decrement_statement),
	420:  uint16(sym_type_constructor_or_function_call_expression),
	421:  uint16(43),
	422:  uint16(7),
	423:  uint16(anon_sym_texture_1d),
	424:  uint16(anon_sym_texture_2d),
	425:  uint16(anon_sym_texture_2d_array),
	426:  uint16(anon_sym_texture_3d),
	427:  uint16(anon_sym_texture_cube),
	428:  uint16(anon_sym_texture_cube_array),
	429:  uint16(anon_sym_texture_multisampled_2d),
	430:  uint16(37),
	431:  uint16(12),
	432:  uint16(anon_sym_bool),
	433:  uint16(anon_sym_u32),
	434:  uint16(anon_sym_i32),
	435:  uint16(anon_sym_f32),
	436:  uint16(anon_sym_f16),
	437:  uint16(anon_sym_sampler),
	438:  uint16(anon_sym_sampler_comparison),
	439:  uint16(anon_sym_texture_depth_2d),
	440:  uint16(anon_sym_texture_depth_2d_array),
	441:  uint16(anon_sym_texture_depth_cube),
	442:  uint16(anon_sym_texture_depth_cube_array),
	443:  uint16(anon_sym_texture_depth_multisampled_2d),
	444:  uint16(47),
	445:  uint16(12),
	446:  uint16(anon_sym_vec2),
	447:  uint16(anon_sym_vec3),
	448:  uint16(anon_sym_vec4),
	449:  uint16(anon_sym_mat2x2),
	450:  uint16(anon_sym_mat2x3),
	451:  uint16(anon_sym_mat2x4),
	452:  uint16(anon_sym_mat3x2),
	453:  uint16(anon_sym_mat3x3),
	454:  uint16(anon_sym_mat3x4),
	455:  uint16(anon_sym_mat4x2),
	456:  uint16(anon_sym_mat4x3),
	457:  uint16(anon_sym_mat4x4),
	458:  uint16(4),
	459:  uint16(131),
	460:  uint16(1),
	461:  uint16(anon_sym_RPAREN),
	462:  uint16(3),
	463:  uint16(2),
	464:  uint16(sym_block_comment),
	465:  uint16(sym_line_comment),
	466:  uint16(129),
	467:  uint16(7),
	468:  uint16(anon_sym_LPAREN),
	469:  uint16(aux_sym_float_literal_token1),
	470:  uint16(aux_sym_float_literal_token2),
	471:  uint16(anon_sym_AMP),
	472:  uint16(anon_sym_STAR),
	473:  uint16(anon_sym_BANG),
	474:  uint16(anon_sym_TILDE),
	475:  uint16(127),
	476:  uint16(43),
	477:  uint16(sym_identifier),
	478:  uint16(sym_int_literal),
	479:  uint16(anon_sym_true),
	480:  uint16(anon_sym_false),
	481:  uint16(anon_sym_bool),
	482:  uint16(anon_sym_u32),
	483:  uint16(anon_sym_i32),
	484:  uint16(anon_sym_f32),
	485:  uint16(anon_sym_f16),
	486:  uint16(anon_sym_array),
	487:  uint16(anon_sym_ptr),
	488:  uint16(anon_sym_sampler),
	489:  uint16(anon_sym_sampler_comparison),
	490:  uint16(anon_sym_texture_depth_2d),
	491:  uint16(anon_sym_texture_depth_2d_array),
	492:  uint16(anon_sym_texture_depth_cube),
	493:  uint16(anon_sym_texture_depth_cube_array),
	494:  uint16(anon_sym_texture_depth_multisampled_2d),
	495:  uint16(anon_sym_texture_1d),
	496:  uint16(anon_sym_texture_2d),
	497:  uint16(anon_sym_texture_2d_array),
	498:  uint16(anon_sym_texture_3d),
	499:  uint16(anon_sym_texture_cube),
	500:  uint16(anon_sym_texture_cube_array),
	501:  uint16(anon_sym_texture_multisampled_2d),
	502:  uint16(anon_sym_texture_storage_1d),
	503:  uint16(anon_sym_texture_storage_2d),
	504:  uint16(anon_sym_texture_storage_2d_array),
	505:  uint16(anon_sym_texture_storage_3d),
	506:  uint16(anon_sym_vec2),
	507:  uint16(anon_sym_vec3),
	508:  uint16(anon_sym_vec4),
	509:  uint16(anon_sym_mat2x2),
	510:  uint16(anon_sym_mat2x3),
	511:  uint16(anon_sym_mat2x4),
	512:  uint16(anon_sym_mat3x2),
	513:  uint16(anon_sym_mat3x3),
	514:  uint16(anon_sym_mat3x4),
	515:  uint16(anon_sym_mat4x2),
	516:  uint16(anon_sym_mat4x3),
	517:  uint16(anon_sym_mat4x4),
	518:  uint16(anon_sym_bitcast),
	519:  uint16(anon_sym_DASH),
	520:  uint16(16),
	521:  uint16(136),
	522:  uint16(1),
	523:  uint16(sym_int_literal),
	524:  uint16(145),
	525:  uint16(1),
	526:  uint16(anon_sym_array),
	527:  uint16(148),
	528:  uint16(1),
	529:  uint16(anon_sym_ptr),
	530:  uint16(38),
	531:  uint16(1),
	532:  uint16(aux_sym_const_expresssion_repeat1),
	533:  uint16(224),
	534:  uint16(1),
	535:  uint16(sym_const_literal),
	536:  uint16(295),
	537:  uint16(1),
	538:  uint16(sym_type_declaration),
	539:  uint16(316),
	540:  uint16(1),
	541:  uint16(sym_const_expression),
	542:  uint16(3),
	543:  uint16(2),
	544:  uint16(sym_block_comment),
	545:  uint16(sym_line_comment),
	546:  uint16(139),
	547:  uint16(2),
	548:  uint16(aux_sym_float_literal_token1),
	549:  uint16(aux_sym_float_literal_token2),
	550:  uint16(142),
	551:  uint16(2),
	552:  uint16(anon_sym_true),
	553:  uint16(anon_sym_false),
	554:  uint16(75),
	555:  uint16(2),
	556:  uint16(sym_float_literal),
	557:  uint16(sym_bool_literal),
	558:  uint16(293),
	559:  uint16(2),
	560:  uint16(sym__vec_prefix),
	561:  uint16(sym__mat_prefix),
	562:  uint16(154),
	563:  uint16(4),
	564:  uint16(anon_sym_texture_storage_1d),
	565:  uint16(anon_sym_texture_storage_2d),
	566:  uint16(anon_sym_texture_storage_2d_array),
	567:  uint16(anon_sym_texture_storage_3d),
	568:  uint16(151),
	569:  uint16(7),
	570:  uint16(anon_sym_texture_1d),
	571:  uint16(anon_sym_texture_2d),
	572:  uint16(anon_sym_texture_2d_array),
	573:  uint16(anon_sym_texture_3d),
	574:  uint16(anon_sym_texture_cube),
	575:  uint16(anon_sym_texture_cube_array),
	576:  uint16(anon_sym_texture_multisampled_2d),
	577:  uint16(157),
	578:  uint16(12),
	579:  uint16(anon_sym_vec2),
	580:  uint16(anon_sym_vec3),
	581:  uint16(anon_sym_vec4),
	582:  uint16(anon_sym_mat2x2),
	583:  uint16(anon_sym_mat2x3),
	584:  uint16(anon_sym_mat2x4),
	585:  uint16(anon_sym_mat3x2),
	586:  uint16(anon_sym_mat3x3),
	587:  uint16(anon_sym_mat3x4),
	588:  uint16(anon_sym_mat4x2),
	589:  uint16(anon_sym_mat4x3),
	590:  uint16(anon_sym_mat4x4),
	591:  uint16(133),
	592:  uint16(13),
	593:  uint16(sym_identifier),
	594:  uint16(anon_sym_bool),
	595:  uint16(anon_sym_u32),
	596:  uint16(anon_sym_i32),
	597:  uint16(anon_sym_f32),
	598:  uint16(anon_sym_f16),
	599:  uint16(anon_sym_sampler),
	600:  uint16(anon_sym_sampler_comparison),
	601:  uint16(anon_sym_texture_depth_2d),
	602:  uint16(anon_sym_texture_depth_2d_array),
	603:  uint16(anon_sym_texture_depth_cube),
	604:  uint16(anon_sym_texture_depth_cube_array),
	605:  uint16(anon_sym_texture_depth_multisampled_2d),
	606:  uint16(16),
	607:  uint16(31),
	608:  uint16(1),
	609:  uint16(sym_int_literal),
	610:  uint16(39),
	611:  uint16(1),
	612:  uint16(anon_sym_array),
	613:  uint16(41),
	614:  uint16(1),
	615:  uint16(anon_sym_ptr),
	616:  uint16(38),
	617:  uint16(1),
	618:  uint16(aux_sym_const_expresssion_repeat1),
	619:  uint16(224),
	620:  uint16(1),
	621:  uint16(sym_const_literal),
	622:  uint16(269),
	623:  uint16(1),
	624:  uint16(sym_const_expression),
	625:  uint16(295),
	626:  uint16(1),
	627:  uint16(sym_type_declaration),
	628:  uint16(3),
	629:  uint16(2),
	630:  uint16(sym_block_comment),
	631:  uint16(sym_line_comment),
	632:  uint16(33),
	633:  uint16(2),
	634:  uint16(aux_sym_float_literal_token1),
	635:  uint16(aux_sym_float_literal_token2),
	636:  uint16(35),
	637:  uint16(2),
	638:  uint16(anon_sym_true),
	639:  uint16(anon_sym_false),
	640:  uint16(75),
	641:  uint16(2),
	642:  uint16(sym_float_literal),
	643:  uint16(sym_bool_literal),
	644:  uint16(293),
	645:  uint16(2),
	646:  uint16(sym__vec_prefix),
	647:  uint16(sym__mat_prefix),
	648:  uint16(45),
	649:  uint16(4),
	650:  uint16(anon_sym_texture_storage_1d),
	651:  uint16(anon_sym_texture_storage_2d),
	652:  uint16(anon_sym_texture_storage_2d_array),
	653:  uint16(anon_sym_texture_storage_3d),
	654:  uint16(43),
	655:  uint16(7),
	656:  uint16(anon_sym_texture_1d),
	657:  uint16(anon_sym_texture_2d),
	658:  uint16(anon_sym_texture_2d_array),
	659:  uint16(anon_sym_texture_3d),
	660:  uint16(anon_sym_texture_cube),
	661:  uint16(anon_sym_texture_cube_array),
	662:  uint16(anon_sym_texture_multisampled_2d),
	663:  uint16(121),
	664:  uint16(12),
	665:  uint16(anon_sym_vec2),
	666:  uint16(anon_sym_vec3),
	667:  uint16(anon_sym_vec4),
	668:  uint16(anon_sym_mat2x2),
	669:  uint16(anon_sym_mat2x3),
	670:  uint16(anon_sym_mat2x4),
	671:  uint16(anon_sym_mat3x2),
	672:  uint16(anon_sym_mat3x3),
	673:  uint16(anon_sym_mat3x4),
	674:  uint16(anon_sym_mat4x2),
	675:  uint16(anon_sym_mat4x3),
	676:  uint16(anon_sym_mat4x4),
	677:  uint16(37),
	678:  uint16(13),
	679:  uint16(sym_identifier),
	680:  uint16(anon_sym_bool),
	681:  uint16(anon_sym_u32),
	682:  uint16(anon_sym_i32),
	683:  uint16(anon_sym_f32),
	684:  uint16(anon_sym_f16),
	685:  uint16(anon_sym_sampler),
	686:  uint16(anon_sym_sampler_comparison),
	687:  uint16(anon_sym_texture_depth_2d),
	688:  uint16(anon_sym_texture_depth_2d_array),
	689:  uint16(anon_sym_texture_depth_cube),
	690:  uint16(anon_sym_texture_depth_cube_array),
	691:  uint16(anon_sym_texture_depth_multisampled_2d),
	692:  uint16(4),
	693:  uint16(160),
	694:  uint16(1),
	695:  uint16(anon_sym_RPAREN),
	696:  uint16(3),
	697:  uint16(2),
	698:  uint16(sym_block_comment),
	699:  uint16(sym_line_comment),
	700:  uint16(129),
	701:  uint16(7),
	702:  uint16(anon_sym_LPAREN),
	703:  uint16(aux_sym_float_literal_token1),
	704:  uint16(aux_sym_float_literal_token2),
	705:  uint16(anon_sym_AMP),
	706:  uint16(anon_sym_STAR),
	707:  uint16(anon_sym_BANG),
	708:  uint16(anon_sym_TILDE),
	709:  uint16(127),
	710:  uint16(43),
	711:  uint16(sym_identifier),
	712:  uint16(sym_int_literal),
	713:  uint16(anon_sym_true),
	714:  uint16(anon_sym_false),
	715:  uint16(anon_sym_bool),
	716:  uint16(anon_sym_u32),
	717:  uint16(anon_sym_i32),
	718:  uint16(anon_sym_f32),
	719:  uint16(anon_sym_f16),
	720:  uint16(anon_sym_array),
	721:  uint16(anon_sym_ptr),
	722:  uint16(anon_sym_sampler),
	723:  uint16(anon_sym_sampler_comparison),
	724:  uint16(anon_sym_texture_depth_2d),
	725:  uint16(anon_sym_texture_depth_2d_array),
	726:  uint16(anon_sym_texture_depth_cube),
	727:  uint16(anon_sym_texture_depth_cube_array),
	728:  uint16(anon_sym_texture_depth_multisampled_2d),
	729:  uint16(anon_sym_texture_1d),
	730:  uint16(anon_sym_texture_2d),
	731:  uint16(anon_sym_texture_2d_array),
	732:  uint16(anon_sym_texture_3d),
	733:  uint16(anon_sym_texture_cube),
	734:  uint16(anon_sym_texture_cube_array),
	735:  uint16(anon_sym_texture_multisampled_2d),
	736:  uint16(anon_sym_texture_storage_1d),
	737:  uint16(anon_sym_texture_storage_2d),
	738:  uint16(anon_sym_texture_storage_2d_array),
	739:  uint16(anon_sym_texture_storage_3d),
	740:  uint16(anon_sym_vec2),
	741:  uint16(anon_sym_vec3),
	742:  uint16(anon_sym_vec4),
	743:  uint16(anon_sym_mat2x2),
	744:  uint16(anon_sym_mat2x3),
	745:  uint16(anon_sym_mat2x4),
	746:  uint16(anon_sym_mat3x2),
	747:  uint16(anon_sym_mat3x3),
	748:  uint16(anon_sym_mat3x4),
	749:  uint16(anon_sym_mat4x2),
	750:  uint16(anon_sym_mat4x3),
	751:  uint16(anon_sym_mat4x4),
	752:  uint16(anon_sym_bitcast),
	753:  uint16(anon_sym_DASH),
	754:  uint16(15),
	755:  uint16(31),
	756:  uint16(1),
	757:  uint16(sym_int_literal),
	758:  uint16(39),
	759:  uint16(1),
	760:  uint16(anon_sym_array),
	761:  uint16(41),
	762:  uint16(1),
	763:  uint16(anon_sym_ptr),
	764:  uint16(224),
	765:  uint16(1),
	766:  uint16(sym_const_literal),
	767:  uint16(291),
	768:  uint16(1),
	769:  uint16(sym_const_expression),
	770:  uint16(295),
	771:  uint16(1),
	772:  uint16(sym_type_declaration),
	773:  uint16(3),
	774:  uint16(2),
	775:  uint16(sym_block_comment),
	776:  uint16(sym_line_comment),
	777:  uint16(33),
	778:  uint16(2),
	779:  uint16(aux_sym_float_literal_token1),
	780:  uint16(aux_sym_float_literal_token2),
	781:  uint16(35),
	782:  uint16(2),
	783:  uint16(anon_sym_true),
	784:  uint16(anon_sym_false),
	785:  uint16(75),
	786:  uint16(2),
	787:  uint16(sym_float_literal),
	788:  uint16(sym_bool_literal),
	789:  uint16(293),
	790:  uint16(2),
	791:  uint16(sym__vec_prefix),
	792:  uint16(sym__mat_prefix),
	793:  uint16(45),
	794:  uint16(4),
	795:  uint16(anon_sym_texture_storage_1d),
	796:  uint16(anon_sym_texture_storage_2d),
	797:  uint16(anon_sym_texture_storage_2d_array),
	798:  uint16(anon_sym_texture_storage_3d),
	799:  uint16(43),
	800:  uint16(7),
	801:  uint16(anon_sym_texture_1d),
	802:  uint16(anon_sym_texture_2d),
	803:  uint16(anon_sym_texture_2d_array),
	804:  uint16(anon_sym_texture_3d),
	805:  uint16(anon_sym_texture_cube),
	806:  uint16(anon_sym_texture_cube_array),
	807:  uint16(anon_sym_texture_multisampled_2d),
	808:  uint16(121),
	809:  uint16(12),
	810:  uint16(anon_sym_vec2),
	811:  uint16(anon_sym_vec3),
	812:  uint16(anon_sym_vec4),
	813:  uint16(anon_sym_mat2x2),
	814:  uint16(anon_sym_mat2x3),
	815:  uint16(anon_sym_mat2x4),
	816:  uint16(anon_sym_mat3x2),
	817:  uint16(anon_sym_mat3x3),
	818:  uint16(anon_sym_mat3x4),
	819:  uint16(anon_sym_mat4x2),
	820:  uint16(anon_sym_mat4x3),
	821:  uint16(anon_sym_mat4x4),
	822:  uint16(37),
	823:  uint16(13),
	824:  uint16(sym_identifier),
	825:  uint16(anon_sym_bool),
	826:  uint16(anon_sym_u32),
	827:  uint16(anon_sym_i32),
	828:  uint16(anon_sym_f32),
	829:  uint16(anon_sym_f16),
	830:  uint16(anon_sym_sampler),
	831:  uint16(anon_sym_sampler_comparison),
	832:  uint16(anon_sym_texture_depth_2d),
	833:  uint16(anon_sym_texture_depth_2d_array),
	834:  uint16(anon_sym_texture_depth_cube),
	835:  uint16(anon_sym_texture_depth_cube_array),
	836:  uint16(anon_sym_texture_depth_multisampled_2d),
	837:  uint16(3),
	838:  uint16(3),
	839:  uint16(2),
	840:  uint16(sym_block_comment),
	841:  uint16(sym_line_comment),
	842:  uint16(129),
	843:  uint16(7),
	844:  uint16(anon_sym_LPAREN),
	845:  uint16(aux_sym_float_literal_token1),
	846:  uint16(aux_sym_float_literal_token2),
	847:  uint16(anon_sym_AMP),
	848:  uint16(anon_sym_STAR),
	849:  uint16(anon_sym_BANG),
	850:  uint16(anon_sym_TILDE),
	851:  uint16(127),
	852:  uint16(43),
	853:  uint16(sym_identifier),
	854:  uint16(sym_int_literal),
	855:  uint16(anon_sym_true),
	856:  uint16(anon_sym_false),
	857:  uint16(anon_sym_bool),
	858:  uint16(anon_sym_u32),
	859:  uint16(anon_sym_i32),
	860:  uint16(anon_sym_f32),
	861:  uint16(anon_sym_f16),
	862:  uint16(anon_sym_array),
	863:  uint16(anon_sym_ptr),
	864:  uint16(anon_sym_sampler),
	865:  uint16(anon_sym_sampler_comparison),
	866:  uint16(anon_sym_texture_depth_2d),
	867:  uint16(anon_sym_texture_depth_2d_array),
	868:  uint16(anon_sym_texture_depth_cube),
	869:  uint16(anon_sym_texture_depth_cube_array),
	870:  uint16(anon_sym_texture_depth_multisampled_2d),
	871:  uint16(anon_sym_texture_1d),
	872:  uint16(anon_sym_texture_2d),
	873:  uint16(anon_sym_texture_2d_array),
	874:  uint16(anon_sym_texture_3d),
	875:  uint16(anon_sym_texture_cube),
	876:  uint16(anon_sym_texture_cube_array),
	877:  uint16(anon_sym_texture_multisampled_2d),
	878:  uint16(anon_sym_texture_storage_1d),
	879:  uint16(anon_sym_texture_storage_2d),
	880:  uint16(anon_sym_texture_storage_2d_array),
	881:  uint16(anon_sym_texture_storage_3d),
	882:  uint16(anon_sym_vec2),
	883:  uint16(anon_sym_vec3),
	884:  uint16(anon_sym_vec4),
	885:  uint16(anon_sym_mat2x2),
	886:  uint16(anon_sym_mat2x3),
	887:  uint16(anon_sym_mat2x4),
	888:  uint16(anon_sym_mat3x2),
	889:  uint16(anon_sym_mat3x3),
	890:  uint16(anon_sym_mat3x4),
	891:  uint16(anon_sym_mat4x2),
	892:  uint16(anon_sym_mat4x3),
	893:  uint16(anon_sym_mat4x4),
	894:  uint16(anon_sym_bitcast),
	895:  uint16(anon_sym_DASH),
	896:  uint16(15),
	897:  uint16(31),
	898:  uint16(1),
	899:  uint16(sym_int_literal),
	900:  uint16(39),
	901:  uint16(1),
	902:  uint16(anon_sym_array),
	903:  uint16(41),
	904:  uint16(1),
	905:  uint16(anon_sym_ptr),
	906:  uint16(224),
	907:  uint16(1),
	908:  uint16(sym_const_literal),
	909:  uint16(295),
	910:  uint16(1),
	911:  uint16(sym_type_declaration),
	912:  uint16(297),
	913:  uint16(1),
	914:  uint16(sym_const_expression),
	915:  uint16(3),
	916:  uint16(2),
	917:  uint16(sym_block_comment),
	918:  uint16(sym_line_comment),
	919:  uint16(33),
	920:  uint16(2),
	921:  uint16(aux_sym_float_literal_token1),
	922:  uint16(aux_sym_float_literal_token2),
	923:  uint16(35),
	924:  uint16(2),
	925:  uint16(anon_sym_true),
	926:  uint16(anon_sym_false),
	927:  uint16(75),
	928:  uint16(2),
	929:  uint16(sym_float_literal),
	930:  uint16(sym_bool_literal),
	931:  uint16(293),
	932:  uint16(2),
	933:  uint16(sym__vec_prefix),
	934:  uint16(sym__mat_prefix),
	935:  uint16(45),
	936:  uint16(4),
	937:  uint16(anon_sym_texture_storage_1d),
	938:  uint16(anon_sym_texture_storage_2d),
	939:  uint16(anon_sym_texture_storage_2d_array),
	940:  uint16(anon_sym_texture_storage_3d),
	941:  uint16(43),
	942:  uint16(7),
	943:  uint16(anon_sym_texture_1d),
	944:  uint16(anon_sym_texture_2d),
	945:  uint16(anon_sym_texture_2d_array),
	946:  uint16(anon_sym_texture_3d),
	947:  uint16(anon_sym_texture_cube),
	948:  uint16(anon_sym_texture_cube_array),
	949:  uint16(anon_sym_texture_multisampled_2d),
	950:  uint16(121),
	951:  uint16(12),
	952:  uint16(anon_sym_vec2),
	953:  uint16(anon_sym_vec3),
	954:  uint16(anon_sym_vec4),
	955:  uint16(anon_sym_mat2x2),
	956:  uint16(anon_sym_mat2x3),
	957:  uint16(anon_sym_mat2x4),
	958:  uint16(anon_sym_mat3x2),
	959:  uint16(anon_sym_mat3x3),
	960:  uint16(anon_sym_mat3x4),
	961:  uint16(anon_sym_mat4x2),
	962:  uint16(anon_sym_mat4x3),
	963:  uint16(anon_sym_mat4x4),
	964:  uint16(37),
	965:  uint16(13),
	966:  uint16(sym_identifier),
	967:  uint16(anon_sym_bool),
	968:  uint16(anon_sym_u32),
	969:  uint16(anon_sym_i32),
	970:  uint16(anon_sym_f32),
	971:  uint16(anon_sym_f16),
	972:  uint16(anon_sym_sampler),
	973:  uint16(anon_sym_sampler_comparison),
	974:  uint16(anon_sym_texture_depth_2d),
	975:  uint16(anon_sym_texture_depth_2d_array),
	976:  uint16(anon_sym_texture_depth_cube),
	977:  uint16(anon_sym_texture_depth_cube_array),
	978:  uint16(anon_sym_texture_depth_multisampled_2d),
	979:  uint16(15),
	980:  uint16(31),
	981:  uint16(1),
	982:  uint16(sym_int_literal),
	983:  uint16(39),
	984:  uint16(1),
	985:  uint16(anon_sym_array),
	986:  uint16(41),
	987:  uint16(1),
	988:  uint16(anon_sym_ptr),
	989:  uint16(224),
	990:  uint16(1),
	991:  uint16(sym_const_literal),
	992:  uint16(295),
	993:  uint16(1),
	994:  uint16(sym_type_declaration),
	995:  uint16(311),
	996:  uint16(1),
	997:  uint16(sym_const_expression),
	998:  uint16(3),
	999:  uint16(2),
	1000: uint16(sym_block_comment),
	1001: uint16(sym_line_comment),
	1002: uint16(33),
	1003: uint16(2),
	1004: uint16(aux_sym_float_literal_token1),
	1005: uint16(aux_sym_float_literal_token2),
	1006: uint16(35),
	1007: uint16(2),
	1008: uint16(anon_sym_true),
	1009: uint16(anon_sym_false),
	1010: uint16(75),
	1011: uint16(2),
	1012: uint16(sym_float_literal),
	1013: uint16(sym_bool_literal),
	1014: uint16(293),
	1015: uint16(2),
	1016: uint16(sym__vec_prefix),
	1017: uint16(sym__mat_prefix),
	1018: uint16(45),
	1019: uint16(4),
	1020: uint16(anon_sym_texture_storage_1d),
	1021: uint16(anon_sym_texture_storage_2d),
	1022: uint16(anon_sym_texture_storage_2d_array),
	1023: uint16(anon_sym_texture_storage_3d),
	1024: uint16(43),
	1025: uint16(7),
	1026: uint16(anon_sym_texture_1d),
	1027: uint16(anon_sym_texture_2d),
	1028: uint16(anon_sym_texture_2d_array),
	1029: uint16(anon_sym_texture_3d),
	1030: uint16(anon_sym_texture_cube),
	1031: uint16(anon_sym_texture_cube_array),
	1032: uint16(anon_sym_texture_multisampled_2d),
	1033: uint16(121),
	1034: uint16(12),
	1035: uint16(anon_sym_vec2),
	1036: uint16(anon_sym_vec3),
	1037: uint16(anon_sym_vec4),
	1038: uint16(anon_sym_mat2x2),
	1039: uint16(anon_sym_mat2x3),
	1040: uint16(anon_sym_mat2x4),
	1041: uint16(anon_sym_mat3x2),
	1042: uint16(anon_sym_mat3x3),
	1043: uint16(anon_sym_mat3x4),
	1044: uint16(anon_sym_mat4x2),
	1045: uint16(anon_sym_mat4x3),
	1046: uint16(anon_sym_mat4x4),
	1047: uint16(37),
	1048: uint16(13),
	1049: uint16(sym_identifier),
	1050: uint16(anon_sym_bool),
	1051: uint16(anon_sym_u32),
	1052: uint16(anon_sym_i32),
	1053: uint16(anon_sym_f32),
	1054: uint16(anon_sym_f16),
	1055: uint16(anon_sym_sampler),
	1056: uint16(anon_sym_sampler_comparison),
	1057: uint16(anon_sym_texture_depth_2d),
	1058: uint16(anon_sym_texture_depth_2d_array),
	1059: uint16(anon_sym_texture_depth_cube),
	1060: uint16(anon_sym_texture_depth_cube_array),
	1061: uint16(anon_sym_texture_depth_multisampled_2d),
	1062: uint16(3),
	1063: uint16(3),
	1064: uint16(2),
	1065: uint16(sym_block_comment),
	1066: uint16(sym_line_comment),
	1067: uint16(164),
	1068: uint16(7),
	1069: uint16(anon_sym_LPAREN),
	1070: uint16(aux_sym_float_literal_token1),
	1071: uint16(aux_sym_float_literal_token2),
	1072: uint16(anon_sym_AMP),
	1073: uint16(anon_sym_STAR),
	1074: uint16(anon_sym_BANG),
	1075: uint16(anon_sym_TILDE),
	1076: uint16(162),
	1077: uint16(43),
	1078: uint16(sym_identifier),
	1079: uint16(sym_int_literal),
	1080: uint16(anon_sym_true),
	1081: uint16(anon_sym_false),
	1082: uint16(anon_sym_bool),
	1083: uint16(anon_sym_u32),
	1084: uint16(anon_sym_i32),
	1085: uint16(anon_sym_f32),
	1086: uint16(anon_sym_f16),
	1087: uint16(anon_sym_array),
	1088: uint16(anon_sym_ptr),
	1089: uint16(anon_sym_sampler),
	1090: uint16(anon_sym_sampler_comparison),
	1091: uint16(anon_sym_texture_depth_2d),
	1092: uint16(anon_sym_texture_depth_2d_array),
	1093: uint16(anon_sym_texture_depth_cube),
	1094: uint16(anon_sym_texture_depth_cube_array),
	1095: uint16(anon_sym_texture_depth_multisampled_2d),
	1096: uint16(anon_sym_texture_1d),
	1097: uint16(anon_sym_texture_2d),
	1098: uint16(anon_sym_texture_2d_array),
	1099: uint16(anon_sym_texture_3d),
	1100: uint16(anon_sym_texture_cube),
	1101: uint16(anon_sym_texture_cube_array),
	1102: uint16(anon_sym_texture_multisampled_2d),
	1103: uint16(anon_sym_texture_storage_1d),
	1104: uint16(anon_sym_texture_storage_2d),
	1105: uint16(anon_sym_texture_storage_2d_array),
	1106: uint16(anon_sym_texture_storage_3d),
	1107: uint16(anon_sym_vec2),
	1108: uint16(anon_sym_vec3),
	1109: uint16(anon_sym_vec4),
	1110: uint16(anon_sym_mat2x2),
	1111: uint16(anon_sym_mat2x3),
	1112: uint16(anon_sym_mat2x4),
	1113: uint16(anon_sym_mat3x2),
	1114: uint16(anon_sym_mat3x3),
	1115: uint16(anon_sym_mat3x4),
	1116: uint16(anon_sym_mat4x2),
	1117: uint16(anon_sym_mat4x3),
	1118: uint16(anon_sym_mat4x4),
	1119: uint16(anon_sym_bitcast),
	1120: uint16(anon_sym_DASH),
	1121: uint16(4),
	1122: uint16(168),
	1123: uint16(1),
	1124: uint16(anon_sym_RPAREN),
	1125: uint16(3),
	1126: uint16(2),
	1127: uint16(sym_block_comment),
	1128: uint16(sym_line_comment),
	1129: uint16(170),
	1130: uint16(2),
	1131: uint16(aux_sym_float_literal_token1),
	1132: uint16(aux_sym_float_literal_token2),
	1133: uint16(166),
	1134: uint16(41),
	1135: uint16(sym_identifier),
	1136: uint16(sym_int_literal),
	1137: uint16(anon_sym_true),
	1138: uint16(anon_sym_false),
	1139: uint16(anon_sym_bool),
	1140: uint16(anon_sym_u32),
	1141: uint16(anon_sym_i32),
	1142: uint16(anon_sym_f32),
	1143: uint16(anon_sym_f16),
	1144: uint16(anon_sym_array),
	1145: uint16(anon_sym_ptr),
	1146: uint16(anon_sym_sampler),
	1147: uint16(anon_sym_sampler_comparison),
	1148: uint16(anon_sym_texture_depth_2d),
	1149: uint16(anon_sym_texture_depth_2d_array),
	1150: uint16(anon_sym_texture_depth_cube),
	1151: uint16(anon_sym_texture_depth_cube_array),
	1152: uint16(anon_sym_texture_depth_multisampled_2d),
	1153: uint16(anon_sym_texture_1d),
	1154: uint16(anon_sym_texture_2d),
	1155: uint16(anon_sym_texture_2d_array),
	1156: uint16(anon_sym_texture_3d),
	1157: uint16(anon_sym_texture_cube),
	1158: uint16(anon_sym_texture_cube_array),
	1159: uint16(anon_sym_texture_multisampled_2d),
	1160: uint16(anon_sym_texture_storage_1d),
	1161: uint16(anon_sym_texture_storage_2d),
	1162: uint16(anon_sym_texture_storage_2d_array),
	1163: uint16(anon_sym_texture_storage_3d),
	1164: uint16(anon_sym_vec2),
	1165: uint16(anon_sym_vec3),
	1166: uint16(anon_sym_vec4),
	1167: uint16(anon_sym_mat2x2),
	1168: uint16(anon_sym_mat2x3),
	1169: uint16(anon_sym_mat2x4),
	1170: uint16(anon_sym_mat3x2),
	1171: uint16(anon_sym_mat3x3),
	1172: uint16(anon_sym_mat3x4),
	1173: uint16(anon_sym_mat4x2),
	1174: uint16(anon_sym_mat4x3),
	1175: uint16(anon_sym_mat4x4),
	1176: uint16(4),
	1177: uint16(174),
	1178: uint16(1),
	1179: uint16(anon_sym_AT),
	1180: uint16(3),
	1181: uint16(2),
	1182: uint16(sym_block_comment),
	1183: uint16(sym_line_comment),
	1184: uint16(47),
	1185: uint16(2),
	1186: uint16(sym_attribute),
	1187: uint16(aux_sym_global_variable_declaration_repeat1),
	1188: uint16(172),
	1189: uint16(41),
	1190: uint16(anon_sym_override),
	1191: uint16(anon_sym_fn),
	1192: uint16(sym_identifier),
	1193: uint16(anon_sym_var),
	1194: uint16(anon_sym_bool),
	1195: uint16(anon_sym_u32),
	1196: uint16(anon_sym_i32),
	1197: uint16(anon_sym_f32),
	1198: uint16(anon_sym_f16),
	1199: uint16(anon_sym_array),
	1200: uint16(anon_sym_ptr),
	1201: uint16(anon_sym_sampler),
	1202: uint16(anon_sym_sampler_comparison),
	1203: uint16(anon_sym_texture_depth_2d),
	1204: uint16(anon_sym_texture_depth_2d_array),
	1205: uint16(anon_sym_texture_depth_cube),
	1206: uint16(anon_sym_texture_depth_cube_array),
	1207: uint16(anon_sym_texture_depth_multisampled_2d),
	1208: uint16(anon_sym_texture_1d),
	1209: uint16(anon_sym_texture_2d),
	1210: uint16(anon_sym_texture_2d_array),
	1211: uint16(anon_sym_texture_3d),
	1212: uint16(anon_sym_texture_cube),
	1213: uint16(anon_sym_texture_cube_array),
	1214: uint16(anon_sym_texture_multisampled_2d),
	1215: uint16(anon_sym_texture_storage_1d),
	1216: uint16(anon_sym_texture_storage_2d),
	1217: uint16(anon_sym_texture_storage_2d_array),
	1218: uint16(anon_sym_texture_storage_3d),
	1219: uint16(anon_sym_vec2),
	1220: uint16(anon_sym_vec3),
	1221: uint16(anon_sym_vec4),
	1222: uint16(anon_sym_mat2x2),
	1223: uint16(anon_sym_mat2x3),
	1224: uint16(anon_sym_mat2x4),
	1225: uint16(anon_sym_mat3x2),
	1226: uint16(anon_sym_mat3x3),
	1227: uint16(anon_sym_mat3x4),
	1228: uint16(anon_sym_mat4x2),
	1229: uint16(anon_sym_mat4x3),
	1230: uint16(anon_sym_mat4x4),
	1231: uint16(4),
	1232: uint16(177),
	1233: uint16(1),
	1234: uint16(anon_sym_RPAREN),
	1235: uint16(3),
	1236: uint16(2),
	1237: uint16(sym_block_comment),
	1238: uint16(sym_line_comment),
	1239: uint16(170),
	1240: uint16(2),
	1241: uint16(aux_sym_float_literal_token1),
	1242: uint16(aux_sym_float_literal_token2),
	1243: uint16(166),
	1244: uint16(41),
	1245: uint16(sym_identifier),
	1246: uint16(sym_int_literal),
	1247: uint16(anon_sym_true),
	1248: uint16(anon_sym_false),
	1249: uint16(anon_sym_bool),
	1250: uint16(anon_sym_u32),
	1251: uint16(anon_sym_i32),
	1252: uint16(anon_sym_f32),
	1253: uint16(anon_sym_f16),
	1254: uint16(anon_sym_array),
	1255: uint16(anon_sym_ptr),
	1256: uint16(anon_sym_sampler),
	1257: uint16(anon_sym_sampler_comparison),
	1258: uint16(anon_sym_texture_depth_2d),
	1259: uint16(anon_sym_texture_depth_2d_array),
	1260: uint16(anon_sym_texture_depth_cube),
	1261: uint16(anon_sym_texture_depth_cube_array),
	1262: uint16(anon_sym_texture_depth_multisampled_2d),
	1263: uint16(anon_sym_texture_1d),
	1264: uint16(anon_sym_texture_2d),
	1265: uint16(anon_sym_texture_2d_array),
	1266: uint16(anon_sym_texture_3d),
	1267: uint16(anon_sym_texture_cube),
	1268: uint16(anon_sym_texture_cube_array),
	1269: uint16(anon_sym_texture_multisampled_2d),
	1270: uint16(anon_sym_texture_storage_1d),
	1271: uint16(anon_sym_texture_storage_2d),
	1272: uint16(anon_sym_texture_storage_2d_array),
	1273: uint16(anon_sym_texture_storage_3d),
	1274: uint16(anon_sym_vec2),
	1275: uint16(anon_sym_vec3),
	1276: uint16(anon_sym_vec4),
	1277: uint16(anon_sym_mat2x2),
	1278: uint16(anon_sym_mat2x3),
	1279: uint16(anon_sym_mat2x4),
	1280: uint16(anon_sym_mat3x2),
	1281: uint16(anon_sym_mat3x3),
	1282: uint16(anon_sym_mat3x4),
	1283: uint16(anon_sym_mat4x2),
	1284: uint16(anon_sym_mat4x3),
	1285: uint16(anon_sym_mat4x4),
	1286: uint16(11),
	1287: uint16(21),
	1288: uint16(1),
	1289: uint16(anon_sym_AT),
	1290: uint16(39),
	1291: uint16(1),
	1292: uint16(anon_sym_array),
	1293: uint16(41),
	1294: uint16(1),
	1295: uint16(anon_sym_ptr),
	1296: uint16(294),
	1297: uint16(1),
	1298: uint16(sym_type_declaration),
	1299: uint16(3),
	1300: uint16(2),
	1301: uint16(sym_block_comment),
	1302: uint16(sym_line_comment),
	1303: uint16(47),
	1304: uint16(2),
	1305: uint16(sym_attribute),
	1306: uint16(aux_sym_global_variable_declaration_repeat1),
	1307: uint16(293),
	1308: uint16(2),
	1309: uint16(sym__vec_prefix),
	1310: uint16(sym__mat_prefix),
	1311: uint16(45),
	1312: uint16(4),
	1313: uint16(anon_sym_texture_storage_1d),
	1314: uint16(anon_sym_texture_storage_2d),
	1315: uint16(anon_sym_texture_storage_2d_array),
	1316: uint16(anon_sym_texture_storage_3d),
	1317: uint16(43),
	1318: uint16(7),
	1319: uint16(anon_sym_texture_1d),
	1320: uint16(anon_sym_texture_2d),
	1321: uint16(anon_sym_texture_2d_array),
	1322: uint16(anon_sym_texture_3d),
	1323: uint16(anon_sym_texture_cube),
	1324: uint16(anon_sym_texture_cube_array),
	1325: uint16(anon_sym_texture_multisampled_2d),
	1326: uint16(121),
	1327: uint16(12),
	1328: uint16(anon_sym_vec2),
	1329: uint16(anon_sym_vec3),
	1330: uint16(anon_sym_vec4),
	1331: uint16(anon_sym_mat2x2),
	1332: uint16(anon_sym_mat2x3),
	1333: uint16(anon_sym_mat2x4),
	1334: uint16(anon_sym_mat3x2),
	1335: uint16(anon_sym_mat3x3),
	1336: uint16(anon_sym_mat3x4),
	1337: uint16(anon_sym_mat4x2),
	1338: uint16(anon_sym_mat4x3),
	1339: uint16(anon_sym_mat4x4),
	1340: uint16(37),
	1341: uint16(13),
	1342: uint16(sym_identifier),
	1343: uint16(anon_sym_bool),
	1344: uint16(anon_sym_u32),
	1345: uint16(anon_sym_i32),
	1346: uint16(anon_sym_f32),
	1347: uint16(anon_sym_f16),
	1348: uint16(anon_sym_sampler),
	1349: uint16(anon_sym_sampler_comparison),
	1350: uint16(anon_sym_texture_depth_2d),
	1351: uint16(anon_sym_texture_depth_2d_array),
	1352: uint16(anon_sym_texture_depth_cube),
	1353: uint16(anon_sym_texture_depth_cube_array),
	1354: uint16(anon_sym_texture_depth_multisampled_2d),
	1355: uint16(11),
	1356: uint16(21),
	1357: uint16(1),
	1358: uint16(anon_sym_AT),
	1359: uint16(39),
	1360: uint16(1),
	1361: uint16(anon_sym_array),
	1362: uint16(41),
	1363: uint16(1),
	1364: uint16(anon_sym_ptr),
	1365: uint16(301),
	1366: uint16(1),
	1367: uint16(sym_type_declaration),
	1368: uint16(3),
	1369: uint16(2),
	1370: uint16(sym_block_comment),
	1371: uint16(sym_line_comment),
	1372: uint16(49),
	1373: uint16(2),
	1374: uint16(sym_attribute),
	1375: uint16(aux_sym_global_variable_declaration_repeat1),
	1376: uint16(293),
	1377: uint16(2),
	1378: uint16(sym__vec_prefix),
	1379: uint16(sym__mat_prefix),
	1380: uint16(45),
	1381: uint16(4),
	1382: uint16(anon_sym_texture_storage_1d),
	1383: uint16(anon_sym_texture_storage_2d),
	1384: uint16(anon_sym_texture_storage_2d_array),
	1385: uint16(anon_sym_texture_storage_3d),
	1386: uint16(43),
	1387: uint16(7),
	1388: uint16(anon_sym_texture_1d),
	1389: uint16(anon_sym_texture_2d),
	1390: uint16(anon_sym_texture_2d_array),
	1391: uint16(anon_sym_texture_3d),
	1392: uint16(anon_sym_texture_cube),
	1393: uint16(anon_sym_texture_cube_array),
	1394: uint16(anon_sym_texture_multisampled_2d),
	1395: uint16(121),
	1396: uint16(12),
	1397: uint16(anon_sym_vec2),
	1398: uint16(anon_sym_vec3),
	1399: uint16(anon_sym_vec4),
	1400: uint16(anon_sym_mat2x2),
	1401: uint16(anon_sym_mat2x3),
	1402: uint16(anon_sym_mat2x4),
	1403: uint16(anon_sym_mat3x2),
	1404: uint16(anon_sym_mat3x3),
	1405: uint16(anon_sym_mat3x4),
	1406: uint16(anon_sym_mat4x2),
	1407: uint16(anon_sym_mat4x3),
	1408: uint16(anon_sym_mat4x4),
	1409: uint16(37),
	1410: uint16(13),
	1411: uint16(sym_identifier),
	1412: uint16(anon_sym_bool),
	1413: uint16(anon_sym_u32),
	1414: uint16(anon_sym_i32),
	1415: uint16(anon_sym_f32),
	1416: uint16(anon_sym_f16),
	1417: uint16(anon_sym_sampler),
	1418: uint16(anon_sym_sampler_comparison),
	1419: uint16(anon_sym_texture_depth_2d),
	1420: uint16(anon_sym_texture_depth_2d_array),
	1421: uint16(anon_sym_texture_depth_cube),
	1422: uint16(anon_sym_texture_depth_cube_array),
	1423: uint16(anon_sym_texture_depth_multisampled_2d),
	1424: uint16(3),
	1425: uint16(3),
	1426: uint16(2),
	1427: uint16(sym_block_comment),
	1428: uint16(sym_line_comment),
	1429: uint16(170),
	1430: uint16(2),
	1431: uint16(aux_sym_float_literal_token1),
	1432: uint16(aux_sym_float_literal_token2),
	1433: uint16(166),
	1434: uint16(41),
	1435: uint16(sym_identifier),
	1436: uint16(sym_int_literal),
	1437: uint16(anon_sym_true),
	1438: uint16(anon_sym_false),
	1439: uint16(anon_sym_bool),
	1440: uint16(anon_sym_u32),
	1441: uint16(anon_sym_i32),
	1442: uint16(anon_sym_f32),
	1443: uint16(anon_sym_f16),
	1444: uint16(anon_sym_array),
	1445: uint16(anon_sym_ptr),
	1446: uint16(anon_sym_sampler),
	1447: uint16(anon_sym_sampler_comparison),
	1448: uint16(anon_sym_texture_depth_2d),
	1449: uint16(anon_sym_texture_depth_2d_array),
	1450: uint16(anon_sym_texture_depth_cube),
	1451: uint16(anon_sym_texture_depth_cube_array),
	1452: uint16(anon_sym_texture_depth_multisampled_2d),
	1453: uint16(anon_sym_texture_1d),
	1454: uint16(anon_sym_texture_2d),
	1455: uint16(anon_sym_texture_2d_array),
	1456: uint16(anon_sym_texture_3d),
	1457: uint16(anon_sym_texture_cube),
	1458: uint16(anon_sym_texture_cube_array),
	1459: uint16(anon_sym_texture_multisampled_2d),
	1460: uint16(anon_sym_texture_storage_1d),
	1461: uint16(anon_sym_texture_storage_2d),
	1462: uint16(anon_sym_texture_storage_2d_array),
	1463: uint16(anon_sym_texture_storage_3d),
	1464: uint16(anon_sym_vec2),
	1465: uint16(anon_sym_vec3),
	1466: uint16(anon_sym_vec4),
	1467: uint16(anon_sym_mat2x2),
	1468: uint16(anon_sym_mat2x3),
	1469: uint16(anon_sym_mat2x4),
	1470: uint16(anon_sym_mat3x2),
	1471: uint16(anon_sym_mat3x3),
	1472: uint16(anon_sym_mat3x4),
	1473: uint16(anon_sym_mat4x2),
	1474: uint16(anon_sym_mat4x3),
	1475: uint16(anon_sym_mat4x4),
	1476: uint16(4),
	1477: uint16(181),
	1478: uint16(1),
	1479: uint16(anon_sym_LPAREN),
	1480: uint16(183),
	1481: uint16(1),
	1482: uint16(anon_sym_AT),
	1483: uint16(3),
	1484: uint16(2),
	1485: uint16(sym_block_comment),
	1486: uint16(sym_line_comment),
	1487: uint16(179),
	1488: uint16(41),
	1489: uint16(anon_sym_override),
	1490: uint16(anon_sym_fn),
	1491: uint16(sym_identifier),
	1492: uint16(anon_sym_var),
	1493: uint16(anon_sym_bool),
	1494: uint16(anon_sym_u32),
	1495: uint16(anon_sym_i32),
	1496: uint16(anon_sym_f32),
	1497: uint16(anon_sym_f16),
	1498: uint16(anon_sym_array),
	1499: uint16(anon_sym_ptr),
	1500: uint16(anon_sym_sampler),
	1501: uint16(anon_sym_sampler_comparison),
	1502: uint16(anon_sym_texture_depth_2d),
	1503: uint16(anon_sym_texture_depth_2d_array),
	1504: uint16(anon_sym_texture_depth_cube),
	1505: uint16(anon_sym_texture_depth_cube_array),
	1506: uint16(anon_sym_texture_depth_multisampled_2d),
	1507: uint16(anon_sym_texture_1d),
	1508: uint16(anon_sym_texture_2d),
	1509: uint16(anon_sym_texture_2d_array),
	1510: uint16(anon_sym_texture_3d),
	1511: uint16(anon_sym_texture_cube),
	1512: uint16(anon_sym_texture_cube_array),
	1513: uint16(anon_sym_texture_multisampled_2d),
	1514: uint16(anon_sym_texture_storage_1d),
	1515: uint16(anon_sym_texture_storage_2d),
	1516: uint16(anon_sym_texture_storage_2d_array),
	1517: uint16(anon_sym_texture_storage_3d),
	1518: uint16(anon_sym_vec2),
	1519: uint16(anon_sym_vec3),
	1520: uint16(anon_sym_vec4),
	1521: uint16(anon_sym_mat2x2),
	1522: uint16(anon_sym_mat2x3),
	1523: uint16(anon_sym_mat2x4),
	1524: uint16(anon_sym_mat3x2),
	1525: uint16(anon_sym_mat3x3),
	1526: uint16(anon_sym_mat3x4),
	1527: uint16(anon_sym_mat4x2),
	1528: uint16(anon_sym_mat4x3),
	1529: uint16(anon_sym_mat4x4),
	1530: uint16(3),
	1531: uint16(187),
	1532: uint16(1),
	1533: uint16(anon_sym_AT),
	1534: uint16(3),
	1535: uint16(2),
	1536: uint16(sym_block_comment),
	1537: uint16(sym_line_comment),
	1538: uint16(185),
	1539: uint16(41),
	1540: uint16(anon_sym_override),
	1541: uint16(anon_sym_fn),
	1542: uint16(sym_identifier),
	1543: uint16(anon_sym_var),
	1544: uint16(anon_sym_bool),
	1545: uint16(anon_sym_u32),
	1546: uint16(anon_sym_i32),
	1547: uint16(anon_sym_f32),
	1548: uint16(anon_sym_f16),
	1549: uint16(anon_sym_array),
	1550: uint16(anon_sym_ptr),
	1551: uint16(anon_sym_sampler),
	1552: uint16(anon_sym_sampler_comparison),
	1553: uint16(anon_sym_texture_depth_2d),
	1554: uint16(anon_sym_texture_depth_2d_array),
	1555: uint16(anon_sym_texture_depth_cube),
	1556: uint16(anon_sym_texture_depth_cube_array),
	1557: uint16(anon_sym_texture_depth_multisampled_2d),
	1558: uint16(anon_sym_texture_1d),
	1559: uint16(anon_sym_texture_2d),
	1560: uint16(anon_sym_texture_2d_array),
	1561: uint16(anon_sym_texture_3d),
	1562: uint16(anon_sym_texture_cube),
	1563: uint16(anon_sym_texture_cube_array),
	1564: uint16(anon_sym_texture_multisampled_2d),
	1565: uint16(anon_sym_texture_storage_1d),
	1566: uint16(anon_sym_texture_storage_2d),
	1567: uint16(anon_sym_texture_storage_2d_array),
	1568: uint16(anon_sym_texture_storage_3d),
	1569: uint16(anon_sym_vec2),
	1570: uint16(anon_sym_vec3),
	1571: uint16(anon_sym_vec4),
	1572: uint16(anon_sym_mat2x2),
	1573: uint16(anon_sym_mat2x3),
	1574: uint16(anon_sym_mat2x4),
	1575: uint16(anon_sym_mat3x2),
	1576: uint16(anon_sym_mat3x3),
	1577: uint16(anon_sym_mat3x4),
	1578: uint16(anon_sym_mat4x2),
	1579: uint16(anon_sym_mat4x3),
	1580: uint16(anon_sym_mat4x4),
	1581: uint16(3),
	1582: uint16(191),
	1583: uint16(1),
	1584: uint16(anon_sym_AT),
	1585: uint16(3),
	1586: uint16(2),
	1587: uint16(sym_block_comment),
	1588: uint16(sym_line_comment),
	1589: uint16(189),
	1590: uint16(41),
	1591: uint16(anon_sym_override),
	1592: uint16(anon_sym_fn),
	1593: uint16(sym_identifier),
	1594: uint16(anon_sym_var),
	1595: uint16(anon_sym_bool),
	1596: uint16(anon_sym_u32),
	1597: uint16(anon_sym_i32),
	1598: uint16(anon_sym_f32),
	1599: uint16(anon_sym_f16),
	1600: uint16(anon_sym_array),
	1601: uint16(anon_sym_ptr),
	1602: uint16(anon_sym_sampler),
	1603: uint16(anon_sym_sampler_comparison),
	1604: uint16(anon_sym_texture_depth_2d),
	1605: uint16(anon_sym_texture_depth_2d_array),
	1606: uint16(anon_sym_texture_depth_cube),
	1607: uint16(anon_sym_texture_depth_cube_array),
	1608: uint16(anon_sym_texture_depth_multisampled_2d),
	1609: uint16(anon_sym_texture_1d),
	1610: uint16(anon_sym_texture_2d),
	1611: uint16(anon_sym_texture_2d_array),
	1612: uint16(anon_sym_texture_3d),
	1613: uint16(anon_sym_texture_cube),
	1614: uint16(anon_sym_texture_cube_array),
	1615: uint16(anon_sym_texture_multisampled_2d),
	1616: uint16(anon_sym_texture_storage_1d),
	1617: uint16(anon_sym_texture_storage_2d),
	1618: uint16(anon_sym_texture_storage_2d_array),
	1619: uint16(anon_sym_texture_storage_3d),
	1620: uint16(anon_sym_vec2),
	1621: uint16(anon_sym_vec3),
	1622: uint16(anon_sym_vec4),
	1623: uint16(anon_sym_mat2x2),
	1624: uint16(anon_sym_mat2x3),
	1625: uint16(anon_sym_mat2x4),
	1626: uint16(anon_sym_mat3x2),
	1627: uint16(anon_sym_mat3x3),
	1628: uint16(anon_sym_mat3x4),
	1629: uint16(anon_sym_mat4x2),
	1630: uint16(anon_sym_mat4x3),
	1631: uint16(anon_sym_mat4x4),
	1632: uint16(3),
	1633: uint16(195),
	1634: uint16(1),
	1635: uint16(anon_sym_AT),
	1636: uint16(3),
	1637: uint16(2),
	1638: uint16(sym_block_comment),
	1639: uint16(sym_line_comment),
	1640: uint16(193),
	1641: uint16(41),
	1642: uint16(anon_sym_override),
	1643: uint16(anon_sym_fn),
	1644: uint16(sym_identifier),
	1645: uint16(anon_sym_var),
	1646: uint16(anon_sym_bool),
	1647: uint16(anon_sym_u32),
	1648: uint16(anon_sym_i32),
	1649: uint16(anon_sym_f32),
	1650: uint16(anon_sym_f16),
	1651: uint16(anon_sym_array),
	1652: uint16(anon_sym_ptr),
	1653: uint16(anon_sym_sampler),
	1654: uint16(anon_sym_sampler_comparison),
	1655: uint16(anon_sym_texture_depth_2d),
	1656: uint16(anon_sym_texture_depth_2d_array),
	1657: uint16(anon_sym_texture_depth_cube),
	1658: uint16(anon_sym_texture_depth_cube_array),
	1659: uint16(anon_sym_texture_depth_multisampled_2d),
	1660: uint16(anon_sym_texture_1d),
	1661: uint16(anon_sym_texture_2d),
	1662: uint16(anon_sym_texture_2d_array),
	1663: uint16(anon_sym_texture_3d),
	1664: uint16(anon_sym_texture_cube),
	1665: uint16(anon_sym_texture_cube_array),
	1666: uint16(anon_sym_texture_multisampled_2d),
	1667: uint16(anon_sym_texture_storage_1d),
	1668: uint16(anon_sym_texture_storage_2d),
	1669: uint16(anon_sym_texture_storage_2d_array),
	1670: uint16(anon_sym_texture_storage_3d),
	1671: uint16(anon_sym_vec2),
	1672: uint16(anon_sym_vec3),
	1673: uint16(anon_sym_vec4),
	1674: uint16(anon_sym_mat2x2),
	1675: uint16(anon_sym_mat2x3),
	1676: uint16(anon_sym_mat2x4),
	1677: uint16(anon_sym_mat3x2),
	1678: uint16(anon_sym_mat3x3),
	1679: uint16(anon_sym_mat3x4),
	1680: uint16(anon_sym_mat4x2),
	1681: uint16(anon_sym_mat4x3),
	1682: uint16(anon_sym_mat4x4),
	1683: uint16(9),
	1684: uint16(39),
	1685: uint16(1),
	1686: uint16(anon_sym_array),
	1687: uint16(41),
	1688: uint16(1),
	1689: uint16(anon_sym_ptr),
	1690: uint16(198),
	1691: uint16(1),
	1692: uint16(sym_type_declaration),
	1693: uint16(3),
	1694: uint16(2),
	1695: uint16(sym_block_comment),
	1696: uint16(sym_line_comment),
	1697: uint16(293),
	1698: uint16(2),
	1699: uint16(sym__vec_prefix),
	1700: uint16(sym__mat_prefix),
	1701: uint16(45),
	1702: uint16(4),
	1703: uint16(anon_sym_texture_storage_1d),
	1704: uint16(anon_sym_texture_storage_2d),
	1705: uint16(anon_sym_texture_storage_2d_array),
	1706: uint16(anon_sym_texture_storage_3d),
	1707: uint16(43),
	1708: uint16(7),
	1709: uint16(anon_sym_texture_1d),
	1710: uint16(anon_sym_texture_2d),
	1711: uint16(anon_sym_texture_2d_array),
	1712: uint16(anon_sym_texture_3d),
	1713: uint16(anon_sym_texture_cube),
	1714: uint16(anon_sym_texture_cube_array),
	1715: uint16(anon_sym_texture_multisampled_2d),
	1716: uint16(121),
	1717: uint16(12),
	1718: uint16(anon_sym_vec2),
	1719: uint16(anon_sym_vec3),
	1720: uint16(anon_sym_vec4),
	1721: uint16(anon_sym_mat2x2),
	1722: uint16(anon_sym_mat2x3),
	1723: uint16(anon_sym_mat2x4),
	1724: uint16(anon_sym_mat3x2),
	1725: uint16(anon_sym_mat3x3),
	1726: uint16(anon_sym_mat3x4),
	1727: uint16(anon_sym_mat4x2),
	1728: uint16(anon_sym_mat4x3),
	1729: uint16(anon_sym_mat4x4),
	1730: uint16(37),
	1731: uint16(13),
	1732: uint16(sym_identifier),
	1733: uint16(anon_sym_bool),
	1734: uint16(anon_sym_u32),
	1735: uint16(anon_sym_i32),
	1736: uint16(anon_sym_f32),
	1737: uint16(anon_sym_f16),
	1738: uint16(anon_sym_sampler),
	1739: uint16(anon_sym_sampler_comparison),
	1740: uint16(anon_sym_texture_depth_2d),
	1741: uint16(anon_sym_texture_depth_2d_array),
	1742: uint16(anon_sym_texture_depth_cube),
	1743: uint16(anon_sym_texture_depth_cube_array),
	1744: uint16(anon_sym_texture_depth_multisampled_2d),
	1745: uint16(9),
	1746: uint16(39),
	1747: uint16(1),
	1748: uint16(anon_sym_array),
	1749: uint16(41),
	1750: uint16(1),
	1751: uint16(anon_sym_ptr),
	1752: uint16(330),
	1753: uint16(1),
	1754: uint16(sym_type_declaration),
	1755: uint16(3),
	1756: uint16(2),
	1757: uint16(sym_block_comment),
	1758: uint16(sym_line_comment),
	1759: uint16(293),
	1760: uint16(2),
	1761: uint16(sym__vec_prefix),
	1762: uint16(sym__mat_prefix),
	1763: uint16(45),
	1764: uint16(4),
	1765: uint16(anon_sym_texture_storage_1d),
	1766: uint16(anon_sym_texture_storage_2d),
	1767: uint16(anon_sym_texture_storage_2d_array),
	1768: uint16(anon_sym_texture_storage_3d),
	1769: uint16(43),
	1770: uint16(7),
	1771: uint16(anon_sym_texture_1d),
	1772: uint16(anon_sym_texture_2d),
	1773: uint16(anon_sym_texture_2d_array),
	1774: uint16(anon_sym_texture_3d),
	1775: uint16(anon_sym_texture_cube),
	1776: uint16(anon_sym_texture_cube_array),
	1777: uint16(anon_sym_texture_multisampled_2d),
	1778: uint16(121),
	1779: uint16(12),
	1780: uint16(anon_sym_vec2),
	1781: uint16(anon_sym_vec3),
	1782: uint16(anon_sym_vec4),
	1783: uint16(anon_sym_mat2x2),
	1784: uint16(anon_sym_mat2x3),
	1785: uint16(anon_sym_mat2x4),
	1786: uint16(anon_sym_mat3x2),
	1787: uint16(anon_sym_mat3x3),
	1788: uint16(anon_sym_mat3x4),
	1789: uint16(anon_sym_mat4x2),
	1790: uint16(anon_sym_mat4x3),
	1791: uint16(anon_sym_mat4x4),
	1792: uint16(37),
	1793: uint16(13),
	1794: uint16(sym_identifier),
	1795: uint16(anon_sym_bool),
	1796: uint16(anon_sym_u32),
	1797: uint16(anon_sym_i32),
	1798: uint16(anon_sym_f32),
	1799: uint16(anon_sym_f16),
	1800: uint16(anon_sym_sampler),
	1801: uint16(anon_sym_sampler_comparison),
	1802: uint16(anon_sym_texture_depth_2d),
	1803: uint16(anon_sym_texture_depth_2d_array),
	1804: uint16(anon_sym_texture_depth_cube),
	1805: uint16(anon_sym_texture_depth_cube_array),
	1806: uint16(anon_sym_texture_depth_multisampled_2d),
	1807: uint16(9),
	1808: uint16(39),
	1809: uint16(1),
	1810: uint16(anon_sym_array),
	1811: uint16(41),
	1812: uint16(1),
	1813: uint16(anon_sym_ptr),
	1814: uint16(359),
	1815: uint16(1),
	1816: uint16(sym_type_declaration),
	1817: uint16(3),
	1818: uint16(2),
	1819: uint16(sym_block_comment),
	1820: uint16(sym_line_comment),
	1821: uint16(293),
	1822: uint16(2),
	1823: uint16(sym__vec_prefix),
	1824: uint16(sym__mat_prefix),
	1825: uint16(45),
	1826: uint16(4),
	1827: uint16(anon_sym_texture_storage_1d),
	1828: uint16(anon_sym_texture_storage_2d),
	1829: uint16(anon_sym_texture_storage_2d_array),
	1830: uint16(anon_sym_texture_storage_3d),
	1831: uint16(43),
	1832: uint16(7),
	1833: uint16(anon_sym_texture_1d),
	1834: uint16(anon_sym_texture_2d),
	1835: uint16(anon_sym_texture_2d_array),
	1836: uint16(anon_sym_texture_3d),
	1837: uint16(anon_sym_texture_cube),
	1838: uint16(anon_sym_texture_cube_array),
	1839: uint16(anon_sym_texture_multisampled_2d),
	1840: uint16(121),
	1841: uint16(12),
	1842: uint16(anon_sym_vec2),
	1843: uint16(anon_sym_vec3),
	1844: uint16(anon_sym_vec4),
	1845: uint16(anon_sym_mat2x2),
	1846: uint16(anon_sym_mat2x3),
	1847: uint16(anon_sym_mat2x4),
	1848: uint16(anon_sym_mat3x2),
	1849: uint16(anon_sym_mat3x3),
	1850: uint16(anon_sym_mat3x4),
	1851: uint16(anon_sym_mat4x2),
	1852: uint16(anon_sym_mat4x3),
	1853: uint16(anon_sym_mat4x4),
	1854: uint16(37),
	1855: uint16(13),
	1856: uint16(sym_identifier),
	1857: uint16(anon_sym_bool),
	1858: uint16(anon_sym_u32),
	1859: uint16(anon_sym_i32),
	1860: uint16(anon_sym_f32),
	1861: uint16(anon_sym_f16),
	1862: uint16(anon_sym_sampler),
	1863: uint16(anon_sym_sampler_comparison),
	1864: uint16(anon_sym_texture_depth_2d),
	1865: uint16(anon_sym_texture_depth_2d_array),
	1866: uint16(anon_sym_texture_depth_cube),
	1867: uint16(anon_sym_texture_depth_cube_array),
	1868: uint16(anon_sym_texture_depth_multisampled_2d),
	1869: uint16(9),
	1870: uint16(39),
	1871: uint16(1),
	1872: uint16(anon_sym_array),
	1873: uint16(41),
	1874: uint16(1),
	1875: uint16(anon_sym_ptr),
	1876: uint16(255),
	1877: uint16(1),
	1878: uint16(sym_type_declaration),
	1879: uint16(3),
	1880: uint16(2),
	1881: uint16(sym_block_comment),
	1882: uint16(sym_line_comment),
	1883: uint16(293),
	1884: uint16(2),
	1885: uint16(sym__vec_prefix),
	1886: uint16(sym__mat_prefix),
	1887: uint16(45),
	1888: uint16(4),
	1889: uint16(anon_sym_texture_storage_1d),
	1890: uint16(anon_sym_texture_storage_2d),
	1891: uint16(anon_sym_texture_storage_2d_array),
	1892: uint16(anon_sym_texture_storage_3d),
	1893: uint16(43),
	1894: uint16(7),
	1895: uint16(anon_sym_texture_1d),
	1896: uint16(anon_sym_texture_2d),
	1897: uint16(anon_sym_texture_2d_array),
	1898: uint16(anon_sym_texture_3d),
	1899: uint16(anon_sym_texture_cube),
	1900: uint16(anon_sym_texture_cube_array),
	1901: uint16(anon_sym_texture_multisampled_2d),
	1902: uint16(121),
	1903: uint16(12),
	1904: uint16(anon_sym_vec2),
	1905: uint16(anon_sym_vec3),
	1906: uint16(anon_sym_vec4),
	1907: uint16(anon_sym_mat2x2),
	1908: uint16(anon_sym_mat2x3),
	1909: uint16(anon_sym_mat2x4),
	1910: uint16(anon_sym_mat3x2),
	1911: uint16(anon_sym_mat3x3),
	1912: uint16(anon_sym_mat3x4),
	1913: uint16(anon_sym_mat4x2),
	1914: uint16(anon_sym_mat4x3),
	1915: uint16(anon_sym_mat4x4),
	1916: uint16(37),
	1917: uint16(13),
	1918: uint16(sym_identifier),
	1919: uint16(anon_sym_bool),
	1920: uint16(anon_sym_u32),
	1921: uint16(anon_sym_i32),
	1922: uint16(anon_sym_f32),
	1923: uint16(anon_sym_f16),
	1924: uint16(anon_sym_sampler),
	1925: uint16(anon_sym_sampler_comparison),
	1926: uint16(anon_sym_texture_depth_2d),
	1927: uint16(anon_sym_texture_depth_2d_array),
	1928: uint16(anon_sym_texture_depth_cube),
	1929: uint16(anon_sym_texture_depth_cube_array),
	1930: uint16(anon_sym_texture_depth_multisampled_2d),
	1931: uint16(9),
	1932: uint16(39),
	1933: uint16(1),
	1934: uint16(anon_sym_array),
	1935: uint16(41),
	1936: uint16(1),
	1937: uint16(anon_sym_ptr),
	1938: uint16(288),
	1939: uint16(1),
	1940: uint16(sym_type_declaration),
	1941: uint16(3),
	1942: uint16(2),
	1943: uint16(sym_block_comment),
	1944: uint16(sym_line_comment),
	1945: uint16(293),
	1946: uint16(2),
	1947: uint16(sym__vec_prefix),
	1948: uint16(sym__mat_prefix),
	1949: uint16(45),
	1950: uint16(4),
	1951: uint16(anon_sym_texture_storage_1d),
	1952: uint16(anon_sym_texture_storage_2d),
	1953: uint16(anon_sym_texture_storage_2d_array),
	1954: uint16(anon_sym_texture_storage_3d),
	1955: uint16(43),
	1956: uint16(7),
	1957: uint16(anon_sym_texture_1d),
	1958: uint16(anon_sym_texture_2d),
	1959: uint16(anon_sym_texture_2d_array),
	1960: uint16(anon_sym_texture_3d),
	1961: uint16(anon_sym_texture_cube),
	1962: uint16(anon_sym_texture_cube_array),
	1963: uint16(anon_sym_texture_multisampled_2d),
	1964: uint16(121),
	1965: uint16(12),
	1966: uint16(anon_sym_vec2),
	1967: uint16(anon_sym_vec3),
	1968: uint16(anon_sym_vec4),
	1969: uint16(anon_sym_mat2x2),
	1970: uint16(anon_sym_mat2x3),
	1971: uint16(anon_sym_mat2x4),
	1972: uint16(anon_sym_mat3x2),
	1973: uint16(anon_sym_mat3x3),
	1974: uint16(anon_sym_mat3x4),
	1975: uint16(anon_sym_mat4x2),
	1976: uint16(anon_sym_mat4x3),
	1977: uint16(anon_sym_mat4x4),
	1978: uint16(37),
	1979: uint16(13),
	1980: uint16(sym_identifier),
	1981: uint16(anon_sym_bool),
	1982: uint16(anon_sym_u32),
	1983: uint16(anon_sym_i32),
	1984: uint16(anon_sym_f32),
	1985: uint16(anon_sym_f16),
	1986: uint16(anon_sym_sampler),
	1987: uint16(anon_sym_sampler_comparison),
	1988: uint16(anon_sym_texture_depth_2d),
	1989: uint16(anon_sym_texture_depth_2d_array),
	1990: uint16(anon_sym_texture_depth_cube),
	1991: uint16(anon_sym_texture_depth_cube_array),
	1992: uint16(anon_sym_texture_depth_multisampled_2d),
	1993: uint16(9),
	1994: uint16(39),
	1995: uint16(1),
	1996: uint16(anon_sym_array),
	1997: uint16(41),
	1998: uint16(1),
	1999: uint16(anon_sym_ptr),
	2000: uint16(240),
	2001: uint16(1),
	2002: uint16(sym_type_declaration),
	2003: uint16(3),
	2004: uint16(2),
	2005: uint16(sym_block_comment),
	2006: uint16(sym_line_comment),
	2007: uint16(293),
	2008: uint16(2),
	2009: uint16(sym__vec_prefix),
	2010: uint16(sym__mat_prefix),
	2011: uint16(45),
	2012: uint16(4),
	2013: uint16(anon_sym_texture_storage_1d),
	2014: uint16(anon_sym_texture_storage_2d),
	2015: uint16(anon_sym_texture_storage_2d_array),
	2016: uint16(anon_sym_texture_storage_3d),
	2017: uint16(43),
	2018: uint16(7),
	2019: uint16(anon_sym_texture_1d),
	2020: uint16(anon_sym_texture_2d),
	2021: uint16(anon_sym_texture_2d_array),
	2022: uint16(anon_sym_texture_3d),
	2023: uint16(anon_sym_texture_cube),
	2024: uint16(anon_sym_texture_cube_array),
	2025: uint16(anon_sym_texture_multisampled_2d),
	2026: uint16(121),
	2027: uint16(12),
	2028: uint16(anon_sym_vec2),
	2029: uint16(anon_sym_vec3),
	2030: uint16(anon_sym_vec4),
	2031: uint16(anon_sym_mat2x2),
	2032: uint16(anon_sym_mat2x3),
	2033: uint16(anon_sym_mat2x4),
	2034: uint16(anon_sym_mat3x2),
	2035: uint16(anon_sym_mat3x3),
	2036: uint16(anon_sym_mat3x4),
	2037: uint16(anon_sym_mat4x2),
	2038: uint16(anon_sym_mat4x3),
	2039: uint16(anon_sym_mat4x4),
	2040: uint16(37),
	2041: uint16(13),
	2042: uint16(sym_identifier),
	2043: uint16(anon_sym_bool),
	2044: uint16(anon_sym_u32),
	2045: uint16(anon_sym_i32),
	2046: uint16(anon_sym_f32),
	2047: uint16(anon_sym_f16),
	2048: uint16(anon_sym_sampler),
	2049: uint16(anon_sym_sampler_comparison),
	2050: uint16(anon_sym_texture_depth_2d),
	2051: uint16(anon_sym_texture_depth_2d_array),
	2052: uint16(anon_sym_texture_depth_cube),
	2053: uint16(anon_sym_texture_depth_cube_array),
	2054: uint16(anon_sym_texture_depth_multisampled_2d),
	2055: uint16(25),
	2056: uint16(107),
	2057: uint16(1),
	2058: uint16(anon_sym_let),
	2059: uint16(109),
	2060: uint16(1),
	2061: uint16(anon_sym_LPAREN),
	2062: uint16(111),
	2063: uint16(1),
	2064: uint16(anon_sym__),
	2065: uint16(113),
	2066: uint16(1),
	2067: uint16(anon_sym_var),
	2068: uint16(197),
	2069: uint16(1),
	2070: uint16(sym_identifier),
	2071: uint16(199),
	2072: uint16(1),
	2073: uint16(anon_sym_LBRACE),
	2074: uint16(201),
	2075: uint16(1),
	2076: uint16(anon_sym_RBRACE),
	2077: uint16(203),
	2078: uint16(1),
	2079: uint16(anon_sym_if),
	2080: uint16(205),
	2081: uint16(1),
	2082: uint16(anon_sym_switch),
	2083: uint16(207),
	2084: uint16(1),
	2085: uint16(anon_sym_fallthrough),
	2086: uint16(209),
	2087: uint16(1),
	2088: uint16(anon_sym_loop),
	2089: uint16(211),
	2090: uint16(1),
	2091: uint16(anon_sym_for),
	2092: uint16(213),
	2093: uint16(1),
	2094: uint16(anon_sym_while),
	2095: uint16(215),
	2096: uint16(1),
	2097: uint16(anon_sym_break),
	2098: uint16(217),
	2099: uint16(1),
	2100: uint16(anon_sym_continue),
	2101: uint16(219),
	2102: uint16(1),
	2103: uint16(anon_sym_return),
	2104: uint16(221),
	2105: uint16(1),
	2106: uint16(anon_sym_discard),
	2107: uint16(154),
	2108: uint16(1),
	2109: uint16(sym_lhs_expression),
	2110: uint16(202),
	2111: uint16(1),
	2112: uint16(aux_sym_lhs_expression_repeat1),
	2113: uint16(246),
	2114: uint16(1),
	2115: uint16(sym_variable_declaration),
	2116: uint16(333),
	2117: uint16(1),
	2118: uint16(sym_fallthrough_statement),
	2119: uint16(3),
	2120: uint16(2),
	2121: uint16(sym_block_comment),
	2122: uint16(sym_line_comment),
	2123: uint16(115),
	2124: uint16(2),
	2125: uint16(anon_sym_AMP),
	2126: uint16(anon_sym_STAR),
	2127: uint16(340),
	2128: uint16(3),
	2129: uint16(sym_assignment_statement),
	2130: uint16(sym_return_statement),
	2131: uint16(sym_variable_statement),
	2132: uint16(66),
	2133: uint16(13),
	2134: uint16(sym__statement),
	2135: uint16(sym_compound_statement),
	2136: uint16(sym_if_statement),
	2137: uint16(sym_switch_statement),
	2138: uint16(sym_loop_statement),
	2139: uint16(sym_for_statement),
	2140: uint16(sym_while_statement),
	2141: uint16(sym_break_statement),
	2142: uint16(sym_continue_statement),
	2143: uint16(sym_discard_statement),
	2144: uint16(sym_increment_statement),
	2145: uint16(sym_decrement_statement),
	2146: uint16(aux_sym_compound_statement_repeat1),
	2147: uint16(25),
	2148: uint16(107),
	2149: uint16(1),
	2150: uint16(anon_sym_let),
	2151: uint16(109),
	2152: uint16(1),
	2153: uint16(anon_sym_LPAREN),
	2154: uint16(111),
	2155: uint16(1),
	2156: uint16(anon_sym__),
	2157: uint16(113),
	2158: uint16(1),
	2159: uint16(anon_sym_var),
	2160: uint16(197),
	2161: uint16(1),
	2162: uint16(sym_identifier),
	2163: uint16(199),
	2164: uint16(1),
	2165: uint16(anon_sym_LBRACE),
	2166: uint16(203),
	2167: uint16(1),
	2168: uint16(anon_sym_if),
	2169: uint16(205),
	2170: uint16(1),
	2171: uint16(anon_sym_switch),
	2172: uint16(209),
	2173: uint16(1),
	2174: uint16(anon_sym_loop),
	2175: uint16(211),
	2176: uint16(1),
	2177: uint16(anon_sym_for),
	2178: uint16(213),
	2179: uint16(1),
	2180: uint16(anon_sym_while),
	2181: uint16(215),
	2182: uint16(1),
	2183: uint16(anon_sym_break),
	2184: uint16(217),
	2185: uint16(1),
	2186: uint16(anon_sym_continue),
	2187: uint16(219),
	2188: uint16(1),
	2189: uint16(anon_sym_return),
	2190: uint16(221),
	2191: uint16(1),
	2192: uint16(anon_sym_discard),
	2193: uint16(223),
	2194: uint16(1),
	2195: uint16(anon_sym_RBRACE),
	2196: uint16(225),
	2197: uint16(1),
	2198: uint16(anon_sym_continuing),
	2199: uint16(154),
	2200: uint16(1),
	2201: uint16(sym_lhs_expression),
	2202: uint16(202),
	2203: uint16(1),
	2204: uint16(aux_sym_lhs_expression_repeat1),
	2205: uint16(246),
	2206: uint16(1),
	2207: uint16(sym_variable_declaration),
	2208: uint16(298),
	2209: uint16(1),
	2210: uint16(sym_continuing_statement),
	2211: uint16(3),
	2212: uint16(2),
	2213: uint16(sym_block_comment),
	2214: uint16(sym_line_comment),
	2215: uint16(115),
	2216: uint16(2),
	2217: uint16(anon_sym_AMP),
	2218: uint16(anon_sym_STAR),
	2219: uint16(340),
	2220: uint16(3),
	2221: uint16(sym_assignment_statement),
	2222: uint16(sym_return_statement),
	2223: uint16(sym_variable_statement),
	2224: uint16(64),
	2225: uint16(13),
	2226: uint16(sym__statement),
	2227: uint16(sym_compound_statement),
	2228: uint16(sym_if_statement),
	2229: uint16(sym_switch_statement),
	2230: uint16(sym_loop_statement),
	2231: uint16(sym_for_statement),
	2232: uint16(sym_while_statement),
	2233: uint16(sym_break_statement),
	2234: uint16(sym_continue_statement),
	2235: uint16(sym_discard_statement),
	2236: uint16(sym_increment_statement),
	2237: uint16(sym_decrement_statement),
	2238: uint16(aux_sym_compound_statement_repeat1),
	2239: uint16(24),
	2240: uint16(227),
	2241: uint16(1),
	2242: uint16(sym_identifier),
	2243: uint16(230),
	2244: uint16(1),
	2245: uint16(anon_sym_let),
	2246: uint16(233),
	2247: uint16(1),
	2248: uint16(anon_sym_LPAREN),
	2249: uint16(236),
	2250: uint16(1),
	2251: uint16(anon_sym_LBRACE),
	2252: uint16(239),
	2253: uint16(1),
	2254: uint16(anon_sym_RBRACE),
	2255: uint16(241),
	2256: uint16(1),
	2257: uint16(anon_sym__),
	2258: uint16(244),
	2259: uint16(1),
	2260: uint16(anon_sym_if),
	2261: uint16(247),
	2262: uint16(1),
	2263: uint16(anon_sym_switch),
	2264: uint16(252),
	2265: uint16(1),
	2266: uint16(anon_sym_loop),
	2267: uint16(255),
	2268: uint16(1),
	2269: uint16(anon_sym_for),
	2270: uint16(258),
	2271: uint16(1),
	2272: uint16(anon_sym_while),
	2273: uint16(261),
	2274: uint16(1),
	2275: uint16(anon_sym_break),
	2276: uint16(264),
	2277: uint16(1),
	2278: uint16(anon_sym_continue),
	2279: uint16(267),
	2280: uint16(1),
	2281: uint16(anon_sym_return),
	2282: uint16(270),
	2283: uint16(1),
	2284: uint16(anon_sym_discard),
	2285: uint16(273),
	2286: uint16(1),
	2287: uint16(anon_sym_var),
	2288: uint16(154),
	2289: uint16(1),
	2290: uint16(sym_lhs_expression),
	2291: uint16(202),
	2292: uint16(1),
	2293: uint16(aux_sym_lhs_expression_repeat1),
	2294: uint16(246),
	2295: uint16(1),
	2296: uint16(sym_variable_declaration),
	2297: uint16(3),
	2298: uint16(2),
	2299: uint16(sym_block_comment),
	2300: uint16(sym_line_comment),
	2301: uint16(250),
	2302: uint16(2),
	2303: uint16(anon_sym_fallthrough),
	2304: uint16(anon_sym_continuing),
	2305: uint16(276),
	2306: uint16(2),
	2307: uint16(anon_sym_AMP),
	2308: uint16(anon_sym_STAR),
	2309: uint16(340),
	2310: uint16(3),
	2311: uint16(sym_assignment_statement),
	2312: uint16(sym_return_statement),
	2313: uint16(sym_variable_statement),
	2314: uint16(64),
	2315: uint16(13),
	2316: uint16(sym__statement),
	2317: uint16(sym_compound_statement),
	2318: uint16(sym_if_statement),
	2319: uint16(sym_switch_statement),
	2320: uint16(sym_loop_statement),
	2321: uint16(sym_for_statement),
	2322: uint16(sym_while_statement),
	2323: uint16(sym_break_statement),
	2324: uint16(sym_continue_statement),
	2325: uint16(sym_discard_statement),
	2326: uint16(sym_increment_statement),
	2327: uint16(sym_decrement_statement),
	2328: uint16(aux_sym_compound_statement_repeat1),
	2329: uint16(25),
	2330: uint16(107),
	2331: uint16(1),
	2332: uint16(anon_sym_let),
	2333: uint16(109),
	2334: uint16(1),
	2335: uint16(anon_sym_LPAREN),
	2336: uint16(111),
	2337: uint16(1),
	2338: uint16(anon_sym__),
	2339: uint16(113),
	2340: uint16(1),
	2341: uint16(anon_sym_var),
	2342: uint16(197),
	2343: uint16(1),
	2344: uint16(sym_identifier),
	2345: uint16(199),
	2346: uint16(1),
	2347: uint16(anon_sym_LBRACE),
	2348: uint16(203),
	2349: uint16(1),
	2350: uint16(anon_sym_if),
	2351: uint16(205),
	2352: uint16(1),
	2353: uint16(anon_sym_switch),
	2354: uint16(209),
	2355: uint16(1),
	2356: uint16(anon_sym_loop),
	2357: uint16(211),
	2358: uint16(1),
	2359: uint16(anon_sym_for),
	2360: uint16(213),
	2361: uint16(1),
	2362: uint16(anon_sym_while),
	2363: uint16(215),
	2364: uint16(1),
	2365: uint16(anon_sym_break),
	2366: uint16(217),
	2367: uint16(1),
	2368: uint16(anon_sym_continue),
	2369: uint16(219),
	2370: uint16(1),
	2371: uint16(anon_sym_return),
	2372: uint16(221),
	2373: uint16(1),
	2374: uint16(anon_sym_discard),
	2375: uint16(225),
	2376: uint16(1),
	2377: uint16(anon_sym_continuing),
	2378: uint16(279),
	2379: uint16(1),
	2380: uint16(anon_sym_RBRACE),
	2381: uint16(154),
	2382: uint16(1),
	2383: uint16(sym_lhs_expression),
	2384: uint16(202),
	2385: uint16(1),
	2386: uint16(aux_sym_lhs_expression_repeat1),
	2387: uint16(246),
	2388: uint16(1),
	2389: uint16(sym_variable_declaration),
	2390: uint16(309),
	2391: uint16(1),
	2392: uint16(sym_continuing_statement),
	2393: uint16(3),
	2394: uint16(2),
	2395: uint16(sym_block_comment),
	2396: uint16(sym_line_comment),
	2397: uint16(115),
	2398: uint16(2),
	2399: uint16(anon_sym_AMP),
	2400: uint16(anon_sym_STAR),
	2401: uint16(340),
	2402: uint16(3),
	2403: uint16(sym_assignment_statement),
	2404: uint16(sym_return_statement),
	2405: uint16(sym_variable_statement),
	2406: uint16(63),
	2407: uint16(13),
	2408: uint16(sym__statement),
	2409: uint16(sym_compound_statement),
	2410: uint16(sym_if_statement),
	2411: uint16(sym_switch_statement),
	2412: uint16(sym_loop_statement),
	2413: uint16(sym_for_statement),
	2414: uint16(sym_while_statement),
	2415: uint16(sym_break_statement),
	2416: uint16(sym_continue_statement),
	2417: uint16(sym_discard_statement),
	2418: uint16(sym_increment_statement),
	2419: uint16(sym_decrement_statement),
	2420: uint16(aux_sym_compound_statement_repeat1),
	2421: uint16(25),
	2422: uint16(107),
	2423: uint16(1),
	2424: uint16(anon_sym_let),
	2425: uint16(109),
	2426: uint16(1),
	2427: uint16(anon_sym_LPAREN),
	2428: uint16(111),
	2429: uint16(1),
	2430: uint16(anon_sym__),
	2431: uint16(113),
	2432: uint16(1),
	2433: uint16(anon_sym_var),
	2434: uint16(197),
	2435: uint16(1),
	2436: uint16(sym_identifier),
	2437: uint16(199),
	2438: uint16(1),
	2439: uint16(anon_sym_LBRACE),
	2440: uint16(203),
	2441: uint16(1),
	2442: uint16(anon_sym_if),
	2443: uint16(205),
	2444: uint16(1),
	2445: uint16(anon_sym_switch),
	2446: uint16(207),
	2447: uint16(1),
	2448: uint16(anon_sym_fallthrough),
	2449: uint16(209),
	2450: uint16(1),
	2451: uint16(anon_sym_loop),
	2452: uint16(211),
	2453: uint16(1),
	2454: uint16(anon_sym_for),
	2455: uint16(213),
	2456: uint16(1),
	2457: uint16(anon_sym_while),
	2458: uint16(215),
	2459: uint16(1),
	2460: uint16(anon_sym_break),
	2461: uint16(217),
	2462: uint16(1),
	2463: uint16(anon_sym_continue),
	2464: uint16(219),
	2465: uint16(1),
	2466: uint16(anon_sym_return),
	2467: uint16(221),
	2468: uint16(1),
	2469: uint16(anon_sym_discard),
	2470: uint16(281),
	2471: uint16(1),
	2472: uint16(anon_sym_RBRACE),
	2473: uint16(154),
	2474: uint16(1),
	2475: uint16(sym_lhs_expression),
	2476: uint16(202),
	2477: uint16(1),
	2478: uint16(aux_sym_lhs_expression_repeat1),
	2479: uint16(246),
	2480: uint16(1),
	2481: uint16(sym_variable_declaration),
	2482: uint16(346),
	2483: uint16(1),
	2484: uint16(sym_fallthrough_statement),
	2485: uint16(3),
	2486: uint16(2),
	2487: uint16(sym_block_comment),
	2488: uint16(sym_line_comment),
	2489: uint16(115),
	2490: uint16(2),
	2491: uint16(anon_sym_AMP),
	2492: uint16(anon_sym_STAR),
	2493: uint16(340),
	2494: uint16(3),
	2495: uint16(sym_assignment_statement),
	2496: uint16(sym_return_statement),
	2497: uint16(sym_variable_statement),
	2498: uint16(64),
	2499: uint16(13),
	2500: uint16(sym__statement),
	2501: uint16(sym_compound_statement),
	2502: uint16(sym_if_statement),
	2503: uint16(sym_switch_statement),
	2504: uint16(sym_loop_statement),
	2505: uint16(sym_for_statement),
	2506: uint16(sym_while_statement),
	2507: uint16(sym_break_statement),
	2508: uint16(sym_continue_statement),
	2509: uint16(sym_discard_statement),
	2510: uint16(sym_increment_statement),
	2511: uint16(sym_decrement_statement),
	2512: uint16(aux_sym_compound_statement_repeat1),
	2513: uint16(24),
	2514: uint16(107),
	2515: uint16(1),
	2516: uint16(anon_sym_let),
	2517: uint16(109),
	2518: uint16(1),
	2519: uint16(anon_sym_LPAREN),
	2520: uint16(111),
	2521: uint16(1),
	2522: uint16(anon_sym__),
	2523: uint16(113),
	2524: uint16(1),
	2525: uint16(anon_sym_var),
	2526: uint16(197),
	2527: uint16(1),
	2528: uint16(sym_identifier),
	2529: uint16(199),
	2530: uint16(1),
	2531: uint16(anon_sym_LBRACE),
	2532: uint16(203),
	2533: uint16(1),
	2534: uint16(anon_sym_if),
	2535: uint16(205),
	2536: uint16(1),
	2537: uint16(anon_sym_switch),
	2538: uint16(209),
	2539: uint16(1),
	2540: uint16(anon_sym_loop),
	2541: uint16(211),
	2542: uint16(1),
	2543: uint16(anon_sym_for),
	2544: uint16(213),
	2545: uint16(1),
	2546: uint16(anon_sym_while),
	2547: uint16(217),
	2548: uint16(1),
	2549: uint16(anon_sym_continue),
	2550: uint16(219),
	2551: uint16(1),
	2552: uint16(anon_sym_return),
	2553: uint16(221),
	2554: uint16(1),
	2555: uint16(anon_sym_discard),
	2556: uint16(283),
	2557: uint16(1),
	2558: uint16(anon_sym_RBRACE),
	2559: uint16(285),
	2560: uint16(1),
	2561: uint16(anon_sym_break),
	2562: uint16(154),
	2563: uint16(1),
	2564: uint16(sym_lhs_expression),
	2565: uint16(202),
	2566: uint16(1),
	2567: uint16(aux_sym_lhs_expression_repeat1),
	2568: uint16(246),
	2569: uint16(1),
	2570: uint16(sym_variable_declaration),
	2571: uint16(337),
	2572: uint16(1),
	2573: uint16(sym_break_if_statement),
	2574: uint16(3),
	2575: uint16(2),
	2576: uint16(sym_block_comment),
	2577: uint16(sym_line_comment),
	2578: uint16(115),
	2579: uint16(2),
	2580: uint16(anon_sym_AMP),
	2581: uint16(anon_sym_STAR),
	2582: uint16(340),
	2583: uint16(3),
	2584: uint16(sym_assignment_statement),
	2585: uint16(sym_return_statement),
	2586: uint16(sym_variable_statement),
	2587: uint16(64),
	2588: uint16(13),
	2589: uint16(sym__statement),
	2590: uint16(sym_compound_statement),
	2591: uint16(sym_if_statement),
	2592: uint16(sym_switch_statement),
	2593: uint16(sym_loop_statement),
	2594: uint16(sym_for_statement),
	2595: uint16(sym_while_statement),
	2596: uint16(sym_break_statement),
	2597: uint16(sym_continue_statement),
	2598: uint16(sym_discard_statement),
	2599: uint16(sym_increment_statement),
	2600: uint16(sym_decrement_statement),
	2601: uint16(aux_sym_compound_statement_repeat1),
	2602: uint16(24),
	2603: uint16(107),
	2604: uint16(1),
	2605: uint16(anon_sym_let),
	2606: uint16(109),
	2607: uint16(1),
	2608: uint16(anon_sym_LPAREN),
	2609: uint16(111),
	2610: uint16(1),
	2611: uint16(anon_sym__),
	2612: uint16(113),
	2613: uint16(1),
	2614: uint16(anon_sym_var),
	2615: uint16(197),
	2616: uint16(1),
	2617: uint16(sym_identifier),
	2618: uint16(199),
	2619: uint16(1),
	2620: uint16(anon_sym_LBRACE),
	2621: uint16(203),
	2622: uint16(1),
	2623: uint16(anon_sym_if),
	2624: uint16(205),
	2625: uint16(1),
	2626: uint16(anon_sym_switch),
	2627: uint16(209),
	2628: uint16(1),
	2629: uint16(anon_sym_loop),
	2630: uint16(211),
	2631: uint16(1),
	2632: uint16(anon_sym_for),
	2633: uint16(213),
	2634: uint16(1),
	2635: uint16(anon_sym_while),
	2636: uint16(217),
	2637: uint16(1),
	2638: uint16(anon_sym_continue),
	2639: uint16(219),
	2640: uint16(1),
	2641: uint16(anon_sym_return),
	2642: uint16(221),
	2643: uint16(1),
	2644: uint16(anon_sym_discard),
	2645: uint16(285),
	2646: uint16(1),
	2647: uint16(anon_sym_break),
	2648: uint16(287),
	2649: uint16(1),
	2650: uint16(anon_sym_RBRACE),
	2651: uint16(154),
	2652: uint16(1),
	2653: uint16(sym_lhs_expression),
	2654: uint16(202),
	2655: uint16(1),
	2656: uint16(aux_sym_lhs_expression_repeat1),
	2657: uint16(246),
	2658: uint16(1),
	2659: uint16(sym_variable_declaration),
	2660: uint16(317),
	2661: uint16(1),
	2662: uint16(sym_break_if_statement),
	2663: uint16(3),
	2664: uint16(2),
	2665: uint16(sym_block_comment),
	2666: uint16(sym_line_comment),
	2667: uint16(115),
	2668: uint16(2),
	2669: uint16(anon_sym_AMP),
	2670: uint16(anon_sym_STAR),
	2671: uint16(340),
	2672: uint16(3),
	2673: uint16(sym_assignment_statement),
	2674: uint16(sym_return_statement),
	2675: uint16(sym_variable_statement),
	2676: uint16(67),
	2677: uint16(13),
	2678: uint16(sym__statement),
	2679: uint16(sym_compound_statement),
	2680: uint16(sym_if_statement),
	2681: uint16(sym_switch_statement),
	2682: uint16(sym_loop_statement),
	2683: uint16(sym_for_statement),
	2684: uint16(sym_while_statement),
	2685: uint16(sym_break_statement),
	2686: uint16(sym_continue_statement),
	2687: uint16(sym_discard_statement),
	2688: uint16(sym_increment_statement),
	2689: uint16(sym_decrement_statement),
	2690: uint16(aux_sym_compound_statement_repeat1),
	2691: uint16(23),
	2692: uint16(107),
	2693: uint16(1),
	2694: uint16(anon_sym_let),
	2695: uint16(109),
	2696: uint16(1),
	2697: uint16(anon_sym_LPAREN),
	2698: uint16(111),
	2699: uint16(1),
	2700: uint16(anon_sym__),
	2701: uint16(113),
	2702: uint16(1),
	2703: uint16(anon_sym_var),
	2704: uint16(197),
	2705: uint16(1),
	2706: uint16(sym_identifier),
	2707: uint16(199),
	2708: uint16(1),
	2709: uint16(anon_sym_LBRACE),
	2710: uint16(203),
	2711: uint16(1),
	2712: uint16(anon_sym_if),
	2713: uint16(205),
	2714: uint16(1),
	2715: uint16(anon_sym_switch),
	2716: uint16(209),
	2717: uint16(1),
	2718: uint16(anon_sym_loop),
	2719: uint16(211),
	2720: uint16(1),
	2721: uint16(anon_sym_for),
	2722: uint16(213),
	2723: uint16(1),
	2724: uint16(anon_sym_while),
	2725: uint16(215),
	2726: uint16(1),
	2727: uint16(anon_sym_break),
	2728: uint16(217),
	2729: uint16(1),
	2730: uint16(anon_sym_continue),
	2731: uint16(219),
	2732: uint16(1),
	2733: uint16(anon_sym_return),
	2734: uint16(221),
	2735: uint16(1),
	2736: uint16(anon_sym_discard),
	2737: uint16(289),
	2738: uint16(1),
	2739: uint16(anon_sym_RBRACE),
	2740: uint16(154),
	2741: uint16(1),
	2742: uint16(sym_lhs_expression),
	2743: uint16(202),
	2744: uint16(1),
	2745: uint16(aux_sym_lhs_expression_repeat1),
	2746: uint16(246),
	2747: uint16(1),
	2748: uint16(sym_variable_declaration),
	2749: uint16(3),
	2750: uint16(2),
	2751: uint16(sym_block_comment),
	2752: uint16(sym_line_comment),
	2753: uint16(115),
	2754: uint16(2),
	2755: uint16(anon_sym_AMP),
	2756: uint16(anon_sym_STAR),
	2757: uint16(340),
	2758: uint16(3),
	2759: uint16(sym_assignment_statement),
	2760: uint16(sym_return_statement),
	2761: uint16(sym_variable_statement),
	2762: uint16(70),
	2763: uint16(13),
	2764: uint16(sym__statement),
	2765: uint16(sym_compound_statement),
	2766: uint16(sym_if_statement),
	2767: uint16(sym_switch_statement),
	2768: uint16(sym_loop_statement),
	2769: uint16(sym_for_statement),
	2770: uint16(sym_while_statement),
	2771: uint16(sym_break_statement),
	2772: uint16(sym_continue_statement),
	2773: uint16(sym_discard_statement),
	2774: uint16(sym_increment_statement),
	2775: uint16(sym_decrement_statement),
	2776: uint16(aux_sym_compound_statement_repeat1),
	2777: uint16(23),
	2778: uint16(107),
	2779: uint16(1),
	2780: uint16(anon_sym_let),
	2781: uint16(109),
	2782: uint16(1),
	2783: uint16(anon_sym_LPAREN),
	2784: uint16(111),
	2785: uint16(1),
	2786: uint16(anon_sym__),
	2787: uint16(113),
	2788: uint16(1),
	2789: uint16(anon_sym_var),
	2790: uint16(197),
	2791: uint16(1),
	2792: uint16(sym_identifier),
	2793: uint16(199),
	2794: uint16(1),
	2795: uint16(anon_sym_LBRACE),
	2796: uint16(203),
	2797: uint16(1),
	2798: uint16(anon_sym_if),
	2799: uint16(205),
	2800: uint16(1),
	2801: uint16(anon_sym_switch),
	2802: uint16(209),
	2803: uint16(1),
	2804: uint16(anon_sym_loop),
	2805: uint16(211),
	2806: uint16(1),
	2807: uint16(anon_sym_for),
	2808: uint16(213),
	2809: uint16(1),
	2810: uint16(anon_sym_while),
	2811: uint16(215),
	2812: uint16(1),
	2813: uint16(anon_sym_break),
	2814: uint16(217),
	2815: uint16(1),
	2816: uint16(anon_sym_continue),
	2817: uint16(219),
	2818: uint16(1),
	2819: uint16(anon_sym_return),
	2820: uint16(221),
	2821: uint16(1),
	2822: uint16(anon_sym_discard),
	2823: uint16(291),
	2824: uint16(1),
	2825: uint16(anon_sym_RBRACE),
	2826: uint16(154),
	2827: uint16(1),
	2828: uint16(sym_lhs_expression),
	2829: uint16(202),
	2830: uint16(1),
	2831: uint16(aux_sym_lhs_expression_repeat1),
	2832: uint16(246),
	2833: uint16(1),
	2834: uint16(sym_variable_declaration),
	2835: uint16(3),
	2836: uint16(2),
	2837: uint16(sym_block_comment),
	2838: uint16(sym_line_comment),
	2839: uint16(115),
	2840: uint16(2),
	2841: uint16(anon_sym_AMP),
	2842: uint16(anon_sym_STAR),
	2843: uint16(340),
	2844: uint16(3),
	2845: uint16(sym_assignment_statement),
	2846: uint16(sym_return_statement),
	2847: uint16(sym_variable_statement),
	2848: uint16(64),
	2849: uint16(13),
	2850: uint16(sym__statement),
	2851: uint16(sym_compound_statement),
	2852: uint16(sym_if_statement),
	2853: uint16(sym_switch_statement),
	2854: uint16(sym_loop_statement),
	2855: uint16(sym_for_statement),
	2856: uint16(sym_while_statement),
	2857: uint16(sym_break_statement),
	2858: uint16(sym_continue_statement),
	2859: uint16(sym_discard_statement),
	2860: uint16(sym_increment_statement),
	2861: uint16(sym_decrement_statement),
	2862: uint16(aux_sym_compound_statement_repeat1),
	2863: uint16(3),
	2864: uint16(3),
	2865: uint16(2),
	2866: uint16(sym_block_comment),
	2867: uint16(sym_line_comment),
	2868: uint16(293),
	2869: uint16(8),
	2871: uint16(anon_sym_SEMI),
	2872: uint16(anon_sym_LPAREN),
	2873: uint16(anon_sym_LBRACE),
	2874: uint16(anon_sym_RBRACE),
	2875: uint16(anon_sym_AT),
	2876: uint16(anon_sym_AMP),
	2877: uint16(anon_sym_STAR),
	2878: uint16(295),
	2879: uint16(20),
	2880: uint16(anon_sym_let),
	2881: uint16(anon_sym_override),
	2882: uint16(anon_sym_type),
	2883: uint16(anon_sym_fn),
	2884: uint16(anon_sym_struct),
	2885: uint16(sym_identifier),
	2886: uint16(anon_sym__),
	2887: uint16(anon_sym_if),
	2888: uint16(anon_sym_else),
	2889: uint16(anon_sym_switch),
	2890: uint16(anon_sym_fallthrough),
	2891: uint16(anon_sym_loop),
	2892: uint16(anon_sym_for),
	2893: uint16(anon_sym_while),
	2894: uint16(anon_sym_break),
	2895: uint16(anon_sym_continue),
	2896: uint16(anon_sym_continuing),
	2897: uint16(anon_sym_return),
	2898: uint16(anon_sym_discard),
	2899: uint16(anon_sym_var),
	2900: uint16(3),
	2901: uint16(3),
	2902: uint16(2),
	2903: uint16(sym_block_comment),
	2904: uint16(sym_line_comment),
	2905: uint16(297),
	2906: uint16(8),
	2908: uint16(anon_sym_SEMI),
	2909: uint16(anon_sym_LPAREN),
	2910: uint16(anon_sym_LBRACE),
	2911: uint16(anon_sym_RBRACE),
	2912: uint16(anon_sym_AT),
	2913: uint16(anon_sym_AMP),
	2914: uint16(anon_sym_STAR),
	2915: uint16(299),
	2916: uint16(20),
	2917: uint16(anon_sym_let),
	2918: uint16(anon_sym_override),
	2919: uint16(anon_sym_type),
	2920: uint16(anon_sym_fn),
	2921: uint16(anon_sym_struct),
	2922: uint16(sym_identifier),
	2923: uint16(anon_sym__),
	2924: uint16(anon_sym_if),
	2925: uint16(anon_sym_else),
	2926: uint16(anon_sym_switch),
	2927: uint16(anon_sym_fallthrough),
	2928: uint16(anon_sym_loop),
	2929: uint16(anon_sym_for),
	2930: uint16(anon_sym_while),
	2931: uint16(anon_sym_break),
	2932: uint16(anon_sym_continue),
	2933: uint16(anon_sym_continuing),
	2934: uint16(anon_sym_return),
	2935: uint16(anon_sym_discard),
	2936: uint16(anon_sym_var),
	2937: uint16(3),
	2938: uint16(3),
	2939: uint16(2),
	2940: uint16(sym_block_comment),
	2941: uint16(sym_line_comment),
	2942: uint16(303),
	2943: uint16(5),
	2944: uint16(anon_sym_LT),
	2945: uint16(anon_sym_GT),
	2946: uint16(anon_sym_PIPE),
	2947: uint16(anon_sym_AMP),
	2948: uint16(anon_sym_SLASH),
	2949: uint16(301),
	2950: uint16(21),
	2951: uint16(anon_sym_SEMI),
	2952: uint16(anon_sym_COMMA),
	2953: uint16(anon_sym_RPAREN),
	2954: uint16(anon_sym_LBRACE),
	2955: uint16(anon_sym_COLON),
	2956: uint16(anon_sym_PIPE_PIPE),
	2957: uint16(anon_sym_AMP_AMP),
	2958: uint16(anon_sym_CARET),
	2959: uint16(anon_sym_EQ_EQ),
	2960: uint16(anon_sym_BANG_EQ),
	2961: uint16(anon_sym_LT_EQ),
	2962: uint16(anon_sym_GT_EQ),
	2963: uint16(anon_sym_LT_LT),
	2964: uint16(anon_sym_GT_GT),
	2965: uint16(anon_sym_PLUS),
	2966: uint16(anon_sym_DASH),
	2967: uint16(anon_sym_STAR),
	2968: uint16(anon_sym_PERCENT),
	2969: uint16(anon_sym_LBRACK),
	2970: uint16(anon_sym_RBRACK),
	2971: uint16(anon_sym_DOT),
	2972: uint16(4),
	2973: uint16(307),
	2974: uint16(1),
	2975: uint16(anon_sym_LPAREN),
	2976: uint16(3),
	2977: uint16(2),
	2978: uint16(sym_block_comment),
	2979: uint16(sym_line_comment),
	2980: uint16(309),
	2981: uint16(5),
	2982: uint16(anon_sym_LT),
	2983: uint16(anon_sym_GT),
	2984: uint16(anon_sym_PIPE),
	2985: uint16(anon_sym_AMP),
	2986: uint16(anon_sym_SLASH),
	2987: uint16(305),
	2988: uint16(20),
	2989: uint16(anon_sym_SEMI),
	2990: uint16(anon_sym_COMMA),
	2991: uint16(anon_sym_RPAREN),
	2992: uint16(anon_sym_LBRACE),
	2993: uint16(anon_sym_PIPE_PIPE),
	2994: uint16(anon_sym_AMP_AMP),
	2995: uint16(anon_sym_CARET),
	2996: uint16(anon_sym_EQ_EQ),
	2997: uint16(anon_sym_BANG_EQ),
	2998: uint16(anon_sym_LT_EQ),
	2999: uint16(anon_sym_GT_EQ),
	3000: uint16(anon_sym_LT_LT),
	3001: uint16(anon_sym_GT_GT),
	3002: uint16(anon_sym_PLUS),
	3003: uint16(anon_sym_DASH),
	3004: uint16(anon_sym_STAR),
	3005: uint16(anon_sym_PERCENT),
	3006: uint16(anon_sym_LBRACK),
	3007: uint16(anon_sym_RBRACK),
	3008: uint16(anon_sym_DOT),
	3009: uint16(3),
	3010: uint16(3),
	3011: uint16(2),
	3012: uint16(sym_block_comment),
	3013: uint16(sym_line_comment),
	3014: uint16(313),
	3015: uint16(5),
	3016: uint16(anon_sym_LT),
	3017: uint16(anon_sym_GT),
	3018: uint16(anon_sym_PIPE),
	3019: uint16(anon_sym_AMP),
	3020: uint16(anon_sym_SLASH),
	3021: uint16(311),
	3022: uint16(21),
	3023: uint16(anon_sym_SEMI),
	3024: uint16(anon_sym_COMMA),
	3025: uint16(anon_sym_RPAREN),
	3026: uint16(anon_sym_LBRACE),
	3027: uint16(anon_sym_COLON),
	3028: uint16(anon_sym_PIPE_PIPE),
	3029: uint16(anon_sym_AMP_AMP),
	3030: uint16(anon_sym_CARET),
	3031: uint16(anon_sym_EQ_EQ),
	3032: uint16(anon_sym_BANG_EQ),
	3033: uint16(anon_sym_LT_EQ),
	3034: uint16(anon_sym_GT_EQ),
	3035: uint16(anon_sym_LT_LT),
	3036: uint16(anon_sym_GT_GT),
	3037: uint16(anon_sym_PLUS),
	3038: uint16(anon_sym_DASH),
	3039: uint16(anon_sym_STAR),
	3040: uint16(anon_sym_PERCENT),
	3041: uint16(anon_sym_LBRACK),
	3042: uint16(anon_sym_RBRACK),
	3043: uint16(anon_sym_DOT),
	3044: uint16(3),
	3045: uint16(3),
	3046: uint16(2),
	3047: uint16(sym_block_comment),
	3048: uint16(sym_line_comment),
	3049: uint16(317),
	3050: uint16(5),
	3051: uint16(anon_sym_LT),
	3052: uint16(anon_sym_GT),
	3053: uint16(anon_sym_PIPE),
	3054: uint16(anon_sym_AMP),
	3055: uint16(anon_sym_SLASH),
	3056: uint16(315),
	3057: uint16(21),
	3058: uint16(anon_sym_SEMI),
	3059: uint16(anon_sym_COMMA),
	3060: uint16(anon_sym_RPAREN),
	3061: uint16(anon_sym_LBRACE),
	3062: uint16(anon_sym_COLON),
	3063: uint16(anon_sym_PIPE_PIPE),
	3064: uint16(anon_sym_AMP_AMP),
	3065: uint16(anon_sym_CARET),
	3066: uint16(anon_sym_EQ_EQ),
	3067: uint16(anon_sym_BANG_EQ),
	3068: uint16(anon_sym_LT_EQ),
	3069: uint16(anon_sym_GT_EQ),
	3070: uint16(anon_sym_LT_LT),
	3071: uint16(anon_sym_GT_GT),
	3072: uint16(anon_sym_PLUS),
	3073: uint16(anon_sym_DASH),
	3074: uint16(anon_sym_STAR),
	3075: uint16(anon_sym_PERCENT),
	3076: uint16(anon_sym_LBRACK),
	3077: uint16(anon_sym_RBRACK),
	3078: uint16(anon_sym_DOT),
	3079: uint16(3),
	3080: uint16(3),
	3081: uint16(2),
	3082: uint16(sym_block_comment),
	3083: uint16(sym_line_comment),
	3084: uint16(321),
	3085: uint16(5),
	3086: uint16(anon_sym_LT),
	3087: uint16(anon_sym_GT),
	3088: uint16(anon_sym_PIPE),
	3089: uint16(anon_sym_AMP),
	3090: uint16(anon_sym_SLASH),
	3091: uint16(319),
	3092: uint16(20),
	3093: uint16(anon_sym_SEMI),
	3094: uint16(anon_sym_COMMA),
	3095: uint16(anon_sym_RPAREN),
	3096: uint16(anon_sym_LBRACE),
	3097: uint16(anon_sym_PIPE_PIPE),
	3098: uint16(anon_sym_AMP_AMP),
	3099: uint16(anon_sym_CARET),
	3100: uint16(anon_sym_EQ_EQ),
	3101: uint16(anon_sym_BANG_EQ),
	3102: uint16(anon_sym_LT_EQ),
	3103: uint16(anon_sym_GT_EQ),
	3104: uint16(anon_sym_LT_LT),
	3105: uint16(anon_sym_GT_GT),
	3106: uint16(anon_sym_PLUS),
	3107: uint16(anon_sym_DASH),
	3108: uint16(anon_sym_STAR),
	3109: uint16(anon_sym_PERCENT),
	3110: uint16(anon_sym_LBRACK),
	3111: uint16(anon_sym_RBRACK),
	3112: uint16(anon_sym_DOT),
	3113: uint16(13),
	3114: uint16(327),
	3115: uint16(1),
	3116: uint16(anon_sym_PIPE),
	3117: uint16(329),
	3118: uint16(1),
	3119: uint16(anon_sym_AMP),
	3120: uint16(341),
	3121: uint16(1),
	3122: uint16(anon_sym_SLASH),
	3123: uint16(343),
	3124: uint16(1),
	3125: uint16(anon_sym_LBRACK),
	3126: uint16(345),
	3127: uint16(1),
	3128: uint16(anon_sym_DOT),
	3129: uint16(3),
	3130: uint16(2),
	3131: uint16(sym_block_comment),
	3132: uint16(sym_line_comment),
	3133: uint16(325),
	3134: uint16(2),
	3135: uint16(anon_sym_LT),
	3136: uint16(anon_sym_GT),
	3137: uint16(331),
	3138: uint16(2),
	3139: uint16(anon_sym_EQ_EQ),
	3140: uint16(anon_sym_BANG_EQ),
	3141: uint16(333),
	3142: uint16(2),
	3143: uint16(anon_sym_LT_EQ),
	3144: uint16(anon_sym_GT_EQ),
	3145: uint16(335),
	3146: uint16(2),
	3147: uint16(anon_sym_LT_LT),
	3148: uint16(anon_sym_GT_GT),
	3149: uint16(337),
	3150: uint16(2),
	3151: uint16(anon_sym_PLUS),
	3152: uint16(anon_sym_DASH),
	3153: uint16(339),
	3154: uint16(2),
	3155: uint16(anon_sym_STAR),
	3156: uint16(anon_sym_PERCENT),
	3157: uint16(323),
	3158: uint16(8),
	3159: uint16(anon_sym_SEMI),
	3160: uint16(anon_sym_COMMA),
	3161: uint16(anon_sym_RPAREN),
	3162: uint16(anon_sym_LBRACE),
	3163: uint16(anon_sym_PIPE_PIPE),
	3164: uint16(anon_sym_AMP_AMP),
	3165: uint16(anon_sym_CARET),
	3166: uint16(anon_sym_RBRACK),
	3167: uint16(7),
	3168: uint16(341),
	3169: uint16(1),
	3170: uint16(anon_sym_SLASH),
	3171: uint16(343),
	3172: uint16(1),
	3173: uint16(anon_sym_LBRACK),
	3174: uint16(345),
	3175: uint16(1),
	3176: uint16(anon_sym_DOT),
	3177: uint16(3),
	3178: uint16(2),
	3179: uint16(sym_block_comment),
	3180: uint16(sym_line_comment),
	3181: uint16(339),
	3182: uint16(2),
	3183: uint16(anon_sym_STAR),
	3184: uint16(anon_sym_PERCENT),
	3185: uint16(327),
	3186: uint16(4),
	3187: uint16(anon_sym_LT),
	3188: uint16(anon_sym_GT),
	3189: uint16(anon_sym_PIPE),
	3190: uint16(anon_sym_AMP),
	3191: uint16(323),
	3192: uint16(16),
	3193: uint16(anon_sym_SEMI),
	3194: uint16(anon_sym_COMMA),
	3195: uint16(anon_sym_RPAREN),
	3196: uint16(anon_sym_LBRACE),
	3197: uint16(anon_sym_PIPE_PIPE),
	3198: uint16(anon_sym_AMP_AMP),
	3199: uint16(anon_sym_CARET),
	3200: uint16(anon_sym_EQ_EQ),
	3201: uint16(anon_sym_BANG_EQ),
	3202: uint16(anon_sym_LT_EQ),
	3203: uint16(anon_sym_GT_EQ),
	3204: uint16(anon_sym_LT_LT),
	3205: uint16(anon_sym_GT_GT),
	3206: uint16(anon_sym_PLUS),
	3207: uint16(anon_sym_DASH),
	3208: uint16(anon_sym_RBRACK),
	3209: uint16(3),
	3210: uint16(3),
	3211: uint16(2),
	3212: uint16(sym_block_comment),
	3213: uint16(sym_line_comment),
	3214: uint16(349),
	3215: uint16(5),
	3216: uint16(anon_sym_LT),
	3217: uint16(anon_sym_GT),
	3218: uint16(anon_sym_PIPE),
	3219: uint16(anon_sym_AMP),
	3220: uint16(anon_sym_SLASH),
	3221: uint16(347),
	3222: uint16(20),
	3223: uint16(anon_sym_SEMI),
	3224: uint16(anon_sym_COMMA),
	3225: uint16(anon_sym_RPAREN),
	3226: uint16(anon_sym_LBRACE),
	3227: uint16(anon_sym_PIPE_PIPE),
	3228: uint16(anon_sym_AMP_AMP),
	3229: uint16(anon_sym_CARET),
	3230: uint16(anon_sym_EQ_EQ),
	3231: uint16(anon_sym_BANG_EQ),
	3232: uint16(anon_sym_LT_EQ),
	3233: uint16(anon_sym_GT_EQ),
	3234: uint16(anon_sym_LT_LT),
	3235: uint16(anon_sym_GT_GT),
	3236: uint16(anon_sym_PLUS),
	3237: uint16(anon_sym_DASH),
	3238: uint16(anon_sym_STAR),
	3239: uint16(anon_sym_PERCENT),
	3240: uint16(anon_sym_LBRACK),
	3241: uint16(anon_sym_RBRACK),
	3242: uint16(anon_sym_DOT),
	3243: uint16(8),
	3244: uint16(341),
	3245: uint16(1),
	3246: uint16(anon_sym_SLASH),
	3247: uint16(343),
	3248: uint16(1),
	3249: uint16(anon_sym_LBRACK),
	3250: uint16(345),
	3251: uint16(1),
	3252: uint16(anon_sym_DOT),
	3253: uint16(3),
	3254: uint16(2),
	3255: uint16(sym_block_comment),
	3256: uint16(sym_line_comment),
	3257: uint16(337),
	3258: uint16(2),
	3259: uint16(anon_sym_PLUS),
	3260: uint16(anon_sym_DASH),
	3261: uint16(339),
	3262: uint16(2),
	3263: uint16(anon_sym_STAR),
	3264: uint16(anon_sym_PERCENT),
	3265: uint16(327),
	3266: uint16(4),
	3267: uint16(anon_sym_LT),
	3268: uint16(anon_sym_GT),
	3269: uint16(anon_sym_PIPE),
	3270: uint16(anon_sym_AMP),
	3271: uint16(323),
	3272: uint16(14),
	3273: uint16(anon_sym_SEMI),
	3274: uint16(anon_sym_COMMA),
	3275: uint16(anon_sym_RPAREN),
	3276: uint16(anon_sym_LBRACE),
	3277: uint16(anon_sym_PIPE_PIPE),
	3278: uint16(anon_sym_AMP_AMP),
	3279: uint16(anon_sym_CARET),
	3280: uint16(anon_sym_EQ_EQ),
	3281: uint16(anon_sym_BANG_EQ),
	3282: uint16(anon_sym_LT_EQ),
	3283: uint16(anon_sym_GT_EQ),
	3284: uint16(anon_sym_LT_LT),
	3285: uint16(anon_sym_GT_GT),
	3286: uint16(anon_sym_RBRACK),
	3287: uint16(3),
	3288: uint16(3),
	3289: uint16(2),
	3290: uint16(sym_block_comment),
	3291: uint16(sym_line_comment),
	3292: uint16(353),
	3293: uint16(5),
	3294: uint16(anon_sym_LT),
	3295: uint16(anon_sym_GT),
	3296: uint16(anon_sym_PIPE),
	3297: uint16(anon_sym_AMP),
	3298: uint16(anon_sym_SLASH),
	3299: uint16(351),
	3300: uint16(20),
	3301: uint16(anon_sym_SEMI),
	3302: uint16(anon_sym_COMMA),
	3303: uint16(anon_sym_RPAREN),
	3304: uint16(anon_sym_LBRACE),
	3305: uint16(anon_sym_PIPE_PIPE),
	3306: uint16(anon_sym_AMP_AMP),
	3307: uint16(anon_sym_CARET),
	3308: uint16(anon_sym_EQ_EQ),
	3309: uint16(anon_sym_BANG_EQ),
	3310: uint16(anon_sym_LT_EQ),
	3311: uint16(anon_sym_GT_EQ),
	3312: uint16(anon_sym_LT_LT),
	3313: uint16(anon_sym_GT_GT),
	3314: uint16(anon_sym_PLUS),
	3315: uint16(anon_sym_DASH),
	3316: uint16(anon_sym_STAR),
	3317: uint16(anon_sym_PERCENT),
	3318: uint16(anon_sym_LBRACK),
	3319: uint16(anon_sym_RBRACK),
	3320: uint16(anon_sym_DOT),
	3321: uint16(3),
	3322: uint16(3),
	3323: uint16(2),
	3324: uint16(sym_block_comment),
	3325: uint16(sym_line_comment),
	3326: uint16(357),
	3327: uint16(5),
	3328: uint16(anon_sym_LT),
	3329: uint16(anon_sym_GT),
	3330: uint16(anon_sym_PIPE),
	3331: uint16(anon_sym_AMP),
	3332: uint16(anon_sym_SLASH),
	3333: uint16(355),
	3334: uint16(20),
	3335: uint16(anon_sym_SEMI),
	3336: uint16(anon_sym_COMMA),
	3337: uint16(anon_sym_RPAREN),
	3338: uint16(anon_sym_LBRACE),
	3339: uint16(anon_sym_PIPE_PIPE),
	3340: uint16(anon_sym_AMP_AMP),
	3341: uint16(anon_sym_CARET),
	3342: uint16(anon_sym_EQ_EQ),
	3343: uint16(anon_sym_BANG_EQ),
	3344: uint16(anon_sym_LT_EQ),
	3345: uint16(anon_sym_GT_EQ),
	3346: uint16(anon_sym_LT_LT),
	3347: uint16(anon_sym_GT_GT),
	3348: uint16(anon_sym_PLUS),
	3349: uint16(anon_sym_DASH),
	3350: uint16(anon_sym_STAR),
	3351: uint16(anon_sym_PERCENT),
	3352: uint16(anon_sym_LBRACK),
	3353: uint16(anon_sym_RBRACK),
	3354: uint16(anon_sym_DOT),
	3355: uint16(11),
	3356: uint16(341),
	3357: uint16(1),
	3358: uint16(anon_sym_SLASH),
	3359: uint16(343),
	3360: uint16(1),
	3361: uint16(anon_sym_LBRACK),
	3362: uint16(345),
	3363: uint16(1),
	3364: uint16(anon_sym_DOT),
	3365: uint16(3),
	3366: uint16(2),
	3367: uint16(sym_block_comment),
	3368: uint16(sym_line_comment),
	3369: uint16(325),
	3370: uint16(2),
	3371: uint16(anon_sym_LT),
	3372: uint16(anon_sym_GT),
	3373: uint16(327),
	3374: uint16(2),
	3375: uint16(anon_sym_PIPE),
	3376: uint16(anon_sym_AMP),
	3377: uint16(333),
	3378: uint16(2),
	3379: uint16(anon_sym_LT_EQ),
	3380: uint16(anon_sym_GT_EQ),
	3381: uint16(335),
	3382: uint16(2),
	3383: uint16(anon_sym_LT_LT),
	3384: uint16(anon_sym_GT_GT),
	3385: uint16(337),
	3386: uint16(2),
	3387: uint16(anon_sym_PLUS),
	3388: uint16(anon_sym_DASH),
	3389: uint16(339),
	3390: uint16(2),
	3391: uint16(anon_sym_STAR),
	3392: uint16(anon_sym_PERCENT),
	3393: uint16(323),
	3394: uint16(10),
	3395: uint16(anon_sym_SEMI),
	3396: uint16(anon_sym_COMMA),
	3397: uint16(anon_sym_RPAREN),
	3398: uint16(anon_sym_LBRACE),
	3399: uint16(anon_sym_PIPE_PIPE),
	3400: uint16(anon_sym_AMP_AMP),
	3401: uint16(anon_sym_CARET),
	3402: uint16(anon_sym_EQ_EQ),
	3403: uint16(anon_sym_BANG_EQ),
	3404: uint16(anon_sym_RBRACK),
	3405: uint16(3),
	3406: uint16(3),
	3407: uint16(2),
	3408: uint16(sym_block_comment),
	3409: uint16(sym_line_comment),
	3410: uint16(361),
	3411: uint16(5),
	3412: uint16(anon_sym_LT),
	3413: uint16(anon_sym_GT),
	3414: uint16(anon_sym_PIPE),
	3415: uint16(anon_sym_AMP),
	3416: uint16(anon_sym_SLASH),
	3417: uint16(359),
	3418: uint16(20),
	3419: uint16(anon_sym_SEMI),
	3420: uint16(anon_sym_COMMA),
	3421: uint16(anon_sym_RPAREN),
	3422: uint16(anon_sym_LBRACE),
	3423: uint16(anon_sym_PIPE_PIPE),
	3424: uint16(anon_sym_AMP_AMP),
	3425: uint16(anon_sym_CARET),
	3426: uint16(anon_sym_EQ_EQ),
	3427: uint16(anon_sym_BANG_EQ),
	3428: uint16(anon_sym_LT_EQ),
	3429: uint16(anon_sym_GT_EQ),
	3430: uint16(anon_sym_LT_LT),
	3431: uint16(anon_sym_GT_GT),
	3432: uint16(anon_sym_PLUS),
	3433: uint16(anon_sym_DASH),
	3434: uint16(anon_sym_STAR),
	3435: uint16(anon_sym_PERCENT),
	3436: uint16(anon_sym_LBRACK),
	3437: uint16(anon_sym_RBRACK),
	3438: uint16(anon_sym_DOT),
	3439: uint16(3),
	3440: uint16(3),
	3441: uint16(2),
	3442: uint16(sym_block_comment),
	3443: uint16(sym_line_comment),
	3444: uint16(365),
	3445: uint16(5),
	3446: uint16(anon_sym_LT),
	3447: uint16(anon_sym_GT),
	3448: uint16(anon_sym_PIPE),
	3449: uint16(anon_sym_AMP),
	3450: uint16(anon_sym_SLASH),
	3451: uint16(363),
	3452: uint16(20),
	3453: uint16(anon_sym_SEMI),
	3454: uint16(anon_sym_COMMA),
	3455: uint16(anon_sym_RPAREN),
	3456: uint16(anon_sym_LBRACE),
	3457: uint16(anon_sym_PIPE_PIPE),
	3458: uint16(anon_sym_AMP_AMP),
	3459: uint16(anon_sym_CARET),
	3460: uint16(anon_sym_EQ_EQ),
	3461: uint16(anon_sym_BANG_EQ),
	3462: uint16(anon_sym_LT_EQ),
	3463: uint16(anon_sym_GT_EQ),
	3464: uint16(anon_sym_LT_LT),
	3465: uint16(anon_sym_GT_GT),
	3466: uint16(anon_sym_PLUS),
	3467: uint16(anon_sym_DASH),
	3468: uint16(anon_sym_STAR),
	3469: uint16(anon_sym_PERCENT),
	3470: uint16(anon_sym_LBRACK),
	3471: uint16(anon_sym_RBRACK),
	3472: uint16(anon_sym_DOT),
	3473: uint16(9),
	3474: uint16(341),
	3475: uint16(1),
	3476: uint16(anon_sym_SLASH),
	3477: uint16(343),
	3478: uint16(1),
	3479: uint16(anon_sym_LBRACK),
	3480: uint16(345),
	3481: uint16(1),
	3482: uint16(anon_sym_DOT),
	3483: uint16(3),
	3484: uint16(2),
	3485: uint16(sym_block_comment),
	3486: uint16(sym_line_comment),
	3487: uint16(335),
	3488: uint16(2),
	3489: uint16(anon_sym_LT_LT),
	3490: uint16(anon_sym_GT_GT),
	3491: uint16(337),
	3492: uint16(2),
	3493: uint16(anon_sym_PLUS),
	3494: uint16(anon_sym_DASH),
	3495: uint16(339),
	3496: uint16(2),
	3497: uint16(anon_sym_STAR),
	3498: uint16(anon_sym_PERCENT),
	3499: uint16(327),
	3500: uint16(4),
	3501: uint16(anon_sym_LT),
	3502: uint16(anon_sym_GT),
	3503: uint16(anon_sym_PIPE),
	3504: uint16(anon_sym_AMP),
	3505: uint16(323),
	3506: uint16(12),
	3507: uint16(anon_sym_SEMI),
	3508: uint16(anon_sym_COMMA),
	3509: uint16(anon_sym_RPAREN),
	3510: uint16(anon_sym_LBRACE),
	3511: uint16(anon_sym_PIPE_PIPE),
	3512: uint16(anon_sym_AMP_AMP),
	3513: uint16(anon_sym_CARET),
	3514: uint16(anon_sym_EQ_EQ),
	3515: uint16(anon_sym_BANG_EQ),
	3516: uint16(anon_sym_LT_EQ),
	3517: uint16(anon_sym_GT_EQ),
	3518: uint16(anon_sym_RBRACK),
	3519: uint16(5),
	3520: uint16(343),
	3521: uint16(1),
	3522: uint16(anon_sym_LBRACK),
	3523: uint16(345),
	3524: uint16(1),
	3525: uint16(anon_sym_DOT),
	3526: uint16(3),
	3527: uint16(2),
	3528: uint16(sym_block_comment),
	3529: uint16(sym_line_comment),
	3530: uint16(327),
	3531: uint16(5),
	3532: uint16(anon_sym_LT),
	3533: uint16(anon_sym_GT),
	3534: uint16(anon_sym_PIPE),
	3535: uint16(anon_sym_AMP),
	3536: uint16(anon_sym_SLASH),
	3537: uint16(323),
	3538: uint16(18),
	3539: uint16(anon_sym_SEMI),
	3540: uint16(anon_sym_COMMA),
	3541: uint16(anon_sym_RPAREN),
	3542: uint16(anon_sym_LBRACE),
	3543: uint16(anon_sym_PIPE_PIPE),
	3544: uint16(anon_sym_AMP_AMP),
	3545: uint16(anon_sym_CARET),
	3546: uint16(anon_sym_EQ_EQ),
	3547: uint16(anon_sym_BANG_EQ),
	3548: uint16(anon_sym_LT_EQ),
	3549: uint16(anon_sym_GT_EQ),
	3550: uint16(anon_sym_LT_LT),
	3551: uint16(anon_sym_GT_GT),
	3552: uint16(anon_sym_PLUS),
	3553: uint16(anon_sym_DASH),
	3554: uint16(anon_sym_STAR),
	3555: uint16(anon_sym_PERCENT),
	3556: uint16(anon_sym_RBRACK),
	3557: uint16(3),
	3558: uint16(3),
	3559: uint16(2),
	3560: uint16(sym_block_comment),
	3561: uint16(sym_line_comment),
	3562: uint16(369),
	3563: uint16(5),
	3564: uint16(anon_sym_LT),
	3565: uint16(anon_sym_GT),
	3566: uint16(anon_sym_PIPE),
	3567: uint16(anon_sym_AMP),
	3568: uint16(anon_sym_SLASH),
	3569: uint16(367),
	3570: uint16(20),
	3571: uint16(anon_sym_SEMI),
	3572: uint16(anon_sym_COMMA),
	3573: uint16(anon_sym_RPAREN),
	3574: uint16(anon_sym_LBRACE),
	3575: uint16(anon_sym_PIPE_PIPE),
	3576: uint16(anon_sym_AMP_AMP),
	3577: uint16(anon_sym_CARET),
	3578: uint16(anon_sym_EQ_EQ),
	3579: uint16(anon_sym_BANG_EQ),
	3580: uint16(anon_sym_LT_EQ),
	3581: uint16(anon_sym_GT_EQ),
	3582: uint16(anon_sym_LT_LT),
	3583: uint16(anon_sym_GT_GT),
	3584: uint16(anon_sym_PLUS),
	3585: uint16(anon_sym_DASH),
	3586: uint16(anon_sym_STAR),
	3587: uint16(anon_sym_PERCENT),
	3588: uint16(anon_sym_LBRACK),
	3589: uint16(anon_sym_RBRACK),
	3590: uint16(anon_sym_DOT),
	3591: uint16(3),
	3592: uint16(3),
	3593: uint16(2),
	3594: uint16(sym_block_comment),
	3595: uint16(sym_line_comment),
	3596: uint16(373),
	3597: uint16(5),
	3598: uint16(anon_sym_LT),
	3599: uint16(anon_sym_GT),
	3600: uint16(anon_sym_PIPE),
	3601: uint16(anon_sym_AMP),
	3602: uint16(anon_sym_SLASH),
	3603: uint16(371),
	3604: uint16(20),
	3605: uint16(anon_sym_SEMI),
	3606: uint16(anon_sym_COMMA),
	3607: uint16(anon_sym_RPAREN),
	3608: uint16(anon_sym_LBRACE),
	3609: uint16(anon_sym_PIPE_PIPE),
	3610: uint16(anon_sym_AMP_AMP),
	3611: uint16(anon_sym_CARET),
	3612: uint16(anon_sym_EQ_EQ),
	3613: uint16(anon_sym_BANG_EQ),
	3614: uint16(anon_sym_LT_EQ),
	3615: uint16(anon_sym_GT_EQ),
	3616: uint16(anon_sym_LT_LT),
	3617: uint16(anon_sym_GT_GT),
	3618: uint16(anon_sym_PLUS),
	3619: uint16(anon_sym_DASH),
	3620: uint16(anon_sym_STAR),
	3621: uint16(anon_sym_PERCENT),
	3622: uint16(anon_sym_LBRACK),
	3623: uint16(anon_sym_RBRACK),
	3624: uint16(anon_sym_DOT),
	3625: uint16(15),
	3626: uint16(329),
	3627: uint16(1),
	3628: uint16(anon_sym_AMP),
	3629: uint16(341),
	3630: uint16(1),
	3631: uint16(anon_sym_SLASH),
	3632: uint16(343),
	3633: uint16(1),
	3634: uint16(anon_sym_LBRACK),
	3635: uint16(345),
	3636: uint16(1),
	3637: uint16(anon_sym_DOT),
	3638: uint16(375),
	3639: uint16(1),
	3640: uint16(anon_sym_AMP_AMP),
	3641: uint16(377),
	3642: uint16(1),
	3643: uint16(anon_sym_PIPE),
	3644: uint16(379),
	3645: uint16(1),
	3646: uint16(anon_sym_CARET),
	3647: uint16(3),
	3648: uint16(2),
	3649: uint16(sym_block_comment),
	3650: uint16(sym_line_comment),
	3651: uint16(325),
	3652: uint16(2),
	3653: uint16(anon_sym_LT),
	3654: uint16(anon_sym_GT),
	3655: uint16(331),
	3656: uint16(2),
	3657: uint16(anon_sym_EQ_EQ),
	3658: uint16(anon_sym_BANG_EQ),
	3659: uint16(333),
	3660: uint16(2),
	3661: uint16(anon_sym_LT_EQ),
	3662: uint16(anon_sym_GT_EQ),
	3663: uint16(335),
	3664: uint16(2),
	3665: uint16(anon_sym_LT_LT),
	3666: uint16(anon_sym_GT_GT),
	3667: uint16(337),
	3668: uint16(2),
	3669: uint16(anon_sym_PLUS),
	3670: uint16(anon_sym_DASH),
	3671: uint16(339),
	3672: uint16(2),
	3673: uint16(anon_sym_STAR),
	3674: uint16(anon_sym_PERCENT),
	3675: uint16(323),
	3676: uint16(6),
	3677: uint16(anon_sym_SEMI),
	3678: uint16(anon_sym_COMMA),
	3679: uint16(anon_sym_RPAREN),
	3680: uint16(anon_sym_LBRACE),
	3681: uint16(anon_sym_PIPE_PIPE),
	3682: uint16(anon_sym_RBRACK),
	3683: uint16(14),
	3684: uint16(329),
	3685: uint16(1),
	3686: uint16(anon_sym_AMP),
	3687: uint16(341),
	3688: uint16(1),
	3689: uint16(anon_sym_SLASH),
	3690: uint16(343),
	3691: uint16(1),
	3692: uint16(anon_sym_LBRACK),
	3693: uint16(345),
	3694: uint16(1),
	3695: uint16(anon_sym_DOT),
	3696: uint16(377),
	3697: uint16(1),
	3698: uint16(anon_sym_PIPE),
	3699: uint16(379),
	3700: uint16(1),
	3701: uint16(anon_sym_CARET),
	3702: uint16(3),
	3703: uint16(2),
	3704: uint16(sym_block_comment),
	3705: uint16(sym_line_comment),
	3706: uint16(325),
	3707: uint16(2),
	3708: uint16(anon_sym_LT),
	3709: uint16(anon_sym_GT),
	3710: uint16(331),
	3711: uint16(2),
	3712: uint16(anon_sym_EQ_EQ),
	3713: uint16(anon_sym_BANG_EQ),
	3714: uint16(333),
	3715: uint16(2),
	3716: uint16(anon_sym_LT_EQ),
	3717: uint16(anon_sym_GT_EQ),
	3718: uint16(335),
	3719: uint16(2),
	3720: uint16(anon_sym_LT_LT),
	3721: uint16(anon_sym_GT_GT),
	3722: uint16(337),
	3723: uint16(2),
	3724: uint16(anon_sym_PLUS),
	3725: uint16(anon_sym_DASH),
	3726: uint16(339),
	3727: uint16(2),
	3728: uint16(anon_sym_STAR),
	3729: uint16(anon_sym_PERCENT),
	3730: uint16(323),
	3731: uint16(7),
	3732: uint16(anon_sym_SEMI),
	3733: uint16(anon_sym_COMMA),
	3734: uint16(anon_sym_RPAREN),
	3735: uint16(anon_sym_LBRACE),
	3736: uint16(anon_sym_PIPE_PIPE),
	3737: uint16(anon_sym_AMP_AMP),
	3738: uint16(anon_sym_RBRACK),
	3739: uint16(12),
	3740: uint16(341),
	3741: uint16(1),
	3742: uint16(anon_sym_SLASH),
	3743: uint16(343),
	3744: uint16(1),
	3745: uint16(anon_sym_LBRACK),
	3746: uint16(345),
	3747: uint16(1),
	3748: uint16(anon_sym_DOT),
	3749: uint16(3),
	3750: uint16(2),
	3751: uint16(sym_block_comment),
	3752: uint16(sym_line_comment),
	3753: uint16(325),
	3754: uint16(2),
	3755: uint16(anon_sym_LT),
	3756: uint16(anon_sym_GT),
	3757: uint16(327),
	3758: uint16(2),
	3759: uint16(anon_sym_PIPE),
	3760: uint16(anon_sym_AMP),
	3761: uint16(331),
	3762: uint16(2),
	3763: uint16(anon_sym_EQ_EQ),
	3764: uint16(anon_sym_BANG_EQ),
	3765: uint16(333),
	3766: uint16(2),
	3767: uint16(anon_sym_LT_EQ),
	3768: uint16(anon_sym_GT_EQ),
	3769: uint16(335),
	3770: uint16(2),
	3771: uint16(anon_sym_LT_LT),
	3772: uint16(anon_sym_GT_GT),
	3773: uint16(337),
	3774: uint16(2),
	3775: uint16(anon_sym_PLUS),
	3776: uint16(anon_sym_DASH),
	3777: uint16(339),
	3778: uint16(2),
	3779: uint16(anon_sym_STAR),
	3780: uint16(anon_sym_PERCENT),
	3781: uint16(323),
	3782: uint16(8),
	3783: uint16(anon_sym_SEMI),
	3784: uint16(anon_sym_COMMA),
	3785: uint16(anon_sym_RPAREN),
	3786: uint16(anon_sym_LBRACE),
	3787: uint16(anon_sym_PIPE_PIPE),
	3788: uint16(anon_sym_AMP_AMP),
	3789: uint16(anon_sym_CARET),
	3790: uint16(anon_sym_RBRACK),
	3791: uint16(14),
	3792: uint16(327),
	3793: uint16(1),
	3794: uint16(anon_sym_PIPE),
	3795: uint16(329),
	3796: uint16(1),
	3797: uint16(anon_sym_AMP),
	3798: uint16(341),
	3799: uint16(1),
	3800: uint16(anon_sym_SLASH),
	3801: uint16(343),
	3802: uint16(1),
	3803: uint16(anon_sym_LBRACK),
	3804: uint16(345),
	3805: uint16(1),
	3806: uint16(anon_sym_DOT),
	3807: uint16(379),
	3808: uint16(1),
	3809: uint16(anon_sym_CARET),
	3810: uint16(3),
	3811: uint16(2),
	3812: uint16(sym_block_comment),
	3813: uint16(sym_line_comment),
	3814: uint16(325),
	3815: uint16(2),
	3816: uint16(anon_sym_LT),
	3817: uint16(anon_sym_GT),
	3818: uint16(331),
	3819: uint16(2),
	3820: uint16(anon_sym_EQ_EQ),
	3821: uint16(anon_sym_BANG_EQ),
	3822: uint16(333),
	3823: uint16(2),
	3824: uint16(anon_sym_LT_EQ),
	3825: uint16(anon_sym_GT_EQ),
	3826: uint16(335),
	3827: uint16(2),
	3828: uint16(anon_sym_LT_LT),
	3829: uint16(anon_sym_GT_GT),
	3830: uint16(337),
	3831: uint16(2),
	3832: uint16(anon_sym_PLUS),
	3833: uint16(anon_sym_DASH),
	3834: uint16(339),
	3835: uint16(2),
	3836: uint16(anon_sym_STAR),
	3837: uint16(anon_sym_PERCENT),
	3838: uint16(323),
	3839: uint16(7),
	3840: uint16(anon_sym_SEMI),
	3841: uint16(anon_sym_COMMA),
	3842: uint16(anon_sym_RPAREN),
	3843: uint16(anon_sym_LBRACE),
	3844: uint16(anon_sym_PIPE_PIPE),
	3845: uint16(anon_sym_AMP_AMP),
	3846: uint16(anon_sym_RBRACK),
	3847: uint16(3),
	3848: uint16(3),
	3849: uint16(2),
	3850: uint16(sym_block_comment),
	3851: uint16(sym_line_comment),
	3852: uint16(383),
	3853: uint16(5),
	3854: uint16(anon_sym_LT),
	3855: uint16(anon_sym_GT),
	3856: uint16(anon_sym_PIPE),
	3857: uint16(anon_sym_AMP),
	3858: uint16(anon_sym_SLASH),
	3859: uint16(381),
	3860: uint16(20),
	3861: uint16(anon_sym_SEMI),
	3862: uint16(anon_sym_COMMA),
	3863: uint16(anon_sym_RPAREN),
	3864: uint16(anon_sym_LBRACE),
	3865: uint16(anon_sym_PIPE_PIPE),
	3866: uint16(anon_sym_AMP_AMP),
	3867: uint16(anon_sym_CARET),
	3868: uint16(anon_sym_EQ_EQ),
	3869: uint16(anon_sym_BANG_EQ),
	3870: uint16(anon_sym_LT_EQ),
	3871: uint16(anon_sym_GT_EQ),
	3872: uint16(anon_sym_LT_LT),
	3873: uint16(anon_sym_GT_GT),
	3874: uint16(anon_sym_PLUS),
	3875: uint16(anon_sym_DASH),
	3876: uint16(anon_sym_STAR),
	3877: uint16(anon_sym_PERCENT),
	3878: uint16(anon_sym_LBRACK),
	3879: uint16(anon_sym_RBRACK),
	3880: uint16(anon_sym_DOT),
	3881: uint16(3),
	3882: uint16(3),
	3883: uint16(2),
	3884: uint16(sym_block_comment),
	3885: uint16(sym_line_comment),
	3886: uint16(387),
	3887: uint16(5),
	3888: uint16(anon_sym_LT),
	3889: uint16(anon_sym_GT),
	3890: uint16(anon_sym_PIPE),
	3891: uint16(anon_sym_AMP),
	3892: uint16(anon_sym_SLASH),
	3893: uint16(385),
	3894: uint16(20),
	3895: uint16(anon_sym_SEMI),
	3896: uint16(anon_sym_COMMA),
	3897: uint16(anon_sym_RPAREN),
	3898: uint16(anon_sym_LBRACE),
	3899: uint16(anon_sym_PIPE_PIPE),
	3900: uint16(anon_sym_AMP_AMP),
	3901: uint16(anon_sym_CARET),
	3902: uint16(anon_sym_EQ_EQ),
	3903: uint16(anon_sym_BANG_EQ),
	3904: uint16(anon_sym_LT_EQ),
	3905: uint16(anon_sym_GT_EQ),
	3906: uint16(anon_sym_LT_LT),
	3907: uint16(anon_sym_GT_GT),
	3908: uint16(anon_sym_PLUS),
	3909: uint16(anon_sym_DASH),
	3910: uint16(anon_sym_STAR),
	3911: uint16(anon_sym_PERCENT),
	3912: uint16(anon_sym_LBRACK),
	3913: uint16(anon_sym_RBRACK),
	3914: uint16(anon_sym_DOT),
	3915: uint16(16),
	3916: uint16(9),
	3917: uint16(1),
	3918: uint16(anon_sym_let),
	3919: uint16(11),
	3920: uint16(1),
	3921: uint16(anon_sym_override),
	3922: uint16(13),
	3923: uint16(1),
	3924: uint16(anon_sym_type),
	3925: uint16(15),
	3926: uint16(1),
	3927: uint16(anon_sym_fn),
	3928: uint16(17),
	3929: uint16(1),
	3930: uint16(anon_sym_struct),
	3931: uint16(19),
	3932: uint16(1),
	3933: uint16(anon_sym_enable),
	3934: uint16(21),
	3935: uint16(1),
	3936: uint16(anon_sym_AT),
	3937: uint16(23),
	3938: uint16(1),
	3939: uint16(anon_sym_var),
	3940: uint16(389),
	3941: uint16(1),
	3943: uint16(391),
	3944: uint16(1),
	3945: uint16(anon_sym_SEMI),
	3946: uint16(243),
	3947: uint16(1),
	3948: uint16(sym_variable_declaration),
	3949: uint16(3),
	3950: uint16(2),
	3951: uint16(sym_block_comment),
	3952: uint16(sym_line_comment),
	3953: uint16(155),
	3954: uint16(2),
	3955: uint16(sym_enable_directive),
	3956: uint16(aux_sym_source_file_repeat1),
	3957: uint16(182),
	3958: uint16(2),
	3959: uint16(sym_attribute),
	3960: uint16(aux_sym_global_variable_declaration_repeat1),
	3961: uint16(342),
	3962: uint16(3),
	3963: uint16(sym_global_variable_declaration),
	3964: uint16(sym_global_constant_declaration),
	3965: uint16(sym_type_alias_declaration),
	3966: uint16(134),
	3967: uint16(4),
	3968: uint16(sym__declaration),
	3969: uint16(sym_function_declaration),
	3970: uint16(sym_struct_declaration),
	3971: uint16(aux_sym_source_file_repeat2),
	3972: uint16(16),
	3973: uint16(329),
	3974: uint16(1),
	3975: uint16(anon_sym_AMP),
	3976: uint16(341),
	3977: uint16(1),
	3978: uint16(anon_sym_SLASH),
	3979: uint16(343),
	3980: uint16(1),
	3981: uint16(anon_sym_LBRACK),
	3982: uint16(345),
	3983: uint16(1),
	3984: uint16(anon_sym_DOT),
	3985: uint16(375),
	3986: uint16(1),
	3987: uint16(anon_sym_AMP_AMP),
	3988: uint16(377),
	3989: uint16(1),
	3990: uint16(anon_sym_PIPE),
	3991: uint16(379),
	3992: uint16(1),
	3993: uint16(anon_sym_CARET),
	3994: uint16(395),
	3995: uint16(1),
	3996: uint16(anon_sym_PIPE_PIPE),
	3997: uint16(3),
	3998: uint16(2),
	3999: uint16(sym_block_comment),
	4000: uint16(sym_line_comment),
	4001: uint16(325),
	4002: uint16(2),
	4003: uint16(anon_sym_LT),
	4004: uint16(anon_sym_GT),
	4005: uint16(331),
	4006: uint16(2),
	4007: uint16(anon_sym_EQ_EQ),
	4008: uint16(anon_sym_BANG_EQ),
	4009: uint16(333),
	4010: uint16(2),
	4011: uint16(anon_sym_LT_EQ),
	4012: uint16(anon_sym_GT_EQ),
	4013: uint16(335),
	4014: uint16(2),
	4015: uint16(anon_sym_LT_LT),
	4016: uint16(anon_sym_GT_GT),
	4017: uint16(337),
	4018: uint16(2),
	4019: uint16(anon_sym_PLUS),
	4020: uint16(anon_sym_DASH),
	4021: uint16(339),
	4022: uint16(2),
	4023: uint16(anon_sym_STAR),
	4024: uint16(anon_sym_PERCENT),
	4025: uint16(393),
	4026: uint16(2),
	4027: uint16(anon_sym_SEMI),
	4028: uint16(anon_sym_RPAREN),
	4029: uint16(17),
	4030: uint16(160),
	4031: uint16(1),
	4032: uint16(anon_sym_RPAREN),
	4033: uint16(329),
	4034: uint16(1),
	4035: uint16(anon_sym_AMP),
	4036: uint16(341),
	4037: uint16(1),
	4038: uint16(anon_sym_SLASH),
	4039: uint16(343),
	4040: uint16(1),
	4041: uint16(anon_sym_LBRACK),
	4042: uint16(345),
	4043: uint16(1),
	4044: uint16(anon_sym_DOT),
	4045: uint16(375),
	4046: uint16(1),
	4047: uint16(anon_sym_AMP_AMP),
	4048: uint16(377),
	4049: uint16(1),
	4050: uint16(anon_sym_PIPE),
	4051: uint16(379),
	4052: uint16(1),
	4053: uint16(anon_sym_CARET),
	4054: uint16(395),
	4055: uint16(1),
	4056: uint16(anon_sym_PIPE_PIPE),
	4057: uint16(397),
	4058: uint16(1),
	4059: uint16(anon_sym_COMMA),
	4060: uint16(3),
	4061: uint16(2),
	4062: uint16(sym_block_comment),
	4063: uint16(sym_line_comment),
	4064: uint16(325),
	4065: uint16(2),
	4066: uint16(anon_sym_LT),
	4067: uint16(anon_sym_GT),
	4068: uint16(331),
	4069: uint16(2),
	4070: uint16(anon_sym_EQ_EQ),
	4071: uint16(anon_sym_BANG_EQ),
	4072: uint16(333),
	4073: uint16(2),
	4074: uint16(anon_sym_LT_EQ),
	4075: uint16(anon_sym_GT_EQ),
	4076: uint16(335),
	4077: uint16(2),
	4078: uint16(anon_sym_LT_LT),
	4079: uint16(anon_sym_GT_GT),
	4080: uint16(337),
	4081: uint16(2),
	4082: uint16(anon_sym_PLUS),
	4083: uint16(anon_sym_DASH),
	4084: uint16(339),
	4085: uint16(2),
	4086: uint16(anon_sym_STAR),
	4087: uint16(anon_sym_PERCENT),
	4088: uint16(17),
	4089: uint16(199),
	4090: uint16(1),
	4091: uint16(anon_sym_LBRACE),
	4092: uint16(329),
	4093: uint16(1),
	4094: uint16(anon_sym_AMP),
	4095: uint16(341),
	4096: uint16(1),
	4097: uint16(anon_sym_SLASH),
	4098: uint16(343),
	4099: uint16(1),
	4100: uint16(anon_sym_LBRACK),
	4101: uint16(345),
	4102: uint16(1),
	4103: uint16(anon_sym_DOT),
	4104: uint16(375),
	4105: uint16(1),
	4106: uint16(anon_sym_AMP_AMP),
	4107: uint16(377),
	4108: uint16(1),
	4109: uint16(anon_sym_PIPE),
	4110: uint16(379),
	4111: uint16(1),
	4112: uint16(anon_sym_CARET),
	4113: uint16(395),
	4114: uint16(1),
	4115: uint16(anon_sym_PIPE_PIPE),
	4116: uint16(123),
	4117: uint16(1),
	4118: uint16(sym_compound_statement),
	4119: uint16(3),
	4120: uint16(2),
	4121: uint16(sym_block_comment),
	4122: uint16(sym_line_comment),
	4123: uint16(325),
	4124: uint16(2),
	4125: uint16(anon_sym_LT),
	4126: uint16(anon_sym_GT),
	4127: uint16(331),
	4128: uint16(2),
	4129: uint16(anon_sym_EQ_EQ),
	4130: uint16(anon_sym_BANG_EQ),
	4131: uint16(333),
	4132: uint16(2),
	4133: uint16(anon_sym_LT_EQ),
	4134: uint16(anon_sym_GT_EQ),
	4135: uint16(335),
	4136: uint16(2),
	4137: uint16(anon_sym_LT_LT),
	4138: uint16(anon_sym_GT_GT),
	4139: uint16(337),
	4140: uint16(2),
	4141: uint16(anon_sym_PLUS),
	4142: uint16(anon_sym_DASH),
	4143: uint16(339),
	4144: uint16(2),
	4145: uint16(anon_sym_STAR),
	4146: uint16(anon_sym_PERCENT),
	4147: uint16(3),
	4148: uint16(3),
	4149: uint16(2),
	4150: uint16(sym_block_comment),
	4151: uint16(sym_line_comment),
	4152: uint16(401),
	4153: uint16(7),
	4154: uint16(anon_sym_SEMI),
	4155: uint16(anon_sym_LPAREN),
	4156: uint16(anon_sym_RPAREN),
	4157: uint16(anon_sym_LBRACE),
	4158: uint16(anon_sym_RBRACE),
	4159: uint16(anon_sym_AMP),
	4160: uint16(anon_sym_STAR),
	4161: uint16(399),
	4162: uint16(15),
	4163: uint16(anon_sym_let),
	4164: uint16(sym_identifier),
	4165: uint16(anon_sym__),
	4166: uint16(anon_sym_if),
	4167: uint16(anon_sym_switch),
	4168: uint16(anon_sym_fallthrough),
	4169: uint16(anon_sym_loop),
	4170: uint16(anon_sym_for),
	4171: uint16(anon_sym_while),
	4172: uint16(anon_sym_break),
	4173: uint16(anon_sym_continue),
	4174: uint16(anon_sym_continuing),
	4175: uint16(anon_sym_return),
	4176: uint16(anon_sym_discard),
	4177: uint16(anon_sym_var),
	4178: uint16(17),
	4179: uint16(199),
	4180: uint16(1),
	4181: uint16(anon_sym_LBRACE),
	4182: uint16(329),
	4183: uint16(1),
	4184: uint16(anon_sym_AMP),
	4185: uint16(341),
	4186: uint16(1),
	4187: uint16(anon_sym_SLASH),
	4188: uint16(343),
	4189: uint16(1),
	4190: uint16(anon_sym_LBRACK),
	4191: uint16(345),
	4192: uint16(1),
	4193: uint16(anon_sym_DOT),
	4194: uint16(375),
	4195: uint16(1),
	4196: uint16(anon_sym_AMP_AMP),
	4197: uint16(377),
	4198: uint16(1),
	4199: uint16(anon_sym_PIPE),
	4200: uint16(379),
	4201: uint16(1),
	4202: uint16(anon_sym_CARET),
	4203: uint16(395),
	4204: uint16(1),
	4205: uint16(anon_sym_PIPE_PIPE),
	4206: uint16(111),
	4207: uint16(1),
	4208: uint16(sym_compound_statement),
	4209: uint16(3),
	4210: uint16(2),
	4211: uint16(sym_block_comment),
	4212: uint16(sym_line_comment),
	4213: uint16(325),
	4214: uint16(2),
	4215: uint16(anon_sym_LT),
	4216: uint16(anon_sym_GT),
	4217: uint16(331),
	4218: uint16(2),
	4219: uint16(anon_sym_EQ_EQ),
	4220: uint16(anon_sym_BANG_EQ),
	4221: uint16(333),
	4222: uint16(2),
	4223: uint16(anon_sym_LT_EQ),
	4224: uint16(anon_sym_GT_EQ),
	4225: uint16(335),
	4226: uint16(2),
	4227: uint16(anon_sym_LT_LT),
	4228: uint16(anon_sym_GT_GT),
	4229: uint16(337),
	4230: uint16(2),
	4231: uint16(anon_sym_PLUS),
	4232: uint16(anon_sym_DASH),
	4233: uint16(339),
	4234: uint16(2),
	4235: uint16(anon_sym_STAR),
	4236: uint16(anon_sym_PERCENT),
	4237: uint16(3),
	4238: uint16(3),
	4239: uint16(2),
	4240: uint16(sym_block_comment),
	4241: uint16(sym_line_comment),
	4242: uint16(405),
	4243: uint16(7),
	4244: uint16(anon_sym_SEMI),
	4245: uint16(anon_sym_LPAREN),
	4246: uint16(anon_sym_RPAREN),
	4247: uint16(anon_sym_LBRACE),
	4248: uint16(anon_sym_RBRACE),
	4249: uint16(anon_sym_AMP),
	4250: uint16(anon_sym_STAR),
	4251: uint16(403),
	4252: uint16(15),
	4253: uint16(anon_sym_let),
	4254: uint16(sym_identifier),
	4255: uint16(anon_sym__),
	4256: uint16(anon_sym_if),
	4257: uint16(anon_sym_switch),
	4258: uint16(anon_sym_fallthrough),
	4259: uint16(anon_sym_loop),
	4260: uint16(anon_sym_for),
	4261: uint16(anon_sym_while),
	4262: uint16(anon_sym_break),
	4263: uint16(anon_sym_continue),
	4264: uint16(anon_sym_continuing),
	4265: uint16(anon_sym_return),
	4266: uint16(anon_sym_discard),
	4267: uint16(anon_sym_var),
	4268: uint16(17),
	4269: uint16(329),
	4270: uint16(1),
	4271: uint16(anon_sym_AMP),
	4272: uint16(341),
	4273: uint16(1),
	4274: uint16(anon_sym_SLASH),
	4275: uint16(343),
	4276: uint16(1),
	4277: uint16(anon_sym_LBRACK),
	4278: uint16(345),
	4279: uint16(1),
	4280: uint16(anon_sym_DOT),
	4281: uint16(375),
	4282: uint16(1),
	4283: uint16(anon_sym_AMP_AMP),
	4284: uint16(377),
	4285: uint16(1),
	4286: uint16(anon_sym_PIPE),
	4287: uint16(379),
	4288: uint16(1),
	4289: uint16(anon_sym_CARET),
	4290: uint16(395),
	4291: uint16(1),
	4292: uint16(anon_sym_PIPE_PIPE),
	4293: uint16(407),
	4294: uint16(1),
	4295: uint16(anon_sym_COMMA),
	4296: uint16(409),
	4297: uint16(1),
	4298: uint16(anon_sym_RPAREN),
	4299: uint16(3),
	4300: uint16(2),
	4301: uint16(sym_block_comment),
	4302: uint16(sym_line_comment),
	4303: uint16(325),
	4304: uint16(2),
	4305: uint16(anon_sym_LT),
	4306: uint16(anon_sym_GT),
	4307: uint16(331),
	4308: uint16(2),
	4309: uint16(anon_sym_EQ_EQ),
	4310: uint16(anon_sym_BANG_EQ),
	4311: uint16(333),
	4312: uint16(2),
	4313: uint16(anon_sym_LT_EQ),
	4314: uint16(anon_sym_GT_EQ),
	4315: uint16(335),
	4316: uint16(2),
	4317: uint16(anon_sym_LT_LT),
	4318: uint16(anon_sym_GT_GT),
	4319: uint16(337),
	4320: uint16(2),
	4321: uint16(anon_sym_PLUS),
	4322: uint16(anon_sym_DASH),
	4323: uint16(339),
	4324: uint16(2),
	4325: uint16(anon_sym_STAR),
	4326: uint16(anon_sym_PERCENT),
	4327: uint16(16),
	4328: uint16(329),
	4329: uint16(1),
	4330: uint16(anon_sym_AMP),
	4331: uint16(341),
	4332: uint16(1),
	4333: uint16(anon_sym_SLASH),
	4334: uint16(343),
	4335: uint16(1),
	4336: uint16(anon_sym_LBRACK),
	4337: uint16(345),
	4338: uint16(1),
	4339: uint16(anon_sym_DOT),
	4340: uint16(375),
	4341: uint16(1),
	4342: uint16(anon_sym_AMP_AMP),
	4343: uint16(377),
	4344: uint16(1),
	4345: uint16(anon_sym_PIPE),
	4346: uint16(379),
	4347: uint16(1),
	4348: uint16(anon_sym_CARET),
	4349: uint16(395),
	4350: uint16(1),
	4351: uint16(anon_sym_PIPE_PIPE),
	4352: uint16(411),
	4353: uint16(1),
	4354: uint16(anon_sym_RPAREN),
	4355: uint16(3),
	4356: uint16(2),
	4357: uint16(sym_block_comment),
	4358: uint16(sym_line_comment),
	4359: uint16(325),
	4360: uint16(2),
	4361: uint16(anon_sym_LT),
	4362: uint16(anon_sym_GT),
	4363: uint16(331),
	4364: uint16(2),
	4365: uint16(anon_sym_EQ_EQ),
	4366: uint16(anon_sym_BANG_EQ),
	4367: uint16(333),
	4368: uint16(2),
	4369: uint16(anon_sym_LT_EQ),
	4370: uint16(anon_sym_GT_EQ),
	4371: uint16(335),
	4372: uint16(2),
	4373: uint16(anon_sym_LT_LT),
	4374: uint16(anon_sym_GT_GT),
	4375: uint16(337),
	4376: uint16(2),
	4377: uint16(anon_sym_PLUS),
	4378: uint16(anon_sym_DASH),
	4379: uint16(339),
	4380: uint16(2),
	4381: uint16(anon_sym_STAR),
	4382: uint16(anon_sym_PERCENT),
	4383: uint16(16),
	4384: uint16(329),
	4385: uint16(1),
	4386: uint16(anon_sym_AMP),
	4387: uint16(341),
	4388: uint16(1),
	4389: uint16(anon_sym_SLASH),
	4390: uint16(343),
	4391: uint16(1),
	4392: uint16(anon_sym_LBRACK),
	4393: uint16(345),
	4394: uint16(1),
	4395: uint16(anon_sym_DOT),
	4396: uint16(375),
	4397: uint16(1),
	4398: uint16(anon_sym_AMP_AMP),
	4399: uint16(377),
	4400: uint16(1),
	4401: uint16(anon_sym_PIPE),
	4402: uint16(379),
	4403: uint16(1),
	4404: uint16(anon_sym_CARET),
	4405: uint16(395),
	4406: uint16(1),
	4407: uint16(anon_sym_PIPE_PIPE),
	4408: uint16(413),
	4409: uint16(1),
	4410: uint16(anon_sym_COMMA),
	4411: uint16(3),
	4412: uint16(2),
	4413: uint16(sym_block_comment),
	4414: uint16(sym_line_comment),
	4415: uint16(325),
	4416: uint16(2),
	4417: uint16(anon_sym_LT),
	4418: uint16(anon_sym_GT),
	4419: uint16(331),
	4420: uint16(2),
	4421: uint16(anon_sym_EQ_EQ),
	4422: uint16(anon_sym_BANG_EQ),
	4423: uint16(333),
	4424: uint16(2),
	4425: uint16(anon_sym_LT_EQ),
	4426: uint16(anon_sym_GT_EQ),
	4427: uint16(335),
	4428: uint16(2),
	4429: uint16(anon_sym_LT_LT),
	4430: uint16(anon_sym_GT_GT),
	4431: uint16(337),
	4432: uint16(2),
	4433: uint16(anon_sym_PLUS),
	4434: uint16(anon_sym_DASH),
	4435: uint16(339),
	4436: uint16(2),
	4437: uint16(anon_sym_STAR),
	4438: uint16(anon_sym_PERCENT),
	4439: uint16(16),
	4440: uint16(329),
	4441: uint16(1),
	4442: uint16(anon_sym_AMP),
	4443: uint16(341),
	4444: uint16(1),
	4445: uint16(anon_sym_SLASH),
	4446: uint16(343),
	4447: uint16(1),
	4448: uint16(anon_sym_LBRACK),
	4449: uint16(345),
	4450: uint16(1),
	4451: uint16(anon_sym_DOT),
	4452: uint16(375),
	4453: uint16(1),
	4454: uint16(anon_sym_AMP_AMP),
	4455: uint16(377),
	4456: uint16(1),
	4457: uint16(anon_sym_PIPE),
	4458: uint16(379),
	4459: uint16(1),
	4460: uint16(anon_sym_CARET),
	4461: uint16(395),
	4462: uint16(1),
	4463: uint16(anon_sym_PIPE_PIPE),
	4464: uint16(415),
	4465: uint16(1),
	4466: uint16(anon_sym_SEMI),
	4467: uint16(3),
	4468: uint16(2),
	4469: uint16(sym_block_comment),
	4470: uint16(sym_line_comment),
	4471: uint16(325),
	4472: uint16(2),
	4473: uint16(anon_sym_LT),
	4474: uint16(anon_sym_GT),
	4475: uint16(331),
	4476: uint16(2),
	4477: uint16(anon_sym_EQ_EQ),
	4478: uint16(anon_sym_BANG_EQ),
	4479: uint16(333),
	4480: uint16(2),
	4481: uint16(anon_sym_LT_EQ),
	4482: uint16(anon_sym_GT_EQ),
	4483: uint16(335),
	4484: uint16(2),
	4485: uint16(anon_sym_LT_LT),
	4486: uint16(anon_sym_GT_GT),
	4487: uint16(337),
	4488: uint16(2),
	4489: uint16(anon_sym_PLUS),
	4490: uint16(anon_sym_DASH),
	4491: uint16(339),
	4492: uint16(2),
	4493: uint16(anon_sym_STAR),
	4494: uint16(anon_sym_PERCENT),
	4495: uint16(16),
	4496: uint16(329),
	4497: uint16(1),
	4498: uint16(anon_sym_AMP),
	4499: uint16(341),
	4500: uint16(1),
	4501: uint16(anon_sym_SLASH),
	4502: uint16(343),
	4503: uint16(1),
	4504: uint16(anon_sym_LBRACK),
	4505: uint16(345),
	4506: uint16(1),
	4507: uint16(anon_sym_DOT),
	4508: uint16(375),
	4509: uint16(1),
	4510: uint16(anon_sym_AMP_AMP),
	4511: uint16(377),
	4512: uint16(1),
	4513: uint16(anon_sym_PIPE),
	4514: uint16(379),
	4515: uint16(1),
	4516: uint16(anon_sym_CARET),
	4517: uint16(395),
	4518: uint16(1),
	4519: uint16(anon_sym_PIPE_PIPE),
	4520: uint16(417),
	4521: uint16(1),
	4522: uint16(anon_sym_SEMI),
	4523: uint16(3),
	4524: uint16(2),
	4525: uint16(sym_block_comment),
	4526: uint16(sym_line_comment),
	4527: uint16(325),
	4528: uint16(2),
	4529: uint16(anon_sym_LT),
	4530: uint16(anon_sym_GT),
	4531: uint16(331),
	4532: uint16(2),
	4533: uint16(anon_sym_EQ_EQ),
	4534: uint16(anon_sym_BANG_EQ),
	4535: uint16(333),
	4536: uint16(2),
	4537: uint16(anon_sym_LT_EQ),
	4538: uint16(anon_sym_GT_EQ),
	4539: uint16(335),
	4540: uint16(2),
	4541: uint16(anon_sym_LT_LT),
	4542: uint16(anon_sym_GT_GT),
	4543: uint16(337),
	4544: uint16(2),
	4545: uint16(anon_sym_PLUS),
	4546: uint16(anon_sym_DASH),
	4547: uint16(339),
	4548: uint16(2),
	4549: uint16(anon_sym_STAR),
	4550: uint16(anon_sym_PERCENT),
	4551: uint16(16),
	4552: uint16(329),
	4553: uint16(1),
	4554: uint16(anon_sym_AMP),
	4555: uint16(341),
	4556: uint16(1),
	4557: uint16(anon_sym_SLASH),
	4558: uint16(343),
	4559: uint16(1),
	4560: uint16(anon_sym_LBRACK),
	4561: uint16(345),
	4562: uint16(1),
	4563: uint16(anon_sym_DOT),
	4564: uint16(375),
	4565: uint16(1),
	4566: uint16(anon_sym_AMP_AMP),
	4567: uint16(377),
	4568: uint16(1),
	4569: uint16(anon_sym_PIPE),
	4570: uint16(379),
	4571: uint16(1),
	4572: uint16(anon_sym_CARET),
	4573: uint16(395),
	4574: uint16(1),
	4575: uint16(anon_sym_PIPE_PIPE),
	4576: uint16(419),
	4577: uint16(1),
	4578: uint16(anon_sym_LBRACE),
	4579: uint16(3),
	4580: uint16(2),
	4581: uint16(sym_block_comment),
	4582: uint16(sym_line_comment),
	4583: uint16(325),
	4584: uint16(2),
	4585: uint16(anon_sym_LT),
	4586: uint16(anon_sym_GT),
	4587: uint16(331),
	4588: uint16(2),
	4589: uint16(anon_sym_EQ_EQ),
	4590: uint16(anon_sym_BANG_EQ),
	4591: uint16(333),
	4592: uint16(2),
	4593: uint16(anon_sym_LT_EQ),
	4594: uint16(anon_sym_GT_EQ),
	4595: uint16(335),
	4596: uint16(2),
	4597: uint16(anon_sym_LT_LT),
	4598: uint16(anon_sym_GT_GT),
	4599: uint16(337),
	4600: uint16(2),
	4601: uint16(anon_sym_PLUS),
	4602: uint16(anon_sym_DASH),
	4603: uint16(339),
	4604: uint16(2),
	4605: uint16(anon_sym_STAR),
	4606: uint16(anon_sym_PERCENT),
	4607: uint16(16),
	4608: uint16(329),
	4609: uint16(1),
	4610: uint16(anon_sym_AMP),
	4611: uint16(341),
	4612: uint16(1),
	4613: uint16(anon_sym_SLASH),
	4614: uint16(343),
	4615: uint16(1),
	4616: uint16(anon_sym_LBRACK),
	4617: uint16(345),
	4618: uint16(1),
	4619: uint16(anon_sym_DOT),
	4620: uint16(375),
	4621: uint16(1),
	4622: uint16(anon_sym_AMP_AMP),
	4623: uint16(377),
	4624: uint16(1),
	4625: uint16(anon_sym_PIPE),
	4626: uint16(379),
	4627: uint16(1),
	4628: uint16(anon_sym_CARET),
	4629: uint16(395),
	4630: uint16(1),
	4631: uint16(anon_sym_PIPE_PIPE),
	4632: uint16(421),
	4633: uint16(1),
	4634: uint16(anon_sym_SEMI),
	4635: uint16(3),
	4636: uint16(2),
	4637: uint16(sym_block_comment),
	4638: uint16(sym_line_comment),
	4639: uint16(325),
	4640: uint16(2),
	4641: uint16(anon_sym_LT),
	4642: uint16(anon_sym_GT),
	4643: uint16(331),
	4644: uint16(2),
	4645: uint16(anon_sym_EQ_EQ),
	4646: uint16(anon_sym_BANG_EQ),
	4647: uint16(333),
	4648: uint16(2),
	4649: uint16(anon_sym_LT_EQ),
	4650: uint16(anon_sym_GT_EQ),
	4651: uint16(335),
	4652: uint16(2),
	4653: uint16(anon_sym_LT_LT),
	4654: uint16(anon_sym_GT_GT),
	4655: uint16(337),
	4656: uint16(2),
	4657: uint16(anon_sym_PLUS),
	4658: uint16(anon_sym_DASH),
	4659: uint16(339),
	4660: uint16(2),
	4661: uint16(anon_sym_STAR),
	4662: uint16(anon_sym_PERCENT),
	4663: uint16(4),
	4664: uint16(427),
	4665: uint16(1),
	4666: uint16(anon_sym_else),
	4667: uint16(3),
	4668: uint16(2),
	4669: uint16(sym_block_comment),
	4670: uint16(sym_line_comment),
	4671: uint16(425),
	4672: uint16(5),
	4673: uint16(anon_sym_LPAREN),
	4674: uint16(anon_sym_LBRACE),
	4675: uint16(anon_sym_RBRACE),
	4676: uint16(anon_sym_AMP),
	4677: uint16(anon_sym_STAR),
	4678: uint16(423),
	4679: uint16(15),
	4680: uint16(anon_sym_let),
	4681: uint16(sym_identifier),
	4682: uint16(anon_sym__),
	4683: uint16(anon_sym_if),
	4684: uint16(anon_sym_switch),
	4685: uint16(anon_sym_fallthrough),
	4686: uint16(anon_sym_loop),
	4687: uint16(anon_sym_for),
	4688: uint16(anon_sym_while),
	4689: uint16(anon_sym_break),
	4690: uint16(anon_sym_continue),
	4691: uint16(anon_sym_continuing),
	4692: uint16(anon_sym_return),
	4693: uint16(anon_sym_discard),
	4694: uint16(anon_sym_var),
	4695: uint16(16),
	4696: uint16(329),
	4697: uint16(1),
	4698: uint16(anon_sym_AMP),
	4699: uint16(341),
	4700: uint16(1),
	4701: uint16(anon_sym_SLASH),
	4702: uint16(343),
	4703: uint16(1),
	4704: uint16(anon_sym_LBRACK),
	4705: uint16(345),
	4706: uint16(1),
	4707: uint16(anon_sym_DOT),
	4708: uint16(375),
	4709: uint16(1),
	4710: uint16(anon_sym_AMP_AMP),
	4711: uint16(377),
	4712: uint16(1),
	4713: uint16(anon_sym_PIPE),
	4714: uint16(379),
	4715: uint16(1),
	4716: uint16(anon_sym_CARET),
	4717: uint16(395),
	4718: uint16(1),
	4719: uint16(anon_sym_PIPE_PIPE),
	4720: uint16(429),
	4721: uint16(1),
	4722: uint16(anon_sym_RBRACK),
	4723: uint16(3),
	4724: uint16(2),
	4725: uint16(sym_block_comment),
	4726: uint16(sym_line_comment),
	4727: uint16(325),
	4728: uint16(2),
	4729: uint16(anon_sym_LT),
	4730: uint16(anon_sym_GT),
	4731: uint16(331),
	4732: uint16(2),
	4733: uint16(anon_sym_EQ_EQ),
	4734: uint16(anon_sym_BANG_EQ),
	4735: uint16(333),
	4736: uint16(2),
	4737: uint16(anon_sym_LT_EQ),
	4738: uint16(anon_sym_GT_EQ),
	4739: uint16(335),
	4740: uint16(2),
	4741: uint16(anon_sym_LT_LT),
	4742: uint16(anon_sym_GT_GT),
	4743: uint16(337),
	4744: uint16(2),
	4745: uint16(anon_sym_PLUS),
	4746: uint16(anon_sym_DASH),
	4747: uint16(339),
	4748: uint16(2),
	4749: uint16(anon_sym_STAR),
	4750: uint16(anon_sym_PERCENT),
	4751: uint16(16),
	4752: uint16(329),
	4753: uint16(1),
	4754: uint16(anon_sym_AMP),
	4755: uint16(341),
	4756: uint16(1),
	4757: uint16(anon_sym_SLASH),
	4758: uint16(343),
	4759: uint16(1),
	4760: uint16(anon_sym_LBRACK),
	4761: uint16(345),
	4762: uint16(1),
	4763: uint16(anon_sym_DOT),
	4764: uint16(375),
	4765: uint16(1),
	4766: uint16(anon_sym_AMP_AMP),
	4767: uint16(377),
	4768: uint16(1),
	4769: uint16(anon_sym_PIPE),
	4770: uint16(379),
	4771: uint16(1),
	4772: uint16(anon_sym_CARET),
	4773: uint16(395),
	4774: uint16(1),
	4775: uint16(anon_sym_PIPE_PIPE),
	4776: uint16(431),
	4777: uint16(1),
	4778: uint16(anon_sym_SEMI),
	4779: uint16(3),
	4780: uint16(2),
	4781: uint16(sym_block_comment),
	4782: uint16(sym_line_comment),
	4783: uint16(325),
	4784: uint16(2),
	4785: uint16(anon_sym_LT),
	4786: uint16(anon_sym_GT),
	4787: uint16(331),
	4788: uint16(2),
	4789: uint16(anon_sym_EQ_EQ),
	4790: uint16(anon_sym_BANG_EQ),
	4791: uint16(333),
	4792: uint16(2),
	4793: uint16(anon_sym_LT_EQ),
	4794: uint16(anon_sym_GT_EQ),
	4795: uint16(335),
	4796: uint16(2),
	4797: uint16(anon_sym_LT_LT),
	4798: uint16(anon_sym_GT_GT),
	4799: uint16(337),
	4800: uint16(2),
	4801: uint16(anon_sym_PLUS),
	4802: uint16(anon_sym_DASH),
	4803: uint16(339),
	4804: uint16(2),
	4805: uint16(anon_sym_STAR),
	4806: uint16(anon_sym_PERCENT),
	4807: uint16(16),
	4808: uint16(329),
	4809: uint16(1),
	4810: uint16(anon_sym_AMP),
	4811: uint16(341),
	4812: uint16(1),
	4813: uint16(anon_sym_SLASH),
	4814: uint16(343),
	4815: uint16(1),
	4816: uint16(anon_sym_LBRACK),
	4817: uint16(345),
	4818: uint16(1),
	4819: uint16(anon_sym_DOT),
	4820: uint16(375),
	4821: uint16(1),
	4822: uint16(anon_sym_AMP_AMP),
	4823: uint16(377),
	4824: uint16(1),
	4825: uint16(anon_sym_PIPE),
	4826: uint16(379),
	4827: uint16(1),
	4828: uint16(anon_sym_CARET),
	4829: uint16(395),
	4830: uint16(1),
	4831: uint16(anon_sym_PIPE_PIPE),
	4832: uint16(433),
	4833: uint16(1),
	4834: uint16(anon_sym_RBRACK),
	4835: uint16(3),
	4836: uint16(2),
	4837: uint16(sym_block_comment),
	4838: uint16(sym_line_comment),
	4839: uint16(325),
	4840: uint16(2),
	4841: uint16(anon_sym_LT),
	4842: uint16(anon_sym_GT),
	4843: uint16(331),
	4844: uint16(2),
	4845: uint16(anon_sym_EQ_EQ),
	4846: uint16(anon_sym_BANG_EQ),
	4847: uint16(333),
	4848: uint16(2),
	4849: uint16(anon_sym_LT_EQ),
	4850: uint16(anon_sym_GT_EQ),
	4851: uint16(335),
	4852: uint16(2),
	4853: uint16(anon_sym_LT_LT),
	4854: uint16(anon_sym_GT_GT),
	4855: uint16(337),
	4856: uint16(2),
	4857: uint16(anon_sym_PLUS),
	4858: uint16(anon_sym_DASH),
	4859: uint16(339),
	4860: uint16(2),
	4861: uint16(anon_sym_STAR),
	4862: uint16(anon_sym_PERCENT),
	4863: uint16(16),
	4864: uint16(329),
	4865: uint16(1),
	4866: uint16(anon_sym_AMP),
	4867: uint16(341),
	4868: uint16(1),
	4869: uint16(anon_sym_SLASH),
	4870: uint16(343),
	4871: uint16(1),
	4872: uint16(anon_sym_LBRACK),
	4873: uint16(345),
	4874: uint16(1),
	4875: uint16(anon_sym_DOT),
	4876: uint16(375),
	4877: uint16(1),
	4878: uint16(anon_sym_AMP_AMP),
	4879: uint16(377),
	4880: uint16(1),
	4881: uint16(anon_sym_PIPE),
	4882: uint16(379),
	4883: uint16(1),
	4884: uint16(anon_sym_CARET),
	4885: uint16(395),
	4886: uint16(1),
	4887: uint16(anon_sym_PIPE_PIPE),
	4888: uint16(435),
	4889: uint16(1),
	4890: uint16(anon_sym_SEMI),
	4891: uint16(3),
	4892: uint16(2),
	4893: uint16(sym_block_comment),
	4894: uint16(sym_line_comment),
	4895: uint16(325),
	4896: uint16(2),
	4897: uint16(anon_sym_LT),
	4898: uint16(anon_sym_GT),
	4899: uint16(331),
	4900: uint16(2),
	4901: uint16(anon_sym_EQ_EQ),
	4902: uint16(anon_sym_BANG_EQ),
	4903: uint16(333),
	4904: uint16(2),
	4905: uint16(anon_sym_LT_EQ),
	4906: uint16(anon_sym_GT_EQ),
	4907: uint16(335),
	4908: uint16(2),
	4909: uint16(anon_sym_LT_LT),
	4910: uint16(anon_sym_GT_GT),
	4911: uint16(337),
	4912: uint16(2),
	4913: uint16(anon_sym_PLUS),
	4914: uint16(anon_sym_DASH),
	4915: uint16(339),
	4916: uint16(2),
	4917: uint16(anon_sym_STAR),
	4918: uint16(anon_sym_PERCENT),
	4919: uint16(16),
	4920: uint16(329),
	4921: uint16(1),
	4922: uint16(anon_sym_AMP),
	4923: uint16(341),
	4924: uint16(1),
	4925: uint16(anon_sym_SLASH),
	4926: uint16(343),
	4927: uint16(1),
	4928: uint16(anon_sym_LBRACK),
	4929: uint16(345),
	4930: uint16(1),
	4931: uint16(anon_sym_DOT),
	4932: uint16(375),
	4933: uint16(1),
	4934: uint16(anon_sym_AMP_AMP),
	4935: uint16(377),
	4936: uint16(1),
	4937: uint16(anon_sym_PIPE),
	4938: uint16(379),
	4939: uint16(1),
	4940: uint16(anon_sym_CARET),
	4941: uint16(395),
	4942: uint16(1),
	4943: uint16(anon_sym_PIPE_PIPE),
	4944: uint16(437),
	4945: uint16(1),
	4946: uint16(anon_sym_SEMI),
	4947: uint16(3),
	4948: uint16(2),
	4949: uint16(sym_block_comment),
	4950: uint16(sym_line_comment),
	4951: uint16(325),
	4952: uint16(2),
	4953: uint16(anon_sym_LT),
	4954: uint16(anon_sym_GT),
	4955: uint16(331),
	4956: uint16(2),
	4957: uint16(anon_sym_EQ_EQ),
	4958: uint16(anon_sym_BANG_EQ),
	4959: uint16(333),
	4960: uint16(2),
	4961: uint16(anon_sym_LT_EQ),
	4962: uint16(anon_sym_GT_EQ),
	4963: uint16(335),
	4964: uint16(2),
	4965: uint16(anon_sym_LT_LT),
	4966: uint16(anon_sym_GT_GT),
	4967: uint16(337),
	4968: uint16(2),
	4969: uint16(anon_sym_PLUS),
	4970: uint16(anon_sym_DASH),
	4971: uint16(339),
	4972: uint16(2),
	4973: uint16(anon_sym_STAR),
	4974: uint16(anon_sym_PERCENT),
	4975: uint16(16),
	4976: uint16(99),
	4977: uint16(1),
	4978: uint16(anon_sym_SEMI),
	4979: uint16(329),
	4980: uint16(1),
	4981: uint16(anon_sym_AMP),
	4982: uint16(341),
	4983: uint16(1),
	4984: uint16(anon_sym_SLASH),
	4985: uint16(343),
	4986: uint16(1),
	4987: uint16(anon_sym_LBRACK),
	4988: uint16(345),
	4989: uint16(1),
	4990: uint16(anon_sym_DOT),
	4991: uint16(375),
	4992: uint16(1),
	4993: uint16(anon_sym_AMP_AMP),
	4994: uint16(377),
	4995: uint16(1),
	4996: uint16(anon_sym_PIPE),
	4997: uint16(379),
	4998: uint16(1),
	4999: uint16(anon_sym_CARET),
	5000: uint16(395),
	5001: uint16(1),
	5002: uint16(anon_sym_PIPE_PIPE),
	5003: uint16(3),
	5004: uint16(2),
	5005: uint16(sym_block_comment),
	5006: uint16(sym_line_comment),
	5007: uint16(325),
	5008: uint16(2),
	5009: uint16(anon_sym_LT),
	5010: uint16(anon_sym_GT),
	5011: uint16(331),
	5012: uint16(2),
	5013: uint16(anon_sym_EQ_EQ),
	5014: uint16(anon_sym_BANG_EQ),
	5015: uint16(333),
	5016: uint16(2),
	5017: uint16(anon_sym_LT_EQ),
	5018: uint16(anon_sym_GT_EQ),
	5019: uint16(335),
	5020: uint16(2),
	5021: uint16(anon_sym_LT_LT),
	5022: uint16(anon_sym_GT_GT),
	5023: uint16(337),
	5024: uint16(2),
	5025: uint16(anon_sym_PLUS),
	5026: uint16(anon_sym_DASH),
	5027: uint16(339),
	5028: uint16(2),
	5029: uint16(anon_sym_STAR),
	5030: uint16(anon_sym_PERCENT),
	5031: uint16(16),
	5032: uint16(329),
	5033: uint16(1),
	5034: uint16(anon_sym_AMP),
	5035: uint16(341),
	5036: uint16(1),
	5037: uint16(anon_sym_SLASH),
	5038: uint16(343),
	5039: uint16(1),
	5040: uint16(anon_sym_LBRACK),
	5041: uint16(345),
	5042: uint16(1),
	5043: uint16(anon_sym_DOT),
	5044: uint16(375),
	5045: uint16(1),
	5046: uint16(anon_sym_AMP_AMP),
	5047: uint16(377),
	5048: uint16(1),
	5049: uint16(anon_sym_PIPE),
	5050: uint16(379),
	5051: uint16(1),
	5052: uint16(anon_sym_CARET),
	5053: uint16(395),
	5054: uint16(1),
	5055: uint16(anon_sym_PIPE_PIPE),
	5056: uint16(439),
	5057: uint16(1),
	5058: uint16(anon_sym_SEMI),
	5059: uint16(3),
	5060: uint16(2),
	5061: uint16(sym_block_comment),
	5062: uint16(sym_line_comment),
	5063: uint16(325),
	5064: uint16(2),
	5065: uint16(anon_sym_LT),
	5066: uint16(anon_sym_GT),
	5067: uint16(331),
	5068: uint16(2),
	5069: uint16(anon_sym_EQ_EQ),
	5070: uint16(anon_sym_BANG_EQ),
	5071: uint16(333),
	5072: uint16(2),
	5073: uint16(anon_sym_LT_EQ),
	5074: uint16(anon_sym_GT_EQ),
	5075: uint16(335),
	5076: uint16(2),
	5077: uint16(anon_sym_LT_LT),
	5078: uint16(anon_sym_GT_GT),
	5079: uint16(337),
	5080: uint16(2),
	5081: uint16(anon_sym_PLUS),
	5082: uint16(anon_sym_DASH),
	5083: uint16(339),
	5084: uint16(2),
	5085: uint16(anon_sym_STAR),
	5086: uint16(anon_sym_PERCENT),
	5087: uint16(16),
	5088: uint16(329),
	5089: uint16(1),
	5090: uint16(anon_sym_AMP),
	5091: uint16(341),
	5092: uint16(1),
	5093: uint16(anon_sym_SLASH),
	5094: uint16(343),
	5095: uint16(1),
	5096: uint16(anon_sym_LBRACK),
	5097: uint16(345),
	5098: uint16(1),
	5099: uint16(anon_sym_DOT),
	5100: uint16(375),
	5101: uint16(1),
	5102: uint16(anon_sym_AMP_AMP),
	5103: uint16(377),
	5104: uint16(1),
	5105: uint16(anon_sym_PIPE),
	5106: uint16(379),
	5107: uint16(1),
	5108: uint16(anon_sym_CARET),
	5109: uint16(395),
	5110: uint16(1),
	5111: uint16(anon_sym_PIPE_PIPE),
	5112: uint16(441),
	5113: uint16(1),
	5114: uint16(anon_sym_RBRACK),
	5115: uint16(3),
	5116: uint16(2),
	5117: uint16(sym_block_comment),
	5118: uint16(sym_line_comment),
	5119: uint16(325),
	5120: uint16(2),
	5121: uint16(anon_sym_LT),
	5122: uint16(anon_sym_GT),
	5123: uint16(331),
	5124: uint16(2),
	5125: uint16(anon_sym_EQ_EQ),
	5126: uint16(anon_sym_BANG_EQ),
	5127: uint16(333),
	5128: uint16(2),
	5129: uint16(anon_sym_LT_EQ),
	5130: uint16(anon_sym_GT_EQ),
	5131: uint16(335),
	5132: uint16(2),
	5133: uint16(anon_sym_LT_LT),
	5134: uint16(anon_sym_GT_GT),
	5135: uint16(337),
	5136: uint16(2),
	5137: uint16(anon_sym_PLUS),
	5138: uint16(anon_sym_DASH),
	5139: uint16(339),
	5140: uint16(2),
	5141: uint16(anon_sym_STAR),
	5142: uint16(anon_sym_PERCENT),
	5143: uint16(3),
	5144: uint16(3),
	5145: uint16(2),
	5146: uint16(sym_block_comment),
	5147: uint16(sym_line_comment),
	5148: uint16(445),
	5149: uint16(5),
	5150: uint16(anon_sym_LPAREN),
	5151: uint16(anon_sym_LBRACE),
	5152: uint16(anon_sym_RBRACE),
	5153: uint16(anon_sym_AMP),
	5154: uint16(anon_sym_STAR),
	5155: uint16(443),
	5156: uint16(15),
	5157: uint16(anon_sym_let),
	5158: uint16(sym_identifier),
	5159: uint16(anon_sym__),
	5160: uint16(anon_sym_if),
	5161: uint16(anon_sym_switch),
	5162: uint16(anon_sym_fallthrough),
	5163: uint16(anon_sym_loop),
	5164: uint16(anon_sym_for),
	5165: uint16(anon_sym_while),
	5166: uint16(anon_sym_break),
	5167: uint16(anon_sym_continue),
	5168: uint16(anon_sym_continuing),
	5169: uint16(anon_sym_return),
	5170: uint16(anon_sym_discard),
	5171: uint16(anon_sym_var),
	5172: uint16(3),
	5173: uint16(3),
	5174: uint16(2),
	5175: uint16(sym_block_comment),
	5176: uint16(sym_line_comment),
	5177: uint16(449),
	5178: uint16(5),
	5179: uint16(anon_sym_LPAREN),
	5180: uint16(anon_sym_LBRACE),
	5181: uint16(anon_sym_RBRACE),
	5182: uint16(anon_sym_AMP),
	5183: uint16(anon_sym_STAR),
	5184: uint16(447),
	5185: uint16(15),
	5186: uint16(anon_sym_let),
	5187: uint16(sym_identifier),
	5188: uint16(anon_sym__),
	5189: uint16(anon_sym_if),
	5190: uint16(anon_sym_switch),
	5191: uint16(anon_sym_fallthrough),
	5192: uint16(anon_sym_loop),
	5193: uint16(anon_sym_for),
	5194: uint16(anon_sym_while),
	5195: uint16(anon_sym_break),
	5196: uint16(anon_sym_continue),
	5197: uint16(anon_sym_continuing),
	5198: uint16(anon_sym_return),
	5199: uint16(anon_sym_discard),
	5200: uint16(anon_sym_var),
	5201: uint16(3),
	5202: uint16(3),
	5203: uint16(2),
	5204: uint16(sym_block_comment),
	5205: uint16(sym_line_comment),
	5206: uint16(453),
	5207: uint16(5),
	5208: uint16(anon_sym_LPAREN),
	5209: uint16(anon_sym_LBRACE),
	5210: uint16(anon_sym_RBRACE),
	5211: uint16(anon_sym_AMP),
	5212: uint16(anon_sym_STAR),
	5213: uint16(451),
	5214: uint16(15),
	5215: uint16(anon_sym_let),
	5216: uint16(sym_identifier),
	5217: uint16(anon_sym__),
	5218: uint16(anon_sym_if),
	5219: uint16(anon_sym_switch),
	5220: uint16(anon_sym_fallthrough),
	5221: uint16(anon_sym_loop),
	5222: uint16(anon_sym_for),
	5223: uint16(anon_sym_while),
	5224: uint16(anon_sym_break),
	5225: uint16(anon_sym_continue),
	5226: uint16(anon_sym_continuing),
	5227: uint16(anon_sym_return),
	5228: uint16(anon_sym_discard),
	5229: uint16(anon_sym_var),
	5230: uint16(3),
	5231: uint16(3),
	5232: uint16(2),
	5233: uint16(sym_block_comment),
	5234: uint16(sym_line_comment),
	5235: uint16(457),
	5236: uint16(5),
	5237: uint16(anon_sym_LPAREN),
	5238: uint16(anon_sym_LBRACE),
	5239: uint16(anon_sym_RBRACE),
	5240: uint16(anon_sym_AMP),
	5241: uint16(anon_sym_STAR),
	5242: uint16(455),
	5243: uint16(15),
	5244: uint16(anon_sym_let),
	5245: uint16(sym_identifier),
	5246: uint16(anon_sym__),
	5247: uint16(anon_sym_if),
	5248: uint16(anon_sym_switch),
	5249: uint16(anon_sym_fallthrough),
	5250: uint16(anon_sym_loop),
	5251: uint16(anon_sym_for),
	5252: uint16(anon_sym_while),
	5253: uint16(anon_sym_break),
	5254: uint16(anon_sym_continue),
	5255: uint16(anon_sym_continuing),
	5256: uint16(anon_sym_return),
	5257: uint16(anon_sym_discard),
	5258: uint16(anon_sym_var),
	5259: uint16(3),
	5260: uint16(3),
	5261: uint16(2),
	5262: uint16(sym_block_comment),
	5263: uint16(sym_line_comment),
	5264: uint16(461),
	5265: uint16(5),
	5266: uint16(anon_sym_LPAREN),
	5267: uint16(anon_sym_LBRACE),
	5268: uint16(anon_sym_RBRACE),
	5269: uint16(anon_sym_AMP),
	5270: uint16(anon_sym_STAR),
	5271: uint16(459),
	5272: uint16(15),
	5273: uint16(anon_sym_let),
	5274: uint16(sym_identifier),
	5275: uint16(anon_sym__),
	5276: uint16(anon_sym_if),
	5277: uint16(anon_sym_switch),
	5278: uint16(anon_sym_fallthrough),
	5279: uint16(anon_sym_loop),
	5280: uint16(anon_sym_for),
	5281: uint16(anon_sym_while),
	5282: uint16(anon_sym_break),
	5283: uint16(anon_sym_continue),
	5284: uint16(anon_sym_continuing),
	5285: uint16(anon_sym_return),
	5286: uint16(anon_sym_discard),
	5287: uint16(anon_sym_var),
	5288: uint16(3),
	5289: uint16(3),
	5290: uint16(2),
	5291: uint16(sym_block_comment),
	5292: uint16(sym_line_comment),
	5293: uint16(465),
	5294: uint16(5),
	5295: uint16(anon_sym_LPAREN),
	5296: uint16(anon_sym_LBRACE),
	5297: uint16(anon_sym_RBRACE),
	5298: uint16(anon_sym_AMP),
	5299: uint16(anon_sym_STAR),
	5300: uint16(463),
	5301: uint16(15),
	5302: uint16(anon_sym_let),
	5303: uint16(sym_identifier),
	5304: uint16(anon_sym__),
	5305: uint16(anon_sym_if),
	5306: uint16(anon_sym_switch),
	5307: uint16(anon_sym_fallthrough),
	5308: uint16(anon_sym_loop),
	5309: uint16(anon_sym_for),
	5310: uint16(anon_sym_while),
	5311: uint16(anon_sym_break),
	5312: uint16(anon_sym_continue),
	5313: uint16(anon_sym_continuing),
	5314: uint16(anon_sym_return),
	5315: uint16(anon_sym_discard),
	5316: uint16(anon_sym_var),
	5317: uint16(3),
	5318: uint16(3),
	5319: uint16(2),
	5320: uint16(sym_block_comment),
	5321: uint16(sym_line_comment),
	5322: uint16(469),
	5323: uint16(5),
	5324: uint16(anon_sym_LPAREN),
	5325: uint16(anon_sym_LBRACE),
	5326: uint16(anon_sym_RBRACE),
	5327: uint16(anon_sym_AMP),
	5328: uint16(anon_sym_STAR),
	5329: uint16(467),
	5330: uint16(15),
	5331: uint16(anon_sym_let),
	5332: uint16(sym_identifier),
	5333: uint16(anon_sym__),
	5334: uint16(anon_sym_if),
	5335: uint16(anon_sym_switch),
	5336: uint16(anon_sym_fallthrough),
	5337: uint16(anon_sym_loop),
	5338: uint16(anon_sym_for),
	5339: uint16(anon_sym_while),
	5340: uint16(anon_sym_break),
	5341: uint16(anon_sym_continue),
	5342: uint16(anon_sym_continuing),
	5343: uint16(anon_sym_return),
	5344: uint16(anon_sym_discard),
	5345: uint16(anon_sym_var),
	5346: uint16(3),
	5347: uint16(3),
	5348: uint16(2),
	5349: uint16(sym_block_comment),
	5350: uint16(sym_line_comment),
	5351: uint16(473),
	5352: uint16(5),
	5353: uint16(anon_sym_LPAREN),
	5354: uint16(anon_sym_LBRACE),
	5355: uint16(anon_sym_RBRACE),
	5356: uint16(anon_sym_AMP),
	5357: uint16(anon_sym_STAR),
	5358: uint16(471),
	5359: uint16(15),
	5360: uint16(anon_sym_let),
	5361: uint16(sym_identifier),
	5362: uint16(anon_sym__),
	5363: uint16(anon_sym_if),
	5364: uint16(anon_sym_switch),
	5365: uint16(anon_sym_fallthrough),
	5366: uint16(anon_sym_loop),
	5367: uint16(anon_sym_for),
	5368: uint16(anon_sym_while),
	5369: uint16(anon_sym_break),
	5370: uint16(anon_sym_continue),
	5371: uint16(anon_sym_continuing),
	5372: uint16(anon_sym_return),
	5373: uint16(anon_sym_discard),
	5374: uint16(anon_sym_var),
	5375: uint16(3),
	5376: uint16(3),
	5377: uint16(2),
	5378: uint16(sym_block_comment),
	5379: uint16(sym_line_comment),
	5380: uint16(477),
	5381: uint16(5),
	5382: uint16(anon_sym_LPAREN),
	5383: uint16(anon_sym_LBRACE),
	5384: uint16(anon_sym_RBRACE),
	5385: uint16(anon_sym_AMP),
	5386: uint16(anon_sym_STAR),
	5387: uint16(475),
	5388: uint16(15),
	5389: uint16(anon_sym_let),
	5390: uint16(sym_identifier),
	5391: uint16(anon_sym__),
	5392: uint16(anon_sym_if),
	5393: uint16(anon_sym_switch),
	5394: uint16(anon_sym_fallthrough),
	5395: uint16(anon_sym_loop),
	5396: uint16(anon_sym_for),
	5397: uint16(anon_sym_while),
	5398: uint16(anon_sym_break),
	5399: uint16(anon_sym_continue),
	5400: uint16(anon_sym_continuing),
	5401: uint16(anon_sym_return),
	5402: uint16(anon_sym_discard),
	5403: uint16(anon_sym_var),
	5404: uint16(3),
	5405: uint16(3),
	5406: uint16(2),
	5407: uint16(sym_block_comment),
	5408: uint16(sym_line_comment),
	5409: uint16(481),
	5410: uint16(5),
	5411: uint16(anon_sym_LPAREN),
	5412: uint16(anon_sym_LBRACE),
	5413: uint16(anon_sym_RBRACE),
	5414: uint16(anon_sym_AMP),
	5415: uint16(anon_sym_STAR),
	5416: uint16(479),
	5417: uint16(15),
	5418: uint16(anon_sym_let),
	5419: uint16(sym_identifier),
	5420: uint16(anon_sym__),
	5421: uint16(anon_sym_if),
	5422: uint16(anon_sym_switch),
	5423: uint16(anon_sym_fallthrough),
	5424: uint16(anon_sym_loop),
	5425: uint16(anon_sym_for),
	5426: uint16(anon_sym_while),
	5427: uint16(anon_sym_break),
	5428: uint16(anon_sym_continue),
	5429: uint16(anon_sym_continuing),
	5430: uint16(anon_sym_return),
	5431: uint16(anon_sym_discard),
	5432: uint16(anon_sym_var),
	5433: uint16(3),
	5434: uint16(3),
	5435: uint16(2),
	5436: uint16(sym_block_comment),
	5437: uint16(sym_line_comment),
	5438: uint16(485),
	5439: uint16(5),
	5440: uint16(anon_sym_LPAREN),
	5441: uint16(anon_sym_LBRACE),
	5442: uint16(anon_sym_RBRACE),
	5443: uint16(anon_sym_AMP),
	5444: uint16(anon_sym_STAR),
	5445: uint16(483),
	5446: uint16(15),
	5447: uint16(anon_sym_let),
	5448: uint16(sym_identifier),
	5449: uint16(anon_sym__),
	5450: uint16(anon_sym_if),
	5451: uint16(anon_sym_switch),
	5452: uint16(anon_sym_fallthrough),
	5453: uint16(anon_sym_loop),
	5454: uint16(anon_sym_for),
	5455: uint16(anon_sym_while),
	5456: uint16(anon_sym_break),
	5457: uint16(anon_sym_continue),
	5458: uint16(anon_sym_continuing),
	5459: uint16(anon_sym_return),
	5460: uint16(anon_sym_discard),
	5461: uint16(anon_sym_var),
	5462: uint16(3),
	5463: uint16(3),
	5464: uint16(2),
	5465: uint16(sym_block_comment),
	5466: uint16(sym_line_comment),
	5467: uint16(489),
	5468: uint16(5),
	5469: uint16(anon_sym_LPAREN),
	5470: uint16(anon_sym_LBRACE),
	5471: uint16(anon_sym_RBRACE),
	5472: uint16(anon_sym_AMP),
	5473: uint16(anon_sym_STAR),
	5474: uint16(487),
	5475: uint16(15),
	5476: uint16(anon_sym_let),
	5477: uint16(sym_identifier),
	5478: uint16(anon_sym__),
	5479: uint16(anon_sym_if),
	5480: uint16(anon_sym_switch),
	5481: uint16(anon_sym_fallthrough),
	5482: uint16(anon_sym_loop),
	5483: uint16(anon_sym_for),
	5484: uint16(anon_sym_while),
	5485: uint16(anon_sym_break),
	5486: uint16(anon_sym_continue),
	5487: uint16(anon_sym_continuing),
	5488: uint16(anon_sym_return),
	5489: uint16(anon_sym_discard),
	5490: uint16(anon_sym_var),
	5491: uint16(14),
	5492: uint16(491),
	5493: uint16(1),
	5495: uint16(493),
	5496: uint16(1),
	5497: uint16(anon_sym_SEMI),
	5498: uint16(496),
	5499: uint16(1),
	5500: uint16(anon_sym_let),
	5501: uint16(499),
	5502: uint16(1),
	5503: uint16(anon_sym_override),
	5504: uint16(502),
	5505: uint16(1),
	5506: uint16(anon_sym_type),
	5507: uint16(505),
	5508: uint16(1),
	5509: uint16(anon_sym_fn),
	5510: uint16(508),
	5511: uint16(1),
	5512: uint16(anon_sym_struct),
	5513: uint16(511),
	5514: uint16(1),
	5515: uint16(anon_sym_AT),
	5516: uint16(514),
	5517: uint16(1),
	5518: uint16(anon_sym_var),
	5519: uint16(243),
	5520: uint16(1),
	5521: uint16(sym_variable_declaration),
	5522: uint16(3),
	5523: uint16(2),
	5524: uint16(sym_block_comment),
	5525: uint16(sym_line_comment),
	5526: uint16(182),
	5527: uint16(2),
	5528: uint16(sym_attribute),
	5529: uint16(aux_sym_global_variable_declaration_repeat1),
	5530: uint16(342),
	5531: uint16(3),
	5532: uint16(sym_global_variable_declaration),
	5533: uint16(sym_global_constant_declaration),
	5534: uint16(sym_type_alias_declaration),
	5535: uint16(132),
	5536: uint16(4),
	5537: uint16(sym__declaration),
	5538: uint16(sym_function_declaration),
	5539: uint16(sym_struct_declaration),
	5540: uint16(aux_sym_source_file_repeat2),
	5541: uint16(14),
	5542: uint16(9),
	5543: uint16(1),
	5544: uint16(anon_sym_let),
	5545: uint16(11),
	5546: uint16(1),
	5547: uint16(anon_sym_override),
	5548: uint16(13),
	5549: uint16(1),
	5550: uint16(anon_sym_type),
	5551: uint16(15),
	5552: uint16(1),
	5553: uint16(anon_sym_fn),
	5554: uint16(17),
	5555: uint16(1),
	5556: uint16(anon_sym_struct),
	5557: uint16(21),
	5558: uint16(1),
	5559: uint16(anon_sym_AT),
	5560: uint16(23),
	5561: uint16(1),
	5562: uint16(anon_sym_var),
	5563: uint16(389),
	5564: uint16(1),
	5566: uint16(517),
	5567: uint16(1),
	5568: uint16(anon_sym_SEMI),
	5569: uint16(243),
	5570: uint16(1),
	5571: uint16(sym_variable_declaration),
	5572: uint16(3),
	5573: uint16(2),
	5574: uint16(sym_block_comment),
	5575: uint16(sym_line_comment),
	5576: uint16(182),
	5577: uint16(2),
	5578: uint16(sym_attribute),
	5579: uint16(aux_sym_global_variable_declaration_repeat1),
	5580: uint16(342),
	5581: uint16(3),
	5582: uint16(sym_global_variable_declaration),
	5583: uint16(sym_global_constant_declaration),
	5584: uint16(sym_type_alias_declaration),
	5585: uint16(132),
	5586: uint16(4),
	5587: uint16(sym__declaration),
	5588: uint16(sym_function_declaration),
	5589: uint16(sym_struct_declaration),
	5590: uint16(aux_sym_source_file_repeat2),
	5591: uint16(14),
	5592: uint16(9),
	5593: uint16(1),
	5594: uint16(anon_sym_let),
	5595: uint16(11),
	5596: uint16(1),
	5597: uint16(anon_sym_override),
	5598: uint16(13),
	5599: uint16(1),
	5600: uint16(anon_sym_type),
	5601: uint16(15),
	5602: uint16(1),
	5603: uint16(anon_sym_fn),
	5604: uint16(17),
	5605: uint16(1),
	5606: uint16(anon_sym_struct),
	5607: uint16(21),
	5608: uint16(1),
	5609: uint16(anon_sym_AT),
	5610: uint16(23),
	5611: uint16(1),
	5612: uint16(anon_sym_var),
	5613: uint16(517),
	5614: uint16(1),
	5615: uint16(anon_sym_SEMI),
	5616: uint16(519),
	5617: uint16(1),
	5619: uint16(243),
	5620: uint16(1),
	5621: uint16(sym_variable_declaration),
	5622: uint16(3),
	5623: uint16(2),
	5624: uint16(sym_block_comment),
	5625: uint16(sym_line_comment),
	5626: uint16(182),
	5627: uint16(2),
	5628: uint16(sym_attribute),
	5629: uint16(aux_sym_global_variable_declaration_repeat1),
	5630: uint16(342),
	5631: uint16(3),
	5632: uint16(sym_global_variable_declaration),
	5633: uint16(sym_global_constant_declaration),
	5634: uint16(sym_type_alias_declaration),
	5635: uint16(132),
	5636: uint16(4),
	5637: uint16(sym__declaration),
	5638: uint16(sym_function_declaration),
	5639: uint16(sym_struct_declaration),
	5640: uint16(aux_sym_source_file_repeat2),
	5641: uint16(3),
	5642: uint16(335),
	5643: uint16(1),
	5644: uint16(sym_texel_format),
	5645: uint16(3),
	5646: uint16(2),
	5647: uint16(sym_block_comment),
	5648: uint16(sym_line_comment),
	5649: uint16(521),
	5650: uint16(16),
	5651: uint16(anon_sym_rgba8unorm),
	5652: uint16(anon_sym_rgba8snorm),
	5653: uint16(anon_sym_rgba8uint),
	5654: uint16(anon_sym_rgba8sint),
	5655: uint16(anon_sym_rgba16uint),
	5656: uint16(anon_sym_rgba16sint),
	5657: uint16(anon_sym_rgba16float),
	5658: uint16(anon_sym_r32uint),
	5659: uint16(anon_sym_r32sint),
	5660: uint16(anon_sym_r32float),
	5661: uint16(anon_sym_rg32uint),
	5662: uint16(anon_sym_rg32sint),
	5663: uint16(anon_sym_rg32float),
	5664: uint16(anon_sym_rgba32uint),
	5665: uint16(anon_sym_rgba32sint),
	5666: uint16(anon_sym_rgba32float),
	5667: uint16(3),
	5668: uint16(144),
	5669: uint16(1),
	5670: uint16(sym_postfix_expression),
	5671: uint16(3),
	5672: uint16(2),
	5673: uint16(sym_block_comment),
	5674: uint16(sym_line_comment),
	5675: uint16(523),
	5676: uint16(14),
	5677: uint16(anon_sym_EQ),
	5678: uint16(anon_sym_RPAREN),
	5679: uint16(anon_sym_PLUS_EQ),
	5680: uint16(anon_sym_DASH_EQ),
	5681: uint16(anon_sym_STAR_EQ),
	5682: uint16(anon_sym_SLASH_EQ),
	5683: uint16(anon_sym_PERCENT_EQ),
	5684: uint16(anon_sym_AMP_EQ),
	5685: uint16(anon_sym_PIPE_EQ),
	5686: uint16(anon_sym_CARET_EQ),
	5687: uint16(anon_sym_PLUS_PLUS),
	5688: uint16(anon_sym_DASH_DASH),
	5689: uint16(anon_sym_LBRACK),
	5690: uint16(anon_sym_DOT),
	5691: uint16(5),
	5692: uint16(527),
	5693: uint16(1),
	5694: uint16(anon_sym_LBRACK),
	5695: uint16(529),
	5696: uint16(1),
	5697: uint16(anon_sym_DOT),
	5698: uint16(153),
	5699: uint16(1),
	5700: uint16(sym_postfix_expression),
	5701: uint16(3),
	5702: uint16(2),
	5703: uint16(sym_block_comment),
	5704: uint16(sym_line_comment),
	5705: uint16(525),
	5706: uint16(12),
	5707: uint16(anon_sym_EQ),
	5708: uint16(anon_sym_RPAREN),
	5709: uint16(anon_sym_PLUS_EQ),
	5710: uint16(anon_sym_DASH_EQ),
	5711: uint16(anon_sym_STAR_EQ),
	5712: uint16(anon_sym_SLASH_EQ),
	5713: uint16(anon_sym_PERCENT_EQ),
	5714: uint16(anon_sym_AMP_EQ),
	5715: uint16(anon_sym_PIPE_EQ),
	5716: uint16(anon_sym_CARET_EQ),
	5717: uint16(anon_sym_PLUS_PLUS),
	5718: uint16(anon_sym_DASH_DASH),
	5719: uint16(5),
	5720: uint16(531),
	5721: uint16(1),
	5722: uint16(anon_sym_LBRACK),
	5723: uint16(533),
	5724: uint16(1),
	5725: uint16(anon_sym_DOT),
	5726: uint16(145),
	5727: uint16(1),
	5728: uint16(sym_postfix_expression),
	5729: uint16(3),
	5730: uint16(2),
	5731: uint16(sym_block_comment),
	5732: uint16(sym_line_comment),
	5733: uint16(523),
	5734: uint16(12),
	5735: uint16(anon_sym_EQ),
	5736: uint16(anon_sym_RPAREN),
	5737: uint16(anon_sym_PLUS_EQ),
	5738: uint16(anon_sym_DASH_EQ),
	5739: uint16(anon_sym_STAR_EQ),
	5740: uint16(anon_sym_SLASH_EQ),
	5741: uint16(anon_sym_PERCENT_EQ),
	5742: uint16(anon_sym_AMP_EQ),
	5743: uint16(anon_sym_PIPE_EQ),
	5744: uint16(anon_sym_CARET_EQ),
	5745: uint16(anon_sym_PLUS_PLUS),
	5746: uint16(anon_sym_DASH_DASH),
	5747: uint16(5),
	5748: uint16(527),
	5749: uint16(1),
	5750: uint16(anon_sym_LBRACK),
	5751: uint16(529),
	5752: uint16(1),
	5753: uint16(anon_sym_DOT),
	5754: uint16(156),
	5755: uint16(1),
	5756: uint16(sym_postfix_expression),
	5757: uint16(3),
	5758: uint16(2),
	5759: uint16(sym_block_comment),
	5760: uint16(sym_line_comment),
	5761: uint16(535),
	5762: uint16(12),
	5763: uint16(anon_sym_EQ),
	5764: uint16(anon_sym_RPAREN),
	5765: uint16(anon_sym_PLUS_EQ),
	5766: uint16(anon_sym_DASH_EQ),
	5767: uint16(anon_sym_STAR_EQ),
	5768: uint16(anon_sym_SLASH_EQ),
	5769: uint16(anon_sym_PERCENT_EQ),
	5770: uint16(anon_sym_AMP_EQ),
	5771: uint16(anon_sym_PIPE_EQ),
	5772: uint16(anon_sym_CARET_EQ),
	5773: uint16(anon_sym_PLUS_PLUS),
	5774: uint16(anon_sym_DASH_DASH),
	5775: uint16(6),
	5776: uint16(307),
	5777: uint16(1),
	5778: uint16(anon_sym_LPAREN),
	5779: uint16(527),
	5780: uint16(1),
	5781: uint16(anon_sym_LBRACK),
	5782: uint16(529),
	5783: uint16(1),
	5784: uint16(anon_sym_DOT),
	5785: uint16(153),
	5786: uint16(1),
	5787: uint16(sym_postfix_expression),
	5788: uint16(3),
	5789: uint16(2),
	5790: uint16(sym_block_comment),
	5791: uint16(sym_line_comment),
	5792: uint16(525),
	5793: uint16(11),
	5794: uint16(anon_sym_EQ),
	5795: uint16(anon_sym_PLUS_EQ),
	5796: uint16(anon_sym_DASH_EQ),
	5797: uint16(anon_sym_STAR_EQ),
	5798: uint16(anon_sym_SLASH_EQ),
	5799: uint16(anon_sym_PERCENT_EQ),
	5800: uint16(anon_sym_AMP_EQ),
	5801: uint16(anon_sym_PIPE_EQ),
	5802: uint16(anon_sym_CARET_EQ),
	5803: uint16(anon_sym_PLUS_PLUS),
	5804: uint16(anon_sym_DASH_DASH),
	5805: uint16(5),
	5806: uint16(527),
	5807: uint16(1),
	5808: uint16(anon_sym_LBRACK),
	5809: uint16(529),
	5810: uint16(1),
	5811: uint16(anon_sym_DOT),
	5812: uint16(152),
	5813: uint16(1),
	5814: uint16(sym_postfix_expression),
	5815: uint16(3),
	5816: uint16(2),
	5817: uint16(sym_block_comment),
	5818: uint16(sym_line_comment),
	5819: uint16(537),
	5820: uint16(12),
	5821: uint16(anon_sym_EQ),
	5822: uint16(anon_sym_RPAREN),
	5823: uint16(anon_sym_PLUS_EQ),
	5824: uint16(anon_sym_DASH_EQ),
	5825: uint16(anon_sym_STAR_EQ),
	5826: uint16(anon_sym_SLASH_EQ),
	5827: uint16(anon_sym_PERCENT_EQ),
	5828: uint16(anon_sym_AMP_EQ),
	5829: uint16(anon_sym_PIPE_EQ),
	5830: uint16(anon_sym_CARET_EQ),
	5831: uint16(anon_sym_PLUS_PLUS),
	5832: uint16(anon_sym_DASH_DASH),
	5833: uint16(3),
	5834: uint16(148),
	5835: uint16(1),
	5836: uint16(sym_postfix_expression),
	5837: uint16(3),
	5838: uint16(2),
	5839: uint16(sym_block_comment),
	5840: uint16(sym_line_comment),
	5841: uint16(539),
	5842: uint16(14),
	5843: uint16(anon_sym_EQ),
	5844: uint16(anon_sym_RPAREN),
	5845: uint16(anon_sym_PLUS_EQ),
	5846: uint16(anon_sym_DASH_EQ),
	5847: uint16(anon_sym_STAR_EQ),
	5848: uint16(anon_sym_SLASH_EQ),
	5849: uint16(anon_sym_PERCENT_EQ),
	5850: uint16(anon_sym_AMP_EQ),
	5851: uint16(anon_sym_PIPE_EQ),
	5852: uint16(anon_sym_CARET_EQ),
	5853: uint16(anon_sym_PLUS_PLUS),
	5854: uint16(anon_sym_DASH_DASH),
	5855: uint16(anon_sym_LBRACK),
	5856: uint16(anon_sym_DOT),
	5857: uint16(5),
	5858: uint16(531),
	5859: uint16(1),
	5860: uint16(anon_sym_LBRACK),
	5861: uint16(533),
	5862: uint16(1),
	5863: uint16(anon_sym_DOT),
	5864: uint16(146),
	5865: uint16(1),
	5866: uint16(sym_postfix_expression),
	5867: uint16(3),
	5868: uint16(2),
	5869: uint16(sym_block_comment),
	5870: uint16(sym_line_comment),
	5871: uint16(539),
	5872: uint16(12),
	5873: uint16(anon_sym_EQ),
	5874: uint16(anon_sym_RPAREN),
	5875: uint16(anon_sym_PLUS_EQ),
	5876: uint16(anon_sym_DASH_EQ),
	5877: uint16(anon_sym_STAR_EQ),
	5878: uint16(anon_sym_SLASH_EQ),
	5879: uint16(anon_sym_PERCENT_EQ),
	5880: uint16(anon_sym_AMP_EQ),
	5881: uint16(anon_sym_PIPE_EQ),
	5882: uint16(anon_sym_CARET_EQ),
	5883: uint16(anon_sym_PLUS_PLUS),
	5884: uint16(anon_sym_DASH_DASH),
	5885: uint16(3),
	5886: uint16(150),
	5887: uint16(1),
	5888: uint16(sym_postfix_expression),
	5889: uint16(3),
	5890: uint16(2),
	5891: uint16(sym_block_comment),
	5892: uint16(sym_line_comment),
	5893: uint16(539),
	5894: uint16(14),
	5895: uint16(anon_sym_EQ),
	5896: uint16(anon_sym_RPAREN),
	5897: uint16(anon_sym_PLUS_EQ),
	5898: uint16(anon_sym_DASH_EQ),
	5899: uint16(anon_sym_STAR_EQ),
	5900: uint16(anon_sym_SLASH_EQ),
	5901: uint16(anon_sym_PERCENT_EQ),
	5902: uint16(anon_sym_AMP_EQ),
	5903: uint16(anon_sym_PIPE_EQ),
	5904: uint16(anon_sym_CARET_EQ),
	5905: uint16(anon_sym_PLUS_PLUS),
	5906: uint16(anon_sym_DASH_DASH),
	5907: uint16(anon_sym_LBRACK),
	5908: uint16(anon_sym_DOT),
	5909: uint16(5),
	5910: uint16(527),
	5911: uint16(1),
	5912: uint16(anon_sym_LBRACK),
	5913: uint16(529),
	5914: uint16(1),
	5915: uint16(anon_sym_DOT),
	5916: uint16(150),
	5917: uint16(1),
	5918: uint16(sym_postfix_expression),
	5919: uint16(3),
	5920: uint16(2),
	5921: uint16(sym_block_comment),
	5922: uint16(sym_line_comment),
	5923: uint16(539),
	5924: uint16(12),
	5925: uint16(anon_sym_EQ),
	5926: uint16(anon_sym_RPAREN),
	5927: uint16(anon_sym_PLUS_EQ),
	5928: uint16(anon_sym_DASH_EQ),
	5929: uint16(anon_sym_STAR_EQ),
	5930: uint16(anon_sym_SLASH_EQ),
	5931: uint16(anon_sym_PERCENT_EQ),
	5932: uint16(anon_sym_AMP_EQ),
	5933: uint16(anon_sym_PIPE_EQ),
	5934: uint16(anon_sym_CARET_EQ),
	5935: uint16(anon_sym_PLUS_PLUS),
	5936: uint16(anon_sym_DASH_DASH),
	5937: uint16(5),
	5938: uint16(527),
	5939: uint16(1),
	5940: uint16(anon_sym_LBRACK),
	5941: uint16(529),
	5942: uint16(1),
	5943: uint16(anon_sym_DOT),
	5944: uint16(149),
	5945: uint16(1),
	5946: uint16(sym_postfix_expression),
	5947: uint16(3),
	5948: uint16(2),
	5949: uint16(sym_block_comment),
	5950: uint16(sym_line_comment),
	5951: uint16(541),
	5952: uint16(12),
	5953: uint16(anon_sym_EQ),
	5954: uint16(anon_sym_RPAREN),
	5955: uint16(anon_sym_PLUS_EQ),
	5956: uint16(anon_sym_DASH_EQ),
	5957: uint16(anon_sym_STAR_EQ),
	5958: uint16(anon_sym_SLASH_EQ),
	5959: uint16(anon_sym_PERCENT_EQ),
	5960: uint16(anon_sym_AMP_EQ),
	5961: uint16(anon_sym_PIPE_EQ),
	5962: uint16(anon_sym_CARET_EQ),
	5963: uint16(anon_sym_PLUS_PLUS),
	5964: uint16(anon_sym_DASH_DASH),
	5965: uint16(5),
	5966: uint16(527),
	5967: uint16(1),
	5968: uint16(anon_sym_LBRACK),
	5969: uint16(529),
	5970: uint16(1),
	5971: uint16(anon_sym_DOT),
	5972: uint16(151),
	5973: uint16(1),
	5974: uint16(sym_postfix_expression),
	5975: uint16(3),
	5976: uint16(2),
	5977: uint16(sym_block_comment),
	5978: uint16(sym_line_comment),
	5979: uint16(543),
	5980: uint16(12),
	5981: uint16(anon_sym_EQ),
	5982: uint16(anon_sym_RPAREN),
	5983: uint16(anon_sym_PLUS_EQ),
	5984: uint16(anon_sym_DASH_EQ),
	5985: uint16(anon_sym_STAR_EQ),
	5986: uint16(anon_sym_SLASH_EQ),
	5987: uint16(anon_sym_PERCENT_EQ),
	5988: uint16(anon_sym_AMP_EQ),
	5989: uint16(anon_sym_PIPE_EQ),
	5990: uint16(anon_sym_CARET_EQ),
	5991: uint16(anon_sym_PLUS_PLUS),
	5992: uint16(anon_sym_DASH_DASH),
	5993: uint16(3),
	5994: uint16(149),
	5995: uint16(1),
	5996: uint16(sym_postfix_expression),
	5997: uint16(3),
	5998: uint16(2),
	5999: uint16(sym_block_comment),
	6000: uint16(sym_line_comment),
	6001: uint16(541),
	6002: uint16(14),
	6003: uint16(anon_sym_EQ),
	6004: uint16(anon_sym_RPAREN),
	6005: uint16(anon_sym_PLUS_EQ),
	6006: uint16(anon_sym_DASH_EQ),
	6007: uint16(anon_sym_STAR_EQ),
	6008: uint16(anon_sym_SLASH_EQ),
	6009: uint16(anon_sym_PERCENT_EQ),
	6010: uint16(anon_sym_AMP_EQ),
	6011: uint16(anon_sym_PIPE_EQ),
	6012: uint16(anon_sym_CARET_EQ),
	6013: uint16(anon_sym_PLUS_PLUS),
	6014: uint16(anon_sym_DASH_DASH),
	6015: uint16(anon_sym_LBRACK),
	6016: uint16(anon_sym_DOT),
	6017: uint16(2),
	6018: uint16(3),
	6019: uint16(2),
	6020: uint16(sym_block_comment),
	6021: uint16(sym_line_comment),
	6022: uint16(545),
	6023: uint16(14),
	6024: uint16(anon_sym_EQ),
	6025: uint16(anon_sym_RPAREN),
	6026: uint16(anon_sym_PLUS_EQ),
	6027: uint16(anon_sym_DASH_EQ),
	6028: uint16(anon_sym_STAR_EQ),
	6029: uint16(anon_sym_SLASH_EQ),
	6030: uint16(anon_sym_PERCENT_EQ),
	6031: uint16(anon_sym_AMP_EQ),
	6032: uint16(anon_sym_PIPE_EQ),
	6033: uint16(anon_sym_CARET_EQ),
	6034: uint16(anon_sym_PLUS_PLUS),
	6035: uint16(anon_sym_DASH_DASH),
	6036: uint16(anon_sym_LBRACK),
	6037: uint16(anon_sym_DOT),
	6038: uint16(2),
	6039: uint16(3),
	6040: uint16(2),
	6041: uint16(sym_block_comment),
	6042: uint16(sym_line_comment),
	6043: uint16(541),
	6044: uint16(14),
	6045: uint16(anon_sym_EQ),
	6046: uint16(anon_sym_RPAREN),
	6047: uint16(anon_sym_PLUS_EQ),
	6048: uint16(anon_sym_DASH_EQ),
	6049: uint16(anon_sym_STAR_EQ),
	6050: uint16(anon_sym_SLASH_EQ),
	6051: uint16(anon_sym_PERCENT_EQ),
	6052: uint16(anon_sym_AMP_EQ),
	6053: uint16(anon_sym_PIPE_EQ),
	6054: uint16(anon_sym_CARET_EQ),
	6055: uint16(anon_sym_PLUS_PLUS),
	6056: uint16(anon_sym_DASH_DASH),
	6057: uint16(anon_sym_LBRACK),
	6058: uint16(anon_sym_DOT),
	6059: uint16(2),
	6060: uint16(3),
	6061: uint16(2),
	6062: uint16(sym_block_comment),
	6063: uint16(sym_line_comment),
	6064: uint16(547),
	6065: uint16(12),
	6066: uint16(anon_sym_EQ),
	6067: uint16(anon_sym_RPAREN),
	6068: uint16(anon_sym_PLUS_EQ),
	6069: uint16(anon_sym_DASH_EQ),
	6070: uint16(anon_sym_STAR_EQ),
	6071: uint16(anon_sym_SLASH_EQ),
	6072: uint16(anon_sym_PERCENT_EQ),
	6073: uint16(anon_sym_AMP_EQ),
	6074: uint16(anon_sym_PIPE_EQ),
	6075: uint16(anon_sym_CARET_EQ),
	6076: uint16(anon_sym_PLUS_PLUS),
	6077: uint16(anon_sym_DASH_DASH),
	6078: uint16(2),
	6079: uint16(3),
	6080: uint16(2),
	6081: uint16(sym_block_comment),
	6082: uint16(sym_line_comment),
	6083: uint16(543),
	6084: uint16(12),
	6085: uint16(anon_sym_EQ),
	6086: uint16(anon_sym_RPAREN),
	6087: uint16(anon_sym_PLUS_EQ),
	6088: uint16(anon_sym_DASH_EQ),
	6089: uint16(anon_sym_STAR_EQ),
	6090: uint16(anon_sym_SLASH_EQ),
	6091: uint16(anon_sym_PERCENT_EQ),
	6092: uint16(anon_sym_AMP_EQ),
	6093: uint16(anon_sym_PIPE_EQ),
	6094: uint16(anon_sym_CARET_EQ),
	6095: uint16(anon_sym_PLUS_PLUS),
	6096: uint16(anon_sym_DASH_DASH),
	6097: uint16(2),
	6098: uint16(3),
	6099: uint16(2),
	6100: uint16(sym_block_comment),
	6101: uint16(sym_line_comment),
	6102: uint16(535),
	6103: uint16(12),
	6104: uint16(anon_sym_EQ),
	6105: uint16(anon_sym_RPAREN),
	6106: uint16(anon_sym_PLUS_EQ),
	6107: uint16(anon_sym_DASH_EQ),
	6108: uint16(anon_sym_STAR_EQ),
	6109: uint16(anon_sym_SLASH_EQ),
	6110: uint16(anon_sym_PERCENT_EQ),
	6111: uint16(anon_sym_AMP_EQ),
	6112: uint16(anon_sym_PIPE_EQ),
	6113: uint16(anon_sym_CARET_EQ),
	6114: uint16(anon_sym_PLUS_PLUS),
	6115: uint16(anon_sym_DASH_DASH),
	6116: uint16(6),
	6117: uint16(549),
	6118: uint16(1),
	6119: uint16(anon_sym_EQ),
	6120: uint16(553),
	6121: uint16(1),
	6122: uint16(anon_sym_PLUS_PLUS),
	6123: uint16(555),
	6124: uint16(1),
	6125: uint16(anon_sym_DASH_DASH),
	6126: uint16(10),
	6127: uint16(1),
	6128: uint16(sym_compound_assignment_operator),
	6129: uint16(3),
	6130: uint16(2),
	6131: uint16(sym_block_comment),
	6132: uint16(sym_line_comment),
	6133: uint16(551),
	6134: uint16(8),
	6135: uint16(anon_sym_PLUS_EQ),
	6136: uint16(anon_sym_DASH_EQ),
	6137: uint16(anon_sym_STAR_EQ),
	6138: uint16(anon_sym_SLASH_EQ),
	6139: uint16(anon_sym_PERCENT_EQ),
	6140: uint16(anon_sym_AMP_EQ),
	6141: uint16(anon_sym_PIPE_EQ),
	6142: uint16(anon_sym_CARET_EQ),
	6143: uint16(4),
	6144: uint16(559),
	6145: uint16(1),
	6146: uint16(anon_sym_enable),
	6147: uint16(3),
	6148: uint16(2),
	6149: uint16(sym_block_comment),
	6150: uint16(sym_line_comment),
	6151: uint16(155),
	6152: uint16(2),
	6153: uint16(sym_enable_directive),
	6154: uint16(aux_sym_source_file_repeat1),
	6155: uint16(557),
	6156: uint16(9),
	6158: uint16(anon_sym_SEMI),
	6159: uint16(anon_sym_let),
	6160: uint16(anon_sym_override),
	6161: uint16(anon_sym_type),
	6162: uint16(anon_sym_fn),
	6163: uint16(anon_sym_struct),
	6164: uint16(anon_sym_AT),
	6165: uint16(anon_sym_var),
	6166: uint16(2),
	6167: uint16(3),
	6168: uint16(2),
	6169: uint16(sym_block_comment),
	6170: uint16(sym_line_comment),
	6171: uint16(537),
	6172: uint16(12),
	6173: uint16(anon_sym_EQ),
	6174: uint16(anon_sym_RPAREN),
	6175: uint16(anon_sym_PLUS_EQ),
	6176: uint16(anon_sym_DASH_EQ),
	6177: uint16(anon_sym_STAR_EQ),
	6178: uint16(anon_sym_SLASH_EQ),
	6179: uint16(anon_sym_PERCENT_EQ),
	6180: uint16(anon_sym_AMP_EQ),
	6181: uint16(anon_sym_PIPE_EQ),
	6182: uint16(anon_sym_CARET_EQ),
	6183: uint16(anon_sym_PLUS_PLUS),
	6184: uint16(anon_sym_DASH_DASH),
	6185: uint16(7),
	6186: uint16(31),
	6187: uint16(1),
	6188: uint16(sym_int_literal),
	6189: uint16(236),
	6190: uint16(1),
	6191: uint16(sym_const_literal),
	6192: uint16(3),
	6193: uint16(2),
	6194: uint16(sym_block_comment),
	6195: uint16(sym_line_comment),
	6196: uint16(33),
	6197: uint16(2),
	6198: uint16(aux_sym_float_literal_token1),
	6199: uint16(aux_sym_float_literal_token2),
	6200: uint16(562),
	6201: uint16(2),
	6202: uint16(anon_sym_LBRACE),
	6203: uint16(anon_sym_COLON),
	6204: uint16(564),
	6205: uint16(2),
	6206: uint16(anon_sym_true),
	6207: uint16(anon_sym_false),
	6208: uint16(75),
	6209: uint16(2),
	6210: uint16(sym_float_literal),
	6211: uint16(sym_bool_literal),
	6212: uint16(7),
	6213: uint16(31),
	6214: uint16(1),
	6215: uint16(sym_int_literal),
	6216: uint16(236),
	6217: uint16(1),
	6218: uint16(sym_const_literal),
	6219: uint16(3),
	6220: uint16(2),
	6221: uint16(sym_block_comment),
	6222: uint16(sym_line_comment),
	6223: uint16(33),
	6224: uint16(2),
	6225: uint16(aux_sym_float_literal_token1),
	6226: uint16(aux_sym_float_literal_token2),
	6227: uint16(564),
	6228: uint16(2),
	6229: uint16(anon_sym_true),
	6230: uint16(anon_sym_false),
	6231: uint16(566),
	6232: uint16(2),
	6233: uint16(anon_sym_LBRACE),
	6234: uint16(anon_sym_COLON),
	6235: uint16(75),
	6236: uint16(2),
	6237: uint16(sym_float_literal),
	6238: uint16(sym_bool_literal),
	6239: uint16(2),
	6240: uint16(3),
	6241: uint16(2),
	6242: uint16(sym_block_comment),
	6243: uint16(sym_line_comment),
	6244: uint16(568),
	6245: uint16(10),
	6247: uint16(anon_sym_SEMI),
	6248: uint16(anon_sym_let),
	6249: uint16(anon_sym_override),
	6250: uint16(anon_sym_type),
	6251: uint16(anon_sym_fn),
	6252: uint16(anon_sym_struct),
	6253: uint16(anon_sym_enable),
	6254: uint16(anon_sym_AT),
	6255: uint16(anon_sym_var),
	6256: uint16(2),
	6257: uint16(3),
	6258: uint16(2),
	6259: uint16(sym_block_comment),
	6260: uint16(sym_line_comment),
	6261: uint16(570),
	6262: uint16(9),
	6264: uint16(anon_sym_SEMI),
	6265: uint16(anon_sym_let),
	6266: uint16(anon_sym_override),
	6267: uint16(anon_sym_type),
	6268: uint16(anon_sym_fn),
	6269: uint16(anon_sym_struct),
	6270: uint16(anon_sym_AT),
	6271: uint16(anon_sym_var),
	6272: uint16(2),
	6273: uint16(3),
	6274: uint16(2),
	6275: uint16(sym_block_comment),
	6276: uint16(sym_line_comment),
	6277: uint16(572),
	6278: uint16(9),
	6280: uint16(anon_sym_SEMI),
	6281: uint16(anon_sym_let),
	6282: uint16(anon_sym_override),
	6283: uint16(anon_sym_type),
	6284: uint16(anon_sym_fn),
	6285: uint16(anon_sym_struct),
	6286: uint16(anon_sym_AT),
	6287: uint16(anon_sym_var),
	6288: uint16(9),
	6289: uint16(21),
	6290: uint16(1),
	6291: uint16(anon_sym_AT),
	6292: uint16(574),
	6293: uint16(1),
	6294: uint16(sym_identifier),
	6295: uint16(576),
	6296: uint16(1),
	6297: uint16(anon_sym_RPAREN),
	6298: uint16(187),
	6299: uint16(1),
	6300: uint16(aux_sym_parameter_list_repeat1),
	6301: uint16(275),
	6302: uint16(1),
	6303: uint16(sym_variable_identifier_declaration),
	6304: uint16(276),
	6305: uint16(1),
	6306: uint16(sym_parameter),
	6307: uint16(305),
	6308: uint16(1),
	6309: uint16(sym_parameter_list),
	6310: uint16(3),
	6311: uint16(2),
	6312: uint16(sym_block_comment),
	6313: uint16(sym_line_comment),
	6314: uint16(194),
	6315: uint16(2),
	6316: uint16(sym_attribute),
	6317: uint16(aux_sym_global_variable_declaration_repeat1),
	6318: uint16(2),
	6319: uint16(3),
	6320: uint16(2),
	6321: uint16(sym_block_comment),
	6322: uint16(sym_line_comment),
	6323: uint16(578),
	6324: uint16(9),
	6326: uint16(anon_sym_SEMI),
	6327: uint16(anon_sym_let),
	6328: uint16(anon_sym_override),
	6329: uint16(anon_sym_type),
	6330: uint16(anon_sym_fn),
	6331: uint16(anon_sym_struct),
	6332: uint16(anon_sym_AT),
	6333: uint16(anon_sym_var),
	6334: uint16(2),
	6335: uint16(3),
	6336: uint16(2),
	6337: uint16(sym_block_comment),
	6338: uint16(sym_line_comment),
	6339: uint16(580),
	6340: uint16(9),
	6342: uint16(anon_sym_SEMI),
	6343: uint16(anon_sym_let),
	6344: uint16(anon_sym_override),
	6345: uint16(anon_sym_type),
	6346: uint16(anon_sym_fn),
	6347: uint16(anon_sym_struct),
	6348: uint16(anon_sym_AT),
	6349: uint16(anon_sym_var),
	6350: uint16(2),
	6351: uint16(3),
	6352: uint16(2),
	6353: uint16(sym_block_comment),
	6354: uint16(sym_line_comment),
	6355: uint16(582),
	6356: uint16(9),
	6358: uint16(anon_sym_SEMI),
	6359: uint16(anon_sym_let),
	6360: uint16(anon_sym_override),
	6361: uint16(anon_sym_type),
	6362: uint16(anon_sym_fn),
	6363: uint16(anon_sym_struct),
	6364: uint16(anon_sym_AT),
	6365: uint16(anon_sym_var),
	6366: uint16(2),
	6367: uint16(3),
	6368: uint16(2),
	6369: uint16(sym_block_comment),
	6370: uint16(sym_line_comment),
	6371: uint16(584),
	6372: uint16(9),
	6374: uint16(anon_sym_SEMI),
	6375: uint16(anon_sym_let),
	6376: uint16(anon_sym_override),
	6377: uint16(anon_sym_type),
	6378: uint16(anon_sym_fn),
	6379: uint16(anon_sym_struct),
	6380: uint16(anon_sym_AT),
	6381: uint16(anon_sym_var),
	6382: uint16(2),
	6383: uint16(3),
	6384: uint16(2),
	6385: uint16(sym_block_comment),
	6386: uint16(sym_line_comment),
	6387: uint16(586),
	6388: uint16(9),
	6390: uint16(anon_sym_SEMI),
	6391: uint16(anon_sym_let),
	6392: uint16(anon_sym_override),
	6393: uint16(anon_sym_type),
	6394: uint16(anon_sym_fn),
	6395: uint16(anon_sym_struct),
	6396: uint16(anon_sym_AT),
	6397: uint16(anon_sym_var),
	6398: uint16(2),
	6399: uint16(3),
	6400: uint16(2),
	6401: uint16(sym_block_comment),
	6402: uint16(sym_line_comment),
	6403: uint16(588),
	6404: uint16(9),
	6406: uint16(anon_sym_SEMI),
	6407: uint16(anon_sym_let),
	6408: uint16(anon_sym_override),
	6409: uint16(anon_sym_type),
	6410: uint16(anon_sym_fn),
	6411: uint16(anon_sym_struct),
	6412: uint16(anon_sym_AT),
	6413: uint16(anon_sym_var),
	6414: uint16(2),
	6415: uint16(3),
	6416: uint16(2),
	6417: uint16(sym_block_comment),
	6418: uint16(sym_line_comment),
	6419: uint16(590),
	6420: uint16(9),
	6422: uint16(anon_sym_SEMI),
	6423: uint16(anon_sym_let),
	6424: uint16(anon_sym_override),
	6425: uint16(anon_sym_type),
	6426: uint16(anon_sym_fn),
	6427: uint16(anon_sym_struct),
	6428: uint16(anon_sym_AT),
	6429: uint16(anon_sym_var),
	6430: uint16(2),
	6431: uint16(3),
	6432: uint16(2),
	6433: uint16(sym_block_comment),
	6434: uint16(sym_line_comment),
	6435: uint16(592),
	6436: uint16(9),
	6438: uint16(anon_sym_SEMI),
	6439: uint16(anon_sym_let),
	6440: uint16(anon_sym_override),
	6441: uint16(anon_sym_type),
	6442: uint16(anon_sym_fn),
	6443: uint16(anon_sym_struct),
	6444: uint16(anon_sym_AT),
	6445: uint16(anon_sym_var),
	6446: uint16(2),
	6447: uint16(3),
	6448: uint16(2),
	6449: uint16(sym_block_comment),
	6450: uint16(sym_line_comment),
	6451: uint16(594),
	6452: uint16(9),
	6454: uint16(anon_sym_SEMI),
	6455: uint16(anon_sym_let),
	6456: uint16(anon_sym_override),
	6457: uint16(anon_sym_type),
	6458: uint16(anon_sym_fn),
	6459: uint16(anon_sym_struct),
	6460: uint16(anon_sym_AT),
	6461: uint16(anon_sym_var),
	6462: uint16(9),
	6463: uint16(21),
	6464: uint16(1),
	6465: uint16(anon_sym_AT),
	6466: uint16(574),
	6467: uint16(1),
	6468: uint16(sym_identifier),
	6469: uint16(596),
	6470: uint16(1),
	6471: uint16(anon_sym_RPAREN),
	6472: uint16(187),
	6473: uint16(1),
	6474: uint16(aux_sym_parameter_list_repeat1),
	6475: uint16(275),
	6476: uint16(1),
	6477: uint16(sym_variable_identifier_declaration),
	6478: uint16(276),
	6479: uint16(1),
	6480: uint16(sym_parameter),
	6481: uint16(339),
	6482: uint16(1),
	6483: uint16(sym_parameter_list),
	6484: uint16(3),
	6485: uint16(2),
	6486: uint16(sym_block_comment),
	6487: uint16(sym_line_comment),
	6488: uint16(194),
	6489: uint16(2),
	6490: uint16(sym_attribute),
	6491: uint16(aux_sym_global_variable_declaration_repeat1),
	6492: uint16(7),
	6493: uint16(31),
	6494: uint16(1),
	6495: uint16(sym_int_literal),
	6496: uint16(204),
	6497: uint16(1),
	6498: uint16(sym_const_literal),
	6499: uint16(221),
	6500: uint16(1),
	6501: uint16(sym_case_selectors),
	6502: uint16(3),
	6503: uint16(2),
	6504: uint16(sym_block_comment),
	6505: uint16(sym_line_comment),
	6506: uint16(33),
	6507: uint16(2),
	6508: uint16(aux_sym_float_literal_token1),
	6509: uint16(aux_sym_float_literal_token2),
	6510: uint16(564),
	6511: uint16(2),
	6512: uint16(anon_sym_true),
	6513: uint16(anon_sym_false),
	6514: uint16(75),
	6515: uint16(2),
	6516: uint16(sym_float_literal),
	6517: uint16(sym_bool_literal),
	6518: uint16(2),
	6519: uint16(3),
	6520: uint16(2),
	6521: uint16(sym_block_comment),
	6522: uint16(sym_line_comment),
	6523: uint16(598),
	6524: uint16(9),
	6526: uint16(anon_sym_SEMI),
	6527: uint16(anon_sym_let),
	6528: uint16(anon_sym_override),
	6529: uint16(anon_sym_type),
	6530: uint16(anon_sym_fn),
	6531: uint16(anon_sym_struct),
	6532: uint16(anon_sym_AT),
	6533: uint16(anon_sym_var),
	6534: uint16(6),
	6535: uint16(31),
	6536: uint16(1),
	6537: uint16(sym_int_literal),
	6538: uint16(236),
	6539: uint16(1),
	6540: uint16(sym_const_literal),
	6541: uint16(3),
	6542: uint16(2),
	6543: uint16(sym_block_comment),
	6544: uint16(sym_line_comment),
	6545: uint16(33),
	6546: uint16(2),
	6547: uint16(aux_sym_float_literal_token1),
	6548: uint16(aux_sym_float_literal_token2),
	6549: uint16(564),
	6550: uint16(2),
	6551: uint16(anon_sym_true),
	6552: uint16(anon_sym_false),
	6553: uint16(75),
	6554: uint16(2),
	6555: uint16(sym_float_literal),
	6556: uint16(sym_bool_literal),
	6557: uint16(2),
	6558: uint16(3),
	6559: uint16(2),
	6560: uint16(sym_block_comment),
	6561: uint16(sym_line_comment),
	6562: uint16(307),
	6563: uint16(8),
	6564: uint16(anon_sym_SEMI),
	6565: uint16(anon_sym_EQ),
	6566: uint16(anon_sym_LPAREN),
	6567: uint16(anon_sym_COMMA),
	6568: uint16(anon_sym_RPAREN),
	6569: uint16(anon_sym_LBRACE),
	6570: uint16(anon_sym_RBRACE),
	6571: uint16(anon_sym_GT),
	6572: uint16(2),
	6573: uint16(3),
	6574: uint16(2),
	6575: uint16(sym_block_comment),
	6576: uint16(sym_line_comment),
	6577: uint16(600),
	6578: uint16(8),
	6579: uint16(anon_sym_SEMI),
	6580: uint16(anon_sym_EQ),
	6581: uint16(anon_sym_LPAREN),
	6582: uint16(anon_sym_COMMA),
	6583: uint16(anon_sym_RPAREN),
	6584: uint16(anon_sym_LBRACE),
	6585: uint16(anon_sym_RBRACE),
	6586: uint16(anon_sym_GT),
	6587: uint16(2),
	6588: uint16(3),
	6589: uint16(2),
	6590: uint16(sym_block_comment),
	6591: uint16(sym_line_comment),
	6592: uint16(602),
	6593: uint16(8),
	6594: uint16(anon_sym_SEMI),
	6595: uint16(anon_sym_EQ),
	6596: uint16(anon_sym_LPAREN),
	6597: uint16(anon_sym_COMMA),
	6598: uint16(anon_sym_RPAREN),
	6599: uint16(anon_sym_LBRACE),
	6600: uint16(anon_sym_RBRACE),
	6601: uint16(anon_sym_GT),
	6602: uint16(2),
	6603: uint16(3),
	6604: uint16(2),
	6605: uint16(sym_block_comment),
	6606: uint16(sym_line_comment),
	6607: uint16(604),
	6608: uint16(8),
	6609: uint16(anon_sym_SEMI),
	6610: uint16(anon_sym_EQ),
	6611: uint16(anon_sym_LPAREN),
	6612: uint16(anon_sym_COMMA),
	6613: uint16(anon_sym_RPAREN),
	6614: uint16(anon_sym_LBRACE),
	6615: uint16(anon_sym_RBRACE),
	6616: uint16(anon_sym_GT),
	6617: uint16(7),
	6618: uint16(606),
	6619: uint16(1),
	6620: uint16(sym_identifier),
	6621: uint16(609),
	6622: uint16(1),
	6623: uint16(anon_sym_AT),
	6624: uint16(180),
	6625: uint16(1),
	6626: uint16(aux_sym_struct_declaration_repeat1),
	6627: uint16(271),
	6628: uint16(1),
	6629: uint16(sym_variable_identifier_declaration),
	6630: uint16(328),
	6631: uint16(1),
	6632: uint16(sym_struct_member),
	6633: uint16(3),
	6634: uint16(2),
	6635: uint16(sym_block_comment),
	6636: uint16(sym_line_comment),
	6637: uint16(197),
	6638: uint16(2),
	6639: uint16(sym_attribute),
	6640: uint16(aux_sym_global_variable_declaration_repeat1),
	6641: uint16(7),
	6642: uint16(21),
	6643: uint16(1),
	6644: uint16(anon_sym_AT),
	6645: uint16(574),
	6646: uint16(1),
	6647: uint16(sym_identifier),
	6648: uint16(186),
	6649: uint16(1),
	6650: uint16(aux_sym_struct_declaration_repeat1),
	6651: uint16(271),
	6652: uint16(1),
	6653: uint16(sym_variable_identifier_declaration),
	6654: uint16(272),
	6655: uint16(1),
	6656: uint16(sym_struct_member),
	6657: uint16(3),
	6658: uint16(2),
	6659: uint16(sym_block_comment),
	6660: uint16(sym_line_comment),
	6661: uint16(197),
	6662: uint16(2),
	6663: uint16(sym_attribute),
	6664: uint16(aux_sym_global_variable_declaration_repeat1),
	6665: uint16(7),
	6666: uint16(21),
	6667: uint16(1),
	6668: uint16(anon_sym_AT),
	6669: uint16(23),
	6670: uint16(1),
	6671: uint16(anon_sym_var),
	6672: uint16(612),
	6673: uint16(1),
	6674: uint16(anon_sym_override),
	6675: uint16(614),
	6676: uint16(1),
	6677: uint16(anon_sym_fn),
	6678: uint16(266),
	6679: uint16(1),
	6680: uint16(sym_variable_declaration),
	6681: uint16(3),
	6682: uint16(2),
	6683: uint16(sym_block_comment),
	6684: uint16(sym_line_comment),
	6685: uint16(47),
	6686: uint16(2),
	6687: uint16(sym_attribute),
	6688: uint16(aux_sym_global_variable_declaration_repeat1),
	6689: uint16(6),
	6690: uint16(616),
	6691: uint16(1),
	6692: uint16(sym_identifier),
	6693: uint16(618),
	6694: uint16(1),
	6695: uint16(sym_int_literal),
	6696: uint16(185),
	6697: uint16(1),
	6698: uint16(aux_sym_attribute_repeat1),
	6699: uint16(3),
	6700: uint16(2),
	6701: uint16(sym_block_comment),
	6702: uint16(sym_line_comment),
	6703: uint16(33),
	6704: uint16(2),
	6705: uint16(aux_sym_float_literal_token1),
	6706: uint16(aux_sym_float_literal_token2),
	6707: uint16(282),
	6708: uint16(2),
	6709: uint16(sym__literal_or_identifier),
	6710: uint16(sym_float_literal),
	6711: uint16(6),
	6712: uint16(620),
	6713: uint16(1),
	6714: uint16(sym_identifier),
	6715: uint16(622),
	6716: uint16(1),
	6717: uint16(sym_int_literal),
	6718: uint16(183),
	6719: uint16(1),
	6720: uint16(aux_sym_attribute_repeat1),
	6721: uint16(3),
	6722: uint16(2),
	6723: uint16(sym_block_comment),
	6724: uint16(sym_line_comment),
	6725: uint16(33),
	6726: uint16(2),
	6727: uint16(aux_sym_float_literal_token1),
	6728: uint16(aux_sym_float_literal_token2),
	6729: uint16(270),
	6730: uint16(2),
	6731: uint16(sym__literal_or_identifier),
	6732: uint16(sym_float_literal),
	6733: uint16(6),
	6734: uint16(624),
	6735: uint16(1),
	6736: uint16(sym_identifier),
	6737: uint16(627),
	6738: uint16(1),
	6739: uint16(sym_int_literal),
	6740: uint16(185),
	6741: uint16(1),
	6742: uint16(aux_sym_attribute_repeat1),
	6743: uint16(3),
	6744: uint16(2),
	6745: uint16(sym_block_comment),
	6746: uint16(sym_line_comment),
	6747: uint16(630),
	6748: uint16(2),
	6749: uint16(aux_sym_float_literal_token1),
	6750: uint16(aux_sym_float_literal_token2),
	6751: uint16(285),
	6752: uint16(2),
	6753: uint16(sym__literal_or_identifier),
	6754: uint16(sym_float_literal),
	6755: uint16(7),
	6756: uint16(21),
	6757: uint16(1),
	6758: uint16(anon_sym_AT),
	6759: uint16(574),
	6760: uint16(1),
	6761: uint16(sym_identifier),
	6762: uint16(180),
	6763: uint16(1),
	6764: uint16(aux_sym_struct_declaration_repeat1),
	6765: uint16(260),
	6766: uint16(1),
	6767: uint16(sym_struct_member),
	6768: uint16(271),
	6769: uint16(1),
	6770: uint16(sym_variable_identifier_declaration),
	6771: uint16(3),
	6772: uint16(2),
	6773: uint16(sym_block_comment),
	6774: uint16(sym_line_comment),
	6775: uint16(197),
	6776: uint16(2),
	6777: uint16(sym_attribute),
	6778: uint16(aux_sym_global_variable_declaration_repeat1),
	6779: uint16(7),
	6780: uint16(21),
	6781: uint16(1),
	6782: uint16(anon_sym_AT),
	6783: uint16(574),
	6784: uint16(1),
	6785: uint16(sym_identifier),
	6786: uint16(188),
	6787: uint16(1),
	6788: uint16(aux_sym_parameter_list_repeat1),
	6789: uint16(262),
	6790: uint16(1),
	6791: uint16(sym_parameter),
	6792: uint16(275),
	6793: uint16(1),
	6794: uint16(sym_variable_identifier_declaration),
	6795: uint16(3),
	6796: uint16(2),
	6797: uint16(sym_block_comment),
	6798: uint16(sym_line_comment),
	6799: uint16(194),
	6800: uint16(2),
	6801: uint16(sym_attribute),
	6802: uint16(aux_sym_global_variable_declaration_repeat1),
	6803: uint16(7),
	6804: uint16(633),
	6805: uint16(1),
	6806: uint16(sym_identifier),
	6807: uint16(636),
	6808: uint16(1),
	6809: uint16(anon_sym_AT),
	6810: uint16(188),
	6811: uint16(1),
	6812: uint16(aux_sym_parameter_list_repeat1),
	6813: uint16(275),
	6814: uint16(1),
	6815: uint16(sym_variable_identifier_declaration),
	6816: uint16(329),
	6817: uint16(1),
	6818: uint16(sym_parameter),
	6819: uint16(3),
	6820: uint16(2),
	6821: uint16(sym_block_comment),
	6822: uint16(sym_line_comment),
	6823: uint16(194),
	6824: uint16(2),
	6825: uint16(sym_attribute),
	6826: uint16(aux_sym_global_variable_declaration_repeat1),
	6827: uint16(3),
	6828: uint16(327),
	6829: uint16(1),
	6830: uint16(sym_address_space),
	6831: uint16(3),
	6832: uint16(2),
	6833: uint16(sym_block_comment),
	6834: uint16(sym_line_comment),
	6835: uint16(639),
	6836: uint16(5),
	6837: uint16(anon_sym_function),
	6838: uint16(anon_sym_private),
	6839: uint16(anon_sym_workgroup),
	6840: uint16(anon_sym_uniform),
	6841: uint16(anon_sym_storage),
	6842: uint16(3),
	6843: uint16(274),
	6844: uint16(1),
	6845: uint16(sym_address_space),
	6846: uint16(3),
	6847: uint16(2),
	6848: uint16(sym_block_comment),
	6849: uint16(sym_line_comment),
	6850: uint16(639),
	6851: uint16(5),
	6852: uint16(anon_sym_function),
	6853: uint16(anon_sym_private),
	6854: uint16(anon_sym_workgroup),
	6855: uint16(anon_sym_uniform),
	6856: uint16(anon_sym_storage),
	6857: uint16(6),
	6858: uint16(109),
	6859: uint16(1),
	6860: uint16(anon_sym_LPAREN),
	6861: uint16(641),
	6862: uint16(1),
	6863: uint16(sym_identifier),
	6864: uint16(202),
	6865: uint16(1),
	6866: uint16(aux_sym_lhs_expression_repeat1),
	6867: uint16(287),
	6868: uint16(1),
	6869: uint16(sym_lhs_expression),
	6870: uint16(3),
	6871: uint16(2),
	6872: uint16(sym_block_comment),
	6873: uint16(sym_line_comment),
	6874: uint16(115),
	6875: uint16(2),
	6876: uint16(anon_sym_AMP),
	6877: uint16(anon_sym_STAR),
	6878: uint16(6),
	6879: uint16(109),
	6880: uint16(1),
	6881: uint16(anon_sym_LPAREN),
	6882: uint16(641),
	6883: uint16(1),
	6884: uint16(sym_identifier),
	6885: uint16(202),
	6886: uint16(1),
	6887: uint16(aux_sym_lhs_expression_repeat1),
	6888: uint16(306),
	6889: uint16(1),
	6890: uint16(sym_lhs_expression),
	6891: uint16(3),
	6892: uint16(2),
	6893: uint16(sym_block_comment),
	6894: uint16(sym_line_comment),
	6895: uint16(115),
	6896: uint16(2),
	6897: uint16(anon_sym_AMP),
	6898: uint16(anon_sym_STAR),
	6899: uint16(5),
	6900: uint16(643),
	6901: uint16(1),
	6902: uint16(anon_sym_RBRACE),
	6903: uint16(645),
	6904: uint16(1),
	6905: uint16(anon_sym_case),
	6906: uint16(648),
	6907: uint16(1),
	6908: uint16(anon_sym_default),
	6909: uint16(3),
	6910: uint16(2),
	6911: uint16(sym_block_comment),
	6912: uint16(sym_line_comment),
	6913: uint16(193),
	6914: uint16(2),
	6915: uint16(sym_switch_body),
	6916: uint16(aux_sym_switch_statement_repeat1),
	6917: uint16(5),
	6918: uint16(21),
	6919: uint16(1),
	6920: uint16(anon_sym_AT),
	6921: uint16(574),
	6922: uint16(1),
	6923: uint16(sym_identifier),
	6924: uint16(263),
	6925: uint16(1),
	6926: uint16(sym_variable_identifier_declaration),
	6927: uint16(3),
	6928: uint16(2),
	6929: uint16(sym_block_comment),
	6930: uint16(sym_line_comment),
	6931: uint16(47),
	6932: uint16(2),
	6933: uint16(sym_attribute),
	6934: uint16(aux_sym_global_variable_declaration_repeat1),
	6935: uint16(4),
	6936: uint16(653),
	6937: uint16(1),
	6938: uint16(anon_sym_RPAREN),
	6939: uint16(655),
	6940: uint16(1),
	6941: uint16(sym_int_literal),
	6942: uint16(3),
	6943: uint16(2),
	6944: uint16(sym_block_comment),
	6945: uint16(sym_line_comment),
	6946: uint16(651),
	6947: uint16(3),
	6948: uint16(sym_identifier),
	6949: uint16(aux_sym_float_literal_token1),
	6950: uint16(aux_sym_float_literal_token2),
	6951: uint16(4),
	6952: uint16(655),
	6953: uint16(1),
	6954: uint16(sym_int_literal),
	6955: uint16(657),
	6956: uint16(1),
	6957: uint16(anon_sym_RPAREN),
	6958: uint16(3),
	6959: uint16(2),
	6960: uint16(sym_block_comment),
	6961: uint16(sym_line_comment),
	6962: uint16(651),
	6963: uint16(3),
	6964: uint16(sym_identifier),
	6965: uint16(aux_sym_float_literal_token1),
	6966: uint16(aux_sym_float_literal_token2),
	6967: uint16(5),
	6968: uint16(21),
	6969: uint16(1),
	6970: uint16(anon_sym_AT),
	6971: uint16(574),
	6972: uint16(1),
	6973: uint16(sym_identifier),
	6974: uint16(261),
	6975: uint16(1),
	6976: uint16(sym_variable_identifier_declaration),
	6977: uint16(3),
	6978: uint16(2),
	6979: uint16(sym_block_comment),
	6980: uint16(sym_line_comment),
	6981: uint16(47),
	6982: uint16(2),
	6983: uint16(sym_attribute),
	6984: uint16(aux_sym_global_variable_declaration_repeat1),
	6985: uint16(2),
	6986: uint16(3),
	6987: uint16(2),
	6988: uint16(sym_block_comment),
	6989: uint16(sym_line_comment),
	6990: uint16(659),
	6991: uint16(5),
	6992: uint16(anon_sym_SEMI),
	6993: uint16(anon_sym_EQ),
	6994: uint16(anon_sym_COMMA),
	6995: uint16(anon_sym_RPAREN),
	6996: uint16(anon_sym_RBRACE),
	6997: uint16(4),
	6998: uint16(199),
	6999: uint16(1),
	7000: uint16(aux_sym_lhs_expression_repeat1),
	7001: uint16(3),
	7002: uint16(2),
	7003: uint16(sym_block_comment),
	7004: uint16(sym_line_comment),
	7005: uint16(661),
	7006: uint16(2),
	7007: uint16(anon_sym_LPAREN),
	7008: uint16(sym_identifier),
	7009: uint16(663),
	7010: uint16(2),
	7011: uint16(anon_sym_AMP),
	7012: uint16(anon_sym_STAR),
	7013: uint16(5),
	7014: uint16(666),
	7015: uint16(1),
	7016: uint16(anon_sym_RBRACE),
	7017: uint16(668),
	7018: uint16(1),
	7019: uint16(anon_sym_case),
	7020: uint16(670),
	7021: uint16(1),
	7022: uint16(anon_sym_default),
	7023: uint16(3),
	7024: uint16(2),
	7025: uint16(sym_block_comment),
	7026: uint16(sym_line_comment),
	7027: uint16(193),
	7028: uint16(2),
	7029: uint16(sym_switch_body),
	7030: uint16(aux_sym_switch_statement_repeat1),
	7031: uint16(5),
	7032: uint16(199),
	7033: uint16(1),
	7034: uint16(anon_sym_LBRACE),
	7035: uint16(672),
	7036: uint16(1),
	7037: uint16(anon_sym_if),
	7038: uint16(121),
	7039: uint16(1),
	7040: uint16(sym_else_statement),
	7041: uint16(3),
	7042: uint16(2),
	7043: uint16(sym_block_comment),
	7044: uint16(sym_line_comment),
	7045: uint16(129),
	7046: uint16(2),
	7047: uint16(sym_compound_statement),
	7048: uint16(sym_if_statement),
	7049: uint16(5),
	7050: uint16(674),
	7051: uint16(1),
	7052: uint16(sym_identifier),
	7053: uint16(676),
	7054: uint16(1),
	7055: uint16(anon_sym_LPAREN),
	7056: uint16(199),
	7057: uint16(1),
	7058: uint16(aux_sym_lhs_expression_repeat1),
	7059: uint16(3),
	7060: uint16(2),
	7061: uint16(sym_block_comment),
	7062: uint16(sym_line_comment),
	7063: uint16(678),
	7064: uint16(2),
	7065: uint16(anon_sym_AMP),
	7066: uint16(anon_sym_STAR),
	7067: uint16(4),
	7068: uint16(680),
	7069: uint16(1),
	7070: uint16(anon_sym_read),
	7071: uint16(326),
	7072: uint16(1),
	7073: uint16(sym_access_mode),
	7074: uint16(3),
	7075: uint16(2),
	7076: uint16(sym_block_comment),
	7077: uint16(sym_line_comment),
	7078: uint16(682),
	7079: uint16(2),
	7080: uint16(anon_sym_write),
	7081: uint16(anon_sym_read_write),
	7082: uint16(4),
	7083: uint16(684),
	7084: uint16(1),
	7085: uint16(anon_sym_COMMA),
	7086: uint16(212),
	7087: uint16(1),
	7088: uint16(aux_sym_case_selectors_repeat1),
	7089: uint16(3),
	7090: uint16(2),
	7091: uint16(sym_block_comment),
	7092: uint16(sym_line_comment),
	7093: uint16(686),
	7094: uint16(2),
	7095: uint16(anon_sym_LBRACE),
	7096: uint16(anon_sym_COLON),
	7097: uint16(4),
	7098: uint16(668),
	7099: uint16(1),
	7100: uint16(anon_sym_case),
	7101: uint16(670),
	7102: uint16(1),
	7103: uint16(anon_sym_default),
	7104: uint16(3),
	7105: uint16(2),
	7106: uint16(sym_block_comment),
	7107: uint16(sym_line_comment),
	7108: uint16(200),
	7109: uint16(2),
	7110: uint16(sym_switch_body),
	7111: uint16(aux_sym_switch_statement_repeat1),
	7112: uint16(5),
	7113: uint16(688),
	7114: uint16(1),
	7115: uint16(sym_identifier),
	7116: uint16(690),
	7117: uint16(1),
	7118: uint16(anon_sym_LT),
	7119: uint16(253),
	7120: uint16(1),
	7121: uint16(sym_variable_qualifier),
	7122: uint16(254),
	7123: uint16(1),
	7124: uint16(sym_variable_identifier_declaration),
	7125: uint16(3),
	7126: uint16(2),
	7127: uint16(sym_block_comment),
	7128: uint16(sym_line_comment),
	7129: uint16(5),
	7130: uint16(199),
	7131: uint16(1),
	7132: uint16(anon_sym_LBRACE),
	7133: uint16(692),
	7134: uint16(1),
	7135: uint16(anon_sym_DASH_GT),
	7136: uint16(171),
	7137: uint16(1),
	7138: uint16(sym_compound_statement),
	7139: uint16(277),
	7140: uint16(1),
	7141: uint16(sym_function_return_type_declaration),
	7142: uint16(3),
	7143: uint16(2),
	7144: uint16(sym_block_comment),
	7145: uint16(sym_line_comment),
	7146: uint16(4),
	7147: uint16(694),
	7148: uint16(1),
	7149: uint16(anon_sym_COMMA),
	7150: uint16(208),
	7151: uint16(1),
	7152: uint16(aux_sym_case_selectors_repeat1),
	7153: uint16(3),
	7154: uint16(2),
	7155: uint16(sym_block_comment),
	7156: uint16(sym_line_comment),
	7157: uint16(697),
	7158: uint16(2),
	7159: uint16(anon_sym_LBRACE),
	7160: uint16(anon_sym_COLON),
	7161: uint16(5),
	7162: uint16(199),
	7163: uint16(1),
	7164: uint16(anon_sym_LBRACE),
	7165: uint16(692),
	7166: uint16(1),
	7167: uint16(anon_sym_DASH_GT),
	7168: uint16(163),
	7169: uint16(1),
	7170: uint16(sym_compound_statement),
	7171: uint16(247),
	7172: uint16(1),
	7173: uint16(sym_function_return_type_declaration),
	7174: uint16(3),
	7175: uint16(2),
	7176: uint16(sym_block_comment),
	7177: uint16(sym_line_comment),
	7178: uint16(2),
	7179: uint16(3),
	7180: uint16(2),
	7181: uint16(sym_block_comment),
	7182: uint16(sym_line_comment),
	7183: uint16(699),
	7184: uint16(4),
	7185: uint16(anon_sym_u32),
	7186: uint16(anon_sym_i32),
	7187: uint16(anon_sym_f32),
	7188: uint16(anon_sym_f16),
	7189: uint16(4),
	7190: uint16(680),
	7191: uint16(1),
	7192: uint16(anon_sym_read),
	7193: uint16(303),
	7194: uint16(1),
	7195: uint16(sym_access_mode),
	7196: uint16(3),
	7197: uint16(2),
	7198: uint16(sym_block_comment),
	7199: uint16(sym_line_comment),
	7200: uint16(682),
	7201: uint16(2),
	7202: uint16(anon_sym_write),
	7203: uint16(anon_sym_read_write),
	7204: uint16(4),
	7205: uint16(701),
	7206: uint16(1),
	7207: uint16(anon_sym_COMMA),
	7208: uint16(208),
	7209: uint16(1),
	7210: uint16(aux_sym_case_selectors_repeat1),
	7211: uint16(3),
	7212: uint16(2),
	7213: uint16(sym_block_comment),
	7214: uint16(sym_line_comment),
	7215: uint16(562),
	7216: uint16(2),
	7217: uint16(anon_sym_LBRACE),
	7218: uint16(anon_sym_COLON),
	7219: uint16(3),
	7220: uint16(655),
	7221: uint16(1),
	7222: uint16(sym_int_literal),
	7223: uint16(3),
	7224: uint16(2),
	7225: uint16(sym_block_comment),
	7226: uint16(sym_line_comment),
	7227: uint16(651),
	7228: uint16(3),
	7229: uint16(sym_identifier),
	7230: uint16(aux_sym_float_literal_token1),
	7231: uint16(aux_sym_float_literal_token2),
	7232: uint16(5),
	7233: uint16(199),
	7234: uint16(1),
	7235: uint16(anon_sym_LBRACE),
	7236: uint16(692),
	7237: uint16(1),
	7238: uint16(anon_sym_DASH_GT),
	7239: uint16(166),
	7240: uint16(1),
	7241: uint16(sym_compound_statement),
	7242: uint16(264),
	7243: uint16(1),
	7244: uint16(sym_function_return_type_declaration),
	7245: uint16(3),
	7246: uint16(2),
	7247: uint16(sym_block_comment),
	7248: uint16(sym_line_comment),
	7249: uint16(5),
	7250: uint16(199),
	7251: uint16(1),
	7252: uint16(anon_sym_LBRACE),
	7253: uint16(692),
	7254: uint16(1),
	7255: uint16(anon_sym_DASH_GT),
	7256: uint16(165),
	7257: uint16(1),
	7258: uint16(sym_compound_statement),
	7259: uint16(259),
	7260: uint16(1),
	7261: uint16(sym_function_return_type_declaration),
	7262: uint16(3),
	7263: uint16(2),
	7264: uint16(sym_block_comment),
	7265: uint16(sym_line_comment),
	7266: uint16(4),
	7267: uint16(680),
	7268: uint16(1),
	7269: uint16(anon_sym_read),
	7270: uint16(319),
	7271: uint16(1),
	7272: uint16(sym_access_mode),
	7273: uint16(3),
	7274: uint16(2),
	7275: uint16(sym_block_comment),
	7276: uint16(sym_line_comment),
	7277: uint16(682),
	7278: uint16(2),
	7279: uint16(anon_sym_write),
	7280: uint16(anon_sym_read_write),
	7281: uint16(4),
	7282: uint16(703),
	7283: uint16(1),
	7284: uint16(anon_sym_SEMI),
	7285: uint16(705),
	7286: uint16(1),
	7287: uint16(anon_sym_EQ),
	7288: uint16(707),
	7289: uint16(1),
	7290: uint16(anon_sym_COLON),
	7291: uint16(3),
	7292: uint16(2),
	7293: uint16(sym_block_comment),
	7294: uint16(sym_line_comment),
	7295: uint16(4),
	7296: uint16(709),
	7297: uint16(1),
	7298: uint16(anon_sym_LPAREN),
	7299: uint16(711),
	7300: uint16(1),
	7301: uint16(anon_sym_LT),
	7302: uint16(82),
	7303: uint16(1),
	7304: uint16(sym_argument_list_expression),
	7305: uint16(3),
	7306: uint16(2),
	7307: uint16(sym_block_comment),
	7308: uint16(sym_line_comment),
	7309: uint16(3),
	7310: uint16(707),
	7311: uint16(1),
	7312: uint16(anon_sym_COLON),
	7313: uint16(3),
	7314: uint16(2),
	7315: uint16(sym_block_comment),
	7316: uint16(sym_line_comment),
	7317: uint16(713),
	7318: uint16(2),
	7319: uint16(anon_sym_SEMI),
	7320: uint16(anon_sym_EQ),
	7321: uint16(2),
	7322: uint16(3),
	7323: uint16(2),
	7324: uint16(sym_block_comment),
	7325: uint16(sym_line_comment),
	7326: uint16(715),
	7327: uint16(3),
	7328: uint16(anon_sym_RBRACE),
	7329: uint16(anon_sym_case),
	7330: uint16(anon_sym_default),
	7331: uint16(4),
	7332: uint16(717),
	7333: uint16(1),
	7334: uint16(anon_sym_LBRACE),
	7335: uint16(719),
	7336: uint16(1),
	7337: uint16(anon_sym_COLON),
	7338: uint16(229),
	7339: uint16(1),
	7340: uint16(sym_case_compound_statement),
	7341: uint16(3),
	7342: uint16(2),
	7343: uint16(sym_block_comment),
	7344: uint16(sym_line_comment),
	7345: uint16(2),
	7346: uint16(3),
	7347: uint16(2),
	7348: uint16(sym_block_comment),
	7349: uint16(sym_line_comment),
	7350: uint16(721),
	7351: uint16(3),
	7352: uint16(anon_sym_SEMI),
	7353: uint16(anon_sym_COMMA),
	7354: uint16(anon_sym_RPAREN),
	7355: uint16(3),
	7356: uint16(707),
	7357: uint16(1),
	7358: uint16(anon_sym_COLON),
	7359: uint16(3),
	7360: uint16(2),
	7361: uint16(sym_block_comment),
	7362: uint16(sym_line_comment),
	7363: uint16(723),
	7364: uint16(2),
	7365: uint16(anon_sym_SEMI),
	7366: uint16(anon_sym_EQ),
	7367: uint16(2),
	7368: uint16(3),
	7369: uint16(2),
	7370: uint16(sym_block_comment),
	7371: uint16(sym_line_comment),
	7372: uint16(725),
	7373: uint16(3),
	7374: uint16(anon_sym_SEMI),
	7375: uint16(anon_sym_COMMA),
	7376: uint16(anon_sym_RPAREN),
	7377: uint16(2),
	7378: uint16(3),
	7379: uint16(2),
	7380: uint16(sym_block_comment),
	7381: uint16(sym_line_comment),
	7382: uint16(727),
	7383: uint16(3),
	7384: uint16(anon_sym_RBRACE),
	7385: uint16(anon_sym_case),
	7386: uint16(anon_sym_default),
	7387: uint16(4),
	7388: uint16(717),
	7389: uint16(1),
	7390: uint16(anon_sym_LBRACE),
	7391: uint16(729),
	7392: uint16(1),
	7393: uint16(anon_sym_COLON),
	7394: uint16(220),
	7395: uint16(1),
	7396: uint16(sym_case_compound_statement),
	7397: uint16(3),
	7398: uint16(2),
	7399: uint16(sym_block_comment),
	7400: uint16(sym_line_comment),
	7401: uint16(4),
	7402: uint16(707),
	7403: uint16(1),
	7404: uint16(anon_sym_COLON),
	7405: uint16(731),
	7406: uint16(1),
	7407: uint16(anon_sym_SEMI),
	7408: uint16(733),
	7409: uint16(1),
	7410: uint16(anon_sym_EQ),
	7411: uint16(3),
	7412: uint16(2),
	7413: uint16(sym_block_comment),
	7414: uint16(sym_line_comment),
	7415: uint16(3),
	7416: uint16(737),
	7417: uint16(1),
	7418: uint16(anon_sym_RBRACE),
	7419: uint16(3),
	7420: uint16(2),
	7421: uint16(sym_block_comment),
	7422: uint16(sym_line_comment),
	7423: uint16(735),
	7424: uint16(2),
	7425: uint16(anon_sym_AT),
	7426: uint16(sym_identifier),
	7427: uint16(2),
	7428: uint16(3),
	7429: uint16(2),
	7430: uint16(sym_block_comment),
	7431: uint16(sym_line_comment),
	7432: uint16(739),
	7433: uint16(3),
	7434: uint16(anon_sym_RBRACE),
	7435: uint16(anon_sym_case),
	7436: uint16(anon_sym_default),
	7437: uint16(2),
	7438: uint16(3),
	7439: uint16(2),
	7440: uint16(sym_block_comment),
	7441: uint16(sym_line_comment),
	7442: uint16(741),
	7443: uint16(3),
	7444: uint16(anon_sym_SEMI),
	7445: uint16(anon_sym_COMMA),
	7446: uint16(anon_sym_RPAREN),
	7447: uint16(3),
	7448: uint16(745),
	7449: uint16(1),
	7450: uint16(anon_sym_RPAREN),
	7451: uint16(3),
	7452: uint16(2),
	7453: uint16(sym_block_comment),
	7454: uint16(sym_line_comment),
	7455: uint16(743),
	7456: uint16(2),
	7457: uint16(anon_sym_AT),
	7458: uint16(sym_identifier),
	7459: uint16(2),
	7460: uint16(3),
	7461: uint16(2),
	7462: uint16(sym_block_comment),
	7463: uint16(sym_line_comment),
	7464: uint16(747),
	7465: uint16(3),
	7466: uint16(anon_sym_RBRACE),
	7467: uint16(anon_sym_case),
	7468: uint16(anon_sym_default),
	7469: uint16(3),
	7470: uint16(749),
	7471: uint16(1),
	7472: uint16(anon_sym_RPAREN),
	7473: uint16(3),
	7474: uint16(2),
	7475: uint16(sym_block_comment),
	7476: uint16(sym_line_comment),
	7477: uint16(743),
	7478: uint16(2),
	7479: uint16(anon_sym_AT),
	7480: uint16(sym_identifier),
	7481: uint16(2),
	7482: uint16(3),
	7483: uint16(2),
	7484: uint16(sym_block_comment),
	7485: uint16(sym_line_comment),
	7486: uint16(751),
	7487: uint16(3),
	7488: uint16(anon_sym_SEMI),
	7489: uint16(anon_sym_COMMA),
	7490: uint16(anon_sym_RPAREN),
	7491: uint16(2),
	7492: uint16(3),
	7493: uint16(2),
	7494: uint16(sym_block_comment),
	7495: uint16(sym_line_comment),
	7496: uint16(753),
	7497: uint16(3),
	7498: uint16(anon_sym_RBRACE),
	7499: uint16(anon_sym_case),
	7500: uint16(anon_sym_default),
	7501: uint16(2),
	7502: uint16(3),
	7503: uint16(2),
	7504: uint16(sym_block_comment),
	7505: uint16(sym_line_comment),
	7506: uint16(697),
	7507: uint16(3),
	7508: uint16(anon_sym_COMMA),
	7509: uint16(anon_sym_LBRACE),
	7510: uint16(anon_sym_COLON),
	7511: uint16(3),
	7512: uint16(755),
	7513: uint16(1),
	7514: uint16(anon_sym_RBRACE),
	7515: uint16(3),
	7516: uint16(2),
	7517: uint16(sym_block_comment),
	7518: uint16(sym_line_comment),
	7519: uint16(735),
	7520: uint16(2),
	7521: uint16(anon_sym_AT),
	7522: uint16(sym_identifier),
	7523: uint16(2),
	7524: uint16(3),
	7525: uint16(2),
	7526: uint16(sym_block_comment),
	7527: uint16(sym_line_comment),
	7528: uint16(757),
	7529: uint16(3),
	7530: uint16(anon_sym_SEMI),
	7531: uint16(anon_sym_COMMA),
	7532: uint16(anon_sym_RPAREN),
	7533: uint16(2),
	7534: uint16(3),
	7535: uint16(2),
	7536: uint16(sym_block_comment),
	7537: uint16(sym_line_comment),
	7538: uint16(759),
	7539: uint16(3),
	7540: uint16(anon_sym_RBRACE),
	7541: uint16(anon_sym_case),
	7542: uint16(anon_sym_default),
	7543: uint16(3),
	7544: uint16(761),
	7545: uint16(1),
	7546: uint16(anon_sym_COMMA),
	7547: uint16(763),
	7548: uint16(1),
	7549: uint16(anon_sym_GT),
	7550: uint16(3),
	7551: uint16(2),
	7552: uint16(sym_block_comment),
	7553: uint16(sym_line_comment),
	7554: uint16(3),
	7555: uint16(731),
	7556: uint16(1),
	7557: uint16(anon_sym_SEMI),
	7558: uint16(733),
	7559: uint16(1),
	7560: uint16(anon_sym_EQ),
	7561: uint16(3),
	7562: uint16(2),
	7563: uint16(sym_block_comment),
	7564: uint16(sym_line_comment),
	7565: uint16(3),
	7566: uint16(765),
	7567: uint16(1),
	7568: uint16(sym_identifier),
	7569: uint16(249),
	7570: uint16(1),
	7571: uint16(sym_variable_identifier_declaration),
	7572: uint16(3),
	7573: uint16(2),
	7574: uint16(sym_block_comment),
	7575: uint16(sym_line_comment),
	7576: uint16(3),
	7577: uint16(767),
	7578: uint16(1),
	7579: uint16(anon_sym_SEMI),
	7580: uint16(769),
	7581: uint16(1),
	7582: uint16(anon_sym_EQ),
	7583: uint16(3),
	7584: uint16(2),
	7585: uint16(sym_block_comment),
	7586: uint16(sym_line_comment),
	7587: uint16(3),
	7588: uint16(707),
	7589: uint16(1),
	7590: uint16(anon_sym_COLON),
	7591: uint16(771),
	7592: uint16(1),
	7593: uint16(anon_sym_EQ),
	7594: uint16(3),
	7595: uint16(2),
	7596: uint16(sym_block_comment),
	7597: uint16(sym_line_comment),
	7598: uint16(3),
	7599: uint16(773),
	7600: uint16(1),
	7601: uint16(anon_sym_COMMA),
	7602: uint16(775),
	7603: uint16(1),
	7604: uint16(anon_sym_RPAREN),
	7605: uint16(3),
	7606: uint16(2),
	7607: uint16(sym_block_comment),
	7608: uint16(sym_line_comment),
	7609: uint16(3),
	7610: uint16(777),
	7611: uint16(1),
	7612: uint16(anon_sym_SEMI),
	7613: uint16(779),
	7614: uint16(1),
	7615: uint16(anon_sym_EQ),
	7616: uint16(3),
	7617: uint16(2),
	7618: uint16(sym_block_comment),
	7619: uint16(sym_line_comment),
	7620: uint16(3),
	7621: uint16(199),
	7622: uint16(1),
	7623: uint16(anon_sym_LBRACE),
	7624: uint16(164),
	7625: uint16(1),
	7626: uint16(sym_compound_statement),
	7627: uint16(3),
	7628: uint16(2),
	7629: uint16(sym_block_comment),
	7630: uint16(sym_line_comment),
	7631: uint16(3),
	7632: uint16(717),
	7633: uint16(1),
	7634: uint16(anon_sym_LBRACE),
	7635: uint16(235),
	7636: uint16(1),
	7637: uint16(sym_case_compound_statement),
	7638: uint16(3),
	7639: uint16(2),
	7640: uint16(sym_block_comment),
	7641: uint16(sym_line_comment),
	7642: uint16(3),
	7643: uint16(703),
	7644: uint16(1),
	7645: uint16(anon_sym_SEMI),
	7646: uint16(705),
	7647: uint16(1),
	7648: uint16(anon_sym_EQ),
	7649: uint16(3),
	7650: uint16(2),
	7651: uint16(sym_block_comment),
	7652: uint16(sym_line_comment),
	7653: uint16(2),
	7654: uint16(3),
	7655: uint16(2),
	7656: uint16(sym_block_comment),
	7657: uint16(sym_line_comment),
	7658: uint16(743),
	7659: uint16(2),
	7660: uint16(anon_sym_AT),
	7661: uint16(sym_identifier),
	7662: uint16(3),
	7663: uint16(781),
	7664: uint16(1),
	7665: uint16(sym_identifier),
	7666: uint16(289),
	7667: uint16(1),
	7668: uint16(sym_variable_identifier_declaration),
	7669: uint16(3),
	7670: uint16(2),
	7671: uint16(sym_block_comment),
	7672: uint16(sym_line_comment),
	7673: uint16(2),
	7674: uint16(3),
	7675: uint16(2),
	7676: uint16(sym_block_comment),
	7677: uint16(sym_line_comment),
	7678: uint16(735),
	7679: uint16(2),
	7680: uint16(anon_sym_AT),
	7681: uint16(sym_identifier),
	7682: uint16(3),
	7683: uint16(783),
	7684: uint16(1),
	7685: uint16(sym_identifier),
	7686: uint16(278),
	7687: uint16(1),
	7688: uint16(sym_variable_identifier_declaration),
	7689: uint16(3),
	7690: uint16(2),
	7691: uint16(sym_block_comment),
	7692: uint16(sym_line_comment),
	7693: uint16(2),
	7694: uint16(3),
	7695: uint16(2),
	7696: uint16(sym_block_comment),
	7697: uint16(sym_line_comment),
	7698: uint16(713),
	7699: uint16(2),
	7700: uint16(anon_sym_SEMI),
	7701: uint16(anon_sym_EQ),
	7702: uint16(3),
	7703: uint16(785),
	7704: uint16(1),
	7705: uint16(anon_sym_COMMA),
	7706: uint16(787),
	7707: uint16(1),
	7708: uint16(anon_sym_GT),
	7709: uint16(3),
	7710: uint16(2),
	7711: uint16(sym_block_comment),
	7712: uint16(sym_line_comment),
	7713: uint16(3),
	7714: uint16(789),
	7715: uint16(1),
	7716: uint16(anon_sym_SEMI),
	7717: uint16(791),
	7718: uint16(1),
	7719: uint16(anon_sym_if),
	7720: uint16(3),
	7721: uint16(2),
	7722: uint16(sym_block_comment),
	7723: uint16(sym_line_comment),
	7724: uint16(3),
	7725: uint16(793),
	7726: uint16(1),
	7727: uint16(sym_identifier),
	7728: uint16(241),
	7729: uint16(1),
	7730: uint16(sym_variable_identifier_declaration),
	7731: uint16(3),
	7732: uint16(2),
	7733: uint16(sym_block_comment),
	7734: uint16(sym_line_comment),
	7735: uint16(3),
	7736: uint16(795),
	7737: uint16(1),
	7738: uint16(sym_identifier),
	7739: uint16(334),
	7740: uint16(1),
	7741: uint16(sym_variable_identifier_declaration),
	7742: uint16(3),
	7743: uint16(2),
	7744: uint16(sym_block_comment),
	7745: uint16(sym_line_comment),
	7746: uint16(3),
	7747: uint16(199),
	7748: uint16(1),
	7749: uint16(anon_sym_LBRACE),
	7750: uint16(170),
	7751: uint16(1),
	7752: uint16(sym_compound_statement),
	7753: uint16(3),
	7754: uint16(2),
	7755: uint16(sym_block_comment),
	7756: uint16(sym_line_comment),
	7757: uint16(3),
	7758: uint16(755),
	7759: uint16(1),
	7760: uint16(anon_sym_RBRACE),
	7761: uint16(797),
	7762: uint16(1),
	7763: uint16(anon_sym_COMMA),
	7764: uint16(3),
	7765: uint16(2),
	7766: uint16(sym_block_comment),
	7767: uint16(sym_line_comment),
	7768: uint16(2),
	7769: uint16(3),
	7770: uint16(2),
	7771: uint16(sym_block_comment),
	7772: uint16(sym_line_comment),
	7773: uint16(799),
	7774: uint16(2),
	7775: uint16(anon_sym_COMMA),
	7776: uint16(anon_sym_RBRACE),
	7777: uint16(3),
	7778: uint16(749),
	7779: uint16(1),
	7780: uint16(anon_sym_RPAREN),
	7781: uint16(801),
	7782: uint16(1),
	7783: uint16(anon_sym_COMMA),
	7784: uint16(3),
	7785: uint16(2),
	7786: uint16(sym_block_comment),
	7787: uint16(sym_line_comment),
	7788: uint16(2),
	7789: uint16(3),
	7790: uint16(2),
	7791: uint16(sym_block_comment),
	7792: uint16(sym_line_comment),
	7793: uint16(803),
	7794: uint16(2),
	7795: uint16(anon_sym_COMMA),
	7796: uint16(anon_sym_RPAREN),
	7797: uint16(3),
	7798: uint16(199),
	7799: uint16(1),
	7800: uint16(anon_sym_LBRACE),
	7801: uint16(168),
	7802: uint16(1),
	7803: uint16(sym_compound_statement),
	7804: uint16(3),
	7805: uint16(2),
	7806: uint16(sym_block_comment),
	7807: uint16(sym_line_comment),
	7808: uint16(3),
	7809: uint16(717),
	7810: uint16(1),
	7811: uint16(anon_sym_LBRACE),
	7812: uint16(229),
	7813: uint16(1),
	7814: uint16(sym_case_compound_statement),
	7815: uint16(3),
	7816: uint16(2),
	7817: uint16(sym_block_comment),
	7818: uint16(sym_line_comment),
	7819: uint16(3),
	7820: uint16(805),
	7821: uint16(1),
	7822: uint16(anon_sym_SEMI),
	7823: uint16(807),
	7824: uint16(1),
	7825: uint16(anon_sym_EQ),
	7826: uint16(3),
	7827: uint16(2),
	7828: uint16(sym_block_comment),
	7829: uint16(sym_line_comment),
	7830: uint16(2),
	7831: uint16(3),
	7832: uint16(2),
	7833: uint16(sym_block_comment),
	7834: uint16(sym_line_comment),
	7835: uint16(809),
	7836: uint16(2),
	7837: uint16(sym_identifier),
	7838: uint16(sym_int_literal),
	7839: uint16(3),
	7840: uint16(811),
	7841: uint16(1),
	7842: uint16(anon_sym_LBRACE),
	7843: uint16(296),
	7844: uint16(1),
	7845: uint16(sym_continuing_compound_statement),
	7846: uint16(3),
	7847: uint16(2),
	7848: uint16(sym_block_comment),
	7849: uint16(sym_line_comment),
	7850: uint16(3),
	7851: uint16(168),
	7852: uint16(1),
	7853: uint16(anon_sym_RPAREN),
	7854: uint16(813),
	7855: uint16(1),
	7856: uint16(anon_sym_COMMA),
	7857: uint16(3),
	7858: uint16(2),
	7859: uint16(sym_block_comment),
	7860: uint16(sym_line_comment),
	7861: uint16(3),
	7862: uint16(815),
	7863: uint16(1),
	7864: uint16(anon_sym_COMMA),
	7865: uint16(817),
	7866: uint16(1),
	7867: uint16(anon_sym_RPAREN),
	7868: uint16(3),
	7869: uint16(2),
	7870: uint16(sym_block_comment),
	7871: uint16(sym_line_comment),
	7872: uint16(2),
	7873: uint16(3),
	7874: uint16(2),
	7875: uint16(sym_block_comment),
	7876: uint16(sym_line_comment),
	7877: uint16(819),
	7878: uint16(2),
	7879: uint16(anon_sym_COMMA),
	7880: uint16(anon_sym_RBRACE),
	7881: uint16(3),
	7882: uint16(821),
	7883: uint16(1),
	7884: uint16(anon_sym_COMMA),
	7885: uint16(823),
	7886: uint16(1),
	7887: uint16(anon_sym_RBRACE),
	7888: uint16(3),
	7889: uint16(2),
	7890: uint16(sym_block_comment),
	7891: uint16(sym_line_comment),
	7892: uint16(2),
	7893: uint16(3),
	7894: uint16(2),
	7895: uint16(sym_block_comment),
	7896: uint16(sym_line_comment),
	7897: uint16(825),
	7898: uint16(2),
	7899: uint16(anon_sym_COMMA),
	7900: uint16(anon_sym_GT),
	7901: uint16(3),
	7902: uint16(827),
	7903: uint16(1),
	7904: uint16(anon_sym_COMMA),
	7905: uint16(829),
	7906: uint16(1),
	7907: uint16(anon_sym_GT),
	7908: uint16(3),
	7909: uint16(2),
	7910: uint16(sym_block_comment),
	7911: uint16(sym_line_comment),
	7912: uint16(2),
	7913: uint16(3),
	7914: uint16(2),
	7915: uint16(sym_block_comment),
	7916: uint16(sym_line_comment),
	7917: uint16(831),
	7918: uint16(2),
	7919: uint16(anon_sym_COMMA),
	7920: uint16(anon_sym_RPAREN),
	7921: uint16(3),
	7922: uint16(833),
	7923: uint16(1),
	7924: uint16(anon_sym_COMMA),
	7925: uint16(835),
	7926: uint16(1),
	7927: uint16(anon_sym_RPAREN),
	7928: uint16(3),
	7929: uint16(2),
	7930: uint16(sym_block_comment),
	7931: uint16(sym_line_comment),
	7932: uint16(3),
	7933: uint16(199),
	7934: uint16(1),
	7935: uint16(anon_sym_LBRACE),
	7936: uint16(174),
	7937: uint16(1),
	7938: uint16(sym_compound_statement),
	7939: uint16(3),
	7940: uint16(2),
	7941: uint16(sym_block_comment),
	7942: uint16(sym_line_comment),
	7943: uint16(2),
	7944: uint16(3),
	7945: uint16(2),
	7946: uint16(sym_block_comment),
	7947: uint16(sym_line_comment),
	7948: uint16(723),
	7949: uint16(2),
	7950: uint16(anon_sym_SEMI),
	7951: uint16(anon_sym_EQ),
	7952: uint16(3),
	7953: uint16(199),
	7954: uint16(1),
	7955: uint16(anon_sym_LBRACE),
	7956: uint16(127),
	7957: uint16(1),
	7958: uint16(sym_compound_statement),
	7959: uint16(3),
	7960: uint16(2),
	7961: uint16(sym_block_comment),
	7962: uint16(sym_line_comment),
	7963: uint16(3),
	7964: uint16(709),
	7965: uint16(1),
	7966: uint16(anon_sym_LPAREN),
	7967: uint16(82),
	7968: uint16(1),
	7969: uint16(sym_argument_list_expression),
	7970: uint16(3),
	7971: uint16(2),
	7972: uint16(sym_block_comment),
	7973: uint16(sym_line_comment),
	7974: uint16(3),
	7975: uint16(707),
	7976: uint16(1),
	7977: uint16(anon_sym_COLON),
	7978: uint16(837),
	7979: uint16(1),
	7980: uint16(anon_sym_EQ),
	7981: uint16(3),
	7982: uint16(2),
	7983: uint16(sym_block_comment),
	7984: uint16(sym_line_comment),
	7985: uint16(3),
	7986: uint16(657),
	7987: uint16(1),
	7988: uint16(anon_sym_RPAREN),
	7989: uint16(839),
	7990: uint16(1),
	7991: uint16(anon_sym_COMMA),
	7992: uint16(3),
	7993: uint16(2),
	7994: uint16(sym_block_comment),
	7995: uint16(sym_line_comment),
	7996: uint16(3),
	7997: uint16(27),
	7998: uint16(1),
	7999: uint16(anon_sym_LPAREN),
	8000: uint16(89),
	8001: uint16(1),
	8002: uint16(sym_parenthesized_expression),
	8003: uint16(3),
	8004: uint16(2),
	8005: uint16(sym_block_comment),
	8006: uint16(sym_line_comment),
	8007: uint16(2),
	8008: uint16(841),
	8009: uint16(1),
	8010: uint16(anon_sym_LT),
	8011: uint16(3),
	8012: uint16(2),
	8013: uint16(sym_block_comment),
	8014: uint16(sym_line_comment),
	8015: uint16(2),
	8016: uint16(843),
	8017: uint16(1),
	8018: uint16(anon_sym_COMMA),
	8019: uint16(3),
	8020: uint16(2),
	8021: uint16(sym_block_comment),
	8022: uint16(sym_line_comment),
	8023: uint16(2),
	8024: uint16(845),
	8025: uint16(1),
	8026: uint16(sym_identifier),
	8027: uint16(3),
	8028: uint16(2),
	8029: uint16(sym_block_comment),
	8030: uint16(sym_line_comment),
	8031: uint16(2),
	8032: uint16(847),
	8033: uint16(1),
	8034: uint16(anon_sym_RPAREN),
	8035: uint16(3),
	8036: uint16(2),
	8037: uint16(sym_block_comment),
	8038: uint16(sym_line_comment),
	8039: uint16(2),
	8040: uint16(849),
	8041: uint16(1),
	8042: uint16(anon_sym_SEMI),
	8043: uint16(3),
	8044: uint16(2),
	8045: uint16(sym_block_comment),
	8046: uint16(sym_line_comment),
	8047: uint16(2),
	8048: uint16(837),
	8049: uint16(1),
	8050: uint16(anon_sym_EQ),
	8051: uint16(3),
	8052: uint16(2),
	8053: uint16(sym_block_comment),
	8054: uint16(sym_line_comment),
	8055: uint16(2),
	8056: uint16(851),
	8057: uint16(1),
	8058: uint16(anon_sym_LT),
	8059: uint16(3),
	8060: uint16(2),
	8061: uint16(sym_block_comment),
	8062: uint16(sym_line_comment),
	8063: uint16(2),
	8064: uint16(415),
	8065: uint16(1),
	8066: uint16(anon_sym_SEMI),
	8067: uint16(3),
	8068: uint16(2),
	8069: uint16(sym_block_comment),
	8070: uint16(sym_line_comment),
	8071: uint16(2),
	8072: uint16(853),
	8073: uint16(1),
	8074: uint16(anon_sym_LPAREN),
	8075: uint16(3),
	8076: uint16(2),
	8077: uint16(sym_block_comment),
	8078: uint16(sym_line_comment),
	8079: uint16(2),
	8080: uint16(711),
	8081: uint16(1),
	8082: uint16(anon_sym_LT),
	8083: uint16(3),
	8084: uint16(2),
	8085: uint16(sym_block_comment),
	8086: uint16(sym_line_comment),
	8087: uint16(2),
	8088: uint16(855),
	8089: uint16(1),
	8090: uint16(anon_sym_LBRACE),
	8091: uint16(3),
	8092: uint16(2),
	8093: uint16(sym_block_comment),
	8094: uint16(sym_line_comment),
	8095: uint16(2),
	8096: uint16(857),
	8097: uint16(1),
	8098: uint16(anon_sym_LPAREN),
	8099: uint16(3),
	8100: uint16(2),
	8101: uint16(sym_block_comment),
	8102: uint16(sym_line_comment),
	8103: uint16(2),
	8104: uint16(859),
	8105: uint16(1),
	8106: uint16(anon_sym_RBRACE),
	8107: uint16(3),
	8108: uint16(2),
	8109: uint16(sym_block_comment),
	8110: uint16(sym_line_comment),
	8111: uint16(2),
	8112: uint16(861),
	8113: uint16(1),
	8114: uint16(anon_sym_SEMI),
	8115: uint16(3),
	8116: uint16(2),
	8117: uint16(sym_block_comment),
	8118: uint16(sym_line_comment),
	8119: uint16(2),
	8120: uint16(863),
	8121: uint16(1),
	8122: uint16(anon_sym_RBRACE),
	8123: uint16(3),
	8124: uint16(2),
	8125: uint16(sym_block_comment),
	8126: uint16(sym_line_comment),
	8127: uint16(2),
	8128: uint16(865),
	8129: uint16(1),
	8130: uint16(anon_sym_LT),
	8131: uint16(3),
	8132: uint16(2),
	8133: uint16(sym_block_comment),
	8134: uint16(sym_line_comment),
	8135: uint16(2),
	8136: uint16(867),
	8137: uint16(1),
	8138: uint16(anon_sym_LT),
	8139: uint16(3),
	8140: uint16(2),
	8141: uint16(sym_block_comment),
	8142: uint16(sym_line_comment),
	8143: uint16(2),
	8144: uint16(869),
	8145: uint16(1),
	8146: uint16(anon_sym_LBRACE),
	8147: uint16(3),
	8148: uint16(2),
	8149: uint16(sym_block_comment),
	8150: uint16(sym_line_comment),
	8151: uint16(2),
	8152: uint16(871),
	8153: uint16(1),
	8154: uint16(anon_sym_LT),
	8155: uint16(3),
	8156: uint16(2),
	8157: uint16(sym_block_comment),
	8158: uint16(sym_line_comment),
	8159: uint16(2),
	8160: uint16(873),
	8161: uint16(1),
	8162: uint16(anon_sym_GT),
	8163: uint16(3),
	8164: uint16(2),
	8165: uint16(sym_block_comment),
	8166: uint16(sym_line_comment),
	8167: uint16(2),
	8168: uint16(707),
	8169: uint16(1),
	8170: uint16(anon_sym_COLON),
	8171: uint16(3),
	8172: uint16(2),
	8173: uint16(sym_block_comment),
	8174: uint16(sym_line_comment),
	8175: uint16(2),
	8176: uint16(875),
	8177: uint16(1),
	8178: uint16(anon_sym_RPAREN),
	8179: uint16(3),
	8180: uint16(2),
	8181: uint16(sym_block_comment),
	8182: uint16(sym_line_comment),
	8183: uint16(2),
	8184: uint16(877),
	8185: uint16(1),
	8186: uint16(anon_sym_RPAREN),
	8187: uint16(3),
	8188: uint16(2),
	8189: uint16(sym_block_comment),
	8190: uint16(sym_line_comment),
	8191: uint16(2),
	8192: uint16(879),
	8193: uint16(1),
	8194: uint16(anon_sym_RPAREN),
	8195: uint16(3),
	8196: uint16(2),
	8197: uint16(sym_block_comment),
	8198: uint16(sym_line_comment),
	8199: uint16(2),
	8200: uint16(881),
	8201: uint16(1),
	8202: uint16(anon_sym_SEMI),
	8203: uint16(3),
	8204: uint16(2),
	8205: uint16(sym_block_comment),
	8206: uint16(sym_line_comment),
	8207: uint16(2),
	8208: uint16(223),
	8209: uint16(1),
	8210: uint16(anon_sym_RBRACE),
	8211: uint16(3),
	8212: uint16(2),
	8213: uint16(sym_block_comment),
	8214: uint16(sym_line_comment),
	8215: uint16(2),
	8216: uint16(883),
	8217: uint16(1),
	8218: uint16(sym_identifier),
	8219: uint16(3),
	8220: uint16(2),
	8221: uint16(sym_block_comment),
	8222: uint16(sym_line_comment),
	8223: uint16(2),
	8224: uint16(885),
	8225: uint16(1),
	8226: uint16(anon_sym_SEMI),
	8227: uint16(3),
	8228: uint16(2),
	8229: uint16(sym_block_comment),
	8230: uint16(sym_line_comment),
	8231: uint16(2),
	8232: uint16(887),
	8233: uint16(1),
	8234: uint16(sym_identifier),
	8235: uint16(3),
	8236: uint16(2),
	8237: uint16(sym_block_comment),
	8238: uint16(sym_line_comment),
	8239: uint16(2),
	8240: uint16(889),
	8241: uint16(1),
	8242: uint16(sym_identifier),
	8243: uint16(3),
	8244: uint16(2),
	8245: uint16(sym_block_comment),
	8246: uint16(sym_line_comment),
	8247: uint16(2),
	8248: uint16(891),
	8249: uint16(1),
	8250: uint16(sym_identifier),
	8251: uint16(3),
	8252: uint16(2),
	8253: uint16(sym_block_comment),
	8254: uint16(sym_line_comment),
	8255: uint16(2),
	8256: uint16(893),
	8257: uint16(1),
	8258: uint16(anon_sym_RBRACE),
	8259: uint16(3),
	8260: uint16(2),
	8261: uint16(sym_block_comment),
	8262: uint16(sym_line_comment),
	8263: uint16(2),
	8264: uint16(895),
	8265: uint16(1),
	8266: uint16(anon_sym_COMMA),
	8267: uint16(3),
	8268: uint16(2),
	8269: uint16(sym_block_comment),
	8270: uint16(sym_line_comment),
	8271: uint16(2),
	8272: uint16(283),
	8273: uint16(1),
	8274: uint16(anon_sym_RBRACE),
	8275: uint16(3),
	8276: uint16(2),
	8277: uint16(sym_block_comment),
	8278: uint16(sym_line_comment),
	8279: uint16(2),
	8280: uint16(897),
	8281: uint16(1),
	8282: uint16(anon_sym_GT),
	8283: uint16(3),
	8284: uint16(2),
	8285: uint16(sym_block_comment),
	8286: uint16(sym_line_comment),
	8287: uint16(2),
	8288: uint16(787),
	8289: uint16(1),
	8290: uint16(anon_sym_GT),
	8291: uint16(3),
	8292: uint16(2),
	8293: uint16(sym_block_comment),
	8294: uint16(sym_line_comment),
	8295: uint16(2),
	8296: uint16(117),
	8297: uint16(1),
	8298: uint16(anon_sym_RPAREN),
	8299: uint16(3),
	8300: uint16(2),
	8301: uint16(sym_block_comment),
	8302: uint16(sym_line_comment),
	8303: uint16(2),
	8304: uint16(549),
	8305: uint16(1),
	8306: uint16(anon_sym_EQ),
	8307: uint16(3),
	8308: uint16(2),
	8309: uint16(sym_block_comment),
	8310: uint16(sym_line_comment),
	8311: uint16(2),
	8312: uint16(899),
	8313: uint16(1),
	8314: uint16(anon_sym_SEMI),
	8315: uint16(3),
	8316: uint16(2),
	8317: uint16(sym_block_comment),
	8318: uint16(sym_line_comment),
	8319: uint16(2),
	8320: uint16(901),
	8321: uint16(1),
	8322: uint16(anon_sym_LBRACE),
	8323: uint16(3),
	8324: uint16(2),
	8325: uint16(sym_block_comment),
	8326: uint16(sym_line_comment),
	8327: uint16(2),
	8328: uint16(903),
	8329: uint16(1),
	8330: uint16(anon_sym_LPAREN),
	8331: uint16(3),
	8332: uint16(2),
	8333: uint16(sym_block_comment),
	8334: uint16(sym_line_comment),
	8335: uint16(2),
	8336: uint16(905),
	8337: uint16(1),
	8338: uint16(anon_sym_EQ),
	8339: uint16(3),
	8340: uint16(2),
	8341: uint16(sym_block_comment),
	8342: uint16(sym_line_comment),
	8343: uint16(2),
	8344: uint16(907),
	8345: uint16(1),
	8346: uint16(anon_sym_GT),
	8347: uint16(3),
	8348: uint16(2),
	8349: uint16(sym_block_comment),
	8350: uint16(sym_line_comment),
	8351: uint16(2),
	8352: uint16(909),
	8353: uint16(1),
	8354: uint16(anon_sym_COMMA),
	8355: uint16(3),
	8356: uint16(2),
	8357: uint16(sym_block_comment),
	8358: uint16(sym_line_comment),
	8359: uint16(2),
	8360: uint16(911),
	8361: uint16(1),
	8362: uint16(anon_sym_COMMA),
	8363: uint16(3),
	8364: uint16(2),
	8365: uint16(sym_block_comment),
	8366: uint16(sym_line_comment),
	8367: uint16(2),
	8368: uint16(913),
	8369: uint16(1),
	8370: uint16(anon_sym_COMMA),
	8371: uint16(3),
	8372: uint16(2),
	8373: uint16(sym_block_comment),
	8374: uint16(sym_line_comment),
	8375: uint16(2),
	8376: uint16(763),
	8377: uint16(1),
	8378: uint16(anon_sym_GT),
	8379: uint16(3),
	8380: uint16(2),
	8381: uint16(sym_block_comment),
	8382: uint16(sym_line_comment),
	8383: uint16(2),
	8384: uint16(915),
	8385: uint16(1),
	8386: uint16(anon_sym_COMMA),
	8387: uint16(3),
	8388: uint16(2),
	8389: uint16(sym_block_comment),
	8390: uint16(sym_line_comment),
	8391: uint16(2),
	8392: uint16(917),
	8393: uint16(1),
	8394: uint16(anon_sym_SEMI),
	8395: uint16(3),
	8396: uint16(2),
	8397: uint16(sym_block_comment),
	8398: uint16(sym_line_comment),
	8399: uint16(2),
	8400: uint16(281),
	8401: uint16(1),
	8402: uint16(anon_sym_RBRACE),
	8403: uint16(3),
	8404: uint16(2),
	8405: uint16(sym_block_comment),
	8406: uint16(sym_line_comment),
	8407: uint16(2),
	8408: uint16(771),
	8409: uint16(1),
	8410: uint16(anon_sym_EQ),
	8411: uint16(3),
	8412: uint16(2),
	8413: uint16(sym_block_comment),
	8414: uint16(sym_line_comment),
	8415: uint16(2),
	8416: uint16(919),
	8417: uint16(1),
	8418: uint16(anon_sym_COMMA),
	8419: uint16(3),
	8420: uint16(2),
	8421: uint16(sym_block_comment),
	8422: uint16(sym_line_comment),
	8423: uint16(2),
	8424: uint16(921),
	8425: uint16(1),
	8426: uint16(anon_sym_RBRACE),
	8427: uint16(3),
	8428: uint16(2),
	8429: uint16(sym_block_comment),
	8430: uint16(sym_line_comment),
	8431: uint16(2),
	8432: uint16(923),
	8433: uint16(1),
	8434: uint16(anon_sym_RBRACE),
	8435: uint16(3),
	8436: uint16(2),
	8437: uint16(sym_block_comment),
	8438: uint16(sym_line_comment),
	8439: uint16(2),
	8440: uint16(123),
	8441: uint16(1),
	8442: uint16(anon_sym_RPAREN),
	8443: uint16(3),
	8444: uint16(2),
	8445: uint16(sym_block_comment),
	8446: uint16(sym_line_comment),
	8447: uint16(2),
	8448: uint16(925),
	8449: uint16(1),
	8450: uint16(anon_sym_RPAREN),
	8451: uint16(3),
	8452: uint16(2),
	8453: uint16(sym_block_comment),
	8454: uint16(sym_line_comment),
	8455: uint16(2),
	8456: uint16(927),
	8457: uint16(1),
	8458: uint16(anon_sym_SEMI),
	8459: uint16(3),
	8460: uint16(2),
	8461: uint16(sym_block_comment),
	8462: uint16(sym_line_comment),
	8463: uint16(2),
	8464: uint16(929),
	8465: uint16(1),
	8466: uint16(anon_sym_SEMI),
	8467: uint16(3),
	8468: uint16(2),
	8469: uint16(sym_block_comment),
	8470: uint16(sym_line_comment),
	8471: uint16(2),
	8472: uint16(931),
	8473: uint16(1),
	8474: uint16(anon_sym_SEMI),
	8475: uint16(3),
	8476: uint16(2),
	8477: uint16(sym_block_comment),
	8478: uint16(sym_line_comment),
	8479: uint16(2),
	8480: uint16(933),
	8481: uint16(1),
	8482: uint16(anon_sym_SEMI),
	8483: uint16(3),
	8484: uint16(2),
	8485: uint16(sym_block_comment),
	8486: uint16(sym_line_comment),
	8487: uint16(2),
	8488: uint16(935),
	8489: uint16(1),
	8490: uint16(anon_sym_RBRACE),
	8491: uint16(3),
	8492: uint16(2),
	8493: uint16(sym_block_comment),
	8494: uint16(sym_line_comment),
	8495: uint16(2),
	8496: uint16(789),
	8497: uint16(1),
	8498: uint16(anon_sym_SEMI),
	8499: uint16(3),
	8500: uint16(2),
	8501: uint16(sym_block_comment),
	8502: uint16(sym_line_comment),
	8503: uint16(2),
	8504: uint16(937),
	8505: uint16(1),
	8506: uint16(anon_sym_RBRACE),
	8507: uint16(3),
	8508: uint16(2),
	8509: uint16(sym_block_comment),
	8510: uint16(sym_line_comment),
	8511: uint16(2),
	8512: uint16(939),
	8513: uint16(1),
	8515: uint16(3),
	8516: uint16(2),
	8517: uint16(sym_block_comment),
	8518: uint16(sym_line_comment),
	8519: uint16(2),
	8520: uint16(941),
	8521: uint16(1),
	8522: uint16(anon_sym_RBRACE),
	8523: uint16(3),
	8524: uint16(2),
	8525: uint16(sym_block_comment),
	8526: uint16(sym_line_comment),
	8527: uint16(2),
	8528: uint16(943),
	8529: uint16(1),
	8530: uint16(anon_sym_RPAREN),
	8531: uint16(3),
	8532: uint16(2),
	8533: uint16(sym_block_comment),
	8534: uint16(sym_line_comment),
	8535: uint16(2),
	8536: uint16(945),
	8537: uint16(1),
	8538: uint16(anon_sym_LPAREN),
	8539: uint16(3),
	8540: uint16(2),
	8541: uint16(sym_block_comment),
	8542: uint16(sym_line_comment),
	8543: uint16(2),
	8544: uint16(947),
	8545: uint16(1),
	8546: uint16(anon_sym_LBRACE),
	8547: uint16(3),
	8548: uint16(2),
	8549: uint16(sym_block_comment),
	8550: uint16(sym_line_comment),
	8551: uint16(2),
	8552: uint16(949),
	8553: uint16(1),
	8554: uint16(anon_sym_RBRACE),
	8555: uint16(3),
	8556: uint16(2),
	8557: uint16(sym_block_comment),
	8558: uint16(sym_line_comment),
	8559: uint16(2),
	8560: uint16(951),
	8561: uint16(1),
	8562: uint16(sym_identifier),
	8563: uint16(3),
	8564: uint16(2),
	8565: uint16(sym_block_comment),
	8566: uint16(sym_line_comment),
	8567: uint16(2),
	8568: uint16(953),
	8569: uint16(1),
	8570: uint16(sym_identifier),
	8571: uint16(3),
	8572: uint16(2),
	8573: uint16(sym_block_comment),
	8574: uint16(sym_line_comment),
	8575: uint16(2),
	8576: uint16(955),
	8577: uint16(1),
	8578: uint16(sym_identifier),
	8579: uint16(3),
	8580: uint16(2),
	8581: uint16(sym_block_comment),
	8582: uint16(sym_line_comment),
	8583: uint16(2),
	8584: uint16(957),
	8585: uint16(1),
	8586: uint16(sym_identifier),
	8587: uint16(3),
	8588: uint16(2),
	8589: uint16(sym_block_comment),
	8590: uint16(sym_line_comment),
	8591: uint16(2),
	8592: uint16(959),
	8593: uint16(1),
	8594: uint16(sym_identifier),
	8595: uint16(3),
	8596: uint16(2),
	8597: uint16(sym_block_comment),
	8598: uint16(sym_line_comment),
	8599: uint16(2),
	8600: uint16(961),
	8601: uint16(1),
	8602: uint16(sym_identifier),
	8603: uint16(3),
	8604: uint16(2),
	8605: uint16(sym_block_comment),
	8606: uint16(sym_line_comment),
	8607: uint16(2),
	8608: uint16(963),
	8609: uint16(1),
	8610: uint16(anon_sym_GT),
	8611: uint16(3),
	8612: uint16(2),
	8613: uint16(sym_block_comment),
	8614: uint16(sym_line_comment),
}

var ts_small_parse_table_map = [328]uint32_t{
	1:   uint32(102),
	2:   uint32(191),
	3:   uint32(280),
	4:   uint32(369),
	5:   uint32(458),
	6:   uint32(520),
	7:   uint32(606),
	8:   uint32(692),
	9:   uint32(754),
	10:  uint32(837),
	11:  uint32(896),
	12:  uint32(979),
	13:  uint32(1062),
	14:  uint32(1121),
	15:  uint32(1176),
	16:  uint32(1231),
	17:  uint32(1286),
	18:  uint32(1355),
	19:  uint32(1424),
	20:  uint32(1476),
	21:  uint32(1530),
	22:  uint32(1581),
	23:  uint32(1632),
	24:  uint32(1683),
	25:  uint32(1745),
	26:  uint32(1807),
	27:  uint32(1869),
	28:  uint32(1931),
	29:  uint32(1993),
	30:  uint32(2055),
	31:  uint32(2147),
	32:  uint32(2239),
	33:  uint32(2329),
	34:  uint32(2421),
	35:  uint32(2513),
	36:  uint32(2602),
	37:  uint32(2691),
	38:  uint32(2777),
	39:  uint32(2863),
	40:  uint32(2900),
	41:  uint32(2937),
	42:  uint32(2972),
	43:  uint32(3009),
	44:  uint32(3044),
	45:  uint32(3079),
	46:  uint32(3113),
	47:  uint32(3167),
	48:  uint32(3209),
	49:  uint32(3243),
	50:  uint32(3287),
	51:  uint32(3321),
	52:  uint32(3355),
	53:  uint32(3405),
	54:  uint32(3439),
	55:  uint32(3473),
	56:  uint32(3519),
	57:  uint32(3557),
	58:  uint32(3591),
	59:  uint32(3625),
	60:  uint32(3683),
	61:  uint32(3739),
	62:  uint32(3791),
	63:  uint32(3847),
	64:  uint32(3881),
	65:  uint32(3915),
	66:  uint32(3972),
	67:  uint32(4029),
	68:  uint32(4088),
	69:  uint32(4147),
	70:  uint32(4178),
	71:  uint32(4237),
	72:  uint32(4268),
	73:  uint32(4327),
	74:  uint32(4383),
	75:  uint32(4439),
	76:  uint32(4495),
	77:  uint32(4551),
	78:  uint32(4607),
	79:  uint32(4663),
	80:  uint32(4695),
	81:  uint32(4751),
	82:  uint32(4807),
	83:  uint32(4863),
	84:  uint32(4919),
	85:  uint32(4975),
	86:  uint32(5031),
	87:  uint32(5087),
	88:  uint32(5143),
	89:  uint32(5172),
	90:  uint32(5201),
	91:  uint32(5230),
	92:  uint32(5259),
	93:  uint32(5288),
	94:  uint32(5317),
	95:  uint32(5346),
	96:  uint32(5375),
	97:  uint32(5404),
	98:  uint32(5433),
	99:  uint32(5462),
	100: uint32(5491),
	101: uint32(5541),
	102: uint32(5591),
	103: uint32(5641),
	104: uint32(5667),
	105: uint32(5691),
	106: uint32(5719),
	107: uint32(5747),
	108: uint32(5775),
	109: uint32(5805),
	110: uint32(5833),
	111: uint32(5857),
	112: uint32(5885),
	113: uint32(5909),
	114: uint32(5937),
	115: uint32(5965),
	116: uint32(5993),
	117: uint32(6017),
	118: uint32(6038),
	119: uint32(6059),
	120: uint32(6078),
	121: uint32(6097),
	122: uint32(6116),
	123: uint32(6143),
	124: uint32(6166),
	125: uint32(6185),
	126: uint32(6212),
	127: uint32(6239),
	128: uint32(6256),
	129: uint32(6272),
	130: uint32(6288),
	131: uint32(6318),
	132: uint32(6334),
	133: uint32(6350),
	134: uint32(6366),
	135: uint32(6382),
	136: uint32(6398),
	137: uint32(6414),
	138: uint32(6430),
	139: uint32(6446),
	140: uint32(6462),
	141: uint32(6492),
	142: uint32(6518),
	143: uint32(6534),
	144: uint32(6557),
	145: uint32(6572),
	146: uint32(6587),
	147: uint32(6602),
	148: uint32(6617),
	149: uint32(6641),
	150: uint32(6665),
	151: uint32(6689),
	152: uint32(6711),
	153: uint32(6733),
	154: uint32(6755),
	155: uint32(6779),
	156: uint32(6803),
	157: uint32(6827),
	158: uint32(6842),
	159: uint32(6857),
	160: uint32(6878),
	161: uint32(6899),
	162: uint32(6917),
	163: uint32(6935),
	164: uint32(6951),
	165: uint32(6967),
	166: uint32(6985),
	167: uint32(6997),
	168: uint32(7013),
	169: uint32(7031),
	170: uint32(7049),
	171: uint32(7067),
	172: uint32(7082),
	173: uint32(7097),
	174: uint32(7112),
	175: uint32(7129),
	176: uint32(7146),
	177: uint32(7161),
	178: uint32(7178),
	179: uint32(7189),
	180: uint32(7204),
	181: uint32(7219),
	182: uint32(7232),
	183: uint32(7249),
	184: uint32(7266),
	185: uint32(7281),
	186: uint32(7295),
	187: uint32(7309),
	188: uint32(7321),
	189: uint32(7331),
	190: uint32(7345),
	191: uint32(7355),
	192: uint32(7367),
	193: uint32(7377),
	194: uint32(7387),
	195: uint32(7401),
	196: uint32(7415),
	197: uint32(7427),
	198: uint32(7437),
	199: uint32(7447),
	200: uint32(7459),
	201: uint32(7469),
	202: uint32(7481),
	203: uint32(7491),
	204: uint32(7501),
	205: uint32(7511),
	206: uint32(7523),
	207: uint32(7533),
	208: uint32(7543),
	209: uint32(7554),
	210: uint32(7565),
	211: uint32(7576),
	212: uint32(7587),
	213: uint32(7598),
	214: uint32(7609),
	215: uint32(7620),
	216: uint32(7631),
	217: uint32(7642),
	218: uint32(7653),
	219: uint32(7662),
	220: uint32(7673),
	221: uint32(7682),
	222: uint32(7693),
	223: uint32(7702),
	224: uint32(7713),
	225: uint32(7724),
	226: uint32(7735),
	227: uint32(7746),
	228: uint32(7757),
	229: uint32(7768),
	230: uint32(7777),
	231: uint32(7788),
	232: uint32(7797),
	233: uint32(7808),
	234: uint32(7819),
	235: uint32(7830),
	236: uint32(7839),
	237: uint32(7850),
	238: uint32(7861),
	239: uint32(7872),
	240: uint32(7881),
	241: uint32(7892),
	242: uint32(7901),
	243: uint32(7912),
	244: uint32(7921),
	245: uint32(7932),
	246: uint32(7943),
	247: uint32(7952),
	248: uint32(7963),
	249: uint32(7974),
	250: uint32(7985),
	251: uint32(7996),
	252: uint32(8007),
	253: uint32(8015),
	254: uint32(8023),
	255: uint32(8031),
	256: uint32(8039),
	257: uint32(8047),
	258: uint32(8055),
	259: uint32(8063),
	260: uint32(8071),
	261: uint32(8079),
	262: uint32(8087),
	263: uint32(8095),
	264: uint32(8103),
	265: uint32(8111),
	266: uint32(8119),
	267: uint32(8127),
	268: uint32(8135),
	269: uint32(8143),
	270: uint32(8151),
	271: uint32(8159),
	272: uint32(8167),
	273: uint32(8175),
	274: uint32(8183),
	275: uint32(8191),
	276: uint32(8199),
	277: uint32(8207),
	278: uint32(8215),
	279: uint32(8223),
	280: uint32(8231),
	281: uint32(8239),
	282: uint32(8247),
	283: uint32(8255),
	284: uint32(8263),
	285: uint32(8271),
	286: uint32(8279),
	287: uint32(8287),
	288: uint32(8295),
	289: uint32(8303),
	290: uint32(8311),
	291: uint32(8319),
	292: uint32(8327),
	293: uint32(8335),
	294: uint32(8343),
	295: uint32(8351),
	296: uint32(8359),
	297: uint32(8367),
	298: uint32(8375),
	299: uint32(8383),
	300: uint32(8391),
	301: uint32(8399),
	302: uint32(8407),
	303: uint32(8415),
	304: uint32(8423),
	305: uint32(8431),
	306: uint32(8439),
	307: uint32(8447),
	308: uint32(8455),
	309: uint32(8463),
	310: uint32(8471),
	311: uint32(8479),
	312: uint32(8487),
	313: uint32(8495),
	314: uint32(8503),
	315: uint32(8511),
	316: uint32(8519),
	317: uint32(8527),
	318: uint32(8535),
	319: uint32(8543),
	320: uint32(8551),
	321: uint32(8559),
	322: uint32(8567),
	323: uint32(8575),
	324: uint32(8583),
	325: uint32(8591),
	326: uint32(8599),
	327: uint32(8607),
}

var ts_parse_actions = [965]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_source_file),
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
		Fstate: uint16(133),
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
		Fstate: uint16(258),
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
		Fstate: uint16(242),
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
		Fstate: uint16(358),
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
		Fstate: uint16(356),
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
		Fstate: uint16(355),
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
		Fstate: uint16(354),
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
		Fstate: uint16(353),
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
		Fstate: uint16(206),
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
		Fstate: uint16(74),
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
		Fstate: uint16(21),
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
		Fstate: uint16(80),
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
		Fstate: uint16(75),
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
		Fstate: uint16(76),
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
		Fstate: uint16(73),
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
		Fstate: uint16(176),
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
		Fstate: uint16(302),
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
		Fstate: uint16(300),
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
		Fstate: uint16(284),
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
		Fstate: uint16(299),
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
		Fcount: uint8(1),
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
		Fstate: uint16(218),
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
		Fstate: uint16(290),
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
		Fstate: uint16(31),
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
		Fcount: uint8(1),
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
		Fstate: uint16(31),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_return_statement),
	})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(74),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(21),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(75),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(76),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(73),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(176),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(302),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(300),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(284),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(299),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(218),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(290),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(31),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Fstate:      uint16(31),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	100: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(140),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(251),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(191),
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
		Fcount: uint8(1),
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
		Fstate: uint16(321),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(206),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_for_header),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(293),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_for_header),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_for_header),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_argument_list_expression_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(aux_sym_const_expresssion_repeat1),
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
		Fstate:      uint16(176),
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
		Fsymbol:      uint16(aux_sym_const_expresssion_repeat1),
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
		Fstate:      uint16(75),
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
		Fsymbol:      uint16(aux_sym_const_expresssion_repeat1),
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
		Fstate:      uint16(76),
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
		Fsymbol:      uint16(aux_sym_const_expresssion_repeat1),
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
		Fstate:      uint16(73),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_const_expresssion_repeat1),
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
		Fstate:      uint16(302),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_const_expresssion_repeat1),
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
		Fstate:      uint16(300),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_const_expresssion_repeat1),
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
		Fstate:      uint16(284),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_const_expresssion_repeat1),
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
		Fstate:      uint16(299),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_const_expresssion_repeat1),
	})))),
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
		Fstate:      uint16(293),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(96),
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
		Fcount: uint8(1),
	}})),
	163: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_compound_assignment_operator),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_compound_assignment_operator),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_const_expresssion_repeat1),
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
		Fstate: uint16(234),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_const_expresssion_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_global_variable_declaration_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_global_variable_declaration_repeat1),
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
		Fstate:      uint16(353),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_attribute),
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
		Fstate: uint16(184),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_attribute),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_attribute),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_attribute),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_attribute),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_attribute),
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
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_attribute),
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
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_attribute),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(20),
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
		Fcount: uint8(1),
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
		Fstate: uint16(332),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(351),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(350),
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
		Fcount: uint8(1),
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
		Fstate: uint16(13),
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
		Fcount: uint8(1),
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
		Fstate: uint16(345),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(343),
	}})))),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(3),
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
		Fcount: uint8(1),
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
		Fstate: uint16(341),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	224: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
	})))),
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
		Fstate:      uint16(137),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
	})))),
	232: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(251),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	234: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
	})))),
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
		Fstate:      uint16(191),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
	})))),
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
		Fstate:      uint16(69),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
	})))),
	241: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	242: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
	})))),
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
		Fstate:      uint16(321),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(25),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(20),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(351),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(350),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(13),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(345),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(343),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(3),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
	})))),
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
		Fstate:      uint16(341),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(206),
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
		Fsymbol:      uint16(aux_sym_compound_statement_repeat1),
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
		Fstate:      uint16(202),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	280: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	281: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(239),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	284: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	285: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(256),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	288: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	289: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	290: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	291: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	292: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	293: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	294: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_compound_statement),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_compound_statement),
	})))),
	297: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	298: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_compound_statement),
	})))),
	299: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	300: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_compound_statement),
	})))),
	301: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	302: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_bool_literal),
	})))),
	303: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	304: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_bool_literal),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	306: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_type_declaration),
	})))),
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
		Fcount: uint8(1),
	}})),
	310: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	311: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	312: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_const_literal),
	})))),
	313: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	314: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_const_literal),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_float_literal),
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
		Fcount: uint8(1),
	}})),
	318: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_float_literal),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_argument_list_expression),
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
		Fcount: uint8(1),
	}})),
	322: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_argument_list_expression),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	324: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(5),
	})))),
	325: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	326: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	327: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	328: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(5),
	})))),
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
		Fcount: uint8(1),
	}})),
	330: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(15),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	334: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	335: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(16),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	338: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	339: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	340: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	341: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(24),
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
		Fstate: uint16(26),
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
		Fstate: uint16(314),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_argument_list_expression),
	})))),
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
		Fcount: uint8(1),
	}})),
	350: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_argument_list_expression),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_type_constructor_or_function_call_expression),
	})))),
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
		Fcount: uint8(1),
	}})),
	354: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_type_constructor_or_function_call_expression),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_unary_expression),
		Fproduction_id: uint16(2),
	})))),
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
		Fcount: uint8(1),
	}})),
	358: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_unary_expression),
		Fproduction_id: uint16(2),
	})))),
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
		Fcount: uint8(1),
	}})),
	362: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_composite_value_decomposition_expression),
		Fproduction_id: uint16(6),
	})))),
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
		Fcount: uint8(1),
	}})),
	366: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_composite_value_decomposition_expression),
		Fproduction_id: uint16(6),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_bitcast_expression),
	})))),
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
		Fcount: uint8(1),
	}})),
	370: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_bitcast_expression),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_subscript_expression),
		Fproduction_id: uint16(10),
	})))),
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
		Fcount: uint8(1),
	}})),
	374: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_subscript_expression),
		Fproduction_id: uint16(10),
	})))),
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
		Fstate: uint16(9),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(23),
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
		Fstate: uint16(30),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_argument_list_expression),
	})))),
	383: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	384: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_argument_list_expression),
	})))),
	385: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	386: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_argument_list_expression),
	})))),
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
		Fcount: uint8(1),
	}})),
	388: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_argument_list_expression),
	})))),
	389: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	390: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	391: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(134),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	394: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_assignment_statement),
		Fproduction_id: uint16(5),
	})))),
	395: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	396: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	397: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	398: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	399: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	400: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_increment_statement),
	})))),
	401: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	402: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_increment_statement),
	})))),
	403: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	404: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_decrement_statement),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_decrement_statement),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_global_constant_declaration),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_global_constant_declaration),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	424: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(14),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(14),
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
		Fcount: uint8(1),
	}})),
	428: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(143),
	}})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_variable_statement),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_return_statement),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_variable_statement),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	444: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_break_statement),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_break_statement),
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
		Fcount: uint8(1),
	}})),
	448: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(17),
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
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(17),
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
		Fcount: uint8(1),
	}})),
	452: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_loop_statement),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_loop_statement),
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
		Fcount: uint8(1),
	}})),
	456: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_while_statement),
		Fproduction_id: uint16(15),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_while_statement),
		Fproduction_id: uint16(15),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_continue_statement),
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
		Fsymbol:      uint16(sym_continue_statement),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_switch_statement),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_switch_statement),
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
		Fcount: uint8(1),
	}})),
	468: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_loop_statement),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_loop_statement),
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
		Fcount: uint8(1),
	}})),
	472: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_for_statement),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_for_statement),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_discard_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_discard_statement),
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
		Fcount: uint8(1),
	}})),
	480: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_else_statement),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_else_statement),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_loop_statement),
	})))),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_loop_statement),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__statement),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__statement),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
	})))),
	495: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(132),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	497: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
	})))),
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
		Fstate:      uint16(258),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
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
		Fstate:      uint16(242),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
	})))),
	504: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(358),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
	})))),
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
		Fstate:      uint16(356),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	509: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
	})))),
	510: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
	})))),
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
		Fstate:      uint16(353),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	515: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_source_file_repeat2),
	})))),
	516: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(206),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fstate: uint16(132),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(331),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_postfix_expression),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_lhs_expression),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(28),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(17),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	534: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(357),
	}})))),
	535: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	536: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_lhs_expression),
	})))),
	537: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	538: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_lhs_expression),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_postfix_expression),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_postfix_expression),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_lhs_expression),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_postfix_expression),
	})))),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_lhs_expression),
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
		Fstate: uint16(10),
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
		Fstate: uint16(45),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	556: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	557: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	558: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	559: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	560: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      uint16(354),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_case_selectors),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_case_selectors),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_enable_directive),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	571: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_struct_declaration),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_struct_declaration),
		Fproduction_id: uint16(4),
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
		Fstate: uint16(304),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(8),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(11),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	583: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(13),
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
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(3),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	587: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__declaration),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(7),
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
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_struct_declaration),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(16),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(9),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(12),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_type_declaration),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	603: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_type_declaration),
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
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_type_declaration),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	607: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_struct_declaration_repeat1),
	})))),
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
		Fstate:      uint16(304),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	610: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_struct_declaration_repeat1),
	})))),
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
		Fstate:      uint16(353),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(257),
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
		Fstate: uint16(313),
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
		Fstate: uint16(282),
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
		Fstate: uint16(282),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(270),
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
		Fcount: uint8(1),
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
		Fstate: uint16(270),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_attribute_repeat1),
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
		Fstate:      uint16(285),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_attribute_repeat1),
	})))),
	629: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(285),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	630: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	631: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_attribute_repeat1),
	})))),
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
		Fstate:      uint16(76),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_parameter_list_repeat1),
	})))),
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
		Fstate:      uint16(304),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_parameter_list_repeat1),
	})))),
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
		Fstate:      uint16(353),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(273),
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
		Fstate: uint16(137),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_switch_statement_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	646: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_switch_statement_repeat1),
	})))),
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
		Fstate:      uint16(173),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	649: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_switch_statement_repeat1),
	})))),
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
		Fstate:      uint16(226),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_attribute_repeat1),
	})))),
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
		Fstate: uint16(55),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_attribute_repeat1),
	})))),
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
		Fstate: uint16(53),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_variable_identifier_declaration),
		Fproduction_id: uint16(1),
	})))),
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
		Fsymbol:      uint16(aux_sym_lhs_expression_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_lhs_expression_repeat1),
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
		Fstate:      uint16(199),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(226),
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
		Fstate: uint16(25),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	675: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	677: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	678: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	679: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	680: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(318),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(318),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(157),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_case_selectors),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	695: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_case_selectors_repeat1),
	})))),
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
		Fstate:      uint16(175),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_case_selectors_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(330),
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
		Fstate: uint16(158),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_global_constant_declaration),
	})))),
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
		Fstate: uint16(18),
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
		Fstate: uint16(56),
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
		Fstate: uint16(2),
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
		Fstate: uint16(57),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_variable_declaration),
	})))),
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
		Fsymbol:      uint16(sym_switch_body),
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
		Fstate: uint16(62),
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
		Fstate: uint16(248),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_const_expression),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_variable_declaration),
	})))),
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
		Fsymbol:      uint16(sym_const_expression),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_case_compound_statement),
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
		Fstate: uint16(265),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_global_constant_declaration),
	})))),
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
		Fstate: uint16(19),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_struct_declaration_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	740: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_switch_body),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_const_expression),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	744: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_parameter_list_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_parameter_list),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_case_compound_statement),
	})))),
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
		Fsymbol:      uint16(sym_parameter_list),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_const_expression),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_switch_body),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	756: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(161),
	}})))),
	757: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	758: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_const_expression),
	})))),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_case_compound_statement),
	})))),
	761: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(267),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(217),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	768: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_global_variable_declaration),
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
		Fstate: uint16(43),
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
		Fstate: uint16(41),
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
		Fstate: uint16(46),
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
		Fstate: uint16(222),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_variable_statement),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	780: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	781: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	782: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	783: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(223),
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
		Fstate: uint16(211),
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
		Fstate: uint16(178),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(8),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(244),
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
		Fstate: uint16(228),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_struct_member),
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
		Fstate: uint16(231),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_parameter),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_global_variable_declaration),
	})))),
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
		Fstate: uint16(44),
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
	810: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(68),
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
		Fstate: uint16(48),
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
		Fstate: uint16(196),
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
		Fstate: uint16(54),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_struct_member),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	823: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	824: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	825: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	826: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_address_space),
	})))),
	827: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: uint16(203),
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
		Fstate: uint16(310),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_parameter),
	})))),
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
		Fstate: uint16(233),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_parameter_list),
	})))),
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
		Fstate: uint16(29),
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
		Fstate: uint16(195),
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
		Fstate: uint16(210),
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
		Fstate: uint16(213),
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
		Fstate: uint16(138),
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
		Fstate: uint16(141),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_type_alias_declaration),
	})))),
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
		Fstate: uint16(58),
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
		Fstate: uint16(172),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_function_return_type_declaration),
	})))),
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
		Fstate: uint16(34),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_continuing_statement),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_global_variable_declaration),
	})))),
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
		Fstate: uint16(130),
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
		Fstate: uint16(135),
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
		Fstate: uint16(189),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_function_return_type_declaration),
	})))),
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
		Fstate: uint16(61),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(209),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	878: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(279),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(5),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	884: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_variable_qualifier),
	})))),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_global_variable_declaration),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_variable_qualifier),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(292),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_continuing_compound_statement),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	898: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_access_mode),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(181),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	906: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(252),
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
		Fstate: uint16(250),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_texel_format),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	918: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	922: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_continuing_compound_statement),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(348),
	}})))),
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
		Fstate: uint16(215),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(131),
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
		Fstate: uint16(128),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(124),
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
	936: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_fallthrough_statement),
	})))),
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
		Fstate: uint16(225),
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
	940: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_continuing_compound_statement),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_for_header),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(32),
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
		Fstate: uint16(65),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_break_if_statement),
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
		Fstate: uint16(52),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(323),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(136),
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
		Fstate: uint16(325),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(283),
	}})))),
}

func tree_sitter_wgsl(tls *libc.TLS) (r uintptr) {
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
	Fkeyword_capture_token:     uint16(sym_identifier),
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
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(ts_lex_keywords)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_wgsl_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_wgsl_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_wgsl_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_wgsl_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_wgsl_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00identifier\x00line_comment\x00;\x00=\x00let\x00override\x00(\x00,\x00)\x00type\x00fn\x00->\x00struct\x00{\x00}\x00enable\x00@\x00_\x00+=\x00-=\x00*=\x00/=\x00%=\x00&=\x00|=\x00^=\x00if\x00else\x00switch\x00case\x00:\x00default\x00fallthrough\x00loop\x00for\x00while\x00break\x00continue\x00continuing\x00return\x00discard\x00var\x00<\x00>\x00++\x00--\x00int_literal\x00float_literal_token1\x00float_literal_token2\x00true\x00false\x00bool\x00u32\x00i32\x00f32\x00f16\x00array\x00ptr\x00sampler\x00sampler_comparison\x00texture_depth_2d\x00texture_depth_2d_array\x00texture_depth_cube\x00texture_depth_cube_array\x00texture_depth_multisampled_2d\x00texture_1d\x00texture_2d\x00texture_2d_array\x00texture_3d\x00texture_cube\x00texture_cube_array\x00texture_multisampled_2d\x00texture_storage_1d\x00texture_storage_2d\x00texture_storage_2d_array\x00texture_storage_3d\x00vec2\x00vec3\x00vec4\x00mat2x2\x00mat2x3\x00mat2x4\x00mat3x2\x00mat3x3\x00mat3x4\x00mat4x2\x00mat4x3\x00mat4x4\x00rgba8unorm\x00rgba8snorm\x00rgba8uint\x00rgba8sint\x00rgba16uint\x00rgba16sint\x00rgba16float\x00r32uint\x00r32sint\x00r32float\x00rg32uint\x00rg32sint\x00rg32float\x00rgba32uint\x00rgba32sint\x00rgba32float\x00function\x00private\x00workgroup\x00uniform\x00storage\x00read\x00write\x00read_write\x00bitcast\x00||\x00&&\x00|\x00^\x00&\x00==\x00!=\x00<=\x00>=\x00<<\x00>>\x00+\x00-\x00*\x00/\x00%\x00!\x00~\x00[\x00]\x00.\x00block_comment\x00source_file\x00_declaration\x00global_variable_declaration\x00global_constant_declaration\x00type_alias_declaration\x00const_expression\x00function_declaration\x00function_return_type_declaration\x00struct_declaration\x00struct_member\x00enable_directive\x00attribute\x00_literal_or_identifier\x00parameter_list\x00parameter\x00_statement\x00compound_statement\x00assignment_statement\x00compound_assignment_operator\x00if_statement\x00else_statement\x00switch_statement\x00switch_body\x00case_selectors\x00case_compound_statement\x00fallthrough_statement\x00loop_statement\x00for_statement\x00for_header\x00while_statement\x00break_statement\x00break_if_statement\x00continue_statement\x00continuing_statement\x00continuing_compound_statement\x00return_statement\x00discard_statement\x00variable_statement\x00variable_declaration\x00variable_qualifier\x00variable_identifier_declaration\x00increment_statement\x00decrement_statement\x00_expression\x00const_literal\x00float_literal\x00bool_literal\x00parenthesized_expression\x00type_constructor_or_function_call_expression\x00type_declaration\x00_vec_prefix\x00_mat_prefix\x00texel_format\x00address_space\x00access_mode\x00argument_list_expression\x00bitcast_expression\x00binary_expression\x00unary_expression\x00postfix_expression\x00subscript_expression\x00lhs_expression\x00composite_value_decomposition_expression\x00source_file_repeat1\x00source_file_repeat2\x00global_variable_declaration_repeat1\x00const_expresssion_repeat1\x00struct_declaration_repeat1\x00attribute_repeat1\x00parameter_list_repeat1\x00compound_statement_repeat1\x00switch_statement_repeat1\x00case_selectors_repeat1\x00argument_list_expression_repeat1\x00lhs_expression_repeat1\x00accessor\x00alternative\x00argument\x00body\x00condition\x00consequence\x00left\x00name\x00parameters\x00right\x00subscript\x00value\x00"
