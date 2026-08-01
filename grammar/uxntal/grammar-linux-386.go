// Code generated for linux/386 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-uxntal/src -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-uxntal -I /home/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /home/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build linux && 386

package grammar_uxntal

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
const FIELD_COUNT = 6
const INT16_MAX = 0x7fff
const INT32_MAX = 0x7fffffff
const INT64_MAX = 0x7fffffffffffffff
const INT8_MAX = 0x7f
const INTMAX_MAX = "INT64_MAX"
const INTMAX_MIN = "INT64_MIN"
const INTPTR_MAX = "INT32_MAX"
const INTPTR_MIN = "INT32_MIN"
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
const LARGE_STATE_COUNT = 25
const MAX_ALIAS_SEQUENCE_LENGTH = 5
const PRODUCTION_ID_COUNT = 7
const PTRDIFF_MAX = "INT32_MAX"
const PTRDIFF_MIN = "INT32_MIN"
const RAND_MAX = 0x7fffffff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT32_MAX"
const STATE_COUNT = 48
const SYMBOL_COUNT = 305
const TOKEN_COUNT = 286
const TREE_SITTER_SERIALIZATION_BUFFER_SIZE = 1024
const UINT16_MAX = 0xffff
const UINT32_MAX = "0xffffffffu"
const UINT64_MAX = "0xffffffffffffffffu"
const UINT8_MAX = 0xff
const UINTMAX_MAX = "UINT64_MAX"
const UINTPTR_MAX = "UINT32_MAX"
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
const _FILE_OFFSET_BITS = 64
const _GNU_SOURCE = 1
const _LP64 = 1
const _REDIR_TIME64 = 1
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
const __LONG_MAX = 0x7fffffff
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

type __predefined_size_t = uint32

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int32

type uintptr_t = uint32

type intptr_t = int32

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

type size_t = uint32

type div_t = struct {
	Fquot int32
	Frem  int32
}

type ldiv_t = struct {
	Fquot int32
	Frem  int32
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

type wctype_t = uint32

type locale_t = uintptr

type wctrans_t = uintptr

type TokenType = int32

const COMMENT = 0

func tree_sitter_uxntal_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_uxntal_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
}

func tree_sitter_uxntal_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_uxntal_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
}

func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
}

func tree_sitter_uxntal_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var is_in_string uint8
	var nesting_depth uint32
	_, _ = is_in_string, nesting_depth
	is_in_string = libc.BoolUint8(false1 != 0)
	for libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
		advance(tls, lexer)
	}
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('"') {
		is_in_string = libc.BoolUint8(true1 != 0)
	}
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('(') && !(is_in_string != 0) {
		advance(tls, lexer)
		nesting_depth = uint32(1)
		for {
			switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
			case int32('\000'):
				return libc.BoolUint8(false1 != 0)
			case int32('('):
				advance(tls, lexer)
				nesting_depth = nesting_depth + 1
			case int32(')'):
				advance(tls, lexer)
				nesting_depth = nesting_depth - 1
				if nesting_depth == uint32(0) {
					(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(COMMENT)
					return libc.BoolUint8(true1 != 0)
				}
			default:
				advance(tls, lexer)
				break
			}
			goto _1
		_1:
		}
	}
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('"') {
		is_in_string = libc.BoolUint8(false1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

type ts_symbol_identifiers = int32

const sym_identifier = 1
const anon_sym_PERCENT = 2
const aux_sym_macro_token1 = 3
const anon_sym_LBRACE = 4
const anon_sym_RBRACE = 5
const anon_sym_TILDE = 6
const aux_sym_include_token1 = 7
const anon_sym_BRK = 8
const anon_sym_INC = 9
const anon_sym_POP = 10
const anon_sym_NIP = 11
const anon_sym_SWP = 12
const anon_sym_ROT = 13
const anon_sym_DUP = 14
const anon_sym_OVR = 15
const anon_sym_EQU = 16
const anon_sym_NEQ = 17
const anon_sym_GTH = 18
const anon_sym_LTH = 19
const anon_sym_JMP = 20
const anon_sym_JCN = 21
const anon_sym_JSR = 22
const anon_sym_STH = 23
const anon_sym_LDZ = 24
const anon_sym_STZ = 25
const anon_sym_LDR = 26
const anon_sym_STR = 27
const anon_sym_LDA = 28
const anon_sym_STA = 29
const anon_sym_DEI = 30
const anon_sym_DEO = 31
const anon_sym_ADD = 32
const anon_sym_SUB = 33
const anon_sym_MUL = 34
const anon_sym_DIV = 35
const anon_sym_AND = 36
const anon_sym_ORA = 37
const anon_sym_EOR = 38
const anon_sym_SFT = 39
const anon_sym_JCI = 40
const anon_sym_INC2 = 41
const anon_sym_POP2 = 42
const anon_sym_NIP2 = 43
const anon_sym_SWP2 = 44
const anon_sym_ROT2 = 45
const anon_sym_DUP2 = 46
const anon_sym_OVR2 = 47
const anon_sym_EQU2 = 48
const anon_sym_NEQ2 = 49
const anon_sym_GTH2 = 50
const anon_sym_LTH2 = 51
const anon_sym_JMP2 = 52
const anon_sym_JCN2 = 53
const anon_sym_JSR2 = 54
const anon_sym_STH2 = 55
const anon_sym_LDZ2 = 56
const anon_sym_STZ2 = 57
const anon_sym_LDR2 = 58
const anon_sym_STR2 = 59
const anon_sym_LDA2 = 60
const anon_sym_STA2 = 61
const anon_sym_DEI2 = 62
const anon_sym_DEO2 = 63
const anon_sym_ADD2 = 64
const anon_sym_SUB2 = 65
const anon_sym_MUL2 = 66
const anon_sym_DIV2 = 67
const anon_sym_AND2 = 68
const anon_sym_ORA2 = 69
const anon_sym_EOR2 = 70
const anon_sym_SFT2 = 71
const anon_sym_JMI = 72
const anon_sym_INCr = 73
const anon_sym_POPr = 74
const anon_sym_NIPr = 75
const anon_sym_SWPr = 76
const anon_sym_ROTr = 77
const anon_sym_DUPr = 78
const anon_sym_OVRr = 79
const anon_sym_EQUr = 80
const anon_sym_NEQr = 81
const anon_sym_GTHr = 82
const anon_sym_LTHr = 83
const anon_sym_JMPr = 84
const anon_sym_JCNr = 85
const anon_sym_JSRr = 86
const anon_sym_STHr = 87
const anon_sym_LDZr = 88
const anon_sym_STZr = 89
const anon_sym_LDRr = 90
const anon_sym_STRr = 91
const anon_sym_LDAr = 92
const anon_sym_STAr = 93
const anon_sym_DEIr = 94
const anon_sym_DEOr = 95
const anon_sym_ADDr = 96
const anon_sym_SUBr = 97
const anon_sym_MULr = 98
const anon_sym_DIVr = 99
const anon_sym_ANDr = 100
const anon_sym_ORAr = 101
const anon_sym_EORr = 102
const anon_sym_SFTr = 103
const anon_sym_JSI = 104
const anon_sym_INC2r = 105
const anon_sym_POP2r = 106
const anon_sym_NIP2r = 107
const anon_sym_SWP2r = 108
const anon_sym_ROT2r = 109
const anon_sym_DUP2r = 110
const anon_sym_OVR2r = 111
const anon_sym_EQU2r = 112
const anon_sym_NEQ2r = 113
const anon_sym_GTH2r = 114
const anon_sym_LTH2r = 115
const anon_sym_JMP2r = 116
const anon_sym_JCN2r = 117
const anon_sym_JSR2r = 118
const anon_sym_STH2r = 119
const anon_sym_LDZ2r = 120
const anon_sym_STZ2r = 121
const anon_sym_LDR2r = 122
const anon_sym_STR2r = 123
const anon_sym_LDA2r = 124
const anon_sym_STA2r = 125
const anon_sym_DEI2r = 126
const anon_sym_DEO2r = 127
const anon_sym_ADD2r = 128
const anon_sym_SUB2r = 129
const anon_sym_MUL2r = 130
const anon_sym_DIV2r = 131
const anon_sym_AND2r = 132
const anon_sym_ORA2r = 133
const anon_sym_EOR2r = 134
const anon_sym_SFT2r = 135
const anon_sym_LIT = 136
const anon_sym_INCk = 137
const anon_sym_POPk = 138
const anon_sym_NIPk = 139
const anon_sym_SWPk = 140
const anon_sym_ROTk = 141
const anon_sym_DUPk = 142
const anon_sym_OVRk = 143
const anon_sym_EQUk = 144
const anon_sym_NEQk = 145
const anon_sym_GTHk = 146
const anon_sym_LTHk = 147
const anon_sym_JMPk = 148
const anon_sym_JCNk = 149
const anon_sym_JSRk = 150
const anon_sym_STHk = 151
const anon_sym_LDZk = 152
const anon_sym_STZk = 153
const anon_sym_LDRk = 154
const anon_sym_STRk = 155
const anon_sym_LDAk = 156
const anon_sym_STAk = 157
const anon_sym_DEIk = 158
const anon_sym_DEOk = 159
const anon_sym_ADDk = 160
const anon_sym_SUBk = 161
const anon_sym_MULk = 162
const anon_sym_DIVk = 163
const anon_sym_ANDk = 164
const anon_sym_ORAk = 165
const anon_sym_EORk = 166
const anon_sym_SFTk = 167
const anon_sym_LIT2 = 168
const anon_sym_INC2k = 169
const anon_sym_POP2k = 170
const anon_sym_NIP2k = 171
const anon_sym_SWP2k = 172
const anon_sym_ROT2k = 173
const anon_sym_DUP2k = 174
const anon_sym_OVR2k = 175
const anon_sym_EQU2k = 176
const anon_sym_NEQ2k = 177
const anon_sym_GTH2k = 178
const anon_sym_LTH2k = 179
const anon_sym_JMP2k = 180
const anon_sym_JCN2k = 181
const anon_sym_JSR2k = 182
const anon_sym_STH2k = 183
const anon_sym_LDZ2k = 184
const anon_sym_STZ2k = 185
const anon_sym_LDR2k = 186
const anon_sym_STR2k = 187
const anon_sym_LDA2k = 188
const anon_sym_STA2k = 189
const anon_sym_DEI2k = 190
const anon_sym_DEO2k = 191
const anon_sym_ADD2k = 192
const anon_sym_SUB2k = 193
const anon_sym_MUL2k = 194
const anon_sym_DIV2k = 195
const anon_sym_AND2k = 196
const anon_sym_ORA2k = 197
const anon_sym_EOR2k = 198
const anon_sym_SFT2k = 199
const anon_sym_LITr = 200
const anon_sym_INCkr = 201
const anon_sym_POPkr = 202
const anon_sym_NIPkr = 203
const anon_sym_SWPkr = 204
const anon_sym_ROTkr = 205
const anon_sym_DUPkr = 206
const anon_sym_OVRkr = 207
const anon_sym_EQUkr = 208
const anon_sym_NEQkr = 209
const anon_sym_GTHkr = 210
const anon_sym_LTHkr = 211
const anon_sym_JMPkr = 212
const anon_sym_JCNkr = 213
const anon_sym_JSRkr = 214
const anon_sym_STHkr = 215
const anon_sym_LDZkr = 216
const anon_sym_STZkr = 217
const anon_sym_LDRkr = 218
const anon_sym_STRkr = 219
const anon_sym_LDAkr = 220
const anon_sym_STAkr = 221
const anon_sym_DEIkr = 222
const anon_sym_DEOkr = 223
const anon_sym_ADDkr = 224
const anon_sym_SUBkr = 225
const anon_sym_MULkr = 226
const anon_sym_DIVkr = 227
const anon_sym_ANDkr = 228
const anon_sym_ORAkr = 229
const anon_sym_EORkr = 230
const anon_sym_SFTkr = 231
const anon_sym_LIT2r = 232
const anon_sym_INC2kr = 233
const anon_sym_POP2kr = 234
const anon_sym_NIP2kr = 235
const anon_sym_SWP2kr = 236
const anon_sym_ROT2kr = 237
const anon_sym_DUP2kr = 238
const anon_sym_OVR2kr = 239
const anon_sym_EQU2kr = 240
const anon_sym_NEQ2kr = 241
const anon_sym_GTH2kr = 242
const anon_sym_LTH2kr = 243
const anon_sym_JMP2kr = 244
const anon_sym_JCN2kr = 245
const anon_sym_JSR2kr = 246
const anon_sym_STH2kr = 247
const anon_sym_LDZ2kr = 248
const anon_sym_STZ2kr = 249
const anon_sym_LDR2kr = 250
const anon_sym_STR2kr = 251
const anon_sym_LDA2kr = 252
const anon_sym_STA2kr = 253
const anon_sym_DEI2kr = 254
const anon_sym_DEO2kr = 255
const anon_sym_ADD2kr = 256
const anon_sym_SUB2kr = 257
const anon_sym_MUL2kr = 258
const anon_sym_DIV2kr = 259
const anon_sym_AND2kr = 260
const anon_sym_ORA2kr = 261
const anon_sym_EOR2kr = 262
const anon_sym_SFT2kr = 263
const anon_sym_PIPE = 264
const anon_sym_DOLLAR = 265
const anon_sym_POUND = 266
const aux_sym_hex_literal_token1 = 267
const anon_sym_AT = 268
const anon_sym_SLASH = 269
const aux_sym_rune_token1 = 270
const anon_sym_COMMA = 271
const anon_sym__ = 272
const anon_sym_DOT = 273
const anon_sym_DASH = 274
const anon_sym_SEMI = 275
const anon_sym_EQ = 276
const anon_sym_BANG = 277
const anon_sym_QMARK = 278
const anon_sym_AMP = 279
const anon_sym_LBRACK = 280
const anon_sym_RBRACK = 281
const anon_sym_DQUOTE = 282
const aux_sym_raw_ascii_token1 = 283
const sym_number = 284
const sym_comment = 285
const sym_program = 286
const sym_memory_execution = 287
const sym_subroutine = 288
const sym__non_toplevel_statement = 289
const sym_macro = 290
const sym_include = 291
const sym_opcode = 292
const sym_absolute_pad_operation = 293
const sym_relative_pad_operation = 294
const sym_hex_literal = 295
const sym_label = 296
const sym_sublabel_reference = 297
const sym_rune = 298
const sym_rune_char = 299
const sym_brackets = 300
const sym_raw_ascii = 301
const aux_sym_program_repeat1 = 302
const aux_sym_memory_execution_repeat1 = 303
const aux_sym_rune_repeat1 = 304

var ts_symbol_names = [305]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 15,
	3:   __ccgo_ts + 4,
	4:   __ccgo_ts + 17,
	5:   __ccgo_ts + 19,
	6:   __ccgo_ts + 21,
	7:   __ccgo_ts + 23,
	8:   __ccgo_ts + 38,
	9:   __ccgo_ts + 42,
	10:  __ccgo_ts + 46,
	11:  __ccgo_ts + 50,
	12:  __ccgo_ts + 54,
	13:  __ccgo_ts + 58,
	14:  __ccgo_ts + 62,
	15:  __ccgo_ts + 66,
	16:  __ccgo_ts + 70,
	17:  __ccgo_ts + 74,
	18:  __ccgo_ts + 78,
	19:  __ccgo_ts + 82,
	20:  __ccgo_ts + 86,
	21:  __ccgo_ts + 90,
	22:  __ccgo_ts + 94,
	23:  __ccgo_ts + 98,
	24:  __ccgo_ts + 102,
	25:  __ccgo_ts + 106,
	26:  __ccgo_ts + 110,
	27:  __ccgo_ts + 114,
	28:  __ccgo_ts + 118,
	29:  __ccgo_ts + 122,
	30:  __ccgo_ts + 126,
	31:  __ccgo_ts + 130,
	32:  __ccgo_ts + 134,
	33:  __ccgo_ts + 138,
	34:  __ccgo_ts + 142,
	35:  __ccgo_ts + 146,
	36:  __ccgo_ts + 150,
	37:  __ccgo_ts + 154,
	38:  __ccgo_ts + 158,
	39:  __ccgo_ts + 162,
	40:  __ccgo_ts + 166,
	41:  __ccgo_ts + 170,
	42:  __ccgo_ts + 175,
	43:  __ccgo_ts + 180,
	44:  __ccgo_ts + 185,
	45:  __ccgo_ts + 190,
	46:  __ccgo_ts + 195,
	47:  __ccgo_ts + 200,
	48:  __ccgo_ts + 205,
	49:  __ccgo_ts + 210,
	50:  __ccgo_ts + 215,
	51:  __ccgo_ts + 220,
	52:  __ccgo_ts + 225,
	53:  __ccgo_ts + 230,
	54:  __ccgo_ts + 235,
	55:  __ccgo_ts + 240,
	56:  __ccgo_ts + 245,
	57:  __ccgo_ts + 250,
	58:  __ccgo_ts + 255,
	59:  __ccgo_ts + 260,
	60:  __ccgo_ts + 265,
	61:  __ccgo_ts + 270,
	62:  __ccgo_ts + 275,
	63:  __ccgo_ts + 280,
	64:  __ccgo_ts + 285,
	65:  __ccgo_ts + 290,
	66:  __ccgo_ts + 295,
	67:  __ccgo_ts + 300,
	68:  __ccgo_ts + 305,
	69:  __ccgo_ts + 310,
	70:  __ccgo_ts + 315,
	71:  __ccgo_ts + 320,
	72:  __ccgo_ts + 325,
	73:  __ccgo_ts + 329,
	74:  __ccgo_ts + 334,
	75:  __ccgo_ts + 339,
	76:  __ccgo_ts + 344,
	77:  __ccgo_ts + 349,
	78:  __ccgo_ts + 354,
	79:  __ccgo_ts + 359,
	80:  __ccgo_ts + 364,
	81:  __ccgo_ts + 369,
	82:  __ccgo_ts + 374,
	83:  __ccgo_ts + 379,
	84:  __ccgo_ts + 384,
	85:  __ccgo_ts + 389,
	86:  __ccgo_ts + 394,
	87:  __ccgo_ts + 399,
	88:  __ccgo_ts + 404,
	89:  __ccgo_ts + 409,
	90:  __ccgo_ts + 414,
	91:  __ccgo_ts + 419,
	92:  __ccgo_ts + 424,
	93:  __ccgo_ts + 429,
	94:  __ccgo_ts + 434,
	95:  __ccgo_ts + 439,
	96:  __ccgo_ts + 444,
	97:  __ccgo_ts + 449,
	98:  __ccgo_ts + 454,
	99:  __ccgo_ts + 459,
	100: __ccgo_ts + 464,
	101: __ccgo_ts + 469,
	102: __ccgo_ts + 474,
	103: __ccgo_ts + 479,
	104: __ccgo_ts + 484,
	105: __ccgo_ts + 488,
	106: __ccgo_ts + 494,
	107: __ccgo_ts + 500,
	108: __ccgo_ts + 506,
	109: __ccgo_ts + 512,
	110: __ccgo_ts + 518,
	111: __ccgo_ts + 524,
	112: __ccgo_ts + 530,
	113: __ccgo_ts + 536,
	114: __ccgo_ts + 542,
	115: __ccgo_ts + 548,
	116: __ccgo_ts + 554,
	117: __ccgo_ts + 560,
	118: __ccgo_ts + 566,
	119: __ccgo_ts + 572,
	120: __ccgo_ts + 578,
	121: __ccgo_ts + 584,
	122: __ccgo_ts + 590,
	123: __ccgo_ts + 596,
	124: __ccgo_ts + 602,
	125: __ccgo_ts + 608,
	126: __ccgo_ts + 614,
	127: __ccgo_ts + 620,
	128: __ccgo_ts + 626,
	129: __ccgo_ts + 632,
	130: __ccgo_ts + 638,
	131: __ccgo_ts + 644,
	132: __ccgo_ts + 650,
	133: __ccgo_ts + 656,
	134: __ccgo_ts + 662,
	135: __ccgo_ts + 668,
	136: __ccgo_ts + 674,
	137: __ccgo_ts + 678,
	138: __ccgo_ts + 683,
	139: __ccgo_ts + 688,
	140: __ccgo_ts + 693,
	141: __ccgo_ts + 698,
	142: __ccgo_ts + 703,
	143: __ccgo_ts + 708,
	144: __ccgo_ts + 713,
	145: __ccgo_ts + 718,
	146: __ccgo_ts + 723,
	147: __ccgo_ts + 728,
	148: __ccgo_ts + 733,
	149: __ccgo_ts + 738,
	150: __ccgo_ts + 743,
	151: __ccgo_ts + 748,
	152: __ccgo_ts + 753,
	153: __ccgo_ts + 758,
	154: __ccgo_ts + 763,
	155: __ccgo_ts + 768,
	156: __ccgo_ts + 773,
	157: __ccgo_ts + 778,
	158: __ccgo_ts + 783,
	159: __ccgo_ts + 788,
	160: __ccgo_ts + 793,
	161: __ccgo_ts + 798,
	162: __ccgo_ts + 803,
	163: __ccgo_ts + 808,
	164: __ccgo_ts + 813,
	165: __ccgo_ts + 818,
	166: __ccgo_ts + 823,
	167: __ccgo_ts + 828,
	168: __ccgo_ts + 833,
	169: __ccgo_ts + 838,
	170: __ccgo_ts + 844,
	171: __ccgo_ts + 850,
	172: __ccgo_ts + 856,
	173: __ccgo_ts + 862,
	174: __ccgo_ts + 868,
	175: __ccgo_ts + 874,
	176: __ccgo_ts + 880,
	177: __ccgo_ts + 886,
	178: __ccgo_ts + 892,
	179: __ccgo_ts + 898,
	180: __ccgo_ts + 904,
	181: __ccgo_ts + 910,
	182: __ccgo_ts + 916,
	183: __ccgo_ts + 922,
	184: __ccgo_ts + 928,
	185: __ccgo_ts + 934,
	186: __ccgo_ts + 940,
	187: __ccgo_ts + 946,
	188: __ccgo_ts + 952,
	189: __ccgo_ts + 958,
	190: __ccgo_ts + 964,
	191: __ccgo_ts + 970,
	192: __ccgo_ts + 976,
	193: __ccgo_ts + 982,
	194: __ccgo_ts + 988,
	195: __ccgo_ts + 994,
	196: __ccgo_ts + 1000,
	197: __ccgo_ts + 1006,
	198: __ccgo_ts + 1012,
	199: __ccgo_ts + 1018,
	200: __ccgo_ts + 1024,
	201: __ccgo_ts + 1029,
	202: __ccgo_ts + 1035,
	203: __ccgo_ts + 1041,
	204: __ccgo_ts + 1047,
	205: __ccgo_ts + 1053,
	206: __ccgo_ts + 1059,
	207: __ccgo_ts + 1065,
	208: __ccgo_ts + 1071,
	209: __ccgo_ts + 1077,
	210: __ccgo_ts + 1083,
	211: __ccgo_ts + 1089,
	212: __ccgo_ts + 1095,
	213: __ccgo_ts + 1101,
	214: __ccgo_ts + 1107,
	215: __ccgo_ts + 1113,
	216: __ccgo_ts + 1119,
	217: __ccgo_ts + 1125,
	218: __ccgo_ts + 1131,
	219: __ccgo_ts + 1137,
	220: __ccgo_ts + 1143,
	221: __ccgo_ts + 1149,
	222: __ccgo_ts + 1155,
	223: __ccgo_ts + 1161,
	224: __ccgo_ts + 1167,
	225: __ccgo_ts + 1173,
	226: __ccgo_ts + 1179,
	227: __ccgo_ts + 1185,
	228: __ccgo_ts + 1191,
	229: __ccgo_ts + 1197,
	230: __ccgo_ts + 1203,
	231: __ccgo_ts + 1209,
	232: __ccgo_ts + 1215,
	233: __ccgo_ts + 1221,
	234: __ccgo_ts + 1228,
	235: __ccgo_ts + 1235,
	236: __ccgo_ts + 1242,
	237: __ccgo_ts + 1249,
	238: __ccgo_ts + 1256,
	239: __ccgo_ts + 1263,
	240: __ccgo_ts + 1270,
	241: __ccgo_ts + 1277,
	242: __ccgo_ts + 1284,
	243: __ccgo_ts + 1291,
	244: __ccgo_ts + 1298,
	245: __ccgo_ts + 1305,
	246: __ccgo_ts + 1312,
	247: __ccgo_ts + 1319,
	248: __ccgo_ts + 1326,
	249: __ccgo_ts + 1333,
	250: __ccgo_ts + 1340,
	251: __ccgo_ts + 1347,
	252: __ccgo_ts + 1354,
	253: __ccgo_ts + 1361,
	254: __ccgo_ts + 1368,
	255: __ccgo_ts + 1375,
	256: __ccgo_ts + 1382,
	257: __ccgo_ts + 1389,
	258: __ccgo_ts + 1396,
	259: __ccgo_ts + 1403,
	260: __ccgo_ts + 1410,
	261: __ccgo_ts + 1417,
	262: __ccgo_ts + 1424,
	263: __ccgo_ts + 1431,
	264: __ccgo_ts + 1438,
	265: __ccgo_ts + 1440,
	266: __ccgo_ts + 1442,
	267: __ccgo_ts + 1444,
	268: __ccgo_ts + 1458,
	269: __ccgo_ts + 1460,
	270: __ccgo_ts + 4,
	271: __ccgo_ts + 1462,
	272: __ccgo_ts + 1464,
	273: __ccgo_ts + 1466,
	274: __ccgo_ts + 1468,
	275: __ccgo_ts + 1470,
	276: __ccgo_ts + 1472,
	277: __ccgo_ts + 1474,
	278: __ccgo_ts + 1476,
	279: __ccgo_ts + 1478,
	280: __ccgo_ts + 1480,
	281: __ccgo_ts + 1482,
	282: __ccgo_ts + 1484,
	283: __ccgo_ts + 1486,
	284: __ccgo_ts + 1503,
	285: __ccgo_ts + 1510,
	286: __ccgo_ts + 1518,
	287: __ccgo_ts + 1526,
	288: __ccgo_ts + 1543,
	289: __ccgo_ts + 1554,
	290: __ccgo_ts + 1578,
	291: __ccgo_ts + 1584,
	292: __ccgo_ts + 1592,
	293: __ccgo_ts + 1599,
	294: __ccgo_ts + 1622,
	295: __ccgo_ts + 1645,
	296: __ccgo_ts + 1657,
	297: __ccgo_ts + 1663,
	298: __ccgo_ts + 1682,
	299: __ccgo_ts + 1687,
	300: __ccgo_ts + 1697,
	301: __ccgo_ts + 1706,
	302: __ccgo_ts + 1716,
	303: __ccgo_ts + 1732,
	304: __ccgo_ts + 1757,
}

var ts_symbol_map = [305]TSSymbol{
	1:   uint16(sym_identifier),
	2:   uint16(anon_sym_PERCENT),
	3:   uint16(sym_identifier),
	4:   uint16(anon_sym_LBRACE),
	5:   uint16(anon_sym_RBRACE),
	6:   uint16(anon_sym_TILDE),
	7:   uint16(aux_sym_include_token1),
	8:   uint16(anon_sym_BRK),
	9:   uint16(anon_sym_INC),
	10:  uint16(anon_sym_POP),
	11:  uint16(anon_sym_NIP),
	12:  uint16(anon_sym_SWP),
	13:  uint16(anon_sym_ROT),
	14:  uint16(anon_sym_DUP),
	15:  uint16(anon_sym_OVR),
	16:  uint16(anon_sym_EQU),
	17:  uint16(anon_sym_NEQ),
	18:  uint16(anon_sym_GTH),
	19:  uint16(anon_sym_LTH),
	20:  uint16(anon_sym_JMP),
	21:  uint16(anon_sym_JCN),
	22:  uint16(anon_sym_JSR),
	23:  uint16(anon_sym_STH),
	24:  uint16(anon_sym_LDZ),
	25:  uint16(anon_sym_STZ),
	26:  uint16(anon_sym_LDR),
	27:  uint16(anon_sym_STR),
	28:  uint16(anon_sym_LDA),
	29:  uint16(anon_sym_STA),
	30:  uint16(anon_sym_DEI),
	31:  uint16(anon_sym_DEO),
	32:  uint16(anon_sym_ADD),
	33:  uint16(anon_sym_SUB),
	34:  uint16(anon_sym_MUL),
	35:  uint16(anon_sym_DIV),
	36:  uint16(anon_sym_AND),
	37:  uint16(anon_sym_ORA),
	38:  uint16(anon_sym_EOR),
	39:  uint16(anon_sym_SFT),
	40:  uint16(anon_sym_JCI),
	41:  uint16(anon_sym_INC2),
	42:  uint16(anon_sym_POP2),
	43:  uint16(anon_sym_NIP2),
	44:  uint16(anon_sym_SWP2),
	45:  uint16(anon_sym_ROT2),
	46:  uint16(anon_sym_DUP2),
	47:  uint16(anon_sym_OVR2),
	48:  uint16(anon_sym_EQU2),
	49:  uint16(anon_sym_NEQ2),
	50:  uint16(anon_sym_GTH2),
	51:  uint16(anon_sym_LTH2),
	52:  uint16(anon_sym_JMP2),
	53:  uint16(anon_sym_JCN2),
	54:  uint16(anon_sym_JSR2),
	55:  uint16(anon_sym_STH2),
	56:  uint16(anon_sym_LDZ2),
	57:  uint16(anon_sym_STZ2),
	58:  uint16(anon_sym_LDR2),
	59:  uint16(anon_sym_STR2),
	60:  uint16(anon_sym_LDA2),
	61:  uint16(anon_sym_STA2),
	62:  uint16(anon_sym_DEI2),
	63:  uint16(anon_sym_DEO2),
	64:  uint16(anon_sym_ADD2),
	65:  uint16(anon_sym_SUB2),
	66:  uint16(anon_sym_MUL2),
	67:  uint16(anon_sym_DIV2),
	68:  uint16(anon_sym_AND2),
	69:  uint16(anon_sym_ORA2),
	70:  uint16(anon_sym_EOR2),
	71:  uint16(anon_sym_SFT2),
	72:  uint16(anon_sym_JMI),
	73:  uint16(anon_sym_INCr),
	74:  uint16(anon_sym_POPr),
	75:  uint16(anon_sym_NIPr),
	76:  uint16(anon_sym_SWPr),
	77:  uint16(anon_sym_ROTr),
	78:  uint16(anon_sym_DUPr),
	79:  uint16(anon_sym_OVRr),
	80:  uint16(anon_sym_EQUr),
	81:  uint16(anon_sym_NEQr),
	82:  uint16(anon_sym_GTHr),
	83:  uint16(anon_sym_LTHr),
	84:  uint16(anon_sym_JMPr),
	85:  uint16(anon_sym_JCNr),
	86:  uint16(anon_sym_JSRr),
	87:  uint16(anon_sym_STHr),
	88:  uint16(anon_sym_LDZr),
	89:  uint16(anon_sym_STZr),
	90:  uint16(anon_sym_LDRr),
	91:  uint16(anon_sym_STRr),
	92:  uint16(anon_sym_LDAr),
	93:  uint16(anon_sym_STAr),
	94:  uint16(anon_sym_DEIr),
	95:  uint16(anon_sym_DEOr),
	96:  uint16(anon_sym_ADDr),
	97:  uint16(anon_sym_SUBr),
	98:  uint16(anon_sym_MULr),
	99:  uint16(anon_sym_DIVr),
	100: uint16(anon_sym_ANDr),
	101: uint16(anon_sym_ORAr),
	102: uint16(anon_sym_EORr),
	103: uint16(anon_sym_SFTr),
	104: uint16(anon_sym_JSI),
	105: uint16(anon_sym_INC2r),
	106: uint16(anon_sym_POP2r),
	107: uint16(anon_sym_NIP2r),
	108: uint16(anon_sym_SWP2r),
	109: uint16(anon_sym_ROT2r),
	110: uint16(anon_sym_DUP2r),
	111: uint16(anon_sym_OVR2r),
	112: uint16(anon_sym_EQU2r),
	113: uint16(anon_sym_NEQ2r),
	114: uint16(anon_sym_GTH2r),
	115: uint16(anon_sym_LTH2r),
	116: uint16(anon_sym_JMP2r),
	117: uint16(anon_sym_JCN2r),
	118: uint16(anon_sym_JSR2r),
	119: uint16(anon_sym_STH2r),
	120: uint16(anon_sym_LDZ2r),
	121: uint16(anon_sym_STZ2r),
	122: uint16(anon_sym_LDR2r),
	123: uint16(anon_sym_STR2r),
	124: uint16(anon_sym_LDA2r),
	125: uint16(anon_sym_STA2r),
	126: uint16(anon_sym_DEI2r),
	127: uint16(anon_sym_DEO2r),
	128: uint16(anon_sym_ADD2r),
	129: uint16(anon_sym_SUB2r),
	130: uint16(anon_sym_MUL2r),
	131: uint16(anon_sym_DIV2r),
	132: uint16(anon_sym_AND2r),
	133: uint16(anon_sym_ORA2r),
	134: uint16(anon_sym_EOR2r),
	135: uint16(anon_sym_SFT2r),
	136: uint16(anon_sym_LIT),
	137: uint16(anon_sym_INCk),
	138: uint16(anon_sym_POPk),
	139: uint16(anon_sym_NIPk),
	140: uint16(anon_sym_SWPk),
	141: uint16(anon_sym_ROTk),
	142: uint16(anon_sym_DUPk),
	143: uint16(anon_sym_OVRk),
	144: uint16(anon_sym_EQUk),
	145: uint16(anon_sym_NEQk),
	146: uint16(anon_sym_GTHk),
	147: uint16(anon_sym_LTHk),
	148: uint16(anon_sym_JMPk),
	149: uint16(anon_sym_JCNk),
	150: uint16(anon_sym_JSRk),
	151: uint16(anon_sym_STHk),
	152: uint16(anon_sym_LDZk),
	153: uint16(anon_sym_STZk),
	154: uint16(anon_sym_LDRk),
	155: uint16(anon_sym_STRk),
	156: uint16(anon_sym_LDAk),
	157: uint16(anon_sym_STAk),
	158: uint16(anon_sym_DEIk),
	159: uint16(anon_sym_DEOk),
	160: uint16(anon_sym_ADDk),
	161: uint16(anon_sym_SUBk),
	162: uint16(anon_sym_MULk),
	163: uint16(anon_sym_DIVk),
	164: uint16(anon_sym_ANDk),
	165: uint16(anon_sym_ORAk),
	166: uint16(anon_sym_EORk),
	167: uint16(anon_sym_SFTk),
	168: uint16(anon_sym_LIT2),
	169: uint16(anon_sym_INC2k),
	170: uint16(anon_sym_POP2k),
	171: uint16(anon_sym_NIP2k),
	172: uint16(anon_sym_SWP2k),
	173: uint16(anon_sym_ROT2k),
	174: uint16(anon_sym_DUP2k),
	175: uint16(anon_sym_OVR2k),
	176: uint16(anon_sym_EQU2k),
	177: uint16(anon_sym_NEQ2k),
	178: uint16(anon_sym_GTH2k),
	179: uint16(anon_sym_LTH2k),
	180: uint16(anon_sym_JMP2k),
	181: uint16(anon_sym_JCN2k),
	182: uint16(anon_sym_JSR2k),
	183: uint16(anon_sym_STH2k),
	184: uint16(anon_sym_LDZ2k),
	185: uint16(anon_sym_STZ2k),
	186: uint16(anon_sym_LDR2k),
	187: uint16(anon_sym_STR2k),
	188: uint16(anon_sym_LDA2k),
	189: uint16(anon_sym_STA2k),
	190: uint16(anon_sym_DEI2k),
	191: uint16(anon_sym_DEO2k),
	192: uint16(anon_sym_ADD2k),
	193: uint16(anon_sym_SUB2k),
	194: uint16(anon_sym_MUL2k),
	195: uint16(anon_sym_DIV2k),
	196: uint16(anon_sym_AND2k),
	197: uint16(anon_sym_ORA2k),
	198: uint16(anon_sym_EOR2k),
	199: uint16(anon_sym_SFT2k),
	200: uint16(anon_sym_LITr),
	201: uint16(anon_sym_INCkr),
	202: uint16(anon_sym_POPkr),
	203: uint16(anon_sym_NIPkr),
	204: uint16(anon_sym_SWPkr),
	205: uint16(anon_sym_ROTkr),
	206: uint16(anon_sym_DUPkr),
	207: uint16(anon_sym_OVRkr),
	208: uint16(anon_sym_EQUkr),
	209: uint16(anon_sym_NEQkr),
	210: uint16(anon_sym_GTHkr),
	211: uint16(anon_sym_LTHkr),
	212: uint16(anon_sym_JMPkr),
	213: uint16(anon_sym_JCNkr),
	214: uint16(anon_sym_JSRkr),
	215: uint16(anon_sym_STHkr),
	216: uint16(anon_sym_LDZkr),
	217: uint16(anon_sym_STZkr),
	218: uint16(anon_sym_LDRkr),
	219: uint16(anon_sym_STRkr),
	220: uint16(anon_sym_LDAkr),
	221: uint16(anon_sym_STAkr),
	222: uint16(anon_sym_DEIkr),
	223: uint16(anon_sym_DEOkr),
	224: uint16(anon_sym_ADDkr),
	225: uint16(anon_sym_SUBkr),
	226: uint16(anon_sym_MULkr),
	227: uint16(anon_sym_DIVkr),
	228: uint16(anon_sym_ANDkr),
	229: uint16(anon_sym_ORAkr),
	230: uint16(anon_sym_EORkr),
	231: uint16(anon_sym_SFTkr),
	232: uint16(anon_sym_LIT2r),
	233: uint16(anon_sym_INC2kr),
	234: uint16(anon_sym_POP2kr),
	235: uint16(anon_sym_NIP2kr),
	236: uint16(anon_sym_SWP2kr),
	237: uint16(anon_sym_ROT2kr),
	238: uint16(anon_sym_DUP2kr),
	239: uint16(anon_sym_OVR2kr),
	240: uint16(anon_sym_EQU2kr),
	241: uint16(anon_sym_NEQ2kr),
	242: uint16(anon_sym_GTH2kr),
	243: uint16(anon_sym_LTH2kr),
	244: uint16(anon_sym_JMP2kr),
	245: uint16(anon_sym_JCN2kr),
	246: uint16(anon_sym_JSR2kr),
	247: uint16(anon_sym_STH2kr),
	248: uint16(anon_sym_LDZ2kr),
	249: uint16(anon_sym_STZ2kr),
	250: uint16(anon_sym_LDR2kr),
	251: uint16(anon_sym_STR2kr),
	252: uint16(anon_sym_LDA2kr),
	253: uint16(anon_sym_STA2kr),
	254: uint16(anon_sym_DEI2kr),
	255: uint16(anon_sym_DEO2kr),
	256: uint16(anon_sym_ADD2kr),
	257: uint16(anon_sym_SUB2kr),
	258: uint16(anon_sym_MUL2kr),
	259: uint16(anon_sym_DIV2kr),
	260: uint16(anon_sym_AND2kr),
	261: uint16(anon_sym_ORA2kr),
	262: uint16(anon_sym_EOR2kr),
	263: uint16(anon_sym_SFT2kr),
	264: uint16(anon_sym_PIPE),
	265: uint16(anon_sym_DOLLAR),
	266: uint16(anon_sym_POUND),
	267: uint16(aux_sym_hex_literal_token1),
	268: uint16(anon_sym_AT),
	269: uint16(anon_sym_SLASH),
	270: uint16(sym_identifier),
	271: uint16(anon_sym_COMMA),
	272: uint16(anon_sym__),
	273: uint16(anon_sym_DOT),
	274: uint16(anon_sym_DASH),
	275: uint16(anon_sym_SEMI),
	276: uint16(anon_sym_EQ),
	277: uint16(anon_sym_BANG),
	278: uint16(anon_sym_QMARK),
	279: uint16(anon_sym_AMP),
	280: uint16(anon_sym_LBRACK),
	281: uint16(anon_sym_RBRACK),
	282: uint16(anon_sym_DQUOTE),
	283: uint16(aux_sym_raw_ascii_token1),
	284: uint16(sym_number),
	285: uint16(sym_comment),
	286: uint16(sym_program),
	287: uint16(sym_memory_execution),
	288: uint16(sym_subroutine),
	289: uint16(sym__non_toplevel_statement),
	290: uint16(sym_macro),
	291: uint16(sym_include),
	292: uint16(sym_opcode),
	293: uint16(sym_absolute_pad_operation),
	294: uint16(sym_relative_pad_operation),
	295: uint16(sym_hex_literal),
	296: uint16(sym_label),
	297: uint16(sym_sublabel_reference),
	298: uint16(sym_rune),
	299: uint16(sym_rune_char),
	300: uint16(sym_brackets),
	301: uint16(sym_raw_ascii),
	302: uint16(aux_sym_program_repeat1),
	303: uint16(aux_sym_memory_execution_repeat1),
	304: uint16(aux_sym_rune_repeat1),
}

var ts_symbol_metadata = [305]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	},
	136: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	137: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	138: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	139: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	140: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	141: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	142: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	143: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	144: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	145: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	146: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	147: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	148: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	149: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	150: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	151: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	152: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	153: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	154: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	155: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	156: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	157: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	158: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	159: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	160: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	161: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	162: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	163: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	164: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	165: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	166: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	167: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	168: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	169: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	170: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	171: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	172: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	173: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	174: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	175: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	176: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	177: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	178: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	179: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	180: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	181: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	182: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	183: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	184: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	185: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	186: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	187: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	188: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	189: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	190: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	191: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	192: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	193: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	194: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	195: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	196: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	197: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	198: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	199: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	200: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	201: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	202: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	203: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	204: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	205: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	206: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	207: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	208: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	209: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	210: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	211: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	212: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	213: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	214: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	215: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	216: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	217: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	218: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	219: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	220: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	221: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	222: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	223: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	224: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	225: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	226: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	227: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	228: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	229: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	230: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	231: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	232: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	233: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	234: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	235: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	236: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	237: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	238: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	239: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	240: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	241: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	242: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	243: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	244: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	245: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	246: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	247: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	248: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	249: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	250: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	251: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	252: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	253: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	254: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	255: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	256: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	257: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	258: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	259: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	260: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	261: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	262: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	263: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	264: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	265: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	266: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	267: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	268: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	269: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	270: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	271: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	272: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	273: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	274: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	275: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	276: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	277: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	278: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	279: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	280: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	281: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	282: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	283: {},
	284: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	285: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	286: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	287: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	288: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	289: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	290: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	291: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	292: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	293: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	294: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	295: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	296: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	297: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	298: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	299: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	300: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	301: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	302: {},
	303: {},
	304: {},
}

