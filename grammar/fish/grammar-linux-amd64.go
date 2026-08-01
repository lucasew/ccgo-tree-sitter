// Code generated for linux/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-fish/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-fish -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && amd64

package grammar_fish

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
const LARGE_STATE_COUNT = 34
const MAX_ALIAS_SEQUENCE_LENGTH = 7
const PRODUCTION_ID_COUNT = 19
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const STATE_COUNT = 432
const SYMBOL_COUNT = 108
const TOKEN_COUNT = 57
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

const CONCAT = 0
const BRACKET_CONCAT = 1
const CONCAT_LIST = 2
const BEGIN_BRACE = 3

func tree_sitter_fish_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_fish_external_scanner_destroy(tls *libc.TLS, p uintptr) {
}

func tree_sitter_fish_external_scanner_reset(tls *libc.TLS, p uintptr) {
}

func tree_sitter_fish_external_scanner_serialize(tls *libc.TLS, p uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_fish_external_scanner_deserialize(tls *libc.TLS, p uintptr, b uintptr, n uint32) {
}

func tree_sitter_fish_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	// BEGIN_BRACE: { followed by whitespace or ; (for begin_statement)
	// Must take priority over internal '{' token used by brace_expansion
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BEGIN_BRACE))) != 0 {
		// Skip leading whitespace (since whitespace is in extras)
		for libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0)) // skip=true for whitespace
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('{') {
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0)) // consume '{'
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(';') || libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BEGIN_BRACE)
				return libc.BoolUint8(true1 != 0)
			}
		}
		// Not matched - return false to let internal lexer try
		// (returning false resets lexer state)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(CONCAT_LIST))) != 0 {
		if !((*TSLexer)(unsafe.Pointer(lexer)).Flookahead == 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('[')) {
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(CONCAT_LIST)
			return libc.BoolUint8(true1 != 0)
		}
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(CONCAT))) != 0 {
		if !((*TSLexer)(unsafe.Pointer(lexer)).Flookahead == 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('>') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('<') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(')') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(';') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('&') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('|') || libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0) {
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(CONCAT)
			return libc.BoolUint8(true1 != 0)
		}
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BRACKET_CONCAT))) != 0 {
		if !((*TSLexer)(unsafe.Pointer(lexer)).Flookahead == 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(')') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('(') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('}') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(',') || libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0) {
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BRACKET_CONCAT)
			return libc.BoolUint8(true1 != 0)
		}
	}
	return libc.BoolUint8(false1 != 0)
}

const sym_word = 1
const anon_sym_and = 2
const anon_sym_or = 3
const anon_sym_PIPE_PIPE = 4
const anon_sym_AMP_AMP = 5
const anon_sym_AMP_PIPE = 6
const anon_sym_2_GT_PIPE = 7
const anon_sym_PIPE = 8
const anon_sym_SEMI = 9
const anon_sym_AMP = 10
const anon_sym_LF = 11
const anon_sym_CR = 12
const anon_sym_CR_LF = 13
const anon_sym_BANG = 14
const anon_sym_not = 15
const anon_sym_DOLLAR = 16
const anon_sym_LPAREN = 17
const anon_sym_RPAREN = 18
const anon_sym_function = 19
const anon_sym_end = 20
const sym_integer = 21
const sym_float = 22
const anon_sym_return = 23
const anon_sym_switch = 24
const anon_sym_case = 25
const sym_break = 26
const sym_continue = 27
const anon_sym_for = 28
const anon_sym_in = 29
const anon_sym_while = 30
const anon_sym_if = 31
const anon_sym_else = 32
const anon_sym_begin = 33
const anon_sym_RBRACE = 34
const sym_comment = 35
const sym_variable_name = 36
const anon_sym_DOT_DOT = 37
const anon_sym_LBRACK = 38
const anon_sym_RBRACK = 39
const anon_sym_LBRACE = 40
const anon_sym_COMMA = 41
const anon_sym_DQUOTE = 42
const aux_sym_double_quote_string_token1 = 43
const anon_sym_SQUOTE = 44
const aux_sym_single_quote_string_token1 = 45
const sym_escape_sequence = 46
const sym_stream_redirect = 47
const sym_direction = 48
const anon_sym_POUND = 49
const sym_home_dir_expansion = 50
const sym_glob = 51
const sym_brace_word = 52
const sym__concat = 53
const sym__brace_concat = 54
const sym__concat_list = 55
const sym__begin_brace = 56
const sym_program = 57
const sym_conditional_execution = 58
const sym_pipe = 59
const sym_redirect_statement = 60
const sym__terminated_statement = 61
const sym__terminated_opt_statement = 62
const sym_negated_statement = 63
const sym__command_substitution_dollar = 64
const sym__command_substitution_inner = 65
const sym_command_substitution = 66
const sym_function_definition = 67
const sym_return = 68
const sym_switch_statement = 69
const sym_case_clause = 70
const sym_for_statement = 71
const sym_while_statement = 72
const sym_if_statement = 73
const sym_else_if_clause = 74
const sym_else_clause = 75
const sym_begin_statement = 76
const sym_variable_expansion = 77
const sym_index = 78
const sym_range = 79
const sym_list_element_access = 80
const sym_brace_expansion = 81
const sym_double_quote_string = 82
const sym_single_quote_string = 83
const sym_command = 84
const sym_file_redirect = 85
const sym__special_character = 86
const sym_concatenation = 87
const sym__expression = 88
const sym_brace_concatenation = 89
const sym__brace_expression = 90
const sym__base_brace_expression = 91
const aux_sym_program_repeat1 = 92
const aux_sym_function_definition_repeat1 = 93
const aux_sym_function_definition_repeat2 = 94
const aux_sym_switch_statement_repeat1 = 95
const aux_sym_case_clause_repeat1 = 96
const aux_sym_for_statement_repeat1 = 97
const aux_sym_while_statement_repeat1 = 98
const aux_sym_if_statement_repeat1 = 99
const aux_sym_variable_expansion_repeat1 = 100
const aux_sym_list_element_access_repeat1 = 101
const aux_sym_brace_expansion_repeat1 = 102
const aux_sym_double_quote_string_repeat1 = 103
const aux_sym_single_quote_string_repeat1 = 104
const aux_sym_command_repeat1 = 105
const aux_sym_concatenation_repeat1 = 106
const aux_sym_brace_concatenation_repeat1 = 107

var ts_symbol_names = [108]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 9,
	3:   __ccgo_ts + 13,
	4:   __ccgo_ts + 16,
	5:   __ccgo_ts + 19,
	6:   __ccgo_ts + 22,
	7:   __ccgo_ts + 25,
	8:   __ccgo_ts + 29,
	9:   __ccgo_ts + 31,
	10:  __ccgo_ts + 33,
	11:  __ccgo_ts + 35,
	12:  __ccgo_ts + 37,
	13:  __ccgo_ts + 39,
	14:  __ccgo_ts + 42,
	15:  __ccgo_ts + 44,
	16:  __ccgo_ts + 48,
	17:  __ccgo_ts + 50,
	18:  __ccgo_ts + 52,
	19:  __ccgo_ts + 54,
	20:  __ccgo_ts,
	21:  __ccgo_ts + 63,
	22:  __ccgo_ts + 71,
	23:  __ccgo_ts + 77,
	24:  __ccgo_ts + 84,
	25:  __ccgo_ts + 91,
	26:  __ccgo_ts + 96,
	27:  __ccgo_ts + 102,
	28:  __ccgo_ts + 111,
	29:  __ccgo_ts + 115,
	30:  __ccgo_ts + 118,
	31:  __ccgo_ts + 124,
	32:  __ccgo_ts + 127,
	33:  __ccgo_ts + 132,
	34:  __ccgo_ts + 138,
	35:  __ccgo_ts + 140,
	36:  __ccgo_ts + 148,
	37:  __ccgo_ts + 162,
	38:  __ccgo_ts + 165,
	39:  __ccgo_ts + 167,
	40:  __ccgo_ts + 169,
	41:  __ccgo_ts + 171,
	42:  __ccgo_ts + 173,
	43:  __ccgo_ts + 175,
	44:  __ccgo_ts + 202,
	45:  __ccgo_ts + 204,
	46:  __ccgo_ts + 231,
	47:  __ccgo_ts + 247,
	48:  __ccgo_ts + 263,
	49:  __ccgo_ts + 273,
	50:  __ccgo_ts + 275,
	51:  __ccgo_ts + 294,
	52:  __ccgo_ts + 4,
	53:  __ccgo_ts + 299,
	54:  __ccgo_ts + 307,
	55:  __ccgo_ts + 321,
	56:  __ccgo_ts + 169,
	57:  __ccgo_ts + 334,
	58:  __ccgo_ts + 342,
	59:  __ccgo_ts + 364,
	60:  __ccgo_ts + 369,
	61:  __ccgo_ts + 388,
	62:  __ccgo_ts + 410,
	63:  __ccgo_ts + 436,
	64:  __ccgo_ts + 454,
	65:  __ccgo_ts + 483,
	66:  __ccgo_ts + 511,
	67:  __ccgo_ts + 532,
	68:  __ccgo_ts + 77,
	69:  __ccgo_ts + 552,
	70:  __ccgo_ts + 569,
	71:  __ccgo_ts + 581,
	72:  __ccgo_ts + 595,
	73:  __ccgo_ts + 611,
	74:  __ccgo_ts + 624,
	75:  __ccgo_ts + 639,
	76:  __ccgo_ts + 651,
	77:  __ccgo_ts + 667,
	78:  __ccgo_ts + 686,
	79:  __ccgo_ts + 692,
	80:  __ccgo_ts + 698,
	81:  __ccgo_ts + 718,
	82:  __ccgo_ts + 734,
	83:  __ccgo_ts + 754,
	84:  __ccgo_ts + 774,
	85:  __ccgo_ts + 782,
	86:  __ccgo_ts + 796,
	87:  __ccgo_ts + 815,
	88:  __ccgo_ts + 829,
	89:  __ccgo_ts + 815,
	90:  __ccgo_ts + 841,
	91:  __ccgo_ts + 859,
	92:  __ccgo_ts + 882,
	93:  __ccgo_ts + 898,
	94:  __ccgo_ts + 926,
	95:  __ccgo_ts + 954,
	96:  __ccgo_ts + 979,
	97:  __ccgo_ts + 999,
	98:  __ccgo_ts + 1021,
	99:  __ccgo_ts + 1045,
	100: __ccgo_ts + 1066,
	101: __ccgo_ts + 1093,
	102: __ccgo_ts + 1121,
	103: __ccgo_ts + 1145,
	104: __ccgo_ts + 1173,
	105: __ccgo_ts + 1201,
	106: __ccgo_ts + 1217,
	107: __ccgo_ts + 1239,
}

var ts_symbol_map = [108]TSSymbol{
	1:   uint16(sym_word),
	2:   uint16(anon_sym_and),
	3:   uint16(anon_sym_or),
	4:   uint16(anon_sym_PIPE_PIPE),
	5:   uint16(anon_sym_AMP_AMP),
	6:   uint16(anon_sym_AMP_PIPE),
	7:   uint16(anon_sym_2_GT_PIPE),
	8:   uint16(anon_sym_PIPE),
	9:   uint16(anon_sym_SEMI),
	10:  uint16(anon_sym_AMP),
	11:  uint16(anon_sym_LF),
	12:  uint16(anon_sym_CR),
	13:  uint16(anon_sym_CR_LF),
	14:  uint16(anon_sym_BANG),
	15:  uint16(anon_sym_not),
	16:  uint16(anon_sym_DOLLAR),
	17:  uint16(anon_sym_LPAREN),
	18:  uint16(anon_sym_RPAREN),
	19:  uint16(anon_sym_function),
	20:  uint16(anon_sym_end),
	21:  uint16(sym_integer),
	22:  uint16(sym_float),
	23:  uint16(anon_sym_return),
	24:  uint16(anon_sym_switch),
	25:  uint16(anon_sym_case),
	26:  uint16(sym_break),
	27:  uint16(sym_continue),
	28:  uint16(anon_sym_for),
	29:  uint16(anon_sym_in),
	30:  uint16(anon_sym_while),
	31:  uint16(anon_sym_if),
	32:  uint16(anon_sym_else),
	33:  uint16(anon_sym_begin),
	34:  uint16(anon_sym_RBRACE),
	35:  uint16(sym_comment),
	36:  uint16(sym_variable_name),
	37:  uint16(anon_sym_DOT_DOT),
	38:  uint16(anon_sym_LBRACK),
	39:  uint16(anon_sym_RBRACK),
	40:  uint16(anon_sym_LBRACE),
	41:  uint16(anon_sym_COMMA),
	42:  uint16(anon_sym_DQUOTE),
	43:  uint16(aux_sym_double_quote_string_token1),
	44:  uint16(anon_sym_SQUOTE),
	45:  uint16(aux_sym_single_quote_string_token1),
	46:  uint16(sym_escape_sequence),
	47:  uint16(sym_stream_redirect),
	48:  uint16(sym_direction),
	49:  uint16(anon_sym_POUND),
	50:  uint16(sym_home_dir_expansion),
	51:  uint16(sym_glob),
	52:  uint16(sym_word),
	53:  uint16(sym__concat),
	54:  uint16(sym__brace_concat),
	55:  uint16(sym__concat_list),
	56:  uint16(anon_sym_LBRACE),
	57:  uint16(sym_program),
	58:  uint16(sym_conditional_execution),
	59:  uint16(sym_pipe),
	60:  uint16(sym_redirect_statement),
	61:  uint16(sym__terminated_statement),
	62:  uint16(sym__terminated_opt_statement),
	63:  uint16(sym_negated_statement),
	64:  uint16(sym__command_substitution_dollar),
	65:  uint16(sym__command_substitution_inner),
	66:  uint16(sym_command_substitution),
	67:  uint16(sym_function_definition),
	68:  uint16(sym_return),
	69:  uint16(sym_switch_statement),
	70:  uint16(sym_case_clause),
	71:  uint16(sym_for_statement),
	72:  uint16(sym_while_statement),
	73:  uint16(sym_if_statement),
	74:  uint16(sym_else_if_clause),
	75:  uint16(sym_else_clause),
	76:  uint16(sym_begin_statement),
	77:  uint16(sym_variable_expansion),
	78:  uint16(sym_index),
	79:  uint16(sym_range),
	80:  uint16(sym_list_element_access),
	81:  uint16(sym_brace_expansion),
	82:  uint16(sym_double_quote_string),
	83:  uint16(sym_single_quote_string),
	84:  uint16(sym_command),
	85:  uint16(sym_file_redirect),
	86:  uint16(sym__special_character),
	87:  uint16(sym_concatenation),
	88:  uint16(sym__expression),
	89:  uint16(sym_concatenation),
	90:  uint16(sym__brace_expression),
	91:  uint16(sym__base_brace_expression),
	92:  uint16(aux_sym_program_repeat1),
	93:  uint16(aux_sym_function_definition_repeat1),
	94:  uint16(aux_sym_function_definition_repeat2),
	95:  uint16(aux_sym_switch_statement_repeat1),
	96:  uint16(aux_sym_case_clause_repeat1),
	97:  uint16(aux_sym_for_statement_repeat1),
	98:  uint16(aux_sym_while_statement_repeat1),
	99:  uint16(aux_sym_if_statement_repeat1),
	100: uint16(aux_sym_variable_expansion_repeat1),
	101: uint16(aux_sym_list_element_access_repeat1),
	102: uint16(aux_sym_brace_expansion_repeat1),
	103: uint16(aux_sym_double_quote_string_repeat1),
	104: uint16(aux_sym_single_quote_string_repeat1),
	105: uint16(aux_sym_command_repeat1),
	106: uint16(aux_sym_concatenation_repeat1),
	107: uint16(aux_sym_brace_concatenation_repeat1),
}

var ts_symbol_metadata = [108]TSSymbolMetadata{
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
	},
	24: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	25: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	36: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	43: {},
	44: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	45: {},
	46: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	47: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	48: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	49: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	50: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	54: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	55: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	57: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	58: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	59: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	60: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	61: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	62: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	63: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	64: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	65: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	91: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	92:  {},
	93:  {},
	94:  {},
	95:  {},
	96:  {},
	97:  {},
	98:  {},
	99:  {},
	100: {},
	101: {},
	102: {},
	103: {},
	104: {},
	105: {},
	106: {},
	107: {},
}

const field_argument = 1
const field_condition = 2
const field_destination = 3
const field_name = 4
const field_operator = 5
const field_option = 6
const field_redirect = 7
const field_value = 8
const field_variable = 9

var ts_field_names = [10]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 1267,
	2: __ccgo_ts + 1276,
	3: __ccgo_ts + 1286,
	4: __ccgo_ts + 1298,
	5: __ccgo_ts + 1303,
	6: __ccgo_ts + 1312,
	7: __ccgo_ts + 1319,
	8: __ccgo_ts + 1328,
	9: __ccgo_ts + 1334,
}

var ts_field_map_slices = [19]TSFieldMapSlice{
	2: {
		Flength: uint16(1),
	},
	4: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(2),
		Flength: uint16(1),
	},
	6: {
		Findex:  uint16(3),
		Flength: uint16(3),
	},
	7: {
		Findex:  uint16(6),
		Flength: uint16(1),
	},
	8: {
		Findex:  uint16(7),
		Flength: uint16(2),
	},
	9: {
		Findex:  uint16(9),
		Flength: uint16(1),
	},
	10: {
		Findex:  uint16(10),
		Flength: uint16(4),
	},
	11: {
		Findex:  uint16(14),
		Flength: uint16(1),
	},
	12: {
		Findex:  uint16(15),
		Flength: uint16(2),
	},
	13: {
		Findex:  uint16(17),
		Flength: uint16(1),
	},
	14: {
		Findex:  uint16(18),
		Flength: uint16(1),
	},
	15: {
		Findex:  uint16(19),
		Flength: uint16(2),
	},
	16: {
		Findex:  uint16(21),
		Flength: uint16(2),
	},
	17: {
		Findex:  uint16(23),
		Flength: uint16(1),
	},
	18: {
		Findex:  uint16(24),
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [26]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_name),
	},
	1: {
		Ffield_id: uint16(field_redirect),
	},
	2: {
		Ffield_id: uint16(field_argument),
	},
	3: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	4: {
		Ffield_id: uint16(field_name),
	},
	5: {
		Ffield_id:    uint16(field_redirect),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	6: {
		Ffield_id: uint16(field_option),
	},
	7: {
		Ffield_id:    uint16(field_destination),
		Fchild_index: uint8(1),
	},
	8: {
		Ffield_id: uint16(field_operator),
	},
	9: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	10: {
		Ffield_id:  uint16(field_argument),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	11: {
		Ffield_id:    uint16(field_argument),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	12: {
		Ffield_id:  uint16(field_redirect),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	13: {
		Ffield_id:    uint16(field_redirect),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	14: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	15: {
		Ffield_id:  uint16(field_option),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	16: {
		Ffield_id:    uint16(field_option),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	17: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
	},
	18: {
		Ffield_id: uint16(field_value),
	},
	19: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	20: {
		Ffield_id:    uint16(field_option),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	21: {
		Ffield_id:  uint16(field_value),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	22: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	23: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(2),
	},
	24: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	25: {
		Ffield_id:    uint16(field_variable),
		Fchild_index: uint8(1),
	},
}

var ts_alias_sequences = [19][7]TSSymbol{
	0: {},
	1: {
		0: uint16(sym_word),
	},
	3: {
		0: uint16(sym_command_substitution),
	},
}

var ts_non_terminal_alias_map = [9]uint16_t{
	0: uint16(sym__command_substitution_dollar),
	1: uint16(2),
	2: uint16(sym__command_substitution_dollar),
	3: uint16(sym_command_substitution),
	4: uint16(sym__special_character),
	5: uint16(2),
	6: uint16(sym__special_character),
	7: uint16(sym_word),
}

var ts_primary_state_ids = [432]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(3),
	4:   uint16(4),
	5:   uint16(5),
	6:   uint16(6),
	7:   uint16(7),
	8:   uint16(8),
	9:   uint16(4),
	10:  uint16(4),
	11:  uint16(11),
	12:  uint16(12),
	13:  uint16(13),
	14:  uint16(14),
	15:  uint16(15),
	16:  uint16(16),
	17:  uint16(17),
	18:  uint16(17),
	19:  uint16(19),
	20:  uint16(17),
	21:  uint16(21),
	22:  uint16(17),
	23:  uint16(19),
	24:  uint16(21),
	25:  uint16(19),
	26:  uint16(17),
	27:  uint16(19),
	28:  uint16(28),
	29:  uint16(17),
	30:  uint16(19),
	31:  uint16(19),
	32:  uint16(17),
	33:  uint16(19),
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
	57:  uint16(36),
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
	68:  uint16(68),
	69:  uint16(69),
	70:  uint16(70),
	71:  uint16(71),
	72:  uint16(72),
	73:  uint16(71),
	74:  uint16(74),
	75:  uint16(71),
	76:  uint16(74),
	77:  uint16(72),
	78:  uint16(78),
	79:  uint16(79),
	80:  uint16(78),
	81:  uint16(79),
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
	97:  uint16(72),
	98:  uint16(72),
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
	116: uint16(116),
	117: uint16(117),
	118: uint16(118),
	119: uint16(119),
	120: uint16(118),
	121: uint16(118),
	122: uint16(118),
	123: uint16(123),
	124: uint16(118),
	125: uint16(123),
	126: uint16(126),
	127: uint16(127),
	128: uint16(128),
	129: uint16(127),
	130: uint16(127),
	131: uint16(127),
	132: uint16(78),
	133: uint16(79),
	134: uint16(78),
	135: uint16(79),
	136: uint16(82),
	137: uint16(84),
	138: uint16(91),
	139: uint16(94),
	140: uint16(89),
	141: uint16(83),
	142: uint16(87),
	143: uint16(88),
	144: uint16(101),
	145: uint16(99),
	146: uint16(105),
	147: uint16(108),
	148: uint16(107),
	149: uint16(102),
	150: uint16(110),
	151: uint16(111),
	152: uint16(112),
	153: uint16(109),
	154: uint16(106),
	155: uint16(100),
	156: uint16(104),
	157: uint16(103),
	158: uint16(158),
	159: uint16(159),
	160: uint16(160),
	161: uint16(82),
	162: uint16(79),
	163: uint16(78),
	164: uint16(78),
	165: uint16(79),
	166: uint16(166),
	167: uint16(166),
	168: uint16(166),
	169: uint16(169),
	170: uint16(169),
	171: uint16(166),
	172: uint16(169),
	173: uint16(166),
	174: uint16(169),
	175: uint16(94),
	176: uint16(89),
	177: uint16(166),
	178: uint16(87),
	179: uint16(91),
	180: uint16(166),
	181: uint16(88),
	182: uint16(84),
	183: uint16(169),
	184: uint16(169),
	185: uint16(185),
	186: uint16(169),
	187: uint16(83),
	188: uint16(106),
	189: uint16(102),
	190: uint16(190),
	191: uint16(191),
	192: uint16(104),
	193: uint16(101),
	194: uint16(107),
	195: uint16(105),
	196: uint16(196),
	197: uint16(112),
	198: uint16(100),
	199: uint16(103),
	200: uint16(99),
	201: uint16(111),
	202: uint16(202),
	203: uint16(110),
	204: uint16(108),
	205: uint16(109),
	206: uint16(206),
	207: uint16(207),
	208: uint16(208),
	209: uint16(209),
	210: uint16(207),
	211: uint16(211),
	212: uint16(212),
	213: uint16(213),
	214: uint16(214),
	215: uint16(215),
	216: uint16(114),
	217: uint16(217),
	218: uint16(218),
	219: uint16(219),
	220: uint16(220),
	221: uint16(221),
	222: uint16(222),
	223: uint16(223),
	224: uint16(208),
	225: uint16(225),
	226: uint16(207),
	227: uint16(227),
	228: uint16(228),
	229: uint16(208),
	230: uint16(230),
	231: uint16(231),
	232: uint16(232),
	233: uint16(233),
	234: uint16(234),
	235: uint16(208),
	236: uint16(236),
	237: uint16(237),
	238: uint16(238),
	239: uint16(207),
	240: uint16(240),
	241: uint16(241),
	242: uint16(242),
	243: uint16(243),
	244: uint16(244),
	245: uint16(245),
	246: uint16(207),
	247: uint16(247),
	248: uint16(208),
	249: uint16(249),
	250: uint16(207),
	251: uint16(251),
	252: uint16(252),
	253: uint16(207),
	254: uint16(254),
	255: uint16(255),
	256: uint16(208),
	257: uint16(257),
	258: uint16(258),
	259: uint16(208),
	260: uint16(260),
	261: uint16(261),
	262: uint16(262),
	263: uint16(260),
	264: uint16(262),
	265: uint16(261),
	266: uint16(260),
	267: uint16(261),
	268: uint16(260),
	269: uint16(79),
	270: uint16(78),
	271: uint16(78),
	272: uint16(79),
	273: uint16(82),
	274: uint16(79),
	275: uint16(78),
	276: uint16(78),
	277: uint16(82),
	278: uint16(91),
	279: uint16(88),
	280: uint16(79),
	281: uint16(83),
	282: uint16(282),
	283: uint16(88),
	284: uint16(284),
	285: uint16(285),
	286: uint16(286),
	287: uint16(287),
	288: uint16(282),
	289: uint16(282),
	290: uint16(287),
	291: uint16(291),
	292: uint16(282),
	293: uint16(293),
	294: uint16(287),
	295: uint16(282),
	296: uint16(83),
	297: uint16(103),
	298: uint16(287),
	299: uint16(287),
	300: uint16(84),
	301: uint16(110),
	302: uint16(108),
	303: uint16(89),
	304: uint16(287),
	305: uint16(101),
	306: uint16(102),
	307: uint16(87),
	308: uint16(282),
	309: uint16(94),
	310: uint16(107),
	311: uint16(91),
	312: uint16(109),
	313: uint16(99),
	314: uint16(106),
	315: uint16(78),
	316: uint16(79),
	317: uint16(109),
	318: uint16(106),
	319: uint16(104),
	320: uint16(82),
	321: uint16(99),
	322: uint16(103),
	323: uint16(100),
	324: uint16(105),
	325: uint16(107),
	326: uint16(326),
	327: uint16(102),
	328: uint16(108),
	329: uint16(110),
	330: uint16(79),
	331: uint16(78),
	332: uint16(101),
	333: uint16(112),
	334: uint16(111),
	335: uint16(335),
	336: uint16(335),
	337: uint16(88),
	338: uint16(91),
	339: uint16(339),
	340: uint16(79),
	341: uint16(335),
	342: uint16(335),
	343: uint16(79),
	344: uint16(78),
	345: uint16(335),
	346: uint16(335),
	347: uint16(82),
	348: uint16(78),
	349: uint16(83),
	350: uint16(350),
	351: uint16(351),
	352: uint16(335),
	353: uint16(353),
	354: uint16(354),
	355: uint16(354),
	356: uint16(353),
	357: uint16(357),
	358: uint16(358),
	359: uint16(359),
	360: uint16(354),
	361: uint16(83),
	362: uint16(109),
	363: uint16(363),
	364: uint16(364),
	365: uint16(106),
	366: uint16(354),
	367: uint16(354),
	368: uint16(99),
	369: uint16(353),
	370: uint16(370),
	371: uint16(353),
	372: uint16(372),
	373: uint16(353),
	374: uint16(374),
	375: uint16(353),
	376: uint16(88),
	377: uint16(377),
	378: uint16(91),
	379: uint16(354),
	380: uint16(107),
	381: uint16(381),
	382: uint16(382),
	383: uint16(110),
	384: uint16(384),
	385: uint16(99),
	386: uint16(386),
	387: uint16(386),
	388: uint16(103),
	389: uint16(108),
	390: uint16(382),
	391: uint16(382),
	392: uint16(392),
	393: uint16(106),
	394: uint16(105),
	395: uint16(382),
	396: uint16(107),
	397: uint16(384),
	398: uint16(386),
	399: uint16(381),
	400: uint16(386),
	401: uint16(384),
	402: uint16(381),
	403: uint16(384),
	404: uint16(386),
	405: uint16(109),
	406: uint16(111),
	407: uint16(407),
	408: uint16(386),
	409: uint16(382),
	410: uint16(381),
	411: uint16(112),
	412: uint16(102),
	413: uint16(381),
	414: uint16(386),
	415: uint16(384),
	416: uint16(101),
	417: uint16(417),
	418: uint16(417),
	419: uint16(417),
	420: uint16(417),
	421: uint16(417),
	422: uint16(417),
	423: uint16(417),
	424: uint16(424),
	425: uint16(425),
	426: uint16(426),
	427: uint16(427),
	428: uint16(428),
	429: uint16(429),
	430: uint16(430),
	431: uint16(431),
}

func anon_sym_LF_character_set_1(tls *libc.TLS, c int32_t) (r uint8) {
	var v1, v2, v3, v4, v6, v7 int32
	var v5, v8 bool
	_, _, _, _, _, _, _, _ = v1, v2, v3, v4, v5, v6, v7, v8
	if c < int32(8192) {
		if c < int32(133) {
			if c < int32(' ') {
				v3 = libc.BoolInt32(c >= int32('\t') && c <= int32('\f'))
			} else {
				v3 = libc.BoolInt32(c <= int32(' '))
			}
			v2 = v3
		} else {
			if v5 = c <= int32(133); !v5 {
				if c < int32(5760) {
					v4 = libc.BoolInt32(c == int32(160))
				} else {
					v4 = libc.BoolInt32(c <= int32(5760))
				}
			}
			v2 = libc.BoolInt32(v5 || v4 != 0)
		}
		v1 = v2
	} else {
		if v8 = c <= int32(8202); !v8 {
			if c < int32(8287) {
				if c < int32(8239) {
					v7 = libc.BoolInt32(c >= int32(8232) && c <= int32(8233))
				} else {
					v7 = libc.BoolInt32(c <= int32(8239))
				}
				v6 = v7
			} else {
				v6 = libc.BoolInt32(c <= int32(8287) || c == int32(12288))
			}
		}
		v1 = libc.BoolInt32(v8 || v6 != 0)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func sym_word_character_set_1(tls *libc.TLS, c int32_t) (r uint8) {
	var v1, v2, v3, v4, v5 int32
	var v6 bool
	_, _, _, _, _, _ = v1, v2, v3, v4, v5, v6
	if c < int32(8192) {
		if c < int32(160) {
			if c < int32(133) {
				v3 = libc.BoolInt32(c >= int32(11) && c <= int32('\f'))
			} else {
				v3 = libc.BoolInt32(c <= int32(133))
			}
			v2 = v3
		} else {
			v2 = libc.BoolInt32(c <= int32(160) || c == int32(5760))
		}
		v1 = v2
	} else {
		if v6 = c <= int32(8202); !v6 {
			if c < int32(8287) {
				if c < int32(8239) {
					v5 = libc.BoolInt32(c >= int32(8232) && c <= int32(8233))
				} else {
					v5 = libc.BoolInt32(c <= int32(8239))
				}
				v4 = v5
			} else {
				v4 = libc.BoolInt32(c <= int32(8287) || c == int32(12288))
			}
		}
		v1 = libc.BoolInt32(v6 || v4 != 0)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func sym_word_character_set_2(tls *libc.TLS, c int32_t) (r uint8) {
	var v1, v2, v3, v4, v6, v7, v8 int32
	var v10, v5, v9 bool
	_, _, _, _, _, _, _, _, _, _ = v1, v10, v2, v3, v4, v5, v6, v7, v8, v9
	if c < int32('$') {
		if c < int32('\r') {
			if c < int32('\t') {
				v3 = libc.BoolInt32(c == 0)
			} else {
				v3 = libc.BoolInt32(c <= int32('\n'))
			}
			v2 = v3
		} else {
			if v5 = c <= int32('\r'); !v5 {
				if c < int32('"') {
					v4 = libc.BoolInt32(c == int32(' '))
				} else {
					v4 = libc.BoolInt32(c <= int32('"'))
				}
			}
			v2 = libc.BoolInt32(v5 || v4 != 0)
		}
		v1 = v2
	} else {
		if v10 = c <= int32('$'); !v10 {
			if c < int32('>') {
				if c < int32(';') {
					v7 = libc.BoolInt32(c >= int32('&') && c <= int32('*'))
				} else {
					v7 = libc.BoolInt32(c <= int32('<'))
				}
				v6 = v7
			} else {
				if v9 = c <= int32('>'); !v9 {
					if c < int32('{') {
						v8 = libc.BoolInt32(c >= int32('[') && c <= int32(']'))
					} else {
						v8 = libc.BoolInt32(c <= int32('~'))
					}
				}
				v6 = libc.BoolInt32(v9 || v8 != 0)
			}
		}
		v1 = libc.BoolInt32(v10 || v6 != 0)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func sym_word_character_set_3(tls *libc.TLS, c int32_t) (r uint8) {
	var v1, v2, v3, v4, v6, v7 int32
	var v5, v8 bool
	_, _, _, _, _, _, _, _ = v1, v2, v3, v4, v5, v6, v7, v8
	if c < int32('$') {
		if c < int32('\r') {
			if c < int32('\t') {
				v3 = libc.BoolInt32(c == 0)
			} else {
				v3 = libc.BoolInt32(c <= int32('\n'))
			}
			v2 = v3
		} else {
			if v5 = c <= int32('\r'); !v5 {
				if c < int32('"') {
					v4 = libc.BoolInt32(c == int32(' '))
				} else {
					v4 = libc.BoolInt32(c <= int32('"'))
				}
			}
			v2 = libc.BoolInt32(v5 || v4 != 0)
		}
		v1 = v2
	} else {
		if v8 = c <= int32('$'); !v8 {
			if c < int32('[') {
				if c < int32(';') {
					v7 = libc.BoolInt32(c >= int32('&') && c <= int32('*'))
				} else {
					v7 = libc.BoolInt32(c <= int32(';'))
				}
				v6 = v7
			} else {
				v6 = libc.BoolInt32(c <= int32(']') || c >= int32('{') && c <= int32('~'))
			}
		}
		v1 = libc.BoolInt32(v8 || v6 != 0)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func sym_word_character_set_4(tls *libc.TLS, c int32_t) (r uint8) {
	var v1, v2, v3, v4, v5 int32
	var v6 bool
	_, _, _, _, _, _ = v1, v2, v3, v4, v5, v6
	if c < int32('&') {
		if c < int32(' ') {
			if c < int32('\t') {
				v3 = libc.BoolInt32(c == 0)
			} else {
				v3 = libc.BoolInt32(c <= int32('\t'))
			}
			v2 = v3
		} else {
			v2 = libc.BoolInt32(c <= int32(' ') || c >= int32('"') && c <= int32('$'))
		}
		v1 = v2
	} else {
		if v6 = c <= int32('*'); !v6 {
			if c < int32('[') {
				if c < int32('>') {
					v5 = libc.BoolInt32(c >= int32(';') && c <= int32('<'))
				} else {
					v5 = libc.BoolInt32(c <= int32('>'))
				}
				v4 = v5
			} else {
				v4 = libc.BoolInt32(c <= int32(']') || c >= int32('{') && c <= int32('~'))
			}
		}
		v1 = libc.BoolInt32(v6 || v4 != 0)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func sym_word_character_set_5(tls *libc.TLS, c int32_t) (r uint8) {
	var v1, v2, v3, v4, v5 int32
	var v6 bool
	_, _, _, _, _, _ = v1, v2, v3, v4, v5, v6
	if c < int32('&') {
		if c < int32(' ') {
			if c < int32('\t') {
				v3 = libc.BoolInt32(c == 0)
			} else {
				v3 = libc.BoolInt32(c <= int32('\r'))
			}
			v2 = v3
		} else {
			v2 = libc.BoolInt32(c <= int32(' ') || c >= int32('"') && c <= int32('$'))
		}
		v1 = v2
	} else {
		if v6 = c <= int32('*'); !v6 {
			if c < int32('[') {
				if c < int32('>') {
					v5 = libc.BoolInt32(c >= int32(';') && c <= int32('<'))
				} else {
					v5 = libc.BoolInt32(c <= int32('>'))
				}
				v4 = v5
			} else {
				v4 = libc.BoolInt32(c <= int32(']') || c >= int32('{') && c <= int32('~'))
			}
		}
		v1 = libc.BoolInt32(v6 || v4 != 0)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func sym_brace_word_character_set_1(tls *libc.TLS, c int32_t) (r uint8) {
	var v1, v2, v3, v4, v6, v7 int32
	var v5, v8 bool
	_, _, _, _, _, _, _, _ = v1, v2, v3, v4, v5, v6, v7, v8
	if c < int32(8192) {
		if c < int32(133) {
			if c < int32(' ') {
				v3 = libc.BoolInt32(c >= int32('\t') && c <= int32('\r'))
			} else {
				v3 = libc.BoolInt32(c <= int32(' '))
			}
			v2 = v3
		} else {
			if v5 = c <= int32(133); !v5 {
				if c < int32(5760) {
					v4 = libc.BoolInt32(c == int32(160))
				} else {
					v4 = libc.BoolInt32(c <= int32(5760))
				}
			}
			v2 = libc.BoolInt32(v5 || v4 != 0)
		}
		v1 = v2
	} else {
		if v8 = c <= int32(8202); !v8 {
			if c < int32(8287) {
				if c < int32(8239) {
					v7 = libc.BoolInt32(c >= int32(8232) && c <= int32(8233))
				} else {
					v7 = libc.BoolInt32(c <= int32(8239))
				}
				v6 = v7
			} else {
				v6 = libc.BoolInt32(c <= int32(8287) || c == int32(12288))
			}
		}
		v1 = libc.BoolInt32(v8 || v6 != 0)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
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
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(113)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(84)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(128)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(108)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(133)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(75)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(60)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(125)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(108)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(75)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(17)
			goto next_state
		}
		if anon_sym_LF_character_set_1(tls, lookahead) != 0 {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('\n') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(24)
			goto next_state
		}
		if anon_sym_LF_character_set_1(tls, lookahead) != 0 {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('\n') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(118)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32(')') && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('|') && lookahead != int32('}') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('\n') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(119)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32(')') && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('|') && lookahead != int32('}') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('\n') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(108)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(75)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(60)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(120)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('\n') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(53)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(121)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32(')') && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('|') && lookahead != int32('}') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(122)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('&') || int32(')') < lookahead) && lookahead != int32(';') && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('|') && lookahead != int32('}') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(123)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('&') || int32(')') < lookahead) && lookahead != int32(';') && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('|') && lookahead != int32('}') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(124)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('&') || int32(')') < lookahead) && lookahead != int32(';') && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('|') && lookahead != int32('}') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(113)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(127)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('&') || int32(')') < lookahead) && lookahead != int32(';') && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('|') && lookahead != int32('}') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32('\n') || lookahead == int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(126)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('&') || int32(')') < lookahead) && lookahead != int32(';') && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('|') && lookahead != int32('}') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(83)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(75)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(63)
			goto next_state
		}
		if sym_brace_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(143)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32(')') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if sym_brace_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(87)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if sym_brace_word_character_set_1(tls, lookahead) != 0 {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(14)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('#') {
			state = uint16(90)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if sym_brace_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(91)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('.') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('<') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(108)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('<') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(109)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('U') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('X') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(99)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(97)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(95)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('a') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('d') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('e') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('e') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('f') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('f') {
			state = uint16(71)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('l') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('s') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('s') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('2') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(30):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(31):
		if eof != 0 {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('\n') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('&') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(75)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(31)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(117)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') && lookahead != int32('>') && lookahead != int32('|') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(32):
		if eof != 0 {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(83)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(26)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(75)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(64)
			goto next_state
		}
		if sym_brace_word_character_set_1(tls, lookahead) != 0 {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_2_GT_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('|') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('&') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(106)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(111)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LF)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(48)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LF)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LF)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(50)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LF)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(51)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LF)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(52)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LF)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(53)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(48)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(50)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(51)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(52)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(53)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(108)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(60)
			goto next_state
		}
		if !(sym_word_character_set_3(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(109)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(60)
			goto next_state
		}
		if !(sym_word_character_set_3(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(140)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(145)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(63)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('$') && (lookahead < int32('\'') || int32('*') < lookahead) && lookahead != int32(',') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_integer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(65)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(66)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('$') && (lookahead < int32('\'') || int32('*') < lookahead) && lookahead != int32(',') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_case)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_case)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_in)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_in)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_if)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_if)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_else)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_else)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_variable_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_double_quote_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(88)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('$') && lookahead != int32('\\') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_double_quote_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(86)
			goto next_state
		}
		if sym_brace_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(87)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('"') || int32('$') < lookahead) && lookahead != int32('\\') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_double_quote_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('$') && lookahead != int32('\\') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_single_quote_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(92)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_single_quote_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(90)
			goto next_state
		}
		if sym_brace_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(91)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_single_quote_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_stream_redirect)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_direction)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_direction)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('&') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_direction)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('&') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(110)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_direction)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('&') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(110)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(106)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_direction)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('&') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_direction)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('>') {
			state = uint16(112)
			goto next_state
		}
		if lookahead == int32('?') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_direction)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('?') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_home_dir_expansion)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_glob)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(117)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\t') && (lookahead < int32(' ') || int32('$') < lookahead) && (lookahead < int32('&') || int32('*') < lookahead) && lookahead != int32(';') && lookahead != int32('<') && lookahead != int32('>') && (lookahead < int32('[') || int32(']') < lookahead) && (lookahead < int32('{') || int32('~') < lookahead) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(118)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\t') && (lookahead < int32(' ') || int32('$') < lookahead) && (lookahead < int32('&') || int32('*') < lookahead) && lookahead != int32(';') && lookahead != int32('<') && lookahead != int32('>') && (lookahead < int32('[') || int32(']') < lookahead) && (lookahead < int32('{') || int32('~') < lookahead) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(119)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\t') && (lookahead < int32(' ') || int32('$') < lookahead) && (lookahead < int32('&') || int32('*') < lookahead) && lookahead != int32(';') && lookahead != int32('<') && lookahead != int32('>') && (lookahead < int32('[') || int32(']') < lookahead) && (lookahead < int32('{') || int32('~') < lookahead) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(61)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(60)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(120)
			goto next_state
		}
		if !(sym_word_character_set_4(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(47)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(53)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(121)
			goto next_state
		}
		if !(sym_word_character_set_4(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(122)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && (lookahead < int32(' ') || int32('$') < lookahead) && (lookahead < int32('&') || int32('*') < lookahead) && lookahead != int32(';') && lookahead != int32('<') && lookahead != int32('>') && (lookahead < int32('[') || int32(']') < lookahead) && (lookahead < int32('{') || int32('~') < lookahead) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(123)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && (lookahead < int32(' ') || int32('$') < lookahead) && (lookahead < int32('&') || int32('*') < lookahead) && lookahead != int32(';') && lookahead != int32('<') && lookahead != int32('>') && (lookahead < int32('[') || int32(']') < lookahead) && (lookahead < int32('{') || int32('~') < lookahead) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(134)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(124)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && (lookahead < int32(' ') || int32('$') < lookahead) && (lookahead < int32('&') || int32('*') < lookahead) && lookahead != int32(';') && lookahead != int32('<') && lookahead != int32('>') && (lookahead < int32('[') || int32(']') < lookahead) && (lookahead < int32('{') || int32('~') < lookahead) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(114)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(84)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(128)
			goto next_state
		}
		if lookahead == int32('2') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(133)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(60)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(125)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && (lookahead < int32(' ') || int32('$') < lookahead) && (lookahead < int32('&') || int32('*') < lookahead) && lookahead != int32(';') && lookahead != int32('<') && lookahead != int32('>') && (lookahead < int32('[') || int32(']') < lookahead) && (lookahead < int32('{') || int32('~') < lookahead) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(126):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(126)
			goto next_state
		}
		if !(sym_word_character_set_5(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(114)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(139)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if sym_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(127)
			goto next_state
		}
		if !(sym_word_character_set_5(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(79)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(137)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(59)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(68)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(132):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(74)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(70)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(72)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(138)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(130)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(130)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(131)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(138):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(132)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(139):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(62)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(140):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(65)
			goto next_state
		}
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(141):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if !(sym_word_character_set_2(tls, lookahead) != 0) {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_brace_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(146)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('$') && (lookahead < int32('\'') || int32('*') < lookahead) && lookahead != int32(',') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(143):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_brace_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(144)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(63)
			goto next_state
		}
		if sym_brace_word_character_set_1(tls, lookahead) != 0 {
			state = uint16(143)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('"') || int32('$') < lookahead) && (lookahead < int32('\'') || int32(',') < lookahead) && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(144):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_brace_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(63)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('$') && (lookahead < int32('\'') || int32('*') < lookahead) && lookahead != int32(',') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_brace_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(66)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('$') && (lookahead < int32('\'') || int32('*') < lookahead) && lookahead != int32(',') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(146):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_brace_word)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('$') && (lookahead < int32('\'') || int32('*') < lookahead) && lookahead != int32(',') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(146)
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
		if lookahead == int32('f') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('w') {
			state = uint16(9)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(133) || lookahead == int32(160) || lookahead == int32(5760) || int32(8192) <= lookahead && lookahead <= int32(8202) || lookahead == int32(8232) || lookahead == int32(8233) || lookahead == int32(8239) || lookahead == int32(8287) || lookahead == int32(12288) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('n') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('e') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('o') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('o') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('o') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('r') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('e') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('w') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('h') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('d') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('g') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('e') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('n') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('r') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('n') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('t') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(17):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_or)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(18):
		if lookahead == int32('t') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('i') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('i') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_and)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(22):
		if lookahead == int32('i') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('a') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('t') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_for)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(26):
		if lookahead == int32('c') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_not)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(28):
		if lookahead == int32('u') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('t') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('l') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('n') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('k') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('i') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('t') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('r') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('c') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('e') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_begin)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_break)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		if lookahead == int32('n') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('i') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('n') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('h') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_while)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		if lookahead == int32('u') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('o') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_return)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_switch)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(49):
		if lookahead == int32('e') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('n') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_continue)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_function)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [432]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	2: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(2),
	},
	3: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(2),
	},
	4: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(2),
	},
	5: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(2),
	},
	6: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(2),
	},
	7: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(2),
	},
	8: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	9: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(2),
	},
	10: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	11: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(2),
	},
	12: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(2),
	},
	13: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(2),
	},
	14: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	15: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(2),
	},
	16: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(2),
	},
	17: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	18: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	19: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	20: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	21: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	22: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	23: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	24: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	25: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	26: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	27: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	28: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	29: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	30: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	31: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	32: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	33: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	34: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(2),
	},
	35: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(2),
	},
	36: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(2),
	},
	37: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	38: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	39: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	40: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	41: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	42: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	43: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	44: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	45: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	46: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	47: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	48: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	49: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	50: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	51: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	52: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	53: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	54: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	55: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	56: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	57: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	58: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	59: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	60: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(2),
	},
	61: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(2),
	},
	62: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(2),
	},
	63: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(2),
	},
	64: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(2),
	},
	65: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(2),
	},
	66: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(2),
	},
	67: {
		Flex_state: uint16(5),
	},
	68: {
		Flex_state: uint16(5),
	},
	69: {
		Flex_state: uint16(5),
	},
	70: {
		Flex_state: uint16(5),
	},
	71: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(2),
	},
	72: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(2),
	},
	73: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(2),
	},
	74: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	75: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	76: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	77: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(2),
	},
	78: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	79: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	80: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	81: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	82: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	83: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	84: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	85: {
		Flex_state: uint16(6),
	},
	86: {
		Flex_state: uint16(6),
	},
	87: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	88: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	89: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	90: {
		Flex_state: uint16(6),
	},
	91: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(3),
	},
	92: {
		Flex_state: uint16(6),
	},
	93: {
		Flex_state: uint16(6),
	},
	94: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	95: {
		Flex_state: uint16(6),
	},
	96: {
		Flex_state: uint16(6),
	},
	97: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(2),
	},
	98: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	99: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	100: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	101: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	102: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	103: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	104: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	105: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	106: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	107: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	108: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	109: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	110: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	111: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	112: {
		Flex_state:          uint16(5),
		Fexternal_lex_state: uint16(4),
	},
	113: {
		Flex_state: uint16(5),
	},
	114: {
		Flex_state: uint16(5),
	},
	115: {
		Flex_state: uint16(5),
	},
	116: {
		Flex_state: uint16(11),
	},
	117: {
		Flex_state: uint16(11),
	},
	118: {
		Flex_state: uint16(12),
	},
	119: {
		Flex_state: uint16(11),
	},
	120: {
		Flex_state: uint16(12),
	},
	121: {
		Flex_state: uint16(12),
	},
	122: {
		Flex_state: uint16(12),
	},
	123: {
		Flex_state: uint16(11),
	},
	124: {
		Flex_state: uint16(12),
	},
	125: {
		Flex_state: uint16(11),
	},
	126: {
		Flex_state: uint16(11),
	},
	127: {
		Flex_state: uint16(10),
	},
	128: {
		Flex_state: uint16(12),
	},
	129: {
		Flex_state: uint16(10),
	},
	130: {
		Flex_state: uint16(10),
	},
	131: {
		Flex_state: uint16(10),
	},
	132: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	133: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	134: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	135: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	136: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	137: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	138: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	139: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	140: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	141: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	142: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	143: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(3),
	},
	144: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	145: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	146: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	147: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	148: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	149: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	150: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	151: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	152: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	153: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	154: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	155: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	156: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	157: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(4),
	},
	158: {
		Flex_state: uint16(6),
	},
	159: {
		Flex_state: uint16(12),
	},
	160: {
		Flex_state: uint16(6),
	},
	161: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(3),
	},
	162: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(3),
	},
	163: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(3),
	},
	164: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(3),
	},
	165: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(3),
	},
	166: {
		Flex_state: uint16(32),
	},
	167: {
		Flex_state: uint16(32),
	},
	168: {
		Flex_state: uint16(32),
	},
	169: {
		Flex_state: uint16(32),
	},
	170: {
		Flex_state: uint16(32),
	},
	171: {
		Flex_state: uint16(32),
	},
	172: {
		Flex_state: uint16(32),
	},
	173: {
		Flex_state: uint16(32),
	},
	174: {
		Flex_state: uint16(32),
	},
	175: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	176: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	177: {
		Flex_state: uint16(32),
	},
	178: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	179: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(3),
	},
	180: {
		Flex_state: uint16(32),
	},
	181: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(3),
	},
	182: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	183: {
		Flex_state: uint16(32),
	},
	184: {
		Flex_state: uint16(32),
	},
	185: {
		Flex_state: uint16(32),
	},
	186: {
		Flex_state: uint16(32),
	},
	187: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(3),
	},
	188: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	189: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	190: {
		Flex_state: uint16(1),
	},
	191: {
		Flex_state: uint16(1),
	},
	192: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	193: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	194: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	195: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	196: {
		Flex_state: uint16(1),
	},
	197: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	198: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	199: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	200: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	201: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	202: {
		Flex_state: uint16(1),
	},
	203: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	204: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	205: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	206: {
		Flex_state: uint16(1),
	},
	207: {
		Flex_state: uint16(1),
	},
	208: {
		Flex_state: uint16(1),
	},
	209: {
		Flex_state: uint16(1),
	},
	210: {
		Flex_state: uint16(1),
	},
	211: {
		Flex_state: uint16(1),
	},
	212: {
		Flex_state: uint16(1),
	},
	213: {
		Flex_state: uint16(1),
	},
	214: {
		Flex_state: uint16(1),
	},
	215: {
		Flex_state: uint16(1),
	},
	216: {
		Flex_state: uint16(1),
	},
	217: {
		Flex_state: uint16(1),
	},
	218: {
		Flex_state: uint16(1),
	},
	219: {
		Flex_state: uint16(1),
	},
	220: {
		Flex_state: uint16(1),
	},
	221: {
		Flex_state: uint16(1),
	},
	222: {
		Flex_state: uint16(1),
	},
	223: {
		Flex_state: uint16(1),
	},
	224: {
		Flex_state: uint16(1),
	},
	225: {
		Flex_state: uint16(1),
	},
	226: {
		Flex_state: uint16(1),
	},
	227: {
		Flex_state: uint16(1),
	},
	228: {
		Flex_state: uint16(1),
	},
	229: {
		Flex_state: uint16(1),
	},
	230: {
		Flex_state: uint16(1),
	},
	231: {
		Flex_state: uint16(1),
	},
	232: {
		Flex_state: uint16(1),
	},
	233: {
		Flex_state: uint16(1),
	},
	234: {
		Flex_state: uint16(1),
	},
	235: {
		Flex_state: uint16(1),
	},
	236: {
		Flex_state: uint16(1),
	},
	237: {
		Flex_state: uint16(1),
	},
	238: {
		Flex_state: uint16(1),
	},
	239: {
		Flex_state: uint16(1),
	},
	240: {
		Flex_state: uint16(1),
	},
	241: {
		Flex_state: uint16(1),
	},
	242: {
		Flex_state: uint16(1),
	},
	243: {
		Flex_state: uint16(1),
	},
	244: {
		Flex_state: uint16(1),
	},
	245: {
		Flex_state: uint16(1),
	},
	246: {
		Flex_state: uint16(1),
	},
	247: {
		Flex_state: uint16(32),
	},
	248: {
		Flex_state: uint16(1),
	},
	249: {
		Flex_state: uint16(1),
	},
	250: {
		Flex_state: uint16(1),
	},
	251: {
		Flex_state: uint16(1),
	},
	252: {
		Flex_state: uint16(32),
	},
	253: {
		Flex_state: uint16(1),
	},
	254: {
		Flex_state: uint16(1),
	},
	255: {
		Flex_state: uint16(1),
	},
	256: {
		Flex_state: uint16(1),
	},
	257: {
		Flex_state: uint16(1),
	},
	258: {
		Flex_state: uint16(1),
	},
	259: {
		Flex_state: uint16(1),
	},
	260: {
		Flex_state: uint16(1),
	},
	261: {
		Flex_state: uint16(1),
	},
	262: {
		Flex_state: uint16(1),
	},
	263: {
		Flex_state: uint16(1),
	},
	264: {
		Flex_state: uint16(1),
	},
	265: {
		Flex_state: uint16(1),
	},
	266: {
		Flex_state: uint16(1),
	},
	267: {
		Flex_state: uint16(1),
	},
	268: {
		Flex_state: uint16(1),
	},
	269: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(5),
	},
	270: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(5),
	},
	271: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(5),
	},
	272: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(5),
	},
	273: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(5),
	},
	274: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	275: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	276: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	277: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	278: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(5),
	},
	279: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(5),
	},
	280: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	281: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(5),
	},
	282: {
		Flex_state: uint16(13),
	},
	283: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	284: {
		Flex_state: uint16(32),
	},
	285: {
		Flex_state: uint16(32),
	},
	286: {
		Flex_state: uint16(32),
	},
	287: {
		Flex_state: uint16(13),
	},
	288: {
		Flex_state: uint16(13),
	},
	289: {
		Flex_state: uint16(13),
	},
	290: {
		Flex_state: uint16(13),
	},
	291: {
		Flex_state: uint16(32),
	},
	292: {
		Flex_state: uint16(13),
	},
	293: {
		Flex_state: uint16(13),
	},
	294: {
		Flex_state: uint16(13),
	},
	295: {
		Flex_state: uint16(13),
	},
	296: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	297: {
		Flex_state: uint16(32),
	},
	298: {
		Flex_state: uint16(13),
	},
	299: {
		Flex_state: uint16(13),
	},
	300: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	301: {
		Flex_state: uint16(32),
	},
	302: {
		Flex_state: uint16(32),
	},
	303: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	304: {
		Flex_state: uint16(13),
	},
	305: {
		Flex_state: uint16(32),
	},
	306: {
		Flex_state: uint16(32),
	},
	307: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	308: {
		Flex_state: uint16(13),
	},
	309: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	310: {
		Flex_state: uint16(32),
	},
	311: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(3),
	},
	312: {
		Flex_state: uint16(32),
	},
	313: {
		Flex_state: uint16(32),
	},
	314: {
		Flex_state: uint16(32),
	},
	315: {
		Flex_state:          uint16(13),
		Fexternal_lex_state: uint16(5),
	},
	316: {
		Flex_state:          uint16(13),
		Fexternal_lex_state: uint16(5),
	},
	317: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	318: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	319: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	320: {
		Flex_state:          uint16(13),
		Fexternal_lex_state: uint16(5),
	},
	321: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	322: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	323: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	324: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	325: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	326: {
		Flex_state: uint16(2),
	},
	327: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	328: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	329: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	330: {
		Flex_state:          uint16(13),
		Fexternal_lex_state: uint16(5),
	},
	331: {
		Flex_state:          uint16(13),
		Fexternal_lex_state: uint16(5),
	},
	332: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	333: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	334: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(4),
	},
	335: {
		Flex_state: uint16(14),
	},
	336: {
		Flex_state: uint16(14),
	},
	337: {
		Flex_state:          uint16(13),
		Fexternal_lex_state: uint16(5),
	},
	338: {
		Flex_state:          uint16(13),
		Fexternal_lex_state: uint16(5),
	},
	339: {
		Flex_state: uint16(32),
	},
	340: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(6),
	},
	341: {
		Flex_state: uint16(14),
	},
	342: {
		Flex_state: uint16(14),
	},
	343: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(6),
	},
	344: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(6),
	},
	345: {
		Flex_state: uint16(14),
	},
	346: {
		Flex_state: uint16(14),
	},
	347: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(6),
	},
	348: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(6),
	},
	349: {
		Flex_state:          uint16(13),
		Fexternal_lex_state: uint16(5),
	},
	350: {
		Flex_state: uint16(2),
	},
	351: {
		Flex_state: uint16(32),
	},
	352: {
		Flex_state: uint16(14),
	},
	353: {
		Flex_state: uint16(15),
	},
	354: {
		Flex_state: uint16(15),
	},
	355: {
		Flex_state: uint16(15),
	},
	356: {
		Flex_state: uint16(15),
	},
	357: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	358: {
		Flex_state: uint16(32),
	},
	359: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	360: {
		Flex_state: uint16(15),
	},
	361: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(6),
	},
	362: {
		Flex_state: uint16(13),
	},
	363: {
		Flex_state: uint16(32),
	},
	364: {
		Flex_state: uint16(13),
	},
	365: {
		Flex_state: uint16(13),
	},
	366: {
		Flex_state: uint16(15),
	},
	367: {
		Flex_state: uint16(15),
	},
	368: {
		Flex_state: uint16(13),
	},
	369: {
		Flex_state: uint16(15),
	},
	370: {
		Flex_state: uint16(15),
	},
	371: {
		Flex_state: uint16(15),
	},
	372: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	373: {
		Flex_state: uint16(15),
	},
	374: {
		Flex_state: uint16(32),
	},
	375: {
		Flex_state: uint16(15),
	},
	376: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(6),
	},
	377: {
		Flex_state: uint16(32),
	},
	378: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(6),
	},
	379: {
		Flex_state: uint16(15),
	},
	380: {
		Flex_state: uint16(13),
	},
	381: {
		Flex_state: uint16(32),
	},
	382: {
		Flex_state: uint16(32),
	},
	383: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	384: {
		Flex_state: uint16(32),
	},
	385: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	386: {
		Flex_state: uint16(14),
	},
	387: {
		Flex_state: uint16(14),
	},
	388: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	389: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	390: {
		Flex_state: uint16(32),
	},
	391: {
		Flex_state: uint16(32),
	},
	392: {
		Flex_state: uint16(32),
	},
	393: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	394: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	395: {
		Flex_state: uint16(32),
	},
	396: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	397: {
		Flex_state: uint16(32),
	},
	398: {
		Flex_state: uint16(14),
	},
	399: {
		Flex_state: uint16(32),
	},
	400: {
		Flex_state: uint16(14),
	},
	401: {
		Flex_state: uint16(32),
	},
	402: {
		Flex_state: uint16(32),
	},
	403: {
		Flex_state: uint16(32),
	},
	404: {
		Flex_state: uint16(14),
	},
	405: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	406: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	407: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	408: {
		Flex_state: uint16(14),
	},
	409: {
		Flex_state: uint16(32),
	},
	410: {
		Flex_state: uint16(32),
	},
	411: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	412: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	413: {
		Flex_state: uint16(32),
	},
	414: {
		Flex_state: uint16(14),
	},
	415: {
		Flex_state: uint16(32),
	},
	416: {
		Flex_state:          uint16(32),
		Fexternal_lex_state: uint16(7),
	},
	417: {
		Flex_state: uint16(32),
	},
	418: {
		Flex_state: uint16(32),
	},
	419: {
		Flex_state: uint16(32),
	},
	420: {
		Flex_state: uint16(32),
	},
	421: {
		Flex_state: uint16(32),
	},
	422: {
		Flex_state: uint16(32),
	},
	423: {
		Flex_state: uint16(32),
	},
	424: {
		Flex_state: uint16(32),
	},
	425: {
		Flex_state: uint16(32),
	},
	426: {
		Flex_state: uint16(32),
	},
	427: {
		Flex_state: uint16(32),
	},
	428: {
		Flex_state: uint16(32),
	},
	429: {
		Flex_state: uint16(14),
	},
	430: {
		Flex_state: uint16(32),
	},
	431: {
		Flex_state: uint16(32),
	},
}

const ts_external_token__concat = 0
const ts_external_token__brace_concat = 1
const ts_external_token__concat_list = 2
const ts_external_token__begin_brace = 3

var ts_external_scanner_symbol_map = [4]TSSymbol{
	0: uint16(sym__concat),
	1: uint16(sym__brace_concat),
	2: uint16(sym__concat_list),
	3: uint16(sym__begin_brace),
}

var ts_external_scanner_states = [8][4]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
		3: libc.BoolUint8(true1 != 0),
	},
	2: {
		3: libc.BoolUint8(true1 != 0),
	},
	3: {
		0: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
	},
	4: {
		0: libc.BoolUint8(true1 != 0),
	},
	5: {
		2: libc.BoolUint8(true1 != 0),
	},
	6: {
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
	},
	7: {
		1: libc.BoolUint8(true1 != 0),
	},
}

var ts_parse_table = [34][108]uint16_t{
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
		35: uint16(3),
		37: uint16(1),
		38: uint16(1),
		39: uint16(1),
		40: uint16(1),
		41: uint16(1),
		42: uint16(1),
		44: uint16(1),
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
	},
	1: {
		0:  uint16(5),
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(11),
		10: uint16(11),
		11: uint16(11),
		12: uint16(11),
		13: uint16(11),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(25),
		27: uint16(25),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		57: uint16(430),
		58: uint16(262),
		59: uint16(262),
		60: uint16(262),
		63: uint16(262),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(262),
		68: uint16(262),
		69: uint16(262),
		71: uint16(262),
		72: uint16(262),
		73: uint16(262),
		76: uint16(262),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(262),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(28),
	},
	2: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(45),
		10: uint16(45),
		11: uint16(45),
		12: uint16(45),
		13: uint16(45),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		20: uint16(47),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(49),
		27: uint16(49),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		32: uint16(51),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(267),
		59: uint16(267),
		60: uint16(267),
		62: uint16(4),
		63: uint16(267),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(267),
		68: uint16(267),
		69: uint16(267),
		71: uint16(267),
		72: uint16(267),
		73: uint16(267),
		74: uint16(339),
		75: uint16(425),
		76: uint16(267),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(267),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(4),
		99: uint16(339),
	},
	3: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(53),
		10: uint16(53),
		11: uint16(53),
		12: uint16(53),
		13: uint16(53),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		20: uint16(55),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(49),
		27: uint16(49),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		32: uint16(51),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(267),
		59: uint16(267),
		60: uint16(267),
		62: uint16(2),
		63: uint16(267),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(267),
		68: uint16(267),
		69: uint16(267),
		71: uint16(267),
		72: uint16(267),
		73: uint16(267),
		74: uint16(351),
		75: uint16(426),
		76: uint16(267),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(267),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(2),
		99: uint16(351),
	},
	4: {
		1:  uint16(57),
		2:  uint16(60),
		3:  uint16(60),
		9:  uint16(63),
		10: uint16(63),
		11: uint16(63),
		12: uint16(63),
		13: uint16(63),
		14: uint16(66),
		15: uint16(66),
		16: uint16(69),
		17: uint16(72),
		19: uint16(75),
		20: uint16(78),
		21: uint16(57),
		22: uint16(57),
		23: uint16(80),
		24: uint16(83),
		26: uint16(86),
		27: uint16(86),
		28: uint16(89),
		30: uint16(92),
		31: uint16(95),
		32: uint16(78),
		33: uint16(98),
		35: uint16(3),
		38: uint16(101),
		39: uint16(101),
		40: uint16(104),
		42: uint16(107),
		44: uint16(110),
		46: uint16(57),
		50: uint16(57),
		51: uint16(57),
		56: uint16(113),
		58: uint16(267),
		59: uint16(267),
		60: uint16(267),
		62: uint16(4),
		63: uint16(267),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(267),
		68: uint16(267),
		69: uint16(267),
		71: uint16(267),
		72: uint16(267),
		73: uint16(267),
		76: uint16(267),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(267),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(4),
	},
	5: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(45),
		10: uint16(45),
		11: uint16(45),
		12: uint16(45),
		13: uint16(45),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		20: uint16(116),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(49),
		27: uint16(49),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		32: uint16(116),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(267),
		59: uint16(267),
		60: uint16(267),
		62: uint16(4),
		63: uint16(267),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(267),
		68: uint16(267),
		69: uint16(267),
		71: uint16(267),
		72: uint16(267),
		73: uint16(267),
		76: uint16(267),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(267),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(4),
	},
	6: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(118),
		10: uint16(118),
		11: uint16(118),
		12: uint16(118),
		13: uint16(118),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		20: uint16(120),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(49),
		27: uint16(49),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		32: uint16(120),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(267),
		59: uint16(267),
		60: uint16(267),
		62: uint16(5),
		63: uint16(267),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(267),
		68: uint16(267),
		69: uint16(267),
		71: uint16(267),
		72: uint16(267),
		73: uint16(267),
		76: uint16(267),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(267),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(5),
	},
	7: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(122),
		10: uint16(122),
		11: uint16(122),
		12: uint16(122),
		13: uint16(122),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		20: uint16(124),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(126),
		27: uint16(126),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(261),
		59: uint16(261),
		60: uint16(261),
		62: uint16(9),
		63: uint16(261),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(261),
		68: uint16(261),
		69: uint16(261),
		71: uint16(261),
		72: uint16(261),
		73: uint16(261),
		76: uint16(261),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(261),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(9),
	},
	8: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(128),
		10: uint16(128),
		11: uint16(128),
		12: uint16(128),
		13: uint16(128),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(130),
		27: uint16(130),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		34: uint16(132),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(222),
		59: uint16(222),
		60: uint16(222),
		62: uint16(14),
		63: uint16(222),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(222),
		68: uint16(222),
		69: uint16(222),
		71: uint16(222),
		72: uint16(222),
		73: uint16(222),
		76: uint16(222),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(222),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(14),
	},
	9: {
		1:  uint16(57),
		2:  uint16(60),
		3:  uint16(60),
		9:  uint16(134),
		10: uint16(134),
		11: uint16(134),
		12: uint16(134),
		13: uint16(134),
		14: uint16(66),
		15: uint16(66),
		16: uint16(69),
		17: uint16(72),
		19: uint16(75),
		20: uint16(78),
		21: uint16(57),
		22: uint16(57),
		23: uint16(80),
		24: uint16(83),
		26: uint16(137),
		27: uint16(137),
		28: uint16(89),
		30: uint16(92),
		31: uint16(95),
		33: uint16(98),
		35: uint16(3),
		38: uint16(101),
		39: uint16(101),
		40: uint16(104),
		42: uint16(107),
		44: uint16(110),
		46: uint16(57),
		50: uint16(57),
		51: uint16(57),
		56: uint16(113),
		58: uint16(261),
		59: uint16(261),
		60: uint16(261),
		62: uint16(9),
		63: uint16(261),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(261),
		68: uint16(261),
		69: uint16(261),
		71: uint16(261),
		72: uint16(261),
		73: uint16(261),
		76: uint16(261),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(261),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(9),
	},
	10: {
		1:  uint16(57),
		2:  uint16(60),
		3:  uint16(60),
		9:  uint16(140),
		10: uint16(140),
		11: uint16(140),
		12: uint16(140),
		13: uint16(140),
		14: uint16(66),
		15: uint16(66),
		16: uint16(69),
		17: uint16(72),
		19: uint16(75),
		21: uint16(57),
		22: uint16(57),
		23: uint16(80),
		24: uint16(83),
		26: uint16(143),
		27: uint16(143),
		28: uint16(89),
		30: uint16(92),
		31: uint16(95),
		33: uint16(98),
		34: uint16(78),
		35: uint16(3),
		38: uint16(101),
		39: uint16(101),
		40: uint16(104),
		42: uint16(107),
		44: uint16(110),
		46: uint16(57),
		50: uint16(57),
		51: uint16(57),
		56: uint16(113),
		58: uint16(265),
		59: uint16(265),
		60: uint16(265),
		62: uint16(10),
		63: uint16(265),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(265),
		68: uint16(265),
		69: uint16(265),
		71: uint16(265),
		72: uint16(265),
		73: uint16(265),
		76: uint16(265),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(265),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(10),
	},
	11: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(122),
		10: uint16(122),
		11: uint16(122),
		12: uint16(122),
		13: uint16(122),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		20: uint16(146),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(126),
		27: uint16(126),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(261),
		59: uint16(261),
		60: uint16(261),
		62: uint16(9),
		63: uint16(261),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(261),
		68: uint16(261),
		69: uint16(261),
		71: uint16(261),
		72: uint16(261),
		73: uint16(261),
		76: uint16(261),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(261),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(9),
	},
	12: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(148),
		10: uint16(148),
		11: uint16(148),
		12: uint16(148),
		13: uint16(148),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		20: uint16(150),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(126),
		27: uint16(126),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(261),
		59: uint16(261),
		60: uint16(261),
		62: uint16(13),
		63: uint16(261),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(261),
		68: uint16(261),
		69: uint16(261),
		71: uint16(261),
		72: uint16(261),
		73: uint16(261),
		76: uint16(261),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(261),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(13),
	},
	13: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(122),
		10: uint16(122),
		11: uint16(122),
		12: uint16(122),
		13: uint16(122),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		20: uint16(152),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(126),
		27: uint16(126),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(261),
		59: uint16(261),
		60: uint16(261),
		62: uint16(9),
		63: uint16(261),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(261),
		68: uint16(261),
		69: uint16(261),
		71: uint16(261),
		72: uint16(261),
		73: uint16(261),
		76: uint16(261),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(261),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(9),
	},
	14: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(154),
		10: uint16(154),
		11: uint16(154),
		12: uint16(154),
		13: uint16(154),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(156),
		27: uint16(156),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		34: uint16(146),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(232),
		59: uint16(232),
		60: uint16(232),
		62: uint16(10),
		63: uint16(232),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(232),
		68: uint16(232),
		69: uint16(232),
		71: uint16(232),
		72: uint16(232),
		73: uint16(232),
		76: uint16(232),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(232),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(10),
	},
	15: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(158),
		10: uint16(158),
		11: uint16(158),
		12: uint16(158),
		13: uint16(158),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		20: uint16(132),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(126),
		27: uint16(126),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(261),
		59: uint16(261),
		60: uint16(261),
		62: uint16(11),
		63: uint16(261),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(261),
		68: uint16(261),
		69: uint16(261),
		71: uint16(261),
		72: uint16(261),
		73: uint16(261),
		76: uint16(261),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(261),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(11),
	},
	16: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(160),
		10: uint16(160),
		11: uint16(160),
		12: uint16(160),
		13: uint16(160),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		20: uint16(162),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(126),
		27: uint16(126),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(261),
		59: uint16(261),
		60: uint16(261),
		62: uint16(7),
		63: uint16(261),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(261),
		68: uint16(261),
		69: uint16(261),
		71: uint16(261),
		72: uint16(261),
		73: uint16(261),
		76: uint16(261),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(261),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		98: uint16(7),
	},
	17: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(164),
		10: uint16(164),
		11: uint16(164),
		12: uint16(164),
		13: uint16(164),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(166),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(168),
		27: uint16(168),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(253),
		59: uint16(253),
		60: uint16(253),
		63: uint16(253),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(253),
		68: uint16(253),
		69: uint16(253),
		71: uint16(253),
		72: uint16(253),
		73: uint16(253),
		76: uint16(253),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(253),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(33),
	},
	18: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(170),
		10: uint16(170),
		11: uint16(170),
		12: uint16(170),
		13: uint16(170),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(172),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(174),
		27: uint16(174),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(250),
		59: uint16(250),
		60: uint16(250),
		63: uint16(250),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(250),
		68: uint16(250),
		69: uint16(250),
		71: uint16(250),
		72: uint16(250),
		73: uint16(250),
		76: uint16(250),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(250),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(19),
	},
	19: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(176),
		10: uint16(176),
		11: uint16(176),
		12: uint16(176),
		13: uint16(176),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(178),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(180),
		27: uint16(180),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(256),
		59: uint16(256),
		60: uint16(256),
		63: uint16(256),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(256),
		68: uint16(256),
		69: uint16(256),
		71: uint16(256),
		72: uint16(256),
		73: uint16(256),
		76: uint16(256),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(256),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(21),
	},
	20: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(182),
		10: uint16(182),
		11: uint16(182),
		12: uint16(182),
		13: uint16(182),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(184),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(186),
		27: uint16(186),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(207),
		59: uint16(207),
		60: uint16(207),
		63: uint16(207),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(207),
		68: uint16(207),
		69: uint16(207),
		71: uint16(207),
		72: uint16(207),
		73: uint16(207),
		76: uint16(207),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(207),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(31),
	},
	21: {
		1:  uint16(188),
		2:  uint16(191),
		3:  uint16(191),
		9:  uint16(194),
		10: uint16(194),
		11: uint16(194),
		12: uint16(194),
		13: uint16(194),
		14: uint16(197),
		15: uint16(197),
		16: uint16(200),
		17: uint16(203),
		18: uint16(206),
		19: uint16(208),
		21: uint16(188),
		22: uint16(188),
		23: uint16(211),
		24: uint16(214),
		26: uint16(217),
		27: uint16(217),
		28: uint16(220),
		30: uint16(223),
		31: uint16(226),
		33: uint16(229),
		35: uint16(3),
		38: uint16(232),
		39: uint16(232),
		40: uint16(235),
		42: uint16(238),
		44: uint16(241),
		46: uint16(188),
		50: uint16(188),
		51: uint16(188),
		56: uint16(244),
		58: uint16(264),
		59: uint16(264),
		60: uint16(264),
		63: uint16(264),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(264),
		68: uint16(264),
		69: uint16(264),
		71: uint16(264),
		72: uint16(264),
		73: uint16(264),
		76: uint16(264),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(264),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(21),
	},
	22: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(247),
		10: uint16(247),
		11: uint16(247),
		12: uint16(247),
		13: uint16(247),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(249),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(251),
		27: uint16(251),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(246),
		59: uint16(246),
		60: uint16(246),
		63: uint16(246),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(246),
		68: uint16(246),
		69: uint16(246),
		71: uint16(246),
		72: uint16(246),
		73: uint16(246),
		76: uint16(246),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(246),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(25),
	},
	23: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(176),
		10: uint16(176),
		11: uint16(176),
		12: uint16(176),
		13: uint16(176),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(253),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(255),
		27: uint16(255),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(235),
		59: uint16(235),
		60: uint16(235),
		63: uint16(235),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(235),
		68: uint16(235),
		69: uint16(235),
		71: uint16(235),
		72: uint16(235),
		73: uint16(235),
		76: uint16(235),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(235),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(21),
	},
	24: {
		0:  uint16(257),
		1:  uint16(188),
		2:  uint16(191),
		3:  uint16(191),
		9:  uint16(259),
		10: uint16(259),
		11: uint16(259),
		12: uint16(259),
		13: uint16(259),
		14: uint16(197),
		15: uint16(197),
		16: uint16(200),
		17: uint16(203),
		19: uint16(208),
		21: uint16(188),
		22: uint16(188),
		23: uint16(211),
		24: uint16(214),
		26: uint16(262),
		27: uint16(262),
		28: uint16(220),
		30: uint16(223),
		31: uint16(226),
		33: uint16(229),
		35: uint16(3),
		38: uint16(232),
		39: uint16(232),
		40: uint16(235),
		42: uint16(238),
		44: uint16(241),
		46: uint16(188),
		50: uint16(188),
		51: uint16(188),
		56: uint16(244),
		58: uint16(262),
		59: uint16(262),
		60: uint16(262),
		63: uint16(262),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(262),
		68: uint16(262),
		69: uint16(262),
		71: uint16(262),
		72: uint16(262),
		73: uint16(262),
		76: uint16(262),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(262),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(24),
	},
	25: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(176),
		10: uint16(176),
		11: uint16(176),
		12: uint16(176),
		13: uint16(176),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(265),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(267),
		27: uint16(267),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(248),
		59: uint16(248),
		60: uint16(248),
		63: uint16(248),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(248),
		68: uint16(248),
		69: uint16(248),
		71: uint16(248),
		72: uint16(248),
		73: uint16(248),
		76: uint16(248),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(248),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(21),
	},
	26: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(269),
		10: uint16(269),
		11: uint16(269),
		12: uint16(269),
		13: uint16(269),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(271),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(273),
		27: uint16(273),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(239),
		59: uint16(239),
		60: uint16(239),
		63: uint16(239),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(239),
		68: uint16(239),
		69: uint16(239),
		71: uint16(239),
		72: uint16(239),
		73: uint16(239),
		76: uint16(239),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(239),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(27),
	},
	27: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(176),
		10: uint16(176),
		11: uint16(176),
		12: uint16(176),
		13: uint16(176),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(275),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(277),
		27: uint16(277),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(259),
		59: uint16(259),
		60: uint16(259),
		63: uint16(259),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(259),
		68: uint16(259),
		69: uint16(259),
		71: uint16(259),
		72: uint16(259),
		73: uint16(259),
		76: uint16(259),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(259),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(21),
	},
	28: {
		0:  uint16(279),
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(281),
		10: uint16(281),
		11: uint16(281),
		12: uint16(281),
		13: uint16(281),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(25),
		27: uint16(25),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(262),
		59: uint16(262),
		60: uint16(262),
		63: uint16(262),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(262),
		68: uint16(262),
		69: uint16(262),
		71: uint16(262),
		72: uint16(262),
		73: uint16(262),
		76: uint16(262),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(262),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(24),
	},
	29: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(283),
		10: uint16(283),
		11: uint16(283),
		12: uint16(283),
		13: uint16(283),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(285),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(287),
		27: uint16(287),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(226),
		59: uint16(226),
		60: uint16(226),
		63: uint16(226),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(226),
		68: uint16(226),
		69: uint16(226),
		71: uint16(226),
		72: uint16(226),
		73: uint16(226),
		76: uint16(226),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(226),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(23),
	},
	30: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(176),
		10: uint16(176),
		11: uint16(176),
		12: uint16(176),
		13: uint16(176),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(289),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(291),
		27: uint16(291),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(208),
		59: uint16(208),
		60: uint16(208),
		63: uint16(208),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(208),
		68: uint16(208),
		69: uint16(208),
		71: uint16(208),
		72: uint16(208),
		73: uint16(208),
		76: uint16(208),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(208),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(21),
	},
	31: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(176),
		10: uint16(176),
		11: uint16(176),
		12: uint16(176),
		13: uint16(176),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(293),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(295),
		27: uint16(295),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(224),
		59: uint16(224),
		60: uint16(224),
		63: uint16(224),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(224),
		68: uint16(224),
		69: uint16(224),
		71: uint16(224),
		72: uint16(224),
		73: uint16(224),
		76: uint16(224),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(224),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(21),
	},
	32: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(297),
		10: uint16(297),
		11: uint16(297),
		12: uint16(297),
		13: uint16(297),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(299),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(301),
		27: uint16(301),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(210),
		59: uint16(210),
		60: uint16(210),
		63: uint16(210),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(210),
		68: uint16(210),
		69: uint16(210),
		71: uint16(210),
		72: uint16(210),
		73: uint16(210),
		76: uint16(210),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(210),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(30),
	},
	33: {
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(9),
		9:  uint16(176),
		10: uint16(176),
		11: uint16(176),
		12: uint16(176),
		13: uint16(176),
		14: uint16(13),
		15: uint16(13),
		16: uint16(15),
		17: uint16(17),
		18: uint16(303),
		19: uint16(19),
		21: uint16(7),
		22: uint16(7),
		23: uint16(21),
		24: uint16(23),
		26: uint16(305),
		27: uint16(305),
		28: uint16(27),
		30: uint16(29),
		31: uint16(31),
		33: uint16(33),
		35: uint16(3),
		38: uint16(35),
		39: uint16(35),
		40: uint16(37),
		42: uint16(39),
		44: uint16(41),
		46: uint16(7),
		50: uint16(7),
		51: uint16(7),
		56: uint16(43),
		58: uint16(229),
		59: uint16(229),
		60: uint16(229),
		63: uint16(229),
		64: uint16(103),
		65: uint16(103),
		66: uint16(84),
		67: uint16(229),
		68: uint16(229),
		69: uint16(229),
		71: uint16(229),
		72: uint16(229),
		73: uint16(229),
		76: uint16(229),
		77: uint16(84),
		81: uint16(84),
		82: uint16(84),
		83: uint16(84),
		84: uint16(229),
		86: uint16(87),
		87: uint16(69),
		88: uint16(69),
		92: uint16(21),
	},
}

var ts_small_parse_table = [12501]uint16_t{
	0:     uint16(26),
	1:     uint16(3),
	2:     uint16(1),
	3:     uint16(sym_comment),
	4:     uint16(15),
	5:     uint16(1),
	6:     uint16(anon_sym_DOLLAR),
	7:     uint16(17),
	8:     uint16(1),
	9:     uint16(anon_sym_LPAREN),
	10:    uint16(19),
	11:    uint16(1),
	12:    uint16(anon_sym_function),
	13:    uint16(21),
	14:    uint16(1),
	15:    uint16(anon_sym_return),
	16:    uint16(23),
	17:    uint16(1),
	18:    uint16(anon_sym_switch),
	19:    uint16(27),
	20:    uint16(1),
	21:    uint16(anon_sym_for),
	22:    uint16(29),
	23:    uint16(1),
	24:    uint16(anon_sym_while),
	25:    uint16(31),
	26:    uint16(1),
	27:    uint16(anon_sym_if),
	28:    uint16(33),
	29:    uint16(1),
	30:    uint16(anon_sym_begin),
	31:    uint16(37),
	32:    uint16(1),
	33:    uint16(anon_sym_LBRACE),
	34:    uint16(39),
	35:    uint16(1),
	36:    uint16(anon_sym_DQUOTE),
	37:    uint16(41),
	38:    uint16(1),
	39:    uint16(anon_sym_SQUOTE),
	40:    uint16(43),
	41:    uint16(1),
	42:    uint16(sym__begin_brace),
	43:    uint16(87),
	44:    uint16(1),
	45:    uint16(sym__special_character),
	46:    uint16(9),
	47:    uint16(2),
	48:    uint16(anon_sym_and),
	49:    uint16(anon_sym_or),
	50:    uint16(13),
	51:    uint16(2),
	52:    uint16(anon_sym_BANG),
	53:    uint16(anon_sym_not),
	54:    uint16(35),
	55:    uint16(2),
	56:    uint16(anon_sym_LBRACK),
	57:    uint16(anon_sym_RBRACK),
	58:    uint16(307),
	59:    uint16(2),
	60:    uint16(anon_sym_end),
	61:    uint16(anon_sym_case),
	62:    uint16(309),
	63:    uint16(2),
	64:    uint16(sym_break),
	65:    uint16(sym_continue),
	66:    uint16(36),
	67:    uint16(2),
	68:    uint16(sym__terminated_statement),
	69:    uint16(aux_sym_function_definition_repeat2),
	70:    uint16(69),
	71:    uint16(2),
	72:    uint16(sym_concatenation),
	73:    uint16(sym__expression),
	74:    uint16(103),
	75:    uint16(2),
	76:    uint16(sym__command_substitution_dollar),
	77:    uint16(sym__command_substitution_inner),
	78:    uint16(84),
	79:    uint16(5),
	80:    uint16(sym_command_substitution),
	81:    uint16(sym_variable_expansion),
	82:    uint16(sym_brace_expansion),
	83:    uint16(sym_double_quote_string),
	84:    uint16(sym_single_quote_string),
	85:    uint16(7),
	86:    uint16(6),
	87:    uint16(sym_integer),
	88:    uint16(sym_float),
	89:    uint16(sym_escape_sequence),
	90:    uint16(sym_home_dir_expansion),
	91:    uint16(sym_glob),
	92:    uint16(sym_word),
	93:    uint16(260),
	94:    uint16(12),
	95:    uint16(sym_conditional_execution),
	96:    uint16(sym_pipe),
	97:    uint16(sym_redirect_statement),
	98:    uint16(sym_negated_statement),
	99:    uint16(sym_function_definition),
	100:   uint16(sym_return),
	101:   uint16(sym_switch_statement),
	102:   uint16(sym_for_statement),
	103:   uint16(sym_while_statement),
	104:   uint16(sym_if_statement),
	105:   uint16(sym_begin_statement),
	106:   uint16(sym_command),
	107:   uint16(26),
	108:   uint16(3),
	109:   uint16(1),
	110:   uint16(sym_comment),
	111:   uint16(15),
	112:   uint16(1),
	113:   uint16(anon_sym_DOLLAR),
	114:   uint16(17),
	115:   uint16(1),
	116:   uint16(anon_sym_LPAREN),
	117:   uint16(19),
	118:   uint16(1),
	119:   uint16(anon_sym_function),
	120:   uint16(21),
	121:   uint16(1),
	122:   uint16(anon_sym_return),
	123:   uint16(23),
	124:   uint16(1),
	125:   uint16(anon_sym_switch),
	126:   uint16(27),
	127:   uint16(1),
	128:   uint16(anon_sym_for),
	129:   uint16(29),
	130:   uint16(1),
	131:   uint16(anon_sym_while),
	132:   uint16(31),
	133:   uint16(1),
	134:   uint16(anon_sym_if),
	135:   uint16(33),
	136:   uint16(1),
	137:   uint16(anon_sym_begin),
	138:   uint16(37),
	139:   uint16(1),
	140:   uint16(anon_sym_LBRACE),
	141:   uint16(39),
	142:   uint16(1),
	143:   uint16(anon_sym_DQUOTE),
	144:   uint16(41),
	145:   uint16(1),
	146:   uint16(anon_sym_SQUOTE),
	147:   uint16(43),
	148:   uint16(1),
	149:   uint16(sym__begin_brace),
	150:   uint16(87),
	151:   uint16(1),
	152:   uint16(sym__special_character),
	153:   uint16(9),
	154:   uint16(2),
	155:   uint16(anon_sym_and),
	156:   uint16(anon_sym_or),
	157:   uint16(13),
	158:   uint16(2),
	159:   uint16(anon_sym_BANG),
	160:   uint16(anon_sym_not),
	161:   uint16(35),
	162:   uint16(2),
	163:   uint16(anon_sym_LBRACK),
	164:   uint16(anon_sym_RBRACK),
	165:   uint16(309),
	166:   uint16(2),
	167:   uint16(sym_break),
	168:   uint16(sym_continue),
	169:   uint16(311),
	170:   uint16(2),
	171:   uint16(anon_sym_end),
	172:   uint16(anon_sym_case),
	173:   uint16(34),
	174:   uint16(2),
	175:   uint16(sym__terminated_statement),
	176:   uint16(aux_sym_function_definition_repeat2),
	177:   uint16(69),
	178:   uint16(2),
	179:   uint16(sym_concatenation),
	180:   uint16(sym__expression),
	181:   uint16(103),
	182:   uint16(2),
	183:   uint16(sym__command_substitution_dollar),
	184:   uint16(sym__command_substitution_inner),
	185:   uint16(84),
	186:   uint16(5),
	187:   uint16(sym_command_substitution),
	188:   uint16(sym_variable_expansion),
	189:   uint16(sym_brace_expansion),
	190:   uint16(sym_double_quote_string),
	191:   uint16(sym_single_quote_string),
	192:   uint16(7),
	193:   uint16(6),
	194:   uint16(sym_integer),
	195:   uint16(sym_float),
	196:   uint16(sym_escape_sequence),
	197:   uint16(sym_home_dir_expansion),
	198:   uint16(sym_glob),
	199:   uint16(sym_word),
	200:   uint16(260),
	201:   uint16(12),
	202:   uint16(sym_conditional_execution),
	203:   uint16(sym_pipe),
	204:   uint16(sym_redirect_statement),
	205:   uint16(sym_negated_statement),
	206:   uint16(sym_function_definition),
	207:   uint16(sym_return),
	208:   uint16(sym_switch_statement),
	209:   uint16(sym_for_statement),
	210:   uint16(sym_while_statement),
	211:   uint16(sym_if_statement),
	212:   uint16(sym_begin_statement),
	213:   uint16(sym_command),
	214:   uint16(26),
	215:   uint16(3),
	216:   uint16(1),
	217:   uint16(sym_comment),
	218:   uint16(322),
	219:   uint16(1),
	220:   uint16(anon_sym_DOLLAR),
	221:   uint16(325),
	222:   uint16(1),
	223:   uint16(anon_sym_LPAREN),
	224:   uint16(328),
	225:   uint16(1),
	226:   uint16(anon_sym_function),
	227:   uint16(333),
	228:   uint16(1),
	229:   uint16(anon_sym_return),
	230:   uint16(336),
	231:   uint16(1),
	232:   uint16(anon_sym_switch),
	233:   uint16(342),
	234:   uint16(1),
	235:   uint16(anon_sym_for),
	236:   uint16(345),
	237:   uint16(1),
	238:   uint16(anon_sym_while),
	239:   uint16(348),
	240:   uint16(1),
	241:   uint16(anon_sym_if),
	242:   uint16(351),
	243:   uint16(1),
	244:   uint16(anon_sym_begin),
	245:   uint16(357),
	246:   uint16(1),
	247:   uint16(anon_sym_LBRACE),
	248:   uint16(360),
	249:   uint16(1),
	250:   uint16(anon_sym_DQUOTE),
	251:   uint16(363),
	252:   uint16(1),
	253:   uint16(anon_sym_SQUOTE),
	254:   uint16(366),
	255:   uint16(1),
	256:   uint16(sym__begin_brace),
	257:   uint16(87),
	258:   uint16(1),
	259:   uint16(sym__special_character),
	260:   uint16(316),
	261:   uint16(2),
	262:   uint16(anon_sym_and),
	263:   uint16(anon_sym_or),
	264:   uint16(319),
	265:   uint16(2),
	266:   uint16(anon_sym_BANG),
	267:   uint16(anon_sym_not),
	268:   uint16(331),
	269:   uint16(2),
	270:   uint16(anon_sym_end),
	271:   uint16(anon_sym_case),
	272:   uint16(339),
	273:   uint16(2),
	274:   uint16(sym_break),
	275:   uint16(sym_continue),
	276:   uint16(354),
	277:   uint16(2),
	278:   uint16(anon_sym_LBRACK),
	279:   uint16(anon_sym_RBRACK),
	280:   uint16(36),
	281:   uint16(2),
	282:   uint16(sym__terminated_statement),
	283:   uint16(aux_sym_function_definition_repeat2),
	284:   uint16(69),
	285:   uint16(2),
	286:   uint16(sym_concatenation),
	287:   uint16(sym__expression),
	288:   uint16(103),
	289:   uint16(2),
	290:   uint16(sym__command_substitution_dollar),
	291:   uint16(sym__command_substitution_inner),
	292:   uint16(84),
	293:   uint16(5),
	294:   uint16(sym_command_substitution),
	295:   uint16(sym_variable_expansion),
	296:   uint16(sym_brace_expansion),
	297:   uint16(sym_double_quote_string),
	298:   uint16(sym_single_quote_string),
	299:   uint16(313),
	300:   uint16(6),
	301:   uint16(sym_integer),
	302:   uint16(sym_float),
	303:   uint16(sym_escape_sequence),
	304:   uint16(sym_home_dir_expansion),
	305:   uint16(sym_glob),
	306:   uint16(sym_word),
	307:   uint16(260),
	308:   uint16(12),
	309:   uint16(sym_conditional_execution),
	310:   uint16(sym_pipe),
	311:   uint16(sym_redirect_statement),
	312:   uint16(sym_negated_statement),
	313:   uint16(sym_function_definition),
	314:   uint16(sym_return),
	315:   uint16(sym_switch_statement),
	316:   uint16(sym_for_statement),
	317:   uint16(sym_while_statement),
	318:   uint16(sym_if_statement),
	319:   uint16(sym_begin_statement),
	320:   uint16(sym_command),
	321:   uint16(26),
	322:   uint16(3),
	323:   uint16(1),
	324:   uint16(sym_comment),
	325:   uint16(15),
	326:   uint16(1),
	327:   uint16(anon_sym_DOLLAR),
	328:   uint16(17),
	329:   uint16(1),
	330:   uint16(anon_sym_LPAREN),
	331:   uint16(19),
	332:   uint16(1),
	333:   uint16(anon_sym_function),
	334:   uint16(21),
	335:   uint16(1),
	336:   uint16(anon_sym_return),
	337:   uint16(23),
	338:   uint16(1),
	339:   uint16(anon_sym_switch),
	340:   uint16(27),
	341:   uint16(1),
	342:   uint16(anon_sym_for),
	343:   uint16(29),
	344:   uint16(1),
	345:   uint16(anon_sym_while),
	346:   uint16(31),
	347:   uint16(1),
	348:   uint16(anon_sym_if),
	349:   uint16(33),
	350:   uint16(1),
	351:   uint16(anon_sym_begin),
	352:   uint16(37),
	353:   uint16(1),
	354:   uint16(anon_sym_LBRACE),
	355:   uint16(39),
	356:   uint16(1),
	357:   uint16(anon_sym_DQUOTE),
	358:   uint16(41),
	359:   uint16(1),
	360:   uint16(anon_sym_SQUOTE),
	361:   uint16(43),
	362:   uint16(1),
	363:   uint16(sym__begin_brace),
	364:   uint16(369),
	365:   uint16(1),
	366:   uint16(anon_sym_end),
	367:   uint16(87),
	368:   uint16(1),
	369:   uint16(sym__special_character),
	370:   uint16(9),
	371:   uint16(2),
	372:   uint16(anon_sym_and),
	373:   uint16(anon_sym_or),
	374:   uint16(13),
	375:   uint16(2),
	376:   uint16(anon_sym_BANG),
	377:   uint16(anon_sym_not),
	378:   uint16(35),
	379:   uint16(2),
	380:   uint16(anon_sym_LBRACK),
	381:   uint16(anon_sym_RBRACK),
	382:   uint16(371),
	383:   uint16(2),
	384:   uint16(sym_break),
	385:   uint16(sym_continue),
	386:   uint16(46),
	387:   uint16(2),
	388:   uint16(sym__terminated_statement),
	389:   uint16(aux_sym_function_definition_repeat2),
	390:   uint16(69),
	391:   uint16(2),
	392:   uint16(sym_concatenation),
	393:   uint16(sym__expression),
	394:   uint16(103),
	395:   uint16(2),
	396:   uint16(sym__command_substitution_dollar),
	397:   uint16(sym__command_substitution_inner),
	398:   uint16(84),
	399:   uint16(5),
	400:   uint16(sym_command_substitution),
	401:   uint16(sym_variable_expansion),
	402:   uint16(sym_brace_expansion),
	403:   uint16(sym_double_quote_string),
	404:   uint16(sym_single_quote_string),
	405:   uint16(7),
	406:   uint16(6),
	407:   uint16(sym_integer),
	408:   uint16(sym_float),
	409:   uint16(sym_escape_sequence),
	410:   uint16(sym_home_dir_expansion),
	411:   uint16(sym_glob),
	412:   uint16(sym_word),
	413:   uint16(263),
	414:   uint16(12),
	415:   uint16(sym_conditional_execution),
	416:   uint16(sym_pipe),
	417:   uint16(sym_redirect_statement),
	418:   uint16(sym_negated_statement),
	419:   uint16(sym_function_definition),
	420:   uint16(sym_return),
	421:   uint16(sym_switch_statement),
	422:   uint16(sym_for_statement),
	423:   uint16(sym_while_statement),
	424:   uint16(sym_if_statement),
	425:   uint16(sym_begin_statement),
	426:   uint16(sym_command),
	427:   uint16(26),
	428:   uint16(3),
	429:   uint16(1),
	430:   uint16(sym_comment),
	431:   uint16(15),
	432:   uint16(1),
	433:   uint16(anon_sym_DOLLAR),
	434:   uint16(17),
	435:   uint16(1),
	436:   uint16(anon_sym_LPAREN),
	437:   uint16(19),
	438:   uint16(1),
	439:   uint16(anon_sym_function),
	440:   uint16(21),
	441:   uint16(1),
	442:   uint16(anon_sym_return),
	443:   uint16(23),
	444:   uint16(1),
	445:   uint16(anon_sym_switch),
	446:   uint16(27),
	447:   uint16(1),
	448:   uint16(anon_sym_for),
	449:   uint16(29),
	450:   uint16(1),
	451:   uint16(anon_sym_while),
	452:   uint16(31),
	453:   uint16(1),
	454:   uint16(anon_sym_if),
	455:   uint16(33),
	456:   uint16(1),
	457:   uint16(anon_sym_begin),
	458:   uint16(37),
	459:   uint16(1),
	460:   uint16(anon_sym_LBRACE),
	461:   uint16(39),
	462:   uint16(1),
	463:   uint16(anon_sym_DQUOTE),
	464:   uint16(41),
	465:   uint16(1),
	466:   uint16(anon_sym_SQUOTE),
	467:   uint16(43),
	468:   uint16(1),
	469:   uint16(sym__begin_brace),
	470:   uint16(373),
	471:   uint16(1),
	472:   uint16(anon_sym_end),
	473:   uint16(87),
	474:   uint16(1),
	475:   uint16(sym__special_character),
	476:   uint16(9),
	477:   uint16(2),
	478:   uint16(anon_sym_and),
	479:   uint16(anon_sym_or),
	480:   uint16(13),
	481:   uint16(2),
	482:   uint16(anon_sym_BANG),
	483:   uint16(anon_sym_not),
	484:   uint16(35),
	485:   uint16(2),
	486:   uint16(anon_sym_LBRACK),
	487:   uint16(anon_sym_RBRACK),
	488:   uint16(371),
	489:   uint16(2),
	490:   uint16(sym_break),
	491:   uint16(sym_continue),
	492:   uint16(49),
	493:   uint16(2),
	494:   uint16(sym__terminated_statement),
	495:   uint16(aux_sym_function_definition_repeat2),
	496:   uint16(69),
	497:   uint16(2),
	498:   uint16(sym_concatenation),
	499:   uint16(sym__expression),
	500:   uint16(103),
	501:   uint16(2),
	502:   uint16(sym__command_substitution_dollar),
	503:   uint16(sym__command_substitution_inner),
	504:   uint16(84),
	505:   uint16(5),
	506:   uint16(sym_command_substitution),
	507:   uint16(sym_variable_expansion),
	508:   uint16(sym_brace_expansion),
	509:   uint16(sym_double_quote_string),
	510:   uint16(sym_single_quote_string),
	511:   uint16(7),
	512:   uint16(6),
	513:   uint16(sym_integer),
	514:   uint16(sym_float),
	515:   uint16(sym_escape_sequence),
	516:   uint16(sym_home_dir_expansion),
	517:   uint16(sym_glob),
	518:   uint16(sym_word),
	519:   uint16(263),
	520:   uint16(12),
	521:   uint16(sym_conditional_execution),
	522:   uint16(sym_pipe),
	523:   uint16(sym_redirect_statement),
	524:   uint16(sym_negated_statement),
	525:   uint16(sym_function_definition),
	526:   uint16(sym_return),
	527:   uint16(sym_switch_statement),
	528:   uint16(sym_for_statement),
	529:   uint16(sym_while_statement),
	530:   uint16(sym_if_statement),
	531:   uint16(sym_begin_statement),
	532:   uint16(sym_command),
	533:   uint16(26),
	534:   uint16(3),
	535:   uint16(1),
	536:   uint16(sym_comment),
	537:   uint16(15),
	538:   uint16(1),
	539:   uint16(anon_sym_DOLLAR),
	540:   uint16(17),
	541:   uint16(1),
	542:   uint16(anon_sym_LPAREN),
	543:   uint16(19),
	544:   uint16(1),
	545:   uint16(anon_sym_function),
	546:   uint16(21),
	547:   uint16(1),
	548:   uint16(anon_sym_return),
	549:   uint16(23),
	550:   uint16(1),
	551:   uint16(anon_sym_switch),
	552:   uint16(27),
	553:   uint16(1),
	554:   uint16(anon_sym_for),
	555:   uint16(29),
	556:   uint16(1),
	557:   uint16(anon_sym_while),
	558:   uint16(31),
	559:   uint16(1),
	560:   uint16(anon_sym_if),
	561:   uint16(33),
	562:   uint16(1),
	563:   uint16(anon_sym_begin),
	564:   uint16(37),
	565:   uint16(1),
	566:   uint16(anon_sym_LBRACE),
	567:   uint16(39),
	568:   uint16(1),
	569:   uint16(anon_sym_DQUOTE),
	570:   uint16(41),
	571:   uint16(1),
	572:   uint16(anon_sym_SQUOTE),
	573:   uint16(43),
	574:   uint16(1),
	575:   uint16(sym__begin_brace),
	576:   uint16(375),
	577:   uint16(1),
	578:   uint16(anon_sym_end),
	579:   uint16(87),
	580:   uint16(1),
	581:   uint16(sym__special_character),
	582:   uint16(9),
	583:   uint16(2),
	584:   uint16(anon_sym_and),
	585:   uint16(anon_sym_or),
	586:   uint16(13),
	587:   uint16(2),
	588:   uint16(anon_sym_BANG),
	589:   uint16(anon_sym_not),
	590:   uint16(35),
	591:   uint16(2),
	592:   uint16(anon_sym_LBRACK),
	593:   uint16(anon_sym_RBRACK),
	594:   uint16(371),
	595:   uint16(2),
	596:   uint16(sym_break),
	597:   uint16(sym_continue),
	598:   uint16(57),
	599:   uint16(2),
	600:   uint16(sym__terminated_statement),
	601:   uint16(aux_sym_function_definition_repeat2),
	602:   uint16(69),
	603:   uint16(2),
	604:   uint16(sym_concatenation),
	605:   uint16(sym__expression),
	606:   uint16(103),
	607:   uint16(2),
	608:   uint16(sym__command_substitution_dollar),
	609:   uint16(sym__command_substitution_inner),
	610:   uint16(84),
	611:   uint16(5),
	612:   uint16(sym_command_substitution),
	613:   uint16(sym_variable_expansion),
	614:   uint16(sym_brace_expansion),
	615:   uint16(sym_double_quote_string),
	616:   uint16(sym_single_quote_string),
	617:   uint16(7),
	618:   uint16(6),
	619:   uint16(sym_integer),
	620:   uint16(sym_float),
	621:   uint16(sym_escape_sequence),
	622:   uint16(sym_home_dir_expansion),
	623:   uint16(sym_glob),
	624:   uint16(sym_word),
	625:   uint16(263),
	626:   uint16(12),
	627:   uint16(sym_conditional_execution),
	628:   uint16(sym_pipe),
	629:   uint16(sym_redirect_statement),
	630:   uint16(sym_negated_statement),
	631:   uint16(sym_function_definition),
	632:   uint16(sym_return),
	633:   uint16(sym_switch_statement),
	634:   uint16(sym_for_statement),
	635:   uint16(sym_while_statement),
	636:   uint16(sym_if_statement),
	637:   uint16(sym_begin_statement),
	638:   uint16(sym_command),
	639:   uint16(26),
	640:   uint16(3),
	641:   uint16(1),
	642:   uint16(sym_comment),
	643:   uint16(15),
	644:   uint16(1),
	645:   uint16(anon_sym_DOLLAR),
	646:   uint16(17),
	647:   uint16(1),
	648:   uint16(anon_sym_LPAREN),
	649:   uint16(19),
	650:   uint16(1),
	651:   uint16(anon_sym_function),
	652:   uint16(21),
	653:   uint16(1),
	654:   uint16(anon_sym_return),
	655:   uint16(23),
	656:   uint16(1),
	657:   uint16(anon_sym_switch),
	658:   uint16(27),
	659:   uint16(1),
	660:   uint16(anon_sym_for),
	661:   uint16(29),
	662:   uint16(1),
	663:   uint16(anon_sym_while),
	664:   uint16(31),
	665:   uint16(1),
	666:   uint16(anon_sym_if),
	667:   uint16(33),
	668:   uint16(1),
	669:   uint16(anon_sym_begin),
	670:   uint16(37),
	671:   uint16(1),
	672:   uint16(anon_sym_LBRACE),
	673:   uint16(39),
	674:   uint16(1),
	675:   uint16(anon_sym_DQUOTE),
	676:   uint16(41),
	677:   uint16(1),
	678:   uint16(anon_sym_SQUOTE),
	679:   uint16(43),
	680:   uint16(1),
	681:   uint16(sym__begin_brace),
	682:   uint16(377),
	683:   uint16(1),
	684:   uint16(anon_sym_end),
	685:   uint16(87),
	686:   uint16(1),
	687:   uint16(sym__special_character),
	688:   uint16(9),
	689:   uint16(2),
	690:   uint16(anon_sym_and),
	691:   uint16(anon_sym_or),
	692:   uint16(13),
	693:   uint16(2),
	694:   uint16(anon_sym_BANG),
	695:   uint16(anon_sym_not),
	696:   uint16(35),
	697:   uint16(2),
	698:   uint16(anon_sym_LBRACK),
	699:   uint16(anon_sym_RBRACK),
	700:   uint16(371),
	701:   uint16(2),
	702:   uint16(sym_break),
	703:   uint16(sym_continue),
	704:   uint16(47),
	705:   uint16(2),
	706:   uint16(sym__terminated_statement),
	707:   uint16(aux_sym_function_definition_repeat2),
	708:   uint16(69),
	709:   uint16(2),
	710:   uint16(sym_concatenation),
	711:   uint16(sym__expression),
	712:   uint16(103),
	713:   uint16(2),
	714:   uint16(sym__command_substitution_dollar),
	715:   uint16(sym__command_substitution_inner),
	716:   uint16(84),
	717:   uint16(5),
	718:   uint16(sym_command_substitution),
	719:   uint16(sym_variable_expansion),
	720:   uint16(sym_brace_expansion),
	721:   uint16(sym_double_quote_string),
	722:   uint16(sym_single_quote_string),
	723:   uint16(7),
	724:   uint16(6),
	725:   uint16(sym_integer),
	726:   uint16(sym_float),
	727:   uint16(sym_escape_sequence),
	728:   uint16(sym_home_dir_expansion),
	729:   uint16(sym_glob),
	730:   uint16(sym_word),
	731:   uint16(263),
	732:   uint16(12),
	733:   uint16(sym_conditional_execution),
	734:   uint16(sym_pipe),
	735:   uint16(sym_redirect_statement),
	736:   uint16(sym_negated_statement),
	737:   uint16(sym_function_definition),
	738:   uint16(sym_return),
	739:   uint16(sym_switch_statement),
	740:   uint16(sym_for_statement),
	741:   uint16(sym_while_statement),
	742:   uint16(sym_if_statement),
	743:   uint16(sym_begin_statement),
	744:   uint16(sym_command),
	745:   uint16(26),
	746:   uint16(3),
	747:   uint16(1),
	748:   uint16(sym_comment),
	749:   uint16(15),
	750:   uint16(1),
	751:   uint16(anon_sym_DOLLAR),
	752:   uint16(17),
	753:   uint16(1),
	754:   uint16(anon_sym_LPAREN),
	755:   uint16(19),
	756:   uint16(1),
	757:   uint16(anon_sym_function),
	758:   uint16(21),
	759:   uint16(1),
	760:   uint16(anon_sym_return),
	761:   uint16(23),
	762:   uint16(1),
	763:   uint16(anon_sym_switch),
	764:   uint16(27),
	765:   uint16(1),
	766:   uint16(anon_sym_for),
	767:   uint16(29),
	768:   uint16(1),
	769:   uint16(anon_sym_while),
	770:   uint16(31),
	771:   uint16(1),
	772:   uint16(anon_sym_if),
	773:   uint16(33),
	774:   uint16(1),
	775:   uint16(anon_sym_begin),
	776:   uint16(37),
	777:   uint16(1),
	778:   uint16(anon_sym_LBRACE),
	779:   uint16(39),
	780:   uint16(1),
	781:   uint16(anon_sym_DQUOTE),
	782:   uint16(41),
	783:   uint16(1),
	784:   uint16(anon_sym_SQUOTE),
	785:   uint16(43),
	786:   uint16(1),
	787:   uint16(sym__begin_brace),
	788:   uint16(379),
	789:   uint16(1),
	790:   uint16(anon_sym_end),
	791:   uint16(87),
	792:   uint16(1),
	793:   uint16(sym__special_character),
	794:   uint16(9),
	795:   uint16(2),
	796:   uint16(anon_sym_and),
	797:   uint16(anon_sym_or),
	798:   uint16(13),
	799:   uint16(2),
	800:   uint16(anon_sym_BANG),
	801:   uint16(anon_sym_not),
	802:   uint16(35),
	803:   uint16(2),
	804:   uint16(anon_sym_LBRACK),
	805:   uint16(anon_sym_RBRACK),
	806:   uint16(371),
	807:   uint16(2),
	808:   uint16(sym_break),
	809:   uint16(sym_continue),
	810:   uint16(57),
	811:   uint16(2),
	812:   uint16(sym__terminated_statement),
	813:   uint16(aux_sym_function_definition_repeat2),
	814:   uint16(69),
	815:   uint16(2),
	816:   uint16(sym_concatenation),
	817:   uint16(sym__expression),
	818:   uint16(103),
	819:   uint16(2),
	820:   uint16(sym__command_substitution_dollar),
	821:   uint16(sym__command_substitution_inner),
	822:   uint16(84),
	823:   uint16(5),
	824:   uint16(sym_command_substitution),
	825:   uint16(sym_variable_expansion),
	826:   uint16(sym_brace_expansion),
	827:   uint16(sym_double_quote_string),
	828:   uint16(sym_single_quote_string),
	829:   uint16(7),
	830:   uint16(6),
	831:   uint16(sym_integer),
	832:   uint16(sym_float),
	833:   uint16(sym_escape_sequence),
	834:   uint16(sym_home_dir_expansion),
	835:   uint16(sym_glob),
	836:   uint16(sym_word),
	837:   uint16(263),
	838:   uint16(12),
	839:   uint16(sym_conditional_execution),
	840:   uint16(sym_pipe),
	841:   uint16(sym_redirect_statement),
	842:   uint16(sym_negated_statement),
	843:   uint16(sym_function_definition),
	844:   uint16(sym_return),
	845:   uint16(sym_switch_statement),
	846:   uint16(sym_for_statement),
	847:   uint16(sym_while_statement),
	848:   uint16(sym_if_statement),
	849:   uint16(sym_begin_statement),
	850:   uint16(sym_command),
	851:   uint16(26),
	852:   uint16(3),
	853:   uint16(1),
	854:   uint16(sym_comment),
	855:   uint16(15),
	856:   uint16(1),
	857:   uint16(anon_sym_DOLLAR),
	858:   uint16(17),
	859:   uint16(1),
	860:   uint16(anon_sym_LPAREN),
	861:   uint16(19),
	862:   uint16(1),
	863:   uint16(anon_sym_function),
	864:   uint16(21),
	865:   uint16(1),
	866:   uint16(anon_sym_return),
	867:   uint16(23),
	868:   uint16(1),
	869:   uint16(anon_sym_switch),
	870:   uint16(27),
	871:   uint16(1),
	872:   uint16(anon_sym_for),
	873:   uint16(29),
	874:   uint16(1),
	875:   uint16(anon_sym_while),
	876:   uint16(31),
	877:   uint16(1),
	878:   uint16(anon_sym_if),
	879:   uint16(33),
	880:   uint16(1),
	881:   uint16(anon_sym_begin),
	882:   uint16(37),
	883:   uint16(1),
	884:   uint16(anon_sym_LBRACE),
	885:   uint16(39),
	886:   uint16(1),
	887:   uint16(anon_sym_DQUOTE),
	888:   uint16(41),
	889:   uint16(1),
	890:   uint16(anon_sym_SQUOTE),
	891:   uint16(43),
	892:   uint16(1),
	893:   uint16(sym__begin_brace),
	894:   uint16(381),
	895:   uint16(1),
	896:   uint16(anon_sym_end),
	897:   uint16(87),
	898:   uint16(1),
	899:   uint16(sym__special_character),
	900:   uint16(9),
	901:   uint16(2),
	902:   uint16(anon_sym_and),
	903:   uint16(anon_sym_or),
	904:   uint16(13),
	905:   uint16(2),
	906:   uint16(anon_sym_BANG),
	907:   uint16(anon_sym_not),
	908:   uint16(35),
	909:   uint16(2),
	910:   uint16(anon_sym_LBRACK),
	911:   uint16(anon_sym_RBRACK),
	912:   uint16(371),
	913:   uint16(2),
	914:   uint16(sym_break),
	915:   uint16(sym_continue),
	916:   uint16(45),
	917:   uint16(2),
	918:   uint16(sym__terminated_statement),
	919:   uint16(aux_sym_function_definition_repeat2),
	920:   uint16(69),
	921:   uint16(2),
	922:   uint16(sym_concatenation),
	923:   uint16(sym__expression),
	924:   uint16(103),
	925:   uint16(2),
	926:   uint16(sym__command_substitution_dollar),
	927:   uint16(sym__command_substitution_inner),
	928:   uint16(84),
	929:   uint16(5),
	930:   uint16(sym_command_substitution),
	931:   uint16(sym_variable_expansion),
	932:   uint16(sym_brace_expansion),
	933:   uint16(sym_double_quote_string),
	934:   uint16(sym_single_quote_string),
	935:   uint16(7),
	936:   uint16(6),
	937:   uint16(sym_integer),
	938:   uint16(sym_float),
	939:   uint16(sym_escape_sequence),
	940:   uint16(sym_home_dir_expansion),
	941:   uint16(sym_glob),
	942:   uint16(sym_word),
	943:   uint16(263),
	944:   uint16(12),
	945:   uint16(sym_conditional_execution),
	946:   uint16(sym_pipe),
	947:   uint16(sym_redirect_statement),
	948:   uint16(sym_negated_statement),
	949:   uint16(sym_function_definition),
	950:   uint16(sym_return),
	951:   uint16(sym_switch_statement),
	952:   uint16(sym_for_statement),
	953:   uint16(sym_while_statement),
	954:   uint16(sym_if_statement),
	955:   uint16(sym_begin_statement),
	956:   uint16(sym_command),
	957:   uint16(26),
	958:   uint16(3),
	959:   uint16(1),
	960:   uint16(sym_comment),
	961:   uint16(15),
	962:   uint16(1),
	963:   uint16(anon_sym_DOLLAR),
	964:   uint16(17),
	965:   uint16(1),
	966:   uint16(anon_sym_LPAREN),
	967:   uint16(19),
	968:   uint16(1),
	969:   uint16(anon_sym_function),
	970:   uint16(21),
	971:   uint16(1),
	972:   uint16(anon_sym_return),
	973:   uint16(23),
	974:   uint16(1),
	975:   uint16(anon_sym_switch),
	976:   uint16(27),
	977:   uint16(1),
	978:   uint16(anon_sym_for),
	979:   uint16(29),
	980:   uint16(1),
	981:   uint16(anon_sym_while),
	982:   uint16(31),
	983:   uint16(1),
	984:   uint16(anon_sym_if),
	985:   uint16(33),
	986:   uint16(1),
	987:   uint16(anon_sym_begin),
	988:   uint16(37),
	989:   uint16(1),
	990:   uint16(anon_sym_LBRACE),
	991:   uint16(39),
	992:   uint16(1),
	993:   uint16(anon_sym_DQUOTE),
	994:   uint16(41),
	995:   uint16(1),
	996:   uint16(anon_sym_SQUOTE),
	997:   uint16(43),
	998:   uint16(1),
	999:   uint16(sym__begin_brace),
	1000:  uint16(383),
	1001:  uint16(1),
	1002:  uint16(anon_sym_end),
	1003:  uint16(87),
	1004:  uint16(1),
	1005:  uint16(sym__special_character),
	1006:  uint16(9),
	1007:  uint16(2),
	1008:  uint16(anon_sym_and),
	1009:  uint16(anon_sym_or),
	1010:  uint16(13),
	1011:  uint16(2),
	1012:  uint16(anon_sym_BANG),
	1013:  uint16(anon_sym_not),
	1014:  uint16(35),
	1015:  uint16(2),
	1016:  uint16(anon_sym_LBRACK),
	1017:  uint16(anon_sym_RBRACK),
	1018:  uint16(371),
	1019:  uint16(2),
	1020:  uint16(sym_break),
	1021:  uint16(sym_continue),
	1022:  uint16(57),
	1023:  uint16(2),
	1024:  uint16(sym__terminated_statement),
	1025:  uint16(aux_sym_function_definition_repeat2),
	1026:  uint16(69),
	1027:  uint16(2),
	1028:  uint16(sym_concatenation),
	1029:  uint16(sym__expression),
	1030:  uint16(103),
	1031:  uint16(2),
	1032:  uint16(sym__command_substitution_dollar),
	1033:  uint16(sym__command_substitution_inner),
	1034:  uint16(84),
	1035:  uint16(5),
	1036:  uint16(sym_command_substitution),
	1037:  uint16(sym_variable_expansion),
	1038:  uint16(sym_brace_expansion),
	1039:  uint16(sym_double_quote_string),
	1040:  uint16(sym_single_quote_string),
	1041:  uint16(7),
	1042:  uint16(6),
	1043:  uint16(sym_integer),
	1044:  uint16(sym_float),
	1045:  uint16(sym_escape_sequence),
	1046:  uint16(sym_home_dir_expansion),
	1047:  uint16(sym_glob),
	1048:  uint16(sym_word),
	1049:  uint16(263),
	1050:  uint16(12),
	1051:  uint16(sym_conditional_execution),
	1052:  uint16(sym_pipe),
	1053:  uint16(sym_redirect_statement),
	1054:  uint16(sym_negated_statement),
	1055:  uint16(sym_function_definition),
	1056:  uint16(sym_return),
	1057:  uint16(sym_switch_statement),
	1058:  uint16(sym_for_statement),
	1059:  uint16(sym_while_statement),
	1060:  uint16(sym_if_statement),
	1061:  uint16(sym_begin_statement),
	1062:  uint16(sym_command),
	1063:  uint16(26),
	1064:  uint16(3),
	1065:  uint16(1),
	1066:  uint16(sym_comment),
	1067:  uint16(15),
	1068:  uint16(1),
	1069:  uint16(anon_sym_DOLLAR),
	1070:  uint16(17),
	1071:  uint16(1),
	1072:  uint16(anon_sym_LPAREN),
	1073:  uint16(19),
	1074:  uint16(1),
	1075:  uint16(anon_sym_function),
	1076:  uint16(21),
	1077:  uint16(1),
	1078:  uint16(anon_sym_return),
	1079:  uint16(23),
	1080:  uint16(1),
	1081:  uint16(anon_sym_switch),
	1082:  uint16(27),
	1083:  uint16(1),
	1084:  uint16(anon_sym_for),
	1085:  uint16(29),
	1086:  uint16(1),
	1087:  uint16(anon_sym_while),
	1088:  uint16(31),
	1089:  uint16(1),
	1090:  uint16(anon_sym_if),
	1091:  uint16(33),
	1092:  uint16(1),
	1093:  uint16(anon_sym_begin),
	1094:  uint16(37),
	1095:  uint16(1),
	1096:  uint16(anon_sym_LBRACE),
	1097:  uint16(39),
	1098:  uint16(1),
	1099:  uint16(anon_sym_DQUOTE),
	1100:  uint16(41),
	1101:  uint16(1),
	1102:  uint16(anon_sym_SQUOTE),
	1103:  uint16(43),
	1104:  uint16(1),
	1105:  uint16(sym__begin_brace),
	1106:  uint16(385),
	1107:  uint16(1),
	1108:  uint16(anon_sym_end),
	1109:  uint16(87),
	1110:  uint16(1),
	1111:  uint16(sym__special_character),
	1112:  uint16(9),
	1113:  uint16(2),
	1114:  uint16(anon_sym_and),
	1115:  uint16(anon_sym_or),
	1116:  uint16(13),
	1117:  uint16(2),
	1118:  uint16(anon_sym_BANG),
	1119:  uint16(anon_sym_not),
	1120:  uint16(35),
	1121:  uint16(2),
	1122:  uint16(anon_sym_LBRACK),
	1123:  uint16(anon_sym_RBRACK),
	1124:  uint16(371),
	1125:  uint16(2),
	1126:  uint16(sym_break),
	1127:  uint16(sym_continue),
	1128:  uint16(57),
	1129:  uint16(2),
	1130:  uint16(sym__terminated_statement),
	1131:  uint16(aux_sym_function_definition_repeat2),
	1132:  uint16(69),
	1133:  uint16(2),
	1134:  uint16(sym_concatenation),
	1135:  uint16(sym__expression),
	1136:  uint16(103),
	1137:  uint16(2),
	1138:  uint16(sym__command_substitution_dollar),
	1139:  uint16(sym__command_substitution_inner),
	1140:  uint16(84),
	1141:  uint16(5),
	1142:  uint16(sym_command_substitution),
	1143:  uint16(sym_variable_expansion),
	1144:  uint16(sym_brace_expansion),
	1145:  uint16(sym_double_quote_string),
	1146:  uint16(sym_single_quote_string),
	1147:  uint16(7),
	1148:  uint16(6),
	1149:  uint16(sym_integer),
	1150:  uint16(sym_float),
	1151:  uint16(sym_escape_sequence),
	1152:  uint16(sym_home_dir_expansion),
	1153:  uint16(sym_glob),
	1154:  uint16(sym_word),
	1155:  uint16(263),
	1156:  uint16(12),
	1157:  uint16(sym_conditional_execution),
	1158:  uint16(sym_pipe),
	1159:  uint16(sym_redirect_statement),
	1160:  uint16(sym_negated_statement),
	1161:  uint16(sym_function_definition),
	1162:  uint16(sym_return),
	1163:  uint16(sym_switch_statement),
	1164:  uint16(sym_for_statement),
	1165:  uint16(sym_while_statement),
	1166:  uint16(sym_if_statement),
	1167:  uint16(sym_begin_statement),
	1168:  uint16(sym_command),
	1169:  uint16(26),
	1170:  uint16(3),
	1171:  uint16(1),
	1172:  uint16(sym_comment),
	1173:  uint16(15),
	1174:  uint16(1),
	1175:  uint16(anon_sym_DOLLAR),
	1176:  uint16(17),
	1177:  uint16(1),
	1178:  uint16(anon_sym_LPAREN),
	1179:  uint16(19),
	1180:  uint16(1),
	1181:  uint16(anon_sym_function),
	1182:  uint16(21),
	1183:  uint16(1),
	1184:  uint16(anon_sym_return),
	1185:  uint16(23),
	1186:  uint16(1),
	1187:  uint16(anon_sym_switch),
	1188:  uint16(27),
	1189:  uint16(1),
	1190:  uint16(anon_sym_for),
	1191:  uint16(29),
	1192:  uint16(1),
	1193:  uint16(anon_sym_while),
	1194:  uint16(31),
	1195:  uint16(1),
	1196:  uint16(anon_sym_if),
	1197:  uint16(33),
	1198:  uint16(1),
	1199:  uint16(anon_sym_begin),
	1200:  uint16(37),
	1201:  uint16(1),
	1202:  uint16(anon_sym_LBRACE),
	1203:  uint16(39),
	1204:  uint16(1),
	1205:  uint16(anon_sym_DQUOTE),
	1206:  uint16(41),
	1207:  uint16(1),
	1208:  uint16(anon_sym_SQUOTE),
	1209:  uint16(43),
	1210:  uint16(1),
	1211:  uint16(sym__begin_brace),
	1212:  uint16(387),
	1213:  uint16(1),
	1214:  uint16(anon_sym_end),
	1215:  uint16(87),
	1216:  uint16(1),
	1217:  uint16(sym__special_character),
	1218:  uint16(9),
	1219:  uint16(2),
	1220:  uint16(anon_sym_and),
	1221:  uint16(anon_sym_or),
	1222:  uint16(13),
	1223:  uint16(2),
	1224:  uint16(anon_sym_BANG),
	1225:  uint16(anon_sym_not),
	1226:  uint16(35),
	1227:  uint16(2),
	1228:  uint16(anon_sym_LBRACK),
	1229:  uint16(anon_sym_RBRACK),
	1230:  uint16(371),
	1231:  uint16(2),
	1232:  uint16(sym_break),
	1233:  uint16(sym_continue),
	1234:  uint16(57),
	1235:  uint16(2),
	1236:  uint16(sym__terminated_statement),
	1237:  uint16(aux_sym_function_definition_repeat2),
	1238:  uint16(69),
	1239:  uint16(2),
	1240:  uint16(sym_concatenation),
	1241:  uint16(sym__expression),
	1242:  uint16(103),
	1243:  uint16(2),
	1244:  uint16(sym__command_substitution_dollar),
	1245:  uint16(sym__command_substitution_inner),
	1246:  uint16(84),
	1247:  uint16(5),
	1248:  uint16(sym_command_substitution),
	1249:  uint16(sym_variable_expansion),
	1250:  uint16(sym_brace_expansion),
	1251:  uint16(sym_double_quote_string),
	1252:  uint16(sym_single_quote_string),
	1253:  uint16(7),
	1254:  uint16(6),
	1255:  uint16(sym_integer),
	1256:  uint16(sym_float),
	1257:  uint16(sym_escape_sequence),
	1258:  uint16(sym_home_dir_expansion),
	1259:  uint16(sym_glob),
	1260:  uint16(sym_word),
	1261:  uint16(263),
	1262:  uint16(12),
	1263:  uint16(sym_conditional_execution),
	1264:  uint16(sym_pipe),
	1265:  uint16(sym_redirect_statement),
	1266:  uint16(sym_negated_statement),
	1267:  uint16(sym_function_definition),
	1268:  uint16(sym_return),
	1269:  uint16(sym_switch_statement),
	1270:  uint16(sym_for_statement),
	1271:  uint16(sym_while_statement),
	1272:  uint16(sym_if_statement),
	1273:  uint16(sym_begin_statement),
	1274:  uint16(sym_command),
	1275:  uint16(26),
	1276:  uint16(3),
	1277:  uint16(1),
	1278:  uint16(sym_comment),
	1279:  uint16(15),
	1280:  uint16(1),
	1281:  uint16(anon_sym_DOLLAR),
	1282:  uint16(17),
	1283:  uint16(1),
	1284:  uint16(anon_sym_LPAREN),
	1285:  uint16(19),
	1286:  uint16(1),
	1287:  uint16(anon_sym_function),
	1288:  uint16(21),
	1289:  uint16(1),
	1290:  uint16(anon_sym_return),
	1291:  uint16(23),
	1292:  uint16(1),
	1293:  uint16(anon_sym_switch),
	1294:  uint16(27),
	1295:  uint16(1),
	1296:  uint16(anon_sym_for),
	1297:  uint16(29),
	1298:  uint16(1),
	1299:  uint16(anon_sym_while),
	1300:  uint16(31),
	1301:  uint16(1),
	1302:  uint16(anon_sym_if),
	1303:  uint16(33),
	1304:  uint16(1),
	1305:  uint16(anon_sym_begin),
	1306:  uint16(37),
	1307:  uint16(1),
	1308:  uint16(anon_sym_LBRACE),
	1309:  uint16(39),
	1310:  uint16(1),
	1311:  uint16(anon_sym_DQUOTE),
	1312:  uint16(41),
	1313:  uint16(1),
	1314:  uint16(anon_sym_SQUOTE),
	1315:  uint16(43),
	1316:  uint16(1),
	1317:  uint16(sym__begin_brace),
	1318:  uint16(389),
	1319:  uint16(1),
	1320:  uint16(anon_sym_end),
	1321:  uint16(87),
	1322:  uint16(1),
	1323:  uint16(sym__special_character),
	1324:  uint16(9),
	1325:  uint16(2),
	1326:  uint16(anon_sym_and),
	1327:  uint16(anon_sym_or),
	1328:  uint16(13),
	1329:  uint16(2),
	1330:  uint16(anon_sym_BANG),
	1331:  uint16(anon_sym_not),
	1332:  uint16(35),
	1333:  uint16(2),
	1334:  uint16(anon_sym_LBRACK),
	1335:  uint16(anon_sym_RBRACK),
	1336:  uint16(371),
	1337:  uint16(2),
	1338:  uint16(sym_break),
	1339:  uint16(sym_continue),
	1340:  uint16(57),
	1341:  uint16(2),
	1342:  uint16(sym__terminated_statement),
	1343:  uint16(aux_sym_function_definition_repeat2),
	1344:  uint16(69),
	1345:  uint16(2),
	1346:  uint16(sym_concatenation),
	1347:  uint16(sym__expression),
	1348:  uint16(103),
	1349:  uint16(2),
	1350:  uint16(sym__command_substitution_dollar),
	1351:  uint16(sym__command_substitution_inner),
	1352:  uint16(84),
	1353:  uint16(5),
	1354:  uint16(sym_command_substitution),
	1355:  uint16(sym_variable_expansion),
	1356:  uint16(sym_brace_expansion),
	1357:  uint16(sym_double_quote_string),
	1358:  uint16(sym_single_quote_string),
	1359:  uint16(7),
	1360:  uint16(6),
	1361:  uint16(sym_integer),
	1362:  uint16(sym_float),
	1363:  uint16(sym_escape_sequence),
	1364:  uint16(sym_home_dir_expansion),
	1365:  uint16(sym_glob),
	1366:  uint16(sym_word),
	1367:  uint16(263),
	1368:  uint16(12),
	1369:  uint16(sym_conditional_execution),
	1370:  uint16(sym_pipe),
	1371:  uint16(sym_redirect_statement),
	1372:  uint16(sym_negated_statement),
	1373:  uint16(sym_function_definition),
	1374:  uint16(sym_return),
	1375:  uint16(sym_switch_statement),
	1376:  uint16(sym_for_statement),
	1377:  uint16(sym_while_statement),
	1378:  uint16(sym_if_statement),
	1379:  uint16(sym_begin_statement),
	1380:  uint16(sym_command),
	1381:  uint16(26),
	1382:  uint16(3),
	1383:  uint16(1),
	1384:  uint16(sym_comment),
	1385:  uint16(15),
	1386:  uint16(1),
	1387:  uint16(anon_sym_DOLLAR),
	1388:  uint16(17),
	1389:  uint16(1),
	1390:  uint16(anon_sym_LPAREN),
	1391:  uint16(19),
	1392:  uint16(1),
	1393:  uint16(anon_sym_function),
	1394:  uint16(21),
	1395:  uint16(1),
	1396:  uint16(anon_sym_return),
	1397:  uint16(23),
	1398:  uint16(1),
	1399:  uint16(anon_sym_switch),
	1400:  uint16(27),
	1401:  uint16(1),
	1402:  uint16(anon_sym_for),
	1403:  uint16(29),
	1404:  uint16(1),
	1405:  uint16(anon_sym_while),
	1406:  uint16(31),
	1407:  uint16(1),
	1408:  uint16(anon_sym_if),
	1409:  uint16(33),
	1410:  uint16(1),
	1411:  uint16(anon_sym_begin),
	1412:  uint16(37),
	1413:  uint16(1),
	1414:  uint16(anon_sym_LBRACE),
	1415:  uint16(39),
	1416:  uint16(1),
	1417:  uint16(anon_sym_DQUOTE),
	1418:  uint16(41),
	1419:  uint16(1),
	1420:  uint16(anon_sym_SQUOTE),
	1421:  uint16(43),
	1422:  uint16(1),
	1423:  uint16(sym__begin_brace),
	1424:  uint16(391),
	1425:  uint16(1),
	1426:  uint16(anon_sym_end),
	1427:  uint16(87),
	1428:  uint16(1),
	1429:  uint16(sym__special_character),
	1430:  uint16(9),
	1431:  uint16(2),
	1432:  uint16(anon_sym_and),
	1433:  uint16(anon_sym_or),
	1434:  uint16(13),
	1435:  uint16(2),
	1436:  uint16(anon_sym_BANG),
	1437:  uint16(anon_sym_not),
	1438:  uint16(35),
	1439:  uint16(2),
	1440:  uint16(anon_sym_LBRACK),
	1441:  uint16(anon_sym_RBRACK),
	1442:  uint16(371),
	1443:  uint16(2),
	1444:  uint16(sym_break),
	1445:  uint16(sym_continue),
	1446:  uint16(57),
	1447:  uint16(2),
	1448:  uint16(sym__terminated_statement),
	1449:  uint16(aux_sym_function_definition_repeat2),
	1450:  uint16(69),
	1451:  uint16(2),
	1452:  uint16(sym_concatenation),
	1453:  uint16(sym__expression),
	1454:  uint16(103),
	1455:  uint16(2),
	1456:  uint16(sym__command_substitution_dollar),
	1457:  uint16(sym__command_substitution_inner),
	1458:  uint16(84),
	1459:  uint16(5),
	1460:  uint16(sym_command_substitution),
	1461:  uint16(sym_variable_expansion),
	1462:  uint16(sym_brace_expansion),
	1463:  uint16(sym_double_quote_string),
	1464:  uint16(sym_single_quote_string),
	1465:  uint16(7),
	1466:  uint16(6),
	1467:  uint16(sym_integer),
	1468:  uint16(sym_float),
	1469:  uint16(sym_escape_sequence),
	1470:  uint16(sym_home_dir_expansion),
	1471:  uint16(sym_glob),
	1472:  uint16(sym_word),
	1473:  uint16(263),
	1474:  uint16(12),
	1475:  uint16(sym_conditional_execution),
	1476:  uint16(sym_pipe),
	1477:  uint16(sym_redirect_statement),
	1478:  uint16(sym_negated_statement),
	1479:  uint16(sym_function_definition),
	1480:  uint16(sym_return),
	1481:  uint16(sym_switch_statement),
	1482:  uint16(sym_for_statement),
	1483:  uint16(sym_while_statement),
	1484:  uint16(sym_if_statement),
	1485:  uint16(sym_begin_statement),
	1486:  uint16(sym_command),
	1487:  uint16(26),
	1488:  uint16(3),
	1489:  uint16(1),
	1490:  uint16(sym_comment),
	1491:  uint16(15),
	1492:  uint16(1),
	1493:  uint16(anon_sym_DOLLAR),
	1494:  uint16(17),
	1495:  uint16(1),
	1496:  uint16(anon_sym_LPAREN),
	1497:  uint16(19),
	1498:  uint16(1),
	1499:  uint16(anon_sym_function),
	1500:  uint16(21),
	1501:  uint16(1),
	1502:  uint16(anon_sym_return),
	1503:  uint16(23),
	1504:  uint16(1),
	1505:  uint16(anon_sym_switch),
	1506:  uint16(27),
	1507:  uint16(1),
	1508:  uint16(anon_sym_for),
	1509:  uint16(29),
	1510:  uint16(1),
	1511:  uint16(anon_sym_while),
	1512:  uint16(31),
	1513:  uint16(1),
	1514:  uint16(anon_sym_if),
	1515:  uint16(33),
	1516:  uint16(1),
	1517:  uint16(anon_sym_begin),
	1518:  uint16(37),
	1519:  uint16(1),
	1520:  uint16(anon_sym_LBRACE),
	1521:  uint16(39),
	1522:  uint16(1),
	1523:  uint16(anon_sym_DQUOTE),
	1524:  uint16(41),
	1525:  uint16(1),
	1526:  uint16(anon_sym_SQUOTE),
	1527:  uint16(43),
	1528:  uint16(1),
	1529:  uint16(sym__begin_brace),
	1530:  uint16(393),
	1531:  uint16(1),
	1532:  uint16(anon_sym_end),
	1533:  uint16(87),
	1534:  uint16(1),
	1535:  uint16(sym__special_character),
	1536:  uint16(9),
	1537:  uint16(2),
	1538:  uint16(anon_sym_and),
	1539:  uint16(anon_sym_or),
	1540:  uint16(13),
	1541:  uint16(2),
	1542:  uint16(anon_sym_BANG),
	1543:  uint16(anon_sym_not),
	1544:  uint16(35),
	1545:  uint16(2),
	1546:  uint16(anon_sym_LBRACK),
	1547:  uint16(anon_sym_RBRACK),
	1548:  uint16(371),
	1549:  uint16(2),
	1550:  uint16(sym_break),
	1551:  uint16(sym_continue),
	1552:  uint16(57),
	1553:  uint16(2),
	1554:  uint16(sym__terminated_statement),
	1555:  uint16(aux_sym_function_definition_repeat2),
	1556:  uint16(69),
	1557:  uint16(2),
	1558:  uint16(sym_concatenation),
	1559:  uint16(sym__expression),
	1560:  uint16(103),
	1561:  uint16(2),
	1562:  uint16(sym__command_substitution_dollar),
	1563:  uint16(sym__command_substitution_inner),
	1564:  uint16(84),
	1565:  uint16(5),
	1566:  uint16(sym_command_substitution),
	1567:  uint16(sym_variable_expansion),
	1568:  uint16(sym_brace_expansion),
	1569:  uint16(sym_double_quote_string),
	1570:  uint16(sym_single_quote_string),
	1571:  uint16(7),
	1572:  uint16(6),
	1573:  uint16(sym_integer),
	1574:  uint16(sym_float),
	1575:  uint16(sym_escape_sequence),
	1576:  uint16(sym_home_dir_expansion),
	1577:  uint16(sym_glob),
	1578:  uint16(sym_word),
	1579:  uint16(263),
	1580:  uint16(12),
	1581:  uint16(sym_conditional_execution),
	1582:  uint16(sym_pipe),
	1583:  uint16(sym_redirect_statement),
	1584:  uint16(sym_negated_statement),
	1585:  uint16(sym_function_definition),
	1586:  uint16(sym_return),
	1587:  uint16(sym_switch_statement),
	1588:  uint16(sym_for_statement),
	1589:  uint16(sym_while_statement),
	1590:  uint16(sym_if_statement),
	1591:  uint16(sym_begin_statement),
	1592:  uint16(sym_command),
	1593:  uint16(26),
	1594:  uint16(3),
	1595:  uint16(1),
	1596:  uint16(sym_comment),
	1597:  uint16(15),
	1598:  uint16(1),
	1599:  uint16(anon_sym_DOLLAR),
	1600:  uint16(17),
	1601:  uint16(1),
	1602:  uint16(anon_sym_LPAREN),
	1603:  uint16(19),
	1604:  uint16(1),
	1605:  uint16(anon_sym_function),
	1606:  uint16(21),
	1607:  uint16(1),
	1608:  uint16(anon_sym_return),
	1609:  uint16(23),
	1610:  uint16(1),
	1611:  uint16(anon_sym_switch),
	1612:  uint16(27),
	1613:  uint16(1),
	1614:  uint16(anon_sym_for),
	1615:  uint16(29),
	1616:  uint16(1),
	1617:  uint16(anon_sym_while),
	1618:  uint16(31),
	1619:  uint16(1),
	1620:  uint16(anon_sym_if),
	1621:  uint16(33),
	1622:  uint16(1),
	1623:  uint16(anon_sym_begin),
	1624:  uint16(37),
	1625:  uint16(1),
	1626:  uint16(anon_sym_LBRACE),
	1627:  uint16(39),
	1628:  uint16(1),
	1629:  uint16(anon_sym_DQUOTE),
	1630:  uint16(41),
	1631:  uint16(1),
	1632:  uint16(anon_sym_SQUOTE),
	1633:  uint16(43),
	1634:  uint16(1),
	1635:  uint16(sym__begin_brace),
	1636:  uint16(395),
	1637:  uint16(1),
	1638:  uint16(anon_sym_end),
	1639:  uint16(87),
	1640:  uint16(1),
	1641:  uint16(sym__special_character),
	1642:  uint16(9),
	1643:  uint16(2),
	1644:  uint16(anon_sym_and),
	1645:  uint16(anon_sym_or),
	1646:  uint16(13),
	1647:  uint16(2),
	1648:  uint16(anon_sym_BANG),
	1649:  uint16(anon_sym_not),
	1650:  uint16(35),
	1651:  uint16(2),
	1652:  uint16(anon_sym_LBRACK),
	1653:  uint16(anon_sym_RBRACK),
	1654:  uint16(371),
	1655:  uint16(2),
	1656:  uint16(sym_break),
	1657:  uint16(sym_continue),
	1658:  uint16(57),
	1659:  uint16(2),
	1660:  uint16(sym__terminated_statement),
	1661:  uint16(aux_sym_function_definition_repeat2),
	1662:  uint16(69),
	1663:  uint16(2),
	1664:  uint16(sym_concatenation),
	1665:  uint16(sym__expression),
	1666:  uint16(103),
	1667:  uint16(2),
	1668:  uint16(sym__command_substitution_dollar),
	1669:  uint16(sym__command_substitution_inner),
	1670:  uint16(84),
	1671:  uint16(5),
	1672:  uint16(sym_command_substitution),
	1673:  uint16(sym_variable_expansion),
	1674:  uint16(sym_brace_expansion),
	1675:  uint16(sym_double_quote_string),
	1676:  uint16(sym_single_quote_string),
	1677:  uint16(7),
	1678:  uint16(6),
	1679:  uint16(sym_integer),
	1680:  uint16(sym_float),
	1681:  uint16(sym_escape_sequence),
	1682:  uint16(sym_home_dir_expansion),
	1683:  uint16(sym_glob),
	1684:  uint16(sym_word),
	1685:  uint16(263),
	1686:  uint16(12),
	1687:  uint16(sym_conditional_execution),
	1688:  uint16(sym_pipe),
	1689:  uint16(sym_redirect_statement),
	1690:  uint16(sym_negated_statement),
	1691:  uint16(sym_function_definition),
	1692:  uint16(sym_return),
	1693:  uint16(sym_switch_statement),
	1694:  uint16(sym_for_statement),
	1695:  uint16(sym_while_statement),
	1696:  uint16(sym_if_statement),
	1697:  uint16(sym_begin_statement),
	1698:  uint16(sym_command),
	1699:  uint16(26),
	1700:  uint16(3),
	1701:  uint16(1),
	1702:  uint16(sym_comment),
	1703:  uint16(15),
	1704:  uint16(1),
	1705:  uint16(anon_sym_DOLLAR),
	1706:  uint16(17),
	1707:  uint16(1),
	1708:  uint16(anon_sym_LPAREN),
	1709:  uint16(19),
	1710:  uint16(1),
	1711:  uint16(anon_sym_function),
	1712:  uint16(21),
	1713:  uint16(1),
	1714:  uint16(anon_sym_return),
	1715:  uint16(23),
	1716:  uint16(1),
	1717:  uint16(anon_sym_switch),
	1718:  uint16(27),
	1719:  uint16(1),
	1720:  uint16(anon_sym_for),
	1721:  uint16(29),
	1722:  uint16(1),
	1723:  uint16(anon_sym_while),
	1724:  uint16(31),
	1725:  uint16(1),
	1726:  uint16(anon_sym_if),
	1727:  uint16(33),
	1728:  uint16(1),
	1729:  uint16(anon_sym_begin),
	1730:  uint16(37),
	1731:  uint16(1),
	1732:  uint16(anon_sym_LBRACE),
	1733:  uint16(39),
	1734:  uint16(1),
	1735:  uint16(anon_sym_DQUOTE),
	1736:  uint16(41),
	1737:  uint16(1),
	1738:  uint16(anon_sym_SQUOTE),
	1739:  uint16(43),
	1740:  uint16(1),
	1741:  uint16(sym__begin_brace),
	1742:  uint16(397),
	1743:  uint16(1),
	1744:  uint16(anon_sym_end),
	1745:  uint16(87),
	1746:  uint16(1),
	1747:  uint16(sym__special_character),
	1748:  uint16(9),
	1749:  uint16(2),
	1750:  uint16(anon_sym_and),
	1751:  uint16(anon_sym_or),
	1752:  uint16(13),
	1753:  uint16(2),
	1754:  uint16(anon_sym_BANG),
	1755:  uint16(anon_sym_not),
	1756:  uint16(35),
	1757:  uint16(2),
	1758:  uint16(anon_sym_LBRACK),
	1759:  uint16(anon_sym_RBRACK),
	1760:  uint16(371),
	1761:  uint16(2),
	1762:  uint16(sym_break),
	1763:  uint16(sym_continue),
	1764:  uint16(57),
	1765:  uint16(2),
	1766:  uint16(sym__terminated_statement),
	1767:  uint16(aux_sym_function_definition_repeat2),
	1768:  uint16(69),
	1769:  uint16(2),
	1770:  uint16(sym_concatenation),
	1771:  uint16(sym__expression),
	1772:  uint16(103),
	1773:  uint16(2),
	1774:  uint16(sym__command_substitution_dollar),
	1775:  uint16(sym__command_substitution_inner),
	1776:  uint16(84),
	1777:  uint16(5),
	1778:  uint16(sym_command_substitution),
	1779:  uint16(sym_variable_expansion),
	1780:  uint16(sym_brace_expansion),
	1781:  uint16(sym_double_quote_string),
	1782:  uint16(sym_single_quote_string),
	1783:  uint16(7),
	1784:  uint16(6),
	1785:  uint16(sym_integer),
	1786:  uint16(sym_float),
	1787:  uint16(sym_escape_sequence),
	1788:  uint16(sym_home_dir_expansion),
	1789:  uint16(sym_glob),
	1790:  uint16(sym_word),
	1791:  uint16(263),
	1792:  uint16(12),
	1793:  uint16(sym_conditional_execution),
	1794:  uint16(sym_pipe),
	1795:  uint16(sym_redirect_statement),
	1796:  uint16(sym_negated_statement),
	1797:  uint16(sym_function_definition),
	1798:  uint16(sym_return),
	1799:  uint16(sym_switch_statement),
	1800:  uint16(sym_for_statement),
	1801:  uint16(sym_while_statement),
	1802:  uint16(sym_if_statement),
	1803:  uint16(sym_begin_statement),
	1804:  uint16(sym_command),
	1805:  uint16(26),
	1806:  uint16(3),
	1807:  uint16(1),
	1808:  uint16(sym_comment),
	1809:  uint16(15),
	1810:  uint16(1),
	1811:  uint16(anon_sym_DOLLAR),
	1812:  uint16(17),
	1813:  uint16(1),
	1814:  uint16(anon_sym_LPAREN),
	1815:  uint16(19),
	1816:  uint16(1),
	1817:  uint16(anon_sym_function),
	1818:  uint16(21),
	1819:  uint16(1),
	1820:  uint16(anon_sym_return),
	1821:  uint16(23),
	1822:  uint16(1),
	1823:  uint16(anon_sym_switch),
	1824:  uint16(27),
	1825:  uint16(1),
	1826:  uint16(anon_sym_for),
	1827:  uint16(29),
	1828:  uint16(1),
	1829:  uint16(anon_sym_while),
	1830:  uint16(31),
	1831:  uint16(1),
	1832:  uint16(anon_sym_if),
	1833:  uint16(33),
	1834:  uint16(1),
	1835:  uint16(anon_sym_begin),
	1836:  uint16(37),
	1837:  uint16(1),
	1838:  uint16(anon_sym_LBRACE),
	1839:  uint16(39),
	1840:  uint16(1),
	1841:  uint16(anon_sym_DQUOTE),
	1842:  uint16(41),
	1843:  uint16(1),
	1844:  uint16(anon_sym_SQUOTE),
	1845:  uint16(43),
	1846:  uint16(1),
	1847:  uint16(sym__begin_brace),
	1848:  uint16(399),
	1849:  uint16(1),
	1850:  uint16(anon_sym_end),
	1851:  uint16(87),
	1852:  uint16(1),
	1853:  uint16(sym__special_character),
	1854:  uint16(9),
	1855:  uint16(2),
	1856:  uint16(anon_sym_and),
	1857:  uint16(anon_sym_or),
	1858:  uint16(13),
	1859:  uint16(2),
	1860:  uint16(anon_sym_BANG),
	1861:  uint16(anon_sym_not),
	1862:  uint16(35),
	1863:  uint16(2),
	1864:  uint16(anon_sym_LBRACK),
	1865:  uint16(anon_sym_RBRACK),
	1866:  uint16(371),
	1867:  uint16(2),
	1868:  uint16(sym_break),
	1869:  uint16(sym_continue),
	1870:  uint16(39),
	1871:  uint16(2),
	1872:  uint16(sym__terminated_statement),
	1873:  uint16(aux_sym_function_definition_repeat2),
	1874:  uint16(69),
	1875:  uint16(2),
	1876:  uint16(sym_concatenation),
	1877:  uint16(sym__expression),
	1878:  uint16(103),
	1879:  uint16(2),
	1880:  uint16(sym__command_substitution_dollar),
	1881:  uint16(sym__command_substitution_inner),
	1882:  uint16(84),
	1883:  uint16(5),
	1884:  uint16(sym_command_substitution),
	1885:  uint16(sym_variable_expansion),
	1886:  uint16(sym_brace_expansion),
	1887:  uint16(sym_double_quote_string),
	1888:  uint16(sym_single_quote_string),
	1889:  uint16(7),
	1890:  uint16(6),
	1891:  uint16(sym_integer),
	1892:  uint16(sym_float),
	1893:  uint16(sym_escape_sequence),
	1894:  uint16(sym_home_dir_expansion),
	1895:  uint16(sym_glob),
	1896:  uint16(sym_word),
	1897:  uint16(263),
	1898:  uint16(12),
	1899:  uint16(sym_conditional_execution),
	1900:  uint16(sym_pipe),
	1901:  uint16(sym_redirect_statement),
	1902:  uint16(sym_negated_statement),
	1903:  uint16(sym_function_definition),
	1904:  uint16(sym_return),
	1905:  uint16(sym_switch_statement),
	1906:  uint16(sym_for_statement),
	1907:  uint16(sym_while_statement),
	1908:  uint16(sym_if_statement),
	1909:  uint16(sym_begin_statement),
	1910:  uint16(sym_command),
	1911:  uint16(26),
	1912:  uint16(3),
	1913:  uint16(1),
	1914:  uint16(sym_comment),
	1915:  uint16(15),
	1916:  uint16(1),
	1917:  uint16(anon_sym_DOLLAR),
	1918:  uint16(17),
	1919:  uint16(1),
	1920:  uint16(anon_sym_LPAREN),
	1921:  uint16(19),
	1922:  uint16(1),
	1923:  uint16(anon_sym_function),
	1924:  uint16(21),
	1925:  uint16(1),
	1926:  uint16(anon_sym_return),
	1927:  uint16(23),
	1928:  uint16(1),
	1929:  uint16(anon_sym_switch),
	1930:  uint16(27),
	1931:  uint16(1),
	1932:  uint16(anon_sym_for),
	1933:  uint16(29),
	1934:  uint16(1),
	1935:  uint16(anon_sym_while),
	1936:  uint16(31),
	1937:  uint16(1),
	1938:  uint16(anon_sym_if),
	1939:  uint16(33),
	1940:  uint16(1),
	1941:  uint16(anon_sym_begin),
	1942:  uint16(37),
	1943:  uint16(1),
	1944:  uint16(anon_sym_LBRACE),
	1945:  uint16(39),
	1946:  uint16(1),
	1947:  uint16(anon_sym_DQUOTE),
	1948:  uint16(41),
	1949:  uint16(1),
	1950:  uint16(anon_sym_SQUOTE),
	1951:  uint16(43),
	1952:  uint16(1),
	1953:  uint16(sym__begin_brace),
	1954:  uint16(401),
	1955:  uint16(1),
	1956:  uint16(anon_sym_end),
	1957:  uint16(87),
	1958:  uint16(1),
	1959:  uint16(sym__special_character),
	1960:  uint16(9),
	1961:  uint16(2),
	1962:  uint16(anon_sym_and),
	1963:  uint16(anon_sym_or),
	1964:  uint16(13),
	1965:  uint16(2),
	1966:  uint16(anon_sym_BANG),
	1967:  uint16(anon_sym_not),
	1968:  uint16(35),
	1969:  uint16(2),
	1970:  uint16(anon_sym_LBRACK),
	1971:  uint16(anon_sym_RBRACK),
	1972:  uint16(371),
	1973:  uint16(2),
	1974:  uint16(sym_break),
	1975:  uint16(sym_continue),
	1976:  uint16(41),
	1977:  uint16(2),
	1978:  uint16(sym__terminated_statement),
	1979:  uint16(aux_sym_function_definition_repeat2),
	1980:  uint16(69),
	1981:  uint16(2),
	1982:  uint16(sym_concatenation),
	1983:  uint16(sym__expression),
	1984:  uint16(103),
	1985:  uint16(2),
	1986:  uint16(sym__command_substitution_dollar),
	1987:  uint16(sym__command_substitution_inner),
	1988:  uint16(84),
	1989:  uint16(5),
	1990:  uint16(sym_command_substitution),
	1991:  uint16(sym_variable_expansion),
	1992:  uint16(sym_brace_expansion),
	1993:  uint16(sym_double_quote_string),
	1994:  uint16(sym_single_quote_string),
	1995:  uint16(7),
	1996:  uint16(6),
	1997:  uint16(sym_integer),
	1998:  uint16(sym_float),
	1999:  uint16(sym_escape_sequence),
	2000:  uint16(sym_home_dir_expansion),
	2001:  uint16(sym_glob),
	2002:  uint16(sym_word),
	2003:  uint16(263),
	2004:  uint16(12),
	2005:  uint16(sym_conditional_execution),
	2006:  uint16(sym_pipe),
	2007:  uint16(sym_redirect_statement),
	2008:  uint16(sym_negated_statement),
	2009:  uint16(sym_function_definition),
	2010:  uint16(sym_return),
	2011:  uint16(sym_switch_statement),
	2012:  uint16(sym_for_statement),
	2013:  uint16(sym_while_statement),
	2014:  uint16(sym_if_statement),
	2015:  uint16(sym_begin_statement),
	2016:  uint16(sym_command),
	2017:  uint16(26),
	2018:  uint16(3),
	2019:  uint16(1),
	2020:  uint16(sym_comment),
	2021:  uint16(15),
	2022:  uint16(1),
	2023:  uint16(anon_sym_DOLLAR),
	2024:  uint16(17),
	2025:  uint16(1),
	2026:  uint16(anon_sym_LPAREN),
	2027:  uint16(19),
	2028:  uint16(1),
	2029:  uint16(anon_sym_function),
	2030:  uint16(21),
	2031:  uint16(1),
	2032:  uint16(anon_sym_return),
	2033:  uint16(23),
	2034:  uint16(1),
	2035:  uint16(anon_sym_switch),
	2036:  uint16(27),
	2037:  uint16(1),
	2038:  uint16(anon_sym_for),
	2039:  uint16(29),
	2040:  uint16(1),
	2041:  uint16(anon_sym_while),
	2042:  uint16(31),
	2043:  uint16(1),
	2044:  uint16(anon_sym_if),
	2045:  uint16(33),
	2046:  uint16(1),
	2047:  uint16(anon_sym_begin),
	2048:  uint16(37),
	2049:  uint16(1),
	2050:  uint16(anon_sym_LBRACE),
	2051:  uint16(39),
	2052:  uint16(1),
	2053:  uint16(anon_sym_DQUOTE),
	2054:  uint16(41),
	2055:  uint16(1),
	2056:  uint16(anon_sym_SQUOTE),
	2057:  uint16(43),
	2058:  uint16(1),
	2059:  uint16(sym__begin_brace),
	2060:  uint16(403),
	2061:  uint16(1),
	2062:  uint16(anon_sym_end),
	2063:  uint16(87),
	2064:  uint16(1),
	2065:  uint16(sym__special_character),
	2066:  uint16(9),
	2067:  uint16(2),
	2068:  uint16(anon_sym_and),
	2069:  uint16(anon_sym_or),
	2070:  uint16(13),
	2071:  uint16(2),
	2072:  uint16(anon_sym_BANG),
	2073:  uint16(anon_sym_not),
	2074:  uint16(35),
	2075:  uint16(2),
	2076:  uint16(anon_sym_LBRACK),
	2077:  uint16(anon_sym_RBRACK),
	2078:  uint16(371),
	2079:  uint16(2),
	2080:  uint16(sym_break),
	2081:  uint16(sym_continue),
	2082:  uint16(43),
	2083:  uint16(2),
	2084:  uint16(sym__terminated_statement),
	2085:  uint16(aux_sym_function_definition_repeat2),
	2086:  uint16(69),
	2087:  uint16(2),
	2088:  uint16(sym_concatenation),
	2089:  uint16(sym__expression),
	2090:  uint16(103),
	2091:  uint16(2),
	2092:  uint16(sym__command_substitution_dollar),
	2093:  uint16(sym__command_substitution_inner),
	2094:  uint16(84),
	2095:  uint16(5),
	2096:  uint16(sym_command_substitution),
	2097:  uint16(sym_variable_expansion),
	2098:  uint16(sym_brace_expansion),
	2099:  uint16(sym_double_quote_string),
	2100:  uint16(sym_single_quote_string),
	2101:  uint16(7),
	2102:  uint16(6),
	2103:  uint16(sym_integer),
	2104:  uint16(sym_float),
	2105:  uint16(sym_escape_sequence),
	2106:  uint16(sym_home_dir_expansion),
	2107:  uint16(sym_glob),
	2108:  uint16(sym_word),
	2109:  uint16(263),
	2110:  uint16(12),
	2111:  uint16(sym_conditional_execution),
	2112:  uint16(sym_pipe),
	2113:  uint16(sym_redirect_statement),
	2114:  uint16(sym_negated_statement),
	2115:  uint16(sym_function_definition),
	2116:  uint16(sym_return),
	2117:  uint16(sym_switch_statement),
	2118:  uint16(sym_for_statement),
	2119:  uint16(sym_while_statement),
	2120:  uint16(sym_if_statement),
	2121:  uint16(sym_begin_statement),
	2122:  uint16(sym_command),
	2123:  uint16(26),
	2124:  uint16(3),
	2125:  uint16(1),
	2126:  uint16(sym_comment),
	2127:  uint16(15),
	2128:  uint16(1),
	2129:  uint16(anon_sym_DOLLAR),
	2130:  uint16(17),
	2131:  uint16(1),
	2132:  uint16(anon_sym_LPAREN),
	2133:  uint16(19),
	2134:  uint16(1),
	2135:  uint16(anon_sym_function),
	2136:  uint16(21),
	2137:  uint16(1),
	2138:  uint16(anon_sym_return),
	2139:  uint16(23),
	2140:  uint16(1),
	2141:  uint16(anon_sym_switch),
	2142:  uint16(27),
	2143:  uint16(1),
	2144:  uint16(anon_sym_for),
	2145:  uint16(29),
	2146:  uint16(1),
	2147:  uint16(anon_sym_while),
	2148:  uint16(31),
	2149:  uint16(1),
	2150:  uint16(anon_sym_if),
	2151:  uint16(33),
	2152:  uint16(1),
	2153:  uint16(anon_sym_begin),
	2154:  uint16(37),
	2155:  uint16(1),
	2156:  uint16(anon_sym_LBRACE),
	2157:  uint16(39),
	2158:  uint16(1),
	2159:  uint16(anon_sym_DQUOTE),
	2160:  uint16(41),
	2161:  uint16(1),
	2162:  uint16(anon_sym_SQUOTE),
	2163:  uint16(43),
	2164:  uint16(1),
	2165:  uint16(sym__begin_brace),
	2166:  uint16(405),
	2167:  uint16(1),
	2168:  uint16(anon_sym_end),
	2169:  uint16(87),
	2170:  uint16(1),
	2171:  uint16(sym__special_character),
	2172:  uint16(9),
	2173:  uint16(2),
	2174:  uint16(anon_sym_and),
	2175:  uint16(anon_sym_or),
	2176:  uint16(13),
	2177:  uint16(2),
	2178:  uint16(anon_sym_BANG),
	2179:  uint16(anon_sym_not),
	2180:  uint16(35),
	2181:  uint16(2),
	2182:  uint16(anon_sym_LBRACK),
	2183:  uint16(anon_sym_RBRACK),
	2184:  uint16(371),
	2185:  uint16(2),
	2186:  uint16(sym_break),
	2187:  uint16(sym_continue),
	2188:  uint16(59),
	2189:  uint16(2),
	2190:  uint16(sym__terminated_statement),
	2191:  uint16(aux_sym_function_definition_repeat2),
	2192:  uint16(69),
	2193:  uint16(2),
	2194:  uint16(sym_concatenation),
	2195:  uint16(sym__expression),
	2196:  uint16(103),
	2197:  uint16(2),
	2198:  uint16(sym__command_substitution_dollar),
	2199:  uint16(sym__command_substitution_inner),
	2200:  uint16(84),
	2201:  uint16(5),
	2202:  uint16(sym_command_substitution),
	2203:  uint16(sym_variable_expansion),
	2204:  uint16(sym_brace_expansion),
	2205:  uint16(sym_double_quote_string),
	2206:  uint16(sym_single_quote_string),
	2207:  uint16(7),
	2208:  uint16(6),
	2209:  uint16(sym_integer),
	2210:  uint16(sym_float),
	2211:  uint16(sym_escape_sequence),
	2212:  uint16(sym_home_dir_expansion),
	2213:  uint16(sym_glob),
	2214:  uint16(sym_word),
	2215:  uint16(263),
	2216:  uint16(12),
	2217:  uint16(sym_conditional_execution),
	2218:  uint16(sym_pipe),
	2219:  uint16(sym_redirect_statement),
	2220:  uint16(sym_negated_statement),
	2221:  uint16(sym_function_definition),
	2222:  uint16(sym_return),
	2223:  uint16(sym_switch_statement),
	2224:  uint16(sym_for_statement),
	2225:  uint16(sym_while_statement),
	2226:  uint16(sym_if_statement),
	2227:  uint16(sym_begin_statement),
	2228:  uint16(sym_command),
	2229:  uint16(26),
	2230:  uint16(3),
	2231:  uint16(1),
	2232:  uint16(sym_comment),
	2233:  uint16(15),
	2234:  uint16(1),
	2235:  uint16(anon_sym_DOLLAR),
	2236:  uint16(17),
	2237:  uint16(1),
	2238:  uint16(anon_sym_LPAREN),
	2239:  uint16(19),
	2240:  uint16(1),
	2241:  uint16(anon_sym_function),
	2242:  uint16(21),
	2243:  uint16(1),
	2244:  uint16(anon_sym_return),
	2245:  uint16(23),
	2246:  uint16(1),
	2247:  uint16(anon_sym_switch),
	2248:  uint16(27),
	2249:  uint16(1),
	2250:  uint16(anon_sym_for),
	2251:  uint16(29),
	2252:  uint16(1),
	2253:  uint16(anon_sym_while),
	2254:  uint16(31),
	2255:  uint16(1),
	2256:  uint16(anon_sym_if),
	2257:  uint16(33),
	2258:  uint16(1),
	2259:  uint16(anon_sym_begin),
	2260:  uint16(37),
	2261:  uint16(1),
	2262:  uint16(anon_sym_LBRACE),
	2263:  uint16(39),
	2264:  uint16(1),
	2265:  uint16(anon_sym_DQUOTE),
	2266:  uint16(41),
	2267:  uint16(1),
	2268:  uint16(anon_sym_SQUOTE),
	2269:  uint16(43),
	2270:  uint16(1),
	2271:  uint16(sym__begin_brace),
	2272:  uint16(407),
	2273:  uint16(1),
	2274:  uint16(anon_sym_end),
	2275:  uint16(87),
	2276:  uint16(1),
	2277:  uint16(sym__special_character),
	2278:  uint16(9),
	2279:  uint16(2),
	2280:  uint16(anon_sym_and),
	2281:  uint16(anon_sym_or),
	2282:  uint16(13),
	2283:  uint16(2),
	2284:  uint16(anon_sym_BANG),
	2285:  uint16(anon_sym_not),
	2286:  uint16(35),
	2287:  uint16(2),
	2288:  uint16(anon_sym_LBRACK),
	2289:  uint16(anon_sym_RBRACK),
	2290:  uint16(371),
	2291:  uint16(2),
	2292:  uint16(sym_break),
	2293:  uint16(sym_continue),
	2294:  uint16(48),
	2295:  uint16(2),
	2296:  uint16(sym__terminated_statement),
	2297:  uint16(aux_sym_function_definition_repeat2),
	2298:  uint16(69),
	2299:  uint16(2),
	2300:  uint16(sym_concatenation),
	2301:  uint16(sym__expression),
	2302:  uint16(103),
	2303:  uint16(2),
	2304:  uint16(sym__command_substitution_dollar),
	2305:  uint16(sym__command_substitution_inner),
	2306:  uint16(84),
	2307:  uint16(5),
	2308:  uint16(sym_command_substitution),
	2309:  uint16(sym_variable_expansion),
	2310:  uint16(sym_brace_expansion),
	2311:  uint16(sym_double_quote_string),
	2312:  uint16(sym_single_quote_string),
	2313:  uint16(7),
	2314:  uint16(6),
	2315:  uint16(sym_integer),
	2316:  uint16(sym_float),
	2317:  uint16(sym_escape_sequence),
	2318:  uint16(sym_home_dir_expansion),
	2319:  uint16(sym_glob),
	2320:  uint16(sym_word),
	2321:  uint16(263),
	2322:  uint16(12),
	2323:  uint16(sym_conditional_execution),
	2324:  uint16(sym_pipe),
	2325:  uint16(sym_redirect_statement),
	2326:  uint16(sym_negated_statement),
	2327:  uint16(sym_function_definition),
	2328:  uint16(sym_return),
	2329:  uint16(sym_switch_statement),
	2330:  uint16(sym_for_statement),
	2331:  uint16(sym_while_statement),
	2332:  uint16(sym_if_statement),
	2333:  uint16(sym_begin_statement),
	2334:  uint16(sym_command),
	2335:  uint16(26),
	2336:  uint16(3),
	2337:  uint16(1),
	2338:  uint16(sym_comment),
	2339:  uint16(15),
	2340:  uint16(1),
	2341:  uint16(anon_sym_DOLLAR),
	2342:  uint16(17),
	2343:  uint16(1),
	2344:  uint16(anon_sym_LPAREN),
	2345:  uint16(19),
	2346:  uint16(1),
	2347:  uint16(anon_sym_function),
	2348:  uint16(21),
	2349:  uint16(1),
	2350:  uint16(anon_sym_return),
	2351:  uint16(23),
	2352:  uint16(1),
	2353:  uint16(anon_sym_switch),
	2354:  uint16(27),
	2355:  uint16(1),
	2356:  uint16(anon_sym_for),
	2357:  uint16(29),
	2358:  uint16(1),
	2359:  uint16(anon_sym_while),
	2360:  uint16(31),
	2361:  uint16(1),
	2362:  uint16(anon_sym_if),
	2363:  uint16(33),
	2364:  uint16(1),
	2365:  uint16(anon_sym_begin),
	2366:  uint16(37),
	2367:  uint16(1),
	2368:  uint16(anon_sym_LBRACE),
	2369:  uint16(39),
	2370:  uint16(1),
	2371:  uint16(anon_sym_DQUOTE),
	2372:  uint16(41),
	2373:  uint16(1),
	2374:  uint16(anon_sym_SQUOTE),
	2375:  uint16(43),
	2376:  uint16(1),
	2377:  uint16(sym__begin_brace),
	2378:  uint16(409),
	2379:  uint16(1),
	2380:  uint16(anon_sym_end),
	2381:  uint16(87),
	2382:  uint16(1),
	2383:  uint16(sym__special_character),
	2384:  uint16(9),
	2385:  uint16(2),
	2386:  uint16(anon_sym_and),
	2387:  uint16(anon_sym_or),
	2388:  uint16(13),
	2389:  uint16(2),
	2390:  uint16(anon_sym_BANG),
	2391:  uint16(anon_sym_not),
	2392:  uint16(35),
	2393:  uint16(2),
	2394:  uint16(anon_sym_LBRACK),
	2395:  uint16(anon_sym_RBRACK),
	2396:  uint16(371),
	2397:  uint16(2),
	2398:  uint16(sym_break),
	2399:  uint16(sym_continue),
	2400:  uint16(50),
	2401:  uint16(2),
	2402:  uint16(sym__terminated_statement),
	2403:  uint16(aux_sym_function_definition_repeat2),
	2404:  uint16(69),
	2405:  uint16(2),
	2406:  uint16(sym_concatenation),
	2407:  uint16(sym__expression),
	2408:  uint16(103),
	2409:  uint16(2),
	2410:  uint16(sym__command_substitution_dollar),
	2411:  uint16(sym__command_substitution_inner),
	2412:  uint16(84),
	2413:  uint16(5),
	2414:  uint16(sym_command_substitution),
	2415:  uint16(sym_variable_expansion),
	2416:  uint16(sym_brace_expansion),
	2417:  uint16(sym_double_quote_string),
	2418:  uint16(sym_single_quote_string),
	2419:  uint16(7),
	2420:  uint16(6),
	2421:  uint16(sym_integer),
	2422:  uint16(sym_float),
	2423:  uint16(sym_escape_sequence),
	2424:  uint16(sym_home_dir_expansion),
	2425:  uint16(sym_glob),
	2426:  uint16(sym_word),
	2427:  uint16(263),
	2428:  uint16(12),
	2429:  uint16(sym_conditional_execution),
	2430:  uint16(sym_pipe),
	2431:  uint16(sym_redirect_statement),
	2432:  uint16(sym_negated_statement),
	2433:  uint16(sym_function_definition),
	2434:  uint16(sym_return),
	2435:  uint16(sym_switch_statement),
	2436:  uint16(sym_for_statement),
	2437:  uint16(sym_while_statement),
	2438:  uint16(sym_if_statement),
	2439:  uint16(sym_begin_statement),
	2440:  uint16(sym_command),
	2441:  uint16(26),
	2442:  uint16(3),
	2443:  uint16(1),
	2444:  uint16(sym_comment),
	2445:  uint16(322),
	2446:  uint16(1),
	2447:  uint16(anon_sym_DOLLAR),
	2448:  uint16(325),
	2449:  uint16(1),
	2450:  uint16(anon_sym_LPAREN),
	2451:  uint16(328),
	2452:  uint16(1),
	2453:  uint16(anon_sym_function),
	2454:  uint16(331),
	2455:  uint16(1),
	2456:  uint16(anon_sym_end),
	2457:  uint16(333),
	2458:  uint16(1),
	2459:  uint16(anon_sym_return),
	2460:  uint16(336),
	2461:  uint16(1),
	2462:  uint16(anon_sym_switch),
	2463:  uint16(342),
	2464:  uint16(1),
	2465:  uint16(anon_sym_for),
	2466:  uint16(345),
	2467:  uint16(1),
	2468:  uint16(anon_sym_while),
	2469:  uint16(348),
	2470:  uint16(1),
	2471:  uint16(anon_sym_if),
	2472:  uint16(351),
	2473:  uint16(1),
	2474:  uint16(anon_sym_begin),
	2475:  uint16(357),
	2476:  uint16(1),
	2477:  uint16(anon_sym_LBRACE),
	2478:  uint16(360),
	2479:  uint16(1),
	2480:  uint16(anon_sym_DQUOTE),
	2481:  uint16(363),
	2482:  uint16(1),
	2483:  uint16(anon_sym_SQUOTE),
	2484:  uint16(366),
	2485:  uint16(1),
	2486:  uint16(sym__begin_brace),
	2487:  uint16(87),
	2488:  uint16(1),
	2489:  uint16(sym__special_character),
	2490:  uint16(316),
	2491:  uint16(2),
	2492:  uint16(anon_sym_and),
	2493:  uint16(anon_sym_or),
	2494:  uint16(319),
	2495:  uint16(2),
	2496:  uint16(anon_sym_BANG),
	2497:  uint16(anon_sym_not),
	2498:  uint16(354),
	2499:  uint16(2),
	2500:  uint16(anon_sym_LBRACK),
	2501:  uint16(anon_sym_RBRACK),
	2502:  uint16(411),
	2503:  uint16(2),
	2504:  uint16(sym_break),
	2505:  uint16(sym_continue),
	2506:  uint16(57),
	2507:  uint16(2),
	2508:  uint16(sym__terminated_statement),
	2509:  uint16(aux_sym_function_definition_repeat2),
	2510:  uint16(69),
	2511:  uint16(2),
	2512:  uint16(sym_concatenation),
	2513:  uint16(sym__expression),
	2514:  uint16(103),
	2515:  uint16(2),
	2516:  uint16(sym__command_substitution_dollar),
	2517:  uint16(sym__command_substitution_inner),
	2518:  uint16(84),
	2519:  uint16(5),
	2520:  uint16(sym_command_substitution),
	2521:  uint16(sym_variable_expansion),
	2522:  uint16(sym_brace_expansion),
	2523:  uint16(sym_double_quote_string),
	2524:  uint16(sym_single_quote_string),
	2525:  uint16(313),
	2526:  uint16(6),
	2527:  uint16(sym_integer),
	2528:  uint16(sym_float),
	2529:  uint16(sym_escape_sequence),
	2530:  uint16(sym_home_dir_expansion),
	2531:  uint16(sym_glob),
	2532:  uint16(sym_word),
	2533:  uint16(263),
	2534:  uint16(12),
	2535:  uint16(sym_conditional_execution),
	2536:  uint16(sym_pipe),
	2537:  uint16(sym_redirect_statement),
	2538:  uint16(sym_negated_statement),
	2539:  uint16(sym_function_definition),
	2540:  uint16(sym_return),
	2541:  uint16(sym_switch_statement),
	2542:  uint16(sym_for_statement),
	2543:  uint16(sym_while_statement),
	2544:  uint16(sym_if_statement),
	2545:  uint16(sym_begin_statement),
	2546:  uint16(sym_command),
	2547:  uint16(26),
	2548:  uint16(3),
	2549:  uint16(1),
	2550:  uint16(sym_comment),
	2551:  uint16(15),
	2552:  uint16(1),
	2553:  uint16(anon_sym_DOLLAR),
	2554:  uint16(17),
	2555:  uint16(1),
	2556:  uint16(anon_sym_LPAREN),
	2557:  uint16(19),
	2558:  uint16(1),
	2559:  uint16(anon_sym_function),
	2560:  uint16(21),
	2561:  uint16(1),
	2562:  uint16(anon_sym_return),
	2563:  uint16(23),
	2564:  uint16(1),
	2565:  uint16(anon_sym_switch),
	2566:  uint16(27),
	2567:  uint16(1),
	2568:  uint16(anon_sym_for),
	2569:  uint16(29),
	2570:  uint16(1),
	2571:  uint16(anon_sym_while),
	2572:  uint16(31),
	2573:  uint16(1),
	2574:  uint16(anon_sym_if),
	2575:  uint16(33),
	2576:  uint16(1),
	2577:  uint16(anon_sym_begin),
	2578:  uint16(37),
	2579:  uint16(1),
	2580:  uint16(anon_sym_LBRACE),
	2581:  uint16(39),
	2582:  uint16(1),
	2583:  uint16(anon_sym_DQUOTE),
	2584:  uint16(41),
	2585:  uint16(1),
	2586:  uint16(anon_sym_SQUOTE),
	2587:  uint16(43),
	2588:  uint16(1),
	2589:  uint16(sym__begin_brace),
	2590:  uint16(414),
	2591:  uint16(1),
	2592:  uint16(anon_sym_end),
	2593:  uint16(87),
	2594:  uint16(1),
	2595:  uint16(sym__special_character),
	2596:  uint16(9),
	2597:  uint16(2),
	2598:  uint16(anon_sym_and),
	2599:  uint16(anon_sym_or),
	2600:  uint16(13),
	2601:  uint16(2),
	2602:  uint16(anon_sym_BANG),
	2603:  uint16(anon_sym_not),
	2604:  uint16(35),
	2605:  uint16(2),
	2606:  uint16(anon_sym_LBRACK),
	2607:  uint16(anon_sym_RBRACK),
	2608:  uint16(371),
	2609:  uint16(2),
	2610:  uint16(sym_break),
	2611:  uint16(sym_continue),
	2612:  uint16(44),
	2613:  uint16(2),
	2614:  uint16(sym__terminated_statement),
	2615:  uint16(aux_sym_function_definition_repeat2),
	2616:  uint16(69),
	2617:  uint16(2),
	2618:  uint16(sym_concatenation),
	2619:  uint16(sym__expression),
	2620:  uint16(103),
	2621:  uint16(2),
	2622:  uint16(sym__command_substitution_dollar),
	2623:  uint16(sym__command_substitution_inner),
	2624:  uint16(84),
	2625:  uint16(5),
	2626:  uint16(sym_command_substitution),
	2627:  uint16(sym_variable_expansion),
	2628:  uint16(sym_brace_expansion),
	2629:  uint16(sym_double_quote_string),
	2630:  uint16(sym_single_quote_string),
	2631:  uint16(7),
	2632:  uint16(6),
	2633:  uint16(sym_integer),
	2634:  uint16(sym_float),
	2635:  uint16(sym_escape_sequence),
	2636:  uint16(sym_home_dir_expansion),
	2637:  uint16(sym_glob),
	2638:  uint16(sym_word),
	2639:  uint16(263),
	2640:  uint16(12),
	2641:  uint16(sym_conditional_execution),
	2642:  uint16(sym_pipe),
	2643:  uint16(sym_redirect_statement),
	2644:  uint16(sym_negated_statement),
	2645:  uint16(sym_function_definition),
	2646:  uint16(sym_return),
	2647:  uint16(sym_switch_statement),
	2648:  uint16(sym_for_statement),
	2649:  uint16(sym_while_statement),
	2650:  uint16(sym_if_statement),
	2651:  uint16(sym_begin_statement),
	2652:  uint16(sym_command),
	2653:  uint16(26),
	2654:  uint16(3),
	2655:  uint16(1),
	2656:  uint16(sym_comment),
	2657:  uint16(15),
	2658:  uint16(1),
	2659:  uint16(anon_sym_DOLLAR),
	2660:  uint16(17),
	2661:  uint16(1),
	2662:  uint16(anon_sym_LPAREN),
	2663:  uint16(19),
	2664:  uint16(1),
	2665:  uint16(anon_sym_function),
	2666:  uint16(21),
	2667:  uint16(1),
	2668:  uint16(anon_sym_return),
	2669:  uint16(23),
	2670:  uint16(1),
	2671:  uint16(anon_sym_switch),
	2672:  uint16(27),
	2673:  uint16(1),
	2674:  uint16(anon_sym_for),
	2675:  uint16(29),
	2676:  uint16(1),
	2677:  uint16(anon_sym_while),
	2678:  uint16(31),
	2679:  uint16(1),
	2680:  uint16(anon_sym_if),
	2681:  uint16(33),
	2682:  uint16(1),
	2683:  uint16(anon_sym_begin),
	2684:  uint16(37),
	2685:  uint16(1),
	2686:  uint16(anon_sym_LBRACE),
	2687:  uint16(39),
	2688:  uint16(1),
	2689:  uint16(anon_sym_DQUOTE),
	2690:  uint16(41),
	2691:  uint16(1),
	2692:  uint16(anon_sym_SQUOTE),
	2693:  uint16(43),
	2694:  uint16(1),
	2695:  uint16(sym__begin_brace),
	2696:  uint16(416),
	2697:  uint16(1),
	2698:  uint16(anon_sym_end),
	2699:  uint16(87),
	2700:  uint16(1),
	2701:  uint16(sym__special_character),
	2702:  uint16(9),
	2703:  uint16(2),
	2704:  uint16(anon_sym_and),
	2705:  uint16(anon_sym_or),
	2706:  uint16(13),
	2707:  uint16(2),
	2708:  uint16(anon_sym_BANG),
	2709:  uint16(anon_sym_not),
	2710:  uint16(35),
	2711:  uint16(2),
	2712:  uint16(anon_sym_LBRACK),
	2713:  uint16(anon_sym_RBRACK),
	2714:  uint16(371),
	2715:  uint16(2),
	2716:  uint16(sym_break),
	2717:  uint16(sym_continue),
	2718:  uint16(57),
	2719:  uint16(2),
	2720:  uint16(sym__terminated_statement),
	2721:  uint16(aux_sym_function_definition_repeat2),
	2722:  uint16(69),
	2723:  uint16(2),
	2724:  uint16(sym_concatenation),
	2725:  uint16(sym__expression),
	2726:  uint16(103),
	2727:  uint16(2),
	2728:  uint16(sym__command_substitution_dollar),
	2729:  uint16(sym__command_substitution_inner),
	2730:  uint16(84),
	2731:  uint16(5),
	2732:  uint16(sym_command_substitution),
	2733:  uint16(sym_variable_expansion),
	2734:  uint16(sym_brace_expansion),
	2735:  uint16(sym_double_quote_string),
	2736:  uint16(sym_single_quote_string),
	2737:  uint16(7),
	2738:  uint16(6),
	2739:  uint16(sym_integer),
	2740:  uint16(sym_float),
	2741:  uint16(sym_escape_sequence),
	2742:  uint16(sym_home_dir_expansion),
	2743:  uint16(sym_glob),
	2744:  uint16(sym_word),
	2745:  uint16(263),
	2746:  uint16(12),
	2747:  uint16(sym_conditional_execution),
	2748:  uint16(sym_pipe),
	2749:  uint16(sym_redirect_statement),
	2750:  uint16(sym_negated_statement),
	2751:  uint16(sym_function_definition),
	2752:  uint16(sym_return),
	2753:  uint16(sym_switch_statement),
	2754:  uint16(sym_for_statement),
	2755:  uint16(sym_while_statement),
	2756:  uint16(sym_if_statement),
	2757:  uint16(sym_begin_statement),
	2758:  uint16(sym_command),
	2759:  uint16(25),
	2760:  uint16(3),
	2761:  uint16(1),
	2762:  uint16(sym_comment),
	2763:  uint16(15),
	2764:  uint16(1),
	2765:  uint16(anon_sym_DOLLAR),
	2766:  uint16(17),
	2767:  uint16(1),
	2768:  uint16(anon_sym_LPAREN),
	2769:  uint16(19),
	2770:  uint16(1),
	2771:  uint16(anon_sym_function),
	2772:  uint16(21),
	2773:  uint16(1),
	2774:  uint16(anon_sym_return),
	2775:  uint16(23),
	2776:  uint16(1),
	2777:  uint16(anon_sym_switch),
	2778:  uint16(27),
	2779:  uint16(1),
	2780:  uint16(anon_sym_for),
	2781:  uint16(29),
	2782:  uint16(1),
	2783:  uint16(anon_sym_while),
	2784:  uint16(31),
	2785:  uint16(1),
	2786:  uint16(anon_sym_if),
	2787:  uint16(33),
	2788:  uint16(1),
	2789:  uint16(anon_sym_begin),
	2790:  uint16(37),
	2791:  uint16(1),
	2792:  uint16(anon_sym_LBRACE),
	2793:  uint16(39),
	2794:  uint16(1),
	2795:  uint16(anon_sym_DQUOTE),
	2796:  uint16(41),
	2797:  uint16(1),
	2798:  uint16(anon_sym_SQUOTE),
	2799:  uint16(43),
	2800:  uint16(1),
	2801:  uint16(sym__begin_brace),
	2802:  uint16(6),
	2803:  uint16(1),
	2804:  uint16(sym__terminated_statement),
	2805:  uint16(87),
	2806:  uint16(1),
	2807:  uint16(sym__special_character),
	2808:  uint16(9),
	2809:  uint16(2),
	2810:  uint16(anon_sym_and),
	2811:  uint16(anon_sym_or),
	2812:  uint16(13),
	2813:  uint16(2),
	2814:  uint16(anon_sym_BANG),
	2815:  uint16(anon_sym_not),
	2816:  uint16(35),
	2817:  uint16(2),
	2818:  uint16(anon_sym_LBRACK),
	2819:  uint16(anon_sym_RBRACK),
	2820:  uint16(418),
	2821:  uint16(2),
	2822:  uint16(sym_break),
	2823:  uint16(sym_continue),
	2824:  uint16(69),
	2825:  uint16(2),
	2826:  uint16(sym_concatenation),
	2827:  uint16(sym__expression),
	2828:  uint16(103),
	2829:  uint16(2),
	2830:  uint16(sym__command_substitution_dollar),
	2831:  uint16(sym__command_substitution_inner),
	2832:  uint16(84),
	2833:  uint16(5),
	2834:  uint16(sym_command_substitution),
	2835:  uint16(sym_variable_expansion),
	2836:  uint16(sym_brace_expansion),
	2837:  uint16(sym_double_quote_string),
	2838:  uint16(sym_single_quote_string),
	2839:  uint16(7),
	2840:  uint16(6),
	2841:  uint16(sym_integer),
	2842:  uint16(sym_float),
	2843:  uint16(sym_escape_sequence),
	2844:  uint16(sym_home_dir_expansion),
	2845:  uint16(sym_glob),
	2846:  uint16(sym_word),
	2847:  uint16(268),
	2848:  uint16(12),
	2849:  uint16(sym_conditional_execution),
	2850:  uint16(sym_pipe),
	2851:  uint16(sym_redirect_statement),
	2852:  uint16(sym_negated_statement),
	2853:  uint16(sym_function_definition),
	2854:  uint16(sym_return),
	2855:  uint16(sym_switch_statement),
	2856:  uint16(sym_for_statement),
	2857:  uint16(sym_while_statement),
	2858:  uint16(sym_if_statement),
	2859:  uint16(sym_begin_statement),
	2860:  uint16(sym_command),
	2861:  uint16(25),
	2862:  uint16(3),
	2863:  uint16(1),
	2864:  uint16(sym_comment),
	2865:  uint16(15),
	2866:  uint16(1),
	2867:  uint16(anon_sym_DOLLAR),
	2868:  uint16(17),
	2869:  uint16(1),
	2870:  uint16(anon_sym_LPAREN),
	2871:  uint16(19),
	2872:  uint16(1),
	2873:  uint16(anon_sym_function),
	2874:  uint16(21),
	2875:  uint16(1),
	2876:  uint16(anon_sym_return),
	2877:  uint16(23),
	2878:  uint16(1),
	2879:  uint16(anon_sym_switch),
	2880:  uint16(27),
	2881:  uint16(1),
	2882:  uint16(anon_sym_for),
	2883:  uint16(29),
	2884:  uint16(1),
	2885:  uint16(anon_sym_while),
	2886:  uint16(31),
	2887:  uint16(1),
	2888:  uint16(anon_sym_if),
	2889:  uint16(33),
	2890:  uint16(1),
	2891:  uint16(anon_sym_begin),
	2892:  uint16(37),
	2893:  uint16(1),
	2894:  uint16(anon_sym_LBRACE),
	2895:  uint16(39),
	2896:  uint16(1),
	2897:  uint16(anon_sym_DQUOTE),
	2898:  uint16(41),
	2899:  uint16(1),
	2900:  uint16(anon_sym_SQUOTE),
	2901:  uint16(43),
	2902:  uint16(1),
	2903:  uint16(sym__begin_brace),
	2904:  uint16(12),
	2905:  uint16(1),
	2906:  uint16(sym__terminated_statement),
	2907:  uint16(87),
	2908:  uint16(1),
	2909:  uint16(sym__special_character),
	2910:  uint16(9),
	2911:  uint16(2),
	2912:  uint16(anon_sym_and),
	2913:  uint16(anon_sym_or),
	2914:  uint16(13),
	2915:  uint16(2),
	2916:  uint16(anon_sym_BANG),
	2917:  uint16(anon_sym_not),
	2918:  uint16(35),
	2919:  uint16(2),
	2920:  uint16(anon_sym_LBRACK),
	2921:  uint16(anon_sym_RBRACK),
	2922:  uint16(420),
	2923:  uint16(2),
	2924:  uint16(sym_break),
	2925:  uint16(sym_continue),
	2926:  uint16(69),
	2927:  uint16(2),
	2928:  uint16(sym_concatenation),
	2929:  uint16(sym__expression),
	2930:  uint16(103),
	2931:  uint16(2),
	2932:  uint16(sym__command_substitution_dollar),
	2933:  uint16(sym__command_substitution_inner),
	2934:  uint16(84),
	2935:  uint16(5),
	2936:  uint16(sym_command_substitution),
	2937:  uint16(sym_variable_expansion),
	2938:  uint16(sym_brace_expansion),
	2939:  uint16(sym_double_quote_string),
	2940:  uint16(sym_single_quote_string),
	2941:  uint16(7),
	2942:  uint16(6),
	2943:  uint16(sym_integer),
	2944:  uint16(sym_float),
	2945:  uint16(sym_escape_sequence),
	2946:  uint16(sym_home_dir_expansion),
	2947:  uint16(sym_glob),
	2948:  uint16(sym_word),
	2949:  uint16(266),
	2950:  uint16(12),
	2951:  uint16(sym_conditional_execution),
	2952:  uint16(sym_pipe),
	2953:  uint16(sym_redirect_statement),
	2954:  uint16(sym_negated_statement),
	2955:  uint16(sym_function_definition),
	2956:  uint16(sym_return),
	2957:  uint16(sym_switch_statement),
	2958:  uint16(sym_for_statement),
	2959:  uint16(sym_while_statement),
	2960:  uint16(sym_if_statement),
	2961:  uint16(sym_begin_statement),
	2962:  uint16(sym_command),
	2963:  uint16(25),
	2964:  uint16(3),
	2965:  uint16(1),
	2966:  uint16(sym_comment),
	2967:  uint16(15),
	2968:  uint16(1),
	2969:  uint16(anon_sym_DOLLAR),
	2970:  uint16(17),
	2971:  uint16(1),
	2972:  uint16(anon_sym_LPAREN),
	2973:  uint16(19),
	2974:  uint16(1),
	2975:  uint16(anon_sym_function),
	2976:  uint16(21),
	2977:  uint16(1),
	2978:  uint16(anon_sym_return),
	2979:  uint16(23),
	2980:  uint16(1),
	2981:  uint16(anon_sym_switch),
	2982:  uint16(27),
	2983:  uint16(1),
	2984:  uint16(anon_sym_for),
	2985:  uint16(29),
	2986:  uint16(1),
	2987:  uint16(anon_sym_while),
	2988:  uint16(31),
	2989:  uint16(1),
	2990:  uint16(anon_sym_if),
	2991:  uint16(33),
	2992:  uint16(1),
	2993:  uint16(anon_sym_begin),
	2994:  uint16(37),
	2995:  uint16(1),
	2996:  uint16(anon_sym_LBRACE),
	2997:  uint16(39),
	2998:  uint16(1),
	2999:  uint16(anon_sym_DQUOTE),
	3000:  uint16(41),
	3001:  uint16(1),
	3002:  uint16(anon_sym_SQUOTE),
	3003:  uint16(43),
	3004:  uint16(1),
	3005:  uint16(sym__begin_brace),
	3006:  uint16(3),
	3007:  uint16(1),
	3008:  uint16(sym__terminated_statement),
	3009:  uint16(87),
	3010:  uint16(1),
	3011:  uint16(sym__special_character),
	3012:  uint16(9),
	3013:  uint16(2),
	3014:  uint16(anon_sym_and),
	3015:  uint16(anon_sym_or),
	3016:  uint16(13),
	3017:  uint16(2),
	3018:  uint16(anon_sym_BANG),
	3019:  uint16(anon_sym_not),
	3020:  uint16(35),
	3021:  uint16(2),
	3022:  uint16(anon_sym_LBRACK),
	3023:  uint16(anon_sym_RBRACK),
	3024:  uint16(418),
	3025:  uint16(2),
	3026:  uint16(sym_break),
	3027:  uint16(sym_continue),
	3028:  uint16(69),
	3029:  uint16(2),
	3030:  uint16(sym_concatenation),
	3031:  uint16(sym__expression),
	3032:  uint16(103),
	3033:  uint16(2),
	3034:  uint16(sym__command_substitution_dollar),
	3035:  uint16(sym__command_substitution_inner),
	3036:  uint16(84),
	3037:  uint16(5),
	3038:  uint16(sym_command_substitution),
	3039:  uint16(sym_variable_expansion),
	3040:  uint16(sym_brace_expansion),
	3041:  uint16(sym_double_quote_string),
	3042:  uint16(sym_single_quote_string),
	3043:  uint16(7),
	3044:  uint16(6),
	3045:  uint16(sym_integer),
	3046:  uint16(sym_float),
	3047:  uint16(sym_escape_sequence),
	3048:  uint16(sym_home_dir_expansion),
	3049:  uint16(sym_glob),
	3050:  uint16(sym_word),
	3051:  uint16(268),
	3052:  uint16(12),
	3053:  uint16(sym_conditional_execution),
	3054:  uint16(sym_pipe),
	3055:  uint16(sym_redirect_statement),
	3056:  uint16(sym_negated_statement),
	3057:  uint16(sym_function_definition),
	3058:  uint16(sym_return),
	3059:  uint16(sym_switch_statement),
	3060:  uint16(sym_for_statement),
	3061:  uint16(sym_while_statement),
	3062:  uint16(sym_if_statement),
	3063:  uint16(sym_begin_statement),
	3064:  uint16(sym_command),
	3065:  uint16(24),
	3066:  uint16(3),
	3067:  uint16(1),
	3068:  uint16(sym_comment),
	3069:  uint16(15),
	3070:  uint16(1),
	3071:  uint16(anon_sym_DOLLAR),
	3072:  uint16(17),
	3073:  uint16(1),
	3074:  uint16(anon_sym_LPAREN),
	3075:  uint16(19),
	3076:  uint16(1),
	3077:  uint16(anon_sym_function),
	3078:  uint16(21),
	3079:  uint16(1),
	3080:  uint16(anon_sym_return),
	3081:  uint16(23),
	3082:  uint16(1),
	3083:  uint16(anon_sym_switch),
	3084:  uint16(27),
	3085:  uint16(1),
	3086:  uint16(anon_sym_for),
	3087:  uint16(29),
	3088:  uint16(1),
	3089:  uint16(anon_sym_while),
	3090:  uint16(31),
	3091:  uint16(1),
	3092:  uint16(anon_sym_if),
	3093:  uint16(33),
	3094:  uint16(1),
	3095:  uint16(anon_sym_begin),
	3096:  uint16(37),
	3097:  uint16(1),
	3098:  uint16(anon_sym_LBRACE),
	3099:  uint16(39),
	3100:  uint16(1),
	3101:  uint16(anon_sym_DQUOTE),
	3102:  uint16(41),
	3103:  uint16(1),
	3104:  uint16(anon_sym_SQUOTE),
	3105:  uint16(43),
	3106:  uint16(1),
	3107:  uint16(sym__begin_brace),
	3108:  uint16(87),
	3109:  uint16(1),
	3110:  uint16(sym__special_character),
	3111:  uint16(9),
	3112:  uint16(2),
	3113:  uint16(anon_sym_and),
	3114:  uint16(anon_sym_or),
	3115:  uint16(13),
	3116:  uint16(2),
	3117:  uint16(anon_sym_BANG),
	3118:  uint16(anon_sym_not),
	3119:  uint16(35),
	3120:  uint16(2),
	3121:  uint16(anon_sym_LBRACK),
	3122:  uint16(anon_sym_RBRACK),
	3123:  uint16(422),
	3124:  uint16(2),
	3125:  uint16(sym_break),
	3126:  uint16(sym_continue),
	3127:  uint16(69),
	3128:  uint16(2),
	3129:  uint16(sym_concatenation),
	3130:  uint16(sym__expression),
	3131:  uint16(103),
	3132:  uint16(2),
	3133:  uint16(sym__command_substitution_dollar),
	3134:  uint16(sym__command_substitution_inner),
	3135:  uint16(84),
	3136:  uint16(5),
	3137:  uint16(sym_command_substitution),
	3138:  uint16(sym_variable_expansion),
	3139:  uint16(sym_brace_expansion),
	3140:  uint16(sym_double_quote_string),
	3141:  uint16(sym_single_quote_string),
	3142:  uint16(7),
	3143:  uint16(6),
	3144:  uint16(sym_integer),
	3145:  uint16(sym_float),
	3146:  uint16(sym_escape_sequence),
	3147:  uint16(sym_home_dir_expansion),
	3148:  uint16(sym_glob),
	3149:  uint16(sym_word),
	3150:  uint16(202),
	3151:  uint16(12),
	3152:  uint16(sym_conditional_execution),
	3153:  uint16(sym_pipe),
	3154:  uint16(sym_redirect_statement),
	3155:  uint16(sym_negated_statement),
	3156:  uint16(sym_function_definition),
	3157:  uint16(sym_return),
	3158:  uint16(sym_switch_statement),
	3159:  uint16(sym_for_statement),
	3160:  uint16(sym_while_statement),
	3161:  uint16(sym_if_statement),
	3162:  uint16(sym_begin_statement),
	3163:  uint16(sym_command),
	3164:  uint16(24),
	3165:  uint16(3),
	3166:  uint16(1),
	3167:  uint16(sym_comment),
	3168:  uint16(15),
	3169:  uint16(1),
	3170:  uint16(anon_sym_DOLLAR),
	3171:  uint16(17),
	3172:  uint16(1),
	3173:  uint16(anon_sym_LPAREN),
	3174:  uint16(19),
	3175:  uint16(1),
	3176:  uint16(anon_sym_function),
	3177:  uint16(21),
	3178:  uint16(1),
	3179:  uint16(anon_sym_return),
	3180:  uint16(23),
	3181:  uint16(1),
	3182:  uint16(anon_sym_switch),
	3183:  uint16(27),
	3184:  uint16(1),
	3185:  uint16(anon_sym_for),
	3186:  uint16(29),
	3187:  uint16(1),
	3188:  uint16(anon_sym_while),
	3189:  uint16(31),
	3190:  uint16(1),
	3191:  uint16(anon_sym_if),
	3192:  uint16(33),
	3193:  uint16(1),
	3194:  uint16(anon_sym_begin),
	3195:  uint16(37),
	3196:  uint16(1),
	3197:  uint16(anon_sym_LBRACE),
	3198:  uint16(39),
	3199:  uint16(1),
	3200:  uint16(anon_sym_DQUOTE),
	3201:  uint16(41),
	3202:  uint16(1),
	3203:  uint16(anon_sym_SQUOTE),
	3204:  uint16(43),
	3205:  uint16(1),
	3206:  uint16(sym__begin_brace),
	3207:  uint16(87),
	3208:  uint16(1),
	3209:  uint16(sym__special_character),
	3210:  uint16(9),
	3211:  uint16(2),
	3212:  uint16(anon_sym_and),
	3213:  uint16(anon_sym_or),
	3214:  uint16(13),
	3215:  uint16(2),
	3216:  uint16(anon_sym_BANG),
	3217:  uint16(anon_sym_not),
	3218:  uint16(35),
	3219:  uint16(2),
	3220:  uint16(anon_sym_LBRACK),
	3221:  uint16(anon_sym_RBRACK),
	3222:  uint16(424),
	3223:  uint16(2),
	3224:  uint16(sym_break),
	3225:  uint16(sym_continue),
	3226:  uint16(69),
	3227:  uint16(2),
	3228:  uint16(sym_concatenation),
	3229:  uint16(sym__expression),
	3230:  uint16(103),
	3231:  uint16(2),
	3232:  uint16(sym__command_substitution_dollar),
	3233:  uint16(sym__command_substitution_inner),
	3234:  uint16(84),
	3235:  uint16(5),
	3236:  uint16(sym_command_substitution),
	3237:  uint16(sym_variable_expansion),
	3238:  uint16(sym_brace_expansion),
	3239:  uint16(sym_double_quote_string),
	3240:  uint16(sym_single_quote_string),
	3241:  uint16(7),
	3242:  uint16(6),
	3243:  uint16(sym_integer),
	3244:  uint16(sym_float),
	3245:  uint16(sym_escape_sequence),
	3246:  uint16(sym_home_dir_expansion),
	3247:  uint16(sym_glob),
	3248:  uint16(sym_word),
	3249:  uint16(191),
	3250:  uint16(12),
	3251:  uint16(sym_conditional_execution),
	3252:  uint16(sym_pipe),
	3253:  uint16(sym_redirect_statement),
	3254:  uint16(sym_negated_statement),
	3255:  uint16(sym_function_definition),
	3256:  uint16(sym_return),
	3257:  uint16(sym_switch_statement),
	3258:  uint16(sym_for_statement),
	3259:  uint16(sym_while_statement),
	3260:  uint16(sym_if_statement),
	3261:  uint16(sym_begin_statement),
	3262:  uint16(sym_command),
	3263:  uint16(24),
	3264:  uint16(3),
	3265:  uint16(1),
	3266:  uint16(sym_comment),
	3267:  uint16(15),
	3268:  uint16(1),
	3269:  uint16(anon_sym_DOLLAR),
	3270:  uint16(17),
	3271:  uint16(1),
	3272:  uint16(anon_sym_LPAREN),
	3273:  uint16(19),
	3274:  uint16(1),
	3275:  uint16(anon_sym_function),
	3276:  uint16(21),
	3277:  uint16(1),
	3278:  uint16(anon_sym_return),
	3279:  uint16(23),
	3280:  uint16(1),
	3281:  uint16(anon_sym_switch),
	3282:  uint16(27),
	3283:  uint16(1),
	3284:  uint16(anon_sym_for),
	3285:  uint16(29),
	3286:  uint16(1),
	3287:  uint16(anon_sym_while),
	3288:  uint16(31),
	3289:  uint16(1),
	3290:  uint16(anon_sym_if),
	3291:  uint16(33),
	3292:  uint16(1),
	3293:  uint16(anon_sym_begin),
	3294:  uint16(37),
	3295:  uint16(1),
	3296:  uint16(anon_sym_LBRACE),
	3297:  uint16(39),
	3298:  uint16(1),
	3299:  uint16(anon_sym_DQUOTE),
	3300:  uint16(41),
	3301:  uint16(1),
	3302:  uint16(anon_sym_SQUOTE),
	3303:  uint16(43),
	3304:  uint16(1),
	3305:  uint16(sym__begin_brace),
	3306:  uint16(87),
	3307:  uint16(1),
	3308:  uint16(sym__special_character),
	3309:  uint16(9),
	3310:  uint16(2),
	3311:  uint16(anon_sym_and),
	3312:  uint16(anon_sym_or),
	3313:  uint16(13),
	3314:  uint16(2),
	3315:  uint16(anon_sym_BANG),
	3316:  uint16(anon_sym_not),
	3317:  uint16(35),
	3318:  uint16(2),
	3319:  uint16(anon_sym_LBRACK),
	3320:  uint16(anon_sym_RBRACK),
	3321:  uint16(426),
	3322:  uint16(2),
	3323:  uint16(sym_break),
	3324:  uint16(sym_continue),
	3325:  uint16(69),
	3326:  uint16(2),
	3327:  uint16(sym_concatenation),
	3328:  uint16(sym__expression),
	3329:  uint16(103),
	3330:  uint16(2),
	3331:  uint16(sym__command_substitution_dollar),
	3332:  uint16(sym__command_substitution_inner),
	3333:  uint16(84),
	3334:  uint16(5),
	3335:  uint16(sym_command_substitution),
	3336:  uint16(sym_variable_expansion),
	3337:  uint16(sym_brace_expansion),
	3338:  uint16(sym_double_quote_string),
	3339:  uint16(sym_single_quote_string),
	3340:  uint16(7),
	3341:  uint16(6),
	3342:  uint16(sym_integer),
	3343:  uint16(sym_float),
	3344:  uint16(sym_escape_sequence),
	3345:  uint16(sym_home_dir_expansion),
	3346:  uint16(sym_glob),
	3347:  uint16(sym_word),
	3348:  uint16(190),
	3349:  uint16(12),
	3350:  uint16(sym_conditional_execution),
	3351:  uint16(sym_pipe),
	3352:  uint16(sym_redirect_statement),
	3353:  uint16(sym_negated_statement),
	3354:  uint16(sym_function_definition),
	3355:  uint16(sym_return),
	3356:  uint16(sym_switch_statement),
	3357:  uint16(sym_for_statement),
	3358:  uint16(sym_while_statement),
	3359:  uint16(sym_if_statement),
	3360:  uint16(sym_begin_statement),
	3361:  uint16(sym_command),
	3362:  uint16(24),
	3363:  uint16(3),
	3364:  uint16(1),
	3365:  uint16(sym_comment),
	3366:  uint16(15),
	3367:  uint16(1),
	3368:  uint16(anon_sym_DOLLAR),
	3369:  uint16(17),
	3370:  uint16(1),
	3371:  uint16(anon_sym_LPAREN),
	3372:  uint16(19),
	3373:  uint16(1),
	3374:  uint16(anon_sym_function),
	3375:  uint16(21),
	3376:  uint16(1),
	3377:  uint16(anon_sym_return),
	3378:  uint16(23),
	3379:  uint16(1),
	3380:  uint16(anon_sym_switch),
	3381:  uint16(27),
	3382:  uint16(1),
	3383:  uint16(anon_sym_for),
	3384:  uint16(29),
	3385:  uint16(1),
	3386:  uint16(anon_sym_while),
	3387:  uint16(31),
	3388:  uint16(1),
	3389:  uint16(anon_sym_if),
	3390:  uint16(33),
	3391:  uint16(1),
	3392:  uint16(anon_sym_begin),
	3393:  uint16(37),
	3394:  uint16(1),
	3395:  uint16(anon_sym_LBRACE),
	3396:  uint16(39),
	3397:  uint16(1),
	3398:  uint16(anon_sym_DQUOTE),
	3399:  uint16(41),
	3400:  uint16(1),
	3401:  uint16(anon_sym_SQUOTE),
	3402:  uint16(43),
	3403:  uint16(1),
	3404:  uint16(sym__begin_brace),
	3405:  uint16(87),
	3406:  uint16(1),
	3407:  uint16(sym__special_character),
	3408:  uint16(9),
	3409:  uint16(2),
	3410:  uint16(anon_sym_and),
	3411:  uint16(anon_sym_or),
	3412:  uint16(13),
	3413:  uint16(2),
	3414:  uint16(anon_sym_BANG),
	3415:  uint16(anon_sym_not),
	3416:  uint16(35),
	3417:  uint16(2),
	3418:  uint16(anon_sym_LBRACK),
	3419:  uint16(anon_sym_RBRACK),
	3420:  uint16(428),
	3421:  uint16(2),
	3422:  uint16(sym_break),
	3423:  uint16(sym_continue),
	3424:  uint16(69),
	3425:  uint16(2),
	3426:  uint16(sym_concatenation),
	3427:  uint16(sym__expression),
	3428:  uint16(103),
	3429:  uint16(2),
	3430:  uint16(sym__command_substitution_dollar),
	3431:  uint16(sym__command_substitution_inner),
	3432:  uint16(84),
	3433:  uint16(5),
	3434:  uint16(sym_command_substitution),
	3435:  uint16(sym_variable_expansion),
	3436:  uint16(sym_brace_expansion),
	3437:  uint16(sym_double_quote_string),
	3438:  uint16(sym_single_quote_string),
	3439:  uint16(7),
	3440:  uint16(6),
	3441:  uint16(sym_integer),
	3442:  uint16(sym_float),
	3443:  uint16(sym_escape_sequence),
	3444:  uint16(sym_home_dir_expansion),
	3445:  uint16(sym_glob),
	3446:  uint16(sym_word),
	3447:  uint16(196),
	3448:  uint16(12),
	3449:  uint16(sym_conditional_execution),
	3450:  uint16(sym_pipe),
	3451:  uint16(sym_redirect_statement),
	3452:  uint16(sym_negated_statement),
	3453:  uint16(sym_function_definition),
	3454:  uint16(sym_return),
	3455:  uint16(sym_switch_statement),
	3456:  uint16(sym_for_statement),
	3457:  uint16(sym_while_statement),
	3458:  uint16(sym_if_statement),
	3459:  uint16(sym_begin_statement),
	3460:  uint16(sym_command),
	3461:  uint16(17),
	3462:  uint16(3),
	3463:  uint16(1),
	3464:  uint16(sym_comment),
	3465:  uint16(435),
	3466:  uint16(1),
	3467:  uint16(anon_sym_DOLLAR),
	3468:  uint16(438),
	3469:  uint16(1),
	3470:  uint16(anon_sym_LPAREN),
	3471:  uint16(444),
	3472:  uint16(1),
	3473:  uint16(anon_sym_LBRACE),
	3474:  uint16(447),
	3475:  uint16(1),
	3476:  uint16(anon_sym_DQUOTE),
	3477:  uint16(450),
	3478:  uint16(1),
	3479:  uint16(anon_sym_SQUOTE),
	3480:  uint16(453),
	3481:  uint16(1),
	3482:  uint16(sym_stream_redirect),
	3483:  uint16(456),
	3484:  uint16(1),
	3485:  uint16(sym_direction),
	3486:  uint16(67),
	3487:  uint16(1),
	3488:  uint16(aux_sym_command_repeat1),
	3489:  uint16(87),
	3490:  uint16(1),
	3491:  uint16(sym__special_character),
	3492:  uint16(115),
	3493:  uint16(1),
	3494:  uint16(sym_file_redirect),
	3495:  uint16(441),
	3496:  uint16(2),
	3497:  uint16(anon_sym_LBRACK),
	3498:  uint16(anon_sym_RBRACK),
	3499:  uint16(103),
	3500:  uint16(2),
	3501:  uint16(sym__command_substitution_dollar),
	3502:  uint16(sym__command_substitution_inner),
	3503:  uint16(113),
	3504:  uint16(2),
	3505:  uint16(sym_concatenation),
	3506:  uint16(sym__expression),
	3507:  uint16(84),
	3508:  uint16(5),
	3509:  uint16(sym_command_substitution),
	3510:  uint16(sym_variable_expansion),
	3511:  uint16(sym_brace_expansion),
	3512:  uint16(sym_double_quote_string),
	3513:  uint16(sym_single_quote_string),
	3514:  uint16(430),
	3515:  uint16(6),
	3516:  uint16(sym_integer),
	3517:  uint16(sym_float),
	3518:  uint16(sym_escape_sequence),
	3519:  uint16(sym_home_dir_expansion),
	3520:  uint16(sym_glob),
	3521:  uint16(sym_word),
	3522:  uint16(433),
	3523:  uint16(12),
	3524:  uint16(anon_sym_PIPE_PIPE),
	3525:  uint16(anon_sym_AMP_AMP),
	3526:  uint16(anon_sym_AMP_PIPE),
	3527:  uint16(anon_sym_2_GT_PIPE),
	3528:  uint16(anon_sym_PIPE),
	3529:  uint16(anon_sym_SEMI),
	3530:  uint16(anon_sym_AMP),
	3531:  uint16(anon_sym_LF),
	3532:  uint16(anon_sym_CR),
	3533:  uint16(anon_sym_CR_LF),
	3534:  uint16(anon_sym_RPAREN),
	3535:  uint16(anon_sym_RBRACE),
	3536:  uint16(17),
	3537:  uint16(3),
	3538:  uint16(1),
	3539:  uint16(sym_comment),
	3540:  uint16(15),
	3541:  uint16(1),
	3542:  uint16(anon_sym_DOLLAR),
	3543:  uint16(17),
	3544:  uint16(1),
	3545:  uint16(anon_sym_LPAREN),
	3546:  uint16(37),
	3547:  uint16(1),
	3548:  uint16(anon_sym_LBRACE),
	3549:  uint16(39),
	3550:  uint16(1),
	3551:  uint16(anon_sym_DQUOTE),
	3552:  uint16(41),
	3553:  uint16(1),
	3554:  uint16(anon_sym_SQUOTE),
	3555:  uint16(461),
	3556:  uint16(1),
	3557:  uint16(sym_stream_redirect),
	3558:  uint16(463),
	3559:  uint16(1),
	3560:  uint16(sym_direction),
	3561:  uint16(67),
	3562:  uint16(1),
	3563:  uint16(aux_sym_command_repeat1),
	3564:  uint16(87),
	3565:  uint16(1),
	3566:  uint16(sym__special_character),
	3567:  uint16(115),
	3568:  uint16(1),
	3569:  uint16(sym_file_redirect),
	3570:  uint16(35),
	3571:  uint16(2),
	3572:  uint16(anon_sym_LBRACK),
	3573:  uint16(anon_sym_RBRACK),
	3574:  uint16(103),
	3575:  uint16(2),
	3576:  uint16(sym__command_substitution_dollar),
	3577:  uint16(sym__command_substitution_inner),
	3578:  uint16(113),
	3579:  uint16(2),
	3580:  uint16(sym_concatenation),
	3581:  uint16(sym__expression),
	3582:  uint16(84),
	3583:  uint16(5),
	3584:  uint16(sym_command_substitution),
	3585:  uint16(sym_variable_expansion),
	3586:  uint16(sym_brace_expansion),
	3587:  uint16(sym_double_quote_string),
	3588:  uint16(sym_single_quote_string),
	3589:  uint16(7),
	3590:  uint16(6),
	3591:  uint16(sym_integer),
	3592:  uint16(sym_float),
	3593:  uint16(sym_escape_sequence),
	3594:  uint16(sym_home_dir_expansion),
	3595:  uint16(sym_glob),
	3596:  uint16(sym_word),
	3597:  uint16(459),
	3598:  uint16(12),
	3599:  uint16(anon_sym_PIPE_PIPE),
	3600:  uint16(anon_sym_AMP_AMP),
	3601:  uint16(anon_sym_AMP_PIPE),
	3602:  uint16(anon_sym_2_GT_PIPE),
	3603:  uint16(anon_sym_PIPE),
	3604:  uint16(anon_sym_SEMI),
	3605:  uint16(anon_sym_AMP),
	3606:  uint16(anon_sym_LF),
	3607:  uint16(anon_sym_CR),
	3608:  uint16(anon_sym_CR_LF),
	3609:  uint16(anon_sym_RPAREN),
	3610:  uint16(anon_sym_RBRACE),
	3611:  uint16(17),
	3612:  uint16(3),
	3613:  uint16(1),
	3614:  uint16(sym_comment),
	3615:  uint16(15),
	3616:  uint16(1),
	3617:  uint16(anon_sym_DOLLAR),
	3618:  uint16(17),
	3619:  uint16(1),
	3620:  uint16(anon_sym_LPAREN),
	3621:  uint16(37),
	3622:  uint16(1),
	3623:  uint16(anon_sym_LBRACE),
	3624:  uint16(39),
	3625:  uint16(1),
	3626:  uint16(anon_sym_DQUOTE),
	3627:  uint16(41),
	3628:  uint16(1),
	3629:  uint16(anon_sym_SQUOTE),
	3630:  uint16(461),
	3631:  uint16(1),
	3632:  uint16(sym_stream_redirect),
	3633:  uint16(463),
	3634:  uint16(1),
	3635:  uint16(sym_direction),
	3636:  uint16(68),
	3637:  uint16(1),
	3638:  uint16(aux_sym_command_repeat1),
	3639:  uint16(87),
	3640:  uint16(1),
	3641:  uint16(sym__special_character),
	3642:  uint16(115),
	3643:  uint16(1),
	3644:  uint16(sym_file_redirect),
	3645:  uint16(35),
	3646:  uint16(2),
	3647:  uint16(anon_sym_LBRACK),
	3648:  uint16(anon_sym_RBRACK),
	3649:  uint16(103),
	3650:  uint16(2),
	3651:  uint16(sym__command_substitution_dollar),
	3652:  uint16(sym__command_substitution_inner),
	3653:  uint16(113),
	3654:  uint16(2),
	3655:  uint16(sym_concatenation),
	3656:  uint16(sym__expression),
	3657:  uint16(84),
	3658:  uint16(5),
	3659:  uint16(sym_command_substitution),
	3660:  uint16(sym_variable_expansion),
	3661:  uint16(sym_brace_expansion),
	3662:  uint16(sym_double_quote_string),
	3663:  uint16(sym_single_quote_string),
	3664:  uint16(7),
	3665:  uint16(6),
	3666:  uint16(sym_integer),
	3667:  uint16(sym_float),
	3668:  uint16(sym_escape_sequence),
	3669:  uint16(sym_home_dir_expansion),
	3670:  uint16(sym_glob),
	3671:  uint16(sym_word),
	3672:  uint16(465),
	3673:  uint16(12),
	3674:  uint16(anon_sym_PIPE_PIPE),
	3675:  uint16(anon_sym_AMP_AMP),
	3676:  uint16(anon_sym_AMP_PIPE),
	3677:  uint16(anon_sym_2_GT_PIPE),
	3678:  uint16(anon_sym_PIPE),
	3679:  uint16(anon_sym_SEMI),
	3680:  uint16(anon_sym_AMP),
	3681:  uint16(anon_sym_LF),
	3682:  uint16(anon_sym_CR),
	3683:  uint16(anon_sym_CR_LF),
	3684:  uint16(anon_sym_RPAREN),
	3685:  uint16(anon_sym_RBRACE),
	3686:  uint16(13),
	3687:  uint16(3),
	3688:  uint16(1),
	3689:  uint16(sym_comment),
	3690:  uint16(471),
	3691:  uint16(1),
	3692:  uint16(anon_sym_DOLLAR),
	3693:  uint16(473),
	3694:  uint16(1),
	3695:  uint16(anon_sym_LPAREN),
	3696:  uint16(477),
	3697:  uint16(1),
	3698:  uint16(anon_sym_LBRACE),
	3699:  uint16(479),
	3700:  uint16(1),
	3701:  uint16(anon_sym_DQUOTE),
	3702:  uint16(481),
	3703:  uint16(1),
	3704:  uint16(anon_sym_SQUOTE),
	3705:  uint16(178),
	3706:  uint16(1),
	3707:  uint16(sym__special_character),
	3708:  uint16(475),
	3709:  uint16(2),
	3710:  uint16(anon_sym_LBRACK),
	3711:  uint16(anon_sym_RBRACK),
	3712:  uint16(199),
	3713:  uint16(2),
	3714:  uint16(sym__command_substitution_dollar),
	3715:  uint16(sym__command_substitution_inner),
	3716:  uint16(255),
	3717:  uint16(2),
	3718:  uint16(sym_concatenation),
	3719:  uint16(sym__expression),
	3720:  uint16(182),
	3721:  uint16(5),
	3722:  uint16(sym_command_substitution),
	3723:  uint16(sym_variable_expansion),
	3724:  uint16(sym_brace_expansion),
	3725:  uint16(sym_double_quote_string),
	3726:  uint16(sym_single_quote_string),
	3727:  uint16(467),
	3728:  uint16(6),
	3729:  uint16(sym_integer),
	3730:  uint16(sym_float),
	3731:  uint16(sym_escape_sequence),
	3732:  uint16(sym_home_dir_expansion),
	3733:  uint16(sym_glob),
	3734:  uint16(sym_word),
	3735:  uint16(469),
	3736:  uint16(14),
	3737:  uint16(anon_sym_PIPE_PIPE),
	3738:  uint16(anon_sym_AMP_AMP),
	3739:  uint16(anon_sym_AMP_PIPE),
	3740:  uint16(anon_sym_2_GT_PIPE),
	3741:  uint16(anon_sym_PIPE),
	3742:  uint16(anon_sym_SEMI),
	3743:  uint16(anon_sym_AMP),
	3744:  uint16(anon_sym_LF),
	3745:  uint16(anon_sym_CR),
	3746:  uint16(anon_sym_CR_LF),
	3747:  uint16(anon_sym_RPAREN),
	3748:  uint16(anon_sym_RBRACE),
	3749:  uint16(sym_stream_redirect),
	3750:  uint16(sym_direction),
	3751:  uint16(3),
	3752:  uint16(3),
	3753:  uint16(1),
	3754:  uint16(sym_comment),
	3755:  uint16(485),
	3756:  uint16(1),
	3757:  uint16(sym__begin_brace),
	3758:  uint16(483),
	3759:  uint16(33),
	3760:  uint16(anon_sym_and),
	3761:  uint16(anon_sym_or),
	3762:  uint16(anon_sym_SEMI),
	3763:  uint16(anon_sym_AMP),
	3764:  uint16(anon_sym_LF),
	3765:  uint16(anon_sym_CR),
	3766:  uint16(anon_sym_CR_LF),
	3767:  uint16(anon_sym_BANG),
	3768:  uint16(anon_sym_not),
	3769:  uint16(anon_sym_DOLLAR),
	3770:  uint16(anon_sym_LPAREN),
	3771:  uint16(anon_sym_function),
	3772:  uint16(anon_sym_end),
	3773:  uint16(sym_integer),
	3774:  uint16(sym_float),
	3775:  uint16(anon_sym_return),
	3776:  uint16(anon_sym_switch),
	3777:  uint16(sym_break),
	3778:  uint16(sym_continue),
	3779:  uint16(anon_sym_for),
	3780:  uint16(anon_sym_while),
	3781:  uint16(anon_sym_if),
	3782:  uint16(anon_sym_else),
	3783:  uint16(anon_sym_begin),
	3784:  uint16(anon_sym_LBRACK),
	3785:  uint16(anon_sym_RBRACK),
	3786:  uint16(anon_sym_LBRACE),
	3787:  uint16(anon_sym_DQUOTE),
	3788:  uint16(anon_sym_SQUOTE),
	3789:  uint16(sym_escape_sequence),
	3790:  uint16(sym_home_dir_expansion),
	3791:  uint16(sym_glob),
	3792:  uint16(sym_word),
	3793:  uint16(3),
	3794:  uint16(3),
	3795:  uint16(1),
	3796:  uint16(sym_comment),
	3797:  uint16(489),
	3798:  uint16(1),
	3799:  uint16(sym__begin_brace),
	3800:  uint16(487),
	3801:  uint16(33),
	3802:  uint16(anon_sym_and),
	3803:  uint16(anon_sym_or),
	3804:  uint16(anon_sym_SEMI),
	3805:  uint16(anon_sym_AMP),
	3806:  uint16(anon_sym_LF),
	3807:  uint16(anon_sym_CR),
	3808:  uint16(anon_sym_CR_LF),
	3809:  uint16(anon_sym_BANG),
	3810:  uint16(anon_sym_not),
	3811:  uint16(anon_sym_DOLLAR),
	3812:  uint16(anon_sym_LPAREN),
	3813:  uint16(anon_sym_function),
	3814:  uint16(anon_sym_end),
	3815:  uint16(sym_integer),
	3816:  uint16(sym_float),
	3817:  uint16(anon_sym_return),
	3818:  uint16(anon_sym_switch),
	3819:  uint16(sym_break),
	3820:  uint16(sym_continue),
	3821:  uint16(anon_sym_for),
	3822:  uint16(anon_sym_while),
	3823:  uint16(anon_sym_if),
	3824:  uint16(anon_sym_else),
	3825:  uint16(anon_sym_begin),
	3826:  uint16(anon_sym_LBRACK),
	3827:  uint16(anon_sym_RBRACK),
	3828:  uint16(anon_sym_LBRACE),
	3829:  uint16(anon_sym_DQUOTE),
	3830:  uint16(anon_sym_SQUOTE),
	3831:  uint16(sym_escape_sequence),
	3832:  uint16(sym_home_dir_expansion),
	3833:  uint16(sym_glob),
	3834:  uint16(sym_word),
	3835:  uint16(3),
	3836:  uint16(3),
	3837:  uint16(1),
	3838:  uint16(sym_comment),
	3839:  uint16(485),
	3840:  uint16(1),
	3841:  uint16(sym__begin_brace),
	3842:  uint16(483),
	3843:  uint16(32),
	3844:  uint16(anon_sym_and),
	3845:  uint16(anon_sym_or),
	3846:  uint16(anon_sym_SEMI),
	3847:  uint16(anon_sym_AMP),
	3848:  uint16(anon_sym_LF),
	3849:  uint16(anon_sym_CR),
	3850:  uint16(anon_sym_CR_LF),
	3851:  uint16(anon_sym_BANG),
	3852:  uint16(anon_sym_not),
	3853:  uint16(anon_sym_DOLLAR),
	3854:  uint16(anon_sym_LPAREN),
	3855:  uint16(anon_sym_function),
	3856:  uint16(anon_sym_end),
	3857:  uint16(sym_integer),
	3858:  uint16(sym_float),
	3859:  uint16(anon_sym_return),
	3860:  uint16(anon_sym_switch),
	3861:  uint16(sym_break),
	3862:  uint16(sym_continue),
	3863:  uint16(anon_sym_for),
	3864:  uint16(anon_sym_while),
	3865:  uint16(anon_sym_if),
	3866:  uint16(anon_sym_begin),
	3867:  uint16(anon_sym_LBRACK),
	3868:  uint16(anon_sym_RBRACK),
	3869:  uint16(anon_sym_LBRACE),
	3870:  uint16(anon_sym_DQUOTE),
	3871:  uint16(anon_sym_SQUOTE),
	3872:  uint16(sym_escape_sequence),
	3873:  uint16(sym_home_dir_expansion),
	3874:  uint16(sym_glob),
	3875:  uint16(sym_word),
	3876:  uint16(3),
	3877:  uint16(3),
	3878:  uint16(1),
	3879:  uint16(sym_comment),
	3880:  uint16(257),
	3881:  uint16(1),
	3882:  uint16(sym__begin_brace),
	3883:  uint16(206),
	3884:  uint16(32),
	3885:  uint16(anon_sym_and),
	3886:  uint16(anon_sym_or),
	3887:  uint16(anon_sym_SEMI),
	3888:  uint16(anon_sym_AMP),
	3889:  uint16(anon_sym_LF),
	3890:  uint16(anon_sym_CR),
	3891:  uint16(anon_sym_CR_LF),
	3892:  uint16(anon_sym_BANG),
	3893:  uint16(anon_sym_not),
	3894:  uint16(anon_sym_DOLLAR),
	3895:  uint16(anon_sym_LPAREN),
	3896:  uint16(anon_sym_RPAREN),
	3897:  uint16(anon_sym_function),
	3898:  uint16(sym_integer),
	3899:  uint16(sym_float),
	3900:  uint16(anon_sym_return),
	3901:  uint16(anon_sym_switch),
	3902:  uint16(sym_break),
	3903:  uint16(sym_continue),
	3904:  uint16(anon_sym_for),
	3905:  uint16(anon_sym_while),
	3906:  uint16(anon_sym_if),
	3907:  uint16(anon_sym_begin),
	3908:  uint16(anon_sym_LBRACK),
	3909:  uint16(anon_sym_RBRACK),
	3910:  uint16(anon_sym_LBRACE),
	3911:  uint16(anon_sym_DQUOTE),
	3912:  uint16(anon_sym_SQUOTE),
	3913:  uint16(sym_escape_sequence),
	3914:  uint16(sym_home_dir_expansion),
	3915:  uint16(sym_glob),
	3916:  uint16(sym_word),
	3917:  uint16(3),
	3918:  uint16(3),
	3919:  uint16(1),
	3920:  uint16(sym_comment),
	3921:  uint16(485),
	3922:  uint16(1),
	3923:  uint16(sym__begin_brace),
	3924:  uint16(483),
	3925:  uint16(32),
	3926:  uint16(anon_sym_and),
	3927:  uint16(anon_sym_or),
	3928:  uint16(anon_sym_SEMI),
	3929:  uint16(anon_sym_AMP),
	3930:  uint16(anon_sym_LF),
	3931:  uint16(anon_sym_CR),
	3932:  uint16(anon_sym_CR_LF),
	3933:  uint16(anon_sym_BANG),
	3934:  uint16(anon_sym_not),
	3935:  uint16(anon_sym_DOLLAR),
	3936:  uint16(anon_sym_LPAREN),
	3937:  uint16(anon_sym_function),
	3938:  uint16(sym_integer),
	3939:  uint16(sym_float),
	3940:  uint16(anon_sym_return),
	3941:  uint16(anon_sym_switch),
	3942:  uint16(sym_break),
	3943:  uint16(sym_continue),
	3944:  uint16(anon_sym_for),
	3945:  uint16(anon_sym_while),
	3946:  uint16(anon_sym_if),
	3947:  uint16(anon_sym_begin),
	3948:  uint16(anon_sym_RBRACE),
	3949:  uint16(anon_sym_LBRACK),
	3950:  uint16(anon_sym_RBRACK),
	3951:  uint16(anon_sym_LBRACE),
	3952:  uint16(anon_sym_DQUOTE),
	3953:  uint16(anon_sym_SQUOTE),
	3954:  uint16(sym_escape_sequence),
	3955:  uint16(sym_home_dir_expansion),
	3956:  uint16(sym_glob),
	3957:  uint16(sym_word),
	3958:  uint16(3),
	3959:  uint16(3),
	3960:  uint16(1),
	3961:  uint16(sym_comment),
	3962:  uint16(257),
	3963:  uint16(2),
	3964:  uint16(sym__begin_brace),
	3966:  uint16(206),
	3967:  uint16(31),
	3968:  uint16(anon_sym_and),
	3969:  uint16(anon_sym_or),
	3970:  uint16(anon_sym_SEMI),
	3971:  uint16(anon_sym_AMP),
	3972:  uint16(anon_sym_LF),
	3973:  uint16(anon_sym_CR),
	3974:  uint16(anon_sym_CR_LF),
	3975:  uint16(anon_sym_BANG),
	3976:  uint16(anon_sym_not),
	3977:  uint16(anon_sym_DOLLAR),
	3978:  uint16(anon_sym_LPAREN),
	3979:  uint16(anon_sym_function),
	3980:  uint16(sym_integer),
	3981:  uint16(sym_float),
	3982:  uint16(anon_sym_return),
	3983:  uint16(anon_sym_switch),
	3984:  uint16(sym_break),
	3985:  uint16(sym_continue),
	3986:  uint16(anon_sym_for),
	3987:  uint16(anon_sym_while),
	3988:  uint16(anon_sym_if),
	3989:  uint16(anon_sym_begin),
	3990:  uint16(anon_sym_LBRACK),
	3991:  uint16(anon_sym_RBRACK),
	3992:  uint16(anon_sym_LBRACE),
	3993:  uint16(anon_sym_DQUOTE),
	3994:  uint16(anon_sym_SQUOTE),
	3995:  uint16(sym_escape_sequence),
	3996:  uint16(sym_home_dir_expansion),
	3997:  uint16(sym_glob),
	3998:  uint16(sym_word),
	3999:  uint16(3),
	4000:  uint16(3),
	4001:  uint16(1),
	4002:  uint16(sym_comment),
	4003:  uint16(489),
	4004:  uint16(1),
	4005:  uint16(sym__begin_brace),
	4006:  uint16(487),
	4007:  uint16(32),
	4008:  uint16(anon_sym_and),
	4009:  uint16(anon_sym_or),
	4010:  uint16(anon_sym_SEMI),
	4011:  uint16(anon_sym_AMP),
	4012:  uint16(anon_sym_LF),
	4013:  uint16(anon_sym_CR),
	4014:  uint16(anon_sym_CR_LF),
	4015:  uint16(anon_sym_BANG),
	4016:  uint16(anon_sym_not),
	4017:  uint16(anon_sym_DOLLAR),
	4018:  uint16(anon_sym_LPAREN),
	4019:  uint16(anon_sym_function),
	4020:  uint16(anon_sym_end),
	4021:  uint16(sym_integer),
	4022:  uint16(sym_float),
	4023:  uint16(anon_sym_return),
	4024:  uint16(anon_sym_switch),
	4025:  uint16(sym_break),
	4026:  uint16(sym_continue),
	4027:  uint16(anon_sym_for),
	4028:  uint16(anon_sym_while),
	4029:  uint16(anon_sym_if),
	4030:  uint16(anon_sym_begin),
	4031:  uint16(anon_sym_LBRACK),
	4032:  uint16(anon_sym_RBRACK),
	4033:  uint16(anon_sym_LBRACE),
	4034:  uint16(anon_sym_DQUOTE),
	4035:  uint16(anon_sym_SQUOTE),
	4036:  uint16(sym_escape_sequence),
	4037:  uint16(sym_home_dir_expansion),
	4038:  uint16(sym_glob),
	4039:  uint16(sym_word),
	4040:  uint16(5),
	4041:  uint16(3),
	4042:  uint16(1),
	4043:  uint16(sym_comment),
	4044:  uint16(493),
	4045:  uint16(1),
	4046:  uint16(sym__concat),
	4047:  uint16(495),
	4048:  uint16(1),
	4049:  uint16(sym__concat_list),
	4050:  uint16(82),
	4051:  uint16(1),
	4052:  uint16(aux_sym_variable_expansion_repeat1),
	4053:  uint16(491),
	4054:  uint16(27),
	4055:  uint16(anon_sym_PIPE_PIPE),
	4056:  uint16(anon_sym_AMP_AMP),
	4057:  uint16(anon_sym_AMP_PIPE),
	4058:  uint16(anon_sym_2_GT_PIPE),
	4059:  uint16(anon_sym_PIPE),
	4060:  uint16(anon_sym_SEMI),
	4061:  uint16(anon_sym_AMP),
	4062:  uint16(anon_sym_LF),
	4063:  uint16(anon_sym_CR),
	4064:  uint16(anon_sym_CR_LF),
	4065:  uint16(anon_sym_DOLLAR),
	4066:  uint16(anon_sym_LPAREN),
	4067:  uint16(anon_sym_RPAREN),
	4068:  uint16(sym_integer),
	4069:  uint16(sym_float),
	4070:  uint16(anon_sym_RBRACE),
	4071:  uint16(anon_sym_LBRACK),
	4072:  uint16(anon_sym_RBRACK),
	4073:  uint16(anon_sym_LBRACE),
	4074:  uint16(anon_sym_DQUOTE),
	4075:  uint16(anon_sym_SQUOTE),
	4076:  uint16(sym_escape_sequence),
	4077:  uint16(sym_stream_redirect),
	4078:  uint16(sym_direction),
	4079:  uint16(sym_home_dir_expansion),
	4080:  uint16(sym_glob),
	4081:  uint16(sym_word),
	4082:  uint16(5),
	4083:  uint16(3),
	4084:  uint16(1),
	4085:  uint16(sym_comment),
	4086:  uint16(495),
	4087:  uint16(1),
	4088:  uint16(sym__concat_list),
	4089:  uint16(499),
	4090:  uint16(1),
	4091:  uint16(sym__concat),
	4092:  uint16(78),
	4093:  uint16(1),
	4094:  uint16(aux_sym_variable_expansion_repeat1),
	4095:  uint16(497),
	4096:  uint16(27),
	4097:  uint16(anon_sym_PIPE_PIPE),
	4098:  uint16(anon_sym_AMP_AMP),
	4099:  uint16(anon_sym_AMP_PIPE),
	4100:  uint16(anon_sym_2_GT_PIPE),
	4101:  uint16(anon_sym_PIPE),
	4102:  uint16(anon_sym_SEMI),
	4103:  uint16(anon_sym_AMP),
	4104:  uint16(anon_sym_LF),
	4105:  uint16(anon_sym_CR),
	4106:  uint16(anon_sym_CR_LF),
	4107:  uint16(anon_sym_DOLLAR),
	4108:  uint16(anon_sym_LPAREN),
	4109:  uint16(anon_sym_RPAREN),
	4110:  uint16(sym_integer),
	4111:  uint16(sym_float),
	4112:  uint16(anon_sym_RBRACE),
	4113:  uint16(anon_sym_LBRACK),
	4114:  uint16(anon_sym_RBRACK),
	4115:  uint16(anon_sym_LBRACE),
	4116:  uint16(anon_sym_DQUOTE),
	4117:  uint16(anon_sym_SQUOTE),
	4118:  uint16(sym_escape_sequence),
	4119:  uint16(sym_stream_redirect),
	4120:  uint16(sym_direction),
	4121:  uint16(sym_home_dir_expansion),
	4122:  uint16(sym_glob),
	4123:  uint16(sym_word),
	4124:  uint16(4),
	4125:  uint16(3),
	4126:  uint16(1),
	4127:  uint16(sym_comment),
	4128:  uint16(82),
	4129:  uint16(1),
	4130:  uint16(aux_sym_variable_expansion_repeat1),
	4131:  uint16(493),
	4132:  uint16(2),
	4133:  uint16(sym__concat),
	4134:  uint16(sym__concat_list),
	4135:  uint16(491),
	4136:  uint16(27),
	4137:  uint16(anon_sym_PIPE_PIPE),
	4138:  uint16(anon_sym_AMP_AMP),
	4139:  uint16(anon_sym_AMP_PIPE),
	4140:  uint16(anon_sym_2_GT_PIPE),
	4141:  uint16(anon_sym_PIPE),
	4142:  uint16(anon_sym_SEMI),
	4143:  uint16(anon_sym_AMP),
	4144:  uint16(anon_sym_LF),
	4145:  uint16(anon_sym_CR),
	4146:  uint16(anon_sym_CR_LF),
	4147:  uint16(anon_sym_DOLLAR),
	4148:  uint16(anon_sym_LPAREN),
	4149:  uint16(anon_sym_RPAREN),
	4150:  uint16(sym_integer),
	4151:  uint16(sym_float),
	4152:  uint16(anon_sym_RBRACE),
	4153:  uint16(anon_sym_LBRACK),
	4154:  uint16(anon_sym_RBRACK),
	4155:  uint16(anon_sym_LBRACE),
	4156:  uint16(anon_sym_DQUOTE),
	4157:  uint16(anon_sym_SQUOTE),
	4158:  uint16(sym_escape_sequence),
	4159:  uint16(sym_stream_redirect),
	4160:  uint16(sym_direction),
	4161:  uint16(sym_home_dir_expansion),
	4162:  uint16(sym_glob),
	4163:  uint16(sym_word),
	4164:  uint16(4),
	4165:  uint16(3),
	4166:  uint16(1),
	4167:  uint16(sym_comment),
	4168:  uint16(80),
	4169:  uint16(1),
	4170:  uint16(aux_sym_variable_expansion_repeat1),
	4171:  uint16(499),
	4172:  uint16(2),
	4173:  uint16(sym__concat),
	4174:  uint16(sym__concat_list),
	4175:  uint16(497),
	4176:  uint16(27),
	4177:  uint16(anon_sym_PIPE_PIPE),
	4178:  uint16(anon_sym_AMP_AMP),
	4179:  uint16(anon_sym_AMP_PIPE),
	4180:  uint16(anon_sym_2_GT_PIPE),
	4181:  uint16(anon_sym_PIPE),
	4182:  uint16(anon_sym_SEMI),
	4183:  uint16(anon_sym_AMP),
	4184:  uint16(anon_sym_LF),
	4185:  uint16(anon_sym_CR),
	4186:  uint16(anon_sym_CR_LF),
	4187:  uint16(anon_sym_DOLLAR),
	4188:  uint16(anon_sym_LPAREN),
	4189:  uint16(anon_sym_RPAREN),
	4190:  uint16(sym_integer),
	4191:  uint16(sym_float),
	4192:  uint16(anon_sym_RBRACE),
	4193:  uint16(anon_sym_LBRACK),
	4194:  uint16(anon_sym_RBRACK),
	4195:  uint16(anon_sym_LBRACE),
	4196:  uint16(anon_sym_DQUOTE),
	4197:  uint16(anon_sym_SQUOTE),
	4198:  uint16(sym_escape_sequence),
	4199:  uint16(sym_stream_redirect),
	4200:  uint16(sym_direction),
	4201:  uint16(sym_home_dir_expansion),
	4202:  uint16(sym_glob),
	4203:  uint16(sym_word),
	4204:  uint16(5),
	4205:  uint16(3),
	4206:  uint16(1),
	4207:  uint16(sym_comment),
	4208:  uint16(503),
	4209:  uint16(1),
	4210:  uint16(sym__concat),
	4211:  uint16(505),
	4212:  uint16(1),
	4213:  uint16(sym__concat_list),
	4214:  uint16(82),
	4215:  uint16(1),
	4216:  uint16(aux_sym_variable_expansion_repeat1),
	4217:  uint16(501),
	4218:  uint16(27),
	4219:  uint16(anon_sym_PIPE_PIPE),
	4220:  uint16(anon_sym_AMP_AMP),
	4221:  uint16(anon_sym_AMP_PIPE),
	4222:  uint16(anon_sym_2_GT_PIPE),
	4223:  uint16(anon_sym_PIPE),
	4224:  uint16(anon_sym_SEMI),
	4225:  uint16(anon_sym_AMP),
	4226:  uint16(anon_sym_LF),
	4227:  uint16(anon_sym_CR),
	4228:  uint16(anon_sym_CR_LF),
	4229:  uint16(anon_sym_DOLLAR),
	4230:  uint16(anon_sym_LPAREN),
	4231:  uint16(anon_sym_RPAREN),
	4232:  uint16(sym_integer),
	4233:  uint16(sym_float),
	4234:  uint16(anon_sym_RBRACE),
	4235:  uint16(anon_sym_LBRACK),
	4236:  uint16(anon_sym_RBRACK),
	4237:  uint16(anon_sym_LBRACE),
	4238:  uint16(anon_sym_DQUOTE),
	4239:  uint16(anon_sym_SQUOTE),
	4240:  uint16(sym_escape_sequence),
	4241:  uint16(sym_stream_redirect),
	4242:  uint16(sym_direction),
	4243:  uint16(sym_home_dir_expansion),
	4244:  uint16(sym_glob),
	4245:  uint16(sym_word),
	4246:  uint16(3),
	4247:  uint16(3),
	4248:  uint16(1),
	4249:  uint16(sym_comment),
	4250:  uint16(503),
	4251:  uint16(2),
	4252:  uint16(sym__concat),
	4253:  uint16(sym__concat_list),
	4254:  uint16(501),
	4255:  uint16(27),
	4256:  uint16(anon_sym_PIPE_PIPE),
	4257:  uint16(anon_sym_AMP_AMP),
	4258:  uint16(anon_sym_AMP_PIPE),
	4259:  uint16(anon_sym_2_GT_PIPE),
	4260:  uint16(anon_sym_PIPE),
	4261:  uint16(anon_sym_SEMI),
	4262:  uint16(anon_sym_AMP),
	4263:  uint16(anon_sym_LF),
	4264:  uint16(anon_sym_CR),
	4265:  uint16(anon_sym_CR_LF),
	4266:  uint16(anon_sym_DOLLAR),
	4267:  uint16(anon_sym_LPAREN),
	4268:  uint16(anon_sym_RPAREN),
	4269:  uint16(sym_integer),
	4270:  uint16(sym_float),
	4271:  uint16(anon_sym_RBRACE),
	4272:  uint16(anon_sym_LBRACK),
	4273:  uint16(anon_sym_RBRACK),
	4274:  uint16(anon_sym_LBRACE),
	4275:  uint16(anon_sym_DQUOTE),
	4276:  uint16(anon_sym_SQUOTE),
	4277:  uint16(sym_escape_sequence),
	4278:  uint16(sym_stream_redirect),
	4279:  uint16(sym_direction),
	4280:  uint16(sym_home_dir_expansion),
	4281:  uint16(sym_glob),
	4282:  uint16(sym_word),
	4283:  uint16(4),
	4284:  uint16(3),
	4285:  uint16(1),
	4286:  uint16(sym_comment),
	4287:  uint16(510),
	4288:  uint16(1),
	4289:  uint16(sym__concat),
	4290:  uint16(94),
	4291:  uint16(1),
	4292:  uint16(aux_sym_concatenation_repeat1),
	4293:  uint16(508),
	4294:  uint16(27),
	4295:  uint16(anon_sym_PIPE_PIPE),
	4296:  uint16(anon_sym_AMP_AMP),
	4297:  uint16(anon_sym_AMP_PIPE),
	4298:  uint16(anon_sym_2_GT_PIPE),
	4299:  uint16(anon_sym_PIPE),
	4300:  uint16(anon_sym_SEMI),
	4301:  uint16(anon_sym_AMP),
	4302:  uint16(anon_sym_LF),
	4303:  uint16(anon_sym_CR),
	4304:  uint16(anon_sym_CR_LF),
	4305:  uint16(anon_sym_DOLLAR),
	4306:  uint16(anon_sym_LPAREN),
	4307:  uint16(anon_sym_RPAREN),
	4308:  uint16(sym_integer),
	4309:  uint16(sym_float),
	4310:  uint16(anon_sym_RBRACE),
	4311:  uint16(anon_sym_LBRACK),
	4312:  uint16(anon_sym_RBRACK),
	4313:  uint16(anon_sym_LBRACE),
	4314:  uint16(anon_sym_DQUOTE),
	4315:  uint16(anon_sym_SQUOTE),
	4316:  uint16(sym_escape_sequence),
	4317:  uint16(sym_stream_redirect),
	4318:  uint16(sym_direction),
	4319:  uint16(sym_home_dir_expansion),
	4320:  uint16(sym_glob),
	4321:  uint16(sym_word),
	4322:  uint16(13),
	4323:  uint16(3),
	4324:  uint16(1),
	4325:  uint16(sym_comment),
	4326:  uint16(516),
	4327:  uint16(1),
	4328:  uint16(anon_sym_DOLLAR),
	4329:  uint16(518),
	4330:  uint16(1),
	4331:  uint16(anon_sym_LPAREN),
	4332:  uint16(522),
	4333:  uint16(1),
	4334:  uint16(anon_sym_LBRACE),
	4335:  uint16(524),
	4336:  uint16(1),
	4337:  uint16(anon_sym_DQUOTE),
	4338:  uint16(526),
	4339:  uint16(1),
	4340:  uint16(anon_sym_SQUOTE),
	4341:  uint16(142),
	4342:  uint16(1),
	4343:  uint16(sym__special_character),
	4344:  uint16(520),
	4345:  uint16(2),
	4346:  uint16(anon_sym_LBRACK),
	4347:  uint16(anon_sym_RBRACK),
	4348:  uint16(157),
	4349:  uint16(2),
	4350:  uint16(sym__command_substitution_dollar),
	4351:  uint16(sym__command_substitution_inner),
	4352:  uint16(93),
	4353:  uint16(3),
	4354:  uint16(sym_concatenation),
	4355:  uint16(sym__expression),
	4356:  uint16(aux_sym_case_clause_repeat1),
	4357:  uint16(514),
	4358:  uint16(5),
	4359:  uint16(anon_sym_SEMI),
	4360:  uint16(anon_sym_AMP),
	4361:  uint16(anon_sym_LF),
	4362:  uint16(anon_sym_CR),
	4363:  uint16(anon_sym_CR_LF),
	4364:  uint16(137),
	4365:  uint16(5),
	4366:  uint16(sym_command_substitution),
	4367:  uint16(sym_variable_expansion),
	4368:  uint16(sym_brace_expansion),
	4369:  uint16(sym_double_quote_string),
	4370:  uint16(sym_single_quote_string),
	4371:  uint16(512),
	4372:  uint16(6),
	4373:  uint16(sym_integer),
	4374:  uint16(sym_float),
	4375:  uint16(sym_escape_sequence),
	4376:  uint16(sym_home_dir_expansion),
	4377:  uint16(sym_glob),
	4378:  uint16(sym_word),
	4379:  uint16(18),
	4380:  uint16(3),
	4381:  uint16(1),
	4382:  uint16(sym_comment),
	4383:  uint16(516),
	4384:  uint16(1),
	4385:  uint16(anon_sym_DOLLAR),
	4386:  uint16(518),
	4387:  uint16(1),
	4388:  uint16(anon_sym_LPAREN),
	4389:  uint16(522),
	4390:  uint16(1),
	4391:  uint16(anon_sym_LBRACE),
	4392:  uint16(524),
	4393:  uint16(1),
	4394:  uint16(anon_sym_DQUOTE),
	4395:  uint16(526),
	4396:  uint16(1),
	4397:  uint16(anon_sym_SQUOTE),
	4398:  uint16(528),
	4399:  uint16(1),
	4400:  uint16(anon_sym_SEMI),
	4401:  uint16(530),
	4402:  uint16(1),
	4403:  uint16(anon_sym_AMP),
	4404:  uint16(532),
	4405:  uint16(1),
	4406:  uint16(anon_sym_LF),
	4407:  uint16(534),
	4408:  uint16(1),
	4409:  uint16(anon_sym_CR),
	4410:  uint16(536),
	4411:  uint16(1),
	4412:  uint16(anon_sym_CR_LF),
	4413:  uint16(96),
	4414:  uint16(1),
	4415:  uint16(aux_sym_function_definition_repeat1),
	4416:  uint16(142),
	4417:  uint16(1),
	4418:  uint16(sym__special_character),
	4419:  uint16(520),
	4420:  uint16(2),
	4421:  uint16(anon_sym_LBRACK),
	4422:  uint16(anon_sym_RBRACK),
	4423:  uint16(157),
	4424:  uint16(2),
	4425:  uint16(sym__command_substitution_dollar),
	4426:  uint16(sym__command_substitution_inner),
	4427:  uint16(160),
	4428:  uint16(2),
	4429:  uint16(sym_concatenation),
	4430:  uint16(sym__expression),
	4431:  uint16(137),
	4432:  uint16(5),
	4433:  uint16(sym_command_substitution),
	4434:  uint16(sym_variable_expansion),
	4435:  uint16(sym_brace_expansion),
	4436:  uint16(sym_double_quote_string),
	4437:  uint16(sym_single_quote_string),
	4438:  uint16(512),
	4439:  uint16(6),
	4440:  uint16(sym_integer),
	4441:  uint16(sym_float),
	4442:  uint16(sym_escape_sequence),
	4443:  uint16(sym_home_dir_expansion),
	4444:  uint16(sym_glob),
	4445:  uint16(sym_word),
	4446:  uint16(4),
	4447:  uint16(3),
	4448:  uint16(1),
	4449:  uint16(sym_comment),
	4450:  uint16(510),
	4451:  uint16(1),
	4452:  uint16(sym__concat),
	4453:  uint16(94),
	4454:  uint16(1),
	4455:  uint16(aux_sym_concatenation_repeat1),
	4456:  uint16(538),
	4457:  uint16(27),
	4458:  uint16(anon_sym_PIPE_PIPE),
	4459:  uint16(anon_sym_AMP_AMP),
	4460:  uint16(anon_sym_AMP_PIPE),
	4461:  uint16(anon_sym_2_GT_PIPE),
	4462:  uint16(anon_sym_PIPE),
	4463:  uint16(anon_sym_SEMI),
	4464:  uint16(anon_sym_AMP),
	4465:  uint16(anon_sym_LF),
	4466:  uint16(anon_sym_CR),
	4467:  uint16(anon_sym_CR_LF),
	4468:  uint16(anon_sym_DOLLAR),
	4469:  uint16(anon_sym_LPAREN),
	4470:  uint16(anon_sym_RPAREN),
	4471:  uint16(sym_integer),
	4472:  uint16(sym_float),
	4473:  uint16(anon_sym_RBRACE),
	4474:  uint16(anon_sym_LBRACK),
	4475:  uint16(anon_sym_RBRACK),
	4476:  uint16(anon_sym_LBRACE),
	4477:  uint16(anon_sym_DQUOTE),
	4478:  uint16(anon_sym_SQUOTE),
	4479:  uint16(sym_escape_sequence),
	4480:  uint16(sym_stream_redirect),
	4481:  uint16(sym_direction),
	4482:  uint16(sym_home_dir_expansion),
	4483:  uint16(sym_glob),
	4484:  uint16(sym_word),
	4485:  uint16(3),
	4486:  uint16(3),
	4487:  uint16(1),
	4488:  uint16(sym_comment),
	4489:  uint16(542),
	4490:  uint16(2),
	4491:  uint16(sym__concat),
	4492:  uint16(sym__concat_list),
	4493:  uint16(540),
	4494:  uint16(27),
	4495:  uint16(anon_sym_PIPE_PIPE),
	4496:  uint16(anon_sym_AMP_AMP),
	4497:  uint16(anon_sym_AMP_PIPE),
	4498:  uint16(anon_sym_2_GT_PIPE),
	4499:  uint16(anon_sym_PIPE),
	4500:  uint16(anon_sym_SEMI),
	4501:  uint16(anon_sym_AMP),
	4502:  uint16(anon_sym_LF),
	4503:  uint16(anon_sym_CR),
	4504:  uint16(anon_sym_CR_LF),
	4505:  uint16(anon_sym_DOLLAR),
	4506:  uint16(anon_sym_LPAREN),
	4507:  uint16(anon_sym_RPAREN),
	4508:  uint16(sym_integer),
	4509:  uint16(sym_float),
	4510:  uint16(anon_sym_RBRACE),
	4511:  uint16(anon_sym_LBRACK),
	4512:  uint16(anon_sym_RBRACK),
	4513:  uint16(anon_sym_LBRACE),
	4514:  uint16(anon_sym_DQUOTE),
	4515:  uint16(anon_sym_SQUOTE),
	4516:  uint16(sym_escape_sequence),
	4517:  uint16(sym_stream_redirect),
	4518:  uint16(sym_direction),
	4519:  uint16(sym_home_dir_expansion),
	4520:  uint16(sym_glob),
	4521:  uint16(sym_word),
	4522:  uint16(4),
	4523:  uint16(3),
	4524:  uint16(1),
	4525:  uint16(sym_comment),
	4526:  uint16(546),
	4527:  uint16(1),
	4528:  uint16(sym__concat),
	4529:  uint16(89),
	4530:  uint16(1),
	4531:  uint16(aux_sym_concatenation_repeat1),
	4532:  uint16(544),
	4533:  uint16(27),
	4534:  uint16(anon_sym_PIPE_PIPE),
	4535:  uint16(anon_sym_AMP_AMP),
	4536:  uint16(anon_sym_AMP_PIPE),
	4537:  uint16(anon_sym_2_GT_PIPE),
	4538:  uint16(anon_sym_PIPE),
	4539:  uint16(anon_sym_SEMI),
	4540:  uint16(anon_sym_AMP),
	4541:  uint16(anon_sym_LF),
	4542:  uint16(anon_sym_CR),
	4543:  uint16(anon_sym_CR_LF),
	4544:  uint16(anon_sym_DOLLAR),
	4545:  uint16(anon_sym_LPAREN),
	4546:  uint16(anon_sym_RPAREN),
	4547:  uint16(sym_integer),
	4548:  uint16(sym_float),
	4549:  uint16(anon_sym_RBRACE),
	4550:  uint16(anon_sym_LBRACK),
	4551:  uint16(anon_sym_RBRACK),
	4552:  uint16(anon_sym_LBRACE),
	4553:  uint16(anon_sym_DQUOTE),
	4554:  uint16(anon_sym_SQUOTE),
	4555:  uint16(sym_escape_sequence),
	4556:  uint16(sym_stream_redirect),
	4557:  uint16(sym_direction),
	4558:  uint16(sym_home_dir_expansion),
	4559:  uint16(sym_glob),
	4560:  uint16(sym_word),
	4561:  uint16(14),
	4562:  uint16(3),
	4563:  uint16(1),
	4564:  uint16(sym_comment),
	4565:  uint16(554),
	4566:  uint16(1),
	4567:  uint16(anon_sym_DOLLAR),
	4568:  uint16(557),
	4569:  uint16(1),
	4570:  uint16(anon_sym_LPAREN),
	4571:  uint16(563),
	4572:  uint16(1),
	4573:  uint16(anon_sym_LBRACE),
	4574:  uint16(566),
	4575:  uint16(1),
	4576:  uint16(anon_sym_DQUOTE),
	4577:  uint16(569),
	4578:  uint16(1),
	4579:  uint16(anon_sym_SQUOTE),
	4580:  uint16(90),
	4581:  uint16(1),
	4582:  uint16(aux_sym_for_statement_repeat1),
	4583:  uint16(142),
	4584:  uint16(1),
	4585:  uint16(sym__special_character),
	4586:  uint16(560),
	4587:  uint16(2),
	4588:  uint16(anon_sym_LBRACK),
	4589:  uint16(anon_sym_RBRACK),
	4590:  uint16(157),
	4591:  uint16(2),
	4592:  uint16(sym__command_substitution_dollar),
	4593:  uint16(sym__command_substitution_inner),
	4594:  uint16(158),
	4595:  uint16(2),
	4596:  uint16(sym_concatenation),
	4597:  uint16(sym__expression),
	4598:  uint16(552),
	4599:  uint16(5),
	4600:  uint16(anon_sym_SEMI),
	4601:  uint16(anon_sym_AMP),
	4602:  uint16(anon_sym_LF),
	4603:  uint16(anon_sym_CR),
	4604:  uint16(anon_sym_CR_LF),
	4605:  uint16(137),
	4606:  uint16(5),
	4607:  uint16(sym_command_substitution),
	4608:  uint16(sym_variable_expansion),
	4609:  uint16(sym_brace_expansion),
	4610:  uint16(sym_double_quote_string),
	4611:  uint16(sym_single_quote_string),
	4612:  uint16(549),
	4613:  uint16(6),
	4614:  uint16(sym_integer),
	4615:  uint16(sym_float),
	4616:  uint16(sym_escape_sequence),
	4617:  uint16(sym_home_dir_expansion),
	4618:  uint16(sym_glob),
	4619:  uint16(sym_word),
	4620:  uint16(3),
	4621:  uint16(3),
	4622:  uint16(1),
	4623:  uint16(sym_comment),
	4624:  uint16(574),
	4625:  uint16(2),
	4626:  uint16(sym__concat),
	4627:  uint16(sym__concat_list),
	4628:  uint16(572),
	4629:  uint16(27),
	4630:  uint16(anon_sym_PIPE_PIPE),
	4631:  uint16(anon_sym_AMP_AMP),
	4632:  uint16(anon_sym_AMP_PIPE),
	4633:  uint16(anon_sym_2_GT_PIPE),
	4634:  uint16(anon_sym_PIPE),
	4635:  uint16(anon_sym_SEMI),
	4636:  uint16(anon_sym_AMP),
	4637:  uint16(anon_sym_LF),
	4638:  uint16(anon_sym_CR),
	4639:  uint16(anon_sym_CR_LF),
	4640:  uint16(anon_sym_DOLLAR),
	4641:  uint16(anon_sym_LPAREN),
	4642:  uint16(anon_sym_RPAREN),
	4643:  uint16(sym_integer),
	4644:  uint16(sym_float),
	4645:  uint16(anon_sym_RBRACE),
	4646:  uint16(anon_sym_LBRACK),
	4647:  uint16(anon_sym_RBRACK),
	4648:  uint16(anon_sym_LBRACE),
	4649:  uint16(anon_sym_DQUOTE),
	4650:  uint16(anon_sym_SQUOTE),
	4651:  uint16(sym_escape_sequence),
	4652:  uint16(sym_stream_redirect),
	4653:  uint16(sym_direction),
	4654:  uint16(sym_home_dir_expansion),
	4655:  uint16(sym_glob),
	4656:  uint16(sym_word),
	4657:  uint16(14),
	4658:  uint16(3),
	4659:  uint16(1),
	4660:  uint16(sym_comment),
	4661:  uint16(516),
	4662:  uint16(1),
	4663:  uint16(anon_sym_DOLLAR),
	4664:  uint16(518),
	4665:  uint16(1),
	4666:  uint16(anon_sym_LPAREN),
	4667:  uint16(522),
	4668:  uint16(1),
	4669:  uint16(anon_sym_LBRACE),
	4670:  uint16(524),
	4671:  uint16(1),
	4672:  uint16(anon_sym_DQUOTE),
	4673:  uint16(526),
	4674:  uint16(1),
	4675:  uint16(anon_sym_SQUOTE),
	4676:  uint16(86),
	4677:  uint16(1),
	4678:  uint16(aux_sym_function_definition_repeat1),
	4679:  uint16(142),
	4680:  uint16(1),
	4681:  uint16(sym__special_character),
	4682:  uint16(520),
	4683:  uint16(2),
	4684:  uint16(anon_sym_LBRACK),
	4685:  uint16(anon_sym_RBRACK),
	4686:  uint16(157),
	4687:  uint16(2),
	4688:  uint16(sym__command_substitution_dollar),
	4689:  uint16(sym__command_substitution_inner),
	4690:  uint16(160),
	4691:  uint16(2),
	4692:  uint16(sym_concatenation),
	4693:  uint16(sym__expression),
	4694:  uint16(576),
	4695:  uint16(5),
	4696:  uint16(anon_sym_SEMI),
	4697:  uint16(anon_sym_AMP),
	4698:  uint16(anon_sym_LF),
	4699:  uint16(anon_sym_CR),
	4700:  uint16(anon_sym_CR_LF),
	4701:  uint16(137),
	4702:  uint16(5),
	4703:  uint16(sym_command_substitution),
	4704:  uint16(sym_variable_expansion),
	4705:  uint16(sym_brace_expansion),
	4706:  uint16(sym_double_quote_string),
	4707:  uint16(sym_single_quote_string),
	4708:  uint16(512),
	4709:  uint16(6),
	4710:  uint16(sym_integer),
	4711:  uint16(sym_float),
	4712:  uint16(sym_escape_sequence),
	4713:  uint16(sym_home_dir_expansion),
	4714:  uint16(sym_glob),
	4715:  uint16(sym_word),
	4716:  uint16(13),
	4717:  uint16(3),
	4718:  uint16(1),
	4719:  uint16(sym_comment),
	4720:  uint16(583),
	4721:  uint16(1),
	4722:  uint16(anon_sym_DOLLAR),
	4723:  uint16(586),
	4724:  uint16(1),
	4725:  uint16(anon_sym_LPAREN),
	4726:  uint16(592),
	4727:  uint16(1),
	4728:  uint16(anon_sym_LBRACE),
	4729:  uint16(595),
	4730:  uint16(1),
	4731:  uint16(anon_sym_DQUOTE),
	4732:  uint16(598),
	4733:  uint16(1),
	4734:  uint16(anon_sym_SQUOTE),
	4735:  uint16(142),
	4736:  uint16(1),
	4737:  uint16(sym__special_character),
	4738:  uint16(589),
	4739:  uint16(2),
	4740:  uint16(anon_sym_LBRACK),
	4741:  uint16(anon_sym_RBRACK),
	4742:  uint16(157),
	4743:  uint16(2),
	4744:  uint16(sym__command_substitution_dollar),
	4745:  uint16(sym__command_substitution_inner),
	4746:  uint16(93),
	4747:  uint16(3),
	4748:  uint16(sym_concatenation),
	4749:  uint16(sym__expression),
	4750:  uint16(aux_sym_case_clause_repeat1),
	4751:  uint16(581),
	4752:  uint16(5),
	4753:  uint16(anon_sym_SEMI),
	4754:  uint16(anon_sym_AMP),
	4755:  uint16(anon_sym_LF),
	4756:  uint16(anon_sym_CR),
	4757:  uint16(anon_sym_CR_LF),
	4758:  uint16(137),
	4759:  uint16(5),
	4760:  uint16(sym_command_substitution),
	4761:  uint16(sym_variable_expansion),
	4762:  uint16(sym_brace_expansion),
	4763:  uint16(sym_double_quote_string),
	4764:  uint16(sym_single_quote_string),
	4765:  uint16(578),
	4766:  uint16(6),
	4767:  uint16(sym_integer),
	4768:  uint16(sym_float),
	4769:  uint16(sym_escape_sequence),
	4770:  uint16(sym_home_dir_expansion),
	4771:  uint16(sym_glob),
	4772:  uint16(sym_word),
	4773:  uint16(4),
	4774:  uint16(3),
	4775:  uint16(1),
	4776:  uint16(sym_comment),
	4777:  uint16(510),
	4778:  uint16(1),
	4779:  uint16(sym__concat),
	4780:  uint16(89),
	4781:  uint16(1),
	4782:  uint16(aux_sym_concatenation_repeat1),
	4783:  uint16(601),
	4784:  uint16(27),
	4785:  uint16(anon_sym_PIPE_PIPE),
	4786:  uint16(anon_sym_AMP_AMP),
	4787:  uint16(anon_sym_AMP_PIPE),
	4788:  uint16(anon_sym_2_GT_PIPE),
	4789:  uint16(anon_sym_PIPE),
	4790:  uint16(anon_sym_SEMI),
	4791:  uint16(anon_sym_AMP),
	4792:  uint16(anon_sym_LF),
	4793:  uint16(anon_sym_CR),
	4794:  uint16(anon_sym_CR_LF),
	4795:  uint16(anon_sym_DOLLAR),
	4796:  uint16(anon_sym_LPAREN),
	4797:  uint16(anon_sym_RPAREN),
	4798:  uint16(sym_integer),
	4799:  uint16(sym_float),
	4800:  uint16(anon_sym_RBRACE),
	4801:  uint16(anon_sym_LBRACK),
	4802:  uint16(anon_sym_RBRACK),
	4803:  uint16(anon_sym_LBRACE),
	4804:  uint16(anon_sym_DQUOTE),
	4805:  uint16(anon_sym_SQUOTE),
	4806:  uint16(sym_escape_sequence),
	4807:  uint16(sym_stream_redirect),
	4808:  uint16(sym_direction),
	4809:  uint16(sym_home_dir_expansion),
	4810:  uint16(sym_glob),
	4811:  uint16(sym_word),
	4812:  uint16(18),
	4813:  uint16(3),
	4814:  uint16(1),
	4815:  uint16(sym_comment),
	4816:  uint16(516),
	4817:  uint16(1),
	4818:  uint16(anon_sym_DOLLAR),
	4819:  uint16(518),
	4820:  uint16(1),
	4821:  uint16(anon_sym_LPAREN),
	4822:  uint16(522),
	4823:  uint16(1),
	4824:  uint16(anon_sym_LBRACE),
	4825:  uint16(524),
	4826:  uint16(1),
	4827:  uint16(anon_sym_DQUOTE),
	4828:  uint16(526),
	4829:  uint16(1),
	4830:  uint16(anon_sym_SQUOTE),
	4831:  uint16(603),
	4832:  uint16(1),
	4833:  uint16(anon_sym_SEMI),
	4834:  uint16(605),
	4835:  uint16(1),
	4836:  uint16(anon_sym_AMP),
	4837:  uint16(607),
	4838:  uint16(1),
	4839:  uint16(anon_sym_LF),
	4840:  uint16(609),
	4841:  uint16(1),
	4842:  uint16(anon_sym_CR),
	4843:  uint16(611),
	4844:  uint16(1),
	4845:  uint16(anon_sym_CR_LF),
	4846:  uint16(90),
	4847:  uint16(1),
	4848:  uint16(aux_sym_for_statement_repeat1),
	4849:  uint16(142),
	4850:  uint16(1),
	4851:  uint16(sym__special_character),
	4852:  uint16(520),
	4853:  uint16(2),
	4854:  uint16(anon_sym_LBRACK),
	4855:  uint16(anon_sym_RBRACK),
	4856:  uint16(157),
	4857:  uint16(2),
	4858:  uint16(sym__command_substitution_dollar),
	4859:  uint16(sym__command_substitution_inner),
	4860:  uint16(158),
	4861:  uint16(2),
	4862:  uint16(sym_concatenation),
	4863:  uint16(sym__expression),
	4864:  uint16(137),
	4865:  uint16(5),
	4866:  uint16(sym_command_substitution),
	4867:  uint16(sym_variable_expansion),
	4868:  uint16(sym_brace_expansion),
	4869:  uint16(sym_double_quote_string),
	4870:  uint16(sym_single_quote_string),
	4871:  uint16(512),
	4872:  uint16(6),
	4873:  uint16(sym_integer),
	4874:  uint16(sym_float),
	4875:  uint16(sym_escape_sequence),
	4876:  uint16(sym_home_dir_expansion),
	4877:  uint16(sym_glob),
	4878:  uint16(sym_word),
	4879:  uint16(14),
	4880:  uint16(3),
	4881:  uint16(1),
	4882:  uint16(sym_comment),
	4883:  uint16(618),
	4884:  uint16(1),
	4885:  uint16(anon_sym_DOLLAR),
	4886:  uint16(621),
	4887:  uint16(1),
	4888:  uint16(anon_sym_LPAREN),
	4889:  uint16(627),
	4890:  uint16(1),
	4891:  uint16(anon_sym_LBRACE),
	4892:  uint16(630),
	4893:  uint16(1),
	4894:  uint16(anon_sym_DQUOTE),
	4895:  uint16(633),
	4896:  uint16(1),
	4897:  uint16(anon_sym_SQUOTE),
	4898:  uint16(96),
	4899:  uint16(1),
	4900:  uint16(aux_sym_function_definition_repeat1),
	4901:  uint16(142),
	4902:  uint16(1),
	4903:  uint16(sym__special_character),
	4904:  uint16(624),
	4905:  uint16(2),
	4906:  uint16(anon_sym_LBRACK),
	4907:  uint16(anon_sym_RBRACK),
	4908:  uint16(157),
	4909:  uint16(2),
	4910:  uint16(sym__command_substitution_dollar),
	4911:  uint16(sym__command_substitution_inner),
	4912:  uint16(160),
	4913:  uint16(2),
	4914:  uint16(sym_concatenation),
	4915:  uint16(sym__expression),
	4916:  uint16(616),
	4917:  uint16(5),
	4918:  uint16(anon_sym_SEMI),
	4919:  uint16(anon_sym_AMP),
	4920:  uint16(anon_sym_LF),
	4921:  uint16(anon_sym_CR),
	4922:  uint16(anon_sym_CR_LF),
	4923:  uint16(137),
	4924:  uint16(5),
	4925:  uint16(sym_command_substitution),
	4926:  uint16(sym_variable_expansion),
	4927:  uint16(sym_brace_expansion),
	4928:  uint16(sym_double_quote_string),
	4929:  uint16(sym_single_quote_string),
	4930:  uint16(613),
	4931:  uint16(6),
	4932:  uint16(sym_integer),
	4933:  uint16(sym_float),
	4934:  uint16(sym_escape_sequence),
	4935:  uint16(sym_home_dir_expansion),
	4936:  uint16(sym_glob),
	4937:  uint16(sym_word),
	4938:  uint16(3),
	4939:  uint16(3),
	4940:  uint16(1),
	4941:  uint16(sym_comment),
	4942:  uint16(489),
	4943:  uint16(1),
	4944:  uint16(sym__begin_brace),
	4945:  uint16(487),
	4946:  uint16(28),
	4947:  uint16(anon_sym_and),
	4948:  uint16(anon_sym_or),
	4949:  uint16(anon_sym_BANG),
	4950:  uint16(anon_sym_not),
	4951:  uint16(anon_sym_DOLLAR),
	4952:  uint16(anon_sym_LPAREN),
	4953:  uint16(anon_sym_function),
	4954:  uint16(anon_sym_end),
	4955:  uint16(sym_integer),
	4956:  uint16(sym_float),
	4957:  uint16(anon_sym_return),
	4958:  uint16(anon_sym_switch),
	4959:  uint16(anon_sym_case),
	4960:  uint16(sym_break),
	4961:  uint16(sym_continue),
	4962:  uint16(anon_sym_for),
	4963:  uint16(anon_sym_while),
	4964:  uint16(anon_sym_if),
	4965:  uint16(anon_sym_begin),
	4966:  uint16(anon_sym_LBRACK),
	4967:  uint16(anon_sym_RBRACK),
	4968:  uint16(anon_sym_LBRACE),
	4969:  uint16(anon_sym_DQUOTE),
	4970:  uint16(anon_sym_SQUOTE),
	4971:  uint16(sym_escape_sequence),
	4972:  uint16(sym_home_dir_expansion),
	4973:  uint16(sym_glob),
	4974:  uint16(sym_word),
	4975:  uint16(3),
	4976:  uint16(3),
	4977:  uint16(1),
	4978:  uint16(sym_comment),
	4979:  uint16(489),
	4980:  uint16(1),
	4981:  uint16(sym__begin_brace),
	4982:  uint16(487),
	4983:  uint16(27),
	4984:  uint16(anon_sym_and),
	4985:  uint16(anon_sym_or),
	4986:  uint16(anon_sym_BANG),
	4987:  uint16(anon_sym_not),
	4988:  uint16(anon_sym_DOLLAR),
	4989:  uint16(anon_sym_LPAREN),
	4990:  uint16(anon_sym_function),
	4991:  uint16(anon_sym_end),
	4992:  uint16(sym_integer),
	4993:  uint16(sym_float),
	4994:  uint16(anon_sym_return),
	4995:  uint16(anon_sym_switch),
	4996:  uint16(sym_break),
	4997:  uint16(sym_continue),
	4998:  uint16(anon_sym_for),
	4999:  uint16(anon_sym_while),
	5000:  uint16(anon_sym_if),
	5001:  uint16(anon_sym_begin),
	5002:  uint16(anon_sym_LBRACK),
	5003:  uint16(anon_sym_RBRACK),
	5004:  uint16(anon_sym_LBRACE),
	5005:  uint16(anon_sym_DQUOTE),
	5006:  uint16(anon_sym_SQUOTE),
	5007:  uint16(sym_escape_sequence),
	5008:  uint16(sym_home_dir_expansion),
	5009:  uint16(sym_glob),
	5010:  uint16(sym_word),
	5011:  uint16(3),
	5012:  uint16(3),
	5013:  uint16(1),
	5014:  uint16(sym_comment),
	5015:  uint16(638),
	5016:  uint16(1),
	5017:  uint16(sym__concat),
	5018:  uint16(636),
	5019:  uint16(27),
	5020:  uint16(anon_sym_PIPE_PIPE),
	5021:  uint16(anon_sym_AMP_AMP),
	5022:  uint16(anon_sym_AMP_PIPE),
	5023:  uint16(anon_sym_2_GT_PIPE),
	5024:  uint16(anon_sym_PIPE),
	5025:  uint16(anon_sym_SEMI),
	5026:  uint16(anon_sym_AMP),
	5027:  uint16(anon_sym_LF),
	5028:  uint16(anon_sym_CR),
	5029:  uint16(anon_sym_CR_LF),
	5030:  uint16(anon_sym_DOLLAR),
	5031:  uint16(anon_sym_LPAREN),
	5032:  uint16(anon_sym_RPAREN),
	5033:  uint16(sym_integer),
	5034:  uint16(sym_float),
	5035:  uint16(anon_sym_RBRACE),
	5036:  uint16(anon_sym_LBRACK),
	5037:  uint16(anon_sym_RBRACK),
	5038:  uint16(anon_sym_LBRACE),
	5039:  uint16(anon_sym_DQUOTE),
	5040:  uint16(anon_sym_SQUOTE),
	5041:  uint16(sym_escape_sequence),
	5042:  uint16(sym_stream_redirect),
	5043:  uint16(sym_direction),
	5044:  uint16(sym_home_dir_expansion),
	5045:  uint16(sym_glob),
	5046:  uint16(sym_word),
	5047:  uint16(3),
	5048:  uint16(3),
	5049:  uint16(1),
	5050:  uint16(sym_comment),
	5051:  uint16(642),
	5052:  uint16(1),
	5053:  uint16(sym__concat),
	5054:  uint16(640),
	5055:  uint16(27),
	5056:  uint16(anon_sym_PIPE_PIPE),
	5057:  uint16(anon_sym_AMP_AMP),
	5058:  uint16(anon_sym_AMP_PIPE),
	5059:  uint16(anon_sym_2_GT_PIPE),
	5060:  uint16(anon_sym_PIPE),
	5061:  uint16(anon_sym_SEMI),
	5062:  uint16(anon_sym_AMP),
	5063:  uint16(anon_sym_LF),
	5064:  uint16(anon_sym_CR),
	5065:  uint16(anon_sym_CR_LF),
	5066:  uint16(anon_sym_DOLLAR),
	5067:  uint16(anon_sym_LPAREN),
	5068:  uint16(anon_sym_RPAREN),
	5069:  uint16(sym_integer),
	5070:  uint16(sym_float),
	5071:  uint16(anon_sym_RBRACE),
	5072:  uint16(anon_sym_LBRACK),
	5073:  uint16(anon_sym_RBRACK),
	5074:  uint16(anon_sym_LBRACE),
	5075:  uint16(anon_sym_DQUOTE),
	5076:  uint16(anon_sym_SQUOTE),
	5077:  uint16(sym_escape_sequence),
	5078:  uint16(sym_stream_redirect),
	5079:  uint16(sym_direction),
	5080:  uint16(sym_home_dir_expansion),
	5081:  uint16(sym_glob),
	5082:  uint16(sym_word),
	5083:  uint16(3),
	5084:  uint16(3),
	5085:  uint16(1),
	5086:  uint16(sym_comment),
	5087:  uint16(646),
	5088:  uint16(1),
	5089:  uint16(sym__concat),
	5090:  uint16(644),
	5091:  uint16(27),
	5092:  uint16(anon_sym_PIPE_PIPE),
	5093:  uint16(anon_sym_AMP_AMP),
	5094:  uint16(anon_sym_AMP_PIPE),
	5095:  uint16(anon_sym_2_GT_PIPE),
	5096:  uint16(anon_sym_PIPE),
	5097:  uint16(anon_sym_SEMI),
	5098:  uint16(anon_sym_AMP),
	5099:  uint16(anon_sym_LF),
	5100:  uint16(anon_sym_CR),
	5101:  uint16(anon_sym_CR_LF),
	5102:  uint16(anon_sym_DOLLAR),
	5103:  uint16(anon_sym_LPAREN),
	5104:  uint16(anon_sym_RPAREN),
	5105:  uint16(sym_integer),
	5106:  uint16(sym_float),
	5107:  uint16(anon_sym_RBRACE),
	5108:  uint16(anon_sym_LBRACK),
	5109:  uint16(anon_sym_RBRACK),
	5110:  uint16(anon_sym_LBRACE),
	5111:  uint16(anon_sym_DQUOTE),
	5112:  uint16(anon_sym_SQUOTE),
	5113:  uint16(sym_escape_sequence),
	5114:  uint16(sym_stream_redirect),
	5115:  uint16(sym_direction),
	5116:  uint16(sym_home_dir_expansion),
	5117:  uint16(sym_glob),
	5118:  uint16(sym_word),
	5119:  uint16(3),
	5120:  uint16(3),
	5121:  uint16(1),
	5122:  uint16(sym_comment),
	5123:  uint16(650),
	5124:  uint16(1),
	5125:  uint16(sym__concat),
	5126:  uint16(648),
	5127:  uint16(27),
	5128:  uint16(anon_sym_PIPE_PIPE),
	5129:  uint16(anon_sym_AMP_AMP),
	5130:  uint16(anon_sym_AMP_PIPE),
	5131:  uint16(anon_sym_2_GT_PIPE),
	5132:  uint16(anon_sym_PIPE),
	5133:  uint16(anon_sym_SEMI),
	5134:  uint16(anon_sym_AMP),
	5135:  uint16(anon_sym_LF),
	5136:  uint16(anon_sym_CR),
	5137:  uint16(anon_sym_CR_LF),
	5138:  uint16(anon_sym_DOLLAR),
	5139:  uint16(anon_sym_LPAREN),
	5140:  uint16(anon_sym_RPAREN),
	5141:  uint16(sym_integer),
	5142:  uint16(sym_float),
	5143:  uint16(anon_sym_RBRACE),
	5144:  uint16(anon_sym_LBRACK),
	5145:  uint16(anon_sym_RBRACK),
	5146:  uint16(anon_sym_LBRACE),
	5147:  uint16(anon_sym_DQUOTE),
	5148:  uint16(anon_sym_SQUOTE),
	5149:  uint16(sym_escape_sequence),
	5150:  uint16(sym_stream_redirect),
	5151:  uint16(sym_direction),
	5152:  uint16(sym_home_dir_expansion),
	5153:  uint16(sym_glob),
	5154:  uint16(sym_word),
	5155:  uint16(3),
	5156:  uint16(3),
	5157:  uint16(1),
	5158:  uint16(sym_comment),
	5159:  uint16(654),
	5160:  uint16(1),
	5161:  uint16(sym__concat),
	5162:  uint16(652),
	5163:  uint16(27),
	5164:  uint16(anon_sym_PIPE_PIPE),
	5165:  uint16(anon_sym_AMP_AMP),
	5166:  uint16(anon_sym_AMP_PIPE),
	5167:  uint16(anon_sym_2_GT_PIPE),
	5168:  uint16(anon_sym_PIPE),
	5169:  uint16(anon_sym_SEMI),
	5170:  uint16(anon_sym_AMP),
	5171:  uint16(anon_sym_LF),
	5172:  uint16(anon_sym_CR),
	5173:  uint16(anon_sym_CR_LF),
	5174:  uint16(anon_sym_DOLLAR),
	5175:  uint16(anon_sym_LPAREN),
	5176:  uint16(anon_sym_RPAREN),
	5177:  uint16(sym_integer),
	5178:  uint16(sym_float),
	5179:  uint16(anon_sym_RBRACE),
	5180:  uint16(anon_sym_LBRACK),
	5181:  uint16(anon_sym_RBRACK),
	5182:  uint16(anon_sym_LBRACE),
	5183:  uint16(anon_sym_DQUOTE),
	5184:  uint16(anon_sym_SQUOTE),
	5185:  uint16(sym_escape_sequence),
	5186:  uint16(sym_stream_redirect),
	5187:  uint16(sym_direction),
	5188:  uint16(sym_home_dir_expansion),
	5189:  uint16(sym_glob),
	5190:  uint16(sym_word),
	5191:  uint16(3),
	5192:  uint16(3),
	5193:  uint16(1),
	5194:  uint16(sym_comment),
	5195:  uint16(656),
	5196:  uint16(1),
	5197:  uint16(sym__concat),
	5198:  uint16(544),
	5199:  uint16(27),
	5200:  uint16(anon_sym_PIPE_PIPE),
	5201:  uint16(anon_sym_AMP_AMP),
	5202:  uint16(anon_sym_AMP_PIPE),
	5203:  uint16(anon_sym_2_GT_PIPE),
	5204:  uint16(anon_sym_PIPE),
	5205:  uint16(anon_sym_SEMI),
	5206:  uint16(anon_sym_AMP),
	5207:  uint16(anon_sym_LF),
	5208:  uint16(anon_sym_CR),
	5209:  uint16(anon_sym_CR_LF),
	5210:  uint16(anon_sym_DOLLAR),
	5211:  uint16(anon_sym_LPAREN),
	5212:  uint16(anon_sym_RPAREN),
	5213:  uint16(sym_integer),
	5214:  uint16(sym_float),
	5215:  uint16(anon_sym_RBRACE),
	5216:  uint16(anon_sym_LBRACK),
	5217:  uint16(anon_sym_RBRACK),
	5218:  uint16(anon_sym_LBRACE),
	5219:  uint16(anon_sym_DQUOTE),
	5220:  uint16(anon_sym_SQUOTE),
	5221:  uint16(sym_escape_sequence),
	5222:  uint16(sym_stream_redirect),
	5223:  uint16(sym_direction),
	5224:  uint16(sym_home_dir_expansion),
	5225:  uint16(sym_glob),
	5226:  uint16(sym_word),
	5227:  uint16(3),
	5228:  uint16(3),
	5229:  uint16(1),
	5230:  uint16(sym_comment),
	5231:  uint16(660),
	5232:  uint16(1),
	5233:  uint16(sym__concat),
	5234:  uint16(658),
	5235:  uint16(27),
	5236:  uint16(anon_sym_PIPE_PIPE),
	5237:  uint16(anon_sym_AMP_AMP),
	5238:  uint16(anon_sym_AMP_PIPE),
	5239:  uint16(anon_sym_2_GT_PIPE),
	5240:  uint16(anon_sym_PIPE),
	5241:  uint16(anon_sym_SEMI),
	5242:  uint16(anon_sym_AMP),
	5243:  uint16(anon_sym_LF),
	5244:  uint16(anon_sym_CR),
	5245:  uint16(anon_sym_CR_LF),
	5246:  uint16(anon_sym_DOLLAR),
	5247:  uint16(anon_sym_LPAREN),
	5248:  uint16(anon_sym_RPAREN),
	5249:  uint16(sym_integer),
	5250:  uint16(sym_float),
	5251:  uint16(anon_sym_RBRACE),
	5252:  uint16(anon_sym_LBRACK),
	5253:  uint16(anon_sym_RBRACK),
	5254:  uint16(anon_sym_LBRACE),
	5255:  uint16(anon_sym_DQUOTE),
	5256:  uint16(anon_sym_SQUOTE),
	5257:  uint16(sym_escape_sequence),
	5258:  uint16(sym_stream_redirect),
	5259:  uint16(sym_direction),
	5260:  uint16(sym_home_dir_expansion),
	5261:  uint16(sym_glob),
	5262:  uint16(sym_word),
	5263:  uint16(3),
	5264:  uint16(3),
	5265:  uint16(1),
	5266:  uint16(sym_comment),
	5267:  uint16(664),
	5268:  uint16(1),
	5269:  uint16(sym__concat),
	5270:  uint16(662),
	5271:  uint16(27),
	5272:  uint16(anon_sym_PIPE_PIPE),
	5273:  uint16(anon_sym_AMP_AMP),
	5274:  uint16(anon_sym_AMP_PIPE),
	5275:  uint16(anon_sym_2_GT_PIPE),
	5276:  uint16(anon_sym_PIPE),
	5277:  uint16(anon_sym_SEMI),
	5278:  uint16(anon_sym_AMP),
	5279:  uint16(anon_sym_LF),
	5280:  uint16(anon_sym_CR),
	5281:  uint16(anon_sym_CR_LF),
	5282:  uint16(anon_sym_DOLLAR),
	5283:  uint16(anon_sym_LPAREN),
	5284:  uint16(anon_sym_RPAREN),
	5285:  uint16(sym_integer),
	5286:  uint16(sym_float),
	5287:  uint16(anon_sym_RBRACE),
	5288:  uint16(anon_sym_LBRACK),
	5289:  uint16(anon_sym_RBRACK),
	5290:  uint16(anon_sym_LBRACE),
	5291:  uint16(anon_sym_DQUOTE),
	5292:  uint16(anon_sym_SQUOTE),
	5293:  uint16(sym_escape_sequence),
	5294:  uint16(sym_stream_redirect),
	5295:  uint16(sym_direction),
	5296:  uint16(sym_home_dir_expansion),
	5297:  uint16(sym_glob),
	5298:  uint16(sym_word),
	5299:  uint16(3),
	5300:  uint16(3),
	5301:  uint16(1),
	5302:  uint16(sym_comment),
	5303:  uint16(668),
	5304:  uint16(1),
	5305:  uint16(sym__concat),
	5306:  uint16(666),
	5307:  uint16(27),
	5308:  uint16(anon_sym_PIPE_PIPE),
	5309:  uint16(anon_sym_AMP_AMP),
	5310:  uint16(anon_sym_AMP_PIPE),
	5311:  uint16(anon_sym_2_GT_PIPE),
	5312:  uint16(anon_sym_PIPE),
	5313:  uint16(anon_sym_SEMI),
	5314:  uint16(anon_sym_AMP),
	5315:  uint16(anon_sym_LF),
	5316:  uint16(anon_sym_CR),
	5317:  uint16(anon_sym_CR_LF),
	5318:  uint16(anon_sym_DOLLAR),
	5319:  uint16(anon_sym_LPAREN),
	5320:  uint16(anon_sym_RPAREN),
	5321:  uint16(sym_integer),
	5322:  uint16(sym_float),
	5323:  uint16(anon_sym_RBRACE),
	5324:  uint16(anon_sym_LBRACK),
	5325:  uint16(anon_sym_RBRACK),
	5326:  uint16(anon_sym_LBRACE),
	5327:  uint16(anon_sym_DQUOTE),
	5328:  uint16(anon_sym_SQUOTE),
	5329:  uint16(sym_escape_sequence),
	5330:  uint16(sym_stream_redirect),
	5331:  uint16(sym_direction),
	5332:  uint16(sym_home_dir_expansion),
	5333:  uint16(sym_glob),
	5334:  uint16(sym_word),
	5335:  uint16(3),
	5336:  uint16(3),
	5337:  uint16(1),
	5338:  uint16(sym_comment),
	5339:  uint16(672),
	5340:  uint16(1),
	5341:  uint16(sym__concat),
	5342:  uint16(670),
	5343:  uint16(27),
	5344:  uint16(anon_sym_PIPE_PIPE),
	5345:  uint16(anon_sym_AMP_AMP),
	5346:  uint16(anon_sym_AMP_PIPE),
	5347:  uint16(anon_sym_2_GT_PIPE),
	5348:  uint16(anon_sym_PIPE),
	5349:  uint16(anon_sym_SEMI),
	5350:  uint16(anon_sym_AMP),
	5351:  uint16(anon_sym_LF),
	5352:  uint16(anon_sym_CR),
	5353:  uint16(anon_sym_CR_LF),
	5354:  uint16(anon_sym_DOLLAR),
	5355:  uint16(anon_sym_LPAREN),
	5356:  uint16(anon_sym_RPAREN),
	5357:  uint16(sym_integer),
	5358:  uint16(sym_float),
	5359:  uint16(anon_sym_RBRACE),
	5360:  uint16(anon_sym_LBRACK),
	5361:  uint16(anon_sym_RBRACK),
	5362:  uint16(anon_sym_LBRACE),
	5363:  uint16(anon_sym_DQUOTE),
	5364:  uint16(anon_sym_SQUOTE),
	5365:  uint16(sym_escape_sequence),
	5366:  uint16(sym_stream_redirect),
	5367:  uint16(sym_direction),
	5368:  uint16(sym_home_dir_expansion),
	5369:  uint16(sym_glob),
	5370:  uint16(sym_word),
	5371:  uint16(3),
	5372:  uint16(3),
	5373:  uint16(1),
	5374:  uint16(sym_comment),
	5375:  uint16(676),
	5376:  uint16(1),
	5377:  uint16(sym__concat),
	5378:  uint16(674),
	5379:  uint16(27),
	5380:  uint16(anon_sym_PIPE_PIPE),
	5381:  uint16(anon_sym_AMP_AMP),
	5382:  uint16(anon_sym_AMP_PIPE),
	5383:  uint16(anon_sym_2_GT_PIPE),
	5384:  uint16(anon_sym_PIPE),
	5385:  uint16(anon_sym_SEMI),
	5386:  uint16(anon_sym_AMP),
	5387:  uint16(anon_sym_LF),
	5388:  uint16(anon_sym_CR),
	5389:  uint16(anon_sym_CR_LF),
	5390:  uint16(anon_sym_DOLLAR),
	5391:  uint16(anon_sym_LPAREN),
	5392:  uint16(anon_sym_RPAREN),
	5393:  uint16(sym_integer),
	5394:  uint16(sym_float),
	5395:  uint16(anon_sym_RBRACE),
	5396:  uint16(anon_sym_LBRACK),
	5397:  uint16(anon_sym_RBRACK),
	5398:  uint16(anon_sym_LBRACE),
	5399:  uint16(anon_sym_DQUOTE),
	5400:  uint16(anon_sym_SQUOTE),
	5401:  uint16(sym_escape_sequence),
	5402:  uint16(sym_stream_redirect),
	5403:  uint16(sym_direction),
	5404:  uint16(sym_home_dir_expansion),
	5405:  uint16(sym_glob),
	5406:  uint16(sym_word),
	5407:  uint16(3),
	5408:  uint16(3),
	5409:  uint16(1),
	5410:  uint16(sym_comment),
	5411:  uint16(680),
	5412:  uint16(1),
	5413:  uint16(sym__concat),
	5414:  uint16(678),
	5415:  uint16(27),
	5416:  uint16(anon_sym_PIPE_PIPE),
	5417:  uint16(anon_sym_AMP_AMP),
	5418:  uint16(anon_sym_AMP_PIPE),
	5419:  uint16(anon_sym_2_GT_PIPE),
	5420:  uint16(anon_sym_PIPE),
	5421:  uint16(anon_sym_SEMI),
	5422:  uint16(anon_sym_AMP),
	5423:  uint16(anon_sym_LF),
	5424:  uint16(anon_sym_CR),
	5425:  uint16(anon_sym_CR_LF),
	5426:  uint16(anon_sym_DOLLAR),
	5427:  uint16(anon_sym_LPAREN),
	5428:  uint16(anon_sym_RPAREN),
	5429:  uint16(sym_integer),
	5430:  uint16(sym_float),
	5431:  uint16(anon_sym_RBRACE),
	5432:  uint16(anon_sym_LBRACK),
	5433:  uint16(anon_sym_RBRACK),
	5434:  uint16(anon_sym_LBRACE),
	5435:  uint16(anon_sym_DQUOTE),
	5436:  uint16(anon_sym_SQUOTE),
	5437:  uint16(sym_escape_sequence),
	5438:  uint16(sym_stream_redirect),
	5439:  uint16(sym_direction),
	5440:  uint16(sym_home_dir_expansion),
	5441:  uint16(sym_glob),
	5442:  uint16(sym_word),
	5443:  uint16(3),
	5444:  uint16(3),
	5445:  uint16(1),
	5446:  uint16(sym_comment),
	5447:  uint16(684),
	5448:  uint16(1),
	5449:  uint16(sym__concat),
	5450:  uint16(682),
	5451:  uint16(27),
	5452:  uint16(anon_sym_PIPE_PIPE),
	5453:  uint16(anon_sym_AMP_AMP),
	5454:  uint16(anon_sym_AMP_PIPE),
	5455:  uint16(anon_sym_2_GT_PIPE),
	5456:  uint16(anon_sym_PIPE),
	5457:  uint16(anon_sym_SEMI),
	5458:  uint16(anon_sym_AMP),
	5459:  uint16(anon_sym_LF),
	5460:  uint16(anon_sym_CR),
	5461:  uint16(anon_sym_CR_LF),
	5462:  uint16(anon_sym_DOLLAR),
	5463:  uint16(anon_sym_LPAREN),
	5464:  uint16(anon_sym_RPAREN),
	5465:  uint16(sym_integer),
	5466:  uint16(sym_float),
	5467:  uint16(anon_sym_RBRACE),
	5468:  uint16(anon_sym_LBRACK),
	5469:  uint16(anon_sym_RBRACK),
	5470:  uint16(anon_sym_LBRACE),
	5471:  uint16(anon_sym_DQUOTE),
	5472:  uint16(anon_sym_SQUOTE),
	5473:  uint16(sym_escape_sequence),
	5474:  uint16(sym_stream_redirect),
	5475:  uint16(sym_direction),
	5476:  uint16(sym_home_dir_expansion),
	5477:  uint16(sym_glob),
	5478:  uint16(sym_word),
	5479:  uint16(3),
	5480:  uint16(3),
	5481:  uint16(1),
	5482:  uint16(sym_comment),
	5483:  uint16(688),
	5484:  uint16(1),
	5485:  uint16(sym__concat),
	5486:  uint16(686),
	5487:  uint16(27),
	5488:  uint16(anon_sym_PIPE_PIPE),
	5489:  uint16(anon_sym_AMP_AMP),
	5490:  uint16(anon_sym_AMP_PIPE),
	5491:  uint16(anon_sym_2_GT_PIPE),
	5492:  uint16(anon_sym_PIPE),
	5493:  uint16(anon_sym_SEMI),
	5494:  uint16(anon_sym_AMP),
	5495:  uint16(anon_sym_LF),
	5496:  uint16(anon_sym_CR),
	5497:  uint16(anon_sym_CR_LF),
	5498:  uint16(anon_sym_DOLLAR),
	5499:  uint16(anon_sym_LPAREN),
	5500:  uint16(anon_sym_RPAREN),
	5501:  uint16(sym_integer),
	5502:  uint16(sym_float),
	5503:  uint16(anon_sym_RBRACE),
	5504:  uint16(anon_sym_LBRACK),
	5505:  uint16(anon_sym_RBRACK),
	5506:  uint16(anon_sym_LBRACE),
	5507:  uint16(anon_sym_DQUOTE),
	5508:  uint16(anon_sym_SQUOTE),
	5509:  uint16(sym_escape_sequence),
	5510:  uint16(sym_stream_redirect),
	5511:  uint16(sym_direction),
	5512:  uint16(sym_home_dir_expansion),
	5513:  uint16(sym_glob),
	5514:  uint16(sym_word),
	5515:  uint16(2),
	5516:  uint16(3),
	5517:  uint16(1),
	5518:  uint16(sym_comment),
	5519:  uint16(690),
	5520:  uint16(27),
	5521:  uint16(anon_sym_PIPE_PIPE),
	5522:  uint16(anon_sym_AMP_AMP),
	5523:  uint16(anon_sym_AMP_PIPE),
	5524:  uint16(anon_sym_2_GT_PIPE),
	5525:  uint16(anon_sym_PIPE),
	5526:  uint16(anon_sym_SEMI),
	5527:  uint16(anon_sym_AMP),
	5528:  uint16(anon_sym_LF),
	5529:  uint16(anon_sym_CR),
	5530:  uint16(anon_sym_CR_LF),
	5531:  uint16(anon_sym_DOLLAR),
	5532:  uint16(anon_sym_LPAREN),
	5533:  uint16(anon_sym_RPAREN),
	5534:  uint16(sym_integer),
	5535:  uint16(sym_float),
	5536:  uint16(anon_sym_RBRACE),
	5537:  uint16(anon_sym_LBRACK),
	5538:  uint16(anon_sym_RBRACK),
	5539:  uint16(anon_sym_LBRACE),
	5540:  uint16(anon_sym_DQUOTE),
	5541:  uint16(anon_sym_SQUOTE),
	5542:  uint16(sym_escape_sequence),
	5543:  uint16(sym_stream_redirect),
	5544:  uint16(sym_direction),
	5545:  uint16(sym_home_dir_expansion),
	5546:  uint16(sym_glob),
	5547:  uint16(sym_word),
	5548:  uint16(2),
	5549:  uint16(3),
	5550:  uint16(1),
	5551:  uint16(sym_comment),
	5552:  uint16(692),
	5553:  uint16(27),
	5554:  uint16(anon_sym_PIPE_PIPE),
	5555:  uint16(anon_sym_AMP_AMP),
	5556:  uint16(anon_sym_AMP_PIPE),
	5557:  uint16(anon_sym_2_GT_PIPE),
	5558:  uint16(anon_sym_PIPE),
	5559:  uint16(anon_sym_SEMI),
	5560:  uint16(anon_sym_AMP),
	5561:  uint16(anon_sym_LF),
	5562:  uint16(anon_sym_CR),
	5563:  uint16(anon_sym_CR_LF),
	5564:  uint16(anon_sym_DOLLAR),
	5565:  uint16(anon_sym_LPAREN),
	5566:  uint16(anon_sym_RPAREN),
	5567:  uint16(sym_integer),
	5568:  uint16(sym_float),
	5569:  uint16(anon_sym_RBRACE),
	5570:  uint16(anon_sym_LBRACK),
	5571:  uint16(anon_sym_RBRACK),
	5572:  uint16(anon_sym_LBRACE),
	5573:  uint16(anon_sym_DQUOTE),
	5574:  uint16(anon_sym_SQUOTE),
	5575:  uint16(sym_escape_sequence),
	5576:  uint16(sym_stream_redirect),
	5577:  uint16(sym_direction),
	5578:  uint16(sym_home_dir_expansion),
	5579:  uint16(sym_glob),
	5580:  uint16(sym_word),
	5581:  uint16(2),
	5582:  uint16(3),
	5583:  uint16(1),
	5584:  uint16(sym_comment),
	5585:  uint16(694),
	5586:  uint16(27),
	5587:  uint16(anon_sym_PIPE_PIPE),
	5588:  uint16(anon_sym_AMP_AMP),
	5589:  uint16(anon_sym_AMP_PIPE),
	5590:  uint16(anon_sym_2_GT_PIPE),
	5591:  uint16(anon_sym_PIPE),
	5592:  uint16(anon_sym_SEMI),
	5593:  uint16(anon_sym_AMP),
	5594:  uint16(anon_sym_LF),
	5595:  uint16(anon_sym_CR),
	5596:  uint16(anon_sym_CR_LF),
	5597:  uint16(anon_sym_DOLLAR),
	5598:  uint16(anon_sym_LPAREN),
	5599:  uint16(anon_sym_RPAREN),
	5600:  uint16(sym_integer),
	5601:  uint16(sym_float),
	5602:  uint16(anon_sym_RBRACE),
	5603:  uint16(anon_sym_LBRACK),
	5604:  uint16(anon_sym_RBRACK),
	5605:  uint16(anon_sym_LBRACE),
	5606:  uint16(anon_sym_DQUOTE),
	5607:  uint16(anon_sym_SQUOTE),
	5608:  uint16(sym_escape_sequence),
	5609:  uint16(sym_stream_redirect),
	5610:  uint16(sym_direction),
	5611:  uint16(sym_home_dir_expansion),
	5612:  uint16(sym_glob),
	5613:  uint16(sym_word),
	5614:  uint16(13),
	5615:  uint16(3),
	5616:  uint16(1),
	5617:  uint16(sym_comment),
	5618:  uint16(516),
	5619:  uint16(1),
	5620:  uint16(anon_sym_DOLLAR),
	5621:  uint16(518),
	5622:  uint16(1),
	5623:  uint16(anon_sym_LPAREN),
	5624:  uint16(522),
	5625:  uint16(1),
	5626:  uint16(anon_sym_LBRACE),
	5627:  uint16(524),
	5628:  uint16(1),
	5629:  uint16(anon_sym_DQUOTE),
	5630:  uint16(526),
	5631:  uint16(1),
	5632:  uint16(anon_sym_SQUOTE),
	5633:  uint16(95),
	5634:  uint16(1),
	5635:  uint16(aux_sym_for_statement_repeat1),
	5636:  uint16(142),
	5637:  uint16(1),
	5638:  uint16(sym__special_character),
	5639:  uint16(520),
	5640:  uint16(2),
	5641:  uint16(anon_sym_LBRACK),
	5642:  uint16(anon_sym_RBRACK),
	5643:  uint16(157),
	5644:  uint16(2),
	5645:  uint16(sym__command_substitution_dollar),
	5646:  uint16(sym__command_substitution_inner),
	5647:  uint16(158),
	5648:  uint16(2),
	5649:  uint16(sym_concatenation),
	5650:  uint16(sym__expression),
	5651:  uint16(137),
	5652:  uint16(5),
	5653:  uint16(sym_command_substitution),
	5654:  uint16(sym_variable_expansion),
	5655:  uint16(sym_brace_expansion),
	5656:  uint16(sym_double_quote_string),
	5657:  uint16(sym_single_quote_string),
	5658:  uint16(512),
	5659:  uint16(6),
	5660:  uint16(sym_integer),
	5661:  uint16(sym_float),
	5662:  uint16(sym_escape_sequence),
	5663:  uint16(sym_home_dir_expansion),
	5664:  uint16(sym_glob),
	5665:  uint16(sym_word),
	5666:  uint16(12),
	5667:  uint16(3),
	5668:  uint16(1),
	5669:  uint16(sym_comment),
	5670:  uint16(516),
	5671:  uint16(1),
	5672:  uint16(anon_sym_DOLLAR),
	5673:  uint16(518),
	5674:  uint16(1),
	5675:  uint16(anon_sym_LPAREN),
	5676:  uint16(522),
	5677:  uint16(1),
	5678:  uint16(anon_sym_LBRACE),
	5679:  uint16(524),
	5680:  uint16(1),
	5681:  uint16(anon_sym_DQUOTE),
	5682:  uint16(526),
	5683:  uint16(1),
	5684:  uint16(anon_sym_SQUOTE),
	5685:  uint16(142),
	5686:  uint16(1),
	5687:  uint16(sym__special_character),
	5688:  uint16(520),
	5689:  uint16(2),
	5690:  uint16(anon_sym_LBRACK),
	5691:  uint16(anon_sym_RBRACK),
	5692:  uint16(157),
	5693:  uint16(2),
	5694:  uint16(sym__command_substitution_dollar),
	5695:  uint16(sym__command_substitution_inner),
	5696:  uint16(85),
	5697:  uint16(3),
	5698:  uint16(sym_concatenation),
	5699:  uint16(sym__expression),
	5700:  uint16(aux_sym_case_clause_repeat1),
	5701:  uint16(137),
	5702:  uint16(5),
	5703:  uint16(sym_command_substitution),
	5704:  uint16(sym_variable_expansion),
	5705:  uint16(sym_brace_expansion),
	5706:  uint16(sym_double_quote_string),
	5707:  uint16(sym_single_quote_string),
	5708:  uint16(512),
	5709:  uint16(6),
	5710:  uint16(sym_integer),
	5711:  uint16(sym_float),
	5712:  uint16(sym_escape_sequence),
	5713:  uint16(sym_home_dir_expansion),
	5714:  uint16(sym_glob),
	5715:  uint16(sym_word),
	5716:  uint16(13),
	5717:  uint16(3),
	5718:  uint16(1),
	5719:  uint16(sym_comment),
	5720:  uint16(696),
	5721:  uint16(1),
	5722:  uint16(anon_sym_DOLLAR),
	5723:  uint16(698),
	5724:  uint16(1),
	5725:  uint16(anon_sym_LPAREN),
	5726:  uint16(702),
	5727:  uint16(1),
	5728:  uint16(anon_sym_RBRACE),
	5729:  uint16(704),
	5730:  uint16(1),
	5731:  uint16(anon_sym_LBRACE),
	5732:  uint16(706),
	5733:  uint16(1),
	5734:  uint16(anon_sym_COMMA),
	5735:  uint16(708),
	5736:  uint16(1),
	5737:  uint16(anon_sym_DQUOTE),
	5738:  uint16(710),
	5739:  uint16(1),
	5740:  uint16(anon_sym_SQUOTE),
	5741:  uint16(399),
	5742:  uint16(1),
	5743:  uint16(aux_sym_brace_expansion_repeat1),
	5744:  uint16(388),
	5745:  uint16(2),
	5746:  uint16(sym__command_substitution_dollar),
	5747:  uint16(sym__command_substitution_inner),
	5748:  uint16(403),
	5749:  uint16(2),
	5750:  uint16(sym_brace_concatenation),
	5751:  uint16(sym__brace_expression),
	5752:  uint16(700),
	5753:  uint16(5),
	5754:  uint16(sym_integer),
	5755:  uint16(sym_float),
	5756:  uint16(sym_escape_sequence),
	5757:  uint16(sym_glob),
	5758:  uint16(sym_brace_word),
	5759:  uint16(357),
	5760:  uint16(6),
	5761:  uint16(sym_command_substitution),
	5762:  uint16(sym_variable_expansion),
	5763:  uint16(sym_brace_expansion),
	5764:  uint16(sym_double_quote_string),
	5765:  uint16(sym_single_quote_string),
	5766:  uint16(sym__base_brace_expression),
	5767:  uint16(12),
	5768:  uint16(3),
	5769:  uint16(1),
	5770:  uint16(sym_comment),
	5771:  uint16(714),
	5772:  uint16(1),
	5773:  uint16(anon_sym_DOLLAR),
	5774:  uint16(716),
	5775:  uint16(1),
	5776:  uint16(anon_sym_LPAREN),
	5777:  uint16(720),
	5778:  uint16(1),
	5779:  uint16(anon_sym_LBRACE),
	5780:  uint16(722),
	5781:  uint16(1),
	5782:  uint16(anon_sym_DQUOTE),
	5783:  uint16(724),
	5784:  uint16(1),
	5785:  uint16(anon_sym_SQUOTE),
	5786:  uint16(307),
	5787:  uint16(1),
	5788:  uint16(sym__special_character),
	5789:  uint16(718),
	5790:  uint16(2),
	5791:  uint16(anon_sym_LBRACK),
	5792:  uint16(anon_sym_RBRACK),
	5793:  uint16(322),
	5794:  uint16(2),
	5795:  uint16(sym__command_substitution_dollar),
	5796:  uint16(sym__command_substitution_inner),
	5797:  uint16(350),
	5798:  uint16(2),
	5799:  uint16(sym_concatenation),
	5800:  uint16(sym__expression),
	5801:  uint16(300),
	5802:  uint16(5),
	5803:  uint16(sym_command_substitution),
	5804:  uint16(sym_variable_expansion),
	5805:  uint16(sym_brace_expansion),
	5806:  uint16(sym_double_quote_string),
	5807:  uint16(sym_single_quote_string),
	5808:  uint16(712),
	5809:  uint16(6),
	5810:  uint16(sym_integer),
	5811:  uint16(sym_float),
	5812:  uint16(sym_escape_sequence),
	5813:  uint16(sym_home_dir_expansion),
	5814:  uint16(sym_glob),
	5815:  uint16(sym_word),
	5816:  uint16(13),
	5817:  uint16(3),
	5818:  uint16(1),
	5819:  uint16(sym_comment),
	5820:  uint16(696),
	5821:  uint16(1),
	5822:  uint16(anon_sym_DOLLAR),
	5823:  uint16(698),
	5824:  uint16(1),
	5825:  uint16(anon_sym_LPAREN),
	5826:  uint16(704),
	5827:  uint16(1),
	5828:  uint16(anon_sym_LBRACE),
	5829:  uint16(706),
	5830:  uint16(1),
	5831:  uint16(anon_sym_COMMA),
	5832:  uint16(708),
	5833:  uint16(1),
	5834:  uint16(anon_sym_DQUOTE),
	5835:  uint16(710),
	5836:  uint16(1),
	5837:  uint16(anon_sym_SQUOTE),
	5838:  uint16(726),
	5839:  uint16(1),
	5840:  uint16(anon_sym_RBRACE),
	5841:  uint16(410),
	5842:  uint16(1),
	5843:  uint16(aux_sym_brace_expansion_repeat1),
	5844:  uint16(388),
	5845:  uint16(2),
	5846:  uint16(sym__command_substitution_dollar),
	5847:  uint16(sym__command_substitution_inner),
	5848:  uint16(397),
	5849:  uint16(2),
	5850:  uint16(sym_brace_concatenation),
	5851:  uint16(sym__brace_expression),
	5852:  uint16(700),
	5853:  uint16(5),
	5854:  uint16(sym_integer),
	5855:  uint16(sym_float),
	5856:  uint16(sym_escape_sequence),
	5857:  uint16(sym_glob),
	5858:  uint16(sym_brace_word),
	5859:  uint16(357),
	5860:  uint16(6),
	5861:  uint16(sym_command_substitution),
	5862:  uint16(sym_variable_expansion),
	5863:  uint16(sym_brace_expansion),
	5864:  uint16(sym_double_quote_string),
	5865:  uint16(sym_single_quote_string),
	5866:  uint16(sym__base_brace_expression),
	5867:  uint16(13),
	5868:  uint16(3),
	5869:  uint16(1),
	5870:  uint16(sym_comment),
	5871:  uint16(696),
	5872:  uint16(1),
	5873:  uint16(anon_sym_DOLLAR),
	5874:  uint16(698),
	5875:  uint16(1),
	5876:  uint16(anon_sym_LPAREN),
	5877:  uint16(704),
	5878:  uint16(1),
	5879:  uint16(anon_sym_LBRACE),
	5880:  uint16(706),
	5881:  uint16(1),
	5882:  uint16(anon_sym_COMMA),
	5883:  uint16(708),
	5884:  uint16(1),
	5885:  uint16(anon_sym_DQUOTE),
	5886:  uint16(710),
	5887:  uint16(1),
	5888:  uint16(anon_sym_SQUOTE),
	5889:  uint16(728),
	5890:  uint16(1),
	5891:  uint16(anon_sym_RBRACE),
	5892:  uint16(402),
	5893:  uint16(1),
	5894:  uint16(aux_sym_brace_expansion_repeat1),
	5895:  uint16(388),
	5896:  uint16(2),
	5897:  uint16(sym__command_substitution_dollar),
	5898:  uint16(sym__command_substitution_inner),
	5899:  uint16(401),
	5900:  uint16(2),
	5901:  uint16(sym_brace_concatenation),
	5902:  uint16(sym__brace_expression),
	5903:  uint16(700),
	5904:  uint16(5),
	5905:  uint16(sym_integer),
	5906:  uint16(sym_float),
	5907:  uint16(sym_escape_sequence),
	5908:  uint16(sym_glob),
	5909:  uint16(sym_brace_word),
	5910:  uint16(357),
	5911:  uint16(6),
	5912:  uint16(sym_command_substitution),
	5913:  uint16(sym_variable_expansion),
	5914:  uint16(sym_brace_expansion),
	5915:  uint16(sym_double_quote_string),
	5916:  uint16(sym_single_quote_string),
	5917:  uint16(sym__base_brace_expression),
	5918:  uint16(13),
	5919:  uint16(3),
	5920:  uint16(1),
	5921:  uint16(sym_comment),
	5922:  uint16(696),
	5923:  uint16(1),
	5924:  uint16(anon_sym_DOLLAR),
	5925:  uint16(698),
	5926:  uint16(1),
	5927:  uint16(anon_sym_LPAREN),
	5928:  uint16(704),
	5929:  uint16(1),
	5930:  uint16(anon_sym_LBRACE),
	5931:  uint16(706),
	5932:  uint16(1),
	5933:  uint16(anon_sym_COMMA),
	5934:  uint16(708),
	5935:  uint16(1),
	5936:  uint16(anon_sym_DQUOTE),
	5937:  uint16(710),
	5938:  uint16(1),
	5939:  uint16(anon_sym_SQUOTE),
	5940:  uint16(730),
	5941:  uint16(1),
	5942:  uint16(anon_sym_RBRACE),
	5943:  uint16(381),
	5944:  uint16(1),
	5945:  uint16(aux_sym_brace_expansion_repeat1),
	5946:  uint16(384),
	5947:  uint16(2),
	5948:  uint16(sym_brace_concatenation),
	5949:  uint16(sym__brace_expression),
	5950:  uint16(388),
	5951:  uint16(2),
	5952:  uint16(sym__command_substitution_dollar),
	5953:  uint16(sym__command_substitution_inner),
	5954:  uint16(700),
	5955:  uint16(5),
	5956:  uint16(sym_integer),
	5957:  uint16(sym_float),
	5958:  uint16(sym_escape_sequence),
	5959:  uint16(sym_glob),
	5960:  uint16(sym_brace_word),
	5961:  uint16(357),
	5962:  uint16(6),
	5963:  uint16(sym_command_substitution),
	5964:  uint16(sym_variable_expansion),
	5965:  uint16(sym_brace_expansion),
	5966:  uint16(sym_double_quote_string),
	5967:  uint16(sym_single_quote_string),
	5968:  uint16(sym__base_brace_expression),
	5969:  uint16(12),
	5970:  uint16(3),
	5971:  uint16(1),
	5972:  uint16(sym_comment),
	5973:  uint16(471),
	5974:  uint16(1),
	5975:  uint16(anon_sym_DOLLAR),
	5976:  uint16(473),
	5977:  uint16(1),
	5978:  uint16(anon_sym_LPAREN),
	5979:  uint16(477),
	5980:  uint16(1),
	5981:  uint16(anon_sym_LBRACE),
	5982:  uint16(479),
	5983:  uint16(1),
	5984:  uint16(anon_sym_DQUOTE),
	5985:  uint16(481),
	5986:  uint16(1),
	5987:  uint16(anon_sym_SQUOTE),
	5988:  uint16(178),
	5989:  uint16(1),
	5990:  uint16(sym__special_character),
	5991:  uint16(475),
	5992:  uint16(2),
	5993:  uint16(anon_sym_LBRACK),
	5994:  uint16(anon_sym_RBRACK),
	5995:  uint16(199),
	5996:  uint16(2),
	5997:  uint16(sym__command_substitution_dollar),
	5998:  uint16(sym__command_substitution_inner),
	5999:  uint16(216),
	6000:  uint16(2),
	6001:  uint16(sym_concatenation),
	6002:  uint16(sym__expression),
	6003:  uint16(182),
	6004:  uint16(5),
	6005:  uint16(sym_command_substitution),
	6006:  uint16(sym_variable_expansion),
	6007:  uint16(sym_brace_expansion),
	6008:  uint16(sym_double_quote_string),
	6009:  uint16(sym_single_quote_string),
	6010:  uint16(467),
	6011:  uint16(6),
	6012:  uint16(sym_integer),
	6013:  uint16(sym_float),
	6014:  uint16(sym_escape_sequence),
	6015:  uint16(sym_home_dir_expansion),
	6016:  uint16(sym_glob),
	6017:  uint16(sym_word),
	6018:  uint16(13),
	6019:  uint16(3),
	6020:  uint16(1),
	6021:  uint16(sym_comment),
	6022:  uint16(696),
	6023:  uint16(1),
	6024:  uint16(anon_sym_DOLLAR),
	6025:  uint16(698),
	6026:  uint16(1),
	6027:  uint16(anon_sym_LPAREN),
	6028:  uint16(704),
	6029:  uint16(1),
	6030:  uint16(anon_sym_LBRACE),
	6031:  uint16(706),
	6032:  uint16(1),
	6033:  uint16(anon_sym_COMMA),
	6034:  uint16(708),
	6035:  uint16(1),
	6036:  uint16(anon_sym_DQUOTE),
	6037:  uint16(710),
	6038:  uint16(1),
	6039:  uint16(anon_sym_SQUOTE),
	6040:  uint16(732),
	6041:  uint16(1),
	6042:  uint16(anon_sym_RBRACE),
	6043:  uint16(413),
	6044:  uint16(1),
	6045:  uint16(aux_sym_brace_expansion_repeat1),
	6046:  uint16(388),
	6047:  uint16(2),
	6048:  uint16(sym__command_substitution_dollar),
	6049:  uint16(sym__command_substitution_inner),
	6050:  uint16(415),
	6051:  uint16(2),
	6052:  uint16(sym_brace_concatenation),
	6053:  uint16(sym__brace_expression),
	6054:  uint16(700),
	6055:  uint16(5),
	6056:  uint16(sym_integer),
	6057:  uint16(sym_float),
	6058:  uint16(sym_escape_sequence),
	6059:  uint16(sym_glob),
	6060:  uint16(sym_brace_word),
	6061:  uint16(357),
	6062:  uint16(6),
	6063:  uint16(sym_command_substitution),
	6064:  uint16(sym_variable_expansion),
	6065:  uint16(sym_brace_expansion),
	6066:  uint16(sym_double_quote_string),
	6067:  uint16(sym_single_quote_string),
	6068:  uint16(sym__base_brace_expression),
	6069:  uint16(12),
	6070:  uint16(3),
	6071:  uint16(1),
	6072:  uint16(sym_comment),
	6073:  uint16(15),
	6074:  uint16(1),
	6075:  uint16(anon_sym_DOLLAR),
	6076:  uint16(17),
	6077:  uint16(1),
	6078:  uint16(anon_sym_LPAREN),
	6079:  uint16(37),
	6080:  uint16(1),
	6081:  uint16(anon_sym_LBRACE),
	6082:  uint16(39),
	6083:  uint16(1),
	6084:  uint16(anon_sym_DQUOTE),
	6085:  uint16(41),
	6086:  uint16(1),
	6087:  uint16(anon_sym_SQUOTE),
	6088:  uint16(87),
	6089:  uint16(1),
	6090:  uint16(sym__special_character),
	6091:  uint16(35),
	6092:  uint16(2),
	6093:  uint16(anon_sym_LBRACK),
	6094:  uint16(anon_sym_RBRACK),
	6095:  uint16(103),
	6096:  uint16(2),
	6097:  uint16(sym__command_substitution_dollar),
	6098:  uint16(sym__command_substitution_inner),
	6099:  uint16(114),
	6100:  uint16(2),
	6101:  uint16(sym_concatenation),
	6102:  uint16(sym__expression),
	6103:  uint16(84),
	6104:  uint16(5),
	6105:  uint16(sym_command_substitution),
	6106:  uint16(sym_variable_expansion),
	6107:  uint16(sym_brace_expansion),
	6108:  uint16(sym_double_quote_string),
	6109:  uint16(sym_single_quote_string),
	6110:  uint16(7),
	6111:  uint16(6),
	6112:  uint16(sym_integer),
	6113:  uint16(sym_float),
	6114:  uint16(sym_escape_sequence),
	6115:  uint16(sym_home_dir_expansion),
	6116:  uint16(sym_glob),
	6117:  uint16(sym_word),
	6118:  uint16(12),
	6119:  uint16(3),
	6120:  uint16(1),
	6121:  uint16(sym_comment),
	6122:  uint16(516),
	6123:  uint16(1),
	6124:  uint16(anon_sym_DOLLAR),
	6125:  uint16(518),
	6126:  uint16(1),
	6127:  uint16(anon_sym_LPAREN),
	6128:  uint16(522),
	6129:  uint16(1),
	6130:  uint16(anon_sym_LBRACE),
	6131:  uint16(524),
	6132:  uint16(1),
	6133:  uint16(anon_sym_DQUOTE),
	6134:  uint16(526),
	6135:  uint16(1),
	6136:  uint16(anon_sym_SQUOTE),
	6137:  uint16(142),
	6138:  uint16(1),
	6139:  uint16(sym__special_character),
	6140:  uint16(520),
	6141:  uint16(2),
	6142:  uint16(anon_sym_LBRACK),
	6143:  uint16(anon_sym_RBRACK),
	6144:  uint16(92),
	6145:  uint16(2),
	6146:  uint16(sym_concatenation),
	6147:  uint16(sym__expression),
	6148:  uint16(157),
	6149:  uint16(2),
	6150:  uint16(sym__command_substitution_dollar),
	6151:  uint16(sym__command_substitution_inner),
	6152:  uint16(137),
	6153:  uint16(5),
	6154:  uint16(sym_command_substitution),
	6155:  uint16(sym_variable_expansion),
	6156:  uint16(sym_brace_expansion),
	6157:  uint16(sym_double_quote_string),
	6158:  uint16(sym_single_quote_string),
	6159:  uint16(512),
	6160:  uint16(6),
	6161:  uint16(sym_integer),
	6162:  uint16(sym_float),
	6163:  uint16(sym_escape_sequence),
	6164:  uint16(sym_home_dir_expansion),
	6165:  uint16(sym_glob),
	6166:  uint16(sym_word),
	6167:  uint16(10),
	6168:  uint16(3),
	6169:  uint16(1),
	6170:  uint16(sym_comment),
	6171:  uint16(471),
	6172:  uint16(1),
	6173:  uint16(anon_sym_DOLLAR),
	6174:  uint16(473),
	6175:  uint16(1),
	6176:  uint16(anon_sym_LPAREN),
	6177:  uint16(477),
	6178:  uint16(1),
	6179:  uint16(anon_sym_LBRACE),
	6180:  uint16(479),
	6181:  uint16(1),
	6182:  uint16(anon_sym_DQUOTE),
	6183:  uint16(481),
	6184:  uint16(1),
	6185:  uint16(anon_sym_SQUOTE),
	6186:  uint16(475),
	6187:  uint16(2),
	6188:  uint16(anon_sym_LBRACK),
	6189:  uint16(anon_sym_RBRACK),
	6190:  uint16(199),
	6191:  uint16(2),
	6192:  uint16(sym__command_substitution_dollar),
	6193:  uint16(sym__command_substitution_inner),
	6194:  uint16(192),
	6195:  uint16(6),
	6196:  uint16(sym_command_substitution),
	6197:  uint16(sym_variable_expansion),
	6198:  uint16(sym_brace_expansion),
	6199:  uint16(sym_double_quote_string),
	6200:  uint16(sym_single_quote_string),
	6201:  uint16(sym__special_character),
	6202:  uint16(734),
	6203:  uint16(7),
	6204:  uint16(sym_integer),
	6205:  uint16(sym_float),
	6206:  uint16(sym_escape_sequence),
	6207:  uint16(anon_sym_POUND),
	6208:  uint16(sym_home_dir_expansion),
	6209:  uint16(sym_glob),
	6210:  uint16(sym_word),
	6211:  uint16(11),
	6212:  uint16(3),
	6213:  uint16(1),
	6214:  uint16(sym_comment),
	6215:  uint16(696),
	6216:  uint16(1),
	6217:  uint16(anon_sym_DOLLAR),
	6218:  uint16(698),
	6219:  uint16(1),
	6220:  uint16(anon_sym_LPAREN),
	6221:  uint16(704),
	6222:  uint16(1),
	6223:  uint16(anon_sym_LBRACE),
	6224:  uint16(708),
	6225:  uint16(1),
	6226:  uint16(anon_sym_DQUOTE),
	6227:  uint16(710),
	6228:  uint16(1),
	6229:  uint16(anon_sym_SQUOTE),
	6230:  uint16(736),
	6231:  uint16(2),
	6232:  uint16(anon_sym_RBRACE),
	6233:  uint16(anon_sym_COMMA),
	6234:  uint16(388),
	6235:  uint16(2),
	6236:  uint16(sym__command_substitution_dollar),
	6237:  uint16(sym__command_substitution_inner),
	6238:  uint16(424),
	6239:  uint16(2),
	6240:  uint16(sym_brace_concatenation),
	6241:  uint16(sym__brace_expression),
	6242:  uint16(700),
	6243:  uint16(5),
	6244:  uint16(sym_integer),
	6245:  uint16(sym_float),
	6246:  uint16(sym_escape_sequence),
	6247:  uint16(sym_glob),
	6248:  uint16(sym_brace_word),
	6249:  uint16(357),
	6250:  uint16(6),
	6251:  uint16(sym_command_substitution),
	6252:  uint16(sym_variable_expansion),
	6253:  uint16(sym_brace_expansion),
	6254:  uint16(sym_double_quote_string),
	6255:  uint16(sym_single_quote_string),
	6256:  uint16(sym__base_brace_expression),
	6257:  uint16(10),
	6258:  uint16(3),
	6259:  uint16(1),
	6260:  uint16(sym_comment),
	6261:  uint16(15),
	6262:  uint16(1),
	6263:  uint16(anon_sym_DOLLAR),
	6264:  uint16(17),
	6265:  uint16(1),
	6266:  uint16(anon_sym_LPAREN),
	6267:  uint16(37),
	6268:  uint16(1),
	6269:  uint16(anon_sym_LBRACE),
	6270:  uint16(39),
	6271:  uint16(1),
	6272:  uint16(anon_sym_DQUOTE),
	6273:  uint16(41),
	6274:  uint16(1),
	6275:  uint16(anon_sym_SQUOTE),
	6276:  uint16(35),
	6277:  uint16(2),
	6278:  uint16(anon_sym_LBRACK),
	6279:  uint16(anon_sym_RBRACK),
	6280:  uint16(103),
	6281:  uint16(2),
	6282:  uint16(sym__command_substitution_dollar),
	6283:  uint16(sym__command_substitution_inner),
	6284:  uint16(104),
	6285:  uint16(6),
	6286:  uint16(sym_command_substitution),
	6287:  uint16(sym_variable_expansion),
	6288:  uint16(sym_brace_expansion),
	6289:  uint16(sym_double_quote_string),
	6290:  uint16(sym_single_quote_string),
	6291:  uint16(sym__special_character),
	6292:  uint16(738),
	6293:  uint16(7),
	6294:  uint16(sym_integer),
	6295:  uint16(sym_float),
	6296:  uint16(sym_escape_sequence),
	6297:  uint16(anon_sym_POUND),
	6298:  uint16(sym_home_dir_expansion),
	6299:  uint16(sym_glob),
	6300:  uint16(sym_word),
	6301:  uint16(10),
	6302:  uint16(3),
	6303:  uint16(1),
	6304:  uint16(sym_comment),
	6305:  uint16(516),
	6306:  uint16(1),
	6307:  uint16(anon_sym_DOLLAR),
	6308:  uint16(518),
	6309:  uint16(1),
	6310:  uint16(anon_sym_LPAREN),
	6311:  uint16(522),
	6312:  uint16(1),
	6313:  uint16(anon_sym_LBRACE),
	6314:  uint16(524),
	6315:  uint16(1),
	6316:  uint16(anon_sym_DQUOTE),
	6317:  uint16(526),
	6318:  uint16(1),
	6319:  uint16(anon_sym_SQUOTE),
	6320:  uint16(520),
	6321:  uint16(2),
	6322:  uint16(anon_sym_LBRACK),
	6323:  uint16(anon_sym_RBRACK),
	6324:  uint16(157),
	6325:  uint16(2),
	6326:  uint16(sym__command_substitution_dollar),
	6327:  uint16(sym__command_substitution_inner),
	6328:  uint16(156),
	6329:  uint16(6),
	6330:  uint16(sym_command_substitution),
	6331:  uint16(sym_variable_expansion),
	6332:  uint16(sym_brace_expansion),
	6333:  uint16(sym_double_quote_string),
	6334:  uint16(sym_single_quote_string),
	6335:  uint16(sym__special_character),
	6336:  uint16(740),
	6337:  uint16(7),
	6338:  uint16(sym_integer),
	6339:  uint16(sym_float),
	6340:  uint16(sym_escape_sequence),
	6341:  uint16(anon_sym_POUND),
	6342:  uint16(sym_home_dir_expansion),
	6343:  uint16(sym_glob),
	6344:  uint16(sym_word),
	6345:  uint16(10),
	6346:  uint16(3),
	6347:  uint16(1),
	6348:  uint16(sym_comment),
	6349:  uint16(714),
	6350:  uint16(1),
	6351:  uint16(anon_sym_DOLLAR),
	6352:  uint16(716),
	6353:  uint16(1),
	6354:  uint16(anon_sym_LPAREN),
	6355:  uint16(720),
	6356:  uint16(1),
	6357:  uint16(anon_sym_LBRACE),
	6358:  uint16(722),
	6359:  uint16(1),
	6360:  uint16(anon_sym_DQUOTE),
	6361:  uint16(724),
	6362:  uint16(1),
	6363:  uint16(anon_sym_SQUOTE),
	6364:  uint16(718),
	6365:  uint16(2),
	6366:  uint16(anon_sym_LBRACK),
	6367:  uint16(anon_sym_RBRACK),
	6368:  uint16(322),
	6369:  uint16(2),
	6370:  uint16(sym__command_substitution_dollar),
	6371:  uint16(sym__command_substitution_inner),
	6372:  uint16(319),
	6373:  uint16(6),
	6374:  uint16(sym_command_substitution),
	6375:  uint16(sym_variable_expansion),
	6376:  uint16(sym_brace_expansion),
	6377:  uint16(sym_double_quote_string),
	6378:  uint16(sym_single_quote_string),
	6379:  uint16(sym__special_character),
	6380:  uint16(742),
	6381:  uint16(7),
	6382:  uint16(sym_integer),
	6383:  uint16(sym_float),
	6384:  uint16(sym_escape_sequence),
	6385:  uint16(anon_sym_POUND),
	6386:  uint16(sym_home_dir_expansion),
	6387:  uint16(sym_glob),
	6388:  uint16(sym_word),
	6389:  uint16(5),
	6390:  uint16(3),
	6391:  uint16(1),
	6392:  uint16(sym_comment),
	6393:  uint16(493),
	6394:  uint16(1),
	6395:  uint16(sym__concat),
	6396:  uint16(744),
	6397:  uint16(1),
	6398:  uint16(sym__concat_list),
	6399:  uint16(136),
	6400:  uint16(1),
	6401:  uint16(aux_sym_variable_expansion_repeat1),
	6402:  uint16(491),
	6403:  uint16(18),
	6404:  uint16(anon_sym_SEMI),
	6405:  uint16(anon_sym_AMP),
	6406:  uint16(anon_sym_LF),
	6407:  uint16(anon_sym_CR),
	6408:  uint16(anon_sym_CR_LF),
	6409:  uint16(anon_sym_DOLLAR),
	6410:  uint16(anon_sym_LPAREN),
	6411:  uint16(sym_integer),
	6412:  uint16(sym_float),
	6413:  uint16(anon_sym_LBRACK),
	6414:  uint16(anon_sym_RBRACK),
	6415:  uint16(anon_sym_LBRACE),
	6416:  uint16(anon_sym_DQUOTE),
	6417:  uint16(anon_sym_SQUOTE),
	6418:  uint16(sym_escape_sequence),
	6419:  uint16(sym_home_dir_expansion),
	6420:  uint16(sym_glob),
	6421:  uint16(sym_word),
	6422:  uint16(5),
	6423:  uint16(3),
	6424:  uint16(1),
	6425:  uint16(sym_comment),
	6426:  uint16(499),
	6427:  uint16(1),
	6428:  uint16(sym__concat),
	6429:  uint16(744),
	6430:  uint16(1),
	6431:  uint16(sym__concat_list),
	6432:  uint16(132),
	6433:  uint16(1),
	6434:  uint16(aux_sym_variable_expansion_repeat1),
	6435:  uint16(497),
	6436:  uint16(18),
	6437:  uint16(anon_sym_SEMI),
	6438:  uint16(anon_sym_AMP),
	6439:  uint16(anon_sym_LF),
	6440:  uint16(anon_sym_CR),
	6441:  uint16(anon_sym_CR_LF),
	6442:  uint16(anon_sym_DOLLAR),
	6443:  uint16(anon_sym_LPAREN),
	6444:  uint16(sym_integer),
	6445:  uint16(sym_float),
	6446:  uint16(anon_sym_LBRACK),
	6447:  uint16(anon_sym_RBRACK),
	6448:  uint16(anon_sym_LBRACE),
	6449:  uint16(anon_sym_DQUOTE),
	6450:  uint16(anon_sym_SQUOTE),
	6451:  uint16(sym_escape_sequence),
	6452:  uint16(sym_home_dir_expansion),
	6453:  uint16(sym_glob),
	6454:  uint16(sym_word),
	6455:  uint16(4),
	6456:  uint16(3),
	6457:  uint16(1),
	6458:  uint16(sym_comment),
	6459:  uint16(136),
	6460:  uint16(1),
	6461:  uint16(aux_sym_variable_expansion_repeat1),
	6462:  uint16(493),
	6463:  uint16(2),
	6464:  uint16(sym__concat),
	6465:  uint16(sym__concat_list),
	6466:  uint16(491),
	6467:  uint16(18),
	6468:  uint16(anon_sym_SEMI),
	6469:  uint16(anon_sym_AMP),
	6470:  uint16(anon_sym_LF),
	6471:  uint16(anon_sym_CR),
	6472:  uint16(anon_sym_CR_LF),
	6473:  uint16(anon_sym_DOLLAR),
	6474:  uint16(anon_sym_LPAREN),
	6475:  uint16(sym_integer),
	6476:  uint16(sym_float),
	6477:  uint16(anon_sym_LBRACK),
	6478:  uint16(anon_sym_RBRACK),
	6479:  uint16(anon_sym_LBRACE),
	6480:  uint16(anon_sym_DQUOTE),
	6481:  uint16(anon_sym_SQUOTE),
	6482:  uint16(sym_escape_sequence),
	6483:  uint16(sym_home_dir_expansion),
	6484:  uint16(sym_glob),
	6485:  uint16(sym_word),
	6486:  uint16(4),
	6487:  uint16(3),
	6488:  uint16(1),
	6489:  uint16(sym_comment),
	6490:  uint16(134),
	6491:  uint16(1),
	6492:  uint16(aux_sym_variable_expansion_repeat1),
	6493:  uint16(499),
	6494:  uint16(2),
	6495:  uint16(sym__concat),
	6496:  uint16(sym__concat_list),
	6497:  uint16(497),
	6498:  uint16(18),
	6499:  uint16(anon_sym_SEMI),
	6500:  uint16(anon_sym_AMP),
	6501:  uint16(anon_sym_LF),
	6502:  uint16(anon_sym_CR),
	6503:  uint16(anon_sym_CR_LF),
	6504:  uint16(anon_sym_DOLLAR),
	6505:  uint16(anon_sym_LPAREN),
	6506:  uint16(sym_integer),
	6507:  uint16(sym_float),
	6508:  uint16(anon_sym_LBRACK),
	6509:  uint16(anon_sym_RBRACK),
	6510:  uint16(anon_sym_LBRACE),
	6511:  uint16(anon_sym_DQUOTE),
	6512:  uint16(anon_sym_SQUOTE),
	6513:  uint16(sym_escape_sequence),
	6514:  uint16(sym_home_dir_expansion),
	6515:  uint16(sym_glob),
	6516:  uint16(sym_word),
	6517:  uint16(5),
	6518:  uint16(3),
	6519:  uint16(1),
	6520:  uint16(sym_comment),
	6521:  uint16(503),
	6522:  uint16(1),
	6523:  uint16(sym__concat),
	6524:  uint16(746),
	6525:  uint16(1),
	6526:  uint16(sym__concat_list),
	6527:  uint16(136),
	6528:  uint16(1),
	6529:  uint16(aux_sym_variable_expansion_repeat1),
	6530:  uint16(501),
	6531:  uint16(18),
	6532:  uint16(anon_sym_SEMI),
	6533:  uint16(anon_sym_AMP),
	6534:  uint16(anon_sym_LF),
	6535:  uint16(anon_sym_CR),
	6536:  uint16(anon_sym_CR_LF),
	6537:  uint16(anon_sym_DOLLAR),
	6538:  uint16(anon_sym_LPAREN),
	6539:  uint16(sym_integer),
	6540:  uint16(sym_float),
	6541:  uint16(anon_sym_LBRACK),
	6542:  uint16(anon_sym_RBRACK),
	6543:  uint16(anon_sym_LBRACE),
	6544:  uint16(anon_sym_DQUOTE),
	6545:  uint16(anon_sym_SQUOTE),
	6546:  uint16(sym_escape_sequence),
	6547:  uint16(sym_home_dir_expansion),
	6548:  uint16(sym_glob),
	6549:  uint16(sym_word),
	6550:  uint16(4),
	6551:  uint16(3),
	6552:  uint16(1),
	6553:  uint16(sym_comment),
	6554:  uint16(749),
	6555:  uint16(1),
	6556:  uint16(sym__concat),
	6557:  uint16(139),
	6558:  uint16(1),
	6559:  uint16(aux_sym_concatenation_repeat1),
	6560:  uint16(508),
	6561:  uint16(18),
	6562:  uint16(anon_sym_SEMI),
	6563:  uint16(anon_sym_AMP),
	6564:  uint16(anon_sym_LF),
	6565:  uint16(anon_sym_CR),
	6566:  uint16(anon_sym_CR_LF),
	6567:  uint16(anon_sym_DOLLAR),
	6568:  uint16(anon_sym_LPAREN),
	6569:  uint16(sym_integer),
	6570:  uint16(sym_float),
	6571:  uint16(anon_sym_LBRACK),
	6572:  uint16(anon_sym_RBRACK),
	6573:  uint16(anon_sym_LBRACE),
	6574:  uint16(anon_sym_DQUOTE),
	6575:  uint16(anon_sym_SQUOTE),
	6576:  uint16(sym_escape_sequence),
	6577:  uint16(sym_home_dir_expansion),
	6578:  uint16(sym_glob),
	6579:  uint16(sym_word),
	6580:  uint16(3),
	6581:  uint16(3),
	6582:  uint16(1),
	6583:  uint16(sym_comment),
	6584:  uint16(574),
	6585:  uint16(2),
	6586:  uint16(sym__concat),
	6587:  uint16(sym__concat_list),
	6588:  uint16(572),
	6589:  uint16(18),
	6590:  uint16(anon_sym_SEMI),
	6591:  uint16(anon_sym_AMP),
	6592:  uint16(anon_sym_LF),
	6593:  uint16(anon_sym_CR),
	6594:  uint16(anon_sym_CR_LF),
	6595:  uint16(anon_sym_DOLLAR),
	6596:  uint16(anon_sym_LPAREN),
	6597:  uint16(sym_integer),
	6598:  uint16(sym_float),
	6599:  uint16(anon_sym_LBRACK),
	6600:  uint16(anon_sym_RBRACK),
	6601:  uint16(anon_sym_LBRACE),
	6602:  uint16(anon_sym_DQUOTE),
	6603:  uint16(anon_sym_SQUOTE),
	6604:  uint16(sym_escape_sequence),
	6605:  uint16(sym_home_dir_expansion),
	6606:  uint16(sym_glob),
	6607:  uint16(sym_word),
	6608:  uint16(4),
	6609:  uint16(3),
	6610:  uint16(1),
	6611:  uint16(sym_comment),
	6612:  uint16(749),
	6613:  uint16(1),
	6614:  uint16(sym__concat),
	6615:  uint16(140),
	6616:  uint16(1),
	6617:  uint16(aux_sym_concatenation_repeat1),
	6618:  uint16(601),
	6619:  uint16(18),
	6620:  uint16(anon_sym_SEMI),
	6621:  uint16(anon_sym_AMP),
	6622:  uint16(anon_sym_LF),
	6623:  uint16(anon_sym_CR),
	6624:  uint16(anon_sym_CR_LF),
	6625:  uint16(anon_sym_DOLLAR),
	6626:  uint16(anon_sym_LPAREN),
	6627:  uint16(sym_integer),
	6628:  uint16(sym_float),
	6629:  uint16(anon_sym_LBRACK),
	6630:  uint16(anon_sym_RBRACK),
	6631:  uint16(anon_sym_LBRACE),
	6632:  uint16(anon_sym_DQUOTE),
	6633:  uint16(anon_sym_SQUOTE),
	6634:  uint16(sym_escape_sequence),
	6635:  uint16(sym_home_dir_expansion),
	6636:  uint16(sym_glob),
	6637:  uint16(sym_word),
	6638:  uint16(4),
	6639:  uint16(3),
	6640:  uint16(1),
	6641:  uint16(sym_comment),
	6642:  uint16(751),
	6643:  uint16(1),
	6644:  uint16(sym__concat),
	6645:  uint16(140),
	6646:  uint16(1),
	6647:  uint16(aux_sym_concatenation_repeat1),
	6648:  uint16(544),
	6649:  uint16(18),
	6650:  uint16(anon_sym_SEMI),
	6651:  uint16(anon_sym_AMP),
	6652:  uint16(anon_sym_LF),
	6653:  uint16(anon_sym_CR),
	6654:  uint16(anon_sym_CR_LF),
	6655:  uint16(anon_sym_DOLLAR),
	6656:  uint16(anon_sym_LPAREN),
	6657:  uint16(sym_integer),
	6658:  uint16(sym_float),
	6659:  uint16(anon_sym_LBRACK),
	6660:  uint16(anon_sym_RBRACK),
	6661:  uint16(anon_sym_LBRACE),
	6662:  uint16(anon_sym_DQUOTE),
	6663:  uint16(anon_sym_SQUOTE),
	6664:  uint16(sym_escape_sequence),
	6665:  uint16(sym_home_dir_expansion),
	6666:  uint16(sym_glob),
	6667:  uint16(sym_word),
	6668:  uint16(3),
	6669:  uint16(3),
	6670:  uint16(1),
	6671:  uint16(sym_comment),
	6672:  uint16(503),
	6673:  uint16(2),
	6674:  uint16(sym__concat),
	6675:  uint16(sym__concat_list),
	6676:  uint16(501),
	6677:  uint16(18),
	6678:  uint16(anon_sym_SEMI),
	6679:  uint16(anon_sym_AMP),
	6680:  uint16(anon_sym_LF),
	6681:  uint16(anon_sym_CR),
	6682:  uint16(anon_sym_CR_LF),
	6683:  uint16(anon_sym_DOLLAR),
	6684:  uint16(anon_sym_LPAREN),
	6685:  uint16(sym_integer),
	6686:  uint16(sym_float),
	6687:  uint16(anon_sym_LBRACK),
	6688:  uint16(anon_sym_RBRACK),
	6689:  uint16(anon_sym_LBRACE),
	6690:  uint16(anon_sym_DQUOTE),
	6691:  uint16(anon_sym_SQUOTE),
	6692:  uint16(sym_escape_sequence),
	6693:  uint16(sym_home_dir_expansion),
	6694:  uint16(sym_glob),
	6695:  uint16(sym_word),
	6696:  uint16(4),
	6697:  uint16(3),
	6698:  uint16(1),
	6699:  uint16(sym_comment),
	6700:  uint16(749),
	6701:  uint16(1),
	6702:  uint16(sym__concat),
	6703:  uint16(139),
	6704:  uint16(1),
	6705:  uint16(aux_sym_concatenation_repeat1),
	6706:  uint16(538),
	6707:  uint16(18),
	6708:  uint16(anon_sym_SEMI),
	6709:  uint16(anon_sym_AMP),
	6710:  uint16(anon_sym_LF),
	6711:  uint16(anon_sym_CR),
	6712:  uint16(anon_sym_CR_LF),
	6713:  uint16(anon_sym_DOLLAR),
	6714:  uint16(anon_sym_LPAREN),
	6715:  uint16(sym_integer),
	6716:  uint16(sym_float),
	6717:  uint16(anon_sym_LBRACK),
	6718:  uint16(anon_sym_RBRACK),
	6719:  uint16(anon_sym_LBRACE),
	6720:  uint16(anon_sym_DQUOTE),
	6721:  uint16(anon_sym_SQUOTE),
	6722:  uint16(sym_escape_sequence),
	6723:  uint16(sym_home_dir_expansion),
	6724:  uint16(sym_glob),
	6725:  uint16(sym_word),
	6726:  uint16(3),
	6727:  uint16(3),
	6728:  uint16(1),
	6729:  uint16(sym_comment),
	6730:  uint16(542),
	6731:  uint16(2),
	6732:  uint16(sym__concat),
	6733:  uint16(sym__concat_list),
	6734:  uint16(540),
	6735:  uint16(18),
	6736:  uint16(anon_sym_SEMI),
	6737:  uint16(anon_sym_AMP),
	6738:  uint16(anon_sym_LF),
	6739:  uint16(anon_sym_CR),
	6740:  uint16(anon_sym_CR_LF),
	6741:  uint16(anon_sym_DOLLAR),
	6742:  uint16(anon_sym_LPAREN),
	6743:  uint16(sym_integer),
	6744:  uint16(sym_float),
	6745:  uint16(anon_sym_LBRACK),
	6746:  uint16(anon_sym_RBRACK),
	6747:  uint16(anon_sym_LBRACE),
	6748:  uint16(anon_sym_DQUOTE),
	6749:  uint16(anon_sym_SQUOTE),
	6750:  uint16(sym_escape_sequence),
	6751:  uint16(sym_home_dir_expansion),
	6752:  uint16(sym_glob),
	6753:  uint16(sym_word),
	6754:  uint16(3),
	6755:  uint16(3),
	6756:  uint16(1),
	6757:  uint16(sym_comment),
	6758:  uint16(646),
	6759:  uint16(1),
	6760:  uint16(sym__concat),
	6761:  uint16(644),
	6762:  uint16(18),
	6763:  uint16(anon_sym_SEMI),
	6764:  uint16(anon_sym_AMP),
	6765:  uint16(anon_sym_LF),
	6766:  uint16(anon_sym_CR),
	6767:  uint16(anon_sym_CR_LF),
	6768:  uint16(anon_sym_DOLLAR),
	6769:  uint16(anon_sym_LPAREN),
	6770:  uint16(sym_integer),
	6771:  uint16(sym_float),
	6772:  uint16(anon_sym_LBRACK),
	6773:  uint16(anon_sym_RBRACK),
	6774:  uint16(anon_sym_LBRACE),
	6775:  uint16(anon_sym_DQUOTE),
	6776:  uint16(anon_sym_SQUOTE),
	6777:  uint16(sym_escape_sequence),
	6778:  uint16(sym_home_dir_expansion),
	6779:  uint16(sym_glob),
	6780:  uint16(sym_word),
	6781:  uint16(3),
	6782:  uint16(3),
	6783:  uint16(1),
	6784:  uint16(sym_comment),
	6785:  uint16(638),
	6786:  uint16(1),
	6787:  uint16(sym__concat),
	6788:  uint16(636),
	6789:  uint16(18),
	6790:  uint16(anon_sym_SEMI),
	6791:  uint16(anon_sym_AMP),
	6792:  uint16(anon_sym_LF),
	6793:  uint16(anon_sym_CR),
	6794:  uint16(anon_sym_CR_LF),
	6795:  uint16(anon_sym_DOLLAR),
	6796:  uint16(anon_sym_LPAREN),
	6797:  uint16(sym_integer),
	6798:  uint16(sym_float),
	6799:  uint16(anon_sym_LBRACK),
	6800:  uint16(anon_sym_RBRACK),
	6801:  uint16(anon_sym_LBRACE),
	6802:  uint16(anon_sym_DQUOTE),
	6803:  uint16(anon_sym_SQUOTE),
	6804:  uint16(sym_escape_sequence),
	6805:  uint16(sym_home_dir_expansion),
	6806:  uint16(sym_glob),
	6807:  uint16(sym_word),
	6808:  uint16(3),
	6809:  uint16(3),
	6810:  uint16(1),
	6811:  uint16(sym_comment),
	6812:  uint16(660),
	6813:  uint16(1),
	6814:  uint16(sym__concat),
	6815:  uint16(658),
	6816:  uint16(18),
	6817:  uint16(anon_sym_SEMI),
	6818:  uint16(anon_sym_AMP),
	6819:  uint16(anon_sym_LF),
	6820:  uint16(anon_sym_CR),
	6821:  uint16(anon_sym_CR_LF),
	6822:  uint16(anon_sym_DOLLAR),
	6823:  uint16(anon_sym_LPAREN),
	6824:  uint16(sym_integer),
	6825:  uint16(sym_float),
	6826:  uint16(anon_sym_LBRACK),
	6827:  uint16(anon_sym_RBRACK),
	6828:  uint16(anon_sym_LBRACE),
	6829:  uint16(anon_sym_DQUOTE),
	6830:  uint16(anon_sym_SQUOTE),
	6831:  uint16(sym_escape_sequence),
	6832:  uint16(sym_home_dir_expansion),
	6833:  uint16(sym_glob),
	6834:  uint16(sym_word),
	6835:  uint16(3),
	6836:  uint16(3),
	6837:  uint16(1),
	6838:  uint16(sym_comment),
	6839:  uint16(672),
	6840:  uint16(1),
	6841:  uint16(sym__concat),
	6842:  uint16(670),
	6843:  uint16(18),
	6844:  uint16(anon_sym_SEMI),
	6845:  uint16(anon_sym_AMP),
	6846:  uint16(anon_sym_LF),
	6847:  uint16(anon_sym_CR),
	6848:  uint16(anon_sym_CR_LF),
	6849:  uint16(anon_sym_DOLLAR),
	6850:  uint16(anon_sym_LPAREN),
	6851:  uint16(sym_integer),
	6852:  uint16(sym_float),
	6853:  uint16(anon_sym_LBRACK),
	6854:  uint16(anon_sym_RBRACK),
	6855:  uint16(anon_sym_LBRACE),
	6856:  uint16(anon_sym_DQUOTE),
	6857:  uint16(anon_sym_SQUOTE),
	6858:  uint16(sym_escape_sequence),
	6859:  uint16(sym_home_dir_expansion),
	6860:  uint16(sym_glob),
	6861:  uint16(sym_word),
	6862:  uint16(3),
	6863:  uint16(3),
	6864:  uint16(1),
	6865:  uint16(sym_comment),
	6866:  uint16(668),
	6867:  uint16(1),
	6868:  uint16(sym__concat),
	6869:  uint16(666),
	6870:  uint16(18),
	6871:  uint16(anon_sym_SEMI),
	6872:  uint16(anon_sym_AMP),
	6873:  uint16(anon_sym_LF),
	6874:  uint16(anon_sym_CR),
	6875:  uint16(anon_sym_CR_LF),
	6876:  uint16(anon_sym_DOLLAR),
	6877:  uint16(anon_sym_LPAREN),
	6878:  uint16(sym_integer),
	6879:  uint16(sym_float),
	6880:  uint16(anon_sym_LBRACK),
	6881:  uint16(anon_sym_RBRACK),
	6882:  uint16(anon_sym_LBRACE),
	6883:  uint16(anon_sym_DQUOTE),
	6884:  uint16(anon_sym_SQUOTE),
	6885:  uint16(sym_escape_sequence),
	6886:  uint16(sym_home_dir_expansion),
	6887:  uint16(sym_glob),
	6888:  uint16(sym_word),
	6889:  uint16(3),
	6890:  uint16(3),
	6891:  uint16(1),
	6892:  uint16(sym_comment),
	6893:  uint16(650),
	6894:  uint16(1),
	6895:  uint16(sym__concat),
	6896:  uint16(648),
	6897:  uint16(18),
	6898:  uint16(anon_sym_SEMI),
	6899:  uint16(anon_sym_AMP),
	6900:  uint16(anon_sym_LF),
	6901:  uint16(anon_sym_CR),
	6902:  uint16(anon_sym_CR_LF),
	6903:  uint16(anon_sym_DOLLAR),
	6904:  uint16(anon_sym_LPAREN),
	6905:  uint16(sym_integer),
	6906:  uint16(sym_float),
	6907:  uint16(anon_sym_LBRACK),
	6908:  uint16(anon_sym_RBRACK),
	6909:  uint16(anon_sym_LBRACE),
	6910:  uint16(anon_sym_DQUOTE),
	6911:  uint16(anon_sym_SQUOTE),
	6912:  uint16(sym_escape_sequence),
	6913:  uint16(sym_home_dir_expansion),
	6914:  uint16(sym_glob),
	6915:  uint16(sym_word),
	6916:  uint16(3),
	6917:  uint16(3),
	6918:  uint16(1),
	6919:  uint16(sym_comment),
	6920:  uint16(680),
	6921:  uint16(1),
	6922:  uint16(sym__concat),
	6923:  uint16(678),
	6924:  uint16(18),
	6925:  uint16(anon_sym_SEMI),
	6926:  uint16(anon_sym_AMP),
	6927:  uint16(anon_sym_LF),
	6928:  uint16(anon_sym_CR),
	6929:  uint16(anon_sym_CR_LF),
	6930:  uint16(anon_sym_DOLLAR),
	6931:  uint16(anon_sym_LPAREN),
	6932:  uint16(sym_integer),
	6933:  uint16(sym_float),
	6934:  uint16(anon_sym_LBRACK),
	6935:  uint16(anon_sym_RBRACK),
	6936:  uint16(anon_sym_LBRACE),
	6937:  uint16(anon_sym_DQUOTE),
	6938:  uint16(anon_sym_SQUOTE),
	6939:  uint16(sym_escape_sequence),
	6940:  uint16(sym_home_dir_expansion),
	6941:  uint16(sym_glob),
	6942:  uint16(sym_word),
	6943:  uint16(3),
	6944:  uint16(3),
	6945:  uint16(1),
	6946:  uint16(sym_comment),
	6947:  uint16(684),
	6948:  uint16(1),
	6949:  uint16(sym__concat),
	6950:  uint16(682),
	6951:  uint16(18),
	6952:  uint16(anon_sym_SEMI),
	6953:  uint16(anon_sym_AMP),
	6954:  uint16(anon_sym_LF),
	6955:  uint16(anon_sym_CR),
	6956:  uint16(anon_sym_CR_LF),
	6957:  uint16(anon_sym_DOLLAR),
	6958:  uint16(anon_sym_LPAREN),
	6959:  uint16(sym_integer),
	6960:  uint16(sym_float),
	6961:  uint16(anon_sym_LBRACK),
	6962:  uint16(anon_sym_RBRACK),
	6963:  uint16(anon_sym_LBRACE),
	6964:  uint16(anon_sym_DQUOTE),
	6965:  uint16(anon_sym_SQUOTE),
	6966:  uint16(sym_escape_sequence),
	6967:  uint16(sym_home_dir_expansion),
	6968:  uint16(sym_glob),
	6969:  uint16(sym_word),
	6970:  uint16(3),
	6971:  uint16(3),
	6972:  uint16(1),
	6973:  uint16(sym_comment),
	6974:  uint16(688),
	6975:  uint16(1),
	6976:  uint16(sym__concat),
	6977:  uint16(686),
	6978:  uint16(18),
	6979:  uint16(anon_sym_SEMI),
	6980:  uint16(anon_sym_AMP),
	6981:  uint16(anon_sym_LF),
	6982:  uint16(anon_sym_CR),
	6983:  uint16(anon_sym_CR_LF),
	6984:  uint16(anon_sym_DOLLAR),
	6985:  uint16(anon_sym_LPAREN),
	6986:  uint16(sym_integer),
	6987:  uint16(sym_float),
	6988:  uint16(anon_sym_LBRACK),
	6989:  uint16(anon_sym_RBRACK),
	6990:  uint16(anon_sym_LBRACE),
	6991:  uint16(anon_sym_DQUOTE),
	6992:  uint16(anon_sym_SQUOTE),
	6993:  uint16(sym_escape_sequence),
	6994:  uint16(sym_home_dir_expansion),
	6995:  uint16(sym_glob),
	6996:  uint16(sym_word),
	6997:  uint16(3),
	6998:  uint16(3),
	6999:  uint16(1),
	7000:  uint16(sym_comment),
	7001:  uint16(676),
	7002:  uint16(1),
	7003:  uint16(sym__concat),
	7004:  uint16(674),
	7005:  uint16(18),
	7006:  uint16(anon_sym_SEMI),
	7007:  uint16(anon_sym_AMP),
	7008:  uint16(anon_sym_LF),
	7009:  uint16(anon_sym_CR),
	7010:  uint16(anon_sym_CR_LF),
	7011:  uint16(anon_sym_DOLLAR),
	7012:  uint16(anon_sym_LPAREN),
	7013:  uint16(sym_integer),
	7014:  uint16(sym_float),
	7015:  uint16(anon_sym_LBRACK),
	7016:  uint16(anon_sym_RBRACK),
	7017:  uint16(anon_sym_LBRACE),
	7018:  uint16(anon_sym_DQUOTE),
	7019:  uint16(anon_sym_SQUOTE),
	7020:  uint16(sym_escape_sequence),
	7021:  uint16(sym_home_dir_expansion),
	7022:  uint16(sym_glob),
	7023:  uint16(sym_word),
	7024:  uint16(3),
	7025:  uint16(3),
	7026:  uint16(1),
	7027:  uint16(sym_comment),
	7028:  uint16(664),
	7029:  uint16(1),
	7030:  uint16(sym__concat),
	7031:  uint16(662),
	7032:  uint16(18),
	7033:  uint16(anon_sym_SEMI),
	7034:  uint16(anon_sym_AMP),
	7035:  uint16(anon_sym_LF),
	7036:  uint16(anon_sym_CR),
	7037:  uint16(anon_sym_CR_LF),
	7038:  uint16(anon_sym_DOLLAR),
	7039:  uint16(anon_sym_LPAREN),
	7040:  uint16(sym_integer),
	7041:  uint16(sym_float),
	7042:  uint16(anon_sym_LBRACK),
	7043:  uint16(anon_sym_RBRACK),
	7044:  uint16(anon_sym_LBRACE),
	7045:  uint16(anon_sym_DQUOTE),
	7046:  uint16(anon_sym_SQUOTE),
	7047:  uint16(sym_escape_sequence),
	7048:  uint16(sym_home_dir_expansion),
	7049:  uint16(sym_glob),
	7050:  uint16(sym_word),
	7051:  uint16(3),
	7052:  uint16(3),
	7053:  uint16(1),
	7054:  uint16(sym_comment),
	7055:  uint16(642),
	7056:  uint16(1),
	7057:  uint16(sym__concat),
	7058:  uint16(640),
	7059:  uint16(18),
	7060:  uint16(anon_sym_SEMI),
	7061:  uint16(anon_sym_AMP),
	7062:  uint16(anon_sym_LF),
	7063:  uint16(anon_sym_CR),
	7064:  uint16(anon_sym_CR_LF),
	7065:  uint16(anon_sym_DOLLAR),
	7066:  uint16(anon_sym_LPAREN),
	7067:  uint16(sym_integer),
	7068:  uint16(sym_float),
	7069:  uint16(anon_sym_LBRACK),
	7070:  uint16(anon_sym_RBRACK),
	7071:  uint16(anon_sym_LBRACE),
	7072:  uint16(anon_sym_DQUOTE),
	7073:  uint16(anon_sym_SQUOTE),
	7074:  uint16(sym_escape_sequence),
	7075:  uint16(sym_home_dir_expansion),
	7076:  uint16(sym_glob),
	7077:  uint16(sym_word),
	7078:  uint16(3),
	7079:  uint16(3),
	7080:  uint16(1),
	7081:  uint16(sym_comment),
	7082:  uint16(656),
	7083:  uint16(1),
	7084:  uint16(sym__concat),
	7085:  uint16(544),
	7086:  uint16(18),
	7087:  uint16(anon_sym_SEMI),
	7088:  uint16(anon_sym_AMP),
	7089:  uint16(anon_sym_LF),
	7090:  uint16(anon_sym_CR),
	7091:  uint16(anon_sym_CR_LF),
	7092:  uint16(anon_sym_DOLLAR),
	7093:  uint16(anon_sym_LPAREN),
	7094:  uint16(sym_integer),
	7095:  uint16(sym_float),
	7096:  uint16(anon_sym_LBRACK),
	7097:  uint16(anon_sym_RBRACK),
	7098:  uint16(anon_sym_LBRACE),
	7099:  uint16(anon_sym_DQUOTE),
	7100:  uint16(anon_sym_SQUOTE),
	7101:  uint16(sym_escape_sequence),
	7102:  uint16(sym_home_dir_expansion),
	7103:  uint16(sym_glob),
	7104:  uint16(sym_word),
	7105:  uint16(3),
	7106:  uint16(3),
	7107:  uint16(1),
	7108:  uint16(sym_comment),
	7109:  uint16(654),
	7110:  uint16(1),
	7111:  uint16(sym__concat),
	7112:  uint16(652),
	7113:  uint16(18),
	7114:  uint16(anon_sym_SEMI),
	7115:  uint16(anon_sym_AMP),
	7116:  uint16(anon_sym_LF),
	7117:  uint16(anon_sym_CR),
	7118:  uint16(anon_sym_CR_LF),
	7119:  uint16(anon_sym_DOLLAR),
	7120:  uint16(anon_sym_LPAREN),
	7121:  uint16(sym_integer),
	7122:  uint16(sym_float),
	7123:  uint16(anon_sym_LBRACK),
	7124:  uint16(anon_sym_RBRACK),
	7125:  uint16(anon_sym_LBRACE),
	7126:  uint16(anon_sym_DQUOTE),
	7127:  uint16(anon_sym_SQUOTE),
	7128:  uint16(sym_escape_sequence),
	7129:  uint16(sym_home_dir_expansion),
	7130:  uint16(sym_glob),
	7131:  uint16(sym_word),
	7132:  uint16(2),
	7133:  uint16(3),
	7134:  uint16(1),
	7135:  uint16(sym_comment),
	7136:  uint16(754),
	7137:  uint16(18),
	7138:  uint16(anon_sym_SEMI),
	7139:  uint16(anon_sym_AMP),
	7140:  uint16(anon_sym_LF),
	7141:  uint16(anon_sym_CR),
	7142:  uint16(anon_sym_CR_LF),
	7143:  uint16(anon_sym_DOLLAR),
	7144:  uint16(anon_sym_LPAREN),
	7145:  uint16(sym_integer),
	7146:  uint16(sym_float),
	7147:  uint16(anon_sym_LBRACK),
	7148:  uint16(anon_sym_RBRACK),
	7149:  uint16(anon_sym_LBRACE),
	7150:  uint16(anon_sym_DQUOTE),
	7151:  uint16(anon_sym_SQUOTE),
	7152:  uint16(sym_escape_sequence),
	7153:  uint16(sym_home_dir_expansion),
	7154:  uint16(sym_glob),
	7155:  uint16(sym_word),
	7156:  uint16(9),
	7157:  uint16(3),
	7158:  uint16(1),
	7159:  uint16(sym_comment),
	7160:  uint16(696),
	7161:  uint16(1),
	7162:  uint16(anon_sym_DOLLAR),
	7163:  uint16(698),
	7164:  uint16(1),
	7165:  uint16(anon_sym_LPAREN),
	7166:  uint16(704),
	7167:  uint16(1),
	7168:  uint16(anon_sym_LBRACE),
	7169:  uint16(708),
	7170:  uint16(1),
	7171:  uint16(anon_sym_DQUOTE),
	7172:  uint16(710),
	7173:  uint16(1),
	7174:  uint16(anon_sym_SQUOTE),
	7175:  uint16(388),
	7176:  uint16(2),
	7177:  uint16(sym__command_substitution_dollar),
	7178:  uint16(sym__command_substitution_inner),
	7179:  uint16(756),
	7180:  uint16(5),
	7181:  uint16(sym_integer),
	7182:  uint16(sym_float),
	7183:  uint16(sym_escape_sequence),
	7184:  uint16(sym_glob),
	7185:  uint16(sym_brace_word),
	7186:  uint16(407),
	7187:  uint16(6),
	7188:  uint16(sym_command_substitution),
	7189:  uint16(sym_variable_expansion),
	7190:  uint16(sym_brace_expansion),
	7191:  uint16(sym_double_quote_string),
	7192:  uint16(sym_single_quote_string),
	7193:  uint16(sym__base_brace_expression),
	7194:  uint16(2),
	7195:  uint16(3),
	7196:  uint16(1),
	7197:  uint16(sym_comment),
	7198:  uint16(758),
	7199:  uint16(18),
	7200:  uint16(anon_sym_SEMI),
	7201:  uint16(anon_sym_AMP),
	7202:  uint16(anon_sym_LF),
	7203:  uint16(anon_sym_CR),
	7204:  uint16(anon_sym_CR_LF),
	7205:  uint16(anon_sym_DOLLAR),
	7206:  uint16(anon_sym_LPAREN),
	7207:  uint16(sym_integer),
	7208:  uint16(sym_float),
	7209:  uint16(anon_sym_LBRACK),
	7210:  uint16(anon_sym_RBRACK),
	7211:  uint16(anon_sym_LBRACE),
	7212:  uint16(anon_sym_DQUOTE),
	7213:  uint16(anon_sym_SQUOTE),
	7214:  uint16(sym_escape_sequence),
	7215:  uint16(sym_home_dir_expansion),
	7216:  uint16(sym_glob),
	7217:  uint16(sym_word),
	7218:  uint16(5),
	7219:  uint16(3),
	7220:  uint16(1),
	7221:  uint16(sym_comment),
	7222:  uint16(503),
	7223:  uint16(1),
	7224:  uint16(sym__concat),
	7225:  uint16(760),
	7226:  uint16(1),
	7227:  uint16(sym__concat_list),
	7228:  uint16(161),
	7229:  uint16(1),
	7230:  uint16(aux_sym_variable_expansion_repeat1),
	7231:  uint16(501),
	7232:  uint16(14),
	7233:  uint16(anon_sym_PIPE_PIPE),
	7234:  uint16(anon_sym_AMP_AMP),
	7235:  uint16(anon_sym_AMP_PIPE),
	7236:  uint16(anon_sym_2_GT_PIPE),
	7237:  uint16(anon_sym_PIPE),
	7238:  uint16(anon_sym_SEMI),
	7239:  uint16(anon_sym_AMP),
	7240:  uint16(anon_sym_LF),
	7241:  uint16(anon_sym_CR),
	7242:  uint16(anon_sym_CR_LF),
	7243:  uint16(anon_sym_RPAREN),
	7244:  uint16(anon_sym_RBRACE),
	7245:  uint16(sym_stream_redirect),
	7246:  uint16(sym_direction),
	7247:  uint16(5),
	7248:  uint16(3),
	7249:  uint16(1),
	7250:  uint16(sym_comment),
	7251:  uint16(499),
	7252:  uint16(1),
	7253:  uint16(sym__concat),
	7254:  uint16(763),
	7255:  uint16(1),
	7256:  uint16(sym__concat_list),
	7257:  uint16(163),
	7258:  uint16(1),
	7259:  uint16(aux_sym_variable_expansion_repeat1),
	7260:  uint16(497),
	7261:  uint16(14),
	7262:  uint16(anon_sym_PIPE_PIPE),
	7263:  uint16(anon_sym_AMP_AMP),
	7264:  uint16(anon_sym_AMP_PIPE),
	7265:  uint16(anon_sym_2_GT_PIPE),
	7266:  uint16(anon_sym_PIPE),
	7267:  uint16(anon_sym_SEMI),
	7268:  uint16(anon_sym_AMP),
	7269:  uint16(anon_sym_LF),
	7270:  uint16(anon_sym_CR),
	7271:  uint16(anon_sym_CR_LF),
	7272:  uint16(anon_sym_RPAREN),
	7273:  uint16(anon_sym_RBRACE),
	7274:  uint16(sym_stream_redirect),
	7275:  uint16(sym_direction),
	7276:  uint16(5),
	7277:  uint16(3),
	7278:  uint16(1),
	7279:  uint16(sym_comment),
	7280:  uint16(493),
	7281:  uint16(1),
	7282:  uint16(sym__concat),
	7283:  uint16(763),
	7284:  uint16(1),
	7285:  uint16(sym__concat_list),
	7286:  uint16(161),
	7287:  uint16(1),
	7288:  uint16(aux_sym_variable_expansion_repeat1),
	7289:  uint16(491),
	7290:  uint16(14),
	7291:  uint16(anon_sym_PIPE_PIPE),
	7292:  uint16(anon_sym_AMP_AMP),
	7293:  uint16(anon_sym_AMP_PIPE),
	7294:  uint16(anon_sym_2_GT_PIPE),
	7295:  uint16(anon_sym_PIPE),
	7296:  uint16(anon_sym_SEMI),
	7297:  uint16(anon_sym_AMP),
	7298:  uint16(anon_sym_LF),
	7299:  uint16(anon_sym_CR),
	7300:  uint16(anon_sym_CR_LF),
	7301:  uint16(anon_sym_RPAREN),
	7302:  uint16(anon_sym_RBRACE),
	7303:  uint16(sym_stream_redirect),
	7304:  uint16(sym_direction),
	7305:  uint16(4),
	7306:  uint16(3),
	7307:  uint16(1),
	7308:  uint16(sym_comment),
	7309:  uint16(161),
	7310:  uint16(1),
	7311:  uint16(aux_sym_variable_expansion_repeat1),
	7312:  uint16(493),
	7313:  uint16(2),
	7314:  uint16(sym__concat),
	7315:  uint16(sym__concat_list),
	7316:  uint16(491),
	7317:  uint16(14),
	7318:  uint16(anon_sym_PIPE_PIPE),
	7319:  uint16(anon_sym_AMP_AMP),
	7320:  uint16(anon_sym_AMP_PIPE),
	7321:  uint16(anon_sym_2_GT_PIPE),
	7322:  uint16(anon_sym_PIPE),
	7323:  uint16(anon_sym_SEMI),
	7324:  uint16(anon_sym_AMP),
	7325:  uint16(anon_sym_LF),
	7326:  uint16(anon_sym_CR),
	7327:  uint16(anon_sym_CR_LF),
	7328:  uint16(anon_sym_RPAREN),
	7329:  uint16(anon_sym_RBRACE),
	7330:  uint16(sym_stream_redirect),
	7331:  uint16(sym_direction),
	7332:  uint16(4),
	7333:  uint16(3),
	7334:  uint16(1),
	7335:  uint16(sym_comment),
	7336:  uint16(164),
	7337:  uint16(1),
	7338:  uint16(aux_sym_variable_expansion_repeat1),
	7339:  uint16(499),
	7340:  uint16(2),
	7341:  uint16(sym__concat),
	7342:  uint16(sym__concat_list),
	7343:  uint16(497),
	7344:  uint16(14),
	7345:  uint16(anon_sym_PIPE_PIPE),
	7346:  uint16(anon_sym_AMP_AMP),
	7347:  uint16(anon_sym_AMP_PIPE),
	7348:  uint16(anon_sym_2_GT_PIPE),
	7349:  uint16(anon_sym_PIPE),
	7350:  uint16(anon_sym_SEMI),
	7351:  uint16(anon_sym_AMP),
	7352:  uint16(anon_sym_LF),
	7353:  uint16(anon_sym_CR),
	7354:  uint16(anon_sym_CR_LF),
	7355:  uint16(anon_sym_RPAREN),
	7356:  uint16(anon_sym_RBRACE),
	7357:  uint16(sym_stream_redirect),
	7358:  uint16(sym_direction),
	7359:  uint16(12),
	7360:  uint16(765),
	7361:  uint16(1),
	7362:  uint16(anon_sym_DOLLAR),
	7363:  uint16(767),
	7364:  uint16(1),
	7365:  uint16(anon_sym_LPAREN),
	7366:  uint16(769),
	7367:  uint16(1),
	7368:  uint16(sym_integer),
	7369:  uint16(771),
	7370:  uint16(1),
	7371:  uint16(sym_comment),
	7372:  uint16(773),
	7373:  uint16(1),
	7374:  uint16(anon_sym_DOT_DOT),
	7375:  uint16(775),
	7376:  uint16(1),
	7377:  uint16(anon_sym_RBRACK),
	7378:  uint16(777),
	7379:  uint16(1),
	7380:  uint16(anon_sym_DQUOTE),
	7381:  uint16(779),
	7382:  uint16(1),
	7383:  uint16(anon_sym_SQUOTE),
	7384:  uint16(284),
	7385:  uint16(1),
	7386:  uint16(sym_index),
	7387:  uint16(185),
	7388:  uint16(2),
	7389:  uint16(sym_range),
	7390:  uint16(aux_sym_list_element_access_repeat1),
	7391:  uint16(297),
	7392:  uint16(2),
	7393:  uint16(sym__command_substitution_dollar),
	7394:  uint16(sym__command_substitution_inner),
	7395:  uint16(285),
	7396:  uint16(4),
	7397:  uint16(sym_command_substitution),
	7398:  uint16(sym_variable_expansion),
	7399:  uint16(sym_double_quote_string),
	7400:  uint16(sym_single_quote_string),
	7401:  uint16(12),
	7402:  uint16(765),
	7403:  uint16(1),
	7404:  uint16(anon_sym_DOLLAR),
	7405:  uint16(767),
	7406:  uint16(1),
	7407:  uint16(anon_sym_LPAREN),
	7408:  uint16(769),
	7409:  uint16(1),
	7410:  uint16(sym_integer),
	7411:  uint16(771),
	7412:  uint16(1),
	7413:  uint16(sym_comment),
	7414:  uint16(773),
	7415:  uint16(1),
	7416:  uint16(anon_sym_DOT_DOT),
	7417:  uint16(777),
	7418:  uint16(1),
	7419:  uint16(anon_sym_DQUOTE),
	7420:  uint16(779),
	7421:  uint16(1),
	7422:  uint16(anon_sym_SQUOTE),
	7423:  uint16(781),
	7424:  uint16(1),
	7425:  uint16(anon_sym_RBRACK),
	7426:  uint16(284),
	7427:  uint16(1),
	7428:  uint16(sym_index),
	7429:  uint16(185),
	7430:  uint16(2),
	7431:  uint16(sym_range),
	7432:  uint16(aux_sym_list_element_access_repeat1),
	7433:  uint16(297),
	7434:  uint16(2),
	7435:  uint16(sym__command_substitution_dollar),
	7436:  uint16(sym__command_substitution_inner),
	7437:  uint16(285),
	7438:  uint16(4),
	7439:  uint16(sym_command_substitution),
	7440:  uint16(sym_variable_expansion),
	7441:  uint16(sym_double_quote_string),
	7442:  uint16(sym_single_quote_string),
	7443:  uint16(12),
	7444:  uint16(765),
	7445:  uint16(1),
	7446:  uint16(anon_sym_DOLLAR),
	7447:  uint16(767),
	7448:  uint16(1),
	7449:  uint16(anon_sym_LPAREN),
	7450:  uint16(769),
	7451:  uint16(1),
	7452:  uint16(sym_integer),
	7453:  uint16(771),
	7454:  uint16(1),
	7455:  uint16(sym_comment),
	7456:  uint16(773),
	7457:  uint16(1),
	7458:  uint16(anon_sym_DOT_DOT),
	7459:  uint16(777),
	7460:  uint16(1),
	7461:  uint16(anon_sym_DQUOTE),
	7462:  uint16(779),
	7463:  uint16(1),
	7464:  uint16(anon_sym_SQUOTE),
	7465:  uint16(783),
	7466:  uint16(1),
	7467:  uint16(anon_sym_RBRACK),
	7468:  uint16(284),
	7469:  uint16(1),
	7470:  uint16(sym_index),
	7471:  uint16(185),
	7472:  uint16(2),
	7473:  uint16(sym_range),
	7474:  uint16(aux_sym_list_element_access_repeat1),
	7475:  uint16(297),
	7476:  uint16(2),
	7477:  uint16(sym__command_substitution_dollar),
	7478:  uint16(sym__command_substitution_inner),
	7479:  uint16(285),
	7480:  uint16(4),
	7481:  uint16(sym_command_substitution),
	7482:  uint16(sym_variable_expansion),
	7483:  uint16(sym_double_quote_string),
	7484:  uint16(sym_single_quote_string),
	7485:  uint16(12),
	7486:  uint16(765),
	7487:  uint16(1),
	7488:  uint16(anon_sym_DOLLAR),
	7489:  uint16(767),
	7490:  uint16(1),
	7491:  uint16(anon_sym_LPAREN),
	7492:  uint16(769),
	7493:  uint16(1),
	7494:  uint16(sym_integer),
	7495:  uint16(771),
	7496:  uint16(1),
	7497:  uint16(sym_comment),
	7498:  uint16(773),
	7499:  uint16(1),
	7500:  uint16(anon_sym_DOT_DOT),
	7501:  uint16(777),
	7502:  uint16(1),
	7503:  uint16(anon_sym_DQUOTE),
	7504:  uint16(779),
	7505:  uint16(1),
	7506:  uint16(anon_sym_SQUOTE),
	7507:  uint16(785),
	7508:  uint16(1),
	7509:  uint16(anon_sym_RBRACK),
	7510:  uint16(284),
	7511:  uint16(1),
	7512:  uint16(sym_index),
	7513:  uint16(173),
	7514:  uint16(2),
	7515:  uint16(sym_range),
	7516:  uint16(aux_sym_list_element_access_repeat1),
	7517:  uint16(297),
	7518:  uint16(2),
	7519:  uint16(sym__command_substitution_dollar),
	7520:  uint16(sym__command_substitution_inner),
	7521:  uint16(285),
	7522:  uint16(4),
	7523:  uint16(sym_command_substitution),
	7524:  uint16(sym_variable_expansion),
	7525:  uint16(sym_double_quote_string),
	7526:  uint16(sym_single_quote_string),
	7527:  uint16(12),
	7528:  uint16(765),
	7529:  uint16(1),
	7530:  uint16(anon_sym_DOLLAR),
	7531:  uint16(767),
	7532:  uint16(1),
	7533:  uint16(anon_sym_LPAREN),
	7534:  uint16(769),
	7535:  uint16(1),
	7536:  uint16(sym_integer),
	7537:  uint16(771),
	7538:  uint16(1),
	7539:  uint16(sym_comment),
	7540:  uint16(773),
	7541:  uint16(1),
	7542:  uint16(anon_sym_DOT_DOT),
	7543:  uint16(777),
	7544:  uint16(1),
	7545:  uint16(anon_sym_DQUOTE),
	7546:  uint16(779),
	7547:  uint16(1),
	7548:  uint16(anon_sym_SQUOTE),
	7549:  uint16(787),
	7550:  uint16(1),
	7551:  uint16(anon_sym_RBRACK),
	7552:  uint16(284),
	7553:  uint16(1),
	7554:  uint16(sym_index),
	7555:  uint16(168),
	7556:  uint16(2),
	7557:  uint16(sym_range),
	7558:  uint16(aux_sym_list_element_access_repeat1),
	7559:  uint16(297),
	7560:  uint16(2),
	7561:  uint16(sym__command_substitution_dollar),
	7562:  uint16(sym__command_substitution_inner),
	7563:  uint16(285),
	7564:  uint16(4),
	7565:  uint16(sym_command_substitution),
	7566:  uint16(sym_variable_expansion),
	7567:  uint16(sym_double_quote_string),
	7568:  uint16(sym_single_quote_string),
	7569:  uint16(12),
	7570:  uint16(765),
	7571:  uint16(1),
	7572:  uint16(anon_sym_DOLLAR),
	7573:  uint16(767),
	7574:  uint16(1),
	7575:  uint16(anon_sym_LPAREN),
	7576:  uint16(769),
	7577:  uint16(1),
	7578:  uint16(sym_integer),
	7579:  uint16(771),
	7580:  uint16(1),
	7581:  uint16(sym_comment),
	7582:  uint16(773),
	7583:  uint16(1),
	7584:  uint16(anon_sym_DOT_DOT),
	7585:  uint16(777),
	7586:  uint16(1),
	7587:  uint16(anon_sym_DQUOTE),
	7588:  uint16(779),
	7589:  uint16(1),
	7590:  uint16(anon_sym_SQUOTE),
	7591:  uint16(789),
	7592:  uint16(1),
	7593:  uint16(anon_sym_RBRACK),
	7594:  uint16(284),
	7595:  uint16(1),
	7596:  uint16(sym_index),
	7597:  uint16(185),
	7598:  uint16(2),
	7599:  uint16(sym_range),
	7600:  uint16(aux_sym_list_element_access_repeat1),
	7601:  uint16(297),
	7602:  uint16(2),
	7603:  uint16(sym__command_substitution_dollar),
	7604:  uint16(sym__command_substitution_inner),
	7605:  uint16(285),
	7606:  uint16(4),
	7607:  uint16(sym_command_substitution),
	7608:  uint16(sym_variable_expansion),
	7609:  uint16(sym_double_quote_string),
	7610:  uint16(sym_single_quote_string),
	7611:  uint16(12),
	7612:  uint16(765),
	7613:  uint16(1),
	7614:  uint16(anon_sym_DOLLAR),
	7615:  uint16(767),
	7616:  uint16(1),
	7617:  uint16(anon_sym_LPAREN),
	7618:  uint16(769),
	7619:  uint16(1),
	7620:  uint16(sym_integer),
	7621:  uint16(771),
	7622:  uint16(1),
	7623:  uint16(sym_comment),
	7624:  uint16(773),
	7625:  uint16(1),
	7626:  uint16(anon_sym_DOT_DOT),
	7627:  uint16(777),
	7628:  uint16(1),
	7629:  uint16(anon_sym_DQUOTE),
	7630:  uint16(779),
	7631:  uint16(1),
	7632:  uint16(anon_sym_SQUOTE),
	7633:  uint16(791),
	7634:  uint16(1),
	7635:  uint16(anon_sym_RBRACK),
	7636:  uint16(284),
	7637:  uint16(1),
	7638:  uint16(sym_index),
	7639:  uint16(171),
	7640:  uint16(2),
	7641:  uint16(sym_range),
	7642:  uint16(aux_sym_list_element_access_repeat1),
	7643:  uint16(297),
	7644:  uint16(2),
	7645:  uint16(sym__command_substitution_dollar),
	7646:  uint16(sym__command_substitution_inner),
	7647:  uint16(285),
	7648:  uint16(4),
	7649:  uint16(sym_command_substitution),
	7650:  uint16(sym_variable_expansion),
	7651:  uint16(sym_double_quote_string),
	7652:  uint16(sym_single_quote_string),
	7653:  uint16(12),
	7654:  uint16(765),
	7655:  uint16(1),
	7656:  uint16(anon_sym_DOLLAR),
	7657:  uint16(767),
	7658:  uint16(1),
	7659:  uint16(anon_sym_LPAREN),
	7660:  uint16(769),
	7661:  uint16(1),
	7662:  uint16(sym_integer),
	7663:  uint16(771),
	7664:  uint16(1),
	7665:  uint16(sym_comment),
	7666:  uint16(773),
	7667:  uint16(1),
	7668:  uint16(anon_sym_DOT_DOT),
	7669:  uint16(777),
	7670:  uint16(1),
	7671:  uint16(anon_sym_DQUOTE),
	7672:  uint16(779),
	7673:  uint16(1),
	7674:  uint16(anon_sym_SQUOTE),
	7675:  uint16(793),
	7676:  uint16(1),
	7677:  uint16(anon_sym_RBRACK),
	7678:  uint16(284),
	7679:  uint16(1),
	7680:  uint16(sym_index),
	7681:  uint16(185),
	7682:  uint16(2),
	7683:  uint16(sym_range),
	7684:  uint16(aux_sym_list_element_access_repeat1),
	7685:  uint16(297),
	7686:  uint16(2),
	7687:  uint16(sym__command_substitution_dollar),
	7688:  uint16(sym__command_substitution_inner),
	7689:  uint16(285),
	7690:  uint16(4),
	7691:  uint16(sym_command_substitution),
	7692:  uint16(sym_variable_expansion),
	7693:  uint16(sym_double_quote_string),
	7694:  uint16(sym_single_quote_string),
	7695:  uint16(12),
	7696:  uint16(765),
	7697:  uint16(1),
	7698:  uint16(anon_sym_DOLLAR),
	7699:  uint16(767),
	7700:  uint16(1),
	7701:  uint16(anon_sym_LPAREN),
	7702:  uint16(769),
	7703:  uint16(1),
	7704:  uint16(sym_integer),
	7705:  uint16(771),
	7706:  uint16(1),
	7707:  uint16(sym_comment),
	7708:  uint16(773),
	7709:  uint16(1),
	7710:  uint16(anon_sym_DOT_DOT),
	7711:  uint16(777),
	7712:  uint16(1),
	7713:  uint16(anon_sym_DQUOTE),
	7714:  uint16(779),
	7715:  uint16(1),
	7716:  uint16(anon_sym_SQUOTE),
	7717:  uint16(795),
	7718:  uint16(1),
	7719:  uint16(anon_sym_RBRACK),
	7720:  uint16(284),
	7721:  uint16(1),
	7722:  uint16(sym_index),
	7723:  uint16(167),
	7724:  uint16(2),
	7725:  uint16(sym_range),
	7726:  uint16(aux_sym_list_element_access_repeat1),
	7727:  uint16(297),
	7728:  uint16(2),
	7729:  uint16(sym__command_substitution_dollar),
	7730:  uint16(sym__command_substitution_inner),
	7731:  uint16(285),
	7732:  uint16(4),
	7733:  uint16(sym_command_substitution),
	7734:  uint16(sym_variable_expansion),
	7735:  uint16(sym_double_quote_string),
	7736:  uint16(sym_single_quote_string),
	7737:  uint16(4),
	7738:  uint16(3),
	7739:  uint16(1),
	7740:  uint16(sym_comment),
	7741:  uint16(797),
	7742:  uint16(1),
	7743:  uint16(sym__concat),
	7744:  uint16(176),
	7745:  uint16(1),
	7746:  uint16(aux_sym_concatenation_repeat1),
	7747:  uint16(601),
	7748:  uint16(14),
	7749:  uint16(anon_sym_PIPE_PIPE),
	7750:  uint16(anon_sym_AMP_AMP),
	7751:  uint16(anon_sym_AMP_PIPE),
	7752:  uint16(anon_sym_2_GT_PIPE),
	7753:  uint16(anon_sym_PIPE),
	7754:  uint16(anon_sym_SEMI),
	7755:  uint16(anon_sym_AMP),
	7756:  uint16(anon_sym_LF),
	7757:  uint16(anon_sym_CR),
	7758:  uint16(anon_sym_CR_LF),
	7759:  uint16(anon_sym_RPAREN),
	7760:  uint16(anon_sym_RBRACE),
	7761:  uint16(sym_stream_redirect),
	7762:  uint16(sym_direction),
	7763:  uint16(4),
	7764:  uint16(3),
	7765:  uint16(1),
	7766:  uint16(sym_comment),
	7767:  uint16(799),
	7768:  uint16(1),
	7769:  uint16(sym__concat),
	7770:  uint16(176),
	7771:  uint16(1),
	7772:  uint16(aux_sym_concatenation_repeat1),
	7773:  uint16(544),
	7774:  uint16(14),
	7775:  uint16(anon_sym_PIPE_PIPE),
	7776:  uint16(anon_sym_AMP_AMP),
	7777:  uint16(anon_sym_AMP_PIPE),
	7778:  uint16(anon_sym_2_GT_PIPE),
	7779:  uint16(anon_sym_PIPE),
	7780:  uint16(anon_sym_SEMI),
	7781:  uint16(anon_sym_AMP),
	7782:  uint16(anon_sym_LF),
	7783:  uint16(anon_sym_CR),
	7784:  uint16(anon_sym_CR_LF),
	7785:  uint16(anon_sym_RPAREN),
	7786:  uint16(anon_sym_RBRACE),
	7787:  uint16(sym_stream_redirect),
	7788:  uint16(sym_direction),
	7789:  uint16(12),
	7790:  uint16(765),
	7791:  uint16(1),
	7792:  uint16(anon_sym_DOLLAR),
	7793:  uint16(767),
	7794:  uint16(1),
	7795:  uint16(anon_sym_LPAREN),
	7796:  uint16(769),
	7797:  uint16(1),
	7798:  uint16(sym_integer),
	7799:  uint16(771),
	7800:  uint16(1),
	7801:  uint16(sym_comment),
	7802:  uint16(773),
	7803:  uint16(1),
	7804:  uint16(anon_sym_DOT_DOT),
	7805:  uint16(777),
	7806:  uint16(1),
	7807:  uint16(anon_sym_DQUOTE),
	7808:  uint16(779),
	7809:  uint16(1),
	7810:  uint16(anon_sym_SQUOTE),
	7811:  uint16(802),
	7812:  uint16(1),
	7813:  uint16(anon_sym_RBRACK),
	7814:  uint16(284),
	7815:  uint16(1),
	7816:  uint16(sym_index),
	7817:  uint16(185),
	7818:  uint16(2),
	7819:  uint16(sym_range),
	7820:  uint16(aux_sym_list_element_access_repeat1),
	7821:  uint16(297),
	7822:  uint16(2),
	7823:  uint16(sym__command_substitution_dollar),
	7824:  uint16(sym__command_substitution_inner),
	7825:  uint16(285),
	7826:  uint16(4),
	7827:  uint16(sym_command_substitution),
	7828:  uint16(sym_variable_expansion),
	7829:  uint16(sym_double_quote_string),
	7830:  uint16(sym_single_quote_string),
	7831:  uint16(4),
	7832:  uint16(3),
	7833:  uint16(1),
	7834:  uint16(sym_comment),
	7835:  uint16(797),
	7836:  uint16(1),
	7837:  uint16(sym__concat),
	7838:  uint16(175),
	7839:  uint16(1),
	7840:  uint16(aux_sym_concatenation_repeat1),
	7841:  uint16(538),
	7842:  uint16(14),
	7843:  uint16(anon_sym_PIPE_PIPE),
	7844:  uint16(anon_sym_AMP_AMP),
	7845:  uint16(anon_sym_AMP_PIPE),
	7846:  uint16(anon_sym_2_GT_PIPE),
	7847:  uint16(anon_sym_PIPE),
	7848:  uint16(anon_sym_SEMI),
	7849:  uint16(anon_sym_AMP),
	7850:  uint16(anon_sym_LF),
	7851:  uint16(anon_sym_CR),
	7852:  uint16(anon_sym_CR_LF),
	7853:  uint16(anon_sym_RPAREN),
	7854:  uint16(anon_sym_RBRACE),
	7855:  uint16(sym_stream_redirect),
	7856:  uint16(sym_direction),
	7857:  uint16(3),
	7858:  uint16(3),
	7859:  uint16(1),
	7860:  uint16(sym_comment),
	7861:  uint16(574),
	7862:  uint16(2),
	7863:  uint16(sym__concat),
	7864:  uint16(sym__concat_list),
	7865:  uint16(572),
	7866:  uint16(14),
	7867:  uint16(anon_sym_PIPE_PIPE),
	7868:  uint16(anon_sym_AMP_AMP),
	7869:  uint16(anon_sym_AMP_PIPE),
	7870:  uint16(anon_sym_2_GT_PIPE),
	7871:  uint16(anon_sym_PIPE),
	7872:  uint16(anon_sym_SEMI),
	7873:  uint16(anon_sym_AMP),
	7874:  uint16(anon_sym_LF),
	7875:  uint16(anon_sym_CR),
	7876:  uint16(anon_sym_CR_LF),
	7877:  uint16(anon_sym_RPAREN),
	7878:  uint16(anon_sym_RBRACE),
	7879:  uint16(sym_stream_redirect),
	7880:  uint16(sym_direction),
	7881:  uint16(12),
	7882:  uint16(765),
	7883:  uint16(1),
	7884:  uint16(anon_sym_DOLLAR),
	7885:  uint16(767),
	7886:  uint16(1),
	7887:  uint16(anon_sym_LPAREN),
	7888:  uint16(769),
	7889:  uint16(1),
	7890:  uint16(sym_integer),
	7891:  uint16(771),
	7892:  uint16(1),
	7893:  uint16(sym_comment),
	7894:  uint16(773),
	7895:  uint16(1),
	7896:  uint16(anon_sym_DOT_DOT),
	7897:  uint16(777),
	7898:  uint16(1),
	7899:  uint16(anon_sym_DQUOTE),
	7900:  uint16(779),
	7901:  uint16(1),
	7902:  uint16(anon_sym_SQUOTE),
	7903:  uint16(804),
	7904:  uint16(1),
	7905:  uint16(anon_sym_RBRACK),
	7906:  uint16(284),
	7907:  uint16(1),
	7908:  uint16(sym_index),
	7909:  uint16(185),
	7910:  uint16(2),
	7911:  uint16(sym_range),
	7912:  uint16(aux_sym_list_element_access_repeat1),
	7913:  uint16(297),
	7914:  uint16(2),
	7915:  uint16(sym__command_substitution_dollar),
	7916:  uint16(sym__command_substitution_inner),
	7917:  uint16(285),
	7918:  uint16(4),
	7919:  uint16(sym_command_substitution),
	7920:  uint16(sym_variable_expansion),
	7921:  uint16(sym_double_quote_string),
	7922:  uint16(sym_single_quote_string),
	7923:  uint16(3),
	7924:  uint16(3),
	7925:  uint16(1),
	7926:  uint16(sym_comment),
	7927:  uint16(542),
	7928:  uint16(2),
	7929:  uint16(sym__concat),
	7930:  uint16(sym__concat_list),
	7931:  uint16(540),
	7932:  uint16(14),
	7933:  uint16(anon_sym_PIPE_PIPE),
	7934:  uint16(anon_sym_AMP_AMP),
	7935:  uint16(anon_sym_AMP_PIPE),
	7936:  uint16(anon_sym_2_GT_PIPE),
	7937:  uint16(anon_sym_PIPE),
	7938:  uint16(anon_sym_SEMI),
	7939:  uint16(anon_sym_AMP),
	7940:  uint16(anon_sym_LF),
	7941:  uint16(anon_sym_CR),
	7942:  uint16(anon_sym_CR_LF),
	7943:  uint16(anon_sym_RPAREN),
	7944:  uint16(anon_sym_RBRACE),
	7945:  uint16(sym_stream_redirect),
	7946:  uint16(sym_direction),
	7947:  uint16(4),
	7948:  uint16(3),
	7949:  uint16(1),
	7950:  uint16(sym_comment),
	7951:  uint16(797),
	7952:  uint16(1),
	7953:  uint16(sym__concat),
	7954:  uint16(175),
	7955:  uint16(1),
	7956:  uint16(aux_sym_concatenation_repeat1),
	7957:  uint16(508),
	7958:  uint16(14),
	7959:  uint16(anon_sym_PIPE_PIPE),
	7960:  uint16(anon_sym_AMP_AMP),
	7961:  uint16(anon_sym_AMP_PIPE),
	7962:  uint16(anon_sym_2_GT_PIPE),
	7963:  uint16(anon_sym_PIPE),
	7964:  uint16(anon_sym_SEMI),
	7965:  uint16(anon_sym_AMP),
	7966:  uint16(anon_sym_LF),
	7967:  uint16(anon_sym_CR),
	7968:  uint16(anon_sym_CR_LF),
	7969:  uint16(anon_sym_RPAREN),
	7970:  uint16(anon_sym_RBRACE),
	7971:  uint16(sym_stream_redirect),
	7972:  uint16(sym_direction),
	7973:  uint16(12),
	7974:  uint16(765),
	7975:  uint16(1),
	7976:  uint16(anon_sym_DOLLAR),
	7977:  uint16(767),
	7978:  uint16(1),
	7979:  uint16(anon_sym_LPAREN),
	7980:  uint16(769),
	7981:  uint16(1),
	7982:  uint16(sym_integer),
	7983:  uint16(771),
	7984:  uint16(1),
	7985:  uint16(sym_comment),
	7986:  uint16(773),
	7987:  uint16(1),
	7988:  uint16(anon_sym_DOT_DOT),
	7989:  uint16(777),
	7990:  uint16(1),
	7991:  uint16(anon_sym_DQUOTE),
	7992:  uint16(779),
	7993:  uint16(1),
	7994:  uint16(anon_sym_SQUOTE),
	7995:  uint16(806),
	7996:  uint16(1),
	7997:  uint16(anon_sym_RBRACK),
	7998:  uint16(284),
	7999:  uint16(1),
	8000:  uint16(sym_index),
	8001:  uint16(166),
	8002:  uint16(2),
	8003:  uint16(sym_range),
	8004:  uint16(aux_sym_list_element_access_repeat1),
	8005:  uint16(297),
	8006:  uint16(2),
	8007:  uint16(sym__command_substitution_dollar),
	8008:  uint16(sym__command_substitution_inner),
	8009:  uint16(285),
	8010:  uint16(4),
	8011:  uint16(sym_command_substitution),
	8012:  uint16(sym_variable_expansion),
	8013:  uint16(sym_double_quote_string),
	8014:  uint16(sym_single_quote_string),
	8015:  uint16(12),
	8016:  uint16(765),
	8017:  uint16(1),
	8018:  uint16(anon_sym_DOLLAR),
	8019:  uint16(767),
	8020:  uint16(1),
	8021:  uint16(anon_sym_LPAREN),
	8022:  uint16(769),
	8023:  uint16(1),
	8024:  uint16(sym_integer),
	8025:  uint16(771),
	8026:  uint16(1),
	8027:  uint16(sym_comment),
	8028:  uint16(773),
	8029:  uint16(1),
	8030:  uint16(anon_sym_DOT_DOT),
	8031:  uint16(777),
	8032:  uint16(1),
	8033:  uint16(anon_sym_DQUOTE),
	8034:  uint16(779),
	8035:  uint16(1),
	8036:  uint16(anon_sym_SQUOTE),
	8037:  uint16(808),
	8038:  uint16(1),
	8039:  uint16(anon_sym_RBRACK),
	8040:  uint16(284),
	8041:  uint16(1),
	8042:  uint16(sym_index),
	8043:  uint16(177),
	8044:  uint16(2),
	8045:  uint16(sym_range),
	8046:  uint16(aux_sym_list_element_access_repeat1),
	8047:  uint16(297),
	8048:  uint16(2),
	8049:  uint16(sym__command_substitution_dollar),
	8050:  uint16(sym__command_substitution_inner),
	8051:  uint16(285),
	8052:  uint16(4),
	8053:  uint16(sym_command_substitution),
	8054:  uint16(sym_variable_expansion),
	8055:  uint16(sym_double_quote_string),
	8056:  uint16(sym_single_quote_string),
	8057:  uint16(12),
	8058:  uint16(771),
	8059:  uint16(1),
	8060:  uint16(sym_comment),
	8061:  uint16(810),
	8062:  uint16(1),
	8063:  uint16(anon_sym_DOLLAR),
	8064:  uint16(813),
	8065:  uint16(1),
	8066:  uint16(anon_sym_LPAREN),
	8067:  uint16(816),
	8068:  uint16(1),
	8069:  uint16(sym_integer),
	8070:  uint16(819),
	8071:  uint16(1),
	8072:  uint16(anon_sym_DOT_DOT),
	8073:  uint16(822),
	8074:  uint16(1),
	8075:  uint16(anon_sym_RBRACK),
	8076:  uint16(824),
	8077:  uint16(1),
	8078:  uint16(anon_sym_DQUOTE),
	8079:  uint16(827),
	8080:  uint16(1),
	8081:  uint16(anon_sym_SQUOTE),
	8082:  uint16(284),
	8083:  uint16(1),
	8084:  uint16(sym_index),
	8085:  uint16(185),
	8086:  uint16(2),
	8087:  uint16(sym_range),
	8088:  uint16(aux_sym_list_element_access_repeat1),
	8089:  uint16(297),
	8090:  uint16(2),
	8091:  uint16(sym__command_substitution_dollar),
	8092:  uint16(sym__command_substitution_inner),
	8093:  uint16(285),
	8094:  uint16(4),
	8095:  uint16(sym_command_substitution),
	8096:  uint16(sym_variable_expansion),
	8097:  uint16(sym_double_quote_string),
	8098:  uint16(sym_single_quote_string),
	8099:  uint16(12),
	8100:  uint16(765),
	8101:  uint16(1),
	8102:  uint16(anon_sym_DOLLAR),
	8103:  uint16(767),
	8104:  uint16(1),
	8105:  uint16(anon_sym_LPAREN),
	8106:  uint16(769),
	8107:  uint16(1),
	8108:  uint16(sym_integer),
	8109:  uint16(771),
	8110:  uint16(1),
	8111:  uint16(sym_comment),
	8112:  uint16(773),
	8113:  uint16(1),
	8114:  uint16(anon_sym_DOT_DOT),
	8115:  uint16(777),
	8116:  uint16(1),
	8117:  uint16(anon_sym_DQUOTE),
	8118:  uint16(779),
	8119:  uint16(1),
	8120:  uint16(anon_sym_SQUOTE),
	8121:  uint16(830),
	8122:  uint16(1),
	8123:  uint16(anon_sym_RBRACK),
	8124:  uint16(284),
	8125:  uint16(1),
	8126:  uint16(sym_index),
	8127:  uint16(180),
	8128:  uint16(2),
	8129:  uint16(sym_range),
	8130:  uint16(aux_sym_list_element_access_repeat1),
	8131:  uint16(297),
	8132:  uint16(2),
	8133:  uint16(sym__command_substitution_dollar),
	8134:  uint16(sym__command_substitution_inner),
	8135:  uint16(285),
	8136:  uint16(4),
	8137:  uint16(sym_command_substitution),
	8138:  uint16(sym_variable_expansion),
	8139:  uint16(sym_double_quote_string),
	8140:  uint16(sym_single_quote_string),
	8141:  uint16(3),
	8142:  uint16(3),
	8143:  uint16(1),
	8144:  uint16(sym_comment),
	8145:  uint16(503),
	8146:  uint16(2),
	8147:  uint16(sym__concat),
	8148:  uint16(sym__concat_list),
	8149:  uint16(501),
	8150:  uint16(14),
	8151:  uint16(anon_sym_PIPE_PIPE),
	8152:  uint16(anon_sym_AMP_AMP),
	8153:  uint16(anon_sym_AMP_PIPE),
	8154:  uint16(anon_sym_2_GT_PIPE),
	8155:  uint16(anon_sym_PIPE),
	8156:  uint16(anon_sym_SEMI),
	8157:  uint16(anon_sym_AMP),
	8158:  uint16(anon_sym_LF),
	8159:  uint16(anon_sym_CR),
	8160:  uint16(anon_sym_CR_LF),
	8161:  uint16(anon_sym_RPAREN),
	8162:  uint16(anon_sym_RBRACE),
	8163:  uint16(sym_stream_redirect),
	8164:  uint16(sym_direction),
	8165:  uint16(3),
	8166:  uint16(3),
	8167:  uint16(1),
	8168:  uint16(sym_comment),
	8169:  uint16(664),
	8170:  uint16(1),
	8171:  uint16(sym__concat),
	8172:  uint16(662),
	8173:  uint16(14),
	8174:  uint16(anon_sym_PIPE_PIPE),
	8175:  uint16(anon_sym_AMP_AMP),
	8176:  uint16(anon_sym_AMP_PIPE),
	8177:  uint16(anon_sym_2_GT_PIPE),
	8178:  uint16(anon_sym_PIPE),
	8179:  uint16(anon_sym_SEMI),
	8180:  uint16(anon_sym_AMP),
	8181:  uint16(anon_sym_LF),
	8182:  uint16(anon_sym_CR),
	8183:  uint16(anon_sym_CR_LF),
	8184:  uint16(anon_sym_RPAREN),
	8185:  uint16(anon_sym_RBRACE),
	8186:  uint16(sym_stream_redirect),
	8187:  uint16(sym_direction),
	8188:  uint16(3),
	8189:  uint16(3),
	8190:  uint16(1),
	8191:  uint16(sym_comment),
	8192:  uint16(650),
	8193:  uint16(1),
	8194:  uint16(sym__concat),
	8195:  uint16(648),
	8196:  uint16(14),
	8197:  uint16(anon_sym_PIPE_PIPE),
	8198:  uint16(anon_sym_AMP_AMP),
	8199:  uint16(anon_sym_AMP_PIPE),
	8200:  uint16(anon_sym_2_GT_PIPE),
	8201:  uint16(anon_sym_PIPE),
	8202:  uint16(anon_sym_SEMI),
	8203:  uint16(anon_sym_AMP),
	8204:  uint16(anon_sym_LF),
	8205:  uint16(anon_sym_CR),
	8206:  uint16(anon_sym_CR_LF),
	8207:  uint16(anon_sym_RPAREN),
	8208:  uint16(anon_sym_RBRACE),
	8209:  uint16(sym_stream_redirect),
	8210:  uint16(sym_direction),
	8211:  uint16(7),
	8212:  uint16(3),
	8213:  uint16(1),
	8214:  uint16(sym_comment),
	8215:  uint16(838),
	8216:  uint16(1),
	8217:  uint16(sym_stream_redirect),
	8218:  uint16(840),
	8219:  uint16(1),
	8220:  uint16(sym_direction),
	8221:  uint16(251),
	8222:  uint16(1),
	8223:  uint16(sym_file_redirect),
	8224:  uint16(832),
	8225:  uint16(2),
	8226:  uint16(anon_sym_PIPE_PIPE),
	8227:  uint16(anon_sym_AMP_AMP),
	8228:  uint16(834),
	8229:  uint16(3),
	8230:  uint16(anon_sym_AMP_PIPE),
	8231:  uint16(anon_sym_2_GT_PIPE),
	8232:  uint16(anon_sym_PIPE),
	8233:  uint16(836),
	8234:  uint16(7),
	8235:  uint16(anon_sym_SEMI),
	8236:  uint16(anon_sym_AMP),
	8237:  uint16(anon_sym_LF),
	8238:  uint16(anon_sym_CR),
	8239:  uint16(anon_sym_CR_LF),
	8240:  uint16(anon_sym_RPAREN),
	8241:  uint16(anon_sym_RBRACE),
	8242:  uint16(3),
	8243:  uint16(3),
	8244:  uint16(1),
	8245:  uint16(sym_comment),
	8246:  uint16(251),
	8247:  uint16(1),
	8248:  uint16(sym_file_redirect),
	8249:  uint16(842),
	8250:  uint16(14),
	8251:  uint16(anon_sym_PIPE_PIPE),
	8252:  uint16(anon_sym_AMP_AMP),
	8253:  uint16(anon_sym_AMP_PIPE),
	8254:  uint16(anon_sym_2_GT_PIPE),
	8255:  uint16(anon_sym_PIPE),
	8256:  uint16(anon_sym_SEMI),
	8257:  uint16(anon_sym_AMP),
	8258:  uint16(anon_sym_LF),
	8259:  uint16(anon_sym_CR),
	8260:  uint16(anon_sym_CR_LF),
	8261:  uint16(anon_sym_RPAREN),
	8262:  uint16(anon_sym_RBRACE),
	8263:  uint16(sym_stream_redirect),
	8264:  uint16(sym_direction),
	8265:  uint16(3),
	8266:  uint16(3),
	8267:  uint16(1),
	8268:  uint16(sym_comment),
	8269:  uint16(656),
	8270:  uint16(1),
	8271:  uint16(sym__concat),
	8272:  uint16(544),
	8273:  uint16(14),
	8274:  uint16(anon_sym_PIPE_PIPE),
	8275:  uint16(anon_sym_AMP_AMP),
	8276:  uint16(anon_sym_AMP_PIPE),
	8277:  uint16(anon_sym_2_GT_PIPE),
	8278:  uint16(anon_sym_PIPE),
	8279:  uint16(anon_sym_SEMI),
	8280:  uint16(anon_sym_AMP),
	8281:  uint16(anon_sym_LF),
	8282:  uint16(anon_sym_CR),
	8283:  uint16(anon_sym_CR_LF),
	8284:  uint16(anon_sym_RPAREN),
	8285:  uint16(anon_sym_RBRACE),
	8286:  uint16(sym_stream_redirect),
	8287:  uint16(sym_direction),
	8288:  uint16(3),
	8289:  uint16(3),
	8290:  uint16(1),
	8291:  uint16(sym_comment),
	8292:  uint16(646),
	8293:  uint16(1),
	8294:  uint16(sym__concat),
	8295:  uint16(644),
	8296:  uint16(14),
	8297:  uint16(anon_sym_PIPE_PIPE),
	8298:  uint16(anon_sym_AMP_AMP),
	8299:  uint16(anon_sym_AMP_PIPE),
	8300:  uint16(anon_sym_2_GT_PIPE),
	8301:  uint16(anon_sym_PIPE),
	8302:  uint16(anon_sym_SEMI),
	8303:  uint16(anon_sym_AMP),
	8304:  uint16(anon_sym_LF),
	8305:  uint16(anon_sym_CR),
	8306:  uint16(anon_sym_CR_LF),
	8307:  uint16(anon_sym_RPAREN),
	8308:  uint16(anon_sym_RBRACE),
	8309:  uint16(sym_stream_redirect),
	8310:  uint16(sym_direction),
	8311:  uint16(3),
	8312:  uint16(3),
	8313:  uint16(1),
	8314:  uint16(sym_comment),
	8315:  uint16(668),
	8316:  uint16(1),
	8317:  uint16(sym__concat),
	8318:  uint16(666),
	8319:  uint16(14),
	8320:  uint16(anon_sym_PIPE_PIPE),
	8321:  uint16(anon_sym_AMP_AMP),
	8322:  uint16(anon_sym_AMP_PIPE),
	8323:  uint16(anon_sym_2_GT_PIPE),
	8324:  uint16(anon_sym_PIPE),
	8325:  uint16(anon_sym_SEMI),
	8326:  uint16(anon_sym_AMP),
	8327:  uint16(anon_sym_LF),
	8328:  uint16(anon_sym_CR),
	8329:  uint16(anon_sym_CR_LF),
	8330:  uint16(anon_sym_RPAREN),
	8331:  uint16(anon_sym_RBRACE),
	8332:  uint16(sym_stream_redirect),
	8333:  uint16(sym_direction),
	8334:  uint16(3),
	8335:  uint16(3),
	8336:  uint16(1),
	8337:  uint16(sym_comment),
	8338:  uint16(660),
	8339:  uint16(1),
	8340:  uint16(sym__concat),
	8341:  uint16(658),
	8342:  uint16(14),
	8343:  uint16(anon_sym_PIPE_PIPE),
	8344:  uint16(anon_sym_AMP_AMP),
	8345:  uint16(anon_sym_AMP_PIPE),
	8346:  uint16(anon_sym_2_GT_PIPE),
	8347:  uint16(anon_sym_PIPE),
	8348:  uint16(anon_sym_SEMI),
	8349:  uint16(anon_sym_AMP),
	8350:  uint16(anon_sym_LF),
	8351:  uint16(anon_sym_CR),
	8352:  uint16(anon_sym_CR_LF),
	8353:  uint16(anon_sym_RPAREN),
	8354:  uint16(anon_sym_RBRACE),
	8355:  uint16(sym_stream_redirect),
	8356:  uint16(sym_direction),
	8357:  uint16(7),
	8358:  uint16(3),
	8359:  uint16(1),
	8360:  uint16(sym_comment),
	8361:  uint16(838),
	8362:  uint16(1),
	8363:  uint16(sym_stream_redirect),
	8364:  uint16(840),
	8365:  uint16(1),
	8366:  uint16(sym_direction),
	8367:  uint16(251),
	8368:  uint16(1),
	8369:  uint16(sym_file_redirect),
	8370:  uint16(832),
	8371:  uint16(2),
	8372:  uint16(anon_sym_PIPE_PIPE),
	8373:  uint16(anon_sym_AMP_AMP),
	8374:  uint16(834),
	8375:  uint16(3),
	8376:  uint16(anon_sym_AMP_PIPE),
	8377:  uint16(anon_sym_2_GT_PIPE),
	8378:  uint16(anon_sym_PIPE),
	8379:  uint16(844),
	8380:  uint16(7),
	8381:  uint16(anon_sym_SEMI),
	8382:  uint16(anon_sym_AMP),
	8383:  uint16(anon_sym_LF),
	8384:  uint16(anon_sym_CR),
	8385:  uint16(anon_sym_CR_LF),
	8386:  uint16(anon_sym_RPAREN),
	8387:  uint16(anon_sym_RBRACE),
	8388:  uint16(3),
	8389:  uint16(3),
	8390:  uint16(1),
	8391:  uint16(sym_comment),
	8392:  uint16(688),
	8393:  uint16(1),
	8394:  uint16(sym__concat),
	8395:  uint16(686),
	8396:  uint16(14),
	8397:  uint16(anon_sym_PIPE_PIPE),
	8398:  uint16(anon_sym_AMP_AMP),
	8399:  uint16(anon_sym_AMP_PIPE),
	8400:  uint16(anon_sym_2_GT_PIPE),
	8401:  uint16(anon_sym_PIPE),
	8402:  uint16(anon_sym_SEMI),
	8403:  uint16(anon_sym_AMP),
	8404:  uint16(anon_sym_LF),
	8405:  uint16(anon_sym_CR),
	8406:  uint16(anon_sym_CR_LF),
	8407:  uint16(anon_sym_RPAREN),
	8408:  uint16(anon_sym_RBRACE),
	8409:  uint16(sym_stream_redirect),
	8410:  uint16(sym_direction),
	8411:  uint16(3),
	8412:  uint16(3),
	8413:  uint16(1),
	8414:  uint16(sym_comment),
	8415:  uint16(642),
	8416:  uint16(1),
	8417:  uint16(sym__concat),
	8418:  uint16(640),
	8419:  uint16(14),
	8420:  uint16(anon_sym_PIPE_PIPE),
	8421:  uint16(anon_sym_AMP_AMP),
	8422:  uint16(anon_sym_AMP_PIPE),
	8423:  uint16(anon_sym_2_GT_PIPE),
	8424:  uint16(anon_sym_PIPE),
	8425:  uint16(anon_sym_SEMI),
	8426:  uint16(anon_sym_AMP),
	8427:  uint16(anon_sym_LF),
	8428:  uint16(anon_sym_CR),
	8429:  uint16(anon_sym_CR_LF),
	8430:  uint16(anon_sym_RPAREN),
	8431:  uint16(anon_sym_RBRACE),
	8432:  uint16(sym_stream_redirect),
	8433:  uint16(sym_direction),
	8434:  uint16(3),
	8435:  uint16(3),
	8436:  uint16(1),
	8437:  uint16(sym_comment),
	8438:  uint16(654),
	8439:  uint16(1),
	8440:  uint16(sym__concat),
	8441:  uint16(652),
	8442:  uint16(14),
	8443:  uint16(anon_sym_PIPE_PIPE),
	8444:  uint16(anon_sym_AMP_AMP),
	8445:  uint16(anon_sym_AMP_PIPE),
	8446:  uint16(anon_sym_2_GT_PIPE),
	8447:  uint16(anon_sym_PIPE),
	8448:  uint16(anon_sym_SEMI),
	8449:  uint16(anon_sym_AMP),
	8450:  uint16(anon_sym_LF),
	8451:  uint16(anon_sym_CR),
	8452:  uint16(anon_sym_CR_LF),
	8453:  uint16(anon_sym_RPAREN),
	8454:  uint16(anon_sym_RBRACE),
	8455:  uint16(sym_stream_redirect),
	8456:  uint16(sym_direction),
	8457:  uint16(3),
	8458:  uint16(3),
	8459:  uint16(1),
	8460:  uint16(sym_comment),
	8461:  uint16(638),
	8462:  uint16(1),
	8463:  uint16(sym__concat),
	8464:  uint16(636),
	8465:  uint16(14),
	8466:  uint16(anon_sym_PIPE_PIPE),
	8467:  uint16(anon_sym_AMP_AMP),
	8468:  uint16(anon_sym_AMP_PIPE),
	8469:  uint16(anon_sym_2_GT_PIPE),
	8470:  uint16(anon_sym_PIPE),
	8471:  uint16(anon_sym_SEMI),
	8472:  uint16(anon_sym_AMP),
	8473:  uint16(anon_sym_LF),
	8474:  uint16(anon_sym_CR),
	8475:  uint16(anon_sym_CR_LF),
	8476:  uint16(anon_sym_RPAREN),
	8477:  uint16(anon_sym_RBRACE),
	8478:  uint16(sym_stream_redirect),
	8479:  uint16(sym_direction),
	8480:  uint16(3),
	8481:  uint16(3),
	8482:  uint16(1),
	8483:  uint16(sym_comment),
	8484:  uint16(684),
	8485:  uint16(1),
	8486:  uint16(sym__concat),
	8487:  uint16(682),
	8488:  uint16(14),
	8489:  uint16(anon_sym_PIPE_PIPE),
	8490:  uint16(anon_sym_AMP_AMP),
	8491:  uint16(anon_sym_AMP_PIPE),
	8492:  uint16(anon_sym_2_GT_PIPE),
	8493:  uint16(anon_sym_PIPE),
	8494:  uint16(anon_sym_SEMI),
	8495:  uint16(anon_sym_AMP),
	8496:  uint16(anon_sym_LF),
	8497:  uint16(anon_sym_CR),
	8498:  uint16(anon_sym_CR_LF),
	8499:  uint16(anon_sym_RPAREN),
	8500:  uint16(anon_sym_RBRACE),
	8501:  uint16(sym_stream_redirect),
	8502:  uint16(sym_direction),
	8503:  uint16(6),
	8504:  uint16(3),
	8505:  uint16(1),
	8506:  uint16(sym_comment),
	8507:  uint16(838),
	8508:  uint16(1),
	8509:  uint16(sym_stream_redirect),
	8510:  uint16(840),
	8511:  uint16(1),
	8512:  uint16(sym_direction),
	8513:  uint16(251),
	8514:  uint16(1),
	8515:  uint16(sym_file_redirect),
	8516:  uint16(834),
	8517:  uint16(3),
	8518:  uint16(anon_sym_AMP_PIPE),
	8519:  uint16(anon_sym_2_GT_PIPE),
	8520:  uint16(anon_sym_PIPE),
	8521:  uint16(846),
	8522:  uint16(9),
	8523:  uint16(anon_sym_PIPE_PIPE),
	8524:  uint16(anon_sym_AMP_AMP),
	8525:  uint16(anon_sym_SEMI),
	8526:  uint16(anon_sym_AMP),
	8527:  uint16(anon_sym_LF),
	8528:  uint16(anon_sym_CR),
	8529:  uint16(anon_sym_CR_LF),
	8530:  uint16(anon_sym_RPAREN),
	8531:  uint16(anon_sym_RBRACE),
	8532:  uint16(3),
	8533:  uint16(3),
	8534:  uint16(1),
	8535:  uint16(sym_comment),
	8536:  uint16(680),
	8537:  uint16(1),
	8538:  uint16(sym__concat),
	8539:  uint16(678),
	8540:  uint16(14),
	8541:  uint16(anon_sym_PIPE_PIPE),
	8542:  uint16(anon_sym_AMP_AMP),
	8543:  uint16(anon_sym_AMP_PIPE),
	8544:  uint16(anon_sym_2_GT_PIPE),
	8545:  uint16(anon_sym_PIPE),
	8546:  uint16(anon_sym_SEMI),
	8547:  uint16(anon_sym_AMP),
	8548:  uint16(anon_sym_LF),
	8549:  uint16(anon_sym_CR),
	8550:  uint16(anon_sym_CR_LF),
	8551:  uint16(anon_sym_RPAREN),
	8552:  uint16(anon_sym_RBRACE),
	8553:  uint16(sym_stream_redirect),
	8554:  uint16(sym_direction),
	8555:  uint16(3),
	8556:  uint16(3),
	8557:  uint16(1),
	8558:  uint16(sym_comment),
	8559:  uint16(672),
	8560:  uint16(1),
	8561:  uint16(sym__concat),
	8562:  uint16(670),
	8563:  uint16(14),
	8564:  uint16(anon_sym_PIPE_PIPE),
	8565:  uint16(anon_sym_AMP_AMP),
	8566:  uint16(anon_sym_AMP_PIPE),
	8567:  uint16(anon_sym_2_GT_PIPE),
	8568:  uint16(anon_sym_PIPE),
	8569:  uint16(anon_sym_SEMI),
	8570:  uint16(anon_sym_AMP),
	8571:  uint16(anon_sym_LF),
	8572:  uint16(anon_sym_CR),
	8573:  uint16(anon_sym_CR_LF),
	8574:  uint16(anon_sym_RPAREN),
	8575:  uint16(anon_sym_RBRACE),
	8576:  uint16(sym_stream_redirect),
	8577:  uint16(sym_direction),
	8578:  uint16(3),
	8579:  uint16(3),
	8580:  uint16(1),
	8581:  uint16(sym_comment),
	8582:  uint16(676),
	8583:  uint16(1),
	8584:  uint16(sym__concat),
	8585:  uint16(674),
	8586:  uint16(14),
	8587:  uint16(anon_sym_PIPE_PIPE),
	8588:  uint16(anon_sym_AMP_AMP),
	8589:  uint16(anon_sym_AMP_PIPE),
	8590:  uint16(anon_sym_2_GT_PIPE),
	8591:  uint16(anon_sym_PIPE),
	8592:  uint16(anon_sym_SEMI),
	8593:  uint16(anon_sym_AMP),
	8594:  uint16(anon_sym_LF),
	8595:  uint16(anon_sym_CR),
	8596:  uint16(anon_sym_CR_LF),
	8597:  uint16(anon_sym_RPAREN),
	8598:  uint16(anon_sym_RBRACE),
	8599:  uint16(sym_stream_redirect),
	8600:  uint16(sym_direction),
	8601:  uint16(2),
	8602:  uint16(3),
	8603:  uint16(1),
	8604:  uint16(sym_comment),
	8605:  uint16(848),
	8606:  uint16(14),
	8607:  uint16(anon_sym_PIPE_PIPE),
	8608:  uint16(anon_sym_AMP_AMP),
	8609:  uint16(anon_sym_AMP_PIPE),
	8610:  uint16(anon_sym_2_GT_PIPE),
	8611:  uint16(anon_sym_PIPE),
	8612:  uint16(anon_sym_SEMI),
	8613:  uint16(anon_sym_AMP),
	8614:  uint16(anon_sym_LF),
	8615:  uint16(anon_sym_CR),
	8616:  uint16(anon_sym_CR_LF),
	8617:  uint16(anon_sym_RPAREN),
	8618:  uint16(anon_sym_RBRACE),
	8619:  uint16(sym_stream_redirect),
	8620:  uint16(sym_direction),
	8621:  uint16(8),
	8622:  uint16(3),
	8623:  uint16(1),
	8624:  uint16(sym_comment),
	8625:  uint16(293),
	8626:  uint16(1),
	8627:  uint16(anon_sym_RPAREN),
	8628:  uint16(838),
	8629:  uint16(1),
	8630:  uint16(sym_stream_redirect),
	8631:  uint16(840),
	8632:  uint16(1),
	8633:  uint16(sym_direction),
	8634:  uint16(251),
	8635:  uint16(1),
	8636:  uint16(sym_file_redirect),
	8637:  uint16(832),
	8638:  uint16(2),
	8639:  uint16(anon_sym_PIPE_PIPE),
	8640:  uint16(anon_sym_AMP_AMP),
	8641:  uint16(834),
	8642:  uint16(3),
	8643:  uint16(anon_sym_AMP_PIPE),
	8644:  uint16(anon_sym_2_GT_PIPE),
	8645:  uint16(anon_sym_PIPE),
	8646:  uint16(850),
	8647:  uint16(5),
	8648:  uint16(anon_sym_SEMI),
	8649:  uint16(anon_sym_AMP),
	8650:  uint16(anon_sym_LF),
	8651:  uint16(anon_sym_CR),
	8652:  uint16(anon_sym_CR_LF),
	8653:  uint16(8),
	8654:  uint16(3),
	8655:  uint16(1),
	8656:  uint16(sym_comment),
	8657:  uint16(838),
	8658:  uint16(1),
	8659:  uint16(sym_stream_redirect),
	8660:  uint16(840),
	8661:  uint16(1),
	8662:  uint16(sym_direction),
	8663:  uint16(852),
	8664:  uint16(1),
	8665:  uint16(anon_sym_RPAREN),
	8666:  uint16(251),
	8667:  uint16(1),
	8668:  uint16(sym_file_redirect),
	8669:  uint16(832),
	8670:  uint16(2),
	8671:  uint16(anon_sym_PIPE_PIPE),
	8672:  uint16(anon_sym_AMP_AMP),
	8673:  uint16(834),
	8674:  uint16(3),
	8675:  uint16(anon_sym_AMP_PIPE),
	8676:  uint16(anon_sym_2_GT_PIPE),
	8677:  uint16(anon_sym_PIPE),
	8678:  uint16(850),
	8679:  uint16(5),
	8680:  uint16(anon_sym_SEMI),
	8681:  uint16(anon_sym_AMP),
	8682:  uint16(anon_sym_LF),
	8683:  uint16(anon_sym_CR),
	8684:  uint16(anon_sym_CR_LF),
	8685:  uint16(2),
	8686:  uint16(3),
	8687:  uint16(1),
	8688:  uint16(sym_comment),
	8689:  uint16(854),
	8690:  uint16(14),
	8691:  uint16(anon_sym_PIPE_PIPE),
	8692:  uint16(anon_sym_AMP_AMP),
	8693:  uint16(anon_sym_AMP_PIPE),
	8694:  uint16(anon_sym_2_GT_PIPE),
	8695:  uint16(anon_sym_PIPE),
	8696:  uint16(anon_sym_SEMI),
	8697:  uint16(anon_sym_AMP),
	8698:  uint16(anon_sym_LF),
	8699:  uint16(anon_sym_CR),
	8700:  uint16(anon_sym_CR_LF),
	8701:  uint16(anon_sym_RPAREN),
	8702:  uint16(anon_sym_RBRACE),
	8703:  uint16(sym_stream_redirect),
	8704:  uint16(sym_direction),
	8705:  uint16(8),
	8706:  uint16(3),
	8707:  uint16(1),
	8708:  uint16(sym_comment),
	8709:  uint16(289),
	8710:  uint16(1),
	8711:  uint16(anon_sym_RPAREN),
	8712:  uint16(838),
	8713:  uint16(1),
	8714:  uint16(sym_stream_redirect),
	8715:  uint16(840),
	8716:  uint16(1),
	8717:  uint16(sym_direction),
	8718:  uint16(251),
	8719:  uint16(1),
	8720:  uint16(sym_file_redirect),
	8721:  uint16(832),
	8722:  uint16(2),
	8723:  uint16(anon_sym_PIPE_PIPE),
	8724:  uint16(anon_sym_AMP_AMP),
	8725:  uint16(834),
	8726:  uint16(3),
	8727:  uint16(anon_sym_AMP_PIPE),
	8728:  uint16(anon_sym_2_GT_PIPE),
	8729:  uint16(anon_sym_PIPE),
	8730:  uint16(850),
	8731:  uint16(5),
	8732:  uint16(anon_sym_SEMI),
	8733:  uint16(anon_sym_AMP),
	8734:  uint16(anon_sym_LF),
	8735:  uint16(anon_sym_CR),
	8736:  uint16(anon_sym_CR_LF),
	8737:  uint16(2),
	8738:  uint16(3),
	8739:  uint16(1),
	8740:  uint16(sym_comment),
	8741:  uint16(854),
	8742:  uint16(14),
	8743:  uint16(anon_sym_PIPE_PIPE),
	8744:  uint16(anon_sym_AMP_AMP),
	8745:  uint16(anon_sym_AMP_PIPE),
	8746:  uint16(anon_sym_2_GT_PIPE),
	8747:  uint16(anon_sym_PIPE),
	8748:  uint16(anon_sym_SEMI),
	8749:  uint16(anon_sym_AMP),
	8750:  uint16(anon_sym_LF),
	8751:  uint16(anon_sym_CR),
	8752:  uint16(anon_sym_CR_LF),
	8753:  uint16(anon_sym_RPAREN),
	8754:  uint16(anon_sym_RBRACE),
	8755:  uint16(sym_stream_redirect),
	8756:  uint16(sym_direction),
	8757:  uint16(2),
	8758:  uint16(3),
	8759:  uint16(1),
	8760:  uint16(sym_comment),
	8761:  uint16(856),
	8762:  uint16(14),
	8763:  uint16(anon_sym_PIPE_PIPE),
	8764:  uint16(anon_sym_AMP_AMP),
	8765:  uint16(anon_sym_AMP_PIPE),
	8766:  uint16(anon_sym_2_GT_PIPE),
	8767:  uint16(anon_sym_PIPE),
	8768:  uint16(anon_sym_SEMI),
	8769:  uint16(anon_sym_AMP),
	8770:  uint16(anon_sym_LF),
	8771:  uint16(anon_sym_CR),
	8772:  uint16(anon_sym_CR_LF),
	8773:  uint16(anon_sym_RPAREN),
	8774:  uint16(anon_sym_RBRACE),
	8775:  uint16(sym_stream_redirect),
	8776:  uint16(sym_direction),
	8777:  uint16(2),
	8778:  uint16(3),
	8779:  uint16(1),
	8780:  uint16(sym_comment),
	8781:  uint16(858),
	8782:  uint16(14),
	8783:  uint16(anon_sym_PIPE_PIPE),
	8784:  uint16(anon_sym_AMP_AMP),
	8785:  uint16(anon_sym_AMP_PIPE),
	8786:  uint16(anon_sym_2_GT_PIPE),
	8787:  uint16(anon_sym_PIPE),
	8788:  uint16(anon_sym_SEMI),
	8789:  uint16(anon_sym_AMP),
	8790:  uint16(anon_sym_LF),
	8791:  uint16(anon_sym_CR),
	8792:  uint16(anon_sym_CR_LF),
	8793:  uint16(anon_sym_RPAREN),
	8794:  uint16(anon_sym_RBRACE),
	8795:  uint16(sym_stream_redirect),
	8796:  uint16(sym_direction),
	8797:  uint16(2),
	8798:  uint16(3),
	8799:  uint16(1),
	8800:  uint16(sym_comment),
	8801:  uint16(860),
	8802:  uint16(14),
	8803:  uint16(anon_sym_PIPE_PIPE),
	8804:  uint16(anon_sym_AMP_AMP),
	8805:  uint16(anon_sym_AMP_PIPE),
	8806:  uint16(anon_sym_2_GT_PIPE),
	8807:  uint16(anon_sym_PIPE),
	8808:  uint16(anon_sym_SEMI),
	8809:  uint16(anon_sym_AMP),
	8810:  uint16(anon_sym_LF),
	8811:  uint16(anon_sym_CR),
	8812:  uint16(anon_sym_CR_LF),
	8813:  uint16(anon_sym_RPAREN),
	8814:  uint16(anon_sym_RBRACE),
	8815:  uint16(sym_stream_redirect),
	8816:  uint16(sym_direction),
	8817:  uint16(2),
	8818:  uint16(3),
	8819:  uint16(1),
	8820:  uint16(sym_comment),
	8821:  uint16(862),
	8822:  uint16(14),
	8823:  uint16(anon_sym_PIPE_PIPE),
	8824:  uint16(anon_sym_AMP_AMP),
	8825:  uint16(anon_sym_AMP_PIPE),
	8826:  uint16(anon_sym_2_GT_PIPE),
	8827:  uint16(anon_sym_PIPE),
	8828:  uint16(anon_sym_SEMI),
	8829:  uint16(anon_sym_AMP),
	8830:  uint16(anon_sym_LF),
	8831:  uint16(anon_sym_CR),
	8832:  uint16(anon_sym_CR_LF),
	8833:  uint16(anon_sym_RPAREN),
	8834:  uint16(anon_sym_RBRACE),
	8835:  uint16(sym_stream_redirect),
	8836:  uint16(sym_direction),
	8837:  uint16(2),
	8838:  uint16(3),
	8839:  uint16(1),
	8840:  uint16(sym_comment),
	8841:  uint16(692),
	8842:  uint16(14),
	8843:  uint16(anon_sym_PIPE_PIPE),
	8844:  uint16(anon_sym_AMP_AMP),
	8845:  uint16(anon_sym_AMP_PIPE),
	8846:  uint16(anon_sym_2_GT_PIPE),
	8847:  uint16(anon_sym_PIPE),
	8848:  uint16(anon_sym_SEMI),
	8849:  uint16(anon_sym_AMP),
	8850:  uint16(anon_sym_LF),
	8851:  uint16(anon_sym_CR),
	8852:  uint16(anon_sym_CR_LF),
	8853:  uint16(anon_sym_RPAREN),
	8854:  uint16(anon_sym_RBRACE),
	8855:  uint16(sym_stream_redirect),
	8856:  uint16(sym_direction),
	8857:  uint16(2),
	8858:  uint16(3),
	8859:  uint16(1),
	8860:  uint16(sym_comment),
	8861:  uint16(862),
	8862:  uint16(14),
	8863:  uint16(anon_sym_PIPE_PIPE),
	8864:  uint16(anon_sym_AMP_AMP),
	8865:  uint16(anon_sym_AMP_PIPE),
	8866:  uint16(anon_sym_2_GT_PIPE),
	8867:  uint16(anon_sym_PIPE),
	8868:  uint16(anon_sym_SEMI),
	8869:  uint16(anon_sym_AMP),
	8870:  uint16(anon_sym_LF),
	8871:  uint16(anon_sym_CR),
	8872:  uint16(anon_sym_CR_LF),
	8873:  uint16(anon_sym_RPAREN),
	8874:  uint16(anon_sym_RBRACE),
	8875:  uint16(sym_stream_redirect),
	8876:  uint16(sym_direction),
	8877:  uint16(2),
	8878:  uint16(3),
	8879:  uint16(1),
	8880:  uint16(sym_comment),
	8881:  uint16(862),
	8882:  uint16(14),
	8883:  uint16(anon_sym_PIPE_PIPE),
	8884:  uint16(anon_sym_AMP_AMP),
	8885:  uint16(anon_sym_AMP_PIPE),
	8886:  uint16(anon_sym_2_GT_PIPE),
	8887:  uint16(anon_sym_PIPE),
	8888:  uint16(anon_sym_SEMI),
	8889:  uint16(anon_sym_AMP),
	8890:  uint16(anon_sym_LF),
	8891:  uint16(anon_sym_CR),
	8892:  uint16(anon_sym_CR_LF),
	8893:  uint16(anon_sym_RPAREN),
	8894:  uint16(anon_sym_RBRACE),
	8895:  uint16(sym_stream_redirect),
	8896:  uint16(sym_direction),
	8897:  uint16(2),
	8898:  uint16(3),
	8899:  uint16(1),
	8900:  uint16(sym_comment),
	8901:  uint16(862),
	8902:  uint16(14),
	8903:  uint16(anon_sym_PIPE_PIPE),
	8904:  uint16(anon_sym_AMP_AMP),
	8905:  uint16(anon_sym_AMP_PIPE),
	8906:  uint16(anon_sym_2_GT_PIPE),
	8907:  uint16(anon_sym_PIPE),
	8908:  uint16(anon_sym_SEMI),
	8909:  uint16(anon_sym_AMP),
	8910:  uint16(anon_sym_LF),
	8911:  uint16(anon_sym_CR),
	8912:  uint16(anon_sym_CR_LF),
	8913:  uint16(anon_sym_RPAREN),
	8914:  uint16(anon_sym_RBRACE),
	8915:  uint16(sym_stream_redirect),
	8916:  uint16(sym_direction),
	8917:  uint16(2),
	8918:  uint16(3),
	8919:  uint16(1),
	8920:  uint16(sym_comment),
	8921:  uint16(862),
	8922:  uint16(14),
	8923:  uint16(anon_sym_PIPE_PIPE),
	8924:  uint16(anon_sym_AMP_AMP),
	8925:  uint16(anon_sym_AMP_PIPE),
	8926:  uint16(anon_sym_2_GT_PIPE),
	8927:  uint16(anon_sym_PIPE),
	8928:  uint16(anon_sym_SEMI),
	8929:  uint16(anon_sym_AMP),
	8930:  uint16(anon_sym_LF),
	8931:  uint16(anon_sym_CR),
	8932:  uint16(anon_sym_CR_LF),
	8933:  uint16(anon_sym_RPAREN),
	8934:  uint16(anon_sym_RBRACE),
	8935:  uint16(sym_stream_redirect),
	8936:  uint16(sym_direction),
	8937:  uint16(2),
	8938:  uint16(3),
	8939:  uint16(1),
	8940:  uint16(sym_comment),
	8941:  uint16(864),
	8942:  uint16(14),
	8943:  uint16(anon_sym_PIPE_PIPE),
	8944:  uint16(anon_sym_AMP_AMP),
	8945:  uint16(anon_sym_AMP_PIPE),
	8946:  uint16(anon_sym_2_GT_PIPE),
	8947:  uint16(anon_sym_PIPE),
	8948:  uint16(anon_sym_SEMI),
	8949:  uint16(anon_sym_AMP),
	8950:  uint16(anon_sym_LF),
	8951:  uint16(anon_sym_CR),
	8952:  uint16(anon_sym_CR_LF),
	8953:  uint16(anon_sym_RPAREN),
	8954:  uint16(anon_sym_RBRACE),
	8955:  uint16(sym_stream_redirect),
	8956:  uint16(sym_direction),
	8957:  uint16(8),
	8958:  uint16(3),
	8959:  uint16(1),
	8960:  uint16(sym_comment),
	8961:  uint16(146),
	8962:  uint16(1),
	8963:  uint16(anon_sym_RBRACE),
	8964:  uint16(838),
	8965:  uint16(1),
	8966:  uint16(sym_stream_redirect),
	8967:  uint16(840),
	8968:  uint16(1),
	8969:  uint16(sym_direction),
	8970:  uint16(251),
	8971:  uint16(1),
	8972:  uint16(sym_file_redirect),
	8973:  uint16(832),
	8974:  uint16(2),
	8975:  uint16(anon_sym_PIPE_PIPE),
	8976:  uint16(anon_sym_AMP_AMP),
	8977:  uint16(834),
	8978:  uint16(3),
	8979:  uint16(anon_sym_AMP_PIPE),
	8980:  uint16(anon_sym_2_GT_PIPE),
	8981:  uint16(anon_sym_PIPE),
	8982:  uint16(866),
	8983:  uint16(5),
	8984:  uint16(anon_sym_SEMI),
	8985:  uint16(anon_sym_AMP),
	8986:  uint16(anon_sym_LF),
	8987:  uint16(anon_sym_CR),
	8988:  uint16(anon_sym_CR_LF),
	8989:  uint16(2),
	8990:  uint16(3),
	8991:  uint16(1),
	8992:  uint16(sym_comment),
	8993:  uint16(868),
	8994:  uint16(14),
	8995:  uint16(anon_sym_PIPE_PIPE),
	8996:  uint16(anon_sym_AMP_AMP),
	8997:  uint16(anon_sym_AMP_PIPE),
	8998:  uint16(anon_sym_2_GT_PIPE),
	8999:  uint16(anon_sym_PIPE),
	9000:  uint16(anon_sym_SEMI),
	9001:  uint16(anon_sym_AMP),
	9002:  uint16(anon_sym_LF),
	9003:  uint16(anon_sym_CR),
	9004:  uint16(anon_sym_CR_LF),
	9005:  uint16(anon_sym_RPAREN),
	9006:  uint16(anon_sym_RBRACE),
	9007:  uint16(sym_stream_redirect),
	9008:  uint16(sym_direction),
	9009:  uint16(8),
	9010:  uint16(3),
	9011:  uint16(1),
	9012:  uint16(sym_comment),
	9013:  uint16(838),
	9014:  uint16(1),
	9015:  uint16(sym_stream_redirect),
	9016:  uint16(840),
	9017:  uint16(1),
	9018:  uint16(sym_direction),
	9019:  uint16(870),
	9020:  uint16(1),
	9021:  uint16(anon_sym_RPAREN),
	9022:  uint16(251),
	9023:  uint16(1),
	9024:  uint16(sym_file_redirect),
	9025:  uint16(832),
	9026:  uint16(2),
	9027:  uint16(anon_sym_PIPE_PIPE),
	9028:  uint16(anon_sym_AMP_AMP),
	9029:  uint16(834),
	9030:  uint16(3),
	9031:  uint16(anon_sym_AMP_PIPE),
	9032:  uint16(anon_sym_2_GT_PIPE),
	9033:  uint16(anon_sym_PIPE),
	9034:  uint16(850),
	9035:  uint16(5),
	9036:  uint16(anon_sym_SEMI),
	9037:  uint16(anon_sym_AMP),
	9038:  uint16(anon_sym_LF),
	9039:  uint16(anon_sym_CR),
	9040:  uint16(anon_sym_CR_LF),
	9041:  uint16(2),
	9042:  uint16(3),
	9043:  uint16(1),
	9044:  uint16(sym_comment),
	9045:  uint16(872),
	9046:  uint16(14),
	9047:  uint16(anon_sym_PIPE_PIPE),
	9048:  uint16(anon_sym_AMP_AMP),
	9049:  uint16(anon_sym_AMP_PIPE),
	9050:  uint16(anon_sym_2_GT_PIPE),
	9051:  uint16(anon_sym_PIPE),
	9052:  uint16(anon_sym_SEMI),
	9053:  uint16(anon_sym_AMP),
	9054:  uint16(anon_sym_LF),
	9055:  uint16(anon_sym_CR),
	9056:  uint16(anon_sym_CR_LF),
	9057:  uint16(anon_sym_RPAREN),
	9058:  uint16(anon_sym_RBRACE),
	9059:  uint16(sym_stream_redirect),
	9060:  uint16(sym_direction),
	9061:  uint16(8),
	9062:  uint16(3),
	9063:  uint16(1),
	9064:  uint16(sym_comment),
	9065:  uint16(253),
	9066:  uint16(1),
	9067:  uint16(anon_sym_RPAREN),
	9068:  uint16(838),
	9069:  uint16(1),
	9070:  uint16(sym_stream_redirect),
	9071:  uint16(840),
	9072:  uint16(1),
	9073:  uint16(sym_direction),
	9074:  uint16(251),
	9075:  uint16(1),
	9076:  uint16(sym_file_redirect),
	9077:  uint16(832),
	9078:  uint16(2),
	9079:  uint16(anon_sym_PIPE_PIPE),
	9080:  uint16(anon_sym_AMP_AMP),
	9081:  uint16(834),
	9082:  uint16(3),
	9083:  uint16(anon_sym_AMP_PIPE),
	9084:  uint16(anon_sym_2_GT_PIPE),
	9085:  uint16(anon_sym_PIPE),
	9086:  uint16(850),
	9087:  uint16(5),
	9088:  uint16(anon_sym_SEMI),
	9089:  uint16(anon_sym_AMP),
	9090:  uint16(anon_sym_LF),
	9091:  uint16(anon_sym_CR),
	9092:  uint16(anon_sym_CR_LF),
	9093:  uint16(2),
	9094:  uint16(3),
	9095:  uint16(1),
	9096:  uint16(sym_comment),
	9097:  uint16(874),
	9098:  uint16(14),
	9099:  uint16(anon_sym_PIPE_PIPE),
	9100:  uint16(anon_sym_AMP_AMP),
	9101:  uint16(anon_sym_AMP_PIPE),
	9102:  uint16(anon_sym_2_GT_PIPE),
	9103:  uint16(anon_sym_PIPE),
	9104:  uint16(anon_sym_SEMI),
	9105:  uint16(anon_sym_AMP),
	9106:  uint16(anon_sym_LF),
	9107:  uint16(anon_sym_CR),
	9108:  uint16(anon_sym_CR_LF),
	9109:  uint16(anon_sym_RPAREN),
	9110:  uint16(anon_sym_RBRACE),
	9111:  uint16(sym_stream_redirect),
	9112:  uint16(sym_direction),
	9113:  uint16(2),
	9114:  uint16(3),
	9115:  uint16(1),
	9116:  uint16(sym_comment),
	9117:  uint16(874),
	9118:  uint16(14),
	9119:  uint16(anon_sym_PIPE_PIPE),
	9120:  uint16(anon_sym_AMP_AMP),
	9121:  uint16(anon_sym_AMP_PIPE),
	9122:  uint16(anon_sym_2_GT_PIPE),
	9123:  uint16(anon_sym_PIPE),
	9124:  uint16(anon_sym_SEMI),
	9125:  uint16(anon_sym_AMP),
	9126:  uint16(anon_sym_LF),
	9127:  uint16(anon_sym_CR),
	9128:  uint16(anon_sym_CR_LF),
	9129:  uint16(anon_sym_RPAREN),
	9130:  uint16(anon_sym_RBRACE),
	9131:  uint16(sym_stream_redirect),
	9132:  uint16(sym_direction),
	9133:  uint16(8),
	9134:  uint16(3),
	9135:  uint16(1),
	9136:  uint16(sym_comment),
	9137:  uint16(838),
	9138:  uint16(1),
	9139:  uint16(sym_stream_redirect),
	9140:  uint16(840),
	9141:  uint16(1),
	9142:  uint16(sym_direction),
	9143:  uint16(876),
	9144:  uint16(1),
	9145:  uint16(anon_sym_RPAREN),
	9146:  uint16(251),
	9147:  uint16(1),
	9148:  uint16(sym_file_redirect),
	9149:  uint16(832),
	9150:  uint16(2),
	9151:  uint16(anon_sym_PIPE_PIPE),
	9152:  uint16(anon_sym_AMP_AMP),
	9153:  uint16(834),
	9154:  uint16(3),
	9155:  uint16(anon_sym_AMP_PIPE),
	9156:  uint16(anon_sym_2_GT_PIPE),
	9157:  uint16(anon_sym_PIPE),
	9158:  uint16(850),
	9159:  uint16(5),
	9160:  uint16(anon_sym_SEMI),
	9161:  uint16(anon_sym_AMP),
	9162:  uint16(anon_sym_LF),
	9163:  uint16(anon_sym_CR),
	9164:  uint16(anon_sym_CR_LF),
	9165:  uint16(2),
	9166:  uint16(3),
	9167:  uint16(1),
	9168:  uint16(sym_comment),
	9169:  uint16(878),
	9170:  uint16(14),
	9171:  uint16(anon_sym_PIPE_PIPE),
	9172:  uint16(anon_sym_AMP_AMP),
	9173:  uint16(anon_sym_AMP_PIPE),
	9174:  uint16(anon_sym_2_GT_PIPE),
	9175:  uint16(anon_sym_PIPE),
	9176:  uint16(anon_sym_SEMI),
	9177:  uint16(anon_sym_AMP),
	9178:  uint16(anon_sym_LF),
	9179:  uint16(anon_sym_CR),
	9180:  uint16(anon_sym_CR_LF),
	9181:  uint16(anon_sym_RPAREN),
	9182:  uint16(anon_sym_RBRACE),
	9183:  uint16(sym_stream_redirect),
	9184:  uint16(sym_direction),
	9185:  uint16(2),
	9186:  uint16(3),
	9187:  uint16(1),
	9188:  uint16(sym_comment),
	9189:  uint16(874),
	9190:  uint16(14),
	9191:  uint16(anon_sym_PIPE_PIPE),
	9192:  uint16(anon_sym_AMP_AMP),
	9193:  uint16(anon_sym_AMP_PIPE),
	9194:  uint16(anon_sym_2_GT_PIPE),
	9195:  uint16(anon_sym_PIPE),
	9196:  uint16(anon_sym_SEMI),
	9197:  uint16(anon_sym_AMP),
	9198:  uint16(anon_sym_LF),
	9199:  uint16(anon_sym_CR),
	9200:  uint16(anon_sym_CR_LF),
	9201:  uint16(anon_sym_RPAREN),
	9202:  uint16(anon_sym_RBRACE),
	9203:  uint16(sym_stream_redirect),
	9204:  uint16(sym_direction),
	9205:  uint16(8),
	9206:  uint16(3),
	9207:  uint16(1),
	9208:  uint16(sym_comment),
	9209:  uint16(838),
	9210:  uint16(1),
	9211:  uint16(sym_stream_redirect),
	9212:  uint16(840),
	9213:  uint16(1),
	9214:  uint16(sym_direction),
	9215:  uint16(880),
	9216:  uint16(1),
	9217:  uint16(anon_sym_RBRACE),
	9218:  uint16(251),
	9219:  uint16(1),
	9220:  uint16(sym_file_redirect),
	9221:  uint16(832),
	9222:  uint16(2),
	9223:  uint16(anon_sym_PIPE_PIPE),
	9224:  uint16(anon_sym_AMP_AMP),
	9225:  uint16(834),
	9226:  uint16(3),
	9227:  uint16(anon_sym_AMP_PIPE),
	9228:  uint16(anon_sym_2_GT_PIPE),
	9229:  uint16(anon_sym_PIPE),
	9230:  uint16(866),
	9231:  uint16(5),
	9232:  uint16(anon_sym_SEMI),
	9233:  uint16(anon_sym_AMP),
	9234:  uint16(anon_sym_LF),
	9235:  uint16(anon_sym_CR),
	9236:  uint16(anon_sym_CR_LF),
	9237:  uint16(2),
	9238:  uint16(3),
	9239:  uint16(1),
	9240:  uint16(sym_comment),
	9241:  uint16(882),
	9242:  uint16(14),
	9243:  uint16(anon_sym_PIPE_PIPE),
	9244:  uint16(anon_sym_AMP_AMP),
	9245:  uint16(anon_sym_AMP_PIPE),
	9246:  uint16(anon_sym_2_GT_PIPE),
	9247:  uint16(anon_sym_PIPE),
	9248:  uint16(anon_sym_SEMI),
	9249:  uint16(anon_sym_AMP),
	9250:  uint16(anon_sym_LF),
	9251:  uint16(anon_sym_CR),
	9252:  uint16(anon_sym_CR_LF),
	9253:  uint16(anon_sym_RPAREN),
	9254:  uint16(anon_sym_RBRACE),
	9255:  uint16(sym_stream_redirect),
	9256:  uint16(sym_direction),
	9257:  uint16(2),
	9258:  uint16(3),
	9259:  uint16(1),
	9260:  uint16(sym_comment),
	9261:  uint16(874),
	9262:  uint16(14),
	9263:  uint16(anon_sym_PIPE_PIPE),
	9264:  uint16(anon_sym_AMP_AMP),
	9265:  uint16(anon_sym_AMP_PIPE),
	9266:  uint16(anon_sym_2_GT_PIPE),
	9267:  uint16(anon_sym_PIPE),
	9268:  uint16(anon_sym_SEMI),
	9269:  uint16(anon_sym_AMP),
	9270:  uint16(anon_sym_LF),
	9271:  uint16(anon_sym_CR),
	9272:  uint16(anon_sym_CR_LF),
	9273:  uint16(anon_sym_RPAREN),
	9274:  uint16(anon_sym_RBRACE),
	9275:  uint16(sym_stream_redirect),
	9276:  uint16(sym_direction),
	9277:  uint16(8),
	9278:  uint16(3),
	9279:  uint16(1),
	9280:  uint16(sym_comment),
	9281:  uint16(838),
	9282:  uint16(1),
	9283:  uint16(sym_stream_redirect),
	9284:  uint16(840),
	9285:  uint16(1),
	9286:  uint16(sym_direction),
	9287:  uint16(884),
	9288:  uint16(1),
	9289:  uint16(anon_sym_RPAREN),
	9290:  uint16(251),
	9291:  uint16(1),
	9292:  uint16(sym_file_redirect),
	9293:  uint16(832),
	9294:  uint16(2),
	9295:  uint16(anon_sym_PIPE_PIPE),
	9296:  uint16(anon_sym_AMP_AMP),
	9297:  uint16(834),
	9298:  uint16(3),
	9299:  uint16(anon_sym_AMP_PIPE),
	9300:  uint16(anon_sym_2_GT_PIPE),
	9301:  uint16(anon_sym_PIPE),
	9302:  uint16(850),
	9303:  uint16(5),
	9304:  uint16(anon_sym_SEMI),
	9305:  uint16(anon_sym_AMP),
	9306:  uint16(anon_sym_LF),
	9307:  uint16(anon_sym_CR),
	9308:  uint16(anon_sym_CR_LF),
	9309:  uint16(2),
	9310:  uint16(3),
	9311:  uint16(1),
	9312:  uint16(sym_comment),
	9313:  uint16(874),
	9314:  uint16(14),
	9315:  uint16(anon_sym_PIPE_PIPE),
	9316:  uint16(anon_sym_AMP_AMP),
	9317:  uint16(anon_sym_AMP_PIPE),
	9318:  uint16(anon_sym_2_GT_PIPE),
	9319:  uint16(anon_sym_PIPE),
	9320:  uint16(anon_sym_SEMI),
	9321:  uint16(anon_sym_AMP),
	9322:  uint16(anon_sym_LF),
	9323:  uint16(anon_sym_CR),
	9324:  uint16(anon_sym_CR_LF),
	9325:  uint16(anon_sym_RPAREN),
	9326:  uint16(anon_sym_RBRACE),
	9327:  uint16(sym_stream_redirect),
	9328:  uint16(sym_direction),
	9329:  uint16(2),
	9330:  uint16(3),
	9331:  uint16(1),
	9332:  uint16(sym_comment),
	9333:  uint16(886),
	9334:  uint16(14),
	9335:  uint16(anon_sym_PIPE_PIPE),
	9336:  uint16(anon_sym_AMP_AMP),
	9337:  uint16(anon_sym_AMP_PIPE),
	9338:  uint16(anon_sym_2_GT_PIPE),
	9339:  uint16(anon_sym_PIPE),
	9340:  uint16(anon_sym_SEMI),
	9341:  uint16(anon_sym_AMP),
	9342:  uint16(anon_sym_LF),
	9343:  uint16(anon_sym_CR),
	9344:  uint16(anon_sym_CR_LF),
	9345:  uint16(anon_sym_RPAREN),
	9346:  uint16(anon_sym_RBRACE),
	9347:  uint16(sym_stream_redirect),
	9348:  uint16(sym_direction),
	9349:  uint16(2),
	9350:  uint16(3),
	9351:  uint16(1),
	9352:  uint16(sym_comment),
	9353:  uint16(888),
	9354:  uint16(14),
	9355:  uint16(anon_sym_PIPE_PIPE),
	9356:  uint16(anon_sym_AMP_AMP),
	9357:  uint16(anon_sym_AMP_PIPE),
	9358:  uint16(anon_sym_2_GT_PIPE),
	9359:  uint16(anon_sym_PIPE),
	9360:  uint16(anon_sym_SEMI),
	9361:  uint16(anon_sym_AMP),
	9362:  uint16(anon_sym_LF),
	9363:  uint16(anon_sym_CR),
	9364:  uint16(anon_sym_CR_LF),
	9365:  uint16(anon_sym_RPAREN),
	9366:  uint16(anon_sym_RBRACE),
	9367:  uint16(sym_stream_redirect),
	9368:  uint16(sym_direction),
	9369:  uint16(8),
	9370:  uint16(3),
	9371:  uint16(1),
	9372:  uint16(sym_comment),
	9373:  uint16(275),
	9374:  uint16(1),
	9375:  uint16(anon_sym_RPAREN),
	9376:  uint16(838),
	9377:  uint16(1),
	9378:  uint16(sym_stream_redirect),
	9379:  uint16(840),
	9380:  uint16(1),
	9381:  uint16(sym_direction),
	9382:  uint16(251),
	9383:  uint16(1),
	9384:  uint16(sym_file_redirect),
	9385:  uint16(832),
	9386:  uint16(2),
	9387:  uint16(anon_sym_PIPE_PIPE),
	9388:  uint16(anon_sym_AMP_AMP),
	9389:  uint16(834),
	9390:  uint16(3),
	9391:  uint16(anon_sym_AMP_PIPE),
	9392:  uint16(anon_sym_2_GT_PIPE),
	9393:  uint16(anon_sym_PIPE),
	9394:  uint16(850),
	9395:  uint16(5),
	9396:  uint16(anon_sym_SEMI),
	9397:  uint16(anon_sym_AMP),
	9398:  uint16(anon_sym_LF),
	9399:  uint16(anon_sym_CR),
	9400:  uint16(anon_sym_CR_LF),
	9401:  uint16(2),
	9402:  uint16(3),
	9403:  uint16(1),
	9404:  uint16(sym_comment),
	9405:  uint16(890),
	9406:  uint16(14),
	9407:  uint16(anon_sym_PIPE_PIPE),
	9408:  uint16(anon_sym_AMP_AMP),
	9409:  uint16(anon_sym_AMP_PIPE),
	9410:  uint16(anon_sym_2_GT_PIPE),
	9411:  uint16(anon_sym_PIPE),
	9412:  uint16(anon_sym_SEMI),
	9413:  uint16(anon_sym_AMP),
	9414:  uint16(anon_sym_LF),
	9415:  uint16(anon_sym_CR),
	9416:  uint16(anon_sym_CR_LF),
	9417:  uint16(anon_sym_RPAREN),
	9418:  uint16(anon_sym_RBRACE),
	9419:  uint16(sym_stream_redirect),
	9420:  uint16(sym_direction),
	9421:  uint16(2),
	9422:  uint16(3),
	9423:  uint16(1),
	9424:  uint16(sym_comment),
	9425:  uint16(890),
	9426:  uint16(14),
	9427:  uint16(anon_sym_PIPE_PIPE),
	9428:  uint16(anon_sym_AMP_AMP),
	9429:  uint16(anon_sym_AMP_PIPE),
	9430:  uint16(anon_sym_2_GT_PIPE),
	9431:  uint16(anon_sym_PIPE),
	9432:  uint16(anon_sym_SEMI),
	9433:  uint16(anon_sym_AMP),
	9434:  uint16(anon_sym_LF),
	9435:  uint16(anon_sym_CR),
	9436:  uint16(anon_sym_CR_LF),
	9437:  uint16(anon_sym_RPAREN),
	9438:  uint16(anon_sym_RBRACE),
	9439:  uint16(sym_stream_redirect),
	9440:  uint16(sym_direction),
	9441:  uint16(2),
	9442:  uint16(3),
	9443:  uint16(1),
	9444:  uint16(sym_comment),
	9445:  uint16(890),
	9446:  uint16(14),
	9447:  uint16(anon_sym_PIPE_PIPE),
	9448:  uint16(anon_sym_AMP_AMP),
	9449:  uint16(anon_sym_AMP_PIPE),
	9450:  uint16(anon_sym_2_GT_PIPE),
	9451:  uint16(anon_sym_PIPE),
	9452:  uint16(anon_sym_SEMI),
	9453:  uint16(anon_sym_AMP),
	9454:  uint16(anon_sym_LF),
	9455:  uint16(anon_sym_CR),
	9456:  uint16(anon_sym_CR_LF),
	9457:  uint16(anon_sym_RPAREN),
	9458:  uint16(anon_sym_RBRACE),
	9459:  uint16(sym_stream_redirect),
	9460:  uint16(sym_direction),
	9461:  uint16(2),
	9462:  uint16(3),
	9463:  uint16(1),
	9464:  uint16(sym_comment),
	9465:  uint16(892),
	9466:  uint16(14),
	9467:  uint16(anon_sym_PIPE_PIPE),
	9468:  uint16(anon_sym_AMP_AMP),
	9469:  uint16(anon_sym_AMP_PIPE),
	9470:  uint16(anon_sym_2_GT_PIPE),
	9471:  uint16(anon_sym_PIPE),
	9472:  uint16(anon_sym_SEMI),
	9473:  uint16(anon_sym_AMP),
	9474:  uint16(anon_sym_LF),
	9475:  uint16(anon_sym_CR),
	9476:  uint16(anon_sym_CR_LF),
	9477:  uint16(anon_sym_RPAREN),
	9478:  uint16(anon_sym_RBRACE),
	9479:  uint16(sym_stream_redirect),
	9480:  uint16(sym_direction),
	9481:  uint16(2),
	9482:  uint16(3),
	9483:  uint16(1),
	9484:  uint16(sym_comment),
	9485:  uint16(890),
	9486:  uint16(14),
	9487:  uint16(anon_sym_PIPE_PIPE),
	9488:  uint16(anon_sym_AMP_AMP),
	9489:  uint16(anon_sym_AMP_PIPE),
	9490:  uint16(anon_sym_2_GT_PIPE),
	9491:  uint16(anon_sym_PIPE),
	9492:  uint16(anon_sym_SEMI),
	9493:  uint16(anon_sym_AMP),
	9494:  uint16(anon_sym_LF),
	9495:  uint16(anon_sym_CR),
	9496:  uint16(anon_sym_CR_LF),
	9497:  uint16(anon_sym_RPAREN),
	9498:  uint16(anon_sym_RBRACE),
	9499:  uint16(sym_stream_redirect),
	9500:  uint16(sym_direction),
	9501:  uint16(2),
	9502:  uint16(3),
	9503:  uint16(1),
	9504:  uint16(sym_comment),
	9505:  uint16(890),
	9506:  uint16(14),
	9507:  uint16(anon_sym_PIPE_PIPE),
	9508:  uint16(anon_sym_AMP_AMP),
	9509:  uint16(anon_sym_AMP_PIPE),
	9510:  uint16(anon_sym_2_GT_PIPE),
	9511:  uint16(anon_sym_PIPE),
	9512:  uint16(anon_sym_SEMI),
	9513:  uint16(anon_sym_AMP),
	9514:  uint16(anon_sym_LF),
	9515:  uint16(anon_sym_CR),
	9516:  uint16(anon_sym_CR_LF),
	9517:  uint16(anon_sym_RPAREN),
	9518:  uint16(anon_sym_RBRACE),
	9519:  uint16(sym_stream_redirect),
	9520:  uint16(sym_direction),
	9521:  uint16(8),
	9522:  uint16(3),
	9523:  uint16(1),
	9524:  uint16(sym_comment),
	9525:  uint16(265),
	9526:  uint16(1),
	9527:  uint16(anon_sym_RPAREN),
	9528:  uint16(838),
	9529:  uint16(1),
	9530:  uint16(sym_stream_redirect),
	9531:  uint16(840),
	9532:  uint16(1),
	9533:  uint16(sym_direction),
	9534:  uint16(251),
	9535:  uint16(1),
	9536:  uint16(sym_file_redirect),
	9537:  uint16(832),
	9538:  uint16(2),
	9539:  uint16(anon_sym_PIPE_PIPE),
	9540:  uint16(anon_sym_AMP_AMP),
	9541:  uint16(834),
	9542:  uint16(3),
	9543:  uint16(anon_sym_AMP_PIPE),
	9544:  uint16(anon_sym_2_GT_PIPE),
	9545:  uint16(anon_sym_PIPE),
	9546:  uint16(850),
	9547:  uint16(5),
	9548:  uint16(anon_sym_SEMI),
	9549:  uint16(anon_sym_AMP),
	9550:  uint16(anon_sym_LF),
	9551:  uint16(anon_sym_CR),
	9552:  uint16(anon_sym_CR_LF),
	9553:  uint16(10),
	9554:  uint16(765),
	9555:  uint16(1),
	9556:  uint16(anon_sym_DOLLAR),
	9557:  uint16(767),
	9558:  uint16(1),
	9559:  uint16(anon_sym_LPAREN),
	9560:  uint16(769),
	9561:  uint16(1),
	9562:  uint16(sym_integer),
	9563:  uint16(771),
	9564:  uint16(1),
	9565:  uint16(sym_comment),
	9566:  uint16(777),
	9567:  uint16(1),
	9568:  uint16(anon_sym_DQUOTE),
	9569:  uint16(779),
	9570:  uint16(1),
	9571:  uint16(anon_sym_SQUOTE),
	9572:  uint16(291),
	9573:  uint16(1),
	9574:  uint16(sym_index),
	9575:  uint16(894),
	9576:  uint16(2),
	9577:  uint16(anon_sym_DOT_DOT),
	9578:  uint16(anon_sym_RBRACK),
	9579:  uint16(297),
	9580:  uint16(2),
	9581:  uint16(sym__command_substitution_dollar),
	9582:  uint16(sym__command_substitution_inner),
	9583:  uint16(285),
	9584:  uint16(4),
	9585:  uint16(sym_command_substitution),
	9586:  uint16(sym_variable_expansion),
	9587:  uint16(sym_double_quote_string),
	9588:  uint16(sym_single_quote_string),
	9589:  uint16(8),
	9590:  uint16(3),
	9591:  uint16(1),
	9592:  uint16(sym_comment),
	9593:  uint16(838),
	9594:  uint16(1),
	9595:  uint16(sym_stream_redirect),
	9596:  uint16(840),
	9597:  uint16(1),
	9598:  uint16(sym_direction),
	9599:  uint16(896),
	9600:  uint16(1),
	9601:  uint16(anon_sym_RPAREN),
	9602:  uint16(251),
	9603:  uint16(1),
	9604:  uint16(sym_file_redirect),
	9605:  uint16(832),
	9606:  uint16(2),
	9607:  uint16(anon_sym_PIPE_PIPE),
	9608:  uint16(anon_sym_AMP_AMP),
	9609:  uint16(834),
	9610:  uint16(3),
	9611:  uint16(anon_sym_AMP_PIPE),
	9612:  uint16(anon_sym_2_GT_PIPE),
	9613:  uint16(anon_sym_PIPE),
	9614:  uint16(850),
	9615:  uint16(5),
	9616:  uint16(anon_sym_SEMI),
	9617:  uint16(anon_sym_AMP),
	9618:  uint16(anon_sym_LF),
	9619:  uint16(anon_sym_CR),
	9620:  uint16(anon_sym_CR_LF),
	9621:  uint16(2),
	9622:  uint16(3),
	9623:  uint16(1),
	9624:  uint16(sym_comment),
	9625:  uint16(898),
	9626:  uint16(14),
	9627:  uint16(anon_sym_PIPE_PIPE),
	9628:  uint16(anon_sym_AMP_AMP),
	9629:  uint16(anon_sym_AMP_PIPE),
	9630:  uint16(anon_sym_2_GT_PIPE),
	9631:  uint16(anon_sym_PIPE),
	9632:  uint16(anon_sym_SEMI),
	9633:  uint16(anon_sym_AMP),
	9634:  uint16(anon_sym_LF),
	9635:  uint16(anon_sym_CR),
	9636:  uint16(anon_sym_CR_LF),
	9637:  uint16(anon_sym_RPAREN),
	9638:  uint16(anon_sym_RBRACE),
	9639:  uint16(sym_stream_redirect),
	9640:  uint16(sym_direction),
	9641:  uint16(8),
	9642:  uint16(3),
	9643:  uint16(1),
	9644:  uint16(sym_comment),
	9645:  uint16(178),
	9646:  uint16(1),
	9647:  uint16(anon_sym_RPAREN),
	9648:  uint16(838),
	9649:  uint16(1),
	9650:  uint16(sym_stream_redirect),
	9651:  uint16(840),
	9652:  uint16(1),
	9653:  uint16(sym_direction),
	9654:  uint16(251),
	9655:  uint16(1),
	9656:  uint16(sym_file_redirect),
	9657:  uint16(832),
	9658:  uint16(2),
	9659:  uint16(anon_sym_PIPE_PIPE),
	9660:  uint16(anon_sym_AMP_AMP),
	9661:  uint16(834),
	9662:  uint16(3),
	9663:  uint16(anon_sym_AMP_PIPE),
	9664:  uint16(anon_sym_2_GT_PIPE),
	9665:  uint16(anon_sym_PIPE),
	9666:  uint16(850),
	9667:  uint16(5),
	9668:  uint16(anon_sym_SEMI),
	9669:  uint16(anon_sym_AMP),
	9670:  uint16(anon_sym_LF),
	9671:  uint16(anon_sym_CR),
	9672:  uint16(anon_sym_CR_LF),
	9673:  uint16(2),
	9674:  uint16(3),
	9675:  uint16(1),
	9676:  uint16(sym_comment),
	9677:  uint16(900),
	9678:  uint16(14),
	9679:  uint16(anon_sym_PIPE_PIPE),
	9680:  uint16(anon_sym_AMP_AMP),
	9681:  uint16(anon_sym_AMP_PIPE),
	9682:  uint16(anon_sym_2_GT_PIPE),
	9683:  uint16(anon_sym_PIPE),
	9684:  uint16(anon_sym_SEMI),
	9685:  uint16(anon_sym_AMP),
	9686:  uint16(anon_sym_LF),
	9687:  uint16(anon_sym_CR),
	9688:  uint16(anon_sym_CR_LF),
	9689:  uint16(anon_sym_RPAREN),
	9690:  uint16(anon_sym_RBRACE),
	9691:  uint16(sym_stream_redirect),
	9692:  uint16(sym_direction),
	9693:  uint16(10),
	9694:  uint16(765),
	9695:  uint16(1),
	9696:  uint16(anon_sym_DOLLAR),
	9697:  uint16(767),
	9698:  uint16(1),
	9699:  uint16(anon_sym_LPAREN),
	9700:  uint16(769),
	9701:  uint16(1),
	9702:  uint16(sym_integer),
	9703:  uint16(771),
	9704:  uint16(1),
	9705:  uint16(sym_comment),
	9706:  uint16(777),
	9707:  uint16(1),
	9708:  uint16(anon_sym_DQUOTE),
	9709:  uint16(779),
	9710:  uint16(1),
	9711:  uint16(anon_sym_SQUOTE),
	9712:  uint16(286),
	9713:  uint16(1),
	9714:  uint16(sym_index),
	9715:  uint16(902),
	9716:  uint16(2),
	9717:  uint16(anon_sym_DOT_DOT),
	9718:  uint16(anon_sym_RBRACK),
	9719:  uint16(297),
	9720:  uint16(2),
	9721:  uint16(sym__command_substitution_dollar),
	9722:  uint16(sym__command_substitution_inner),
	9723:  uint16(285),
	9724:  uint16(4),
	9725:  uint16(sym_command_substitution),
	9726:  uint16(sym_variable_expansion),
	9727:  uint16(sym_double_quote_string),
	9728:  uint16(sym_single_quote_string),
	9729:  uint16(8),
	9730:  uint16(3),
	9731:  uint16(1),
	9732:  uint16(sym_comment),
	9733:  uint16(303),
	9734:  uint16(1),
	9735:  uint16(anon_sym_RPAREN),
	9736:  uint16(838),
	9737:  uint16(1),
	9738:  uint16(sym_stream_redirect),
	9739:  uint16(840),
	9740:  uint16(1),
	9741:  uint16(sym_direction),
	9742:  uint16(251),
	9743:  uint16(1),
	9744:  uint16(sym_file_redirect),
	9745:  uint16(832),
	9746:  uint16(2),
	9747:  uint16(anon_sym_PIPE_PIPE),
	9748:  uint16(anon_sym_AMP_AMP),
	9749:  uint16(834),
	9750:  uint16(3),
	9751:  uint16(anon_sym_AMP_PIPE),
	9752:  uint16(anon_sym_2_GT_PIPE),
	9753:  uint16(anon_sym_PIPE),
	9754:  uint16(850),
	9755:  uint16(5),
	9756:  uint16(anon_sym_SEMI),
	9757:  uint16(anon_sym_AMP),
	9758:  uint16(anon_sym_LF),
	9759:  uint16(anon_sym_CR),
	9760:  uint16(anon_sym_CR_LF),
	9761:  uint16(2),
	9762:  uint16(3),
	9763:  uint16(1),
	9764:  uint16(sym_comment),
	9765:  uint16(854),
	9766:  uint16(14),
	9767:  uint16(anon_sym_PIPE_PIPE),
	9768:  uint16(anon_sym_AMP_AMP),
	9769:  uint16(anon_sym_AMP_PIPE),
	9770:  uint16(anon_sym_2_GT_PIPE),
	9771:  uint16(anon_sym_PIPE),
	9772:  uint16(anon_sym_SEMI),
	9773:  uint16(anon_sym_AMP),
	9774:  uint16(anon_sym_LF),
	9775:  uint16(anon_sym_CR),
	9776:  uint16(anon_sym_CR_LF),
	9777:  uint16(anon_sym_RPAREN),
	9778:  uint16(anon_sym_RBRACE),
	9779:  uint16(sym_stream_redirect),
	9780:  uint16(sym_direction),
	9781:  uint16(2),
	9782:  uint16(3),
	9783:  uint16(1),
	9784:  uint16(sym_comment),
	9785:  uint16(904),
	9786:  uint16(14),
	9787:  uint16(anon_sym_PIPE_PIPE),
	9788:  uint16(anon_sym_AMP_AMP),
	9789:  uint16(anon_sym_AMP_PIPE),
	9790:  uint16(anon_sym_2_GT_PIPE),
	9791:  uint16(anon_sym_PIPE),
	9792:  uint16(anon_sym_SEMI),
	9793:  uint16(anon_sym_AMP),
	9794:  uint16(anon_sym_LF),
	9795:  uint16(anon_sym_CR),
	9796:  uint16(anon_sym_CR_LF),
	9797:  uint16(anon_sym_RPAREN),
	9798:  uint16(anon_sym_RBRACE),
	9799:  uint16(sym_stream_redirect),
	9800:  uint16(sym_direction),
	9801:  uint16(8),
	9802:  uint16(3),
	9803:  uint16(1),
	9804:  uint16(sym_comment),
	9805:  uint16(838),
	9806:  uint16(1),
	9807:  uint16(sym_stream_redirect),
	9808:  uint16(840),
	9809:  uint16(1),
	9810:  uint16(sym_direction),
	9811:  uint16(906),
	9812:  uint16(1),
	9813:  uint16(anon_sym_RPAREN),
	9814:  uint16(251),
	9815:  uint16(1),
	9816:  uint16(sym_file_redirect),
	9817:  uint16(832),
	9818:  uint16(2),
	9819:  uint16(anon_sym_PIPE_PIPE),
	9820:  uint16(anon_sym_AMP_AMP),
	9821:  uint16(834),
	9822:  uint16(3),
	9823:  uint16(anon_sym_AMP_PIPE),
	9824:  uint16(anon_sym_2_GT_PIPE),
	9825:  uint16(anon_sym_PIPE),
	9826:  uint16(850),
	9827:  uint16(5),
	9828:  uint16(anon_sym_SEMI),
	9829:  uint16(anon_sym_AMP),
	9830:  uint16(anon_sym_LF),
	9831:  uint16(anon_sym_CR),
	9832:  uint16(anon_sym_CR_LF),
	9833:  uint16(2),
	9834:  uint16(3),
	9835:  uint16(1),
	9836:  uint16(sym_comment),
	9837:  uint16(854),
	9838:  uint16(14),
	9839:  uint16(anon_sym_PIPE_PIPE),
	9840:  uint16(anon_sym_AMP_AMP),
	9841:  uint16(anon_sym_AMP_PIPE),
	9842:  uint16(anon_sym_2_GT_PIPE),
	9843:  uint16(anon_sym_PIPE),
	9844:  uint16(anon_sym_SEMI),
	9845:  uint16(anon_sym_AMP),
	9846:  uint16(anon_sym_LF),
	9847:  uint16(anon_sym_CR),
	9848:  uint16(anon_sym_CR_LF),
	9849:  uint16(anon_sym_RPAREN),
	9850:  uint16(anon_sym_RBRACE),
	9851:  uint16(sym_stream_redirect),
	9852:  uint16(sym_direction),
	9853:  uint16(2),
	9854:  uint16(3),
	9855:  uint16(1),
	9856:  uint16(sym_comment),
	9857:  uint16(854),
	9858:  uint16(14),
	9859:  uint16(anon_sym_PIPE_PIPE),
	9860:  uint16(anon_sym_AMP_AMP),
	9861:  uint16(anon_sym_AMP_PIPE),
	9862:  uint16(anon_sym_2_GT_PIPE),
	9863:  uint16(anon_sym_PIPE),
	9864:  uint16(anon_sym_SEMI),
	9865:  uint16(anon_sym_AMP),
	9866:  uint16(anon_sym_LF),
	9867:  uint16(anon_sym_CR),
	9868:  uint16(anon_sym_CR_LF),
	9869:  uint16(anon_sym_RPAREN),
	9870:  uint16(anon_sym_RBRACE),
	9871:  uint16(sym_stream_redirect),
	9872:  uint16(sym_direction),
	9873:  uint16(8),
	9874:  uint16(3),
	9875:  uint16(1),
	9876:  uint16(sym_comment),
	9877:  uint16(838),
	9878:  uint16(1),
	9879:  uint16(sym_stream_redirect),
	9880:  uint16(840),
	9881:  uint16(1),
	9882:  uint16(sym_direction),
	9883:  uint16(908),
	9884:  uint16(1),
	9885:  uint16(anon_sym_RPAREN),
	9886:  uint16(251),
	9887:  uint16(1),
	9888:  uint16(sym_file_redirect),
	9889:  uint16(832),
	9890:  uint16(2),
	9891:  uint16(anon_sym_PIPE_PIPE),
	9892:  uint16(anon_sym_AMP_AMP),
	9893:  uint16(834),
	9894:  uint16(3),
	9895:  uint16(anon_sym_AMP_PIPE),
	9896:  uint16(anon_sym_2_GT_PIPE),
	9897:  uint16(anon_sym_PIPE),
	9898:  uint16(850),
	9899:  uint16(5),
	9900:  uint16(anon_sym_SEMI),
	9901:  uint16(anon_sym_AMP),
	9902:  uint16(anon_sym_LF),
	9903:  uint16(anon_sym_CR),
	9904:  uint16(anon_sym_CR_LF),
	9905:  uint16(7),
	9906:  uint16(3),
	9907:  uint16(1),
	9908:  uint16(sym_comment),
	9909:  uint16(838),
	9910:  uint16(1),
	9911:  uint16(sym_stream_redirect),
	9912:  uint16(840),
	9913:  uint16(1),
	9914:  uint16(sym_direction),
	9915:  uint16(251),
	9916:  uint16(1),
	9917:  uint16(sym_file_redirect),
	9918:  uint16(832),
	9919:  uint16(2),
	9920:  uint16(anon_sym_PIPE_PIPE),
	9921:  uint16(anon_sym_AMP_AMP),
	9922:  uint16(834),
	9923:  uint16(3),
	9924:  uint16(anon_sym_AMP_PIPE),
	9925:  uint16(anon_sym_2_GT_PIPE),
	9926:  uint16(anon_sym_PIPE),
	9927:  uint16(910),
	9928:  uint16(5),
	9929:  uint16(anon_sym_SEMI),
	9930:  uint16(anon_sym_AMP),
	9931:  uint16(anon_sym_LF),
	9932:  uint16(anon_sym_CR),
	9933:  uint16(anon_sym_CR_LF),
	9934:  uint16(7),
	9935:  uint16(3),
	9936:  uint16(1),
	9937:  uint16(sym_comment),
	9938:  uint16(838),
	9939:  uint16(1),
	9940:  uint16(sym_stream_redirect),
	9941:  uint16(840),
	9942:  uint16(1),
	9943:  uint16(sym_direction),
	9944:  uint16(251),
	9945:  uint16(1),
	9946:  uint16(sym_file_redirect),
	9947:  uint16(832),
	9948:  uint16(2),
	9949:  uint16(anon_sym_PIPE_PIPE),
	9950:  uint16(anon_sym_AMP_AMP),
	9951:  uint16(834),
	9952:  uint16(3),
	9953:  uint16(anon_sym_AMP_PIPE),
	9954:  uint16(anon_sym_2_GT_PIPE),
	9955:  uint16(anon_sym_PIPE),
	9956:  uint16(912),
	9957:  uint16(5),
	9958:  uint16(anon_sym_SEMI),
	9959:  uint16(anon_sym_AMP),
	9960:  uint16(anon_sym_LF),
	9961:  uint16(anon_sym_CR),
	9962:  uint16(anon_sym_CR_LF),
	9963:  uint16(7),
	9964:  uint16(3),
	9965:  uint16(1),
	9966:  uint16(sym_comment),
	9967:  uint16(838),
	9968:  uint16(1),
	9969:  uint16(sym_stream_redirect),
	9970:  uint16(840),
	9971:  uint16(1),
	9972:  uint16(sym_direction),
	9973:  uint16(251),
	9974:  uint16(1),
	9975:  uint16(sym_file_redirect),
	9976:  uint16(832),
	9977:  uint16(2),
	9978:  uint16(anon_sym_PIPE_PIPE),
	9979:  uint16(anon_sym_AMP_AMP),
	9980:  uint16(834),
	9981:  uint16(3),
	9982:  uint16(anon_sym_AMP_PIPE),
	9983:  uint16(anon_sym_2_GT_PIPE),
	9984:  uint16(anon_sym_PIPE),
	9985:  uint16(914),
	9986:  uint16(5),
	9987:  uint16(anon_sym_SEMI),
	9988:  uint16(anon_sym_AMP),
	9989:  uint16(anon_sym_LF),
	9990:  uint16(anon_sym_CR),
	9991:  uint16(anon_sym_CR_LF),
	9992:  uint16(7),
	9993:  uint16(3),
	9994:  uint16(1),
	9995:  uint16(sym_comment),
	9996:  uint16(838),
	9997:  uint16(1),
	9998:  uint16(sym_stream_redirect),
	9999:  uint16(840),
	10000: uint16(1),
	10001: uint16(sym_direction),
	10002: uint16(251),
	10003: uint16(1),
	10004: uint16(sym_file_redirect),
	10005: uint16(832),
	10006: uint16(2),
	10007: uint16(anon_sym_PIPE_PIPE),
	10008: uint16(anon_sym_AMP_AMP),
	10009: uint16(834),
	10010: uint16(3),
	10011: uint16(anon_sym_AMP_PIPE),
	10012: uint16(anon_sym_2_GT_PIPE),
	10013: uint16(anon_sym_PIPE),
	10014: uint16(916),
	10015: uint16(5),
	10016: uint16(anon_sym_SEMI),
	10017: uint16(anon_sym_AMP),
	10018: uint16(anon_sym_LF),
	10019: uint16(anon_sym_CR),
	10020: uint16(anon_sym_CR_LF),
	10021: uint16(7),
	10022: uint16(3),
	10023: uint16(1),
	10024: uint16(sym_comment),
	10025: uint16(838),
	10026: uint16(1),
	10027: uint16(sym_stream_redirect),
	10028: uint16(840),
	10029: uint16(1),
	10030: uint16(sym_direction),
	10031: uint16(251),
	10032: uint16(1),
	10033: uint16(sym_file_redirect),
	10034: uint16(832),
	10035: uint16(2),
	10036: uint16(anon_sym_PIPE_PIPE),
	10037: uint16(anon_sym_AMP_AMP),
	10038: uint16(834),
	10039: uint16(3),
	10040: uint16(anon_sym_AMP_PIPE),
	10041: uint16(anon_sym_2_GT_PIPE),
	10042: uint16(anon_sym_PIPE),
	10043: uint16(850),
	10044: uint16(5),
	10045: uint16(anon_sym_SEMI),
	10046: uint16(anon_sym_AMP),
	10047: uint16(anon_sym_LF),
	10048: uint16(anon_sym_CR),
	10049: uint16(anon_sym_CR_LF),
	10050: uint16(7),
	10051: uint16(3),
	10052: uint16(1),
	10053: uint16(sym_comment),
	10054: uint16(838),
	10055: uint16(1),
	10056: uint16(sym_stream_redirect),
	10057: uint16(840),
	10058: uint16(1),
	10059: uint16(sym_direction),
	10060: uint16(251),
	10061: uint16(1),
	10062: uint16(sym_file_redirect),
	10063: uint16(832),
	10064: uint16(2),
	10065: uint16(anon_sym_PIPE_PIPE),
	10066: uint16(anon_sym_AMP_AMP),
	10067: uint16(834),
	10068: uint16(3),
	10069: uint16(anon_sym_AMP_PIPE),
	10070: uint16(anon_sym_2_GT_PIPE),
	10071: uint16(anon_sym_PIPE),
	10072: uint16(866),
	10073: uint16(5),
	10074: uint16(anon_sym_SEMI),
	10075: uint16(anon_sym_AMP),
	10076: uint16(anon_sym_LF),
	10077: uint16(anon_sym_CR),
	10078: uint16(anon_sym_CR_LF),
	10079: uint16(7),
	10080: uint16(3),
	10081: uint16(1),
	10082: uint16(sym_comment),
	10083: uint16(838),
	10084: uint16(1),
	10085: uint16(sym_stream_redirect),
	10086: uint16(840),
	10087: uint16(1),
	10088: uint16(sym_direction),
	10089: uint16(251),
	10090: uint16(1),
	10091: uint16(sym_file_redirect),
	10092: uint16(832),
	10093: uint16(2),
	10094: uint16(anon_sym_PIPE_PIPE),
	10095: uint16(anon_sym_AMP_AMP),
	10096: uint16(834),
	10097: uint16(3),
	10098: uint16(anon_sym_AMP_PIPE),
	10099: uint16(anon_sym_2_GT_PIPE),
	10100: uint16(anon_sym_PIPE),
	10101: uint16(918),
	10102: uint16(5),
	10103: uint16(anon_sym_SEMI),
	10104: uint16(anon_sym_AMP),
	10105: uint16(anon_sym_LF),
	10106: uint16(anon_sym_CR),
	10107: uint16(anon_sym_CR_LF),
	10108: uint16(7),
	10109: uint16(3),
	10110: uint16(1),
	10111: uint16(sym_comment),
	10112: uint16(838),
	10113: uint16(1),
	10114: uint16(sym_stream_redirect),
	10115: uint16(840),
	10116: uint16(1),
	10117: uint16(sym_direction),
	10118: uint16(251),
	10119: uint16(1),
	10120: uint16(sym_file_redirect),
	10121: uint16(832),
	10122: uint16(2),
	10123: uint16(anon_sym_PIPE_PIPE),
	10124: uint16(anon_sym_AMP_AMP),
	10125: uint16(834),
	10126: uint16(3),
	10127: uint16(anon_sym_AMP_PIPE),
	10128: uint16(anon_sym_2_GT_PIPE),
	10129: uint16(anon_sym_PIPE),
	10130: uint16(920),
	10131: uint16(5),
	10132: uint16(anon_sym_SEMI),
	10133: uint16(anon_sym_AMP),
	10134: uint16(anon_sym_LF),
	10135: uint16(anon_sym_CR),
	10136: uint16(anon_sym_CR_LF),
	10137: uint16(7),
	10138: uint16(3),
	10139: uint16(1),
	10140: uint16(sym_comment),
	10141: uint16(838),
	10142: uint16(1),
	10143: uint16(sym_stream_redirect),
	10144: uint16(840),
	10145: uint16(1),
	10146: uint16(sym_direction),
	10147: uint16(251),
	10148: uint16(1),
	10149: uint16(sym_file_redirect),
	10150: uint16(832),
	10151: uint16(2),
	10152: uint16(anon_sym_PIPE_PIPE),
	10153: uint16(anon_sym_AMP_AMP),
	10154: uint16(834),
	10155: uint16(3),
	10156: uint16(anon_sym_AMP_PIPE),
	10157: uint16(anon_sym_2_GT_PIPE),
	10158: uint16(anon_sym_PIPE),
	10159: uint16(922),
	10160: uint16(5),
	10161: uint16(anon_sym_SEMI),
	10162: uint16(anon_sym_AMP),
	10163: uint16(anon_sym_LF),
	10164: uint16(anon_sym_CR),
	10165: uint16(anon_sym_CR_LF),
	10166: uint16(3),
	10167: uint16(771),
	10168: uint16(1),
	10169: uint16(sym_comment),
	10170: uint16(271),
	10171: uint16(1),
	10172: uint16(aux_sym_variable_expansion_repeat1),
	10173: uint16(499),
	10174: uint16(8),
	10175: uint16(sym__concat_list),
	10176: uint16(anon_sym_DOLLAR),
	10177: uint16(anon_sym_LPAREN),
	10178: uint16(sym_integer),
	10179: uint16(anon_sym_DOT_DOT),
	10180: uint16(anon_sym_RBRACK),
	10181: uint16(anon_sym_DQUOTE),
	10182: uint16(anon_sym_SQUOTE),
	10183: uint16(4),
	10184: uint16(771),
	10185: uint16(1),
	10186: uint16(sym_comment),
	10187: uint16(924),
	10188: uint16(1),
	10189: uint16(sym__concat_list),
	10190: uint16(273),
	10191: uint16(1),
	10192: uint16(aux_sym_variable_expansion_repeat1),
	10193: uint16(493),
	10194: uint16(7),
	10195: uint16(anon_sym_DOLLAR),
	10196: uint16(anon_sym_LPAREN),
	10197: uint16(sym_integer),
	10198: uint16(anon_sym_DOT_DOT),
	10199: uint16(anon_sym_RBRACK),
	10200: uint16(anon_sym_DQUOTE),
	10201: uint16(anon_sym_SQUOTE),
	10202: uint16(3),
	10203: uint16(771),
	10204: uint16(1),
	10205: uint16(sym_comment),
	10206: uint16(273),
	10207: uint16(1),
	10208: uint16(aux_sym_variable_expansion_repeat1),
	10209: uint16(493),
	10210: uint16(8),
	10211: uint16(sym__concat_list),
	10212: uint16(anon_sym_DOLLAR),
	10213: uint16(anon_sym_LPAREN),
	10214: uint16(sym_integer),
	10215: uint16(anon_sym_DOT_DOT),
	10216: uint16(anon_sym_RBRACK),
	10217: uint16(anon_sym_DQUOTE),
	10218: uint16(anon_sym_SQUOTE),
	10219: uint16(4),
	10220: uint16(771),
	10221: uint16(1),
	10222: uint16(sym_comment),
	10223: uint16(924),
	10224: uint16(1),
	10225: uint16(sym__concat_list),
	10226: uint16(270),
	10227: uint16(1),
	10228: uint16(aux_sym_variable_expansion_repeat1),
	10229: uint16(499),
	10230: uint16(7),
	10231: uint16(anon_sym_DOLLAR),
	10232: uint16(anon_sym_LPAREN),
	10233: uint16(sym_integer),
	10234: uint16(anon_sym_DOT_DOT),
	10235: uint16(anon_sym_RBRACK),
	10236: uint16(anon_sym_DQUOTE),
	10237: uint16(anon_sym_SQUOTE),
	10238: uint16(4),
	10239: uint16(771),
	10240: uint16(1),
	10241: uint16(sym_comment),
	10242: uint16(926),
	10243: uint16(1),
	10244: uint16(sym__concat_list),
	10245: uint16(273),
	10246: uint16(1),
	10247: uint16(aux_sym_variable_expansion_repeat1),
	10248: uint16(503),
	10249: uint16(7),
	10250: uint16(anon_sym_DOLLAR),
	10251: uint16(anon_sym_LPAREN),
	10252: uint16(sym_integer),
	10253: uint16(anon_sym_DOT_DOT),
	10254: uint16(anon_sym_RBRACK),
	10255: uint16(anon_sym_DQUOTE),
	10256: uint16(anon_sym_SQUOTE),
	10257: uint16(4),
	10258: uint16(3),
	10259: uint16(1),
	10260: uint16(sym_comment),
	10261: uint16(275),
	10262: uint16(1),
	10263: uint16(aux_sym_variable_expansion_repeat1),
	10264: uint16(499),
	10265: uint16(2),
	10266: uint16(sym__concat),
	10267: uint16(sym__concat_list),
	10268: uint16(497),
	10269: uint16(5),
	10270: uint16(anon_sym_SEMI),
	10271: uint16(anon_sym_AMP),
	10272: uint16(anon_sym_LF),
	10273: uint16(anon_sym_CR),
	10274: uint16(anon_sym_CR_LF),
	10275: uint16(4),
	10276: uint16(3),
	10277: uint16(1),
	10278: uint16(sym_comment),
	10279: uint16(277),
	10280: uint16(1),
	10281: uint16(aux_sym_variable_expansion_repeat1),
	10282: uint16(493),
	10283: uint16(2),
	10284: uint16(sym__concat),
	10285: uint16(sym__concat_list),
	10286: uint16(491),
	10287: uint16(5),
	10288: uint16(anon_sym_SEMI),
	10289: uint16(anon_sym_AMP),
	10290: uint16(anon_sym_LF),
	10291: uint16(anon_sym_CR),
	10292: uint16(anon_sym_CR_LF),
	10293: uint16(5),
	10294: uint16(3),
	10295: uint16(1),
	10296: uint16(sym_comment),
	10297: uint16(493),
	10298: uint16(1),
	10299: uint16(sym__concat),
	10300: uint16(929),
	10301: uint16(1),
	10302: uint16(sym__concat_list),
	10303: uint16(277),
	10304: uint16(1),
	10305: uint16(aux_sym_variable_expansion_repeat1),
	10306: uint16(491),
	10307: uint16(5),
	10308: uint16(anon_sym_SEMI),
	10309: uint16(anon_sym_AMP),
	10310: uint16(anon_sym_LF),
	10311: uint16(anon_sym_CR),
	10312: uint16(anon_sym_CR_LF),
	10313: uint16(5),
	10314: uint16(3),
	10315: uint16(1),
	10316: uint16(sym_comment),
	10317: uint16(503),
	10318: uint16(1),
	10319: uint16(sym__concat),
	10320: uint16(931),
	10321: uint16(1),
	10322: uint16(sym__concat_list),
	10323: uint16(277),
	10324: uint16(1),
	10325: uint16(aux_sym_variable_expansion_repeat1),
	10326: uint16(501),
	10327: uint16(5),
	10328: uint16(anon_sym_SEMI),
	10329: uint16(anon_sym_AMP),
	10330: uint16(anon_sym_LF),
	10331: uint16(anon_sym_CR),
	10332: uint16(anon_sym_CR_LF),
	10333: uint16(2),
	10334: uint16(771),
	10335: uint16(1),
	10336: uint16(sym_comment),
	10337: uint16(574),
	10338: uint16(8),
	10339: uint16(sym__concat_list),
	10340: uint16(anon_sym_DOLLAR),
	10341: uint16(anon_sym_LPAREN),
	10342: uint16(sym_integer),
	10343: uint16(anon_sym_DOT_DOT),
	10344: uint16(anon_sym_RBRACK),
	10345: uint16(anon_sym_DQUOTE),
	10346: uint16(anon_sym_SQUOTE),
	10347: uint16(2),
	10348: uint16(771),
	10349: uint16(1),
	10350: uint16(sym_comment),
	10351: uint16(542),
	10352: uint16(8),
	10353: uint16(sym__concat_list),
	10354: uint16(anon_sym_DOLLAR),
	10355: uint16(anon_sym_LPAREN),
	10356: uint16(sym_integer),
	10357: uint16(anon_sym_DOT_DOT),
	10358: uint16(anon_sym_RBRACK),
	10359: uint16(anon_sym_DQUOTE),
	10360: uint16(anon_sym_SQUOTE),
	10361: uint16(5),
	10362: uint16(3),
	10363: uint16(1),
	10364: uint16(sym_comment),
	10365: uint16(499),
	10366: uint16(1),
	10367: uint16(sym__concat),
	10368: uint16(929),
	10369: uint16(1),
	10370: uint16(sym__concat_list),
	10371: uint16(276),
	10372: uint16(1),
	10373: uint16(aux_sym_variable_expansion_repeat1),
	10374: uint16(497),
	10375: uint16(5),
	10376: uint16(anon_sym_SEMI),
	10377: uint16(anon_sym_AMP),
	10378: uint16(anon_sym_LF),
	10379: uint16(anon_sym_CR),
	10380: uint16(anon_sym_CR_LF),
	10381: uint16(2),
	10382: uint16(771),
	10383: uint16(1),
	10384: uint16(sym_comment),
	10385: uint16(503),
	10386: uint16(8),
	10387: uint16(sym__concat_list),
	10388: uint16(anon_sym_DOLLAR),
	10389: uint16(anon_sym_LPAREN),
	10390: uint16(sym_integer),
	10391: uint16(anon_sym_DOT_DOT),
	10392: uint16(anon_sym_RBRACK),
	10393: uint16(anon_sym_DQUOTE),
	10394: uint16(anon_sym_SQUOTE),
	10395: uint16(7),
	10396: uint16(3),
	10397: uint16(1),
	10398: uint16(sym_comment),
	10399: uint16(934),
	10400: uint16(1),
	10401: uint16(anon_sym_DOLLAR),
	10402: uint16(936),
	10403: uint16(1),
	10404: uint16(anon_sym_DQUOTE),
	10405: uint16(938),
	10406: uint16(1),
	10407: uint16(aux_sym_double_quote_string_token1),
	10408: uint16(940),
	10409: uint16(1),
	10410: uint16(sym_escape_sequence),
	10411: uint16(364),
	10412: uint16(1),
	10413: uint16(sym__command_substitution_dollar),
	10414: uint16(298),
	10415: uint16(2),
	10416: uint16(sym_variable_expansion),
	10417: uint16(aux_sym_double_quote_string_repeat1),
	10418: uint16(3),
	10419: uint16(3),
	10420: uint16(1),
	10421: uint16(sym_comment),
	10422: uint16(542),
	10423: uint16(2),
	10424: uint16(sym__concat),
	10425: uint16(sym__concat_list),
	10426: uint16(540),
	10427: uint16(5),
	10428: uint16(anon_sym_SEMI),
	10429: uint16(anon_sym_AMP),
	10430: uint16(anon_sym_LF),
	10431: uint16(anon_sym_CR),
	10432: uint16(anon_sym_CR_LF),
	10433: uint16(3),
	10434: uint16(771),
	10435: uint16(1),
	10436: uint16(sym_comment),
	10437: uint16(944),
	10438: uint16(1),
	10439: uint16(anon_sym_DOT_DOT),
	10440: uint16(942),
	10441: uint16(6),
	10442: uint16(anon_sym_DOLLAR),
	10443: uint16(anon_sym_LPAREN),
	10444: uint16(sym_integer),
	10445: uint16(anon_sym_RBRACK),
	10446: uint16(anon_sym_DQUOTE),
	10447: uint16(anon_sym_SQUOTE),
	10448: uint16(2),
	10449: uint16(771),
	10450: uint16(1),
	10451: uint16(sym_comment),
	10452: uint16(946),
	10453: uint16(7),
	10454: uint16(anon_sym_DOLLAR),
	10455: uint16(anon_sym_LPAREN),
	10456: uint16(sym_integer),
	10457: uint16(anon_sym_DOT_DOT),
	10458: uint16(anon_sym_RBRACK),
	10459: uint16(anon_sym_DQUOTE),
	10460: uint16(anon_sym_SQUOTE),
	10461: uint16(2),
	10462: uint16(771),
	10463: uint16(1),
	10464: uint16(sym_comment),
	10465: uint16(894),
	10466: uint16(7),
	10467: uint16(anon_sym_DOLLAR),
	10468: uint16(anon_sym_LPAREN),
	10469: uint16(sym_integer),
	10470: uint16(anon_sym_DOT_DOT),
	10471: uint16(anon_sym_RBRACK),
	10472: uint16(anon_sym_DQUOTE),
	10473: uint16(anon_sym_SQUOTE),
	10474: uint16(7),
	10475: uint16(3),
	10476: uint16(1),
	10477: uint16(sym_comment),
	10478: uint16(934),
	10479: uint16(1),
	10480: uint16(anon_sym_DOLLAR),
	10481: uint16(948),
	10482: uint16(1),
	10483: uint16(anon_sym_DQUOTE),
	10484: uint16(950),
	10485: uint16(1),
	10486: uint16(aux_sym_double_quote_string_token1),
	10487: uint16(952),
	10488: uint16(1),
	10489: uint16(sym_escape_sequence),
	10490: uint16(364),
	10491: uint16(1),
	10492: uint16(sym__command_substitution_dollar),
	10493: uint16(293),
	10494: uint16(2),
	10495: uint16(sym_variable_expansion),
	10496: uint16(aux_sym_double_quote_string_repeat1),
	10497: uint16(7),
	10498: uint16(3),
	10499: uint16(1),
	10500: uint16(sym_comment),
	10501: uint16(934),
	10502: uint16(1),
	10503: uint16(anon_sym_DOLLAR),
	10504: uint16(954),
	10505: uint16(1),
	10506: uint16(anon_sym_DQUOTE),
	10507: uint16(956),
	10508: uint16(1),
	10509: uint16(aux_sym_double_quote_string_token1),
	10510: uint16(958),
	10511: uint16(1),
	10512: uint16(sym_escape_sequence),
	10513: uint16(364),
	10514: uint16(1),
	10515: uint16(sym__command_substitution_dollar),
	10516: uint16(299),
	10517: uint16(2),
	10518: uint16(sym_variable_expansion),
	10519: uint16(aux_sym_double_quote_string_repeat1),
	10520: uint16(7),
	10521: uint16(3),
	10522: uint16(1),
	10523: uint16(sym_comment),
	10524: uint16(934),
	10525: uint16(1),
	10526: uint16(anon_sym_DOLLAR),
	10527: uint16(960),
	10528: uint16(1),
	10529: uint16(anon_sym_DQUOTE),
	10530: uint16(962),
	10531: uint16(1),
	10532: uint16(aux_sym_double_quote_string_token1),
	10533: uint16(964),
	10534: uint16(1),
	10535: uint16(sym_escape_sequence),
	10536: uint16(364),
	10537: uint16(1),
	10538: uint16(sym__command_substitution_dollar),
	10539: uint16(287),
	10540: uint16(2),
	10541: uint16(sym_variable_expansion),
	10542: uint16(aux_sym_double_quote_string_repeat1),
	10543: uint16(7),
	10544: uint16(3),
	10545: uint16(1),
	10546: uint16(sym_comment),
	10547: uint16(934),
	10548: uint16(1),
	10549: uint16(anon_sym_DOLLAR),
	10550: uint16(950),
	10551: uint16(1),
	10552: uint16(aux_sym_double_quote_string_token1),
	10553: uint16(952),
	10554: uint16(1),
	10555: uint16(sym_escape_sequence),
	10556: uint16(966),
	10557: uint16(1),
	10558: uint16(anon_sym_DQUOTE),
	10559: uint16(364),
	10560: uint16(1),
	10561: uint16(sym__command_substitution_dollar),
	10562: uint16(293),
	10563: uint16(2),
	10564: uint16(sym_variable_expansion),
	10565: uint16(aux_sym_double_quote_string_repeat1),
	10566: uint16(2),
	10567: uint16(771),
	10568: uint16(1),
	10569: uint16(sym_comment),
	10570: uint16(968),
	10571: uint16(7),
	10572: uint16(anon_sym_DOLLAR),
	10573: uint16(anon_sym_LPAREN),
	10574: uint16(sym_integer),
	10575: uint16(anon_sym_DOT_DOT),
	10576: uint16(anon_sym_RBRACK),
	10577: uint16(anon_sym_DQUOTE),
	10578: uint16(anon_sym_SQUOTE),
	10579: uint16(7),
	10580: uint16(3),
	10581: uint16(1),
	10582: uint16(sym_comment),
	10583: uint16(934),
	10584: uint16(1),
	10585: uint16(anon_sym_DOLLAR),
	10586: uint16(970),
	10587: uint16(1),
	10588: uint16(anon_sym_DQUOTE),
	10589: uint16(972),
	10590: uint16(1),
	10591: uint16(aux_sym_double_quote_string_token1),
	10592: uint16(974),
	10593: uint16(1),
	10594: uint16(sym_escape_sequence),
	10595: uint16(364),
	10596: uint16(1),
	10597: uint16(sym__command_substitution_dollar),
	10598: uint16(290),
	10599: uint16(2),
	10600: uint16(sym_variable_expansion),
	10601: uint16(aux_sym_double_quote_string_repeat1),
	10602: uint16(7),
	10603: uint16(3),
	10604: uint16(1),
	10605: uint16(sym_comment),
	10606: uint16(976),
	10607: uint16(1),
	10608: uint16(anon_sym_DOLLAR),
	10609: uint16(979),
	10610: uint16(1),
	10611: uint16(anon_sym_DQUOTE),
	10612: uint16(981),
	10613: uint16(1),
	10614: uint16(aux_sym_double_quote_string_token1),
	10615: uint16(984),
	10616: uint16(1),
	10617: uint16(sym_escape_sequence),
	10618: uint16(364),
	10619: uint16(1),
	10620: uint16(sym__command_substitution_dollar),
	10621: uint16(293),
	10622: uint16(2),
	10623: uint16(sym_variable_expansion),
	10624: uint16(aux_sym_double_quote_string_repeat1),
	10625: uint16(7),
	10626: uint16(3),
	10627: uint16(1),
	10628: uint16(sym_comment),
	10629: uint16(934),
	10630: uint16(1),
	10631: uint16(anon_sym_DOLLAR),
	10632: uint16(950),
	10633: uint16(1),
	10634: uint16(aux_sym_double_quote_string_token1),
	10635: uint16(952),
	10636: uint16(1),
	10637: uint16(sym_escape_sequence),
	10638: uint16(987),
	10639: uint16(1),
	10640: uint16(anon_sym_DQUOTE),
	10641: uint16(364),
	10642: uint16(1),
	10643: uint16(sym__command_substitution_dollar),
	10644: uint16(293),
	10645: uint16(2),
	10646: uint16(sym_variable_expansion),
	10647: uint16(aux_sym_double_quote_string_repeat1),
	10648: uint16(7),
	10649: uint16(3),
	10650: uint16(1),
	10651: uint16(sym_comment),
	10652: uint16(934),
	10653: uint16(1),
	10654: uint16(anon_sym_DOLLAR),
	10655: uint16(989),
	10656: uint16(1),
	10657: uint16(anon_sym_DQUOTE),
	10658: uint16(991),
	10659: uint16(1),
	10660: uint16(aux_sym_double_quote_string_token1),
	10661: uint16(993),
	10662: uint16(1),
	10663: uint16(sym_escape_sequence),
	10664: uint16(364),
	10665: uint16(1),
	10666: uint16(sym__command_substitution_dollar),
	10667: uint16(294),
	10668: uint16(2),
	10669: uint16(sym_variable_expansion),
	10670: uint16(aux_sym_double_quote_string_repeat1),
	10671: uint16(3),
	10672: uint16(3),
	10673: uint16(1),
	10674: uint16(sym_comment),
	10675: uint16(503),
	10676: uint16(2),
	10677: uint16(sym__concat),
	10678: uint16(sym__concat_list),
	10679: uint16(501),
	10680: uint16(5),
	10681: uint16(anon_sym_SEMI),
	10682: uint16(anon_sym_AMP),
	10683: uint16(anon_sym_LF),
	10684: uint16(anon_sym_CR),
	10685: uint16(anon_sym_CR_LF),
	10686: uint16(2),
	10687: uint16(771),
	10688: uint16(1),
	10689: uint16(sym_comment),
	10690: uint16(654),
	10691: uint16(7),
	10692: uint16(anon_sym_DOLLAR),
	10693: uint16(anon_sym_LPAREN),
	10694: uint16(sym_integer),
	10695: uint16(anon_sym_DOT_DOT),
	10696: uint16(anon_sym_RBRACK),
	10697: uint16(anon_sym_DQUOTE),
	10698: uint16(anon_sym_SQUOTE),
	10699: uint16(7),
	10700: uint16(3),
	10701: uint16(1),
	10702: uint16(sym_comment),
	10703: uint16(934),
	10704: uint16(1),
	10705: uint16(anon_sym_DOLLAR),
	10706: uint16(950),
	10707: uint16(1),
	10708: uint16(aux_sym_double_quote_string_token1),
	10709: uint16(952),
	10710: uint16(1),
	10711: uint16(sym_escape_sequence),
	10712: uint16(995),
	10713: uint16(1),
	10714: uint16(anon_sym_DQUOTE),
	10715: uint16(364),
	10716: uint16(1),
	10717: uint16(sym__command_substitution_dollar),
	10718: uint16(293),
	10719: uint16(2),
	10720: uint16(sym_variable_expansion),
	10721: uint16(aux_sym_double_quote_string_repeat1),
	10722: uint16(7),
	10723: uint16(3),
	10724: uint16(1),
	10725: uint16(sym_comment),
	10726: uint16(934),
	10727: uint16(1),
	10728: uint16(anon_sym_DOLLAR),
	10729: uint16(950),
	10730: uint16(1),
	10731: uint16(aux_sym_double_quote_string_token1),
	10732: uint16(952),
	10733: uint16(1),
	10734: uint16(sym_escape_sequence),
	10735: uint16(997),
	10736: uint16(1),
	10737: uint16(anon_sym_DQUOTE),
	10738: uint16(364),
	10739: uint16(1),
	10740: uint16(sym__command_substitution_dollar),
	10741: uint16(293),
	10742: uint16(2),
	10743: uint16(sym_variable_expansion),
	10744: uint16(aux_sym_double_quote_string_repeat1),
	10745: uint16(4),
	10746: uint16(3),
	10747: uint16(1),
	10748: uint16(sym_comment),
	10749: uint16(999),
	10750: uint16(1),
	10751: uint16(sym__concat),
	10752: uint16(309),
	10753: uint16(1),
	10754: uint16(aux_sym_concatenation_repeat1),
	10755: uint16(508),
	10756: uint16(5),
	10757: uint16(anon_sym_SEMI),
	10758: uint16(anon_sym_AMP),
	10759: uint16(anon_sym_LF),
	10760: uint16(anon_sym_CR),
	10761: uint16(anon_sym_CR_LF),
	10762: uint16(2),
	10763: uint16(771),
	10764: uint16(1),
	10765: uint16(sym_comment),
	10766: uint16(680),
	10767: uint16(7),
	10768: uint16(anon_sym_DOLLAR),
	10769: uint16(anon_sym_LPAREN),
	10770: uint16(sym_integer),
	10771: uint16(anon_sym_DOT_DOT),
	10772: uint16(anon_sym_RBRACK),
	10773: uint16(anon_sym_DQUOTE),
	10774: uint16(anon_sym_SQUOTE),
	10775: uint16(2),
	10776: uint16(771),
	10777: uint16(1),
	10778: uint16(sym_comment),
	10779: uint16(672),
	10780: uint16(7),
	10781: uint16(anon_sym_DOLLAR),
	10782: uint16(anon_sym_LPAREN),
	10783: uint16(sym_integer),
	10784: uint16(anon_sym_DOT_DOT),
	10785: uint16(anon_sym_RBRACK),
	10786: uint16(anon_sym_DQUOTE),
	10787: uint16(anon_sym_SQUOTE),
	10788: uint16(4),
	10789: uint16(3),
	10790: uint16(1),
	10791: uint16(sym_comment),
	10792: uint16(1001),
	10793: uint16(1),
	10794: uint16(sym__concat),
	10795: uint16(303),
	10796: uint16(1),
	10797: uint16(aux_sym_concatenation_repeat1),
	10798: uint16(544),
	10799: uint16(5),
	10800: uint16(anon_sym_SEMI),
	10801: uint16(anon_sym_AMP),
	10802: uint16(anon_sym_LF),
	10803: uint16(anon_sym_CR),
	10804: uint16(anon_sym_CR_LF),
	10805: uint16(7),
	10806: uint16(3),
	10807: uint16(1),
	10808: uint16(sym_comment),
	10809: uint16(934),
	10810: uint16(1),
	10811: uint16(anon_sym_DOLLAR),
	10812: uint16(950),
	10813: uint16(1),
	10814: uint16(aux_sym_double_quote_string_token1),
	10815: uint16(952),
	10816: uint16(1),
	10817: uint16(sym_escape_sequence),
	10818: uint16(1004),
	10819: uint16(1),
	10820: uint16(anon_sym_DQUOTE),
	10821: uint16(364),
	10822: uint16(1),
	10823: uint16(sym__command_substitution_dollar),
	10824: uint16(293),
	10825: uint16(2),
	10826: uint16(sym_variable_expansion),
	10827: uint16(aux_sym_double_quote_string_repeat1),
	10828: uint16(2),
	10829: uint16(771),
	10830: uint16(1),
	10831: uint16(sym_comment),
	10832: uint16(646),
	10833: uint16(7),
	10834: uint16(anon_sym_DOLLAR),
	10835: uint16(anon_sym_LPAREN),
	10836: uint16(sym_integer),
	10837: uint16(anon_sym_DOT_DOT),
	10838: uint16(anon_sym_RBRACK),
	10839: uint16(anon_sym_DQUOTE),
	10840: uint16(anon_sym_SQUOTE),
	10841: uint16(2),
	10842: uint16(771),
	10843: uint16(1),
	10844: uint16(sym_comment),
	10845: uint16(650),
	10846: uint16(7),
	10847: uint16(anon_sym_DOLLAR),
	10848: uint16(anon_sym_LPAREN),
	10849: uint16(sym_integer),
	10850: uint16(anon_sym_DOT_DOT),
	10851: uint16(anon_sym_RBRACK),
	10852: uint16(anon_sym_DQUOTE),
	10853: uint16(anon_sym_SQUOTE),
	10854: uint16(4),
	10855: uint16(3),
	10856: uint16(1),
	10857: uint16(sym_comment),
	10858: uint16(999),
	10859: uint16(1),
	10860: uint16(sym__concat),
	10861: uint16(309),
	10862: uint16(1),
	10863: uint16(aux_sym_concatenation_repeat1),
	10864: uint16(538),
	10865: uint16(5),
	10866: uint16(anon_sym_SEMI),
	10867: uint16(anon_sym_AMP),
	10868: uint16(anon_sym_LF),
	10869: uint16(anon_sym_CR),
	10870: uint16(anon_sym_CR_LF),
	10871: uint16(7),
	10872: uint16(3),
	10873: uint16(1),
	10874: uint16(sym_comment),
	10875: uint16(934),
	10876: uint16(1),
	10877: uint16(anon_sym_DOLLAR),
	10878: uint16(1006),
	10879: uint16(1),
	10880: uint16(anon_sym_DQUOTE),
	10881: uint16(1008),
	10882: uint16(1),
	10883: uint16(aux_sym_double_quote_string_token1),
	10884: uint16(1010),
	10885: uint16(1),
	10886: uint16(sym_escape_sequence),
	10887: uint16(364),
	10888: uint16(1),
	10889: uint16(sym__command_substitution_dollar),
	10890: uint16(304),
	10891: uint16(2),
	10892: uint16(sym_variable_expansion),
	10893: uint16(aux_sym_double_quote_string_repeat1),
	10894: uint16(4),
	10895: uint16(3),
	10896: uint16(1),
	10897: uint16(sym_comment),
	10898: uint16(999),
	10899: uint16(1),
	10900: uint16(sym__concat),
	10901: uint16(303),
	10902: uint16(1),
	10903: uint16(aux_sym_concatenation_repeat1),
	10904: uint16(601),
	10905: uint16(5),
	10906: uint16(anon_sym_SEMI),
	10907: uint16(anon_sym_AMP),
	10908: uint16(anon_sym_LF),
	10909: uint16(anon_sym_CR),
	10910: uint16(anon_sym_CR_LF),
	10911: uint16(2),
	10912: uint16(771),
	10913: uint16(1),
	10914: uint16(sym_comment),
	10915: uint16(668),
	10916: uint16(7),
	10917: uint16(anon_sym_DOLLAR),
	10918: uint16(anon_sym_LPAREN),
	10919: uint16(sym_integer),
	10920: uint16(anon_sym_DOT_DOT),
	10921: uint16(anon_sym_RBRACK),
	10922: uint16(anon_sym_DQUOTE),
	10923: uint16(anon_sym_SQUOTE),
	10924: uint16(3),
	10925: uint16(3),
	10926: uint16(1),
	10927: uint16(sym_comment),
	10928: uint16(574),
	10929: uint16(2),
	10930: uint16(sym__concat),
	10931: uint16(sym__concat_list),
	10932: uint16(572),
	10933: uint16(5),
	10934: uint16(anon_sym_SEMI),
	10935: uint16(anon_sym_AMP),
	10936: uint16(anon_sym_LF),
	10937: uint16(anon_sym_CR),
	10938: uint16(anon_sym_CR_LF),
	10939: uint16(2),
	10940: uint16(771),
	10941: uint16(1),
	10942: uint16(sym_comment),
	10943: uint16(676),
	10944: uint16(7),
	10945: uint16(anon_sym_DOLLAR),
	10946: uint16(anon_sym_LPAREN),
	10947: uint16(sym_integer),
	10948: uint16(anon_sym_DOT_DOT),
	10949: uint16(anon_sym_RBRACK),
	10950: uint16(anon_sym_DQUOTE),
	10951: uint16(anon_sym_SQUOTE),
	10952: uint16(2),
	10953: uint16(771),
	10954: uint16(1),
	10955: uint16(sym_comment),
	10956: uint16(638),
	10957: uint16(7),
	10958: uint16(anon_sym_DOLLAR),
	10959: uint16(anon_sym_LPAREN),
	10960: uint16(sym_integer),
	10961: uint16(anon_sym_DOT_DOT),
	10962: uint16(anon_sym_RBRACK),
	10963: uint16(anon_sym_DQUOTE),
	10964: uint16(anon_sym_SQUOTE),
	10965: uint16(2),
	10966: uint16(771),
	10967: uint16(1),
	10968: uint16(sym_comment),
	10969: uint16(664),
	10970: uint16(7),
	10971: uint16(anon_sym_DOLLAR),
	10972: uint16(anon_sym_LPAREN),
	10973: uint16(sym_integer),
	10974: uint16(anon_sym_DOT_DOT),
	10975: uint16(anon_sym_RBRACK),
	10976: uint16(anon_sym_DQUOTE),
	10977: uint16(anon_sym_SQUOTE),
	10978: uint16(5),
	10979: uint16(3),
	10980: uint16(1),
	10981: uint16(sym_comment),
	10982: uint16(493),
	10983: uint16(1),
	10984: uint16(aux_sym_double_quote_string_token1),
	10985: uint16(1012),
	10986: uint16(1),
	10987: uint16(sym__concat_list),
	10988: uint16(320),
	10989: uint16(1),
	10990: uint16(aux_sym_variable_expansion_repeat1),
	10991: uint16(491),
	10992: uint16(3),
	10993: uint16(anon_sym_DOLLAR),
	10994: uint16(anon_sym_DQUOTE),
	10995: uint16(sym_escape_sequence),
	10996: uint16(5),
	10997: uint16(3),
	10998: uint16(1),
	10999: uint16(sym_comment),
	11000: uint16(499),
	11001: uint16(1),
	11002: uint16(aux_sym_double_quote_string_token1),
	11003: uint16(1012),
	11004: uint16(1),
	11005: uint16(sym__concat_list),
	11006: uint16(315),
	11007: uint16(1),
	11008: uint16(aux_sym_variable_expansion_repeat1),
	11009: uint16(497),
	11010: uint16(3),
	11011: uint16(anon_sym_DOLLAR),
	11012: uint16(anon_sym_DQUOTE),
	11013: uint16(sym_escape_sequence),
	11014: uint16(3),
	11015: uint16(3),
	11016: uint16(1),
	11017: uint16(sym_comment),
	11018: uint16(676),
	11019: uint16(1),
	11020: uint16(sym__concat),
	11021: uint16(674),
	11022: uint16(5),
	11023: uint16(anon_sym_SEMI),
	11024: uint16(anon_sym_AMP),
	11025: uint16(anon_sym_LF),
	11026: uint16(anon_sym_CR),
	11027: uint16(anon_sym_CR_LF),
	11028: uint16(3),
	11029: uint16(3),
	11030: uint16(1),
	11031: uint16(sym_comment),
	11032: uint16(664),
	11033: uint16(1),
	11034: uint16(sym__concat),
	11035: uint16(662),
	11036: uint16(5),
	11037: uint16(anon_sym_SEMI),
	11038: uint16(anon_sym_AMP),
	11039: uint16(anon_sym_LF),
	11040: uint16(anon_sym_CR),
	11041: uint16(anon_sym_CR_LF),
	11042: uint16(3),
	11043: uint16(3),
	11044: uint16(1),
	11045: uint16(sym_comment),
	11046: uint16(656),
	11047: uint16(1),
	11048: uint16(sym__concat),
	11049: uint16(544),
	11050: uint16(5),
	11051: uint16(anon_sym_SEMI),
	11052: uint16(anon_sym_AMP),
	11053: uint16(anon_sym_LF),
	11054: uint16(anon_sym_CR),
	11055: uint16(anon_sym_CR_LF),
	11056: uint16(5),
	11057: uint16(3),
	11058: uint16(1),
	11059: uint16(sym_comment),
	11060: uint16(503),
	11061: uint16(1),
	11062: uint16(aux_sym_double_quote_string_token1),
	11063: uint16(1014),
	11064: uint16(1),
	11065: uint16(sym__concat_list),
	11066: uint16(320),
	11067: uint16(1),
	11068: uint16(aux_sym_variable_expansion_repeat1),
	11069: uint16(501),
	11070: uint16(3),
	11071: uint16(anon_sym_DOLLAR),
	11072: uint16(anon_sym_DQUOTE),
	11073: uint16(sym_escape_sequence),
	11074: uint16(3),
	11075: uint16(3),
	11076: uint16(1),
	11077: uint16(sym_comment),
	11078: uint16(638),
	11079: uint16(1),
	11080: uint16(sym__concat),
	11081: uint16(636),
	11082: uint16(5),
	11083: uint16(anon_sym_SEMI),
	11084: uint16(anon_sym_AMP),
	11085: uint16(anon_sym_LF),
	11086: uint16(anon_sym_CR),
	11087: uint16(anon_sym_CR_LF),
	11088: uint16(3),
	11089: uint16(3),
	11090: uint16(1),
	11091: uint16(sym_comment),
	11092: uint16(654),
	11093: uint16(1),
	11094: uint16(sym__concat),
	11095: uint16(652),
	11096: uint16(5),
	11097: uint16(anon_sym_SEMI),
	11098: uint16(anon_sym_AMP),
	11099: uint16(anon_sym_LF),
	11100: uint16(anon_sym_CR),
	11101: uint16(anon_sym_CR_LF),
	11102: uint16(3),
	11103: uint16(3),
	11104: uint16(1),
	11105: uint16(sym_comment),
	11106: uint16(642),
	11107: uint16(1),
	11108: uint16(sym__concat),
	11109: uint16(640),
	11110: uint16(5),
	11111: uint16(anon_sym_SEMI),
	11112: uint16(anon_sym_AMP),
	11113: uint16(anon_sym_LF),
	11114: uint16(anon_sym_CR),
	11115: uint16(anon_sym_CR_LF),
	11116: uint16(3),
	11117: uint16(3),
	11118: uint16(1),
	11119: uint16(sym_comment),
	11120: uint16(660),
	11121: uint16(1),
	11122: uint16(sym__concat),
	11123: uint16(658),
	11124: uint16(5),
	11125: uint16(anon_sym_SEMI),
	11126: uint16(anon_sym_AMP),
	11127: uint16(anon_sym_LF),
	11128: uint16(anon_sym_CR),
	11129: uint16(anon_sym_CR_LF),
	11130: uint16(3),
	11131: uint16(3),
	11132: uint16(1),
	11133: uint16(sym_comment),
	11134: uint16(668),
	11135: uint16(1),
	11136: uint16(sym__concat),
	11137: uint16(666),
	11138: uint16(5),
	11139: uint16(anon_sym_SEMI),
	11140: uint16(anon_sym_AMP),
	11141: uint16(anon_sym_LF),
	11142: uint16(anon_sym_CR),
	11143: uint16(anon_sym_CR_LF),
	11144: uint16(3),
	11145: uint16(3),
	11146: uint16(1),
	11147: uint16(sym_comment),
	11148: uint16(1019),
	11149: uint16(1),
	11150: uint16(anon_sym_if),
	11151: uint16(1017),
	11152: uint16(5),
	11153: uint16(anon_sym_SEMI),
	11154: uint16(anon_sym_AMP),
	11155: uint16(anon_sym_LF),
	11156: uint16(anon_sym_CR),
	11157: uint16(anon_sym_CR_LF),
	11158: uint16(3),
	11159: uint16(3),
	11160: uint16(1),
	11161: uint16(sym_comment),
	11162: uint16(650),
	11163: uint16(1),
	11164: uint16(sym__concat),
	11165: uint16(648),
	11166: uint16(5),
	11167: uint16(anon_sym_SEMI),
	11168: uint16(anon_sym_AMP),
	11169: uint16(anon_sym_LF),
	11170: uint16(anon_sym_CR),
	11171: uint16(anon_sym_CR_LF),
	11172: uint16(3),
	11173: uint16(3),
	11174: uint16(1),
	11175: uint16(sym_comment),
	11176: uint16(672),
	11177: uint16(1),
	11178: uint16(sym__concat),
	11179: uint16(670),
	11180: uint16(5),
	11181: uint16(anon_sym_SEMI),
	11182: uint16(anon_sym_AMP),
	11183: uint16(anon_sym_LF),
	11184: uint16(anon_sym_CR),
	11185: uint16(anon_sym_CR_LF),
	11186: uint16(3),
	11187: uint16(3),
	11188: uint16(1),
	11189: uint16(sym_comment),
	11190: uint16(680),
	11191: uint16(1),
	11192: uint16(sym__concat),
	11193: uint16(678),
	11194: uint16(5),
	11195: uint16(anon_sym_SEMI),
	11196: uint16(anon_sym_AMP),
	11197: uint16(anon_sym_LF),
	11198: uint16(anon_sym_CR),
	11199: uint16(anon_sym_CR_LF),
	11200: uint16(4),
	11201: uint16(3),
	11202: uint16(1),
	11203: uint16(sym_comment),
	11204: uint16(331),
	11205: uint16(1),
	11206: uint16(aux_sym_variable_expansion_repeat1),
	11207: uint16(499),
	11208: uint16(2),
	11209: uint16(sym__concat_list),
	11210: uint16(aux_sym_double_quote_string_token1),
	11211: uint16(497),
	11212: uint16(3),
	11213: uint16(anon_sym_DOLLAR),
	11214: uint16(anon_sym_DQUOTE),
	11215: uint16(sym_escape_sequence),
	11216: uint16(4),
	11217: uint16(3),
	11218: uint16(1),
	11219: uint16(sym_comment),
	11220: uint16(320),
	11221: uint16(1),
	11222: uint16(aux_sym_variable_expansion_repeat1),
	11223: uint16(493),
	11224: uint16(2),
	11225: uint16(sym__concat_list),
	11226: uint16(aux_sym_double_quote_string_token1),
	11227: uint16(491),
	11228: uint16(3),
	11229: uint16(anon_sym_DOLLAR),
	11230: uint16(anon_sym_DQUOTE),
	11231: uint16(sym_escape_sequence),
	11232: uint16(3),
	11233: uint16(3),
	11234: uint16(1),
	11235: uint16(sym_comment),
	11236: uint16(646),
	11237: uint16(1),
	11238: uint16(sym__concat),
	11239: uint16(644),
	11240: uint16(5),
	11241: uint16(anon_sym_SEMI),
	11242: uint16(anon_sym_AMP),
	11243: uint16(anon_sym_LF),
	11244: uint16(anon_sym_CR),
	11245: uint16(anon_sym_CR_LF),
	11246: uint16(3),
	11247: uint16(3),
	11248: uint16(1),
	11249: uint16(sym_comment),
	11250: uint16(688),
	11251: uint16(1),
	11252: uint16(sym__concat),
	11253: uint16(686),
	11254: uint16(5),
	11255: uint16(anon_sym_SEMI),
	11256: uint16(anon_sym_AMP),
	11257: uint16(anon_sym_LF),
	11258: uint16(anon_sym_CR),
	11259: uint16(anon_sym_CR_LF),
	11260: uint16(3),
	11261: uint16(3),
	11262: uint16(1),
	11263: uint16(sym_comment),
	11264: uint16(684),
	11265: uint16(1),
	11266: uint16(sym__concat),
	11267: uint16(682),
	11268: uint16(5),
	11269: uint16(anon_sym_SEMI),
	11270: uint16(anon_sym_AMP),
	11271: uint16(anon_sym_LF),
	11272: uint16(anon_sym_CR),
	11273: uint16(anon_sym_CR_LF),
	11274: uint16(6),
	11275: uint16(771),
	11276: uint16(1),
	11277: uint16(sym_comment),
	11278: uint16(1021),
	11279: uint16(1),
	11280: uint16(anon_sym_DOLLAR),
	11281: uint16(1023),
	11282: uint16(1),
	11283: uint16(anon_sym_LPAREN),
	11284: uint16(1025),
	11285: uint16(1),
	11286: uint16(sym_variable_name),
	11287: uint16(340),
	11288: uint16(1),
	11289: uint16(sym_variable_expansion),
	11290: uint16(385),
	11291: uint16(1),
	11292: uint16(sym__command_substitution_inner),
	11293: uint16(6),
	11294: uint16(771),
	11295: uint16(1),
	11296: uint16(sym_comment),
	11297: uint16(1027),
	11298: uint16(1),
	11299: uint16(anon_sym_DOLLAR),
	11300: uint16(1029),
	11301: uint16(1),
	11302: uint16(anon_sym_LPAREN),
	11303: uint16(1031),
	11304: uint16(1),
	11305: uint16(sym_variable_name),
	11306: uint16(316),
	11307: uint16(1),
	11308: uint16(sym_variable_expansion),
	11309: uint16(368),
	11310: uint16(1),
	11311: uint16(sym__command_substitution_inner),
	11312: uint16(3),
	11313: uint16(3),
	11314: uint16(1),
	11315: uint16(sym_comment),
	11316: uint16(542),
	11317: uint16(2),
	11318: uint16(sym__concat_list),
	11319: uint16(aux_sym_double_quote_string_token1),
	11320: uint16(540),
	11321: uint16(3),
	11322: uint16(anon_sym_DOLLAR),
	11323: uint16(anon_sym_DQUOTE),
	11324: uint16(sym_escape_sequence),
	11325: uint16(3),
	11326: uint16(3),
	11327: uint16(1),
	11328: uint16(sym_comment),
	11329: uint16(574),
	11330: uint16(2),
	11331: uint16(sym__concat_list),
	11332: uint16(aux_sym_double_quote_string_token1),
	11333: uint16(572),
	11334: uint16(3),
	11335: uint16(anon_sym_DOLLAR),
	11336: uint16(anon_sym_DQUOTE),
	11337: uint16(sym_escape_sequence),
	11338: uint16(5),
	11339: uint16(771),
	11340: uint16(1),
	11341: uint16(sym_comment),
	11342: uint16(1033),
	11343: uint16(1),
	11344: uint16(anon_sym_end),
	11345: uint16(1035),
	11346: uint16(1),
	11347: uint16(anon_sym_else),
	11348: uint16(431),
	11349: uint16(1),
	11350: uint16(sym_else_clause),
	11351: uint16(363),
	11352: uint16(2),
	11353: uint16(sym_else_if_clause),
	11354: uint16(aux_sym_if_statement_repeat1),
	11355: uint16(4),
	11356: uint16(771),
	11357: uint16(1),
	11358: uint16(sym_comment),
	11359: uint16(1037),
	11360: uint16(1),
	11361: uint16(sym__concat_list),
	11362: uint16(348),
	11363: uint16(1),
	11364: uint16(aux_sym_variable_expansion_repeat1),
	11365: uint16(499),
	11366: uint16(3),
	11367: uint16(sym__brace_concat),
	11368: uint16(anon_sym_RBRACE),
	11369: uint16(anon_sym_COMMA),
	11370: uint16(6),
	11371: uint16(771),
	11372: uint16(1),
	11373: uint16(sym_comment),
	11374: uint16(1039),
	11375: uint16(1),
	11376: uint16(anon_sym_DOLLAR),
	11377: uint16(1041),
	11378: uint16(1),
	11379: uint16(anon_sym_LPAREN),
	11380: uint16(1043),
	11381: uint16(1),
	11382: uint16(sym_variable_name),
	11383: uint16(280),
	11384: uint16(1),
	11385: uint16(sym_variable_expansion),
	11386: uint16(321),
	11387: uint16(1),
	11388: uint16(sym__command_substitution_inner),
	11389: uint16(6),
	11390: uint16(771),
	11391: uint16(1),
	11392: uint16(sym_comment),
	11393: uint16(1045),
	11394: uint16(1),
	11395: uint16(anon_sym_DOLLAR),
	11396: uint16(1047),
	11397: uint16(1),
	11398: uint16(anon_sym_LPAREN),
	11399: uint16(1049),
	11400: uint16(1),
	11401: uint16(sym_variable_name),
	11402: uint16(79),
	11403: uint16(1),
	11404: uint16(sym_variable_expansion),
	11405: uint16(99),
	11406: uint16(1),
	11407: uint16(sym__command_substitution_inner),
	11408: uint16(3),
	11409: uint16(771),
	11410: uint16(1),
	11411: uint16(sym_comment),
	11412: uint16(344),
	11413: uint16(1),
	11414: uint16(aux_sym_variable_expansion_repeat1),
	11415: uint16(499),
	11416: uint16(4),
	11417: uint16(sym__brace_concat),
	11418: uint16(sym__concat_list),
	11419: uint16(anon_sym_RBRACE),
	11420: uint16(anon_sym_COMMA),
	11421: uint16(3),
	11422: uint16(771),
	11423: uint16(1),
	11424: uint16(sym_comment),
	11425: uint16(347),
	11426: uint16(1),
	11427: uint16(aux_sym_variable_expansion_repeat1),
	11428: uint16(493),
	11429: uint16(4),
	11430: uint16(sym__brace_concat),
	11431: uint16(sym__concat_list),
	11432: uint16(anon_sym_RBRACE),
	11433: uint16(anon_sym_COMMA),
	11434: uint16(6),
	11435: uint16(767),
	11436: uint16(1),
	11437: uint16(anon_sym_LPAREN),
	11438: uint16(771),
	11439: uint16(1),
	11440: uint16(sym_comment),
	11441: uint16(1051),
	11442: uint16(1),
	11443: uint16(anon_sym_DOLLAR),
	11444: uint16(1053),
	11445: uint16(1),
	11446: uint16(sym_variable_name),
	11447: uint16(272),
	11448: uint16(1),
	11449: uint16(sym_variable_expansion),
	11450: uint16(313),
	11451: uint16(1),
	11452: uint16(sym__command_substitution_inner),
	11453: uint16(6),
	11454: uint16(771),
	11455: uint16(1),
	11456: uint16(sym_comment),
	11457: uint16(1055),
	11458: uint16(1),
	11459: uint16(anon_sym_DOLLAR),
	11460: uint16(1057),
	11461: uint16(1),
	11462: uint16(anon_sym_LPAREN),
	11463: uint16(1059),
	11464: uint16(1),
	11465: uint16(sym_variable_name),
	11466: uint16(133),
	11467: uint16(1),
	11468: uint16(sym_variable_expansion),
	11469: uint16(145),
	11470: uint16(1),
	11471: uint16(sym__command_substitution_inner),
	11472: uint16(4),
	11473: uint16(771),
	11474: uint16(1),
	11475: uint16(sym_comment),
	11476: uint16(1061),
	11477: uint16(1),
	11478: uint16(sym__concat_list),
	11479: uint16(347),
	11480: uint16(1),
	11481: uint16(aux_sym_variable_expansion_repeat1),
	11482: uint16(503),
	11483: uint16(3),
	11484: uint16(sym__brace_concat),
	11485: uint16(anon_sym_RBRACE),
	11486: uint16(anon_sym_COMMA),
	11487: uint16(4),
	11488: uint16(771),
	11489: uint16(1),
	11490: uint16(sym_comment),
	11491: uint16(1037),
	11492: uint16(1),
	11493: uint16(sym__concat_list),
	11494: uint16(347),
	11495: uint16(1),
	11496: uint16(aux_sym_variable_expansion_repeat1),
	11497: uint16(493),
	11498: uint16(3),
	11499: uint16(sym__brace_concat),
	11500: uint16(anon_sym_RBRACE),
	11501: uint16(anon_sym_COMMA),
	11502: uint16(3),
	11503: uint16(3),
	11504: uint16(1),
	11505: uint16(sym_comment),
	11506: uint16(503),
	11507: uint16(2),
	11508: uint16(sym__concat_list),
	11509: uint16(aux_sym_double_quote_string_token1),
	11510: uint16(501),
	11511: uint16(3),
	11512: uint16(anon_sym_DOLLAR),
	11513: uint16(anon_sym_DQUOTE),
	11514: uint16(sym_escape_sequence),
	11515: uint16(2),
	11516: uint16(3),
	11517: uint16(1),
	11518: uint16(sym_comment),
	11519: uint16(1064),
	11520: uint16(5),
	11521: uint16(anon_sym_SEMI),
	11522: uint16(anon_sym_AMP),
	11523: uint16(anon_sym_LF),
	11524: uint16(anon_sym_CR),
	11525: uint16(anon_sym_CR_LF),
	11526: uint16(5),
	11527: uint16(771),
	11528: uint16(1),
	11529: uint16(sym_comment),
	11530: uint16(1035),
	11531: uint16(1),
	11532: uint16(anon_sym_else),
	11533: uint16(1066),
	11534: uint16(1),
	11535: uint16(anon_sym_end),
	11536: uint16(425),
	11537: uint16(1),
	11538: uint16(sym_else_clause),
	11539: uint16(363),
	11540: uint16(2),
	11541: uint16(sym_else_if_clause),
	11542: uint16(aux_sym_if_statement_repeat1),
	11543: uint16(6),
	11544: uint16(771),
	11545: uint16(1),
	11546: uint16(sym_comment),
	11547: uint16(1068),
	11548: uint16(1),
	11549: uint16(anon_sym_DOLLAR),
	11550: uint16(1070),
	11551: uint16(1),
	11552: uint16(anon_sym_LPAREN),
	11553: uint16(1072),
	11554: uint16(1),
	11555: uint16(sym_variable_name),
	11556: uint16(162),
	11557: uint16(1),
	11558: uint16(sym_variable_expansion),
	11559: uint16(200),
	11560: uint16(1),
	11561: uint16(sym__command_substitution_inner),
	11562: uint16(5),
	11563: uint16(3),
	11564: uint16(1),
	11565: uint16(sym_comment),
	11566: uint16(1074),
	11567: uint16(1),
	11568: uint16(anon_sym_SQUOTE),
	11569: uint16(1076),
	11570: uint16(1),
	11571: uint16(aux_sym_single_quote_string_token1),
	11572: uint16(1078),
	11573: uint16(1),
	11574: uint16(sym_escape_sequence),
	11575: uint16(370),
	11576: uint16(1),
	11577: uint16(aux_sym_single_quote_string_repeat1),
	11578: uint16(5),
	11579: uint16(3),
	11580: uint16(1),
	11581: uint16(sym_comment),
	11582: uint16(1080),
	11583: uint16(1),
	11584: uint16(anon_sym_SQUOTE),
	11585: uint16(1082),
	11586: uint16(1),
	11587: uint16(aux_sym_single_quote_string_token1),
	11588: uint16(1084),
	11589: uint16(1),
	11590: uint16(sym_escape_sequence),
	11591: uint16(356),
	11592: uint16(1),
	11593: uint16(aux_sym_single_quote_string_repeat1),
	11594: uint16(5),
	11595: uint16(3),
	11596: uint16(1),
	11597: uint16(sym_comment),
	11598: uint16(1086),
	11599: uint16(1),
	11600: uint16(anon_sym_SQUOTE),
	11601: uint16(1088),
	11602: uint16(1),
	11603: uint16(aux_sym_single_quote_string_token1),
	11604: uint16(1090),
	11605: uint16(1),
	11606: uint16(sym_escape_sequence),
	11607: uint16(371),
	11608: uint16(1),
	11609: uint16(aux_sym_single_quote_string_repeat1),
	11610: uint16(5),
	11611: uint16(3),
	11612: uint16(1),
	11613: uint16(sym_comment),
	11614: uint16(1076),
	11615: uint16(1),
	11616: uint16(aux_sym_single_quote_string_token1),
	11617: uint16(1078),
	11618: uint16(1),
	11619: uint16(sym_escape_sequence),
	11620: uint16(1092),
	11621: uint16(1),
	11622: uint16(anon_sym_SQUOTE),
	11623: uint16(370),
	11624: uint16(1),
	11625: uint16(aux_sym_single_quote_string_repeat1),
	11626: uint16(4),
	11627: uint16(771),
	11628: uint16(1),
	11629: uint16(sym_comment),
	11630: uint16(1096),
	11631: uint16(1),
	11632: uint16(sym__brace_concat),
	11633: uint16(372),
	11634: uint16(1),
	11635: uint16(aux_sym_brace_concatenation_repeat1),
	11636: uint16(1094),
	11637: uint16(2),
	11638: uint16(anon_sym_RBRACE),
	11639: uint16(anon_sym_COMMA),
	11640: uint16(4),
	11641: uint16(771),
	11642: uint16(1),
	11643: uint16(sym_comment),
	11644: uint16(1098),
	11645: uint16(1),
	11646: uint16(anon_sym_end),
	11647: uint16(1100),
	11648: uint16(1),
	11649: uint16(anon_sym_case),
	11650: uint16(377),
	11651: uint16(2),
	11652: uint16(sym_case_clause),
	11653: uint16(aux_sym_switch_statement_repeat1),
	11654: uint16(4),
	11655: uint16(771),
	11656: uint16(1),
	11657: uint16(sym_comment),
	11658: uint16(1104),
	11659: uint16(1),
	11660: uint16(sym__brace_concat),
	11661: uint16(359),
	11662: uint16(1),
	11663: uint16(aux_sym_brace_concatenation_repeat1),
	11664: uint16(1102),
	11665: uint16(2),
	11666: uint16(anon_sym_RBRACE),
	11667: uint16(anon_sym_COMMA),
	11668: uint16(5),
	11669: uint16(3),
	11670: uint16(1),
	11671: uint16(sym_comment),
	11672: uint16(1107),
	11673: uint16(1),
	11674: uint16(anon_sym_SQUOTE),
	11675: uint16(1109),
	11676: uint16(1),
	11677: uint16(aux_sym_single_quote_string_token1),
	11678: uint16(1111),
	11679: uint16(1),
	11680: uint16(sym_escape_sequence),
	11681: uint16(369),
	11682: uint16(1),
	11683: uint16(aux_sym_single_quote_string_repeat1),
	11684: uint16(2),
	11685: uint16(771),
	11686: uint16(1),
	11687: uint16(sym_comment),
	11688: uint16(503),
	11689: uint16(4),
	11690: uint16(sym__brace_concat),
	11691: uint16(sym__concat_list),
	11692: uint16(anon_sym_RBRACE),
	11693: uint16(anon_sym_COMMA),
	11694: uint16(3),
	11695: uint16(3),
	11696: uint16(1),
	11697: uint16(sym_comment),
	11698: uint16(676),
	11699: uint16(1),
	11700: uint16(aux_sym_double_quote_string_token1),
	11701: uint16(674),
	11702: uint16(3),
	11703: uint16(anon_sym_DOLLAR),
	11704: uint16(anon_sym_DQUOTE),
	11705: uint16(sym_escape_sequence),
	11706: uint16(4),
	11707: uint16(771),
	11708: uint16(1),
	11709: uint16(sym_comment),
	11710: uint16(1113),
	11711: uint16(1),
	11712: uint16(anon_sym_end),
	11713: uint16(1115),
	11714: uint16(1),
	11715: uint16(anon_sym_else),
	11716: uint16(363),
	11717: uint16(2),
	11718: uint16(sym_else_if_clause),
	11719: uint16(aux_sym_if_statement_repeat1),
	11720: uint16(3),
	11721: uint16(3),
	11722: uint16(1),
	11723: uint16(sym_comment),
	11724: uint16(1120),
	11725: uint16(1),
	11726: uint16(aux_sym_double_quote_string_token1),
	11727: uint16(1118),
	11728: uint16(3),
	11729: uint16(anon_sym_DOLLAR),
	11730: uint16(anon_sym_DQUOTE),
	11731: uint16(sym_escape_sequence),
	11732: uint16(3),
	11733: uint16(3),
	11734: uint16(1),
	11735: uint16(sym_comment),
	11736: uint16(664),
	11737: uint16(1),
	11738: uint16(aux_sym_double_quote_string_token1),
	11739: uint16(662),
	11740: uint16(3),
	11741: uint16(anon_sym_DOLLAR),
	11742: uint16(anon_sym_DQUOTE),
	11743: uint16(sym_escape_sequence),
	11744: uint16(5),
	11745: uint16(3),
	11746: uint16(1),
	11747: uint16(sym_comment),
	11748: uint16(1122),
	11749: uint16(1),
	11750: uint16(anon_sym_SQUOTE),
	11751: uint16(1124),
	11752: uint16(1),
	11753: uint16(aux_sym_single_quote_string_token1),
	11754: uint16(1126),
	11755: uint16(1),
	11756: uint16(sym_escape_sequence),
	11757: uint16(375),
	11758: uint16(1),
	11759: uint16(aux_sym_single_quote_string_repeat1),
	11760: uint16(5),
	11761: uint16(3),
	11762: uint16(1),
	11763: uint16(sym_comment),
	11764: uint16(1128),
	11765: uint16(1),
	11766: uint16(anon_sym_SQUOTE),
	11767: uint16(1130),
	11768: uint16(1),
	11769: uint16(aux_sym_single_quote_string_token1),
	11770: uint16(1132),
	11771: uint16(1),
	11772: uint16(sym_escape_sequence),
	11773: uint16(353),
	11774: uint16(1),
	11775: uint16(aux_sym_single_quote_string_repeat1),
	11776: uint16(3),
	11777: uint16(3),
	11778: uint16(1),
	11779: uint16(sym_comment),
	11780: uint16(638),
	11781: uint16(1),
	11782: uint16(aux_sym_double_quote_string_token1),
	11783: uint16(636),
	11784: uint16(3),
	11785: uint16(anon_sym_DOLLAR),
	11786: uint16(anon_sym_DQUOTE),
	11787: uint16(sym_escape_sequence),
	11788: uint16(5),
	11789: uint16(3),
	11790: uint16(1),
	11791: uint16(sym_comment),
	11792: uint16(1076),
	11793: uint16(1),
	11794: uint16(aux_sym_single_quote_string_token1),
	11795: uint16(1078),
	11796: uint16(1),
	11797: uint16(sym_escape_sequence),
	11798: uint16(1134),
	11799: uint16(1),
	11800: uint16(anon_sym_SQUOTE),
	11801: uint16(370),
	11802: uint16(1),
	11803: uint16(aux_sym_single_quote_string_repeat1),
	11804: uint16(5),
	11805: uint16(3),
	11806: uint16(1),
	11807: uint16(sym_comment),
	11808: uint16(1136),
	11809: uint16(1),
	11810: uint16(anon_sym_SQUOTE),
	11811: uint16(1138),
	11812: uint16(1),
	11813: uint16(aux_sym_single_quote_string_token1),
	11814: uint16(1141),
	11815: uint16(1),
	11816: uint16(sym_escape_sequence),
	11817: uint16(370),
	11818: uint16(1),
	11819: uint16(aux_sym_single_quote_string_repeat1),
	11820: uint16(5),
	11821: uint16(3),
	11822: uint16(1),
	11823: uint16(sym_comment),
	11824: uint16(1076),
	11825: uint16(1),
	11826: uint16(aux_sym_single_quote_string_token1),
	11827: uint16(1078),
	11828: uint16(1),
	11829: uint16(sym_escape_sequence),
	11830: uint16(1144),
	11831: uint16(1),
	11832: uint16(anon_sym_SQUOTE),
	11833: uint16(370),
	11834: uint16(1),
	11835: uint16(aux_sym_single_quote_string_repeat1),
	11836: uint16(4),
	11837: uint16(771),
	11838: uint16(1),
	11839: uint16(sym_comment),
	11840: uint16(1096),
	11841: uint16(1),
	11842: uint16(sym__brace_concat),
	11843: uint16(359),
	11844: uint16(1),
	11845: uint16(aux_sym_brace_concatenation_repeat1),
	11846: uint16(1146),
	11847: uint16(2),
	11848: uint16(anon_sym_RBRACE),
	11849: uint16(anon_sym_COMMA),
	11850: uint16(5),
	11851: uint16(3),
	11852: uint16(1),
	11853: uint16(sym_comment),
	11854: uint16(1076),
	11855: uint16(1),
	11856: uint16(aux_sym_single_quote_string_token1),
	11857: uint16(1078),
	11858: uint16(1),
	11859: uint16(sym_escape_sequence),
	11860: uint16(1148),
	11861: uint16(1),
	11862: uint16(anon_sym_SQUOTE),
	11863: uint16(370),
	11864: uint16(1),
	11865: uint16(aux_sym_single_quote_string_repeat1),
	11866: uint16(4),
	11867: uint16(771),
	11868: uint16(1),
	11869: uint16(sym_comment),
	11870: uint16(1100),
	11871: uint16(1),
	11872: uint16(anon_sym_case),
	11873: uint16(1150),
	11874: uint16(1),
	11875: uint16(anon_sym_end),
	11876: uint16(358),
	11877: uint16(2),
	11878: uint16(sym_case_clause),
	11879: uint16(aux_sym_switch_statement_repeat1),
	11880: uint16(5),
	11881: uint16(3),
	11882: uint16(1),
	11883: uint16(sym_comment),
	11884: uint16(1076),
	11885: uint16(1),
	11886: uint16(aux_sym_single_quote_string_token1),
	11887: uint16(1078),
	11888: uint16(1),
	11889: uint16(sym_escape_sequence),
	11890: uint16(1152),
	11891: uint16(1),
	11892: uint16(anon_sym_SQUOTE),
	11893: uint16(370),
	11894: uint16(1),
	11895: uint16(aux_sym_single_quote_string_repeat1),
	11896: uint16(2),
	11897: uint16(771),
	11898: uint16(1),
	11899: uint16(sym_comment),
	11900: uint16(542),
	11901: uint16(4),
	11902: uint16(sym__brace_concat),
	11903: uint16(sym__concat_list),
	11904: uint16(anon_sym_RBRACE),
	11905: uint16(anon_sym_COMMA),
	11906: uint16(4),
	11907: uint16(771),
	11908: uint16(1),
	11909: uint16(sym_comment),
	11910: uint16(1154),
	11911: uint16(1),
	11912: uint16(anon_sym_end),
	11913: uint16(1156),
	11914: uint16(1),
	11915: uint16(anon_sym_case),
	11916: uint16(377),
	11917: uint16(2),
	11918: uint16(sym_case_clause),
	11919: uint16(aux_sym_switch_statement_repeat1),
	11920: uint16(2),
	11921: uint16(771),
	11922: uint16(1),
	11923: uint16(sym_comment),
	11924: uint16(574),
	11925: uint16(4),
	11926: uint16(sym__brace_concat),
	11927: uint16(sym__concat_list),
	11928: uint16(anon_sym_RBRACE),
	11929: uint16(anon_sym_COMMA),
	11930: uint16(5),
	11931: uint16(3),
	11932: uint16(1),
	11933: uint16(sym_comment),
	11934: uint16(1159),
	11935: uint16(1),
	11936: uint16(anon_sym_SQUOTE),
	11937: uint16(1161),
	11938: uint16(1),
	11939: uint16(aux_sym_single_quote_string_token1),
	11940: uint16(1163),
	11941: uint16(1),
	11942: uint16(sym_escape_sequence),
	11943: uint16(373),
	11944: uint16(1),
	11945: uint16(aux_sym_single_quote_string_repeat1),
	11946: uint16(3),
	11947: uint16(3),
	11948: uint16(1),
	11949: uint16(sym_comment),
	11950: uint16(668),
	11951: uint16(1),
	11952: uint16(aux_sym_double_quote_string_token1),
	11953: uint16(666),
	11954: uint16(3),
	11955: uint16(anon_sym_DOLLAR),
	11956: uint16(anon_sym_DQUOTE),
	11957: uint16(sym_escape_sequence),
	11958: uint16(4),
	11959: uint16(771),
	11960: uint16(1),
	11961: uint16(sym_comment),
	11962: uint16(1165),
	11963: uint16(1),
	11964: uint16(anon_sym_RBRACE),
	11965: uint16(1167),
	11966: uint16(1),
	11967: uint16(anon_sym_COMMA),
	11968: uint16(392),
	11969: uint16(1),
	11970: uint16(aux_sym_brace_expansion_repeat1),
	11971: uint16(4),
	11972: uint16(771),
	11973: uint16(1),
	11974: uint16(sym_comment),
	11975: uint16(1167),
	11976: uint16(1),
	11977: uint16(anon_sym_COMMA),
	11978: uint16(1169),
	11979: uint16(1),
	11980: uint16(anon_sym_RBRACE),
	11981: uint16(392),
	11982: uint16(1),
	11983: uint16(aux_sym_brace_expansion_repeat1),
	11984: uint16(2),
	11985: uint16(771),
	11986: uint16(1),
	11987: uint16(sym_comment),
	11988: uint16(680),
	11989: uint16(3),
	11990: uint16(sym__brace_concat),
	11991: uint16(anon_sym_RBRACE),
	11992: uint16(anon_sym_COMMA),
	11993: uint16(4),
	11994: uint16(771),
	11995: uint16(1),
	11996: uint16(sym_comment),
	11997: uint16(1165),
	11998: uint16(1),
	11999: uint16(anon_sym_RBRACE),
	12000: uint16(1167),
	12001: uint16(1),
	12002: uint16(anon_sym_COMMA),
	12003: uint16(390),
	12004: uint16(1),
	12005: uint16(aux_sym_brace_expansion_repeat1),
	12006: uint16(2),
	12007: uint16(771),
	12008: uint16(1),
	12009: uint16(sym_comment),
	12010: uint16(638),
	12011: uint16(3),
	12012: uint16(sym__brace_concat),
	12013: uint16(anon_sym_RBRACE),
	12014: uint16(anon_sym_COMMA),
	12015: uint16(4),
	12016: uint16(771),
	12017: uint16(1),
	12018: uint16(sym_comment),
	12019: uint16(1039),
	12020: uint16(1),
	12021: uint16(anon_sym_DOLLAR),
	12022: uint16(1171),
	12023: uint16(1),
	12024: uint16(sym_variable_name),
	12025: uint16(274),
	12026: uint16(1),
	12027: uint16(sym_variable_expansion),
	12028: uint16(4),
	12029: uint16(771),
	12030: uint16(1),
	12031: uint16(sym_comment),
	12032: uint16(1045),
	12033: uint16(1),
	12034: uint16(anon_sym_DOLLAR),
	12035: uint16(1173),
	12036: uint16(1),
	12037: uint16(sym_variable_name),
	12038: uint16(81),
	12039: uint16(1),
	12040: uint16(sym_variable_expansion),
	12041: uint16(2),
	12042: uint16(771),
	12043: uint16(1),
	12044: uint16(sym_comment),
	12045: uint16(654),
	12046: uint16(3),
	12047: uint16(sym__brace_concat),
	12048: uint16(anon_sym_RBRACE),
	12049: uint16(anon_sym_COMMA),
	12050: uint16(2),
	12051: uint16(771),
	12052: uint16(1),
	12053: uint16(sym_comment),
	12054: uint16(672),
	12055: uint16(3),
	12056: uint16(sym__brace_concat),
	12057: uint16(anon_sym_RBRACE),
	12058: uint16(anon_sym_COMMA),
	12059: uint16(4),
	12060: uint16(771),
	12061: uint16(1),
	12062: uint16(sym_comment),
	12063: uint16(1167),
	12064: uint16(1),
	12065: uint16(anon_sym_COMMA),
	12066: uint16(1175),
	12067: uint16(1),
	12068: uint16(anon_sym_RBRACE),
	12069: uint16(392),
	12070: uint16(1),
	12071: uint16(aux_sym_brace_expansion_repeat1),
	12072: uint16(4),
	12073: uint16(771),
	12074: uint16(1),
	12075: uint16(sym_comment),
	12076: uint16(1167),
	12077: uint16(1),
	12078: uint16(anon_sym_COMMA),
	12079: uint16(1177),
	12080: uint16(1),
	12081: uint16(anon_sym_RBRACE),
	12082: uint16(392),
	12083: uint16(1),
	12084: uint16(aux_sym_brace_expansion_repeat1),
	12085: uint16(4),
	12086: uint16(771),
	12087: uint16(1),
	12088: uint16(sym_comment),
	12089: uint16(1179),
	12090: uint16(1),
	12091: uint16(anon_sym_RBRACE),
	12092: uint16(1181),
	12093: uint16(1),
	12094: uint16(anon_sym_COMMA),
	12095: uint16(392),
	12096: uint16(1),
	12097: uint16(aux_sym_brace_expansion_repeat1),
	12098: uint16(2),
	12099: uint16(771),
	12100: uint16(1),
	12101: uint16(sym_comment),
	12102: uint16(664),
	12103: uint16(3),
	12104: uint16(sym__brace_concat),
	12105: uint16(anon_sym_RBRACE),
	12106: uint16(anon_sym_COMMA),
	12107: uint16(2),
	12108: uint16(771),
	12109: uint16(1),
	12110: uint16(sym_comment),
	12111: uint16(660),
	12112: uint16(3),
	12113: uint16(sym__brace_concat),
	12114: uint16(anon_sym_RBRACE),
	12115: uint16(anon_sym_COMMA),
	12116: uint16(4),
	12117: uint16(771),
	12118: uint16(1),
	12119: uint16(sym_comment),
	12120: uint16(1167),
	12121: uint16(1),
	12122: uint16(anon_sym_COMMA),
	12123: uint16(1184),
	12124: uint16(1),
	12125: uint16(anon_sym_RBRACE),
	12126: uint16(392),
	12127: uint16(1),
	12128: uint16(aux_sym_brace_expansion_repeat1),
	12129: uint16(2),
	12130: uint16(771),
	12131: uint16(1),
	12132: uint16(sym_comment),
	12133: uint16(668),
	12134: uint16(3),
	12135: uint16(sym__brace_concat),
	12136: uint16(anon_sym_RBRACE),
	12137: uint16(anon_sym_COMMA),
	12138: uint16(4),
	12139: uint16(771),
	12140: uint16(1),
	12141: uint16(sym_comment),
	12142: uint16(1167),
	12143: uint16(1),
	12144: uint16(anon_sym_COMMA),
	12145: uint16(1186),
	12146: uint16(1),
	12147: uint16(anon_sym_RBRACE),
	12148: uint16(391),
	12149: uint16(1),
	12150: uint16(aux_sym_brace_expansion_repeat1),
	12151: uint16(4),
	12152: uint16(771),
	12153: uint16(1),
	12154: uint16(sym_comment),
	12155: uint16(1021),
	12156: uint16(1),
	12157: uint16(anon_sym_DOLLAR),
	12158: uint16(1188),
	12159: uint16(1),
	12160: uint16(sym_variable_name),
	12161: uint16(343),
	12162: uint16(1),
	12163: uint16(sym_variable_expansion),
	12164: uint16(4),
	12165: uint16(771),
	12166: uint16(1),
	12167: uint16(sym_comment),
	12168: uint16(1167),
	12169: uint16(1),
	12170: uint16(anon_sym_COMMA),
	12171: uint16(1190),
	12172: uint16(1),
	12173: uint16(anon_sym_RBRACE),
	12174: uint16(392),
	12175: uint16(1),
	12176: uint16(aux_sym_brace_expansion_repeat1),
	12177: uint16(4),
	12178: uint16(771),
	12179: uint16(1),
	12180: uint16(sym_comment),
	12181: uint16(1051),
	12182: uint16(1),
	12183: uint16(anon_sym_DOLLAR),
	12184: uint16(1192),
	12185: uint16(1),
	12186: uint16(sym_variable_name),
	12187: uint16(269),
	12188: uint16(1),
	12189: uint16(sym_variable_expansion),
	12190: uint16(4),
	12191: uint16(771),
	12192: uint16(1),
	12193: uint16(sym_comment),
	12194: uint16(1167),
	12195: uint16(1),
	12196: uint16(anon_sym_COMMA),
	12197: uint16(1194),
	12198: uint16(1),
	12199: uint16(anon_sym_RBRACE),
	12200: uint16(382),
	12201: uint16(1),
	12202: uint16(aux_sym_brace_expansion_repeat1),
	12203: uint16(4),
	12204: uint16(771),
	12205: uint16(1),
	12206: uint16(sym_comment),
	12207: uint16(1167),
	12208: uint16(1),
	12209: uint16(anon_sym_COMMA),
	12210: uint16(1194),
	12211: uint16(1),
	12212: uint16(anon_sym_RBRACE),
	12213: uint16(392),
	12214: uint16(1),
	12215: uint16(aux_sym_brace_expansion_repeat1),
	12216: uint16(4),
	12217: uint16(771),
	12218: uint16(1),
	12219: uint16(sym_comment),
	12220: uint16(1167),
	12221: uint16(1),
	12222: uint16(anon_sym_COMMA),
	12223: uint16(1190),
	12224: uint16(1),
	12225: uint16(anon_sym_RBRACE),
	12226: uint16(395),
	12227: uint16(1),
	12228: uint16(aux_sym_brace_expansion_repeat1),
	12229: uint16(4),
	12230: uint16(771),
	12231: uint16(1),
	12232: uint16(sym_comment),
	12233: uint16(1068),
	12234: uint16(1),
	12235: uint16(anon_sym_DOLLAR),
	12236: uint16(1196),
	12237: uint16(1),
	12238: uint16(sym_variable_name),
	12239: uint16(165),
	12240: uint16(1),
	12241: uint16(sym_variable_expansion),
	12242: uint16(2),
	12243: uint16(771),
	12244: uint16(1),
	12245: uint16(sym_comment),
	12246: uint16(676),
	12247: uint16(3),
	12248: uint16(sym__brace_concat),
	12249: uint16(anon_sym_RBRACE),
	12250: uint16(anon_sym_COMMA),
	12251: uint16(2),
	12252: uint16(771),
	12253: uint16(1),
	12254: uint16(sym_comment),
	12255: uint16(684),
	12256: uint16(3),
	12257: uint16(sym__brace_concat),
	12258: uint16(anon_sym_RBRACE),
	12259: uint16(anon_sym_COMMA),
	12260: uint16(2),
	12261: uint16(771),
	12262: uint16(1),
	12263: uint16(sym_comment),
	12264: uint16(1102),
	12265: uint16(3),
	12266: uint16(sym__brace_concat),
	12267: uint16(anon_sym_RBRACE),
	12268: uint16(anon_sym_COMMA),
	12269: uint16(4),
	12270: uint16(771),
	12271: uint16(1),
	12272: uint16(sym_comment),
	12273: uint16(1055),
	12274: uint16(1),
	12275: uint16(anon_sym_DOLLAR),
	12276: uint16(1198),
	12277: uint16(1),
	12278: uint16(sym_variable_name),
	12279: uint16(135),
	12280: uint16(1),
	12281: uint16(sym_variable_expansion),
	12282: uint16(4),
	12283: uint16(771),
	12284: uint16(1),
	12285: uint16(sym_comment),
	12286: uint16(1167),
	12287: uint16(1),
	12288: uint16(anon_sym_COMMA),
	12289: uint16(1200),
	12290: uint16(1),
	12291: uint16(anon_sym_RBRACE),
	12292: uint16(392),
	12293: uint16(1),
	12294: uint16(aux_sym_brace_expansion_repeat1),
	12295: uint16(4),
	12296: uint16(771),
	12297: uint16(1),
	12298: uint16(sym_comment),
	12299: uint16(1167),
	12300: uint16(1),
	12301: uint16(anon_sym_COMMA),
	12302: uint16(1186),
	12303: uint16(1),
	12304: uint16(anon_sym_RBRACE),
	12305: uint16(392),
	12306: uint16(1),
	12307: uint16(aux_sym_brace_expansion_repeat1),
	12308: uint16(2),
	12309: uint16(771),
	12310: uint16(1),
	12311: uint16(sym_comment),
	12312: uint16(688),
	12313: uint16(3),
	12314: uint16(sym__brace_concat),
	12315: uint16(anon_sym_RBRACE),
	12316: uint16(anon_sym_COMMA),
	12317: uint16(2),
	12318: uint16(771),
	12319: uint16(1),
	12320: uint16(sym_comment),
	12321: uint16(650),
	12322: uint16(3),
	12323: uint16(sym__brace_concat),
	12324: uint16(anon_sym_RBRACE),
	12325: uint16(anon_sym_COMMA),
	12326: uint16(4),
	12327: uint16(771),
	12328: uint16(1),
	12329: uint16(sym_comment),
	12330: uint16(1167),
	12331: uint16(1),
	12332: uint16(anon_sym_COMMA),
	12333: uint16(1202),
	12334: uint16(1),
	12335: uint16(anon_sym_RBRACE),
	12336: uint16(392),
	12337: uint16(1),
	12338: uint16(aux_sym_brace_expansion_repeat1),
	12339: uint16(4),
	12340: uint16(771),
	12341: uint16(1),
	12342: uint16(sym_comment),
	12343: uint16(1027),
	12344: uint16(1),
	12345: uint16(anon_sym_DOLLAR),
	12346: uint16(1204),
	12347: uint16(1),
	12348: uint16(sym_variable_name),
	12349: uint16(330),
	12350: uint16(1),
	12351: uint16(sym_variable_expansion),
	12352: uint16(4),
	12353: uint16(771),
	12354: uint16(1),
	12355: uint16(sym_comment),
	12356: uint16(1167),
	12357: uint16(1),
	12358: uint16(anon_sym_COMMA),
	12359: uint16(1202),
	12360: uint16(1),
	12361: uint16(anon_sym_RBRACE),
	12362: uint16(409),
	12363: uint16(1),
	12364: uint16(aux_sym_brace_expansion_repeat1),
	12365: uint16(2),
	12366: uint16(771),
	12367: uint16(1),
	12368: uint16(sym_comment),
	12369: uint16(646),
	12370: uint16(3),
	12371: uint16(sym__brace_concat),
	12372: uint16(anon_sym_RBRACE),
	12373: uint16(anon_sym_COMMA),
	12374: uint16(3),
	12375: uint16(771),
	12376: uint16(1),
	12377: uint16(sym_comment),
	12378: uint16(1206),
	12379: uint16(1),
	12380: uint16(anon_sym_LBRACK),
	12381: uint16(141),
	12382: uint16(1),
	12383: uint16(sym_list_element_access),
	12384: uint16(3),
	12385: uint16(771),
	12386: uint16(1),
	12387: uint16(sym_comment),
	12388: uint16(1208),
	12389: uint16(1),
	12390: uint16(anon_sym_LBRACK),
	12391: uint16(349),
	12392: uint16(1),
	12393: uint16(sym_list_element_access),
	12394: uint16(3),
	12395: uint16(771),
	12396: uint16(1),
	12397: uint16(sym_comment),
	12398: uint16(1210),
	12399: uint16(1),
	12400: uint16(anon_sym_LBRACK),
	12401: uint16(361),
	12402: uint16(1),
	12403: uint16(sym_list_element_access),
	12404: uint16(3),
	12405: uint16(771),
	12406: uint16(1),
	12407: uint16(sym_comment),
	12408: uint16(1212),
	12409: uint16(1),
	12410: uint16(anon_sym_LBRACK),
	12411: uint16(187),
	12412: uint16(1),
	12413: uint16(sym_list_element_access),
	12414: uint16(3),
	12415: uint16(771),
	12416: uint16(1),
	12417: uint16(sym_comment),
	12418: uint16(1214),
	12419: uint16(1),
	12420: uint16(anon_sym_LBRACK),
	12421: uint16(281),
	12422: uint16(1),
	12423: uint16(sym_list_element_access),
	12424: uint16(3),
	12425: uint16(771),
	12426: uint16(1),
	12427: uint16(sym_comment),
	12428: uint16(1216),
	12429: uint16(1),
	12430: uint16(anon_sym_LBRACK),
	12431: uint16(83),
	12432: uint16(1),
	12433: uint16(sym_list_element_access),
	12434: uint16(3),
	12435: uint16(771),
	12436: uint16(1),
	12437: uint16(sym_comment),
	12438: uint16(1218),
	12439: uint16(1),
	12440: uint16(anon_sym_LBRACK),
	12441: uint16(296),
	12442: uint16(1),
	12443: uint16(sym_list_element_access),
	12444: uint16(2),
	12445: uint16(771),
	12446: uint16(1),
	12447: uint16(sym_comment),
	12448: uint16(1179),
	12449: uint16(2),
	12450: uint16(anon_sym_RBRACE),
	12451: uint16(anon_sym_COMMA),
	12452: uint16(2),
	12453: uint16(771),
	12454: uint16(1),
	12455: uint16(sym_comment),
	12456: uint16(1033),
	12457: uint16(1),
	12458: uint16(anon_sym_end),
	12459: uint16(2),
	12460: uint16(771),
	12461: uint16(1),
	12462: uint16(sym_comment),
	12463: uint16(1066),
	12464: uint16(1),
	12465: uint16(anon_sym_end),
	12466: uint16(2),
	12467: uint16(771),
	12468: uint16(1),
	12469: uint16(sym_comment),
	12470: uint16(1220),
	12471: uint16(1),
	12472: uint16(anon_sym_if),
	12473: uint16(2),
	12474: uint16(771),
	12475: uint16(1),
	12476: uint16(sym_comment),
	12477: uint16(1222),
	12478: uint16(1),
	12479: uint16(anon_sym_in),
	12480: uint16(2),
	12481: uint16(771),
	12482: uint16(1),
	12483: uint16(sym_comment),
	12484: uint16(1224),
	12485: uint16(1),
	12486: uint16(sym_variable_name),
	12487: uint16(2),
	12488: uint16(771),
	12489: uint16(1),
	12490: uint16(sym_comment),
	12491: uint16(1226),
	12492: uint16(1),
	12494: uint16(2),
	12495: uint16(771),
	12496: uint16(1),
	12497: uint16(sym_comment),
	12498: uint16(1228),
	12499: uint16(1),
	12500: uint16(anon_sym_end),
}

var ts_small_parse_table_map = [398]uint32_t{
	1:   uint32(107),
	2:   uint32(214),
	3:   uint32(321),
	4:   uint32(427),
	5:   uint32(533),
	6:   uint32(639),
	7:   uint32(745),
	8:   uint32(851),
	9:   uint32(957),
	10:  uint32(1063),
	11:  uint32(1169),
	12:  uint32(1275),
	13:  uint32(1381),
	14:  uint32(1487),
	15:  uint32(1593),
	16:  uint32(1699),
	17:  uint32(1805),
	18:  uint32(1911),
	19:  uint32(2017),
	20:  uint32(2123),
	21:  uint32(2229),
	22:  uint32(2335),
	23:  uint32(2441),
	24:  uint32(2547),
	25:  uint32(2653),
	26:  uint32(2759),
	27:  uint32(2861),
	28:  uint32(2963),
	29:  uint32(3065),
	30:  uint32(3164),
	31:  uint32(3263),
	32:  uint32(3362),
	33:  uint32(3461),
	34:  uint32(3536),
	35:  uint32(3611),
	36:  uint32(3686),
	37:  uint32(3751),
	38:  uint32(3793),
	39:  uint32(3835),
	40:  uint32(3876),
	41:  uint32(3917),
	42:  uint32(3958),
	43:  uint32(3999),
	44:  uint32(4040),
	45:  uint32(4082),
	46:  uint32(4124),
	47:  uint32(4164),
	48:  uint32(4204),
	49:  uint32(4246),
	50:  uint32(4283),
	51:  uint32(4322),
	52:  uint32(4379),
	53:  uint32(4446),
	54:  uint32(4485),
	55:  uint32(4522),
	56:  uint32(4561),
	57:  uint32(4620),
	58:  uint32(4657),
	59:  uint32(4716),
	60:  uint32(4773),
	61:  uint32(4812),
	62:  uint32(4879),
	63:  uint32(4938),
	64:  uint32(4975),
	65:  uint32(5011),
	66:  uint32(5047),
	67:  uint32(5083),
	68:  uint32(5119),
	69:  uint32(5155),
	70:  uint32(5191),
	71:  uint32(5227),
	72:  uint32(5263),
	73:  uint32(5299),
	74:  uint32(5335),
	75:  uint32(5371),
	76:  uint32(5407),
	77:  uint32(5443),
	78:  uint32(5479),
	79:  uint32(5515),
	80:  uint32(5548),
	81:  uint32(5581),
	82:  uint32(5614),
	83:  uint32(5666),
	84:  uint32(5716),
	85:  uint32(5767),
	86:  uint32(5816),
	87:  uint32(5867),
	88:  uint32(5918),
	89:  uint32(5969),
	90:  uint32(6018),
	91:  uint32(6069),
	92:  uint32(6118),
	93:  uint32(6167),
	94:  uint32(6211),
	95:  uint32(6257),
	96:  uint32(6301),
	97:  uint32(6345),
	98:  uint32(6389),
	99:  uint32(6422),
	100: uint32(6455),
	101: uint32(6486),
	102: uint32(6517),
	103: uint32(6550),
	104: uint32(6580),
	105: uint32(6608),
	106: uint32(6638),
	107: uint32(6668),
	108: uint32(6696),
	109: uint32(6726),
	110: uint32(6754),
	111: uint32(6781),
	112: uint32(6808),
	113: uint32(6835),
	114: uint32(6862),
	115: uint32(6889),
	116: uint32(6916),
	117: uint32(6943),
	118: uint32(6970),
	119: uint32(6997),
	120: uint32(7024),
	121: uint32(7051),
	122: uint32(7078),
	123: uint32(7105),
	124: uint32(7132),
	125: uint32(7156),
	126: uint32(7194),
	127: uint32(7218),
	128: uint32(7247),
	129: uint32(7276),
	130: uint32(7305),
	131: uint32(7332),
	132: uint32(7359),
	133: uint32(7401),
	134: uint32(7443),
	135: uint32(7485),
	136: uint32(7527),
	137: uint32(7569),
	138: uint32(7611),
	139: uint32(7653),
	140: uint32(7695),
	141: uint32(7737),
	142: uint32(7763),
	143: uint32(7789),
	144: uint32(7831),
	145: uint32(7857),
	146: uint32(7881),
	147: uint32(7923),
	148: uint32(7947),
	149: uint32(7973),
	150: uint32(8015),
	151: uint32(8057),
	152: uint32(8099),
	153: uint32(8141),
	154: uint32(8165),
	155: uint32(8188),
	156: uint32(8211),
	157: uint32(8242),
	158: uint32(8265),
	159: uint32(8288),
	160: uint32(8311),
	161: uint32(8334),
	162: uint32(8357),
	163: uint32(8388),
	164: uint32(8411),
	165: uint32(8434),
	166: uint32(8457),
	167: uint32(8480),
	168: uint32(8503),
	169: uint32(8532),
	170: uint32(8555),
	171: uint32(8578),
	172: uint32(8601),
	173: uint32(8621),
	174: uint32(8653),
	175: uint32(8685),
	176: uint32(8705),
	177: uint32(8737),
	178: uint32(8757),
	179: uint32(8777),
	180: uint32(8797),
	181: uint32(8817),
	182: uint32(8837),
	183: uint32(8857),
	184: uint32(8877),
	185: uint32(8897),
	186: uint32(8917),
	187: uint32(8937),
	188: uint32(8957),
	189: uint32(8989),
	190: uint32(9009),
	191: uint32(9041),
	192: uint32(9061),
	193: uint32(9093),
	194: uint32(9113),
	195: uint32(9133),
	196: uint32(9165),
	197: uint32(9185),
	198: uint32(9205),
	199: uint32(9237),
	200: uint32(9257),
	201: uint32(9277),
	202: uint32(9309),
	203: uint32(9329),
	204: uint32(9349),
	205: uint32(9369),
	206: uint32(9401),
	207: uint32(9421),
	208: uint32(9441),
	209: uint32(9461),
	210: uint32(9481),
	211: uint32(9501),
	212: uint32(9521),
	213: uint32(9553),
	214: uint32(9589),
	215: uint32(9621),
	216: uint32(9641),
	217: uint32(9673),
	218: uint32(9693),
	219: uint32(9729),
	220: uint32(9761),
	221: uint32(9781),
	222: uint32(9801),
	223: uint32(9833),
	224: uint32(9853),
	225: uint32(9873),
	226: uint32(9905),
	227: uint32(9934),
	228: uint32(9963),
	229: uint32(9992),
	230: uint32(10021),
	231: uint32(10050),
	232: uint32(10079),
	233: uint32(10108),
	234: uint32(10137),
	235: uint32(10166),
	236: uint32(10183),
	237: uint32(10202),
	238: uint32(10219),
	239: uint32(10238),
	240: uint32(10257),
	241: uint32(10275),
	242: uint32(10293),
	243: uint32(10313),
	244: uint32(10333),
	245: uint32(10347),
	246: uint32(10361),
	247: uint32(10381),
	248: uint32(10395),
	249: uint32(10418),
	250: uint32(10433),
	251: uint32(10448),
	252: uint32(10461),
	253: uint32(10474),
	254: uint32(10497),
	255: uint32(10520),
	256: uint32(10543),
	257: uint32(10566),
	258: uint32(10579),
	259: uint32(10602),
	260: uint32(10625),
	261: uint32(10648),
	262: uint32(10671),
	263: uint32(10686),
	264: uint32(10699),
	265: uint32(10722),
	266: uint32(10745),
	267: uint32(10762),
	268: uint32(10775),
	269: uint32(10788),
	270: uint32(10805),
	271: uint32(10828),
	272: uint32(10841),
	273: uint32(10854),
	274: uint32(10871),
	275: uint32(10894),
	276: uint32(10911),
	277: uint32(10924),
	278: uint32(10939),
	279: uint32(10952),
	280: uint32(10965),
	281: uint32(10978),
	282: uint32(10996),
	283: uint32(11014),
	284: uint32(11028),
	285: uint32(11042),
	286: uint32(11056),
	287: uint32(11074),
	288: uint32(11088),
	289: uint32(11102),
	290: uint32(11116),
	291: uint32(11130),
	292: uint32(11144),
	293: uint32(11158),
	294: uint32(11172),
	295: uint32(11186),
	296: uint32(11200),
	297: uint32(11216),
	298: uint32(11232),
	299: uint32(11246),
	300: uint32(11260),
	301: uint32(11274),
	302: uint32(11293),
	303: uint32(11312),
	304: uint32(11325),
	305: uint32(11338),
	306: uint32(11355),
	307: uint32(11370),
	308: uint32(11389),
	309: uint32(11408),
	310: uint32(11421),
	311: uint32(11434),
	312: uint32(11453),
	313: uint32(11472),
	314: uint32(11487),
	315: uint32(11502),
	316: uint32(11515),
	317: uint32(11526),
	318: uint32(11543),
	319: uint32(11562),
	320: uint32(11578),
	321: uint32(11594),
	322: uint32(11610),
	323: uint32(11626),
	324: uint32(11640),
	325: uint32(11654),
	326: uint32(11668),
	327: uint32(11684),
	328: uint32(11694),
	329: uint32(11706),
	330: uint32(11720),
	331: uint32(11732),
	332: uint32(11744),
	333: uint32(11760),
	334: uint32(11776),
	335: uint32(11788),
	336: uint32(11804),
	337: uint32(11820),
	338: uint32(11836),
	339: uint32(11850),
	340: uint32(11866),
	341: uint32(11880),
	342: uint32(11896),
	343: uint32(11906),
	344: uint32(11920),
	345: uint32(11930),
	346: uint32(11946),
	347: uint32(11958),
	348: uint32(11971),
	349: uint32(11984),
	350: uint32(11993),
	351: uint32(12006),
	352: uint32(12015),
	353: uint32(12028),
	354: uint32(12041),
	355: uint32(12050),
	356: uint32(12059),
	357: uint32(12072),
	358: uint32(12085),
	359: uint32(12098),
	360: uint32(12107),
	361: uint32(12116),
	362: uint32(12129),
	363: uint32(12138),
	364: uint32(12151),
	365: uint32(12164),
	366: uint32(12177),
	367: uint32(12190),
	368: uint32(12203),
	369: uint32(12216),
	370: uint32(12229),
	371: uint32(12242),
	372: uint32(12251),
	373: uint32(12260),
	374: uint32(12269),
	375: uint32(12282),
	376: uint32(12295),
	377: uint32(12308),
	378: uint32(12317),
	379: uint32(12326),
	380: uint32(12339),
	381: uint32(12352),
	382: uint32(12365),
	383: uint32(12374),
	384: uint32(12384),
	385: uint32(12394),
	386: uint32(12404),
	387: uint32(12414),
	388: uint32(12424),
	389: uint32(12434),
	390: uint32(12444),
	391: uint32(12452),
	392: uint32(12459),
	393: uint32(12466),
	394: uint32(12473),
	395: uint32(12480),
	396: uint32(12487),
	397: uint32(12494),
}

var ts_parse_actions = [1230]TSParseActionEntry{
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
		Fcount: uint8(1),
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
		Fsymbol:     uint16(sym_program),
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
		Fstate: uint16(84),
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
		Fstate: uint16(66),
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
		Fstate: uint16(28),
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
		Fstate: uint16(63),
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
		Fcount: uint8(1),
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
		Fstate: uint16(342),
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
		Fstate: uint16(20),
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
		Fstate: uint16(126),
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
		Fstate: uint16(70),
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
		Fstate: uint16(119),
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
		Fstate: uint16(262),
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
		Fstate: uint16(429),
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
		Fstate: uint16(61),
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
		Fstate: uint16(62),
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
		Fstate: uint16(15),
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
		Fstate: uint16(100),
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
		Fstate: uint16(120),
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
		Fstate: uint16(288),
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
		Fstate: uint16(355),
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
		Fstate: uint16(8),
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
		Fstate: uint16(4),
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
		Fstate: uint16(233),
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
		Fstate: uint16(267),
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
		Fcount: uint8(1),
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
		Fstate: uint16(326),
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
		Fstate: uint16(2),
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
		Fstate: uint16(213),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
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
		Fstate:      uint16(84),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
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
		Fstate:      uint16(66),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
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
		Fstate:      uint16(4),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
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
		Fstate:      uint16(63),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
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
		Fstate:      uint16(342),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
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
		Fstate:      uint16(20),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
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
		Fstate:      uint16(126),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fcount: uint8(2),
	}})),
	81: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fstate:      uint16(70),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
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
		Fstate:      uint16(119),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
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
		Fstate:      uint16(267),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fstate:      uint16(429),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	93: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
	94: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
	97: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fstate:      uint16(15),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fstate:      uint16(100),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	105: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fstate:      uint16(120),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	108: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fstate:      uint16(288),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	111: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fstate:      uint16(355),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fstate:      uint16(8),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_else_if_clause),
		Fproduction_id: uint16(17),
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
		Fstate: uint16(5),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_else_if_clause),
		Fproduction_id: uint16(17),
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
		Fstate: uint16(9),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_else_clause),
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
		Fstate: uint16(261),
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
		Fstate: uint16(14),
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
		Fstate: uint16(222),
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
		Fstate: uint16(243),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fstate:      uint16(9),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	138: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fstate:      uint16(261),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fstate:      uint16(10),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_while_statement_repeat1),
	})))),
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
		Fstate:      uint16(265),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: uint16(237),
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
		Fcount: uint8(1),
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
		Fstate: uint16(13),
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
		Fstate: uint16(214),
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
		Fcount: uint8(1),
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
		Fstate: uint16(206),
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
		Fcount: uint8(1),
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
		Fstate: uint16(10),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(11),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_else_clause),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(253),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(19),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(21),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(31),
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
		Fstate: uint16(106),
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
		Fcount: uint8(1),
	}})),
	187: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	189: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	190: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      uint16(66),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	195: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	196: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	198: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      uint16(63),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      uint16(342),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      uint16(20),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(126),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(70),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(119),
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
		Fcount: uint8(2),
	}})),
	218: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      uint16(264),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	221: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      uint16(429),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(61),
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
		Fcount: uint8(2),
	}})),
	227: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	228: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      uint16(15),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	233: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	234: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(100),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	235: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	236: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      uint16(120),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	240: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(288),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(355),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fstate:      uint16(8),
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
		Fcount: uint8(1),
	}})),
	248: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	249: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(365),
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
		Fcount: uint8(1),
	}})),
	252: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	253: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(405),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(235),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	259: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	260: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      uint16(24),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	263: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
	264: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(262),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	265: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(362),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(248),
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
		Fcount: uint8(1),
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
		Fstate: uint16(27),
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
		Fcount: uint8(1),
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
		Fstate: uint16(154),
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
		Fstate: uint16(239),
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
		Fstate: uint16(153),
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
		Fcount: uint8(1),
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
		Fstate: uint16(259),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_program),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(24),
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
		Fcount: uint8(1),
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
		Fstate: uint16(23),
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
		Fstate: uint16(393),
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
		Fcount: uint8(1),
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
		Fstate: uint16(226),
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
		Fcount: uint8(1),
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
		Fstate: uint16(317),
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
		Fcount: uint8(1),
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
		Fstate: uint16(208),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(109),
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
		Fcount: uint8(1),
	}})),
	296: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	298: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	302: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(229),
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
		Fcount: uint8(1),
	}})),
	308: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_case_clause),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	312: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_case_clause),
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
		Fcount: uint8(2),
	}})),
	314: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(84),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	317: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
	318: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(66),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	320: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(63),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	323: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
	324: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(342),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	326: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(20),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	329: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(126),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
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
		Fstate:      uint16(70),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(119),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	340: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(260),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(429),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	346: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(61),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	349: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(62),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(15),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	355: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(100),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	358: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(120),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(288),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	364: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(355),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	367: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(8),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: uint16(254),
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
		Fcount: uint8(1),
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
		Fstate: uint16(263),
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
		Fcount: uint8(1),
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
		Fstate: uint16(258),
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
		Fcount: uint8(1),
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
		Fstate: uint16(220),
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
		Fstate: uint16(228),
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
		Fcount: uint8(1),
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
		Fstate: uint16(240),
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
		Fcount: uint8(1),
	}})),
	382: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(241),
	}})))),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(219),
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
		Fcount: uint8(1),
	}})),
	388: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	390: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(217),
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
		Fcount: uint8(1),
	}})),
	394: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(242),
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
		Fcount: uint8(1),
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
		Fstate: uint16(215),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(236),
	}})))),
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
		Fcount: uint8(1),
	}})),
	402: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(221),
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
		Fcount: uint8(1),
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
		Fstate: uint16(209),
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
		Fcount: uint8(1),
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
		Fstate: uint16(227),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_function_definition_repeat2),
	})))),
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
		Fstate:      uint16(263),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	415: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(268),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(202),
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
		Fcount: uint8(1),
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
		Fstate: uint16(191),
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
		Fcount: uint8(1),
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
		Fstate: uint16(190),
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
		Fcount: uint8(1),
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
		Fstate: uint16(196),
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
		Fcount: uint8(2),
	}})),
	431: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_command_repeat1),
		Fproduction_id: uint16(10),
	})))),
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
		Fstate:      uint16(84),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	434: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_command_repeat1),
		Fproduction_id: uint16(10),
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
		Fcount: uint8(2),
	}})),
	436: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_command_repeat1),
		Fproduction_id: uint16(10),
	})))),
	437: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(342),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	439: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_command_repeat1),
		Fproduction_id: uint16(10),
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
		Fstate:      uint16(20),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_command_repeat1),
		Fproduction_id: uint16(10),
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
		Fstate:      uint16(100),
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
		Fcount: uint8(2),
	}})),
	445: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_command_repeat1),
		Fproduction_id: uint16(10),
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
		Fstate:      uint16(120),
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
		Fcount: uint8(2),
	}})),
	448: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_command_repeat1),
		Fproduction_id: uint16(10),
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
		Fstate:      uint16(288),
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
		Fcount: uint8(2),
	}})),
	451: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_command_repeat1),
		Fproduction_id: uint16(10),
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
		Fstate:      uint16(355),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_command_repeat1),
		Fproduction_id: uint16(10),
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
		Fstate:      uint16(115),
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
		Fcount: uint8(2),
	}})),
	457: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_command_repeat1),
		Fproduction_id: uint16(10),
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
		Fstate:      uint16(125),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_command),
		Fproduction_id: uint16(6),
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
		Fcount: uint8(1),
	}})),
	462: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_command),
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
		Fcount: uint8(1),
	}})),
	468: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	470: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_return),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(17),
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
		Fcount: uint8(1),
	}})),
	476: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(360),
	}})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__terminated_opt_statement),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__terminated_opt_statement),
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
		Fsymbol:      uint16(sym__terminated_statement),
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
		Fsymbol:      uint16(sym__terminated_statement),
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
		Fcount: uint8(1),
	}})),
	492: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_variable_expansion),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_variable_expansion),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(422),
	}})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_variable_expansion),
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
		Fsymbol:      uint16(sym_variable_expansion),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_variable_expansion_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_variable_expansion_repeat1),
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
		Fsymbol:      uint16(aux_sym_variable_expansion_repeat1),
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
		Fstate:      uint16(422),
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
		Fcount: uint8(1),
	}})),
	509: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: uint16(129),
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
		Fcount: uint8(1),
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
		Fstate: uint16(137),
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
		Fcount: uint8(1),
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
		Fstate: uint16(35),
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
		Fstate: uint16(346),
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
		Fcount: uint8(1),
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
		Fstate: uint16(26),
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
		Fcount: uint8(1),
	}})),
	521: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	525: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(37),
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
		Fcount: uint8(1),
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
		Fstate: uint16(55),
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
		Fcount: uint8(1),
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
		Fstate: uint16(38),
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
		Fcount: uint8(1),
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
		Fstate: uint16(53),
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
		Fcount: uint8(1),
	}})),
	537: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__expression),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	541: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_list_element_access),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	543: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_list_element_access),
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
		Fcount: uint8(1),
	}})),
	545: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_concatenation_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	547: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_concatenation_repeat1),
	})))),
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
		Fstate:      uint16(129),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	550: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_for_statement_repeat1),
		Fproduction_id: uint16(16),
	})))),
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
		Fstate:      uint16(137),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fsymbol:        uint16(aux_sym_for_statement_repeat1),
		Fproduction_id: uint16(16),
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
		Fcount: uint8(2),
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
		Fsymbol:        uint16(aux_sym_for_statement_repeat1),
		Fproduction_id: uint16(16),
	})))),
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
		Fstate:      uint16(346),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	558: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_for_statement_repeat1),
		Fproduction_id: uint16(16),
	})))),
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
		Fstate:      uint16(26),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	561: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_for_statement_repeat1),
		Fproduction_id: uint16(16),
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
		Fstate:      uint16(155),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	564: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_for_statement_repeat1),
		Fproduction_id: uint16(16),
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
		Fstate:      uint16(124),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	567: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_for_statement_repeat1),
		Fproduction_id: uint16(16),
	})))),
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
		Fstate:      uint16(308),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	570: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_for_statement_repeat1),
		Fproduction_id: uint16(16),
	})))),
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
		Fstate:      uint16(354),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	573: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list_element_access),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list_element_access),
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
		Fcount: uint8(1),
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
		Fstate: uint16(54),
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
		Fcount: uint8(2),
	}})),
	579: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_case_clause_repeat1),
	})))),
	580: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	582: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_case_clause_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_case_clause_repeat1),
	})))),
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
		Fstate:      uint16(346),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_case_clause_repeat1),
	})))),
	588: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(26),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	589: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	590: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_case_clause_repeat1),
	})))),
	591: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(155),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_case_clause_repeat1),
	})))),
	594: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	595: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	596: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_case_clause_repeat1),
	})))),
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
		Fstate:      uint16(308),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_case_clause_repeat1),
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
		Fstate:      uint16(354),
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
		Fcount: uint8(1),
	}})),
	602: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_concatenation),
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
		Fcount: uint8(1),
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
		Fstate: uint16(51),
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
		Fcount: uint8(1),
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
		Fstate: uint16(58),
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
		Fcount: uint8(1),
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
		Fstate: uint16(42),
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
		Fcount: uint8(1),
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
		Fstate: uint16(40),
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
		Fcount: uint8(1),
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
		Fstate: uint16(56),
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
		Fcount: uint8(2),
	}})),
	614: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_function_definition_repeat1),
		Fproduction_id: uint16(12),
	})))),
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
		Fstate:      uint16(137),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	617: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_function_definition_repeat1),
		Fproduction_id: uint16(12),
	})))),
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
		Fcount: uint8(2),
	}})),
	619: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_function_definition_repeat1),
		Fproduction_id: uint16(12),
	})))),
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
		Fstate:      uint16(346),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	622: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_function_definition_repeat1),
		Fproduction_id: uint16(12),
	})))),
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
		Fstate:      uint16(26),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_function_definition_repeat1),
		Fproduction_id: uint16(12),
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
		Fstate:      uint16(155),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_function_definition_repeat1),
		Fproduction_id: uint16(12),
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
		Fstate:      uint16(124),
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
		Fcount: uint8(2),
	}})),
	631: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_function_definition_repeat1),
		Fproduction_id: uint16(12),
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
		Fstate:      uint16(308),
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
		Fcount: uint8(2),
	}})),
	634: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_function_definition_repeat1),
		Fproduction_id: uint16(12),
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
		Fstate:      uint16(354),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__command_substitution_dollar),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__command_substitution_dollar),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__special_character),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__special_character),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_double_quote_string),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_double_quote_string),
	})))),
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
		Fcount: uint8(1),
	}})),
	649: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_single_quote_string),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_single_quote_string),
	})))),
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
		Fcount: uint8(1),
	}})),
	653: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_command_substitution),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	655: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_command_substitution),
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
		Fsymbol:      uint16(aux_sym_concatenation_repeat1),
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
		Fcount: uint8(1),
	}})),
	659: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_brace_expansion),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_brace_expansion),
	})))),
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
		Fcount: uint8(1),
	}})),
	663: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__command_substitution_inner),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__command_substitution_inner),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__command_substitution_inner),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__command_substitution_inner),
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
		Fcount: uint8(1),
	}})),
	671: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_single_quote_string),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_single_quote_string),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__command_substitution_inner),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	677: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__command_substitution_inner),
	})))),
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
		Fcount: uint8(1),
	}})),
	679: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_double_quote_string),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	681: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_double_quote_string),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_brace_expansion),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_brace_expansion),
	})))),
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
		Fcount: uint8(1),
	}})),
	687: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_brace_expansion),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_brace_expansion),
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
		Fcount: uint8(1),
	}})),
	691: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_command_repeat1),
		Fproduction_id: uint16(5),
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
		Fcount: uint8(1),
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
		Fsymbol:        uint16(sym_file_redirect),
		Fproduction_id: uint16(8),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_command_repeat1),
		Fproduction_id: uint16(4),
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
		Fcount: uint8(1),
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
		Fstate: uint16(335),
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
		Fcount: uint8(1),
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
		Fstate: uint16(29),
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
		Fcount: uint8(1),
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
		Fstate: uint16(357),
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
		Fcount: uint8(1),
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
		Fstate: uint16(201),
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
		Fcount: uint8(1),
	}})),
	705: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	709: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(367),
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
		Fcount: uint8(1),
	}})),
	713: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	715: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(32),
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
		Fcount: uint8(1),
	}})),
	719: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	723: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(379),
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
		Fcount: uint8(1),
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
		Fstate: uint16(111),
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
		Fcount: uint8(1),
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
		Fstate: uint16(406),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(334),
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
		Fcount: uint8(1),
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
		Fstate: uint16(151),
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
		Fcount: uint8(1),
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
		Fstate: uint16(192),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_brace_expansion_repeat1),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(104),
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
		Fcount: uint8(1),
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
		Fstate: uint16(156),
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
		Fstate: uint16(319),
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
		Fstate: uint16(417),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	747: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_variable_expansion_repeat1),
	})))),
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
		Fstate:      uint16(417),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_concatenation_repeat1),
	})))),
	753: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(130),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(1),
	}})),
	755: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_for_statement_repeat1),
		Fproduction_id: uint16(14),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(407),
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
		Fcount: uint8(1),
	}})),
	759: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_function_definition_repeat1),
		Fproduction_id: uint16(7),
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
		Fsymbol:      uint16(aux_sym_variable_expansion_repeat1),
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
		Fstate:      uint16(420),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(420),
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
		Fstate: uint16(345),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(285),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(252),
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
		Fstate: uint16(338),
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
		Fstate: uint16(289),
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
		Fstate: uint16(366),
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
		Fstate: uint16(311),
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
		Fstate: uint16(278),
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
		Fstate: uint16(143),
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
		Fstate: uint16(279),
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
		Fstate: uint16(378),
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
		Fstate: uint16(376),
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
		Fstate: uint16(138),
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
		Fstate: uint16(283),
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
		Fstate: uint16(127),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_concatenation_repeat1),
	})))),
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
		Fstate:      uint16(127),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(179),
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
		Fstate: uint16(91),
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
		Fstate: uint16(337),
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
		Fstate: uint16(181),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	811: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_element_access_repeat1),
	})))),
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
		Fstate:      uint16(345),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	814: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_element_access_repeat1),
	})))),
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
		Fstate:      uint16(18),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	817: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_element_access_repeat1),
	})))),
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
		Fstate:      uint16(285),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_element_access_repeat1),
	})))),
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
		Fstate:      uint16(252),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_element_access_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_list_element_access_repeat1),
	})))),
	826: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(289),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	828: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_list_element_access_repeat1),
	})))),
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
		Fstate:      uint16(366),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(88),
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
		Fcount: uint8(1),
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
		Fstate: uint16(65),
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
		Fcount: uint8(1),
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
		Fstate: uint16(64),
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
		Fcount: uint8(1),
	}})),
	837: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_conditional_execution),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(251),
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
		Fcount: uint8(1),
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
		Fstate: uint16(123),
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
		Fcount: uint8(1),
	}})),
	843: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_pipe),
	})))),
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
		Fcount: uint8(1),
	}})),
	845: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_conditional_execution),
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
		Fcount: uint8(1),
	}})),
	847: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_negated_statement),
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
		Fcount: uint8(1),
	}})),
	849: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_while_statement),
		Fproduction_id: uint16(9),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(74),
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
		Fcount: uint8(1),
	}})),
	853: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	855: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_function_definition),
		Fproduction_id: uint16(15),
	})))),
	856: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	857: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_begin_statement),
	})))),
	858: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	859: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(9),
	})))),
	860: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	861: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_while_statement),
		Fproduction_id: uint16(9),
	})))),
	862: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	863: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_for_statement),
		Fproduction_id: uint16(18),
	})))),
	864: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	865: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_function_definition),
		Fproduction_id: uint16(11),
	})))),
	866: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	867: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	868: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	869: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(9),
	})))),
	870: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: uint16(107),
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
		Fcount: uint8(1),
	}})),
	873: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_switch_statement),
		Fproduction_id: uint16(13),
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
		Fcount: uint8(1),
	}})),
	875: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_for_statement),
		Fproduction_id: uint16(18),
	})))),
	876: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	877: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	878: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	879: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_switch_statement),
		Fproduction_id: uint16(13),
	})))),
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
		Fcount: uint8(1),
	}})),
	881: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	883: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(9),
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
		Fcount: uint8(1),
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
		Fstate: uint16(396),
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
		Fcount: uint8(1),
	}})),
	887: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_begin_statement),
	})))),
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
		Fcount: uint8(1),
	}})),
	889: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_function_definition),
		Fproduction_id: uint16(11),
	})))),
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
		Fcount: uint8(1),
	}})),
	891: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_function_definition),
		Fproduction_id: uint16(15),
	})))),
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
		Fcount: uint8(1),
	}})),
	893: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_begin_statement),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_range),
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
		Fcount: uint8(1),
	}})),
	897: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(380),
	}})))),
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
		Fcount: uint8(1),
	}})),
	899: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(9),
	})))),
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
		Fcount: uint8(1),
	}})),
	901: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_redirect_statement),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_range),
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
		Fcount: uint8(1),
	}})),
	905: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_return),
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
		Fcount: uint8(1),
	}})),
	907: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	909: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	911: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	913: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(76),
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
		Fcount: uint8(1),
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
		Fstate: uint16(98),
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
		Fcount: uint8(1),
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
		Fstate: uint16(77),
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
		Fcount: uint8(1),
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
		Fstate: uint16(71),
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
		Fcount: uint8(1),
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
		Fstate: uint16(72),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(421),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	927: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_variable_expansion_repeat1),
	})))),
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
		Fstate:      uint16(421),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(423),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_variable_expansion_repeat1),
	})))),
	933: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(423),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(1),
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
		Fstate: uint16(336),
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
		Fcount: uint8(1),
	}})),
	937: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: uint16(298),
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
		Fcount: uint8(1),
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
		Fstate: uint16(298),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_list_element_access_repeat1),
	})))),
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
		Fstate: uint16(247),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_index),
	})))),
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
		Fcount: uint8(1),
	}})),
	949: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(305),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(293),
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
		Fcount: uint8(1),
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
		Fstate: uint16(110),
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
		Fstate: uint16(299),
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
		Fcount: uint8(1),
	}})),
	959: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(301),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: uint16(287),
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
		Fcount: uint8(1),
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
		Fstate: uint16(416),
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
		Fcount: uint8(1),
	}})),
	971: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(383),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	975: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	977: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_double_quote_string_repeat1),
	})))),
	978: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	980: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_double_quote_string_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	982: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_double_quote_string_repeat1),
	})))),
	983: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	985: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_double_quote_string_repeat1),
	})))),
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
		Fstate:      uint16(293),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: uint16(332),
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
		Fstate: uint16(329),
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
		Fstate: uint16(294),
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
		Fcount: uint8(1),
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
		Fstate: uint16(294),
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
		Fcount: uint8(1),
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
		Fstate: uint16(193),
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
		Fcount: uint8(1),
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
		Fstate: uint16(101),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1002: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_concatenation_repeat1),
	})))),
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
		Fstate:      uint16(131),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1004: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1005: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1006: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1007: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1008: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1009: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1010: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1011: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1012: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1013: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(418),
	}})))),
	1014: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1015: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_variable_expansion_repeat1),
	})))),
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
		Fstate:      uint16(418),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: uint16(16),
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
		Fcount: uint8(1),
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
		Fstate: uint16(60),
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
		Fstate: uint16(398),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(29),
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
		Fstate: uint16(340),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(414),
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
		Fstate: uint16(22),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(316),
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
		Fstate: uint16(249),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(326),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(419),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(386),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(32),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(280),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(387),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(20),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(79),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(400),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(272),
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
		Fstate: uint16(408),
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
		Fstate: uint16(26),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: uint16(133),
	}})))),
	1061: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1062: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_variable_expansion_repeat1),
	})))),
	1063: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(419),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1064: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1065: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(374),
	}})))),
	1066: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1067: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1068: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1069: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(404),
	}})))),
	1070: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1071: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1072: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1073: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1074: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1075: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(412),
	}})))),
	1076: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1077: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(370),
	}})))),
	1078: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1079: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(370),
	}})))),
	1080: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1081: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1082: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1083: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1084: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1085: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1086: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1087: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1088: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1089: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(371),
	}})))),
	1090: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1091: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(371),
	}})))),
	1092: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1093: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1094: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1095: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__brace_expression),
	})))),
	1096: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1097: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1098: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1099: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1100: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1101: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1102: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1103: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_brace_concatenation_repeat1),
	})))),
	1104: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1105: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_brace_concatenation_repeat1),
	})))),
	1106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(159),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1107: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1108: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1109: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1110: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(369),
	}})))),
	1111: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1112: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(369),
	}})))),
	1113: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1114: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_if_statement_repeat1),
	})))),
	1115: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1116: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_if_statement_repeat1),
	})))),
	1117: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(427),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1118: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1119: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_double_quote_string_repeat1),
		Fproduction_id: uint16(3),
	})))),
	1120: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1121: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_double_quote_string_repeat1),
		Fproduction_id: uint16(3),
	})))),
	1122: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1123: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1124: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1125: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(375),
	}})))),
	1126: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1127: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(375),
	}})))),
	1128: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1129: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(389),
	}})))),
	1130: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1131: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1132: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1133: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1134: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1135: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1136: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1137: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_single_quote_string_repeat1),
	})))),
	1138: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1139: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_single_quote_string_repeat1),
	})))),
	1140: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(370),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1141: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1142: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_single_quote_string_repeat1),
	})))),
	1143: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(370),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1144: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1145: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1146: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1147: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_brace_concatenation),
	})))),
	1148: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1149: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1150: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1151: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1152: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1153: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1154: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1155: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	1156: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1157: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	1158: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(117),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1159: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1160: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(328),
	}})))),
	1161: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1162: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(373),
	}})))),
	1163: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1164: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(373),
	}})))),
	1165: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1166: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(333),
	}})))),
	1167: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1168: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1169: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(394),
	}})))),
	1171: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1172: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1173: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1174: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1175: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1176: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1177: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1178: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1179: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1180: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_brace_expansion_repeat1),
	})))),
	1181: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1182: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_brace_expansion_repeat1),
	})))),
	1183: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      uint16(128),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1184: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1185: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1186: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1187: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1188: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1189: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1190: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1191: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1192: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1193: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1194: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1195: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(411),
	}})))),
	1196: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1197: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1198: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1199: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1200: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1201: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1202: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1203: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(152),
	}})))),
	1204: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1205: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1206: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1207: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1208: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1209: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1210: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1211: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1212: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1213: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1214: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1215: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1216: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1217: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1218: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1219: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1220: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1221: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1222: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1223: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1224: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1225: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: uint16(428),
	}})))),
	1226: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1227: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	1228: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1229: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

func tree_sitter_fish(tls *libc.TLS) (r uintptr) {
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
	Fkeyword_capture_token:     uint16(sym_word),
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
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_fish_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_fish_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_fish_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_fish_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_fish_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00word\x00and\x00or\x00||\x00&&\x00&|\x002>|\x00|\x00;\x00&\x00\n\x00\r\x00\r\n\x00!\x00not\x00$\x00(\x00)\x00function\x00integer\x00float\x00return\x00switch\x00case\x00break\x00continue\x00for\x00in\x00while\x00if\x00else\x00begin\x00}\x00comment\x00variable_name\x00..\x00[\x00]\x00{\x00,\x00\"\x00double_quote_string_token1\x00'\x00single_quote_string_token1\x00escape_sequence\x00stream_redirect\x00direction\x00#\x00home_dir_expansion\x00glob\x00_concat\x00_brace_concat\x00_concat_list\x00program\x00conditional_execution\x00pipe\x00redirect_statement\x00_terminated_statement\x00_terminated_opt_statement\x00negated_statement\x00_command_substitution_dollar\x00_command_substitution_inner\x00command_substitution\x00function_definition\x00switch_statement\x00case_clause\x00for_statement\x00while_statement\x00if_statement\x00else_if_clause\x00else_clause\x00begin_statement\x00variable_expansion\x00index\x00range\x00list_element_access\x00brace_expansion\x00double_quote_string\x00single_quote_string\x00command\x00file_redirect\x00_special_character\x00concatenation\x00_expression\x00_brace_expression\x00_base_brace_expression\x00program_repeat1\x00function_definition_repeat1\x00function_definition_repeat2\x00switch_statement_repeat1\x00case_clause_repeat1\x00for_statement_repeat1\x00while_statement_repeat1\x00if_statement_repeat1\x00variable_expansion_repeat1\x00list_element_access_repeat1\x00brace_expansion_repeat1\x00double_quote_string_repeat1\x00single_quote_string_repeat1\x00command_repeat1\x00concatenation_repeat1\x00brace_concatenation_repeat1\x00argument\x00condition\x00destination\x00name\x00operator\x00option\x00redirect\x00value\x00variable\x00"