type ts_field_identifiers = int32

const field_absolute = 1
const field_immediate = 2
const field_relative = 3
const field_rune_start = 4
const field_sublabel = 5
const field_zero_page = 6

var ts_field_names = [7]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 1770,
	2: __ccgo_ts + 1779,
	3: __ccgo_ts + 1789,
	4: __ccgo_ts + 1798,
	5: __ccgo_ts + 1809,
	6: __ccgo_ts + 1818,
}

var ts_field_map_slices = [7]TSFieldMapSlice{
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
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(4),
		Flength: uint16(1),
	},
	6: {
		Findex:  uint16(5),
		Flength: uint16(1),
	},
}

var ts_field_map_entries = [6]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_relative),
	},
	1: {
		Ffield_id: uint16(field_zero_page),
	},
	2: {
		Ffield_id: uint16(field_absolute),
	},
	3: {
		Ffield_id: uint16(field_immediate),
	},
	4: {
		Ffield_id: uint16(field_sublabel),
	},
	5: {
		Ffield_id: uint16(field_rune_start),
	},
}

var ts_alias_sequences = [7][5]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [48]TSStateId{
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
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _ = eof, i, i1, i2, lookahead, result, skip
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
			state = uint16(15)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(i < libc.Uint32FromInt64(80)/libc.Uint32FromInt64(2)) {
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
			state = uint16(0)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(53)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('*') || lookahead == int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(1):
		i1 = uint32(0)
		for {
			if !(i1 < libc.Uint32FromInt64(40)/libc.Uint32FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(2):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(3):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(52)
			goto next_state
		}
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(54)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('*') || int32('/') <= lookahead && lookahead <= int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(4):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(52)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(2)
			goto next_state
		}
		return result
	case int32(5):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(52)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(4)
			goto next_state
		}
		return result
	case int32(6):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('*') || lookahead == int32('/') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(17)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(7):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(8):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('*') || int32('/') <= lookahead && lookahead <= int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(9):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('.') || lookahead == int32('/') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(10):
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(11):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(12):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('*') || lookahead == int32('/') || lookahead == int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(14):
		if eof != 0 {
			state = uint16(15)
			goto next_state
		}
		i2 = uint32(0)
		for {
			if !(i2 < libc.Uint32FromInt64(72)/libc.Uint32FromInt64(2)) {
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
		if int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(53)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('*') || int32('/') <= lookahead && lookahead <= int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(15):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(16):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(17):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_macro_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('*') || lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(18):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_macro_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('*') || lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(19):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_macro_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('*') || lookahead == int32('/') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(20):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_macro_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') || lookahead == int32('-') || int32('/') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(22):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(23):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TILDE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(24):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_include_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('-') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_literal_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_hex_literal_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_rune_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_rune_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('*') || int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_rune_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('*') || int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_rune_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('*') || int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_rune_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('*') || int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_rune_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') || lookahead == int32('/') {
			state = uint16(35)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_rune_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') || int32('/') <= lookahead && lookahead <= int32('9') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_rune_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') || int32('/') <= lookahead && lookahead <= int32('9') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_raw_ascii_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32(' ') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(55)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(52)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('#') || lookahead == int32('*') || lookahead == int32('-') || lookahead == int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(55)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(52)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('#') || lookahead == int32('*') || lookahead == int32('-') || lookahead == int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('#') || lookahead == int32('*') || lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(52)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('#') || lookahead == int32('*') || lookahead == int32('-') || lookahead == int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(52)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('#') || lookahead == int32('*') || lookahead == int32('-') || lookahead == int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('#') || lookahead == int32('*') || lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') || lookahead == int32('*') || lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32(':') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(59)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [40]uint16_t{
	0:  uint16('!'),
	1:  uint16(45),
	2:  uint16('"'),
	3:  uint16(50),
	4:  uint16('#'),
	5:  uint16(27),
	6:  uint16('$'),
	7:  uint16(26),
	8:  uint16('%'),
	9:  uint16(16),
	10: uint16('&'),
	11: uint16(47),
	12: uint16(','),
	13: uint16(40),
	14: uint16('-'),
	15: uint16(42),
	16: uint16('.'),
	17: uint16(41),
	18: uint16('/'),
	19: uint16(31),
	20: uint16(';'),
	21: uint16(43),
	22: uint16('='),
	23: uint16(44),
	24: uint16('?'),
	25: uint16(46),
	26: uint16('@'),
	27: uint16(30),
	28: uint16('['),
	29: uint16(48),
	30: uint16(']'),
	31: uint16(49),
	32: uint16('{'),
	33: uint16(21),
	34: uint16('|'),
	35: uint16(25),
	36: uint16('}'),
	37: uint16(22),
	38: uint16('~'),
	39: uint16(23),
}

var map_token1 = [20]uint16_t{
	0:  uint16('!'),
	1:  uint16(45),
	2:  uint16('&'),
	3:  uint16(47),
	4:  uint16(','),
	5:  uint16(40),
	6:  uint16('-'),
	7:  uint16(42),
	8:  uint16('.'),
	9:  uint16(41),
	10: uint16(';'),
	11: uint16(43),
	12: uint16('='),
	13: uint16(44),
	14: uint16('?'),
	15: uint16(46),
	16: uint16('*'),
	17: uint16(36),
	18: uint16('/'),
	19: uint16(36),
}

var map_token2 = [36]uint16_t{
	0:  uint16('!'),
	1:  uint16(45),
	2:  uint16('"'),
	3:  uint16(50),
	4:  uint16('#'),
	5:  uint16(27),
	6:  uint16('$'),
	7:  uint16(26),
	8:  uint16('%'),
	9:  uint16(16),
	10: uint16('&'),
	11: uint16(47),
	12: uint16(','),
	13: uint16(40),
	14: uint16('-'),
	15: uint16(42),
	16: uint16('.'),
	17: uint16(41),
	18: uint16(';'),
	19: uint16(43),
	20: uint16('='),
	21: uint16(44),
	22: uint16('?'),
	23: uint16(46),
	24: uint16('@'),
	25: uint16(30),
	26: uint16('['),
	27: uint16(48),
	28: uint16(']'),
	29: uint16(49),
	30: uint16('|'),
	31: uint16(25),
	32: uint16('}'),
	33: uint16(22),
	34: uint16('~'),
	35: uint16(23),
}

func ts_lex_keywords(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i uint32_t
	var lookahead int32_t
	_, _, _, _, _ = eof, i, lookahead, result, skip
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
		i = uint32(0)
		for {
			if !(i < libc.Uint32FromInt64(60)/libc.Uint32FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token3[i]) == lookahead {
				state = map_token3[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('D') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('N') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('R') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('E') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('I') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('U') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('O') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('Q') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('T') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('N') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('C') {
			state = uint16(26)
			goto next_state
		}
		if lookahead == int32('M') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('S') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('D') {
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('I') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('U') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('E') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('I') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('R') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('V') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('O') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('O') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('F') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('U') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('W') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(15):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym__)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(16):
		if lookahead == int32('D') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('D') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('K') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('I') {
			state = uint16(46)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('V') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('P') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('R') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('U') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('H') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('C') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('I') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('N') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('I') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('P') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('I') {
			state = uint16(58)
			goto next_state
		}
		if lookahead == int32('R') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('A') {
			state = uint16(60)
			goto next_state
		}
		if lookahead == int32('R') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('Z') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('T') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('H') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('L') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('Q') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('P') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('A') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('R') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('P') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('T') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('T') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('A') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('H') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('R') {
			state = uint16(75)
			goto next_state
		}
		if lookahead == int32('Z') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('B') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('P') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ADD)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(83)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BRK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEO)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DIV)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(91)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(92)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DUP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(94)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(95)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EOR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(98)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQU)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(100)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(101)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GTH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INC)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(106)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JCI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JCN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(109)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(110)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(112)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(113)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JSI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JSR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(118)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(119)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(122)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDZ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(124)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(125)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LIT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(127)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LTH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(129)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(130)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MUL)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(132)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(133)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NEQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(136)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NIP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(138)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ORA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_OVR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(145)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(147)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ROT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(150)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(151)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SFT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(153)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(154)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(156)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(159)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(160)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(163)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STZ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(165)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(166)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SUB)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(168)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SWP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('2') {
			state = uint16(171)
			goto next_state
		}
		if lookahead == int32('k') {
			state = uint16(172)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ADD2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(174)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ADDk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ADDr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AND2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ANDk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ANDr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEI2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(180)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEIk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(182)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEIr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEO2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(183)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEOk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEOr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DIV2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(186)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DIVk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DIVr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DUP2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(189)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DUPk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DUPr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EOR2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(192)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EORk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EORr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQU2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(195)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQUk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQUr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GTH2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(198)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GTHk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GTHr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INC2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(201)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INCk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INCr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JCN2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(204)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JCNk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JCNr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JMP2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(207)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JMPk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JMPr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JSR2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(210)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JSRk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(212)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JSRr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDA2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(213)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDAk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(215)
			goto next_state
		}
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDAr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDR2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(216)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(217)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDRk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(218)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDRr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDZ2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(219)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(220)
			goto next_state
		}
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDZk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(221)
			goto next_state
		}
		return result
	case int32(126):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDZr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LIT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(222)
			goto next_state
		}
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LITr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LTH2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(223)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(224)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LTHk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(225)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LTHr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(132):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MUL2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(226)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(227)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MULk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(228)
			goto next_state
		}
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MULr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NEQ2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(229)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(230)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NEQk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(231)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NEQr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(138):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NIP2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(232)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(233)
			goto next_state
		}
		return result
	case int32(139):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NIPk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(234)
			goto next_state
		}
		return result
	case int32(140):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NIPr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(141):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ORA2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(235)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(236)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ORAk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(237)
			goto next_state
		}
		return result
	case int32(143):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ORAr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(144):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_OVR2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(238)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(239)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_OVRk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(240)
			goto next_state
		}
		return result
	case int32(146):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_OVRr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(147):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POP2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(241)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(242)
			goto next_state
		}
		return result
	case int32(148):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POPk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(243)
			goto next_state
		}
		return result
	case int32(149):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POPr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(150):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ROT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(244)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(245)
			goto next_state
		}
		return result
	case int32(151):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ROTk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(246)
			goto next_state
		}
		return result
	case int32(152):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ROTr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(153):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SFT2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(247)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(248)
			goto next_state
		}
		return result
	case int32(154):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SFTk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(249)
			goto next_state
		}
		return result
	case int32(155):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SFTr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(156):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STA2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(250)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(251)
			goto next_state
		}
		return result
	case int32(157):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(252)
			goto next_state
		}
		return result
	case int32(158):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(159):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STH2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(253)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(254)
			goto next_state
		}
		return result
	case int32(160):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STHk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(255)
			goto next_state
		}
		return result
	case int32(161):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STHr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(162):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STR2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(256)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(257)
			goto next_state
		}
		return result
	case int32(163):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STRk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(258)
			goto next_state
		}
		return result
	case int32(164):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STRr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(165):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STZ2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(259)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(260)
			goto next_state
		}
		return result
	case int32(166):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STZk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(167):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STZr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(168):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SUB2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(262)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(263)
			goto next_state
		}
		return result
	case int32(169):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SUBk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(264)
			goto next_state
		}
		return result
	case int32(170):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SUBr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(171):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SWP2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('k') {
			state = uint16(265)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(266)
			goto next_state
		}
		return result
	case int32(172):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SWPk)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(267)
			goto next_state
		}
		return result
	case int32(173):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SWPr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(174):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ADD2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(175):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ADD2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(176):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ADDkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(177):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AND2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(178):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AND2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(179):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ANDkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(180):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEI2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(181):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEI2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(182):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEIkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(183):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEO2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(184):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEO2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(185):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEOkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(186):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DIV2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(272)
			goto next_state
		}
		return result
	case int32(187):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DIV2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(188):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DIVkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(189):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DUP2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(190):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DUP2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(191):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DUPkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(192):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EOR2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(274)
			goto next_state
		}
		return result
	case int32(193):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EOR2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(194):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EORkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(195):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQU2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(275)
			goto next_state
		}
		return result
	case int32(196):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQU2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(197):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQUkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(198):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GTH2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(276)
			goto next_state
		}
		return result
	case int32(199):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GTH2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(200):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GTHkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(201):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INC2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(277)
			goto next_state
		}
		return result
	case int32(202):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INC2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(203):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INCkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(204):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JCN2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(278)
			goto next_state
		}
		return result
	case int32(205):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JCN2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(206):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JCNkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(207):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JMP2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(279)
			goto next_state
		}
		return result
	case int32(208):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JMP2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(209):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JMPkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(210):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JSR2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(280)
			goto next_state
		}
		return result
	case int32(211):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JSR2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(212):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JSRkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(213):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDA2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(281)
			goto next_state
		}
		return result
	case int32(214):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDA2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(215):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDAkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(216):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDR2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(282)
			goto next_state
		}
		return result
	case int32(217):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDR2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(218):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDRkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(219):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDZ2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(283)
			goto next_state
		}
		return result
	case int32(220):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDZ2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(221):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDZkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(222):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LIT2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(223):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LTH2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(284)
			goto next_state
		}
		return result
	case int32(224):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LTH2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(225):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LTHkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(226):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MUL2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(285)
			goto next_state
		}
		return result
	case int32(227):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MUL2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(228):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MULkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(229):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NEQ2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(286)
			goto next_state
		}
		return result
	case int32(230):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NEQ2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(231):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NEQkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(232):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NIP2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(287)
			goto next_state
		}
		return result
	case int32(233):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NIP2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(234):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NIPkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(235):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ORA2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(288)
			goto next_state
		}
		return result
	case int32(236):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ORA2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(237):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ORAkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(238):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_OVR2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(289)
			goto next_state
		}
		return result
	case int32(239):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_OVR2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(240):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_OVRkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(241):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POP2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(290)
			goto next_state
		}
		return result
	case int32(242):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POP2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(243):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POPkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(244):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ROT2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(291)
			goto next_state
		}
		return result
	case int32(245):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ROT2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(246):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ROTkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(247):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SFT2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(292)
			goto next_state
		}
		return result
	case int32(248):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SFT2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(249):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SFTkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(250):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STA2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(293)
			goto next_state
		}
		return result
	case int32(251):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STA2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(252):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(253):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STH2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(294)
			goto next_state
		}
		return result
	case int32(254):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STH2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(255):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STHkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(256):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STR2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(295)
			goto next_state
		}
		return result
	case int32(257):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STR2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(258):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STRkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(259):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STZ2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(260):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STZ2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(261):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STZkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(262):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SUB2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(297)
			goto next_state
		}
		return result
	case int32(263):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SUB2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(264):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SUBkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(265):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SWP2k)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(298)
			goto next_state
		}
		return result
	case int32(266):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SWP2r)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(267):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SWPkr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(268):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ADD2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(269):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AND2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(270):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEI2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(271):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DEO2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(272):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DIV2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(273):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DUP2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(274):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EOR2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(275):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQU2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(276):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GTH2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(277):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_INC2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(278):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JCN2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(279):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JMP2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(280):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_JSR2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(281):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDA2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(282):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDR2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(283):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LDZ2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(284):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LTH2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(285):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MUL2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(286):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NEQ2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(287):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NIP2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(288):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ORA2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(289):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_OVR2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(290):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POP2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(291):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ROT2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(292):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SFT2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(293):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STA2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(294):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STH2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(295):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STR2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(296):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STZ2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(297):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SUB2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(298):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SWP2kr)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token3 = [30]uint16_t{
	0:  uint16('A'),
	1:  uint16(1),
	2:  uint16('B'),
	3:  uint16(2),
	4:  uint16('D'),
	5:  uint16(3),
	6:  uint16('E'),
	7:  uint16(4),
	8:  uint16('G'),
	9:  uint16(5),
	10: uint16('I'),
	11: uint16(6),
	12: uint16('J'),
	13: uint16(7),
	14: uint16('L'),
	15: uint16(8),
	16: uint16('M'),
	17: uint16(9),
	18: uint16('N'),
	19: uint16(10),
	20: uint16('O'),
	21: uint16(11),
	22: uint16('P'),
	23: uint16(12),
	24: uint16('R'),
	25: uint16(13),
	26: uint16('S'),
	27: uint16(14),
	28: uint16('_'),
	29: uint16(15),
}

var ts_lex_modes = [48]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Fexternal_lex_state: uint16(1),
	},
	2: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	3: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	4: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	5: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	6: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	7: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	8: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	9: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	10: {
		Flex_state:          uint16(14),
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
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	15: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	16: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	17: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	18: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	19: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	20: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	21: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	22: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	23: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	24: {
		Flex_state:          uint16(14),
		Fexternal_lex_state: uint16(1),
	},
	25: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	26: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	27: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	28: {
		Fexternal_lex_state: uint16(1),
	},
	29: {
		Fexternal_lex_state: uint16(1),
	},
	30: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	31: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	32: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	33: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	34: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(1),
	},
	35: {
		Fexternal_lex_state: uint16(1),
	},
	36: {
		Fexternal_lex_state: uint16(1),
	},
	37: {
		Fexternal_lex_state: uint16(1),
	},
	38: {
		Flex_state:          uint16(6),
		Fexternal_lex_state: uint16(1),
	},
	39: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(1),
	},
	40: {
		Fexternal_lex_state: uint16(1),
	},
	41: {
		Fexternal_lex_state: uint16(1),
	},
	42: {
		Fexternal_lex_state: uint16(1),
	},
	43: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(1),
	},
	44: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(1),
	},
	45: {
		Flex_state:          uint16(9),
		Fexternal_lex_state: uint16(1),
	},
	46: {
		Fexternal_lex_state: uint16(1),
	},
	47: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(1),
	},
}

var ts_parse_table = [25][305]uint16_t{
	0: {
		0:   uint16(1),
		1:   uint16(1),
		2:   uint16(1),
		4:   uint16(1),
		5:   uint16(1),
		6:   uint16(1),
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
		46:  uint16(1),
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
		135: uint16(1),
		136: uint16(1),
		137: uint16(1),
		138: uint16(1),
		139: uint16(1),
		140: uint16(1),
		141: uint16(1),
		142: uint16(1),
		143: uint16(1),
		144: uint16(1),
		145: uint16(1),
		146: uint16(1),
		147: uint16(1),
		148: uint16(1),
		149: uint16(1),
		150: uint16(1),
		151: uint16(1),
		152: uint16(1),
		153: uint16(1),
		154: uint16(1),
		155: uint16(1),
		156: uint16(1),
		157: uint16(1),
		158: uint16(1),
		159: uint16(1),
		160: uint16(1),
		161: uint16(1),
		162: uint16(1),
		163: uint16(1),
		164: uint16(1),
		165: uint16(1),
		166: uint16(1),
		167: uint16(1),
		168: uint16(1),
		169: uint16(1),
		170: uint16(1),
		171: uint16(1),
		172: uint16(1),
		173: uint16(1),
		174: uint16(1),
		175: uint16(1),
		176: uint16(1),
		177: uint16(1),
		178: uint16(1),
		179: uint16(1),
		180: uint16(1),
		181: uint16(1),
		182: uint16(1),
		183: uint16(1),
		184: uint16(1),
		185: uint16(1),
		186: uint16(1),
		187: uint16(1),
		188: uint16(1),
		189: uint16(1),
		190: uint16(1),
		191: uint16(1),
		192: uint16(1),
		193: uint16(1),
		194: uint16(1),
		195: uint16(1),
		196: uint16(1),
		197: uint16(1),
		198: uint16(1),
		199: uint16(1),
		200: uint16(1),
		201: uint16(1),
		202: uint16(1),
		203: uint16(1),
		204: uint16(1),
		205: uint16(1),
		206: uint16(1),
		207: uint16(1),
		208: uint16(1),
		209: uint16(1),
		210: uint16(1),
		211: uint16(1),
		212: uint16(1),
		213: uint16(1),
		214: uint16(1),
		215: uint16(1),
		216: uint16(1),
		217: uint16(1),
		218: uint16(1),
		219: uint16(1),
		220: uint16(1),
		221: uint16(1),
		222: uint16(1),
		223: uint16(1),
		224: uint16(1),
		225: uint16(1),
		226: uint16(1),
		227: uint16(1),
		228: uint16(1),
		229: uint16(1),
		230: uint16(1),
		231: uint16(1),
		232: uint16(1),
		233: uint16(1),
		234: uint16(1),
		235: uint16(1),
		236: uint16(1),
		237: uint16(1),
		238: uint16(1),
		239: uint16(1),
		240: uint16(1),
		241: uint16(1),
		242: uint16(1),
		243: uint16(1),
		244: uint16(1),
		245: uint16(1),
		246: uint16(1),
		247: uint16(1),
		248: uint16(1),
		249: uint16(1),
		250: uint16(1),
		251: uint16(1),
		252: uint16(1),
		253: uint16(1),
		254: uint16(1),
		255: uint16(1),
		256: uint16(1),
		257: uint16(1),
		258: uint16(1),
		259: uint16(1),
		260: uint16(1),
		261: uint16(1),
		262: uint16(1),
		263: uint16(1),
		264: uint16(1),
		265: uint16(1),
		266: uint16(1),
		268: uint16(1),
		269: uint16(1),
		271: uint16(1),
		272: uint16(1),
		273: uint16(1),
		274: uint16(1),
		275: uint16(1),
		276: uint16(1),
		277: uint16(1),
		278: uint16(1),
		279: uint16(1),
		280: uint16(1),
		281: uint16(1),
		282: uint16(1),
		284: uint16(1),
		285: uint16(3),
	},
	1: {
		0:   uint16(5),
		2:   uint16(7),
		6:   uint16(9),
		264: uint16(11),
		268: uint16(13),
		285: uint16(3),
		286: uint16(42),
		287: uint16(29),
		288: uint16(29),
		290: uint16(29),
		291: uint16(29),
		293: uint16(3),
		296: uint16(4),
		302: uint16(29),
	},
	2: {
		0:   uint16(15),
		1:   uint16(17),
		2:   uint16(15),
		5:   uint16(15),
		6:   uint16(15),
		8:   uint16(20),
		9:   uint16(23),
		10:  uint16(23),
		11:  uint16(23),
		12:  uint16(23),
		13:  uint16(23),
		14:  uint16(23),
		15:  uint16(23),
		16:  uint16(23),
		17:  uint16(23),
		18:  uint16(23),
		19:  uint16(23),
		20:  uint16(23),
		21:  uint16(23),
		22:  uint16(23),
		23:  uint16(23),
		24:  uint16(23),
		25:  uint16(23),
		26:  uint16(23),
		27:  uint16(23),
		28:  uint16(23),
		29:  uint16(23),
		30:  uint16(23),
		31:  uint16(23),
		32:  uint16(23),
		33:  uint16(23),
		34:  uint16(23),
		35:  uint16(23),
		36:  uint16(23),
		37:  uint16(23),
		38:  uint16(23),
		39:  uint16(23),
		40:  uint16(20),
		41:  uint16(23),
		42:  uint16(23),
		43:  uint16(23),
		44:  uint16(23),
		45:  uint16(23),
		46:  uint16(23),
		47:  uint16(23),
		48:  uint16(23),
		49:  uint16(23),
		50:  uint16(23),
		51:  uint16(23),
		52:  uint16(23),
		53:  uint16(23),
		54:  uint16(23),
		55:  uint16(23),
		56:  uint16(23),
		57:  uint16(23),
		58:  uint16(23),
		59:  uint16(23),
		60:  uint16(23),
		61:  uint16(23),
		62:  uint16(23),
		63:  uint16(23),
		64:  uint16(23),
		65:  uint16(23),
		66:  uint16(23),
		67:  uint16(23),
		68:  uint16(23),
		69:  uint16(23),
		70:  uint16(23),
		71:  uint16(23),
		72:  uint16(20),
		73:  uint16(20),
		74:  uint16(20),
		75:  uint16(20),
		76:  uint16(20),
		77:  uint16(20),
		78:  uint16(20),
		79:  uint16(20),
		80:  uint16(20),
		81:  uint16(20),
		82:  uint16(20),
		83:  uint16(20),
		84:  uint16(20),
		85:  uint16(20),
		86:  uint16(20),
		87:  uint16(20),
		88:  uint16(20),
		89:  uint16(20),
		90:  uint16(20),
		91:  uint16(20),
		92:  uint16(20),
		93:  uint16(20),
		94:  uint16(20),
		95:  uint16(20),
		96:  uint16(20),
		97:  uint16(20),
		98:  uint16(20),
		99:  uint16(20),
		100: uint16(20),
		101: uint16(20),
		102: uint16(20),
		103: uint16(20),
		104: uint16(20),
		105: uint16(20),
		106: uint16(20),
		107: uint16(20),
		108: uint16(20),
		109: uint16(20),
		110: uint16(20),
		111: uint16(20),
		112: uint16(20),
		113: uint16(20),
		114: uint16(20),
		115: uint16(20),
		116: uint16(20),
		117: uint16(20),
		118: uint16(20),
		119: uint16(20),
		120: uint16(20),
		121: uint16(20),
		122: uint16(20),
		123: uint16(20),
		124: uint16(20),
		125: uint16(20),
		126: uint16(20),
		127: uint16(20),
		128: uint16(20),
		129: uint16(20),
		130: uint16(20),
		131: uint16(20),
		132: uint16(20),
		133: uint16(20),
		134: uint16(20),
		135: uint16(20),
		136: uint16(23),
		137: uint16(23),
		138: uint16(23),
		139: uint16(23),
		140: uint16(23),
		141: uint16(23),
		142: uint16(23),
		143: uint16(23),
		144: uint16(23),
		145: uint16(23),
		146: uint16(23),
		147: uint16(23),
		148: uint16(23),
		149: uint16(23),
		150: uint16(23),
		151: uint16(23),
		152: uint16(23),
		153: uint16(23),
		154: uint16(23),
		155: uint16(23),
		156: uint16(23),
		157: uint16(23),
		158: uint16(23),
		159: uint16(23),
		160: uint16(23),
		161: uint16(23),
		162: uint16(23),
		163: uint16(23),
		164: uint16(23),
		165: uint16(23),
		166: uint16(23),
		167: uint16(23),
		168: uint16(23),
		169: uint16(23),
		170: uint16(23),
		171: uint16(23),
		172: uint16(23),
		173: uint16(23),
		174: uint16(23),
		175: uint16(23),
		176: uint16(23),
		177: uint16(23),
		178: uint16(23),
		179: uint16(23),
		180: uint16(23),
		181: uint16(23),
		182: uint16(23),
		183: uint16(23),
		184: uint16(23),
		185: uint16(23),
		186: uint16(23),
		187: uint16(23),
		188: uint16(23),
		189: uint16(23),
		190: uint16(23),
		191: uint16(23),
		192: uint16(23),
		193: uint16(23),
		194: uint16(23),
		195: uint16(23),
		196: uint16(23),
		197: uint16(23),
		198: uint16(23),
		199: uint16(23),
		200: uint16(20),
		201: uint16(20),
		202: uint16(20),
		203: uint16(20),
		204: uint16(20),
		205: uint16(20),
		206: uint16(20),
		207: uint16(20),
		208: uint16(20),
		209: uint16(20),
		210: uint16(20),
		211: uint16(20),
		212: uint16(20),
		213: uint16(20),
		214: uint16(20),
		215: uint16(20),
		216: uint16(20),
		217: uint16(20),
		218: uint16(20),
		219: uint16(20),
		220: uint16(20),
		221: uint16(20),
		222: uint16(20),
		223: uint16(20),
		224: uint16(20),
		225: uint16(20),
		226: uint16(20),
		227: uint16(20),
		228: uint16(20),
		229: uint16(20),
		230: uint16(20),
		231: uint16(20),
		232: uint16(20),
		233: uint16(20),
		234: uint16(20),
		235: uint16(20),
		236: uint16(20),
		237: uint16(20),
		238: uint16(20),
		239: uint16(20),
		240: uint16(20),
		241: uint16(20),
		242: uint16(20),
		243: uint16(20),
		244: uint16(20),
		245: uint16(20),
		246: uint16(20),
		247: uint16(20),
		248: uint16(20),
		249: uint16(20),
		250: uint16(20),
		251: uint16(20),
		252: uint16(20),
		253: uint16(20),
		254: uint16(20),
		255: uint16(20),
		256: uint16(20),
		257: uint16(20),
		258: uint16(20),
		259: uint16(20),
		260: uint16(20),
		261: uint16(20),
		262: uint16(20),
		263: uint16(20),
		264: uint16(15),
		265: uint16(26),
		266: uint16(29),
		268: uint16(15),
		271: uint16(32),
		272: uint16(32),
		273: uint16(35),
		274: uint16(35),
		275: uint16(38),
		276: uint16(38),
		277: uint16(41),
		278: uint16(41),
		279: uint16(44),
		280: uint16(47),
		281: uint16(15),
		282: uint16(50),
		284: uint16(53),
		285: uint16(3),
		289: uint16(2),
		292: uint16(2),
		294: uint16(2),
		295: uint16(2),
		297: uint16(2),
		298: uint16(2),
		299: uint16(25),
		300: uint16(2),
		301: uint16(2),
		303: uint16(2),
	},
	3: {
		0:   uint16(56),
		1:   uint16(58),
		2:   uint16(56),
		6:   uint16(56),
		8:   uint16(60),
		9:   uint16(62),
		10:  uint16(62),
		11:  uint16(62),
		12:  uint16(62),
		13:  uint16(62),
		14:  uint16(62),
		15:  uint16(62),
		16:  uint16(62),
		17:  uint16(62),
		18:  uint16(62),
		19:  uint16(62),
		20:  uint16(62),
		21:  uint16(62),
		22:  uint16(62),
		23:  uint16(62),
		24:  uint16(62),
		25:  uint16(62),
		26:  uint16(62),
		27:  uint16(62),
		28:  uint16(62),
		29:  uint16(62),
		30:  uint16(62),
		31:  uint16(62),
		32:  uint16(62),
		33:  uint16(62),
		34:  uint16(62),
		35:  uint16(62),
		36:  uint16(62),
		37:  uint16(62),
		38:  uint16(62),
		39:  uint16(62),
		40:  uint16(60),
		41:  uint16(62),
		42:  uint16(62),
		43:  uint16(62),
		44:  uint16(62),
		45:  uint16(62),
		46:  uint16(62),
		47:  uint16(62),
		48:  uint16(62),
		49:  uint16(62),
		50:  uint16(62),
		51:  uint16(62),
		52:  uint16(62),
		53:  uint16(62),
		54:  uint16(62),
		55:  uint16(62),
		56:  uint16(62),
		57:  uint16(62),
		58:  uint16(62),
		59:  uint16(62),
		60:  uint16(62),
		61:  uint16(62),
		62:  uint16(62),
		63:  uint16(62),
		64:  uint16(62),
		65:  uint16(62),
		66:  uint16(62),
		67:  uint16(62),
		68:  uint16(62),
		69:  uint16(62),
		70:  uint16(62),
		71:  uint16(62),
		72:  uint16(60),
		73:  uint16(60),
		74:  uint16(60),
		75:  uint16(60),
		76:  uint16(60),
		77:  uint16(60),
		78:  uint16(60),
		79:  uint16(60),
		80:  uint16(60),
		81:  uint16(60),
		82:  uint16(60),
		83:  uint16(60),
		84:  uint16(60),
		85:  uint16(60),
		86:  uint16(60),
		87:  uint16(60),
		88:  uint16(60),
		89:  uint16(60),
		90:  uint16(60),
		91:  uint16(60),
		92:  uint16(60),
		93:  uint16(60),
		94:  uint16(60),
		95:  uint16(60),
		96:  uint16(60),
		97:  uint16(60),
		98:  uint16(60),
		99:  uint16(60),
		100: uint16(60),
		101: uint16(60),
		102: uint16(60),
		103: uint16(60),
		104: uint16(60),
		105: uint16(60),
		106: uint16(60),
		107: uint16(60),
		108: uint16(60),
		109: uint16(60),
		110: uint16(60),
		111: uint16(60),
		112: uint16(60),
		113: uint16(60),
		114: uint16(60),
		115: uint16(60),
		116: uint16(60),
		117: uint16(60),
		118: uint16(60),
		119: uint16(60),
		120: uint16(60),
		121: uint16(60),
		122: uint16(60),
		123: uint16(60),
		124: uint16(60),
		125: uint16(60),
		126: uint16(60),
		127: uint16(60),
		128: uint16(60),
		129: uint16(60),
		130: uint16(60),
		131: uint16(60),
		132: uint16(60),
		133: uint16(60),
		134: uint16(60),
		135: uint16(60),
		136: uint16(62),
		137: uint16(62),
		138: uint16(62),
		139: uint16(62),
		140: uint16(62),
		141: uint16(62),
		142: uint16(62),
		143: uint16(62),
		144: uint16(62),
		145: uint16(62),
		146: uint16(62),
		147: uint16(62),
		148: uint16(62),
		149: uint16(62),
		150: uint16(62),
		151: uint16(62),
		152: uint16(62),
		153: uint16(62),
		154: uint16(62),
		155: uint16(62),
		156: uint16(62),
		157: uint16(62),
		158: uint16(62),
		159: uint16(62),
		160: uint16(62),
		161: uint16(62),
		162: uint16(62),
		163: uint16(62),
		164: uint16(62),
		165: uint16(62),
		166: uint16(62),
		167: uint16(62),
		168: uint16(62),
		169: uint16(62),
		170: uint16(62),
		171: uint16(62),
		172: uint16(62),
		173: uint16(62),
		174: uint16(62),
		175: uint16(62),
		176: uint16(62),
		177: uint16(62),
		178: uint16(62),
		179: uint16(62),
		180: uint16(62),
		181: uint16(62),
		182: uint16(62),
		183: uint16(62),
		184: uint16(62),
		185: uint16(62),
		186: uint16(62),
		187: uint16(62),
		188: uint16(62),
		189: uint16(62),
		190: uint16(62),
		191: uint16(62),
		192: uint16(62),
		193: uint16(62),
		194: uint16(62),
		195: uint16(62),
		196: uint16(62),
		197: uint16(62),
		198: uint16(62),
		199: uint16(62),
		200: uint16(60),
		201: uint16(60),
		202: uint16(60),
		203: uint16(60),
		204: uint16(60),
		205: uint16(60),
		206: uint16(60),
		207: uint16(60),
		208: uint16(60),
		209: uint16(60),
		210: uint16(60),
		211: uint16(60),
		212: uint16(60),
		213: uint16(60),
		214: uint16(60),
		215: uint16(60),
		216: uint16(60),
		217: uint16(60),
		218: uint16(60),
		219: uint16(60),
		220: uint16(60),
		221: uint16(60),
		222: uint16(60),
		223: uint16(60),
		224: uint16(60),
		225: uint16(60),
		226: uint16(60),
		227: uint16(60),
		228: uint16(60),
		229: uint16(60),
		230: uint16(60),
		231: uint16(60),
		232: uint16(60),
		233: uint16(60),
		234: uint16(60),
		235: uint16(60),
		236: uint16(60),
		237: uint16(60),
		238: uint16(60),
		239: uint16(60),
		240: uint16(60),
		241: uint16(60),
		242: uint16(60),
		243: uint16(60),
		244: uint16(60),
		245: uint16(60),
		246: uint16(60),
		247: uint16(60),
		248: uint16(60),
		249: uint16(60),
		250: uint16(60),
		251: uint16(60),
		252: uint16(60),
		253: uint16(60),
		254: uint16(60),
		255: uint16(60),
		256: uint16(60),
		257: uint16(60),
		258: uint16(60),
		259: uint16(60),
		260: uint16(60),
		261: uint16(60),
		262: uint16(60),
		263: uint16(60),
		264: uint16(56),
		265: uint16(64),
		266: uint16(66),
		268: uint16(56),
		271: uint16(68),
		272: uint16(68),
		273: uint16(70),
		274: uint16(70),
		275: uint16(72),
		276: uint16(72),
		277: uint16(74),
		278: uint16(74),
		279: uint16(76),
		280: uint16(78),
		282: uint16(80),
		284: uint16(82),
		285: uint16(3),
		289: uint16(6),
		292: uint16(6),
		294: uint16(6),
		295: uint16(6),
		297: uint16(6),
		298: uint16(6),
		299: uint16(25),
		300: uint16(6),
		301: uint16(6),
		303: uint16(6),
	},
	4: {
		0:   uint16(84),
		1:   uint16(58),
		2:   uint16(84),
		6:   uint16(84),
		8:   uint16(60),
		9:   uint16(62),
		10:  uint16(62),
		11:  uint16(62),
		12:  uint16(62),
		13:  uint16(62),
		14:  uint16(62),
		15:  uint16(62),
		16:  uint16(62),
		17:  uint16(62),
		18:  uint16(62),
		19:  uint16(62),
		20:  uint16(62),
		21:  uint16(62),
		22:  uint16(62),
		23:  uint16(62),
		24:  uint16(62),
		25:  uint16(62),
		26:  uint16(62),
		27:  uint16(62),
		28:  uint16(62),
		29:  uint16(62),
		30:  uint16(62),
		31:  uint16(62),
		32:  uint16(62),
		33:  uint16(62),
		34:  uint16(62),
		35:  uint16(62),
		36:  uint16(62),
		37:  uint16(62),
		38:  uint16(62),
		39:  uint16(62),
		40:  uint16(60),
		41:  uint16(62),
		42:  uint16(62),
		43:  uint16(62),
		44:  uint16(62),
		45:  uint16(62),
		46:  uint16(62),
		47:  uint16(62),
		48:  uint16(62),
		49:  uint16(62),
		50:  uint16(62),
		51:  uint16(62),
		52:  uint16(62),
		53:  uint16(62),
		54:  uint16(62),
		55:  uint16(62),
		56:  uint16(62),
		57:  uint16(62),
		58:  uint16(62),
		59:  uint16(62),
		60:  uint16(62),
		61:  uint16(62),
		62:  uint16(62),
		63:  uint16(62),
		64:  uint16(62),
		65:  uint16(62),
		66:  uint16(62),
		67:  uint16(62),
		68:  uint16(62),
		69:  uint16(62),
		70:  uint16(62),
		71:  uint16(62),
		72:  uint16(60),
		73:  uint16(60),
		74:  uint16(60),
		75:  uint16(60),
		76:  uint16(60),
		77:  uint16(60),
		78:  uint16(60),
		79:  uint16(60),
		80:  uint16(60),
		81:  uint16(60),
		82:  uint16(60),
		83:  uint16(60),
		84:  uint16(60),
		85:  uint16(60),
		86:  uint16(60),
		87:  uint16(60),
		88:  uint16(60),
		89:  uint16(60),
		90:  uint16(60),
		91:  uint16(60),
		92:  uint16(60),
		93:  uint16(60),
		94:  uint16(60),
		95:  uint16(60),
		96:  uint16(60),
		97:  uint16(60),
		98:  uint16(60),
		99:  uint16(60),
		100: uint16(60),
		101: uint16(60),
		102: uint16(60),
		103: uint16(60),
		104: uint16(60),
		105: uint16(60),
		106: uint16(60),
		107: uint16(60),
		108: uint16(60),
		109: uint16(60),
		110: uint16(60),
		111: uint16(60),
		112: uint16(60),
		113: uint16(60),
		114: uint16(60),
		115: uint16(60),
		116: uint16(60),
		117: uint16(60),
		118: uint16(60),
		119: uint16(60),
		120: uint16(60),
		121: uint16(60),
		122: uint16(60),
		123: uint16(60),
		124: uint16(60),
		125: uint16(60),
		126: uint16(60),
		127: uint16(60),
		128: uint16(60),
		129: uint16(60),
		130: uint16(60),
		131: uint16(60),
		132: uint16(60),
		133: uint16(60),
		134: uint16(60),
		135: uint16(60),
		136: uint16(62),
		137: uint16(62),
		138: uint16(62),
		139: uint16(62),
		140: uint16(62),
		141: uint16(62),
		142: uint16(62),
		143: uint16(62),
		144: uint16(62),
		145: uint16(62),
		146: uint16(62),
		147: uint16(62),
		148: uint16(62),
		149: uint16(62),
		150: uint16(62),
		151: uint16(62),
		152: uint16(62),
		153: uint16(62),
		154: uint16(62),
		155: uint16(62),
		156: uint16(62),
		157: uint16(62),
		158: uint16(62),
		159: uint16(62),
		160: uint16(62),
		161: uint16(62),
		162: uint16(62),
		163: uint16(62),
		164: uint16(62),
		165: uint16(62),
		166: uint16(62),
		167: uint16(62),
		168: uint16(62),
		169: uint16(62),
		170: uint16(62),
		171: uint16(62),
		172: uint16(62),
		173: uint16(62),
		174: uint16(62),
		175: uint16(62),
		176: uint16(62),
		177: uint16(62),
		178: uint16(62),
		179: uint16(62),
		180: uint16(62),
		181: uint16(62),
		182: uint16(62),
		183: uint16(62),
		184: uint16(62),
		185: uint16(62),
		186: uint16(62),
		187: uint16(62),
		188: uint16(62),
		189: uint16(62),
		190: uint16(62),
		191: uint16(62),
		192: uint16(62),
		193: uint16(62),
		194: uint16(62),
		195: uint16(62),
		196: uint16(62),
		197: uint16(62),
		198: uint16(62),
		199: uint16(62),
		200: uint16(60),
		201: uint16(60),
		202: uint16(60),
		203: uint16(60),
		204: uint16(60),
		205: uint16(60),
		206: uint16(60),
		207: uint16(60),
		208: uint16(60),
		209: uint16(60),
		210: uint16(60),
		211: uint16(60),
		212: uint16(60),
		213: uint16(60),
		214: uint16(60),
		215: uint16(60),
		216: uint16(60),
		217: uint16(60),
		218: uint16(60),
		219: uint16(60),
		220: uint16(60),
		221: uint16(60),
		222: uint16(60),
		223: uint16(60),
		224: uint16(60),
		225: uint16(60),
		226: uint16(60),
		227: uint16(60),
		228: uint16(60),
		229: uint16(60),
		230: uint16(60),
		231: uint16(60),
		232: uint16(60),
		233: uint16(60),
		234: uint16(60),
		235: uint16(60),
		236: uint16(60),
		237: uint16(60),
		238: uint16(60),
		239: uint16(60),
		240: uint16(60),
		241: uint16(60),
		242: uint16(60),
		243: uint16(60),
		244: uint16(60),
		245: uint16(60),
		246: uint16(60),
		247: uint16(60),
		248: uint16(60),
		249: uint16(60),
		250: uint16(60),
		251: uint16(60),
		252: uint16(60),
		253: uint16(60),
		254: uint16(60),
		255: uint16(60),
		256: uint16(60),
		257: uint16(60),
		258: uint16(60),
		259: uint16(60),
		260: uint16(60),
		261: uint16(60),
		262: uint16(60),
		263: uint16(60),
		264: uint16(84),
		265: uint16(64),
		266: uint16(66),
		268: uint16(84),
		271: uint16(68),
		272: uint16(68),
		273: uint16(70),
		274: uint16(70),
		275: uint16(72),
		276: uint16(72),
		277: uint16(74),
		278: uint16(74),
		279: uint16(76),
		280: uint16(78),
		282: uint16(80),
		284: uint16(86),
		285: uint16(3),
		289: uint16(5),
		292: uint16(5),
		294: uint16(5),
		295: uint16(5),
		297: uint16(5),
		298: uint16(5),
		299: uint16(25),
		300: uint16(5),
		301: uint16(5),
		303: uint16(5),
	},
	5: {
		0:   uint16(88),
		1:   uint16(58),
		2:   uint16(88),
		6:   uint16(88),
		8:   uint16(60),
		9:   uint16(62),
		10:  uint16(62),
		11:  uint16(62),
		12:  uint16(62),
		13:  uint16(62),
		14:  uint16(62),
		15:  uint16(62),
		16:  uint16(62),
		17:  uint16(62),
		18:  uint16(62),
		19:  uint16(62),
		20:  uint16(62),
		21:  uint16(62),
		22:  uint16(62),
		23:  uint16(62),
		24:  uint16(62),
		25:  uint16(62),
		26:  uint16(62),
		27:  uint16(62),
		28:  uint16(62),
		29:  uint16(62),
		30:  uint16(62),
		31:  uint16(62),
		32:  uint16(62),
		33:  uint16(62),
		34:  uint16(62),
		35:  uint16(62),
		36:  uint16(62),
		37:  uint16(62),
		38:  uint16(62),
		39:  uint16(62),
		40:  uint16(60),
		41:  uint16(62),
		42:  uint16(62),
		43:  uint16(62),
		44:  uint16(62),
		45:  uint16(62),
		46:  uint16(62),
		47:  uint16(62),
		48:  uint16(62),
		49:  uint16(62),
		50:  uint16(62),
		51:  uint16(62),
		52:  uint16(62),
		53:  uint16(62),
		54:  uint16(62),
		55:  uint16(62),
		56:  uint16(62),
		57:  uint16(62),
		58:  uint16(62),
		59:  uint16(62),
		60:  uint16(62),
		61:  uint16(62),
		62:  uint16(62),
		63:  uint16(62),
		64:  uint16(62),
		65:  uint16(62),
		66:  uint16(62),
		67:  uint16(62),
		68:  uint16(62),
		69:  uint16(62),
		70:  uint16(62),
		71:  uint16(62),
		72:  uint16(60),
		73:  uint16(60),
		74:  uint16(60),
		75:  uint16(60),
		76:  uint16(60),
		77:  uint16(60),
		78:  uint16(60),
		79:  uint16(60),
		80:  uint16(60),
		81:  uint16(60),
		82:  uint16(60),
		83:  uint16(60),
		84:  uint16(60),
		85:  uint16(60),
		86:  uint16(60),
		87:  uint16(60),
		88:  uint16(60),
		89:  uint16(60),
		90:  uint16(60),
		91:  uint16(60),
		92:  uint16(60),
		93:  uint16(60),
		94:  uint16(60),
		95:  uint16(60),
		96:  uint16(60),
		97:  uint16(60),
		98:  uint16(60),
		99:  uint16(60),
		100: uint16(60),
		101: uint16(60),
		102: uint16(60),
		103: uint16(60),
		104: uint16(60),
		105: uint16(60),
		106: uint16(60),
		107: uint16(60),
		108: uint16(60),
		109: uint16(60),
		110: uint16(60),
		111: uint16(60),
		112: uint16(60),
		113: uint16(60),
		114: uint16(60),
		115: uint16(60),
		116: uint16(60),
		117: uint16(60),
		118: uint16(60),
		119: uint16(60),
		120: uint16(60),
		121: uint16(60),
		122: uint16(60),
		123: uint16(60),
		124: uint16(60),
		125: uint16(60),
		126: uint16(60),
		127: uint16(60),
		128: uint16(60),
		129: uint16(60),
		130: uint16(60),
		131: uint16(60),
		132: uint16(60),
		133: uint16(60),
		134: uint16(60),
		135: uint16(60),
		136: uint16(62),
		137: uint16(62),
		138: uint16(62),
		139: uint16(62),
		140: uint16(62),
		141: uint16(62),
		142: uint16(62),
		143: uint16(62),
		144: uint16(62),
		145: uint16(62),
		146: uint16(62),
		147: uint16(62),
		148: uint16(62),
		149: uint16(62),
		150: uint16(62),
		151: uint16(62),
		152: uint16(62),
		153: uint16(62),
		154: uint16(62),
		155: uint16(62),
		156: uint16(62),
		157: uint16(62),
		158: uint16(62),
		159: uint16(62),
		160: uint16(62),
		161: uint16(62),
		162: uint16(62),
		163: uint16(62),
		164: uint16(62),
		165: uint16(62),
		166: uint16(62),
		167: uint16(62),
		168: uint16(62),
		169: uint16(62),
		170: uint16(62),
		171: uint16(62),
		172: uint16(62),
		173: uint16(62),
		174: uint16(62),
		175: uint16(62),
		176: uint16(62),
		177: uint16(62),
		178: uint16(62),
		179: uint16(62),
		180: uint16(62),
		181: uint16(62),
		182: uint16(62),
		183: uint16(62),
		184: uint16(62),
		185: uint16(62),
		186: uint16(62),
		187: uint16(62),
		188: uint16(62),
		189: uint16(62),
		190: uint16(62),
		191: uint16(62),
		192: uint16(62),
		193: uint16(62),
		194: uint16(62),
		195: uint16(62),
		196: uint16(62),
		197: uint16(62),
		198: uint16(62),
		199: uint16(62),
		200: uint16(60),
		201: uint16(60),
		202: uint16(60),
		203: uint16(60),
		204: uint16(60),
		205: uint16(60),
		206: uint16(60),
		207: uint16(60),
		208: uint16(60),
		209: uint16(60),
		210: uint16(60),
		211: uint16(60),
		212: uint16(60),
		213: uint16(60),
		214: uint16(60),
		215: uint16(60),
		216: uint16(60),
		217: uint16(60),
		218: uint16(60),
		219: uint16(60),
		220: uint16(60),
		221: uint16(60),
		222: uint16(60),
		223: uint16(60),
		224: uint16(60),
		225: uint16(60),
		226: uint16(60),
		227: uint16(60),
		228: uint16(60),
		229: uint16(60),
		230: uint16(60),
		231: uint16(60),
		232: uint16(60),
		233: uint16(60),
		234: uint16(60),
		235: uint16(60),
		236: uint16(60),
		237: uint16(60),
		238: uint16(60),
		239: uint16(60),
		240: uint16(60),
		241: uint16(60),
		242: uint16(60),
		243: uint16(60),
		244: uint16(60),
		245: uint16(60),
		246: uint16(60),
		247: uint16(60),
		248: uint16(60),
		249: uint16(60),
		250: uint16(60),
		251: uint16(60),
		252: uint16(60),
		253: uint16(60),
		254: uint16(60),
		255: uint16(60),
		256: uint16(60),
		257: uint16(60),
		258: uint16(60),
		259: uint16(60),
		260: uint16(60),
		261: uint16(60),
		262: uint16(60),
		263: uint16(60),
		264: uint16(88),
		265: uint16(64),
		266: uint16(66),
		268: uint16(88),
		271: uint16(68),
		272: uint16(68),
		273: uint16(70),
		274: uint16(70),
		275: uint16(72),
		276: uint16(72),
		277: uint16(74),
		278: uint16(74),
		279: uint16(76),
		280: uint16(78),
		282: uint16(80),
		284: uint16(90),
		285: uint16(3),
		289: uint16(2),
		292: uint16(2),
		294: uint16(2),
		295: uint16(2),
		297: uint16(2),
		298: uint16(2),
		299: uint16(25),
		300: uint16(2),
		301: uint16(2),
		303: uint16(2),
	},
	6: {
		0:   uint16(92),
		1:   uint16(58),
		2:   uint16(92),
		6:   uint16(92),
		8:   uint16(60),
		9:   uint16(62),
		10:  uint16(62),
		11:  uint16(62),
		12:  uint16(62),
		13:  uint16(62),
		14:  uint16(62),
		15:  uint16(62),
		16:  uint16(62),
		17:  uint16(62),
		18:  uint16(62),
		19:  uint16(62),
		20:  uint16(62),
		21:  uint16(62),
		22:  uint16(62),
		23:  uint16(62),
		24:  uint16(62),
		25:  uint16(62),
		26:  uint16(62),
		27:  uint16(62),
		28:  uint16(62),
		29:  uint16(62),
		30:  uint16(62),
		31:  uint16(62),
		32:  uint16(62),
		33:  uint16(62),
		34:  uint16(62),
		35:  uint16(62),
		36:  uint16(62),
		37:  uint16(62),
		38:  uint16(62),
		39:  uint16(62),
		40:  uint16(60),
		41:  uint16(62),
		42:  uint16(62),
		43:  uint16(62),
		44:  uint16(62),
		45:  uint16(62),
		46:  uint16(62),
		47:  uint16(62),
		48:  uint16(62),
		49:  uint16(62),
		50:  uint16(62),
		51:  uint16(62),
		52:  uint16(62),
		53:  uint16(62),
		54:  uint16(62),
		55:  uint16(62),
		56:  uint16(62),
		57:  uint16(62),
		58:  uint16(62),
		59:  uint16(62),
		60:  uint16(62),
		61:  uint16(62),
		62:  uint16(62),
		63:  uint16(62),
		64:  uint16(62),
		65:  uint16(62),
		66:  uint16(62),
		67:  uint16(62),
		68:  uint16(62),
		69:  uint16(62),
		70:  uint16(62),
		71:  uint16(62),
		72:  uint16(60),
		73:  uint16(60),
		74:  uint16(60),
		75:  uint16(60),
		76:  uint16(60),
		77:  uint16(60),
		78:  uint16(60),
		79:  uint16(60),
		80:  uint16(60),
		81:  uint16(60),
		82:  uint16(60),
		83:  uint16(60),
		84:  uint16(60),
		85:  uint16(60),
		86:  uint16(60),
		87:  uint16(60),
		88:  uint16(60),
		89:  uint16(60),
		90:  uint16(60),
		91:  uint16(60),
		92:  uint16(60),
		93:  uint16(60),
		94:  uint16(60),
		95:  uint16(60),
		96:  uint16(60),
		97:  uint16(60),
		98:  uint16(60),
		99:  uint16(60),
		100: uint16(60),
		101: uint16(60),
		102: uint16(60),
		103: uint16(60),
		104: uint16(60),
		105: uint16(60),
		106: uint16(60),
		107: uint16(60),
		108: uint16(60),
		109: uint16(60),
		110: uint16(60),
		111: uint16(60),
		112: uint16(60),
		113: uint16(60),
		114: uint16(60),
		115: uint16(60),
		116: uint16(60),
		117: uint16(60),
		118: uint16(60),
		119: uint16(60),
		120: uint16(60),
		121: uint16(60),
		122: uint16(60),
		123: uint16(60),
		124: uint16(60),
		125: uint16(60),
		126: uint16(60),
		127: uint16(60),
		128: uint16(60),
		129: uint16(60),
		130: uint16(60),
		131: uint16(60),
		132: uint16(60),
		133: uint16(60),
		134: uint16(60),
		135: uint16(60),
		136: uint16(62),
		137: uint16(62),
		138: uint16(62),
		139: uint16(62),
		140: uint16(62),
		141: uint16(62),
		142: uint16(62),
		143: uint16(62),
		144: uint16(62),
		145: uint16(62),
		146: uint16(62),
		147: uint16(62),
		148: uint16(62),
		149: uint16(62),
		150: uint16(62),
		151: uint16(62),
		152: uint16(62),
		153: uint16(62),
		154: uint16(62),
		155: uint16(62),
		156: uint16(62),
		157: uint16(62),
		158: uint16(62),
		159: uint16(62),
		160: uint16(62),
		161: uint16(62),
		162: uint16(62),
		163: uint16(62),
		164: uint16(62),
		165: uint16(62),
		166: uint16(62),
		167: uint16(62),
		168: uint16(62),
		169: uint16(62),
		170: uint16(62),
		171: uint16(62),
		172: uint16(62),
		173: uint16(62),
		174: uint16(62),
		175: uint16(62),
		176: uint16(62),
		177: uint16(62),
		178: uint16(62),
		179: uint16(62),
		180: uint16(62),
		181: uint16(62),
		182: uint16(62),
		183: uint16(62),
		184: uint16(62),
		185: uint16(62),
		186: uint16(62),
		187: uint16(62),
		188: uint16(62),
		189: uint16(62),
		190: uint16(62),
		191: uint16(62),
		192: uint16(62),
		193: uint16(62),
		194: uint16(62),
		195: uint16(62),
		196: uint16(62),
		197: uint16(62),
		198: uint16(62),
		199: uint16(62),
		200: uint16(60),
		201: uint16(60),
		202: uint16(60),
		203: uint16(60),
		204: uint16(60),
		205: uint16(60),
		206: uint16(60),
		207: uint16(60),
		208: uint16(60),
		209: uint16(60),
		210: uint16(60),
		211: uint16(60),
		212: uint16(60),
		213: uint16(60),
		214: uint16(60),
		215: uint16(60),
		216: uint16(60),
		217: uint16(60),
		218: uint16(60),
		219: uint16(60),
		220: uint16(60),
		221: uint16(60),
		222: uint16(60),
		223: uint16(60),
		224: uint16(60),
		225: uint16(60),
		226: uint16(60),
		227: uint16(60),
		228: uint16(60),
		229: uint16(60),
		230: uint16(60),
		231: uint16(60),
		232: uint16(60),
		233: uint16(60),
		234: uint16(60),
		235: uint16(60),
		236: uint16(60),
		237: uint16(60),
		238: uint16(60),
		239: uint16(60),
		240: uint16(60),
		241: uint16(60),
		242: uint16(60),
		243: uint16(60),
		244: uint16(60),
		245: uint16(60),
		246: uint16(60),
		247: uint16(60),
		248: uint16(60),
		249: uint16(60),
		250: uint16(60),
		251: uint16(60),
		252: uint16(60),
		253: uint16(60),
		254: uint16(60),
		255: uint16(60),
		256: uint16(60),
		257: uint16(60),
		258: uint16(60),
		259: uint16(60),
		260: uint16(60),
		261: uint16(60),
		262: uint16(60),
		263: uint16(60),
		264: uint16(92),
		265: uint16(64),
		266: uint16(66),
		268: uint16(92),
		271: uint16(68),
		272: uint16(68),
		273: uint16(70),
		274: uint16(70),
		275: uint16(72),
		276: uint16(72),
		277: uint16(74),
		278: uint16(74),
		279: uint16(76),
		280: uint16(78),
		282: uint16(80),
		284: uint16(90),
		285: uint16(3),
		289: uint16(2),
		292: uint16(2),
		294: uint16(2),
		295: uint16(2),
		297: uint16(2),
		298: uint16(2),
		299: uint16(25),
		300: uint16(2),
		301: uint16(2),
		303: uint16(2),
	},
	7: {
		1:   uint16(58),
		8:   uint16(60),
		9:   uint16(62),
		10:  uint16(62),
		11:  uint16(62),
		12:  uint16(62),
		13:  uint16(62),
		14:  uint16(62),
		15:  uint16(62),
		16:  uint16(62),
		17:  uint16(62),
		18:  uint16(62),
		19:  uint16(62),
		20:  uint16(62),
		21:  uint16(62),
		22:  uint16(62),
		23:  uint16(62),
		24:  uint16(62),
		25:  uint16(62),
		26:  uint16(62),
		27:  uint16(62),
		28:  uint16(62),
		29:  uint16(62),
		30:  uint16(62),
		31:  uint16(62),
		32:  uint16(62),
		33:  uint16(62),
		34:  uint16(62),
		35:  uint16(62),
		36:  uint16(62),
		37:  uint16(62),
		38:  uint16(62),
		39:  uint16(62),
		40:  uint16(60),
		41:  uint16(62),
		42:  uint16(62),
		43:  uint16(62),
		44:  uint16(62),
		45:  uint16(62),
		46:  uint16(62),
		47:  uint16(62),
		48:  uint16(62),
		49:  uint16(62),
		50:  uint16(62),
		51:  uint16(62),
		52:  uint16(62),
		53:  uint16(62),
		54:  uint16(62),
		55:  uint16(62),
		56:  uint16(62),
		57:  uint16(62),
		58:  uint16(62),
		59:  uint16(62),
		60:  uint16(62),
		61:  uint16(62),
		62:  uint16(62),
		63:  uint16(62),
		64:  uint16(62),
		65:  uint16(62),
		66:  uint16(62),
		67:  uint16(62),
		68:  uint16(62),
		69:  uint16(62),
		70:  uint16(62),
		71:  uint16(62),
		72:  uint16(60),
		73:  uint16(60),
		74:  uint16(60),
		75:  uint16(60),
		76:  uint16(60),
		77:  uint16(60),
		78:  uint16(60),
		79:  uint16(60),
		80:  uint16(60),
		81:  uint16(60),
		82:  uint16(60),
		83:  uint16(60),
		84:  uint16(60),
		85:  uint16(60),
		86:  uint16(60),
		87:  uint16(60),
		88:  uint16(60),
		89:  uint16(60),
		90:  uint16(60),
		91:  uint16(60),
		92:  uint16(60),
		93:  uint16(60),
		94:  uint16(60),
		95:  uint16(60),
		96:  uint16(60),
		97:  uint16(60),
		98:  uint16(60),
		99:  uint16(60),
		100: uint16(60),
		101: uint16(60),
		102: uint16(60),
		103: uint16(60),
		104: uint16(60),
		105: uint16(60),
		106: uint16(60),
		107: uint16(60),
		108: uint16(60),
		109: uint16(60),
		110: uint16(60),
		111: uint16(60),
		112: uint16(60),
		113: uint16(60),
		114: uint16(60),
		115: uint16(60),
		116: uint16(60),
		117: uint16(60),
		118: uint16(60),
		119: uint16(60),
		120: uint16(60),
		121: uint16(60),
		122: uint16(60),
		123: uint16(60),
		124: uint16(60),
		125: uint16(60),
		126: uint16(60),
		127: uint16(60),
		128: uint16(60),
		129: uint16(60),
		130: uint16(60),
		131: uint16(60),
		132: uint16(60),
		133: uint16(60),
		134: uint16(60),
		135: uint16(60),
		136: uint16(62),
		137: uint16(62),
		138: uint16(62),
		139: uint16(62),
		140: uint16(62),
		141: uint16(62),
		142: uint16(62),
		143: uint16(62),
		144: uint16(62),
		145: uint16(62),
		146: uint16(62),
		147: uint16(62),
		148: uint16(62),
		149: uint16(62),
		150: uint16(62),
		151: uint16(62),
		152: uint16(62),
		153: uint16(62),
		154: uint16(62),
		155: uint16(62),
		156: uint16(62),
		157: uint16(62),
		158: uint16(62),
		159: uint16(62),
		160: uint16(62),
		161: uint16(62),
		162: uint16(62),
		163: uint16(62),
		164: uint16(62),
		165: uint16(62),
		166: uint16(62),
		167: uint16(62),
		168: uint16(62),
		169: uint16(62),
		170: uint16(62),
		171: uint16(62),
		172: uint16(62),
		173: uint16(62),
		174: uint16(62),
		175: uint16(62),
		176: uint16(62),
		177: uint16(62),
		178: uint16(62),
		179: uint16(62),
		180: uint16(62),
		181: uint16(62),
		182: uint16(62),
		183: uint16(62),
		184: uint16(62),
		185: uint16(62),
		186: uint16(62),
		187: uint16(62),
		188: uint16(62),
		189: uint16(62),
		190: uint16(62),
		191: uint16(62),
		192: uint16(62),
		193: uint16(62),
		194: uint16(62),
		195: uint16(62),
		196: uint16(62),
		197: uint16(62),
		198: uint16(62),
		199: uint16(62),
		200: uint16(60),
		201: uint16(60),
		202: uint16(60),
		203: uint16(60),
		204: uint16(60),
		205: uint16(60),
		206: uint16(60),
		207: uint16(60),
		208: uint16(60),
		209: uint16(60),
		210: uint16(60),
		211: uint16(60),
		212: uint16(60),
		213: uint16(60),
		214: uint16(60),
		215: uint16(60),
		216: uint16(60),
		217: uint16(60),
		218: uint16(60),
		219: uint16(60),
		220: uint16(60),
		221: uint16(60),
		222: uint16(60),
		223: uint16(60),
		224: uint16(60),
		225: uint16(60),
		226: uint16(60),
		227: uint16(60),
		228: uint16(60),
		229: uint16(60),
		230: uint16(60),
		231: uint16(60),
		232: uint16(60),
		233: uint16(60),
		234: uint16(60),
		235: uint16(60),
		236: uint16(60),
		237: uint16(60),
		238: uint16(60),
		239: uint16(60),
		240: uint16(60),
		241: uint16(60),
		242: uint16(60),
		243: uint16(60),
		244: uint16(60),
		245: uint16(60),
		246: uint16(60),
		247: uint16(60),
		248: uint16(60),
		249: uint16(60),
		250: uint16(60),
		251: uint16(60),
		252: uint16(60),
		253: uint16(60),
		254: uint16(60),
		255: uint16(60),
		256: uint16(60),
		257: uint16(60),
		258: uint16(60),
		259: uint16(60),
		260: uint16(60),
		261: uint16(60),
		262: uint16(60),
		263: uint16(60),
		265: uint16(64),
		266: uint16(66),
		271: uint16(68),
		272: uint16(68),
		273: uint16(70),
		274: uint16(70),
		275: uint16(72),
		276: uint16(72),
		277: uint16(74),
		278: uint16(74),
		279: uint16(76),
		280: uint16(78),
		281: uint16(94),
		282: uint16(80),
		284: uint16(96),
		285: uint16(3),
		289: uint16(9),
		292: uint16(9),
		294: uint16(9),
		295: uint16(9),
		297: uint16(9),
		298: uint16(9),
		299: uint16(25),
		300: uint16(9),
		301: uint16(9),
		303: uint16(9),
	},
	8: {
		1:   uint16(58),
		5:   uint16(98),
		8:   uint16(60),
		9:   uint16(62),
		10:  uint16(62),
		11:  uint16(62),
		12:  uint16(62),
		13:  uint16(62),
		14:  uint16(62),
		15:  uint16(62),
		16:  uint16(62),
		17:  uint16(62),
		18:  uint16(62),
		19:  uint16(62),
		20:  uint16(62),
		21:  uint16(62),
		22:  uint16(62),
		23:  uint16(62),
		24:  uint16(62),
		25:  uint16(62),
		26:  uint16(62),
		27:  uint16(62),
		28:  uint16(62),
		29:  uint16(62),
		30:  uint16(62),
		31:  uint16(62),
		32:  uint16(62),
		33:  uint16(62),
		34:  uint16(62),
		35:  uint16(62),
		36:  uint16(62),
		37:  uint16(62),
		38:  uint16(62),
		39:  uint16(62),
		40:  uint16(60),
		41:  uint16(62),
		42:  uint16(62),
		43:  uint16(62),
		44:  uint16(62),
		45:  uint16(62),
		46:  uint16(62),
		47:  uint16(62),
		48:  uint16(62),
		49:  uint16(62),
		50:  uint16(62),
		51:  uint16(62),
		52:  uint16(62),
		53:  uint16(62),
		54:  uint16(62),
		55:  uint16(62),
		56:  uint16(62),
		57:  uint16(62),
		58:  uint16(62),
		59:  uint16(62),
		60:  uint16(62),
		61:  uint16(62),
		62:  uint16(62),
		63:  uint16(62),
		64:  uint16(62),
		65:  uint16(62),
		66:  uint16(62),
		67:  uint16(62),
		68:  uint16(62),
		69:  uint16(62),
		70:  uint16(62),
		71:  uint16(62),
		72:  uint16(60),
		73:  uint16(60),
		74:  uint16(60),
		75:  uint16(60),
		76:  uint16(60),
		77:  uint16(60),
		78:  uint16(60),
		79:  uint16(60),
		80:  uint16(60),
		81:  uint16(60),
		82:  uint16(60),
		83:  uint16(60),
		84:  uint16(60),
		85:  uint16(60),
		86:  uint16(60),
		87:  uint16(60),
		88:  uint16(60),
		89:  uint16(60),
		90:  uint16(60),
		91:  uint16(60),
		92:  uint16(60),
		93:  uint16(60),
		94:  uint16(60),
		95:  uint16(60),
		96:  uint16(60),
		97:  uint16(60),
		98:  uint16(60),
		99:  uint16(60),
		100: uint16(60),
		101: uint16(60),
		102: uint16(60),
		103: uint16(60),
		104: uint16(60),
		105: uint16(60),
		106: uint16(60),
		107: uint16(60),
		108: uint16(60),
		109: uint16(60),
		110: uint16(60),
		111: uint16(60),
		112: uint16(60),
		113: uint16(60),
		114: uint16(60),
		115: uint16(60),
		116: uint16(60),
		117: uint16(60),
		118: uint16(60),
		119: uint16(60),
		120: uint16(60),
		121: uint16(60),
		122: uint16(60),
		123: uint16(60),
		124: uint16(60),
		125: uint16(60),
		126: uint16(60),
		127: uint16(60),
		128: uint16(60),
		129: uint16(60),
		130: uint16(60),
		131: uint16(60),
		132: uint16(60),
		133: uint16(60),
		134: uint16(60),
		135: uint16(60),
		136: uint16(62),
		137: uint16(62),
		138: uint16(62),
		139: uint16(62),
		140: uint16(62),
		141: uint16(62),
		142: uint16(62),
		143: uint16(62),
		144: uint16(62),
		145: uint16(62),
		146: uint16(62),
		147: uint16(62),
		148: uint16(62),
		149: uint16(62),
		150: uint16(62),
		151: uint16(62),
		152: uint16(62),
		153: uint16(62),
		154: uint16(62),
		155: uint16(62),
		156: uint16(62),
		157: uint16(62),
		158: uint16(62),
		159: uint16(62),
		160: uint16(62),
		161: uint16(62),
		162: uint16(62),
		163: uint16(62),
		164: uint16(62),
		165: uint16(62),
		166: uint16(62),
		167: uint16(62),
		168: uint16(62),
		169: uint16(62),
		170: uint16(62),
		171: uint16(62),
		172: uint16(62),
		173: uint16(62),
		174: uint16(62),
		175: uint16(62),
		176: uint16(62),
		177: uint16(62),
		178: uint16(62),
		179: uint16(62),
		180: uint16(62),
		181: uint16(62),
		182: uint16(62),
		183: uint16(62),
		184: uint16(62),
		185: uint16(62),
		186: uint16(62),
		187: uint16(62),
		188: uint16(62),
		189: uint16(62),
		190: uint16(62),
		191: uint16(62),
		192: uint16(62),
		193: uint16(62),
		194: uint16(62),
		195: uint16(62),
		196: uint16(62),
		197: uint16(62),
		198: uint16(62),
		199: uint16(62),
		200: uint16(60),
		201: uint16(60),
		202: uint16(60),
		203: uint16(60),
		204: uint16(60),
		205: uint16(60),
		206: uint16(60),
		207: uint16(60),
		208: uint16(60),
		209: uint16(60),
		210: uint16(60),
		211: uint16(60),
		212: uint16(60),
		213: uint16(60),
		214: uint16(60),
		215: uint16(60),
		216: uint16(60),
		217: uint16(60),
		218: uint16(60),
		219: uint16(60),
		220: uint16(60),
		221: uint16(60),
		222: uint16(60),
		223: uint16(60),
		224: uint16(60),
		225: uint16(60),
		226: uint16(60),
		227: uint16(60),
		228: uint16(60),
		229: uint16(60),
		230: uint16(60),
		231: uint16(60),
		232: uint16(60),
		233: uint16(60),
		234: uint16(60),
		235: uint16(60),
		236: uint16(60),
		237: uint16(60),
		238: uint16(60),
		239: uint16(60),
		240: uint16(60),
		241: uint16(60),
		242: uint16(60),
		243: uint16(60),
		244: uint16(60),
		245: uint16(60),
		246: uint16(60),
		247: uint16(60),
		248: uint16(60),
		249: uint16(60),
		250: uint16(60),
		251: uint16(60),
		252: uint16(60),
		253: uint16(60),
		254: uint16(60),
		255: uint16(60),
		256: uint16(60),
		257: uint16(60),
		258: uint16(60),
		259: uint16(60),
		260: uint16(60),
		261: uint16(60),
		262: uint16(60),
		263: uint16(60),
		265: uint16(64),
		266: uint16(66),
		271: uint16(68),
		272: uint16(68),
		273: uint16(70),
		274: uint16(70),
		275: uint16(72),
		276: uint16(72),
		277: uint16(74),
		278: uint16(74),
		279: uint16(76),
		280: uint16(78),
		282: uint16(80),
		284: uint16(90),
		285: uint16(3),
		289: uint16(2),
		292: uint16(2),
		294: uint16(2),
		295: uint16(2),
		297: uint16(2),
		298: uint16(2),
		299: uint16(25),
		300: uint16(2),
		301: uint16(2),
		303: uint16(2),
	},
	9: {
		1:   uint16(58),
		8:   uint16(60),
		9:   uint16(62),
		10:  uint16(62),
		11:  uint16(62),
		12:  uint16(62),
		13:  uint16(62),
		14:  uint16(62),
		15:  uint16(62),
		16:  uint16(62),
		17:  uint16(62),
		18:  uint16(62),
		19:  uint16(62),
		20:  uint16(62),
		21:  uint16(62),
		22:  uint16(62),
		23:  uint16(62),
		24:  uint16(62),
		25:  uint16(62),
		26:  uint16(62),
		27:  uint16(62),
		28:  uint16(62),
		29:  uint16(62),
		30:  uint16(62),
		31:  uint16(62),
		32:  uint16(62),
		33:  uint16(62),
		34:  uint16(62),
		35:  uint16(62),
		36:  uint16(62),
		37:  uint16(62),
		38:  uint16(62),
		39:  uint16(62),
		40:  uint16(60),
		41:  uint16(62),
		42:  uint16(62),
		43:  uint16(62),
		44:  uint16(62),
		45:  uint16(62),
		46:  uint16(62),
		47:  uint16(62),
		48:  uint16(62),
		49:  uint16(62),
		50:  uint16(62),
		51:  uint16(62),
		52:  uint16(62),
		53:  uint16(62),
		54:  uint16(62),
		55:  uint16(62),
		56:  uint16(62),
		57:  uint16(62),
		58:  uint16(62),
		59:  uint16(62),
		60:  uint16(62),
		61:  uint16(62),
		62:  uint16(62),
		63:  uint16(62),
		64:  uint16(62),
		65:  uint16(62),
		66:  uint16(62),
		67:  uint16(62),
		68:  uint16(62),
		69:  uint16(62),
		70:  uint16(62),
		71:  uint16(62),
		72:  uint16(60),
		73:  uint16(60),
		74:  uint16(60),
		75:  uint16(60),
		76:  uint16(60),
		77:  uint16(60),
		78:  uint16(60),
		79:  uint16(60),
		80:  uint16(60),
		81:  uint16(60),
		82:  uint16(60),
		83:  uint16(60),
		84:  uint16(60),
		85:  uint16(60),
		86:  uint16(60),
		87:  uint16(60),
		88:  uint16(60),
		89:  uint16(60),
		90:  uint16(60),
		91:  uint16(60),
		92:  uint16(60),
		93:  uint16(60),
		94:  uint16(60),
		95:  uint16(60),
		96:  uint16(60),
		97:  uint16(60),
		98:  uint16(60),
		99:  uint16(60),
		100: uint16(60),
		101: uint16(60),
		102: uint16(60),
		103: uint16(60),
		104: uint16(60),
		105: uint16(60),
		106: uint16(60),
		107: uint16(60),
		108: uint16(60),
		109: uint16(60),
		110: uint16(60),
		111: uint16(60),
		112: uint16(60),
		113: uint16(60),
		114: uint16(60),
		115: uint16(60),
		116: uint16(60),
		117: uint16(60),
		118: uint16(60),
		119: uint16(60),
		120: uint16(60),
		121: uint16(60),
		122: uint16(60),
		123: uint16(60),
		124: uint16(60),
		125: uint16(60),
		126: uint16(60),
		127: uint16(60),
		128: uint16(60),
		129: uint16(60),
		130: uint16(60),
		131: uint16(60),
		132: uint16(60),
		133: uint16(60),
		134: uint16(60),
		135: uint16(60),
		136: uint16(62),
		137: uint16(62),
		138: uint16(62),
		139: uint16(62),
		140: uint16(62),
		141: uint16(62),
		142: uint16(62),
		143: uint16(62),
		144: uint16(62),
		145: uint16(62),
		146: uint16(62),
		147: uint16(62),
		148: uint16(62),
		149: uint16(62),
		150: uint16(62),
		151: uint16(62),
		152: uint16(62),
		153: uint16(62),
		154: uint16(62),
		155: uint16(62),
		156: uint16(62),
		157: uint16(62),
		158: uint16(62),
		159: uint16(62),
		160: uint16(62),
		161: uint16(62),
		162: uint16(62),
		163: uint16(62),
		164: uint16(62),
		165: uint16(62),
		166: uint16(62),
		167: uint16(62),
		168: uint16(62),
		169: uint16(62),
		170: uint16(62),
		171: uint16(62),
		172: uint16(62),
		173: uint16(62),
		174: uint16(62),
		175: uint16(62),
		176: uint16(62),
		177: uint16(62),
		178: uint16(62),
		179: uint16(62),
		180: uint16(62),
		181: uint16(62),
		182: uint16(62),
		183: uint16(62),
		184: uint16(62),
		185: uint16(62),
		186: uint16(62),
		187: uint16(62),
		188: uint16(62),
		189: uint16(62),
		190: uint16(62),
		191: uint16(62),
		192: uint16(62),
		193: uint16(62),
		194: uint16(62),
		195: uint16(62),
		196: uint16(62),
		197: uint16(62),
		198: uint16(62),
		199: uint16(62),
		200: uint16(60),
		201: uint16(60),
		202: uint16(60),
		203: uint16(60),
		204: uint16(60),
		205: uint16(60),
		206: uint16(60),
		207: uint16(60),
		208: uint16(60),
		209: uint16(60),
		210: uint16(60),
		211: uint16(60),
		212: uint16(60),
		213: uint16(60),
		214: uint16(60),
		215: uint16(60),
		216: uint16(60),
		217: uint16(60),
		218: uint16(60),
		219: uint16(60),
		220: uint16(60),
		221: uint16(60),
		222: uint16(60),
		223: uint16(60),
		224: uint16(60),
		225: uint16(60),
		226: uint16(60),
		227: uint16(60),
		228: uint16(60),
		229: uint16(60),
		230: uint16(60),
		231: uint16(60),
		232: uint16(60),
		233: uint16(60),
		234: uint16(60),
		235: uint16(60),
		236: uint16(60),
		237: uint16(60),
		238: uint16(60),
		239: uint16(60),
		240: uint16(60),
		241: uint16(60),
		242: uint16(60),
		243: uint16(60),
		244: uint16(60),
		245: uint16(60),
		246: uint16(60),
		247: uint16(60),
		248: uint16(60),
		249: uint16(60),
		250: uint16(60),
		251: uint16(60),
		252: uint16(60),
		253: uint16(60),
		254: uint16(60),
		255: uint16(60),
		256: uint16(60),
		257: uint16(60),
		258: uint16(60),
		259: uint16(60),
		260: uint16(60),
		261: uint16(60),
		262: uint16(60),
		263: uint16(60),
		265: uint16(64),
		266: uint16(66),
		271: uint16(68),
		272: uint16(68),
		273: uint16(70),
		274: uint16(70),
		275: uint16(72),
		276: uint16(72),
		277: uint16(74),
		278: uint16(74),
		279: uint16(76),
		280: uint16(78),
		281: uint16(100),
		282: uint16(80),
		284: uint16(90),
		285: uint16(3),
		289: uint16(2),
		292: uint16(2),
		294: uint16(2),
		295: uint16(2),
		297: uint16(2),
		298: uint16(2),
		299: uint16(25),
		300: uint16(2),
		301: uint16(2),
		303: uint16(2),
	},
	10: {
		1:   uint16(58),
		5:   uint16(102),
		8:   uint16(60),
		9:   uint16(62),
		10:  uint16(62),
		11:  uint16(62),
		12:  uint16(62),
		13:  uint16(62),
		14:  uint16(62),
		15:  uint16(62),
		16:  uint16(62),
		17:  uint16(62),
		18:  uint16(62),
		19:  uint16(62),
		20:  uint16(62),
		21:  uint16(62),
		22:  uint16(62),
		23:  uint16(62),
		24:  uint16(62),
		25:  uint16(62),
		26:  uint16(62),
		27:  uint16(62),
		28:  uint16(62),
		29:  uint16(62),
		30:  uint16(62),
		31:  uint16(62),
		32:  uint16(62),
		33:  uint16(62),
		34:  uint16(62),
		35:  uint16(62),
		36:  uint16(62),
		37:  uint16(62),
		38:  uint16(62),
		39:  uint16(62),
		40:  uint16(60),
		41:  uint16(62),
		42:  uint16(62),
		43:  uint16(62),
		44:  uint16(62),
		45:  uint16(62),
		46:  uint16(62),
		47:  uint16(62),
		48:  uint16(62),
		49:  uint16(62),
		50:  uint16(62),
		51:  uint16(62),
		52:  uint16(62),
		53:  uint16(62),
		54:  uint16(62),
		55:  uint16(62),
		56:  uint16(62),
		57:  uint16(62),
		58:  uint16(62),
		59:  uint16(62),
		60:  uint16(62),
		61:  uint16(62),
		62:  uint16(62),
		63:  uint16(62),
		64:  uint16(62),
		65:  uint16(62),
		66:  uint16(62),
		67:  uint16(62),
		68:  uint16(62),
		69:  uint16(62),
		70:  uint16(62),
		71:  uint16(62),
		72:  uint16(60),
		73:  uint16(60),
		74:  uint16(60),
		75:  uint16(60),
		76:  uint16(60),
		77:  uint16(60),
		78:  uint16(60),
		79:  uint16(60),
		80:  uint16(60),
		81:  uint16(60),
		82:  uint16(60),
		83:  uint16(60),
		84:  uint16(60),
		85:  uint16(60),
		86:  uint16(60),
		87:  uint16(60),
		88:  uint16(60),
		89:  uint16(60),
		90:  uint16(60),
		91:  uint16(60),
		92:  uint16(60),
		93:  uint16(60),
		94:  uint16(60),
		95:  uint16(60),
		96:  uint16(60),
		97:  uint16(60),
		98:  uint16(60),
		99:  uint16(60),
		100: uint16(60),
		101: uint16(60),
		102: uint16(60),
		103: uint16(60),
		104: uint16(60),
		105: uint16(60),
		106: uint16(60),
		107: uint16(60),
		108: uint16(60),
		109: uint16(60),
		110: uint16(60),
		111: uint16(60),
		112: uint16(60),
		113: uint16(60),
		114: uint16(60),
		115: uint16(60),
		116: uint16(60),
		117: uint16(60),
		118: uint16(60),
		119: uint16(60),
		120: uint16(60),
		121: uint16(60),
		122: uint16(60),
		123: uint16(60),
		124: uint16(60),
		125: uint16(60),
		126: uint16(60),
		127: uint16(60),
		128: uint16(60),
		129: uint16(60),
		130: uint16(60),
		131: uint16(60),
		132: uint16(60),
		133: uint16(60),
		134: uint16(60),
		135: uint16(60),
		136: uint16(62),
		137: uint16(62),
		138: uint16(62),
		139: uint16(62),
		140: uint16(62),
		141: uint16(62),
		142: uint16(62),
		143: uint16(62),
		144: uint16(62),
		145: uint16(62),
		146: uint16(62),
		147: uint16(62),
		148: uint16(62),
		149: uint16(62),
		150: uint16(62),
		151: uint16(62),
		152: uint16(62),
		153: uint16(62),
		154: uint16(62),
		155: uint16(62),
		156: uint16(62),
		157: uint16(62),
		158: uint16(62),
		159: uint16(62),
		160: uint16(62),
		161: uint16(62),
		162: uint16(62),
		163: uint16(62),
		164: uint16(62),
		165: uint16(62),
		166: uint16(62),
		167: uint16(62),
		168: uint16(62),
		169: uint16(62),
		170: uint16(62),
		171: uint16(62),
		172: uint16(62),
		173: uint16(62),
		174: uint16(62),
		175: uint16(62),
		176: uint16(62),
		177: uint16(62),
		178: uint16(62),
		179: uint16(62),
		180: uint16(62),
		181: uint16(62),
		182: uint16(62),
		183: uint16(62),
		184: uint16(62),
		185: uint16(62),
		186: uint16(62),
		187: uint16(62),
		188: uint16(62),
		189: uint16(62),
		190: uint16(62),
		191: uint16(62),
		192: uint16(62),
		193: uint16(62),
		194: uint16(62),
		195: uint16(62),
		196: uint16(62),
		197: uint16(62),
		198: uint16(62),
		199: uint16(62),
		200: uint16(60),
		201: uint16(60),
		202: uint16(60),
		203: uint16(60),
		204: uint16(60),
		205: uint16(60),
		206: uint16(60),
		207: uint16(60),
		208: uint16(60),
		209: uint16(60),
		210: uint16(60),
		211: uint16(60),
		212: uint16(60),
		213: uint16(60),
		214: uint16(60),
		215: uint16(60),
		216: uint16(60),
		217: uint16(60),
		218: uint16(60),
		219: uint16(60),
		220: uint16(60),
		221: uint16(60),
		222: uint16(60),
		223: uint16(60),
		224: uint16(60),
		225: uint16(60),
		226: uint16(60),
		227: uint16(60),
		228: uint16(60),
		229: uint16(60),
		230: uint16(60),
		231: uint16(60),
		232: uint16(60),
		233: uint16(60),
		234: uint16(60),
		235: uint16(60),
		236: uint16(60),
		237: uint16(60),
		238: uint16(60),
		239: uint16(60),
		240: uint16(60),
		241: uint16(60),
		242: uint16(60),
		243: uint16(60),
		244: uint16(60),
		245: uint16(60),
		246: uint16(60),
		247: uint16(60),
		248: uint16(60),
		249: uint16(60),
		250: uint16(60),
		251: uint16(60),
		252: uint16(60),
		253: uint16(60),
		254: uint16(60),
		255: uint16(60),
		256: uint16(60),
		257: uint16(60),
		258: uint16(60),
		259: uint16(60),
		260: uint16(60),
		261: uint16(60),
		262: uint16(60),
		263: uint16(60),
		265: uint16(64),
		266: uint16(66),
		271: uint16(68),
		272: uint16(68),
		273: uint16(70),
		274: uint16(70),
		275: uint16(72),
		276: uint16(72),
		277: uint16(74),
		278: uint16(74),
		279: uint16(76),
		280: uint16(78),
		282: uint16(80),
		284: uint16(104),
		285: uint16(3),
		289: uint16(8),
		292: uint16(8),
		294: uint16(8),
		295: uint16(8),
		297: uint16(8),
		298: uint16(8),
		299: uint16(25),
		300: uint16(8),
		301: uint16(8),
		303: uint16(8),
	},
	11: {
		0:   uint16(106),
		1:   uint16(108),
		2:   uint16(106),
		5:   uint16(106),
		6:   uint16(106),
		8:   uint16(106),
		9:   uint16(108),
		10:  uint16(108),
		11:  uint16(108),
		12:  uint16(108),
		13:  uint16(108),
		14:  uint16(108),
		15:  uint16(108),
		16:  uint16(108),
		17:  uint16(108),
		18:  uint16(108),
		19:  uint16(108),
		20:  uint16(108),
		21:  uint16(108),
		22:  uint16(108),
		23:  uint16(108),
		24:  uint16(108),
		25:  uint16(108),
		26:  uint16(108),
		27:  uint16(108),
		28:  uint16(108),
		29:  uint16(108),
		30:  uint16(108),
		31:  uint16(108),
		32:  uint16(108),
		33:  uint16(108),
		34:  uint16(108),
		35:  uint16(108),
		36:  uint16(108),
		37:  uint16(108),
		38:  uint16(108),
		39:  uint16(108),
		40:  uint16(106),
		41:  uint16(108),
		42:  uint16(108),
		43:  uint16(108),
		44:  uint16(108),
		45:  uint16(108),
		46:  uint16(108),
		47:  uint16(108),
		48:  uint16(108),
		49:  uint16(108),
		50:  uint16(108),
		51:  uint16(108),
		52:  uint16(108),
		53:  uint16(108),
		54:  uint16(108),
		55:  uint16(108),
		56:  uint16(108),
		57:  uint16(108),
		58:  uint16(108),
		59:  uint16(108),
		60:  uint16(108),
		61:  uint16(108),
		62:  uint16(108),
		63:  uint16(108),
		64:  uint16(108),
		65:  uint16(108),
		66:  uint16(108),
		67:  uint16(108),
		68:  uint16(108),
		69:  uint16(108),
		70:  uint16(108),
		71:  uint16(108),
		72:  uint16(106),
		73:  uint16(106),
		74:  uint16(106),
		75:  uint16(106),
		76:  uint16(106),
		77:  uint16(106),
		78:  uint16(106),
		79:  uint16(106),
		80:  uint16(106),
		81:  uint16(106),
		82:  uint16(106),
		83:  uint16(106),
		84:  uint16(106),
		85:  uint16(106),
		86:  uint16(106),
		87:  uint16(106),
		88:  uint16(106),
		89:  uint16(106),
		90:  uint16(106),
		91:  uint16(106),
		92:  uint16(106),
		93:  uint16(106),
		94:  uint16(106),
		95:  uint16(106),
		96:  uint16(106),
		97:  uint16(106),
		98:  uint16(106),
		99:  uint16(106),
		100: uint16(106),
		101: uint16(106),
		102: uint16(106),
		103: uint16(106),
		104: uint16(106),
		105: uint16(106),
		106: uint16(106),
		107: uint16(106),
		108: uint16(106),
		109: uint16(106),
		110: uint16(106),
		111: uint16(106),
		112: uint16(106),
		113: uint16(106),
		114: uint16(106),
		115: uint16(106),
		116: uint16(106),
		117: uint16(106),
		118: uint16(106),
		119: uint16(106),
		120: uint16(106),
		121: uint16(106),
		122: uint16(106),
		123: uint16(106),
		124: uint16(106),
		125: uint16(106),
		126: uint16(106),
		127: uint16(106),
		128: uint16(106),
		129: uint16(106),
		130: uint16(106),
		131: uint16(106),
		132: uint16(106),
		133: uint16(106),
		134: uint16(106),
		135: uint16(106),
		136: uint16(108),
		137: uint16(108),
		138: uint16(108),
		139: uint16(108),
		140: uint16(108),
		141: uint16(108),
		142: uint16(108),
		143: uint16(108),
		144: uint16(108),
		145: uint16(108),
		146: uint16(108),
		147: uint16(108),
		148: uint16(108),
		149: uint16(108),
		150: uint16(108),
		151: uint16(108),
		152: uint16(108),
		153: uint16(108),
		154: uint16(108),
		155: uint16(108),
		156: uint16(108),
		157: uint16(108),
		158: uint16(108),
		159: uint16(108),
		160: uint16(108),
		161: uint16(108),
		162: uint16(108),
		163: uint16(108),
		164: uint16(108),
		165: uint16(108),
		166: uint16(108),
		167: uint16(108),
		168: uint16(108),
		169: uint16(108),
		170: uint16(108),
		171: uint16(108),
		172: uint16(108),
		173: uint16(108),
		174: uint16(108),
		175: uint16(108),
		176: uint16(108),
		177: uint16(108),
		178: uint16(108),
		179: uint16(108),
		180: uint16(108),
		181: uint16(108),
		182: uint16(108),
		183: uint16(108),
		184: uint16(108),
		185: uint16(108),
		186: uint16(108),
		187: uint16(108),
		188: uint16(108),
		189: uint16(108),
		190: uint16(108),
		191: uint16(108),
		192: uint16(108),
		193: uint16(108),
		194: uint16(108),
		195: uint16(108),
		196: uint16(108),
		197: uint16(108),
		198: uint16(108),
		199: uint16(108),
		200: uint16(106),
		201: uint16(106),
		202: uint16(106),
		203: uint16(106),
		204: uint16(106),
		205: uint16(106),
		206: uint16(106),
		207: uint16(106),
		208: uint16(106),
		209: uint16(106),
		210: uint16(106),
		211: uint16(106),
		212: uint16(106),
		213: uint16(106),
		214: uint16(106),
		215: uint16(106),
		216: uint16(106),
		217: uint16(106),
		218: uint16(106),
		219: uint16(106),
		220: uint16(106),
		221: uint16(106),
		222: uint16(106),
		223: uint16(106),
		224: uint16(106),
		225: uint16(106),
		226: uint16(106),
		227: uint16(106),
		228: uint16(106),
		229: uint16(106),
		230: uint16(106),
		231: uint16(106),
		232: uint16(106),
		233: uint16(106),
		234: uint16(106),
		235: uint16(106),
		236: uint16(106),
		237: uint16(106),
		238: uint16(106),
		239: uint16(106),
		240: uint16(106),
		241: uint16(106),
		242: uint16(106),
		243: uint16(106),
		244: uint16(106),
		245: uint16(106),
		246: uint16(106),
		247: uint16(106),
		248: uint16(106),
		249: uint16(106),
		250: uint16(106),
		251: uint16(106),
		252: uint16(106),
		253: uint16(106),
		254: uint16(106),
		255: uint16(106),
		256: uint16(106),
		257: uint16(106),
		258: uint16(106),
		259: uint16(106),
		260: uint16(106),
		261: uint16(106),
		262: uint16(106),
		263: uint16(106),
		264: uint16(106),
		265: uint16(106),
		266: uint16(106),
		268: uint16(106),
		269: uint16(110),
		271: uint16(106),
		272: uint16(106),
		273: uint16(106),
		274: uint16(106),
		275: uint16(106),
		276: uint16(106),
		277: uint16(106),
		278: uint16(106),
		279: uint16(106),
		280: uint16(106),
		281: uint16(106),
		282: uint16(106),
		284: uint16(106),
		285: uint16(3),
	},
	12: {
		0:   uint16(112),
		1:   uint16(114),
		2:   uint16(112),
		5:   uint16(112),
		6:   uint16(112),
		8:   uint16(112),
		9:   uint16(114),
		10:  uint16(114),
		11:  uint16(114),
		12:  uint16(114),
		13:  uint16(114),
		14:  uint16(114),
		15:  uint16(114),
		16:  uint16(114),
		17:  uint16(114),
		18:  uint16(114),
		19:  uint16(114),
		20:  uint16(114),
		21:  uint16(114),
		22:  uint16(114),
		23:  uint16(114),
		24:  uint16(114),
		25:  uint16(114),
		26:  uint16(114),
		27:  uint16(114),
		28:  uint16(114),
		29:  uint16(114),
		30:  uint16(114),
		31:  uint16(114),
		32:  uint16(114),
		33:  uint16(114),
		34:  uint16(114),
		35:  uint16(114),
		36:  uint16(114),
		37:  uint16(114),
		38:  uint16(114),
		39:  uint16(114),
		40:  uint16(112),
		41:  uint16(114),
		42:  uint16(114),
		43:  uint16(114),
		44:  uint16(114),
		45:  uint16(114),
		46:  uint16(114),
		47:  uint16(114),
		48:  uint16(114),
		49:  uint16(114),
		50:  uint16(114),
		51:  uint16(114),
		52:  uint16(114),
		53:  uint16(114),
		54:  uint16(114),
		55:  uint16(114),
		56:  uint16(114),
		57:  uint16(114),
		58:  uint16(114),
		59:  uint16(114),
		60:  uint16(114),
		61:  uint16(114),
		62:  uint16(114),
		63:  uint16(114),
		64:  uint16(114),
		65:  uint16(114),
		66:  uint16(114),
		67:  uint16(114),
		68:  uint16(114),
		69:  uint16(114),
		70:  uint16(114),
		71:  uint16(114),
		72:  uint16(112),
		73:  uint16(112),
		74:  uint16(112),
		75:  uint16(112),
		76:  uint16(112),
		77:  uint16(112),
		78:  uint16(112),
		79:  uint16(112),
		80:  uint16(112),
		81:  uint16(112),
		82:  uint16(112),
		83:  uint16(112),
		84:  uint16(112),
		85:  uint16(112),
		86:  uint16(112),
		87:  uint16(112),
		88:  uint16(112),
		89:  uint16(112),
		90:  uint16(112),
		91:  uint16(112),
		92:  uint16(112),
		93:  uint16(112),
		94:  uint16(112),
		95:  uint16(112),
		96:  uint16(112),
		97:  uint16(112),
		98:  uint16(112),
		99:  uint16(112),
		100: uint16(112),
		101: uint16(112),
		102: uint16(112),
		103: uint16(112),
		104: uint16(112),
		105: uint16(112),
		106: uint16(112),
		107: uint16(112),
		108: uint16(112),
		109: uint16(112),
		110: uint16(112),
		111: uint16(112),
		112: uint16(112),
		113: uint16(112),
		114: uint16(112),
		115: uint16(112),
		116: uint16(112),
		117: uint16(112),
		118: uint16(112),
		119: uint16(112),
		120: uint16(112),
		121: uint16(112),
		122: uint16(112),
		123: uint16(112),
		124: uint16(112),
		125: uint16(112),
		126: uint16(112),
		127: uint16(112),
		128: uint16(112),
		129: uint16(112),
		130: uint16(112),
		131: uint16(112),
		132: uint16(112),
		133: uint16(112),
		134: uint16(112),
		135: uint16(112),
		136: uint16(114),
		137: uint16(114),
		138: uint16(114),
		139: uint16(114),
		140: uint16(114),
		141: uint16(114),
		142: uint16(114),
		143: uint16(114),
		144: uint16(114),
		145: uint16(114),
		146: uint16(114),
		147: uint16(114),
		148: uint16(114),
		149: uint16(114),
		150: uint16(114),
		151: uint16(114),
		152: uint16(114),
		153: uint16(114),
		154: uint16(114),
		155: uint16(114),
		156: uint16(114),
		157: uint16(114),
		158: uint16(114),
		159: uint16(114),
		160: uint16(114),
		161: uint16(114),
		162: uint16(114),
		163: uint16(114),
		164: uint16(114),
		165: uint16(114),
		166: uint16(114),
		167: uint16(114),
		168: uint16(114),
		169: uint16(114),
		170: uint16(114),
		171: uint16(114),
		172: uint16(114),
		173: uint16(114),
		174: uint16(114),
		175: uint16(114),
		176: uint16(114),
		177: uint16(114),
		178: uint16(114),
		179: uint16(114),
		180: uint16(114),
		181: uint16(114),
		182: uint16(114),
		183: uint16(114),
		184: uint16(114),
		185: uint16(114),
		186: uint16(114),
		187: uint16(114),
		188: uint16(114),
		189: uint16(114),
		190: uint16(114),
		191: uint16(114),
		192: uint16(114),
		193: uint16(114),
		194: uint16(114),
		195: uint16(114),
		196: uint16(114),
		197: uint16(114),
		198: uint16(114),
		199: uint16(114),
		200: uint16(112),
		201: uint16(112),
		202: uint16(112),
		203: uint16(112),
		204: uint16(112),
		205: uint16(112),
		206: uint16(112),
		207: uint16(112),
		208: uint16(112),
		209: uint16(112),
		210: uint16(112),
		211: uint16(112),
		212: uint16(112),
		213: uint16(112),
		214: uint16(112),
		215: uint16(112),
		216: uint16(112),
		217: uint16(112),
		218: uint16(112),
		219: uint16(112),
		220: uint16(112),
		221: uint16(112),
		222: uint16(112),
		223: uint16(112),
		224: uint16(112),
		225: uint16(112),
		226: uint16(112),
		227: uint16(112),
		228: uint16(112),
		229: uint16(112),
		230: uint16(112),
		231: uint16(112),
		232: uint16(112),
		233: uint16(112),
		234: uint16(112),
		235: uint16(112),
		236: uint16(112),
		237: uint16(112),
		238: uint16(112),
		239: uint16(112),
		240: uint16(112),
		241: uint16(112),
		242: uint16(112),
		243: uint16(112),
		244: uint16(112),
		245: uint16(112),
		246: uint16(112),
		247: uint16(112),
		248: uint16(112),
		249: uint16(112),
		250: uint16(112),
		251: uint16(112),
		252: uint16(112),
		253: uint16(112),
		254: uint16(112),
		255: uint16(112),
		256: uint16(112),
		257: uint16(112),
		258: uint16(112),
		259: uint16(112),
		260: uint16(112),
		261: uint16(112),
		262: uint16(112),
		263: uint16(112),
		264: uint16(112),
		265: uint16(112),
		266: uint16(112),
		268: uint16(112),
		269: uint16(110),
		271: uint16(112),
		272: uint16(112),
		273: uint16(112),
		274: uint16(112),
		275: uint16(112),
		276: uint16(112),
		277: uint16(112),
		278: uint16(112),
		279: uint16(112),
		280: uint16(112),
		281: uint16(112),
		282: uint16(112),
		284: uint16(112),
		285: uint16(3),
	},
	13: {
		0:   uint16(116),
		1:   uint16(118),
		2:   uint16(116),
		5:   uint16(116),
		6:   uint16(116),
		8:   uint16(116),
		9:   uint16(118),
		10:  uint16(118),
		11:  uint16(118),
		12:  uint16(118),
		13:  uint16(118),
		14:  uint16(118),
		15:  uint16(118),
		16:  uint16(118),
		17:  uint16(118),
		18:  uint16(118),
		19:  uint16(118),
		20:  uint16(118),
		21:  uint16(118),
		22:  uint16(118),
		23:  uint16(118),
		24:  uint16(118),
		25:  uint16(118),
		26:  uint16(118),
		27:  uint16(118),
		28:  uint16(118),
		29:  uint16(118),
		30:  uint16(118),
		31:  uint16(118),
		32:  uint16(118),
		33:  uint16(118),
		34:  uint16(118),
		35:  uint16(118),
		36:  uint16(118),
		37:  uint16(118),
		38:  uint16(118),
		39:  uint16(118),
		40:  uint16(116),
		41:  uint16(118),
		42:  uint16(118),
		43:  uint16(118),
		44:  uint16(118),
		45:  uint16(118),
		46:  uint16(118),
		47:  uint16(118),
		48:  uint16(118),
		49:  uint16(118),
		50:  uint16(118),
		51:  uint16(118),
		52:  uint16(118),
		53:  uint16(118),
		54:  uint16(118),
		55:  uint16(118),
		56:  uint16(118),
		57:  uint16(118),
		58:  uint16(118),
		59:  uint16(118),
		60:  uint16(118),
		61:  uint16(118),
		62:  uint16(118),
		63:  uint16(118),
		64:  uint16(118),
		65:  uint16(118),
		66:  uint16(118),
		67:  uint16(118),
		68:  uint16(118),
		69:  uint16(118),
		70:  uint16(118),
		71:  uint16(118),
		72:  uint16(116),
		73:  uint16(116),
		74:  uint16(116),
		75:  uint16(116),
		76:  uint16(116),
		77:  uint16(116),
		78:  uint16(116),
		79:  uint16(116),
		80:  uint16(116),
		81:  uint16(116),
		82:  uint16(116),
		83:  uint16(116),
		84:  uint16(116),
		85:  uint16(116),
		86:  uint16(116),
		87:  uint16(116),
		88:  uint16(116),
		89:  uint16(116),
		90:  uint16(116),
		91:  uint16(116),
		92:  uint16(116),
		93:  uint16(116),
		94:  uint16(116),
		95:  uint16(116),
		96:  uint16(116),
		97:  uint16(116),
		98:  uint16(116),
		99:  uint16(116),
		100: uint16(116),
		101: uint16(116),
		102: uint16(116),
		103: uint16(116),
		104: uint16(116),
		105: uint16(116),
		106: uint16(116),
		107: uint16(116),
		108: uint16(116),
		109: uint16(116),
		110: uint16(116),
		111: uint16(116),
		112: uint16(116),
		113: uint16(116),
		114: uint16(116),
		115: uint16(116),
		116: uint16(116),
		117: uint16(116),
		118: uint16(116),
		119: uint16(116),
		120: uint16(116),
		121: uint16(116),
		122: uint16(116),
		123: uint16(116),
		124: uint16(116),
		125: uint16(116),
		126: uint16(116),
		127: uint16(116),
		128: uint16(116),
		129: uint16(116),
		130: uint16(116),
		131: uint16(116),
		132: uint16(116),
		133: uint16(116),
		134: uint16(116),
		135: uint16(116),
		136: uint16(118),
		137: uint16(118),
		138: uint16(118),
		139: uint16(118),
		140: uint16(118),
		141: uint16(118),
		142: uint16(118),
		143: uint16(118),
		144: uint16(118),
		145: uint16(118),
		146: uint16(118),
		147: uint16(118),
		148: uint16(118),
		149: uint16(118),
		150: uint16(118),
		151: uint16(118),
		152: uint16(118),
		153: uint16(118),
		154: uint16(118),
		155: uint16(118),
		156: uint16(118),
		157: uint16(118),
		158: uint16(118),
		159: uint16(118),
		160: uint16(118),
		161: uint16(118),
		162: uint16(118),
		163: uint16(118),
		164: uint16(118),
		165: uint16(118),
		166: uint16(118),
		167: uint16(118),
		168: uint16(118),
		169: uint16(118),
		170: uint16(118),
		171: uint16(118),
		172: uint16(118),
		173: uint16(118),
		174: uint16(118),
		175: uint16(118),
		176: uint16(118),
		177: uint16(118),
		178: uint16(118),
		179: uint16(118),
		180: uint16(118),
		181: uint16(118),
		182: uint16(118),
		183: uint16(118),
		184: uint16(118),
		185: uint16(118),
		186: uint16(118),
		187: uint16(118),
		188: uint16(118),
		189: uint16(118),
		190: uint16(118),
		191: uint16(118),
		192: uint16(118),
		193: uint16(118),
		194: uint16(118),
		195: uint16(118),
		196: uint16(118),
		197: uint16(118),
		198: uint16(118),
		199: uint16(118),
		200: uint16(116),
		201: uint16(116),
		202: uint16(116),
		203: uint16(116),
		204: uint16(116),
		205: uint16(116),
		206: uint16(116),
		207: uint16(116),
		208: uint16(116),
		209: uint16(116),
		210: uint16(116),
		211: uint16(116),
		212: uint16(116),
		213: uint16(116),
		214: uint16(116),
		215: uint16(116),
		216: uint16(116),
		217: uint16(116),
		218: uint16(116),
		219: uint16(116),
		220: uint16(116),
		221: uint16(116),
		222: uint16(116),
		223: uint16(116),
		224: uint16(116),
		225: uint16(116),
		226: uint16(116),
		227: uint16(116),
		228: uint16(116),
		229: uint16(116),
		230: uint16(116),
		231: uint16(116),
		232: uint16(116),
		233: uint16(116),
		234: uint16(116),
		235: uint16(116),
		236: uint16(116),
		237: uint16(116),
		238: uint16(116),
		239: uint16(116),
		240: uint16(116),
		241: uint16(116),
		242: uint16(116),
		243: uint16(116),
		244: uint16(116),
		245: uint16(116),
		246: uint16(116),
		247: uint16(116),
		248: uint16(116),
		249: uint16(116),
		250: uint16(116),
		251: uint16(116),
		252: uint16(116),
		253: uint16(116),
		254: uint16(116),
		255: uint16(116),
		256: uint16(116),
		257: uint16(116),
		258: uint16(116),
		259: uint16(116),
		260: uint16(116),
		261: uint16(116),
		262: uint16(116),
		263: uint16(116),
		264: uint16(116),
		265: uint16(116),
		266: uint16(116),
		268: uint16(116),
		269: uint16(110),
		271: uint16(116),
		272: uint16(116),
		273: uint16(116),
		274: uint16(116),
		275: uint16(116),
		276: uint16(116),
		277: uint16(116),
		278: uint16(116),
		279: uint16(116),
		280: uint16(116),
		281: uint16(116),
		282: uint16(116),
		284: uint16(116),
		285: uint16(3),
	},
	14: {
		0:   uint16(112),
		1:   uint16(114),
		2:   uint16(112),
		5:   uint16(112),
		6:   uint16(112),
		8:   uint16(112),
		9:   uint16(114),
		10:  uint16(114),
		11:  uint16(114),
		12:  uint16(114),
		13:  uint16(114),
		14:  uint16(114),
		15:  uint16(114),
		16:  uint16(114),
		17:  uint16(114),
		18:  uint16(114),
		19:  uint16(114),
		20:  uint16(114),
		21:  uint16(114),
		22:  uint16(114),
		23:  uint16(114),
		24:  uint16(114),
		25:  uint16(114),
		26:  uint16(114),
		27:  uint16(114),
		28:  uint16(114),
		29:  uint16(114),
		30:  uint16(114),
		31:  uint16(114),
		32:  uint16(114),
		33:  uint16(114),
		34:  uint16(114),
		35:  uint16(114),
		36:  uint16(114),
		37:  uint16(114),
		38:  uint16(114),
		39:  uint16(114),
		40:  uint16(112),
		41:  uint16(114),
		42:  uint16(114),
		43:  uint16(114),
		44:  uint16(114),
		45:  uint16(114),
		46:  uint16(114),
		47:  uint16(114),
		48:  uint16(114),
		49:  uint16(114),
		50:  uint16(114),
		51:  uint16(114),
		52:  uint16(114),
		53:  uint16(114),
		54:  uint16(114),
		55:  uint16(114),
		56:  uint16(114),
		57:  uint16(114),
		58:  uint16(114),
		59:  uint16(114),
		60:  uint16(114),
		61:  uint16(114),
		62:  uint16(114),
		63:  uint16(114),
		64:  uint16(114),
		65:  uint16(114),
		66:  uint16(114),
		67:  uint16(114),
		68:  uint16(114),
		69:  uint16(114),
		70:  uint16(114),
		71:  uint16(114),
		72:  uint16(112),
		73:  uint16(112),
		74:  uint16(112),
		75:  uint16(112),
		76:  uint16(112),
		77:  uint16(112),
		78:  uint16(112),
		79:  uint16(112),
		80:  uint16(112),
		81:  uint16(112),
		82:  uint16(112),
		83:  uint16(112),
		84:  uint16(112),
		85:  uint16(112),
		86:  uint16(112),
		87:  uint16(112),
		88:  uint16(112),
		89:  uint16(112),
		90:  uint16(112),
		91:  uint16(112),
		92:  uint16(112),
		93:  uint16(112),
		94:  uint16(112),
		95:  uint16(112),
		96:  uint16(112),
		97:  uint16(112),
		98:  uint16(112),
		99:  uint16(112),
		100: uint16(112),
		101: uint16(112),
		102: uint16(112),
		103: uint16(112),
		104: uint16(112),
		105: uint16(112),
		106: uint16(112),
		107: uint16(112),
		108: uint16(112),
		109: uint16(112),
		110: uint16(112),
		111: uint16(112),
		112: uint16(112),
		113: uint16(112),
		114: uint16(112),
		115: uint16(112),
		116: uint16(112),
		117: uint16(112),
		118: uint16(112),
		119: uint16(112),
		120: uint16(112),
		121: uint16(112),
		122: uint16(112),
		123: uint16(112),
		124: uint16(112),
		125: uint16(112),
		126: uint16(112),
		127: uint16(112),
		128: uint16(112),
		129: uint16(112),
		130: uint16(112),
		131: uint16(112),
		132: uint16(112),
		133: uint16(112),
		134: uint16(112),
		135: uint16(112),
		136: uint16(114),
		137: uint16(114),
		138: uint16(114),
		139: uint16(114),
		140: uint16(114),
		141: uint16(114),
		142: uint16(114),
		143: uint16(114),
		144: uint16(114),
		145: uint16(114),
		146: uint16(114),
		147: uint16(114),
		148: uint16(114),
		149: uint16(114),
		150: uint16(114),
		151: uint16(114),
		152: uint16(114),
		153: uint16(114),
		154: uint16(114),
		155: uint16(114),
		156: uint16(114),
		157: uint16(114),
		158: uint16(114),
		159: uint16(114),
		160: uint16(114),
		161: uint16(114),
		162: uint16(114),
		163: uint16(114),
		164: uint16(114),
		165: uint16(114),
		166: uint16(114),
		167: uint16(114),
		168: uint16(114),
		169: uint16(114),
		170: uint16(114),
		171: uint16(114),
		172: uint16(114),
		173: uint16(114),
		174: uint16(114),
		175: uint16(114),
		176: uint16(114),
		177: uint16(114),
		178: uint16(114),
		179: uint16(114),
		180: uint16(114),
		181: uint16(114),
		182: uint16(114),
		183: uint16(114),
		184: uint16(114),
		185: uint16(114),
		186: uint16(114),
		187: uint16(114),
		188: uint16(114),
		189: uint16(114),
		190: uint16(114),
		191: uint16(114),
		192: uint16(114),
		193: uint16(114),
		194: uint16(114),
		195: uint16(114),
		196: uint16(114),
		197: uint16(114),
		198: uint16(114),
		199: uint16(114),
		200: uint16(112),
		201: uint16(112),
		202: uint16(112),
		203: uint16(112),
		204: uint16(112),
		205: uint16(112),
		206: uint16(112),
		207: uint16(112),
		208: uint16(112),
		209: uint16(112),
		210: uint16(112),
		211: uint16(112),
		212: uint16(112),
		213: uint16(112),
		214: uint16(112),
		215: uint16(112),
		216: uint16(112),
		217: uint16(112),
		218: uint16(112),
		219: uint16(112),
		220: uint16(112),
		221: uint16(112),
		222: uint16(112),
		223: uint16(112),
		224: uint16(112),
		225: uint16(112),
		226: uint16(112),
		227: uint16(112),
		228: uint16(112),
		229: uint16(112),
		230: uint16(112),
		231: uint16(112),
		232: uint16(112),
		233: uint16(112),
		234: uint16(112),
		235: uint16(112),
		236: uint16(112),
		237: uint16(112),
		238: uint16(112),
		239: uint16(112),
		240: uint16(112),
		241: uint16(112),
		242: uint16(112),
		243: uint16(112),
		244: uint16(112),
		245: uint16(112),
		246: uint16(112),
		247: uint16(112),
		248: uint16(112),
		249: uint16(112),
		250: uint16(112),
		251: uint16(112),
		252: uint16(112),
		253: uint16(112),
		254: uint16(112),
		255: uint16(112),
		256: uint16(112),
		257: uint16(112),
		258: uint16(112),
		259: uint16(112),
		260: uint16(112),
		261: uint16(112),
		262: uint16(112),
		263: uint16(112),
		264: uint16(112),
		265: uint16(112),
		266: uint16(112),
		268: uint16(112),
		271: uint16(112),
		272: uint16(112),
		273: uint16(112),
		274: uint16(112),
		275: uint16(112),
		276: uint16(112),
		277: uint16(112),
		278: uint16(112),
		279: uint16(112),
		280: uint16(112),
		281: uint16(112),
		282: uint16(112),
		284: uint16(112),
		285: uint16(3),
	},
	15: {
		0:   uint16(120),
		1:   uint16(122),
		2:   uint16(120),
		5:   uint16(120),
		6:   uint16(120),
		8:   uint16(120),
		9:   uint16(122),
		10:  uint16(122),
		11:  uint16(122),
		12:  uint16(122),
		13:  uint16(122),
		14:  uint16(122),
		15:  uint16(122),
		16:  uint16(122),
		17:  uint16(122),
		18:  uint16(122),
		19:  uint16(122),
		20:  uint16(122),
		21:  uint16(122),
		22:  uint16(122),
		23:  uint16(122),
		24:  uint16(122),
		25:  uint16(122),
		26:  uint16(122),
		27:  uint16(122),
		28:  uint16(122),
		29:  uint16(122),
		30:  uint16(122),
		31:  uint16(122),
		32:  uint16(122),
		33:  uint16(122),
		34:  uint16(122),
		35:  uint16(122),
		36:  uint16(122),
		37:  uint16(122),
		38:  uint16(122),
		39:  uint16(122),
		40:  uint16(120),
		41:  uint16(122),
		42:  uint16(122),
		43:  uint16(122),
		44:  uint16(122),
		45:  uint16(122),
		46:  uint16(122),
		47:  uint16(122),
		48:  uint16(122),
		49:  uint16(122),
		50:  uint16(122),
		51:  uint16(122),
		52:  uint16(122),
		53:  uint16(122),
		54:  uint16(122),
		55:  uint16(122),
		56:  uint16(122),
		57:  uint16(122),
		58:  uint16(122),
		59:  uint16(122),
		60:  uint16(122),
		61:  uint16(122),
		62:  uint16(122),
		63:  uint16(122),
		64:  uint16(122),
		65:  uint16(122),
		66:  uint16(122),
		67:  uint16(122),
		68:  uint16(122),
		69:  uint16(122),
		70:  uint16(122),
		71:  uint16(122),
		72:  uint16(120),
		73:  uint16(120),
		74:  uint16(120),
		75:  uint16(120),
		76:  uint16(120),
		77:  uint16(120),
		78:  uint16(120),
		79:  uint16(120),
		80:  uint16(120),
		81:  uint16(120),
		82:  uint16(120),
		83:  uint16(120),
		84:  uint16(120),
		85:  uint16(120),
		86:  uint16(120),
		87:  uint16(120),
		88:  uint16(120),
		89:  uint16(120),
		90:  uint16(120),
		91:  uint16(120),
		92:  uint16(120),
		93:  uint16(120),
		94:  uint16(120),
		95:  uint16(120),
		96:  uint16(120),
		97:  uint16(120),
		98:  uint16(120),
		99:  uint16(120),
		100: uint16(120),
		101: uint16(120),
		102: uint16(120),
		103: uint16(120),
		104: uint16(120),
		105: uint16(120),
		106: uint16(120),
		107: uint16(120),
		108: uint16(120),
		109: uint16(120),
		110: uint16(120),
		111: uint16(120),
		112: uint16(120),
		113: uint16(120),
		114: uint16(120),
		115: uint16(120),
		116: uint16(120),
		117: uint16(120),
		118: uint16(120),
		119: uint16(120),
		120: uint16(120),
		121: uint16(120),
		122: uint16(120),
		123: uint16(120),
		124: uint16(120),
		125: uint16(120),
		126: uint16(120),
		127: uint16(120),
		128: uint16(120),
		129: uint16(120),
		130: uint16(120),
		131: uint16(120),
		132: uint16(120),
		133: uint16(120),
		134: uint16(120),
		135: uint16(120),
		136: uint16(122),
		137: uint16(122),
		138: uint16(122),
		139: uint16(122),
		140: uint16(122),
		141: uint16(122),
		142: uint16(122),
		143: uint16(122),
		144: uint16(122),
		145: uint16(122),
		146: uint16(122),
		147: uint16(122),
		148: uint16(122),
		149: uint16(122),
		150: uint16(122),
		151: uint16(122),
		152: uint16(122),
		153: uint16(122),
		154: uint16(122),
		155: uint16(122),
		156: uint16(122),
		157: uint16(122),
		158: uint16(122),
		159: uint16(122),
		160: uint16(122),
		161: uint16(122),
		162: uint16(122),
		163: uint16(122),
		164: uint16(122),
		165: uint16(122),
		166: uint16(122),
		167: uint16(122),
		168: uint16(122),
		169: uint16(122),
		170: uint16(122),
		171: uint16(122),
		172: uint16(122),
		173: uint16(122),
		174: uint16(122),
		175: uint16(122),
		176: uint16(122),
		177: uint16(122),
		178: uint16(122),
		179: uint16(122),
		180: uint16(122),
		181: uint16(122),
		182: uint16(122),
		183: uint16(122),
		184: uint16(122),
		185: uint16(122),
		186: uint16(122),
		187: uint16(122),
		188: uint16(122),
		189: uint16(122),
		190: uint16(122),
		191: uint16(122),
		192: uint16(122),
		193: uint16(122),
		194: uint16(122),
		195: uint16(122),
		196: uint16(122),
		197: uint16(122),
		198: uint16(122),
		199: uint16(122),
		200: uint16(120),
		201: uint16(120),
		202: uint16(120),
		203: uint16(120),
		204: uint16(120),
		205: uint16(120),
		206: uint16(120),
		207: uint16(120),
		208: uint16(120),
		209: uint16(120),
		210: uint16(120),
		211: uint16(120),
		212: uint16(120),
		213: uint16(120),
		214: uint16(120),
		215: uint16(120),
		216: uint16(120),
		217: uint16(120),
		218: uint16(120),
		219: uint16(120),
		220: uint16(120),
		221: uint16(120),
		222: uint16(120),
		223: uint16(120),
		224: uint16(120),
		225: uint16(120),
		226: uint16(120),
		227: uint16(120),
		228: uint16(120),
		229: uint16(120),
		230: uint16(120),
		231: uint16(120),
		232: uint16(120),
		233: uint16(120),
		234: uint16(120),
		235: uint16(120),
		236: uint16(120),
		237: uint16(120),
		238: uint16(120),
		239: uint16(120),
		240: uint16(120),
		241: uint16(120),
		242: uint16(120),
		243: uint16(120),
		244: uint16(120),
		245: uint16(120),
		246: uint16(120),
		247: uint16(120),
		248: uint16(120),
		249: uint16(120),
		250: uint16(120),
		251: uint16(120),
		252: uint16(120),
		253: uint16(120),
		254: uint16(120),
		255: uint16(120),
		256: uint16(120),
		257: uint16(120),
		258: uint16(120),
		259: uint16(120),
		260: uint16(120),
		261: uint16(120),
		262: uint16(120),
		263: uint16(120),
		264: uint16(120),
		265: uint16(120),
		266: uint16(120),
		268: uint16(120),
		271: uint16(120),
		272: uint16(120),
		273: uint16(120),
		274: uint16(120),
		275: uint16(120),
		276: uint16(120),
		277: uint16(120),
		278: uint16(120),
		279: uint16(120),
		280: uint16(120),
		281: uint16(120),
		282: uint16(120),
		284: uint16(120),
		285: uint16(3),
	},
	16: {
		0:   uint16(124),
		1:   uint16(126),
		2:   uint16(124),
		5:   uint16(124),
		6:   uint16(124),
		8:   uint16(124),
		9:   uint16(126),
		10:  uint16(126),
		11:  uint16(126),
		12:  uint16(126),
		13:  uint16(126),
		14:  uint16(126),
		15:  uint16(126),
		16:  uint16(126),
		17:  uint16(126),
		18:  uint16(126),
		19:  uint16(126),
		20:  uint16(126),
		21:  uint16(126),
		22:  uint16(126),
		23:  uint16(126),
		24:  uint16(126),
		25:  uint16(126),
		26:  uint16(126),
		27:  uint16(126),
		28:  uint16(126),
		29:  uint16(126),
		30:  uint16(126),
		31:  uint16(126),
		32:  uint16(126),
		33:  uint16(126),
		34:  uint16(126),
		35:  uint16(126),
		36:  uint16(126),
		37:  uint16(126),
		38:  uint16(126),
		39:  uint16(126),
		40:  uint16(124),
		41:  uint16(126),
		42:  uint16(126),
		43:  uint16(126),
		44:  uint16(126),
		45:  uint16(126),
		46:  uint16(126),
		47:  uint16(126),
		48:  uint16(126),
		49:  uint16(126),
		50:  uint16(126),
		51:  uint16(126),
		52:  uint16(126),
		53:  uint16(126),
		54:  uint16(126),
		55:  uint16(126),
		56:  uint16(126),
		57:  uint16(126),
		58:  uint16(126),
		59:  uint16(126),
		60:  uint16(126),
		61:  uint16(126),
		62:  uint16(126),
		63:  uint16(126),
		64:  uint16(126),
		65:  uint16(126),
		66:  uint16(126),
		67:  uint16(126),
		68:  uint16(126),
		69:  uint16(126),
		70:  uint16(126),
		71:  uint16(126),
		72:  uint16(124),
		73:  uint16(124),
		74:  uint16(124),
		75:  uint16(124),
		76:  uint16(124),
		77:  uint16(124),
		78:  uint16(124),
		79:  uint16(124),
		80:  uint16(124),
		81:  uint16(124),
		82:  uint16(124),
		83:  uint16(124),
		84:  uint16(124),
		85:  uint16(124),
		86:  uint16(124),
		87:  uint16(124),
		88:  uint16(124),
		89:  uint16(124),
		90:  uint16(124),
		91:  uint16(124),
		92:  uint16(124),
		93:  uint16(124),
		94:  uint16(124),
		95:  uint16(124),
		96:  uint16(124),
		97:  uint16(124),
		98:  uint16(124),
		99:  uint16(124),
		100: uint16(124),
		101: uint16(124),
		102: uint16(124),
		103: uint16(124),
		104: uint16(124),
		105: uint16(124),
		106: uint16(124),
		107: uint16(124),
		108: uint16(124),
		109: uint16(124),
		110: uint16(124),
		111: uint16(124),
		112: uint16(124),
		113: uint16(124),
		114: uint16(124),
		115: uint16(124),
		116: uint16(124),
		117: uint16(124),
		118: uint16(124),
		119: uint16(124),
		120: uint16(124),
		121: uint16(124),
		122: uint16(124),
		123: uint16(124),
		124: uint16(124),
		125: uint16(124),
		126: uint16(124),
		127: uint16(124),
		128: uint16(124),
		129: uint16(124),
		130: uint16(124),
		131: uint16(124),
		132: uint16(124),
		133: uint16(124),
		134: uint16(124),
		135: uint16(124),
		136: uint16(126),
		137: uint16(126),
		138: uint16(126),
		139: uint16(126),
		140: uint16(126),
		141: uint16(126),
		142: uint16(126),
		143: uint16(126),
		144: uint16(126),
		145: uint16(126),
		146: uint16(126),
		147: uint16(126),
		148: uint16(126),
		149: uint16(126),
		150: uint16(126),
		151: uint16(126),
		152: uint16(126),
		153: uint16(126),
		154: uint16(126),
		155: uint16(126),
		156: uint16(126),
		157: uint16(126),
		158: uint16(126),
		159: uint16(126),
		160: uint16(126),
		161: uint16(126),
		162: uint16(126),
		163: uint16(126),
		164: uint16(126),
		165: uint16(126),
		166: uint16(126),
		167: uint16(126),
		168: uint16(126),
		169: uint16(126),
		170: uint16(126),
		171: uint16(126),
		172: uint16(126),
		173: uint16(126),
		174: uint16(126),
		175: uint16(126),
		176: uint16(126),
		177: uint16(126),
		178: uint16(126),
		179: uint16(126),
		180: uint16(126),
		181: uint16(126),
		182: uint16(126),
		183: uint16(126),
		184: uint16(126),
		185: uint16(126),
		186: uint16(126),
		187: uint16(126),
		188: uint16(126),
		189: uint16(126),
		190: uint16(126),
		191: uint16(126),
		192: uint16(126),
		193: uint16(126),
		194: uint16(126),
		195: uint16(126),
		196: uint16(126),
		197: uint16(126),
		198: uint16(126),
		199: uint16(126),
		200: uint16(124),
		201: uint16(124),
		202: uint16(124),
		203: uint16(124),
		204: uint16(124),
		205: uint16(124),
		206: uint16(124),
		207: uint16(124),
		208: uint16(124),
		209: uint16(124),
		210: uint16(124),
		211: uint16(124),
		212: uint16(124),
		213: uint16(124),
		214: uint16(124),
		215: uint16(124),
		216: uint16(124),
		217: uint16(124),
		218: uint16(124),
		219: uint16(124),
		220: uint16(124),
		221: uint16(124),
		222: uint16(124),
		223: uint16(124),
		224: uint16(124),
		225: uint16(124),
		226: uint16(124),
		227: uint16(124),
		228: uint16(124),
		229: uint16(124),
		230: uint16(124),
		231: uint16(124),
		232: uint16(124),
		233: uint16(124),
		234: uint16(124),
		235: uint16(124),
		236: uint16(124),
		237: uint16(124),
		238: uint16(124),
		239: uint16(124),
		240: uint16(124),
		241: uint16(124),
		242: uint16(124),
		243: uint16(124),
		244: uint16(124),
		245: uint16(124),
		246: uint16(124),
		247: uint16(124),
		248: uint16(124),
		249: uint16(124),
		250: uint16(124),
		251: uint16(124),
		252: uint16(124),
		253: uint16(124),
		254: uint16(124),
		255: uint16(124),
		256: uint16(124),
		257: uint16(124),
		258: uint16(124),
		259: uint16(124),
		260: uint16(124),
		261: uint16(124),
		262: uint16(124),
		263: uint16(124),
		264: uint16(124),
		265: uint16(124),
		266: uint16(124),
		268: uint16(124),
		271: uint16(124),
		272: uint16(124),
		273: uint16(124),
		274: uint16(124),
		275: uint16(124),
		276: uint16(124),
		277: uint16(124),
		278: uint16(124),
		279: uint16(124),
		280: uint16(124),
		281: uint16(124),
		282: uint16(124),
		284: uint16(124),
		285: uint16(3),
	},
	17: {
		0:   uint16(128),
		1:   uint16(130),
		2:   uint16(128),
		5:   uint16(128),
		6:   uint16(128),
		8:   uint16(128),
		9:   uint16(130),
		10:  uint16(130),
		11:  uint16(130),
		12:  uint16(130),
		13:  uint16(130),
		14:  uint16(130),
		15:  uint16(130),
		16:  uint16(130),
		17:  uint16(130),
		18:  uint16(130),
		19:  uint16(130),
		20:  uint16(130),
		21:  uint16(130),
		22:  uint16(130),
		23:  uint16(130),
		24:  uint16(130),
		25:  uint16(130),
		26:  uint16(130),
		27:  uint16(130),
		28:  uint16(130),
		29:  uint16(130),
		30:  uint16(130),
		31:  uint16(130),
		32:  uint16(130),
		33:  uint16(130),
		34:  uint16(130),
		35:  uint16(130),
		36:  uint16(130),
		37:  uint16(130),
		38:  uint16(130),
		39:  uint16(130),
		40:  uint16(128),
		41:  uint16(130),
		42:  uint16(130),
		43:  uint16(130),
		44:  uint16(130),
		45:  uint16(130),
		46:  uint16(130),
		47:  uint16(130),
		48:  uint16(130),
		49:  uint16(130),
		50:  uint16(130),
		51:  uint16(130),
		52:  uint16(130),
		53:  uint16(130),
		54:  uint16(130),
		55:  uint16(130),
		56:  uint16(130),
		57:  uint16(130),
		58:  uint16(130),
		59:  uint16(130),
		60:  uint16(130),
		61:  uint16(130),
		62:  uint16(130),
		63:  uint16(130),
		64:  uint16(130),
		65:  uint16(130),
		66:  uint16(130),
		67:  uint16(130),
		68:  uint16(130),
		69:  uint16(130),
		70:  uint16(130),
		71:  uint16(130),
		72:  uint16(128),
		73:  uint16(128),
		74:  uint16(128),
		75:  uint16(128),
		76:  uint16(128),
		77:  uint16(128),
		78:  uint16(128),
		79:  uint16(128),
		80:  uint16(128),
		81:  uint16(128),
		82:  uint16(128),
		83:  uint16(128),
		84:  uint16(128),
		85:  uint16(128),
		86:  uint16(128),
		87:  uint16(128),
		88:  uint16(128),
		89:  uint16(128),
		90:  uint16(128),
		91:  uint16(128),
		92:  uint16(128),
		93:  uint16(128),
		94:  uint16(128),
		95:  uint16(128),
		96:  uint16(128),
		97:  uint16(128),
		98:  uint16(128),
		99:  uint16(128),
		100: uint16(128),
		101: uint16(128),
		102: uint16(128),
		103: uint16(128),
		104: uint16(128),
		105: uint16(128),
		106: uint16(128),
		107: uint16(128),
		108: uint16(128),
		109: uint16(128),
		110: uint16(128),
		111: uint16(128),
		112: uint16(128),
		113: uint16(128),
		114: uint16(128),
		115: uint16(128),
		116: uint16(128),
		117: uint16(128),
		118: uint16(128),
		119: uint16(128),
		120: uint16(128),
		121: uint16(128),
		122: uint16(128),
		123: uint16(128),
		124: uint16(128),
		125: uint16(128),
		126: uint16(128),
		127: uint16(128),
		128: uint16(128),
		129: uint16(128),
		130: uint16(128),
		131: uint16(128),
		132: uint16(128),
		133: uint16(128),
		134: uint16(128),
		135: uint16(128),
		136: uint16(130),
		137: uint16(130),
		138: uint16(130),
		139: uint16(130),
		140: uint16(130),
		141: uint16(130),
		142: uint16(130),
		143: uint16(130),
		144: uint16(130),
		145: uint16(130),
		146: uint16(130),
		147: uint16(130),
		148: uint16(130),
		149: uint16(130),
		150: uint16(130),
		151: uint16(130),
		152: uint16(130),
		153: uint16(130),
		154: uint16(130),
		155: uint16(130),
		156: uint16(130),
		157: uint16(130),
		158: uint16(130),
		159: uint16(130),
		160: uint16(130),
		161: uint16(130),
		162: uint16(130),
		163: uint16(130),
		164: uint16(130),
		165: uint16(130),
		166: uint16(130),
		167: uint16(130),
		168: uint16(130),
		169: uint16(130),
		170: uint16(130),
		171: uint16(130),
		172: uint16(130),
		173: uint16(130),
		174: uint16(130),
		175: uint16(130),
		176: uint16(130),
		177: uint16(130),
		178: uint16(130),
		179: uint16(130),
		180: uint16(130),
		181: uint16(130),
		182: uint16(130),
		183: uint16(130),
		184: uint16(130),
		185: uint16(130),
		186: uint16(130),
		187: uint16(130),
		188: uint16(130),
		189: uint16(130),
		190: uint16(130),
		191: uint16(130),
		192: uint16(130),
		193: uint16(130),
		194: uint16(130),
		195: uint16(130),
		196: uint16(130),
		197: uint16(130),
		198: uint16(130),
		199: uint16(130),
		200: uint16(128),
		201: uint16(128),
		202: uint16(128),
		203: uint16(128),
		204: uint16(128),
		205: uint16(128),
		206: uint16(128),
		207: uint16(128),
		208: uint16(128),
		209: uint16(128),
		210: uint16(128),
		211: uint16(128),
		212: uint16(128),
		213: uint16(128),
		214: uint16(128),
		215: uint16(128),
		216: uint16(128),
		217: uint16(128),
		218: uint16(128),
		219: uint16(128),
		220: uint16(128),
		221: uint16(128),
		222: uint16(128),
		223: uint16(128),
		224: uint16(128),
		225: uint16(128),
		226: uint16(128),
		227: uint16(128),
		228: uint16(128),
		229: uint16(128),
		230: uint16(128),
		231: uint16(128),
		232: uint16(128),
		233: uint16(128),
		234: uint16(128),
		235: uint16(128),
		236: uint16(128),
		237: uint16(128),
		238: uint16(128),
		239: uint16(128),
		240: uint16(128),
		241: uint16(128),
		242: uint16(128),
		243: uint16(128),
		244: uint16(128),
		245: uint16(128),
		246: uint16(128),
		247: uint16(128),
		248: uint16(128),
		249: uint16(128),
		250: uint16(128),
		251: uint16(128),
		252: uint16(128),
		253: uint16(128),
		254: uint16(128),
		255: uint16(128),
		256: uint16(128),
		257: uint16(128),
		258: uint16(128),
		259: uint16(128),
		260: uint16(128),
		261: uint16(128),
		262: uint16(128),
		263: uint16(128),
		264: uint16(128),
		265: uint16(128),
		266: uint16(128),
		268: uint16(128),
		271: uint16(128),
		272: uint16(128),
		273: uint16(128),
		274: uint16(128),
		275: uint16(128),
		276: uint16(128),
		277: uint16(128),
		278: uint16(128),
		279: uint16(128),
		280: uint16(128),
		281: uint16(128),
		282: uint16(128),
		284: uint16(128),
		285: uint16(3),
	},
	18: {
		0:   uint16(132),
		1:   uint16(134),
		2:   uint16(132),
		5:   uint16(132),
		6:   uint16(132),
		8:   uint16(132),
		9:   uint16(134),
		10:  uint16(134),
		11:  uint16(134),
		12:  uint16(134),
		13:  uint16(134),
		14:  uint16(134),
		15:  uint16(134),
		16:  uint16(134),
		17:  uint16(134),
		18:  uint16(134),
		19:  uint16(134),
		20:  uint16(134),
		21:  uint16(134),
		22:  uint16(134),
		23:  uint16(134),
		24:  uint16(134),
		25:  uint16(134),
		26:  uint16(134),
		27:  uint16(134),
		28:  uint16(134),
		29:  uint16(134),
		30:  uint16(134),
		31:  uint16(134),
		32:  uint16(134),
		33:  uint16(134),
		34:  uint16(134),
		35:  uint16(134),
		36:  uint16(134),
		37:  uint16(134),
		38:  uint16(134),
		39:  uint16(134),
		40:  uint16(132),
		41:  uint16(134),
		42:  uint16(134),
		43:  uint16(134),
		44:  uint16(134),
		45:  uint16(134),
		46:  uint16(134),
		47:  uint16(134),
		48:  uint16(134),
		49:  uint16(134),
		50:  uint16(134),
		51:  uint16(134),
		52:  uint16(134),
		53:  uint16(134),
		54:  uint16(134),
		55:  uint16(134),
		56:  uint16(134),
		57:  uint16(134),
		58:  uint16(134),
		59:  uint16(134),
		60:  uint16(134),
		61:  uint16(134),
		62:  uint16(134),
		63:  uint16(134),
		64:  uint16(134),
		65:  uint16(134),
		66:  uint16(134),
		67:  uint16(134),
		68:  uint16(134),
		69:  uint16(134),
		70:  uint16(134),
		71:  uint16(134),
		72:  uint16(132),
		73:  uint16(132),
		74:  uint16(132),
		75:  uint16(132),
		76:  uint16(132),
		77:  uint16(132),
		78:  uint16(132),
		79:  uint16(132),
		80:  uint16(132),
		81:  uint16(132),
		82:  uint16(132),
		83:  uint16(132),
		84:  uint16(132),
		85:  uint16(132),
		86:  uint16(132),
		87:  uint16(132),
		88:  uint16(132),
		89:  uint16(132),
		90:  uint16(132),
		91:  uint16(132),
		92:  uint16(132),
		93:  uint16(132),
		94:  uint16(132),
		95:  uint16(132),
		96:  uint16(132),
		97:  uint16(132),
		98:  uint16(132),
		99:  uint16(132),
		100: uint16(132),
		101: uint16(132),
		102: uint16(132),
		103: uint16(132),
		104: uint16(132),
		105: uint16(132),
		106: uint16(132),
		107: uint16(132),
		108: uint16(132),
		109: uint16(132),
		110: uint16(132),
		111: uint16(132),
		112: uint16(132),
		113: uint16(132),
		114: uint16(132),
		115: uint16(132),
		116: uint16(132),
		117: uint16(132),
		118: uint16(132),
		119: uint16(132),
		120: uint16(132),
		121: uint16(132),
		122: uint16(132),
		123: uint16(132),
		124: uint16(132),
		125: uint16(132),
		126: uint16(132),
		127: uint16(132),
		128: uint16(132),
		129: uint16(132),
		130: uint16(132),
		131: uint16(132),
		132: uint16(132),
		133: uint16(132),
		134: uint16(132),
		135: uint16(132),
		136: uint16(134),
		137: uint16(134),
		138: uint16(134),
		139: uint16(134),
		140: uint16(134),
		141: uint16(134),
		142: uint16(134),
		143: uint16(134),
		144: uint16(134),
		145: uint16(134),
		146: uint16(134),
		147: uint16(134),
		148: uint16(134),
		149: uint16(134),
		150: uint16(134),
		151: uint16(134),
		152: uint16(134),
		153: uint16(134),
		154: uint16(134),
		155: uint16(134),
		156: uint16(134),
		157: uint16(134),
		158: uint16(134),
		159: uint16(134),
		160: uint16(134),
		161: uint16(134),
		162: uint16(134),
		163: uint16(134),
		164: uint16(134),
		165: uint16(134),
		166: uint16(134),
		167: uint16(134),
		168: uint16(134),
		169: uint16(134),
		170: uint16(134),
		171: uint16(134),
		172: uint16(134),
		173: uint16(134),
		174: uint16(134),
		175: uint16(134),
		176: uint16(134),
		177: uint16(134),
		178: uint16(134),
		179: uint16(134),
		180: uint16(134),
		181: uint16(134),
		182: uint16(134),
		183: uint16(134),
		184: uint16(134),
		185: uint16(134),
		186: uint16(134),
		187: uint16(134),
		188: uint16(134),
		189: uint16(134),
		190: uint16(134),
		191: uint16(134),
		192: uint16(134),
		193: uint16(134),
		194: uint16(134),
		195: uint16(134),
		196: uint16(134),
		197: uint16(134),
		198: uint16(134),
		199: uint16(134),
		200: uint16(132),
		201: uint16(132),
		202: uint16(132),
		203: uint16(132),
		204: uint16(132),
		205: uint16(132),
		206: uint16(132),
		207: uint16(132),
		208: uint16(132),
		209: uint16(132),
		210: uint16(132),
		211: uint16(132),
		212: uint16(132),
		213: uint16(132),
		214: uint16(132),
		215: uint16(132),
		216: uint16(132),
		217: uint16(132),
		218: uint16(132),
		219: uint16(132),
		220: uint16(132),
		221: uint16(132),
		222: uint16(132),
		223: uint16(132),
		224: uint16(132),
		225: uint16(132),
		226: uint16(132),
		227: uint16(132),
		228: uint16(132),
		229: uint16(132),
		230: uint16(132),
		231: uint16(132),
		232: uint16(132),
		233: uint16(132),
		234: uint16(132),
		235: uint16(132),
		236: uint16(132),
		237: uint16(132),
		238: uint16(132),
		239: uint16(132),
		240: uint16(132),
		241: uint16(132),
		242: uint16(132),
		243: uint16(132),
		244: uint16(132),
		245: uint16(132),
		246: uint16(132),
		247: uint16(132),
		248: uint16(132),
		249: uint16(132),
		250: uint16(132),
		251: uint16(132),
		252: uint16(132),
		253: uint16(132),
		254: uint16(132),
		255: uint16(132),
		256: uint16(132),
		257: uint16(132),
		258: uint16(132),
		259: uint16(132),
		260: uint16(132),
		261: uint16(132),
		262: uint16(132),
		263: uint16(132),
		264: uint16(132),
		265: uint16(132),
		266: uint16(132),
		268: uint16(132),
		271: uint16(132),
		272: uint16(132),
		273: uint16(132),
		274: uint16(132),
		275: uint16(132),
		276: uint16(132),
		277: uint16(132),
		278: uint16(132),
		279: uint16(132),
		280: uint16(132),
		281: uint16(132),
		282: uint16(132),
		284: uint16(132),
		285: uint16(3),
	},
	19: {
		0:   uint16(136),
		1:   uint16(138),
		2:   uint16(136),
		5:   uint16(136),
		6:   uint16(136),
		8:   uint16(136),
		9:   uint16(138),
		10:  uint16(138),
		11:  uint16(138),
		12:  uint16(138),
		13:  uint16(138),
		14:  uint16(138),
		15:  uint16(138),
		16:  uint16(138),
		17:  uint16(138),
		18:  uint16(138),
		19:  uint16(138),
		20:  uint16(138),
		21:  uint16(138),
		22:  uint16(138),
		23:  uint16(138),
		24:  uint16(138),
		25:  uint16(138),
		26:  uint16(138),
		27:  uint16(138),
		28:  uint16(138),
		29:  uint16(138),
		30:  uint16(138),
		31:  uint16(138),
		32:  uint16(138),
		33:  uint16(138),
		34:  uint16(138),
		35:  uint16(138),
		36:  uint16(138),
		37:  uint16(138),
		38:  uint16(138),
		39:  uint16(138),
		40:  uint16(136),
		41:  uint16(138),
		42:  uint16(138),
		43:  uint16(138),
		44:  uint16(138),
		45:  uint16(138),
		46:  uint16(138),
		47:  uint16(138),
		48:  uint16(138),
		49:  uint16(138),
		50:  uint16(138),
		51:  uint16(138),
		52:  uint16(138),
		53:  uint16(138),
		54:  uint16(138),
		55:  uint16(138),
		56:  uint16(138),
		57:  uint16(138),
		58:  uint16(138),
		59:  uint16(138),
		60:  uint16(138),
		61:  uint16(138),
		62:  uint16(138),
		63:  uint16(138),
		64:  uint16(138),
		65:  uint16(138),
		66:  uint16(138),
		67:  uint16(138),
		68:  uint16(138),
		69:  uint16(138),
		70:  uint16(138),
		71:  uint16(138),
		72:  uint16(136),
		73:  uint16(136),
		74:  uint16(136),
		75:  uint16(136),
		76:  uint16(136),
		77:  uint16(136),
		78:  uint16(136),
		79:  uint16(136),
		80:  uint16(136),
		81:  uint16(136),
		82:  uint16(136),
		83:  uint16(136),
		84:  uint16(136),
		85:  uint16(136),
		86:  uint16(136),
		87:  uint16(136),
		88:  uint16(136),
		89:  uint16(136),
		90:  uint16(136),
		91:  uint16(136),
		92:  uint16(136),
		93:  uint16(136),
		94:  uint16(136),
		95:  uint16(136),
		96:  uint16(136),
		97:  uint16(136),
		98:  uint16(136),
		99:  uint16(136),
		100: uint16(136),
		101: uint16(136),
		102: uint16(136),
		103: uint16(136),
		104: uint16(136),
		105: uint16(136),
		106: uint16(136),
		107: uint16(136),
		108: uint16(136),
		109: uint16(136),
		110: uint16(136),
		111: uint16(136),
		112: uint16(136),
		113: uint16(136),
		114: uint16(136),
		115: uint16(136),
		116: uint16(136),
		117: uint16(136),
		118: uint16(136),
		119: uint16(136),
		120: uint16(136),
		121: uint16(136),
		122: uint16(136),
		123: uint16(136),
		124: uint16(136),
		125: uint16(136),
		126: uint16(136),
		127: uint16(136),
		128: uint16(136),
		129: uint16(136),
		130: uint16(136),
		131: uint16(136),
		132: uint16(136),
		133: uint16(136),
		134: uint16(136),
		135: uint16(136),
		136: uint16(138),
		137: uint16(138),
		138: uint16(138),
		139: uint16(138),
		140: uint16(138),
		141: uint16(138),
		142: uint16(138),
		143: uint16(138),
		144: uint16(138),
		145: uint16(138),
		146: uint16(138),
		147: uint16(138),
		148: uint16(138),
		149: uint16(138),
		150: uint16(138),
		151: uint16(138),
		152: uint16(138),
		153: uint16(138),
		154: uint16(138),
		155: uint16(138),
		156: uint16(138),
		157: uint16(138),
		158: uint16(138),
		159: uint16(138),
		160: uint16(138),
		161: uint16(138),
		162: uint16(138),
		163: uint16(138),
		164: uint16(138),
		165: uint16(138),
		166: uint16(138),
		167: uint16(138),
		168: uint16(138),
		169: uint16(138),
		170: uint16(138),
		171: uint16(138),
		172: uint16(138),
		173: uint16(138),
		174: uint16(138),
		175: uint16(138),
		176: uint16(138),
		177: uint16(138),
		178: uint16(138),
		179: uint16(138),
		180: uint16(138),
		181: uint16(138),
		182: uint16(138),
		183: uint16(138),
		184: uint16(138),
		185: uint16(138),
		186: uint16(138),
		187: uint16(138),
		188: uint16(138),
		189: uint16(138),
		190: uint16(138),
		191: uint16(138),
		192: uint16(138),
		193: uint16(138),
		194: uint16(138),
		195: uint16(138),
		196: uint16(138),
		197: uint16(138),
		198: uint16(138),
		199: uint16(138),
		200: uint16(136),
		201: uint16(136),
		202: uint16(136),
		203: uint16(136),
		204: uint16(136),
		205: uint16(136),
		206: uint16(136),
		207: uint16(136),
		208: uint16(136),
		209: uint16(136),
		210: uint16(136),
		211: uint16(136),
		212: uint16(136),
		213: uint16(136),
		214: uint16(136),
		215: uint16(136),
		216: uint16(136),
		217: uint16(136),
		218: uint16(136),
		219: uint16(136),
		220: uint16(136),
		221: uint16(136),
		222: uint16(136),
		223: uint16(136),
		224: uint16(136),
		225: uint16(136),
		226: uint16(136),
		227: uint16(136),
		228: uint16(136),
		229: uint16(136),
		230: uint16(136),
		231: uint16(136),
		232: uint16(136),
		233: uint16(136),
		234: uint16(136),
		235: uint16(136),
		236: uint16(136),
		237: uint16(136),
		238: uint16(136),
		239: uint16(136),
		240: uint16(136),
		241: uint16(136),
		242: uint16(136),
		243: uint16(136),
		244: uint16(136),
		245: uint16(136),
		246: uint16(136),
		247: uint16(136),
		248: uint16(136),
		249: uint16(136),
		250: uint16(136),
		251: uint16(136),
		252: uint16(136),
		253: uint16(136),
		254: uint16(136),
		255: uint16(136),
		256: uint16(136),
		257: uint16(136),
		258: uint16(136),
		259: uint16(136),
		260: uint16(136),
		261: uint16(136),
		262: uint16(136),
		263: uint16(136),
		264: uint16(136),
		265: uint16(136),
		266: uint16(136),
		268: uint16(136),
		271: uint16(136),
		272: uint16(136),
		273: uint16(136),
		274: uint16(136),
		275: uint16(136),
		276: uint16(136),
		277: uint16(136),
		278: uint16(136),
		279: uint16(136),
		280: uint16(136),
		281: uint16(136),
		282: uint16(136),
		284: uint16(136),
		285: uint16(3),
	},
	20: {
		0:   uint16(140),
		1:   uint16(142),
		2:   uint16(140),
		5:   uint16(140),
		6:   uint16(140),
		8:   uint16(140),
		9:   uint16(142),
		10:  uint16(142),
		11:  uint16(142),
		12:  uint16(142),
		13:  uint16(142),
		14:  uint16(142),
		15:  uint16(142),
		16:  uint16(142),
		17:  uint16(142),
		18:  uint16(142),
		19:  uint16(142),
		20:  uint16(142),
		21:  uint16(142),
		22:  uint16(142),
		23:  uint16(142),
		24:  uint16(142),
		25:  uint16(142),
		26:  uint16(142),
		27:  uint16(142),
		28:  uint16(142),
		29:  uint16(142),
		30:  uint16(142),
		31:  uint16(142),
		32:  uint16(142),
		33:  uint16(142),
		34:  uint16(142),
		35:  uint16(142),
		36:  uint16(142),
		37:  uint16(142),
		38:  uint16(142),
		39:  uint16(142),
		40:  uint16(140),
		41:  uint16(142),
		42:  uint16(142),
		43:  uint16(142),
		44:  uint16(142),
		45:  uint16(142),
		46:  uint16(142),
		47:  uint16(142),
		48:  uint16(142),
		49:  uint16(142),
		50:  uint16(142),
		51:  uint16(142),
		52:  uint16(142),
		53:  uint16(142),
		54:  uint16(142),
		55:  uint16(142),
		56:  uint16(142),
		57:  uint16(142),
		58:  uint16(142),
		59:  uint16(142),
		60:  uint16(142),
		61:  uint16(142),
		62:  uint16(142),
		63:  uint16(142),
		64:  uint16(142),
		65:  uint16(142),
		66:  uint16(142),
		67:  uint16(142),
		68:  uint16(142),
		69:  uint16(142),
		70:  uint16(142),
		71:  uint16(142),
		72:  uint16(140),
		73:  uint16(140),
		74:  uint16(140),
		75:  uint16(140),
		76:  uint16(140),
		77:  uint16(140),
		78:  uint16(140),
		79:  uint16(140),
		80:  uint16(140),
		81:  uint16(140),
		82:  uint16(140),
		83:  uint16(140),
		84:  uint16(140),
		85:  uint16(140),
		86:  uint16(140),
		87:  uint16(140),
		88:  uint16(140),
		89:  uint16(140),
		90:  uint16(140),
		91:  uint16(140),
		92:  uint16(140),
		93:  uint16(140),
		94:  uint16(140),
		95:  uint16(140),
		96:  uint16(140),
		97:  uint16(140),
		98:  uint16(140),
		99:  uint16(140),
		100: uint16(140),
		101: uint16(140),
		102: uint16(140),
		103: uint16(140),
		104: uint16(140),
		105: uint16(140),
		106: uint16(140),
		107: uint16(140),
		108: uint16(140),
		109: uint16(140),
		110: uint16(140),
		111: uint16(140),
		112: uint16(140),
		113: uint16(140),
		114: uint16(140),
		115: uint16(140),
		116: uint16(140),
		117: uint16(140),
		118: uint16(140),
		119: uint16(140),
		120: uint16(140),
		121: uint16(140),
		122: uint16(140),
		123: uint16(140),
		124: uint16(140),
		125: uint16(140),
		126: uint16(140),
		127: uint16(140),
		128: uint16(140),
		129: uint16(140),
		130: uint16(140),
		131: uint16(140),
		132: uint16(140),
		133: uint16(140),
		134: uint16(140),
		135: uint16(140),
		136: uint16(142),
		137: uint16(142),
		138: uint16(142),
		139: uint16(142),
		140: uint16(142),
		141: uint16(142),
		142: uint16(142),
		143: uint16(142),
		144: uint16(142),
		145: uint16(142),
		146: uint16(142),
		147: uint16(142),
		148: uint16(142),
		149: uint16(142),
		150: uint16(142),
		151: uint16(142),
		152: uint16(142),
		153: uint16(142),
		154: uint16(142),
		155: uint16(142),
		156: uint16(142),
		157: uint16(142),
		158: uint16(142),
		159: uint16(142),
		160: uint16(142),
		161: uint16(142),
		162: uint16(142),
		163: uint16(142),
		164: uint16(142),
		165: uint16(142),
		166: uint16(142),
		167: uint16(142),
		168: uint16(142),
		169: uint16(142),
		170: uint16(142),
		171: uint16(142),
		172: uint16(142),
		173: uint16(142),
		174: uint16(142),
		175: uint16(142),
		176: uint16(142),
		177: uint16(142),
		178: uint16(142),
		179: uint16(142),
		180: uint16(142),
		181: uint16(142),
		182: uint16(142),
		183: uint16(142),
		184: uint16(142),
		185: uint16(142),
		186: uint16(142),
		187: uint16(142),
		188: uint16(142),
		189: uint16(142),
		190: uint16(142),
		191: uint16(142),
		192: uint16(142),
		193: uint16(142),
		194: uint16(142),
		195: uint16(142),
		196: uint16(142),
		197: uint16(142),
		198: uint16(142),
		199: uint16(142),
		200: uint16(140),
		201: uint16(140),
		202: uint16(140),
		203: uint16(140),
		204: uint16(140),
		205: uint16(140),
		206: uint16(140),
		207: uint16(140),
		208: uint16(140),
		209: uint16(140),
		210: uint16(140),
		211: uint16(140),
		212: uint16(140),
		213: uint16(140),
		214: uint16(140),
		215: uint16(140),
		216: uint16(140),
		217: uint16(140),
		218: uint16(140),
		219: uint16(140),
		220: uint16(140),
		221: uint16(140),
		222: uint16(140),
		223: uint16(140),
		224: uint16(140),
		225: uint16(140),
		226: uint16(140),
		227: uint16(140),
		228: uint16(140),
		229: uint16(140),
		230: uint16(140),
		231: uint16(140),
		232: uint16(140),
		233: uint16(140),
		234: uint16(140),
		235: uint16(140),
		236: uint16(140),
		237: uint16(140),
		238: uint16(140),
		239: uint16(140),
		240: uint16(140),
		241: uint16(140),
		242: uint16(140),
		243: uint16(140),
		244: uint16(140),
		245: uint16(140),
		246: uint16(140),
		247: uint16(140),
		248: uint16(140),
		249: uint16(140),
		250: uint16(140),
		251: uint16(140),
		252: uint16(140),
		253: uint16(140),
		254: uint16(140),
		255: uint16(140),
		256: uint16(140),
		257: uint16(140),
		258: uint16(140),
		259: uint16(140),
		260: uint16(140),
		261: uint16(140),
		262: uint16(140),
		263: uint16(140),
		264: uint16(140),
		265: uint16(140),
		266: uint16(140),
		268: uint16(140),
		271: uint16(140),
		272: uint16(140),
		273: uint16(140),
		274: uint16(140),
		275: uint16(140),
		276: uint16(140),
		277: uint16(140),
		278: uint16(140),
		279: uint16(140),
		280: uint16(140),
		281: uint16(140),
		282: uint16(140),
		284: uint16(140),
		285: uint16(3),
	},
	21: {
		0:   uint16(144),
		1:   uint16(146),
		2:   uint16(144),
		5:   uint16(144),
		6:   uint16(144),
		8:   uint16(144),
		9:   uint16(146),
		10:  uint16(146),
		11:  uint16(146),
		12:  uint16(146),
		13:  uint16(146),
		14:  uint16(146),
		15:  uint16(146),
		16:  uint16(146),
		17:  uint16(146),
		18:  uint16(146),
		19:  uint16(146),
		20:  uint16(146),
		21:  uint16(146),
		22:  uint16(146),
		23:  uint16(146),
		24:  uint16(146),
		25:  uint16(146),
		26:  uint16(146),
		27:  uint16(146),
		28:  uint16(146),
		29:  uint16(146),
		30:  uint16(146),
		31:  uint16(146),
		32:  uint16(146),
		33:  uint16(146),
		34:  uint16(146),
		35:  uint16(146),
		36:  uint16(146),
		37:  uint16(146),
		38:  uint16(146),
		39:  uint16(146),
		40:  uint16(144),
		41:  uint16(146),
		42:  uint16(146),
		43:  uint16(146),
		44:  uint16(146),
		45:  uint16(146),
		46:  uint16(146),
		47:  uint16(146),
		48:  uint16(146),
		49:  uint16(146),
		50:  uint16(146),
		51:  uint16(146),
		52:  uint16(146),
		53:  uint16(146),
		54:  uint16(146),
		55:  uint16(146),
		56:  uint16(146),
		57:  uint16(146),
		58:  uint16(146),
		59:  uint16(146),
		60:  uint16(146),
		61:  uint16(146),
		62:  uint16(146),
		63:  uint16(146),
		64:  uint16(146),
		65:  uint16(146),
		66:  uint16(146),
		67:  uint16(146),
		68:  uint16(146),
		69:  uint16(146),
		70:  uint16(146),
		71:  uint16(146),
		72:  uint16(144),
		73:  uint16(144),
		74:  uint16(144),
		75:  uint16(144),
		76:  uint16(144),
		77:  uint16(144),
		78:  uint16(144),
		79:  uint16(144),
		80:  uint16(144),
		81:  uint16(144),
		82:  uint16(144),
		83:  uint16(144),
		84:  uint16(144),
		85:  uint16(144),
		86:  uint16(144),
		87:  uint16(144),
		88:  uint16(144),
		89:  uint16(144),
		90:  uint16(144),
		91:  uint16(144),
		92:  uint16(144),
		93:  uint16(144),
		94:  uint16(144),
		95:  uint16(144),
		96:  uint16(144),
		97:  uint16(144),
		98:  uint16(144),
		99:  uint16(144),
		100: uint16(144),
		101: uint16(144),
		102: uint16(144),
		103: uint16(144),
		104: uint16(144),
		105: uint16(144),
		106: uint16(144),
		107: uint16(144),
		108: uint16(144),
		109: uint16(144),
		110: uint16(144),
		111: uint16(144),
		112: uint16(144),
		113: uint16(144),
		114: uint16(144),
		115: uint16(144),
		116: uint16(144),
		117: uint16(144),
		118: uint16(144),
		119: uint16(144),
		120: uint16(144),
		121: uint16(144),
		122: uint16(144),
		123: uint16(144),
		124: uint16(144),
		125: uint16(144),
		126: uint16(144),
		127: uint16(144),
		128: uint16(144),
		129: uint16(144),
		130: uint16(144),
		131: uint16(144),
		132: uint16(144),
		133: uint16(144),
		134: uint16(144),
		135: uint16(144),
		136: uint16(146),
		137: uint16(146),
		138: uint16(146),
		139: uint16(146),
		140: uint16(146),
		141: uint16(146),
		142: uint16(146),
		143: uint16(146),
		144: uint16(146),
		145: uint16(146),
		146: uint16(146),
		147: uint16(146),
		148: uint16(146),
		149: uint16(146),
		150: uint16(146),
		151: uint16(146),
		152: uint16(146),
		153: uint16(146),
		154: uint16(146),
		155: uint16(146),
		156: uint16(146),
		157: uint16(146),
		158: uint16(146),
		159: uint16(146),
		160: uint16(146),
		161: uint16(146),
		162: uint16(146),
		163: uint16(146),
		164: uint16(146),
		165: uint16(146),
		166: uint16(146),
		167: uint16(146),
		168: uint16(146),
		169: uint16(146),
		170: uint16(146),
		171: uint16(146),
		172: uint16(146),
		173: uint16(146),
		174: uint16(146),
		175: uint16(146),
		176: uint16(146),
		177: uint16(146),
		178: uint16(146),
		179: uint16(146),
		180: uint16(146),
		181: uint16(146),
		182: uint16(146),
		183: uint16(146),
		184: uint16(146),
		185: uint16(146),
		186: uint16(146),
		187: uint16(146),
		188: uint16(146),
		189: uint16(146),
		190: uint16(146),
		191: uint16(146),
		192: uint16(146),
		193: uint16(146),
		194: uint16(146),
		195: uint16(146),
		196: uint16(146),
		197: uint16(146),
		198: uint16(146),
		199: uint16(146),
		200: uint16(144),
		201: uint16(144),
		202: uint16(144),
		203: uint16(144),
		204: uint16(144),
		205: uint16(144),
		206: uint16(144),
		207: uint16(144),
		208: uint16(144),
		209: uint16(144),
		210: uint16(144),
		211: uint16(144),
		212: uint16(144),
		213: uint16(144),
		214: uint16(144),
		215: uint16(144),
		216: uint16(144),
		217: uint16(144),
		218: uint16(144),
		219: uint16(144),
		220: uint16(144),
		221: uint16(144),
		222: uint16(144),
		223: uint16(144),
		224: uint16(144),
		225: uint16(144),
		226: uint16(144),
		227: uint16(144),
		228: uint16(144),
		229: uint16(144),
		230: uint16(144),
		231: uint16(144),
		232: uint16(144),
		233: uint16(144),
		234: uint16(144),
		235: uint16(144),
		236: uint16(144),
		237: uint16(144),
		238: uint16(144),
		239: uint16(144),
		240: uint16(144),
		241: uint16(144),
		242: uint16(144),
		243: uint16(144),
		244: uint16(144),
		245: uint16(144),
		246: uint16(144),
		247: uint16(144),
		248: uint16(144),
		249: uint16(144),
		250: uint16(144),
		251: uint16(144),
		252: uint16(144),
		253: uint16(144),
		254: uint16(144),
		255: uint16(144),
		256: uint16(144),
		257: uint16(144),
		258: uint16(144),
		259: uint16(144),
		260: uint16(144),
		261: uint16(144),
		262: uint16(144),
		263: uint16(144),
		264: uint16(144),
		265: uint16(144),
		266: uint16(144),
		268: uint16(144),
		271: uint16(144),
		272: uint16(144),
		273: uint16(144),
		274: uint16(144),
		275: uint16(144),
		276: uint16(144),
		277: uint16(144),
		278: uint16(144),
		279: uint16(144),
		280: uint16(144),
		281: uint16(144),
		282: uint16(144),
		284: uint16(144),
		285: uint16(3),
	},
	22: {
		0:   uint16(106),
		1:   uint16(108),
		2:   uint16(106),
		5:   uint16(106),
		6:   uint16(106),
		8:   uint16(106),
		9:   uint16(108),
		10:  uint16(108),
		11:  uint16(108),
		12:  uint16(108),
		13:  uint16(108),
		14:  uint16(108),
		15:  uint16(108),
		16:  uint16(108),
		17:  uint16(108),
		18:  uint16(108),
		19:  uint16(108),
		20:  uint16(108),
		21:  uint16(108),
		22:  uint16(108),
		23:  uint16(108),
		24:  uint16(108),
		25:  uint16(108),
		26:  uint16(108),
		27:  uint16(108),
		28:  uint16(108),
		29:  uint16(108),
		30:  uint16(108),
		31:  uint16(108),
		32:  uint16(108),
		33:  uint16(108),
		34:  uint16(108),
		35:  uint16(108),
		36:  uint16(108),
		37:  uint16(108),
		38:  uint16(108),
		39:  uint16(108),
		40:  uint16(106),
		41:  uint16(108),
		42:  uint16(108),
		43:  uint16(108),
		44:  uint16(108),
		45:  uint16(108),
		46:  uint16(108),
		47:  uint16(108),
		48:  uint16(108),
		49:  uint16(108),
		50:  uint16(108),
		51:  uint16(108),
		52:  uint16(108),
		53:  uint16(108),
		54:  uint16(108),
		55:  uint16(108),
		56:  uint16(108),
		57:  uint16(108),
		58:  uint16(108),
		59:  uint16(108),
		60:  uint16(108),
		61:  uint16(108),
		62:  uint16(108),
		63:  uint16(108),
		64:  uint16(108),
		65:  uint16(108),
		66:  uint16(108),
		67:  uint16(108),
		68:  uint16(108),
		69:  uint16(108),
		70:  uint16(108),
		71:  uint16(108),
		72:  uint16(106),
		73:  uint16(106),
		74:  uint16(106),
		75:  uint16(106),
		76:  uint16(106),
		77:  uint16(106),
		78:  uint16(106),
		79:  uint16(106),
		80:  uint16(106),
		81:  uint16(106),
		82:  uint16(106),
		83:  uint16(106),
		84:  uint16(106),
		85:  uint16(106),
		86:  uint16(106),
		87:  uint16(106),
		88:  uint16(106),
		89:  uint16(106),
		90:  uint16(106),
		91:  uint16(106),
		92:  uint16(106),
		93:  uint16(106),
		94:  uint16(106),
		95:  uint16(106),
		96:  uint16(106),
		97:  uint16(106),
		98:  uint16(106),
		99:  uint16(106),
		100: uint16(106),
		101: uint16(106),
		102: uint16(106),
		103: uint16(106),
		104: uint16(106),
		105: uint16(106),
		106: uint16(106),
		107: uint16(106),
		108: uint16(106),
		109: uint16(106),
		110: uint16(106),
		111: uint16(106),
		112: uint16(106),
		113: uint16(106),
		114: uint16(106),
		115: uint16(106),
		116: uint16(106),
		117: uint16(106),
		118: uint16(106),
		119: uint16(106),
		120: uint16(106),
		121: uint16(106),
		122: uint16(106),
		123: uint16(106),
		124: uint16(106),
		125: uint16(106),
		126: uint16(106),
		127: uint16(106),
		128: uint16(106),
		129: uint16(106),
		130: uint16(106),
		131: uint16(106),
		132: uint16(106),
		133: uint16(106),
		134: uint16(106),
		135: uint16(106),
		136: uint16(108),
		137: uint16(108),
		138: uint16(108),
		139: uint16(108),
		140: uint16(108),
		141: uint16(108),
		142: uint16(108),
		143: uint16(108),
		144: uint16(108),
		145: uint16(108),
		146: uint16(108),
		147: uint16(108),
		148: uint16(108),
		149: uint16(108),
		150: uint16(108),
		151: uint16(108),
		152: uint16(108),
		153: uint16(108),
		154: uint16(108),
		155: uint16(108),
		156: uint16(108),
		157: uint16(108),
		158: uint16(108),
		159: uint16(108),
		160: uint16(108),
		161: uint16(108),
		162: uint16(108),
		163: uint16(108),
		164: uint16(108),
		165: uint16(108),
		166: uint16(108),
		167: uint16(108),
		168: uint16(108),
		169: uint16(108),
		170: uint16(108),
		171: uint16(108),
		172: uint16(108),
		173: uint16(108),
		174: uint16(108),
		175: uint16(108),
		176: uint16(108),
		177: uint16(108),
		178: uint16(108),
		179: uint16(108),
		180: uint16(108),
		181: uint16(108),
		182: uint16(108),
		183: uint16(108),
		184: uint16(108),
		185: uint16(108),
		186: uint16(108),
		187: uint16(108),
		188: uint16(108),
		189: uint16(108),
		190: uint16(108),
		191: uint16(108),
		192: uint16(108),
		193: uint16(108),
		194: uint16(108),
		195: uint16(108),
		196: uint16(108),
		197: uint16(108),
		198: uint16(108),
		199: uint16(108),
		200: uint16(106),
		201: uint16(106),
		202: uint16(106),
		203: uint16(106),
		204: uint16(106),
		205: uint16(106),
		206: uint16(106),
		207: uint16(106),
		208: uint16(106),
		209: uint16(106),
		210: uint16(106),
		211: uint16(106),
		212: uint16(106),
		213: uint16(106),
		214: uint16(106),
		215: uint16(106),
		216: uint16(106),
		217: uint16(106),
		218: uint16(106),
		219: uint16(106),
		220: uint16(106),
		221: uint16(106),
		222: uint16(106),
		223: uint16(106),
		224: uint16(106),
		225: uint16(106),
		226: uint16(106),
		227: uint16(106),
		228: uint16(106),
		229: uint16(106),
		230: uint16(106),
		231: uint16(106),
		232: uint16(106),
		233: uint16(106),
		234: uint16(106),
		235: uint16(106),
		236: uint16(106),
		237: uint16(106),
		238: uint16(106),
		239: uint16(106),
		240: uint16(106),
		241: uint16(106),
		242: uint16(106),
		243: uint16(106),
		244: uint16(106),
		245: uint16(106),
		246: uint16(106),
		247: uint16(106),
		248: uint16(106),
		249: uint16(106),
		250: uint16(106),
		251: uint16(106),
		252: uint16(106),
		253: uint16(106),
		254: uint16(106),
		255: uint16(106),
		256: uint16(106),
		257: uint16(106),
		258: uint16(106),
		259: uint16(106),
		260: uint16(106),
		261: uint16(106),
		262: uint16(106),
		263: uint16(106),
		264: uint16(106),
		265: uint16(106),
		266: uint16(106),
		268: uint16(106),
		271: uint16(106),
		272: uint16(106),
		273: uint16(106),
		274: uint16(106),
		275: uint16(106),
		276: uint16(106),
		277: uint16(106),
		278: uint16(106),
		279: uint16(106),
		280: uint16(106),
		281: uint16(106),
		282: uint16(106),
		284: uint16(106),
		285: uint16(3),
	},
	23: {
		0:   uint16(148),
		1:   uint16(150),
		2:   uint16(148),
		6:   uint16(148),
		8:   uint16(148),
		9:   uint16(150),
		10:  uint16(150),
		11:  uint16(150),
		12:  uint16(150),
		13:  uint16(150),
		14:  uint16(150),
		15:  uint16(150),
		16:  uint16(150),
		17:  uint16(150),
		18:  uint16(150),
		19:  uint16(150),
		20:  uint16(150),
		21:  uint16(150),
		22:  uint16(150),
		23:  uint16(150),
		24:  uint16(150),
		25:  uint16(150),
		26:  uint16(150),
		27:  uint16(150),
		28:  uint16(150),
		29:  uint16(150),
		30:  uint16(150),
		31:  uint16(150),
		32:  uint16(150),
		33:  uint16(150),
		34:  uint16(150),
		35:  uint16(150),
		36:  uint16(150),
		37:  uint16(150),
		38:  uint16(150),
		39:  uint16(150),
		40:  uint16(148),
		41:  uint16(150),
		42:  uint16(150),
		43:  uint16(150),
		44:  uint16(150),
		45:  uint16(150),
		46:  uint16(150),
		47:  uint16(150),
		48:  uint16(150),
		49:  uint16(150),
		50:  uint16(150),
		51:  uint16(150),
		52:  uint16(150),
		53:  uint16(150),
		54:  uint16(150),
		55:  uint16(150),
		56:  uint16(150),
		57:  uint16(150),
		58:  uint16(150),
		59:  uint16(150),
		60:  uint16(150),
		61:  uint16(150),
		62:  uint16(150),
		63:  uint16(150),
		64:  uint16(150),
		65:  uint16(150),
		66:  uint16(150),
		67:  uint16(150),
		68:  uint16(150),
		69:  uint16(150),
		70:  uint16(150),
		71:  uint16(150),
		72:  uint16(148),
		73:  uint16(148),
		74:  uint16(148),
		75:  uint16(148),
		76:  uint16(148),
		77:  uint16(148),
		78:  uint16(148),
		79:  uint16(148),
		80:  uint16(148),
		81:  uint16(148),
		82:  uint16(148),
		83:  uint16(148),
		84:  uint16(148),
		85:  uint16(148),
		86:  uint16(148),
		87:  uint16(148),
		88:  uint16(148),
		89:  uint16(148),
		90:  uint16(148),
		91:  uint16(148),
		92:  uint16(148),
		93:  uint16(148),
		94:  uint16(148),
		95:  uint16(148),
		96:  uint16(148),
		97:  uint16(148),
		98:  uint16(148),
		99:  uint16(148),
		100: uint16(148),
		101: uint16(148),
		102: uint16(148),
		103: uint16(148),
		104: uint16(148),
		105: uint16(148),
		106: uint16(148),
		107: uint16(148),
		108: uint16(148),
		109: uint16(148),
		110: uint16(148),
		111: uint16(148),
		112: uint16(148),
		113: uint16(148),
		114: uint16(148),
		115: uint16(148),
		116: uint16(148),
		117: uint16(148),
		118: uint16(148),
		119: uint16(148),
		120: uint16(148),
		121: uint16(148),
		122: uint16(148),
		123: uint16(148),
		124: uint16(148),
		125: uint16(148),
		126: uint16(148),
		127: uint16(148),
		128: uint16(148),
		129: uint16(148),
		130: uint16(148),
		131: uint16(148),
		132: uint16(148),
		133: uint16(148),
		134: uint16(148),
		135: uint16(148),
		136: uint16(150),
		137: uint16(150),
		138: uint16(150),
		139: uint16(150),
		140: uint16(150),
		141: uint16(150),
		142: uint16(150),
		143: uint16(150),
		144: uint16(150),
		145: uint16(150),
		146: uint16(150),
		147: uint16(150),
		148: uint16(150),
		149: uint16(150),
		150: uint16(150),
		151: uint16(150),
		152: uint16(150),
		153: uint16(150),
		154: uint16(150),
		155: uint16(150),
		156: uint16(150),
		157: uint16(150),
		158: uint16(150),
		159: uint16(150),
		160: uint16(150),
		161: uint16(150),
		162: uint16(150),
		163: uint16(150),
		164: uint16(150),
		165: uint16(150),
		166: uint16(150),
		167: uint16(150),
		168: uint16(150),
		169: uint16(150),
		170: uint16(150),
		171: uint16(150),
		172: uint16(150),
		173: uint16(150),
		174: uint16(150),
		175: uint16(150),
		176: uint16(150),
		177: uint16(150),
		178: uint16(150),
		179: uint16(150),
		180: uint16(150),
		181: uint16(150),
		182: uint16(150),
		183: uint16(150),
		184: uint16(150),
		185: uint16(150),
		186: uint16(150),
		187: uint16(150),
		188: uint16(150),
		189: uint16(150),
		190: uint16(150),
		191: uint16(150),
		192: uint16(150),
		193: uint16(150),
		194: uint16(150),
		195: uint16(150),
		196: uint16(150),
		197: uint16(150),
		198: uint16(150),
		199: uint16(150),
		200: uint16(148),
		201: uint16(148),
		202: uint16(148),
		203: uint16(148),
		204: uint16(148),
		205: uint16(148),
		206: uint16(148),
		207: uint16(148),
		208: uint16(148),
		209: uint16(148),
		210: uint16(148),
		211: uint16(148),
		212: uint16(148),
		213: uint16(148),
		214: uint16(148),
		215: uint16(148),
		216: uint16(148),
		217: uint16(148),
		218: uint16(148),
		219: uint16(148),
		220: uint16(148),
		221: uint16(148),
		222: uint16(148),
		223: uint16(148),
		224: uint16(148),
		225: uint16(148),
		226: uint16(148),
		227: uint16(148),
		228: uint16(148),
		229: uint16(148),
		230: uint16(148),
		231: uint16(148),
		232: uint16(148),
		233: uint16(148),
		234: uint16(148),
		235: uint16(148),
		236: uint16(148),
		237: uint16(148),
		238: uint16(148),
		239: uint16(148),
		240: uint16(148),
		241: uint16(148),
		242: uint16(148),
		243: uint16(148),
		244: uint16(148),
		245: uint16(148),
		246: uint16(148),
		247: uint16(148),
		248: uint16(148),
		249: uint16(148),
		250: uint16(148),
		251: uint16(148),
		252: uint16(148),
		253: uint16(148),
		254: uint16(148),
		255: uint16(148),
		256: uint16(148),
		257: uint16(148),
		258: uint16(148),
		259: uint16(148),
		260: uint16(148),
		261: uint16(148),
		262: uint16(148),
		263: uint16(148),
		264: uint16(148),
		265: uint16(148),
		266: uint16(148),
		268: uint16(148),
		271: uint16(148),
		272: uint16(148),
		273: uint16(148),
		274: uint16(148),
		275: uint16(148),
		276: uint16(148),
		277: uint16(148),
		278: uint16(148),
		279: uint16(148),
		280: uint16(148),
		282: uint16(148),
		284: uint16(148),
		285: uint16(3),
	},
	24: {
		0:   uint16(152),
		1:   uint16(154),
		2:   uint16(152),
		6:   uint16(152),
		8:   uint16(152),
		9:   uint16(154),
		10:  uint16(154),
		11:  uint16(154),
		12:  uint16(154),
		13:  uint16(154),
		14:  uint16(154),
		15:  uint16(154),
		16:  uint16(154),
		17:  uint16(154),
		18:  uint16(154),
		19:  uint16(154),
		20:  uint16(154),
		21:  uint16(154),
		22:  uint16(154),
		23:  uint16(154),
		24:  uint16(154),
		25:  uint16(154),
		26:  uint16(154),
		27:  uint16(154),
		28:  uint16(154),
		29:  uint16(154),
		30:  uint16(154),
		31:  uint16(154),
		32:  uint16(154),
		33:  uint16(154),
		34:  uint16(154),
		35:  uint16(154),
		36:  uint16(154),
		37:  uint16(154),
		38:  uint16(154),
		39:  uint16(154),
		40:  uint16(152),
		41:  uint16(154),
		42:  uint16(154),
		43:  uint16(154),
		44:  uint16(154),
		45:  uint16(154),
		46:  uint16(154),
		47:  uint16(154),
		48:  uint16(154),
		49:  uint16(154),
		50:  uint16(154),
		51:  uint16(154),
		52:  uint16(154),
		53:  uint16(154),
		54:  uint16(154),
		55:  uint16(154),
		56:  uint16(154),
		57:  uint16(154),
		58:  uint16(154),
		59:  uint16(154),
		60:  uint16(154),
		61:  uint16(154),
		62:  uint16(154),
		63:  uint16(154),
		64:  uint16(154),
		65:  uint16(154),
		66:  uint16(154),
		67:  uint16(154),
		68:  uint16(154),
		69:  uint16(154),
		70:  uint16(154),
		71:  uint16(154),
		72:  uint16(152),
		73:  uint16(152),
		74:  uint16(152),
		75:  uint16(152),
		76:  uint16(152),
		77:  uint16(152),
		78:  uint16(152),
		79:  uint16(152),
		80:  uint16(152),
		81:  uint16(152),
		82:  uint16(152),
		83:  uint16(152),
		84:  uint16(152),
		85:  uint16(152),
		86:  uint16(152),
		87:  uint16(152),
		88:  uint16(152),
		89:  uint16(152),
		90:  uint16(152),
		91:  uint16(152),
		92:  uint16(152),
		93:  uint16(152),
		94:  uint16(152),
		95:  uint16(152),
		96:  uint16(152),
		97:  uint16(152),
		98:  uint16(152),
		99:  uint16(152),
		100: uint16(152),
		101: uint16(152),
		102: uint16(152),
		103: uint16(152),
		104: uint16(152),
		105: uint16(152),
		106: uint16(152),
		107: uint16(152),
		108: uint16(152),
		109: uint16(152),
		110: uint16(152),
		111: uint16(152),
		112: uint16(152),
		113: uint16(152),
		114: uint16(152),
		115: uint16(152),
		116: uint16(152),
		117: uint16(152),
		118: uint16(152),
		119: uint16(152),
		120: uint16(152),
		121: uint16(152),
		122: uint16(152),
		123: uint16(152),
		124: uint16(152),
		125: uint16(152),
		126: uint16(152),
		127: uint16(152),
		128: uint16(152),
		129: uint16(152),
		130: uint16(152),
		131: uint16(152),
		132: uint16(152),
		133: uint16(152),
		134: uint16(152),
		135: uint16(152),
		136: uint16(154),
		137: uint16(154),
		138: uint16(154),
		139: uint16(154),
		140: uint16(154),
		141: uint16(154),
		142: uint16(154),
		143: uint16(154),
		144: uint16(154),
		145: uint16(154),
		146: uint16(154),
		147: uint16(154),
		148: uint16(154),
		149: uint16(154),
		150: uint16(154),
		151: uint16(154),
		152: uint16(154),
		153: uint16(154),
		154: uint16(154),
		155: uint16(154),
		156: uint16(154),
		157: uint16(154),
		158: uint16(154),
		159: uint16(154),
		160: uint16(154),
		161: uint16(154),
		162: uint16(154),
		163: uint16(154),
		164: uint16(154),
		165: uint16(154),
		166: uint16(154),
		167: uint16(154),
		168: uint16(154),
		169: uint16(154),
		170: uint16(154),
		171: uint16(154),
		172: uint16(154),
		173: uint16(154),
		174: uint16(154),
		175: uint16(154),
		176: uint16(154),
		177: uint16(154),
		178: uint16(154),
		179: uint16(154),
		180: uint16(154),
		181: uint16(154),
		182: uint16(154),
		183: uint16(154),
		184: uint16(154),
		185: uint16(154),
		186: uint16(154),
		187: uint16(154),
		188: uint16(154),
		189: uint16(154),
		190: uint16(154),
		191: uint16(154),
		192: uint16(154),
		193: uint16(154),
		194: uint16(154),
		195: uint16(154),
		196: uint16(154),
		197: uint16(154),
		198: uint16(154),
		199: uint16(154),
		200: uint16(152),
		201: uint16(152),
		202: uint16(152),
		203: uint16(152),
		204: uint16(152),
		205: uint16(152),
		206: uint16(152),
		207: uint16(152),
		208: uint16(152),
		209: uint16(152),
		210: uint16(152),
		211: uint16(152),
		212: uint16(152),
		213: uint16(152),
		214: uint16(152),
		215: uint16(152),
		216: uint16(152),
		217: uint16(152),
		218: uint16(152),
		219: uint16(152),
		220: uint16(152),
		221: uint16(152),
		222: uint16(152),
		223: uint16(152),
		224: uint16(152),
		225: uint16(152),
		226: uint16(152),
		227: uint16(152),
		228: uint16(152),
		229: uint16(152),
		230: uint16(152),
		231: uint16(152),
		232: uint16(152),
		233: uint16(152),
		234: uint16(152),
		235: uint16(152),
		236: uint16(152),
		237: uint16(152),
		238: uint16(152),
		239: uint16(152),
		240: uint16(152),
		241: uint16(152),
		242: uint16(152),
		243: uint16(152),
		244: uint16(152),
		245: uint16(152),
		246: uint16(152),
		247: uint16(152),
		248: uint16(152),
		249: uint16(152),
		250: uint16(152),
		251: uint16(152),
		252: uint16(152),
		253: uint16(152),
		254: uint16(152),
		255: uint16(152),
		256: uint16(152),
		257: uint16(152),
		258: uint16(152),
		259: uint16(152),
		260: uint16(152),
		261: uint16(152),
		262: uint16(152),
		263: uint16(152),
		264: uint16(152),
		265: uint16(152),
		266: uint16(152),
		268: uint16(152),
		271: uint16(152),
		272: uint16(152),
		273: uint16(152),
		274: uint16(152),
		275: uint16(152),
		276: uint16(152),
		277: uint16(152),
		278: uint16(152),
		279: uint16(152),
		280: uint16(152),
		282: uint16(152),
		284: uint16(152),
		285: uint16(3),
	},
}

var ts_small_parse_table = [370]uint16_t{
	0:   uint16(10),
	1:   uint16(3),
	2:   uint16(1),
	3:   uint16(sym_comment),
	4:   uint16(76),
	5:   uint16(1),
	6:   uint16(anon_sym_AMP),
	7:   uint16(156),
	8:   uint16(1),
	9:   uint16(sym_identifier),
	10:  uint16(158),
	11:  uint16(1),
	12:  uint16(aux_sym_rune_token1),
	13:  uint16(14),
	14:  uint16(1),
	15:  uint16(sym_sublabel_reference),
	16:  uint16(68),
	17:  uint16(2),
	18:  uint16(anon_sym_COMMA),
	19:  uint16(anon_sym__),
	20:  uint16(70),
	21:  uint16(2),
	22:  uint16(anon_sym_DOT),
	23:  uint16(anon_sym_DASH),
	24:  uint16(72),
	25:  uint16(2),
	26:  uint16(anon_sym_SEMI),
	27:  uint16(anon_sym_EQ),
	28:  uint16(74),
	29:  uint16(2),
	30:  uint16(anon_sym_BANG),
	31:  uint16(anon_sym_QMARK),
	32:  uint16(26),
	33:  uint16(2),
	34:  uint16(sym_rune_char),
	35:  uint16(aux_sym_rune_repeat1),
	36:  uint16(10),
	37:  uint16(3),
	38:  uint16(1),
	39:  uint16(sym_comment),
	40:  uint16(76),
	41:  uint16(1),
	42:  uint16(anon_sym_AMP),
	43:  uint16(160),
	44:  uint16(1),
	45:  uint16(sym_identifier),
	46:  uint16(162),
	47:  uint16(1),
	48:  uint16(aux_sym_rune_token1),
	49:  uint16(22),
	50:  uint16(1),
	51:  uint16(sym_sublabel_reference),
	52:  uint16(68),
	53:  uint16(2),
	54:  uint16(anon_sym_COMMA),
	55:  uint16(anon_sym__),
	56:  uint16(70),
	57:  uint16(2),
	58:  uint16(anon_sym_DOT),
	59:  uint16(anon_sym_DASH),
	60:  uint16(72),
	61:  uint16(2),
	62:  uint16(anon_sym_SEMI),
	63:  uint16(anon_sym_EQ),
	64:  uint16(74),
	65:  uint16(2),
	66:  uint16(anon_sym_BANG),
	67:  uint16(anon_sym_QMARK),
	68:  uint16(27),
	69:  uint16(2),
	70:  uint16(sym_rune_char),
	71:  uint16(aux_sym_rune_repeat1),
	72:  uint16(9),
	73:  uint16(3),
	74:  uint16(1),
	75:  uint16(sym_comment),
	76:  uint16(164),
	77:  uint16(1),
	78:  uint16(sym_identifier),
	79:  uint16(166),
	80:  uint16(1),
	81:  uint16(aux_sym_rune_token1),
	82:  uint16(180),
	83:  uint16(1),
	84:  uint16(anon_sym_AMP),
	85:  uint16(168),
	86:  uint16(2),
	87:  uint16(anon_sym_COMMA),
	88:  uint16(anon_sym__),
	89:  uint16(171),
	90:  uint16(2),
	91:  uint16(anon_sym_DOT),
	92:  uint16(anon_sym_DASH),
	93:  uint16(174),
	94:  uint16(2),
	95:  uint16(anon_sym_SEMI),
	96:  uint16(anon_sym_EQ),
	97:  uint16(177),
	98:  uint16(2),
	99:  uint16(anon_sym_BANG),
	100: uint16(anon_sym_QMARK),
	101: uint16(27),
	102: uint16(2),
	103: uint16(sym_rune_char),
	104: uint16(aux_sym_rune_repeat1),
	105: uint16(9),
	106: uint16(3),
	107: uint16(1),
	108: uint16(sym_comment),
	109: uint16(183),
	110: uint16(1),
	112: uint16(185),
	113: uint16(1),
	114: uint16(anon_sym_PERCENT),
	115: uint16(188),
	116: uint16(1),
	117: uint16(anon_sym_TILDE),
	118: uint16(191),
	119: uint16(1),
	120: uint16(anon_sym_PIPE),
	121: uint16(194),
	122: uint16(1),
	123: uint16(anon_sym_AT),
	124: uint16(3),
	125: uint16(1),
	126: uint16(sym_absolute_pad_operation),
	127: uint16(4),
	128: uint16(1),
	129: uint16(sym_label),
	130: uint16(28),
	131: uint16(5),
	132: uint16(sym_memory_execution),
	133: uint16(sym_subroutine),
	134: uint16(sym_macro),
	135: uint16(sym_include),
	136: uint16(aux_sym_program_repeat1),
	137: uint16(9),
	138: uint16(3),
	139: uint16(1),
	140: uint16(sym_comment),
	141: uint16(7),
	142: uint16(1),
	143: uint16(anon_sym_PERCENT),
	144: uint16(9),
	145: uint16(1),
	146: uint16(anon_sym_TILDE),
	147: uint16(11),
	148: uint16(1),
	149: uint16(anon_sym_PIPE),
	150: uint16(13),
	151: uint16(1),
	152: uint16(anon_sym_AT),
	153: uint16(197),
	154: uint16(1),
	156: uint16(3),
	157: uint16(1),
	158: uint16(sym_absolute_pad_operation),
	159: uint16(4),
	160: uint16(1),
	161: uint16(sym_label),
	162: uint16(28),
	163: uint16(5),
	164: uint16(sym_memory_execution),
	165: uint16(sym_subroutine),
	166: uint16(sym_macro),
	167: uint16(sym_include),
	168: uint16(aux_sym_program_repeat1),
	169: uint16(3),
	170: uint16(3),
	171: uint16(1),
	172: uint16(sym_comment),
	173: uint16(199),
	174: uint16(1),
	175: uint16(sym_identifier),
	176: uint16(201),
	177: uint16(10),
	178: uint16(aux_sym_rune_token1),
	179: uint16(anon_sym_COMMA),
	180: uint16(anon_sym__),
	181: uint16(anon_sym_DOT),
	182: uint16(anon_sym_DASH),
	183: uint16(anon_sym_SEMI),
	184: uint16(anon_sym_EQ),
	185: uint16(anon_sym_BANG),
	186: uint16(anon_sym_QMARK),
	187: uint16(anon_sym_AMP),
	188: uint16(3),
	189: uint16(3),
	190: uint16(1),
	191: uint16(sym_comment),
	192: uint16(203),
	193: uint16(1),
	194: uint16(sym_identifier),
	195: uint16(205),
	196: uint16(10),
	197: uint16(aux_sym_rune_token1),
	198: uint16(anon_sym_COMMA),
	199: uint16(anon_sym__),
	200: uint16(anon_sym_DOT),
	201: uint16(anon_sym_DASH),
	202: uint16(anon_sym_SEMI),
	203: uint16(anon_sym_EQ),
	204: uint16(anon_sym_BANG),
	205: uint16(anon_sym_QMARK),
	206: uint16(anon_sym_AMP),
	207: uint16(3),
	208: uint16(3),
	209: uint16(1),
	210: uint16(sym_comment),
	211: uint16(207),
	212: uint16(1),
	213: uint16(sym_identifier),
	214: uint16(209),
	215: uint16(10),
	216: uint16(aux_sym_rune_token1),
	217: uint16(anon_sym_COMMA),
	218: uint16(anon_sym__),
	219: uint16(anon_sym_DOT),
	220: uint16(anon_sym_DASH),
	221: uint16(anon_sym_SEMI),
	222: uint16(anon_sym_EQ),
	223: uint16(anon_sym_BANG),
	224: uint16(anon_sym_QMARK),
	225: uint16(anon_sym_AMP),
	226: uint16(3),
	227: uint16(3),
	228: uint16(1),
	229: uint16(sym_comment),
	230: uint16(211),
	231: uint16(1),
	232: uint16(sym_identifier),
	233: uint16(213),
	234: uint16(10),
	235: uint16(aux_sym_rune_token1),
	236: uint16(anon_sym_COMMA),
	237: uint16(anon_sym__),
	238: uint16(anon_sym_DOT),
	239: uint16(anon_sym_DASH),
	240: uint16(anon_sym_SEMI),
	241: uint16(anon_sym_EQ),
	242: uint16(anon_sym_BANG),
	243: uint16(anon_sym_QMARK),
	244: uint16(anon_sym_AMP),
	245: uint16(3),
	246: uint16(3),
	247: uint16(1),
	248: uint16(sym_comment),
	249: uint16(215),
	250: uint16(1),
	251: uint16(sym_identifier),
	252: uint16(217),
	253: uint16(10),
	254: uint16(aux_sym_rune_token1),
	255: uint16(anon_sym_COMMA),
	256: uint16(anon_sym__),
	257: uint16(anon_sym_DOT),
	258: uint16(anon_sym_DASH),
	259: uint16(anon_sym_SEMI),
	260: uint16(anon_sym_EQ),
	261: uint16(anon_sym_BANG),
	262: uint16(anon_sym_QMARK),
	263: uint16(anon_sym_AMP),
	264: uint16(2),
	265: uint16(3),
	266: uint16(1),
	267: uint16(sym_comment),
	268: uint16(219),
	269: uint16(5),
	271: uint16(anon_sym_PERCENT),
	272: uint16(anon_sym_TILDE),
	273: uint16(anon_sym_PIPE),
	274: uint16(anon_sym_AT),
	275: uint16(2),
	276: uint16(3),
	277: uint16(1),
	278: uint16(sym_comment),
	279: uint16(221),
	280: uint16(5),
	282: uint16(anon_sym_PERCENT),
	283: uint16(anon_sym_TILDE),
	284: uint16(anon_sym_PIPE),
	285: uint16(anon_sym_AT),
	286: uint16(2),
	287: uint16(3),
	288: uint16(1),
	289: uint16(sym_comment),
	290: uint16(223),
	291: uint16(5),
	293: uint16(anon_sym_PERCENT),
	294: uint16(anon_sym_TILDE),
	295: uint16(anon_sym_PIPE),
	296: uint16(anon_sym_AT),
	297: uint16(3),
	298: uint16(3),
	299: uint16(1),
	300: uint16(sym_comment),
	301: uint16(225),
	302: uint16(1),
	303: uint16(sym_identifier),
	304: uint16(227),
	305: uint16(1),
	306: uint16(aux_sym_macro_token1),
	307: uint16(2),
	308: uint16(3),
	309: uint16(1),
	310: uint16(sym_comment),
	311: uint16(229),
	312: uint16(1),
	313: uint16(aux_sym_raw_ascii_token1),
	314: uint16(2),
	315: uint16(3),
	316: uint16(1),
	317: uint16(sym_comment),
	318: uint16(231),
	319: uint16(1),
	320: uint16(anon_sym_LBRACE),
	321: uint16(2),
	322: uint16(3),
	323: uint16(1),
	324: uint16(sym_comment),
	325: uint16(233),
	326: uint16(1),
	327: uint16(sym_number),
	328: uint16(2),
	329: uint16(3),
	330: uint16(1),
	331: uint16(sym_comment),
	332: uint16(235),
	333: uint16(1),
	335: uint16(2),
	336: uint16(3),
	337: uint16(1),
	338: uint16(sym_comment),
	339: uint16(237),
	340: uint16(1),
	341: uint16(sym_identifier),
	342: uint16(2),
	343: uint16(3),
	344: uint16(1),
	345: uint16(sym_comment),
	346: uint16(239),
	347: uint16(1),
	348: uint16(sym_identifier),
	349: uint16(2),
	350: uint16(3),
	351: uint16(1),
	352: uint16(sym_comment),
	353: uint16(241),
	354: uint16(1),
	355: uint16(aux_sym_include_token1),
	356: uint16(2),
	357: uint16(3),
	358: uint16(1),
	359: uint16(sym_comment),
	360: uint16(243),
	361: uint16(1),
	362: uint16(sym_number),
	363: uint16(2),
	364: uint16(3),
	365: uint16(1),
	366: uint16(sym_comment),
	367: uint16(245),
	368: uint16(1),
	369: uint16(aux_sym_hex_literal_token1),
}

var ts_small_parse_table_map = [23]uint32_t{
	1:  uint32(36),
	2:  uint32(72),
	3:  uint32(105),
	4:  uint32(137),
	5:  uint32(169),
	6:  uint32(188),
	7:  uint32(207),
	8:  uint32(226),
	9:  uint32(245),
	10: uint32(264),
	11: uint32(275),
	12: uint32(286),
	13: uint32(297),
	14: uint32(307),
	15: uint32(314),
	16: uint32(321),
	17: uint32(328),
	18: uint32(335),
	19: uint32(342),
	20: uint32(349),
	21: uint32(356),
	22: uint32(363),
}

var ts_parse_actions = [247]TSParseActionEntry{
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(41)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
	})))),
	19: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	20: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	21: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(15)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
	})))),
	25: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(15)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	26: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	27: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(46)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	30: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
	})))),
	31: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(47)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	32: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	33: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(34)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(30)),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		Fsymbol:      uint16(aux_sym_memory_execution_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_memory_execution),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(34)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_subroutine),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_subroutine),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_memory_execution),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(17)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	107: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_rune),
		Fproduction_id: uint16(6),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_rune),
		Fproduction_id: uint16(6),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_rune),
		Fproduction_id: uint16(6),
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
		Fcount: uint8(1),
	}})),
	115: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_rune),
		Fproduction_id: uint16(6),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__non_toplevel_statement),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__non_toplevel_statement),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_opcode),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_opcode),
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
		Fsymbol:      uint16(sym_sublabel_reference),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_sublabel_reference),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_brackets),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_brackets),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_raw_ascii),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_raw_ascii),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_brackets),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_brackets),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_hex_literal),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_hex_literal),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_relative_pad_operation),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_relative_pad_operation),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_absolute_pad_operation),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_absolute_pad_operation),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_label),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_label),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
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
		Fsymbol:      uint16(aux_sym_rune_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rune_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rune_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(34)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	172: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_rune_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(33)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_rune_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		Fsymbol:      uint16(aux_sym_rune_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fsymbol:      uint16(aux_sym_rune_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(30)),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(38)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Fcount:    uint8(2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(41)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	198: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	200: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_rune_char),
		Fproduction_id: uint16(5),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_rune_char),
		Fproduction_id: uint16(5),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_rune_char),
		Fproduction_id: uint16(4),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	206: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_rune_char),
		Fproduction_id: uint16(4),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_rune_char),
		Fproduction_id: uint16(3),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_rune_char),
		Fproduction_id: uint16(3),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_rune_char),
		Fproduction_id: uint16(2),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_rune_char),
		Fproduction_id: uint16(2),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_rune_char),
		Fproduction_id: uint16(1),
	})))),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_rune_char),
		Fproduction_id: uint16(1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_macro),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_include),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_macro),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	230: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	236: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	237: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	242: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	243: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	244: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	245: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
	}})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token_comment = 0

var ts_external_scanner_symbol_map = [1]TSSymbol{
	0: uint16(sym_comment),
}

var ts_external_scanner_states = [2][1]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_uxntal(tls *libc.TLS) (r uintptr) {
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
	Fprimary_state_ids: uintptr(unsafe.Pointer(&ts_primary_state_ids)),
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 92)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 96)) = __ccgo_fp(ts_lex_keywords)
	*(*uintptr)(unsafe.Add(p, 112)) = __ccgo_fp(tree_sitter_uxntal_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 116)) = __ccgo_fp(tree_sitter_uxntal_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 120)) = __ccgo_fp(tree_sitter_uxntal_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 124)) = __ccgo_fp(tree_sitter_uxntal_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 128)) = __ccgo_fp(tree_sitter_uxntal_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00identifier\x00%\x00{\x00}\x00~\x00include_token1\x00BRK\x00INC\x00POP\x00NIP\x00SWP\x00ROT\x00DUP\x00OVR\x00EQU\x00NEQ\x00GTH\x00LTH\x00JMP\x00JCN\x00JSR\x00STH\x00LDZ\x00STZ\x00LDR\x00STR\x00LDA\x00STA\x00DEI\x00DEO\x00ADD\x00SUB\x00MUL\x00DIV\x00AND\x00ORA\x00EOR\x00SFT\x00JCI\x00INC2\x00POP2\x00NIP2\x00SWP2\x00ROT2\x00DUP2\x00OVR2\x00EQU2\x00NEQ2\x00GTH2\x00LTH2\x00JMP2\x00JCN2\x00JSR2\x00STH2\x00LDZ2\x00STZ2\x00LDR2\x00STR2\x00LDA2\x00STA2\x00DEI2\x00DEO2\x00ADD2\x00SUB2\x00MUL2\x00DIV2\x00AND2\x00ORA2\x00EOR2\x00SFT2\x00JMI\x00INCr\x00POPr\x00NIPr\x00SWPr\x00ROTr\x00DUPr\x00OVRr\x00EQUr\x00NEQr\x00GTHr\x00LTHr\x00JMPr\x00JCNr\x00JSRr\x00STHr\x00LDZr\x00STZr\x00LDRr\x00STRr\x00LDAr\x00STAr\x00DEIr\x00DEOr\x00ADDr\x00SUBr\x00MULr\x00DIVr\x00ANDr\x00ORAr\x00EORr\x00SFTr\x00JSI\x00INC2r\x00POP2r\x00NIP2r\x00SWP2r\x00ROT2r\x00DUP2r\x00OVR2r\x00EQU2r\x00NEQ2r\x00GTH2r\x00LTH2r\x00JMP2r\x00JCN2r\x00JSR2r\x00STH2r\x00LDZ2r\x00STZ2r\x00LDR2r\x00STR2r\x00LDA2r\x00STA2r\x00DEI2r\x00DEO2r\x00ADD2r\x00SUB2r\x00MUL2r\x00DIV2r\x00AND2r\x00ORA2r\x00EOR2r\x00SFT2r\x00LIT\x00INCk\x00POPk\x00NIPk\x00SWPk\x00ROTk\x00DUPk\x00OVRk\x00EQUk\x00NEQk\x00GTHk\x00LTHk\x00JMPk\x00JCNk\x00JSRk\x00STHk\x00LDZk\x00STZk\x00LDRk\x00STRk\x00LDAk\x00STAk\x00DEIk\x00DEOk\x00ADDk\x00SUBk\x00MULk\x00DIVk\x00ANDk\x00ORAk\x00EORk\x00SFTk\x00LIT2\x00INC2k\x00POP2k\x00NIP2k\x00SWP2k\x00ROT2k\x00DUP2k\x00OVR2k\x00EQU2k\x00NEQ2k\x00GTH2k\x00LTH2k\x00JMP2k\x00JCN2k\x00JSR2k\x00STH2k\x00LDZ2k\x00STZ2k\x00LDR2k\x00STR2k\x00LDA2k\x00STA2k\x00DEI2k\x00DEO2k\x00ADD2k\x00SUB2k\x00MUL2k\x00DIV2k\x00AND2k\x00ORA2k\x00EOR2k\x00SFT2k\x00LITr\x00INCkr\x00POPkr\x00NIPkr\x00SWPkr\x00ROTkr\x00DUPkr\x00OVRkr\x00EQUkr\x00NEQkr\x00GTHkr\x00LTHkr\x00JMPkr\x00JCNkr\x00JSRkr\x00STHkr\x00LDZkr\x00STZkr\x00LDRkr\x00STRkr\x00LDAkr\x00STAkr\x00DEIkr\x00DEOkr\x00ADDkr\x00SUBkr\x00MULkr\x00DIVkr\x00ANDkr\x00ORAkr\x00EORkr\x00SFTkr\x00LIT2r\x00INC2kr\x00POP2kr\x00NIP2kr\x00SWP2kr\x00ROT2kr\x00DUP2kr\x00OVR2kr\x00EQU2kr\x00NEQ2kr\x00GTH2kr\x00LTH2kr\x00JMP2kr\x00JCN2kr\x00JSR2kr\x00STH2kr\x00LDZ2kr\x00STZ2kr\x00LDR2kr\x00STR2kr\x00LDA2kr\x00STA2kr\x00DEI2kr\x00DEO2kr\x00ADD2kr\x00SUB2kr\x00MUL2kr\x00DIV2kr\x00AND2kr\x00ORA2kr\x00EOR2kr\x00SFT2kr\x00|\x00$\x00#\x00hex_lit_value\x00@\x00/\x00,\x00_\x00.\x00-\x00;\x00=\x00!\x00?\x00&\x00[\x00]\x00\"\x00raw_ascii_token1\x00number\x00comment\x00program\x00memory_execution\x00subroutine\x00_non_toplevel_statement\x00macro\x00include\x00opcode\x00absolute_pad_operation\x00relative_pad_operation\x00hex_literal\x00label\x00sublabel_reference\x00rune\x00rune_char\x00brackets\x00raw_ascii\x00program_repeat1\x00memory_execution_repeat1\x00rune_repeat1\x00absolute\x00immediate\x00relative\x00rune_start\x00sublabel\x00zero_page\x00"
